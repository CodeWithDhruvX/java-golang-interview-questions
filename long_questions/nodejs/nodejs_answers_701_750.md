## 🟢 Hybrid, Service Mesh & Long-Running (Questions 701-750)

### Question 701: How do you use Node.js within a polyglot microservices architecture?

**Answer:**
Node.js acts as one service (e.g., API Gateway or Frontend-For-Backend).
Communicates with other languages (Java, Go, Python) via:
1.  **HTTP/REST:** Standard.
2.  **gRPC:** Efficient internal calls.
3.  **Message Queues:** RabbitMQ/Kafka.

---

### Question 702: How do you communicate between a Node.js app and a service written in Rust?

**Answer:**
If separate services: gRPC or REST.
If within same process: Use FFI (Foreign Function Interface) or compile Rust to `cdylib` and load as Native Addon (`.node`).

---

### Question 703: What is the benefit of using WebAssembly modules from Node.js?

**Answer:**
Performance.
Run heavy CPU logic (Image Processing, Crypto) in Wasm (compiled from C/Rust) near native speed, while using JS for IO/glue code. Secure sandbox.

---

### Question 704: How do you bridge native code and JavaScript via `node-addon-api`?

**Answer:**
Write C++ code using N-API macros.
Compile with `node-gyp`.
`require('./build/Release/addon.node')`.
The wrapper handles converting V8 types (Object, Number) to C++ types.

---

### Question 705: How do you containerize multi-language services with Node.js entrypoints?

**Answer:**
Dockerfile uses multi-stage builds.
Stage 1: Build Rust/Go binary.
Stage 2: Build Node app.
Stage 3: Copy Binary to Node container.
Node uses `child_process.spawn('./binary')` to run handling logic.

---

### Question 706: How do you integrate a Node.js app with Istio?

**Answer:**
Deploy Node app in Kubernetes with Istio sidecar injection enabled.
Istio installs `envoy` proxy next to Node container.
Node app talks to `localhost` (Envoy intercepts traffic).
No code change needed in Node app.

---

### Question 707: How do sidecars affect Node.js app performance?

**Answer:**
Slight network latency (headers overhead, proxy hop).
Increased memory usage in pod (Envoy container).
However, offloads mTLS, Retry, Circuit Breaking logic from the Node Single Thread, potentially **improving** throughput.

---

### Question 708: What headers must a Node.js app forward in a mesh for tracing?

**Answer:**
Zipkin/B3 headers:
`x-request-id`, `x-b3-traceid`, `x-b3-spanid`, `x-b3-parentspanid`, `x-b3-sampled`, `x-b3-flags`, `x-ot-span-context`.
You must copy these from Incoming Req -> Outgoing Req.

---

### Question 709: How do you implement mutual TLS (mTLS) with Node.js in a mesh?

**Answer:**
Ideally, let the Mesh (Istio) handle it.
If doing manually in Node:
`tls.createServer({ requestCert: true, rejectUnauthorized: true, ca: [trustedCA], key, cert })`.

---

### Question 710: How does a Node.js app behave with circuit breaking in Envoy?

**Answer:**
Envoy returns 503 if circuit opens.
Node app sees 503 and should handle it gracefully (Error page / Retry).
Node doesn't need its own Circuit Breaker for that specific route.

---

### Question 711: How do you prevent memory bloat in a long-running Node process?

**Answer:**
Limit cache sizes (LRU).
Close unused connections.
Monitor `process.memoryUsage()`.
Do not attach unlimited listeners (`emitter.on`).

---

### Question 712: How do you write a system daemon using Node.js?

**Answer:**
(Linux) Use `systemd`.
Create `/etc/systemd/system/myapp.service`.
`ExecStart=/usr/bin/node /opt/myapp/index.js`.
`Restart=always`.

---

### Question 713: How do you manage orphaned child processes from a Node app?

**Answer:**
If Parent dies, Child becomes zombie or re-parented to Init.
**Fix:** Explicitly kill children on exit.
`process.on('exit', () => child.kill())`.
Or use a package like `tree-kill`.

---

### Question 714: How do you daemonize a Node.js script on Linux?

**Answer:**
Historically `daemon` module (double-fork).
Modern: Use process managers (PM2) or Systemd.
`pm2 start app.js`.

---

### Question 715: How do you watch a directory and respond to file system events robustly?

