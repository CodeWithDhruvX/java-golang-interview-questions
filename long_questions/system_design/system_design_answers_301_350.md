## 🔸 Modern Architecture Patterns (Questions 301-310)

### Question 301: What is an event-sourcing architecture? Where would you use it?

**Answer:**
*   **Concept:** Storing the *sequence of events* (changes) rather than the current state.
*   **State:** Reconstructed by replaying events.
*   **Use Cases:** Banking (Ledge: Deposit, Withdraw), Audit Trails, Legal scenarios where history is critical.
*   **Benefits:** Complete history, temporal queries ("Balance as of last Tuesday"), high write throughput.

### Question 302: How would you design a CQRS-based system?

**Answer:**
**Command Query Responsibility Segregation.**
*   **Command Side (Write):** Handles updates (`CreateOrder`). Optimised for writes/logic. (e.g., Event Sourcing + RabbitMQ).
*   **Query Side (Read):** Handles reads (`GetOrders`). Optimised for fast retrieval (e.g., Elasticsearch / Denormalized SQL View).
*   **Sync:** Async event projection from Command -> Query side.

### Question 303: How does hexagonal (ports and adapters) architecture work?

**Answer:**
*   **Core:** Business Logic sits in the center, independent of frameworks/DBs.
*   **Ports:** Interfaces defined by the Core (e.g., `RepositoryPort`).
*   **Adapters:** Implementations of Ports (e.g., `PostgresAdapter`, `RestApiAdapter`).
*   **Benefit:** Testability (Mock adapters easily) and flexibility (Swap MySQL for Mongo without touching Core logic).

### Question 304: What is a service mesh and when is it useful?

**Answer:**
*   **Definition:** An infrastructure layer (Sidecar proxies like Envoy) that handles service-to-service communication.
*   **Responsibilities:** Observability (Metrics), Traffic Control (Canary, Retries, Timeout), Security (mTLS).
*   **Useful When:** You have 50+ microservices and need consistent policy enforcement without duplicating code in every service.

### Question 305: Compare monorepo vs polyrepo in large-scale systems.

**Answer:**
*   **Monorepo (Google, Meta):**
    *   *Pros:* Atomic commits (change API + Clients together), Shared code visibility, Standardized tooling.
    *   *Cons:* Build times scale poorly, Large checkout size.
*   **Polyrepo (Netflix, Amazon):**
    *   *Pros:* Decentralized autonomy, Isolation.
    *   *Cons:* Dependency hell ("Diamond Dependency"), Synchronization overhead.

### Question 306: How do you design for graceful degradation?

**Answer:**
*   **Core vs Non-Core:** Identify critical paths (Checkout) vs optional paths (Recommendations).
*   **Fallback:** If "Recommendations" service is down, catch exception and return Empty List (or Cached list) instead of 500 Error.
*   **UI:** Show "Some features unavailable" instead of a blank page.

### Question 307: What are design considerations for zero-downtime deployments?

**Answer:**
*   **Rolling Updates:** Update K8s pods 1 by 1.
*   **Database:** Backward Compatibility.
    *   Don't rename columns. (Add new column -> Duplicate Write -> Migrates data -> Switch Read -> Delete old column).
*   **Health Checks:** Load Balancer must not route traffic until NEW version is literally ready (`readinessProbe`).

### Question 308: How to implement blue-green deployments?

**Answer:**
*   **Setup:** Two identical environments (Blue=Live, Green=Idle).
*   **Deploy:** Push v2 to Green. Test Green thoroughly.
*   **Switch:** Update Load Balancer to point to Green. (Instant switch).
*   **Rollback:** If v2 sucks, point LB back to Blue.

### Question 309: What is canary deployment and when to use it?

**Answer:**
*   **Process:** Deploy v2 to a small subset of users (1% -> 5% -> 25%).
*   **Filter:** Based on IP, UserID, or HTTP Header.
*   **Monitor:** Compare Errors/Latency of v2 vs v1.
*   **Use:** High-risk changes where tests might miss real-world edge cases.

### Question 310: How would you design a system for high observability?

**Answer:**
**Three Pillars:**
1.  **Logs:** Standardized JSON, containing `TraceID`. (ELK/Loki).
2.  **Metrics:** Prometheus (Infrastructure + Business metrics).
3.  **Traces:** Jaeger/OpenTelemetry (Request flow).
*   **Correlation:** Ability to click a metric spike -> Jump to Trace -> Jump to Logs.

---

## 🔸 Low-Level Design Scenarios (Questions 311-320)

### Question 311: Design a URL parsing library.

**Answer:**
*   **Component:** State Machine.
*   **States:** `Protocol`, `Host`, `Port`, `Path`, `Query`, `Fragment`.
*   **Logic:**
    1.  Read until `://` -> Protocol.
    2.  Read until `/` -> Host (+ Port if `:`).
    3.  Read until `?` -> Path.
    4.  Read until `#` -> Query Params (Split by `&`).
