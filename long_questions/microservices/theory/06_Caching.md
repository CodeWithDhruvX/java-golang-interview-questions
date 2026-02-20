# 🟢 **91–100: Caching**

### 91. What is caching?
"Caching is the practice of storing frequently accessed data in a temporary, extremely fast storage layer (usually RAM) so that future requests for that data can be served dramatically faster than querying the primary database.

If a 'Top 10 Products' API query takes 500ms to calculate via SQL joins, I run the query once and store the JSON result in a cache. For the next 5 minutes, every user asking for that data gets it from the cache in 5ms. 

I use caching aggressively to protect backend databases from massive read traffic and provide sub-millisecond latencies to end users."

#### Indepth
Caches operate at multiple levels: Browser Caching (HTTP Cache-Control headers), Edge Caching (CDNs capturing static assets), API Gateway Caching (short-circuiting requests before they hit microservices), and Application-Level Caching (Redis/Memcached). Effective caching can reduce primary database loads by over 90%.

---

### 92. Distributed caching?
"Distributed caching is a cache shared by multiple instances of an application, typically hosted on a separate cluster of servers (like Redis or Memcached).

If I ran an 'in-memory' cache (like Java's `ConcurrentHashMap`), and I have 5 instances of my Product microservice, they would all have to query the database independently to build their own local caches. Worse, if someone updates a product price, one instance might update its cache while the other 4 still serve the old price.

With a distributed cache like Redis, all 5 microservice instances point to one central, blazing-fast Redis cluster. If one instance updates the price in Redis, all other 4 instances instantly see the new price on their next read."

#### Indepth
Distributed caches themselves must be clustered to avoid becoming a single point of failure. Redis handles this via Master-Slave replication and Redis Sentinel for automatic failover. This enables the cache to survive hardware failures seamlessly while servicing millions of reads per second.

---

### 93. What is cache aside pattern?
"Cache Aside (or 'Lazy Loading') is the most common caching pattern I use. The application code is responsible for managing the cache and the database directly.

When my service receives a request for User #5, it first checks the cache. 
- If User #5 is found (a **Cache Hit**), it returns it immediately. 
- If not found (a **Cache Miss**), the service queries the database, retrieves User #5, explicitly writes it into the cache, and then returns it.

I like this because it's simple to implement and ensures that only data that is actually requested ever takes up valuable cache memory."

#### Indepth
In this pattern, the cache acts strictly as a "sidekick" (aside). The underlying database is completely unaware of the cache's existence. The massive downside is that a cache miss incurs a huge latency penalty for that specific unlucky user request, as they have to wait for three network hops (check cache -> check DB -> write cache -> return).

---

### 94. What is write-through cache?
"In a Write-Through cache, the application treats the Cache as the primary data store. 

When my application wants to save a new Order, it writes that Order object directly into the Cache. The Cache software itself then synchronously writes that data to the backend database before acknowledging the write to the application.

I rarely use this because every 'Write' payload suffers a double-latency penalty (writing to RAM + writing to Disk). However, it guarantees that the cache and the database are always 100% perfectly synchronized, which is fantastic for critical read-heavy data."

#### Indepth
This pattern is often implemented using features native to highly advanced caching platforms (e.g., Hazelcast or Redis Enterprise) wrapping a slower relational DB. Because data is written to the cache first, subsequent reads are incredibly fast and never result in a Cache Miss penalty for newly created data.

---

### 95. What is write-behind cache?
"A Write-Behind (or Write-Back) cache is an asynchronous variation of the Write-Through pattern.

When my application saves a new Order, it writes it exclusively to the Cache, and the Cache immediately acknowledges success. It doesn't write to the database yet. Behind the scenes, the Cache software batches up hundreds of these changes periodically (say, every 5 seconds) and asynchronously Flushes them down into the primary database.

I use this for ultra-high-throughput write scenarios (like a gaming leaderboard or massive telemetry ingestion). The database never gets strained by individual row inserts because everything is batched beautifully."

#### Indepth
This pattern enables staggering write performance because to the application, all writes are occurring purely at RAM speeds. The terrifying trade-off is data loss: if the Cache server suddenly loses power before it flushes its asynchronous batch to the persistent database, those 5 seconds of orders are permanently obliterated.

---

### 96. What is cache invalidation?
"Cache invalidation is the process of deliberately deleting or updating stale data in the cache to ensure clients aren't reading obsolete information.

It notoriously holds the title of 'one of the two hardest problems in computer science'. If my inventory system receives new stock but the API still serves the cached 'Out of Stock' response for the next hour, business is lost.

When my microservice executes an `UPDATE` or `DELETE` on a database row, I immediately write application logic to execute a `DELETE` command against that specific key in Redis, forcing the next read request to fetch the fresh data from the database."

#### Indepth
Manual invalidation logic is deeply prone to race conditions (e.g., if the DB transaction rolls back after the cache was already deleted). For mission-critical syncing, CDC (Change Data Capture) tools like Debezium are highly preferred. They read the raw MySQL binary logs and instantly publish cache-invalidation events via Kafka with zero application-code interference.

---

### 97. What is TTL?
"TTL (Time To Live) is the simplest, most automated form of cache expiration. 

When I save a key like `User:5` into Redis, I attach an absolute expiration time to it (e.g., 60 minutes). After exactly 60 minutes, Redis automatically deletes the key. 

I aggressively add a TTL to almost *every* piece of cached data. It acts as the ultimate fail-safe 'safety net'. Even if my manual cache invalidation code fails due to a bug, the stale data will eventually destroy itself when the TTL expires."

#### Indepth
Strategic TTL configurations depend heavily on data volatility. A stock price ticker might have a TTL of 5 seconds. A blog article's content might have a TTL of 24 hours. A user's JWT blocklist might have a TTL matching precisely the exact remaining lifetime of the token itself.

---

### 98. How to prevent cache stampede?
"A Cache Stampede (or Dog-Piling) happens under heavy load when an incredibly popular cached item expires (its TTL hits 0).

If 10,000 users are simultaneously viewing a viral video, and the video's metadata cache expires, all 10,000 requests experience a Cache Miss simultaneously. All 10,000 requests instantly query the backend database for the exact same metadata, instantly crashing the database.

I prevent this using standard **Mutex Locks** (or Distributed Locks). When the cache misses, the first thread attempts to grab a lock in Redis. It succeeds, queries the DB, and repopulates the cache. The other 9,999 threads fail to grab the lock and are instructed to sleep for 50ms and try reading the cache again."

#### Indepth
Another modern strategy is **Probabilistic Early Expiration**. Using an algorithm (like XFetch), the application randomly decides to re-compute and update the cache *before* the TTL actually expires. If the TTL is 60m, at 59m, 1% of incoming requests act as if it expired, silently refreshing the background cache while the other 99% continue reading the existing item undisturbed.

---

### 99. How to use Redis in microservices?
"Redis is an incredibly versatile, ultra-fast in-memory data store. I use it for far more than simple caching.

I use Redis for **Distributed Session Storage**. If my User authenticates on Pod 1, their session is saved in Redis. If their next request routes to Pod 2, Pod 2 still knows they are authenticated.

I use it for **Rate Limiting** API Gateways using fast atomic operations (`INCR`). I use it for **Distributed Locks** (via Redlock algorithm) to prevent duplicate cron-job executions. And I use it for blazing-fast **Leaderboards** via its native Sorted Sets feature."

#### Indepth
Redis is fundamentally single-threaded (for executing commands). This guarantees native atomic operations without complex locking logic. An `INCR` command is mathematically guaranteed to never encounter a race condition between two competing microservices. This makes Redis exceptionally powerful as a distributed coordination mechanism.

---

### 100. What is cache warming?
"Cache warming (or pre-warming) is the proactive process of loading data into the cache *before* users actually request it.

If I run a major e-commerce site, and our 'Black Friday Sale' page launches at midnight, I anticipate a million simultaneous clicks at 12:00:01. If the cache is empty, my database will be utterly destroyed by Cache Misses on the first second.

At 11:50 PM, I execute a script that runs all the massive SQL queries necessary to render the Black Friday page and forcefully injects the final JSON into Redis. By the time midnight hits, the cache is already fully 'warm', seamlessly absorbing the devastating load."

#### Indepth
Cache warming is essentially the scheduled, defensive generation of materialized views in memory. It is incredibly common in batch-processing pipelines where a nightly Hadoop or Spark job crunches massive ML models and directly pushes the finalized recommendation lists precisely into the Redis caches right before peak morning hours.
