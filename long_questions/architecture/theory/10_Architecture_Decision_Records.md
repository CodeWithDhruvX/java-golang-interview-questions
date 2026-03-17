# 📋 Architecture Decision Records & Trade-off Analysis — Questions 1–10

> **Level:** 🔴 Senior
> **Asked at:** Amazon (Principal/Staff SDE), Google (L5+), senior architecture roles, engineering manager interviews

---

### 1. What is an Architecture Decision Record (ADR)?

"An ADR is a **document that captures an important architectural decision** — the context that led to it, the decision itself, and the consequences. It's the 'commit history' for architectural choices.

The problem it solves: Six months after adopting Kafka for messaging, a new engineer asks 'why do we use Kafka instead of SQS?' Nobody remembers. The architect who made the decision has moved on. The knowledge is lost. With an ADR, future team members can understand the reasoning without reconstructing it from scratch.

Format: short, structured. Title, Status (proposed/accepted/deprecated), Context (what forced this decision), Decision (what we decided), Consequences (good and bad outcomes). Stored in the repository alongside the code."

#### 🗣️ How to Explain in Interview
**Interviewer:** What is an Architecture Decision Record (ADR)?
**Your Response:** "An ADR is a short, version-controlled document that captures an essential architectural choice and the **'why'** behind it. I view it as the **collective memory** of the engineering organization. We’ve all encountered complex configurations where we wonder, 'Why was this specific path taken?' If the original architect has moved on, the team is left guessing. 

By maintaining a `/docs/adr` folder in the repo, we preserve the constraints, alternatives considered, and most importantly, the **trade-offs** that were accepted at the time. This drastically speeds up senior onboarding and prevents the team from circular arguments every six months. It’s not about bureaucracy; it’s about providing the necessary context for whoever has to operate and evolve the system years down the line."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Engineering leadership, principal engineer interviews

#### Indepth
ADR template:
```markdown
# ADR-007: Use Kafka for Order Events

## Status
Accepted (2024-03-15)

## Context
Our Order Service needs to notify 5+ downstream services (Inventory, Payment, Delivery,
Notification, Analytics) when an order is placed. Current approach uses synchronous HTTP
calls — any downstream failure blocks order confirmation. Downstream services can't be 
added without modifying Order Service.

## Decision
Adopt Apache Kafka as the message broker for order events. Order Service publishes 
`OrderPlaced` events to a Kafka topic. Each downstream service subscribes independently.

## Alternatives Considered
- **AWS SQS + SNS fan-out:** Simpler, managed, but vendor lock-in and no event replay.
- **RabbitMQ:** Less operational overhead than Kafka, but no event replay, lower throughput ceiling.
- **Synchronous REST chains:** Current approach. Tight coupling, cascading failures.

## Consequences
✅ Downstream services are independently deployable and scalable  
✅ Event replay supports new services catching up on historical data  
✅ Decoupled failure domains — payment outage doesn't block order confirmation  
❌ Eventual consistency — downstream state may lag by milliseconds  
❌ Operational complexity — Kafka cluster management required  
❌ Added consumer group monitoring and dead letter queue management  
```

---

### 2. What is trade-off analysis in architecture?

"Every architectural decision involves trade-offs — gaining one quality attribute at the cost of another. There are no free lunches in architecture. The architect's job is to make these trade-offs **explicit, reasoned, and aligned with business priorities**.

Classic trade-offs: Consistency vs Availability (CAP theorem), Performance vs Security (encryption adds overhead), Simplicity vs Scalability (monolith is simpler, microservices scale better), Consistency vs Latency (synchronous replication is consistent but slow), Flexibility vs Standards (custom protocol is optimal, HTTP/REST is universally understood).

When asked 'what would you do?' in an architecture interview, the red flag answer is a confident single solution with no acknowledgment of trade-offs. The green flag answer is 'it depends — here are the trade-offs, and given X constraint, I'd choose Y'."

#### 🗣️ How to Explain in Interview
**Interviewer:** What is trade-off analysis in architecture?
**Your Response:** "In software architecture, **there are no perfect solutions, only trade-offs.** Every gain in a quality attribute like Availability usually comes at a cost in another, such as Consistency or absolute Performance. My job as a senior architect isn't to find the 'best' tool in a vacuum, but the one that aligns most closely with our specific business drivers.

