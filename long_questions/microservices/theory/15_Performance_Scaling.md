# 🟢 **236–250: Performance & Scaling**

### 236. How do you profile a microservice?
"Profiling is the deep, often intrusive analysis of a microservice's execution to identify performance bottlenecks. It goes beyond simple metrics (CPU/Memory usage) by inspecting the actual code execution paths.

In Java, I use tools like JProfiler or Java Flight Recorder (JFR) to take CPU snapshots (flame graphs) to see exactly which methods are burning CPU cycles, or take heap dumps to analyze object allocation rates and memory leaks. In Go, I use `pprof` to serve profiling data over HTTP.

Profiling is heavy. I never run it continuously in production. Instead, I use Continuous Profiling tools (like Datadog Profiler or Pyroscope) that periodically sample the application stack traces with very low overhead (~1%), allowing me to retroactively investigate why a service spiked at 3 AM."

#### Indepth
Memory leaks are fundamentally rarely caused structurally by the GC missing dead objects (as long as references are cleared). They are normally caused by long-lived data structures (like global Caches or `ThreadLocal` variables) aggressively retaining objects and casually lacking explicit TTL expiration policies natively.

---

### 237. What is Autoscaling and how do you configure it?
"Autoscaling dynamically adjusts the number of running instances of a microservice based on current load, ensuring performance during traffic spikes and saving money during low traffic.

In Kubernetes, this is handled by the **Horizontal Pod Autoscaler (HPA)**. I configure the HPA to monitor a specific metric—typically target CPU utilization (e.g., scale up if average CPU > 70%).

If a marketing push causes a burst of traffic, the CPU hits 80%. HPA automatically provisions 5 more pods. The traffic spreads out, CPU drops back to 50%, and the system stabilizes. Later that night, when traffic drops, HPA scales the pods back down to the minimum count to save cloud costs."

#### Indepth
Scaling based on CPU/Memory is often insufficient for asynchronous workers (e.g., a service consuming from Kafka). A worker processing complex images might use 100% CPU to process just 1 message per second. To autoscale workers effectively, I use **KEDA (Kubernetes Event-driven Autoscaling)** to scale based on the *custom metric* of Kafka Consumer Group Lag (e.g., "Scale up if there are more than 10,000 unprocessed messages in the queue").

---

### 238. Explain Amdahl's Law in the context of scaling.
"Amdahl's Law is a formula that dictates the theoretical maximum speedup you can achieve by parallelizing a given task. 

It states that the overall performance improvement is limited by the **strictly sequential (non-parallelizable) portion** of the workload.

If a batch job takes 100 minutes to run on 1 thread, and 5 minutes of that job is strictly sequential (e.g., writing the final summary report to a single locked database row), the maximum theoretical speedup is limited by that 5 minutes. Even if I throw 1,000 parallel microservice instances at the remaining 95 minutes of work to finish it instantly, the job will *still* take at least 5 minutes. You can't reach infinite scalability."

#### Indepth
This is why database locks and synchronous orchestrators become massive bottlenecks at scale. Designing for scale means ruthlessly hunting down sequential chokepoints and redesigning them to be parallel—e.g., using Eventual Consistency or Sharding instead of single-row locking.

---

### 239. What is DB Connection Pooling and why is it critical?
"Opening a brand new TCP connection to a database, authenticating, and initiating a session is an extremely expensive and slow network operation (often taking 20-50ms). 

If a microservice handles 1,000 requests per second and attempts to open a new DB connection for every single request, it will exhaust its own CPU and immediately crash the database server due to connection overhead.

A **Connection Pool** (like HikariCP in Java or `pgxpool` in Go) maintains a pool of pre-established, 'warm' database connections in memory. When a request needs to query the DB, it 'borrows' a connection from the pool, runs the query in 1ms, and instantly returns the connection to the pool for the next request to use. It fundamentally protects the database from connection storms."

#### Indepth
Pool sizing is counter-intuitive. A pool of 10-20 connections per microservice is often vastly superior to a pool of 1,000. PostgreSQL, for instance, thrives with a small number of highly active connections. Massive connection pools cause excessive context switching at the database kernel level, severely degrading actual throughput. Formula: `Connections = ((core_count * 2) + effective_spindle_count)`.

