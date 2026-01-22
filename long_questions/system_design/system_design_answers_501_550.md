## 🔸 Machine Learning & AI System Design (Questions 501-510)

### Question 501: How would you design a recommendation engine?

**Answer:**
*   **Data Collection:** User logs (clicks, views) -> Kafka -> Data Lake.
*   **Candidate Generation (Recall):** Fast filtering.
    *   Collaborative Filtering (Matrix Factorization).
    *   Content-Based (Tags).
    *   Selects Top 1000 items from 1M.
*   **Ranking:** Slow, precise.
    *   GBDT / Deep Learning Model (e.g., Wide & Deep).
    *   Scores items based on user probability to click. Selects Top 10.
*   **Serving:** Returns ranked list.

### Question 502: Design a real-time fraud detection system using ML.

**Answer:**
(See Q237 & Q356).
*   **Feature Store:** Real-time features (Velocity: `Trans_Count_Last_Hr`) served from Redis.
*   **Inference:**
    *   Call Model Service (GBDT) with features.
    *   Latency < 200ms.
*   **Fallback:** If Model times out, fallback to Rule Engine.

### Question 503: How to deploy and scale machine learning models in production?

**Answer:**
*   **Serving Platforms:** TensorFlow Serving, TorchServe, Triton Inference Server.
*   **Strategy:**
    *   **CPU:** Good for decision trees/simple models.
    *   **GPU:** Required for Deep Learning / LLMs.
*   **Batching:** Server batches requests (wait 5ms to group 10 requests) to maximize GPU throughput.
*   **Scaling:** KPA (Knative Pod Autoscaler) based on `Concurrency` or `GPU_Utilization`.

### Question 504: How do you handle model versioning in a large system?

**Answer:**
*   **Model Registry:** (MLflow / WandB).
    *   `v1.0` -> `s3://models/fraud/v1.pkl`
    *   `v1.1` -> `s3://models/fraud/v1.1.pkl` (Staged)
*   **Config:** Service loads version from Config Server.
*   **Rollback:** Instant config change points back to `v1.0`.

### Question 505: Design an A/B testing platform for ML models.

**Answer:**
(See Q235).
*   **Assignment:** Hash UserID to Bucket (0-100).
*   **Routing:**
    *   Bucket 0-50 -> Model A (Control).
    *   Bucket 51-100 -> Model B (Variant).
*   **Evaluation:** Compare CTR / Conversion Rate metrics after 1 week.

### Question 506: Build a system for real-time personalized content.

**Answer:**
*   **Profile:** Updates in Real-time (User clicked "Sports").
*   **Process:**
    1.  User Click -> API -> Kafka -> Flink -> Updates Redis User Profile.
    2.  Next Refresh -> Rec Engine reads updated Profile -> Returns related news.

### Question 507: How would you design a feature store?

**Answer:**
(See Q239).
*   **Offline Store:** (Glue/BigQuery). Historical features for training.
*   **Online Store:** (Redis/DynamoDB). Latest feature values for inference.
*   **Sync:** Job ensures Online store matches Offline definition.

### Question 508: How do you manage data drift in ML systems?

**Answer:**
(See Q232).
*   **Schema Validation:** Ensure input range (Age 0-100) hasn't shifted (e.g., suddenly receiving -1).
*   **Distribution Check:** KS-Test (Komogorov-Smirnov) comparing Training set distribution vs Live window.

### Question 509: Build a self-learning chatbot backend.

**Answer:**
1.  **Conversation:** User talks. Model replies.
2.  **Feedback:** User clicks "Thumbs Up/Down" or "Solved".
3.  **RLHF (Reinforcement Learning from Human Feedback):**
    *   Store `(Prompt, Reply, Reward)`.
    *   Fine-tune model periodically using PPO (Proximal Policy Optimization).

### Question 510: How do you log and monitor ML inference performance?

**Answer:**
*   **Metrics:** Latency, Throughput.
*   **Model Specific:**
    *   **Confidence Distribution:** If model usually is 90% confident, but drops to 50%, alert.
    *   **Prediction Drift:** Is the model predicting "Fraud" 50% of the time (vs normal 1%)?

---

## 🔸 Edge Computing & IoT (Questions 511-520)

### Question 511: Design a smart traffic light control system.

**Answer:**
*   **Sensors:** Inductive loops / Cameras at intersection.
*   **Edge:** Controller at the light runs logic.
    *   "If ambulance detected (Siren/Visual) -> Green".
    *   "If Car waiting > 2 mins -> Green".
*   **Central:** Syncs timing across the city for "Green Wave" optimization.

### Question 512: How would you build a fleet tracking system using GPS devices?

**Answer:**
(See Q259).
*   **Optimization:** Map Matching. GPS is noisy. Snap the raw point to the nearest road segment graph.

### Question 513: Build a system to sync IoT data from millions of sensors.

**Answer:**
*   **Protocol:** MQTT over WebSockets.
*   **Broker:** HiveMQ / AWS IoT Core. Supports millions of concurrent connections.
*   **Backpressure:** Devices buffer data if Cloud is slow.