**Answer:**
`fs.watch` is flaky across OS.
Use **Chokidar**.
It handles debouncing, varying OS implementations (fsevents on Mac, inotify on Linux), and recursive watching.

---

### Question 716: What flags help tune V8 garbage collection for large heaps?

**Answer:**
`--max-old-space-size=N`: Heap limit.
`--gc-interval=100`: Force GC after allocation.
`--optimize_for_size`: Tradeoff speed for RAM.

---

### Question 717: How do you profile heap growth in production Node apps?

**Answer:**
Take Heap Snapshots periodically (using `heapdump`) or triggering via Signal (`SIGUSR2`).
Or use APM tools (NewRelic) that track Object counts.

---

### Question 718: What’s the difference between minor and major GC in V8?

**Answer:**
*   **Minor (Scavenge):** Fast. Copies surviving young objects to Old Space.
*   **Major (Mark-Sweep):** Slow. Pauses execution. Goes through entire Heap to free space.

---

### Question 719: How do you find and resolve detached DOM-like memory leaks?

**Answer:**
In Node (JSDOM contexts) or Electron.
An object is detached if JS has reference, but it's not in the Tree.
Heap Snapshot shows "Detached DOM tree". NULL the reference to fix.

---

### Question 720: How does object shape affect memory usage in Node?

**Answer:**
Objects with consistent shapes share the "Hidden Class" descriptor.
Objects with random property order (`{a:1, b:2}` vs `{b:2, a:1}`) create separate Hidden Classes, wasting memory and deoptimizing code.

---

### Question 721: How do you handle `Expect: 100-continue` requests?

**Answer:**
Client sends Header + No Body. Waits for 100 Status. Then sends Body.
Node `http` Server handles this automatically (replies 100).
You can listen to `checkContinue` event to add custom logic (Auth check before body upload).

---

### Question 722: What is the impact of `Transfer-Encoding: chunked` on Node streams?

**Answer:**
Node uses it by default for HTTP responses without `Content-Length`.
It allows streaming data to client immediately as `res.write()` is called, keeping connection open.

---

### Question 723: How do you throttle uploads in a Node server?

**Answer:**
Pipe the incoming request stream through a throttler (Transform stream).
`req.pipe(throttle(10kb/s)).pipe(file)`.

---

### Question 724: How do you respond to `OPTIONS` requests for CORS preflights?

**Answer:**
Return Headers: `Access-Control-Allow-Origin`, `Methods`, `Headers`.
Status: `204 No Content`.
`cors` middleware does this.

---

### Question 725: How do you spoof IPs during development for geo-based testing?

**Answer:**
Manually set `X-Forwarded-For` header in Postman/Curl.
Configure middleware to read this header.

---

### Question 726: How do you cache `node_modules` effectively in CI pipelines?

**Answer:**
Key cache by `package-lock.json` hash.
`key: node-modules-${{ hashFiles('package-lock.json') }}`.
If lockfile matches, restore cache. Run `npm ci` (fast, finds existing).

---

### Question 727: What’s the difference between `npm ci` and `npm install` in CI?

**Answer:**
*   **Install:** Can update lockfile. Slower.
*   **CI:** Requires lockfile. Fails if inconsistent. Deletes `node_modules`. Reliable.

---

### Question 728: How do you run partial test suites using tags or filters?

**Answer:**
`jest -t "API Tests"` (Regex match name).
Or split files into folders.

---

### Question 729: How do you automate semantic versioning based on commit messages?

**Answer:**
(See Q652). Semantic Release.
`fix:` -> Patch. `feat:` -> Minor. `BREAKING CHANGE:` -> Major.

---

### Question 730: How do you run tests in parallel across CI runners?

**Answer:**
Use Sharding (Q591).
Or external tools like `Knapsack Pro` to dynamically distribute tests based on timing.

---

### Question 731: How do you implement OAuth 2.1 PKCE flow in a Node.js backend?

**Answer:**
1.  Client generates Code Verifier & Challenge.
2.  Sends Challenge to Auth Server.
3.  Server redirects back with Code.
4.  Client exchanges Code + Verifier for Token.
Backend (Node) just facilitates request calls if it acts as Client.

---

### Question 732: How do you validate signed cookies in stateless JWT auth?

**Answer:**
Cookie has signature (HMAC).
`cookie-parser` verifies signature using secret.
Then verify JWT inside the cookie.

---

### Question 733: What is token chaining and how do you implement it in Node?

