# 🟣 **661–680: Caching & Storage Systems**

### 661. How do you cache database query results in Go?
"I use the **Cache-Aside** pattern.
1.  Check Redis: `val, err := rdb.Get(ctx, key)`.
2.  If found (Hit), unmarshal and return.
3.  If missing (Miss), Query DB.
4.  Write to Redis (SetEX with TTL).
5.  Return data.
I wrap this logic in a generic function `GetCached[T]` to avoid repetition."

#### Indepth
**Thundering Herd**. If 1000 users request the same key simultaneously (Cache Miss), 1000 queries hit the DB. Use `singleflight` (Chapter 29) *before* the DB query. This ensures only 1 query runs, and the result is shared with all 1000 waiters. This is mandatory for hot keys.

---

### 662. How do you use Redis with Go for distributed caching?
"I use the `go-redis/v9` library.
It supports connection pooling and cluster mode automatically.
`rdb := redis.NewClient(...)`.
I prefer storing data as **Protobuf** or **MsgPack** (binary) instead of JSON to save space and CPU.
I always use a TTL to prevent the cache from filling up with stale data forever."

#### Indepth
**Pipelining**. If you need to set 50 keys, don't do 50 Round Trips. Use `rdb.Pipeline()`. It batches commands into a single TCP packet. `pipe.Set(...)`; `pipe.Exec(ctx)`. This reduces network latency by 50x for bulk operations.

---

### 663. How do you implement LRU cache in Go?
"I use `hashicorp/golang-lru` or implement it with a Map + Doubly Linked List.
Map stores `Key -> NodePtr`.
List stores nodes in order of access.
**Get**: Move node to front of list.
**Set**: If full, remove last node (tail), delete from map, add new node to front.
This guarantees O(1) operations."

#### Indepth
**Generics Implementation**. `hashicorp/golang-lru` uses `interface{}` (heap allocation). In Go 1.18+, you can write a type-safe generic LRU `[K comparable, V any]` to avoid these allocations. Combine this with a `sync.Pool` for the nodes to achieve zero-alloc LRU cache.

---

### 664. How do you ensure cache invalidation on data update?
"It's the hardest problem in CS.
Strategies:
1.  **TTL**: Accept staleness for N minutes.
2.  **Write-Through**: Update DB and deletions Cache in same transaction.
3.  **CDC (Change Data Capture)**: Listen to DB binlog (Debezium) -> Kafka -> Go Worker -> `redis.Del(key)`. This is the most robust web-scale solution."

#### Indepth
**Broadcasting**. If using Local Cache (in-memory), clearing Redis isn't enough. You must broadcast a "Invalidate Key X" message to *all* app instances (via Redis Pub/Sub). Each instance hears the event and deletes Key X from its local RAM.

---

### 665. How do you handle stale reads in Go apps with caching?
"I use the **Probabilistic Early Recomputation** (X-Fetch) pattern.
Store `beta` value.
If `TTL - now < beta * log(rand)`, I recompute the value in the background *while returning the stale value*.
This prevents the 'Thundering Herd' where everyone misses the cache at the exact same second."

#### Indepth
**Grace Period**. Another trick is "Soft vs Hard TTL". Item expires at 5 mins (Soft), but is kept for 10 mins (Hard). If accessed between 5-10m, return the Stale item immediately, but kick off a background refresh. This is cleaner than probabilistic math.

---

### 666. How do you implement a write-through cache in Go?
"My application wraps the Cache and DB.
`func SaveUser(u User) { tx := db.Begin(); repo.Save(tx, u); cache.Set(u.ID, u); tx.Commit() }`.
The downside is latency (writing to two places).
If Cache write fails, I often log a warning and let the cache be stale (eventually inconsistent) rather than failing the user request."

#### Indepth
**Write-Behind**. Write to Cache *only*, and let a background worker flush to DB. This is extremely fast (RAM speed) but risky (data loss if server crashes). Only use for non-critical data (like 'Like Counts' or 'User Presence').

---

### 667. How do you handle concurrency in in-memory caches?
"If using a simple map, I need `sync.RWMutex`.
`mu.RLock()` for getting. `mu.Lock()` for setting.
For high contention, I use **Sharding**.
`dgraph-io/ristretto` is a high-performance Go cache. It uses tiny LFUs and sharded locks to handle millions of Ops/sec without contention bottlenecks."

