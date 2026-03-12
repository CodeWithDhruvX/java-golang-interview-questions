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

### Explanation
The hmap struct is the runtime representation of a Go map. It contains metadata like count, bucket count (B), hash seed, and pointers to bucket arrays. The oldbuckets pointer enables incremental resizing without blocking.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does a Go map work internally?
**Your Response:** "A Go map is actually a pointer to an hmap struct in the runtime. This struct tracks the number of elements, bucket configuration, and hash seed. The B field determines how many buckets there are - 2^B buckets total. The buckets field points to the actual bucket array, and oldbuckets is used during resizing. The hash0 field is a random seed that makes hash attacks harder. When I create a map, Go allocates this hmap structure and initializes the buckets. The map variable I use in Go code is just a pointer to this runtime structure. This design allows maps to grow dynamically and handle collisions efficiently."

### Bucket Structure
Each bucket holds **8 key/value pairs**.
*   **Top Hash (tophash)**: First 8 bits of the hash. Used for fast comparison inside a bucket.
*   **Keys**: Stored sequentially (`key0, key1...`).
*   **Values**: Stored sequentially after keys (`val0, val1...`). *Why? To eliminate padding and save memory.*
*   **Overflow Pointer**: If a bucket is full ( > 8 keys), it points to an **overflow bucket**.

### Explanation
Map buckets store 8 key/value pairs efficiently. Keys and values are stored separately to minimize memory padding. The tophash array stores the first 8 bits of each key's hash for fast initial comparison before full key comparison.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How are map buckets organized?
**Your Response:** "Each map bucket holds exactly 8 key/value pairs. The keys are stored sequentially, followed by the values - this layout eliminates memory padding that would occur if I interleaved them. Each bucket also has a tophash array containing the first 8 bits of each key's hash. When looking up a key, Go first compares the tophash to quickly eliminate non-matches before doing the full key comparison. If a bucket gets more than 8 elements due to hash collisions, it creates an overflow bucket linked to the original. This design balances memory efficiency with lookup speed - the tophash provides a fast filter, and the sequential storage minimizes wasted space."

### Lookup Process
1.  Hash the Key.
2.  **Low-Order Bits (LOB)** determine the **Bucket Index**.
3.  **High-Order Bits (HOB)** are compared against **tophash** array inside the bucket.
4.  If tophash matches, compare the actual key.

### Explanation
Map lookup uses the hash value's low-order bits to select the bucket, then high-order bits (tophash) for fast comparison within the bucket. Only if tophash matches does Go perform the full key comparison.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does map lookup work internally?
**Your Response:** "When I look up a key in a map, Go first hashes the key. The low-order bits of the hash determine which bucket to check - this ensures even distribution. Inside the bucket, Go compares the high-order bits (stored in tophash) to quickly find potential matches. If the tophash matches, only then does Go do the full key comparison. This two-step process is efficient - the tophash comparison eliminates most non-matches quickly. If the key isn't found in the main bucket, Go checks any overflow buckets. The hash seed ensures different runs of the program produce different hash values, preventing hash collision attacks. This design gives O(1) average lookup time while handling collisions gracefully."

### Evacuation (Resizing)
*   **Trigger**: When `Load Factor > 6.5` (Average items per bucket).
*   **Growth**: Size doubles.
*   **Incremental**: Data is NOT copied all at once (that would freeze the app). Instead, when you "Access" (Assigment/Delete) a key, 2 buckets are moved from `oldbuckets` to `buckets`.

### Explanation
Map resizing (evacuation) is triggered when the load factor exceeds 6.5. The map doubles in size, but evacuation is incremental - only 2 buckets are moved per access to avoid blocking the application during resize.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does map resizing work?
**Your Response:** "Maps resize when the load factor exceeds 6.5, meaning there are too many elements per bucket. When this happens, the map doubles in size. But crucially, Go doesn't copy all the data at once - that would freeze the application. Instead, Go uses incremental evacuation. The old buckets are kept around, and each time I access the map, Go moves 2 buckets from the old array to the new one. This spreads the resizing cost over many operations. During evacuation, new elements go to the bigger buckets, while existing elements are gradually moved. This design ensures maps remain responsive even during growth. The oldbuckets pointer keeps track of the old bucket array until all elements are evacuated."

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

