# 🛠️ 04 — Scaling & Microservices
> **Most Asked in Product-Based Companies** | 🛠️ Difficulty: Hard

---

## 🔑 Must-Know Topics
- API Gateway and Service Mesh patterns
- Monolithic vs Microservices architecture
- Handling distributed transactions (Saga Pattern)
- Node.js scaling strategies (Horizontal vs Vertical, Load Balancing)
- Service Discovery and Circuit Breakers (Resilience)

---

## ❓ Frequently Asked Questions

### Q1. How do you implement the Saga Pattern for distributed transactions in a Node.js microservices architecture?

**Answer:**
In a microservices architecture, operations often span multiple services (e.g., An Order Service, a Payment Service, and an Inventory Service). Since you can't use traditional ACID transactions across different databases, you use a **Saga**.

A Saga is a sequence of local transactions. Each service performs a local transaction and publishes an event that triggers the next step in the saga.

**Two implementation approaches:**
1. **Choreography (Event-Driven):** Services subscribe to each other's events.
   - *Pros:* Decentralized, no single point of failure.
   - *Cons:* Hard to track the whole workflow for complex sagas.
   - *Example:* Order Service creates 'Pending' order -> Emits `OrderCreated` -> Payment Service listens, charges card -> Emits `PaymentCleared` -> Inventory Service listens, reserves items -> Emits `InventoryReserved` -> Order Service marks 'Complete'.
2. **Orchestration (Command-Driven):** A central "Orchestrator" service tells other services what to do.
   - *Pros:* Easier to manage complex workflows and centralize state.
   - *Cons:* The orchestrator becomes a potential bottleneck/single point of failure.

**Handling Failures (Compensating Transactions):**
If a step fails (e.g., Payment clears, but Inventory is out of stock), the saga executes a series of *compensating transactions* to undo the previous steps (e.g., an event is emitted to refund the payment and mark the Order as 'Failed').

---

### Q2. What is the Circuit Breaker pattern and why use it in Node.js?

**Answer:**
When a microservice A depends on microservice B, if Service B goes down or becomes extremely slow, requests from Service A will queue up, consuming resources, and eventually causing Service A to crash (cascading failure).

The **Circuit Breaker** (like the `opossum` package in Node.js) prevents this by monitoring the failure rate of the external call.

**States:**
1. **Closed:** Requests flow normally. If errors exceed a threshold (e.g., 50% in 10s), it trips to *Open*.
2. **Open:** All requests instantly fail (fast-fail) and return a fallback response. It stops hammering the failing Service B, giving it time to recover.
3. **Half-Open:** After a timeout, the circuit lets a few requests through to test if Service B has recovered. If they succeed, it closes the circuit. If they fail, it trips back to Open.

---

### Q3. Compare GraphQL vs REST vs gRPC in Node.js Microservices.

**Answer:**

1. **REST (JSON over HTTP/1.1):**
   - **Pros:** Standard, cacheable, easy to debug, widespread tooling.
   - **Cons:** Over-fetching or under-fetching of data. Not strongly typed by default.
   - **Best for:** Public-facing APIs.

2. **GraphQL (Query over HTTP):**
   - **Pros:** Clients ask for exactly what they need in a single request. Solves over/under-fetching. Strongly typed schema.
   - **Cons:** Harder to cache at the HTTP level. Complex to set up (N+1 query problems).
   - **Best for:** Mobile apps, complex frontends (React/Next.js) aggregating data from multiple services.

3. **gRPC (Protobufs over HTTP/2):**
   - **Pros:** Extremely fast (binary format, much smaller payloads), multiplexed streams via HTTP/2, strongly typed contracts (`.proto` files).
   - **Cons:** Binary payloads are harder to debug without tools. Browser support is limited (requires grpc-web proxy).
   - **Best for:** Internal Microservice-to-Microservice communication.

---

### Q4. What is Service Discovery and why is it needed?

**Answer:**
In modern cloud environments, microservices are deployed dynamically in containers (like Docker/Kubernetes). Their IP addresses and ports change frequently due to scaling, updates, or failures. Hardcoding URLs in configuration files no longer works.

**Service Discovery** (e.g., Consul, Eureka, or Kubernetes DNS) allows services to find each other dynamically.

**How it works (Client-side discovery):**
1. Service A needs to call Service B.
2. Service A queries the Service Registry (e.g., Consul) asking, "Where is Service B?".
3. The Registry returns a list of healthy IP/Port combinations.
4. Service A uses a local load balancer algorithm to pick one and makes the request.

*(Note: In Kubernetes systems, this is abstracted away using internal DNS and Services, meaning the Node.js developer just makes a request to `http://service-b-name`).*
