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

### Explanation
Database query result caching in Go follows a cache-aside pattern where you first check the cache, return if found, otherwise query the database and populate the cache. This is typically implemented as a decorator or repository wrapper around the data access layer, using Redis or similar for fast lookups.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you cache database query results in Go?
**Your Response:** "I implement database query result caching using the cache-aside pattern. First, I check if the data exists in Redis cache - if it's a hit, I return the cached result immediately. If it's a miss, I query the database, populate the cache with the result, and then return it. I typically implement this as a decorator or repository wrapper around my data access layer. The cache key is usually something like 'user:123' and I set a reasonable TTL like 10 minutes. This approach significantly reduces database load for frequently accessed data while keeping the implementation simple and maintainable. The pattern ensures that the cache is populated lazily - only when data is actually requested."

---

### Question 662: How do you use Redis with Go for distributed caching?

**Answer:**
Use `go-redis/redis`.
Supports connection pooling, sentinel (HA), and cluster mode out of the box.
Use `pipelining` to send multiple cache sets in one RTT.

### Explanation
Redis with Go uses the go-redis/redis library which provides connection pooling, high availability with sentinel mode, and cluster mode support. Pipelining allows sending multiple commands in a single round-trip, significantly improving performance for bulk operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use Redis with Go for distributed caching?
**Your Response:** "I use Redis with Go through the `go-redis/redis` library which is the standard client. It provides connection pooling out of the box, supports high availability with sentinel mode, and can scale horizontally with cluster mode. For performance optimization, I use pipelining to send multiple cache operations in a single round-trip to Redis. This is especially useful when I need to set multiple cache keys at once. The library handles all the complexity of Redis protocol communication while providing a clean, idiomatic Go interface. I can easily switch between single-instance, sentinel, or cluster configurations just by changing the client options."

---

### Question 663: How do you implement LRU cache in Go?

**Answer:**
Use `hashicorp/golang-lru` or implement:
- **Map:** For O(1) access.
- **Doubly Linked List:** To track usage order.
Move item to front on access. Remove from tail when full.

### Explanation
LRU cache in Go uses a hash map for O(1) access combined with a doubly linked list to track usage order. The hashicorp/golang-lru library provides a robust implementation. Items are moved to the front when accessed and removed from the tail when the cache is full.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement LRU cache in Go?
**Your Response:** "I implement LRU cache using a combination of a hash map and doubly linked list. The hash map gives me O(1) access to cache items, while the doubly linked list tracks the usage order - recently used items are at the front, least recently used at the back. When an item is accessed, I move it to the front of the list. When the cache is full, I remove items from the tail. In practice, I usually use the `hashicorp/golang-lru` library which provides a battle-tested implementation. This data structure is perfect for caching scenarios where I want to keep frequently accessed items in memory while automatically evicting the least useful ones when space is needed."

---

### Question 664: How do you ensure cache invalidation on data update?

**Answer:**
Hardest problem in CS.
1.  **Write-Through:** Update DB and Cache simultaneously.
2.  **Delete-on-Write:** Update DB, *Delete* the cache key (Next read will refill it). Better because avoiding "stale update" races is easier than perfect synchronization.

### Explanation
Cache invalidation is challenging because write-through updates both DB and cache simultaneously but risks stale data, while delete-on-write updates the DB and deletes the cache key, allowing the next read to repopulate fresh data. Delete-on-write is preferred as it avoids complex synchronization issues.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you ensure cache invalidation on data update?
**Your Response:** "Cache invalidation is one of the hardest problems in computer science. I use two main approaches: write-through and delete-on-write. Write-through updates both the database and cache simultaneously, but this can lead to race conditions where stale data persists. I prefer delete-on-write where I update the database and delete the cache key, so the next read will repopulate the cache with fresh data. This approach is simpler and more reliable because avoiding 'stale update' race conditions is easier than achieving perfect synchronization between cache and database. The trade-off is a brief cache miss, but it ensures data consistency and prevents serving stale data."

---

### Question 665: How do you handle stale reads in Go apps with caching?

