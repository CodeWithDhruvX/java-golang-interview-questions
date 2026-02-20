# 🟡 **JavaScript Functions, Closures & Scope — Junior / Mid Level**

> **Target Companies:** Wipro, HCL, LTIMindtree, Tech Mahindra (Service-Based) + Mid-size Product Companies, SDE-1/SDE-2 roles

---

### 31. What is a pure function?

"A **pure function** is one that:
1. Always returns the **same output** for the same input.
2. Has **no side effects** — it doesn't modify external state, make API calls, or touch the DOM.

```js
// Pure
const add = (a, b) => a + b;

// Impure (side effect: modifies external state)
let total = 0;
const addToTotal = (n) => { total += n; };
```

I love pure functions because they're **predictable, testable, and composable**. They're the foundation of functional programming and React's component model."

#### Indepth
Pure functions are **referentially transparent** — you can replace a call with its return value without changing program behavior. This enables **memoization** (caching results), **parallel execution** without race conditions, and **time-travel debugging** in Redux. The hardest part in real apps is quarantining side effects to the edges of your system.

---

### 32. What is memoization?

"**Memoization** is a performance optimization where you **cache the results** of expensive function calls and return the cached result when the same inputs occur again.

```js
function memoize(fn) {
  const cache = new Map();
  return function(...args) {
    const key = JSON.stringify(args);
    if (cache.has(key)) return cache.get(key);
    const result = fn.apply(this, args);
    cache.set(key, result);
    return result;
  };
}

const slowSquare = (n) => { /* heavy computation */ return n * n; };
const fastSquare = memoize(slowSquare);
fastSquare(10); // computes
fastSquare(10); // returns cached result instantly
```"

#### Indepth
The key challenge in memoization is the **cache key**. `JSON.stringify` works for simple args but fails for functions, circular references, or `undefined`. For production, libraries like `lodash.memoize` use `Map` with the first argument as the key by default. **React.memo**, **useMemo**, and **useCallback** are React's built-in memoization for components and values.

---

### 33. What are IIFEs (Immediately Invoked Function Expressions)?

"An **IIFE** is a function that is defined and called **immediately** at the same time.

```js
(function() {
  const secret = 'hidden';
  // Everything here is private
  console.log('IIFE ran!');
})();

// Arrow IIFE
(() => {
  console.log('Arrow IIFE');
})();
```

Before ES6 modules, IIFEs were the primary way to create **private scope** and avoid polluting the global namespace. jQuery itself used an IIFE."

#### Indepth
The wrapping `()` is necessary because the JS parser would otherwise see `function` at the start of a line as a **function declaration** (which can't be immediately invoked). The outer parens force it to be interpreted as an **expression**. Today, ES6 modules replaced IIFEs for encapsulation, but IIFEs still appear in bundler output (like Webpack's IIFE for each module) and in polyfills.

---

### 34. How do closures preserve data?

"A closure captures **references** to outer variables, not copies. The captured variables live on the **heap** as long as the closure exists.

```js
function makeAdder(x) {
  return (y) => x + y; // 'x' is captured by closure
}
const add5 = makeAdder(5);
add5(3); // 8 — 'x' (5) is still alive in memory
```

Even after `makeAdder` returns, the inner arrow function holds a live reference to `x`. This is the essence of closure — **functions carrying their environment with them**."

#### Indepth
Each call to `makeAdder` creates a **new closure** with its own `x`. This is how you avoid the classic `var` in loop bug — each iteration gets its own binding. Under the hood, V8 allocates the closed-over variable on the heap via **escape analysis**. Variables not captured by any closure stay on the stack.

---

### 35. What is currying?

"**Currying** transforms a function with multiple arguments into a series of functions, each taking **one argument** at a time.

```js
// Normal function
const multiply = (a, b) => a * b;

// Curried version
const curriedMultiply = (a) => (b) => a * b;

const double = curriedMultiply(2);
double(5);  // 10
double(7);  // 14
```

Currying enables **partial application** — pre-filling some arguments to create specialized functions. It's widely used in functional programming libraries like Ramda and in building reusable utility functions."

#### Indepth
A generalized curry function that works for any arity:
```js
function curry(fn) {
  return function curried(...args) {
    if (args.length >= fn.length) return fn(...args);
    return (...more) => curried(...args, ...more);
  };
}
```
This compares accumulated args to `fn.length` (the number of declared parameters). Libraries like lodash's `_.curry` also support **placeholder arguments** (`_`) for gaps in partial application.

