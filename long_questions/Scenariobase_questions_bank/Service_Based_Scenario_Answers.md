# Service-Based Scenario Answers (1-70)

## 🟢 Production & Debugging Scenarios (1–30)

### Question 1: A deployed service suddenly starts returning 500 errors. Logs are unclear. How do you debug?

**Answer:**
1.  **Check Infrastructure Metrics:** CPU, Memory, Disk I/O to see if the server is overloaded.
2.  **Check Dependency Health:** Verify if Database, Redis, or downstream APIs are reachable and healthy.
3.  **Enable Verbose Logging:** Temporarily increase log level to DEBUG if possible without redeploying (e.g., via dynamic config).
4.  **Reproduce Locally:** Try to reproduce with the same input in a lower environment.
5.  **Check Recent Changes:** Did a deployment or config change fail just before? Rollback if necessary.

---

### Question 2: CPU usage is normal, but response time has doubled. What do you check first?

**Answer:**
1.  **I/O Wait:** The thread might be waiting for Disk or Network (DB calls, external APIs).
2.  **Database Locking:** Slow queries or lock contention in the database.
3.  **Connection Pool Exhaustion:** Threads waiting to get a DB connection.
4.  **Garbage Collection:** Frequent "Stop-the-world" GC pauses (check JVM logs).
5.  **Network Latency:** Check if simple pings to dependencies are slow.

---

### Question 3: Memory usage keeps increasing over time. How do you identify a memory leak?

**Answer:**
1.  **Monitor Heap Usage:** Check if Heap usage grows without ever dropping back down after GC.
2.  **Heap Dump:** Take a Heap Dump (using `jmap` or similar tools) when memory is high.
3.  **Analyze Dump:** Use tools like Eclipse Memory Analyzer (MAT) or VisualVM to find objects retaining the most memory (Dominator Tree).
4.  **Check Static References:** Look for static collections (Maps, Lists) that keep growing and never get cleared.

---

### Question 4: An application crashes only under load, not in dev. How do you reproduce and fix it?

**Answer:**
1.  **Load Testing:** Use tools like JMeter, Gatling, or Apache Benchmark to simulate production-level traffic in a Staging environment.
2.  **Analyze Crash Logs:** Check JVM crash logs (`hs_err_pid.log`) or OOM errors.
3.  **Resource Limits:** Check if Prod has stricter limits (Memory, File Descriptors) than Dev.
4.  **Concurrency Issues:** Race conditions often only appear under high parallelism. Review code for thread safety.

---

### Question 5: A background job stops running after deployment. How do you investigate?

**Answer:**
1.  **Check Scheduler Logs:** Is the scheduler (Cron, Quartz, K8s CronJob) actually triggering the job?
2.  **Check Job Status:** Did it start and immediately fail (check logs for startup errors)?
3.  **Configuration:** Verify if the "Enabled" flag for the job was accidentally set to false in the new config.
4.  **Stuck Thread:** Is the previous execution still running (stuck) and blocking the new one? (Thread Dump).

---

### Question 6: Users report intermittent failures, but monitoring shows everything “green”. What do you do?

**Answer:**
1.  **Check Client-Side:** Is it a browser issue, specific OS, or network provider (CORS, DNS)?
2.  **Load Balancer Logs:** Check LB logs for 5xx errors that might not reach the app server.
3.  **Specific Edge Cases:** The monitoring might test a "Health Check" endpoint, but real users might fail on specific data inputs.
4.  **Timeouts:** Users might be timing out before the server responds (server shows 200 OK eventually, but user saw error).

---

### Question 7: A service works locally but fails in production. How do you debug environment issues?

**Answer:**
1.  **Configuration Diff:** Compare application.properties/yaml, env vars between Local and Prod.
2.  **Network Policies:** Firewalls or Security Groups might block ports in Prod that are open Locally.
3.  **Data Differences:** Prod data might have nulls or formats your local code doesn't handle.
4.  **Versions:** Check Java/Node/Python runtime versions and library versions.

---

### Question 8: After a config change, performance degrades. How do you roll back safely?

**Answer:**
1.  **Immediate Revert:** Re-deploy the previous known-good configuration artifact.
2.  **Feature Flags:** If the config was controlled by a flag, toggle it off.
3.  **Blue/Green Deployment:** If using Blue/Green, switch traffic back to the "Blue" (old) environment.
4.  **Verify:** Ensure performance metrics recover after rollback to confirm the config was the root cause.

