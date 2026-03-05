# ⚡ Caching Fundamentals — Interview Questions (Service-Based Companies)

This document covers caching concepts commonly tested at service-based companies like TCS, Infosys, Wipro, Capgemini, HCL. Targeted at 1–5 years of experience rounds.

---

### Q1: What is caching and why is it used?

**Answer:**
**Caching** is storing frequently accessed data in a fast-access storage layer (memory) so future requests for that data are served faster, reducing latency and load on origin systems (database, external APIs).

**Why caching is used:**
- **Reduce latency**: Memory access is ~100ns vs disk I/O ~10ms (100,000x faster).
- **Reduce database load**: Serve repeated queries from cache, not DB.
- **Improve throughput**: Handle more requests per second with same backend resources.
- **Cost savings**: Fewer DB queries = lower DB compute/cloud costs.

**Cache hit vs miss:**
- **Cache Hit**: Requested data is found in cache → served immediately.
- **Cache Miss**: Data not in cache → fetch from origin, store in cache, return to client.

**Hit Rate**: `Cache Hits / Total Requests × 100%`. Aim for >80-95% for effectiveness.

---

### Q2: What are the common cache eviction policies?

**Answer:**
When the cache is full, an eviction policy decides which items to remove.

| Policy | Full Name | How it works | Best for |
|---|---|---|---|
| **LRU** | Least Recently Used | Remove item least recently accessed | General purpose (most common) |
| **LFU** | Least Frequently Used | Remove item accessed fewest times | Long-lived, stable access patterns |
| **FIFO** | First In First Out | Remove oldest inserted item | Simple queues |
| **MRU** | Most Recently Used | Remove item most recently accessed | Scanning large datasets |
| **TTL** | Time To Live | Items expire after a set time period | Time-sensitive data |
| **Random** | Random Replacement | Remove random item | Simple, low overhead |

**LRU Implementation (conceptual):**
```
[HEAD (most recent)] → [B] → [A] → [C] → [TAIL (evict next)]
Access D (miss) → evict C → [D] → [B] → [A]
Access A (hit) → move to head → [A] → [D] → [B]
```

**Redis eviction policies:**
```
maxmemory-policy: allkeys-lru     # LRU across all keys (most common)
                  volatile-lru    # LRU only among keys with TTL set
                  allkeys-lfu     # LFU across all keys
                  noeviction      # Return error on write when full
```

---

### Q3: What is the difference between write-through, write-behind (write-back), and write-around caching?

**Answer:**

**Write-Through:**
- Write to cache AND database synchronously. Both updated before confirming to client.
- Pros: Cache is always consistent with DB. No data loss.
- Cons: Higher write latency (must wait for both cache + DB write).
- Use case: Data that is frequently read after write.

```
Client → Cache [ write ] → DB [ write ] → OK
```

**Write-Behind (Write-Back):**
- Write to cache only. Database is updated asynchronously (after a short delay or batch).
- Pros: Very fast writes (only cache write in critical path).
- Cons: Risk of data loss if cache fails before async DB write.
- Use case: Write-heavy workloads where slight inconsistency is acceptable (shopping cart, analytics counters).

```
Client → Cache [ write ] → OK  (DB updated asynchronously later)
```

**Write-Around:**
- Write directly to database, bypassing cache.
- Cache is populated only when the data is next read.
- Pros: Avoids caching data that might not be read again.
- Cons: First read after write always causes a cache miss.
- Use case: Large files, batch imports, data written once and rarely read soon after.

```
Client → DB [ write ] → OK  (cache NOT updated)
Next read → cache miss → fetch from DB → populate cache
```

---

### Q4: What are cache stampede / thundering herd and how do you prevent them?

**Answer:**
**Cache stampede** (Thundering Herd): When a popular cached item expires, many concurrent requests simultaneously miss the cache, hit the database at once, and overwhelm it.

```
10,000 req/sec on cached homepage
Cache expires → all 10,000 requests simultaneously hit DB
DB overloaded → cascading failure
```

**Prevention strategies:**

**1. Mutex / Cache Lock:**
- First request to miss acquires a lock, fetches from DB, populates cache.
- Other requests wait for the lock, then hit cache.
- Downside: Waiting adds latency.

**2. Cache-Aside with Early Expiry (Probabilistic Early Recomputation):**
- Before expiry, randomly start refreshing the cache early in a small % of requests.
- Distributes the refresh load.

**3. Background Refresh:**
- Never let cache truly expire. A background job proactively refreshes keys before they expire.
- Cache always has data; background refresh is decoupled.

**4. Stale-While-Revalidate:**
- Return stale (expired) data immediately.
- Refresh cache asynchronously in background.
- `Cache-Control: max-age=300, stale-while-revalidate=60`

**5. Jitter (Random TTL):**
- Instead of all items expiring at the exact same time, add random jitter to TTL.
  ```
  TTL = base_ttl + random(0, 30_seconds)
  ```
- Prevents mass simultaneous expiry of related cached items.

---

### Q5: What is Redis and what data structures does it support?

**Answer:**
**Redis** (Remote Dictionary Server) is an in-memory, open-source data store used as a cache, message broker, session store, and real-time data platform. All data is stored in RAM → ultra-low latency.

**Data structures:**

| Data Structure | Commands | Use Case |
|---|---|---|
| **String** | GET, SET, INCR, APPEND | Simple key-value, counters, session tokens |
| **List** | LPUSH/RPUSH, LPOP/RPOP, LRANGE | Queues, stacks, activity feeds |
| **Set** | SADD, SMEMBERS, SINTER, SUNION | Unique tags, followers list, deduplication |
| **Sorted Set (ZSet)** | ZADD, ZRANGE, ZRANK | Leaderboards, priority queues, rate limiting |
| **Hash** | HSET, HGET, HMGET | User profile object, product attributes |
| **Bitmap** | SETBIT, GETBIT, BITCOUNT | Feature flags, user activity tracking |
| **Stream** | XADD, XREAD, XGROUP | Event log, message queue (like Kafka, basic) |

**Common use cases:**
```bash
# Session storage
SET session:abc123 '{"userId":42,"role":"admin"}' EX 3600

# Rate limiter
INCR rate:user:42 → EXPIRE rate:user:42 60

# Leaderboard
ZADD game:scores 9800 "player_1"
ZREVRANGE game:scores 0 9 WITHSCORES  # Top 10

# Pub/Sub for real-time notifications
PUBLISH chat:room1 "Hello!"
SUBSCRIBE chat:room1
```

---

*Prepared for technical screening at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, Tech Mahindra).*
