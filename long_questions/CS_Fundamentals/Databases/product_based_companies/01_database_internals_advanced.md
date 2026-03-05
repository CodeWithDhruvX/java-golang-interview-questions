# 🗄️ Database Internals — Advanced Interview Questions (Product-Based Companies)

This document covers advanced database internals for product-based company interviews (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay). Targeted at 3–10 years of experience rounds.

---

### Q1: How does PostgreSQL implement MVCC (Multi-Version Concurrency Control)?

**Answer:**
**MVCC** allows multiple transactions to read and write concurrently without blocking each other by maintaining multiple versions of each row.

**How PostgreSQL implements MVCC:**

Each row has hidden system columns:
- **`xmin`**: Transaction ID (XID) that inserted/created this row version.
- **`xmax`**: Transaction ID that deleted/updated this row version (0 if currently live).

**On UPDATE:**
- PostgreSQL does NOT update in-place.
- Creates a NEW row version with the new data and `xmin` = current XID.
- Marks old row with `xmax` = current XID (logically deleted, still on disk).

**Visibility rules:**
A transaction with XID `T` can see row version if:
- `xmin` is committed AND `xmin < T` (inserted before T started)
- `xmax` is 0 OR `xmax` is not yet committed when T started

**Implication — Dead Tuples:**
Old row versions accumulate on disk even after commits. `VACUUM` process cleans up dead tuples and reclaims space.

```sql
-- See dead tuples in a table
SELECT n_dead_tup, n_live_tup FROM pg_stat_user_tables WHERE relname = 'orders';

-- Manual vacuum
VACUUM ANALYZE orders;
```

**Isolation levels in PostgreSQL:**
| Level | Dirty Read | Non-Repeatable Read | Phantom Read |
|---|---|---|---|
| Read Committed | ✗ | ✓ (can happen) | ✓ |
| Repeatable Read | ✗ | ✗ | ✗ (PG uses snapshot) |
| Serializable | ✗ | ✗ | ✗ |

---

### Q2: Explain Write-Ahead Logging (WAL). How does it ensure durability?

**Answer:**
**WAL** is the mechanism databases use to ensure durability and enable crash recovery. The core principle: **log the change BEFORE applying it to data files**.

**How WAL works in PostgreSQL:**
1. Transaction begins.
2. All changes (inserts, updates, deletes) are first written to the **WAL buffer** in memory.
3. On `COMMIT`, WAL buffer is **flushed to disk** (WAL segment file — `pg_wal/` directory). This is the durable point.
4. AFTER WAL is durable, the transaction is considered committed (client gets ACK).
5. Modified data pages may still be in memory (buffer pool) — asynchronously flushed to data files later.

**Crash Recovery:**
- On restart, PostgreSQL reads WAL from the last checkpoint.
- Replays all WAL records to bring data files up-to-date.
- This is **redo recovery**.

**WAL uses:**
- **Replication**: PostgreSQL streaming replication ships WAL to standbys in real-time. Standbys replay WAL to stay current.
- **Point-In-Time Recovery (PITR)**: Archive WAL + base backup → restore to any point in time.
- **Logical Decoding**: Tools like Debezium read WAL to stream changes to Kafka (CDC — Change Data Capture).

---

### Q3: How does a database query optimizer work? What is the difference between a sequential scan and an index scan?

**Answer:**
The **query optimizer** transforms a SQL query into the most efficient execution plan.

**Steps:**
1. **Parsing**: SQL text → Parse tree (AST).
2. **Rewriting**: Apply transformation rules (subquery flattening, view expansion).
3. **Planning/Optimization**: Generate candidate plans, estimate costs, pick lowest cost.
4. **Execution**: Execute the chosen plan.

**Cost estimation:**
The optimizer uses **statistics** (row counts, value distributions, histogram, null fraction) maintained in `pg_statistics` to estimate:
- How many rows a predicate will filter.
- Which join algorithm (nested loop, hash join, merge join) is cheapest.
- Whether an index scan or sequential scan is cheaper.

