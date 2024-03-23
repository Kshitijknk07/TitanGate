# 🚀 TitanGate API Gateway

A modern, high-performance API Gateway built with Go & Fiber. Perfect for managing and securing your microservices.

## ✨ Features

- **Smart Load Balancing**
  - Round-robin distribution
  - Automatic health checks
  - Service discovery support

- **Security & Protection**
  - JWT authentication
  - Rate limiting
  - Request/response transformation
  - IP filtering

- **Performance**
  - Redis caching
  - In-memory LRU cache
  - Request/response compression

- **Monitoring**
  - Real-time metrics dashboard
  - Prometheus integration
  - Health status monitoring

## 🚀 Quick Start

1. **Install Dependencies**
   ```bash
   # Install Go 1.21+
   # Install Redis
   ```

2. **Setup**
   ```bash
   git clone https://github.com/yourusername/TitanGate.git
   cd TitanGate/backend
   go mod download
   ```

3. **Configure**
   Create `.env` in `backend/`:
   ```
   PORT=8081
   REDIS_HOST=localhost
   REDIS_PORT=6379
   JWT_SECRET=your_secret
   RATE_LIMIT=100
   CACHE_TTL=300
   LOG_LEVEL=info
   ```

4. **Run**
   ```bash
   go run cmd/api/main.go
   ```

## 📊 Dashboard

- Main Dashboard: `http://localhost:8081/`
- Metrics: `http://localhost:8081/metrics`

## 🔧 API Endpoints

- `GET /user` - User endpoint
- `GET /health` - Health check
- `GET /metrics` - Prometheus metrics

## 🏗️ Project Structure

```
backend/
├── cmd/          # Application entry points
├── internal/     # Core packages
│   ├── cache/    # Caching logic
│   ├── config/   # Configuration
│   ├── handlers/ # HTTP handlers
│   ├── middleware/ # Middleware components
│   └── services/ # Business logic
└── static/       # Dashboard assets
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

MIT License

---

Built with ❤️ using [Fiber](https://github.com/gofiber/fiber), [Redis](https://redis.io/), and [Prometheus](https://prometheus.io/)
