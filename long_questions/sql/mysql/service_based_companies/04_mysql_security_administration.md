# 🔐 04 — MySQL Security & Administration
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

> Security and DBA skills — commonly tested for **DevOps, DBA, and senior backend roles** at service companies.

---

## 🔑 Must-Know Topics
- User management: CREATE USER, GRANT, REVOKE, ROLES
- Authentication plugins (caching_sha2_password vs mysql_native_password)
- Privilege levels: global, database, table, column
- SSL/TLS encryption in transit
- Backup strategies: mysqldump, mysqlpump, Percona XtraBackup
- Point-in-time recovery with binlog
- performance_schema and monitoring
- MySQL configuration tuning (my.cnf)

---

## ❓ Most Asked Questions

### Q1. How do you create users and manage privileges?

```sql
-- Create a user (MySQL 8.0+)
CREATE USER 'app_user'@'10.0.0.%'           -- allow from this IP range
    IDENTIFIED BY 'StrongPassword!123'
    PASSWORD EXPIRE INTERVAL 90 DAY         -- force rotation every 90 days
    FAILED_LOGIN_ATTEMPTS 5                 -- lock after 5 failed attempts
    PASSWORD_LOCK_TIME 1;                   -- locked for 1 day

-- Grant specific privileges (principle of least privilege)
GRANT SELECT, INSERT, UPDATE ON myapp.* TO 'app_user'@'10.0.0.%';
GRANT SELECT ON myapp.logs TO 'readonly_user'@'%';
GRANT EXECUTE ON PROCEDURE myapp.TransferFunds TO 'app_user'@'10.0.0.%';

-- Column-level grants (very fine-grained)
GRANT SELECT (id, name, email) ON myapp.users TO 'limited_user'@'%';

-- Superuser / DBA account (separate from app account)
GRANT ALL PRIVILEGES ON *.* TO 'dba_admin'@'localhost' WITH GRANT OPTION;

-- View grants
SHOW GRANTS FOR 'app_user'@'10.0.0.%';

-- Revoke privileges
REVOKE INSERT, UPDATE ON myapp.* FROM 'app_user'@'10.0.0.%';

-- Flush privileges (needed after direct mysql.user table edits, not needed after GRANT/REVOKE)
FLUSH PRIVILEGES;

-- Drop user
DROP USER IF EXISTS 'old_user'@'%';
```

---

### Q2. How do you use MySQL roles (8.0+)?

```sql
-- Roles: named collection of privileges — simplify managing multiple users
CREATE ROLE 'app_read', 'app_write', 'app_admin';

GRANT SELECT ON myapp.* TO 'app_read';
GRANT SELECT, INSERT, UPDATE, DELETE ON myapp.* TO 'app_write';
GRANT ALL PRIVILEGES ON myapp.* TO 'app_admin';

-- Assign roles to users
CREATE USER 'alice'@'%' IDENTIFIED BY 'AlicePass!';
GRANT 'app_write' TO 'alice'@'%';

CREATE USER 'bob'@'%' IDENTIFIED BY 'BobPass!';
GRANT 'app_read' TO 'bob'@'%';

-- Set default roles (auto-activated on login)
SET DEFAULT ROLE 'app_write' TO 'alice'@'%';

-- User activates a role in session
SET ROLE 'app_read';
SET ROLE ALL;        -- activate all granted roles

-- Check active roles
SELECT CURRENT_ROLE();

-- Show role grants
SHOW GRANTS FOR 'alice'@'%' USING 'app_write';
```

---

### Q3. How do you perform a MySQL backup? (mysqldump)

```bash
# Full database backup with mysqldump
mysqldump \
    --host=localhost \
    --user=backup_user \
    --password \
    --single-transaction \        # consistent snapshot without locking (InnoDB)
    --routines \                  # include stored procedures & functions
    --triggers \                  # include triggers
    --events \                    # include events
    --hex-blob \                  # blob data as hex (safer)
    --set-gtid-purged=OFF \       # don't include GTID info (for non-GTID replicas)
    myapp_db > myapp_backup_$(date +%Y%m%d_%H%M%S).sql

# Backup specific tables
mysqldump myapp_db orders order_items customers > transactions.sql

# Backup all databases
mysqldump --all-databases > all_dbs_backup.sql

# Compressed backup
mysqldump -u root -p myapp_db | gzip > myapp_$(date +%F).sql.gz

# Restore from backup
mysql -u root -p myapp_db < myapp_backup_20240115.sql

# Restore compressed
gunzip < myapp_20240115.sql.gz | mysql -u root -p myapp_db
```

---

### Q4. How do you do point-in-time recovery?

