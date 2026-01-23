## 🟢 Runtime & Internals (Questions 401-450)

### Question 401: What are native addons in Node.js and how are they written?

**Answer:**
Native addons are shared objects (written in C/C++) that can be loaded into Node.js using `require()`. They provide interface between JS and C/C++ libraries.

**Tools:** `node-gyp`, `N-API`.
**Usage:** High performance calculation, System APIs.

---

### Question 402: What is the role of `node-addon-api`?

**Answer:**
It is a C++ wrapper for N-API (the stable functionality of the V8 API) that isolates addons from changes in the underlying JavaScript engine. It allows addons to work across different Node versions without recompiling.

---

### Question 403: What are bootstrap loaders in Node.js?

**Answer:**
Scripts that run before the main application code to set up the environment (like patching global variables or enabling instrumentation).
**Flag:** `node --require ./setup.js app.js`.

---

### Question 404: How does Node.js handle uncaught exceptions in ES modules?

**Answer:**
Similar to CommonJS, but since ESM is async, top-level errors might be rejected Promises. `uncaughtException` catches synchronous errors. `unhandledRejection` catches async errors.

---

### Question 405: What is the purpose of the `--experimental` flag in Node.js?

**Answer:**
Enables features that are in development and subject to change (e.g., `--experimental-fetch` in older versions). Use with caution in production.

---

### Question 406: What’s the difference between synchronous and asynchronous garbage collection?

**Answer:**
*   **Sync:** STW (Stop-the-World). Pauses execution to clean memory.
*   **Async:** Runs concurrently with the main thread (mostly marking phase) to reduce pause times. V8 uses a mix.

---

### Question 407: How does Node.js integrate with libuv?

**Answer:**
Node.js binds V8 (JS Engine) to libuv (Event Loop).
When you call `fs.readFile` (JS), Node calls C++ bindings, which call `uv_fs_read` (libuv). Libuv uses a thread pool. When done, it pushes callback to loop.

---

### Question 408: What is the difference between `queueMicrotask` and `process.nextTick()`?

**Answer:**
*   **`process.nextTick`**: Specific to Node.js. Runs *before* the next event loop phase. Higher priority.
*   **`queueMicrotask`**: Standard Web API. Runs *after* the current task finishes.

Order: `nextTick` -> `microtask` -> `macrotask`.

---

### Question 409: What does the `--trace-sync-io` flag do?

**Answer:**
It prints a warning/stack trace whenever your code performs synchronous I/O after the first turn of the event loop. Helps detect blocking code in production.

---

### Question 410: How does Node.js maintain a reference to the event loop?

**Answer:**
Specifically, `libuv` maintains a reference count of active handles (timers, sockets).
If `refs > 0`, the process stays alive.
If `refs == 0`, the process exits.
`unref()` decreases this count (allowing exit even if handle is active).

---

### Question 411: How do you force garbage collection in a development environment?

**Answer:**
Run with `node --expose-gc app.js`.
Then call `global.gc()`.

---

### Question 412: What are retained objects in heap snapshots?

**Answer:**
Objects that are not "garbage" because they are reachable from the GC Root (Global, Active Stack).
**Retained Size:** Size of object + size of everything it keeps alive.

---

### Question 413: What is memory fragmentation and how does it impact long-lived apps?

**Answer:**
When free memory is split into small non-contiguous blocks. It forces the OS to allocate new pages even if total free space is sufficient. V8 Compacts memory (moves objects) to fix this.

---

### Question 414: How can memory leaks occur via closures in async code?

**Answer:**
If an async callback captures a large scope, that scope remains in memory until the callback runs. If the callback hangs or is stored in a global list, the scope leaks.

---

### Question 415: What tools does Node.js offer for memory debugging?

**Answer:**
1.  `v8.getHeapStatistics()`
2.  `process.memoryUsage()`
3.  `--inspect` (Chrome snapshots)
4.  `heapdump` module.

---

### Question 416: How would you trace the path of an asynchronous call from start to finish?

**Answer:**
Use `async_hooks` module to assign IDs to execution contexts and log the parent-child relationship of triggers.
Or use APM tools (Datadog) which do this automatically.

