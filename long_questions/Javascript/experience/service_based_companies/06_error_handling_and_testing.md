# 📘 06 — Error Handling, Testing & Code Quality
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- `try/catch/finally` and custom error classes
- Error types and error propagation
- Basic testing concepts (unit, integration)
- Code quality: naming, comments, pure functions
- Console methods
- Debugging techniques

---

## ❓ Most Asked Questions

### Q1. How does error handling work in JavaScript?

```javascript
// try/catch/finally: handle runtime errors without crashing

// Basic error handling
try {
    const data = JSON.parse(invalidString); // throws SyntaxError
    processData(data);
} catch (error) {
    console.error("Error type:", error.name);    // "SyntaxError"
    console.error("Message:", error.message);    // details
    console.error("Stack:", error.stack);        // stack trace
} finally {
    // ALWAYS runs — good for cleanup
    closeLoader();
}

// Error types in JavaScript:
// - Error: generic base error
// - SyntaxError: invalid JS syntax
// - TypeError: wrong type (calling non-function, accessing property of null)
// - ReferenceError: undefined variable
// - RangeError: value out of range (e.g., new Array(-1))
// - URIError: invalid encodeURIComponent input

// Specific error handling
try {
    doSomething();
} catch (err) {
    if (err instanceof TypeError) {
        console.log("Type error:", err.message);
    } else if (err instanceof RangeError) {
        console.log("Range error:", err.message);
    } else {
        throw err; // re-throw unknown errors
    }
}

// Custom error class
class ValidationError extends Error {
    constructor(field, message) {
        super(message);
        this.name = "ValidationError";
        this.field = field;
    }
}

class ApiError extends Error {
    constructor(status, message) {
        super(message);
        this.name = "ApiError";
        this.status = status;
    }
}

// Usage
function validateAge(age) {
    if (typeof age !== "number") throw new ValidationError("age", "Age must be a number");
    if (age < 0 || age > 150) throw new ValidationError("age", "Age out of range");
    return true;
}

try {
    validateAge("twenty");
} catch (err) {
    if (err instanceof ValidationError) {
        console.log(`Validation failed for field '${err.field}': ${err.message}`);
    }
}
```

---

### Q2. How do you handle async errors properly?

```javascript
// Promise error handling

// .catch() at end of chain
fetchUserData(userId)
    .then(user => processUser(user))
    .then(result => saveResult(result))
    .catch(err => {
        // Catches errors from ANY step in the chain
        logger.error("Pipeline failed:", err);
    });

// ⚠️ Unhandled promise rejection: catch ALL promises!
// Without .catch(), unhandled rejections cause warnings (Node) or errors (browsers)
fetch('/api/data'); // ❌ no .catch() — potential silent failure

// async/await with proper error handling
async function loadUserProfile(userId) {
    try {
        const user = await fetchUser(userId);

        if (!user) {
            throw new Error(`User ${userId} not found`);
        }

        const profile = await fetchProfile(user.id);
        return { user, profile };

    } catch (error) {
        if (error instanceof ApiError && error.status === 404) {
            return null; // user not found — return null, don't throw
        }
        // Unknown error — re-throw for caller to handle
        throw error;
    }
}

// Global error handlers (safety net — not a substitute for try/catch!)
// Browser
window.addEventListener("unhandledrejection", (event) => {
    console.error("Unhandled promise rejection:", event.reason);
    event.preventDefault(); // suppress console error (optional)
});

window.addEventListener("error", (event) => {
    console.error("Uncaught error:", event.error);
    reportToMonitoring(event.error);
});

// Node.js
process.on("unhandledRejection", (reason, promise) => {
    console.error("Unhandled Rejection:", reason);
});

process.on("uncaughtException", (error) => {
    console.error("Uncaught Exception:", error);
    process.exit(1); // gracefully restart via PM2/systemd
});
```

---

### Q3. What are the console debugging methods?

