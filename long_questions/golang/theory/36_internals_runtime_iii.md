# ⚙️ Go Theory Questions: 701–720 Go Internals & Runtime

## 701. How does the Go scheduler work internally?

**Answer:**
Go uses a **Work-Stealing Scheduler** based on the GMP model.
**G (Goroutine)**: The code to run.
**M (Machine)**: The OS Thread.
**P (Processor)**: The resource context (queue) needed to execute a G.

P holds a Local Run Queue of Gs. M binds to P and executes Gs from P's queue.
If P's queue is empty, the M attempts to "steal" half the Gs from another P's queue. This balances the load across all CPU cores dynamically.

---

## 702. What are GOMAXPROCS and how does it affect performance?

**Answer:**
`GOMAXPROCS` limits the number of **P**s (Processors) that can execute user-level code simultaneously.
Default = `runtime.NumCPU()`.
If set to 1: The program runs concurrently but not in parallel (no two lines of code run at the exact same nanosecond).
Increasing it > NumCPU usually degrades performance due to Context Switching overhead, unless the app is heavily I/O bound and not using the Netpoller (which is rare).

---

## 703. What’s the internal structure of a goroutine?

**Answer:**
A goroutine is defined by the `g` struct in the runtime.
Key components:
1.  **Stack**: A pointer to the dynamically growing stack (starts at 2KB).
2.  **sched**: Keeps the Program Counter (PC) and Stack Pointer (SP) for context switching.
3.  **status**: State (Runnable, Running, Waiting, Dead).
4.  **m**: Pointer to the Machine (M) currently running it (if any).

---

## 704. How does garbage collection work in Go’s runtime?

**Answer:**
Go uses a **Concurrent Mark-Sweep (CMS)** collector with a **Write Barrier**.
It is non-generational (mostly) and non-compacting.
1.  **Mark**: Traverses the object graph from Roots. Marks live objects.
2.  **Sweep**: Scans the heap. Reclaims unmarked memory.
The "Pacer" algorithm determines when to start the next cycle triggered by the Heap Backup ratio (GOGC), targeting a specific CPU usage (25% during GC) to minimize latency.

---

## 705. What are safepoints in the Go runtime?

**Answer:**
A **Safepoint** is a location in code where the runtime knows the stack is in a consistent state and can be scanned by the GC.
Go inserts safepoint checks (preamble) at function calls.
In previous versions, tight loops `for { i++ }` had no function calls, so the GC could not stop them (causing STW delays). Go 1.14 introduced **Asynchronous Preemption** (using Unix signals) to interrupt goroutines anywhere, even in tight loops.

---

## 706. What is cooperative scheduling in Go?

**Answer:**
Cooperative means the Goroutine *voluntarily* yields control to the scheduler.
Yield points:
1.  Channel send/receive.
2.  I/O operations (Network/File).
3.  `runtime.Gosched()`.
4.  Function calls (Check stack guard).
Before 1.14 (Preemptive), if you wrote a loop that did none of these, you hogged the thread forever. Now, the scheduler is preemptive, but it still heavily favors cooperative points for efficiency.

---

## 707. What are the stages of Go's garbage collector?

**Answer:**
1.  **Sweep Termination**: Stop the World (STW). Finish up the last cycle.
2.  **Mark Phase (Concurrent)**: Turn on Write Barrier. Scan roots. Traverse heap. App keeps running (albeit slightly slower due to barrier).
3.  **Mark Termination**: STW. Rescan globals/changed stacks. Close Write Barrier.
4.  **Sweep (Concurrent)**: Reclaim memory in the background as the app allocates new stuff.

---

## 708. What is the role of the `runtime` package?

**Answer:**
The `runtime` is Go's "VM" (though it compiles to machine code).
It handles:
1.  **Memory Management**: Allocator (malloc) and GC.
2.  **Concurrency**: Scheduler (GMP) and Channels.
3.  **Reflection**: Type system implementation.
It replaces the standard C library (`libc`) for these tasks, interacting directly with syscalls, which allows Go binaries to be static and self-contained.

---

## 709. How does Go handle stack traces?

**Answer:**
Go stacks are text-friendly frames.
When a panic occurs, the runtime walks the stack (via frame pointers).
It resolves the Program Counter (PC) to Function Name + Line Number using the **DWARF** symbol table embedded in the binary. This is why stripping symbols (`-ldflags="-s -w"`) reduces binary size but makes stack traces useless (just hex addresses).

---

## 710. How does the Go runtime manage memory allocation?

