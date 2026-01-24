# JavaScript Interview Questions & Answers (51-100)

## 🔹 Intermediate (Continued)

### Error Handling & Debugging

#### Question 51: What are common JavaScript errors?

**Answer:**
JavaScript has several standard error types:
*   **`ReferenceError`**: Accessing a variable that hasn't been declared (or before initialization in TDZ).
*   **`TypeError`**: Operating on a value of the wrong type (e.g., `null.prop`, invoking non-function).
*   **`SyntaxError`**: Invalid code sent to parser.
*   **`RangeError`**: Numeric variable or parameter is outside its valid range.
*   **`URIError`**: `encodeURI()` passed invalid parameters.

**Example:**
```javascript
// ReferenceError
console.log(x); 

// TypeError
const num = 10;
num.toUpperCase(); 
```

---

#### Question 52: How to handle errors globally?

**Answer:**
*   **Browser (Synchronous/DOM):** `window.onerror` or `window.addEventListener('error')`.
*   **Unhandled Promises:** `window.addEventListener('unhandledrejection')`.
*   **Node.js:** `process.on('uncaughtException')` and `process.on('unhandledRejection')`.

**Example:**
```javascript
window.addEventListener('error', (event) => {
    console.log("Global error:", event.message);
});

window.addEventListener('unhandledrejection', (event) => {
    console.log("Unhandled Promise:", event.reason);
});
```

---

#### Question 53: How does `try...catch` with `finally` work?

**Answer:**
*   **`try`**: Executes code.
*   **`catch`**: Executes if exception is thrown in `try`.
*   **`finally`**: *Always* executes after try/catch, regardless of error or return statements. Ideal for cleanup (closing connections, hiding loaders).

**Example:**
```javascript
function demo() {
    try {
        console.log("Try");
        throw new Error("Oops");
    } catch (e) {
        console.log("Catch");
        return; // Finally still runs!
    } finally {
        console.log("Finally");
    }
}
demo();
// Output: Try -> Catch -> Finally
```

---

#### Question 54: What are custom errors?

**Answer:**
Creating user-defined error classes by extending the built-in `Error` class. Useful for specific application logic errors.

**Example:**
```javascript
class ValidationError extends Error {
    constructor(message) {
        super(message);
        this.name = "ValidationError";
    }
}

try {
    throw new ValidationError("Invalid input");
} catch (e) {
    if (e instanceof ValidationError) {
        console.log("Validation failed:", e.message);
    }
}
```

---

### Data Structures in JS

#### Question 55: What is a Set and how is it different from an array?

