## 🟢 Internals, Hardware & Experimental (Questions 801-850)

### Question 801: What is the role of the `Bootstrapper` phase in Node.js startup?

**Answer:**
The phase where Node.js sets up the execution environment **before** running user code.
It initializes V8, `libuv`, loads built-in modules (`fs`, `http`), and sets up the global scope (`process`, `console`).

---

### Question 802: How does Node.js initialize the event loop in C++ bindings?

**Answer:**
In `src/node_main.cc`:
1.  Initialize libuv default loop (`uv_default_loop()`).
2.  Create `Environment`.
3.  Start Loop (`uv_run()`).

---

### Question 803: What is libuv and how does Node.js use it under the hood?

**Answer:**
Cross-platform C library for async I/O.
It handles TCP/UDP, Files (thread pool), DNS, and Signals.
Node delegates all I/O to libuv, which notifies Node via callbacks when done.

---

### Question 804: How does Node.js map JavaScript timers to the libuv event loop?

**Answer:**
Node uses a Min-Heap (priority queue) logic in JS (`lib/internal/timers.js`).
It calculates the closest expiry time.
It sets a *single* libuv timer handle (`uv_timer_start`) to wake up at that time.

---

### Question 805: How is the microtask queue managed differently from the macrotask queue?

**Answer:**
Macrotasks (Timers/IO) are processed by libuv phases.
Microtasks (Promises) are processed by V8 (conceptually) immediately after the current stack empties, **before** the next Macrotask starts.

---

### Question 806: What is the internal representation of `Buffer` in memory?

**Answer:**
A `Buffer` is a `Uint8Array`.
Internally backed by an `ArrayBuffer`.
Memory is allocated outside V8 Heap (in C++ land) but points to it.
For small buffers, it might use a slice of a pre-allocated 8KB pool.

---

### Question 807: What is `process.binding()` and when should it be avoided?

**Answer:**
Returns internal C++ bindings (e.g., `process.binding('fs')`).
**Avoid:** It is internal, undocumented, and API changes frequently. Use public modules (`require('fs')`).

---

### Question 808: How does Node.js bridge the V8 heap and native heap?

**Answer:**
`ArrayBuffer` allows shared access.
Or via `External` objects in V8 API, which hold pointers to C++ memory structures.

---

### Question 809: What’s the role of `libcrypto` in Node.js native modules?

**Answer:**
Part of OpenSSL dynamic library linked to Node.
Provides the backing for the `crypto` module (Hashing, HMAC, Cipher).

---

### Question 810: How does Node.js avoid blocking the loop when interfacing with file descriptors?

**Answer:**
Disk I/O is hard to do non-blocking on some OSes.
Libuv maintains a **Thread Pool** (default 4 threads).
FS requests are sent to a thread. The thread blocks. Main loop continues.

---

### Question 811: How do you simulate starvation in the Node event loop?

**Answer:**
Run a heavy synchronous computation.
`while(true) {}`.
Or recurse `process.nextTick`.
Observe that `setInterval` stops firing.

---

### Question 812: What happens if a promise chain throws asynchronously?

**Answer:**
If error bubbles to top and is unhandled: `unhandledRejection` event.
It does **not** crash the process in older Node, but crashes in newer Node (Exit 1).

---

### Question 813: How can `setImmediate()` behave differently on different OS platforms?

**Answer:**
Usually consistent now.
Historically, execution order relative to I/O callbacks could handle differently depending on system load and polling speed.

---

### Question 814: How does the event loop prioritize nextTick vs. setTimeout?

**Answer:**
`nextTick` queue is drained **fully** between phases.
`setTimeout` is checked only in Timers phase.
So `nextTick` always wins (and can block timers).

---

### Question 815: What are the risks of deeply recursive `process.nextTick()` calls?

**Answer:**
Starvation. The event loop never proceeds to the I/O phase. Server stops accepting requests.

---

### Question 816: What are Node.js snapshots and what use cases do they solve?

**Answer:**
(See Q192).
Snapshot built at build-time captures initialized heap.
Use case: CLI tools start instantly. fast-booting microservices.

---

### Question 817: How do you enable and test a new experimental Node.js feature?

**Answer:**
Check `node --help`. Find flag (e.g., `--experimental-test-runner`).
Run: `node --experimental-test-runner app.js`.

---

### Question 818: How does Node.js support policies (`--experimental-policy`)?

**Answer:**
Integrity Checks.
JSON Policy file defines which resources (files) can be loaded by the application.
Prevents processing malicious files if they don't match hash in policy.

---

### Question 819: What is the purpose of the `--frozen-intrinsics` flag?

**Answer:**
Freezes `Array.prototype`, `Object.prototype`, etc.
Prevents 3rd party code/polyfills from modifying built-ins (Prototype Pollution defense).

---

### Question 820: What is the `node:test` module and how does it differ from popular testing libraries?

**Answer:**
Built-in Test Runner (Node 18+).
No need for `npm install jest`.
`import { test, describe } from 'node:test'`.
Fast startup. Supports TAP output.

---

### Question 821: How do you interact with USB devices using Node.js?

**Answer:**
Library `usb` (libusb bindings).
List devices, open endpoints, read/write output reports.

---

### Question 822: What’s the role of `node-hid` and how is it used?

**Answer:**
Interacting with Human Interface Devices (Keyboard, Mouse, Gamepad).
Read events directly from device even if app is in background.

---

### Question 823: How do you read serial port data in Node?

**Answer:**
Library `serialport`.
`const port = new SerialPort({ path: '/dev/tty-usbserial1', baudRate: 9600 })`.
Listen to `data` events.

---

### Question 824: How do you write a Node.js app that responds to GPIO input?

