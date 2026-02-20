# 🟢 **JavaScript Basics — Fresher / Junior Level**

> **Target Companies:** TCS, Infosys, Wipro, Cognizant, Capgemini, Accenture (Service-Based) + Early-stage startups, Junior Frontend roles (Product-Based)

---

### 1. What are the different data types in JavaScript?

"JavaScript has **8 data types** split into two categories.

**Primitive types**: `string`, `number`, `bigint`, `boolean`, `undefined`, `symbol`, and `null`.
**Non-primitive**: `object` (which includes arrays, functions, dates, etc.).

I remember it as — anything that's not an object is a primitive. Primitives are **immutable** and stored by value on the stack. Objects are stored by reference on the heap."

#### Indepth
JavaScript uses dynamic typing, so a variable can hold any type at any time. The `typeof` operator reveals the type at runtime. One legacy bug: `typeof null === 'object'` (it should be `'null'`) — this exists for backward compatibility since JavaScript 1.0.

---

### 2. What is the difference between `var`, `let`, and `const`?

"`var` is function-scoped and hoisted (initialized as `undefined`). `let` and `const` are block-scoped and hoisted but **not initialized** (causing a Temporal Dead Zone).

`const` must be assigned at declaration and cannot be reassigned, but the object it points to **can still be mutated**.

I always use `const` by default, `let` when reassignment is needed, and try to avoid `var` entirely to prevent accidental scope leakage."

#### Indepth
`var` declarations are added as properties to the global `window` object in browsers. `let` and `const` are stored in a separate declarative environment record and are NOT properties of `window`. This is why `let x = 1; console.log(window.x)` returns `undefined`.

---

### 3. What is hoisting in JavaScript?

"Hoisting is JavaScript's behavior of **moving declarations to the top** of their scope during the compilation phase, before code executes.

`var` declarations are hoisted and initialized as `undefined`. `function` declarations are hoisted completely (both declaration and body). `let` and `const` are hoisted but remain in the **Temporal Dead Zone** until their line is reached.

I rely on function hoisting to call utility functions defined later in a file, but I never rely on `var` hoisting — it's a classic source of bugs."

#### Indepth
Hoisting doesn't literally move code. The JS engine does two passes: a **compilation phase** (creates the memory space for declarations) and an **execution phase** (runs the code line by line). `let`/`const` are allocated but deliberately not initialized, causing a `ReferenceError` if accessed early — this is safer behavior than `var`'s silent `undefined`.

---

### 4. How does type coercion work in JavaScript?

"Type coercion is JavaScript's **automatic conversion** of values from one type to another.

It happens in two ways: **implicit** (JS does it for you, like `'5' + 1 = '51'`) and **explicit** (you do it yourself, like `Number('5')`).

The `+` operator triggers string coercion if either operand is a string. Comparison operators like `==` trigger numeric coercion. This is why I always use `===` (strict equality) which skips coercion."

#### Indepth
The coercion rules follow the abstract `ToPrimitive` operation. For objects, JavaScript calls `valueOf()` first, then `toString()`. This is why `[] + {}` gives `"[object Object]"` and `{} + []` gives `0` (the `{}` is parsed as an empty block, not an object, in statement position).

---

### 5. What are truthy and falsy values?

"**Falsy values** are values that evaluate to `false` in a boolean context. There are exactly 6: `false`, `0`, `''` (empty string), `null`, `undefined`, and `NaN`.

Everything else is **truthy** — including `[]`, `{}`, `'0'`, and `'false'`!

I use this knowledge for concise guard clauses: `if (user && user.name)` instead of verbose null checks."

#### Indepth
This matters in real-world code. A common bug: checking `if (count)` fails when `count === 0` which is a valid number but falsy. The fix is `if (count !== undefined)` or `if (count != null)`. The double-equals null check `!= null` is one of the few cases where `==` is actually preferred over `===`.

---

### 6. What is the difference between `==` and `===`?

