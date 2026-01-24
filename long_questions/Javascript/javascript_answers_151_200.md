# JavaScript Interview Questions & Answers (151-200)

## 🔹 ES2020+ Features (221–250)

### Question 151: What is the nullish coalescing operator (`??`)?

**Answer:**
The nullish coalescing operator `??` returns the right-hand operand only if the left-hand operand is `null` or `undefined`. Unlike `||`, it does not return the right side for other falsy values like `0`, `""`, or `false`.

**Example:**
```javascript
const count = 0;
const defaultCount = count || 10; // 10 (Wrong, 0 is falsy)
const explicitCount = count ?? 10; // 0 (Correct)
```

---

### Question 152: What is optional chaining and how does it work?

**Answer:**
The `?.` operator allows reading the value of a property located deep within a chain of connected objects without checking each reference in the chain. If a reference is nullish (`null`/`undefined`), it stops and returns `undefined`.

**Example:**
```javascript
const user = {};
// const city = user.address.city; // TypeError
const city = user?.address?.city; // undefined (No error)
```

---

### Question 153: What are private class fields in JavaScript?

**Answer:**
Class fields prefixed with `#` are private to the class. They cannot be accessed or modified from outside the class instance.

**Example:**
```javascript
class Counter {
    #count = 0;
    increment() { this.#count++; }
    getValue() { return this.#count; }
}
const c = new Counter();
c.increment();
console.log(c.getValue()); // 1
// console.log(c.#count); // SyntaxError: Private field
```

---

### Question 154: What is top-level await?

**Answer:**
Examples of using the `await` keyword outside of `async` functions within ES Modules. The module execution waits for the promise to resolve before imports are available to consumers.

**Example:**
```javascript
// data.js
const response = await fetch('https://api.example.com');
export const data = await response.json();
```

---

### Question 155: What are logical assignment operators (`&&=`, `||=`, `??=`)?

**Answer:**
They combine logical operations with assignment.
*   `a ||= b` → `a = a || b` (Assign if `a` is falsy)
*   `a &&= b` → `a = a && b` (Assign if `a` is truthy)
*   `a ??= b` → `a = a ?? b` (Assign if `a` is nullish)

**Example:**
```javascript
let msg = "";
msg ||= "Default"; // "Default" (since "" is falsy)
```

---

### Question 156: What is `BigInt` in JavaScript?

**Answer:**
A primitive wrapper object used to represent integers larger than `2^53 - 1` (which is `Number.MAX_SAFE_INTEGER`). Created by appending `n` to the end of an integer.

**Example:**
```javascript
const big = 9007199254740991n;
const bigger = big + 1n; // 9007199254740992n
// console.log(10n + 10); // TypeError: Cannot mix BigInt and Number
```

---

### Question 157: What is `WeakRef`?

**Answer:**
`WeakRef` allows holding a weak reference to an object without preventing it from being garbage collected. Used implementation of caches or mappings.

**Example:**
```javascript
let obj = { data: "Large" };
const ref = new WeakRef(obj);
obj = null; // Object can be GC'd now

// Later...
const cached = ref.deref();
if (cached) console.log(cached.data);
```

---

### Question 158: What is `FinalizationRegistry`?

**Answer:**
A utility that lets you request a callback when an object is garbage collected.

**Example:**
```javascript
const registry = new FinalizationRegistry((heldValue) => {
    console.log(`Object with ${heldValue} was finalized`);
});
let obj = {};
registry.register(obj, "ID-123");
obj = null; // Callback runs eventually after GC
```

---

### Question 159: How does `Promise.allSettled()` work?

**Answer:**
It takes an iterable of promises and returns a promise that resolves after **all** of the given promises have either fulfilled or rejected. The result is an array of objects describing the outcome (`{ status: 'fulfilled', value: ... }` or `{ status: 'rejected', reason: ... }`).

---

### Question 160: How does dynamic import (`import()`) work?

