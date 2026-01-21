## 🟣 Caching & Storage Systems (Questions 661-680)

### Question 661: How do you cache database query results in Go?

**Answer:**
Fetch, Check Cache (Redis), Return if Hit. If Miss, Query DB, Set Cache, Return.
**Pattern:** Decorator/Repository wrapper.

```go
func GetUser(id int) {
    val, err := redis.Get(key).Result()
    if err == nil { return decode(val) }
    
    u := db.GetUser(id)
    redis.Set(key, encode(u), 10*time.Minute)
    return u
}
```

---

### Question 662: How do you use Redis with Go for distributed caching?

**Answer:**
Use `go-redis/redis`.
Supports connection pooling, sentinel (HA), and cluster mode out of the box.
Use `pipelining` to send multiple cache sets in one RTT.

---

### Question 663: How do you implement LRU cache in Go?

**Answer:**
Use `hashicorp/golang-lru` or implement:
- **Map:** For O(1) access.
- **Doubly Linked List:** To track usage order.
Move item to front on access. Remove from tail when full.

---

### Question 664: How do you ensure cache invalidation on data update?

**Answer:**
Hardest problem in CS.
1.  **Write-Through:** Update DB and Cache simultaneously.
2.  **Delete-on-Write:** Update DB, *Delete* the cache key (Next read will refill it). Better because avoiding "stale update" races is easier than perfect synchronization.

---

### Question 665: How do you handle stale reads in Go apps with caching?

**Answer:**
- Accept it (Eventual Consistency).
- Use **Lease/Gutter:** If cache is down/missing, only allow single flight (singleflight group) to DB to prevent Thundering Herd.

---

### Question 666: How do you implement a write-through cache in Go?

**Answer:**
Encapsulate logic in a `Store` struct.
`Save(item)` method must:
1.  Start Tx.
2.  Write DB.
3.  Commit Tx.
4.  Write Cache.
If 4 fails, cache is stale (set short TTL to mitigate).

---

### Question 667: How do you handle concurrency in in-memory caches?

**Answer:**
- **Regular Map:** Not thread safe. Panic.
- **`sync.RWMutex`:** Protect map.
- **`sync.Map`:** Good for stable keys.
- **Sharding:** BigCache / FreeCache use sharding (many locks) to reduce contention on high write loads.

---

### Question 668: How do you use bloom filters in Go?

**Answer:**
Probabilistic structure to test "Is X definitely NOT in set?".
Efficient (bits vs bytes).
Library: `willf/bloom`.
Use before querying DB to save IO on non-existent keys (e.g., Checking used usernames).

---

### Question 669: How do you build a TTL-based memory cache?

**Answer:**
Struct with `value` and `expiration time`.
Background goroutine (cleaner) ticks every minute, locks map, iterates, and deletes expired items. (Or use a heap/timer for precision).

---

### Question 670: How do you use memcached in Go?

**Answer:**
Library: `bradfitz/gomemcache`.
Protocol is simpler than Redis (Text/Binary). No complex types (lists/sets), just Set/Get bytes.

---

### Question 671: How do you store large binary blobs in Go?

**Answer:**
Do not store in RAM or DB (usually).
Stream to Object Storage (S3/MinIO).
In Go, `io.Reader` interface allows streaming file upload directly from HTTP Request to S3 Request without loading whole file into memory.

---

### Question 672: How do you build an append-only log file storage in Go?

**Answer:**
Open file with `os.O_APPEND|os.O_WRONLY`.
Goroutine safe for writes (mostly, depending on OS atomicity guarantees for small writes).
Used for WAL (Write Ahead Logs).

---

### Question 673: How do you use BoltDB or BadgerDB in Go?

**Answer:**
Embedded Key-Value stores (No server needed, runs inside Go process).
**BoltDB:** Read-optimized (B+ Tree). Good for config/meta.
**BadgerDB:** Write-optimized (LSM Tree). Fast.
```go
db, _ := bolt.Open("my.db", 0600, nil)
db.Update(func(tx *bolt.Tx) error { ... })
```

---

### Question 674: How do you structure a file-based key-value store in Go?

**Answer:**
Simple: Directory per "Bucket", File per "Key", Content is Value.
Fast for reads (OS file cache).
Hard to list keys efficiently if directory is huge.

---

### Question 675: How do you handle distributed caching with Go?

**Answer:**
Use consistent hashing to map Keys -> Nodes.
Library: `stathat/consistent`.
Determine which cache server to query based on key hash.

---

### Question 676: How do you monitor cache hit/miss ratios in Go?

**Answer:**
Metrics.
```go
if cached {
    metrics.CacheHits.Inc()
} else {
    metrics.CacheMiss.Inc()
}
```
Visualize in Grafana. If hit rate drops < 80%, investigate.

---

### Question 677: How do you use consistent hashing for distributed caching?

**Answer:**
Ring architecture.
Hash keys and Node IPs into the same ring.
Walk clockwise to find the owner node.
Scaling: Adding a node only redistributes K/N keys, not all keys.

---

### Question 678: How do you build a cache warming strategy in Go?

**Answer:**
On deployment/startup, run a worker that queries the most popular keys (from analytics) and populates the cache *before* opening the HTTP port to traffic.
Prevents "Cold Cache" latency spikes.

---

### Question 679: How do you use S3-compatible storage APIs in Go?

**Answer:**
Use `minio-go` or `aws-sdk-go`.
They speak the S3 XML protocol.
Allows swapping AWS S3 with on-prem MinIO seamlessly.

---

### Question 680: How do you implement local persistent disk caching?

**Answer:**
Use `diskv` or `bigcache` (if configured for persistence).
Stores cached JSON responses on disk (`/tmp/cache/md5(url)`).
Serve from disk if fresh, else fetch remote.

---