```javascript
// Essential console methods for service-based interviews

// Basic logging
console.log("Simple log");
console.info("Info message");
console.warn("Warning!");      // yellow in DevTools
console.error("Error!");       // red in DevTools, includes stack trace

// Object inspection
console.log(user);             // may show [Object]
console.dir(user);             // shows full object tree

// Grouped output
console.group("User Details");
console.log("Name:", user.name);
console.log("Age:", user.age);
console.groupEnd();

console.groupCollapsed("API Response"); // starts collapsed
console.log(response);
console.groupEnd();

// Performance measurement
console.time("dataProcessing");
processLargeDataset(data);
console.timeEnd("dataProcessing"); // "dataProcessing: 45.2ms"

// Count calls
function trackClicks() {
    console.count("Button clicked");
}
// Logs: "Button clicked: 1", "Button clicked: 2", ...

// Assert: logs error if condition is false
console.assert(array.length > 0, "Array should not be empty");

// Table: display array of objects as table
console.table([
    { name: "Alice", age: 25, role: "admin" },
    { name: "Bob",   age: 30, role: "viewer" }
]);

// Stack trace
function a() { b(); }
function b() { c(); }
function c() { console.trace("Trace from c"); }
a(); // shows call stack: c → b → a

// ⚠️ Remove console.log before production — use a proper logger
// Production-safe logging
const logger = {
    log:   (...args) => process.env.NODE_ENV !== "production" && console.log(...args),
    error: (...args) => console.error(...args), // always log errors
    warn:  (...args) => console.warn(...args)
};
```

---

### Q4. What are common JavaScript bugs and how to avoid them?

```javascript
// Bug 1: Mutating function arguments (pure function violation)
function addItem(cart, item) {
    cart.items.push(item); // ❌ mutates original array!
    return cart;
}

// ✅ Fix: return new object
function addItemSafe(cart, item) {
    return { ...cart, items: [...cart.items, item] };
}

// Bug 2: Floating-point arithmetic
0.1 + 0.2 === 0.3; // false! (0.30000000000000004)

// ✅ Fix: use toFixed or Math.round for display; compare with epsilon
Math.abs(0.1 + 0.2 - 0.3) < Number.EPSILON; // true
(0.1 + 0.2).toFixed(2) === "0.30"; // true
// For currency: use integers (store in cents)
const price = 1050; // ₹10.50 stored as paise

// Bug 3: Asynchronous bugs — stale closures
function makeCounter() {
    let count = 0;
    return () => {
        console.log(count);
        setTimeout(() => count++, 1000);
    };
}
// Multiple clicks before 1s: all log 0 (stale closure)

// Bug 4: Object reference vs value
const a = { x: 1 };
const b = a;          // same reference!

b.x = 99;
console.log(a.x); // 99 — unexpected mutation!

// ✅ Fix: always clone
const b2 = { ...a };
b2.x = 99;
console.log(a.x); // 1 — safe

// Bug 5: parseInt radix
parseInt("08");      // in old browsers: 0 (octal)! 
parseInt("08", 10);  // ✅ 8 — always specify radix

// Bug 6: typeof null
typeof null === "object"; // true — JS bug!
// ✅ Check for null explicitly
function isObject(val) {
    return val !== null && typeof val === "object";
}

// Bug 7: Array sorting default (lexicographic)
[10, 1, 100, 2].sort();          // [1, 10, 100, 2] — string sort!
[10, 1, 100, 2].sort((a, b) => a - b); // ✅ [1, 2, 10, 100] — numeric
```

---

### Q5. Explain basic testing concepts with Jest examples.

```javascript
// Testing: verify code behaves as expected
// Test types:
// Unit tests: test one function in isolation
// Integration tests: test how multiple pieces work together
// E2E tests: simulate full user journeys

// Jest: popular testing framework

// Unit test example (utils.test.js)
const { formatCurrency, calculateDiscount } = require('./utils');

// describe: groups related tests
describe('formatCurrency', () => {
    // test / it: individual test case
    test('formats whole number correctly', () => {
        expect(formatCurrency(1050)).toBe('₹1,050');
    });

    test('formats decimal correctly', () => {
        expect(formatCurrency(99.5)).toBe('₹99.50');
    });

    test('handles zero', () => {
        expect(formatCurrency(0)).toBe('₹0');
    });
});

describe('calculateDiscount', () => {
    test('returns discounted price', () => {
        expect(calculateDiscount(1000, 20)).toBe(800); // 20% off
    });

    test('throws for invalid discount', () => {
        expect(() => calculateDiscount(1000, 110)).toThrow('Discount cannot exceed 100%');
    });
});

// Async test
test('fetches user by ID', async () => {
    // Mock the fetch call
    global.fetch = jest.fn().mockResolvedValueOnce({
        ok: true,
        json: async () => ({ id: 1, name: 'Alice' })
    });

    const user = await getUserById(1);

    expect(user.name).toBe('Alice');
    expect(fetch).toHaveBeenCalledWith('/api/users/1'); // verify call
    expect(fetch).toHaveBeenCalledTimes(1);             // called once
});

// Common matchers:
expect(value).toBe(exact);          // ===
expect(value).toEqual(deepEqual);   // deep equality
expect(value).toBeNull();
expect(value).toBeDefined();
expect(value).toBeTruthy();
expect(arr).toContain(item);
expect(arr).toHaveLength(5);
expect(str).toMatch(/pattern/);
```

