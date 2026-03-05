# ⚙️ Microservices Fundamentals — Interview Questions (Service-Based Companies)

This document covers microservices concepts commonly tested at service-based companies like TCS, Infosys, Wipro, Capgemini, HCL. Targeted at 1–5 years of experience rounds.

---

### Q1: What are Microservices? How do they compare to a Monolithic architecture?

**Answer:**
**Microservices** is an architectural style that structures an application as a collection of small, autonomous services modeled around business domains.

| Feature | Monolithic Architecture | Microservices Architecture |
|---|---|---|
| Deployment | Single deployable unit (e.g., one WAR/EAR file) | Independent deployments per service |
| Database | One shared database for all modules | Each service has its own database (Decentralized Data Management) |
| Tech Stack | Tied to one language/framework stack | Polyglot — different tech stacks can be used per service |
| Scaling | Scale the entire monolith | Scale individual bottlenecks separately |
| Fault Isolation| Bug in one module (memory leak) crashes the whole app | Bug in one service doesn't crash others |
| Complexity | Low initial complexity, becomes tangled over time | High initial complexity (distributed systems overhead) |

**When to use Microservices:** Large teams, complex business domains, necessity for independent scaling. (Not recommended for small startups finding product-market fit).

---

### Q2: Why should each Microservice have its own Database?

**Answer:**
This rule is known as the **Database-per-service pattern**.

**Why it's essential:**
1. **Loose Coupling:** If multiple services share a database, changing the schema for Service A might break Service B. With private databases, schemas can evolve independently.
2. **Independent Scaling:** The order database might need a NoSQL datastore (high write throughput), while the product catalog might need a relational database or search engine.
3. **Encapsulation:** Services can only communicate via well-defined APIs. Direct database access breaks encapsulation.

**Challenges this introduces:**
- Difficult to do distributed transactions (requires Sagas).
- Difficult to query data that spans multiple services (requires API composition or CQRS).

---

### Q3: What is an API Gateway? Why do we need it?

**Answer:**
An **API Gateway** acts as a reverse proxy, heavily utilized as the single entry point for all client requests into a microservices architecture.

**Why it's needed:**
Without an API Gateway, a mobile client might need to make 5 different HTTP calls to 5 different services (User, Order, Inventory, Payment, Notification) to render one screen. This is slow and exposes internal IPs.

**Key Responsibilities of an API Gateway:**
- **Routing:** Directs `/api/users` to the User Service, `/api/orders` to the Order Service.
- **Authentication & Authorization:** Validates JWTs at the edge so backend services don't have to duplicate logic.
- **Rate Limiting:** Prevents DDoS and brute-force attacks.
- **Response Aggregation (BFF - Backend for Frontend):** Gathers data from 4 services, combines it, and sends one consolidated response to the client, reducing roundtrips.
- **Protocol Translation:** Can translate external HTTP REST into internal gRPC.

**Examples:** Spring Cloud Gateway, Kong, NGINX, AWS API Gateway.

---

### Q4: Explain the Circuit Breaker pattern.

**Answer:**
The **Circuit Breaker** pattern prevents an application from repeatedly trying to execute an operation that is likely to fail, preventing cascading failures in a distributed system.

**How it works (3 States):**

1. **CLOSED (Normal):** Everything is healthy. Requests flow through. If a request fails, a failure counter increments.
2. **OPEN (Failing):** If failures reach a defined threshold (e.g., 50% failure over 10 seconds), the circuit trips to OPEN. All further requests **fail fast** immediately without calling the backend. This gives the failing downstream service time to recover.
3. **HALF-OPEN (Testing):** After a timeout period, the circuit allows a limited number of test requests through.
   - If they succeed → Circuit resets to CLOSED.
   - If they fail → Circuit trips back to OPEN and resets the timeout.

**Fallback:** When a circuit is OPEN, the service can execute a fallback method (e.g., return cached data, or an empty list, or a generic error) instead of throwing an exception.

**Tools:** Resilience4j, Netflix Hystrix (deprecated).

---

### Q5: How do Microservices communicate with each other? (Synchronous vs Asynchronous)

**Answer:**

**1. Synchronous Communication:**
- **How:** The caller sends a request and blocks (waits) until a response is received.
- **Protocols:** HTTP/REST, gRPC, GraphQL.
- **Pros:** Simple, immediate feedback, easy to trace.
- **Cons:** Creates tight coupling. If Service B is down, Service A fails (chain of failure). Latencies add up.
- **Use case:** User querying their order status.

**2. Asynchronous Communication (Event-Driven):**
- **How:** Service A emits a message/event and continues its work. Service B consumes it whenever it's ready.
- **Protocols:** Message Brokers (RabbitMQ, Apache Kafka), AWS SQS.
- **Pros:** Loose coupling, high availability (Service B can be down, messages queue up and process later), excellent for traffic spikes.
- **Cons:** System is only eventually consistent. Harder to trace and debug.
- **Use case:** Processing an order after payment completes, sending welcome emails.

**Best Practice:** Use Asynchronous communication for state-changing operations (writes). Use Synchronous communication for simple data retrieval (reads).

---

*Prepared for technical screening at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, Tech Mahindra).*
