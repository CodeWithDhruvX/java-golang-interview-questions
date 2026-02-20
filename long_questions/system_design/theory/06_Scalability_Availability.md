# 🔴 Scalability & Availability — Questions 51–60

> **Level:** 🔴 Senior (5+ yrs)
> **Asked at:** FAANG, Flipkart, PhonePe, Razorpay, Swiggy — for systems serving millions of users

---

### 51. How do you design a system that handles millions of users?
"Designing for millions of users is about building systems that scale horizontally, eliminate single points of failure, and use every caching and async trick in the book.

My approach: start with requirements — DAU (Daily Active Users), peak QPS, data volume, latency SLA. Then design from the outside in: globally distributed CDN for static content, geo-aware DNS routing, load-balanced stateless app servers, a caching layer (Redis) to absorb most reads, and finally the database tier which is the hardest to scale.

At 1M users, a single well-tuned Postgres with a read replica is likely enough. At 10M users, I'm looking at multiple read replicas, Redis caching. At 100M users, I'm thinking about sharding, async processing for heavy writes, and probably separate services for hot paths."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Google, Flipkart, Swiggy — classic "scale" interview question

#### Indepth
Scale milestones and architecture evolution:

| Scale | Architecture |
|---|---|
| ~ 1K users | Single server (VPS) + managed DB |
| ~ 10K users | App server + separate DB server |
| ~ 100K users | Load-balanced stateless app servers (2-3) + read replica |
| ~ 1M users | + Redis cache, CDN, async job queues |
| ~ 10M users | + DB sharding or read-heavy NoSQL, separate microservices for hot paths |
| ~ 100M users | Multi-region, data centers in multiple continents, Kafka event streaming |
| ~ 1B users | Netflix/Facebook/Google scale — custom hardware, global Anycast, massive data pipeline |

Back-of-envelope math example:
- 10M users, 10% online at peak → 1M concurrent users
- Each user makes 5 requests/minute → 5M req/min = ~83K QPS
- Each server handles 1K QPS → Need 83 app servers
- DB read ratio 80% → Absorb 80% with Redis → DB sees 16K QPS → 2-3 Postgres primaries

---

### 52. How to scale a system with read-heavy workload?
"Read-heavy systems are the most common scenario — social feeds, product catalogs, news sites. The solution hierarchy is: cache aggressively, replicate reads, denormalize for fast retrieval.

Step one: **Cache at every layer**. For a social feed, cache the computed feed in Redis. Cache individual post data. Cache user profile data. A 95% cache hit rate means only 5% of reads touch the DB.

Step two: **Read replicas**. All writes go to the primary DB; all reads distribute across N read replicas. PostgreSQL streaming replication adds replicas in minutes. A read-heavy setup might have 1 primary + 10 replicas.

Step three: **CDN for static content**. Images, videos, static HTML — serve from the edge. The DB never sees these requests."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Hotstar, Netflix, Swiggy, Flipkart product catalog teams

#### Indepth
Read scaling toolkit:
1. **In-process cache (L1):** Small LRU cache in app memory. Hit rate: very high for hot data. Evicted on service restart. No network hop.
2. **Distributed cache (Redis/Memcached):** Shared across all app instances. Persistent across restarts. One network hop.
3. **Read replicas:** Async-replicated DB copies. Good for query-heavy reads that can't be cached (complex filtering). Replication lag: typically <100ms.
4. **CQRS Read Model:** Materialized views or separate read-optimized DB (Elasticsearch for search, ClickHouse for analytics) updated asynchronously. Best hit rate for complex queries.
5. **Eventual consistency at the read tier:** Serve stale data for non-critical reads. Use `Cache-Control: stale-while-revalidate` at CDN level.

**Read-your-writes consistency:** After a user posts a photo, they must see their own post immediately (even though the read DB might lag). Solution: route the posting user's reads to the primary for a short window (e.g., 10 seconds), then route to replicas.

---

### 53. How to scale a system with write-heavy workload?
"Write-heavy systems are harder than read-heavy — you can't cache away writes. Examples: IoT sensor ingestion, real-time analytics events, chat message delivery, financial transactions.

The primary strategies: **write batching** (buffer writes in memory, flush periodically — trades durability for throughput), **async writes** (write to a fast message queue like Kafka, let a consumer write to DB asynchronously at its own pace), and **LSM-tree storage engines** (Cassandra, RocksDB optimized for append-heavy writes).

For pure transactional write-heavy systems (payment platform), I use **DB sharding** — split writes across multiple DB primaries. For time-series / event data, I use **Kafka + stream processing** — events are written to Kafka at millions/second, consumers batch-write to Cassandra or ClickHouse."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, PhonePe, Twitter/X (tweet writes), Hotstar (event tracking), IoT companies

