# 🟢 Basic System Design Concepts — Questions 1–10

> **Level:** 🟢 Junior (0–2 yrs)
> **Asked at:** Service companies (TCS, Infosys, Wipro, Capgemini) as core theory, and as warm-up rounds at product companies (Amazon, Flipkart, Swiggy).

---

### 1. What is system design?
"System design is the process of defining the **architecture, components, modules, interfaces, and data flow** of a system to satisfy given functional and non-functional requirements.

When an interviewer asks me to 'design YouTube', they're not asking me to write code. They want to know how I think about the big picture — what databases I'd use, how I'd handle millions of concurrent users, how I'd store and stream video files.

It sits at the intersection of architecture, distributed systems, and product thinking. I break it into two layers: **High-Level Design (HLD)** — the bird's eye view of services, databases, queues, and CDNs — and **Low-Level Design (LLD)** — the detailed design of classes, APIs, and DB schemas."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys, Wipro (as theory) | Amazon, Flipkart (as intro to deeper rounds)

#### Indepth
System design interviews assess your ability to handle ambiguity. The key framework is:
1. **Clarify requirements** — Ask: DAU? Read-heavy or write-heavy? Consistency critical?
2. **Estimate scale** — Storage, QPS, bandwidth back-of-envelope calculations
3. **Define APIs** — What are the endpoints?
4. **Core components** — DB choice, caching, queuing, CDN
5. **Deep dive** — Interviewer will drill into one component
Always start with requirements clarification — jumping to a solution is the #1 mistake.

---

### 2. Difference between high-level and low-level design.
"High-Level Design (HLD) is the **macro view** — it answers 'what services exist and how they talk to each other'. Low-Level Design (LLD) is the **micro view** — it answers 'how does a specific service work internally'.

For example, if I'm designing WhatsApp: HLD shows User Service, Message Service, WebSocket Gateway, Kafka queues, and Cassandra DB on a diagram. LLD then zooms into the Message Service — showing the class diagram, the `Message` schema, and the exact API contract.

In interviews at product companies, HLD is enough for a 45-minute session. LLD questions come in separate coding rounds or design rounds focused on 'Design a Parking Lot'."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** Infosys, Wipro (as direct theory) | Myntra, Ajio (as part of design rounds)

#### Indepth
| Aspect | HLD | LLD |
|---|---|---|
| Focus | Architecture, system topology | Class diagrams, algorithm, DB schema |
| Audience | Architects, PMs, Senior Devs | Developers, Testers |
| Outcome | Component diagram, tech stack | Class/Sequence diagrams, pseudo-code |
| Example | "Use Kafka between Order & Payment service" | "OrderService has method `placeOrder(userId, items)` which calls `inventoryService.reserve()` then `paymentService.charge()`" |

HLD decisions directly impact scalability. LLD decisions directly impact maintainability and code quality.

---

### 3. What is scalability? Types of scalability?
"Scalability is the system's ability to **handle growing load** — more users, more data, more requests — without a proportional degradation in performance.

There are two types: **Vertical scaling** (scale up) — adding more CPU, RAM, or disk to a single machine. I did this for an internal dashboard that started getting more traffic; it worked quickly but has a hard ceiling. **Horizontal scaling** (scale out) — adding more machines. This is what Google, Amazon and every major platform does. It requires stateless services and a load balancer in front.

I always ask: 'Can this service scale horizontally?' If it stores state locally (sessions, files), it can't — and that's a design smell."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** Amazon (SDE-1, SDE-2), Flipkart, Zomato, Swiggy

#### Indepth
Beyond vertical and horizontal, there are two more types:
- **Diagonal scaling:** A mix — scale up first, then scale out. Often the most cost-efficient real-world approach.
- **Database scalability** is the hardest part. Stateless app servers scale easily behind a load balancer. The DB is usually the bottleneck. Techniques: read replicas, sharding, caching.

Key metric to quantify scalability: **QPS (Queries Per Second)**. A well-designed system should be able to scale from 1K QPS to 1M QPS by adding machines — with no code changes.

---

### 4. What is a load balancer?
"A load balancer is a server or software that **distributes incoming traffic across multiple backend servers** to ensure no single server is overwhelmed.

Think of it like a traffic cop at a highway interchange. Without it, every car (request) would pile into one lane (server) until it crashes. With it, traffic flows evenly.

