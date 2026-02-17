# 🟣 Go Theory Questions: 661–680 Caching & Storage Systems

## 661. How do you cache database query results in Go?

**Answer:**
We use the **Cache-Aside Pattern**.
1.  Check Cache: `val, err := redis.Get(key)`.
2.  If Hit: Return `val`.
3.  If Miss: Query DB. `val = db.Query(...)`.
4.  Set Cache: `redis.Set(key, val, 10*time.Minute)`.
We usually marshal the Go struct to JSON or Protobuf before storing it in the cache.

---

## 662. How do you use Redis with Go for distributed caching?

**Answer:**
We use `go-redis/redis`.
Client setup:
`rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})`.

We use it in middleware or service layer.
For High Availability, we use `FailoverClient` (Sentinel) or `ClusterClient` (Redis Cluster). Go's concurrency model pairs perfectly with Redis pipelining (`rdb.Pipeline()`) to send 100 commands in one round-trip.

---

## 663. How do you implement LRU cache in Go?

**Answer:**
LRU (Least Recently Used) requires a **Doubly Linked List** (for O(1) move-to-front) and a **Map** (for O(1) lookup).
We use `hashicorp/golang-lru`.

`cache, _ := lru.New(128)`
`cache.Add("key", "val")`
`val, ok := cache.Get("key")`
When `Add` exceeds 128 items, it automatically evicts the tail (oldest accessed). We use this for in-process memory caching.

---

## 664. How do you ensure cache invalidation on data update?

**Answer:**
There are two hard things in CS...
Strategy: **Delete on Write**.
When `UpdateUser(id)` succeeds in DB:
We call `redis.Del("user:"+id)`.
Next read will miss and fetch fresh data.
We do *not* try to `Set` the new value immediately because of race conditions involved in concurrent writes. Deleting is safer (eventual consistency).

---

## 665. How do you handle stale reads in Go apps with caching?

**Answer:**
We accept them, or we version them.
Time-To-Live (TTL) is the main mechanism. "Data is fresh for 5 minutes".
If we need strict freshness, we check `Last-Modified` headers or use a `version` field in the key: `redis.Get("user:123:v5")`.
If the DB moves to v6, the cache key changes, forcing a miss.

---

## 666. How do you implement a write-through cache in Go?

**Answer:**
In Write-Through, the application writes to the **Cache Wrapper**, which writes to DB *and* Cache synchronously.

```go
func (s *Store) Save(u User) {
   db.Save(u) // 1. Writes Truth
   cache.Set(u.ID, u) // 2. Updates Cache
}
```
This ensures the cache is always warm. The downside is write latency (2 network calls).

---

## 667. How do you handle concurrency in in-memory caches?

**Answer:**
Standard `map` is not thread-safe.
1.  **sync.RWMutex**: `mu.RLock()` for reading, `mu.Lock()` for writing.
2.  **sync.Map**: Optimized for stable keys and high read volume.
3.  **Sharding**: `BigCache` or `FreeCache` shard the map into 256 sub-maps with 256 locks to reduce lock contention on high-core-count machines.

---

## 668. How do you use bloom filters in Go?

**Answer:**
Bloom Filter: Probabilistic set. "Possibly in set" or "Definitely not".
Use case: Avoid hitting DB for non-existent UUIDs.

We use `willf/bloom`.
`filter.Add(key)`.
`if !filter.Test(key) { return 404 }`.
If `Test` returns true, it *might* be there, so we check Redis/DB. This saves 99% of wasted lookups for invalid keys.

---

## 669. How do you build a TTL-based memory cache?

**Answer:**
Background cleanup goroutine.
Item struct: `{ Value, Expiry time.Time }`.

```go
go func() {
    for range time.Tick(1 * time.Minute) {
        mu.Lock()
        for k, v := range cache {
            if time.Now().After(v.Expiry) { delete(cache, k) }
        }
        mu.Unlock()
    }
}()
```
For high performance, we use a Priority Queue (Heap) ordered by expiry time, so we don't have to scan the whole map.

---

## 670. How do you use memcached in Go?

**Answer:**
`bradfitz/gomemcache`.
Similar to Redis but simpler (Key-Value only, no structures).
`mc.Set(&memcache.Item{Key: "foo", Value: []byte("bar")})`.
We use Memcached when we need multi-threaded architecture (Redis is single-threaded) or simple raw speed for massive object caching without persistence needs.