**Answer:**
Based on **TCMalloc** (Thread-Caching Malloc).
Hierarchy:
1.  **mcache**: Per-P cache (No Lock). Fast path for small objects.
2.  **mcentral**: Shared list of spans. (Lock).
3.  **mheap**: Global heap requesting 64MB chunks from OS. (Lock).
This hierarchy ensures multiple threads can allocate memory simultaneously without fighting over a global mutex 99% of the time.

---

## 711. What’s the difference between a green thread and a goroutine?

**Answer:**
"Green Thread" is a generic term for user-space threads (Java pre-1.2, Python Gevent).
Goroutines are Go's implementation of green threads.
Differences from traditional green threads:
1.  **Stack**: Dynamic (Grow/Shrink).
2.  **Scheduling**: Multicore (M:N). Many green thread implementations are 1:N (single kernel thread), meaning they can't parallellize interpretation. Go's can.

---

## 712. What are finalizers in Go and how do they work?

**Answer:**
`runtime.SetFinalizer(obj, func(obj *Type))`.
This function runs when `obj` is garbage collected.
It is **NOT** deterministic. It might run seconds later, or never (if program exits).
Use logic: Cleaning up CGO resources (freeing `malloc`'d C strings) that the Go GC doesn't know about. Never use it for essential cleanup like flushing file buffers (use `defer` instead).

---

## 713. What is the role of `goexit` internally?

**Answer:**
`runtime.Goexit()` terminates the calling goroutine immediately.
It does **not** crash the app.
It runs all `defer` statements before quitting.
This is what happens when you call `t.Fatal()` in a test: it kills the test goroutine but leaves the test runner alive.

---

## 714. How does Go avoid stop-the-world pauses?

**Answer:**
By doing the heavy lifting **Concurrently**.
Older GCs paused everything to Mark.
Go's GC marks *while* you mutate the graph.
To prevent you from hiding an object from the scanner, Go uses a **Hybrid Write Barrier**: "If user code moves a pointer during GC, color the target Grey (to be scanned)". This ensures consistency without stopping the world, keeping pauses sub-millisecond.

---

## 715. How does memory fragmentation affect Go programs?

**Answer:**
Go minimizes fragmentation via **Size Classes**.
Allocating 17 bytes actually allocates 32 bytes (the nearest class).
All 32-byte objects live in a dedicated "32-byte Span".
This means there are no "holes" of odd sizes.
However, **Virtual Memory Fragmentation** can happen if the heap grows huge (100GB) and then shrinks; Go might release memory to OS (`MADV_FREE`), but if a span has just 1 live object, the whole 8KB page is kept hostage.

---

## 716. What’s the meaning of “non-preemptible” code in Go?

**Answer:**
Code that cannot be paused.
Historically, tightness math loops.
Currently, **CGO** calls and **Syscalls** can be non-preemptible.
If a Goroutine calls a C function that sleeps for 1 hour, the Go runtime cannot pause it. The Scheduler handles this by detaching the M (thread) from the P, spinning up a new M to service the other Goroutines in the queue to prevent starvation.

---

## 717. What are M:N scheduling models and how does Go implement it?

**Answer:**
1:1 (Java/C++): 1 User Thread = 1 Kernel Thread. (Expensive).
N:1 (Python/JS): N User Threads = 1 Kernel Thread. (No Parallelism).
**M:N (Go)**: M User Goroutines map to N Kernel Threads.
Go achieves this via the **Scheduler**. It multiplexes thousands of Gs onto `GOMAXPROCS` threads, getting the best of both worlds: high concurrency + parallelism + low memory footprint.

---

## 718. How does Go detect deadlocks at runtime?

**Answer:**
The runtime has a background check.
If **all** Goroutines are asleep (waiting on a channel or mutex) and no other external event (timer/netpoller) can wake them up, the runtime panics: `fatal error: all goroutines are asleep - deadlock!`.
This only detects global deadlocks that freeze the entire app. It does not detect partial deadlocks (where 2 goroutines are stuck but the rest of the app is fine).

---

## 719. What are the internal states of a goroutine?

**Answer:**
1.  `_Gidle`: Just allocated.
2.  `_Grunnable`: In a run queue (Local or Global). Ready to execute.
3.  `_Grunning`: Currently executing on an M.
4.  `_Gsyscall`: Executing a system call.
5.  `_Gwaiting`: Blocked (Channel/Mutex/Network).
6.  `_Gdead`: Finished. Can be reused (freelist).

---

## 720. What happens when a panic occurs inside a goroutine?

**Answer:**
1.  Execution stops.
2.  Stack unwinding begins. `defer`s run in LIFO order.
3.  If `recover()` is called, execution resumes at the recover point.
4.  If not recovered, the **Entire Program Crashes** (exit 2).
A panic in one goroutine kills *all* goroutines unless recovered. This is why top-level goroutines in HTTP servers are wrapped in a `recover()` block.
