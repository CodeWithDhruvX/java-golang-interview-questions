## 🔸 Advanced Distributed Systems (Questions 401-410)

### Question 401: How would you handle leader election in distributed systems?

**Answer:**
*   **Algorithms:**
    *   **Bully Algorithm:** Higher ID wins.
    *   **Raft/Paxos:** Use a consensus cluster (Etcd/ZooKeeper) to elect a leader. 
    *   **Lease:** Nodes try to acquire a lock in DB/Redis. Winner is Leader for X seconds.
*   **Scenario:** Master-Slave DB, Job Schedulers.

### Question 402: Design a gossip-based messaging protocol.

**Answer:**
*   **Concept:** Nodes pick K random neighbors and share info periodically (like a virus spreading).
*   **Use:** Cluster membership (Cassandra), Failure detection (SWIM).
*   **Mechanism:** `Node A -> Node B: "I suspect Node C is dead"`. `Node B` propagates. If `Node C` doesn't refute, it's marked Dead.

### Question 403: How would you achieve quorum-based consensus?

**Answer:**
*   **Quorum:** `(N/2) + 1`.
*   **Write:** Must succeed on W nodes.
*   **Read:** Must read from R nodes.
*   **Condition:** `W + R > N` guarantees intersection (Pigeonhole Principle).
*   **Example:** DynamoDB, Cassandra.

### Question 404: Explain vector clocks and how you’d use them.

**Answer:**
(See Q103). Captures causality.
*   `[A:1, B:0]` happens before `[A:1, B:1]`.
*   `[A:1, B:0]` and `[A:0, B:1]` are concurrent (conflict).

### Question 405: How to detect split-brain scenarios?

**Answer:**
(See Q106).
*   **Quorum Check:** If a partition can't see the majority (e.g., 3 out of 5), it steps down. "I am in the minority, I cannot be leader."

### Question 406: How to handle write skew in distributed databases?

**Answer:**
*   **Write Skew:** An anomaly where two transactions read overlapping data, make disjoint writes, but violate constraints when combined. (e.g., On-call roster: "At least one person must be on call").
*   **Fix:**
    *   **Materialize Conflict:** Create a dummy row used as a lock.
    *   **SSI (Serializable Snapshot Isolation):** DB detects conflict at commit time.
    *   **For Update:** `SELECT * ... FOR UPDATE`.

### Question 407: Design a system using Raft consensus algorithm.

**Answer:**
*   **Roles:** Leader, Follower, Candidate.
*   **Log Replication:**
    1.  Client -> Leader: writes `cmd`.
    2.  Leader -> Followers: `AppendEntries(cmd)`.
    3.  Followers -> Leader: `ACK`.
    4.  If Majority ACKs: Leader commits -> Client Success.

### Question 408: How do you manage schema evolution in distributed systems?

**Answer:**
*   **Format:** Avro / Protobuf (Schema Registry).
*   **Rules:**
    *   **Forward Compatible:** Old code can read New data (Reader ignores unknown fields).
    *   **Backward Compatible:** New code can read Old data (Reader handles missing fields with defaults).
*   **Registry:** Kafka Schema Registry enforces compatibility checks on Produce.

### Question 409: How would you optimize for low tail latency?

**Answer:**
(Tail Latency: The slowest 1% of requests).
*   **Hedged Requests:** Send request to 2 replicas. Use the first response. Cancel the other.
*   **Jitter:** Retry quickly if P99 threshold crossed.
*   **Resource Isolation:** Ensure GC pauses or noisy neighbors don't starve the thread.

### Question 410: Explain read-repair mechanism in eventual consistency.

**Answer:**
*   **Scenario:** Client reads from 3 nodes (Quorum read).
    *   Node A: `v2`
    *   Node B: `v2`
    *   Node C: `v1` (Stale)
*   **Repair:** Coordinator compares versions. Returns `v2` to client. Asynchronously sends `v2` to Node C to fix it.

---

## 🔸 Real-Time & Event-Driven Systems (Questions 411-420)

### Question 411: Design a real-time gaming leaderboard system.

**Answer:**
(See Q191). Redis Sorted Set.

### Question 412: How would you build a high-frequency trading system?

**Answer:**
*   **Latency:** Minimize network hops. Colocation (Server in same building as Exchange).
*   **OS:** Kernel bypass (DPDK/Solarflare) to read packets directly from NIC to Userspace (Skip TCP/IP stack overhead).
*   **Language:** C++ / Rust / Java (Zero GC).
*   **Logic:** Simple algorithms (Arbitrage). No DB calls. Everything in RAM.

