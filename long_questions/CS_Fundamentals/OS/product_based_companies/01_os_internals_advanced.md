# 🖥️ OS Internals — Advanced Interview Questions (Product-Based Companies)

This document covers advanced OS internals for product-based company interviews (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay). Targeted at 3–10 years of experience rounds.

---

### Q1: How does the Linux kernel handle system calls? What is the cost of a syscall?

**Answer:**
A **system call** is the interface between user-space code and the kernel. User programs cannot directly access hardware — they must ask the kernel.

**System call flow (x86-64 Linux):**
1. User program puts syscall number in `rax` register, arguments in `rdi`, `rsi`, `rdx`, etc.
2. Executes `syscall` instruction.
3. CPU switches from **Ring 3 (user mode)** to **Ring 0 (kernel mode)** — privilege level change.
4. CPU saves user-space registers, switches to kernel stack.
5. Kernel executes the syscall handler (e.g., `sys_read`, `sys_write`).
6. Kernel sets return value in `rax`, executes `sysret`.
7. CPU switches back to Ring 3, restores registers.
8. User program resumes.

**Cost of a syscall:**
- Typically **~100-300 nanoseconds** for a simple syscall (like `getpid()`).
- Breakdown: ~50ns mode switch + kernel execution + mode switch back.
- **Meltdown/Spectre patches (2018)**: Added Page Table Isolation (PTI/KPTI) — flushing TLB on every kernel entry/exit. Added **~100-400ns overhead** per syscall on affected systems.

**Minimizing syscall overhead:**
- **io_uring**: Submit multiple I/O operations in batches via shared ring buffers — drastically fewer syscalls.
- **vDSO (Virtual Dynamic Shared Object)**: Maps frequently called, read-only syscalls (like `gettimeofday`) into user space. Executes without mode switch at all.
- **Batch operations**: Use `sendfile`, `splice`, vectorized I/O (`readv`/`writev`).

---

### Q2: Explain memory-mapped files (mmap). How is it used in databases and message queues?

**Answer:**
**`mmap()`** maps a file (or device) into the process's virtual address space. Reading/writing to that memory region directly reads/writes the file — without explicit `read()`/`write()` syscalls.

**How mmap works:**
1. `mmap(addr, len, PROT_READ|PROT_WRITE, MAP_SHARED, fd, offset)` creates a virtual memory mapping.
2. Initially, the pages are not loaded into RAM.
3. First access to a page triggers a **page fault** — OS loads the page from disk.
4. Subsequent accesses are served from RAM (no syscall needed).
5. On `munmap()` or process exit, dirty pages are written back to disk.

**Advantages:**
- Zero-copy reads: Data goes from disk cache → user space without kernel→user copy.
- OS manages caching automatically via page cache.
- Enables sharing between processes (`MAP_SHARED`).

**Used in production systems:**
- **Kafka**: Log segments are mmaped — producer writes to virtual memory, OS manages when to flush to disk. Enables high-throughput sequential writes.
- **RocksDB/LevelDB (LSM stores)**: Use mmap for SSTable reads.
- **SQLite**: WAL (Write-Ahead Log) mode uses mmap for readers.
- **Elasticsearch/Lucene**: mmap for inverted index segments.
- **Shared memory IPC**: Multiple processes mapping the same file/shmem object.

**Disadvantages:**
- Page faults for large files cause latency spikes.
- Difficult to reason about when writes are durable (need `msync()`).
- 32-bit systems: limited address space.

---

### Q3: What is the difference between user-space and kernel-space threading? Explain Go's M:N threading model.

**Answer:**

**Kernel Threads:**
- Managed by the OS kernel.
- Each thread has its own OS-level thread (1:1 mapping).
- Context switch goes through kernel — expensive (~5-10μs).
- Examples: Java threads (mapped 1:1 to OS threads prior to Loom).

**User-space Threads (Green Threads / Coroutines / Goroutines):**
- Managed by user-space runtime — kernel doesn't know about them individually.
- Much lighter: context switch is cooperative, just registers save/restore in user space (~100ns).
- Examples: Go goroutines (initially), Python gevent, old Java green threads.

**Go's M:N Model:**
- **G (Goroutine)**: User-space coroutine. Can be millions of them.
- **M (Machine)**: OS kernel thread. Typically = `GOMAXPROCS` (number of CPU cores).
- **P (Processor)**: Logical processor that holds runqueues. GOMAXPROCS = number of Ps.

```
G G G G G G G ← many goroutines (M:N)
    ↓  ↓
   P0  P1     ← each P holds a local run queue
    |   |
   M0  M1     ← OS threads (bound to CPU cores)
```

