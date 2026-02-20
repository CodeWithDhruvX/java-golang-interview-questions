# 🔵 Advanced MySQL Interview Questions (Q61–100)

---

### 61. What are isolation levels in MySQL?

"MySQL (InnoDB) supports four transaction isolation levels, defined by the SQL standard:

- **READ UNCOMMITTED**: Can read uncommitted changes from other transactions (dirty reads). Almost never used.
- **READ COMMITTED**: Only reads committed data. Allows non-repeatable reads.
- **REPEATABLE READ**: Default in InnoDB. Guarantees the same row data within a transaction. Prevents dirty and non-repeatable reads.
- **SERIALIZABLE**: Strictest. All reads become locking reads. Prevents phantom reads.

Change with: `SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;`"

#### Indepth
InnoDB's `REPEATABLE READ` prevents phantom reads for **snapshot reads** (regular `SELECT`) using MVCC. However, for **locking reads** (`SELECT ... FOR UPDATE`), phantoms are still possible unless gap locks are held. `SERIALIZABLE` upgrades all plain reads to locking reads, completely preventing phantoms but at the cost of reduced concurrency. In practice, `READ COMMITTED` is often preferred in high-concurrency OLTP systems to reduce lock contention, accepting the non-repeatable read tradeoff.

---

### 62. Explain dirty read, non-repeatable read, and phantom read.

"These are three types of **read anomalies** in concurrent transactions:

- **Dirty Read**: Transaction A reads data that Transaction B has modified but **not yet committed**. If B rolls back, A has read invalid data. (Prevented by READ COMMITTED and above.)
- **Non-Repeatable Read**: Transaction A reads a row twice; between the reads, Transaction B commits an update. A sees different values. (Prevented by REPEATABLE READ and above.)
- **Phantom Read**: Transaction A reads a set of rows matching a condition; Transaction B inserts new matching rows and commits. A's second read finds new 'phantom' rows. (Prevented by SERIALIZABLE.)

Each isolation level is defined by which of these anomalies it permits."

#### Indepth
InnoDB's MVCC cleverly prevents dirty reads and non-repeatable reads without locking by giving each transaction a **consistent read view** (a snapshot) at the start of the transaction. All regular SELECTs read from this snapshot, not the live data. Only `SELECT ... FOR UPDATE` and `SELECT ... LOCK IN SHARE MODE` read current data and acquire locks. This is why InnoDB's REPEATABLE READ has excellent concurrency despite its strong guarantee.

---

### 63. What is a composite primary key?

"A **composite primary key** uses two or more columns together to uniquely identify a row. Neither column alone is unique — their combination is.

```sql
CREATE TABLE enrollment (
  student_id INT,
  course_id INT,
  enrolled_at DATETIME,
  PRIMARY KEY (student_id, course_id)
);
```

A student can appear multiple times (in different courses) and a course can appear multiple times (for different students), but the combination is unique.

I use composite primary keys for **junction/association tables** in many-to-many relationships."

#### Indepth
In InnoDB, the composite primary key is the clustered index. The physical row order is sorted by `(student_id, course_id)`. All secondary indexes contain both PKcomponents as their row pointer, increasing secondary index size proportionally. When using a composite PK, be deliberate about which column comes first — put the most commonly queried first to enable efficient range scans and prefix index usage.

---

### 64. What is indexing in detail?

"An **index** is a separate data structure that MySQL maintains alongside a table to speed up data retrieval. Without an index, querying requires a full table scan — O(n) for every query.

MySQL's primary index type is **B+Tree**: a balanced tree where leaf nodes contain the actual index entries (and in InnoDB's clustered index, the actual row data). Lookups, range scans, and ORDER BY operations are all O(log n) with a proper index.

Other types: **Hash** (Memory engine, exact lookups only), **FULLTEXT** (inverted index for text search), **Spatial** (R-Tree for geo data)."

#### Indepth
The cost of index maintenance during writes is proportional to the number of indexes. Each INSERT must update all indexes. Each UPDATE that modifies an indexed column updates those indexes. InnoDB's **change buffer** amortizes index maintenance for secondary indexes by buffering page changes when the index page is not in the buffer pool, applying them lazily when the page is next read. This significantly reduces write amplification on I/O-heavy secondary index updates.

---

### 65. What is B-Tree indexing?

"A **B+Tree** (the actual structure MySQL uses internally) is a balanced, ordered tree where:
- All data is stored in **leaf nodes**.
- Leaf nodes are **linked** in a doubly linked list, enabling efficient range scans.
- Internal nodes contain only key values used for navigation.

This structure provides O(log n) lookups and efficient range traversals — critical for `WHERE col BETWEEN x AND y` and `ORDER BY col` queries.

I understand B+Tree mechanics because it directly explains index design rules: leftmost prefix, avoided leading wildcards, and why sorted inserts (AUTO_INCREMENT) are faster."

#### Indepth
In InnoDB, B+Tree pages are **16KB by default** (`innodb_page_size`). A page stores many index entries — a typical integer-keyed index holds hundreds of entries per page. The tree is typically 3–4 levels tall even for tables with hundreds of millions of rows. This means most lookups require only 3–4 I/O operations (one per level), which is why B+Tree indexes are so efficient even on massive tables.

