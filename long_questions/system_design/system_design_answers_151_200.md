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

### Question 153: Design a BI dashboard backend.

**Answer:**
*   **Requirement:** Low latency queries on massive data.
*   **Architecture:**
    *   **Data Lake (S3):** Raw data storage.
    *   **ETL:** Spark jobs transform raw data into star schemas.
    *   **Data Warehouse (Snowflake/BigQuery):** Stores structured data.
    *   **OLAP Cube / Materialized Views:** Pre-compute common aggregations (Sales by Region) to serve dashboards instantly.

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

### Question 155: How would you implement an A/B testing platform?

**Answer:**
*   **Assignment Service:** Maps `UserID` -> `Variant` (Control/Test).
    *   Use Hash: `hash(UserID + ExperimentID) % 100`. (Stateless, deterministic).
*   **Configuration:** Store experiments in DB/Cache (Start/End time, rollout %).
*   **Tracking:** Services log events (`Click`, `Purchase`) tagged with `ExperimentID`.
*   **Analysis:** Data pipeline aggregates metrics (Conversion Rate) for Control vs Test groups.

### Question 156: What is a star schema in data warehousing?

**Answer:**
A modeling technique with a central **Fact Table** surrounded by **Dimension Tables**.
*   **Fact Table:** Contains measurements (Price, Quantity) and Foreign Keys. (e.g., `Sales`).
*   **Dimension Tables:** Contain descriptive attributes (e.g., `Product`, `Time`, `Store`).
*   **Benefit:** Simpler queries (fewer joins than normalized snowflake schema), faster aggregations.

### Question 157: How to collect and store clickstream data?

**Answer:**
*   **Frontend:** JS SDK buffers events, flushes batch to API Gateway.
*   **Ingestion:** API Gateway -> Kafka (Decouples producer/consumer).
*   **Storage (Speed Layer):** Real-time stream to Redis/Druid for live dashboards.
*   **Storage (Batch Layer):** Kafka Connect -> S3 (Parquet format) -> BigQuery for historical analysis.

### Question 158: Design a time-based log analytics tool.

**Answer:**
(e.g., Splunk / ELK)
*   **Ingestion:** high throughput.
*   **Indexing:** Inverted Index (for text search) + BKD Trees (for numeric/time range).
*   **Sharding:** Create a new index every day (`logs-2023-10-01`).
*   **Lifecycle:** Move old indices to cheaper cold storage after 30 days.
*   **Query:** Scatter-gather (query all relevant shards, aggregate results).

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

### Question 160: How do you build a data pipeline using Spark?

**Answer:**
*   **DAG:** Spark builds a Directed Acyclic Graph of transformations.
*   **Lazy Evaluation:** Operations (map, filter) aren't executed until an Action (count, save) is called.
*   **RDD/DataFrame:** Distributed collection of data.
*   **Stages:** Spark splits job into stages based on "Shuffles" (wide dependencies like `groupBy`).
*   **Cluster Manager:** YARN/K8s allocates resources.

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

### Question 162: Design a multiplayer game server backend.

**Answer:**
*   **State Sync:** Use **UDP** (or reliable UDP) for speed. Dropped packets are okay (old coordinates).
*   **Architecture:**
    *   **Game Loop:** Server runs loop (60 ticks/sec), updates physics, broadcasts world state to clients.
    *   **State:** Kept in memory.
*   **Matchmaking:** Separate service that groups players by skill/latency and spins up a Game Server instance.

### Question 163: How would you implement a typing indicator?

**Answer:**
*   **Ephemeral:** No need to store in DB.
*   **Flow:**
    1.  User starts typing. Client sends `TypingStart` event over WebSocket.
    2.  Server broadcast to room participants.
    3.  User stops (or debounce timeout). Client sends `TypingStop`.
*   **Optimization:** Don't send every keystroke. Send event only if `last_sent > 2 seconds` ago.

### Question 164: Design a presence detection system.

**Answer:**
(Online/Offline status).
*   **Heartbeat:** client sends "I'm alive" ping every 30s.
*   **Server:** Updates `LastSeen` timestamp in Redis with TTL=45s.
*   **Query:** If `CurrentTime - LastSeen < 45s`, User is Online.
*   **Disconnect:** On WebSocket close, set status to Offline immediately.

