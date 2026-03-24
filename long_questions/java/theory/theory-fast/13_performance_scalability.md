# Performance & Scalability Interview Questions (141-148)

## Performance Tuning

### 141. How do you identify performance bottlenecks?
"I start from the **user's perspective** and work backward.

1.  **Observability**: I check the P99 latency. Is it the API, the DB, or the external service?
2.  **Profiling**: If it's the application (CPU high), I use a profiler (like Async Profiler or JProfiler) to see where the threads are spending time. Is it a specific method?
3.  **Database**: If it's the DB, I check slow query logs.
4.  **Network**: I check for network latency or packet loss.

The key is measuring, not guessing. Fixing the wrong bottleneck is a waste of time."

**Spoken Format:**
"Finding performance bottlenecks is like being a detective for a slow system.

You start with the user's perspective - they're complaining something is slow, but you don't know what.

The investigation process is:

1. **Observability** is like checking security cameras - you look at metrics dashboards to see if the problem is widespread or isolated.

2. **Profiling** is like dusting for fingerprints - you use tools to see exactly where the application is spending its time.

3. **Database Analysis** is like checking traffic patterns - you look at slow query logs to see if certain queries are causing traffic jams.

4. **Network Analysis** is like checking road conditions - you look for latency or packet loss affecting performance.

The key insight: Don't guess based on assumptions. Measure first, identify the actual bottleneck, then fix the right problem. Fixing the wrong bottleneck is like treating a headache when the real issue is a broken leg!""

### 142. Difference between throughput and latency?
"**Latency** is *speed*: How fast does one request return? (e.g., 50ms). It matters to the individual user.

**Throughput** is *capacity*: How many requests can we handle per second? (e.g., 10,000 RPS). It matters to the system owner.

You can have low latency but low throughput (a fast server that crashes after 10 users). Or high throughput but high latency (a batch job that processes 1M records but takes an hour to finish)."

**Spoken Format:**
"Latency and throughput are like two different ways to measure a highway's performance.

**Latency** is like measuring how fast one car travels from point A to point B. If the trip takes 30 minutes, that's high latency - individual drivers care about this.

**Throughput** is like measuring how many cars can travel on the highway at the same time. If 1000 cars can travel per hour, that's high throughput - the highway authority cares about this.

The key difference:
- Low latency but low throughput: A few cars travel very fast, but highway can handle many more
- High latency but high throughput: Many cars travel, but each trip takes a long time

For user experience, latency matters more. For system capacity, throughput matters more. You need both for a truly performant system!""

### 143. How do you handle high traffic spikes?
"Scalability handles gradual growth; **Elasticity** handles spikes.

1.  **Auto-scaling**: Configure AWS/K8s to add more instances when CPU > 70%.
2.  **Caching**: Aggressively cache read data to protect the DB.
3.  **Rate Limiting**: Protect the system by rejecting excess traffic (429s).
4. **Asynchronous Processing**: Put heavy write operations (like 'Process Video') into a Queue (Kafka) and process them later. This keeps the API responsive even if the backend is swamped."

**Spoken Format:**
"Handling traffic spikes is like preparing your restaurant for a sudden rush of customers.

**Auto-scaling** is like having extra staff on call - when the restaurant gets busy, more waiters and chefs automatically appear to handle the crowd.

**Caching** is like having popular dishes pre-prepared - instead of cooking everything from scratch during rush hour, you serve popular items quickly from cache.

**Rate Limiting** is like having a bouncer who limits how many people can enter at once - prevents the kitchen from being overwhelmed.

**Asynchronous Processing** is like taking orders now but cooking later - you accept the order immediately so customers don't wait, but the actual cooking happens in the background.

The combination ensures that even during sudden rushes, your system remains responsive and doesn't crash under pressure!"

### 144. What is backpressure?
"Backpressure is a feedback mechanism where a consumer tells a producer to slow down.

Imagine a fast API pushing data to a slow database. If the API keeps pushing, the system will run out of memory buffering the data.

Reactive Streams (like Project Reactor in Spring WebFlux) handle this automatically. The subscriber requests only N items. The publisher sends N and waits for next request. It prevents `OutOfMemoryError` by matching the flow rate."

**Spoken Format:**
"Backpressure is like having a smart traffic control system for data flow.

Imagine a fast water pump (producer) trying to fill a small bucket (consumer). If the pump keeps pouring, the bucket overflows and water spills everywhere.

**Reactive Streams** solve this by:
- The consumer tells the pump 'I can only handle N drops per second'
- The pump adjusts its flow rate to match what the consumer can handle
- No more overflowing, no more wasted water

This is like having a smart shower that automatically adjusts water pressure based on how much you can drink.

