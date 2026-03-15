## 🔸 Analytics & Big Data (Questions 151-160)

### Question 151: Design a metrics collection system like Prometheus.

**Answer:**
*   **Architecture:** Pull-based (Scraping).
*   **Components:**
    *   **Scraper:** Periodically polls `/metrics` endpoint of services.
    *   **TSDB (Storage):** Optimized for high write throughput (append-only) and compression (delta-of-delta).
    *   **Rule Evaluator:** Checks alerts (e.g., `CPU > 80%`).
*   **Discovery:** Uses Service Discovery (K8s/Consul) to find targets.
*   **Query:** PromQL engine for aggregation.

### Question 152: How do you store and query billions of rows efficiently?

**Answer:**
*   **Columnar Storage:** (e.g., Cassandra, ClickHouse, Redshift). Stores data by column rather than row. Excellent for aggregating specific fields (e.g., "Sum of Price").
*   **Partitioning:** Split data by Time (Day/Month) or Key.
*   **Indexing:** Bitmap indexes or Bloom Filters to skip irrelevant data blocks.
*   **Compression:** Columnar usage allows high compression ratios (10x).

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you store and query billions of rows efficiently?

**Your Response:** "To store and query billions of rows efficiently, I'd use a columnar storage system, which stores data by column rather than row. This allows for excellent aggregation performance on specific fields. I'd also use partitioning to split the data by time or key, which reduces the amount of data that needs to be scanned. Additionally, I'd use indexing techniques like bitmap indexes or Bloom Filters to skip irrelevant data blocks. Finally, I'd take advantage of columnar storage's high compression ratios to reduce storage costs."

### Question 153: Design a BI dashboard backend.

**Answer:**
*   **Requirement:** Low latency queries on massive data.
*   **Architecture:**
    *   **Data Lake (S3):** Raw data storage.
    *   **ETL:** Spark jobs transform raw data into star schemas.
    *   **Data Warehouse (Snowflake/BigQuery):** Stores structured data.
    *   **OLAP Cube / Materialized Views:** Pre-compute common aggregations (Sales by Region) to serve dashboards instantly.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a BI dashboard backend?

**Your Response:** "A BI dashboard backend needs to provide low-latency queries on massive data. I'd use a data lake to store raw data, and then use ETL jobs to transform the data into star schemas. The data would be stored in a data warehouse, which would provide fast query performance. To further improve performance, I'd use OLAP cubes or materialized views to pre-compute common aggregations, such as sales by region. This would allow the dashboard to serve queries instantly."

### Question 154: What is OLAP vs OLTP?

**Answer:**
*   **OLTP (Online Transaction Processing):**
    *   Focus: Daily operations (Insert, Update-heavy).
    *   Pattern: Row-oriented, ACID, fast point-lookups.
    *   DB: PostgreSQL, MySQL.
*   **OLAP (Online Analytical Processing):**
    *   Focus: Analysis, Reporting (Read-heavy).
    *   Pattern: Column-oriented, complex aggregations.
    *   DB: Redshift, Snowflake, ClickHouse.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is OLAP vs OLTP?

**Your Response:** "OLTP and OLAP serve different purposes in data systems. OLTP, or Online Transaction Processing, is for day-to-day operations like inserting orders, updating user profiles - it's optimized for lots of small, fast transactions with ACID compliance. Think of your e-commerce checkout process. OLAP, or Online Analytical Processing, is for business intelligence and analytics - it's optimized for complex queries and aggregations over large datasets. Think of generating monthly sales reports. OLTP databases like PostgreSQL use row-oriented storage for fast point lookups, while OLAP databases like Snowflake use column-oriented storage for efficient aggregations. Most companies use both - OLTP for the application database and OLAP for the data warehouse."

### Question 155: How would you implement an A/B testing platform?

**Answer:**
*   **Assignment Service:** Maps `UserID` -> `Variant` (Control/Test).
    *   Use Hash: `hash(UserID + ExperimentID) % 100`. (Stateless, deterministic).
*   **Configuration:** Store experiments in DB/Cache (Start/End time, rollout %).
*   **Tracking:** Services log events (`Click`, `Purchase`) tagged with `ExperimentID`.
*   **Analysis:** Data pipeline aggregates metrics (Conversion Rate) for Control vs Test groups.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you implement an A/B testing platform?

**Your Response:** "An A/B testing platform needs to reliably assign users to different variants and track their behavior. I'd use a hash function like `hash(UserID + ExperimentID) % 100` to deterministically assign users to control or test groups - this ensures the same user always gets the same variant. The experiment configurations would be stored in a database with start/end times and rollout percentages. As users interact with the application, I'd log events like clicks and purchases tagged with the experiment ID. Finally, a data pipeline would aggregate these metrics to calculate conversion rates for each variant. The key challenges are ensuring consistent assignment across multiple servers, handling experiments that run simultaneously, and detecting statistical significance in the results."

### Question 156: What is a star schema in data warehousing?

**Answer:**
A modeling technique with a central **Fact Table** surrounded by **Dimension Tables**.
*   **Fact Table:** Contains measurements (Price, Quantity) and Foreign Keys. (e.g., `Sales`).
*   **Dimension Tables:** Contain descriptive attributes (e.g., `Product`, `Time`, `Store`).
*   **Benefit:** Simpler queries (fewer joins than normalized snowflake schema), faster aggregations.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is a star schema in data warehousing?

**Your Response:** "A star schema is a data modeling pattern that's optimized for analytics queries. It has a central fact table containing measurements and foreign keys, surrounded by dimension tables containing descriptive attributes. For example, a sales fact table might contain price and quantity with foreign keys to product, time, and store dimension tables. This structure makes queries much simpler and faster because most analytical queries only need to join the fact table with a few dimension tables, rather than complex multi-level joins in normalized schemas. It's called a star schema because when you draw it, it looks like a star with the fact table in the center and dimension tables radiating outward."

### Question 157: How to collect and store clickstream data?

**Answer:**
*   **Frontend:** JS SDK buffers events, flushes batch to API Gateway.
*   **Ingestion:** API Gateway -> Kafka (Decouples producer/consumer).
*   **Storage (Speed Layer):** Real-time stream to Redis/Druid for live dashboards.
*   **Storage (Batch Layer):** Kafka Connect -> S3 (Parquet format) -> BigQuery for historical analysis.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you collect and store clickstream data?