**Key features:**
- **Work stealing**: Idle P steals goroutines from busy P's runqueue.
- **Non-blocking syscalls**: When a goroutine blocks on I/O, the runtime parks it, and the M moves to run another goroutine. A separate thread pool handles the actual blocking syscall.
- **Stack growth**: Goroutines start with ~2KB stack, grow dynamically (vs 8MB fixed OS thread stack).

---

### Q4: Explain Copy-on-Write (CoW) forking. How does it work and where is it used?

**Answer:**
**Copy-on-Write (CoW)** is an optimization where two processes share the same physical memory pages after a `fork()`. Pages are only physically copied when one process **writes** to them.

**Without CoW:**
`fork()` would need to copy the entire parent address space immediately — prohibitively expensive for large processes.

**With CoW:**
1. `fork()` creates a new process that shares all pages with the parent (marked read-only, copy-on-write).
2. Both parent and child read from the same physical pages.
3. When EITHER writes to a page:
   - **Page fault** triggered (write to read-only page).
   - OS copies that specific page to a new physical frame.
   - Updates the writing process's page table to point to the new copy.
   - Both processes now have independent copies of that page.
4. Pages never written remain shared — zero copy cost.

```
BEFORE write:
Parent PTE → Physical Page A ← Child PTE (shared, read-only)

AFTER write by child:
Parent PTE → Physical Page A
Child PTE  → Physical Page B (copy of A, modified)
```

**Used in:**
- **Redis `BGSAVE`**: Forks a child to write snapshot to disk. While child writes RDB file, parent continues serving writes using CoW — only modified pages are duplicated.
- **Process creation on Unix**: Every `fork()` + `exec()` uses CoW.
- **Container runtimes**: Docker layers use overlay filesystems with CoW semantics.
- **Programming langes**: Python's multiprocessing, Ruby's Unicorn web server.

---

### Q5: What is the OOM killer in Linux? How does it decide which process to kill?

**Answer:**
When the Linux system runs out of physical memory and swap, the **OOM (Out-Of-Memory) Killer** is invoked to free memory by killing a process.

**When OOM killer triggers:**
- Memory allocation fails and memory reclaim (page eviction, swap) cannot satisfy it.
- Linux tries to reclaim pages aggressively first (evict page cache, swap out anonymous pages).
- If still not enough → OOM killer runs.

**OOM score calculation:**
Each process has an **`oom_score`** (0–1000). Higher score = more likely to be killed.

Components of oom_score:
- **Memory usage**: Percentage of physical RAM the process uses (major factor).
- **Running time**: Longer-running processes get slightly lower score (children of long-running processes protected).
- **`oom_score_adj`**: User-adjustable (-1000 to +1000). `-1000` = never kill; `+1000` = kill first.

**Viewing OOM scores:**
```bash
cat /proc/<PID>/oom_score
cat /proc/<PID>/oom_score_adj
```

**Protecting critical processes:**
```bash
echo -1000 > /proc/<PID>/oom_score_adj  # Never kill this process (root only)
```

**In practice:**
- Kubernetes sets `oom_score_adj` based on **QoS class**:
  - `Guaranteed` pods: oom_score_adj = -997 (highly protected)
  - `BestEffort` pods: oom_score_adj = 1000 (killed first)
  - `Burstable` pods: oom_score_adj = 2-999 (based on memory limit ratio)

---

### Q6: What is io_uring and how does it differ from epoll for high-performance I/O?

**Answer:**

**epoll (Linux 2.6+):**
- **Event-driven** I/O: Register file descriptors (sockets, pipes), get notified when they're ready for read/write.
- Each I/O operation still requires separate `read()`/`write()` syscalls.
- Per-operation syscall overhead adds up at high RPS.

**io_uring (Linux 5.1+):**
- **Submission/Completion queue pairs** in shared memory (ring buffers) between kernel and user space.
- User prepares I/O operations in **Submission Queue (SQ)** — then calls one syscall (`io_uring_enter()`) to submit all of them (or zero syscalls in polling mode).
- Kernel completes the operations and puts results in **Completion Queue (CQ)** — user polls the CQ.

**Comparison:**

| Feature | epoll | io_uring |
|---|---|---|
| Mechanism | Readiness notification | Completion notification |
| Syscalls per I/O | ~2-3 (epoll_wait + read + write) | ~0 (with SQPOLL kernel thread) |
| Supported ops | Network I/O only | Network + File I/O + custom ops |
| Zero-copy | No | Yes (with registered buffers) |
| Kernel version | 2.6+ | 5.1+, full features in 5.10+ |

**Real-world adoption:**
- **Tokio (Rust async runtime)**: Moving from epoll to io_uring.
- **Node.js libuv**: Evaluating io_uring.
- **RocksDB**: io_uring for file I/O.
- **PostgreSQL 16+**: io_uring for WAL writes.

---

*Prepared for technical rounds at product-based companies (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay, Groww).*
