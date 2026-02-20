# 🔴 Architecture Patterns — Questions 41–50

> **Level:** 🔴 Senior (5+ yrs) — Core design patterns for scalable distributed systems
> **Asked at:** Uber, Lyft, Google, Twitter/X, Swiggy, Meesho (SDE-3 and Principal level)

---

### 41. What is microservices architecture?
"Microservices is an architectural style where an application is built as a **collection of small, independent, and loosely coupled services**, each owning its own data and business capability.

Each service is deployable independently. The Payment service team deploys 20 times a day without touching the User service. This is what unlocks developer velocity at scale. Amazon has thousands of microservices — each owned by a two-pizza team.

The classic principle: each microservice should be responsible for one business capability — the Order service manages orders, the Inventory service manages stock, the Notification service manages alerts. They communicate via APIs (REST/gRPC) or events (Kafka)."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Almost every product company mid-to-large scale — Flipkart, Swiggy, Uber, Amazon, Razorpay

#### Indepth
Key characteristics of microservices (per Martin Fowler):
1. **Componentization via services:** Each service is a separate deployable process. Failure of one service doesn't crash others.
2. **Organized around business capabilities:** Not technical layers (UI, DB, Logic) but business functions (Orders, Users, Payments).
3. **Decentralized data management:** Each service owns its own DB. Order service has its own MySQL. Inventory service has its own Postgres. No shared DB — this is sacrosanct.
4. **Smart endpoints, dumb pipes:** Unlike SOA's ESB (Enterprise Service Bus), microservices use simple transport (HTTP/Kafka). Business logic is in the service, not the middleware.
5. **Design for failure:** Every service can fail. Every inter-service call must have timeouts, retries, and circuit breakers.

**When NOT to use microservices:** At small scale (<20 engineers), the operational overhead (service discovery, distributed tracing, orchestration) outweighs the benefits. Start as a modular monolith, extract services when you feel specific scaling or team pain.

---

### 42. What is monolithic architecture?
"A monolith is an application where **all functionality is in one codebase, built as one deployable artifact** — one JAR, one binary, one process.

It's not inherently bad. For a startup or a small product, a well-structured monolith is often the right choice. It's simpler to develop, debug, test, and deploy. You don't have distributed system problems (no network calls between services, no eventual consistency, no distributed tracing needed).

The problems emerge at scale: if the Order module has a memory leak, it crashes the *entire* application — including the User module that was working fine. Scaling requires scaling *everything*, even if only the search module needs more resources. And 100 developers working in one repository creates merge conflicts and slow builds."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Service companies (TCS, Infosys) where legacy monolith modernization is common; also at product companies when asked to compare to microservices

#### Indepth
Monolith variants:
- **Modular Monolith:** Single deployable but code organized into well-defined modules with clear boundaries. This is the sweet spot before microservices are needed. Each module could be extracted to a service later. Examples: Shopify started as a Rails modular monolith and stayed that way for years.
- **Distributed Monolith (Anti-pattern):** Looks like microservices (separate services) but they share a database or have tight synchronous coupling. Worst of both worlds — microservice complexity without microservice benefits.

Migration path: **Strangler Fig Pattern** — gradually replace a monolith with microservices by building new functionality as services, and moving existing functionality piece by piece. The monolith "strangles" over time. Never do a big-bang rewrite — it almost always fails.

---

### 43. Difference between microservices and SOA.
"Both SOA and microservices decompose systems into services but differ fundamentally in scope and communication.

**SOA (Service Oriented Architecture)** is an enterprise pattern from the early 2000s. Services communicate through a centralized **ESB (Enterprise Service Bus)** — a heavy middleware that handles routing, transformation, orchestration, and protocol mediation. SOA was designed for enterprise integration (connecting SAP, Oracle, mainframes).

**Microservices** are a modern evolution: lighter communication (HTTP/REST/gRPC or message queues directly), decentralized data, team autonomy. No centralized ESB. The philosophy: 'smart endpoints, dumb pipes'."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Enterprise service companies (Infosys, Wipro clients), banks migrating off SOA — also FAANG design comparisons

