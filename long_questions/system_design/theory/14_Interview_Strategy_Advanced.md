# 🔴 Interview Strategy, Estimation & Advanced Concepts — Questions 131–140

> **Level:** 🟢 All levels — Interview framework, estimation, and advanced topics not covered elsewhere
> **Asked at:** All companies — these are meta-skills and advanced concepts that separate good candidates from great ones

---

### 131. How to approach a system design interview?
"My interview framework is: **Clarify → Estimate → Architect → Deep Dive**. This structured approach prevents the #1 mistake candidates make: jumping to a solution before understanding the problem.

**Clarify (5 min):** Ask about scale (DAU, QPS, data size), functional requirements (what does the system do?), non-functional requirements (latency SLA, consistency requirements, availability target), and constraints (read-heavy or write-heavy? global or regional?). Never design blindly — a 10-user internal tool and a 10M-user consumer app have radically different architectures.

**Estimate (5 min):** Back-of-envelope calculations. Storage per day, peak QPS, bandwidth. These numbers inform every subsequent decision — a system with 1K QPS doesn't need Kafka; a system with 1M QPS probably does.

**High-Level Architecture (10-15 min):** Draw the main components on the whiteboard — clients, LB, API servers, cache, DB, message queue. Explain the flow of a request end-to-end for the most critical use case.

**Deep Dive (15-20 min):** The interviewer will pick one or two areas to explore deeply — database design, the most complex component, a specific failure scenario. This is where senior candidates shine."

#### 🏢 Company Context
**Level:** All levels | **Asked at:** Every company — this is the meta-skill for interviews

#### Indepth
Detailed interview structure (45-minute session):

```
0-5 min:   Requirements clarification
           - Who are the users?
           - What are the core features? (MVP scope)
           - Non-functional: scale (DAU, QPS, storage), SLA (latency <100ms?), availability (99.99%?)
           
5-10 min:  Capacity estimation (very rough is fine)
           - Daily active users × actions per user × data per action = storage/day
           - Peak QPS = total daily requests / seconds_in_day × peak_factor (3-5x)
           
10-25 min: High-level design
           - Draw the request path for the primary use case
           - Identify the key components (API gateway, services, DB, cache, queue, CDN)
           - Explicitly call out trade-offs you're making
           
25-45 min: Deep dive (interview-directed)
           - Database schema design
           - Scaling the hardest bottleneck
           - Failure scenarios and recovery
           - Specific algorithm/data structure (consistent hashing, etc.)
```

**Common mistakes to avoid:**
- Not clarifying requirements → over-engineer or under-engineer
- Jumping straight to tech choices → "I'll use Kafka" before knowing the scale
- Going silent → narrate your thinking even when exploring options
- Not discussing trade-offs → surface-level answers lack depth
- Over-designing → if 10K QPS, you don't need sharding; a single DB with read replica is fine

**What interviewers look for:** Problem decomposition, ability to handle ambiguity, knowledge of trade-offs (not just solutions), ability to change direction based on constraints, communication clarity.

---

### 132. How to do back-of-envelope estimation?
"Back-of-envelope (BOE) estimation is approximating system scale with quick mental math — getting to 'right order of magnitude' without a calculator.

The goal isn't precision. A system that handles 1M QPS vs 10M QPS needs fundamentally different architectures. Estimating within 1 order of magnitude is sufficient.

My approach: round aggressively (10M users, not 8.7M), use powers of 10 for mental math, and know a few key numbers by heart: 1 million seconds ≈ 11 days, 1 billion seconds ≈ 31 years, daily QPS ≈ daily-events / 10^5 (100K seconds in a day — round 86400)."

#### 🏢 Company Context
**Level:** All levels | **Asked at:** Every system design interview — Amazon particularly loves this; they have the WORKING BACKWARDS mechanism

#### Indepth
Numbers every engineer should know:
| Resource | Approximate Values |
|---|---|
| L1 cache reference | 1 ns |
| L2 cache reference | 4 ns |
| RAM read | 100 ns |
| Read 1MB from RAM | 0.1 ms |
| SSD random read | 0.1 ms |
| Network round trip (same DC) | 0.5 ms |
| HDD seek | 10 ms |
| Network round trip (cross-DC) | 30-150 ms |

