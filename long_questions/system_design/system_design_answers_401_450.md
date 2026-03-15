## 🔸 Advanced Distributed Systems (Questions 401-410)

### Question 401: How would you handle leader election in distributed systems?

**Answer:**
*   **Algorithms:**
    *   **Bully Algorithm:** Higher ID wins.
    *   **Raft/Paxos:** Use a consensus cluster (Etcd/ZooKeeper) to elect a leader. 
    *   **Lease:** Nodes try to acquire a lock in DB/Redis. Winner is Leader for X seconds.
*   **Scenario:** Master-Slave DB, Job Schedulers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you handle leader election in distributed systems?

**Your Response:** "I'd use different approaches based on the scenario. For simple cases, the Bully Algorithm where the node with the highest ID wins works well. For production systems, I'd use consensus algorithms like Raft or Paxos with a cluster of Etcd or ZooKeeper nodes to elect a leader.

Another approach is using leases - nodes compete to acquire a lock in Redis or the database, and whoever gets the lock becomes leader for a fixed time period. This is useful for master-slave databases or job schedulers where we need exactly one leader coordinating work. The key is ensuring that if the leader fails, another node can quickly take over to maintain system availability."

### Question 402: Design a gossip-based messaging protocol.

**Answer:**
*   **Concept:** Nodes pick K random neighbors and share info periodically (like a virus spreading).
*   **Use:** Cluster membership (Cassandra), Failure detection (SWIM).
*   **Mechanism:** `Node A -> Node B: "I suspect Node C is dead"`. `Node B` propagates. If `Node C` doesn't refute, it's marked Dead.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you design a gossip-based messaging protocol?

**Your Response:** "I'd design it like how information spreads in a social network - each node periodically picks K random neighbors and shares what it knows. This creates an epidemic-like spread of information throughout the cluster.

It's perfect for cluster membership in systems like Cassandra or failure detection using protocols like SWIM. For example, if Node A suspects Node C is dead, it tells Node B, which tells other nodes. If Node C doesn't refute the suspicion quickly, the cluster marks it as dead. This approach is highly scalable and fault-tolerant - even if some nodes fail, the gossip continues to spread. It's decentralized, so there's no single point of failure, and it naturally handles network partitions."

### Question 403: How would you achieve quorum-based consensus?

**Answer:**
*   **Quorum:** `(N/2) + 1`.
*   **Write:** Must succeed on W nodes.
*   **Read:** Must read from R nodes.
*   **Condition:** `W + R > N` guarantees intersection (Pigeonhole Principle).
*   **Example:** DynamoDB, Cassandra.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you achieve quorum-based consensus?

**Your Response:** "I'd use the quorum approach where we define N as total nodes, W as write quorum, and R as read quorum. The key formula is W + R > N, which guarantees that reads and writes intersect at least one common node.

For example, in a 5-node cluster, quorum is 3. If we write to 3 nodes and read from 3 nodes, we're guaranteed to hit at least one node that has the latest data. This is the pigeonhole principle in action. Systems like DynamoDB and Cassandra use this approach to balance consistency and availability. By adjusting W and R values, we can tune the system - for strong consistency, we'd use W=3, R=3, while for higher availability we might use W=2, R=2."

### Question 404: Explain vector clocks and how you’d use them.

**Answer:**
(See Q103). Captures causality.
*   `[A:1, B:0]` happens before `[A:1, B:1]`.
*   `[A:1, B:0]` and `[A:0, B:1]` are concurrent (conflict).

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain vector clocks and how you’d use them.

**Your Response:** "Vector clocks capture causality in distributed systems. Each process maintains a counter, and the vector shows what each process knows. For example, if we have [A:1, B:0], we know this happened before [A:1, B:1] because B's counter increased.

But [A:1, B:0] and [A:0, B:1] are concurrent - they happened independently, so there's a conflict. I'd use vector clocks in distributed databases to detect conflicts during eventual consistency. When two nodes update the same data concurrently, we can compare vector clocks to see if there's a conflict and trigger resolution logic like last-write-wins or manual merge. This helps maintain data integrity without requiring strong consistency across all nodes."

### Question 405: How to detect split-brain scenarios?

**Answer:**
(See Q106).
*   **Quorum Check:** If a partition can't see the majority (e.g., 3 out of 5), it steps down. "I am in the minority, I cannot be leader."

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to detect split-brain scenarios?

**Your Response:** "I'd implement quorum checks to prevent split-brain. In a 5-node cluster, if a partition can only see 2 nodes instead of the required majority of 3, it automatically steps down from leadership.

The logic is simple: if I'm in the minority, I cannot be the leader. This prevents two different leaders from making conflicting decisions in different partitions. The minority partition becomes read-only or stops serving requests until it can re-establish communication with the majority. This approach ensures that even during network partitions, there's at most one leader making decisions, maintaining consistency across the system. It's a fundamental pattern in distributed systems like ZooKeeper and etcd."

