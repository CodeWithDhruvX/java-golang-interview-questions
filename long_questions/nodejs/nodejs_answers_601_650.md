## 🟢 Runtimes, Interop & Stability (Questions 601-650)

### Question 601: How do you polyfill Node.js core modules in non-Node runtimes (like Deno)?

**Answer:**
Deno provides the `std/node` compatibility layer.
`import { readFile } from "https://deno.land/std/node/fs.ts";`
Or use specialized polyfill packages like `browserify-zlib` or `events` on npm.

---

### Question 602: What is the `node:` prefix in module imports, and why was it introduced?

**Answer:**
Explicitly identifies Core Modules.
`import fs from 'node:fs';`
Prevents conflict if you accidentally installed a package named `fs` from npm. Also technically faster (skips node_modules resolution).

---

### Question 603: How does Node.js interact with WebAssembly-based modules?

**Answer:**
Native support.
1.  Read `.wasm` file to Buffer.
2.  `WebAssembly.instantiate(buffer, imports)`.
3.  Call exported Wasm functions.
Or use experimental Wasm module support (`import func from './lib.wasm'`).

---

### Question 604: How do you check runtime compatibility for packages targeting Node and Bun?

**Answer:**
Check `engines` field in `package.json`.
Test in CI against both runtimes.
Avoid native addons that rely deeply on V8 C++ API (Bun uses JavaScriptCore).

---

### Question 605: What are the limitations of using ESM-only packages in older Node versions?

**Answer:**
Node < 12 doesn't support ESM.
Node 12 requires `.mjs` or flag.
**Fix:** You must transpile (Babel) ESM code to CommonJS to support older Node versions.

---

### Question 606: How do you handle dynamic import paths in hybrid ESM/CJS apps?

**Answer:**
`import()` works in both.
However, `__dirname` is missing in ESM.
Construct path using `import.meta.url`.
`const path = new URL('./file.js', import.meta.url).pathname;`

---

### Question 607: What tools help ensure cross-runtime compatibility (Node, Deno, Bun)?

**Answer:**
**Unbuild** or **Bundlers** (Rollup).
They can output different formats for different targets.
**WinterCG** (Web-interoperable Runtimes Community Group) defines standard APIs shared by all.

---

### Question 608: What is a runtime shim and when do you need it?

**Answer:**
Code that intercepts API calls to provide missing functionality.
e.g., Shimming `window` or `fetch` in old Node versions to allow frontend code to run on server.

---

### Question 609: How do you simulate browser APIs (like `localStorage`) in Node for SSR?

**Answer:**
Use `node-localstorage` or define a global Mock.
```javascript
global.localStorage = {
  getItem: () => null,
  setItem: () => {}
};
```

---

### Question 610: What happens when you try to use `window` in Node.js?

**Answer:**
`ReferenceError: window is not defined`.
This is often the first error seen when trying to run client-side libraries in SSR without checking `if (typeof window !== 'undefined')`.

---

### Question 611: How do you call a Rust function from Node.js?

**Answer:**
1.  Write Rust function.
2.  Use **Neon** or **napi-rs**.
3.  Build to `.node` binary.
4.  `make` / `require` it in JS.

---

### Question 612: How can you interface Node.js with a Python script?

**Answer:**
1.  **Spawn:** `child_process.spawn('python', ['script.py'])`. Communicate via stdio (JSON strings).
2.  **ZeroRPC:** Use RPC library.
3.  **PythonShell:** Library wrapping spawn.

---

### Question 613: What is `grpc-node` and how does it connect to services written in Go or Java?

**Answer:**
gRPC uses Protobuf (language neutral).
Node acts as gRPC client.
Loads `.proto` file.
Calls GetUser() -> Network -> gRPC Server (Java/Go) -> Returns Protobuf -> Node.

---

### Question 614: How do you expose a Node.js library to be used from Java?

**Answer:**
Usually via HTTP API (REST).
If in-process embedding is needed (GraalVM), Java can execute JS code directly.

---

### Question 615: What are foreign function interfaces (FFIs) and how are they used in Node.js?

**Answer:**
`node-ffi` allow loading dynamic libraries (`.dll`, `.so`) and calling C functions without writing C++ code.
`const lib = ffi.Library('libm', { 'ceil': [ 'double', [ 'double' ] ] });`

---

### Question 616: How do you implement a WebSocket server in pure Node.js?

