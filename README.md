# TitanGate - High-Performance API Gateway



## Overview

TitanGate is a high-performance, scalable, and secure API Gateway built using Fastify. Designed to optimize API management, it ensures smooth communication between clients and backend services with robust features such as rate limiting, caching, authentication, analytics, and load balancing.

## Key Features

- **Rate Limiting**: Protects APIs from abuse by limiting requests per client.
- **Caching**: Enhances response times and reduces server load with intelligent caching mechanisms.
- **Authentication & Authorization**: Implements JWT-based security for endpoint access control.
- **API Versioning**: Supports multiple API versions for seamless service evolution.
- **API Analytics**: Provides real-time monitoring and performance tracking with Prometheus.
- **Load Balancing**: Distributes incoming requests across multiple backend services for better availability.
- **GraphQL API Gateway (Upcoming)**: Converts REST APIs dynamically into GraphQL for flexible data querying.

---

## Project Structure

### Backend Architecture

```
backend/
├── src/
│   ├── controllers/
│   │   ├── authController.js      # Handles authentication logic
│   │   ├── userController.js      # Manages user-related endpoints
│   │   └── protectedController.js # Secured API routes
│   │
│   ├── loadbalancer/
│   │   └── loadBalancer.js       # Implements traffic distribution strategies
│   │
│   ├── plugins/
│   │   ├── analytics.js         # Logs and tracks API metrics
│   │   ├── caching.js           # Implements caching mechanisms
│   │   ├── jwt.js               # Manages authentication middleware
│   │   ├── rateLimit.js         # Rate limiting logic
│   │   └── graphqlGateway.js    # Converts REST to GraphQL (Upcoming Feature)
│   │
│   ├── routes/
│   │   ├── apiRoutes.js         # Defines all API endpoints
│   │   ├── authRoutes.js        # Authentication endpoints
│   │   └── userRoutes.js        # User management endpoints
│   │
│   ├── v1/                      # API Version 1 Implementation
│   ├── v2/                      # API Version 2 Implementation
│   └── server.js                 # Application entry point
│
├── package.json                 # Project dependencies
├── pnpm-lock.yaml               # Dependency lockfile
└── .gitignore                   # Files and folders to be ignored by Git
```

### Frontend Architecture

```
frontend/
├── src/
│   ├── components/
│   │   ├── APIAnalytics.jsx    # Displays API analytics & metrics
│   │   ├── Header.jsx          # Reusable header component
│   │   ├── Sidebar.jsx         # Sidebar navigation
│   │   ├── RateLimit.jsx       # Displays rate limiting information
│   │   ├── GraphQLExplorer.jsx # UI for interacting with GraphQL API
│   │
│   ├── hooks/
│   │   ├── useAuth.js         # Authentication state management
│   │   ├── useFetch.js        # Reusable API request hook
│   │
│   ├── pages/
│   │   ├── Dashboard.jsx      # Displays key metrics and analytics
│   │   ├── Home.jsx           # Homepage with overview details
│   │   ├── Login.jsx          # Login page
│   │   └── GraphQLPlayground.jsx # UI for GraphQL queries (Upcoming)
│   │
│   ├── utils/
│   │   └── api.js             # Handles API requests
│   │
│   ├── app.jsx                 # Main application file
│   ├── index.css               # Global styles
│   └── main.jsx                # Root entry file
```

---

## Project Status

### ✅ Completed Features:

- **Rate Limiting**: Protects endpoints from excessive requests.
- **Caching**: Optimizes performance with response caching.
- **Authentication**: Implements JWT-based security.
- **API Versioning**: Supports multiple API versions.
- **API Analytics**: Provides monitoring via Prometheus.
- **Load Balancing**: Distributes traffic for high availability.

### 🚀 In Progress:

- **GraphQL API Gateway**: Dynamically converts REST APIs into GraphQL.
- **Improved Admin Dashboard**: Enhanced UI for API monitoring & management.

---

## Getting Started

### Prerequisites

Ensure you have the following installed:

- Node.js (>=16.x)
- pnpm (or npm/yarn)
- Docker (optional for running services)

### Installation

```sh
git clone https://github.com/Kshitijknk07/TitanGate.git
cd TitanGate
pnpm i
```

### Running the Backend

```sh
cd backend
pnpm start
```

Once the backend starts, you should see output similar to:

```
🚀 SERVER IS RUNNING ON http://localhost:3000
📊 Metrics available at http://localhost:3000/metrics
```

### Running the Frontend

```sh
cd frontend
pnpm run dev
```

Once the frontend starts, you should see output similar to:

```
VITE v6.1.0  ready in 503 ms
➜  Local:   http://localhost:5173/
➜  Network: use --host to expose
```

---
## 📝 **Usage**

- **Rate Limiting**: Your API requests are limited to **100 requests per minute**. Exceeding the limit will result in a **429 Too Many Requests** error.
- **Caching**: Common responses are cached and served quickly, reducing the time it takes to handle requests.
- **Authentication**: You can authenticate using the **/login** endpoint by providing a valid **username** and **password**. Upon successful authentication, you will receive a **JWT token** that must be included in the **Authorization** header to access protected routes.
- **API Versioning**: You can access different API versions. By default, **v1** is available, and **v2** can be toggled on or off using feature flags.
- **API Analytics**: Integrated analytics and logging to track request metrics, API performance, and error rates in real time.
- **Load Balancing**: Incoming traffic is distributed evenly across multiple backend services using a round-robin algorithm to ensure high availability and better resource utilization.

### Example Request Flow:
1. **Login** to get the token:
   - **POST** `/login`
   - Request Body:
     ```json
     {
       "username": "admin",
       "password": "password"
     }
     ```
   - Response:
     ```json
     {
       "token": "your-jwt-token"
     }
     ```

2. **Access protected route** with token:
   - **GET** `/protected`
   - Add the following header in your request:
     ```
     Authorization: Bearer your-jwt-token
     ```
   - Response (if authorized):
     ```json
     {
       "message": "You are authorized"
     }
     ```
   - Response (if not authorized):
     ```json
     {
       "message": "You are not authorized"
     }
     ```

### API Versioning:
- **v1**: Available by default at `/v1`.
- **v2**: Can be accessed at `/v2` if enabled in the feature flags (`featureFlags.v2Enabled = true`).
- **Note**: If **v2** is disabled via the feature flag, attempting to access `/v2` will result in a **404 Not Found** error.

### API Analytics:
- **Endpoint**: `GET /metrics`
- **Response Example**:
  ```json
  {
    "requests_total": 1500,
    "error_rate": 0.02,
    "avg_response_time": "120ms"
  }
  ```
- Analytics provides insights into API usage, error occurrences, and response performance.
- **Load Balancing**:

    How it works: The load balancer automatically distributes incoming traffic across multiple backend services using a round-robin approach. This ensures high availability and better performance by balancing the load among your backend servers.
    In case of failure: If a backend service becomes unavailable, the load balancer will attempt to forward the request to the next available backend, ensuring minimal disruption.
    Backend Services: The system will rotate between backend services like http://backend-service-1, http://backend-service-2, and http://backend-service-3 for a more resilient and scalable architecture.

To test, make requests to the following endpoints:

- `GET /` - Returns a simple message confirming that the API Gateway is running.
- `GET /metrics` - Returns real-time API analytics data.
- `GET /v1` - Access the v1 API version.
- `GET /v2` - Access the v2 API version (if enabled).

  **MORE TO COME**