**Answer:**
- Accept it (Eventual Consistency).
- Use **Lease/Gutter:** If cache is down/missing, only allow single flight (singleflight group) to DB to prevent Thundering Herd.

### Explanation
Stale reads in cached Go applications are handled through eventual consistency acceptance or lease/gutter patterns. The singleflight group ensures only one request hits the database when cache is down, preventing thundering herd problems where multiple requests simultaneously try to repopulate cache.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle stale reads in Go apps with caching?
**Your Response:** "I handle stale reads by either accepting eventual consistency or implementing a lease/gutter pattern. Most of the time, I accept that cached data might be slightly stale - this is the trade-off for performance. However, when cache failures occur, I use a singleflight group to ensure only one request hits the database to repopulate the cache. This prevents the thundering herd problem where multiple requests simultaneously try to refresh the same data. The singleflight pattern coordinates concurrent requests, allowing the first one to fetch data while others wait for the result. This approach balances performance with data safety and prevents overwhelming my database during cache recovery."

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

### Explanation
Write-through cache in Go encapsulates logic in a Store struct where the Save method performs database operations within a transaction before updating the cache. If cache write fails, the data remains stale but short TTL mitigates this. The transaction ensures database consistency before cache population.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a write-through cache in Go?
**Your Response:** "I implement write-through cache by encapsulating the logic in a Store struct. The Save method follows a specific sequence: first start a database transaction, write to the database, commit the transaction, and then write to the cache. If the cache write fails, the database operation is already committed, so the cache might be stale temporarily. I mitigate this by using a short TTL so the cache will be refreshed quickly. The transaction ensures database consistency before I attempt to update the cache. This approach provides strong consistency for the database while accepting brief cache inconsistencies, which is a reasonable trade-off for most applications."

---

### Question 667: How do you handle concurrency in in-memory caches?

**Answer:**
- **Regular Map:** Not thread safe. Panic.
- **`sync.RWMutex`:** Protect map.
- **`sync.Map`:** Good for stable keys.
- **Sharding:** BigCache / FreeCache use sharding (many locks) to reduce contention on high write loads.

### Explanation
Concurrency in in-memory caches requires different approaches based on access patterns. Regular maps are not thread-safe, sync.RWMutex protects maps with reader-writer locks, sync.Map works well for stable key sets, and sharding reduces contention by using multiple locks as seen in BigCache/FreeCache for high-write scenarios.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle concurrency in in-memory caches?
**Your Response:** "I handle concurrency in in-memory caches using several approaches depending on the use case. For simple cases, I use a regular map protected by `sync.RWMutex` which allows multiple concurrent readers but exclusive writers. For scenarios with stable key sets, `sync.Map` is optimized for that pattern and provides better performance. For high-write loads, I use sharding where I split the cache into multiple segments, each with its own lock - this is how libraries like BigCache and FreeCache achieve high performance. The choice depends on the access pattern - RWMutex for general use, sync.Map for read-heavy stable data, and sharding for write-intensive scenarios. Each approach trades off complexity for performance characteristics."

---

### Question 668: How do you use bloom filters in Go?

**Answer:**
Probabilistic structure to test "Is X definitely NOT in set?".
Efficient (bits vs bytes).
Library: `willf/bloom`.
Use before querying DB to save IO on non-existent keys (e.g., Checking used usernames).

### Explanation
Bloom filters in Go are probabilistic data structures that efficiently test whether an item is definitely not in a set. They use minimal space (bits vs bytes) and are useful for avoiding expensive database lookups for non-existent keys. The willf/bloom library provides a Go implementation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use bloom filters in Go?
**Your Response:** "I use bloom filters as a probabilistic data structure to efficiently test whether an item is definitely not in a set. They're incredibly space-efficient, using bits instead of bytes to represent membership. I typically use the `willf/bloom` library in Go. The key insight is that bloom filters can tell me with certainty that something is NOT in a set, though they might have false positives. This makes them perfect for use cases like checking if a username is already taken - if the bloom filter says it's definitely not taken, I can skip the database lookup entirely. Only when it might be taken do I need to query the database. This saves a lot of unnecessary IO operations for non-existent keys."

---

