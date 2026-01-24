# Javascript Interview Questions & Answers

## 🔹 1. Beginner (Questions 1-30)

### Syntax & Basics

**Q1: What are the different data types in JavaScript?**
JavaScript has 8 data types:
1.  **Primitive**: `String`, `Number`, `BigInt`, `Boolean`, `Undefined`, `Null`, `Symbol`.
2.  **Non-Primitive**: `Object` (includes Arrays, Functions).

**Q2: What is the difference between `var`, `let`, and `const`?**
*   `var`: Function-scoped, can be redeclared/updated, hoisted with `undefined`.
*   `let`: Block-scoped, can be updated but not redeclared, hoisted to TDZ (Temporal Dead Zone).
*   `const`: Block-scoped, cannot be updated or redeclared, must be initialized during declaration.

**Q3: What is hoisting in JavaScript?**
Hoisting is a mechanism where variable and function declarations are moved to the top of their scope before code execution.
*   Functions are fully hoisted.
*   `var` variables are hoisted with `undefined`.
*   `let` and `const` are hoisted but stay in the TDZ until initialized.

**Q4: How does type coercion work in JavaScript?**
Type coercion is the automatic conversion of values from one data type to another (e.g., string to number). It happens in:
*   **Implicit Coercion**: `1 + '2'` → `'12'` (Number to String).
*   **Explicit Coercion**: `Number('5')`, `String(123)`.

**Q5: What are truthy and falsy values?**
*   **Falsy**: `false`, `0`, `-0`, `""` (empty string), `null`, `undefined`, `NaN`, `0n` (BigInt).
*   **Truthy**: Everything else (e.g., `"0"`, `[]`, `{}`, `function(){}`).

**Q6: What is the difference between `==` and `===`?**
*   `==` (Loose quality): Performs type coercion before comparison (e.g., `5 == '5'` is `true`).
*   `===` (Strict equality): Checks both value and type (e.g., `5 === '5'` is `false`).

**Q7: Explain the difference between `null` and `undefined`.**
*   `undefined`: A variable has been declared but not assigned a value.
*   `null`: An assignment value that represents "no value" or "empty".
*   `typeof undefined` is `"undefined"`, while `typeof null` is `"object"` (a known JS bug).

**Q8: How does `typeof` work?**
It returns a string indicating the type of the operand.
```javascript
typeof 42; // "number"
typeof "hello"; // "string"
typeof true; // "boolean"
typeof null; // "object"
typeof []; // "object"
```

