## 🟢 Concurrency, Jobs & Tooling (Questions 451-500)

### Question 451: What is the `Atomics` module and how does it work with shared memory?

**Answer:**
`Atomics` provides atomic operations (add, load, store, wait) on a `SharedArrayBuffer` to avoid race conditions between Worker Threads.

**Code:**
```javascript
const buffer = new SharedArrayBuffer(4);
const u8 = new Uint8Array(buffer);
// Thread 1
Atomics.add(u8, 0, 1);
// Atomically ensures no other thread interrupts the addition
```

---

### Question 452: What is the purpose of `SharedArrayBuffer` in Node.js?

**Answer:**
It allows multiple worker threads to read/write to the **same** block of memory. This is zero-copy communication, much faster than `postMessage`.

---

### Question 453: How do you implement a mutex in Node.js?

**Answer:**
Using `Atomics.wait` and `Atomics.notify` on an Int32Array.

**Concept:**
1.  Try to change lock value 0 -> 1 (compareExchange).
2.  If fail, `Atomics.wait`.
3.  When owner releases, `Atomics.notify`.

---

### Question 454: How can you coordinate multiple async workers to share a task queue?

**Answer:**
1.  **Shared Memory:** Implement a lock-free queue in SharedArrayBuffer (Hard).
2.  **Coordinator:** Main thread holds the queue, sends items to workers when they say "ready".
3.  **External:** Redis List.

---

### Question 455: When should you use worker threads over a message broker?

**Answer:**
*   **Worker Threads:** For CPU intensive tasks within the *same* machine/instance. Low latency.
*   **Message Broker (RabbitMQ):** For distributing work across *multiple* machines/services. Durability.

---

### Question 456: What’s the difference between `bull`, `kue`, and `agenda`?

**Answer:**
*   **Bull:** Redis based. Fast. Active.
*   **Kue:** Redis based. Archived/Deprecated.
*   **Agenda:** MongoDB based. simpler setup if you already use Mongo.

---

### Question 457: How do you persist job states across restarts?

**Answer:**
Use a persistent queue library (Bull/Redis). The jobs are stored in the DB. When Node restarts and queue consumer starts, it picks up pending jobs. In-memory queues lose data.

---

### Question 458: How do you ensure at-least-once job delivery?

**Answer:**
The worker must **acknowledge** the job only after processing is done.
If worker crashes, the queue detects timeout/lack of ack, and re-queues the job.
**Idempotency** is required in the worker logic.

---

### Question 459: How would you build a delayed job scheduler without external dependencies?

**Answer:**
Use `setTimeout`.
**Problem:** Volatile (lost on restart).
**Solution:** Store task (Time, Data) in a persistent file/DB. On startup, read DB and `setTimeout` for remaining duration.

---

### Question 460: How do you monitor background job failure rates?

**Answer:**
Subscribe to Queue events.
`queue.on('failed', (job, err) => alertSystem(err));`
Push metrics to Prometheus/Datadog.

---

### Question 461: How do you read a binary file byte-by-byte?

**Answer:**
Read into a Buffer, then access via index.

**Code:**
```javascript
const buf = fs.readFileSync('file.bin');
for (const byte of buf) {
  console.log(byte); // 0-255
}
```

---

### Question 462: How do you convert a binary buffer to a hexadecimal string?

**Answer:**
`buffer.toString('hex')`.

**Code:**
```javascript
const b = Buffer.from([255, 0, 15]);
console.log(b.toString('hex')); // 'ff000f'
```

---

### Question 463: What’s the difference between `utf8` and `utf16le` in Node?

**Answer:**
Encoding schemes.
*   **UTF-8:** Variable width (1-4 bytes). Standard for Web.
*   **UTF-16LE:** 2 or 4 bytes. Used by Windows internally.
Node Buffers support both.

---

### Question 464: How does base64 encoding inflate data size?

**Answer:**
Base64 uses 4 characters to represent 3 bytes.
Inflation is roughly **33%** (4/3).

---

### Question 465: How do you detect and handle corrupted binary input?

**Answer:**
Use Checksums (CRC32, MD5) or Magic Numbers (File Signatures).
Read first few bytes to verify header (e.g., PNG starts with `89 50 4E 47`).

---

### Question 466: What is a dual-module package?