---

### 66. What is a covering index?

"A **covering index** is an index that contains ALL columns a query needs — so MySQL can satisfy the query entirely from the index without ever visiting the actual table rows.

```sql
CREATE INDEX idx_covering ON orders(customer_id, status, total);
SELECT status, total FROM orders WHERE customer_id = 42;
```

The query only needs `customer_id`, `status`, and `total` — all in the index. MySQL reads only the compact index structure instead of the full heap rows.

EXPLAIN shows `Using index` in the Extra column when a covering index is used."

#### Indepth
Covering indexes are one of the highest-impact optimizations available. Index pages are smaller and more cacheable than full table pages, so covering index scans fit more data in the buffer pool. For read-heavy analytics queries, a well-placed covering index can reduce I/O by 10×–100×. The strategy: identify the most expensive SELECT queries, find the WHERE + ORDER BY + SELECT columns, and build a composite index containing all of them in optimal order.

---

### 67. What is a query execution plan?

"A **query execution plan** is the sequence of steps MySQL decides to use to execute a query — which indexes to use (or not), what join algorithms to apply, what order to join tables in, whether to sort or use filesort.

`EXPLAIN` reveals the plan. Key fields:
- `type`: Access method (const, ref, range, ALL)
- `key`: Index chosen
- `rows`: Estimated rows scanned
- `Extra`: Additional operations (filesort, temporary, index)

Understanding execution plans is the foundation of query tuning."

#### Indepth
MySQL's **cost-based optimizer** (CBO) evaluates multiple plan alternatives and picks the one with the lowest estimated cost (in I/O units). It uses **statistics** stored in `mysql.innodb_table_stats` and `innodb_index_stats`. These statistics are sampled, not exact. When data is highly skewed (e.g., 90% of rows have `status='active'`), the optimizer may pick a bad plan. Use `ANALYZE TABLE` to refresh statistics or `FORCE INDEX` to override the optimizer when necessary.

---

### 68. How does MySQL handle concurrency?

"MySQL (InnoDB) handles concurrency through:

1. **MVCC (Multi-Version Concurrency Control)**: Readers don't block writers. Each transaction sees a consistent snapshot. No read locks needed.
2. **Row-level locking**: Writers lock only the specific rows they modify — not the entire table.
3. **Isolation levels**: Control the visibility rules between concurrent transactions.
4. **Deadlock detection**: Automatically detects and resolves circular lock dependencies.

Together, these mechanisms allow thousands of concurrent connections to read and write simultaneously with minimal contention."

#### Indepth
InnoDB uses a **lock manager** that maintains a list of active locks. Row locks are stored in memory structures (not on disk). The `INFORMATION_SCHEMA.INNODB_LOCKS` and `INNODB_LOCK_WAITS` tables expose current locks — invaluable for diagnosing lock contention in production. The `performance_schema.data_locks` table in MySQL 8.0 provides even more detailed lock information. Understanding lock types (S locks, X locks, intention locks, gap locks) is essential for advanced MySQL tuning.

---

### 69. What is row-level locking?

"**Row-level locking** means a write operation (INSERT/UPDATE/DELETE) only acquires a lock on the specific row(s) it modifies, not the entire table.

This maximizes concurrency: Transaction A can update row 1 while Transaction B simultaneously updates row 2 — they don't block each other.

InnoDB always uses row-level locking, which is why it vastly outperforms MyISAM (table-level locking) under concurrent write workloads."

#### Indepth
InnoDB implements row-level locking via **record locks** on index entries, not on actual rows. If a query doesn't use an index, InnoDB may lock ALL rows (because it must scan all index entries to find matching rows). This is a critical gotcha: a poorly written UPDATE without a WHERE index condition effectively becomes a table lock. Always ensure UPDATE and DELETE queries use indexed conditions, verified with `EXPLAIN`.

---

### 70. What is table-level locking?

"**Table-level locking** locks the entire table for a write operation, blocking all other reads AND writes until the lock is released.

MyISAM uses table-level locking exclusively. One concurrent write blocks all other table access.

InnoDB acquires table-level **metadata locks** (MDL) for DDL operations (ALTER TABLE, DROP TABLE). Even in InnoDB, an ALTER TABLE running on a busy table can cause a lock queue to build up, starving all subsequent queries. This is why online DDL (`ALTER TABLE ... ALGORITHM=INPLACE`) was crucial."

#### Indepth
MySQL 5.5 introduced **metadata locking (MDL)**. Every query acquires a shared MDL on the tables it accesses. A DDL operation (ALTER TABLE) waits for an exclusive MDL — meaning it waits for ALL in-progress queries to finish. If long-running transactions are open, ALTER TABLE will wait indefinitely and block all subsequent queries behind it in the MDL queue. Tools like `pt-online-schema-change` or `gh-ost` avoid this by applying schema changes without holding long MDL.

---

### 71. What is optimistic vs pessimistic locking?

