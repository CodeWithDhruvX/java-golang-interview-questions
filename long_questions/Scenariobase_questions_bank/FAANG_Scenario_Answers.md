# FAANG-Level Scenario Answers (71-100)

## 🔴 Distributed Systems & Scale Scenarios

### Question 71: A cache stampede brings down your database. How do you prevent it?

**Answer:**
1.  **Locking:** Use a mutex (lock) so only *one* process fetches data from DB to populate cache. Others wait.
2.  **Probabilistic Early Expiration:** Refresh the cache entry *before* it strictly expires (e.g., if TTL is 60s, fetch at 50s with 10% probability).
3.  **Singleflight:** Use a "Singleflight" pattern (Go/Java libs) to coalesce duplicate requests for the same key.

---

### Question 72: A leader node crashes during job processing. What happens next?

**Answer:**
1.  **Detection:** Followers detect missing heartbeats (Lease Expiration).
2.  **Election:** A new Leader Election is triggered (Raft/Paxos/Bully Algorithm).
3.  **Recovery:** The job must be re-processed (if not checkpointed) or resumed from the last checkpoint (WAL).
4.  **Consistency:** Ensure the old leader (if it comes back) steps down (fencing token).

---

### Question 73: Network partition splits your cluster. How does your system behave?

**Answer:**
1.  **CAP Theorem:** Decide between CP (Consistency) or AP (Availability).
2.  **CP System (e.g., ZooKeeper):** The minority partition stops accepting writes to prevent Split Brain.
3.  **AP System (e.g., Cassandra):** Both sides accept writes. You must handle conflict resolution (Vector Clocks, LWW) when the partition heals.

---

### Question 74: Hot keys overload a single cache node. How do you fix it?

**Answer:**
1.  **Local Cache:** Add an in-memory (L1) cache on the application instances for that specific hot key.
2.  **Key Splitting:** Replicate the key. Store `key_1`, `key_2`, `key_3` on different nodes. Randomly read from one.
3.  **Routing:** Use a sidecar/proxy to coalesce requests.

---

### Question 75: A global service shows higher latency in one region. How do you debug?

**Answer:**
1.  **Geo-DNS:** Are users being routed to the wrong datacenter?
2.  **Connectivity:** Is the trans-oceanic cable/link optimized? Check packet loss.
3.  **Dependency:** Is a local dependency (e.g., EU-West DB) overloaded or slow?
4.  **Deployment:** Did a bad canary deploy land in that region specifically?

---

### Question 76: Event consumers fall behind producers. What do you do?

**Answer:**
1.  **Scale Consumers:** Increase the number of Consumer Groups/Threads (up to the number of Partitions).
2.  **Backpressure:** Slow down the producer/API ingestion.
3.  **Partitioning:** Increase Kafka partitions to allow more parallel consumers.
4.  **Prioritization:** Move slow/large messages to a separate "Slow Lane" topic.

---

### Question 77: Duplicate messages appear in a queue. How do you handle idempotency?

**Answer:**
1.  **Idempotency Key:** Every message has a unique ID (UUID).
2.  **Deduplication:** Before processing, check DB/Redis if `ID` was already processed.
3.  **Atomic Operation:** Use `INSERT IGNORE` or conditional updates in DB (`UPDATE ... WHERE status = 'PENDING'`).

---

### Question 78: Exactly-once processing is required. How do you design for failures?

**Answer:**
1.  **Impossible?** Pure "Exactly-Once" delivery is impossible (Two Generals Problem). We simulate it.
2.  **Idempotent Consumer:** Effectively Example-Once processing. (See Q77).
3.  **Transactional Outbox:** Write to DB and Outbox Table in one transaction. Publisher reads Outbox.
4.  **Kafka Transactions:** Use `enable.idempotence=true` and transactional producers/consumers.

---

### Question 79: Redis goes down in a critical path. How does your system recover?

**Answer:**
1.  **Circuit Breaker:** Detect failure, stop calling Redis.
2.  **Fallback:** Fetch directly from DB (if DB can handle the load).
3.  **Local Cache:** Use in-memory cache as a temporary degradation.
4.  **Sentinel/Cluster:** Redis Sentinel should auto-promote a Replica to Master.

