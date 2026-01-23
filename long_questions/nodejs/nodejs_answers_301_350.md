## 🟢 Uncommon Concepts & Internals (Questions 301-350)

### Question 301: What happens when you `require()` a JSON file?

**Answer:**
Node.js treats it as a JSON file, parses the content using `JSON.parse()`, and returns the resulting JavaScript object.

**Code:**
```javascript
// config.json: { "port": 3000 }
const config = require('./config.json');
console.log(config.port); // 3000
```
*Note: This is synchronous and cached.*

---

### Question 302: How does Node.js internally handle DNS lookups?

**Answer:**
It uses `c-ares` (a C library) for non-blocking DNS lookups. However, `dns.lookup` (used by `net.connect`) uses the system resolver via the thread pool (blocking mechanism managed by libuv), while `dns.resolve` uses c-ares directly.

---

### Question 303: What is the `vm` module and when would you use it?

**Answer:**
The `vm` module allows compiling and running code within V8 virtual machine contexts.
**Use Case:** Sandboxing code (e.g., running user-submitted scripts), but it is **not secure** for untrusted code.

**Code:**
```javascript
const vm = require('vm');
const context = { x: 2 };
vm.createContext(context);
vm.runInContext('x += 40', context);
console.log(context.x); // 42
```

---

### Question 304: Can you dynamically load a module at runtime?

**Answer:**
Yes.
*   **CommonJS:** `require(variable)` works anywhere.
*   **ESM:** `await import(variable)` works anywhere.

---

### Question 305: What is module hot-reloading and how do you implement it?

**Answer:**
Hot-reloading means updating the module code without restarting the process.
**Mechanism:** Delete the module from `require.cache` and re-require it.

**Code:**
```javascript
function load() {
  delete require.cache[require.resolve('./module')];
  return require('./module');
}
```

---

### Question 306: How do you share a single instance of a module across files?

**Answer:**
Node.js modules are cached. The first time `require('./a')` runs, it executes. The second time, it returns the *same object*.
So, simply capturing state in the module scope makes it a Singleton.

**store.js:**
```javascript
const data = [];
module.exports = { data };
```
**a.js:** `require('./store').data.push(1)`
**b.js:** `require('./store').data` has `[1]`.

---

### Question 307: What is the `inspector` module used for?

**Answer:**
It allows you to attach to the running V8 inspector from within the code to profile execution or take heap snapshots.

**Code:**
```javascript
const inspector = require('inspector');
const session = new inspector.Session();
session.connect();
// Send commands to V8...
```

---

### Question 308: How do you track asynchronous context in Node.js?

**Answer:**
Using `AsyncLocalStorage` (from `async_hooks` module). It allows storing data (like Request ID) that is accessible deeply within the async call stack without passing it as arguments.

**Code:**
```javascript
const { AsyncLocalStorage } = require('async_hooks');
const storage = new AsyncLocalStorage();

storage.run({ requestId: 123 }, () => {
  // Deeply nested function
  console.log(storage.getStore().requestId); // 123
});
```

---

### Question 309: What is the significance of `globalThis` in Node.js?

**Answer:**
`globalThis` is a standard property providing the global `this` value across environments (Browser `window`, Node `global`). It makes code portable.

---

### Question 310: What is a symbol and how is it used in Node.js?

**Answer:**
A Symbol is a unique and immutable primitive. It is used to create "hidden" or unique object properties that don't appear in standard iteration.

**Usage:**
```javascript
const kKeepAlive = Symbol('keep-alive');
class Request {
  [kKeepAlive] = true;
}
```

---

### Question 311: What does `process.exitCode = 1` do compared to `process.exit(1)`?

**Answer:**
*   **`process.exit(1)`**: Forces the process to terminate **immediately**, potentially truncating stdout/stderr writes.
*   **`process.exitCode = 1`**: Sets the exit code but allows the process to exit naturally (when the event loop empties). **Preferred**.

---

### Question 312: What is the use of `process.stdin.setRawMode()`?

**Answer:**
It switches the terminal to "raw mode", where input is passed character-by-character (including special keys like Arrows) instead of line-by-line. Used for building interactive CLIs.

---

### Question 313: What are the arguments available in `process.argv`?

**Answer:**
An array containing command-line arguments.
1.  Path to Node executable.
2.  Path to Script file.
3.  ...Args.

---

### Question 314: How can you access the current memory usage of a Node.js process?