### Question 406: How to handle write skew in distributed databases?

**Answer:**
*   **Write Skew:** An anomaly where two transactions read overlapping data, make disjoint writes, but violate constraints when combined. (e.g., On-call roster: "At least one person must be on call").
*   **Fix:**
    *   **Materialize Conflict:** Create a dummy row used as a lock.
    *   **SSI (Serializable Snapshot Isolation):** DB detects conflict at commit time.
    *   **For Update:** `SELECT * ... FOR UPDATE`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle write skew in distributed databases?

**Your Response:** "Write skew is a subtle anomaly where two transactions read overlapping data, make different changes, and together violate a constraint. For example, in an on-call roster where at least one person must be on call, two doctors might both see the other is on call and both go off duty.

To fix this, I'd use several approaches. Materialize conflict by creating a dummy row that both transactions must lock. Or use Serializable Snapshot Isolation where the database detects the conflict at commit time. The simplest is using SELECT FOR UPDATE to lock the rows we read. This ensures that if two transactions try to create the same conflict scenario, one will block until the other commits, preventing the constraint violation."

### Question 407: Design a system using Raft consensus algorithm.

**Answer:**
*   **Roles:** Leader, Follower, Candidate.
*   **Log Replication:**
    1.  Client -> Leader: writes `cmd`.
    2.  Leader -> Followers: `AppendEntries(cmd)`.
    3.  Followers -> Leader: `ACK`.
    4.  If Majority ACKs: Leader commits -> Client Success.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design a system using Raft consensus algorithm.

**Your Response:** "I'd implement Raft with three roles: Leader, Follower, and Candidate. The Leader handles all client requests and coordinates replication. When a client writes, the Leader appends the command to its log and sends AppendEntries to all Followers.

Followers respond with ACKs, and once the Leader receives acknowledgments from a majority, it commits the entry and returns success to the client. If the Leader fails, Followers timeout and become Candidates, starting a new election. This ensures strong consistency - all committed entries are replicated on a majority of nodes before being acknowledged to clients. It's simpler to understand than Paxos while providing the same guarantees."

### Question 408: How do you manage schema evolution in distributed systems?

**Answer:**
*   **Format:** Avro / Protobuf (Schema Registry).
*   **Rules:**
    *   **Forward Compatible:** Old code can read New data (Reader ignores unknown fields).
    *   **Backward Compatible:** New code can read Old data (Reader handles missing fields with defaults).
*   **Registry:** Kafka Schema Registry enforces compatibility checks on Produce.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage schema evolution in distributed systems?

**Your Response:** "I'd use structured formats like Avro or Protobuf with a Schema Registry. The key is maintaining both forward and backward compatibility. Forward compatibility means old code can read new data - readers ignore unknown fields. Backward compatibility means new code can read old data - readers handle missing fields with defaults.

I'd use Kafka Schema Registry to enforce these compatibility checks when producers send data. This prevents breaking changes from reaching consumers. For example, adding a new optional field is safe, but removing or changing a field type requires careful coordination. The registry acts as a contract enforcement layer, ensuring that schema changes don't break existing services in production."

### Question 409: How would you optimize for low tail latency?

**Answer:**
(Tail Latency: The slowest 1% of requests).
*   **Hedged Requests:** Send request to 2 replicas. Use the first response. Cancel the other.
*   **Jitter:** Retry quickly if P99 threshold crossed.
*   **Resource Isolation:** Ensure GC pauses or noisy neighbors don't starve the thread.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you optimize for low tail latency?

**Your Response:** "Tail latency refers to the slowest 1% of requests, which disproportionately impact user experience. I'd use several techniques. Hedged requests - send the same request to multiple replicas and use the first response, canceling the others. This helps when one replica is slow.

I'd also add jitter to retries - if we cross the P99 threshold, retry immediately with a small random delay to avoid thundering herd problems. Most importantly, resource isolation - ensure garbage collection pauses or noisy neighbors don't starve critical request threads. This might mean using separate thread pools, dedicated cores, or even separate processes for latency-sensitive operations. The goal is to prevent a few slow requests from affecting the overall user experience."

### Question 410: Explain read-repair mechanism in eventual consistency.

**Answer:**
*   **Scenario:** Client reads from 3 nodes (Quorum read).
    *   Node A: `v2`
    *   Node B: `v2`
    *   Node C: `v1` (Stale)
*   **Repair:** Coordinator compares versions. Returns `v2` to client. Asynchronously sends `v2` to Node C to fix it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain read-repair mechanism in eventual consistency.

**Your Response:** "Read-repair is how eventually consistent systems heal themselves. When a client reads from multiple nodes in a quorum, the coordinator might see different versions - say Node A and B have v2, but Node C still has the stale v1.

