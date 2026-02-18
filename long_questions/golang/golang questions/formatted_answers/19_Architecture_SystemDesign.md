# 🟢 **361–380: Architecture and System Design**

### 361. How do you design a scalable Go application?
"I adhere to the **Shared Nothing Architecture**.

Each instance of my Go service is stateless.
Any state (User Session, Shopping Cart) is pushed to a persistent store (Postgres) or a fast cache (Redis).
This allows me to scale horizontally: if traffic doubles, I just double the number of pods. I also heavily use **Message Queues** (Kafka) to decouple components so a load spike becomes a backlog, not an outage."

#### Indepth
Statelessness is a spectrum. Even "stateless" services often have local caches (for repeated DB lookups). Ensure you have a **Cache Invalidation Strategy** (e.g., TTLs or Pub/Sub events) so that one pod doesn't serve stale data while others serve new data.

---

### 362. What is Clean Architecture in Go?
"It’s a layered approach that enforces the **Dependency Rule**: dependencies only point inward.

*   **Entities (Core)**: Pure business objects (`User`). No SQL tags, no JSON tags.
*   **Use Cases**: Application logic (`CreateUser`). Depends only on Entities.
*   **Adapters**: HTTP handlers, SQL implementations. Depend on Use Cases.
*   **Infrastructure**: The DB driver, the Web framework.
This ensures I can swap Postgres for MongoDB without changing a single line of my core business logic."

#### Indepth
Clean Architecture comes with a cost: **Boilerplate**. You will map `UserDTO` (API) to `User` (Domain) to `UserModel` (DB). For simple CRUD apps, this is overkill. Don't be afraid to use a "Vertical Slice Architecture" (Transaction Script) for simple features, and upgrade to Clean Architecture only for complex domains.

---

### 363. How do you implement Domain-Driven Design (DDD) in Go?
"I organize code by **Domain**, not by Layer.

Instead of `packages/models` and `packages/controllers`, I have:
`packages/billing` (contains its own models, logic, and repo interfaces).
`packages/inventory`.
I define **Aggregates** (Transaction Boundaries) and enforce that changes to an aggregate only happen through its methods. This prevents 'anemic domains' where logic leaks into the service layer."

#### Indepth
A common DDD mistake in Go is creating "God Aggregates". If your `User` aggregate contains `Orders`, `Payments`, and `Reviews`, you seal the database row for `User` every time you add a Review. Keep aggregates small. `Order` should be its own aggregate and reference `UserID` by ID, not by embedding.

---

### 364. What is the Hexagonal Architecture (Ports and Adapters)?
"It’s about decoupling the app from the technology.

**Ports**: Interfaces defined by the Core.
*   Driving Port: `OrderService` interface (called by HTTP).
*   Driven Port: `OrderRepository` interface (implemented by SQL).

**Adapters**: The glue code.
*   `HttpHandler` adapter plugs into the Driving Port.
*   `PostgresRepo` adapter plugs into the Driven Port.
This allows me to run the entire application logic in a unit test with a `MemoryRepo` adapter, in milliseconds."

#### Indepth
Hexagonal architecture makes **Contract Testing** easier. You can write a test suite for the `Repository` interface and run it against both the `MemoryRepo` (for unit tests) and the `PostgresRepo` (for integration tests). This ensures your fake and real implementations actually behave identically.

---

### 365. How do you handle configuration management in distributed systems?
"I treat config as **Immutable**.

I use environment variables (12-Factor App) injected by Kubernetes ConfigMaps.
I read them on startup using `viper` or `kelseyhightower/envconfig`.
I avoid 'hot reloading' config because it leads to split-brain states (half the pods have new config, half have old). If I need to change config, I deploy a new revision of the pod."

#### Indepth
Feature Flags are a form of dynamic config. Use a system like **LaunchDarkly** or **Unleash** to toggle features on/off without redeploying. This decouples "Deployment" (moving code) from "Release" (enabling features), allowing for safer rollouts.

---

### 366. How do you design for failure in Go systems?
"I assume the network is hostile.

1.  **Timeouts**: Every `http.Client` and `sql.DB` call has a strict context timeout.
2.  **Retries**: I use exponential backoff with jitter for transient errors.
3.  **Circuit Breakers**: If the Payment Service is failing 50% of the time, I stop calling it for 30s to let it recover.
4.  **Bulkheads**: I isolate thread pools so a slow feature doesn't starve the whole app."

#### Indepth
"Fail Open" vs "Fail Closed". If your Fraud Check service is down, do you block the user (Fail Closed) or let them pass (Fail Open)? For high-value transactions, Fail Closed. For low-stakes features (Recommendations), Fail Open. Document this decision explicitly.

---

### 367. What is the CQRS pattern and when to use it?
"**Command Query Responsibility Segregation**.

