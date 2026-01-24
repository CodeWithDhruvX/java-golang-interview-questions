# JavaScript Interview Questions & Answers (751-850)

## 🔹 19. Symbols, Meta & Reflection (Questions 751-780)

**Q751: What are the use cases of `Symbol.toPrimitive`?**
It allows an object to customize its conversion to a primitive value. It gets called with a hint (`"number"`, `"string"`, or `"default"`).
```javascript
const obj = {
  [Symbol.toPrimitive](hint) {
    return hint === "number" ? 10 : "text";
  }
};
console.log(+obj); // 10
console.log(`${obj}`); // "text"
```

**Q752: How does `Symbol.hasInstance` customize `instanceof`?**
It lets a class/object decide if another object is an instance of it.
```javascript
class Even {
  static [Symbol.hasInstance](obj) {
    return Number(obj) % 2 === 0;
  }
}
console.log(2 instanceof Even); // true
```

**Q753: What is `Symbol.isConcatSpreadable`?**
A boolean property that controls if an object (like an Array-like object) should be flattened when using `Array.prototype.concat`.

**Q754: What does `Symbol.match` override?**
It identifies a regex-like object that can be updated for use with `String.prototype.match()`.

**Q755: What’s the purpose of `Symbol.iterator` vs `Symbol.asyncIterator`?**
*   `Symbol.iterator`: Used by `for...of`. Returns `{ value, done }`.
*   `Symbol.asyncIterator`: Used by `for await...of`. Returns `Promise<{ value, done }>`.

**Q756: How to use `Symbol.toStringTag` in objects?**
It customizes the output of `Object.prototype.toString.call(obj)`.
```javascript
const user = { [Symbol.toStringTag]: "User" };
console.log(Object.prototype.toString.call(user)); // "[object User]"
```

**Q757: What does `Symbol.unscopables` do?**
It specifies object properties to be excluded from the `with` environment bindings. (Rarely used).

**Q758: What’s the purpose of `Reflect.get()`?**
It gets a property from an object, like `obj[prop]`, but as a function. It specifically handles getters better when using a `receiver` (this context).

**Q759: What is the difference between `Reflect.set()` and direct assignment?**
`Reflect.set()` returns a **boolean** (`true` if successful, `false` if failed), whereas assignment (`=`) returns the assigned value (or throws in strict mode if read-only).

**Q760: Why use `Reflect.ownKeys()` over `Object.keys()`?**
`Object.keys()` returns enumerable string properties. `Reflect.ownKeys()` returns **all** own keys, including non-enumerable and Symbols.

---

## 🔹 20. Proxy Deep Dive (Questions 781-810)
*(Source list 761-770 in Header 781-810)*

**Q761: What are traps in JavaScript Proxy?**
Traps are internal methods you can intercept in a Proxy handler. Examples: `get`, `set`, `has`, `deleteProperty`, `construct`.

**Q762: How does `get` and `set` trap work?**
*   `get(target, prop, receiver)`: Intercepts reading properties.
*   `set(target, prop, value, receiver)`: Intercepts writing properties.

**Q763: What’s the `has` trap and when is it triggered?**
It intercepts the `in` operator.
```javascript
const p = new Proxy({}, {
  has(target, key) { return key === "secret"; }
});
console.log("secret" in p); // true
console.log("other" in p); // false
```

**Q764: How can Proxy be used to validate object values?**
By using the `set` trap to check the value type or range before assigning it to the target. If invalid, throw error or return false.

**Q765: How to observe property access using Proxy?**
Use the `get` trap to log specific access patterns.

**Q766: How to use Proxy for auto-completion or fallback values?**
In `get`, if the property doesn't exist on the target, return a default value or calculate similarity match.

**Q767: How can you use Proxy for access logging?**
Wrap an object and log every `get`/`set` call with timestamps.

**Q768: How to prevent new properties from being added with a Proxy?**
Use `preventExtensions` or `set` trap returning `false` for new keys.

**Q769: What is the role of `defineProperty` trap in Proxy?**
Intercepts `Object.defineProperty()`. Can be used to prevent changing property descriptors (e.g., making something read-only dynamically).

**Q770: Can you use a Proxy as a wrapper for functions?**
Yes. Use the `apply` trap to intercept function calls and `construct` trap for `new` calls.

---

## 🔹 21. Memory, GC, and Performance Internals (Questions 811-840)
*(Source list 771-780)*

