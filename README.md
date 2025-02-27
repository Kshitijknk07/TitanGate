# 🚀 TitanGate API Gateway

> A powerful, modern API Gateway built with Go & Fiber

## 🌟 What is TitanGate?

TitanGate is an  API Gateway that acts as a unified entry point for all your microservices. It handles routing, load balancing, caching, and monitoring - allowing you to focus on building your services while it manages the traffic.

## ⚡ Core Features

### 🔄 Intelligent Load Balancing
- Round-robin distribution
- Health checking of backends
- Automatic failover
- Weighted distribution
- Connection tracking

### 🛡️ Rate Limiting
- IP-based rate limiting
- Configurable thresholds
- Automatic blocking
- Rate limit analytics

### 📦 Smart Caching
- Redis-based caching
- LRU cache implementation
- Configurable TTL
- Cache invalidation
- Hit/miss metrics

### 📊 Real-time Metrics
- Request counts
- Response times
- Cache hit rates
- Backend health
- Live dashboard

### 🔢 API Versioning
- Multiple API versions
- Header-based routing
- Path-based versioning
- Default version fallback

## 💫 How It Works

```mermaid
graph LR
    Client --> TitanGate
    TitanGate --> RateLimit[Rate Limiter]
    RateLimit --> Cache[Cache Layer]
    Cache --> LoadBalancer[Load Balancer]
    LoadBalancer --> Service1[Backend Service 1]
    LoadBalancer --> Service2[Backend Service 2]
    LoadBalancer --> Service3[Backend Service 3]


