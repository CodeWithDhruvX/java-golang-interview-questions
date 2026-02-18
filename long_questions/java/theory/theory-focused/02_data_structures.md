# Data Structures & Collections - Interview Answers

> 🎯 **Focus:** These answers focus on *why* and *when* you choose a specific data structure.

### 1. ArrayList vs LinkedList?
"It comes down to memory layout and how the CPU accesses data.

**ArrayList** is backed by a dynamic array. Reading from it is super fast—O(1)—because the data is contiguous in memory, so the CPU can pre-fetch it efficiently. However, inserting or deleting in the middle is slow because it has to shift all the elements over.

**LinkedList** is a doubly-linked list. Inserting is fast—O(1)—because you just update pointers. But reading is slow—O(n)—because you have to traverse the nodes, which are scattered in memory, causing cache misses.

Honestly, in modern Java, `ArrayList` is almost always better due to **cache locality**. I haven't used a `LinkedList` in production in years; even for queues, `ArrayDeque` is faster."

---

### 2. HashMap vs TreeMap vs LinkedHashMap?
"These are all Maps, but they differ in how they **order** the keys.

**HashMap** has no order. It’s the default choice because it’s the fastest—O(1) lookup.

**LinkedHashMap** maintains **insertion order**. It remembers the sequence in which you added keys. It’s slightly slower but very useful when I’m returning a JSON response and the order of fields matters, or when implementing an LRU cache.

**TreeMap** keeps keys **sorted** (using natural order or a Comparator). It’s slower—O(log n)—because it uses a Red-Black tree internally. I use this only when I specifically need to iterate over data in a sorted way, like printing a daily report."

---

### 3. How does HashMap work internally?
"It maps a key to a value using **Hashing**.

When I call `put(Key, Value)`, the map uses `Key.hashCode()` to calculate an index and places the node in a specific 'bucket' in an array.

If two keys hash to the same bucket—which is a **collision**—they form a Linked List. When I call `get()`, it finds the bucket and iterates through that list using `equals()` to find the right key.

A cool improvement in Java 8 is that if a bucket gets too full (over 8 nodes), it converts that list into a **Red-Black Tree**. This ensures that even with many collisions, performance remains O(log n) instead of degrading to O(n)."

---

### 4. What is the difference between HashSet and TreeSet?
"It’s the same distinction as HashMap vs TreeMap, because internally, a **Set is just a wrapper around a Map** (with a dummy value).

**HashSet** is unordered and fast. It uses a HashMap. I use it when I just need to deduplicate a list of IDs and don't care about order.

**TreeSet** is sorted. It uses a TreeMap. I use it when I need to perform range operations—like identifying all values between 1 and 10 using `subSet()`—or when I need the data to stay sorted automatically."

---

### 5. Array vs Collection (List)?
"**Arrays** are low-level, fixed-size structures. They are slightly faster and memory-efficient for primitives (`int[]`), but they are rigid—you can’t resize them once created.

**Collections** (like `ArrayList`) are dynamic. They grow automatically and provide rich utility methods like `contains()`, `remove()`, and Stream support.

In business logic, I almost always use `Lists` because they are easier to read and safer. Arrays are covariant (meaning `String[]` is an `Object[]`), which can lead to runtime errors. Collections are invariant, so the compiler catches type mismatches early. I basically only use arrays for performance-critical buffers or var-args."

---

### 6. What is a PriorityQueue?
"It’s a special queue that orders elements based on **priority** rather than just arrival time.

When you call `poll()`, you don't necessarily get the oldest element; you get the 'highest priority' one (based on a Comparator). Internally, it uses a **Binary Heap** to keep the highest priority element efficiently at the top.

It’s perfect for task scheduling systems. For example, if I'm processing support tickets, I want to process 'Critical' tickets before 'Low' priority ones, even if the Low ones arrived first. PriorityQueue handles that logic automatically."

---

### 7. What is ConcurrentHashMap?
"It’s a thread-safe implementation of Map that is significantly faster than the old `Hashtable` or `Collections.synchronizedMap`.

Older approaches locked the **entire** map for every write, creating a bottleneck. `ConcurrentHashMap` uses **Lock Stripping**—it splits the map into segments and only locks the specific segment being written to.

Even better, reads are generally lock-free. So in a web application where you have many threads reading from a cache and only occasional updates, `ConcurrentHashMap` provides massive performance gains."

---

### 8. Comparable vs Comparator?
"Both are used for sorting, but they differ in where the logic lives.

**Comparable** is for **Natural Ordering**. The class itself implements it (`implements Comparable`). You define `compareTo` inside the class. For example, a `Date` object naturally compares itself to another Date.

**Comparator** is for **Custom Ordering**. It’s an external implementation—usually a separate class or a lambda.

I use Comparable for the default sort (like sorting Users by ID). But if I have a UI that needs to sort Users by 'Name' or 'Join Date', I wouldn't change the User class; I’d just pass a specific `Comparator` to the sort method."

---

### 9. What is the contract between hashCode() and equals()?
"The contract is: **If two objects are equal, they MUST have the same hashCode.**

If `a.equals(b)` is true, then `a.hashCode()` must match `b.hashCode()`.

If you break this—for example, by overriding `equals` but not `hashCode`—hash-based collections like `HashMap` and `HashSet` will break. You might put an object in, but because the hash code is wrong, you’ll never be able to retrieve it again, or you'll end up with duplicate entries in a Set.

That's why I always let my IDE or Lombok generate both methods together."

---

### 10. CopyOnWriteArrayList?
"It’s a thread-safe List designed specifically for **read-heavy** scenarios.

Every time you add or update an element, it actually creates a fresh **copy** of the underlying array, modifies it, and then swaps the reference. This means that readers never get blocked—they continue reading the 'old' version safely while the write happens.

It’s very expensive for writes, so I wouldn't use it for a high-frequency event log. But it's perfect for things like a list of Event Listeners, where you iterate and notify often, but rarely add new listeners."
