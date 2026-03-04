# ⚡ 04 — Performance, Memory & Advanced JS Internals
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- V8 engine internals: hidden classes, inline caching, JIT
- Garbage collection and memory management
- `WeakRef` and `FinalizationRegistry`
- Virtual scrolling / performant rendering
- `requestAnimationFrame`
- Memory profiling techniques
- `structuredClone` vs JSON.parse

---

## ❓ Most Asked Questions

### Q1. How does V8's JIT and hidden classes work?

```javascript
// V8 optimizes JS by creating "hidden classes" (shapes) for objects
// with the same property layout — enables C++-like property access

// ✅ Optimizable: consistent property order at construction
function Point(x, y) {
    this.x = x; // always same order → V8 creates one hidden class
    this.y = y;
}
const p1 = new Point(1, 2); // HiddenClass C0 → C1(x) → C2(x,y)
const p2 = new Point(3, 4); // Reuses same hidden class chain — FAST!

// ❌ Deoptimizes hidden classes: inconsistent property assignment
function Point2(x, y) {
    if (x > 0) this.x = x;  // conditional — not all instances have 'x'
    this.y = y;
    this.extra = undefined; // deleting or adding props later = new hidden class
}
delete p1.x; // ❌ forces transition to new hidden class

// ❌ Property order matters!
const obj1 = { x: 1, y: 2 }; // HiddenClass A
const obj2 = { y: 2, x: 1 }; // HiddenClass B (different order = different class)
// obj1 and obj2 can't share optimized code

// Inline Caching (IC): V8 caches the hidden class at each call site
function getX(obj) { return obj.x; } // V8 assumes obj always has same hidden class
getX(p1); // caches HiddenClass of p1
getX(p2); // same class? → fast monomorphic IC
// Polymorphic (2-4 different types): slower; Megamorphic (>4): gives up optimization

// V8 optimization pipeline:
// Source → Parser (AST) → Ignition (bytecode) → TurboFan (JIT compiled) → Machine code
// "Deopt": if assumption broken, falls back to interpreter
```

---

### Q2. Explain Garbage Collection in V8 — Generational GC.

```javascript
// V8's Orinoco Garbage Collector uses Generational GC:
// Two spaces:
//   New Space (Young Generation) — small, fast, frequently collected (Scavenge GC)
//   Old Space (Old Generation) — large objects, survived 2 GC cycles, Mark-Sweep-Compact

// Scavenge (Minor GC): Copy-based — fast because only touches live objects
// Step 1: Mark live objects in "from" semi-space
// Step 2: Copy live objects to "to" semi-space
// Step 3: Swap semi-spaces
// Objects surviving 2 scavenges promoted to Old Space

// Mark-Sweep-Compact (Major GC):
// Step 1: Mark all reachable objects from roots (global, stack)
// Step 2: Sweep and add unreachable to free list
// Step 3 (optional): Compact — move live objects together (reduces fragmentation)

// ✅ GC-friendly patterns:
// Avoid creating short-lived objects in hot paths
function renderLoop() {
    // ❌ creates new object every frame (GC pressure)
    // context.fillStyle = { r: 255, g: 0, b: 0 };

    // ✅ reuse objects
    const color = { r: 0, g: 0, b: 0 };
    return function render() {
        color.r = (color.r + 1) % 256; // mutate, not allocate
        // render frame
        requestAnimationFrame(render);
    };
}

// Object pooling to reduce GC
class ObjectPool {
    #pool = [];
    #create;

    constructor(factory, initialSize = 10) {
        this.#create = factory;
        this.#pool = Array.from({ length: initialSize }, factory);
    }

    acquire() {
        return this.#pool.pop() ?? this.#create();
    }

    release(obj) {
        this.#pool.push(obj);
    }
}

const particlePool = new ObjectPool(() => ({ x: 0, y: 0, vx: 0, vy: 0 }));
const particle = particlePool.acquire();
// use particle...
particlePool.release(particle); // returned to pool, not GC'd
```

---

### Q3. What are memory leaks in JavaScript and how to detect them?

