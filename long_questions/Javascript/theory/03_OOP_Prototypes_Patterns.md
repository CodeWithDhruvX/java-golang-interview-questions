# 🟠 **JavaScript OOP & Prototypes — Mid / Senior Level**

> **Target Companies:** Mindtree, Mphasis, Hexaware (Service-Based) + Mid-Senior SDE roles at Zomato, Swiggy, Paytm, Razorpay (Product-Based)

---

### 56. What is a Map?

"A **Map** is a key-value collection where **keys can be any type** — objects, functions, primitives — unlike plain objects where keys are always strings or symbols.

```js
const map = new Map();
map.set('name', 'Dhruv');
map.set(42, 'the answer');
map.set({ id: 1 }, 'user object as key');

map.get('name');         // 'Dhruv'
map.size;                // 3
map.has(42);             // true
map.delete('name');
map.forEach((val, key) => console.log(key, val));
```

Map preserves **insertion order** and is iterable directly. I use it over plain objects when I need non-string keys or when I need reliable `.size` and iteration."

#### Indepth
Map is significantly faster than Object for frequent add/delete operations because objects carry prototype overhead. `Map.prototype` has 5 key methods: `get`, `set`, `has`, `delete`, `clear`. Converting: `Object.fromEntries(map)` → plain object; `new Map(Object.entries(obj))` → Map. `WeakMap` keys must be objects, are weakly referenced (GC-able), and have no `.size` or iteration — perfect for storing private metadata on DOM nodes.

---

### 57. How to implement a stack/queue using JavaScript?

"Both can be implemented using an **array**:

```js
// Stack (LIFO) — use push/pop
const stack = [];
stack.push(1); stack.push(2); stack.push(3);
stack.pop();  // 3 (last in, first out)

// Queue (FIFO) — use push/shift
const queue = [];
queue.push(1); queue.push(2); queue.push(3);
queue.shift(); // 1 (first in, first out)
```

For performance-critical queues, `shift()` is O(n) because it re-indexes the array. A linked-list-based queue is O(1) for both enqueue and dequeue."

#### Indepth
A O(1) queue using a linked list:
```js
class Queue {
  #head = null; #tail = null; #size = 0;
  enqueue(val) {
    const node = { val, next: null };
    this.#tail ? (this.#tail.next = node) : (this.#head = node);
    this.#tail = node; this.#size++;
  }
  dequeue() {
    if (!this.#head) return undefined;
    const val = this.#head.val;
    this.#head = this.#head.next;
    if (!this.#head) this.#tail = null;
    this.#size--; return val;
  }
  get size() { return this.#size; }
}
```
Node.js's internal event queues use efficient data structures rather than arrays for this exact reason.

---

### 58. What are WeakMap and WeakSet?

"**WeakMap** is like a Map but keys must be **objects** and are held **weakly** — if no other reference to the key object exists, it can be garbage collected.
**WeakSet** is like a Set but only stores objects, also held weakly.

```js
const cache = new WeakMap();
let user = { id: 1 };
cache.set(user, { profile: '...' });

user = null; // user object can now be GC'd
// cache entry is automatically cleaned up
```

They have NO `size`, NO iteration, and NO `clear()` method. This is intentional — you can't enumerate them because GC timing is non-deterministic."

#### Indepth
Primary use cases: 1) **Private per-instance data** for class instances (before `#` fields). 2) **DOM node caching** — when the DOM node is removed, the cached data is automatically cleaned up. 3) **Memoization caches** where the key is an object — avoids holding references that prevent GC. Libraries like React internally use WeakMap to associate component metadata with DOM nodes.

---

### 59. How do you format a date in JavaScript?

"Modern way using `Intl.DateTimeFormat`:
```js
const date = new Date();

// Locale-aware formatting
new Intl.DateTimeFormat('en-IN', {
  year: 'numeric', month: 'long', day: 'numeric'
}).format(date);
// 'February 20, 2026'

// Quick toLocaleDateString
date.toLocaleDateString('en-US'); // '2/20/2026'

// ISO string (for APIs/storage)
date.toISOString(); // '2026-02-20T09:02:30.000Z'
```

I use `Intl.DateTimeFormat` for display and ISO strings for storage/API. For complex date manipulation, I use **date-fns** (tree-shakeable) over Moment.js (deprecated, large bundle)."

