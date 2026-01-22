## 🟢 Basic System Design Concepts (Questions 1-10)

### Question 1: What is system design?

**Answer:**
System design is the process of defining the architecture, components, modules, interfaces, and data for a system to satisfy specified requirements. It involves making high-level decisions about the system's structure and behavior.

*   **Key Aspects:**
    *   **Architecture:** The overall structure (e.g., Monolithic, Microservices).
    *   **Scalability:** Ability to handle growth.
    *   **Reliability:** Ensuring the system works correctly and consistently.
    *   **Availability:** Ensuring the system is operational when needed.
    *   **Maintainability:** Ease of modification and updates.

### Question 2: Difference between high-level and low-level design.

**Answer:**

| Feature | High-Level Design (HLD) | Low-Level Design (LLD) |
| :--- | :--- | :--- |
| **Focus** | Overall system architecture and component interactions. | Internal logic of individual components. |
| **Scope** | Macroscopic view (Databases, Servers, APIs). | Microscopic view (Classes, Functions, Algorithms). |
| **Audience** | Architects, Stakeholders. | Developers, Testers. |
| **Outcome** | Architecture diagrams, Technology stack. | Class diagrams, Pseudo-code, Database schemas. |

### Question 3: What is scalability? Types of scalability?

**Answer:**
Scalability is the capability of a system to handle a growing amount of work by adding resources to the system.

*   **Vertical Scalability (Scaling Up):** Adding more power (CPU, RAM) to an existing server.
    *   *Pros:* Simple to implement.
    *   *Cons:* Limited by hardware capacity, single point of failure.
*   **Horizontal Scalability (Scaling Out):** Adding more machines to the pool of resources.
    *   *Pros:* Unlimited growth, fault tolerance.
    *   *Cons:* Complex to manage (data consistency, load balancing).

### Question 4: What is a load balancer?

**Answer:**
A load balancer is a device or service that distributes network or application traffic across a cluster of servers.

*   **Purpose:**
    *   Increases capacity (concurrent users) and reliability.
    *   Prevents any single server from becoming a bottleneck.
*   **Types:**
    *   **Layer 4 (Transport):** Based on IP/Port (e.g., TCP/UDP).
    *   **Layer 7 (Application):** Based on content (e.g., HTTP headers, cookies).

### Question 5: What is caching? Where can it be applied?

**Answer:**
Caching is the process of storing copies of files or data in a temporary storage location (cache) so that they can be accessed more quickly.

*   **Where it is applied:**
    *   **Browser/Client:** Local storage, Cookies.
    *   **CDN:** Static assets closer to users.
    *   **Load Balancer:** Reverse proxy caching.
    *   **Application Layer:** In-memory objects.
    *   **Database:** Query cache (e.g., Redis, Memcached).

### Question 6: What is CDN and how does it work?

**Answer:**
A Content Delivery Network (CDN) is a geographically distributed network of proxy servers and their data centers.

*   **How it works:**
    *   Users request content (images, videos, CSS).
    *   The request is routed to the nearest edge server (Point of Presence - PoP).
    *   If the content is cached, it's served immediately.
    *   If not, the CDN fetches it from the origin server, caches it, and serves it.
*   **Benefits:** Reduced latency, lower bandwidth costs, higher availability.

### Question 7: What is a reverse proxy?

**Answer:**
A reverse proxy is a server that sits in front of one or more web servers, intercepting requests from clients.

*   **Functions:**
    *   **Load Balancing:** Distributes traffic.
    *   **Security:** Hides backend server IPs, handles SSL/TLS termination.
    *   **Caching:** Serves static content.
    *   **Compression:** Compresses outgoing data (e.g., Gzip).

### Question 8: What is a message queue?

**Answer:**
A message queue is a form of asynchronous service-to-service communication used in serverless and microservices architectures.

