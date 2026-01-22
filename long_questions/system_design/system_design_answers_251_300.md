## 🔸 IoT & Edge Computing (Questions 251-260)

### Question 251: Design a smart home system.

**Answer:**
*   **Components:** Sensors (Temp, Motion), Actuators (Light, Lock), Hub/Gateway.
*   **Protocol:** Zigbee/Z-Wave (Low power) -> Hub -> MQTT (Over Wi-Fi) -> Cloud.
*   **Architecture:**
    *   **Local Processing:** Hub handles fast rules ("Motion detected -> On Light"). Works offline.
    *   **Cloud Processing:** Long-term storage, complex analytics, remote control via App.
*   **Security:** mTLS between Hub and Cloud.

### Question 252: How would you handle intermittent connectivity in IoT devices?

**Answer:**
*   **Local Caching:** Store data on the device (SQLite/Flash) when offline.
*   **Queueing:** MQTT "QoS 1" (At least once) or "QoS 2" (Exactly once). Message stays in local queue until ack received.
*   **Sync:** When online, upload batch. Use timestamps to resolve conflicts (Last-Write-Wins or append-only log).

### Question 253: How do you securely update firmware over the air?

**Answer:**
**OTA (Over-The-Air) Update Process:**
1.  **Build:** Manufacturer signs firmware binary with Private Key.
2.  **Distribute:** Upload to CDN. Push notification to device via MQTT.
3.  **Verify:** Device downloads binary. Verifies signature using pre-installed Public Key. Validates Checksum.
4.  **Install:** A/B Partitioning. Install to Partition B. Reboot. If boot fails, rollback to Partition A.

### Question 254: How do you process data at the edge?

**Answer:**
*   **Concept:** Process data on the device or a nearby gateway instead of sending everything to the cloud.
*   **Tools:**
    *   **AWS Greengrass:** Runs Lambda functions locally on the hub.
    *   **TFLite:** run ML inference (e.g., Object Detection) on the camera chip.
*   **Benefit:** Low latency (millisecond decisions), privacy, reduced bandwidth cost.

### Question 255: Design a low-latency data pipeline for a smart city.

**Answer:**
*   **Ingest:** 5G/LoRaWAN towers receive sensor data (Traffic, Air Quality).
*   **Edge:** Multi-Access Edge Computing (MEC) nodes at cell towers filter noise.
*   **Transport:** MQTT over UDP (for speed).
*   **Core:** Kafka -> Flink (Stream Processing) -> Traffic Light Controller.
*   **Latency Goal:** < 50ms from Sensor to Traffic Light.

### Question 256: What is fog computing?

**Answer:**
A decentralized computing infrastructure placed between the Cloud and IoT devices (Edge).
*   **Hierarchy:** Cloud (Global) -> Fog (Regional/City) -> Edge (Device/Gateway).
*   **Role:** Fog nodes (e.g., Identifying a valid user at the building entrance) offload the Cloud but have more power than the Edge.

### Question 257: How do you handle device authentication at scale?

**Answer:**
*   **X.509 Certificates:** Factory provisions a unique certificate burned into a TPM (Hardware Security) chip on each device.
*   **Protocol:** mTLS. Device presents cert to IoT Core. Core validates signature against Root CA.
*   **Lifecycle:** Revocation Lists (CRL) or OCSP to block stolen devices.

### Question 258: How would you reduce network usage in edge computing?

**Answer:**
1.  **Filtering:** Don't send "Temp = 72" every second. Send only if "Temp changes by > 1 degree".
2.  **Aggregation:** Send `Avg(Temp)` every minute instead of raw data.
3.  **Compression:** Use Protobuf instead of JSON.
4.  **Delta Updates:** Send only changed fields.

### Question 259: Design a fleet tracking system for delivery vehicles.

