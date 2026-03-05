# ⚡ Caching — Advanced Interview Questions (Product-Based Companies)

This document covers advanced caching concepts for product-based company interviews (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto). Targeted at 3–10 years of experience rounds.

---

### Q1: How does Redis achieve high availability? Explain Redis Sentinel vs Redis Cluster.

**Answer:**

**Redis Sentinel (HA for single-shard):**
- Monitors a Redis primary + replica setup.
- If primary fails: Sentinels vote and promote a replica to primary automatically (leader election among Sentinels requires quorum).
- Clients reconnect to new primary via Sentinel-provided address.

```
[Sentinel 1] [Sentinel 2] [Sentinel 3] ← 3 Sentinels for quorum (2/3 needed)
           ↓ monitoring
    [Primary] ← [Replica 1] ← [Replica 2]
```

**Failover process:**
1. Sentinel detects primary as `SDOWN` (subjectively down).
2. If quorum of Sentinels agree → `ODOWN` (objectively down).
3. Sentinels elect a Sentinel leader.
4. Leader promotes the most up-to-date replica.
5. Reconfigures other replicas to follow new primary.
6. Notifies clients of new primary address.

**Redis Cluster (sharding + HA):**
- Horizontally partitions data across multiple shards.
- **16384 hash slots** distributed across master nodes.
- Each master has replicas (promoted on failure via Raft-like gossip protocol).
- Clients can connect to any node — nodes redirect with `MOVED` or `ASK` responses.

```
[Shard A: slots 0-5460]     [Shard B: slots 5461-10922]    [Shard C: slots 10923-16383]
 Primary A + Replica A        Primary B + Replica B           Primary C + Replica C
```

**Key difference:**
- Sentinel: High availability only, single shard (all data on one machine). Good for up to ~100GB.
- Cluster: Sharding + HA, multiple shards. Good for TBs of data.

---

### Q2: Explain how you would design a distributed cache with consistent hashing and what challenges you'd face.

**Answer:**

**Design:**
1. Hash each cache node to positions on a virtual ring (0 to 2^32).
2. Use **virtual nodes** (vnodes): Each physical node gets 150 positions on the ring for better distribution.
3. Hash each cache key → position on ring → find successor node → that node stores the key.

```python
# Python-like pseudocode
ring = SortedDict()  # position → node

for node in nodes:
    for vnode in range(150):
        pos = hash(f"{node}:{vnode}") % 2**32
        ring[pos] = node

def get_node(key):
    pos = hash(key) % 2**32
    # Find first position >= pos (clockwise)
    idx = ring.bisect_left(pos)
    if idx == len(ring): idx = 0
    return ring.peekitem(idx)[1]  # return node name
```

**Challenges:**

**1. Hot keys (celebrity problem):**
- A single highly-accessed key overwhelms the node that owns it.
- Solution: **Key replication** — store hot keys on multiple nodes and route reads round-robin.
- Solution: **Local caching in client** (L1 cache) to absorb hot key reads.

**2. Cache warm-up after node failure:**
- New node starts cold — all its hashed keys miss initially → stampede on DB.
- Solution: Pre-warm from replicas, use stale-while-revalidate, gradual traffic shift.

**3. Data locality for multi-key operations:**
- Keys on different nodes can't be atomically updated.
- Redis MGET across cluster → requires `{hash_tag}` for co-location: `{user:42}:cart`, `{user:42}:session` → same slot.

**4. Rebalancing cost:**
- Adding a node: only 1/N portion of keys move. But the actual data migration is still expensive at scale.
- Cassandra uses virtual nodes with token ranges; rebalancing happens gradually.

---

### Q3: How do you implement a cache-aside pattern with cache invalidation for a microservices architecture?

**Answer:**
**Cache-aside (Lazy Loading):** Application manages cache explicitly.
```
Read:
  val = cache.get(key)
  if val is None:
    val = db.get(key)
    cache.set(key, val, ttl=300)
  return val

Write:
  db.update(key, new_val)
  cache.delete(key)  # Invalidate, not update — prevents stale serving
```

**Why invalidate (delete) rather than update on write:**
- Writing to cache on every DB write risks serving the cache get between write and cache update (race condition).
- Delete forces next read to get fresh data from DB.

**Distributed cache invalidation challenges:**

**1. Cache-DB inconsistency window:**
- Sequence: DB write succeeds → cache delete fails (Redis down) → cache serves stale data indefinitely.
- Solution: **Short TTLs** as a fallback. Accept small inconsistency window.
- Solution: **Write-through** for high-consistency data.

**2. Cascade invalidation:**
- A single entity update may need to invalidate multiple cache keys (user data cached by multiple views).
- Solution: **Tag-based invalidation** — group related cache keys under a tag, invalidate the tag.

**3. Cross-service cache invalidation:**
- Service A writes to DB; Service B caches the same entity.
- Solution: **Event-driven invalidation** — publish change event to Kafka. All services subscribe and invalidate their caches on event receipt.

```
Order Service writes order → publishes "order.updated" event to Kafka
Reporting Service (subscriber) → invalidates order-related cache keys
Dashboard Service (subscriber) → invalidates dashboard cache
```

---

### Q4: How does CPU cache work? Why does cache-friendly code matter for performance?

**Answer:**
CPUs have a hierarchy of caches (L1, L2, L3) between registers and RAM.

| Cache | Size | Latency | Notes |
|---|---|---|---|
| L1 | 32–128 KB | ~4 cycles (~1ns) | Per core, split data/instruction |
| L2 | 256 KB–2 MB | ~12 cycles (~4ns) | Per core |
| L3 | 4–64 MB | ~40 cycles (~15ns) | Shared across cores |
| RAM | GBs | ~200 cycles (~60-100ns) | Main memory |
| SSD | TBs | ~100,000 cycles (~30μs) | |

**Cache line:** CPU fetches data in 64-byte chunks (cache lines). If you read `array[0]`, the CPU loads the entire cache line containing `array[0]` through `array[7]` (for int64).

**Cache-friendly code patterns:**

**Row-major vs column-major traversal:**
```go
// Cache-FRIENDLY (row-major — sequential memory access)
for i := 0; i < N; i++ {
    for j := 0; j < M; j++ {
        sum += matrix[i][j]  // sequential, each access hits cache
    }
}

// Cache-UNFRIENDLY (column-major — jumps across rows)
for j := 0; j < M; j++ {
    for i := 0; i < N; i++ {
        sum += matrix[i][j]  // cache miss every N elements
    }
}
```
Column-major can be **10-100x slower** for large matrices due to cache misses.

**Data structure layout (struct of arrays vs array of structs):**
```go
// Array of structs (AoS) — bad if only position needed
type Particle struct { X, Y, Z, Mass, Charge float64 }
particles []*Particle

// Struct of arrays (SoA) — cache-friendly for position-only loop
type Particles struct { X, Y, Z, Mass, Charge []float64 }
// Processing positions: X, Y, Z are contiguous — excellent cache use
```

**False cache line sharing (false sharing):**
When two goroutines/threads write to different variables that happen to be on the same cache line → one invalidates the other's cache → performance degradation.
```go
// Bad: counter[0] and counter[1] may share a cache line
var counter [2]int64

// Good: pad to separate cache lines
type PaddedCounter struct {
    value int64
    _     [56]byte  // pad to 64 bytes (full cache line)
}
```

---

*Prepared for technical rounds at product-based companies (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay, Groww).*
