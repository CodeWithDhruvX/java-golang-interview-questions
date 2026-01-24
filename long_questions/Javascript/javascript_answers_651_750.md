# JavaScript Interview Questions & Answers (651-750)

## 🔹 13. Language Quirks & Internals (Questions 651-680)

**Q651: What is the difference between `Object.is()` and `===`?**
They are almost identical strict equality checks, but `Object.is()` handles two edge cases correctly:
1.  `Object.is(NaN, NaN)` is **true** (while `NaN === NaN` is false).
2.  `Object.is(+0, -0)` is **false** (while `+0 === -0` is true).

**Q652: What happens if you override `toString()` in an object?**
It affects how the object acts in string coercion contexts (e.g., `alert(obj)`, `obj + ""`).
```javascript
const obj = { toString: () => "Custom" };
console.log(obj + " String"); // "Custom String"
```

**Q653: How does `valueOf()` affect type coercion?**
`valueOf` is preferred over `toString` for numeric conversion. If an object has both, `valueOf` is called for primitives (except Date, which prefers toString).
```javascript
const obj = { valueOf: () => 10 };
console.log(obj * 2); // 20
```

**Q654: What is boxing and unboxing in JavaScript?**
*   **Boxing:** Automatic wrapping of primitive values (string, number, boolean) into their Object wrappers so you can access methods (e.g., `"abc".toUpperCase()`).
*   **Unboxing:** Converting the Object wrapper back to a primitive value (via `valueOf()`/`toPrimitive`).

**Q655: What is an arguments object shadowing named parameters?**
(Legacy behavior). In non-strict mode, modifying `arguments[0]` updates the named parameter `a`, and vice-versa. In strict mode (`"use strict"`), they are decoupled.

**Q656: Can an object have duplicate keys?**
Not in a standard Object. If you define duplicate keys, the **last one wins** (overwrites valid predecessors).
```javascript
const obj = { a: 1, a: 2 };
console.log(obj); // { a: 2 }
```

**Q657: How are floating-point numbers represented internally in JS?**
Standard IEEE-754 Double Precision (64-bit binary format).
*   1 bit sign
*   11 bits exponent
*   52 bits mantissa (fraction).

**Q658: How does `toPrecision()` differ from `toFixed()`?**
*   `toFixed(n)`: Limits **decimal places** (digits after decimal point).
*   `toPrecision(n)`: Limits **significant digits** (total digits).
```javascript
const num = 123.456;
num.toFixed(2);     // "123.46"
num.toPrecision(2); // "1.2e+2" (or "120")
```

**Q659: Why does `0.1 + 0.2 !== 0.3` in JavaScript?**
Due to binary floating-point precision errors. 0.1 and 0.2 cannot be represented exactly in binary. The sum is slightly larger than 0.3 (`0.30000000000000004`).

**Q660: How do you make a number always display two decimal places?**
`num.toFixed(2)`. Only for display (returns a String).

---

## 🔹 14. Advanced Array & Object Behavior (Questions 681-710)
*(Source numbered 661-690 usually matches section header numbers, sticking to source list numbers 661-710)*

**Q661: What is the output of `typeof []` and `typeof {}`?**
Both return `"object"`. Arrays are specialized objects in JS.

**Q662: How does `delete arr[index]` affect array length?**
It removes the value (making it `empty`/undefined) but **does not** update the `length`. It creates a "sparse array".
```javascript
const arr = [1, 2];
delete arr[0];
console.log(arr.length); // 2
console.log(arr[0]); // undefined
```

**Q663: How does sparse array differ from a regular array?**
Sparse arrays have "holes" (indices with no keys). Methods like `forEach`, `map`, `filter` **skip** holes. `find` and `for...of` do NOT skip holes (treat them as undefined).

**Q664: What is the difference between `Object.entries()` and `Object.values()`?**
*   `Object.values(obj)`: Returns array of values `[v1, v2]`.
*   `Object.entries(obj)`: Returns array of pairs `[[k1, v1], [k2, v2]]`.

**Q665: What’s the output of spreading a string into an array?**
It splits the string into characters.
`[..."Hello"]` -> `['H', 'e', 'l', 'l', 'o']`.

**Q666: What is the behavior of `Array.from()` with a string?**
Similar to spread syntax. Creates an array of characters.
`Array.from("123")` -> `['1', '2', '3']`.

**Q667: How does array destructuring skip elements?**
Using commas to omit indices.
```javascript
const [first, , third] = [1, 2, 3];
// first = 1, third = 3
```

**Q668: What’s the result of `new Array(5).map(x => 1)`?**
`[empty × 5]`. `map` skips empty slots in sparse arrays. Result is still an empty/sparse array.
**Fix:** `Array.from({length: 5}).map(...)` or `new Array(5).fill(0).map(...)`.

