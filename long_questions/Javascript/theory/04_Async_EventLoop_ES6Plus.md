# 🔴 **JavaScript Async, Event Loop & ES6+ — Senior Level**

> **Target Companies:** Infosys (Senior), Accenture FSI, IBM (Service-Based) + Razorpay, PhonePe, Flipkart, Meesho, Juspay (Product-Based)

---

### 81. What is the Temporal Dead Zone in JavaScript?

"The **Temporal Dead Zone (TDZ)** is the period between when a `let` or `const` variable's **scope begins** and when its **declaration is reached** during execution — accessing the variable in this zone causes a `ReferenceError`.

```js
{
  console.log(x); // ✅ undefined (var is hoisted & initialized)
  console.log(y); // ❌ ReferenceError: Cannot access 'y' before initialization

  var x = 10;
  let y = 20; // TDZ for 'y' ends here
}
```

The TDZ makes `let`/`const` safer than `var` by turning silent bugs (undefined) into loud errors (ReferenceError)."

#### Indepth
Technically, `let`/`const` ARE hoisted — they are registered in the scope at compilation time. But their binding is explicitly left uninitialized. The TDZ is the gap between scope entry and initialization. This matters for `typeof`: `typeof undeclaredVar` returns `'undefined'`, but `typeof tdz-let-var` inside TDZ throws a `ReferenceError` — a subtle difference from truly undeclared variables.

---

### 82. How does `Object.freeze()` and `Object.seal()` differ?

"`Object.freeze()` makes an object **completely immutable** — cannot add, remove, or modify properties.
`Object.seal()` prevents adding/removing properties but allows **modifying** existing ones.

```js
const frozen = Object.freeze({ a: 1, b: { c: 2 } });
frozen.a = 99;         // silently fails (or throws in strict mode)
frozen.d = 4;          // silently fails
frozen.b.c = 99;       // ✅ WORKS! Freeze is shallow!

const sealed = Object.seal({ x: 10 });
sealed.x = 99;         // ✅ allowed (modify)
sealed.y = 20;         // ❌ silently fails (can't add)
delete sealed.x;       // ❌ silently fails (can't delete)
```"

#### Indepth
Both operations are **shallow** — nested objects are not frozen/sealed. For deep freeze:
```js
function deepFreeze(obj) {
  Object.getOwnPropertyNames(obj).forEach(name => {
    if (typeof obj[name] === 'object') deepFreeze(obj[name]);
  });
  return Object.freeze(obj);
}
```
`Object.isFrozen()` and `Object.isSealed()` check the state. In strict mode, mutations throw `TypeError` instead of silently failing. Redux uses frozen objects in dev mode to catch accidental mutations.

---

### 83. What is the purpose of `use strict`?

"`'use strict'` opts your code into a **stricter mode** that catches common mistakes and disables problematic features:

```js
'use strict';

x = 10;                     // ❌ ReferenceError — undeclared variable
delete Object.prototype;    // ❌ TypeError
function f() { return this; } f(); // undefined (not global object)
var obj = { a: 1, a: 2 };  // ❌ SyntaxError (duplicate keys)
arguments.caller;           // ❌ TypeError
```

Benefits: catches silent errors early, prevents accidental globals, enables safer syntax for future JS features."

#### Indepth
ES6 **modules** are always in strict mode — you don't need the directive. ES6 **class bodies** are also always strict. `'use strict'` can be applied per-file or per-function. In strict mode, `this` inside a regular function call is `undefined` instead of the global object — a critical difference for understanding `this` binding. Strict mode also restricts `with` statement and `eval`'s ability to introduce new variables.

---

### 84. How do symbols work in JavaScript?

"**Symbols** are a primitive type that creates **guaranteed unique** values — they're primarily used as object property keys to avoid name collisions.

```js
const id = Symbol('description');
const obj = { [id]: 123, name: 'Dhruv' };

console.log(id.toString());   // 'Symbol(description)'
console.log(obj[id]);         // 123
console.log(id === Symbol('description')); // false — always unique!

// Global symbols (shared across modules)
const globalId = Symbol.for('id');
Symbol.keyFor(globalId); // 'id'
```"

