# Serverless and Event-Driven Architecture (Product-Based Companies)

## 1. What are the key challenges of Serverless architectures at scale, and how do you mitigate them?

**Expected Answer:**
Serverless (like AWS Lambda) offers auto-scaling and zero server management, but introduces specific challenges at scale:

*   **Cold Starts:** Initialization latency when a new container spins up.
    *   *Mitigation:* Provisioned Concurrency (keeps instances warm), using lighter runtimes (Go/Rust/Node.js over Java/.NET), optimizing package size, keeping deployment artifacts small.
*   **State Management:** Serverless functions are stateless by nature.
    *   *Mitigation:* Use external state stores (Redis, DynamoDB). For workflow orchestration across functions, use AWS Step Functions or Azure Durable Functions rather than chaining functions together (which leads to the "Serverless Trilemma" or spaghetti architecture).
*   **Connection Pooling Exhaustion:** Thousands of concurrent Lambdas can exhaust connections to relational databases (like PostgreSQL).
    *   *Mitigation:* Use database proxies (AWS RDS Proxy, PgBouncer) or migrate to connectionless DBs (DynamoDB) or HTTP-based data APIs (Aurora Serverless Data API).
*   **Vendor Lock-in:** Code and architecture become tightly coupled to specific cloud provider services (e.g., SQS + Lambda + DynamoDB).
    *   *Mitigation:* Use Hexagonal Architecture to separate business logic from cloud-specific handlers. Consider abstraction frameworks like Serverless Framework, though runtime lock-in is often accepted as a trade-off for speed.
*   **Debugging and Observability:** Tracing requests across independent functions and events is extremely difficult.
    *   *Mitigation:* Mandatory distributed tracing (AWS X-Ray, Datadog), structured logging with correlation IDs passed through all events.

## 2. In an Event-Driven Architecture (EDA), how do you ensure message delivery guarantees and handle failures?

**Expected Answer:**
In EDA, components communicate via asynchronous events. Handling delivery guarantees is critical:

*   **Delivery Guarantees:**
    *   *At-most-once:* Fire and forget. Acceptable for metrics/analytics.
    *   *At-least-once:* Standard in most brokers (Kafka, SQS). Requires consumer to be **idempotent** because duplicates *will* happen.
    *   *Exactly-once:* Extremely difficult/expensive. Kafka supports exactly-once semantics (EOS) within its own ecosystem, but end-to-end exactly-once usually requires idempotency at the final destination (e.g., database upsert with a unique transaction ID).
*   **Handling Failures:**
    *   **Dead Letter Queues (DLQ):** Unprocessable messages are moved to a DLQ after $N$ retries for manual inspection or automated replay.
    *   **Retry with Exponential Backoff & Jitter:** Prevents overwhelming downstream services during a recovery phase.
    *   **Circuit Breaker:** If a downstream service is failing, stop pulling events or pause processing to give the service time to recover.
    *   **Poison Pill Handling:** A message that consistently causes consumer crashes. Must be caught, logged, and moved to a DLQ without blocking the rest of the queue.

## 3. Explain the CQRS (Command Query Responsibility Segregation) pattern. When should it be explicitly avoided?

**Expected Answer:**
**CQRS** separates the models used to update data (Commands) from the models used to read data (Queries).

*   **How it works:** You write to an optimized write store (e.g., Relational DB for ACID transactions). An event is published. A worker consumes the event and updates an optimized read store (e.g., Elasticsearch for text search, Redis for fast key-value, or a denormalized NoSQL view).
*   **Benefits:** Independent scaling of reads and writes. Optimized data models for specific UI views without complex SQL joins.
*   **When to AVOID it (The Anti-Patterns):**
    *   **Simple CRUD applications:** If the UI directly maps to the database tables, CQRS adds massive, unjustified overhead.
    *   **Strong Consistency Requirements:** CQRS inherently relies on **Eventual Consistency**. If the business requires the user to *immediately* read their own write across all systems globally, CQRS introduces complex polling or WebSockets to simulate strong consistency.
    *   **Small Teams/Startups:** The operational complexity of maintaining multiple data stores, message brokers, and event handlers is too high for teams seeking rapid product-market fit.

## 4. How does Event Sourcing differ from a traditional database state model?

**Expected Answer:**
*   **Traditional State (CRUD):** The database stores the *current state* of an entity. When an update happens, the old state is overwritten. History is lost unless explicitly kept in an audit table.
*   **Event Sourcing:** The database stores a sequence of immutable *events* that describe changes to the entity over time. The "current state" is derived by replaying these events from the beginning.
*   **Differences & Benefits:**
    *   **Auditability:** 100% accurate historical record (e.g., accounting ledgers).
    *   **Time Travel:** You can reconstruct the state of the system at any given point in the past.
    *   **Debugging:** If a bug corrupted the state, you can fix the bug, delete the derived state, and replay the events to get the correct state.
*   **Challenges:** Event schema evolution (how to handle events recorded in V1 format when the system is on V3), and read performance (requires "Snapshots" to avoid replaying millions of events).

## 5. Describe the Saga Pattern. Compare Choreography vs. Orchestration.

**Expected Answer:**
The **Saga Pattern** manages distributed transactions across multiple microservices without using 2PC (Two-Phase Commit, which locks resources globally and scales poorly). A Saga is a sequence of local transactions. If one local transaction fails, a series of **compensating transactions** are triggered to undo the work of the preceding steps.

*   **Choreography (Event-Based):**
    *   Each service publishes an event after its local transaction. Other services listen to that event and trigger their own transactions.
    *   *Pros:* Decentralized, no single point of failure.
    *   *Cons:* "Spaghetti architecture." Hard to track the overall flow. Cyclic dependencies can occur. Best for simple Sagas (2-4 steps).
*   **Orchestration (Command-Based):**
    *   A central coordinator (Orchestrator) manages the Saga. It explicitly commands services to execute local transactions and listens for their replies. If a failure happens, the Orchestrator commands the appropriate services to execute compensations.
    *   *Pros:* Centralized logic, easy to monitor, easy to see the state of the transaction. Avoids cyclic dependencies.
    *   *Cons:* Orchestrator can become a bottleneck or a monolithic god-class if not designed well.
    *   *Tools:* AWS Step Functions, Temporal, Camunda.
    *   *Best for:* Complex workflows with many steps.
