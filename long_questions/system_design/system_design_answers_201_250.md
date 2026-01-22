## 🔸 Fault Tolerance & Resilience (Questions 201-210)

### Question 201: How would you design a self-healing system?

**Answer:**
A system that detects and corrects faults without human intervention.
1.  **Detection:** Health checks, liveness probes (K8s), outlier detection.
2.  **Mitigation:**
    *   **Auto-Restart:** K8s restarts crashed pods.
    *   **Failover:** Switch to a healthy standby.
    *   **Rate Limiting:** Shed load to prevent total collapse.
3.  **Correction:** Re-replication (if a data node dies, create a new copy elsewhere).

### Question 202: How do you detect and recover from cascading failures?

**Answer:**
Cascading failure: A small failure triggers a chain reaction (e.g., Service A fails -> Slower response -> Callers retry -> Callers overload -> Service B fails).
*   **Detection:** High latency, 500s, exponential rise in traffic (retries).
*   **Prevention:**
    *   **Circuit Breakers:** Stop calling a failing service early.
    *   **Backoff & Jitter:** Randomize retries to avoid "thundering herd".
    *   **Bulkheading:** Isolate dependencies (e.g., separate thread pools).

### Question 203: How to isolate failures in microservices?

**Answer:**
*   **Bulkheading Pattern:** Like a ship's watertight compartments.
    *   *Implementation:* Use separate Thread Pools for separate downstream services.
    *   *Result:* If Service A is slow and exhausts its thread pool, Service B (different pool) remains unaffected.
*   **Timeouts:** Fail fast if downstream is unresponsive.

### Question 204: What is bulkheading in system design?

**Answer:**
An implementation of the isolation principle.
*   **Scenario:** Application depends on Database A and API B.
*   **Problem:** API B becomes slow. All threads wait for API B. No threads left for Database A.
*   **Bulkhead:** Allocate 10 threads for API B, 10 for Database A. If API B jams, only its 10 threads are blocked. The rest of the app continues working.

### Question 205: How do you simulate failures for testing?

**Answer:**
**Chaos Engineering.**
*   **Chaos Monkey:** Randomly kills instances in production.
*   **Fault Injection:**
    *   Inject latency (sleep 5s) in network layer (Istio/Service Mesh).
    *   Simulate disk failure / packet loss (tc command).
    *   Return 500 errors for random requests.
*   **Goal:** Verify that alerts fire and redundancy works.

### Question 206: What is circuit breaking and when to use it?

**Answer:**
(Repeated concept, focused implementation).
*   **When:** Calling external services, databases, or APIs over network.
*   **Logic:**
    *   Count consecutive failures.
    *   If Count > Threshold -> OPEN circuit (Throw error immediately).
    *   Sleep for Timeout.
    *   Allow ONE request (Half-Open).
    *   If success -> CLOSE. If fail -> OPEN again.

### Question 207: What happens when a database goes down?

**Answer:**
1.  **App Layer:** Connection pool timeout -> Returns 500 to users.
2.  **Failover (If configured):**
    *   Monitoring detects Master is down.
    *   Election promotes a Slave to Master.
    *   DNS/VirtualIP processes update to point to new Master.
    *   App reconnects. (Downtime: 30s - 2min).

### Question 208: How to design systems for disaster recovery?

**Answer:**
*   **RPO (Recovery Point Objective):** Max data loss allowed. (Strategy: Async replication to another region).
*   **RTO (Recovery Time Objective):** Max downtime allowed. (Strategy: Active-Active or Warm Standby).
*   **Backup:** Point-in-time snapshots (daily/hourly).
*   **Drill:** Regular "Game Days" to practice restoring from backup.

### Question 209: What’s the difference between RTO and RPO?

**Answer:**
*   **RPO (Recovery Point Objective):** "How much data can I lose?"
    *   If RPO = 1 hour, you back up every hour. Failure at 2:59 means losing data since 2:00.
*   **RTO (Recovery Time Objective):** "How fast must we get back up?"
    *   If RTO = 4 hours, system must be live within 4 hours of crash.

### Question 210: How to handle region-wide cloud outages?