**Answer:**
*   **Device:** GPS module sends `(lat, lon, speed)` every 10s via 4G.
*   **Ingestion:** MQTT Broker -> Kafka.
*   **Storage:**
    *   **Hot:** Redis GEO (Current location).
    *   **Cold:** Cassandra/TimescaleDB (Trip history).
*   **Query:** "Where is Truck 5?" -> Redis. "Show path taken yesterday" -> Cassandra.

### Question 260: How do you prevent sensor spoofing?

**Answer:**
(Attacker injecting fake data).
*   **Physical Security:** Tamper-proof hardware.
*   **Crypto:** Sign every data packet with device's Private Key.
*   **Anomaly Detection:** ML model on the cloud detects impossible physics (e.g., Temp jumps from 20C to 100C in 1s, or Location jumps 500km).

---

## 🔸 Mobile & Offline Systems (Questions 261-270)

### Question 261: How would you build an app that works offline-first?

**Answer:**
*   **Architecture:** App reads/writes to Local DB (SQLite/Realm), NOT the Network.
*   **Sync Engine:** Background process syncs Local DB with Remote DB.
*   **UI:** Optimistic UI. Show "Done" immediately. Show spinner only for the sync status icon.

### Question 262: How to sync data efficiently between mobile and server?

**Answer:**
*   **Delta Sync:** Store `LastSyncTimestamp`.
    *   Client asks: "Give me changes since T1".
    *   Server queries: `SELECT * FROM data WHERE updated_at > T1`.
*   **Soft Deletes:** Don't delete rows. Set `deleted_at` so clients can download the "deletion event".

### Question 263: How would you implement conflict resolution in sync?

**Answer:**
*   **Last Write Wins (LWW):** Compare timestamps. Newest overwrites. (Simple, but can lose data).
*   **Manual Merge:** Flag conflict, ask user to choose version. (Git style).
*   **CRDTs:** Mathematically mergeable data structures. (Complex).
*   **Server Authority:** Server creates a "Merge Commit" and forces client to re-fetch.

### Question 264: How do you compress data for slow networks?

**Answer:**
*   **Transport:** Gzip / Brotli for HTTP.
*   **Format:** Binary (Protobuf/FlatBuffers) instead of JSON.
*   **Images:** WebP/AVIF.
*   **Resizing:** Request specific size (`img.jpg?w=300`) from CDN.

### Question 265: Design a mobile wallet system.

**Answer:**
*   **Security:** Biometric Auth (FaceID) needed to open app.
*   **Tokenization:** Don't store Card Number. Store Token provided by Payment Processor (Stripe).
*   **Offline:** Display QR Code (TOTP) for merchant to scan (requires no internet on phone).
*   **Transaction:** Ledger pattern (Double Entry Bookkeeping).

### Question 266: How do push notifications work at scale?

**Answer:**
*   **Registration:** App gets DeviceToken from OS. Sends to Backend.
*   **Send:** Backend -> Queue -> Worker -> Calls APNS (Apple) / FCM (Google).
*   **Optimization:** Batch requests (send 1000 tokens in one HTTP/2 call to APNS).
*   **Cleanup:** Remove invalid tokens (User uninstalled app) based on APNS feedback mechanism.

### Question 267: Design a live location-sharing feature.

**Answer:**
*   **Protocol:** WebSocket. Client sends `LocationUpdate` every 2s.
*   **Backend:** Redis `GEOADD`.
*   **Privacy:** Ephemeral sharing (Redis Key TTL = 1 hour).
*   **Scale:** If 1M users, Redis Cluster sharded by UserID.

### Question 268: How to optimize battery usage in mobile applications?

**Answer:**
1.  **Batch Networking:** Make one big request instead of 10 small ones (wakes up radio fewer times).
2.  **Background Processing:** Use OS Job Schedulers (WorkManager) that run only when charging/Wi-Fi connected.
3.  **Location:** Use Geofencing (passive) instead of GPS polling (active).

### Question 269: How to handle mobile version compatibility?

