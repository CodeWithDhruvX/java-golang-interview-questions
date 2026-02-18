# System Design Interview Questions (71-80)

## Scalability & Distributed Systems

### 71. How do you design a scalable REST service?
"Scalability starts with **Statelessness**.

The server shouldn't store client context (like 'User is on page 2') in memory. If I have 3 servers, a request might hit Server A, and the next one hits Server B. If state is on A, B fails.

So, I use tokens (JWT) for auth and store any session data in an external store like Redis.

Then, I put a **Load Balancer** in front to distribute traffic. I decouple heavy processing using **Message Queues** (Kafka/RabbitMQ) so the API stays responsive. And I use **Caching** (CDN for static assets, Redis for data) to reduce DB load."

### 72. How does load balancing work?
"A Load Balancer (LB) sits between the client and your servers. It acts as a traffic cop.

It accepts the incoming request and forwards it to one of the healthy backend servers based on an algorithm—usually **Round Robin** (take turns), or **Least Connections** (send to the server doing the least work).

It also does **Health Checks**. If Server 3 crashes, the LB detects it (via a /health endpoint) and stops sending traffic there until it recovers. This ensures high availability."

### 73. What is horizontal vs vertical scaling?
"**Vertical Scaling (Scaling Up)** is making the single server stronger: adding more RAM, more CPU. It's the easiest way to solve a problem—'just upgrade the AWS instance type'. But it has a hard limit (you can't add infinite RAM) and is a single point of failure.

**Horizontal Scaling (Scaling Out)** is adding more servers. Instead of one super-computer, you have 10 commodity servers. This is harder to implement (you need load balancers, stateless apps, distributed DBs), but the ceiling is virtually infinite. You can just keep adding nodes."

### 74. How would you design a URL shortener?
"This is a classic. The core is a mapping between a `long_url` and a `short_code`.

I'd use a SQL database (PostgreSQL) or NoSQL (DynamoDB) to store `{id, long_url, short_code, created_at}`.

The tricky part is generating the `short_code`. I wouldn't use a hash function (too long). I’d use **Base62 encoding** (A-Z, a-z, 0-9) on a unique ID.
But distributing unique IDs is hard. I’d use a dedicated ID Generator service (like Snowflake) or a pre-generation service (Token Service) to hand out ranges of IDs to servers.

For reads, since URLs rarely change, I’d cache the redirects heavily in Redis. 99% of traffic would hit the cache, not the DB."

### 75. How would you design a rate limiter?
"I'd implement this using Redis because it’s fast and atomic.

A popular algorithm is the **Token Bucket**. Every user gets a 'bucket' that fills with tokens at a constant rate. Each request costs a token. If the bucket is empty, request rejected (429).

Alternatively, the **Sliding Window** algorithm interacts better with burst traffic. I'd store a sorted set (ZSET) in Redis for each user, containing timestamps of their last N requests. When a new request comes, I remove timestamps older than the window (e.g., 1 minute) and count the remaining. If count < limit, allow.

For a global distributed limiter, I'd use Redis + Lua scripts to ensure atomicity."

### 76. Where would you use caching?
"I use caching at every layer where data is read frequently but changes rarely.

1.  **Browser/CDN**: Cache static assets (JS, CSS, Images). This is 'free' speed.
2.  **API Gateway / Reverse Proxy**: Cache entire API responses for public, static data (e.g., 'List of Countries').
3.  **Application / Service**: Internal object caching.
4.  **Database**: Buffer pool (DB does this automatically).

The most common specific use case is caching expensive database queries or results of complex calculations."

### 77. Redis vs in-memory cache?
"**In-memory cache** (like Caffeine or a simple HashMap) lives inside your JVM heap. It’s the fastest possible cache (nanoseconds). But:
1.  It consumes your application's RAM.
2.  It’s local. Server A doesn't seeServer B's cache. You might have data inconsistency.
3.  It dies when the app restarts.

**Distributed cache** (Redis) lives on a separate server. It’s slightly slower (network call ~1ms), but:
1.  It’s shared across all instances. If Server A updates the cache, Server B sees it.
2.  It persists (can survive app restarts).

I use in-memory for tiny, static reference data (like country codes). I use Redis for almost everything else."

### 78. Database vs cache consistency?
"This is the hardest problem in computer science.

There are three main patterns:

1.  **Cache-Aside (Lazy Loading)**: Read from cache. Miss? Read DB -> Write Cache. Update? Update DB -> *Delete* Cache. Next read re-populates it. This is the safest and most common.
2.  **Write-Through**: Write to Cache and DB synchronously. Strong consistency, but higher write latency.
3.  **Write-Back**: Write only to Cache. Cache async writes to DB. Super fast, but risky—if Cache dies, data is lost.

I almost always stick to **Cache-Aside** with a **TTL (Time To Live)** as a safety net. Even if code fails to invalidate cache, the TTL ensures it eventually corrects itself."

### 79. Microservices pros and cons?
"**Pros**:
-   **Independent Scaling**: Scale the 'Checkout' service without scaling 'User Profile'.
-   **Technology Agnostic**: Use Java for backend, Python for ML, Node for I/O.
-   **Fault Isolation**: If 'Recommendations' crash, 'Checkout' still works.
-   **Faster Deployments**: Compile and deploy just one small service.

**Cons**:
-   **Complexity**: Distributed systems are hard. Network failures, latency, eventual consistency.
-   **Operational Overhead**: You need robust monitoring, logging, and deployment pipelines (Kubernetes).
-   **Data consistency**: No ACID transactions across services (Sagas are complex).

It’s not a silver bullet; it’s a trade-off."

### 80. Monolith vs microservices — when to choose what?
"Start with a Monolith. Always.

A **Modular Monolith** is faster to develop, easier to test, and trivially easy to deploy. You keep domains separate via packages/modules, not network calls.

Move to Microservices only when:
1.  **Team Scaling**: You have 50+ devs stepping on each other's toes in the same repo.
2.  **Tech Requirements**: One part needs massive CPU scaling (video encoding) while another needs huge memory (in-memory analytics).
3.  **Fault Tolerance**: One feature is critical (Payments) and must not be taken down by a memory leak in a non-critical feature (Comments).

Don't do microservices just because Netflix does."