---

### Question 9: Logs show timeouts when calling another service. How do you troubleshoot?

**Answer:**
1.  **Check Downstream Service:** Is the other service down or overloaded?
2.  **Network Trace:** Use `ping`, `traceroute`, or `tcpdump` to check for packet loss or network congestion.
3.  **Timeout Configuration:** Is the timeout set too low (e.g., 100ms) for normal variability?
4.  **Resource Contention:** Is the downstream service throttling your requests?

---

### Question 10: An API works for some users but not others. How do you isolate the issue?

**Answer:**
1.  **Segment Users:** Is it specific to a Region, Browser, Account Type (Free vs Paid), or Data Shard?
2.  **Inspect Request Data:** Compare headers and payloads of successful vs failed requests.
3.  **Check Permissions:** Is it an authorization issue (403 Forbidden) for specific roles?
4.  **Sticky Sessions:** If one server in the cluster is bad, users stuck to that server will fail.

---

### Question 11: A deployment caused partial outage. How do you minimize impact?

**Answer:**
1.  **Stop the Rollout:** Halt any ongoing deployment to new nodes.
2.  **Rollback:** Immediately revert updated nodes to the previous version.
3.  **Circuit Breaker:** If a specific feature is breaking, open the circuit breaker to disable just that feature.
4.  **Status Page:** Update status page to inform users (Transparency).

---

### Question 12: A cron job runs twice unexpectedly. How do you find the root cause?

**Answer:**
1.  **Multiple Instances:** Are there multiple instances of the app running (e.g., scaled up) and all have the scheduler enabled?
2.  **Retry logic:** Did the job fail (or timeout) and the scheduler retried it?
3.  **Clock Skew:** If running on multiple servers with unsynchronized clocks (rare with NTP).
4.  **Fix:** Implement distributed locking (e.g., via Redis or DB) so only one node can run the job.

---

### Question 13: An application hangs without crashing. What debugging steps do you take?

**Answer:**
1.  **Thread Dump:** Take a thread dump to see what threads are doing.
2.  **Analyze Blocks:** Look for Deadlocks (two threads waiting on each other) or threads stuck on I/O (waiting forever for a socket).
3.  **Infinite Loop:** Check for threads in RUNNABLE state consuming 100% CPU in the dump.
4.  **Memory:** Full GC loop (Application pauses repeatedly to clear memory).

---

### Question 14: You see thread pool exhaustion. How do you fix it?

**Answer:**
1.  **Identify Blocker:** Why are threads stuck? (Slow DB, Slow 3rd party API).
2.  **Increase Pool Size:** Temporary fix, but dangerous (can overload downstream).
3.  **Shorten Timeouts:** Fail fast so threads are released back to the pool quicker.
4.  **Async Processing:** Switch to non-blocking I/O (Reactive) if the load is high-concurrency/low-compute.

---

### Question 15: Garbage collection pauses are causing latency spikes. How do you diagnose?

**Answer:**
1.  **Enable GC Logs:** `-Xlog:gc*`.
2.  **Analyze Logs:** Use tools like GCeasy.io to see Pause Times and Frequency.
3.  **Tune Heap:** Increase Heap size if frequent Full GCs are due to lack of space.
4.  **Tune Algorithm:** Switch GC algorithms (e.g., from Parallel to G1GC or ZGC for lower latency).

---

### Question 16: A service restarts frequently in production. What do you check?

**Answer:**
1.  **OOM Killer:** Check `dmesg` or system logs to see if Linux killed the process due to Out Of Memory.
2.  **Health Check Failures:** Is Kubernetes/Load Balancer killing and restarting it because the health endpoint timed out?
3.  **Fatal Errors:** Check app logs for unhandled exceptions causing main thread crash.
4.  **Liveness Probe:** Is the liveness probe configuration too aggressive?

---

### Question 17: You receive a “disk full” alert on production. What actions do you take?

**Answer:**
1.  **Clean Logs:** Check `/var/log` or app log directories. gzip or delete old logs.
2.  **Clean Temp:** Check `/tmp` for leftover files.
3.  **Docker Prune:** Remove unused images/containers occupying space.
4.  **Root Cause:** Fix log rotation policy (ensure logs adhere to max size/days).

---

### Question 18: Logs are missing for some requests. How do you debug logging issues?

