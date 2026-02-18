# System Design & Microservices - Interview Answers

> 🎯 **Focus:** These answers move beyond code to architecture—how systems live in production.

### 1. Monolith vs Microservices?
"A **Monolith** is a single deployable unit. Everything—UI, Service, DB access—runs in one process. It’s easy to develop and deploy, but hard to scale independently.

**Microservices** break the app into small, independent services (User Service, Order Service) that communicate via APIs.
The pro is scalability—I can scale the Order Service during Black Friday without touching the User Service.
The con is complexity. You now have to deal with distributed tracing, network failures, and data consistency."

---

### 2. What is CAP Theorem?
"It states that in a distributed system, you can only pick two: **Consistency**, **Availability**, and **Partition Tolerance**.

Since network partitions (P) are inevitable in the cloud, you really only choose between CP or AP.
**CP (Consistency)**: If a node goes down, the system blocks updates to ensure data remains accurate (like a Bank).
**AP (Availability)**: The system always responds, even if it might serve slightly stale data (like a Facebook timeline).

Most of my microservices favor Availability (AP) and settle for 'Eventual Consistency'."

---

### 3. SQL vs NoSQL?
"**SQL** (Relational) is best for structured data with strict relationships and ACID transactions. If I’m building a financial ledger or an e-commerce order system with complex joins, I pick SQL (Postgres/MySQL).

**NoSQL** (Non-relational) is for flexible schemas or massive throughput.
I use **MongoDB** (Document store) for catalogs or content management where inputs vary.
I use **Redis** (Key-Value) for caching.
I use **Cassandra** (Wide Column) for massive write-heavy logs."

---

### 4. How do you handle Distributed Transactions?
"In a monolith, we just use `@Transactional`. In microservices, it’s harder because databases are separate.
We usually avoid distributed transactions (like Two-Phase Commit) because they block and don't scale.

Instead, we use the **Saga Pattern**.
It’s a sequence of local transactions. Service A does its work and publishes an event. Service B listens and does its work.
If Service B fails, it fires a 'Compensation Event', and Service A listens to that to undo the changes (like issuing a refund). It relies on eventual consistency."

---

### 5. Horizontal vs Vertical Scaling?
"**Vertical Scaling** (Scaling Up) means buying a bigger machine—more RAM, more CPU. It’s easy but has a hardware limit.

**Horizontal Scaling** (Scaling Out) means adding *more* machines. This is the cloud-native way. We run 10 tiny instances of our service behind a Load Balancer. If load increases, we just spin up 5 more instances. This is theoretically infinite scaling."

---

### 6. What is a Load Balancer?
"It sits in front of your servers and distributes incoming traffic across them.

It ensures no single server gets overwhelmed. It also handles **Health Checks**—if Server 3 dies, the Load Balancer detects it and stops sending traffic there.
We typically use NGINX or cloud managed LBs (like AWS ALB)."

---

### 7. Caching Strategies?
"Caching helps reduce database load.

**Look-aside Cache** (most common): App checks Cache. If miss, it asks DB, then populates Cache.
**Write-through**: App writes to Cache, and Cache writes to DB. Keeps data consistent but slow writes.
**Eviction Policies** are critical. I usually use LRU (Least Recently Used) with a TTL (Time To Live). If you don't expire cache keys, you risk showing stale data to the user."

---

### 8. Explain JWT (JSON Web Token)?
"It’s a stateless way to handle authentication.
Instead of storing a Session ID in the server's memory (which is hard to scale), we generate a signed Token containing the user's ID and Roles. We send this to the client.

The client sends it back in the `Authorization` header for every request. The server just verifies the signature to know it's valid. It doesn't need to look up anything in the DB. This makes the auth layer completely scalable."

---

### 9. Circuit Breaker Pattern?
"It prevents cascading failures.
If Service A calls Service B, and Service B is down or slow, Service A shouldn't keep waiting and blocking threads.

A **Circuit Breaker** (like Resilience4j) detects the failures. After a threshold (say, 50% failure), it 'opens the circuit' and strictly fails fast immediately, without even calling Service B. This gives Service B time to recover and prevents the whole system from hanging."
