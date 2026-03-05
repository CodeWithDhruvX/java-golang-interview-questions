# 🗄️ Database Internals — Product-Based Companies

> **Level:** 🔴 Senior / Staff
> **Asked at:** Google, Amazon, Uber, PhonePe, Razorpay — senior data/backend roles

---

## Q1. How does a B-Tree work? Why do databases use B-Trees for indexes?

"A B-Tree (Balanced Tree) is the data structure underlying almost every relational database index — PostgreSQL, MySQL InnoDB, Oracle. It keeps data sorted and balanced, enabling O(log N) search, insert, and delete operations regardless of the dataset size.

The key insight is the branching factor: a B-Tree node can hold hundreds of keys and hundreds of child pointers. A B-Tree with branching factor 500 and depth 4 can index **62.5 billion records** while doing at most **4 disk reads** to find any record. This is why B-Tree indexes are so effective: they minimize disk I/O."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Google, Amazon, Uber — common in senior data engineering and database internals discussions

#### Deep Dive
**B-Tree structure:**
```
A node at depth 0 (root):
    [ptr | key15 | ptr | key40 | ptr | key75 | ptr]
     ↓             ↓             ↓             ↓
  keys<15      15≤k<40       40≤k<75        keys≥75

Each internal node: stores sorted keys + pointers to children
Each leaf node: stores keys + pointers to actual data rows (heap file)

Properties:
  - All leaves at same depth (balanced)
  - Every node ≥ 50% full (except root)
  - Branching factor (order) = typically 100-500 in production
```

**Why not BST (Binary Search Tree)?**
```
BST: branching factor = 2, height = log₂(N)
     1 billion records → height = 30 → 30 disk reads per lookup

B-Tree: branching factor = 500, height = log₅₀₀(N)
     1 billion records → height = 4 → 4 disk reads per lookup

Disk reads are ~1ms each. BST: 30ms. B-Tree: 4ms. That's the difference.
```

**B-Tree vs B+Tree (what PostgreSQL and MySQL actually use):**
```
B-Tree: Internal nodes store keys AND data → more memory per node
B+Tree: Internal nodes store ONLY keys (navigation) → more branching factor
        ALL data stored in leaf nodes → leaves linked in a sorted linked list

B+Tree advantages:
  1. Range scans: traverse the leaf linked list without going up/down the tree
  2. More branching factor → shallower tree → fewer disk reads
  3. All records at same depth → predictable performance

SELECT * FROM orders WHERE created_at BETWEEN '2024-01-01' AND '2024-12-31'
→ B+Tree: find first leaf for '2024-01-01', walk linked list to '2024-12-31'
→ Sequential disk reads → very efficient
```

**Index impact on write performance:**
```
INSERT INTO orders VALUES (...):
  1. Find correct leaf node in B+Tree (log N reads)
  2. Insert key into leaf
  3. If leaf is full → SPLIT: create new leaf, promote middle key to parent
  4. If parent is full → split propagates up the tree (rare but expensive)
  5. Write WAL log entry
  6. Write dirty pages to disk (deferred — page cache)

Per-index overhead: each additional index → additional B+Tree to update on every INSERT/UPDATE/DELETE.
Rule of thumb: index every column you query on, but no more. Tables with 10 indexes on a high-write table will bottleneck on index maintenance.
```

---

## Q2. How does an LSM-Tree (Log-Structured Merge Tree) work? Why do Cassandra and RocksDB use it?

"LSM-Trees optimize for **write-heavy workloads** by making writes always sequential and never in-place. Sequential disk writes are 10-100x faster than random writes, which is why LSM-Trees dominate write-heavy databases (Cassandra, RocksDB, LevelDB, ClickHouse).

The core idea: never update data in-place. All writes go to an in-memory buffer (MemTable), which is periodically flushed to disk as immutable sorted files (SSTables). Reads must merge data across multiple SSTables to find the current value."

