## 🔸 Security-Centric Designs (Questions 351-360)

### Question 351: Design a 2FA system.

**Answer:**
*   **TOTP (Time-based One-Time Password):**
    *   **Setup:** Server generates Secret Key (`SK`). Client scans QR.
    *   **Login:** Client and Server compute `HMAC(SK, TimeBucket)`. If match -> specific user.
*   **SMS/Email:**
    *   Generate random code -> Store in Redis (`userID:code` TTL 5m) -> Send via Twilio/SendGrid.
    *   Verify: Input code matches Redis value.

### Question 352: How to securely store access tokens?

**Answer:**
*   **Browser:**
    *   **HttpOnly Cookie:** Best against XSS. JS cannot read it. Vulnerable to CSRF (Mitigation: SameSite=Strict).
    *   **LocalStorage:** Vulnerable to XSS (Malicious lib can read `localStorage.accessToken`).
*   **Recommendation:** HttpOnly Cookie for Refresh Token. Short-lived Access Token in Memory.

### Question 353: What’s the difference between OAuth2 and OpenID Connect?

**Answer:**
*   **OAuth2 (Authorization):** "I want to access your photos on Google Drive." (Access Token). Delegated Access.
*   **OIDC (Authentication):** "I want to know WHO you are." (ID Token). Identity Layer on top of OAuth2. Returns JWT with user profile.

### Question 354: Design a secrets management service.

**Answer:**
(e.g., Vault).
*   **Storage:** Encrypted backend (Consul/Etcd).
*   **Access:** Client authenticates (K8s ServiceAccount) -> Gets temporary token.
*   **Dynamic Secrets:**
    *   Vault creates a *temporary* Database User for the App.
    *   Lease expires -> Vault deletes the DB User.
*   **Unsealing:** Shamir's Secret Sharing (Need 3 of 5 keys to start Vault).

### Question 355: How would you design an audit trail for admin actions?

**Answer:**
*   **Middleware:** Intercept all mutable requests (`POST`, `PUT`, `DELETE`).
*   **Context:** Extract `AdminID`, `TargetResource`, `OldValue`, `NewValue`.
*   **Sink:** Asynchronously push to immutable ledger (S3 Object Lock / Blockchain).
*   **Alerting:** Trigger alert if `DeleteUser` called > 5 times in 1 min.

### Question 356: Design a DDoS protection layer.

**Answer:**
*   **Edge (Cloudflare):** Absorbs Volumetric attacks (L3/L4).
*   **WAF (L7):** Blocks SQL Injection, Bad User-Agents.
*   **Rate Limiting:** IP-based.
*   **Challenge:** CAPTCHA / JS Challenge for suspicious traffic.
*   **Infrastructure:** Auto-scaling groups to absorb legitimate traffic spikes.

### Question 357: How do you implement permission inheritance?

**Answer:**
*   **RBAC Hierarchy:**
    *   `Admin` inherits `Editor` inherits `Viewer`.
*   **Groups:**
    *   User is in `TeamLeader` group.
    *   `TeamLeader` group is member of `Employee` group.
*   **Resolution:** Graph traversal (BFS) to check if User reaches required Permission.

### Question 358: How to detect and prevent replay attacks?

**Answer:**
*   **Nonce:** Client sends unique `Nonce` (UUID). Server tracks "seen nonces" in Redis. Rejects duplicates.
*   **Timestamp:** Request must include `Timestamp`. Server rejects if `Now - Timestamp > 5 min` (Window of acceptance).
*   **Signature:** Sign `(Body + Nonce + Timestamp)` to prevent tampering.

### Question 359: How would you secure WebSockets?

**Answer:**
*   **Handshake Auth:**
    *   Client connects `wss://api.com?token=JWT`.
    *   Server validates JWT before upgrading HTTP to WS.
*   **Origin Check:** Validate `Origin` header to prevent CSWSH (Cross-Site WebSocket Hijacking).
*   **WSS:** Always use TLS.