### Question 669: How do you build a TTL-based memory cache?

**Answer:**
Struct with `value` and `expiration time`.
Background goroutine (cleaner) ticks every minute, locks map, iterates, and deletes expired items. (Or use a heap/timer for precision).

### Explanation
TTL-based memory cache in Go uses structs containing values and expiration times. A background goroutine periodically scans and removes expired items. For more precision, a heap/timer approach can be used instead of periodic scanning. This provides automatic cleanup of stale cache entries.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a TTL-based memory cache?
**Your Response:** "I build TTL-based memory caches using structs that store both the value and expiration time. I run a background goroutine that periodically scans the cache - typically every minute - and removes expired items. The cleaner goroutine locks the map during iteration to ensure thread safety. For more precise expiration handling, I can use a heap with timers instead of periodic scanning, which triggers cleanup exactly when items expire. The periodic approach is simpler and works well for most use cases, while the heap/timer approach provides better precision at the cost of complexity. This automatic cleanup prevents memory leaks from expired cache entries building up over time."

---

### Question 670: How do you use memcached in Go?

**Answer:**
Library: `bradfitz/gomemcache`.
Protocol is simpler than Redis (Text/Binary). No complex types (lists/sets), just Set/Get bytes.

### Explanation
Memcached in Go uses the bradfitz/gomemcache library which implements the simple text/binary protocol. Unlike Redis, memcached only supports basic Set/Get operations on byte data without complex data types like lists or sets, making it simpler but less feature-rich.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use memcached in Go?
**Your Response:** "I use memcached in Go with the `bradfitz/gomemcache` library. It implements the memcached protocol which is simpler than Redis - it uses either text or binary protocols and only supports basic Set and Get operations on byte data. Unlike Redis, memcached doesn't have complex data types like lists or sets, which makes it simpler but less feature-rich. I choose memcached when I need straightforward key-value caching without the complexity of Redis's advanced features. The library is lightweight and easy to use, making it perfect for simple caching scenarios where I just need to store and retrieve serialized data."

---

### Question 671: How do you store large binary blobs in Go?

**Answer:**
Do not store in RAM or DB (usually).
Stream to Object Storage (S3/MinIO).
In Go, `io.Reader` interface allows streaming file upload directly from HTTP Request to S3 Request without loading whole file into memory.

### Explanation
Large binary blobs in Go should not be stored in RAM or databases. Instead, stream them to object storage like S3 or MinIO. The io.Reader interface enables direct streaming from HTTP requests to S3 without loading entire files into memory, making it memory-efficient for large file handling.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you store large binary blobs in Go?
**Your Response:** "I avoid storing large binary blobs in RAM or databases. Instead, I stream them directly to object storage like AWS S3 or MinIO. The beauty of Go is that the `io.Reader` interface allows me to stream files directly from an HTTP request to an S3 request without ever loading the entire file into memory. This approach is incredibly memory-efficient and can handle files of any size. I just connect the reader from the upload to the writer for S3, and Go handles the streaming automatically. This prevents memory exhaustion and allows my application to handle large file uploads efficiently, even with limited memory resources."

---

### Question 672: How do you build an append-only log file storage in Go?

**Answer:**
Open file with `os.O_APPEND|os.O_WRONLY`.
Goroutine safe for writes (mostly, depending on OS atomicity guarantees for small writes).
Used for WAL (Write Ahead Logs).

### Explanation
Append-only log file storage in Go uses os.Open with O_APPEND and O_WRONLY flags. This is goroutine-safe for writes depending on OS atomicity guarantees for small writes. This pattern is commonly used for Write Ahead Logs (WAL) in database systems for durability and crash recovery.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build an append-only log file storage in Go?
**Your Response:** "I build append-only log storage by opening files with `os.O_APPEND|os.O_WRONLY` flags. This ensures all writes are appended to the end of the file. The writes are generally goroutine-safe, especially for small writes, thanks to OS-level atomicity guarantees. This pattern is perfect for Write Ahead Logs (WAL) in database systems where I need durable, crash-recoverable storage. The append-only nature makes writes very fast since I don't need to seek, and it provides a natural audit trail of all operations. This approach is widely used in databases and message queues for ensuring data durability and supporting recovery after crashes."

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