**Q669: How do object keys get stringified in maps?**
In `Object`, keys are always Strings (or Symbols). `obj[1]` becomes `obj["1"]`.
In `Map`, keys can be any type and are **not** coerced to strings.

**Q670: Can you use objects as keys in regular JS objects?**
No. They are coerced to `"[object Object]"`. So all objects usually map to the same key. Use `Map` instead.

---

## 🔹 15. Real-World Async Problems (Questions 711-740)
*(Source list 671-680)*

**Q671: How to implement a retry mechanism for failed async requests?**
Recursive loop or loop with `await` inside.
```javascript
async function fetchWithRetry(url, retries=3) {
    for (let i=0; i<retries; i++) {
        try {
            return await fetch(url);
        } catch(err) {
            if (i === retries-1) throw err;
        }
    }
}
```

**Q672: How to limit concurrent fetch calls?**
Use a semaphore or a batching library (like `p-limit`).
Maintain a counter of active requests. If active < limit, start next request.

**Q673: What’s the difference between polling and long-polling?**
*   **Polling:** Periodic requests (every 5s) regardless of update.
*   **Long-Polling:** Request stays open until server has data. Client immediately sends new request upon response.

**Q674: How to time out a Promise after N milliseconds?**
Use `Promise.race` with a timeout promise that rejects.
```javascript
const timeout = new Promise((_, rej) => setTimeout(rej, 1000, 'Timeout'));
await Promise.race([fetchData(), timeout]);
```

**Q675: How to implement a queue with Promises?**
Chain promises (`p = p.then(nextTask)`) or use an async/await loop consuming a task array.

**Q676: What is a deferred object? Can you simulate one?**
An object exposing `resolve` and `reject` externally.
```javascript
function defer() {
    let resolve, reject;
    const promise = new Promise((res, rej) => { resolve = res; reject = rej; });
    return { promise, resolve, reject };
}
```

**Q677: How to run async functions sequentially?**
Use `for...of` loop with `await`.
```javascript
for (const item of items) {
    await process(item);
}
```

**Q678: What is the difference between async iterator and regular iterator?**
*   Regular iterator: `next()` returns `{ val, done }`.
*   Async iterator: `next()` returns `Promise<{ val, done }>`. Used in `for await...of`.

**Q679: What happens when you `await` inside a loop?**
Inside `forEach`: Does NOT wait. (Fire and forget).
Inside `for/of`: Waits for each iteration to finish before starting next (Sequential).

**Q680: How does the async stack trace differ from a sync one?**
Async stack traces can lose context because the initiation of the task (call stack) is gone when the callback runs. Modern engines (Async Stack Traces) stitch them together to show the "async cause".

---

## 🔹 16. Custom Implementations (Questions 741-770)
*(Source list 681-690)*

**Q681: How would you implement a custom `forEach()`?**
```javascript
Array.prototype.myForEach = function(cb) {
    for (let i = 0; i < this.length; i++) {
        // Check if index exists (sparse array check)
        if (Object.prototype.hasOwnProperty.call(this, i)) {
            cb(this[i], i, this);
        }
    }
};
```

**Q682: Can you write your own version of `Array.prototype.map()`?**
```javascript
Array.prototype.myMap = function(cb) {
    const res = []; // Or new Array(this.length)
    for (let i = 0; i < this.length; i++) {
        if (i in this) {
            res[i] = cb(this[i], i, this);
        }
    }
    return res;
};
```

**Q683: Implement a deep equality checker (`deepEqual()`).**
Recursive comparison.
1. Check strictly `===`.
2. Check `typeof`.
3. Check `keys.length`.
4. Recursively call `deepEqual(a[key], b[key])`.

**Q684: How to write a function like `JSON.stringify()` (simple case)?**
Recursive function handling types (String -> quotes, Array -> `[]`, Object -> `{}`).

**Q685: Write a polyfill for `Promise.all`.**
```javascript
function promiseAll(promises) {
    return new Promise((resolve, reject) => {
        let results = [], completed = 0;
        if (promises.length === 0) resolve([]);
        
        promises.forEach((p, i) => {
            Promise.resolve(p).then(val => {
                results[i] = val;
                completed++;
                if (completed === promises.length) resolve(results);
            }).catch(reject);
        });
    });
}
```

**Q686: How would you implement `throttle()` without setTimeout?**
Using timestamps.
```javascript
function throttle(fn, delay) {
    let last = 0;
    return (...args) => {
        const now = Date.now();
        if (now - last >= delay) {
            last = now;
            fn(...args);
        }
    }
}
```

