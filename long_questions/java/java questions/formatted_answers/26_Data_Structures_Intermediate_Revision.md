# 26. Data Structures (Intermediate Revision)

**Q: Regex Validation (matches)**
> "Stop writing complex loops to validate emails. Use Regex.
> `str.matches(\"\\\\d+\")` returns true if the string is all digits.
>
> **Performance Tip**: `String.matches()` re-compiles the pattern every time. If you validate thousands of strings, compile the `Pattern` object once as a `static final` constant and reuse it."

**Indepth:**
> **DOS Attack**: Be careful with Regex. A poorly written regex (nested quantifiers like `(a+)+`) can cause "Catastrophic Backtracking" if the input is malicious, hanging your CPU. This is a common Denial of Service vector.


---

**Q: ArrayList vs LinkedList (Real World)**
> "In textbooks, LinkedList is faster for adding/removing. In the real world, **ArrayList is almost always faster**.
>
> Why? **Cache Locality**.
> ArrayList elements are next to each other in memory. The CPU fetches a chunk of memory and processes it efficiently.
> LinkedList nodes are scattered everywhere. The CPU wastes time waiting for memory fetches (Cache Misses).
> Only use LinkedList if you are building a Queue/Deque or doing heavy splits/merges."

**Indepth:**
> **Memory Overhead**: LinkedList nodes have significant overhead. Each node stores the data object reference + next pointer + previous pointer. That's 24 bytes of overhead per element (on 64-bit JVM) compared to 0 bytes for ArrayList.


---

**Q: List.of() vs Arrays.asList()**
> "**Arrays.asList()** is a bridge to legacy arrays. It allows `set()` but not `add()`, and it passes changes through to the original array.
>
> **List.of()** (Java 9) is the modern standard for constants.
> *   It is **Truly Immutable** (No `set`, no `add`).
> *   It creates a highly optimized internal class (not a standard ArrayList).
> *   It does **not** allow nulls."

**Indepth:**
> **Best Practice**: Use `List.of` generally, but be aware of the "No Nulls" rule. If you need to store nulls (rare), you must use `Arrays.asList` or a standard `ArrayList`.


---

**Q: HashSet vs TreeSet vs LinkedHashSet**
> "You choose based on **Ordering**:
>
> 1.  **HashSet**: 'I don't care about order, just give me speed.' (O(1)).
> 2.  **LinkedHashSet**: 'I want them in the order I inserted them.' (Preserves insertion order).
> 3.  **TreeSet**: 'I want them sorted (A-Z, 1-10).' (O(log n))."

**Indepth:**
> **Consistency**: `TreeSet` uses `compareTo` (or Comparator) to determine equality, *not* `equals()`. If `compareTo` returns 0, TreeSet considers the elements duplicates, even if `equals()` returns false. This can lead to weird bugs.


---

**Q: HashMap vs TreeMap**
> "Same logic as Sets.
> *   **HashMap** for speed (O(1) lookup). Keys are jumbled.
> *   **TreeMap** for sorted keys (O(log n) lookup).
>
> Use `TreeMap` when you need features like `firstKey()`, `lastKey()`, or `subMap()`. For example, getting all events that happened 'between 10 AM and 11 AM'."

**Indepth:**
> **Balancing**: TreeMap uses a Red-Black Tree. This guarantees `log(n)` time for operations, but re-balancing the tree after insertions is more expensive than simply dumping an item into a HashMap bucket.


---

**Q: computeIfAbsent**
> "This method saves you lines of code and enhances performance.
>
> Instead of:
> `if (!map.containsKey(key)) map.put(key, new ArrayList()); return map.get(key);`
>
> You write:
> `return map.computeIfAbsent(key, k -> new ArrayList());`
>
> It's atomic, clean, and ensures the value is created only when needed."

**Indepth:**
> **Concurrent**: `computeIfAbsent` is atomic in `ConcurrentHashMap`. This makes it the perfect tool for implementing local caches without complex `synchronized` blocks.


---

**Q: HashMap Collision Handling**
> "If two keys hash to the same bucket, HashMap starts a **Linked List** in that bucket.
>
> If that list gets too long (more than 8 items), Java 8 automatically transforms it into a **Red-Black Tree**.
> This improves the worst-case performance from O(n) (scanning a long list) to O(log n) (searching a tree)."

**Indepth:**
> **Attack**: This switch to Trees (JEP 180) was done to prevent "Hash Flooding" attacks. Attackers could send thousands of requests with keys that all hash to the same bucket, turning your server into a slow O(n) crawler. The tree protects against this.


---

**Q: Queue vs Deque**
> "**Queue** (First-In-First-Out): Use it for tasks like 'Processing jobs in order'. Methods: `offer()`, `poll()`.
>
> "**Deque** (Double-Ended Queue): You can add/remove from **both** ends.
> Use `ArrayDeque` as your go-to Stack implementation. It is faster than the legacy `Stack` class because it is not synchronized."

**Indepth:**
> **Stack vs Deque**: `Stack` is a class, `Deque` is an interface. `Stack` extends `Vector`, which means every method is synchronized (slow). `ArrayDeque` is the modern, unsynchronized, faster replacement.


---

**Q: PriorityQueue**
> "This isn't a normal queue. It doesn't follow FIFO.
> It keeps elements ordered by **Priority** (Smallest to Largest by default).
>
> Useful for scheduling systems: 'Process the High Priority VIP job before the Regular job, even if the Regular job came first'."

**Indepth:**
> **Implementation**: PriorityQueue uses an array-based **Binary Heap**. It does *not* keep the array sorted. It only guarantees that `array[k] <= array[2*k+1]` and `array[k] <= array[2*k+2]`. This partial ordering is enough for fast polling.


---

**Q: Tree Traversal (Pre, In, Post)**
> "If this is a coding question, remember the position of the **Root**:
>
> 1.  **Pre-Order**: **Root**, Left, Right. (Used for copying trees).
> 2.  **In-Order**: Left, **Root**, Right. (Gives sorted order for BSTs).
> 3.  **Post-Order**: Left, Right, **Root**. (Used for deleting trees/garbage collection)."

**Indepth:**
> **Iterative**: Any recursive traversal can be converted to an iterative one using an explicit `Stack`. In interviews, knowing the iterative approach (especially for Pre-Order) is a big plus.


---

**Q: Streams: map vs flatMap**
> "**map** is for 1-to-1 conversion.
> `Stream<String> -> map(s -> s.length()) -> Stream<Integer>`.
>
> **flatMap** is for flattening nested structures.
> `Stream<List<String>> -> flatMap(List::stream) -> Stream<String>`.
> It 'unwraps' the inner lists so you have one big smooth stream of elements."

**Indepth:**
> **Stream Handling**: `flatMap` is essential when working with `Optional` inside a stream. `stream.map(Optional::stream)` gives you `Stream<Stream<T>>`. `stream.flatMap(Optional::stream)` gives you a clean `Stream<T>`.


---

**Q: Parallel Stream**
> "Don't just add `.parallel()` thinking it makes everything faster.
>
> *   **Good for**: Huge datasets, CPU-intensive tasks (like complex math on every element).
> *   **Bad for**: Small lists, IO operations (network/DB calls), or tasks that need order preservation.
>
> Overhead of splitting tasks and joining threads can often make Parallel Streams *slower* than simple Sequential Streams."

**Indepth:**
> **Spliterator**: Parallel streams effectively rely on the `Spliterator`. If your data source splits poorly (like a LinkedList), parallel performance will be terrible. `ArrayList` splits perfectly.

