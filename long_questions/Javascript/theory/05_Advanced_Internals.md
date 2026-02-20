# 🔴 **JavaScript Advanced Internals — Senior / Expert Level**

> **Target Companies:** Infosys Elite, Wipro Turbo, Cognizant Intelligent Operations (Service-Based) + Google, Microsoft, Amazon, Atlassian, Stripe (Product-Based / FAANG)

---

### 101. What are labeled statements?

"**Labeled statements** allow you to name a loop or block, enabling `break` or `continue` to target specific outer loops.

```js
outer: for (let i = 0; i < 3; i++) {
  for (let j = 0; j < 3; j++) {
    if (j === 1) break outer;  // breaks the OUTER loop
    console.log(i, j);
  }
}
// Output: 0 0

// Without label, break only exits the inner loop
```

I rarely use labeled statements in application code — they can reduce clarity. They appear in generated code or matrix algorithms where breaking multiple loop levels is necessary."

#### Indepth
Labels are NOT `goto` — they only influence `break` and `continue` for loops and can't jump to arbitrary code positions. The label scope is lexical and block-scoped. In modern code, breaking out of nested loops is usually better achieved by refactoring the inner loop into a function (and using `return`) or using `Array.some()` which stops iteration when the callback returns `true`.

---

### 102. What is `Object.defineProperty()`?

"`Object.defineProperty()` gives you **fine-grained control** over a property's behavior using property descriptors.

```js
const obj = {};
Object.defineProperty(obj, 'name', {
  value: 'Dhruv',
  writable: false,     // can't reassign
  enumerable: false,   // won't show in for..in or Object.keys()
  configurable: false, // can't delete or redefine
});

obj.name = 'Other'; // silently fails (throws in strict mode)
Object.keys(obj);   // [] — 'name' is not enumerable

// Accessor property (getter/setter)
Object.defineProperty(obj, 'loggedName', {
  get() { console.log('accessed!'); return this._name; },
  set(val) { this._name = val.toUpperCase(); },
  enumerable: true,
  configurable: true,
});
```"

#### Indepth
The **two types of property descriptors**: 1) **Data descriptors** (`value`, `writable`) — stores a value directly. 2) **Accessor descriptors** (`get`, `set`) — uses functions for access. They're mutually exclusive. `Object.defineProperties()` defines multiple at once. `Object.getOwnPropertyDescriptor(obj, 'name')` retrieves them. Vue 2's reactivity system was built entirely on `Object.defineProperty` getters/setters. Vue 3 switched to `Proxy` which is more powerful.

---

### 103. What is the `void` operator?

"The `void` operator **evaluates an expression** and always returns `undefined`.

```js
void 0;                    // undefined
void 'anything';           // undefined
void someFunction();       // evaluates function, returns undefined

// Classic use: in <a> tags to prevent navigation
// <a href=\"javascript:void(0)\">Click me</a>

// Guaranteeing undefined (before ES5, 'undefined' could be reassigned)
if (x === void 0) { /* x is undefined */ }
```

Modern use: in async IIFEs where you want to explicitly discard the return value and signal intention."

#### Indepth
`void` is important for **async IIFEs**: `void async function() { await doSomething(); }()`. This avoids returning a Promise that might be mistakenly `await`ed. Also important: before ES5, `undefined` was a writable global, so `void 0` was the safe way to get the `undefined` value. Today, you can just use `undefined` safely. In minifiers, `void 0` is preferred over `undefined` because it's 4 characters shorter.

---

### 104. How does `instanceof` really work internally?

"`instanceof` checks if `Constructor.prototype` exists **anywhere in the prototype chain** of an object.

```js
const arr = [1, 2, 3];
arr instanceof Array;    // true
arr instanceof Object;   // true (Array.prototype's [[Prototype]] is Object.prototype)
arr instanceof Map;      // false

// Internally equivalent to:
Array.prototype.isPrototypeOf(arr); // true
```

It can be **customized** with `Symbol.hasInstance`:

```js
class EvenNumber {
  static [Symbol.hasInstance](num) {
    return typeof num === 'number' && num % 2 === 0;
  }
}
4 instanceof EvenNumber; // true
```"

#### Indepth
`instanceof` fails across **different realms** (iframes, separate Node.js `vm` contexts) because each realm has its own `Array.prototype`. A value created in an iframe with `new Array()` fails `instanceof Array` in the main window. This is why `Array.isArray()` is safer — it checks the `[[Class]]` internal slot rather than the prototype chain, working correctly across realms.

