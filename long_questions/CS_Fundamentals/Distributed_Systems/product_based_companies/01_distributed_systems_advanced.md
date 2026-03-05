# 📡 Distributed Systems — Advanced Interview Questions (Product-Based Companies)

This document covers advanced distributed systems concepts for product-based company interviews (Google, Meta, Amazon, Flipkart, Swiggy, CRED). Targeted at 3–10 years of experience rounds.

---

### Q1: Explain the PACELC theorem. How is it an improvement over CAP?

**Answer:**
**PACELC** extends CAP by addressing the latency-consistency tradeoff that CAP ignores when the system is operating normally (no partition).

**CAP limitation:** CAP only discusses what happens during a **partition (P)**. It ignores the normal-operation tradeoff.

**PACELC says:**
- **If (P)artition**: choose between **A**vailability and **C**onsistency (like CAP)
- **E**lse (normal operation): choose between **L**atency and **C**onsistency

```
PACELC = {PA or PC} + {EL or EC}
```

| System | Partition Behavior | Else (Normal) |
|---|---|---|
| DynamoDB / Cassandra | PA (prefers Availability) | EL (prefers Latency) → PA/EL |
| Zookeeper / HBase | PC (prefers Consistency) | EC (prefers Consistency) → PC/EC |
| MongoDB (v3+) | PC | EC (linearizable reads from primary) → PC/EC |
| CockroachDB | PC | EC → PC/EC |

**Why it matters:**
Even without partitions, achieving consistency requires coordination (replication protocol round trips), which adds latency. PACELC makes this tradeoff explicit. Most real-world systems accept higher latency for coordination-based consistency (QUORUM reads/writes).

---

### Q2: What is vector clocks and how are they used for conflict detection in distributed systems?

**Answer:**
**Vector clocks** are a mechanism for tracking causal relationships between events in a distributed system — determining if event A happened before B, or if they're concurrent (potential conflict).

**Structure:** Each node maintains a vector of counters, one per node in the system.

**Rules:**
- When a node does a local event: increment its own counter.
- When sending a message: attach current vector clock.
- When receiving a message: take element-wise maximum of received clock and local clock, then increment own counter.

**Example with 3 nodes (A, B, C):**
```
A writes: A[1,0,0]  → sends to B
B receives: B takes max([1,0,0],[0,0,0]) → [1,0,0], increments own: [1,1,0]
C writes independently: C[0,0,1]
```

**Comparing two vector clocks:**
- **VC1 < VC2** (causally before): VC1[i] ≤ VC2[i] for all i, and strict < for at least one.
- **Concurrent** (potential conflict): Neither VC1 < VC2 nor VC2 < VC1.

**Used in:**
- **Amazon DynamoDB/Dynamo** (original paper): Detects conflicting writes on the same key. If conflict detected, both versions returned to client for application-level resolution.
- **Riak**: Siblings (concurrent writes) tracked via vector clocks.

**Alternative: Hybrid Logical Clocks (HLC):** Combines physical time (wall clock) and logical time. Used in CockroachDB for MVCC timestamps.

---

### Q3: Explain the Two-Phase Commit (2PC) protocol. What are its failure modes?

**Answer:**
**2PC** is a distributed protocol to achieve atomicity across multiple nodes — either all commit or all abort.

**Phase 1 — Prepare:**
1. Coordinator sends `PREPARE` to all Participants.
2. Each Participant locks resources, writes to WAL (ready to commit).
3. Participant responds `VOTE_COMMIT` or `VOTE_ABORT`.

**Phase 2 — Commit:**
1. If ALL voted `COMMIT` → Coordinator sends `COMMIT` to all. Participants commit and release locks.
2. If ANY voted `ABORT` → Coordinator sends `ABORT` to all. Participants rollback and release locks.

**Failure modes:**