**Answer:**
1.  **Log Level:** Is the log level set to INFO but those requests are only logging at DEBUG?
2.  **Sampling:** Is the logging framework (or sidecar like Fluentd) sampling logs (dropping 90%) to save space?
3.  **Buffer Full:** If async logging is used, the buffer might be full and dropping events.
4.  **Crash:** Did the app crash before writing the log buffer to disk?

---

### Question 19: A feature works for admin users but not normal users. How do you debug?

**Answer:**
1.  **Permissions:** Check RBAC (Role Based Access Control) logic.
2.  **Data Scope:** Admins might see all data; users might have a `WHERE user_id = ?` clause that is buggy.
3.  **UI Code:** Is the front-end conditionally rendering broken code for verify-users?
4.  **Reproduce:** Impersonate a normal user in Staging to reproduce.

---

### Question 20: The application works during the day but fails at night. Why might this happen?

**Answer:**
1.  **Batch Jobs:** Heavy nightly ETL/Backup jobs sharing the same DB causing locks or high load.
2.  **Auto-Scaling:** Traffic drops at night, auto-scaler reduces nodes too much, causing overload on the remaining few.
3.  **3rd Party Maintenance:** External APIs might have maintenance windows at night.
4.  **Timezones:** Date rolling bugs (e.g., UTC vs Local time).

---

### Question 21: After scaling up instances, performance worsens. How do you debug?

**Answer:**
1.  **Database Bottleneck:** More app instances = More DB connections. The DB might be overwhelmed (Connection limit hit).
2.  **Cache Contention:** Thundering herd on the cache.
3.  **Lock Contention:** More nodes trying to acquire the same distributed locks.
4.  **Cold Start:** New instances need time to warm up (JIT compilation, connection pooling).

---

### Question 22: A service becomes slow only during peak hours. What metrics do you examine?

**Answer:**
1.  **Throughput (RPS):** Is traffic simply exceeding capacity?
2.  **Queue Depth:** Are threads/requests queuing up?
3.  **Latency Distribution:** Is it p99 spiking (bottleneck) or average (general load)?
4.  **Dependency Latency:** Are downstream services slowing down under load?

---

### Question 23: Requests queue up but workers are idle. What could be wrong?

**Answer:**
1.  **Configuration mismatch:** Queue has tasks, but workers are listening to the wrong queue/topic.
2.  **Connection Issue:** Workers lost connection to the message broker.
3.  **Prefetch Limit:** Workers might have a prefetch limit of 1 and are stuck on a "poison pill" message they can't process/ack.
4.  **Deadlock:** Worker threads are stuck waiting for something, appearing "idle" (not using CPU) but unable to pick new work.

---

### Question 24: You see connection pool exhaustion. How do you resolve it?

**Answer:**
1.  **Leak Check:** Are connections being closed (`conn.close()`) in `finally` blocks?
2.  **Pool Size:** Increase `max-connections` if the DB handles it.
3.  **Query Speed:** Slow queries hold connections longer. Optimize SQL.
4.  **Timeout:** Reduce `connection-timeout` so the app fails fast instead of hanging.

---

### Question 25: An app crashes when a specific input is sent. How do you debug safely in prod?

**Answer:**
1.  **Logs:** Find the stack trace associated with the crash.
2.  **Sanitize:** Identify the payload. Is it a "Zip Bomb", huge JSON, or SQL Injection attempt?
3.  **Fix Locally:** Do NOT test in Prod. Extract the payload and run unit tests locally.
4.  **WAF:** Block that specific pattern at the WAF (Web Application Firewall) level temporarily.

---

### Question 26: Monitoring shows normal latency but users complain of slowness. Why?

**Answer:**
1.  **Averages Deceive:** You are monitoring Average latency, but p99 is high (1% of users suffer).
2.  **Network Path:** The latency is between User and Server (CDN, ISP), not *on* the Server.
3.  **Rendering:** API is fast, but the Browser JS rendering is slow.
4.  **Missing Metrics:** You aren't measuring a specific critical blocking step (e.g., waiting for connection pool).

---

### Question 27: After JVM upgrade, memory usage increases. How do you analyze?

**Answer:**
1.  **Default Flags:** New Java versions change defaults (e.g., different GC, String Deduplication).
2.  **Compare Baseline:** Run same load test on old vs new version.
3.  **Native Memory:** Newer JVMs might use more native memory (Metaspace, CodeCache).
4.  **Heap Dump:** Compare dumps to see if object headers/sizes changed.

