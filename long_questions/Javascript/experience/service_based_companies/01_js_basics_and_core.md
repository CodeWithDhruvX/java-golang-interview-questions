# 📘 01 — JavaScript Basics & Core Concepts
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Data types, `var`/`let`/`const`, hoisting
- Type coercion, `==` vs `===`
- Closures and scope
- Arrow functions and `this`
- Array methods: `map`, `filter`, `reduce`, `forEach`
- Object methods: `Object.keys`, `Object.entries`, spread operator
- Template literals, destructuring

---

## ❓ Most Asked Questions

### Q1. What are the different data types in JavaScript?

```javascript
// JavaScript has 8 data types:

// Primitive (immutable, stored by value):
let str    = "Hello";        // String
let num    = 42;             // Number
let big    = 9007199254740992n; // BigInt
let bool   = true;           // Boolean
let undef;                   // Undefined
let empty  = null;           // Null
let sym    = Symbol("id");   // Symbol

// Non-primitive (mutable, stored by reference):
let obj    = { name: "Alice" }; // Object
let arr    = [1, 2, 3];         // Array (type of object)
let fn     = function() {};     // Function (type of object)

// Checking type:
typeof "text"       // "string"
typeof 42           // "number"
typeof undefined    // "undefined"
typeof null         // "object"  ← famous JS bug (null is not an object)
typeof []           // "object"
typeof function(){} // "function"

// Better null/array check:
Array.isArray([]);                           // true
Object.prototype.toString.call(null);        // "[object Null]"
Object.prototype.toString.call([]);          // "[object Array]"
```

---

### Q2. Difference between `var`, `let`, and `const`.

```javascript
// var: function-scoped, hoisted and initialized to undefined, re-declarable
function demoVar() {
    console.log(x); // undefined (hoisted)
    var x = 10;
    if (true) {
        var x = 20; // same 'x' — function scoped
    }
    console.log(x); // 20
}

// let: block-scoped, hoisted but in TDZ, not re-declarable
function demoLet() {
    // console.log(y); // ReferenceError (TDZ)
    let y = 10;
    if (true) {
        let y = 20; // different 'y' — block scoped
        console.log(y); // 20
    }
    console.log(y); // 10
}

// const: block-scoped, must be initialized, cannot be reassigned
const PI = 3.14159;
// PI = 3; // TypeError

// ⚠️ const with objects: reference is constant, not the content!
const user = { name: "Alice" };
user.name = "Bob";  // ✅ mutating the object is fine
// user = {};       // ❌ reassigning the variable is not

// Best practice: use const by default, let when you need to reassign, avoid var
```

---

### Q3. What are closures? Give a practical example.

```javascript
// Closure: a function that remembers the variables from its outer scope
// even after the outer function has finished executing.

// Example: Counter
function createCounter() {
    let count = 0; // this variable is "closed over"

    return {
        increment() { return ++count; },
        decrement() { return --count; },
        getCount()  { return count; }
    };
}

const counter = createCounter();
counter.increment(); // 1
counter.increment(); // 2
counter.decrement(); // 1
// count is private — cannot be accessed from outside

// Practical use: default parameters with configuration
function createApiCaller(baseUrl) {
    return function(endpoint) {
        return fetch(baseUrl + endpoint); // baseUrl is closed over
    };
}

const callApi = createApiCaller("https://api.example.com");
callApi("/users");    // fetches https://api.example.com/users
callApi("/products"); // fetches https://api.example.com/products

// ⚠️ Common closure pitfall with loops:
for (var i = 1; i <= 3; i++) {
    setTimeout(() => console.log(i), i * 1000); // Prints 4, 4, 4!
    // Fix: use let
}
for (let i = 1; i <= 3; i++) {
    setTimeout(() => console.log(i), i * 1000); // Prints 1, 2, 3 ✅
}
```

---

### Q4. Explain arrow functions and how `this` differs.

