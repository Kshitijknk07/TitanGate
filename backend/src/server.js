import Fastify from "fastify";
import cors from "@fastify/cors";
import fetch from "node-fetch";
import rateLimit from "./plugins/rateLimit.js";
import caching from "./plugins/caching.js";
import jwt from "./plugins/jwt.js";
import analytics from "./plugins/analytics.js";
import apiRoutes from "./routes/apiRoutes.js";
import v1 from "./v1/index.js";
import v2 from "./v2/index.js";
import { getNextBackend } from "./loadbalancer/loadBalancer.js";
import errorHandler from "./plugins/errorHandler.js"; 

// Initialize Fastify instance with logging enabled
const fastify = Fastify({ logger: true });

// Register plugins
fastify.register(rateLimit);
fastify.register(caching);
fastify.register(jwt);
fastify.register(analytics);
fastify.register(errorHandler);

// Register versioned API routes
fastify.register(v1, { prefix: "/v1" });
fastify.register(v2, { prefix: "/v2" });
fastify.register(apiRoutes);

// Health check endpoint
fastify.get("/health", async (request, reply) => {
  reply.send({ status: "ok", uptime: process.uptime() });
});

// Register CORS
fastify.register(cors, {
  origin: ["http://localhost:5173", "http://localhost:3000"],
  credentials: true,
  methods: ["GET", "POST", "PUT", "DELETE", "PATCH"],
});

// Load Balancer Route
fastify.route({
  method: ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"],
  url: "/*",
  handler: async (request, reply) => {
    const target = getNextBackend();  // Get the next backend server for load balancing
    try {
      const response = await fetch(target + request.url, {
        method: request.method,
        headers: { ...request.headers, "x-forwarded-for": request.ip },
        body: request.method !== "GET" && request.method !== "HEAD" ? JSON.stringify(request.body) : null,
      });
      
      reply.code(response.status);
      Object.entries(response.headers.raw()).forEach(([key, value]) => reply.header(key, value));
      reply.send(await response.text());
    } catch (error) {
      fastify.log.error(`Load balancer error: ${error.message}`);
      reply.code(500).send({ error: "Backend service unavailable" });
    }
  },
});

// Graceful shutdown
const gracefulShutdown = async () => {
  try {
    await fastify.close();
    console.log("Server shut down gracefully");
    process.exit(0);
  } catch (err) {
    console.error("Error during shutdown", err);
    process.exit(1);
  }
};

// Handle process signals for graceful shutdown
process.on("SIGTERM", gracefulShutdown);
process.on("SIGINT", gracefulShutdown);

// Handle global errors
process.on("unhandledRejection", (err) => {
  fastify.log.error(`Unhandled Rejection: ${err}`);
  process.exit(1);
});

process.on("uncaughtException", (err) => {
  fastify.log.error(`Uncaught Exception: ${err}`);
  process.exit(1);
});

// Start server
const start = async () => {
  try {
    await fastify.listen({ port: 3000, host: "0.0.0.0" });
    console.log("🚀 Server running on http://localhost:3000");
    console.log("📊 Metrics available at http://localhost:3000/metrics");
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

start();