**Answer:**
Handle the HTTP `Upgrade` event.
Read/Write raw frames (Opcode, Masking, Length).
**Note:** Very complex. Use `ws` library which implements the protocol on top of native http.

---

### Question 617: How would you build an MQTT client in Node?

**Answer:**
Use `mqtt.js`.
`const client = mqtt.connect('mqtt://broker.hivemq.com');`
`client.subscribe('topic');`
MQTT runs over TCP.

---

### Question 618: How do you use Node.js to communicate over raw TCP or UDP?

**Answer:**
*   **TCP:** `net` module. `net.connect(port)`.
*   **UDP:** `dgram` module. `dgram.createSocket('udp4')`.

---

### Question 619: How do you build a gRPC reflection server in Node?

**Answer:**
Use `@grpc/reflection`.
It exposes the available services/methods to clients dynamically (like `grpcurl`), so they don't need the `.proto` file beforehand.

---

### Question 620: How do you serve and consume GraphQL over WebSockets using Node.js?

**Answer:**
Use `graphql-ws` library.
It implements the standard GraphQL-over-WebSocket protocol (subscriptions).
Wraps the `ws` server and handles connection init, start (query), data, stop.

---

### Question 621: How do you isolate tenant data in a Node.js SaaS app?

**Answer:**
1.  **Logical:** `WHERE tenant_id = ?` in every query.
2.  **Schema:** Postgres Schema per tenant.
3.  **Physical:** DB per tenant.

---

### Question 622: What is a tenant-aware database connection, and how do you implement it?

**Answer:**
A function/middleware that resolves the `tenant_id` (from subdomain/header) and returns a DB Client configured for that tenant.

---

### Question 623: How do you sandbox plugins per tenant in a shared Node.js runtime?

**Answer:**
Very hard. (See Q303).
Use **VM2** (deprecated/insecure) or **Isolated-VM**.
Best: Run tenant code in Ephemeral Containers (Docker) or specialized micro-VMs (Firecracker).

---

### Question 624: How do you enforce per-tenant rate limits?

**Answer:**
Redis Keys: `ratelimit:{tenantId}`.
Middleware checks this key.
Allows tiering (Free tenant: 100 req/m, Pro: 1000).

---

### Question 625: How would you structure a monolith-to-microservices migration in Node?

**Answer:**
**Strangler Pattern.**
1.  Put Proxy in front of Monolith.
2.  Build new Service (e.g., Users) in Node.
3.  Update Proxy to route `/users` to new service.
4.  Repeat.

---

### Question 626: How do you validate JWTs with rotating public keys?

**Answer:**
Use JWKS (JSON Web Key Set).
The Authorization Server exposes `/.well-known/jwks.json`.
Node.js resource server fetches this, finds the key matching the JWT's `kid` (Key ID), and verifies signature.
Library: `jwks-rsa`.

---

### Question 627: What are token replay attacks, and how do you prevent them?

**Answer:**
Attacker captures a valid token and uses it.
**Prevention:**
1.  HTTPS (prevent capture).
2.  Short Expiry (limit window).
3.  Token Binding (bind to fingerprint).
4.  JTI (Nonce) tracking (prevent using same token twice - requires state).

---

### Question 628: How do you securely rotate API keys in a Node.js system?

**Answer:**
Support multiple active keys.
1.  Add new key.
2.  Update clients to use new key.
3.  Wait.
4.  Disable old key.

---

### Question 629: How do you implement HMAC signature verification?

**Answer:**
Used for Webhooks (Stripe/GitHub).
1.  Get Payload + Secret.
2.  Compute `crypto.createHmac('sha256', secret).update(payload).digest('hex')`.
3.  Compare with `X-Signature` header.

---

### Question 630: How do you prevent internal SSRF attacks in internal-only Node APIs?

**Answer:**
Validating URLs provided by users.
Block Private IP ranges (127.0.0.1, 192.168.x.x, 10.x.x.x).
Use a library like `ssrf-req-filter`.

---

### Question 631: How do you handle WebSocket connections during a Node.js deploy?

**Answer:**
Hard, because restarting kills connections.
**Strategy:**
1.  Stop accepting new connections.
2.  Wait for existing to close (or Force close after timeout).
3.  Client logic usually handles reconnection (Socket.IO does this auto).
4.  External State (Redis) ensures data isn't lost during reconnect.