#### Indepth
Write optimization techniques:

- **Write Buffering:** Batch multiple writes into one DB transaction. Write 1000 events in one `INSERT` instead of 1000 separate `INSERTs`. Throughput improvement: 100-1000x (fewer round trips, fewer transaction overhead).
- **Kafka as Write Buffer:** Producers write to Kafka at millions of events/second. Kafka consumers write to DB in batches. DB handles batch writes at its own pace. Kafka provides durability + replay capability.
- **Cassandra LSM Trees:** Writes go to MemTable (in-memory) then sequentially flushed to SSTables on disk. Sequential disk writes are 10-100x faster than random B-tree writes. Cassandra handles 1M+ writes/second on modest hardware.
- **Sharding by write key:** Shard payment records by `merchant_id` — each merchant's transactions go to a specific shard. No cross-shard writes for single-merchant operations.
- **CQRS Command Model:** Write to a normalized, write-optimized command DB. Separate read projections updated asynchronously. Each model is tuned independently.

**Write amplification trap:** SSDs have a finite write endurance. Databases with B-trees write the same data multiple times (write amplification factor ~10x for MySQL InnoDB). LSM-tree DBs have lower write amplification. Important consideration for write-heavy SSD deployments.

---

### 54. What is replication and when to use it?
"Database replication is the process of copying data from one DB node (primary/master) to one or more other nodes (replicas/secondaries). The primary accepts writes; replicas continuously apply those writes to stay in sync.

I use replication for two purposes: **read scaling** (distribute read queries across replicas) and **high availability** (if the primary fails, promote a replica to be the new primary with minimal downtime).

In PostgreSQL, streaming replication is built-in and near-instant — replicas are typically 10-100ms behind primary. For global deployments, replication across continents introduces more lag (200-500ms), which is the fundamental physics-imposed consistency trade-off."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Any company with a non-trivial DB footprint — Amazon, Flipkart, Razorpay, Google

#### Indepth
Replication types:
- **Synchronous Replication:** Primary waits for at least one replica to confirm the write before acknowledging to the client. Zero data loss but higher write latency. Used in financial systems where zero data loss is mandatory.
- **Asynchronous Replication:** Primary acknowledges write immediately; replica applies it later. Lower write latency but potential data loss of committed transactions if primary crashes before replica applies them. Used for read replicas and DR replicas.
- **Semi-synchronous (MySQL):** Primary waits for one replica to receive (not apply) the write. Balances latency and data safety.
- **Single-leader (Master-Slave):** One primary. Simple, most common (PostgreSQL, MySQL).
- **Multi-leader (Multi-master):** Multiple primaries accept writes. Complex conflict resolution needed. Avoids single-region dependency. Used in CockroachDB, DynamoDB Global Tables.
- **Leaderless (Dynamo-style):** Any node can accept writes. Uses quorum reads/writes (e.g., W=2, R=2, N=3). Cassandra, DynamoDB.

---

### 55. How to make a system fault-tolerant?
"Fault tolerance is designing a system so that **individual component failures don't cause system-wide failures**. The core principle: assume everything will fail — hardware, network, software — and design accordingly.

