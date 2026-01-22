## 🔸 Search, Filtering, and Indexing (Questions 651-660)

### Question 651: Design a full-text search engine with autocomplete.

**Answer:**
*   **Indexing:** Inverted Index (`Term -> DocIDs`).
*   **Autocomplete:** Trie System.
    *   **Prefix Search:** O(L) to find prefix node.
    *   **Cache:** Store Top 10 hottest suggestions at each node.
    *   **Tool:** Redis (ZSET) or Elasticsearch (Completion Suggester).

### Question 652: Build a system for personalized search ranking.

**Answer:**
*   **Retrieval (L1):** Get Top 500 relevant items (BM25 score).
*   **Re-Ranking (L2):** Learning to Rank (LTR) model (LambdaMART).
    *   Features: User History, CTR, Recency.
    *   Output: Sorted Top 500.

### Question 653: How to implement federated search across services?

**Answer:**
*   **Aggregator:** Search Service.
*   **Fan-out:** Call `UserSearchService`, `ProductSearchService`, `OrderSearchService`.
*   **Merge:** Receive `[Results, Score]`. Normalize scores. Sort.
*   **Timeout:** If `ProductSearch` is slow, return partial results from others.

### Question 654: Design an e-commerce filter engine (facets, range, categories).

**Answer:**
*   **Bitmaps:** `Color_Red: 101`, `Size_M: 011`.
*   **Query:** "Red AND M" -> `101 & 011 = 001`. (Item 3).
*   **Roaring Bitmaps:** Compressed bitmap format used by Lucene/Elasticsearch for speed.

### Question 655: Implement typo-tolerant search indexing.

**Answer:**
*   **Fuzzy Search:** Levenshtein Automaton.
*   **N-Grams:** Index "iphone" as `ip, ph, ho, on, ne`.
    *   Query "iphoe" -> `ip, ph, ho, oe`.
    *   High overlap coefficient = Match.

### Question 656: Design a system for trending keyword detection.

**Answer:**
*   **Stream:** Kafka topic of Search Queries.
*   **Processing:**
    *   Sliding Window (1 hr).
    *   Count-Min Sketch (Approximate Frequency).
*   **Trending:** `Velocity = (CurrentCount - LastHourCount)`. Sort by Velocity.

### Question 657: How would you enable real-time reindexing without downtime?

**Answer:**
(See Q470 - Aliasing).

### Question 658: Build a tagging + filtering engine for massive datasets.

**Answer:**
*   **Data Structure:** Postings List.
    *   `Tag:AI -> [Doc1, Doc5, Doc9...]`.
*   **Compression:** Delta Encoding (Store `1, +4, +4` instead of `1, 5, 9`).
*   **Intersection:** Skip Pointers to jump ahead during list intersection.

### Question 659: Design a “related searches” suggestion system.

**Answer:**
*   **Co-occurrence Matrix:**
    *   Users who searched X also searched Y within 5 minutes.
    *   `Map<Query, Map<RelatedQuery, Count>>`.
*   **Graph:** Random Walk on Query Graph.

### Question 660: How to support multi-language search effectively?

**Answer:**
*   **Analysis:** Detect Language (`cld2` library).
*   **Index:** Separate indices: `products_en`, `products_fr`.
*   **Analyzer:** Use language-specific Stemmers (Snowball) and Stop-words.

---

## 🔸 Rules Engines & Automation (Questions 661-670)

### Question 661: Design a rules engine for fraud detection.

**Answer:**
*   **DSL:** Define rules in JSON/YAML. (`Amount > 1000 AND Country != US`).
*   **AST:** Parse DSL into Abstract Syntax Tree.
*   **Evaluation:** Traverse AST with Context (`Transaction` object).
*   **Rete Algorithm:** Optimization for matching many patterns against many objects.

### Question 662: Build a workflow automation engine like Zapier.

**Answer:**
*   **Trigger:** Webhook / Polling.
*   **Action:** API Client.
*   **Pipeline:** `Trigger -> Filter -> Transform -> Action`.
*   **State:** Store `StepID` in DB. Resume from failure.

### Question 663: Implement a rule-based alerting system.

**Answer:**
*   **Evaluator:** Prometheus Alertmanager.
*   **Tick:** Every 15s.
*   **Condition:** `rate(errors[1m]) > 5`.
*   **Grouping:** Group alerts by `ClusterID`. Wait 1m to batch.

### Question 664: Design a user-triggered automation builder.

**Answer:**
*   **UI:** Drag and drop blocks.
*   **Backend:** Save as Directed Acyclic Graph (DAG) JSON.
*   **Execution:** Topological Sort of DAG -> Execute sequentially.

### Question 665: Build a no-code workflow designer backend.

**Answer:**
(Similar to Q664).
*   **Validation:** Check for Cycles in Graph. Check Types (Output of Step A must match Input of Step B).