#### Indepth
| Feature | SOA | Microservices |
|---|---|---|
| Scope | Enterprise-wide integration | Application-level decomposition |
| Communication | Centralized ESB (SOAP, WSDL, XML) | Direct (REST/gRPC) or messaging (Kafka) |
| Data | Often shared DB / data warehouse | Decentralized — DB per service |
| Team size | Cross-functional enterprise teams | Small autonomous two-pizza teams |
| Governance | Centralized (ESB team controls) | Decentralized (each team owns service) |
| Granularity | Coarser (Order Processing System) | Finer (Order Service, Order Item Service) |
| Example | IBM WebSphere ESB | Netflix OSS, Kubernetes microservices |

ESBs are a major anti-pattern in modern systems because they concentrate complexity in one place — a smart "pipe" that knows too much about the business logic of the services it connects. If the ESB is wrong or slow, everything is wrong and slow.

---

### 44. What is service discovery?
"In a static environment, you hardcode service IPs. In a dynamic environment (Kubernetes, auto-scaling), service IPs change constantly. Service discovery is the mechanism by which services **find each other's current network addresses at runtime**.

Think of it like a phone book that's always up-to-date, updated in real-time as services come up and go down. Services register their IP:port on startup; they deregister on shutdown. Other services query the registry to find where to send requests.

There are two patterns: **Client-side discovery** where the client queries the registry and then load-balances directly. **Server-side discovery** where the client just calls the LB, and the LB queries the registry."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Kubernetes-heavy shops — Google, Amazon (EKS), Flipkart, Swiggy, any company running microservices

#### Indepth
Service discovery implementations:
- **Client-Side Discovery (Netflix Eureka + Ribbon):** Service registers with Eureka on startup. Ribbon (client library) queries Eureka to get list of healthy instances and applies its own LB algorithm. Eureka is now largely replaced by...
- **Kubernetes DNS:** Every Service in K8s gets a DNS entry: `service-name.namespace.svc.cluster.local`. Kubernetes CoreDNS resolves this. No client library needed — it's transparent.
- **Consul:** Multi-DC service registry with health checking, K/V store, and service mesh capabilities. Works across cloud providers and on-prem.
- **Server-Side Discovery:** Client calls Kubernetes Service (ClusterIP). K8s iptables (kube-proxy) routes to a healthy pod. The *client never knows* about individual pod IPs.