My toolkit: **redundancy** (no single points of failure — every critical component has a backup), **circuit breakers** (stop cascading failures from a failing dependency), **timeouts and retries with exponential backoff** (don't wait forever for a response, and don't hammer a struggling service), **graceful degradation** (return a useful partial response even when some components are down), and **bulkheads** (isolate failures so a problem in one service pool doesn't exhaust resources for all).

Netflix famously runs **Chaos Monkey** in production — a service that randomly terminates EC2 instances during business hours to ensure the system is genuinely fault-tolerant, not just theoretically."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix (chaos engineering originators), Amazon (AWS reliability), Google SRE, PhonePe (payment uptime is critical)

#### Indepth
The four key resilience patterns:
1. **Timeout:** Every network call has a deadline (e.g., 200ms). If exceeded, fail fast. Don't let slow dependencies tie up threads. `context.WithTimeout(ctx, 200*time.Millisecond)` in Go.
2. **Retry with exponential backoff + jitter:** On failure, retry after 1s → 2s → 4s → 8s with random jitter. Jitter prevents retry storms (all clients retrying simultaneously = thundering herd 2.0).
3. **Circuit Breaker:** As explained in Q47. After N failures, stop calling the failed service. Prevents cascading failures.
4. **Bulkhead:** Isolate resource pools per service. Payment service gets its own thread pool (100 threads) and order service gets its own (100 threads). If Payment service hangs all its threads, order service threads are unaffected.

Additional tools:
- **Fallback Cache:** When DB is down, serve stale cached data rather than returning errors
- **Feature Flags:** Quickly disable a non-critical feature under load
- **Multi-AZ / Multi-Region:** No correlated failures — AZ-level failures affect one DC, not all

---

### 56. What is failover?
"Failover is the **automatic switching from a failed primary component to a backup** with minimal service interruption.

Database failover: primary DB goes down → replica promotion to primary (handled by Patroni, AWS RDS Multi-AZ, PgBouncer). Application servers: LB's health check detects dead server → stops routing traffic to it → routes all traffic to healthy servers. Region failover: Route 53 health check detects unhealthy region → updates DNS to point to DR region.

The key metrics are **RTO (Recovery Time Objective)** — how fast must failover complete? And **RPO (Recovery Point Objective)** — how much data loss is acceptable? AWS RDS Multi-AZ achieves RTO of ~60-120 seconds and RPO of near-zero (synchronous replication)."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (RDS/Aurora HA design), PhonePe, Razorpay (payment uptime), Google (SRE concepts)

#### Indepth
Failover strategies by layer:
- **Database Failover with Patroni (PostgreSQL HA):** Patroni is a template for managing PostgreSQL HA. It uses etcd/Consul/ZooKeeper for distributed consensus. When primary dies, Patroni holds an election, promotes the most up-to-date replica, and updates the connection URL (via HAProxy or DNS). RTO: ~30 seconds.
- **AWS RDS Multi-AZ:** AWS manages synchronous standby in a different AZ. On primary failure, AWS automatically updates the DNS CNAME to point to standby. Application reconnects transparently (if using connection pooling). RTO: 60-120 seconds.
- **AWS Aurora Multi-AZ:** Aurora separates storage (shared across AZs, 6 copies) from compute. Failover within storage cluster is instant; compute failover to read replica is ~30s.

**Split-brain prevention:** During network partition, both primary and backup might think they're the primary. This is catastrophic for data consistency. Solutions: STONITH (Shoot The Other Node In The Head — physically fence the failed node), quorum-based consensus (etcd ensures only one node becomes primary).

---

### 57. What is high availability?
"High Availability (HA) means a system is **operational for a very high percentage of time** — typically expressed as 'nines': 99.9% (three nines) = 8.7 hours downtime/year, 99.99% = 52 minutes/year, 99.999% = 5 minutes/year.

HA is achieved by eliminating single points of failure, implementing fast failover, and designing for fault isolation. It's not just about not going down — it's about recovering *quickly* when things inevitably do go wrong.

Designing for 99.99% uptime means: every component has N+1 redundancy, failover is automated (not manual paging at 3am), deployments are zero-downtime, and you have fire drills (chaos engineering) to validate that your HA mechanisms actually work."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** SRE roles at Google, Amazon, LinkedIn, Razorpay, PhonePe (payment SLA commitments to clients)

#### Indepth
SLA/SLO/SLI framework (Google SRE's model):
- **SLI (Service Level Indicator):** The actual measured metric. Examples: request success rate, P99 latency, error rate.
- **SLO (Service Level Objective):** Internal target. "99.9% of requests succeed in <300ms."
- **SLA (Service Level Agreement):** External contract with customers with financial penalties. "We guarantee 99.9% uptime per month."
- **Error Budget:** The remaining tolerance for downtime. At 99.9% SLO, you have 8.7 hours/year of error budget. If you've used most of it, no more risk releases until next period.

HA design checklist for a service:
- [ ] Load balanced (≥2 instances in ≥2 AZs)
- [ ] Auto-healing (unhealthy instances automatically replaced)
- [ ] DB in Multi-AZ or with replica + automated failover
- [ ] Cache is HA (Redis Sentinel or Redis Cluster)
- [ ] Message queue is replicated (Kafka with replication factor ≥3)
- [ ] CDN has automatic origin failover
- [ ] DNS TTL low enough for fast failover (<60 seconds)

---

### 58. Difference between active-passive and active-active systems.
"Active-passive and active-active are two HA architectures.

**Active-Passive:** One node handles all traffic (active), the other is idle on standby (passive). On failure of the active, the passive takes over. Simple to implement and reason about. The downside: you're paying for a server that does nothing 99.9% of the time.

**Active-Active:** Both nodes handle traffic simultaneously. More efficient resource utilization. More complex because both nodes process writes simultaneously, requiring data synchronization and conflict resolution. If one node fails, the other takes 100% of the load (must be provisioned for 2x capacity).

For load balancers: Active-Active is standard. For databases: Active-Passive is safer and more common (to avoid write conflicts). Multi-master DB replication is possible but complex."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Google, Flipkart — infrastructure design discussions

#### Indepth
| Aspect | Active-Passive | Active-Active |
|---|---|---|
| Traffic | All to active | Distributed to both |
| Resource utilization | 50% wasted (idle passive) | 100% utilized |
| Failover complexity | Simple — passive promotes | Complex — must handle in-flight requests + state sync |
| Data conflicts | Not possible (only one writer) | Possible (both write → conflict resolution) |
| Failover time | Seconds (IP takeover / DNS update) | Zero (traffic re-weighted instantly) |
| Common for | DB primaries, master services | LBs, stateless app servers |

**Active-Active for stateless services:** App servers are stateless (sessions in Redis). Running 3 instances in active-active means any instance can handle any request. LB distributes evenly. If one dies, LB redistributes. True active-active with no conflicts. This is the ideal design for modern microservices.

---

### 59. What is graceful degradation?
"Graceful degradation means a system continues to **provide partial functionality even when some components fail**, rather than failing completely.

Amazon's product page is a classic example: if the recommendations engine is down, the page still loads with product details, images, and price. You just don't see 'Customers also bought' section. If the reviews service is down, the product still appears, minus the reviews. The core buying experience is preserved even under partial failure.

This requires careful triage of features: which are 'must-have' (payment, product display) and which are 'nice-to-have' (recommendations, cross-sells, analytics). Must-have features must never depend on nice-to-have services."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (pioneered this), Netflix, Hotstar, Flipkart, Swiggy

#### Indepth
Implementation techniques:
1. **Bulkhead isolation:** Nice-to-have features (recommendations) run in a separate thread pool. If they hang, they don't consume threads needed for core checkout.
2. **Default fallbacks:** `getRecommendations()` returns `[]` (empty list) on timeout instead of propagating error. UI handles empty list gracefully ("No suggestions available").
3. **Feature flags:** Toggle off expensive/failing features at runtime without deploying code.
4. **Circuit breakers:** Automatically stop calling failing services. Return fallback immediately.
5. **Client-side fallback:** If API call fails, render cached data from before. Progressive web apps store critical assets in Service Worker cache.

**Netflix's Chaos Monkey** validates graceful degradation: it kills random services and checks if the main streaming experience still works (it should, degraded but functional). This is the only way to truly verify degradation works in production — testing in staging doesn't replicate production conditions.

---

### 60. What is a throttling mechanism?
"Throttling (Rate Limiting) is controlling the rate at which requests are processed or resources are consumed — protecting the system from being overwhelmed and ensuring fair usage.

I implement rate limiting at the API gateway layer so every service benefits without individual implementation. The most common algorithm: **Token Bucket** — each client gets a bucket of N tokens that refill at rate R per second. Each request costs one token. If the bucket is empty, the request is rejected (HTTP 429).

For a payment API: allow 100 requests/minute per API key. This prevents:
1. Abusive clients from hammering the API
2. A bug in a client that creates an infinite retry loop
3. DDoS protection (not complete but reduces impact)
4. Fair usage across all customers"

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, Stripe (API-first companies), AWS (API Gateway rate limiting), Swiggy/Zomato (restaurant order throttle)

#### Indepth
Rate limiting algorithms:

1. **Token Bucket:** Bucket holds N tokens. Tokens fill at rate R/second. Request consumes 1 token. Empty bucket = reject. Allows short bursts (bucket can be full) up to N requests. Stripe and Razorpay use this.

2. **Leaky Bucket:** Requests enter a fixed-size queue. Requests leave at a constant rate. Burst requests go into queue; if queue is full, reject. Smooths traffic — no bursts allowed. Used in network traffic shaping.

3. **Fixed Window Counter:** Count requests in a fixed time window (1 minute). If counter > limit, reject. Simple but has edge case: 100 requests in last second of minute 1 + 100 in first second of minute 2 = 200 in 2 seconds (double the limit).

4. **Sliding Window Log:** Log every request timestamp. Count requests in last N seconds from current timestamp. Perfect accuracy. Memory-intensive (store all timestamps).

5. **Sliding Window Counter:** Combination — approximate sliding window using two fixed-window counters + interpolation. Best balance of accuracy and efficiency.

**Distributed rate limiting:** Multiple app servers share a Redis counter. `INCR key; EXPIRE key 60` — but this has race conditions. Use Redis `EVAL` Lua scripts for atomic counter + check operation. Redis pipelines reduce network round trips.