"**Pessimistic locking**: Assumes conflicts will happen. Acquires locks upfront to prevent concurrent modification.
```sql
SELECT * FROM inventory WHERE id = 1 FOR UPDATE; -- Locks the row
```

**Optimistic locking**: Assumes conflicts are rare. Doesn't lock. Instead, verifies nothing changed before committing:
```sql
UPDATE inventory SET quantity = 9, version = 6
WHERE id = 1 AND version = 5; -- If version changed, update fails
```

I prefer optimistic locking for read-heavy systems where contention is low. For write-heavy or financial systems where conflicts are common, pessimistic locking prevents lost updates cleanly."

#### Indepth
Optimistic locking is commonly implemented with a **version number** or **updated_at timestamp** column. The application reads the row with its version, modifies data in memory, then issues the UPDATE with the original version in the WHERE clause. If `affected_rows = 0`, a conflict occurred and the application retries. This pattern works well in web applications but requires application-level retry logic. Pessimistic locking (`FOR UPDATE`) is safer for financial transactions where data correctness is paramount.

---

### 72. What is sharding?

"**Sharding** is a horizontal scaling technique that partitions data across multiple separate database servers (shards), each holding a subset of the total data. Each shard is an independent MySQL instance.

Example: Users with ID 1–1M on Shard A, 1M–2M on Shard B.

I use sharding when a single MySQL server can't handle the write throughput or dataset size. It's a complex architectural decision — sharding adds application complexity, makes cross-shard queries and transactions very difficult, and complicates operations."

#### Indepth
Sharding strategies: **range-based** (partition by value ranges), **hash-based** (partition by hash of key, distributes uniformly), **directory-based** (lookup table maps keys to shards, most flexible). The biggest challenge is **cross-shard queries** — JOINs across shards require application-level aggregation. Also, rebalancing data when adding new shards is operationally complex. Most teams use managed sharding platforms (Vitess, PlanetScale) rather than implementing sharding from scratch.

---

### 73. What is database clustering?

"**Database clustering** in MySQL context typically refers to **MySQL Cluster (NDB Cluster)** — a distributed, in-memory database that partitions data across multiple nodes for high availability and scalability.

More broadly, it also refers to topologies like **InnoDB Cluster** (Group Replication + MySQL Router + MySQL Shell) that provide automatic primary election, failover, and read scaling.

I use InnoDB Cluster for production HA setups where automated failover and no data loss are required."

#### Indepth
MySQL NDB Cluster stores all data in **shared memory** across data nodes and is designed for extremely low latency and high write throughput for simple key-value lookups. However, it performs poorly for complex JOIN queries and is difficult to operate. InnoDB Cluster (using Group Replication) is a more practical HA solution for typical web workloads — it uses synchronous replication with Paxos consensus, guaranteeing zero data loss on primary failure.

---

### 74. What is a materialized view?

"A **materialized view** is a pre-computed, physically stored result set of a query. Unlike regular views (which recompute every time), a materialized view is stored on disk and refreshed periodically or on demand.

**MySQL doesn't natively support materialized views.** The workaround is to create a summary table and refresh it via:
- A scheduled `event` (MySQL Event Scheduler)
- A trigger on the source table
- Application-side logic after writes

