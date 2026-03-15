## 🟢 Scalability and Availability (Questions 51-60)

### Question 51: How do you design a system that handles millions of users?

**Answer:**
Designing for millions of users requires a distributed architecture focusing on horizontal scaling and decoupling.
1.  **Load Balancing:** Distribute traffic across multiple application servers.
2.  **Database Scaling:** Use Master-Slave replication for reads, Sharding for writes, or NoSQL for massive scale.
3.  **Caching:** Aggressively cache at all layers (CDN, Redis, Browser) to offload the DB.
4.  **Asynchronous Processing:** Use Message Queues (Kafka/RabbitMQ) for non-critical tasks (emails, report generation).
5.  **Microservices:** Break monolithic apps into smaller services to scale independently.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you design a system that handles millions of users?

**Your Response:** "Designing for millions of users requires a distributed architecture from the start. I'd begin with load balancing to distribute traffic across multiple application servers. For the database, I'd use a combination of read replicas to handle the read-heavy workload and sharding for write scalability. Caching is crucial - I'd implement Redis at multiple levels to reduce database load. For non-critical operations like sending emails or generating reports, I'd use message queues like Kafka to process them asynchronously. Finally, I'd break the application into microservices so each component can scale independently based on its specific needs. The key is identifying bottlenecks and designing each layer to handle horizontal scaling."

### Question 52: How to scale a system read-heavy workload?

**Answer:**
*   **Caching:** The most effective strategy. Use Redis/Memcached to serve frequent queries from memory.
*   **Database Replication:** Add multiple Read Replicas (Slaves). Point read queries to slaves and writes to the master.
*   **CDN:** Serve static content (images, CSS) from the edge.
*   **Denormalization:** Structure database tables to avoid expensive joins during reads.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How to scale a system read-heavy workload?

**Your Response:** "For read-heavy systems, caching is my first line of defense. I'd implement Redis or Memcached to serve frequently accessed queries directly from memory. Then I'd add multiple read replicas to the database - all writes go to the master, but reads are distributed across several read-only copies. I'd also use a CDN to serve static content like images and CSS from edge locations closer to users. In the database itself, I might denormalize some tables to avoid complex joins that slow down reads. The goal is to minimize database hits by serving data from the fastest possible source at each layer."

### Question 53: How to scale a system write-heavy workload?

**Answer:**
Write-heavy systems are harder to scale than read-heavy ones.
*   **Sharding:** Distribute data across multiple database nodes based on a shard key (e.g., UserID).
*   **NoSQL:** Use write-optimized databases like Cassandra (LSM Trees) or DynamoDB.
*   **Async Writes (Write-Behind):** Write to a message queue first (Kafka) and process/persist to DB asynchronously.
*   **Bulk/Batch Inserts:** Group small writes into fewer large batches to reduce I/O overhead.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How to scale a system write-heavy workload?

**Your Response:** "Write-heavy systems are much harder to scale than read-heavy ones. My primary strategy would be sharding - distributing data across multiple database servers based on a shard key like user ID or geographic region. I'd also consider using NoSQL databases like Cassandra that are optimized for high write throughput with their log-structured storage. For some use cases, I'd implement write-behind caching where I write to a fast cache first and asynchronously persist to the database. I'd also batch small writes together into larger, more efficient database operations. The key is distributing the write load across multiple machines and optimizing the write path at every level."

### Question 54: What is replication and when to use it?

**Answer:**
Replication is keeping a copy of the same data on multiple machines.
*   **Uses:**
    *   **High Availability:** If one node fails, others can serve data.
    *   **Latency:** Replicate data to regions closer to users.
    *   **Read Scaling:** Distribute read traffic across replicas.
*   **Types:** Active-Passive (Master-Slave), Active-Active (Master-Master).

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is replication and when to use it?

**Your Response:** "Replication is creating and maintaining copies of the same data across multiple machines. I use it for three main reasons: high availability - if one server fails, others can serve the data; latency reduction - I can replicate data to different geographic regions so users access it from closer locations; and read scaling - I can distribute read queries across multiple replicas. The common patterns are active-passive where one server handles all traffic and others are on standby, and active-active where multiple servers share the load. The choice depends on whether I need failover capability, performance improvement, or both."

### Question 55: How to make a system fault-tolerant?

**Answer:**
Fault tolerance is the ability of a system to continue operating without interruption when one or more of its components fail.
*   **Redundancy:** Eliminate Single Points of Failure (SPOF) by having backup components (e.g., multiple server instances, DB replicas).
*   **Replication:** Replicate data across zones/regions.
*   **Circuit Breaker:** Fail fast and recover gracefully to prevent cascading failures.
*   **Health Checks & Auto-recovery:** Kubernetes restarts failed pods automatically.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How to make a system fault-tolerant?