#### Company Context & Level
**Level:** 🔴 Staff | **Asked at:** Cassandra-heavy companies, Uber, Google, Amazon — asked when discussing database selection or Cassandra internals

#### Deep Dive
**LSM-Tree architecture:**
```
WRITE PATH (always append → always fast):
  1. Write → WAL (Write-Ahead Log on disk) — crash safety
  2. Write → MemTable (in-memory sorted tree — Red-Black tree or SkipList)
  3. When MemTable hits size limit (128MB): flush to disk as SSTable (Sorted String Table)
  4. SSTable: immutable, sorted by key, with Bloom filter and index

READ PATH (must check multiple levels):
  1. Check MemTable (newest data)
  2. Check L0 SSTables (recent disk flushes, may have overlapping key ranges)
  3. Check L1 SSTables (compacted, non-overlapping)
  4. Check L2, L3, ... (older, larger compaction levels)
  5. Merge results: latest timestamp wins
```

**SSTable + Bloom filter:**
```
SSTable on disk:
  [Index block: {key → byte offset in data block}]
  [Data block: sorted key-value pairs in binary format]
  [Bloom filter: probabilistic — "does key X exist here?"]
  [Footer: pointers to index, bloom filter blocks]

Bloom filter magic:
  Before reading an SSTable: "Is key 'user:456' in this file?"
  Bloom filter says NO → skip this file entirely (no disk read!)
  Bloom filter says YES → read the file (may be false positive, but never false negative)
  
  99% of "not found" lookups eliminated by Bloom filters.
  Without Bloom filters: read all SSTables → O(number of SSTables) disk reads
  With Bloom filters: read ~1-2 SSTables on average
```

**Compaction — background merge process:**
```
Without compaction: SSTables pile up, reads get slower (more files to check)

Compaction: merge multiple SSTables into one larger, deduplicated SSTable
  Input: SSTable_1 [a=1, b=2, c=3], SSTable_2 [b=9, d=4]  (b was updated)
  Output: SSTable_merged [a=1, b=9, c=3, d=4]  (b=2 is a tombstone — discarded)

Compaction strategies:
  Size-Tiered (Cassandra default): merge SSTables of similar size → good for write-heavy
  Leveled (RocksDB default): maintain level size ratio → better for reads, more write amplification

WRITE AMPLIFICATION: data written multiple times (memory → L0 → L1 → L2...)
  RocksDB typical write amplification: 10-30x
  Trade-off: accept write amplification for sequential I/O gains
```

**B-Tree vs LSM-Tree — the fundamental trade-off:**
```
                B-Tree          LSM-Tree
Write latency   Higher (random) Lower (sequential)
Read latency    Lower           Higher (merge across levels)
Space amp       ~1x             ~2-4x (stale versions before compaction)
Write amp       ~2-4x           ~10-30x
Use case        OLTP reads      Write-heavy (IoT, timeseries, Cassandra)
Examples        PostgreSQL, MySQL, Oracle   Cassandra, RocksDB, LevelDB
```

---

## Q3. What is Write-Ahead Logging (WAL) and why is it essential for database crash recovery?

"WAL is the mechanism that allows databases to guarantee durability (the D in ACID) without writing data to the actual table files on every transaction. Instead, every change is first written to a sequential append-only WAL file. If the database crashes, it replays the WAL to recover its state.

Sequential writes to WAL are extremely fast (10-100x vs random writes to data files). The database can defer writing to the actual table files (heap pages) — writing only when convenient (during checkpoint). Meanwhile, WAL provides crash safety."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Senior backend roles, any company using PostgreSQL at scale, Amazon