---

### 240. How do you implement Rate Limiting?
"Rate Limiting protects my APIs from brute-force attacks, DDoS, and aggressive clients by restricting the number of API calls a user or IP can make within a specific time window.

I implement rate limiting at the **API Gateway level**, preventing bad traffic from ever reaching the internal microservices. The standard algorithm is the **Token Bucket**.

Using a centralized Redis instance (so rate limits work across all Gateway nodes), each User ID gets a bucket with 100 tokens. Every API call consumes 1 token. Tokens refill at a rate of 10 per second. If the bucket hits 0, the Gateway rejects the request with an HTTP 429 'Too Many Requests' status, completely shielding the backend."

#### Indepth
To implement this atomically in Redis without race conditions, you must use Lua scripts. A single `EVAL` command checks the token count, calculates refill based on the current timestamp, decrements, and updates the TTL in one transaction.

---

### 241. What is the N+1 Query Problem?
"The N+1 Query Problem is a notorious performance killer, usually caused by naive ORM (Object-Relational Mapping) usage like Hibernate or GORM.

It happens when I query a database for a list of $N$ items (1 query), and then, while looping through those items to display them, the ORM lazy-loads a relationship for each item, firing $N$ additional secondary queries. 
To retrieve a list of 100 Orders and their associated Customer names, the ORM fires 1 query for Orders, plus 100 separate queries for Customers. Total: 101 queries. 

The fix is **Eager Loading** (e.g., `JOIN FETCH` in JPA, or `Preload()` in GORM) to fetch everything in 1 massive efficient SQL `JOIN`, dropping the query count from 101 to exactly 1."

#### Indepth
In microservices API Composition (or GraphQL), this exact same concept exists as "The N+1 API Request Problem". If I fetch 100 orders, and make 100 separate HTTP calls to the Customer Service to resolve their names, I bring down the network. The solution is the **Batch Loader / DataLoader** pattern, which collects all 100 required Customer IDs and makes exactly 1 bulk HTTP request `GET /customers?ids=1,2,3...100`.

---

### 242. How does Pagination impact API Performance?
"Return payload size heavily affects API throughput. You cannot return 1,000,000 rows of user data in one JSON payload; it will cause OutOfMemory errors on the server and crash the client parsing it.

APIs must use **Pagination**. The naive approach is **Offset-based pagination** (`OFFSET 5000 LIMIT 100`). This is disastrously slow at scale because the database engine still has to fetch and count the first 5,000 rows internally just to throw them away and return the next 100.

The highly scalable approach is **Cursor-based pagination** (or Keyset pagination). You pass a specific reference point (e.g., `WHERE id > 5000 LIMIT 100`). Because the database can use an Index to instantly jump exactly to ID 5000, fetching page 1,000 is just as blazing fast as fetching page 1."

#### Indepth
Cursor pagination makes jumping to a specific page number impossible (e.g., you can't click "Page 50" on a UI, you can only click "Next"). Modern infinite-scroll UIs (like Twitter feeds or TikTok) strictly rely on Cursor pagination to maintain sub-millisecond query performance regardless of how deeply you scroll.

---

### 243. What is Data Sharding?
"When a microservice's single underlying database grows too massive to fit on one physical disk or handle the write volume, we implement **Data Sharding**.

Sharding horizontally partitions the rows of a database across multiple independent, smaller database servers. 

For a massive User Service, I select a **Shard Key** (e.g., User ID). Using consistent hashing, User IDs 1-1,000,000 go to Shard A (DB Server 1), and 1,000,001-2,000,000 go to Shard B (DB Server 2). 

This allows infinite horizontal scaling of write throughput. However, any query that doesn't include the Shard Key (e.g., 'Find all users born in May') forces an expensive 'Scatter-Gather' query, hitting every shard simultaneously and merging the results in memory."

#### Indepth
Resharding (adding a new Shard C when A and B get full) is the most dangerous operational procedure in databases. It requires migrating terabytes of live data transparently while the application continues to write. Modern "NewSQL" databases like CockroachDB or TiDB handle this auto-sharding natively, abstracting the nightmare away from developers.
