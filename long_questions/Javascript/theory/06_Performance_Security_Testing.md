# 🔴 **JavaScript Performance, Security & Testing — Expert Level**

> **Target Companies:** Capgemini, NTT Data, DXC Technology (Service-Based Senior) + Netflix, Uber, Swiggy, CRED, Groww, Zepto (Product-Based Senior/Staff Engineer)

---

### 116. What is the purpose of the `arguments` object?

"The `arguments` object is an **array-like object** (not a real array) available inside every regular function that contains all passed arguments — even those beyond the declared parameters.

```js
function sum() {
  let total = 0;
  for (let i = 0; i < arguments.length; i++) {
    total += arguments[i];
  }
  return total;
}
sum(1, 2, 3, 4); // 10

// Modern replacement: rest parameters
function sum(...nums) {
  return nums.reduce((a, b) => a + b, 0);
}
```

`arguments` does NOT exist in arrow functions. I avoid it in new code and prefer rest parameters."

#### Indepth
`arguments` is array-like (has `.length` and indexed access) but lacks array methods like `map`, `filter`. Convert with: `Array.from(arguments)` or `[...arguments]`. In non-strict mode, `arguments` mirrors the named parameters — modifying `arguments[0]` also modifies `param0` and vice versa (aliasing). In strict mode, this aliasing is disabled. `arguments.callee` (references the current function) is deprecated — use named function expressions instead.

---

### 117. What is the nullish coalescing operator (`??`)?

"The **nullish coalescing operator** (`??`) returns the right-hand side only when the left side is `null` or `undefined` — NOT for other falsy values like `0`, `''`, or `false`.

```js
// Problem with || (short-circuits on ANY falsy value)
const count = userCount || 10; // ❌ 0 uses fallback 10 even though 0 is valid!

// Solution with ?? (only nullish)
const count = userCount ?? 10; // ✅ 0 stays 0, only null/undefined → 10

// Real world examples
const name = user.name ?? 'Anonymous';
const port = config.port ?? 3000;
const items = response.data ?? [];
```"

#### Indepth
`??` (ES2020) vs `||`: `||` treats ANY falsy value as missing; `??` only treats `null`/`undefined` as missing. This matters enormously for: port numbers (`0` is falsy), boolean flags (`false` means disabled, not missing), and counts. **Logical nullish assignment** (`??=`): `a ??= b` assigns `b` to `a` only if `a` is null/undefined. Combined: `user.settings ??= {}` initializes settings only if not already set.

---

### 118. What are logical assignment operators?

"ES2021 logical assignment operators combine a logical operation with assignment:

```js
// &&= — assign only if left is TRUTHY
let a = 1;
a &&= 5;    // a = 5 (because a was truthy)
let b = 0;
b &&= 5;    // b = 0 (no assignment because b was falsy)

// ||= — assign only if left is FALSY
let c = null;
c ||= 'default';  // c = 'default'
let d = 'value';
d ||= 'default';  // d = 'value' (no change)

// ??= — assign only if left is NULL/UNDEFINED
let e;
e ??= 'initial'; // e = 'initial'
e ??= 'again';   // e = 'initial' (no change — already defined)
```"

#### Indepth
These operators **short-circuit** — the right side is only evaluated if needed. This is important when the right side has side effects: `obj.count ??= heavyComputation()` — `heavyComputation()` only runs if `obj.count` is null/undefined. These operators are particularly useful for **initializing nested config objects**: `config.db ??= {}; config.db.pool ??= 5;`. TypeScript fully supports them from v4.0.

---

### 119. What is `WeakRef`?

"`WeakRef` holds a **weak reference** to an object — it doesn't prevent the object from being garbage collected. Use it to cache things you'd like to keep IF they exist, but won't fight the GC to keep alive.

```js
let cache = new WeakRef(bigDataObject);

function getData() {
  const data = cache.deref(); // deref() returns the object or undefined
  if (data) {
    return data; // still in memory
  }
  // Object was GC'd, recreate it
  const newData = loadData();
  cache = new WeakRef(newData);
  return newData;
}
```"

#### Indepth
`WeakRef.prototype.deref()` returns the referenced object or `undefined` if GC'd. The timing of GC is non-deterministic — you can't rely on when `deref()` returns `undefined`. `WeakRef` is best paired with `FinalizationRegistry` which executes a callback after an object is collected. Primary use cases: caching expensive computed objects (images, buffers), observer patterns where observed objects can die naturally. Avoid using `WeakRef` as a substitute for proper lifetime management.

