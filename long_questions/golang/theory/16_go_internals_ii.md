# 🟢 Go Theory Questions: 301–320 Go Internals and Runtime

## 301. How does the Go scheduler work?

**Answer:**
The Go scheduler uses an **M:N** model to map M goroutines onto N OS threads.

It functions as a **Cooperative Scheduler** (historically, though now preemptive via async signals as of Go 1.14). It maintains a Global Run Queue and Local Run Queues for each Processor (P).

When a goroutine blocks (e.g., waiting for network IO), the scheduler parks it and picks another runnable goroutine from the local queue. This "Work Stealing" algorithm ensures that if one processor runs out of work, it steals half the goroutines from another processor, keeping all CPU cores saturated efficiently.

---

## 302. What is M:N scheduling in Golang?

**Answer:**
M:N means **M Goroutines** are multiplexed onto **N Kernel Threads**.

This is crucial because Kernel Threads are expensive (around 1MB stack, slow context switch), whereas Goroutines are cheap (2KB stack, fast switch).

Go creates a small number of Kernel Threads (typically equal to the number of CPU cores, `GOMAXPROCS`). It then schedules thousands of Goroutines onto these few threads. This allows Go programs to handle concurrent operations (like 100k HTTP connections) that would crash a standard Thread-per-Client server due to memory exhaustion.

---

## 303. How does the Go garbage collector work?

**Answer:**
Go uses a **Concurrent, Tri-color Mark-and-Sweep** collector.

It works in phases. First, strictly concurrently, it "marks" objects as White (candidate for deletion), Grey (needs checking), or Black (in use). It traverses the heap starting from "Roots" (global variables, stack frames).

Critically, it uses a **Write Barrier** to track pointers changed during the marking phase so it doesn't accidentally delete an object that was just assigned. Its main goal is **Low Latency** (pauses < 500 microseconds), often sacrificing raw throughput to avoid the massive "Stop The World" pauses seen in older Java GCs.

---

## 304. What are STW (stop-the-world) events in GC?

**Answer:**
STW events are brief moments when the Runtime pauses all user goroutines to perform safe operations.

In Go's modern GC, there are two very short STW phases: strictly at the **start** (to enable write barriers) and at the **end** (to terminate the mark phase).

Historically, these pauses were long (hundreds of ms). Now, they are sub-millisecond. However, if you allocate millions of tiny objects extremely fast, you can still force the GC to work so hard that it steals significant CPU time from your application ("GC Thrashing").

---

## 305. How are goroutines implemented under the hood?

**Answer:**
A goroutine is a `struct` (called `g` in the runtime source) that holds its own **Stack Pointer**, **Program Counter**, and scheduling info.

Unlike an OS thread, it does not map 1:1 to a kernel resource. It starts with a dynamic 2KB stack.

The Runtime manages them using three main structs:
1.  **G (Goroutine)**: The code to run.
2.  **M (Machine)**: The actual OS thread executing the code.
3.  **P (Processor)**: A resource required to execute Go code (matches `GOMAXPROCS`).
This abstraction allows Go to swap Gs onto Ms extremely fast (nanoseconds) completely in user-space.

---

## 306. How does stack growth work in Go?

**Answer:**
Go stacks are **Contiguous and Dynamic**. They start small (2KB).

When a function call would overflow the current stack, the runtime detects this (using a "stack guard" check in the function prologue). It halts the goroutine, allocates a new, larger stack chunk (usually 2x size), copies all existing data to the new stack, and updates all pointers to point to the new addresses.

This "Stack Copying" is why you can have deep recursion in Go without StackOverflow errors, up to the limit of available system memory (or `SetMaxStack`).

---

## 307. What is the difference between blocking and non-blocking channels internally?

**Answer:**
Internally, a channel is a `hchan` struct protected by a Mutex (`lock`).