```bash
# Scenario: database corrupted at 2:30 PM, last backup was at midnight
# Goal: restore backup + replay binlog up to 2:29 PM

# Step 1: Restore the last full backup
mysql -u root -p myapp_db < /backups/myapp_20240115_000000.sql

# Step 2: Find binlog files created after backup
SHOW BINARY LOGS;
# Find the binlog file & position from the backup (in dump file header)

# Step 3: Replay binlog up to just BEFORE the corruption
mysqlbinlog \
    --start-datetime="2024-01-15 00:00:00" \
    --stop-datetime="2024-01-15 14:29:00" \
    /var/lib/mysql/mysql-bin.000042 \
    /var/lib/mysql/mysql-bin.000043 \
    | mysql -u root -p myapp_db

# With GTID: even simpler
mysqlbinlog \
    --include-gtids="source_uuid:1-12345" \
    /var/lib/mysql/mysql-bin.000042 \
    | mysql -u root -p

# Percona XtraBackup (for zero-downtime backups of large databases):
# xtrabackup --backup --target-dir=/backups/full
# xtrabackup --prepare --target-dir=/backups/full
# xtrabackup --copy-back --target-dir=/backups/full
```

---

### Q5. How do you configure SSL/TLS in MySQL?

```sql
-- Check if SSL is configured
SHOW VARIABLES LIKE '%ssl%';
SHOW STATUS LIKE 'Ssl_cipher';
SELECT * FROM performance_schema.replication_connection_configuration
WHERE CHANNEL_NAME = '' AND SSL_ALLOWED = 'YES';

-- Check current connection's SSL
SHOW SESSION STATUS LIKE 'Ssl_cipher';  -- empty = no SSL

-- Require SSL for a user (force encryption)
ALTER USER 'app_user'@'%' REQUIRE SSL;
-- Or:
GRANT SELECT ON myapp.* TO 'app_user'@'%' REQUIRE SSL;

-- More granular SSL requirements:
ALTER USER 'admin'@'%'
    REQUIRE SUBJECT '/C=US/ST=CA/O=MyCompany/CN=admin-cert'
    AND ISSUER '/C=US/ST=CA/O=MyCA/CN=root-cert';

-- my.cnf: enable SSL on server
-- [mysqld]
-- ssl-ca   = /etc/mysql/ssl/ca-cert.pem
-- ssl-cert = /etc/mysql/ssl/server-cert.pem
-- ssl-key  = /etc/mysql/ssl/server-key.pem

-- Connect with SSL from client:
-- mysql --ssl-ca=ca-cert.pem --ssl-cert=client-cert.pem --ssl-key=client-key.pem \
--       -u app_user -p -h dbhost myapp_db
```

---

### Q6. How do you monitor MySQL using performance_schema?

```sql
-- performance_schema: zero-overhead diagnostic tables (enabled by default in 5.7+)

-- Top queries by total execution time
SELECT
    DIGEST_TEXT                                           AS query,
    COUNT_STAR                                            AS executions,
    ROUND(AVG_TIMER_WAIT / 1e12, 3)                      AS avg_sec,
    ROUND(SUM_TIMER_WAIT / 1e12, 3)                      AS total_sec,
    ROUND(SUM_ROWS_EXAMINED / COUNT_STAR)                 AS avg_rows_examined,
    ROUND(SUM_ROWS_SENT / COUNT_STAR)                     AS avg_rows_sent
FROM performance_schema.events_statements_summary_by_digest
WHERE SCHEMA_NAME = 'myapp'
ORDER BY SUM_TIMER_WAIT DESC
LIMIT 10;

-- Tables with most I/O
SELECT object_name, count_star, sum_timer_wait/1e12 AS total_sec
FROM performance_schema.table_io_waits_summary_by_table
WHERE object_schema = 'myapp'
ORDER BY sum_timer_wait DESC LIMIT 10;

-- Current connections and running queries
SELECT
    id, user, host, db, command, time,
    LEFT(info, 100) AS query
FROM information_schema.processlist
WHERE command != 'Sleep'
ORDER BY time DESC;

-- Kill long-running query
KILL QUERY 1234;    -- kills the query but keeps connection
KILL 1234;          -- kills the entire connection + rolls back transaction
```

---

### Q7. What are the most important MySQL configuration parameters?

