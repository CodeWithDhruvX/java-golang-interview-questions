# 🟣 **JavaScript Proxy, Reflect, Generators & Web Workers — Staff / Principal Level**

> **Target Companies:** Deloitte, Accenture Tech Centers (Service-Based) + Microsoft, Atlassian, Adobe, Freshworks, Hasura, Browserstack (Product-Based — Staff/Principal Engineer)

---

### 131. What are traps in JavaScript Proxy?

"**Traps** are method hooks on a Proxy handler that intercept fundamental JavaScript operations. Each trap corresponds to an **internal method** like `[[Get]]`, `[[Set]]`, `[[Has]]` etc.

```js
const handler = {
  get(target, prop, receiver) { /* intercept property reads */ },
  set(target, prop, value, receiver) { /* intercept property writes */ },
  has(target, prop) { /* intercept 'in' operator */ },
  deleteProperty(target, prop) { /* intercept 'delete' */ },
  apply(target, thisArg, args) { /* intercept function calls */ },
  construct(target, args, newTarget) { /* intercept 'new' */ },
  ownKeys(target) { /* intercept Object.keys(), for..in */ },
  getOwnPropertyDescriptor(target, prop) { /* intercept descriptor access */ },
};

const proxy = new Proxy(target, handler);
```"

#### Indepth
The 13 Proxy traps map exactly to the 13 **essential internal methods** in the ECMAScript spec. Using `Reflect.X()` in each trap calls the default behavior — this is critical for maintaining the **invariants** (rules) that proxies must follow. For example, the `get` trap cannot return a different value for a `non-writable, non-configurable` data property — violating this throws a `TypeError`. Always pair `Proxy` with `Reflect` for correct behavior.

---

### 132. How can Proxy be used to validate object values?

"Create a type-safe object using Proxy's `set` trap:

```js
function createTypedObject(schema) {
  return new Proxy({}, {
    set(target, prop, value) {
      const expectedType = schema[prop];
      if (!expectedType) throw new Error(`Unknown property: ${prop}`);
      if (typeof value !== expectedType) {
        throw new TypeError(`${prop} must be ${expectedType}, got ${typeof value}`);
      }
      return Reflect.set(target, prop, value);
    },
    get(target, prop) {
      if (!(prop in target) && prop in schema) return undefined;
      return Reflect.get(target, prop);
    }
  });
}

const user = createTypedObject({ name: 'string', age: 'number' });
user.name = 'Dhruv';  // ✅
user.age = 'old';     // ❌ TypeError: age must be number
user.email = 'x';     // ❌ Error: Unknown property
```"

#### Indepth
This pattern demonstrates how **Vue 3's reactivity** works internally — but Vue also tracks dependents to trigger re-renders. Proxies can validate, log, react, or transform every interaction. A production validator would check not just typeof but: required fields, min/max for numbers, regex for strings, enum values. Libraries like `zod` and `yup` achieve similar validation at parse-time rather than runtime, which is more efficient.

---

### 133. How can you use Proxy for access logging?

"A transparent logging proxy that records all property accesses:

```js
function withLogging(target, label = 'Object') {
  return new Proxy(target, {
    get(t, prop, receiver) {
      const value = Reflect.get(t, prop, receiver);
      if (typeof value === 'function') {
        // Wrap methods to log calls
        return function(...args) {
          console.log(`[${label}] ${String(prop)}(${args.map(JSON.stringify).join(', ')})`);
          const result = value.apply(this === receiver ? t : this, args);
          console.log(`[${label}] → ${JSON.stringify(result)}`);
          return result;
        };
      }
      console.log(`[${label}] GET .${String(prop)} → ${JSON.stringify(value)}`);
      return value;
    },
    set(t, prop, value) {
      console.log(`[${label}] SET .${String(prop)} = ${JSON.stringify(value)}`);
      return Reflect.set(t, prop, value);
    }
  });
}

const user = withLogging({ name: 'Dhruv', greet() { return `Hi, ${this.name}`; } }, 'User');
user.name;     // logs: [User] GET .name → "Dhruv"
user.greet();  // logs: [User] greet() ... [User] → "Hi, Dhruv"
```"

