# System Design Interview Questions & Answers (Summary Version)

> **Quick reference guide with concise explanations for System Design interview questions**

---

## 🔹 Basic System Design Concepts

**Q1: What is system design?**
System design is the process of defining the architecture, components, modules, interfaces, and data for a system to satisfy specified requirements. It involves high-level decisions about the structure and behavior of the system.

**Q2: Difference between high-level and low-level design.**
High-Level Design (HLD) focuses on the overall system architecture, component interactions, and data flow. Low-Level Design (LLD) details the implementation logic, class diagrams, and database schema for individual components initiated in the HLD.

**Q3: What is scalability? Types of scalability?**
Scalability is the capability of a system to handle a growing amount of work.
*   **Vertical Scalability (Scaling Up):** Adding resources (CPU, RAM) to a single server.
*   **Horizontal Scalability (Scaling Out):** Adding more servers to the pool of resources.

**Q4: What is a load balancer?**
A load balancer is a device or software that distributes network or application traffic across a cluster of servers to improve responsiveness and availability, preventing any single server from becoming a bottleneck.

**Q5: What is caching? Where can it be applied?**
Caching is the temporary storage of copies of data to allow for faster retrieval. It can be applied at various layers: client (browser), CDN, load balancer, web server, application, database, and distributed cache (like Redis).

**Q6: What is CDN and how does it work?**
A Content Delivery Network (CDN) is a group of geographically distributed servers that speed up the delivery of web content by bringing it closer to where users are. It caches static content like images, CSS, and JS files.

**Q7: What is a reverse proxy?**
A reverse proxy is a server that sits before web servers and forwards client requests to those web servers. It is used for load balancing, security, caching, and SSL termination.

**Q8: What is a message queue?**
A message queue is a form of asynchronous service-to-service communication used in serverless and microservices architectures. Messages are stored on the queue until they are processed and deleted by a consumer.

**Q9: What is sharding?**
Sharding is a method of splitting and storing a single logical dataset in multiple databases. It spreads the load across multiple servers, helping in horizontal scaling.

**Q10: Difference between vertical and horizontal scaling.**
Vertical scaling means adding more power to an existing machine, while horizontal scaling means adding more machines to the network. Vertical is easier but has a limit; horizontal is more complex but offers theoretically infinite scale.

---

## 🔹 Database Design & Management

**Q11: SQL vs NoSQL – when to use what?**
Use SQL (Relational) for structured data, complex queries, and ACID transactions (e.g., financial systems). Use NoSQL for unstructured data, high write throughput, horizontal scaling, and flexible schemas (e.g., content management, real-time analytics).

**Q12: How do you scale a database?**
Databases can be scaled vertically (upgrading hardware) or horizontally (sharding, read replicas, federation). Caching can also reduce database load.

**Q13: What is denormalization and why is it useful?**
Denormalization involves adding redundant data to a database to improve read performance by avoiding complex joins. It trades write performance and storage space for faster reads.

**Q14: CAP theorem – explain and give examples.**
CAP states a distributed system can only provide two of three guarantees: Consistency, Availability, and Partition Tolerance.
*   **CP:** MongoDB, HBase (Consistent, Partition Tolerant).
*   **AP:** Cassandra, DynamoDB (Available, Partition Tolerant).

**Q15: What is eventual consistency?**
A consistency model used in distributed computing to achieve high availability. It guarantees that, if no new updates are made to a given data item, eventually all accesses to that item will return the last updated value.

**Q16: How would you design a schema for a social network?**
Use a combination of SQL (for users, relationships) and NoSQL (for feed, posts). Tables/Collections might include Users, Posts, Follows, Likes. Graph databases (like Neo4j) are excellent for managing relationships (who follows whom).

**Q17: Explain database partitioning.**
Partitioning is dividing a database into distinct, independent parts.
*   **Horizontal partitioning (Sharding):** Splitting rows involves putting different rows into different tables.
*   **Vertical partitioning:** Splitting columns involves creating tables with fewer columns.

**Q18: How do you handle schema migrations?**
Use migration tools (like Flyway, Liquibase) to version control schema changes. Perform backward-compatible changes (e.g., add column before populating it), ensuring zero downtime during deployment.

**Q19: What is a write-ahead log?**
A Write-Ahead Log (WAL) is a technique where modifications are written to a log file before they are applied to the main database files. It ensures data integrity and recovery in case of a crash.

**Q20: How to design a time-series database?**
Optimize for high write ingestion and time-range queries. Use column-oriented storage, data compression, and downsampling (rolling up old data). Examples: InfluxDB, Prometheus.

---

## 🔹 Caching & Performance

**Q21: How to cache data effectively?**
Identify frequently accessed static data or expensive queries. Set appropriate TTL (Time-To-Live). access patterns (Read-Through, Write-Around).

**Q22: What is cache eviction policy? (LRU, LFU, FIFO)**
Rules to decide which data to remove when the cache is full.
*   **LRU (Least Recently Used):** Removes items not used for the longest time.
*   **LFU (Least Frequently Used):** Removes items with the fewest hits.
*   **FIFO (First In First Out):** Removes the oldest items.

**Q23: Difference between Redis and Memcached.**
*   **Memcached:** Multithreaded, simple key-value store, suitable for simple caching.
*   **Redis:** Single-threaded, supports complex data structures (lists, sets, hashes), persistence, and pub/sub.

**Q24: What are the downsides of caching?**
Cache invalidation complexity, memory costs, consistency issues (stale data), and "cold start" performance penalties.

**Q25: How to handle cache invalidation?**
Strategies include:
*   **Write-through:** Update cache and DB simultaneously.
*   **TTL:** Expire keys after a set time.
*   **Explicit Deletion:** Remove cache entries on data updates.

**Q26: What is a write-through vs write-back cache?**
*   **Write-through:** Data is written to cache and DB at the same time. Slower writes, stronger consistency.
*   **Write-back (Write-behind):** Data is written to cache first and asynchronously to DB. Faster writes, risk of data loss on crash.

**Q27: What happens when cache is full?**
The cache eviction policy (e.g., LRU) takes over to remove existing items to make space for new ones.

**Q28: How do you prevent cache stampede?**
Occurs when many requests hit the DB simultaneously because a cache key expired. Solutions: Locking (only one process recomputes), probabilistic early expiration, or "thundering herd" protection.

**Q29: What is CDN caching vs local caching?**
*   **CDN Caching:** Distributed globally, caches static assets closer to users.
*   **Local Caching:** In-memory on the application server/browser, faster access but limited scope and coherence issues.

**Q30: Where do you place caching in a web architecture?**
Everywhere: Browser, CDN (edge), Load Balancer, Web Server (local cache), Application (Redis/Memcached), Database (buffer pool).

---

## 🔹 Load Balancing

**Q31: What is a load balancer?**
A component that distributes incoming network traffic across multiple servers to ensure no single server bears too much load.

**Q32: Types of load balancing strategies.**
*   **Round Robin:** Sequential request distribution.
*   **Least Connections:** Sends traffic to the server with fewest active connections.
*   **IP Hash:** Maps IP address to a specific server (sticky).

**Q33: How do you implement sticky sessions?**
Use cookies or IP hashing to ensure a user's session always routes to the same server. Useful for stateful applications but can cause uneven load distribution.

**Q34: What are health checks in load balancing?**
Periodic checks (pings/requests) sent to backend servers to ensure they are available. If a server fails, the LB stops sending traffic to it.

**Q35: Difference between Layer 4 and Layer 7 load balancers.**
*   **Layer 4 (Transport):** Routes based on IP and Port (TCP/UDP). Faster, less context-aware.
*   **Layer 7 (Application):** Routes based on content (URL, Headers, Cookies). Slower (CPU intensive) but smarter routing.

**Q36: What is DNS load balancing?**
Distributing traffic at the DNS level by configuring multiple A records for a domain. The DNS server returns different IP addresses to different clients (often in Round-Robin).

**Q37: Explain round-robin vs least connections algorithm.**
Round-robin simply cycles through the list of servers. Least connections is dynamic; it checks which server is currently handling the fewest requests and sends new traffic there.

**Q38: How to handle load balancer failures?**
Use a high-availability pair (Active-Passive or Active-Active) with a heartbeat mechanism. If the active LB fails, the passive one takes over the Virtual IP.

**Q39: What is geo-load balancing?**
Distributing traffic based on the user's geographic location. DNS servers resolve the domain to the IP address of the data center closest to the user.

**Q40: How to design a multi-region load balancing setup?**
Use Global Traffic Management (DNS-based GSLB) to route users to the nearest region. Within the region, use local load balancers to distribute traffic to application servers.

---

## 🔹 Design Patterns & Architecture

**Q41: What is microservices architecture?**
An architectural style where an application is structured as a collection of loosely coupled, independently deployable services organized around business capabilities.

**Q42: What is monolithic architecture?**
A traditional unified model where all distinct components of the application (UI, business logic, data access) are built and deployed as a single unit or codebase.

**Q43: Difference between microservices and SOA.**
SOA (Service-Oriented Architecture) focuses on enterprise-wide service reuse and uses an Enterprise Service Bus (ESB). Microservices focus on decoupling and agility, often using lightweight protocols (HTTP/REST) and are bounded by context.

**Q44: What is service discovery?**
Automatic detection of devices and services on a computer network. In microservices, it allows services to find each other dynamically without hardcoded IPs (e.g., using Consul, Eureka, or Kubernetes DNS).

**Q45: How do services communicate in microservices?**
*   **Synchronous:** HTTP/REST, gRPC.
*   **Asynchronous:** Message queues (RabbitMQ, Kafka), Pub/Sub.

**Q46: What is an API gateway?**
A server that acts as a single entry point for the system. It handles routing, composition, authentication, rate limiting, and caching for requests from clients to backend services.

**Q47: What is the circuit breaker pattern?**
A design pattern used to detect failures and encapsulate the logic of preventing a failure from constantly recurring (e.g., stopping requests to a dead service to prevent cascading failures).

**Q48: What is the saga pattern?**
A sequence of local transactions where each updates data and publishes an event or message to trigger the next transaction in the saga. Used for managing distributed transactions across microservices.

**Q49: What is eventual consistency in microservices?**
The concept that data changes in one service might take time to propagate to others. System components become consistent over time rather than immediately.

**Q50: How to ensure idempotency?**
Use unique idempotency keys (like a UUID) in requests. The server checks if the key was already processed; if so, it returns the previous result instead of re-executing the operation.

---

## 🔹 Scalability and Availability

**Q51: How do you design a system that handles millions of users?**
Focus on horizontal scaling, decoupling components (microservices), using caching layers, load balancing, database sharding/partitioning, and asynchronous processing (queues).

**Q52: How to scale a system read-heavy workload?**
Use Caching (Redis/CDN), Read Replicas for databases, and Denormalization.

**Q53: How to scale a system write-heavy workload?**
Use Sharding, Asynchronous processing (Message Queues), and write-optimized databases (NoSQL like Cassandra or LSM-tree based).

**Q54: What is replication and when to use it?**
Copying data across multiple servers. Use it to increase data availability, redundancy (backup), and read performance (read replicas).

**Q55: How to make a system fault-tolerant?**
Eliminate single points of failure (redundancy), implement automatic failover, use circuit breakers, and ensure data replication.

**Q56: What is failover?**
The automatic switching to a standby computer server, system, hardware component, or network upon the failure or abnormal termination of the previously active application.

**Q57: What is high availability?**
A system design approach (measured in "nines", e.g., 99.999%) ensuring the system is operational and accessible for a high percentage of time. Achieved via redundancy and failover.

**Q58: Difference between active-passive and active-active systems.**
*   **Active-Passive:** One node handles traffic; the other waits as backup.
*   **Active-Active:** Both nodes handle traffic, improving performance and redundancy.

**Q59: What is graceful degradation?**
The ability of a system to maintain limited functionality even when a large portion of it has been destroyed or rendered inoperative.

**Q60: What is a throttling mechanism?**
Controlling the rate at which requests are processed to prevent server overload. It can delay or reject requests exceeding a defined limit.

---

## 🔹 Security & Authentication

**Q61: How to design a secure login system?**
Use HTTPS, salt and hash passwords (bcrypt/Argon2), enforce strong password policies, implement MFA (Multi-Factor Authentication), and protect against brute-force attacks.

**Q62: What is OAuth 2.0?**
An authorization framework that enables applications to obtain limited access to user accounts on an HTTP service without exposing user credentials (e.g., "Log in with Google").

**Q63: What is JWT?**
JSON Web Token. A compact, URL-safe means of representing claims to be transferred between two parties. Commonly used for stateless authentication.

**Q64: How do you store passwords securely?**
Never store in plain text. Use a strong cryptographic hash function (like bcrypt, scrypt, or Argon2) with a unique salt for each user.

**Q65: What is rate limiting?**
Restricting the number of requests a user or client can make to a server within a specified time frame to prevent abuse and ensure stability.

