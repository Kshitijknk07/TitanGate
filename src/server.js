import Fastify from "fastify";
import rateLimit from "./plugins/rateLimit";
import caching from "./plugins/caching";
import jwt from "./plugins/jwt";
import apiRoutes from "./routes/apiRoutes";

const fastify = Fastify({
  logger: true,
});

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