"`==` (loose equality) compares values **after type coercion**. `===` (strict equality) compares value AND type **without coercion**.

`0 == '0'` is `true` because JS converts the string to a number. `0 === '0'` is `false` because the types differ.

I always use `===`. The only exception is `null == undefined` which is a clean idiom to check for both null and undefined in one go."

#### Indepth
The `==` comparison follows the Abstract Equality Comparison algorithm in the ECMAScript spec. It has 12 steps with rules for all type combinations. `===` is the Strict Equality Comparison — much simpler: if types differ, return `false`. Using `===` everywhere also helps linting tools and TypeScript catch type errors.

---

### 7. Explain the difference between `null` and `undefined`.

"`undefined` means a variable has been declared but **not assigned a value**. It's the engine's default. `null` is an intentional assignment meaning **no value** — it's programmer-defined emptiness.

A function that doesn't explicitly return a value returns `undefined`. I set a variable to `null` when I want to actively signal 'this intentionally has no value', like `user = null` after logout.

`typeof undefined` is `'undefined'`. `typeof null` is `'object'` (a famous bug)."

#### Indepth
`null == undefined` is `true` (loose equality), but `null === undefined` is `false`. In practice, APIs often return `null` for 'not found' and `undefined` for 'not set'. The distinction matters for database operations where `null` has semantic meaning (nullable column) vs `undefined` meaning the field wasn't provided at all.

---

### 8. How does `typeof` work?

"`typeof` is a **unary operator** that returns a string indicating the type of its operand.

Common results: `typeof 'hello'` → `'string'`, `typeof 42` → `'number'`, `typeof true` → `'boolean'`, `typeof undefined` → `'undefined'`, `typeof function(){}` → `'function'`.

The gotcha: `typeof null` → `'object'` (historical bug) and `typeof []` → `'object'`. To check for arrays, I use `Array.isArray()`."

#### Indepth
`typeof` is safe to use even on undeclared variables — it won't throw a `ReferenceError`. `typeof undeclaredVar` returns `'undefined'`. This was historically used for feature detection like `if (typeof window !== 'undefined')`. Today, optional chaining and environment checks are more idiomatic.

---

### 9. What are template literals in JavaScript?

"Template literals are strings defined with **backticks** (`` ` ``) instead of quotes. They support **multi-line strings** and **expression interpolation** using `${expression}`.

```js
const name = 'Dhruv';
const msg = `Hello, ${name}! Today is ${new Date().toDateString()}.`;
```

I use them constantly — they're cleaner than string concatenation and the multi-line support alone makes them worth it for generating HTML snippets or SQL queries."

#### Indepth
Template literals can also be used as **tagged templates**: a function placed before the backtick receives the string parts and interpolated values as separate arguments. Libraries like `gql`, `css`, and `html` from lit-element use this to create DSLs inside JavaScript. It's an underused but powerful feature.

---

### 10. How does the `switch` statement work?

"`switch` evaluates an expression once and then matches it against `case` labels using **strict equality** (`===`). It executes the matching block and continues until it hits a `break` or the end.

The key gotcha is **fall-through**: if you omit `break`, execution continues into the next `case`. This can be intentional (grouping cases) or a bug.

I use `switch` for clean multi-way branching on an enum-like value, but for complex logic I often prefer an **object dispatch table** which is more readable."

#### Indepth
The `default` case runs if no `case` matches, similar to an `else` clause. Fall-through is sometimes used intentionally: grouping multiple cases to share the same handler. `switch` uses strict equality, so `switch('1') { case 1: }` won't match. A common pattern to avoid switch entirely: `const handlers = { a: fnA, b: fnB }; handlers[key]?.()`.

---

### 11. What is the difference between function declaration and function expression?

"A **function declaration** is hoisted completely and can be called before its definition. A **function expression** is assigned to a variable — only the variable is hoisted (`undefined` for `var`, TDZ for `let/const`).

