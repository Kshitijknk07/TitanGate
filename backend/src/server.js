// src/server.js
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
import cors from "@fastify/cors";

// Create Fastify instance
const fastify = Fastify({ logger: true });

// Register core plugins
fastify.register(rateLimit);
fastify.register(caching);
fastify.register(jwt);
fastify.register(analytics);

// Register versioned routes
fastify.register(v1, { prefix: "/v1" });
fastify.register(v2, { prefix: "/v2" });

// Register API routes
fastify.register(apiRoutes);  

// Health Check Route
fastify.get("/health", async (request, reply) => {
  try {
    reply.send({ status: "ok", uptime: process.uptime() });
  } catch (err) {
    reply.code(500).send({ error: "Health check failed", details: err.message });
  }
});

// Register CORS plugin
fastify.register(cors, {
  origin: "http://localhost:5173", // Replace with your frontend URL
  methods: ["GET", "POST", "PUT", "DELETE", "PATCH"], // Allowed methods (OPTIONS is handled automatically)
  preflightContinue: true // Let CORS plugin handle OPTIONS
});

// Load Balancer Route
// Note: We use fastify.route with an explicit method list that excludes OPTIONS.
fastify.route({
  method: ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"],
  url: "/*",
  handler: async (request, reply) => {
    const target = getNextBackend();
    try {
      const response = await fetch(target + request.url, {
        method: request.method,
        headers: {
          ...request.headers,
          "x-forwarded-for": request.ip,
        },
        // Only attach a body for methods that support it
        body: request.method !== "GET" && request.method !== "HEAD" ? JSON.stringify(request.body) : null,
      });

      // Set response headers and send back the response body
      reply.code(response.status);
      // response.headers.raw() returns an object of header arrays; set each header accordingly
      for (const [key, value] of Object.entries(response.headers.raw())) {
        reply.header(key, value);
      }
      reply.send(await response.text());
    } catch (error) {
      fastify.log.error(`Error in load balancing: ${error.message}`);
      reply.code(500).send({ error: "Backend service unavailable" });
    }
  }
});

// Graceful Shutdown Handler
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

// Global error handling for unhandled rejections and exceptions
process.on("unhandledRejection", (err) => {
  fastify.log.error(`Unhandled Rejection: ${err}`);
  process.exit(1);
});

process.on("uncaughtException", (err) => {
  fastify.log.error(`Uncaught Exception: ${err}`);
  process.exit(1);
});

// Start the server
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

// Listen for shutdown signals
process.on("SIGTERM", gracefulShutdown);
process.on("SIGINT", gracefulShutdown);

start();