#### Indepth
**BigCache**. If you have GBs of data and GC pauses are killing you, use `bigcache`. It bypasses the GC by allocating a massive `[]byte` arena and storing entries as serialized bytes (pointers are hidden from GC). It effectively manually manages memory in Go.

---

### 668. How do you use bloom filters in Go?
"I use `bits-and-blooms/bloom`.
It’s a probabilistic set.
If `Test` says No, the key is **definitely not** in the set.
If `Test` says Yes, it *might* be (False Positive).
I use it before querying DB/Disk: 'Do we have user X?'. If No, save the disk IO. If Yes, check disk to confirm."

#### Indepth
**Sizing**. A Bloom Filter needs size (N) and error rate (P) upfront. If you underestimate N, the filter fills up and the False Positive rate spikes to 100%, rendering it useless. Always overprovision or use Scalable Bloom Filters that grow automatically.

---

### 669. How do you build a TTL-based memory cache?
"I store items with `Expiration int64`.
I run a background goroutine (Cleaner).
`ticker := time.NewTicker(1 * time.Minute)`.
`for range ticker.C { mu.Lock(); deleteExpired(); mu.Unlock() }`.
This 'Stop the World' cleanup is bad for large caches. Better approach: randomized sampling (Redis style) or a heap-based priority queue."

#### Indepth
**Active Expiration**. In user-land Go, rely on `Get()` to check expiry. `val, ok := map[key]; if val.Expired() { delete; return nil }`. The background cleaner is just a safety net for keys that are *never* accessed again to prevent memory leaks.

---

### 670. How do you use memcached in Go?
"I use `bradfitz/gomemcache`.
It’s simpler than Redis (Key-Value only, no structures).
It uses Consistent Hashing on the client side to distribute keys across multiple servers.
It’s great for raw HTML fragment caching (`GetMulti`) where Redis features aren't needed."

#### Indepth
**Slab Allocation**. Memcached never fragments memory because it uses Slabs (fixed size chunks). Redis uses `malloc`, which can lead to fragmentation over years. For pure, dumb, high-throughput caching of uniform objects, Memcached is technically superior in memory efficiency.

---

### 671. How do you store large binary blobs in Go?
"Not in the DB!
I store metadata in Postgres (`file_url`, `size`).
I stream the blob to **S3 / MinIO**.
Client Upload -> Go (generates Presigned URL) -> Client uploads to S3 directly.
This saves my Go server bandwidth and CPU."

#### Indepth
**Range Requests**. Using `http.ServeContent` with an `io.ReadSeeker` allows clients to request bytes 0-100 (Range Header). S3 supports this natively. If you proxy S3 through Go, ensure you pass the `Range` header through so video players can seek efficiently.

---

### 672. How do you build an append-only log file storage in Go?
"I open file with `O_APPEND`.
Write: `f.Write(entry); f.Sync()`.
To read efficiently, I maintain an in-memory index: `OffsetMap[ID] -> FilePosition`.
`f.Seek(pos, 0); f.Read(...)`.
This is exactly how Kafka and Bitcask storage engines work."

#### Indepth
**Log Rotation**. You can't write to one file forever. Implement rotation: when file reaches 1GB, rename to `data.1`, open new `data.active`. The reading index must now track `FileID + Offset`. Old files can be compacted (garbage collected) by removing deleted keys.

---

### 673. How do you use BoltDB or BadgerDB in Go?
"They are **Embedded KV Stores** (pure Go).
No external server process. Ideal for single-node apps.
**BoltDB**: B+Tree, read-heavy.
**BadgerDB**: LSM Tree, write-heavy.
I use them when deployment simplicity is priority (just one binary, no docker-compose for Redis)."

#### Indepth
**Read/Write Amplification**. BoltDB (B+Tree) does 1 read per page, but random writes are slow (updating pages). Badger (LSM) writes strictly sequentially (fast), but reads might need to check multiple SSTables. Choose Bolt for Read-Heavy (Config, CMS), Badger for Write-Heavy (Logs, Metrics).

---

