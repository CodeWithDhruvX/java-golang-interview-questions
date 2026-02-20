# ⚙️ **701–720: Go Internals & Runtime**

### 701. How does the Go scheduler work internally?
"Go uses an **M:N Scheduler**, meaning it multiplexes M OS threads onto N Goroutines.

It revolves around three entities: **G** (Goroutine), **M** (Machine/OS Thread), and **P** (Processor/Context).
The **P** holds a local queue of runnable **G**s. The **M** attaches to a **P** to execute them. If an M blocks (like in a syscall), the P detaches and grabs a new M to keep running the other Gs.

This architecture solves the 'thread-per-request' problem. I can have 100,000 idle Gs, and the scheduler efficiently parks them without consuming OS resources, waking them only when work is available."

#### Indepth
**Work Stealing**. If a P's local queue is empty, it doesn't just sit idle. It tries to "steal" half of the goroutines from another P's queue. This balances the load dynamically across cores, ensuring no CPU core is idle while there is work to be done anywhere in the system.

---

### 702. What are GOMAXPROCS and how does it affect performance?
"**GOMAXPROCS** controls the number of **P**s (Processors) available to execute Go code simultaneously. By default, it equals `runtime.NumCPU()`.

Technically, it limits *parallelism* (executing code on multiple cores at the exact same nanosecond), not *concurrency* (managing multiple tasks).

I rarely change this manually. However, in containerized environments (Kubernetes) with strict CPU quotas, leaving it at default can cause performace issues (throttling). I use `automaxprocs` to ensure Go sees the *quota* limit, not the host machine's total cores."

#### Indepth
**IO Polling**. The Go scheduler uses "Netpoller" (kqueue/epoll/IOCP). When a goroutine reads from the network, it is parked and moved to the Netpoller. The M (thread) is freed to run other Gs. When data arrives, the Netpoller moves the G back to a P's run queue. This is why Go is so efficient at IO-bound workloads.

---

### 703. What’s the internal structure of a goroutine?
"A goroutine is represented by a `g` struct (defined in `runtime/runtime2.go`).

It’s surprisingly simple. It contains its own **stack** (starting at 2KB), the **instruction pointer** (PC), the **stack pointer** (SP), and its status (Running, Waiting, Runnable).

Unlike OS threads which have a fixed large stack (1-2MB), the `g` struct allows the stack to grow and shrink dynamically. This lightweight nature is why I can treat goroutines as 'cheap' resources compared to Java threads."

#### Indepth
**Stack Copying**. When a stack grows (e.g. from 2KB to 4KB), the runtime allocates a new, larger segment and *copies* the entire stack contents to the new location. It then updates all pointers to the stack. This is transparent to the user but has a small performance cost during deep recursion.

---

### 704. How does garbage collection work in Go’s runtime?
"Go uses a **Concurrent, Tri-color Mark-and-Sweep** collector.

It views the heap as a graph.
1.  **Mark**: It starts from roots (global variables, active stack frames) and colors reachable objects 'Black'.
2.  **Sweep**: It walks the heap and reclaims memory from 'White' (unreachable) objects.

The key innovation is that it runs **concurrently** with my application code. It uses a 'Write Barrier' to track changes made while the GC is running, keeping pause times (STW) incredibly low—usually under 100 microseconds."

#### Indepth
**GOGC**. This variable controls the GC frequency. Default is 100 (run GC when heap grows 100%). Setting `GOGC=200` makes GC run less often (trading RAM for CPU). Setting `GOGC=off` disables it entirely. **Go 1.19** introduced a "Soft Memory Limit" (`GOMEMLIMIT`) which is a safer way to tune GC than raw percentages.

---

### 705. What are safepoints in the Go runtime?
"Safepoints are locations in the code where the runtime can safely stop a goroutine (for GC or scheduling).

Historically, these were placed at function preambles. This meant a tight loop like `for {}` could block the GC forever because it never hit a safepoint.

Since Go 1.14, the runtime uses **Asynchronous Preemption**. It sends a signal (`SIGURG` on Linux) to the thread, forcing it to stop at any instruction. This solved the 'tight loop' latency spikes I used to see in older Go versions."