```javascript
// Common memory leak sources:

// 1. ❌ Detached DOM nodes in closures
function setupButton() {
    const button = document.getElementById("btn");
    const handler = () => doSomething(button); // closure holds 'button'
    button.addEventListener('click', handler);

    // ✅ Fix: remove listener when done
    return () => button.removeEventListener('click', handler);
}

// 2. ❌ Global variable accumulation
var cache = {}; // grows forever
function process(key, data) {
    cache[key] = { data, timestamp: Date.now() }; // never cleaned
}

// ✅ Fix: bounded Map or WeakMap
const cache = new Map();
const MAX_CACHE = 1000;
function processFixed(key, data) {
    if (cache.size >= MAX_CACHE) {
        const firstKey = cache.keys().next().value;
        cache.delete(firstKey); // evict oldest
    }
    cache.set(key, data);
}

// 3. ❌ setInterval without cleanup
const id = setInterval(() => {
    // holds closure reference + interval runs forever if not cleared
}, 1000);
// ✅ clearInterval(id) when component unmounts

// 4. ❌ Closures capturing large objects
function bigLeak() {
    const HUGE_DATA = new Array(1000000).fill("data");

    return function() {
        // Only uses tiny part but closure keeps all of HUGE_DATA alive!
        return HUGE_DATA[0];
    };
}

// ✅ Fix: only capture what you need
function noLeak() {
    const first = new Array(1000000).fill("data")[0]; // capture only what's needed
    return () => first; // HUGE_DATA can be GC'd
}

// WeakRef: hold reference without preventing GC
class Cache {
    #map = new Map();

    set(key, value) {
        this.#map.set(key, new WeakRef(value));
    }

    get(key) {
        const ref = this.#map.get(key);
        if (!ref) return undefined;
        const value = ref.deref(); // returns value or undefined if GC'd
        if (!value) this.#map.delete(key); // cleanup dead entry
        return value;
    }
}
```

---

### Q4. Implement Virtual Scrolling for a list of 100,000 items.