---

## 671. How do you store large binary blobs in Go?

**Answer:**
We **stream** them.
We do not load `[]byte` into RAM.
Interface: `io.Reader`.
We use `io.Copy(dst, src)`.
Destination is usually **Object Storage** (S3/MinIO) or disk. We rarely store blobs in the DB (Postgres TOAST is okay for small text, bad for images).

---

## 672. How do you build an append-only log file storage in Go?

**Answer:**
(Like Kafka/WAL).
1.  Open file in Append Mode: `os.OpenFile(..., os.O_APPEND|os.O_WRONLY)`.
2.  Write: `file.Write(data)`.
3.  Index: Maintain a Map `Offset -> FilePosition`.
This allows O(1) writes (no distinct seeking) and O(1) lookups if we keep the index in memory.

---

## 673. How do you use BoltDB or BadgerDB in Go?

**Answer:**
These are **Embedded Key-Value Stores** (Pure Go).
They run inside your process (no TCP).
**Bolt** (bbolt): Read-heavy, B+Tree, ACID (uses file locking).
**Badger**: Write-heavy, LSM Tree (Log Structured Merge), faster writes.
We use them for building databases (like Etcd uses Bolt) or for single-node apps where running a separate Postgres server is overkill.

---

## 674. How do you structure a file-based key-value store in Go?

**Answer:**
Naive approach: JSON file.
Better approach: **Directory Hash**.
Key: `abc12345`.
Path: `data/abc/12/34/5.json`.
This avoids having 1 million files in one folder (which chokes the OS).
Go `os.WriteFile` is atomic (on most filesystems if you write temp and rename), ensuring we don't have partial writes.

---

## 675. How do you handle distributed caching with Go?

**Answer:**
(See Q 237 - Consistent Hashing).
If we have 5 Redis nodes.
We don't want to query all 5.
We use `groupcache` (written by Google for dl.google.com).
It coordinates peers. If Node A needs Key X, and Key X belongs to Node B, A asks B. B fetches from DB and caches it. A caches it too (hot).
It eliminates "Thundering Herd" because only one node fetches from DB.

---

## 676. How do you monitor cache hit/miss ratios in Go?

**Answer:**
We use **Prometheus Counters**.
`cache_hits_total` and `cache_misses_total`.
Middleware/Wrapper:
```go
func Get(k string) {
    if found {
        metrics.Hits.Inc()
        return val
    }
    metrics.Misses.Inc()
    // fetch...
}
```
Grafana query: `rate(hits) / (rate(hits) + rate(misses))`. If this drops below 80%, we investigate.

---

## 677. How do you use consistent hashing for distributed caching?

**Answer:**
We map keys to a Ring (0-360 degrees).
We map Nodes to points on the Ring.
Key X belongs to the next Node clockwise.
Library: `stathat/consistent`.
When we add a Node, it takes over only 1/N keys from its neighbor. Standard Modulo hashing (`hash % N`) would reshuffle 100% of keys, causing massive cache misses.

---

## 678. How do you build a cache warming strategy in Go?

**Answer:**
**Cache Warming** pre-fills the cache on startup.
1.  **Static**: Iterate a list of "top 100 products" and `Set` them.
2.  **Shadow Traffic**: Replay yesterday's access logs against the new server before flipping the DNS switch.
In Go, we launch a `go func() { Warmup() }` on startup. Use a Semaphore/WorkerPool to avoid hammering the DB during warmup.

---

## 679. How do you use S3-compatible storage APIs in Go?

**Answer:**
We use `minio-go` or `aws-sdk-go`.
MinIO SDK is cleaner for generic S3 usage.
`minioClient.PutObject(ctx, bucket, objectName, reader, size, opts)`.
It handles **Multipart Uploads** automatically for large files (splitting 1GB file into 5MB chunks and uploading in parallel goroutines), which is a huge performance win.

---

## 680. How do you implement local persistent disk caching?

**Answer:**
We use a library like **Peterbourgon/diskv** or simple file logic.
We verify:
1.  **Cache Size**: Check directory usage. If > 1GB, delete oldest files (file mod time).
2.  **Crash Safety**: Use temp files + `os.Rename`.
This is useful for caching generated reports or images where re-generating is CPU expensive but fetching from network is also slow.
