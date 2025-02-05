import fp from "fastify-plugin";
import client from "prom-client";

/**
 * Analytics plugin for Fastify to track HTTP request metrics and expose Prometheus metrics.
 * @param {import('fastify').FastifyInstance} fastify - Fastify instance
 * @param {Object} options - Plugin options (optional)
 */
async function analytics(fastify, options) {
  // Create a new Prometheus registry
  const register = new client.Registry();

  // Define metrics
  const httpRequestDuration = new client.Histogram({
    name: "http_request_duration_seconds",
    help: "Duration of HTTP requests in seconds",
    labelNames: ["method", "route", "status_code"],
    buckets: [0.1, 0.3, 0.5, 1, 2, 5, 10], // Time buckets in seconds
  });

  const httpRequestCount = new client.Counter({
    name: "http_requests_total",
    help: "Total number of HTTP requests received",
    labelNames: ["method", "route", "status_code"],
  });

  // Register metrics
  register.registerMetric(httpRequestDuration);
  register.registerMetric(httpRequestCount);
  register.setDefaultLabels({ app: "TitanGate" });

  // Collect default system metrics (CPU, Memory, etc.)
  client.collectDefaultMetrics({ register });

  // Middleware to track request timing
  fastify.addHook("onRequest", async (request) => {
    request.startTime = process.hrtime(); // Start tracking request time
  });

  fastify.addHook("onResponse", async (request, reply) => {
    if (!request.startTime) {
      fastify.log.warn("Request start time not found. Skipping metrics collection.");
      return;
    }

    const [seconds, nanoseconds] = process.hrtime(request.startTime);
    const responseTimeInSeconds = seconds + nanoseconds / 1e9;

    const { method, routerPath, url, statusCode } = request;

    // Log request metrics
    try {
      httpRequestDuration
        .labels(method, routerPath || url, statusCode)
        .observe(responseTimeInSeconds);

      httpRequestCount
        .labels(method, routerPath || url, statusCode)
        .inc();
    } catch (error) {
      fastify.log.error("Error recording request metrics:", error);
    }
  });

  // Expose `/metrics` endpoint for Prometheus
  fastify.get("/metrics", async (req, reply) => {
    try {
      reply.header("Content-Type", register.contentType);
      return register.metrics();
    } catch (error) {
      fastify.log.error("Error generating metrics:", error);
      reply.status(500).send("Error generating metrics");
    }
  });
}

export default fp(analytics, {
  name: "analytics-plugin", // Plugin name
  fastify: "4.x", // Specify compatible Fastify version
});