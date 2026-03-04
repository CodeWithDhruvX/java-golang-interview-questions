# ⚙️ 01 — Node.js Internals & V8 Engine
> **Most Asked in Product-Based Companies** | ⚙️ Difficulty: Hard

---

## 🔑 Must-Know Topics
- Libuv architecture (Thread Pool, Event Loop internals)
- V8 Engine (JIT Compilation, Ignition, TurboFan)
- Garbage Collection in V8 (Orinoco, Scavenger, Mark-Sweep)
- Managing Memory and Buffer instances
- Native C++ Addons with node-gyp or N-API

---

## ❓ Frequently Asked Questions

### Q1. Deep dive into the V8 Engine. How does it compile JavaScript to machine code?

**Answer:**
Google's V8 engine takes JavaScript code and compiles it directly into machine code before executing it, which is much faster than interpreting it line by line.

**The V8 Pipeline:**
1. **Parser & AST:** The JS source code is parsed into an Abstract Syntax Tree (AST).
2. **Ignition (Interpreter):** Takes the AST and compiles it rapidly into unoptimized bytecode, which begins executing immediately.
3. **TurboFan (Optimizing Compiler):** While the bytecode is running, V8 watches for "hot" functions (functions executed many times) and the types of data passed to them.
   - It sends these hot functions to *TurboFan*, which generates highly optimized machine code based on the observed data types.
4. **Deoptimization:** If the assumptions made by TurboFan turn out to be false (e.g., a function suddenly receives strings instead of integers), V8 *deoptimizes* the code, reverting it back to Ignition's unoptimized bytecode.

---

### Q2. How does Garbage Collection (GC) work in V8?

**Answer:**
V8 uses a Generational Garbage Collector called *Orinoco*. Memory is divided into two generations:

1. **Young Generation (Nursery / Scavenger):**
   - New objects are allocated here. It is very small (1-8 MB) and fills up quickly.
   - Most objects "die young" (become unreachable quickly).
   - GC here uses the *Scavenge* algorithm (copying semi-space). It divides the young space into "From" and "To" spaces. Live objects are copied to the "To" space, and the "From" space is completely freed.
   - If an object survives two scavenge cycles, it is *promoted* (tenured) to the Old Generation.

2. **Old Generation:**
   - Holds long-lived objects. It is much larger.
   - GC here uses a *Mark-and-Sweep* (and Compact) algorithm.
   - **Mark:** Traverses the object graph completely, marking live objects.
   - **Sweep:** Traverses memory sequentially, freeing memory of dead (unmarked) objects.
   - **Compact:** Moves surviving objects to contiguous memory to reduce fragmentation.

*Modern V8 GC runs concurrently and in parallel across multiple helper threads to keep GC pauses under 1 millisecond.*

---

### Q3. Explain Libuv and its role in Node.js. Does Node.js really only use one thread?

**Answer:**
Libuv is a multi-platform C library focused on asynchronous I/O. It provides the Event Loop to Node.js and acts as the bridge to the underlying operating system.

**The "Single-Threaded" Myth:**
Node.js executes *JavaScript code* on a single thread (the Main Thread/Event Loop). However, Node.js itself is multithreaded. 

For asynchronous operations that cannot be handled non-blockingly at the OS level (like file system operations `fs.*`, DNS lookups, or crypto hashing), Libuv offloads the work to a **Worker Pool (Thread Pool)**.

- **Thread Pool:** By default, Libuv maintains a pool of 4 worker threads (can be increased up to 1024 via `process.env.UV_THREADPOOL_SIZE`).
- When an `fs.readFile` is called, the Main Thread passes it to the Thread Pool. A thread in the pool reads the file blockingly. Once done, it notifies the Event Loop, and the callback executes on the Main Thread.

---

### Q4. What is a Memory Leak in Node.js, and how do you trace it?

**Answer:**
A memory leak occurs when objects are no longer needed but remain reachable from the GC root, preventing the Garbage Collector from freeing their memory. Over time, memory consumption grows until the process runs out of memory and crashes (`FATAL ERROR: Ineffective mark-compacts near heap limit Allocation failed`).

**Common Causes:**
1. Storing data in global variables indefinitely.
2. Unclosed Event Listeners (e.g., `emitter.on('event', cb)` called repeatedly without `.off()`).
3. Closures that hold onto large objects unintentionally.
4. Caching data in memory without a size limit or TTL.

**How to trace it:**
1. **Heap Snapshots:** Use node `--inspect` and Chrome DevTools to take Heap Snapshots at different times. Compare them to see which objects are sticking around.
2. **`process.memoryUsage()`:** Monitor `heapUsed` and `rss`.
3. **Clinic.js:** Use `clinic memory` to profile and visualize the memory usage of the application.

---

### Q5. What are Node.js Buffers and why do we need them when we have arrays?

**Answer:**
The `Buffer` class in Node.js is designed to handle raw binary data outside V8's heap memory. It is a globally available class.

**Why use Buffers instead of Arrays?**
- V8's native JavaScript Arrays are not designed for purely binary data. They can be slow and memory-inefficient for large streams of raw bits and bytes.
- Buffers allocate a fixed chunk of memory (outside the V8 heap, in C++ land) which makes them extremely fast and memory-efficient for reading from files, network sockets, or handling cryptography data.

**Example:**
```javascript
// Allocating a 10-byte buffer filled with zeros
const buf = Buffer.alloc(10); 

// Creating a buffer from a string
const strBuf = Buffer.from('hello world', 'utf8');

// Modifying binary directly
buf[0] = 255; 
```
