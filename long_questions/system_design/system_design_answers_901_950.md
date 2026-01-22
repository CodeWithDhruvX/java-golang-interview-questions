## 🔸 Serverless & Event-Driven Patterns (Questions 901-910)

### Question 901: Design a serverless backend for image resizing.

**Answer:**
*   **Trigger:** S3 `ObjectCreated` event.
*   **Compute:** Lambda Function.
*   **Layer:** ImageMagick / Sharp layer (Node.js).
*   **Storage:** Write resized image to `bucket-resized/`. This prevents infinite loop triggers.

### Question 902: How do you handle "Cold Start" in FaaS?

**Answer:**
*   **Keep-Warm:** Ping function every 5 mins.
*   **Provisioned Concurrency:** AWS feature to keep N instances initialized.
*   **Language:** Use Go/Rust instead of Java/Spring (JVM startup time).

### Question 903: Build a stateful workflow using serverless functions.

**Answer:**
*   **Orchestrator:** AWS Step Functions / Azure Logic Apps.
*   **State:** Externalize state to DynamoDB. Function is stateless.
*   **Pass:** Output of Func A passed as JSON input to Func B.

### Question 904: Design a serverless API with rate limiting per tenant.

**Answer:**
*   **Gateway:** API Gateway (Usage Plans).
*   **Custom:** Lambda Authorizer checks Redis `Rate:{TenantID}`. Returns `429` / `Allow` policy.

### Question 905: How to coordinate distributed transactions in serverless?

**Answer:**
*   **Saga Pattern:** (See Q329).
*   **Step Functions:** `Try -> Catch -> Compensate` block.
    *   State `ProcessPayment` fails -> Jump to `RefundInventory`.

### Question 906: Build a real-time chat using serverless websockets.

**Answer:**
*   **Connection:** API Gateway WebSocket API maintains connection ID.
*   **Store:** `ConnectionID` -> `UserID` in DynamoDB.
*   **Push:** `POST https://.../@connections/{conn_id}` to send message.

### Question 907: Design a serverless map-reduce job.

**Answer:**
*   **Map:** S3 Batch Operations invokes Lambda for 1M objects.
*   **Shuffle:** Mapper writes intermediate results to DynamoDB/S3.
*   **Reduce:** DynamoDB Streams trigger Reducer Lambda.

### Question 908: How to secret management in serverless?

**Answer:**
*   **Environment Variables:** Encrypted at rest.
*   **Secrets Manager:** Fetch at runtime. Cache in `/tmp` for warm reuse. avoids API cost.

### Question 909: Design a fan-out notification system using serverless.

**Answer:**
*   **SNS:** Topic `NewPost`.
*   **Subscribers:**
    *   Lambda A (Email).
    *   Lambda B (Push).
    *   SQS C (Analytics).
*   **Scaling:** SNS invokes Lambdas in parallel (thousands concurrent).

### Question 910: Build a serverless URL shortener.

**Answer:**
*   **Write:** API Gateway -> Lambda -> DynamoDB (`Hash` -> `URL`).
*   **Read:** API Gateway -> Lambda -> 301 Redirect.
*   **CDN:** Cache Redirect at CloudFront Edge. Reduce Lambda hits by 99%.

---

## 🔸 Big Data & Analytics Deep Dive (Questions 911-920)

### Question 911: Design a Data Lake architecture on cloud.

**Answer:**
*   **Zones:**
    *   `Raw`: Immutable original JSON/CSV.
    *   `Curated`: Parquet/Avro (Cleaned).
    *   `Consumer`: Aggregated tables for BI.
*   **Catalog:** AWS Glue / Hive Metastore.

### Question 912: How to handle schema evolution in Parquet files?

**Answer:**
*   **Discovery:** Crawler detects new columns.
*   **Table Format:** Apache Iceberg / Hudi / Delta Lake.
    *   Metadata files track "Schema v1 applies to file A", "Schema v2 applies to file B". Merge on read.

### Question 913: Design a real-time dashboard for log analytics.

**Answer:**
*   **Ingest:** Filebeat -> Kafka.
*   **Index:** Elasticsearch (Hot/Warm architecture).
*   **Visualize:** Kibana.

### Question 914: Build a privacy-compliant data deletion pipeline (GDPR).

**Answer:**
*   **Problem:** Deleting 1 row from Parquet is expensive (Have to rewrite 1GB file).
*   **Solution:** Delta Lake `DELETE`. Logic:
    *   Write `Tombstone` record.
    *   Async `VACUUM` job rewrites files to physically remove data.

### Question 915: Design a high-cardinality metrics storage engine.

**Answer:**
(See Cortex / Thanos).
*   **TSDB:** Inverted Index for Labels (`IP=10.0.0.1` -> `[SeriesID]`).
*   **Problem:** Too many unique IPs = explosion.
*   **Solution:** Roll up (Drop IP dimension) after 24h.

### Question 916: How to deduplicate data in a Kappa architecture?