I split my application into two parts:
*   **Write Side (Command)**: Handles `CreateOrder`. Sticktly consistent, normalized data.
*   **Read Side (Query)**: Handles `GetOrderHistory`. Eventually consistent, denormalized (e.g., Elasticsearch or a specific Read View table).
I use it when the read load is massive (1000:1 read:write ratio) or when the query patterns are too complex for the normalized write schema."

#### Indepth
CQRS enables **polyglot persistence**. The Write model can be a normalized PostgreSQL (3rd Normal Form) for data integrity. The Read model can be a denormalized MongoDB document, or a Redis Hash, optimized specifically for the "User Profile" screen. You project events from SQL to Mongo asynchronously.

---

### 368. How do you implement Event Sourcing in Go?
"Instead of storing the 'Current State' (Balance: $50), I store the **Sequence of Events** (Deposited $100, Withdrew $50).

To get the balance, I replay the events: $0 + $100 - $50 = $50.
In Go, I append these events to an immutable log (Kafka/EventStore).
It gives me a perfect audit trail and allows 'Time Travel' debugging, but it adds massive complexity (snapshots, schema evolution), so I only use it for critical financial/audit systems."

#### Indepth
Event Sourcing has a "snapshots" problem. To get the current balance, you can't replay 1 million events every time. You verify a Snapshot every 100 events. The generic `Aggregate.Fold()` logic in Go is typically: `state = Snapshot + apply(events_since_snapshot)`.

---

### 369. How do you handle database migrations in a microservices architecture?
"Each service **owns** its schema.

Service A effectively checks out Service A's DB. Service B cannot touch it.
I run migrations (using `golang-migrate`) as an **Init Container** in Kubernetes.
When a pod starts, it upgrades the schema *before* the app boots.
I strictly write **Backward Compatible** migrations (Add Column -> Deploy Code -> Backfill -> Remove Column) to support Zero Downtime deployments."

#### Indepth
Distributed locks are dangerous for migrations. If two pods try to run migrations simultaneously, they might corrupt the DB. Use `advisory_locks` (Postgres) or a dedicated K8s Job `helm install --wait`. Never let the application pod run migrations on startup in a multi-replica setup.

---

### 370. What is the Strangler Fig pattern?
"It’s how I kill a legacy Monolith.

1.  Put a Proxy (API Gateway) in front of the Monolith.
2.  Write a new Go Microservice for *one* feature (e.g., Search).
3.  Route `/search` traffic to the new Go service.
4.  Keep `/users`, `/billing` going to the Monolith.
5.  Repeat until the Monolith is empty.
This minimizes risk compared to a 'Big Bang' rewrite."

#### Indepth
The hardest part of Strangler Fig is **Data Sync**. If the monolith and microservice share the same DB, you have a "Distributed Monolith". If they split the DB, you need a sync mechanism (CDC or Dual Write) until the migration is complete. Prefer splitting the DB early if possible.

---

### 371. How do you design API rate limiting for high traffic?
"I use the **Token Bucket** algorithm backed by Redis.

Each request decrements a counter in Redis (`INCR key` is atomic).
Keys are sharded by UserID or IP.
If the count > Limit, I return **HTTP 429**.
I prefer Redis because it works across distributed instances. For ultra-high scale (DDOS protection), I rely on a specialized WAF (Cloudflare) at the edge rather than Go middleware."

#### Indepth
Token bucket is bursty. If you want smooth traffic, use **Leaky Bucket**. If you want a hard cap, use **Fixed Window**. If you want a hard cap without boundary issues (the "double spike" at minute boundaries), use **Sliding Window Log**, though it's memory expensive. **Sliding Window Counter** is the best memory/accuracy trade-off.

---

### 372. How do you manage distributed sessions?
"I go **Stateless** with JWTs.

The session data is signed and stored in the token on the client side. The server just verifies the signature.
If I absolutely need server-side state (e.g., to revoke a login instantly), I store the Session ID in **Redis**.
I never store sessions in application memory because Sticky Sessions break load balancing and complicate autoscaling."

#### Indepth
JWTs cannot be revoked easily. If you need revocation (e.g., "Sign out all devices"), you must keep a blacklist of revoked JTI (JWT IDs) in Redis, which checks on every request. This re-introduces state, negating some benefits of stateless JWTs, but it's a necessary trade-off for security.

---

### 373. What is the Outbox Pattern?
"It solves the 'Dual Write' problem (Write to DB + Publish to Kafka).

1.  Start SQL Transaction.
2.  Insert `User`.
3.  Insert `Msg` into an `outbox` table in the *same* DB.
4.  Commit Transaction.
5.  A background poller reads the `outbox` table and publishes to Kafka.
This guarantees that **if** the user is created, the event **will** be published, assuming eventually consistency."

#### Indepth
Debezium is a popular tool for this. It reads the Postgres Write-Ahead Log (WAL) and streams changes to Kafka. This assumes you don't even need an explicit "Outbox" table if you treat the domain table changes as events, but the Outbox table is safer for keeping internal schema private.