---

### 36. What is the prototype chain?

"Every JavaScript object has a hidden `[[Prototype]]` link to another object. When you access a property that doesn't exist on an object, JS looks up this chain until it finds it or reaches `null`.

```js
const arr = [1, 2, 3];
// arr → Array.prototype → Object.prototype → null
arr.push(4);       // found on Array.prototype
arr.toString();    // found on Object.prototype
arr.nonExistent;   // undefined (end of chain reached)
```

This is how **inheritance** works in JavaScript — not by copying, but by linking."

#### Indepth
The prototype chain is the fundamental mechanism of JS inheritance. `Object.getPrototypeOf(arr) === Array.prototype` is `true`. Setting `Object.prototype.myProp = 'shared'` affects ALL objects — this is **prototype pollution**, a serious security vulnerability. `Object.create(null)` creates objects with NO prototype, useful for safe dictionary objects.

---

### 37. What is the difference between `__proto__` and `prototype`?

"`prototype` is a **property on function objects** (constructors). It's the object that will be assigned as `[[Prototype]]` for instances created by that function.

`__proto__` is a **property on instances** that points to their prototype — it's an accessor for `[[Prototype]]`.

```js
function Dog(name) { this.name = name; }
Dog.prototype.bark = function() { return 'Woof!'; };

const rex = new Dog('Rex');
rex.__proto__ === Dog.prototype; // true
```"

#### Indepth
`__proto__` was non-standard but so widely used that ES6 standardized it as an accessor on `Object.prototype`. The formally correct way to access/set the prototype is `Object.getPrototypeOf(obj)` and `Object.setPrototypeOf(obj, proto)` — but `setPrototypeOf` is slow and should be avoided in performance-critical code. The `class` syntax in ES6 is syntax sugar over the prototype chain.

---

### 38. What is `Object.create()` used for?

"`Object.create(proto)` creates a **new object** with the specified object as its `[[Prototype]]`. It's the foundation of prototypal inheritance.

```js
const animal = {
  speak() { console.log(`${this.name} makes a sound.`); }
};

const dog = Object.create(animal);
dog.name = 'Rex';
dog.speak(); // 'Rex makes a sound.' (found via prototype chain)
```

I use `Object.create(null)` to create a pure dictionary object with no inherited properties — safer for dynamic key storage since you can't accidentally access `toString`, `hasOwnProperty`, etc."

#### Indepth
`Object.create` accepts a second argument — property descriptors. This is how `class` inheritance works internally: `Object.create(Parent.prototype)` sets up the child's prototype. Before ES6 classes, this was the standard way to implement inheritance in frameworks like Backbone.js.

---

### 39. How does inheritance work in JavaScript?

"JavaScript uses **prototypal inheritance** — objects inherit directly from other objects via the prototype chain, not from classes like in Java.

```js
class Animal {
  constructor(name) { this.name = name; }
  speak() { return `${this.name} makes a noise.`; }
}

class Dog extends Animal {
  speak() { return `${this.name} barks.`; }
}

const d = new Dog('Rex');
d.speak(); // 'Rex barks.'
d instanceof Animal; // true
```

ES6 `class` and `extends` are syntactic sugar — under the hood, it's still prototype chaining."

#### Indepth
`extends` sets up two prototype chains: `Dog.prototype → Animal.prototype` (for instance methods) and `Dog → Animal` (for static methods). `super()` calls the parent constructor and is **mandatory** in a `constructor` that extends another class — forgetting it causes a ReferenceError because `this` is not initialized until `super()` runs.

---

### 40. What are ES6 classes?

"ES6 classes are **syntactic sugar** over JavaScript's prototype-based system. They make OOP patterns more familiar to developers coming from Java or Python.

```js
class Person {
  #age; // private field (ES2022)

  constructor(name, age) {
    this.name = name;
    this.#age = age;
  }

  greet() { return `Hi, I'm ${this.name}`; }
  get age() { return this.#age; }

  static create(name, age) { return new Person(name, age); }
}
```

They support constructors, methods, static methods, getters/setters, private fields (`#`), and inheritance. But they're still prototypal under the hood."

#### Indepth
Key differences from Java classes: JS classes are **first-class values** (can be passed as arguments), class bodies are always in strict mode, and class declarations are **NOT hoisted** (in the same way functions are). Private fields (`#name`) are genuinely private — even reflection can't access them, unlike the old closure-based privacy pattern.

---

