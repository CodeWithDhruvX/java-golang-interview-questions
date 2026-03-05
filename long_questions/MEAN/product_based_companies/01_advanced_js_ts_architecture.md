# Advanced JavaScript, TypeScript & Architecture (Product-Based Companies)

Product-based companies expect a deep understanding of how the language works under the hood, how browser engines compile JS, memory management, and advanced TypeScript patterns for enterprise-scale applications.

## V8 Engine & JavaScript Internals

### 1. How does the V8 Engine compile and execute JavaScript?
V8 compiles JavaScript directly to native machine code before executing it, instead of interpreting it in real-time or compiling it to bytecode.
*   **Parser**: Parses JS code into an Abstract Syntax Tree (AST).
*   **Ignition (Interpreter)**: Generates bytecode from the AST. It runs quickly but the bytecode isn't highly optimized.
*   **TurboFan (Optimizing Compiler)**: Takes hot (frequently executed) bytecode from Ignition and compiles it into highly optimized machine code. If assumptions made during optimization turn out to be false (e.g., a function expects numbers but suddenly gets a string), TurboFan "de-optimizes" the code back to bytecode.

### 2. How does Garbage Collection work in V8? Explain Memory Leaks in JavaScript.
V8 divides memory into Generation spaces: **Young Generation** (newly allocated objects) and **Old Generation** (long-lived objects).
*   **Minor GC (Scavenger)**: Cleans up the Young Generation quickly and frequently.
*   **Major GC (Mark-and-Sweep/Compact)**: Cleans up the Old Generation. It "marks" live objects starting from the root (global object) and "sweeps" (frees) unreferenced objects.
**Memory leaks** occur when objects are no longer needed but are still referenced by the application, preventing the GC from collecting them.
**Common causes of memory leaks in JS:**
*   Accidental global variables (`window.myVar = data`).
*   Forgotten timers/intervals (`setInterval` that is never cleared).
*   Closures holding onto large scopes unnecessarily.
*   Detached DOM elements (referencing a DOM node in JS after it's removed from the document).

### 3. Explain the concept of "Macrotasks" and "Microtasks" in the Event Loop in detail.
*   **Macrotasks (Task Queue)**: `setTimeout`, `setInterval`, `setImmediate` (Node), I/O, UI rendering.
*   **Microtasks (Microtask Queue)**: `Promise.then()/.catch()/.finally()`, `MutationObserver`, `process.nextTick` (Node).
**Execution Order**:
1.  Execute synchronous code (Call Stack).
2.  Check the Microtask Queue. If not empty, execute *all* microtasks. (If a microtask adds another microtask, it gets executed in the same cycle, potentially blocking the event loop).
3.  Render UI (if needed in browsers).
4.  Pick ONE task from the Macrotask Queue and execute it.
5.  Repeat.

## Advanced Design Patterns

### 4. How would you implement a Publisher/Subscriber (Pub/Sub) pattern in vanilla JS?
The Pub/Sub pattern promotes loose coupling. Components communicate via an event channel without knowing about each other.
```javascript
class EventEmitter {
    constructor() {
        this.events = {};
    }
    subscribe(eventName, callback) {
        if (!this.events[eventName]) {
            this.events[eventName] = [];
        }
        this.events[eventName].push(callback);
        // Return an unsubscribe function
        return () => {
            this.events[eventName] = this.events[eventName].filter(cb => cb !== callback);
        }
    }
    publish(eventName, data) {
        if (this.events[eventName]) {
            this.events[eventName].forEach(cb => cb(data));
        }
    }
}
```

### 5. Explain the Singleton pattern and its drawbacks in Node.js.
A Singleton ensures a class has only one instance and provides a global point of access to it.
*   **In Node.js:** Modules are cached after the first time they are loaded (`require()`). This inherently acts like a Singleton.
*   **Drawbacks**: Singletons introduce global state into an application, which makes unit testing very difficult (state leaks between tests) and can cause hard-to-track bugs in highly concurrent applications. Dependency Injection is often preferred over Singletons.

## Advanced TypeScript

### 6. What are Utility Types in TypeScript? Explain `Partial`, `Pick`, `Omit`, and `Record`.
Utility types facilitate common type transformations.
*   `Partial<T>`: Constructs a type with all properties of `T` set to optional.
*   `Pick<T, K>`: Constructs a type by picking the set of properties `K` from `T`.
*   `Omit<T, K>`: Constructs a type by picking all properties of `T` and then removing `K`.
*   `Record<K, T>`: Constructs an object type whose property keys are `K` and whose property values are `T`. (e.g., `Record<string, number>` is a dictionary of string-number pairs).

### 7. What is a Discriminated Union?
A Discriminated Union is a pattern in TS where you combine a Union Type with a common literal field (the "discriminator") to allow the compiler to narrow down the specific type in a `switch` or `if` statement.
```typescript
interface Circle { kind: 'circle'; radius: number; }
interface Square { kind: 'square'; sideLength: number; }
type Shape = Circle | Square;

function getArea(shape: Shape) {
    // 'kind' is the discriminator
    switch (shape.kind) {
        case 'circle':
            return Math.PI * shape.radius ** 2; // TS knows this is a Circle
        case 'square':
            return shape.sideLength ** 2; // TS knows this is a Square
    }
}
```

### 8. Explain Type Guards and the `is` operator.
Type guards are expressions that perform a runtime check that guarantees the type in some scope.
```typescript
interface Fish { swim(): void; }
interface Bird { fly(): void; }

// Custom type guard using 'is'
function isFish(pet: Fish | Bird): pet is Fish {
    return (pet as Fish).swim !== undefined;
}

const myPet: Fish | Bird = getPet();
if (isFish(myPet)) {
    myPet.swim(); // TS knows myPet is Fish
} else {
    myPet.fly(); // TS knows myPet is Bird
}
```