**Blocking**: If you send to a full channel, the runtime adds your goroutine (`G`) to the channel's `recvq` (wait queue) and calls into the scheduler to park you (`gopark`). You consume 0 CPU while waiting.

**Non-blocking** (via `select` with `default`): The compiler generates a special runtime call (`selectnbsend`). It checks the lock; if the buffer is full, it returns `false` immediately instead of parking the goroutine. This avoids the expensive context switch associated with blocking.

---

## 308. What is a GOMAXPROCS and how does it affect execution?

**Answer:**
`GOMAXPROCS` controls the number of **P** (Processors) in the scheduler, which limits how many user-level Go threads can execute code *simultaneously* on the CPU.

By default, it equals `runtime.NumCPU()`.

If you set `GOMAXPROCS=1`, your program becomes strictly concurrent but not parallel (time-slicing on a single core). Increasing it beyond the physical core count usually hurts performance due to cache contention and context switching overhead, unless your program is heavily I/O blocked and syscall-heavy (though the netpoller handles most I/O efficiently anyway).

---

## 309. How does Go manage memory fragmentation?

**Answer:**
Go uses a **TCMalloc-based** (Thread-Caching Malloc) allocator designed to minimize fragmentation.

It divides memory into "span classes" based on object size (e.g., separate spans for 8-byte objects, 16-byte objects). When you need 12 bytes, it rounds up to 16 and gives you a slot from the 16-byte span.

Because all objects in a span are the exact same size, there are no "holes" of unusable size between them. Additionally, `mcentral` and `mheap` coalesce free spans to return large chunks of memory to the OS if they remain unused for a long time (`MADV_FREE`).

---

## 310. How are maps implemented internally in Go?

**Answer:**
A Go map is a Hash Table implemented as an array of **Buckets**.

Each bucket holds up to 8 key/value pairs. When you look up a key, Go hashes it. The **Low-Order bits** of the hash select the bucket. The **High-Order bits** (Top Hash) are stored inside the bucket to allow fast comparisons without Dereferencing the full key.

relationships. If a bucket overflows (more than 8 collisions), Go chains an "overflow bucket". When the total load factor is too high (avg 6.5 items/bucket), the map grows (doubles in size) and incrementally evacuates keys to the new buckets.

---

## 311. How does slice backing array reallocation work?

**Answer:**
When you `append()` to a full slice, Go creates a new backing array.

The growth algorithm is roughly:
1.  If capacity < 256, it doubles (2x).
2.  If capacity > 256, it grows by a factor of ~1.25x (optimized for smoother transitions).

It then **copies** the old data to the new array and returns a slice header pointing to the new memory. The old array is eventually garbage collected if no other slice refers to it. This copying cost is why pre-allocating slices (`make([]int, 0, 1000)`) is a major performance optimization.

---

## 312. What is the zero value concept in Go?

**Answer:**
Go guarantees that all memory is initialized to a known state, never random garbage.

*   Pointers/Interfaces/Maps/Channels: `nil`
*   Numbers: `0`
*   Booleans: `false`
*   Strings: `""` (empty string, not nil)
*   Structs: Recursive zero value for all fields.

This design eliminates "Uninitialized Variable" bugs common in C. However, it requires careful design: our structs must be "ready to use" with their zero values (e.g., `sync.Mutex` works without a constructor, but a custom specific type might need a `New()` function).

---

## 313. How does Go avoid data races with its memory model?

**Answer:**
The **Go Memory Model** defines "Happens-Before" relationships.

It guarantees that a write in one goroutine is observed by a read in another *only if* there is an explicit synchronization event (Channel send/receive, Mutex Lock/Unlock, WaitGroup Wait).

Without these events, the compiler and CPU are free to reorder instructions for speed. Go does *not* behave like `volatile` in Java or C++. If you share a variable without a lock, you have undefined behavior. We use the **Race Detector** (`-race`) to strictly enforce these rules during testing.

---

## 314. What is escape analysis and how can you visualize it?