*   **Validation:** Check allowed characters (RFC 3986).

### Question 312: Design a rate limiter.

**Answer:**
*   **Algorithms:** Token Bucket, Leaky Bucket, Sliding Window Log.
*   **Data Structure:**
    *   Redis Hash: `key: UserID`, `val: {tokens: 5, last_refill: timestamp}`.
    *   Atomic/Lua Script: Calculate refill since `last_refill`, decrement token, save.

### Question 313: Build a cron job scheduler.

**Answer:**
*   **Min-Heap:** Store tasks ordered by `NextExecutionTime`.
*   **Thread:** Sleep until `Heap.Peek().Time`.
*   **Execution:** Pop task -> Run (in ThreadPool) -> Calculate Next Time -> Push back to Heap.
*   **Distribution:** Use Leader Election (Etcd). Only Leader executes; if Leader dies, Follower takes over.

### Question 314: Design a thread-safe in-memory cache.

**Answer:**
*   **Storage:** `ConcurrentHashMap` (Java) or `sync.Map` (Go).
*   **Eviction (LRU):** Doubly Linked List + HashMap.
    *   *Access:* Move node to Head.
    *   *Insert:* Add to Head. If size > Max, remove Tail.
    *   *Locking:* Read-Write Lock (`RWMutex`). Allow multiple readers, one writer (for LRU updates).

### Question 315: Implement a retry mechanism with exponential backoff.

**Answer:**
*   **Loop:** `for i in 0..MaxRetries`:
    *   Call Service.
    *   If Success -> Break.
    *   If Generic Error -> Throw.
    *   If Transient Error (503/Timeout) -> Wait `Base * 2^i + Jitter`.
*   **Jitter:** Add random `±100ms` so all failed clients don't retry at the EXACT same millisecond.

### Question 316: Design a key-value store from scratch.

**Answer:**
*   **In-Memory:** Hash Map.
*   **Persistence (WAL):** Write-Ahead Log. Append operation (SET K V) to file before updating memory.
*   **Snapshot:** Periodically dump Map to disk.
*   **LSM Tree (LevelDB/RocksDB):** Optimized for write heavy loads. MemTable -> Flush to SSTable on disk.

### Question 317: Implement a feature toggle mechanism.

**Answer:**
*   **Interface:** `bool isEnabled(String featureName, User context)`.
*   **Config source:** Database / Redis / JSON file.
*   **Caching:** Cache config locally for 1 min.
*   **Strategy:** Boolean toggle, Percentage rollout (`hash(user) % 100 < 20`), or User Whitelist.

### Question 318: Design a library to handle pagination in APIs.

**Answer:**
*   **Input:** `PageRequest(page, size, sort)`.
*   **Output:** `PageResponse(content[], totalElements, totalPages, hasNext)`.
*   **Logic:**
    *   SQL: `LIMIT size OFFSET (page * size)`.
    *   Check Max Page Size (security).

### Question 319: Create an in-memory log aggregator.

**Answer:**
*   **Buffer:** Circular Buffer (Ring Buffer) of size N strings. O(1) Write. Oldest overwritten.
*   **Concurrent:** Use spin-lock or atomic index.
*   **Flush:** Background thread wakes up every X seconds -> Dumps buffer to Disk/Network.

### Question 320: Design a UUID generation service.

**Answer:**
*   **Type 4 (Random):** 122 random bits. Collision theoretical but negligible.
*   **Snowflake (Twitter):** 64-bit ID (Sortable by time).
    *   `1 bit` (Unused) | `41 bits` (Timestamp) | `10 bits` (MachineID) | `12 bits` (Sequence).
    *   Guarantees unique, chronologically sortable IDs distributed across machines.

---

## 🔸 Enterprise Systems (Questions 321-330)

### Question 321: Design an SSO (Single Sign-On) solution.

**Answer:**
(e.g., CAS / SAML / OIDC).
1.  User visits App A. Not logged in.
2.  Redirect to Identity Provider (IdP).
3.  User logs in at IdP. IdP sets global session cookie.
4.  IdP redirects back to App A with `AuthCode`.
5.  App A exchanges Code for Token.
6.  User visits App B. Redirect to IdP. IdP sees cookie -> Auto-redirects back with Code (No login screen).

### Question 322: How would you build a document management system?

**Answer:**
*   **Storage:** S3 (Blobs).
*   **Metadata:** Postgres (Author, Tags, Folder Structure).
*   **Versioning:** Store `DocID`, `VersionID`. S3 supports versioning natively.
*   **Search:** Tika (extract text from PDF) -> Elasticsearch.
*   **Permissions:** ACLs per folder/file.

### Question 323: Design a workflow engine.

