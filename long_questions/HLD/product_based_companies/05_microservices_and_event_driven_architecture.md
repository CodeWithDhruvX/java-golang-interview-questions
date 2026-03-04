# High-Level Design (HLD): Microservices and Event-Driven Architecture

Modern product-based companies heavily utilize Microservices to scale their organizational teams as well as their software.

## 1. Monolith vs. Microservices. Why and when do we migrate?
**Answer:**
*   **Monolith:** A single executable containing all business logic. Good for early stage (fast iteration, simple deployment, no network latency between modules).
*   **Microservices:** Breaking the monolith into small, independently deployable services organized around business capabilities.
*   **When to Migrate:**
    *   When the team grows too large for a single codebase (merge conflicts, slow PR reviews).
    *   When different parts of the system require distinct scaling (e.g., the Image Processing module needs GPU VMs, while the User Login module needs lightweight CPU instances).
    *   When you want to use polyglot technologies (one service in Go, one in Node, one in Python).

## 2. Explain API Gateway and its responsibilities.
**Answer:**
In a microservices architecture, having the client app talk directly to 50 different microservices is unmanageable (complex networking, multiple domains, security exposure).
*   **API Gateway:** A single entry point that sits between the client and the microservices.
*   **Responsibilities:**
    *   *Routing / Reverse Proxying:* Directing the `/users` endpoint to the User Service, and `/payments` to the Payment Service.
    *   *Authentication/Authorization:* Validating JWT tokens before the request hits backend services.
    *   *Rate Limiting & Throttling:* Protecting backends from DDOS or overload.
    *   *Cross-Cutting Concerns:* Logging, metrics aggregation, SSL termination.
    *   *Response Aggregation (BFF - Backend for Frontend):* Making 3 simultaneous API calls to different services and composing them into a single JSON payload for a mobile client.

## 3. How do microservices communicate? (Synchronous vs Asynchronous)
**Answer:**
*   **Synchronous Communication (REST / gRPC):**
    *   Service A calls Service B and blocks waiting for a response.
    *   *gRPC:* Uses Protobufs over HTTP/2. Highly compressed binary data, much faster than REST JSON. Used heavily for internal microservice-to-microservice communication.
    *   *Drawback:* High coupling. If Service B is down, Service A fails (unless Circuit Breakers are used).
*   **Asynchronous Communication (Message Queues / Pub-Sub):**
    *   Service A publishes an event to a broker (Kafka/RabbitMQ) and immediately returns. Service B, C, and D consume the event at their own pace.
    *   *Benefits:* Loose coupling, fault isolation (if B goes down, messages sit in the queue until B recovers), ability to handle traffic spikes.

## 4. What is a Circuit Breaker pattern?
**Answer:**
If Service A calls Service B, and Service B is experiencing massive DB delays, Service B will take 30 seconds to respond. Very quickly, all threads on Service A will be blocked waiting for B, causing cascading failures that take down the entire system.
*   **Circuit Breaker:** Sits between A and B. It monitors failures.
*   *Closed State:* Normal operation. Requests flow freely.
*   *Open State:* If failures exceed a threshold (e.g., 50% failures over 10 seconds), the circuit "opens." Service A immediately fails fast or returns a fallback cache, rather than making the network call to B. This gives B time to recover.
*   *Half-Open State:* After a configurable timeout, it lets a few test requests through. If they succeed, it closes; if they fail, it remains open.

## 5. How do you handle Distributed Transactions? (Saga Pattern)
**Answer:**
Microservices advocate "Database per Service." Therefore, traditional ACID SQL transactions (two-phase commit) spanning multiple services do not work.
**Saga Pattern:**
A Saga is a sequence of local transactions. Each service updates its own database and publishes an event that triggers the next local transaction in the saga.
*   *Choreography (Event-Based):* Services subscribe to each other's events. (Good for simple workflows). E.g., OrderService creates Order -> sends `OrderCreated` -> PaymentService charges card.
*   *Orchestration (Command-Based):* A central Saga Orchestrator service commands the microservices to execute local transactions. (Better for complex workflows).
*   **Compensating Transactions:** If a step fails (e.g., Order created, Payment successful, but Inventory fails), the Saga must execute compensating transactions (refund payment, cancel order) to rollback the distributed state.

## 6. What is Event Sourcing?
**Answer:**
Instead of storing the current state of an entity in a database (e.g., User balance = $50), we store a sequence of immutable events representing state changes (e.g., Deposit +100, Withdraw -50).
*   *Rebuilding State:* The current state is derived by replaying the events.
*   *Benefits:* Perfect audit log, ability to time-travel to any past state. Highly scalable writes (append-only logs).
*   *Related Pattern:* CQRS (Command Query Responsibility Segregation). Writes go to the event log, background workers update a read-optimized view database.
