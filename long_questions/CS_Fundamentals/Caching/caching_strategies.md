# Caching Strategies

## 1. Cache-Aside (Lazy Loading)
The most common pattern. The application is responsible for reading from and writing to the cache.

### Flow
1.  **Read**: App checks Cache.
    *   *Hit*: Return data.
    *   *Miss*: App reads from Database, writes to Cache, then returns.
2.  **Write**: App writes to Database. Deletes (Invalidates) the key in Cache.

*   **Pros**: Resilient (if Cache fails, DB still works). Data model in cache can differ from DB.
*   **Cons**: First request is always slow (Cache miss). Stale data potential if Invalidation fails.

## 2. Read-Through
The application treats the Cache as the main data store. The Cache is responsible for fetching from DB.

### Flow
1.  **Read**: App asks Cache.
    *   *Hit*: Return data.
    *   *Miss*: **Cache** (not App) invokes a loader to read from DB, updates itself, and returns.

*   **Pros**: Simpler application code (App doesn't know about DB). Solves Thundering Herd problem better.
*   **Cons**: Cache provider needs a plugin/connector for the DB.

## 3. Write-Through
Data is written to the Cache and the Database at the same time.

### Flow
1.  **Write**: App writes to Cache. Cache synchronously writes to DB. Returns success only if both succeed.

*   **Pros**: High Consistency (Cache == DB). No stale data.
*   **Cons**: Higher Write Latency (2 writes).

## 4. Write-Reference (Write-Behind / Write-Back)
High performance, lower consistency.

### Flow
1.  **Write**: App writes to Cache. Cache immediately confirms success.
2.  **Async**: Cache asynchronously pushes updates to DB in background (batching).

*   **Pros**: Extremely fast writes. Reduced load on DB (deduplicated writes).
*   **Cons**: **Data Loss Risk**. If Cache crashes before syncing to DB, data is lost.
*   **Use Case**: Analytics counters, Likes.

## 5. Comparison Table

| Strategy | Read Speed | Write Speed | Consistency | Reliability |
| :--- | :--- | :--- | :--- | :--- |
| **Cache-Aside** | Fast (after miss) | Fast | Eventual | High |
| **Read-Through** | Fast (after miss) | N/A | Eventual | High |
| **Write-Through** | Fast | Slow | Strong | High |
| **Write-Back** | Fast | Very Fast | Weak | Low (Risk) |

## 6. Interview Questions
1.  **How to handle Cache Stampede (Thundering Herd)?**
    *   *Ans*: When a popular key expires, 1000 requests hit the DB simultaneously.
    *   *Solution 1*: **Probabilistic Early Expiration** (Simulate expiry before actual time).
    *   *Solution 2*: **Locking**. Only 1 thread queries DB, others wait.
2.  **Where to place a Cache?**
    *   *Ans*:
        *   **Client Side**: Browser Cache / CDN. (Fastest, Hardest to validate).
        *   **App Side**: In-memory (Ehcache). (Fast, but consumes App RAM).
        *   **Distributed**: Redis/Memcached. (Shared, Scalable).