---

### 120. How does `Promise.prototype.finally()` work?

"`.finally()` runs a callback when a promise settles (either resolves or rejects), useful for cleanup code that must run regardless of outcome.

```js
async function fetchData(url) {
  showLoading(true);
  try {
    const data = await fetch(url).then(r => r.json());
    updateUI(data);
    return data;
  } catch (err) {
    showError(err.message);
    throw err; // re-throw so caller knows it failed
  } finally {
    showLoading(false); // ALWAYS hide loader
  }
}

// With promise chain
fetch('/api/data')
  .then(processData)
  .catch(handleError)
  .finally(() => hideSpinner());
```"

#### Indepth
Key behavior: `.finally()` **does not receive any argument** (no value or reason — it's cleanup, not result processing). It **passes through** the resolved value or rejection reason to the next handler — unlike `.then(fn, fn)` which could swallow values. `finally` with a `return` or `throw` **overrides** the original result. `finally` with no explicit return lets the original result flow through. This makes it safe for cleanup without accidentally changing the promise result.

---

### 121. How do you cancel a fetch request?

"Use the **AbortController** API to cancel fetch requests:

```js
const controller = new AbortController();
const { signal } = controller;

try {
  const response = await fetch('/api/data', { signal });
  const data = await response.json();
  updateUI(data);
} catch (err) {
  if (err.name === 'AbortError') {
    console.log('Fetch was cancelled'); // expected
  } else {
    throw err; // unexpected error
  }
}

// Cancel the request (e.g., user navigated away, or new search started)
controller.abort();
```

In React: call `controller.abort()` in the `useEffect` cleanup function."

#### Indepth
`AbortController` sends a signal that `fetch` (and other async APIs) listen to. Aborting triggers an `AbortError` (a `DOMException`) in the fetch promise. One controller can abort multiple requests: `Promise.all([fetch(url1, { signal }), fetch(url2, { signal })])` — one `abort()` cancels both. AbortController is also supported by `addEventListener` (auto-remove listeners on abort), making it a powerful lifecycle management tool.

---

### 122. What are microtasks vs macrotasks with examples?

"The event loop processes tasks in strict priority order:

**Synchronous code** (always first):
```js
console.log('1 — sync');
```

**Microtasks** (drained COMPLETELY before next macrotask):
```js
Promise.resolve().then(() => console.log('2 — microtask'));
queueMicrotask(() => console.log('3 — microtask'));
```

**Macrotasks** (one per event loop iteration):
```js
setTimeout(() => console.log('4 — macrotask'), 0);
setInterval(() => console.log('5 — macrotask'), 0);
```

**Output**: `1 — sync` → `2 — microtask` → `3 — microtask` → `4 — macrotask`"

#### Indepth
Why this matters: If you queue microtasks in a loop (e.g., a Promise that resolves and queues another microtask), you can **starve macrotasks** — the UI never updates. This is the async equivalent of an infinite loop. `MutationObserver` callbacks are microtasks — they run synchronously after DOM changes but before the browser paints. In Node.js, `process.nextTick()` runs before microtasks — it has even higher priority. `setImmediate()` is a macrotask that runs after I/O callbacks.

---

### 123. How to implement a retry mechanism for failed async requests?

"Implement exponential backoff for robust retry logic:

```js
async function fetchWithRetry(url, options = {}, maxRetries = 3) {
  for (let attempt = 1; attempt <= maxRetries; attempt++) {
    try {
      const response = await fetch(url, options);
      if (!response.ok) throw new Error(`HTTP ${response.status}`);
      return await response.json();
    } catch (err) {
      if (attempt === maxRetries) throw err; // final attempt, give up
      
      const delay = Math.pow(2, attempt) * 1000; // 2s, 4s, 8s
      const jitter = Math.random() * 1000;       // add randomness
      await new Promise(resolve => setTimeout(resolve, delay + jitter));
      console.log(`Attempt ${attempt} failed, retrying in ${delay}ms...`);
    }
  }
}
```"

#### Indepth
**Exponential backoff with jitter** (randomness) is crucial for distributed systems — if all clients retry at the same time (without jitter), they create synchronized thundering herds that overwhelm the server. Libraries like `axios-retry` and `p-retry` implement sophisticated retry logic with: max retries, retry conditions (only retry on 5xx, not 4xx), delay strategies, and `onRetry` callbacks. AWS, GCP, and Stripe recommend exponential backoff for all API clients.

---

### 124. Implement a deep equality checker (`deepEqual()`).

"A `deepEqual` function that recursively compares objects and arrays:

```js
function deepEqual(a, b) {
  // Primitives and same reference
  if (a === b) return true;
  
  // Handle null (typeof null === 'object')
  if (a === null || b === null) return false;
  
  // Type mismatch
  if (typeof a !== typeof b) return false;
  if (typeof a !== 'object') return a === b;
  
  // Handle arrays
  if (Array.isArray(a) !== Array.isArray(b)) return false;
  
  const keysA = Object.keys(a);
  const keysB = Object.keys(b);
  if (keysA.length !== keysB.length) return false;
  
  return keysA.every(key => deepEqual(a[key], b[key]));
}

deepEqual({ a: 1, b: [2, 3] }, { a: 1, b: [2, 3] }); // true
deepEqual({ a: 1 }, { a: 2 }); // false
```"

#### Indepth
Production-grade `deepEqual` must handle: `NaN === NaN` (use `Object.is`), `Date` objects (compare `.getTime()`), `RegExp` (compare `.source` and `.flags`), `Map`/`Set` (custom comparison), circular references (track visited objects with a `WeakSet`), `Symbol` keys (`Object.getOwnPropertySymbols`). Libraries: `lodash.isEqual` handles all these cases. Fast-deep-equal npm package benchmarks at 4x faster than Lodash for common cases.

---

### 125. Write a polyfill for `Promise.all`.

"Implementing `Promise.all` from scratch:

```js
Promise.myAll = function(promises) {
  return new Promise((resolve, reject) => {
    const results = [];
    let remaining = promises.length;
    
    if (remaining === 0) {
      resolve(results);
      return;
    }
    
    promises.forEach((promise, index) => {
      Promise.resolve(promise).then(value => {
        results[index] = value;
        remaining--;
        if (remaining === 0) resolve(results);
      }).catch(reject); // first rejection rejects the whole thing
    });
  });
};

// Test
Promise.myAll([
  Promise.resolve(1),
  fetch('/api/user').then(r => r.json()),
  42 // non-promise values wrapped with Promise.resolve
]).then(console.log); // [1, userData, 42]
```"

#### Indepth
Key considerations: 1) Must handle non-promise values (wrap with `Promise.resolve`). 2) Results must maintain **input order**, not resolution order — that's why we use `results[index]`, not `results.push`. 3) Empty array resolves with `[]` synchronously (edge case). 4) First rejection immediately rejects, other promises continue but their results are ignored. This interviewable implementation shows understanding of Promise mechanics and async coordination patterns.

