# 🟡 02 — Asynchronous Programming in Node.js
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Synchronous vs Asynchronous code
- Callbacks and Callback Hell
- Promises and their specific states
- `async` / `await`
- Error handling in async operations
- `.then()`, `.catch()`, `.finally()`
- `Promise.all()`, `Promise.allSettled()`, `Promise.race()`

---

## ❓ Frequently Asked Questions

### Q1. What is the difference between Synchronous and Asynchronous functions in Node.js?

**Answer:**

**Synchronous (Blocking):**
Code is executed line by line. The next line of code cannot run until the current line has finished executing. This blocks the single thread in Node.js.
```javascript
const fs = require('fs');

console.log('Start');
// Blocks execution until file is completely read
const data = fs.readFileSync('file.txt', 'utf8');
console.log(data);
console.log('End');
```

**Asynchronous (Non-Blocking):**
Code does not wait for an operation to complete. It registers a callback/promise and moves to the next line. When the operation finishes in the background, the callback is pushed to the callback queue.
```javascript
const fs = require('fs');

console.log('Start');
fs.readFile('file.txt', 'utf8', (err, data) => {
    if (err) throw err;
    console.log(data);
});
console.log('End'); // Prints before the file data
```

---

### Q2. What is Callback Hell and how do you avoid it?

**Answer:**
Callback Hell (often called the "Pyramid of Doom") occurs when multiple asynchronous operations depend on each other, leading to deeply nested callbacks. It makes the code hard to read, maintain, and debug.

**Example of Callback Hell:**
```javascript
getData(function(a){
    getMoreData(a, function(b){
        getMoreData(b, function(c){ 
            getMoreData(c, function(d){ 
                console.log(d);
            });
        });
    });
});
```

**How to avoid it:**
1. **Use Promises:** Chain `.then()` to keep code flat.
2. **Use `async/await`:** Write asynchronous code that looks synchronous.
3. **Modularization:** Break callbacks into named, separate functions instead of anonymous inline arrow functions.

---

### Q3. Explain Promises in Node.js. What are its different states?

**Answer:**
A Promise is an object representing the eventual completion (or failure) of an asynchronous operation and its resulting value.

**States of a Promise:**
1. **Pending:** The initial state. The operation has not completed yet.
2. **Fulfilled (Resolved):** The operation completed successfully.
3. **Rejected:** The operation failed (e.g., network error, file not found).

**Example:**
```javascript
const fs = require('fs').promises;

fs.readFile('file.txt', 'utf8')
    .then(data => {
        console.log("File read successfully: ", data); // Fulfilled
    })
    .catch(err => {
        console.error("Error reading file: ", err); // Rejected
    })
    .finally(() => {
        console.log("Operation finished."); // Always runs
    });
```

---

### Q4. How does `async/await` work? Is it different from Promises?

**Answer:**
`async / await` is syntactic sugar over Promises. It provides a more readable, synchronous-looking way to write asynchronous code. It does not replace Promises; it uses them under the hood.

- `async` keyword placed before a function ensures the function always returns a Promise.
- `await` keyword can only be used inside an `async` function. It pauses the execution of the function until the Promise settles (resolves or rejects).

```javascript
async function fetchUser() {
    try {
        console.log('Fetching...');
        const user = await db.getUser(1); // Pauses here until resolved
        console.log(user);
    } catch (error) {
        console.error('Failed to fetch user:', error);
    }
}
```

---

### Q5. What is the difference between `Promise.all()` and `Promise.allSettled()`?

**Answer:**

**`Promise.all(iterable)`:**
- Takes an array of promises.
- **Resolves:** When *all* promises in the array resolve. Returns an array of resolved values.
- **Rejects:** If *any single* promise rejects, `Promise.all` immediately rejects with that error, ignoring the rest (fail-fast behavior).
```javascript
Promise.all([promise1, promise2])
    .then(results => console.log(results))
    .catch(err => console.log('One failed:', err));
```

**`Promise.allSettled(iterable)`:**
- Takes an array of promises.
- **Resolves:** When *all* promises settle (either resolve or reject).
- Never rejects early. Returns an array of objects describing the outcome of each promise (`{status: "fulfilled", value: val}` or `{status: "rejected", reason: err}`).
```javascript
Promise.allSettled([promise1, promise2])
    .then(results => {
        results.forEach(result => console.log(result.status));
    });
```

---

### Q6. How do you handle errors in asynchronous code?

**Answer:**
1. **In Callbacks:** The "error-first callback" pattern is standard in Node.js. The first argument of the callback is reserved for the error object.
   ```javascript
   fs.readFile('data.txt', (err, data) => {
       if (err) return console.error("Failed:", err);
       console.log("Success:", data);
   });
   ```
2. **In Promises:** Use `.catch()` at the end of the promise chain.
   ```javascript
   fetchData().then(data => process(data)).catch(err => console.error(err));
   ```
3. **In Async/Await:** Use standard `try...catch` blocks.
   ```javascript
   async function doWork() {
       try {
           const data = await fetchData();
       } catch (error) {
           console.error("Caught an error:", error);
       }
   }
   ```