**Answer:**
*   **Force Update:** API checks `MinSupportedVersion`. If app is too old, block usage and show "Please Update" dialog.
*   **API Versioning:** Backend supports v1, v2, v3.
*   **Feature Flags:** "Enable New UI" flag is false for old versions.

### Question 270: How would you secure sensitive data in mobile storage?

**Answer:**
*   **Keychain/Keystore:** Use OS-provided secure storage for Tokens/Passwords. Encrypted by hardware.
*   **Encryption:** SQLChiper for SQLite.
*   **No Caching:** Disable HTTP caching for sensitive API endpoints (`Cache-Control: no-store`).

---

## 🔸 Observability & Reliability (Questions 271-280)

### Question 271: Design a log correlation engine.

**Answer:**
*   **Goal:** Connect App Logs, LB Logs, and DB Logs for a single request.
*   **Trace ID:** Injected at Ingress (Load Balancer). Passed via HTTP Headers (`X-Trace-ID`) to all downstream services.
*   **Logging:** Every log line includes `[TraceID]`.
*   **UI:** Splunk/Kibana groups logs by this ID.

### Question 272: What is distributed tracing? How would you implement it?

**Answer:**
*   **Tools:** Jaeger, Zipkin, OpenTelemetry.
*   **Spans:** Each operation (DB Query, Http Call) is a "Span" with StartTime, EndTime, ParentID.
*   **Visualization:** Gantt chart showing where time was spent.
*   **Sampling:** Trace only 0.1% of requests to save storage.

### Question 273: How do you detect slow queries?

**Answer:**
*   **Database:** Enable Slow Query Log (`slow_query_log_file` in MySQL) for queries > 1s.
*   **Application:** APM (New Relic/Datadog) instruments JDBC/ORM drivers.
*   **Metrics:** Histogram of query duration. High P99 suggests slow queries.

### Question 274: Design a custom alerting system.

**Answer:**
*   **Rules:** Defined in YAML (`IF avg(cpu) > 90 FOR 5m`).
*   **Evaluator:** Cron job queries TSDB every minute.
*   **State:** Maintain state (Alert Firing vs Resolved).
*   **Notification:** Deduplicate alerts -> Route to PagerDuty/Slack.
*   **Silence:** Support "Maintenance Mode".

### Question 275: How to create a health dashboard for microservices?

**Answer:**
*   **Discovery:** Poll K8s API for all Services.
*   **Checks:** Call `/health` endpoint of each service.
*   **Aggregator:** Compute "Overall Status" (Green/Yellow/Red).
*   **Dependency Map:** Visualize Service A depends on Service B. Determining "Root Cause".

### Question 276: What is SLO, SLA, and SLI?

**Answer:**
*   **SLI (Indicator):** The metric. (e.g., Latency).
*   **SLO (Objective):** Internal goal. (e.g., "99% requests < 200ms").
*   **SLA (Agreement):** Contract with customer. (e.g., "If < 99%, we refund 10%").
*   *Note:* SLA is looser than SLO.

### Question 277: How to track business metrics from logs?

**Answer:**
*   **Log:** `Order Placed amount=100 currency=USD`.
*   **Metric Sink:** Use **Grok Exporter** (Prometheus) or **CloudWatch Metric Filter**.
*   **Pattern:** Regex match `amount=(\d+)`.
*   **Result:** Counter `orders_total` increments; Histogram `order_value` observes 100.

### Question 278: Design a queryable log storage system.

**Answer:**
*   **ELK Stack:**
    *   **Elasticsearch:** Indexing.
    *   **Hot/Warm/Cold:** Move logs to cheaper storage as they age.
*   **Loki:**
    *   Indexes *only* metadata (labels), not the text content.
    *   Much cheaper. Grep at query time.

### Question 279: What metrics would you monitor for a payment system?

**Answer:**
1.  **Success Rate:** `(Success / Total)`. Alarm if drops below 98%.
2.  **Latency:** P99 processing time.
3.  **Decline Reasons:** Sharp spike in "Insufficient Funds" might mean a bug or fraud.
4.  **Wallet Balance:** Integrity check (Total User Balance == Bank Account Balance).

