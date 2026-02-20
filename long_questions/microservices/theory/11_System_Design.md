# 🟢 **166–180: System Design (Architecture)**

### 166. Design a URL shortener.
"The core requirement is generating a unique, short alias for a long URL and ensuring ultra-fast redirects (`301 Moved Permanently`). 
I'd use a Base62 encoding algorithm (A-Z, a-z, 0-9) because 7 characters yield $62^7$ (3.5 trillion) unique URLs.

To prevent two identical long URLs from generating the same hash, I use a dedicated 'Token Generation Service' (like Apache Zookeeper or a standalone database sequence) that just dispenses mathematically guaranteed unique integer IDs. The application server grabs an ID (e.g., `1234567`), converts it to Base62 (e.g., `aB3dE`), and saves the mapping in a NoSQL database like MongoDB (for fast lookups) or Cassandra (for massive writes).

For sub-millisecond redirect performance, I cache the top 20% most frequently accessed URLs in a globally distributed Redis cluster."

#### Indepth
If a user requests a custom alias (`myapp.com/mycustomname`), the Base62 integer mechanism is bypassed. The system attempts a direct database insert. If it violates a unique index constraint, it rejects the request. The database acts as the ultimate transactional source of truth against race conditions.

---

### 167. Design a chat application.
"A chat app requires persistent, real-time bi-directional communication. HTTP is too slow (due to polling overhead), so I explicitly use WebSockets.

A user's phone opens a persistent WebSocket connection to a 'Chat Gateway' server. Because there are millions of users, there will be hundreds of Gateway servers. I must use a massive clustered Redis Pub/Sub deployment as the central nervous system.

When Alice sends a message to Bob, it hits Gateway A. Gateway A doesn't know where Bob is connected. So Gateway A publishes the message `{"to":"Bob", "msg":"Hi"}` to the central Redis Pub/Sub topic. Gateway D (where Bob happens to be connected) sees the message and instantly pushes it down Bob's open WebSocket."

#### Indepth
To handle offline users, messages must be persisted to a deeply partitioned NoSQL database like Cassandra before they are published to Redis. If the user is offline, the Gateway skips the WebSocket push. When the user eventually reconnects, their device pulls all missed messages from Cassandra chronologically based on a 'last-seen' timestamp.

---

### 168. Design an e-commerce platform.
"An e-commerce platform is fundamentally a suite of decoupled microservices: User Profile, Product Catalog, Shopping Cart, Inventory, and Order Processing.

The 'Product Catalog' is fiercely read-heavy. I load all product data into Elasticsearch for advanced facet filtering (e.g., 'Size: M, Color: Blue') and aggressive CDN caching for images.
The 'Shopping Cart' is fiercely write-heavy but highly volatile. I store active carts entirely in a highly available Redis cluster (with a 7-day TTL) for sub-millisecond responsiveness.

When a user clicks 'Buy', the 'Order Processing' pipeline begins. Crucially, I use the Saga Pattern across Kafka. The Order Service emits an 'Order Created' event. The Inventory Service consumes it, deducts stock, and emits 'Inventory Secured'. The Payment Service consumes that, charges the card, and emits 'Payment Succeeded'. If payment fails, it emits a compensating event, and Inventory gracefully restores the stock."

#### Indepth
The absolute worst bottleneck is Flash Sales (e.g., limited edition sneakers). The Inventory database will crash if 50,000 users click 'Buy' on 10 available shirts simultaneously. A pre-calculated token bucket in Redis is required to aggressively shed 49,990 requests at the edge API Gateway instantly, allowing only 10 lucky requests to ever touch the fragile relational database.

---

### 169. Design a ride-sharing app.
"The architecture pivots on spatial indexing and massive real-time GPS telemetry ingestion.

Millions of active drivers continuously ping their GPS coordinates every 3 seconds. I pipe this firehose of UDP traffic directly into Apache Kafka. A stream-processing engine (like Apache Flink) ingests this and updates a geospatial database (like Redis with GEO commands or PostgreSQL with PostGIS).

