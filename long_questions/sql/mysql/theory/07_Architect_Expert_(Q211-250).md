# 🔴 Architect / Expert Level MySQL Questions (Q211–250)

---

### 211. How would you design a highly available MySQL architecture?

"My target HA architecture for production:

1. **InnoDB Cluster** (Group Replication + MySQL Router): Provides automatic primary election and failover (60–90 second recovery).
2. Or: **Primary + 2 replicas with semi-sync replication** + Orchestrator for automated failover.
3. **MySQL Router / ProxySQL** in front: Routes write traffic to primary, reads to replicas. Automatic reroute on failover.
4. **Cross-AZ placement**: Primary in AZ-1, replicas in AZ-2 and AZ-3. Survives single availability zone failure.
5. **Automated backups**: Daily XtraBackup + continuous binary log shipping to object storage.

This achieves RPO~0 (zero data loss with semi-sync) and RTO < 2 minutes (automated failover)."

#### Indepth
The hardest part of HA is **split-brain prevention** — two nodes both believing they're the primary after a network partition. InnoDB Cluster uses Paxos consensus (requires majority quorum) to prevent split-brain: if a partition isolates one node from the majority, it can't accept writes. Orchestrator uses a similar approach with topology locks. Without split-brain prevention, a failover can result in two primaries simultaneously, causing data divergence that's extremely difficult to resolve. Always use odd-numbered node counts (3, 5) for quorum-based HA.

---

### 212. How would you scale MySQL to handle millions of QPS?

"Scaling MySQL to millions of QPS requires a layered approach:

1. **Caching layer (Redis)**: Cache hot data at the application layer. 80% of read queries often hit < 20% of data (Pareto principle). Cache this 20%.
2. **Read replicas**: Horizontal read scaling. 10 replicas = ~10× read capacity. Route via ProxySQL.
3. **Connection pooling (ProxySQL)**: Handle millions of application-level connections with hundreds of MySQL connections.
4. **Vertical scaling**: Largest available instance first — 128+ CPU cores, 1TB+ RAM, NVMe SSDs.
5. **Write sharding**: When writes saturate a single primary, shard by tenant/region/user-range.
6. **CQRS**: Separate read models from write models using event-driven architecture."

#### Indepth
MySQL's single-primary write throughput ceiling is approximately **100,000–300,000 simple transactions/second** on top hardware. Beyond this, sharding is unavoidable. The key metric is **writes per second** — reads scale linearly with replicas. Most applications hit read limits before write limits (typically 80–90% reads). For write-heavy systems (ad auction bidding, IoT telemetry), MySQL alone is insufficient beyond ~500K writes/second regardless of hardware — at that point, use specialized systems (Kafka → ClickHouse, DynamoDB, or custom sharded MySQL) designed for extreme write throughput.

---

### 213. When would you choose sharding over replication?

"Choose sharding when:
- **Write throughput** exceeds a single primary's capacity (replication doesn't help writes).
- **Dataset size** exceeds a single server's storage/RAM (buffer pool can't hold working set).
- **Query latency** for a specific tenant's data must be isolated from other tenants.

Continue with replication when:
- **Read** load needs to scale (replication scales reads linearly).
- Dataset fits on one well-tuned server.
- Cross-table queries are critical (sharding breaks JOIN across shards).

I only recommend sharding after exhausting all single-server optimizations — sharding adds 10× operational complexity."

#### Indepth
A common mistake: sharding too early. Before sharding, exhaust: vertical scaling, read replicas, caching, query optimization, connection pooling, Vitess (transparent MySQL sharding middleware), and business logic changes to reduce write contention. Sharding requires the application to know about shard topology (or use a transparent proxy like Vitess/ProxySQL). Cross-shard transactions require distributed 2PC or application-level saga patterns. Cross-shard JOINs require scatter-gather (query all shards). Each of these is a significant engineering and operational burden — defer sharding as long as possible.

---

### 214. How do you design shard keys?

"A shard key (partition key) determines which shard stores which data. Good shard key properties:

1. **High cardinality**: Enough distinct values to distribute data evenly across shards.
2. **Even distribution**: Avoid hot spots — shard key should spread write load equally.
3. **Query alignment**: Most queries should be bounded to a single shard (minimizing scatter-gather).
4. **Immutability**: Never changes post-insertion (changing shard key requires moving data across shards).

Common choices: `user_id` for user-centric applications, `tenant_id` for SaaS, `created_at` time bucket for time-series data."

#### Indepth
The worst shard key anti-patterns: (1) **Monotonically increasing keys** (like timestamps or AUTO_INCREMENT IDs) — all writes go to the latest shard, creating a hot shard. Use **consistent hashing** of the shard key to distribute load. (2) **Low cardinality keys** (status, country) — few distinct values means few shards can hold the data; can't rebalance. (3) **Compound keys that change** — moving data between shards during rebalancing is expensive. Design shard keys to be `hash(user_id) % shard_count` using a consistent hashing ring, allowing smooth shard addition without full resharding.

---

### 215. What is consistent hashing?

"**Consistent hashing** is an algorithm for distributing data across nodes such that when nodes are added or removed, **only a fraction of keys need to be redistributed** — not all of them.

In a traditional hash ring: hash(key) % N. Adding a new node changes N and invalidates ~100% of mappings.

In consistent hashing: keys and nodes are mapped to positions on a circular ring. A key is assigned to the nearest clockwise node. Adding a node only redistributes keys between it and its predecessor — ~1/N keys on average.

I use consistent hashing for shard key assignment and also for cache cluster distribution (Redis Cluster, Memcached)."

#### Indepth
To prevent uneven distribution due to sparse node placement, consistent hashing uses **virtual nodes (vnodes)**: each physical node is represented by many (e.g., 150) positions on the ring. This creates a more statistically uniform distribution of key assignments. When a node is added, it takes over some vnodes from adjacent nodes — distributing the rebalancing load across many nodes rather than one. This is exactly how **Cassandra**, **DynamoDB**, and **Redis Cluster** implement their sharding — all based on consistent hashing with vnodes.

---

### 216. What is the CAP theorem in the context of MySQL?

