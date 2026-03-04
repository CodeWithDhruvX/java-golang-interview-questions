# 🗣️ Theory — MySQL Security & Administration
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How do you manage MySQL users and their privileges?"

> *"I follow the principle of least privilege — every user gets the minimum permissions needed for their role. For an application connecting to the database, I create a dedicated user — never root — and grant only SELECT, INSERT, UPDATE, DELETE on the specific database it needs. If the app uses stored procedures, I grant EXECUTE on those procedures rather than broad table-level DML. I scope the connection by host — app_user@10.0.0.% means this user can only connect from that IP range, not from any machine in the world. GRANT ALL to any non-admin user is a red flag. I use SHOW GRANTS FOR user@host to audit what's been granted. MySQL 8.0 introduced roles — named collections of privileges — which make managing privileges for many users much cleaner. You define a role like 'app_read' with SELECT, assign the role to multiple users, and if the permissions need to change, you update the role once instead of updating every user individually."*

---

## Q: "What is the most important MySQL database configuration parameter?"

> *"Without question, it's innodb_buffer_pool_size. The InnoDB buffer pool is the in-memory cache where InnoDB stores data pages and index pages. Every read goes through the buffer pool — if the page is already there, it's served from memory. If not, MySQL reads it from disk. Maximizing the buffer pool hit rate is the single best thing you can do for MySQL performance. The typical recommendation is to set it to 70 to 80 percent of total available RAM — you need to leave memory for the OS, the MySQL process itself, connections, and sort buffers. On a dedicated 16 GB database server, I'd set it to 12 GB. After that, the next most impactful settings are: innodb_log_file_size — larger means better write throughput but longer recovery time after crash. innodb_flush_log_at_trx_commit — setting it to 1 is fully ACID durable, setting it to 2 gives better performance at the cost of potential 1-second data loss on power failure. max_connections — set it to what your connection pool actually needs, not an arbitrarily large number."*

---

## Q: "How do you back up a MySQL database? What is the --single-transaction flag?"

> *"The primary tool for logical backups is mysqldump. It generates a SQL file with CREATE TABLE statements and INSERT statements that fully recreate the database. The most important flag for InnoDB tables is --single-transaction, which wraps the entire dump in a single transaction using MVCC — this gives you a consistent point-in-time snapshot without locking any tables. The dump sees the database exactly as it was when the transaction started, even as new writes continue in the background. Without --single-transaction, mysqldump uses table locks for consistency — which blocks all write operations on those tables for the entire duration of the dump. That's unacceptable in production. Other important flags: --routines to include stored procedures, --triggers to include triggers, --events to include scheduled events, --set-gtid-purged=OFF when restoring to a replica that manages its own GTID tracking. For databases too large for mysqldump to be practical — multi-hundred-gigabyte databases — I use Percona XtraBackup, which does hot physical backups with no table locks and much faster backup and restore times."*

---

## Q: "How does point-in-time recovery work in MySQL?"

> *"Point-in-time recovery is how you recover a database to a specific moment — say, just before an accidental DROP TABLE or a bad data migration. It requires two things: a full backup and the binary logs that start from when that backup was taken. The process: first restore the full backup to get the database to its state at backup time. Then use the mysqlbinlog tool to extract and replay all the binary log events from the backup timestamp up to just before the disaster occurred — stopping before the harmful event. The result is the database restored to the exact state it was in at that cutoff point. With GTID replication, this is even cleaner — you can specify to replay up to a specific GTID instead of a timestamp, which is more precise than time since timestamps can be ambiguous. This is why I always ensure the binary log is enabled in production — it's not just for replication, it's your recovery lifeline. And I regularly test that backup + binlog recovery works consistently, because a backup you've never tested restoring is not really a backup."*

---

## Q: "How do you identify and fix performance problems in MySQL?"

> *"My debugging flow starts with the slow query log — I enable it and set long_query_time to something like 0.5 seconds to catch queries taking more than half a second. I also enable log_queries_not_using_indexes for development environments. Once I have the slow queries, I analyze them with mysqldumpslow to get the highest-impact candidates — sorted by total execution time. Then I bring those queries into EXPLAIN. I look at the 'type' column first — ALL or index on large tables means trouble. I check the 'key' column to see which index was actually used. 'Extra' column tells me if there's a filesort or temporary table — both are expensive signals. If I see a full scan where I expect an index scan, I check whether the index exists, whether the query is written in a way that blocks index usage — like a function on the column — and whether index statistics are stale. ANALYZE TABLE refreshes statistics. For persistent monitoring, I use performance_schema.events_statements_summary_by_digest to see aggregated query patterns — total executions, average time, and rows examined — across the lifetime of the server."*

---

## Q: "What security hardening steps do you apply to a fresh MySQL installation?"

> *"Several things immediately. First, run mysql_secure_installation — it prompts to set the root password, remove anonymous users, remove the test database, and disable remote root login. These are the low-hanging fruit. Then: create application-specific users with narrow privileges — never have applications using root. Restrict user connections to specific IP addresses or ranges rather than '%' wildcard where possible. Enable SSL/TLS for all client connections — use REQUIRE SSL on accounts or configure the server to enforce it globally. Set password expiration policies with PASSWORD EXPIRE INTERVAL and FAILED_LOGIN_ATTEMPTS. Keep MySQL updated — security patches are released regularly and running a years-old version is a known vulnerability. Limit the MySQL process's OS-level privileges — run as a dedicated non-root system user. Disable LOAD DATA INFILE if not needed — it can be used to read arbitrary files. Audit user accounts periodically with SELECT user, host FROM mysql.user to remove stale accounts. For regulated environments, enable the MySQL Audit Log plugin or equivalent to record all queries for compliance."*

---

## Q: "How do you monitor MySQL health in production?"

> *"I use multiple layers. For real-time visibility: SHOW PROCESSLIST or information_schema.processlist shows running queries — I look for long-running queries or a pile-up of threads waiting in 'Locked' state. SHOW ENGINE INNODB STATUS gives a detailed snapshot — lock wait information, deadlock history, buffer pool stats, and the insert buffer. For historical metrics, I expose MySQL metrics to a monitoring stack — either via the mysqld_exporter for Prometheus and Grafana, or CloudWatch if on AWS RDS. The key metrics I watch: Queries per second and slow queries per second. InnoDB buffer pool hit rate — should be above 99% in steady state. Replication lag — Seconds_Behind_Master on replicas. Thread count and waiting threads. Disk I/O wait. Long-running transactions via information_schema.innodb_trx. Connection count vs max_connections — if you're near the limit, connections start refusing. Temp tables created on disk — means tmp_table_size is too low or queries are doing hash aggregations beyond memory. I alert on: replication lag > 30s, buffer pool hit rate < 95%, connections > 80% of max, any long-running transaction > 60 seconds."*