**Answer:**
*   **Kappa:** Stream only. No batch layer.
*   **Dedup:** RocksDB state store (Flink). Store `MessageHash` for 7 days. Drop duplicates.

### Question 917: Design a system to join two infinite streams.

**Answer:**
(See Q557).
*   **Constraint:** Must define Window (e.g., "Join if Click happens within 10 mins of Impression").
*   **State:** Buffer Impression.

### Question 918: Build a distributed crawler for the web.

**Answer:**
*   **Frontier:** URL Queue (Redis Priority Queue).
*   **Politeness:** `Map<Domain, LastAccessTime>`. Delay if too frequent.
*   **Storage:** WARC files in S3.
*   **DNS:** Caching DNS resolver.

### Question 919: How to optimize top-K queries on big data?

**Answer:**
*   **Approx:** Count-Min Sketch.
*   **Merge:** MapReduce. Each mapper finds local Top-K. Reducer merges lists.

### Question 920: Design a data lineage tracking system.

**Answer:**
*   **Graph:** `Job A reads Table X, writes Table Y`.
*   **Ingest:** SQL Parser (extracts `FROM` / `JOIN` tables).
*   **Visual:** "Table Y is tainted because Table X has bad data".

---

## 🔸 AR/VR & Spatial Computing (Questions 921-930)

### Question 921: Design a Point Cloud streaming system for AR.

**Answer:**
*   **Format:** PLY / PCD.
*   **Octree:** Divide 3D space. Stream only "Visible" nodes based on Camera Frustum.
*   **LOD:** Stream low detail first -> refine.

### Question 922: Build a Spatial Anchor synchronization system.

**Answer:**
*   **Anchor:** `ID`, `Descriptor (Visual Features)`, `Pose (x,y,z)`.
*   **Look up:** Device sends Camera Features. Server matches Descriptor. Returns Pose relative to World.
*   **Cloud:** Azure Spatial Anchors / Google Cloud Anchors.

### Question 923: Design a backend for a massive multiplayer VR world (Metaverse).

**Answer:**
*   **Partitioning:** Spatial Hashing. Server A handles (0,0) to (100,100).
*   **Handoff:** User walks across boundary -> Transfer session A to B.
*   **Voice:** Proximity Chat (Volume decreases with distance).

### Question 924: How to handle physics sync in networked VR?

**Answer:**
*   **Authority:** Server authoritative.
*   **Prediction:** Client predicts physics locally. Packet arrives -> Correction (Lerp).
*   **Ownership:** "User holding the cup" has authority over Cup physics to ensure 0 latency interaction.

### Question 925: Design a collaborative 3D modeling platform.

**Answer:**
*   **CRDT:** Tree structure of Scene Graph (`Root -> Transform -> Cube`).
*   **Locking:** User selects Vertex. Lock vertex.

### Question 926: Build a low-latency 360-video streaming service.

**Answer:**
*   **Projection:** Equirectangular.
*   **Tiling:** Cut 4K video into 4x4 tiles.
*   **ABR:** Based on Head Orientation ("Look at"), download tiles in view at High Res, others at Low Res.

### Question 927: How to persistent AR content in the real world?

**Answer:**
*   **Localization:** GPS (Coarse) + VPS (Visual Positioning System using Street View).
*   **DB:** `GeoHash` -> `List[ContentID]`.

### Question 928: Design a gesture recognition pipeline.

**Answer:**
*   **Edge:** Hand Tracking model (MediaPipe) runs on device (30fps).
*   **Output:** Skeleton (21 points).
*   **Classifier:** "Index + Thumb touching" = Click.

### Question 929: Build a backend for digital avatars and inventory.

**Answer:**
(See Q764). Interoperable standard (VRM).

### Question 930: Design a spatial audio engine for VR.

**Answer:**
*   **Input:** `SourcePos`, `ListenerPos`, `RoomGeometry`.
*   **HRTF:** Head Related Transfer Function.
*   **Calc:** Done Client-Side (CPU heavy). Server only syncs positions.

---

## 🔸 Gaming & Real-Time Simulation (Questions 931-940)

### Question 931: Design a matchmaking system for ranked games.

**Answer:**
*   **Pool:** Users with `MMR` (Matchmaking Rating).
*   **Bucket:** Group by `MMR +/- 100` and `Ping < 50ms`.
*   **Expansion:** If no match in 30s, expand range to `+/- 200`.

### Question 932: Build a state replication system for FPS games.

**Answer:**
*   **Snapshot:** Send full state every 50ms? (Too much bandwidth).
*   **Delta:** Send only changes.
*   **Compression:** Quantize Floats (Send angle as Byte 0-255).

### Question 933: Design a leaderboards system with seasonal resets.

**Answer:**
*   **Redis:** `Rank:Season_1`.
*   **Archive:** At season end, rename key to `Rank:Season_1_Final`. Create empty `Rank:Season_2`.

### Question 934: How to detect aimbots and wallhacks?