"The **CAP theorem** states that a distributed data system can guarantee at most two of three properties simultaneously:
- **Consistency (C)**: Every read receives the most recent write or an error.
- **Availability (A)**: Every request receives a (non-error) response — maybe stale.
- **Partition tolerance (P)**: The system continues operating during network partitions.

**MySQL** is a **CP system**: it prioritizes Consistency over Availability. On a network partition, a MySQL primary will stop accepting writes before risking data inconsistency (with semi-sync and Group Replication quorum requirements). Read replicas can be AP in async mode (stale reads but available)."

#### Indepth
CAP is often simplified; **PACELC** extends it: even without partitions, there's a latency-consistency tradeoff. MySQL with `sync_binlog=1` + `innodb_flush_log_at_trx_commit=1` favors strict consistency at the cost of write latency. Setting these to 0 or 2 favors latency at the cost of potential data loss — shifting toward the AP side temporarily. Modern distributed SQL databases (Spanner, CockroachDB) achieve global consistency with Paxos consensus, providing CP guarantees at global scale — something standard MySQL replication can't provide due to asynchronous replication lag.

---

### 217. How do you implement multi-region MySQL deployment?

"Multi-region MySQL deployment adds latency and consistency challenges:

1. **Active-Passive**: Primary in Region A, read replicas in Region B/C. All writes to Region A. Cross-region replication follows async. RTO: 2–10 minutes on failover.

2. **Active-Active with conflict resolution**: Writes accepted in any region. Use **Group Replication** or **Galera Cluster** for synchronous multi-primary. Cross-region latency (50–200ms) becomes commit latency — typically acceptable only for reads with eventual consistency.

3. **Geo-partitioned sharding**: Users in the US shard to US MySQL; EU users to EU MySQL. Each region's MySQL is a single-primary cluster. Cross-region queries are rare by design."

#### Indepth
The fundamental challenge is **cross-region commit latency**. Synchronous replication (Group Replication) between US and EU has 100ms+ RTT for each transaction ACK — catastrophic for OLTP (most databases commit < 10ms locally). Options: (1) Accept asynchronous replication and the RPO this implies. (2) Use NewSQL databases (Spanner, CockroachDB, TiDB) designed for global synchronous replication with TrueTime or Raft-over-WAN. (3) Design application to be region-local with eventual consistency between regions and strict consistency only within a region.

---

### 218. What is data partitioning strategy?

"A **data partitioning strategy** defines how data is physically divided across storage units (partitions, shards, nodes).

Types:
- **Horizontal partitioning (sharding)**: Rows split across multiple tables/servers by a partition key.
- **Vertical partitioning**: Columns split — rarely accessed columns in separate tables (reduces row width).
- **Range partitioning**: Rows with key values in specific ranges go to specific partitions (e.g., by year, by ID range).
- **Hash partitioning**: Key is hashed to uniformly distribute rows.
- **List partitioning**: Specific discrete values go to specific partitions.

I choose the strategy based on query patterns and growth projections."

#### Indepth
MySQL's built-in table partitioning (RANGE, LIST, HASH, KEY) is for single-server partitioning — it improves `partition pruning` for queries that filter on the partition key and simplifies archiving (drop a partition instead of a slow bulk DELETE). It does NOT distribute load across servers. For multi-server distribution, use application-level sharding or Vitess. Partition pruning only works when the `WHERE` clause can be statically evaluated against the partition function — `WHERE created_at > CURDATE() - INTERVAL 7 DAY` enables pruning; `WHERE created_at > user_input_date` may not, depending on how MySQL evaluates the expression.

---

### 219. What is table partition pruning?

"**Partition pruning** is a MySQL optimization where the query planner identifies which partitions can possibly contain matching rows and skips all others — scanning only relevant partitions.

```sql
-- Table partitioned by RANGE on year(order_date):
SELECT * FROM orders WHERE order_date BETWEEN '2023-01-01' AND '2023-12-31';
-- MySQL scans only the 2023 partition, skips all others.
```

EXPLAIN shows `partitions` column — lists which partitions will be scanned.

This makes queries on time-series partitioned tables dramatically faster — scanning 1 partition out of 10 is a 10× I/O reduction."

#### Indepth
Partition pruning only works when the partition key is directly included in the `WHERE` clause with a **constant expression** that MySQL can evaluate at optimization time. `WHERE YEAR(order_date) = 2023` might NOT prune if the partition function is `RANGE (TO_DAYS(order_date))` — the expressions must match. Also: `JOIN` queries may prevent pruning if the join condition references the partition key from the other table (not a constant). Always verify pruning with `EXPLAIN` and look at the `partitions` column. Unexpected multi-partition scans are a common performance regression in partitioned tables.

---

### 220. What is global transaction ID consistency across shards?

"In a sharded system, each shard has independent transactions. There's no built-in global transaction management across shards.

Challenges:
- Reads that span multiple shards may see data from different transaction points in time.
- Writes that span multiple shards (cross-shard transactions) require coordination.

Solutions:
- **Application-level 2PC**: Lock + prepare on all shards, then commit on all. Complex, failure-prone.
- **Saga pattern**: Break cross-shard work into a sequence of local transactions with compensating rollback actions.
- **Single shard design**: Design data model to avoid cross-shard operations.
- **Vitess/PlanetScale**: Provides VTGate which handles some cross-shard consistency via VReplication."

#### Indepth
Cross-shard consistency is the hardest unsolved problem in sharded MySQL architectures. The Saga pattern is the most practical solution: each step is a local transaction; if a step fails, execute compensating transactions (inverse operations) in reverse order to undo previous steps. This provides eventual consistency rather than ACID atomicity. For use cases requiring strict atomicity across shards, consider NewSQL databases (Spanner, CockroachDB) that natively support globally consistent distributed transactions with Paxos-based 2PC that's crash-safe.

---

### 221. How do you prevent hot partitions?

"A **hot partition** occurs when disproportionate traffic concentrates on one shard/partition — usually because the shard key doesn't distribute load evenly.

