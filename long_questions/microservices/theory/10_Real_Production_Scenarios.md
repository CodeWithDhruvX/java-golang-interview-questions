# 🟢 **151–165: Real Production Scenario Questions**

### 151. How did you handle production outage?
"During a major outage where our Payment microservice started returning 500s, my immediate priority was *restoration*, not root-cause analysis. 

The API Gateway alerts fired in Datadog. I immediately joined the incident bridge. We checked the APM metrics and noticed the Payment pods were repeatedly crashing due to OutOfMemory (OOM) errors. 

Rather than debugging the heap dump while users failed to checkout, we instantly initiated a fallback routine. We manually scaled the pod count from 10 to 30 to dilute the sudden memory load and bought the service enough breathing room to stabilize. Once stabilized, we diverted traffic to a stable 'Blue' cluster, downloaded the heap dumps, and identified a runaway list allocation introduced in the previous day's deployment."

#### Indepth
Handling an outage requires strict incident command structure: an Incident Commander (IC) solely focused on coordinating communication and keeping stakeholders updated, distinct from the engineering 'Resolvers' actively querying logs and pushing mitigations. Post-incident, a blameless 'Postmortem' document is mandatory to prevent exact recurrence.

---

### 152. How did you debug memory leak?
"Our Spring Boot Order Service was restarting every 4 hours. Grafana showed the JVM heap usage steadily climbing diagonally upward without ever plateauing after Garbage Collection.

We couldn't reproduce it locally with minor traffic. So, we connected to a struggling production pod directly using Java Flight Recorder (JFR) and triggered a heap dump via Spring Boot Actuator (`/actuator/heapdump`).

Analyzing the 2GB `.hprof` file in Eclipse MAT, it became immediately obvious. The `Leak Suspects` report highlighted millions of orphaned `OrderDTO` objects retained permanently inside a static `ConcurrentHashMap` intended as an ad-hoc local cache that tragically had no eviction mechanism (TTL) implemented."

#### Indepth
Memory leaks in microservices are often caused by improper usage of `ThreadLocal` variables (especially in reactive frameworks or web servers like Tomcat where thread-pools are reused endlessly), or by adding large objects to global Maps acting as makeshift caches that grow indefinitely until `OutOfMemoryError` crashes the container.

---

### 153. How did you improve performance?
"The 'Generate Monthly Report' API in our Billing microservice was timing out constantly, taking over 15 seconds.

I traced the API call in Jaeger and saw that the logic executed 5,000 distinct SQL `SELECT` queries sequentially to fetch user details for each billable item—the classic N+1 query problem.

I refactored the Hibernate repository to use a bulk `JOIN FETCH` query, reducing 5,000 database round-trips to exactly 1. I then realized the report data only changed daily, so I cached the finalized JSON response in Redis with a 24-hour TTL. Latency dropped from 15,000ms down to 12ms."

#### Indepth
Performance optimization should strictly follow measurements, not guesses. Using an APM tool to identify the slowest spans usually reveals that the bottleneck is almost entirely I/O bound (database reads, external HTTP network latency, sluggish DNS resolution), not CPU-bound application logic executing algorithms.

---

### 154. How did you scale system?
"When our food delivery platform went viral, our monolithic Node.js backend entirely locked up during the 6:00 PM dinner rush.

First, we attacked the database: we upgraded (vertically scaled) the PostgreSQL instance and shifted all read queries (menu scanning) to three Read Replicas, freeing the primary solely for writes (saving orders).

Second, we tackled the application tier: we introduced Kubernetes. We wrapped the Node app in a Docker container and set up a Horizontal Pod Autoscaler (HPA). At 5:30 PM, K8s detected CPU load rising and automatically spawned 40 new instances of the API. By 8:30 PM, it gracefully killed them off. We handled 10x our normal traffic flawlessly."

#### Indepth
For horizontal scaling to work effectively, the application must be religiously stateless. Any session data or local memory storage prevents load balancers from indiscriminately routing user requests arbitrarily across the 40 new application nodes. Data must be pushed out to a distributed cache (Redis).

---

### 155. How did you design rate limiting?
"Our public weather API was frequently abused by scraping bots, degrading service for paying customers. 

We deployed the API Gateway (Kong) at the very perimeter of our AWS VPC. I implemented a 'Token Bucket' rate-limiting plugin mapping directly to the client's API Key. 

Free-tier users were strictly capped at 60 requests per minute. If they hit 61, Kong short-circuited the request, didn't even forward it to our internal Spring Boot microservices, and instantly returned an HTTP `429 Too Many Requests`. The backend microservices experienced drastically cleaner, predictive load patterns."

#### Indepth
Implementing rate limiting locally in application memory (e.g., using Google Guava) is a flawed distributed systems design because if you deploy 10 pods, the user can theoretically hit 600 requests. Distributed counters utilizing a blazing fast external system (like Redis `INCR` or Redis Lua scripts) ensure cluster-wide strict enforcement.