The coordinator returns the latest version v2 to the client, but importantly, it also asynchronously sends v2 to Node C to bring it up to date. This happens in the background without blocking the client. Over time, read-repair operations ensure all nodes converge to the same data state. It's a key mechanism that makes eventual consistency practical - the system automatically fixes inconsistencies during normal read operations, reducing the need for separate anti-entropy processes."

---

## 🔸 Real-Time & Event-Driven Systems (Questions 411-420)

### Question 411: Design a real-time gaming leaderboard system.

**Answer:**
(See Q191). Redis Sorted Set.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design a real-time gaming leaderboard system.

**Your Response:** "I'd use Redis Sorted Sets for real-time leaderboards. Each player's score would be stored as a member in the sorted set with their score as the ranking value. Redis provides O(log N) performance for updates and O(log N + M) for retrieving top M players.

For real-time updates, when a player completes a game, we'd update their score using ZADD. The leaderboard can be retrieved with ZRANGE for top players. Redis handles the sorting automatically, so we don't need to maintain separate ranking logic. This approach scales to millions of players with sub-millisecond latency, perfect for gaming leaderboards where players expect to see their rank updated instantly after each game."

### Question 412: How would you build a high-frequency trading system?

**Answer:**
*   **Latency:** Minimize network hops. Colocation (Server in same building as Exchange).
*   **OS:** Kernel bypass (DPDK/Solarflare) to read packets directly from NIC to Userspace (Skip TCP/IP stack overhead).
*   **Language:** C++ / Rust / Java (Zero GC).
*   **Logic:** Simple algorithms (Arbitrage). No DB calls. Everything in RAM.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you build a high-frequency trading system?

**Your Response:** "For HFT, every microsecond counts. I'd colocate servers in the same building as exchanges to minimize network latency. I'd use kernel bypass technologies like DPDK to read packets directly from the NIC to userspace, skipping the TCP/IP stack overhead.

I'd choose zero-GC languages like C++ or Rust, or carefully tuned Java with minimal garbage collection. The trading logic would be simple arbitrage algorithms with everything kept in RAM - no database calls during trading. The focus is on raw speed - we need to make trading decisions in microseconds. This architecture prioritizes latency over everything else, as even a millisecond delay can mean losing millions in trading opportunities."

### Question 413: How would you ensure ordering in an event stream?

**Answer:**
(See Q122). Partition Key.
*   Kafka guarantees order *within a partition*. All events for `Order_123` must go to Partition 5.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you ensure ordering in an event stream?

**Your Response:** "I'd use Kafka partition keys to guarantee ordering. Kafka maintains strict ordering within each partition, but not across partitions. So I'd ensure all events for the same entity - like Order_123 - go to the same partition every time.

The partition key would be the order ID, which Kafka hashes to determine the partition number. This means all events for Order_123 (created, updated, paid, shipped) will be processed in the exact order they occurred. If we need to scale beyond one partition's capacity, we might shard by customer ID instead, but the principle remains the same - related events must share the same partition key to maintain their ordering."

### Question 414: What’s the difference between event sourcing and event streaming?

**Answer:**
*   **Event Sourcing (Storage Pattern):** DB structure. Store `[Deposited, Withdrew]` instead of `Balance=10`. Replayable source of truth.
*   **Event Streaming (Transport Pattern):** Moving data via Kafka. You can stream events without using event sourcing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between event sourcing and event streaming?

**Your Response:** "Event sourcing is a storage pattern where we store the sequence of events that led to the current state, rather than storing the state itself. Instead of storing Balance=10, we store Deposited and Withdrew events. This gives us a complete audit trail and the ability to replay events to reconstruct state.

Event streaming is about moving events between systems using technologies like Kafka. You can have event streaming without event sourcing - for example, streaming sensor data to analytics systems. Conversely, you can use event sourcing without streaming by storing events in a traditional database. They're complementary patterns - event sourcing solves the persistence problem, while event streaming solves the distribution problem."

### Question 415: Design a telemetry system for autonomous vehicles.

**Answer:**
*   **Volume:** TBs of data.
*   **Strategy:**
    *   **Critical:** Stream immediately (Crash detected) via 5G/MQTT.
    *   **Non-Critical:** Buffer on local SSD. Upload bulk via Wi-Fi when car is charging at night.
*   **Compression:** High efficiency (Delta encodings).

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design a telemetry system for autonomous vehicles.

**Your Response:** "Autonomous vehicles generate terabytes of data, so I'd use a tiered strategy. Critical events like crash detection would stream immediately via 5G or MQTT for real-time response. Non-critical telemetry like sensor readings would buffer on local SSD and upload in bulk when the car connects to Wi-Fi, typically while charging overnight.