---

### Question 417: What is async context propagation?

**Answer:**
Passing data (like UserID) through the async call chain without arguments.
Implemented via `AsyncLocalStorage`.

---

### Question 418: What is a task queue vs. a job queue in Node.js?

**Answer:**
*   **Task Queue (Macrotask):** SetTimeout, I/O.
*   **Job Queue (Microtask):** Promises.
Jobs (Micro) run before Tasks (Macro).

---

### Question 419: What types of timers are handled in the “timers” phase?

**Answer:**
Only `setTimeout` and `setInterval`.
`setImmediate` is handled in "Check" phase.

---

### Question 420: When does a Node.js process _not_ exit automatically?

**Answer:**
When the Event Loop has active handles (Open Server, Database Connection, Interval).
You must `close()` servers or `unref()` timers to allow exit.

---

### Question 421: What is `NODE_OPTIONS` used for?

**Answer:**
An environment variable to pass CLI flags to the Node binary without changing the launch command.
`NODE_OPTIONS='--max-old-space-size=4096' node app.js`

---

### Question 422: What does `--inspect-brk` do?

**Answer:**
Starts the inspector and **pauses** execution at the very first line. Useful for debugging startup scripts.

---

### Question 423: What is the purpose of `--no-deprecation`?

**Answer:**
Silences warning messages about deprecated APIs (e.g., `Buffer()` constructor).
Cleaner logs in production.

---

### Question 424: How do you detect memory limits set via CLI?

**Answer:**
`v8.getHeapStatistics().heap_size_limit`.

---

### Question 425: What’s the impact of `--max-old-space-size`?

**Answer:**
Sets the limit for the V8 Old Space (Major Heap). If usage exceeds this, V8 crashes with OOM.
Default is conservative (approx 2GB). In containers, set it to ~75% of container RAM.

---

### Question 426: How do you watch a directory tree recursively?

**Answer:**
`fs.watch(dir, { recursive: true })` (Supported on Windows/macOS, partial on Linux).
Better: `chokidar.watch(dir)`.

---

### Question 427: How do symbolic links behave in `fs.readdir()`?

**Answer:**
It lists the link name itself, not the target content.
Use `fs.readlink()` to find target.

---

### Question 428: What’s the difference between `fs.readSync()` and `fs.readFileSync()`?

**Answer:**
*   **`fs.readFileSync(path)`**: High-level. Opens, reads, closes, returns content.
*   **`fs.readSync(fd, buffer, ...)`**: Low-level. Requires a File Descriptor (from open). Reads into an existing buffer.

---

### Question 429: How do you safely handle file descriptor limits?

**Answer:**
Use `graceful-fs` (queues `open` calls).
Manually: Implement a queue to limit concurrent open operations.

---

### Question 430: How do you prevent concurrent file writes from corrupting data?

**Answer:**
Use a lock file (`proper-lockfile`).
Or atomic write: Write to `temp.txt` then `fs.rename('temp.txt', 'real.txt')`.

---

### Question 431: How do you create a TCP server in Node.js?

**Answer:**
Using `net` module.

**Code:**
```javascript
const net = require('net');
const server = net.createServer((socket) => {
  socket.write('Echo server\r\n');
  socket.pipe(socket);
});
server.listen(1337);
```

---

### Question 432: What’s the difference between `net.Socket` and `tls.TLSSocket`?

**Answer:**
`tls.TLSSocket` wraps a `net.Socket` and handles Encryption/Decryption logic transparently.

---

### Question 433: How would you handle slowloris attacks at the Node.js level?

**Answer:**
(Slow headers attack).
1.  Set `server.headersTimeout`.
2.  Use Nginx/Reverse Proxy (better suited).
3.  Track connection time and destroy if too slow.

---

### Question 434: How does the `dns` module work differently in `resolve` vs. `lookup`?

**Answer:**
*   **`dns.lookup`**: Uses OS mechanism (`/etc/hosts`, syscall). Same as `ping`. Blocks libuv thread.
*   **`dns.resolve`**: Uses network DNS query. Purely async.