**Answer:**
*   **Definition:** BPMN / JSON DAG.
*   **Execution:**
    *   **Orchestrator:** Maintains state of process instance (`CurrentStep: Approval`).
    *   **Queue:** Push "Job: Approval" to Queue.
    *   **Worker:** completing job notifies Orchestrator.
    *   **State:** Use Event Sourcing to track history.

### Question 324: Design a multi-department HRMS backend.

**Answer:**
*   **Tenant:** Multi-tenant or Single-tenant per Corp.
*   **Modules:** Payroll, Leave, Profile.
*   **Access:** Complex RBAC (`Manager` sees Team Salaries, `Employee` sees own).
*   **Audit:** Strict logging of all changes.

### Question 325: Build a financial transaction ledger.

**Answer:**
*   **Immutability:** INSERT ONLY. Never UPDATE.
*   **Double Entry:** Every transaction has 2 legs (Debit A, Credit B). Sum must be 0.
*   **Consistency:** ACID Transaction (Postgres).
*   **Checksum:** Store hash of row + hash of previous row (Blockchain-lite) to detect tampering.

### Question 326: How would you handle multi-currency accounting?

**Answer:**
*   **Storage:** `amount` (Decimal), `currency` (String).
*   **Exchange:** Maintain `ExchangeRate` table with historical rates.
*   **Reporting:** Convert all to Base Currency (USD) using the rate *at the time of transaction*.
*   **Rounding:** Use Banker's Rounding.

### Question 327: How would you enforce access control in enterprise apps?

**Answer:**
*   **Pattern:** Policy Based Access Control (PBAC) / ABAC (Attribute Based).
*   **Policy:** `CanView(User, Doc) IF User.Dept == Doc.Dept AND Time < 5PM`.
*   **Engine:** Open Policy Agent (OPA). Call OPA with JSON inputs, get Allow/Deny decision.

### Question 328: Design a time tracking and approval system.

**Answer:**
*   **Timesheet:** JSON blob of hours per day/task.
*   **Submission:** State defaults to `Submitted`.
*   **Notification:** Email Manager.
*   **Approval:** Manager API (`POST /timesheet/123/approve`). Changes state to `Approved`. Locks record (ReadOnly).

### Question 329: How would you implement business process orchestration?

**Answer:**
*   **Saga Pattern:** Sequence of local transactions.
*   **Orchestrator (Camunda/Temporal):**
    *   Call Service A (Reserve Hotel).
    *   Call Service B await result (Reserve Flight).
    *   If B fails -> Execute Compensating Transaction A (Cancel Hotel).

### Question 330: How would you version business rules?

**Answer:**
*   **Rules Engine:** (Drools).
*   **Versioning:**
    *   Key: `tax_rule_v1` vs `tax_rule_v2`.
    *   Effective Dates: `Rule: Tax=20%, Start=2023-01-01, End=2023-12-31`.
*   **Selection:** Query rule valid for `TransactionDate`.

---

## 🔸 Trade-offs & Decision Making (Questions 331-340)

### Question 331: When do you denormalize data?

**Answer:**
*   **Read-Heavy:** When reading involves 5+ Joins and is slow.
*   **Example:** Storing `AuthorName` in `Book` table instead of joining `Author` table.
*   **Trade-off:** Write Complexity. Updating Author name requires updating all their Books.

### Question 332: When to choose SQL vs NoSQL?

**Answer:**
*   **SQL:** Structured data, Complex Joins, ACID required (Financial), Ad-hoc queries.
*   **NoSQL:** Unstructured/Flexible schema (Documents), Extreme Scale (Cassandra), High Write throughput, Simple Access Patterns (Key-Value), Graph relationships.

### Question 333: How do you choose between vertical and horizontal scaling?

**Answer:**
*   **Vertical (Scale Up):** Simpler (no code change). Good for moderate load. Limited by hardware capabilities ($$).
*   **Horizontal (Scale Out):** Complex (sharding, distributed state). Unlimited theoretical scale. Good for massive load.
*   **Strategy:** Start Vertical. Switch to Horizontal when cost/performance bottleneck hits.

### Question 334: When to use eventual consistency?

**Answer:**
*   **Scenario:** High Availability is more important than immediate Consistency (CAP Theorem - AP).
*   **Examples:** Social Feed (OK if friend's post appears 5s late), Analytics, Search Index.
*   **Avoid:** Payments, Inventory (Overselling).

### Question 335: When is strong consistency a must?

**Answer:**
*   **Scenario:** Correctness is critical. (CAP Theorem - CP).
*   **Examples:** Bank Transfer, Ticket Booking (prevent double booking), Leader Election (only 1 leader).

### Question 336: When would you prefer gRPC over REST?

**Answer:**
*   **Internal Microservices:** Lower latency (Protobuf < JSON), multiplexing (HTTP/2), type safety.
*   **Client SDKs:** Auto-generate client code.
*   **Streaming:** Bi-directional streaming requirements.
*   **Note:** Harder for external public APIs (browser support issues).

### Question 337: Should you build or buy a third-party tool?

**Answer:**
*   **Core Competency:** Build if it's your specific "Secret Sauce" (e.g., Uber's matching algo).
*   **Generic:** Buy (e.g., Auth0 for Auth, Stripe for Payments, Twilio for SMS). "Don't reinvent the wheel unless you need a rounder wheel."

