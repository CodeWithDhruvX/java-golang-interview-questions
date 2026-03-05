# ⚖️ Distributed Systems Trade-offs — Product-Based Companies

> **Level:** 🔴 Senior / Staff
> **Asked at:** Google, Uber, Amazon, Flipkart, PhonePe, Razorpay

---

## Q1. What is consistent hashing and why is it used in distributed systems?

"Consistent hashing is a technique for distributing keys across a cluster of nodes such that **adding or removing a node only remaps a fraction of the keys** — proportional to 1/N where N is the number of nodes — rather than remapping everything.

Classic modulo hashing (`key % N`) requires remapping nearly all keys when a node joins or leaves. In a Redis cluster with 10 shards, adding an 11th shard would invalidate ~91% of cached entries — a cache stampede.

With consistent hashing: nodes and keys are mapped to positions on a circular ring (hash ring). A key is assigned to the first node clockwise from its position. When a node is added, only keys between the new node and its predecessor migrate. When a node is removed, its keys redistribute only to the successor."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart, Uber, any company running distributed caches or databases

#### Deep Dive
**The hash ring implementation:**
```
Ring positions: 0 ───────────────────────────── 2^32-1 (wraps around)

Nodes placed at: Node A = 12%, Node B = 45%, Node C = 78%

Key K hashes to 50% → assigned to Node C (next clockwise)
Key K hashes to 80% → assigned to Node A (wraps around to start)

Add Node D at 60%:
  - All keys from 45% to 60% move from Node C → Node D
  - All other keys unaffected
```

**Virtual nodes (vnodes) — solving the hotspot problem:**
Without vnodes, physical placement is uneven. Node A might cover 10% of the ring while Node B covers 40%.

Solution: Each physical node gets multiple virtual node positions on the ring (e.g., 150 vnodes per physical node). Keys are spread evenly across all physical nodes regardless of hash distribution.

```
Physical: 3 nodes
Virtual:  Node A → vnodes at [5%, 30%, 65%, 81%, ...]
          Node B → vnodes at [12%, 48%, 72%, 90%, ...]
          Node C → vnodes at [20%, 55%, 78%, 95%, ...]
→ Each node ends up serving ~33% of keys
```

**Used by:** Amazon DynamoDB, Apache Cassandra, Redis Cluster (partial), CDN edge routing.

**Trade-off with consistent hashing:** Hotspots can still occur when a single key is read-heavy (celebrity problem). Solution: replicate that key to multiple nodes and read from any replica.

---

## Q2. Explain the Raft consensus algorithm. How does leader election work?

"Raft is a consensus algorithm that allows a cluster of servers to agree on a sequence of values (log entries) even in the presence of failures. It's designed to be more understandable than Paxos while providing equivalent guarantees.

Every Raft cluster has one **leader** at a time. All writes go through the leader. The leader replicates log entries to followers. A quorum (majority) must acknowledge an entry before the leader commits it and returns success to the client.

This guarantees: **if a write is acknowledged, it will survive any single-node failure** — because a majority of nodes have it."

#### Company Context & Level
**Level:** 🔴 Staff/Principal | **Asked at:** Google, Uber, any company building distributed systems from scratch

#### Deep Dive
**Leader election — the Raft state machine:**
```
All nodes start as Followers.
  ↓ (election timeout expires — no heartbeat from leader)
Follower → Candidate
  ↓ sends RequestVote RPC to all peers
If receives votes from majority → becomes Leader
If receives AppendEntries from valid leader → reverts to Follower

Leader sends periodic heartbeat (empty AppendEntries) to reset followers' election timers.
If leader dies → followers' timers expire → new election → new leader elected in ~150-300ms.
```

**Term numbers — the "logical clock" of Raft:**
- Every election starts a new **term** (monotonically increasing integer)
- Nodes reject messages from lower-term leaders (prevents stale leaders)
- If a node sees a higher term → immediately reverts to Follower

**Log replication — the core operation:**
```
Client sends: "SET key=value"
Leader:
  1. Appends entry to local log: [term=3, index=47, cmd="SET key=value"]
  2. Sends AppendEntries to all followers in parallel
  3. Waits for majority (e.g., 3/5) acknowledgments
  4. Commits entry (advances commitIndex)
  5. Applies to state machine
  6. Returns success to client
  7. Next heartbeat notifies followers to commit
```