---

### 105. How is the `new` keyword implemented?

"The `new` operator does 4 things:
1. Creates a new empty object `{}`
2. Sets the object's `[[Prototype]]` to `Constructor.prototype`
3. Calls the constructor with `this` set to the new object
4. Returns the new object (or the constructor's explicit return if it returns an object)

```js
// Implementing 'new' manually:
function myNew(Constructor, ...args) {
  // Step 1 & 2: Create object with correct prototype
  const obj = Object.create(Constructor.prototype);
  // Step 3: Call constructor
  const result = Constructor.apply(obj, args);
  // Step 4: Return object or constructor's return value if it's an object
  return result instanceof Object ? result : obj;
}

function Dog(name) { this.name = name; }
const rex = myNew(Dog, 'Rex');
rex instanceof Dog; // true
```"

#### Indepth
The override behavior (Step 4) is why `return { custom: true }` in a constructor replaces the new instance — any returned **object** overrides it. Returning a **primitive** is ignored. `new.target` (inside a constructor or function) tells you if the function was called with `new` — it's the constructor function itself if called with `new`, `undefined` if called normally. Use `new.target` to enforce constructor usage.

---

### 106. What is a "polyfill" in JavaScript?

"A **polyfill** is code that implements a feature that the browser **doesn't natively support** — it fills the gap in older environments.

```js
// Polyfill for Array.prototype.includes (ES7)
if (!Array.prototype.includes) {
  Array.prototype.includes = function(searchElement, fromIndex) {
    if (this == null) throw new TypeError('"this" is null or undefined');
    const arr = Object(this);
    const len = arr.length >>> 0;
    if (len === 0) return false;
    // ... full implementation
    return false;
  };
}
```

Modern approach: Use **core-js** (comprehensive polyfill library) via Babel's `useBuiltIns: 'usage'` which automatically includes only the polyfills needed for target browsers."

#### Indepth
**Polyfill** vs **Shim** vs **Transpiler**: A polyfill adds missing functionality. A shim adapts old API to a new interface. A transpiler transforms syntax (syntax can't be polyfilled — you can't polyfill arrow functions or `async/await`, only transpile them). Services like **Polyfill.io** deliver customized polyfill bundles based on the `User-Agent`. In 2024, most production apps target modern browsers and need fewer polyfills — but they're still critical for enterprise apps supporting IE11 or older mobile browsers.

---

### 107. What is feature detection in JavaScript?

"**Feature detection** checks if a browser supports a feature before using it, rather than checking the browser name/version (user-agent sniffing).

```js
// ❌ Browser sniffing (fragile, breaks with browser updates)
if (navigator.userAgent.includes('Chrome')) { ... }

// ✅ Feature detection (robust)
if ('IntersectionObserver' in window) {
  // Use IntersectionObserver
} else {
  // Fallback implementation
}

// Check API method availability
if (navigator.clipboard?.writeText) {
  await navigator.clipboard.writeText(text);
} else {
  document.execCommand('copy'); // old fallback
}
```

**Modernizr** is a library that performs many feature detections automatically."

#### Indepth
Feature detection is the foundation of **progressive enhancement** — building a baseline experience and enhancing it for capable browsers. For CSS features: `@supports` in CSS or `CSS.supports('display', 'grid')` in JS. For APIs that need user permission (geolocation, clipboard), use try/catch rather than `in` check since permissions can be denied even if the API exists.

---

### 108. What is function composition?

"**Function composition** chains multiple functions where each output feeds the next, creating a new combined function. `compose` applies right-to-left, `pipe` applies left-to-right.

```js
const compose = (...fns) => x => fns.reduceRight((v, fn) => fn(v), x);
const pipe = (...fns) => x => fns.reduce((v, fn) => fn(v), x);

const sanitize = s => s.trim();
const lowercase = s => s.toLowerCase();
const capitalize = s => s[0].toUpperCase() + s.slice(1);

const format = pipe(sanitize, lowercase, capitalize);
format('  hElLo WoRlD  '); // 'Hello world'
```

Real-world: React middleware, Redux reducers, Express middleware, RxJS operators all use composition."

#### Indepth
Composition works best with **pure functions** of arity 1 (one argument). For multi-argument functions, use currying first. The `pipe` operator proposal (`|>`) at TC39 would allow native syntax: `'  HELLO  ' |> sanitize |> lowercase |> capitalize`. Libraries: **Ramda** is fully curried and lazy; **lodash/fp** provides functional-style utilities. Function composition is the key pattern in **middleware** architectures.

---

### 109. What is prototype pollution and how to prevent it?

"**Prototype pollution** occurs when an attacker manipulates `Object.prototype`, adding or modifying properties that all objects inherit — affecting the entire application.

```js
// Vulnerable merge function
function merge(target, source) {
  for (let key in source) {
    target[key] = source[key]; // ❌
  }
}

merge({}, JSON.parse('{"__proto__": {"admin": true}}'));
({}).admin; // true — ALL objects now have admin: true!
```

**Prevention**:
```js
// Check before assigning
if (key === '__proto__' || key === 'constructor' || key === 'prototype') continue;

// Use Object.create(null) for dictionaries
const safeCache = Object.create(null);

// Use Map instead of plain object for dynamic keys
const map = new Map();
```"

#### Indepth
Real-world CVEs: `lodash.merge` had prototype pollution vulnerabilities. `npm audit` found hundreds of packages affected. Prevention strategies: 1) Validate/sanitize input before merging. 2) Use `hasOwnProperty` check: `Object.prototype.hasOwnProperty.call(target, key)`. 3) Use `Object.create(null)` for data dictionaries. 4) Use `Object.freeze(Object.prototype)` in dev to detect mutations. The `__proto__` key in `JSON.parse` IS safe by default — `JSON.parse` doesn't set `[[Prototype]]`, only string-keyed properties.

