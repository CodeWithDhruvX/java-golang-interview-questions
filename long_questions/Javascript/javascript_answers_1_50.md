# JavaScript Interview Questions & Answers (1-50)

## 🔹 Beginner (1–30)

### Syntax & Basics

#### Question 1: What are the different data types in JavaScript?

**Answer:**
JavaScript has 8 standard data types, divided into primitives and non-primitives (objects).

**Primitives (Immutable):**
*   **String**: Textual data (`"hello"`).
*   **Number**: Integer or floating-point numbers (`42`, `3.14`).
*   **BigInt**: Integers larger than `2^53 - 1` (`9007199254740991n`).
*   **Boolean**: Logical entities (`true`, `false`).
*   **Undefined**: Variable declared but not assigned (`undefined`).
*   **Null**: Intentional absence of any object value (`null`).
*   **Symbol**: Unique and immutable identifier (`Symbol('id')`).

**Non-Primitive (Mutable):**
*   **Object**: Collection of key-value pairs (includes Arrays, Functions, Dates, etc.).

**Example:**
```javascript
let str = "Hello";       // String
let num = 100;           // Number
let big = 10n;           // BigInt
let bool = true;         // Boolean
let undef;               // Undefined
let empty = null;        // Null
let sym = Symbol("id");  // Symbol

let obj = { x: 1 };      // Object
let arr = [1, 2, 3];     // Array (Object)
```

---

#### Question 2: What is the difference between `var`, `let`, and `const`?

**Answer:**
They differ in scoping, hoisting, and re-assignment capabilities.

| Feature | `var` | `let` | `const` |
| :--- | :--- | :--- | :--- |
| **Scope** | Function Scope | Block Scope | Block Scope |
| **Reassignable** | Yes | Yes | No |
| **Redeclarable** | Yes | No | No |
| **Hoisting** | Yes (initialized `undefined`) | Yes (TDZ - ReferenceError) | Yes (TDZ - ReferenceError) |

**Example:**
```javascript
function demo() {
    if (true) {
        var a = 1;
        let b = 2;
        const c = 3;
    }
    console.log(a); // 1 (Function scoped)
    // console.log(b); // ReferenceError (Block scoped)
    // console.log(c); // ReferenceError (Block scoped)
}
```

---

#### Question 3: What is hoisting in JavaScript?

**Answer:**
Hoisting is the JavaScript engine's behavior of moving variable and function declarations to the top of their containing scope during the compilation phase.

*   **Function Declarations:** Fully hoisted. Can call before definition.
*   **`var` variables:** Hoisted and initialized with `undefined`.
*   **`let` and `const`:** Hoisted but stay in "Temporal Dead Zone" (TDZ) until the line of initialization is executed. Accessing them before throws `ReferenceError`.

**Example:**
```javascript
// Function Hoisting
sayHello(); // Works: "Hello"
function sayHello() { console.log("Hello"); }

// var Hoisting
console.log(myVar); // undefined
var myVar = 10;

// let Hoisting (TDZ)
// console.log(myLet); // ReferenceError: Cannot access 'myLet' before initialization
let myLet = 20;
```

---

#### Question 4: How does type coercion work in JavaScript?

**Answer:**
Type coercion is the automatic or implicit conversion of values from one data type to another (e.g., string to number).

**Implicit Coercion:**
Happens automatically with operators.
```javascript
console.log("5" + 2);  // "52" (String concatenation)
console.log("5" - 2);  // 3 (String to Number)
console.log("5" * "2"); // 10 (Both to Number)
console.log(true + 1); // 2 (true becomes 1)
```

**Explicit Coercion:**
Manually converting types.
```javascript
let num = Number("123");
let str = String(123);
let bool = Boolean(1);
```

---

#### Question 5: What are truthy and falsy values?

**Answer:**
In boolean contexts (like `if` statements), values are coerced to `true` (truthy) or `false` (falsy).

**Falsy Values (Total 7):**
*   `false`
*   `0`
*   `-0`
*   `0n` (BigInt zero)
*   `""` (Empty string)
*   `null`
*   `undefined`
*   `NaN`

