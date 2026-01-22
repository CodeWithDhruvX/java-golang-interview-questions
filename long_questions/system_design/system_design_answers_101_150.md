## 🔸 Advanced Distributed Systems (Questions 101-110)

### Question 101: How does a distributed hash table (DHT) work?

**Answer:**
A Distributed Hash Table (DHT) provides a lookup service similar to a hash table: (key, value) pairs are stored in the DHT, and any participating node can efficiently retrieve the value associated with a given key.
*   **Mechanism:**
    *   **Consistent Hashing:** Keys and Nodes are mapped to the same ID space (e.g., a ring).
    *   **Key storage:** A key is stored on the node with an ID closest to the key's ID.
    *   **Routing:** Nodes maintain a routing table (e.g., Finger Table in Chord) to find the responsible node in `O(log N)` hops.
*   **Examples:** Kademlia, Chord, Cassandra (uses DHT concepts).

### Question 102: Explain the concept of quorum in distributed systems.

**Answer:**
Quorum is the minimum number of votes that a distributed transaction has to obtain to be allowed to perform an operation in a distributed system.
*   **Formula:** `R + W > N`
    *   `N` = Total replicas.
    *   `R` = Read Quorum (nodes to read from).
    *   `W` = Write Quorum (nodes to write to).
*   **Example:** If `N=3`, `W=2`, `R=2`. A read is guaranteed to see the latest write because the Read set and Write set must overlap by at least one node.

### Question 103: What is vector clock? How is it used in conflict resolution?

**Answer:**
A vector clock is an algorithm for generating a partial ordering of events in a distributed system and detecting causality violations.
*   **Structure:** A list of `[NodeA: 1, NodeB: 2, NodeC: 1]`.
*   **Usage:**
    *   When a node updates data, it increments its own counter in the vector.
    *   **Conflict:** If two objects have vector clocks where neither dominates the other (e.g., `[A:1, B:0]` vs `[A:0, B:1]`), it's a conflict (concurrent writes). The app must resolve this (e.g., merge or last-write-wins).

### Question 104: Compare Raft vs Paxos.

**Answer:**

| Feature | Paxos | Raft |
| :--- | :--- | :--- |
| **Goal** | Consensus. | Consensus + Understandability. |
| **Structure** | Leaderless (basic), or Leader-based (Multi-Paxos). | Strong Leader. |
| **Complexity** | Extremely complex to implement correctly. | Designed to be modular (Leader Election, Log Replication). |
| **Adoption** | Google Spanner, Chubby. | Etcd, Consul, Kubernetes. |

### Question 105: What is gossip protocol?

**Answer:**
A peer-to-peer communication protocol where nodes share information by periodically sending it to a few random neighbors.
*   **Analogy:** Viral spread of a rumor (epidemic protocol).
*   **Usage:**
    *   **Failure Detection:** "I haven't heard from Node A, have you?" (Cassandra).
    *   **Membership:** Discovering new nodes in the cluster.
*   **Benefits:** Highly scalable (O(log N)), robust, no single point of failure.

### Question 106: How do you detect and recover from split brain in distributed systems?

**Answer:**
Split-Brain occurs when a network partition causes two groups of nodes to believe they are the "primary" or "leader."
*   **Detection:** Heartbeat timeouts between regions.
*   **Prevention:**
    *   **Quorum (Majority Vote):** A leader is only valid if it can talk to >50% of nodes. If a partition has <50%, it steps down. (Requires odd number of nodes: 3, 5, 7).
    *   **Fencing Tokens:** Prevent the "zombie" leader from writing to storage by incrementing a lock version.

### Question 107: How to ensure consistency in leader-follower replication?

**Answer:**
1.  **Synchronous Replication:** Leader waits for Follower to acknowledge write before responding to client. (Guarantees consistency, risks latency/availability).
2.  **Semi-Synchronous:** Wait for at least 1 follower to ack.
3.  **Read-Your-Own-Writes (Session Consistency):** Ensure a user reads from the Leader or an up-to-date Follower for their *own* data.

### Question 108: What is lease-based leader election?