---

### Question 28: You detect zombie processes on a server. What steps do you take?

**Answer:**
1.  **Identify Parent:** Zombies are dead processes waiting for Parent to read exit code. Find parent PID.
2.  **Fix Parent:** The parent app is buggy (not waiting/reaping children).
3.  **Kill Parent:** Killing the parent will reparent zombies to `init` (PID 1), which reaps them.
4.  **Restart:** Reboot server if zombies consume all PID slots.

---

### Question 29: A third-party API suddenly becomes slow. How do you protect your system?

**Answer:**
1.  **Circuit Breaker:** Detect slowness and "open" circuit to stop calling them immediately.
2.  **Timeouts:** Ensure strict timeouts (fail fast).
3.  **Fallback:** Return cached data or default values instead of the API response.
4.  **Async:** Decouple: Accept user request, queue it, process API call in background so user doesn't wait.

---

### Question 30: How do you debug a production issue when you have no access to the server?

**Answer:**
1.  **Centralized Logs:** Rely on Splunk, ELK, or CloudWatch logs.
2.  **Metrics:** Use dashboards (Grafana, Datadog) to infer state (CPU, memory, custom metrics).
3.  **APM:** Use Application Performance Monitoring (New Relic, Dynatrace) for traces and stack profiles.
4.  **Replica:** Spin up a snapshot of the Prod Environment (if possible) to separate VPC you have access to.

---

## 🟡 Database & Backend Scenarios (31–55)

### Question 31: A database query suddenly becomes slow after data growth. What do you do?

**Answer:**
1.  **Explain Plan:** Run `EXPLAIN ANALYZE` to see if the query is doing a Full Table Scan.
2.  **Index:** Add missing indexes on columns used in `WHERE`, `JOIN`, or `ORDER BY` clauses.
3.  **Partitioning:** If the table is huge, consider partitioning by date or ID.
4.  **Archive:** Move old data to an archive table/warehouse to keep the active table small.

---

### Question 32: Deadlocks start appearing in the database. How do you analyze them?

**Answer:**
1.  **DB Logs:** Enable deadlock logging (e.g., `innodb_print_all_deadlocks` in MySQL).
2.  **Transaction Graph:** Identify the two transactions involved. A locks Row 1 then wants Row 2. B locks Row 2 then wants Row 1.
3.  **Lock Order:** Ensure all code acquires locks in the same order (e.g., always lock Table A before Table B).
4.  **Transaction Size:** Keep transactions short and fast.

---

### Question 33: High read latency but low CPU usage in DB. What might be wrong?

**Answer:**
1.  **I/O Bound:** The DB is waiting for the Disk to read data (IOPS limit reached).
2.  **Network:** Network bandwidth saturation or high latency.
3.  **Lock Waiting:** Queries are stuck waiting for locks held by other write transactions.
4.  **Client:** The application client might be overwhelmed, not reading the response fast enough.

---

### Question 34: Database connections are exhausted. How do you fix it?

**Answer:**
1.  **Connection Pooling:** Ensure app uses a pool (HikariCP) and reuses connections.
2.  **Leak:** Check for code not closing connections in `finally` block.
3.  **Scale:** Increase `max_connections` on DB (carefully, check RAM).
4.  **Cache:** Cache frequent read queries to reduce hits to the DB.

---

### Question 35: An index improves reads but slows writes. How do you decide?

**Answer:**
1.  **Read/Write Ratio:** If the app is 90% reads, the index is worth it. If 50% writes, reconsider.
2.  **Write Impact:** Measure how much slower the write is (10ms to 15ms? Acceptable. 10ms to 200ms? No.).
3.  **Async Index:** Some DBs allow async indexing or you can offload complex search to ElasticSearch.

---

### Question 36: A report query times out in production. How do you optimize it?

**Answer:**
1.  **Read Replica:** Run heavy reports on a Read Replica to avoid blocking the main DB.
2.  **Pre-aggregation:** Use Materialized Views to pre-calculate the data.
3.  **ETL:** Move reporting to an OLAP database (Data Warehouse) like Snowflake/BigQuery.
4.  **Optimization:** Rewrite SQL (avoid `SELECT *`, `OR`, `LIKE %...`).

---