### Question 280: How do you handle noisy alerts?

**Answer:**
*   **Thresholding:** Stop alerting on spikes; alert on trends.
*   **Hysteresis:** Alert on >90%; Resolve on <80% (Prevent flapping).
*   **Grouping:** Group 100 "Pod Failed" alerts into 1 "Cluster Issue" notification.
*   **Routing:** Send Info/Warn to Slack; Critical to PagerDuty.

---

## 🔸 Product-specific Designs (Questions 281-290)

### Question 281: Design a digital signature service.

**Answer:**
(e.g., DocuSign).
*   **Identity:** Email verification / 2FA to prove signer identity.
*   **Storage:** Secure storage of PDF.
*   **Cryptography:**
    *   Hash the PDF content.
    *   Sign Hash with Service's Private Key (Timestamped).
*   **Audit Trail:** Log IP, Time, Email for every view/sign action.
*   **Verification:** Anyone can verify Service's signature.

### Question 282: How to build a collaborative calendar system?

**Answer:**
*   **Data Model:** `Event` (Start, End, Owner, Invitees).
*   **Conflict:** "Double Booking".
    *   Use Optimistic Locking (`WHERE version = v`).
    *   Constraint checking (`WHERE NOT OVERLAPS`).
*   **Recurrence:** Store rule (`RRULE:FREQ=WEEKLY`), calculate instances on read.
*   **Timezones:** Store everything in UTC. Convert to User TZ on UI.

### Question 283: Design a voting/polling platform with live results.

**Answer:**
*   **Write:** High volume bursts (TV Show voting).
    *   Ingest via Kafka.
    *   Aggregator (Flink) counts votes in 1s windows.
    *   Write increments to Redis/Cassandra.
*   **Read:** Clients poll JSON from S3/CDN (updated every 5s) or WebSocket.
*   **Integrity:** One vote per UserID/IP (deduplication in Flink).

### Question 284: Design a document approval system.

**Answer:**
*   **Workflow Engine:** (Camunda/Temporal).
*   **State Machine:** `Draft -> Pending_Mgr -> Pending_Director -> Approved`.
*   **Action:** Manager clicks "Approve" -> Triggers Webhook -> Transition State -> Notify Director.
*   **Timeout:** If Manager doesn't act in 2 days -> Escalate / Auto-reject.

### Question 285: How to design an API monetization platform?

**Answer:**
(e.g., RapidAPI).
*   **Gateway:** Kong/Apigee.
*   **Metering:** Sidecar/Plugin counts requests per API Key.
*   **Quota:** `If RequestCount > PlanLimit -> 429`.
*   **Billing:** Async job aggregates usage daily -> Charges Stripe.

### Question 286: Design a cloud cost monitoring tool.

**Answer:**
*   **Ingest:** Pull Cost Usage Reports (CUR) from AWS S3 (CSV format).
*   **Process:** ETL (Glue/Spark) to normalize data.
*   **Tagging:** Group costs by `Project`, `Team`.
*   **Anomaly:** Machine Learning (Forecast vs Actual). "Why did EC2 spend double today?".

### Question 287: Build a digital content watermarking system.

**Answer:**
*   **Visible:** Overlay text (FFmpeg filter).
*   **Invisible:** Steganography (Modify least significant bits of image pixels).
*   **Dynamic:** Embed `UserID` in the watermark on download. If leaked, decode watermark to find the leaker.

### Question 288: Design a stock price alert system.

**Answer:**
*   **Input:** Real-time stream from Exchange.
*   **Rule:** User wants "Alert if Apple > $150".
*   **Matching:**
    *   Store rules in a Trie or Interval Tree.
    *   For each price tick, check matching rules.
*   **Notification:** High priority Push.

### Question 289: Design a plagiarism checker backend.

**Answer:**
(See Q195).

