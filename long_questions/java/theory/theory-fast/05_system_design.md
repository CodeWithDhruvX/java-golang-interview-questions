# System Design Interview Questions (71-80)

## Scalability & Distributed Systems

### 71. How do you design a scalable REST service?
"Scalability starts with **Statelessness**.

The server shouldn't store client context (like 'User is on page 2') in memory. If I have 3 servers, a request might hit Server A, and the next one hits Server B. If state is on A, B fails.

So, I use tokens (JWT) for auth and store any session data in an external store like Redis.

Then, I put a **Load Balancer** in front to distribute traffic. I decouple heavy processing using **Message Queues** (Kafka/RabbitMQ) so the API stays responsive. And I use **Caching** (CDN for static assets, Redis for data) to reduce DB load."

**Spoken Format:**
"Designing a scalable REST service is like building a highway system that can handle 10x more traffic without falling apart.

The foundation is **Statelessness** - your servers shouldn't remember who the user is. Each request should be self-contained. If I hit Server A with my request, then Server B should handle the next request just fine.

Why? Because if servers remember state, you have to worry about which server gets which requests. It's like having multiple checkout counters that need to stay synchronized.

For authentication, I use tokens (JWT) instead of sessions. The token contains all the user info, and I don't need to store session data.

For traffic distribution, I use a **Load Balancer** - think of it as a traffic cop directing cars to different lanes. It uses algorithms like Round Robin (taking turns) or Least Connections (sending to least busy server).

For heavy work, I use **Message Queues** like Kafka - instead of making the user wait for slow processing, I put the task in a queue and respond immediately.

And **Caching** is like having local storage for frequently requested items - like keeping popular products in a local warehouse instead of going to the main warehouse every time.

The goal is that no single server becomes a bottleneck, and the system can grow by adding more servers horizontally!""

### 72. How does load balancing work?
"A Load Balancer (LB) sits between the client and your servers. It acts as a traffic cop.

It accepts the incoming request and forwards it to one of the healthy backend servers based on an algorithm—usually **Round Robin** (take turns), or **Least Connections** (send to the server doing the least work).

It also does **Health Checks**. If Server 3 crashes, the LB detects it (via a /health endpoint) and stops sending traffic there until it recovers. This ensures high availability."

**Spoken Format:**
"A Load Balancer is like having a smart traffic manager for a busy parking lot.

Imagine cars arriving at your parking lot entrance. The Load Balancer is the attendant who directs each car to an available parking spot.

It uses different strategies:

**Round Robin** - like taking turns: Car 1 goes to Spot A, Car 2 to Spot B, Car 3 to Spot C, then back to Spot A. Fair distribution.

**Least Connections** - like sending cars to the least crowded areas: if Spot A has 5 cars and Spot B has 2, send the next car to Spot B.

The magic is the **Health Checks** - the Load Balancer constantly checks if each parking spot (server) is still available. If Spot C crashes, the attendant stops sending cars there until it's fixed.

This ensures high availability - even if one parking area fails, cars are still directed to the working spots.

The Load Balancer also handles **session affinity** when needed - if a user is in the middle of shopping, you want to keep sending them to the same server so their shopping cart stays there.

It’s like having a really smart parking attendant who not only directs traffic but also monitors the health of the entire parking lot!"

### 73. What is horizontal vs vertical scaling?
"**Vertical Scaling (Scaling Up)** is making the single server stronger: adding more RAM, more CPU. It’s the easiest way to solve a problem—‘just upgrade the AWS instance type’. But it has a hard limit (you can’t add infinite RAM) and is a single point of failure.

**Horizontal Scaling (Scaling Out)** is adding more servers. Instead of one super-computer, you have 10 commodity servers. This is harder to implement (you need load balancers, stateless apps, distributed DBs), but the ceiling is virtually infinite. You can just keep adding nodes."

**Spoken Format:**
"This is like choosing between making one employee super strong versus hiring more employees.

**Vertical Scaling** is like upgrading one employee’s computer - giving them more RAM, faster CPU, bigger monitor. They can handle more work, but there’s a limit - you can’t upgrade infinitely.

It’s the simple solution when you have one big task that needs more power. ‘The database is slow, just upgrade the server.’

**Horizontal Scaling** is like hiring more employees instead of upgrading existing ones. If you have 10 employees each handling 10 requests per second, you can handle 100 requests per second by hiring 10 more employees.

The ceiling is virtually infinite - you can keep hiring more people as your business grows.

Horizontal scaling is harder because you need to deal with:
- Load balancers to distribute work
- Making your application stateless so any employee can handle any request
- Distributed databases so everyone has access to the same data
- More complex deployment and monitoring

But for most modern applications, horizontal scaling is the only way to handle massive growth!"

### 74. How would you design a URL shortener?
"This is a classic. The core is a mapping between a `long_url` and a `short_code`.

I’d use a SQL database (PostgreSQL) or NoSQL (DynamoDB) to store `{id, long_url, short_code, created_at}`.

The tricky part is generating the `short_code`. I wouldn’t use a hash function (too long). I’d use **Base62 encoding** (A-Z, a-z, 0-9) on a unique ID.
But distributing unique IDs is hard. I’d use a dedicated ID Generator service (like Snowflake) or a pre-generation service (Token Service) to hand out ranges of IDs to servers.