**Q9: What are template literals in JavaScript?**
Enclosed by backticks (`` ` ``), they allow embedded expressions (`${expression}`) and multi-line strings.
```javascript
const name = "John";
console.log(`Hello, ${name}`);
```

**Q10: How does the `switch` statement work?**
It evaluates an expression and executes code matching a `case` clause. It uses strict equality (`===`).
```javascript
switch(x) {
  case 1: // code
    break;
  default: // code
}
```

### Functions & Scope

**Q11: What is the difference between function declaration and function expression?**
*   **Declaration**: `function foo() {}`. Hoisted completely.
*   **Expression**: `const foo = function() {}`. Only the variable is hoisted (if `var`, `undefined`; if `const`, TDZ), not the function definition.

**Q12: What is a callback function?**
A function passed as an argument to another function, which is then invoked inside the outer function to complete some kind of routine or action.

**Q13: What are arrow functions?**
A shorter syntax for writing functions using `=>`. They do not have their own `this`, `arguments`, or `super`.
```javascript
const add = (a, b) => a + b;
```

**Q14: What is lexical scoping?**
A function's scope is determined by its physical location in the source code. Inner functions have access to variables declared in their outer scope.

**Q15: What is the use of closures in JavaScript?**
A closure is a function that remembers its outer variables and can access them even when the outer function has finished execution. Used for data privacy, factories, etc.

**Q16: How do default parameters work?**
Allow initializing named parameters with default values if no value or `undefined` is passed.
```javascript
function greet(name = "User") { ... }
```

### Arrays & Objects

**Q17: How to clone an object in JavaScript?**
*   Shallow copy: `Object.assign({}, obj)` or spread `{...obj}`.
*   Deep copy: `JSON.parse(JSON.stringify(obj))` or `structuredClone(obj)`.

**Q18: What is destructuring?**
Unpacking values from arrays or properties from objects into distinct variables.
```javascript
const { name, age } = person;
const [first, second] = numbers;
```

**Q19: What are array methods like `map`, `filter`, `reduce`?**
*   `map`: Creates a new array by applying a function to each element.
*   `filter`: Creates a new array with elements that pass the test.
*   `reduce`: Reduces the array to a single value (accumulator).

**Q20: How to remove duplicates from an array?**
Using `Set`:
```javascript
const unique = [...new Set(array)];
```

### Control Flow

**Q21: How does the event loop work?**
It monitors the Call Stack and Callback Queue. If the Call Stack is empty, it pushes the first event from the queue to the stack.

**Q22: What is the difference between synchronous and asynchronous code?**
*   **Synchronous**: Executed line-by-line, blocking.
*   **Asynchronous**: Executed later (e.g., via callbacks, promises), non-blocking.

**Q23: What are promises?**
Objects representing the eventual completion (or failure) of an asynchronous operation. States: Pending, Fulfilled, Rejected.

**Q24: What is async/await?**
Syntactic sugar over Promises. `async` makes a function return a Promise, `await` pauses execution until the Promise resolves.

**Q25: How does error handling with `try...catch` work?**
Code in `try` block is executed. If an error occurs, control shifts to `catch` block.
```javascript
try {
  // risky code
} catch (error) {
  // handle error
}
```

### DOM Basics

**Q26: How to select elements in the DOM?**
*   `document.getElementById('id')`
*   `document.querySelector('.class')`
*   `document.querySelectorAll('tag')`

**Q27: How to attach an event listener?**
```javascript
element.addEventListener('click', function() { ... });
```

**Q28: What is the difference between `innerHTML` and `textContent`?**
*   `innerHTML`: Parses content as HTML (can result in XSS).
*   `textContent`: Treats content as raw text (safer).

**Q29: What is `event.preventDefault()` and `event.stopPropagation()`?**
*   `preventDefault()`: Stops the browser's default behavior (e.g., form submission).
*   `stopPropagation()`: Stops the event from bubbling up the DOM tree.

**Q30: How do you manipulate classes using JavaScript?**
```javascript
element.classList.add('active');
element.classList.remove('active');
element.classList.toggle('active');
```

## 🔹 2. Intermediate (Questions 31-70)

### Advanced Functions & Closures

**Q31: What is a pure function?**
A function that always returns the same result given the same arguments and has no side effects (doesn't modify external state).

**Q32: What is memoization?**
An optimization technique used to speed up function calls by caching the results of expensive function calls and returning the cached result when the same inputs occur again.

**Q33: What are IIFEs (Immediately Invoked Function Expressions)?**
Functions that are executed as soon as they are defined. Used to create a local scope and avoid polluting the global namespace.
```javascript
(function() {
  // code
})();
```

**Q34: How do closures preserve data?**
Closures retain references to the variables in their lexical scope even after the outer function has returned, effectively "preserving" that data.

**Q35: What is currying?**
The process of checking a function that takes multiple arguments into a sequence of functions that each take a single argument.
```javascript
const add = a => b => a + b;
add(1)(2); // 3
```

### Objects & Prototypes

**Q36: What is the prototype chain?**
The mechanism by which objects allow inheritance. If a property isn't found on an object, JS looks up its prototype, and so on, until it reaches `null`.

**Q37: What is the difference between `__proto__` and `prototype`?**
*   `__proto__`: The actual object that is used in the lookup chain to resolve methods, etc. (instance level).
*   `prototype`: The object that is used to build `__proto__` when you create an object with `new` (function/constructor level).

**Q38: What is `Object.create()` used for?**
It creates a new object, using an existing object as the prototype of the newly created object.

**Q39: How does inheritance work in JavaScript?**
Primarily through the prototype chain. Objects inherit properties and methods from their prototype. ES6 `class...extends` syntax is sugar over this.

**Q40: What are ES6 classes?**
Syntactic sugar over JavaScript's existing prototype-based inheritance. They provide a cleaner syntax to create objects and deal with inheritance.
```javascript
class Animal {
  constructor(name) { this.name = name; }
}
```

### ES6+ Features

**Q41: What are generators in JavaScript?**
Functions that can be paused and resumed. They are declared with `function*` and use `yield` to return values.

**Q42: What is the spread operator?**
The `...` operator expands an iterable (like an array) into individual elements.
```javascript
const arr = [1, 2, 3];
const newArr = [...arr, 4, 5];
```

**Q43: What is the rest parameter?**
The `...` operator allows a function to accept an indefinite number of arguments as an array.
```javascript
function sum(...args) { return args.reduce((a, b) => a + b); }
```

**Q44: How do modules (`import`/`export`) work?**
ES6 Modules allow splitting code into separate files. `export` exposes variables/functions, and `import` brings them into another file.

**Q45: What is optional chaining?**
The `?.` operator permits reading the value of a property located deep within a chain of connected objects without validation.
```javascript
const user = {};
console.log(user?.address?.street); // undefined (no error)
```

### Async & Event Loop

**Q46: What is the microtask queue vs macrotask queue?**
*   **Microtasks**: High priority (Promises, `queueMicrotask`, MutationObserver). Executed immediately after the current script and before rendering.
*   **Macrotasks**: Lower priority (setTimeout, setInterval, I/O). Executed one per loop iteration.

**Q47: How does `setTimeout` actually work?**
It registers a callback to be executed after a minimum delay. The callback is pushed to the Macrotask Queue and executed when the Call Stack is empty.

**Q48: What is a race condition? How do you avoid it?**
Occurs when the behavior of software depends on the timing of uncontrollable events (like network requests). Avoid by using `Promise.all`, proper synchronization, or ensure operations happen serially if needed.

**Q49: How does JavaScript handle concurrency?**
Via the Event Loop. JS is single-threaded but offloads async operations (I/O, timers) to the browser/node APIs, handling callbacks when they complete.

**Q50: What is a debounce vs throttle function?**
*   **Debounce**: Ensures a function is not called again until a certain amount of time has passed since the last call (e.g., search input).
*   **Throttle**: Ensures a function is called at most once in a specified time period (e.g., scroll events).

### Error Handling & Debugging

**Q51: What are common JavaScript errors?**
`ReferenceError` (var not found), `TypeError` (wrong type operation), `SyntaxError` (invalid code), `RangeError` (number out of range).

**Q52: How to handle errors globally?**
*   Browser: `window.onerror` or `window.addEventListener('error')`.
*   Unhandled Promises: `window.addEventListener('unhandledrejection')`.

**Q53: How does `try...catch` with `finally` work?**
`finally` block executes regardless of whether an error was thrown or caught. Useful for cleanup.
```javascript
try { ... } catch (e) { ... } finally { /* cleanup */ }
```

**Q54: What are custom errors?**
Classes extending the built-in `Error` class to create specific error types.
```javascript
class ValidationError extends Error {
  constructor(message) { super(message); this.name = "ValidationError"; }
}
```

### Data Structures in JS

**Q55: What is a Set and how is it different from an array?**
A collection of values where each value must be unique. Unlike arrays, sets do not use indexes.

**Q56: What is a Map?**
A collection of keyed data items, similar to an Object. However, `Map` allows keys of any type (including objects) and maintains insertion order.

**Q57: How to implement a stack/queue using JavaScript?**
*   **Stack** (LIFO): Use Array `push()` and `pop()`.
*   **Queue** (FIFO): Use Array `push()` and `shift()`.

**Q58: What are WeakMap and WeakSet?**
Collections that hold weak references to objects (keys in WeakMap, values in WeakSet). They do not prevent garbage collection if the key/value is not referenced elsewhere.

### Date & Time

**Q59: How do you format a date in JavaScript?**
*   `Date.prototype.toLocaleDateString()`
*   `Intl.DateTimeFormat`
*   Libraries like date-fns or Moment.js (legacy).

**Q60: How to get the difference between two dates?**
Subtract date objects to get the difference in milliseconds.
```javascript
const diff = date2 - date1;
const days = diff / (1000 * 60 * 60 * 24);
```

### Regular Expressions

**Q61: How do regular expressions work in JavaScript?**
Patterns used to match character combinations in strings. Created with `/pattern/flags` or `new RegExp()`. Used with methods like `test()`, `exec()`, `match()`.

**Q62: How to validate an email using regex?**
A simple (non-exhaustive) pattern:
```javascript
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
emailRegex.test("test@example.com");
```

### Miscellaneous

**Q63: What is NaN?**
"Not-a-Number". A special property of the global object representing a value that is not a legal number. `typeof NaN` is `"number"`.

**Q64: What is the difference between `parseInt()` and `Number()`?**
*   `parseInt()`: Parses a string and returns an integer. Stops at the first non-digit character (e.g., `parseInt("10px")` -> `10`).
*   `Number()`: Converts the entire value to a number. Returns `NaN` if mixed content (e.g., `Number("10px")` -> `NaN`).

**Q65: How do you deep clone an object?**
*   Modern: `structuredClone(obj)`.
*   Legacy: `JSON.parse(JSON.stringify(obj))` (has limitations like Functions/Dates).
*   Libraries: Lodash `_.cloneDeep(obj)`.

## 🔹 3. Advanced (Questions 71-100)

### Event Delegation & Propagation

**Q71: How does event bubbling/capturing work?**
*   **Capturing (Down)**: Event goes from Window → Target.
*   **Bubbling (Up)**: Event goes from Target → Window.
*   By default, listeners are registered in the bubbling phase. Use `{ capture: true }` for capturing.

**Q72: What is event delegation?**
Attaching a single event listener to a parent element to manage events for all its children (existing and future). Uses event bubbling.
```javascript
document.getElementById('list').addEventListener('click', e => {
  if(e.target.matches('li')) { /* ... */ }
});
```

**Q73: What are synthetic events in React (JS background)?**
(JS Context): Wrappers around the browser's native events to ensure cross-browser consistency. React reuses these event objects for performance.

### Memory & Performance

**Q74: How does garbage collection work in JS?**
Typically uses "Mark-and-Sweep". It marks "reachable" objects (from roots like global vars, stack) and sweeps (deletes) unreachable ones.

**Q75: What are memory leaks and how to avoid them?**
Unintentional retention of objects in memory. Causes: Global variables, forgotten timers/listeners, closures holding large scopes.

**Q76: How to profile JavaScript performance?**
Use browser DevTools (Performance tab) to record runtime performance, analyze the flame chart, and identify bottlenecks (long tasks, layout thrashing, excessive GC).

### Browser APIs

**Q77: What is the Fetch API?**
A modern, Promise-based API for making network requests (replacing XMLHttpRequest).
```javascript
fetch(url).then(res => res.json()).then(data => console.log(data));
```

**Q78: How does `localStorage` differ from `sessionStorage`?**
*   `localStorage`: Persists even after browser is closed (no expiration).
*   `sessionStorage`: Persists only for the page session (cleared when tab is closed).

**Q79: How does the Clipboard API work?**
Async API to read/write to system clipboard. Requires user permission/interaction.
```javascript
navigator.clipboard.writeText("Hello");
```

**Q80: What is `requestAnimationFrame`?**
Tells the browser you want to perform an animation. The browser calls the callback before the next repaint (usually 60fps), optimizing performance and battery.

### Security

**Q81: What is XSS and how do you prevent it in JS?**
**Cross-Site Scripting**: Attacker injects malicious scripts.
Prevent by: Escaping/sanitizing user input, using `textContent` instead of `innerHTML`, and CSP (Content Security Policy).

**Q82: What is CSRF and how does JavaScript help mitigate it?**
**Cross-Site Request Forgery**: Attacker tricks user into performing unwanted actions.
Mitigate by: Anti-CSRF tokens (sent in headers) and SameSite cookie attributes.

**Q83: How do you sanitize user input?**
Remove unsafe characters (tags, scripts). Use libraries like DOMPurify or built-in browser APIs (if careful).

### Design Patterns in JS

**Q84: What is the module pattern?**
Using closures to create private scope and expose only public API parts.
```javascript
const Module = (() => {
  let privateVar = 1;
  return { publicMethod: () => privateVar };
})();
```

**Q85: How does the singleton pattern work in JS?**
Ensures a class has only one instance.
```javascript
const Singleton = { /* object literal is a singleton */ };
```
Or checking if instance exists in constructor.

**Q86: What is the observer pattern?**
Subscribe/Publish model. An object (subject) maintains a list of dependents (observers) and notifies them of state changes.

**Q87: What is the factory pattern?**
A function that creates and returns objects, without requiring the consumer to specify the exact class of object to be created.

### Build Tools & Environment

**Q88: What is transpilation (e.g., Babel)?**
Converting source code from one language syntax (e.g., modern ES6+) to another (e.g., ES5) for backward compatibility.

**Q89: What is tree shaking?**
Dead code elimination during the build process. Bundlers remove unused exports from the final bundle to reduce size.

**Q90: How does bundling work with Webpack?**
It takes modules with dependencies and generates static assets representing those modules. It builds a dependency graph and merges files.

### Testing

**Q91: What are unit tests in JavaScript?**
Testing individual units of source code (functions, classes) in isolation to determine if they are fit for use.

**Q92: How do you mock dependencies?**
Replacing real modules/functions with simulated versions (mocks/spies) to isolate the code being tested. Jest `jest.mock()`.

**Q93: What are common testing libraries in JS?**
*   **Runners/Assertion**: Jest, Mocha, Jasmine.
*   **E2E**: Cypress, Playwright.

### Type Systems

**Q94: What is TypeScript and how does it extend JavaScript?**
A superset of JavaScript that adds static typing. It compiles down to plain JavaScript.

**Q95: What is the difference between static and dynamic typing?**
*   **Static** (TS): Types checked at compile time.
*   **Dynamic** (JS): Types checked at runtime.

### Misc/Tricky

**Q96: What is the output of `[] + []` or `{} + []`?**
*   `[] + []` -> `""` (Empty string).
*   `{} + []` -> `"[object Object]"` (if parsed as expression) or `0` (if block + array coercion in some REPLs). (Standard JS context: string concatenation).

**Q97: How does `this` behave in arrow functions?**
Arrow functions do not have their own `this`. They inherit `this` from the surrounding lexical scope (where they were defined).

**Q98: What are tagged template literals?**
Advanced form of template literals allowing you to parse template strings with a function.
```javascript
tag`Hello ${name}`;
```

**Q99: How do you implement your own `bind` function?**
```javascript
Function.prototype.myBind = function(context, ...args) {
  const fn = this;
  return function(...newArgs) {
    return fn.apply(context, [...args, ...newArgs]);
  };
};
```

**Q100: Explain `Promise.all`, `Promise.race`, and `Promise.any`.**
*   `all`: Resolves when **all** promises resolve (or one rejects).
*   `race`: Resolves/Rejects as soon as **first** promise settles.
*   `any`: Resolves as soon as **any** promise resolves (rejects only if all reject).

---

## 🔹 4. Core Concepts & Syntax (Questions 101-130)

**Q101: What is the Temporal Dead Zone (TDZ)?**
The period between entering the scope of a variable declared with `let` or `const` and the actual line where it is initialized. Accessing the variable in this zone throws a `ReferenceError`.

**Q102: What are labeled statements?**
A statement with an identifier label that can be referred to by `break` or `continue` statements to control the flow of nested loops.
```javascript
loop1: for(let i=0; i<3; i++) {
  break loop1;
}
```

**Q103: What’s the difference between `Object.freeze()` and `Object.seal()`?**
*   `freeze()`: Object is immutable. Cannot add, delete, or modify properties.
*   `seal()`: Cannot add or delete properties, but *can* modify existing property values.

**Q104: How does `Object.defineProperty()` work?**
It defines a new property or modifies an existing one on an object, allowing fine-grained control over descriptors like `configurable`, `enumerable`, `writable`, and `value`.

**Q105: How to make a property read-only in JavaScript?**
Using `Object.defineProperty()` with `writable: false`.
```javascript
Object.defineProperty(obj, 'prop', { value: 42, writable: false });
```

**Q106: What is the role of the `with` statement (and why is it discouraged)?**
Extends the scope chain for a statement. Discouraged because it causes performance issues and confusing scope ambiguity (blocked in strict mode).

**Q107: How does JavaScript handle integer vs float precision?**
JS has only one number type: IEEE 754 64-bit floating point. Integers are safe up to `2^53 - 1` (`Number.MAX_SAFE_INTEGER`). Floats can have precision errors (e.g., `0.1 + 0.2 !== 0.3`).

**Q108: What is the difference between `delete` and setting `undefined`?**
*   `delete obj.key`: Removes the property entirely from the object.
*   `obj.key = undefined`: The property remains but holds the value `undefined`.

**Q109: What is the purpose of `use strict`?**
Enforces stricter parsing and error handling. Prevents accidental globals, disallows `with`, and makes assignment to non-writable properties throw errors.

**Q110: How do symbols work in JavaScript?**
unique and immutable primitive values, often used as object keys to avoid name collisions and create "hidden" properties.

**Q111: What is `Symbol.iterator`?**
A well-known symbol that specifies the default iterator for an object. Used by `for...of`.

**Q112: What are well-known symbols in JavaScript?**
Predefined symbols like `Symbol.iterator`, `Symbol.toStringTag`, `Symbol.toPrimitive` that customize object behavior.

**Q113: What is the `void` operator?**
Evaluates an expression and returns `undefined`. Often used as `void(0)` to prevent side effects in hrefs or arrow functions.

**Q114: What is the `in` operator used for?**
Checks if a property exists in an object or its prototype chain. `prop in object`.

**Q115: How does `instanceof` really work internally?**
It checks if the `prototype` property of a constructor appears anywhere in the prototype chain of an object.

**Q116: How is the `new` keyword implemented?**
1. Creates a new empty object.
2. Links it to the constructor's prototype.
3. Binds `this` to the new object and calls the constructor.
4. Returns the object (unless constructor returns a non-primitive).

**Q117: What is a “polyfill” in JavaScript?**
Code that implements a feature on web browsers that do not natively support it.

**Q118: What is feature detection in JavaScript?**
Checking if a browser supports a specific feature (API/Method) before using it.
```javascript
if ('geolocation' in navigator) { /* usage */ }
```

**Q119: What is a transpiler and how is it different from a compiler?**
A transpiler (like Babel) translates source code to another source code (ES6 -> ES5). A compiler usually translates to machine/byte code.

**Q120: What is the difference between `escape()` and `encodeURIComponent()`?**
*   `escape()`: Deprecated.
*   `encodeURIComponent()`: Standard. Encodes special characters (including `?`, `&`, `/`) for use in URL components.

### Execution Context & Scope (Questions 121-150)

**Q121: What are the phases of execution in JS?**
1.  **Creation Phase**: Window/Global object created, `this` defined, variables declared (hoisting).
2.  **Execution Phase**: Code executed line by line, values assigned.

**Q122: How does the call stack work in JavaScript?**
LIFO (Last In, First Out) structure. Tracks function calls. When a function receives a call, it's pushed; when it returns, it's popped.

**Q123: What is the difference between `globalThis`, `window`, and `self`?**
*   `window`: Global object in Browser.
*   `self`: Refers to global scope in Workers (and Browser).
*   `global`: Global object in Node.js.
*   `globalThis`: Standardized universal reference to the global object across environments.

**Q124: How do block scopes behave inside loops?**
Using `let` in a `for` loop creates a new lexical scope for each iteration, allowing closures to capture the correct value of `i`.

**Q125: What is tail call optimization (TCO)?**
An ES6 optimization where a recursive call in the tail position (last action) reuse the stack frame, preventing stack overflow. (Rarely supported).

**Q126: What is the purpose of the `arguments` object?**
An array-like object accessible inside functions (non-arrow) containing the values of the arguments passed to that function.

**Q127: How does `eval()` affect scope?**
It executes code in the local scope where it's called, potentially modifying local variables (unless indirect eval). Slow and insecure.

**Q128: How does `Function()` constructor work?**
Creates a new Function object. Parsed in the global scope (doesn't access local closures).
```javascript
const sum = new Function('a', 'b', 'return a + b');
```

**Q129: What is the difference between `apply()`, `call()`, and `bind()`?**
*   `call(this, arg1, arg2)`: Invokes immediately with args.
*   `apply(this, [args])`: Invokes immediately with array of args.
*   `bind(this, args)`: Returns a new function with bound `this`, invoked later.

**Q130: What is the scope of variables declared with `var` in `for` loops?**
Function (or Global) scoped. The variable is shared across all iterations. Often leads to bugs with async callbacks inside loops without closures/IIFE.

**Q131: What is function composition?**
Combining two or more functions to produce a new function. `f(g(x))`.
```javascript
const compose = (f, g) => x => f(g(x));
```

**Q132: How to implement a debounce function manually?**
```javascript
function debounce(fn, delay) {
  let timeout;
  return function(...args) {
    clearTimeout(timeout);
    timeout = setTimeout(() => fn.apply(this, args), delay);
  };
}
```

**Q133: How to implement throttling manually?**
```javascript
function throttle(fn, limit) {
  let inThrottle;
  return function(...args) {
    if (!inThrottle) {
      fn.apply(this, args);
      inThrottle = true;
      setTimeout(() => inThrottle = false, limit);
    }
  };
}
```

**Q134: What is an arity of a function?**
The number of arguments a function expects (`func.length`).

**Q135: What is a higher-order function?**
A function that takes another function as an argument or returns a function (e.g., `map`, `filter`, `bind`).

**Q136: What is function chaining?**
Pattern where methods return `this` (the object itself), allowing multiple calls to be chained in a single statement. `obj.method1().method2()`.

**Q137: What is the difference between `.call(this)` and `.apply(this)`?**
(Duplicate of Q129). `call` takes arguments separately, `apply` takes them as an array.

**Q138: How to memoize a recursive function?**
Pass the cache as an argument or wrap the function in a memoizer.
```javascript
const memo = {};
function fib(n) {
  if (n in memo) return memo[n];
  if (n <= 1) return n;
  return memo[n] = fib(n-1) + fib(n-2);
}
```

**Q139: How can you implement `once()` behavior in JS?**
```javascript
function once(fn) {
  let called = false;
  return function(...args) {
    if (!called) {
      called = true;
      return fn.apply(this, args);
    }
  };
}
```

**Q140: What are function decorators in JavaScript?**
(Stage 3 Proposal). Functions that wrap other classes/methods to extend usage. `@decorator`. Useful for logging, binding, etc.

### OOP & Prototypes (Questions 141-150)

**Q141: How to implement multiple inheritance in JavaScript?**
JS doesn't support multiple inheritance directly. Use **Mixins** (copying methods from multiple objects to a target prototype) or Class Composition.

**Q142: What’s the difference between `super()` and `super.prop()`?**
*   `super()`: Calls the parent class constructor. Must be called before `this`.
*   `super.method()`: Calls a method on the parent class prototype.

**Q143: How does prototype pollution happen?**
When attacker can modify the prototype of a base object (like `Object.prototype`), affecting all objects in the app. Often occurs during unsafe recursive merges.

**Q144: What is `constructor.name`?**
A read-only property of the function that reference the name of the function (class).
```javascript
class Foo {}
console.log(new Foo().constructor.name); // "Foo"
```

**Q145: How to dynamically add methods to a prototype?**
Assign a function to the `prototype` object.
```javascript
Array.prototype.last = function() { return this[this.length - 1]; };
```

**Q146: What’s the difference between static and instance methods?**
*   **Static**: Called on the class itself (`Class.method()`). Cannot access `this` instance.
*   **Instance**: Called on instances of the class (`new Class().method()`).

**Q147: What is mixin pattern in JavaScript?**
A pattern to add properties/methods from one object to another without using inheritance.
```javascript
Object.assign(User.prototype, sayHiMixin);
```

**Q148: What is object augmentation?**
Adding new properties or methods to an existing object (or prototype) after it has been created.

**Q149: How is classical inheritance different from prototypal?**
*   **Classical**: Class is a blueprint. Objects are instances. Classes inherit from classes.
*   **Prototypal**: Objects inherit directly from other objects.

**Q150: What is the purpose of `Object.getPrototypeOf()`?**
Standard method to get the prototype (`[[Prototype]]`) of an object. Safer than `__proto__`.

### ES2020+ Features (Questions 151-160)

**Q151: What is the nullish coalescing operator (`??`)?**
Returns the right-hand side operand when the left-hand side is `null` or `undefined` (not just falsy).
```javascript
const val = input ?? "default";
```

**Q152: What is optional chaining and how does it work?**
(Duplicate of Q45). Access deep properties safely. `user?.profile?.name`.

**Q153: What are private class fields in JavaScript?**
Fields prefixed with `#` are private to the class and cannot be accessed from outside.
```javascript
class Counter { #count = 0; }
```

**Q154: What is top-level await?**
Allows using `await` outside of async functions in modules. The module waits for the promise to resolve before being imported by others.

**Q155: What are logical assignment operators (`&&=`, `||=`, `??=`)?**
Combine logical operators with assignment.
`x ||= y` is `x = x || y` (assigns only if x is falsy).

**Q156: What is `BigInt` in JavaScript?**
A primitive for integers larger than `2^53 - 1`. Created by appending `n` to a number. `10n`.

**Q157: What is `WeakRef`?**
Creates a weak reference to an object, allowing it to be garbage collected. Uses `ref.deref()` to access.

**Q158: What is `FinalizationRegistry`?**
Allows you to request a callback when an object is garbage collected.

**Q159: How does `Promise.allSettled()` work?**
Waits for all promises to settle (resolve OR reject). Returns an array of objects describing the outcome of each.

**Q160: How does dynamic import (`import()`) work?**
Loads a module asynchronously and returns a Promise. Useful for code splitting.
```javascript
import('./module.js').then(module => ...);
```

### DOM & BOM Specific (Questions 161-170)

**Q161: What is the difference between `document.body` and `document.documentElement`?**
*   `document.body`: Refers to the `<body>` element.
*   `document.documentElement`: Refers to the `<html>` element (root).

**Q162: How do mutation observers work?**
An API to watch for changes to the DOM tree (attributes, childList, subtree).
```javascript
const observer = new MutationObserver(cb);
observer.observe(target, { attributes: true });
```

**Q163: What is event delegation vs event bubbling?**
*   **Bubbling**: The mechanism (events go up).
*   **Delegation**: The pattern using bubbling to handle events on a parent.

**Q164: How do you implement custom events?**
Using `CustomEvent` constructor.
```javascript
const event = new CustomEvent('myEvent', { detail: { foo: 'bar' } });
element.dispatchEvent(event);
```

**Q165: What is the difference between `focus()` and `blur()`?**
*   `focus()`: Sets focus on an element (e.g., input).
*   `blur()`: Removes focus from an element.

**Q166: How do you programmatically scroll an element into view?**
```javascript
element.scrollIntoView({ behavior: 'smooth' });
```

**Q167: What are data attributes in HTML/JS?**
Custom attributes starting with `data-`. Accessed via `dataset`.
`<div data-id="123"></div>` -> `div.dataset.id`.

**Q168: What are the risks of using `innerHTML`?**
High risk of XSS (Cross-Site Scripting) if user input is mistakenly included without sanitization.

**Q169: How do `getBoundingClientRect()` and layout work?**
Returns the size of an element and its position relative to the viewport. Forces browser to calculate layout (reflow), so can be expensive.

**Q170: How do you listen for form field changes with JavaScript?**
Use the `input` event (fires on every keystroke/change) or `change` event (fires on blur/commit).

### Performance Optimization (Questions 171-180)

**Q171: What is the difference between lazy loading and preloading?**
*   **Lazy Loading**: Delay loading resources until they are needed (e.g., images in viewport).
*   **Preloading**: Load resources early because they will be needed soon (`<link rel="preload">`).

**Q172: What is critical rendering path?**
The sequence of steps the browser goes through to convert HTML, CSS, and JS into pixels on the screen. Optimizing this improves First Contentful Paint.

**Q173: What is layout thrashing?**
Occurs when JS repeatedly reads and writes to the DOM in a loop, causing the browser to recalculate layout (reflow) multiple times unnecessarily.

**Q174: What is the `requestIdleCallback()` API?**
queues a function to be called during a browser's idle periods (when it's not doing high-priority work). Good for analytics/background tasks.

