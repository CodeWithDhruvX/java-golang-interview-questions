## 🔸 Multimedia & Real-Time Media Systems (Questions 701-710)

### Question 701: Design a system like YouTube with video upload, transcoding, and streaming.

**Answer:**
*   **Upload:** Resumeable Upload (TUS Protocol) to S3.
*   **Process:** Lambda triggers AWS MediaConvert. Transcodes to HLS (360p, 720p, 1080p).
*   **Delivery:** CloudFront CDN.
*   **Metadata:** DynamoDB stores `VideoID`, `Duration`, `HLS_URL`.

### Question 702: How would you build a real-time video conferencing app backend?

**Answer:**
*   **Protocol:** WebRTC (Peer-to-Peer) for 1:1.
*   **Group Call:** SFU (Selective Forwarding Unit) like Jitsi/Mediasoup. Server receives 1 stream from User A, forwards to B, C, D.
*   **Signaling:** WebSocket to exchange SDP (Session Description Protocol) and ICE Candidates.

### Question 703: Design a podcast hosting and distribution platform.

**Answer:**
*   **Feed:** Generate RSS XML (`feed.xml`).
*   **Hosting:** S3 for MP3 files.
*   **Analytics:**
    *   **Method:** Server-Side Log Analysis (CloudFront Logs).
    *   **Range Request:** Detect "Download" vs "Stream" (Partial content 206).

### Question 704: Build a low-latency live streaming infrastructure.

**Answer:**
*   **Protocols:** RTMP (Ingest) -> WebRTC / LL-HLS (Low Latency HLS) for Playback.
*   **Latency:** Standard HLS (~10s). LL-HLS (~2s). WebRTC (< 500ms).
*   **Edge:** Transcode at the Edge (Cloudflare Workers) to minimize hop.

### Question 705: How do you implement video thumbnail generation at scale?

**Answer:**
*   **Trigger:** Upload Complete.
*   **Job:** FFMPEG extracts frame 0, 10s, 50% mark.
*   **Sprite:** Stitch frames into a single sprite sheet image (for hover preview).

### Question 706: Design an audio transcription service with multi-language support.

**Answer:**
*   **Queue:** Upload -> SQS -> Worker.
*   **Model:** Whisper (OpenAI).
*   **Chunks:** Split 1 hour audio into 30s chunks. Process in parallel. Stitch text.
*   **Timestamps:** Output VTT/SRT format `00:01 --> 00:05 Hello`.

### Question 707: Build a backend for short video creation and sharing (like TikTok).

**Answer:**
*   **Pre-fetch:** Feed algorithm predicts next videos. Client downloads top 5.
*   **Feed:** Real-time ranking using Flink.
*   **Upload:** Client-side compression / filter application (GLSL).

### Question 708: How to handle content moderation for user-uploaded media?

**Answer:**
*   **Automated:**
    *   Hash Matching (PDQ Hash) against known bad images.
    *   ML (safety-detectors) for Nudity/Violence.
*   **Manual:** Queue flagged items for human review.

### Question 709: Design a distributed video encoding pipeline.

**Answer:**
*   **Split:** Split input.mp4 into 5-minute segments.
*   **Distribute:** Workers encode segments in parallel.
*   **Merge:** `ffmpeg -f concat` to stitch encoded segments.
*   **Speed:** Reduces 1 hour encoding to minutes.

### Question 710: Build a collaborative video annotation and feedback system.

**Answer:**
*   **Model:** `Annotation` (VideoID, Timestamp, User, Text, Rect {x,y,w,h}).
*   **Player:** Pauses at timestamp. Draws SVG layer over video.
*   **Sync:** WebSocket pushes new annotation to active viewers.

---

## 🔸 Geo-Distributed & Multi-Region Architectures (Questions 711-720)

### Question 711: Design a global ride-sharing platform with geo-aware matchmaking.