For reads, since URLs rarely change, I’d cache redirects heavily in Redis. 99% of traffic would hit cache, not DB."

**Spoken Format:**
"URL shorteners are like creating a personalized forwarding service for mail.

The core challenge is generating unique, short codes that are hard to guess.

I’d store in a database: `{id, long_url, short_code, created_at}`. When someone wants to create a short URL, I generate a random code like ‘abc123’.

The problem is distributing these codes across multiple servers. If Server A generates ‘abc123’ and Server B generates ‘abc123’, you’ll have conflicts!

The solutions are:

1. **Dedicated ID Generator** - a separate service whose only job is to hand out unique ranges of IDs to each server.

2. **Snowflake IDs** - generate IDs that are unique across all servers without needing coordination. Each ID contains timestamp and machine ID.

For performance, I’d cache the redirects heavily in Redis. Since the long URL never changes, 99% of requests would hit the cache instead of the database.

It’s like having a directory assistance that remembers everyone’s extension - once you look up ‘abc123’, you never need to ask again!"

### 75. How would you design a rate limiter?
"I’d implement this using Redis because it’s fast and atomic.

A popular algorithm is the **Token Bucket**. Every user gets a ‘bucket’ that fills with tokens at a constant rate. Each request costs a token. If the bucket is empty, request rejected (429).

Alternatively, the **Sliding Window** algorithm interacts better with burst traffic. I’d store a sorted set (ZSET) in Redis for each user, containing timestamps of their last N requests. When a new request comes, I remove timestamps older than the window (e.g., 1 minute) and count the remaining. If count < limit, allow.

For a global distributed limiter, I’d use Redis + Lua scripts to ensure atomicity."

**Spoken Format:**
"Rate limiting is like having a bouncer at a club who ensures everyone gets fair access.

The **Token Bucket** algorithm is like giving each person a set of tokens when they enter the club. Each request costs one token. If they run out, they have to wait until next month when they get more tokens.

The bouncer checks their token count - if they have tokens left, they can enter. If not, they’re blocked until next month.

The **Sliding Window** is smarter for handling bursts. Imagine someone tries to make 10 requests quickly. The sliding window counts requests in the last minute. If they haven’t exceeded the limit in that window, they can still enter even if they used all their monthly tokens.

For distributed systems, I’d use Redis with Lua scripts to ensure the token counting is atomic across multiple servers.

It’s like having multiple bouncers at different doors who all communicate to ensure no one can sneak in by going to a different bouncer!"

### 76. Where would you use caching?
"I use caching at every layer where data is read frequently but changes rarely.

1.  **Browser/CDN**: Cache static assets (JS, CSS, Images). This is ‘free’ speed.
2.  **API Gateway / Reverse Proxy**: Cache entire API responses for public, static data (e.g., ‘List of Countries’).
3.  **Application / Service**: Internal object caching.
4.  **Database**: Buffer pool (DB does this automatically).

The most common specific use case is caching expensive database queries or results of complex calculations."

**Spoken Format:**
"Caching is like having a photographic memory for things you look up often.

I cache at multiple levels based on how frequently data changes:

**Browser/CDN** - for static assets that never change (logo images, CSS files). This is like having these assets stored locally in every user’s browser - super fast access.

**API Gateway** - for expensive API responses that multiple users request (like country lists, exchange rates). Cache for a day, serve thousands of users from cache.

**Application Layer** - for complex business logic results (like recommendation algorithms, report generation). Cache for an hour, avoid recalculating the same expensive operations.

**Database** - most databases have built-in query result caching. This is like the database remembering the answer to a question it already solved.

The key is knowing when to cache:
- Read-heavy, write-rarely → cache aggressively
- Write-heavy, read-occasionally → cache lightly or not at all
- Time-sensitive data → use TTL (Time To Live) so cache expires automatically

It’s like having different types of memory - short-term memory for immediate needs, long-term memory for reference data!"

### 77. Redis vs in-memory cache?
"**In-memory cache** (like Caffeine or a simple HashMap) lives inside your JVM heap. It’s the fastest possible cache (nanoseconds). But:
1. It consumes your application’s RAM.
2. It’s local. Server A doesn’t see Server B’s cache. You might have data inconsistency.
3. It dies when the app restarts.

**Distributed cache** (Redis) lives on a separate server. It’s slightly slower (network call ~1ms), but:
1. It’s shared across all instances. If Server A updates cache, Server B sees it.
2. It persists (can survive app restarts).

I use in-memory for tiny, static reference data (like country codes). I use Redis for almost everything else."

**Spoken Format:**
"Redis vs in-memory cache is like choosing between a personal notebook and a shared whiteboard.

**In-memory cache** is like a personal notebook - it’s super fast, but only you can see it. If you restart your app, the notebook is gone.

**Distributed cache** (Redis) is like a shared whiteboard - it’s slightly slower, but everyone can see it. If one person updates the board, everyone sees the change.

