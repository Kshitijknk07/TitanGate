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

// Create Fastify instance
const fastify = Fastify({ logger: true });

// Feature flags, ideally from environment variables
const featureFlags = {
  v2Enabled: process.env.FEATURE_V2_ENABLED === 'true' || true, // Default to true if not defined
};

// Middleware to check if a feature is enabled
fastify.addHook("onRequest", async (request, reply) => {
  if (request.url.includes("/v2") && !featureFlags.v2Enabled) {
    reply.code(404).send("Not Found");
  }
});

// Register routes with versioning
fastify.register(v1, { prefix: "/v1" }); // http://localhost:3000/v1
fastify.register(v2, { prefix: "/v2" }); // http://localhost:3000/v2

// Register plugins
fastify.register(rateLimit);
fastify.register(caching);
fastify.register(jwt);
fastify.register(analytics);

// Register the API routes
fastify.register(apiRoutes);

// Health Check Route
fastify.get("/health", async (request, reply) => {
  try {
    // Simple health check logic, you could add more checks (DB, external services)
    reply.send({ status: "ok", uptime: process.uptime() });
  } catch (err) {
    reply.code(500).send({ error: "Health check failed", details: err.message });
  }
});

// Load Balancer Middleware: Handle all requests and forward to the backend
fastify.all("/*", async (request, reply) => {
  const target = getNextBackend();

  try {
    // Forward the request to the selected backend
    const response = await fetch(target + request.url, {
      method: request.method,
      headers: {
        ...request.headers, 
        'x-forwarded-for': request.ip, // Optional: Forward real IP for better logging
      },
      body: request.method !== "GET" && request.method !== "HEAD" ? JSON.stringify(request.body) : null,
    });

    // Forward the backend response back to the client
    reply.code(response.status).headers(response.headers.raw()).send(await response.text());
  } catch (error) {
    // In case of an error (e.g., backend service unavailable)
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

// Catch unhandled rejections and terminate
process.on("unhandledRejection", (err) => {
  fastify.log.error(`Unhandled Rejection: ${err}`);
  process.exit(1);
});

// Catch uncaught exceptions and terminate
process.on("uncaughtException", (err) => {
  fastify.log.error(`Uncaught Exception: ${err}`);
  process.exit(1);
});

// Start the server
const start = async () => {
  try {
    await fastify.listen({ port: 3000, host: "0.0.0.0" }); // Listen on all interfaces for containerized environments
    console.log("🚀 SERVER IS RUNNING ON http://localhost:3000");
    console.log("📊 Metrics available at http://localhost:3000/metrics");
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

// Register graceful shutdown hook
process.on("SIGTERM", gracefulShutdown);
process.on("SIGINT", gracefulShutdown);

start();