**Truthy Values:**
Everything else is truthy, including:
*   `"0"` (String containing zero)
*   `"false"` (String)
*   `[]` (Empty array)
*   `{}` (Empty object)
*   `function(){}`

**Example:**
```javascript
if ([]) console.log("Arrays are truthy"); // Prints
if ("") console.log("Empty string is truthy"); // Does ONLY not print
```

---

#### Question 6: What is the difference between `==` and `===`?

**Answer:**
*   **`==` (Loose Equality):** Performs type coercion before comparison. If types differ, it attempts to convert them to a common type.
*   **`===` (Strict Equality):** Checks both value and type. No coercion is allowed. Always preferred.

**Example:**
```javascript
console.log(5 == "5");   // true (String "5" converted to number 5)
console.log(5 === "5");  // false (Number != String)

console.log(null == undefined); // true
console.log(null === undefined); // false
```

---

#### Question 7: Explain the difference between `null` and `undefined`.

**Answer:**
*   **`undefined`:** Represents the state of a variable that has been declared but not assigned a value. It is the default return value of functions without a return statement.
*   **`null`:** An assignment value representing "no value" or "empty object". It must be explicitly assigned.

**Code:**
```javascript
let a;
console.log(a); // undefined

let b = null;
console.log(b); // null

// Type check quirk
console.log(typeof undefined); // "undefined"
console.log(typeof null);      // "object" (Legacy bug in JS)
```

---

#### Question 8: How does `typeof` work?

**Answer:**
The `typeof` operator returns a string indicating the type of the operand.

**Rules:**
```javascript
typeof "text"      // "string"
typeof 42          // "number"
typeof true        // "boolean"
typeof undefined   // "undefined"
typeof Symbol()    // "symbol"
typeof 10n         // "bigint"
typeof function(){} // "function"

// Special cases
typeof null        // "object"
typeof []          // "object"
typeof {}          // "object"
typeof NaN         // "number"
```

---

#### Question 9: What are template literals in JavaScript?

