## 🟢 Web Standards & Edge AI (Questions 901-950)

### Question 901: How do you use `fetch()` streaming with ReadableStreams in Node?

**Answer:**
`response.body` is a web `ReadableStream`.
```javascript
const res = await fetch(url);
for await (const chunk of res.body) {
  // process chunk (Uint8Array)
}
```

---

### Question 902: How does Node.js support `ReadableStream` from Web Streams Standard?

**Answer:**
Global `ReadableStream` class.
Interoperable with `buffer`, `fetch`, etc.
Can be converted to Node stream via `Readable.fromWeb()`.

---

### Question 903: What is the `Blob` object in Node.js and when is it used?

**Answer:**
Immutable raw data.
Used in `fetch` API, creating Object URLs (in browser logic ported to Node), or constructing `FormData`.

---

### Question 904: How do you handle `FormData` server-side in Node using Web standard APIs?

**Answer:**
Node 18 `Request` object has `.formData()` method.
```javascript
// if using a standard-compliant framework
const data = await req.formData();
const file = data.get('file'); // Blob/File
```

---

### Question 905: How does Node implement `URLPattern` for matching routes?

**Answer:**
Experimental global `URLPattern`.
`const pattern = new URLPattern({ pathname: '/users/:id' });`
`pattern.exec(url)`.
Standard way to parse routes without regex libs.

---

### Question 906: What are Node.js isolates in edge computing environments?

**Answer:**
V8 Isolates. Light containers.
Own Heap/GC.
Fast startup.
Shared underlying OS thread.

---

### Question 907: How do you share memory between multiple isolates safely?

**Answer:**
Generally restricted.
Use `SharedArrayBuffer` if allowed by runtime.
Or external KV store (Redis on Edge).

---

### Question 908: What are multi-tenancy constraints when deploying Node on edge runtimes?

**Answer:**
Strict CPU limits (50ms).
Memory limits (128MB).
No long running processes.
Code size limits (1MB).

---

### Question 909: How do you dynamically load WebAssembly at runtime in edge environments?

**Answer:**
Standard `WebAssembly.instantiateStreaming(fetch('mod.wasm'))`.
Efficient compilation during download.

---

### Question 910: How do you detect and handle runtime limitations (e.g. memory/performance) in edge Node?

**Answer:**
`try/catch` around heavy ops.
Check memory usage.
If close to limit, return 503 or Partial content.

---

### Question 911: How do you share data between a parent and child process securely?

**Answer:**
`IPC` channel (`send`/`on('message')`).
Pass only **serialized** data (JSON).
Validate schema on receive.

---

### Question 912: What are the pitfalls of serializing large objects via IPC?

**Answer:**
JSON.stringify/parse is sync and slow. Blocks both processes.
**Fix:** Use shared memory or pipe a stream.

---

### Question 913: How do you revoke shared buffers to prevent memory leaks?

**Answer:**
If using `Transferable` (ArrayBuffer), transferring it to Worker "neutrals" (empties) it in the Main thread. Ownership moves. No GC leak.

---

### Question 914: How do you implement capability-based security between Node worker threads?

**Answer:**
Only pass specific `MessagePort` to the worker.
Worker can only talk to that port, not the global scope.

---

### Question 915: How can circular references affect structured clone in IPC?

**Answer:**
Node's IPC (serialization) usually handles circular refs (Structured Clone Algorithm).
`JSON.stringify` fails, but `postMessage` succeeds.

---

### Question 916: How would you build a WebSocket-based geolocation tracker in Node?

**Answer:**
Client sends GPS `{lat, lng}` via WS.
Server buffers last locations in Redis `GEOADD`.
Server broadcasts to viewers.

---

### Question 917: How do you perform geospatial queries at scale from Node?

**Answer:**
Redis: `GEORADIUS key longitude latitude 10 km`.
MongoDB: `$near` query.
PostGIS: SQL geography types.

---

### Question 918: How do you handle latency-sensitive GPS data streams in Node apps?

**Answer:**
UDP instead of TCP (Node `dgram`).
Drop old packets (real-time matters more than complete history).

---

### Question 919: How do you manage geofencing logic in a distributed Node system?

**Answer:**
Shard users by Region (Geohash).
Check entering/exiting polygon logic on the node responsible for that geohash.

---

### Question 920: How do you simulate geo‑failover in real‑time Node clusters?

**Answer:**
Manually partition network between regions.
Verify clients reconnect to secondary region.

---

### Question 921: How do you stream synthetized audio in real time using Node?

**Answer:**
(Text-to-Speech).
Pipe TTS Engine output (PCM Stream) -> `FFmpeg` (to ensure MP3/Opus) -> HTTP Response.

---

### Question 922: How do you support live video transcoding pipelines in Node backends?

**Answer:**
Spawn `ffmpeg`.
Input: RTMP stream.
Output: HLS (.m3u8 + .ts segments) written to disk/S3.
Node serves the playlist file.

---

### Question 923: How do you merge multiple live media streams into one feed?

**Answer:**
Video Mixing.
Complex FFMPEG filter `[0:v][1:v]overlay`.
Heavy CPU.
Better to do client-side via WebRTC.

---

### Question 924: How do you control media quality adaptively over unstable networks?

**Answer:**
HLS / DASH.
Generate 3 qualities (Low, Med, High).
Client switches based on bandwidth.

---

### Question 925: What is the role of `MediaStreamTrackProcessor` in Node?

