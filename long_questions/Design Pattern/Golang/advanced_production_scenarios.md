# ⚔️ Advanced Production Scenarios ("War Stories")

This guide covers **Senior/Principal level** scenarios. These are not "textbook" questions; they require understanding **internals, failure modes, and trade-offs**.

---

## 🔥 1. Production Debugging & Observability

### Scenario 1: "The Memory Leak"
**Problem:** A service running in Kubernetes keeps getting OOM Killed (Out of Memory) every 6 hours.
**Investigation & Solution:**
1.  **Observability First:** Check Grafana/Prometheus. Is it a slow linear growth (Eventual Leak) or a spike (Load related)?
2.  **Profiling:** Use `pprof`.
    *   Enable `net/http/pprof`.
    *   Take a **heap profile** when memory is normal.
    *   Take another **heap profile** just before it crashes (if possible) or when high.
    *   `go tool pprof -http=:8080 diff base.prof current.prof`.
3.  **Common Causes in Go:**
    *   **Goroutine Leaks:** Goroutines waiting on a channel that nobody writes to. Use `runtime.NumGoroutine()` or `pprof/goroutine` to spot 100k+ goroutines.
    *   **Timer Leaks:** `time.After()` in heavy loops (before Go 1.23). Use `time.NewTimer()` instead.
    *   **Sub-slice Retention:** Storing a small slice of a very large array keeps the large array in memory.

### Scenario 2: "The CPU Spike"
**Problem:** CPU usage jumps to 100%, causing latency to spike, but traffic volume is normal.
**Investigation:**
1.  **Profiling:** Capture a **CPU profile** (`go tool pprof profile`).
2.  **Analysis:** Look for:
    *   **Garbage Collection (GC) thrashing:** If `runtime.mallocgc` is top, you are allocating too much on the heap. **Fix**: Optimize structs, use `sync.Pool`.
    *   **Serialization:** JSON parsing (`encoding/json`) is expensive. **Fix**: Switch to `easyjson` or `Protobuf`.
    *   **Validations:** Regex compilation inside loops. **Fix**: `regexp.MustCompile` globally.

### Scenario 3: "The Deadlock"
**Problem:** The application is running (PID exists) but stopped responding to HTTP requests. Health check times out.
**Solution:**
1.  **Trace:** trigger a **Full Goroutine Dump** (`GET /debug/pprof/goroutine?debug=2`).
2.  **Analyze Blockage:** Look for goroutines stuck in `semacquire` (Mutex wait) or `chan receive`.
3.  **Cyclic Wait:** Service A calls B, B calls A. **Fix**: Remove cycle, use timeouts.
4.  **Mutex Logic:** `defer mu.Unlock()` is safe, but manual unlocking in complex if-else branches often leads to deadlocks.

---

## 🚀 2. Database Performance & Internals

### Scenario 4: "The Slow Query"
**Problem:** An endpoint fetching users was fast (50ms) but now takes 3s as the table grew to 10M rows.
**Investigation & Solution:**
1.  **Explain Analyze:** Run `EXPLAIN ANALYZE SELECT ...`.
    *   Look for **Seq Scan** (Full Table Scan) vs **Index Scan**.
2.  **Indexing Strategy:**
    *   **Composite Index:** If query is `WHERE region='US' AND status='active'`, an index on `(region)` is not enough. You need `(region, status)`.
    *   **Cardinality:** Indexing a Boolean column (True/False) is useless (Database still scans 50% of rows).
3.  **Partial Indexes:** Index only `active` users (`WHERE status = 'active'`) to save index size.

### Scenario 5: "The Hot Partition"
**Problem:** You use DynamoDB/Cassandra. Overall traffic is low, but requests for one specific customer are failing.
**Root Cause:** **Hot Key / Data Skew**. All traffic for "Big Customer" hits the same shard/partition.
**Solution:**
1.  **Write Sharding:** Append a random suffix to the Shard Key (e.g., `CustomerID_1`, `CustomerID_2`).
2.  **Read Aggregation:** Read from all suffixes `_1` to `_N` and merge results.
3.  **Caching:** Cache the hot data aggressively.

### Scenario 6: "Connection Pool Exhaustion"
**Problem:** During a traffic spike, DB throws "Too Many Connections".
**Bad Fix:** Increasing `max_connections` (DB CPU will choke).
**Good Fix:**
1.  **Queueing:** Implement a semaphore or worker pool in the Go app to limit concurrent DB requests.
2.  **Circuit Breaker:** If DB is slow, fail fast instead of piling up connections.
3.  **Proxy:** Use **PgBouncer** (for Postgres) to multiplex thousands of app connections into few DB connections.

---

## 🌐 3. Advanced Distributed Systems

### Scenario 7: "Distributed Locking"
**Problem:** Cron job runs on 5 replicas. You need to ensure the "Daily Report" is generated only ONCE.
**Solution:**
1.  **Optimistic Locking (DB):** `UPDATE jobs SET status='RUNNING' WHERE id=1 AND status='PENDING'`. If RowsAffected == 0, someone else took it.
2.  **Redis (Redlock):** Fast, but careful with TTL. If the job takes longer than TTL, lock releases, and another worker enters. **Fix:** Watchdog process to extend TTL.
3.  **Etcd/ZooKeeper:** Strong consistency locks. Best for critical correctness.

### Scenario 8: "Cache Stampede" (Thundering Herd)
**Problem:** A popular cache key expires. 10,000 requests hit the DB simultaneously to re-fetch it. DB crashes.
**Solution:**
1.  **X-Cache-Lock:** Local mutex. Only ONE goroutine fetches from DB, others wait.
2.  **Probabilistic Early Expiration:** If TTL is 60s, fetch probability increases from 50s to 60s. A request at 55s might re-fetch before it's actually empty.

### Scenario 9: "Idempotency in Payments"
**Problem:** User clicks "Pay" twice. Or, the response "Success" is lost in network, so client retries. User is charged twice.
**Solution:**
1.  **Idempotency Key:** Client generates a UUID (`req_123`).
2.  **Logic:**
    *   API checks unique constraint on `req_123` in SQL.
    *   If exists -> Return previous stored response.
    *   If new -> Process Payment -> Store Response -> Return.
3.  **Critical:** The "Check" and "Insert" must be atomic (Database Transaction).

---

## 💡 4. Go Specific Optimizations

### Scenario 10: "Sync.Pool for GC Pressure"
**Problem:** High GC pauses because you allocate generic `buffer []byte` for every HTTP request.
**Solution:**
*   Use `sync.Pool`.
*   `Get()` a buffer from the pool.
*   `Reset()` it.
*   `Put()` it back.
*   **Result:** Zero allocations on the hot path.

### Scenario 11: "Context Propagation"
**Problem:** A request takes 10s. The user cancels (closes browser). The backend keeps processing for 10s, wasting resources.
**Solution:**
*   Always pass `context.Context`.
*   In DB/HTTP calls, use `req.WithContext(ctx)`.
*   Listen to `<-ctx.Done()` in long loops.
*   If context cancels, abort work immediately.
