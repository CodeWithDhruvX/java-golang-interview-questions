# High-Level Design (HLD): Caching Patterns and Strategies

Proper caching is the easiest way to improve read performance in an enterprise application, but it introduces the hardest problem in computer science: Cache Invalidation.

## 1. What are the common Caching Patterns? (Cache Aside vs Write-Through)
**Answer:**
*   **Cache-Aside (Look-Aside):** The most common approach. The application checks the cache first. If it's a "hit", return the data. If it's a "miss", the application queries the database, puts the result into the cache, and returns the data. 
    *   *Pros:* Cache only contains data that is actually requested (memory efficient). If cache goes down, app can still query DB directly (resilient).
    *   *Cons:* Data can become stale if the DB is updated but the cache is not invalidated.
*   **Read-Through:** The application asks the cache for data. If it's a miss, the *cache provider itself* (not the application code) queries the DB, updates itself, and returns to the app. (Usually requires specific cache plugins).
*   **Write-Through:** The application writes data directly to the cache. The cache synchronously writes the data to the DB before returning success to the app.
    *   *Pros:* Ensures data in the cache is NEVER stale. Absolute consistency.
    *   *Cons:* Introduces extra latency on every write operation, since two systems must be updated synchronously.
*   **Write-Behind (Write-Back):** Application writes data to the cache, and the cache immediately returns success. Asynchronously, a background worker batches these updates and flushes them to the DB.
    *   *Pros:* Incredibly fast write speeds (you are writing at RAM speed).
    *   *Cons:* High risk of data loss. If the cache server crashes before flushing to the DB, the data is permanently gone.

## 2. Explain Cache Eviction Policies.
**Answer:**
Caches (like Redis or Memcached) have limited RAM. When the cache is full, old data must be evicted to make room for new data.
*   **LRU (Least Recently Used):** Evict the item that hasn't been accessed for the longest time. (Standard choice for web applications).
*   **LFU (Least Frequently Used):** Evict the item that has been accessed the fewest times overall.
*   **FIFO (First In, First Out):** Evict the oldest item in the cache, regardless of how often it's accessed.
*   **TTL (Time To Live):** A proactive approach. Every key is inserted with an expiration time (e.g., 60 minutes). The cache automatically deletes it when time runs out. Excellent for ensuring data eventually refreshes.

## 3. What is a Cache Stampede (Thundering Herd) and how do you prevent it?
**Answer:**
*   **The Problem:** Imagine a highly popular key (e.g., "Trending Topics on Twitter") is cached with a TTL. Thousands of users are requesting it constantly, so the cache is hit 100%. Suddenly, the TTL expires, and the key is deleted from the cache. Instantly, 1,000 concurrent requests all experience a "Cache Miss". All 1,000 requests bypass the cache and hit the database simultaneously to run the heavy "calculate trending" query. The database is crushed and goes offline.
*   **Solutions:**
    *   **Mutex Locks (Distributed Locking):** When the 1,000 requests miss the cache, force them to acquire a Redis lock before hitting the DB. Only Thread 1 gets the lock. Thread 1 queries the DB and repopulates the cache. Threads 2-1000 wait for 50ms, check the cache again (it's now repopulated!), and return.
    *   **Probabilistic Early Expiration (Cache Jitter):** Add randomness to the TTL. Instead of setting TTL to exactly 60 minutes for every item, set it to 60mins +/- a randomized percentage. This ensures large batches of keys do not all expire at the exact same millisecond. 
    *   **Background Refresh (Cron job):** For insanely popular keys that *must never* miss, do not put a TTL on them. Instead, have a background cron job forcefully overwrite the cache key with fresh DB data every 5 minutes. The user requests never trigger the database query directly.
