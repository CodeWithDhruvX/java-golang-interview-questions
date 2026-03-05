# Capacity Planning and Estimation (Product-Based Companies)

## 1. Walk me through a back-of-the-envelope calculation for a URL Shortener service (like bit.ly).

**Expected Answer:**
*This demonstrates the ability to translate vague requirements into technical constraints.*

**1. Assumptions & Traffic Estimation:**
*   **Traffic:** Let's assume 100 million new URLs generated per month.
*   **Read/Write Ratio:** URL shorteners are read-heavy. Let's assume 100:1.
*   **URL generations/month:** 100M.
*   **URL redirections/month:** 100M * 100 = 10 Billion.
*   **QPS (Queries Per Second) - Write:** 100M / (30 days * 24 hours * 3600 seconds) = ~40 URLs/second.
*   **QPS (Queries Per Second) - Read:** 10B / (30 * 24 * 3600) = ~4000 reads/second.

**2. Storage Estimation:**
*   Assume each URL mapping (original URL + short hash + metadata like user_id, timestamp) takes about 500 bytes.
*   We want to keep data for 5 years.
*   **Total URLs:** 100M/month * 12 months/year * 5 years = 6 Billion URLs.
*   **Total Storage:** 6 Billion * 500 bytes = 3 Terabytes (3TB).
    *   *Conclusion:* 3TB is easily handled by modern databases (even a single RDBMS node, though we'd shard for high availability and read throughput).

**3. Bandwidth Estimation:**
*   **Write Bandwidth:** 40 URLs/sec * 500 bytes = 20 KB/sec.
*   **Read Bandwidth:** 4000 URLs/sec * 500 bytes = 2 MB/sec.
    *   *Conclusion:* Very low bandwidth; will not be a bottleneck.

**4. Memory/Cache Estimation:**
*   Following the 80/20 rule, 20% of URLs generate 80% of traffic. We should cache the top 20% of daily read requests.
*   **Daily Read Requests:** 4000 reads/sec * 86400 seconds = ~350M requests/day.
*   **Calculate 20%:** 20% of 350M = 70M URLs.
*   **Cache Memory Needed:** 70M * 500 bytes = 35 Gigabytes (35GB).
    *   *Conclusion:* 35GB fits easily into a modern Redis cluster (or even a single beefy Redis node, though a cluster provides HA).

## 2. How do you map DAU/MAU to QPS to determine server requirements?

**Expected Answer:**
Translating Daily Active Users (DAU) to Queries Per Second (QPS) is crucial for provisioning.

1.  **Understand User Behavior:** Define what an "active user" does. E.g., for Twitter, maybe an active user views 50 tweets, posts 1 tweet, and likes 5 tweets per day.
2.  **Calculate Total Daily Operations:** DAU * Operations per user.
    *   Example: 10M DAU * 50 views = 500M views/day.
3.  **Calculate Baseline QPS:** Total Daily Operations / 86400 (seconds in a day).
    *   Example: 500M / 86400 ≈ 5,800 QPS.
4.  **Calculate Peak QPS:** Traffic is never uniform. A common heuristic is **Peak QPS = 2x or 3x Baseline QPS** depending on the application's timezone alignment (e.g., food delivery peaks strictly at lunch/dinner; global social media is more spread out).
    *   Example: Peak QPS = ~11,600 to 17,400 QPS.
5.  **Server Sizing:** Determine how many requests a single server instance (e.g., standard 4 vCPU, 16GB RAM) can handle. This depends on the runtime (Node, Java, Go) and workload (I/O bound vs. CPU bound). Assume 1 server handles 500 QPS.
    *   **Servers required:** 17,400 / 500 = 35 servers minimum (plus buffer for redundancy).

## 3. How do you prepare an architecture for extreme traffic spikes (e.g., Black Friday, Flash Sales, Super Bowl Ads)?

**Expected Answer:**
Standard auto-scaling is often too slow for instantaneous spikes (thundering herds). Preparation requires a multi-layered approach:

*   **Pre-scaling / Over-provisioning:** Schedule auto-scaling groups to increase minimum instances hours *before* the expected event. Don't rely on reactive CPU-based scaling.
*   **Aggressive Edge Caching (CDN):** Serve as much static and quasi-static content as possible from the CDN. Cache API responses using `Cache-Control` headers for endpoints that don't need real-time accuracy (e.g., product reviews, category listing).
*   **Queueing and Asynchronous Processing:** For write-heavy bursts (e.g., placing orders), do not write directly to the DB synchronously. Place the order event in a highly scalable message queue (Kafka, SQS) and return a "Accepted/Processing" status to the user. Workers drain the queue at the maximum safe rate the DB can handle.
*   **Degraded Application States (Shedding Load):** Implement feature flags to turn off non-critical features. During a flash sale, disable recommendations, user reviews, or complex search functionalities to reserve all compute power for the checkout flow.
*   **Circuit Breakers & Rate Limiting:** Strictly enforce rate limits at the API Gateway to prevent malicious or accidental DDoS. Use circuit breakers to protect downstream services from cascading failures.
*   **Database Read Replicas:** Scale out read replicas heavily to handle the read traffic surge.

## 4. What are the practical limitations of scaling a relational database, and how do you calculate when you need to shard?

**Expected Answer:**
Relational databases (like PostgreSQL/MySQL) primarily scale vertically (bigger machines).
*   **Limitations:**
    *   *Storage Limit:* A single disk volume has maximum size and IOPS limits.
    *   *CPU/Memory Limit:* You can only buy so much RAM/CPU on a single box.
    *   *Connection Limits:* OS-level limits on TCP connections; DBs spend too much time context-switching if there are tens of thousands of connections.
*   **Calculating Sharding Needs:**
    1.  **The Working Set Threshold:** Does the "active" data (indexes + frequently accessed rows) no longer fit into RAM? If caching and vertical scaling max out, data must be partitioned.
    2.  **Write Throughput (IOPS):** Read replicas scale reads, but all *writes* must go to the primary node. If the primary node hits its disk IOPS limit (~64k-256k on cloud SSDs depending on provider), you *must* shard to distribute writes across multiple disks/primaries.
    3.  **Storage Costs/Backup times:** When backups take > 24 hours or table rebuilds/migrations lock columns for unacceptable durations, the dataset is too large for a single instance.
*   *Note:* Sharding is extremely complex (joins break, global transactions break). Before sharding, attempt: functional partitioning (putting totally isolated tables on different DBs), caching, materialized views, and archiving old data.