**Answer:**
Web API (Experimental). Allows manipulating raw video/audio frames in JS streams.
Mainly a browser API, but relevant if using headless browser or compatible polyfills.

---

### Question 926: How do you leverage hardware-backed key stores from Node?

**Answer:**
PKCS#11 interface.
Use `graphene-pk11` library.
Talk to HSM (Hardware Security Module) to sign data without exposing private key.

---

### Question 927: How do you interface with TPM or secure enclave modules in Node?

**Answer:**
TPM 2.0 tools (`tpm2-tss`).
Native addon wrappers.

---

### Question 928: How do you accelerate cryptographic operations with native bindings?

**Answer:**
Node `crypto` already uses OpenSSL (Native). Use it.
Don't use JS-only crypto libs (`crypto-js`) for heavy lifting.

---

### Question 929: How do you handle hardware‑backed signature verification in Node?

**Answer:**
Receive public key from hardware (certificate).
Verify normally in software (OpenSSL).

---

### Question 930: How do you offload crypto to GPU or hardware modules securely?

**Answer:**
CUDA bindings.
Or offloading to a sidecar proxy that handles SSL termination (Nginx/Envoy).

---

### Question 931: How do you integrate TypeScript runtime type validation?

**Answer:**
Libraries like `runtypes` or `zod`.
Define schema once. Infers TS type. Validates JS object at runtime.

---

### Question 932: How do you auto-generate validation middleware from an OpenAPI spec?

**Answer:**
`express-openapi-validator`.
Reads `openapi.yaml`.
Automatically validates every incoming request body against the spec before hitting controller.

---

### Question 933: How do you enforce type-safe GraphQL schemas in Node runtime?

**Answer:**
`graphql-codegen`.
Generates TS types for Resolvers.
Ensures Resolver return value matches Schema type.

---

### Question 934: How do you support end-to-end type safety between front-end and Node backends?

**Answer:**
**tRPC**.
Shared Router Type definition.
Frontend imports Type (not code).
`client.getUser.query()` results in typed response effortlessly.

---

### Question 935: How do you validate WebSocket message payloads at runtime with type generators?

**Answer:**
Attach a schema ID to message.
`{ type: 'CHAT', payload: ... }`.
Switch(type), select Zod schema, parse payload.

---

### Question 936: How do you manage sub‑millisecond latency in WebSocket data feeds via Node?

**Answer:**
1.  Kernel tuning (sysctl).
2.  Disable Nagle (`setNoDelay`).
3.  Use buffer pools (no GC).
4.  `uWebSockets.js` (C++ based WS server) - faster than native Node WS.

---

### Question 937: How do you aggregate real-time market data streams in Node?

**Answer:**
Sliding Window.
Keep data in ring buffer.
Calculate Average/Volume on emit.
Publish summary.

---

### Question 938: How do you avoid GC pauses in latency‑sensitive trading systems?

**Answer:**
**Object Pooling**. Reuse objects instead of `new Object()`.
Avoid closures.
Pre-allocate buffers.

---

### Question 939: What techniques reduce latency in Node serialization/deserialization?

**Answer:**
Schema-based serialization (fast-json-stringify).
Avoids scanning object keys. Generates string builder code.
Or use binary (SBE - Simple Binary Encoding).

---

### Question 940: How do you recover from failed heartbeats in live trading system processes?

**Answer:**
Watchdog process.
If logic thread hangs, Kill -9 and restart immediately (fail fast).
Failover to hot standby.

---

### Question 941: How do you manage cross‑cloud cluster communication between Node services?

**Answer:**
VPN / VPC Peering.
Service Mesh (consul connect).
Or Public Internet with mTLS.

---

### Question 942: How do you route data based on cloud‑region affinity?

**Answer:**
DNS Geolocation.
And App Logic: "User is in EU, write to EU-DB".

---

### Question 943: How do you ensure data consistency across multi-cloud Node databases?

**Answer:**
Distributed SQL (CockroachDB/Yugabyte).
Or eventual consistency (Replication).

---

### Question 944: How do you handle failover for Node APIs between cloud providers?

**Answer:**
Global Load Balancer (Cloudflare/Akamai).
Health check origin A (AWS). If down, route to origin B (GCP).

---

### Question 945: How do you unify logging and tracing across multi‑cloud Node deployments?

**Answer:**
Ship logs to a 3rd party SaaS (Datadog/Logz.io).
Don't store locally.

---

### Question 946: How do you aggressively pre-render API responses at build time?

**Answer:**
**SSG (Static Site Generation)** paradigm applied to API.
Generate JSON files (`products.json`). Upload to CDN.
API endpoint just redirects to CDN URL.

---

### Question 947: How do you invalidate edge‑cached HTML generated by Node?

**Answer:**
`Purge` API calls to CDN.
Or use Tag-based validation (`Surrogate-Key`).

---

### Question 948: How do you handle real‑time user‑personalized content via edge rendering?

**Answer:**
**ESI** (Edge Side Includes).
Or Cache HTML Shell. Fetch user-data via JS.
Or "Edge Functions" that inject user name into cached HTML before serving.

---

### Question 949: How do you detect bot crawlers for caching logic in Node SSR?

**Answer:**
User-Agent regex.
If Bot: Server Render full page (Cache heavily).
If User: Client Render (Interactive).

---

### Question 950: How do you fallback to live Node render only when cache misses?

**Answer:**
**ISR (Incremental Static Regeneration)**.
CDN attempts serve stale.
Background: Node regenerates page. Updates Cache.
Next user gets new page.
