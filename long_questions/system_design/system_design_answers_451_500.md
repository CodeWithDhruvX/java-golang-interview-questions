## 🔸 E-commerce / Fintech Specific (Questions 451-460)

### Question 451: Design a flash sale system that avoids overselling.

**Answer:**
(See Q383).
*   **Lua Script (Redis):** `GET stock`, `IF stock > 0 THEN DECR stock; RETURN Success ELSE RETURN Fail`.
*   **Atomicity:** Redis guarantees single-threaded execution of the script. No race condition.
*   **Post-Process:** Success -> Queue -> Payment Service.

### Question 452: How would you ensure idempotency in a payment API?

**Answer:**
*   **Idempotency Key:** Client generates UUID (`IDempotency-Key: abc-123`).
*   **DB:** `PaymentRequests` table (Key, Status, Response).
*   **Logic:**
    1.  Check if Key exists.
    2.  If exists and Status=Success, return *saved* response.
    3.  If not exists, process payment -> Insert Key -> Return Response.

### Question 453: Design a refund processing pipeline.

**Answer:**
*   **State Machine:** `RefundRequested -> PendingGateway -> GatewayAck -> Completed`.
*   **Async:** Don't block HTTP. Return `202 Accepted`.
*   **Consistency:** If Gateway says "Success" but DB update fails -> Cron job reconciles status later.

### Question 454: How to handle high-value transactions securely?

**Answer:**
*   **2FA/MFA:** Step-up auth required.
*   **Fraud Check:** Blocking call to Risk Engine (Velocity checks, Device fingerprint).
*   **Audit:** WORM storage.
*   **Encryption:** Field-Level Encryption (FLE) for Card Numbers in DB.

### Question 455: Design an invoicing and tax calculation microservice.

**Answer:**
*   **Tax:** Integration with Avalara / Vertex (Rules are too complex to build).
*   **Invoice:** Immutable PDF generation.
*   **Storage:** S3 (with Object Lock for compliance).
*   **Model:** `InvoiceLineItem` (Qty, UnitPrice, TaxRate, TaxAmount).

### Question 456: How do you build a secure shopping cart across devices?

**Answer:**
*   **Anonymous:** Cart stored in `Redis` key `session:123`.
*   **Login:** Merge logic.
    *   Load `UserCart` from DB.
    *   Load `SessionCart` from Redis.
    *   Combine items.
    *   Save to DB.
    *   Clear Redis.

### Question 457: Design a pricing engine for a global marketplace.

**Answer:**
*   **Layers:**
    1.  Base Price (USD).
    2.  FX Rate (Lookup Daily).
    3.  Country Adjustment (Purchasing Power Parity).
    4.  Tax (VAT).
*   **Cache:** Pre-calculate final prices for Top 10 Currencies. Cache in Redis.

### Question 458: Build a credit scoring engine based on activity data.

**Answer:**
*   **Data Points:** Repayment History, Wallet Usage, App Behavior.
*   **Batch:** Spark job aggregates last 6 months data -> Feature Vector.
*   **Inference:** Random Forest model -> Score (300-900).
*   **Serving:** API `GetScore(User)` returns pre-computed score.

### Question 459: How would you design EMI loan repayment tracking?

**Answer:**
*   **Schedule:** Generate N rows in `RepaymentSchedule` table (`DueDate`, `Amount`, `Status`).
*   **Reminder:** Cron runs daily `WHERE DueDate = Tomorrow`.
*   **Collection:** Auto-debit (ACH/Card) on DueDate.
*   **Late:** If fail -> Retry 3 times -> Mark Overdue -> Penalty.

### Question 460: How to handle currency fluctuation in real-time checkout?

**Answer:**
*   **Quote:** User sees `$100 (in BTC)`. System generates `QuoteID` valid for 10 minutes.
*   **Lock:** Exchange locks rate with Liquidity Provider.
*   **Expire:** If user pays after 10 mins, `QuoteID` is invalid. Re-quote.

---

## 🔸 Database-Specific Questions (Questions 461-470)

### Question 461: How would you implement TTL for database records?

**Answer:**
*   **NoSQL (Dynamo/Redis/Mongo):** Native support.
*   **SQL:**
    *   Partition by Time (`orders_2023_01`). Drop old partitions (`DROP TABLE`). (Fast).
    *   `DELETE WHERE created_at < X`. (Slow, causes fragmentation/locking).

### Question 462: What is LSM Tree and where is it used?

