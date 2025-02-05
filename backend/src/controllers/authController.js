export async function login(request, reply) {
  const { username, password } = request.body;

  if (username === "admin" && password === "password") {
    const token = request.jwt.sign({ username });
    return reply.send({ token });
  }
  return reply.status(401).send({ message: "Invalid username or password" });
}