Prevention strategies:
- **Avoid monotonic shard keys**: Auto-increment IDs or timestamps → all writes go to the latest partition. Use hash(id) instead.
- **High-cardinality keys**: Low-cardinality keys (country, category) create too few buckets.
- **Virtual nodes**: Assign many virtual node positions per physical shard to smooth distribution.
- **Monitor key distribution**: Regularly measure per-shard query rate and data size. Rebalance early.
- **Shard splitting**: Identify hot shards and split them into two sub-shards."

#### Indepth
Detecting hot partitions: monitor CPU, IOPS, and QPS per MySQL shard instance. A hot shard shows much higher CPU/IOPS than peers. In Vitess, the VTTablet exposes per-shard metrics through Prometheus. For time-series partitioned tables, prevent hot partitions with **sub-partitioning**: `RANGE` by year, sub-partitioned by `HASH(user_id)`. This spreads the current year's writes across hash sub-partitions. The trade-off: more partitions = more metadata overhead and file handles. Balance based on measured write distribution.

---

### 222. How do you handle schema changes in distributed systems?

"Schema changes in distributed/sharded systems are complex because shards are independent databases. Strategies:

1. **Backward-compatible migrations**: Always add columns with defaults; never rename or drop columns in one step. Use expand-then-contract pattern:
   - Step 1 (Expand): Add new column. Old code still works.
   - Step 2 (Migrate): Backfill data. Deploy new code that uses both old and new.
   - Step 3 (Contract): Remove old column after old code is gone.

2. **Rolling migrations**: Apply schema changes shard-by-shard while the system remains live.
3. **Online DDL tools**: gh-ost or pt-osc on each shard sequentially."

#### Indepth
The **expand-then-contract (parallel change)** pattern is essential for zero-downtime schema evolution. It ensures: (1) the schema is always compatible with both the old and new application version simultaneously, enabling safe deploys. (2) No blocking DDL during peak traffic — changes are phased over multiple deploy cycles. Tools like **Liquibase** or **Flyway** support this by allowing you to version schema changes and apply them in coordinated phases. For sharded systems, run migrations on a canary shard first, validate, then roll out to all shards via automation.

---

### 223. How do you handle zero-downtime deployment?

"Zero-downtime database deployment requires application and schema changes to be decoupled:

1. **Backward-compatible schema change**: Deploy the schema change first. Ensure the new schema works with the OLD application code.
2. Deploy the new application code (rolling deploy — old + new code run simultaneously).
3. After all pods/servers run the new code, optionally make a final schema cleanup.

Key principle: Database changes must always be backward-compatible with the previous application version.

Tools: GitHub's Scientist library (dark reads for validation), feature flags to gradually enable new DB paths."

#### Indepth
The most common zero-downtime deployment failure: a developer adds a `NOT NULL` column without a default and deploys the schema before deploying the application code that populates it. Old app code trying to insert without the new column fails. Solution checklist before every migration: (1) Can the old app code work with the new schema? (2) Can the new app code work with the old schema? If both are YES, you have a safe zero-downtime migration. Platform-level db migration tools (Rails strong_migrations, Laravel Safe Migrations) encode these rules and block unsafe migration patterns automatically.

---

### 224. What is blue-green deployment in database context?

"**Blue-green deployment** for databases means maintaining two production environments — Blue (current) and Green (new) — where one is live and the other is idle.

Application deployment:
- Green environment gets the new schema/code first.
- Green is validated (smoke tests, load tests).
- Traffic is cut over from Blue to Green atomically (DNS/load balancer switch).
- Blue remains as rollback target for a short period.

For databases, the challenge is **data synchronization**: during cutover, writes to Blue must also appear in Green. Achieved via replication or dual-write."

#### Indepth
Pure blue-green for databases is harder than for stateless apps because database state is mutable and shared. Real-world approach: (1) Set up Green as a replica of Blue. (2) While replicating, Blue is live and Green is catching up. (3) At cutover: make Blue read-only, wait for Green to fully sync (seconds with low lag), update DNS/proxy to point to Green, make Green the new primary. (4) Maintain Blue as a read-only backup for rollback for 15–30 minutes. Blue-green works best when combined with GTID-based replication for precise cutover point tracking.

---

### 225. What is canary database deployment?

"A **canary database deployment** gradually rolls out schema or query changes to a small percentage of traffic first, monitors for issues, then expands rollout.

Database canary approaches:
- **Invisible index**: Mark the new index as INVISIBLE, validate with specific queries using `USE INDEX`, then make it VISIBLE.
- **New column dark launch**: Add a new column but only write to it in 1% of requests. Monitor errors/latency, then ramp to 100%.
- **Read shadow traffic**: Duplicate reads to a test database with the new schema, compare results with production.
- **Feature flags**: Enable new query paths for 1% of users, comparing performance and correctness.

This reduces the blast radius of schema changes."

#### Indepth
Canary deployments require **observability infrastructure**: metrics per query type, error rates, latency percentiles. MySQL's `performance_schema.events_statements_summary_by_digest` tracks per-query statistics automatically. Compare digest performance between canary (1% of queries using new schema pattern) and control (99% using old pattern). Automated canary monitoring with automatic rollback (if error rate or P99 latency spikes) makes canary deployments safe enough to run continuously in CI/CD pipelines without human monitoring during off-hours.

---

### 226. How do you monitor MySQL in production?

"My MySQL monitoring stack:

1. **Metrics collection**: `mysqld_exporter` (Prometheus) + Grafana dashboard (PMM or dedicated MySQL dashboards).
2. **Key metrics**: QPS, TPS, connection count, buffer pool hit rate, replication lag, InnoDB row ops, slow query count.
3. **Slow query log** → `pt-query-digest` for query analysis.
4. **Alert thresholds**: Replication lag > 30s, buffer pool hit rate < 95%, connections > 80% of max, disk > 80%.
5. **Percona Monitoring and Management (PMM)**: All-in-one monitoring specifically for MySQL.
6. **Real-time**: `SHOW ENGINE INNODB STATUS`, `SHOW PROCESSLIST`, `sys.innodb_lock_waits`."