### Explanation
A channel is represented by the hchan struct containing a circular buffer (buf), queues of waiting goroutines (recvq, sendq), and a mutex for thread safety. The qcount tracks elements in the buffer, dataqsiz is buffer capacity.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does a Go channel work internally?
**Your Response:** "A channel is actually a pointer to an hchan struct in the runtime. This struct contains a circular buffer for buffered channels, queues for waiting senders and receivers, and a mutex for thread safety. The buf field points to the actual array that stores data. recvq and sendq are queues of goroutines waiting to receive or send. When I use a channel, Go locks this struct, manipulates the queues and buffer, then unlocks. The design is thread-safe by default - I don't need to add my own synchronization. For unbuffered channels, the buffer is empty and operations directly match senders with receivers. This structure enables all the channel operations I use in Go code."

### Buffered Channel (Circular Queue)
*   Uses a **Ring Buffer** (array) to store elements.
*   **Send**: Acquires lock -> Copies data to `buf` -> Release lock.
*   **Receive**: Acquires lock -> Copies data from `buf` -> Release lock.

### Explanation
Buffered channels use a ring buffer (circular queue) implemented as an array. Send and receive operations acquire the channel's mutex, copy data to/from the buffer, and release the lock, ensuring thread-safe operation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do buffered channels work internally?
**Your Response:** "Buffered channels use a ring buffer - essentially a circular array. When I send a value, Go locks the channel, copies the data into the buffer at the write position, advances the write pointer, and unlocks. When I receive, Go locks the channel, copies data from the read position, advances the read pointer, and unlocks. The ring buffer design means when the pointers reach the end, they wrap around to the beginning. This is efficient because it avoids shifting elements. The buffer size determines how many sends can happen without blocking. The mutex ensures thread safety, though Go optimizes by using direct transfers when possible instead of going through the buffer."

### Direct Send/Receive (Unbuffered / Bypass)
*   **Optimization**: If a receiver is already waiting, the sender **writes directly** to the receiver's stack memory!
*   **Sudog**: The struct representing a waiting goroutine in `recvq` or `sendq`.

### Explanation
Unbuffered channels optimize by directly copying data between goroutines when both are ready. The sender writes directly to the receiver's stack memory, bypassing the buffer entirely. Sudog structs represent waiting goroutines in the queues.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do unbuffered channels optimize performance?
**Your Response:** "Unbuffered channels have a clever optimization - when a sender and receiver are both ready, Go bypasses the buffer entirely. The sender copies the data directly to the receiver's stack memory. This is much faster than going through a buffer. The waiting goroutines are represented by sudog structs in the sendq and recvq queues. When both parties are ready, Go matches them and does the direct copy. This is why unbuffered channels synchronize - the sender blocks until a receiver is ready, then the transfer happens immediately. This direct transfer optimization makes unbuffered channels very efficient for handoffs between goroutines. The sudog structure tracks the goroutine's state and stack location for this direct copy operation."

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

### Explanation
The GMP scheduler consists of G (goroutines), M (OS threads), and P (processors). Gs are lightweight with 2KB stacks. Ms are real OS threads. Ps provide the context to run Go code and have local run queues. GOMAXPROCS controls P count.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Go's scheduler work?
**Your Response:** "Go uses an M:N scheduler with three components: G for goroutines, M for OS threads, and P for processors. A goroutine is lightweight with only a 2KB stack containing its execution state. An M is a real OS thread that the kernel schedules. A P is a logical processor that provides the context to run Go code - it has a local run queue of goroutines. By default, Go creates one P per CPU core. The scheduler's job is to efficiently map many goroutines onto fewer OS threads. This design allows Go to run thousands of goroutines efficiently without overwhelming the system with too many threads."

### Execution Flow
*   `M` must acquire a `P` to run `G`.
*   `M` picks a `G` from `P`'s **Local Run Queue**.
*   If Local Queue is empty, `M` tries to **Steal** work (half the Gs) from another `P`'s queue. (Work Stealing).
*   If all empty, it checks the **Global Run Queue** (Checked occasionally to prevent starvation).

