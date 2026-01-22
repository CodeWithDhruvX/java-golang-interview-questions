## 🟢 Scalability and Availability (Questions 51-60)

### Question 51: How do you design a system that handles millions of users?

**Answer:**
Designing for millions of users requires a distributed architecture focusing on horizontal scaling and decoupling.
1.  **Load Balancing:** Distribute traffic across multiple application servers.
2.  **Database Scaling:** Use Master-Slave replication for reads, Sharding for writes, or NoSQL for massive scale.
3.  **Caching:** Aggressively cache at all layers (CDN, Redis, Browser) to offload the DB.
4.  **Asynchronous Processing:** Use Message Queues (Kafka/RabbitMQ) for non-critical tasks (emails, report generation).
5.  **Microservices:** Break monolithic apps into smaller services to scale independently.

### Question 52: How to scale a system read-heavy workload?

**Answer:**
*   **Caching:** The most effective strategy. Use Redis/Memcached to serve frequent queries from memory.
*   **Database Replication:** Add multiple Read Replicas (Slaves). Point read queries to slaves and writes to the master.
*   **CDN:** Serve static content (images, CSS) from the edge.
*   **Denormalization:** Structure database tables to avoid expensive joins during reads.

### Question 53: How to scale a system write-heavy workload?

**Answer:**
Write-heavy systems are harder to scale than read-heavy ones.
*   **Sharding:** Distribute data across multiple database nodes based on a shard key (e.g., UserID).
*   **NoSQL:** Use write-optimized databases like Cassandra (LSM Trees) or DynamoDB.
*   **Async Writes (Write-Behind):** Write to a message queue first (Kafka) and process/persist to DB asynchronously.
*   **Bulk/Batch Inserts:** Group small writes into fewer large batches to reduce I/O overhead.

### Question 54: What is replication and when to use it?

**Answer:**
Replication is keeping a copy of the same data on multiple machines.
*   **Uses:**
    *   **High Availability:** If one node fails, others can serve data.
    *   **Latency:** Replicate data to regions closer to users.
    *   **Read Scaling:** Distribute read traffic across replicas.
*   **Types:** Active-Passive (Master-Slave), Active-Active (Master-Master).

### Question 55: How to make a system fault-tolerant?

**Answer:**
Fault tolerance is the ability of a system to continue operating without interruption when one or more of its components fail.
*   **Redundancy:** Eliminate Single Points of Failure (SPOF) by having backup components (e.g., multiple server instances, DB replicas).
*   **Replication:** Replicate data across zones/regions.
*   **Circuit Breaker:** Fail fast and recover gracefully to prevent cascading failures.
*   **Health Checks & Auto-recovery:** Kubernetes restarts failed pods automatically.

### Question 56: What is failover?

**Answer:**
Failover is the automatic switching to a redundant or standby computer server, system, hardware component, or network upon the failure of the previously active application.
*   **Example:** If the Primary DB crashes, the system promotes a Read Replica to be the new Primary.
*   **Automation:** Usually handled by Load Balancers or Orchestrators (K8s) using heartbeat monitoring.

### Question 57: What is high availability?

**Answer:**
High Availability (HA) refers to systems that are durable and likely to operate continuously without failure for a long time.
*   **Measurement:** "Nines" of availability.
    *   99.9% uptime = 8.76 hours downtime/year.
    *   99.999% (Five 9s) = 5.26 minutes downtime/year.
*   **Achieved via:** Redundancy, Load Balancing, Clustering, and Failover mechanisms.

### Question 58: Difference between active-passive and active-active systems.

**Answer:**
*   **Active-Passive:**
    *   One node handles traffic (Active); the other waits on standby (Passive).
    *   Passive node only takes over if Active fails.
    *   *Pros:* Simpler logic. *Cons:* Resource waste (passive node sits idle).
*   **Active-Active:**
    *   All nodes handle traffic simultaneously.
    *   *Pros:* Better resource utilization, higher throughput.
    *   *Cons:* Complex synchronization (avoiding data conflicts if both write).