**Answer:**
A mechanism where a leader holds a "lease" (time-bound lock) on a shared resource (like a key in Etcd/Zookeeper).
*   **Workflow:**
    1.  Node acquires lease (e.g., for 10 seconds).
    2.  Node becomes Leader.
    3.  Node must renew lease before it expires to stay Leader.
    4.  If lease expires (node dies or network fail), another node can acquire it.
*   **Benefit:** Handles crashes gracefully without complex voting every time.

### Question 109: Explain the difference between synchronous and asynchronous replication.

**Answer:**
*   **Synchronous:** Write is committed to Primary AND Replicas before Success.
    *   *Pros:* Zero data loss (RPO=0).
    *   *Cons:* Slow write; if Replica is down, write fails/hangs.
*   **Asynchronous:** Write is committed to Primary -> Success. Replicated to Replicas in background.
    *   *Pros:* Fast write.
    *   *Cons:* Potential data loss if Primary crashes before replication.

### Question 110: What is two-phase commit protocol?

**Answer:**
A distributed algorithm to coordinate all the processes that participate in a distributed transaction on whether to commit or abort (rollback).
*   **Phase 1 (Prepare):** Coordinator asks all participants: "Can you commit?" Participants lock resources and vote Yes/No.
*   **Phase 2 (Commit):**
    *   If all voted Yes: Coordinator sends "Commit".
    *   If any voted No: Coordinator sends "Abort".
*   **Drawback:** It's a blocking protocol. If Coordinator crashes, participants are stuck holding locks.

---

## 🔸 Advanced Caching & Optimization (Questions 111-120)

### Question 111: How do you design a multi-layer cache system?

**Answer:**
A strategy using different types of caches at different layers to maximize speed and minimize cost.
*   **L1 (Local Cache):** In-memory (RAM) inside the app server (e.g., Caffeine/Guava). Fastest, but small and not shared.
*   **L2 (Distributed Cache):** Remote cluster (Redis/Memcached). Shared across servers, larger size, slightly slower (network call).
*   **L3 (CDN):** Caching static assets at the edge.
*   **Flow:** App checks L1 -> Miss -> Checks L2 -> Miss -> DB -> Updates L2 -> Updates L1.

### Question 112: Explain cache coherence in distributed cache.

**Answer:**
Ensuring that data in different cache replicas reflects the most recent write.
*   **Problem:** If you update DB but fail to update Cache, or update one cache node but not others.
*   **Solutions:**
    *   **Invalidation:** Delete key on write. Next read fetches fresh data.
    *   **Pub/Sub:** Broadcast "invalidation" messages to all local caches when data changes.

### Question 113: What is cache warming?

**Answer:**
The process of pre-populating the cache with data before the system goes live or handles traffic.
*   **Why:** An empty cache means 100% miss rate, causing a "thundering herd" to the DB, potentially crashing it.
*   **How:** Run a script that queries the most popular keys (top 20% of data) to load them into Redis/Memcached.

### Question 114: What is negative caching?