### 41. What are generators in JavaScript?

"A **generator** is a function that can **pause** and **resume** its execution. It uses the `function*` syntax and `yield` to pause.

```js
function* counter() {
  let i = 0;
  while (true) {
    yield i++;
  }
}

const gen = counter();
gen.next().value; // 0
gen.next().value; // 1
gen.next().value; // 2
```

Generators are perfect for **lazy evaluation** (generating infinite sequences), **custom iterators**, and **async flows** (they were the precursor to async/await via co.js)."

#### Indepth
Calling a generator function returns a **generator object** that implements the **iterator protocol** (`next()` method). Each `yield` suspends the function and returns `{ value, done }`. `generator.return(val)` terminates it. `generator.throw(err)` injects an error at the yield point. Redux Saga uses generators extensively to handle side effects.

---

### 42. What is the spread operator?

"The spread operator (`...`) expands an **iterable** (array, string, Set, Map) into individual elements, or spreads object properties.

```js
// Arrays
const a = [1, 2, 3];
const b = [...a, 4, 5]; // [1, 2, 3, 4, 5]

// Objects (shallow merge)
const obj1 = { a: 1 };
const merged = { ...obj1, b: 2 }; // { a: 1, b: 2 }

// Function calls
Math.max(...[1, 5, 3]); // 5
```

I use it everywhere for immutable updates — especially in React state: `setState({ ...prevState, name: 'new' })`."

#### Indepth
Object spread follows property order: later spreads overwrite earlier ones. `{ ...defaults, ...overrides }` is a clean config merge pattern. Spread creates a **shallow copy** — nested objects are still shared by reference. Crucially, object spread only copies **own enumerable** properties, not prototype methods.

---

### 43. What is the rest parameter?

"The rest parameter (`...`) collects **remaining arguments** into an array. It must be the last parameter.

```js
function sum(...numbers) {
  return numbers.reduce((a, b) => a + b, 0);
}
sum(1, 2, 3, 4); // 10

// With other params
function first(x, y, ...rest) {
  console.log(rest); // all args after x and y
}
```

This replaces the old `arguments` object pattern. Unlike `arguments`, the rest parameter is a **real array** with all array methods."

#### Indepth
`arguments` is an array-like object (not a real array), doesn't work in arrow functions, and includes ALL arguments. Rest parameters only collect the **remaining** ones and return a true `Array`. The spread operator and rest parameter look identical (`...`) but work in opposite directions — spread expands, rest collects.

---

### 44. How do modules (`import`/`export`) work?

"ES Modules (`ESM`) let you split code into files and share functionality using `export` and `import`.

```js
// math.js — named exports
export const PI = 3.14;
export function add(a, b) { return a + b; }

// utils.js — default export
export default function greet(name) { return `Hello, ${name}`; }

// main.js — importing
import greet from './utils.js';
import { PI, add } from './math.js';
import * as Math from './math.js'; // namespace import
```"

#### Indepth
ES Modules are **statically analyzed** — the imports/exports must be at the top level, not inside conditionals. This allows bundlers to perform **tree-shaking** (dead code elimination). They're also **singletons** — the same module is only evaluated once, no matter how many times it's imported. This is different from CommonJS (`require`) which can be dynamic and returns cached exports.

---

### 45. What is optional chaining?

"Optional chaining (`?.`) lets you safely access **deeply nested properties** without throwing if an intermediate property is `null` or `undefined`.

```js
const user = { address: null };

// Without optional chaining — throws TypeError
user.address.city; // ❌

// With optional chaining — returns undefined safely
user.address?.city; // ✅ undefined

// Also works for methods and arrays
user.getAge?.();         // calls only if getAge exists
arr?.[0]?.name;          // safe array access
```"

#### Indepth
Optional chaining **short-circuits** — if the left side is `null/undefined`, the entire expression immediately returns `undefined` without evaluating the right side. It can be combined with nullish coalescing: `user?.profile?.bio ?? 'No bio'`. TypeScript was the one that popularized optional chaining before it became a JS standard in ES2020.

---

### 46. What is the microtask queue vs macrotask queue?

"The **microtask queue** (job queue) holds: Promise callbacks (`.then`/`.catch`/`.finally`), `queueMicrotask()`, and MutationObserver callbacks.
The **macrotask queue** (task queue) holds: `setTimeout`, `setInterval`, `setImmediate` (Node), and I/O callbacks.

**Priority**: After each macrotask, the JS engine drains the **entire** microtask queue before picking the next macrotask.

