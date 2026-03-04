# 🏗️ 04 — MySQL Replication & Architecture
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

> System design and replication — deeply probed in **architecture rounds at Amazon, Ola, Zomato, Nykaa, BYJU'S**.

---

## 🔑 Must-Know Topics
- Asynchronous replication (master–slave / primary–replica)
- Semi-synchronous replication
- GTID-based replication
- Binary log (binlog) formats: STATEMENT, ROW, MIXED
- Sharding strategies: range, hash, directory-based
- Vertical vs horizontal scaling
- Table partitioning: RANGE, LIST, HASH, KEY
- MySQL Group Replication & InnoDB Cluster
- ProxySQL / MySQL Router for connection routing

---

## ❓ Most Asked Questions

### Q1. How does MySQL replication work?

```sql
-- MySQL Replication: Primary (source) → one or more Replicas

-- HOW IT WORKS:
-- 1. Every write on Primary is recorded in the Binary Log (binlog)
-- 2. Replica's IO Thread connects to Primary, downloads binlog events → stores in Relay Log
-- 3. Replica's SQL Thread reads Relay Log, re-executes the events locally
-- This is ASYNCHRONOUS by default (Primary doesn't wait for replica confirmation)

-- Set up primary (my.cnf):
-- [mysqld]
-- server-id = 1
-- log-bin = mysql-bin
-- binlog_format = ROW  -- safest format for replication

-- Set up replica (my.cnf):
-- [mysqld]
-- server-id = 2        -- must be unique across all servers

-- On Primary:
CREATE USER 'replication_user'@'%' IDENTIFIED BY 'SecurePassword!';
GRANT REPLICATION SLAVE ON *.* TO 'replication_user'@'%';
FLUSH TABLES WITH READ LOCK;
SHOW MASTER STATUS;   -- note File and Position
-- Take snapshot (mysqldump --single-transaction), then:
UNLOCK TABLES;

-- On Replica:
CHANGE MASTER TO
    MASTER_HOST = '10.0.0.1',
    MASTER_USER = 'replication_user',
    MASTER_PASSWORD = 'SecurePassword!',
    MASTER_LOG_FILE = 'mysql-bin.000001',
    MASTER_LOG_POS = 154;
START SLAVE;
SHOW SLAVE STATUS\G  -- check Seconds_Behind_Master, Slave_IO_Running, Slave_SQL_Running
```

---

### Q2. What is GTID-based replication? Why is it preferred?

```sql
-- GTID = Global Transaction Identifier
-- Format: source_server_uuid:transaction_number
-- e.g., 3E11FA47-71CA-11E1-9E33-C80AA9429562:1-100

-- Benefits over position-based replication:
-- ✅ Each transaction has universally unique ID — no need to track binlog file+position
-- ✅ Failover is simple — replica just needs the GTID to resume from
-- ✅ Easier to set up replicas without knowing exact binlog position
-- ✅ MySQL auto-skips already-applied transactions (idempotent)

-- Enable GTID (my.cnf):
-- gtid_mode = ON
-- enforce_gtid_consistency = ON

-- Set up replica with GTID:
CHANGE MASTER TO
    MASTER_HOST = '10.0.0.1',
    MASTER_USER = 'replication_user',
    MASTER_PASSWORD = 'SecurePassword!',
    MASTER_AUTO_POSITION = 1;  -- 🎯 GTID: no file/position needed!
START SLAVE;

-- Check GTID state
SELECT @@gtid_executed;    -- GTIDs already applied on this server
SELECT @@gtid_purged;      -- GTIDs purged from binlog (history)
SHOW MASTER STATUS;        -- shows Executed_Gtid_Set
```

---

### Q3. What are binlog formats? What is the difference?

```sql
-- Binary log records all changes for replication and point-in-time recovery

-- 3 formats (controlled by binlog_format):

-- 1. STATEMENT — logs the SQL statement itself
--    Pros: compact, human-readable
--    Cons: non-deterministic statements (NOW(), RAND(), UUID()) can cause inconsistency

-- 2. ROW (recommended) — logs actual row changes (before/after values)
--    Pros: accurate, safe for all statements including non-deterministic
--    Cons: larger binlog size (one record per changed row)
--    Best for: production replication

-- 3. MIXED — uses STATEMENT normally, switches to ROW for non-deterministic queries
--    Pros: balance of compactness and safety
--    Cons: unpredictable size, harder to debug

SET GLOBAL binlog_format = 'ROW';

-- Use binlog for point-in-time recovery:
-- mysqlbinlog --start-datetime="2024-01-15 10:00:00"
--             --stop-datetime="2024-01-15 11:00:00"
--             /var/lib/mysql/mysql-bin.000042 | mysql -u root -p

-- Inspect binlog contents:
-- mysqlbinlog /var/lib/mysql/mysql-bin.000001 | head -100
SHOW BINLOG EVENTS IN 'mysql-bin.000001' LIMIT 20;
```