### Question 360: How to design secure file uploads?

**Answer:**
1.  **Validation:** Check Magic Bytes (don't trust extension). Re-encode image (strips malicious payloads).
2.  **Storage:** Store in S3, not local disk.
3.  **Permissions:** Make S3 bucket Private. Serve via Presigned URL.
4.  **Sandbox:** Run virus scan (ClamAV) in isolated container before marking "Safe".

---

## 🔸 Component and Data Modeling (Questions 361-370)

### Question 361: How would you model user roles in a large SaaS platform?

**Answer:**
**Schema:**
*   `User` (ID, Name)
*   `Organization` (ID, Plan)
*   `Role` (ID, Name: "Admin", OrgID)
*   `Permission` (ID, Code: "billing.read")
*   `RolePermission` (RoleID, PermissionID)
*   `UserOrgRole` (UserID, OrgID, RoleID) -> *User can be Admin in Org A but Viewer in Org B*.

### Question 362: Model a product inventory with variants and warehouses.

**Answer:**
*   `Product` (T-Shirt).
*   `Variant` (Red, Size M).
*   `Warehouse` (NY, LA).
*   `Inventory`:
    *   `VariantID`
    *   `WarehouseID`
    *   `Quantity` (Available)
    *   `Reserved` (In Carts)

### Question 363: How to model time-based entitlements?

**Answer:**
(e.g., Netflix Subscription, free trial).
*   `Subscription`:
    *   `UserID`
    *   `PlanID`
    *   `Status` (Active, Canceled).
    *   `CurrentPeriodStart` (2023-01-01).
    *   `CurrentPeriodEnd` (2023-02-01).
*   **Check:** `if Status == Active AND Now < CurrentPeriodEnd`.

### Question 364: Model a data-sharing agreement between companies.

**Answer:**
(B2B Data Pipe).
*   `Agreement` (SourceOrg, TargetOrg, Expiry).
*   `DataSet` (ID, Schema).
*   `AccessGrant` (AgreementID, DataSetID, Filter: "Region=US").
*   **Enforcement:** API Gateway checks `AccessGrant` before returning data.

### Question 365: Design a dynamic pricing model.

**Answer:**
*   `BasePrice` (Static).
*   `PricingRule`:
    *   `Condition` (JSON Logic: `User.Loyalty == Gold`).
    *   `Action` (`Multiplier: 0.9` or `FlatOff: 10`).
    *   `Priority` (1).
*   **Engine:** Fetch all rules -> Filter applicable -> Apply in Priority order.

### Question 366: Model a messaging inbox with threads and participants.

**Answer:**
*   `Thread` (ID, LastMessageAt).
*   `Participant` (ThreadID, UserID, LastReadAt).
*   `Message` (ThreadID, SenderID, Content, CreatedAt).
*   **Unread Count:** `Wait for Message WHERE ThreadID IN (MyThreads) AND CreatedAt > Participant.LastReadAt`.

### Question 367: Design a schema for customer support tickets.

**Answer:**
*   `Ticket` (ID, Subject, Status, RequesterID, AssigneeID).
*   `Comment` (TicketID, AuthorID, Body, IsInternal).
*   `StatusHistory` (TicketID, FromStatus, ToStatus, ChangedBy, Timestamp).
*   **SLA:** `Ticket.ReplyDueBy`.

### Question 368: Model a booking system with cancellation windows.

**Answer:**
*   `Booking` (ID, ResourceID, Start, End, Status).
*   `Policy` (ResourceID, `RefundPercentage`, `HoursBefore`).
*   **Logic:**
    *   User requests Cancel at `T_cancel`.
    *   `Gap = Booking.Start - T_cancel`.
    *   `Refund = lookup Policy where Gap > HoursBefore`.

### Question 369: Design a permission hierarchy tree.

**Answer:**
(Folder permissions).
*   `Resource` (ID, ParentID, Type).
*   `ACL` (ResourceID, UserID, Permission).
*   **Inheritance:**
    *   Recursive CTE (Common Table Expression) to traverse up: `Result = ACL(ID) OR ACL(ParentID) ...`
    *   Materialized Path: Store path `/root/marketing/2023`. Query `ACL WHERE path matches prefix`.

### Question 370: How would you design an entity history tracker?

**Answer:**
(CDC - Change Data Capture).
*   **Shadow Table:** `Users_History` (Same columns as `Users` + `Version`, `ModifiedBy`).
*   **Trigger:** On Update `Users` -> Insert into `Users_History`.
*   **Hibernate Envers:** Java library that does this automatically.

---

## 🔸 API & Integration Design (Questions 371-380)

### Question 371: How would you build an API gateway?

**Answer:**
*   **Core:** Reverse Proxy (Nginx/Envoy).
*   **Plugins:** Chain of Responsibility Pattern.
    *   Auth Plugin (Validate JWT).
    *   RateLimit Plugin (Redis).
    *   Routing Plugin (Path -> Service).
*   **Config:** Control Plane pushes config to Data Plane (Proxy).

### Question 372: How to design a bulk import/export API?

**Answer:**
*   **Sync:** Bad. (Timeout).
*   **Async Pattern:**
    1.  `POST /export` -> Returns `202 Accepted` + `JobID`.
    2.  Worker generates CSV, uploads to S3.
    3.  `GET /jobs/{id}` -> Returns `Status: Processing`.
    4.  Processing complete -> Returns `Status: Done` + `DownloadURL`.

### Question 373: How do you handle partial failures in batch APIs?

**Answer:**
*   **Design:** `POST /batch`. Body: `[Item1, Item2, Item3]`.
*   **Response (207 Multi-Status):**
    ```json
    [
      {"id": 1, "status": "success"},
      {"id": 2, "status": "error", "msg": "invalid"},
      {"id": 3, "status": "success"}
    ]
    ```
*   **All or Nothing:** Wrap all in one DB transaction. If one fails, Rollback all.

### Question 374: How do you expose webhooks securely?

**Answer:**
(See Q179).
*   **Verify Sender:** HMAC Signature.
*   **Verify Receiver:** Mutual TLS (mTLS).
*   **IP Whitelist:** Allow requests only from known IP ranges (e.g., Stripe IPs).

### Question 375: How would you make APIs backward-compatible?

**Answer:**
*   **Add:** Adding a field is safe. (Clients ignore unknown fields).
*   **Delete/Rename:** Unsafe.
    *   Keep old field. Mark Deprecated.
    *   Populate BOTH old and new fields in backend.
*   **Semantics:** Don't change behavior (e.g., logic of status=Active).

### Question 376: Design an API discovery mechanism.

**Answer:**
*   **Internal:** Service Registry (Consul/K8s DNS).
*   **Public:** Developer Portal (Backstage).
*   **Schema:** Iterate all services -> Fetch `swagger.json` -> Aggregate into central UI.

### Question 377: How would you design async APIs using polling?

**Answer:**
(See Q372). The "Request-Reply via Polling" pattern.

### Question 378: How to build an OAuth2-based authorization flow?

**Answer:**
(Authorization Code Grant).
1.  Client redirects User to Auth Server.
2.  User approves.
3.  Auth Server redirects to Callback URL with `code`.
4.  Client backend POSTs `code` + `client_secret` to Auth Server.
5.  Auth Server returns `access_token` + `refresh_token`.

### Question 379: Design a billing API integration layer.

**Answer:**
*   **Abstraction:** `IMaymentProvider`. Implement `StripeAdapter`, `PayPalAdapter`.
*   **Webhooks:** Normalize incoming webhooks (Charge Succeeded) into internal Event: `PaymentSuccess`.
*   **Reconciliation:** Daily job comparing Our DB vs Stripe Reports to find mismatched states.

### Question 380: How to prevent API abuse from internal clients?

**Answer:**
*   **Quotas:** Even internal teams need quotas (Prevent loop DDOS).
*   **Contracts:** Consumer Driven Contracts. If Team A calls Team B, they must agree on schema.
*   **Circuit Breakers:** If Team A abuses Team B, Team B breaks circuit to protect itself.

---

## 🔸 Real-World Product Backends (Questions 381-390)

### Question 381: Design the backend of a podcast platform.

**Answer:**
*   **Hosting:** S3 for MP3s.
*   **RSS:** Generate RSS XML dynamically or cache it (CDN).
*   **Analytics:**
    *   Client pings every 10s: `Ping(User, Episode, Time)`.
    *   Server Aggregates: `CompletedListen`, `DropOffRate`.
*   **Search:** Transcribe audio (Whisper) -> Index Text in Elasticsearch.

### Question 382: Design a platform like Duolingo (adaptive learning).

**Answer:**
*   **Knowledge Graph:** Dependent concepts (`Words -> Grammar -> Sentence`).
*   **Spaced Repetition:** SuperMemo / Smem2 algorithm. Schedule review of a word based on "Forgetting Curve".
*   **Gamification:** Leaderboards (Redis ZSET). Streaks (Bitmap).

### Question 383: Build the backend of a flash sales system (like limited-time discounts).

**Answer:**
*   **Challenge:** Massive concurrency. "First come first served".
*   **Queue:** Put users in "Waiting Room" (Netomite).
*   **Inventory:** Redis `DECR stock`. If < 0, Sold Out.
*   **Async:** Only winners process to Checkout DB.

### Question 384: Design a loyalty program with points and tiers.

**Answer:**
*   **Event:** `OrderCompleted(Amount)`.
*   **Calculator:** `Points = Amount * TierMultiplier`.
*   **Accumulator:** Update `UserPoints`.
*   **Trigger:** If `UserPoints > GoldThreshold` -> Upgrade Tier -> Send Email.
*   **Expiry:** Scheduled job to expire points > 1 year old.

### Question 385: How would you build a serverless blog CMS?

**Answer:**
*   **Frontend:** Next.js hosted on Vercel/S3.
*   **Backend:** AWS Lambda.
*   **DB:** DynamoDB (Single Table Design).
*   **Images:** S3 + Lambda Trigger (Resize).
*   **Cost:** Near zero when idle.

### Question 386: Design an online judge system like Leetcode.

**Answer:**
*   **Sandbox:** Docker / gVisor / Firecracker. Isolates code execution.
*   **Limits:** `docker run --memory=128m --cpus=0.5`. Time limit via `timeout` command.
*   **Security:** Block syscalls (Networking, Filesystem) using seccomp profile.

### Question 387: Design the backend of a QR code-based payment system.

**Answer:**
*   **QR:** Encodes `MerchantID`.
*   **Flow:**
    1.  User scans QR -> App gets `MerchantID`.
    2.  User enters Amount -> Auth (PIN).
    3.  Backend moves money `User -> Merchant`.
    4.  Notify Merchant (WebSocket/Push).

### Question 388: Build a birthday/anniversary reminder service.

**Answer:**
*   **DB:** `Events` (UserID, Date: MM-DD).
*   **Scheduler:** Daily Scan. `SELECT * FROM Events WHERE Month=Now.Month AND Day=Now.Day`.
*   **Queue:** Push to Notification Queue.
*   **Scale:** Shard by Day of Year (366 shards).

### Question 389: Design an ad delivery engine.

**Answer:**
*   **Bidding:** Real Time Bidding (RTB). < 100ms.
*   **Index:** Inverted Index of Targeting criteria (Age, Geo, Interest).
*   **Selection:** Filter Eligible Ads -> Calculate eCPM (Bid * CTR probability) -> Select Winner.
*   **Pacing:** Don't show budget in 1 hour. Smooth delivery.

### Question 390: Build a real-time sports score platform.

**Answer:**
*   **Source:** Sport Data Provider (Opta/Sportradar) Push Feed.
*   **Ingest:** Webhook -> Redis Pub/Sub.
*   **Push:** WebSocket Server subscribes to Redis. Broadcasts to 1M connected clients.
*   **Optimization:** Delta updates ("Score changed 1-0", not full object).

---

## 🔸 DevOps & Deployment Systems (Questions 391-400)

### Question 391: Design an internal CI/CD system.

**Answer:**
*   **Pipeline as Code:** YAML (`.gitlab-ci.yml`).
*   **Runner:** Agent executing jobs (Docker container).
*   **Stages:** Build (Compile) -> Test (Unit) -> Package (Docker Build) -> Deploy (Helm Upgrade).
*   **Artifacts:** Store Jars/Docker Images in Artifactory/Registry.

### Question 392: How would you build infrastructure provisioning using Terraform?

**Answer:**
*   **State:** Store `.tfstate` in S3 with Locking (DynamoDB).
*   **Modules:** Reusable components (`vpc`, `rds`, `k8s`).
*   **Workflow:** `terraform plan` (Review changes) -> `terraform apply`.
*   **Drift:** Periodic check if Real Infra diverges from Code.

### Question 393: Design a centralized logging solution for 500 microservices.

**Answer:**
(See Q125 & Q310). Focus on Scale.
*   **Tiering:** Kafka handles burst.
*   **Standardization:** Enforce JSON schema libraries across teams.

### Question 394: How to auto-scale based on CPU and memory metrics?

**Answer:**
*   **HPA (K8s):**
    *   Metrics Server scrapes cAdvisor.
    *   HPA Controller checks `Current / Target`.
    *   `DesiredReplicas = CurrentReplicas * (CurrentMetric / TargetMetric)`.
*   **Lag:** Takes 1-2 mins.

### Question 395: How to secure deployments using GitOps?

**Answer:**
(ArgoCD / Flux).
*   **Pull Model:** Cluster pulls config from Git. CI Pipeline does NOT have `kubeconfig` access (Security+).
*   **Sync:** Operator ensures Cluster State == Git State.
*   **Rollback:** `git revert`.

### Question 396: How do you perform rolling upgrades for Kubernetes services?

**Answer:**
*   **Strategy:** `maxUnavailable: 25%`, `maxSurge: 25%`.
*   **Process:**
    1.  Create new Pod. Wait for Ready.
    2.  Kill old Pod.
    3.  Repeat.
*   **Result:** Capacity never drops below 75%. Service remains available.

### Question 397: How would you design a secrets rotation system?

**Answer:**
*   **Automation:** Lambda function triggered by CloudWatch Event (every 30 days).
*   **Action:**
    1.  Generate New Key. Store in Secrets Manager.
    2.  Update Database User password.
    3.  Restart App (or App re-fetches secret).
    4.  Deactivate Old Key.

### Question 398: How to isolate noisy containers in a shared cluster?

**Answer:**
*   **Resources:** Requests (Guaranteed) and Limits (Cap).
*   **QoS Class:**
    *   **Guaranteed:** Request == Limit. (High Priority).
    *   **Burstable:** Request < Limit. (Throttled first).
    *   **BestEffort:** No limits. (Killed first on OOM).

### Question 399: Design a system for monitoring container resource limits.

**Answer:**
*   **Metric:** `container_cpu_cfs_throttled_seconds_total` (Prometheus).
*   **Alert:** If throttling > 5% of time, increase Limit.
*   **OOM:** Monitor `container_last_seen` or `OOMKilled` exit code.

### Question 400: Build a dashboard to track deployment status across environments.

**Answer:**
*   **Data:** Query K8s API (`images` running in `Prod` namespace) + Git Commit Hash.
*   **Visual:** Matrix. Rows=Services, Cols=Envs (Dev, Stage, Prod). Cells=Version.
*   **Diff:** Highlight if Prod Version < Stage Version.