---

### Question 435: What is socket multiplexing and does Node.js support it?

**Answer:**
Sending multiple signals over one socket.
HTTP/2 supports it natively.
WebSockets need sub-protocol/libraries (`mux-demux`).
Native TCP: No, stream is raw bytes.

---

### Question 436: What is the Fetch API in Node.js and how is it different from browsers?

**Answer:**
Included in Node 18+ (Globals). Based on `undici`.
Similar API (`fetch`, `Request`, `Response`).
Differs in details: No cookies jar, handling of CORS (server side doesn't enforce CORS).

---

### Question 437: What are Web Streams and how are they implemented in Node.js?

**Answer:**
Standard streams (`ReadableStream`, `WritableStream`) used in Browsers.
Node.js implemented them to align with Web Standards. You can convert: `stream.Readable.toWeb(nodeStream)`.

---

### Question 438: What is the File API and does Node.js support it natively?

**Answer:**
Web API `File` and `Blob`.
Node 20+ supports `Blob`. `File` is also available. Useful for interop with `fetch(url, { body: formData })`.

---

### Question 439: How is `AbortController` integrated into Node.js APIs?

**Answer:**
Passed as `signal` option.
Supported in `fs` (since recent versions), `net`, `http`, `timers`.
`fs.readFile(path, { signal }, cb)`.

---

### Question 440: What is the `WebAssembly` module in Node.js?

**Answer:**
Native object to compile and instantiate Wasm binaries.
Allows running Rust/C++ code near-native speed.

```javascript
const wasmBuffer = fs.readFileSync('module.wasm');
const { instance } = await WebAssembly.instantiate(wasmBuffer);
instance.exports.add(1, 2);
```

---

### Question 441: What are hidden classes and how do they affect performance?

**Answer:**
V8 creates hidden classes (Shapes) for objects with same property layout.
Monorphic code (same shape) is 100x faster than Polymorphic.
**Tip:** Don't delete properties. Assign `null` instead.

---

### Question 442: What is the cost of polymorphism in hot V8 functions?

**Answer:**
If a function accepts objects of many different shapes (Polymorphic), V8 abandons Inline Caching and does a slow dictionary lookup for properties.

---

### Question 443: How does V8 optimize frequently called functions?

**Answer:**
**TurboFan** compiler.
1.  Interpreter (Ignition) runs bytecode.
2.  If function is hot, TurboFan compiles it to optimized machine code.
3.  If assumption (Types) fails, it De-optimizes back to bytecode.

---

### Question 444: What’s the difference between inline and megamorphic calls?

**Answer:**
*   **Inline:** V8 copies function body to call site (No call overhead).
*   **Megamorphic:** Function called with >4 different shapes. Generic slow path.

---

### Question 445: What are deoptimization triggers in V8?

**Answer:**
1.  Changing argument types (Int -> String).
2.  Adding properties to objects after creation.
3.  Using `try/catch` (in very old V8, now fixed).

---

### Question 446: How do you use `node --inspect` with Chrome DevTools?

**Answer:**
1.  Run `node --inspect script.js`.
2.  Open Chrome `chrome://inspect`.
3.  Click "Open dedicated DevTools for Node".
4.  You can set breakpoints and view memory.

---

### Question 447: What is the role of `v8.getHeapStatistics()`?

**Answer:**
Returns details about V8 memory limits and current usage.
Useful for monitoring if you are close to the limit (`heap_size_limit`).

---

### Question 448: How do you analyze a core dump from a Node.js crash?

**Answer:**
Use `lldb` (Mac/Linux) + `llnode` plugin.
It allows inspecting JS stack frames and objects from the binary core dump file.

---

### Question 449: What does `node --trace-deprecation` show you?

**Answer:**
Prints a stack trace whenever a deprecated API is used, helping you find *where* in your code (or libraries) the old API is being called.

---

### Question 450: How do you view async stack traces in modern Node?

**Answer:**
Node.js enables them by default for Promises/Async-Await.
Example: Error stack will show `async function x` frames even if the error happened in a future tick.