**Your Response:** "Fault tolerance means the system keeps working even when components fail. I achieve this by eliminating single points of failure - having redundant components at every level. For servers, I'd run multiple instances behind a load balancer. For data, I'd replicate it across different availability zones or even regions. I'd implement circuit breakers so failures don't cascade through the system. And I'd use health checks with auto-recovery mechanisms like Kubernetes that automatically restart failed services. The goal is designing the system to expect and handle failures gracefully rather than trying to prevent them entirely."

### Question 56: What is failover?

**Answer:**
Failover is the automatic switching to a redundant or standby computer server, system, hardware component, or network upon the failure of the previously active application.
*   **Example:** If the Primary DB crashes, the system promotes a Read Replica to be the new Primary.
*   **Automation:** Usually handled by Load Balancers or Orchestrators (K8s) using heartbeat monitoring.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is failover?

**Your Response:** "Failover is the system's ability to automatically switch from a failed component to a backup one. For example, if my primary database server crashes, failover would automatically promote one of the read replicas to become the new primary so the application can continue working. This is usually handled by load balancers or orchestrators like Kubernetes that continuously monitor the health of components using heartbeats. When they detect a failure, they automatically redirect traffic to healthy backup components. The key is that this happens automatically and quickly so users don't even notice there was a problem."

### Question 57: What is high availability?

**Answer:**
High Availability (HA) refers to systems that are durable and likely to operate continuously without failure for a long time.
*   **Measurement:** "Nines" of availability.
    *   99.9% uptime = 8.76 hours downtime/year.
    *   99.999% (Five 9s) = 5.26 minutes downtime/year.
*   **Achieved via:** Redundancy, Load Balancing, Clustering, and Failover mechanisms.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is high availability?

**Your Response:** "High availability means the system is up and running most of the time, typically measured in 'nines'. For example, 99.9% uptime means the system can be down for about 8 hours per year, while 99.999% or 'five nines' means only about 5 minutes of downtime per year. I achieve this through redundancy - having multiple servers, databases, and network paths. I use load balancing to distribute traffic and failover mechanisms to switch to backup components when something fails. The specific availability target depends on the business requirements - financial systems might need five nines, while a blog might be fine with 99.9%."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** What's the difference between active-passive and active-active systems?

**Your Response:** "In active-passive, one server handles all the traffic while the backup sits idle waiting to take over if the primary fails. It's simpler to implement but wastes resources since the backup isn't being used. In active-active, all servers share the traffic load simultaneously, which gives me better resource utilization and higher throughput. However, it's more complex because I need to handle data synchronization to avoid conflicts when multiple servers can write to the same data. I'd choose active-passive for simpler failover scenarios and active-active when I need maximum performance and can handle the complexity."

### Question 59: What is graceful degradation?

**Answer:**
Graceful degradation allows a system to maintain limited functionality even when a large portion of it has been destroyed or is inoperative.
*   **Example:** If the "Recommendations" service fails on an E-commerce site, the main page still loads "Products," but the "Recommended for You" section is empty or hidden, rather than showing a 500 Error Page.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is graceful degradation?

**Your Response:** "Graceful degradation is designing the system to provide reduced functionality rather than complete failure when something goes wrong. For example, if the recommendation service on an e-commerce site fails, instead of showing users a 500 error page, I'd still show the main product listings but leave the recommendations section empty or hide it. The core functionality still works, just with fewer features. This is much better than a complete outage because users can still accomplish their main goals. I implement this by identifying non-critical features and designing fallbacks for when they're unavailable."

### Question 60: What is a throttling mechanism?

**Answer:**
Throttling is the process of limiting the number of actions a user or component can perform in a given time frame to prevent overuse of resources.
*   **Purpose:** Protects the system from becoming unresponsive due to high load (DoS attacks or noisy neighbors).
*   **Implementation:** Token Bucket algorithm, Leaky Bucket.
*   **Result:** Requests exceeding the limit are rejected (HTTP 429 Too Many Requests).

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is a throttling mechanism?

**Your Response:** "Throttling is how I protect the system from being overwhelmed by too many requests. I limit how many requests a user or service can make in a given time period - for example, 100 requests per minute per IP address. This prevents both accidental overload and malicious attacks like denial-of-service. When the limit is exceeded, I return an HTTP 429 'Too Many Requests' error. I typically implement this using algorithms like token bucket or leaky bucket, and I'd place it at the API gateway level. It's a crucial protection mechanism for any public-facing API."

---

## 🟢 Security & Authentication (Questions 61-70)

### Question 61: How to design a secure login system?