---

### Question 80: Distributed lock causes throughput drop. How do you redesign?

**Answer:**
1.  **Granularity:** Lock specific resource (Row ID) instead of the whole object/table.
2.  **Optimistic Locking:** Remove the lock. Use Version numbers (`UPDATE ... WHERE version=v`). Retry on failure.
3.  **Sharding:** Shard the lock manager (if using Redis/Zookeeper) to distribute load.

---

### Question 81: Schema change breaks older services. How do you deploy safely?

**Answer:**
1.  **Expand-Contract:**
    *   **Phase 1:** Add new column/table (Code writes to both).
    *   **Phase 2:** Backfill old data.
    *   **Phase 3:** Switch code to read new column.
    *   **Phase 4:** Remove old column.
2.  **Protobuf/Avro:** Use forward/backward compatible schema evolution rules.

---

### Question 82: Thundering herd during cache warm-up. How do you solve it?

**Answer:**
1.  **Pre-warming:** Run a script to populate cache before switching traffic.
2.  **Request Coalescing:** (Same as Cache Stampede).
3.  **Jitter:** If cache expires, don't let all nodes refresh instantly. Add random delay.

---

### Question 83: A background reprocessing job overloads production. How do you control it?

**Answer:**
1.  **Rate Limiting:** Limit the job to N requests/sec.
2.  **Priority:** Run the job with lower priority (Quality of Service).
3.  **Isolation:** Point the job to a Read Replica, not the Primary DB.
4.  **Feedback Loop:** Pause job if high latency is detected in main app.

---

### Question 84: Partial failures cause cascading outages. How do you stop them?

