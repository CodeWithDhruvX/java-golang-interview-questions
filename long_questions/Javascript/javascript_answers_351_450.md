# JavaScript Interview Questions & Answers (351-450)

## 🔹 3. Language Internals & Deep Dive (Questions 351-380)

**Q351: How is the `this` context determined in `setTimeout` / `setInterval`?**
By default, the callback executed by `setTimeout` uses the `window` (or global) object as `this`, because it is called as a standalone function invocation by the timer mechanism, not as a method of an object.
**Fix:** Use arrow functions (which capture lexical `this`) or `.bind(this)`.

**Q352: How does automatic semicolon insertion (ASI) work?**
The JavaScript parser automatically inserts semicolons at the end of lines if it encounters a parsing error or a line terminator (enter) that looks like the end of a statement.
**Gotcha:**
```javascript
return
{
   a: 1
}
// Parsed as: return; { a: 1 } -> Returns undefined!
```

**Q353: What is tail call optimization (TCO), and why is it not widely supported?**
TCO allows a recursive function to reuse the current stack frame if the recursive call is the final action (tail position). This prevents stack overflow. Ideally part of ES6, but only Safari (JavaScriptCore) implements it reliably. V8 dropped it due to implementation complexity and debugging stack trace issues.

**Q354: How do you implement your own version of `new` keyword?**
```javascript
function myNew(Constructor, ...args) {
    const instance = Object.create(Constructor.prototype);
    const result = Constructor.apply(instance, args);
    return (typeof result === 'object' && result !== null) ? result : instance;
}
```

**Q355: What is the difference between shallow copy and deep copy?**
*   **Shallow:** Copies properties. If a property is an object, it copies the *reference*. (`{...obj}`, `Object.assign`).
*   **Deep:** Recursively copies all nested objects. (`structuredClone`, `JSON.parse(JSON.stringify)`).

**Q356: How do object property descriptors work?**
They describe attributes of a property: `value`, `writable`, `enumerable`, `configurable`. Access via `Object.getOwnPropertyDescriptor(obj, 'prop')`.

**Q357: What is an accessor property vs data property?**
*   **Data Property:** Has a `value` and `writable` attribute.
*   **Accessor Property:** Has `get` and `set` functions instead of a value. defined via `Object.defineProperty` or `get prop() {}` syntax.

**Q358: What does it mean that JavaScript is a single-threaded language?**
It has a single Call Stack. It can only execute one piece of code at a time. Concurrency is handled via the Event Loop and non-blocking I/O (Web APIs).

**Q359: How does JavaScript manage memory under the hood?**
*   **Allocation:** Memory reserved when objects/vars are created.
*   **Usage:** Read/Write.
*   **Deallocation (GC):** "Mark-and-Sweep" algorithm finds unreachable objects and frees memory.

**Q360: How does the V8 engine optimize JavaScript?**
*   **JIT Compilation:** Compiles JS to machine code at runtime (Ignition interpreter + TurboFan optimizer).
*   **Hidden Classes:** Optimizes object property access by creating hidden structures (shapes) shared by objects with same properties.
*   **Inline Caching:** Caches results of property lookups.

---

## 🔹 4. Closures & Scope Tricks (Questions 381-410)
*(Note: Source numbering labeled this section 381-410, but questions are 361-370)*

**Q361: How do closures help in data privacy?**
They allow creating "private" variables that are inaccessible from the outside but accessible to privileged methods returned by the closure.
```javascript
function createBank() {
    let balance = 0; // Private
    return { deposit: (n) => balance += n };
}
```

**Q362: Can closures lead to memory leaks? How?**
Yes. If a closure is stored in a long-lived object (like a global handler) and it references large objects in its lexical scope, those large objects cannot be garbage collected even if not actively used.

**Q363: How can you use closures to implement private variables?**
(See Q361). By defining variables inside a function and returning only the functions that need to access them.

**Q364: What will be the output of a closure inside a loop with `var`?**
`var` is function-scoped, so the variable is shared across iterations. By the time the closure runs, the loop has finished, and the variable holds the final value.
```javascript
for (var i=0; i<3; i++) setTimeout(() => console.log(i), 100); 
// Output: 3, 3, 3
```

