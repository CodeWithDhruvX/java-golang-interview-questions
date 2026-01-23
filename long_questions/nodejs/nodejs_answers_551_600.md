## 🟢 Modern Web & Native Integration (Questions 551-600)

### Question 551: How do you serve pre-rendered HTML in a Node.js SSR app?

**Answer:**
In SSR (e.g., Next.js or Custom Express):
1.  Read the template `index.html`.
2.  Render React/Vue app to String.
3.  Inject string into the template div `<div id="root">...</div>`.
4.  Send final HTML.

---

### Question 552: How do you hydrate browser-side components from a Node server?

**Answer:**
The Node server sends the **Initial State** (JSON data) in a `<script>window.__DATA__ = ...</script>` tag.
The Client JS reads this state to initialize components without re-fetching data, ensuring DOM matches server HTML.

---

### Question 553: How do you use `puppeteer` or `playwright` for browser automation in Node.js?

**Answer:**
Headless Browser automation.
**Code:**
```javascript
const browser = await puppeteer.launch();
const page = await browser.newPage();
await page.goto('https://example.com');
await page.screenshot({ path: 'example.png' });
await browser.close();
```

---

### Question 554: How do you create a browser extension using a Node.js build system?

**Answer:**
Extensions are HTML/JS. Use Node (Webpack/Vite) to bundle the code.
Configure `manifest.json`.
Build step copies output to `dist/` which is loaded into Chrome.

---

### Question 555: How do you bridge WebRTC functionality between Node and browser clients?

**Answer:**
Node.js acts as the **Signaling Server** (using WebSockets) to exchange SDP offers/answers between clients.
Actual media (Video/Audio) goes P2P.
For server-side media processing, use `wrtc` (Node WebRTC bindings).

---

### Question 556: How do you use CSP headers with Node.js?

**Answer:**
Content Security Policy prevents XSS.
Use `helmet` middleware.
```javascript
app.use(helmet.contentSecurityPolicy({
  directives: {
    defaultSrc: ["'self'"],
    scriptSrc: ["'self'", "trusted.com"]
  }
}));
```

---

### Question 557: How do you detect dependency confusion attacks in npm?

**Answer:**
Attackers publish public packages with the same name as your private internal packages.
**Prevent:** Use Scoped packages (`@myorg/utils`) and configure `.npmrc` to strictly map scopes to your private registry.

---

### Question 558: What is token binding and how could it be used in Node.js?

**Answer:**
Binds a session token to a specific TLS connection or Client Certificate.
Hard to hijack token even if stolen.
**Impl:** Validate `req.socket.getPeerCertificate()` matches the claim in the JWT.

---

### Question 559: How do you implement certificate pinning in a Node HTTPS client?

**Answer:**
Explicitly check the server certificate fingerprint during handshake.
`checkServerIdentity` option in `tls.connect` / `https.request`.

---

### Question 560: What is a secure enclave and how can Node.js interact with one?

**Answer:**
Hardware isolated memory (AWS Nitro, Intel SGX).
Node.js cannot access it directly easily. Usually requires a C++ addon or specialized SDK (like AWS KMS) to offload secrets processing to the enclave.

---

### Question 561: What’s the difference between `ffi-napi` and `node-gyp`?

**Answer:**
*   **`node-gyp`:** Compiles C++ code into a native addon (`.node`) during `npm install`. High performance. Harder dev experience.
*   **`ffi-napi`:** Loads dynamic libraries (`.dll`, `.so`) at runtime and calls functions via JS. Slower, but no compilation needed if you have the DLL.

---

### Question 562: How do you wrap a C++ library for use in Node.js?

**Answer:**
Use **N-API** (Node API).
Write C++ wrapper using `napi.h`. Define `Init` function.
Export functions that convert JS Args -> C++ Args -> Return.

---

### Question 563: What are napi threadsafe functions?

**Answer:**
Allow C++ threads (background workers) to call back into JavaScript (Main Thread) safely.
Standard N-API functions can only be called from the JS thread.

---

### Question 564: How does Node.js handle cross-compiling native modules?

**Answer:**
Using `node-pre-gyp`.
You compile binaries for all targets (Win, Mac, Linux) on CI. Upload to S3.
When user installs module, it downloads the correct binary for their OS instead of compiling locally.

---

### Question 565: How do you debug a segmentation fault in a native Node.js module?