```js
setTimeout(() => console.log('macro'), 0);
Promise.resolve().then(() => console.log('micro'));
console.log('sync');
// Output: sync → micro → macro
```"

#### Indepth
This ordering matters for correctness in frameworks. React's batched state updates leverage microtasks to synchronize DOM updates after all state changes in an event handler. In Node.js, `process.nextTick()` runs even BEFORE the microtask queue — it's a special ultra-high-priority queue. `setImmediate()` is a macrotask that runs after I/O callbacks.

---

### 47. How does `setTimeout` actually work?

"`setTimeout` is a **browser/Node API** — it's not part of the JS language itself. It schedules a callback to run **after at least** N milliseconds (not exactly N).

```js
setTimeout(() => console.log('fired'), 0);
// Still runs AFTER all synchronous code and microtasks
```

When the timer fires, the callback is placed in the **macrotask queue**. It won't run until the call stack is empty. This is why `setTimeout(fn, 0)` doesn't mean 'immediately'."

#### Indepth
Browsers enforce a **minimum delay of 4ms** for nested `setTimeout` calls (5+ levels deep) per the HTML spec. Node.js uses `libuv`'s timer implementation which is more granular. In practice, timers can fire **later** than requested but never earlier. `setTimeout` returning the macrotask queue (not microtask) is a key concept for understanding execution order.

---

### 48. What is a race condition? How do you avoid it?

"A **race condition** occurs when the outcome depends on the **timing/order** of multiple async operations, leading to unpredictable results.

```js
// Classic race condition in search
let lastRequestId = 0;
async function search(query) {
  const id = ++lastRequestId;
  const result = await fetch(`/search?q=${query}`);
  if (id !== lastRequestId) return; // Ignore stale results
  updateUI(result);
}
```

Common solutions: **AbortController** to cancel previous requests, **debouncing** to limit how often requests fire, and **request IDs** to ignore stale responses."

#### Indepth
`Promise.race()` is related — it resolves/rejects with whichever promise settles first. Race conditions in JS mainly occur with async UI updates (search, forms). **AbortController** + `signal` passed to `fetch` lets you cancel in-flight requests: `fetch(url, { signal: controller.signal })`. The search pattern above (ID checking) is called **optimistic concurrency control**.

---

### 49. How does JavaScript handle concurrency?

"JavaScript achieves **concurrency** (doing multiple things) through its **non-blocking event loop**, not through multi-threading.

I/O operations (fetch, disk reads) are offloaded to the browser's or Node's runtime APIs (written in C++). When complete, their callbacks enter the queue. The event loop picks them up when the stack is free.

For true parallelism (multi-core), JavaScript uses **Web Workers** (browser) or **Worker Threads** (Node.js), which run on separate threads without shared memory (except SharedArrayBuffer)."

#### Indepth
JS's concurrency model is called the **Reactor pattern** (vs Java's thread-based approach). It's excellent for I/O-heavy workloads (like a Node.js web server). It struggles with CPU-heavy tasks (image processing, crypto) on the main thread. Node.js handles thousands of concurrent requests with a single thread via event loop, which is why it outperforms thread-per-request models for I/O-bound work.

---

### 50. What is a debounce vs throttle function?

"**Debounce** delays execution until a specified time has passed since the **last** invocation. Good for: search-as-you-type, resize handlers.
**Throttle** ensures a function runs **at most once** per specified interval. Good for: scroll handlers, rate-limiting API calls.

```js
// Debounce: search fires 300ms after user stops typing
const onSearch = debounce((query) => fetchResults(query), 300);

// Throttle: scroll handler fires at most every 100ms
const onScroll = throttle(() => updateUI(), 100);
```

The key difference: debounce **resets** the timer on each call; throttle **ignores** calls that happen within the interval."

#### Indepth
Implementation:
```js
function debounce(fn, delay) {
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => fn(...args), delay);
  };
}

function throttle(fn, limit) {
  let inThrottle;
  return (...args) => {
    if (!inThrottle) {
      fn(...args);
      inThrottle = true;
      setTimeout(() => inThrottle = false, limit);
    }
  };
}
```
`requestAnimationFrame` is an alternative to throttle for visual updates — it syncs with the browser's repaint cycle (~60fps), providing natural throttling.

---

### 51. What are common JavaScript errors?

