# ⚡ 01 — JavaScript Core & Closures
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Closure internals and memory implications
- Scope chain, hoisting nuances
- IIFE patterns
- Currying and partial application
- Function composition
- Memoization
- `this` binding rules (call, apply, bind)

---

## ❓ Most Asked Questions

### Q1. Explain closure with a practical real-world use case.

```javascript
// Closure: inner function retains access to outer function's scope
// even after the outer function has returned.

// ✅ Real-world: Rate limiter using closures
function createRateLimiter(maxCalls, windowMs) {
    let callCount = 0;
    let windowStart = Date.now();

    return function(fn) {
        const now = Date.now();
        if (now - windowStart > windowMs) {
            callCount = 0;          // reset window
            windowStart = now;
        }
        if (callCount < maxCalls) {
            callCount++;
            return fn();
        } else {
            throw new Error(`Rate limit exceeded: ${maxCalls} calls per ${windowMs}ms`);
        }
    };
}

const limiter = createRateLimiter(5, 1000); // 5 calls per second
// callCount and windowStart are "private" — only accessible via limiter

// ✅ Real-world: Event counter
function makeCounter(initial = 0) {
    let count = initial;
    return {
        increment: () => ++count,
        decrement: () => --count,
        reset: ()    => (count = initial),
        get value()  { return count; }
    };
}

const counter = makeCounter(10);
counter.increment(); // 11
counter.reset();     // back to 10
// count is inaccessible directly — true encapsulation

// ⚠️ Closure pitfall: loop + setTimeout
for (var i = 0; i < 3; i++) {
    setTimeout(() => console.log(i), 100); // logs 3 3 3 (var is function scoped)
}

// Fix 1: use let (block scoped)
for (let i = 0; i < 3; i++) {
    setTimeout(() => console.log(i), 100); // logs 0 1 2
}

// Fix 2: IIFE to capture current i
for (var i = 0; i < 3; i++) {
    (function(j) {
        setTimeout(() => console.log(j), 100);
    })(i);
}
```

---

### Q2. How does `this` binding work? Explain all four rules.

```javascript
// Rule 1: Default binding — 'this' is global (or undefined in strict mode)
function show() {
    console.log(this); // Window (browser) or global (Node)
}
show();

// Rule 2: Implicit binding — 'this' is the object before the dot
const obj = {
    name: "Alice",
    greet() { console.log(this.name); }
};
obj.greet(); // "Alice"

// Implicit binding lost!
const fn = obj.greet;
fn(); // undefined — no object before the dot

// Rule 3: Explicit binding — call, apply, bind
function greet(greeting, punctuation) {
    return `${greeting}, ${this.name}${punctuation}`;
}
greet.call({ name: "Bob" }, "Hi", "!");         // "Hi, Bob!"
greet.apply({ name: "Bob" }, ["Hi", "!"]);      // "Hi, Bob!"
const boundGreet = greet.bind({ name: "Bob" }); // creates new function
boundGreet("Hello", ".");                        // "Hello, Bob."

// Rule 4: new binding — 'this' is the newly created object
function Person(name) {
    this.name = name;  // 'this' refers to new object
}
const p = new Person("Charlie");
console.log(p.name); // "Charlie"

// Arrow functions: NO own 'this' — inherits from lexical scope
class Timer {
    constructor() { this.seconds = 0; }

    start() {
        // Arrow function — 'this' is Timer instance
        setInterval(() => {
            this.seconds++;
            console.log(this.seconds);
        }, 1000);
    }
}
```

---

### Q3. Implement `Function.prototype.bind` from scratch.

```javascript
// Custom bind implementation — frequently asked in FAANG/product interviews
Function.prototype.myBind = function(context, ...args) {
    const fn = this; // the original function

    return function(...newArgs) {
        // Handle 'new' invocation: if used as constructor, 'this' is the new instance
        if (new.target) {
            return new fn(...args, ...newArgs);
        }
        return fn.apply(context, [...args, ...newArgs]);
    };
};

// Test
function multiply(a, b, c) {
    return `${this.prefix}: ${a * b * c}`;
}
const triple = multiply.myBind({ prefix: "Result" }, 3);
console.log(triple(4, 5)); // "Result: 60"

// Custom call
Function.prototype.myCall = function(context = {}, ...args) {
    const sym = Symbol(); // unique key to avoid property collision
    context[sym] = this;
    const result = context[sym](...args);
    delete context[sym];
    return result;
};

// Custom apply
Function.prototype.myApply = function(context = {}, args = []) {
    const sym = Symbol();
    context[sym] = this;
    const result = context[sym](...args);
    delete context[sym];
    return result;
};
```

---

### Q4. Explain currying and implement a generic `curry` function.