I use summary tables (MySQL's equivalent) for expensive analytics queries that don't need real-time data."

#### Indepth
Without native materialized views, MySQL teams typically maintain **denormalized aggregate tables** manually. For example, a `daily_revenue` table refreshed every hour. The tradeoff: freshness vs query performance. For real-time reporting, some teams use MySQL 8.0 **generated columns** combined with indexes as a lightweight alternative, or stream data changes to a dedicated OLAP system (ClickHouse, BigQuery) for analytics.

---

### 75. How do you optimize slow queries?

"My systematic process for slow query optimization:
1. **Identify**: Enable slow query log (`slow_query_log`, `long_query_time`). Use `pt-query-digest` to find the top offenders.
2. **Analyze**: Run `EXPLAIN` / `EXPLAIN ANALYZE` on the query.
3. **Index**: Add appropriate indexes (composite, covering).
4. **Rewrite**: Simplify subqueries to JOINs, eliminate unnecessary ORDER BY or DISTINCT.
5. **Schema**: Consider denormalization or partitioning for hot tables.
6. **Cache**: Add application-level caching for repeated identical queries.
7. **Validate**: Verify improvement with `EXPLAIN ANALYZE` and actual timing."

#### Indepth
In MySQL 8.0, `EXPLAIN ANALYZE` runs the query and returns **actual vs estimated** row counts and timing at each step — critical for diagnosing cases where the optimizer has wrong statistics. Also check `performance_schema.events_statements_summary_by_digest` for historical query statistics including average execution times, without needing slow query log. For persistent slow queries, look at `InnoDB status` for lock waits contributing to the latency.

---

### 76. What are MySQL logs?

"MySQL maintains several log types:
- **Error log**: MySQL server errors, warnings, startup/shutdown messages.
- **General query log**: Every query received (including `SELECT`). Very verbose, usually disabled in production.
- **Slow query log**: Queries exceeding `long_query_time`. Essential for optimization.
- **Binary log (binlog)**: All data-changing statements/events. Used for replication and point-in-time recovery.
- **Relay log**: On replicas, stores events downloaded from the primary's binlog before applying.
- **InnoDB redo log**: Crash recovery for InnoDB.

I always enable the slow query log in production (`long_query_time=1`)."

#### Indepth
The **binary log** is the most operationally critical log. It can be in three formats: `STATEMENT` (logs SQL text — compact but non-deterministic for some functions), `ROW` (logs actual row changes — safe but verbose), and `MIXED` (uses statement where safe, row otherwise). For replication correctness, use `ROW` format. `binlog_row_image=MINIMAL` reduces row-based binlog size by only logging before/after values of changed columns, not all columns.

---

### 77. What is binary logging?

"**Binary logging (binlog)** is MySQL's mechanism for recording all data modification events. The binlog is a sequence of binary-encoded events describing changes to the database.

It serves two critical purposes:
1. **Replication**: Replicas read the binlog to replicate changes from the primary.
2. **Point-in-time recovery (PITR)**: Replay binlog events after restoring a backup to recover to an exact moment before a failure or error.

Enable with: `log_bin = /var/log/mysql/mysql-bin.log` in `my.cnf`."

#### Indepth
Binlog events are identified by a **position** (log file + offset) or a **GTID** (Global Transaction ID, a globally unique identifier per transaction). GTID-based replication (`gtid_mode=ON`) simplifies failover dramatically — replicas know exactly what they've applied and where to start reading without needing to know log file names and positions. Always configure GTID in modern MySQL deployments.

---

### 78. What is the difference between statement-based and row-based replication?

"**Statement-Based Replication (SBR)**: The primary logs the actual SQL statements. Replicas re-execute the same statements.
- 📦 Compact logs. But non-deterministic functions (`NOW()`, `UUID()`, `RAND()`) can produce different results on replicas → data drift.

**Row-Based Replication (RBR)**: The primary logs the actual row changes (before and after images).
- 🔐 Deterministic — replicas always apply the exact same data changes. But logs are larger.

**Mixed**: Uses SBR by default, switches to RBR for statements containing non-deterministic elements.

I always use ROW-based replication in production for correctness."

#### Indepth
Row-based replication with `binlog_row_image=MINIMAL` is the modern best practice. MINIMAL mode writes only the PK (before image) and changed columns (after image), significantly reducing binlog size vs FULL mode which logs all columns. For high-write workloads, large binlogs increase disk I/O, replication bandwidth, and point-in-time recovery time. Monitor binlog size with `SHOW BINARY LOGS` and implement binlog rotation with `expire_logs_days`.

---

### 79. What is MySQL Workbench?

"**MySQL Workbench** is MySQL's official GUI tool for database administration, development, and design.

Features include:
- **SQL editor** with syntax highlighting and auto-complete
- **EER (Enhanced Entity-Relationship) diagram** for visual schema design
- **Query execution** with EXPLAIN visualization
- **Server administration** (user management, server start/stop)
- **Migration wizard** for importing from other databases
- **Performance dashboard** for monitoring metrics

I use it for schema design and exploring databases locally. For production, I prefer CLI tools (`mysql`, `mysqladmin`, `pt-*`) and monitoring dashboards like Grafana."

#### Indepth
MySQL Workbench's **EXPLAIN tree visualization** is particularly valuable — it renders the execution plan as a graphical tree, making it much easier to spot bottlenecks than reading tabular EXPLAIN output. The **Performance Schema Dashboard** tab connects to `performance_schema` and displays live query statistics, wait events, and InnoDB metrics. However, Workbench is resource-heavy and slow for remote connections. Many DBAs prefer lightweight tools like DBeaver or TablePlus.

---

### 80. What are common performance tuning techniques?

"Key MySQL performance tuning areas:
1. **Query-level**: Add indexes, rewrite queries, eliminate N+1 patterns.
2. **InnoDB buffer pool**: Size it to hold frequently accessed data in memory (`innodb_buffer_pool_size` = 70–80% of RAM).
3. **Connection management**: Use connection pooling (ProxySQL, PgBouncer-equivalent). Set appropriate `max_connections`.
4. **I/O tuning**: `innodb_flush_log_at_trx_commit=2` for non-critical workloads. SSD storage.
5. **Schema design**: Optimal data types, avoid unnecessary NULL columns, partition large tables.
6. **Read/Write splitting**: Route reads to replicas to offload primary."

#### Indepth
The **InnoDB buffer pool** is the single most impactful tuning parameter. It caches frequently accessed table and index pages. A large buffer pool means more data fits in memory, reducing disk I/O dramatically. Monitor buffer pool hit rate: `(1 - (innodb_buffer_pool_reads / innodb_buffer_pool_read_requests)) × 100%`. Target > 99%. Also tune `innodb_log_file_size` (larger = better write throughput but slower crash recovery) and ensure `innodb_flush_method=O_DIRECT` to avoid double-buffering with OS disk cache.

---

### 81. How do you handle large datasets?

"For large datasets I apply multiple strategies:
- **Indexing**: Ensure all query predicates use indexes. Use covering indexes.
- **Partitioning**: Partition large tables by date/range to enable partition pruning.
- **Archiving**: Move old data to archive tables or cold storage to keep hot tables small.
- **Pagination**: Use keyset pagination instead of OFFSET for deep paginated reads.
- **Read replicas**: Distribute read queries across replicas.
- **Caching**: Application-layer cache (Redis) for frequently read large result sets.
- **Sharding**: Horizontal partitioning across multiple servers for extreme scale."

#### Indepth
Large datasets expose **statistics accuracy issues** — MySQL's optimizer samples only a fraction of index pages for statistics. On billion-row tables, estimates can be off by 100×. Increase `innodb_stats_persistent_sample_pages` (default 20) to improve estimate accuracy. Also consider **online archiving**: move rows older than N months to a partitioned archive table in low-traffic windows, or use `pt-archiver` from Percona Toolkit for online, non-blocking row archiving.

---

### 82. What is a deadlock detection mechanism?

"InnoDB uses a **wait-for graph** to detect deadlocks:

Each transaction is a node in the graph. An edge from A → B means A is waiting for a lock held by B. InnoDB periodically traverses this graph looking for cycles. When a cycle is found (deadlock), it selects the transaction with the **least undo log data** (cheapest to roll back) as the victim and terminates it with error `ERROR 1213`.

The victim receives an error and must retry. The surviving transaction proceeds normally.

View the last deadlock: `SHOW ENGINE INNODB STATUS`."

#### Indepth
InnoDB's deadlock detection is **synchronous** — when a transaction requests a lock held by another, it immediately checks for deadlocks before waiting. This adds overhead in extremely high-concurrency systems. MySQL 8.0 added `innodb_deadlock_detect=OFF` to disable detection and rely on lock timeout (`innodb_lock_wait_timeout`) instead — useful in high-contention scenarios where detection overhead is measurable. Always investigate and fix the root cause of frequent deadlocks rather than just relying on the retry mechanism.

---

### 83. What are foreign key constraints and their rules?

"Foreign key constraints enforce **referential integrity** between tables:
- **RESTRICT** (default): Prevents deleting a parent row if child rows reference it.
- **CASCADE**: Automatically deletes child rows when the parent is deleted.
- **SET NULL**: Sets child FK column to NULL when parent is deleted.
- **NO ACTION**: Like RESTRICT but deferred checking (in MySQL, same as RESTRICT).
- **SET DEFAULT**: Sets child FK to default value (rarely supported).

```sql
FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
```"

#### Indepth
Foreign key checks carry a **performance cost**: every INSERT/UPDATE on the child table triggers a lookup in the parent table, and every DELETE on the parent triggers a lookup in the child table (to check/cascade). On high-write workloads, disable FK checks during bulk loads with `SET FOREIGN_KEY_CHECKS=0` (re-enable after and verify integrity separately). Also: the FK column in the child table should be indexed — without an index, parent-side DELETEs trigger full child table scans.

---

### 84. What is a schema in MySQL?

"In MySQL, the terms **schema** and **database** are **synonymous** — `CREATE SCHEMA` and `CREATE DATABASE` do the same thing. A schema is a namespace containing tables, views, stored procedures, functions, and triggers.

In other RDBMS (like PostgreSQL or SQL Server), a schema is a sub-namespace within a database — you can have multiple schemas in one database.

I always name schemas after the application or environment: `myapp_production`, `myapp_staging`."

#### Indepth
MySQL conflates schema and database, which surprises developers from PostgreSQL or Oracle backgrounds. In PostgreSQL, a database contains schemas which contain tables — three hierarchy levels. MySQL has only two: server → database/schema → tables. This means MySQL can't namespace tables within the same logical database, requiring separate database instances for multi-tenant isolation, or prefixing table names as a workaround.

---

### 85. What are events in MySQL?

"MySQL **events** are scheduled tasks managed by the **Event Scheduler** — MySQL's built-in cron-like system.

```sql
CREATE EVENT cleanup_old_sessions
ON SCHEDULE EVERY 1 HOUR
DO
  DELETE FROM sessions WHERE created_at < NOW() - INTERVAL 24 HOUR;
```

I use events for: cleaning up expired records, refreshing summary/aggregate tables, rotating logs, sending reminder flags.

Enable the scheduler: `SET GLOBAL event_scheduler = ON;`"

#### Indepth
MySQL events are stored in the `mysql.event` table and run in the context of a dedicated **event scheduler thread**. They execute with the privileges of the `DEFINER` user. Events are **not replicated by default** — setting `log_bin_trust_function_creators=1` and logging binary log as ROW is needed for proper replication of event-triggered changes. In HA setups with failover, ensure events are re-enabled on the new primary after promotion, as they may default to disabled on replicas.

---

### 86. What is a composite unique index?

"A **composite unique index** enforces uniqueness across a combination of multiple columns — no two rows can have the same values for that specific set of columns together. Individual columns may repeat.

```sql
CREATE UNIQUE INDEX idx_unique_booking
ON bookings(user_id, event_id, date);
```

A user can book different events on the same date, or the same event on different dates — but not the same event on the same date twice.

I use composite unique indexes to enforce business rules like 'a user can only have one active subscription per plan'."

#### Indepth
A composite unique index creates both a **uniqueness constraint** (enforced at write time) and a **B+Tree secondary index** (usable for query optimization) with a single structure. The leftmost prefix rule applies: `(user_id, event_id, date)` can serve queries filtering on `user_id` alone or `(user_id, event_id)` but not `event_id` alone. Design composite unique indexes with both the uniqueness requirement AND query optimization needs in mind.

---

### 87. What is indexing selectivity?

"**Selectivity** is the ratio of distinct values in an indexed column to the total number of rows. Expressed as: `distinct_values / total_rows`.

High selectivity (close to 1.0): `user_id`, `email` — index is very useful.
Low selectivity (close to 0): `status` (only 3 values like active/inactive/suspended) — index may be useless.

MySQL's optimizer uses selectivity to decide whether using an index is worth it. For low-selectivity columns, a full table scan may actually be faster than an index scan."

#### Indepth
MySQL's index selectivity is stored as **cardinality** in `INFORMATION_SCHEMA.STATISTICS`. View it with `SHOW INDEX FROM tablename`. The optimizer estimates selectivity by dividing the stored cardinality by the table row count. For skewed data distributions, even a high-cardinality column may have poor selectivity for a specific value — e.g., in a multi-tenant table, `tenant_id = 'big_client'` may match 80% of rows. **Histograms** (MySQL 8.0+) solve this by capturing value distribution statistics, enabling the optimizer to make better choices on skewed data.

---

### 88. What is query cache?

"MySQL previously had a **query cache** that stored the full result sets of `SELECT` queries. If the same query was executed again and the table hadn't changed, the cached result was returned directly, bypassing the query engine.

However, the query cache was **deprecated in MySQL 5.7.20** and **removed in MySQL 8.0** because:
- It was a global mutex bottleneck — every read/write acquired a lock on the cache.
- Any write to a table invalidated ALL cached queries for that table — poor cache efficiency.
- Under high concurrency, the cache actually degraded performance.

For caching, use application-level solutions: **Redis**, **Memcached**."

#### Indepth
The query cache's failure was a fundamental design flaw: cache invalidation was at the **table granularity**. One `INSERT` into `orders` invalidated all `orders`-related cached queries, even if the affected rows weren't part of those queries. In concurrent write-heavy environments, the cache thrashed constantly. Application caches like Redis can implement **fine-grained invalidation** (by specific data patterns) and don't require a global mutex — making them vastly superior for all practical use cases.

---

### 89. How does MySQL store data internally?

"InnoDB stores data in **tablespace files** (`.ibd` files per table when `innodb_file_per_table=ON`). Data is organized as:
- **Segments**: Large allocations of space (one data segment, one index segment per table).
- **Extents**: 64 consecutive 16KB pages (1MB total).
- **Pages**: 16KB unit of I/O. The default unit for all reads and writes.
- **Rows**: Variable-length records within pages.

The table data (rows) is stored in the **clustered index B+Tree**, where leaf nodes hold the actual row data."

#### Indepth
InnoDB's **row format** (`DYNAMIC` by default since MySQL 5.7) handles large column values: values too large to fit on a page are stored off-page in **overflow pages**, with a 20-byte pointer inline on the main row. `COMPRESSED` row format adds zlib compression per page — reduces storage and I/O for cold data but adds CPU overhead. Understanding row formats helps diagnose row overflow errors and optimize storage for `TEXT`/`BLOB`-heavy tables.

---

### 90. What is a buffer pool?

"The **InnoDB buffer pool** is an in-memory cache for table data and index pages. It's the most important performance component of MySQL.

When MySQL accesses a page, it's loaded from disk into the buffer pool. Subsequent accesses to the same page hit the cache — no disk I/O. Modified pages become 'dirty' and are flushed to disk periodically by background threads.

Size it with: `innodb_buffer_pool_size`. Recommended: 70–80% of total RAM.

A well-sized buffer pool turns a disk-bound workload into a memory-bound one — 10,000× faster."

#### Indepth
The buffer pool uses an **LRU (Least Recently Used) variant** called "Midpoint Insertion Strategy" to prevent large table scans from evicting warm working-set data. The pool is split into a **young sublist** (recently accessed, default 5/8 of pool) and an **old sublist** (recently loaded, 3/8). New pages enter the old sublist. Only pages accessed again after `innodb_old_blocks_time` (default 1 second) are promoted to the young sublist. This prevents a `SELECT *` full table scan from evicting all your hot index pages.

---

### 91. What is the difference between logical and physical backup?

"**Logical backup**: Exports data as SQL statements (`INSERT`, `CREATE TABLE`). Tool: `mysqldump`, MySQL Shell's `dumpInstance`.
- ✅ Portable across MySQL versions and platforms.
- ❌ Slow to restore for large databases (must re-execute SQL).

**Physical backup**: Copies the raw data files (InnoDB `.ibd` files, redo logs).
- ✅ Fast backup and restore — just copy files.
- ❌ Less portable, must match same MySQL version/OS in some cases.
- Tools: Percona XtraBackup, MySQL Enterprise Backup.

I use physical backups for production (speed) and logical for exports/migrations."

#### Indepth
Physical backups with **Percona XtraBackup** are **non-blocking** — XtraBackup copies data files while MySQL runs normally, then applies redo log changes accumulated during the copy (the 'apply log' step). This enables consistent hot backups without downtime. Restoring a physical backup of a 1TB database takes minutes (file copy) vs hours (replaying SQL with `mysqldump`). For point-in-time recovery, combine physical backup with binary log replay.

---

### 92. What is mysqldump?

"`mysqldump` is MySQL's built-in command-line tool for logical backups. It exports database structure and data as SQL statements.

```bash
mysqldump -u root -p --single-transaction --databases mydb > backup.sql
```

`--single-transaction` uses `START TRANSACTION WITH CONSISTENT SNAPSHOT` to create a consistent backup of InnoDB tables without locking them.

I use `mysqldump` for small databases, schema migrations, and data exports. For large (>10GB) databases, Percona XtraBackup is far more efficient."

#### Indepth
`--single-transaction` works for InnoDB because it leverages MVCC — the entire dump runs inside a single transaction with a consistent read view. This means no row locks are held during the dump. However, `--single-transaction` **doesn't work for MyISAM tables** (which don't support transactions). For mixed-engine databases, add `--lock-tables` for MyISAM tables. Also add `--master-data=2` or `--source-data=2` to record the binary log position for use in setting up replicas or PITR.