#### Deep Dive
**PostgreSQL WAL flow:**
```
BEGIN TRANSACTION;
  UPDATE accounts SET balance = balance - 500 WHERE id = 1;
  UPDATE accounts SET balance = balance + 500 WHERE id = 2;
COMMIT;

What PostgreSQL does:
  1. Acquires row locks
  2. Modifies pages in shared_buffers (in-memory page cache)
     → Pages are now "dirty" (modified in memory, not yet on disk)
  3. Writes WAL record to WAL buffer:
     WAL record: {LSN: 0x1F4A2, table: accounts, page: 42, offset: 128, old: 10000, new: 9500}
     WAL record: {LSN: 0x1F4A3, table: accounts, page: 67, offset: 64, old: 5000, new: 5500}
  4. On COMMIT: WAL buffer flushed to WAL file on disk (fsync)
     → Transaction is DURABLE at this point
     → Client receives "commit successful"
  5. Dirty buffer pages written to heap files later (during checkpoint)
     → Can be any time — crash before this is fine, WAL has it

Crash recovery:
  On startup after crash:
    1. Find last checkpoint in WAL
    2. Replay all WAL records after checkpoint
    3. Database reaches consistent state
    4. Open for connections
```

**WAL use beyond crash recovery:**
```
1. Replication (streaming replication):
   Primary writes WAL → WAL shipped to standbys → standbys replay WAL
   This is how PostgreSQL streaming replication works.

2. Logical replication / CDC (Change Data Capture):
   Debezium reads WAL → publishes to Kafka → downstream consumers (Elasticsearch, data warehouse)
   Without WAL, you'd need application-level change tracking.

3. Point-in-time recovery (PITR):
   Daily base backup + continuous WAL archiving to S3
   Restore to any point: restore base backup + replay WAL up to target LSN
   "Restore to exactly 2:47:33 PM on March 4th"
```

**WAL performance tuning:**
```
wal_level = replica          -- Enables replication; 'minimal' reduces WAL volume
synchronous_commit = off     -- Don't wait for WAL fsync on COMMIT (faster, risk last 2 txns on crash)
wal_buffers = 64MB           -- WAL buffer size (default is low — increase for write-heavy workloads)
checkpoint_completion_target = 0.9  -- Spread checkpoint I/O over 90% of checkpoint interval
```

---

## Q4. What is MVCC (Multi-Version Concurrency Control)? How does PostgreSQL implement it?

"MVCC is the mechanism that allows multiple transactions to run concurrently without blocking each other. The key insight: instead of blocking readers when a writer modifies a row, keep multiple **versions** of the same row. Each transaction sees a snapshot of the database at the time it started.

PostgreSQL calls this: 'readers never block writers, writers never block readers.' This is fundamentally different from lock-based approaches where a write would block all concurrent reads of that row."

#### Company Context & Level
**Level:** 🔴 Staff | **Asked at:** Senior PostgreSQL roles, Google, Uber — database internals interviews

#### Deep Dive
**MVCC row versioning in PostgreSQL:**
```
Physical row storage includes hidden columns:
  xmin: transaction ID that INSERTED or CREATED this row version
  xmax: transaction ID that DELETED or UPDATED this row version (0 = still live)

Example: UPDATE accounts SET balance = 9500 WHERE id = 1;

Before UPDATE (row version 1):
  [id=1, balance=10000, xmin=100, xmax=0]    -- created by txn 100, still alive

After UPDATE by transaction 200:
  [id=1, balance=10000, xmin=100, xmax=200]  -- old version: deleted by txn 200
  [id=1, balance=9500,  xmin=200, xmax=0]    -- new version: created by txn 200

A concurrent reader (transaction 150, started before txn 200):
  Reads the row → sees xmin=100 (< 150, so it was committed before I started)
               → sees xmax=200 (200 > 150, so it was deleted AFTER I started → ignore)
  → Sees balance=10000 (old version, as expected for snapshot isolation)

A later reader (transaction 300, started after txn 200 committed):
  Reads the row → first version: xmax=200 (< 300, so deleted before I started → invisible)
               → second version: xmin=200 (< 300, created before I started → visible)
  → Sees balance=9500 (new version, correct)
```

