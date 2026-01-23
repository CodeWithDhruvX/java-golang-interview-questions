## 🟢 Security, Edge & Platform (Questions 751-800)

### Question 751: What’s a prototype pollution vulnerability, and how does Node help mitigate it?

**Answer:**
Attacker feeds JSON `{"__proto__": {"admin": true}}`. Merging this into an object modifies the global `Object.prototype`.
**Mitigation:**
1.  Use `Object.create(null)` for maps.
2.  Use `--disable-proto=delete` flag (Node 20+) to disable `__proto__`.
3.  Freeze prototype `Object.freeze(Object.prototype)`.

---

### Question 752: How do you defend against path traversal when serving files?

**Answer:**
Attack: `GET /../../etc/passwd`.
**Defense:**
1.  `path.normalize(input)`.
2.  Check if `resolvedPath.startsWith(baseDir)`.
3.  Use secure libraries like `send` or `express.static` (they handle this).

---

### Question 753: How do you securely evaluate dynamic code (e.g., user-generated math)?

**Answer:**
**Never** use `eval`.
Use a parser library (mathjs) to evaluate expressions safely.
If logic is complex, run in **Isolated-VM** or separate container.

---

### Question 754: What tools detect dependency-level vulnerabilities in Node.js?

**Answer:**
1.  `npm audit`.
2.  **Snyk**.
3.  **Dependabot** (GitHub).

---

### Question 755: What are security headers, and how do you implement them in Node?

**Answer:**
HSTS, X-Frame-Options, CSP, X-Content-Type-Options.
Use `helmet` middleware in Express.

---

### Question 756: How do you cache SSR responses at the CDN layer from Node?

**Answer:**
Set `Cache-Control` header.
`res.set('Cache-Control', 'public, max-age=300, s-maxage=3600');`
`s-maxage` tells CDN to cache for 1 hour.

---

### Question 757: How do you support stale-while-revalidate in Node?

**Answer:**
Header: `Cache-Control: public, s-maxage=60, stale-while-revalidate=600`.
CDN serves stale content while fetching new content from Node in background.

---

### Question 758: How do you differentiate bot vs. user traffic in edge Node apps?

**Answer:**
Inspect `User-Agent` header.
Use Middleware to check against known bot list (Googlebot, Bingbot).
Or use headers provided by CDN (Cloudflare `CF-Worker`).

---

### Question 759: How do you run a Node function across multiple PoPs?

**Answer:**
Deploy to Edge Platform (Lambda@Edge, Fly.io, Cloudflare).
They automatically replicate/route code execution to the Point of Presence closest to user.

---

### Question 760: How do you implement rate limiting at the CDN edge with Node fallback?

**Answer:**
CDN (Cloudflare) usually handles DDOS.
Node falls back to application-logic limits (e.g., "5 posts per day").
Both layers are needed.

---

### Question 761: How do you build a Node SDK for a third-party API?

**Answer:**
1.  Wrap `fetch` calls in a Class.
2.  Provide Typed methods (`getUser(id)`).
3.  Handle Auth internally.
4.  Normalize Errors.

---

### Question 762: How do you version APIs exposed via an SDK?

**Answer:**
Allow passing version in config.
`new SDK({ version: 'v2' })`.
Or publish separate packages `@sdk/v1`, `@sdk/v2`.

---

### Question 763: How do you support plugin hooks in a Node SDK?

**Answer:**
Allow user to pass `onRequest` or `onResponse` callbacks.
Use `tapable` or `EventEmitter`.

---

### Question 764: How do you enforce request schema validation in SDK usage?

**Answer:**
Use TS Types for compile time.
Use `zod` or `joi` at runtime in the SDK method to validate arguments before making the network call.

---

### Question 765: How do you handle breaking API changes in SDKs?

**Answer:**
SemVer.
If API v2 breaks response shape, release SDK v2.0.0.
Maintain v1.x branch for old API.

---

### Question 766: How do you retry failed API calls with exponential backoff?

**Answer:**
Wait time = `base * (2 ^ attempt)`.
Loop + `setTimeout`.
Avoid "Thundering Herd" by adding Jitter (Randomness).

---

### Question 767: How do you implement a circuit breaker in Node?

**Answer:**
(See Q550).
State: Closed (Normal), Open (Fail).
If Errors > Threshold, Switch to Open (Fail immediately).
After timeout, Half-Open (Try one request). If success -> Closed.

---

### Question 768: What is a bulkhead pattern, and how does it apply to Node?

**Answer:**
Isolating resources.
If "Image Processing" is heavy, limit it to X concurrent requests.
This ensures "User Login" (fast) isn't starved by Image Processing requests clogging the event loop.

---

### Question 769: How do you isolate failed dependencies in a microservice?

**Answer:**
Wrap calls in Try/Catch.
Return fallback data (Default Avatar if Avatar Service fails).
Don't crash the orchestrator.

---

### Question 770: How do you track downstream dependency health?

**Answer:**
Health Check endpoint probes downstream services.
Or passive monitoring (track error rates of outgoing calls).

---

### Question 771: How do you mark a package as private vs. public?

**Answer:**
`package.json`: `"private": true`. (Prevents excessive publishing).
To publish public: Remove it or set `"publishConfig": { "access": "public" }`.

---

### Question 772: How do you publish beta tags like `@next` or `@alpha`?

**Answer:**
`npm publish --tag next`.
Users install via `npm install pkg@next`.

---

### Question 773: How do you deprecate a published npm version?

**Answer:**
`npm deprecate my-package@"< 1.0.2" "Critical bug found, update to 1.0.2"`.
Prints warning on install.

---

### Question 774: What are the security risks of installing unscoped packages?