**Answer:**
1.  **Circuit Breakers:** Stop calling the failing service.
2.  **Bulkheads:** Thread pool isolation. (Service A down shouldn't consume threads for Service B).
3.  **Timeouts:** Aggressive timeouts to free up resources.

---

### Question 85: A single shard becomes overloaded. How do you rebalance?

**Answer:**
1.  **Consistent Hashing:** Use Virtual Nodes to redistribute load more evenly.
2.  **Split Shard:** Split the hot shard into two (requires data movement).
3.  **Isolate Tenant:** If one tenant is the cause, move them to a dedicated shard/hardware.

---

### Question 86: You need to debug a bug that happens once a day at scale. How?

**Answer:**
1.  **Structured Logging:** Ensure correlation IDs exist.
2.  **Metrics:** Correlate the error timestamp with Cron jobs, GC events, or traffic spikes.
3.  **Trace Sampling:** Increase trace sampling rate around that time.
4.  **Canary:** Isolate traffic to one node to "catch" it with a debugger (or profiler) attached.

---

### Question 87: A rolling deployment causes inconsistent reads. Why?

**Answer:**
1.  **Version Mix:** User hits v2 server (writes new data format), then hits v1 server (reads old format -> fails).
2.  **Cache:** v1 cached an object structure that v2 doesn't understand (or vice versa).
3.  **Fix:** Ensure Backward Compatibility. Use Sticky Sessions during deployment (risky).

---

### Question 88: Clock skew causes ordering issues. How do you fix?

**Answer:**
1.  **Logical Clocks:** Use Lamport Timestamps or Vector Clocks to order events causally, not by wall clock.
2.  **NTP:** Ensure NTP is running (mitigation, not fix).
3.  **Google Spanner approach:** TrueTime API (atomic clocks + GPS) to bound the error.

---

### Question 89: A metrics system lies during outages. How do you validate data?

**Answer:**
1.  **Cross-Check:** Compare APM metrics (Latency) with Load Balancer logs (5xx count) and DB metrics.
2.  **External Probe:** Use Blackbox monitoring (Runscope, Pingdom) to test from outside.
3.  **Resolution:** Metrics might be aggregated (1 min avg) hiding 10s spikes. Look at raw logs or higher resolution.

---

### Question 90: Distributed tracing adds overhead. How do you balance visibility vs performance?

**Answer:**
1.  **Sampling:** Only trace 0.1% or 1% of requests.
2.  **Head-based Sampling:** Decide at the start of request (Random).
3.  **Tail-based Sampling:** Decide *after* request finishes (Keep trace only if it was an Error or Slow). (Expensive to buffer).

---

### Question 91: Leader election flaps frequently. What are the consequences?

**Answer:**
1.  **Downtime:** System is unavailable during election (Write pauses).
2.  **Stale Reads:** Clients might read from old leader.
3.  **Fix:** Increase heartbeat timeout. Check network stability between Zookeeper/Etcd nodes.

---

### Question 92: A service must degrade gracefully under overload. How?

**Answer:**
1.  **Load Shedding:** Drop low-priority requests (e.g., Background stats).
2.  **Feature Toggle:** Disable expensive features (e.g., "People who viewed this also viewed...").
3.  **Staleness:** serve stale data from cache instead of hitting DB.

---

### Question 93: Data loss is reported in an eventually consistent system. How do you investigate?

**Answer:**
1.  **Replication Delay:** Was it read from a lagging replica?
2.  **Conflict Resolution:** Did "Last Write Wins" overwrite the data silently?
3.  **Quorum:** Was the write acknowledged by `W=1` but that node crashed before replicating? Need `W=Quorum`.

---

### Question 94: A retry storm worsens an outage. How do you design retries?

**Answer:**
1.  **Exponential Backoff:** Wait 1s, 2s, 4s, 8s...
2.  **Jitter:** Wait `Random(1s, 1.5s)`. Prevents synchronized retries.
3.  **Circuit Breaker:** Stop retrying globally if the service is confirmed Down.
4.  **Retry Budget:** Limit retries to 10% of total traffic.

---

### Question 95: Global rate limiting is needed. How do you implement it?

**Answer:**
1.  **Central Store:** Redis Cluster with sliding window counters.
2.  **Performance:** Reading Redis for every request is slow. Use "baching" (Local agent aggregates 100 reqs -> updates Redis).
3.  **Token Bucket:** Standard algorithm for bursty traffic.

---

### Question 96: Backpressure is missing and causes crashes. How do you add it?

**Answer:**
1.  **Queue Limits:** Bounded queues (BlockingQueue). If full, reject new work.
2.  **TCP Backpressure:** Stop reading from the socket (client TCP window closes).
3.  **Reactive Streams:** Use libraries (RxJava, Project Reactor) that support backpressure natively.

---

### Question 97: Multi-tenant service sees noisy neighbor issues. How do you isolate?

**Answer:**
1.  **Quotas:** Enforce strict RPS/Storage limits per tenant.
2.  **Shuffle Sharding:** Isolate tenants to specific subsets of shards. If one shard dies, only a few tenants are affected.
3.  **dedicated Pool:** Move the VIP tenant to a dedicated hardware pool.

---

### Question 98: Blue-green deployment causes traffic imbalance. How do you fix?

**Answer:**
1.  **DNS Caching:** Clients/ISPs cache the old IP. (Lower TTL before switch).
2.  **Keepalive:** Long-lived TCP connections stay connected to the old Green fleet. (Force reconnect).
3.  **Load Balancer:** Use Weighted Target Groups (ALB) instead of DNS switch.

---

### Question 99: An SLO is violated intermittently. How do you root-cause it?

**Answer:**
1.  **Heatmap:** Look at latency usage heatmap.
2.  **Dependency:** Is it correlated with a backup job, a cron job, or a dependency deploying?
3.  **Noise:** Is it one specific heavy customer?

---

### Question 100: You’re on call and multiple alerts fire at once. How do you prioritize?

**Answer:**
1.  **Impact:** Triage based on Customer Impact (Data Loss > Downtime > High Latency > Internal Tool).
2.  **Root Cause:** Identify the "Source" alert vs "Symptom" alerts (DB Down causes 50 Service Alerts). Fix the DB.
3.  **Mitigate:** Focus on restoring service (Rollback/Restart/Scale) before detailed debugging.
4.  **Communicate:** Acknowledge incident to stakeholders.

---
