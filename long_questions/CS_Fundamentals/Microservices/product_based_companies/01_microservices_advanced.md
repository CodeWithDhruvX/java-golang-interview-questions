# ⚙️ Microservices — Advanced Interview Questions (Product-Based Companies)

This document covers advanced microservices architecture concepts for product-based company interviews (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto). Targeted at 3–10 years of experience rounds.

---

### Q1: How do you handle distributed transactions? Explain the Saga Pattern.

**Answer:**
Because microservices use decentralized databases (Database-per-service), traditional ACID transactions (like 2-Phase Commit / 2PC) cannot be used efficiently due to locking overhead and CAP theorem constraints.

**The Saga Pattern:**
A Saga is a sequence of **local transactions**. Each local transaction updates the database and publishes a message/event to trigger the next local transaction in the saga.

**If a step fails:** The saga executes a series of **compensating transactions** to undo the changes made by the preceding local transactions.

**Two implementation styles:**

**1. Choreography (Event-based):**
- No central coordinator. Services listen to events and react.
- Order Service publishes `OrderCreated`.
- Payment Service listens, charges credit card, publishes `PaymentSuccessful`.
- Inventory Service listens, reserves stock.
- *Pros:* Decoupled, simple for small workflows.
- *Cons:* Hard to track the overall status. Complex sagas become a "tangled web" of events.

**2. Orchestration (Command-based):**
- A central **Saga Orchestrator** manages the workflow flow.
- Orchestrator tells Payment Service to `ChargeCard`.
- Payment replies with `CardCharged`.
- Orchestrator tells Inventory to `ReserveStock`. If it fails, Orchestrator commands Payment to `RefundCard`.
- *Pros:* Clear state tracking, easier to understand complex workflows, centralized error handling.
- *Cons:* Orchestrator can become a single point of failure/bottleneck if not designed well.

---

### Q2: What is CQRS and Event Sourcing? Why are they often used together?

**Answer:**

**CQRS (Command Query Responsibility Segregation):**
Separates the data mutation operations (Commands: POST, PUT, DELETE) from the data retrieval operations (Queries: GET).
- **Command Side:** Optimized for complex validation and writes.
- **Query Side:** Uses a completely different, denormalized read-model database (e.g., Elasticsearch, Redis) optimized purely for fast reads. No joins required.
- The Query DB is asynchronously updated via events from the Command DB.

**Event Sourcing:**
Instead of storing the *current state* of an entity, you store an append-only log of *all events* that altered the entity's state.
- Traditional: Update `balance = 100` to `balance = 150`. Old state is lost.
- Event Source: Store `[AccountCreated(0), Deposited(100), Deposited(50)]`.
- The current state is derived by "replaying" the event log.

**Why used together:**
If the datastore is an append-only log of events (Event Sourcing), querying current state is extremely slow (you have to replay all events).
Therefore, you **must apply CQRS**:
- The event log serves as the Command/Write model.
- Dedicated materialized views (Query/Read models) are continually built by consuming the event log, providing O(1) read access.

---

### Q3: Explain Distributed Tracing and MDC (Mapped Diagnostic Context).

**Answer:**
In microservices, a single user request can propagate across 10 different services. If an error occurs, finding where the chain broke in scattered logs is impossible.

**Distributed Tracing (e.g., Jaeger, OpenTelemetry, Zipkin):**
1. When a request hits the API Gateway, it generates a unique **Trace ID**.
2. As the request moves through services (via HTTP headers like `X-B3-TraceId` or W3C `traceparent`), the Trace ID is preserved.
3. Each operation within a service generates a **Span ID** (linked to the Trace ID).
4. All logs and metrics are exported to a tracing backend.
5. SREs can view a Waterfall chart showing exactly how long each hop took and where failures occurred.

**MDC (Mapped Diagnostic Context):**
A concept in logging frameworks (Logback, Log4j). It allows you to inject the Trace ID into the thread context.
- Once injected, *every* log statement emitted by that thread automatically includes the `[TraceID]` prefix.
- Vital for querying centralized log aggregators (Elasticsearch, Splunk) — searching for `trace_id=12345` pulls up application logs from ALL services involved in that request.

---

### Q4: What is the Strangler Fig pattern for migrating monoliths?

**Answer:**
The **Strangler Fig Pattern** is a strategy for safely migrating a legacy monolithic application to microservices without doing a risky "Big Bang" rewrite.

**How it works:**
1. Place a routing layer (API Gateway) in front of the legacy monolith. Initially, 100% of traffic goes to the monolith.
2. Identify a specific, isolated capability (e.g., User Profiles).
3. Build the new "Profile Microservice" alongside the monolith, with its own database.
4. Update the API Gateway to route `/api/profiles` traffic to the new microservice, while all other traffic still goes to the monolith.
5. Gradually, piece by piece, new microservices "strangle" the monolith.
6. Eventually, the monolith handles zero traffic and can be safely decommissioned.

**Challenges:** Data synchronization between the legacy database and the new microservice databases during the migration phase (often handled via CDC - Change Data Capture tools like Debezium).

---

### Q5: How do you design an API rate limiter for a distributed gateway? (Algorithms)

**Answer:**
Rate limiting protects backend services from abuse and SLA violations.

**1. Token Bucket:**
- A bucket holds up to N tokens. Tokens are added at a constant rate.
- Each request removes 1 token. If the bucket is empty, request is dropped (429 Too Many Requests).
- *Pros:* Allows for short bursts of traffic. Easy to implement in Redis.

**2. Leaky Bucket:**
- Requests enter a queue (bucket). They are processed at a strictly constant rate.
- If queue is full, new requests leak (drop).
- *Pros:* Smooths out traffic. *Cons:* Does not allow bursts.

**3. Fixed Window Counter:**
- Keep a counter for a fixed time window (e.g., 00:00 to 00:01).
- Increments on request. If threshold hit, block until next window.
- *Flaw:* Bursting at window edges (2x traffic limit if spikes occur at 00:00:59 and 00:01:01).

**4. Sliding Window Log:**
- Store the timestamp of every request in a Redis Sorted Set.
- Remove timestamps older than 1 minute. Count the remaining.
- *Pros:* Perfectly accurate. *Cons:* High memory footprint.

**5. Sliding Window Counter (Best balance):**
- Combines Fixed Window and Sliding Log. Tracks counters per fixed window, but calculates a weighted average based on overlapping time.
- Implementation: Redis + Lua scripts to ensure read/update/write is atomic.

---

*Prepared for technical rounds at product-based companies (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay, Groww).*