#### Indepth
**sysmon**. The system monitor thread runs outside the scheduler. It wakes up every 20us-10ms. Its job is to detect long-running Gs (preemption), Retake Ps from blocked syscalls (handoff), and force GC if it hasn't run in 2 minutes. It is the "watchdog" of the runtime.

---

### 706. What is cooperative scheduling in Go?
"**Cooperative scheduling** means the running task must voluntarily yield control to the scheduler.

In early Go, a goroutine only yielded when it called a function or blocked on I/O. If it just crunched numbers, it hogged the CPU.

Go has moved away from purely cooperative scheduling to **Preemptive Scheduling**. Now, the `sysmon` background thread detects if a goroutine has been running too long (>10ms) and forcefully deschedules it, ensuring fairness so that other request handlers don't starve."

#### Indepth
**Fairness vs Throughput**. Preemption hurts throughput (context switching overhead) but guarantees fairness (latency). For batch processing where latency doesn't matter, you might strictly prefer cooperative scheduling, but you can't turn off preemption in Go. The 10ms timeslice is hardcoded.

---

### 707. What are the stages of Go's garbage collector?
"The GC cycle has four main phases:
1.  **Sweep Termination**: A short Stop-The-World (STW) pause to finish the previous cycle.
2.  **Marking**: Concurrent. The GC scans memory while the app runs (Write Barriers are on).
3.  **Mark Termination**: A second STW pause to finish marking and closing the barriers.
4.  **Sweeping**: Concurrent. The runtime reclaims memory in the background.

My mental model is that the 'Marking' phase burns CPU (stealing about 25% of cycles), but keeps the app responsive. The STW pauses are practically invisible for most web workloads."

#### Indepth
**Mark Assist**. If the app allocates memory faster than the GC can mark it, the GC forces the allocating goroutine to *help* mark. This slows down the allocation, applying backpressure. This prevents the app from outrunning the GC and causing an OOM.

---

### 708. What is the role of the `runtime` package?
"The `runtime` package is the engine that powers Go. It is linked into every binary.

It handles:
1.  **Memory Allocation** (`mallocgc` relies on `tcmalloc` principles).
2.  **Scheduling** (`proc.go` manages M, P, G).
3.  **Garbage Collection** (`mgc.go`).
4.  **Stack Management** (growing/shrinking stacks).

I never import `runtime` for business logic (except `GOMAXPROCS`), but understanding it helps me debug performance issues—knowing *why* a stack split happens or *why* the GC is thrashing."

#### Indepth
`runtime.KeepAlive(x)`. This function is critical when using `unsafe` or `SetFinalizer`. It tells the GC "Consider object X reachable up to this point". Without it, the aggressive compiler might notice X isn't used anymore *before* the function ends, and collect it while you are still using its raw pointer in a C call.

---

### 709. How does Go handle stack traces?
"When a panic occurs, the runtime walks the stack frames of the current goroutine.

It uses metadata generated by the compiler (the **PC-Value table**) to map the raw Instruction Pointer (PC) to the file name, line number, and function name.

This is why stripped binaries are harder to debug—if that symbol table is gone, the runtime can't print the nice `main.go:42` message. I rely on these stack traces heavily; they are usually the only clue needed to fix a crash."

#### Indepth
**Panic Parsing**. Many log aggregators (Datadog/Sentry) have special parsers for Go stack traces. They group "Error at line 42" and "Error at line 43" as different issues. Consistent, unstripped stack traces are vital for error grouping. Use `-ldflags="-s -w"` only for production binaries where size matters more than debuggability (or use creating sourcemaps).

---

### 710. How does the Go runtime manage memory allocation?
"Go’s allocator is based on **TCMalloc** (Thread-Caching Malloc). It prioritizes low lock contention.

It acts as a hierarchy:
1.  **mcache**: A per-P (per-thread) cache. Allocating small objects here is lock-free and extremely fast.
2.  **mcentral**: A shared list of spans (pages). Requires a lock.
3.  **mheap**: The global heap that requests memory from the OS (`mmap`).