When a rider requests a car, the Rider Service queries the spatial database: `SELECT drivers WHERE distance < 3 miles`. The system identifies 5 nearby drivers and pushes the ride offer to their mobile apps via WebSockets. The first driver to accept locks the ride via a distributed transaction, and the others receive a 'Ride no longer available' push notification."

#### Indepth
Calculating the price involves complex ML models predicting ETA, current dynamic traffic flow, and localized surge-pricing multipliers. This computational heaviness is asynchronously calculated utilizing historical data pipelines (Hadoop) and the final materialized view is pushed to a fast cache for the active Rider Service to query instantaneously.

---

### 170. Design a ticket booking system.
"A system like Ticketmaster requires absolute transactional integrity while managing catastrophic, spiky concurrency. 

If Taylor Swift tickets drop at 10:00 AM, five million people will simultaneously target exactly 50,000 specific database rows (seats). I cannot just let them all hit the database.

I implement a massive 'Virtual Waiting Room' utilizing an edge CDN (Cloudflare) to absorb the initial flood. Traffic is dripped slowly into the core API. 
When a user clicks 'Seat A1', I immediately lock that specific seat using an atomic Redis `SETNX` lock with a rigid 5-minute TTL. For 5 minutes, that seat is cryptographically theirs. If they finish payment via Stripe within the window, the reservation is permanently committed to PostgreSQL. If they dawdle, the TTL expires, the lock vanishes, and the seat instantly becomes available for the next fan."

#### Indepth
Data structure choice is crucial. A relational database mapping `Venues -> 1:M -> Sections -> 1:M -> Rows -> 1:M -> Seats` mathematically guarantees integrity but suffers under extreme join pressure. Aggregating the entire seating chart into a pre-computed JSON BLOB stored in Redis guarantees instant UI rendering.

---

### 171. Design a notification system.
"A notification system must intelligently route millions of short messages (SMS, Email, Push) to varying devices without losing data.

The frontend API receives a `SendNotification` payload. It absolutely does not send the email synchronously. It validates the request and tosses it into a highly durable Apache Kafka topic (e.g., `topic-notifications-pending`), immediately returning a 200 OK to the client.

I spin up dedicated consumer groups for each channel: `EmailWorkers`, `SMSWorkers`, and `APNSWorkers` (Apple Push). The `EmailWorkers` pull from the topic, format the HTML templates, and interface with an external provider (like SendGrid). If SendGrid is down, the worker simply pauses consumption or throws the message into a Dead Letter Queue (DLQ) for automatic retry later. No messages are ever lost."

#### Indepth
Users get extremely annoyed by duplicate notifications. Because the consumers (Workers) might experience network crashes right before acknowledging a message to Kafka, the workers strictly verify a unique `notification_hash` against a fast Redis cache before dispatching the payload to SendGrid, guaranteeing mathematical Idempotency.

---

### 172. Design a payment gateway.
"A payment gateway is the ultimate 'CP' (Consistent and Partition Tolerant) system. It must absolutely prioritize mathematical accuracy over blazing speed.

I utilize strictly ACID-compliant relational databases (like PostgreSQL) configured with synchronous replication. No eventual consistency is permitted here.

Communication with bank APIs (Visa, Mastercard) is inherently flaky and slow. Therefore, the core of the gateway is a hyper-robust Async State Machine. When a payment is submitted, the DB row saves as `STATE: PENDING`. Background Quartz schedulers or temporal workflows pick up the pending rows and carefully execute the external HTTP bank calls. If the bank times out, the system safely retries with strictly regulated Exponential Backoff."

#### Indepth
PCI-DSS compliance strictly dictates that credit card Primary Account Numbers (PANs) never touch the application's physical disk or logging statements. The system must immediately convert the PAN into an opaque 'Token' using an isolated, highly-secured internal vault, ensuring developers only ever interact with meaningless tokens (`tok_12345`).

---

### 173. Design a metric monitoring system.
"Like Datadog, a monitoring system must ingest billions of tiny data points per second, compress them, and query them rapidly.

My microservices don't push metrics randomly; they expose a `/metrics` text endpoint. A scraping agent (like Prometheus) pulls these endpoints every 15 seconds. 

