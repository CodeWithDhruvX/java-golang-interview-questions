# ⚡ 02 — Async JavaScript & Event Loop
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Event Loop internals (call stack, microtask queue, macrotask queue)
- Promise internals and chaining
- `async/await` error handling patterns
- `Promise.all`, `Promise.allSettled`, `Promise.race`, `Promise.any`
- Generators and iterators
- Observable pattern
- Concurrency patterns

---

## ❓ Most Asked Questions

### Q1. Explain the Event Loop in detail — output prediction.

```javascript
// Event Loop Priority:
// 1. Call Stack (synchronous)
// 2. Microtasks: Promise callbacks, queueMicrotask, MutationObserver
// 3. Macrotasks: setTimeout, setInterval, setImmediate, I/O, UI events

console.log("1");                           // Call stack

setTimeout(() => console.log("2"), 0);     // Macrotask queue

Promise.resolve()
    .then(() => console.log("3"))          // Microtask
    .then(() => console.log("4"));         // Microtask (chained)

queueMicrotask(() => console.log("5"));    // Microtask

console.log("6");                           // Call stack

// Output: 1, 6, 3, 5, 4, 2
// Explanation:
// Sync: 1, 6
// Microtasks drained completely: 3 → schedules 4 → 5 → 4
// Macrotask: 2

// ⚠️ Nested setTimeout vs Promise.resolve in a loop:
setTimeout(() => {
    console.log("timeout");
    Promise.resolve().then(() => console.log("microtask in timeout"));
}, 0);
// Output: "timeout", "microtask in timeout"
// Microtasks run immediately after each macrotask!
```

---

### Q2. Implement `Promise.all`, `Promise.allSettled`, `Promise.race`, `Promise.any` from scratch.

```javascript
// Promise.all — fails fast on first rejection
function promiseAll(promises) {
    return new Promise((resolve, reject) => {
        if (!promises.length) return resolve([]);
        
        const results = [];
        let pending = promises.length;

        promises.forEach((promise, i) => {
            Promise.resolve(promise).then(value => {
                results[i] = value;
                if (--pending === 0) resolve(results);
            }).catch(reject); // short-circuit on first error
        });
    });
}

// Promise.allSettled — waits for all, never rejects
function promiseAllSettled(promises) {
    return Promise.all(promises.map(p =>
        Promise.resolve(p)
            .then(value  => ({ status: "fulfilled", value }))
            .catch(reason => ({ status: "rejected", reason }))
    ));
}

// Promise.race — first to settle (resolve OR reject) wins
function promiseRace(promises) {
    return new Promise((resolve, reject) => {
        promises.forEach(p => Promise.resolve(p).then(resolve).catch(reject));
    });
}

// Promise.any — first to resolve wins; rejects only if ALL reject
function promiseAny(promises) {
    return new Promise((resolve, reject) => {
        let errors = [];
        let pending = promises.length;

        if (!pending) return reject(new AggregateError([], "All promises were rejected"));

        promises.forEach((p, i) => {
            Promise.resolve(p)
                .then(resolve) // first resolve wins
                .catch(err => {
                    errors[i] = err;
                    if (--pending === 0) {
                        reject(new AggregateError(errors, "All promises were rejected"));
                    }
                });
        });
    });
}
```

---

### Q3. What is the difference between microtask and macrotask queues?

```javascript
// Microtask Queue (higher priority):
// - Promise .then/.catch/.finally callbacks
// - queueMicrotask()
// - MutationObserver callbacks
// ALL microtasks drain before next macrotask runs!

// Macrotask Queue (lower priority):
// - setTimeout / setInterval
// - setImmediate (Node.js)
// - I/O callbacks
// - UI rendering events
// - MessageChannel

// Practical order in browser:
// 1. Current script runs to completion (call stack cleared)
// 2. ALL microtasks run (including new microtasks added during drain)
// 3. Browser may render (between macrotasks)
// 4. ONE macrotask runs
// 5. Go to step 2

// Starvation: infinite microtask loop blocks macrotasks and rendering!
function starve() {
    queueMicrotask(starve); // never lets setTimeout run!
}

// Node.js specific queue order:
// nextTick queue > Promise microtask > macrotask (I/O) > setImmediate
process.nextTick(() => console.log("nextTick")); // first!
Promise.resolve().then(() => console.log("promise"));
setImmediate(() => console.log("setImmediate")); // last
```

---

### Q4. Implement a `retry` utility for async functions.