---

### 126. What is the Revealing Module Pattern?

"The **Revealing Module Pattern** is a variant of the module pattern that defines all functions privately and only returns an object **mapping private names to public interfaces**:

```js
const apiService = (() => {
  // Private implementation
  const BASE_URL = 'https://api.example.com';
  let authToken = null;

  function setToken(token) { authToken = token; }
  function makeRequest(endpoint, options = {}) {
    return fetch(`${BASE_URL}${endpoint}`, {
      headers: { Authorization: `Bearer ${authToken}` },
      ...options,
    });
  }
  function getUsers() { return makeRequest('/users'); }
  function getUser(id) { return makeRequest(`/users/${id}`); }

  // Reveal only what's public
  return { setToken, getUsers, getUser };
})();
```"

#### Indepth
The advantage over the basic module pattern: ALL logic is written consistently as private functions. The public API is a **thin mapping layer**, making it easy to see exactly what's exposed without scanning for method definitions. The disadvantage: revealed functions hold references to their original private functions — if you override a public method from outside, internal calls still use the private version (unlike regular module pattern methods that use `this`).

---

### 127. What is the Proxy pattern in modern JS frameworks?

"The **Proxy** object intercepts and customizes fundamental object operations (property access, assignment, function invocation).

```js
const handler = {
  get(target, prop, receiver) {
    console.log(`Getting: ${prop}`);
    return Reflect.get(target, prop, receiver);
  },
  set(target, prop, value) {
    if (typeof value !== 'number') throw new TypeError(`${prop} must be a number`);
    return Reflect.set(target, prop, value);
  }
};

const state = new Proxy({ count: 0 }, handler);
state.count;     // logs: "Getting: count" → returns 0
state.count = 5; // sets to 5
state.count = 'x'; // throws TypeError
```