### Question 514: Design a firmware update system for IoT devices.

**Answer:**
(See Q253).

### Question 515: How would you handle data validation on the edge?

**Answer:**
*   **Schema Check:** Reject malformed JSON locally. Save bandwidth.
*   **Sanity Check:** If `Temp > 1000C`, mark as `SensorError` locally. Don't trigger Cloud Alarm for fire.

### Question 516: Design a real-time fire alert system using sensors.

**Answer:**
*   **Priority:** Life Safety.
*   **Path 1 (Local):** Smoke Sensor -> Local Alarm Siren (Hardwire). Failsafe.
*   **Path 2 (Digital):** Sensor -> Gateway -> Cloud -> Push Notification -> Fire Dept API.

### Question 517: Build an architecture for offline-first mobile apps.

**Answer:**
(See Q261).
*   **Database:** WatermelonDB / RxDB (Syncs with backend automatically).

### Question 518: How to handle syncing data between edge and cloud?

**Answer:**
*   **Digital Twin:** Cloud maintains a JSON document representing the "Desired State" and "Reported State".
*   **Sync:**
    1.  Device reports `State: Off`.
    2.  User sets `Desired: On`.
    3.  Cloud calculates `Delta`. Sends command `Turn On`.
    4.  Device turns on. Reports `State: On`.

### Question 519: Design a smart home command center backend.

**Answer:**
(See Q251).

### Question 520: Build a local-first video processing system on edge devices.

**Answer:**
*   **Hardware:** Jetson Nano / Coral TPU.
*   **Pipeline:** Camera -> Frame -> Object Detection (YOLO Tiny) -> If Person -> Upload Clip to Cloud.
*   **Privacy:** Faces blurred locally before upload.

---

## 🔸 Collaboration Platforms (Questions 521-530)

### Question 521: Design a live collaborative whiteboard.

**Answer:**
(e.g., Miro / Excalidraw).
*   **Data Structure:** List of Elements (Line, Rect).
*   **Communication:** WebSocket. Broadcast delta updates.
*   **Conflict:**
    *   **Last Write Wins:** Simplest.
    *   **Fractional Indexing:** Insert element between `0.1` and `0.2` -> `0.15`. prevents index shifting.

### Question 522: How would you build real-time Google Docs-style editing?

**Answer:**
(See Q168). Operational Transformation (OT) or CRDT (Yjs / Automerge).

### Question 523: Design a system for version-controlled document editing.

**Answer:**
*   **Snapshots:** Save full content every 10 mins or major change.
*   **Deltas:** Store list of operations (`KeyStroke`) between snapshots.
*   **Restore:** Load Snapshot + Replay Deltas.

### Question 524: How do you manage permissions in collaborative workspaces?

**Answer:**
*   **Model:** `Workspace -> Team -> Project -> Page`.
*   **Inheritance:** User added to Workspace inherits access to all Public Teams.
*   **Sharing:** Share Link token grants access to specific Page (External Guest).

### Question 525: Build a project management app backend (like Trello).

**Answer:**
*   **Board:** Columns (Lists) + Cards.
*   **Ordering:** Lexorank (Jira) or Fractional Indexing.
    *   Move Card A between B and C. New Rank = `(B.rank + C.rank) / 2`.
    *   Avoids updating all card ranks.

### Question 526: Design a task assignment and notification engine.

**Answer:**
*   **Mention:** Parse text for `@username`.
*   **Notify:**
    *   If user online -> WebSocket toast.
    *   If offline -> Email aggregation (Summary of mentions).

### Question 527: Build a calendar invite and availability detection system.

**Answer:**
*   **Availability:**
    *   Bitmasking (30 min slots).
    *   Query: `User A(Bitmap) & User B(Bitmap) & User C(Bitmap)`.
    *   Result: `1` bits are common free slots. fast intersection.

### Question 528: Design a shared annotation system for PDFs/images.

**Answer:**
*   **Coordinate:** Store `(x, y, page_num, comment_id)`.
*   **Overlay:** Client renders PDF, then renders React components at absolute coordinates.
*   **Anchor:** If text flows (HTML), anchor to `DOM Selector` / `Text Range` instead of coordinates.

### Question 529: Design a time-tracking system with team analytics.

**Answer:**
*   **Timer:** Start/Stop events.
*   **Aggregation:**
    *   `TimesheetLine` (Task, Date, Hours).
    *   Materialized View: `Sum(Hours) Group By Project`.
*   **Timezone:** Store Start/Stop in UTC. Display in User Local.

### Question 530: How to ensure consistency in multi-user interactions?

**Answer:**
*   **Optimistic Locking:**
    *   Read `version=5`.
    *   Edit.
    *   Save `WHERE version=5`.
    *   If 0 rows updated (someone else saved `version=6`) -> Throw Error "Data Changed, Refresh".

---

## 🔸 Design Tradeoff Scenarios (Questions 531-540)