I use in-memory for tiny, static data that never changes (like country codes). I use Redis for everything else - user sessions, product catalogs, recommendation results.

The choice depends on your use case: speed vs. persistence vs. sharing vs. size. Each has its place in a good caching strategy!"

### 78. Database vs cache consistency?
"This is the hardest problem in computer science.

There are three main patterns:

1.  **Cache-Aside (Lazy Loading)**: Read from cache. Miss? Read DB -> Write Cache. Update? Update DB -> *Delete* Cache. Next read re-populates it. This is the safest and most common.
2.  **Write-Through**: Write to Cache and DB synchronously. Strong consistency, but higher write latency.
3.  **Write-Back**: Write only to Cache. Cache async writes to DB. Super fast, but risky—if Cache dies, data is lost.

I almost always stick to **Cache-Aside** with a **TTL (Time To Live)** as a safety net. Even if code fails to invalidate cache, TTL ensures it eventually corrects itself."

**Spoken Format:**
"Cache consistency is like keeping multiple whiteboards in a classroom synchronized - it’s surprisingly hard!

**Cache-Aside (Lazy Loading)**: Read from cache first. If cache miss, read from database and update cache. When data changes, update database then delete cache. Next read repopulates cache.

This is the safest approach because cache and database eventually converge. Like having a rule: always check the whiteboard first, if it’s empty, ask the teacher.

**Write-Through**: Write to cache and database at the same time. Strong consistency, but slower writes because you’re updating two places.

**Write-Back**: Write only to cache, then asynchronously update database. Super fast writes, but risky - if cache fails, you lose data.

I almost always use Cache-Aside with TTL - it’s like having an automatic eraser for the whiteboard. Even if a student forgets to erase it after updating, the whiteboard will automatically clear itself after a set time.

The TTL ensures that even if your code has bugs and forgets to invalidate cache, the cache won’t stay wrong forever!"

### 79. Microservices pros and cons?
"**Pros**:
-   **Independent Scaling**: Scale the ‘Checkout’ service without scaling ‘User Profile’.
-   **Technology Agnostic**: Use Java for backend, Python for ML, Node for I/O.
-   **Fault Isolation**: If ‘Recommendations’ crash, ‘Checkout’ still works.
-   **Faster Deployments**: Compile and deploy just one small service.

**Cons**:
-   **Complexity**: Distributed systems are hard. Network failures, latency, eventual consistency.
-   **Operational Overhead**: You need robust monitoring, logging, and deployment pipelines (Kubernetes).
-   **Data consistency**: No ACID transactions across services (Sagas are complex).

It’s not a silver bullet; it’s a trade-off."

**Spoken Format:**
"Microservices are like breaking your big company into small, specialized teams instead of having one giant department.

**The Pros:**
- Each team can scale independently - if the ‘Recommendations’ team needs more power, just upgrade their servers
- Each team can use different technology - ML team uses Python, ‘Payments’ team uses Java
- If one team fails, others keep working - the ‘Comments’ service can go down without taking down the whole app
- Faster deployments - small changes are quicker to test and deploy

**The Cons:**
- **Complexity** - now you need to handle network failures, service discovery, distributed transactions
- **Data Consistency** - no ACID guarantees across services. You need complex patterns like Sagas
- **Operational Overhead** - you need monitoring, logging, Kubernetes, service mesh
- **Deployment** - what used to be one deploy command is now 50+ microservices

**When to choose:**

Start with a **Modular Monolith** - organize your single application into well-defined modules. It’s faster to develop and easier to understand.

Only move to microservices when you have:
- Different scaling needs for different parts
- Multiple teams working independently
- A feature that’s so critical it needs its own infrastructure

Don’t do microservices just because they’re trendy - do it because they solve real problems for your specific situation!"

### 80. Monolith vs microservices — when to choose what?
"Start with a Monolith. Always.

A **Modular Monolith** is faster to develop, easier to test, and trivially easy to deploy. You keep domains separate via packages/modules, not network calls.

Move to Microservices only when:
1.  **Team Scaling**: You have 50+ devs stepping on each other's toes in the same repo.
2.  **Tech Requirements**: One part needs massive CPU scaling (video encoding) while another needs huge memory (in-memory analytics).
3.  **Fault Tolerance**: One feature is critical (Payments) and must not be taken down by a memory leak in a non-critical feature (Comments).

Don't do microservices just because Netflix does."

**Spoken Format:**
"Starting with a monolith is like having one big restaurant with different sections.

A **Modular Monolith** keeps everything organized - the kitchen is separate from the dining area, the bar is separate from the payment area. But it's all one connected building with one shared kitchen.

The benefits are:
- **Faster development** - one codebase, one deployment, easier debugging
- **Simpler testing** - everything is in one process, no network calls to test
- **Easier operations** - one database, one transaction, no distributed consistency issues

Move to microservices only when:

1. **Team scaling** - you have 50+ developers and different teams need different technology stacks
2. **Different scaling needs** - video processing needs huge servers, user management needs fast databases
3. **Fault isolation** - one critical feature shouldn't take down the entire system

The key insight: microservices solve organizational and scaling problems, but they create technical problems. Make sure you're solving real problems, not just following architecture trends!""