**Answer:**
Service A receives Token T1. Calls Service B.
Option 1: Pass T1 (if Audience allows).
Option 2: Exchange T1 for T2 (On-Behalf-Of flow) via Auth Server, then call B.

---

### Question 734: How do you handle refresh token theft in a Node app?

**Answer:**
**Reuse Detection.**
If a Refresh Token is used twice (once by hacker, once by legit user):
Invalidate the **entire chain** (Family). Force login for user.

---

### Question 735: How do you federate identity from multiple providers?

**Answer:**
Use a Gateway Protocol (e.g., Passport.js uses generic profile normalization).
Or use Auth0/Keycloak to handle federation, receiving a standard OIDC token in Node.

---

### Question 736: What’s the difference between hoisting and non-hoisting strategies?

**Answer:**
*   **Hoisting:** Moving common dependencies to Root `node_modules`. (Lerna/Yarn default). Saves space. Risk: Phantom dependencies.
*   **No-Hoist:** Each package has own `node_modules`. Safe isolation. Uses more space.

---

### Question 737: How do you share environment config securely across multiple packages?

**Answer:**
Do not commit `.env`.
Use a shared standard config package that loads from `process.env`.
Inject Env vars at the root CI/Server level.

---

### Question 738: How do you manage package versioning inside a monorepo?

**Answer:**
If "Fixed" mode: All packages satisfy a single version.
If "Independent": Each package has own version.
Use changesets or lerna.

---

### Question 739: How do you publish internal-only npm packages?

**Answer:**
`"private": true` (Never publish).
Or Publish to Private Registry (Artifactory/Github Packages).
Use scoped name `@corp/lib`.

---

### Question 740: How do you avoid dependency hell in a monorepo with many teams?

**Answer:**
Use strict peer dependencies.
Use tools like **Syncpack** to enforce consistent versions of dependencies (e.g., ensure all packages use React 17, not mix of 16/17).

---

### Question 741: How do you build an image CDN using Node.js?

**Answer:**
Route: `/image/:options/:file`.
1.  Check Cache (S3/CDN).
2.  If miss, Fetch Original.
3.  Use **Sharp** to resize/transcode.
4.  Save Cache. Return Stream.

---

### Question 742: How do you transcode audio/video files in a Node worker?

**Answer:**
Use `ffmpeg-static` or `fluent-ffmpeg`.
Spawn process.
`ffmpeg -i input.mp4 output.webm`.
Running in Worker thread doesn't help much as FFMPEG is separate process. `spawn` is non-blocking.

---

### Question 743: What are the performance implications of image manipulation via `sharp`?

**Answer:**
Sharp uses libvips (native C library).
It is much faster and memory efficient than JS-only libraries (Jimp).
Does not block Event Loop significantly (runs in thread pool).

---

### Question 744: How do you bundle WASM modules into Node packages?

**Answer:**
Include `.wasm` file in package.
Use `fs.readFileSync(__dirname + '/lib.wasm')` to load it.

---

### Question 745: How do you offload asset compression to a background process?

**Answer:**
Producer (Upload API) -> Save to S3 -> Send event to Queue (SQS).
Consumer (Worker) -> Download -> Compress -> Re-upload.

---

### Question 746: How do you mock `fs` operations without breaking `realpath`?

**Answer:**
`mock-fs` sometimes breaks native modules using `realpath`.
**Fix:** Use `jest.mock` on `fs` methods selectively, or write files to a temporary directory (`os.tmpdir()`) for real IO tests.

---

### Question 747: How do you simulate backpressure in integration tests?

**Answer:**
Create a slow Writable stream.
`_write(chunk, enc, cb) { setTimeout(cb, 100); }`
Pipe data to it. Verify source stream pauses (`drain` event).

---

### Question 748: How do you test middleware order and side effects?

**Answer:**
Spy on middleware.
`app.use(spy1); app.use(spy2);`
Request.
`expect(spy1).toHaveBeenCalledBefore(spy2)`.

---

### Question 749: How do you test streaming responses in HTTP servers?

**Answer:**
Use a client that reads stream.
`http.get(url, res => { res.on('data', chunk => ... ) })`.
Assert chunks arrive over time.

---

### Question 750: How do you verify telemetry output in tests?

**Answer:**
Mock the transport (e.g., console or HTTP).
`const spy = jest.spyOn(console, 'log');`
Run action.
Check if log contained `{ "trace_id": "..." }`.
