# Performance and Caching Basics (Service-Based Companies)

## 1. What are the basic strategies for improving API response times?

**Expected Answer:**
To improve an API that is responding slowly, I would look at the following areas:

1.  **Database Optimization:** Implement missing database indexes on queried columns. Avoid `N+1` query problems by using joins or batch fetching via ORM.
2.  **Caching:** Implement caching for frequently accessed, rarely changing data (using Redis or Memcached) to prevent hitting the database for every request.
3.  **Pagination:** Do not return massive arrays of objects. Implement RESTful pagination (Offset/Limit or Cursor-based).
4.  **Asynchronous Processing:** Move slow, non-blocking tasks (like sending emails, generating PDFs) to a background worker queue (e.g., Celery, RabbitMQ) instead of blocking the HTTP response.
5.  **Payload Reduction:** Only return the JSON fields the client actually needs (DTOs) and enable gzip/brotli compression at the API Gateway or web server level.

## 2. Explain the difference between Redis and Memcached.

**Expected Answer:**
Both are fast, in-memory key-value stores used primarily for caching, but they have distinct differences:

*   **Memcached:**
    *   Very simple, designed purely as a volatile string cache.
    *   Supports only simple key-value pairs (strings).
    *   Multithreaded architecture (highly efficient for simple GET/SET).
*   **Redis:**
    *   An in-memory data structure server.
    *   Supports complex data types: Strings, Hashes, Lists, Sets, Sorted Sets.
    *   Supports persistence (can save data to disk).
    *   Single-threaded event loop architecture (though highly optimized).
    *   Supports advanced features like Pub/Sub messaging, Lua scripting, and Geospatial indexes.
    *   *Conclusion:* Redis is the industry standard for almost all new projects due to its vastly superior feature set.

## 3. What is a CDN (Content Delivery Network), and why is it important?

**Expected Answer:**
*   **What it is:** A CDN (like Cloudflare, AWS CloudFront, Akamai) is a globally distributed network of proxy servers.
*   **How it works:** When a user requests a static asset (images, CSS, JS, videos), the CDN routes the request to the geographic "Edge Server" closest to the user, rather than routing the request all the way back to the application's Origin server.
*   **Importance:**
    1.  **High Performance/Low Latency:** Assets load much faster for users geographically far from the primary data center.
    2.  **Reduced Load:** Offloads massive amounts of traffic and bandwidth consumption from the main infrastructure.
    3.  **DDoS Protection:** CDNs naturally absorb massive traffic spikes and offer web application firewalls (WAF) to block malicious traffic at the edge.

## 4. Explain HTTP Caching headers (`Cache-Control`, `ETag`).

**Expected Answer:**
HTTP caching allows browsers and intermediate proxies to cache responses, saving network trips.

*   `Cache-Control`: The most important header. It dictates *who* can cache the response and for *how long*.
    *   `max-age=3600`: The resource is valid for 3600 seconds.
    *   `no-cache`: The browser can store the cache, but must validate with the server before using it.
    *   `no-store`: Do not cache this under any circumstances (crucial for sensitive banking data).
*   `ETag` (Entity Tag): A unique hash or string representing the specific version of the resource.
    *   When the browser requests a resource it has cached, it sends the header `If-None-Match: <ETag value>`.
    *   If the server sees the file hasn't changed, it replies with a `304 Not Modified` and an empty body, saving significant bandwidth.

## 5. What are Database Indexes, and what are their trade-offs?

**Expected Answer:**
*   **What they are:** An index is a data structure (usually a B-Tree) created on specific columns in a database table. It acts like an index at the back of a book, allowing the DB engine to quickly locate rows without scanning every single row in the table (a full table scan).
*   **Pros:** Radically speeds up `SELECT` queries, `WHERE` clauses, and `JOIN` operations.
*   **Trade-offs (Cons):**
    *   **Write Penalty:** Every time an `INSERT`, `UPDATE`, or `DELETE` occurs, the index must also be updated. This slows down write operations.
    *   **Storage Cost:** Indexes consume additional disk space and memory.
    *   *Rule of Thumb:* Only index columns that are frequently used in search conditions or joins; do not index every column.

## 6. What is a Cache Stampede (Thundering Herd), and how do you prevent it?

**Expected Answer:**
*   **The Problem:** A Cache Stampede occurs when a highly popular cached item (e.g., the homepage data) suddenly expires (cache miss). Simultaneously, hundreds of concurrent requests arrive. They all notice the cache miss, and they all query the expensive database at the exact same time to regenerate the data, causing a massive traffic spike that crashes the database.
*   **Prevention:**
    *   **Locking (Mutex):** The first request to notice the cache miss acquires a lock. Subsequent requests must wait for the lock to be released (meaning the first request has populated the cache) before they proceed, at which point they will hit the cache.
    *   **Stale-While-Revalidate:** Serve stale (slightly expired) data to the users while a background asynchronous thread updates the cache from the database.