**Q66: How to secure APIs?**
Use HTTPS, Authentication (OAuth2/JWT), Rate Limiting, Input Validation, and proper error handling (don't leak sensitive info).

**Q67: What is CORS?**
Cross-Origin Resource Sharing. A browser security feature that restricts web pages from making requests to a different domain than the one that served the web page, unless the server explicitly allows it.

**Q68: Explain SSL/TLS in web communication.**
Protocols for establishing authenticated and encrypted links between networked computers. It ensures data privacy and integrity during transmission.

**Q69: What is cross-site request forgery (CSRF)?**
An attack that forces an end user to execute unwanted actions on a web application in which they are currently authenticated. Prevented using anti-CSRF tokens.

**Q70: What is cross-site scripting (XSS)?**
An attack where malicious scripts are injected into otherwise benign and trusted websites. Prevented by escaping/sanitizing user input.

---

## 🔹 Design Specific Systems (Popular Questions)

**Q71: Design YouTube**
Focus on: BLOB storage for videos, CDN for streaming, separate services for upload/transcoding, metadata database (SQL/NoSQL), and recommendation engine.

**Q72: Design WhatsApp**
Focus on: Real-time messaging (WebSocket/XMPP), E2E encryption (Signal Protocol), temporary message storage (store-and-forward), and offline handling.

**Q73: Design Twitter**
Focus on: Newsfeed generation (Fanout-on-write vs Fanout-on-read), high availability, timeline cache (Redis), and dealing with "celebrity" nodes (hot partitioning).

**Q74: Design Uber**
Focus on: Geospatial database (Quadtree/Geohash), real-time location tracking (WebSockets), matchmaking service, and state machine for trip status.

**Q75: Design Instagram**
Focus on: Image storage (S3), metadata (PostgreSQL/NoSQL), relationships (Graph DB), and feed generation (Pre-computed).

**Q76: Design a URL Shortener (like bit.ly)**
Focus on: Hashing algorithms (Base62), high read throughput, separation of redirection service (fast) and management service, and key generation service (KGS).

**Q77: Design a file storage system (like Dropbox or Google Drive)**
Focus on: Block-level storage, chunking files, de-duplication, synchronization client, and meta-data database (to track chunks).

**Q78: Design a news feed (like Facebook)**
Focus on: Aggregation service, ranking algorithms, feed publishing (Pull vs Push model), and caching recent feeds in memory.

**Q79: Design a video streaming service (like Netflix)**
Focus on: Adaptive Bitrate Streaming, CDN distribution, content encryption (DRM), and Open Connect appliances.

**Q80: Design an e-commerce platform (like Amazon)**
Focus on: Product catalog (NoSQL), Inventory management (ACID/SQL), Shopping cart (Redis), Checkout process, and recommendation systems.

---

## 🔹 Monitoring, Logging & DevOps

**Q81: How to monitor a distributed system?**
Use a combination of Metrics (Prometheus), Logging (ELK Stack), and Tracing (Jaeger/Zipkin). Monitor the four golden signals: Latency, Traffic, Errors, and Saturation.

**Q82: What are metrics, logs, and traces?**
*   **Metrics:** Aggregatable numerical data (CPU %, Req/sec).
*   **Logs:** Discrete events (text) happening at specific times.
*   **Traces:** The path of a specific request through multiple microservices.

**Q83: What is Prometheus and Grafana?**
Prometheus is a monitoring system and time-series database. Grafana is an analytics and visualization platform used to chart Prometheus data.

**Q84: What is centralized logging?**
Aggregating logs from all servers/containers into a single location (like Elasticsearch or Splunk) for easy searching and analysis.

**Q85: How to detect system bottlenecks?**
Analyze performance metrics (CPU, Memory, I/O) and use distributed tracing to find which service or query is taking the most time in a request chain.

**Q86: How do you debug a distributed system?**
Use Distributed Tracing (Correlation IDs), centralized logs, and reproduce issues in a staging environment.

**Q87: What is canary deployment?**
Releasing a change to a small subset of users first to test functionality and performance before rolling it out to everyone.

**Q88: What is blue-green deployment?**
Maintaining two identical production environments (Blue and Green). One is live, the other is idle. New code is deployed to the idle ecosystem, tested, and then traffic is switched.

**Q89: How do you handle configuration in distributed systems?**
Use centralized configuration services (like Consul, Etcd, or Spring Cloud Config) to manage settings dynamically without redeploying.

**Q90: What is chaos engineering?**
The discipline of experimenting on a system (e.g., randomly killing services/pods) to build confidence in the system's capability to withstand turbulent conditions in production.

---

## 🔹 Miscellaneous & Advanced Topics

**Q91: What is a web socket? How is it different from HTTP?**
WebSocket provides a persistent, full-duplex communication channel over a single TCP connection, ideal for real-time apps. HTTP is request-response and usually stateless.

**Q92: How do you design rate limiting?**
Algorithms: Token Bucket, Leaky Bucket, Fixed Window, Sliding Window. Use a fast store like Redis to track counts per IP/User.

**Q93: How to handle distributed transactions?**
Use Two-Phase Commit (2PC) (rigid, blocking) or Sagas (flexible, eventual consistency) with compensating transactions.

**Q94: What is the role of Zookeeper in distributed systems?**
It acts as a centralized service for maintaining configuration information, naming, providing distributed synchronization, and group services.

**Q95: What is eventual vs strong consistency?**
*   **Strong:** Reads always return the most recent write.
*   **Eventual:** Reads might return stale data, but eventually, all nodes will become consistent.

**Q96: What is leader election?**
The process of designating a single process as the organizer of some task among several distributed computers (nodes). Algorithms: Paxos, Raft, Bully Algorithm.

**Q97: What is CRDT?**
Conflict-free Replicated Data Type. A data structure that can be replicated across multiple computers in a network, where the replicas can be updated independently and resolve inconsistencies mathematically without coordination.

**Q98: Explain Raft or Paxos consensus algorithms.**
Protocols to allow a network of unreliable processors to agree on a single value/state. Raft is designed to be easier to understand than Paxos, using a leader-based approach.

**Q99: How to design a cron job scheduler?**
Use a distributed scheduler (like Quartz) or a simple DB-backed solution with a lock. For scale, use a persistent queue and workers.

**Q100: What is the difference between throughput and latency?**
*   **Throughput:** How much work (requests) the system can perform in a given time.
*   **Latency:** The time it takes for a single request to be processed and a response received.

---

## 🔸 Advanced Distributed Systems

**Q101: How does a distributed hash table (DHT) work?**
A DHT is a decentralized storage system where keys are mapped to nodes using a hash function (like Consistent Hashing). It allows for efficient lookup (typically O(log N)) without a central coordinator (e.g., Chord, Kademlia).

**Q102: Explain the concept of quorum in distributed systems.**
Quorum is the minimum number of votes that a distributed transaction has to obtain to be allowed to perform an operation in a distributed system. For a system with N nodes, a quorum is usually (N/2 + 1) to ensure consistency.

**Q103: What is vector clock? How is it used in conflict resolution?**
Vector clocks are an algorithm for generating a partial ordering of events in a distributed system and detecting causality violations. They attach a vector of counters to each message to track versions across nodes, helping resolve conflicts (e.g., in DynamoDB).

**Q104: Compare Raft vs Paxos.**
Both are consensus algorithms. Paxos is the original, mathematically proven but complex to implement. Raft was designed for understandability, using a strong leader approach and separating leader election from log replication.

**Q105: What is gossip protocol?**
APeers in a distributed system periodically exchange state information with random other peers. It spreads information (like node health or configuration updates) through the network like a virus/gossip (e.g., Cassandra).

**Q106: How do you detect and recover from split brain in distributed systems?**
Split brain occurs when network failure divides a cluster into two sub-clusters, both believing they are the active master. Solutions: Quorum (majority vote), Fencing (isolating the old leader), or Tie-breaker nodes.

**Q107: How to ensure consistency in leader-follower replication?**
Use synchronous replication (leader waits for follower ack) for strong consistency. For eventual consistency, use asynchronous replication and handle conflicts on read or write.

**Q108: What is lease-based leader election?**
Nodes acquire a "lease" (lock with a timeout) from a central store (like Zookeeper/ETCD). The node holding the lease is the leader. It must renew the lease before it expires to maintain leadership.

**Q109: Explain the difference between synchronous and asynchronous replication.**
*   **Synchronous:** Write is committed only after being written to both primary and replica. Slower, zero data loss.
*   **Asynchronous:** Write is committed on primary immediately; replica updates later. Faster, potential data loss on failover.

**Q110: What is two-phase commit protocol?**
A distributed algorithm to ensure all nodes in a transaction commit or abort. Phase 1: Coordinator asks "Can you commit?" (Prepare). Phase 2: If all say yes, "Commit"; otherwise "Abort".

---

## 🔸 Advanced Caching & Optimization

**Q111: How do you design a multi-layer cache system?**
Use a hierarchy: L1 (Local/In-Memory Cache on App Server), L2 (Distributed Cache likes Redis), L3 (CDN/Database Cache). Check L1 -> L2 -> L3 -> DB.

**Q112: Explain cache coherence in distributed cache.**
Ensuring that multiple caches sharing common resources reflect the same data. In distributed systems, this is often loosened to eventual consistency or managed via invalidation messages (pub/sub) when data changes.

**Q113: What is cache warming?**
The process of pre-populating the cache with data (e.g., after a restart or deployment) before it starts handling live traffic, to prevent a performance dip from cache misses.

**Q114: What is negative caching?**
Caching the result of a "miss" or "not found" (e.g., 404 error) to prevent repeated lookups for data that is known to not exist.

**Q115: How does content-based cache invalidation work?**
Updating or removing cache entries based on changes to the underlying content (tags/keys) rather than just time. E.g., when "Product A" is updated, invalidate all cache keys associated with "Product A".

**Q116: How would you cache personalized data?**
Personalized data (like "My Feed") is hard to cache. Strategy: Cache the common components (content metadata) globally, and compute the personalized list on the fly, or store the pre-computed list ID per user with a short TTL.

**Q117: Design a global caching strategy for a multi-region app.**
Use a CDN for static assets. For dynamic data, use region-specific Redis clusters.Replicate frequently accessed, low-change global data across regions.

**Q118: How do you handle stale cache data?**
Accept it for some use cases (Eventual Consistency). For strict cases, use short TTLs, versioning (Key-v2), or write-through strategies.

**Q119: What is TTL (Time to Live) in caching and how to set it optimally?**
TTL is the duration data stays in cache. Set it based on data volatility: Static data = Long TTL; Real-time data = Very Short/No TTL.

**Q120: How to handle write-heavy cache scenarios?**
Write-heavy data causes frequent invalidation. Strategy: Write-Back caching (buffer writes in cache, flush to DB periodically) or don't cache deeply; just cache reads.

---

## 🔸 Data Flow & Streaming Systems

**Q121: How does Apache Kafka work? Design considerations?**
Kafka is a distributed event store and stream-processing platform. It uses a commit log architecture, partitioned topics, and consumer groups. Design for partition count (parallelism) and retention policies.

**Q122: How do you ensure message ordering in a distributed system?**
In Kafka, order is guaranteed only *within* a partition. To ensure total global order is hard/slow; usually, you ensure related data (e.g., same User ID) goes to the same partition using a partition key.

**Q123: How to design a real-time data analytics pipeline?**
Source (API/DB) -> Ingestion (Kafka/Kinesis) -> Processing (Flink/Spark Streaming) -> Storage (Druid/ClickHouse) -> Visualization (Grafana/Superset).

**Q124: Difference between stream processing vs batch processing.**
*   **Stream:** Real-time, low latency, processes one event at a time (or small windows).
*   **Batch:** High latency, processes large chunks of data at scheduled intervals.

**Q125: Design a log aggregation system.**
Agents (Filebeat/Fluentd) collect logs -> Buffer (Kafka) -> Indexer (Logstash) -> Storage (Elasticsearch) -> UI (Kibana).

**Q126: Explain publish-subscribe vs message queue pattern.**
*   **Queue:** Point-to-point. One producer, one consumer (competitors). Message is gone once consumed.
*   **Pub/Sub:** One producer, multiple subscribers. Message is broadcast to all interested parties.

**Q127: How would you handle duplicate messages in streaming systems?**
Use Idempotent Consumers. Track processed Message IDs in a fast store (Redis). If ID exists, skip processing.

**Q128: Design an event-driven architecture.**
Services communicate by emitting events ("OrderPlaced") rather than direct calls. Other services listen and react. Decouples systems but increases complexity in tracing flow.

**Q129: What is backpressure in streaming systems?**
When the consumer cannot keep up with the producer. The system must slow down the producer or buffer data (up to a limit) to prevent crash/OOM.

**Q130: How to scale Kafka consumers?**
Add more consumer instances to the Consumer Group. Note: You can only scale up to the number of partitions in the topic.

---

## 🔸 File & Media Systems

**Q131: How to design a file deduplication system?**
Calculate a hash (SHA-256) of the file content. If hash exists, verify byte-by-byte (to avoid collisions) and point to existing file instead of storing a duplicate.

**Q132: How to store large files efficiently (e.g., videos)?**
Use Object Storage (AWS S3, GCS). Don't store files in the database; store the reference URL/Path in the DB.

**Q133: Explain content-addressable storage.**
A storage mechanism where data is retrieved based on its content (hash) rather than its location. If two files have the same content, they share the same address.

**Q134: Design a secure file sharing system.**
Generate a unique, random URL token. Set expiration. Enforce ACLs. Encrypt file at rest. For sharing, send a signed URL with limited validity.

**Q135: How to implement resumable uploads?**
Split file into chunks. Upload chunks sequentially or in parallel. Track successful chunks on server. If failed, retry only missing chunks (e.g., TUS protocol).

**Q136: How to handle concurrent file uploads?**
Use a load balancer to distribute traffic. Stream data directly to Object Storage (S3 Presigned URLs) to avoid burdening app servers.

**Q137: How do you store and serve user-generated media?**
Upload to S3. Trigger lambda/worker to resize/compress. Store metadata in DB. Serve via CDN.

**Q138: What is chunking and how does it help in file uploads?**
Breaking a large file into smaller pieces. Allows for parallel uploads, retry on failure, and lower memory usage.

**Q139: Explain design of thumbnail generation service.**
User uploads image -> Event to Queue -> Worker picks job -> Generates thumbnails (using ImageMagick) -> Stores in S3 -> Updates DB status.

**Q140: How would you stream a video in real-time?**
Use protocols like HLS (HTTP Live Streaming) or DASH. Break video into small TS chunks (.ts) and a manifest file (.m3u8). Client downloads chunks in sequence.

---

## 🔸 Search & Recommendation Engines

**Q141: How would you design a search autocomplete system?**
Use a Trie (Prefix Tree) data structure. For scale, store top-k frequent terms in nodes. Use Redis or Elasticsearch for fast prefix lookups.

**Q142: Design a spell checker.**
Use Levenshtein Distance or Soundex to find similar words. Contextual spell check uses N-Gram models or ML.

**Q143: How does a recommendation engine work?**
Two main types:
*   **Content-Based:** Recommends items similar to what user liked.
*   **Collaborative Filtering:** Recommends what similar users liked.

**Q144: How to store and index documents for full-text search?**
Use an Inverted Index (maps words to document IDs). Tools: Elasticsearch, Solr, Lucene.

**Q145: What is Elasticsearch and how does it work?**
A distributed, RESTful search and analytics engine based on Lucene. It stores data as JSON documents and builds an inverted index for fast keyword search.

**Q146: Design a trending hashtag system.**
Streaming pipeline (Kafka) -> Count-Min Sketch or Sliding Window Counter (Spark/Flink) -> Store top K sorted hashtags in Redis/DB.

**Q147: How do you handle typos in search queries?**
Fuzzy search. Edit distance algorithms allowing 1-2 character differences.

**Q148: How do you personalize search results?**
Combine relevance score (BM25) with user profile score (previous clicks, location, interests) using Learning to Rank (LTR).

**Q149: How would you design a vector similarity search system?**
Convert items to embeddings (vectors). Use ANN (Approximate Nearest Neighbor) libraries like FAISS or vector databases (Pinecone, Milvus) to find similar vectors.

**Q150: How to optimize search performance in a large dataset?**
Sharding (split index), Caching (common queries), removing stop words, optimizing mappings, and using appropriate hardware.

---

## 🔸 Analytics & Big Data

**Q151: Design a metrics collection system like Prometheus.**
Pull-based model. Central server scrapes metrics endpoints exposed by services. Stores in Time-Series DB. Offers Query Language (PromQL).

**Q152: How do you store and query billions of rows efficiently?**
Use Columnar Storage (Parquet, Cassandra, Redshift). Compress data. Partition by time or key. Pre-aggregate data for dashboards.

**Q153: Design a BI dashboard backend.**
Don't query raw DB. Use an OLAP cube or materialized views. Cache results. Use WebSocket for real-time updates.

**Q154: What is OLAP vs OLTP?**
*   **OLTP (Online Transaction Processing):** Row-oriented, fast writes, short txns (e.g., Banking).
*   **OLAP (Online Analytical Processing):** Column-oriented, complex queries, historical analysis (e.g., Reports).

**Q155: How would you implement an A/B testing platform?**
User requests feature -> Configuration Server checks user hash/bucket -> Assigns Variation A or B -> App logs events -> Analytics compares conversion rates.

**Q156: What is a star schema in data warehousing?**
A modeling technique with a central Fact Table (events) connected to multiple Dimension Tables (attributes like dates, products). fast for queries.

**Q157: How to collect and store clickstream data?**
Frontend JS -> API Gateway -> Kafka -> Data Lake (S3). Batch process later for analysis.

**Q158: Design a time-based log analytics tool.**
ELK Stack (Elasticsearch, Logstash, Kibana) or Splunk. Ingest logs, parse structured fields, index by timestamp.

**Q159: What is the role of data lakes vs warehouses?**
*   **Data Lake:** Stores raw, unstructured data (S3). Cheap, flexible.
*   **Data Warehouse:** Stores structured, processed data (Redshift/Snowflake). Optimized for SQL analysis.

**Q160: How do you build a data pipeline using Spark?**
Read source (HDFS/S3) -> Transformations (Map/Reduce/Filter) -> Write sink (DB/S3). Orchestrate with Airflow.

---

## 🔸 Realtime & Communication Systems

**Q161: Design a real-time chat system.**
Services: Chat Service, Presence Service. Storage: Cassandra/HBase for chat logs (write heavy). Real-time: WebSockets/XMPP. Notifications: Push.

**Q162: Design a multiplayer game server backend.**
State synchronization is key. Use UDP for speed (lossy acceptable). Authoritative Server architecture prevents cheating. Lockstep or Interpolation for smooth movement.

**Q163: How would you implement a typing indicator?**
Ephemeral status. Client sends "typing_start" event -> Server broadcasts to room -> "typing_end" after timeout or send. Don't store in DB; use Redis or memory.

**Q164: Design a presence detection system.**
User connects -> Heartbeat to Redis (set key with TTL). If heartbeat stops, key expires -> User marked offline.

**Q165: How do you handle rate control in WebSocket communication?**
Implement buffering or throttling on the server side. Drop irrelevant or old messages if the client is slow.

**Q166: Design a push notification system.**
App Server -> Notification Service -> Queue -> Workers -> 3rd Party Gateway (FCM/APNS). Handle retries and token management.

**Q167: How would you scale a WebRTC server?**
WebRTC is P2P but needs Signaling Servers (for setup) and TURN servers (relay). Scale Signaling with Load Balancers/Redis PubSub. TURN servers consume bandwidth, scale horizontally.

**Q168: Design a collaborative editing system (like Google Docs).**
Operational Transformation (OT) or CRDTs to handle concurrent edits. WebSocket for real-time sync.

**Q169: What is Operational Transformation (OT)?**
A system of algorithms to allow multiple users to edit the same specific document without conflicts. It transforms operations based on other concurrent operations.

**Q170: Explain CRDT for real-time collaboration.**
Conflict-free Replicated Data Types. Data structures (like PN-Counter, LWW-Element-Set) that automatically resolve conflicts mathematically, allowing independent updates.

---

## 🔸 API Design & Protocols

**Q171: REST vs gRPC – which one and when?**
*   **REST:** JSON over HTTP/1.1. Simple, universal, human-readable. Good for public APIs.
*   **gRPC:** Protobuf over HTTP/2. Binary, strictly typed, fast. Good for internal microservices.

**Q172: How do you version APIs?**
URI Versioning (`/v1/users`), Header Versioning (`Accept-version: v1`), or Parameter Versioning. URI is most common and cache-friendly.

**Q173: How to design an idempotent API?**
Client sends a unique `Idempotency-Key` header. Server checks if key processed. If yes, return saved response. If no, process and save.

**Q174: How to handle breaking changes in APIs?**
Avoid them if possible. If necessary, create a new version (v2). Deprecate v1 with a sunset period. Explain changes clearly in docs.

**Q175: What is HATEOAS in REST?**
Hypermedia As The Engine Of Application State. The API response includes links to valid next actions/states, helping clients navigate the API dynamically.

**Q176: How do you design pagination in APIs?**
*   **Offset-based:** `limit=10&offset=20`. Simple, but slow for large datasets.
*   **Cursor-based:** `limit=10&cursor=xyz`. Faster, stable for realtime data.

**Q177: How do you secure public APIs?**
API Keys (for identification), OAuth2 (for authorization), HTTPS (encryption), Rate Limiting (DDoS protection), and Input Validation.

**Q178: How to implement rate limiting on APIs?**
Token Bucket algorithm on the API Gateway. Identify caller by IP or API Key. Return 429 Too Many Requests.

**Q179: Design a webhook system.**
User registers URL. Event happens -> Queue -> Worker sends HTTP POST to User URL. Retry with backoff if it fails. Verify signature (HMAC) for security.

**Q180: How to validate and document APIs?**
Use OpenAPI (Swagger) Specification. Auto-generate docs. Use schema validation middleware to enforce request structure.

---

## 🔸 Cloud-Native & Kubernetes

**Q181: Design a multi-tenant SaaS application.**
Database per tenant (isolated) or Shared Database with TenantID (cheap). Application logic checks Tenant Context on every request.

**Q182: How do you handle secrets in Kubernetes?**
Use K8s Secrets objects (base64 encoded). Ideally, integrate with external Vault (HashiCorp) and inject as environment variables or mounted volumes.

**Q183: How would you scale stateful apps in Kubernetes?**
Use StatefulSets. They provide stable network IDs and persistent storage. Scaling is harder than stateless; often requires manual operator intervention or specific Operators.

**Q184: What are sidecars in microservices and why use them?**
A helper container running alongside the main app container in the same Pod. Used for logging, proxying (Envoy), or config watching without changing app code.

**Q185: How do you design stateless services?**
Store no session data in the app memory/disk. Store state in external stores (Redis/DB). Allows any instance to handle any request.

**Q186: What is a service mesh?**
An infrastructure layer (like Istio/Linkerd) for handling service-to-service communication. Features: mTLS, observability, traffic splitting, retries.

**Q187: How do you monitor containers at scale?**
Prometheus scrapes metrics. Grafana visualizes. Advisor (cAdvisor) runs on nodes. Logging via Fluentd/DaemonSets.

**Q188: What is the role of Helm in cloud-native deployments?**
Package manager for Kubernetes. Defines "Charts" (templates of YAMLs) to manage complex applications easily.

**Q189: What are Kubernetes Operators?**
Software extensions to K8s that use custom resources to manage applications and their components. Detailed knowledge of how to deploy/manage a specific app (like a DB).

**Q190: How would you design an auto-scaling mechanism?**
Horizontal Pod Autoscaler (HPA) based on CPU/Memory/Custom Metrics (requests/sec). Cluster Autoscaler adds nodes when pods are pending.

---

## 🔸 Miscellaneous System Design Challenges

**Q191: Design a leaderboard system.**
Use Redis Sorted Sets (`ZADD`, `ZRANK`, `ZRANGE`). Fast updates and retrieval of rank. Backup to DB for persistence.

**Q192: Design a notification system.**
Pluggable sender architecture (Email, SMS, Push). Queue based decoupling. User preference store.

**Q193: How to build a feature flag system?**
Checks flag status at runtime. Central config service or DB. Cache flags locally in app with short TTL. Rollout updates incrementally.

**Q194: Design a dynamic pricing engine.**
Inputs: Demand, Supply, Time, User Segment. Rules Engine calculates price. Cache price for short window.

**Q195: How would you build a plagiarism detection system?**
Split text into N-grams or Shingling. Hash chunks (MinHash). Compare Jaccard similarity with database of known docs.

**Q196: Design a subscription billing system.**
State machine (Active, PastDue, Canceled). Recurring jobs (cron) to charge based on period. Payment Gateway integration (Stripe API).

**Q197: Design a Q&A platform (like Stack Overflow).**
Relational DB for Questions/Answers. Redis for Vote counts. Elasticsearch for Text Search. Reputation calculation async job.

**Q198: Design a crowdfunding platform.**
"All-or-Nothing" logic. Hold funds in escrow or authorize card but capture only when goal met. High concurrency checking on funding closing seconds.

**Q199: Design a loyalty and rewards engine.**
Event sourcing system. User Action -> Event -> Rule Engine -> Grant Points -> Ledger.

**Q200: Design a privacy settings system for social media.**
Granular permissions (Public, Friends, Only Me). Check edges in Graph DB. "Can User A see Object B?" query on every access.

---

## 🔹 Fault Tolerance & Resilience

**Q201: How would you design a self-healing system?**
Implement health checks, automatic restarts (via orchestration like K8s), circuit breakers to stop cascading failures, and auto-scaling to replace unhealthy nodes.

**Q202: How do you detect and recover from cascading failures?**
Use Bulkheads (pool isolation), Circuit Breakers, and Timeouts. Detect via spike in error rates/latency. Recover by shedding load (throttling) and restarting components gradually.

**Q203: How to isolate failures in microservices?**
Decouple services using async messaging queues. Ensure failure in one service (e.g., Email service) doesn't block the main flow (e.g., Checkout).

**Q204: What is bulkheading in system design?**
Isolating resources (like thread pools or connection pools) into distinct groups so that failure in one group doesn't exhaust resources for others (like compartments in a ship).

**Q205: How do you simulate failures for testing?**
Use Chaos Engineering tools (Chaos Monkey, Gremlin). Inject network latency, packet loss, or kill processes in staging/prod to verify resilience.

**Q206: What is circuit breaking and when to use it?**
A pattern that "trips" (stops requests) when downstream failures exceed a threshold. Prevents wasting resources on a dead service. Use it for all synchronous external calls.

**Q207: What happens when a database goes down?**
Read replicas can promote to Master (failover). Application should queue writes or return "ReadOnly" mode error.

**Q208: How to design systems for disaster recovery?**
Replicate data to a different geographic region (`us-east` to `us-west`). Automate failover DNS switching. Define RTO (Time objective) and RPO (Point objective).

**Q209: What’s the difference between RTO and RPO?**
*   **RTO (Recovery Time Objective):** Max acceptable downtime magnitude. (How long to come back up?)
*   **RPO (Recovery Point Objective):** Max acceptable data loss. (How much recent data can we lose?)

**Q210: How to handle region-wide cloud outages?**
Multi-region active-active or active-passive deployment. Global Load Balancer routes traffic to the surviving region. Data must be replicated cross-region.

---

## 🔹 Networking & Protocols

**Q211: How does TCP work under high latency?**
TCP performance drops because of the "Bandwidth-Delay Product". It waits for ACKs. optimization involves Window Scaling and Selective Acknowledgments (SACK) or switching to UDP-based protocols like QUIC.

**Q212: What happens during a TCP handshake?**
Three-way handshake:
1.  **SYN:** Client sends synchronization packet.
2.  **SYN-ACK:** Server acknowledges and sends its own SYN.
3.  **ACK:** Client establishes connection.

**Q213: What is long polling vs short polling?**
*   **Short:** Client asks "New data?" repeatedly every X seconds.
*   **Long:** Client asks; server holds the connection open until data is available or timeout. Better for real-time.

**Q214: How would you design a protocol over UDP?**
Add necessary reliability features at the application layer: Sequence numbers for ordering, checksums for integrity, and retransmission logic for critical packets (like Game State).

**Q215: How to reduce latency in cross-country communication?**
Use a CDN to serve static content. Use Edge Computing. Optimize TCP/TLS handshake (TLS 1.3). Route traffic via private backbone networks (like AWS Global Accelerator).

**Q216: What is QUIC and how does it compare to HTTP/2?**
QUIC is built on UDP. It eliminates Head-of-Line blocking (a packet loss doesn't stop other streams). Faster connection setup (0-RTT). HTTP/3 uses QUIC.

**Q217: How do NAT and firewalls affect distributed systems?**
NAT hides internal IPs, complicating P2P connections (requires STUN/TURN). Firewalls block ports, requiring explicit allow-listing for inter-service communication.

**Q218: How do CDNs route traffic efficiently?**
Anycast DNS. The same IP address is announced from multiple locations. The internet routing protocol (BGP) directs the user to the topologically nearest server.

**Q219: How to handle packet loss in real-time systems?**
Use Forward Error Correction (FEC) (sending redundant data to reconstruct lost packets) or conceal errors (interpolation in audio/video) instead of retransmitting (too slow).

**Q220: Explain DNS resolution and its failure points.**
Browser -> OS Cache -> ISP Resolver -> Root -> TLD -> Authoritative NS -> IP. Failure points: Cache poisoning, DDoS on DNS servers, propagation delay.

---

## 🔹 Search, Indexing & Metadata

**Q221: Design a tag-based search system.**
Use Inverted Index: Map `Tag -> List[DocID]`. Intersection of lists gives results for multiple tags (`TagA AND TagB`).

**Q222: How do you implement "did you mean" suggestions?**
Spell check (Levenshtein distance) on query terms against a dictionary of common/valid terms. Suggestions are terms with low edit distance.

**Q223: Design a distributed indexing system.**
Split index into shards (partitions). Map document ID to shard. Scatter-gather: Query all shards, aggregate and rank results.

**Q224: How to handle synonyms and stemming in search?**
Apply Token Filters during indexing. Stemming: `running` -> `run`. Synonyms: `maintain -> expand -> ["support", "keep"]`. store expansions or map at query time.

**Q225: How would you build faceted search?**
During indexing, extract attributes (Color, Size). During query, aggregate counts of these attributes for the result set. Elasticsearch "Aggregations" feature does this efficiently.

**Q226: What is inverted indexing?**
A mapping from content (words/terms) to its location in the database (document IDs). The core structure of search engines.

**Q227: How do you store and search metadata at scale?**
NoSQL (Document store like MongoDB or Wide-column like Cassandra) is good for flexible, large-scale metadata. Index critical fields for search.

**Q228: How to optimize autocomplete suggestions with popularity?**
Store `(prefix, suggestion, score)` triples. use a Trie. Store the max score in each node to quickly traverse the most popular branch.

**Q229: How to merge search indexes from multiple data sources?**
ETL pipeline standardizes data into a common schema -> Ingest into single Search Cluster. OR Federated Search: Query independent indexes and merge results (slower).

**Q230: Design a location-aware search engine.**
Use Geospatial Indexing (QuadTrees, Geohashes, or KD-Trees). Filter points within radius. Rank by distance.

---

## 🔹 AI/ML System Design

**Q231: How do you deploy a machine learning model to production?**
Wrap model in API (Flask/FastAPI/TF Serving). Containerize (Docker). Deploy to K8s. Alternatively, export to ONNX/TensorRT for optimized inference.

**Q232: How to monitor model drift?**
Compare the statistical distribution of training data (baseline) vs live inference data. If deviation (KL Divergence) exceeds threshold, trigger warning/retraining.

**Q233: Design a recommendation system for a video platform.**
Candidate Generation (Two-Tower model retrieval) -> Ranking (Deep learning model with user/video features) -> Re-ranking (filter watched, add diversity).

**Q234: How to serve real-time predictions with low latency?**
Cache common predictions. Use Model Quantization (reduce precision float32->int8). Use specialized hardware (TPUs/GPUs). Keep feature store (Redis) close.

**Q235: How to do A/B testing for ML models?**
Route X% traffic to Model A, Y% to Model B. Measure business metric (CTR, Conversion). Ensure user stickiness (same user gets same model).

**Q236: How do you version and roll back ML models?**
Use Model Registry (MLflow). Tag versions (v1, v2). Deploy via Blue/Green. If v2 metrics drop, easy switch back to v1.

**Q237: Design a fraud detection system using ML.**
Ingest Txn -> Feature Engineering (velocity, location) -> Inference (Random Forest/GBM) -> Score -> Rules Engine (Block/Allow/2FA). Latency < 200ms.

**Q238: How to retrain models automatically with new data?**
Airflow pipeline: Daily job fetches new labels -> Retrains model -> Evaluates against test set -> If better, promote to Registry.

**Q239: How do you manage feature engineering pipelines?**
Use a Feature Store (Feast/Tecton). Define features once. Serve point-in-time correct values for training, and low-latency value lookup for inference.

**Q240: Design an AI-powered personal assistant.**
ASR (Speech to Text) -> NLP/NLU (Intent Recognition) -> Dialog Manager (State Machine) -> Fulfillment (API calls) -> TTS (Text to Speech).

---

## 🔹 Blockchain & Decentralized Systems

**Q241: Design a decentralized identity verification system.**
User holds private key. ID data hashed and signed on chain (DID - Decentralized Identifier). Verifiable Credentials issued by Authorities off-chain but anchored on-chain.

**Q242: How would you store large files on the blockchain?**
You don't. Store file on IPFS/Arweave. Store the hash of the file on the blockchain.

**Q243: What is Merkle Tree and where is it used?**
A binary tree of hashes. Leaf nodes are data blocks. Parent is hash of children. Used in Blockchain (Bitcoin/Ethereum) to verify integrity of transactions efficiently without downloading full block.

**Q244: How would you design a crypto wallet backend?**
Manage keys (HSM/MPC). Index blockchain data (scan blocks to detect user transactions). Broadcast signed transactions.

**Q245: Design a smart contract-based subscription service.**
User approves contract to spend X tokens. Contract has `withdraw()` function callable by Service Provider every month.

**Q246: How do consensus algorithms work in blockchain?**
Nodes agree on the next valid block. Proof of Work (Compute puzzle), Proof of Stake (Economic stake), Proof of Authority (Reputation).

**Q247: Design a blockchain explorer like Etherscan.**
Node (Geth) syncs chain -> ETL extracts blocks/txns -> Relational DB (Postgres) for querying balances/history -> Frontend API.

**Q248: What is proof-of-stake vs proof-of-work?**
*   **PoW:** Security via energy expenditure (miners).
*   **PoS:** Security via capital locking (validators). PoS is greener and faster.

**Q249: How do you build a decentralized messaging app?**
P2P network (Libp2p). Messages encrypted with recipient public key. Stored on decentralized storage (IPFS) or relayed via gossip.

**Q250: How do you ensure integrity in a blockchain network?**
Chained hashing (header contains hash of previous block). Public verification (anyone can check POw/PoS). Longest chain rule.

---

## 🔹 IoT & Edge Computing

**Q251: Design a smart home system.**
Hub (Gateway) connects devices (Zigbee/WiFi). Hub syncs state to Cloud (MQTT Shadow). App talks to Cloud (or Hub locally).

**Q252: How would you handle intermittent connectivity in IoT devices?**
Store and Forward. Device stores messages locally in buffer. Pushes to cloud when connection restores. Use lightweight protocol (MQTT).

**Q253: How do you securely update firmware over the air?**
Cloud signs firmware image. Device downloads, verifies signature (Public Key), flashes to secondary slot, reboots. Fallback if boot fails.

**Q254: How do you process data at the edge?**
Run lightweight compute (AWS Greengrass, K3s) on gateway. Filter/Aggregate data (e.g., avg temp) before sending only anomalies/summary to cloud.

**Q255: Design a low-latency data pipeline for a smart city.**
Sensors -> 5G Edge Server (Process critical events like accidents) -> Cloud (Archival).

**Q256: What is fog computing?**
A decentralized computing structure located between the cloud and devices (the "fog" near the ground). Reduces latency and bandwidth.

**Q257: How do you handle device authentication at scale?**
Mutual TLS (mTLS). Each device burned with a unique X.509 certificate during manufacturing.

**Q258: How would you reduce network usage in edge computing?**
Compression, Protocol optimization (Protobuf/MQTT), and limit transmission frequency (Change-of-Value reporting only).

**Q259: Design a fleet tracking system for delivery vehicles.**
GPS module -> MQTT -> Broker -> Stream Processor (Updates Loc) -> DB (Current Loc) + TimeSeries DB (History).

**Q260: How do you prevent sensor spoofing?**
Cryptographic signing of data packets at the hardware level (Secure Element / TPM). Anomaly detection on the server side.

---

## 🔹 Mobile & Offline Systems

**Q261: How would you build an app that works offline-first?**
Local DB (SQLite/Realm) as source of truth for UI. Sync Engine enables background sync with API when online.

**Q262: How to sync data efficiently between mobile and server?**
Delta Sync. Client sends "Last Updated Timestamp". Server returns only records changed since then.

**Q263: How would you implement conflict resolution in sync?**
Last-Write-Wins (Timestamp) or Server-Wins. For complex logic, keep both versions and ask user to resolve.

**Q264: How do you compress data for slow networks?**
Gzip/Brotli for HTTP. WebP for images. Protobuf for payload.

**Q265: Design a mobile wallet system.**
Security paramount. Store keys in Secure Enclave/KeyStore. Biometrics for access. transactions signed locally, broadcast to backend.

**Q266: How do push notifications work at scale?**
Backend -> Queue -> Notification Service -> Platform Gateway (APNs/FCM) -> Device. Maintain mapping of UserId -> DeviceToken.

**Q267: Design a live location-sharing feature.**
WebSocket/MQTT connection. Send lat/long every X seconds. Server broadcasts to subscribers (Designated friends).

**Q268: How to optimize battery usage in mobile applications?**
Batch network requests. Use system schedulers (JobScheduler/WorkManager) instead of wake locks. Reduce polling.

**Q269: How to handle mobile version compatibility?**
API Versioning. Feature Flags (disable new features on old apps). Force Update screen for breaking changes.

**Q270: How would you secure sensitive data in mobile storage?**
Encrypt database (SQLCipher). Never store API secrets in code. Use OS-provided secure storage (Keychain).

---

## 🔹 Observability & Reliability

**Q271: Design a log correlation engine.**
Inject `Trace-ID` at ingress. Propagate via HTTP headers across services. Logger includes `Trace-ID`. Engine groups logs by ID.

**Q272: What is distributed tracing? How would you implement it?**
Tracking a request flow across microservices. Tools: OpenTelemetry, Jaeger. Instruments code to send spans (start/end time) to collector.

**Q273: How do you detect slow queries?**
Database Slow Query Log. APM tools (New Relic/Datadog) instrument SQL driver. Threshold alerting (> 500ms).

**Q274: Design a custom alerting system.**
Metrics Aggregator -> Eval Rule (CPU > 80% for 5m) -> Alert Manager -> Router -> Integrations (Slack/PagerDuty). Deduplicate alerts.

**Q275: How to create a health dashboard for microservices?**
Central dashboard (Grafana). Visualization of uptime, error rate, latency (p95/p99) for each service. Traffic light System (Green/Red).

**Q276: What is SLO, SLA, and SLI?**
*   **SLI (Indicator):** The metric (Availability).
*   **SLO (Objective):** The goal (99.9%).
*   **SLA (Agreement):** The contract with customer (Penalty if < 99.9%).

**Q277: How to track business metrics from logs?**
Log structured events (`event: "purchase", amount: 50`). Log processor extracts fields -> increments counters in Prometheus/StatsD.

**Q278: Design a queryable log storage system.**
Inverted Index on log tokens (Elasticsearch/Loki). Partition by time (Indices per day). Compress old logs.

**Q279: What metrics would you monitor for a payment system?**
Transaction Rate, Success/Failure Ratio, Latency, Cart Abandonment, Payment Gateway Errors.

**Q280: How do you handle noisy alerts?**
Tune thresholds. Implement "Hysteresis" (must be bad for X mins). Group related alerts. Use AI Ops to suppress anomaly noise.

---

## 🔹 Product-specific Designs

**Q281: Design a digital signature service.**
Upload Doc -> Hashing -> Sign Hash with User Private Key -> Timestamp -> Embed Signature. Validation: Decrypt hash with Public Key.

**Q282: How to build a collaborative calendar system?**
Data Model: Events (Start, End, Users). Conflict detection (Overlapping intervals). Notifications. CalDAV protocol support.

**Q283: Design a voting/polling platform with live results.**
High write ingestion (Votes). Message Queue (Kafka) to buffer. Stream processing to aggregate counts. WebSocket to push updates to clients.

**Q284: Design a document approval system.****
Workflow Engine (State Machine). Draft -> Pending -> Approved/Rejected. Audit log of transitions. Email notifications.

**Q285: How to design an API monetization platform?**
API Gateway tracks usage (calls count). Billing Service aggregates usage -> Calculates cost -> Charges customer via Stripe. Quota enforcement.

**Q286: Design a cloud cost monitoring tool.**
Ingest billing reports (AWS CUR). Parse large CSVs. Aggregate by Tag/Service. Store in OLAP DB. Dashboard visualization.

**Q287: Build a digital content watermarking system.**
On upload/download: Transcoder overlays UUID/UserHash invisibly or visibly onto image/video frames.

**Q288: Design a stock price alert system.**
Ingest price stream. Rule Engine matches price against millions of user triggers (`IF AAPL > 150`). Queue notifications.

**Q289: Design a plagiarism checker backend.**
Rolling Hash (Rabin-Karp) to fingerprint text. Store fingerprints in DB. Compare new doc fingerprints against DB (Jaccard Index).

**Q290: How would you build an auction system?**
Real-time constraints. WebSocket for bid updates. Redis Atomic Increments or Lua script to ensure bid validity (New > Current). Timer services.

---

## 🔹 Privacy, Compliance & Governance

**Q291: How do you handle GDPR data deletion?**
"Right to be forgotten". propagate "Delete User" event. All services delete user data. For backups: crypto-shredding (delete the encryption key for that user's data).

**Q292: How to log access to sensitive data?**
Audit Logging sidecar/middleware. Logs: Who, What, When. Store in tamper-evident storage (WORM - Write Once Read Many).

**Q293: What is data masking and where to apply it?**
Obfuscating data (e.g., `****-****-1234`). Apply at Presentation Layer (UI), Logs, and Non-Prod DBs.

**Q294: How do you design audit logs?**
Immutable, Append-only. structured info. High durability. Secure access.

**Q295: How do you implement RBAC (Role-Based Access Control)?**
Users have Roles. Roles have Permissions. Middleware checks: `User.Role.Permissions.contains("edit_post")`.

**Q296: How to encrypt data at rest and in transit?**
*   **Rest:** AES-256. Database encryption (TDE) or Disk encryption. Key Management Service (KMS).
*   **Transit:** TLS 1.2/1.3 (HTTPS).

**Q297: Design a user consent management system.**
Store consents `(User, ConsentType, Version, Timestamp, Granted)`. Versioning is key (Privacy Policy v1 vs v2).

**Q298: How to implement “Right to be forgotten”?**
(Same as Q291). Centralized deletion orchestration.

**Q299: How do you classify sensitive vs public data?**
Data Catalog/Governance tool. Tag schemas (PII, PCI, Public). Enforce policies based on tags.

**Q300: How to enforce data residency rules in cloud apps?**
Tag tenants with "Region". App logic ensures data for "EU" tenant is only written to "EU" DB/Bucket. Error if attempting cross-region write.

---

## ⚙️ Modern Architecture Patterns

**Q301: What is an event-sourcing architecture? Where would you use it?**
Persisting the state of a business entity as a sequence of state-changing events (e.g., "MoneDeposited", "MoneyWithdrawn") rather than just the current state. Useful for banking, audit logs, and complex domains allowing "time travel" debugging.

**Q302: How would you design a CQRS-based system?**
Command Query Responsibility Segregation. Split the model into two: Write Model (Command) optimized for validation/updates, and Read Model (Query) optimized for fast retrieval (e.g., Views). Sync via events.

**Q303: How does hexagonal (ports and adapters) architecture work?**
Decouples core business logic (center) from external concerns (UI, DB, API) using "Ports" (interfaces) and "Adapters" (implementations). Allows swapping DB or UI without changing core logic.

**Q304: What is a service mesh and when is it useful?**
(Repeated concept, depth added). A dedicated infrastructure layer for handling service-to-service communication. Useful in complex microservice landscapes for Observability, Traffic Control (Canary), and Security (mTLS) offloading.

**Q305: Compare monorepo vs polyrepo in large-scale systems.**
*   **Monorepo:** Single repo for all code. atomic commits, easy code sharing. tooling complexity.
*   **Polyrepo:** Repo per service. Strong boundaries, independent CI. Dependency hell.

**Q306: How do you design for graceful degradation?**
Ensure core functionality works even if auxiliary features fail. E.g., if "Recommendations" fail, show "Popular items" or an empty list, but don't crash the homepage.

**Q307: What are design considerations for zero-downtime deployments?**
Backward compatibility (DB schema, API). Rolling updates. Health checks. Stateless application servers.

**Q308: How to implement blue-green deployments?**
(Repeated concept). Provision new environment (Green), deploy v2. Switch router/LB to Green. Validate. If bad, switch back to Blue.

**Q309: What is canary deployment and when to use it?**
(Repeated concept). Roll out to 1-5% of traffic. Monitor metrics (Error rate). Gradient increase. Use for high-risk changes.

**Q310: How would you design a system for high observability?**
Standardized logging (JSON). Distributed tracing coverage. Granular metrics (Red methodology). Centralized dashboard. Correlation ID propagation.

---

## 🧠 Low-Level Design Scenarios

**Q311: Design a URL parsing library.**
State machine approach. identifying properties (Protocol, Host, Port, Path, Query). Handle edge cases (relative URLs, encoding). API: `Parser.parse(string) -> URLObject`.

**Q312: Design a rate limiter.**
Class `RateLimiter`. Methods: `allowRequest(clientId)`. Logic: Bucket algorithm. Store timestamps/tokens in synchronized map or Redis.

**Q313: Build a cron job scheduler.**
Priority Queue (Min-Heap) storing `(ExecutionTime, Job)`. Thread sleeps until top job time. Execute -> Calculate next time -> Re-insert.

**Q314: Design a thread-safe in-memory cache.**
`ConcurrentHashMap` for storage. `DoublyLinkedList` for LRU eviction. Locks (Read/Write) for consistency during updates.

**Q315: Implement a retry mechanism with exponential backoff.**
Loop with `Thread.sleep(base * 2^attempt)`. Handle `MaxRetries` and transient vs terminal errors. Jitter (randomness) preventing thundering herd.

**Q316: Design a key-value store from scratch.**
In-Memory Hash Map. Persistence: Write Ahead Log (WAL) + Snapshots. Indexing: Hash Index or LSM Tree. Network interface (GET/PUT).

**Q317: Implement a feature toggle mechanism.**
Configuration Provider (JSON/DB/Env). `FeatureManager.isEnabled("feature_key", context)`. Strategy pattern for activation rules (UserID %, Beta Group).

**Q318: Design a library to handle pagination in APIs.**
Classes representing `PageRequest` (limit, offset/cursor) and `PageResult` (items, total, nextLink). Abstraction over DB queries.

**Q319: Create an in-memory log aggregator.**
Buffer (Ring Buffer/Queue). Flush policy (Time-based or Size-based). Sink interface (File, Console, Network). Non-blocking append.

**Q320: Design a UUID generation service.**
Snowflake algorithm (Twitter). 64-bit integar: Timestamp + Machine ID + Sequence Number. K-ordered (sortable by time), unique distributed.

---

## 🏢 Enterprise Systems

**Q321: Design an SSO (Single Sign-On) solution.**
Central Identity Provider (IdP) (e.g., Keycloak). Service Providers (SP) redirect to IdP. IdP validates creds, sets Global Session Cookie, returns Auth Code/Token to SP.

**Q322: How would you build a document management system?**
Storage: S3. Hierarchy: Relational DB (Folders/Files). Indexing: Elasticsearch (Content). Permissions: ACLs. Features: Versioning, Search, Preview.

**Q323: Design a workflow engine.**
Define processes (BPMN/DAG). Engine interprets definitions. State persistence (DB). Task dispatching (Queues). History logging.

**Q324: Design a multi-department HRMS backend.**
Employee Entity (Core). Modules: Payroll, Leave, Performance. Multi-tenancy (Department isolation). RBAC. Audit logging.

**Q325: Build a financial transaction ledger.**
Double-Entry Bookkeeping. Immutable Journal (Debits/Credits). Sum of Debits must equal Credits. ACID transactions mandatory.

**Q326: How would you handle multi-currency accounting?**
Store Amount + CurrencyCode. Exchange Rate Service (historical rates). Reporting in Base Currency. Handle precision (BigDecimal).

**Q327: How would you enforce access control in enterprise apps?**
RBAC (Role Based). Hierarchy (Manager sees Direct Reports). Middleware interceptors. Centralized Policy Decision Point (OPA - Open Policy Agent).

**Q328: Design a time tracking and approval system.**
Timesheet handling. Submission -> Notification to Manager -> Approval Action -> State Change. Reminder Jobs.

**Q329: How would you implement business process orchestration?**
Use an Orchestrator (Camunda/Temporal). Code defines workflows. State is persisted. Retries and Sagas handled by the engine.

**Q330: How would you version business rules?**
Store rules in DB/Git with VersionID. "Effective Date" ranges. App loads active rules for current date.

---

## ⚖️ Trade-offs & Decision Making

**Q331: When do you denormalize data?**
When Read performance is critical and Joins are too slow. When data is read frequently but updated rarely.

**Q332: When to choose SQL vs NoSQL?**
*   **SQL:** Complex queries (Joins), Strict Schema, ACID.
*   **NoSQL:** High Throughput, Dynamic Schema, Hierarchical/Graph data, Scalability > Consistency.

**Q333: How do you choose between vertical and horizontal scaling?**
*   **Vertical:** Start here. Simple, no code changes. Costly at high end.
*   **Horizontal:** When hitting vertical limits or need HA. Requires stateless app/sharding.

**Q334: When to use eventual consistency?**
High Availability requirements. Geo-distributed systems. Non-critical data (Social feeds, Likes).

**Q335: When is strong consistency a must?**
Financial transactions. Inventory counts. Authentication/Access Control.

**Q336: When would you prefer gRPC over REST?**
Internal Microservices. High performance (low latency/bandwidth). Polyglot environments. Streaming needs.

**Q337: Should you build or buy a third-party tool?**
*   **Buy:** Commodity, non-core competency (Email, Auth, Billing).
*   **Build:** Core business differentiator, unique requirements, too expensive at scale.

**Q338: When to go for cloud-native vs self-hosted?**
*   **Cloud-Native:** Speed, Managed Services, OpEx model.
*   **Self-Hosted:** Compliance, Latency (On-prem), massive scale cost optimization (Dropbox scenario).

**Q339: How do you decide database sharding strategy?**
Based on query patterns.
*   **Range:** Good for range queries, bad for hotspots.
*   **Hash:** Good distribution, no range queries.
*   **Directory:** Flexible, lookup overhead.

**Q340: How to decide between pub/sub and queue?**
*   **Queue:** Work distribution (Load balancing). One worker processes msg.
*   **Pub/Sub:** Broadcast state changes. Multiple systems react to same event.

---

## 🧪 Testing, QA & Simulation

**Q341: How do you test distributed systems?**
Unit Tests, Integration Tests (Docker Compose), Contract Consumer Tests (Pact), End-to-End Tests, Chaos Testing.

**Q342: How to simulate network partitions?**
Tools like `iptables` (drop packets), `tc` (add delay), or proxy middleware (Toxiproxy) to cut connections between specific containers.

**Q343: How would you test failover in production?**
Game Days. Scheduled exercises where you intentionally kill a primary DB/Node and verify if the system recovers within RTO.

**Q344: How to design test cases for a messaging system?**
Test ordering, duplicate handling, poison pill (bad format) processing, DLQ routing, retry backoff logic.

**Q345: How do you mock time-sensitive features?**
Abstract time via a `Clock` interface. In tests, inject a `FakeClock` that you can advance manually.

**Q346: How would you test eventual consistency?**
Polling assertion. "Perform Action -> Assert Result (Wait/Retry up to X seconds)".

**Q347: Design a chaos testing module.**
Agent running on nodes. Configurable attacks (CPU burn, Kill Process). Scheduler/Randomizer. Safety stop button.

**Q348: How to test an analytics system for accuracy?**
Send known synthetic events. Run pipeline. Compare Output Query Result with Expected Aggregation.

**Q349: What is the role of synthetic monitoring?**
Scripted bots simulating user journeys (Login -> Checkout) running periodically from global locations to detect outages/performance issues.

**Q350: How do you simulate high concurrency?**
Load testing tools (JMeter, K6, Locust). Distributed load generators. Ramp up Virtual Users to stress limits.

---

## 🛡️ Security-Centric Designs

**Q351: Design a 2FA system.**
User + Password (1st factor) -> Generate TOP/HOTP Code -> Send via SMS/Email/App -> Verify Code -> Issue Token. Rate limit verification.

**Q352: How to securely store access tokens?**
*   **Browser:** HttpOnly Secure Cookies (Prevents XSS).
*   **Mobile:** Secure Storage (Keychain/Keystore). Avoid LocalStorage.

**Q353: What’s the difference between OAuth2 and OpenID Connect?**
*   **OAuth2:** Authorization (Access). "Keys to the car".
*   **OIDC:** Authentication (Identity). "Badge showing who you are". (Built on top of OAuth2).

**Q354: Design a secrets management service.**
(Like Vault). Encrypt secrets at rest. Dynamic secrets (generated on demand). Lease/Revocation. Access via specific Policies.

**Q355: How would you design an audit trail for admin actions?**
Interceptor on Admin API. Log `(Who, Action, Resource, Diff, Timestamp)`. Write to append-only storage.

**Q356: Design a DDoS protection layer.**
Edge filtering (Cloudflare/AWS Shield). Rate Limiting. IP Reputation blocking. Syn Cookies. Auto-scaling to absorb traffic.

**Q357: How do you implement permission inheritance?**
Graph approach. `Admin > Manager > User`. Check path in graph. Or flatten effective permissions at assignment time.

**Q358: How to detect and prevent replay attacks?**
Timestamp + Nonce (valid only once) in request signature. Server tracks seen nonces within valid time window.

**Q359: How would you secure WebSockets?**
WSS (TLS). Authenticate on connect (pass Token). Authorize subscriptions (Can User A subscribe to Channel B?).

**Q360: How to design secure file uploads?**
Validate file type (Magic bytes, not extension). Virus scan. Store outside web root. Randomize filename. Serve via CDN/Presigned URL with `Content-Disposition`.

---

## 🧩 Component and Data Modeling

**Q361: How would you model user roles in a large SaaS platform?**
`Users` N:N `Roles`. `Roles` N:N `Permissions` (or 1:N). `Tenant` context.

**Q362: Model a product inventory with variants and warehouses.**
`Product` (T-Shirt). `Variant` (Red-L). `Warehouse`. `InventoryItem` (WarehouseID, VariantID, Qty).

**Q363: How to model time-based entitlements?**
`Entitlement` (UserID, Resource, ValidFrom, ValidTo). Query: `Select ... where Now() BETWEEN ValidFrom AND ValidTo`.

**Q364: Model a data-sharing agreement between companies.**
`Organization`, `Agreement` (SourceOrg, TargetOrg, DataSope, Expiry). Access logic checks valid Agreement exists.

**Q365: Design a dynamic pricing model.**
`BasePrice`. `Adjustments` (Season, DemandFactor, UserDiscount). `FinalPrice = Base * factors`. History table for auditing.

**Q366: Model a messaging inbox with threads and participants.**
`Thread` (ID). `Participant` (ThreadID, UserID, LastReadMessageID). `Message` (ThreadID, SenderID, Content).

**Q367: Design a schema for customer support tickets.**
`Ticket` (ID, Status, Assignee). `Comment` (TicketID, Author, Text). `History` (TicketID, FieldChanged, Old, New).

**Q368: Model a booking system with cancellation windows.**
`Booking` (ID, Time). `CancellationPolicy` (ResourceID, HoursBefore, Refund%). Logic checks `Now < BookingTime - Hours` to calculate refund.

**Q369: Design a permission hierarchy tree.**
`Permission` (ID, ParentID). Closure Table or Materialized Path (`/root/admin/edit`) for efficient subtree queries.

**Q370: How would you design an entity history tracker?**
`EntityTable` (Current). `EntityHistoryTable` (all fields + Version + ModifiedDate). On Update: Copy current to History, Update Current.

---

## 🔄 API & Integration Design

**Q371: How would you build an API gateway?**
Reverse Proxy (Nginx/Envoy). Modules for Authentication, Rate Limiting, Routing, Request Aggregation/Transformation.

**Q372: How to design a bulk import/export API?**
Async Pattern. POST `/jobs/import` -> returns `JobID`. Client polls `/jobs/{id}`. Server updates status/result URL.

**Q373: How do you handle partial failures in batch APIs?**
`207 Multi-Status`. Response body contains status for each item: `[{id:1, status: 200}, {id:2, status: 500}]`.

**Q374: How do you expose webhooks securely?**
Sign payloads (HMAC). Support HTTPS only. Retry logic. IP whitelisting (if possible).

**Q375: How would you make APIs backward-compatible?**
Additive changes only. Never rename/remove fields. Use deprecation warnings. Versioning for structural changes.

**Q376: Design an API discovery mechanism.**
Central Portal (Swagger UI). Developer registration. Service Catalogue.

**Q377: How would you design async APIs using polling?**
(Similar to Q372). Return `202 Accepted` + `Location: /status/123`. Client GETs status until `200 OK`.

**Q378: How to build an OAuth2-based authorization flow?**
Implement Authorization Server (Endpoints: /authorize, /token, /introspect). Resource Owner (User) grants consent. Client gets Token.

**Q379: Design a billing API integration layer.**
Adapter pattern. Wrapper around Stripe/PayPal. Internal definition of "Charge". Adapter translates to provider specific API.

**Q380: How to prevent API abuse from internal clients?**
Even internal clients need ID/Secret. Quotas per client. Service-to-Service auth (mTLS/JWT).

---

## 🌐 Real-World Product Backends

**Q381: Design the backend of a podcast platform.**
RSS Feed ingestion (Parser). Audio Storage (S3). Search (Metadata). Analytics (Download/Listen tracking).

**Q382: Design a platform like Duolingo (adaptive learning).**
Content Graph (Skills dependency). User State (Progress). Spaced Repetition Algorithm (predict forgetfulness).

**Q383: Build the backend of a flash sales system (like limited-time discounts).**
High Concurrency inventory check. Redis Lua Script for atomic decrement `DECR inventory`. Queue for successful "claims" -> Order Processing.

**Q384: Design a loyalty program with points and tiers.**
Rules Engine (Spend $ -> Points). Async Calculator. Tier Evaluation Job. Expiry Job.

**Q385: How would you build a serverless blog CMS?**
Lambda for logic (GetPost, CreatePost). DynamoDB for storage. S3 for static site generation (trigger build on update).

**Q386: Design an online judge system like Leetcode.**
Submission -> Queue -> Sandbox (Docker/Firecracker) -> Execute Code -> Compare Output -> Return Result. Security (Process isolation, Timeouts) is key.

**Q387: Design the backend of a QR code-based payment system.**
QR contains MerchantID/Amount. Payer scans -> App calls PaymentAPI -> Transfer Money -> Notify Merchant via WebSocket.

**Q388: Build a birthday/anniversary reminder service.**
DB stores Dates. Daily Job checks `Day/Month == Today`. Sends notification. timezone handling.

**Q389: Design an ad delivery engine.**
Request -> Auction Service (Real Time Bidding) -> Select Winner -> Return Ad Markup. Maximize eCPM. Latency budget < 100ms.

**Q390: Build a real-time sports score platform.**
Data Feed Provider (Push/Pull) -> parsing -> Update Cache (Redis) -> Pub/Sub -> WebSocket Gateway -> Clients.

---

## ☁️ DevOps & Deployment Systems

**Q391: Design an internal CI/CD system.**
Repo Webhook -> Build Server (Jenkins/Tekton). Steps: Checkout, Test, Lint, Build Artifact (Docker), Push Registry, Deploy Manifest.

**Q392: How would you build infrastructure provisioning using Terraform?**
Define State (S3 backend). Modules (VPC, Cluster, DB). `terraform plan` (review changes). `terraform apply`.

**Q393: Design a centralized logging solution for 500 microservices.**
Sidecar (Fluentd) in every pod -> Aggregator layer -> Kafka -> Elasticsearch. Retention policies (Hot/Warm/Cold).

**Q394: How to auto-scale based on CPU and memory metrics?**
Metrics Server collects usage. HPA (Horizontal Pod Autoscaler) query metrics. Adjust Replicas.

**Q395: How to secure deployments using GitOps?**
Git is Source of Truth. ArgoCD works inside cluster, pulls from Git. No CI server needs direct cluster access (Security+).

**Q396: How do you perform rolling upgrades for Kubernetes services?**
Strategy: `RollingUpdate`. `maxUnavailable: 25%`, `maxSurge: 25%`. K8s replaces pods incrementally.

**Q397: How would you design a secrets rotation system?**
Secrets Manager (Vault). Agent generates new DB credential. Updates Vault. Restarts/Notifies App to reload config.

**Q398: How to isolate noisy containers in a shared cluster?**
Resource Limits/Requests (CPU, RAM). QoS classes (Guaranteed vs BestEffort). Dedicated Nodes (Taints/Tolerations).

**Q399: Design a system for monitoring container resource limits.****
cAdvisor -> Prometheus. Alert if `container_memory_usage` > 90% of `limit`. Detect OOMKilled events (KubeStateMetrics).

**Q400: Build a dashboard to track deployment status across environments.**
API to query GitOps status / CI pipeline status. Visualize Version Number per Service per Env (Dev: v1.2, Prod: v1.1).

---

## 🌀 Advanced Distributed Systems

**Q401: How would you handle leader election in distributed systems?**
Use algorithms like Raft, Paxos, or Bully Algorithm. Alternatively, use a coordination service like ZooKeeper or Etcd where nodes compete to create an ephemeral node; the winner becomes leader.

**Q402: Design a gossip-based messaging protocol.**
Nodes periodically pick random peers to share state improvements. Epidemic spreading ensures all nodes eventually receive the update. Optimized with "Push-Pull" strategies.

**Q403: How would you achieve quorum-based consensus?**
Define Read (R) and Write (W) quorums such that R + W > N (Total nodes). This overlap ensures that every read sees at least one write from the latest update.

**Q404: Explain vector clocks and how you’d use them.**
(Repeated concept). Use to track causal history of data updates in distributed stores (like Riak/Dynamo). Compare vectors to determine if version A happened before B, or if they are concurrent (conflict).

**Q405: How to detect split-brain scenarios?**
Monitor heartbeats. If partition occurs, use a Mediator or Quorum. If a node cannot contact the quorum, it fences itself (stops processing writes) to prevent data corruption.

**Q406: How to handle write skew in distributed databases?**
Write skew happens in Snapshot Isolation when two transactions read overlapping data but update disjoint data. Prevent using Serializable Isolation or materializing conflict (creating a dummy row to lock).

**Q407: Design a system using Raft consensus algorithm.**
Cluster of 3-5 nodes. One Leader, others Followers. Client writes to Leader. Leader replicates Log Entry. Once majority Ack, Leader commits and responds.

**Q408: How do you manage schema evolution in distributed systems?**
Use schema registries (Avro). Ensure forward/backward compatibility (e.g., add optional fields). Rolling updates of services where new consumers can read old data and viz-versa.

**Q409: How would you optimize for low tail latency?**
Hedge Requests (send to 2 replicas, use first response). Good partitioning. Isolate noisy neighbors. Keep data in memory.

**Q410: Explain read-repair mechanism in eventual consistency.**
When a client reads data from multiple nodes and detects inconsistencies (stale versions), the updated value is written back to the stale nodes to converge them.

---

## ⏱️ Real-Time & Event-Driven Systems

**Q411: Design a real-time gaming leaderboard system.**
(Repeated). Use Redis Sorted Sets. Score update = `ZINCRBY`. Rank lookup = `ZREVRANK`. Partition by score range or user ID if massive.

**Q412: How would you build a high-frequency trading system?**
Co-location (servers near exchange). FPGA/ASIC for speed. Kernel bypass networking (Solarflare). Limit-less lock-free data structures. C++.

**Q413: How would you ensure ordering in an event stream?**
Partition data by key (e.g., UserId). Ensure all events for that key go to the same partition (Kafka). Consumer processes partition sequentially.

**Q414: What’s the difference between event sourcing and event streaming?**
*   **Sourcing:** System state is derived from replaying events.
*   **Streaming:** Moving data continuously between systems for processing (ETL, Analytics).

**Q415: Design a telemetry system for autonomous vehicles.**
Vehicle processes raw data (LiDAR) at Edge. Sends summary/anomalies via 5G/LTE to Cloud (MQTT/HTTP). Cloud stores in Time-Series DB.

**Q416: How would you handle late-arriving events?**
Use Event Time (when it happened), not Processing Time. Watermarking in stream processors (Flink) allows waiting for a period before closing the window.

**Q417: How do you design time-window-based analytics?**
Tumbling Window (fixed, non-overlapping) or Sliding Window (overlapping). Stream processors maintain state for the window duration.

**Q418: How to handle out-of-order events?**
Buffer events. Reorder based on timestamp. Discard events that are "too late" (past watermark).

**Q419: Build a real-time bidding system for ads.**
(Repeated). Bidder Service must respond < 100ms. Timeout enforcement. In-memory profile lookup. Asynchronous logging.

**Q420: How do you debounce vs throttle events at scale?**
*   **Debounce:** Group bursts, emit only after silence.
*   **Throttle:** Emit at most once every X seconds. Implement via Redis keys with expiry.

---

## 🌍 Cross-Region and Global Systems

**Q421: Design a global distributed file system.**
Namespace federation. Caching at edges. Metadata services in every region. Replication for durability. Strong or Eventual consistency based on file type.

**Q422: How to replicate user data across continents?**
Multi-Master (complex conflict resolution) or Master-Slave with Read Replicas globally. Geo-partitioning (User pinned to home region).

**Q423: How to design a geo-aware DNS routing system?**
Use Latency-Based Routing (Amazon Route53). DNS returns IP of region with lowest latency to user.

**Q424: What’s the role of Anycast in global systems?**
Single IP address announced from multiple locations. Network routes request to nearest topological ingress point.

**Q425: How to maintain GDPR compliance in multi-region architecture?**
Data Residency. Store EU users in EU region databases. Prevent replication to US. Logic layer enforces routing.

**Q426: How would you geo-fence user data storage?**
Tag data with location. Storage layer enforces placement policies. "User A (DE)" -> Only write to `db-frankfurt`.

**Q427: How to reduce cross-region data transfer costs?**
Compress data. Batch updates. Replicate only essential data. Use Backbone networks (cloud provider internal) instead of public internet.

**Q428: Design a worldwide leaderboard that updates in near-real-time.**
Local Leaderboards (Region) aggregate scores -> Send Top K to Global Leaderboard Service (Redis) -> Broadcast updates back to regions.

**Q429: How would you architect a global media delivery service?**
Origin Storage (S3). tiered Caching: Regional Edge Caches -> Local Edge Caches (ISP). Intelligent Request Routing.

**Q430: What’s the design of a follow-the-sun support platform?**
Global Ticket workflow. Shift handovers based on time. "US Team signs off, transfers active queue to Asia Team". Shared state DB.

---

## 🔒 Privacy-First & Regulated Systems

**Q431: How would you design a system to log data access without logging the data itself?**
Log `(UserID, AccessedResourceID, Timestamp, Operation)`. Do not include the payload.

**Q432: Design a healthcare record system compliant with HIPAA.**
Encryption everywhere (At rest/transit). Strict Access Logs. RBAC. Physical security (managed by Cloud Provider). BAA (Business Associate Agreement).

**Q433: How to tokenize PII data and still allow searching?**
Deterministic Encryption (always encrypts to same ciphertext) allows exact match search. Or store Token -> ID mapping in secure vault.

**Q434: What’s differential privacy and how would you implement it?**
Adding noise to datasets so individual records cannot be reverse-engineered, while aggregate statistics remain accurate.

**Q435: Design a zero-knowledge authentication system.**
Prover demonstrates knowledge of a secret (password) to Verifier without revealing it. (e.g., SRP - Secure Remote Password protocol).

**Q436: How would you audit access patterns without violating user privacy?**
Anonymize logs. Aggregate metrics. Use PII scanners to ensure no leaks in logs.

**Q437: How to build a KYC/AML-compliant data platform?**
Know Your Customer / Anti-Money Laundering. Verify Identity (Docs). Screen against Watchlists. Monitor Transaction Patterns. Keep Audit Trails for 5-7 years.

**Q438: Design a parental control system with granular rules.**
Policy Engine. Rules: Time Limits, Content Categories. Network filter (DNS/VPN) or Device Agent enforces rules.

**Q439: How do you design audit-only access for sensitive systems?**
"Break Glass" procedure. Admin requests temporary elevate access. All actions recorded. notification sent to Security Team.

**Q440: How to comply with "right to data portability"?**
Build "Export Data" feature. Aggregates all user data (DB, Logs, Files) into a standard format (JSON/ZIP) for download.

---

## 🧭 Hybrid Architectures & Integration

**Q441: Design a cloud-bursting architecture between on-prem and AWS.**
Base load on On-Prem. Monitoring detects threshold. Spin up AWS EC2 instances. Load Balancer routes overflow traffic to AWS. VPN/Direct Connect for secure link.

**Q442: How would you integrate a legacy monolith into a new microservices stack?**
Strangler Fig Pattern. Put Proxy/Gateway in front. Route specific paths (/api/new) to Microservices. Keep (/api/old) to Monolith. Gradually migrate.

**Q443: How do you handle API contracts between polyglot services?**
Schema-first design (Protobuf/OpenAPI). Generate client/server code from schema. Enforce compatibility checks in CI.

**Q444: Design a system to sync SaaS app data across multiple tenants.**
(Ambiguous question, assuming syncing 'master' data to tenants). Publish-Subscribe. Master updates -> Event -> Tenant Listeners update local stores.

**Q445: How to build a unified notification system with email, SMS, and push?**
Abstraction layer: `sendNotification(user, type, content)`. Routing logic selects provider (Twilio, SendGrid, FCM). Template engine formats content.

**Q446: How would you migrate a live database without downtime?**
Dual Write. 1. App writes to Old & New. 2. Backfill historic data. 3. Verify consistency. 4. Switch Reads to New. 5. Remove Old.

**Q447: Design a hybrid cloud-native + edge deployment model.**
Control Plane in Cloud (K8s Master). Worker Nodes at Edge. GitOps syncs manifest to Cloud, Edge pulls updates.

**Q448: Build an adapter layer for integrating 3rd-party APIs.**
Interface `PaymentProvider`. Implementations `StripeAdapter`, `PayPalAdapter`. App depends on Interface. Adapter handles auth, mapping, retries.

**Q449: How to version inter-service contracts in gRPC?**
Package names `v1`, `v2`. Run both versions in parallel if breaking change. Client negotiates or upgrades.

**Q450: Design a feature rollout system that integrates web and mobile.**
Central Feature Flag service. Web pools on load. Mobile fetches headers/config on startup. Consistent hashing ensures user sees same variation on both.

---

## 🛒 E-commerce / Fintech Specific

**Q451: Design a flash sale system that avoids overselling.**
Inventory mgmt is key. Redis `DECR`. If result >= 0, success. Async update DB. Queue orders.

**Q452: How would you ensure idempotency in a payment API?**
Mandatory `Idempotency-Key` header. DB constraint on Key. Check before process. Store response.

**Q453: Design a refund processing pipeline.**
Trigger Refund. Update Order Status (Refund_Pending). Call Gateway. If success, Status Refunded, Ledger updated. If fail, Retry/Alert.

**Q454: How to handle high-value transactions securely?**
Multi-approval workflow (Maker-Checker). Fraud checks. Hardware Security Module (HSM) for signing. Real-time alerts.

**Q455: Design an invoicing and tax calculation microservice.**
Input: Items, Location. Call Tax Provider (Avalara/Vertex). Generate PDF. Store. Email User.

**Q456: How do you build a secure shopping cart across devices?**
Store cart in DB (not local storage). Keyed by UserID. Merge Guest Cart (SessionID) with User Cart upon login.

**Q457: Design a pricing engine for a global marketplace.**
Base Price per Product. Currency Conversion. Country Logic (VAT/GST). Promotions application.

**Q458: Build a credit scoring engine based on activity data.**
Ingest signals (Payment history, Usage). Run Scoring Model. Update User Score profile. Use score for credit decisions.

**Q459: How would you design EMI loan repayment tracking?**
Schedule generation (Dates, Amounts). Daily Job checks due dates. Triggers Auto-Debit. Updates Schedule status. Handles failures (Late fees).

**Q460: How to handle currency fluctuation in real-time checkout?**
Fix rate at start of checkout session (guarantee for 15 mins). Hedge the risk or use real-time spot rates with buffer.

---

## 🧱 Database-Specific Questions

**Q461: How would you implement TTL for database records?**
Some DBs support native TTL (DynamoDB, MongoDB, Redis). For SQL, run a background cleanup job `DELETE FROM table WHERE expiry < NOW()`.

**Q462: What is LSM Tree and where is it used?**
Log-Structured Merge-tree. Optimizes for writes (append only). Used in Cassandra, RocksDB, HBase. Data buffered in Memory (MemTable), flushed to Disk (SSTable), compacted later.

**Q463: How do you design a query cache with invalidation?**
Store Query Hash -> Result. Invalidate when tables involved are updated. (Hard to get right, better to cache objects).

**Q464: Design a schema for audit logs with replay capability.**
Store `InitialState`, list of `Events` (Diffs). To replay: Load Initial, Apply Diffs sequentially.

**Q465: Design a write-heavy logging system with efficient read support.**
Use detailed index (Elasticsearch) for search. Use fast append log (Kafka) for ingestion.

**Q466: How would you enforce uniqueness across multiple fields?**
Composite Unique Constraint in SQL (`UNIQUE(col1, col2)`). In NoSQL, derive a unique key `col1_col2` and use it as primary key.

**Q467: When to choose columnar databases?**
For Analytical workloads (OLAP). queries scanning many rows but few columns (e.g., "Average Salary"). Cassandra (wide-column) or Redshift.

**Q468: How to manage schema evolution in NoSQL?**
Application handles schema (Schema-on-read). Code must handle missing fields (defaults) or deprecated fields.

**Q469: Build a DB partitioning strategy based on activity.**
Hot/Cold partitioning. Recent data (Hot) on high-performance DB. Old data (Cold) moved to cheaper storage/archive tables.

**Q470: How to perform online reindexing of large datasets?**
Create new Index. Double-write to Old and New. Backfill data to New. Switch Reads to New. Drop Old.

---

## ⛓️ Concurrency & Parallelism

**Q471: How to design a system with concurrent writers and readers?**
Use MVCC (Multi-Version Concurrency Control) database. Writers create new version; Readers read old version. No blocking.

**Q472: How to handle deadlocks in distributed lock management?**
Lease/Timeouts on locks. Deadlock detection (wait-for graph) is hard distributed. Prevention (ordering resource acquisition).

**Q473: Design a job queue with priority and fairness.**
Multiple queues (High, Med, Low). Consumers poll High first. Or Weighted Round Robin. Starvation prevention (promote old low-priority jobs).

**Q474: How to parallelize file processing safely?**
Split file into chunks. Put chunk info in Queue. Workers process chunks. Coordinator tracks completion. Idempotent processing imperative.

**Q475: Build a system for running background jobs with retry, delay, and cancel.**
(Like Sidekiq/Quartz). DB table for Jobs. Poller thread picks due jobs. Execute. On fail, update `retry_at` with exponential backoff.

**Q476: What are the issues in shared-nothing architecture?**
Data partitioning complexity. Cross-partition transactions are expensive. Rebalancing data when adding nodes.

**Q477: How would you implement mutex in a distributed system?**
(Repeated). Redis `SETNX` (Set if Not Exists) with expiry. Or Zookeeper ephemeral nodes.

**Q478: Design a bulk data processing system using worker pools.**
Producer fills Queue. Fixed pool of Workers consume. Monitor Queue depth to scale pool.

**Q479: How to implement exponential backoff and jitter across clients?**
Client logic: `wait = min(cap, base * 2^attempt)`. `wait_with_jitter = wait / 2 + random(0, wait/2)`. Prevents synchronized retries.

**Q480: Design a checkpointing system in long-running pipelines.**
Periodically save state (offset, partial results) to persistent store. On restart, load last checkpoint and resume.

---

## ⚡ User Experience-Oriented

**Q481: Design a “Save for later” system in a shopping app.**
Move Item from Cart Table to SavedItems Table. (UserID, ProductID, DateAdded). Suggest moving back when available/discounted.

**Q482: How would you build "undo" functionality in UI-backed system?**
Delay action execution (optimistic UI). Or Command Pattern (Action and InverseAction). Store history stack.

**Q483: Design a progress-aware file uploader.**
Client sends chunks. Server acknowledges bytes received. Client calculates % = `acked / total`. Update UI.

**Q484: How to cache partial page loads with API responses?**
GraphQL or Fragment Caching (ESI - Edge Side Includes). Cache specific API calls (e.g., Header, Sidebar) independently.

**Q485: How would you design smart auto-refresh in dashboards?**
Adaptive polling. If user active, poll fast. If idle, poll slow. Or use WebSocket/SSE for push updates only on change.

**Q486: Design a “Recently Used” item system.**
List per user. On Access, remove item if exists, push to front. Trim to Size K. Store in Redis/DB.

**Q487: Build an in-app notification center with read/unread sync.**
Notification Table (ID, User, Content, IsRead). Endpoint `markAsRead(ids)`. WebSocket pushes new count.

**Q488: How to implement real-time typing indicators in chat?**
(Repeated). Ephemeral events. "User X started typing". Client shows UI. Timeout clears UI.

**Q489: Design an infinite scrolling backend.**
Cursor-based pagination. Client requests Next Page. Server returns items + Cursor. Pre-fetch next page on client.

**Q490: Build a clipboard history syncing system across devices.**
Daemon watches clipboard. Encrypts. Pushes to cloud. Other devices pull/push updates. E2E encryption needed.

---

## 🧬 Analytics & Insights Systems

**Q491: Design a funnel tracking system.**
Log events `(User, Step, Time)`. Query: Count users at Step 1. Count users at Step 2 who were in Step 1... Calculate Drop-off %.

**Q492: Build an ad performance dashboard backend.**
Aggregated tables (Pre-calculation). `Sum(Impressions), Sum(Clicks)` grouped by `AdID, Hour`. Fast read queries.

**Q493: Design a cohort analysis pipeline.**
Group users by `Acquisition Date`. Track retention over weeks. Matrix query: `Rows=Weeks since signup, Cols=Cohort`.

**Q494: How to build a delayed-event ingestion pipeline?**
Handle events arriving days late. Update historical aggregates. Idempotency is crucial. Re-processing capability.

**Q495: Design a clickstream data storage system.**
High write volume. Kafka -> S3 (Parquet). Partition by Date. Query with Athena/Presto.

**Q496: Design a time-series aggregation engine.**
Stream processor (Flink). Aggregate raw points into 1m, 1h, 1d buckets. Store in TSDB.

**Q497: How to ensure accuracy in sampled analytics?**
Use consistent sampling (Hash UserID). Scale factors (if sampling 1/10, multiply results by 10). Estimate error margins.

**Q498: How would you process multi-tenant analytics securely?**
Row-level security. Every query must include `WHERE TenantID = ?`. Physically separate tables if high security needed.

**Q499: Build a sentiment analysis dashboard with trending topics.**
Ingest Social Feed. NLP Service scores texts. Aggregate scores by Topic. Rank topics by volume/negativity.

**Q500: How to design an “anomaly detection” system for business metrics?**
Forecast expected value (Holt-Winters/ARIMA). Compare Actual vs Forecast. If deviation > 3 sigma, alert.

---

## 🧠 Machine Learning & AI System Design

**Q501: How would you design a recommendation engine?**
(Repeated concept). Ingest user events. Train Collaborative Filtering (Matrix Factorization) or Deep Learning models. Serve via API (Candidate Generation -> Ranking).

**Q502: Design a real-time fraud detection system using ML.**
(Repeated). Feature extraction (Aggregates). Inference Call. Low latency reqs (< 200ms). Fallback to rules if ML fails/timeout.

**Q503: How to deploy and scale machine learning models in production?**
Model-as-a-Service (Docker container exposing REST/gRPC). Orchestrate with K8s (Seldon Core/KServe). Scale based on GPU utilization/QPS.

**Q504: How do you handle model versioning in a large system?**
Model Registry (MLflow). Associate artifacts with metadata (Metrics, Hyperparams, Commit SHA). Deploy specific version tags.

**Q505: Design an A/B testing platform for ML models.**
Split traffic. Assign User -> Bucket. Bucket A gets Model v1, Bucket B gets Model v2. Log outcomes. Analyze significance with T-Test.

**Q506: Build a system for real-time personalized content.**
User Profile Service (interests). Content Store (metadata). Matcher service scores content against profile. Cache top results.

**Q507: How would you design a feature store?**
Central repository for ML features. Online Store (Redis) for low-latency serving. Offline Store (Parquet/S3) for training. Consistency between them.

**Q508: How do you manage data drift in ML systems?**
Monitor distributions of inputs. If P(X) changes significantly, alert. Retrain model on recent data.

**Q509: Build a self-learning chatbot backend.**
NLP (Intent Classification). Dialog Management (Reinforcement Learning or State Machine). Feedback loop (User says "Wrong answer" -> Label data -> Retrain).

**Q510: How do you log and monitor ML inference performance?**
Log `Input, Prediction, GroundTruth (delayed), Latency`. Monitor Prediction Distribution (to detect bias/drift).

---

## 🌐 Edge Computing & IoT

**Q511: Design a smart traffic light control system.**
Sensors (Camera/Induction loop) at intersection. Local Edge Controller processes video (Computer Vision) to detect cars. Adjusts lights locally. Syncs stats to cloud.

**Q512: How would you build a fleet tracking system using GPS devices?**
(Repeated). Device -> MQTT -> Cloud. Geospatial DB. Geofencing services.

**Q513: Build a system to sync IoT data from millions of sensors.**
Sharded MQTT Brokers. Load Balancer. Streaming pipeline (Kafka) to handle write spikes. Batch write to DB.

**Q514: Design a firmware update system for IoT devices.**
(Repeated). Repo of signed binaries. Device checks for update. Downloads. Verifies. Installs. Reports status. Phased rollout.

**Q515: How would you handle data validation on the edge?**
Lightweight schema checks (JSON Schema) on the device. Discard malformed data to save bandwidth.

**Q516: Design a real-time fire alert system using sensors.**
Sensor detects smoke -> Local Alarm (Urgent) -> Message to Edge Gateway -> Priority Push to Cloud -> Notification Service (SMS/Call).

**Q517: Build an architecture for offline-first mobile apps.**
(Repeated). Local DB (SQLite). Sync Queue. Conflict resolution strategy. background sync Job.

**Q518: How to handle syncing data between edge and cloud?**
Bi-directional sync. Edge pushes telemetry. Cloud pushes configuration/commands. Shadow Device pattern.

**Q519: Design a smart home command center backend.**
User App -> Cloud API -> Device Shadow -> MQTT Command -> Home Gateway -> Device (Lightbulb).

**Q520: Build a local-first video processing system on edge devices.**
Camera -> Local AI Accelerator (Jetson) -> Object Detection -> Send Metadata to Cloud (not full video). Upload video clip only on "Event".

---

## 📎 Collaboration Platforms

**Q521: Design a live collaborative whiteboard.**
Canvas state as list of Shapes/Paths. Broadcast delta updates (`{action: "draw", path: [...]}`) via WebSocket.

**Q522: How would you build real-time Google Docs-style editing?**
OT (Operational Transformation). Server enforces order of operations. Client transforms local op against incoming ops.

**Q523: Design a system for version-controlled document editing.**
Store snapshots (v1, v2) and change logs (deltas). "Save" creates new snapshot. "History" replays deltas.

**Q524: How do you manage permissions in collaborative workspaces?**
ACLs (Access Control Lists) on Document/Folder level. `User A has WRITE on Doc X`. Cache checks for speed.

**Q525: Build a project management app backend (like Trello).**
Board -> Lists -> Cards. Drag & Drop = Update `ListID` and `Position` of Card. Real-time update to other viewers.

**Q526: Design a task assignment and notification engine.**
Assignment Event -> Notification Service. Fanout to Email/In-App. Check user preferences.

**Q527: Build a calendar invite and availability detection system.**
Interval Tree data structure to find free slots. `FindFree(User, Duration)`. Check overlaps.

**Q528: Design a shared annotation system for PDFs/images.**
Store annotations as overlay coordinates `(x, y, text)`. Anchored to document version/page.

**Q529: Design a time-tracking system with team analytics.**
Timer Start/Stop events. Duration calculation. Aggregation by Project/User/Tag.

**Q530: How to ensure consistency in multi-user interactions?**
Optimistic Locking (Version numbers). Check `CurrentVersion == ReadVersion` before update. If separate, reject/merge.

---

## ⚔️ Design Tradeoff Scenarios

**Q531: When would you use peer-to-peer instead of client-server?**
File sharing (BitTorrent), VoIP (WebRTC), Censorship resistance, Reducing server bandwidth costs. Tradeoff: Complexity, Security, Consistency.

**Q532: When would you store derived data vs calculate on the fly?**
*   **Store:** Expensive calculation, frequent reads (Dashboard).
*   **Calc:** Cheap calculation, real-time data needed, rare reads.

**Q533: Compare column-oriented vs row-oriented databases.**
*   **Row:** OLTP. Good for fetching single record (Select *).
*   **Column:** OLAP. Good for aggregates (Sum(Salary)). Better compression.

**Q534: Tradeoffs between batch processing and stream processing.**
*   **Batch:** Good for complex logic, high throughput, reprocessing. High latency.
*   **Stream:** Low latency. Complex error handling/state management.

**Q535: When to choose eventual consistency over strict consistency?**
Where Availability and Partition Tolerance are preferred (CAP). Scalability needs. User experience tolerates stale data (Likes count).

**Q536: When to use webhooks vs polling?**
*   **Webhooks:** Event-driven, low latency, less server load (no empty checks).
*   **Polling:** Simpler client (firewall friendly), control over rate.

**Q537: Tradeoffs between push vs pull messaging.**
*   **Push:** Real-time, server burden (keeping connections).
*   **Pull:** Client control, easier to scale server (stateless), latency.

**Q538: When is an in-memory cache harmful?**
Stale data risk. Memory cost. Complexity of invalidation. Cold start issues.

**Q539: When to replicate vs partition data?**
*   **Replicate:** Read scalability, Availability.
*   **Partition:** Write scalability, Large dataset size. Usually do both.

**Q540: When is premature optimization justified?**
Rarely. Only key architectural decisions (DB choice) that are hard to change later. Avoid micro-optimizations.

---

## 🛠️ Dev Tooling, Developer Experience & Platforms

**Q541: Build a CI/CD pipeline for ML models.**
Trigger on Code/Data change. Train. Eval. Register. Deploy to Staging. Integration Test. Promote to Prod.

**Q542: Design a feature-flag management platform.**
Admin UI to set flags. SDK for clients (`getFlag(user)`). Polling or Streaming updates to SDK.

**Q543: Build a schema registry and validation system.**
Central store for Avro/Protobuf schemas. Producers register schema. Consumers download. Enforce compatibility rules.

**Q544: How would you build a local-first developer playground?**
Docker-based. CLI tool spins up containers (DB, API) locally matching prod versions. Mount code volumes.

**Q545: Build a backend for a real-time log viewer.**
Tail logs. Push via WebSocket to browser. Filter at source/server to save bandwidth.

**Q546: How to design a system that captures runtime metrics?**
Agent/Library in app. Samples CPU/Memory/Reqs. Periodic flush to Metrics Backend (Prometheus).

**Q547: Build an automated changelog generator.**
Parse Git Commits (Conventional Commits: `feat:`, `fix:`). Aggregate by tag. Render Markdown.

**Q548: Design a developer API key management platform.**
Generate secure keys (High entropy). Hash storage. Scope permissions. Usage tracking/Billing. Revocation.

**Q549: How to allow secure plugin support in SaaS?**
Webhooks (Events). IFrames (UI). Sandbox execution (WebAssembly/Functions) for server-side logic.

**Q550: Design a sandbox environment manager for services.**
Ephemeral namespaces in K8s. Spin up dependency graph. TTL auto-cleanup.

---

## 🧮 High-Throughput Data Systems

**Q551: Design a high-frequency sensor data processing pipeline.**
Sensor -> UDP/MQTT -> Kafka (Partitioned) -> Flink (Windowing/Filtering) -> TSDB. compression is vital.

**Q552: How would you ingest and index billions of rows daily?**
Bulk loading. Disable indexes during load. Partitioning. Log-structured storage (LSM).

**Q553: Build a log compaction service.**
Read logs. Dedup. Merge small files into large blocks (Parquet). Upload to Cold Storage (Glacier).

**Q554: Design a document deduplication system.**
MinHash / SimHash. Compare signatures. Group near-duplicates.

**Q555: How to throttle high-throughput API clients?**
Distributed counters (Redis). Token Bucket. Return 429.

**Q556: Design a fast data tagging and labeling system.**
Stream processing. Apply Regex/Rules/ML Classifier. Append tags. Index tags immediately.

**Q557: How to build a stream join service?**
Join Stream A and Stream B. Buffer events in State Store (RocksDB) within window T. Emit Joined event.

**Q558: Design a timestamp alignment engine for time-series.**
Resampling. Interpolation (Linear/Last-Known). Align to grid (every 1 min).

**Q559: Build a multi-source data stitching engine.**
Common identifier (Email/Phone). Graph resolution. Merge attributes with precedence rules.

**Q560: Design a dynamic data warehouse ingestion system.**
Schema inferencing. Auto-create tables. Handle schema drift (add columns). Copy data.

---

## ⚙️ Workflow, Pipelines & Job Systems

**Q561: Design a DAG-based workflow scheduler.**
(Like Airflow). Parse DAG. Determine dependencies. Queue ready tasks. Workers execute. Update state. Handle failure/retries.

**Q562: Build a distributed task execution system with retries.**
Task Queue. Workers. Ack mechanism. If no Ack (crash), task becomes visible again (At-least-once). Dead Letter Queue.

**Q563: Design a system to pause and resume data pipelines.**
Checkpoints. Save offset/state. On resume, load state.

**Q564: How to orchestrate cross-platform pipelines (e.g., ML + ETL)?**
Unified control plane. Agents on platform (Spark cluster, GPU cluster). API to submit jobs/poll status.

**Q565: Build a cron-job metrics dashboard.**
Parse cron logs. Track `Start, End, ExitCode`. Visualize Duration and Failure Rate. Alert on `Missing`.

**Q566: How to version and rollback pipeline logic?**
Git-backed DAG definitions. Deployment updates the scheduler. Rollback = Git Revert.

**Q567: Build a human-in-the-loop approval pipeline.**
Task suspends -> Email User -> User clicks "Approve" -> Webhook callbacks Scheduler -> Task resumes.

**Q568: How to manage retries with poison message queues?**
DLQ (Dead Letter Queue). Store failed msg + error. Admin UI to inspect, fix payload, and replay.

**Q569: Build a system to auto-reschedule failed jobs intelligently.**
Analyze error. Transient (Network)? Retry fast. Resource (OOM)? Retry with more RAM / different node. Logic? No retry.

**Q570: How would you implement transactional workflows?**
Saga Pattern. Orchestrator executes steps. If step N fails, execute Compensating Transactions for N-1..1.

---

## 🔁 Repeatable Patterns at Scale

**Q571: How would you auto-scale background workers?**
Monitor Queue Depth (Lag). Scale Metric: `Lag / TargetProcessingTime`. Add workers to meet target.

**Q572: Design a document translation platform with async queues.**
Upload -> S3 -> Queue -> Translation Service (Long running) -> S3 -> Callback/Notification.

**Q573: How do you log, trace, and monitor microservice chains?**
(Repeated). OpenTelemetry. Distributed Tracing.

**Q574: Design a user impersonation feature for admins.**
Generate "Impersonation Token" (JWT with `sub: target_user`, `actor: admin`). Log all actions with Actor ID.

**Q575: How would you build a platform for email campaigns?**
Template Engine. Personalization. List Mgmt. Batch sending (MTA warm-up). Tracking (Pixel/Link Rewrite).

**Q576: Build a system for uploading, converting, and hosting documents.**
Upload -> Convert (LibreOffice/Pandoc) -> PDF/HTML -> CDN.

**Q577: Design a feature to restore deleted user data.**
Soft Delete (`is_deleted = true`). Undelete = set to false. Hard delete after 30 days (Garbage Collection).

**Q578: Build a bulk data import engine with validation and rollback.**
Staging table. Validate rows. If interactions pass, Copy to Main table inside Transaction. Errors? Reject batch or partial success report.

**Q579: Design a tag suggestion engine for uploaded content.**
Image Classification / NLP on text. Return confidence scores. User confirms.

**Q580: How to sync user data between mobile and web in real-time?**
WebSocket. Push changes. Local DB on mobile updates UI.

---

## 🧑‍💼 Enterprise & SaaS-Oriented

**Q581: Design a tenant-aware SaaS with strong isolation.**
Separate Schema/DB per tenant (Gold Tier) or Discriminator Column (Standard Tier). Middleware injects TenantID.

**Q582: Build a customizable dashboard backend for users.**
Store Layout widgets config (JSON). Dynamic queries based on widget type. Aggregation API.

**Q583: Design an SLA tracker for customers.**
Monitor SLIs. Calculate Uptime/ErrorBudget. Generate Report. Calculate Credit if SLA breached.

**Q584: Build a per-client throttling system.**
Ratelimit per TenantID. Tiers (Free: 10rps, Pro: 100rps). Redis Token Bucket.

**Q585: Design a notification preference manager.**
Table `Preferences` (User, Channel, MessageType, Enabled). Filter notification requests against this table.

**Q586: Build a billing system for tiered plans + metered usage.**
Stripe Subscription + Usage Records. Metering Service aggregates usage events. Reports to Stripe.

**Q587: How to offer usage insights to business users?**
Analytics pipeline -> OLAP DB -> Embedded Analytics (Charts) in App.

**Q588: Design a secure cross-tenant admin view.**
SuperAdmin Role. Ignored Tenant Filter. Strict Logging.

**Q589: Build a system to manage legal agreements by geography.**
Versioned Terms. Map `Country -> TermsVersion`. On Login, check `UserCurrentTerms < Latest`. Force accept modal.

**Q590: How would you audit privileged actions by admins?**
(Repeated). Immutable Audit Log.

---

## ✨ Real-World Inspired Challenges

**Q591: Build the backend of a meditation app.**
Audio Streaming. Progress tracking (Streaks). Offline support.

**Q592: Design a backend for a celebrity live-stream Q&A app.**
Massive Fanout (One broadcaster, million viewers). Chat moderation (AI + Human). Question Upvoting queue.

**Q593: How would you build a fantasy sports league platform?**
Ingest Stats. Calculate Points. Update League Standings. Locking rosters before game time.

**Q594: Design a carbon footprint tracking system.**
Integrate with Travel/Purchase APIs. Carbon values DB. Calculator. Visualization.

**Q595: Build a QR-based restaurant menu + ordering system.**
Static/Dynamic QR. Menu API. Cart. Order -> Kitchen Display System (KDS). Payment.

**Q596: Design a system for online multiplayer quiz games.**
WebSocket Lobby. Sync Question Push. Timer on server. Collect Answers. Broadcast Score.

**Q597: Build a secure health report sharing system.**
E2E Encryption. Time-bound access links. Audit everything.

**Q598: Design an architecture for emergency SOS alerts.**
High Availability. Priority SMS/Push gateways. Location tracking. Integration with responders.

**Q599: How would you build a personalized trip planner backend?**
Constraint Solver (Budget, Time, POIs). TSP (Traveling Salesman) approximation. Maps API.

**Q600: Design a backend for a local community bulletin board.**
Geospatial search (Posts near me). Moderation (Text filter). User Reputation.

---

## 🏥 Healthcare, Legal & Education Domains

**Q601: Design an appointment booking system for hospitals.**
Resource Scheduling. Doctors have slots. ACID transactions to prevent double booking. Notification reminders. Integration with EMR (Electronic Medical Records).

**Q602: How would you build a prescription refill tracker?**
State Machine: Prescribed -> Requested -> Approved -> Filled -> Picked Up. Compliance checks against remaining refills.

**Q603: Build a vaccination record platform with auditability.**
Immutable Ledger (Blockchain or QLDB) for integrity. User access via Health ID. QR code generation for verification.

**Q604: Design a HIPAA-compliant messaging app for doctors.**
E2E Encryption. Ephemeral messages options. Strict access logging. No data in notifications (just "New Message"). Remote wipe capability.

**Q605: Build a legal document workflow & e-signing system.**
Document Storage (encrypted at rest). Versioning. Signature capture (Audit trail of IP, Time). PDF sealing (Digital Signature).

**Q606: Design an online examination platform with proctoring.**
Secure Browser (Lockdown). Webcam streaming to Proctor/AI. Question Bank randomization. Scalable submission backend.

**Q607: How to prevent cheating in online exams?**
AI Monitoring (Gaze tracking, multiple faces). Plagiarism check on text answers. Disable Copy/Paste. Watermarking questions.

**Q608: Design a medical image storage & retrieval system.**
PACS (Picture Archiving and Communication System). Store DICOM files in Object Store. Metadata in DB. Streaming viewer (Tile based loading).

**Q609: Build a therapy session scheduling and journal app.**
Privacy first. Encrypted Journals (User key). Video call integration (WebRTC). Calendar sync.

**Q610: Design a patient-doctor chat system with language translation.**
Real-time translate (Google Translate API) in pipeline. Store original and translated. Disclaimer on accuracy.

---

## 🧩 API Design and Management

**Q611: How would you implement API rate limiting per user and API key?**
(Repeated). Redis Lua script. Check `User:Count` and `Key:Count`. Decrement/Reset on window.

**Q612: Design an API gateway with auth, logging, and caching.**
Ingress controller (Nginx/Kong). Plugins for JWT Auth, Request logging to ELK, Proxy Caching for GET.

**Q613: How to support API versioning for backward compatibility?**
(Repeated). URL (/v1/).

**Q614: Design a public-facing developer portal for APIs.**
Documentation (Swagger UI). Sign up/API Key management. Sandbox/Try-it-out. Usage dashboard.

**Q615: How to throttle APIs based on time of day or region?**
Dynamic config for Rate Limiter. `If (Region == EU and Time == Peak) Limit = 100 else Limit = 500`.

**Q616: Design a webhook delivery system with retries and dead letter queues.**
(Repeated). Reliable delivery is key.

**Q617: How do you build an internal API dependency graph?**
Tracing data (OpenTelemetry). Visualize Service Map (Service A calls B, C). Detect cycles.

**Q618: Build an API change detection & alerting platform.**
Monitor OpenAPI specs in Git. Diff specs. If breaking change detected -> Alert PR Author.

**Q619: How would you design a GraphQL gateway over microservices?**
Schema Stitching or Apollo Federation. Gateway merges subgraphs from microservices into one Supergraph.

**Q620: Implement usage analytics for each endpoint in an API platform.**
Log request path + status. Aggregation pipeline. Dashboard: "Top 10 Slowest Endpoints", "Most used Endpoints".

---

## ✅ Data Quality & Integrity

**Q621: Design a system for real-time data validation.**
Schema Registry (Avro). Producer validates before sending. Consumer validates on read. DLQ on failure.

**Q622: How would you ensure consistency in duplicate data across services?**
Single Source of Truth (Master Data Management). Other services reference ID or subscribe to updates from Master.

**Q623: Design a service for cleaning and deduplicating customer records.**
ETL job. Fuzzy matching (Name, Address). Merge logic (Golden Record).

**Q624: Build a platform for automated schema consistency checks.**
CI/CD step. Check DB migration scripts against rules. Check Code-DB alignment.

**Q625: Design a data poisoning detection mechanism in ML pipelines.**
Statistical Analysis of Training Data. Outlier detection. Track provenance of data (User Reputation).

**Q626: Build a distributed checksum validation system.**
Periodically read data chunks. Calculate CRC32. Compare with stored checksum. Alert on mismatch (Bit rot).

**Q627: How do you detect silent data corruption?**
(Same as Q626). End-to-end checksums (Application level).

**Q628: Implement a rollback-safe write-ahead log.**
Append-only log. Checkpoint (Sequence Number). On crash, replay from last checkpoint.

**Q629: Design a “data quarantine” zone for suspect records.**
If validation fails, route to "Quarantine Table/Topic" instead of discarding. Ops review and fix.

**Q630: How do you verify external data imports are safe?**
Sandboxed parsing. Schema validation. Value range checks. Virus scan.

---

## 🧪 Testing, Monitoring, and Observability

**Q631: Design a chaos testing platform.**
(Repeated). Orchestrate attacks. measure steady state deviation.

**Q632: Build a synthetic monitoring tool for uptime checks.**
Distributed runners (Lambda/Agents). HTTP GET /health. Assert Status 200, Latency < 1s.

**Q633: Design a traceable, testable deployment pipeline.**
Commit -> Build (ID) -> Test -> Deploy (ID). Trace ID links Code, Artifact, and Release.

**Q634: How do you ensure test coverage for service-to-service interactions?**
Integration tests with Mocks/Stubs. Contract Testing (Pact).

**Q635: Design a zero-downtime release system.**
(Repeated). Rolling/Blue-Green.

**Q636: Build a real-time alerting system for performance regressions.**
Compare current deployment metrics vs previous deployment baseline. Anomaly detection.

**Q637: How to record detailed request/response traces for debugging?**
 Sampling (1% of requests) detailed application logs + headers. Store in specialized store (Jaeger).

**Q638: Design a system to test rollback and forward compatibility.**
CI Job: Migrate DB Up -> Run App v2 -> Migrate DB Down -> Run App v1 -> Verify health.

**Q639: Build a shadow traffic replay system for staging environments.**
Capture Prod Traffic -> Async Replay to Staging (stripping PII). Compare results (Diffy).

**Q640: Design a service dependency graph with fault isolation.**
(Repeated). Circuit breaker status reporting creates live isolation map.

---

## 🛡️ Security & Compliance

**Q641: Design a secure file upload service.**
(Repeated).

**Q642: How would you encrypt large datasets with minimal latency?**
Envelope Encryption. Generate DEK (Data Encryption Key). Encrypt Data with DEK (AES-GCM). Encrypt DEK with Master Key (KMS). Store EncryptedData + EncryptedDEK.

**Q643: Design a secure OAuth2 flow for mobile and web.**
PKCE (Proof Key for Code Exchange). Prevents interception of Auth Code.

**Q644: Implement data access policies based on roles and geography.**
ABAC (Attribute Based Access Control). Policy: `Allow if Role=Manager AND User.Location == Resource.Location`.

**Q645: Build a system to audit and revoke stale credentials.**
Scanner detects unused Keys (> 90 days). Notify owner. Disable. Delete.

**Q646: How would you secure long-lived background processes?**
Identify workload identity (Service Account). Rotate tokens automatically. Least Privilege.

**Q647: Design a phishing detection system for emails.**
Analyze Sender SPF/DKIM/DMARC. Analyze Body Links (Domain reputation). NLP on Content (Urgency/Threats).

**Q648: Build a service to manage and rotate secrets.**
(Repeated). Vault.

**Q649: How to detect unusual login patterns across geographies?**
Velocity Check: Login London 10:00, Login NYC 10:05. Impossible Travel. Trigger 2FA/Block.

**Q650: Design a system for 2FA backup codes and recovery.**
Generate 10 random codes. Hash them. Store Hashes. User verifies identity (Email/Support) to recover.

---

## 🔍 Search, Filtering, and Indexing

**Q651: Design a full-text search engine with autocomplete.**
(Repeated).

**Q652: Build a system for personalized search ranking.**
(Repeated). LTR (Learning to Rank).

**Q653: How to implement federated search across services?**
Gateway scatters query to Service A, B, C. Gathers results. Normalizes ranking. Returns merged list.

**Q654: Design an e-commerce filter engine (facets, range, categories).**
Elasticsearch Aggregations. Bitmaps for boolean filters (Fast Intersection).

**Q655: Implement typo-tolerant search indexing.**
(Repeated). Fuzzy search.

**Q656: Design a system for trending keyword detection.**
(Repeated). Streaming counters.

**Q657: How would you enable real-time reindexing without downtime?**
Zero-downtime alias switching. Build new index. switch alias `prod_start` to `new_index`.

**Q658: Build a tagging + filtering engine for massive datasets.**
Inverted Index.

**Q659: Design a “related searches” suggestion system.**
Analysis of Query logs. "Users who searched X also searched Y". Co-occurrence matrix.

**Q660: How to support multi-language search effectively?**
Language detection. Specific Analyzers (Stemming/Stopwords) per language. Index `title_en`, `title_es`. Query relevant field.

---

## ⚙️ Rules Engines & Automation

**Q661: Design a rules engine for fraud detection.**
DSL (Domain Specific Language) for rules (`amount > 5000 AND location != home`). Rete Algorithm for pattern matching.

**Q662: Build a workflow automation engine like Zapier.**
Triggers (Poller/Webhook) -> Action Pipeline. Connector abstraction (Auth, API mapping).

**Q663: Implement a rule-based alerting system.**
(Repeated).

**Q664: Design a user-triggered automation builder.**
UI constructs Logic Tree (JSON). Backend interpreter traverses tree executing nodes.

**Q665: Build a no-code workflow designer backend.**
(Same as Q664/Zapier).

**Q666: Design a credit card fraud rules platform.**
(Same as Q661).

**Q667: Implement rule versioning with rollback support.**
Store Rule Sets as immutable versions. Evaluate Context against RuleSet v123.

**Q668: Build an engine that triggers actions based on threshold breaches.**
Metric Stream -> Threshold Check -> Action Dispatcher. State (Breach Start Time) to prevent flapping.

**Q669: Design a low-latency decision engine for recommendations.**
Pre-compute Rules into Decision Trees/Lookup Tables. Load into memory.

**Q670: How would you design a conflict resolution system for overlapping rules?**
Priority / Salience. Run highest priority rule. Or "First Match". Or "Best Match".

---

## 🧾 Document Processing & Content Systems

**Q671: Build a resume parsing and ranking engine.**
OCR/Text Extraction. Named Entity Recognition (Extract Name, Skills). Match against Job Description keywords.

**Q672: Design a contract redlining and version tracking system.**
Store Document + Change Sets. UI renders diffs. Audit trail of who changed what.

**Q673: How would you convert scanned documents to searchable formats?**
OCR Pipeline (Tesseract/Textract). Index output text in Elasticsearch. Overlay text on image in PDF.

**Q674: Design a collaborative Markdown document editor.**
(Repeated).

**Q675: Build a service to extract tables and charts from PDFs.**
Computer Vision (Layout Analysis) to detect table boundaries. Extract cells. Reconstruct structure.

**Q676: How to handle bulk uploads and format normalization?**
Async Job. Detect MimeType. Converters (Magick/Pandoc) to standard format. Validation.

**Q677: Design a plagiarism detection platform.**
(Repeated).

**Q678: Build a CMS for publishing across mobile and web.**
Headless CMS. Content = JSON blocks. API serves content. Clients render.

**Q679: Design a privacy-aware document sharing system.**
(Repeated).

**Q680: Build a legal compliance document delivery tracker.**
Track delivery (Email opened/Link clicked). Require e-signature/Ack.

---

## ⛓️ Chain of Events & State Machines

**Q681: Build a state machine for user onboarding.**
States: EmailPending, ProfilePending, TeamPending, Active. Events transition states. Persist current state.

**Q682: Design a job status tracker with retry and escalation.**
(Repeated).

**Q683: Implement a finite state machine for delivery systems.**
Order Placed -> Preparing -> OutForDelivery -> Delivered. Valid transitions only.

**Q684: Build a versioned state transition audit trail.**
Log every transition. Replay to debug.

**Q685: How to validate illegal transitions in distributed systems?**
Optimistic locking on State. `UPDATE orders SET status='Shipped' WHERE id=1 AND status='Prepared'`. If row count 0, illegal.

**Q686: Design a resume-from-checkpoint download system.**
(Repeated).

**Q687: How would you visualize system transitions and workflows?**
Generate Graphviz/Mermaid diagram from State definitions. realtime overlay of current instance position.

**Q688: Build a system to pause/resume workflows on external triggers.**
Wait State. Webhook wakes up workflow.

**Q689: How to build real-time state dashboards for business operations?**
Change Data Capture (CDC) from DB -> WebSocket -> Dashboard.

**Q690: Implement state reconciliation across distributed replicas.**
Vector Clocks / Version Vectors. Merge logic.

---

## ⏳ Time-Sensitive & Temporal Systems

**Q691: Build a daily digest email generator.**
Collect events for user throughout day in Bucket. At 9 AM local time, render Template with bucket items. Send. Empty bucket.

**Q692: Design a system to expire items at dynamic times.**
Priority Queue (Delay Queue). Redis Keyspace Notifications (Expire event).

**Q693: How to implement TTL with accuracy and efficiency?**
(Repeated).

**Q694: Build a high-resolution calendar scheduler.**
Bitmask for time slots? Or Segment Tree for intervals.

**Q695: Design a system for recurring job execution with drift correction.**
Calculate NextRun = LastRun + Interval. (Fixed Rate vs Fixed Delay).

**Q696: How to detect temporal anomalies (e.g., no events)?**
"Dead Man's Switch". If no heartbeat/event in X time, trigger alert.

**Q697: Implement a system for calculating historical state timelines.**
Store valid time ranges `[Start, End]`. Temporal join. "What was status at Time T?".

**Q698: How would you enable users to “rewind” system state?**
Event Sourcing. Replay events up to Timestamp T.

**Q699: Design a time-aware feed with relevance and freshness.**
Score = Relevance * DecayFunction(Time).

**Q700: How to manage time zones and DST in event scheduling?**
Store everything in UTC. Convert to User TZ for display. Compute recurring rules using User TZ library (IANA DB).

---

## 🎥 Multimedia & Real-Time Media Systems

**Q701: Design a system like YouTube with video upload, transcoding, and streaming.**
Uploads to S3. Lambda triggers MediaConvert (transcoding to HLS/DASH different bitrates). CDN distributes. Metadata in SQL/NoSQL.

**Q702: How would you build a real-time video conferencing app backend?**
WebRTC for P2P media. SFU (Selective Forwarding Unit) (like Jitsi/Kurento) for scaling multiparty calls. Signaling via WebSocket.

**Q703: Design a podcast hosting and distribution platform.**
Upload Audio. Generate RSS Feed. Serve RSS via CDN. Analytics via Log Analysis of Access Logs (Range requests handling).

**Q704: Build a low-latency live streaming infrastructure.**
RTMP Ingest -> Transcode -> WebRTC/HLS-LL Out. CDN Edge Compute to relay streams.

**Q705: How do you implement video thumbnail generation at scale?**
(Repeated). FFMPEG worker.

**Q706: Design an audio transcription service with multi-language support.**
Audio Upload -> Split -> STT Service (Whisper/Google Speech) -> Text. Mapping timestamps to text.

**Q707: Build a backend for short video creation and sharing (like TikTok).**
Feed Generation crucial. Pre-compute feeds. Heavy CDN usage. Video recommendation engine.

**Q708: How to handle content moderation for user-uploaded media?**
AI Pipeline (NSFW detection, Copyright hash match). Flag for Human Review.

**Q709: Design a distributed video encoding pipeline.**
Split video into chunks. Parallel workers encode chunks. Merge chunks. (MapReduce style).

**Q710: Build a collaborative video annotation and feedback system.**
Store comments with `(timestamp, xy_coord)`. Overlay on player. Sync via WebSocket.

---

## 🌍 Geo-Distributed & Multi-Region Architectures

**Q711: Design a global ride-sharing platform with geo-aware matchmaking.**
Sharded by Geohash. Matchmaking looks in adjacent geohashes. Global user profile replication.

**Q712: How to replicate databases across continents with low latency?**
(Repeated). Async replication.

**Q713: Build a global DNS management platform.**
(Repeated). Anycast.

**Q714: Design a latency-aware content delivery network (CDN).**
Edges in ISPs. Route user to best Edge via DNS or HTTP Redirect based on real-time network measurements (RUM).

**Q715: Build a user session store that’s accessible worldwide.**
Replicated Redis (CRDT based like Redis Enterprise) or DynamoDB Global Tables.

**Q716: Design a disaster-tolerant data backup system across regions.**
Cross-Region Replication (CRR) features of S3/DynamoDB. Periodic snapshot transfer.

**Q717: How to route traffic based on geographic failover policies?**
DNS Health Checks. If EU endpoint fails, DNS updates to point EU traffic to US East.

**Q718: Design a multi-region checkout flow for an e-commerce site.**
Cart must be consistent. Pin session to region. If region fails, cart might be lost (tradeoff) or lazily replicated.

**Q719: How to resolve conflicts in distributed systems with time skew?**
Use Google TrueTime (Atomic clocks) or Logic Clocks (Lamport/Vector) instead of wall clock.

**Q720: Build a privacy-aware geo-location logging system.**
Fuzz location (truncate lat/long). Drop precision / blur regions.

---

## 🔐 Access Control & Permissions

**Q721: Design a role-based access control system (RBAC).**
(Repeated).

**Q722: How would you implement attribute-based access control (ABAC)?**
(Repeated).

**Q723: Build a secure permission audit trail system.**
(Repeated).

**Q724: How to manage temporary access to sensitive data?**
Time-bound tokens. TTL on permissions.

**Q725: Design a permission hierarchy system for nested organizations.**
Materialized Path or Nested Sets model in DB. Inherit permissions from Ancestors.

**Q726: How do you revoke user access instantly across services?**
Short Token TTL (5 mins). Valid "Revocation List" broadcast to all services or checked at Gateway.

**Q727: Build a user delegation system (acting on behalf of another user).**
Token exchange. Admin gets token with scopes of User X. Audit logs record "Admin acting as User X".

**Q728: How to handle access control for shared resources?**
ACL on resource. `Resource.ACL = [UserA: R, UserB: RW]`.

**Q729: Implement fine-grained permissions at object-level scope.**
(Same as Q728). Policy decision point.

**Q730: Design a service for reviewing and approving access requests.**
Workflow engine. Request access -> Manager approves -> Provisioning.

---

## 🎨 UX & Frontend-Driven Backend Design

**Q731: Design a backend for a drag-and-drop form builder.**
Store Form definition as JSON Schema. Validation rules stored. Submission handler interprets schema.

**Q732: Build a recommendation engine for a product configurator.**
Constraint Satisfaction Problem. "If Engine=V8, Transmission!=Manual". Returns valid options.

**Q733: How to support undo/redo functionality for multi-user apps?**
Command Pattern. Stack of Actions. Reversing action must be possible.

**Q734: Design a system to store user dashboards and widgets.**
Widget Config + Layout coordinates.

**Q735: Build a flexible notification preference backend.**
(Repeated).

**Q736: Design a theme and layout manager for SaaS products.**
CSS Variables / JSON Theme config. Served with App.

**Q737: How to implement live cursor tracking (like Figma)?**
WebSocket Broadcast. ephemeral state. High frequency updates (interpolate on client).

**Q738: Design an interface versioning system for backward compatibility.**
Serve older JS bundles. API maintains old contracts.

**Q739: Build a progressive onboarding backend system.**
Track user progress (Steps completed). Unlock features based on progress.

**Q740: Design a feature-tour trigger system based on user behavior.**
Analytics events -> Trigger Rule -> Return "ShowTour" flag to client.

---

## 🏢 Enterprise-Scale Operations

**Q741: Design a unified audit log system across services.**
(Repeated).

**Q742: Build a unified identity provider for SSO integration.**
(Repeated).

**Q743: Design a dashboard for tracking KPIs in real time.**
(Repeated).

**Q744: How to support multi-department billing and reporting?**
Tag resources with CostCenter. Aggregation pipeline.

**Q745: Build a security incident reporting system.**
Ticket system specialized for SecOps. Workflow. Evidence attachment.

**Q746: Design a deployment pipeline with approval workflows.**
Manual Gate in CI/CD.

**Q747: Build a productivity insights system using activity data.**
Analyze commits, tickets, meetings. (Privacy concerns critical).

**Q748: How to manage configurations across thousands of tenants?**
Hierarchical Config: Default -> Regional -> Tenant -> User Override. Merge at runtime.

**Q749: Design a compliance reporting engine with export capabilities.**
Scanners -> Database -> Report Generator (PDF/CSV).

**Q750: Build a system for managing employee hardware inventory.**
Asset Management DB. Assignment/Return workflow. Depreciation calculation.

---

## 📉 Metrics, Alerts & Reliability

**Q751: Build a dynamic SLAs tracking system.**
(Repeated).

**Q752: How would you detect cascading failures in microservices?**
Dependency graph health.

**Q753: Design a health-check system with progressive fallbacks.**
If Primary DB down -> Switch to ReadOnly. If Cache down -> Go to DB. If both down -> Static Site.

**Q754: Build a usage spike detection system.**
Rate of change of Request Count.

**Q755: How do you throttle noisy or failing components automatically?**
Circuit Breakers.

**Q756: Design a smart alerting system to reduce false positives.**
(Repeated).

**Q757: Build an adaptive retry strategy system.**
Backoff based on server load headers.

**Q758: Design a system that auto-pauses non-critical jobs during outages.**
"Panic Mode" switch. Job workers check switch before processing.

**Q759: Implement global service status pages across services.**
Public Status Page -> Aggregates Health checks.

**Q760: Build a root cause analysis suggestion engine using logs.**
Correlate Error Logs with Change Events (Deployments/Config changes).

---

## 🧬 Emerging Use Cases & Next-Gen Apps

**Q761: Design a decentralized identity verification platform.**
(Repeated).

**Q762: Build an AI-based resume ranking system.**
(Repeated).

**Q763: Design a real-time multiplayer chess engine.**
Server validates moves. Maintenance of FEN string (Board state). Clock management.

**Q764: Build a backend for digital collectibles and NFT marketplace.**
Web3 Integration (Minting). Indexing Blockchain events. Off-chain metadata.

**Q765: Design a blockchain transaction explorer system.**
(Repeated).

**Q766: How to create a conversational UI backend for banking?**
LLM RAG (Retrieval Augmented Generation). Strict Guardrails.

**Q767: Build a system for auto-generating social media posts.**
Templates + Content Inputs -> LLM -> Image Gen -> Scheduling.

**Q768: Design a marketplace for prompt engineering and AI tools.**
Catalog of Prompts. Versioning. Testing playground.

**Q769: Build a cloud cost optimization recommendation engine.**
Analyze usage patterns. Recommendation: "Downsize EC2", "Buy Reserved Instance".

**Q770: Design a decentralized content publishing platform.**
IPFS for storage. Blockchain for ownership registry.

---

## ⚖️ Consistency, Transactions & Tradeoffs

**Q771: Design a distributed locking system.**
Redis (Redlock). Zookeeper.

**Q772: How would you implement eventual consistency in user profile sync?**
Queue updates. Workers apply.

**Q773: Design a system with guaranteed at-least-once delivery semantics.**
Ack only after processing.

**Q774: How do you handle partial failure in multi-step workflows?**
Compensating transactions.

**Q775: Build a system that supports distributed transactions (2PC/SAGA).**
(Repeated).

**Q776: Design a data reconciliation system between microservices.**
Periodic scan and compare. Repair diffs.

**Q777: Implement an audit-safe compensation mechanism.**
Log compensation logic.

**Q778: Design a data expiration and soft-delete system with recovery.**
TTL + Trash Bin folder.

**Q779: Build a versioned update system with rollback capabilities.**
(Repeated).

**Q780: Design a rate consistency checker for billing systems.**
Verify charged amount against calculated amount.

---

## 📦 Feature Management & Experimentation

**Q781: Build a feature flag platform with targeting rules.**
(Repeated).

**Q782: Design an experiment platform with multiple variant support.**
Multivariate testing.

**Q783: How would you test conflicting experiments safely?**
Exclusion groups. (User in Exp A cannot be in Exp B).

**Q784: Build a kill-switch system for unstable features.**
Global Flag override.

**Q785: How to implement real-time experiment exposure logging?**
Log `Exposure` event when user sees variant.

**Q786: Design a backend to run multi-armed bandit experiments.**
Algo updates weights (traffic allocation) based on real-time success metrics to favor winning variant.

**Q787: Build a system for ramping features based on metrics.**
Auto-ramp: 1% -> 5% -> 10% if Error Rate stable.

**Q788: Design a platform for internal beta testing.**
Dogfooding groups.

**Q789: Build a self-service experimentation platform for PMs.**
UI for defining Hypothesis, Variations, Metrics.

**Q790: Design a feedback loop system for post-launch analysis.**
Automated report generation.

---

## 🧠 Intelligence & User Modeling

**Q791: Build a user interest graph based on actions.**
Graph DB (User -LIKES-> Topic). Weighted edges.

**Q792: Design a real-time cohorting system.**
Classify user on session start.

**Q793: Build a behavioral pattern recognition system.**
Sequence mining.

**Q794: How to create a personalized notification prioritization system?**
Rank notifications by P(Click).

**Q795: Design a model for churn prediction and mitigation.**
Predictive score. Trigger retention workflow.

**Q796: Build a backend for smart autocomplete and prediction.**
(Repeated).

**Q797: Design a “smart mute” system that suppresses irrelevant alerts.**
Feedback loop "Not Helpful".

**Q798: Build a system that learns from user search failures.**
Analyze "No Results" queries. Create synonyms or content gaps.

**Q799: Design a content feed that adapts based on scroll behavior.**
Dwell time tracking. Negatively weight skipped content.

**Q800: Build a real-time intent detection engine for helpdesk chat.**
NLP classifier on first message. Route to Specialist.

---

## 🤖 AI/ML Infrastructure & Data Pipelines

**Q801: Design a system for training ML models on petabytes of data.**
Distributed Training (Data Parallelism). Use Horovod/PyTorch DDP. Store data in S3/HDFS. High-bandwidth networking (InfiniBand/EFA). Checkpointing.

**Q802: Build a feature store for ML pipelines.**
(Repeated). Feast/Tecton.

**Q803: Design a pipeline for real-time ML model inference.**
Load Balancer -> Inference Service (TFServing) -> Model. Monitoring sidecar.

**Q804: How would you version models, datasets, and parameters?**
(Repeated). DVC (Data Version Control) + MLflow.

**Q805: Design an experiment tracking system for data scientists.**
(Repeated).

**Q806: How to ensure reproducibility in ML training pipelines?**
Docker containers. Pin all library versions. Seed RNGs. Version control data/code.

**Q807: Build a scalable hyperparameter tuning platform.**
Controller (Vizier/Katib) spawns trials (workers). Search strategy (Bayesian Optimization). Early stopping.

**Q808: Design a real-time fraud detection ML pipeline.**
(Repeated).

**Q809: Build a system to serve multiple models with low latency.**
Multi-Model Serving (Triton Inference Server). Shared memory.

**Q810: Design a pipeline for explainable AI (XAI) results.**
Inference -> SHAP/LIME calculation job (Async) -> Store explanation.

---

## 💳 Financial, Banking & Trading Systems

**Q811: Design a system to detect suspicious banking transactions.**
Rule Engine + ML Anomaly Detection. Real-time stream processing.

**Q812: Build a micro-lending platform with KYC and credit scoring.**
KYC API integration. Loan DB. Interest calculation engine.

**Q813: Design a stock order matching engine.**
In-memory Order Book (Double Linked List/RB Tree). Matching algorithm (Price/Time priority). Single-threaded for determinism (LMAX Disruptor).

**Q814: How would you design a cross-border payment gateway?**
SWIFT/SEPA integration. FX conversion service. Compliance checks.

**Q815: Build a reconciliation system for bank transactions.**
(Repeated).

**Q816: Design a subscription billing engine with retries and proration.**
(Repeated).

**Q817: Build a system for invoice generation and PDF delivery.**
Async worker. Template engine -> PDF Lib -> Email Service.

**Q818: Design a ledger system for recording financial transactions.**
(Repeated).

**Q819: How do you ensure transactional integrity across currencies?**
Atomic transactions. Storing rates at time of transaction.

**Q820: Build a fraud-resistant peer-to-peer payment system.**
Risk scoring on sender/receiver. Device fingerprinting. 3DSecure.

---

## 📦 Logistics, Supply Chain & Transportation

**Q821: Design a package tracking system for courier networks.**
Scan events at Hubs. Event Stream -> State Update. Customer Query -> Current Status.

**Q822: Build a warehouse inventory management system.**
SKU tracking. Location (Aisle/Bin). First-In-First-Out (FIFO) logic. Barcode scanning integration.

**Q823: Design a last-mile delivery optimization system.**
Route optimization (VRP - Vehicle Routing Problem) solver. Driver App. Real-time traffic adjustment.

**Q824: How would you route trucks in real-time based on traffic?**
Graph matching. Integration with Maps/Traffic API. Re-calculation trigger on delay.

**Q825: Build a return pickup and refund processing system.**
Reverse Logistic workflow. Condition check -> Refund trigger.

**Q826: Design a multi-hop shipment system with dependencies.**
Graph data structure for route. `Hub A -> Hub B -> Hub C`. Track readiness at each hop.

**Q827: How do you track items across warehouses globally?**
Centralized Inventory DB. Distributed lock for transfers.

**Q828: Build a delivery time prediction engine using live data.**
ML Regression Model. Features: Distance, Traffic, Driver History, Weather.

**Q829: Design a cold-chain logistics system with sensor alerts.**
IoT Sensors (Temp). Alert if `Temp > Threshold`. Audit trail for compliance.

**Q830: Build a logistics system for multi-vendor fulfillment.**
Order Splitter. Route sub-orders to Vendor APIs. Aggregate status.

---

## 🌐 Multi-Tenant SaaS & Admin Platforms

**Q831: Design a multi-tenant SaaS backend with strict data isolation.**
(Repeated).

**Q832: Build a tenant-aware rate limiting and quota system.**
(Repeated).

**Q833: How do you manage schema evolution per tenant?**
Migration scripts run per tenant DB. Version tracking table in each tenant DB.

**Q834: Design a system for custom branding per customer.**
Store Assets (Logo, CSS) in CDN under TenantID path. Frontend loads assets dynamically.

**Q835: Build an admin dashboard for monitoring tenant activity.**
(Repeated).

**Q836: Design audit trails for actions performed by tenant admins.**
(Repeated).

**Q837: How to implement access delegation to external consultants?**
Temporary User creation or "Ghost" login capability with audit.

**Q838: Design a tenant lifecycle management platform.**
Onboarding (Provisioning). Suspension (Payment fail). Offboarding (Data Delete/Export).

**Q839: Build a multi-tenant notification preference system.**
(Repeated).

**Q840: Design a config-as-a-service platform for multiple tenants.**
Central Config Service. Hierarchy `App/Tenant/Env`.

---

## 📱 Mobile-First & Offline-First Systems

**Q841: Design a sync service for mobile apps with offline mode.**
(Repeated).

**Q842: Build a mobile push notification delivery system.**
(Repeated).

**Q843: How would you sync mobile caches with eventual consistency?**
ETags / Version numbers.

**Q844: Design a mobile usage analytics pipeline.**
Batch upload from device (minimize radio usage). Ingest to Kafka.

**Q845: Build a delta update system for reducing mobile data usage.**
(Repeated).

**Q846: Design a system for location-based push campaigns.**
Geofencing (OS level). Wake up app -> Check criteria -> Trigger Local Notification.

**Q847: Build a system to handle device ID rotations and merges.**
Stable Internal UserID. Link new DeviceID to UserID. Merge history.

**Q848: How to queue and sync user actions taken offline?**
Local command queue. Replay on connect. Retry on error.

**Q849: Design a multi-device session management system.**
List valid refresh tokens per User. Revoke one or all.

**Q850: Build a mobile telemetry event aggregation system.**
(Repeated).

---

## 🧱 Graphs, Relationships & Social Features

**Q851: Design a system to recommend mutual connections.**
Graph Traversal (BFS depth 2). "Friends of Friends". Rank by number of mutuals.

**Q852: Build a social graph service with depth-based queries.**
Graph DB (Neo4j). Specialized for "Shortest Path" or "K-Hops".

**Q853: How to detect influencer clusters in a network?**
Community Detection algorithms (Louvain/Label Propagation).

**Q854: Design a system for follow/unfollow with eventual consistency.**
Async jobs to update follower counts and feed fanout.

**Q855: Build a friend suggestion engine using graph traversal.**
(Same as Q851).

**Q856: Design a graph-based spam detection platform.**
Analyze connection patterns. "Star" pattern (one node connects to many unrelated nodes fast) = Spammer.

**Q857: Build a group and subgroup system with scoped permissions.**
DAG of Groups. Permission inheritance.

**Q858: Design a real-time "who viewed your profile" system.**
Stream view events. HyperLogLog for unique viewers. Store latest X viewers.

**Q859: Build a common connections insight engine.**
Intersection of adjacency lists.

**Q860: Design a relationship recommendation system using vector similarity.**
Embed User Graph + Interests. Vector Search.

---

## ⚙️ Automation, Agents & Worker Systems

**Q861: Design a cron-as-a-service platform.**
User submits Schedule + Webhook. Scheduler triggers Webhook.

**Q862: Build a task queue system with retry, delay, and dependencies.**
(Repeated).

**Q863: Design a system for prioritizing queued tasks by SLA.**
Priority Queue. Separate queues for Gold/Silver users.

**Q864: Build a distributed worker pool manager.**
Register Workers. Health check. Dispatch tasks.

**Q865: Design a workflow engine with failure recovery.**
(Repeated).

**Q866: Build a job monitoring and alerting platform.**
(Repeated).

**Q867: Design a dynamic workload auto-scaling system.**
(Repeated).

**Q868: Build an orchestrator for data processing pipelines.**
(Repeated).

**Q869: How to detect zombie or stuck background workers?**
Heartbeat requirement. If no heartbeat, mark dead, re-queue task.

**Q870: Design a backpressure-aware task dispatch system.**
Check queue depth or worker latency. Reject/Defer new tasks if saturated.

---

## ⚡ Real-Time & Event-Driven Systems

**Q871: Design a real-time leaderboard system.**
(Repeated).

**Q872: Build a system for real-time fraud alerts via SMS/email.**
Stream processor -> Rule -> Notification Svc.

**Q873: How would you handle event ordering in distributed systems?**
(Repeated).

**Q874: Build a webhook processor with retries and exponential backoff.**
(Repeated).

**Q875: Design a publish/subscribe system with guaranteed delivery.**
Persistent log (Kafka/Pulsar). Ack tracking.

**Q876: Build a system for synchronizing document edits in real-time.**
(Repeated).

**Q877: Design a real-time feed with aggregation and deduplication.**
Windowing function. "X and 3 others liked your post".

**Q878: Build a real-time sentiment analysis system for news.**
Stream News -> NLP -> Sentiment Score -> Dashboard.

**Q879: Design a collaborative whiteboard backend.**
(Repeated).

**Q880: Build a real-time alerting system for critical operations.**
(Repeated).

---

## 📡 IoT, Sensors & Device Management

**Q881: Design a system to track millions of connected devices.**
Device Registry DB. State (Online/Offline) in Redis.

**Q882: Build an alert system for abnormal sensor behavior.**
(Repeated).

**Q883: How do you handle firmware updates over-the-air?**
(Repeated).

**Q884: Design a telemetry ingestion pipeline for edge devices.**
(Repeated).

**Q885: Build a digital twin platform for connected devices.**
Maintain virtual state JSON matching physical state. Simulation capable.

**Q886: Design a rule engine to trigger actions from sensor data.**
(Repeated).

**Q887: Build a configuration push system for IoT endpoints.**
MQTT Retained Messages. Device subscribes to config topic.

**Q888: Design a geofencing service for IoT-enabled vehicles.**
(Repeated).

**Q889: Build a device registration and provisioning platform.**
Secure manufacturing process. Claiming process (Proof of Ownership).

**Q890: Design a secure data sync protocol for unreliable networks.**
Chunked, checksummed, resumable transfer.

---

## ⚖️ Edge, Scale & Chaos Engineering

**Q891: Design a system for edge caching with invalidation strategies.**
Cache at CDN edge. purge-by-tag.

**Q892: How to architect a resilient service mesh?**
Decentralized sidecars. Control plane High Availability. Fallback to direct connection if sidecar fails.

**Q893: Design a chaos engineering tool to inject latency and failure.**
(Repeated). Traffic interception (Service Mesh or eBPF). Rules for injection.
