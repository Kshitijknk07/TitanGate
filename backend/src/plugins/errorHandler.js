// Placeholder error handler plugin for Fastify

export default async function errorHandler(fastify, options) {
  fastify.setErrorHandler((error, request, reply) => {
    fastify.log.error(error);
    reply.status(500).send({ error: 'Internal Server Error' });
  });
} 