# 🚀 Microservices — Deployment & Resilience Patterns (Product Companies)

> **Level:** 🔴 Senior
> **Asked at:** Uber, CRED, Razorpay, Flipkart, Swiggy, Zepto, PhonePe, Netflix

---

## Q1. How do you implement a distributed rate limiter that works correctly across 50 API Gateway nodes?

"A naive per-node in-memory rate limiter will fail because each of your 50 gateway nodes tracks its own counter. A user can bypass the limit by simply routing requests across different nodes — 50 requests per node means 2,500 effective requests per second.

A **distributed rate limiter** requires a shared, atomic counter. My standard approach:

**Implementation using Redis + Lua Script (Token Bucket or Sliding Window Log):**

Instead of the simple INCR command (which has a race condition between the check and the increment), I use a **Lua script executed atomically on a Redis node**. Lua scripts in Redis are guaranteed to run atomistically.

```lua
-- Sliding Window Counter using Redis
local key = KEYS[1]          -- e.g., "rate_limit:user:12345"
local window = tonumber(ARGV[1]) -- e.g., 60 seconds
local limit = tonumber(ARGV[2])  -- e.g., 100 requests

local current = tonumber(redis.call('GET', key) or '0')
if current >= limit then
    return 0  -- REJECTED
else
    redis.call('INCR', key)
    redis.call('EXPIRE', key, window)
    return 1  -- ALLOWED
end
```

**Algorithms to know:**
1. **Token Bucket:** A bucket fills at a fixed rate. Each request takes a token. Allows short bursts.
2. **Sliding Window Log:** Store timestamps of each request. Count items within the window. Most accurate but memory-heavy.
3. **Fixed Window Counter:** Simplest. Resets every window. Has a boundary burst problem (2x traffic allowed at window edges).

For fintech and APIs with strict SLAs, the **sliding window** approach is usually preferred.

**Redis Cluster considerations:** If you run Redis Cluster, you must ensure all requests for the same user's rate limit key route to the same shard using hash tags: `{user:12345}:rate_limit`."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, CRED, Uber — payment and high-traffic API platforms where rate limiting is a core business requirement, not just a nice-to-have.

#### Indepth
**Leaky Bucket:** A variant where requests enter a queue (the bucket) and are processed at a fixed output rate. It smooths out bursty traffic perfectly. The downside is higher latency for burst requests as they wait in the queue, and requests must be dropped when the queue is full.

---

## Q2. Explain Blue/Green, Canary, and Rolling deployments. When do you use each?

"All three solve the problem of deploying new code to production without taking the system down (Zero Downtime Deployment), but they have different risk profiles and rollback strategies.

**1. Rolling Deployment:**
Gradually replace old instances with new ones. If you have 10 pods, update 2 at a time: 8 old + 2 new → 6 old + 4 new → ... → 10 new.
- *Pros:* Requires no extra infrastructure. Default K8s RollingUpdate strategy.
- *Cons:* During the rollout, **two versions coexist**. If your API has a breaking change (e.g., a database column was renamed), old and new pods are simultaneously processing requests, causing data inconsistency issues.
- *Use when:* Backward-compatible updates with no schema changes.

**2. Blue/Green Deployment:**
Maintain two identical production environments: Blue (live) and Green (idle). Deploy the new version to Green, run smoke tests, then switch the Load Balancer to point 100% traffic to Green. Blue stays idle for instant rollback.
- *Pros:* Instant rollback (flip the LB back). No mixed-version traffic.
- *Cons:* Requires 2x the infrastructure cost during the transition. Database migrations must be backward-compatible (both Blue and Green shared the same DB).
- *Use when:* Major releases, database schema changes, or compliance-required instant rollback capability.

**3. Canary Deployment:**
Route a small percentage of live production traffic (e.g., 5%) to the new version while 95% still goes to the old version. Monitor error rates, latency, and business metrics. Gradually increase the percentage if metrics look good.
- *Pros:* Real user validation. Limits the blast radius of a bad deployment. Allows A/B testing.
- *Cons:* More complex routing logic (usually requires a Service Mesh like Istio or feature-flag system). Your application must handle two versions simultaneously.
- *Use when:* High-risk features, testing at scale, or iterating on user-facing changes."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Swiggy, Uber, Flipkart — companies running continuous deployment pipelines and needing robust rollback mechanisms.

#### Indepth
**Feature Flags (as a Deployment Strategy):** Considered an evolution beyond Canary. The new code is deployed to 100% of servers, but the feature is hidden behind a flag that's only enabled for a specific percentage of users (or specific user IDs, or internal users only). This decouples deployment from release and gives product teams, not just ops, control over rollout — without a re-deploy.

---

## Q3. What is Event Sourcing? How is it different from just storing the current state in a database?

"In a standard database (`CRUD model`), you store the **current state** of an entity. When an order goes from `CREATED` to `PAID` to `SHIPPED`, you UPDATE the `orders` table row. The history of what happened is lost.

**Event Sourcing** inverts this. Instead of storing the current state, you store an **immutable, append-only log of all events** that happened to that entity.

```
events table:
----------------------------------
| order_id | event_type     | payload                    | timestamp  |
|----------|----------------|----------------------------|------------|
| O-123    | OrderCreated   | {user: ..., items: ...}    | 10:00:01   |
| O-123    | PaymentReceived| {amount: 500, method: UPI} | 10:00:05   |
| O-123    | OrderShipped   | {tracking_id: TRK456}      | 10:02:00   |
```

The **current state** is derived by replaying all events from the beginning (or from a snapshot). This is the `OrderAggregate`.