**Answer:**
A package that supports both CommonJS (`require`) and ESM (`import`) consumers.
Implemented using `exports` in package.json.

---

### Question 467: How do you publish both CommonJS and ESM versions of a module?

**Answer:**
Transpile TS/Source to two folders: `dist/cjs` and `dist/mjs`.
Point `exports` to them.

```json
"exports": {
  "import": "./dist/mjs/index.js",
  "require": "./dist/cjs/index.js"
}
```

---

### Question 468: What is the `exports` map in `package.json` used for?

**Answer:**
(See Q318). Defines public entry points. Prevents deep imports unless explicitly exposed.

---

### Question 469: How do you bundle a Node library using Rollup or esbuild?

**Answer:**
Tools usually for frontend, but useful for Node to minifiy or bundle dependencies.
Config target: `node`. Format: `cjs` or `esm`.
Mark `dependencies` as `external` to avoid bundling `node_modules`.

---

### Question 470: What’s the difference between a named export and a default export in ESM?

**Answer:**
*   **Named:** `export const x = 1`. Import with `{ x }`. Tree-shakeable.
*   **Default:** `export default x`. Import with `anyName`. Harder to refactor.

---

### Question 471: How do you handle Windows file paths vs POSIX paths?

**Answer:**
Windows uses Backslash `\`. POSIX uses Forward Slash `/`.
**Fix:** Always use `path.join()` or `path.normalize()`.
`path.join('a', 'b')` -> `a/b` or `a\b` automatically.

---

### Question 472: How do you detect the current platform in Node.js?

**Answer:**
`process.platform`.
Values: `'darwin'`, `'freebsd'`, `'linux'`, `'openbsd'`, `'sunos'`, `'win32'`.

---

### Question 473: How do you set environment variables cross-platform?

**Answer:**
Use `cross-env` package.
`"start": "cross-env NODE_ENV=production node app.js"`
Works on Windows (cmd/powershell) and Linux (bash).

---

### Question 474: How does signal handling differ on Windows vs Linux in Node?

**Answer:**
Windows does not support signals like `SIGTERM` or `SIGINT` natively in the same way.
However, Node.js emulates `SIGINT` (Ctrl+C). `SIGTERM` support is limited on Windows.

---

### Question 475: What are EOL issues and how do you handle them?

**Answer:**
End of Line characters.
Windows: `CRLF` (`\r\n`).
Linux: `LF` (`\n`).
**Fix:** Use `.gitattributes` to enforce LF. Use `os.EOL` constant in code when writing strings.

---

### Question 476: What is the `perf_hooks` module used for?

**Answer:**
(See Q194). Performance measurement API compatible with W3C Performance Timing API.
High resolution timing, Performance Observers.

---

### Question 477: What is the `readline/promises` API?

**Answer:**
New in Node 17. Allows reading input line-by-line using `await`.

```javascript
const rl = require('readline/promises').createInterface({ input });
const answer = await rl.question('What is your name? ');
```

---

### Question 478: How do you use the `url` module for parsing and formatting?

**Answer:**
Use the WHATWG `URL` class (Global).

```javascript
const myUrl = new URL('https://example.com:8000/path?name=User');
console.log(myUrl.hostname); // example.com
myUrl.searchParams.append('age', '20');
```

---

### Question 479: How does the `assert` module handle deep comparisons?

**Answer:**
`assert.deepStrictEqual(actual, expected)`.
Recursively checks own enumerable properties. Primitives are compared with `===`.

---

### Question 480: What is the purpose of the `tty` module?

**Answer:**
Detects if Node is running in a Text Terminal (TTY).
`process.stdout.isTTY` returns true if outputting to console, false if piped to file (`node app.js > out.txt`). Use this to toggle colors off for files.

---

### Question 481: How do you test code that uses `process.exit()`?

**Answer:**
Stub `process.exit`.
**Jest:**
```javascript
const mockExit = jest.spyOn(process, 'exit').mockImplementation(() => {});
myFunction();
expect(mockExit).toHaveBeenCalledWith(1);
```

---

### Question 482: How do you mock system time or timezone?

**Answer:**
Time: Fake Timers (Q366).
Timezone: Set `process.env.TZ = 'UTC'` before running tests (works on Limit/Mac). Or use mocking libraries.

---

### Question 483: How do you assert log output during tests?

**Answer:**
Spy on `console.log` or your logger's transport.

```javascript
const spy = jest.spyOn(console, 'log');
fn();
expect(spy).toHaveBeenCalledWith('Success');
```

---

### Question 484: How do you create dynamic test cases programmatically?

**Answer:**
Loop inside the `describe` block.

```javascript
[1, 2, 3].forEach(num => {
  test(`should handle ${num}`, () => {
    expect(fn(num)).toBe(true);
  });
});
```

---

### Question 485: What’s the benefit of property-based testing in Node?

**Answer:**
Libraries like `fast-check` generate thousands of random inputs to find edge cases you didn't think of.

---

### Question 486: What are some alternatives to npm (other than yarn)?

**Answer:**
*   **pnpm:** Uses hard links/symlinks to save disk space. Very fast.
*   **cnpm:** China mirror.
*   **Bun / Deno:** Have their own installers.

---

### Question 487: What is a monorepo and how do tools like Lerna or Nx help manage it?

**Answer:**
(See Q165).
Lerna: Bootstraps packages (links them).
Nx: Builds graph, caches tasks, runs affected tests only.

---

### Question 488: What’s the role of `npx` in rapid tooling?

**Answer:**
Runs binaries from npm registry without installing.
`npx cowsay hello` -> downloads cowsay -> runs it -> deletes it.

---

### Question 489: How do you consume private npm packages securely?

**Answer:**
Use `.npmrc` with an Auth Token.
`//registry.npmjs.org/:_authToken=${NPM_TOKEN}`
Inject token in CI via Env Vars.