**Answer:**
*   **Sharding:** S2 Geometry (Google). Cell ID at Level 12 (~2km).
*   **Match:** Query `Drivers` index for `S2_CellID` neighbors of `Rider`.
*   **State:** Redis Cluster with Geo module.

### Question 712: How to replicate databases across continents with low latency?

**Answer:**
(See Q133/Q422).
*   **X-Region Read Replica:** Async replication. Local Reads.
*   **Write:** Single Master (US) or Multi-Master (Aurora Global / DynamoDB Global Tables).

### Question 713: Build a global DNS management platform.

**Answer:**
*   **Anycast:** 1 IP announced from 50 POPs.
*   **Config:** Distributed Key-Value Store propagates records to edge nodes.
*   **Logic:** Edge node runs BIND/CoreDNS. Answers queries from local memory.

### Question 714: Design a latency-aware content delivery network (CDN).

**Answer:**
*   **Request:** User -> Edge Node.
*   **Cache:**
    1.  Check RAM (Hot).
    2.  Check SSD (Warm).
    3.  Check Origin (Miss).
*   **Routing:** BGP routing ensures User hits closest Edge.

### Question 715: Build a user session store that’s accessible worldwide.

**Answer:**
*   **Global Table:** DynamoDB Global Table.
*   **Strategy:** Write `Session` to local region. AWS replicates to others.
*   **Conflict:** Last Write Wins. (Login in NY overwrites Login in London).

### Question 716: Design a disaster-tolerant data backup system across regions.

**Answer:**
*   **RPO:** Recovery Point (Data loss). **RTO:** Recovery Time (Downtime).
*   **Strategy:** Cross-Region Replication (CRR) on S3 buckets.
*   **Compliance:** Only replicate to "Allowed" regions (GDPR).

### Question 717: How to route traffic based on geographic failover policies?

**Answer:**
*   **DNS:** Route53.
*   **Health:** Associate Health Check with US-East Record.
*   **Failover:** If Health Check fails, remove US-East IP, return EU-West IP. TTL must be low (60s).

### Question 718: Design a multi-region checkout flow for an e-commerce site.

**Answer:**
*   **Inventory:** Global Inventory DB (Pinned to US).
*   **Reservation:**
    *   Read Local Cache.
    *   Hard Reserve against Global DB (Cross-region call required, latency incurred for correctness).
*   **Optimization:** Regional Allowances (Allocated 100 iPhone to EU warehouse).

### Question 719: How to resolve conflicts in distributed systems with time skew?

**Answer:**
*   **Google TrueTime:** Spanner uses Atomic Clocks + GPS to bound error.
*   **HLC (Hybrid Logical Clocks):** Combines Physical Time + Logical Counter.
*   **Logic:** `HLC.Now() = Max(Physical, Parent.HLC)`.

### Question 720: Build a privacy-aware geo-location logging system.

**Answer:**
*   **Fuzzing:** Truncate Lat/Lon to 2 decimal places (1km accuracy).
*   **Cloaking:** Randomly shift point within circle.
*   **Storage:** `Geohash`.

---

## 🔸 Access Control & Permissions (Questions 721-730)

### Question 721: Design a role-based access control system (RBAC).

**Answer:**
(See Q295).

### Question 722: How would you implement attribute-based access control (ABAC)?

**Answer:**
(See Q327 for OPA).
*   **Attributes:** Subject (User), Object (File), Environment (Time/Location).
*   **Rule:** `Grant IF User.Level >= File.Class AND Time in OfficeHours`.

### Question 723: Build a secure permission audit trail system.

**Answer:**
*   **Event:** `PermissionGranted(Admin, TargetUser, Role, Time)`.
*   **Chain:** Hash Chaining (`Hash(PrevEvent + CurrentEvent)`).
*   **Verify:** Recompute hashes to detect deleted logs.

### Question 724: How to manage temporary access to sensitive data?

