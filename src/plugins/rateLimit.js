import ratelimit from "@fastify/rate-limit";

export default async function (fastify, options) {
  fastify.register(ratelimit, {
    max: 100, // Maximum number of requests
    timeWindow: "1 minute", // Time window for the limit
  });
}