**Answer:**
Restoring a "miss" validation or an error result in the cache.
*   **Scenario:** Client queries for `User_999` (doesn't exist). DB returns "Not Found".
*   **Without Negative Cache:** Every request for `User_999` hits the DB.
*   **With Negative Cache:** Store `Key: User_999, Value: NotFound` with a short TTL (e.g., 5 min).

### Question 115: How does content-based cache invalidation work?

**Answer:**
Invalidation based on the content changing (often using hashing).
*   **Technique (Web):** `style.css?v=hash123`. When file changes, content hash becomes `hash456`.
*   **Browser/CDN:** Sees a new URL, so fetches the new file. No need for explicit "purge" command.

### Question 116: How would you cache personalized data?

**Answer:**
Caching user-specific content (e.g., "My Orders", "Recommended for You").
*   **Challenge:** Low cache hit ratio (User A's data is useless to User B).
*   **Strategy:**
    *   Cache fragments, not whole pages. (Cache the "Product" block, but fetch "Price for User" dynamically).
    *   Use **Edge Side Includes (ESI):** Assemble cached fragments at the CDN/LB.
    *   Store Session/User data in a dedicated fast store (Redis) keyed by UserID.

### Question 117: Design a global caching strategy for a multi-region app.

**Answer:**
1.  **Local Redis:** Each region has its own Redis cluster for fast, region-local access.
2.  **Replication:** Use Active-Active replication (e.g., Redis Enterprise CRDT or DynamoDB Global Tables) to sync write-heavy cache data across regions (with latency).
3.  **CDN:** Cache static content globally.
4.  **Route:** Users route to nearest Region; fetch from local cache. If miss, fetch from local DB (which might be replicated).

### Question 118: How do you handle stale cache data?

**Answer:**
Stale data is inevitable in eventual consistency.
*   **Short TTL:** Reduce the window of staleness.
*   **Versioning:** Store `version` in DB and Cache. Client sends expected version. If mismatch, fetch from DB.
*   **Soft TTL (Graceful):** Store two TTLs.
    *   `Soft`: "Data is old, but serve it anyway while fetching new data in background."
    *   `Hard`: "Data is too dead, return miss."

### Question 119: What is TTL (Time to Live) in caching and how to set it optimally?

**Answer:**
TTL governs how long data stays in cache.
*   **Static Data:** Long TTL (Hours/Days). (e.g., Product Description).
*   **Dynamic Data:** Short TTL (Seconds/Minutes). (e.g., Prices, Stock).
*   **Jitter:** Add random variance to TTL (e.g., 60s ± 5s) to prevent **Cache Stampede** (all keys expiring at once).

### Question 120: How to handle write-heavy cache scenarios?

**Answer:**
If data changes rapidly (e.g., live viewer count), updating cache on every write is expensive.
*   **Write-Back:** Update cache in-memory, flush to DB every N seconds.
*   **Batching:** Aggregate writes (increments) and write once.
*   **Just-in-Time:** Don't write to cache on update. Just delete the key (Invalidate). Next read will repopulate it (less race conditions).

---

## 🔸 Data Flow & Streaming Systems (Questions 121-130)

### Question 121: How does Apache Kafka work? Designing considerations?

**Answer:**
Kafka is a distributed event streaming platform.
*   **Concept:** Log-based storage. Browsers (Producers) write to the end of a log. Servers (Consumers) read from a specific offset.
*   **Structure:**
    *   **Topic:** Feed of messages.
    *   **Partition:** Topic is split into partitions for scalability.
    *   **Broker:** Server holding partitions.
*   **Design Considerations:**
    *   Number of partitions = Max number of parallel consumers.
    *   Replication factor (usually 3) for durability.
    *   Retention policy (Time-based or Size-based cleanup).

### Question 122: How do you ensure message ordering in a distributed system?

**Answer:**
In distributed queues (like Kafka/SQS), global ordering is hard.
*   **Kafka Strategy:** Parallelism is achieved by Partitions. Ordering is guaranteed **only within a Partition**.
*   **Solution:** Use a `Partition Key` (e.g., `UserID`). All events for User A go to Partition 1. Therefore, User A's events are processed in order.
*   **Trade-off:** You cannot have ordering AND parallel consumers for the *same* entity.

### Question 123: How to design a real-time data analytics pipeline?

**Answer:**
1.  **Ingestion:** Kafka / AWS Kinesis (Buffer high-velocity data).
2.  **Processing:**
    *   **Stream Processor:** Apache Flink / Spark Streaming / Kafka Streams.
    *   **Logic:** Windowing (Calculate avg every 1 min), Filtering, Enrichment.
3.  **Storage:**
    *   **Hot Path (Real-time):** Write result to Redis/TimeSeriesDB (InfluxDB) for dashboards.
    *   **Cold Path (History):** Dump raw data to Data Lake (S3/HDFS).
4.  **Visualize:** Grafana / Tableau.

### Question 124: Difference between stream processing vs batch processing.

**Answer:**
*   **Stream (Real-time):**
    *   Process data as it arrives (event-by-event).
    *   Low latency (ms/seconds).
    *   Tool: Flink, Storm, Kafka Streams.
    *   Use Case: Fraud detection, Monitoring.
*   **Batch (Offline):**
    *   Process large volume of accumulated data at intervals (e.g., every night).
    *   High Latency (hours), High Throughput.
    *   Tool: Hadoop MapReduce, Spark.
    *   Use Case: Payroll, Daily reporting.

### Question 125: Design a log aggregation system.

**Answer:**
(See Question 84 for simple version. Detailed design here).
*   **Agent (Sidecar):** Runs on every server (Fluentd/Filebeat). Tails log files.
*   **Buffer:** Kafka. Agents push logs to Kafka (handles bursts).
*   **Indexer:** Logstash reads from Kafka, parses JSON, sends to Elasticsearch.
*   **Storage Refinement:**
    *   Hot Nodes (SSD) for recent 7 days (Fast search).
    *   Warm Nodes (HDD) for 30 days.
    *   Archive (S3) for 1 year (Cheap).

### Question 126: Explain publish-subscribe vs message queue pattern.

**Answer:**
*   **Message Queue (Point-to-Point):**
    *   Sender -> Queue -> **One** Receiver.
    *   If Receiver A reads msg, Receiver B won't seeing it.
    *   Use: Work distribution (Job queues).
*   **Pub/Sub (Broadcast):**
    *   Publisher -> Topic -> **All** Subscribers.
    *   Message is cloned for Subscriber A and Subscriber B.
    *   Use: Notifications, Event updates (Order Placed -> Email Service AND Inventory Service).

### Question 127: How would you handle duplicate messages in streaming systems?

**Answer:**
Duplicates happen due to retries (At-least-once delivery).
*   **Idempotency (Consumer Side):**
    *   Store `MessageID` in a distinct table or Redis.
    *   Before processing, check if `MessageID` exists.
    *   If processed, skip.
*   **Transaction (Producer Side):** Kafka supports "Exactly-Once Semantics" (EOS) using transactional writes, but it's complex and has performance cost.

### Question 128: Design an event-driven architecture.

**Answer:**
*   **Core:** Components trigger events (`UserRegistered`, `PaymentFailed`) rather than calling APIs.
*   **Bus:** Central message broker (Kafka/EventBridge).
*   **Components:**
    *   **Producers:** Emit events.
    *   **Consumers:** React to events.
*   **Benefits:** Decoupling (Producer doesn't know who listens), Scalability.
*   **Challenges:** Debugging (Hard to trace flow), Event Schema evolution.

### Question 129: What is backpressure in streaming systems?

**Answer:**
A mechanism where a consumer signals the producer to slow down because it cannot keep up with the data rate.
*   **Scenario:** Producer sends 1000 msg/sec. Consumer can only process 500 msg/sec. Memory fills up -> Crash.
*   **Backpressure:**
    *   **TCP Flow Control:** Built-in windowing.
    *   **Reactive Streams:** Explicit `request(n)` signals in app code.
    *   **Buffering:** Use Kafka to buffer the surge; Consumer reads at its own pace.

### Question 130: How to scale Kafka consumers?

**Answer:**
*   **Partitioning:** The unit of parallelism in Kafka is the Partition.
    *   If Topic has 10 partitions, you can have up to 10 consumers in a Consumer Group.
    *   Adding 11th consumer does nothing (it sits idle).
*   **Scaling Up:**
    1.  Increase Topic Partitions (e.g., 10 -> 20).
    2.  Add more Consumer instances (up to 20).
*   **Bottleneck:** If a single partition is too slow, optimize the consumer code (multi-threading within one consumer).

---

## 🔸 File & Media Systems (Questions 131-140)

### Question 131: How to design a file deduplication system?

**Answer:**
Storing only one copy of duplicate data.
*   **Client-Side:** Calculate Hash (SHA-256) of file. Ask Server "Do you have this hash?". If yes, skip upload.
*   **Block-Level:** Split file into chunks. Hash each chunk.
    *   Store chunks in Blob Storage (S3).
    *   Store File Meta: `FileA -> [ChunkHash1, ChunkHash2]`.
    *   If File B has same chunks, it just points to existing hashes.

### Question 132: How to store large files efficiently (e.g., videos)?

**Answer:**
*   **Storage:** Object Storage (AWS S3, Google GCS, Azure Blob). Not Database!
*   **Access:** Generate Pre-signed URLs. Client uploads directly to S3 (bypassing App Server) to save bandwidth.
*   **Delivery:** Use CDN.
*   **Archival:** Move old/rarely watched videos to Cold Storage (AWS Glacier) to save cost.

### Question 133: Explain content-addressable storage.

**Answer:**
A storage system where data is retrieved based on its content (Hash) rather than its location (Path).
*   **Example:** IPFS, Git.
*   **Key:** The address of the file is `Hash(FileContent)`.
*   **Prop:** Integrity is guaranteed. If content changes, Hash changes, so Address changes.

### Question 134: Design a secure file sharing system.

**Answer:**
*   **Access Control:** ACLs in DB (`UserA can read FileX`).
*   **Encryption at Rest:** Encrypt file Key (DEK) with Master Key (KEK).
*   **Encryption in Transit:** TLS.
*   **Sharing:**
    *   Generate a unique random Token (UUID).
    *   Link Token to FileID + ExpiryTime.
    *   Share Link: `app.com/share?token=xyz`.

### Question 135: How to implement resumable uploads?

**Answer:**
Essential for large files on unreliable networks.
*   **Protocol:** TUS Protocol is the standard.
*   **Logic:**
    1.  Client splits file into chunks (e.g., 1MB).
    2.  Upload Chunk 1. Server acknowledges `Offset: 1MB`.
    3.  Network fail.
    4.  Client asks "Where was I?". Server says `Offset: 1MB`.
    5.  Client starts uploading Chunk 2.

### Question 136: How to handle concurrent file uploads?

**Answer:**
Multiple users uploading multiple files.
*   **Direct-to-Cloud:** Don't proxy through your backend. Generate Presigned URL -> Client uploads to S3. Parallelizes purely on S3's scale.
*   **Backend Scaling:** If must use backend, use Async I/O (Node.js/Go) to handle thousands of open connections without blocking threads.

### Question 137: How do you store and serve user-generated media?

**Answer:**
*   **Upload:** Direct to S3 (Raw Bucket).
*   **Trigger:** S3 Event triggers Lambda function.
*   **Process:** Sanitize (strip EXIF metadata), Resize (Thumbnail, Medium, Large), Compress (WebP).
*   **Store:** Processed images to S3 (Public Bucket).
*   **Serve:** Via CloudFront (CDN).

### Question 138: What is chunking and how does it help in file uploads?

**Answer:**
Breaking a file into smaller pieces.
*   **Parallelism:** Upload 5 chunks at once (saturate bandwidth).
*   **Resiliency:** If one chunk fails, retry just that chunk, not the whole 10GB file.
*   **Browser:** `File.slice()` API in Javascript.

### Question 139: Explain design of thumbnail generation service.

**Answer:**
*   **Async Worker Pattern:**
    1.  User uploads image. App returns "202 Accepted".
    2.  Push "ImageID" to Queue (SQS).
    3.  Worker (ImageMagick/FFmpeg) pulls message.
    4.  Worker downloads image -> Resizes -> Uploads Thumbnail -> Updates DB.
    5.  Client long-polls or gets WebSocket update: "Thumbnail Ready".

### Question 140: How would you stream a video in real-time?

**Answer:**
(Live Streaming).
*   **Protocol:**
    *   **RTMP:** (Old standard) Ingest from Camera to Server.
    *   **HLS/DASH:** (Modern, Delivery) Server chunks stream into `.ts` files (e.g., 2-second segments) and updates a manifest `.m3u8` file.
*   **Latency:** Standard HLS has 10-30s delay. For low latency (meetings), use **WebRTC** (UDP-based, sub-second latency).

---

## 🔸 Search & Recommendation Engines (Questions 141-150)

### Question 141: How would you design a search autocomplete system?

**Answer:**
*   **Data Structure:** **Trie (Prefix Tree)**.
    *   Root -> 'a' -> 'p' -> 'p' -> 'l' -> 'e'.
    *   Store "Apple" at the last node.
    *   Store "Top 5 popular searches" at each node to return quickly.
*   **Optimization:**
    *   **Typeahead Service:** Keeps Trie in memory (or Redis).
    *   **Pre-computation:** Don't search full DB. Build the Trie offline from logs.

### Question 142: Design a spell checker.

**Answer:**
*   **Edit Distance (Levenshtein):** Calculate how many operations (insert, delete, replace) to turn "helo" into "hello" (Distance=1).
*   **Soundex:** Phonetic matching ("Smith" vs "Smyth").
*   **BK-Tree:** A tree structure optimized for searching metric spaces (like edit distance).
*   **Approach:** Find dictionary words with Edit Distance <= 2 to the input.

### Question 143: How does a recommendation engine work?

**Answer:**
1.  **Collaborative Filtering:** "Users like you also liked..." (User-based) or "People who bought X also bought Y" (Item-based). (Matrix Factorization).
2.  **Content-Based:** "You liked Sci-Fi, here is more Sci-Fi" (Based on tags/metadata).
3.  **Hybrid:** Combine both.
4.  **Deep Learning:** YouTube/Netflix use neural networks considering context (time of day, device).

### Question 144: How to store and index documents for full-text search?

**Answer:**
*   **Inverted Index:** The core structure of Lucene/Elasticsearch.
    *   Map `Word -> List[DocumentIDs]`.
    *   "Burger" -> `[Doc1, Doc5]`
    *   "King" -> `[Doc5, Doc9]`
    *   Query "Burger King" -> Intersection(`[Doc1, Doc5]`, `[Doc5, Doc9]`) -> `Doc5`.

### Question 145: What is Elasticsearch and how does it work?

**Answer:**
A distributed, RESTful search and analytics engine based on Apache Lucene.
*   **Distributed:** Shards and Replicas.
*   **JSON:** Stores schema-less documents.
*   **NRT (Near Real-Time):** Searchable within 1 second of indexing.
*   **Use Case:** Log analytics, Full-text search, Security analysis.

### Question 146: Design a trending hashtag system.

**Answer:**
(Stream Processing problem).
1.  **Ingest:** Twitter Firehose -> Kafka.
2.  **Window:** "Last 1 hour".
3.  **Count:** Stream Processor (Flink) counts hashtags in Sliding Window (e.g., every 5 min, calc count for last 60 min).
4.  **Top-K:** Maintain a Min-Heap of size K (e.g., 10). If count > heap.min, replace.
5.  **Serve:** Query Redis for current Top 10 list.

### Question 147: How do you handle typos in search queries?

**Answer:**
*   **Fuzzy Search:** Elasticsearch supports fuzzy queries using Edit Distance (`fuzziness: AUTO`).
*   **N-Grams:** Break "Google" into "Goo", "oog", "ogl", "gle". Index these. A typo "Gogle" shares many n-grams, so it matches.

### Question 148: How do you personalize search results?

**Answer:**
*   **Re-ranking:**
    1.  Perform generic search -> Get Top 1000 results.
    2.  Apply personalised scoring model (LightGBM/XGBoost).
    *   *Features:* User's location, past clicks, price preference.
    3.  Sort by New Score -> Return Top 10.

### Question 149: How would you design a vector similarity search system?

**Answer:**
Used for semantic search ("Show me red dress" finds images of red dresses without text match).
*   **Concept:** Convert text/image to Vector (Embedding) using ML (BERT/ResNet).
*   **Search:** Find vectors "closest" to query vector (Cosine Similarity in high-dimensional space).
*   **Index:** KNN (K-Nearest Neighbors). Approximations: HNSW (Hierarchical Navigable Small World), FAISS (Facebook AI Similarity Search).
*   **DB:** Pinecone, Milvus, Weaviate, pgvector.

### Question 150: How to optimize search performance in a large dataset?

**Answer:**
1.  **Sharding:** Split index across nodes. Parallelize search.
2.  **Filter First:** Apply cheap filters (Category=Shoes) before expensive text match.
3.  **Index Optimization:** Remove Stop Words ("the", "and"), excessive fields.
4.  **Result Caching:** Cache results for common queries ("iPhone 15").
5.  **Timeout:** Return "Best Effort" results if query takes too long (Partial results > No results).
