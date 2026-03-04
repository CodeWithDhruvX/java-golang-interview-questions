# 📘 04 — Asynchronous JavaScript
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Event loop basics
- Callbacks and callback hell
- Promises: `.then`, `.catch`, `.finally`
- `async`/`await` with error handling
- `Promise.all`, `Promise.allSettled`
- AJAX vs Fetch API

---

## ❓ Most Asked Questions

### Q1. How does the JavaScript Event Loop work?

```javascript
// JavaScript is single-threaded — one operation runs at a time
// Event loop coordinates execution across different queues

// Queues:
// 1. Call Stack: where JS code executes
// 2. Web APIs: browser handles async operations (timers, fetch, DOM events)
// 3. Microtask Queue: Promise callbacks — higher priority
// 4. Macrotask (Callback) Queue: setTimeout, setInterval, events — lower priority

// Execution order:
// 1. Run all synchronous code (clear call stack)
// 2. Process ALL microtasks (Promises)
// 3. Render (browser only, between macrotasks)
// 4. Process ONE macrotask (setTimeout, etc.)
// 5. Repeat from step 2

console.log("Start");    // 1st — synchronous

setTimeout(() => {
    console.log("Timer"); // 4th — macrotask
}, 0);

Promise.resolve()
    .then(() => console.log("Promise 1"))  // 3rd — microtask
    .then(() => console.log("Promise 2")); // still microtask (chain)

console.log("End");      // 2nd — synchronous

// Output: Start → End → Promise 1 → Promise 2 → Timer

// Why this matters: don't rely on setTimeout to be "exactly" after a delay
// Actual delay = specified delay + time to clear current task + microtasks
```

---

### Q2. What is callback hell and how to avoid it?

```javascript
// Callback hell: deeply nested callbacks (Pyramid of Doom)
// Common in Node.js file operations and old-style APIs

// ❌ Callback hell (hard to read, error-prone)
getUser(userId, (err, user) => {
    if (err) return handleError(err);
    
    getOrders(user.id, (err, orders) => {
        if (err) return handleError(err);
        
        getOrderDetails(orders[0].id, (err, details) => {
            if (err) return handleError(err);
            
            sendConfirmation(user.email, details, (err, result) => {
                if (err) return handleError(err);
                console.log("Done:", result); // indentation keeps growing!
            });
        });
    });
});

// ✅ Solution 1: Named functions (flatten nesting)
function handleOrderDetails(err, details) {
    if (err) return handleError(err);
    sendConfirmation(user.email, details, handleConfirmation);
}

// ✅ Solution 2: Promises (chain instead of nesting)
getUser(userId)
    .then(user   => getOrders(user.id))
    .then(orders => getOrderDetails(orders[0].id))
    .then(details => sendConfirmation(user.email, details))
    .then(result  => console.log("Done:", result))
    .catch(err    => handleError(err));

// ✅ Solution 3: async/await (most readable — looks synchronous)
async function processOrder(userId) {
    const user    = await getUser(userId);
    const orders  = await getOrders(user.id);
    const details = await getOrderDetails(orders[0].id);
    const result  = await sendConfirmation(user.email, details);
    console.log("Done:", result);
}
```

---

### Q3. Explain `async/await` error handling patterns.

```javascript
// Pattern 1: try/catch (most common)
async function fetchProfile(userId) {
    try {
        const user   = await getUser(userId);
        const profile = await getProfile(user.id);
        return profile;
    } catch (err) {
        // Catches errors from ANY await in the try block
        console.error("Failed to load profile:", err.message);
        return null; // graceful fallback
    }
}

// Pattern 2: .catch on individual awaits (handle each separately)
async function fetchData(url) {
    const response = await fetch(url).catch(err => {
        console.error("Network error:", err);
        return null;
    });

    if (!response) return { error: "Network failed" };

    const data = await response.json().catch(() => null);
    return data || { error: "Invalid JSON" };
}

// Pattern 3: Result pattern (no exceptions, explicit error state)
async function safeAsync(asyncFn) {
    try {
        const data = await asyncFn();
        return { data, error: null };
    } catch (error) {
        return { data: null, error };
    }
}

// Usage
const { data: users, error } = await safeAsync(() => fetch('/api/users').then(r => r.json()));
if (error) {
    showError(error.message);
} else {
    displayUsers(users);
}

// Pattern 4: finally for cleanup (runs regardless of success/failure)
async function uploadFile(file) {
    setLoading(true);
    try {
        const url = await uploadToS3(file);
        await saveUrl(url); // save to database
        showSuccess("Upload complete!");
        return url;
    } catch (err) {
        showError("Upload failed: " + err.message);
        return null;
    } finally {
        setLoading(false); // ALWAYS runs — even if error thrown
    }
}
```

---

### Q4. Implement a simple `fetch`-based API service class.

```javascript
// Practical: reusable API service for service-based projects

class ApiService {
    #baseUrl;
    #headers;

    constructor(baseUrl) {
        this.#baseUrl = baseUrl;
        this.#headers = { 'Content-Type': 'application/json' };
    }

    setAuthToken(token) {
        this.#headers['Authorization'] = `Bearer ${token}`;
        return this;
    }

    async #request(method, endpoint, data = null) {
        const config = {
            method,
            headers: { ...this.#headers }
        };

        if (data) config.body = JSON.stringify(data);

        const response = await fetch(this.#baseUrl + endpoint, config);

        // Handle non-2xx responses
        if (!response.ok) {
            let errorMessage = `HTTP ${response.status}`;
            try {
                const err = await response.json();
                errorMessage = err.message || errorMessage;
            } catch {}
            throw new Error(errorMessage);
        }

        // Handle empty responses (204 No Content)
        if (response.status === 204) return null;

        return response.json();
    }

    get(endpoint)           { return this.#request('GET', endpoint); }
    post(endpoint, data)    { return this.#request('POST', endpoint, data); }
    put(endpoint, data)     { return this.#request('PUT', endpoint, data); }
    patch(endpoint, data)   { return this.#request('PATCH', endpoint, data); }
    delete(endpoint)        { return this.#request('DELETE', endpoint); }
}

// Usage in a project
const api = new ApiService("https://api.yourcompany.com");
api.setAuthToken(localStorage.getItem("token"));

// User operations
const users    = await api.get('/users');
const newUser  = await api.post('/users', { name: "Alice", email: "alice@example.com" });
const updated  = await api.put(`/users/${userId}`, { name: "Alice Smith" });
await api.delete(`/users/${userId}`);
```