### Question 165: How do you handle rate control in WebSocket communication?

**Answer:**
*   **Problem:** One user spamming messages prevents others from being heard.
*   **Leaky Bucket:** Associate a bucket with each WebSocket connection.
*   **Enforcement:**
    *   If bucket full, drop message or close connection.
    *   Send warning message to client.

### Question 166: Design a push notification system.

**Answer:**
*   **Components:**
    *   **Notification Service:** API to accept notifications.
    *   **Queue:** Kafka (prioritize transactional emails vs marketing).
    *   **Dispatcher:** Workers pull from Queue.
    *   **Gateways:** Interface with APNS (Apple) and FCM (Android).
*   **Retry:** Exponential backoff if APNS/FCM is down.

### Question 167: How would you scale a WebRTC server?

**Answer:**
WebRTC (P2P) doesn't scale well for large groups (Mesh topology: N*N connections).
*   **SFU (Selective Forwarding Unit):** Architecture where Clients send media to Server (SFU). SFU clones and forwards to other participants.
    *   *Client Upload:* 1 stream.
    *   *Client Download:* N-1 streams.
*   **MCU (Multipoint Control Unit):** Server mixes all streams into 1 video. Saves bandwidth but CPU expensive.

### Question 168: Design a collaborative editing system (like Google Docs).

**Answer:**
*   **Challenge:** Concurrency. User A types "Hello", User B types "World" at same index.
*   **Algorithm:** **Operational Transformation (OT)** or **CRDT**.
    *   **OT:** Central server transforms operations (e.g., "Insert at index 5" becomes "Insert at index 10") based on history so everyone converges.
*   **Communication:** WebSocket.

### Question 169: What is Operational Transformation (OT)?

**Answer:**
A family of algorithms for conflict resolution in collaborative editing.
*   **Logic:**
    *   `Op1`: User A inserts 'X' at pos 0.
    *   `Op2`: User B deletes char at pos 0.
    *   If Server receives `Op1` first, it transforms incoming `Op2` to account for the shift caused by `Op1`.
*   **Complexity:** Requires central server to order operations.

### Question 170: Explain CRDT for real-time collaboration.