#### Indepth
Symbol-keyed properties are **not enumerable** in `Object.keys()`, `for...in`, or `JSON.stringify()` — they're truly hidden. Only `Object.getOwnPropertySymbols()` or `Reflect.ownKeys()` reveals them. Well-known symbols are predefined: `Symbol.iterator` (makes objects iterable), `Symbol.toPrimitive` (custom type conversion), `Symbol.hasInstance` (customizes `instanceof`). Libraries use symbols for internal properties to avoid clashing with user code.

---

### 85. What are the phases of execution in JS?

"JavaScript execution has two main phases:

1. **Creation/Compilation Phase**: The engine scans the code, creates the **execution context**, and sets up the **scope chain**. Variable declarations (`var`) are hoisted and initialized to `undefined`. Function declarations are fully hoisted.

2. **Execution Phase**: Code runs line by line. Values are assigned, functions are called, pushing new execution contexts onto the **call stack**.

```
Call Stack:
┌──────────────┐
│ inner()      │  ← currently executing
├──────────────┤
│ outer()      │
├──────────────┤
│ global       │
└──────────────┘
```"

#### Indepth
Each **Execution Context** contains: a Variable Environment (for `var`), a Lexical Environment (for `let`/`const`), and a `this` binding. The **scope chain** is established during creation. **Closures** capture the Lexical Environment. When a function returns, its context is popped from the stack — but its Lexical Environment survives if a closure holds a reference to it.

---

### 86. How does the call stack work in JavaScript?

"The **call stack** is a LIFO data structure that tracks function execution. Each function call pushes a **stack frame**; when the function returns, its frame is popped.

```js
function a() { b(); }
function b() { c(); }
function c() { console.trace(); } // Shows stack trace

a();
// Call stack: a → b → c (c is top/current)
```

**Stack overflow** occurs when recursion is too deep — the stack runs out of space. This is why infinite recursion throws a `RangeError: Maximum call stack size exceeded`."

#### Indepth
V8's default stack size is ~15,000–20,000 frames. You can increase it in Node.js with `--stack-size=65536`. **Tail call optimization (TCO)** would allow recursive functions that tail-call themselves to reuse the current frame — the ES6 spec mandates TCO, but only Safari fully implements it. Alternative for deep recursion: **trampolining** (converting recursion to iteration with a while loop driving a thunk).

---

### 87. What is the difference between `apply()`, `call()`, and `bind()`?

"All three control what `this` refers to inside a function:

- `call(thisArg, arg1, arg2)` — calls immediately with args passed **individually**
- `apply(thisArg, [arg1, arg2])` — calls immediately with args as **array**
- `bind(thisArg, arg1, arg2)` — returns a **new function** with `this` permanently bound

```js
function greet(greeting, punctuation) {
  return `${greeting}, ${this.name}${punctuation}`;
}
const user = { name: 'Dhruv' };

greet.call(user, 'Hello', '!');    // 'Hello, Dhruv!'
greet.apply(user, ['Hi', '?']);    // 'Hi, Dhruv?'
const boundGreet = greet.bind(user, 'Hey');
boundGreet('...');                 // 'Hey, Dhruv...'
```"

#### Indepth
`bind` is used heavily in React class components: `this.handleClick = this.handleClick.bind(this)` in constructor. Arrow functions don't need this — they inherit `this` lexically. `bind` also supports **partial application** — pre-filling arguments. Internally, `bind` creates a new function object with a fixed `[[BoundThis]]` internal slot. Arrow functions cannot be re-bound even with `call`/`apply` — the `thisArg` is silently ignored.

---

### 88. What is function composition?

"**Function composition** is creating a new function by combining multiple functions where the output of one becomes the input of the next.

```js
// compose: applies right-to-left
const compose = (...fns) => x => fns.reduceRight((acc, fn) => fn(acc), x);

// pipe: applies left-to-right (more readable)
const pipe = (...fns) => x => fns.reduce((acc, fn) => fn(acc), x);

const double = x => x * 2;
const addOne = x => x + 1;
const square = x => x * x;

const transform = pipe(double, addOne, square);
transform(3); // double(3)=6 → addOne(6)=7 → square(7)=49
```"

#### Indepth
Composition is the core of **functional programming** and point-free style. Libraries like **Ramda** and **lodash/fp** provide curried, composable utilities. In React, HOCs and hooks are forms of function composition. The **pipe operator** (`|>`) is a TC39 proposal that would enable: `3 |> double |> addOne |> square`. `RxJS` uses composition for observable transformations.

---