---

### 156. How did you optimize database?
"Our primary PostgreSQL database was approaching 95% CPU, causing timeouts.

Before sharding (which is complex and risky), I analyzed the `pg_stat_statements` table to find the most expensive, frequently run queries. One query searching user telemetry was taking 2,000ms. I executed an `EXPLAIN ANALYZE` and realized it was performing a massive Sequential Scan (reading every row on disk).

I simply added a composite B-Tree index covering the `status` and `timestamp` columns. The query flipped to an Index Seek, executing in 5ms. The overall database CPU utilization dropped instantly from 95% to twenty percent, postponing the need to shard by an entire year."

#### Indepth
Another aggressive optimization is horizontal partitioning. If the `orders` table holds 100 million rows, but users only actively query orders from the last 30 days, dividing the table mathematically by month reduces the query engine's scanning surface area significantly, radically improving I/O throughput.

---

### 157. How did you migrate legacy system?
"We had to decommission a 10-year-old monolithic Java application handling the company's core inventory.

We strictly utilized the Strangler Fig Pattern. I deployed an NGINX API Gateway routing 100% of traffic to the monolith. Over two months, my team wrote a lean, brand-new 'Warehouse Microservice' purely mimicking the existing APIs exactly. 

We deployed it alongside the monolith but initially sent *zero* traffic to it. Then, using 'Dark Launching', we shadowed live traffic to the new microservice passively to test performance without affecting real users. Finally, we flipped the Gateway routing table to send all `/warehouse` URLs to the new Go service, entirely bypassing the monolith. We repeated this process feature by feature."

#### Indepth
Data migration remains the hardest element. Often, two-way sync tools (like Debezium) are temporarily employed to keep the new isolated microservice database constantly synchronized perfectly with the legacy monolith database, ensuring a rollback is instantly possible without losing the data ingested during the new system's operational window.

---

### 158. How did you handle distributed deadlock?
"In our microservice ecosystem, the Order service locked an inventory row, then synchronously called the Shipping service. Simultaneously, the Shipping service locked a logistics row and synchronously called the Order service. 

Both services blocked instantly, waiting for the other to release the database lock. Standard local databases cannot detect this because the locks are held across two totally different servers over HTTP. The 30-second HTTP timeouts eventually severed the connection, failing transactions completely.

We fixed it by destroying the synchronous HTTP link. We switched to an asynchronous Saga choreography pattern over Kafka. The Order service committed locally and fired an event, never blocking or holding database locks while waiting for Shipping. 'Hold-and-Wait' distributed anti-patterns were eliminated entirely."

#### Indepth
Imposing a strict, rigid topological ordering on service interaction (Service A can ALWAYS call Service B, but Service B is architecturally forbidden from EVER calling Service A) inherently prevents cyclic dependency graphs, which mathematically guarantees distributed deadlocks cannot form.

---

### 159. How did you implement retry safely?
"We needed our API Gateway to automatically retry failed network calls to the backend Payment provider, as their API frequently dropped connections.

I implemented Resilience4j with an **Exponential Backoff and Jitter** algorithm. Let's say the Payment API hiccuped and rejected 5,000 simultaneous requests. If we retried them all blindly exactly 1 second later, we would effectively accidentally DDoS the struggling Payment provider.

Our strategy randomly delayed retry #1 between 1s and 2s, retry #2 between 2s and 4s, and then aborted. Crucially, I verified the upstream API endpoint was strictly idempotent, ensuring no user was ever double-charged if our first network response was simply lost in transit."

#### Indepth
Safe retries fundamentally enforce an upper limit of attempts (usually 3 max). They are often deliberately wrapped inside a catastrophic Circuit Breaker. If the Circuit Breaker trips "Open", all internal retries immediately abort without executing, decisively shielding the downstream service from futile requests.

---

### 160. What trade-offs did you make?
"The biggest trade-off in distributed architecture is always Consistency versus Availability (CAP Theorem).

We were designing a high-velocity Social Media feed where thousands of tweets per second were ingested. I explicitly sacrificed Strong Database Consistency (which required locking tables and utilizing 2PC, totally shattering our throughput) in favor of High Availability and Eventual Consistency.

When a user posted a comment, we returned a 'Success 200' immediately and updated the UI locally on the client-side, while pushing the actual save operation asynchronously into Kafka. Sometimes other users couldn't see the comment for 1-2 seconds until the databases caught up, but our API never crashed during the Super Bowl traffic spike."

#### Indepth
Every microservice is a trade-off. Splitting a monolith into twenty services trades "Ease of initial deployment and debugging" for "Decoupled domain ownership and independent scaling speed." Choosing gRPC over REST sacrifices "Browser curl readability" in pursuit of "Maximum binary throughput". There is no perfect architecture, only the right compromises.

---