---

### Q5. What is the difference between synchronous and asynchronous code?

```javascript
// Synchronous: code runs line by line — next line waits for current to finish
// ✅ Predictable ❌ Blocks execution during long operations

console.log("Step 1");
const result = computeHeavily(); // blocks for 2 seconds
console.log("Step 2"); // runs only after computation

// Asynchronous: initiates operation, continues execution, handles result later
// ✅ Non-blocking ❌ More complex control flow

console.log("Step 1");
fetch('/api/data').then(r => r.json()).then(data => {
    console.log("Data received:", data); // runs later when response arrives
});
console.log("Step 2"); // runs immediately, before data is received
// Output: Step 1 → Step 2 → Data received

// When to use what:
// Synchronous: calculations, in-memory operations, config reading
// Asynchronous: API calls, file I/O, database queries, timers

// Common async patterns in modern JS:
// 1. Callbacks (old style — axios, event handlers)
// 2. Promises (fetch API, modern libraries)
// 3. async/await (for clean, readable async code)

// Practical: loading user dashboard
async function initDashboard() {
    showSkeleton(); // immediately show loading state

    const [user, notifications] = await Promise.all([
        api.get('/me'),
        api.get('/notifications')
    ]);

    hideSkeleton();
    renderUser(user);
    renderNotifications(notifications);
}
```

---

### Q6. How do you make multiple API calls efficiently?

```javascript
// Sequential: each call waits for previous (slow)
async function slow() {
    const user     = await fetchUser(1);    // 300ms
    const products = await fetchProducts(); // 300ms
    const cart     = await fetchCart(1);    // 300ms
    // Total: ~900ms
}

// Parallel: all start simultaneously (fast)
async function fast() {
    const [user, products, cart] = await Promise.all([
        fetchUser(1),
        fetchProducts(),
        fetchCart(1)
    ]);
    // Total: ~300ms (longest single request)
}

// When sequential is needed: when one depends on another
async function sequential() {
    const user = await fetchUser(1);           // need user first
    const orders = await fetchOrders(user.id); // needs user.id
}

// Partial success: load what's available even if some fail
async function loadDashboard() {
    const [userResult, productsResult, cartResult] = await Promise.allSettled([
        fetchUser(1),
        fetchProducts(),
        fetchCart(1)
    ]);

    return {
        user:     userResult.status === 'fulfilled'     ? userResult.value     : null,
        products: productsResult.status === 'fulfilled' ? productsResult.value : [],
        cart:     cartResult.status === 'fulfilled'     ? cartResult.value     : { items: [] }
    };
}

// Batch requests: collect IDs and fetch in one call
const pendingRequests = new Map(); // id → [resolve, reject][]
let batchTimer;

function getUserBatched(id) {
    return new Promise((resolve, reject) => {
        if (!pendingRequests.has(id)) pendingRequests.set(id, []);
        pendingRequests.get(id).push({ resolve, reject });

        clearTimeout(batchTimer);
        batchTimer = setTimeout(async () => {
            const ids = [...pendingRequests.keys()];
            const users = await fetchUsersInBatch(ids); // single API call!
            users.forEach(user => {
                pendingRequests.get(user.id)?.forEach(({ resolve }) => resolve(user));
            });
            pendingRequests.clear();
        }, 10); // batch requests in 10ms window
    });
}
```

---

### Q7. What are microtasks vs macrotasks?

```javascript
// Microtasks: high priority, processed immediately after current task
// Sources: Promise.then/catch/finally, queueMicrotask(), MutationObserver

// Macrotasks (tasks): lower priority, one processed per event loop cycle
// Sources: setTimeout, setInterval, setImmediate, I/O callbacks, UI events

// ⚠️ All microtasks drain before ANY macrotask runs
setTimeout(() => console.log("Macrotask"), 0);

Promise.resolve()
    .then(() => console.log("Microtask 1"))
    .then(() => console.log("Microtask 2"))
    .then(() => console.log("Microtask 3"));

console.log("Synchronous");

// Output: Synchronous → Microtask 1 → Microtask 2 → Microtask 3 → Macrotask

// Practical impact: async/await always defers to microtask queue
async function demo() {
    console.log("1 - start async");
    const result = await Promise.resolve("done"); // suspends here
    console.log("3 - after await:", result); // runs in microtask queue
}
demo();
console.log("2 - after calling async fn"); // runs synchronously
// Output: 1 → 2 → 3

// Why it matters:
// - Stacking microtasks can delay rendering (UI updates happen between macrotasks)
// - Promise.resolve().then() is guaranteed to run before setTimeout
// - Useful for deferring logic without delaying the critical render path

// Safe microtask scheduling
function nextTick(callback) {
    return Promise.resolve().then(callback);
}
nextTick(() => {
    // Runs after current synchronous code, before next timer
    // Safe to read DOM state updated this tick
});
```