### 674. How do you structure a file-based key-value store in Go?
"Naive: JSON file. Read all, modify, write all. Slow.
Better: **Log Structured Merge (LSM)**.
Writes go to MemTable (RAM). When full, flush to SSTable (Disk).
Compaction process merges old SSTables.
This implementation is complex; I usually defer to `LevelDB` or `Badger` unless learning."

#### Indepth
**Wal (Write Ahead Log)**. Before adding to MemTable, write to a pure append-only WAL file to survive crashes. On startup, replay the WAL to reconstruct the MemTable. This guarantees Durability (D in ACID).

---

### 675. How do you handle distributed caching with Go?
"I use **Groupcache** (by Google).
It’s a library, not a server.
Peers talk to each other.
If Peer A needs Key X, and Peer B owns Key X, A asks B.
It eliminates the 'Hot Key' problem because the hot key is stored in memory on the owner node, and others request it. It has no eviction, designed for static content."

#### Indepth
Groupcache does NOT support value updates (Immutable). It fails if you need to `Update("user:1", new_data)`. It works best for content that never changes (like a Blob based on Content-Hash) or changes very rarely. It powers Google's "dl.google.com" downloads.

---

### 676. How do you monitor cache hit/miss ratios in Go?
"I wrap my cache client.
`func Get(k) { increment('cache_total'); val := cache.Get(k); if val { increment('cache_hit') } else { increment('cache_miss') } }`.
I graph `hit / total` in Grafana.
If Hit Ratio drops < 80%, I investigate (TTL too low? Key eviction? Bad access pattern?)."

#### Indepth
**Metrics Cardinality**. Do NOT tag metrics with the *Key* (`cache_miss{key="user:123"}`). This creates infinite metrics and kills Prometheus. Tag by "Type" (`cache_miss{type="users"}`). You want to know if the "User Cache" is healthy, not specific keys.

---

### 677. How do you use consistent hashing for distributed caching?
"I map keys to a Ring of servers (0-360 degrees).
`hash(key)` finds the point on the ring.
I walk clockwise to find the first Server.
This ensures that adding a new Cache Node only invalidates 1/N keys, not ALL keys.
I use `stathat/consistent` or `serialx/hashring`."

#### Indepth
**Virtual Nodes**. To avoid lopsided distribution (one server getting 80% keys by bad luck), mapping 1 server to 100 "Virtual Nodes" on the ring. This statistically smoothes out the distribution so keys are spread evenly even with a small number of physical servers.

---

### 678. How do you build a cache warming strategy in Go?
"On startup, my app is cold (slow).
I implement a **Warmer**.
It reads the 'Top 1000 Accessed Keys' (from yesterday's logs).
It proactively fetches them from DB and populates Redis *before* the app marks itself as `/ready`.
This prevents the deployment latency spike."

#### Indepth
**Startup Probes**. K8s `startupProbe` is perfect here. It can fail for 60s while the cache warms up. Only once the warmer finishes does the app switch to `readinessProbe`, allowing traffic. This ensures users never hit a cold cache.

---

### 679. How do you use S3-compatible storage APIs in Go?
"I use the **MinIO SDK** or AWS SDK.
MinIO SDK is cleaner.
`minioClient.PutObject(ctx, bucket, name, reader, size, opts)`.
I ensure I handle **Context Cancellation**: if the user closes the connection, the upload to S3 should abort to save bandwidth."

#### Indepth
**Multipart Uploads**. For files > 100MB, splits them into 5MB chunks and upload in parallel goroutines. `minio` SDK handles this automatically. It also allows resuming failed uploads. Never upload a 1GB file in a single PUT request; a network blip at 99% forces a full restart.

---

### 680. How do you implement local persistent disk caching?
"I use a directory structure: `cache/ab/cd/abcdef...` (sharded folders).
Check file existence. Check mod time.
If expired or missing, fetch and write file.
I verify to use atomic file writes (`ioutil.TempFile` + `os.Rename`) so I never serve a half-written file to a concurrent reader."

#### Indepth
**LRU Deletion**. You can't keep writing files forever. Run a background garbage collector that checks disk usage. If Usage > 80%, delete the oldest files (based on mtime/atime) until Usage < 70%. `syscall.Statfs` gives you disk free space info.