### Question 413: How would you ensure ordering in an event stream?

**Answer:**
(See Q122). Partition Key.
*   Kafka guarantees order *within a partition*. All events for `Order_123` must go to Partition 5.

### Question 414: What’s the difference between event sourcing and event streaming?

**Answer:**
*   **Event Sourcing (Storage Pattern):** DB structure. Store `[Deposited, Withdrew]` instead of `Balance=10`. Replayable source of truth.
*   **Event Streaming (Transport Pattern):** Moving data via Kafka. You can stream events without using event sourcing.

### Question 415: Design a telemetry system for autonomous vehicles.

**Answer:**
*   **Volume:** TBs of data.
*   **Strategy:**
    *   **Critical:** Stream immediately (Crash detected) via 5G/MQTT.
    *   **Non-Critical:** Buffer on local SSD. Upload bulk via Wi-Fi when car is charging at night.
*   **Compression:** High efficiency (Delta encodings).

### Question 416: How would you handle late-arriving events?

**Answer:**
(Stream Processing problem).
*   **Watermarks:** A heuristic asserting "No events older than T will arrive now".
*   **Windowing:** "Hourly Window".
*   **Late Data:**
    *   *Allowed:* Update the already-computed window (if within 1 hour).
    *   *Dropped:* Determine cutoff (e.g., > 1 day late).

### Question 417: How do you design time-window-based analytics?

**Answer:**
*   **Tumbling Window:** Non-overlapping. `[12:00-12:05]`, `[12:05-12:10]`.
*   **Sliding Window:** Overlapping. "Last 5 mins, updated every 1 min".
*   **Session Window:** "While user is active + 30m timeout".

### Question 418: How to handle out-of-order events?

**Answer:**
*   **Buffer:** Stream processor holds events in memory buffer.
*   **Sort:** Re-order based on `EventTime` (not `ProcessingTime`).
*   **Watermark:** Trigger processing when Watermark passes.

### Question 419: Build a real-time bidding system for ads.

**Answer:**
(See Q389).

### Question 420: How do you debounce vs throttle events at scale?

**Answer:**
*   **Throttle (Rate Limit):** "Max 1 request per second". (Drop extras).
*   **Debounce (Delay):** "Wait for 1s of silence". If user types 'A', wait. Types 'AB', reset timer. Types 'ABC'... stop. Wait 1s -> Send 'ABC'.
*   **Implementation:** Client-side via JS. Server-side via Redis keys expiration.

---

## 🔸 Cross-Region and Global Systems (Questions 421-430)

### Question 421: Design a global distributed file system.

**Answer:**
(e.g., Dropbox / Google Drive).
*   **Metadata:** Global DB (Spanner) or Region-sharded DB.
*   **Block Store:** S3 buckets in each region.
*   **Edge:** Caches popular blocks.
*   **Sync:** Differential sync (rsync style).

### Question 422: How to replicate user data across continents?

**Answer:**
*   **Async Replication:** Master in US. Slaves in EU/Asia. Fast write, potential stale read in EU.
*   **Multi-Master (CRDTs):** Write to local region. Background sync merges connection. (DynamoDB Global Tables).
*   **Geo-Partitioning:** EU users strictly pinned to EU DB. (GDPR compliance).

### Question 423: How to design a geo-aware DNS routing system?

**Answer:**
*   **Route53 (AWS):** Geolocation Routing Policy.
*   **Logic:** "If IP is from France, value=1.2.3.4 (Paris LB). If IP from US, value=5.6.7.8 (N.Virginia LB)."
*   **Failover:** If Paris health check fails, route France users to London.

### Question 424: What’s the role of Anycast in global systems?

**Answer:**
*   **Unicast:** 1 IP = 1 Server.
*   **Anycast:** 1 IP = 50 Servers globally.
*   **Routing:** BGP routes user to the topological nearest server.
*   **Benefit:** DDoS mitigation (attack traffic is decentralized/diluted), Low latency without complex DNS.

### Question 425: How to maintain GDPR compliance in multi-region architecture?

**Answer:**
*   **Data Residency:** PII must stay in EU.
*   **Tagging:** Column level tags: `pii=true`.
*   **Filtering:** Replication job checks tags. `if pii && dest != EU` -> skip.
*   **Pseudonymization:** Tokenize PII in US; keep mapping only in EU.

### Question 426: How would you geo-fence user data storage?

