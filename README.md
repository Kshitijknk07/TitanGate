# 🚀 TitanGate API Gateway

> A high-performance API Gateway built with Go & Fiber, featuring real-time monitoring, load balancing, caching, and security middleware.

## 🌟 Overview

TitanGate provides a unified entry point for your microservices, with features for traffic management, monitoring, and security.  
It is built for performance and extensibility.

---

## ⚡ Core Features

- **Load Balancing:**  
  - Round-robin (default, weighted and other algorithms can be added)
  - Automatic backend health checks

- **Security & Rate Limiting:**  
  - IP-based rate limiting (configurable)
  - JWT-based authentication

- **Caching:**  
  - Redis integration for distributed caching
  - In-memory LRU cache

- **Monitoring:**  
  - Prometheus metrics endpoint (`/metrics`)
  - Real-time dashboard (served from `/static/`)

- **Basic Routing:**  
  - No API versioning (all endpoints are flat, e.g., `/user`, `/health`)

---

## 🏗️ Architecture

```
Client
  │
  ▼
TitanGate (Fiber)
  ├── Auth Middleware (JWT)
  ├── Rate Limiter
  ├── Cache Layer (Redis/LRU)
  ├── Load Balancer (Round Robin)
  ├── Health Checker
  ├── Prometheus Metrics
  └── Static Dashboard
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Redis server (for caching and rate limiting)

### Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/TitanGate.git
    cd TitanGate
    ```

2. **Install Go dependencies:**
    ```sh
    go mod download
    ```

3. **Configure environment variables:**
    - Copy or create a `.env` file in the `backend/` directory:
      ```
      PORT=8081
      REDIS_HOST=localhost
      REDIS_PORT=6379
      JWT_SECRET=your_jwt_secret_here
      RATE_LIMIT=100
      CACHE_TTL=300
      LOG_LEVEL=info
      ```
    - Adjust values as needed.

4. **Start the server:**
    ```sh
    cd backend
    go run cmd/api/main.go
    ```

---

## 📊 Dashboard

- Access the dashboard at [http://localhost:8081/](http://localhost:8081/)
- Prometheus metrics available at [http://localhost:8081/metrics](http://localhost:8081/metrics)

---

## 🔧 API Endpoints

- `GET /user` — Example user endpoint
- `GET /health` — Health check endpoint
- `GET /metrics` — Prometheus metrics

> **Note:**  
> There is currently **no API versioning** and no management API for backends.  
> All endpoints are flat (e.g., `/user`, `/health`).

---

## 🛠️ Development

### Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── cache/
│   ├── config/
│   ├── handlers/
│   ├── loadbalancer/
│   ├── metrics/
│   ├── middleware/
│   ├── routes/
│   └── services/
├── static/
│   ├── index.html
│   ├── css/
│   └── js/
└── .env
```

---

## ⚠️ Project Status

- **Core gateway functionality is implemented and works.**
- **API versioning and management endpoints are NOT implemented.**
- **Only round-robin load balancing is available by default.**
- **Dashboard is available but may be basic.**
- **Some advanced features listed in the original README are not present or are placeholders.**

---

## 🤝 Contributing

Pull requests are welcome! Please open issues for bugs or feature requests.

---

## 📝 License

MIT License

---

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [Redis](https://redis.io/)
- [Prometheus](https://prometheus.io/)
- [Chart.js](https://www.chartjs.org/)
