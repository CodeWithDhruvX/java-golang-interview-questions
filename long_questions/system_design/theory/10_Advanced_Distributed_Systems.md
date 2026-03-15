# 🔴 Advanced Distributed Systems — Questions 91–100

> **Level:** 🔴 Senior – Principal
> **Asked at:** Google, Amazon, Meta, LinkedIn, Uber — Staff/Principal Engineer and Architect roles

---

### 91. What is CAP theorem?
"The CAP theorem states that a distributed data store can only guarantee **two of three** properties simultaneously: **Consistency** (C), **Availability** (A), and **Partition Tolerance** (P).

**Consistency:** Every read returns the most recent write or an error — no stale data served. **Availability:** Every request receives a response (no error), but it might be stale. **Partition Tolerance:** The system continues operating even if some messages between nodes are lost or delayed.

In reality, network partitions are unavoidable in distributed systems — so P is not optional. The real trade-off is **CP vs AP** during a partition: do you return an error (preferring consistency over availability) or serve stale data (preferring availability over consistency)?

Payments must be CP. Social media feeds can be AP."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (DynamoDB's eventual consistency), Google (Spanner's strong consistency), any distributed DB design discussion

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is CAP theorem?
**Your Response:** The CAP theorem states that a distributed data store can only guarantee two of three properties simultaneously: Consistency, Availability, and Partition Tolerance. Consistency means every read returns the most recent write or an error - no stale data. Availability means every request receives a response, but it might be stale. Partition Tolerance means the system continues operating even if some messages between nodes are lost or delayed. In reality, network partitions are unavoidable, so partition tolerance is not optional. The real trade-off is CP vs AP during a partition - do you return an error to prefer consistency over availability, or serve stale data to prefer availability over consistency? Payments systems must be CP, while social media feeds can be AP.

#### Indepth
CAP in practice:

| System | Choice | Why |
|---|---|---|
| PostgreSQL (master) | CP — errors on partition | ACID single node, consistency prioritized |
| MySQL with sync replication | CP | Waits for replica ack |
| Cassandra | AP | Tunable consistency; default is eventual |
| DynamoDB | AP (eventual) or CP (strong reads) | User configurable per request |
| ZooKeeper / etcd | CP | Leader election and coordination must be consistent |
| Redis (cluster) | AP | Async replication → potential stale reads |
| Google Spanner | CP (borderline) | TrueTime + 2-phase commit buys near-synchronous global consistency |

**PACELC extension:** Better model than CAP. When system is **P**artitioned → **A** vs **C** trade-off (CAP). When **E**lse (no partition) → **L**atency vs **C**onsistency trade-off. Spanner and CockroachDB sacrifice latency for global consistency even without partitions. Cassandra sacrifices consistency for low latency.

---

### 92. What is the ACID property?
"ACID is the set of properties that guarantee **database transactions are processed reliably** even in the face of errors and failures. Every serious financial system depends on these.

**Atomicity:** A transaction is all-or-nothing. Transfer $100 from A to B: either both debit A and credit B happen, or neither does. No half-transactions. **Consistency:** A transaction brings the DB from one valid state to another — all constraints, foreign keys, and business rules are satisfied. **Isolation:** Concurrent transactions appear to run sequentially — one transaction doesn't see another's uncommitted changes. **Durability:** Once a transaction commits, it's permanent — even if the server crashes immediately after, the data survives (persisted to disk/WAL)."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Razorpay, Paytm, PhonePe (financial transactions), Amazon, any company with DB design questions

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the ACID property?
**Your Response:** ACID is the set of properties that guarantee database transactions are processed reliably even with errors and failures. Every serious financial system depends on these. Atomicity means a transaction is all-or-nothing - like transferring $100 from A to B, either both the debit and credit happen, or neither does. Consistency means a transaction brings the database from one valid state to another, satisfying all constraints and business rules. Isolation means concurrent transactions appear to run sequentially - one transaction doesn't see another's uncommitted changes. Durability means once a transaction commits, it's permanent even if the server crashes immediately after.