**Q365: How does JavaScript handle block-level closures with `let`?**
`let` creates a new execution context (binding) for *each iteration* of the loop. Each closure captures the specific instance of the variable for that iteration.
```javascript
for (let i=0; i<3; i++) setTimeout(() => console.log(i), 100);
// Output: 0, 1, 2
```

**Q366: What is a closure trap and how to fix it?**
Accidentally capturing an old variable (stale closure) typically in React hooks or loops. Fix: Use Functional State Updates, `useRef`, or dependency arrays correctly.

**Q367: Can closures capture updated variable values?**
Yes. Closures capture the *reference* to the variable, not the value at creation time. If the variable changes, the closure sees the new value (unless the closure ran before the change).

**Q368: How can you simulate a module using closures?**
**Revealing Module Pattern**:
```javascript
const Module = (() => {
    const p = "private";
    return {
        get: () => p
    };
})();
```

**Q369: Explain scope chaining with a nested closure.**
Inner Scope -> Outer Function Scope -> Global Scope. The engine looks up the chain until it finds the variable.

**Q370: How does garbage collection affect closed-over variables?**
Normally local vars are GC'd when function ends. If a closure exists and is reachable (e.g., assigned to a global), the variables in its scope are kept alive in the Heap.

---

## 🔹 5. ES2021–ES2024+ Concepts (Questions 411-440)
*(Source questions 371-380)*

**Q371: What are logical assignment operators and when to use them?**
`&&=`, `||=`, `??=`.
Use them for concise default assignment or short-circuit updates. `config.timeout ??= 3000;`.

**Q372: What are class static initialization blocks?**
`static { ... }` block inside a class. useful for evaluating statements (like try-catch) to initialize private static fields.

**Q373: What are top-level await pitfalls?**
It blocks the execution of the entire module (and modules importing it) until the promise resolves. Can delay application startup if not used carefully (e.g. awaiting slow network calls).

**Q374: How does `at()` method improve indexing?**
Allows negative indexing for Arrays and Strings. `arr.at(-1)` vs `arr[arr.length - 1]`.

**Q375: What are the benefits of `Array.prototype.groupBy()`?**
(Now `Object.groupBy`). Transforms an array into an object where items are grouped by a specific key. Replaces manual `reduce` logic.

**Q376: What is a pipeline operator (`|>`) and its status?**
Standardizes chaining functions: `x |> f |> g` is `g(f(x))`. Still a Proposal (Stage 2/3), not standard yet.

**Q377: How does the `do` expression proposal work?**
Allows executing a block of statements and returning a completion value. `const x = do { if(cond) 1; else 2; }`. (Stage 1).

**Q378: What are Records and Tuples (Stage 2 proposal)?**
Immutable primitive versions of Objects and Arrays. Compares by value, not reference. `#{ a: 1 } === #{ a: 1 }` is true.

**Q379: What is the shadow realm proposal in JavaScript?**
Allows creating a distinct global environment (like an iframe or Worker) but synchronously within the same JS thread. Useful for sandboxing plugins.

**Q380: What’s the difference between `import.meta` and `import()`?**
*   `import.meta`: Object containing metadata about the current module (e.g., `import.meta.url`).
*   `import()`: Function to dynamically load modules.

---

## 🔹 6. Functions, Context, and Binding (Questions 441-470)
*(Source questions 381-390)*

**Q381: How does function hoisting differ from variable hoisting?**
Function declarations are hoisted *with* their definition (usable instantly). `var` variables are hoisted but initialized to `undefined`.

**Q382: Why are arrow functions not suitable as constructors?**
They do not have a `prototype` property and strictly bind `this` lexically. Calling them with `new` throws `TypeError`.

**Q383: Can you override `this` in arrow functions?**
No. `call`, `apply`, and `bind` can pass arguments but cannot change the value of `this` in an arrow function.

**Q384: How do fat-arrow functions differ in scope handling?**
They do not create their own execution context for `this`. They use the `this` from the surrounding code block.

**Q385: Can you write a function that returns itself?**
Yes. Useful for chaining or recursion.
```javascript
function foo() { return foo; }
foo()()() === foo; // true
```