**Q175: How do you debounce a resize event?**
Wrap the resize handler in a debounce function so it only fires after the user stops resizing for N ms.

**Q176: What is the cost of deep object cloning?**
High CPU and Memory usage. Use carefully. `structuredClone` is faster than `JSON` methods but still O(n).

**Q177: What’s the difference between async vs defer in script tags?**
*   `async`: Downloads in parallel, executes as soon as downloaded (blocks parsing). Order not guaranteed.
*   `defer`: Downloads in parallel, executes strictly after HTML parsing is complete. Order preserved.

**Q178: How does lazy evaluation work?**
Delays the evaluation of an expression until its value is needed. JS Generators are a form of lazy evaluation.

**Q179: What is the fastest way to iterate over large arrays?**
Generally, a standard `for` loop (`for(let i=0; i<len; i++)`) is faster than `forEach` or `map` due to less function overhead.

**Q180: How to avoid memory leaks in long-running JS apps?**
Clean up event listeners, clear intervals/timeouts, nullify references to DOM elements when removed, use WeakMaps for caching.

### Testing, Tooling & Build (Questions 181-190)

**Q181: What’s the difference between unit, integration, and e2e testing?**
*   **Unit**: Tests isolated functions/components.
*   **Integration**: Tests interaction between modules.
*   **E2E**: Tests the full application flow from user perspective.

**Q182: What is mocking in JavaScript testing?**
Simulating the behavior of real objects/modules (like API calls) to isolate the code under test.

**Q183: How does code coverage work?**
Measures the percentage of code (lines, functions, branches) executed during tests.

**Q184: How do you test async code?**
Return the Promise, use `async/await` in test case, or use `done` callback (legacy).
```javascript
test('async', async () => {
  const data = await fetchData();
  expect(data).toBe('foo');
});
```

**Q185: What is the purpose of `tsconfig.json` in JS projects?**
Configures the TypeScript compiler (or VS Code JS checking). defines options like target ES version, module system, and strictness.

**Q186: What is a bundler vs a transpiler?**
*   **Transpiler** (Babel): Translates source code to source code (Types -> JS, New JS -> Old JS).
*   **Bundler** (Webpack): Combines multiple files/assets into a single (or few) output files.

**Q187: How does tree shaking work?**
(Duplicate of Q89). Removes unused exports using static analysis of ES modules.

**Q188: What is hot module replacement (HMR)?**
A feature in bundlers that updates modules in the running application (in browser) without a full page reload.

**Q189: What is linting and how does ESLint help?**
Static analysis to find problems in code (syntax errors, style violations, potential bugs) before execution.

**Q190: What is the difference between npm and yarn?**
Both are package managers. npm is default. Yarn was created for speed/determinism (lockfile), but npm has caught up.

### Browser, Events, Storage, Network (Questions 191-200)

**Q191: How does Service Worker work?**
A script running in the background, separate from the web page. Intercepts network requests (proxy), enabling offline support and caching.

**Q192: How do you cache API data in the browser?**
Use `localStorage`, `Cache API` (via Service Worker), or `IndexedDB`.

**Q193: What is the difference between cookies, localStorage, and sessionStorage?**
*   **Cookies**: Sent with every HTTP request, small (4KB), strict expiration.
*   **Storage**: Client-side only, larger (5MB+).
    *   **Local**: Persistent.
    *   **Session**: Tab-scoped.

**Q194: What is a Broadcast Channel API?**
Allows communication between different browsing contexts (tabs, windows, iframes) of the same origin.
```javascript
const bc = new BroadcastChannel('test_channel');
bc.postMessage('hello');
```

**Q195: What are CORS preflight requests?**
Browser sends an `OPTIONS` request before the actual request (if non-simple) to check if the server permits the action.

**Q196: What is the Fetch KeepAlive option?**
`fetch(url, { keepalive: true })`. Allows the request to outlive the page. Useful for analytics on unload.

**Q197: What is the purpose of `navigator` object?**
Exposes state and identity of the user agent (browser). Contains `userAgent`, `language`, `geolocation`, `clipboard`, etc.

**Q198: How do you detect if the user is online or offline?**
`navigator.onLine` property. Listen for `online` and `offline` events on `window`.

**Q199: What is the difference between long polling and WebSockets?**
*   **Long Polling**: Client requests, server holds until data valid, returns, client requests again. (HTTP overhead).
*   **WebSockets**: Persistent, full-duplex TCP connection. Low latency.

**Q200: What is the `Performance` API?**
Provides access to performance-related information. `performance.now()`, `performance.mark()`, `performance.measure()`.

---

## 🔹 5. Language Internals & Deep Dive (Questions 351-380)

**Q351: How is the `this` context determined in setTimeout/setInterval?**
In non-strict mode, it defaults to the global object (`window`). In strict mode, it's `undefined`. Arrow functions inherit `this` from the enclosing scope.

**Q352: How does automatic semicolon insertion (ASI) work?**
The JS engine inserts semicolons at the end of lines if it encounters a syntax error, a line break, or a closing brace. Can cause bugs (e.g., `return` on its own line).

**Q353: What is tail call optimization, and why is it not widely supported?**
Optimization where the last function call in a stack frame replaces the current frame. Not widely supported due to difficulty in debugging (stack traces get lost).