**Answer:**
`import()` loads a module asynchronously and returns a promise that resolves to the module namespace object. Useful for code splitting/lazy loading.

**Example:**
```javascript
btn.addEventListener('click', async () => {
    const { func } = await import('./module.js');
    func();
});
```

---

## 🔹 DOM & BOM Specific (251–280)

### Question 161: What is the difference between `document.body` and `document.documentElement`?

**Answer:**
*   **`document.body`**: Returns the `<body>` element.
*   **`document.documentElement`**: Returns the `<html>` element (the root element of the document).

---

### Question 162: How do mutation observers work?

**Answer:**
`MutationObserver` creates an interface that provides the ability to watch for changes being made to the DOM tree (attributes, childList, subtree).

**Example:**
```javascript
const observer = new MutationObserver((mutations) => {
    mutations.forEach(m => console.log(m.type));
});
observer.observe(document.body, { childList: true, subtree: true });
```

---

### Question 163: What is event delegation vs event bubbling?

**Answer:**
*   **Bubbling** is the mechanism where an event triggers on the deepest element and propagates up to parents.
*   **Delegation** is the pattern of using a single listener on a parent to handle events for multiple children (leveraging bubbling).

---

### Question 164: How do you implement custom events?

**Answer:**
Using the `CustomEvent` constructor. Can carry custom data in the `detail` property.

**Example:**
```javascript
const event = new CustomEvent('userLogin', { detail: { username: 'Alice' } });
window.addEventListener('userLogin', e => console.log(e.detail.username));
window.dispatchEvent(event);
```

---

### Question 165: What is the difference between `focus()` and `blur()`?

**Answer:**
*   **`focus()`**: Sets focus on the specified element (if it can be focused).
*   **`blur()`**: Removes keyboard focus from the current element.
    *   Note: `focusin` and `focusout` events bubble, while `focus` and `blur` do not.

---

### Question 166: How do you programmatically scroll an element into view?

**Answer:**
Using `element.scrollIntoView()`.

**Example:**
```javascript
element.scrollIntoView({ 
    behavior: 'smooth', 
    block: 'center' 
});
```

---

### Question 167: What are data attributes in HTML/JS?

**Answer:**
Custom attributes prefixed with `data-` (e.g., `data-id="123"`). Accessed in JS via the `dataset` property.

**Example:**
```javascript
// HTML: <div id="user" data-id="123" data-role="admin"></div>
const el = document.getElementById('user');
console.log(el.dataset.id); // "123"
console.log(el.dataset.role); // "admin"
```

---

### Question 168: What are the risks of using `innerHTML`?

**Answer:**
*   **XSS (Cross-Site Scripting):** If user input is inserted without sanitization, malicious scripts can execute.
*   **Performance:** Re-parses and rebuilds all child nodes, detaching existing event listeners.

---

### Question 169: How do `getBoundingClientRect()` and layout work?

**Answer:**
It returns the size of an element and its position relative to the **viewport**. Calling it forces the browser to calculate layout (Reflow), which can be expensive if done repeatedly in a loop.

---

### Question 170: How do you listen for form field changes with JavaScript?

**Answer:**
*   **`input` event:** Fires immediately when the value changes (typing).
*   **`change` event:** Fires when the element loses focus or selection is committed (enter/click away).

---

## 🔹 Performance Optimization (281–310)

### Question 171: What is the difference between lazy loading and preloading?

**Answer:**
*   **Lazy Loading:** Loading resources (images, scripts) **only when needed** (e.g., when scrolled into view). Reduces initial load time.
*   **Preloading:** Loading critical resources **sooner** than the browser discovers them (e.g., `<link rel="preload">`). Optimizes rendering speed.

---

### Question 172: What is critical rendering path?