**Q771: What is the mark-and-sweep garbage collection strategy?**
(See Q74). Roots -> Mark reachable -> Sweep/Free unreachable.

**Q772: What causes memory leaks in closures?**
If a closure is stored globally but references large variables from its parent scope which are no longer needed, the GC cannot clean them up.

**Q773: How to identify detached DOM nodes?**
Use DevTools "Heap Snapshot". Look for nodes marked "Detached". This happens when JS holds a reference to a node that was removed from the DOM tree.

**Q774: How do WeakRefs help with caching?**
`WeakRef` holds an object without preventing GC. If memory is tight, the GC collects the object. The cache logic checks `ref.deref()`: if undefined, re-fetch; else use cached.

**Q775: Why are WeakMaps better for DOM-based caches?**
If the DOM node (key) is removed from the document and has no other references, the WeakMap entry is automatically eventually GC'd, preventing leaks.

**Q776: What are finalizers in JavaScript and how are they used?**
Using `FinalizationRegistry`, you can register a callback to run after an object is GC'd. Useful for cleaning up external resources (WebAssembly memory, texture buffers).

**Q777: How can you prevent memory bloat in single-page applications?**
Clean up event listeners (`removeEventListener`), clear intervals, and nullify references in component teardown/unmount lifecycle.

**Q778: What does Chrome’s performance tab measure in JS?**
Execution time (Flame Chart), Layout/Reflow costs, Paint costs, Network waterfalls, Memory usage over time.

**Q779: How does JS engine handle inline caching?**
When a function accesses a property `obj.x` repeatedly with objects of the same structure (Hidden Class), V8 caches the memory offset of `x`, skipping the hash lookup.

**Q780: What is hidden class optimization in V8?**
V8 creates internal "shapes" (Hidden Classes) for objects. If two objects share the same properties added in the same order, they share a Hidden Class, speeding up property access.

---

## 🔹 22. Web Workers & Shared Memory (Questions 841-870)
*(Source list 781-790)*

**Q781: What is the purpose of a Web Worker?**
To run JavaScript code in a background thread, preventing CPU-intensive tasks (image processing, data sorting) from blocking the main thread (UI).

**Q782: What are the communication methods between workers and main thread?**
*   `postMessage()` / `onmessage` (Async, copying data).
*   `SharedArrayBuffer` (Concurrent access).

**Q783: Can you pass functions into Web Workers?**
No. Data passed via `postMessage` must be serializable (Structured Clone Algorithm). Functions cannot be cloned.

**Q784: How to terminate a Web Worker from the main thread?**
`worker.terminate()`. (Immediately kills it).

**Q785: What is a Blob Worker?**
A Worker created from a Blob URL (containing JS code string) instead of an external file.
`new Worker(URL.createObjectURL(new Blob(["...code..."])))`.

**Q786: What’s the difference between dedicated and shared workers?**
*   **Dedicated (`new Worker`)**: Linked to one script/tab.
*   **Shared (`new SharedWorker`)**: Can be accessed by multiple scripts/tabs of the same origin.

**Q787: What is the Transferable interface in JS?**
An interface allowing ownership of objects (ArrayBuffer, MessagePort) to be moved to a Worker. `postMessage(data, [transferables])`.

**Q788: What is SharedArrayBuffer and how is it used?**
A buffer whose memory is shared between Main Thread and Workers. Both can read/write directly. Fast but requires synchronization.

**Q789: What are Atomics and how do they ensure safety?**
Static methods (`Atomics.add`, `load`, `store`, `wait`) to perform thread-safe operations on `SharedArrayBuffer` (Int32Array views) to prevent race conditions.

**Q790: How do you implement a thread-safe counter with Atomics?**
```javascript
const sharedBuffer = new SharedArrayBuffer(4);
const int32 = new Int32Array(sharedBuffer);
// In workers:
Atomics.add(int32, 0, 1); // Increments index 0 safely
```

---

## 🔹 23. Edge API Features (Questions 871-900)
*(Source list 791-800) Note: The source section size was smaller than usual.*

**Q791: What is the Payment Request API?**
Standardized API to collect payment info (credit card, specific payment apps) from the user via the browser's native UI.

**Q792: What is the Battery Status API and why was it deprecated?**
Allowed accessing battery level/charging status. Deprecated due to privacy concerns (fingerprinting users).

**Q793: How does the Web Share API work?**
Invokes the native OS share dialog. `navigator.share({ title, text, url })`. Requires user interaction (click).