**Answer:**
**Log-Structured Merge-Tree.**
*   **Write:** Append to MemTable (RAM). Fast.
*   **Flush:** Sorted String Table (SSTable) on disk. Immutable.
*   **Read:** Check MemTable -> Check SSTables (Bloom Filter optimizes).
*   **Used In:** Cassandra, RocksDB, LevelDB. (Write-optimized).

### Question 463: How do you design a query cache with invalidation?

**Answer:**
*   **Look-aside:** App checks Cache.
*   **Invalidation:**
    *   **Exact Match:** `DEL "query:select*from_users"`. (Hard to know all permutations).
    *   **Table Tagging:** Tag cached key with `table:users`. When `users` table updates, increment `tag_version`. Query key includes `tag_version` so old keys become misses.

### Question 464: Design a schema for audit logs with replay capability.

**Answer:**
*   **columns:** `SeqID` (BigInt auto-inc), `Timestamp`, `Actor`, `Action`, `Payload` (JSON).
*   **Replay:**
    *   Reader tracks `LastSeqID`.
    *   Fetch `WHERE SeqID > LastSeqID ORDER BY SeqID ASC`.

### Question 465: Design a write-heavy logging system with efficient read support.

**Answer:**
*   **Write:** Kafka (Append only).
*   **Read:** ClickHouse (Columnar).
    *   Kafka Connect sinks data to ClickHouse in batches.
    *   ClickHouse indexes columns for fast aggregations.

### Question 466: How would you enforce uniqueness across multiple fields?

**Answer:**
*   **DB Constraint:** `UNIQUE(col1, col2)`.
*   **Application (High Scale):**
    *   Create `hash = SHA256(col1 + col2)`.
    *   Store `hash` in a Key-Value store with `SETNX` (Set if Not Exists).

### Question 467: When to choose columnar databases?

**Answer:**
*   **OLAP:** Analytical queries. "Sum of Sales where Region=US".
*   **Why:** Reads only the `Sales` and `Region` columns from disk. Skips `CustomerName`, `Address`. 100x faster IO.
*   **Examples:** Redshift, Snowflake, BigQuery.

### Question 468: How to manage schema evolution in NoSQL?

**Answer:**
*   **Pattern:** Versioning.
    *   Document has `v: 1`.
    *   Code: `if doc.v == 1 { data = migrate(doc); }`.
*   **Lazy Migration:** Update schema when the record is *saved* next time.

### Question 469: Build a DB partitioning strategy based on activity.

**Answer:**
(Hot/Cold data).
*   **Hot:** Recent data in High-Performance NVMe DB.
*   **Warm:** Data > 1 month in HDD DB.
*   **Cold:** Data > 1 year in S3 (Parquet).
*   **App:** Query router decides which DB to hit based on Date Range.

### Question 470: How to perform online reindexing of large datasets?

**Answer:**
*   **Aliasing (Elasticsearch):**
    1.  Create `Index_V2`.
    2.  Double-write to `Index_V1` and `Index_V2`.
    3.  Backfill V2 from V1.
    4.  Atomic Switch Alias `Production` -> `Index_V2`.
    5.  Delete V1.

---

## 🔸 Concurrency & Parallelism (Questions 471-480)

### Question 471: How to design a system with concurrent writers and readers?

**Answer:**
*   **MVCC (Multi-Version Concurrency Control):**
    *   Writers create new version of row. Don't block Readers.
    *   Readers see a consistent snapshot of the "old" version until commit.
*   **Locking:**
    *   `ReadWriteLock`: Multiple readers allowed. 1 Writer exclusive.

### Question 472: How to handle deadlocks in distributed lock management?

**Answer:**
*   **Prevention:** Acquire locks in predefined order (Sort resource IDs).
*   **Detection:** Wait-for Graph. Cycle = Deadlock.
*   **Timeout:** `lock.acquire(timeout=5s)`. If fail, release all held locks and retry.

### Question 473: Design a job queue with priority and fairness.

**Answer:**
*   **Priority:** 3 Queues (High, Med, Low).
*   **Worker:** Checks High first. If empty, check Med.
*   **Starvation:** High priority flood prevents Low from running.
*   **Fairness (Weighted Round Robin):** Process 5 High, 3 Med, 1 Low per cycle.

### Question 474: How to parallelize file processing safely?

