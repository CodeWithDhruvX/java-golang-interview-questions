## 🟢 Modern Build Tools & Edge (Questions 501-550)

### Question 501: How do you use `esbuild` in a Node.js app?

**Answer:**
`esbuild` is an extremely fast bundler.
**Usage:**
1.  Install: `npm install esbuild`.
2.  Script:
```javascript
require('esbuild').buildSync({
  entryPoints: ['app.js'],
  bundle: true,
  platform: 'node',
  outfile: 'dist/out.js',
});
```

---

### Question 502: What’s the difference between transpiling and bundling?

**Answer:**
*   **Transpiling:** Converting source code (TS, ES6+) to another version (ES5, JS). One-to-one file mapping usually.
*   **Bundling:** Combining multiple files and dependencies into a single (or few) output files.

---

### Question 503: How does Vite handle Node.js APIs?

**Answer:**
Vite is primarily a frontend tool. When building for Node (SSR), it externalizes Node built-ins. Use `vite build --ssr` to produce a bundle compatible with Node runtime.

---

### Question 504: How do you pre-bundle dependencies in a Node.js monorepo?

**Answer:**
Use tools like `tsp` or `esbuild` in a "prebuild" step for shared packages.
Alternatively, use **Yarn Berry** (PnP) or **pnpm** which essentially pre-link dependencies.

---

### Question 505: How do you create a custom plugin for Rollup in a Node.js project?

**Answer:**
A Rollup plugin is an object with hooks.

```javascript
export default function myPlugin() {
  return {
    name: 'my-plugin',
    resolveId(source) {
      if (source === 'virtual-module') return source;
      return null;
    },
    load(id) {
      if (id === 'virtual-module') return 'export default "hi"';
      return null;
    }
  };
}
```

---

### Question 506: What’s the purpose of `tsup` in Node.js development?

**Answer:**
`tsup` is a zero-config bundler powered by `esbuild`. It is designed to bundle TypeScript libraries quickly without needing a complex `webpack` or `rollup` config.

---

### Question 507: How do you reduce cold start time for serverless Node functions?

**Answer:**
1.  **Bundle:** Use `esbuild` to remove unused code/files.
2.  **Lazy Load:** Require heavy modules inside the handler, not top-level (if supported).
3.  **Keep Warm:** Ping the function periodically.
4.  **Use lighter framework:** Avoid Express if possible; use `middy`.

---

### Question 508: How can `SWC` be used with Node.js for fast transpilation?

**Answer:**
`swc` is a Rust-based compiler (alternative to Babel).
**Usage:** `swc-node` allows running TS files directly (like `ts-node` but faster).
`npx swc-node src/index.ts`

---

### Question 509: What’s the benefit of lazy importing modules?

**Answer:**
Reduces startup time and memory usage. Modules are only loaded when the specific code path is executed.
`const heavy = require('heavy-lib')` inside a route handler.

---

### Question 510: How do you structure a hybrid TypeScript and JavaScript Node app?

**Answer:**
Enable `allowJs: true` in `tsconfig.json`.
TS compiles JS files too (or leaves them).
Gradually rename `.js` to `.ts` adding types.

---

### Question 511: How does Node.js differ when deployed at the edge?

**Answer:**
Edge runtimes (Cloudflare, Vercel Edge) often run on **V8 Isolates**, not full Node.js.
**Differences:** No `fs` access, no native C++ modules, standard Web Fetch API instead of `http` module.

---

### Question 512: What are the Node.js limitations in Cloudflare Workers?

**Answer:**
1.  **No `eval()`** or `new Function`.
2.  **No Native Modules.**
3.  **Limited `Buffer` support** (use `Uint8Array`).
4.  **No TCP Sockets** (until recently with `connect()` API).

---

### Question 513: How do you handle cold starts in AWS Lambda with Node.js?

**Answer:**
(See Q507).
Also, enable **Provisioned Concurrency** (AWS feature) to keep initialized instances ready.

---

### Question 514: How do you manage shared dependencies in a serverless bundle?

**Answer:**
Use **Lambda Layers**.
Create a layer with `node_modules`. Attach it to multiple functions.
Reduces deployment size of individual functions.

---

### Question 515: How do you reduce package size for edge deployment?

**Answer:**
1.  Dev dependencies excluded.
2.  Tree-shaking (esbuild/webpack).
3.  Minification.
4.  Replace large libs (`moment`) with small ones (`dayjs`).