#### Indepth
`new Date()` gives a timestamp in **local time**, `.toISOString()` always gives UTC. A common bug: `new Date('2026-02-20')` is parsed as UTC midnight, but `new Date('2026/02/20')` is parsed as local midnight — different times! Always use `date-fns`, `Luxon`, or the upcoming **Temporal API** for reliable date math. The Temporal API (TC39 Stage 3) will replace `Date` with an immutable, timezone-aware API.

---

### 60. How to get the difference between two dates?

"Subtract timestamps and convert milliseconds to your desired unit:

```js
const start = new Date('2026-01-01');
const end = new Date('2026-02-20');

const diffMs = end - start;                    // milliseconds
const diffDays = diffMs / (1000 * 60 * 60 * 24); // 50 days
const diffHours = diffMs / (1000 * 60 * 60);

// With date-fns (recommended)
import { differenceInDays, differenceInHours } from 'date-fns';
differenceInDays(end, start); // 50
```

For production, I always use **date-fns** or **Luxon** to avoid DST (Daylight Saving Time) edge cases that raw millisecond math gets wrong."

#### Indepth
DST is the main pitfall: when clocks change, a 'day' might be 23 or 25 hours. Raw `diffMs / 86400000` can give wrong answers around DST transitions. `date-fns`'s `differenceInDays` accounts for this by comparing calendar dates, not raw milliseconds. For duration formatting: `Intl.DurationFormat` (newer) or `date-fns`'s `formatDuration`.

---

### 61. How do regular expressions work in JavaScript?

"A regular expression is a **pattern** used to match, search, or replace strings.

```js
const regex = /^[a-z]+@[a-z]+\.[a-z]{2,}$/i;

regex.test('dhruv@gmail.com');   // true
'Hello World'.match(/\w+/g);     // ['Hello', 'World']
'a1b2c3'.replace(/\d/g, '#');    // 'a#b#c#'
'one two three'.split(/\s+/);    // ['one', 'two', 'three']
```

Flags: `g` (global), `i` (case-insensitive), `m` (multiline), `s` (dotAll — `.` matches newlines)."

#### Indepth
Regex in JS uses the **NFA (Non-deterministic Finite Automaton)** engine, which supports backtracking. Catastrophic backtracking (ReDoS) is a real security risk — a malicious input can cause exponential time matching. Example: `/^(a+)+$/` on `'aaaaaab'` can hang. Mitigate with: atomic groups (when available), possessive quantifiers, or input length limits. Use `RegExp.prototype.exec()` in a loop with `g` flag for iterative matching with capture groups.

---

### 62. How to validate an email using regex?

"A reasonable email validation regex:
```js
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

function isValidEmail(email) {
  return emailRegex.test(email.toLowerCase());
}

isValidEmail('dhruv@gmail.com');    // true
isValidEmail('invalid@');           // false
isValidEmail('no-at-sign.com');     // false
```

I keep the regex simple on the frontend and do **proper validation on the backend** (by actually sending a verification email). No regex can fully validate email per RFC 5321."

#### Indepth
The RFC 5321/5322 compliant email regex is several hundred characters long and still doesn't handle all edge cases (e.g., `"user name"@example.com` is technically valid). The only true validation is **sending a confirmation email**. Frontend regex is just for UX (early feedback), not security. HTML5 `<input type="email">` provides browser-native validation that handles most common cases well.

---

### 63. What is NaN?

"`NaN` stands for **Not a Number** — a special numeric value representing an invalid math operation result.

```js
typeof NaN;           // 'number' (paradox!)
NaN === NaN;          // false (NaN is not equal to itself)
Number('hello');      // NaN
0 / 0;                // NaN
Math.sqrt(-1);        // NaN

// Proper check
Number.isNaN(NaN);    // true  ✅ (strict)
isNaN('hello');       // true  ⚠️ (coerces first — less reliable)
```"

#### Indepth
`NaN` is the only value in JavaScript that is **not equal to itself** — this is mandated by the IEEE 754 floating-point standard. `Number.isNaN()` (ES6) is always preferred over global `isNaN()` because the global version coerces the argument first: `isNaN('hello')` is `true` because `Number('hello')` is `NaN`. `Object.is(NaN, NaN)` returns `true`.