### Question 666: Design a credit card fraud rules platform.

**Answer:**
*   **Hot Deploy:** Rules stored in Redis/Cache.
*   **Shadow Mode:** Test new rule on live traffic but log result instead of blocking. Measure False Positives.
*   **Priority:** `Block List > Velocity Check > ML Score`.

### Question 667: Implement rule versioning with rollback support.

**Answer:**
*   **Immutable:** `RuleID: 123`, `Version: 5`.
*   **Reference:** `Policy: Fraud_Check` points to `Rule: 123_v5`.
*   **Rollback:** Update Policy to point to `Rule: 123_v4`.

### Question 668: Build an engine that triggers actions based on threshold breaches.

**Answer:**
*   **Input:** Metric Stream using Kafka.
*   **State:** Flink `ValueState`.
*   **Logic:**
    *   `Current = GetState()`.
    *   `Current += event.value`.
    *   `if Current > Threshold -> Emit Alert`.

### Question 669: Design a low-latency decision engine for recommendations.

**Answer:**
*   **Pre-compute:** "For User X, Candidates are [A, B, C]". Store in Redis.
*   **Runtime:** Only apply filtering (Out of stock) and light ranking. Return < 10ms.

### Question 670: How would you design a conflict resolution system for overlapping rules?

**Answer:**
*   **Specificity:** Rule `User=Alice` beats Rule `Group=Engineering`.
*   **Priority:** Explicit integer rank (`100` beats `10`).
*   **Order:** First Match Wins.

---

## 🔸 Document Processing & Content Systems (Questions 671-680)

### Question 671: Build a resume parsing and ranking engine.

**Answer:**
*   **Parse:** OCR / PDFToText.
*   **Entity Extraction:** NER (Named Entity Recognition) to find `Skills`, `Experience`.
*   **Ranking:** Vector Similarity. Embedding(`JobDesc`) dot Embedding(`Resume`).

### Question 672: Design a contract redlining and version tracking system.

**Answer:**
*   **Diff:** `diff-match-patch` algorithm (Google).
*   **Storage:** Save XML/JSON representing the doc structure.
*   **Track Changes:** Annotate text ranges with `DeletedBy: UserA`.

### Question 673: How would you convert scanned documents to searchable formats?

**Answer:**
*   **Pipeline:** S3 Upload -> Lambda (Tesseract OCR) -> HOCR (HTML with coords).
*   **Index:** Elasticsearch ingest attachment plugin.

### Question 674: Design a collaborative Markdown document editor.

**Answer:**
(See Q522).

### Question 675: Build a service to extract tables and charts from PDFs.

**Answer:**
*   **Tools:** Amazon Textract / Tabula.
*   **Heuristic:** Look for grid lines.
*   **Output:** Convert to CSV/JSON.

### Question 676: How to handle bulk uploads and format normalization?

**Answer:**
*   **Queue:** Bulk Upload = Background Job.
*   **Normalization:**
    *   CSV? Excel? JSON?
    *   Convert all to intermediate DataFrame.
    *   Apply Schema Validation.

### Question 677: Design a plagiarism detection platform.

**Answer:**
*   **Fingerprint:** Winnowing Algorithm (Select k-grams).
*   **DB:** Rolling Hash Database.
*   **Compare:** Find sequences of matching hashes.

### Question 678: Build a CMS for publishing across mobile and web.

**Answer:**
(Headless CMS).
*   **Content:** JSON API (`GET /posts/1`).
*   **Presentation:** Decoupled. React renders JSON for Web. SwiftUI renders JSON for iOS.

### Question 679: Design a privacy-aware document sharing system.

**Answer:**
*   **ACL:** `Resource: Doc1`, `User: Alice`, `Action: Read`.
*   **Time-bomb:** `ExpiresAt` field. Background job revokes access.
*   **Watermark:** Embed `User: Alice` invisible watermark in PDF download.

### Question 680: Build a legal compliance document delivery tracker.

**Answer:**
*   **Requirement:** Proof of Delivery.
*   **Email:** Track Open Pixel.
*   **Portal:** User must click "I Acknowledge".
*   **Audit:** Store `UserAgent`, `IP`, `Timestamp` of click in immutable ledger.

---

## 🔸 Chain of Events & State Machines (Questions 681-690)

### Question 681: Build a state machine for user onboarding.

**Answer:**
*   **States:** `Created -> EmailVerified -> ProfileFilled -> TeamJoined -> Active`.
*   **Events:** `VerifyEmail`, `SaveProfile`.
*   **Guard:** Can't go `Created -> Active` without `EmailVerified`.

### Question 682: Design a job status tracker with retry and escalation.

**Answer:**
*   **Retry:** State `Failed`. Transition to `Retrying`. Increment `RetryCount`.
*   **Escalate:** If `RetryCount > 3`, transition to `Escalated`. Notify Manager.

### Question 683: Implement a finite state machine for delivery systems.