#### Indepth
ACID implementation details:
- **Atomicity:** Implemented via **WAL (Write-Ahead Log)**. Every operation is first written to an append-only WAL file. On commit, WAL is marked committed. On crash before commit, WAL shows incomplete entries — DB rolls them back on recovery.
- **Consistency:** Enforced by **constraint checking** at commit time: NOT NULL, UNIQUE, FOREIGN KEY, CHECK constraints. Application-level consistency (business rules) is the developer's responsibility.
- **Isolation:** Implemented via **MVCC (Multi-Version Concurrency Control)** in PostgreSQL, MySQL InnoDB. Each transaction sees a consistent snapshot of the DB. Reading doesn't block writing. Standard isolation levels: Read Uncommitted → Read Committed → Repeatable Read → Serializable.
- **Durability:** `fsync()` call flushes WAL to durable storage before commit acknowledgment. Cloud databases (Aurora, Spanner) write to distributed storage (6 replicas) for durability without single-disk dependency.

**Serializable Snapshot Isolation (SSI):** PostgreSQL's Serializable level uses SSI — detects serialization anomalies without locks. Google Spanner uses external consistency (beyond serializable) — all distributed transactions globally serializable by timestamp (TrueTime API).

---

### 93. What is BASE?
"BASE (Basically Available, Soft state, Eventually consistent) is the alternative to ACID in distributed NoSQL systems. It reflects the trade-offs made to achieve high availability and horizontal scalability.

**Basically Available:** The system responds to every request, but the response may be stale data or a partial answer — it's always available, just not always consistent. **Soft state:** The system state may change over time even without input (due to eventual consistency propagation). **Eventually consistent:** Given no new updates, all nodes will eventually converge to the same state — but there's a window where they disagree.

Cassandra, DynamoDB, CouchDB — these are BASE systems. They sacrifice consistency for availability and partition tolerance. The critical insight: many applications *don't need* strong consistency. A social media like count doesn't need to be exact to the millisecond."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (DynamoDB discussions), system design interviews involving NoSQL decisions — Flipkart, Swiggy, Hotstar

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is BASE?
**Your Response:** BASE stands for Basically Available, Soft state, and Eventually consistent - it's the alternative to ACID in distributed NoSQL systems. Basically Available means the system responds to every request, but the response may be stale data or a partial answer - it's always available, just not always consistent. Soft state means the system state may change over time even without input due to eventual consistency propagation. Eventually consistent means given no new updates, all nodes will eventually converge to the same state, but there's a window where they disagree. Systems like Cassandra, DynamoDB, and CouchDB are BASE systems that sacrifice consistency for availability and partition tolerance.