**Answer:**
The sequence of steps the browser performs to convert HTML, CSS, and JS into pixels on the screen.
Steps: HTML -> DOM -> CSSOM -> Render Tree -> Layout -> Paint.
Optimizing CRP involves minimizing render-blocking resources.

---

### Question 173: What is layout thrashing?

**Answer:**
Occurs when JS repeatedly reads and writes to the DOM in the same frame/loop, causing the browser to recalculate layout (Reflow) multiple times unnecessarily.

**Fix:** Batch reads and writes.

---

### Question 174: What is the `requestIdleCallback()` API?

**Answer:**
A window method that queues a function to be called during a browser's idle periods. Used for low-priority background tasks (analytics, prefetching) without affecting interaction latency.

---

### Question 175: How do you debounce a resize event?

**Answer:**
Resize events fire rapidly. Wrap the handler in a debounce function so it only runs once after the specific pause in resizing.

**Code:**
```javascript
window.addEventListener('resize', debounce(() => {
    console.log("Resize finished");
}, 200));
```

---

### Question 176: What is the cost of deep object cloning?

**Answer:**
Deep cloning is expensive (O(N) where N is total nodes). `JSON.parse/stringify` is slow and lossy. `structuredClone` is better but still costly for large objects. Avoid deep cloning in performance-critical loops.

---

### Question 177: What’s the difference between async vs defer in script tags?

**Answer:**
*   **Normal:** HTML parsing pauses → Script downloads & runs → HTML parsing resumes.
*   **`async`:** Script downloads parallel to HTML parsing. Runs **immediately** when downloaded (pausing HTML). Order not guaranteed.
*   **`defer`:** Script downloads parallel. Runs **after** HTML parsing is complete. Order guaranteed.

---

### Question 178: How does lazy evaluation work?

**Answer:**
Delays the evaluation of an expression until its value is needed.
Examples: Generators (yield), Short-circuiting (`&&`, `||`). Not a built-in feature for all variable types like in Haskell.

---

### Question 179: What is the fastest way to iterate over large arrays?

**Answer:**
Standard `for` loop (`for (let i=0; i<len; i++)`) is generally the fastest (raw performance). `forEach` and `map` have function call overhead. However, modern engines optimize `for...of` and array methods well enough that readability usually wins unless in hot code paths.

---

### Question 180: How to avoid memory leaks in long-running JS apps?

**Answer:**
*   Remove event listeners when elements are removed (`removeEventListener`).
*   Clear timers (`clearInterval`).
*   Nullify references to large objects/DOM nodes no longer needed.
*   Use `WeakMap`/`WeakRef` for caches.

---

## 🔹 Testing, Tooling & Build (311–330)

### Question 181: What’s the difference between unit, integration, and e2e testing?

**Answer:**
*   **Unit**: Tests smallest parts (functions/components) in isolation. Fast.
*   **Integration**: Tests interaction between units (Database + API).
*   **E2E (End-to-End)**: Tests full application flow from user perspective (Browser automation). Slow.

---

### Question 182: What is mocking in JavaScript testing?

**Answer:**
Replacing dependencies (APIs, Database calls, Imported modules) with controlled replacements (mocks/spies) to test the code in isolation without external side effects (e.g., `jest.mock()`).

---

### Question 183: How does code coverage work?

**Answer:**
Measures the percentage of code lines, branches, and functions executed during tests. Tools like Istanbul (used by Jest) instrument the code with counters to track execution.

---

### Question 184: How do you test async code?

**Answer:**
Testing frameworks support `async/await`.
**Jest Example:**
```javascript
test('fetch data', async () => {
    const data = await fetchData();
    expect(data).toBe('Success');
});
```

---

### Question 185: What is the purpose of `tsconfig.json` in JS projects?

**Answer:**
Configuration file for TypeScript projects. It specifies root files and compiler options (target JS version, module system, strict mode rules). Can also be used in JS projects (via `allowJs`) to enable VS Code intellisense.

---

### Question 186: What is a bundler vs a transpiler?