### Explanation
BoltDB and BadgerDB are embedded key-value stores that run within the Go process without requiring external servers. BoltDB is read-optimized using B+ trees, suitable for configuration and metadata. BadgerDB is write-optimized using LSM trees, providing high performance for write-heavy workloads.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use BoltDB or BadgerDB in Go?
**Your Response:** "I use BoltDB and BadgerDB as embedded key-value stores that run inside my Go process without needing external servers. BoltDB is read-optimized using B+ trees, making it great for configuration and metadata where I do more reading than writing. BadgerDB is write-optimized using LSM trees, providing excellent performance for write-heavy workloads. I choose between them based on my access patterns - BoltDB for read-heavy scenarios, BadgerDB for write-heavy applications. Both provide transaction support and ACID guarantees, making them suitable for reliable local storage. The embedded nature means no separate database server to manage, which simplifies deployment and reduces operational complexity."

---

### Question 674: How do you structure a file-based key-value store in Go?

**Answer:**
Simple: Directory per "Bucket", File per "Key", Content is Value.
Fast for reads (OS file cache).
Hard to list keys efficiently if directory is huge.

### Explanation
File-based key-value stores in Go use a simple structure where each bucket is a directory and each key is a file within that directory. This leverages OS file caching for fast reads but becomes inefficient for listing keys when directories become very large.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you structure a file-based key-value store in Go?
**Your Response:** "I structure file-based key-value stores using a simple approach where each bucket becomes a directory and each key becomes a file containing the value. This leverages the OS file cache for fast reads since the operating system handles caching automatically. The structure is intuitive - I can navigate the filesystem to inspect data manually. However, this approach has limitations - listing keys efficiently becomes difficult when directories contain thousands of files. I use this pattern for simple applications where the key count is manageable and I want the ability to inspect data directly through the filesystem. It's great for development and debugging, though not suitable for high-performance production scenarios with large key sets."

---

### Question 675: How do you handle distributed caching with Go?

**Answer:**
Use consistent hashing to map Keys -> Nodes.
Library: `stathat/consistent`.
Determine which cache server to query based on key hash.

### Explanation
Distributed caching in Go uses consistent hashing to map keys to cache nodes. The stathat/consistent library provides this functionality, determining which cache server to query based on key hash. This approach minimizes data redistribution when nodes are added or removed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle distributed caching with Go?
**Your Response:** "I handle distributed caching using consistent hashing to map keys to cache nodes. I use the `stathat/consistent` library which provides a robust implementation of consistent hashing. When I need to store or retrieve a value, I hash the key and the consistent hashing algorithm tells me which cache server should handle that key. The beauty of this approach is that when I add or remove cache nodes, only a small portion of keys need to be remapped - specifically K/N keys where K is total keys and N is number of nodes. This minimizes cache disruption during scaling operations and ensures even distribution of load across the cache cluster."

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

### Explanation
Cache hit/miss ratio monitoring in Go uses metrics counters to track cache performance. Cache hits and misses are incremented accordingly and visualized in Grafana. A hit rate below 80% typically indicates caching issues that need investigation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you monitor cache hit/miss ratios in Go?
**Your Response:** "I monitor cache hit/miss ratios using metrics counters. I increment a CacheHits counter when data is found in cache and a CacheMiss counter when it's not. I expose these metrics through Prometheus and visualize them in Grafana. I set up alerts to notify me if the hit rate drops below 80%, which usually indicates caching problems. This could mean my cache keys aren't optimal, TTLs are too short, or my cache size is insufficient. Monitoring these metrics helps me understand caching effectiveness and optimize performance. The hit/miss ratio is a key indicator of how well my caching strategy is working and whether users are getting the performance benefits I expect."

---

### Question 677: How do you use consistent hashing for distributed caching?

**Answer:**
Ring architecture.
Hash keys and Node IPs into the same ring.
Walk clockwise to find the owner node.
Scaling: Adding a node only redistributes K/N keys, not all keys.