**Answer:**
A Segfault crashes the process immediately.
Use `gdb` (Linux) or `lldb` (Mac).
`gdb --args node app.js`.
Run `bt` (backtrace) after crash to see C++ stack.

---

### Question 566: What are the challenges in making a Node app portable across OSes?

**Answer:**
1.  **File Paths:** (`\` vs `/`).
2.  **Line Endings:** (`CRLF` vs `LF`).
3.  **Shell Commands:** (`rm -rf` doesn't work on Win cmd).
4.  **Native Modules:** Need binaries for all OSes.

---

### Question 567: How do you handle Node.js binary compatibility across ARM and x64?

**Answer:**
Node.js (and native modules) must be compiled for the specific architecture (Apple Silicon M1 is ARM64).
Use Docker multi-arch builds (`buildx`) to produce images for both.

---

### Question 568: What is the role of Docker multi-arch builds for Node.js?

**Answer:**
Allows one tag (`my-app:latest`) to run on AWS (x64) and Raspberry Pi/Mac M1 (ARM64).
`docker buildx build --platform linux/amd64,linux/arm64 .`

---

### Question 569: How do you ensure consistent `node_modules` across platforms?

**Answer:**
You generally **cannot** share `node_modules` between OSes (due to binary addons).
**Dev:** Use `npm ci` to install fresh on each OS.
**Docker:** Build inside the container (Linux) to guarantee consistency.

---

### Question 570: What are Windows-specific pitfalls when developing with Node.js?

**Answer:**
1.  **Max Path Length:** 260 chars limit (often hit by nested node_modules).
2.  **Signals:** No SIGTERM.
3.  **Environment Variables:** `SET X=1` vs `export X=1`. Use `cross-env`.

---

### Question 571: How do you implement conditional responses using ETag headers?

**Answer:**
1.  Generate hash (ETag) of response body.
2.  Check `req.headers['if-none-match']`.
3.  If match, return `304 Not Modified` (Empty body).
4.  Else, return `200` + Body + ETag.

---

### Question 572: What is a schema-first approach to API development in Node?

**Answer:**
Define output (OpenAPI/GraphQL Schema) **before** writing code.
Use tools to auto-generate types/routes/validation from the schema.
Ensures API matches documentation.

---

### Question 573: How do you dynamically generate OpenAPI docs from your Node.js app?

**Answer:**
Use `swagger-jsdoc`.
Add JSDoc comments above routes.
The tool scans code and generates `swagger.json` at runtime.

---

### Question 574: How do you structure an API gateway in a Node monorepo?

**Answer:**
Gateway package imports types/schemas from Service packages.
It routes requests to services (via HTTP or function call if monolithic deployment).
Centralizes Auth, Logging.

---

### Question 575: How do you handle feature deprecation safely in APIs?

**Answer:**
1.  Add `Deprecation` header.
2.  Log usage of deprecated endpoints to identify active users.
3.  Support both Old and New versions for X months.

---

### Question 576: How do you create CPU flamegraphs for Node.js?

**Answer:**
1.  Use `0x` tool.
2.  `0x app.js`.
3.  Generates an interactive HTML flamegraph showing CPU stack depth/width.

---

### Question 577: What is an event loop utilization metric?

**Answer:**
`performance.eventLoopUtilization()`.
Returns % of time the loop was active vs idle.
High utilization (>90%) = Loop is blocked/busy.

---

### Question 578: How do you identify high-GC zones in your Node app?

**Answer:**
Run with `--trace-gc`.
Analyze logs. Frequent Scavenge (minor GC) is fine. Frequent Mark-sweep (major GC) indicates memory pressure/leak.

---

### Question 579: What is the role of `v8-profiler-node8` or `clinic.js`?

**Answer:**
**Clinic.js** is a suite (Doctor, Flame, Heaps).
It wraps profiling tools into a GUI to diagnose performance issues (I/O blocks, CPU, Memory).

---

### Question 580: How do you build a performance heatmap for your routes?

**Answer:**
Middleware records `duration` and `route_path`.
Aggregate data (P95, P99) per route.
Vizualize: Red = Slow routes.

---

### Question 581: How do you create an animated CLI using Node.js?

**Answer:**
Use `ora` (spinners) or `listr` (task lists).
They update the same terminal line (using carriage return `\r` and ANSI codes) to create animation.

---

### Question 582: What is an interactive TUI (text UI) in Node.js?

**Answer:**
Rich UI in terminal (Boxes, Buttons, Dashboard).
**Lib:** `blessed` or `ink` (React for CLI).

---

### Question 583: How do you support tab-completion for custom Node CLIs?

**Answer:**
Library `omelette` or `yargs` completion.
It generates a shell script. User adds it to `.bashrc`.
Shell calls your CLI with special flag to get completion options.

---

### Question 584: How do you manage multi-language CLI output in Node.js?

**Answer:**
Use `i18n` library.
Load JSON file based on User Locale.
`console.log(i18n.__('WELCOME'))`.

---

### Question 585: What is the purpose of chalk, ora, and inquirer combined?

**Answer:**
The "Holy Trinity" of Node CLIs.
*   **Chalk:** Colors.
*   **Ora:** Spinners/Loading.
*   **Inquirer:** Input Prompts.

---

### Question 586: What is the impact of WebContainers on Node.js development?

**Answer:**
(StackBlitz technology). Runs Node.js **inside the browser** (using WebAssembly).
Allows instant dev environments without local setup.

---

### Question 587: How do you build native apps using Node.js + Tauri?

**Answer:**
Tauri uses Rust for backend + Webview.
You build your frontend (JS) as usual.
Smaller/Faster than Electron (which bundles Node+Chrome).

---

### Question 588: How do you use EdgeDB or SurrealDB with Node?

**Answer:**
Connect via their official Node.js drivers.
These are modern DBs (Graph-Relational).

---

### Question 589: How does Bun’s compatibility layer affect Node.js libraries?

**Answer:**
Bun creates polyfills for Node APIs (`fs`, `http`, `path`).
Most Node libs work in Bun, but native addons (`node-gyp`) or deep internal APIs (`v8`) might fail.

---

### Question 590: How do you benchmark Node.js vs Rust in network-bound scenarios?

**Answer:**
Use a load tester (`wrk` or `k6`) against minimal servers in both.
Node.js performs surprisingly well (close to Rust) in pure I/O due to ephemeral nature of I/O wait. Rust wins in CPU / Parsing.

---

### Question 591: How do you run test sharding across multiple Node.js runners?

**Answer:**
Split test list into N chunks.
Runner 1 runs Chunk 1.
Jest supports `--shard=1/4`.

---

### Question 592: How do you snapshot test CLI apps?

**Answer:**
Capture `stdout`.
Run CLI with args. Compare output string against saved snapshot file.
**Lib:** `strip-ansi` (remove colors) before comparing.

---

### Question 593: How do you write fuzz tests for JSON input handling?

**Answer:**
Pass random garbage / malformed JSON / huge strings to API.
Ensure app returns 400, **never** crashes (e.g., `JSON.parse` throw).

---

### Question 594: What are smoke tests in the context of Node deployment?

**Answer:**
Lightweight tests run **production** immediately after deploy.
e.g., `GET /health`, `GET /api/me`.
If fail, auto-rollback.

---

### Question 595: How do you test a Node.js app behind a feature flag system?

**Answer:**
Run test suite twice.
1.  `FLAGS=feature=true npm test`
2.  `FLAGS=feature=false npm test`

---

### Question 596: What is the Node.js Technical Steering Committee (TSC)?

**Answer:**
The governing body. They manage high-level direction, disputes, and permissions for the core Node.js project.

---

### Question 597: What are the benefits of LTS in Node.js versions?

**Answer:**
**Long Term Support** (Even numbered versions: 18, 20).
Guaranteed security updates for 30 months. Stable API.
Recommended for Production.

---

### Question 598: How do Node.js release schedules impact library maintainers?

**Answer:**
Maintainers must support all Active LTS versions.
When Node 14 goes EOL, maintainers drop support (major version bump) to use new syntax (node 16 features).

---

### Question 599: How do you become a Node.js module maintainer?

**Answer:**
Start by contributing. Answer issues. Fix bugs. The existing author may grant publish rights.
Or publish your own useful package.

---

### Question 600: What are the biggest current challenges in the Node.js ecosystem?

**Answer:**
1.  **CJS vs ESM split:** Painful transition.
2.  **Competition:** Bun/Deno/Go splitting the user base.
3.  **Security:** Supply chain attacks in NPM.
