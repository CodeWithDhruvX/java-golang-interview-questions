# 🗣️ Theory — MySQL Replication & Architecture
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does MySQL replication work?"

> *"MySQL replication works on the publish-subscribe model using the binary log. Every write operation on the primary server — INSERT, UPDATE, DELETE, DDL — gets recorded into the binary log in sequence. The replica has an IO thread that connects to the primary, reads those binlog events continuously, and writes them into a local relay log on the replica. A separate SQL thread then reads from the relay log and replays each event on the replica. By default, this is asynchronous — the primary writes to its binlog and immediately confirms the commit to the client, without waiting for any replica. The replica catches up as fast as it can, but there's always inherent lag. This design gives high write throughput on the primary at the cost of potential data loss if the primary crashes before the replica has caught up. The common architecture is one or two read replicas handling all SELECT traffic while the primary handles all writes."*

---

## Q: "What is GTID and why is it better than position-based replication?"

> *"GTID stands for Global Transaction Identifier. Every transaction committed on any server in a replication topology gets a universally unique ID — the server's UUID plus a sequence number, like a UUID colon 1 through N. With position-based replication, when you need to add a new replica or promote a replica to primary after a failover, you have to know the exact binary log file name and byte offset position where the new replica should start reading. Getting this wrong means data loss or duplicate data. With GTIDs, a replica just tells the primary 'I have already executed these GTIDs, send me everything else.' The primary figures out the right starting point automatically. Failed automatic failover scenarios that used to require DBA intervention become automatable. MySQL Group Replication and orchestration tools like Orchestrator or MHA work much more reliably with GTIDs. The only setup cost is enabling gtid_mode and enforce_gtid_consistency in my.cnf and ensuring no non-GTID-safe operations are executed."*

---

## Q: "What are the three binlog formats? Which do you recommend?"

> *"STATEMENT format logs the SQL statement itself. It's compact and human-readable — but it breaks down for non-deterministic statements. If you execute DELETE FROM logs ORDER BY RAND() LIMIT 100, the primary deletes a random 100 rows. The replica re-executes the same SQL but gets a different random 100 rows. Now your replica has different data than your primary — replication is broken and you might not even notice. ROW format logs the actual row changes — the before and after values for each affected row. It's accurate for every possible SQL statement including non-deterministic ones, and it's the recommended format for production. The downside is larger binlog size — a single UPDATE affecting 1 million rows generates 1 million row change records. MIXED format uses STATEMENT normally for efficiency, and automatically switches to ROW when it detects non-deterministic statements. It's a reasonable compromise but harder to reason about. My recommendation: always use ROW in production — storage is cheap, correctness is not."*

---

## Q: "How do you approach database sharding? What are the tradeoffs?"

> *"Sharding is horizontal partitioning — splitting your data across multiple database servers so the total capacity and write throughput grow with the number of shards. It's a last resort when a single powerful MySQL server and its replicas can no longer handle the write load or data volume. The main sharding strategies: range-based sharding splits by value ranges — user IDs 1 to 1 million on shard 1, 1 to 2 million on shard 2. It's simple and range queries stay on one shard, but new data accumulates on the last shard — you get hot spots. Hash-based sharding routes rows by computing a hash of the key modulo the number of shards — more even distribution, but adding shards requires rehashing and moving data. Directory-based sharding uses a lookup table to map entities to shards — the most flexible, but the lookup table itself becomes a single point of failure. The hard truths about sharding: cross-shard JOINs are painful or impossible, cross-shard transactions require distributed transaction protocols, and sharding significantly complicates your application code. I'd exhaust vertical scaling, read replicas, caching, and partitioning before accepting the operational complexity of sharding."*

---

## Q: "What is table partitioning? How is it different from sharding?"

> *"Table partitioning splits a single logical table into physically separate storage segments — but it all still lives on one server, one MySQL instance. The application sees one table and queries it normally. MySQL internally looks at the partition key in the WHERE clause and only reads the relevant partitions — this is called partition pruning. So a query like WHERE order_date = '2024-03-01' on a date-partitioned table might only scan the March 2024 partition instead of years of data. This is great for time-series data where old data is no longer queried — you can TRUNCATE or DROP old partitions instantly instead of running slow DELETE statements. Sharding, in contrast, splits data across multiple servers — the application or a proxy layer must route queries to the right server. Partitioning is a single-server query optimization. Sharding is a distributed systems architecture for scaling beyond one server. I use partitioning frequently for log tables, time-series events, and archive data. I recommend sharding only when partitioning and vertical scaling are exhausted."*

---

## Q: "What is replication lag and how do you deal with it?"

> *"Replication lag is how far behind a replica is from the primary — measured in seconds by the Seconds_Behind_Master field in SHOW SLAVE STATUS. A replica with 30 seconds of lag is serving data that's 30 seconds stale. This matters because any read sent to the replica during that window might return outdated results. The causes: heavy write load on the primary that the replica's single SQL thread can't keep up with, large transactions that take a long time to replay, inefficient queries on the replica that run slowly due to missing indexes, or just network latency. The fixes: enable parallel replication — MySQL 5.7+ can replay non-conflicting transactions in parallel using multiple SQL threads. Keep transactions small and avoid massive single-transaction writes. Ensure replicas have the same or equivalent hardware as the primary — replicas do real work. For the application: implement 'read your own writes' — route queries immediately after writes to the primary for a short window. Use WAIT_FOR_EXECUTED_GTID_SET to make the application wait until the replica has caught up to a specific transaction before reading from it."*

---

## Q: "What is MySQL Group Replication?"

> *"MySQL Group Replication is MySQL's built-in high-availability solution with automatic failover. It uses a consensus protocol — similar to Paxos — to ensure all group members agree on the order of transactions before any member commits them. You need at minimum three nodes for a proper quorum. In single-primary mode, one node is the primary accepting writes, and the others are secondaries that can handle reads. If the primary fails, the group automatically elects a new primary from the surviving members — no manual DBA intervention needed. In multi-primary mode, all nodes can accept writes simultaneously, with built-in conflict detection that rolls back conflicting transactions. The trade-off of Group Replication versus traditional async replication: writes have higher latency because they must be certified across all members before committing — you're paying a network round-trip per transaction. InnoDB Cluster layers MySQL Shell and MySQL Router on top of Group Replication — the Router automatically directs writes to the current primary and can distribute reads across secondaries, so your application just connects to the Router and doesn't need to know the topology."*