When I evaluate a design, I look at both **First-Order and Second-Order consequences.** Adding a cache (First-Order) improves latency, but it also increases complexity and introduces the risk of stale data (Second-Order). I make these trade-offs explicit in my design docs. If a stakeholder wants both absolute real-time accuracy and extreme high scale at low cost, I have to explain the technical reality: we can optimize for two, but we’re always trading off the third."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Staff/Principal engineer interviews where there's no 'correct' answer

#### Indepth
Framework for trade-off analysis (ATAM — Architecture Trade-off Analysis Method):
1. **Identify quality attributes:** Which NFRs matter most? (availability, performance, security, maintainability)
2. **Identify architectural drivers:** What constraints force certain decisions? (team size, budget, existing infrastructure, compliance)
3. **Map decisions to quality attributes:** For each major decision, analyse its impact on each quality attribute
4. **Identify sensitivity points:** Decisions that have large impact on quality (shard key choice hugely impacts scalability)
5. **Identify trade-off points:** Decisions that positively impact one quality and negatively impact another
6. **Prioritize:** Given business priorities, which trade-offs are acceptable?

Example matrix for messaging choice:
| Quality Attribute | Kafka | SQS | RabbitMQ |
|------------------|-------|-----|----------|
| Throughput | +++ | ++ | + |
| Operational simplicity | - | +++ | ++ |
| Message replay | +++ | - | - |
| At-least-once guarantee | +++ | +++ | +++ |
| Vendor lock-in | ✅ (open source) | ❌ (AWS) | ✅ |
| Cost (managed) | High (self-hosted) | Low (pay-per-message) | Medium |

---

### 3. How do you evaluate build vs buy vs open source decisions?

"Build vs Buy vs Open Source is a recurring architectural decision for every new capability. The right answer depends on: **core competency** (is this component core to your competitive advantage?), **time to market** (how fast do you need it?), **total cost of ownership** (licensing, maintenance, operational), and **control** (how much customization do you need?).

Rule of thumb: if it's a solved problem (logging, monitoring, API gateway), buy or use open source. If it's core to your business differentiation (recommendation algorithm, fraud detection model), build it. Building a general-purpose search engine when Elasticsearch exists is vanity engineering.

Specific consideration: open source has zero license cost but non-zero operational cost (running Kafka is not free when you account for engineering time to operate, monitor, and upgrade it). AWS MSK turns Kafka into a managed service — you pay more per message but your engineers focus on business logic."

#### 🗣️ How to Explain in Interview
**Interviewer:** How do you evaluate build vs buy vs open source decisions?
**Your Response:** "I follow a simple heuristic: **'Build for differentiation, buy for utility.'** If a component is core to our company's competitive advantage—like a proprietary matching algorithm—we build it. But for commodity features like Authentication, Email, or API Gateways, we buy or leverage mature open-source solutions.

The biggest trap I see is underestimating the **Total Cost of Ownership (TCO)** of open source. A 'free' license doesn't mean a free service; you must account for the engineering hours spent on patching, scaling, and on-call rotations. I often advocate for a 'SaaS-first' approach for infrastructure (like AWS managed services) so that our top engineers can focus 100% of their energy on shipping features that drive revenue, rather than managing a Postgres cluster."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Engineering leadership, staff/principal roles

#### Indepth
Framework:
| Factor | Build | Buy (SaaS) | Open Source |
|--------|-------|-----------|-------------|
| Core competency? | Yes → Build | No → Don't Build | Depends |
| Time to market | Slowest | Fastest | Medium |
| Customization | Full | Limited by vendor | High |
| Ongoing cost | Engineering time | Subscription | Operational overhead |
| Vendor dependency | None | High | None |
| Compliance | Full control | Vendor's compliance | Full control |
| Team expertise | Builds internal | Outsourced | Team owns operations |

