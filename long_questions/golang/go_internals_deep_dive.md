# 🧠 Go Internals Deep Dive

This section helps you answer "How does X work under the hood?" questions. These are **critical** for L5/L6 (Senior/Staff) roles at FAANG and top product companies.

---

## 🗺️ 1. Map Internals

### The `hmap` Struct
A Go map is a pointer to a `hmap` struct in the runtime.
```go
type hmap struct {
    count     int    // Live cells = len(map)
    flags     uint8
    B         uint8  // log_2 of # of buckets (2^B buckets)
    noverflow uint16 // Approx number of overflow buckets
    hash0     uint32 // Hash seed
    buckets   unsafe.Pointer // Array of 2^B buckets
    oldbuckets unsafe.Pointer // Previous buckets (used during resizing)
    ...
}
```

### Bucket Structure
Each bucket holds **8 key/value pairs**.
*   **Top Hash (tophash)**: First 8 bits of the hash. Used for fast comparison inside a bucket.
*   **Keys**: Stored sequentially (`key0, key1...`).
*   **Values**: Stored sequentially after keys (`val0, val1...`). *Why? To eliminate padding and save memory.*
*   **Overflow Pointer**: If a bucket is full ( > 8 keys), it points to an **overflow bucket**.

### Lookup Process
1.  Hash the Key.
2.  **Low-Order Bits (LOB)** determine the **Bucket Index**.
3.  **High-Order Bits (HOB)** are compared against **tophash** array inside the bucket.
4.  If tophash matches, compare the actual key.

### Evacuation (Resizing)
*   **Trigger**: When `Load Factor > 6.5` (Average items per bucket).
*   **Growth**: Size doubles.
*   **Incremental**: Data is NOT copied all at once (that would freeze the app). Instead, when you "Access" (Assigment/Delete) a key, 2 buckets are moved from `oldbuckets` to `buckets`.

---

## 📡 2. Channel Internals

### The `hchan` Struct
A channel is a pointer to a `hchan` struct. It is **thread-safe** by design.
```go
type hchan struct {
    qcount   uint           // total data in queue
    dataqsiz uint           // size of circular queue
    buf      unsafe.Pointer // points to an array (Circular Queue)
    elemsize uint16
    closed   uint32
    recvq    waitq          // list of recv waiters
    sendq    waitq          // list of send waiters
    lock     mutex          // Protects all fields
}
```

### Buffered Channel (Circular Queue)
*   Uses a **Ring Buffer** (array) to store elements.
*   **Send**: Acquires lock -> Copies data to `buf` -> Release lock.
*   **Receive**: Acquires lock -> Copies data from `buf` -> Release lock.

### Direct Send/Receive (Unbuffered / Bypass)
*   **Optimization**: If a receiver is already waiting, the sender **writes directly** to the receiver's stack memory!
*   **Sudog**: The struct representing a waiting goroutine in `recvq` or `sendq`.

---

## ⚙️ 3. The Scheduler (GMP Model)

Go uses an **M:N Scheduler** (Runs M OS threads on N Cores, multiplexing thousands of Goroutines).

### The Components
1.  **G (Goroutine)**:
    *   Lightweight (Starts at 2KB stack).
    *   Contains Stack Pointer (SP), Program Counter (PC), and Status (Runnable, Waiting).
2.  **M (Machine)**:
    *   Real **OS Thread**.
    *   Executed by the OS.
3.  **P (Processor)**:
    *   Logical Context required to run Go code.
    *   Has a **Local Run Queue** of Gs.
    *   Default `GOMAXPROCS` = Number of CPU Cores.

### Execution Flow
*   `M` must acquire a `P` to run `G`.
*   `M` picks a `G` from `P`'s **Local Run Queue**.
*   If Local Queue is empty, `M` tries to **Steal** work (half the Gs) from another `P`'s queue. (Work Stealing).
*   If all empty, it checks the **Global Run Queue** (Checked occasionally to prevent starvation).

### Preemption (Cooperative & Asynchronous)
*   **Old**: Cooperative. Occurred at function calls. Loops could hang the scheduler.
*   **New (Go 1.14+)**: **Asynchronous Preemption**. Sys signals (`SIGURG`) interrupt a goroutine running > 10ms.

---

## 🔌 4. Interfaces Internals

### `eface` (Empty Interface `interface{}`)
Used for `var x interface{} = ...`.
```go
type eface struct {
    _type *_type          // Information about the dynamic type (int, string, User...)
    data  unsafe.Pointer  // Pointer to the actual data
}
```

### `iface` (Non-Empty Interface)
Used for `var r Shape` (where `Shape` has methods).
```go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
```
*   **itab (Interface Table)**: Caches the method pointers.
    *   Contains the specific `type` (e.g., `*Square`).
    *   Contains list of function pointers (e.g., `Area()`, `Perimeter()`) specialized for that type.
    *   *Significance*: Calling a method on an interface is a **dynamic dispatch** (O(1) lookup in itab table), slightly slower than direct call but very fast.

---

## ✂️ 5. Slices Internals

### The `slice` Header
A slice is just a small struct passed by value.
```go
type slice struct {
    array unsafe.Pointer // Pointer to the underlying backing array
    len   int            // Number of elements accessible
    cap   int            // Capacity allocated
}
```
*   **Grow**: When `append` exceeds `cap`, Go allocates a **new** larger array (usually 2x), copies data, and updates the pointer.
*   **Gotcha**: Review "Sub-slice Memory Leak" (keeping a reference to a small part of a huge array keeps the huge array alive in GC).