```ini
# my.cnf — key InnoDB and global settings

[mysqld]
# Memory
innodb_buffer_pool_size = 12G         # 70-80% of RAM — most important setting!
innodb_buffer_pool_instances = 8      # parallel buffer pool access
innodb_log_file_size = 1G             # larger = better write throughput, slower recovery
tmp_table_size = 256M                 # in-memory temp table threshold
max_heap_table_size = 256M

# Connections
max_connections = 500                 # max simultaneous connections
thread_cache_size = 128               # reuse threads to avoid creation overhead

# I/O
innodb_flush_log_at_trx_commit = 1    # ACID: 1=fully durable. 2=may lose 1 sec on crash
innodb_flush_method = O_DIRECT        # bypass OS cache (avoids double-caching)
innodb_io_capacity = 2000             # IOPS available to InnoDB (SSD: 2000-10000)
innodb_io_capacity_max = 4000

# Slow query log
slow_query_log = ON
slow_query_log_file = /var/log/mysql/slow.log
long_query_time = 0.5                 # log queries > 500ms
log_queries_not_using_indexes = ON

# Replication
server-id = 1
log-bin = mysql-bin
binlog_format = ROW
binlog_row_image = MINIMAL             # log only changed columns (smaller binlog)
expire_logs_days = 7                   # auto-purge binlog after 7 days
```

---

### Q8. How do you rotate and maintain MySQL error logs and binary logs?

```sql
-- Binary log maintenance
SHOW BINARY LOGS;                           -- list all binlog files and sizes
PURGE BINARY LOGS TO 'mysql-bin.000042';    -- delete logs up to (not including) this file
PURGE BINARY LOGS BEFORE '2024-01-01';     -- delete logs older than date

-- Auto-expire binlogs (set in my.cnf):
-- expire_logs_days = 7  (MySQL 5.7)
-- binlog_expire_logs_seconds = 604800  (MySQL 8.0 — 7 days)

-- Rotate error log without restarting:
FLUSH ERROR LOGS;
-- Or rename and flush:
-- mv /var/log/mysql/error.log /var/log/mysql/error.log.$(date +%F)
-- mysqladmin -u root -p flush-logs

-- General query log (normally off — very verbose)
SET GLOBAL general_log = 'OFF';
SET GLOBAL general_log_file = '/var/log/mysql/general.log';

-- Check log file locations
SHOW VARIABLES LIKE '%log%file%';
SHOW VARIABLES LIKE 'log_error';

-- Flush all logs
FLUSH LOGS;           -- rotates all logs, generates new binlog file
FLUSH BINARY LOGS;    -- rotate binlog only
```

---

### Q9. How do you audit and secure MySQL access?

```sql
-- View all users and their authentication plugins
SELECT user, host, plugin, account_locked, password_expired
FROM mysql.user
ORDER BY user;

-- Check for blank passwords (security risk)
SELECT user, host FROM mysql.user WHERE authentication_string = '' OR authentication_string IS NULL;

-- Check for users with global privileges (should be minimal)
SELECT user, host FROM mysql.user WHERE Super_priv = 'Y';

-- Rename root user (obscurity measure)
RENAME USER 'root'@'localhost' TO 'dba_admin'@'localhost';

-- Remove anonymous users
DELETE FROM mysql.user WHERE User = '';
FLUSH PRIVILEGES;

-- Disable remote root login
DELETE FROM mysql.user WHERE User = 'root' AND Host != 'localhost';
FLUSH PRIVILEGES;

-- MySQL Audit Plugin (Enterprise) or community alternatives:
-- MariaDB Audit Plugin, McAfee MySQL Audit Plugin
-- Logs: who logged in, what queries ran, from which IP

-- Enable connection logging in application (important for forensics):
-- Log: timestamp, user, query_hash, rows_affected, execution_time, source_IP
```

---

### Q10. How do you check and optimize table health?

```sql
-- CHECK TABLE — verify table integrity (use on MyISAM; for InnoDB use innodb_check)
CHECK TABLE users, orders, products;
-- Returns: status OK or errors

-- ANALYZE TABLE — update index statistics (does NOT modify data)
ANALYZE TABLE employees;
-- Run after large data changes so optimizer makes correct choices

-- OPTIMIZE TABLE — defragment table (reclaim space from deletes/updates)
OPTIMIZE TABLE orders;  -- rebuilds table: reclaims space, reorganizes rows
-- ⚠️ Locks table during rebuild! Use pt-online-schema-change for large tables
-- Use tool: pt-online-schema-change --alter "ENGINE=InnoDB" users

-- REPAIR TABLE — fix corrupted MyISAM tables (InnoDB is usually self-healing)
REPAIR TABLE old_myisam_table;

-- Check table sizes
SELECT
    table_name,
    ROUND((data_length + index_length) / 1024 / 1024, 2) AS total_mb,
    ROUND(data_length / 1024 / 1024, 2)                  AS data_mb,
    ROUND(index_length / 1024 / 1024, 2)                 AS index_mb,
    table_rows                                            AS est_rows
FROM information_schema.tables
WHERE table_schema = 'myapp'
ORDER BY total_mb DESC;
```