### Explanation
Consistent hashing for distributed caching uses a ring architecture where both keys and node IPs are hashed onto the same ring. To find the owner node, you walk clockwise from the key position. Adding a node only redistributes K/N keys rather than all keys, minimizing disruption.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use consistent hashing for distributed caching?
**Your Response:** "I implement consistent hashing using a ring architecture. I hash both the cache keys and the node IP addresses onto the same circular ring. To determine which node should store a particular key, I find the key's position on the ring and walk clockwise until I encounter the first node - that node owns the key. The beauty of this approach is that when I add a new node, only K/N keys need to be redistributed, where K is the total number of keys and N is the number of nodes. This minimal redistribution is much better than traditional hashing where adding a node would require moving almost all keys. This makes scaling the cache cluster much more efficient and reduces the performance impact of adding or removing nodes."

---

### Question 678: How do you build a cache warming strategy in Go?

**Answer:**
On deployment/startup, run a worker that queries the most popular keys (from analytics) and populates the cache *before* opening the HTTP port to traffic.
Prevents "Cold Cache" latency spikes.

### Explanation
Cache warming strategies in Go involve running workers during deployment/startup that query popular keys from analytics and populate the cache before opening the HTTP port to traffic. This prevents cold cache latency spikes by ensuring frequently accessed data is already cached when the application starts serving requests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a cache warming strategy in Go?
**Your Response:** "I build cache warming strategies by running workers during application deployment or startup that pre-populate the cache with popular data. I analyze usage patterns or analytics to identify the most frequently accessed keys, then have workers query these keys and populate the cache before I open the HTTP port to traffic. This prevents the 'cold cache' problem where the first users after deployment experience high latency because the cache is empty. By warming the cache proactively, I ensure that when the application starts serving requests, the most important data is already cached and available instantly. This significantly improves user experience right after deployments and prevents performance degradation during the cache warm-up period."

---

### Question 679: How do you use S3-compatible storage APIs in Go?

**Answer:**
Use `minio-go` or `aws-sdk-go`.
They speak the S3 XML protocol.
Allows swapping AWS S3 with on-prem MinIO seamlessly.

### Explanation
S3-compatible storage APIs in Go use minio-go or aws-sdk-go libraries that implement the S3 XML protocol. This enables seamless swapping between AWS S3 and on-premises MinIO deployments without code changes, providing flexibility in storage infrastructure.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use S3-compatible storage APIs in Go?
**Your Response:** "I use S3-compatible storage APIs in Go with either `minio-go` or `aws-sdk-go` libraries. Both libraries speak the S3 XML protocol, which means I can swap between AWS S3 and on-premises MinIO deployments without changing my code. I just need to update the endpoint configuration. This flexibility is incredibly valuable - I can develop using MinIO locally for cost savings and speed, then deploy to AWS S3 in production. Or I can use MinIO for on-premises deployments while maintaining the same API. The libraries handle all the complexity of the S3 protocol, providing a clean Go interface for uploading, downloading, and managing objects. This abstraction makes my storage layer portable across different providers."

---

### Question 680: How do you implement local persistent disk caching?

**Answer:**
Use `diskv` or `bigcache` (if configured for persistence).
Stores cached JSON responses on disk (`/tmp/cache/md5(url)`).
Serve from disk if fresh, else fetch remote.

### Explanation
Local persistent disk caching in Go uses diskv or bigcache configured for persistence. Cached JSON responses are stored on disk using URL-based file paths like `/tmp/cache/md5(url)`. Fresh data is served from disk, while stale data triggers remote fetching and cache refresh.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement local persistent disk caching?
**Your Response:** "I implement local persistent disk caching using libraries like `diskv` or `bigcache` configured for persistence. I store cached JSON responses on disk using URL-based paths like `/tmp/cache/md5(url)` where the filename is the MD5 hash of the URL. When a request comes in, I first check if the cached file exists and is fresh - if so, I serve it directly from disk. If the cache is stale or missing, I fetch the data from the remote source and update the disk cache. This approach provides persistence across application restarts and reduces network traffic for frequently accessed data. It's especially useful for API responses, static assets, or any data that doesn't change frequently but is expensive to fetch."

---