**Sequential Scan vs Index Scan:**

| | Sequential Scan | Index Scan | Bitmap Index Scan |
|---|---|---|---|
| How | Read entire table sequentially | Use index to find specific rows, fetch randomly | Collect row pointers in bitmap, then heap fetch |
| I/O pattern | Sequential (fast, prefetchable) | Random (expensive for HDD, OK for SSD) | Semi-sequential |
| Best for | Large % of rows (>5-20% of table) | Highly selective (< 1-5% of rows) | Medium selectivity + multiple indexes |

**Forcing plan inspection:**
```sql
EXPLAIN (ANALYZE, BUFFERS, FORMAT TEXT) 
SELECT * FROM orders WHERE user_id = 42 AND status = 'PENDING';
```
Look at: `Seq Scan` vs `Index Scan`, actual vs estimated rows, buffer hits vs reads.

---

### Q4: What are LSM Trees and how do they differ from B-Trees for write-heavy workloads?

**Answer:**

**B-Tree (used by PostgreSQL, MySQL InnoDB):**
- In-place updates. Writes require random disk I/O.
- Read performance excellent.
- Write amplification: Each write may touch multiple tree nodes.
- Fragmentation over time — need to rebuild.

**LSM Tree (used by RocksDB, Cassandra, LevelDB, HBase):**
- **MemTable**: Writes go to in-memory sorted structure (very fast).
- When MemTable is full → flush to an immutable SSTable file on disk (sequential write — very fast).
- **Compaction**: Background process merges SSTables, removes deleted/stale entries.

**LSM Write Flow:**
```
Write → WAL (crash safety) + MemTable (L0 in-memory)
      → MemTable full → SSTable (L1)
      → Compaction merges L1 → L2 → L3 ...
```

**Comparison:**

| | B-Tree | LSM Tree |
|---|---|---|
| Write performance | Lower (random I/O) | Higher (sequential writes) |
| Read performance | Higher (direct lookup) | Lower (may check multiple SSTables) |
| Write amplification | Medium | Higher during compaction |
| Space amplification | Lower | Higher (deleted keys still in old SSTables until compaction) |
| Typical use | OLTP (PostgreSQL) | High-write workloads (Cassandra, RocksDB, Kafka log segment) |

**Bloom filters:** LSM databases use bloom filters per SSTable to answer "is this key in this SSTable?" without reading the file. Reduces unnecessary I/O for reads.

---

### Q5: How do distributed databases handle consensus? Explain the Raft protocol.

**Answer:**
In distributed databases, multiple nodes must agree on the same data. **Raft** is a consensus algorithm that ensures a cluster of nodes agrees on a sequence of operations (replacing the complex Paxos).

**Raft roles:**
- **Leader**: Receives all writes, replicates to followers.
- **Follower**: Receives replicated log entries from leader.
- **Candidate**: Temporarily becomes candidate during election.

**Leader Election:**
1. On startup, all nodes are followers.
2. Each follower has a random election timeout (150-300ms).
3. First timeout expires → node becomes Candidate. Sends `RequestVote` RPC.
4. Majority votes → becomes Leader. Sends heartbeats to maintain authority.
5. If 2 leaders tie → new election with higher term.

**Log Replication:**
1. Client sends write to Leader.
2. Leader appends to its log, sends `AppendEntries` to all Followers in parallel.
3. Once **majority** (n/2+1) acknowledges → Leader commits the entry.
4. Leader responds to client. Tells followers to commit in next heartbeat.

**Used in:**
- **etcd**: Kubernetes configuration store (uses Raft).
- **CockroachDB**: Distributed ACID transactions using Raft per range.
- **TiKV**: Distributed key-value store (Raft groups).

---

*Prepared for technical rounds at product-based companies (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay, Groww).*
