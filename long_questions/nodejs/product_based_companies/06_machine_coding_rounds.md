# 💻 06 — Machine Coding Rounds
> **Most Asked in Product-Based Companies** | 💻 Difficulty: Hard

---

## 🔑 Must-Know Topics
These questions require you to write fully working, clean, production-ready code with proper error handling and architecture within 60-90 minutes.

- Writing Custom Middlewares
- Rate Limiting implementations
- High Concurrency Handling (Task Queues)
- In-memory Pub/Sub systems

---

## ❓ Frequently Asked Questions

### Q1. Implement an In-Memory Rate Limiter Middleware for Express. By IP, Max 100 requests per 15 minutes.

**Answer:**

**The Sliding Window Log OR Token Bucket approach is ideal.** Here is a straightforward implementation using an in-memory Store with `setInterval` purging to avoid memory leaks.

```javascript
const express = require('express');
const app = express();

// In-memory store
const requestStore = new Map();

// Configuration
const WINDOW_SIZE_MS = 15 * 60 * 1000; // 15 mins
const MAX_REQUESTS = 100;

// Middleware
const rateLimiter = (req, res, next) => {
    const ip = req.ip || req.connection.remoteAddress;
    const now = Date.now();

    if (!requestStore.has(ip)) {
        requestStore.set(ip, [now]);
        return next();
    }

    let timestamps = requestStore.get(ip);
    
    // Filter out requests older than the window
    timestamps = timestamps.filter(time => time > now - WINDOW_SIZE_MS);
    
    if (timestamps.length >= MAX_REQUESTS) {
        requestStore.set(ip, timestamps); // update filtered list
        return res.status(429).json({ error: "Too many requests, please try again later." });
    }

    timestamps.push(now);
    requestStore.set(ip, timestamps);
    
    // Set headers
    res.set('X-RateLimit-Limit', MAX_REQUESTS);
    res.set('X-RateLimit-Remaining', MAX_REQUESTS - timestamps.length);

    next();
};

// Cleanup routine to prevent Memory Leaks! (Crucial for Interviews)
setInterval(() => {
    const now = Date.now();
    for (let [ip, timestamps] of requestStore.entries()) {
        const validTimestamps = timestamps.filter(t => t > now - WINDOW_SIZE_MS);
        if (validTimestamps.length === 0) {
            requestStore.delete(ip); // clear inactive users
        } else {
            requestStore.set(ip, validTimestamps);
        }
    }
}, WINDOW_SIZE_MS);

// Usage
app.use(rateLimiter);

app.get('/api/data', (req, res) => res.json({ secret: 'data' }));

app.listen(3000, () => console.log('Server running.'));
```
*Note: In production, explicitly state you would use Redis, but implement this purely in JS to demonstrate algorithmic knowledge.*

---

### Q2. Implement an asynchronous Task Queue capable of running N concurrent tasks.

**Answer:**
Product companies love testing your knowledge of concurrency control, Promises, and the event loop.

We need a Queue that accepts Promise-returning tasks and executes them with a specified concurrency limit, guaranteeing that no more than `limit` tasks are running simultaneously.

```javascript
class TaskQueue {
    constructor(concurrencyLimit) {
        this.limit = concurrencyLimit;
        this.queue = [];
        this.running = 0;
    }

    // Add a task (a function returning a Promise or value)
    enqueue(task) {
        return new Promise((resolve, reject) => {
            // Push wrapper to the queue
            this.queue.push(async () => {
                try {
                    const result = await task();
                    resolve(result);
                } catch (error) {
                    reject(error);
                }
            });

            this.processNext(); // Attempt to start processing
        });
    }

    processNext() {
        if (this.running >= this.limit || this.queue.length === 0) {
            return; // We are at capacity or idle
        }
        
        // Take next task
        const nextTask = this.queue.shift();
        this.running++;

        // Execute it
        nextTask().finally(() => {
            this.running--;
            this.processNext(); // Trigger the next one when finished
        });
    }
}

// === Testing the Queue ===

const q = new TaskQueue(2); // Only allow 2 tasks concurrently

// A mock async task factory
const createTask = (id, delayMs) => {
    return () => new Promise(resolve => {
        console.log(`Task ${id} starting...`);
        setTimeout(() => {
            console.log(`Task ${id} COMPLETED after ${delayMs}ms.`);
            resolve(`Result ${id}`);
        }, delayMs);
    });
};

q.enqueue(createTask(1, 2000)).then(console.log);
q.enqueue(createTask(2, 2000)).then(console.log);
q.enqueue(createTask(3, 1000)).then(console.log);
// Task 1 and 2 start immediately.
// Task 3 is queued.
// After 2s, Task 1 and 2 complete. Task 3 starts.
// Task 3 completes 1s later.
```

---

### Q3. Write a Custom Event Emitter in JavaScript.

**Answer:**
This question tests your understanding of the Observer Pattern, Hash Maps, and JavaScript arrays.

```javascript
class CustomEmitter {
    constructor() {
        this.events = {}; // { eventName: [listeners] }
    }

    on(eventName, listener) {
        if (!this.events[eventName]) {
            this.events[eventName] = [];
        }
        this.events[eventName].push(listener);
    }

    emit(eventName, ...args) {
        if (this.events[eventName]) {
            // Loop through all registered listeners and execute them
            this.events[eventName].forEach(listener => {
                listener.apply(this, args);
            });
        }
    }

    off(eventName, listenerToRemove) {
        if (this.events[eventName]) {
            this.events[eventName] = this.events[eventName].filter(
                listener => listener !== listenerToRemove
            );
        }
    }

    once(eventName, listener) {
        const onceWrapper = (...args) => {
            listener.apply(this, args); // Execute the actual listener
            this.off(eventName, onceWrapper); // Immediately remove the wrapper
        };
        this.on(eventName, onceWrapper);
    }
}

// === Usage ===
const emitter = new CustomEmitter();

const callback = (data) => console.log('Received:', data);
emitter.on('data', callback);

emitter.emit('data', 'Hello World'); // Prints: Received: Hello World
emitter.off('data', callback);
emitter.emit('data', 'Hidden'); // Does nothing

emitter.once('startup', () => console.log('Booting!'));
emitter.emit('startup'); // Prints: Booting!
emitter.emit('startup'); // Does nothing
```