---

### 64. What is the difference between `parseInt()` and `Number()`?

"`parseInt(str, radix)` parses a string and returns an **integer**, stopping at the first non-numeric character. `Number(val)` converts the **entire value** and returns `NaN` for anything invalid.

```js
parseInt('42px');     // 42  (stops at 'p')
Number('42px');       // NaN (entire string must be valid)

parseInt('0xFF', 16); // 255 (hex parsing)
parseInt('010');      // 10  (in modern engines; always specify radix!)

Number('');           // 0
parseInt('');         // NaN
```"

#### Indepth
Always specify the **radix** (second argument) for `parseInt`. Without it, old engines treated strings starting with `0` as octal — `parseInt('010')` would return `8`. `.parseFloat()` parses floating-point numbers similarly to `parseInt`. For strict numeric conversion: `Number(val)` or the unary `+` operator (`+'42'`). For user input: validate first, then parse.

---

### 65. How do you deep clone an object?

"**Modern best practice**: Use `structuredClone()` — available natively since Chrome 98, Node 17:

```js
const original = { a: 1, b: { c: [1, 2, 3] }, d: new Date() };
const clone = structuredClone(original);
clone.b.c.push(4);
original.b.c; // [1, 2, 3] — untouched ✅
```

**Legacy alternatives**:
- `JSON.parse(JSON.stringify(obj))` — fast but loses functions, Dates, undefined, Symbol
- `lodash.cloneDeep(obj)` — handles edge cases, battle-tested

I use `structuredClone` for most cases today. It supports: nested objects, arrays, Dates, Maps, Sets, RegExp, ArrayBuffers, and circular references."

#### Indepth
`structuredClone` uses the **HTML Structured Clone Algorithm** and correctly handles: circular references (no stack overflow), `Date`, `Map`, `Set`, `RegExp`, `ArrayBuffer`, `BigInt`. It does NOT clone: functions, DOM nodes, class instances (loses prototype). For class instances, you need a custom serialization strategy (like `toJSON`/`fromJSON`). `JSON.parse(JSON.stringify(obj))` is O(n) space and O(n) time — `structuredClone` is similar but more correct.

---

### 66. What is the module pattern?

"The **module pattern** uses an IIFE with closures to create **public and private members**:

```js
const counter = (() => {
  let _count = 0; // private

  return {
    increment() { _count++; },
    decrement() { _count--; },
    getCount() { return _count; }
  };
})();

counter.increment();
counter.getCount(); // 1
counter._count;     // undefined (private!)
```

Before ES6 modules, this was THE way to organize code without polluting the global namespace. jQuery's entire codebase is wrapped in a module pattern IIFE."

#### Indepth
The **Revealing Module Pattern** variation returns an object that maps private function names to public ones: `return { add: increment, get: getCount }`. This gives cleaner internal naming. ES6 modules (`export`/`import`) supersede this pattern for new code — they have lexical scoping, static analysis, and tree-shaking support. But understanding the module pattern explains how pre-module JS codebases work.

---

### 67. How does the singleton pattern work in JS?

"A **singleton** ensures only one instance of a class/object is created:

```js
class Database {
  static #instance = null;

  constructor() {
    if (Database.#instance) return Database.#instance;
    // Initialize DB connection
    this.connection = createConnection();
    Database.#instance = this;
  }

  static getInstance() {
    return Database.#instance ?? new Database();
  }
}

const db1 = new Database();
const db2 = new Database();
db1 === db2; // true ✅
```"

#### Indepth
ES modules are natural singletons — a module is only evaluated once, and the same exports are shared. So a module that creates and exports a connection is a de-facto singleton: `export const db = new Database()`. In Node.js, modules are cached in `require.cache`, making them singletons within a process. The classic singleton anti-pattern: using it for mutable global state makes testing hard (use **dependency injection** instead).

---

### 68. What is the observer pattern?

"The **observer pattern** (also called pub/sub) defines a **one-to-many relationship** — when one object (subject) changes state, all its dependents (observers) are notified.