**Q354: How do you implement your own version of `new` keyword?**
1. Create new object inheriting from constructor's prototype.
2. Call constructor with `this` bound to new object.
3. Return object (or constructor's return value).

**Q355: What is the difference between shallow copy and deep copy?**
*   **Shallow**: Copies properties sharing references (nested objects point to same memory).
*   **Deep**: Recursively copies all nested objects (fully independent).

**Q356: How do object property descriptors work?**
Metadata describing a property: `value`, `writable`, `enumerable`, `configurable`. Access via `Object.getOwnPropertyDescriptor()`.

**Q357: What is an accessor property vs data property?**
*   **Data**: Has `value` and `writable`.
*   **Accessor**: Has `get` and `set` functions (no `value`).

**Q358: What does it mean that JavaScript is a single-threaded language?**
It has one call stack and one memory heap. It executes code one line at a time. Concurrency is handled via the Event Loop.

**Q359: How does JavaScript manage memory under the hood?**
Allocates memory (heap/stack) when objects are created and frees it via Garbage Collection (Reachability/Mark-and-Sweep) when not used.

**Q360: How does the V8 engine optimize JavaScript?**
Uses JIT (Just-In-Time) compilation. Components:
*   **Ignition**: Interpreter (generates bytecode).
*   **TurboFan**: Optimizing compiler (generates machine code based on profiling).

### Closures & Scope Tricks (Questions 381-410)

**Q361: How do closures help in data privacy?**
Variables defined in the outer function are inaccessible from the global scope but accessible to the inner function (closure), effectively making them private.

**Q362: Can closures lead to memory leaks? How?**
Yes. If a closure keeps a reference to a large object (e.g., DOM node) that is no longer needed, GC cannot free it because it's still "reachable".

**Q363: How can you use closures to implement private variables?**
```javascript
function Counter() {
  let count = 0;
  return {
    increment: () => ++count,
    get: () => count
  };
}
```

**Q364: What will be the output of a closure inside a loop with `var`?**
It prints the *last* value of the index for all iterations (because `var` is function-scoped). Fix with `let` or IIFE.

**Q365: How does JavaScript handle block-level closures with `let`?**
For each iteration of a loop, JS creates a new lexical environment for the `let` variable, so the closure captures the correct per-iteration value.

**Q366: What is a closure trap and how to fix it?**
Refers to common mistakes like the "var in loop" issue. Fix using block scope (`let`) or passing data into a factory function.

**Q367: Can closures capture updated variable values?**
Yes. Closures capture the *reference* to the variable, not the value at the time of creation. If the variable changes, the closure sees the new value.

**Q368: How can you simulate a module using closures?**
Use an IIFE that returns an object exposing public methods, while keeping variables inside the IIFE private (Module Pattern).

**Q369: Explain scope chaining with a nested closure.**
Inner scope -> Outer Scope -> ... -> Global Scope. JS looks up variables in this chain. Deep nesting keeps the entire chain alive.

**Q370: How does garbage collection affect closed-over variables?**
As long as the closure is alive (referenced), the variables it closes over are considered roots/reachable and won't be collected.

### ES2021–ES2024+ Concepts (Questions 411-440)

**Q371: What are logical assignment operators and when to use them?**
`a ||= b` (assign if falsy), `a &&= b` (assign if truthy), `a ??= b` (assign if nullish). Concise state updates.

**Q372: What are class static initialization blocks?**
A block inside a class (`static { ... }`) to perform complex static field initialization.

**Q373: What are top-level await pitfalls?**
Can block module loading. If module A awaits, module B importing A is also delayed. Circular deps with await can deadlock.

**Q374: How does `at()` method improve indexing?**
Allows negative indexing. `arr.at(-1)` gets the last element. (Cleaner than `arr[arr.length-1]`).

**Q375: What are the benefits of `Array.prototype.groupBy()`?**
(Replaced by `Object.groupBy`). Groups array items into an object based on a callback key.
```javascript
Object.groupBy(items, item => item.category);
```

**Q376: What is a pipeline operator (`|>`) and its status?**
(Stage 2). Syntactic sugar to feed the result of one function into another. `x |> f |> g` === `g(f(x))`.

**Q377: How does the `do` expression proposal work?**
Allows using blocks (if/else, loops) as expressions that return a value.
```javascript
let x = do { if(cond) 1; else 2; };
```

**Q378: What are Records and Tuples (Stage 2 proposal)?**
Immutable data structures. Records `#{ x: 1 }` (like objects) and Tuples `#[1, 2]` (like arrays). Compared by value, not reference.

**Q379: What is the shadow realm proposal in JavaScript?**
API to create a separate global environment (realm) to evaluate code without polluting the main global scope.

**Q380: What’s the difference between `import.meta` and `import()`?**
*   `import.meta`: Object with metadata about the *current* module (e.g., `url`).
*   `import()`: Function to load *other* modules dynamically.

### Functions, Context, and Binding (Questions 441-470)

**Q381: How does function hoisting differ from variable hoisting?**
Functions Declarations are hoisted *with* their definition (usable before line). `var` variables are hoisted as `undefined`.

**Q382: Why are arrow functions not suitable as constructors?**
They lack the internal `[[Construct]]` method and strict `prototype` property. Calling them with `new` throws an error.

**Q383: Can you override `this` in arrow functions?**
No. `bind`, `call`, and `apply` do not change the `this` value of an arrow function (it remains lexically bound).

**Q384: How do fat-arrow functions differ in scope handling?**
They capture `this`, `arguments`, `super`, and `new.target` from the surrounding lexical context.

**Q385: Can you write a function that returns itself?**
Yes. `function f() { return f; }`.

**Q386: How to implement a function spy/logger?**
Wrap the target function.
```javascript
function spy(fn) {
  return function(...args) {
    console.log('Called with', args);
    return fn.apply(this, args);
  }
}
```

**Q387: How to make a function self-memoizing?**
Attach a cache property to the function object itself.
```javascript
function f(x) {
  if (!f.cache) f.cache = {};
  if (x in f.cache) return f.cache[x];
  // compute...
}
```

**Q388: What’s a trampoline function in recursion?**
A utility to flatten recursion into a loop. Functions return "thunks" (next step) instead of calling themselves, preventing stack overflow.

**Q389: What are generator functions and how do they work internally?**
Functions that return an iterator (`Generator`). They maintain internal state and yield execution control back to the caller using `yield`.

**Q390: What is the difference between generator and async generator?**
*   **Generator** (`function*`): Yields values synchronously. Returns `{ value, done }`.
*   **Async Generator** (`async function*`): Yields values asynchronously. Returns a Promise resolving to `{ value, done }`.

### Advanced Async & Concurrency (Questions 471-500)

**Q391: How does `Promise.prototype.finally()` work?**
Execute a callback when promise settles (fulfilled or rejected). Returns a new promise resolving with original value (transparent).

**Q392: How do you cancel a fetch request?**
Use `AbortController`.
```javascript
const controller = new AbortController();
fetch(url, { signal: controller.signal });
controller.abort();
```

**Q393: What is the AbortController API?**
A general-purpose API to abort asynchronous tasks (DOM requests, Fetch, Streams).

**Q394: What is a scheduler in JavaScript?**
(Proposal `scheduler.postTask`). API to prioritize and schedule tasks (user-blocking, background) in the event loop.

**Q395: How does cooperative multitasking work in JS?**
Breaking long tasks into smaller chunks (using `setTimeout` or `yield`) to allow other events/rendering to interleave (yielding to main thread).

**Q396: How does concurrency differ from parallelism in JS?**
*   **Concurrency**: Handling multiple tasks in overlapping time periods (Task A starts, pauses, Task B runs). JS Main Thread.
*   **Parallelism**: Running multiple tasks simultaneously. Web Workers.

**Q397: What are microtasks vs macrotasks with examples?**
*   **Micro**: `Promise.then`, `MutationObserver`, `queueMicrotask`. (Ran immediately after script).
*   **Macro**: `setTimeout`, `setInterval`, I/O. (Ran next loop tick).

**Q398: How does event delegation help async UIs?**
Allows handling events for elements that haven't loaded yet (e.g., list items fetched from API) by placing one listener on the container.

**Q399: How do you chain multiple fetch requests properly?**
Return the next Fetch Promise inside the `.then()`.
```javascript
fetch(A).then(r => r.json()).then(data => fetch(B));
```

**Q400: What happens when you `await` a non-promise?**
JS wraps it in a resolved Promise (`Promise.resolve(value)`) and waits for the microtask queue to process it (suspends function execution).

### Regex, Parsing, String Ops (Questions 401-410)

**Q401: How do you escape special characters in a RegExp?**
Prepend with backslash `\`. No built-in function, but can use replace: `str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')`.

**Q402: What is lazy vs greedy matching in regex?**
*   **Greedy** (default): Matches as much as possible (`.*`).
*   **Lazy** (`?`): Matches as little as possible (`.*?`).

**Q403: How to create a named capture group?**
Use `(?<name>pattern)`.
```javascript
const match = /(?<year>\d{4})/.exec("2024");
console.log(match.groups.year);
```

**Q404: What are the use cases for `String.prototype.matchAll()`?**
Iterating over *all* matches of a regex (with capturing groups) using a loop. Returns an iterator.

**Q405: What is the difference between `match()` and `exec()`?**
*   `str.match(regex)`: Returns array of matches (or null). If global `/g`, doesn't capture groups details.
*   `regex.exec(str)`: Returns one match details with index/groups. Stateful (advances `lastIndex`).

**Q406: How do you validate a strong password with regex?**
Lookaheads. `/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&]).{8,}/`

**Q407: What is the difference between `search()` and `indexOf()`?**
*   `search(regex)`: Accepts Regex. Returns index.
*   `indexOf(str)`: Accepts String. Faster for simple substrings.

**Q408: How does `padStart()` help with formatting?**
Pads the start of the string with another string until it reaches input length. ` "5".padStart(2, "0") // "05"`.

**Q409: What’s the difference between `slice()`, `substring()` and `substr()`?**
*   `slice(start, end)`: Handles negative indices. (Preferred).
*   `substring(start, end)`: Swaps arguments if start > end. No negatives.
*   `substr(start, len)`: Deprecated. Second arg is length.

**Q410: How do tagged template literals work?**
(Duplicate of Q98). Function parses the template parts and expression values.

### Data Handling, Encoding, JSON (Questions 411-420)

**Q411: How to safely stringify an object with circular references?**
Use `JSON.stringify` with a custom replacer or a library like `flatted`. Native stringify throws TypeError.

**Q412: How does `structuredClone()` differ from `JSON.stringify()`?**
It handles cyclic references, `Date`, `Map`, `Set`, `Blob`, etc., which JSON stringify causes data loss (Date->String, Map->{}) or errors.

**Q413: What is base64 encoding and how to do it in JS?**
Binary-to-text encoding.
*   Node: `Buffer.from(str).toString('base64')`.
*   Browser: `btoa(str)` (encode) / `atob(encoded)`.

**Q414: How do you parse and validate nested JSON?**
`JSON.parse()` handles nested structures automatically. Wrap in `try...catch` to validate syntax.

**Q415: How to handle large JSON file parsing efficiently?**
Use streaming parsers (like Oboe.js or Clarinet) instead of loading the huge string ensuring memory isn't exhausted.

**Q416: What’s the difference between JSONP and CORS?**
*   **JSONP**: Hack using `<script>` tags to bypass Same-Origin Policy (GET only, insecure).
*   **CORS**: Standard HTTP mechanism (headers) to allow cross-origin requests safely.

**Q417: What is the `reviver` function in `JSON.parse()`?**
A second argument function `(key, value) => ...` that transforms the results. Useful for restoring `Date` objects from strings.

**Q418: What’s the difference between parsing a JSON string vs object?**
Wait... parsing an object isn't a thing. You stringify an object to JSON string, and parse a JSON string to an object.

**Q419: How do you serialize data with Dates or BigInt?**
*   `Date`: `toJSON()` converts to ISO string.
*   `BigInt`: `JSON.stringify` fails. Use a `.toJSON` method or replacer to convert to String.

**Q420: How do you send FormData with fetch?**
Pass the `FormData` object as body. Browser automatically sets `Content-Type: multipart/form-data`.
```javascript
fetch(url, { method: 'POST', body: formData });
```

### Memory, Storage, Workers (Questions 421-430)

**Q421: How does the JS event loop interact with Web Workers?**
Workers have their *own* separate Event Loop and Memory Heap. They communicate with the main thread via `postMessage`.

**Q422: What is SharedArrayBuffer and Atomics?**
*   `SharedArrayBuffer`: Memory shared between Main Thread and Workers.
*   `Atomics`: Operations to read/write safely (prevent race conditions) on shared numbers.

**Q423: What are transferable objects in Web Workers?**
Objects (ArrayBuffer, MessagePort) whose ownership is *transferred* (moved) to the worker. Zero-copy (fast), but original context loses access.

**Q424: What is the IndexedDB API and how does it compare to localStorage?**
Async, transactional, excessive object store. Can store large amounts of structured data (files/blobs), unlike localStorage (string-only, 5MB limit).

**Q425: How do you handle blob URLs and object URLs?**
`URL.createObjectURL(blob)`. Reference to a file in memory. Note: Must revoke (`URL.revokeObjectURL`) to free memory.

**Q426: What’s the lifetime of data in sessionStorage?**
Until the tab/window is closed. Refreshes generally keep the session.

**Q427: How to store and retrieve objects in localStorage?**
Must serialize/deserialize.
`setItem('key', JSON.stringify(obj))`
`JSON.parse(getItem('key'))`

**Q428: What are Service Workers and how do they cache?**
(Duplicate of Q191). They use the `Cache Storage API` to programmatically store Request/Response pairs.

**Q429: What is the Cache API?**
Storage mechanism for Request/Response object pairs. Available in Window and Workers. `caches.open('name').then(c => c.put(req, res))`.

**Q430: How do you monitor memory usage in JS?**
`performance.memory` (Chrome specific). Or DevTools Memory Profiler (Heapsnapshot). `process.memoryUsage()` in Node.

### Security & Edge Cases (Questions 431-440)

**Q431: What is prototype pollution and how to prevent it?**
(Duplicate of Q143). Prevent by using `Object.create(null)`, freezing prototypes, or using safe merge libraries.

**Q432: What is script injection and how to sanitize input?**
Injecting malicious JS. Sanitize HTML (DOMPurify), use CSP, encode entities.

**Q433: What is DOM-based XSS?**
Vulnerability where the attack payload is executed as a result of modifying the DOM "environment" in the client via JS (e.g., `location.hash` -> `innerHTML`).

**Q434: What is CSP (Content Security Policy)?**
HTTP Header `Content-Security-Policy`. Restricts sources (domains) from which content (scripts, images) can be loaded.

**Q435: How to secure client-side tokens?**
Don't store in localStorage (XSS vulnerable). Store in `httpOnly` `SameSite` cookies.

**Q436: What is clickjacking and how can JS prevent it?**
Attacker overlays an invisible iframe. Prevent with `X-Frame-Options` header. (JS frame-busting scripts are unreliable).

**Q437: How does `innerHTML` open you to security risks?**
It evaluates scripts inside `<img onerror=...>` tags (but blocks direct `<script>` tags usually).

**Q438: What are best practices for safely handling user input?**
Treat all input as untrusted. Validate type/format. Output encode specific to context (HTML, JS, CSS).

**Q439: What is the difference between `eval()` and `Function()` constructor in terms of security?**
Both are dangerous. `Function()` is slightly "safer" regarding scope (runs in global), but still allows arbitrary code execution.

**Q440: Why is it unsafe to use user input directly in template literals?**
If used to generate HTML/SQL without sanitization, it leads to XSS/Injection.

### Math, Date, Utility (Questions 441-450)

**Q441: How do you generate a UUID in JavaScript?**
`crypto.randomUUID()` (Modern). Or usage of `Math.random` (Legacy/insecure).

**Q442: What is the difference between `Math.floor()` and `Math.trunc()`?**
*   `Math.floor()`: Rounds down (-5.5 -> -6).
*   `Math.trunc()`: Removes decimal part (-5.5 -> -5). Same for positive numbers.

**Q443: How to implement a random number in a range?**
`Math.floor(Math.random() * (max - min + 1)) + min`.

**Q444: How to find time difference between two ISO date strings?**
`new Date(d2) - new Date(d1)` gives milliseconds.

**Q445: What is the best way to clone a date object?**
`new Date(date.getTime())` or simply `new Date(date)`.

**Q446: What is `Intl.DateTimeFormat` used for?**
Language-sensitive date and time formatting.
`new Intl.DateTimeFormat('en-US').format(date)`.

**Q447: How to sort an array of objects by date?**
`arr.sort((a, b) => new Date(a.date) - new Date(b.date))`.

**Q448: How to throttle a function using `requestAnimationFrame()`?**
For visual updates, wrap the update in `rAF`. Ensures it only runs once per frame.

**Q449: How does `setTimeout(..., 0)` really behave?**
Pushes callback to the Macrotask queue to run as soon as stack is empty. Min delay is typically 4ms (clamped).

**Q450: How to pause execution for 5 seconds in an async function?**
`await new Promise(r => setTimeout(r, 5000));`

---

## 🔹 6. Language Quirks & Internals (Questions 651-680)

**Q651: What is the difference between `Object.is()` and `===`?**
*   `===`: `-0 === +0` is true, `NaN === NaN` is false.
*   `Object.is()`: `-0` is not `+0`, `NaN` is `NaN`. More precise for edge cases.

**Q652: What happens if you override `toString()` in an object?**
It changes how the object is converted to a primitive string (e.g., in alerts or string concatenation).

**Q653: How does `valueOf()` affect type coercion?**
It is called when an object is converted to a primitive value (number or string). If returns a primitive, that value is used.

**Q654: What is boxing and unboxing in JavaScript?**
*   **Boxing**: Wrapping a primitive (`"str"`) in an Object (`new String("str")`) to access methods. (JS does this auto).
*   **Unboxing**: Extracting the primitive value from the object wrapper (`obj.valueOf()`).

**Q655: What is an arguments object shadowing named parameters?**
If a non-strict function has a parameter named `arguments`, it shadows the built-in `arguments` object, hiding it.

**Q656: Can an object have duplicate keys?**
No. In strict mode/modern JS literals, duplicate keys overwrite the previous one (last one wins). `JSON.parse` behavior varies by implementation but usually last wins.

**Q657: How are floating-point numbers represented internally in JS?**
IEEE 754 Double Precision (64-bit). Sign bit, 11-bit Exponent, 52-bit Mantissa (Fraction).

**Q658: How does `toPrecision()` differ from `toFixed()`?**
*   `toFixed(n)`: Formats to `n` digits *after* decimal point.
*   `toPrecision(n)`: Formats to `n` significant digits (total digits).

**Q659: Why does `0.1 + 0.2 !== 0.3` in JavaScript?**
Due to floating-point precision errors in binary representation. `0.1 + 0.2` is `0.30000000000000004`.

**Q660: How do you make a number always display two decimal places?**
`num.toFixed(2)`. Returns a string string.

### Advanced Array & Object Behavior (Questions 681-710)

**Q661: What is the output of `typeof []` and `typeof {}`?**
Both return `"object"`.

**Q662: How does `delete arr[index]` affect array length?**
It does *not* affect length. It leaves a "hole" (empty slot) at that index. `arr[index]` becomes `undefined`.

**Q663: How does sparse array differ from a regular array?**
Sparse arrays have gaps (empty slots). Iterators like `forEach` skip empty slots, but standard `for` loops don't.

**Q664: What is the difference between `Object.entries()` and `Object.values()`?**
*   `entries()`: Returns `[[key, value], [key, value]]`.
*   `values()`: Returns `[value, value]`.

**Q665: What’s the output of spreading a string into an array?**
`[..."hello"]` -> `['h', 'e', 'l', 'l', 'o']`. Spread iterates over the string characters.

**Q666: What is the behavior of `Array.from()` with a string?**
Same as spread. `Array.from("foo")` -> `['f', 'o', 'o']`.

**Q667: How does array destructuring skip elements?**
Using commas. `const [first, , third] = arr;`.

**Q668: What’s the result of `new Array(5).map(x => 1)`?**
`[empty x 5]`. `map` skips empty slots. You get an array of 5 empty slots, not 1s. Use `Array(5).fill(1)` or `Array.from({length:5}, () => 1)`.

**Q669: How do object keys get stringified in maps?**
In `Object`, keys are converted to strings (`obj[{}]` -> `obj['[object Object]']`). In `Map`, keys keep their type (reference).

**Q670: Can you use objects as keys in regular JS objects?**
No. They are coerced to `"[object Object]"`. Only Strings and Symbols can be real keys in Objects.

### Real-World Async Problems (Questions 711-740)

**Q671: How to implement a retry mechanism for failed async requests?**
Recursive function or loop that catches error, checks attempt count, waits, and calls itself again.

**Q672: How to limit concurrent fetch calls?**
Use a semaphore or a library like `p-limit`. Process queue: start N requests, when one finishes, start next.

**Q673: What’s the difference between polling and long-polling?**
*   **Polling**: Periodic request every X seconds.
*   **Long-Polling**: Request stays open until server has data (or timeout), then client immediately requests again.

**Q674: How to time out a Promise after N milliseconds?**
`Promise.race([ originalPromise, new Promise((_, r) => setTimeout(r, N, 'Timeout')) ])`.

**Q675: How to implement a queue with Promises?**
Chain them. `p = p.then(() => nextTask())`. Or use `async` loop processing an array of tasks.

**Q676: What is a deferred object? Can you simulate one?**
An object exposing `resolve` and `reject` methods externally.
```javascript
let resolve, reject;
const p = new Promise((res, rej) => { resolve = res; reject = rej; });
return { promise: p, resolve, reject };
```

**Q677: How to run async functions sequentially?**
`for...of` loop with `await`. `reduce` with promise chain.
```javascript
for (const task of tasks) await task();
```

**Q678: What is the difference between async iterator and regular iterator?**
*   **Iterator**: `next()` returns `{ val, done }`.
*   **Async Iterator**: `next()` returns `Promise({ val, done })`. Used in `for await...of`.

**Q679: What happens when you `await` inside a loop?**
It pauses the loop execution at that line until promise resolves. Makes loop sequential (slower than `Promise.all` parallel).

**Q680: How does the async stack trace differ from a sync one?**
Async traces often lose the caller context (where the async task was initiated) unless the engine supports "Zero-cost async stack traces" (modern V8).

### Custom Implementations (Questions 741-770)

**Q681: How would you implement a custom `forEach()`?**
```javascript
Array.prototype.myEach = function(cb) {
  for (let i = 0; i < this.length; i++) cb(this[i], i, this);
};
```

**Q682: Can you write your own version of `Array.prototype.map()`?**
```javascript
Array.prototype.myMap = function(cb) {
  const res = [];
  for (let i = 0; i < this.length; i++) res.push(cb(this[i], i, this));
  return res;
};
```

**Q683: Implement a deep equality checker (`deepEqual()`).**
Recursively check keys and values. Handle Primitives, Date, RegExp.
```javascript
function isDeepEqual(o1, o2) {
  if(o1 === o2) return true;
  if(typeof o1 !== 'object' || typeof o2 !== 'object') return false;
  const keys1 = Object.keys(o1), keys2 = Object.keys(o2);
  if(keys1.length !== keys2.length) return false;
  return keys1.every(key => isDeepEqual(o1[key], o2[key]));
}
```

**Q684: How to write a function like `JSON.stringify()` (simple case)?**
Recursive function checking type. If object, map keys to `"key": value`. If string, wrap in quotes.

**Q685: Write a polyfill for `Promise.all`.**
Return new Promise. Track count of resolved. If one rejects, reject main. If all resolve, resolve main with array.

**Q686: How would you implement `throttle()` without setTimeout?**
Using `Date.now()`. Check `now - lastRun >= limit`.

**Q687: Implement a deep clone function.**
(See Q65/Q683 ideas). Recursive copy of properties.

**Q688: Implement your own `bind()` method.**
(Duplicate of Q99).

**Q689: Implement a tiny reactive store in JavaScript.**
Object with state and listeners array. `setState` updates state and calls all listeners.

**Q690: Implement an event emitter in JavaScript.**
Map of eventNames to array of callbacks. `on(name, cb)`, `emit(name, data)`.

### Real-World Design Patterns (Questions 771-800)

**Q691: What is the Revealing Module Pattern?**
Module pattern where you define functions privately and return an object literal referencing only the ones you want to be public.

**Q692: How to implement a publish/subscribe pattern in JS?**
(Same as Event Emitter). Decouples sender and receiver.

**Q693: How is the Proxy pattern used in modern JS frameworks?**
Vue 3 / MobX. Wraps data objects to intercept get/set operations for dependency tracking and auto-updates.

**Q694: What is the Facade pattern in frontend dev?**
Providing a simplified API to a complex underlying subsystem (e.g., a simple API wrapper around a complex `fetch` setup).

**Q695: How does dependency injection work in JavaScript?**
Passing dependencies (services, config) into a function/class rather than importing/creating them inside.

**Q696: Explain how the Factory pattern works with JS classes.**
A method/function that returns instances of different classes based on input logic.

**Q697: How do decorators work in JS (TC39 proposal)?**
Functions that modify class/method behavior at definition time (wrapping them).

**Q698: What’s the difference between mixin and inheritance?**
*   **Inheritance**: "Is-a" relationship key (rigid hierarchy).
*   **Mixin**: "Has-a" / "Can-do" capability (flexible composition).

**Q699: How is the Command pattern used in undo/redo systems?**
Encapsulate actions as objects (`{ execute(), undo() }`). Store them in a stack.

**Q700: What is the Strategy pattern and how can JS implement it?**
Swapping algorithms at runtime. pass a function (strategy) as an argument.
`sort(array, algorithmFn)`.

---

## 🔹 7. Browser-Related (Questions 701-750)

**Q701: How does the browser parse HTML+JS during page load?**
1. HTML parser builds DOM.
2. If `<script>` found, pauses parsing, fetches, executes (unless async/defer).
3. If CSS link found, fetches (blocks render).
4. Resume parsing after JS.

**Q702: What are critical rendering path optimizations?**
Minimize critical resources, inline critical CSS, defer non-critical JS, compress images, use CDN.

**Q703: What’s the difference between DOMContentLoaded and load?**
*   **DOMContentLoaded**: HTML loaded/parsed. Scripts executed. (Images/Stylesheets may still be loading).
*   **load**: Everything loaded (including images, styles).

**Q704: How does `window.history.pushState()` work?**
Adds an entry to the browser's session history stack without reloading the page. Used in SPAs for routing.

**Q705: What is `location.hash` used for?**
The anchor part of the URL (`#section`). Does not trigger reload. Used for in-page navigation or client-side routing.

**Q706: What are intersection observers?**
An API to asynchronously observe changes in the intersection of a target element with an ancestor or viewport. (Lazy loading, infinite scroll).

**Q707: How does the `ResizeObserver` API work?**
Notifies when an element's dimensions change. More performant than listening to `window.resize`.

**Q708: What is the difference between passive and non-passive event listeners?**
`{ passive: true }` tells the browser that the listener will *not* call `preventDefault()`. Allows browser to optimize scrolling performance.

**Q709: How do you detect when a tab/window becomes hidden or inactive?**
`document.visibilityState` ('visible', 'hidden') and `visibilitychange` event.

**Q710: What are performance bottlenecks with heavy DOM manipulation?**
Frequent Reflows (Layout calculation) and Repaints. Fix by batching updates (DocumentFragment) or using Virtual DOM.

### Framework-Agnostic Questions (Questions 711-730)

**Q711: What are single-spa micro frontends and how does JavaScript handle routing between them?**
Framework for orchestrating multiple MFEs. A root config handles routing and mounts/unmounts "applications" (JS bundles) based on URL.

**Q712: How to create your own templating engine in JavaScript?**
String replacement (`.replace(/{{prop}}/g, ...)`), or `Function` constructor with string interpolation for logic.

**Q713: What is hydration in frontend frameworks?**
The process of attaching event listeners and state to server-rendered HTML to make it interactive on the client.

**Q714: How do Virtual DOMs optimize JS performance?**
They update a lightweight JS object tree first, compare (diff) with the previous tree, and only apply changed parts to the real DOM (patching).

**Q715: How can you diff two DOM trees?**
Recursive algorithm comparing node types, attributes, and children. (React uses a heuristic O(n) approach using keys).

**Q716: What is the shadow DOM and how is it used in Web Components?**
Encapsulated DOM tree attached to an element. Styles/Scripts inside don't leak out, and global styles don't leak in.

**Q717: How does two-way binding differ from one-way binding?**
*   **One-way**: Model -> View. (React).
*   **Two-way**: Model <-> View. UI updates Model, Model updates UI automatically (Angular/Vue).

**Q718: How is dirty checking implemented?**
Storing old values and periodically comparing them with new values. If different, trigger update. (AngularJS $digest).

**Q719: What are reactive primitives in state libraries?**
Signals, Observables, or Ref/Reactive objects. They track dependencies (who used this value?) and notify them on change.

**Q720: How does lazy hydration improve performance?**
Delays hydration of components until they are visible or interacted with, reducing initial JS execution time (TTI).

### Weird/Tricky Output Puzzles (Questions 861-890 - labeled 721-750 in source list)

**Q721: What is the output of `[] == ![]`?**
`true`. `![]` is `false`. `[] == false` converts `[]` to `""`, then `""` to `0`. `false` to `0`. `0 == 0`.

**Q722: Why is `NaN !== NaN` true?**
By IEEE 754 standard, NaN is not equal to any value, including itself. Use `isNaN` or `Number.isNaN`.

**Q723: What’s the output of `typeof null`?**
`"object"`. (Historical bug in JS).

**Q724: Explain the result of `1 < 2 < 3` and `3 > 2 > 1`.**
*   `1 < 2 < 3`: `true < 3` -> `1 < 3` -> `true`.
*   `3 > 2 > 1`: `true > 1` -> `1 > 1` -> `false`.

**Q725: What is the output of `true + false`?**
`1`. (`1 + 0`).

**Q726: Why does `[].toString()` return an empty string?**
Arrays stringify by joining elements with commas. Empty array has no elements -> `""`.

**Q727: Why does `{}` + `[]` return `"[object Object]"`?**
(Wait, usually `0` in block context in some REPLs, but string concat in expressions). `[object Object]` + `""`.

**Q728: What’s the result of `0 == '0'`, `false == '0'`, `false == undefined`?**
*   `0 == '0'`: `true`.
*   `false == '0'`: `true` (`0 == 0`).
*   `false == undefined`: `false`.

**Q729: What happens when you do `new Boolean(false)`?**
It creates a Boolean *object* wrapper. The object itself is truthy in `if` conditions. `if (new Boolean(false)) { logs }`.

**Q730: What’s the result of `[...'hello']`?**
`['h', 'e', 'l', 'l', 'o']`.

### Tooling, Standards, Ecosystem (Questions 891-920 - labeled 731-750 in source list)

**Q731: What’s the role of `.babelrc` or `babel.config.js`?**
Configuration file for Babel. Defines presets (env, react) and plugins for transpilation.

**Q732: What is the difference between `babel-polyfill` and `babel-runtime`?**
*   **Polyfill**: Modifies global prototype (pollutes global).
*   **Runtime**: Aliases methods (sandboxed). Better for libraries.

**Q733: What is a monorepo in JavaScript?**
Single repository containing multiple projects (packages). Managed by tools like Turborepo, Nx, Lerna, Yarn Workspaces.

**Q734: How does tree-shaking work in Rollup/Webpack?**
Relies on ES6 `import/export` static structure. Marks unused exports and excludes them from bundle.

**Q735: What is the purpose of source maps?**
Maps minified/transpiled code back to original source code. Allows debugging the original code in browser.

**Q736: What is a code-splitting strategy?**
Splitting bundle into smaller chunks.
*   **Vendor**: Libraries (cacheable).
*   **Route-based**: Load per page.
*   **Dynamic**: `import()`.

**Q737: How does a service worker update lifecycle work?**
Install -> Waiting (until old SW stops controlling clients) -> Activate. Use `skipWaiting()` to force update.

**Q738: What’s the difference between yarn.lock and package-lock.json?**
Both lock dependency versions. `yarn.lock` (Yarn), `package-lock.json` (npm). Don't mix them.

**Q739: How do Node.js and browser environments differ in JS APIs?**
Node: `fs`, `http`, `process`, `global`. No DOM, `window`.
Browser: DOM, `window`, `navigator`. No `fs`.

**Q740: What’s the role of `.nvmrc`?**
File containing the Node.js version number (`18.16.0`). Used by `nvm` to switch to the correct version automatically.

**Q741: What’s the difference between deep freeze and shallow freeze?**
*   `Object.freeze()` is shallow (nested objects are mutable).
*   Deep freeze recursively applies freeze to all nested objects.

**Q742: What are rehydration bugs?**
Mismatch between the server-rendered HTML and the client-side VDOM initial render. React throws errors/warnings.

**Q743: How does object spread handle symbol properties?**
It copies own enumerable Symbol properties.

**Q744: How do you debounce input in vanilla JS without third-party libs?**
(See Q132). Reset a `setTimeout` on each call.

**Q745: What is the difference between `Function.prototype.toString()` and `.name`?**
*   `toString()`: Returns the full source code of the function.
*   `.name`: Returns the name of the function.

**Q746: How does V8 optimize tail calls?**
(Wait, Q353 says it's not widely supported). It reuses the stack frame. Safari supports it. V8 dropped it mostly.

**Q747: How do WeakMap keys behave with garbage collection?**
If the key (object) is not referenced anywhere else, it gets GC'd, and the entry is removed from the WeakMap automatically.

**Q748: What is cross-site script inclusion (XSSI)?**
Attack reading data from JSON APIs by including them as `<script>` tags. Prevent with parser-breaking prefixes `)]}',\n`.

**Q749: What is a null prototype object?**
Object with no prototype. `Object.create(null)`. No methods like `toString` or `hasOwnProperty`. safe for maps.

**Q750: How to convert an arguments object to an array?**
`[...arguments]` or `Array.from(arguments)`.

---

## 🔹 8. Symbols, Meta & Reflection (Questions 751-780)

**Q751: What are the use cases of `Symbol.toPrimitive`?**
Customizing how an object is converted to a primitive value (e.g., in math ops).
```javascript
obj[Symbol.toPrimitive] = function(hint) { return hint === 'number' ? 10 : 'obj'; }
```

**Q752: How does `Symbol.hasInstance` customize `instanceof`?**
Allows an object to decide if another object is an instance of it.
```javascript
class Even {
  static [Symbol.hasInstance](obj) { return obj % 2 === 0; }
}
10 instanceof Even // true
```

**Q753: What is `Symbol.isConcatSpreadable`?**
Boolean to configure if an object should be flattened to its array elements when using `Array.prototype.concat()`.

**Q754: What does `Symbol.match` override?**
Used by `String.prototype.match()`. Use it to create custom matchers (like Regex objects).

**Q755: What’s the purpose of `Symbol.iterator` vs `Symbol.asyncIterator`?**
*   `iterator`: For synchronous loops (`for...of`).
*   `asyncIterator`: For asynchronous loops (`for await...of`).

**Q756: How to use `Symbol.toStringTag` in objects?**
Customizes the output of `Object.prototype.toString.call(obj)`.
```javascript
obj[Symbol.toStringTag] = 'MyObj'; // [object MyObj]
```

**Q757: What does `Symbol.unscopables` do?**
Specifies object properties to be excluded from the `with` environment bindings.

**Q758: What’s the purpose of `Reflect.get()`?**
Gets a property value. Like `obj[key]` but as a function. `Reflect.get(target, key, receiver)`.

**Q759: What is the difference between `Reflect.set()` and direct assignment?**
Returns a Boolean indicating success (`true/false`). Assignment throws strict mode errors or fails silently.

**Q760: Why use `Reflect.ownKeys()` over `Object.keys()`?**
`Reflect.ownKeys()` returns *all* keys (enumerable, non-enumerable, string, AND symbol keys). `Object.keys` only enumerable strings.

### Proxy Deep Dive (Questions 781-810 - labeled 761-790 in source list)

**Q761: What are traps in JavaScript Proxy?**
Methods in the handler object that intercept operations (e.g., `get`, `set`, `apply`, `construct`).

**Q762: How does `get` and `set` trap work?**
*   `get(target, prop, receiver)`: Intercepts property access.
*   `set(target, prop, val, receiver)`: Intercepts property assignment.

**Q763: What’s the `has` trap and when is it triggered?**
Intercepts the `in` operator. `prop in proxy`.

**Q764: How can Proxy be used to validate object values?**
In the `set` trap, check if the value is valid (e.g., `typeof value === 'number'`) before applying strictly to target.

**Q765: How to observe property access using Proxy?**
Log inside `get` / `set` traps.

**Q766: How to use Proxy for auto-completion or fallback values?**
In `get` trap, if property doesn't exist (`!Reflect.has(target, prop)`), return a default value or suggested key.

**Q767: How can you use Proxy for access logging?**
Simply `console.log` every operation in the traps.

**Q768: How to prevent new properties from being added with a Proxy?**
`preventExtensions` trap or `set` trap returning `false` for new keys.

**Q769: What is the role of `defineProperty` trap in Proxy?**
Intercepts `Object.defineProperty()`. Can prevent property definition or enforce descriptors.

**Q770: Can you use a Proxy as a wrapper for functions?**
Yes. Use the `apply` trap to intercept function calls and `construct` for `new` calls.

### Memory, GC, and Performance Internals (Questions 811-840 - labeled 771-800 in source list)

**Q771: What is the mark-and-sweep garbage collection strategy?**
(Duplicate of Q74). Roots -> Mark reachable -> Sweep (free) unreachable.

**Q772: What causes memory leaks in closures?**
(Duplicate of Q362). Closures holding references to large scopes.

**Q773: How to identify detached DOM nodes?**
Use Heap Snapshot in DevTools. Look for "Detached" nodes (red). They are removed from DOM but referenced by JS.

**Q774: How do WeakRefs help with caching?**
Hold a reference to an object without preventing GC. If GC runs, the cache entry is cleared automatically.

**Q775: Why are WeakMaps better for DOM-based caches?**
If the DOM node is removed, the WeakMap entry (where key=node) is also removed (eventually), preventing leaks.

**Q776: What are finalizers in JavaScript and how are they used?**
`FinalizationRegistry`. Callback runs after object is GC'd. Used for resource cleanup (e.g., closing Wasm handles).

**Q777: How can you prevent memory bloat in single-page applications?**
Virtualize long lists, lazy load data, clean up event listeners/timers on component unmount.

**Q778: What does Chrome’s performance tab measure in JS?**
Frame rate (FPS), CPU usage (Scripting, Rendering, Painting), Memory Heap.

**Q779: How does JS engine handle inline caching?**
Optimizes property access. V8 remembers the "shape" (hidden class) of objects and the offset of properties to skip lookups.

**Q780: What is hidden class optimization in V8?**
V8 creates internal C++ classes (Shapes) for objects with same keys. Adding keys in different order creates different hidden classes (slower).

### Web Workers & Shared Memory (Questions 841-870 - labeled 781-810 in source list)

**Q781: What is the purpose of a Web Worker?**
Run scripts in background threads, keeping the UI thread (Main) responsive/unblocked.

**Q782: What are the communication methods between workers and main thread?**
`postMessage()` and `onmessage` (Structured Cloning). Or `SharedArrayBuffer` (Shared Memory).

**Q783: Can you pass functions into Web Workers?**
No. Functions cannot be structured-cloned. You must pass logic as strings (eval - bad) or load separate files.

**Q784: How to terminate a Web Worker from the main thread?**
`worker.terminate()`.

**Q785: What is a Blob Worker?**
Creating a worker from a Blob URL (string of code) instead of an external file.
`new Worker(URL.createObjectURL(new Blob([code])))`.

**Q786: What’s the difference between dedicated and shared workers?**
*   **Dedicated**: Linked to one script/tab.
*   **Shared**: Accessible by multiple scripts/tabs (same origin).

**Q787: What is the Transferable interface in JS?**
(Duplicate of Q423). Moving ownership of buffer.

**Q788: What is SharedArrayBuffer and how is it used?**
(Duplicate of Q422). Shared binary memory.

**Q789: What are Atomics and how do they ensure safety?**
(Duplicate of Q422). `Atomics.add`, `Atomics.wait`, `Atomics.notify`.

**Q790: How do you implement a thread-safe counter with Atomics?**
`Atomics.add(typedArray, index, 1)`.

### Edge API Features (Questions 871-900 - labeled 791-820 in source list)

**Q791: What is the Payment Request API?**
Standardized browser API for handling payment UI (credit cards, addresses) natively.

**Q792: What is the Battery Status API and why was it deprecated?**
Allowed reading battery level. Deprecated due to privacy (fingerprinting) risks.

**Q793: How does the Web Share API work?**
Invokes the native OS share dialog. `navigator.share({ title, url })`. Requires user interaction.

**Q794: What is the Permissions API?**
Query status of permissions (geo, mic, camera). `navigator.permissions.query({ name: 'geolocation' })`.

**Q795: How to detect clipboard read/write availability?**
Use Permissions API `clipboard-read`, `clipboard-write`.

**Q796: What is the Page Visibility API?**
(Duplicate of Q709). `document.hidden`.

**Q797: What is the Beacon API used for?**
`navigator.sendBeacon(url, data)`. Sends small data asynchronously (POST) to server. Guaranteed to send even if page unloads (analytics).

**Q798: How does the Wake Lock API prevent screen sleeping?**
`navigator.wakeLock.request('screen')`. Keeps screen on (e.g., for video or presentation apps).

**Q799: How does the Vibration API work?**
`navigator.vibrate(200)`. Vibrate device for 200ms.

**Q800: What is the Idle Detection API?**
Detects if user is idle (no input) or screen locked. Privacy-sensitive.

### UI + Event Loop Integration (Questions 901-930 - labeled 801-830 in source list)

**Q801: How to defer non-critical JS for better Time to Interactive?**
`defer`, dynamic `import()`, `requestIdleCallback`.

**Q802: What is the difference between `requestIdleCallback()` and `requestAnimationFrame()`?**
*   `rAF`: Runs before paint (visual updates). High priority.
*   `rIC`: Runs when idle (background tasks). Low priority.

**Q803: How can JS simulate batching DOM reads and writes?**
(FastDOM library concept). Group all reads (measure), then all writes (mutate) to prevent layout thrashing.

**Q804: What is layout thrashing and how to avoid it?**
(Duplicate of Q173).

**Q805: What’s a mutation observer and how does it affect reactivity?**
(See Q162). It catches DOM changes. Used by frameworks to sync state with DOM if 3rd party scripts modify it.

**Q806: How does JavaScript throttle input rendering on scroll?**
Using `passive` listeners or throttling the scroll handler to `rAF`.

**Q807: What is the difference between synchronous vs async click handling?**
*   **Sync**: `onclick`. Blocks invalidation.
*   **Async**: If logic involves `await`, the event finishes propagation before logic completes.

**Q808: What happens if you trigger layout inside a JS loop?**
Performance disaster. Browser re-calculates layout N times.

**Q809: Why should you avoid frequent style recalculations?**
High CPU cost.

**Q810: How to use ResizeObserver for responsive elements?**
Observe element size changes (container queries concept) instead of viewport changes.

### Obscure JS Constructs & Pitfalls (Questions 931-960 - labeled 811-840 in source list)

**Q811: What does `new.target` represent in a constructor?**
Reference to the constructor that was invoked with `new`. useful to detect if class was subclassed or called without new.

**Q812: What’s the result of `typeof function* () {}`?**
`"function"`.

**Q813: How is a generator paused and resumed internally?**
Use `yield`. The execution context is saved (stack frame suspended) and restored on `next()`.

**Q814: What happens if you yield inside a try block?**
The generator pauses. `try` block remains active. If you throw into the generator (`gen.throw()`), it catches in that `try` block.

**Q815: What is the “temporal dead zone” with `let`?**
(Duplicate of Q101).

**Q816: Can you redeclare `let` or `const` in the same block?**
No. SyntaxError.

**Q817: What is the return value of `void` operator?**
`undefined`.

**Q818: How is `.length` of a function determined?**
Number of named parameters (excluding default parameters and rest).

**Q819: What does `arguments.callee` do and why is it deprecated?**
References the currently executing function. Deprecated because it breaks optimizations and strict mode.

**Q820: How does `Function.length` differ from `arguments.length`?**
*   `Function.length`: Expected parameters (definition).
*   `arguments.length`: Actual passed arguments (invocation).

### Async Patterns and Edge Handling (Questions 961-990 - labeled 821-850 in source list)

**Q821: What is an async IIFE?**
`(async () => { await ... })()`. Allows usages of await in older environments (pre top-level await).

**Q822: What happens if you `await` inside a constructor?**
SyntaxError. Class constructors cannot be async.

**Q823: How to make a function return synchronously or asynchronously based on context?**
Check input. If promise, return promise chain. If value, return value. (Zalgo anti-pattern, usually bad).

**Q824: How do you catch errors thrown from an async generator?**
Wrap the `for await...of` loop in a `try...catch`.

**Q825: What’s the difference between rejecting a promise vs throwing?**
Inside async function/then: same effect (Promise rejects).
Outside: Throwing crashes script (synchronous), Rejecting returns rejected Promise.

**Q826: How does `queueMicrotask()` differ from `setTimeout()`?**
(Duplicate Q46/397). `queueMicrotask` runs ASAP (before render). `setTimeout` runs next cycle (after render/paint usually).

**Q827: How do browser rendering and JS event loop interact?**
Task -> Microtasks -> Render (Style/Layout/Paint) -> Task. (If VSync).

**Q828: Can you pause a generator with `await`?**
Yes. `yield await promise`. It waits for promise, then yields value.

**Q829: How do you cancel an async generator?**
`iterator.return()`. Terminates the generator loop.

**Q830: How do you prioritize microtasks in a browser?**
You can't "prioritize" them manually. They always run before Macrotasks.

### New & Experimental Features (Questions 991-1020 - labeled 831-860 in source list)

**Q831: What is the `Observable` proposal?**
Standard interface for handling streams of data (push-based), like RxJS but native.

**Q832: What are decorators?**
(Duplicate).

**Q833: What are Records & Tuples?**
(Duplicate).

**Q834: What is module attributes proposal?**
(Now "Import Attributes"). `import json from './data.json' with { type: 'json' }`.

**Q835: How will pattern matching improve control flow?**
`match (val) { when (1) -> ... }`. More powerful switch.

**Q836: What is `Array.prototype.with()` and how is it different from `splice()`?**
`arr.with(index, value)`. Returns a *new* array with the item replaced (Non-mutating). `splice` mutates.

**Q837: What is `Map.groupBy()`?**
(See Q375). Group items into Map.

**Q838: What is `Promise.withResolvers()`?**
Returns `{ promise, resolve, reject }`. Avoiding the `new Promise` constructor callback nesting.

**Q839: What is `ArrayBuffer.transfer()` and how is it useful?**
Resizing/moving ArrayBuffers without copying.

**Q840: What is the Temporal API and how does it replace Date?**
New standard date/time API fixing `Date` quirks (immutable, timezone-aware, distinct types for Date/Time/Zoned).

### Bonus: Debugging, Observability, DX (Questions 1021-1050 - labeled 841-850 in source list)

**Q841: What is a source map?**
(Duplicate Q735).

**Q842: How to log deep objects safely without circular ref crashes?**
`console.dir(obj, { depth: null })` (Node). Or stringify with cycle replacer.

**Q843: How to monitor async stack traces across nested promises?**
Enable "Async Stack Traces" in DevTools.

**Q844: What are break-on-access debugging strategies?**
Right click in DevTools -> "Break on... attribute modification / subtree mod" or `debug(fn)`.

**Q845: What is a memory snapshot and how do you analyze it?**
Capture of heap state. Compare snapshots to see what objects were created but not collected (leaks).

**Q846: What does the “Retainers” tab show in Chrome DevTools?**
Shows what object is holding a reference to the selected object, preventing it from GC.

**Q847: What is the Timeline flame chart?**
Visualizes call stack over time. Width = duration. Stack = depth.

**Q848: How can you simulate offline in DevTools?**
Network Tab -> Offline.

**Q849: What does “Long Task” mean in performance profiling?**
Task taking > 50ms. Blocks main thread, causing jank.

**Q850: How can you track event listeners registered on an element?**
DevTools Elements -> Event Listeners tab. Or `getEventListeners(node)` in Console API.

---

## 🔹 9. Language Internals, Quirks & Grammar (Questions 851-880)

**Q851: Why is `null` an object in JavaScript?**
Legacy bug. The first 3 bits of type tags were 000 for objects, and null was all zeros. It won't be fixed because it would break existing web sites.

**Q852: What is the lexical environment in JavaScript?**
Internal structure that holds identifier-variable mapping. Consists of Environment Record (storage) and Reference to Outer Environment (scope chain).

**Q853: What are environment records in execution contexts?**
The actual storage place where variables and function declarations are stored within a Lexical Environment.

**Q854: What is a binding identifier?**
The identifier (name) used to reference a variable/constant in a specific scope/environment.

**Q855: How are function parameters stored in memory?**
Stored as variables in the function's Execution Context (Environment Record).

**Q856: How do arrow functions capture `this` differently than normal functions?**
They don't have their own `this` binding. They resolve `this` from the enclosing lexical scope (like a variable).

**Q857: What is the purpose of `super()` and when must it be called?**
Called in constructor of derived class. Must be called before accessing `this`. Initializes the parent class.

**Q858: Why is `typeof NaN === 'number'`?**
NaN is defined as a numeric value ("Not a Number") in IEEE 754 standard. It's a special invalid state of the Number type.

**Q859: What is a completion record in JavaScript?**
Internal specification type used to explain control flow (return, break, continue, throw). `{ type: normal|return|throw, value, target }`.

**Q860: How does JavaScript handle dangling commas in objects and arrays?**
It ignores them. Allowed in Arrays, Objects, and (ES2017) Function parameters. `[1, 2, ]` length is 2.

### Error Handling, Custom Errors & Try/Catch Deep Dive (Questions 881-910 - labeled 861-890 in source list)

**Q861: What’s the difference between a runtime and syntax error?**
*   **Syntax**: Code cannot be parsed (typos). Happens before execution.
*   **Runtime**: Code parses but fails during execution (e.g., `undefined.prop`).

**Q862: How do you throw a custom error in JavaScript?**
`throw new Error('msg')` or extend Error class. `class MyError extends Error {}`.

**Q863: What is the `Error.captureStackTrace()` method?**
(Node.js/V8 specific). Creates a `.stack` property on the target object. useful for custom error constructors to hide implementation details from stack trace.

**Q864: How does the stack trace behave in async functions?**
Historically broken. Modern engines stitch together stack traces ("Async Stack Frames") so you see the caller even if it awaited.

**Q865: What happens if you throw inside a finally block?**
The original error (if any) is overwritten. The new error bubbles up.

**Q866: What is a rethrow pattern?**
Catching an error, doing something (log), and then throwing it again `throw err` so higher-level handlers can deal with it.

**Q867: Can `finally` override the return value?**
Yes. If `finally` returns a value, it overrides any `return` or `throw` from `try/catch`.

**Q868: How do you handle promise rejections without `catch()`?**
You can't really (it will cause Unhandled Rejection). Listen to `unhandledrejection` event globally.

**Q869: What is `window.onerror` used for?**
Global event handler for uncaught exceptions. Legacy API.

**Q870: How do you define a global error handler in Node.js?**
`process.on('uncaughtException', cb)` and `process.on('unhandledRejection', cb)`.

### Object-Oriented Programming (Questions 911-940 - labeled 871-900 in source list)

**Q871: What is method overriding in JavaScript classes?**
Defining a method in a child class with the same name as the parent class. Child method replaces parent's.

**Q872: How does inheritance work with ES6 classes?**
`class Child extends Parent {}`. Uses prototype chain under the hood.

**Q873: How do static properties differ from instance properties?**
*   **Static**: Belongs to Class. shared.
*   **Instance**: Belongs to Object. unique per instance.

**Q874: What is a class expression in JavaScript?**
Defining a class as an expression to assign to a variable. `const MyClass = class {};`.

**Q875: Can you instantiate a class without `new`?**
No. ES6 Classes invoke a check and throw TypeError: "Class constructor cannot be invoked without 'new'".

**Q876: What’s the difference between `super.method()` and `this.method()`?**
*   `super.method()`: Calls method from Parent prototype.
*   `this.method()`: Calls method from current instance (potentially overridden).

**Q877: How can a class extend a built-in like `Array`?**
`class MyArray extends Array {}`. Works in ES6. `map`, `filter` return instances of `MyArray`.

**Q878: What happens when you `return` an object from a constructor?**
That object overrides the `this` instance being created. `new Foo()` returns the explicit object.

**Q879: What’s the output if you call a class without `new`?**
TypeError.

**Q880: How do private class fields differ from closure-based privacy?**
*   **#private**: Native, cleaner syntax, memory efficient (per instance).
*   **Closures**: Robust, but memory overhead (per method created per instance unless shared).

### Functional Programming Concepts (Questions 941-970 - labeled 881-910 in source list)

**Q881: What is a monad in JavaScript (simple explanation)?**
A wrapper around a value (like Promise or Array) that provides a standard way (`.then`, `.flatMap`) to transform the value and return a new wrapper.

**Q882: What’s the difference between `compose` and `pipe` functions?**
*   **Compose**: Right-to-Left. `f(g(x))` -> `compose(f, g)(x)`.
*   **Pipe**: Left-to-Right. `g(f(x))` -> `pipe(g, f)(x)`.

**Q883: How can you implement a curry function manually?**
Recursively return a function until args count matches arity.
```javascript
const curry = (fn, ...args) => args.length >= fn.length ? fn(...args) : (...next) => curry(fn, ...args, ...next);
```

**Q884: What’s the difference between partial application and currying?**
*   **Curry**: Converts `f(a,b,c)` to `f(a)(b)(c)`.
*   **Partial**: Fixes some arguments producing a new function with smaller arity. `f(1, b, c)`.

**Q885: What is immutability and how do you enforce it in JS?**
State cannot change after creation. Enforce with `Object.freeze()`, `const` (limited), or libraries like Immutable.js/Immer.

**Q886: What are transducers?**
Composable algorithmic transformations (map+filter etc) that process data in a single pass, independent of the input data structure.

**Q887: How does `reduceRight()` differ from `reduce()`?**
Processes array from Right to Left (End to Start).

**Q888: What is a functor in functional programming with JavaScript?**
An object/container that implements a `map` method (e.g., Array).

**Q889: How would you implement `compose()` in JavaScript?**
`funcs.reduce((a, b) => (...args) => a(b(...args)))`.

**Q890: What are point-free functions?**
Functions that don't explicitly mention their arguments. `const newFn = compose(f, g)` instead of `x => f(g(x))`.

### Internationalization (i18n) and Formatting APIs (Questions 971-990 - labeled 891-910 in source list)

**Q891: How does `Intl.NumberFormat` work?**
Formats numbers based on locale (currency, percent, separators).
`new Intl.NumberFormat('de-DE', { style: 'currency', currency: 'EUR' }).format(1234.56)`.

**Q892: What is the difference between `toLocaleString()` and `Intl.NumberFormat()`?**
`toLocaleString` creates a new formatter instance every call (slower in loops). `Intl` allows reusing the formatter instance (faster).

**Q893: How to localize date/time for multiple time zones?**
`new Intl.DateTimeFormat('en-US', { timeZone: 'Asia/Tokyo' })`.

**Q894: What is the purpose of `Intl.PluralRules`?**
helps choose the correct plural form (zero, one, two, few, many, other) for a given number/locale.

**Q895: What are BCP-47 locale strings?**
Standard for language tags. `en-US`, `fr-CA`.

**Q896: How does `Intl.ListFormat` work?**
Formats lists linguistically. `['A', 'B', 'C']` -> "A, B, and C".

**Q897: What is the `Intl.RelativeTimeFormat`?**
"5 minutes ago", "in 2 days".

**Q898: How can you dynamically switch between languages in a JS app?**
Reload page or use a router/context to swap the locale string passed to `Intl` or i18n library.

**Q899: What is `Intl.DisplayNames` and its use case?**
Translates language/region/script codes to human readable names. "US" -> "United States".

**Q900: How does `Intl.Segmenter` help with word tokenization?**
Splits string into words/sentences correctly handling locale rules (e.g., languages without spaces).

### Testing JavaScript (Questions 991-1020 - labeled 901-910 in source list)

**Q901: What are the differences between mocking and stubbing?**
*   **Stub**: Canned answer to a call (force return value).
*   **Mock**: Expectations about the call (verify it was called X times with Y args).

**Q902: What’s the purpose of `jest.fn()`?**
Creates a mock function to spy on calls or implementation.

**Q903: How to test async functions using `done()` vs `async/await`?**
(Duplicate Q184).

**Q904: What is snapshot testing?**
Comparing the rendered output (JSON/HTML) against a saved "snapshot" file to detect unintended changes.

**Q905: How does `sinon.spy()` work?**
Wraps a function to record arguments, return value, `this`, and exception. Original function still runs.

**Q906: How to test DOM events using testing libraries?**
`fireEvent.click(node)` or `userEvent.click(node)`.

**Q907: What are fake timers and when should you use them?**
(Jest `useFakeTimers`). Fast-forward time (setTimeout, setInterval) instantly in tests instead of waiting meant time.

**Q908: What is the difference between shallow and deep rendering?**
*   **Shallow**: Renders component one level deep (mocks children).
*   **Deep**: Renders full tree (real children).

**Q909: How to mock fetch in a unit test?**
`global.fetch = jest.fn(() => Promise.resolve(...))`.

**Q910: What is mutation testing in JavaScript?**
(e.g., Stryker). Tool modifies your code (mutants) and runs tests. If tests pass, the mutant "survived" (bad coverage).

### Security-Oriented (Questions 1021-1050 - labeled 911-920 in source list)

**Q911: How do you protect against clickjacking with JavaScript?**
Frame busting (`top.location != self.location`), but headers (`X-Frame-Options`) are better.

**Q912: What is JavaScript sandboxing?**
Running untrusted code in an isolated environment (iframe, worker, vm module) with limited access to globals/DOM.

**Q913: What is the difference between innerHTML and textContent (security)?**
(Duplicate). `innerHTML` parses HTML (XSS risk). `textContent` is text only.

**Q914: What are DOMPurify or similar libraries used for?**
Sanitizing HTML strings (removing script/bad tags) before inserting into DOM.

**Q915: How can Content Security Policy (CSP) block XSS?**
By banning inline scripts (`script-src 'self'`) and restricting domains.

**Q916: Why should you avoid eval in modern apps?**
(Duplicate). Performance + Security.

**Q917: What is the difference between escaping and encoding?**
*   **Escape**: Making characters safe for a context (HTML entities `<` -> `&lt;`).
*   **Encode**: Converting format (URL Encoding space -> `%20`).

**Q918: How can JavaScript open up vulnerabilities when used with user input?**
If input controls logic (`eval`, `setTimeout`), DB queries (NoSQL injection), or DOM (`innerHTML`).

**Q919: How do CSP nonces work with inline scripts?**
Server generates random nonce, adds to header. matching `<script nonce="...">` allowed.

**Q920: What is subresource integrity (SRI)?**
`<script src="..." integrity="sha384-...">`. Browser verifies hash of fetched file ensures it wasn't tampered with (CDN hack).

---

## 🔹 10. Advanced Modules, Imports & Bundling (Questions 951-960)

**Q951: What is the difference between `import` and `require`?**
*   `require`: CommonJS (Node), Dynamic, Synchronous.
*   `import`: ES Module (Standard), Static structure, Async loading support.

**Q952: How does tree shaking work with ES modules?**
Bundlers analyze `import/export` graph statically. Unused exports are dropped. `require` cannot be tree-shaken easily.

**Q953: What are dynamic imports and how do they improve performance?**
`import()`. Loads code on demand (lazy loading), reducing initial bundle size.

**Q954: Can you `import` conditionally based on runtime values?**
Only with dynamic `import()`. Static `import` must be at top-level.

**Q955: What is the default module loading strategy in browsers?**
Defer. (Fetched in parallel, executed in order after parsing).

**Q956: How do circular imports behave in ES Modules?**
They work (mostly). The module instance is shared. If variable is `let`, accessing it before initialization throws TDZ. `var` or `function` works.

**Q957: What is the difference between `export default` and `named exports`?**
*   **Named**: Importing must match name (`{ foo }`). Multiple per file.
*   **Default**: Import can be named anything (`import anyName`). One per file.

**Q958: How does `import.meta` provide context in ES modules?**
Exposes `import.meta.url`, `import.meta.resolve()`. Context about the module file itself.

**Q959: What is module hoisting in bundlers like Webpack?**
(Scope Hoisting). Concatenating modules into a single closure to reduce function overhead and improve execution speed.

**Q960: What are the advantages of ES modules over CommonJS?**
Standard (Browser + Node), Async, Tree-shakable, Cyclic dependency support is better, Syntax.

### Animation, Timing, and Visuals (Questions 961-970)

**Q961: How does `requestAnimationFrame()` optimize rendering?**
Syncs with monitor refresh rate (VSync).

**Q962: What is the difference between `setInterval()` and `requestAnimationFrame()`?**
`setInterval` runs regardless of rendering (can cause jank/dropped frames). `rAF` runs only when browser is ready to paint.

**Q963: How can you create a smooth progress bar with JS?**
Use `rAF` to interpolate width value smoothly over time, or CSS transitions triggered by JS class change.

**Q964: How does JavaScript throttling improve animation performance?**
Limits event frequency (like scroll) to prevent overloading the main thread.

**Q965: What are scroll-driven animations and how can you implement them?**
Animations linked to scroll position. Use `IntersectionObserver` or new CSS Scroll-Linked Animations API (polyfill with JS).

**Q966: How to synchronize CSS animations with JS logic?**
Listen to `animationend` / `transitionend` events.

**Q967: How to pause/resume an animation using JS?**
`element.style.animationPlayState = 'paused' | 'running'`.

**Q968: How do you animate with `transform` vs `top/left` properties?**
`transform` uses GPU (Compositor thread), does not trigger reflow. `top/left` triggers reflow (CPU), slower.

**Q969: What is the FLIP animation technique in JavaScript?**
**F**irst, **L**ast, **I**nvert, **P**lay. Pre-calculate the change in position and apply a transform to simulate movement cheaply.

**Q970: How does `IntersectionObserver` help with lazy animations?**
Only start/play animation when element is visible in viewport.

### Mobile, Accessibility, Device APIs (Questions 971-980)

**Q971: How do you detect device orientation changes in JavaScript?**
`window.addEventListener('orientationchange')` or matching media query `(orientation: portrait)`.

**Q972: What is the DeviceMotionEvent and when is it useful?**
Access accelerometer/gyroscope data. (Games, compass).

**Q973: How to detect touchscreen devices using JS?**
`'ontouchstart' in window` or `navigator.maxTouchPoints > 0`.

**Q974: How do you improve accessibility with ARIA attributes?**
Add `aria-label`, `aria-hidden`, `role` to help screen readers understand custom widgets.

**Q975: What is the difference between `pointerdown`, `mousedown`, and `touchstart`?**
*   **PointerEvents**: Unified. Handles mouse, touch, pen.
*   **TouchEvents**: Touch only.
*   **MouseEvents**: Mouse only (or emulated touch).

**Q976: How to detect screen reader usage with JS (and why it's hard)?**
Explicitly hard/impossible for privacy reasons. focus tracking is one indirect clue.

**Q977: What is responsive font scaling using `window.devicePixelRatio`?**
Adjusting canvas resolution or loading retina images based on pixel density.

**Q978: How do you enable keyboard-only navigation detection?**
Listen for Tab key presses. Add `.focus-visible` class styles.

**Q979: How do you support dark mode toggling in JS?**
`matchMedia('(prefers-color-scheme: dark)')` to detect OS setting. Toggle class on `<body>`.

**Q980: How do you detect and respond to viewport resize on mobile?**
`resize` event. Be careful of virtual keyboard appearing (shrinks viewport). use `visualViewport` API.

### Real-World Scenarios & Behavioral Patterns (Questions 981-990)

**Q981: How do you detect if a user is idle?**
Listen for mousemove/keydown. If no event for X mins, trigger idle. Reset timer on event.

**Q982: How do you debounce API requests on search input?**
(Duplicate Q50).

**Q983: How to manage state between tabs using `localStorage`?**
Listen to `storage` event. When one tab updates storage, others receive event with new value.

**Q984: How to handle stale cache data in a SPA?**
Versioning. If API version > Cached version, clear cache. Or `Cache-Control` headers.

**Q985: What happens if two tabs write to localStorage at once?**
It is synchronous and blocking (for that origin). No race condition on the individual write, but last write wins.

**Q986: How to persist scroll position across navigation in a SPA?**
Save `window.scrollY` to `sessionStorage` on unload/routeChange. Restore on load. (Or standard `history.scrollRestoration`).

**Q987: How to sync Redux or global state with URL query parameters?**
Listener on state change -> `history.replaceState`. Listener on `popstate` -> generic action to update store.

**Q988: How to manage undo/redo stack in JS?**
(See Q699 Commmand Pattern). Stack of previous states.

**Q989: How to prevent form resubmission on reload?**
Post/Redirect/Get pattern. Or `history.replaceState()` to remove query params/state after submit.

**Q990: How do you implement optimistic UI updates?**
Update UI immediately (optimistic). Send request. If fail, revert UI (rollback).

### Web APIs & Network Strategies (Questions 991-1000)

**Q991: What is HTTP/2 Push and how can JavaScript use it?**
Server sends resources before client asks. JS doesn't "use" it directly; browser cache handles it. (Deprecated in Chrome 106+).

**Q992: What is fetch keepalive used for?**
(Duplicate Q196).

**Q993: What is a streaming response in fetch and when should you use it?**
`response.body.getReader()`. Read streams chunk by chunk. (Video, huge JSON, progress indicators).

**Q994: What is the purpose of `navigator.sendBeacon()`?**
(Duplicate Q797).

**Q995: What is the difference between offline-first and cache-first strategies?**
*   **Offline-first**: App assumes offline. Serves from cache, syncs later.
*   **Cache-first**: Try cache. If missing, fetch network. (Speed focus).

**Q996: How do you update cached assets in a PWA?**
Service Worker versioning. If `sw.js` changes, new one installs.

**Q997: What is stale-while-revalidate in service workers?**
Serve from cache immediately (fast), but update cache from network in background (fresh next time).

**Q998: How do you track network failures and retries in JavaScript?**
`window.addEventListener('online/offline')`. Catch fetch errors.

**Q999: What is background sync and how is it implemented?**
Service Worker `sync` event. Queues requests when offline, sends them when connectivity returns.

**Q1000: How to handle slow internet or flaky connections in JS UIs?**
Show skeletons/spinners. Use timeouts. Retry logic. Optimistic UI.

---