```javascript
// Exponential backoff retry — production pattern
async function retry(fn, {
    retries = 3,
    delay = 1000,
    backoff = 2,
    onRetry = null
} = {}) {
    let lastError;
    let currentDelay = delay;

    for (let attempt = 1; attempt <= retries + 1; attempt++) {
        try {
            return await fn();
        } catch (error) {
            lastError = error;
            if (attempt > retries) break;

            if (onRetry) onRetry(error, attempt);

            // Exponential backoff with jitter
            const jitter = Math.random() * 0.3 * currentDelay;
            await sleep(currentDelay + jitter);
            currentDelay *= backoff;
        }
    }
    throw lastError;
}

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

// Usage
const data = await retry(
    () => fetch("https://api.example.com/data").then(r => r.json()),
    {
        retries: 3,
        delay: 500,
        backoff: 2,
        onRetry: (err, attempt) => console.log(`Retry ${attempt}: ${err.message}`)
    }
);

// With timeout
async function withTimeout(fn, timeoutMs) {
    const timeout = new Promise((_, reject) =>
        setTimeout(() => reject(new Error("Timeout")), timeoutMs)
    );
    return Promise.race([fn(), timeout]);
}
```

---

### Q5. Explain Generators and their use in async control flow.

```javascript
// Generator: function that can pause (yield) and resume
function* counter() {
    let i = 0;
    while (true) {
        const reset = yield i++;   // yield returns value, receives next() arg
        if (reset) i = 0;
    }
}

const gen = counter();
gen.next();        // { value: 0, done: false }
gen.next();        // { value: 1, done: false }
gen.next(true);    // { value: 0, done: false } — reset!

// ✅ Real use: Pagination with generator
async function* fetchPaginatedData(endpoint) {
    let page = 1;
    while (true) {
        const response = await fetch(`${endpoint}?page=${page}`);
        const { data, hasMore } = await response.json();
        yield data;
        if (!hasMore) break;
        page++;
    }
}

// Consuming async generator
for await (const batch of fetchPaginatedData("/api/users")) {
    processBatch(batch);
}

// ✅ Infinite sequence without memory issues
function* fibonacci() {
    let [a, b] = [0, 1];
    while (true) {
        yield a;
        [a, b] = [b, a + b];
    }
}

const fib = fibonacci();
const first10 = Array.from({ length: 10 }, () => fib.next().value);
// [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
```

---

### Q6. What are common `async/await` anti-patterns?

```javascript
// ❌ Anti-pattern 1: Sequential awaits when parallel is possible
async function bad() {
    const user = await fetchUser(1);      // waits ~100ms
    const posts = await fetchPosts(1);    // waits ANOTHER ~100ms
    // Total: ~200ms
}

// ✅ Fix: parallel with Promise.all
async function good() {
    const [user, posts] = await Promise.all([
        fetchUser(1),    // both start simultaneously
        fetchPosts(1)    // Total: ~100ms
    ]);
}

// ❌ Anti-pattern 2: Swallowed errors (no try/catch)
async function swallowed() {
    const data = await riskyOperation(); // if this throws, unhandled rejection!
    return data;
}

// ✅ Fix: proper error handling
async function handled() {
    try {
        return await riskyOperation();
    } catch (err) {
        logger.error(err);
        throw new AppError("Operation failed", { cause: err });
    }
}

// ❌ Anti-pattern 3: async in forEach
async function bad2(ids) {
    ids.forEach(async id => {
        await processId(id); // forEach doesn't await these!
    });
    // Function returns before any processing is done
}

// ✅ Fix: use for...of or Promise.all with map
async function good2(ids) {
    await Promise.all(ids.map(id => processId(id))); // parallel
    // OR
    for (const id of ids) {
        await processId(id); // sequential if needed
    }
}

// ❌ Anti-pattern 4: new Promise wrapping async functions (Promise hell)
function bad3() {
    return new Promise(async (resolve, reject) => { // ❌ antipattern
        const data = await fetch("url"); // async inside new Promise constructor
        resolve(data);
    });
}

// ✅ Fix: just use async function directly
async function good3() {
    return fetch("url"); // Promise is returned automatically
}
```

---

### Q7. Implement a `TaskQueue` with concurrency limits.

```javascript
// Concurrency limiter: run at most N async tasks simultaneously
class TaskQueue {
    #concurrency;
    #running = 0;
    #queue = [];

    constructor(concurrency = 4) {
        this.#concurrency = concurrency;
    }

    async add(fn) {
        return new Promise((resolve, reject) => {
            this.#queue.push({ fn, resolve, reject });
            this.#run();
        });
    }

    #run() {
        while (this.#running < this.#concurrency && this.#queue.length) {
            const { fn, resolve, reject } = this.#queue.shift();
            this.#running++;

            Promise.resolve(fn())
                .then(resolve)
                .catch(reject)
                .finally(() => {
                    this.#running--;
                    this.#run(); // process next in queue
                });
        }
    }
}

// Usage: download 100 files but max 5 at a time
const queue = new TaskQueue(5);
const urls = Array.from({ length: 100 }, (_, i) => `https://api.com/file/${i}`);