```javascript
// Arrow functions: concise syntax, no own 'this'

// Regular function
function greet(name) {
    return "Hello, " + name;
}

// Arrow function (equivalent)
const greetArrow = (name) => "Hello, " + name;
const double = n => n * 2;          // single param: no parentheses needed
const square = n => ({ value: n * n }); // returning object: wrap in ()

// KEY difference: 'this' binding
const person = {
    name: "Alice",
    
    // ❌ Arrow function: 'this' inherits from outer scope (module/window — NOT person)
    greetArrow: () => {
        console.log(`Hi, I'm ${this?.name}`); // undefined!
    },

    // ✅ Regular function: 'this' is the object that called the method
    greetRegular() {
        console.log(`Hi, I'm ${this.name}`); // "Alice"
    },

    // ✅ Arrow inside regular method: captures 'this' correctly
    startTimer() {
        setTimeout(() => {
            console.log(`${this.name} is done`); // "Alice" — arrow captures this
        }, 1000);
    }
};

// Arrow functions also lack:
// - arguments object
// - 'new' (cannot be a constructor)
// - prototype property
```

---

### Q5. Explain `map`, `filter`, `reduce` with examples.

```javascript
const products = [
    { id: 1, name: "Laptop",  price: 50000, category: "electronics" },
    { id: 2, name: "T-Shirt", price: 500,   category: "clothing" },
    { id: 3, name: "Phone",   price: 30000, category: "electronics" },
    { id: 4, name: "Jeans",   price: 1500,  category: "clothing" },
];

// map: transform each element → new array (same length)
const names = products.map(p => p.name);
// ["Laptop", "T-Shirt", "Phone", "Jeans"]

const withTax = products.map(p => ({ ...p, priceWithTax: p.price * 1.18 }));
// Each product now has priceWithTax field

// filter: keep elements passing a condition → new array (≤ same length)
const electronics = products.filter(p => p.category === "electronics");
// [Laptop, Phone]

const affordable = products.filter(p => p.price < 10000);
// [T-Shirt, Jeans]

// reduce: accumulate into single value
const totalValue = products.reduce((sum, p) => sum + p.price, 0);
// 82000

// Group by category (reduce into object)
const grouped = products.reduce((acc, p) => {
    acc[p.category] = acc[p.category] ?? [];
    acc[p.category].push(p);
    return acc;
}, {});
// { electronics: [Laptop, Phone], clothing: [T-Shirt, Jeans] }

// Chaining: get total of electronics priced > 20000
const electronicsTotal = products
    .filter(p => p.category === "electronics" && p.price > 20000)
    .reduce((sum, p) => sum + p.price, 0);
// 80000
```

---

### Q6. What is the spread operator and rest parameters?

```javascript
// Spread (...): expand iterable into individual elements
// Rest (...): collect multiple elements into an array

// Spread: merge arrays
const arr1 = [1, 2, 3];
const arr2 = [4, 5, 6];
const merged = [...arr1, ...arr2]; // [1, 2, 3, 4, 5, 6]

// Spread: copy array (shallow)
const original = [1, 2, 3];
const copy = [...original]; // new array, not same reference

// Spread: merge objects
const defaults = { theme: "light", lang: "en", fontSize: 14 };
const userPrefs = { lang: "hi", fontSize: 16 };
const settings = { ...defaults, ...userPrefs };
// { theme: "light", lang: "hi", fontSize: 16 } — userPrefs overrides

// Spread: function call
const nums = [3, 1, 4, 1, 5, 9];
Math.max(...nums); // equivalent to Math.max(3, 1, 4, 1, 5, 9) → 9

// Rest parameters: collect remaining args into array
function sum(first, second, ...rest) {
    console.log(first, second); // 1, 2
    console.log(rest);          // [3, 4, 5]
    return [first, second, ...rest].reduce((a, b) => a + b, 0);
}
sum(1, 2, 3, 4, 5); // 15

// Destructuring with rest
const [head, ...tail] = [1, 2, 3, 4, 5];
// head = 1, tail = [2, 3, 4, 5]

const { name, ...rest } = { name: "Alice", age: 25, city: "Delhi" };
// name = "Alice", rest = { age: 25, city: "Delhi" }
```

---

### Q7. What is event bubbling and capturing?

```javascript
// Event propagation has 3 phases:
// 1. Capture phase: event travels from root → target
// 2. Target phase: event reaches target element
// 3. Bubble phase: event travels back from target → root

// HTML structure: body > div#parent > button#child

document.getElementById("parent").addEventListener("click", () => {
    console.log("Parent clicked");
}, false); // false = bubble phase (default)

document.getElementById("child").addEventListener("click", (e) => {
    console.log("Child clicked");
    e.stopPropagation(); // stops event from bubbling to parent
}, false);

