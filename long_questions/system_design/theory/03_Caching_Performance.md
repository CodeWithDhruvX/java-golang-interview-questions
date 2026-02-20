# 🟡 Caching & Performance — Questions 21–30

> **Level:** 🟡 Mid-level (2–5 yrs)
> **Asked at:** Google, Meta, Netflix, Amazon, Swiggy, Hotstar — any system serving millions of users

---

### 21. How to cache data effectively?
"Caching effectively is about identifying *what* to cache, *where* to cache it, and *for how long*.

My mental model: cache data that is **frequently read, rarely written, and expensive to compute or fetch**. The product catalog on Swiggy? Cache it. A user's real-time wallet balance? Don't cache it — it changes too often and accuracy is critical.

For the 'where': I start with the layer closest to the user. Can a CDN serve it? Great — sub-millisecond from edge. Can a Redis cluster at the application tier serve it? Good — avoids the DB. The further from the user, the higher the latency.

For 'how long': TTL should match how often the data changes. For a restaurant menu, 5 minutes is fine. For a flight price, maybe 30 seconds. For a Slack user profile, an hour."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Swiggy, Zomato (menu/restaurant data), Hotstar, Netflix (content metadata), Razorpay (pricing rules)

#### Indepth
Five rules for effective caching:
1. **Cache at the right layer:** CDN → Reverse Proxy → Application Memory → Distributed Cache → DB query cache — each layer improves performance but adds complexity.
2. **Set meaningful TTLs:** Too long → stale data. Too short → DB pressure. Use business-domain knowledge to set TTL, not arbitrary numbers.
3. **Track metrics:** Monitor cache **hit rate** (>90% is good), miss rate, eviction rate. A 50% hit rate means caching isn't helping — your key distribution is wrong.
4. **Warm the cache:** On service startup or after clearing, the first requests will all be cache misses. Pre-populate critical keys at startup (cache warming) to avoid a cold start stampede.
5. **Test cache invalidation paths:** The most common bug is stale data in production after an update — ensure every write path invalidates the relevant cache key.

---

### 22. What is cache eviction policy? (LRU, LFU, FIFO)
"When a cache reaches its memory limit, it must remove something to make room. The eviction policy decides what gets removed.

**LRU (Least Recently Used)** removes the item that hasn't been read in the longest time. This is the most popular general-purpose policy — it operates on the intuition that 'if you haven't looked at it recently, you probably don't need it soon'. Redis defaults to LRU.

**LFU (Least Frequently Used)** removes the item that has been accessed the fewest times overall. Better for access patterns where some data is popular for a long time (like a viral YouTube video vs. a brand new one). But it can be slow to adapt to changing popularity.

**FIFO** just removes the oldest item regardless of access. Simple but rarely optimal."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Google (caching interviews), Amazon (ElastiCache design), Netflix (CDN cache policy)

