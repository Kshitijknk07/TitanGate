import Fastify from "fastify";
import rateLimit from "./plugins/rateLimit.js";
import caching from "./plugins/caching.js";
import jwt from "./plugins/jwt.js";
import analytics from "./plugins/analytics.js";
import apiRoutes from "./routes/apiRoutes.js";
import v1 from "./v1/index.js";
import v2 from "./v2/index.js";

const fastify = Fastify({ logger: true });

// Feature flags
const featureFlags = {
  v2Enabled: true, // Enable v2
};

// Middleware to check if a feature is enabled
fastify.addHook("onRequest", async (request, reply) => {
  if (request.url.includes("/v2") && !featureFlags.v2Enabled) {
    reply.code(404).send("Not Found");
  }
});

// Register versions
fastify.register(v1, { prefix: "/v1" }); // http://localhost:3000/v1
fastify.register(v2, { prefix: "/v2" }); // http://localhost:3000/v2

// Register plugins
fastify.register(rateLimit);
fastify.register(caching);
fastify.register(jwt);
fastify.register(analytics); // Register analytics

// Register routes
fastify.register(apiRoutes);

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

start();
