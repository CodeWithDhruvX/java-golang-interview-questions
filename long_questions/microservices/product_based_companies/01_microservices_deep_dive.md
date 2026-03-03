# 🏗️ Microservices — Product-Based Companies Deep Dive

> **Level:** 🔴 Senior
> **Asked at:** Amazon, Google, Flipkart, Uber, Swiggy, Netflix, Razorpay, PhonePe

---

## Q1. In a microservices architecture, how do you handle distributed transactions across multiple services when 2PC is not an option?

"Two-Phase Commit (2PC) is a blocking protocol that doesn't scale well in distributed systems due to its synchronous nature and reliance on a central coordinator. Instead, for product-scale systems, I use the **Saga Pattern**.

A Saga is a sequence of local transactions. Each local transaction updates the database and publishes a message or event to trigger the next local transaction in the saga.

If a local transaction fails because it violates a business rule, the saga executes **compensating transactions** to undo the impact of the preceding local transactions.

Two ways to implement Sagas:
**1. Choreography:** Each service publishes events, and other services subscribe to those events. There is no central coordinator.
*Pros:* Simple for a few services, no SPOF.
*Cons:* Complex to track flow, risk of cyclic dependencies.

**2. Orchestration:** An orchestrator (like an AWS Step Function or Axon Saga manager) tells participating services what local transactions to execute.
*Pros:* Centralized logic, easier to understand and test, avoids cyclic dependencies.
*Cons:* Orchestrator can become a design SPOF (not necessarily infrastructure, but logic coupling).

For complex flows (like an e-commerce order: Order -> Payment -> Inventory -> Shipping), Orchestration is almost always preferred so the Order Service (or a dedicated Order Saga Orchestrator) controls the entire lifecycle."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Fintech (Razorpay, PhonePe) and E-commerce (Flipkart, Amazon) where transaction integrity is critical but 2PC is too slow.

#### Indepth
**Compensating Transactions must be Idempotent:** If the network fails during a compensation (e.g., refunding a payment), the orchestrator will retry. If the refund API isn't idempotent, the customer gets refunded twice. You must pass the original Transaction ID/Saga ID and check if it has already been compensated before executing the logic.

---

## Q2. How do you prevent dual-write problems when a microservice needs to update its database and publish an event to Kafka?

"The dual-write problem occurs when a service needs to do two things: write to its local database and publish a message to a broker (like Kafka). Since these involve two different systems without a shared transaction manager, one might succeed while the other fails, leading to an inconsistent state.

I solve this using the **Transactional Outbox Pattern**.

Instead of writing to the DB and sending to Kafka directly, the service does the following in **one single local database transaction**:
1. Update the business entity (e.g., `Order` table).
2. Insert an event describing the change into an `Outbox` table.

Since both happen in the same local DB transaction, they are atomic. Either both succeed, or both roll back.

Then, a separate background process reads the `Outbox` table and publishes the events to Kafka. Once published, the outbox record is marked as processed or deleted.

To implement the background process securely and instantly, I use **Change Data Capture (CDC)** tools like Debezium. Debezium reads the database transaction log (like MySQL binlog or Postgres WAL) and streams the outbox inserts directly to Kafka."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Swiggy, Zepto — systems heavily relying on event-driven architectures.

#### Indepth
**Listen to Yourself Pattern:** A variation where the service *only* writes to the event bus (Kafka), and then its own consumer reads the event back from Kafka to update its local database. This makes Kafka the single source of truth. However, it means local reads immediately after a write will be stale (eventual consistency), which complicates UI logic requiring immediate feedback.

---

## Q3. Explain how you would implement resilient communication between microservices. What happens when a downstream service is struggling?

"If Service A calls Service B, and Service B becomes slow or unresponsive, Service A's threads will block waiting for a response. If traffic is high, Service A will exhaust its thread pool and also crash. This is a cascading failure.

To prevent this, I implement three primary resilience patterns:

**1. Timeouts and Retries with Exponential Backoff:**
Never make an external call without a hard timeout. If a call fails, retry, but use exponential backoff (wait 1s, then 2s, 4s...) to avoid hammering a struggling service. Add *Jitter* (randomness) to the backoff to prevent thundering herds when multiple instances retry simultaneously.

**2. Circuit Breaker:**
I use libraries like Resilience4j. If Service B fails consistently (e.g., >50% error rate in a sliding window), the circuit breaker 'opens'. Service A stops calling Service B entirely for a cooldown period, failing fast. This saves Service A's threads and gives Service B time to recover.

**3. Bulkheads:**
I isolate resources. If Service A talks to Service B and Service C, I allocate a specific thread pool (or semaphore limit) just for Service B. If B is dead, only B's thread pool is exhausted, allowing Service A to continue serving requests that only rely on Service C."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Amazon, Hotstar — companies dealing with massive scale and partial failure management.

#### Indepth
**Fallbacks:** When a circuit breaker is open, what do you return? You should have a fallback strategy. E.g., if the personalized recommendation service is down, the fallback returns a cached list of globally trending items. It's a degraded experience, but better than an error page.

