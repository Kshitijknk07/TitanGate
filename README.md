# 🚀 **TitanGate** - High-Performance API Gateway with Rate Limiting, Caching, and Authentication and More

TitanGate is a powerful and scalable API Gateway built with Fastify, designed to efficiently manage and optimize your backend services. With a focus on performance, security, and developer experience, TitanGate provides a robust set of features to ensure your APIs are reliable, scalable, and easy to maintain.

---

🛠 **Current Features**
- ✅ **Rate Limiting**: Implemented rate limiting to ensure that API consumers don't overwhelm your services by making too many requests in a short period.
- ✅ **Caching**: Added caching functionality to store frequently requested data for quick retrieval, improving performance and reducing load on your servers.
- ✅ **Authentication**: Integrated JWT-based authentication, allowing secure login and protected routes. Users can log in and receive a token for authorization to access protected resources.
- ✅ **API Versioning**: Managing different API versions to maintain compatibility with older clients while allowing new features to be introduced.
- ✅ 📊 API Analytics: Integrated detailed request logging and performance tracking using Prometheus, allowing real-time monitoring of API traffic, response times, and error rates.


---

## 📈 **Upcoming Features**

The following features are planned for future releases:

- **⚙️ Load Balancing**: Distributing incoming traffic evenly across multiple backend services to ensure high availability and better resource utilization.
- **🔗 GraphQL API Gateway: Dynamically convert REST APIs into GraphQL, allowing clients to fetch only the data they need while improving efficiency and flexibility.
 
---

## 🏁 **Project Status**

- **✅ Completed**:
    - **Rate Limiting**: Protects APIs from too many requests within a short period.
    - **Caching**: Stores commonly used responses to minimize computation time and reduce load on backend servers.
    - **Authentication**: Integrated JWT-based authentication for secure user login and protected routes.
    - **API Versioning**: Managing different API versions to maintain compatibility with older clients while allowing new features to be introduced.
    - 📊 API Analytics: Integrated detailed analytics and logging using Prometheus to track request metrics, API performance, and error rates in real time.

- **🚧 In Progress**:
    - **Load Balancing**: Distributing incoming traffic evenly across multiple backend services for better resource utilization and high availability.
    - 🚀 GraphQL API Gateway: Dynamically converting REST APIs into GraphQL endpoints, providing more flexible and efficient API consumption.

---

## 🧑‍💻 **How to Get Started**

### Prerequisites

Make sure you have the following installed on your machine:

- **Node.js** (v16.x or higher)
- **pnpm** (preferred package manager)

---

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/Kshitijknk07/TitanGate.git
   cd TitanGate
   ```
2. install dependencies:

   ```bash
   pnpm install
   ```
3. Run the project:

   ```bash
   pnpm start
   ```
4. Your API Gateway will be live at:
   
   ```bash
    http://localhost:3000
   ```
   

## 📝 **Usage**

- **Rate Limiting**: Your API requests are limited to **100 requests per minute**. Exceeding the limit will result in a **429 Too Many Requests** error.
- **Caching**: Common responses are cached and served quickly, reducing the time it takes to handle requests.
- **Authentication**: You can authenticate using the **/login** endpoint by providing a valid **username** and **password**. Upon successful authentication, you will receive a **JWT token** that must be included in the **Authorization** header to access protected routes.
- **API Versioning**: You can access different API versions. By default, **v1** is available, and **v2** can be toggled on or off using feature flags.
- **API Analytics**: Integrated analytics and logging to track request metrics, API performance, and error rates in real time.

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

To test, make requests to the following endpoints:

- `GET /` - Returns a simple message confirming that the API Gateway is running.
- `GET /metrics` - Returns real-time API analytics data.



## 🛠 **Tech Stack**

- **Backend Framework**: [Fastify](https://www.fastify.io/)
- **Rate Limiting**: [@fastify/rate-limit](https://www.npmjs.com/package/@fastify/rate-limit)
- **Caching**: [@fastify/caching](https://www.npmjs.com/package/@fastify/caching)
- **Authentication**: [fastify-jwt](https://www.npmjs.com/package/fastify-jwt) (JWT-based authentication for secure API access)
- **API Analytics**: [fastify-metrics](https://www.npmjs.com/package/fastify-metrics) (Collecting detailed API request logs, performance metrics, and error tracking)





