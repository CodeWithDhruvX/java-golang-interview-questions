## 🔸 AI/ML Infrastructure & Data Pipelines (Questions 801-810)

### Question 801: Design a system for training ML models on petabytes of data.

**Answer:**
*   **Storage:** Data Lake (S3/HDFS).
*   **Compute:** Distributed Training (Ray / Horovod / Spark).
*   **Strategy:** Data Parallelism. Split data into chunks. Each node trains a replica of the model. Gradient updates synced via Parameter Server or Ring-AllReduce.

### Question 802: Build a feature store for ML pipelines.

**Answer:**
(See Q507).
*   **Feast:** Open Source Feature Store.
*   **Registry:** Metadata of all features.
*   **Serving:** `get_online_features()` (Redis) for low latency. `get_offline_features()` (BigQuery) for training.

### Question 803: Design a pipeline for real-time ML model inference.

**Answer:**
*   **Request:** API Gateway -> Load Balancer -> Model Service.
*   **Optimization:** Batching (Micro-batches of 16-32 requests).
*   **Hardware:** NVIDIA Triton Inference Server on GPU instances.

### Question 804: How would you version models, datasets, and parameters?

**Answer:**
*   **DVC (Data Version Control):** Git for data.
*   **MLflow:** Tracks `RunID`, `Params`, `Metrics`, `ArtifactURI`.
*   **Lineage:** Trace back `Model v2` -> `Data v3` -> `Code commit SHA`.

### Question 805: Design an experiment tracking system for data scientists.

**Answer:**
*   **Dashboard:** Weights & Biases / TensorBoard.
*   **Metadata DB:** Postgres stores hyperparams (`lr=0.01`).
*   **Artifact Store:** S3 stores model weights (`model.pth`).

### Question 806: How to ensure reproducibility in ML training pipelines?

**Answer:**
*   **Containerization:** Docker image contains exact library versions (`requirements.txt` with hashes).
*   **Seed:** Fix random seeds (`torch.manual_seed(42)`).
*   **Data Snapshot:** Tag dataset version.

### Question 807: Build a scalable hyperparameter tuning platform.

**Answer:**
*   **Search:** Bayesian Optimization / Grid Search.
*   **Orchestrator:** Katib (K8s). Spawns 100 trials in parallel.
*   **Early Stopping:** If accuracy isn't improving, kill the pod (save cost).

### Question 808: Design a real-time fraud detection ML pipeline.

**Answer:**
(See Q502).

### Question 809: Build a system to serve multiple models with low latency.

**Answer:**
*   **Multi-Model Server:** TorchServe.
*   **Routing:** `Header: x-model-version: v2`.
*   **Resource Sharing:** Pack multiple small models into GPU VRAM (MIG - Multi-Instance GPU).

### Question 810: Design a pipeline for explainable AI (XAI) results.

**Answer:**
*   **Recall:** Fetch Prediction.
*   **Explain:** Run SHAP/LIME kernel on the input.
*   **Compute:** Expensive. Run async background worker.
*   **Store:** Save Explanation JSON (`Feature A contributed +0.3`).

---

## 🔸 Financial, Banking & Trading Systems (Questions 811-820)

### Question 811: Design a system to detect suspicious banking transactions.

**Answer:**
*   **Rules:** Pre-check (Velocity).
*   **ML:** Post-check (Pattern Matching).
*   **Graph:** "Is money flowing to known bad actor?" (Graph DB traversal).

### Question 812: Build a micro-lending platform with KYC and credit scoring.

**Answer:**
*   **KYC:** Third-party API (Onfido) verifies ID.
*   **Score:** Alternative Data (Utility bills, SMS logs if permissioned).
*   **Ledger:** Double-entry accounting. `Debit LoanReceivable`, `Credit Cash`.

### Question 813: Design a stock order matching engine.

**Answer:**
*   **Data Structure:** Limit Order Book. Two Min/Max Heaps (Bids/Asks).
*   **Speed:** Single Threaded (LMAX Architecture) to avoid locking overhead. In-Memory.
*   **Log:** Write commands to Ring Buffer (Disruptor) before processing for recovery.

### Question 814: How would you design a cross-border payment gateway?

**Answer:**
*   **FX:** Integrate with Liquidity Providers.
*   **Messaging:** ISO 20022 / SWIFT.
*   **Compliane:** OFAC Screening.

### Question 815: Build a reconciliation system for bank transactions.

**Answer:**
(See Q776).
*   **3-Way:** Database vs Payment Gateway vs Bank Statement.
*   **Format:** Parse MT940 / CAMT.053 files.

### Question 816: Design a subscription billing engine with retries and proration.

**Answer:**
*   **Proration:** Upgrade mid-month. `Credit Remaining` + `Charge New_Plan_Remaining`.
*   **Dunning:** Smart Retries. If card fail "Insufficient Funds", retry on Payday (15th/30th).