**Q794: What is the Permissions API?**
Way to query the status of permissions (geolocation, camera) without triggering a prompt. `navigator.permissions.query({ name: 'geolocation' })`.

**Q795: How to detect clipboard read/write availability?**
Use `navigator.permissions` with `'clipboard-read'` or `'clipboard-write'`.

**Q796: What is the Page Visibility API?**
(See Q709). `document.hidden` and `visibilitychange`.

**Q797: What is the Beacon API used for?**
(See Q196). `navigator.sendBeacon(url, data)` sends small data asynchronously to server before unload. Non-blocking.

**Q798: How does the Wake Lock API prevent screen sleeping?**
`navigator.wakeLock.request('screen')`. Prevents device from dimming/locking screen during critical tasks (presentation, navigation).

**Q799: How does the Vibration API work?**
`navigator.vibrate([200, 100, 200])`. Vibrates device (Mobile only).

**Q800: What is the Idle Detection API?**
Detects if user is away from keyboard/screen. `new IdleDetector()`. Requires permission.

---

## 🔹 24. UI + Event Loop Integration (Questions 901-930)
*(Source list 801-810)*

**Q801: How to defer non-critical JS for better Time to Interactive?**
Use `defer` or `async` on scripts. Use Code Splitting. Use `requestIdleCallback` for initialization tasks.

**Q802: What is the difference between `requestIdleCallback()` and `requestAnimationFrame()`?**
*   `rAF`: Runs every frame (High priority, visual).
*   `rIC`: Runs when browser is idle (Low priority, background).

**Q803: How can JS simulate batching DOM reads and writes?**
Use a library like FastDOM or manually group reads (`offsetWidth`) first, then writes (`style.width`) frame-by-frame to avoid thrashing.

**Q804: What is layout thrashing and how to avoid it?**
(Duplicate Q173). Interleaved read/write causing reflows. Batch them.

**Q805: What’s a mutation observer and how does it affect reactivity?**
(Duplicate Q162). It fires microtasks when DOM changes. Can implement "reactivity" by watching attribute changes.

**Q806: How does JavaScript throttle input rendering on scroll?**
Use `passive` listeners or `requestAnimationFrame` to decouple the scroll handler from the rendering update.

**Q807: What is the difference between synchronous vs async click handling?**
Native clicks are sync (stack usually empty). If you trigger click programmatically `el.click()`, it executes handlers synchronously.

**Q808: What happens if you trigger layout inside a JS loop?**
Performance disaster. Each iteration forces a Reflow (Layout calculation) if you read layout properties after writing style.

**Q809: Why should you avoid frequent style recalculations?**
They are expensive (CPU).

**Q810: How to use ResizeObserver for responsive elements?**
Does not rely on window resize. Fires when specific element size changes (e.g. container query polyfill).

---

## 🔹 25. Obscure JS Constructs & Pitfalls (Questions 931-960)
*(Source list 811-820)*

**Q811: What does `new.target` represent in a constructor?**
It refers to the constructor function invoked by `new`. If a function is called without `new`, it is undefined. Useful for enforcing usage of `new` or checking subclassing.

**Q812: What’s the result of `typeof function* () {}`?**
`"function"`. Generators are functions.

**Q813: How is a generator paused and resumed internally?**
It saves its execution context (Variable Environment, stack pointer) when `yield` is called. `next()` restores the context.

**Q814: What happens if you yield inside a try block?**
The generator pauses. `next()` resumes inside `try`. `throw()` injects an error at the yield point, triggering `catch`.

**Q815: What is the “temporal dead zone” with `let`?**
(Duplicate Q101).

**Q816: Can you redeclare `let` or `const` in the same block?**
No. `SyntaxError`.

**Q817: What is the return value of `void` operator?**
`undefined`.

**Q818: How is `.length` of a function determined?**
Number of formal parameters (excluding rest parameters and defaults).

**Q819: What does `arguments.callee` do and why is it deprecated?**
Refers to the currently executing function. Deprecated because it hinders optimization and breaks in strict mode. Use named function expressions instead.

**Q820: How does `Function.length` differ from `arguments.length`?**
*   `Function.length`: Expected parameters (definition).
*   `arguments.length`: Actual passed arguments (execution).

---

## 🔹 26. Async Patterns and Edge Handling (Questions 961-990)
*(Source list 821-830)*