**Answer:**
*   **Multi-Region Strategy:**
    *   **Active-Passive:** Region A is live. Region B has data (replicated) but no traffic. If A dies, spin up servers in B and switch DNS.
    *   **Active-Active:** Both regions serve traffic. Global Load Balancer routes users. If A dies, route all to B. (Requires resolving data conflicts).

---

## 🔹 Networking & Protocols (Questions 211-220)

### Question 211: How does TCP work under high latency?

**Answer:**
*   **Problem:** TCP waits for ACK before sending more packets. In high latency (long RTT), throughput drops because the pipeline is empty.
*   **Window Scaling:** Increases window size (bytes in flight) to keep pipe full.
*   **Congestion Control:** Algorithms like BBR (Bottleneck Bandwidth and Round-trip propagation time) optimize for throughput rather than packet loss.

### Question 212: What happens during a TCP handshake?

**Answer:**
3-Way Handshake (SYN, SYN-ACK, ACK).
1.  **Client:** Sends `SYN` (Synchronize) packet with random Sequence Number `A`.
2.  **Server:** Receives `SYN`. Sends `SYN-ACK` with Ack Number `A+1` and its own Sequence Number `B`.
3.  **Client:** Receives `SYN-ACK`. Sends `ACK` with Ack Number `B+1`.
*   Connection Established.

### Question 213: What is long polling vs short polling?

**Answer:**
*   **Short Polling:**
    *   Client: "Any new data?" -> Server: "No" (immediate).
    *   Client waits 1s, asks again.
    *   *Cons:* Wasted resources, high traffic.
*   **Long Polling:**
    *   Client: "Any new data?" -> Server: Holds connection open...
    *   ...Data arrives... -> Server: "Yes, here it is".
    *   *Pros:* Less traffic, instant update.

### Question 214: How would you design a protocol over UDP?

**Answer:**
(e.g., Implementing reliability on UDP for a game).
1.  **Sequencing:** Add Sequence Number to packet header to detect out-of-order.
2.  **Acks:** Receiver sends ACK for critical packets.
3.  **Retransmission:** Sender resends if ACK not received within timeout.
4.  **Flow Control:** Sender limits rate to avoid flooding receiver.
*   *Note:* This reinvents TCP, but with control over "what to drop" (e.g., drop old move packets, keep chat packets).

### Question 215: How to reduce latency in cross-country communication?

**Answer:**
1.  **CDN:** Serve static content from edge.
2.  **Edge Compute:** Move logic (e.g., auth check) to edge.
3.  **Global Accelerator (Anycast):** Route user to closest entry point on AWS network, then traverse dedicated fiber backbone (bypassing public internet hops).
4.  **Protocol:** Use HTTP/2 or HTTP/3 (QUIC) to reduce connection overhead.

### Question 216: What is QUIC and how does it compare to HTTP/2?

**Answer:**
*   **QUIC:** A transport protocol built on top of UDP.
*   **Problems with HTTP/2 (TCP):** Head-of-line blocking. If one packet is lost, all streams over that TCP connection wait.
*   **QUIC Benefits:**
    *   **No Head-of-Line Blocking:** Streams are independent. Lost packet only delays that stream.
    *   **Faster Handshake:** 0-RTT (Zero, Round Trip Time) resumes.
    *   **Connection Migration:** Survive IP change (Wi-Fi to 4G) seamlessly.

### Question 217: How do NAT and firewalls affect distributed systems?

**Answer:**
*   **NAT (Network Address Translation):** Hides internal IPs.
    *   *Issue:* P2P systems (WebRTC) can't connect directly.
    *   *Fix:* STUN/TURN servers to discover public IP or relay traffic.
*   **Firewalls:** Block ports.
    *   *Design:* Ensure required ports (e.g., 443, 8080) are open in Security Groups.

### Question 218: How do CDNs route traffic efficiently?

**Answer:**
*   **Anycast DNS:** Multiple servers share the same IP address. BGP routing directs the user to the topologically nearest server.
*   **DNS Resolution:** The DNS server detects the resolver's IP and returns the CNAME of the nearest Edge Location.

### Question 219: How to handle packet loss in real-time systems?

**Answer:**
*   **Video/Audio:**
    *   **FEC (Forward Error Correction):** Send redundant data. Receiver reconstructs lost packet without asking for retransmission.
    *   **Concealment:** Interpolate (guess) the missing frame.