### Question 817: Build a system for invoice generation and PDF delivery.

**Answer:**
*   **Queue:** Async job.
*   **PDF:** Headless Chrome (Puppeteer) renders HTML to PDF. Pixel perfect.
*   ** Delivery:** Email via SES / SendGrid.

### Question 818: Design a ledger system for recording financial transactions.

**Answer:**
*   **Immutability:** Append-only table.
*   **Integrity:** Row includes `Hash(PrevRowHash + CurrentRow)`.
*   **Balance:** Snapshot `Balance` table updated by triggers.

### Question 819: How do you ensure transactional integrity across currencies?

**Answer:**
*   **Atomic:** All money movement in one DB Transaction.
*   **Precision:** Use `BigDecimal` or Integer (Cents). Never Float.
*   **Conversion:** Store Exchange Rate used at moment of transaction.

### Question 820: Build a fraud-resistant peer-to-peer payment system.

**Answer:**
(Venmo/CashApp).
*   **Risk:** Device Fingerprinting.
*   **Challenge:** If New Device + High Amount -> Require FaceID / OTP.
*   **Cool-down:** Delay funds availability for 24h for new users.

---

## 🔸 Logistics, Supply Chain & Transportation (Questions 821-830)

### Question 821: Design a package tracking system for courier networks.

**Answer:**
*   **Event Sourcing:** `ScannedAtHub`, `OutForDelivery`.
*   **Cassandra:** High write throughput. Partition Key `TrackingID`.
*   **Search:** Elasticsearch for "Find all packages in Chicago Hub".

### Question 822: Build a warehouse inventory management system.

**Answer:**
*   **Bin:** `Location: A-01-02`.
*   **SKU:** `Item: Shampoo`.
*   **Locking:** Optimistic. Picker A reserves item.
*   **Sync:** Handheld scanner updates DB in real-time.

### Question 823: Design a last-mile delivery optimization system.

**Answer:**
*   **Problem:** VRP (Vehicle Routing Problem).
*   **Input:** 100 packages, 5 drivers.
*   **Geocoding:** Convert addresses to Lat/Lon.
*   **Cluster/Route:** K-Means clustering -> TSP per cluster.

### Question 824: How would you route trucks in real-time based on traffic?

**Answer:**
*   **Provider:** Google Maps Platform / Mapbox Traffic API.
*   **Re-route:** If `ETA_New > ETA_Old + 10 mins`, suggest detour.

### Question 825: Build a return pickup and refund processing system.

**Answer:**
*   **Init:** User generates QR code.
*   **Scan:** Courier scans QR at pickup. Trigger `ReturnInitiated` event.
*   **Refund:** "Instant Refund" (Risk based) or "Refund on Arrival" (Scan at warehouse).

### Question 826: Design a multi-hop shipment system with dependencies.

**Answer:**
*   **Graph:** `Origin -> Hub A -> Hub B -> Dest`.
*   **SLA:** Calculate `CutoffTime` at each Hub.
*   **Miss:** If package misses Hub A truck, recalculate ETA.

### Question 827: How do you track items across warehouses globally?

**Answer:**
*   **Global View:** Aggregated nightly.
*   **Transfer:** `TransferOrder(Source, Dest, SKU, Qty)`.
*   **Transit:** Inventory "In Transit" is virtual warehouse.

### Question 828: Build a delivery time prediction engine using live data.

**Answer:**
*   **ML:** Regression Model.
*   **Features:** Distance, DayOfWeek, Weather, Traffic, CourierSpeed.
*   **Training:** Historical Deliveries.

### Question 829: Design a cold-chain logistics system with sensor alerts.

**Answer:**
*   **IoT:** Sensor sends `Temp` every 5 mins.
*   **Stream:** Compare `Temp` vs `AllowedRange(SKU)`. (Ice Cream -20C, Bananas 12C).
*   **Alert:** Push Notification to Driver app "Check AC!".

### Question 830: Build a logistics system for multi-vendor fulfillment.

**Answer:**
(Dropshipping).
*   **Order Router:** Split order. Item A -> Vendor 1. Item B -> Vendor 2.
*   **Integration:** EDI / API connection to Vendors.
*   **Consolidation:** Optionally route to Cross-dock facility to combine.

---

## 🔸 Multi-Tenant SaaS & Admin Platforms (Questions 831-840)

### Question 831: Design a multi-tenant SaaS backend with strict data isolation.

**Answer:**
*   **Row Security:** `WHERE tenant_id = ?` Mandatory.
*   **Schema:** `SET search_path = tenant_schema`. (Postgres).
*   **Database:** Separate DB for VIP tenants.

### Question 832: Build a tenant-aware rate limiting and quota system.