**Answer:**
Typosquatting (`react` vs `raect`).
Dependency Confusion.
Scoped packages (`@angular/core`) are safer as they are owned by verified orgs.

---

### Question 775: How do you sign npm packages for verification?

**Answer:**
npm now supports **Sigstore**.
`npm publish --provenance`.
Links package to the GitHub Action run that built it, providing verifiable trail.

---

### Question 776: How do you pass complex data between threads in Node?

**Answer:**
`worker.postMessage(obj)`.
Uses Structured Clone Algorithm (Copies data).
For Zero-copy (huge data), use `SharedArrayBuffer` or `Transferable` (ArrayBuffer).

---

### Question 777: How do you measure CPU load in a multithreaded Node app?

**Answer:**
Main process usage doesn't show worker usage.
In Worker: call `process.cpuUsage()`. Send stats to Main. Main aggregates.

---

### Question 778: How do you handle thread crashes in a safe way?

**Answer:**
Worker emits `error` and `exit` event.
Main logic:
1.  Log error.
2.  Restart worker (Spawning pool).
3.  Dead-letter queue the task that crashed it.

---

### Question 779: When should you use `SharedArrayBuffer` in Node?

**Answer:**
For high-performance shared state (e.g., Matrix for ML, Counters).
(See Q452).

---

### Question 780: How do you terminate a runaway thread?

**Answer:**
`worker.terminate()`.
Immediately stops the thread. Data processing mid-way is lost.

---

### Question 781: How do you direct requests to region-specific backends in Node?

**Answer:**
DNS-level (GeoDNS).
Or API Gateway checks User IP -> Resolves Region -> Proxies to `us-east.api.com`.

---

### Question 782: How do you manage latency-aware routing in a Node API?

**Answer:**
Usually Infrastructure job (AWS Latency Routing).
In Node code: connect to nearest DB replica (Read) based on config.

---

### Question 783: How do you sync state across regions in a distributed Node app?

**Answer:**
Global Database (DynamoDB Global Tables / CockroachDB).
Or Async Replication (Event Bus).
Do not try to sync in-memory state across regions immediately (CAP theorem).

---

### Question 784: What are the challenges of sticky sessions in multi-region Node deployments?

**Answer:**
User hops region (travel/failover).
Session data is in Region A's Redis. Region B doesn't have it.
**Fix:** Stateless JWT or Global Redis (slow).

---

### Question 785: How do you invalidate region-specific cache entries?

**Answer:**
Broadcast "Purge" event globally.
Message Bus (Kafka) propagates to all regions.
Each Region's Node Cluster subscribes and clears local cache.

---

### Question 786: How do you implement `fetch()` with full spec compliance in Node?

**Answer:**
Use standard `fetch`. (Node 18+).
It is compliant (mostly).
Differences: File uploads might use `fs` streams instead of Blobs in older polyfills.

---

### Question 787: What is the AbortController pattern and how is it used?

**Answer:**
(See Q195).
Pass `signal` to async operation. Call `abort()`.
Listener `signal.addEventListener('abort')` triggers cleanup.

---

### Question 788: How do you polyfill newer web APIs in older Node versions?

**Answer:**
Use libraries like `node-fetch`, `abort-controller`, `web-streams-polyfill`.
Import and attach to global if needed.

---

### Question 789: How do you use Headers and FormData objects in Node streams?

**Answer:**
Native in Node 18.
`const fd = new FormData(); fd.append('file', blob);`
`fetch(url, { body: fd })` automatically sets multipart headers.

---

### Question 790: How do you simulate a service worker in Node.js for SSR?

**Answer:**
Difficult. Service Worker intercepts network.
In Node, you can wrap `fetch` to check a Cache before going to network.
Libraries: `msw` (Mock Service Worker) works in Node to intercept requests.

---

### Question 791: How do you implement real-time event tracking in Node?

**Answer:**
Streaming pipeline.
API Endpoint -> Kafka Producer -> Data Lake.
Fire-and-forget (don't wait for ACK if speed > reliability).

---

### Question 792: How do you debounce high-volume client events server-side?

**Answer:**
(See Q236).
Aggregator pattern.
Store events in Redis List.
Worker processes list every 5 seconds.

---

### Question 793: How do you batch telemetry for export to a data warehouse?

**Answer:**
Buffer in memory.
`if (buffer.length > 1000 || time > 1min) flush()`.
Bulk Insert to DB.

---

### Question 794: How do you scrub PII before logging user behavior?

**Answer:**
Middleware or Logger Transform.
Recursively traverse object.
Mask keys: `password`, `email`, `ssn`.
`email: "r***@gmail.com"`.

---

### Question 795: How do you A/B test features at the API level?

**Answer:**
Assign user to Bucket (Hash-based or Random).
`if (bucket == 'B') return newFeature()`.
Log "Exposure Event" for analytics.

---

### Question 796: What is the Node.js Release Working Group?

**Answer:**
Team responsible for managing release schedules, LTS policies, and producing the binary releases.

---

### Question 797: How do you propose a new feature to core Node?

**Answer:**
Open an Issue (Feature Request).
Draft a PR.
Or create an RFC in `nodejs/prevention`.

---

### Question 798: How do you write a good README for a Node library?

**Answer:**
Badges (CI/NPM).
Installation.
Usage (Clear Code Snippets).
API Reference.
License.

---

### Question 799: What are stability indices in Node.js core modules?

**Answer:**
Docs show Stability:
0 - Deprecated.
1 - Experimental.
2 - Stable. (Safe to use).

---

### Question 800: How do you contribute test coverage to the Node.js core repo?

**Answer:**
Clone `nodejs/node`.
Build.
Run tests (`make test`).
Find gaps (Code Coverage report).
Add test case. Submit PR.
