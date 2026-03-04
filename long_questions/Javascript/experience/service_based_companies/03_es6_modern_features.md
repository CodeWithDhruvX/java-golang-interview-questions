# 📘 03 — ES6+ Features & Modern JavaScript
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Destructuring (arrays and objects)
- Template literals
- `class` syntax
- `Promise`, `async/await`
- Modules (import/export)
- Optional chaining, nullish coalescing
- `Set`, `Map`, `WeakMap`

---

## ❓ Most Asked Questions

### Q1. What is destructuring? Give practical examples.

```javascript
// Destructuring: unpack values from arrays/objects into variables

// Array Destructuring
const [first, second, , fourth] = [10, 20, 30, 40];
// first=10, second=20, fourth=40 (30 is skipped)

// With default values
const [a = 1, b = 2] = [100]; // a=100, b=2 (default used)

// Swapping variables
let x = 1, y = 2;
[x, y] = [y, x]; // x=2, y=1 (no temp variable needed!)

// Rest in array
const [head, ...tail] = [1, 2, 3, 4, 5];
// head=1, tail=[2,3,4,5]

// Object Destructuring
const user = { id: 1, name: "Alice", age: 25, city: "Delhi" };
const { name, age } = user;

// Renaming
const { name: userName, age: userAge } = user; // rename: name → userName

// Default values
const { role = "viewer", name: displayName } = user;
// role="viewer" (not in user object), displayName="Alice"

// Nested destructuring
const order = {
    id: "ORD-001",
    customer: { name: "Bob", address: { city: "Mumbai" } }
};
const { customer: { name: customerName, address: { city } } } = order;
// customerName="Bob", city="Mumbai"

// Function parameter destructuring
function displayUser({ name, age, role = "user" }) {
    return `${name} (${age}) - ${role}`;
}
displayUser({ name: "Alice", age: 25 }); // "Alice (25) - user"

// API response handling
async function fetchUserProfile(id) {
    const { data: { user, permissions }, status } = await api.get(`/users/${id}`);
    return { user, permissions, ok: status === 200 };
}
```

---

### Q2. Explain ES6 class syntax with inheritance.

```javascript
// ES6 class: cleaner syntax for prototype-based OOP

class Animal {
    // Class fields (ES2022)
    #sound = "..."; // private field

    constructor(name, species) {
        this.name    = name;
        this.species = species;
    }

    // Method (on prototype)
    speak() {
        return `${this.name} says ${this.#sound}`;
    }

    // Getter and setter
    get label() { return `[${this.species}] ${this.name}`; }
    set label(val) { [this.species, this.name] = val.split(' '); }

    // Static method (on class, not instances)
    static create(type, name) {
        const sounds = { dog: "Woof", cat: "Meow", cow: "Moo" };
        const animal = new Animal(name, type);
        animal.#sound = sounds[type] ?? "...";
        return animal;
    }
}

// Inheritance
class Dog extends Animal {
    #tricks = [];

    constructor(name) {
        super(name, "dog"); // call parent constructor
        this.breed = "";
    }

    learn(trick) {
        this.#tricks.push(trick);
        return this; // method chaining
    }

    showTricks() {
        return `${this.name} knows: ${this.#tricks.join(', ') || 'nothing yet'}`;
    }

    // Override parent method
    speak() {
        return `${super.speak()} (really loudly!)`;
    }
}

const rex = new Dog("Rex");
rex.learn("sit").learn("shake").learn("roll over");
rex.speak();       // "Rex says Woof (really loudly!)"
rex.showTricks();  // "Rex knows: sit, shake, roll over"
rex.label;         // "[dog] Rex"
rex instanceof Dog;    // true
rex instanceof Animal; // true (prototype chain)
```

---

### Q3. Explain `Set` and `Map` with use cases.

```javascript
// Set: collection of UNIQUE values (any type)
const uniqueIds = new Set([1, 2, 3, 2, 1, 3]);
uniqueIds.size;          // 3 (duplicates removed)
uniqueIds.has(2);        // true
uniqueIds.add(4);
uniqueIds.delete(1);
uniqueIds.forEach(id => console.log(id));
[...uniqueIds];          // [2, 3, 4] — convert to array

// ✅ Use case 1: remove duplicates
const words = ["apple", "banana", "apple", "cherry", "banana"];
const unique = [...new Set(words)]; // ["apple", "banana", "cherry"]

// ✅ Use case 2: tag system (intersection, union, difference)
const techTags = new Set(["js", "python", "java"]);
const jobTags  = new Set(["js", "react", "java"]);

// Intersection: common tags
const common = new Set([...techTags].filter(t => jobTags.has(t)));
// Set {"js", "java"}