---

### 110. What is the difference between `apply()`, `call()`, and `bind()`?

"All three set the `this` context for a function call:

| Method | Invocation | Arg format | Returns |
|---|---|---|---|
| `call` | Immediate | Spread args | Result |
| `apply` | Immediate | Args array | Result |
| `bind` | Deferred | Spread args (partial) | New function |

```js
function introduce(greeting, punctuation) {
  return `${greeting}, I'm ${this.name}${punctuation}`;
}
const ctx = { name: 'Dhruv' };

introduce.call(ctx, 'Hello', '!');     // immediate, spread
introduce.apply(ctx, ['Hello', '!']); // immediate, array
const fn = introduce.bind(ctx, 'Hey'); // returns new function
fn('.');                                // 'Hey, I'm Dhruv.'
```"

#### Indepth
**`apply` use case**: Math.max with an array: `Math.max.apply(null, [1, 5, 3])` — though spread (`Math.max(...arr)`) is the modern replacement. **`call` use case**: borrowing array methods: `Array.prototype.slice.call(arguments)`. **`bind` use case**: event handlers in class components: `this.handleClick = this.handleClick.bind(this)`. Arrow functions made most `bind` use cases obsolete in modern code.

---

### 111. What is the mark-and-sweep garbage collection strategy?

"**Mark-and-Sweep** is the primary GC algorithm in modern JS engines:

**Phase 1 — Mark**: Starting from **GC roots** (stack variables, globals, closures), the GC traverses the object graph and marks every **reachable** object.

**Phase 2 — Sweep**: All **unmarked** (unreachable) objects are freed. Their memory is reclaimed and returned to the memory pool.

```
GC Root → objA → objB → objC  (all marked, not collected)
                  ↳ objD      (marked, not collected)
