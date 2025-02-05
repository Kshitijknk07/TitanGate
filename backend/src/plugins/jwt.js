import fastifyPlugin from 'fastify-plugin';
import fastifyJwt from '@fastify/jwt';

async function jwtPlugin(fastify, options) {
  fastify.register(fastifyJwt, {
    secret: 'your-secret-key'
  });

  fastify.decorate('authenticate', async function (request, reply) {
    try {
      await request.jwtVerify();
    } catch (err) {
      reply.send(err);
    }
  });
}

export default fastifyPlugin(jwtPlugin);