// Union: all tags
const all = new Set([...techTags, ...jobTags]);

// Map: key-value pairs where keys can be ANY type (not just strings)
const userRoles = new Map();
userRoles.set("alice@example.com", "admin");
userRoles.set("bob@example.com", "viewer");

userRoles.get("alice@example.com"); // "admin"
userRoles.has("bob@example.com");   // true
userRoles.size;                      // 2

// Map preserves insertion order, iteratable
for (const [email, role] of userRoles) {
    console.log(`${email}: ${role}`);
}

// ✅ Use case: counting word frequency
function wordFrequency(text) {
    const freq = new Map();
    text.split(/\s+/).forEach(word => {
        freq.set(word, (freq.get(word) ?? 0) + 1);
    });
    return freq;
}

// Map vs Object:
// Map: any key type, ordered, size property, iterable, no inherited keys
// Object: string/symbol keys only, prototype chain risk, JSON-serializable
```

---

### Q4. How do optional chaining (`?.`) and nullish coalescing (`??`) work?

```javascript
// Optional chaining: access nested properties safely without null checks
const user = {
    profile: {
        address: {
            city: "Mumbai"
        }
    }
};

// Old way (verbose)
const city1 = user && user.profile && user.profile.address && user.profile.address.city;

// ✅ Optional chaining
const city2 = user?.profile?.address?.city;     // "Mumbai"
const zip   = user?.profile?.address?.zipCode;  // undefined (no error)

// Works with methods and array access
const firstOrder = user?.orders?.[0];            // undefined if no orders
const total = user?.getCartTotal?.();            // undefined if method doesn't exist

// Nullish coalescing: use right side ONLY if left is null or undefined
// Unlike ||, it doesn't treat 0, "", false as falsy
const score = 0;

const result1 = score || 10;   // 10 — wrong! 0 is falsy
const result2 = score ?? 10;   // 0 — correct! 0 is not null/undefined

// Combining
const config = null;
const timeout = config?.timeout ?? 5000;    // 5000 (config is null)
const retries = config?.retries ?? 3;       // 3

// Logical assignment operators (ES2021)
let a = null;
a ??= "default";     // a = "default" (only assigns if null/undefined)

let b = false;
b ||= "fallback";    // b = "fallback" (assigns if falsy)

let c = 5;
c &&= c * 2;         // c = 10 (assigns if truthy)

// Real-world: safe data display in React
function UserCard({ user }) {
    return (
        <div>
            <h2>{user?.name ?? "Anonymous"}</h2>
            <p>{user?.profile?.bio ?? "No bio provided"}</p>
            <span>{user?.followerCount ?? 0} followers</span>
        </div>
    );
}
```

---

### Q5. How do ES6 modules (import/export) work?

```javascript
// Named exports: multiple exports per file
// math.js
export const PI = 3.14159;
export function add(a, b) { return a + b; }
export function multiply(a, b) { return a * b; }

// Default export: one per file
// user.service.js
export default class UserService {
    async getUser(id) { /* ... */ }
    async createUser(data) { /* ... */ }
}

// Importing
import UserService from './user.service.js';           // default import
import { PI, add, multiply } from './math.js';         // named imports
import { add as sum } from './math.js';                // alias
import * as MathUtils from './math.js';                // namespace import
import UserService, { PI } from './utils.js';          // both default + named

// Dynamic imports: code splitting (loads module on demand)
async function loadChart() {
    const { default: Chart } = await import('./chart.js'); // lazy
    return new Chart();
}

// Conditional import
const locale = navigator.language;
const translations = await import(`./locales/${locale}.js`).catch(
    () => import('./locales/en.js') // fallback to English
);

// Re-exporting (barrel files): index.js
export { default as UserService } from './user.service.js';
export { add, PI } from './math.js';
export * from './helpers.js';
// Consumer:
import { UserService, add } from './services'; // one import for all

// CommonJS vs ESM:
// CommonJS: require/module.exports — synchronous, Node.js default
// ESM: import/export — static, asynchronous, browser native
```

---

### Q6. What are Generators? When do you use them?

```javascript
// Generator: function that can pause execution and yield values one at a time

function* fibonacci() {
    let [a, b] = [0, 1];
    while (true) {
        yield a;
        [a, b] = [b, a + b];
    }
}

const fib = fibonacci();
fib.next(); // { value: 0, done: false }
fib.next(); // { value: 1, done: false }
fib.next(); // { value: 1, done: false }
fib.next(); // { value: 2, done: false }