```js
class EventEmitter {
  #events = {};

  on(event, listener) {
    this.#events[event] ??= [];
    this.#events[event].push(listener);
  }

  emit(event, ...args) {
    this.#events[event]?.forEach(fn => fn(...args));
  }

  off(event, listener) {
    this.#events[event] = this.#events[event]?.filter(l => l !== listener);
  }
}

const emitter = new EventEmitter();
emitter.on('data', (payload) => console.log(payload));
emitter.emit('data', { id: 1 }); // logs { id: 1 }
```"

#### Indepth
Node.js's `EventEmitter` class is the canonical implementation. React's state system is essentially an observer pattern — component re-renders when subscribed state changes. In browsers, the DOM event system is also observer-based. The risk: **memory leaks** if listeners aren't removed (`.off()` / `removeEventListener`). React's `useEffect` cleanup solves this for component lifecycles.

---

### 69. What is the factory pattern?

"The **factory pattern** uses a function (or class) to create and return objects without exposing the instantiation logic:

```js
function createUser(role) {
  const base = { name: '', permissions: [] };

  const roles = {
    admin: { ...base, permissions: ['read', 'write', 'delete'] },
    editor: { ...base, permissions: ['read', 'write'] },
    viewer: { ...base, permissions: ['read'] },
  };

  return roles[role] ?? { ...base };
}

const admin = createUser('admin');
const viewer = createUser('viewer');
```"

#### Indepth
Factory functions are powerful in JS because they can return any object without `new`. They also naturally support **closures** (private state). Abstract factories create families of related objects. In React, components are essentially factories — they return JSX elements. HOCs (Higher-Order Components) are factory functions that take a component and return an enhanced component.

---

### 70. What is transpilation (e.g., Babel)?

"**Transpilation** converts code from one language version to another at the same abstraction level — as opposed to compilation which goes to a lower level (like machine code).

**Babel** transpiles modern ES2022+ JavaScript to older ES5 that older browsers can understand:

```js
// Modern (input)
const greet = (name = 'World') => `Hello, ${name}!`;

// Transpiled ES5 output
'use strict';
var greet = function() {
  var name = arguments.length > 0 ? arguments[0] : 'World';
  return 'Hello, ' + name + '!';
};
```

Babel plugins transform specific syntax; presets are bundles of plugins."

#### Indepth
Babel works in 3 phases: **parse** (code → AST), **transform** (AST → modified AST via plugins), **generate** (AST → code + source maps). Source maps link generated code back to original, enabling browser DevTools to show original source. **SWC** (Rust-based) and **esbuild** (Go-based) are modern Babel alternatives that are 10-100x faster, making them preferred in Vite and Next.js.

---

### 71. How does event bubbling/capturing work?

"DOM events propagate in **3 phases**:
1. **Capture phase**: event travels from `window` → down to the target element
2. **Target phase**: event is at the actual target element
3. **Bubble phase**: event travels back up from target → `window`

By default, `addEventListener` listens in the **bubble phase**. Pass `{ capture: true }` to listen in the capture phase.

```js
// Capture phase (top-down)
document.addEventListener('click', handler, { capture: true });
// Bubble phase (bottom-up) — default
document.addEventListener('click', handler);
```"

#### Indepth
Most events bubble, but some don't: `focus`, `blur`, `load`, `unload`, `scroll`. Their alternatives (`focusin`, `focusout`) do bubble. Event **delegation** leverages bubbling — attach one listener on the parent to handle events from many children. `event.target` is the element that triggered the event; `event.currentTarget` is the element the listener is attached to. `stopPropagation()` stops further propagation in both directions.

---

### 72. What is event delegation?

"**Event delegation** attaches a **single event listener** on a parent element to handle events from all its children, using event bubbling.

```js
// Instead of 100 separate listeners on each <li>
document.getElementById('todo-list').addEventListener('click', (e) => {
  if (e.target.matches('li.todo-item')) {
    toggleTodo(e.target.dataset.id);
  }
  if (e.target.matches('.delete-btn')) {
    deleteItem(e.target.closest('li').dataset.id);
  }
});
```

Benefits: fewer listeners (better memory), works for dynamically added elements, simpler cleanup."

#### Indepth
`e.target.matches(selector)` checks if the clicked element matches a CSS selector. `e.target.closest(selector)` traverses up the DOM to find the nearest ancestor matching the selector — useful for clicking inside a complex card and still knowing which card was clicked. Virtual DOM frameworks (React) implement their own synthetic event system using delegation on the root container.