| Failure | Effect | Recovery |
|---|---|---|
| Participant crashes before PREPARE | Coordinator doesn't get vote → aborts. Safe. | Participant aborts on recovery |
| Participant crashes after PREPARE (sent VOTE_COMMIT) | Participant holds locks, uncertain if committed | Must wait for coordinator to recover — **blocking problem** |
| Coordinator crashes after PREPARE but before COMMIT | All participants are stuck waiting — locks held indefinitely | **Blocking failure — major problem with 2PC** |
| Coordinator crashes after sending some COMMITs | Some committed, some not — inconsistent state | Coordinator recovery + redo from log |

**The blocking problem:** If the coordinator fails after `PREPARE`, participants are stuck holding resource locks forever. This is the fundamental weakness of 2PC.

**Solutions:**
- **3PC (Three-Phase Commit)**: Adds a `PreCommit` phase to break the blocking, but still not partition-tolerant.
- **Saga pattern**: Break distributed transaction into compensatable local transactions with events.
- **Raft-based atomic commit** (CockroachDB): Use Raft consensus instead of 2PC coordinator.

---

### Q4: What is consistent hashing? How is it used in distributed caches and databases?

**Answer:**
**Consistent hashing** distributes keys across nodes such that adding or removing a node only requires remapping a small fraction of keys (1/n on average).

**How it works:**
1. Map both nodes and keys to positions on a **hash ring** (0 to 2^32-1).
2. Each key is assigned to the first node clockwise from its position on the ring.

```
Ring:
    0° ─── Node A (at hash 100)
            │
            ▼
Node C     Key X (hash 300 → go to D)
(at 250)    
            ▼
      Node D (at hash 350)
```

**Adding/removing a node:**
- Old hash ring (N nodes): Adding a node only affects keys between the new node and its predecessor → ~1/N keys remapped.
- Traditional consistent hashing: Adding 1 node → ALL keys must be remapped (100% disruption).

**Virtual nodes (vnodes):**
- Each physical node is represented by multiple virtual positions on the ring (e.g., 150 vnodes per node).
- Better load distribution — prevents "hot spots" when nodes are unevenly distributed on the ring.

**Used in:**
- **Cassandra**: Data sharded using consistent hashing with vnodes. Each row's partition key is hashed to a token on the ring.
- **Redis Cluster**: 16384 hash slots distributed using consistent hashing.
- **Memcached** client-side sharding.
- **DynamoDB** (original Amazon design).

---

### Q5: How do you design an idempotent API? Why is idempotency critical in distributed systems?

**Answer:**
**Idempotent**: An operation can be applied multiple times without changing the result beyond the first application.

**Why critical:**
In distributed systems, requests can be retried due to:
- Network timeouts (didn't get response — did it succeed?)
- Load balancer retries
- Client retry logic
- Message queue at-least-once delivery

Without idempotency, retries cause duplicate side effects (double charges, double orders).

**HTTP method idempotency:**

| Method | Idempotent? | Safe? | Reason |
|---|---|---|---|
| GET | ✓ | ✓ | Read-only |
| PUT | ✓ | ✗ | Full replace — same result every time |
| DELETE | ✓ | ✗ | Deleting already-deleted = same final state |
| POST | ✗ | ✗ | Creates new resource each time |
| PATCH | ✗ (usually) | ✗ | Depends on implementation |

**Implementing idempotent POST (payment example):**

```http
POST /payments
Idempotency-Key: <client-generated-UUID>  ← unique per intended operation

{
  "amount": 1000,
  "account": "ACC123"
}
```

**Server logic:**
1. Hash the `Idempotency-Key` as a lookup key in a store (Redis/DB).
2. If key **not seen before**: Execute the operation, store result with key.
3. If key **already seen**: Return the cached result immediately — don't re-execute.
4. Set TTL on the key (e.g., 24 hours).

**Database-level idempotency:**
```sql
-- UPSERT — idempotent insert
INSERT INTO payments (idempotency_key, amount, status)
VALUES ($key, $amount, 'PENDING')
ON CONFLICT (idempotency_key) DO NOTHING;
```

---

*Prepared for technical rounds at product-based companies (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay, Groww).*