### 89. What is `Symbol.iterator`?

"`Symbol.iterator` is a **well-known symbol** that defines the default iterator for an object. Any object with `[Symbol.iterator]()` that returns an iterator becomes **iterable** — usable with `for...of`, spread, and destructuring.

```js
class Range {
  constructor(start, end) { this.start = start; this.end = end; }

  [Symbol.iterator]() {
    let current = this.start;
    const end = this.end;
    return {
      next() {
        return current <= end
          ? { value: current++, done: false }
          : { value: undefined, done: true };
      }
    };
  }
}

for (const n of new Range(1, 5)) console.log(n); // 1 2 3 4 5
[...new Range(1, 3)]; // [1, 2, 3]
```"

#### Indepth
Built-in iterables: `Array`, `String`, `Map`, `Set`, `NodeList`, `arguments` — they all implement `[Symbol.iterator]`. Generators automatically implement both the iterator and iterable protocols. `Symbol.asyncIterator` is the async version — used with `for await...of` for async data streams (e.g., streaming HTTP responses in Node.js).

---

### 90. What are private class fields in JavaScript?

"Private class fields use the `#` prefix and are **truly private** — inaccessible from outside the class. This is a hard-enforcement by the JavaScript engine, unlike the old `_` naming convention.

```js
class BankAccount {
  #balance = 0; // private field
  #transactionLog = []; // private

  deposit(amount) {
    if (amount <= 0) throw new Error('Invalid amount');
    this.#balance += amount;
    this.#log(`Deposited ${amount}`);
  }

  #log(msg) { // private method
    this.#transactionLog.push({ msg, time: Date.now() });
  }

  get balance() { return this.#balance; }
}

const acc = new BankAccount();
acc.deposit(100);
acc.#balance; // ❌ SyntaxError — genuinely private
```"

#### Indepth
Private fields are stored in a special **slot** associated with the class, not on the prototype or instance directly. They are not enumerable, not JSON-serializable, not accessible via `Object.keys()` or `Proxy`. `#` fields must be declared at the class level before use. You can check if an instance has a private field using `#field in obj` (ES2022 ergonomic brand checks): `if (#balance in obj) {...}`.

---

### 91. What is top-level await?

"**Top-level await** allows using `await` directly in the top-level of an **ES module** file — without wrapping in an async function.

```js
// module.js (ES module with top-level await)
const config = await fetch('/config.json').then(r => r.json());
const db = await connectDB(config.dbUrl);

export { db, config };
```

This allows modules to do async initialization before their exports are used by importing modules. The importing module will pause until the awaited value resolves."

#### Indepth
Top-level await only works in **ES modules** (files with `type: 'module'` or `.mjs`). It makes the entire module behave as a promise — importing modules wait for it to resolve. This can introduce **deadlocks** if module A awaits module B which imports module A (circular dependency with top-level await). It's commonly used for config loading, database connections, and feature flag fetching at module startup.

---

### 92. What is `BigInt` in JavaScript?

"`BigInt` is a built-in type for representing integers **larger than** `Number.MAX_SAFE_INTEGER` (2⁵³ - 1) with **arbitrary precision**.

```js
const big = 9007199254740991n;      // BigInt literal
const bigger = BigInt('9007199254740992');

big + 1n;           // 9007199254740992n
big * 2n;           // 18014398509481982n

// Cannot mix BigInt and Number
big + 1;            // ❌ TypeError
big + BigInt(1);    // ✅

Number(big);        // loses precision for very large values
typeof big;         // 'bigint'
```"

#### Indepth
BigInt is important for: **cryptographic algorithms** (large key math), **financial calculations** (avoiding floating-point errors), and working with **64-bit integer IDs** from APIs (which exceed `MAX_SAFE_INTEGER`). BigInt cannot be serialized to JSON natively (throws TypeError) — use `JSON.stringify(obj, (_, v) => typeof v === 'bigint' ? v.toString() : v)`. The **Temporal API** uses BigInt internally for nanosecond precision timestamps.

---

### 93. What is `Promise.allSettled()` vs `Promise.all()`?

"`Promise.all()` **fails fast** — if ANY promise rejects, the entire thing rejects immediately and other results are ignored.
`Promise.allSettled()` **waits for all** promises to settle (resolve or reject) and returns an array of status objects.