**Answer:**
Template literals are string literals enclosed by backticks (`` ` ``) that allow:
1.  **String Interpolation:** Embedding expressions using `${}`.
2.  **Multi-line Strings:** Preserves newlines without `\n`.

**Example:**
```javascript
const name = "Alice";
const age = 25;

// Old way
const str1 = "Hello " + name + ", you are " + age + ".";

// Template Literal
const str2 = `Hello ${name},
you are ${age}.`;

console.log(str2);
```

---

#### Question 10: How does the `switch` statement work?

**Answer:**
The `switch` statement evaluates an expression and matches its value against `case` clauses using **strict equality** (`===`).

**Example:**
```javascript
const day = 2;

switch (day) {
    case 1:
        console.log("Monday");
        break;
    case 2:
        console.log("Tuesday");
        break; // Stops execution here
    default:
        console.log("Other day");
}
```
If `break` is omitted, execution "falls through" to the next case regardless of match.

---

### Functions & Scope

#### Question 11: What is the difference between function declaration and function expression?

**Answer:**
*   **Function Declaration:** Defined with `function name() {}`. Hoisted to top of scope.
*   **Function Expression:** Assigned to a variable `const func = function() {}`. Not hoisted (variable hoisting rules apply).

**Example:**
```javascript
// Declaration
console.log(add(2, 3)); // Works
function add(a, b) { return a + b; }

// Expression
// console.log(sub(5, 2)); // Error: sub is not defined (or undefined if var)
const sub = function(a, b) { return a - b; };
```

---

#### Question 12: What is a callback function?

**Answer:**
A callback is a function passed as an argument to another function, to be executed ("called back") at a later time (synchronously or asynchronously).

**Example:**
```javascript
function processData(input, callback) {
    console.log("Processing " + input);
    callback();
}

function done() {
    console.log("Done!");
}

processData("file.txt", done);
```

---

#### Question 13: What are arrow functions?

**Answer:**
Arrow functions (`=>`) provide a concise syntax for writing functions. They have semantic differences:
1.  **No `this` binding:** They capture `this` from the enclosing lexical context.
2.  **No `arguments` object.**
3.  **Implicit return** for single expressions.

**Example:**
```javascript
// Standard
const add = function(a, b) {
    return a + b;
};

// Arrow
const addArrow = (a, b) => a + b;

// Lexical 'this'
function Person() {
    this.age = 0;
    // Arrow function captures 'this' from Person
    setInterval(() => {
        this.age++; 
    }, 1000);
}
```

---

#### Question 14: What is lexical scoping?

**Answer:**
Lexical scoping means that a function's scope is determined by its physical location in the source code. Inner functions have access to variables declared in their outer parent scopes.

**Example:**
```javascript
const outerVar = "I am outer";

function outer() {
    const innerVar = "I am inner";
    
    function inner() {
        // Accesses outerVar from lexical scope
        console.log(outerVar); 
        console.log(innerVar);
    }
    inner();
}
outer();
```

---

#### Question 15: What is the use of closures in JavaScript?

**Answer:**
A closure is created when a function is defined inside another function, allowing the inner function to access the outer function's variables even after the outer function has returned.

**Uses:**
*   Data privacy/Emulating private methods.
*   Function factories.
*   Memoization.

**Example:**
```javascript
function createCounter() {
    let count = 0; // Private variable via closure
    return {
        increment: () => ++count,
        getValue: () => count
    };
}

const counter = createCounter();
console.log(counter.increment()); // 1
console.log(counter.getValue());  // 1
// console.log(counter.count);    // undefined (inaccessible)
```

---

#### Question 16: How do default parameters work?

**Answer:**
Default parameters allow initializing (formal) function parameters with default values if no value or `undefined` is passed.

**Example:**
```javascript
function greet(name = "Guest") {
    return `Hello, ${name}`;
}

console.log(greet("Alice")); // "Hello, Alice"
console.log(greet());        // "Hello, Guest"
console.log(greet(undefined)); // "Hello, Guest"
console.log(greet(null));    // "Hello, null" (null is a value, so default not used)
```

---

### Arrays & Objects

#### Question 17: How to clone an object in JavaScript?

**Answer:**
*   **Shallow Copy:**
    *   Creates a new object but nested objects are still references.
    *   `Object.assign({}, original)`
    *   Spread syntax: `{...original}`
*   **Deep Copy:**
    *   Creates a fully independent copy.
    *   `structuredClone(original)` (Modern standard)
    *   `JSON.parse(JSON.stringify(original))` (Limitations: No dates, functions, undefined)

**Example:**
```javascript
const obj = { a: 1, nested: { b: 2 } };

// Shallow
const shallow = { ...obj };
shallow.nested.b = 99; // Affects obj!

// Deep
const deep = structuredClone(obj);
deep.nested.b = 100; // Does not affect obj
```

---

#### Question 18: What is destructuring?

**Answer:**
Destructuring allow unpacking values from arrays or properties from objects into distinct variables.

**Example:**
```javascript
// Array Destructuring
const arr = [10, 20];
const [x, y] = arr;

// Object Destructuring
const user = { id: 1, name: "Alice" };
const { name, id } = user;

// Renaming and Defaults
const { name: userName, role = "User" } = user;
```

---

#### Question 19: What are array methods like `map`, `filter`, `reduce`?

**Answer:**
These are Higher-Order Functions for processing arrays.
*   **`map()`**: Transforms every element and returns a new array of the same size.
*   **`filter()`**: Returns a new array with elements that pass the condition.
*   **`reduce()`**: Accumulates array elements into a single value.

**Example:**
```javascript
const nums = [1, 2, 3, 4];

const doubled = nums.map(n => n * 2); // [2, 4, 6, 8]
const evens = nums.filter(n => n % 2 === 0); // [2, 4]
const sum = nums.reduce((acc, curr) => acc + curr, 0); // 10
```

---

#### Question 20: How to remove duplicates from an array?

**Answer:**
The most efficient way in modern JS is using the `Set` object, which stores unique values.

**Example:**
```javascript
const nums = [1, 2, 2, 3, 1, 4];

// Using Set and Spread
const unique = [...new Set(nums)]; 
// [1, 2, 3, 4]

// Using filter (Older method, slower O(N^2))
const uniqueOld = nums.filter((v, i, a) => a.indexOf(v) === i);
```

---

### Control Flow & Async

#### Question 21: How does the event loop work?

**Answer:**
JavaScript is single-threaded. The event loop orchestrates execution:
1.  **Call Stack:** Executes synchronous code.
2.  **Web APIs:** handles async operations (timers, fetch).
3.  **Callback Queue (Task Queue):** Holds callbacks (setTimeout) waiting for Stack to clear.
4.  **Microtask Queue:** Holds Promises. Has higher priority than Task Queue.

**Cycle:**
*   Execute Stack until empty.
*   Check Microtask Queue -> Execute ALL.
*   Check Task Queue -> Execute ONE.
*   Repeat.

---

#### Question 22: What is the difference between synchronous and asynchronous code?

**Answer:**
*   **Synchronous:** Blocking. Code executes line-by-line. The next line waits for the previous one to finish.
*   **Asynchronous:** Non-blocking. Long-running tasks (IO, Timers) start, and code execution continues immediately. Result is handled later (Callback/Promise).

**Example:**
```javascript
// Sync
console.log(1);
console.log(2); // Runs after 1

// Async
console.log(1);
setTimeout(() => console.log(2), 1000); // Runs later
console.log(3); // Runs before 2
```

---

#### Question 23: What are promises?

**Answer:**
A Promise is an object representing the eventual completion (or failure) of an asynchronous operation.

**States:**
*   **Pending:** Initial state.
*   **Fulfilled:** Operation successful (calls `.then()`).
*   **Rejected:** Operation failed (calls `.catch()`).

**Example:**
```javascript
const myPromise = new Promise((resolve, reject) => {
    let success = true;
    if (success) resolve("Done!");
    else reject("Error!");
});

myPromise
    .then(res => console.log(res))
    .catch(err => console.error(err));
```

---

#### Question 24: What is async/await?

**Answer:**
`async/await` is syntactic sugar built on top of Promises to write asynchronous code that looks synchronous.
*   `async`: Ensures function returns a Promise.
*   `await`: Pauses execution until the Promise resolves.

**Example:**
```javascript
async function fetchData() {
    try {
        const response = await fetch('url'); // Pauses here
        const data = await response.json();
        console.log(data);
    } catch (error) {
        console.error(error);
    }
}
```

---

#### Question 25: How does error handling with `try...catch` work?

**Answer:**
It allows handling runtime errors gracefully to prevent the program from crashing.
*   **try:** Block of code to test.
*   **catch:** Block executes if error occurs.
*   **finally:** (Optional) Block executes regardless of result (cleanup).

**Example:**
```javascript
try {
    let result = problematicFunction();
} catch (err) {
    console.log("Error caught:", err.message);
} finally {
    console.log("Cleanup operations");
}
```

---

### DOM Basics

#### Question 26: How to select elements in the DOM?

**Answer:**
Common methods:
*   `document.getElementById('id')`: Best performance for IDs.
*   `document.querySelector('.class')`: First match (CSS selector).
*   `document.querySelectorAll('.class')`: All matches (NodeList).
*   `document.getElementsByClassName('class')`: Live HTMLCollection.

**Example:**
```javascript
const btn = document.querySelector('#submitBtn');
const items = document.querySelectorAll('.list-item');
```

---

#### Question 27: How to attach an event listener?

**Answer:**
Using `addEventListener()`. It attaches an event handler without overwriting existing handlers.

**Example:**
```javascript
const btn = document.querySelector('button');

btn.addEventListener('click', (event) => {
    console.log("Button clicked!");
    console.log(event.target);
});
```

---

#### Question 28: What is the difference between `innerHTML` and `textContent`?

**Answer:**
*   **`innerHTML`**: Gets/Sets HTML markup. Parses string as HTML. Risk of XSS (Cross-Site Scripting).
*   **`textContent`**: Gets/Sets raw text. Does not parse HTML tags. Safer and faster.

**Example:**
```javascript
div.innerHTML = "<b>Bold</b>"; // Renders bold text
div.textContent = "<b>Bold</b>"; // Renders literal string "<b>Bold</b>"
```

---

#### Question 29: What is `event.preventDefault()` and `event.stopPropagation()`?

**Answer:**
*   **`preventDefault()`**: Stops the default browser behavior for the event (e.g., preventing form submission or link navigation).
*   **`stopPropagation()`**: Stops the event from bubbling up the DOM tree to parent elements.

**Example:**
```javascript
form.addEventListener('submit', (e) => {
    e.preventDefault(); // Prevents page reload
});

child.addEventListener('click', (e) => {
    e.stopPropagation(); // Parent click handler won't fire
});
```

---

#### Question 30: How do you manipulate classes using JavaScript?

**Answer:**
Using the `classList` API on an element.

**Methods:**
*   `add('class')`
*   `remove('class')`
*   `toggle('class')`
*   `contains('class')`

**Example:**
```javascript
const box = document.querySelector('.box');
box.classList.add('active');
box.classList.toggle('hidden');
if (box.classList.contains('active')) { /* ... */ }
```

---

## 🔹 Intermediate (31–70)

### Advanced Functions & Closures

#### Question 31: What is a pure function?

**Answer:**
A pure function must satisfy two conditions:
1.  **Deterministic:** Same inputs always produce same output.
2.  **No Side Effects:** Does not modify external state, DOM, or global variables.

**Example:**
```javascript
// Pure
const add = (a, b) => a + b;

// Impure (uses external state)
let count = 0;
const addToCount = (num) => count += num;
```

---

#### Question 32: What is memoization?

**Answer:**
Memoization is an optimization technique to speed up function execution by **caching** results of expensive calls and returning the cached result when the same inputs occur again.

**Example:**
```javascript
function memoize(fn) {
    const cache = {};
    return function(n) {
        if (n in cache) return cache[n];
        const result = fn(n);
        cache[n] = result;
        return result;
    }
}
```

---

#### Question 33: What are IIFEs (Immediately Invoked Function Expressions)?

**Answer:**
Functions that execute as soon as they are defined. Pattern: `(function(){ ... })();`. Commonly used to create a private scope.

**Example:**
```javascript
(function() {
    const privateVar = "Secret";
    console.log("Ran immediately");
})();
// privateVar is not accessible here
```

---

#### Question 34: How do closures preserve data?

**Answer:**
Since inner functions retain references to the scope in which they were created, variables in that scope are not garbage collected as long as the inner function exists. This allows "preserving" state.

**Example:**
```javascript
function timer() {
    let start = Date.now();
    return function() {
        return Date.now() - start; // 'start' is preserved
    }
}
```

---

#### Question 35: What is currying?

**Answer:**
Currying is transforming a function `f(a, b, c)` into sequences of functions `f(a)(b)(c)`.

**Example:**
```javascript
// Normal
const sum = (a, b) => a + b;

// Curried
const curriedSum = (a) => (b) => a + b;

const add5 = curriedSum(5);
console.log(add5(2)); // 7
```

---

### Objects & Prototypes

#### Question 36: What is the prototype chain?

**Answer:**
It is the mechanism for inheritance in JS. When accessing a property of an object, if not found, JS looks at the object's `__proto__`, then that prototype's `__proto__`, forming a chain until null is reached.

**Example:**
```javascript
const parent = { greet: () => "Hello" };
const child = Object.create(parent);
console.log(child.greet()); // "Hello" (found on prototype)
```

---

#### Question 37: What is the difference between `__proto__` and `prototype`?

**Answer:**
*   **`__proto__`** (Instance): The actual object pointer on a specific instance (e.g., `obj.__proto__`). It points to the parent.
*   **`prototype`** (Constructor): The property on a Function (`Func.prototype`) that will be assigned as `__proto__` to instances created with `new Func()`.

---

#### Question 38: What is `Object.create()` used for?

**Answer:**
It creates a new object with a specific object explicitly set as its prototype.

**Example:**
```javascript
const dog = { bark: () => console.log("Woof") };
const pug = Object.create(dog); // pug.__proto__ === dog
pug.bark(); // Works
```

---

#### Question 39: How does inheritance work in JavaScript?

**Answer:**
JavaScript acts via **Prototypal Inheritance**. Objects inherit directly from other objects via links (`prototype`). ES6 classes are syntactical sugar over this system.

**Example (ES5):**
```javascript
function Animal(name) { this.name = name; }
Animal.prototype.speak = function() { return "Noise"; };

function Dog(name) { Animal.call(this, name); }
Dog.prototype = Object.create(Animal.prototype);
```

---

#### Question 40: What are ES6 classes?

**Answer:**
Classes provide a cleaner syntax for creating objects and dealing with inheritance, constructor functions, and prototypes.

**Example:**
```javascript
class User {
    constructor(name) {
        this.name = name;
    }
    sayHi() {
        console.log(`Hi ${this.name}`);
    }
}
```

---

### ES6+ Features

#### Question 41: What are generators in JavaScript?

**Answer:**
Functions that can be paused and resumed using `function*` syntax and the `yield` keyword. They return an Iterator.

**Example:**
```javascript
function* idGen() {
    let id = 1;
    while(true) yield id++;
}
const gen = idGen();
console.log(gen.next().value); // 1
console.log(gen.next().value); // 2
```

---

#### Question 42: What is the spread operator?

**Answer:**
The `...` syntax expands iterables (arrays/strings) into individual elements.

**Usage:**
```javascript
// Arrays
const arr = [1, 2];
const newArr = [...arr, 3, 4]; // [1, 2, 3, 4]

// Objects
const obj = { a: 1 };
const newObj = { ...obj, b: 2 };
```

---

#### Question 43: What is the rest parameter?

**Answer:**
The `...` syntax in function parameters collects multiple arguments into an array.

**Example:**
```javascript
function sum(...numbers) {
    return numbers.reduce((a, b) => a + b, 0);
}
console.log(sum(1, 2, 3)); // 6
```

---

#### Question 44: How do modules (`import`/`export`) work?

**Answer:**
ES Modules (ESM) allow organizing code into files.
*   `export` variables/functions to make them public.
*   `import` to use them in other files.

**Example:**
```javascript
// math.js
export const add = (a, b) => a + b;

// main.js
import { add } from './math.js';
```

---

#### Question 45: What is optional chaining?

**Answer:**
The `?.` operator safely accesses nested object properties. If a reference is `null` or `undefined`, it short-circuits and returns `undefined` instead of throwing an error.

**Example:**
```javascript
const user = { details: null };
// console.log(user.details.age); // Error
console.log(user.details?.age);   // undefined (Safe)
```

---

### Async & Event Loop

#### Question 46: What is the microtask queue vs macrotask queue?

**Answer:**
*   **Microtasks (High Priority):** `Promise.then`, `queueMicrotask`, `MutationObserver`. Processed immediately after current script, before rendering.
*   **Macrotasks (Task Queue):** `setTimeout`, `setInterval`, `setImmediate`. Processed one per loop cycle.

**Order:** Sync Code -> Microtasks -> UI Render -> One Macrotask.

---

#### Question 47: How does `setTimeout` actually work?

**Answer:**
It is not part of the JS engine (V8) but a Web API. It sets a timer; when the timer expires, the callback is pushed to the **Task Queue** (Macrotask queue). It runs only when the Call Stack is empty.

**Note:** `setTimeout(..., 0)` does not run immediately, it defers execution to the end of the stack.

---

#### Question 48: What is a race condition? How do you avoid it?

**Answer:**
Occurs when the outcome depends on the timing of uncontrollable events (e.g., two async requests modifying the same variable).
**Avoid:** Using `Promise.all` for dependent data, ensuring atomic updates, or proper state management.

---

#### Question 49: How does JavaScript handle concurrency?

**Answer:**
Despite being single-threaded, JS handles concurrency using the **Event Loop** and **Asynchronous Callbacks**. It offloads blocking I/O operations (network, disk) to the browser/OS threads (Web APIs) and processes results when they are ready.

---

#### Question 50: What is a debounce vs throttle function?

**Answer:**
Both limit the rate of function execution.
*   **Debounce:** Delays execution until 'n' ms have passed since the last call. (Good for Search bars).
*   **Throttle:** Ensures function executes at most once every 'n' ms. (Good for Scroll events).

**Example (Debounce):**
```javascript
function debounce(fn, delay) {
    let timeout;
    return (...args) => {
        clearTimeout(timeout);
        timeout = setTimeout(() => fn(...args), delay);
    }
}
```
