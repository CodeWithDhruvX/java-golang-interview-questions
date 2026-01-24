# JavaScript Interview Questions & Answers (101-150)

## 🔹 Core Concepts & Syntax (101–130)

### Question 101: What is the Temporal Dead Zone (TDZ) in JavaScript?

**Answer:**
The Temporal Dead Zone (TDZ) is the period between the start of a block and the initialization of a `let` or `const` variable declared within that block. Accessing the variable during this phase throws a `ReferenceError`.

**Example:**
```javascript
{
    // TDZ starts here
    // console.log(x); // ReferenceError: Cannot access 'x' before initialization
    
    let x = 10; // TDZ ends here
    console.log(x); // 10
}
```

---

### Question 102: What are labeled statements?

**Answer:**
A label provides an identifier for a statement (usually a loop) that allows you to `break` or `continue` that specific loop from within nested loops.

**Example:**
```javascript
outerLoop: for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
        if (i === 1 && j === 1) {
            break outerLoop; // Breaks the outer loop, not just the inner one
        }
        console.log(`i=${i}, j=${j}`);
    }
}
```

---

### Question 103: What’s the difference between `Object.freeze()` and `Object.seal()`?

**Answer:**
*   **`Object.freeze(obj)`**: Makes the object **immutable**. Cannot add, delete, or modify properties. (Deep freeze requires recursion).
*   **`Object.seal(obj)`**: Prevents adding or deleting properties, but **can modify** existing properties.

**Example:**
```javascript
const frozen = Object.freeze({ prop: 42 });
frozen.prop = 33; // Fails (succeeds silently or TypeError in strict mode)

const sealed = Object.seal({ prop: 42 });
sealed.prop = 33; // Works
delete sealed.prop; // Fails
```

---

### Question 104: How does `Object.defineProperty()` work?

**Answer:**
It allows defining a new property or modifying an existing one with precise control over its descriptors (`value`, `writable`, `enumerable`, `configurable`).

**Example:**
```javascript
const obj = {};
Object.defineProperty(obj, 'prop', {
    value: 42,
    writable: false, // Cannot be changed
    enumerable: true // Shows in loops
});

obj.prop = 100;
console.log(obj.prop); // 42
```

---

### Question 105: How to make a property read-only in JavaScript?

**Answer:**
Using `Object.defineProperty()` with `writable: false`.

**Example:**
```javascript
const user = { name: "Alice" };
Object.defineProperty(user, "id", {
    value: 123,
    writable: false
});
user.id = 456; // Ignored or Error
```

---

### Question 106: What is the role of the `with` statement (and why is it discouraged)?

**Answer:**
`with (obj) { ... }` extends the scope chain for a statement, making properties of `obj` available as local variables.
**Discouraged because:**
*   It creates ambiguity (variable vs property).
*   Performance (disables optimizations).
*   Forbidden in strict mode (`"use strict"`).

---

### Question 107: How does JavaScript handle integer vs float precision?

**Answer:**
JavaScript does not have distinct `int` and `float` types. All numbers are **IEEE 754 Double Precision Floats**. This leads to precision errors like `0.1 + 0.2 !== 0.3`.
Integers are safe up to `Number.MAX_SAFE_INTEGER` (`2^53 - 1`). Use `BigInt` for larger integers.

---

### Question 108: What is the difference between `delete` and setting `undefined`?

**Answer:**
*   **`delete obj.key`**: Removes the property entirely from the object. The key no longer exists.
*   **`obj.key = undefined`**: The property still exists, but its value is `undefined`.

**Example:**
```javascript
const obj = { a: 1, b: 2 };
delete obj.a; 
console.log('a' in obj); // false

obj.b = undefined;
console.log('b' in obj); // true
```

---

### Question 109: What is the purpose of `use strict`?

**Answer:**
It enforces a stricter parsing and error handling model for JavaScript code.
*   Prevents using undeclared variables (`x = 10`).
*   Disallows `with`.
*   Throws errors for assigning to read-only properties.
*   Fixes `this` in simple calls (undefined instead of global object).