```js
const [p1, p2, p3] = [
  fetch('/api/user'),     // resolves
  fetch('/api/missing'),  // rejects with 404
  fetch('/api/settings'), // resolves
];

// all() — fails on first rejection
await Promise.all([p1, p2, p3]); // ❌ Throws on p2 rejection

// allSettled() — all results
const results = await Promise.allSettled([p1, p2, p3]);
results.forEach(r => {
  if (r.status === 'fulfilled') use(r.value);
  if (r.status === 'rejected') logError(r.reason);
});
```"

#### Indepth
When to use each: `Promise.all` — when you need ALL results and partial failure means the operation failed (like a multi-part form submission). `Promise.allSettled` — when you want partial results and should handle each independently (like a dashboard loading multiple widgets). `Promise.any` resolves on the first fulfillment; `Promise.race` resolves/rejects on the first settlement. These four cover all concurrency patterns.

---

### 94. What is optional chaining and how does it work?

"Optional chaining (`?.`) provides a safe way to access deeply nested object properties without throwing if an intermediate value is `null` or `undefined`. It **short-circuits** to `undefined` instead.

```js
const user = { profile: null };

// Without optional chaining
const bio = user && user.profile && user.profile.bio; // verbose

// With optional chaining
const bio = user?.profile?.bio;              // undefined (safe)
const age = user?.getAge?.();               // safe method call
const tag = user?.tags?.[0];               // safe array access
const name = user?.name ?? 'Anonymous';    // with nullish coalescing
```"

#### Indepth
Optional chaining is **not the same** as falsy checking. `obj?.prop` only short-circuits on `null`/`undefined` — not on `0`, `''`, or `false`. It also **doesn't swallow errors** — if `user.getAge()` throws an error inside the function, the `?.` won't catch it; it only prevents the `TypeError: Cannot read properties of null`. TypeScript uses `?.` heavily in type narrowing and the `strict` mode encourages its use.

---

### 95. What is TypeScript and how does it extend JavaScript?

"**TypeScript** is a typed superset of JavaScript maintained by Microsoft. All valid JS is valid TypeScript. It adds:
- **Static type annotations**: `const name: string = 'Dhruv'`
- **Interfaces and type aliases**: `interface User { name: string; age: number; }`
- **Generics**: `function identity<T>(arg: T): T { return arg; }`
- **Access modifiers**: `private`, `protected`, `readonly`
- **Strict null checks**: prevents `null`/`undefined` errors

TypeScript **compiles to JavaScript** — it's erased at runtime. The type system is purely a development-time tool for better tooling and catching bugs early."

#### Indepth
TypeScript uses **structural typing** (duck typing) — if an object has the required shape, it satisfies the type, regardless of its declared type. This differs from Java's **nominal typing**. TypeScript's type system is **Turing-complete** — you can compute types at compile time. Key features: `Union types` (`string | number`), `Intersection types` (`A & B`), `Conditional types` (`T extends U ? X : Y`), `Mapped types`, `Template literal types`. In 2024, TypeScript is essentially mandatory for large JS codebases.

---

### 96. What is the output of `[] + []` or `{} + []`?

"This tests understanding of type coercion:

```js
[] + []         // '' (both coerce to empty string, then concat)
[] + {}         // '[object Object]' (array→'', object→'[object Object]')
{} + []         // 0  (in statement position, {} is empty block, then +[] = 0)
({}) + []       // '[object Object]' (force {} as expression)
+[]             // 0 (unary +, [] coerces to '', Number('') = 0)
+{}             // NaN (Number('[object Object]') = NaN)
```

These are classic coercion gotchas because `+` operator tries string concatenation OR addition depending on operand types."

#### Indepth
The rules: The `+` operator calls `ToPrimitive` on both sides. Arrays call `toString()` which joins elements (empty → `''`). Objects call `valueOf()` (returns object, not primitive) then `toString()` (returns `'[object Object]'`). If either side is a string after coercion, string concatenation occurs. The `{} + []` ambiguity is specific to the REPL/global context where `{}` at statement start is parsed as a block, not an object literal.

---

### 97. How does `this` behave in arrow functions?

"Arrow functions **do not have their own `this`** — they inherit `this` from their **enclosing lexical scope** at definition time. It cannot be changed with `call`, `apply`, or `bind`.