orphan             objE       (unreachable → collected)
```"

#### Indepth
V8 uses a **generational** approach: new objects go to **young generation** (Scavenger GC, triggered frequently, fast). Survivors are promoted to **old generation** (Major GC, Mark-Sweep-Compact, triggered less frequently). **Concurrent marking** runs on background threads while JS executes, reducing pause time. **Incremental marking** splits the mark phase into small slices. **Write barriers** track pointer mutations during concurrent marking.

---

### 112. How does V8 optimize JavaScript?

"V8 uses **JIT (Just-In-Time)** compilation — it starts by interpreting bytecode (Ignition interpreter), and as code runs more, the **TurboFan** compiler takes hot functions and compiles them to optimized machine code.

Key optimizations:
- **Inline caching**: caches method lookups for common object shapes
- **Hidden classes**: structural type system for fast property access
- **Escape analysis**: determines if object can stay on stack
- **Function inlining**: replaces function calls with the function body

```js
// Hidden class optimization
const p1 = { x: 1, y: 2 }; // Shape C1: {x, y}
const p2 = { x: 3, y: 4 }; // Reuses shape C1 — optimized!
const p3 = { y: 1, x: 2 }; // Shape C2: {y, x} — different order = new class!
```"

#### Indepth
**Hidden classes** (V8) / **Shapes** (SpiderMonkey): when objects have the same properties in the same order, they share a hidden class, enabling O(1) property access via offsets instead of hash lookups. Adding properties in different orders or with `delete` fragments hidden classes, causing **deoptimization**. Best practices: 1) Initialize all object properties in the constructor. 2) Maintain consistent property order. 3) Avoid `delete` — set to `undefined` instead. 4) Avoid polymorphic functions — functions called with many different object shapes can't be optimized.

---

### 113. What is the difference between lazy loading and preloading?

"**Lazy loading** defers loading resources until they're actually needed (e.g., when images scroll into view). **Preloading** fetches resources early, before they're needed, to ensure they're ready immediately when required.

```js
// Lazy loading images (native HTML5)
<img src='photo.jpg' loading='lazy' alt='...'>

// Lazy loading JS modules (dynamic import)
button.addEventListener('click', async () => {
  const { Modal } = await import('./Modal.js');
  new Modal().open(); // loaded on demand
});

// Preloading critical resources (HTML)
<link rel='preload' href='hero.jpg' as='image'>
<link rel='prefetch' href='next-page.js'>
```"

#### Indepth
**`preload`** vs **`prefetch`**: `preload` is for **current page** critical resources (high priority, required soon). `prefetch` is for **next navigation** resources (low priority, might need later). Lazy loading with `IntersectionObserver` is more powerful than `loading='lazy'` — you can customize the threshold, add animations, and handle error states. Next.js's `Image` component and React Router's code splitting use these patterns extensively.

---

### 114. What is layout thrashing?

"**Layout thrashing** (forced synchronous layout) occurs when JavaScript alternately **reads** and **writes** to DOM layout properties in a loop, forcing the browser to recalculate layout on every read.

```js
// ❌ Layout thrashing — read after write forces reflow
const boxes = document.querySelectorAll('.box');
boxes.forEach(box => {
  box.style.width = box.offsetWidth * 2 + 'px'; // write then read then write
});

// ✅ Batch reads first, then writes
const widths = [...boxes].map(b => b.offsetWidth);  // all reads
boxes.forEach((box, i) => box.style.width = widths[i] * 2 + 'px'); // all writes
```

FastDOM library and React's virtual DOM batch reads and writes to prevent this."

#### Indepth
Layout-triggering properties: `offsetWidth/Height`, `clientWidth/Height`, `scrollTop/Left`, `getBoundingClientRect()`, computed styles. The browser optimizes by deferring layout calculations. But accessing layout properties mid-script **forces an immediate layout** (reflow). The **RAIL model** (Response, Animation, Idle, Load) says each animation frame should complete in <16ms. Layout thrashing can easily exceed this, causing dropped frames.

---

### 115. What is tail call optimization?

"**Tail call optimization (TCO)** allows recursive functions where the recursive call is the **last operation** to reuse the current stack frame instead of creating a new one — preventing stack overflow.

```js
// Regular recursion — O(n) stack frames, can overflow
function factorial(n) {
  if (n <= 1) return 1;
  return n * factorial(n - 1); // NOT a tail call — multiply happens after
}

// Tail recursive — TCO eligible
function factorial(n, acc = 1) {
  if (n <= 1) return acc;
  return factorial(n - 1, n * acc); // tail call — nothing after recursive call
}
```

TCO is mandated by ES6 spec but only **Safari** fully implements it. For other engines, use **trampolining**."

#### Indepth
**Trampolining** simulates TCO without engine support:
```js
function trampoline(fn) {
  return (...args) => {
    let result = fn(...args);
    while (typeof result === 'function') result = result();
    return result;
  };
}

const factorial = trampoline(function f(n, acc = 1) {
  if (n <= 1) return acc;
  return () => f(n - 1, n * acc); // return thunk instead of calling
});
factorial(100000); // works without stack overflow
```
The thunk pattern returns a function instead of recursing, turning deep recursion into iteration driven by a while loop.