---

### 73. How does garbage collection work in JS?

"JavaScript uses **automatic garbage collection** — you don't manually free memory. The most common algorithm is **Mark and Sweep**:

1. **Mark phase**: starting from 'roots' (global variables, call stack), the GC marks all reachable objects
2. **Sweep phase**: all unmarked (unreachable) objects are freed

```js
let obj = { data: 'heavy' };
// obj is reachable — NOT collected

obj = null;
// original object is now unreachable — eligible for GC
```"

#### Indepth
V8 uses a generational GC. Short-lived objects go to **young generation** (collected frequently, fast). Long-lived objects are promoted to **old generation** (collected less frequently). This aligns with the **generational hypothesis**: most objects die young. GC runs in the background but can cause **pause jank** — V8 uses **concurrent/parallel marking** to minimize main-thread pauses. Avoid memory leaks by removing event listeners, clearing timers, and nulling references.

---

### 74. What are memory leaks and how to avoid them?

"A **memory leak** occurs when objects are no longer needed but still referenced, preventing GC from reclaiming memory. Common causes:

1. **Forgotten event listeners** — add without removing
2. **Detached DOM nodes** — removed from DOM but still referenced in JS
3. **Closures holding large data** — closure keeps reference alive
4. **Static/global collections** that grow infinitely

```js
// Leak: event listener never removed
button.addEventListener('click', handler); // 🔴
// Fix:
button.removeEventListener('click', handler); // ✅
// Or use AbortController to remove all at once
```"

#### Indepth
To detect leaks: Chrome DevTools Memory tab → take Heap Snapshots → compare between snapshots → look for growing object counts. The **Retainers** panel shows what's keeping an object alive. `WeakMap` and `WeakRef` help by allowing GC to collect objects even if you hold a reference. In SPAs (React, Vue), component unmount cleanup (`useEffect` return function) is critical — forgetting to clean up timers or subscriptions is the most common React memory leak.

---

### 75. How to profile JavaScript performance?

"**Browser DevTools Performance tab** is the primary tool:
1. Open Chrome DevTools → Performance tab
2. Click Record → interact with your app → Stop
3. Analyze the **flame chart** — wide bars indicate slow functions
4. Look for **Long Tasks** (>50ms), **layout thrashing**, and **forced reflows**

```js
// Code-level profiling
console.time('myOperation');
doHeavyWork();
console.timeEnd('myOperation'); // logs: 'myOperation: 120ms'

// Precision measurement
performance.mark('start');
doWork();
performance.measure('work', 'start');
console.log(performance.getEntriesByName('work')[0].duration);
```"

#### Indepth
The **Performance API** (`window.performance`) gives sub-millisecond precision. `performance.now()` returns a high-resolution timestamp. In Node.js: `process.hrtime.bigint()` for nanosecond precision. Key metrics to watch: **FCP** (First Contentful Paint), **LCP** (Largest Contentful Paint), **TTI** (Time to Interactive), **CLS** (Cumulative Layout Shift) — these are the **Web Vitals** measured by Google.

---

### 76. What is the Fetch API?

"The **Fetch API** is the modern way to make HTTP requests in the browser, replacing `XMLHttpRequest`:

```js
async function getUser(id) {
  const response = await fetch(`/api/users/${id}`, {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${token}` },
  });

  if (!response.ok) {
    throw new Error(`HTTP error: ${response.status}`);
  }

  return response.json();
}

// POST example
await fetch('/api/users', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ name: 'Dhruv' }),
});
```"

#### Indepth
Key gotcha: `fetch` **only rejects on network failure** — a 404 or 500 response still resolves! Always check `response.ok` (true for 200–299). `response.json()` is an async method that reads the stream. Other response methods: `.text()`, `.blob()`, `.arrayBuffer()`, `.formData()`. For cancellation, pass an `AbortSignal`: `fetch(url, { signal: controller.signal })`.

---

### 77. How does `localStorage` differ from `sessionStorage`?

"`localStorage` persists data **indefinitely** (until explicitly cleared) across browser sessions and tabs.
`sessionStorage` persists data only for the **current tab session** — it's cleared when the tab is closed.

Both have the same API:
```js
localStorage.setItem('theme', 'dark');
localStorage.getItem('theme');         // 'dark'
localStorage.removeItem('theme');
localStorage.clear(); // removes all