### Question 531: When would you use peer-to-peer instead of client-server?

**Answer:**
*   **P2P:**
    *   Heavy Bandwidth (File Sharing - BitTorrent).
    *   Privacy/Censorship Resistance (Blockchain).
    *   Latency (WebRTC Video Call).
*   **Client-Server:**
    *   Centralized Control/Auth.
    *   Complex Logic/Search.

### Question 532: When would you store derived data vs calculate on the fly?

**Answer:**
*   **Calculate:** Cheap logic (`Price * Qty`). Less storage, always consistent.
*   **Store:** Expensive logic (Aggregates, ML Inference). Fast Read, Risk of Stale Data.

### Question 533: Compare column-oriented vs row-oriented databases.

**Answer:**
(See Q152).
*   **Row (MySQL):** Good for transaction (Fetch one User).
*   **Col (Cassandra/Redshift):** Good for Analytics (Avg Age of all Users).

### Question 534: Tradeoffs between batch processing and stream processing.

**Answer:**
*   **Batch:** High Throughput, Simple, High Latency. (Accuracy > Speed).
*   **Stream:** Low Latency, Complex, Lower Throughput per node. (Speed > Accuracy).

### Question 535: When to choose eventual consistency over strict consistency?

**Answer:**
(See Q334).

### Question 536: When to use webhooks vs polling?

**Answer:**
*   **Webhooks:** Event-driven. Real-time. Server pushes. (Best for "Tell me when X happens").
*   **Polling:** Client pulls. (Best for "Status check" or if Client is behind firewall/NAT and can't receive webhook).

### Question 537: Tradeoffs between push vs pull messaging.

**Answer:**
*   **Pull (Kafka):** Consumer controls rate (Backpressure). Harder to achieve low latency.
*   **Push (RabbitMQ):** Low latency. Risk of overwhelming consumer.

### Question 538: When is an in-memory cache harmful?

**Answer:**
*   **Stale Data:** If business logic requires strict consistency (Bank Balance).
*   **Cold Start:** Empty cache kills DB.
*   **Complexity:** Synchronization bugs.

### Question 539: When to replicate vs partition data?

**Answer:**
*   **Replicate:** For Read Scaling & Availability. (Copy whole dataset).
*   **Partition:** For Write Scaling & Typesize. (Split dataset).

### Question 540: When is premature optimization justified?

**Answer:**
*   **Never:** Usually.
*   **Exception:** Schema Design in Databases (Hard to change later). Core Data Structures (Trie vs Map).

---

## 🔸 Dev Tooling, Developer Experience & Platforms (Questions 541-550)

### Question 541: Build a CI/CD pipeline for ML models.

**Answer:**
(See Q238).

### Question 542: Design a feature-flag management platform.

**Answer:**
(See Q193).

### Question 543: Build a schema registry and validation system.

**Answer:**
(See Q408).
*   **Validation:** Producer checks Registry (`POST /schemas/ids/1`). Serializes data.
*   **Consumer:** Deserializes using ID 1.

### Question 544: How would you build a local-first developer playground?

**Answer:**
(e.g., LocalStack).
*   **Mocking:** Emulate AWS APIs (S3/Dynamo) on `localhost:4566`.
*   **Storage:** File system acting as S3 buckets. SQLite acting as DynamoDB.

### Question 545: Build a backend for a real-time log viewer.

**Answer:**
*   **File Watch:** Tail file (`inotify` on Linux).
*   **Transport:** WebSocket sends lines to Browser.
*   **Search:** `grep` implementation in backend or grep in browser memory (if small).

### Question 546: How to design a system that captures runtime metrics?

**Answer:**
*   **Instrumentation:** Library (Prometheus Client).
*   **Hooks:** HTTP Middleware (Duration), GC Hooks, DB Driver Hooks.
*   **Exposition:** `/metrics` endpoint.

### Question 547: Build an automated changelog generator.

**Answer:**
*   **Source:** Git Commit Messages.
*   **Convention:** Conventional Commits (`feat: add login`, `fix: bug`).
*   **Parser:** Group by type (`Features`, `Fixes`).
*   **Output:** Markdown file.

### Question 548: Design a developer API key management platform.

**Answer:**
*   **Key Gen:** Secure Random string (`sk_live_...`).
*   **Storage:** `Hash(Key)` in DB. (Never store plain).
*   **Scope:** `Key -> Scopes ["read", "write"]`.

### Question 549: How to allow secure plugin support in SaaS?

**Answer:**
*   **Webhooks:** Safest. Plugin registers URL.
*   **WASM (WebAssembly):** Run untrusted code in sandbox inside the App. (Figma uses this).
*   **iFrames:** UI Extensions.

### Question 550: Design a sandbox environment manager for services.

**Answer:**
*   **Isolation:** Namespace per Sandbox (K8s).
*   **Data:** Copy subset of anonymized Prod data.
*   **Router:** Host header `sandbox-1.api.com` routes to Sandbox namespace.
