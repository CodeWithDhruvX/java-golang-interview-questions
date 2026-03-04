# High-Level Design (HLD): Database and Caching Strategies

Understanding when data should be stored, how it should be accessed, and how it should be cached is critical for HLD interviews at product-based companies.

## 1. When do you choose SQL over NoSQL?
**Answer:**
*   **SQL (Relational - MySQL, PostgreSQL):**
    *   *When to use:* When ACID compliance is strictly required (financial applications). 
    *   *When data is structured:* Schema is predefined, strict relationships exist.
    *   *When you need complex queries:* Heavy reliance on `JOIN`s, `GROUP BY`, and aggregations.
*   **NoSQL (Non-relational - MongoDB, Cassandra, DynamoDB):**
    *   *When to use:* Rapid development where schema flexibility is required.
    *   *Massive Scale:* Need massive horizontal scalability for unstructured or semi-structured data (e.g., logs, JSON documents, sensor data).
    *   *Specific Access Patterns:* E.g., Key-Value lookups (Redis/DynamoDB), Wide-column for time-series (Cassandra), Graph for relationships (Neo4j).

## 2. Explain standard Caching Strategies (Cache Aside, Read-Through, Write-Through, Write-Back).
**Answer:**
*   **Cache Aside (Lazy Loading):** Application requests data from cache. If hit, return. If miss, app fetches from DB, puts it into cache, then returns. 
    *   *Pros:* Cache only holds requested data. *Cons:* Initial miss incurs 3 network trips.
*   **Read-Through:** Application requests from cache. If miss, the *cache provider* automatically fetches from DB, updates itself, and returns to app.
    *   *Pros:* Simplifies app code.
*   **Write-Through:** App writes to Cache, and Cache immediately writes synchronously to DB.
    *   *Pros:* Data is always consistent. *Cons:* Write latency is higher.
*   **Write-Back (Write-Behind):** App writes ONLY to Cache. The Cache asynchronously flushes data to the DB in batches at intervals.
    *   *Pros:* Extremely fast writes (good for heavy write loads). *Cons:* Risk of data loss if the cache server crashes before flushing to DB.

## 3. How do you handle Cache Eviction?
**Answer:**
When memory is full, the cache must remove old data.
*   **LRU (Least Recently Used):** Evicts the item that hasn't been accessed for the longest time. Generally the most popular default policy.
*   **LFU (Least Frequently Used):** Evicts the item with the lowest access count.
*   **FIFO (First In First Out):** Evicts the oldest item inserted.
*   **TTL (Time To Live):** Items automatically expire after a set time. Crucial for data that naturally becomes stale (e.g., a stock price quote).

## 4. What is the Thundering Herd Problem (Cache Stampede) and how do you solve it?
**Answer:**
**Problem:** A highly popular cached item (e.g., homepage configuration) expires. Simultaneously, thousands of concurrent requests attempt to read it, get a cache miss, and all query the underlying Database at the exact same millisecond, bringing the DB down.
**Solutions:**
1.  **Mutex/Distributed Lock:** When a miss occurs, the first request acquires a lock to query the DB and update the cache. All other requests wait, and then read from the cache once the lock is released.
2.  **Probabilistic Early Expiration (Cache warming):** A background process re-computes and updates the cache *before* the TTL actually expires.
3.  **Stale-While-Revalidate:** Serve stale cached content to users instantly while a background thread asynchronously queries the DB to freshen the cache.

## 5. Cassandra vs. MongoDB vs. Redis: Compare uses cases.
**Answer:**
*   **Redis (In-memory KV Store):** Sub-millisecond latency. Use for caching, session management, leaderboard (Sorted Sets), pub/sub messaging, rate limiting. Data is volatile (though persistence mechanisms exist).
*   **MongoDB (Document Store):** High availability, dynamic schema (BSON/JSON). Use for Content Management Systems, catalog data, user profiles, where documents map directly to application objects. AP system (can be tuned).
*   **Cassandra (Wide-column Store):** Masterless decentralized architecture. Highly tuned for write-heavy workloads and linear scale. AP system. Use for time-series data, IOT sensor logs, scalable user activity tracking.

## 6. How does Database Replication work, and what is the Replication Lag problem?
**Answer:**
*   **Master-Slave:** Writes go to Master, which writes to an oplog/binlog. Slaves read this log to apply updates asynchronously.
*   **Replication Lag:** Because replication is usually asynchronous, a slave might be slightly behind the master. If a user updates their profile picture (write to master) and immediately refreshes the page (read from slave), they might see the old picture.
*   *Solutions for Replication Lag:*
    *   *Read-your-own-writes consistency:* If the user modifies data, route their reads for that specific data back to the master for a short window (e.g., 5 seconds). Let all other users' reads hit the slaves.
    *   Synchronous replication (causes severe write latency but guarantees consistency).