I'd implement high-efficiency compression using delta encodings to reduce bandwidth usage. The system needs to handle both real-time critical data and massive batch uploads. This approach balances the need for immediate response to critical events with the reality that streaming terabytes of data continuously would be prohibitively expensive. The local buffering ensures we don't lose data even if connectivity is temporarily lost."

### Question 416: How would you handle late-arriving events?

**Answer:**
(Stream Processing problem).
*   **Watermarks:** A heuristic asserting "No events older than T will arrive now".
*   **Windowing:** "Hourly Window".
*   **Late Data:**
    *   *Allowed:* Update the already-computed window (if within 1 hour).
    *   *Dropped:* Determine cutoff (e.g., > 1 day late).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you handle late-arriving events?

**Your Response:** "Late-arriving events are common in stream processing. I'd use watermarks - a heuristic that asserts no events older than a certain timestamp will arrive. For example, after processing events up to 2 PM, we set a watermark saying we won't accept events older than 1:55 PM.

For time windows like hourly aggregations, I'd allow late events within a grace period - say one hour - to update already-computed windows. Events arriving later than the cutoff would be dropped. The key is balancing completeness with timeliness - we want accurate results but can't wait forever for late data. This approach handles network delays, out-of-order delivery, and system restarts while keeping the stream processing pipeline moving forward."

### Question 417: How do you design time-window-based analytics?

**Answer:**
*   **Tumbling Window:** Non-overlapping. `[12:00-12:05]`, `[12:05-12:10]`.
*   **Sliding Window:** Overlapping. "Last 5 mins, updated every 1 min".
*   **Session Window:** "While user is active + 30m timeout".

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you design time-window-based analytics?

**Your Response:** "I'd use different window types based on the use case. Tumbling windows are non-overlapping time buckets like 5-minute intervals - perfect for hourly sales reports. Sliding windows overlap and update frequently, like 'last 5 minutes updated every minute' - great for real-time dashboards.

Session windows are dynamic, based on user activity with a timeout - like 'while the user is active plus 30 minutes after their last action'. The choice depends on whether we need fixed time periods, continuous real-time views, or activity-based grouping. Each window type serves different analytical needs - tumbling for periodic reporting, sliding for real-time monitoring, and session for user behavior analysis."

### Question 418: How to handle out-of-order events?

**Answer:**
*   **Buffer:** Stream processor holds events in memory buffer.
*   **Sort:** Re-order based on `EventTime` (not `ProcessingTime`).
*   **Watermark:** Trigger processing when Watermark passes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle out-of-order events?

**Your Response:** "I'd buffer events in memory and sort them by EventTime rather than ProcessingTime. The stream processor holds events until it's confident no earlier events will arrive, using watermarks as the trigger.

For example, if we're processing events up to 2:00 PM, the watermark might be 1:55 PM, meaning we assume no events older than 1:55 PM will arrive. Events in the buffer get sorted by their actual event time, not when they were processed. When the watermark passes, we process the sorted events. This ensures we compute accurate time-based aggregations even when events arrive out of order due to network delays or system issues."

### Question 419: Build a real-time bidding system for ads.

**Answer:**
(See Q389).

### How to Explain in Interview (Spoken style format)
**Interviewer:** Build a real-time bidding system for ads.

**Your Response:** "I'd build it around Real-Time Bidding with sub-100ms response times. The key is an inverted index that maps targeting criteria like age, location, and interests to eligible ads.

When a bid request comes in, I'd filter eligible ads based on targeting, calculate the eCPM for each by multiplying their bid by predicted click-through rate, and select the winner. For budget management, I'd implement pacing algorithms to ensure advertisers don't exhaust their budget in the first hour. The system needs to be extremely fast while handling complex targeting logic and massive concurrent requests from multiple ad exchanges simultaneously."

### Question 420: How do you debounce vs throttle events at scale?

**Answer:**
*   **Throttle (Rate Limit):** "Max 1 request per second". (Drop extras).
*   **Debounce (Delay):** "Wait for 1s of silence". If user types 'A', wait. Types 'AB', reset timer. Types 'ABC'... stop. Wait 1s -> Send 'ABC'.
*   **Implementation:** Client-side via JS. Server-side via Redis keys expiration.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you debounce vs throttle events at scale?

**Your Response:** "Throttling limits the rate of events - like 'maximum 1 request per second' - dropping any extras that exceed the limit. This is good for API rate limiting.

Debouncing waits for a period of silence before processing - like waiting 1 second after the user stops typing. If they type 'A', we wait; they type 'AB', we reset the timer; when they finally stop and 1 second passes, we send 'ABC'. This is perfect for search suggestions. I'd implement throttling server-side using Redis for rate limiting, and debouncing client-side with JavaScript timers. Both help manage system load, but for different use cases."

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