This means mostly, allocation is just bumping a pointer in the local cache. That’s why I don't shy away from creating short-lived structs—it’s cheap."

#### Indepth
**Tiny Allocator**. For objects < 16 bytes (like `bool`, `int`), Go packs them together into a single 16-byte block. This reduces fragmentation and improves cache locality. It is a specific optimization inside `mallocgc`.

---

### 711. What’s the difference between a green thread and a goroutine?
"Conceptually, they are similar (user-space threads), but Goroutines are an evolution.

Classic Green Threads (like in early Java) often had fixed stack sizes (e.g., 64KB). If you allocated 100k of them, you ran out of RAM even if they were doing nothing.

**Goroutines** utilize **segmented stacks** (now contiguous, resizeable stacks). They start at 2KB. This dynamic sizing is the key difference that allows Go to support *millions* of concurrent routines on a single machine where other runtimes failed."

#### Indepth
**Context Switching Cost**. Switching between OS threads takes ~1-2 microseconds (kernel involved). Switching between Goroutines takes ~200 nanoseconds (user space). This 10x difference is why Go's concurrency model is superior for high-frequency switching tasks (like network servers).

---

### 712. What are finalizers in Go and how do they work?
"A finalizer is a function attached to an object via `runtime.SetFinalizer`. The GC calls it when the object is about to be collected.

**I strongly avoid them.**

They are unpredictable. There is no guarantee *when* they will run, or if they will run at all (e.g., if the program exits). Relying on them for cleanup (like closing a file) leads to resource leaks. I always use explicit `Close()` methods or `defer`."

#### Indepth
**Resurrection**. A common bug with Finalizers is "resurrecting" the object. If the finalizer assigns the object to a global variable, the object becomes reachable again! The GC has to track this state, making finalizers expensive and complicated. Just don't use them.

---

### 713. What is the role of `goexit` internally?
"`runtime.Goexit()` terminates the goroutine that calls it.

Unlike `return`, it executes all deferred calls immediately before quitting. It does *not* stop the program (like `os.Exit`), just the goroutine.

It’s actually how the runtime handles the end of a normal goroutine function (it implicitly calls `goexit`). I've only used it manually in niche testing scenarios where I needed to fail a test helper without panicking."

#### Indepth
**Panic vs Goexit**. `panic` crashes the program if not recovered. `Goexit` silently terminates the goroutine. If the main goroutine calls `Goexit`, the program continues running until all *other* goroutines finish (or crash/deadlock). It's a weird state, mostly for internal runtime use.

---

### 714. How does Go avoid stop-the-world pauses?
"It minimizes them, but doesn't completely avoid them. It does this by making the most expensive part—**Marking**—concurrent.

The challenge is: 'What if the app changes a pointer while the GC is scanning it?'
Go solves this with a **Hybrid Write Barrier**. It intercepts every pointer write during the GC cycle and makes sure the new object is marked/colored correctly, preserving the 'Tri-color Invariant'. This allows the world to keep moving while the GC cleans up."

#### Indepth
**Barrier Overhead**. The Write Barrier adds a small overhead to *every* pointer assignment in your code (checking flags, maybe coloring). This is why code with tons of pointer mutations runs slightly slower when GC is active. Stack writes do not have barriers (for performance), which necessitates the "Stack Rescan" (STW) phase.

---

### 715. How does memory fragmentation affect Go programs?
"Fragmentation happens when there is free memory, but it's in small, scattered chunks that can't be used for a large allocation.

Go uses **Size Classes** (allocating objects into fixed-size buckets like 8B, 16B, 32B) to minimize *external* fragmentation.

However, Go creates *virtual memory fragmentation*. The runtime might hold onto a large virtual address space even if physical RAM usage is low. This is why `top` sometimes shows huge VIRT usage for Go apps—it’s usually benign, but can be confusing."