---

### Q6. What are pure functions and side effects?

```javascript
// Pure function:
// 1. Same inputs → always same output (deterministic)
// 2. No side effects (doesn't modify external state)

// ✅ Pure
function add(a, b) { return a + b; }
function formatName(first, last) { return `${first} ${last}`; }
function getTotal(items) { return items.reduce((sum, item) => sum + item.price, 0); }

// ❌ Impure (has side effects)
let tax = 0.18;
function priceWithTax(amount) {
    return amount * (1 + tax); // depends on external 'tax' variable
}

let count = 0;
function increment() { return ++count; } // modifies external state

function fetchUser(id) { return fetch(`/api/users/${id}`); } // I/O side effect

// Side effects (not always bad, but should be isolated):
// - DOM manipulation
// - Fetch/HTTP calls
// - localStorage operations
// - console.log
// - Modifying variables outside function scope
// - Math.random(), Date.now() (non-deterministic)

// ✅ Best practice: keep business logic pure, isolate side effects
// Core logic — pure, testable
function calculateOrderTotal(items, taxRate) {
    const subtotal = items.reduce((sum, item) => sum + item.price * item.qty, 0);
    return subtotal * (1 + taxRate);
}

// Coordinating function — handles side effects
async function processOrder(orderId) {
    const order = await fetchOrder(orderId);       // side effect: fetch
    const total = calculateOrderTotal(order.items, 0.18); // pure!
    await saveOrder({ ...order, total });           // side effect: save
    sendEmail(order.customerEmail, total);          // side effect: email
}
// calculateOrderTotal can be unit-tested without mocking anything
```

---

### Q7. How do you write clean, maintainable JavaScript code?

```javascript
// Clean Code Principles for JavaScript:

// 1. Meaningful names
// ❌
const d = new Date();
const arr = users.filter(u => u.a > 18);

// ✅
const currentDate = new Date();
const adultUsers = users.filter(user => user.age > 18);

// 2. Functions should do ONE thing (SRP)
// ❌ function doing too much
async function processUserRegistration(data) {
    validateEmail(data.email);
    const hashedPwd = hashPassword(data.password);
    const user = await db.users.create({ ...data, password: hashedPwd });
    await sendWelcomeEmail(user.email);
    await createDefaultSettings(user.id);
    return user;
}

// ✅ Break into focused functions
async function registerUser(userData) {
    const validated  = validateRegistrationData(userData);
    const user       = await createUserAccount(validated);
    await onboardNewUser(user);
    return user;
}

// 3. Avoid deep nesting — early returns
// ❌
function processOrder(order) {
    if (order) {
        if (order.items && order.items.length > 0) {
            if (order.customer) {
                // main logic
            }
        }
    }
}

// ✅ Guard clauses
function processOrder(order) {
    if (!order) return null;
    if (!order.items?.length) throw new Error("Empty order");
    if (!order.customer) throw new Error("No customer");
    // main logic here
}

// 4. Consistent error handling
// 5. Avoid magic numbers/strings
const MAX_RETRIES   = 3;
const TIMEOUT_MS    = 5000;
const STATUS_ACTIVE = "active";

// 6. Use const by default, let when needed, avoid var
// 7. Prefer functional methods: map, filter, reduce over imperative loops
// 8. Destructure for cleaner code
```