**Answer:**
*   **Server-Side:** "Heuristics".
    *   Did cursor snap 180 degrees in 1 frame?
    *   Did bullet hit head through "Unbreakable Wall"?
*   **Client-Side:** Anti-Cheat kernel driver (e.g., Vanguard) checks memory injection.

### Question 935: Design a global game server scaling strategy.

**Answer:**
*   **Agones:** K8s Custom Resource for Game Servers.
*   **Allocation:** Request game server -> Returns IP:Port.
*   **Lifecycle:** Server State `Allocated` -> `Ready` -> `Shutdown`.

### Question 936: Build an in-game economy and trading system.

**Answer:**
*   **ACID:** DB Transaction essential.
*   **Audit:** Record `Transfer(From: A, To: B, Item: Sword)`.
*   **Binding:** Some items "Soulbound" (Cannot trade).

### Question 937: Design a replay system for strategy games.

**Answer:**
*   **Deterministic:** Save `RandomSeed` and `List[PlayerInputs]`.
*   **Playback:** Re-simulate game from Seed applying inputs at exact frames. Result is identical.

### Question 938: How to minimize lag compensation (Rewind)?

**Answer:**
*   **Hitbox:** Server stores history of hitboxes for last 1s.
*   **Check:** "At T=100, Player shot. At T=100, where was Enemy?" Rewind enemy to position at T=100. Check collision.

### Question 939: Design a system for seamless map loading (Open World).

**Answer:**
*   **Chunking:** Divide world into grids.
*   **Streaming:** Load Grid (0,1) when player approaches within 100m. Unload Grid (0,-1).

### Question 940: Build a cross-platform save system (PC/Console/Mobile).

**Answer:**
*   **Conflict:** Timestamp check.
*   **Schema:** Binary Blob or JSON?
*   **Storage:** Cloud Key-Value store `SaveSlot_1`.

---

## 🔸 Web3 & Decentralized Applications (Questions 941-950)

### Question 941: Design a decentralized file storage system (IPFS inspired).

**Answer:**
*   **Addressing:** CIA (Content Addressed). `Hash(Content)`.
*   **DHT:** Kademlia DHT to find "Who has this hash?".
*   **Swarm:** Download chunks from multiple peers.

### Question 942: Build an indexing service for blockchain events (The Graph).

**Answer:**
*   **Ingest:** Connect to ETH Node RPC.
*   **Scan:** Filter logs by `ContractAddress` and `Topic`.
*   **Map:** Apply mapping logic (`Transfer -> UserBalance`).
*   **Store:** Postgres.
*   **Serve:** GraphQL.

### Question 943: Design a cryptocurrency wallet backend (Custodial).

**Answer:**
*   **Hot Wallet:** Online. Holds 5% funds for withdrawals.
*   **Cold Wallet:** Offline/HSM. Holds 95%.
*   **Signing:** Key never leaves HSM. Transaction sent to HSM -> Signed Tx returned.

### Question 944: How to implement "Sign in with Ethereum"?

**Answer:**
*   **Challenge:** Server generates `Nonce`.
*   **Sign:** User signs message `Nonce + Domain` with Private Key.
*   **Verify:** Server runs `ecrecover(Sig)`. If Address matches User, Auth successful.

### Question 945: Design a decentralized oracle (Chainlink).

**Answer:**
*   **Aggregation:** 20 Nodes fetch "Price of ETH" from different APIs (Binance, Coinbase).
*   **Commit:** Nodes submit value on-chain.
*   **Consensus:** Smart Contract takes Median of values.
*   **Reward:** Nodes paid in link for honest data.

### Question 946: Build a Layer 2 rollup relayer.

**Answer:**
*   **Accumulate:** Collect 1000 txs off-chain.
*   **Proof:** Generate ZK-Proof (Validity) or Merkle Root (Optimistic).
*   **Post:** Submit Root + CallData to L1. Saves gas.

### Question 947: Design a Merkle Tree for whitelist verification.

**Answer:**
*   **Off-chain:** Hash all 10k whitelist addresses -> Root.
*   **On-chain:** Store only Root.
*   **Proof:** User sends `[Hash A, Hash B...]` (Merkle Path) to prove they are in the tree.

### Question 948: Build a decentralized voting system (DAO).

**Answer:**
*   **Snapshot:** Off-chain signing. Free.
*   **Weight:** `BalanceOf(Token)` at Block Height X.
*   **Storage:** Store votes on IPFS.

### Question 949: Design a bridge between two blockchains.

**Answer:**
*   **Lock:** User locks Token A on Chain 1.
*   **Witness:** Relayer detects Lock Event.
*   **Mint:** Relayer calls Contract on Chain 2 to Mint Wrapped Token A.
*   **Risk:** 51% attack on Bridge Validators (Multi-sig).

### Question 950: How to mitigate MEV (Miner Extractable Value) in block building?

**Answer:**
*   **Private Mempool:** Flashbots.
*   **Bundle:** Searchers submit "Bundle" of txs.
*   **Promise:** Miner guarantees bundle order and no front-running, or reverts bundle.