*   **Gaming:** Interpolate position (dead reckoning).

### Question 220: Explain DNS resolution and its failure points.

**Answer:**
*   **Flow:** Browser check cache -> OS Cache -> Recursive Resolver (ISP) -> Root Server -> TLD Server (.com) -> Authoritative Server (example.com).
*   **Failure Points:**
    *   **DDoS on DNS Provider:** (e.g., Dyn attack). Site becomes unreachable even if servers are up.
    *   **Cache Poisoning:** Attacker inserts fake IP into cache.
    *   **Propagation Delay:** TTL prevents users from seeing IP update immediately.

---

## 🔹 Search, Indexing & Metadata (Questions 221-230)

### Question 221: Design a tag-based search system.

**Answer:**
*   **Schema:** `ItemTags` table (ItemID, TagID).
*   **Query:** "Find items with Tag A OR Tag B".
    *   `SELECT DISTINCT ItemID FROM ItemTags WHERE TagID IN (A, B)`.
*   **Search Engine:** Using Inverted Index is faster. `Tag A -> [Item1, Item2]`.

### Question 222: How do you implement "did you mean" suggestions?

**Answer:**
1.  **Fuzzy Matching:** Query existing index with Levenshtein distance 1 or 2.
2.  **Log Analysis:** Look at query logs. "Users who searched 'iphnoe' mostly clicked 'iphone' results next." Build a correction map.
3.  **Spell Checker:** (See Q142).

### Question 223: Design a distributed indexing system.

**Answer:**
*   **Sharding:** Split documents into shards (by Hash/DocID).
*   **Indexing:** Each node builds a local inverted index for its documents.
*   **Query:** Scatter-Gather.
    *   Coordinator sends query to ALL shards.
    *   Each shard returns top 10 matches (TF-IDF score).
    *   Coordinator merges, sorts, returns top 10 global.

### Question 224: How to handle synonyms and stemming in search?

**Answer:**
*   **Stemming:** Reduce words to root. "Running", "Ran", "Runs" -> "Run". Index "Run". Query "Runner" -> "Run".
*   **Synonyms:** Expansion.
    *   Query "Phone" -> Internally expand to "Phone OR Mobile OR Cellphone".
    *   Use a dictionary/thesaurus during query processing.

### Question 225: How would you build faceted search?

**Answer:**
(e.g., Amazon sidebar: Filter by Brand, Price, Size).
*   **Datastructure:** Search engine (Elasticsearch) Aggregations.
*   **Query:** "Query: Shoes". "Aggregations: Brand (count), Size (count)".
*   **Result:** Returns list of shoes AND buckets: `Nike (50), Adidas (30)`.

### Question 226: What is inverted indexing?

**Answer:**
(Core search concept).
*   **Forward Index:** `Doc1 -> "apple banana"`.
*   **Inverted Index:**
    *   `"apple" -> [Doc1, Doc3]`
    *   `"banana" -> [Doc1, Doc2]`
*   **Search "apple":** Lookup logic is O(1) (Hash map look up), then return list.

### Question 227: How do you store and search metadata at scale?

**Answer:**
(e.g., Dropbox file metadata).
*   **Requirement:** ACID transactions (move file), strong consistency.
*   **Solution:** Sharded SQL (MySQL/Vitess) or NewSQL (CockroachDB).
*   **Index:** Separate Service. Metadata DB updates -> Stream -> Elasticsearch (for fuzzy search).

### Question 228: How to optimize autocomplete suggestions with popularity?

**Answer:**
*   **Trie Node:** Store "Top 5" at every node.
*   **Update:**
    *   Accumulate search frequency in logs.
    *   Offline Job (MapReduce) aggregates frequency.
    *   Rebuild/Update Trie with new weights.
*   **Performance:** Lookup is O(Prefix Length), extremely fast.

### Question 229: How to merge search indexes from multiple data sources?

**Answer:**
(Federated Search).
1.  **Standardize:** Convert UserDB, ProductDB, SupportLog into a common JSON schema.
2.  **Ingest:** ETL pipeline pushes standardized docs into a single Elasticsearch cluster.
3.  **Alias:** Use Index Aliases to query `user_index` and `product_index` as one `search_all` index.

### Question 230: Design a location-aware search engine.