### Question 59: What is graceful degradation?

**Answer:**
Graceful degradation allows a system to maintain limited functionality even when a large portion of it has been destroyed or is inoperative.
*   **Example:** If the "Recommendations" service fails on an E-commerce site, the main page still loads "Products," but the "Recommended for You" section is empty or hidden, rather than showing a 500 Error Page.

### Question 60: What is a throttling mechanism?

**Answer:**
Throttling is the process of limiting the number of actions a user or component can perform in a given time frame to prevent overuse of resources.
*   **Purpose:** Protects the system from becoming unresponsive due to high load (DoS attacks or noisy neighbors).
*   **Implementation:** Token Bucket algorithm, Leaky Bucket.
*   **Result:** Requests exceeding the limit are rejected (HTTP 429 Too Many Requests).

---

## 🟢 Security & Authentication (Questions 61-70)

### Question 61: How to design a secure login system?

**Answer:**
1.  **HTTPS:** Encrypt all traffic using SSL/TLS.
2.  **Hashing:** Never store plain text passwords. Use bcrypt, Argon2, or SCRAM with a unique salt per user.
3.  **MFA:** Implement Multi-Factor Authentication (SMS/TOTP).
4.  **Session Management:** Use secure, HTTPOnly, Secure cookies for session IDs to prevent XSS.
5.  **Rate Limiting:** Prevent Brute Force attacks by limiting failed login attempts.

### Question 62: What is OAuth 2.0?

**Answer:**
The industry-standard protocol for authorization.
*   **Concept:** Allows a user to grant a third-party application access to their resources on another service (e.g., "Log in with Google") without sharing their password.
*   **Roles:** Resource Owner (User), Client (App), Authorization Server (Google), Resource Server (API).
*   **Flow:** User authenticates with Auth Server -> Auth Server issues Access Token -> Client uses Token to access Resource Server.

### Question 63: What is JWT?

**Answer:**
JSON Web Token (JWT) is a compact, URL-safe means of representing claims to be transferred between two parties.
*   **Structure:** `Header.Payload.Signature`
*   **Stateless:** The server doesn't need to store session data; it verifies the signature to trust the token.
*   **Usage:** Authentication (Logged in user), Information Exchange.
*   **Security:** Always use HTTPS. Don't put sensitive secrets in the Payload (it's only Base64 encoded, not encrypted).

### Question 64: How do you store passwords securely?

**Answer:**
1.  **Hashing:** One-way transformation (cannot be reversed).
2.  **Salting:** Add a unique random string to each password before hashing to defeat Rainbow Table attacks.
3.  **Algorithm:** Use slow algorithms like **bcrypt**, **scrypt**, or **Argon2** to make brute-force attacks computationally expensive.
4.  **Pepper:** (Optional) Add a secret key stored separately from the DB to the hash.

### Question 65: What is rate limiting?

**Answer:**
A strategy for limiting network traffic. It sets a cap on how many requests a sender can issue in a specific time window.
*   **Algorithms:**
    *   **Token Bucket:** Tokens are added at a fixed rate; request consumes a token.
    *   **Leaky Bucket:** Requests enter a queue and are processed at a constant rate.
    *   **Fixed Window Counter:** Count requests per minute.
    *   **Sliding Window Log:** More accurate timestamp tracking.

### Question 66: How to secure APIs?

**Answer:**
1.  **Authentication:** Identify the caller (API Keys, OAuth, JWT).
2.  **Authorization:** Ensure caller has permission (RBAC, Scopes).
3.  **Encryption:** Force HTTPS everywhere.
4.  **Input Validation:** Sanitize all inputs to prevent SQL Injection/XSS.
5.  **Rate Limiting:** Prevent DoS.
6.  **CORS:** Restrict which domains can call your API.

### Question 67: What is CORS?