**Key advantages:**
1. **Complete Audit Trail:** Every change is recorded. For a bank transaction, you can tell exactly what happened, when, and in what sequence. This is legally required in fintech.
2. **Temporal Queries:** 'What was the state of this order at 10:00:03?' — Just replay events up to that timestamp.
3. **Event Replay for New Systems:** If you add a new Analytics Service next year, you can replay the entire event history from Day 1 and populate the analytics database retrospectively.
4. **Natural Fit with CQRS:** Event Sourcing is almost always used with CQRS. Events are published to Kafka; read-model consumers build optimized query databases from them.

**The downside:** Querying current state requires replaying events (mitigated by periodic Snapshots). The concept is also harder for most developers to reason about than simple CRUD."

#### 🏢 Company Context
**Level:** 🔴 Senior/Principal | **Asked at:** Fintech companies (Razorpay, PhonePe, Groww) where auditing, compliance, and rollback of financial state are critical non-functional requirements.

#### Indepth
**Event Sourcing vs. CQRS vs. Outbox Pattern:** These are three distinct concepts often confused.
- **Outbox Pattern:** Solves dual-write atomicity (write to DB + publish to Kafka atomically).
- **CQRS:** Separates read models from write models.
- **Event Sourcing:** Replaces the write model itself — instead of updating state, you append events.
You can use CQRS without Event Sourcing (separate read/write DBs, still CRUD-based). But Event Sourcing almost always implies CQRS is also used.

---

## Q4. How do you implement Chaos Engineering? What is the philosophy behind it?

"The core philosophy of Chaos Engineering, pioneered by Netflix, is: **'If you don't intentionally break your system in a controlled way, production will break it in an uncontrolled way.'**

Traditional testing proves that things work when the happy path is followed. Chaos Engineering deliberately injects failures to prove your system handles them gracefully.

**The scientific method:**
1. **Define a Steady State:** What does 'normal' look like? (e.g., 99.9% requests succeed with p99 latency < 200ms, and checkout completion rate is 98%)
2. **Form a Hypothesis:** 'If the Payment Service pod on Node-3 crashes, checkouts will still succeed because the circuit breaker routes to a healthy Payment Service pod.'
3. **Inject Failure:** Kill the pod, inject network latency (500ms), create a CPU spike, corrupt DNS responses.
4. **Observe:** Did the steady state hold? Did the metrics deviate?
5. **Learn and Fix:** Harden the system. Automate the test into your CI/CD pipeline.

**Tools:**
- **Chaos Monkey (Netflix):** Randomly terminates VM instances in production.
- **Chaos Toolkit / Litmus (for Kubernetes):** Inject pod kills, network partitions, and resource exhaustion into K8s workloads.
- **Gremlin:** Commercial platform for structured chaos experiments.

**What to test:**
- Pod/node failure (killed instances)
- Network latency and packet loss between services
- Dependency timeouts (3rd party API goes down)
- Disk I/O saturation
- Memory pressure (OOM killer scenarios)"

#### 🏢 Company Context
**Level:** 🔴 Senior/SRE | **Asked at:** Netflix-style companies, Hotstar, Uber — organizations that invest in SRE and platform reliability. This question signals maturity in handling production incidents.

#### Indepth
**GameDay:** A structured event where SRE and Development teams deliberately run chaos experiments during business hours with everyone watching. It's a practice drill — like a fire drill for infrastructure — that builds confidence in incident response playbooks and tooling before a real outage occurs.

---

## Q5. How do you implement distributed tracing across 8 microservices? Walk me through OpenTelemetry.

"When a user's checkout request fails after 3 seconds, how do you know which of the 8 microservices in its path is responsible?

**Distributed Tracing** creates a unified view of a single request as it flows across all services. The concepts:

- **Trace:** The entire journey of one user request. Identified by a globally unique `TraceId`.
- **Span:** One unit of work within that trace (e.g., 'Order Service processing', 'DB query in Order Service'). Each span has its own `SpanId` and records start time, end time, and metadata (status codes, errors).
- **Parent-Child:** Spans have parent-child relationships. 'Order Service calls Payment Service' means the Order span is the parent of the Payment span.

**OpenTelemetry (OTel) is the standard for implementation:**
1. **Instrumentation:** You add the OTel SDK to each microservice (auto-instrumentation via agents for Java, or manual for custom spans). It automatically captures HTTP, gRPC, and DB calls.
2. **Propagation:** OTel automatically injects `traceparent` headers into every outgoing HTTP/gRPC request. The receiving service extracts it and continues the same trace.
3. **Collector:** Each service sends spans to a local OTel Collector sidecar, which batches, filters, and exports them to your backend (Jaeger, Zipkin, Grafana Tempo, or Datadog).
4. **Visualization:** In Jaeger UI, you search by `TraceId` and see a Gantt chart of every service hop, database query, and its latency.

**In Spring Boot:**
```xml
<dependency>
  <groupId>io.micrometer</groupId>
  <artifactId>micrometer-tracing-bridge-otel</artifactId>
</dependency>
```
Zero-code configuration: Spring Boot 3.x auto-instruments HTTP calls and propagates the `TraceId` automatically."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Platform engineering and SRE interviews at Flipkart, Swiggy, Amazon — any company with more than 10 microservices where debugging production issues requires more than reading individual log files.

#### Indepth
**Sampling:** In a high-traffic system, tracing every single request generates enormous data volumes. **Head-based sampling** makes a random decision at the start (trace 1% of traffic). **Tail-based sampling** buffers spans and only exports traces that ended in an error or exceeded a latency threshold — far smarter and more useful for debugging.