### Explanation
The scheduler execution flow requires M to acquire P, then picks G from P's local queue. If empty, M steals work from other P's queues (work stealing). If all local queues are empty, M checks the global queue to prevent starvation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the Go scheduler execute goroutines?
**Your Response:** "The scheduler follows a specific flow: an M must acquire a P to run a G. The M picks a goroutine from the P's local run queue. If the local queue is empty, the M tries to steal work from other P's queues - it takes half the goroutines to balance the load. This work stealing ensures all CPUs stay busy. If all local queues are empty, the M checks the global run queue. This hierarchy - local queue first, then work stealing, then global queue - optimizes for cache locality and load balancing. The work stealing is particularly important because it prevents one CPU from being idle while others are overloaded. This design efficiently distributes work across all available processors."

### Preemption (Cooperative & Asynchronous)
*   **Old**: Cooperative. Occurred at function calls. Loops could hang the scheduler.
*   **New (Go 1.14+)**: **Asynchronous Preemption**. Sys signals (`SIGURG`) interrupt a goroutine running > 10ms.

### Explanation
Go's preemption evolved from cooperative (only at function calls) to asynchronous (Go 1.14+) using SIGURG signals. Asynchronous preemption interrupts goroutines running longer than 10ms, preventing infinite loops from hanging the scheduler.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Go handle preemption?
**Your Response:** "Go's preemption has evolved significantly. Originally, Go used cooperative preemption - goroutines could only be preempted at function calls. This meant tight loops could hang the entire scheduler. Since Go 1.14, Go uses asynchronous preemption with system signals. If a goroutine runs for more than 10ms without yielding, the runtime sends a SIGURG signal to interrupt it. This allows the scheduler to preempt even goroutines stuck in infinite loops. The change was crucial for fairness and responsiveness - it prevents a single misbehaving goroutine from starving others. This asynchronous preemption works with the operating system's signal handling to safely pause goroutines and schedule others."

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

### Explanation
Empty interfaces (eface) store type information and a data pointer. The _type field describes the dynamic type, while data points to the actual value. This enables Go to hold any type in an interface{} variable.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do empty interfaces work internally?
**Your Response:** "An empty interface is represented by an eface struct with two fields: _type and data. The _type field contains information about the actual type stored - whether it's an int, string, or custom type. The data field is a pointer to the actual value. When I assign something to an interface{}, Go creates this eface structure. This design allows the interface to hold any type while preserving type information for later use. The type information enables type assertions and reflection. The data pointer might point directly to the value or to a copy, depending on the value's size. This two-word structure is the foundation of Go's type system."

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

### Explanation
Non-empty interfaces (iface) use an itab (interface table) containing method pointers. The itab caches the concrete type and specialized method implementations, enabling O(1) dynamic dispatch for interface method calls.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do non-empty interfaces work internally?
**Your Response:** "Non-empty interfaces use an iface struct with an itab and data pointer. The itab is the interface table that contains the concrete type information and pointers to the method implementations. When I assign a type to an interface, Go creates an itab that caches the method pointers for that specific type. This enables dynamic dispatch - when I call a method on the interface, Go looks up the method pointer in the itab table and calls it. This is an O(1) operation, so interface method calls are very fast. The itab is created once per type-interface pair and cached for reuse. This design balances flexibility with performance - interfaces give me polymorphism while keeping method call overhead minimal."

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

### Explanation
A slice is a 3-word struct containing a pointer to the backing array, length, and capacity. When append exceeds capacity, Go allocates a new larger array (typically double), copies elements, and updates the slice header to point to the new array.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do slices work internally?
**Your Response:** "A slice is just a small struct with three fields: a pointer to the underlying array, the length, and the capacity. When I pass a slice to a function, I'm copying this struct, not the array. This is why slices are cheap to pass around. When I append beyond capacity, Go allocates a new array - usually double the size - copies all elements, and updates the pointer in the slice header. The old array gets garbage collected if there are no other references. A common gotcha is sub-slice memory leaks - if I keep a reference to a small part of a huge array, the entire huge array stays in memory because the slice header still points to it. This design makes slices efficient while maintaining the flexibility of dynamic arrays."
