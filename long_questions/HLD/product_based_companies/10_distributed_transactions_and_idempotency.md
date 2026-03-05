# High-Level Design (HLD): Distributed Transactions and Idempotency

In microservices architectures, ensuring data consistency across multiple databases without standard ACID guarantees is one of the hardest problems you will be asked to solve in an interview.

## 1. What is a Distributed Transaction? Why is it hard?
**Answer:**
A distributed transaction is a database transaction in which two or more network hosts are involved. 
*   **The Problem:** In a monolithic app with one database, you just use `BEGIN; UPDATE A; UPDATE B; COMMIT;`. If anything fails, it all rolls back. In a microservices app, the Order Service has its own DB, and the Payment Service has its own DB. You cannot do a standard ACID commit across two independent databases.

## 2. Two-Phase Commit (2PC) Protocol
**Answer:**
A classic approach to distributed transactions that provides strong consistency. It requires a central "Coordinator" node.
*   **Phase 1: Prepare Phase:** The Coordinator asks all participating nodes (e.g., Order DB, Payment DB) "Are you ready to commit?". The nodes lock their resources and reply "Yes" or "No".
*   **Phase 2: Commit Phase:** If ALL nodes replied "Yes", the Coordinator sends a "Commit" message to all nodes. If ANY node replied "No" or timed out, the Coordinator sends a "Rollback" message to all nodes.
*   **Drawbacks:** It is a **blocking** protocol. While waiting for the Coordinator's final decision, databases hold row locks, reducing system throughput severely. It also introduces a single point of failure (the Coordinator). Rarely used in modern high-scale microservices.

## 3. The Saga Pattern
**Answer:**
The modern solution for distributed transactions. A Saga is a sequence of local transactions. Each local transaction updates the database and publishes a message or event to trigger the next local transaction in the saga.
*   **Compensating Transactions:** If a local transaction fails, the saga executes a series of "compensating transactions" that undo the changes made by the preceding local transactions.
*   **Two execution styles:**
    *   **Choreography (Event-Driven):** No central coordinator. Service A completes work, fires an event to Kafka. Service B listens, does its work, fires an event. If Service C fails, it fires a "Failure Event," and A and B listen for it to trigger their own undo logic. (Hard to track and debug).
    *   **Orchestration:** A central "Saga Orchestrator" service manages the workflow. It explicitly tells Service A to do X, then tells Service B to do Y. If Y fails, the orchestrator explicitly tells Service A to undo X. (Easier to manage, but adds coordination overhead).

## 4. Designing an Idempotent API
**Answer:**
**Idempotency** means that making multiple identical requests has the same effect as making a single request. This is absolutely critical in payment flows (e.g., what if the user clicks "Pay" twice, or the network drops the HTTP 200 OK so the client retries the request?).
*   **How to implement it:**
    *   **Idempotency Key:** The client generates a unique ID (UUID) for the specific action (e.g., `Charge-UUID-1234`) and passes it in an HTTP header (e.g., `Idempotency-Key: Charge-UUID-1234`).
    *   **The Server Logic:**
        1.  When a request hits the Payment Service, it first checks a fast storage (like Redis or a specific DB table) for the `Idempotency-Key`.
        2.  *If not found:* It creates a record `(Key: 1234, Status: IN_PROGRESS)`. It processes the payment. Once successful, it updates the record to `Status: COMPLETED, Response: <json>`. It then returns the response.
        3.  *If found and Status == IN_PROGRESS:* The server knows an identical request is currently being processed by another thread. It returns an HTTP 409 Conflict.
        4.  *If found and Status == COMPLETED:* The server returns the saved `Response: <json>` immediately without charging the credit card a second time.

## 5. Event Sourcing and CQRS
**Answer:**
Often used alongside Sagas for absolute auditability.
*   **Event Sourcing:** Instead of storing the *current state* of an entity, you store a sequence of *state-changing events*. To get the current bank account balance, you don't read a `balance = $100` column; you replay all deposit and withdrawal events (`+50, +200, -150`) to calculate the $100 balance. Provides a 100% reliable audit log.
*   **CQRS (Command Query Responsibility Segregation):** Separates the read model from the write model. The write model (Command) might process heavy business logic and append events to an Event Store. A background processor listens to these events and updates a highly optimized Read Database (Query view). This allows you to scale read and write databases independently.