```js
// Declaration — works before definition due to hoisting
greet(); // ✅
function greet() { return 'hello'; }

// Expression — NOT hoisted
greet(); // ❌ ReferenceError
const greet = function() { return 'hello'; };
```

I prefer function declarations for top-level named utilities and expressions (especially arrow functions) for callbacks and inline logic."

#### Indepth
Function declarations are also block-scoped in strict mode. In non-strict mode, a function declaration inside an `if` block behaves differently across engines — this is one of the most inconsistent areas of JS. In production code, always use `'use strict'` or an ES module (which are always strict).

---

### 12. What is a callback function?

"A **callback** is a function passed as an argument to another function, to be called **later** — either synchronously or asynchronously.

```js
[1, 2, 3].forEach(num => console.log(num)); // synchronous callback
setTimeout(() => console.log('done'), 1000); // async callback
```

Callbacks were the original async pattern in JavaScript. Their downside is **callback hell** — deeply nested callbacks that are hard to read. Promises and async/await were introduced to solve this."

#### Indepth
Callbacks follow the **continuation-passing style (CPS)**. Node.js standardized the **error-first callback** pattern: `callback(err, data)`. The first argument is always an error (or null if none). Libraries like `async.js` were built to manage complex callback flows before Promises became standard.

---

### 13. What are arrow functions?

"Arrow functions are a **shorter syntax** for function expressions. They also have two key differences: they do **not** have their own `this` (they inherit from the enclosing scope) and they cannot be used as constructors.

```js
// Traditional
const add = function(a, b) { return a + b; };
// Arrow
const add = (a, b) => a + b;
```

I use arrow functions for callbacks and methods where I need lexical `this`, like event handlers inside class methods."

#### Indepth
Arrow functions also don't have their own `arguments` object. They cannot be used with `new` (no prototype). They cannot be used as generator functions. Because of lexical `this`, they're ideal for React class components or event listeners where losing `this` context was a classic bug that required `.bind(this)`.

---

### 14. What is lexical scoping?

"**Lexical scoping** means a variable's scope is determined by **where it is written** in the source code, not where it is called.

A function can access variables from its outer (enclosing) scope. This is resolved at definition time, not at call time.

```js
const name = 'global';
function outer() {
  const name = 'outer';
  function inner() { console.log(name); } // 'outer', not 'global'
  inner();
}
```"

#### Indepth
This is the foundation of **closures**. The JS engine builds a **scope chain** at parse time. When a variable is accessed, it walks up the chain from the innermost to the outermost scope. This is different from **dynamic scoping** (used in Bash, Perl) where scope is determined by the call stack at runtime.

---

### 15. What is the use of closures in JavaScript?

"A **closure** is when an inner function **remembers** the variables of its outer function even after the outer function has returned.

```js
function makeCounter() {
  let count = 0;
  return () => ++count;
}
const counter = makeCounter();
counter(); // 1
counter(); // 2
```

I use closures for **data privacy** (like the counter above), **partial application**, and **memoization**. They're fundamental to JS — even event handlers and module patterns are closures."

#### Indepth
Memory-wise, closures keep the entire variable environment alive as long as the closure exists. This can cause memory leaks if closures are stored in long-lived collections (like arrays of event handlers) and the outer variables are large objects. Always dereference (`= null`) closures when they're no longer needed.

---

### 16. How do default parameters work?

"Default parameters let you specify a fallback value for a function argument if it's `undefined`.

```js
function greet(name = 'World') {
  return `Hello, ${name}!`;
}
greet();          // 'Hello, World!'
greet('Dhruv');   // 'Hello, Dhruv!'
```

The default is only used when the argument is `undefined` — NOT when it's `null`. This is important; passing `null` explicitly won't trigger the default."

#### Indepth
Default parameters are evaluated at **call time**, not definition time. This means you can use expressions: `function fn(arr = [])` creates a **new** empty array for each call, unlike Python where `def fn(arr=[])` shares one array across calls. You can even reference earlier parameters: `function fn(x, y = x * 2)`.