**Your Response:** "For clickstream data collection, I'd start with a JavaScript SDK on the frontend that buffers events and periodically flushes them in batches to an API Gateway. This reduces the number of network requests. The API Gateway would then push the events to Kafka, which decouples the producers from consumers and handles backpressure. For real-time analytics like live dashboards, I'd stream the data directly to Redis or Druid. For long-term storage and historical analysis, I'd use Kafka Connect to move the data to S3 in Parquet format, then load it into BigQuery. This lambda architecture gives us both real-time and batch processing capabilities."

### Question 158: Design a time-based log analytics tool.

**Answer:**
(e.g., Splunk / ELK)
*   **Ingestion:** high throughput.
*   **Indexing:** Inverted Index (for text search) + BKD Trees (for numeric/time range).
*   **Sharding:** Create a new index every day (`logs-2023-10-01`).
*   **Lifecycle:** Move old indices to cheaper cold storage after 30 days.
*   **Query:** Scatter-gather (query all relevant shards, aggregate results).

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a time-based log analytics tool?

**Your Response:** "For a log analytics tool like Splunk or ELK, I need to handle high-throughput ingestion and fast querying. I'd use an inverted index for text search and BKD trees for numeric and time range queries. To manage the volume, I'd shard the data by creating a new index each day, like `logs-2023-10-01`. This makes time-based queries very efficient. For storage optimization, I'd implement a lifecycle policy that moves older indices to cheaper cold storage after 30 days. When querying, I'd use a scatter-gather approach where the query coordinator sends requests to all relevant shards and aggregates the results. This design scales horizontally and provides fast search capabilities across massive log datasets."

### Question 159: What is the role of data lakes vs warehouses?

**Answer:**
*   **Data Lake (S3, ADLS):**
    *   Stores raw, unstructured data (JSON, CSV, Logs, Images).
    *   Schema-on-Read (Define schema when you query).
    *   Cheap.
*   **Data Warehouse (Snowflake):**
    *   Stores processed, structured data (Tables).
    *   Schema-on-Write (Must fit schema).
    *   Optimized for SQL performance.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is the role of data lakes vs warehouses?

**Your Response:** "Data lakes and warehouses serve different purposes in the data ecosystem. A data lake, like S3 or ADLS, is for storing raw, unstructured data in its original format - JSON files, CSVs, logs, images. It uses schema-on-read, meaning you define the schema when you query the data, which gives you flexibility but requires more processing. Data lakes are cost-effective for storing massive amounts of raw data. In contrast, a data warehouse like Snowflake stores processed, structured data in tables with predefined schemas - this is schema-on-write. The warehouse optimizes for SQL query performance and is ideal for business intelligence and analytics. Most companies use both: the data lake as the raw data repository, and the warehouse as the curated, query-optimized layer for reporting."

### Question 160: How do you build a data pipeline using Spark?

**Answer:**
*   **DAG:** Spark builds a Directed Acyclic Graph of transformations.
*   **Lazy Evaluation:** Operations (map, filter) aren't executed until an Action (count, save) is called.
*   **RDD/DataFrame:** Distributed collection of data.
*   **Stages:** Spark splits job into stages based on "Shuffles" (wide dependencies like `groupBy`).
*   **Cluster Manager:** YARN/K8s allocates resources.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you build a data pipeline using Spark?

**Your Response:** "When building a data pipeline with Spark, I leverage its core architecture. Spark builds a Directed Acyclic Graph of all transformations, which allows for optimization. It uses lazy evaluation - transformations like map and filter don't execute until an action like count or save is called. I work with DataFrames, which are distributed collections of data with a schema. Spark automatically splits the job into stages based on shuffles - operations that require data redistribution like groupBy. The cluster manager, whether YARN or Kubernetes, handles resource allocation. The key benefits are fault tolerance through lineage, in-memory processing for speed, and the ability to process petabytes of data across a cluster."

---

## 🔸 Realtime & Communication Systems (Questions 161-170)

### Question 161: Design a real-time chat system.

**Answer:**
*   **Protocol:** WebSocket (Persistent bi-directional connection).
*   **Server:** Stateful "Chat Server" holds connections.
*   **Routing:** When User A sends msg to User B:
    1.  Look up which Chat Server holds User B's connection (in Redis).
    2.  Forward message to that server via Pub/Sub (Redis).
    3.  Server pushes to User B.
*   **Persistence:** Cassandra/HBase (User A writes to Inbox A; User B writes to Inbox B - "Inbox Search Pattern").

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a real-time chat system?

**Your Response:** "For a real-time chat system, I'd use WebSockets for persistent bi-directional connections between clients and chat servers. Each chat server maintains the state of connected clients. When User A sends a message to User B, the system needs to route it correctly. First, I'd look up in Redis which chat server is holding User B's connection. Then I'd forward the message to that server using Redis pub/sub. Finally, that server pushes the message to User B. For persistence, I'd use Cassandra or HBase with an inbox pattern - each message gets written to both the sender's and receiver's inbox, which makes message retrieval fast and supports offline messaging. This architecture scales horizontally as we can add more chat servers as user count grows."

### Question 162: Design a multiplayer game server backend.

**Answer:**
*   **State Sync:** Use **UDP** (or reliable UDP) for speed. Dropped packets are okay (old coordinates).
*   **Architecture:**
    *   **Game Loop:** Server runs loop (60 ticks/sec), updates physics, broadcasts world state to clients.
    *   **State:** Kept in memory.
*   **Matchmaking:** Separate service that groups players by skill/latency and spins up a Game Server instance.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a multiplayer game server backend?

**Your Response:** "For multiplayer games, performance and low latency are critical. I'd use UDP for state synchronization because speed is more important than reliability - if a coordinate update is dropped, the next one will correct it anyway. The game server would run a tight loop at 60 ticks per second, updating physics and broadcasting the world state to all connected clients. The game state would be kept in memory for fast access. For matchmaking, I'd have a separate service that groups players based on skill level and latency, then spins up dedicated game server instances for each match. This ensures fair matches and isolates games from each other. The architecture scales by adding more game server instances as player count increases."

### Question 163: How would you implement a typing indicator?

**Answer:**
*   **Ephemeral:** No need to store in DB.
*   **Flow:**
    1.  User starts typing. Client sends `TypingStart` event over WebSocket.
    2.  Server broadcast to room participants.
    3.  User stops (or debounce timeout). Client sends `TypingStop`.