I use Nginx as a reverse proxy/load balancer for internal services. For production cloud deployments, I use AWS ALB (Application Load Balancer) which operates at Layer 7 — it can route `/api` to one set of servers and `/static` to another. This is something Netflix and Amazon use heavily to separate their microservices traffic."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** Infosys, Accenture, TCS (as theory) | Amazon, Meesho, Razorpay (as part of design)

#### Indepth
Load balancers operate at two OSI layers:
- **L4 (Transport Layer):** Routes based on IP + Port. Very fast. Can't see HTTP headers. Example: AWS NLB.
- **L7 (Application Layer):** Routes based on URL path, HTTP headers, cookies. Smarter but slightly more CPU-intensive. Example: AWS ALB, Nginx.

Key algorithms: **Round Robin** (default), **Least Connections** (best for long-lived connections like WebSocket), **IP Hash** (sticky sessions), **Weighted** (route more to powerful servers).

The LB itself must be HA — deployed in Active-Passive or Active-Active pairs using VRRP/Keepalived to avoid a Single Point of Failure (SPOF).

---

### 5. What is caching? Where can it be applied?
"Caching is storing a **copy of data in a fast storage layer** so future requests for the same data can be served faster — avoiding a slow re-computation or DB query.

The classic example: a product page on Amazon. If 10,000 users request `/product/iphone-15`, there's no point hitting the DB 10,000 times. Cache the response once, serve it from Redis for the next 10 minutes.

I've used caching at multiple layers: browser cache for static assets, Redis at the application layer for session management and API response caching, and a CDN (Cloudflare) for geographic content distribution. The closer the cache is to the user, the faster the response."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** Flipkart, Amazon, Google, Swiggy, Netflix

#### Indepth
Caching layers in a typical web architecture (outermost to innermost):
1. **Browser** — `Cache-Control` headers control TTL
2. **CDN** — Nearest PoP serves static files
3. **Reverse Proxy** — Nginx stores cached responses
4. **Application Layer** — In-process (LRU Map) or distributed (Redis)
5. **DB Query Cache** — MySQL's internal query cache (deprecated in 8.0)
6. **CPU Cache** — L1/L2/L3 (hardware level)

Cache **hit rate** is the primary metric. A 99% hit rate means only 1% of requests touch the DB. LinkedIn maintains ~99% hit rates for profile data through aggressive caching.

---

### 6. What is CDN and how does it work?
"A CDN (Content Delivery Network) is a **globally distributed network of servers** that caches static content close to the end user to reduce latency.

When you open Netflix in Mumbai, the video data doesn't come from a server in the US — it comes from the nearest CDN edge node, maybe in Chennai or Pune. That's why Netflix streams at high quality even in India.

The process: User makes a request → DNS resolves to nearest CDN edge (PoP) → If the edge has it cached, serve it instantly → If not (cache miss), fetch from origin server, cache it, then serve it. First user gets slightly slower response; all subsequent users are fast."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🔴 Senior | **Asked at:** Netflix, Hotstar, Amazon, Cloudflare, Akamai roles

#### Indepth
CDN components:
- **PoP (Point of Presence):** Edge data centers in major cities
- **Anycast Routing:** DNS returns the IP of the nearest PoP — this is how the routing "magic" happens
- **Push vs Pull CDN:**
  - *Pull:* CDN fetches from origin on first miss (lazy). Used by most.
  - *Push:* You proactively upload to CDN. Used for large, predictable assets (videos, software releases).

Cache invalidation on a CDN is hard. Common approaches: versioned URLs (`main.v2.js`), short TTL, API-based cache purge (Cloudflare API). Cache-busting via URL versioning is the most reliable pattern used by Google and Facebook.

---

### 7. What is a reverse proxy?
"A reverse proxy is a server that **sits in front of your backend servers**, intercepting all incoming requests and forwarding them to the appropriate backend.

From the client's perspective, they're talking to one server. Behind the scenes, that server could be routing to 50 different microservices. Nginx is the most popular reverse proxy.