**Answer:**
`process.memoryUsage()`.
Returns object with: `rss` (Resident Set Size), `heapTotal`, `heapUsed`, `external`.

---

### Question 315: What happens if a Node.js script exceeds the call stack size?

**Answer:**
It throws a `RangeError: Maximum call stack size exceeded`.
This usually happens with infinite recursion.

---

### Question 316: What are conditional exports in `package.json`?

**Answer:**
It allows a package to export different files based on how it is imported (`require` vs `import`) or the environment (`node` vs `browser`).

**package.json:**
```json
"exports": {
  "require": "./index.cjs",
  "import": "./index.mjs"
}
```

---

### Question 317: How does Node.js resolve module paths?

**Answer:**
1.  Core module? (http).
2.  File? (./ start).
3.  `node_modules` folder (current dir, then parent, then parent...).
4.  Global modules (rarely used now).

---

### Question 318: What is `exports` field in package.json and how does it differ from `main`?

**Answer:**
*   **`main`**: Defines the entry point (legacy).
*   **`exports`**: Modern (Node 12+). Encapsulates the package. Only files explicitly listed in `exports` can be imported. It prevents users from importing internal files (`pkg/lib/private.js`).

---

### Question 319: Can Node.js load `.mjs` and `.cjs` files in the same project?

**Answer:**
Yes. Node determines the module system by extension:
*   `.mjs` -> ESM
*   `.cjs` -> CommonJS
*   `.js` -> Depends on `"type"` in `package.json`.

---

### Question 320: How does dynamic `import()` differ from static `require()`?

**Answer:**
`import()` returns a **Promise**. It loads the module asynchronously. `require()` loads synchronously.

---

### Question 321: How do you handle stream piping errors?

**Answer:**
`src.pipe(dest)` does **not** forward errors. If `src` errors, `dest` is not closed.
**Fix:** Use `stream.pipeline(src, dest, cb)`.

---

### Question 322: What’s the difference between `highWaterMark` and backpressure?

**Answer:**
*   **highWaterMark:** The threshold (size in bytes) at which the stream buffer is considered "full".
*   **Backpressure:** The mechanism (pausing reading) triggered when the buffer hits the `highWaterMark`.

---

### Question 323: How do you convert a stream to a buffer or string?

**Answer:**
You must consume the stream.
**Node 17+:** `streamConsumers.buffer(stream)`.
**Manual:**
```javascript
const chunks = [];
for await (const chunk of stream) {
  chunks.push(chunk);
}
const buffer = Buffer.concat(chunks);
```

---

### Question 324: What is the `finished()` utility in stream handling?

**Answer:**
`stream.finished(stream, callback)` functions as a safe way to detect when a stream is done (success or error). It handles cleanup logic.

---

### Question 325: Can a stream emit events after it has ended?

**Answer:**
Ideally no `data` events, but `close` or `error` might fire. Once `end` (Readable) or `finish` (Writable) fires, the data flow is done.

---

### Question 326: How do you share state between worker threads?

**Answer:**
1.  **`workerData`**: Pass initial data (Copy).
2.  **`SharedArrayBuffer`**: Share memory area. Use `Atomics` to manage race conditions.
3.  **`MessagePort`**: Send messages back and forth.

---

### Question 327: What is a message channel in worker_threads?

**Answer:**
`MessageChannel` creates two entangled ports (`port1`, `port2`). You can pass one port to a worker and keep the other to establish a direct communication line.

---

### Question 328: How do you handle graceful restart of a clustered Node.js app?

**Answer:**
(See Q211).
Master sends signal to Worker. Worker stops accepting connections, finishes existing ones, exits. Master spawns new Worker.

---

### Question 329: When would you prefer clustering over load balancing with NGINX?

**Answer:**
Cluster is for multi-core usage on a **single** machine. Nginx is for balancing across **multiple** machines (or ports).
Usually, you use **both**: Nginx -> balances to Machine A, B -> Machine A uses Cluster for its 8 cores.

---

### Question 330: How do you handle sticky sessions with clustering?

**Answer:**
If using Socket.IO, clients must connect to the *same* worker process.
This requires a **Sticky Store** (Redis) or IP-Hash balancing at the Nginx level, not just round-robin.

---

### Question 331: How do you create a recursive directory structure in Node.js?

**Answer:**
Use `mkdir` with `recursive: true`.

**Code:**
```javascript
fs.mkdirSync('./a/b/c', { recursive: true });
```

---

### Question 332: What’s the difference between `fs.exists()` and `fs.access()`?