**Q386: How to implement a function spy/logger?**
Wrap the target function.
```javascript
function spy(fn) {
    return function(...args) {
        console.log("Called with", args);
        return fn.apply(this, args);
    }
}
```

**Q387: How to make a function self-memoizing?**
Attach the cache to the function object itself.
```javascript
function fib(n) {
    if (!fib.cache) fib.cache = {};
    if (fib.cache[n]) return fib.cache[n];
    // compute...
    return fib.cache[n] = result;
}
```

**Q388: What’s a trampoline function in recursion?**
A utility to prevent stack overflow. It wraps a recursive function so that it returns a "thunk" (function) instead of calling itself immediately. The trampoline loop executes thunks iteratively.

**Q389: What are generator functions and how do they work internally?**
Defined with `function*`. They return a Generator object (Iterator). Maintaining internal state via a state machine logic, yielding values and pausing execution context until `next()` is called.

**Q390: What is the difference between generator and async generator?**
*   **Generator:** `yield` returns value synchronously.
*   **Async Generator:** `yield` (or `yield await`) returns a Promise. Iterated via `for await...of`.

---

## 🔹 7. Advanced Async & Concurrency (Questions 471-500)
*(Source questions 391-400)*

**Q391: How does `Promise.prototype.finally()` work?**
Executes a callback when the promise settles (fulfilled or rejected). It passes through the previous result/error to the next handler.

**Q392: How do you cancel a fetch request?**
Using `AbortController`.
```javascript
const controller = new AbortController();
fetch(url, { signal: controller.signal });
controller.abort(); // Rejects fetch with AbortError
```

**Q393: What is the AbortController API?**
A generic API to signal cancellation to async tasks (Fetch, Event Listeners).

**Q394: What is a scheduler in JavaScript?**
Often refers to the **Event Loop** or the new `scheduler.postTask()` API (experimental) for prioritizing tasks.

**Q395: How does cooperative multitasking work in JS?**
Tasks yield control back to the main thread (e.g. via `setTimeout(0)` or `postTask`) to allow UI rendering, preventing freezing. Not pre-emptive.

**Q396: How does concurrency differ from parallelism in JS?**
*   **Concurrency:** Handling multiple tasks in overlapping timeframes (Event Loop switching).
*   **Parallelism:** Running tasks simultaneously (Web Workers, multi-core).

**Q397: What are microtasks vs macrotasks with examples?**
*   **Microtask:** `Promise.then`, `queueMicrotask`. (Run ASAP).
*   **Macrotask:** `setTimeout`, I/O. (Run next loop).

**Q398: How does event delegation help async UIs?**
You attach one listener to a container. When async content loads (new child elements), the listener automatically handles their events via bubbling, without re-attaching listeners.

**Q399: How do you chain multiple fetch requests properly?**
Return the Promise from `.then()`.
```javascript
fetch(url1)
  .then(res => res.json())
  .then(data => fetch(url2 + data.id))
  .then(res => res.json());
```

**Q400: What happens when you `await` a non-promise?**
JavaScript wraps it in `Promise.resolve()`. It still suspends execution of the async function until the next microtask tick.
```javascript
await 42; // Treated as Promise.resolve(42)
```

---

## 🔹 8. Regex, Parsing, String Ops (Questions 501-530)
*(Source questions 401-410)*

**Q401: How do you escape special characters in a RegExp?**
Prepend with `\`. E.g., `\.`, `\*`. To functionize: `str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')`.

**Q402: What is lazy vs greedy matching in regex?**
*   **Greedy (Default):** Matches as much as possible `.*`.
*   **Lazy (Reluctant):** Matches as little as possible `.*?`.

**Q403: How to create a named capture group?**
Syntax `(?<name>...)`. Access via `matchResult.groups.name`.
```javascript
const re = /(?<year>\d{4})/;
const match = re.exec("2026");
console.log(match.groups.year); // "2026"
```

**Q404: What are the use cases for `String.prototype.matchAll()`?**
Iterating over all matches of a global regex, including access to capture groups for each match. Returns an iterator.

**Q405: What is the difference between `match()` and `exec()`?**
*   `str.match(re)`: Returns array of string matches (no groups if /g).
*   `re.exec(str)`: Stateful loop. Returns detailed match object one by one.