---

## Q4. How do you handle database scaling in a microservices environment where multiple services might need access to related data?

"The golden rule of microservices is **Database per Service**. Services must not share databases; they must only interact via APIs. Shared databases lead to tight coupling, preventing independent schema evolution and scaling.

However, data isolation creates challenges for querying and joining data. My approaches for scaling and querying isolated data:

**1. CQRS (Command Query Responsibility Segregation) for querying across services:**
If the Frontend needs a dashboard combining user data (User Service) and order history (Order Service), I don't make them run complex distributed joins on the fly. 
Instead, I create a dedicated 'View' or 'BFF' (Backend for Frontend) database. As the User and Order services emit state changes via Kafka, a consumer updates this read-optimized Elasticsearch or NoSQL database. The dashboard queries this read model instantly.

**2. Database Sharding for write scale:**
If a single service's database (like the Order DB) grows too large for one master node to handle writes, I apply **Sharding**. I partition the data across multiple database instances based on a Shard Key (e.g., `user_id`). 
Consistent Hashing is used to distribute the data evenly and minimize data movement when adding new shards."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** High-growth scale-ups (CRED, Groww) and large e-commerce dealing with massive data sets.

#### Indepth
**API Composition (API Gateway Pattern):** For simple queries that don't warrant the overhead of CQRS infrastructure, an API Gateway or GraphQL layer can make parallel synchronous calls to the User and Order services and aggregate the JSON in memory before returning to the client. This is fine for simple lookups but scales poorly for large datasets or complex filtering/sorting.

---

## Q5. Describe your approach to testing a massive microservices architecture. How do you catch integration issues without spinning up 50 services locally?

"Testing microservices via end-to-end (E2E) UI tests is a known anti-pattern at scale. They are brittle, slow, and non-deterministic ("flaky"). You cannot spin up 50 services reliably on a developer's laptop or a standard CI pipeline.

Instead, I rely on the **Testing Pyramid skewed towards Contracts**:

**1. Unit Testing:** Extensive testing of business logic in isolation within each service. Fast and reliable.

**2. Component Testing / Integration Testing (Local):** Testing the service interacting with its *own* database (using Testcontainers to spin up ephemeral Postgres/Redis via Docker).

**3. Consumer-Driven Contract Testing (CDC):** This is the replacement for E2E tests. I use a tool like **Pact**. 
The 'Consumer' (Service A) writes a test defining the exact Request it will send and the Response it expects. This generates a 'Contract' definition file.
The 'Provider' (Service B) downloads this Contract in its CI pipeline and runs a test against it. If Service B changes an API field that Service A relies on, Service B's build fails *before* deployment.
This guarantees API compatibility without ever spinning up Service A and B at the same time.

**4. Production Testing / Observability:** You can't catch everything in QA. I use Canary Deployments (rolling out to 5% of traffic), feature flags, and heavy observability (tracing, metrics) to detect and roll back issues instantly in production."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Tech-first companies with mature CI/CD pipelines aiming for multiple deployments per day.

#### Indepth
**Distributed Tracing (e.g., Jaeger, OpenTelemetry):** When a production issue does happen, logs aren't enough because a single user request spans 8 microservices. Tracing injects a `TraceId` at the API Gateway, passing it in HTTP headers to every downstream service. This allows visualizing the exact latency and failure point of a specific request path across the entire fleet.

---

## Q6. What is a Service Mesh, and when should a company adopt one?

"A Service Mesh (like Istio or Linkerd) is a dedicated infrastructure layer for managing service-to-service communication. It operates using a **sidecar proxy** deployed alongside every microservice instance.

Instead of developers writing code for retries, circuit breaking, mTLS (mutual TLS security), and tracing context propagation inside their applications (using libraries like Spring Cloud or Netflix OSS), the Service Mesh handles it at the network layer.

**When to adopt:**
You *should not* adopt a service mesh when you only have 10-20 microservices. The operational complexity of running Istio is massive. 

You *should* adopt a service mesh when:
1. You have heterogeneous tech stacks (Go, Node, Java). Rewriting resilience libraries in every language is impossible; a sidecar solves it uniformly.
2. You have strict compliance requirements demanding mTLS encryption for all internal network traffic. The mesh handles certificate rotation automatically.
3. You have hundreds of microservices where managing traffic routing (canary deployments, blue/green) and observability at scale becomes a dedicated platform team issue."

#### 🏢 Company Context
**Level:** 🔴 Senior/Principal | **Asked at:** Platform Engineering roles, Google, companies transitioning to Kubernetes-first infrastructure.

#### Indepth
**API Gateway vs. Service Mesh:** An API Gateway manages **North-South** traffic (external Internet -> internal cluster). It handles edge security, rate limiting by user, and aggregation. A Service Mesh manages **East-West** traffic (Internal Service A -> Internal Service B). It handles internal routing rules, service mTLS, and internal retries.