**Answer:**
*   `fs.exists()`: Deprecated.
*   `fs.access()`: Checks permissions (Read/Write).

**Best Practice:** Do not check existence before opening. Just `open()` and handle the error. (Avoids Race Conditions).

---

### Question 333: How do you prevent race conditions when reading and writing files?

**Answer:**
If two requests read-mod-write a file, data is lost.
**Fix:** Use a file lock (libraries like `proper-lockfile`) or append-only logs.

---

### Question 334: How do you efficiently copy large files using streams?

**Answer:**
Don't read into memory. Use `copyFile` (fastest, uses kernel call) or `pipeline`.

**Code:**
```javascript
fs.copyFileSync('src', 'dest'); 
// or
stream.pipeline(read, write, cb);
```

---

### Question 335: What is the use of `fs.promises`?

**Answer:**
Provides Promise-based (async/await) versions of FS methods.

```javascript
const fs = require('fs/promises');
await fs.writeFile('file.txt', 'data');
```

---

### Question 336: What happens if you `await` a non-promise value?

**Answer:**
JS wraps it in a resolved Promise. It inserts a microtask delay.

```javascript
const val = await 42; // Same as Promise.resolve(42)
```

---

### Question 337: How does `async` function execution differ from regular functions?

**Answer:**
An `async` function **always** returns a Promise. Even if you `return 1`, it becomes `Promise.resolve(1)`. It allows `await` keyword.

---

### Question 338: How do you make a custom class thenable?

**Answer:**
Implement a `then` method. `await` will recognize it.

**Code:**
```javascript
class MyTask {
  then(resolve, reject) {
    resolve('Done');
  }
}

console.log(await new MyTask()); // "Done"
```

---

### Question 339: How do you cancel a Promise in Node.js?

**Answer:**
Promises are not cancellable by default.
Use **AbortController** (passed to the async operation).

---

### Question 340: Can you make `setTimeout` cancellable?

**Answer:**
Yes, `clearTimeout(id)`.
Or using AbortSignal in usage Node versions:
`setTimeout(1000, null, { signal })`.

---

### Question 341: How do you profile memory leaks in Node.js?

**Answer:**
1.  Isolate the leak (load test).
2.  Take **Heap Snapshots** (start vs end).
3.  Compare in Chrome DevTools. Look for objects retained (New Allocations).

---

### Question 342: What is the role of generational garbage collection in V8?

**Answer:**
Splits heap into **New Space** (Young) and **Old Space**.
Young objects are cheap to collect (Scavenge).
Old objects are expensive (Mark-Sweep).
Most objects die young, making GC efficient.

---

### Question 343: How can you limit memory usage per request?

**Answer:**
Node.js doesn't natively limit memory *per request* (only per process).
**Workaround:** Run code in a confined environment (Child Process, Worker, or VM) where you can measure/kill it.

---

### Question 344: What are weak references and when are they useful?

**Answer:**
`WeakMap` / `WeakSet`.
They hold references to objects **without** preventing Garbage Collection.
**Use:** Associating metadata with an object without causing a memory leak if the object is deleted elsewhere.

---

### Question 345: What is the memory cost of closures in long-lived processes?

**Answer:**
If a closure is stored in a global array/handler, it keeps the **entire Scope** (and parent scopes) alive. This is a common leak source.

---

### Question 346: How do you build dynamic route middleware?

**Answer:**
Return a middleware function based on config.

**Code:**
```javascript
const validate = (schema) => (req, res, next) => {
  if (schema(req.body)) next();
  else res.status(400).send();
};

app.post('/user', validate(userSchema), ...);
```

---

### Question 347: How do you attach data to `res.locals` in Express?

**Answer:**
`res.locals` is an object scoped to the request. It's useful for passing data between middleware and views.

```javascript
app.use((req, res, next) => {
  res.locals.user = req.user;
  next();
});
```

---

### Question 348: How do you skip to the next route in Express?

**Answer:**
Call `next('route')`. This skips remaining middleware in the *current* route stack and goes to the *next* route matching handling.

---

### Question 349: How do you write conditional middleware execution logic?

**Answer:**
Wrap it.

```javascript
const conditional = (condition, middleware) => (req, res, next) => {
  if (condition(req)) return middleware(req, res, next);
  next();
};
```

---

### Question 350: What happens if a middleware hangs and never calls `next()`?

**Answer:**
The request hangs indefinitely (until client timeouts). The browser just spins.
**Always** ensure every code path calls `next()` or `res.send()`.