**Answer:**
Cross-Origin Resource Sharing (CORS) is a browser security mechanism that restricts cross-origin HTTP requests.
*   **Mechanism:** When JS on `domainA.com` calls API on `domainB.com`, the browser sends a pre-flight `OPTIONS` request.
*   **Headers:** Server must respond with `Access-Control-Allow-Origin: domainA.com` (or `*`) for the browser to allow the actual request.
*   **Goal:** Prevents malicious sites from reading data from other sites where the user is logged in.

### Question 68: Explain SSL/TLS in web communication.

**Answer:**
SSL (Secure Sockets Layer) and its successor TLS (Transport Layer Security) encrypt data between client and server.
*   **Handshake:**
    1.  Client sends "Hello" + supported ciphers.
    2.  Server sends "Hello" + Certificate (Public Key).
    3.  Client verifies Certificate with CA (Certificate Authority).
    4.  Client creates a symmetric Session Key, encrypts it with Server's Public Key, and sends it.
    5.  Server decrypts Session Key with Private Key.
    6.  Both parties communicate using the symmetric Session Key (encrypted tunnel).

### Question 69: What is cross-site request forgery (CSRF)?

**Answer:**
An attack that forces an end user to execute unwanted actions on a web application in which they are currently authenticated.
*   **Example:** Malicious site has a hidden form that POSTs to `bank.com/transfer`. If user is logged into bank, the browser sends session cookies, and the transfer happens.
*   **Prevention:** Use **CSRF Tokens** (random values injected into forms/headers) that the malicious site cannot guess.

### Question 70: What is cross-site scripting (XSS)?

**Answer:**
An attack where malicious scripts are injected into trusted websites.
*   **Reflected XSS:** Script is in the URL (e.g., search query) and executed when victim clicks link.
*   **Stored XSS:** Script is saved in DB (e.g., comment section) and executed on every visitor's browser.
*   **Prevention:**
    *   **Sanitize HTML:** Escape special characters (`<` becomes `&lt;`).
    *   **CSP (Content Security Policy):** Restrict sources of executable scripts.

---

## 🟢 Design Specific Systems (Questions 71-80)

### Question 71: Design YouTube

**Answer:**
*   **Core Features:** Upload video, Stream video, Search, Comments/Likes.
*   **Architecture:** Microservices.
*   **Storage:**
    *   **Blob Storage (S3):** Store actual video files.
    *   **Metadata DB (SQL):** Video title, owner, view count.
*   **Processing:** Background workers transcode video into multiple formats/resolutions (720p, 1080p, HLS/DASH).
*   **Streaming:** Use **CDN** heavily to serve video chunks from closest edge location.

### Question 72: Design WhatsApp

**Answer:**
*   **Core:** 1-on-1 Chat, Group Chat, Sent/Delivered/Read Status.
*   **Protocol:** **WebSocket** or **MQTT** for real-time persistent connection.
*   **Storage:**
    *   **Ephemeral:** Messages are stored on server only until delivered. Once delivered, deleted from server (in theory).
    *   **Local DB:** SQLite on user's device stores chat history.
*   **Encryption:** End-to-End Encryption (E2EE) using Signal Protocol.

### Question 73: Design Twitter

**Answer:**
*   **Core:** Post Tweet, Timeline (Feed), Follow.
*   **Read vs Write:** Read-heavy (1000:1 ratio).
*   **Data Model:**
    *   User Table, Tweet Table, Follow Table.
*   **Timeline Generation:**
    *   **Pull Model (Fan-out on Load):** Query all followees' tweets and sort. (Slow).
    *   **Push Model (Fan-out on Write):** When User A tweets, push ID to all followers' pre-computed timeline caches (Redis). (Fast reads).
*   **Hybrid:** Push for normal users, Pull for celebrities (millions of followers).

### Question 74: Design Uber

**Answer:**
*   **Core:** Rider matches Driver, Location Tracking, Payments.
*   **Location Service:**
    *   Drivers send location every 5s.
    *   **Geospatial Index:** Use **QuadTree** or **Google S2** to efficiently find "drivers within X km".
*   **Matching:** Algorithm considers distance, driver rating, ETA.
*   **State Machine:** Trip states (Requested -> Matched -> Started -> Ended).

