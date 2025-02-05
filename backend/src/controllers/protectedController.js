export async function protectedRoute(request, reply) {
  try {
    await request.jwtVerify();
    return { message: "You are authorized" };
  } catch (err) {
    return reply.status(401).send({ message: "You are not authorized" });
  }
}