*   **Mechanism:** Components send messages to a queue (Producer) and other components retrieve them (Consumer).
*   **Benefits:**
    *   **Decoupling:** Producers and consumers don't need to interact directly.
    *   **Scalability:** Consumers can be scaled independently.
    *   **Reliability:** Messages persist until processed.
*   **Examples:** RabbitMQ, Kafka, AWS SQS.

### Question 9: What is sharding?

**Answer:**
Sharding is a method of splitting and storing a single logical dataset in multiple databases.

*   **Process:** Breaking a large database into smaller, faster, more easily managed parts called "shards".
*   **Key Concept:** `Shard Key` determines which shard a row is stored in.
*   **Pros:** Horizontal scaling of writes and storage.
*   **Cons:** Complex queries (joins across shards), uneven data distribution (hotspots).

### Question 10: Difference between vertical and horizontal scaling.

**Answer:**

| Feature | Vertical Scaling (Scale Up) | Horizontal Scaling (Scale Out) |
| :--- | :--- | :--- |
| **Method** | Add resources (CPU, RAM) to one node. | Add more nodes to the system. |
| **Complexity** | Low. | High (requires LB, distributed data). |
| **Downtime** | May require downtime for upgrades. | No downtime (add nodes dynamically). |
| **Cost** | Exponential (high-end hardware is pricey). | Linear (commodity hardware). |
| **Limit** | Hard limit (hardware max). | Theoretically unlimited. |

---

## 🟢 Database Design & Management (Questions 11-20)

### Question 11: SQL vs NoSQL – when to use what?

**Answer:**

*   **Use SQL (Relational) when:**
    *   You need ACID compliance (Atomicity, Consistency, Isolation, Durability).
    *   Data is structured and schema remains constant.
    *   Complex queries (JOINs) are required.
    *   Examples: PostgreSQL, MySQL.
*   **Use NoSQL (Non-Relational) when:**
    *   Data is unstructured or semi-structured.
    *   You need high throughput and horizontal scalability.
    *   Flexible schema is needed (rapid development).
    *   Examples: MongoDB (Document), Redis (Key-Value), Cassandra (Wide-Column).

### Question 12: How do you scale a database?

**Answer:**
1.  **Vertical Scaling:** Upgrade the server hardware (limited).
2.  **Read Replicas:** Master-Slave architecture where writes go to Master and reads are distributed to Slaves.
3.  **Sharding (Horizontal Scaling):** Partition data across multiple servers.
4.  **Partitioning:** Split tables into smaller pieces (Vertical partitioning: split columns; Horizontal partitioning: split rows).
5.  **Caching:** Use Redis/Memcached to reduce DB load.

### Question 13: What is denormalization and why is it useful?

**Answer:**
Denormalization is the process of adding redundant data to a normalized database to reduce the number of joins and improve read performance.

*   **Why useful:**
    *   Optimizes read-heavy workloads.
    *   Avoids complex joins.
*   **Trade-off:**
    *   Writes become more complex (need to update multiple places).
    *   Risk of data inconsistency.

### Question 14: CAP theorem – explain and give examples.

**Answer:**
The CAP theorem states that a distributed system can only provide two of the following three guarantees:
1.  **Consistency:** Every read receives the most recent write or an error.
2.  **Availability:** Every request receives a (non-error) response, without the guarantee that it contains the most recent write.
3.  **Partition Tolerance:** The system survives network failures (message loss between nodes).

*   **Combinations:**
    *   **CP (Consistency + Partition Tolerance):** MongoDB, HBase (System unavailable if partition occurs to preserve consistency).
    *   **AP (Availability + Partition Tolerance):** Cassandra, DynamoDB (System always available, but might return stale data - Eventual Consistency).
    *   **CA:** Not possible in distributed systems (Network partitions are inevitable).

### Question 15: What is eventual consistency?

**Answer:**
A consistency model used in distributed computing to achieve high availability. It guarantees that, if no new updates are made to a given data item, eventually all accesses to that item will return the last updated value.

