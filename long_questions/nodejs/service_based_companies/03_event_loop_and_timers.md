# 🚥 03 — Event Loop & Timers
> **Most Asked in Service-Based Companies** | 🚥 Difficulty: Medium to Hard

---

## 🔑 Must-Know Topics
- Explain the Node.js Event Loop
- Phases of the Event Loop (Timers, Pending, Poll, Check, Close)
- Microtasks (`process.nextTick()`, Promises) vs Macrotasks
- `setTimeout()` vs `setImmediate()` vs `process.nextTick()`
- How Node.js handles I/O operations non-blockingly

---

## ❓ Frequently Asked Questions

### Q1. What is the Event Loop in Node.js and why is it important?

**Answer:**
Node.js is single-threaded, meaning it can only execute one operation at a time. The **Event Loop** is the mechanism that allows Node.js to perform non-blocking I/O operations despite being single-threaded. It achieves this by offloading operations to the system kernel whenever possible.

When an operation (like reading a file or making a network request) is initiated, Node.js hands it off to the C++ layer (libuv). While the operation happens in the background, Node.js continues executing the rest of the script. Once the background operation completes, its callback is pushed onto a queue, and the Event Loop eventually picks it up and executes it.

---

### Q2. What are the different phases of the Event Loop?

**Answer:**
The Event Loop executes in specific phases, processing different types of queues. The core phases (managed by libuv) are:

1. **Timers:** Executes callbacks scheduled by `setTimeout()` and `setInterval()`.
2. **Pending Callbacks:** Executes I/O callbacks deferred to the next loop iteration (e.g., some TCP errors).
3. **Idle, Prepare:** Only used internally.
4. **Poll:** Retrieves new I/O events; executes I/O related callbacks (almost all with the exception of close callbacks, the ones scheduled by timers, and `setImmediate()`); node will block here when appropriate.
5. **Check:** Executes callbacks scheduled by `setImmediate()`.
6. **Close Callbacks:** Executes `close` callbacks (e.g., `socket.on('close', ...)`).

Between each phase, Node.js checks the **Microtask Queue** (callbacks from `process.nextTick()` and resolved Promises) and executes all of them before moving to the next phase.

---

### Q3. What is the difference between `setTimeout` and `setImmediate`?

**Answer:**
Both are used to schedule code execution for the future, but they are executed in different phases of the Event Loop.

- **`setTimeout(callback, delay)`:** Schedules a callback to be executed after a minimum threshold in ms has elapsed. It is executed in the **Timers** phase.
- **`setImmediate(callback)`:** Schedules a callback to execute immediately *after* the current Poll phase completes. It is executed in the **Check** phase.

**Output without I/O context (unpredictable):**
```javascript
setTimeout(() => console.log('timeout'), 0);
setImmediate(() => console.log('immediate'));
// Execution order depends on machine performance constraints
```

**Output within an I/O context (predictable):**
```javascript
const fs = require('fs');

fs.readFile('data.txt', () => {
    setTimeout(() => console.log('timeout'), 0);
    setImmediate(() => console.log('immediate'));
});
// 'immediate' will ALWAYS run first because the Check phase immediately follows the Poll phase (where the file read callback runs).
```

---

### Q4. What is `process.nextTick()`? How is it different from `setTimeout(fn, 0)`?

**Answer:**

**`process.nextTick(callback)`** is not technically part of the Event Loop. It adds the callback to the **"nextTickQueue"**. In Node.js, the `nextTickQueue` is always checked *after the current operation completes* and *before* the Event Loop continues to its next phase.

- **`process.nextTick()`:** Executes immediately after the current synchronous code finishes, completely bypassing the Event Loop phases. It is the highest priority asynchronous operation.
- **`setTimeout(fn, 0)`:** Executes in the Timers phase of the Event Loop, after the current execution context and the microtask queues have been cleared.

**Example:**
```javascript
console.log('Start');

setTimeout(() => console.log('Timeout'), 0);
process.nextTick(() => console.log('NextTick'));
Promise.resolve().then(() => console.log('Promise'));

console.log('End');

// Output:
// Start
// End
// NextTick
// Promise
// Timeout
```

---

### Q5. Can you starve the Event Loop? How?

**Answer:**
Yes. Because Node.js is single-threaded, if a task takes too long, it will block the Event Loop, preventing other tasks from executing.

**Ways to starve the Event Loop:**
1. **CPU Intensive Tasks:** Running heavy synchronous calculations (like image processing, large loops, or cryptography) on the main thread.
2. **Infinite Loops:** An accidental `while(true)` loop.
3. **Synchronous I/O:** Using `fs.readFileSync()` instead of `fs.readFile()` in a web server request handler.
4. **Recursive `process.nextTick()`:** If an operation recursively calls `process.nextTick()`, it will prevent the Event Loop from ever moving to the next phase (like Poll or Timers), starving all I/O operations.

```javascript
// Example of recursive nextTick (bad practice)
function recursiveTick() {
    process.nextTick(recursiveTick);
}
// This will block the event loop forever!
```

**Solution:** Use Worker Threads for CPU intensive tasks, or break operations into smaller chunks using `setImmediate()`.