**Q406: How do you validate a strong password with regex?**
Use lookaheads.
`/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&]).{8,}/`

**Q407: What is the difference between `search()` and `indexOf()`?**
*   `search()`: Takes a RegExp. Returns index.
*   `indexOf()`: Takes a String. Returns index. Faster for simple substrings.

**Q408: How does `padStart()` help with formatting?**
Pads start of string to target length.
`"5".padStart(2, "0")` -> `"05"`. Great for dates/IDs.

**Q409: What’s the difference between `slice()`, `substring()` and `substr()`?**
*   `slice(start, end)`: Supports negative indices. (Recommended).
*   `substring(start, end)`: Swaps start/end if start > end. No negative support.
*   `substr(start, length)`: **Deprecated**.

**Q410: How do tagged template literals work?**
(Duplicate of Q98). Function receives parts of string and expressions.

---

## 🔹 9. Data Handling, Encoding, JSON (Questions 531-560)
*(Source questions 411-420)*

**Q411: How to safely stringify an object with circular references?**
Use a custom `replacer` function or a library (`flatted`).
`JSON.stringify(obj, (key, value) => ...)` (Hard to do manually).

**Q412: How does `structuredClone()` differ from `JSON.stringify()`?**
`structuredClone` supports circular references, `Date`, `Map`, `Set`, `BigInt`, `RegExp`. `JSON` does not.

**Q413: What is base64 encoding and how to do it in JS?**
Encodes binary data to text.
*   Browser: `btoa()` (Encode), `atob()` (Decode).
*   Node: `Buffer.from(str).toString('base64')`.

**Q414: How do you parse and validate nested JSON?**
Use `try-catch` around `JSON.parse()`. Manually check properties or use a schema validator (Zod, Ajv).

**Q415: How to handle large JSON file parsing efficiently?**
Use streaming parsers (like SAX/Oboe.js) instead of loading full string into memory.

**Q416: What’s the difference between JSONP and CORS?**
*   **CORS:** Standard modern security mechanism. Server sends headers.
*   **JSONP:** Legacy hack using `<script>` tags to bypass Same-Origin Policy. Insecure.

**Q417: What is the `reviver` function in `JSON.parse()`?**
Second argument. Transforms the result.
`JSON.parse(json, (key, val) => key === 'date' ? new Date(val) : val)`.

**Q418: What’s the difference between parsing a JSON string vs object?**
You can only parse a String. If you pass an object to `JSON.parse`, it coerces to string (often `"[object Object]"` -> fails).

**Q419: How do you serialize data with Dates or BigInt?**
Use `JSON.stringify` with a `replacer`.
```javascript
JSON.stringify({ n: 10n }, (k, v) => typeof v === 'bigint' ? v.toString() + 'n' : v);
```

**Q420: How do you send FormData with fetch?**
```javascript
const fd = new FormData();
fd.append('file', fileInput.files[0]);
fetch('/upload', { method: 'POST', body: fd }); 
// Note: Do NOT set Content-Type header manually (browser sets boundary).
```

---

## 🔹 10. Memory, Storage, Workers (Questions 561-590)
*(Source questions 421-430)*

**Q421: How does the JS event loop interact with Web Workers?**
Workers run in a separate thread with their own Event Loop. They communicate with the main thread loop via messages (`postMessage`).

**Q422: What is SharedArrayBuffer and Atomics?**
*   `SharedArrayBuffer`: Memory shared between Worker and Main thread (no copying).
*   `Atomics`: Operations to ensure thread-safe reading/writing to that shared memory.

**Q423: What are transferable objects in Web Workers?**
Objects (ArrayBuffer, MessagePort) whose ownership is transferred to the worker. Zero-copy (instant), but main thread loses access.

**Q424: What is the IndexedDB API?**
Low-level API for client-side storage of significant amounts of structured data (files/blobs). Transactional and Async.

**Q425: How do you handle blob URLs and object URLs?**
`URL.createObjectURL(blob)`. Creates a pseudo-URL referencing data in memory. **Must** revoke with `URL.revokeObjectURL()` to free memory.

**Q426: What’s the lifetime of data in sessionStorage?**
Until the tab/window is closed (Page Session). Survives reloads.