### Question 75: Design Instagram

**Answer:**
*   **Core:** Photo upload, Feed, Stories.
*   **Storage:**
    *   Photos/Videos -> Object Storage (S3).
    *   Metadata -> PostgreSQL (sharded by UserID) or Cassandra.
    *   Relations (Follows) -> Graph DB or standard SQL association table.
*   **Feed Generation:** Similar to Twitter (Pre-generated feeds stored in Redis cache).

### Question 76: Design a URL Shortener (like bit.ly)

**Answer:**
*   **Goal:** Convert long URL -> Short URL (e.g., `bit.ly/xyz123`).
*   **Encoding:** Base62 (a-z, A-Z, 0-9).
*   **DB Schema:** `id`(Auto-increment), `long_url`, `short_code`.
*   **Logic:**
    *   Convert DB `id` to Base62 string. ID `100` -> `1C`.
    *   Or use a Distributed ID Generator (Snowflake) to get unique ID, then Base62 it.
*   **Redirect:** HTTP 301 (Permanent) or 302 (Temporary). 301 caches at browser (saving server load but losing analytics). 302 hits server every time (better analytics).

### Question 77: Design a file storage system (like Dropbox or Google Drive)

**Answer:**
*   **Core:** Upload, Sync, Share, Versioning.
*   **Chunking:** Split large files into 4MB chunks. Store chunks in S3.
*   **Deduplication:** Check hash of chunk. If chunk exists, don't upload again; just reference it. Saves massive storage.
*   **Metadata DB:** Tracks file hierarchy (Folder -> File -> Lists of Chunks).
*   **Sync:** Client polls or keeps long-polling connection to detect changes.

### Question 78: Design a news feed (like Facebook)

**Answer:**
*   **Components:** Feed generation, Feed publishing, News Feed API.
*   **Algorithm:** Score = (Affinity * Weight * Time Decay).
    *   Show relevant posts, not just chronological.
*   **Architecture:** Fan-out-on-write (Push) acts as a cache.
*   **Pagination:** Cursor-based pagination (not offset-based) for infinite scroll efficiency.

### Question 79: Design a video streaming service (like Netflix)

**Answer:**
*   **Content:** Movies/Shows (Static, High Quality). Nothing is live.
*   **Processing:**
    *   Ingest raw video.
    *   **Transcoding:** Convert to different resolutions (4K, 1080p, Mobile) and codecs (H.264, VP9).
    *   **Packaging:** HLS or DASH (Adaptive Bitrate Streaming).