*   **Use Case:** DNS, Social Media feeds (it's okay if a post appears a few seconds later for some users).
*   **Contrast:** Strong Consistency (RDBMS), where data is instantly consistent across all nodes.

### Question 16: How would you design a schema for a social network?

**Answer:**
*   **Users Table:** `user_id` (PK), `username`, `email`, `password_hash`, `created_at`.
*   **Friendships Table:** `user_id1`, `user_id2`, `status` (pending, accepted), `created_at`.
*   **Posts Table:** `post_id` (PK), `user_id` (FK), `content`, `image_url`, `created_at`.
*   **Comments Table:** `comment_id`, `post_id` (FK), `user_id` (FK), `text`.
*   **Likes Table:** `post_id`, `user_id` (composite PK).
*   **Optimization:** Use Graph DB (Neo4j) for complex relationship queries (friends of friends).

### Question 17: Explain database partitioning.

**Answer:**
Database partitioning is the process of dividing a large database object (like a table) into smaller, manageable pieces, but still treating them as a single logical entity.

*   **Methods:**
    *   **Range Partitioning:** Based on a range of values (e.g., Dates: 2023, 2024).
    *   **List Partitioning:** Based on a list of values (e.g., Regions: US, EU, Asia).
    *   **Hash Partitioning:** Based on a hash function of a key (distributes data evenly).
    *   **Vertical Partitioning:** Moving infrequently used columns to another table.

### Question 18: How do you handle schema migrations?

**Answer:**
Schema migrations are managed versions of the database schema.
1.  **Versioning:** Each change is a script with a version number (V1__init.sql, V2__add_column.sql).
2.  **backward Compatibility:** Ensure changes don't break existing application code (e.g., add new column first, deploy code, then remove old column).
3.  **Tools:** Flyway, Liquibase.
4.  **Zero-downtime Approach:**
    *   Add nullable column.
    *   Dual write (write to both old and new).
    *   Backfill data.
    *   Switch reads to new.

### Question 19: What is a write-ahead log?

**Answer:**
The Write-Ahead Log (WAL) is a standard method for ensuring data integrity using the "logs" before "data" principle.

*   **Mechanism:** Modifications are written to a secure log file *before* they are applied to the database pages.
*   **Purpose:** Atomicity and Durability. If the system crashes, the database can re-execute the log to restore the data to a consistent state.

### Question 20: How to design a time-series database?

**Answer:**
A Time-Series Database (TSDB) is optimized for handling time-stamped data (metrics, events).

*   **Design Considerations:**
    *   **Write Heavy:** Optimized for high ingestion rates (append-only).
    *   **Storage:** efficient compression (delta encoding).
    *   **Query Patterns:** Range queries, aggregations (avg, max over time).
    *   **Retention:** Downsampling (rollups) and automatic expiration of old data.
*   **Example Technologies:** InfluxDB, Prometheus, TimescaleDB.

---

## 🟢 Caching & Performance (Questions 21-30)

### Question 21: How to cache data effectively?

**Answer:**
1.  **Identify Hot Data:** Cache frequently accessed, rarely changing data.
2.  **Right Layer:** Choose where to cache (CDN, API Gateway, App, DB).
3.  **TTL (Time To Live):** Set appropriate expiration to balance freshness and hit rate.
4.  **Eviction Policies:** Use LRU (Least Recently Used) for general purpose.
5.  **Serialization:** Store compact formats (Protobuf/MessagePack) to save space.

### Question 22: What is cache eviction policy? (LRU, LFU, FIFO)

**Answer:**
Determines which item to discard when the cache is full to make room for new items.
*   **LRU (Least Recently Used):** Discards the least recently used items first. (Most common).
*   **LFU (Least Frequently Used):** Discards items used least often. (Good for stable access patterns).
*   **FIFO (First In First Out):** Discards the oldest items first. (Simple queue).

### Question 23: Difference between Redis and Memcached.

**Answer:**

| Feature | Redis | Memcached |
| :--- | :--- | :--- |
| **Data Types** | Strings, Hashes, Lists, Sets, Sorted Sets. | Strings only. |
| **Persistence** | Yes (RDB snapshots, AOF logs). | No (In-memory only). |
| **Replication** | Yes (Master-Slave). | No. |
| **Pub/Sub** | Yes. | No. |
| **Use Case** | Complex Caching, Message Broker, Leaderboards. | Simple  Key-Value Caching. |

### Question 24: What are the downsides of caching?

**Answer:**
1.  **Stale Data:** Users might see outdated information.
2.  **Complexity:** Cache invalidation is hard to implement correctly.
3.  **Memory Cost:** RAM is expensive compared to disk.
4.  **Cache Penalties:** A "cache miss" adds latency (lookup cache + lookup DB).
5.  **Thundering Herd:** If cache goes down, DB might get overwhelmed.

### Question 25: How to handle cache invalidation?

**Answer:**
*   **Write-through:** Write to cache and DB simultaneously. (Consistent, but slow write).
*   **Cache-Aside (Lazy Loading):** App checks cache; if miss, reads DB and updates cache. (Inconsistent for a short time).
*   **TTL (Time-to-Live):** Auto-expire keys. (Simplest).
*   **Write-behind (Write-back):** Write to cache, async write to DB. (Fast write, risk of data loss).

### Question 26: What is a write-through vs write-back cache?

**Answer:**
*   **Write-Through:** Data is written to the cache and the backing store (DB) at the same time.
    *   *Pros:* Data consistency, reliability.
    *   *Cons:* Higher write latency.
*   **Write-Back (Write-Behind):** Data is written only to the cache initially and confirmed. The cache writes to the DB later in the background.
    *   *Pros:* Low write latency.
    *   *Cons:* Risk of data loss if cache fails before sync.

### Question 27: What happens when cache is full?

**Answer:**
When the cache reaches its memory limit, the **Eviction Policy** kicks in to remove items.
*   If **LRU** is configured, it removes the item accessed longest ago.
*   If **No Eviction** is configured, the cache will return errors on write operations.

### Question 28: How do you prevent cache stampede?

**Answer:**
Cache stampede (or thundering herd) occurs when a popular cache key expires, and many requests simultaneously hit the DB to regenerate it.
*   **Solutions:**
    *   **Locking/Mutex:** Only one process regenerates the key; others wait.
    *   **Probabilistic Early Expiration:** Expire key slightly before actual TTL randomly.
    *   **Use existing value:** Serve stale value while background process updates it.

### Question 29: What is CDN caching vs local caching?

**Answer:**
*   **CDN Caching:** Caches static content (images, JS, CSS) at edge servers globally. Reduces latency for geographically distributed users.
*   **Local Caching:** Caches data within the client's browser (localStorage) or the application server's memory. Reduces network requests to the backend.

### Question 30: Where do you place caching in a web architecture?

**Answer:**
1.  **Client Side:** Browser Cache.
2.  **DNS:** DNS Caching.
3.  **Web Server:** Reverse Proxy Cache (Nginx).
4.  **Application:** In-memory Cache (local map).
5.  **Distributed Cache:** Redis/Memcached cluster.
6.  **Database:** Internal Buffer Pool/Query Cache.

---

## 🟢 Load Balancing (Questions 31-40)

### Question 31: What is a load balancer?

**Answer:**
(Repeated concept but typically asked deeply)
A component that distributes incoming network traffic across multiple servers to ensure no single server bears too much load. It ensures high availability and reliability.

### Question 32: Types of load balancing strategies.

**Answer:**
1.  **Round Robin:** Sequential request distribution.
2.  **Least Connections:** Sends request to server with fewest active connections.
3.  **IP Hash:** Hashes client IP to assign a specific server (sticky).
4.  **Weighted Round Robin:** Assigns more requests to powerful servers.
5.  **Least Response Time:** select server with lowest latency.

### Question 33: How do you implement sticky sessions?

**Answer:**
Sticky sessions (Session Affinity) ensure a user always connects to the same server.
*   **Methods:**
    *   **Source IP Hashing:** Load balancer hash client IP.
    *   **Cookies:** LB injects a tracking cookie (e.g., `SERVERID`) into the response. Subsequent requests with this cookie are routed to the same server.
*   **Trade-off:** If that server fails, session data might be lost (unless stored in external Redis).

### Question 34: What are health checks in load balancing?

**Answer:**
Health checks are periodic probes sent by the LB to backend servers to ensure they are available.
*   **Active Check:** LB pings `/health` endpoint every X seconds.
*   **Passive Check:** LB monitors real traffic; if a server returns errors (5xx), it's marked unhealthy.
*   If a server is unhealthy, the LB stops sending traffic to it until it recovers.

### Question 35: Difference between Layer 4 and Layer 7 load balancers.

**Answer:**

| Feature | Layer 4 (Transport) | Layer 7 (Application) |
| :--- | :--- | :--- |
| **OSI Layer** | Transport Layer (TCP/UDP). | Application Layer (HTTP/HTTPS). |
| **Visibility** | Sees IP and Port only. | Sees URL, Headers, Cookies, Payroll. |
| **Logic** | Simple packet routing. | Smart routing (e.g., `/api` -> Service A, `/images` -> Service B). |
| **Performance** | Very High. | High (CPU intensive due to SSL/Parsing). |

### Question 36: What is DNS load balancing?

**Answer:**
Distributing traffic using the Domain Name System (DNS).
*   **Mechanism:** When a user resolves `example.com`, the DNS server returns one IP from a list of multiple server IPs (often typically Round Robin).
*   **Pros:** Simple, cheap.
*   **Cons:** DNS caching (by ISPs/Browsers) makes it hard to remove a down server instantly.

### Question 37: Explain round-robin vs least connections algorithm.

**Answer:**
*   **Round Robin:** 
    *   *Logic:* Cyclic order (Server A -> B -> C -> A).
    *   *Best for:* Servers with identical specs and stateless connections.
*   **Least Connections:**
    *   *Logic:* Dynamic. Routing to the server with the lowest current load (active connections).
    *   *Best for:* Long-lived connections (WebSocket) or varying request processing times.

### Question 38: How to handle load balancer failures?

**Answer:**
To prevent the LB from becoming a Single Point of Failure (SPOF):
*   **Active-Passive Setup:** Two LBs; one is active, the other is standby. They monitor each other using heartbeats (e.g., VRRP/Keepalived). If active fails, passive takes over the Virtual IP (VIP).
*   **Active-Active:** Both LBs accept traffic, distributed by DNS.

### Question 39: What is geo-load balancing?

**Answer:**
Distributing traffic based on the user's geographic location.
*   **Mechanism:** DNS or Global Traffic Manager (GTM) detects user IP.
*   **Routing:** Routes user to the nearest data center.
*   **Benefit:** Lowest latency, compliance with data residency laws.

### Question 40: How to design a multi-region load balancing setup?

**Answer:**
1.  **DNS Level (GSLB):** Route user to the nearest Region (e.g., US-East vs EU-West).
2.  **Regional LB:** At the entry of the region (AWS ALB/Gateway).
3.  **Local LB:** Distributes to internal microservices/pods.
4.  **Failover:** If an entire region goes down, GSLB updates DNS to route traffic to the next closest healthy region.

---

## 🟢 Design Patterns & Architecture (Questions 41-50)

### Question 41: What is microservices architecture?

**Answer:**
An architectural style where an application is structured as a collection of loosely coupled services.
*   **Characteristics:**
    *   Independently deployable.
    *   Highly testable and maintainable.
    *   Organized around business capabilities (Order Service, User Service).
    *   Owned by small teams.

### Question 42: What is monolithic architecture?

**Answer:**
A traditional unified model where the entire application is built as a single unit.
*   **Characteristics:** Single codebase, single build artifact (JAR/WAR), single deployment.
*   **Pros:** Simple to develop initially, easy debugging, no network latency between calls.
*   **Cons:** Hard to scale (must scale whole app), tight coupling, technology lock-in.

### Question 43: Difference between microservices and SOA.

**Answer:**

| Feature | Microservices | SOA (Service Oriented Architecture) |
| :--- | :--- | :--- |
| **Scope** | App-specific modularization. | Enterprise-wide integration. |
| **Communication** | Lightweight (HTTP/REST, gRPC). | Heavyweight (SOAP, ESB). |
| **Data** | Decentralized (Database per service). | Shared Data / Common Schema. |
| **Coupling** | Decoupled. | Loosely coupled via ESB. |

### Question 44: What is service discovery?

**Answer:**
The mechanism for services to find each other in a dynamic environment (like Kubernetes) where IPs change frequently.
*   **Client-Side Discovery:** Client queries Service Registry (e.g., Eureka) to get IP, then calls service.
*   **Server-Side Discovery:** Client calls Load Balancer; LB queries Registry and routes.
*   **Service Registry:** The database of available service instances (e.g., Consul, Etcd, Zookeeper).

### Question 45: How do services communicate in microservices?

**Answer:**
1.  **Synchronous:**
    *   **HTTP/REST:** Standard standard, human-readable.
    *   **gRPC:** High performance, binary (Protobuf), strict contract.
2.  **Asynchronous:**
    *   **Message Queues:** RabbitMQ, SQS (Point-to-Point).
    *   **Pub/Sub:** Kafka, SNS (Event-driven, one-to-many).

### Question 46: What is an API gateway?

**Answer:**
A server that acts as a single entry point into the system for clients.
*   **Responsibilities:**
    *   Request Routing (Reverse Proxy).
    *   Authentication & Authorization.
    *   Rate Limiting.
    *   SSL Termination.
    *   Request/Response Transformation.
*   Example: AWS API Gateway, Kong, Zuul.

### Question 47: What is the circuit breaker pattern?

**Answer:**
A design pattern used to detect failures and encapsulate the logic of preventing a failure from constantly recurring.
*   **States:**
    *   **Closed:** Requests flow normally.
    *   **Open:** Recent fault count exceeded threshold. Requests are blocked immediately (fail fast) to give the downstream service time to recover.
    *   **Half-Open:** Allow a few test requests. If successful, close circuit; else open again.

### Question 48: What is the saga pattern?

**Answer:**
A pattern for managing distributed transactions that span multiple microservices.
*   **Problem:** strict ACID is hard across services.
*   **Solution:** A sequence of local transactions. If one fails, execute **Compensating Transactions** to undo changes made by previous steps.
    *   *Choreography:* Events trigger next steps (Decentralized).
    *   *Orchestration:* Central coordinator manages flow (Centralized).

### Question 49: What is eventual consistency in microservices?

**Answer:**
Ensuring that data across different microservices becomes consistent over time, but not necessarily at the exact same instant.
*   **Usage:** Service A updates Order; publishes event `OrderCreated`. Service B consumes event and updates Inventory.
*   **Gap:** For a few milliseconds, Order exists but Inventory isn't updated. This is acceptable in many business flows.

### Question 50: How to ensure idempotency?

**Answer:**
Idempotency ensures that making the same request multiple times produces the same result (e.g., charging a card only once).
*   **Implementation:**
    *   Client sends a unique `idempotency-key` (UUID) with the request.
    *   Server checks if key exists in DB/Cache.
        *   If yes, return stored response (don't process again).
        *   If no, process and store the key + response.