The core storage cannot be MySQL; it would utterly choke on billions of inserts. The system requires a specialized Time-Series Database (TSDB) like InfluxDB or VictoriaMetrics. These databases aggressively compress sequential timestamps (e.g., using Delta-of-Delta encoding) so 10,000 metrics take virtually zero disk space. A Grafana frontend then executes rapid read queries against this TSDB to visualize the dashboards."

#### Indepth
Long-term retention requires mathematical 'Downsampling'. Keeping 1-second resolution metrics for 3 years is financially ruinous. After 7 days, a background job aggregates the highly-granular 1-second data into summarized 5-minute averages, deletes the raw data, and archives the dense summaries to cheap AWS S3 cold storage.

---

### 174. Design a distributed cache.
"Similar to Redis, a distributed cache requires blazing speed and horizontal scalability.

Data is stored entirely in physical RAM using sophisticated `HashMap` arrays natively in C or C++. Because one server cannot hold 10 Terabytes of RAM, I must cluster multiple servers together using **Consistent Hashing**. 

Consistent hashing maps all Cache Nodes (Server A, B, C) and all Data Keys (`User1`, `User2`) cleanly onto a virtual ring. This mathematically determines exactly which server holds `User1`. Crucially, if Server B crashes, only the fraction of keys specifically assigned to Server B are remapped to other nodes. The vast majority of the cache remains entirely undisturbed."

#### Indepth
Because RAM is finite, an eviction policy is strictly mandatory. An LRU (Least Recently Used) algorithm utilizes a Doubly-Linked List pointing to the HashMap entries. Whenever a key is read or written, it is ripped out and moved to the 'Head' of the list. When the server hits 100% RAM, the application ruthlessly deletes whatever key is sitting at the 'Tail' of the list.

---

### 175. Design a load balancer.
"A Load Balancer aggressively distributes incoming network traffic across a group of backend servers to prevent any single server from drowning.

I place the Load Balancer (like NGINX or HAProxy) at the edge of the VPC. It maintains a stateful table of all healthy backend IPs. It continuously executes HTTP `GET /health` pings against them. 

When a client hits the LB, it uses an algorithm (like Round-Robin, Least Connections, or IP Hashing) to swiftly select a backend server. In 'Reverse Proxy' mode, the LB opens a distinct TCP connection to the backend server, pulls the HTTP response, and forwards it flawlessly back to the client."

#### Indepth
Layer 4 (L4) Load Balancers operate at the Transport layer, indiscriminately forwarding raw TCP/UDP packets incredibly quickly without peeking inside. Layer 7 (L7) Load Balancers operate at the Application layer; they open the packet, literally read the HTTP headers and URL path, and can intelligently route `/images` traffic specifically to dedicated high-disk-space Image Servers.

---

### 176. Design a search engine.
"A search engine (or a product search feature) requires ultra-fast full-text querying capable of handling typos, synonyms, and stemming (e.g., matching 'running' to 'ran').

Relational databases using `SELECT * WHERE text LIKE '%run%'` execute excruciatingly slow Full Table Scans. I utilize specifically designed Search Engines like Elasticsearch or Apache Solr (built on Apache Lucene).

These engines build an **Inverted Index**. Instead of a row pointing to a sentence, the index breaks the sentence into distinct words (tokens). The index maps the token 'run' directly to an array of Document IDs `[14, 55, 902]`. When a user types 'run', the engine instantly executes a mathematically perfect O(1) array lookup."

#### Indepth
Elasticsearch handles horizontal scaling via Index 'Sharding'. A single massive dataset is forcefully sliced into 5 mathematical Shards distributed across 5 physical servers. When a user executes a search, the Coordinator Node parallelizes the search, asking all 5 servers simultaneously, aggregating the local results, ranking them via TF-IDF (Term Frequency-Inverse Document Frequency), and delivering the final page.

---

### 177. Design a video streaming service.
"A Netflix-style service demands colossal bandwidth and flawless low-latency video playback.