---

### Question 516: How do you handle sessions in stateless serverless functions?

**Answer:**
Sessions cannot be stored in-memory (RAM).
**Solution:** Store Session Data in external fast DB (Redis/DynamoDB) and pass Session ID via Cookie/Header. Or use **JWT** (Client-side session).

---

### Question 517: What’s the difference between regional and edge Node.js deployments?

**Answer:**
*   **Regional:** Runs in specific data centers (e.g., us-east-1). High latency for distant users. Full Node environment.
*   **Edge:** Runs in hundreds of cities closer to user. Low latency. Restricted runtime (Isolates).

---

### Question 518: How do you maintain execution context across stateless requests?

**Answer:**
You cannot. Each request is independent.
You must pass context state via DB or Client (Tokens).
If you need stateful "Actors", use **Durable Objects** (Cloudflare) or **Azure Durable Functions**.

---

### Question 519: What’s the impact of V8 isolates in edge runtimes?

**Answer:**
Isolates share the same process/memory but have separate heaps.
**Pros:** Instant startup (milliseconds), low overhead.
**Cons:** Cannot share mutable state easily between isolates without external store.

---

### Question 520: What are best practices for debugging serverless Node.js apps?

**Answer:**
1.  **Local Emulation:** `serverless-offline` or `sam local`.
2.  **Structured Logs:** CloudWatch.
3.  **Distributed Tracing:** X-Ray or OpenTelemetry.

---

### Question 521: How can Node.js consume a RESTful AI model?

**Answer:**
Using `fetch` or `axios` to call the Python/FastAPI model server.
Send JSON/Image, receive prediction.

---

### Question 522: What’s the purpose of `onnxruntime-node`?

**Answer:**
Runs ONNX (Open Neural Network Exchange) models natively in Node.js.
Allows you to run models trained in PyTorch/TensorFlow directly in Node without Python.

---

### Question 523: How do you run inference from a TensorFlow.js model in Node?

**Answer:**
Install `@tensorflow/tfjs-node`.
Load model: `await tf.loadGraphModel('file://model.json')`.
Predict: `model.predict(tensor)`.
It binds to TensorFlow C binary for speed.

---

### Question 524: How do you handle streaming OpenAI responses in Node?

**Answer:**
Set `stream: true` in API call.
Handle response as a stream.

```javascript
const res = await openai.createCompletion({ stream: true, ... }, { responseType: 'stream' });
res.data.on('data', data => console.log(data.toString()));
```

---

### Question 525: How do you fine-tune a chatbot using Node.js tooling?

**Answer:**
Node.js acts as the orchestrator.
1.  Prepare JSONL dataset.
2.  Upload file via OpenAI API (Node SDK).
3.  Trigger Fine-tune job.

---

### Question 526: How do you implement rate limiting when accessing AI APIs?

**Answer:**
Use a token bucket library (`limiter`).
OpenAI has strict RPM/TPM limits. Implement exponential backoff retry logic.

---

### Question 527: What’s the role of `worker_threads` in running AI tasks in Node.js?

**Answer:**
AI inference (even with native bindings) is CPU heavy.
Run it in a Worker Thread to prevent blocking the main Express/HTTP server loop.

---

### Question 528: How do you integrate a HuggingFace model with Node.js?

**Answer:**
Use **Hugging Face Inference API** (HTTP).
Or use `@xenova/transformers` (JavaScript port of Transformers library) to run models locally in Node.

---

### Question 529: What are the memory implications of running large models in Node?

**Answer:**
Large models (GBs) might crash V8.
**Fix:** Run outside JS Heap (C++ bindings usually handle this, utilizing system RAM). Ensure machine has enough RAM.

---

### Question 530: How do you serialize/deserialize AI model input/output in Node?

**Answer:**
Inputs are usually Tensors (Arrays of Numbers).
Convert Buffer/Image -> Float32Array -> Tensor.

---

### Question 531: How do you design a pluggable architecture in Node.js?

**Answer:**
Define a standard Interface (e.g., `start()`, `stop()`).
Load plugins from a folder using `require()`.
Store them in a list and iterate.

---

### Question 532: How would you build a plugin loader using dynamic imports?

**Answer:**
```javascript
async function loadPlugins(dir) {
  const files = fs.readdirSync(dir);
  return Promise.all(files.map(f => import(path.join(dir, f))));
}
```