**Answer:**
(e.g., "Restaurants near me").
*   **Geo-hashing:** Divide world into grid (Geohash or QuadTree).
*   **Index:** Store `RestaurantID` in the grid cell bucket.
*   **Query:**
    1.  User at `(lat, lon)`. Identify cell.
    2.  Query items in that cell (and 8 neighbors).
    3.  Filter by radius.
*   **Tool:** Elasticsearch `geo_point` type, Redis `GEOADD` / `GEORADIUS`.

---

## 🔹 AI/ML System Design (Questions 231-240)

### Question 231: How do you deploy a machine learning model to production?

**Answer:**
1.  **Containerize:** Package model code + dependencies + weight file (pickle/onnx) in Docker.
2.  **Orchestrate:** Deploy on K8s (Kubeflow/Seldon) or Serverless (SageMaker/Lambda).
3.  **API:** Expose via REST/gRPC endpoint (`/predict`).

### Question 232: How to monitor model drift?

**Answer:**
Drift: The model's performance degrades because real-world data changes (Concept Drift).
*   **Monitor:** Compare statistical distribution of Training Data vs Live Inference Data.
*   **Metric:** KL Divergence, PSI (Population Stability Index).
*   **Action:** If drift detected -> Trigger retraining pipeline.

### Question 233: Design a recommendation system for a video platform.

**Answer:**
*   **Candidate Generation (Recall):** Fast. Select 1000 candidates from millions (Collaborative Filtering / Two-Tower Model).
*   **Ranking:** Slow. detailed scoring of 1000 candidates (Deep Neural Net with many features).
*   **Re-ranking:** Apply business rules (remove watched, diversify categories).

### Question 234: How to serve real-time predictions with low latency?

**Answer:**
*   **Optimize Model:** Quantization (Float32 -> Int8), Pruning (remove weak neurons). Use ONNX Runtime / TensorRT.
*   **Architecture:**
    *   **CPU vs GPU:** Use GPU for batch processing, CPU for single low-latency request (sometimes).
    *   **Caching:** Cache prediction for same inputs.

### Question 235: How to do A/B testing for ML models?

**Answer:**
*   **Canary:** Routes 1% traffic to New Model (Model B).
*   **Shadow Mode:** Route 100% traffic to Model A (return result) AND Model B (log result). Compare B's log against actual user outcome (did they click?) without affecting user.
*   **Interleaving:** Show results from A and B mixed. See which one gets clicked.

### Question 236: How do you version and roll back ML models?

**Answer:**
*   **Model Registry:** (MLflow). Tracks version `v1.0`, `v1.1` binary artifacts.
*   **Deployment:** Use K8s/Serving mesh.
    *   Deploy `v1.1`.
    *   If error rate increases or latency spikes -> Revert traffic to `v1.0` endpoint automatically.

### Question 237: Design a fraud detection system using ML.

**Answer:**
*   **Features:** Transaction Amount, IP Location, DeviceID, Velocity (transactions per min).
*   **Pipeline:**
    *   **Synchronous:** Simple Rule Engine (Amount > 10k).
    *   **Asynchronous:** ML Model scores transaction. If High Probability -> Flag for Manual Review / Block card.
*   **Graph:** Use Graph ML to detect rings of fraudsters.

### Question 238: How to retrain models automatically with new data?

**Answer:**
**CI/CD for ML (CT - Continuous Training).**
*   **Trigger:** Schedule (Weekly) or Drift-Triggered.
*   **Pipeline:** Airflow/Kubeflow chain:
    1.  Extract new data.
    2.  Validate data.
    3.  Train Model.
    4.  Evaluate (Accuracy > threshold?).
    5.  Push to Registry.

### Question 239: How do you manage feature engineering pipelines?

**Answer:**
*   **Feature Store:** (Feast / Tecton).
*   **Problem:** Python code calculates "Avg Spend" for training. Production needs SAME logic for inference.
*   **Solution:** Feature Store calculates feature once -> Stores in Offline (Training) and Online (Redis for Inference) stores consistently.

### Question 240: Design an AI-powered personal assistant.

**Answer:**
*   **ASR (Speech to Text):** Convert audio.
*   **NLU (Natural Language Understanding):** Intent Classification ("Play Music") + Slot Filling ("Song: Despacito").
*   **Dialog Manager:** Maintains context/state machine.
*   **TTS (Text to Speech):** Generate audio response.
*   **Privacy:** On-device processing for "Wake Word".