The core architecture completely avoids sending massive 4GB `mp4` files over the internet. Instead, when a video is aggressively uploaded, background ML servers transcode the video into dozens of different resolutions (4K, 1080p, 480p) and physically slice those videos into tiny 5-second segments (called chunks).

If an iPhone user has a weak 3G connection, the player software downloads the raw 480p 5-second chunk. If they switch to fast WiFi, the player smoothly requests the 1080p 5-second chunk for the very next segment. This is called Adaptive Bitrate Streaming (HLS or DASH). All these billions of tiny chunks are aggressively cached on a massive global CDN."

#### Indepth
To manage metadata (Titles, Cast, Watch History), a heavily decentralized Cassandra NoSQL database is utilized for near-infinite horizontal scaling and Active-Active multi-region availability, ensuring that a user's 'Continue Watching' timestamp synchronizes flawlessly whether they open the app in New York or Tokyo.

---

### 178. Design a feed system.
"Designing a Twitter/Instagram social feed requires generating a chronological list of updates strictly tailored to who the user follows.

For massive celebrities natively, waiting until a user opens their app to mathematically calculate their feed (the 'Pull Model') is cripplingly slow. It requires joining the `User` table, `Follows` table, and `Tweets` table across billions of rows.

Instead, I use the 'Push Model' (Fanout-on-Write). When regular User A tweets, a background worker iterates through their 50 followers and proactively injects the raw Tweet ID directly into 50 distinct Redis Lists representing their followers' individual active timelines. When Follower B opens the app, their personalized feed is instantly pre-computed and sitting flawlessly in Redis RAM."

#### Indepth
The Push Model catastrophically breaks for celebrities with 100 million followers (the 'Justin Bieber' problem). A single tweet would generate 100 million Redis writes, locking the entire data center. Hybrid architectures utilize the Push model purely for normal users, but actively intercept celebrity tweets. They are stored centrally, and the timeline is mathematically computed dynamically (Pull Model) only when an active user actually logs in.

---

### 179. Design a real-time leaderboard.
"A gaming leaderboard must aggressively rank millions of active users and flawlessly fetch a specific user's specific global rank instantly.

SQL `ORDER BY score DESC LIMIT 10` becomes atrociously slow as the user table grows, especially if executed aggressively 1,000 times per second.

I completely abandon relational databases for this specific feature and utilize Redis **Sorted Sets** (`ZSET`). A Sorted Set natively maintains its data precisely mathematically sorted in memory via an advanced Skip List data structure. When a user scores a point, `ZINCRBY leaderboard 1 user_id` executes in $O(\log(N))$ time. Fetching the top 10 players (`ZREVRANGE`) or instantly calculating a user's exact integer rank out of 5 million (`ZREVRANK`) executes in less than a single millisecond."

#### Indepth
If the leaderboard must reflect "All-Time" scores, storing millions of users in a single Redis key might eventually exhaust physical RAM boundaries. The architecture evolves into utilizing background batch jobs (Spark) to compute historical all-time ranks daily in a data warehouse, while utilizing Redis purely to calculate fast, highly volatile "Daily" or "Weekly" dynamic leaderboards.

---

### 180. Design an inventory system.
"An inventory system tracks physical item stock precisely across multiple massive warehouses globally while preventing over-selling during unprecedented traffic spikes.

Relying on eventual consistency (NoSQL) is fundamentally unacceptable here; selling an item that no longer exists physically forces an aggressive business refund and completely shatters customer trust. The architecture enforces rigid ACID transactional guarantees utilizing a distributed SQL database (like CockroachDB or AWS Aurora).

When a checkout initiates, the application aggressively executes a `SELECT ... FOR UPDATE` row-level database lock. This physically forces all other concurrent transactions attempting to buy that specific SKU into a strict queue, ensuring mathematically perfect subtraction."

#### Indepth
To avoid massive API latency caused by database row-locking contention during Flash Sales, developers implement edge-caching 'stock estimation'. If Redis estimates 5,000 units remain, it bypasses locks and allows checkout requests aggressively into the Kafka pipeline. Only when Redis drops to 10 units does the system abruptly clamp down, forcing strict synchronous database validation for the final few purchases.