### Question 338: When to go for cloud-native vs self-hosted?

**Answer:**
*   **Cloud-Native (Serverless/Managed):** Speed to market, low maintenance. (Startups).
*   **Self-Hosted (EC2/On-Prem):** Compliance (Govt/Medical), predictable cost at massive scale (Dropbox moved off AWS to save money), unique hardware needs.

### Question 339: How do you decide database sharding strategy?

**Answer:**
*   **Key Selection:** Critical. Must distribute load evenly.
    *   *Bad:* Shard by `Date` (All write traffic hits "Today's" shard).
    *   *Good:* Shard by `UserID` (Random distribution).
*   **Re-sharding:** Consistent Hashing allows adding nodes with minimal data movement.

### Question 340: How to decide between pub/sub and queue?

**Answer:**
*   **Queue (Point-to-Point):** 1 Producer -> 1 Consumer (of many workers). Load balancing jobs. Each job processed ONCE.
*   **Pub/Sub (Broadcast):** 1 Producer -> Many Subscribers. Each subscriber gets a COPY. Decoupling services (e.g., Order Created -> triggers Email Service, Inventory Service, Analytics).

---

## 🔸 Testing, QA & Simulation (Questions 341-350)

### Question 341: How do you test distributed systems?

**Answer:**
*   **Unit/Integration Tests:** Standard.
*   **Contract Tests (Pact):** Ensure Service A sends JSON that Service B expects.
*   **End-to-End:** Smoke tests on Staging.
*   **Chaos Testing:** (Production) Ensure resilience.

### Question 342: How to simulate network partitions?

**Answer:**
*   **Tools:** `iptables` (Linux firewall), `tc` (Traffic Control).
*   **Action:** Drop all packets from IP range of Region B.
*   **Toxiproxy:** A proxy that sits between services specifically to simulate latency/cuts.

### Question 343: How would you test failover in production?

**Answer:**
*   **Game Day:** Scheduled event.
*   **Action:** Manually terminate the Primary Database instance.
*   **Verification:** Measure time until App recovers. Verify no data loss. Verify alerts fired.

### Question 344: How to design test cases for a messaging system?

**Answer:**
*   **Ordering:** Send 1, 2, 3. Verify received 1, 2, 3.
*   **Duplication:** Send 1 message. Configuring consumer to crash before Ack. Restart. Verify message is redelivered.
*   **Volume:** Send 10k/sec. Verify lag/latency.

### Question 345: How do you mock time-sensitive features?

**Answer:**
*   **Abstract Clock:** Don't call `Date.now()` directly. Inject a `Clock` service.
*   **Test:** Inject a `FixedClock` or `FastForwardClock`.
*   **Scenario:** Testing "Expire after 24h". Set mock clock to `Now + 25h`. Verify expiry.

### Question 346: How would you test eventual consistency?

**Answer:**
*   **Polling:** Write data.
*   **Loop:** Read endpoint every 100ms.
*   **Assert:** Data appears within X seconds (SLA).
*   **Fail:** If data not present after Timeout.

### Question 347: Design a chaos testing module.

**Answer:**
*   **Agent:** Runs on sidecar.
*   **Config:** `chaos_enabled: true`, `failure_rate: 1%`.
*   **Interceptor:** Intercepts HTTP outgoing calls.
*   **Logic:** `if (rand() < 0.01) return 500 error`.

### Question 348: How to test an analytics system for accuracy?

**Answer:**
*   **Gold Standard:** Create a small controlled dataset (10 events) where the result (Sum=50) is manually calculated.
*   **Pipeline:** Push these 10 events.
*   **Verify:** Query dashboard. Assert Sum == 50.

### Question 349: What is the role of synthetic monitoring?

**Answer:**
*   **Proactive:** Robot user (Selenium script) runs periodically (every 5 min) against Prod.
*   **Scenario:** Log in -> Add to Cart -> Checkout.
*   **Goal:** Detect broken flows before real customers complain.

### Question 350: How do you simulate high concurrency?

**Answer:**
*   **Load Testing Tools:** JMeter, K6, Locust.
*   **Distributed Load:** Run the tool on 10 EC2 instances to generate 100k concurrent users.
*   **Target:** Hit the Staging environment API/Websocket.