---

## 🔹 Blockchain & Decentralized Systems (Questions 241-250)

### Question 241: Design a decentralized identity verification system.

**Answer:**
*   **Standard:** DID (Decentralized Identifiers) and Verifiable Credentials (VC).
*   **Flow:**
    1.  Issuer (Govt) signs a VC ("Passport").
    2.  Holder (User) stores VC in Wallet.
    3.  Verifier (Hotel) asks for proof.
    4.  Holder presents VC. Verifier checks signature on Blockchain.
*   **Privacy:** Zero-Knowledge Proofs (prove "Age > 18" without revealing birthdate).

### Question 242: How would you store large files on the blockchain?

**Answer:**
*   **You Don't.** Storing data on-chain is prohibitively expensive ($1000s per MB).
*   **Solution:** IPFS (InterPlanetary File System).
    1.  Upload file to IPFS -> Get Hash (CID).
    2.  Store just the CID (32 bytes) on the Blockchain smart contract.

### Question 243: What is Merkle Tree and where is it used?

**Answer:**
A binary tree of hashes.
*   **Structure:** Leaves = Hash of data blocks. Parent = Hash(Left Child + Right Child). Root = Merkle Root.
*   **Usage:**
    *   **Blockchain:** Root stored in block header. Allows "Light Clients" to verify a transaction exists in a block without downloading the whole block (Merkle Proof).
    *   **Dynamo/Cassandra:** Merkle Trees compare data between replicas to find differences quickly (Anti-entropy).

### Question 244: How would you design a crypto wallet backend?

**Answer:**
*   **Hot Wallet:** Connected to internet, signs user transactions automatically. (Risky).
*   **Cold Wallet:** HSM (Hardware Security Module) / Air-gapped machine signing.
*   **Design:**
    *   DB stores User Balances (Internal Ledger).
    *   Watcher Service listens to Blockchain for deposits.
    *   Signing Service holds Private Keys (Encrypted).

### Question 245: Design a smart contract-based subscription service.

**Answer:**
*   **Push vs Pull:** Crypto is "Push" only (User pushes money). You can't pull from their wallet.
*   **Solution:**
    *   **Escrow:** User deposits 12 months' fee into Contract. Contract releases to Merchant monthly.
    *   **Allowance:** User `approve()` contract to spend X tokens. Contract calls `transferFrom()` monthly.

### Question 246: How do consensus algorithms work in blockchain?

**Answer:**
Solves: "Which block is next?" in a trustless network.
*   **PoW (Proof of Work):** Solve hard math puzzle (CPU intensive). First solver wins. Secure but wasteful.
*   **PoS (Proof of Stake):** Validators lock up money (Stake). Random selection biased by stake size. Efficient.

### Question 247: Design a blockchain explorer like Etherscan.

**Answer:**
*   **Indexer:** Node connects to Ethereum. Reads every block.
*   **Parser:** Decodes transactions, events, internal calls.
*   **DB:** Relational DB (Postgres) essential for queries like "All txs for Address X" (Blockchain is a linked list, bad for queries).
*   **Frontend:** Queries Postgres API.

### Question 248: What is proof-of-stake vs proof-of-work?

**Answer:**
*   **PoW:** Security via Energy. Cost to attack > Reward. (Bitcoin).
*   **PoS:** Security via Economic Value. Attacker needs >51% of total coins. If they attack, their stake is "slashed" (destroyed). (Ethereum).

### Question 249: How do you build a decentralized messaging app?

**Answer:**
*   **Protocol:** Whispers / Waku (Ethereum P2P layer) or Matrix.
*   **Storage:** IPFS for media.
*   **Encryption:** Double Ratchet Algorithm (Signal) for E2EE.
*   **Network:** No central server. Messages propagate via Gossip.

### Question 250: How do you ensure integrity in a blockchain network?

**Answer:**
*   **Chaining:** Block contains Hash of Previous Block. Changing block 5 invalidated hashes of 6, 7, 8...
*   **Consensus:** Honest majority rejects invalid chains.
*   **Digital Signatures:** Every transaction signed by sender's Private Key. No one can spoof a transaction.