```javascript
// Virtual scrolling: only render visible items in the DOM
// Dramatically reduces DOM nodes and improves performance

class VirtualList {
    #container;
    #items;
    #itemHeight;
    #visibleCount;
    #scrollTop = 0;

    constructor(container, items, itemHeight = 50) {
        this.#container = container;
        this.#items = items;
        this.#itemHeight = itemHeight;
        this.#visibleCount = Math.ceil(container.clientHeight / itemHeight) + 2; // +2 buffer

        this.#setup();
    }

    #setup() {
        const totalHeight = this.#items.length * this.#itemHeight;

        // Outer div: real scroll height
        this.#container.style.overflow = "auto";
        this.#container.style.position = "relative";

        // Spacer: creates scroll range without DOM nodes
        this.spacer = Object.assign(document.createElement("div"), {
            style: `height: ${totalHeight}px; position: absolute; width: 100%;`
        });

        // Viewport: holds only visible nodes
        this.viewport = Object.assign(document.createElement("div"), {
            style: `position: sticky; top: 0;`
        });

        this.#container.append(this.spacer, this.viewport);
        this.#container.addEventListener("scroll", this.#onScroll.bind(this), { passive: true });
        this.#render();
    }

    #onScroll = throttle(() => {
        this.#scrollTop = this.#container.scrollTop;
        this.#render();
    }, 16); // 60fps

    #render() {
        const startIndex = Math.floor(this.#scrollTop / this.#itemHeight);
        const endIndex = Math.min(startIndex + this.#visibleCount, this.#items.length);

        const fragment = document.createDocumentFragment();
        for (let i = startIndex; i < endIndex; i++) {
            const el = document.createElement("div");
            el.style.cssText = `
                position: absolute;
                top: ${i * this.#itemHeight}px;
                height: ${this.#itemHeight}px;
                width: 100%;
            `;
            el.textContent = this.#items[i].name;
            el.dataset.index = i;
            fragment.appendChild(el);
        }

        this.viewport.innerHTML = "";
        this.viewport.appendChild(fragment); // single reflow
    }
}

// Usage — 100,000 items, only ~20 DOM nodes at any time
const list = new VirtualList(
    document.getElementById("container"),
    Array.from({ length: 100_000 }, (_, i) => ({ id: i, name: `Item ${i}` })),
    40 // 40px per item
);
```

---

### Q5. Optimize rendering with `requestAnimationFrame` and `IntersectionObserver`.

```javascript
// requestAnimationFrame: schedule work before next paint (60fps = 16ms window)

// ❌ Bad: multiple DOM reads/writes causing forced reflow (layout thrashing)
function badAnimation(elements) {
    elements.forEach(el => {
        const width = el.offsetWidth;      // READ — triggers layout
        el.style.width = width + 1 + "px"; // WRITE — invalidates layout
    });
    // Each iteration reads AFTER write → forced synchronous layout (expensive!)
}

// ✅ Good: batch reads then writes (FastDOM pattern)
function goodAnimation(elements) {
    requestAnimationFrame(() => {
        // READ phase
        const widths = elements.map(el => el.offsetWidth);
        // WRITE phase
        elements.forEach((el, i) => {
            el.style.width = widths[i] + 1 + "px";
        });
    });
}

// ✅ Smooth counter animation using rAF
function animateCounter(target, duration = 1000, element) {
    const start = performance.now();
    const from = parseInt(element.textContent) || 0;

    function update(now) {
        const elapsed = now - start;
        const progress = Math.min(elapsed / duration, 1);
        // Easing: ease-out cubic
        const eased = 1 - Math.pow(1 - progress, 3);
        element.textContent = Math.round(from + (target - from) * eased);
        if (progress < 1) requestAnimationFrame(update);
    }

    requestAnimationFrame(update);
}

// IntersectionObserver: lazy loading / infinite scroll without scroll listeners
const lazyImages = document.querySelectorAll("img[data-src]");
const imageObserver = new IntersectionObserver((entries, observer) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            const img = entry.target;
            img.src = img.dataset.src;
            img.removeAttribute("data-src");
            observer.unobserve(img); // stop observing once loaded
        }
    });
}, {
    rootMargin: "200px 0px", // start loading 200px before entering viewport
    threshold: 0
});

lazyImages.forEach(img => imageObserver.observe(img));
```

---

### Q6. What is `structuredClone` and when should you use it?

```javascript
// structuredClone: native deep clone (JS spec, available browsers 2022+, Node 17+)
// Handles: Date, Map, Set, ArrayBuffer, TypedArray, RegExp, Error
// Does NOT handle: functions, DOM nodes, class instances (loses methods)

const original = {
    date: new Date(),
    map: new Map([["key", "value"]]),
    set: new Set([1, 2, 3]),
    arr: new Uint8Array([1, 2, 3]),
    nested: { deep: { value: 42 } },
    regex: /pattern/gi
};

const clone = structuredClone(original);

// Date preserved as Date (not string like JSON.stringify)
clone.date instanceof Date; // true
// Map/Set preserved
clone.map instanceof Map;   // true

// ❌ structuredClone limitations
const withFn = { fn: () => "hello", value: 42 };
// structuredClone(withFn); // DataCloneError — can't clone functions

// Comparison table:
// JSON.parse/stringify: handles primitives only; dates → strings; no Map/Set/undefined
// Object.assign/{...}: shallow copy; references nested objects
// Lodash _.cloneDeep: works for most; handles functions (as-is); slower
// structuredClone: fast native; handles types JSON misses; no functions/DOM

// ✅ Use structuredClone for:
// - Web Worker message passing
// - Redux state immutability helpers
// - Caching API responses that include Date/Map/Set

// Transferable objects (zero-copy via postMessage):
const buffer = new ArrayBuffer(1024 * 1024); // 1MB
worker.postMessage({ buffer }, [buffer]); // transfers ownership, no copy!
// buffer is now empty in main thread — zero-copy transfer!
```

---

### Q7. Explain Proxy traps for meta-programming.

```javascript
// All available Proxy traps:
const handler = {
    get(target, prop, receiver) {},       // property access
    set(target, prop, value, receiver) {}, // property assignment
    has(target, prop) {},                  // 'in' operator
    deleteProperty(target, prop) {},       // delete operator
    apply(target, thisArg, args) {},       // function call
    construct(target, args, newTarget) {}, // new operator
    getPrototypeOf(target) {},
    setPrototypeOf(target, proto) {},
    isExtensible(target) {},
    preventExtensions(target) {},
    getOwnPropertyDescriptor(target, prop) {},
    defineProperty(target, prop, descriptor) {},
    ownKeys(target) {},
};

// ✅ Practical: Reactive state for UI (mini Vue/MobX)
function reactive(obj, onChange) {
    return new Proxy(obj, {
        set(target, key, value) {
            const old = target[key];
            target[key] = value;
            if (old !== value) onChange(key, value, old);
            return true;
        },
        get(target, key) {
            const val = target[key];
            // Deep reactivity: wrap nested objects
            if (typeof val === 'object' && val !== null) {
                return reactive(val, onChange);
            }
            return val;
        }
    });
}

const state = reactive({ user: { name: "Alice", age: 25 } }, (key, val) => {
    console.log(`[Reactive] ${key} changed to:`, val);
    rerenderComponent(); // trigger UI update
});

state.user.name = "Bob"; // [Reactive] name changed to: Bob
state.user.age = 26;     // [Reactive] age changed to: 26

// ✅ Read-only proxy
function readOnly(obj) {
    return new Proxy(obj, {
        set() { throw new TypeError("Object is read-only"); },
        deleteProperty() { throw new TypeError("Object is read-only"); },
    });
}

const config = readOnly({ apiUrl: "https://api.example.com" });
config.apiUrl = "hack"; // TypeError: Object is read-only
```