// Click on button → "Child clicked" (parent handler doesn't fire)

// event.target: element that triggered the event
// event.currentTarget: element the listener is attached to

// Capture: { capture: true } — fires before bubble
document.getElementById("parent").addEventListener("click", () => {
    console.log("Parent (capture)");
}, { capture: true }); // fires BEFORE child's handler

// Event delegation: listen on parent, handle child events
// More efficient than adding listeners to each child
document.getElementById("list").addEventListener("click", (e) => {
    if (e.target.matches("li")) {
        console.log("List item clicked:", e.target.textContent);
    }
    if (e.target.closest("button.delete")) {
        removeItem(e.target.closest("li").dataset.id);
    }
});
```

---

### Q8. What is the difference between `null`, `undefined`, and `NaN`?

```javascript
// undefined: variable declared but not assigned; missing property; function without return
let x;               // undefined
const obj = {};
console.log(obj.name); // undefined — property doesn't exist

function noReturn() {} 
noReturn(); // undefined

// null: intentional absence of value — explicitly set by developer
let user = null; // "no user currently"

// NaN: "Not a Number" — result of invalid numeric operation
parseInt("abc");     // NaN
0 / 0;              // NaN
Math.sqrt(-1);      // NaN

// Tricky: NaN is not equal to itself!
NaN === NaN;         // false
NaN == NaN;          // false
Number.isNaN(NaN);   // true (safer than global isNaN())
isNaN("hello");      // true (coerces "hello" → NaN first) — unreliable!
Number.isNaN("hello"); // false (no coercion — more reliable)

// null checks
null == undefined;   // true  (loose equality)
null === undefined;  // false (strict equality)
null == false;       // false (null only loosely equals null and undefined)

// Practical null/undefined check
const value = null ?? "default"; // "default" — nullish coalescing
const safe  = user?.name;        // undefined (optional chaining, no error)
```

---

### Q9. Explain Promises and async/await with error handling.

```javascript
// Promise: represents eventual completion or failure of async operation

// Creating a promise
function delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function fetchUser(id) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            if (id > 0) resolve({ id, name: "Alice" });
            else reject(new Error("Invalid ID"));
        }, 500);
    });
}

// Promise chaining
fetchUser(1)
    .then(user => {
        console.log(user.name);
        return fetchOrders(user.id); // can return another promise
    })
    .then(orders => console.log(orders))
    .catch(err => console.error("Error:", err.message))
    .finally(() => console.log("Done")); // always runs

// async/await (cleaner syntax)
async function loadUserData(userId) {
    try {
        const user   = await fetchUser(userId);
        const orders = await fetchOrders(user.id);
        return { user, orders };
    } catch (err) {
        console.error("Failed:", err.message);
        return null; // graceful fallback
    } finally {
        setLoading(false); // always runs
    }
}

// Parallel requests (faster)
async function loadDashboard(userId) {
    const [user, products, notifications] = await Promise.all([
        fetchUser(userId),
        fetchProducts(),
        fetchNotifications(userId)
    ]);
    return { user, products, notifications };
}
```

---

### Q10. How does `this` work in JavaScript?

```javascript
// 'this' depends on HOW a function is called, not where it's defined

// 1. Global context
console.log(this); // window (browser) / global (Node.js)

// 2. Object method
const obj = {
    name: "Alice",
    getName() { return this.name; } // this = obj
};
obj.getName(); // "Alice"

// 3. Lost binding (common interview trick)
const fn = obj.getName;
fn(); // undefined — 'this' is global, not obj

// 4. Fix: bind
const boundFn = obj.getName.bind(obj);
boundFn(); // "Alice"

// 5. call / apply
function introduce(greeting) {
    return `${greeting}, I'm ${this.name}`;
}
introduce.call({ name: "Bob" }, "Hi");   // "Hi, I'm Bob"
introduce.apply({ name: "Bob" }, ["Hi"]); // "Hi, I'm Bob"

// 6. Constructor
function Person(name) {
    this.name = name; // 'this' = newly created object
}
const alice = new Person("Alice");
alice.name; // "Alice"

// 7. Arrow function: inherits 'this' from enclosing scope
class Timer {
    constructor() { this.count = 0; }
    start() {
        setInterval(() => this.count++, 1000); // arrow: 'this' = Timer instance
    }
}
```