---

### Question 632: How do you implement graceful shutdown with open file handles?

**Answer:**
Track open resources.
On `SIGTERM`:
1.  Stop File Watchers.
2.  Flush Write Streams (`stream.end()`).
3.  Wait for `finish` event.
4.  Exit.

---

### Question 633: What is connection draining and how is it used in load-balanced Node apps?

**Answer:**
The Load Balancer stops sending new traffic to the instance being updated.
The instance finishes processing current requests.
Once active connections = 0, the instance is terminated.

---

### Question 634: How do you persist in-flight jobs across rolling deployments?

**Answer:**
Don't keep jobs in RAM. Use a persistent queue (Redis).
If process dies during job, the job times out (lock expires) and is re-queued for another worker.

---

### Question 635: What’s the role of process signaling in safe Node.js shutdown?

**Answer:**
Orchestrator sends `SIGTERM`.
Node app catches `SIGTERM`.
Runs cleanup.
If stuck, Orchestrator sends `SIGKILL` after 30s.

---

### Question 636: How do you batch GraphQL queries on the server?

**Answer:**
**DataLoader**.
Solves N+1 problem.
Collects IDs from multiple resolvers in one tick (`process.nextTick`).
Executes one Batch DB query (`WHERE id IN (...)`).
Distributes results back to promises.

---

### Question 637: What is persisted GraphQL and how does it work in Node.js?

**Answer:**
Client sends a Hash (SHA256) of the query instead of the huge Query string.
Server looks up Hash.
**Pros:** Bandwidth saving, Security (Whitelist only known hashes).

---

### Question 638: How do you enforce query complexity limits?

**Answer:**
Assign "cost" to fields.
Calculate total cost of query before execution.
If > Max, reject.
Plugin: `graphql-query-complexity`.

---

### Question 639: How do you write GraphQL subscriptions using `apollo-server`?

**Answer:**
Define `Subscription` type.
Resolver uses `PubSub` engine.
`subscribe: () => pubsub.asyncIterator(['POST_ADDED'])`.

---

### Question 640: How would you log GraphQL resolver performance?

**Answer:**
Use Apollo Plugins (Tracing).
Or middleware that measures `start` and `end` time of generic resolver functions.

---

### Question 641: How do you implement multilingual support in a Node.js CLI app?

**Answer:**
Detect `LANG` environment variable.
Load corresponding JSON strings.
(See Q584).

---

### Question 642: What is i18next and how is it used in Node apps?

**Answer:**
Popular internalization framework.
Initialize with resources (translations).
`t('key')` returns translated string.
Supports interpolation, pluralization.

---

### Question 643: How do you localize logs and error messages per user context?

**Answer:**
Do **not** translate logs (Keep them English for devs).
Translate Error Messages sent to Client.
Pass `Accept-Language` header to error handler. Return translated message.

---

### Question 644: How do you detect and support right-to-left (RTL) content in SSR?

**Answer:**
Based on Locale.
If `ar` (Arabic) or `he` (Hebrew):
Set `<html dir="rtl">`.

---

### Question 645: What is the role of Unicode normalization in Node.js string handling?

**Answer:**
`'café'` can be written in two ways (composed vs decomposed).
`str.normalize('NFC')` ensures consistent byte representation for comparisons.

---

### Question 646: What is a structured log, and why is it important for observability?

**Answer:**
JSON logs. (See Q356).
Important because it allows searching by fields: `log.duration > 500`. Text logs need regex parsing.

---

### Question 647: How do you implement log sampling in high-throughput Node apps?

**Answer:**
Logger checks random number.
`if (Math.random() < 0.1) logger.info(...)`.
10% sampling.

---

### Question 648: What is the difference between log aggregation and log forwarding?

**Answer:**
*   **Forwarding:** Node writes to `stdout`. Fluentd/Filebeat reads output and pushes to storage. (Best practice).
*   **Aggregation:** Sending logs to a central place (ES/Splunk).

---

### Question 649: How do you correlate logs across multiple Node.js services?

**Answer:**
Trace ID. (See Q357).

---

### Question 650: How do you implement request tracing across async boundaries?

**Answer:**
`AsyncLocalStorage`. (See Q308).
Store the Trace ID on request entry. Retrieve it in logger formatting.
Ensures even `setTimeout` callbacks log the correct Trace ID.