```js
class Timer {
  constructor() { this.seconds = 0; }

  start() {
    // Arrow function: 'this' is Timer instance (lexical)
    setInterval(() => {
      this.seconds++; // ✅ correct 'this'
    }, 1000);

    // Regular function: 'this' would be undefined in strict mode
    setInterval(function() {
      this.seconds++; // ❌ 'this' is undefined/global
    }, 1000);
  }
}
```"

#### Indepth
This is the primary reason arrow functions were introduced — to solve the callback `this` problem that plagued pre-ES6 code. Before arrows, the pattern was `var self = this` or `.bind(this)`. Arrow functions are also useful in: array method callbacks (`map`, `filter`), Promise chains, and React functional component handlers. But avoid arrow functions as **object literal methods** — the `this` would be the outer scope (usually `window`), not the object.

---

### 98. What are tagged template literals?

"Tagged templates let you process a template literal with a **function** (the 'tag'). The function receives the string parts and interpolated values separately.

```js
function highlight(strings, ...values) {
  return strings.reduce((result, str, i) =>
    result + str + (values[i] ? `<mark>${values[i]}</mark>` : ''), '');
}

const name = 'Dhruv';
const score = 95;
highlight`Player: ${name} scored ${score} points!`;
// 'Player: <mark>Dhruv</mark> scored <mark>95</mark> points!'
```

Real-world uses: `gql\`...\`` (GraphQL), `css\`...\`` (CSS-in-JS), `html\`...\`` (safe HTML templating), `sql\`...\`` (parameterized SQL)."

#### Indepth
The tag function receives `strings` (a frozen array of string parts with `strings.raw` for unescaped strings) and rest `values` (the interpolated expressions). `strings.length === values.length + 1` always. Tagged templates enable **DSLs** (domain-specific languages) inside JavaScript. The `String.raw` built-in tag is used to get raw strings (backslashes not processed): `` String.raw`\n\t` `` returns `\\n\\t` instead of a newline + tab.

---

### 99. How do you implement your own `bind` function?

"Implementing `bind` manually:

```js
Function.prototype.myBind = function(thisArg, ...outerArgs) {
  const fn = this; // the original function

  return function(...innerArgs) {
    // Respect 'new' invocation — 'this' from 'new' takes priority
    const context = this instanceof boundFn ? this : thisArg;
    return fn.apply(context, [...outerArgs, ...innerArgs]);
  };

  // For instanceof support: set prototype
  function boundFn() {}
  boundFn.prototype = fn.prototype;
  return boundFn;
};

// Simpler version for most use cases:
Function.prototype.myBind = function(thisArg, ...preset) {
  const fn = this;
  return (...args) => fn.apply(thisArg, [...preset, ...args]);
};
```"

#### Indepth
The full native `bind` must handle: 1) Partial application (pre-filling args), 2) When the bound function is used with `new` — the `thisArg` is ignored and `this` is the new instance. The simplified version using an arrow function doesn't handle the `new` case correctly because arrows can't be constructors. The `length` property of a bound function should be `originalFn.length - presetArgs.length` (clamped to 0).

---

### 100. Explain `Promise.all`, `Promise.race`, and `Promise.any`.

"These static methods compose multiple promises:

```js
const p1 = fetch('/api/a').then(r => r.json());
const p2 = fetch('/api/b').then(r => r.json());
const p3 = fetch('/api/c').then(r => r.json());

// Promise.all — ALL must resolve; fails on first rejection
const [a, b, c] = await Promise.all([p1, p2, p3]);

// Promise.race — first to SETTLE (resolve OR reject) wins
const fastest = await Promise.race([p1, p2, p3]);

// Promise.any — first to RESOLVE wins; fails only if ALL reject
const firstSuccess = await Promise.any([p1, p2, p3]);
// throws AggregateError if all reject

// Promise.allSettled — waits for ALL, never rejects
const results = await Promise.allSettled([p1, p2, p3]);
```"

#### Indepth
| Method | Fulfills when | Rejects when |
|---|---|---|
| `all` | ALL resolve | ANY rejects |
| `race` | First settles | First settles |
| `any` | First resolves | ALL reject |
| `allSettled` | ALL settle | Never |

`Promise.any` (ES2021) fills the gap: "I want the first success, and I don't fail unless everything fails." Use `Promise.race` for **timeout patterns**: `Promise.race([fetch(url), timeout(5000)])` where `timeout` rejects after 5 seconds.