// Take first 10 Fibonacci numbers
function take(generator, n) {
    const result = [];
    for (const value of generator) {
        result.push(value);
        if (result.length === n) break;
    }
    return result;
}
take(fibonacci(), 7); // [0, 1, 1, 2, 3, 5, 8]

// ✅ Use case: paginated data (load next page on demand)
function* paginatedFetcher(endpoint) {
    let page = 1;
    while (true) {
        const data = yield { loading: true };    // signal loading
        yield { data, loading: false };          // yield data
        page++;
    }
}

// ✅ Use case: ID generator
function* makeIdGenerator(prefix = "ID") {
    let id = 1;
    while (true) {
        yield `${prefix}-${String(id++).padStart(6, '0')}`;
    }
}

const idGen = makeIdGenerator("USR");
idGen.next().value; // "USR-000001"
idGen.next().value; // "USR-000002"

// ✅ Use case: custom range iterable
function* range(start = 0, end, step = 1) {
    for (let i = start; i < end; i += step) yield i;
}

[...range(0, 10, 2)]; // [0, 2, 4, 6, 8]
for (const n of range(1, 5)) console.log(n); // 1, 2, 3, 4
```

---

### Q7. Explain `Promise.all`, `Promise.race`, `Promise.allSettled`.

```javascript
// Multiple async operations

const fetchUser     = (id) => fetch(`/api/users/${id}`).then(r => r.json());
const fetchProducts = ()   => fetch('/api/products').then(r => r.json());
const fetchSettings = ()   => fetch('/api/settings').then(r => r.json());

// Promise.all: run in PARALLEL, wait for ALL
// ❌ Fails fast: if ONE rejects, whole thing rejects
async function loadDashboard(userId) {
    try {
        const [user, products, settings] = await Promise.all([
            fetchUser(userId),
            fetchProducts(),
            fetchSettings()
        ]);
        return { user, products, settings }; // all succeeded
    } catch (err) {
        // One of them failed — don't know which
        console.error("Dashboard load failed:", err);
    }
}

// Promise.allSettled: NEVER rejects — waits for all, reports each result
async function loadDashboardSafely(userId) {
    const results = await Promise.allSettled([
        fetchUser(userId),
        fetchProducts(),
        fetchSettings()
    ]);

    return {
        user:     results[0].status === 'fulfilled' ? results[0].value : null,
        products: results[1].status === 'fulfilled' ? results[1].value : [],
        settings: results[2].status === 'fulfilled' ? results[2].value : defaultSettings
    };
}

// Promise.race: first to SETTLE (resolve or reject) wins
async function fetchWithFallback(primaryUrl, fallbackUrl) {
    const controller = new AbortController();
    const timeout    = new Promise((_, reject) =>
        setTimeout(() => reject(new Error("Timeout")), 3000)
    );

    try {
        return await Promise.race([
            fetch(primaryUrl),
            timeout      // times out if fetch takes >3s
        ]);
    } catch {
        return fetch(fallbackUrl); // try fallback
    }
}

// Promise.any: first to RESOLVE wins (ignores rejections)
// Perfect for redundant requests (hit multiple servers, take fastest)
const fastest = await Promise.any([
    fetch("https://server1.example.com/data"),
    fetch("https://server2.example.com/data"),
    fetch("https://server3.example.com/data")
]);
```

---

### Q8. What is the spread operator vs `Object.assign`?

```javascript
// Shallow merge: both create new object by copying properties

const defaults = { theme: "light", lang: "en", timeout: 5000 };
const overrides = { theme: "dark", timeout: 10000 };

// Object.assign: mutates first argument
const config1 = Object.assign({}, defaults, overrides);
// { theme: "dark", lang: "en", timeout: 10000 }

// Spread: always creates new object (more readable)
const config2 = { ...defaults, ...overrides };
// { theme: "dark", lang: "en", timeout: 10000 }

// ⚠️ Both are SHALLOW — nested objects are still references!
const obj = { a: 1, nested: { b: 2 } };
const copy = { ...obj };

copy.nested.b = 99;    // modifies original too!
copy.a = 100;          // does NOT modify original

console.log(obj.nested.b); // 99 — shared reference!
console.log(obj.a);         // 1 — primitive, own copy

// Deep clone alternatives:
const deepCopy1 = structuredClone(obj); // modern, handles Date/Map/Set
const deepCopy2 = JSON.parse(JSON.stringify(obj)); // simple, no Date/fn/undefined

// Practical: updating nested state (React pattern)
const state = {
    user: { name: "Alice", age: 25 },
    settings: { theme: "light" }
};

// Update only user.age — must spread at each level
const newState = {
    ...state,
    user: { ...state.user, age: 26 }  // spread nested object too
};
// state.user.age is still 25 — no mutation!
```