---

### 374. How do you handle idempotency in API design?
"I require an **Idempotency-Key** header (UUID) for critical POST requests.

On the server:
1.  Check Redis: 'Have I seen this Key?'
2.  If yes: Return the *cached response* from the previous success. Do not execute logic.
3.  If no: Execute logic, save response to Redis, return it.
This ensures that if a client retries a payment due to a network timeout, we don't charge them twice."

#### Indepth
Idempotency keys should expire (e.g., 24 hours). Also, the check must be atomic. Use `SET key val NX EX 86400` in Redis. If it returns false, the key exists—fetch the old response. If true, proceed. This avoids the race condition of "Check -> Logic -> Set".

---

### 375. How do you design for eventual consistency?
"I embrace the lag.

In the UI, I use **Optimistic Updates** (show 'Done' immediately implies success locally).
In the backend, I use **Sagas**.
If Service A updates, it emits an event. Service B listens and updates.
If Service B fails, it emits a 'Failed' event, and Service A executes a compensation logic (undo).
I monitor the 'Replication Lag' to ensure 'Eventually' doesn't mean 'Tomorrow'."

#### Indepth
A common pattern is "Read Your Own Writes". When a user updates their profile, pin them to the master DB (or the updated replica) for a few seconds via a cookie. This ensures they don't see their old profile name immediately after changing it, which builds user trust.

---

### 376. What is the Sidecar pattern?
"It’s deploying a helper container in the same Pod as my Go app.

Example: **Envoy Proxy**.
My Go app talks to `localhost:80`. Envoy intercepts it, handles mTLS, retries, and tracing, then forwards it to the destination.
This creates a **Service Mesh**. It keeps my Go code clean of infrastructure concerns like cert rotation or circuit breaking."

#### Indepth
Sidecars have resource overhead. If your Go app uses 10MB RAM but the Envoy sidecar uses 100MB, you are wasting money. For simple setups, "Proxyless Service Mesh" (gRPC library directly talking to the control plane) is emerging as a lighter alternative.

---

### 377. How to implement a circuit breaker in Go?
"I use `gobreaker` or `hystrix-go`.

I wrap unreliable calls:
`cb.Execute(func() error { return http.Get(...) })`.
If failure rate > 50%, the breaker **trips** (Open).
Subsequent calls fail immediately (Fast Failure) without touching the network.
After a sleep window, it lets one request through (Half-Open). If it succeeds, it resets (Closed)."

#### Indepth
State persistence is tricky. If you have 10 pods, each has its own local circuit breaker. If the DB is down, all 10 pods must individually trip their breakers (letting some requests fail). Distributed circuit breakers (using Redis) exist but add latency; usually, local breakers are fine.

---

### 378. How do you handle cascading failures?
"I use **Load Shedding** and **Timeouts**.

If my CPU is > 90%, I start rejecting requests immediately (HTTP 503) to save the remaining capacity for in-flight requests.
I also use **Timeouts** everywhere. If the DB is slow, my API shouldn't hang forever; it should timeout fast and free up the thread.
This prevents a failure in a low-priority service from taking down the critical path."

#### Indepth
Implement **Priority Queues**. If the system is overloaded, drop "Background Sync" traffic but keep "User Checkout" traffic alive. This requires checking the "Task Priority" header at the ingress or middleware layer and maintaining separate semaphores for each tier.

---

### 379. What is API Gateway versus Service Mesh?
"**API Gateway** (North-South): The front door. Handles external clients, AuthN, Rate Limiting, and Routing to microservices.
**Service Mesh** (East-West): The internal wiring. Handles mTLS, Retries, and Tracing *between* microservices.
I use a Gateway (Kong) to let users in, and a Mesh (Istio) to let services talk to each other safely."

#### Indepth
API Gateways can handle **AuthN** (Who are you? - validating the JWT signature) but should delegate **AuthZ** (Are you allowed to do this? - checking scopes/roles) to the service or an external policy engine like OPA (Open Policy Agent). The Gateway is too coarse-grained for complex business rules.

---

### 380. How do you secure data in transit and at rest?
"**Transit**: TLS 1.3 everywhere.
External: HTTPS (Let's Encrypt).
Internal: mTLS (Mutual Auth) so services prove identity to each other.

**Rest**: Encryption at the storage layer (AWS EBS / RDS encryption).
For sensitive columns (SSN/PII), I use **Envelope Encryption**: I encrypt the data with a generated Data Key (DEK) and store the DEK encrypted by a Master Key (KMS) alongside the data."

#### Indepth
Key Rotation is essential. With envelope encryption, you only rotate the Master Key (KMS). You don't need to re-encrypt terabytes of database rows (which are encrypted with DEKs). You just re-encrypt the DEKs with the new Master Key if you want full rotation, or simply use the new Master Key for new data.