**Answer:**
*   **API Gateway:** Inspects JWT `country` claim.
*   **Routing:**
    *   `US-User` -> `us-east-1` DB.
    *   `DE-User` -> `eu-central-1` DB.
*   **Middleware:** Rejects request if `US-User` tries to write to `EU-DB`.

### Question 427: How to reduce cross-region data transfer costs?

**Answer:**
*   **Cache:** Don't fetch from US DB if data is static. Cache in Local Region (Redis).
*   **Batching:** Send large compressed batches instead of chatty calls.
*   **Proxy:** S3 Transfer Acceleration or CloudFront (uses AWS backbone, cheaper than public internet transfer sometimes).

### Question 428: Design a worldwide leaderboard that updates in near-real-time.

**Answer:**
*   **Approach:** Count-Min Sketch (Approximate) or Hierarchical aggregation.
*   **Hierarchy:**
    1.  City Servers aggregate counts (Redis).
    2.  Country Servers aggregate City counts.
    3.  Global Server aggregates Country counts (Global Top 100).
*   **Latency:** Global board lags by 5-10s.

### Question 429: How would you architect a global media delivery service?

**Answer:**
(Netflix CDN).
*   **Open Connect:** Place storage appliances inside ISP data centers.
*   **Prediction:** Pre-fetch "Stranger Things" to ISP box at 4 AM local time.
*   **Serving:** User streams from their own ISP (Latency < 5ms).

### Question 430: What’s the design of a follow-the-sun support platform?

**Answer:**
*   **Goal:** 24/7 support.
*   **Shift Handoff:**
    *   Shift 1 (India) ends.
    *   Tickets in "Open" state are re-assigned to Shift 2 (Europe).
*   **DB:** Single Global DB (or replicated) so Europe sees India's notes immediately.

---

## 🔸 Privacy-First & Regulated Systems (Questions 431-440)

### Question 431: How would you design a system to log data access without logging the data itself?

**Answer:**
*   **Log:** `User=Alice Action=Read ResourceID=123`.
*   **Anti-Pattern:** Do NOT log `Payload={"cc": "4111..."}`.
*   **Hashing:** If needed, log `Hash(Payload)` to prove integrity without revealing data.

### Question 432: Design a healthcare record system compliant with HIPAA.

**Answer:**
*   **Encryption:** At Rest (KMS) + In Transit.
*   **Audit:** Every view must be logged.
*   **Auth:** MFA mandatory.
*   **BAA:** Use only cloud services that sign a Business Associate Agreement.

### Question 433: How to tokenize PII data and still allow searching?

**Answer:**
*   **Deterministic Encryption:** `Enc("Alice")` always equals `Xy9z...`.
*   **Search:** Index `Xy9z...`.
*   **Query:** User searches "Alice" -> App encrypts to `Xy9z...` -> DB finds match.
*   **Trade-off:** Susceptible to frequency analysis attacks (if many people are named "Alice").

### Question 434: What’s differential privacy and how would you implement it?

**Answer:**
*   **Goal:** Share aggregate stats (e.g., "Avg Salary") without revealing individual data ("Bob's Salary").
*   **Method:** Add mathematical noise (Laplace noise) to result.
    *   True Avg: $50,000.
    *   Reported: $50,023.
*   **Privacy Budget:** Limit number of queries to prevent noise cancellation.

### Question 435: Design a zero-knowledge authentication system.

**Answer:**
*   **Concept:** Prove I know the password without sending the password.
*   **Protocol:** SRP (Secure Remote Password) protocol.
*   **Verifier:** Server stores `v = g^x`. Client proves knowledge of `x`.

### Question 436: How would you audit access patterns without violating user privacy?

**Answer:**
*   **Anonymization:** Strip UserIDs from logs used for analytics. Replace with `SessionID` (hashed).
*   **Aggregation:** Report "100 users viewed page" rather than "User A, B... viewed page".

### Question 437: How to build a KYC/AML-compliant data platform?

**Answer:**
(Know Your Customer / Anti-Money Laundering).
*   **Identity:** Store ID documents (Passport) in encrypted S3 (Object Lock).
*   **Checks:** Run name against Sanctions List (OFAC).
*   **Risk:** Score user (High/Med/Low). High requires manual review.

### Question 438: Design a parental control system with granular rules.

**Answer:**
*   **Policy:** `User(Child) -> Block(Category=Gambling, Time=22:00-08:00)`.
*   **Enforcement:**
    *   DNS Level (NextDNS): Resolve `badsite.com` to `0.0.0.0`.
    *   App Level: Local VPN on device intercepts traffic.