### Question 290: How would you build an auction system?

**Answer:**
*   **Real-time:** WebSocket.
*   **Concurrency:** "Last second bidding wars".
*   **Order:** Redis atomic increments/scripts or In-memory matching engine.
*   **Timer:** Server-side authoritative clock. When `Time == End`, stop accepting bids.

---

## 🔸 Privacy, Compliance & Governance (Questions 291-300)

### Question 291: How do you handle GDPR data deletion?

**Answer:**
"Right to be Forgotten".
*   **Architecture:**
    *   **Fact Table:** Store PII.
    *   **Event Store:** Store standard events referencing `UserID`.
*   **Deletion:**
    *   **Hard Delete:** Delete row in PII table.
    *   **Crypto Shredding:** Encrypt PII with per-user key. Destroy the key. Data remains (encrypted) but is unreadable (effectively deleted).

### Question 292: How to log access to sensitive data?

**Answer:**
*   **Middleware:** Intercepts Reads.
*   **Log:** Structured log: `User=Alice Accessed=MedicalRecord ID=Bob Reason=Support`.
*   **Storage:** WORM (Write Once Read Many) storage (S3 Object Lock) to prevent tampering.

### Question 293: What is data masking and where to apply it?

**Answer:**
Hiding parts of data. (e.g., `4111-xxxx-xxxx-1234`).
*   **Dynamic:** Database proxy masks data on-the-fly based on User Role. (Support sees masked, Admin sees clear).
*   **Static:** Mask data when copying from Production to Staging DB.

### Question 294: How do you design audit logs?

**Answer:**
*   **Immutability:** Ensures logs cannot be changed.
*   **Completeness:** Log "Who, What, Where, When, Why".
*   **Storage:** Separate secure bucket.
*   **Retention:** Keep for 7 years (Compliance).

### Question 295: How do you implement RBAC (Role-Based Access Control)?

**Answer:**
*   **Entities:** `User`, `Role` (Admin, Editor), `Permission` (Read, Write).
*   **Mapping:** `User -> Roles -> Permissions`.
*   **Check:** `hasPermission(User, 'Article:Write')`.
*   **JWT:** Embed Roles/Permissions in Token for stateless checking (or verify against DB).

### Question 296: How to encrypt data at rest and in transit?

**Answer:**
*   **Transit:** TLS 1.2/1.3 everywhere.
*   **Rest:**
    *   **Disk Encryption:** AWS EBS Encryption / BitLocker.
    *   **Application Encryption:** Encrypt specific columns (SSN) before inserting.
*   **Key Management:** Use KMS (Key Management Service). Rotate keys annually.

### Question 297: Design a user consent management system.

**Answer:**
(Cookie Banner Backend).
*   **Schema:** `UserConsent` (UserID, CookieCategory, Granted, Date, IP).
*   **API:** Javascript asks "Can I run Analytics?". Backend answers based on `Granted` status.
*   **Audit:** Prove valid consent was given if audited.

### Question 298: How to implement “Right to be forgotten”?

**Answer:**
(See Q291).
It requires mapping all data locations. A central "Deletion Service" publishes `DeleteUser` event. All services (Email, Order, Logs) consume event and scrub data.

### Question 299: How do you classify sensitive vs public data?

**Answer:**
*   **Discovery Tool:** Scan DBs/S3 for patterns (Credit Card Regex, SSN format). Use AWS Macie / Google DLP.
*   **Tagging:** Tag schema/buckets with `Confidentiality: High/Medium/Low`.
*   **Policy:** Deny public access to buckets tagged `High`.

### Question 300: How to enforce data residency rules in cloud apps?

**Answer:**
"German user data must stay in Germany".
*   **Partitioning:** Shard User DB by Region (EU shard vs US shard).
*   **Routing:** Route German users to EU Data Center.
*   **Storage:** Configure S3 buckets in `eu-central-1` to disable Cross-Region Replication to outside regions.
