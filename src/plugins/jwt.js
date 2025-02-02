import jwt from "fastify-jwt";

export default async function (fastify, options) {
  fastify.register(jwt, {
    secret: "supersecret",
  });
}
