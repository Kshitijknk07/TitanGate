import Fastify from "fastify";

const v2 = async (fastify, options) => {
  fastify.get("/users", async (request, reply) => {
    return {
      message: "This is version 2 of the users endpoint with new features!",
    };
  });

  fastify.get("/products", async (request, reply) => {
    return {
      message: "This is version 2 of the products endpoint with new features!",
    };
  });
};

export default v2;