**Answer:**
A `Set` is a collection of **unique** values.
*   **Differences:**
    *   No duplicates allowed.
    *   No index-based access (`set[0]` doesn't work).
    *   Faster lookups `has()` (O(1)) compared to Array `includes()` (O(N)).

**Example:**
```javascript
const set = new Set([1, 2, 2, 3]);
set.add(4);
console.log(set.has(2)); // true
console.log(set.size);   // 4
```

---

#### Question 56: What is a Map?

**Answer:**
A `Map` is a collection of key-value pairs.
*   **Key types:** Allows any type (Objects, Functions) as keys, unlike Objects (Strings/Symbols only).
*   **Order:** Remembers insertion order.
*   **Size:** Has `.size` property.

**Example:**
```javascript
const map = new Map();
const keyObj = { id: 1 };

map.set(keyObj, "User Data");
console.log(map.get(keyObj)); // "User Data"
```

---

#### Question 57: How to implement a stack/queue using JavaScript?

**Answer:**
Using Arrays.
*   **Stack (LIFO):** `push()` to add, `pop()` to remove.
*   **Queue (FIFO):** `push()` to add, `shift()` to remove.

**Example:**
```javascript
// Stack
const stack = [];
stack.push(1);
stack.pop(); // 1

// Queue
const queue = [];
queue.push(1);
queue.shift(); // 1
```

---

#### Question 58: What are WeakMap and WeakSet?

**Answer:**
Collections that hold **weak references** to objects.
*   **Garbage Collection:** If the key (object) is solely referenced by the WeakMap/Set, it can be garbage collected.
*   **Constraints:** Keys must be objects. Not iterable (no `forEach` or `size`).

**Usage:** DOM Node metadata, private data storage.

---

### Date & Time

#### Question 59: How do you format a date in JavaScript?

**Answer:**
*   **`toLocaleDateString()`**: Localized format.
*   **`Intl.DateTimeFormat`**: High-performance, customizable internationalization.
*   **Libraries:** date-fns, Day.js (Moment.js is legacy).

**Example:**
```javascript
const date = new Date();
console.log(date.toLocaleDateString('en-US')); // "1/24/2026"

const formatter = new Intl.DateTimeFormat('en-GB', { 
    year: 'numeric', month: 'long', day: 'numeric' 
});
console.log(formatter.format(date)); // "24 January 2026"
```

---

#### Question 60: How to get the difference between two dates?

**Answer:**
Subtracting two Date objects returns the difference in milliseconds.

**Example:**
```javascript
const date1 = new Date('2026-01-01');
const date2 = new Date('2026-01-05');

const diffTime = Math.abs(date2 - date1);
const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24)); 
console.log(diffDays); // 4
```

---

### Regular Expressions

#### Question 61: How do regular expressions work in JavaScript?

**Answer:**
Patterns used to match character combinations in strings.
*   **Literal:** `/pattern/flags`
*   **Constructor:** `new RegExp('pattern', 'flags')`
*   **Methods:** `test()`, `exec()` (RegExp); `match()`, `replace()`, `search()` (String).

**Example:**
```javascript
const regex = /hello/i; // Case insensitive
console.log(regex.test("Hello World")); // true
```

---

#### Question 62: How to validate an email using regex?

**Answer:**
A basic pattern checks for `chars @ chars . chars`.

**Example:**
```javascript
function validateEmail(email) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(email);
}
console.log(validateEmail("test@example.com")); // true
```

---

### Miscellaneous

#### Question 63: What is NaN?

**Answer:**
"Not-a-Number". A special numeric value indicating an invalid operation (e.g., `Math.sqrt(-1)`, `"abc" * 2`).
*   **`typeof NaN`** is `"number"`.
*   **`NaN === NaN`** is `false`. Use `Number.isNaN()` to check.

---

#### Question 64: What is the difference between `parseInt()` and `Number()`?

**Answer:**
*   **`parseInt()`**: Parses a string character-by-character until it hits a non-digit. Good for "10px".
*   **`Number()`**: Strict conversion. Returns `NaN` if the string contains *any* invalid characters (except whitespace).

**Example:**
```javascript
console.log(parseInt("10px")); // 10
console.log(Number("10px"));   // NaN
```

---

#### Question 65: How do you deep clone an object?

**Answer:**
See Question 17. Use `structuredClone()` (Modern) or `JSON.parse(JSON.stringify())` (Basic) or recursion/libraries.

---

## 🔹 Advanced (71–100)

### Event Delegation & Propagation

#### Question 71: How does event bubbling/capturing work?

**Answer:**
Event propagation has three phases:
1.  **Capturing:** Root -> Target. (Traverse down).
2.  **Target:** Event reaches element.
3.  **Bubbling:** Target -> Root. (Traverse up - Default).

**Example:**
```javascript
// Capturing listener
el.addEventListener('click', fn, { capture: true });
```

---

#### Question 72: What is event delegation?

**Answer:**
A technique of adding a single event listener to a parent element instead of multiple listeners to child elements. It leverages **event bubbling** to detect which child was clicked via `event.target`.

**Benefits:** Memory efficiency, handles dynamic elements.

**Example:**
```javascript
document.querySelector('#list').addEventListener('click', (e) => {
    if (e.target.matches('li')) {
        console.log("List item clicked:", e.target.textContent);
    }
});
```

---

#### Question 73: What are synthetic events in React (JS background)?

**Answer:**
React wraps native browser events in a cross-browser wrapper called `SyntheticEvent`. This unifies behavior across browsers and allows React to pool events for performance.

---

### Memory & Performance

#### Question 74: How does garbage collection work in JS?

**Answer:**
JS engines use automatic GC.
*   **Reachability:** Roots (Global vars, stack) are assumed alive.
*   **Mark-and-Sweep:** The engine traverses from roots to mark accessible objects. Unmarked objects are considered garbage and memory is reclaimed.

---

#### Question 75: What are memory leaks and how to avoid them?

**Answer:**
When memory is allocated but not freed.
**Common Causes:**
1.  **Global variables:** Accidental `window.x`.
2.  **Forgotten intervals:** `setInterval` running forever.
3.  **Closures:** Holding huge scopes.
4.  **Detached DOM elements:** JS ref exists, but removed from DOM.

---

#### Question 76: How to profile JavaScript performance?

**Answer:**
*   **DevTools:** Performance Tab (Record runtime), Memory Tab (Heap snapshots).
*   **Console:** `console.time('label')` / `console.timeEnd('label')`.
*   **Performance API:** `performance.now()`.

---

### Browser APIs

#### Question 77: What is the Fetch API?

**Answer:**
A modern interface for making HTTP requests (AJAX). It uses Promises and is cleaner than `XMLHttpRequest`.

**Example:**
```javascript
fetch('https://api.example.com/data')
    .then(res => res.json())
    .then(data => console.log(data))
    .catch(err => console.error(err));
```

---

#### Question 78: How does `localStorage` differ from `sessionStorage`?

**Answer:**
Both store key-value strings (5MB limit).
*   **`localStorage`**: Persistent. Data remains until explicitly deleted.
*   **`sessionStorage`**: Temporary. Data persists only for the **window session** (cleared when tab closes).

---

#### Question 79: How does the Clipboard API work?

**Answer:**
Modern async API to read/write clipboard. Requires secure context (HTTPS) and permissions.

**Example:**
```javascript
navigator.clipboard.writeText("Copy me!")
    .then(() => console.log("Copied"));
```

---

#### Question 80: What is `requestAnimationFrame`?

**Answer:**
A method to perform animations efficiently. It schedules a callback to run before the next browser repaint (usually 60fps). It pauses in background tabs to save battery.

**Example:**
```javascript
function animate() {
    element.style.left = (parseInt(element.style.left) + 1) + 'px';
    requestAnimationFrame(animate);
}
requestAnimationFrame(animate);
```

---

### Security

#### Question 81: What is XSS and how do you prevent it in JS?

**Answer:**
**Cross-Site Scripting (XSS):** Injecting malicious scripts into a webpage.
**Prevention:**
*   Escape user input.
*   Use `textContent` instead of `innerHTML`.
*   Content Security Policy (CSP).

---

#### Question 82: What is CSRF and how does JavaScript help mitigate it?

**Answer:**
**Cross-Site Request Forgery:** Attacker tricks a user's browser into making unwanted requests.
**Mitigation:**
*   Use **Anti-CSRF tokens** (JS reads token from meta/cookie and sends in header).
*   SameSite Cookies.

---

#### Question 83: How do you sanitize user input?

**Answer:**
Never trust input.
*   Sanitize HTML tags (DOMPurify library).
*   Encode output (HTML entities).
*   Validate types.

---

### Design Patterns in JS

#### Question 84: What is the module pattern?

**Answer:**
Uses closures (often via IIFE) to create private scope and return a public API object.

**Example:**
```javascript
const Module = (function() {
    let private = "Hidden";
    return {
        public: () => private
    };
})();
```

---

#### Question 85: How does the singleton pattern work in JS?

**Answer:**
Ensures a class has only one instance.
**Example:**
```javascript
const Singleton = {
    method() { ... }
};
// Object literals are singletons by default.
```

---

#### Question 86: What is the observer pattern?

**Answer:**
Pub/Sub model. An object (Subject) maintains a list of dependents (Observers) and notifies them of state changes. Used in DOM events (`addEventListener`).

---

#### Question 87: What is the factory pattern?

**Answer:**
A function that creates and returns objects, abstracting specific character instantiation.

**Example:**
```javascript
function createUser(role) {
    return { role, createdAt: Date.now() };
}
```

---

### Build Tools & Environment

#### Question 88: What is transpilation (e.g., Babel)?

**Answer:**
Converting newer source code (ES6+) into older syntax (ES5) to ensure compatibility with older browsers.

---

#### Question 89: What is tree shaking?

**Answer:**
A build optimization (Webpack/Rollup) that eliminates dead code (unused exports) from the final bundle to reduce file size. Relies on ES Modules (`import`/`export`).

---

#### Question 90: How does bundling work with Webpack?

**Answer:**
Webpack takes modules with dependencies (JS, CSS, Images) and generates static assets representing those modules. It builds a dependency graph and merges files into one or more bundles.

---

### Testing & Types

#### Question 91: What are unit tests in JavaScript?

**Answer:**
Testing individual units (functions/classes) in isolation.

#### Question 92: How do you mock dependencies?

**Answer:**
Replacing real modules with simulated versions (spies/stubs) to isolate the unit under test. (e.g., `jest.fn()`, `jest.mock()`).

#### Question 93: What are common testing libraries in JS?

**Answer:**
*   **Jest**: All-in-one runner/assertion/mocking.
*   **Mocha**: Runner.
*   **Chai**: Assertions.
*   **Cypress/Playwright**: E2E testing.

#### Question 94: What is TypeScript?

**Answer:**
A superset of JavaScript adding static typing. Can catch errors at compile time.

#### Question 95: Static vs Dynamic Typing?

**Answer:**
*   **Static (TS):** Types checked at compile time.
*   **Dynamic (JS):** Types checked at runtime. variables can hold any type.

---

### Misc/Tricky

#### Question 96: What is the output of `[] + []` or `{} + []`?

**Answer:**
*   `[] + []`: `""` (Empty string). Arrays convert to strings (`""`) and concatenate.
*   `{} + []`: `0` (interactive REPLs interpret `{}` as block) OR `"[object Object]"` (as expression).

---

#### Question 97: How does `this` behave in arrow functions?

**Answer:**
Arrow functions **do not** have their own `this`. They inherit `this` lexically from the scope in which they were defined. `call`/`apply`/`bind` **cannot** change it.

---

#### Question 98: What are tagged template literals?

**Answer:**
Functions that parse template literals.
**Example:**
```javascript
function tag(strings, ...values) {
    console.log(strings, values);
}
tag`Hello ${name}`;
```

---

#### Question 99: How do you implement your own `bind` function?

**Answer:**
```javascript
Function.prototype.myBind = function(context, ...args) {
    const fn = this;
    return function(...innerArgs) {
        return fn.apply(context, [...args, ...innerArgs]);
    }
}
```

---

#### Question 100: Explain `Promise.all`, `Promise.race`, and `Promise.any`.

**Answer:**
*   **`all(iterable)`**: Resolves when **ALL** resolve. Rejects if **ONE** rejects.
*   **`race(iterable)`**: Settles as soon as the **FIRST** promise settles (resolve or reject).
*   **`any(iterable)`**: Resolves as soon as **ONE** resolves. Rejects if **ALL** reject.
