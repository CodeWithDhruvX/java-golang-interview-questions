# 🖥️ OS Fundamentals — Interview Questions (Service-Based Companies)

This document covers Operating Systems concepts commonly tested at service-based companies like TCS, Infosys, Wipro, Capgemini, HCL, Cognizant. Targeted at 1–5 years of experience rounds.

---

### Q1: What is the difference between a Process and a Thread?

**Answer:**

| Feature | Process | Thread |
|---|---|---|
| Definition | Independent program in execution | Lightweight unit of execution within a process |
| Memory | Each has its own memory space (heap, stack, code) | Shares the process's memory (heap, code); own stack |
| Communication | IPC (pipes, sockets, shared memory) — complex | Shared memory — easy but risky |
| Creation cost | Heavy (fork/exec — duplicates address space) | Lightweight (just a new stack + registers) |
| Failure isolation | Process crash doesn't affect others | Thread crash can crash the entire process |
| Context switch | Expensive (switch entire memory map) | Cheaper (same address space) |
| Example | Two `java` programs running | Two threads inside one Spring Boot app |

**Use-case:**
- Use **processes** for full isolation (web server workers, separate microservices).
- Use **threads** for parallelism within a single application (handling multiple HTTP requests in one JVM, Go goroutines internally use OS threads).

---

### Q2: What is a context switch? What is the overhead involved?

**Answer:**
A **context switch** is when the CPU stops executing one process/thread and starts executing another.

**Steps in a context switch:**
1. **Save state**: CPU registers, program counter, stack pointer of the current process → saved to its PCB (Process Control Block).
2. **Update OS data structures**: Scheduler marks the current process as Waiting/Ready.
3. **Select next process**: Scheduler picks the next process to run.
4. **Restore state**: Load the next process's CPU state from its PCB.
5. **TLB flush** (for process switch): Translation Lookaside Buffer is invalidated — next memory access pays a page walk penalty.
6. **Resume execution** from where the new process was paused.

**Overhead:**
- Direct cost: ~1–10 microseconds for saving/restoring registers.
- Indirect cost: Cache and TLB invalidation — newly scheduled process finds its data evicted from CPU cache (cache miss penalty can be 100–200 clock cycles per miss).

**Reducing context switches:**
- Use thread pools (avoid creating too many threads).
- Use non-blocking/async I/O (event loop like Nginx, Node.js) — fewer blocking waits.
- Batch work to run longer without yielding.

---

### Q3: What is a deadlock? What are the four conditions for a deadlock?

**Answer:**
A **deadlock** is a state where two or more processes are stuck, each waiting for a resource that another holds — none can proceed.

**Classic example:**
```
Process A holds Lock_1, waits for Lock_2
Process B holds Lock_2, waits for Lock_1
→ Both stuck forever
```

**Coffman's 4 conditions** (ALL must hold for deadlock):

| Condition | Meaning |
|---|---|
| **Mutual Exclusion** | Resources cannot be shared; only one process at a time |
| **Hold and Wait** | Process holds at least one resource AND waits for others |
| **No Preemption** | Resources cannot be forcibly taken from a process |
| **Circular Wait** | Process A waits for B, B waits for C, C waits for A (cycle) |

**Prevention strategies** (break one condition):
- **Lock ordering**: Always acquire locks in the same order across all threads → breaks Circular Wait.
- **Timeout**: If lock not acquired in N ms, release all held locks and retry → breaks Hold and Wait partially.
- **Deadlock detection**: Allow deadlocks, detect via cycle detection in resource allocation graph, kill a process.
- **Try-lock pattern**: `tryLock()` with timeout → non-blocking.

---

### Q4: What is virtual memory? What is paging and what is a page fault?

**Answer:**
**Virtual Memory** is an abstraction that gives each process the illusion that it has a large, contiguous block of memory, even if physical RAM is limited or fragmented.

**How it works:**
- Each process has a **virtual address space** (e.g., 0 to 2^48 bytes on 64-bit Linux).
- The OS/CPU maps virtual addresses → physical addresses via a **Page Table**.
- Physical RAM is divided into fixed-size blocks called **frames** (typically 4KB).
- Virtual memory is divided into same-size **pages**.

**Paging:**
- Virtual pages are mapped to physical frames.
- Not all pages need to be in RAM — infrequently used pages are stored on disk (swap space).
- The **TLB (Translation Lookaside Buffer)** caches recent virtual→physical address translations for speed.

**Page Fault:**
Occurs when a process accesses a virtual address whose page is NOT currently in physical RAM.

1. CPU triggers page fault exception.
2. OS page fault handler runs.
3. OS finds the page on swap disk.
4. Loads it into a free physical frame.
5. Updates page table.
6. Resumes the instruction.

**Types:**
- **Minor page fault**: Page is in memory but not mapped yet (no disk I/O needed).
- **Major page fault**: Page must be loaded from disk — expensive (milliseconds).

---

### Q5: What are the common CPU scheduling algorithms?

**Answer:**

| Algorithm | Description | Pros | Cons |
|---|---|---|---|
| **FCFS** (First Come First Serve) | Run in arrival order | Simple | Convoy effect — short jobs wait behind long ones |
| **SJF** (Shortest Job First) | Run shortest burst time next | Optimal avg wait time | Requires knowing burst time in advance (impractical) |
| **Round Robin (RR)** | Each process gets a time quantum (e.g., 10ms), then preempted | Fair, good for time-sharing | High context switches if quantum too small |
| **Priority Scheduling** | Run highest priority process first | Allows urgent tasks | Starvation of low-priority processes |
| **MLFQ** (Multi-Level Feedback Queue) | Multiple queues with different priorities + time quantum; demotes CPU-heavy processes | Best for interactive systems, adapts dynamically | Complex |

**Real-world:**
- Linux uses **CFS (Completely Fair Scheduler)** — gives each process a fair share of CPU time using a red-black tree ordered by "virtual runtime" (time spent on CPU).
- Windows uses a **priority-based preemptive scheduler**.

---

### Q6: What is the difference between mutex, semaphore, and monitor?

**Answer:**

| Feature | Mutex | Semaphore | Monitor |
|---|---|---|---|
| Ownership | Only the acquiring thread can release | Any thread can release | Encapsulated ownership |
| Counter | Binary (0 or 1) | Integer (0 to N) | Binary (like mutex) |
| Use case | Mutual exclusion (one thread at a time) | Controlling access to N resources | High-level synchronization construct |
| Language support | POSIX pthread_mutex | POSIX sem_t | Java `synchronized`, Python `threading.Lock()` |

**Mutex:**
```java
mutex.lock();
// critical section
mutex.unlock();
```

**Semaphore (counting):**
- A semaphore with value=3 allows 3 threads simultaneously (e.g., a pool of 3 database connections).
- `wait()` decrements (blocks if 0). `signal()` increments.

**Monitor:**
- Object that combines a mutex + condition variable.
- Java `synchronized` blocks + `wait()`/`notify()` = monitor.
- Higher-level abstraction — easier to use correctly.

---

*Prepared for technical screening at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, Tech Mahindra).*