---

### Question 110: How do symbols work in JavaScript?

**Answer:**
`Symbol` is a primitive data type that is unique and immutable. Often used as object keys to avoid name collisions and create "hidden" properties (skipped by `for...in` and `JSON.stringify`).

**Example:**
```javascript
const id = Symbol("id");
const obj = {
    [id]: 12345
};
console.log(obj[id]); // 12345
```

---

### Question 111: What is `Symbol.iterator`?

**Answer:**
A well-known symbol that specifies the default iterator for an object. If an object has this property, it can be iterated using `for...of` loops.

**Example:**
```javascript
const iterable = {
    *[Symbol.iterator]() {
        yield 1;
        yield 2;
    }
};
for (const x of iterable) console.log(x); // 1, 2
```

---

### Question 112: What are well-known symbols in JavaScript?

**Answer:**
Predefined symbols that allow you to customize language behavior.
*   `Symbol.iterator` (Looping)
*   `Symbol.toPrimitive` (Type coercion)
*   `Symbol.toStringTag` (String description)
*   `Symbol.species` (Derived object constructor)

---

### Question 113: What is the `void` operator?

**Answer:**
Evaluates an expression and returns `undefined`. Often used in URIs (`javascript:void(0)`) to prevent side effects like page reload or navigation.

**Example:**
```javascript
const result = void(1 + 1); // undefined (calculation happened)
```

---

### Question 114: What is the `in` operator used for?

**Answer:**
Checks if a property exists in an object or its prototype chain.

**Example:**
```javascript
const car = { make: "Toyota" };
console.log("make" in car); // true
console.log("toString" in car); // true (inherited)
```

---

### Question 115: How does `instanceof` really work internally?

**Answer:**
It checks if the `prototype` property of the constructor appears anywhere in the object's prototype chain.
`obj instanceof Constructor` checks `Constructor.prototype` against `obj.__proto__` chain.

---

### Question 116: How is the `new` keyword implemented?

**Answer:**
1.  Creates a new empty object.
2.  Sets the object's prototype (`__proto__`) to the constructor's `prototype`.
3.  Calls the constructor with `this` bound to the new object.
4.  Returns the object (unless constructor returns a non-primitive).

**Code:**
```javascript
function myNew(Constructor, ...args) {
    const obj = Object.create(Constructor.prototype);
    const result = Constructor.apply(obj, args);
    return (typeof result === 'object' && result !== null) ? result : obj;
}
```

---

### Question 117: What is a “polyfill” in JavaScript?

**Answer:**
A piece of code (usually JavaScript on the web) used to provide modern functionality on older browsers that do not natively support it (e.g., adding `Array.prototype.forEach` to IE8).

---

### Question 118: What is feature detection in JavaScript?

**Answer:**
Checking if a browser supports a specific feature (API/Method) before using it, rather than checking the browser name/version.

**Example:**
```javascript
if ('geolocation' in navigator) {
    // Safe to use geolocation
}
```

---

### Question 119: What is a transpiler and how is it different from a compiler?

**Answer:**
A **transpiler** (Source-to-Source Compiler) translates source code from one language to another at similar abstraction levels (e.g., TS -> JS, ES6 -> ES5 via Babel). A compiler typically translates high-level code to low-level machine code/bytecode.

---

### Question 120: What is the difference between `escape()` and `encodeURIComponent()`?

**Answer:**
*   **`escape()`**: Deprecated. Does not encode `@`, `*`, `+`, etc.
*   **`encodeURIComponent()`**: Standard. Encodes special characters including `?`, `&`, `=` for use in URL query parameters.

---

## 🔹 Execution Context & Scope (131–160)

### Question 121: What are the phases of execution in JS?

**Answer:**
1.  **Creation Phase:** Global object/`this` created. Variables and functions declared (Hoisting).
2.  **Execution Phase:** Code executes line-by-line. Values assignments and function calls happen.

---

### Question 122: How does the call stack work in JavaScript?

