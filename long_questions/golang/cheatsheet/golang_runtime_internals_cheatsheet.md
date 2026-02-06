# Golang Runtime & Internals Cheatsheet

The "Black Box" knowledge typically asked in Senior/Staff engineer interviews. Use this to explain *why* Go is fast.

## 🟢 The Scheduler (GMP Model)

Go uses a **M:N scheduler** (M OS threads run N Goroutines).

### 1. The Components
- **G (Goroutine):** A lightweight thread (starts at ~2KB stack). Contains stack, instruction pointer, and scheduling info.
- **M (Machine):** An OS Thread. The actual worker executing code on the CPU.
- **P (Processor):** A resource required to execute Go code. Contains a **Local Run Queue** of Gs.
  - Default `GOMAXPROCS` = Number of CPU Cores.

### 2. How it Works
1. **M** must acquire a **P** to run a **G**.
2. **P** holds a queue of runnable **G**s.
3. **M** picks a **G** from **P**'s queue and executes it.

### 3. Key Optimizations
- **Work Stealing:** If a P's queue is empty, it "steals" half the Gs from another P's queue (randomly). Keeps all cores busy.
- **Hand Off:** If a G makes a blocking syscall (e.g., file I/O), the M blocks. The P detaches and moves to a distinct (new or idle) M to keep running other Gs.
- **Preemption:** Long-running Gs (loops) are paused (~10ms) to prevent starvation (Async Preemption in Go 1.14+).

---

## 🟡 Memory Management

### 1. Stack vs Heap (Escape Analysis)
Go decides where to allocate memory at compile time.
- **Stack:** Fast, strictly LIFO. Automatically cleaned up when function returns.
- **Heap:** Slower, GC managed. Needed for data that lives longer than the function scope.

**Escape Analysis Rules:**
1. **Returns Pointer:** If you return a pointer to a local variable, it **escapes to heap**.
2. **Dynamic Size:** Slices with dynamic size (`make([]int, n)`) often escape.
3. **Interface Assignment:** Variables assigned to `interface{}` often escape.

```bash
# Check escape analysis
go build -gcflags="-m" main.go
```

### 2. Allocation Hierarchy (TCMalloc style)
1. **mcache:** Per-P (Processor) cache. **No lock needed**. Fast allocation for small objects.
2. **mcentral:** Shared central list of spans. Requires locking.
3. **mheap:** Global heap. Requests memory from OS.

---

## 🔴 Garbage Collection (GC)

Go uses a **Concurrent Tri-Color Mark & Sweep** collector.

### 1. The Phases
1. **Mark Setup (STW):** Turn on Write Barriers. Brief "Stop The World" (~10-30µs).
2. **Marking (Concurrent):**
   - **White:** Potential garbage.
   - **Grey:** Active, but children not scanned.
   - **Black:** Active and fully scanned.
   - GC marks roots (stacks, globals) as Grey, then traverses pointers turning them Black.
3. **Mark Termination (STW):** Clean up remaining work.
4. **Sweep (Concurrent):** Reclaim White objects (memory).

### 2. Write Barrier (Hybrid)
To prevent the "user program" (Mutator) from hiding an object from the GC while it runs:
- If you move a pointer, the Write Barrier ensures the object is marked Grey (reachable) so it isn't accidentally deleted.

### 3. Tuning (GOGC)
- Default `GOGC=100`.
- GC triggers when heap grows by 100% of its size after the last GC.
- `GOGC=200`: GC less often (more memory usage).
- `GOGC=50`: GC more often (less memory usage, more CPU).
- `GOMEMLIMIT` (Go 1.19+): Set a "Soft" memory limit to trigger GC more aggressively as you near OOM.

---

## 🔵 Goroutines vs Threads

| Feature | OS Thread | Goroutine |
| :--- | :--- | :--- |
| **Size** | ~1MB (Fixed) | ~2KB (Dynamic growable) |
| **Creation/Teardown** | Slow (Syscalls) | Fast (User space) |
| **Switching Cost** | High (Context switch) | Low (Pointer swap) |
| **Identity** | Has ID | No exposed ID |
| **Scheduling** | By OS Kernel | By Go Runtime |
