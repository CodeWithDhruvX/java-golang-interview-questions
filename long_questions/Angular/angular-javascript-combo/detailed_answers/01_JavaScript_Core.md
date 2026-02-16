# 🧠 JavaScript Core Fundamentals (Detailed Answers)

## 1. What is the difference between var, let, and const?

| Feature | `var` | `let` | `const` |
| :--- | :--- | :--- | :--- |
| **Scope** | Function Scope | Block Scope (`{}`) | Block Scope (`{}`) |
| **Hoisting** | Yes (initialized as `undefined`) | Yes (in **TDZ** - Temporal Dead Zone) | Yes (in **TDZ**) |
| **Re-declaration** | Allowed | Not Allowed | Not Allowed |
| **Re-assignment** | Allowed | Allowed | Not Allowed |
| **Global Object Property** | Yes (window.var) | No | No |

**Example:**
```javascript
// Scope
if (true) {
  var a = 10;
  let b = 20;
}
console.log(a); // 10
// console.log(b); // ReferenceError: b is not defined

// Hoisting
console.log(x); // undefined
var x = 5;

// console.log(y); // ReferenceError: Cannot access 'y' before initialization
let y = 10;
```

---

## 2. What is hoisting?
Hoisting is a JavaScript mechanism where variable and function declarations are moved to the top of their scope using the compilation phase.

*   **Variable Hoisting**: `var` is hoisted and initialized with `undefined`. `let` and `const` are hoisted but stay in the **Temporal Dead Zone** (uninitialized).
*   **Function Hoisting**: Function declarations are fully hoisted (can be called before definition). Arrow functions behave like variables (var/let/const rules apply).

**Example:**
```javascript
greet(); // "Hello"
function greet() {
  console.log("Hello");
}

// sayHi(); // TypeError: sayHi is not a function
var sayHi = function() {
  console.log("Hi");
};
```

---

## 3. What is Temporal Dead Zone (TDZ)?
The TDZ is the time usage between the entering of scope and the variable's actual declaration. Accessing a `let` or `const` variable in its TDZ throws a `ReferenceError`.

**Example:**
```javascript
{
  // TDZ starts here
  // console.log(z); // ReferenceError
  let z = 10; // TDZ ends here
  console.log(z); // 10
}
```

---

## 4. What is closure? Give a real use case.
A **Closure** is a function bundled together with references to its surrounding state (the **lexical environment**). It allows an inner function to access variables from an outer function's scope even after the outer function has finished executing.

**Real Use Case: Data Privacy / Encapsulation**
```javascript
function createCounter() {
  let count = 0; // Private variable
  return {
    increment: function() {
      count++;
      console.log(count);
    },
    getCount: function() {
      return count;
    }
  };
}

const counter = createCounter();
counter.increment(); // 1
counter.increment(); // 2
console.log(counter.count); // undefined (cannot access directly)
```

---

## 5. What is the Event Loop?
JavaScript is single-threaded. The **Event Loop** is the mechanism that allows it to perform non-blocking operations.

**How it works:**
1.  **Call Stack**: Executes synchronous code.
2.  **Web APIs**: Handles async operations (setTimeout, fetch, DOM events).
3.  **Callback Queue (Task Queue)**: Stores callbacks from Web APIs (e.g., setTimeout).
4.  **Microtask Queue**: Stores high-priority tasks (Promises, queueMicrotask).
5.  **Role of Event Loop**: It constantly checks if the **Call Stack** is empty. If empty, it prioritizes the **Microtask Queue**, executing all tasks there, and *then* moves one task from the **Callback Queue** to the Stack.

---

## 6. What is the Call Stack?
Data structure that records where in the program we are. If we step into a function, we put it on the top of the stack. If we return from a function, we pop it off the top of the stack.

---

## 7. What are Microtasks and Macrotasks?
*   **Microtasks (High Priority)**: `Promise.then`, `queueMicrotask`, `MutationObserver`.
*   **Macrotasks (Low Priority)**: `setTimeout`, `setInterval`, `setImmediate`, `I/O`.

**Order of Execution:**
1.  Sync Code (Call Stack)
2.  All Microtasks
3.  One Macrotask
4.  All Microtasks (again)
5.  ...repeat

**Example:**
```javascript
console.log('Start'); // 1. Sync

setTimeout(() => {
  console.log('Timeout'); // 4. Macrotask
}, 0);

Promise.resolve().then(() => {
  console.log('Promise'); // 3. Microtask
});

console.log('End'); // 2. Sync

// Output: Start, End, Promise, Timeout
```

---

## 8. Difference between == and ===?
*   `==` (Abstract Equality): Converts types (type coercion) before comparing. `5 == '5'` is `true`.
*   `===` (Strict Equality): Checks value **and** type. `5 === '5'` is `false`. **Always use this.**

---

## 9. What is type coercion?
Automatic or implicit conversion of values from one data type to another.
*   **Explicit**: `Number('123')`
*   **Implicit**: `'5' - 3` (String converted to Number, result 2) or `'5' + 3` (Number converted to String, result '53').

---

## 10. What are primitive and non-primitive types?
*   **Primitive (Value Types)**: Stored directly in the stack. Immutable.
    *   `string`, `number`, `boolean`, `null`, `undefined`, `symbol`, `bigint`.
*   **Non-Primitive (Reference Types)**: Stored in the heap, reference stored in stack. Mutable.
    *   `Object`, `Array`, `Function`.

---

## 11. What is undefined vs null?
*   **undefined**: "I exist but have no value yet." Default value of uninitialized variables.
*   **null**: "I have a value, and that value is *nothing*." Intentionally assigned.

---

## 12. What is NaN?
"Not a Number". It's a numeric data type containing an invalid number.
*   `typeof NaN === 'number'`
*   `NaN === NaN` is `false`.
*   Check using `Number.isNaN(value)`.

---

## 13. What is truthy and falsy?
Values that translate to `true` or `false` in a boolean context.
*   **Falsy**: `false`, `0`, `""` (empty string), `null`, `undefined`, `NaN`.
*   **Truthy**: Everything else (e.g., `'0'`, `'false'`, `[]`, `{}`).

---

## 14. What is a higher-order function?
A function that either:
1.  Takes one or more functions as arguments (e.g., `map`, `filter`, `setTimeout`).
2.  Returns a function (e.g., Closure).

**Example:**
```javascript
function multiplier(factor) {
  return function(number) {
    return number * factor;
  };
}
const double = multiplier(2);
console.log(double(5)); // 10
```

---

## 15. What is callback hell?
A situation where callbacks are nested within other callbacks several levels deep, making code difficult to read and maintain.
**Solution**: Use `Promises` or `async/await`.