**Answer:**
*   **TTL:** Grant access with Expiry.
*   **JIT (Just-In-Time):**
    *   User requests access.
    *   Approver approves.
    *   System adds User to Group.
    *   Scheduled Job removes User from Group after 2 hours.

### Question 725: Design a permission hierarchy system for nested organizations.

**Answer:**
*   **Tree:** `RootOrg -> SubOrg -> Team`.
*   **Propagation:** ACL on Root check implies ACL on SubOrg?
*   **Graph:** Use Graph DB (Neo4j) to traverse `User -> MemberOf -> Team -> ChildOf -> SubOrg`.

### Question 726: How do you revoke user access instantly across services?

**Answer:**
*   **Stateless JWT:** Cannot revoke instantly (Wait for expiry).
*   **Blacklist:** Push `JTI` (Token ID) to Redis "Revocation List" on all Gateways.
*   **Version:** User has `token_version` in DB. JWT has `v: 1`. Increment DB version on logout. Gateway checks DB (Cache) vs JWT.

### Question 727: Build a user delegation system (acting on behalf of another user).

**Answer:**
*   **Token:** Exchange Admin Token for "Impersonation Token".
*   **Scope:** Limit Impersonation Token scope (Read-Only).
*   **Audit:** Key requirement. Log `RealUser` AND `ImpersonatedUser`.

### Question 728: How to handle access control for shared resources?

**Answer:**
(Google Docs sharing).
*   **ACL:** List of `(User/Group, Permission)`.
*   **Link Sharing:** `Token -> Permission`.
*   **Check:** `if NotInACL(User) AND NoLinkToken() -> Deny`.

### Question 729: Implement fine-grained permissions at object-level scope.

**Answer:**
*   **ReBAC (Relationship Based):** Zanzibar (Google).
*   **Tuples:** `(User:Alice, viewer, Doc:123)`.
*   **Check:** `check(Alice, viewer, Doc:123)`.
*   **Scale:** Optimized for billions of objects.

### Question 730: Design a service for reviewing and approving access requests.

**Answer:**
(Identity Governance).
*   **Workflow:** `Request -> Manager Approval -> Owner Approval -> Provision`.
*   **Notification:** Slack Bot / Email.
*   **Escalation:** If Manager doesn't approve in 24h -> Skip or Remind.

---

## 🔸 UX & Frontend-Driven Backend Design (Questions 731-740)

### Question 731: Design a backend for a drag-and-drop form builder.

**Answer:**
*   **Schema:** JSON. `fields: [{ type: "text", label: "Name", required: true }]`.
*   **Version:** `FormID`, `Version`.
*   **Render:** FE loops over JSON array to render components.
*   **Submission:** Validate payload against JSON schema.

### Question 732: Build a recommendation engine for a product configurator.

**Answer:**
(Car Builder).
*   **Constraints:** `Engine:V8` requires `Chassis:Sport`.
*   **CSP (Constraint Satisfaction Problem):** Solver engine.
*   **API:** `GET /options?selected=[V8]`. Returns available chassis, grays out others.

### Question 733: How to support undo/redo functionality for multi-user apps?

**Answer:**
(See Q682). Command Pattern + OT/CRDT.

### Question 734: Design a system to store user dashboards and widgets.

**Answer:**
(See Q582).

### Question 735: Build a flexible notification preference backend.

**Answer:**
(See Q130/Q445).

### Question 736: Design a theme and layout manager for SaaS products.

**Answer:**
*   **CSS Variables:** Backend stores `{"primary": "#ff0000", "font": "Roboto"}`.
*   **Injection:** Frontend fetches config, applies to `:root` style.
*   **Compiling:** SASS compiler on server to generate `client-123.css` (Performance).

### Question 737: How to implement live cursor tracking (like Figma)?

**Answer:**
*   **Transport:** WebSocket / UDP (Unreliable ok).
*   **Optimization:** Throttling (Send 10 times/sec). Dead Reckoning (Client predicts movement between points).
*   **Ephemeral:** Don't save to DB. Pub/Sub only.