**Q687: Implement a deep clone function.**
(See Q65). Using recursion. Handle Array vs Object. Handle Date/RegExp special cases if needed.

**Q688: Implement your own `bind()` method.**
(Duplicate Q99).

**Q689: Implement a tiny reactive store in JavaScript.**
Use `Proxy` or Getters/Setters.
```javascript
const state = { value: 0 };
const listeners = new Set();

const reactive = new Proxy(state, {
    set(target, prop, val) {
        target[prop] = val;
        listeners.forEach(fn => fn(val));
        return true;
    }
});
```

**Q690: Implement an event emitter in JavaScript.**
Store listeners in a Map/Object: `{ "event": [cb1, cb2] }`.
`on`: push to array. `emit`: loop array and call.

---

## 🔹 17. Real-World Design Patterns (Questions 771-800)
*(Source list 691-700)*

**Q691: What is the Revealing Module Pattern?**
A variation of Module Pattern where you define all functions privately and return an object exposing pointers to the private functions you want public.

**Q692: How to implement a publish/subscribe pattern in JS?**
Similar to Event Emitter/Observer. Decouples sender and receiver.

**Q693: How is the Proxy pattern used in modern JS frameworks?**
Vue 3 uses `Proxy` to wrap state objects. When properties are accessed (get) or modified (set), the Proxy intercepts these operations to track dependencies and trigger UI updates.

**Q694: What is the Facade pattern in frontend dev?**
Providing a simplified API to a complex subsystem. Example: Creating a `Network` service that abstracts `fetch`, `headers`, `auth`, handling complexity internally.

**Q695: How does dependency injection work in JavaScript?**
Passing dependencies (services) into a function/class constructor rather than importing/creating them inside. Enhances testability.

**Q696: Explain how the Factory pattern works with JS classes.**
A static method or separate function responsible for creating instances, abstracting the `new` keyword logic or choosing subclass.

**Q697: How do decorators work in JS (TC39 proposal)?**
Functions that modify class behavior at design time. `@readonly` on a method calls the decorator function with the target descriptor, allowing modification (e.g. `writable: false`).

**Q698: What’s the difference between mixin and inheritance?**
*   **Inheritance:** 'is-a' relationship (rigid hierarchy).
*   **Mixin:** 'has-a' capability. Composing behavior into a class (flexible).

**Q699: How is the Command pattern used in undo/redo systems?**
Encapsulate actions as Objects (`{ execute(), undo() }`). Store history of commands. To Undo, pop command and call `undo()`.

**Q700: What is the Strategy pattern and how can JS implement it?**
Selecting an algorithm at runtime.
```javascript
const strategies = {
    json: data => JSON.stringify(data),
    xml: data => toXML(data)
};
strategies[type](data);
```

---

## 🔹 18. Browser-Related (Questions 801-830)
*(Source list 701-710)*

**Q701: How does the browser parse HTML+JS during page load?**
Parses HTML -> Builds DOM. If `<script>` found -> Pause DOM construction -> Download JS -> Execute JS -> Resume DOM (unless `defer`/`async`).

**Q702: What are critical rendering path optimizations?**
Minimize critical resources (CSS/JS files blocking render). Inline critical CSS. Defer non-critical JS. Compress assets.

**Q703: What’s the difference between DOMContentLoaded and load?**
*   **`DOMContentLoaded`**: HTML parsed, DOM built. (External resources like images may not be loaded).
*   **`load`**: Everything loaded (Images, CSS, Frames).

**Q704: How does `window.history.pushState()` work?**
Changes the URL in the browser bar **without** reloading the page. Adds an entry to history stack. Used by SPA routers.

**Q705: What is `location.hash` used for?**
The part of URL after `#`. Does not trigger server reload. Used for anchors or old-school client-side routing.

**Q706: What are intersection observers?**
API to asynchronously observe changes in the intersection of a target element with an ancestor/viewport. (e.g., Infinite Scroll, Lazy Loading).

**Q707: How does the `ResizeObserver` API work?**
Observes changes to Element's content rect (dimensions). More performant than `window.onresize`.

**Q708: What is the difference between passive and non-passive event listeners?**
`{ passive: true }` tells browser the handler will **not** call `preventDefault()`. Allows browser to optimize scrolling performance (does not wait for listener to finish).

**Q709: How do you detect when a tab/window becomes hidden or inactive?**
`document.visibilityState` ('visible'/'hidden') and `visibilitychange` event.

**Q710: What are performance bottlenecks with heavy DOM manipulation?**
Frequent **Reflows** (Layout calc) and **Repaints**. Reading layout properties (`offsetHeight`) after writing forces synchronous reflow (Layout Thrashing).