**Answer:**
*   **Order:** `Placed -> Preparing -> PickedUp -> OutForDelivery -> Delivered`.
*   **GPS:** Driver App sends `Location`. Geofencing triggers state changes.

### Question 684: Build a versioned state transition audit trail.

**Answer:**
*   **Table:** `Transitions` (`EntityID`, `OldState`, `NewState`, `Trigger`, `Timestamp`).
*   **Query:** Reconstruct history.

### Question 685: How to validate illegal transitions in distributed systems?

**Answer:**
*   **Optimistic Lock:** `UPDATE orders SET state='Shipped', v=2 WHERE id=1 AND state='Paid' AND v=1`.
*   **Result:** If 0 rows updated, transition failed (State was not 'Paid').

### Question 686: Design a resume-from-checkpoint download system.

**Answer:**
*   **Head:** `HEAD /file`. Get `Content-Length: 1000`.
*   **Range:** `GET /file Range: bytes=0-500`.
*   **Resume:** Connection drops. Client checks local file size (500). Sends `GET /file Range: bytes=500-1000`.

### Question 687: How would you visualize system transitions and workflows?

**Answer:**
*   **Graphviz:** Generate DOT file from State Machine definition.
*   **Sankey Diagram:** Flow of users between states.

### Question 688: Build a system to pause/resume workflows on external triggers.

**Answer:**
(See Q563).
*   **Wait Node:** Workflow Engine supports "Wait for Event". Persists state to DB. Hibernates process.
*   **Wake:** Webhook matches Event ID. Restores process from DB.

### Question 689: How to build real-time state dashboards for business operations?

**Answer:**
*   **CDC:** Debezium listens to DB changes.
*   **Aggregator:** Flink aggregates `Count(State)` per minute.
*   **Push:** WebSocket to Dashboard.

### Question 690: Implement state reconciliation across distributed replicas.

**Answer:**
*   **Merkle Tree:** (See Q243).
*   **Vector Clock:** Detect conflict.
*   **LWW:** Last Write Wins (if acceptable).

---

## 🔸 Time-Sensitive & Temporal Systems (Questions 691-700)

### Question 691: Build a daily digest email generator.

**Answer:**
*   **Events:** Store `Activity` throughout day.
*   **Cron:** At User's Local 9 AM.
*   **Batch:** Scan `Activity` for last 24h.
*   **Render:** Generate HTML. Send.

### Question 692: Design a system to expire items at dynamic times.

**Answer:**
*   **Redis:** `EXPIREAT key timestamp`.
*   **DynamoDB:** `TTL` attribute. AWS deletes within 48h.
*   **Precise:** Priority Queue (Min Heap). background thread peeks top.

### Question 693: How to implement TTL with accuracy and efficiency?

**Answer:**
*   **Lazy:** Check TTL on Read. If expired, return null and delete.
*   **Active:** Background sweeper deletes expired keys (Probabilistic approach like Redis).

### Question 694: Build a high-resolution calendar scheduler.

**Answer:**
*   **Resolution:** 1 minute.
*   **Storage:** `Start` and `End` timestamps (BIGINT epoch).
*   **Overlap:** `WHERE A.Start < B.End AND A.End > B.Start`. Index on `(Start, End)`.

### Question 695: Design a system for recurring job execution with drift correction.

**Answer:**
*   **Drift:** `Sleep(24h)` accumulates error (24h + 10ms execution).
*   **Correction:** Calculate `NextRun = StartTime + N * Interval`. Sleep `NextRun - Now`.

### Question 696: How to detect temporal anomalies (e.g., no events)?

**Answer:**
(Dead Man's Switch).
*   **Expected:** 1 event per minute.
*   **Monitor:** `time() - last_seen_time > 5min`. Alert.

### Question 697: Implement a system for calculating historical state timelines.

**Answer:**
*   **Bi-Temporal Modeling:**
    *   `ValidTime`: When the fact is true in real world.
    *   `TransactionTime`: When the DB recorded it.
*   **Query:** "What did we *think* the address was yesterday?"

### Question 698: How would you enable users to “rewind” system state?

**Answer:**
*   **Event Sourcing:** Replay events from 0 to T.
*   **Snapshotting:** Restore snapshot closest to T. Replay events from Snapshot to T.

### Question 699: Design a time-aware feed with relevance and freshness.

**Answer:**
*   **Decay Function:** `Score = Relevance / (Age + 1)^Gravity`.
*   **Hacker News Algo:** New items bubble up. Old items sink rapidly.

### Question 700: How to manage time zones and DST in event scheduling?

**Answer:**
*   **Storage:** Always store UTC.
*   **Display:** Convert to User Local.
*   **Recurring:** Store "9 AM America/New_York".
    *   Compute UTC for next instance.
    *   Handles DST transition automatically (9 AM might be 13:00 or 14:00 UTC).