---

### 93. What is point-in-time recovery?

"**Point-in-time recovery (PITR)** is the ability to restore a database to its exact state at any specific moment — for example, recovering to 2 minutes before an accidental `DROP TABLE`.

Process:
1. Restore the most recent full backup.
2. Replay binary log events from the backup point up to the target time (excluding the bad event).

```bash
mysqlbinlog --start-datetime='...' --stop-datetime='...' binlog.000001 | mysql
```

PITR requires binary logging to be enabled. Without binlogs, you can only restore to the last full backup timestamp."

#### Indepth
PITR is only as good as your binlog retention. Set `binlog_expire_logs_seconds` (MySQL 8.0+) or `expire_logs_days` to retain binlogs long enough to cover your recovery window — typically 7–14 days. Also replicate binlogs to a separate server or cloud storage immediately (using `mysqlbinlog --read-from-remote-server`) in case the primary server is destroyed. PITR is the safety net that turns catastrophic data loss into a 5-minute recovery.

---

### 94. What is a surrogate key?

"A **surrogate key** is an artificial, system-generated key with no business meaning, used as the primary key. The most common example is `INT AUTO_INCREMENT` or `UUID`.

Contrast with a **natural key**, which uses real-world attributes (e.g., SSN, email) as the primary key.

I almost always use surrogate keys because:
- Natural keys can change (email addresses change, SSNs get corrected).
- Natural keys can be long strings (bad for index performance).
- Surrogate keys keep the PK stable and storage-efficient."