It gives me several benefits in one place: load balancing, SSL termination (so backend doesn't deal with TLS overhead), compression (gzip), caching static responses, and security (I can block bad IPs at the Nginx layer without hitting the backend)."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** Infosys, Wipro, Accenture | Razorpay, PhonePe, Meesho

#### Indepth
**Forward proxy vs Reverse proxy:**
- **Forward Proxy:** Sits in front of *clients*. Client → Forward Proxy → Internet. Used for anonymity, filtering (corporate firewalls). Example: VPN, Squid.
- **Reverse Proxy:** Sits in front of *servers*. Client → Reverse Proxy → Backend. Used for LB, SSL termination, caching.

A reverse proxy can also perform **request transformation** (modifying headers before passing to backend, injecting auth tokens) and **API gateway** duties. Tools: Nginx, HAProxy, Traefik, AWS CloudFront (as a reverse proxy for APIs), Envoy (used in service meshes like Istio).

---

### 8. What is a message queue?
"A message queue is an **asynchronous communication mechanism** that decouples the producer of a message from the consumer.

The classic example: A user places an order on Flipkart. The Payment service doesn't need to wait for the Notification service to send an email before confirming the order. Instead, it drops a message `{ orderId, userId, amount }` into a queue. The Notification service picks it up in its own time and sends the email.

This decoupling means services can fail independently, scale independently, and be deployed independently. I use RabbitMQ for simpler task queues and Kafka for high-throughput event streaming."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Zomato, Swiggy, Amazon, Uber, Razorpay

#### Indepth
Key properties to evaluate message queues on:
- **Durability:** Are messages persisted to disk if the broker crashes? (Kafka: yes. RabbitMQ: configurable)
- **Ordering:** Is delivery order guaranteed? (Kafka: within a partition. SQS: no guarantee in standard queues)
- **Delivery semantics:**
  - *At-most-once:* Message may be lost, never re-delivered
  - *At-least-once:* Message delivered ≥1 time, consumer must be idempotent (most common)
  - *Exactly-once:* Strongest guarantee, highest cost (Kafka transactions, not always needed)
- **Consumers:** Point-to-point (one consumer per message — SQS, RabbitMQ) vs Pub/Sub (multiple consumer groups — Kafka, Google Pub/Sub)

---

### 9. What is sharding?
"Sharding is **horizontal partitioning of data across multiple databases or nodes** so that each node holds only a subset of the total data.

Imagine you have a Users table with 500 million rows. A single DB server can't store or query that fast. With sharding, I split it: users with IDs 1–100M on Shard 1, 100M–200M on Shard 2, and so on. Each query goes to the specific shard that owns the data.

I choose the shard key very carefully — a bad shard key creates **hotspots** (one shard gets all the traffic). A good shard key distributes load evenly. At Uber, driver location data is sharded by geographic region."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Amazon, Flipkart, PhonePe, Google, Twitter/X

#### Indepth
Sharding strategies:
- **Range-based:** Shard by value range (userID 1–1M on S1). Simple but risks hotspots if new users cluster at the top.
- **Hash-based:** `shard = hash(userID) % num_shards`. Even distribution but resharding is painful (consistent hashing solves this).
- **Directory-based:** A lookup table maps keys to shards. Most flexible, but the lookup table is a SPOF.
- **Consistent Hashing:** Used by DynamoDB, Cassandra. Adding new shards only moves a fraction of data (not all of it).

Major challenges: Cross-shard JOINs are expensive, enforcing uniqueness across shards is hard, and distributed transactions across shards require the saga pattern. This is why sharding is a last resort after vertical scaling, read replicas, and caching are exhausted.

---

### 10. Difference between vertical and horizontal scaling.
"Vertical scaling (scale up) means **making one machine more powerful** — adding more CPU cores, RAM, or faster disks. Horizontal scaling (scale out) means **adding more machines** to share the load.

For a startup with sudden traffic growth, vertical scaling is my first instinct — it's fast and requires no code changes. Upgrade from a t2.medium to an m5.4xlarge on AWS and you're done. But there's a hard limit on how powerful one machine can be.

Horizontal scaling is the long-term answer. I make my services stateless, put an LB in front, and spin up as many instances as needed. This is how Amazon handles millions of requests per second on Prime Day — not one giant server, but hundreds of smaller ones."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Wipro, Infosys (as theory) | Amazon, Flipkart (in context of system design)

#### Indepth
| Feature | Vertical Scaling | Horizontal Scaling |
|---|---|---|
| Method | Add CPU/RAM/Disk to one node | Add more nodes |
| Complexity | Low — no code change | High — needs LB, stateless services, distributed data |
| Downtime | Requires restart for hardware upgrade | Zero downtime (add nodes live) |
| Cost | Exponential (enterprise hardware premium) | Linear (commodity servers) |
| Limit | Hard hardware ceiling (~192 cores, 24TB RAM) | Theoretically unlimited |
| Failure | Single point of failure | Fault tolerant |

**The Database exception:** Stateless app servers scale horizontally with zero effort. Databases are inherently stateful and much harder to scale horizontally — this is why DB optimization (indexing, caching, read replicas) is always done first. Only when those are exhausted should you shard.