**Transaction isolation levels in PostgreSQL:**
```
READ COMMITTED (default): Each query sees committed data as of query start time.
   → Two queries in same transaction may see different data if committed between them.

REPEATABLE READ: All queries see data as of transaction start time.
   → Consistent snapshot. Can't see writes committed during your transaction.
   → Immune to non-repeatable reads and phantom reads.

SERIALIZABLE: Full isolation. Transactions execute as if serialized.
   → Prevents write skew. Uses SSI (Serializable Snapshot Isolation).
   → Higher overhead — tracks read/write dependencies.

SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;  -- For financial calculations
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;   -- Default for most operations
```

**VACUUM — MVCC's garbage collection:**
```
Problem: Old row versions (dead tuples) accumulate → table bloat → slower full scans

VACUUM process:
  1. Scans the table
  2. Identifies dead tuples (xmax is a committed transaction)
  3. Marks them as free space (doesn't reclaim disk immediately)
  4. Updates the Free Space Map (FSM) so new rows can reuse the space

VACUUM FULL:
  1. Rewrites the entire table, reclaiming all dead space
  2. Requires exclusive lock → tables unavailable during VACUUM FULL
  3. Use sparingly — only when table bloat is extreme

autovacuum: PostgreSQL runs VACUUM automatically in the background.
  autovacuum_vacuum_scale_factor = 0.2  -- Trigger when 20% of table is dead tuples
  autovacuum_analyze_scale_factor = 0.1 -- Trigger ANALYZE for query planner stats
```

---

## Q5. Explain the difference between an index scan, bitmap scan, and sequential scan. When does the query planner choose each?

"The query planner chooses a scan strategy based on estimated cost — primarily how many rows it expects to return and whether those rows are clustered on disk.

**Sequential scan:** Read every page of the table from disk. O(table size). Sounds slow, but sequential disk reads are fast (disk prefetch is effective). Best when reading >5-10% of the table.

**Index scan:** Navigate the B-Tree to find matching row pointers, then fetch each row individually (random I/O into the heap). O(log N + matching rows). Best when matching rows are <1-2% of table.

**Bitmap scan:** Build a bitmap of matching page locations from the index, sort them, fetch pages in order. Middle ground. Best for 1-10% of table with multiple conditions."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Backend roles at companies with PostgreSQL at scale (Razorpay, PhonePe, Swiggy)

#### Deep Dive
```sql
-- EXPLAIN ANALYZE to see what the planner chose:
EXPLAIN ANALYZE SELECT * FROM orders WHERE user_id = 456;

-- Low selectivity (user has 3 orders out of 1M total rows):
→ Index Scan using idx_orders_user_id on orders
   Index Cond: (user_id = 456)
   Rows: 3, cost estimates low → index scan is best

-- High selectivity (status = 'pending' matches 500K out of 1M rows):
EXPLAIN ANALYZE SELECT * FROM orders WHERE status = 'pending';
→ Seq Scan on orders
   Filter: (status = 'pending')
   Rows: 500000 → reading half the table → sequential is cheaper

-- Multiple conditions (bitmap scan):
EXPLAIN ANALYZE SELECT * FROM orders WHERE user_id = 456 AND created_at > '2024-01-01';
→ Bitmap Heap Scan on orders
   Recheck Cond: (user_id = 456 AND created_at > '2024-01-01')
   → Bitmap Index Scan on idx_orders_user_id
   → Bitmap Index Scan on idx_orders_created_at
   → BitmapAnd (intersect both bitmaps)
   → Fetch only matching pages in order
```

**Covering index — eliminating heap access entirely:**
```sql
-- Non-covering index: need heap access for extra columns
SELECT user_id, total_amount, status FROM orders WHERE user_id = 456;
-- Index on user_id → finds rows → heap access for total_amount, status (extra I/O)

-- Covering index: all needed columns in the index
CREATE INDEX idx_orders_covering ON orders(user_id) INCLUDE (total_amount, status);
-- Index only scan → no heap access needed → faster for read-heavy queries

-- PostgreSQL: Index Only Scan (no heap access if visibility map says pages are all-visible)
```