### Question 738: Design an interface versioning system for backward compatibility.

**Answer:**
*   **Problem:** User loads Old UI (Cached) -> Calls New API.
*   **Feature Flags:** Gating UI components.
*   **Endpoint:** API supports both formats or transforms response based on `Accept-Version` header.

### Question 739: Build a progressive onboarding backend system.

**Answer:**
*   **State:** `UserSteps: { "tutorial": true, "first_post": false }`.
*   **Logic:** `NextStep = FindFirstFalse(OrderedSteps)`.
*   **API:** `POST /step/complete { name: "first_post" }`.

### Question 740: Design a feature-tour trigger system based on user behavior.

**Answer:**
*   **Events:** "User hovered 'Export' button".
*   **Rule:** `If Hover > 3s AND TourNotSeen('export') -> Trigger Tour`.
*   **State:** Store `SeenTours` to prevent annoyance.

---

## 🔸 Enterprise-Scale Operations (Questions 741-750)

### Question 741: Design a unified audit log system across services.

**Answer:**
*   **Sidecar:** Fluentd sidecar in every Pod.
*   **Format:** Structured Common Event Format (CEF).
*   **Pipeline:** Sidecar -> Kafka -> S3 (Archive) + Elastic (Search).

### Question 742: Build a unified identity provider for SSO integration.

**Answer:**
(Build your own Auth0).
*   **Federation:** Connect upstream to Google/AD/Okta.
*   **Protocol:** Expose OIDC endpoints (`/authorize`, `/token`, `/userinfo`, `/jwks.json`).
*   **Session:** Centralized session management.

### Question 743: Design a dashboard for tracking KPIs in real time.

**Answer:**
(See Q492). Redis Timeseries or Druid.

### Question 744: How to support multi-department billing and reporting?

**Answer:**
*   **Tagging:** Every resource (EC2/DB) must have `CostCenter` tag.
*   **Reapers:** Scripts kill resources without tags.
*   **Report:** Aggregation by `CostCenter`.

### Question 745: Build a security incident reporting system.

**Answer:**
*   **Intake:** Web Form, Email, API.
*   **Triage:** Assign Severity (P0-P4).
*   **Workflow:** Integration with Jira/ServiceNow.
*   **SLA:** Timer based on Severity. P0 = 15 mins response.

### Question 746: Design a deployment pipeline with approval workflows.

**Answer:**
*   **Gate:** CI passes -> Wait for Approval.
*   **Auth:** Only `Role:ReleaseManager` can POST `/deploy/approve`.
*   **Audit:** Record who approved.

### Question 747: Build a productivity insights system using activity data.

**Answer:**
*   **Sources:** Git Commits, Jira Tickets, Slack Activity.
*   **Privacy:** Aggregation Only (Team Level).
*   **Metric:** Cycle Time (First Commit to Deploy). Deploy Frequency.

### Question 748: How to manage configurations across thousands of tenants?

**Answer:**
*   **Hierarchy:** `GlobalConfig` -> `RegionConfig` -> `TenantConfig` -> `UserConfig`.
*   **Overlay:** Merge strategy. Specific overrides General.
*   **Tool:** LaunchDarkly feature flags context.

### Question 749: Design a compliance reporting engine with export capabilities.

**Answer:**
*   **Report:** Snapshot of "Who had access to what" at Date D.
*   **Generation:** Temporal Query on Audit Logs.
*   **Format:** PDF/CSV. Signed.

### Question 750: Build a system for managing employee hardware inventory.

**Answer:**
*   **Asset:** `Laptop`, `Serial`, `AssignedTo`.
*   **Lifecycle:** `Procured -> Assigned -> Repair -> Retired`.
*   **Scan:** MDM (Jamf) reports actual serials.
*   **Reconcile:** Diff MDM report vs Database. Alert on missing laptops.