"The most common are:
- **ReferenceError**: accessing an undeclared variable (`x is not defined`)
- **TypeError**: calling a non-function, accessing property of `null/undefined`
- **SyntaxError**: invalid JS syntax (caught at parse time)
- **RangeError**: value out of valid range (e.g., `new Array(-1)`)
- **URIError**: invalid URI encoding

The most common in production: `TypeError: Cannot read property 'x' of undefined` — almost always a null/undefined check missing."

#### Indepth
All standard errors extend the `Error` base class. Custom errors can extend it: `class ValidationError extends Error { constructor(msg) { super(msg); this.name = 'ValidationError'; } }`. The `name` property distinguishes error types. In production, use error monitoring tools (Sentry, Datadog) to capture `error.stack` traces which show the exact call chain.

---

### 52. How to handle errors globally?

"In the browser:
```js
// Uncaught synchronous errors
window.onerror = (msg, src, line, col, err) => { logToServer(err); };

// Unhandled Promise rejections
window.addEventListener('unhandledrejection', (event) => {
  logToServer(event.reason);
  event.preventDefault(); // Prevents console error
});
```

In Node.js:
```js
process.on('uncaughtException', (err) => { logger.fatal(err); process.exit(1); });
process.on('unhandledRejection', (reason) => { logger.error(reason); });
```

These are safety nets — I still handle errors at the source, using these only for unexpected failures."

#### Indepth
For production apps, integrate **Sentry** or similar: `import * as Sentry from '@sentry/browser'; Sentry.init({ dsn: '...' });`. It automatically hooks into these global handlers. The `unhandledrejection` event prevents silent promise failures — a huge problem before Node.js started crashing on them by default. Always `process.exit(1)` after `uncaughtException` because the process state may be corrupt.

---

### 53. How does `try...catch` with `finally` work?

"`finally` always executes after `try` and `catch`, regardless of whether an error was thrown or not — even if the `try` block has a `return` statement.

```js
function fetchData() {
  try {
    return riskyOperation();
  } catch (err) {
    handleError(err);
    return null;
  } finally {
    hideLoadingSpinner(); // ALWAYS runs
  }
}
```

If `finally` also has a `return`, it **overrides** the `try` or `catch` return value — a surprising behavior to be aware of."

#### Indepth
The `finally` block runs even in these cases: 1) `try` returns, 2) `catch` returns, 3) an error is thrown by `catch`. The JavaScript engine evaluates it as: put the completion value on hold, run `finally`, then use the finally's completion value if it has one (override), or use the original value. This override behavior makes it dangerous to return from `finally`.

---

### 54. What are custom errors?

"Custom errors extend the built-in `Error` class to represent specific error types in your domain:

```js
class ValidationError extends Error {
  constructor(message, field) {
    super(message);
    this.name = 'ValidationError';
    this.field = field;
  }
}

class NotFoundError extends Error {
  constructor(resource) {
    super(`${resource} not found`);
    this.name = 'NotFoundError';
    this.statusCode = 404;
  }
}

// Usage
try {
  throw new ValidationError('Email is invalid', 'email');
} catch (err) {
  if (err instanceof ValidationError) {
    highlightField(err.field);
  } else {
    throw err; // Re-throw unknown errors
  }
}
```"

#### Indepth
Always set `this.name` explicitly — subclasses don't automatically get the class name as the error name. Also call `Error.captureStackTrace(this, this.constructor)` in Node.js to clean the stack trace (removes the constructor from the trace). In TypeScript, you need `Object.setPrototypeOf(this, CustomError.prototype)` for `instanceof` to work correctly due to a transpilation quirk.

---

### 55. What is a Set and how is it different from an array?

"A **Set** is a collection of **unique values**. Duplicates are automatically ignored.

```js
const set = new Set([1, 2, 2, 3, 3]);
set.size;              // 3
set.has(2);            // true
set.add(4);
set.delete(1);
[...set];              // convert to array: [2, 3, 4]
```

Key differences from Array:
- No duplicate values
- No index-based access (`set[0]` doesn't work)
- `has()` is O(1) vs array `includes()` which is O(n)
- Iteration order is **insertion order** (like arrays)"

#### Indepth
`Set` uses **SameValueZero** comparison (same as `===` but `NaN === NaN` here). `Set` of objects compares by **reference**, not value — `new Set([{a:1}, {a:1}])` has 2 items because they're different object references. For fast membership testing on large datasets, `Set` is dramatically faster than `Array`. `WeakSet` stores objects only, allows GC (no `.size` or iteration).

---