#### Indepth
The debate between **UUID vs AUTO_INCREMENT** as surrogate keys involves real performance tradeoffs: AUTO_INCREMENT integers are sequential → efficient clustered index inserts (no page splits). UUIDs are random → cause B+Tree page splits and fragmentation at high insert rates, increasing write I/O and buffer pool pressure. Solutions: use **UUIDv7** (time-ordered UUID) or MySQL 8.0's `UUID_TO_BIN(uuid, 1)` (rearranges bytes for temporal ordering) to get UUID uniqueness with sequential insertion behavior.

---

### 95. What are common MySQL errors and how to fix them?

"Common MySQL errors I encounter:

| Error | Cause | Fix |
|---|---|---|
| `1045 Access denied` | Wrong credentials/privileges | `GRANT` correct privileges |
| `1062 Duplicate entry` | Unique/PK constraint violation | Deduplicate data or use `ON DUPLICATE KEY UPDATE` |
| `1213 Deadlock found` | Circular lock dependency | Fix transaction order; retry application |
| `1040 Too many connections` | `max_connections` reached | Increase limit or add connection pooling |
| `1366 Incorrect integer` | Type mismatch | Fix application data or schema type |
| `2006 MySQL server has gone away` | Timeout or packet too large | Increase `wait_timeout`, `max_allowed_packet` |"