---

### 17. How to clone an object in JavaScript?

"There are several approaches depending on depth:

**Shallow clone**: `{ ...obj }` or `Object.assign({}, obj)` — copies only top-level properties.
**Deep clone**: `structuredClone(obj)` (modern, recommended) or `JSON.parse(JSON.stringify(obj))` (simple but loses functions, Dates become strings).

I always use `structuredClone()` now for deep cloning. It handles circular references, Dates, Maps, Sets, and TypedArrays correctly — which `JSON` doesn't."

#### Indepth
`JSON.stringify/parse` fails silently for: `undefined` values (dropped), functions (dropped), `Symbol` keys (dropped), `Date` objects (become strings), `Infinity`/`NaN` (become `null`). `structuredClone()` uses the Structured Clone Algorithm, the same one browsers use for `postMessage`. For libraries, `lodash.cloneDeep` is a battle-tested alternative.

---

### 18. What is destructuring?

"Destructuring is syntax that lets you **unpack** values from arrays or properties from objects into distinct variables.

```js
// Object destructuring
const { name, age = 25 } = user;

// Array destructuring
const [first, , third] = [1, 2, 3];

// Nested
const { address: { city } } = user;
```

I use it constantly — especially for function parameters: `function render({ title, body }) {...}` is cleaner than accessing `props.title` repeatedly."

#### Indepth
Destructuring with renaming: `const { name: userName } = user` — this assigns `user.name` to a variable called `userName`. The `:` is NOT a type annotation here. Destructuring also works in `for...of` loops: `for (const [key, val] of Object.entries(obj))`. Combined with rest: `const { a, ...rest } = obj` collects remaining properties.

---

### 19. What are array methods like `map`, `filter`, `reduce`?

"These are **higher-order array methods** — they accept a callback and return a new value without mutating the original.

- `map`: transforms each element, returns new array of same length
- `filter`: returns new array with only elements where callback returns `true`
- `reduce`: accumulates all elements into a single value

```js
const nums = [1, 2, 3, 4, 5];
const doubled = nums.map(n => n * 2);        // [2,4,6,8,10]
const evens = nums.filter(n => n % 2 === 0); // [2,4]
const sum = nums.reduce((acc, n) => acc + n, 0); // 15
```"

#### Indepth
These methods come from **functional programming**. They are composable — you can chain them: `arr.filter(...).map(...).reduce(...)`. For performance-critical code on large arrays, a single `reduce` or `for` loop is faster than chaining (avoids creating multiple intermediate arrays). The `flatMap()` method combines `map` and `flat(1)` in one pass.

---

### 20. How to remove duplicates from an array?

"The cleanest way is using a **Set**, which only stores unique values:

```js
const arr = [1, 2, 2, 3, 3, 3];
const unique = [...new Set(arr)]; // [1, 2, 3]
```

For arrays of objects (where uniqueness is by a property), I use `filter` with an index check or `reduce` with a Map:

```js
const unique = arr.filter((item, idx, self) =>
  idx === self.findIndex(t => t.id === item.id)
);
```"

#### Indepth
`Set` uses the **SameValueZero** algorithm for comparison (like `===` but treats `+0` and `-0` as equal, and `NaN` as equal to itself). For large arrays, `Set` is O(n) vs `indexOf`-based dedup which is O(n²). The `Map`-based approach for objects is O(n) and is the most performant option.

---

### 21. How does the event loop work?

"JavaScript is **single-threaded** — it has one call stack. The event loop is the mechanism that allows non-blocking async operations.

When an async operation (setTimeout, fetch) completes, its callback is placed in a **queue**. The event loop checks: if the call stack is empty, it picks the next callback from the queue and pushes it to the stack.

There are two queues: the **microtask queue** (Promises, `queueMicrotask`) which is fully drained before the **macrotask queue** (setTimeout, setInterval, I/O)."

