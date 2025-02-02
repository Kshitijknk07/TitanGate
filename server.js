import Fastify from "fastify";
import ratelimit from "@fastify/rate-limit";
import caching from "@fastify/caching";

const fastify = Fastify({
  logger: true,
});

// RATE LIMITING
fastify.register(ratelimit, {
  max: 100, // MAXIMUM NUMBER OF REQUESTS
  timeWindow: "1 minute", // TIME WINDOW FOR THE LIMIT
});

// CACHING
fastify.register(cashing, {
  expiresIn: 1000 * 60, // CACHE EXPIRATION TIME
});

// ROUTES
fastify.get("/", async (Request, reply) => {
  return { message: "API GATEWAY IS RUNNING!!" };
});

const start = async () => {
  try {
    await fastify.listen({ port: 3000 });
    console.log("SERVER IS RUNNING ON http://localhost:3000");
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

start();