#### Indepth
Advanced eviction policies:
- **LRU-K:** Instead of recency of last access, uses recency of K-th last access. More accurate popularity signal, especially used in database buffer pool management (PostgreSQL's clock-sweep algorithm).
- **ARC (Adaptive Replacement Cache):** Maintains two LRU lists — one for frequently accessed items (keeps those), one for recently accessed items. Self-tuning and used in ZFS file system.
- **SLRU (Segmented LRU):** Two-segment LRU — cached items start in "probation" segment; promoted to "protected" segment on second access. Used in Caffeine (Java) and inspired by TinyLFU.
- **TinyLFU + W-TinyLFU:** Used in Caffeine (Java in-process cache) and Ristretto (Go). Uses a probabilistic frequency sketch to estimate access frequency with very low memory overhead. Near-optimal hit rates for real workloads.

Redis supports: `allkeys-lru`, `allkeys-lfu`, `allkeys-random`, `volatile-lru` (evict only keys with TTL set). Choose based on your access pattern.

---

### 23. Difference between Redis and Memcached.
"Both are in-memory key-value caches, but they serve different needs.

I choose **Redis** for almost everything because it's dramatically more capable: it supports rich data types (strings, hashes, lists, sets, sorted sets, streams), persistence (RDB + AOF), Pub/Sub, and replication. Redis Sorted Sets are what enable leaderboards and rate limiting in a single data structure.

I'd only choose **Memcached** if I need simple string caching at extremely low per-server overhead with multi-threading that scales linearly with cores. Memcached is slightly faster for pure string gets at high concurrency — but for 99% of use cases, Redis is the better choice."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon (ElastiCache product team), Flipkart, Swiggy

#### Indepth
| Feature | Redis | Memcached |
|---|---|---|
| Data Types | String, Hash, List, Set, Sorted Set, Stream, Bitmap, HyperLogLog | String only |
| Persistence | Yes (RDB snapshots + AOF log) | No — purely in-memory |
| Replication | Yes (Master-Replica + Sentinel) | No |
| Clustering | Redis Cluster (built-in sharding) | Client-side sharding only |
| Pub/Sub | Yes | No |
| Scripting | Lua scripts (atomic multi-command) | No |
| Multi-threading | Single-threaded event loop (I/O multi-threaded from v6) | Multi-threaded (scales with cores) |
| Use Case | Caching + sessions + queues + leaderboards + rate limiting | Pure high-volume string caching |

**When Memcached wins:** Memcached uses less memory per stored item (simpler metadata) and can handle multi-threaded GET storms better for simple key→string workloads. Some teams at Facebook historically used Memcached for this reason. But Redis's versatility typically wins in modern architectures.

---

### 24. What are the downsides of caching?
"Caching is powerful but comes with real risks that I'm careful to account for in design.

The biggest risk is **stale data** — showing a user an outdated version of something that's changed. If a product's price drops but the cache still has the old price, a user might see an inflated number. Cache invalidation strategy must be designed upfront.

The second risk is **cache stampede** — when many requests simultaneously miss a popular cache key (after it expires or is cleared) and flood the DB. This can crash the DB. The third is **data inconsistency** in write-behind caches — the cache is updated but the DB write fails, leading to a permanent inconsistency."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Anywhere caching is used at scale — Netflix, Hotstar, Amazon, Flipkart

#### Indepth
Complete list of caching risks:
1. **Stale data:** Cache doesn't reflect the latest DB state. Solution: shorter TTL, event-driven invalidation.
2. **Cache stampede (thundering herd):** Many simultaneous cache misses → DB overload. Solution: mutex lock, probabilistic early refresh, or background refresh.
3. **Cache penetration:** Queries for data that doesn't exist in cache *or* DB (e.g., probing with random IDs). Each miss hits the DB. Solution: **Bloom filter** in front of cache — filter out known non-existent keys.
4. **Cache avalanche:** Many cached keys expire simultaneously (e.g., after a full cache flush/restart). DB is flooded. Solution: add **jitter** (random offset) to TTL values so evictions are distributed over time.
5. **Memory pressure:** Cache grows unbounded if eviction policy isn't tuned. Solution: maxmemory config + appropriate eviction policy.
6. **Increased operational complexity:** Another distributed system to monitor, scale, and debug.

---

### 25. How to handle cache invalidation?
"Cache invalidation is famously hard — Phil Karlton said 'there are only two hard things in computer science: cache invalidation and naming things'.

My preferred strategy for most systems is **event-driven invalidation**: whenever data changes in the DB, publish an event (via Kafka or Redis Pub/Sub) that invalidates the corresponding cache keys. This gives near-immediate consistency without coupling the writer to the cache.

For simpler systems, TTL-based expiry is sufficient — just accept that users may see data up to N minutes old. The key is to tune TTL to business tolerance for stale data, not to technical convenience."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon (strongly focused on their Dynamo paper themes), Google, Razorpay, Swiggy

#### Indepth
Four cache update strategies (the classic taxonomy):

- **Cache-Aside (Lazy Loading):** App checks cache → miss → read DB → write to cache → return. Pros: Only caches what's actually used. Cons: First request is always slow (cold start); potential for stale data between DB write and next cache miss.

- **Write-Through:** Every DB write also writes to cache. Pros: Cache always consistent. Cons: Write latency doubled; cache polluted with data that may not be read.

- **Write-Behind (Write-Back):** Write to cache only; async flush to DB later. Pros: Lowest write latency. Cons: Risk of data loss if cache crashes before DB sync. Only use when data loss is acceptable.

- **Refresh-Ahead:** Proactively refresh cache before TTL expires (based on predicted access). Used by Netflix for popular content. Reduces cold miss latency for predictable hot data.

**Event-driven invalidation** using CDC (Debezium reading Postgres WAL → Kafka → cache invalidation service) is the cleanest approach at scale — it decouples the cache invalidation from application code.

---

### 26. What is a write-through vs write-back cache?
"Both are about *when* data gets written to the backing store (DB) relative to the cache.

**Write-through:** Every write hits *both* cache and DB synchronously. The write only succeeds when both confirm. This gives you perfect cache-DB consistency at the cost of write latency — every write is as slow as the DB.

**Write-back (write-behind):** The write only updates the cache, and the cache asynchronously flushes to the DB later. Writes are very fast (just RAM speed) but there's a window where the DB is out of date. If the cache crashes in that window, you lose data.

I use write-through for anything where data loss is unacceptable (financial transactions, user data). Write-back for less critical systems where throughput matters more than durability (game leaderboards, analytics counters)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Paytm, PhonePe, Razorpay (financial data), Swiggy (leaderboards, analytics)

#### Indepth
A third variant: **Write-Around** — writes go directly to DB, bypassing the cache. Useful when you're writing large amounts of data that won't be re-read soon (like saving log files or bulk imports). Avoids polluting the cache with write-once data.

Combining strategies: Many systems use cache-aside for reads + write-through for writes. This is the most common real-world pattern:
1. On read: check cache → miss → read DB → populate cache
2. On write: write DB → immediately invalidate (or update) cache key
This avoids the write-behind's data loss risk while keeping reads fast.

**CPU cache analogy:** CPU L1/L2/L3 caches are write-back by default in modern processors — CPU cores write to cache, OS/hardware flushes dirty cache lines to RAM asynchronously. This is why `volatile` and memory barriers exist in concurrent programming — to force cache flushes.

---

### 27. What happens when cache is full?
"When the cache reaches maxmemory, the configured eviction policy kicks in. The cache engine evicts one or more existing keys to make room for the new key.

In Redis, the default behavior when maxmemory is hit without `maxmemory-policy` set is to **return an error on writes** — this is often a surprise to teams who haven't configured it. In production I always set `maxmemory-policy allkeys-lru` so Redis always evicts the least recently used key rather than rejecting writes.

The important thing is to monitor cache eviction rates. A high eviction rate is a signal that your cache is undersized or you're storing data that shouldn't be cached."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any Redis/caching deep-dive — Amazon ElastiCache configuration, Swiggy, Hotstar

#### Indepth
Redis `maxmemory-policy` options:
- `noeviction` (default): Returns error when memory limit reached. Safe but can break your application.
- `allkeys-lru`: Evict least recently used key from all keys. Best general purpose.
- `allkeys-lfu`: Evict least frequently used key. Better for long-lived hot data.
- `allkeys-random`: Evict random key. Bad — evicts hot data randomly.
- `volatile-lru`: Evict LRU only among keys with expiry set. Good when some keys must never be evicted.
- `volatile-lfu`: Evict LFU only among keys with expiry set.
- `volatile-ttl`: Evict key with shortest remaining TTL first.

**Memory optimization tips:** Use Redis data types wisely — store 100 field:value pairs in one Hash rather than 100 separate String keys. Hash with fewer than 128 fields uses ziplist encoding (very memory efficient). Use `OBJECT ENCODING key` to inspect. Tools like `redis-cli --bigkeys` and `redis-cli --hotkeys` (stat mode) help identify memory sinks.

---

### 28. How do you prevent cache stampede?
"Cache stampede (also called thundering herd) happens when a popular cache key expires, and hundreds or thousands of requests simultaneously see a cache miss and all run to the DB to recompute the value — crashing or slowing the DB dramatically.

My preferred solution is the **mutex (lock) pattern**: when a cache miss occurs, only one process acquires a lock to recompute the value. All other concurrent requests for the same key either wait for that process, or are served the slightly stale value (if it still exists). Redis `SET NX EX` implements a distributed mutex cleanly.

Another elegant approach is **probabilistic early expiration**: the cache proactively refreshes a key slightly *before* it expires, based on a random probability that increases as TTL approaches zero. This distributes the refresh work smoothly rather than having a cliff-edge stampede."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Hotstar (live streaming spike traffic), Netflix, Amazon, Google — any high-traffic system

#### Indepth
Detailed solutions for cache stampede:
1. **Mutex/Locking:** First request on miss: `SET lock:key placeholder NX EX 10` (set-if-not-exists with 10s expiry). If acquired → compute → set cache → release lock. Others: wait and poll, or serve stale.
2. **Probabilistic Early Expiration (XFetch algorithm):** `if (TTL - delta * beta * log(rand())) < 0 → refresh now`. Where `delta` = time to compute the value, `beta` ≈ 1. Elegant, no locks needed, spreads load.
3. **Promise/Future cache:** First request creates a Future, stores it in cache, starts async computation. Subsequent requests return the same Future and wait. Ensures single computation even under high concurrency.
4. **Eternal cache with background refresh:** Keys never expire (no TTL). A background job periodically refreshes hot keys before they go stale. Netflix uses this for the home page content tiles.
5. **Layered cache TTL with jitter:** Add random jitter to TTL: `TTL = base_ttl + random(0, 0.2 * base_ttl)`. Prevents synchronized mass expiration.

---

### 29. What is CDN caching vs local caching?
"CDN caching and local caching solve different problems, and a mature system uses both.

**CDN caching** caches content at geographically distributed edge nodes, solving the problem of **latency between users and your origin**. A user in Mumbai fetching your site's JS bundle from a Cloudflare edge in Pune gets it in <5ms vs 150ms from a server in the US. CDN caches static, public content — JS, CSS, images, videos.

**Local caching** (in-process or distributed Redis) caches computed data near your application server. It solves the problem of **expensive DB/service calls**. It caches dynamic, often personalized data — user sessions, API responses, computed aggregations. It's not geo-distributed; it's purpose-built to reduce load on your DB."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Netflix, Hotstar, Amazon, Flipkart (especially for high-traffic sale events)

#### Indepth
Complete view of where each type excels:

| Aspect | CDN Caching | Local/Distributed Caching |
|---|---|---|
| Location | Edge nodes globally (100s of PoPs) | Within or near your data center |
| Data Type | Static public content | Dynamic, often personalized data |
| Keys | URL-based | Application-defined key |
| Invalidation | API call (slow to propagate) or versioned URLs | Instant key deletion via app |
| Personalization | Not possible — same content for all | Fully personalized per user |
| Examples | Cloudflare, Akamai, AWS CloudFront | Redis, Memcached, in-process LRU |

**CDN for dynamic content:** Modern CDNs (Fastly, Cloudflare Workers) can run lightweight Edge Functions. Hotstar uses Fastly to serve personalized content lists from the edge by executing logic in CDN nodes — blurring the line between CDN caching and application-level caching.

---

### 30. Where do you place caching in a web architecture?
"Caching should exist at multiple layers simultaneously — each layer catches different types of requests and provides different trade-offs.

In a typical web architecture, I place caching: (1) **Browser** — via `Cache-Control` headers for static assets, (2) **CDN edge** — for globally distributed static content, (3) **Reverse proxy** — Nginx can cache static or semi-dynamic API responses, (4) **Application-level distributed cache** — Redis cluster for session data and database query results, (5) **Database buffer pool** — the DB itself caches frequently accessed pages in RAM.

The key is understanding that each layer complements the others, and a request only falls through to the next layer on a miss."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Netflix (caching architecture), Hotstar, Amazon, Flipkart, Swiggy

#### Indepth
The full caching stack for a web request to `GET /products/iphone-15`:

```
User Browser
  ↓ Cache-Control: max-age=3600 (hit → no request sent)
CDN (Cloudflare)
  ↓ Cache static assets, product images (TTL: 1hr)
Load Balancer (Nginx)
  ↓ Micro-cache: cache GET responses for 1s (helps with traffic spikes)
Application Server
  ↓ Check Redis: key = "product:iphone-15"
     Hit → return JSON in <1ms
     Miss ↓
  Database (PostgreSQL)
     → Query result returned
     → Application stores in Redis with TTL=5min
     → Returns result
  ↓ Application adds Cache-Control: max-age=300 to HTTP response
CDN caches API response for 5 min
```

**Design decision:** Should API responses be API-cached at the CDN? Yes, for public, non-personalized responses. No, for authenticated/personalized responses. Use `Vary: Authorization` header to ensure CDN doesn't serve one user's data to another.