#### Indepth
The order of execution: synchronous code → microtasks (all of them) → one macrotask → microtasks → one macrotask... Each macrotask is run one at a time, but ALL microtasks queued during a macrotask are run before the next macrotask. This means `Promise.resolve().then(fn)` always runs before `setTimeout(fn, 0)`.

---

### 22. What is the difference between synchronous and asynchronous code?

"**Synchronous** code executes line by line, blocking the next line until the current one finishes. **Asynchronous** code initiates an operation and moves on, handling the result via callbacks, promises, or async/await when it completes.

In a browser, synchronous JS blocks the UI — a `while(true)` loop will freeze the page. Async operations like `fetch` allow the browser to remain responsive while waiting for the network.

The key insight: async doesn't mean multi-threaded in JS. It's still single-threaded, using the event loop."

#### Indepth
CPU-bound tasks (heavy computation) still block the event loop even with async/await. `async/await` only helps with I/O-bound tasks (network, filesystem). For CPU-intensive work, the solution is **Web Workers** (browser) or **Worker Threads** (Node.js), which run on separate threads.

---

### 23. What are Promises?

"A **Promise** is an object representing the eventual completion (or failure) of an async operation. It has three states: `pending`, `fulfilled`, or `rejected`.

```js
const p = new Promise((resolve, reject) => {
  setTimeout(() => resolve('done'), 1000);
});
p.then(result => console.log(result))
 .catch(err => console.error(err));
```

Promises solved **callback hell** by making async code chainable. `Promise.all`, `Promise.race`, `Promise.allSettled`, and `Promise.any` allow powerful composition of multiple async operations."

#### Indepth
Once settled (resolved or rejected), a Promise's state is **immutable**. `.then()` always runs asynchronously as a microtask, even if the promise is already resolved. This guarantees consistent async behavior. Unhandled promise rejections (no `.catch()`) cause unhandledRejection events — in Node.js, this can crash the process in newer versions.

---

### 24. What is async/await?

"`async/await` is **syntactic sugar** over Promises. An `async` function always returns a Promise. `await` pauses execution inside the async function until the Promise resolves, making async code look synchronous.

```js
async function fetchUser(id) {
  try {
    const res = await fetch(`/api/users/${id}`);
    const user = await res.json();
    return user;
  } catch (err) {
    console.error('Failed:', err);
  }
}
```

I use async/await for all my async code. It's dramatically more readable than `.then()` chains, especially for error handling with `try/catch`."

#### Indepth
`await` can only be used directly inside an `async` function (or at the top level in ES modules with **top-level await**). Under the hood, the async function is compiled into a state machine using Promises and generators. `await` actually suspends the function and yields control back to the event loop, allowing other code to run.

---

### 25. How does error handling with `try...catch` work?

"`try/catch` lets you run code that might throw an error and handle it gracefully without crashing the program.

```js
try {
  const data = JSON.parse(invalidJson); // throws SyntaxError
  processData(data);
} catch (err) {
  console.error(err.name, err.message); // Handle error
} finally {
  cleanup(); // Always runs, whether or not there was an error
}
```

I always include a `finally` block for cleanup (closing connections, hiding loaders) because it runs regardless of success or failure."

#### Indepth
`catch` only catches **synchronous** errors and **awaited** async errors. A `.catch()` is needed for unhandled Promise rejections. The `err` object has `name`, `message`, and `stack` properties. To re-throw only specific errors: `catch (err) { if (!(err instanceof MyError)) throw err; ... }`. In modern environments, `Error.cause` option lets you chain errors: `new Error('wrap', { cause: originalErr })`.

---

### 26. How to select elements in the DOM?

"There are several DOM selection methods:

```js
document.getElementById('myId');           // single element by id
document.querySelector('.myClass');        // first match (CSS selector)
document.querySelectorAll('div.card');     // NodeList of all matches
document.getElementsByClassName('card');  // HTMLCollection (live)
document.getElementsByTagName('p');       // HTMLCollection (live)
```