Real examples:
- **Authentication:** Almost never build. Use Auth0, Cognito, or open-source Keycloak. Auth is too critical, too complex, and too regulatory to get wrong.
- **Search:** Use Elasticsearch if standard queries. Build custom only if you have ultra-specific search semantics (e.g., Google's web crawler + ranking).
- **Payment processing:** Use Stripe/Razorpay for the payment rail. But fraud detection ML → often built internally because it's core IP.

---

### 4. How do you approach technical debt in architecture?

"Technical debt is the **accumulated cost of sub-optimal architectural decisions** made in the past — often deliberately, to move faster. Like financial debt, it accrues interest: the longer you ignore it, the more expensive it becomes to resolve.

I categorize debt by severity: (1) **Intentional debt** — consciously incurred to meet a deadline, with a plan to pay it back. Acceptable. (2) **Accidental debt** — didn't know better at the time; discovered later. Needs prioritization. (3) **Bit rot** — formerly good decisions that time has made obsolete (a technology the team has grown out of, an approach that no longer fits the scale).

The error is treating all technical debt as bad. Debt that speeds up delivery without blocking future work is good financial management. Debt that makes every new feature take 3x longer is existential."

#### 🗣️ How to Explain in Interview
**Interviewer:** How do you approach technical debt in architecture?
**Your Response:** "I view technical debt exactly like financial debt. It's often a **strategic tool**—you borrow against the future to hit a critical market window today. That’s perfectly fine as long as you account for the 'interest rate.' If a shortcut makes every future feature take 20% longer to build, that interest will eventually bankrupt your team's velocity.

I categorize debt into three buckets: **Intentional** (strategic shortcuts), **Accidental** (evolving requirements), and **Bit Rot** (aging tech stacks). I advocate for a 'Debt Register' and try to negotiate a fixed percentage of every sprint—usually 15-20%—specifically for paying down high-interest debt. It’s about keeping the system's 'cogs' clean so we can maintain a high, predictable shipping velocity over the long term."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Engineering leadership, engineering manager interviews

#### Indepth
Technical debt management strategies:
1. **Debt register:** Track debt items like you track bugs — with severity, owner, and estimated remediation cost. Make it visible to leadership.
2. **20% rule:** Dedicate 20% of every sprint to tech debt reduction. Not 100% new features, not 0% debt. Netflix and Google maintain this ratio.
3. **Strangler fig for architecture debt:** Gradually replace the debt-laden system rather than big-bang rewrites.
4. **Boy scout rule:** "Leave the campsite cleaner than you found it." Refactor as you touch code for new features.
5. **Definition of Done includes debt:** No new feature is "done" if it introduces architectural debt without a tracking issue.

The **broken windows theory** applied to code: A codebase with visible technical debt (broken windows) encourages further shortcuts. Maintaining zero-tolerance for new debt while systematically paying down old debt changes team culture.

---

### 5. What is the difference between scalability and performance?

"**Performance** is about how fast the system responds to a single request — latency, throughput for a given load. **Scalability** is about how the system's performance holds up as load increases.

A system can be performant but not scalable: A single-threaded Go HTTP server might respond in 1ms per request and handle 100 concurrent requests beautifully. Add 10,000 concurrent requests — it falls apart. It was performant at low load, but not scalable.

A system can be scalable but not performant: If adding more nodes linearly increases throughput but each request still takes 2 seconds, you're scalable but you have a performance problem.

Both matter, but they require different solutions: Performance is about algorithmic efficiency, caching, query optimization. Scalability is about statelessness, horizontal partitioning, load balancing."

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the difference between scalability and performance?
**Your Response:** "This is a fundamental distinction. **Performance** is how fast a system responds to a single request under ideal conditions—'How many milliseconds for this page to load?' **Scalability** is how that performance holds up as you add a thousand more users at the same time.

You can have a performant app that isn't scalable—like a single-node Python server that responds in 10ms but crashes at 100 concurrent connections. Conversely, you can have a scalable app that performs poorly—I can add 100 nodes to handle the load, but if every request still takes 5 seconds due to a slow DB query, we have a performance problem. At a senior level, I optimize code for performance (caching, SQL tuning), but I design architectures (statelessness, partitioning) for scalability."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** All product company interviews

#### Indepth
Amdahl's Law: The speedup from parallelization is limited by the sequential (non-parallelizable) portion of the task. If 5% of your computation is inherently sequential, you'll never exceed 20x speedup no matter how many cores you add.

Performance dimensions:
- **Latency:** Time for a single request (p50, p95, p99)
- **Throughput:** Requests per second at steady state
- **Resource efficiency:** CPU/memory per request

Scalability dimensions:
- **Linear scalability:** Doubling servers doubles throughput (ideal)
- **Sub-linear scalability:** Doubling servers gives less than 2x throughput (common — due to coordination overhead)
- **Super-linear scalability:** Rare. Usually means the system was poorly configured before scaling.

**Back-of-envelope calculations matter in interviews:** "Can a single MySQL instance handle 50K RPM?" (Yes, with proper indexing — MySQL can handle ~10K QPS on a well-tuned instance). "Can one Redis node handle 1M writes/sec?" (No — Redis is single-threaded, max ~100K ops/sec per thread). These mental models show architectural maturity.

---

### 6. How do you approach an architecture interview question?

"I use a structured framework for system design / architecture questions to ensure I cover the problem completely without jumping to a solution before understanding the requirements.

Step 1 — **Clarify requirements** (3-5 min): Ask about scale (users, QPS, data size), consistency requirements, SLAs, geographic distribution, who the clients are. Never start designing without understanding scale.

Step 2 — **Define scope** (2 min): 'For this session, I'll focus on the core order placement flow. I'll set aside notifications and analytics.' Don't try to design the entire company.

Step 3 — **High-level design** (10-15 min): Draw the components — clients, API gateway, services, databases, caches. Talk through the request flow.

Step 4 — **Deep dive** (follow interviewer's lead): Go deep on the component they care about — usually the hardest part.

Step 5 — **Identify bottlenecks and trade-offs** (5 min): What will break at 10x scale? What are the trade-offs in your design?"

#### 🗣️ How to Explain in Interview
**Interviewer:** How do you approach an architecture interview question?
**Your Response:** "My framework is **'Requirements first, Design second.'** I never touch a whiteboard until I've quantified the problem. I’ll start by asking about the Read/Write ratio, the expected p99 latency, and whether we prioritize **Consistency or Availability** given the business context.

Once I have those constraints, I walk through a high-level design to set the baseline and then proactively dive into the **bottlenecks**. For example, I might note, 'This database becomes a SPOF at 10k QPS, so here’s how we’d implement sharding or a cache layer.' Showing that you can anticipate and mitigate systemic failure is much more important than just drawing a perfect diagram. I treat it as a peer-to-peer collaboration with another engineer."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Every company with a system design round

#### Indepth
Common interview mistakes:
1. **Jumping to solution immediately:** Designing before clarifying requirements. "I'd use microservices with Kafka" before asking what the load is.
2. **Ignoring non-functional requirements:** Designing only for functionality, not for scale, availability, or consistency.
3. **No trade-offs acknowledged:** Proposing one solution as if it's the only option. Shows limited experience.
4. **Too much detail too soon:** Spending 20 minutes on the database schema before discussing the high-level architecture.
5. **Unclear about bottlenecks:** Can't identify what will fail first under load.

What interviewers actually evaluate:
- **Communication:** Can you explain complex concepts clearly?
- **Structured thinking:** Do you approach the problem methodically?
- **Trade-off awareness:** Do you acknowledge what you're giving up?
- **Scale intuition:** Do you understand what systems can handle what load?
- **Depth of knowledge:** Can you go deep when asked?

---

### 7. How do you architect for multi-tenancy?

"Multi-tenancy is designing a single system to **serve multiple isolated customers (tenants)** while sharing the same infrastructure. SaaS products are almost always multi-tenant: one Salesforce instance serves millions of companies; each company sees only its own data.

Three isolation models: **Siloed (separate everything per tenant):** Separate databases, separate deployments. Maximum isolation, highest cost, no resource sharing. Good for enterprise customers with strict data residency requirements. **Pooled (share everything):** All tenants in one shared DB. Data separated by a `tenant_id` column. Maximum resource efficiency, lowest isolation. Good for small/medium tenants. **Hybrid (bridge model):** Large tenants get dedicated databases; small tenants share a pooled database. Cost-effective and appropriately isolated."

#### 🗣️ How to Explain in Interview
**Interviewer:** How do you architect for multi-tenancy?
**Your Response:** "Multi-tenancy is about balancing **isolation** with **resource efficiency**. On one end is the 'Silo' model, where every customer gets dedicated infrastructure. It's the most secure and provides the best 'no noisy neighbor' guarantee, but it's expensive to maintain. On the other end is the 'Pool' model, where all tenants share a database and we separate data via a `tenant_id`.

I usually advocate for a **Tiered approach**. Enterprise customers might get an isolated Silo for compliance, while smaller accounts share a pooled infrastructure to keep costs down. A critical part of my design is enforcing isolation at the database layer itself using tools like **PostgreSQL Row Level Security (RLS)**, ensuring a bug in the app layer can't cross tenant boundaries."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** SaaS companies (Freshworks, Chargebee, Zoho, Postman)

#### Indepth
Data isolation strategies:
1. **Database per tenant:** Maximum isolation. Simple queries (no tenant_id filter). Expensive (thousands of DBs for thousands of tenants). Hard to aggregate cross-tenant analytics.
2. **Schema per tenant (PostgreSQL):** Shared DB server, separate schema per tenant. Good balance. Limits: PostgreSQL supports ~10K schemas efficiently.
3. **Row-level security (RLS):** Shared database and schema, row-level filter by tenant_id. PostgreSQL RLS policies enforce tenant isolation at DB level automatically. Most efficient but requires careful implementation.

Row-level security in PostgreSQL:
```sql
CREATE POLICY tenant_isolation ON orders
    USING (tenant_id = current_setting('app.tenant_id')::UUID);
ALTER TABLE orders ENABLE ROW LEVEL SECURITY;
-- All queries to "orders" table automatically filter by current tenant
```

Noisy neighbor problem: In pooled model, one large tenant's heavy queries slow down other tenants. Mitigations: Query timeouts per tenant, read replicas for heavy tenants, dedicated connection pools per tenant, rate limiting at the application layer.

---

### 8. How do you design for geographic distribution?

"Geographic distribution is the architecture for serving users across multiple regions with low latency and high availability, while handling the consistency challenges of geographically separated databases.

Active-passive multi-region: Primary region handles all writes, secondary regions are read replicas. Reads are served locally (low latency). Writes always go to primary (adds latency for remote write). Failure of primary → promote secondary to primary (30-60 seconds of disruption). Used by most companies as the first step to multi-region.

Active-active multi-region: Both regions accept writes. Requires conflict resolution (last-write-wins, CRDT, application-level resolution). Highest availability, lowest latency for writes, but most complex. CockroachDB and Cassandra support active-active natively."

#### 🗣️ How to Explain in Interview
**Interviewer:** How do you design for geographic distribution?
**Your Response:** "Geo-distribution is usually driven by two factors: **Latency** (UX) and **Data Sovereignty** (Compliance). I typically start with an 'Active-Passive' setup with local read replicas to keep latency low globally. 

If we truly need 'Active-Active'—where we accept writes in multiple regions—we face the hard limits of the speed of light. Data consistency becomes the primary challenge. I prefer strategies like **Geo-sharding**, where a user's data is 'homed' in the region closest to them. We also have to design for regional disaster recovery (DR)—if an entire AWS region goes dark, we need a battle-tested plan to pivot traffic without causing massive data corruption."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Companies with global user base or data sovereignty requirements

#### Indepth
Data sovereignty considerations: GDPR requires EU user data to remain in the EU. India's data localization requirements (DPDP Act 2023) require certain data categories to remain in India. Multi-region architecture must ensure user data is stored in the appropriate region — routing layer must ensure an Indian user's PII doesn't accidentally go to a US region.

Regional failover (Route 53 / Cloudflare):
```
Normal: DNS → Mumbai region (primary)
During failover: Health check fails → DNS → Singapore region (secondary)
Failover time: DNS TTL + health check polling interval (typically 60-90 seconds)
```

**Geo-routing:** Route users to the nearest region by IP geolocation. Indian users → Mumbai. SEA users → Singapore. EU users → Frankfurt. This reduces latency by 50-200ms depending on the geography.

**Data replication lag:** In active-passive, replication lag from Mumbai to Singapore might be 50-100ms. Reads from Singapore during normal operation may return data that's 100ms stale. For eventual consistency workloads (social feeds, product catalog), acceptable. For financial data, must route reads to primary.

---

### 9. What is event-driven microservices vs transactions?

"This is the fundamental consistency tension in microservices: **ACID transactions** are easy with a single database (all-or-nothing, strongly consistent), but impossible to achieve natively across services with separate databases.

The trade-off: use **synchronous communication + distributed transactions (2PC/Saga)** if you need strong consistency, accepting higher latency and tighter coupling. Or use **event-driven architecture with eventual consistency**, accepting that state across services will converge over time.

My default: prefer eventual consistency via event-driven architecture for most business operations. Very few business processes actually require immediate strong consistency across services. 'Order placed' can be eventually consistent — it's fine if inventory is updated 200ms later. 'Money transferred' requires much stronger guarantees — but still solvable with Saga + idempotency rather than 2PC."

#### 🗣️ How to Explain in Interview
**Interviewer:** What is event-driven microservices vs transactions?
**Your Response:** "In a monolith, you have ACID transactions that make consistency 'easy.' In microservices, once we split our data, we lose those atomic guarantees. My default is to embrace **Event-Driven Architecture** and accept **Eventual Consistency** for non-critical paths.

Instead of forcing a distributed transaction that would hurt our availability, we use the **Saga Pattern**. For an order flow, Service A performs its local transaction and publishes an 'OrderCreated' event. If a downstream service (like Payment) fails, it publishes a 'Compensation Event' to undo the work of the first service. It’s more complex to implement, but it’s the only way to build a decoupled, highly available system at scale."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Fintech companies, e-commerce, any transactional system design

#### Indepth
When to use each:
| Scenario | Approach | Reason |
|---------|----------|--------|
| Place order + reserve inventory | Saga (eventual) | Acceptable 200ms lag |
| Fund transfer (internal) | Saga + idempotency | No 2PC needed if idempotent |
| User profile update | Eventual consistency | Staleness of seconds acceptable |
| Booking reservation (seats) | Strong consistency | No double-booking allowed |
| Payment authorization | Synchronous + strong | Cannot be eventually consistent |

**The Saga vs 2PC decision:**
- **2PC (Two-Phase Commit):** Coordinator asks all participants to "prepare" → if all agree, "commit". One participant blocking = all blocked. Highly coupled, poor availability.
- **Saga:** Short local transactions with compensating transactions. Eventually consistent. Better availability. Works across service boundaries without a distributed coordinator.

Google Spanner achieves distributed ACID without 2PC using **TrueTime** (atomic clocks and GPS) to order transactions globally. Expensive, but enables globally consistent transactions at scale.

---

### 10. How do you handle breaking changes in distributed systems?

"Breaking changes in distributed systems are dangerous because **not all consumers upgrade simultaneously**. During the rollout window, old and new versions of consumers run side-by-side. Your API/event schema must support both.

The expand-contract pattern: *Expand* — add the new field/behavior alongside the old (support both versions). *Migrate* — update all consumers to use the new format. *Contract* — remove the old field/behavior (now nothing depends on it).

This applies to: API field renames, DB schema changes, message format changes, configuration changes. The rule: never remove before confirming all consumers have migrated. Blue-green deployments help but don't eliminate the problem — external consumers on old contracts still need support."

#### 🗣️ How to Explain in Interview
**Interviewer:** How do you handle breaking changes in distributed systems?
**Your Response:** "In distributed systems, you can't assume that all consumers will upgrade at the same time. I always follow the **'Expand and Contract'** pattern to handle breaking changes safely. If I need to rename an API field, I don't just swap it; I first 'Expand' by adding the new field while maintaining support for the old one.

Only after I’ve verified that all consumers have migrated to the new field do I 'Contract' and remove the old one. I also use **Consumer-Driven Contracts** to catch breaking changes in our CI/CD pipelines before they reach production. Versioning (like `/v1/` to `/v2/`) is our last resort for massive structural shifts, ensuring that we never break a partner's integration unexpectedly."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Platform teams, public API teams, companies with many internal consumers

#### Indepth
Breaking change patterns by scenario:

**API response field rename (`userName` → `fullName`):**
1. Add `fullName` alongside `userName` (both present)
2. Deprecate `userName` in docs, monitor usage
3. After grace period, remove `userName`

**Database column rename (`user_name` → `full_name`):**
1. Add `full_name` column (nullable)
2. Deploy app to write to both columns
3. Backfill `full_name` from `user_name`
4. Deploy app to read from `full_name` only
5. Remove `user_name` column

**Kafka event schema change (add required field):**
1. Add field with a default value (backward compatible)
2. Update all producers to populate the field
3. Update all consumers to use the field
4. Remove the default (field is now truly required)

Never: add a new required field without a default — consumers reading old events that lack the field will fail.

**Feature flags for breaking changes:** Deploy the new behavior behind a feature flag. Enable for internal testing, then 1%, 10%, 100% of traffic. Instant rollback by toggling the flag without redeployment.