**Answer:**
(On Raspberry Pi).
Library `onoff` or `rpio`.
Watch pin state change (Button press).
`button.watch((err, value) => ...)`

---

### Question 825: How do you integrate Node.js with camera/microphone on Linux?

**Answer:**
Spawn `ffmpeg` to capture stream from `/dev/video0`.
Or use specialized modules like `v4l2camera`.

---

### Question 826: How do you hook into OS-level metrics from Node?

**Answer:**
`os` module (Basic).
For deep stats (CPU temp, Fan speed), use `systeminformation` package.

---

### Question 827: What are the limitations of using `os` module for memory stats?

**Answer:**
`os.freemem()` might be misleading (cached memory is technically freeable but counts as used).
Does not show Node's own contribution vs System.

---

### Question 828: How do you monitor file descriptor usage over time?

**Answer:**
Poller loop.
Read `/proc/self/fd` count (Linux).
`fs.readdirSync('/proc/self/fd').length`.

---

### Question 829: What tools exist for live heap snapshotting in Node?

**Answer:**
`heapdump`.
`inspector` module.
`node-report` (triggers on crash).

---

### Question 830: How do you track garbage collection frequency and duration?

**Answer:**
`perf_hooks`.
monitor `PerformanceObserver` for `gc` entry types.

```javascript
const obs = new PerformanceObserver((list) => {
  console.log(list.getEntries()[0].duration);
});
obs.observe({ entryTypes: ['gc'] });
```

---

### Question 831: How do you control TCP socket options (e.g., `keepAlive`, `noDelay`) in Node?

**Answer:**
`socket.setKeepAlive(true, initialDelay)`.
`socket.setNoDelay(true)` (Disables Nagle's algorithm - sends packets immediately).

---

### Question 832: What’s the effect of `highWaterMark` in backpressure scenarios?

**Answer:**
Determines buffer size.
If buffer fills > highWaterMark, `res.write()` returns `false`.
Sender should pause until `drain` event.

---

### Question 833: How do you detect socket timeouts and resets at runtime?

**Answer:**
`socket.on('timeout', ...)`
`socket.on('error', (err) => { if (err.code === 'ECONNRESET') ... }`.

---

### Question 834: How do you manage half-open sockets properly?

**Answer:**
TCP allows one side to end write but keep read open.
Node `net` socket `allowHalfOpen: true`.
Usually you want `false` (Auto close).

---

### Question 835: What is `SO_REUSEADDR`, and how can you enable it from Node?

**Answer:**
Allows binding to a port in `TIME_WAIT` state (fast restart).
Node Cluster uses this to let multiple workers listen on same port. `net` module enables it by default when clustering.

---

### Question 836: How do you minimize CPU usage in event-driven IoT Node apps?

**Answer:**
Avoid polling. Use interrupts (GPIO watch).
Sleep/Pause execution logic.
Use efficient Streams.

---

### Question 837: How do you queue sensor data reliably during network loss?

**Answer:**
Local persistence (SQLite / LevelDB) on the device.
When Network Up: Drain DB to Cloud.

---

### Question 838: How do you implement OTA (over-the-air) updates in Node?

**Answer:**
Device polls Server for version.
Downloads tar.gz.
Extracts.
Restarts service.
Dual-partition strategy is safer (Flash A, boot A. If fail, boot B).

---

### Question 839: How do you write an energy-efficient sensor polling system?

**Answer:**
Increase interval (Sleep longer).
Process data in batches.
Use hardware triggers when possible.

---

### Question 840: What serialization formats are most efficient for IoT Node apps?

**Answer:**
Binary. **Protobuf**, **MessagePack**, or **CBOR**.
JSON is too verbose and CPU heavy to parse on low-power device.

---

### Question 841: How do you enforce pure functions in a Node.js codebase?

**Answer:**
Linting (`eslint-plugin-functional`).
Immutable data structures (`Immutable.js`).
Code Review.

---

### Question 842: How do you implement lazy evaluation in Node pipelines?

**Answer:**
Streams / Generators.
`function* generate()` yields items one by one. processing loop consumes one by one.
Values aren't computed until pulled.

---

### Question 843: What is a monad and how would you emulate one in Node.js?

**Answer:**
Design pattern: Wrapper around value with `map`/`flatMap`.
Example: Promise is (basically) a Monad. `Promise.resolve(v).then(f)`.

---

### Question 844: How can functional constructs help with async error handling?

**Answer:**
`Ether` / `Maybe` types (libs like `fp-ts`).
Avoids throw/catch. Function returns `Either<Error, Success>`. You map over success path.

---

### Question 845: How do you structure a fully functional Node.js service layer?

**Answer:**
Services are just functions.
Inject dependencies as arguments (Currying).
`const createService = (db) => (id) => db.find(id)`.

---

### Question 846: How do you run TensorFlow or ONNX models from Node.js?

**Answer:**
(See Q522, Q523). Use native bindings libraries.

---

### Question 847: How can you offload AI tasks from Node.js to Python efficiently?

**Answer:**
ZeroMQ / gRPC.
Python service stays running (saving model load time). Node sends data via IPC.

---

### Question 848: What are the tradeoffs of using `@tensorflow/tfjs-node`?

**Answer:**
**Pros:** Fast (C++ backend), GPU support (via distinct package).
**Cons:** Large binary size, `node-gyp` build issues on some OS.

---

### Question 849: How do you use WebAssembly to run ML models in Node?

**Answer:**
Compile model runtime (like ONNX) to Wasm.
Run in Node. Slower than Native C++, but safer (Sandbox) and portable.

---

### Question 850: How do you handle large matrix computations in a streaming Node setup?

**Answer:**
Don't load full matrix. Process row-by-row.
Or use `gpu.js` to offload to GPU.
Or offload to Worker Thread using `SharedArrayBuffer` to avoid copying data.