*   **Optimization:** Don't send every keystroke. Send event only if `last_sent > 2 seconds` ago.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you implement a typing indicator?

**Your Response:** "For typing indicators, I keep it simple and ephemeral - no database storage needed. When a user starts typing, their client sends a 'TypingStart' event over WebSocket. The server broadcasts this to all participants in the chat room. When the user stops typing or after a debounce timeout, the client sends 'TypingStop'. To optimize network usage, I wouldn't send events for every keystroke. Instead, I'd throttle it to only send if the last event was more than 2 seconds ago. This prevents flooding the network with typing events while still providing a responsive user experience. The indicator automatically disappears after a timeout if no stop event is received."

### Question 164: Design a presence detection system.

**Answer:**
(Online/Offline status).
*   **Heartbeat:** client sends "I'm alive" ping every 30s.
*   **Server:** Updates `LastSeen` timestamp in Redis with TTL=45s.
*   **Query:** If `CurrentTime - LastSeen < 45s`, User is Online.
*   **Disconnect:** On WebSocket close, set status to Offline immediately.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a presence detection system?

**Your Response:** "For presence detection showing online/offline status, I'd use a heartbeat mechanism. The client sends an 'I'm alive' ping every 30 seconds. The server updates a LastSeen timestamp in Redis with a TTL of 45 seconds. To check if a user is online, I simply verify if the current time minus LastSeen is less than 45 seconds. The TTL in Redis automatically cleans up offline users. When a WebSocket connection closes properly, I immediately set the status to offline. This design is efficient because Redis handles the expiration automatically, and it scales well since presence data is small and accessed frequently. The 45-second window accounts for network latency and temporary disconnections."

### Question 165: How do you handle rate control in WebSocket communication?

**Answer:**
*   **Problem:** One user spamming messages prevents others from being heard.
*   **Leaky Bucket:** Associate a bucket with each WebSocket connection.
*   **Enforcement:**
    *   If bucket full, drop message or close connection.
    *   Send warning message to client.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you handle rate control in WebSocket communication?

**Your Response:** "To prevent users from spamming WebSocket messages, I'd implement rate limiting using the leaky bucket algorithm. Each WebSocket connection would have its own bucket that fills up at a constant rate and has a maximum capacity. When a user sends a message, I check if their bucket has tokens available. If the bucket is full, I either drop the message or close the connection after sending a warning. This approach smooths out bursts of messages while allowing normal conversation flow. Unlike fixed window counters, leaky bucket prevents users from sending all their messages at once and then being silent. It's fair to all users and prevents any single user from dominating the conversation."

### Question 166: Design a push notification system.

**Answer:**
*   **Components:**
    *   **Notification Service:** API to accept notifications.
    *   **Queue:** Kafka (prioritize transactional emails vs marketing).
    *   **Dispatcher:** Workers pull from Queue.
    *   **Gateways:** Interface with APNS (Apple) and FCM (Android).
*   **Retry:** Exponential backoff if APNS/FCM is down.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a push notification system?

**Your Response:** "For push notifications, I'd design a multi-component system. The Notification Service provides an API for applications to send notifications. These go into a Kafka queue, which helps handle backpressure and allows prioritization - for example, transactional notifications get higher priority than marketing ones. Dispatcher workers pull messages from the queue and interface with the platform-specific gateways - APNS for Apple devices and FCM for Android. If the gateways are temporarily down, I'd implement exponential backoff retry logic to avoid overwhelming them. The system needs to handle device token management, track delivery status, and respect user preferences. Kafka ensures reliability even if some workers go down."

### Question 167: How would you scale a WebRTC server?

**Answer:**
WebRTC (P2P) doesn't scale well for large groups (Mesh topology: N*N connections).
*   **SFU (Selective Forwarding Unit):** Architecture where Clients send media to Server (SFU). SFU clones and forwards to other participants.
    *   *Client Upload:* 1 stream.
    *   *Client Download:* N-1 streams.
*   **MCU (Multipoint Control Unit):** Server mixes all streams into 1 video. Saves bandwidth but CPU expensive.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you scale a WebRTC server?

**Your Response:** "WebRTC's peer-to-peer nature doesn't scale well for large groups because it creates a mesh topology where each user connects to every other user. For group calls, I'd use either an SFU or MCU approach. With an SFU or Selective Forwarding Unit, clients send their media stream to the server, which then clones and forwards it to other participants. Each client only uploads one stream but downloads N-1 streams. This is more bandwidth-efficient on the client side. Alternatively, an MCU or Multipoint Control Unit would mix all video streams into a single composite video on the server. This saves client bandwidth but is CPU-intensive on the server. For most use cases, SFU is preferred because it offloads processing to clients and scales better horizontally."

### Question 168: Design a collaborative editing system (like Google Docs).

**Answer:**
*   **Challenge:** Concurrency. User A types "Hello", User B types "World" at same index.
*   **Algorithm:** **Operational Transformation (OT)** or **CRDT**.
    *   **OT:** Central server transforms operations (e.g., "Insert at index 5" becomes "Insert at index 10") based on history so everyone converges.
*   **Communication:** WebSocket.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a collaborative editing system like Google Docs?

**Your Response:** "The main challenge in collaborative editing is handling concurrency - when multiple users edit the same document simultaneously. I'd use either Operational Transformation or CRDT algorithms. With OT, a central server transforms operations based on their history. For example, if User A inserts text at index 5 and User B inserts at the same index, the server transforms one of the operations to account for the other's change, ensuring all users see the same final document. Communication would happen over WebSockets for real-time updates. The server maintains the authoritative document state and broadcasts transformed operations to all clients. Each client applies these operations to their local copy, giving the illusion of instantaneous collaboration."

### Question 169: What is Operational Transformation (OT)?

**Answer:**
A family of algorithms for conflict resolution in collaborative editing.
*   **Logic:**
    *   `Op1`: User A inserts 'X' at pos 0.
    *   `Op2`: User B deletes char at pos 0.
    *   If Server receives `Op1` first, it transforms incoming `Op2` to account for the shift caused by `Op1`.
*   **Complexity:** Requires central server to order operations.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is Operational Transformation (OT)?