**Answer:**
(See Q584).

### Question 833: How do you manage schema evolution per tenant?

**Answer:**
*   **Standard:** All tenants on same schema (easiest).
*   **Custom Fields:** JSONB column for tenant-specific data.
*   **Enterprise:** Migration script loops over 1000 tenant schemas to apply `ALTER TABLE`.

### Question 834: Design a system for custom branding per customer.

**Answer:**
(See Q736).
*   **CDN:** Serve `logo_{tenant_id}.png`.

### Question 835: Build an admin dashboard for monitoring tenant activity.

**Answer:**
*   **Metrics:** `RequestsPerSec`, `StorageUsed`, `ErrorRate` per Tenant.
*   **Aggregation:** Cortex / Thanos for multi-tenant Prometheus.

### Question 836: Design audit trails for actions performed by tenant admins.

**Answer:**
*   **Target:** `User: Admin_A`, `Action: DeleteUser`, `Target: User_B`.
*   **Export:** CSV download available to Tenant Owner.

### Question 837: How to implement access delegation to external consultants?

**Answer:**
*   **Invite:** Tenant invites Consultant Email.
*   **Role:** `Role: Consultant` (Restricted).
*   **Expiry:** Automatic revocation after Project End Date.

### Question 838: Design a tenant lifecycle management platform.

**Answer:**
*   **Onboard:** `POST /tenants`. Provision DB, S3 bucket, Stripe Sub.
*   **Suspend:** Block API access.
*   **Offboard:** Archive data (Cold Storage). Delete hot resources.

### Question 839: Build a multi-tenant notification preference system.

**Answer:**
*   **Default:** System Defaults.
*   **Tenant:** Tenant Admin overrides defaults for Org.
*   **User:** User overrides Org settings (if allowed).
*   **Resolution:** `User > Tenant > System`.

### Question 840: Design a config-as-a-service platform for multiple tenants.

**Answer:**
*   **GitOps:** Repo per tenant? Or Single Repo with folders.
*   **Propagation:** Agent polls Config Service.
*   **Validation:** CUE / Jsonnet to validate config structure.

---

## 🔸 Mobile-First & Offline-First Systems (Questions 841-850)

### Question 841: Design a sync service for mobile apps with offline mode.

**Answer:**
(See Q261).
*   **Sync:** `LastSyncTimestamp`.
*   **Pull:** Server sends records modified since Timestamp.

### Question 842: Build a mobile push notification delivery system.

**Answer:**
*   **Tokens:** Store `FCM_Token` / `APNS_Token` mapping to UserID.
*   **Topic:** Pub/Sub.
*   **Worker:** Calls FCM/APNS API. Handles `InvalidToken` response (cleanup DB).

### Question 843: How would you sync mobile caches with eventual consistency?

**Answer:**
*   **TTL:** Short cache time.
*   **Etag:** Conditional GET.
*   **Silent Push:** Server wakes app on critical update to fetch fresh data.

### Question 844: Design a mobile usage analytics pipeline.

**Answer:**
*   **Batch:** Events stored in SQLite locally.
*   **Upload:** Upload on WiFi/Charging or every 1 hour.
*   **Format:** Compressed JSON (GZIP).

### Question 845: Build a delta update system for reducing mobile data usage.

**Answer:**
*   **Binary Diff:** `bsdiff` for app updates.
*   **Data Diff:** Server computes JSON Patch between `Version_Old` and `Version_New`. Sends Patch.

### Question 846: Design a system for location-based push campaigns.

**Answer:**
*   **Geofence:** OS handles monitoring.
*   **Trigger:** App wakes on "Enter Region".
*   **Local Notification:** App displays alert immediately (No network needed) or pings server for contextual deal.

### Question 847: Build a system to handle device ID rotations and merges.

**Answer:**
*   **IDFA/GAID:** Resettable.
*   **Fingerprint:** Use `InstallID` (UUID generated on first run).
*   **Link:** When User logs in, link `InstallID` to `UserID`. Attribute anon history to User.

### Question 848: How to queue and sync user actions taken offline?

**Answer:**
*   **Queue:** Persistent Job Queue (Redux Offline / WorkManager).
*   **Order:** FIFO.
*   **Conflict:** Server rejects if state changed. Client must rebase/prompt user.

### Question 849: Design a multi-device session management system.

**Answer:**
*   **List:** `Sessions` table. `DeviceModel`, `LastActive`, `IP`.
*   **Remote Logout:** User clicks "Log out other devices". Server deletes session ID from Redis.

### Question 850: Build a mobile telemetry event aggregation system.

**Answer:**
*   **Client Agg:** `Count(Crashes)`++ locally.
*   **Send:** Send summary, not raw events (if high volume), unless "Detailed Telemetry" enabled.