```javascript
// Currying: f(a, b, c) → f(a)(b)(c)
// Enables partial application, function composition, reusability

function curry(fn) {
    return function curried(...args) {
        if (args.length >= fn.length) {
            // All arguments supplied — invoke fn
            return fn.apply(this, args);
        }
        // Return partially applied function
        return function(...moreArgs) {
            return curried.apply(this, [...args, ...moreArgs]);
        };
    };
}

// Example usage
const add = curry((a, b, c) => a + b + c);
add(1)(2)(3);       // 6
add(1, 2)(3);       // 6
add(1)(2, 3);       // 6
add(1, 2, 3);       // 6

// Real-world: URL builder
const buildUrl = curry((protocol, domain, path) =>
    `${protocol}://${domain}/${path}`
);

const httpsBuilder = buildUrl("https");
const apiBuilder   = httpsBuilder("api.example.com");
apiBuilder("users");    // "https://api.example.com/users"
apiBuilder("products"); // "https://api.example.com/products"
```

---

### Q5. What is function composition? Implement `compose` and `pipe`.

```javascript
// compose: applies functions right-to-left
// pipe: applies functions left-to-right

const compose = (...fns) => x => fns.reduceRight((v, f) => f(v), x);
const pipe    = (...fns) => x => fns.reduce((v, f) => f(v), x);

// Example: Data transformation pipeline
const trim       = str => str.trim();
const capitalize = str => str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
const addSuffix  = str => `${str}!`;

const transform = pipe(trim, capitalize, addSuffix);
transform("  hello world  "); // "Hello world!"

// Real-world middleware-style example
const withLogging = fn => (...args) => {
    console.log("Calling with:", args);
    const result = fn(...args);
    console.log("Result:", result);
    return result;
};

const withTiming = fn => (...args) => {
    const start = performance.now();
    const result = fn(...args);
    console.log(`Took ${performance.now() - start}ms`);
    return result;
};

const processData = pipe(
    withLogging,
    withTiming
)(data => data.map(x => x * 2));
```

---

### Q6. Implement a memoize function with cache size limit.

```javascript
// Production-grade memoization with LRU cache limit
function memoize(fn, { maxSize = 100 } = {}) {
    const cache = new Map(); // Map preserves insertion order

    return function(...args) {
        const key = JSON.stringify(args);

        if (cache.has(key)) {
            // LRU: move to end (most recently used)
            const value = cache.get(key);
            cache.delete(key);
            cache.set(key, value);
            return value;
        }

        const result = fn.apply(this, args);

        if (cache.size >= maxSize) {
            // Delete least recently used (first entry in Map)
            const firstKey = cache.keys().next().value;
            cache.delete(firstKey);
        }

        cache.set(key, result);
        return result;
    };
}

// Usage: expensive computation
const memoizedFib = memoize(function fib(n) {
    if (n <= 1) return n;
    return memoizedFib(n - 1) + memoizedFib(n - 2);
});

memoizedFib(40); // computed once, then cached
memoizedFib(40); // instant from cache

// ⚠️ Limitation: JSON.stringify doesn't handle circular refs, functions, etc.
// Production alternative: use WeakMap for object args
function memoizeForObjects(fn) {
    const cache = new WeakMap();
    return function(obj) {
        if (cache.has(obj)) return cache.get(obj);
        const result = fn(obj);
        cache.set(obj, result); // WeakMap — obj can still be garbage collected
        return result;
    };
}
```

---

### Q7. What is the Temporal Dead Zone (TDZ)?

```javascript
// TDZ: let/const variables exist in scope but are NOT accessible
// from the start of the block until the declaration is reached

function example() {
    // TDZ for 'x' starts here...
    console.log(typeof x); // ReferenceError (unlike var which would be 'undefined')
    // TDZ ends here ↓
    let x = 10;
    console.log(x); // 10
}

// Why TDZ exists: prevents using variables before they are properly initialized
// var: hoisted AND initialized to undefined (dangerous)
// let/const: hoisted but NOT initialized (TDZ guards until declaration)

// ⚠️ Constructor TDZ subtlety
class Animal {
    legs = 4; // class fields are initialized in constructor order
}
class Dog extends Animal {
    // super() must be called before 'this' — else TDZ-like error
    constructor() {
        // console.log(this.legs); // Error: must call super before accessing 'this'
        super();
        console.log(this.legs); // 4 — works after super()
    }
}
```

---

### Q8. Deep dive into Hoisting — functions vs variables vs classes.

```javascript
// 1. Function Declarations — fully hoisted (definition + body)
console.log(add(2, 3)); // 5 — works before declaration
function add(a, b) { return a + b; }

// 2. var — hoisted (declaration only, initialized to undefined)
console.log(x); // undefined (not ReferenceError)
var x = 10;
console.log(x); // 10

// 3. let / const — hoisted but in TDZ
// console.log(y); // ReferenceError
let y = 20;

