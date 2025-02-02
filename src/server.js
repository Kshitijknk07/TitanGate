import Fastify from "fastify";
import rateLimit from "./plugins/rateLimit";
import caching from "./plugins/caching";
import jwt from "./plugins/jwt";
import apiRoutes from "./routes/apiRoutes";
import v1 from "./v1";
import v2 from "./v2";

const fastify = Fastify({
  logger: true,
});

// Feature flags
const featureFlags = {
  v2Enabled: true, // Enable v2
};

// Middeleware to check if a feature is enabled
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

// Register routes
fastify.register(apiRoutes);

// Start the server
const start = async () => {
  try {
    await fastify.listen(3000);
    console.log("SERVER IS RUNNING ON http://localhost:3000");
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

start();
