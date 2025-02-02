import Fastify from "fastify";

const v1 = async (fastify, options) => {
  fastify.get("/users", async (request, reply) => {
    return { message: "This is version 1 of the users endpoint" };
  });

  fastify.get("/products", async (request, reply) => {
    return { message: "This is version 1 of the products endpoint" };
  });
};

export default v1;