**Answer:**
1.  **HTTPS:** Encrypt all traffic using SSL/TLS.
2.  **Hashing:** Never store plain text passwords. Use bcrypt, Argon2, or SCRAM with a unique salt per user.
3.  **MFA:** Implement Multi-Factor Authentication (SMS/TOTP).
4.  **Session Management:** Use secure, HTTPOnly, Secure cookies for session IDs to prevent XSS.
5.  **Rate Limiting:** Prevent Brute Force attacks by limiting failed login attempts.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you design a secure login system?

**Your Response:** "For a secure login system, I'd start with HTTPS to encrypt all traffic between the client and server. For passwords, I'd never store them in plain text - I'd use strong hashing algorithms like bcrypt or Argon2 with a unique salt for each user. I'd implement multi-factor authentication to add an extra layer of security beyond just passwords. For session management, I'd use secure, HTTP-only cookies to prevent XSS attacks. Finally, I'd add rate limiting to prevent brute force attacks by limiting failed login attempts per IP address. The goal is defense in depth - multiple layers of security so if one fails, others still protect the system."

### Question 62: What is OAuth 2.0?

**Answer:**
The industry-standard protocol for authorization.
*   **Concept:** Allows a user to grant a third-party application access to their resources on another service (e.g., "Log in with Google") without sharing their password.
*   **Roles:** Resource Owner (User), Client (App), Authorization Server (Google), Resource Server (API).
*   **Flow:** User authenticates with Auth Server -> Auth Server issues Access Token -> Client uses Token to access Resource Server.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is OAuth 2.0?

**Your Response:** "OAuth 2.0 is the industry standard for authorization that lets users grant third-party applications access to their data without sharing passwords. Think of 'Login with Google' - you're not giving your Google password to the app, but Google gives the app a temporary access token with limited permissions. There are four main players: the resource owner (that's you, the user), the client (the app you're using), the authorization server (like Google), and the resource server (the API with your data). The flow is simple: you authenticate with the authorization server, it gives you an access token, and the client uses that token to access your data from the resource server."

### Question 63: What is JWT?

**Answer:**
JSON Web Token (JWT) is a compact, URL-safe means of representing claims to be transferred between two parties.
*   **Structure:** `Header.Payload.Signature`
*   **Stateless:** The server doesn't need to store session data; it verifies the signature to trust the token.
*   **Usage:** Authentication (Logged in user), Information Exchange.
*   **Security:** Always use HTTPS. Don't put sensitive secrets in the Payload (it's only Base64 encoded, not encrypted).

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is JWT?

**Your Response:** "JWT is a compact, self-contained token that carries user information between services. It has three parts: a header, a payload with the user claims, and a signature that proves the token hasn't been tampered with. The beauty of JWT is that it's stateless - the server doesn't need to store session information because it can verify the token's signature. I use JWT for authentication and passing user information between microservices. But I'm careful not to put sensitive data in the payload since it's just base64 encoded, not encrypted, and I always use HTTPS to protect the token in transit."

### Question 64: How do you store passwords securely?

**Answer:**
1.  **Hashing:** One-way transformation (cannot be reversed).
2.  **Salting:** Add a unique random string to each password before hashing to defeat Rainbow Table attacks.
3.  **Algorithm:** Use slow algorithms like **bcrypt**, **scrypt**, or **Argon2** to make brute-force attacks computationally expensive.
4.  **Pepper:** (Optional) Add a secret key stored separately from the DB to the hash.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you store passwords securely?

**Your Response:** "I never store passwords in plain text. First, I hash them using a one-way algorithm so they can't be reversed. I add a unique random salt to each password before hashing to prevent rainbow table attacks. I use slow hashing algorithms like bcrypt or Argon2 that are computationally expensive, making brute force attacks impractical. Optionally, I might add a pepper - a secret key stored separately from the database. The key is that even if the database is compromised, attackers can't easily recover the original passwords because each password is uniquely salted and hashed with a slow algorithm."

### Question 65: What is rate limiting?

**Answer:**
A strategy for limiting network traffic. It sets a cap on how many requests a sender can issue in a specific time window.
*   **Algorithms:**
    *   **Token Bucket:** Tokens are added at a fixed rate; request consumes a token.
    *   **Leaky Bucket:** Requests enter a queue and are processed at a constant rate.
    *   **Fixed Window Counter:** Count requests per minute.
    *   **Sliding Window Log:** More accurate timestamp tracking.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is rate limiting?

**Your Response:** "Rate limiting is how I control the amount of traffic a user or service can send in a given time period. I implement it using different algorithms depending on the use case. The token bucket algorithm adds tokens at a fixed rate and each request consumes a token - this allows bursts but maintains an average rate. The leaky bucket processes requests at a constant rate, smoothing out traffic. For simpler cases, I might use a fixed window counter that resets every minute, or a sliding window log for more accurate tracking. The goal is preventing abuse and ensuring fair resource allocation among all users."

