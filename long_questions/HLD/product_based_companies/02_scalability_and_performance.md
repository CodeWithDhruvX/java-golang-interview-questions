# High-Level Design (HLD): Scalability and Performance

This module focuses on how top tech product companies scale their systems to handle millions of persistent connections, terabytes of data, and extreme throughput requirements.

## 1. How do you scale a system from 0 to 1 Million Users? (Evolution)
**Answer:**
This is a classic evolutionary architecture question.
1.  **Single Server:** Everything (Web, App, DB) on one machine.
2.  **Separate DB:** Move Database to a separate server for better security, scaling, and backups.
3.  **Load Balancer + Multiple App Servers:** Introduce a LB. Application servers become stateless. (Horizontal Scaling).
4.  **Database Replication:** Setup Master-Slave replication. Route read traffic to read replicas to offload the master (which handles writes).
5.  **Caching Layer:** Introduce Memcached/Redis between App and DB to reduce DB read load.
6.  **CDN (Content Delivery Network):** Serve static assets (images, JS, CSS, videos) from edge servers close to users.
7.  **Message Queues:** Decouple monolithic workflows. Use Kafka/RabbitMQ for async processing (e.g., sending emails, generating reports).
8.  **Database Sharding:** When single master can't handle writes, horizontally partition the DB (Sharding).
9.  **Microservices:** Break down monolithic applications into domain-specific independent services.

## 2. Horizontal vs. Vertical Scaling. When to use which?
**Answer:**
*   **Vertical Scaling (Scale-Up):** Adding more CPU, RAM, or Disk to an existing server.
    *   *Pros:* Extremely simple, no architecture changes needed, no distributed data consistency issues.
    *   *Cons:* Hard hardware limit, single point of failure (SPOF), downtime required for upgrades.
    *   *When to use:* Early-stage startups, databases (initially, before sharding becomes strictly necessary).
*   **Horizontal Scaling (Scale-Out):** Adding more servers to the resource pool.
    *   *Pros:* Theoretically infinite scaling, built-in redundancy, high availability.
    *   *Cons:* High architectural complexity, requires load balancers, stateless servers, and distributed data handling.
    *   *When to use:* Modern web applications, large scale processing systems.

## 3. What are the common Database Scaling Techniques?
**Answer:**
1.  **Indexing:** Adding indexes to frequently queried columns speeds up read access dramatically (at the cost of slower writes and extra storage).
2.  **Denormalization:** Reducing joins by adding redundant data. Trades space for read speed.
3.  **Read/Write Splitting:** Master node routes writes, while multiple Read Replicas serve read queries. Good for read-heavy apps (e.g., blogs, e-commerce product pages).
4.  **Database Federation:** Splitting up databases by domain/function (e.g., one DB for users, one for product catalog).
5.  **Sharding (Horizontal Partitioning):** Splitting table data across multiple databases.

## 4. How do you handle High Concurrency (e.g., Flash Sales, Ticket Booking)?
**Answer:**
*   **Queueing:** Never let sudden burst traffic hit the DB directly. Push incoming HTTP requests/orders into a high-throughput message queue (like Kafka) and process them asynchronously at a rate the DB can handle.
*   **Caching with Atomic Operations:** Use Redis for inventory management. Redis is single-threaded and supports atomic operations like `DECR`.
    *   Example: Hold inventory in Redis. When order comes, `DECR` inventory. If `value >= 0`, accept order. Async worker updates DB.
*   **Rate Limiting / Virtual Waiting Rooms:** Prevent overload at the edge. Let only a configurable number of users into the payment flow, putting others in a queue.
*   **Optimistic Concurrency Control:** When updating DB records, avoid locking rows. Use a version number column. (e.g., `UPDATE tickets SET status=sold, version=2 WHERE id=123 AND version=1`).

## 5. What are the typical causes of high latency, and how do you reduce it?
**Answer:**
*   **Network/Distance:** Data traveling across the globe. -> *Solution:* CDNs, Multi-region deployments, Geo-routing.
*   **Database Queries:** Missing indexes, large joins, N+1 query problems. -> *Solution:* Caching, query optimization, read replicas.
*   **Synchronous Processing:** Waiting for slow 3rd party APIs or heavy compute tasks. -> *Solution:* Message queues, background workers.
*   **Garbage Collection Pauses:** Especially in Java/Go. -> *Solution:* Tuning GC, writing memory-efficient code, pooling objects.

## 6. Connection Pooling: Why is it important?
**Answer:**
Opening and closing database connections over TCP is expensive (3-way handshake, authentication, setup).
*   **Connection Pool:** Maintains a pool of active database connections in memory.
*   When a request needs DB access, it borrows a connection from the pool. Once done, it returns it rather than closing it.
*   *Benefit:* Drastically reduces latency and CPU overhead on both application and database servers. Limits the maximum number of concurrent connections the DB has to handle, preventing crashes.