**Answer:**
Escape Analysis is a compiler phase that decides where to allocate a variable: **Stack** or **Heap**.

If a variable's reference "escapes" the function (e.g., returned to a caller, stored in a global, sent to a channel), it must go to the **Heap** (GC managed). If it stays local, it goes to the **Stack** (free/fast hygiene).

We visualize it using `go build -gcflags="-m"`. The compiler will print decisions like `moved to heap: x` or `new(Result) does not escape`. Optimizing code often involves tweaking logic to keep variables on the stack ("Zero Allocation").

---

## 315. How are method sets determined in Go?

**Answer:**
Method sets depend on whether you have a **Value** `T` or a **Pointer** `*T`.

*   The method set of `T` contains only methods with receiver `(t T)`.
*   The method set of `*T` contains methods with receiver `(t *T)` **AND** `(t T)`.

This is why you can call a pointer method on a value (if addressable), because Go automatically takes the address (`&val.Method()`). But you *cannot* assign a value `T` to an interface that requires a `*T` method, because the interface only sees the strictly defined method set of `T`.

---

## 316. What is the difference between pointer receiver and value receiver at runtime?

**Answer:**
At runtime, the difference is strictly about **Copying** and **Mutability**.

**Value Receiver**: Examples copies the struct. If the struct is large (e.g., 1KB), this copy is expensive. Any changes inside the method affect only the copy, not the original.

**Pointer Receiver**: Copies only the pointer (8 bytes). It’s fast. Modifications affect the original struct.

The Runtime treats them differently in interfaces: a Pointer Receiver method *cannot* satisfy an interface if stored as a Value, because the Value might not be addressable (e.g., a map element or temporary intermediate value).

---

## 317. How does Go handle panics internally?

**Answer:**
A `panic()` stops normal execution flow and begins "Unwinding the Stack."

Internally, the runtime runs all deferred functions in LIFO order. If one of them calls `recover()`, the unwinding stops, and execution resumes normally *from the point of return of the recovering function*.

If the stack unwinds all the way to the top of the goroutine without recovery, the runtime prints the stack trace and exits the program (typically with exit code 2). This is distinct from C++ exceptions—it is designed for "unrecoverable" errors (bugs), not flow control.

---

## 318. How is reflection implemented in Go?

**Answer:**
Reflection maps the empty interface `interface{}` to internal types.

An interface is a pair: `(Type, Value)`. The `reflect` package inspects this pair.
`reflect.TypeOf(i)` returns the `Type` metadata (struct fields, method signatures) stored in the binary's read-only section.
`reflect.ValueOf(i)` wraps the actual data, allowing read/write access (if mutable).

It is slow because it defeats compiler optimizations (inlining) and relies on heavy runtime type checks. We avoid it in hot paths, but it's essential for generic serialization libraries like `encoding/json`.

---

## 319. What is type identity in Go?

**Answer:**
Type Identity rules determine if two types are interchangeable.

Two named types (`type A int`, `type B int`) are **Different**, even if underlying types match. You must explicitly cast `A(b)`.

However, if at least one type is **Unnamed** (just `[]int` or `struct{x int}`), they are identical if their underlying structures match.
Crucially, type aliases (`type A = B`) are **Identical**. They are literally the same type at compile time, used solely for refactoring or readability (like `byte` is just `uint8`).

---

## 320. How are interface values represented in memory?

**Answer:**
An interface is a two-word struct (16 bytes on 64-bit systems).
1.  `tab` (ITable): A pointer to a table containing the specific Type info and a list of Function Pointers for the methods efficiently mapped to that type.
2.  `data`: A pointer to the actual concrete value (or a copy of it on the heap if it’s huge).

Calling an interface method `i.Method()` compiles to `i.tab.fun[0](i.data)`. This "Virtual Table" lookup is faster than reflection but slightly slower than a direct call due to CPU branch prediction misses.