Estimation template — Design Twitter:
- 200M DAU
- Average user posts 0.1 tweets/day = 20M tweets/day
- 20M / 86400 = ~230 tweets/second (writes)
- Read:write ratio = 100:1 → 23,000 reads/second
- Tweet size: 1KB → 20M tweets × 1KB = 20GB/day, 7TB/year
- Storage with media: assume 10% tweets have media, avg 200KB → 400GB/day

This tells me: write is trivial for one DB. Reads at 23K QPS need Redis caching. Storage needs planning after year 5. Typical tweet volume doesn't need Kafka for the tweet write path (though it does for fan-out).

---

### 133. What are the SOLID principles in system design?
"SOLID is a set of five principles for writing maintainable, scalable object-oriented code — but they apply equally at the service and system level.

**S — Single Responsibility:** A service should do one thing well. A User service handles users; a Payment service handles payments. Mixing concerns into one service creates a distributed monolith.

**O — Open/Closed:** Systems should be open for extension, closed for modification. Adding a new payment provider shouldn't require changing existing code — add a new provider plugin that implements the `PaymentProvider` interface.

**L — Liskov Substitution:** Any implementation of an interface should be substitutable without breaking the system. Any SQL database (MySQL/Postgres/SQLite) should be pluggable through a common `DatabaseInterface`.

**I — Interface Segregation:** Don't force clients to implement interfaces they don't use. A read-only service shouldn't be forced to implement write methods.

**D — Dependency Inversion:** High-level modules shouldn't depend on low-level modules. Both should depend on abstractions (interfaces). Your order handler depends on a `PaymentServiceInterface`, not the concrete `StripeService`."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** LLD (Low-Level Design) rounds — Flipkart, Amazon, Uber interview rounds focused on OOP design

#### Indepth
SOLID at the microservice level (not just OOP):
- **SRP → Service per business capability:** Payment Service, Order Service, Notification Service. Each evolves independently.
- **OCP → Plugin architecture:** Adding a new DeliveryPartner (Dunzo, Swiggy Genie, Shadowfax) adds a class implementing `DeliveryPartner` interface — no change to core order logic.
- **LSP → Consistent APIs across services:** All product-category services implement the same `CatalogService` interface. API gateway can route transparently.
- **ISP → Consumer-driven contracts:** Service A should only depend on the subset of Service B's API it uses. Pact (contract testing) formalizes this.
- **DIP → Message passing over direct service calls:** Service A depends on `EventQueue` interface, not `KafkaProducer` directly. Swap Kafka for SQS with zero business code change.

These principles lead to the same conclusion as Domain-Driven Design (DDD): model your services and code around business capabilities, build clear boundaries, and use interfaces/events to decouple.

---

### 134. What is the strangler fig pattern?
"The Strangler Fig pattern is a migration strategy for gradually replacing a monolith with microservices — borrowing from the strangler fig tree, which grows around a host tree and eventually replaces it entirely.

