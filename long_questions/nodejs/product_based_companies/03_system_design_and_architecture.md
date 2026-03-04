# 🏛️ 03 — System Design & Architecture
> **Most Asked in Product-Based Companies** | 🏛️ Difficulty: Hard

---

## 🔑 Must-Know Topics
- Designing highly scalable Node.js applications
- API Gateway vs Reverse Proxy (Nginx, HAProxy)
- Message Brokers (RabbitMQ, Apache Kafka, Redis Pub/Sub)
- WebSockets and scalable real-time systems
- CAP Theorem and Eventual Consistency
- Database Sharding and Read Replicas

---

## ❓ Frequently Asked Questions

### Q1. How do you scale a real-time Node.js chat application capable of handling 1M+ concurrent WebSocket connections?

**Answer:**
A single Node.js instance (process) cannot handle 1M concurrent WebSocket connections due to port limitations (max ~65k) and eventual CPU/Memory limits.

**Architecture for scaling:**
1. **Load Balancing:** Use a Layer 7 Load Balancer (like HAProxy or Nginx) configured with sticky sessions (if using Long Polling fallback) or just pure WebSocket forwarding.
2. **Horizontal Scaling:** Deploy multiple instances of the Node.js application across different servers/containers (e.g., using Kubernetes HPA).
3. **The Pub/Sub Message Broker:** 
   - If User A connects to Instance 1, and User B connects to Instance 2, Instance 1 cannot directly emit a socket event to User B.
   - **Solution:** Use **Redis Pub/Sub**, Kafka, or RabbitMQ. All Node.js instances publish messages to the central broker. All instances also subscribe to the broker. 
   - When User A sends a message to User B, Instance 1 publishes to Redis. Redis broadcasts it to all instances. Instance 2 receives it and pushes it out to User B's socket.
4. **OS Optimization:** Increase file descriptor limits (`ulimit -n`) at the OS level, as each WebSocket connection opens a file descriptor.

---

### Q2. What is an API Gateway? How does it differ from a Load Balancer?

**Answer:**
- **Load Balancer:** Operates mostly at Layer 4 (Network) or Layer 7 (Application). Its primary job is to distribute incoming raw network traffic evenly across a group of backend servers to prevent overload.
- **API Gateway:** Operates strictly at Layer 7. It sits between the clients and a massive ecosystem of microservices. It acts as a "reverse proxy" with advanced features aimed specifically at managing APIs.

**Features of an API Gateway (e.g., Kong, AWS API Gateway):**
1. **Routing:** Maps a single endpoint `/api/billing` to the internal Billing microservice.
2. **Authentication/Authorization:** Validates JWTs centrally before forwarding requests to internal services.
3. **Rate Limiting & Throttling:** Prevents DDoS and API abuse per user/IP.
4. **Response Caching:** Caches GET requests.
5. **Protocol Translation:** Can translate REST/HTTP from the client to gRPC for the internal microservices.

---

### Q3. Explain the Outbox Pattern in Microservices. Why is it necessary?

**Answer:**
In a microservices architecture, you often need to **update a database** *and* **publish a message/event** to a message broker (like Kafka/RabbitMQ) at the same time. 

If you update the database and then the service crashes before publishing the message, the system enters an inconsistent state. Since you cannot wrap a database transaction and a network call to a broker in a single ACID transaction (Two-Phase Commit is slow and unsupported by many NoSQL DBs), we use the **Outbox Pattern**.

**How it works:**
1. Create a table/collection called `Outbox` in the same database where the primary entity is updated.
2. Start a local database transaction.
3. Update the primary entity (e.g., `Users` table).
4. Insert a record describing the event into the `Outbox` table.
5. Commit the transaction (guarantees Atomicity).
6. An asynchronous background worker (or a tool like Debezium using CDC) continuously polls the `Outbox` table.
7. The worker reads the event, publishes it to the Message Broker, and then deletes/marks the record as processed in the `Outbox` table.

---

### Q4. Describe the design of a URL Shortener (like Bitly) built with Node.js.

**Answer:**
A URL shortener is a read-heavy system.

1. **Unique ID Generation:** 
   - Generating a random 7-character string (Base62: A-Z, a-z, 0-9) allows for `62^7 = 3.5 trillion` URLs.
   - Use centralized ID generation like Twitter Snowflake or a ZooKeeper counter, or pre-generate ranges of IDs and distribute them to Node.js workers to avoid DB collisions on insert.

2. **Database:** 
   - A NoSQL database like MongoDB or Cassandra is excellent for high-volume, simple Key-Value lookups (`short_url_hash : long_url`).

3. **Node.js API:**
   - `POST /shorten` -> Checks cache. If exists, return. Else, get ID, save to DB, cache it.
   - `GET /:hash` -> High throughput required.

4. **Caching Layer (Crucial):**
   - Use **Redis** to cache recently accessed short URLs. 
   - In Node.js: Check Redis first -> Cache Hit (redirect 301/302). Cache Miss -> Check DB -> Update Redis -> Redirect.
   - Eviction policy: LRU (Least Recently Used).

5. **Redirect Types:**
   - **301 (Permanent):** Browser caches it forever. Good for server load, bad for analytics.
   - **302 (Temporary):** Hits the Node server every time. Good for tracking click analytics, higher server load.

---

### Q5. In a distributed system, how do you handle Node.js application state?

**Answer:**
A golden rule of scalable distributed systems is that **backend servers must be stateless**.

The Node.js application should not store session data, user context, or temporary file uploads in memory or on the local disk. If a server crashes or a load balancer directs the user's next request to a different server, that state would be lost.

**Handling State:**
1. **Session State (e.g., Auth):** Store sessions in a centralized, fast, in-memory datastore like **Redis** or Memcached. Alternatively, use stateless JWTs where the state is stored cryptographically on the client.
2. **Application State/Data:** Store persistent data in central Databases (SQL/NoSQL).
3. **File Uploads:** Stream file uploads directly to distributed object storage like **AWS S3** instead of saving them to the local server disk.