---

### Question 490: What are the differences between CommonJS loaders and ES module loaders?

**Answer:**
*   **CJS Loader:** Reads file, wraps in function `(function(exports, require...))`, executes. Sync.
*   **ESM Loader:** Parsed into AST, Dependency Graph built, Linked, then Executed. Async. Hooks available (`--loader`).

---

### Question 491: How do you run Node.js in a Deno environment (if at all)?

**Answer:**
Deno now supports `npm:` specifiers and has Node compatibility layer (`std/node`).
`import fs from 'node:fs';` works in Deno.

---

### Question 492: What are the challenges in porting a Node.js project to ESM-only?

**Answer:**
1.  Cannot use `require`.
2.  `__dirname` is missing.
3.  Dependencies might be CJS-only (usually fine via default import).
4.  Mocking libraries based on `require.cache` fail.

---

### Question 493: What’s the role of edge computing in Node.js backends?

**Answer:**
Running Node.js (or subset) close to the user (CDN Edge). Reduces latency.
examples: Lambda@Edge, Cloudflare Workers (V8 Isolate, not full Node).

---

### Question 494: How do you run Node.js in a Cloudflare Worker or similar runtime?

**Answer:**
You don't run *full* Node.js. You run V8 compatible JS.
However, recent "Node.js Compatibility" flags in Workers allow using `AsyncLocalStorage`, `EventEmitter`, `Buffer`.

---

### Question 495: How does Node.js compare with Bun in cold start times?

**Answer:**
Bun is significantly faster (written in Zig, focused on startup).
Node.js startup is slower but improving (Snapshots).

---

### Question 496: How do you precompile TypeScript for Node.js deployment?

**Answer:**
Run `tsc` (TypeScript Compiler) during build phase.
Deploy the output `dist/` folder (JS files) and `package.json`. Do not run `ts-node` in production.

---

### Question 497: What are build-time vs. runtime dependencies?

**Answer:**
*   **Build-time:** TypeScript, Webpack, Babel. (Saved in `devDependencies`). Not needed in Prod image.
*   **Runtime:** Express, Lodash. (Saved in `dependencies`). Needed in Prod.

---

### Question 498: How do you automate changelogs with Node.js tools?

**Answer:**
Use **Standard Version** or **Semantic Release**.
They analyze commit messages (Conventional Commits) and generate `CHANGELOG.md`.

---

### Question 499: What is a zero-dependency package and why does it matter?

**Answer:**
A package with `dependencies: {}`.
Matters for security (Supply Chain Attacks) and install speed/size.

---

### Question 500: What are best practices for publishing scoped packages?

**Answer:**
Scoped packages (`@myorg/pkg`) are private by default.
To publish publicly:
`npm publish --access public`.
Naming: Use scopes to avoid name collision.