#### Indepth
Error `1040 Too many connections` is especially dangerous because even the DBA can't connect to diagnose it. MySQL reserves one extra connection for the `SUPER` privilege user (`root`). In emergency access scenarios, use `mysql -u root` which gets this reserved connection slot. Prevention: use connection pooling (ProxySQL, HikariCP) to reuse connections efficiently, and set `max_connections` based on `RAM / (per-connection memory)` where per-connection memory is the sum of all per-session buffers.

---

### 96. How does MySQL handle NULL in indexes?

"MySQL **does** index NULL values in B+Tree indexes, unlike some other databases. Rows with NULL in an indexed column are stored at the 'start' of the index (before all non-NULL values in the B+Tree ordering).

This means queries like `WHERE col IS NULL` CAN use an index.

However, **unique indexes allow multiple NULLs** — since SQL specifies NULL ≠ NULL, multiple NULLs don't violate uniqueness.

A primary key column cannot be NULL (NULLs are forbidden in primary keys)."

#### Indepth
In composite indexes, NULLs are handled per column. A composite index `(a, b)` with `a = NULL` stores the entry at the position corresponding to NULL for `a`. For NULL handling in queries, `WHERE a IS NOT NULL` can be optimized using an index if the majority of rows have non-NULL values. InnoDB's MVCC and purge threads handle NULL-valued rows the same as non-NULL rows — NULLs don't get any special treatment in the transaction or locking layer.

