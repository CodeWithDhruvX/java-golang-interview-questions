## 🔸 Healthcare, Legal & Education Domains (Questions 601-610)

### Question 601: Design an appointment booking system for hospitals.

**Answer:**
*   **Entities:** `Doctor`, `Patient`, `Slot`, `Appointment`.
*   **Concurrency:** (See Q282/Q472). `SELECT FOR UPDATE` on Slot row.
*   **Features:**
    *   **Notification:** SMS reminder 1 day before.
    *   **Waitlist:** Priority Queue.

### Question 602: How would you build a prescription refill tracker?

**Answer:**
*   **State:** `Active`, `RefillRequested`, `Approved`, `Dispensed`.
*   **Integration:** Pharmacy API (ePrescribing standards like HL7/FHIR).
*   **Alert:** "Refill Due" calculated based on `LastDispensedDate` + `DaySupply`.

### Question 603: Build a vaccination record platform with auditability.

**Answer:**
*   **Integrity:** Blockchain / Immutable Ledger (QLDB) to prevent tampering.
*   **Identity:** Hash(SSN + Name) as User Key.
*   **Query:** "Show records for UserID 123" -> Verifiable History.

### Question 604: Design a HIPAA-compliant messaging app for doctors.

**Answer:**
*   **Encryption:** E2EE (Signal Protocol) mandatory. Server cannot see messages.
*   **Audit:** Log *metadata* (Dr A spoke to Dr B at Time T) but not content.
*   **Storage:** Ephemeral. Messages deleted after 30 days.
*   **Device:** Remote wipe capability if phone lost.

### Question 605: Build a legal document workflow & e-signing system.

**Answer:**
(See Q281).
*   **Workflow:** `Draft -> Review -> Sign -> Notarize`.
*   **Audit:** Record IP, timestamp, and Email verification for every step.
*   **Storage:** WORM (Write Once Read Many).

### Question 606: Design an online examination platform with proctoring.

**Answer:**
*   **Scale:** 100k students start at 9:00 AM. (Thundering Herd).
*   **Proctoring:**
    *   **WebRTC:** Stream webcam to Proctor server.
    *   **AI:** Face detection (Is there 1 person? Are they looking away?).
*   **Save:** Auto-save answers to Redis every 30s.

### Question 607: How to prevent cheating in online exams?

**Answer:**
*   **Browser:** Focus Event detection (Did user Alt-Tab?).
*   **System:** Lockdown Browser (Kills external processes like Screen Sharing).
*   **Network:** Detect multiple logins.

### Question 608: Design a medical image storage & retrieval system.

**Answer:**
(PACS - Picture Archiving and Communication System).
*   **Format:** DICOM.
*   **Storage:** S3 (Large blobs).
*   **Viewing:** Tile-based rendering (Like Google Maps). Backend renders high-res tiles requested by frontend.

### Question 609: Build a therapy session scheduling and journal app.

**Answer:**
*   **Journal:** Encrypted Client-Side with User's Password. Even backend can't read.
*   **Schedule:** Recurring appointments (Every Wed 3 PM).
*   **Video:** WebRTC P2P for privacy.

### Question 610: Design a patient-doctor chat system with language translation.

**Answer:**
*   **Pipeline:**
    1.  User sends "Hola".
    2.  Server calls Translate API -> "Hello".
    3.  Store original and translated.
    4.  Send "Hello" to Doctor.
*   **Latancy:** < 1s acceptable.

---

## 🔸 API Design and Management (Questions 611-620)

### Question 611: How would you implement API rate limiting per user and API key?

**Answer:**
(See Q312).
*   **Composite Key:** `limit:apikey:123` and `limit:user:456`.
*   **Check:** Both buckets must have tokens.

### Question 612: Design an API gateway with auth, logging, and caching.

**Answer:**
(See Q371).
*   **Kong / APIGee:**
    *   **Auth:** JWT validation.
    *   **Log:** Async push to Splunk.
    *   **Cache:** Redis-based response caching (for GETs).

### Question 613: How to support API versioning for backward compatibility?

**Answer:**
(See Q172).

### Question 614: Design a public-facing developer portal for APIs.

**Answer:**
*   **Docs:** Swagger UI / Redoc (Auto-generated).
*   **Console:** "Try it out" button (Proxies request to Sandbox).
*   **Dashboard:** Usage charts, Billing history, API Key management.

### Question 615: How to throttle APIs based on time of day or region?

**Answer:**
*   **Policy:** "US-East Limit: 1000 RPS (Day), 100 RPS (Night)".
*   **Time:** Server checks Local Time of Req Origin.
*   **Config:** Distributed Config Store (Consul) updates limit rules dynamically.

### Question 616: Design a webhook delivery system with retries and dead letter queues.

**Answer:**
(See Q179).