---

### Question 533: What is sandboxing, and how can it apply to plugins?

**Answer:**
Running untrusted plugin code in a restricted environment.
**Tools:** `vm` module (weak), or **Wasm** (strong sandbox).

---

### Question 534: How do you isolate plugin execution to prevent memory leaks?

**Answer:**
Run each plugin in a separate **Child Process** or **Worker Thread**.
If it leaks or crashes, you can kill/restart just that worker without affecting the main app.

---

### Question 535: How do you dynamically load and unload modules safely?

**Answer:**
1.  **Load:** `require` / `import`.
2.  **Unload:** Remove from `require.cache` (CJS).
    *   *Note: ESM modules cannot be unloaded from memory in V8 currently.* Potential leak source.

---

### Question 536: What are red metrics and how do you apply them to Node.js?

**Answer:**
**RED**: Rate, Errors, Duration.
Middleware tracks:
1.  Count requests (Rate).
2.  Count 5xx (Errors).
3.  Measure Response Time (Duration).
Expose via Prometheus.

---

### Question 537: How do you track async performance bottlenecks?

**Answer:**
Trace Event Loop Lag. If high, something is blocking.
Use `perf_hooks` to wrap async functions and measure real duration vs CPU duration.

---

### Question 538: What are the challenges in tracing requests through worker threads?

**Answer:**
Context (Trace ID) is lost when crossing thread boundary.
**Fix:** Pass the Trace ID explicitly in the `postMessage` payload and restore it in the worker's AsyncLocalStorage.

---

### Question 539: How would you log structured user actions across multiple services?

**Answer:**
Standard schema (JSON).
`{ event: "ORDER_PLACED", user: "u1", trace: "t1", service: "s1" }`.
Send to centralized Log Store (Elasticsearch).

---

### Question 540: What is distributed tracing and how do you implement it with OpenTelemetry?

**Answer:**
Visualizing a request as it hops services.
**Node:** `opentelemetry/sdk-node`.
Auto-instrumentation patches `http` to inject `traceparent` headers into outgoing requests.

---

### Question 541: How do you prevent resource exhaustion in a long-lived Node.js process?

**Answer:**
1.  **Memory:** restart periodically (PM2 max-memory-restart).
2.  **Sockets:** Set timeouts.
3.  **Descriptors:** Graceful-fs.

---

### Question 542: What happens if you perform thousands of concurrent file reads?

**Answer:**
`EMFILE`: Error/Too many open files.
Thread pool saturation: Other async tasks (DNS) stall.

---

### Question 543: How do you protect your Node.js app against infinite request loops?

**Answer:**
(e.g., Service A calls B, B calls A).
**Prevention:** Pass a `Loop-Detect` header count. Decrement it. If 0, reject.

---

### Question 544: How would you simulate an out-of-memory crash for testing?

**Answer:**
```javascript
const dump = [];
setInterval(() => {
  dump.push(alloc(10 * 1024 * 1024)); // Alloc 10MB repeatedly
}, 100);
```

---

### Question 545: How do you build a recovery strategy after an API rate limit block?

**Answer:**
Respect `Retry-After` header.
Implement Circuit Breaker (stop calling for X minutes).

---

### Question 546: How do you apply feature flags at runtime in Node.js?

**Answer:**
Poll a config server (LaunchDarkly, Firebase Remote Config) or DB.
If `flags.newFeature` is true, run new code path.

---

### Question 547: What is a dynamic configuration system and how can Node.js support it?

**Answer:**
Allows changing config (LogLevel, Timeouts) without restart.
**Impl:** Watch a `config.json` file key or Etcd key. Update global config object on change.

---

### Question 548: How do you handle live configuration reloads without restarting the app?

**Answer:**
Use a Singleton Config class.
When update event occurs, update the class properties.
Ensure code reads from the class, not a const captured at startup.

---

### Question 549: How do you implement a “kill switch” for a Node.js feature in production?

**Answer:**
A Feature Flag that defaults to TRUE. If bug found, set to FALSE in DB/Redis. API checks flag and returns 503 or legacy behavior.

---

### Question 550: What’s the role of circuit breakers in fault-tolerant Node.js design?

**Answer:**
Prevents cascading failures. If Service B is failing fast, Service A stops calling it and returns default error immediately, saving resources on both sides.
**Lib:** `opossum`.