### Question 37: Duplicate records appear in a table. How do you prevent this?

**Answer:**
1.  **Constraint:** Add a `UNIQUE` constraint or index on the relevant columns.
2.  **Idempotency:** Ensure the application logic handles retries correctly (check if exists before insert).
3.  **Cleanup:** Write a script to identifying duplicates (using `GROUP BY ... HAVING COUNT > 1`) and remove the newer ones.

---

### Question 38: Transactions are taking too long to commit. Why?

**Answer:**
1.  **Write Volume:** Too many rows being updated in one transaction.
2.  **Disk Sync:** DB flushing WAL (Write Ahead Log) to disk is slow (check Disk Latency).
3.  **Foreign Keys:** Checking FK constraints on huge tables.
4.  **Locking:** Waiting for other transactions to release locks.

---

### Question 39: A batch job locks tables and blocks users. How do you redesign it?

**Answer:**
1.  **Chunking:** Process data in small batches (e.g., 1000 rows), commit, sleep 10ms, repeat.
2.  **Timing:** Run during off-peak hours (night).
3.  **Row Locking:** Ensure the job uses Row-level locking instead of Table-level locking.
4.  **Isolation:** Read data with `READ UNCOMMITTED` if accuracy isn't critical for the batch.

---

### Question 40: Database size grows unexpectedly. How do you investigate?

**Answer:**
1.  **Analyze Tables:** Find largest tables (`SELECT table_name, data_length...`).
2.  **Audit Logs:** Is the app logging full payloads to a DB table?
3.  **Blobs:** storing large binary files (Images/PDFs) in DB instead of S3.
4.  **Fragmentation:** High delete rate leaving holes (need `OPTIMIZE TABLE` / `VACUUM`).

---

### Question 41: An API update causes DB load to spike. How do you find the cause?

**Answer:**
1.  **N+1 Problem:** Did the new code introduce an N+1 query issue (looping and querying DB per item)?
2.  **Missing Index:** New query filters by a column without an index.
3.  **Cache Miss:** Did the update break the cache logic?
4.  **Slow Query Log:** Check DB slow query log for new entries.

---

### Question 42: Foreign key constraints cause failures in prod. How do you debug?

**Answer:**
1.  **Order:** Are you inserting Child before Parent?
2.  **Data Integrity:** Does the Parent ID actually exist?
3.  **Orphans:** Are you deleting a Parent that still has Children (and no `CASCADE DELETE`)?
4.  **Fix:** Fix insertion order or use soft-deletes to maintain referential integrity.

---

### Question 43: A migration fails halfway in production. What are your next steps?

**Answer:**
1.  **Transactional DDL:** If using Postgres, DDL is transactional. It rolls back automatically. Safe.
2.  **Manual Rollback:** If MySQL (non-transactional DDL), determine which step failed. Manually reverse the changes (DROP the half-created table).
3.  **Fix Script:** Fix the error (e.g., data type mismatch) and re-run.
4.  **Backup:** Always backup before risky migrations.

---

### Question 44: Data is inconsistent across environments. How do you fix it?

**Answer:**
1.  **Sanitized Dump:** Restore a sanitized (PII removed) Prod backup to Staging.
2.  **Seed Data:** Use automated seed scripts to ensure baseline data exists everywhere.
3.  **Version Control:** Manage reference data (e.g., Country codes) via Listeners/Migrations code.

---

### Question 45: Slow deletes are impacting performance. How do you optimize?

**Answer:**
1.  **Soft Delete:** Update a flag `is_deleted = true` (Fast) instead of `DELETE`.
2.  **Batch Delete:** Delete in chunks (`DELETE FROM logs LIMIT 1000`) to avoid locking the table for seconds.
3.  **Partitioning:** Drop an entire old partition (Instant) instead of strictly deleting rows.

---

### Question 46: A service uses too many DB round trips. How do you detect and reduce them?

**Answer:**
1.  **Detect:** Use APM tools or ORM logging (Hibernate Statistics) to count queries per request.
2.  **Batching:** Use `IN` clauses (`SELECT * FROM users WHERE id IN (...)`) instead of loop.
3.  **JOINs:** Fetch related data in one query using JOINs.
4.  **Cache:** Cache the result of frequent lookups.

---

### Question 47: Read replicas lag behind primary. What problems can this cause?