**Answer:**
*   **Split:** Divide file ranges (0-1GB, 1GB-2GB).
*   **Idempotency:** Output filename `part_1.csv`.
*   **Locking:** Use File Locking or lease in DynamoDB to ensure Worker A owns Chunk 1.

### Question 475: Build a system for running background jobs with retry, delay, and cancel.

**Answer:**
(e.g., Sidekiq / BullMQ).
*   **Retry:** Exponential Backoff (ZSET in Redis with timestamp).
*   **Delay:** Add to `Scheduled` ZSET. Poller moves to `Active` List when time arrives.
*   **Cancel:** Check `isCancelled` flag in Redis at start of job execution.

### Question 476: What are the issues in shared-nothing architecture?

**Answer:**
*   **State:** No central DB. Data partitioned across nodes.
*   **Issue:**
    *   **Cross-Node Joins:** Expensive shuffling of data over network.
    *   **Rebalancing:** Adding a node requires moving TBs of data.

### Question 477: How would you implement mutex in a distributed system?

**Answer:**
(e.g., Redis `SETNX`).
*   **Acquire:** `SET lock_key unique_id NX PX 10000`.
*   **Release:** Lua Script: `if GET lock_key == unique_id then DEL lock_key`. (Prevent deleting someone else's lock if yours expired).
*   **Redlock:** Acquire lock on N/2+1 independent Redis nodes for safety.

### Question 478: Design a bulk data processing system using worker pools.

**Answer:**
*   **Pattern:** Fan-out / Fan-in.
*   **Master:** Reads input, pushes 1000 tasks to Queue.
*   **Workers:** Pull task, process, push result to ResultQueue.
*   **Aggregator:** Reads ResultQueue, writes final output.

### Question 479: How to implement exponential backoff and jitter across clients?

**Answer:**
*   **Backoff:** `Sleep = min(Cap, Base * 2^Attempt)`.
*   **Jitter:** `Sleep = Sleep / 2 + random(0, Sleep / 2)`.
*   **Why:** Prevents waves of synchronized retries (Thundering Herd).

### Question 480: Design a checkpointing system in long-running pipelines.

**Answer:**
*   **State:** Store usage `offset` or `processed_ids` in external store (Redis/Zookeeper).
*   **Commit:** Every 100 items, save state.
*   **Recovery:** On crash, read state. Resume from `offset + 1`.

---

## 🔸 User Experience-Oriented (Questions 481-490)

### Question 481: Design a “Save for later” system in a shopping app.

**Answer:**
*   **DB:** `SavedItems` table (UserID, ProductID, AddedAt).
*   **Move:**
    *   Transaction: Remove from `Cart` table -> Insert into `SavedItems`.
*   **Notify:** If Product Price drops -> Query `SavedItems` -> Send Push.

### Question 482: How would you build "undo" functionality in UI-backed system?

**Answer:**
*   **Optimistic UI:** UI shows "Done".
*   **Delay:** Wait 5 seconds before making API call.
*   **Cancel:** If user clicks Undo -> Cancel timer.
*   **Compensating Action:** If API already called -> Call `Revert` API.

### Question 483: Design a progress-aware file uploader.

**Answer:**
*   **XHR:** `onprogress` event gives bytes sent.
*   **Resumable:** If failure, server returns `offset`. Client seeks file stream to `offset` and continues.

### Question 484: How to cache partial page loads with API responses?

**Answer:**
*   **ETag:** Server returns `ETag: "v1"`.
*   **Client:** Sends `If-None-Match: "v1"`.
*   **Server:** If data unchanged -> Returns `304 Not Modified` (Empty Body). Client uses cache.

### Question 485: How would you design smart auto-refresh in dashboards?

**Answer:**
*   **Adaptive:**
    *   Active Tab: Poll every 10s.
    *   Background Tab: Poll every 5m.
*   **Push:** WebSocket "Data Changed" event triggers a specific widget refresh.

### Question 486: Design a “Recently Used” item system.

**Answer:**
*   **LRU:** Client stores list `[ID1, ID2, ID3]` in LocalStorage.
*   **Sync:** Periodically push to `UserPreferences` DB column.
*   **Privacy:** Allow user to clear history.

### Question 487: Build an in-app notification center with read/unread sync.

**Answer:**
*   **Schema:** `Notifications` (ID, UserID, IsRead, Payload).
*   **Sync:**
    *   `GET /notifications?since=timestamp`.
    *   `POST /notifications/read { ids: [...] }`.
*   **Badge:** `COUNT(WHERE IsRead=False)`.

### Question 488: How to implement real-time typing indicators in chat?

**Answer:**
(See Q163). Ephemeral WebSocket events. No DB storage.

### Question 489: Design an infinite scrolling backend.

**Answer:**
*   **Cursor:** `GET /feed?cursor=msg_123`.
*   **Query:** `SELECT * ... WHERE id < 123 LIMIT 20`.
*   **Why:** Better than Offset (no duplicate items if new items arrive at top).

### Question 490: Build a clipboard history syncing system across devices.

**Answer:**
*   **Security:** E2E Encryption (Key generated from Password). Server sees encrypted blobs.
*   **Limit:** Keep last 50 items.
*   **Push:** Copy on Phone -> Push to Server -> Server Pushes to Laptop (FCM).

---

## 🔸 Analytics & Insights Systems (Questions 491-500)

### Question 491: Design a funnel tracking system.

**Answer:**
(Home -> Search -> Cart -> Checkout).
*   **Log:** Users emit events with `SessionID`.
*   **Query:**
    *   Count distinct `SessionID` where Event=Home.
    *   Count distinct `SessionID` where Home AND Search (Sequence Match).
*   **Tool:** ClickHouse / Mixpanel.

### Question 492: Build an ad performance dashboard backend.

**Answer:**
*   **Dimensions:** Campaign, Creative, Geo.
*   **Metrics:** Impressions, Clicks, Cost.
*   **Pre-Aggregation:**
    *   Rollup every minute: `Dim(C1, G1) -> Imp: 100, Clicks: 2`.
    *   Store in Druid/Pinot for sub-second querying.

### Question 493: Design a cohort analysis pipeline.

**Answer:**
(Retention: Users who joined in Jan, how many active in Feb?).
*   **Bitmaps:**
    *   `Active_Jan`: `101010...`
    *   `Active_Feb`: `100010...`
    *   `Retained`: `AND` operation.
    *   `Count`: PopCount.
*   **Efficiency:** Redis Bitmaps are extremely fast for this.

### Question 494: How to build a delayed-event ingestion pipeline?

**Answer:**
(Mobile devices upload logs when Wi-Fi connects).
*   **Problem:** Files arrive 2 days late.
*   **Partitioning:** Use `EventTime` (when it happened) not `IngestTime` (when it arrived).
*   **Idempotency:** Re-process the day's partition if late data arrives.

### Question 495: Design a clickstream data storage system.

**Answer:**
*   **Volume:** Massive (TBs).
*   **Format:** Parquet/Avro (Columnar, Compressed).
*   **Partition:** `s3://logs/date=2023-01-01/hour=10/`.
*   **Query:** Athena / Presto / SparkSQL.

### Question 496: Design a time-series aggregation engine.

**Answer:**
*   **Downsampling:**
    *   Raw: 10s resolution (Keep 7 days).
    *   Rollup 1h: Avg/Max/Min (Keep 1 year).
    *   Rollup 1d: Keep forever.
*   **Automatic:** InfluxDB Continuous Queries / TimescaleDB Continuous Aggregates.

### Question 497: How to ensure accuracy in sampled analytics?

**Answer:**
*   **Sampling:** Log 1% of requests.
*   **Upscale:** Multiply results by 100.
*   **Problem:** Rare events (Errors) might be missed.
*   **Adaptive Sampling:** Sample 100% of Errors, 1% of Successes. Store sampling rate `rate=100` in the record to un-weight later.

### Question 498: How would you process multi-tenant analytics securely?

**Answer:**
*   **Row Level Security (RLS):**
    *   Postgres/Redshift: `CREATE POLICY ... USING (tenant_id = current_user_tenant)`.
*   **Physical:** Separate table per tenant (if few large tenants).

### Question 499: Build a sentiment analysis dashboard with trending topics.

**Answer:**
*   **Stream:** Twitter Firehose.
*   **NLP:** BERT model inference (Classify Positive/Negative).
*   **Aggregator:** Sliding window (1hr). Group by `Topic`. Avg(Sentiment).
*   **Visualization:** Word Cloud.

### Question 500: How to design an “anomaly detection” system for business metrics?

**Answer:**
*   **Model:** Holt-Winters (Seasonality). detects if "Today's Traffic" deviates from "Last 4 Tuesdays".
*   **Z-Score:** `(Current - Mean) / StdDev`. If > 3, Alarm.
*   **Prophet:** Facebook's library for forecasting time series.