### 161. How did you handle high traffic spike?
"We were featured unexpectedly on national television, and traffic to our website surged 50x in under one minute.

Our Auto-Scaling groups (HPA and EC2 Autoscalers) take roughly 3 minutes to spin up new pods and VMs. That was too slow; our API Gateway was drowning, and the database queued thousands of HTTP threads.

We immediately enabled our 'Load Shedding' toggle. At the edge CDN (Cloudflare), we deployed a heavily cached static version of our home page and deliberately disabled our expensive, dynamic 'Recommended Products' widget. This radically reduced backend database calls to near zero, providing our infrastructure the crucial 5 minutes it needed to scale up smoothly and resume full functionality."

#### Indepth
Relying purely on reactive autoscaling is deeply dangerous because traffic often scales vertically faster than software can boot. Proactive strategies rely on scheduled scaling (warming up 500 pods the night before Black Friday) and prioritizing critical API functionality while ruthlessly sacrificing auxiliary features via feature flags during catastrophic usage surges.

---

### 162. How did you prevent cascading failure?
"A cascading failure occurs when Service A fails, causing Service B (which relies on A) to exhaustion-crash, dragging down Service C, terminating the entire company's ecosystem.

In our old architecture, the 'Profile Image Resizer' service froze due to a bad library update. The 'Core API' continuously waited dynamically for image formatting, tying up every available Tomcat thread until the Core API itself went unresponsive, crashing the fundamental login flow.

I prevented this by strictly injecting Circuit Breakers and Bulkhead thread isolation. I separated the 'Image Resizer' API calls into their own tiny thread pool. When the resizer went down again, the circuit abruptly opened, instantly failing image requests cleanly. Crucially, the Login flow—running on a different thread pool— remained unaffected, preserving core system survival."

#### Indepth
Timeouts are the vital "first line of defense". An infinite timeout guarantees thread exhaustion when a downstream dependency vanishes. Setting aggressive timeouts (e.g., 2000 milliseconds max) ensures the request aborts and frees the executing thread before the memory queue congests entirely.

---

### 163. How did you implement idempotency?
"We processed thousands of incoming webhook events from Stripe for payments. Occasionally, Stripe experienced network issues and re-transmitted the exact same Webhook identically three times, causing our system to create three duplicate 'Payment Received' ledger entries.

I implemented an Idempotency Filter. I extracted the unique `stripe_event_id` and executed a Redis `SETNX` (Set if Not Exists) command with a 24-hour expiration. 

If `SETNX` returned True, our app processed the webhook and updated the database securely. If `SETNX` returned False, we instantly returned an HTTP `200 OK` to Stripe, completely ignoring the payload because we immediately recognized it as a duplicate event, preventing dirty data."

#### Indepth
For critical financial systems, relying solely on volatile Redis for idempotency checks is risky. A more robust implementation involves creating a dedicated `processed_events` table within the primary relational PostgreSQL database. Utilizing unique index constraints on the `event_id` ensures the transaction will structurally reject duplicates with mathematically perfect ACID safety.

---

### 164. How did you test microservices?
"Testing microservices purely end-to-end dynamically is a fragile, flaky anti-pattern because the test environment requires 50 services all functioning perfectly to pass. 

I entirely shifted the testing paradigm 'Left'. 

We focused aggressively on massive Unit Test coverage for localized business logic. Then, to ensure the services interacted correctly, we implemented Consumer-Driven Contract Testing using 'Pact'. The Consumer API asserts exactly what JSON response it expects. The Provider API runs these assertions internally during its own CI/CD build process. This mathematically guarantees the two services can communicate successfully in production without actually requiring them to talk over a live network in the testing stage."

#### Indepth
Proper testing "pyramids" for microservices shrink costly integration/E2E tests deliberately. Chaos Engineering supplements late-stage testing by rigorously testing infrastructure resilience (like terminating active pods arbitrarily during test-suite execution) rather than purely asserting successful business paths.

---

### 165. How did you handle breaking API changes?
"We needed to completely restructure our JSON response for the `User Profile` endpoint, renaming `firstname` and `lastname` to a nested `name` object.

Modifying the JSON in place immediately breaks all millions of older Mobile App clients currently installed on user phones that haven't updated yet. 

I resolved this gracefully using strictly URI Versioning. I deployed the new codebase exposing `/api/v2/users` alongside the untouched `/api/v1/users` endpoint. The old mobile applications comfortably consumed `v1` undisturbed. New web clients eagerly utilized `v2`. Over twelve months, as older mobile clients updated across the App Store, traffic on `v1` gradually evaporated to zero, and we safely deleted the `v1` code entirely."

#### Indepth
Semantic Versioning (SemVer) dictates that only breaking changes (removing fields, altering data-types detrimentally) demand a major version bump. Adding completely new, optional fields to an existing JSON response is fundamentally backward-compatible and does not warrant introducing a heavily burdensome `v2` endpoint structure.
