# High-Level System Design (HLD) for Product-Based Companies

For SDE-2 (3+ yoe) and SDE-3 roles, High-Level System Design is often the deciding factor. You are expected to draw architecture diagrams, discuss trade-offs, and design for scale, availability, and reliability.

---

## 🏗️ 1. Core Concepts & Building Blocks

Before attempting a full design, you must understand these foundational concepts.

### Scalability (Vertical vs. Horizontal)
*   **Vertical Scaling (Scale-Up):** Adding more CPU/RAM to a single server. Limited by hardware and has a single point of failure (SPOF).
*   **Horizontal Scaling (Scale-Out):** Adding more servers to a cluster. Requires a Load Balancer and stateless applications. This is the modern standard.

### CAP Theorem
It states that a distributed system can only guarantee two out of three characteristics:
1.  **Consistency (C):** Every read receives the most recent write or an error.
2.  **Availability (A):** Every request receives a (non-error) response, without the guarantee that it contains the most recent write.
3.  **Partition Tolerance (P):** The system continues to operate despite an arbitrary number of messages being dropped by the network.

*Note: In the real world, network partitions (P) are unavoidable, so you must choose between CP (e.g., MongoDB, HBase) or AP (e.g., Cassandra, DynamoDB).*

### Load Balancing
Distributes incoming network traffic across multiple servers.
*   **Layer 4 (Transport):** Routes traffic based on IP and Port (e.g., AWS Network Load Balancer). Fast, but less intelligent.
*   **Layer 7 (Application):** Routes traffic based on HTTP headers, URLs, or cookies (e.g., Nginx, AWS ALB). Slower but smarter.
*   **Algorithms:** Round Robin, Least Connections, IP Hash.

### Caching
Storing frequently accessed data in memory (Redis, Memcached) to reduce database load.
*   **Strategies:** Read-Through, Write-Through, Write-Behind (Write-Back), Cache Aside.
*   **Eviction Policies:** LRU (Least Recently Used), LFU (Least Frequently Used), FIFO.

### Databases
*   **Relational (SQL):** MySQL, PostgreSQL. Use when ACID properties and complex relationships/joins are required (e.g., banking systems).
*   **NoSQL (Document):** MongoDB. Use for flexible schemas and rapid development.
*   **NoSQL (Key-Value):** Redis, DynamoDB. Use for ultra-fast lookups (e.g., user sessions, caching).
*   **NoSQL (Wide-Column):** Cassandra, HBase. Use for massive write-heavy workloads (e.g., time-series data, chat apps).
*   **Replication vs. Sharding:** Replication copies data for High Availability and Read-Scaling. Sharding partitions data across multiple servers for Write-Scaling and massive datasets.

### Message Queues & Event Streaming
Decoupling services and handling asynchronous processing.
*   **Message Queues (RabbitMQ, SQS):** Good for point-to-point, task-based workloads. Messages are deleted after processing.
*   **Event Streams (Kafka, Kinesis):** Good for publish-subscribe, high-throughput, replayable event logs. Messages persist for a defined retention period.

---

## 🛠️ 2. The 5-Step System Design Framework

When asked "Design Facebook," follow this structured approach:

**Step 1: Understand the Goal & Scope (5-10 mins)**
*   Clarify functional requirements (What should the system do? E.g., Users can post tweets, follow others, view feed).
*   Clarify non-functional requirements (Is it read-heavy or write-heavy? Highly available or strictly consistent? Acceptable latency?).
*   Define out-of-scope features to keep the design manageable.

**Step 2: Capacity Estimation (Back-of-the-envelope calculations) (5 mins)**
*   Estimate Daily Active Users (DAU), read/write ratio.
*   Estimate Storage requirements (per day/year).
*   Estimate Network Bandwidth (Requests per second - RPS).
*   *(Tip: Don't get bogged down in math; keep it simple and state your assumptions).*

**Step 3: High-Level API Design (5 mins)**
*   Define the core REST or gRPC endpoints.
*   `POST /v1/tweets (userId, content, mediaIds)`
*   `GET /v1/feed (userId, cursor, limit)`

**Step 4: High-Level Architecture Diagram (10-15 mins)**
*   Draw the components: DNS, CDN, Load Balancer, API Gateway, Microservices, Caching Layer, Database, Message Queues.
*   Show the main data flow for the core use cases.

**Step 5: Deep Dive & Trade-offs (10-15 mins)**
*   The interviewer will drill into specific bottlenecks (e.g., "How do you handle a celebrity with 50M followers posting a tweet?").
*   Discuss Database Schema (SQL vs NoSQL).
*   Discuss Caching strategies.
*   Discuss single points of failure (SPOF) and how to mitigate them.

---

## 🚀 3. Top 10 Frequently Asked System Design Questions

### 1. Design a URL Shortener (TinyURL)
*   **Focus:** Generating unique IDs (Base62 encoding, Snowflake ID, or distributed counters), high read throughput (caching), and fast redirects.

### 2. Design WhatsApp / Facebook Messenger
*   **Focus:** Real-time bidirectional communication (WebSockets, Server-Sent Events), presence service (online/offline status), read receipts, and message persistence (Cassandra).

### 3. Design Twitter / News Feed System
*   **Focus:** Fan-out problem.
*   **Fan-out on write (Push):** Pre-compute the feed for active users when a tweet is posted (Good for normal users).
*   **Fan-out on read (Pull):** Compute on the fly when the user logs in (Good for celebrities like Elon Musk to avoid overwhelming the system). Usually requires a hybrid approach.

### 4. Design Uber / Lyft
*   **Focus:** Location tracking, geospatial indexing (QuadTrees, Geohashing), dispatch system matching riders and drivers in real-time.

### 5. Design Netflix / YouTube
*   **Focus:** Content Delivery Networks (CDNs), video chunking/transcoding (adaptive bitrate streaming), metadata storage, and recommendation engines.

### 6. Design an E-commerce Checkout System (Amazon)
*   **Focus:** Data consistency (Saga Pattern/Two-Phase Commit), Idempotency (preventing double charges), and dealing with massive traffic spikes (Black Friday sales) using queues.

### 7. Design a Ticket Booking System (BookMyShow)
*   **Focus:** Concurrency handling and locking mechanisms (Pessimistic vs. Optimistic locking) to ensure two users don't book the same seat simultaneously.

### 8. Design a Rate Limiter
*   **Focus:** Algorithms (Token Bucket, Leaky Bucket, Sliding Window Log, Sliding Window Counter) and distributed implementation using Redis (e.g., Redis Lua scripts).

### 9. Design an Instagram / Photo Sharing App
*   **Focus:** Scalable object storage (AWS S3) for images, generating thumbnails (asynchronous processing via queues), and feed generation.

### 10. Design a Distributed Web Crawler
*   **Focus:** URL frontier (queueing URLs to visit), politeness policy (rate-limiting requests per domain), duplicate detection (Bloom Filters), and massive horizontal scaling.
