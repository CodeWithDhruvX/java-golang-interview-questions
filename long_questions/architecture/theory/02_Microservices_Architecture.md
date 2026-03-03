# 🔧 Microservices Architecture — Questions 1–12

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Amazon, Flipkart, Swiggy, Zomato, PhonePe, Razorpay, Uber

---

### 1. How do you decompose a monolith into microservices?

"The most important decomposition strategy is **by business domain (Bounded Context from DDD)** — not by technical layer. A 'Payment Service' is a good microservice boundary. A 'Database Access Service' is a terrible one.

My process: Start with an **Event Storming** session to map out the business domain — what events happen, what commands trigger them, what aggregates handle them. This gives you natural seams. Then draw bounded context boundaries around groups of related aggregates. Each bounded context is a microservice candidate.

Secondary decomposition strategies: by **volatility** (components that change frequently should be separate so they can deploy independently), by **scalability** (separate what needs to scale differently — image processing vs auth), and by **team boundaries** (Conway's Law)."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart, senior architecture roles

#### Indepth
Decomposition approach — the Strangler Fig Pattern:
1. **Map the monolith:** Identify domain boundaries within the codebase
2. **Start at the edges:** Extract the least-coupled services first (notifications, billing)
3. **API Gateway facade:** Route some paths to new services, rest to monolith
4. **Strangle incrementally:** As each piece is migrated, the monolith shrinks
5. **Kill the monolith:** When all functionality is migrated

Anti-patterns to avoid:
- **Too many services, too soon:** Each service is a deployment unit — 50 services means 50 CI/CD pipelines, 50 monitoring dashboards. Premature decomposition is as bad as no decomposition.
- **Shared database:** If two services share tables, they're not independent — they're a distributed monolith.
- **Chatty services:** Service A calling B calling C calling D synchronously — 4-hop chain with cascading failure risk.

---

### 2. What communication patterns exist between microservices?

"Two fundamental styles: **Synchronous** (the caller waits for a response) and **Asynchronous** (the caller fires and forgets).

Synchronous: REST/HTTP (simple, universal, but tight coupling) and gRPC (binary protocol, typed contracts, 10x faster than JSON REST, but needs protobuf schema management). I use gRPC for internal service-to-service calls where performance matters and REST for external/public APIs.

Asynchronous via message brokers (Kafka, RabbitMQ): The caller publishes an event and moves on. The consumer processes it in its own time. This provides temporal decoupling — if the consumer is down, the message waits in the queue. Critical for things like 'order placed → email notification → update analytics' where each step can fail independently."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Swiggy, Razorpay, Zepto

#### Indepth
| Pattern | Protocol | Coupling | Latency | Use Case |
|---------|----------|----------|---------|----------|
| REST/HTTP | HTTP/JSON | Tight (sync) | Medium | External APIs, simple CRUD |
| gRPC | HTTP/2 + Protobuf | Tight (sync) | Low | Internal high-perf calls |
| Message Queue | AMQP/Kafka | Loose (async) | High (decoupled) | Work queues, task distribution |
| Event Streaming | Kafka | Loose (async) | High (decoupled) | Event sourcing, audit logs |
| GraphQL | HTTP | Tight (sync) | Medium | Complex query aggregation |
| WebSocket | TCP | Stateful | Very Low | Real-time (chat, live updates) |

**Orchestration vs Choreography:**
- **Orchestration:** A central orchestrator (Order Saga) tells other services what to do step by step. Easy to trace. Risk: orchestrator becomes a bottleneck.
- **Choreography:** Each service reacts to events from other services with no central controller. Highly decoupled. Hard to trace flow.

---

### 3. What is the API Gateway pattern?

"An API Gateway is a **single entry point for all client requests** to the backend microservices. Instead of the mobile app knowing the addresses of 20 different services, it talks to one gateway.

The gateway handles: **auth/authorization** (validate JWT once, not in every service), **rate limiting** (protect backend from traffic spikes), **request routing** (forward `/payments/*` to Payment Service, `/orders/*` to Order Service), **response aggregation** (call multiple services and combine results for one client request), and **observability** (centralized logging of all requests).

AWS API Gateway, Kong, Netflix Zuul, and Envoy are common implementations. I've used Kong for enterprise projects — it's plugin-based so you can add rate limiting, auth, and compression without touching backend code."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Flipkart, Swiggy, PhonePe

#### Indepth
API Gateway responsibilities:
1. **SSL termination:** Decrypt HTTPS at gateway, pass HTTP internally (reduces cert management overhead)
2. **Authentication:** Validate JWT or API keys. Inject user identity as header downstream.
3. **Rate limiting:** Per-user, per-IP, per-API-key limits. Kong: 100 req/min per user.
4. **Request/Response transformation:** Add/remove headers, transform bodies.
5. **Circuit breaking:** If downstream service is unhealthy, return fallback.
6. **Load balancing:** Gateway knows about service instances (via service registry).
7. **Caching:** Cache GET responses to avoid hitting backend.

**Gateway pitfall:** Don't put business logic in the gateway. It becomes a bottleneck and a single point of failure. Auth validation is acceptable; order calculation is not.

---

### 4. What is the Backend for Frontend (BFF) pattern?

"BFF (Backend for Frontend) is a variant of the API Gateway pattern where you create **one dedicated backend per client type** instead of one generic backend for all clients.

The problem: A mobile app needs `user + feed + notifications` in one request to avoid multiple round trips over 4G. The desktop web app needs different data. A generic API can't optimize for both without becoming bloated.

Solution: Mobile BFF aggregates and formats data specifically for the mobile app. Web BFF does the same for web. Each BFF is owned by the team that owns the frontend — they can iterate on it without coordinating with other backend teams."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Flipkart, Amazon, Netflix — companies with multiple client types

#### Indepth
BFF vs single API Gateway:
| Aspect | Single Gateway | BFF |
|--------|---------------|-----|
| Ownerhsip | Platform team | Frontend teams |
| Optimization | Generic, lowest common denominator | Client-specific |
| Data | Client filters what it needs | BFF returns exactly what client needs |
| Coupling | All clients coupled to one gateway | Each client has independent BFF |
| Failure | One BFF failing doesn't affect others | Isolated |

**When to use BFF:** When you have distinctly different clients (mobile, web, TV app, third-party partner) with different data requirements, connectivity constraints, and interaction patterns. Netflix has separate BFFs for iOS, Android, web, smart TV, and gaming consoles.

When NOT to use BFF: If your clients are nearly identical, a BFF per client adds operational overhead for minimal benefit.

---

### 5. What is service discovery in microservices?

"Service discovery is the mechanism by which **services find each other's network locations** (IP + port) dynamically, without hardcoding addresses.

In a containerized environment, service IPs change every time a container restarts or scales. Service A can't have `payment-service.example.com:8080` hardcoded — that IP won't be valid tomorrow.

Two flavors: **Client-side discovery** — the service registry (Consul, Eureka) returns a list of healthy instances, and the client does its own load balancing (Netflix uses this with Ribbon + Eureka). **Server-side discovery** — the load balancer queries the registry and routes traffic (Kubernetes Service + kube-proxy does this). Kubernetes makes server-side discovery the default — you just call `payment-service:8080` and k8s handles the rest."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Swiggy, Zomato, Razorpay, any company using microservices

#### Indepth
Service registry options:
- **Consul (HashiCorp):** Service discovery + health checking + key-value store. Multi-region support.
- **etcd:** Distributed key-value store. Used internally by Kubernetes.
- **Eureka (Netflix):** Peer-to-peer service registry. AP-focused — prioritizes availability over consistency.
- **Kubernetes DNS:** In k8s, every Service gets a DNS name automatically (`service-name.namespace.svc.cluster.local`). The simplest form of service discovery.

**Health checking:** The registry needs to know which instances are healthy. Methods: HTTP endpoint (`/health` returns 200), TCP check, TTL (instance must renew its registration every N seconds or be removed).

---

### 6. What is the Saga pattern?

"The Saga pattern manages **distributed transactions across multiple microservices** without using 2-phase commit (2PC).

The problem: Placing an order requires: (1) Reserve inventory, (2) Charge payment, (3) Schedule delivery. If payment fails after inventory is reserved, you need to unreserve the inventory. But these are 3 separate services with 3 separate databases — you can't do a single ACID transaction.

Saga solution: Define the sequence as a series of local transactions with **compensating transactions** for rollback. If step 2 fails, step 2's compensation runs, then step 1's compensation runs. Either all steps succeed (Saga completes) or you roll back each step in reverse order."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart, Uber, Razorpay — companies with complex transactional workflows

#### Indepth
Two Saga implementation styles:

**Choreography-based Saga:**
- Each service publishes events and reacts to events from others
- `OrderService` → publishes `OrderCreated` → `InventoryService` listens → reserves stock → publishes `StockReserved` → `PaymentService` listens → charges → publishes `PaymentCompleted`
- Pro: Simple, no central orchestrator. Con: Hard to trace, circular event dependency risk.

**Orchestration-based Saga:**
- A central `OrderSaga` orchestrator issues commands to each service
- Saga sends `ReserveStock` to Inventory → on success, sends `ChargePayment` to Payment → on success, sends `ScheduleDelivery` to Delivery
- On failure at any step, sends compensating commands in reverse
- Pro: Easy to trace, clear ownership. Con: Orchestrator risks becoming a bottleneck.

Tools: Temporal.io (workflow orchestration), AWS Step Functions (serverless saga), Axon Framework (Java saga).

---

### 7. What is a circuit breaker pattern?

"A circuit breaker prevents a failing service from bringing down the entire system by **stopping calls to that service when it's unhealthy**.

Without it: Service A calls Service B which is slow, every request to A takes 30 seconds waiting for B's timeout, thread pool exhausts, A goes down too. Cascading failure.

With circuit breaker: After 5 consecutive failures to B, the circuit breaks (opens). For the next 30 seconds, A immediately returns an error/fallback instead of calling B. After 30 seconds, the circuit goes 'half-open' — allows one test request through. If it succeeds, circuit closes. If it fails, stays open. At Netflix, Hystrix was the gold-standard circuit breaker. Today, Resilience4j is the Java standard."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Flipkart, Swiggy — essential for microservices reliability

#### Indepth
Circuit breaker states:
1. **Closed (normal):** All requests pass through. Failure rate is tracked.
2. **Open (failing):** After threshold (e.g., 50% failure rate in 10 seconds), circuit opens. All requests immediately fail fast.
3. **Half-open (testing):** After timeout, allow N requests through. If they succeed, close the circuit. If they fail, reopen.

**Resilience4j configuration:**
```java
CircuitBreakerConfig config = CircuitBreakerConfig.custom()
    .failureRateThreshold(50)          // Open when 50% requests fail
    .waitDurationInOpenState(30s)      // Stay open 30s before half-open
    .permittedCallsInHalfOpenState(5)  // Allow 5 test calls in half-open
    .build();
```

**Fallback strategies:**
- Return cached data (last successful response)
- Return a default/degraded response ("Payment system is busy, please try again")
- Fail silently (for non-critical paths like recommendations)

---

### 8. What is the strangler fig pattern?

"The Strangler Fig is a migration pattern for safely transitioning from a **monolith to microservices** by gradually replacing functionality — like a strangler fig vine that slowly replaces a host tree.

The process: (1) Put an API Gateway/facade in front of the monolith. (2) Identify the first service to extract. (3) Build the new microservice. (4) Update the gateway to route that functionality to the new service. (5) Remove the code from the monolith. Repeat until the monolith is dead.

The beauty: at no point do you do a 'big bang' rewrite. The system is always live. The monolith shrinks gradually. This is how Shopify, Netflix, and Amazon all migrated. The alternative — 'let's rewrite everything' — almost never works."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Companies currently migrating from monoliths

#### Indepth
Strangler fig implementation steps:
1. **Asset capture:** Map all API endpoints and their business functionality
2. **Choose the first seam:** Extract a domain with clear boundaries and low inter-dependencies (user profiles, notifications — not core order processing)
3. **Shadow mode:** Run new service in parallel with monolith, compare outputs — confidence building
4. **Traffic shifting:** 5% → 25% → 50% → 100% traffic shift to new service using feature flags
5. **Remove from monolith:** Delete the old code path
6. **Repeat**

Risks: **Dual writes** (writes must go to both old and new system during transition). **Data migration** (historical data in monolith DB must be accessible by new service). **Cross-boundary transactions** (saga patterns needed).

---

### 9. What is the outbox pattern?

"The outbox pattern solves the **dual-write problem** in microservices: how do you atomically update your database AND publish an event to Kafka?

The naive approach: Commit DB transaction, then publish to Kafka. If the service crashes between the two steps, the DB is updated but no event is published — downstream services are out of sync.

Outbox solution: In the same DB transaction as your business update, also write the event to an `outbox` table. A separate **transactional outbox poller** (or Debezium CDC) reads new records from the outbox table and publishes them to Kafka. If it fails, it retries until published, then marks the outbox record as processed. Since the outbox write is part of the original transaction, it's atomic — either both the business data and the outbox record are written, or neither is."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** PhonePe, Razorpay, Groww — fintech companies that can't afford message loss

#### Indepth
Outbox pattern implementation:
```sql
-- Transaction:
BEGIN;
UPDATE orders SET status = 'PAID' WHERE id = 123;
INSERT INTO outbox (event_type, payload, status) 
  VALUES ('OrderPaid', '{"orderId": 123}', 'PENDING');
COMMIT;
-- At-least-once delivery across DB and Kafka guaranteed
```

**CDC (Change Data Capture) approach:** Debezium watches the outbox table's binary log and publishes events to Kafka automatically. Zero polling overhead, millisecond latency.

**Idempotency at the consumer:** Since the outbox pattern provides at-least-once delivery (a retry may publish the same event twice), consumers must be idempotent. Use a unique `event_id` and a `processed_events` table to deduplicate.

---

### 10. What is distributed tracing and why is it essential in microservices?

"Distributed tracing lets you follow a single request as it propagates across **multiple microservices** — seeing exactly where time was spent and where errors occurred.

In a monolith, you have a stack trace. In microservices, a user request might touch 8 services. Without tracing, when a request is slow, you have no idea which service is the bottleneck. With tracing (Jaeger, Zipkin, AWS X-Ray), you get a 'trace' — a timeline showing each service hop, its duration, and any errors.

Each service propagates a `trace-id` header (like `X-B3-TraceId`) so all logs from all services for one request share the same ID. I can then search for `traceId=abc123` and see the full journey."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** All companies using microservices

#### Indepth
Key concepts:
- **Trace:** The end-to-end journey of one request (has a unique `trace_id`)
- **Span:** One unit of work within a trace (one service call, one DB query). Parent-child relationships form a tree.
- **Context propagation:** Trace ID is passed via HTTP headers (`traceparent` in W3C standard) or gRPC metadata

OpenTelemetry stack:
- **Collector:** Receives spans from services
- **Instrumentation libraries:** Auto-instrument HTTP clients, DB clients (OTel SDKs for Go, Java, Python)
- **Backend:** Jaeger (open source), Tempo (Grafana), AWS X-Ray, Datadog APM

**The three pillars of observability:**
1. **Metrics:** What is the system doing? (QPS, error rate, p99 latency) → Prometheus + Grafana
2. **Logs:** What happened? (Event log with timestamps) → ElasticSearch + Kibana / Loki
3. **Traces:** How did it happen? (Cross-service request journey) → Jaeger / Zipkin

---

### 11. How do you handle authentication and authorization in microservices?

"The standard approach: **Authentication at the gateway, authorization at the service**.

The API Gateway validates the JWT token — checks signature, expiry, issuer. It then injects user identity claims as HTTP headers to downstream services (`X-User-Id`, `X-User-Role`). Services trust these headers and apply their own authorization rules without re-validating the token.

This avoids the N×N problem: without a gateway, every service needs to validate JWTs, maintain OAuth libraries, and call the Auth service. With a gateway, auth logic exists in one place.

For authorization, I use **RBAC (Role-Based Access Control)** for coarse-grained access and **ABAC (Attribute-Based Access Control)** for fine-grained policies (e.g., 'user can only edit their own orders')."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Razorpay, PhonePe, Groww, all security-sensitive companies

#### Indepth
Auth patterns in microservices:
1. **Centralized auth (gateway pattern):** Gateway validates JWT, injects user context. Services are stateless — no session.
2. **Service-to-service auth (mTLS):** Each service has a client certificate. Mutual TLS ensures both parties verify each other's identity. Used in Istio service mesh.
3. **OAuth 2.0 + OIDC:** Industry standard. Authorization Code flow for web apps, Client Credentials for service-to-service.

JWT claims in headers (after gateway decodes):
```
X-User-Id: 12345
X-User-Email: dhruv@example.com
X-User-Roles: admin,customer
```
Services extract these headers without needing to validate the JWT signature themselves.

**OPA (Open Policy Agent):** Decentralizes authorization logic. Instead of baking auth rules into each service, services query OPA's policy engine with context. Policy changes don't require redeployment of services.

---

### 12. What is a service mesh and when do you need one?

"A service mesh is a **dedicated infrastructure layer** that handles all service-to-service communication concerns — like mTLS, load balancing, circuit breaking, distributed tracing, and retries — without changing application code.

The mesh works via a **sidecar proxy** (Envoy) injected alongside every service pod. All traffic in and out of a pod goes through this sidecar. The sidecars are collectively managed by a **control plane** (Istio, Linkerd).

I'd use a service mesh when: I have 20+ services and I don't want to implement retry/circuit-breaking/mTLS logic in every service individually. The mesh makes these cross-cutting concerns declarative — I configure them in YAML, not in code. Uber, Lyft, and Netflix use service meshes at massive scale."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Google, Uber — platform engineering roles

#### Indepth
Service mesh capabilities:
| Feature | Without Mesh | With Mesh (Istio/Linkerd) |
|---------|-------------|--------------------------|
| mTLS | Each service implements | Transparent, auto-rotated certs |
| Retries | Code in each service | VirtualService YAML config |
| Circuit breaking | Resilience4j in each service | DestinationRule YAML config |
| Load balancing | Client-side Ribbon | Sidecar handles it |
| Distributed tracing | OTel SDK in each service | Automatic (sidecar instruments) |
| Traffic splitting | Code-level feature flags | Canary YAML (10% to v2) |

**Istio components:**
- **Envoy sidecar:** Data plane. Intercepts all traffic.
- **Istiod (control plane):** Pushes configuration to sidecars, manages certificate rotation.
- **Virtual Services:** Define traffic routing rules (canary, A/B testing).
- **Destination Rules:** Define load balancing policy, circuit breaking config per service.

**Cost:** Service mesh adds ~50MB memory per pod (Envoy) and 2-5ms latency per hop. Only justified at scale (50+ services).