**Raft vs Paxos:**
| Aspect | Raft | Paxos (Multi-Paxos) |
|--------|------|---------------------|
| Understandability | Designed for clarity | Notoriously hard |
| Single leader | Yes (at any time) | Yes (after leader election) |
| Log handling | Log-centric design | Log built on top |
| Used by | etcd, CockroachDB, TiKV | Google Chubby, Zookeeper (ZAB variant) |

**etcd and Kubernetes:** Kubernetes control plane stores all cluster state (pods, services, config) in etcd, which uses Raft. This is why etcd requires an odd number of nodes (3 or 5) — you need a quorum to elect a leader.

---

## Q3. What are vector clocks and how do they handle causality in distributed systems?

"Vector clocks solve the problem of ordering events in a distributed system where there's no global clock. Each node maintains a vector of counters — one counter per node in the cluster. When an event occurs or a message is sent, the vector clock is incremented.

Two events can be compared: if every counter in clock A is ≤ the corresponding counter in clock B, then A happened-before B. If neither is ≤ the other — they are **concurrent** (neither caused the other). This is the classic conflict that distributed databases must resolve."

#### Company Context & Level
**Level:** 🔴 Staff | **Asked at:** Amazon (DynamoDB design discussion), Uber, distributed systems roles

#### Deep Dive
**Vector clock example — shopping cart conflict:**
```
Initial: {A:0, B:0, C:0}

User on Node A adds "shoes":       A's clock → {A:1, B:0, C:0}
User on Node B (network partition)
  adds "shirt":                    B's clock → {A:0, B:1, C:0}

Partition heals. A and B sync:
  A's event: {A:1, B:0, C:0}
  B's event: {A:0, B:1, C:0}
  Neither is ≤ the other → CONCURRENT CONFLICT

DynamoDB's solution: return both versions to the client.
Client (or application) must merge them → cart = ["shoes", "shirt"].
This is the "shopping cart always merges" strategy (Last-Write-Wins would lose data).
```

**The happen-before relationship (Lamport clocks → Vector clocks evolution):**
- **Lamport clocks:** Single integer. Guarantees partial ordering but can't detect concurrency.
- **Vector clocks:** N integers (one per node). Fully captures causality and concurrency.

**Conflict resolution strategies:**
1. **Last Write Wins (LWW):** Use wall-clock timestamp to pick the latest write. Risk: clock skew loses data. Used by Cassandra default.
2. **Vector clock + client merge:** Return both to client, merge intelligently. Used by DynamoDB (original version 3 paper).
3. **CRDT (Conflict-Free Replicated Data Types):** Mathematically guaranteed to always merge consistently. Used for collaborative editors (Google Docs uses a variant). No conflicts by definition.

---

## Q4. Explain different database sharding strategies. What are the trade-offs of each?

"Sharding splits a large dataset across multiple database nodes, with each node owning a partition. The sharding key determines which node a record belongs to. The choice of sharding key and strategy is one of the most consequential architecture decisions — get it wrong and you can't fix it without migrating all your data."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Flipkart, Amazon, Razorpay, any company with DB scale problems

#### Deep Dive
**Sharding strategy comparison:**

| Strategy | How it works | When to use | Problem |
|----------|-------------|-------------|---------|
| **Range-based** | Rows 1–1M → Shard 1, 1M–2M → Shard 2 | Time-series data, sequential IDs | Hotspot: recent data all hits one shard |
| **Hash-based** | `shard = hash(key) % N` | Even distribution needed | Adding shards requires full reshard |
| **Consistent hashing** | Hash ring, vnodes | Caches, key-value stores | Cross-shard queries are expensive |
| **Directory-based** | A lookup table maps keys to shards | Full control, easy migration | Lookup table is a bottleneck + SPOF |
| **Geo-based** | Indian users → India shard | Data residency requirements | Uneven if user distribution is uneven |

**Choosing the shard key — the most critical decision:**