### Question 439: How do you design audit-only access for sensitive systems?

**Answer:**
*   **Break-Glass Protocol:**
    *   Admin usually has NO access.
    *   Emergency: Admin requests access. Manager approves.
    *   Access granted for 1 hour.
    *   ALL actions screen-recorded and logged to security team.

### Question 440: How to comply with "right to data portability"?

**Answer:**
*   **Export API:** `GET /user/export`.
*   **Format:** Standard machine-readable (JSON/CSV).
*   **Scope:** Photos, Posts, Friends.
*   **Delivery:** Async generation -> Email secure download link (Password protected zip).

---

## 🔸 Hybrid Architectures & Integration (Questions 441-450)

### Question 441: Design a cloud-bursting architecture between on-prem and AWS.

**Answer:**
*   **VPN/Direct Connect:** Bridge On-prem and VPC.
*   **Scaling:**
    *   Base load handled by On-prem (Fixed Cost).
    *   Spike load triggers K8s Autoscaler to add nodes in AWS EKS.
*   **State:** Database must be accessible to both (latency limits). Ideally stateless app tier bursts.

### Question 442: How would you integrate a legacy monolith into a new microservices stack?

**Answer:**
*   **Strangler Fig Pattern:**
    1.  Place API Gateway in front of Monolith.
    2.  Build new "Order Service".
    3.  Gateway routes `/orders` to New Service. `/users` to Monolith.
    4.  Repeat until Monolith is gone.

### Question 443: How do you handle API contracts between polyglot services?

**Answer:**
*   **Standard:** Protobuf / OpenAPI / Thrift.
*   **Code Gen:** Generate Java Client and Go Server from same `.proto` file.
*   **Registry:** Central repo versioning the contracts.

### Question 444: Design a system to sync SaaS app data across multiple tenants.

**Answer:**
(e.g., Slack Shared Channels - Company A talks to Company B).
*   **Bridge:** A specific "Shared Context" database.
*   **Copy:** Message in Company A is copied to Shared DB -> Synced to Company B.
*   **Permissions:** Union of policies.

### Question 445: How to build a unified notification system with email, SMS, and push?

**Answer:**
*   **Abstraction:** `sendNotification(User, Template, Params)`.
*   **router:** `User.Preferences` says "Email for Marketing, SMS for OTP".
*   **Providers:**
    *   Email -> SES.
    *   SMS -> Twilio.
    *   Push -> FCM.

### Question 446: How would you migrate a live database without downtime?

**Answer:**
1.  **Dual Write:** App writes to Old + New DB.
2.  **Backfill:** Batch copy Old data to New (ignoring updates already handled by Dual Write).
3.  **Verify:** Check counts/checksums.
4.  **Read Switch:** App reads from New.
5.  **Write Switch:** App writes only to New. Stop Old.

### Question 447: Design a hybrid cloud-native + edge deployment model.

**Answer:**
(e.g., Retail Stores).
*   **Central:** AWS manages Catalogue, Analytics.
*   **Edge:** Store Server (Kubernetes K3s) runs POS (Point of Sale).
*   **Sync:** Edge operates offline. Syncs transactions to Cloud when internet is up.

### Question 448: Build an adapter layer for integrating 3rd-party APIs.

**Answer:**
*   **Facade Pattern:** Create `IPaymentGateway` interface.
*   **Adapters:** `StripeAdapter`, `PayPalAdapter` wrap messy external API logic.
*   **Resilience:** Wrap calls in Hystrix/Resilience4j (Timeout, Retry, Circuit Breaker).
*   **Testing:** WireMock external APIs to test Adapter logic.

### Question 449: How to version inter-service contracts in gRPC?

**Answer:**
*   **Package Versioning:** `package com.myco.order.v1;` in `.proto`.
*   **Breaking Change:** Create `order.v2` package.
*   **Server:** Implement both `OrderServiceV1` and `OrderServiceV2`.
*   **Client:** Upgrade independently.

### Question 450: Design a feature rollout system that integrates web and mobile.

**Answer:**
*   **Central Server:** Feature Flag Service (LaunchDarkly).
*   **Mobile:** SDK fetches flags on app launch (`enable_dark_mode=true`). Caches it.
*   **Web:** API fetches flags on Session Start.
*   **Consistency:** Ensure User ID hashing is consistent across platforms so User A sees the same feature on Phone and Desktop.