**Vue 3's reactivity system** is built on Proxy — when you set `state.count++`, Vue detects it via the `set` trap and re-renders affected components."

#### Indepth
13 proxy traps: `get`, `set`, `has`, `deleteProperty`, `apply`, `construct`, `getPrototypeOf`, `setPrototypeOf`, `isExtensible`, `preventExtensions`, `getOwnPropertyDescriptor`, `defineProperty`, `ownKeys`. **Reflect** mirrors Proxy traps and provides the default behavior — always use `Reflect` in handlers to maintain invariants. Proxy is NOT polyfillable — it requires native support. This is why Vue 3 dropped IE11 support (Proxy requires modern browsers).

---

### 128. What's the difference between unit, integration, and e2e testing?

"**Unit tests** test individual functions/components in isolation, with all dependencies mocked:
```js
test('add function', () => {
  expect(add(2, 3)).toBe(5);
});
```

**Integration tests** test multiple units working together, with some real dependencies:
```js
test('UserService.create', async () => {
  const user = await UserService.create({ name: 'Dhruv' });
  expect(user.id).toBeDefined(); // real DB, real service
});
```

**E2E tests** test the entire application through a real browser:
```js
// Playwright / Cypress
await page.click('button#submit');
await expect(page.locator('.success-msg')).toBeVisible();
```"

#### Indepth
The **test pyramid**: many unit tests (fast, cheap), fewer integration tests (slower, more setup), fewest E2E tests (slowest, most brittle). In JavaScript: **Jest** for unit/integration, **Playwright** or **Cypress** for E2E. **Testing Library** (`@testing-library/react`) philosophy: test the DOM as users see it, not implementation details. **Storybook** enables isolated component testing. The test pyramid ratio typically: 70% unit / 20% integration / 10% E2E.

---

### 129. What is mocking in JavaScript testing?

"**Mocking** replaces real dependencies (API calls, databases, modules) with controlled fake versions during testing:

```js
// Jest mock modules
jest.mock('./fetchUser', () => ({
  fetchUser: jest.fn().mockResolvedValue({ id: 1, name: 'Dhruv' })
}));

// Mock a specific method
jest.spyOn(UserService, 'create').mockResolvedValue({ id: 42 });

// Mock fetch globally
global.fetch = jest.fn().mockResolvedValue({
  ok: true,
  json: () => Promise.resolve({ data: 'test' })
});

// After test
jest.restoreAllMocks(); // clean up
```"

#### Indepth
**Mock vs Stub vs Spy**: Mock — verifiable fake (you check it was called correctly). Stub — simple replacement with hardcoded return. Spy — wraps real implementation, records calls. `jest.fn()` creates a mock function; `jest.spyOn` creates a spy. **MSW (Mock Service Worker)** is the modern approach for API mocking — it intercepts real HTTP requests at the network level, giving more realistic tests. **Fake timers**: `jest.useFakeTimers()` controls `setTimeout`/`setInterval` for deterministic async tests.

---

### 130. What are best practices for safely handling user input?

"Input safety is a multi-layer concern:

**1. Validate at every boundary**:
```js
// Whitelist validation
function validateAge(input) {
  const age = Number(input);
  if (!Number.isInteger(age) || age < 0 || age > 150) {
    throw new ValidationError('Invalid age');
  }
  return age;
}
```

**2. Sanitize before rendering**:
```js
import DOMPurify from 'dompurify';
element.innerHTML = DOMPurify.sanitize(userGeneratedHTML);
```

**3. Parameterize database queries** (backend):
```js
// Instead of string concatenation:
db.query('SELECT * FROM users WHERE id = $1', [userId]);
```

**4. Encode for output context** (URL, HTML, JS):
```js
encodeURIComponent(userInput); // for URLs
```"

#### Indepth
**Defense in depth**: validate input (right type/format), sanitize (remove dangerous parts), encode output (context-specific escaping). The **OWASP Top 10** lists injection attacks as #1. Client-side validation improves UX but is NOT a security measure — always validate on the server too. For rich text editors (like `<quill>` or `<tiptap>`), always sanitize the output HTML with DOMPurify configured for your allowed tags before saving to DB and before rendering.
