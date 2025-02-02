import { login } from "../controllers/authController";
import { protectedRoute } from "../controllers/protectedController";

export default async function (fastify, options) {
  // Authentication route
  fastify.post("/login", login);

  // Protected route
  fastify.get("/protected", protectedRoute);

  // Public route
  fastify.get("/", async (request, reply) => {
    return { message: "API GATEWAY IS RUNNING!!" };
  });
}