**Answer:**
*   **Bundler (Webpack, Vite):** Combines multiple files (JS, CSS) into a few files (bundles) optimized for browsers. Handles imports.
*   **Transpiler (Babel, SWC):** Translates source code syntax (ES6 -> ES5, TS -> JS) but preserves file structure (usually).

---

### Question 187: How does tree shaking work?

**Answer:**
It analyzes the static import/export statements (ESM) to build a dependency graph. Any export that is not imported or used is marked as dead code and excluded from the final bundle.

---

### Question 188: What is hot module replacement (HMR)?

**Answer:**
A feature in bundlers (Webpack/Vite) that exchanges, adds, or removes modules while the application is running, without a full page reload. Retains application state.

---

### Question 189: What is linting and how does ESLint help?

**Answer:**
Linting works by statically analyzing code to find problems (syntax errors, bad patterns, style violations). ESLint enforces rules to ensure code consistency and prevent bugs.

---

### Question 190: What is the difference between npm and yarn?

**Answer:**
Both are package managers.
*   **npm:** Default with Node.js.
*   **yarn:** Introduced for speed/determinism (mostly caught up by npm now).
*   Differences in lock file format (`package-lock.json` vs `yarn.lock`) and command syntax (`npm install` vs `yarn add`).

---

## 🔹 Browser, Events, Storage, Network (331–350)

### Question 191: How does Service Worker work?

**Answer:**
A script that runs in the background, separate from the webpage. It intercepts network requests (acting as a proxy), enabling features like Offline support (caching), Push Notifications, and Background Sync.

---

### Question 192: How do you cache API data in the browser?

**Answer:**
*   **Service Worker (Cache API):** For offline capability.
*   **`localStorage`/`sessionStorage`:** For simple key-value data.
*   **IndexedDB:** For large structured data.
*   **HTTP Cache Headers:** (`Cache-Control`) configured on server.

---

### Question 193: What is the difference between cookies, localStorage, and sessionStorage?

**Answer:**
*   **Cookies:** Small (4KB), sent with every HTTP request. Used for Auth.
*   **`localStorage`:** Large (5MB+), persistent on client-side.
*   **`sessionStorage`:** Large, cleared when tab closes.

---

### Question 194: What is a Broadcast Channel API?

**Answer:**
Allows communication (message bus) between different browsing contexts (tabs, windows, iframes) of the same origin.

**Example:**
```javascript
const bc = new BroadcastChannel('test_channel');
bc.postMessage('Hello Tabs!');
```

---

### Question 195: What are CORS preflight requests?

**Answer:**
When a browser makes a cross-origin request that isn't "simple" (e.g., uses `PUT`, custom headers), it sends an `OPTIONS` request first to verify if the server permits the actual request.

---

### Question 196: What is the Fetch KeepAlive option?

**Answer:**
`fetch(url, { keepalive: true })`. Allows the request to outlive the page. Useful for sending analytics data when the user unloads the page (navigates away). Replacement for `navigator.sendBeacon`.

---

### Question 197: What is the purpose of `navigator` object?

**Answer:**
Contains information about the browser environment (User Agent, Platform, Language, Online status, Hardware concurrency).

---

### Question 198: How do you detect if the user is online or offline?

**Answer:**
*   Property: `navigator.onLine` (boolean).
*   Events: `window.addEventListener('online', ...)` / `offline`.

---

### Question 199: What is the difference between long polling and WebSockets?

**Answer:**
*   **Long Polling:** Client requests, server holds open until data is available, then responds. Client requests again. (HTTP overhead).
*   **WebSockets:** Persistent, full-duplex TCP connection. Low latency, real-time bi-directional communication.

---

### Question 200: What is the `Performance` API?

**Answer:**
Provides access to high-resolution timing data (`performance.now()`) and metrics (Navigation Timing, Resource Timing) to measure webpage performance accurately.
