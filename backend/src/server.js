import Fastify from "fastify";
import rateLimit from "./plugins/rateLimit.js";
import caching from "./plugins/caching.js";
import jwt from "./plugins/jwt.js";
import analytics from "./plugins/analytics.js";
import apiRoutes from "./routes/apiRoutes.js";
import v1 from "./v1/index.js";
import v2 from "./v2/index.js";
import fetch from "node-fetch";
import { getNextBackend } from "./loadbalancer/loadBalancer.js";
import cors from "@fastify/cors"; // Import CORS plugin

// Create Fastify instance
const fastify = Fastify({ logger: true });

// Register plugins
fastify.register(rateLimit);
fastify.register(caching);
fastify.register(jwt);
fastify.register(analytics);

// Register routes with versioning
fastify.register(v1, { prefix: "/v1" });
fastify.register(v2, { prefix: "/v2" });

// Register the API routes
fastify.register(apiRoutes);

// Health Check Route
fastify.get("/health", async (request, reply) => {
  try {
    reply.send({ status: "ok", uptime: process.uptime() });
  } catch (err) {
    reply.code(500).send({ error: "Health check failed", details: err.message });
  }
});

// Register CORS plugin for specific routes
fastify.register(cors, {
  origin: "http://localhost:5173", // Replace with your frontend URL
  methods: ["GET", "POST", "PUT", "DELETE"],
  preflightContinue: true, // Avoids automatic handling of OPTIONS requests
});

// Load Balancer Middleware: Handle all requests and forward to the backend
fastify.all("/*", async (request, reply) => {
  const target = getNextBackend();

  try {
    const response = await fetch(target + request.url, {
      method: request.method,
      headers: {
        ...request.headers,
        "x-forwarded-for": request.ip,
      },
      body: request.method !== "GET" && request.method !== "HEAD" ? JSON.stringify(request.body) : null,
    });

    reply.code(response.status).headers(response.headers.raw()).send(await response.text());
  } catch (error) {
    fastify.log.error(`Error in load balancing: ${error.message}`);
    reply.code(500).send({ error: "Backend service unavailable" });
  }
});

// Graceful Shutdown
const gracefulShutdown = async () => {
  try {
    await fastify.close();
    console.log("Server gracefully shut down");
    process.exit(0);
  } catch (err) {
    console.error("Error during graceful shutdown", err);
    process.exit(1);
  }
};

process.on("unhandledRejection", (err) => {
  fastify.log.error(`Unhandled Rejection: ${err}`);
  process.exit(1);
});

process.on("uncaughtException", (err) => {
  fastify.log.error(`Uncaught Exception: ${err}`);
  process.exit(1);
});

const start = async () => {
  try {
    await fastify.listen({ port: 3000, host: "0.0.0.0" });
    console.log("🚀 SERVER IS RUNNING ON http://localhost:3000");
    console.log("📊 Metrics available at http://localhost:3000/metrics");
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

process.on("SIGTERM", gracefulShutdown);
process.on("SIGINT", gracefulShutdown);

start();