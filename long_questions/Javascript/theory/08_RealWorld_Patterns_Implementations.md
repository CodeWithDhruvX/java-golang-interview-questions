# 🟣 **JavaScript Real-World Patterns & Custom Implementations — Expert / Architect Level**

> **Target Companies:** Goldman Sachs, JP Morgan (Service-Based BFSI) + Zerodha, Groww, CoinDCX, Razorpay, Stripe India (Product-Based Fintech + Senior/Staff/Principal)

---

### 146. How to implement a queue with Promises?

"A **Promise Queue** processes async tasks one at a time, ensuring sequential execution:

```js
class PromiseQueue {
  #queue = [];
  #isProcessing = false;

  add(task) {
    return new Promise((resolve, reject) => {
      this.#queue.push({ task, resolve, reject });
      this.#process();
    });
  }

  async #process() {
    if (this.#isProcessing || this.#queue.length === 0) return;
    this.#isProcessing = true;

    const { task, resolve, reject } = this.#queue.shift();
    try {
      resolve(await task());
    } catch (err) {
      reject(err);
    } finally {
      this.#isProcessing = false;
      this.#process(); // process next item
    }
  }
}

// Usage
const q = new PromiseQueue();
q.add(() => fetch('/api/a').then(r => r.json())); // runs first
q.add(() => fetch('/api/b').then(r => r.json())); // runs after a completes
```"

#### Indepth
This pattern is used when: 1) API rate limiting requires sequential requests. 2) State mutations must be serialized (like a DB write queue). 3) Animations must run one after another. For **concurrent queues** with a max concurrency limit:
```js
class ConcurrentQueue {
  constructor(concurrency) { this.concurrency = concurrency; this.running = 0; this.queue = []; }
  add(task) { /* runs task if running < concurrency, else queues */ }
}
```
Libraries like `p-queue` provide production-ready implementations with priority, pause/resume, and events.

---

### 147. How to limit concurrent fetch calls?

"Use a semaphore to control concurrency:

```js
async function runWithConcurrency(tasks, limit = 5) {
  const results = [];
  const executing = new Set();

  for (const task of tasks) {
    const promise = task().then(result => {
      executing.delete(promise);
      return result;
    });

    executing.add(promise);
    results.push(promise);

    if (executing.size >= limit) {
      await Promise.race(executing); // wait for one to finish
    }
  }

  return Promise.all(results); // get all in order
}

// Usage: fetch 100 items but max 5 at once
const urls = Array.from({ length: 100 }, (_, i) => `/api/item/${i}`);
const tasks = urls.map(url => () => fetch(url).then(r => r.json()));
const data = await runWithConcurrency(tasks, 5);
```"

#### Indepth
**Why limit concurrency?** Browsers limit simultaneous connections per origin (Chrome: 6 per domain). Too many parallel requests trigger **rate limiting** on APIs. Uncontrolled parallelism exhausts memory. `p-limit` is the production-grade npm package with 1.8M weekly downloads. In Node.js for database operations: connection pool size limits effective concurrency naturally — no point spawning 100 DB queries if the pool only has 10 connections.

---

### 148. How to time out a Promise after N milliseconds?

"Combine `Promise.race` with a rejection timer:

```js
function withTimeout(promise, ms, message = 'Operation timed out') {
  const timeoutPromise = new Promise((_, reject) => {
    const id = setTimeout(() => reject(new Error(message)), ms);
    // Clean up timer if main promise settles first (prevents timer leak)
    promise.finally(() => clearTimeout(id));
  });
  return Promise.race([promise, timeoutPromise]);
}

// Usage
try {
  const data = await withTimeout(
    fetch('/api/slow-endpoint').then(r => r.json()),
    5000, // 5 second timeout
    'Request timed out after 5s'
  );
} catch (err) {
  if (err.message.includes('timed out')) {
    showTimeoutError();
  }
}
```"

#### Indepth
The `finally` cleanup is critical — without it, the `setTimeout` fires after the test completes, causing warnings and potential memory leaks in test environments. For production, pair with `AbortController`:
```js
const controller = new AbortController();
const timeout = setTimeout(() => controller.abort(), 5000);
fetch(url, { signal: controller.signal }).finally(() => clearTimeout(timeout));
```
`AbortController` actually cancels the network request (not just ignores the response), which saves bandwidth. This is the correct production pattern.