**Answer:**
A LIFO (Last In, First Out) stack that tracks function calls. When a function is invoked, a frame is pushed; when it returns, the frame is popped. If the stack exceeds limit (infinite recursion), a "Stack Overflow" occurs.

---

### Question 123: What is the difference between `globalThis`, `window`, and `self`?

**Answer:**
*   **`window`**: Global object in DOM/Browsers.
*   **`self`**: Global object in Web Workers (and Browsers).
*   **`global`**: Global object in Node.js.
*   **`globalThis`**: ES2020 standard way to access the global object across all environments.

---

### Question 124: How do block scopes behave inside loops?

**Answer:**
Using `let` in a `for` loop creates a distinct variable binding for *each iteration*, preventing the classic closure issue where all closures share the same `i`.

**Example:**
```javascript
for (let i = 0; i < 3; i++) {
    setTimeout(() => console.log(i), 100); // 0, 1, 2
}
// with var: 3, 3, 3
```

---

### Question 125: What is tail call optimization (TCO)?

**Answer:**
An optimization where the engine reuses the current stack frame for a function call in the tail position (last action), preventing stack overflow in recursion.
*Note: Only Safari fully implements TCO in JS engines currently.*

---

### Question 126: What is the purpose of the `arguments` object?

**Answer:**
An array-like object available inside non-arrow functions containing all passed arguments. Useful for variadic functions (though `...rest` is preferred now).

**Example:**
```javascript
function sum() {
    return Array.from(arguments).reduce((a, b) => a + b, 0);
}
```

---

### Question 127: How does `eval()` affect scope?

**Answer:**
`eval()` executes code in the **local scope** where it is called (unless indirect call). It can modify local variables, preventing the JS engine from optimizing scope lookups effectively.

---

### Question 128: How does `Function()` constructor work?

**Answer:**
Creates a new function dynamically. Unlike `eval`, functions created with `new Function()` are created in the **global scope** and do not have access to the local closure context.

---

### Question 129: What is the difference between `apply()`, `call()`, and `bind()`?

**Answer:**
*   **`call(thisArg, arg1, arg2)`**: Invokes function immediately with args.
*   **`apply(thisArg, [args])`**: Invokes function immediately with array of args.
*   **`bind(thisArg, arg1)`**: Returns a **new function** with `this` permanently bound.

---

### Question 130: What is the scope of variables declared with `var` in `for` loops?

**Answer:**
`var` is function-scoped. In a loop, the variable leaks to the parent scope and is shared across all iterations.

---

## 🔹 Functions & Functional Programming (161–190)

### Question 131: What is function composition?

**Answer:**
The process of combining two or more functions to produce a new function: `f(g(x))`.

**Example:**
```javascript
const compose = (f, g) => x => f(g(x));
const add1 = x => x + 1;
const square = x => x * x;
const addThenSquare = compose(square, add1);
console.log(addThenSquare(2)); // (2+1)^2 = 9
```

---

### Question 132: How to implement a debounce function manually?

**Answer:**
Delays execution until inactivity.

**Code:**
```javascript
function debounce(fn, delay) {
    let timeoutId;
    return function(...args) {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => {
            fn.apply(this, args);
        }, delay);
    };
}
```

---

### Question 133: How to implement throttling manually?

**Answer:**
Ensures execution at most once per interval.

**Code:**
```javascript
function throttle(fn, interval) {
    let lastTime = 0;
    return function(...args) {
        const now = Date.now();
        if (now - lastTime >= interval) {
            lastTime = now;
            fn.apply(this, args);
        }
    };
}
```

---

### Question 134: What is an arity of a function?

**Answer:**
The number of arguments a function expects. Stored in the `function.length` property.

**Example:**
```javascript
function add(a, b) {}
console.log(add.length); // 2
```

---

### Question 135: What is a higher-order function?

**Answer:**
A function that either:
1.  Takes a function as an argument (e.g., `map`, `setTimeout`).
2.  Returns a function (e.g., `bind`, closures).

---

### Question 136: What is function chaining?

**Answer:**
A pattern where methods return `this` (the object itself), allowing multiple calls to be linked in a single line.