---

### Q4. What is semi-synchronous replication?

```sql
-- Standard replication: ASYNCHRONOUS
-- Primary commits → immediately returns success → replica catches up eventually
-- Risk: if Primary crashes before replica gets the event → DATA LOSS

-- Semi-synchronous replication:
-- Primary commits → WAITS for at least 1 replica to ACK it wrote to relay log
-- Then returns success to client
-- Risk window: still small (network round-trip), but much safer than async

-- Enable semi-sync (requires plugin):
INSTALL PLUGIN rpl_semi_sync_master SONAME 'semisync_master.so';  -- on Primary
INSTALL PLUGIN rpl_semi_sync_slave  SONAME 'semisync_slave.so';   -- on Replica

SET GLOBAL rpl_semi_sync_master_enabled = 1;
SET GLOBAL rpl_semi_sync_slave_enabled  = 1;

-- Timeout: if no replica ACKs within timeout, falls back to async
SET GLOBAL rpl_semi_sync_master_timeout = 10000;  -- milliseconds (10 seconds)

-- Check semi-sync status:
SHOW STATUS LIKE 'Rpl_semi_sync%';

-- Lossless semi-sync (MySQL 5.7+): Primary waits for replica ACK BEFORE committing
-- Zero RPO (Recovery Point Objective) — no data loss even on Primary crash
```

---

### Q5. How do you handle read/write scaling with replication?

```sql
-- Architecture pattern: Write to Primary, Read from Replica(s)
-- Application or proxy layer routes:
--   INSERT/UPDATE/DELETE → Primary
--   SELECT              → Replica (for non-critical reads)

-- ProxySQL configuration example (conceptual):
-- Frontend: applications connect to ProxySQL (127.0.0.1:6033)
-- Backend: ProxySQL routes based on rules

-- Rule: reads to group 20 (replicas), writes to group 10 (primary)
-- INSERT INTO mysql_query_rules: match_pattern='^SELECT' → destination=group 20

-- Replication lag problem:
-- If app writes → Primary, then immediately reads → Replica, it may NOT see the write yet!
-- Solutions:
-- 1. Read your own writes: route specific users' reads to primary after writes
-- 2. Session variables: after write, force reads to primary: SET SESSION read_from_master=1
-- 3. Read propagation wait: SELECT WAIT_FOR_EXECUTED_GTID_SET('gtid:1-N', timeout)
-- 4. Cache: write to cache immediately after DB write; read from cache
SHOW SLAVE STATUS\G  -- check Seconds_Behind_Master — acceptable lag for your use case?
```

---

### Q6. What is database sharding? What strategies exist?

```sql
-- Sharding: horizontally split data across multiple MySQL servers (shards)
-- Each shard has a subset of rows — enables horizontal scaling beyond one server

-- Strategy 1: RANGE sharding
-- Shard 1: user_id 1      – 1,000,000
-- Shard 2: user_id 1,000,001 – 2,000,000
-- Pros: simple, range queries stay on one shard
-- Cons: uneven distribution ("hot shard" for new users), rebalancing is hard

-- Strategy 2: HASH sharding
-- shard_number = hash(user_id) % num_shards
-- Shard 0: users where user_id % 4 = 0
-- Shard 1: users where user_id % 4 = 1 ...
-- Pros: even distribution
-- Cons: range queries hit ALL shards; changing num_shards requires full reshard

-- Strategy 3: DIRECTORY / Lookup sharding
-- Lookup table maps entity → shard: user 12345 → shard 3
-- Pros: flexible, can move entities between shards
-- Cons: lookup table is a bottleneck / single point of failure

-- Strategy 4: Functional sharding (by feature)
-- Users DB | Orders DB | Products DB | Analytics DB
-- Not technically row-level sharding — separates by domain

-- Application-level sharding:
-- Application code computes target shard, maintains N connection pools
-- Vitess (YouTube), Citus (PostgreSQL), PlanetScale — handle sharding transparently
```

---

### Q7. What is table partitioning in MySQL?