**Q427: How to store and retrieve objects in localStorage?**
Convert to string: `setItem('key', JSON.stringify(obj))`.
Retrieve: `JSON.parse(getItem('key'))`.

**Q428: What are Service Workers and how do they cache?**
Proxy scripts. They use the **Cache API** (`caches.open`, `put`, `match`) to store network responses programmatically for offline use.

**Q429: What is the Cache API?**
Programmatic storage for Request/Response pairs. Used by Service Workers.

**Q430: How do you monitor memory usage in JS?**
*   `performance.memory` (Chrome specific).
*   DevTools Memory Profiler (Snapshot).
*   `memwatch` (Node.js).

---

## 🔹 11. Security & Edge Cases (Questions 591-620)
*(Source questions 431-440)*

**Q431: What is prototype pollution?**
(Duplicate Q143). Merging unsafe objects into `__proto__`.

**Q432: What is script injection and how to sanitize input?**
Injecting executable JS code. Sanitize using libraries like DOMPurify before rendering HTML.

**Q433: What is DOM-based XSS?**
Vulnerability where the attack payload is executed as a result of modifying the DOM "on client-side" (e.g., retrieving param from URI and writing to innerHTML without validation).

**Q434: What is CSP (Content Security Policy)?**
HTTP header that tells the browser which sources of executable scripts are approved (e.g., "only my domain"). Prevents XSS.

**Q435: How to secure client-side tokens?**
*   HttpOnly Cookies (Best, prevents JS access).
*   Memory (cleared on refresh).
*   Avoid localStorage for sensitive tokens (accessible to XSS).

**Q436: What is clickjacking and how can JS prevent it?**
Invisible iframe overlay tricking users.
Prevent via Header: `X-Frame-Options: DENY`.
JS: Frame busting `if (top != window) top.location = window.location`.

**Q437: How does `innerHTML` open you to security risks?**
It parses strings as HTML. If string contains `<img src=x onerror=alert(1)>`, script runs.

**Q438: Best practices for handling user input?**
*   Validation (Type check).
*   Sanitization (Strip tags).
*   Encoding (Render as text, not HTML).

**Q439: `eval()` vs `Function()` security?**
*   `eval()`: Access local scope (Very bad).
*   `Function()`: Access global scope only (Still bad, allows arbitrary code execution).

**Q440: Why is it unsafe to use user input directly in template literals?**
If the literal is generating HTML/SQL code, injecting input directly facilitates Injection attacks. Use Tagged Templates to sanitize values.

---

## 🔹 12. Math, Date, Utility (Questions 621-650)
*(Source questions 441-450)*

**Q441: How do you generate a UUID in JavaScript?**
`crypto.randomUUID()` (Modern). Or legacy Math.random hacks.

**Q442: Difference between `Math.floor()` and `Math.trunc()`?**
*   `floor`: Rounds down (3.9 -> 3, -3.9 -> -4).
*   `trunc`: Removes decimal (3.9 -> 3, -3.9 -> -3).

**Q443: How to implement a random number in a range?**
`Math.random() * (max - min) + min`.

**Q444: How to find time difference between two ISO date strings?**
`new Date(str2) - new Date(str1)` (Result in ms).

**Q445: Best way to clone a date object?**
`new Date(oldDate.getTime())` or `new Date(oldDate)`.

**Q446: What is `Intl.DateTimeFormat` used for?**
Formatting dates for specific locales/languages.

**Q447: How to sort an array of objects by date?**
`arr.sort((a, b) => new Date(a.date) - new Date(b.date))`.

**Q448: How to throttle a function using `requestAnimationFrame()`?**
Ensures function runs at most once per frame (visual updates).
```javascript
let scheduled = false;
function onScroll() {
    if (scheduled) return;
    scheduled = true;
    requestAnimationFrame(() => {
        // do work
        scheduled = false;
    });
}
```

**Q449: How does `setTimeout(..., 0)` really behave?**
Pushes callback to end of Task Queue. Runs after stack clears and microtasks. Minimum delay is ~4ms (browser limit for nested timeouts).

**Q450: How to pause execution for 5 seconds in an async function?**
`await new Promise(r => setTimeout(r, 5000));`.