#### Indepth
ACID vs BASE design decision framework:
- Use **ACID** when: financial transactions, inventory management (can't oversell), election/consensus (only one leader), referential integrity is required.
- Use **BASE** when: social feeds, product recommendations, analytics, user activity logs, anything where slightly stale data has low business cost.

**Real-world ACID + BASE hybrid:** Most large systems use both. Payment tables → PostgreSQL (ACID). User activity feed → Cassandra (BASE). Product catalog → MySQL (ACID for inventory counts, eventual consistency for search rankings). The trick is knowing which business requirement needs which guarantee.

**Tunable consistency in DynamoDB and Cassandra:**
- DynamoDB: default eventual consistency reads. `ConsistentRead=true` for strong consistency reads (double the read capacity units). Choose per-operation.
- Cassandra: `ConsistencyLevel` per query: `ONE` (fastest, least consistent) → `QUORUM` (majority of nodes must respond) → `ALL` (all nodes must respond, highest consistency, lowest availability). Formula: if `W + R > N` → strong consistency guaranteed.

---

### 94. What is consistent hashing?
"Consistent hashing is an algorithm for **distributing data (or requests) across a dynamic set of nodes** such that when nodes are added or removed, only a minimal fraction of keys need to be remapped.

The classic problem: you have N cache servers. You hash every key to `key % N` to find which server holds it. A server goes down → N becomes N-1 → `key % N-1` redistributes almost every key → cache miss storm — your DB gets hammered.

Consistent hashing solves this: both servers AND keys are mapped to positions on a circular hash ring (0 to 2^32). Each key is served by the first server clockwise on the ring. When a server is added or removed, only the keys that fall in its range are remapped — roughly K/N keys, where K is total keys and N is total servers. Adding a server only moves a fraction of keys."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (DynamoDB, Dynamo paper), Discord, Cassandra design discussions, Lyft, Uber — distributed storage questions

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is consistent hashing?
**Your Response:** Consistent hashing is an algorithm for distributing data across a dynamic set of nodes such that when nodes are added or removed, only a minimal fraction of keys need to be remapped. The classic problem is you have N cache servers and hash every key to key % N to find which server holds it. When a server goes down, N becomes N-1 and almost every key gets redistributed, causing a cache miss storm. Consistent hashing solves this by mapping both servers and keys to positions on a circular hash ring. Each key is served by the first server clockwise on the ring. When a server is added or removed, only the keys in its range are remapped - roughly K/N keys, where K is total keys and N is total servers.

#### Indepth
Consistent hashing in practice:
- **Virtual nodes (vnodes):** A real server doesn't occupy just one point on the ring — it occupies many virtual positions (e.g., 150 virtual nodes per server). This distributes load more evenly (prevents hotspots when ring positions cluster). Cassandra, DynamoDB use vnodes.
- **Load imbalance without vnodes:** A server added to the ring might get a disproportionately large segment. With 150 vnodes spread randomly, each server's total responsibility is statistically uniform.
- **Replication with consistent hashing:** In Cassandra, data is replicated to the next N servers clockwise on the ring (where N = replication factor). With `RF=3`, data is stored on 3 consecutive servers. If one dies, the others still have the data.

```
Ring: S1(20) → S2(50) → S3(80) → S4(120) → [back to S1]
Key "user:123" → hash → 65 → goes to S3 (first server clockwise)
If S3 goes down:
  Key 65 now goes to S4 (next clockwise)
  Only keys in S3's segment (50-80) are remapped to S4
  All other keys unaffected
```

Used by: Cassandra, DynamoDB, Akamai (CDN routing), Discord (chat server assignment), Redis Cluster (hash slots, not classic consistent hashing but the same idea).

---

### 95. What is a distributed lock?
"A distributed lock is a **cross-process mutual exclusion mechanism** that ensures only one process across multiple machines executes a critical section simultaneously.

The classic use case: a distributed cron job. If I run 5 instances of a job scheduler, I only want one instance to run the job at any given time. A distributed lock stored in Redis ensures mutual exclusion.

Redis-based lock: `SET lock-key unique-value NX EX 30` — set if not exists with 30-second expiry. NX ensures atomicity. EX ensures the lock is auto-released if the process crashes (preventing deadlock). Release: Lua script `if redis.call('GET', key) == myValue then redis.call('DEL', key) end` — release only if I own the lock."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, PhonePe (payment deduplication), Flipkart (inventory reservation), Uber (trip assignment)

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a distributed lock?
**Your Response:** A distributed lock is a cross-process mutual exclusion mechanism that ensures only one process across multiple machines executes a critical section simultaneously. The classic use case is a distributed cron job - if I run 5 instances of a job scheduler, I only want one instance to run the job at any given time. A distributed lock stored in Redis ensures mutual exclusion. The Redis-based approach uses SET lock-key unique-value NX EX 30 - set if not exists with a 30-second expiry. NX ensures atomicity and EX ensures the lock is auto-released if the process crashes to prevent deadlock. To release, I use a Lua script that only deletes the key if I still own the lock.

#### Indepth
Distributed lock approaches and caveats:
1. **Redis SET NX EX (single node):** Simple, fast. Problem: if Redis master fails before replica syncs the lock → two processes both think they have the lock (split-brain). Acceptable for non-critical locks.
2. **Redlock (Redis multi-node):** Write lock to majority of N Redis nodes independently. Lock valid only if granted by majority (N/2+1) within `validity_time = TTL - elapsed`. Controversial (Martin Kleppmann's critique: clock skew can still cause safety violations). Used cautiously.
3. **ZooKeeper / etcd ephemeral nodes:** Create an ephemeral node → if holder dies, session expires → node deleted → lock released. Strongly consistent (ZooKeeper uses Zab consensus; etcd uses Raft). Safer than Redis for critical locks. Higher latency.
4. **Database-based lock:** `UPDATE locks SET holder=myId WHERE resource=X AND holder=NULL`. DB constraints guarantee atomicity. Simplest but DB becomes a bottleneck for high-frequency locking.

**Martin Kleppmann vs antirez debate (2016):** Kleppmann argued Redlock is unsafe due to process pauses and clock drift. antirez (Redis creator) defended Redlock's practical safety. The community consensus: Redlock is fine for lock-based coordination where occasional incorrect behavior is acceptable (e.g., cache invalidation). For correctness-critical systems (financial deduplication), use ZooKeeper or etcd.

---

### 96. What is leader election?
"Leader election is the process by which **a distributed set of nodes chooses one node to be the designated leader** for coordination purposes — the node responsible for making authoritative decisions.

Examples: Kafka uses leader election per partition topic (Zookeeper → KRaft). Database systems use leader election for primary selection (Patroni uses etcd). Kubernetes uses leader election for controller managers (only one controller manager should run reconciliation at a time).

The algorithm used is typically a **consensus protocol** — Raft or Paxos. These protocols guarantee that even if the network splits, only one leader will ever be elected at a time (safety), and a leader will eventually be elected (liveness)."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Kafka internals discussions, Kubernetes/distributed systems infrastructure roles at Amazon, Google

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is leader election?
**Your Response:** Leader election is the process by which a distributed set of nodes chooses one node to be the designated leader for coordination purposes - the node responsible for making authoritative decisions. Examples include Kafka using leader election per partition topic, database systems using it for primary selection, and Kubernetes using it for controller managers where only one should run reconciliation at a time. The algorithm used is typically a consensus protocol like Raft or Paxos, which guarantee that even if the network splits, only one leader will ever be elected at a time for safety, and a leader will eventually be elected for liveness.

#### Indepth
Raft leader election (simplified):
1. All nodes start as **Followers**. If no heartbeat received from leader within election timeout (150-300ms random), a follower becomes a **Candidate**.
2. Candidate increments term counter, votes for itself, sends `RequestVote` RPCs to all other nodes.
3. Nodes grant vote if: (a) candidate's term > their term, and (b) candidate's log is at least as up-to-date as theirs. Each node votes for at most one candidate per term.
4. If candidate receives majority votes → becomes **Leader**. Sends heartbeats to all followers to prevent new elections.
5. If no candidate receives majority (split vote) → all candidates timeout (random delay prevents persistent split), increment term, retry.

**etcd's role in Kubernetes:** etcd stores all cluster state. Kubernetes API server is stateless — all state in etcd. etcd uses Raft internally, so it's itself highly available. API server leader election: multiple API server instances use an etcd lock (via `coordination.k8s.io/Lease` resource) to elect one leader for certain coordination tasks.

---

### 97. What is a distributed transaction?
"A distributed transaction is a transaction that **spans multiple databases or services** — where all participants must commit or all must roll back to maintain ACID semantics.

The classic protocol for this is **2-Phase Commit (2PC)**: Phase 1 (Prepare) — coordinator asks all participants to prepare and vote. If all vote YES, move to phase 2. Phase 2 (Commit) — coordinator sends commit to all participants.

The problem with 2PC: if the coordinator crashes after participants commit local but before sending the final commit message, participants are stuck waiting (blocking). 2PC is blocking and has low availability. For this reason, modern distributed systems often prefer **Sagas** (as discussed in Q48) — a sequence of local transactions with compensating actions, accepting eventual consistency."

#### 🏢 Company Context
**Level:** 🔴 Senior / Principal | **Asked at:** Amazon (Dynamo paper explicitly avoided 2PC), Flipkart, Razorpay, Google (Spanner's 2PC implementation)

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a distributed transaction?
**Your Response:** A distributed transaction is a transaction that spans multiple databases or services, where all participants must commit or all must roll back to maintain ACID semantics. The classic protocol is 2-Phase Commit: Phase 1 is Prepare where the coordinator asks all participants to prepare and vote. If all vote yes, we move to Phase 2 which is Commit where the coordinator sends commit to all participants. The problem with 2PC is that if the coordinator crashes after participants commit locally but before sending the final commit message, participants get stuck waiting - it's blocking and has low availability. For this reason, modern systems often prefer Sagas - a sequence of local transactions with compensating actions, accepting eventual consistency.

#### Indepth
Distributed transaction protocols:
- **2PC (Two-Phase Commit):** Coordinator-based, blocking, synchronous. Used in: distributed SQL DBs (MySQL NDB Cluster, Postgres with foreign data wrappers). Problem: coordinator is a SPOF; blocking on coordinator failure.
- **3PC (Three-Phase Commit):** Adds a pre-commit phase to prevent blocking on coordinator failure. Never widely adopted in practice — too complex, still has edge cases.
- **Saga Pattern:** Sequence of local transactions + compensating transactions. No distributed lock held across services. Accepts eventual consistency. Preferred in microservices. (Covered in Q48)
- **Google Spanner's Approach:** Spanner *does* support distributed transactions using 2PC + TrueTime. TrueTime gives bounded clock uncertainty. Spanner waits for the uncertainty window (5-10ms) at commit time to guarantee external consistency. This is how Google achieves global strong consistency.
- **Amazon Aurora Global Transactions:** Aurora supports XA transactions (distributed transaction protocol). Used for cross-region transactions in Aurora Global Database.

---

### 98. What is the Paxos algorithm?
"Paxos is a **consensus algorithm** that allows a distributed set of nodes to agree on a single value, even if some nodes fail or messages are delayed. It's the theoretical foundation for distributed consensus.

The problem it solves: imagine 5 database nodes. A user writes value `X`. How do we ensure all 5 nodes agree on the same value, even if Node 3 is temporarily partitioned? Paxos provides a mathematically provable algorithm for this.

Paxos has two roles: **Proposers** (suggest values to accept) and **Acceptors** (vote on values). Two phases: Phase 1 (Prepare/Promise) — proposer asks for a promise. Phase 2 (Accept/Learn) — proposer proposes a value; if majority accepts, consensus is reached.

In practice, Raft is now preferred over Paxos — it was designed to be more understandable and equally correct."

#### 🏢 Company Context
**Level:** 🔴 Senior / Principal | **Asked at:** Google (pioneered Paxos use with Chubby), distributed systems theory questions

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the Paxos algorithm?
**Your Response:** Paxos is a consensus algorithm that allows a distributed set of nodes to agree on a single value, even if some nodes fail or messages are delayed. It's the theoretical foundation for distributed consensus. Imagine 5 database nodes and a user writes value X - how do we ensure all 5 nodes agree on the same value even if Node 3 is temporarily partitioned? Paxos provides a mathematically provable algorithm for this. Paxos has two roles: Proposers who suggest values and Acceptors who vote on values. It has two phases: Phase 1 is Prepare/Promise where the proposer asks for a promise, and Phase 2 is Accept/Learn where the proposer proposes a value and if majority accepts, consensus is reached. In practice, Raft is now preferred over Paxos as it was designed to be more understandable.

#### Indepth
Paxos vs Raft:
| Feature | Paxos | Raft |
|---|---|---|
| Author | Leslie Lamport, 1989 | Ongaro & Ousterhout, 2014 |
| Goal | Consensus (theoretical) | Understandable consensus (practical) |
| Leader | Flexible (multiple proposers) | Single strong leader per term |
| Log Replication | Not prescribed | Clear log structure |
| Understandability | Hard (Lamport himself said so) | Designed for clarity |
| Adoption | Google Chubby, MySQL Group Replication | etcd, CockroachDB, TiKV, Consul |

Multi-Paxos (used in practice): Basic Paxos agrees on one value. Multi-Paxos extends it to agree on a log of values (a sequence of commands). Leader elected once, then proposes values in phase 2 only (skip phase 1 for subsequent values). More efficient for continuous log replication.

**Real-world implementations:**
- **Chubby (Google):** Distributed lock service using Paxos. Used for Bigtable master election, GFS chunk server leases.
- **Zab (Zookeeper):** Zookeeper's consensus protocol, similar to Multi-Paxos.
- **Raft:** etcd (used by Kubernetes), CockroachDB, TiKV (TiDB), Consul, RethinkDB.

---

### 99. What is the split-brain problem?
"Split-brain occurs when **a distributed system partitions into two separate groups, each believing the other is down** and both attempting to operate as the primary — leading to conflicting writes and data corruption.

In a primary-secondary DB setup: a network partition between primary and secondary makes the secondary think the primary is dead → secondary promotes itself to primary. Now both think they're primary and accept writes → two diverging versions of truth. When the partition heals, you have conflicting data with no automated way to merge.

Split-brain is one of the most dangerous failure modes in distributed systems. Prevention is more reliable than detection."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Database infrastructure discussions — Amazon RDS HA, PostgreSQL Patroni, Redis Sentinel — any HA DB design

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the split-brain problem?
**Your Response:** Split-brain occurs when a distributed system partitions into two separate groups, each believing the other is down and both attempting to operate as the primary - leading to conflicting writes and data corruption. In a primary-secondary database setup, a network partition makes the secondary think the primary is dead, so it promotes itself to primary. Now both think they're primary and accept writes, creating two diverging versions of truth. When the partition heals, you have conflicting data with no automated way to merge. Split-brain is one of the most dangerous failure modes in distributed systems, and prevention is more reliable than detection.

#### Indepth
Split-brain prevention strategies:
1. **Quorum (majority) requirement:** A node can only become primary if it receives confirmation from a majority (N/2+1) of nodes. A partitioned minority of nodes can never form quorum → they don't promote themselves.
   - etcd, Raft, ZooKeeper all use quorum.
   - AWS RDS Multi-AZ: uses a single AZ witness to break ties. Primary must have quorum; if primary loses contact with witness, it steps down.
2. **STONITH (Shoot The Other Node In The Head):** When a primary suspects the secondary has taken over (or vice versa), it sends a "fence" command to physically power off or reboot the other node. Brutal but effective — only one node can run.
3. **Fencing tokens:** Each time a leader is elected, it's given a monotonically increasing token. Operations must include this token. Storage backend rejects operations with outdated tokens. Even if the old leader thinks it's primary, its writes are rejected.
4. **Witness / Arbitrator node:** A third node (witness) doesn't store data but participates in quorum decisions. Breaks 2-node split-brain. AWS Aurora uses this via storage quorum across 6 nodes in 3 AZs.

---

### 100. What is a quorum in distributed systems?
"A quorum is a **minimum number of nodes that must agree before a read or write operation is considered successful**. It's the fundamental mechanism for ensuring consistency while tolerating node failures.

Given N nodes, a typical quorum configures: Write quorum `W` nodes must confirm a write. Read quorum `R` nodes must be consulted on a read. If `W + R > N`, then every read will see at least one node that has the latest write — read-write overlap is guaranteed. This is called **strong consistency via quorum**.

In DynamoDB and Cassandra, N=3 is common. `W=2, R=2, N=3` → `W+R=4 > 3` → strong consistency. `W=1, R=1` → eventually consistent. The choice is configurable per operation."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (DynamoDB strong reads), Cassandra internals, distributed systems design at Google, Uber

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a quorum in distributed systems?
**Your Response:** A quorum is a minimum number of nodes that must agree before a read or write operation is considered successful. It's the fundamental mechanism for ensuring consistency while tolerating node failures. Given N nodes, we configure a write quorum W where W nodes must confirm a write, and a read quorum R where R nodes must be consulted on a read. If W + R > N, then every read will see at least one node that has the latest write - this guarantees read-write overlap and is called strong consistency via quorum. In DynamoDB and Cassandra, N=3 is common. With W=2, R=2, N=3, we have W+R=4 which is greater than 3, so we get strong consistency. With W=1, R=1, we get eventual consistency.

#### Indepth
Quorum arithmetic and trade-offs:

| Config | W | R | N | W+R>N? | Consistency | Write Latency | Read Latency |
|---|---|---|---|---|---|---|---|
| Strong | 2 | 2 | 3 | Yes (4>3) | Strong | Higher (wait for 2) | Higher |
| Eventual | 1 | 1 | 3 | No (2<3) | Eventual | Low | Low |
| Write optimized | 1 | 3 | 3 | Yes (4>3) | Strong | Lowest | Highest |
| Read optimized | 3 | 1 | 3 | Yes (4>3) | Strong | Highest | Lowest |

**Tunable consistency in Cassandra:**
- `ConsistencyLevel.ONE`: Contact 1 node. Fastest, eventual consistency.
- `ConsistencyLevel.QUORUM`: Contact majority. Strong consistency.
- `ConsistencyLevel.ALL`: Contact all nodes. Strongest, least available (one node failure breaks all writes).
- `ConsistencyLevel.LOCAL_QUORUM`: Quorum within local DC only. Used in multi-DC to avoid cross-DC latency.

**Raft's quorum:** In a 5-node Raft cluster, a leader must receive `APPEND_ENTRIES` confirmation from 3 of 5 nodes (≥ majority) before committing a log entry. This means the cluster can tolerate 2 node failures (still has 3 remaining). No separate W+R configuration needed — it's built into the consensus protocol.