**Answer:**
1.  **Stale Reads:** User updates profile, refreshes, sees old name (violates Read-Your-Writes).
2.  **Business Logic Errors:** Job reads "Pending" status from Replica, but it's actually "Done" on Primary.
3.  **Fix:** Use specific "Sticky" routing for recent writes, or force critical reads to Primary.

---

### Question 48: Application performance degrades after adding joins. How do you tune?

**Answer:**
1.  **Index Keys:** Ensure both columns used in the JOIN condition are indexed.
2.  **Denormalize:** If too expensive, combine data into one table (reduce 3NF).
3.  **Analyze:** Check if the DB is choosing a bad execution plan (e.g., Nested Loop vs Hash Join).

---

### Question 49: A query works fine with small data but fails at scale. Why?

**Answer:**
1.  **Full Table Scan:** It was fast for 100 rows, but scans 10M rows now.
2.  **Memory Sort:** `ORDER BY` spills to disk (Temp tables) because it doesn't fit in RAM.
3.  **Timeout:** The query simply exceeds the configured query timeout.

---

### Question 50: DB CPU is high but queries look simple. What do you check?

**Answer:**
1.  **Frequency:** A simple 1ms query run 100,000 times/sec can kill CPU.
2.  **Context Switching:** Too many concurrent connections/threads fighting for CPU.
3.  **Parsing:** Pre-compile implementation (Prepared Statements) might be missing, causing DB to parse SQL every time.

---

### Question 51: Pagination causes performance issues. How do you redesign?

**Answer:**
1.  **Problem:** `OFFSET 1000000` requires DB to read and discard 1M rows.
2.  **Keyset Pagination:** Use `WHERE id > last_seen_id LIMIT 10` (Seek Method).
3.  **Limit Depth:** Don't allow users to jump to page 10,000.

---

### Question 52: Large transactions cause rollback storms. How do you prevent?

**Answer:**
1.  **Scope:** Keep transactions as small as possible.
2.  **Ordering:** Acquire locks in consistent order to prevent deadlocks (which force rollbacks).
3.  **Retry Logic:** Add Jitter (random delay) to retries so they don't all come back instantly.

---

### Question 53: Data corruption is reported by users. How do you investigate?

**Answer:**
1.  **Scope:** Is it one record or thousands?
2.  **Logs:** Trace the request UUID that wrote the bad data.
3.  **Audit:** Enable DB query logging to see the exact UPDATE statement.
4.  **Code Review:** Check for logic bugs or race conditions (e.g., check-then-act without lock).

---

### Question 54: An audit shows missing records. How do you trace the issue?

**Answer:**
1.  **Application Logs:** Did the app report "Success" but fail to commit?
2.  **Silent Failures:** Caught exception but didn't log/re-throw?
3.  **Transaction Rollback:** Valid logic, but a later error rolled back the entire transaction.

---

### Question 55: Database backup causes performance degradation. How do you mitigate?

**Answer:**
1.  **Replica Backup:** Run the backup snapshot on the Read Replica, not the Primary.
2.  **Scheduling:** Run during absolute minimum traffic window.
3.  **Throttling:** Limit the I/O bandwidth used by the backup tool.

---

## 🔵 API, Integration & Deployment Scenarios (56–70)

### Question 56: A downstream service changes its API and breaks yours. How do you handle it?

**Answer:**
1.  **Immediate Fix:** Deploy a patch to adapt to the new API format.
2.  **Contact:** Reach out to the team (Internal) or Support (External).
3.  **Prevention:** Implement **Consumer Driven Contract (CDC)** tests (e.g., Pact) so they can't break you without failing CI.
4.  **Versioning:** Enforce strict API versioning policies.

---

### Question 57: An API returns inconsistent responses. How do you debug?

**Answer:**
1.  **Cache:** Is a stale cache layer returning old data intermittently?
2.  **Load Balancing:** Are different server nodes running different versions of the code?
3.  **Race Condition:** Concurrent requests modifying the same data.
4.  **Replication Lag:** Reading from a lagging DB replica.

---

### Question 58: Versioning issues appear after a release. How do you fix them?

**Answer:**
1.  **Strategy:** Adopt a clear strategy: URL Versioning (`/v1/user`, `/v2/user`) or Header Versioning (`Accept: application/vnd.app.v2+json`).
2.  **Backward Compatibility:** Ensure existing clients still work (non-breaking changes).
3.  **Deprecation:** Announce deprecation of old versions proactively.

