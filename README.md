# 🚀 **TitanGate** - High-Performance API Gateway with Rate Limiting, Caching, and Authentication and More

Welcome to **TitanGate**, a powerful and scalable **API Gateway** built with **Fastify**. TitanGate handles key features like **Rate Limiting**, **Caching**, and **Authentication** to efficiently manage and optimize your backend services.

---

## 🛠 **Current Features**

- **✅ Rate Limiting**: Implemented rate limiting to ensure that API consumers don't overwhelm your services by making too many requests in a short period.
- **✅ Caching**: Added caching functionality to store frequently requested data for quick retrieval, improving performance and reducing load on your servers.
- **✅ Authentication**: Integrated JWT-based authentication, allowing secure login and protected routes. Users can log in and receive a token for authorization to access protected resources.

---

## 📈 **Upcoming Features**

The following features are planned for future releases:

- **⚙️ Load Balancing**: Distributing incoming traffic evenly across multiple backend services to ensure high availability and better resource utilization.
- **⚡ API Analytics**: Integrating detailed analytics and logging to track request metrics, API performance, and error rates.
- **🔄 API Versioning**: Managing different API versions to maintain compatibility with older clients while allowing new features to be introduced.

---

## 🏁 **Project Status**

- **✅ Completed**:
    - Rate Limiting: Protects APIs from too many requests within a short period.
    - Caching: Stores commonly used responses to minimize computation time and reduce load on backend servers.

- **🚧 In Progress**:
    - Authentication and other advanced features such as load balancing, API analytics, and versioning will be implemented soon!

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
    http://localhost:3000!
   ```
   

## 📝 **Usage**

- **Rate Limiting**: Your API requests are limited to **100 requests per minute**. Exceeding the limit will result in a **429 Too Many Requests** error.
- **Caching**: Common responses are cached and served quickly, reducing the time it takes to handle requests.
- **Authentication**: You can authenticate using the **/login** endpoint by providing a valid **username** and **password**. Upon successful authentication, you will receive a **JWT token** that must be included in the **Authorization** header to access protected routes.

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

To test, make requests to the following endpoints:

- `GET /` - Returns a simple message confirming that the API Gateway is running.



## 🛠 **Tech Stack**

- **Backend Framework**: [Fastify](https://www.fastify.io/)
- **Rate Limiting**: [@fastify/ratelimit](https://www.npmjs.com/package/@fastify/ratelimit)
- **Caching**: [@fastify/caching](https://www.npmjs.com/package/@fastify/caching)
- **Authentication (Coming Soon)**: JWT, OAuth2, etc.