### Question 66: How to secure APIs?

**Answer:**
1.  **Authentication:** Identify the caller (API Keys, OAuth, JWT).
2.  **Authorization:** Ensure caller has permission (RBAC, Scopes).
3.  **Encryption:** Force HTTPS everywhere.
4.  **Input Validation:** Sanitize all inputs to prevent SQL Injection/XSS.
5.  **Rate Limiting:** Prevent DoS.
6.  **CORS:** Restrict which domains can call your API.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How to secure APIs?

**Your Response:** "Securing APIs requires multiple layers of protection. First, I implement authentication to verify who is calling the API - using API keys, OAuth tokens, or JWTs. Then I add authorization to ensure they only have access to what they're allowed to see. I enforce HTTPS everywhere to encrypt all traffic. I validate and sanitize all inputs to prevent injection attacks. I add rate limiting to prevent denial of service attacks. And I configure CORS to restrict which domains can make cross-origin requests. Security is about defense in depth - if one layer fails, others still protect the API."

### Question 67: What is CORS?

**Answer:**
Cross-Origin Resource Sharing (CORS) is a browser security mechanism that restricts cross-origin HTTP requests.
*   **Mechanism:** When JS on `domainA.com` calls API on `domainB.com`, the browser sends a pre-flight `OPTIONS` request.
*   **Headers:** Server must respond with `Access-Control-Allow-Origin: domainA.com` (or `*`) for the browser to allow the actual request.
*   **Goal:** Prevents malicious sites from reading data from other sites where the user is logged in.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is CORS?

**Your Response:** "CORS is a browser security feature that prevents malicious websites from making requests to other websites on your behalf. When JavaScript on domainA.com tries to call an API on domainB.com, the browser first sends a pre-flight OPTIONS request to check if the cross-origin request is allowed. The server must respond with the appropriate CORS headers, specifically Access-Control-Allow-Origin, listing which domains are permitted. This prevents a malicious site from reading your bank account data just because you're logged into your bank in another tab. It's the browser enforcing the same-origin policy to protect user privacy and security."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** Explain SSL/TLS in web communication.

**Your Response:** "SSL/TLS is what creates the secure HTTPS connection you see in your browser. It starts with a handshake where the client and server negotiate encryption capabilities. The server sends its SSL certificate containing its public key, which the client verifies with a trusted certificate authority. The client then generates a symmetric session key, encrypts it with the server's public key, and sends it over. Only the server can decrypt this with its private key. Now both parties have the same symmetric key and can communicate securely. The beauty is that asymmetric encryption is only used briefly to exchange the symmetric key, then faster symmetric encryption handles the actual data transfer."

### Question 69: What is cross-site request forgery (CSRF)?

**Answer:**
An attack that forces an end user to execute unwanted actions on a web application in which they are currently authenticated.
*   **Example:** Malicious site has a hidden form that POSTs to `bank.com/transfer`. If user is logged into bank, the browser sends session cookies, and the transfer happens.
*   **Prevention:** Use **CSRF Tokens** (random values injected into forms/headers) that the malicious site cannot guess.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is cross-site request forgery (CSRF)?

**Your Response:** "CSRF is an attack where a malicious website tricks your browser into making unwanted requests to a site where you're authenticated. For example, if you're logged into your bank and visit a malicious site, it could have a hidden form that automatically submits a transfer request to your bank. Since your browser sends your bank's cookies with the request, the bank thinks it's you making the transfer. I prevent this by implementing CSRF tokens - random values that are included in forms and verified on the server. Since the malicious site can't read or guess these tokens, it can't forge valid requests."

### Question 70: What is cross-site scripting (XSS)?

**Answer:**
An attack where malicious scripts are injected into trusted websites.
*   **Reflected XSS:** Script is in the URL (e.g., search query) and executed when victim clicks link.
*   **Stored XSS:** Script is saved in DB (e.g., comment section) and executed on every visitor's browser.
*   **Prevention:**
    *   **Sanitize HTML:** Escape special characters (`<` becomes `&lt;`).
    *   **CSP (Content Security Policy):** Restrict sources of executable scripts.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is cross-site scripting (XSS)?

**Your Response:** "XSS is when attackers inject malicious scripts into web pages that other users trust. There are two main types: reflected XSS, where the script comes from the URL and executes immediately when someone clicks a malicious link; and stored XSS, where the script is saved in the database and runs on every visitor's browser. I prevent XSS by sanitizing all user input - escaping special characters like < and > so they're treated as text, not HTML. I also implement Content Security Policy headers that restrict which domains can execute scripts. The key is never trusting user input and always treating it as potentially malicious."

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