---

### Question 59: A deployment works in staging but fails in prod. Why?

**Answer:**
1.  **Data Volume:** Prod usually has much more data (exposing slow queries).
2.  **Concurrency:** Prod has real traffic concurrency (exposing race conditions).
3.  **Config Drift:** Env vars, firewall rules, or hardware specs differ.
4.  **Third-Party:** Prod keys for Stripe/AWS might be invalid or expire.

---

### Question 60: A rollback also fails. What do you do?

**Answer:**
1.  **Fix Forward:** If rollback fails (e.g., DB schema changed non-backward compatibly), you must write a hotfix and deploy it forward.
2.  **Restore DB:** If data is corrupted, restore DB from backup (High Downtime).
3.  **Communication:** Escalate immediately and update Status Page.

---

### Question 61: API latency increases due to serialization. How do you diagnose?

**Answer:**
1.  **Profiling:** Cpu Profiler shows time spent in Jackson/Gson/JSON.stringify.
2.  **Payload Size:** Are you returning 5MB of JSON? Start formatting/filtering fields.
3.  **Protocol:** Consider switching to Protobuf or gRPC for internal high-volume calls (much faster/smaller).

---

### Question 62: Rate limits are exceeded unexpectedly. How do you investigate?

**Answer:**
1.  **Client Bug:** Is a client in an infinite retry loop?
2.  **Shared IP:** If rate limiting by IP, are multiple users behind a corporate NAT sharing one IP?
3.  **Configuration:** Is the limit per-instance (and you just scaled down) or global (Redis)?

---

### Question 63: A client sends malformed data. How do you protect your service?

**Answer:**
1.  **Validation:** Strong input validation (Schema validation like Zod, Joi, Bean Validation).
2.  **Reject Early:** Return `400 Bad Request` immediately.
3.  **Sanitize:** Strip dangerous chars to prevent XSS/SQLi.

---

### Question 64: Authentication works intermittently. How do you debug?

**Answer:**
1.  **Distributed State:** If using Sessions, is the session stored in Redis? (If in-memory, it fails on multi-node).
2.  **Clock Skew:** JWT tokens might be "expired" if server clocks differ.
3.  **CORS:** Browser rejecting requests on some subdomains.

---

### Question 65: A feature flag causes unexpected behavior. How do you isolate it?

**Answer:**
1.  **Audit:** Check the flag management audit log (Who changed it? When?).
2.  **Targeting:** Is the flag enabled for "100%" or just "Beta Users"?
3.  **Toggle:** Turn it to "False" globally to confirm if it resolves the issue.

---

### Question 66: An API gateway becomes a bottleneck. What do you do?

**Answer:**
1.  **Scale:** Horizontally scale the Gateway instances.
2.  **Cache:** Enable caching at the Gateway level for static responses.
3.  **Bypass:** Allow internal services to talk directly (Mesh) instead of hair-pinning through Gateway.

---

### Question 67: Network latency suddenly increases. How do you troubleshoot?

**Answer:**
1.  **MTR:** Run MTR (My Traceroute) to pinpoint the hop causing delay.
2.  **Cloud Status:** Check AWS/Azure status dashboard.
3.  **Bandwidth:** Is a large file transfer saturating the pipe?

---

### Question 68: Canary deployment shows mixed results. How do you decide to proceed?

**Answer:**
1.  **Compare Metrics:** Compare Error Rate and Latency of Canary vs Baseline.
2.  **Logs:** specific errors involved in Canary?
3.  **Decision:** If **any** Regression is found -> Rollback Canary. Don't gamble.

---

### Question 69: Clients complain after a schema change. How do you manage backward compatibility?

**Answer:**
1.  **Additive Changes:** Only **add** new fields. Never rename or delete.
2.  **Transformers:** Use a transformation layer to map DB columns to API response fields to maintain contract.
3.  **Contract Tests:** Ensure tests cover the response shape expected by clients.

---

### Question 70: Service health checks pass but functionality is broken. Why?

**Answer:**
1.  **Shallow Check:** Health check just returns `200 OK` (static).
2.  **Deep Check:** Update Liveness/Readiness probe to actually ping the DB and Critical Dependencies. (If DB down -> Return 503).
3.  **SLA:** Monitor "Synthetic Transactions" (fake user flow) to verify actual functionality.

---