The approach: instead of a big-bang rewrite (risky — you're rewriting a running system, which is notoriously hard to do correctly), you **incrementally extract functionality** from the monolith piece by piece.

Put a proxy (API gateway) in front of the monolith. New features are built as independent microservices, which the proxy routes to. Existing features are gradually extracted to services — one endpoint at a time. The monolith shrinks as services grow. Eventually, the monolith is retired."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Companies modernizing legacy systems — TCS, Infosys client projects, Flipkart's historic Oracle-to-microservices migration, any legacy modernization project

#### Indepth
Strangler fig implementation steps:
1. **Add façade/proxy:** Route all traffic through an API gateway in front of the monolith. Clients never call the monolith directly — always through the gateway.
2. **Identify the first extraction candidate:** Choose a bounded context that's: high value (frequently changed), loosely coupled to the rest of the monolith, well-understood. Start with the simplest extraction.
3. **Build new service alongside monolith:** The new service can co-exist — monolith still handles this feature. Test new service in parallel.
4. **Switch traffic:** Update the gateway to route this feature's traffic to the new service. Monitor for errors. Keep monolith as fallback.
5. **Remove from monolith:** Delete the code, tables, and logic from the monolith for this feature.
6. **Repeat** for the next feature.

**Pitfalls:**
- **Shared database:** If monolith and new service share the same DB, you haven't achieved true decoupling. New service should have its own DB, with a migration period where data is dual-written.
- **Choosing too large a first extraction:** First extraction should be small and low-risk. Success builds confidence and teaches the team the process.
- **Premature stop:** Many teams extract 3-4 services and stop. The monolith remains but is now a 'distributed monolith' since it still shares some state. Commit to full extraction or don't start.

---

### 135. What is event-driven architecture?
"Event-driven architecture (EDA) is a design paradigm where **services communicate by producing and consuming events** rather than making direct synchronous calls.

An event is an immutable fact: 'OrderPlaced', 'PaymentProcessed', 'InventoryDeducted'. The producer emits the event and doesn't wait for or care about consumers. Multiple consumers can react to the same event independently. This creates **loose coupling** — the Order service emits `OrderPlaced` and doesn't need to know that three services (Inventory, Notification, Fulfillment) consume it.

Kafka is the backbone of most EDA systems. Events are durable (retained for days/weeks), replayable (replay historical events to rebuild a service), and ordered within a partition."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Swiggy, Zomato, Amazon, Flipkart — any microservices system at scale

#### Indepth
EDA patterns:
- **Event Notification:** Lightweight event to tell other services something happened. Carries minimal data (just the ID). Consumer must fetch full details via API. Simple but causes additional API calls.
- **Event-Carried State Transfer:** Event carries the full current state of the entity. Consumer doesn't need to call back — they can update their own store. Higher data volume but fewer round trips.
- **Event Sourcing:** Persist ALL events (don't store current state, store the log of changes). Current state = replay of all events. Full audit trail, time-travel debugging. DB becomes the event log (Kafka). Complexity: state reconstruction, schema evolution of historical events.
- **CQRS + Event Sourcing:** Commands (writes) → events → state. Queries → projections. Each projection optimized for its read pattern. Powerful but complex.

EDA benefits:
- Loose coupling (services don't know about each other)
- Scalability (consumers scale independently)
- Resilience (consumer failures don't affect producer)
- Auditability (event log is a complete history)

EDA challenges:
- Eventual consistency — consumers lag behind
- Difficult to trace flow across multiple services (requires distributed tracing with trace IDs in event headers)
- Testing is harder (need to spin up Kafka, multiple services)
- Event schema evolution (backward-compatible schema changes with Avro/Protobuf + schema registry)

---

### 136. What is service mesh?
"A service mesh is a **dedicated infrastructure layer** for handling service-to-service communication in microservices — offloaded from application code into a sidecar proxy alongside each service.

Without a service mesh: every service must implement retry logic, circuit breaking, mTLS, service discovery, load balancing, and distributed tracing itself — duplicating complex infrastructure code across 50+ services in 5+ languages.

With a service mesh (Istio + Envoy): the sidecar proxy (Envoy) intercepts all in/out traffic and handles these concerns transparently. The application code is business logic only. The mesh control plane (Istio) configures all sidecars centrally."

#### 🏢 Company Context
**Level:** 🔴 Senior / Principal | **Asked at:** Companies operating large Kubernetes-based microservices — Swiggy, Meesho, Google, Amazon, Lyft (Envoy creators)

#### Indepth
Service mesh capabilities:
- **Traffic management:** A/B testing (route 5% traffic to v2), canary releases, fault injection for chaos testing, retries, timeouts, circuit breaking — all configured via CRDs (Custom Resource Definitions), no code changes.
- **Security (mTLS):** All service-to-service communication automatically encrypted with mutual TLS. Service identity attested by certificates (SPIFFE/SPIRE). No more per-service TLS certificate management.
- **Observability:** Every request is automatically traced (trace ID propagated through mesh). Metrics (request volume, latency, error rate) emitted per service pair. Grafana dashboards without any application instrumentation code.
- **Access control:** Policy: "Only the Order Service is allowed to call the Payment Service." Enforced at the network level by Envoy sidecars — even if a misconfigured service tries to call Payment directly, the sidecar blocks it.

Istio architecture:
- **Data plane:** Envoy proxy sidecar in every pod (iptables redirects all traffic through sidecar)
- **Control plane:** Istiod — configures all Envoy sidecars via xDS (discovery service protocol). Manages certificate rotation, service discovery, traffic routing rules.

Cost: Sidecar adds ~10ms latency per hop (two sidecar hops per service call = ~20ms overhead). CPU and memory overhead per sidecar. Operational complexity of managing Istio itself. Worth it at >20-30 services. Overkill for small microservices deployments.

---

### 137. What is a content moderation system?
"Content moderation at scale (social media, UGC platforms) requires a combination of automated ML filtering and human review — no single approach works alone.

The pipeline: user submits content → **pre-moderation** (ML classifier scores safety probability) → if confidence is high (>95% safe or >95% violating), auto-approve or auto-reject → if confidence is low (ambiguous), route to **human review queue** → human makes final decision → outcome feeds back to improve the ML model.

The key metrics: **precision** (of flagged content, what % truly violates?) and **recall** (of truly violating content, what % did we catch?). Optimize precision to protect user experience (false positives remove good content). Optimize recall to protect safety (false negatives let violating content through)."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Meta, Twitter/X, TikTok, ShareChat, MX Player, YouTube — any UGC platform

#### Indepth
Content moderation system components:
1. **Input processing:** Text, images, videos, audio each require different ML models. Image → CNN-based image classifier (NSFW, violence detection). Text → NLP transformer model (hate speech, spam). Video → frame-by-frame image classification + audio transcription.
2. **ML classification pipeline:**
   - Hashing (PhotoDNA for known CSAM content): exact hash match → immediate block, no ML needed
   - Perceptual hashing: fuzzy match for near-duplicates of known bad content
   - ML classification: confidence score → threshold routing
3. **Human review tooling:** Queue of flagged content. Reviewer makes approve/reject decision. Decision is logged with reviewer ID, decision time, grounds for decision.
4. **Appeals system:** User can appeal a moderation decision. Goes to senior reviewer queue.
5. **Feedback loop:** Human reviewer decisions → training data → retrain models weekly.

Scale challenge: YouTube receives 500 hours of video per minute. Manual review of everything is impossible. ML automation handles 99%+ of content; humans focus on the ambiguous 1% and appeals.

Content moderation is also a **trust and safety** domain involving legal compliance (GDPR's right to explanation for automated decisions, DPDP Act in India), regional rules (content legal in the US, illegal in Germany), and platform policy enforcement.

---

### 138. What is a recommendation system?
"A recommendation system predicts **what content, products, or people a user will want to see next** — personalized to that individual based on their past behavior and similarity to other users.

Two foundational approaches: **Collaborative Filtering** (what similar users liked → recommend to this user) and **Content-Based Filtering** (user likes products with these attributes → recommend more products with similar attributes).

Netflix/Amazon use hybrid approaches: collaborative filtering identifies 'users similar to you liked X', content-based ensures diverse recommendations (not just everything from the same director), and deep learning models (neural collaborative filtering) learn complex interaction patterns that linear models miss."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (product recommendations), Netflix (content), Flipkart (product), Swiggy (restaurant), Spotify (music) — any personalization-heavy product

#### Indepth
Recommendation pipeline:
1. **Candidate Generation:** From millions of items, generate a shortlist of candidates for this user (~1000 items). Fast and approximate. Method: nearest neighbor in embedding space (Matrix Factorization, Word2Vec-style item embeddings using user interaction history).

2. **Scoring/Ranking:** Score each candidate with a more complex model. Logistic regression or a DNN trained on features: user history, item features, contextual signals (time of day, device, location). Output: probability of engagement per candidate.

3. **Post-processing:** Diversity (don't recommend 10 items from same category), freshness boost (slightly up-rank new items), business rules (demote irrelevant items, boost sponsored), deduplication (remove already-seen items).

4. **Cold Start Problem:** New users have no history → collaborative filtering fails. Mitigate with: onboarding flow (ask user preferences), demographic-based defaults (popular in your city), most popular items by category.

Matrix Factorization (ALS — Alternating Least Squares): Decompose user-item interaction matrix into user embedding matrix U (M×K) and item embedding matrix V (N×K). Predicted rating for user u, item i = U[u] · V[i] (dot product). Train by minimizing squared error over known ratings. K (latent factors) = 64-256 typically.

---

### 139. What is a distributed cache?
"A distributed cache is a **cache system that spans multiple nodes**, providing a shared, fast data store for all instances of a horizontally scaled application.

When you have 20 app server instances, you can't use in-process cache — each instance has its own private memory and would have different cached data (cache inconsistency). A distributed cache (Redis Cluster, Memcached cluster) provides a single shared cache that all 20 instances use consistently.

Redis Cluster: data is sharded across 16 nodes using consistent hashing of 16384 hash slots. Each node is responsible for a subset of slots. Replication: each primary node has at least one replica. If a primary fails, its replica promotes. Clients connect using the Redis Cluster protocol which handles slot routing."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Any company with a horizontally scaled application — Swiggy, Flipkart, Amazon, Netflix

#### Indepth
Redis Cluster architecture:
```
Client → Smart Client (slot routing)
         ↓ hash_slot = CRC16(key) % 16384
         ↓ slot 5000 → Node 3?
         
         Node 1: slots 0-5460  (+ Replica 1)
         Node 2: slots 5461-10922 (+ Replica 2)
         Node 3: slots 10923-16383 (+ Replica 3)
       
If client routes to wrong node → MOVED error → redirect
If node is down → ASK error → redirect to replica
```

Redis Sentinel (HA without clustering):
- Separate to Redis Cluster (which also shards)
- Sentinel monitors Redis primary/replica pair
- If primary fails, Sentinels elect new primary among replicas (majority vote)
- Clients connect to Sentinel first, get current primary IP
- Used when you want HA but don't need horizontal sharding (dataset fits in one server)

**Cache key design:** 
- Use namespaced keys: `user:{userId}:profile`, `order:{orderId}:status`. Prevents collisions within a shared Redis.
- Consider key cardinality for memory: `user:*` with 100M users = 100M cache keys. Budget memory accordingly.
- Short keys: `u:{id}:p` vs `user:{userId}:profile` saves bytes × millions of keys = significant memory saving in large clusters.

---

### 140. What is a real-time notification system?
"Real-time notifications need to deliver messages to users **immediately when an event occurs** — WhatsApp's 'message delivered' tick, Swiggy's 'your order is out for delivery', stock price alerts.

The delivery mechanism depends on whether the user is online: (1) **Online users** receive notifications via WebSocket or long-polling. (2) **Offline users** receive push notifications (APNs for iOS, FCM for Android, Web Push for browsers).

The pipeline: event occurs (order status changes) → Notification Service receives event from Kafka → determines notification type and recipients → checks if user is online (via WebSocket server presence registry in Redis) → if online, push via WebSocket → if offline, send push notification via APNs/FCM."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Swiggy, Zomato, Amazon, WhatsApp (Meta), Razorpay (payment alerts), any real-time user-facing product

#### Indepth
Real-time notification system architecture:

```
Event (OrderStatusChanged) → Kafka → Notification Service
                                           ↓
                              User presence check (Redis)
                                           ↓
                    ┌──────────────────────┼──────────────────────┐
                    ▼                      ▼                      ▼
              User ONLINE           User OFFLINE            Multiple Devices
              WebSocket push         FCM / APNs               Push to all
              (instant)             push notification         registered tokens
                                    (may be silenced
                                     by OS/app rules)
```

Notification delivery guarantees:
- **At-most-once:** Firebase FCM semantics — if delivery fails, no retry (acceptable for non-critical notifications like promotional offers).
- **At-least-once + deduplication:** Store notification ID in DB. If client receives duplicate (due to retry), idempotent dedup on `notification_id` prevents showing it twice.

Notification preferences:
- Users can opt out of specific notification categories. Store preferences per user per category in DB.
- Respect quiet hours: user set 'no notifications 10pm-8am' → notification service checks preference + user timezone before delivering.
- Batching: instead of sending 20 low-priority notifications (engagement emails), batch into one daily digest.

**Scaling WebSocket for notifications:**
- Long-lived WebSocket connections → each server handles thousands.
- `connection_registry: { user_id → {server_id, socket_id} }` in Redis.
- Notification service asks Redis "which server is user X connected to?", sends message to that server via Redis Pub/Sub, which then pushes to the client.
- Horizontal scaling: add WebSocket servers, each registers/deregisters user connections in Redis.