---

### 149. What is a deferred object? Can you simulate one?

"A **deferred** exposes a Promise's `resolve` and `reject` externally, allowing external code to settle the promise.

```js
function createDeferred() {
  let resolve, reject;
  const promise = new Promise((res, rej) => {
    resolve = res;
    reject = rej;
  });
  return { promise, resolve, reject };
}

// Usage
const deferred = createDeferred();

// Some external code will resolve it later
button.addEventListener('click', () => deferred.resolve('clicked!'));
setTimeout(() => deferred.reject(new Error('timeout')), 5000);

const result = await deferred.promise; // awaits until resolved/rejected
```

Modern equivalent: `Promise.withResolvers()` (ES2024) does exactly this natively."

#### Indepth
`Promise.withResolvers()` (ES2024): `const { promise, resolve, reject } = Promise.withResolvers()` — same pattern, built-in. Use cases for deferred: 1) **Testing** — resolve a promise from outside to simulate callbacks. 2) **User confirmation modals** — return a deferred, resolve when user confirms, reject on cancel. 3) **Coordination** between unrelated code paths. The Deferred pattern is considered an **antipattern** when overused — most async flows fit naturally into `new Promise(executor)`. Use it sparingly.

---

### 150. How to run async functions sequentially?

"Several approaches for sequential async execution:

```js
const tasks = [fetchA, fetchB, fetchC]; // each returns a Promise

// 1. for...of with await (cleanest)
for (const task of tasks) {
  await task();
}

// 2. reduce (functional style)
await tasks.reduce(async (prevPromise, task) => {
  await prevPromise;
  return task();
}, Promise.resolve());

// 3. Recursive approach
async function runSequentially(tasks, index = 0) {
  if (index >= tasks.length) return;
  await tasks[index]();
  return runSequentially(tasks, index + 1);
}

// 4. For-loop (explicit)
const results = [];
for (const task of tasks) {
  results.push(await task()); // each awaited before next
}
```"

#### Indepth
A common **mistake**: `tasks.forEach(async task => await task())` — `forEach` ignores the returned promises! Each iteration fires without waiting. Use `for...of` or `Array.prototype.reduce`. The `reduce` approach is elegant but has a quirk: `reduce` calls the callback synchronously for all elements, building up a chain of Promises before any run. For sequential with results: `const results = []; for (const t of tasks) results.push(await t())`. For optimal performance, use parallel when order doesn't matter: `Promise.all(tasks.map(t => t()))`.

---

### 151. What is mixin pattern in JavaScript?

"**Mixins** are a way to add reusable behaviors to classes **without inheritance**:

```js
// Mixin factories
const Serializable = (Base) => class extends Base {
  serialize() { return JSON.stringify(this); }
  static deserialize(json) { return Object.assign(new this(), JSON.parse(json)); }
};

const Validatable = (Base) => class extends Base {
  validate() {
    const errors = [];
    for (const [field, rule] of Object.entries(this.constructor.rules ?? {})) {
      if (!rule(this[field])) errors.push(`Invalid: ${field}`);
    }
    return errors;
  }
};

// Composing mixins
class User extends Serializable(Validatable(class {})) {
  static rules = {
    name: v => typeof v === 'string' && v.length > 0,
    age: v => Number.isInteger(v) && v > 0,
  };
  constructor(name, age) { super(); this.name = name; this.age = age; }
}

const u = new User('Dhruv', 28);
u.validate();   // []
u.serialize();  // '{"name":"Dhruv","age":28}'
```"

#### Indepth
Mixins solve the **diamond problem** of multiple inheritance — JS classes can only extend one base, but you can compose unlimited mixins. This is the pattern behind React's HOC pattern and Vue's mixin/Composition API. The **functional mixin pattern** (without class): `Object.assign(target, source)` copies methods directly — simpler but no `instanceof` support. TypeScript supports mixins via the same `(Base: Constructor) => class extends Base` pattern with proper typing.

---

### 152. How does dirty checking work?

"**Dirty checking** is an approach to detect changes by **comparing current state to a known previous snapshot**. AngularJS 1.x used this extensively.