**Q821: What is an async IIFE?**
`(async () => { await ... })()`. Allows using top-level await logic in older environments or for scoping async logic.

**Q822: What happens if you `await` inside a constructor?**
Syntax Error. Constructors cannot be async. Only methods/functions.

**Q823: How to make a function return synchronously or asynchronously based on context?**
Bad practice (Zalgo anti-pattern). Returns typically should be consistent (Always Promise). If strictly needed, check input type.

**Q824: How do you catch errors thrown from an async generator?**
Wrap the `for await...of` loop in a `try/catch` block.

**Q825: What’s the difference between rejecting a promise vs throwing?**
Inside an `async` function, `throw Error()` is equivalent to returning a rejected Promise.

**Q826: How does `queueMicrotask()` differ from `setTimeout()`?**
`queueMicrotask` schedules on the Microtask queue (runs before next render/macrotask). `setTimeout` schedules on Macrotask queue.

**Q827: How do browser rendering and JS event loop interact?**
Render pipeline (Style/Layout/Paint) typically runs **after** microtasks are cleared and before the next Macrotask (Loop).

**Q828: Can you pause a generator with `await`?**
Yes. `yield await promise` pauses the generator until the promise resolves, then yields the result.

**Q829: How do you cancel an async generator?**
Call `generator.return()`. Breaks the loop consuming it.

**Q830: How do you prioritize microtasks in a browser?**
They are inherently prioritized over Macrotasks and Rendering. Blocking the main thread with too many microtasks hangs the browser.

---

## 🔹 27. New & Experimental Features (Questions 991-1020)
*(Source list 831-840)*

**Q831: What is the `Observable` proposal?**
Reactive programming primitive (like RxJS) for handling streams of async data. Not standard yet.

**Q832: What are decorators and their current TC39 status?**
Stage 3. syntax `@deco` to meta-program classes/methods.

**Q833: What are Records & Tuples?**
(See Q378). Immutable deep structures. `#{}` and `#[]`.

**Q834: What is module attributes proposal?**
(Now Import Attributes). `import json from './data.json' with { type: 'json' }`.

**Q835: How will pattern matching improve control flow?**
`match (val) { when (pattern): ... }`. Logic similar to Rust/Haskell match. Powerful switch replacement.

**Q836: What is `Array.prototype.with()` and how is it different from `splice()`?**
`arr.with(index, value)` returns a **new** array with the element replaced (Immutable). `splice` modifies in-place.

**Q837: What is `Map.groupBy()`?**
(See Q375). Static method `Map.groupBy(items, callback)` returns a Map grouping items.

**Q838: What is `Promise.withResolvers()`?**
Returns `{ promise, resolve, reject }`. Eliminates the "deferred" pattern boilerplate `let res; new Promise(r => res = r)`.

**Q839: What is `ArrayBuffer.transfer()`?**
Moves the memory block to a new buffer (zero-copy), detaching the old one. Useful for performant resizing or transfer.

**Q840: What is the Temporal API and how does it replace Date?**
A modern, consistent API for Date/Time. Fixes `Date` object issues (mutability, parsing, timezones). `Temporal.Now.instant()`.

---

## 🔹 28. Debugging & Observability (Questions 1021-1050)
*(Source list 841-850)*

**Q841: What is a source map?**
A file mapping minified/transpiled code back to original source lines.

**Q842: How to log deep objects safely?**
`console.dir(obj, { depth: null })` (Node). Browser: Console expands lazy references. To snapshot: `JSON.stringify(obj)`.

**Q843: How to monitor async stack traces?**
Enable "Async Stack Traces" in DevTools.

**Q844: What are break-on-access debugging strategies?**
DevTools: "Break on attribute modification". Code: `debugger`. Data Breakpoints.

**Q845: What is a memory snapshot?**
A dump of the JS heap at a specific moment. Used to compare snapshots and find leaks.

**Q846: What does the “Retainers” tab show in Chrome DevTools?**
Which objects are holding a reference to the selected object (preventing GC).

**Q847: What is the Timeline flame chart?**
Visualizes the Call Stack over time during recording.

**Q848: How can you simulate offline in DevTools?**
Network Tab -> Dropdown -> "Offline".

**Q849: What does “Long Task” mean?**
A task taking >50ms. Blocks main thread.

**Q850: How can you track event listeners?**
DevTools Elements Tab -> "Event Listeners" sidebar. Or `getEventListeners(node)` in Console API.