```
BAD shard key: created_at (date)
  → All new orders land on the "today" shard
  → 99% of writes go to 1 shard (hotspot)
  → Other shards are idle

GOOD shard key: user_id or order_id (high cardinality, even distribution)
  → Writes distribute across all shards
  → Customer support queries: "show all orders for user_123" → hits exactly 1 shard
```

**The cross-shard query problem:**
```
"Show all orders placed in the last 24 hours across all shards"
→ Must query ALL shards and merge results
→ O(N shards) database calls → slow

Solutions:
1. Denormalization: maintain a global "recent orders" table in one place
2. CQRS: maintain an Elasticsearch read model with all orders indexed by date
3. Accept the limitation: design queries to include the shard key
```

**Amazon's approach in DynamoDB:**
- Partition key = shard key (required for every operation)
- Sort key = for range queries within a partition
- Forces every query to include the partition key → guaranteed single-shard lookup
- Cross-partition scans are expensive (full table scan = reads every partition)

---

## Q5. What is the difference between replication and sharding? When do you use each?

"**Replication** creates copies of the same data on multiple nodes. **Sharding** splits different data across multiple nodes. They solve different problems and are often used together.

Replication solves: **availability and read scalability**. If one replica dies, another serves traffic. Multiple replicas can serve reads in parallel.

Sharding solves: **write scalability and storage**. When a single node can't store all your data or handle all writes, you partition across nodes."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Flipkart, Amazon, Swiggy — standard scale question

#### Deep Dive
**Replication models:**
- **Single-leader (Primary-Replica):** All writes → leader, reads → any replica. Simple, but leader is write bottleneck. MySQL, PostgreSQL default.
- **Multi-leader:** Writes accepted at multiple nodes, sync asynchronously. Better write throughput across regions. Conflict resolution required (CouchDB, active-active regions).
- **Leaderless (Dynamo-style):** Writes go to N nodes simultaneously, quorum required. DynamoDB original, Cassandra.

**Quorum reads and writes (W + R > N):**
```
N = 3 replicas, W = 2 (write quorum), R = 2 (read quorum)
W + R = 4 > N = 3 → guaranteed to overlap → strong consistency

N=3, W=1, R=1 → eventual consistency (faster, but stale reads possible)
N=3, W=3, R=1 → very durable writes, fast reads
```

**Combined: sharding + replication:**
```
DynamoDB architecture:
  4 shards, each shard replicated 3 times across AZs
  → Total 12 nodes
  → Writes: go to leader of relevant shard
  → Reads: from any replica of relevant shard
  → If shard leader fails: one replica elected as new leader
  → If entire AZ fails: other AZs still have replicas
```

---

## Q6. What is the difference between strong consistency, eventual consistency, and tunable consistency?

"Strong consistency guarantees that after a write, every subsequent read — from any node — returns the new value. Eventual consistency guarantees that if no new updates are made, all replicas will eventually converge to the same value, but reads can be stale in the interim.

Tunable consistency lets you choose the level per operation. DynamoDB and Cassandra support this: you can do a strongly-consistent read (higher latency) for a payment check, and an eventually-consistent read (lower latency) for showing a product catalog."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart, PhonePe, Razorpay — fundamental distributed systems question

#### Deep Dive
**The PACELC theorem (extension of CAP):**
```
CAP: During Partition → choose Consistency or Availability
PACELC: Even without partition → choose Latency or Consistency

DynamoDB: PA/EL → On partition: Available; Otherwise: Low latency (eventual)
CockroachDB: PC/EC → On partition: Consistent; Otherwise: Consistent (strong)
```

**Practical consistency choices in real systems:**
| Use Case | Consistency Level | Why |
|----------|-----------------|-----|
| Bank balance check | Strong | Money must not show stale balance |
| Inventory count (flash sale) | Strong | Prevent overselling |
| Product catalog | Eventual | 1-second stale is acceptable |
| Social media feed | Eventual | Followers seeing post 2 seconds late is fine |
| Like/view counts | Eventual | Approximate is acceptable |
| Session token validation | Strong | Security — can't use stale session |

**MongoDB: Read Concern levels:**
```java
// Strong consistency
collection.find(query).readConcern(ReadConcern.MAJORITY)
// Eventual (fastest — read from any secondary)
collection.find(query).readConcern(ReadConcern.LOCAL)
```