### Question 617: How do you build an internal API dependency graph?

**Answer:**
*   **Tracing:** Application Performance Monitoring (APM) agents (Datadog/NewRelic) automatically map `Service A calls Service B`.
*   **Static Analysis:** Scan code for HTTP calls to known endpoints.

### Question 618: Build an API change detection & alerting platform.

**Answer:**
*   **Contract:** OpenAPI Spec stored in Git.
*   **Diff:** CI Pipeline runs `openapi-diff`.
*   **Alert:** If `Breaking Change` detected (Deleted field) -> Fail Build + Notify Team.

### Question 619: How would you design a GraphQL gateway over microservices?

**Answer:**
*   **Federation (Apollo):**
    *   Service A (User) defines `type User`.
    *   Service B (Reviews) extends `type User { reviews: [Review] }`.
*   **Gateway:** Stitches schemas. Queries A for User, then B for Reviews.

### Question 620: Implement usage analytics for each endpoint in an API platform.

**Answer:**
*   **middleware:** Start timer. On response, emit metric: `api_request_duration{endpoint="/users", status="200"}`.
*   **Aggregator:** Prometheus.

---

## 🔸 Data Quality & Integrity (Questions 621-630)

### Question 621: Design a system for real-time data validation.

**Answer:**
*   **Interceptor:** Kafka Stream Processor.
*   **Rules:** JSON Schema / Great Expectations.
*   **Action:**
    *   Valid -> Transformation Topic.
    *   Invalid -> Dead Letter Topic.

### Question 622: How would you ensure consistency in duplicate data across services?

**Answer:**
*   **Master Data Management (MDM):**
    *   Service A has "John Doe".
    *   Service B has "J. Doe".
*   **Resolution:** Central MDM Service assigns `GlobalID`. Updates A and B to link to `GlobalID`.

### Question 623: Design a service for cleaning and deduplicating customer records.

**Answer:**
*   **Batch:** Spark Job runs nightly.
*   **Logic:** Fuzzy Matching (Levenshtein Distance on Name + Exact Match on Phone).
*   **Merge:** Create "Golden Record" with best data from duplicates.

### Question 624: Build a platform for automated schema consistency checks.

**Answer:**
*   **Registry:** Kafka Schema Registry.
*   **Enforcement:** Producer CANNOT publish message if schema doesn't match Registry.
*   **Migration:** CI Check ensures DB Migration SQL is compatible with current Code.

### Question 625: Design a data poisoning detection mechanism in ML pipelines.

**Answer:**
*   **Outlier Detection:** Before training, run Isolation Forest methods to find anomalous training examples.
*   **Provenance:** Track origin of every data point. If Source X provides 90% bad data, block Source X.

### Question 626: Build a distributed checksum validation system.

**Answer:**
*   **Ingest:** Calculate MD5/SHA256 of file.
*   **Transfer:** Send file + Hash.
*   **Verify:** Receiver calculates Hash. If mismatch -> Retransmit.
*   **Periodic:** Background "Scrubbing" task reads disk blocks and verifies CRC.

### Question 627: How do you detect silent data corruption?

**Answer:**
*   **block csum:** Filesystems (ZFS) store checksums.
*   **App Level:** Store `Hash(Row_Content)` in a separate column. On Read, verify `Hash(cols) == StoredHash`.

### Question 628: Implement a rollback-safe write-ahead log.

**Answer:**
*   **Structure:** `LSN (Log Sequence Number) | PrevLSN | TransactionID | Operation`.
*   **Checkpoint:** Truncate log up to Checkpoint LSN.
*   **Rollback:** Read backwards from End to Start, undoing changes.

### Question 629: Design a “data quarantine” zone for suspect records.

**Answer:**
*   **Staging Area:** S3 Bucket `quarantine/`.
*   **Review:** UI to inspect JSON.
*   **Action:** "Fix & Replay" (Edit JSON -> Push to Input Queue) or "Discard".

### Question 630: How do you verify external data imports are safe?

**Answer:**
*   **Sandbox:** Process in isolated container.
*   **Limits:** Check file size, row count boundaries.
*   **Sanitization:** Strip HTML/Script tags (XSS prevention).

---

## 🔸 Testing, Monitoring, and Observability (Questions 631-640)

### Question 631: Design a chaos testing platform.

**Answer:**
(See Q205).

### Question 632: Build a synthetic monitoring tool for uptime checks.

**Answer:**
(See Q349).
*   **Runner:** AWS Lambda scheduled every 1 min.
*   **Checks:** HTTP Get, DNS Resolve, SSL Cert Expiry check.
*   **Result:** Push metrics to CloudWatch.

### Question 633: Design a traceable, testable deployment pipeline.