const results = await Promise.all(
    urls.map(url => queue.add(() => fetch(url).then(r => r.blob())))
);
```

---

### Q8. Explain `AbortController` and cancellable fetch.

```javascript
// AbortController: cancel async operations (fetch, EventListeners)

// Basic fetch cancellation
async function fetchWithCancel(url) {
    const controller = new AbortController();
    const { signal } = controller;

    // Cancel after 5 seconds
    const timeout = setTimeout(() => controller.abort(), 5000);

    try {
        const response = await fetch(url, { signal });
        clearTimeout(timeout);
        return await response.json();
    } catch (err) {
        if (err.name === 'AbortError') {
            console.log("Fetch cancelled");
            return null;
        }
        throw err;
    }
}

// React/SPA: cancel previous search on new keystroke
function useSearch(query) {
    useEffect(() => {
        const controller = new AbortController();

        async function search() {
            try {
                const res = await fetch(`/api/search?q=${query}`, {
                    signal: controller.signal
                });
                setResults(await res.json());
            } catch (err) {
                if (err.name !== 'AbortError') setError(err);
            }
        }

        search();

        return () => controller.abort(); // cleanup: abort on unmount or query change
    }, [query]);
}

// Chaining controllers: abort parent cancels all children
const parentController = new AbortController();
// All child requests use parent's signal
const childFetches = urls.map(url =>
    fetch(url, { signal: parentController.signal })
);
parentController.abort(); // cancels ALL child fetches
```

---

### Q9. Implement an event emitter (Observable pattern).

```javascript
// EventEmitter from scratch — common interview question
class EventEmitter {
    #events = new Map();

    on(event, listener) {
        if (!this.#events.has(event)) {
            this.#events.set(event, new Set());
        }
        this.#events.get(event).add(listener);
        return this; // chainable
    }

    once(event, listener) {
        const wrapper = (...args) => {
            listener(...args);
            this.off(event, wrapper);
        };
        wrapper._original = listener; // for removal tracking
        return this.on(event, wrapper);
    }

    emit(event, ...args) {
        if (!this.#events.has(event)) return false;
        this.#events.get(event).forEach(listener => {
            try { listener(...args); }
            catch (err) { console.error(`Listener error on ${event}:`, err); }
        });
        return true;
    }

    off(event, listener) {
        if (!this.#events.has(event)) return this;
        const listeners = this.#events.get(event);
        listeners.forEach(l => {
            if (l === listener || l._original === listener) {
                listeners.delete(l);
            }
        });
        return this;
    }

    removeAllListeners(event) {
        if (event) this.#events.delete(event);
        else this.#events.clear();
        return this;
    }
}

// Usage
const bus = new EventEmitter();
bus.on("data", data => console.log("Received:", data));
bus.once("connect", () => console.log("Connected!")); // fires only once
bus.emit("data", { id: 1, name: "Alice" });
bus.emit("connect"); // "Connected!" — listener auto-removed
bus.emit("connect"); // nothing
```

---

### Q10. What are Web Workers and how do they relate to the Event Loop?

```javascript
// Main thread: JS is single-threaded — long CPU tasks block the event loop
// Web Workers: separate thread — runs JS in parallel without blocking UI

// main.js
const worker = new Worker('worker.js');

// Send data to worker (structured clone — no shared memory by default)
worker.postMessage({ type: "COMPUTE", payload: { n: 40 } });

// Receive results asynchronously
worker.onmessage = ({ data }) => {
    console.log("Fibonacci result:", data.result);
};

worker.onerror = err => console.error("Worker error:", err);

// Terminate when done
// worker.terminate();

// worker.js (separate file, separate thread)
self.onmessage = ({ data }) => {
    if (data.type === "COMPUTE") {
        const result = fib(data.payload.n); // CPU-intensive — won't block main thread
        self.postMessage({ result });
    }
};

function fib(n) {
    if (n <= 1) return n;
    return fib(n - 1) + fib(n - 2);
}

// SharedArrayBuffer: ACTUAL shared memory between threads (requires COOP/COEP headers)
const sharedBuffer = new SharedArrayBuffer(4);
const sharedArray = new Int32Array(sharedBuffer);
// Use Atomics for safe concurrent access
Atomics.store(sharedArray, 0, 42);
const value = Atomics.load(sharedArray, 0); // 42
```