*   **Delivery:** Open Connect (Netflix's custom CDN) placed directly in ISP networks to reduce latency.

### Question 80: Design an e-commerce platform (like Amazon)

**Answer:**
*   **Services:** Product, Search, Cart, Order, Payment, Inventory.
*   **Search:** Elasticsearch for full-text search and filtering.
*   **Inventory Management:** Strict consistency needed. Use database locks or optimistic locking to prevent overselling ("Last item" problem).
*   **Cart:** Persisted in Redis (fast access) + DB (long-term).
*   **Checkout:** Distributed transaction (Saga pattern) across Order, Payment, and Inventory services.

---

## 🟢 Monitoring, Logging & DevOps (Questions 81-90)

### Question 81: How to monitor a distributed system?

**Answer:**
Monitoring involves collecting, aggregating, and analyzing metrics to check system health.
*   **Four Golden Signals:**
    1.  **Latency:** Time taken to service a request.
    2.  **Traffic:** Demand on the system (req/sec).
    3.  **Errors:** Rate of failing requests (5xx).
    4.  **Saturation:** How "full" the service is (CPU/Memory usage).
*   **Tools:** Prometheus, Datadog, Nagios.

### Question 82: What are metrics, logs, and traces?

**Answer:**
The "Three Pillars of Observability":
1.  **Metrics:** Aggregatable numerical data (e.g., CPU=80%, TPS=500). "What is happening?"
2.  **Logs:** discrete events/records (e.g., "Error: DB connection failed at 10:00"). "Why is it happening?"
3.  **Traces:** Request flow across multiple microservices. "Where is the latency/error?"

### Question 83: What is Prometheus and Grafana?

**Answer:**
*   **Prometheus:** A time-series database and monitoring system. It *pulls* (scrapes) metrics from services at intervals.
*   **Grafana:** A visualization tool. It queries Prometheus (or other sources) and displays data in beautiful dashboards (Graphs, Gauges, Heatmaps).
*   **Integration:** Typical stack: App exposes `/metrics` -> Prometheus scrapes -> Grafana visualizes.

### Question 84: What is centralized logging?

**Answer:**
In distributed systems, looking at logs on individual machines is impossible.
*   **Solution:** Aggregate logs from all services into a central location.
*   **ELK Stack:**
    *   **E**lasticsearch (Search & Store).
    *   **L**ogstash (Collect & Transform).
    *   **K**ibana (Visualize).
*   **Modern alternative:** EFK (Fluentd instead of Logstash) or PLG (Promtail, Loki, Grafana).

### Question 85: How to detect system bottlenecks?

**Answer:**
1.  **Load Testing:** Use tools like JMeter/Locust to stress the system.
2.  **Profiling:** Use CPU/Memory profilers (pprof in Go) to find hot functions.
3.  **Tracing:** Use Jaeger/Zipkin to see which service or DB query is taking the most time in a request chain.
4.  **Database Analysis:** Check slow query logs and explain plans.

### Question 86: How do you debug a distributed system?

**Answer:**
Debugging is hard because state is spread across nodes.
1.  **Correlation ID:** Assign a unique ID to every incoming request and pass it to all downstream services. Log this ID everywhere.
2.  **Distributed Tracing:** Visualize the request path.
3.  **Centralised Logging:** Search logs by Correlation ID.
4.  **Reproducibility:** Capture state to reproduce issues in stage/dev.

### Question 87: What is canary deployment?

**Answer:**
A deployment strategy where the new version is rolled out to a small subset of users (e.g., 5%) first.
*   **Process:** Monitor the canary metrics. If stable, gradually increase percentage (10% -> 50% -> 100%).
*   **Benefit:** Reduces risk. If the new version has a bug, only a few users are affected.

### Question 88: What is blue-green deployment?

**Answer:**
A strategy using two identical environments: Blue (Active/Old) and Green (Idle/New).
*   **Process:**
    1.  Deploy new version to Green.
    2.  Test Green.
    3.  Switch Load Balancer to point to Green.
*   **Benefit:** Zero downtime, instant rollback (switch LB back to Blue).
*   **Cost:** Requires double the infrastructure resources.

### Question 89: How do you handle configuration in distributed systems?

**Answer:**
Hardcoding configs is bad.
*   **Environment Variables:** 12-Factor App methodology.
*   **Centralized Config Server:** Consul, Etcd, Spring Cloud Config.
    *   Services fetch config on startup.
    *   Supports dynamic updates (change config without restart).
*   **Kubernetes Secrets/ConfigMaps:** Native K8s way to inject configs.

### Question 90: What is chaos engineering?

**Answer:**
The discipline of experimenting on a system to build confidence in its capability to withstand turbulent conditions.
*   **Practice:** Intentionally injecting failures (kill pods, add latency, partition network) in production or staging.
*   **Tool:** Chaos Monkey (Netflix).
*   **Goal:** Verify that the system recovers automatically (resilience).

---

## 🟢 Miscellaneous & Advanced Topics (Questions 91-100)

### Question 91: What is a web socket? How is it different from HTTP?

**Answer:**
*   **HTTP:** Request/Response. Client asks, Server answers. Connection closes (stateless). Unidirectional (mostly).
*   **WebSocket:** Full-duplex communication channel over a single TCP connection.
    *   **Persistent:** Connection stays open.
    *   **Bidirectional:** Server can push data to Client anytime.
    *   **Use Case:** Chat apps, Real-time feeds, Multiplayer games.

### Question 92: How do you design rate limiting?

**Answer:**
(Partial overlap with Q65, focusing on design here).
*   **Where:** API Gateway or Sidecar.
*   **Storage:** Redis (fast increment/read).
*   **Key:** API Key, IP Address, or UserID.
*   **Algorithm:** Sliding Window Log (most accurate).
*   **Response:** HTTP 429 + `Retry-After` header.

### Question 93: How to handle distributed transactions?

**Answer:**
*   **Two-Phase Commit (2PC):**
    *   Prepare Phase: Coordinator asks everyone "Can you commit?"
    *   Commit Phase: If all say yes, Coordinator says "Commit".
    *   *Problem:* Blocking, single point of failure.
*   **Saga Pattern (Preferred):**
    *   Sequence of local transactions.
    *   If step fails, trigger Compensating Transactions (Undos).
    *   Non-blocking, eventual consistency.

### Question 94: What is the role of Zookeeper in distributed systems?

**Answer:**
Apache ZooKeeper is a centralized service for maintaining configuration information, naming, providing distributed synchronization, and group services.
*   **Uses:**
    *   **Leader Election:** Deciding who is the Master.
    *   **Service Discovery:** Keeping list of live nodes.
    *   **Distributed Locks:** Ensuring only one process does a task.
*   **Note:** Being replaced by etcd in modern stacks (like K8s).

### Question 95: What is eventual vs strong consistency?

**Answer:**
*   **Strong Consistency:** After a write, ANY subsequent read returns the new value. (e.g., RDBMS). Easier to program, scales poorly (latency).
*   **Eventual Consistency:** After a write, reads *might* return old value for a while. Eventually, all nodes sync up. (e.g., DNS, Cassandra). Scales well, harder to debug.

### Question 96: What is leader election?

**Answer:**
In a cluster of nodes, one node is often designated as the "Leader" (or Master) to coordinate tasks or handle writes.
*   **Process:** Nodes talk to each other to vote/decide who is leader.
*   **Algorithms:** Bully Algorithm, Raft, Paxos.
*   **Scenario:** If Leader dies, a new election is triggered to pick a new Leader.

### Question 97: What is CRDT?

**Answer:**
Conflict-free Replicated Data Type.
*   **Concept:** Data structures that can be replicated across multiple computers, updated independently/concurrently, and resolve inconsistencies mathematically without coordination.
*   **Use Case:** Collaborative editing (Google Docs), Offline sync.
*   **Examples:** G-Counter (Grow-only counter), LWW-Element-Set (Last-Write-Wins).

### Question 98: Explain Raft or Paxos consensus algorithms.

**Answer:**
Algorithms to get a distributed system to agree on a single value (consensus).
*   **Paxos:** The original, proven algorithm. Very complex, hard to implement.
*   **Raft:** Designed to be understandable.
    *   Uses Leader Election.
    *   Log Replication (Leader sends logs to Followers).
    *   Safety (Committed only if majority acknowledge).
    *   Used in Etcd, Consul.

### Question 99: How to design a cron job scheduler?

**Answer:**
*   **Requirement:** Run tasks at specific times reliably.
*   **Single Node:** Linux cron (SPOF).
*   **Distributed:**
    *   **Master-Worker:** Master checks schedule, pushes task to Queue. Workers consume Queue.
    *   **Leader Election:** Use Redis/Zookeeper to pick one node to check schedule and dispatch tasks.
    *   **Deduplication:** Ensure task runs only once (Idempotency).

### Question 100: What is the difference between throughput and latency?

**Answer:**
*   **Latency:** Time taken to process a *single* request. (Measured in ms). "How fast is it?"
*   **Throughput:** Number of requests processed per unit of time. (Measured in RPS / TPS). "How much can it handle?"
*   **Analogy:**
    *   Latency = Travel time of one car.
    *   Throughput = Number of cars passing the bridge per hour.
*   *Optimization:* Pipelining increases throughput but often increases individual latency.