**Your Response:** "Operational Transformation is a set of algorithms for resolving conflicts in collaborative editing. The core idea is that operations are transformed based on what has happened before them. For example, if User A inserts 'X' at position 0 and User B deletes the character at position 0, the order matters. If the server receives User A's insert first, it transforms User B's delete operation to account for the shift - so instead of deleting position 0, it might delete position 1. This ensures all clients converge to the same document state regardless of the order in which operations are received. OT requires a central server to maintain operation history and order them, which adds complexity but guarantees consistency."

### Question 170: Explain CRDT for real-time collaboration.

**Answer:**
Conflict-free Replicated Data Type.
*   **Advantage:** Decentralized. No central server needed for logic.
*   **Operations:** Commutative (Order doesn't matter). `A + B = B + A`.
*   **Types:**
    *   **LWW-Set:** Last Writer Wins.
    *   **RGA (Replicated Growable Array):** For text editing.
*   **Trade-off:** Memory overhead (store tombstones for deletions).

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain CRDT for real-time collaboration?

**Your Response:** "CRDT or Conflict-free Replicated Data Type is an alternative to OT for collaborative editing. The key advantage is that it's decentralized - operations are commutative, meaning the order doesn't matter. Whether you do operation A then B, or B then A, you get the same result. This eliminates the need for a central server to coordinate operations. For text editing, I'd use an RGA or Replicated Growable Array, which handles insertions and deletions. Other types include LWW-Set where the last writer wins for resolving conflicts. The main trade-off is memory overhead - CRDTs need to store tombstones for deleted items to handle concurrent deletions correctly. This approach is great for offline-first applications and reduces server complexity, though it requires more memory than OT."

---

## 🔸 API Design & Protocols (Questions 171-180)

### Question 171: REST vs gRPC – which one and when?

**Answer:**
*   **REST (JSON/HTTP):**
    *   *Pros:* Human readable, broad support, browser compatible, cacheable.
    *   *Cons:* Verbose (JSON footprint), no strict schema enforcement.
    *   *Use:* Public APIs, Web frontends.
*   **gRPC (Protobuf/HTTP2):**
    *   *Pros:* High performance (binary), strict contract (.proto), streaming support, code generation.
    *   *Cons:* Not browser native.
    *   *Use:* Internal microservices communication.

### How to Explain in Interview (Spoken style format)

**Interviewer:** When would you choose REST vs gRPC?

**Your Response:** "I'd choose REST for public APIs and web frontends because it's human-readable, has broad browser support, and benefits from HTTP caching. JSON is easy to debug and work with. However, for internal microservices communication, I'd prefer gRPC. It uses Protocol Buffers which provide strict contracts and generate type-safe client code. gRPC is binary-encoded, making it much more performant than JSON. It also supports bidirectional streaming which REST doesn't handle well. The trade-off is that gRPC isn't natively supported by browsers, so you need a gateway for external clients. Most companies use both - REST for the external API layer and gRPC for internal service-to-service communication."

### Question 172: How do you version APIs?

**Answer:**
1.  **URI Versioning:** `/api/v1/users`. (Most common, easy caching).
2.  **Header Versioning:** `Accept: application/vnd.myapi.v1+json`. (Cleaner URL, harder to explore in browser).
3.  **Parameter:** `/api/users?version=1`.
*   **Strategy:** Support N and N-1 versions. Deprecate N-2 with warning headers.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you version APIs?

**Your Response:** "For API versioning, I prefer URI versioning like `/api/v1/users` because it's the most common approach and plays well with HTTP caching. It's clear and easy for developers to understand. Header versioning using the Accept header is cleaner from a URL perspective but harder to explore in browsers. Parameter versioning with query strings is another option but can get messy. My strategy is to always support the current version N and the previous version N-1 simultaneously. When version N-2 is ready to be deprecated, I'd send deprecation warning headers to give clients time to migrate. This gradual transition approach minimizes breaking changes while keeping the API maintainable."

### Question 173: How to design an idempotent API?

**Answer:**
*   **Safe Methods:** GET, HEAD, PUT, DELETE should be idempotent by definition.
*   **POST:** Not idempotent naturally.
*   **Solution:** Client generates `Consistency-Key` (UUID) -> Server stores Key -> If Key exists, return cached response.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design an idempotent API?

**Your Response:** "For idempotent APIs, I ensure that safe HTTP methods like GET, HEAD, PUT, and DELETE are naturally idempotent - calling them multiple times has the same effect as calling once. The challenge is with POST, which isn't idempotent by nature. To make POST idempotent, I'd have the client generate a unique consistency key or UUID with each request. The server stores this key and if it sees the same key again, it returns the cached response instead of processing the request again. This is crucial for operations like payment processing where retrying due to network issues shouldn't result in duplicate charges. The key is to store these keys long enough to handle retry scenarios but not so long that they cause memory issues."

### Question 174: How to handle breaking changes in APIs?

**Answer:**
*   **Parallel Versions:** Run V1 and V2 side-by-side.
*   **Evolution:**
    *   Add new fields (non-breaking).
    *   Mark old fields as `@deprecated`.
*   **Communication:** Announce sunset timeline.
*   **Adapter:** Backend handles new logic; API Gateway transforms V1 request to V2 format internally.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you handle breaking changes in APIs?

**Your Response:** "When dealing with breaking changes, I avoid them whenever possible by adding new fields rather than removing or changing existing ones. But when breaking changes are unavoidable, I run parallel versions - V1 and V2 side-by-side. I'd mark old fields as deprecated to give clients notice. Communication is key - I'd announce a clear sunset timeline well in advance. Internally, I'd use an adapter pattern where the backend handles the new logic and an API Gateway transforms V1 requests to V2 format. This allows gradual migration without forcing all clients to update at once. The goal is to give developers ample time to migrate while maintaining backward compatibility."

### Question 175: What is HATEOAS in REST?

**Answer:**
Hypermedia As The Engine Of Application State.
*   **Concept:** API response includes links to valid next actions.
*   **Example:**
    ```json
    {
      "id": 1,
      "status": "pending",
      "links": [
        {"rel": "pay", "href": "/orders/1/pay"},
        {"rel": "cancel", "href": "/orders/1/cancel"}
      ]
    }
    ```
*   **Benefit:** Client doesn't need to hardcode logic (e.g., "Can I cancel logic?").

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is HATEOAS in REST?

**Your Response:** "HATEOAS stands for Hypermedia As The Engine Of Application State. It's a REST constraint where the API response includes links to valid next actions. For example, when you fetch an order, the response doesn't just contain the order data, but also links like 'pay' and 'cancel' with their respective URLs. The benefit is that clients don't need to hardcode business logic about what actions are available. The API itself tells the client what it can do next. This makes the API more flexible - if we add new actions or change URLs, clients automatically discover them through the links. It's like how web pages work with hyperlinks - you discover what you can do by following the links provided."

### Question 176: How do you design pagination in APIs?

**Answer:**
1.  **Offset/Limit:** `?offset=100&limit=10`.
    *   *Cons:* Slow DB query (`OFFSET 1M`), inconsistent if rows added/deleted while paging.
2.  **Cursor-Based (Keyset):** `?after_id=105&limit=10`.
    *   *Pros:* `WHERE id > 105` is fast (uses index), consistent results.
    *   *Cons:* Can't jump to page 10.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you design pagination in APIs?

**Your Response:** "For API pagination, I prefer cursor-based pagination using keyset pagination like `?after_id=105&limit=10`. This uses a WHERE clause like `WHERE id > 105` which is very fast with database indexes and provides consistent results even if data changes while paging. The trade-off is that users can't jump to a specific page number. The alternative is offset/limit pagination, but this becomes slow with large offsets because the database still has to scan through all previous rows. It also provides inconsistent results if rows are added or deleted while paging. For most applications, cursor-based pagination is better because it scales well and handles real-time data gracefully. I'd include both next and previous cursor links in the response for easy navigation."

### Question 177: How do you secure public APIs?

**Answer:**
1.  **API Keys:** Identify the client App.
2.  **OAuth2/OIDC:** Identify the User.
3.  **Rate Limiting:** Per IP/Key.
4.  **HTTPS:** Mandatory.
5.  **Input Validation:** Reject malformed JSON.
6.  **Signature:** Client signs payload with secret (HMAC) to verify integrity (like Stripe Webhooks).

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you secure public APIs?

**Your Response:** "For public API security, I implement multiple layers of protection. First, API keys identify the client application, while OAuth2/OIDC identifies the actual user. I enforce rate limiting per IP address and API key to prevent abuse. HTTPS is mandatory for all communications to prevent eavesdropping. Input validation is crucial - I reject any malformed JSON or unexpected data structures. For critical APIs like webhooks, I'd implement payload signing where the client signs the request with an HMAC using a shared secret, similar to how Stripe webhooks work. This verifies the request wasn't tampered with. The key is defense in depth - multiple security measures so if one fails, others still protect the API."

### Question 178: How to implement rate limiting on APIs?

**Answer:**
*   **Fixed Window:** 100 req/minute. (Burst at minute boundary issues).
*   **Sliding Window:** Smooths out bursts.
*   **Headers:** Return `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` to be polite to clients.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement rate limiting on APIs?

**Your Response:** "For rate limiting, I prefer using a sliding window algorithm rather than fixed window. Fixed windows like 100 requests per minute can cause issues at boundaries - users could send 100 requests at 11:59:59 and another 100 at 12:00:01. Sliding windows smooth out these bursts by tracking requests over a rolling time period. I implement this in Redis for distributed systems. I also include rate limit headers in responses - X-RateLimit-Limit, X-RateLimit-Remaining, and X-RateLimit-Reset. This is polite to clients because they know when they can make their next request and can build better user experiences. For different tiers of users, I might have different rate limits - higher for premium customers, lower for free tier."

### Question 179: Design a webhook system.

**Answer:**
*   **Registration:** Client provides Callback URL + Secret.
*   **Delivery:**
    1.  Event occurs.
    2.  Push to Queue (Job).
    3.  Worker POSTs payload to Callback URL.
*   **Security:** Sign payload (`X-Signature`) using the Secret so client can verify sender.
*   **Reliability:** Retry with exponential backoff if Client returns 500. Disable after N failures.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a webhook system?

**Your Response:** "For a webhook system, I'd start with a registration process where clients provide their callback URL and a secret. When an event occurs in our system, I push it to a queue as a job. Worker processes then POST the payload to the registered callback URLs. For security, I'd sign each payload using the client's secret and include it in an X-Signature header, so clients can verify the webhook actually came from us. Reliability is crucial - if a client returns a 500 error, I'd retry with exponential backoff. After a certain number of failures, I'd disable that webhook and notify the client. The queue ensures we don't lose events even if the webhook endpoints are temporarily down. This design provides reliable, secure event delivery to external systems."

### Question 180: How to validate and document APIs?

**Answer:**
*   **Validation:**
    *   Schema library (Zod, Joi).
    *   OpenAPI (Swagger) Validator.
*   **Documentation:**
    *   **OpenAPI Specification (Swagger):** Standard JSON/YAML describing endpoints.
    *   **Autogen:** Generate docs from code (e.g., Swashbuckle for .NET, SpringDoc for Java).
    *   **Portal:** Developer portal (Postman/Redoc).

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you validate and document APIs?

**Your Response:** "For API validation, I use schema libraries like Zod or Joi to validate incoming requests against predefined schemas. I also implement OpenAPI validators that ensure requests conform to the Swagger specification. For documentation, I use the OpenAPI Specification which provides a standard way to describe APIs using JSON or YAML. The best approach is to autogenerate documentation directly from the code using tools like Swashbuckle for .NET or SpringDoc for Java. This ensures docs always stay in sync with the actual API. Finally, I'd create a developer portal using tools like Postman or Redoc that provides interactive documentation where developers can actually try out the API endpoints. This combination ensures both validation accuracy and comprehensive, up-to-date documentation."

---

## 🔸 Cloud-Native & Kubernetes (Questions 181-190)

### Question 181: Design a multi-tenant SaaS application.

**Answer:**
*   **Database Isolation:**
    1.  **Separate DB:** High isolation, expensive (Physical separation).
    2.  **Separate Schema:** Shared DB, separate schema per tenant (Logical separation).
    3.  **Discriminator Column:** Shared Table with `tenant_id` column. Cheap, risk of data leak (Application level separation).
*   **App:** Shared app instances. Middleware extracts `TenantID` from Subdomain (`tenant1.app.com`) or JWT.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a multi-tenant SaaS application?

**Your Response:** "For multi-tenant SaaS, I need to consider both data isolation and application architecture. For database isolation, I have three options: separate databases per tenant provide the highest isolation but are expensive; separate schemas in a shared database offer logical separation at a lower cost; or using a discriminator column in shared tables is cheapest but risks data leaks if queries aren't properly filtered. For the application layer, I'd use shared app instances with middleware that extracts the TenantID from either the subdomain like 'tenant1.app.com' or from a JWT token. This middleware ensures all database queries include the tenant filter. The key is choosing the right isolation level based on security requirements and cost constraints, while ensuring the application code never accidentally exposes data across tenants."

### Question 182: How do you handle secrets in Kubernetes?

**Answer:**
*   **K8s Secrets:** Base64 encoded (not encrypted by default).
*   **Secure Implementation:**
    *   Enable **Encryption at Rest** in etcd.
    *   **External Store:** Use Vault / AWS Secrets Manager.
    *   **CSI Driver:** Mount secrets from External Store into Pod as volumes.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you handle secrets in Kubernetes?

**Your Response:** "Kubernetes Secrets by default are just Base64 encoded, not encrypted, so anyone with access to etcd can read them. For proper secret management, I'd enable encryption at rest in etcd to encrypt secrets before storing them. Even better, I'd use an external secret store like HashiCorp Vault or AWS Secrets Manager. The CSI driver approach is really elegant - it mounts secrets directly from the external store into pods as volumes, so the secrets never actually exist as K8s objects. This provides better security, audit trails, and integration with existing secret management systems. The key principle is to avoid storing secrets in plain text anywhere in the cluster and use proper secret management tools that provide rotation, auditing, and fine-grained access control."

### Question 183: How would you scale stateful apps in Kubernetes?

**Answer:**
Use **StatefulSets** (not Deployments).
*   **Features:**
    *   Stable Network ID (`web-0`, `web-1`).
    *   Stable Storage (**VolumeClaimTemplates**). Each pod gets its own PVC.
    *   Ordered Deployment/Termination.
*   **Applications:** Databases (Postgres), Queues (Kafka).

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you scale stateful apps in Kubernetes?

**Your Response:** "For stateful applications in Kubernetes, I'd use StatefulSets instead of regular Deployments. StatefulSets provide three key features: stable network identifiers like 'web-0' and 'web-1' that remain consistent even if pods are rescheduled; stable storage using VolumeClaimTemplates so each pod gets its own persistent volume that follows it around; and ordered deployment and termination. This is crucial for applications like databases or Kafka clusters where pod identity and data persistence matter. Unlike Deployments where pods are interchangeable, StatefulSets maintain the identity and state of each pod. When scaling up, new pods are created sequentially, and when scaling down, they're terminated in reverse order, ensuring graceful handling of stateful operations."

### Question 184: What are sidecars in microservices and why use them?

**Answer:**
A helper container running alongside the main application container in the same Pod.
*   **Uses:**
    *   **Proxy:** Envoy (Service Mesh) handles traffic/mTLS.
    *   **Logging:** Fluentd tails logs and pushes to Splunk.
    *   **Config:** Updates config file from remote source.
*   **Why:** Decouples infrastructure logic from business logic.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What are sidecars in microservices and why use them?

**Your Response:** "A sidecar is a helper container that runs alongside the main application container in the same Kubernetes pod. They share the same network namespace and storage, making it easy for them to communicate. Common uses include running a proxy like Envoy for service mesh functionality, a logging agent like Fluentd to collect and ship logs, or a config updater that pulls configuration from remote sources. The main benefit is separation of concerns - the application container focuses purely on business logic while sidecars handle cross-cutting concerns like networking, logging, or configuration. This makes the application cleaner and more portable since infrastructure concerns are moved to reusable sidecar containers."

### Question 185: How do you design stateless services?

**Answer:**
*   **Principle:** Treat servers like cattle, not pets.
*   **Storage:** Don't store session/files on local disk/ram. Use Redis/S3.
*   **Benefit:** Any request can go to any instance. Easy to auto-scale.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you design stateless services?

**Your Response:** "When designing stateless services, I follow the principle of treating servers like cattle rather than pets - they should be disposable and interchangeable. The key is to avoid storing any state on the local disk or in memory. Instead, I use external storage like Redis for session data, S3 for files, and databases for persistent data. This means any request can be handled by any instance, which makes horizontal scaling trivial - I can just add more instances behind a load balancer. If an instance dies, it doesn't matter because the state is elsewhere. This design also simplifies deployments and rolling updates since I don't need to worry about migrating state between versions."

### Question 186: What is a service mesh?

**Answer:**
An infrastructure layer for handling service-to-service communication.
*   **Implementation:** Sidecar proxy (Envoy/Linkerd) injected into every pod.
*   **Features:**
    *   Traffic Management (Canary, Retries).
    *   Security (mTLS).
    *   Observability (Tracing, Metrics).
*   **Examples:** Istio, Linkerd.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is a service mesh?

**Your Response:** "A service mesh is an infrastructure layer that handles all service-to-service communication in a microservices architecture. It's implemented by injecting a sidecar proxy like Envoy or Linkerd into every pod. This proxy intercepts all network traffic and provides capabilities like traffic management for canary deployments and automatic retries, security through mutual TLS for encrypted communication, and observability with distributed tracing and metrics collection. The beauty is that application code doesn't need to change - all these capabilities are provided transparently by the mesh. Popular implementations include Istio and Linkerd. Service meshes are especially valuable in complex microservices environments where managing communication between dozens of services becomes challenging."

### Question 187: How do you monitor containers at scale?

**Answer:**
*   **cAdvisor:** Runs on nodes to collect container stats.
*   **Prometheus:** Scrapes metrics.
*   **DaemonSet:** Deploy logging/monitoring agent on every Node.
*   **Auto-Discovery:** Monitor new pods dynamically via K8s API.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you monitor containers at scale?

**Your Response:** "For monitoring containers at scale, I use a combination of tools. cAdvisor runs on each node to collect container resource statistics. Prometheus then scrapes these metrics from all nodes for centralized monitoring and alerting. I deploy monitoring agents as DaemonSets, which ensures one agent runs on every node in the cluster. This provides complete coverage. The key is auto-discovery - instead of manually configuring monitoring for each pod, the agents dynamically discover new pods through the Kubernetes API and start monitoring them automatically. This approach scales well because adding new nodes or pods doesn't require any configuration changes. The system automatically adapts to the changing cluster topology."

### Question 188: What is the role of Helm in cloud-native deployments?

**Answer:**
Package Manager for Kubernetes.
*   **Chart:** package containing templates (YAMLs) for Deployment, Service, ConfigMap.
*   **Values:** Separate config (`values.yaml`) from structure.
*   **Versioning:** Install `app-v1.0`, rollback to `app-v0.9` easily.

### How to Explain in Interview (Spoken style format)

**Interviewer:** What is the role of Helm in cloud-native deployments?

**Your Response:** "Helm is essentially the package manager for Kubernetes applications. A Helm chart packages all the Kubernetes YAML manifests - Deployments, Services, ConfigMaps - into a reusable template. The key benefit is separating configuration from structure through the values.yaml file. This means I can deploy the same application to different environments just by changing the values file. Helm also provides versioning, so I can install app-v1.0 and easily rollback to app-v0.9 if needed. It simplifies complex deployments by managing dependencies and providing hooks for lifecycle management. Instead of managing dozens of individual YAML files, I work with a single chart that encapsulates the entire application."

### Question 189: What are Kubernetes Operators?

**Answer:**
A method of packaging, deploying, and managing a Kubernetes application.
*   **Concept:** Extends K8s API with **Custom Resource Definitions (CRDs)**.
*   **Logic:** A controller loop that ensures the actual state matches the desired state.
*   **Use:** "I want a Postgres Cluster". Operator handles the complex logic (replication, failover, backup).

### How to Explain in Interview (Spoken style format)

**Interviewer:** What are Kubernetes Operators?

**Your Response:** "Kubernetes Operators are a pattern for packaging, deploying, and managing complex applications on Kubernetes. The concept extends the Kubernetes API with Custom Resource Definitions, so I can create new resource types like 'PostgresCluster'. An Operator is essentially a controller loop that continuously works to make the actual state match the desired state defined in the CRD. For example, instead of manually setting up database replication, failover, and backups, I'd just create a PostgresCluster resource and the Operator handles all that complexity automatically. Operators encode operational knowledge into software, making it easier to manage stateful applications like databases, message queues, or monitoring systems that require deep domain expertise to run properly."

### Question 190: How would you design an auto-scaling mechanism?

**Answer:**
*   **HPA (Horizontal Pod Autoscaler):** Adds Pods based on CPU/Memory/Custom Metrics (Requests per sec).
*   **VPA (Vertical Pod Autoscaler):** Resizes Pod (CPU/RAM) recommendations.
*   **Cluster Autoscaler:** Adds Nodes (VMs) when Pods can't be scheduled (Pending state).

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design an auto-scaling mechanism?

**Your Response:** "For auto-scaling in Kubernetes, I use a multi-layered approach. The Horizontal Pod Autoscaler adds or removes pods based on metrics like CPU, memory usage, or custom metrics like requests per second. This handles scaling within the existing cluster capacity. The Vertical Pod Autoscaler monitors resource usage and recommends optimal CPU and memory settings for pods, helping to right-size applications. Finally, the Cluster Autoscaler adds entire new nodes to the cluster when pods can't be scheduled due to resource constraints. This three-tier approach ensures applications can handle varying loads efficiently - HPA handles immediate scaling needs, VPA optimizes resource utilization, and the Cluster Autoscaler ensures we have enough infrastructure capacity. The key is setting appropriate thresholds and metrics for each layer to prevent thrashing."

---

## 🔸 Miscellaneous System Design Challenges (Questions 191-200)

### Question 191: Design a leaderboard system.

**Answer:**
*   **Requirement:** Real-time ranking by score.
*   **Solution:** **Redis Sorted Set (ZSET)**.
    *   `ZADD leaderboard <score> <user_id>`: O(log N).
    *   `ZREVRANGE leaderboard 0 9`: Get Top 10. O(log N + K).
    *   `ZRANK leaderboard <user_id>`: Get user rank.
*   **Scale:** Partition by `GameID` or range of scores if massive.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a leaderboard system?

**Your Response:** "For a real-time leaderboard, I'd use Redis Sorted Sets which are perfect for ranking data. I can add scores with ZADD which runs in logarithmic time, get the top 10 players with ZREVRANGE, or find a specific user's rank with ZRANK. All operations are very fast even with millions of users. If the system needs to handle massive scale, I'd partition the leaderboards by GameID or by score ranges. The key is that Redis keeps the data sorted in memory, so reads are instant. For persistence, I'd periodically snapshot the data to a database. The challenge is handling updates efficiently - when a user's score changes, I update their score in the sorted set, and Redis automatically maintains the ordering. This design provides real-time rankings with sub-millisecond response times."

### Question 192: Design a notification system.

**Answer:**
(See Question 166).

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a notification system?

**Your Response:** "For a notification system, I'd build on the push notification architecture I mentioned earlier. The system needs to handle different types of notifications - email, SMS, push notifications, and in-app alerts. I'd use a queue-based approach where notification events are pushed to a message queue like Kafka. Worker processes then pull from the queue and route notifications through the appropriate channels. For email, I'd integrate with services like SendGrid; for SMS, Twilio; and for push notifications, APNS and FCM. The key is making it reliable with retry logic and handling failures gracefully. I'd also implement user preferences to respect notification settings and frequency limits to avoid spamming users."

### Question 193: How to build a feature flag system?

**Answer:**
*   **DB:** `FlagID`, `Status` (On/Off), `Rules` (JSON: `UserCountry == 'US'`).
*   **SDK:** Polls rules every minute. Evaluates rules locally (Latency: 0ms).
*   **Audit:** Log changes to flags.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you build a feature flag system?

**Your Response:** "For a feature flag system, I'd store flag configurations in a database with fields like FlagID, Status, and Rules in JSON format. The rules might target specific users based on attributes like country or user type. The key component is the SDK that applications embed. This SDK polls the flag server every minute for the latest rules and caches them locally. When the application needs to check a flag, it evaluates the rules against the current user context entirely on the client side, which means zero latency. I'd also implement comprehensive audit logging to track who changed which flags and when. This design allows for instant feature rollouts, gradual rollouts based on user segments, and quick rollbacks if something goes wrong, all without deploying new code."

### Question 194: Design a dynamic pricing engine.

**Answer:**
(e.g., Uber Surge).
*   **Inputs:** Demand (Users opening app), Supply (Drivers available), History.
*   **Partition:** Geohash (S2 Cell).
*   **Calculation:** Run every 5 mins. `PriceMultiplier = f(Demand / Supply)`.
*   **Consistency:** Price valid for X minutes. User locks price on request.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a dynamic pricing engine?

**Your Response:** "For dynamic pricing like Uber's surge pricing, I'd consider multiple inputs: current demand from users opening the app, available supply of drivers, and historical patterns. I'd partition the geographic area using geohashes or S2 cells to calculate pricing for specific zones. The calculation runs every few minutes, computing a price multiplier based on the demand-supply ratio. To ensure consistency, once a user requests a ride, the price is locked for a certain period even if the multiplier changes. This prevents users from seeing prices fluctuate rapidly. The system needs to handle edge cases like major events or weather conditions that might cause unusual demand spikes. I'd also implement caps to prevent excessive pricing and maintain user trust."

### Question 195: How would you build a plagiarism detection system?

**Answer:**
*   **Fingerprinting:** Winnowing algorithm / localized hashing (Rabin-Karp).
*   **Storage:** Store document hashes in Inverted Index.
*   **Query:** Hash new document -> Find overlapping hashes in DB -> Calculate similarity %.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you build a plagiarism detection system?

**Your Response:** "For plagiarism detection, I'd use fingerprinting techniques like the Winnowing algorithm or localized hashing with Rabin-Karp. These algorithms break documents into smaller chunks and create hashes for each chunk, which serve as fingerprints. I'd store these hashes in an inverted index that maps each hash to the documents containing it. When checking a new document for plagiarism, I'd hash it using the same algorithm, then query the index to find documents with overlapping hashes. The more hashes two documents share, the higher their similarity percentage. This approach is efficient because it reduces the comparison problem from O(N²) to much faster lookups. It also handles cases where plagiarists try to evade detection by making small changes to the text."

### Question 196: Design a subscription billing system.

**Answer:**
*   **State Machine:** Active, PastDue, Canceled, Trialing.
*   **Cron:** Daily job checks "Renews Today".
*   **Payment:** Call Stripe/Gateway.
    *   Success: Extend `NextBillingDate`.
    *   Fail: Enter `Retry` logic (Dunning management).
*   **Idempotency:** Critical.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a subscription billing system?

**Your Response:** "For subscription billing, I'd implement it as a state machine with states like Active, PastDue, Canceled, and Trialing. A daily cron job identifies subscriptions that renew today and initiates payment processing through Stripe or other payment gateways. On successful payment, I extend the NextBillingDate. If payment fails, I enter a retry logic with dunning management - trying again after a few days, then a week, and eventually canceling if payment never succeeds. Idempotency is critical here to avoid charging customers multiple times for the same period. I'd also handle edge cases like upgrades, downgrades, and prorations. The system needs to be reliable because billing errors directly impact revenue and customer trust."

### Question 197: Design a Q&A platform (like Stack Overflow).

**Answer:**
*   **Search:** Elasticsearch.
*   **DB:** SQL (Strong consistency for votes/integrity).
*   **Tags:** Many-to-Many relationship.
*   **Reputation:** Event-driven. Upvote Event -> Async worker increments `Reputation`.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a Q&A platform like Stack Overflow?

**Your Response:** "For a Q&A platform, I'd use Elasticsearch for search functionality because it provides powerful full-text search and relevance scoring. The primary database would be SQL like PostgreSQL because I need strong consistency for votes, answers, and maintaining data integrity. I'd implement tags as a many-to-many relationship between questions and tags for categorization. For reputation systems, I'd use an event-driven approach - when someone upvotes, it generates an event that an async worker processes to increment the user's reputation. This decouples reputation calculation from the main request flow. The challenge is handling high read loads, so I'd cache popular questions and use read replicas. For voting, I need to prevent duplicate votes and ensure vote counts are accurate."

### Question 198: Design a crowdfunding platform.

**Answer:**
*   **Transaction:** "All or Nothing".
*   **Pledge:** Auth hold on card (don't charge yet).
*   **Completion:** If Goal met by Deadline -> Capture all charges. If not -> Release holds.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a crowdfunding platform?

**Your Response:** "For crowdfunding, I'd implement an 'all or nothing' model where projects only get funded if they reach their goal by the deadline. When someone pledges, I place an authorization hold on their card but don't charge immediately. This reserves the funds without actually taking money. If the project meets its goal by the deadline, I capture all the charges simultaneously. If the goal isn't met, I release all the authorization holds so backers aren't charged. This approach requires careful timing and coordination with payment processors. I'd also need to handle edge cases like cards expiring between the hold and capture, or insufficient funds at capture time. The system must be reliable because it deals with people's money."

### Question 199: Design a loyalty and rewards engine.

**Answer:**
*   **Events:** "Purchase Created".
*   **Rules Engine:** `If Amount > $100 AND Category == 'Shoes', Give 50 Points`.
*   **Balance:** Transactional Log (`+50`, `-20`). Sum to get Balance.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a loyalty and rewards engine?

**Your Response:** "For a loyalty system, I'd use an event-driven architecture. When events like 'Purchase Created' occur, they trigger the rules engine. The rules engine evaluates conditions like 'If Amount > $100 AND Category == Shoes, Give 50 Points'. I'd store point balances as a transactional log with entries like '+50' for points earned and '-20' for points redeemed. The current balance is simply the sum of all transactions. This approach provides full auditability and makes it easy to reconstruct balances if needed. The rules engine needs to be flexible to support different promotions, expiration policies, and tier calculations. I'd also implement batch processing for some rules to avoid slowing down the main purchase flow."

### Question 200: Design a privacy settings system for social media.

**Answer:**
*   **Model:** Access Control List (ACL).
*   **Defaults:** Public, Friends Only, Private.
*   **Check:** At Read Time. `CanUserSeePost(Viewer, Post)`.
    *   Check Post privacy level.
    *   Check relationship (Friend?).
    *   Check Allow/Block lists.

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you design a privacy settings system for social media?

**Your Response:** "For privacy settings, I'd implement an Access Control List model. Users can set default privacy levels like Public, Friends Only, or Private for their content. The privacy check happens at read time through a function like CanUserSeePost(Viewer, Post). This function checks multiple conditions: the post's privacy level, the relationship between viewer and post owner (are they friends?), and any specific allow or block lists. The key is performing these checks efficiently since every content view requires a privacy check. I'd cache frequently accessed relationships and privacy settings to improve performance. The system needs to handle edge cases like tagged photos, shared content, and inherited privacy settings from groups or pages."