#### Indepth
The most critical MySQL metrics to alert on: (1) `Threads_running` sustained > CPU cores (query pileup). (2) `innodb_buffer_pool_reads` / `innodb_buffer_pool_read_requests` ratio < 99% (working set not in memory). (3) `Seconds_Behind_Master` > threshold (replication lag). (4) Disk I/O utilization > 80% on the MySQL data volume. (5) `innodb_row_lock_waits` rate increasing (lock contention). Set up Grafana with pre-built MySQL dashboards from Percona (free), or use PMM (Percona Monitoring and Management) for full stack MySQL observability including query execution data.

---

### 227. What metrics are critical for MySQL health?

"Key MySQL health metrics by category:

**Query Performance**: `Questions/s` (QPS), `Slow_queries/s`, `Select_full_join`, `Sort_rows/s`.

**InnoDB**: Buffer pool hit rate, `Innodb_row_lock_waits`, `Innodb_os_log_fsyncs`, dirty page ratio.

**Connections**: `Threads_connected`, `Threads_running`, `Connection_errors_max_connections`.

**Replication**: `Seconds_Behind_Master`, GTID gap, I/O thread and SQL thread status.

**Disk/Storage**: Data volume utilization, binary log size, undo tablespace size.

**Process**: CPU usage, memory usage, disk IOPS."

#### Indepth
**`Threads_running`** is the single most actionable real-time health indicator. It shows threads actively executing queries (not just connected). In a healthy system, this should be well below the number of CPU cores. If `Threads_running` approaches CPU core count and stays high, queries are piling up — imminent performance degradation. Compare with `Threads_connected` (all open connections, mostly idle). A healthy ratio: `Threads_running / Threads_connected < 0.1`. When the ratio approaches 1.0, the system is under extreme query pressure and intervention is urgent.

---

### 228. How do you capacity plan for MySQL?

"MySQL capacity planning process:

1. **Baseline current metrics**: QPS, storage growth rate/month, buffer pool hit rate, CPU%, peak connection count.
2. **Growth projection**: Estimate traffic growth 12–24 months out. Common metrics: new users/month, orders/day, event rate.
3. **Headroom targets**: Plan for 40–50% CPU headroom and 50% storage headroom at all times.
4. **Storage**: `current_size + growth_rate * months_ahead + safety_factor`.
5. **Read replicas**: Add replicas when primary CPU > 60% sustained from reads.
6. **Vertical scaling triggers**: When buffer pool hit rate drops below 95%, upgrade RAM first."

#### Indepth
**Storage is the most predictable dimension**: measure current DB size + binlog retention + backup size, and project growth linearly. **Memory is the most impactful**: every time the buffer pool hit rate drops below 99%, the system is doing significantly more disk I/O. Calculate working set size from `innodb_buffer_pool_reads` and `innodb_buffer_pool_read_requests` trends — when hit rate starts declining, it's time to scale RAM. **CPU is the most variable**: depends on query mix changes, not just traffic. Profile the top 10 query patterns for CPU cost and model CPU demand as a function of business growth metrics.

---

### 229. How do you benchmark MySQL?

"Tools for MySQL benchmarking:

**sysbench**: The standard MySQL benchmark tool for OLTP workloads.
```bash
sysbench oltp_read_write --table-size=1000000 --threads=64 \
  --mysql-host=localhost --mysql-user=root --time=60 run
```

**mysqlslap**: Built-in MySQL tool for simple concurrency benchmarks.

**custom workload replay**: Capture production queries with the slow query log, replay them at controlled rates.

I always benchmark on hardware identical to production, with realistic data sizes and distributions, not tiny test datasets."

#### Indepth
Benchmark pitfalls to avoid: (1) **too-small dataset** — the entire table fits in buffer pool, masking real I/O patterns. Use at least 10× the buffer pool size. (2) **uniform data distribution** — synthetic data doesn't capture skew (hot rows, hot indexes) that production workloads create. (3) **ignoring replication overhead** — benchmark on a primary+replica topology, not standalone. (4) **short duration** — MySQL performance shifts over time as InnoDB buffer pool warms up, adaptive hash index builds, and checkpoint flushing kicks in. Run benchmarks for at least 30–60 minutes after a warm-up period.

---

### 230. What is sysbench?

"**sysbench** is an open-source multi-threaded database benchmarking tool designed specifically for OLTP database workloads.

Key OLTP benchmarks:
- `oltp_read_only`: Simulates read-intensive workload (SELECTs, index lookups, range queries).
- `oltp_write_only`: Simulates write-intensive workload (INSERT, UPDATE, DELETE).
- `oltp_read_write`: Mixed workload (the most realistic default).
- `oltp_point_select`: Pure primary key lookups — tests maximum throughput for simple queries.

Output: TPS (transactions/second), latency distribution (min/avg/max/95th percentile), errors."

#### Indepth
sysbench's `oltp_read_write` benchmark models a simplified bank transaction: SELECT balance, read 10 rows, UPDATE 2 rows with indexed updates, DELETE 1 row, INSERT 1 row. It's a good approximation of common web application patterns. For more realistic benchmarks, record production traffic (via `pt-query-digest` against the slow log) and **replay it at controlled rates** using Percona's `pt-log-player`. This replay approach captures actual query distribution, hot spots, and access patterns that sysbench's synthetic workload cannot replicate.

---

### 231. What is pt-query-digest?

"**pt-query-digest** is a Percona Toolkit tool that analyzes MySQL slow query logs and provides detailed performance analysis.

```bash
pt-query-digest /var/log/mysql/mysql-slow.log
```

Output includes for each unique query fingerprint: total execution count, total time, average/max/percentile latencies, rows examined vs sent ratio, and the normalized query text.

I run `pt-query-digest` weekly on the slow query log to identify the top 10 queries by total time consumed — these are the highest-impact optimization targets."

#### Indepth
`pt-query-digest` can analyze not just slow query logs but also general query logs, binary logs (via `mysqlbinlog --verbose`), and `SHOW PROCESSLIST` captures. The most important metric in its output: **total time** (not just avg time) — a query that takes 5ms on average but runs 1 million times/day consumes 5000 seconds of total database time and is a high-priority optimization target even though it seems fast individually. Also output includes queries with `Full_table_scans` and `Full_joins` highlighted, directly actionable for index optimization.