**Answer:**
Conflict-free Replicated Data Type.
*   **Advantage:** Decentralized. No central server needed for logic.
*   **Operations:** Commutative (Order doesn't matter). `A + B = B + A`.
*   **Types:**
    *   **LWW-Set:** Last Writer Wins.
    *   **RGA (Replicated Growable Array):** For text editing.
*   **Trade-off:** Memory overhead (store tombstones for deletions).

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

### Question 172: How do you version APIs?

**Answer:**
1.  **URI Versioning:** `/api/v1/users`. (Most common, easy caching).
2.  **Header Versioning:** `Accept: application/vnd.myapi.v1+json`. (Cleaner URL, harder to explore in browser).
3.  **Parameter:** `/api/users?version=1`.
*   **Strategy:** Support N and N-1 versions. Deprecate N-2 with warning headers.

### Question 173: How to design an idempotent API?

**Answer:**
*   **Safe Methods:** GET, HEAD, PUT, DELETE should be idempotent by definition.
*   **POST:** Not idempotent naturally.
*   **Solution:** Client generates `Consistency-Key` (UUID) -> Server stores Key -> If Key exists, return cached response.

### Question 174: How to handle breaking changes in APIs?

**Answer:**
*   **Parallel Versions:** Run V1 and V2 side-by-side.
*   **Evolution:**
    *   Add new fields (non-breaking).
    *   Mark old fields as `@deprecated`.
*   **Communication:** Announce sunset timeline.
*   **Adapter:** Backend handles new logic; API Gateway transforms V1 request to V2 format internally.

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

### Question 176: How do you design pagination in APIs?

**Answer:**
1.  **Offset/Limit:** `?offset=100&limit=10`.
    *   *Cons:* Slow DB query (`OFFSET 1M`), inconsistent if rows added/deleted while paging.
2.  **Cursor-Based (Keyset):** `?after_id=105&limit=10`.
    *   *Pros:* `WHERE id > 105` is fast (uses index), consistent results.
    *   *Cons:* Can't jump to page 10.

### Question 177: How do you secure public APIs?

**Answer:**
1.  **API Keys:** Identify the client App.
2.  **OAuth2/OIDC:** Identify the User.
3.  **Rate Limiting:** Per IP/Key.
4.  **HTTPS:** Mandatory.
5.  **Input Validation:** Reject malformed JSON.
6.  **Signature:** Client signs payload with secret (HMAC) to verify integrity (like Stripe Webhooks).

### Question 178: How to implement rate limiting on APIs?

**Answer:**
*   **Fixed Window:** 100 req/minute. (Burst at minute boundary issues).
*   **Sliding Window:** Smooths out bursts.
*   **Headers:** Return `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` to be polite to clients.

### Question 179: Design a webhook system.

**Answer:**
*   **Registration:** Client provides Callback URL + Secret.
*   **Delivery:**
    1.  Event occurs.
    2.  Push to Queue (Job).
    3.  Worker POSTs payload to Callback URL.
*   **Security:** Sign payload (`X-Signature`) using the Secret so client can verify sender.
*   **Reliability:** Retry with exponential backoff if Client returns 500. Disable after N failures.

### Question 180: How to validate and document APIs?

**Answer:**
*   **Validation:**
    *   Schema library (Zod, Joi).
    *   OpenAPI (Swagger) Validator.
*   **Documentation:**
    *   **OpenAPI Specification (Swagger):** Standard JSON/YAML describing endpoints.
    *   **Autogen:** Generate docs from code (e.g., Swashbuckle for .NET, SpringDoc for Java).
    *   **Portal:** Developer portal (Postman/Redoc).

---

## 🔸 Cloud-Native & Kubernetes (Questions 181-190)

### Question 181: Design a multi-tenant SaaS application.

**Answer:**
*   **Database Isolation:**
    1.  **Separate DB:** High isolation, expensive (Physical separation).
    2.  **Separate Schema:** Shared DB, separate schema per tenant (Logical separation).
    3.  **Discriminator Column:** Shared Table with `tenant_id` column. Cheap, risk of data leak (Application level separation).
*   **App:** Shared app instances. Middleware extracts `TenantID` from Subdomain (`tenant1.app.com`) or JWT.

### Question 182: How do you handle secrets in Kubernetes?

**Answer:**
*   **K8s Secrets:** Base64 encoded (not encrypted by default).
*   **Secure Implementation:**
    *   Enable **Encryption at Rest** in etcd.
    *   **External Store:** Use Vault / AWS Secrets Manager.
    *   **CSI Driver:** Mount secrets from External Store into Pod as volumes.

### Question 183: How would you scale stateful apps in Kubernetes?

**Answer:**
Use **StatefulSets** (not Deployments).
*   **Features:**
    *   Stable Network ID (`web-0`, `web-1`).
    *   Stable Storage (**VolumeClaimTemplates**). Each pod gets its own PVC.
    *   Ordered Deployment/Termination.
*   **Applications:** Databases (Postgres), Queues (Kafka).

### Question 184: What are sidecars in microservices and why use them?

**Answer:**
A helper container running alongside the main application container in the same Pod.
*   **Uses:**
    *   **Proxy:** Envoy (Service Mesh) handles traffic/mTLS.
    *   **Logging:** Fluentd tails logs and pushes to Splunk.
    *   **Config:** Updates config file from remote source.
*   **Why:** Decouples infrastructure logic from business logic.

### Question 185: How do you design stateless services?

**Answer:**
*   **Principle:** Treat servers like cattle, not pets.
*   **Storage:** Don't store session/files on local disk/ram. Use Redis/S3.
*   **Benefit:** Any request can go to any instance. Easy to auto-scale.

### Question 186: What is a service mesh?

**Answer:**
An infrastructure layer for handling service-to-service communication.
*   **Implementation:** Sidecar proxy (Envoy/Linkerd) injected into every pod.
*   **Features:**
    *   Traffic Management (Canary, Retries).
    *   Security (mTLS).
    *   Observability (Tracing, Metrics).
*   **Examples:** Istio, Linkerd.

### Question 187: How do you monitor containers at scale?

**Answer:**
*   **cAdvisor:** Runs on nodes to collect container stats.
*   **Prometheus:** Scrapes metrics.
*   **DaemonSet:** Deploy logging/monitoring agent on every Node.
*   **Auto-Discovery:** Monitor new pods dynamically via K8s API.

### Question 188: What is the role of Helm in cloud-native deployments?

**Answer:**
Package Manager for Kubernetes.
*   **Chart:** package containing templates (YAMLs) for Deployment, Service, ConfigMap.
*   **Values:** Separate config (`values.yaml`) from structure.
*   **Versioning:** Install `app-v1.0`, rollback to `app-v0.9` easily.

### Question 189: What are Kubernetes Operators?

**Answer:**
A method of packaging, deploying, and managing a Kubernetes application.
*   **Concept:** Extends K8s API with **Custom Resource Definitions (CRDs)**.
*   **Logic:** A controller loop that ensures the actual state matches the desired state.
*   **Use:** "I want a Postgres Cluster". Operator handles the complex logic (replication, failover, backup).

### Question 190: How would you design an auto-scaling mechanism?

**Answer:**
*   **HPA (Horizontal Pod Autoscaler):** Adds Pods based on CPU/Memory/Custom Metrics (Requests per sec).
*   **VPA (Vertical Pod Autoscaler):** Resizes Pod (CPU/RAM) recommendations.
*   **Cluster Autoscaler:** Adds Nodes (VMs) when Pods can't be scheduled (Pending state).

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

### Question 192: Design a notification system.

**Answer:**
(See Question 166).

### Question 193: How to build a feature flag system?

**Answer:**
*   **DB:** `FlagID`, `Status` (On/Off), `Rules` (JSON: `UserCountry == 'US'`).
*   **SDK:** Polls rules every minute. Evaluates rules locally (Latency: 0ms).
*   **Audit:** Log changes to flags.

### Question 194: Design a dynamic pricing engine.

**Answer:**
(e.g., Uber Surge).
*   **Inputs:** Demand (Users opening app), Supply (Drivers available), History.
*   **Partition:** Geohash (S2 Cell).
*   **Calculation:** Run every 5 mins. `PriceMultiplier = f(Demand / Supply)`.
*   **Consistency:** Price valid for X minutes. User locks price on request.

### Question 195: How would you build a plagiarism detection system?

**Answer:**
*   **Fingerprinting:** Winnowing algorithm / localized hashing (Rabin-Karp).
*   **Storage:** Store document hashes in Inverted Index.
*   **Query:** Hash new document -> Find overlapping hashes in DB -> Calculate similarity %.

### Question 196: Design a subscription billing system.

**Answer:**
*   **State Machine:** Active, PastDue, Canceled, Trialing.
*   **Cron:** Daily job checks "Renews Today".
*   **Payment:** Call Stripe/Gateway.
    *   Success: Extend `NextBillingDate`.
    *   Fail: Enter `Retry` logic (Dunning management).
*   **Idempotency:** Critical.

### Question 197: Design a Q&A platform (like Stack Overflow).

**Answer:**
*   **Search:** Elasticsearch.
*   **DB:** SQL (Strong consistency for votes/integrity).
*   **Tags:** Many-to-Many relationship.
*   **Reputation:** Event-driven. Upvote Event -> Async worker increments `Reputation`.

### Question 198: Design a crowdfunding platform.

**Answer:**
*   **Transaction:** "All or Nothing".
*   **Pledge:** Auth hold on card (don't charge yet).
*   **Completion:** If Goal met by Deadline -> Capture all charges. If not -> Release holds.

### Question 199: Design a loyalty and rewards engine.

**Answer:**
*   **Events:** "Purchase Created".
*   **Rules Engine:** `If Amount > $100 AND Category == 'Shoes', Give 50 Points`.
*   **Balance:** Transactional Log (`+50`, `-20`). Sum to get Balance.

### Question 200: Design a privacy settings system for social media.

**Answer:**
*   **Model:** Access Control List (ACL).
*   **Defaults:** Public, Friends Only, Private.
*   **Check:** At Read Time. `CanUserSeePost(Viewer, Post)`.
    *   Check Post privacy level.
    *   Check relationship (Friend?).
    *   Check Allow/Block lists.