```js
class DirtyChecker {
  #watched = new Map(); // element → { lastKnown, callback }

  watch(key, getValue, callback) {
    this.#watched.set(key, { lastKnown: getValue(), getValue, callback });
  }

  // Called periodically (e.g., requestAnimationFrame or $digest)
  check() {
    for (const [key, entry] of this.#watched) {
      const current = entry.getValue();
      if (!Object.is(current, entry.lastKnown)) {
        entry.callback(current, entry.lastKnown);
        entry.lastKnown = current;
      }
    }
  }
}

const checker = new DirtyChecker();
checker.watch('count', () => store.count, (n, o) => console.log(`${o} → ${n}`));
setInterval(() => checker.check(), 50); // $digest cycle equivalent
```"

#### Indepth
AngularJS ran the `$digest` cycle on every user event, HTTP response, and timer. If anything changed, it ran again — up to 10 iterations — until no more changes (convergence). Too many watchers (>2000) caused noticeable lag. Angular 2+ switched to **unidirectional data flow** + `Zone.js` (monkey-patches async APIs to detect when async operations complete). React's virtual DOM diffing and Vue's Proxy-based reactivity are fundamentally different and far more efficient than dirty checking.

---

### 153. What are reactive primitives in state libraries?

"**Reactive primitives** are fine-grained atoms of state that automatically track their consumers and re-run them on change:

```js
// Vanilla reactive primitive (simplified Solid.js / Jotai concept)
function createSignal(initialValue) {
  let value = initialValue;
  const subscribers = new Set();

  const read = () => {
    if (currentEffect) subscribers.add(currentEffect);
    return value;
  };

  const write = (newValue) => {
    value = newValue;
    subscribers.forEach(fn => fn());
  };

  return [read, write];
}

let currentEffect = null;
function createEffect(fn) {
  currentEffect = fn;
  fn(); // run to track dependencies
  currentEffect = null;
}

// Usage
const [count, setCount] = createSignal(0);
createEffect(() => console.log('count is:', count())); // 'count is: 0'
setCount(5); // automatically logs: 'count is: 5'
```"

#### Indepth
This is the exact pattern used by **Solid.js** — the most performant reactive framework as of 2024 (no virtual DOM, no re-renders, just precise DOM updates). **Jotai** and **Recoil** use atoms (minimal reactive state units) in React. **MobX** uses class-based observables. The key idea: instead of diffing a virtual tree, track exactly which computations depend on which state — update only those when state changes. Preact Signals bring this to React's ecosystem.

---

### 154. How to detect if a user is idle?

"Track user activity events and set a timer for inactivity:

```js
class IdleDetector {
  #idleTime = 0;
  #threshold;
  #onIdle;
  #onActive;
  #interval;
  #events = ['mousedown', 'mousemove', 'keydown', 'scroll', 'touchstart'];
  #controller = new AbortController();

  constructor({ threshold = 30000, onIdle, onActive }) {
    this.#threshold = threshold;
    this.#onIdle = onIdle;
    this.#onActive = onActive;
    this.#start();
  }

  #start() {
    const reset = () => {
      if (this.#idleTime >= this.#threshold) this.#onActive?.();
      this.#idleTime = 0;
    };

    this.#events.forEach(event =>
      document.addEventListener(event, reset, { signal: this.#controller.signal, passive: true })
    );

    this.#interval = setInterval(() => {
      this.#idleTime += 1000;
      if (this.#idleTime >= this.#threshold) this.#onIdle?.();
    }, 1000);
  }

  destroy() {
    this.#controller.abort(); // removes all listeners
    clearInterval(this.#interval);
  }
}

const detector = new IdleDetector({
  threshold: 30000, // 30 seconds
  onIdle: () => showSessionWarning(),
  onActive: () => hideWarning(),
});
```"

#### Indepth
The native **Idle Detection API** (`IdleDetector` browser API) provides OS-level idle detection (keyboard, mouse, screen lock) but requires explicit permission and is only supported in Chrome. For production: combine the custom implementation with **Page Visibility API** (`visibilitychange` event) — don't count time when the tab is hidden. Session timeout after inactivity is a **PCI-DSS and banking regulation** requirement for financial apps.

---

### 155. How to manage undo/redo stack in JS?

"A Command pattern-based undo/redo system:

```js
class UndoRedoManager {
  #history = [];
  #future = [];
  #maxHistory;

  constructor(maxHistory = 100) { this.#maxHistory = maxHistory; }

  execute(command) {
    command.execute();
    this.#history.push(command);
    this.#future = []; // clear redo stack on new action

    if (this.#history.length > this.#maxHistory) {
      this.#history.shift(); // keep memory bounded
    }
  }

  undo() {
    if (!this.#history.length) return;
    const command = this.#history.pop();
    command.undo();
    this.#future.push(command);
  }

  redo() {
    if (!this.#future.length) return;
    const command = this.#future.pop();
    command.execute();
    this.#history.push(command);
  }

  get canUndo() { return this.#history.length > 0; }
  get canRedo() { return this.#future.length > 0; }
}

// A command (text editor insert)
class InsertTextCommand {
  constructor(editor, position, text) {
    this.editor = editor; this.position = position; this.text = text;
  }
  execute() { this.editor.insert(this.position, this.text); }
  undo() { this.editor.delete(this.position, this.text.length); }
}
```"

#### Indepth
Real text editors (VS Code, Monaco) use this exact Command pattern. **Immer.js** takes a different approach — it creates **immutable patches** (pairs of `patch` and `inversePatch`) allowing undo by applying inverse patches. Redux DevTools' time-travel works by replaying action history from initial state. For collaborative editors, use **Operational Transformation (OT)** or **CRDTs** which allow concurrent edits from multiple users to merge without conflicts.

---

### 156. How to implement optimistic UI updates?

"Optimistic UI shows the expected result immediately, then reconciles with the server response:

```js
async function toggleTodo(todo) {
  // 1. Optimistically update UI
  const previousState = [...todos]; // save for rollback
  setTodos(prev => prev.map(t =>
    t.id === todo.id ? { ...t, done: !t.done } : t
  ));

  // 2. Sync with server
  try {
    await api.patch(`/todos/${todo.id}`, { done: !todo.done });
    // Server confirmed — no further action needed
  } catch (err) {
    // 3. Rollback on failure
    setTodos(previousState);
    showError('Failed to save. Changes rolled back.');
  }
}
```

React Query and SWR have built-in optimistic update APIs that handle this pattern automatically."

#### Indepth
React Query's `useMutation` `onMutate`/`onError`/`onSettled` hooks implement this pattern:
```js
const mutation = useMutation(updateTodo, {
  onMutate: async (newTodo) => {
    await queryClient.cancelQueries(['todos']);
    const prev = queryClient.getQueryData(['todos']);
    queryClient.setQueryData(['todos'], old => updateData(old, newTodo));
    return { prev }; // context for rollback
  },
  onError: (err, vars, context) => queryClient.setQueryData(['todos'], context.prev),
  onSettled: () => queryClient.invalidateQueries(['todos']),
});
```
This is standard practice in modern UX — Notion, Linear, and Figma all use optimistic updates for a snappy feel.

---

### 157. How to sync Redux or global state with URL query parameters?

"Bidirectional sync between URL and state:

```js
// Custom hook (React) for URL-synced state
function useQueryState(key, defaultValue) {
  const [value, setValue] = useState(() => {
    const params = new URLSearchParams(window.location.search);
    return params.get(key) ?? defaultValue;
  });

  const setQueryValue = useCallback((newVal) => {
    setValue(newVal);
    const params = new URLSearchParams(window.location.search);
    if (newVal === defaultValue) {
      params.delete(key);
    } else {
      params.set(key, newVal);
    }
    // Replace without pushing new history entry
    window.history.replaceState({}, '', `?${params.toString()}`);
  }, [key, defaultValue]);

  // Listen for browser back/forward navigation
  useEffect(() => {
    const handler = () => {
      const params = new URLSearchParams(window.location.search);
      setValue(params.get(key) ?? defaultValue);
    };
    window.addEventListener('popstate', handler);
    return () => window.removeEventListener('popstate', handler);
  }, [key, defaultValue]);

  return [value, setQueryValue];
}
```"

#### Indepth
Libraries: `nuqs` (Next.js type-safe URL state), `use-query-params`, `react-router` `useSearchParams`. Key use cases: **shareable filtered views** (e-commerce filters, dashboard configurations), **pagination state**, **selected tab**. Gotcha: `history.pushState` creates a new browser history entry (back button goes there); `replaceState` updates in-place. Use `pushState` for significant navigation, `replaceState` for ephemeral filter/sort state.