---

### 232. What is replication conflict resolution?

"**Replication conflict** occurs in multi-primary replication (Group Replication, Galera) when two nodes simultaneously modify the same row. Conflict resolution determines which transaction wins.

Group Replication resolution:
- Uses **certification-based** conflict detection: before committing, each transaction's write-set is broadcast to all nodes.
- If two transactions modify the same rows, only the first-committed transaction wins.
- The loser is **rolled back** with `ER_TRANSACTION_ROLLBACK_DUE_TO_CONFLICT` (error 3101).
- The client must retry the rolled-back transaction.

Single-primary mode avoids conflicts entirely (only one node accepts writes)."

#### Indepth
Conflict rate is a key health metric in multi-primary replication. A high conflict rate indicates poor shard key design (many cross-node conflicts) or hot rows accessed by many clients simultaneously. Monitor with `group_replication_stats.count_conflict_detection_rows_killed`. To minimize conflicts: design your application to route writes for a specific entity (user, tenant) to a specific primary node, creating logical partitioning even in a multi-primary cluster. This transforms a conflict-prone multi-primary scenario into effectively several single-primary zones — minimal conflicts, maximum throughput.

---

### 233. What is eventual consistency?

"**Eventual consistency** means that if no new writes occur, all replicas will eventually converge to the same value — but at any given moment, different replicas may return different values for the same key.

In MySQL async replication: the primary commits a write. Replicas apply it milliseconds to seconds later. During the replication delay, reading from a replica returns stale data. This is eventual consistency.

Applications designed for eventual consistency must tolerate stale reads. Use cases: social media timelines, product catalog reads, read-your-writes (route own reads to primary)."

#### Indepth
Eventual consistency isn't a binary — it has **degrees** quantified by **RPO (Recovery Point Objective)**: how stale is the stale data? Async MySQL replication typically lags 10ms–2 seconds under normal load (excellent eventual consistency). During heavy write spikes or replica maintenance, lag can reach minutes (poor, potentially disruptive eventual consistency). Applications should be designed with explicit **read-your-writes consistency** for critical user paths (a user who just updated their profile should always see the update) while tolerating stale data for non-critical paths (leaderboard rankings, recommendation counts).

---

### 234. What is the split-brain problem?

"**Split-brain** occurs in a clustered system when a **network partition** causes nodes to lose contact with each other, and each partition believes the other has failed — both becoming primary simultaneously.

With two primary MySQL servers accepting writes simultaneously: different data is written to each. When the partition heals, both primaries have divergent histories that can't be automatically merged (unlike CRDTs). This is **data corruption**.

Prevention: **quorum requirement** (Paxos/Raft). A node becomes primary only if it can communicate with the majority of cluster nodes. A 3-node cluster requires 2 nodes; if a node is partitioned alone (1 node), it cannot get quorum and refuses to become primary."

#### Indepth
InnoDB Cluster (Group Replication) prevents split-brain via Paxos quorum: a group of 3 accepts writes only when 2+ nodes can communicate. If the network splits into {1 node} and {2 nodes}, only the majority partition continues accepting writes. The single-node partition freezes. When connectivity restores, the single node rejoins as a replica and catches up from the majority partition's binlog. **Orchestrator** (used without Group Replication) prevents split-brain with `--recoveryIgnoreHostnameFilters` and anti-fencing mechanisms — it refuses to promote a replica if it can't confirm the current primary is truly unavailable.

---

### 235. How do you secure MySQL against SQL injection?

"Defense-in-depth against SQL injection:

1. **Primary defense — Parameterized queries / Prepared statements**: Never concatenate user input into SQL. Use `?` placeholders and let the driver handle escaping.

```go
db.Query("SELECT * FROM users WHERE email = ?", userInput)
```

2. **Least privilege**: App DB user has only `SELECT, INSERT, UPDATE, DELETE` — not `DROP`, `CREATE`, `FILE`.
3. **Input validation**: Sanitize and whitelist input types.
4. **WAF (Web Application Firewall)**: Block common SQLi patterns at the network layer.
5. **ProxySQL Firewall**: Block specific query patterns at the database proxy."

#### Indepth
SQL injection prevention is **100% application-layer** — the database itself can't detect injected SQL because by the time MySQL receives the query, it's just SQL. The only foolproof protection is parameterized queries — they separate the query structure from the data, making it structurally impossible for data to modify the query logic. ORMs (GORM, Hibernate, SQLAlchemy) use parameterized queries by default but provide raw query escape hatches that are dangerous (`db.Raw()`, `db.Exec()`). Audit every use of raw query execution in codebases for injection vulnerabilities.

---

### 236. What is data masking?

"**Data masking** is the process of replacing sensitive data (PII, credit card numbers, SSNs) with realistic but non-sensitive substitutes, for use in non-production environments (development, testing, staging).

MySQL Enterprise Edition includes a **Data Masking and De-identification** plugin with functions:
```sql
SELECT mask_ssn('123-45-6789');  -- Returns: 'XXX-XX-6789'
SELECT mask_pan('4111111111111111');  -- Returns: 'XXXXXXXXXXXX1111'
```

I use data masking to let developers work with production-scale data without exposing real customer information."

#### Indepth
Data masking is a **GDPR/compliance requirement** for any organization handling PII. Giving developers actual production data (names, emails, phone numbers) violates GDPR's data minimization principle. A proper data masking pipeline: (1) Extract production backup. (2) Apply masking transformations (swap real PII with synthetic equivalents using consistent hashing so relationships remain valid). (3) Import masked dataset into dev/staging. Open-source tools: **anonymizer** (Go), **PostgreSQL Anonymizer** (works similarly for MySQL via scripts). Key requirement: **referential consistency** — masked user_id must be the same masked value everywhere it appears.

---

### 237. What is audit logging?

"**Audit logging** records who did what to the database — every login, query execution, privilege change, and data access — for security compliance and forensic investigation.

MySQL Enterprise Edition: **MySQL Enterprise Audit** plugin logs to XML or JSON format.