#### Indepth
This pattern is used in: **debugging proxies** during development, **ORMs** (like Objection.js, MikroORM) that intercept model attribute access, **permission systems** (check ACL on every property access), and **change tracking** (like MobX observables or Immer's draft). The `receiver` parameter in `get`/`set` traps handles proper `this` binding for prototype chain accesses — always pass it to `Reflect` calls.

---

### 134. What is `Reflect.get()` and when to use it?

"`Reflect` provides methods that mirror the 13 Proxy traps, offering a functional API for fundamental object operations.

```js
const obj = { name: 'Dhruv', get fullName() { return `${this.name} Shah`; } };

// Traditional property access
obj.name; // 'Dhruv'

// Reflect.get — same but handles receiver for getters correctly
Reflect.get(obj, 'name');           // 'Dhruv'
Reflect.get(obj, 'fullName', obj);  // 'Dhruv Shah' (correct this)

// Why Reflect matters in Proxy handlers:
const proxy = new Proxy(obj, {
  get(target, prop, receiver) {
    // ✅ Receiver ensures getters use proxy as 'this'
    return Reflect.get(target, prop, receiver);
    // ❌ target[prop] would bypass proxy for getter's 'this'
  }
});
```"

#### Indepth
`Reflect` methods return **success/failure** instead of throwing (for most operations): `Reflect.set(obj, 'x', 1)` returns `true`/`false` vs `Object.defineProperty` which throws on failure. This matches what Proxy trap handlers should return. `Reflect.ownKeys()` vs `Object.keys()`: Reflect.ownKeys returns ALL own keys including non-enumerable and Symbol keys. This makes it the most complete enumeration method available.

---

### 135. What is the purpose of a Web Worker?

"**Web Workers** run JavaScript in a **separate background thread**, keeping the main thread free for UI interactions.

```js
// main.js
const worker = new Worker('worker.js');

worker.postMessage({ data: hugeArray, operation: 'sort' });

worker.onmessage = ({ data }) => {
  console.log('Sorted result:', data);
};

worker.onerror = (err) => console.error(err);

// worker.js (separate file)
self.onmessage = ({ data }) => {
  const { data: arr, operation } = data;
  
  if (operation === 'sort') {
    const sorted = [...arr].sort((a, b) => a - b);
    self.postMessage(sorted);
  }
};
```"

#### Indepth
Workers have **no access** to: `window`, `document`, DOM APIs, `localStorage`. They communicate only via `postMessage` which uses the **Structured Clone Algorithm** (copies data, no shared references). For shared memory: use `SharedArrayBuffer` with `Atomics` for thread-safe operations. Worker types: **Dedicated** (one page), **Shared** (multiple pages), **Service** (background, network proxy). In Node.js, use `worker_threads` module — similar concept but can share memory more easily.

---

### 136. What is SharedArrayBuffer and Atomics?

"`SharedArrayBuffer` provides **shared memory** between the main thread and Web Workers. `Atomics` provides **atomic operations** to safely read/write that shared memory without race conditions.

```js
// Create shared memory (256 bytes)
const sab = new SharedArrayBuffer(1024);
const sharedArray = new Int32Array(sab);

// Main thread
worker.postMessage({ buffer: sab });
sharedArray[0] = 42;

// Worker (receives same sab, same memory)
self.onmessage = ({ data }) => {
  const shared = new Int32Array(data.buffer);
  
  // Atomic operations — thread-safe
  Atomics.add(shared, 0, 1);   // atomic increment
  Atomics.store(shared, 1, 99); // atomic write
  const val = Atomics.load(shared, 0); // atomic read
  
  // Wait/notify for synchronization
  Atomics.wait(shared, 0, 42); // pause until shared[0] !== 42
  Atomics.notify(shared, 0, 1); // wake up one waiter
};
```"

#### Indepth
`SharedArrayBuffer` was disabled in 2018 due to **Spectre vulnerability** (timing side-channels). It was re-enabled in 2020 only for **cross-origin isolated** contexts (requires `Cross-Origin-Opener-Policy: same-origin` and `Cross-Origin-Embedder-Policy: require-corp` headers). Atomics ensures operations like compare-and-swap are indivisible — critical for lock-free data structures. `Atomics.wait()` blocks the thread (not allowed on the main thread, only workers).

---

### 137. What is `Symbol.toPrimitive`?

"`Symbol.toPrimitive` lets you define **custom type conversion** for objects. When JS needs a primitive (for arithmetic, comparison, string concatenation), it calls this method with a hint: `'number'`, `'string'`, or `'default'`.

```js
class Money {
  constructor(amount, currency) {
    this.amount = amount;
    this.currency = currency;
  }

  [Symbol.toPrimitive](hint) {
    if (hint === 'number') return this.amount;
    if (hint === 'string') return `${this.amount} ${this.currency}`;
    return this.amount; // 'default' (used for + operator)
  }
}

const price = new Money(100, 'USD');
+price;              // 100 (number hint)
`${price}`;          // '100 USD' (string hint)
price + 50;          // 150 (default hint → number)
```"

#### Indepth
Without `Symbol.toPrimitive`, JS uses `valueOf()` (for numeric) then `toString()` (for string). With `Symbol.toPrimitive`, it completely overrides both. The `hint` parameter distinguishes contexts: `'number'` for unary `+` and arithmetic; `'string'` for template literals and string methods; `'default'` for `==`, `+` operator with mixed types, `?:`. Libraries like `Decimal.js` use this to make decimal instances behave naturally in numeric contexts.

---

### 138. What are generator functions and how do they work internally?

"A generator function (`function*`) returns a **generator object** that implements both the iterable and iterator protocols. It uses `yield` to pause execution and return values one at a time.

```js
function* fibonacci() {
  let [a, b] = [0, 1];
  while (true) {
    yield a;
    [a, b] = [b, a + b];
  }
}

const fib = fibonacci();
fib.next(); // { value: 0, done: false }
fib.next(); // { value: 1, done: false }
fib.next(); // { value: 1, done: false }
fib.next(); // { value: 2, done: false }

// Get first 10
const first10 = [...Array(10)].map(() => fib.next().value);
```"

#### Indepth
Internally, generators compile to a **state machine** with numbered states. `yield` saves the current state (local variables, position in code) and returns to the caller. `next(value)` resumes — the `value` argument becomes the **result** of the `yield` expression inside the generator. This enables two-way communication. `generator.throw(err)` injects an error at the suspended yield point. Redux Saga uses this to express complex async flows declaratively.

---

### 139. What is the difference between generator and async generator?

"A regular **generator** (`function*`) is synchronous — `yield` produces values immediately.
An **async generator** (`async function*`) can `await` promises AND `yield` values — perfect for async data streams.

```js
// Regular generator — synchronous
function* syncCounter() {
  yield 1; yield 2; yield 3;
}

// Async generator — yields async values
async function* fetchPages(url) {
  let page = 1;
  while (true) {
    const data = await fetch(`${url}?page=${page}`).then(r => r.json());
    if (!data.length) return;
    yield data; // yield one page at a time
    page++;
  }
}

// Consume with for await...of
for await (const page of fetchPages('/api/items')) {
  displayPage(page);
}
```"

#### Indepth
Async generators power streaming APIs in Node.js: `fs.createReadStream()` is an async iterable. `fetch` in Node 18+ returns a `Response` with `body` as a `ReadableStream` which is async iterable. Use cases: **paginated APIs** (fetch page by page without loading all), **live data streams** (SSE, WebSocket), **file processing** (read line by line from large files). `for await...of` on a regular (sync) generator also works — it just wraps each value in a resolved Promise.

---

### 140. How do you implement an event emitter in JavaScript?

"A full EventEmitter implementation:

```js
class EventEmitter {
  #events = new Map();
  #maxListeners = 10;

  on(event, listener) {
    if (!this.#events.has(event)) this.#events.set(event, []);
    const listeners = this.#events.get(event);
    if (listeners.length >= this.#maxListeners) {
      console.warn(`MaxListeners exceeded for event: ${event}`);
    }
    listeners.push(listener);
    return this; // for chaining
  }

  once(event, listener) {
    const wrapper = (...args) => {
      listener(...args);
      this.off(event, wrapper);
    };
    wrapper._original = listener; // preserve original for off()
    return this.on(event, wrapper);
  }

  off(event, listener) {
    const listeners = this.#events.get(event) ?? [];
    this.#events.set(event, listeners.filter(l => 
      l !== listener && l._original !== listener
    ));
    return this;
  }

  emit(event, ...args) {
    const listeners = this.#events.get(event) ?? [];
    listeners.forEach(fn => fn(...args));
    return listeners.length > 0;
  }

  listenerCount(event) {
    return this.#events.get(event)?.length ?? 0;
  }
}
```"

#### Indepth
Node.js's `events.EventEmitter` is the backbone of core modules (`Stream`, `net`, `http`). Key features: `once()` for one-time listeners, `setMaxListeners` to avoid memory leak warnings, `removeAllListeners()` for cleanup. **Memory leak detection**: if you add more than the max listeners (default 10) to a single event, Node.js warns — this usually means you forgot to remove listeners. In browsers, `EventTarget` provides similar functionality natively.

---

### 141. Implement an event emitter / pub-sub in JavaScript.

"A lightweight publish-subscribe pattern (decoupled variant of observer):

```js
class PubSub {
  #subscriptions = new Map();

  subscribe(topic, callback) {
    if (!this.#subscriptions.has(topic)) {
      this.#subscriptions.set(topic, new Set());
    }
    this.#subscriptions.get(topic).add(callback);
    
    // Return unsubscribe function (cleanup pattern)
    return () => this.#subscriptions.get(topic)?.delete(callback);
  }

  publish(topic, data) {
    this.#subscriptions.get(topic)?.forEach(cb => {
      try { cb(data); } catch (err) { console.error(err); }
    });
  }
}

const events = new PubSub();

const unsub = events.subscribe('user:login', (user) => {
  console.log(`${user.name} logged in`);
});

events.publish('user:login', { name: 'Dhruv' }); // logs the event
unsub(); // clean up subscription
```"

#### Indepth
PubSub vs Observer: Observer knows who observes it (direct coupling). PubSub uses an intermediary (event bus) — publisher and subscriber don't know each other (loose coupling). Using `Set` instead of `Array` for subscribers: O(1) deletion vs O(n) for arrays. Returning an **unsubscribe function** (instead of `removeSubscriber(cb)`) is the modern pattern — matching React's `useEffect` cleanup, RxJS `subscription.unsubscribe()`, and browser's `AbortController`.

---

### 142. What is the difference between `for...in` and `for...of`?

"`for...in` iterates over **enumerable property keys** of an object (including inherited ones from prototype chain).
`for...of` iterates over **iterable values** (arrays, strings, Set, Map, generators) using the `Symbol.iterator` protocol.

```js
const arr = [10, 20, 30];
arr.customProp = 'meta';

for (const key in arr) {
  console.log(key); // '0', '1', '2', 'customProp' ← includes ALL enumerable keys
}

for (const val of arr) {
  console.log(val); // 10, 20, 30 ← only values, no customProp
}

// for..in on plain objects (correct use)
const obj = { a: 1, b: 2 };
for (const key in obj) {
  if (Object.hasOwn(obj, key)) { /* filter inherited */ }
}
```"

#### Indepth
**`for...in` on arrays is an anti-pattern** — it includes non-index properties and doesn't guarantee numeric order. Always use `for...of` or `forEach` for arrays. For objects: `for...in` with `hasOwnProperty` check, or better: `Object.keys(obj)` (enumerable own), `Object.values()`, `Object.entries()`. `for...of` with a plain object throws `TypeError` — objects are not iterable by default. Make them iterable by implementing `[Symbol.iterator]`.

---

### 143. What is the Temporal API?

"The **Temporal API** (TC39 Stage 3) is the long-awaited successor to `Date` that fixes its fundamental design flaws:

```js
// Current Date API problems:
new Date('2026-01-01');           // Parsed as UTC
new Date('01/01/2026');           // Parsed as LOCAL — ambiguous!
new Date().getMonth();            // 0-indexed (January = 0) — confusing

// Temporal API — precise, immutable, timezone-aware
const today = Temporal.Now.plainDateISO();  // '2026-02-20'
const meeting = Temporal.ZonedDateTime.from({
  year: 2026, month: 3, day: 15,
  hour: 10, minute: 0,
  timeZone: 'Asia/Kolkata',
});

// Arithmetic
const nextWeek = today.add({ weeks: 1 });
const diff = meeting.until(Temporal.Now.zonedDateTimeISO('Asia/Kolkata'));
diff.hours; // exact hours until meeting
```"

#### Indepth
Temporal introduces **5 date/time types**: `PlainDate` (no time/zone), `PlainTime`, `PlainDateTime`, `ZonedDateTime` (full timezone awareness), `Instant` (nanosecond precision UTC). All are **immutable**. BREAKING: Temporal uses **1-indexed months** (January = 1). The `@js-temporal/polyfill` package is available for early adoption. `date-fns` and `Luxon` fill the gap today. Node.js 22+ has experimental Temporal support.

---

### 144. What is `import.meta` in ES Modules?

"`import.meta` is an object that provides **context-specific information** about the current module. Its properties depend on the runtime environment.

```js
// Browser (ES module loaded via <script type='module'>)
import.meta.url; // 'https://example.com/path/to/module.js'

// Resolving relative paths
const url = new URL('./config.json', import.meta.url);
fetch(url).then(r => r.json());

// Vite-specific: environment variables
import.meta.env.MODE;     // 'development' or 'production'
import.meta.env.VITE_API; // custom env var

// Node.js (ESM)
import.meta.url;   // 'file:///path/to/module.js'
import.meta.dirname; // '/path/to' (Node 20+)
import.meta.filename; // '/path/to/module.js' (Node 20+)
```"

#### Indepth
`import.meta.url` replaces the CommonJS `__filename` and `__dirname` in ESM. Previously, ESM modules had no way to get their file path. Pattern: `const __dirname = new URL('.', import.meta.url).pathname`. Bundlers like Vite and Webpack use `import.meta.env` for environment variable injection — these are statically replaced at build time (similar to `process.env` in Node). `import.meta.hot` in Vite enables HMR (Hot Module Replacement) logic.

---

### 145. How do you implement a tiny reactive store in JavaScript?

"A minimal reactive state management store using Proxy:

```js
function createStore(initialState) {
  const subscribers = new Set();
  
  const state = new Proxy(initialState, {
    set(target, prop, value, receiver) {
      const result = Reflect.set(target, prop, value, receiver);
      subscribers.forEach(fn => fn({ ...target })); // notify all
      return result;
    }
  });

  return {
    state,
    subscribe(fn) {
      subscribers.add(fn);
      fn({ ...state }); // immediate initial notification
      return () => subscribers.delete(fn); // unsubscribe
    },
    getSnapshot() { return { ...state }; },
  };
}

// Usage
const store = createStore({ count: 0, user: null });

const unsub = store.subscribe(state => {
  console.log('State changed:', state);
});

store.state.count = 1; // triggers all subscribers
store.state.user = { name: 'Dhruv' }; // triggers again
unsub(); // stop listening
```"

#### Indepth
This is conceptually how **MobX** and **Vue 3 ref/reactive** work. MobX adds: dependency tracking (only re-runs observers that actually READ the changed value), batch updates, computed values, and reactions. Zustand (React) uses a subscription model without proxies. Jotai and Recoil use React's concurrent mode features. The Proxy approach has a footgun: it notifies on EVERY property change — batching with `queueMicrotask` or a scheduler prevents excessive re-renders.