**Example:**
```javascript
class Calc {
    constructor() { this.val = 0; }
    add(n) { this.val += n; return this; }
    sub(n) { this.val -= n; return this; }
}
new Calc().add(10).sub(2);
```

---

### Question 137: What is the difference between `.call(this)` and `.apply(this)`?

**Answer:**
(Duplicate of 129 context). `call` accepts arguments separated by commas, `apply` accepts arguments as an array. `apply` is useful for variadic functions before spread operator existed.

---

### Question 138: How to memoize a recursive function?

**Answer:**
Pass a cache or wrap the logic.

**Example (Fibonacci):**
```javascript
const memo = {};
function fib(n) {
    if (n in memo) return memo[n];
    if (n <= 1) return n;
    return memo[n] = fib(n-1) + fib(n-2);
}
```

---

### Question 139: How can you implement `once()` behavior in JS?

**Answer:**
A higher-order function that runs the target function only once.

**Code:**
```javascript
function once(fn) {
    let ran = false;
    let result;
    return function(...args) {
        if (!ran) {
            ran = true;
            result = fn.apply(this, args);
        }
        return result;
    };
}
```

---

### Question 140: What are function decorators in JavaScript?

**Answer:**
Currently a Stage 3 proposal (common in TS). A decorator is a function that wraps a class, method, or property to modify its behavior (logging, validation) using `@decorator` syntax.

---

## 🔹 OOP & Prototypes (191–220)

### Question 141: How to implement multiple inheritance in JavaScript?

**Answer:**
JS does not support it natively. Uses **Mixins** (copying methods from multiple objects to a target) or Class inheritance chains.

**Example:**
```javascript
const sayHi = { hi() { console.log('Hi'); } };
const sayBye = { bye() { console.log('Bye'); } };

class User {}
Object.assign(User.prototype, sayHi, sayBye);
```

---

### Question 142: What’s the difference between `super()` and `super.prop()`?

**Answer:**
*   **`super()`**: Calls the parent constructor. Mandatory in derived classes before accessing `this`.
*   **`super.method()`**: Calls a specific method on the parent prototype.

---

### Question 143: How does prototype pollution happen?

**Answer:**
An attack where an attacker modifies `Object.prototype` (e.g., via unsafe recursive merges with `__proto__` payload), causing every object in the application to have the injected property. Can lead to DoS or RCE.

---

### Question 144: What is `constructor.name`?

**Answer:**
A read-only property of a function/class that returns its name as a string.

**Example:**
```javascript
class Foo {}
console.log(new Foo().constructor.name); // "Foo"
```

---

### Question 145: How to dynamically add methods to a prototype?

**Answer:**
Assigning functions to the `prototype` object of the constructor.

**Example:**
```javascript
Array.prototype.last = function() {
    return this[this.length - 1];
};
console.log([1, 2, 3].last()); // 3
```

---

### Question 146: What’s the difference between static and instance methods?

**Answer:**
*   **Static:** Defined with `static`. Called on the Class (`User.compare()`). Cannot access `this` instance.
*   **Instance:** Defined normally. Called on objects (`user.sayHi()`).

---

### Question 147: What is mixin pattern in JavaScript?

**Answer:**
A way to add properties/methods to an object from another object without using inheritance. Useful for sharing behavior across unrelated classes.

---

### Question 148: What is object augmentation?

**Answer:**
Adding new properties or methods to an object after it has been created (monkey patching).

---

### Question 149: How is classical inheritance different from prototypal?

**Answer:**
*   **Classical (Java/C++):** Classes are blueprints. Objects are instances. Rigid hierarchy.
*   **Prototypal (JS):** Objects inherit from other objects. More flexible/dynamic.

---

### Question 150: What is the purpose of `Object.getPrototypeOf()`?

**Answer:**
The standard, recommended way to access the prototype (`[[Prototype]]`) of an object (instead of the deprecated `__proto__`).

**Example:**
```javascript
const proto = {};
const obj = Object.create(proto);
console.log(Object.getPrototypeOf(obj) === proto); // true
```