Open-source alternatives: **McAfee MySQL Audit Plugin**, **Percona Audit Log Plugin** (included in Percona Server).

```sql
-- Enable audit (Percona):
INSTALL PLUGIN audit_log SONAME 'audit_log.so';
SET GLOBAL audit_log_policy = ALL;  -- Log everything
```

I enable audit logging for all production databases handling PII or financial data."

#### Indepth
Audit log configuration requires careful policy design — logging ALL queries on a high-throughput database generates massive log volumes (GBs/hour). Best practice: (1) Log all privileged operations (DDL, GRANT, DROP, etc.) always. (2) For DML on sensitive tables, use **table-level filtering**: only audit `SELECT` on `users`, `payments` tables. (3) Audit all failed authentications. (4) Ship logs to a central SIEM (Splunk, ELK) for retention and anomaly detection. Never store audit logs on the same server — a compromised MySQL server could delete them. Ship to write-only remote storage immediately.

---

### 238. What is GDPR compliance in databases?

"GDPR (General Data Protection Regulation) requires databases handling EU citizens' data to implement:

1. **Data minimization**: Only store data that's necessary.
2. **Right to erasure (Right to be Forgotten)**: Be able to delete all data related to a specific person on request.
3. **Data portability**: Export all data for a user in a machine-readable format.
4. **Consent management**: Track and store consent for data processing.
5. **Pseudonymization / Encryption**: Protect PII at rest and in transit.
6. **Audit trails**: Log who accessed PII data.
7. **Data breach notification**: Detect and report breaches within 72 hours."

#### Indepth
The most technically complex GDPR requirement for MySQL: **Right to Erasure**. Simply deleting a user from the `users` table doesn't erase their data if user_id references appear in dozens of other tables (orders, reviews, logs, backups). A proper erasure must scrub all PII while preserving referential integrity (e.g., keep the order record but replace user PII with anonymized placeholders). Build an **erasure choreography service** that knows every table containing PII and executes the correct anonymization/deletion. Test regularly — GDPR erasure failures carry fines up to 4% of global annual turnover.

---

### 239. How do you archive historical data?

"Data archiving moves old, rarely-accessed data from hot production tables to cheaper, slower cold storage while keeping the production table small and fast.

Strategies:
1. **Partition-based archiving**: Drop an old partition (`ALTER TABLE t DROP PARTITION p2020`) — instant, no row-by-row deletion.
2. **Copy-then-delete**: Insert old rows into an archive table, then delete from hot table. Use `pt-archiver` for online, non-blocking archival.
3. **Cold storage export**: Export old data to S3 as Parquet/CSV files. Query with Athena/BigQuery when needed.
4. **Separate archive database**: Keep a read-only archive MySQL instance with pre-configured retention.

I archive based on business-defined retention windows (e.g., orders > 3 years old)."

#### Indepth
**Partition-based archiving** is the most efficient: dropping a partition is a metadata operation — instant regardless of partition size. The prerequisite: the table must be partitioned by the archival key (usually a date). Design partitioning at schema creation time with archival in mind. **`pt-archiver`** is the tool for non-partitioned tables: it deletes rows in configurable-size batches during low-traffic windows, rate-limiting to avoid impacting production. It optionally writes archived rows to a destination table or file before deletion. Combined with a cron job, it creates a fully automated rolling-window archive.

---

### 240. How do you handle 10TB+ databases?

"For 10TB+ MySQL databases I apply multiple layers:

1. **Vertical scaling first**: Use the largest available instance. 10TB on a 192-core / 3TB RAM / high-IOPS NVMe machine — the working set often fits in memory.
2. **InnoDB compression**: `COMPRESSION='zlib'` on large, cold text-heavy tables can halve storage.
3. **Partitioning + archival**: Keep hot partitions (< 1 year) in production, archive older partitions to cheaper storage.
4. **Read replicas**: 10+ replicas for read traffic distribution. Hot read replicas can use SSD, warm ones can use HDD.
5. **Sharding**: If write throughput or working set can't fit on one server, shard by a natural key (tenant, region).
6. **Transparent sharding**: Vitess manages shard topology, allowing gradual resharding without app changes."

#### Indepth
At 10TB+, **backup and recovery duration** becomes a critical concern. A 10TB mysqldump restore takes 20+ hours. Percona XtraBackup physical restore is faster (~2–3 hours for file copy) but still significant. Strategies: (1) **Incremental backups**: full weekly + daily incrementals (XtraBackup has native incremental support). (2) **Multiple backup replicas**: maintain a replica whose sole purpose is backup — never promoted, never serving reads. (3) **Object storage streaming**: stream XtraBackup output directly to S3/GCS without local temp storage. (4) **Automated restore testing**: monthly drill to restore the latest backup and verify data integrity.

---

### 241. How do you design a multi-tenant database architecture?

"Three approaches to multi-tenant MySQL architecture:

1. **Shared schema (single database, all tenants)**: Simplest. Add `tenant_id` to every table. All tenants share the same tables. Easy to scale, hard to isolate.

2. **Separate schema per tenant (same server)**: Each tenant gets their own schema/database on the same MySQL server. Good isolation, simpler than separate servers. MySQL's schema-per-tenant model limits to ~hundreds of tenants practical.

3. **Separate server per tenant**: Maximum isolation. Each tenant has a dedicated MySQL instance. Expensive, best for enterprise/SaaS with strict data isolation requirements.

I use shared schema for SMB SaaS, separate servers for enterprise/regulated tenants."

#### Indepth
Shared schema is the most commonly used approach because it achieves economies of scale. Key design rule: **every single table must have `tenant_id` as the first column of its composite primary key or as the leading column of all indexes**. This ensures all queries are tenant-scoped and partition pruning works. The biggest risk: a missing `WHERE tenant_id = ?` in a query accidentally exposes one tenant's data to another — **data leakage**. Enforce tenant scoping at the ORM level (always inject tenant context into every query) and use Row Level Security if available (PostgreSQL does this natively; MySQL requires application enforcement).

---

### 242. How do you reduce write amplification?

"**Write amplification** occurs when one logical write causes multiple physical writes — to the redo log, the actual data page, the secondary index pages, the binary log, and potentially, the doublewrite buffer.