---

### 97. What is a JSON data type in MySQL?

"MySQL 5.7.8 introduced the **JSON data type** for storing native JSON documents. It validates JSON on insert, stores it in an optimized binary format, and enables direct path-based access without parsing.

```sql
SELECT data->>'$.name' FROM users;
-- Or equivalently:
SELECT JSON_EXTRACT(data, '$.name') FROM users;
```

I use the JSON type for:
- Storing flexible, schema-less attributes alongside structured data.
- Configuration objects.
- Event payloads that vary by type."

#### Indepth
MySQL stores JSON in a **binary format** (not raw text) that enables fast path-based access without reparsing the entire document. You can also create **functional indexes** on JSON paths in MySQL 8.0:
```sql
CREATE INDEX idx_json_name ON users((data->>'$.name'));
```
This enables indexed lookups into JSON fields — a powerful hybrid of relational and document storage. However, JSON columns are stored in the `DYNAMIC` row format and large documents trigger off-page storage, so keep JSON payloads compact for performance.

---

### 98. What are window functions in MySQL?

"**Window functions** (MySQL 8.0+) perform calculations across a 'window' of rows related to the current row without collapsing them into a single output row (unlike GROUP BY aggregates).

```sql
SELECT name, salary,
  RANK() OVER (PARTITION BY department ORDER BY salary DESC) AS dept_rank,
  AVG(salary) OVER (PARTITION BY department) AS dept_avg_salary
FROM employees;
```

This returns all employee rows with their rank and department average alongside — no GROUP BY collapse.

I use window functions for rankings, running totals, moving averages, and percentile calculations — they replace complex self-joins and subqueries with clean, efficient SQL."

#### Indepth
Window functions are evaluated **after** WHERE, FROM, and GROUP BY but **before** ORDER BY and LIMIT. This means you can't directly filter on a window function result in a WHERE clause — you must wrap it in a subquery or CTE:
```sql
SELECT * FROM (
  SELECT name, RANK() OVER (...) AS rnk FROM employees
) t WHERE rnk <= 5;
```
Performance-wise: window functions with `PARTITION BY` on large datasets can be expensive if there's no index on the partition column. Always check `EXPLAIN` for filesort operations on window function queries.

---

### 99. What is CTE (Common Table Expression)?

"A **CTE** is a named temporary result set defined within a query using a `WITH` clause. It improves readability by breaking complex queries into named logical steps.

```sql
WITH top_customers AS (
  SELECT customer_id, SUM(total) AS total_spent
  FROM orders GROUP BY customer_id
  HAVING total_spent > 10000
)
SELECT c.name, tc.total_spent
FROM customers c
JOIN top_customers tc ON c.id = tc.customer_id;
```

CTEs can also be **recursive** (for hierarchical data). I use CTEs to avoid repeating complex subqueries and to make long queries self-documenting."

#### Indepth
MySQL 8.0 supports CTEs and **recursive CTEs** (`WITH RECURSIVE`). Recursive CTEs enable tree/hierarchy traversal without knowing the depth in advance — replacing cumbersome stored procedures with closure tables. Performance: MySQL may materialize the CTE into a temporary table (optimizer decides). Use `EXPLAIN` to check. Non-recursive CTEs used multiple times in the same query are materialized once, avoiding repeated execution — better than equivalent subqueries which may re-execute.

---

### 100. How do you secure a MySQL database?

"My production MySQL security checklist:

1. **Authentication**: Use strong passwords. Disable anonymous accounts. Remove test database (`DROP DATABASE test`).
2. **Least Privilege**: Grant only required permissions per user/service.
3. **Network**: Bind MySQL to private IP only (`bind-address`). Use firewall rules. Never expose port 3306 publicly.
4. **Encryption**: Enable TLS/SSL for all connections. Enable encryption-at-rest for sensitive data.
5. **Audit**: Enable audit logging (MySQL Enterprise Audit or `general_log` + alerting).
6. **Updates**: Apply MySQL security patches promptly.
7. **SQL Injection**: Use prepared statements in all applications.
8. **Backups**: Test backups regularly. Encrypt backup files."

#### Indepth
Run `mysql_secure_installation` immediately after any fresh MySQL install — it interactively sets the root password, removes anonymous accounts, restricts root login to localhost, and drops the test database. Also enable `validate_password` plugin to enforce password complexity. For production, consider **MySQL Enterprise Edition** for its Firewall (blocks unexpected queries), Audit plugin, Transparent Data Encryption (TDE), and role-based access management. For open-source TDE, Percona Server and MariaDB both offer alternatives.

---