// 4. Function Expressions — follow variable hoisting rules
// console.log(multiply(2, 3)); // TypeError (var) or ReferenceError (let/const)
var multiply = function(a, b) { return a * b; }; // multiply is undefined until this line

// 5. Classes — hoisted but in TDZ (like let)
// const obj = new Foo(); // ReferenceError
class Foo {}
const obj = new Foo(); // works

// Interview order prediction
console.log(foo); // undefined (var hoisted)
var foo = "bar";
function foo() {}  // function declaration hoisted above var
// After hoisting: function foo is available, then var foo = "bar" overrides it
// At runtime: foo is "bar"
```

---

### Q9. Explain WeakMap and WeakSet — when and why to use them.

```javascript
// WeakMap: keys must be objects; entries are weakly held (GC can collect them)
// Use case: private data, caching without memory leaks

// Private data pattern with WeakMap
const _private = new WeakMap();

class BankAccount {
    constructor(balance) {
        _private.set(this, { balance, transactions: [] });
    }

    deposit(amount) {
        const data = _private.get(this);
        data.balance += amount;
        data.transactions.push({ type: "deposit", amount });
    }

    get balance() { return _private.get(this).balance; }
}

const account = new BankAccount(1000);
account.deposit(500);
console.log(account.balance); // 1500
// account._private — undefined (truly private!)

// When account is GC'd, WeakMap entry is also cleaned up — no memory leak!

// WeakSet: holds weak references to objects
// Use case: tracking without preventing GC
const visited = new WeakSet();

function processNode(node) {
    if (visited.has(node)) return; // cycle detection
    visited.add(node);
    // process...
}

// ⚠️ Cannot iterate WeakMap/WeakSet — by design (GC can remove at any time)
// ✅ Use Map/Set when you need iteration; WeakMap/WeakSet for memory-safe caching
```

---

### Q10. Implement `debounce` and `throttle` from scratch.

```javascript
// Debounce: delays execution until N ms of inactivity
// Use: search input, resize handler, form validation

function debounce(fn, delay, { leading = false } = {}) {
    let timeoutId;
    let isLeadingCalled = false;

    return function(...args) {
        const context = this;

        if (leading && !isLeadingCalled) {
            fn.apply(context, args);
            isLeadingCalled = true;
        }

        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => {
            if (!leading) fn.apply(context, args);
            isLeadingCalled = false;
        }, delay);
    };
}

// Throttle: executes at most once per N ms
// Use: scroll handler, button click prevention, API rate limiting

function throttle(fn, interval) {
    let lastTime = 0;
    let timeoutId;

    return function(...args) {
        const context = this;
        const now = Date.now();
        const remaining = interval - (now - lastTime);

        if (remaining <= 0) {
            clearTimeout(timeoutId);
            lastTime = now;
            fn.apply(context, args);
        } else {
            // Trailing call — ensures final value is processed
            clearTimeout(timeoutId);
            timeoutId = setTimeout(() => {
                lastTime = Date.now();
                fn.apply(context, args);
            }, remaining);
        }
    };
}

// Usage
const debouncedSearch = debounce(searchAPI, 300);
input.addEventListener('input', debouncedSearch);

const throttledScroll = throttle(handleScroll, 100);
window.addEventListener('scroll', throttledScroll);
```

---

### Q11. What is the difference between `Object.freeze`, `Object.seal`, and `Object.preventExtensions`?

```javascript
// preventExtensions: no new properties can be added; existing can be modified/deleted
const obj1 = { a: 1 };
Object.preventExtensions(obj1);
obj1.b = 2;       // silently fails (or throws in strict mode)
obj1.a = 10;      // ✅ allowed
delete obj1.a;    // ✅ allowed

// seal: no new properties; existing can be modified but NOT deleted
const obj2 = { a: 1, b: 2 };
Object.seal(obj2);
obj2.c = 3;       // fails — no new properties
obj2.a = 10;      // ✅ allowed — modification ok
delete obj2.b;    // fails — can't delete

// freeze: no new properties, no modification, no deletion — fully immutable (shallow)
const obj3 = { a: 1, nested: { b: 2 } };
Object.freeze(obj3);
obj3.a = 5;              // fails
obj3.nested.b = 99;      // ✅ WORKS! freeze is SHALLOW
obj3.nested = {};        // fails  

// Deep freeze implementation (interview question)
function deepFreeze(obj) {
    Object.getOwnPropertyNames(obj).forEach(name => {
        const value = obj[name];
        if (typeof value === 'object' && value !== null) {
            deepFreeze(value);
        }
    });
    return Object.freeze(obj);
}

const config = deepFreeze({ db: { host: "localhost", port: 5432 } });
config.db.host = "remote"; // silently fails — truly immutable
```