Strategies to reduce:

1. **Fewer secondary indexes**: Each secondary index multiplies physical writes per row change. Only index what query patterns require.
2. **`innodb_change_buffer_max_size`**: Increase for write-heavy batch workloads — buffers secondary index changes.
3. **Batch writes**: One `INSERT ... VALUES (a), (b), (c)` generates less WAL and index overhead than three separate INSERTs.
4. **Large redo log files**: Fewer, larger checkpoints reduce cumulative write I/O.
5. **`binlog_row_image=MINIMAL`**: Log only changed columns, not full row images."

#### Indepth
The theoretical minimum write amplification for InnoDB is ~3: redo log + doublewrite buffer + actual data page. In practice, for a row with 5 secondary indexes, it's 3 + 5 secondary index updates = 8× amplification per write. SSDs suffer more from write amplification than HDDs because of internal flash write amplification (cell-level), compounding MySQL's application-level amplification. Using `ALGORITHM=INPLACE` for schema changes avoids full table rebuilds (which would amplify writes for every row). Write amplification is the primary reason NVMe SSDs are preferred for high-write MySQL — their high endurance (TBW rating) tolerates sustained application + hardware write amplification.

---

### 243. How do you design for high write workloads?

"Design patterns for high-write MySQL systems:

1. **Batching**: Accumulate writes in the application and perform bulk inserts instead of individual row inserts. `INSERT INTO t VALUES (a), (b), (c), ...` is orders of magnitude faster.
2. **Async writes via queue**: Application writes to Kafka/Redis; a consumer writes to MySQL. Decouples write spikes from the database.
3. **Write-optimized schema**: Minimize secondary indexes on hot-write tables. Add them on a read replica instead.
4. **InnoDB tuning**: Large buffer pool, large redo log, `innodb_flush_log_at_trx_commit=2`, `innodb_io_capacity=high`.
5. **Partitioning**: Spread writes across multiple partitions to reduce B+Tree contention on the latest leaf page."