sessionStorage.setItem('token', 'abc');
sessionStorage.getItem('token');
```

Both store string-only values (remember to `JSON.stringify`/`JSON.parse` objects), and both have a ~5MB limit per origin."

#### Indepth
Storage is **per origin** (protocol + domain + port). `localStorage` is **shared across all tabs** of the same origin — use `storage` event to sync: `window.addEventListener('storage', e => { if(e.key==='theme') apply(e.newValue); })`. Security: never store JWT tokens or sensitive data in localStorage — it's vulnerable to XSS. Use **HttpOnly cookies** for auth tokens. `IndexedDB` is the alternative for large structured data storage.

---

### 78. What is `requestAnimationFrame`?

"`requestAnimationFrame(callback)` schedules a callback to run **before the browser's next repaint** — typically ~60fps (every 16.67ms). It's the optimal way to perform animations.

```js
function animate(timestamp) {
  // 'timestamp' is a DOMHighResTimeStamp
  element.style.left = `${Math.sin(timestamp / 1000) * 100}px`;
  requestAnimationFrame(animate); // schedule next frame
}

requestAnimationFrame(animate); // start the loop

// Cancel it:
const id = requestAnimationFrame(animate);
cancelAnimationFrame(id);
```"

#### Indepth
RAF is superior to `setInterval` for animations because: 1) **Pauses when tab is hidden** (battery saving), 2) **Syncs with display refresh rate** (no tearing), 3) **Batches visual updates** efficiently. The callback receives a `DOMHighResTimeStamp` — use it to calculate elapsed time for smooth framerate-independent animations. For non-visual background tasks, `requestIdleCallback` is the alternative.

---

### 79. What is XSS and how do you prevent it in JS?

"**Cross-Site Scripting (XSS)** is an attack where malicious scripts are injected into your web page and executed in users' browsers.

```js
// Vulnerable — user input rendered as HTML
element.innerHTML = userInput; // 🔴 XSS if input is '<script>...'

// Safe — renders as text
element.textContent = userInput; // ✅

// Safe HTML injection — sanitize first
import DOMPurify from 'dompurify';
element.innerHTML = DOMPurify.sanitize(userInput); // ✅
```

Prevention: Use `textContent`, CSP headers, sanitize input with DOMPurify, avoid `eval`, and use HttpOnly cookies."

#### Indepth
Three types of XSS: 1) **Stored XSS** (malicious script saved in DB, served to all users), 2) **Reflected XSS** (script in URL, reflected by server), 3) **DOM-based XSS** (script injected through client-side JS reading URL params). **Content Security Policy (CSP)** is the most powerful defense: `Content-Security-Policy: script-src 'self'` prevents loading scripts from other origins. Use `nonce` for inline scripts: `<script nonce="random-nonce">` paired with the CSP header.

---

### 80. What is CSRF and how does JavaScript help mitigate it?

"**Cross-Site Request Forgery (CSRF)** tricks a user's browser into making unwanted requests to a site where they're authenticated (using their session cookies).

JavaScript-based mitigations:
1. **CSRF tokens** — the server generates a unique token, JS reads it from a meta tag and includes it in every state-changing request header
2. **SameSite cookies** — `Set-Cookie: session=...; SameSite=Strict` (server-side but JS can't override)
3. **Custom headers** — CSRF attacks can't set custom headers (like `X-Requested-With`) from other origins

```js
// Including CSRF token in fetch
const csrfToken = document.querySelector('[name=csrf-token]').content;
fetch('/api/transfer', {
  method: 'POST',
  headers: { 'X-CSRF-Token': csrfToken },
  body: JSON.stringify({ amount: 100 }),
});
```"

#### Indepth
Modern mitigation: `SameSite=Strict` or `SameSite=Lax` cookie attribute prevents the browser from sending cookies on cross-site requests — effectively neutralizing most CSRF attacks without tokens. `SameSite=Lax` is now the browser default. JWT tokens stored in memory (not cookies) are immune to CSRF since they must be explicitly added to headers (something foreign sites can't do without XSS).
