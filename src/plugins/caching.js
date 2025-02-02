import caching from "@fastify/caching";

export default async function (fastify, options) {
  fastify.register(caching, {
    expiresIn: 1000 * 60, // Cache expiration time (1 minute)
  });
}
