# JavaScript & TypeScript Fundamentals (Service-Based Companies)

Service-based companies often test your core JavaScript and TypeScript knowledge to ensure you can contribute effectively to enterprise applications. The focus is usually on understanding standard features, ES6+ syntax, asynchronous programming, and basic typing.

## Basic JavaScript Concepts

### 1. What are the differences between `var`, `let`, and `const`?
*   **`var`**: Function-scoped or globally scoped. Can be re-declared and updated. Hoisted and initialized with `undefined`.
*   **`let`**: Block-scoped. Can be updated but not re-declared in the same scope. Hoisted but not initialized (Temporal Dead Zone).
*   **`const`**: Block-scoped. Cannot be updated or re-declared. Must be initialized during declaration. For objects/arrays, the reference cannot change, but properties/elements can be modified.

### 2. Can you explain Closures in JavaScript? Give a practical use case.
A closure is a function that remembers the variables from its lexical scope even after the outer function has finished executing.
**Practical Use Case**: Data encapsulation/privacy.
```javascript
function createCounter() {
    let count = 0; // 'count' is private
    return {
        increment: function() { count++; return count; },
        decrement: function() { count--; return count; }
    };
}
const counter = createCounter();
console.log(counter.increment()); // 1
console.log(counter.count); // undefined (cannot access directly)
```

### 3. What is Hoisting?
Hoisting is a JavaScript mechanism where variable and function declarations are moved to the top of their containing scope before code execution.
*   Function declarations are fully hoisted.
*   `var` is hoisted and initialized as `undefined`.
*   `let` and `const` are hoisted to the block's top but remain in a "Temporal Dead Zone" until their actual declaration line is executed.

### 4. Explain the Event Loop and how JavaScript handles asynchronous operations.
JavaScript is single-threaded. The Event Loop allows Node.js/browsers to perform non-blocking I/O operations by offloading operations to the system kernel whenever possible.
*   **Call Stack**: Executes synchronous code.
*   **Web APIs / Node APIs**: Handles async tasks like `setTimeout` or HTTP requests.
*   **Callback Queue / Task Queue**: Holds callbacks ready to be executed (from `setTimeout`, etc.).
*   **Microtask Queue**: Holds Promise `.then()` & `.catch()` callbacks (higher priority than Callback Queue).
*   **Event Loop**: Continuously checks if the Call Stack is empty. If it is, it pushes tasks from the Microtask Queue first, then the Task Queue, onto the Call Stack.

### 5. What is the difference between `==` and `===`?
*   `==` (Loose equality): Compares values after performing type coercion if they are of different types. (`'5' == 5` is `true`)
*   `===` (Strict equality): Compares both value and type without coercion. (`'5' === 5` is `false`)

## ES6+ Features

### 6. What are Arrow Functions, and how do they differ from normal functions?
Arrow functions provide a shorter syntax for writing function expressions.
**Differences**:
*   They don't have their own `this` binding; they inherit `this` from the parent lexical scope.
*   They don't have an `arguments` object.
*   They cannot be used as constructors (no `new` keyword).

### 7. Explain Destructuring, Rest parameters, and Spread syntax.
*   **Destructuring**: Extracting properties from objects or elements from arrays into variables easily.
    `const { name, age } = userObject;`
*   **Rest parameter (`...args`)**: Condenses multiple elements into a single array (used in function arguments).
    `function log(first, ...rest) { }`
*   **Spread syntax (`...arr`)**: Expands an iterable (like an array or object) into individual elements.
    `const combined = [...arr1, ...arr2];`

### 8. How do Promises work? What are `async/await`?
*   **Promise**: An object representing the eventual completion (or failure) of an asynchronous operation. It has three states: Pending, Fulfilled, Rejected.
*   **`async/await`**: Syntactic sugar over Promises introduced in ES2017. `async` makes a function return a Promise, and `await` pauses the execution of the `async` function until the Promise settles, making async code look synchronous and easier to read.

## TypeScript Fundamentals

### 9. Why use TypeScript over plain JavaScript?
*   **Static Typing**: Catches errors at compile-time rather than run-time.
*   **Better IDE Support**: Provides enhanced intellisense, autocompletion, and refactoring tools.
*   **Code Maintainability**: Self-documenting code through types and interfaces, crucial for large teams and enterprise projects.
*   **Modern Features**: Allows using the latest JavaScript features while compiling down to older, browser-compatible versions.

### 10. What is the difference between an `Interface` and a `Type` in TypeScript?
While largely interchangeable, there are subtle differences:
*   **Interface**: Best used for declaring shapes of objects or class contracts. They support *declaration merging* (multiple declarations with the same name merge into one).
*   **Type Alias**: Can represent primitive types, unions (`string | number`), intersections, and tuples. They do *not* support declaration merging.
Generally, use `interface` for object shapes and `type` for complex type definitions (unions/primitives).

### 11. What are Generics in TypeScript? Provide an example.
Generics allow you to create reusable components that can work over a variety of types rather than a single one, while preserving type safety.
```typescript
function getFirstElement<T>(arr: T[]): T {
    return arr[0];
}

const num = getFirstElement<number>([1, 2, 3]); // returns a number
const str = getFirstElement<string>(['a', 'b', 'c']); // returns a string
```

### 12. Explain the `any`, `unknown`, and `never` types.
*   `any`: Opts out of type checking. Should be avoided as it defeats the purpose of TS.
*   `unknown`: The type-safe counterpart of `any`. You can assign anything to `unknown`, but you cannot perform operations on it until you assert or narrow its type.
*   `never`: Represents a value that never occurs. Often used as the return type for functions that throw an error or contain infinite loops, or in exhaustive `switch` statements.