I almost always use `querySelector`/`querySelectorAll` — they're the most flexible, accepting any CSS selector. The older methods are slightly faster but less versatile."

#### Indepth
`querySelectorAll` returns a **static NodeList** (snapshot, not live). `getElementsBy*` methods return **live HTMLCollections** that update automatically as the DOM changes — which can cause infinite loops if you're iterating and modifying simultaneously. Always convert live collections to arrays: `[...document.getElementsByClassName('x')]`.

---

### 27. How to attach an event listener?

"Use `addEventListener(event, handler)` on any DOM element:

```js
const btn = document.getElementById('submit');
btn.addEventListener('click', function(event) {
  event.preventDefault();
  console.log('Button clicked!');
});
```

I prefer `addEventListener` over inline HTML event attributes (`onclick='...'`) or assigning to `element.onclick` because it allows multiple listeners on the same element and makes it easy to remove listeners with `removeEventListener`."

#### Indepth
The third parameter of `addEventListener` is either a boolean (useCapture) or an options object: `{ once: true, passive: true, capture: false }`. `once: true` automatically removes the listener after the first call. `passive: true` tells the browser the handler won't call `preventDefault()`, allowing scroll performance optimizations for touch events — it's critical for smooth mobile scrolling.

---

### 28. What is the difference between `innerHTML` and `textContent`?

"`textContent` gets/sets the **raw text content** of an element — it's safe from XSS because it doesn't parse HTML.
`innerHTML` gets/sets **parsed HTML** — it can inject and render HTML tags, including `<script>` tags, which makes it an XSS risk.

```js
elem.textContent = '<b>bold</b>'; // Renders as literal text: <b>bold</b>
elem.innerHTML = '<b>bold</b>';   // Renders as: bold (in bold)
```

For untrusted user input, **always use textContent**. Use innerHTML only for trusted, sanitized HTML."

#### Indepth
`innerHTML` triggers a full parse and DOM reconstruction for its subtree — it's slow for frequent updates. `textContent` is fast because it skips parsing. For safe HTML injection, use `element.insertAdjacentHTML('beforeend', safeHtml)` which is slightly faster than `innerHTML` and more flexible. Libraries like **DOMPurify** sanitize HTML before using `innerHTML`.

---

### 29. What is `event.preventDefault()` and `event.stopPropagation()`?

"`event.preventDefault()` cancels the **browser's default behavior** for an event — like stopping a link from navigating, or a form from submitting.
`event.stopPropagation()` stops the event from **bubbling up** to parent elements.

```js
form.addEventListener('submit', (e) => {
  e.preventDefault(); // Prevent form submission/page reload
  validate();
});
child.addEventListener('click', (e) => {
  e.stopPropagation(); // Prevent click from reaching parent
});
```"

#### Indepth
`stopPropagation` stops bubbling but not other listeners on the same element. `stopImmediatePropagation()` stops ALL listeners (including siblings) on the same element. There's also `event.cancelable` to check if `preventDefault()` has any effect. For passive event listeners (e.g., scroll), calling `preventDefault()` throws a warning and is ignored.

---

### 30. How do you manipulate classes using JavaScript?

"Use the `classList` API — it's the modern, clean way to add, remove, and toggle CSS classes:

```js
const el = document.querySelector('.card');
el.classList.add('active');
el.classList.remove('hidden');
el.classList.toggle('expanded');        // adds if absent, removes if present
el.classList.contains('active');        // true/false check
el.classList.replace('old', 'new');     // replace one class with another
```

I never manipulate `element.className` as a string — it's error-prone and overwrites all existing classes."

#### Indepth
`classList.toggle(className, force)` — the optional second boolean argument forces add (`true`) or remove (`false`) regardless of current state. This is useful for: `el.classList.toggle('dark', prefersDark)`. `classList` operations are more performant than string manipulation because they don't trigger unnecessary DOM parsing.