In technical terms, it prevents `OutOfMemoryError` by ensuring that producers don't overwhelm consumers. It's like having a self-regulating system that prevents crashes!"

### 145. When would you use async processing?
"I use async whenever the user **doesn't need the result immediately**.

Example: User uploads a CSV for bulk import.
Synchronous: User waits 5 minutes with a spinning loader. Connection times out. Bad.
Async: API accepts the file, returns '202 Accepted', and queues a background job. We email the user when it's done.

I also use `CompletableFuture` to run parallel tasks—like fetching User Profile and User Orders at the same time—to reduce total latency."

**Spoken Format:**
"Async processing is like having multiple checkout counters in a supermarket.

**Synchronous processing** is like having only one checkout counter. If you have 10 people in line, each person waits for everyone ahead of them to finish.

**Asynchronous processing** is like having 10 checkout counters. People can start checking out immediately, and the actual processing happens in parallel.

The benefits are:
- **Better user experience** - No one waits unnecessarily
- **Higher throughput** - Multiple tasks run simultaneously
- **Resource efficiency** - System resources are used optimally

**CompletableFuture** is like having a smart coordinator that manages multiple checkout counters and tells you when everything is done.

The key is to use async when users don't need immediate results but want the system to be responsive overall!"

### 146. How do you tune JVM for performance?
"Honestly, modern JVMs (Java 17+) are very good out of the box. Premature tuning is bad.

But when I do tune, I focus on:
1.  **Heap Size**: Setting `-Xms` and `-Xmx` equal to prevent resizing jitter.
2.  **GC Algorithm**: Switching to G1GC (default) or ZGC if latency is critical.
3.  **String Deduplication**: Enabling `-XX:+UseStringDeduplication` if heap dumps show many duplicate strings.

I monitor GC logs: if the app spends >5% of time in GC, then I tune."

**Spoken Format:**
"JVM tuning is like optimizing a car's engine for better performance.

Modern JVMs are like modern cars - they're already well-tuned out of the factory. Premature tuning is like trying to modify a brand new car's engine before even driving it.

When I do tune, I focus on:

**Heap Size** is like deciding how big your fuel tank should be. Too small and you run out of fuel frequently. Too big and the car is heavy.

**GC Algorithm** is like choosing between automatic and manual transmission. Different algorithms work better for different driving patterns.

**String Deduplication** is like having a smart system that reuses common phrases instead of storing duplicates.

The key insight: Monitor first, then tune. If you see the engine spending too much time on fuel management (GC), then you know what to adjust. Don't tune blindly - tune based on actual performance data!"

### 147. How do you design for read-heavy systems?
"Read-heavy means >90% reads (like Twitter or a News site).

1.  **Caching**: Redis/Memcached is mandatory. Cache aggressively.
2.  **Read Replicas**: One Primary DB (for writes) and 5 Read Replicas. All read traffic goes to replicas.
3.  **CDN**: Content Delivery Network for all static assets.
4. **Materialized Views**: Pre-calculate complex reports so that read query is essentially `SELECT * FROM report`."

**Spoken Format:**
"Designing read-heavy systems is like building a library optimized for researchers, not for book-lending.

**Caching** is like having a popular books section at the front - researchers can grab frequently accessed books instantly without going to the stacks.

**Read Replicas** is like having multiple copies of popular books - multiple researchers can read the same book simultaneously without waiting.

**CDN** is like having regional library branches - researchers get books from the nearest branch instead of traveling to the main library.

**Materialized Views** are like having pre-written research summaries - instead of searching through thousands of pages for information, researchers just read the summary.

The strategy is to minimize database reads by serving data from the fastest possible location. It's like having a smart librarian who knows exactly where each piece of information is stored!"

### 148. What are hot keys in Redis?
"A Hot Key is a specific key (like `global_config` or `justin_bieber_profile`) that gets requested thousands of times per second.

This overloads the single Redis shard holding that key, causing a bottleneck while other shards are idle.

To fix it, we use **Local Caching** (in-memory on the app server) for that specific key with a short TTL. This offloads traffic from Redis to the application instances."

**Spoken Format:**
"Hot keys in Redis are like having one book that everyone wants to read at the same time.

Imagine a library where everyone wants to read the same popular book simultaneously. If the book is stored on one shelf (Redis shard), everyone has to wait in line to read it.

**Local Caching** solves this by:
- Each library branch (application server) keeps its own copy of the popular book
- Readers get the book instantly from their local branch
- No one has to wait for the central book

This is like having multiple copies of the same book distributed across different library branches.

The result: Better performance for readers, and reduced load on the central book storage system. It's like having a smart distribution system that prevents bottlenecks!"
