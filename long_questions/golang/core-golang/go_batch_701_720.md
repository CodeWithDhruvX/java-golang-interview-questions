## ⚙️ Go Internals & Runtime (Questions 701-720)

### Question 701: How does the Go scheduler work internally?

**Answer:**
Go uses an **M:N scheduler**.
- **G (Goroutine):** Application-level thread (stack + instruction pointer).
- **M (Machine):** OS Thread.
- **P (Processor):** Resource required to run Go code. `GOMAXPROCS` determines the number of 'P's.
The scheduler distributes **G**s across **P**s, which run on **M**s. It uses **Work Stealing** to balance load (if one P is idle, it steals Gs from another P's local queue).

---

### Question 702: What are GOMAXPROCS and how does it affect performance?

**Answer:**
`GOMAXPROCS` limits the number of OS threads that can execute user-level Go code simultaneously.
- **Default:** Number of CPU cores.
- **Performance:** Increasing it > CPUs doesn't help CPU-bound tasks (thrashing), but might help I/O bound tasks in older Go versions (though modern Netpoller handles I/O efficiently without blocking threads).

---

### Question 703: What’s the internal structure of a goroutine?

**Answer:**
A `g` struct (defined in `runtime/runtime2.go`) containing:
- **Stack:** Pointer to stack memory (start/current/end).
- **Sched:** Metadata for context switching (SP, PC).
- **Status:** (Waiting, Runnable, Running).
- **ID:** Unique ID.

---

### Question 704: How does stack growth work in Go?

**Answer:**
Go stacks are dynamic (start at 2KB).
Before a function call, the runtime checks if there's enough stack space.
- **If NO:** It allocates a new, larger stack (2x size), copies data from the old stack, updates pointers, and frees the old stack.
- This allows huge recursion without StackOverflow (up to 1GB usually).

---

### Question 705: How does garbage collection work in Go’s runtime?

**Answer:**
Go uses a **Concurrent, Tri-color Mark-and-Sweep** collector.
1.  **Mark Phase:** Traverses heap starting from Roots (Globals, Stack vars). Marks reachable objects (Grey -> Black).
2.  **Sweep Phase:** Reclaims White (unreachable) objects.
It avoids long "Stop-The-World" (STW) pauses by doing most work concurrently with the app, using a **Write Barrier** to track pointer changes during marking.

---

### Question 706: What are safepoints in the Go runtime?

**Answer:**
Points in the code where the runtime can safely stop a goroutine (for GC or Preemption).
- Traditionally: Function calls.
- Go 1.14+: **Asynchronous Preemption** allows stopping loops without function calls using OS signals (SIGURG).

---

### Question 707: What is cooperative scheduling in Go?

**Answer:**
Historically (pre-1.14), Go relied on goroutines strictly yielding control voluntarily (e.g., at function calls, channel ops, syscalls).
If a goroutine ran a tight loop `for {}`, it could starve others.
Modern Go is **Preemptive** (interrupts long-running Gs efficiently).

---

### Question 708: What are the stages of Go's garbage collector?

**Answer:**
1.  **Sweep Termination:** Stop the World (STW). Prepare.
2.  **Mark (Concurrent):** Scan roots, traverse heap. App runs alongside GC.
3.  **Mark Termination:** STW. Finish marking remaining items (dirty from Write Barrier).
4.  **Sweep (Concurrent):** Reclaim memory in background.

---

### Question 709: What is the role of the runtime package?

**Answer:**
It interacts with Go's runtime system.
- `runtime.Goexit()`: Kill current goroutine.
- `runtime.Gosched()`: Yield processor.
- `runtime.GC()`: Force garbage collection.
- `runtime.Caller()`: Inspect stack trace/program counters.

---

### Question 710: How does Go handle stack traces?

**Answer:**
The runtime unwinds the stack pointers stored in the `g` struct.
Since stacks can move (grow), Go uses a "Stack Map" generated at compile-time to know where pointers live on the stack, allowing accurate tracing and GC scanning.

---

### Question 711: How does the Go runtime manage memory allocation?

**Answer:**
Based on **TCMalloc** (Thread-Caching Malloc).
- **mspan:** Basic unit of memory (run of pages).
- **mcache:** Per-P (Processor) cache for small objects (no locks needed).
- **mcentral:** Shared global lists of spans (requires locks).
- **mheap:** Large allocations (direct from OS).

---

### Question 712: What’s the difference between a green thread and a goroutine?

**Answer:**
Conceptually similar (User-space threads).
- **Goroutine:** Managed by Go Runtime. Dynamically sized stack (2KB to 1GB). Fast startup.
- **Green Thread (Java/Python old):** Often fixed stack size. Mapped to OS threads differently.
Go's implementation is highly optimized for huge concurrency (millions).

---

### Question 713: What are finalizers in Go and how do they work?

**Answer:**
`runtime.SetFinalizer(obj, func(obj *T))`.
Runs when `obj` is garbage collected.
**Caveat:** No guarantee when (or if) it runs. Do not use for critical cleanup (files/sockets). Use `defer` or explicit `Close()` instead.

---

### Question 714: What is the role of `goexit` internally?

**Answer:**
`runtime.Goexit` unwinds the stack of the current goroutine, running all deferred functions, then terminates the goroutine.
It does **not** panic (recover won't catch it).
The main function logic (implicitly) calls this when it finishes.

---

### Question 715: How does Go avoid stop-the-world pauses?

**Answer:**
By running the **Marking** phase concurrently.
The "Write Barrier" ensures that if the app modifies pointers while GC is scanning, those objects are re-queued (greyed) so GC doesn't miss them.
STW is now extremely short (<0.5ms generally) just to start/finish the cycle.

---

### Question 716: How does memory fragmentation affect Go programs?

**Answer:**
Go's allocator (TCMalloc style) segregates objects by size class (e.g., all 32-byte objects go in one span).
This minimizes fragmentation significantly compared to a simple free-list allocator.

---

### Question 717: What’s the meaning of “non-preemptible” code in Go?

**Answer:**
Code that cannot be interrupted by the scheduler.
Pre-1.14: Tight loops.
Post-1.14: Code holding `unsafe.Pointer` logic or certain system calls might still delay preemption slightly, but true non-preemptible code is rare now.

---

### Question 718: What are M:N scheduling models and how does Go implement it?

**Answer:**
**M** Goroutines mapping to **N** OS Threads.
Go Multiplexes Gs onto Ms.
If a G blocks on I/O (System call), the M might block, so the scheduler spins up a new M (or reuses idle one) to keep other Ps running.
When I/O finishes, G joins the run queue.

---

### Question 719: How does Go detect deadlocks at runtime?

**Answer:**
The runtime tracks waiting goroutines.
If **all** goroutines are asleep (waiting on channels/mutexes) and no other work (timers/netpoll) is active, the runtime panics: `fatal error: all goroutines are asleep - deadlock!`.

---

### Question 720: What are the internal states of a goroutine?

**Answer:**
- **_Gidle:** Just allocated.
- **_Grunnable:** In a run queue, waiting for P.
- **_Grunning:** Executing on M.
- **_Gsyscall:** In a system call.
- **_Gwaiting:** Blocked (channel/lock/timer).
- **_Gdead:** Unused/Free.

---