```sql
-- Partitioning: divide ONE table into physically separate files on disk
-- Improves query performance via "partition pruning" (only scan relevant partition)
-- Transparent to application — still queries one logical table

-- RANGE partitioning by year
CREATE TABLE orders (
    id         INT NOT NULL,
    order_date DATE NOT NULL,
    total      DECIMAL(10,2),
    PRIMARY KEY (id, order_date)    -- partition key MUST be part of PK
) PARTITION BY RANGE (YEAR(order_date)) (
    PARTITION p2021 VALUES LESS THAN (2022),
    PARTITION p2022 VALUES LESS THAN (2023),
    PARTITION p2023 VALUES LESS THAN (2024),
    PARTITION p2024 VALUES LESS THAN (2025),
    PARTITION pmax  VALUES LESS THAN MAXVALUE
);

-- Query only hits p2023 partition:
SELECT * FROM orders WHERE order_date BETWEEN '2023-01-01' AND '2023-12-31';
EXPLAIN PARTITIONS SELECT * FROM orders WHERE YEAR(order_date) = 2023;
-- partitions: p2023 ← partition pruning!

-- LIST partitioning by region
CREATE TABLE sales (
    id     INT,
    region VARCHAR(10)
) PARTITION BY LIST COLUMNS(region) (
    PARTITION p_us  VALUES IN ('US', 'CA', 'MX'),
    PARTITION p_eu  VALUES IN ('DE', 'FR', 'GB'),
    PARTITION p_apac VALUES IN ('IN', 'JP', 'SG')
);

-- Maintenance operations on partitions:
ALTER TABLE orders TRUNCATE PARTITION p2021;     -- fast delete old data
ALTER TABLE orders DROP PARTITION p2021;          -- remove partition + data
ALTER TABLE orders ADD PARTITION (PARTITION p2025 VALUES LESS THAN (2026));
```

---

### Q8. What is MySQL Group Replication / InnoDB Cluster?

```sql
-- MySQL Group Replication: multi-primary or single-primary cluster
-- Transactions are certified across all members before commit
-- Automatic failover — if primary fails, new primary elected automatically
-- Uses Paxos-based consensus protocol (requires 3+ nodes for quorum)

-- Single-primary mode (default): one writer, others replicate to replicas
-- Multi-primary mode: all nodes can accept writes (conflict resolution built-in)

-- Setup overview:
-- 1. Configure group_replication plugin on all 3 nodes
-- 2. All nodes have same server-id range, GTIDs enabled
-- plugin-load-add = group_replication.so
-- group_replication_group_name = 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'
-- group_replication_start_on_boot = OFF
-- group_replication_local_address = 'node1:33061'
-- group_replication_group_seeds = 'node1:33061,node2:33061,node3:33061'

-- Start replication:
-- SET GLOBAL group_replication_bootstrap_group=ON;  -- first node only!
-- START GROUP_REPLICATION;
-- SET GLOBAL group_replication_bootstrap_group=OFF;

-- Check group membership:
SELECT * FROM performance_schema.replication_group_members;

-- InnoDB Cluster = Group Replication + MySQL Shell + MySQL Router
-- MySQL Router: reads → any member, writes → primary (automatic routing)
-- AdminAPI: management via MySQL Shell (cluster.addInstance, cluster.status())
```

---

### Q9. What are read replicas and how do you monitor replication lag?

```sql
-- Replication lag: how far behind a replica is from the primary
-- Critical metric — a lagging replica serves STALE data

-- Monitor lag (on replica):
SHOW SLAVE STATUS\G
-- Key columns:
-- Seconds_Behind_Master: lag in seconds (NULL = replica is stopped or disconnected)
-- Slave_IO_Running: Yes/No (is IO thread running?)
-- Slave_SQL_Running: Yes/No (is SQL thread running?)
-- Last_Error: shows last replication error

-- Performance Schema monitoring:
SELECT * FROM performance_schema.replication_connection_status;
SELECT * FROM performance_schema.replication_applier_status_by_worker;

-- Causes of replication lag:
-- 1. Heavy writes on primary (LOT of data to replicate)
-- 2. Single-threaded SQL applier (fix: parallel replication)
-- 3. Large transactions (long-running UPDATE ... WHERE no index → slow on replica)
-- 4. Network latency between primary and replica

-- Enable parallel replication (MySQL 5.7+):
SET GLOBAL slave_parallel_workers = 4;        -- 4 SQL threads
SET GLOBAL slave_parallel_type = 'LOGICAL_CLOCK';  -- safest parallel mode

-- Alert when lag exceeds threshold:
-- Check Seconds_Behind_Master > 30 → trigger PagerDuty/Slack alert
```