---

### 158. What is the FLIP animation technique in JavaScript?

"**FLIP** (First, Last, Invert, Play) is a technique for creating smooth, performant animations by calculating the difference between two states and using CSS transforms:

```js
function flipAnimate(element, callback) {
  // First: record initial position
  const first = element.getBoundingClientRect();

  // Run DOM change
  callback();

  // Last: get final position after DOM change
  const last = element.getBoundingClientRect();

  // Invert: calculate delta transforms
  const deltaX = first.left - last.left;
  const deltaY = first.top - last.top;
  const deltaW = first.width / last.width;
  const deltaH = first.height / last.height;

  // Apply inverted transform (puts element back at original position visually)
  element.animate([
    { transform: `translate(${deltaX}px, ${deltaY}px) scale(${deltaW}, ${deltaH})` },
    { transform: 'translate(0, 0) scale(1, 1)' }
  ], {
    duration: 300,
    easing: 'ease-in-out',
  });
}
```"

#### Indepth
FLIP is brilliant because **CSS transforms are GPU-accelerated** and don't trigger layout recalculations — only composition. The trick: after a layout change, move the element back to where it was (via transform), then animate the transform to `none` — the browser smoothly composites the animation without touching the layout engine. Libraries: `Framer Motion`'s `layout` prop, `AutoAnimate`, and `Vue`'s `<transition>` group all implement FLIP. This is what makes React's `react-flip-toolkit` library work.

---

### 159. How does the Virtual DOM optimize JS performance?

"The **Virtual DOM** is an in-memory JavaScript representation of the actual DOM. Instead of directly manipulating the DOM on every state change, UI libraries (React, Vue 2) compute the **diff** between old and new virtual DOM trees and batch-apply only the minimal real DOM changes.

```
State Change → New VDOM Tree (JS object, cheap)
                     ↓
        Diff(OldVDOM, NewVDOM) → patch set
                     ↓
        Apply patches to real DOM (minimal operations)
```

The real DOM is slow because accessing/modifying it triggers layout, style, and paint recalculations. Virtual DOM allows JS to do all computation on plain objects (fast) before touching the DOM (slow)."

#### Indepth
React's **Fiber** reconciler allows interruptible rendering — long reconciliation work can be paused and resumed, preventing UI jank. React 18's **concurrent mode** uses `startTransition` to mark non-urgent updates. Virtual DOM's edge cases: 1) **key prop** is critical for list reconciliation — without it, React re-creates elements instead of reusing them. 2) **Reconciliation is O(n)** with heuristics (same component type = update, different = replace). **Solid.js** shows that fine-grained reactivity (no VDOM, no reconciliation) can be even faster.

---

### 160. What is `structuredClone()` vs `JSON.parse(JSON.stringify())`?

"`structuredClone()` is the **correct, modern** deep clone method using the HTML Structured Clone Algorithm.

| Feature | `JSON` round-trip | `structuredClone` |
|---|---|---|
| `Date` | → string (broken) | ✅ preserved as Date |
| `undefined` | dropped | ✅ preserved |
| Functions | dropped | ❌ throws |
| `Map`/`Set` | → `{}` or `[]` | ✅ preserved |
| `RegExp` | → `{}` | ✅ preserved |
| Circular refs | throws | ✅ handled |
| `Symbol` keys | dropped | ❌ dropped |
| `Infinity`/`NaN` | → `null` | ✅ preserved |
| Class instances | loses prototype | ❌ loses prototype |
| Performance | faster for small objects | faster for large complex |

```js
const data = { date: new Date(), map: new Map([['a', 1]]), set: new Set([1, 2]) };
const clone1 = JSON.parse(JSON.stringify(data)); // Dates are strings, Map is {}, Set is {}
const clone2 = structuredClone(data); // All preserved correctly ✅
```"

#### Indepth
`structuredClone` was standardized in 2021, available in Node.js 17+ and all modern browsers. For environments without it: `core-js` polyfill or `rfdc` (really fast deep clone) library. Alternatives: `lodash.cloneDeep` handles most cases but doesn't handle circular references that reference themselves multiple times the same way. For class instances requiring deep clone: implement a `clone()` method that explicitly calls `structuredClone` on serializable parts and reconstructs the class instance.