**Answer:**
*   **Traceability:** Git Commit -> Build ID -> Docker Image Tag -> K8s Deployment -> Pod.
*   **Audit:** "Who clicked Deploy?" logged.

### Question 634: How do you ensure test coverage for service-to-service interactions?

**Answer:**
*   **Contract Testing (Pact):**
    *   Consumer defines expectations.
    *   Provider verifies it meets expectations.
*   **Mocking:** Use WireMock in Integration tests.

### Question 635: Design a zero-downtime release system.

**Answer:**
(See Q307 rolling updates).

### Question 636: Build a real-time alerting system for performance regressions.

**Answer:**
*   **Baseline:** Calculate avg latency for last 7 days.
*   **Compare:** If `Current_1h_Avg > Baseline * 1.5` -> Alert "Performance degraded by 50%".
*   **Deploy:** Compare Canary vs Baseline.

### Question 637: How to record detailed request/response traces for debugging?

**Answer:**
*   **Sampling:** 100% too expensive. Sample 1% of success, 100% of errors.
*   **Storage:** Large blobs (Bodies). Store in S3, indexed by TraceID in Elasticsearch.

### Question 638: Design a system to test rollback and forward compatibility.

**Answer:**
*   **Forward:** Test `Client_New` against `Server_Old`.
*   **Backward:** Test `Client_Old` against `Server_New`.
*   **Suite:** Matrix of Client(v1, v2) x Server(v1, v2).

### Question 639: Build a shadow traffic replay system for staging environments.

**Answer:**
*   **Tool:** GoReplay / Envoy Shadowing.
*   **Prod:** Copy request stream -> Send to Staging (Fire and Forget).
*   **Staging:** Process request. Discard response/side-effects (Mock payment gateways).

### Question 640: Design a service dependency graph with fault isolation.

**Answer:**
*   **Discovery:** Service Mesh (Istio) builds map automatically.
*   **Isolation:** (See Q203 Bulkheading).

---

## 🔸 Security & Compliance (Questions 641-650)

### Question 641: Design a secure file upload service.

**Answer:**
(See Q360).

### Question 642: How would you encrypt large datasets with minimal latency?

**Answer:**
*   **Envelope Encryption:**
    1.  Generate DEK (Data Encryption Key) locally. Fast.
    2.  Encrypt Data with DEK (AES-GCM).
    3.  Encrypt DEK with KEK (Key Encryption Key from KMS).
    4.  Store `EncryptedData + EncryptedDEK`.

### Question 643: Design a secure OAuth2 flow for mobile and web.

**Answer:**
*   **PKCE (Proof Key for Code Exchange):**
    1.  Client generates `CodeVerifier` and `CodeChallenge`.
    2.  Send `Challenge` in Auth Request.
    3.  Send `Verifier` in Token Request.
    4.  Server hashes Verifier. If matches Challenge, issue Token.
    *   Prevents Code Interception attacks.

### Question 644: Implement data access policies based on roles and geography.

**Answer:**
*   **ABAC (Attribute Based Access Control):**
    *   `Allow if User.Role == 'HR' AND User.Location == 'EU' AND Data.Location == 'EU'`.

### Question 645: Build a system to audit and revoke stale credentials.

**Answer:**
*   **Scanner:** Checks `LastUsedDate` of IAM Keys / API Tokens.
*   **Policy:** If `Now - LastUsed > 90 days`:
    1.  Disable Key (Soft).
    2.  Notify User.
    3.  Delete after 7 days (Hard).

### Question 646: How would you secure long-lived background processes?

**Answer:**
*   **Identity:** Give the Process its own Identity (Service Account).
*   **Least Privilege:** Grant ONLY permission to read Queue and write DB. No SSH, no S3 admin.
*   **Rotation:** Rotate Service Account keys automatically.

### Question 647: Design a phishing detection system for emails.

**Answer:**
*   **Headers:** Check SPF, DKIM, DMARC. (Spoofing check).
*   **Content:** ML Model analyzes Text and Links ("Click here to reset").
*   **Domain:** Check if domain `g0ogle.com` is distinct from `google.com` (Levenshtein check).

### Question 648: Build a service to manage and rotate secrets.

**Answer:**
(See Q354 Vault).

### Question 649: How to detect unusual login patterns across geographies?

**Answer:**
*   **Impossible Travel:**
    *   Login 1: London, 10:00 AM.
    *   Login 2: New York, 10:05 AM.
    *   Distance: 3000 miles. Time: 5 mins. Speed > Plane? Yes -> Alert.

### Question 650: Design a system for 2FA backup codes and recovery.

**Answer:**
*   **Generation:** 10 random 8-digit codes.
*   **Storage:** `Hash(Code)` in DB.
*   **Usage:** User enters code. Server verifies Checksum. Marks code as `Used`.
*   **Rate Limit:** Max 3 attempts per hour (Prevent brute force).
