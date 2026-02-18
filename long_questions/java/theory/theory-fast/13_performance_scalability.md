# Performance & Scalability Interview Questions (141-148)

## Performance Tuning

### 141. How do you identify performance bottlenecks?
"I start from the **user's perspective** and work backward.

1.  **Observability**: I check the P99 latency. Is it the API, the DB, or the external service?
2.  **Profiling**: If it's the application (CPU high), I use a profiler (like Async Profiler or JProfiler) to see where the threads are spending time. Is it a specific method?
3.  **Database**: If it's the DB, I check slow query logs.
4.  **Network**: I check for network latency or packet loss.

The key is measuring, not guessing. Fixing the wrong bottleneck is a waste of time."

### 142. Difference between throughput and latency?
"**Latency** is *speed*: How fast does one request return? (e.g., 50ms). It matters to the individual user.

**Throughput** is *capacity*: How many requests can we handle per second? (e.g., 10,000 RPS). It matters to the system owner.

You can have low latency but low throughput (a fast server that crashes after 10 users). Or high throughput but high latency (a batch job that processes 1M records but takes an hour to finish)."

### 143. How do you handle high traffic spikes?
"Scalability handles gradual growth; **Elasticity** handles spikes.

1.  **Auto-scaling**: Configure AWS/K8s to add more instances when CPU > 70%.
2.  **Caching**: Aggressively cache read data to protect the DB.
3.  **Rate Limiting**: Protect the system by rejecting excess traffic (429s).
4.  **Asynchronous Processing**: Put heavy write operations (like 'Process Video') into a Queue (Kafka) and process them later. This keeps the API responsive even if the backend is swamped."

### 144. What is backpressure?
"Backpressure is a feedback mechanism where a consumer tells a producer to slow down.

Imagine a fast API pushing data to a slow database. If the API keeps pushing, the system will run out of memory buffering the data.

Reactive Streams (like Project Reactor in Spring WebFlux) handle this automatically. The subscriber requests only N items. The publisher sends N and waits for the next request. It prevents `OutOfMemoryError` by matching the flow rate."

### 145. When would you use async processing?
"I use async whenever the user **doesn't need the result immediately**.

Example: User uploads a CSV for bulk import.
Synchronous: User waits 5 minutes with a spinning loader. Connection times out. Bad.
Async: API accepts the file, returns '202 Accepted', and queues a background job. We email the user when it's done.

I also use `CompletableFuture` to run parallel tasks—like fetching User Profile and User Orders at the same time—to reduce total latency."

### 146. How do you tune JVM for performance?
"Honestly, modern JVMs (Java 17+) are very good out of the box. Premature tuning is bad.

But when I do tune, I focus on:
1.  **Heap Size**: Setting `-Xms` and `-Xmx` equal to prevent resizing jitter.
2.  **GC Algorithm**: Switching to G1GC (default) or ZGC if latency is critical.
3.  **String Deduplication**: Enabling `-XX:+UseStringDeduplication` if heap dumps show many duplicate strings.

I monitor GC logs: if the app spends >5% of time in GC, then I tune."

### 147. How do you design for read-heavy systems?
"Read-heavy means >90% reads (like Twitter or a News site).

1.  **Caching**: Redis/Memcached is mandatory. Cache aggressively.
2.  **Read Replicas**: One Primary DB (for writes) and 5 Read Replicas. All read traffic goes to replicas.
3.  **CDN**: Content Delivery Network for all static assets.
4.  **Materialized Views**: Pre-calculate complex reports so the read query is essentially `SELECT * FROM report`."

### 148. What are hot keys in Redis?
"A Hot Key is a specific key (like `global_config` or `justin_bieber_profile`) that gets requested thousands of times per second.

This overloads the single Redis shard holding that key, causing a bottleneck while other shards are idle.

To fix it, we use **Local Caching** (in-memory on the app server) for that specific key with a short TTL. This offloads the traffic from Redis to the application instances."