**Sidecar proxy pattern:** Envoy sidecar proxy intercepts all traffic from/to a pod. It queries a control plane (Istio's Pilot/xDS) for service discovery and routing rules. This moves service discovery completely out of application code.

---

### 45. How do services communicate in microservices?
"Service communication in microservices falls into two categories: synchronous and asynchronous.

**Synchronous (request-response):** Service A calls Service B and waits for a response. I use **REST over HTTP/2** for external APIs and cross-team services — it's human-readable and widely supported. For internal high-performance calls, I prefer **gRPC** — binary Protobuf encoding is 5-10x smaller than JSON, has strict schema contracts, and supports bidirectional streaming.

**Asynchronous (event-driven):** Service A publishes an event to Kafka. Service B subscribes and processes it in its own time. This decouples services completely — Service A doesn't wait for B, doesn't even know B exists. Perfect for workflows where eventual consistency is acceptable: order placed → payment service picks it up → inventory service picks it up."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Lyft, Amazon, Flipkart, Swiggy — any microservices design discussion

#### Indepth
Communication patterns and when to use each:
- **REST:** Use for external-facing APIs, when human readability matters, or when callers use diverse languages/platforms.
- **gRPC:** Use for internal service-to-service communication needing high throughput and strict contracts. Bi-directional streaming for real-time (chat, live tracking).
- **GraphQL:** Use for client-driven data fetching where different clients (mobile vs web) need different data shapes. API gateway aggregation.
- **Kafka/Event Streaming:** Use for decoupled async workflows, audit logs, event sourcing, data pipeline integrations.
- **Message Queue (RabbitMQ/SQS):** Use for task queues (email sending, image processing) where at-least-once delivery and retry are needed.

**Anti-pattern: Synchronous chains.** If Service A calls B which calls C which calls D, you have cascading latency (each adds latency) and cascading failure (D going down takes A down). Break these chains with async events or aggregate in an API gateway (BFF pattern).

---

### 46. What is an API gateway?
"An API gateway is a **single entry point for all clients into a backend system**. Instead of clients knowing about 50 microservices, they talk to one gateway URL. The gateway routes, transforms, and manages the request.

It consolidates cross-cutting concerns that every service needs: authentication (verify JWT token once at the gateway, not in every service), rate limiting (client can only make 1000 requests/minute), SSL termination, logging, and request routing based on URL path. This keeps microservices lean and focused on business logic.

Kong, AWS API Gateway, and Netflix Zuul are popular options. I've used the BFF (Backend for Frontend) pattern with API gateways — separate gateways for mobile app and web app, each aggregating data differently."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (API Gateway product team), Stripe (API-first company), Razorpay, Flipkart

#### Indepth
API gateway responsibilities:
1. **Routing / Reverse Proxy:** `GET /orders` → Order Service, `GET /products` → Product Service
2. **Auth / AuthZ:** Verify JWT/API key. Decode user identity. Pass user context headers to downstream service.
3. **Rate Limiting:** Token bucket or sliding window per client/API key.
4. **Request/Response Transformation:** Translate REST to gRPC, add/remove headers, merge responses from multiple services.
5. **Circuit Breaking:** Stop forwarding to a failing service.
6. **Logging / Tracing:** Inject correlation IDs into all requests for distributed tracing.
7. **Caching:** Cache frequently requested responses.

**BFF (Backend for Frontend):** Instead of one gateway, have one per client type: Mobile BFF (optimized for bandwidth-constrained mobile), Web BFF (richer responses for browser), Partner API (standardized external API). This lets each BFF be tailored to its consumer's needs.

---

### 47. What is the circuit breaker pattern?
"The circuit breaker is a design pattern that **prevents cascading failures** in microservices. Like an electrical circuit breaker that trips when there's a fault, it detects when a downstream service is failing and stops sending requests to it — failing fast instead of waiting for timeouts.

The circuit breaker has three states: **Closed** (normal operation — all requests flow through), **Open** (downstream is failing — requests immediately return an error without touching the downstream, like returning a cached response or a graceful error), and **Half-Open** (after a cooldown period, let a test request through — if it succeeds, circuit closes; if it fails, it opens again).

The timeout problem is what circuit breakers solve: without one, if the Payment service is taking 30s to respond (instead of 200ms), all my goroutines/threads are blocked for 30s waiting. The thread pool fills up. My service starts failing too. This cascading failure took down Twitter multiple times in its early days."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix (invented Hystrix for this), Uber, Amazon, Flipkart — any resilience-focused interview

#### Indepth
State machine detail:
```
CLOSED ──(failures > threshold)──► OPEN
  ▲                                   │
  │                                   ▼
  └──(test request succeeds)── HALF-OPEN
                                (let 1 request through)
```

Parameters to tune:
- **Failure threshold:** How many failures in a time window to trip. (e.g., 50% error rate in 10 seconds)
- **Slow request threshold:** Treating timeouts as failures (e.g., >1s = failure)
- **Wait duration in Open state:** How long before trying Half-Open (e.g., 30 seconds)
- **Half-Open permitted calls:** How many test requests to allow

**Fallback strategies when circuit is Open:**
- Return cached data (stale but better than nothing)
- Default/empty response ("No recommendations available")
- Queue the request for later processing
- Error with a user-friendly message

Libraries: **Resilience4j** (Java), **Hystrix** (Netflix, now deprecated), **Polly** (.NET), **go-resilience** (Go). Istio/Envoy can implement circuit breaking at the service mesh layer — no application code needed.

---

### 48. What is the saga pattern?
"The saga pattern solves the **distributed transaction problem** in microservices. ACID transactions are easy in a single DB — the DB handles atomicity with a lock and rollback. But across multiple microservices with separate DBs, you can't use a single DB transaction.

The saga breaks a distributed transaction into a **sequence of local transactions**, each in a single service. If a step fails, you execute **compensating transactions** (reverse operations) for all previous successful steps to undo them.

For a hotel booking system: (1) Reserve room (Hotel Service), (2) Charge card (Payment Service), (3) Send confirmation (Notification Service). If step 2 fails, run compensation for step 1: `cancelRoomReservation`. If step 3 fails, do nothing (notification failure doesn't warrant refund)."

#### 🏢 Company Context
**Level:** 🔴 Senior – Principal | **Asked at:** Flipkart (e-commerce checkout), Razorpay, PhonePe (payment flows), Amazon (order fulfillment)

#### Indepth
Two saga implementation styles:
- **Choreography (Event-driven):** Each service publishes an event when its local transaction completes. The next service listens and starts its own transaction. No central coordinator. Decoupled, but hard to visualize the complete flow and debug failures.
  ```
  OrderService → event: OrderCreated
    → PaymentService → event: PaymentCompleted
      → InventoryService → event: InventoryDeducted
        → ShipmentService → event: ShipmentCreated
  ```
- **Orchestration (Centralized):** A Saga Orchestrator (a dedicated service or workflow engine like Temporal, AWS Step Functions) commands each service step-by-step. If a step fails, the orchestrator issues compensating commands. Easier to understand, easier to handle failures.

**Compensating transactions must be:**
- **Idempotent:** Same compensation run twice produces the same result
- **Semantic rollbacks only:** You can't always literally reverse a DB operation (email already sent → can't un-send → instead, send a cancellation email)

SAGAs are complex. Many teams overuse them. First ask: can the business tolerate eventual consistency with simple retries? Often the answer is yes, and a saga is overkill.

---

### 49. What is eventual consistency in microservices?
"In microservices, different services have different databases. When Service A updates its DB and needs to communicate that update to Service B, there's a delay — during which Service A and Service B have different views of the world. This inconsistency window is what eventual consistency describes.

The pragmatic approach: model your domain so most operations don't need cross-service strong consistency. A user's post appearing 200ms later on a follower's feed is acceptable. A double-charge on a credit card is not.

The mechanism: Service A completes its local transaction and publishes an event to Kafka. Service B consumes the event and updates its own DB. During the Kafka lag (usually milliseconds), the data is inconsistent between services. This is accepted and designed for."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (their Dynamo paper is foundational), Flipkart, Swiggy, Uber

#### Indepth
Patterns for managing eventual consistency in microservices:
- **Outbox Pattern:** Avoid losing events if the service crashes after writing to DB but before publishing to Kafka. Solution: write the event to an `outbox` table *in the same local DB transaction* as the business data. A separate poller reads the outbox and publishes to Kafka. Atomicity guaranteed by local transaction; event delivery to Kafka is best-effort with retries.
- **Event Sourcing:** Instead of storing current state, store the history of all events. The current state is derived by replaying events. Kafka becomes the system of record. Any service can subscribe and build its own query-optimized projection.
- **CQRS (Command Query Responsibility Segregation):** Separate the write model (Commands → normalized DB) from the read model (Queries → denormalized, read-optimized projections updated asynchronously). The read model may be slightly stale — that's the eventual consistency trade-off.

**Handling inconsistency in the UI:** Optimistic UI updates — show the user the result immediately as if it succeeded (positive feedback), while the async operation completes in the background. If it fails, rollback the UI. WhatsApp uses this for message sending.

---

### 50. How to ensure idempotency?
"Idempotency means an operation can be **performed multiple times without changing the result beyond the first application**. This is critical for payment APIs, order placement, and any mutation operation in distributed systems.

The pattern: the client generates a globally unique **idempotency key** (UUID) and includes it in the request header. The server checks if this key was already processed. If yes, return the *same stored response* without re-executing. If no, process the request, store the result, and return it.

For a payment API: if a client makes a payment request and the network times out, they don't know if the payment went through. Without idempotency, retrying might double-charge the user. With an idempotency key, the server just returns the stored `{status: 'success', transactionId: 'xyz'}` from the first attempt."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, Paytm, Stripe (payments), Amazon (order placement), any API-centric company

#### Indepth
Implementation steps:
1. Client generates `Idempotency-Key: <UUID>` header
2. Server checks Redis/DB: `GET idempotency:uuid`
   - **Hit:** Return stored response immediately (don't re-process)
   - **Miss:** Process request, store result with a TTL (e.g., 24 hours), return result
3. Handle concurrent duplicate requests: The first request acquired a DB lock / Redis SET NX (`SET key placeholder NX EX 30`). Others wait or receive a 409 Conflict.

Built-in idempotency in HTTP verbs: GET, PUT, DELETE are inherently idempotent (calling them multiple times has the same effect). POST is not — that's why explicit idempotency keys matter for POST requests.

**Stripe's implementation:** Stripe stores idempotency keys with the full request body hash. If the same key is reused with a different request body (different amount), it returns a 400 error — preventing deliberate misuse where someone reuses a key to change the amount on a successful charge.

**Database-level idempotency (Upsert):** `INSERT INTO payments ... ON CONFLICT (idempotency_key) DO NOTHING RETURNING *` — the DB's unique constraint guarantees at-most-once insertion even under concurrent requests.