#### Indepth
The **write-queue pattern** (application → Kafka → MySQL) is the most impactful for extreme write throughput. A consumer can batch 1000 Kafka messages into a single MySQL bulk insert, achieving near-disk write throughput. The tradeoff: eventual consistency (reads don't see writes until the consumer flushes). For systems requiring sync write-then-read consistency, use **dual writes** (write to both MySQL and a fast write buffer like Redis) with MySQL used for durable storage and Redis for low-latency reads until MySQL catches up. This is the CQRS (Command Query Responsibility Segregation) pattern applied to MySQL.

---

### 244. What are the limitations of MySQL?

"MySQL's key limitations compared to other databases:

1. **Weak optimizer**: Less sophisticated than PostgreSQL's optimizer for complex queries (poorly handles some JOIN orderings, subqueries).
2. **No transactional DDL**: DDL causes implicit commit — can't roll back a failed schema migration.
3. **Single-primary writes**: Scaling writes beyond one server requires sharding (vs Spanner/CockroachDB which natively scale writes).
4. **Limited data types**: No native UUID type, limited array/composite type support.
5. **Subpar JSON querying**: JSON storage and querying is weaker than PostgreSQL's JSONB.
6. **No partial indexes**: Can't index only rows WHERE condition IS TRUE.
7. **Expensive COUNT(*)**: No materialized row count for InnoDB tables.
8. **FULLTEXT is basic**: Not suitable for production-grade search at scale."

#### Indepth
The single-primary limitation is MySQL's most significant architectural constraint. PostgreSQL is in the same boat for standard OLTP, but PostgreSQL's logical replication and ecosystem tools for multi-master are more mature. MySQL's alternative is Vitess (transparent sharding + online schema change) which is battle-tested at YouTube/PlanetScale scale. For truly globally-distributed consistent writes, MySQL is the wrong choice — Spanner, CockroachDB, or TiDB (MySQL-compatible distributed SQL) are architecturally superior. The choice to use MySQL should reflect the team's operational expertise and confidence in vertical scaling + read replica scaling for the foreseeable future.

---

### 245. When should you NOT use MySQL?

"Situations where MySQL is a poor choice:

1. **Globally distributed, low-latency writes**: Cross-region write latency kills OLTP performance. Use Spanner, CockroachDB.
2. **Complex analytical queries on massive datasets** (OLAP): MySQL is an OLTP database. Use ClickHouse, BigQuery, Redshift.
3. **Graph data with deep relationship traversal**: Neo4j or Dgraph outperform MySQL for multi-hop graph queries.
4. **Unstructured, schema-free document storage**: MongoDB or CouchDB are more natural. (Though MySQL's JSON type helps.)
5. **Time-series telemetry at extreme scale** (millions of inserts/sec): TimescaleDB, InfluxDB, or OpenTSDB.
6. **Full-text search at production quality**: Elasticsearch or Meilisearch dramatically outperform MySQL FULLTEXT."

#### Indepth
The most common mistake: using MySQL as an analytics database for complex aggregations over hundreds of millions of rows. MySQL's row-store architecture (InnoDB reads full rows, not columns) is extremely inefficient for analytical aggregations that touch only a few columns of a wide table. **ClickHouse**, a columnar database, can be 100–1000× faster for the same analytical query because it only reads the columns involved in the query. The correct pattern for most teams: MySQL for OLTP (transactional operations), ClickHouse/BigQuery for OLAP (analytics) — sync data using CDC (Debezium or MySQL binlog streaming).

---

### 246. What are the differences between MySQL and PostgreSQL?

"Key MySQL vs PostgreSQL differences:

| Feature | MySQL | PostgreSQL |
|---|---|---|
| License | GPL / Commercial | PostgreSQL License (fully open) |
| DDL transactions | ❌ No (implicit commit) | ✅ Yes |
| JSON support | Basic JSON type | Advanced JSONB (binary, indexed, operators) |
| Partial indexes | ❌ No | ✅ Yes |
| Table inheritance | ❌ No | ✅ Yes |
| Parallel queries | Limited | ✅ Full parallel query |
| Full-text search | Basic | Advanced (with weights, headlines, ranking) |
| Extensions | Limited | Rich (PostGIS, pgvector, Timescale) |
| Ecosystem/tooling | Excellent (ProxySQL, Vitess, XtraBackup) | Growing fast |

MySQL excels at: simple OLTP, wide tooling, massive operational knowledge base."

#### Indepth
The most critical practical difference for teams: **transactional DDL in PostgreSQL**. In PostgreSQL, an `ALTER TABLE` can be inside a `BEGIN ... COMMIT` block and rolled back on failure — zero-risk schema migrations. In MySQL, this is impossible. This single difference makes PostgreSQL's operational story for schema evolution significantly safer. MySQL's advantage: a much larger pool of experienced DBAs and battle-tested tooling (Vitess, Orchestrator, ProxySQL, XtraBackup) with decades of production hardening at massive scale that PostgreSQL equivalents are still catching up to.

---

### 247. How does MySQL compare with MongoDB?

"MySQL (relational) vs MongoDB (document):

| Feature | MySQL | MongoDB |
|---|---|---|
| Schema | Fixed, strict (DDL) | Flexible, schema-less |
| ACID transactions | ✅ Full (single + multi-document) | ✅ Full (since 4.0) |
| Query language | SQL (powerful joins, aggregates) | MQL (powerful, but no joins) |
| Horizontal scaling | Shard manually or via Vitess | Native built-in sharding |
| Indexing | B+Tree, fulltext, spatial | B+Tree, text, geo, partial |
| Best for | Structured relational data | Hierarchical documents, varied schema |

I choose MySQL when data has clear relationships and schema is stable. MongoDB when schema is unpredictable and nested document model fits naturally."

#### Indepth
The MongoDB vs MySQL choice is often ideological rather than technical. Modern MongoDB supports ACID transactions (though with performance overhead), and modern MySQL supports JSON types (though less elegantly than MongoDB). The real differentiation: **developer experience for different data shapes**. An order with line items, shipping address, and payment details in MongoDB is one document — natural to query. In MySQL, it's 4–5 tables with JOINs. Conversely, financial ledgers in MongoDB lack the relational enforcement and JOIN power MySQL provides. Pick based on your **primary entity structure**, not technology hype.

---

### 248. How do you integrate MySQL with caching systems (Redis)?

"MySQL + Redis integration patterns:

1. **Cache-aside (lazy loading)**: App checks Redis first. On miss, reads from MySQL, stores result in Redis with TTL.

2. **Write-through**: App writes to both Redis and MySQL on every write. Cache always fresh but adds write latency.

3. **Read-replica caching**: Direct hot read queries to Redis; write queries and cache misses to MySQL.

4. **Session caching**: Store user sessions in Redis (fast TTL management) while persisting durable data in MySQL.

5. **Counter caching**: Increment counters in Redis (atomic `INCR`); periodically flush totals to MySQL."

#### Indepth
Cache invalidation is the hardest problem: when MySQL data changes, the Redis cache for that data must be invalidated or updated. Strategies: (1) **TTL-based expiry**: cache items expire after N seconds. Simple but allows stale reads for up to N seconds. (2) **Write-invalidation**: explicitly DELETE the cache key on every MySQL write. Application must ensure consistency. (3) **Change Data Capture (CDC)**: use Debezium to stream MySQL binlog changes into a Kafka topic, consume it to invalidate specific Redis keys. CDC provides an event-driven, decoupled invalidation system without adding cache invalidation to every write path.

---

### 249. How do you design a disaster recovery plan?

"A comprehensive MySQL Disaster Recovery (DR) plan includes:

**RPO (Recovery Point Objective)**: Define max acceptable data loss. 0 = no loss (semi-sync), 5min = async replication.

**RTO (Recovery Time Objective)**: Define max acceptable downtime. 30s = automated failover. 4h = restore from backup.

**Components**:
1. Regular automated backups (XtraBackup daily, incremental hourly) stored off-site.
2. Binary log shipping to a separate DR region in real-time.
3. Standby replica in DR region (async or semi-sync from primary).
4. Automated failover tooling (Orchestrator, InnoDB Cluster).
5. Documented runbooks for each failure scenario.
6. Regular DR drills — test restoration monthly."

#### Indepth
The most overlooked DR element: **restoration testing**. Backups that have never been restored are worthless in an emergency — you don't know if they're valid until you try. Automate monthly or weekly restore tests: take the latest backup, restore it to a test instance, run integrity checks (`mysqlcheck --all-databases`), verify row counts match production estimates. Also test PITR: restore backup + replay N hours of binlog and verify the target state. DR plans that exist only on paper fail in real incidents — the team needs muscle memory from drills to execute under pressure.

---

### 250. What are best practices for production MySQL?

"My production MySQL best practices checklist:

**Schema Design**:
- Use InnoDB always. Define proper primary keys. Use `utf8mb4`.
- Set explicit NOT NULL and DEFAULT on all columns.

**Performance**:
- Right-size buffer pool (70–80% RAM). Use covering indexes for hot query paths.
- Enable slow query log. Run `pt-query-digest` weekly.

**Reliability**:
- `innodb_flush_log_at_trx_commit=1`, `sync_binlog=1`. Enable GTID replication.
- 2+ replicas minimum. Automated failover (Orchestrator/InnoDB Cluster).
- Daily backups + continuous binlog archiving. Monthly restore tests.

**Security**:
- Least-privilege users. SSL/TLS for all connections. No root login remotely.
- Run `mysql_secure_installation` on every fresh install.

**Monitoring**:
- Alert on: replication lag, connection count, buffer pool hit rate, slow query rate."

#### Indepth
The most-violated production best practice: teams disable `sync_binlog=1` and `innodb_flush_log_at_trx_commit=1` for performance without understanding the risk. A single unexpected OS crash or power loss with these disabled can result in committed transactions being permanently lost (no recovery possible). The performance improvement from disabling these is real (up to 2× write throughput) but the risk is catastrophic data loss in a failure scenario that WILL eventually happen. Use NVMe SSDs with DRAM write cache and BBU to get the performance of async flushing with the safety of synchronous durability instead of compromising durability.

---
