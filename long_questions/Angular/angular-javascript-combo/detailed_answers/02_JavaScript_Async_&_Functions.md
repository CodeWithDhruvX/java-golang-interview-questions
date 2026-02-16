# ⚡ JavaScript Async, Functions & 'this' (Detailed Answers)

## 16. What is a Promise?
A **Promise** is an object representing the eventual completion or failure of an asynchronous operation.

**States of a Promise:**
1.  **Pending**: Initial state, neither fulfilled nor rejected.
2.  **Fulfilled**: Operation completed successfully.
3.  **Rejected**: Operation failed.

## 17. How do you construct a Promise?
```javascript
const myPromise = new Promise((resolve, reject) => {
  const success = true; // Simulating async operation
  if (success) {
    resolve("Data fetched!");
  } else {
    reject("Failed to fetch.");
  }
});

myPromise
  .then(data => console.log(data))
  .catch(error => console.error(error));
```

## 18. What is async/await?
Syntactic sugar built on top of Promises. It allows you to write asynchronous code that looks like synchronous code.

**Key difference**: `await` pauses the execution of the `async` function until the Promise settles.

## 19. Difference between Promise and async/await?
*   **Promise**: Chain `.then()` and `.catch()`. Can lead to nesting (callback hell lite).
*   **async/await**: Linear, readable code. Use `try/catch` for error handling.

## 20. What happens if you don't use await?
The function returns immediately with a **pending Promise**. The code following the call executes **before** the asynchronous operation completes.

## 21. How does setTimeout work internally?
It registers a callback with the Web API which starts a timer. Once the timer expires, the callback is pushed to the **Macrotask Queue**. The Event Loop waits for the Call Stack to clear before executing it.

## 22. What is fetch API?
A modern, Promise-based interface for making HTTP requests (Native in browsers).
*   Use `response.ok` (boolean) to check for success (status 200-299).
*   Use `response.json()` (async) to parse JSON body.

## 23. How do you handle API errors?
With `async/await`, always wrap the code in a `try/catch` block.

```javascript
async function fetchData() {
  try {
    const response = await fetch('https://api.example.com/data');
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Fetch error:', error);
  }
}
```

## 24. What is Promise.all?
Takes an iterable of Promises and returns a single Promise that fulfills when **all** of the promises fulfill.
*   **Wait**: Parallel execution, waits for slowest.
*   **Reject**: Fails immediately if **any** promise rejects.

## 25. What is Promise.race?
Returns a promise that settles as soon as the **first** promise in the iterable settles (resolves or rejects). Useful for timeouts.

---

## 26. What is `this` in JavaScript?
`this` refers to the object that is executing the current function. Its value depends on **how** the function is called.

## 27. What is `this` in global scope, object method, and arrow function?
| Context | refer to |
| :--- | :--- |
| **Global Scope** | Global Object (`window` in browser, `global` in Node). |
| **Object Method** | The object that owns the method. |
| **Function Call** | Global Object (undefined in strict mode). |
| **Constructor** | The newly created instance. |
| **Arrow Function** | **Lexical `this`** (inherits from surrounding scope). |

## 28. Example of Arrow Function `this`
```javascript
const obj = {
  name: "MyObject",
  regularFunc: function() {
    console.log(this.name); // "MyObject"
  },
  arrowFunc: () => {
    console.log(this.name); // undefined (inherits from global scope)
  }
};
```

## 29. What is bind, call, apply?
Used to explicitly set the value of `this`.

*   **Call**: Invokes function immediately. Arguments passed individually. `func.call(obj, 1, 2)`.
*   **Apply**: Invokes function immediately. Arguments passed as an array. `func.apply(obj, [1, 2])`.
*   **Bind**: Returns a **new function** with `this` permanently set. `const newFunc = func.bind(obj)`.

## 30. What is function currying?
A technique of evaluating a function with multiple arguments, into a sequence of functions with a single argument.
`f(a, b, c)` -> `f(a)(b)(c)`

```javascript
const add = (a) => (b) => a + b;
const add5 = add(5);
console.log(add5(10)); // 15
```

## 31. What are IIFE (Immediately Invoked Function Expressions)?
Functions that run as soon as they are defined.
Used to create a local scope and avoid polluting the global namespace.
```javascript
(function() {
    var secret = "I am hidden";
})();
// console.log(secret); // ReferenceError
```

## 32. What is prototype?
Every JavaScript object has a private property which holds a link to another object called its **prototype**. That prototype object has a prototype of its own, and so on until an object is reached with `null` as its prototype.

## 33. What is prototypal inheritance?
When you access a property on an object, JavaScript first checks the object itself. If not found, it looks at the object's prototype, then the prototype's prototype, and so on up the prototype chain.

```javascript
const animal = { eats: true };
const rabbit = Object.create(animal);
console.log(rabbit.eats); // true (inherited)
```