#### Indepth
`MADV_FREE` vs `MADV_DONTNEED`. Go releases memory to the OS. On Linux, it uses `MADV_FREE` (lazy release). The OS shows the memory as "used" until it effectively needs it. This caused a lot of "Go has a memory leak!" complaints. Go 1.16+ reverted to `MADV_DONTNEED` in some cases to make RSS metrics look more accurate to users.

---

### 716. What’s the meaning of “non-preemptible” code in Go?
"Non-preemptible code is code that the scheduler cannot pause.

In older Go versions, a tight math loop `for { x++ }` was non-preemptible because it made no function calls. It could freeze the GC and the Scheduler.

Today, thanks to asynchronous preemption, very little code is truly non-preemptible, except for **cgo** calls and some atomic operations. If I call a C function that sleeps for an hour, the Go scheduler can't touch that thread until it returns."

#### Indepth
`runtime.LockOSThread()`. You can manually make a goroutine "non-preemptible" in the sense that it is bound to a specific OS thread. This is mandatory for libraries that use Thread-Local Storage (TLS) or GUI frameworks (OpenGL) that require all calls to happen from the "Main Thread".

---

### 717. What are M:N scheduling models and how does Go implement it?
"M:N provides a middle ground between 1:1 (Native Threads) and N:1 (Event Loop / Async).

Go implements it by mapping **N** Goroutines onto **M** Kernel Threads.
This allows Go to handle IO blocking efficiently. If a Goroutine blocks on a file read, the Runtime parks it (N side) but keeps the Kernel Thread (M side) busy running other Goroutines.

This abstraction enables synchronous-looking code (`file.Read`) to behave asynchronously under the hood."

#### Indepth
**Blocking Syscalls**. Not all syscalls are async. If a G does a blocking syscall (like `chmod` or `getpakname` on some OSs), the M serves it and blocks. The P detaches (handoff) and starts a *new* M to run other Gs. This is why a Go program might spawn 1000 OS threads if you do 1000 blocking syscalls simultaneously.

---

### 718. How does Go detect deadlocks at runtime?
"The runtime has a global detector. If **all** goroutines are in a waiting state (sleeping, waiting on channel, Waiting on mutex) and no system/network polling is active, it knows progress is impossible.

It panics with `fatal error: all goroutines are asleep - deadlock!`.

Note that it only detects **global** deadlocks. If 2 goroutines are deadlocked but a 3rd is happily running a web server, the runtime won't catch it. That’s why I rely on timeouts."

#### Indepth
**Dumping Goroutines**. To debug a partial deadlock, send `SIGQUIT` (Ctrl+\) to the Go process. It dumps the stack trace of *all* goroutines to stdout. You can then see exactly which two are stuck waiting on each other's locks. This is a builtin runtime feature.

---

### 719. What are the internal states of a goroutine?
"A goroutine moves through states like a state machine:
1.  **_Gidle**: Just initialized.
2.  **_Grunnable**: Sitting in a run queue (Local P or Global), waiting for a CPU.
3.  **_Grunning**: Currently executing on an M.
4.  **_Gwaiting**: Blocked (waiting for channel, mutex, or IO).
5.  **_Gdead**: Finished execution.

Understanding this helps when reading `go tool trace`—seeing 10,000 goroutines in `_Grunnable` means my CPU is saturated."

#### Indepth
**_Gcopying**. There is a transient state during stack growth. The G is paused while its stack is copied to a new location. You rarely see this unless you are really digging into the scheduler internals or have extremely deep recursion issues.

---

### 720. What is `go:linkname` and when is it used?
"`//go:linkname` is a complier directive that acts like a wormhole. It links a local symbol to a private (unexported) symbol in another package.

For example, `time.Sleep` uses it to call private runtime functions.

It is **dangerous**. It bypasses the type system and modularity. I strictly avoid it in application code because it relies on internal implementation details that can change in any Go release, breaking my build."

#### Indepth
**Standard Lib Usage**. You see this often in `sync/atomic` or `syscall`. They link to assembly implementations or runtime intrinsics. It is the bridge between "Go Code" and "Runtime Magic". Unless you are writing an alternative compiler or a debugger, you have no reason to use it.
