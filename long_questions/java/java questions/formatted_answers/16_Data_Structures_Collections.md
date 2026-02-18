# 16. Data Structures (Collections Framework)

**Q: ArrayList vs LinkedList**
> "Think of **ArrayList** as a dynamic array. It forces all elements to sit next to each other in memory.
> *   **Access** is instant (O(1)) because the computer knows exactly where index 5 is.
> *   **Adding/Removing** is slow (O(n)) because if you delete element 0, the computer has to shift every other element one step to the left to fill the gap.
>
> **LinkedList** is a chain of nodes. Each node holds data and a pointer to the next node.
> *   **Access** is slow (O(n)) because to get to the 5th element, you have to start at the head and hop 5 times.
> *   **Adding/Removing** is fast (O(1)) *if you are already at that spot*, because you just change two pointers. No shifting required.
>
> **Verdict**: Use `ArrayList` 99% of the time. Modern CPU caches love contiguous memory, so `ArrayList` is usually faster even for inserts unless the list is massive."

**Indepth:**
> **Resize Strategy**: When ArrayList is full, it creates a new array 50% larger (1.5x) and copies elements. This resizing is costly (O(n)), so always initialize with a capacity if you know the size (`new ArrayList<>(1000)`).


---

**Q: List.of() vs Arrays.asList()**
> "**List.of()** (Java 9+) creates a truly **Immutable List**.
> *   You cannot add/remove elements.
> *   You cannot even set/replace elements.
> *   It doesn't allow `null` elements.
>
> **Arrays.asList()** creates a **Fixed-Size List** backed by an array.
> *   You cannot add/remove (throws exception).
> *   But you **can** replace items (`set()`), and those changes write through to the original array.
> *   It allows `null`."

**Indepth:**
> **Best Practice**: Prefer `List.of()` for constants. Use `Arrays.asList()` only when working with legacy APIs that expect it, or when you need a write-through view of an array.


---

**Q: SubList() caveat**
> "`subList(from, to)` doesn't create a new list. It returns a **View** of the original list.
>
> If you start modifying the *original* list (adding/removing items) after creating a sublist, the sublist becomes undefined and will likely throw a `ConcurrentModificationException` the next time you touch it.
>
> Always treat the sublist as temporary, or wrap it in a `new ArrayList<>(list.subList(...))` to detach it."

**Indepth:**
> **Memory Leak**: In old Java versions, `subList` held a reference to the entire original parent list, preventing GC. New versions copy or are smarter, but referencing sublists is still risky if the parent list is long-lived.


---

**Q: Iterator vs For-Each**
> "For-each loops are syntactic sugar. Under the hood, they use an Iterator.
>
> The big difference is **Modification**.
> If you are looping through a list and try to do `list.remove(item)`, you will crash with a `ConcurrentModificationException`.
> To remove items safely while looping, you **must** use the `Iterator` explicitly and call `iterator.remove()`."

**Indepth:**
> **Java 8**: `Collection.removeIf(filter)` is the modern, thread-safe, and readable way to remove elements. `list.removeIf(s -> s.isEmpty())` creates an iterator internally and handles it correctly.


---

**Q: HashSet vs TreeSet vs LinkedHashSet**
> "It's all about **Order**.
>
> 1.  **HashSet**: The fastest (O(1)). It uses a HashMap internally. It makes **no guarantees** about order. You put items in, they come out in a random jumbled mess.
> 2.  **LinkedHashSet**: Slightly slower. It maintains **Insertion Order**. If you put in [A, B, C], you iterate out [A, B, C]. It uses a Doubly Linked List running through the entries.
> 3.  **TreeSet**: The slowest (O(log n)). It keeps elements **Sorted** (Natural order or custom Comparator). It uses a Red-Black Tree. Useful if you need range queries (like 'give me all numbers greater than 50')."

**Indepth:**
> **LinkedHashSet**: It maintains a doubly-linked list running through all its entries. This adds memory overhead (two extra pointers per entry) but gives predictable iteration order.


---

**Q: HashMap Internal Working**
> "A HashMap is an array of 'Buckets' (Node<K,V>).
>
> 1.  **Put(K, V)**: We calculate `hash(Key)`. This gives us an index (bucket location).
> 2.  **Collision**: If two keys land in the same bucket, we store them as a Linked List (or a Tree if the list gets too long, Java 8+).
> 3.  **Get(K)**: We go to the bucket, walk through the chain, and compare keys using `.equals()`.
>
> This is why `hashCode()` and `equals()` must be consistent. If they disagree, you might put an object in one bucket but look for it in another, getting `null`."

**Indepth:**
> **Load Factor**: The default load factor is 0.75. When the map is 75% full, it resizes (doubles the bucket array). Setting it higher (1.0) saves memory but increases collisions. Setting it lower (0.5) reduces collisions but wastes memory.


---

**Q: HashMap vs TreeMap**
> "Similar to the Set comparison:
>
> *   **HashMap**: Fast (O(1)). Unordered. Uses hashing.
> *   **TreeMap**: Slower (O(log n)). Sorted by Key. Uses a Red-Black Tree.
>
> Use `TreeMap` only if you need to iterate keys in alphabetical order or find 'the key just higher than X'."

**Indepth:**
> **NavigableMap**: TreeMap implements `NavigableMap`, giving you powerful methods like `ceilingKey(K)`, `floorKey(K)`, `higherKey(K)`, `lowerKey(K)`.


---

**Q: computeIfAbsent vs putIfAbsent**
> "Both try to add a value if the key is missing.
>
> **putIfAbsent(key, value)**: You pass the *actual value*. Even if the key exists, that value object is created (and then ignored), which might be wasteful if creating it is expensive.
>
> **computeIfAbsent(key, function)**: You pass a *function*. The function is **only** executed if the key is missing. This is lazy and much more efficient for expensive objects."

**Indepth:**
> **Concurrency**: `computeIfAbsent` is atomic in `ConcurrentHashMap`. It guarantees the computation happens only once, even if multiple threads race to compute the value for the same key.


---

**Q: Queue vs Deque vs Stack**
> "**Queue** is standard FIFO (First-In-First-Out). You join the back, leave from the front.
>
> **Deque** (Double Ended Queue) is the superhero interface. You can add/remove from **both** ends.
>
> **Stack**: The `Stack` class remains in Java for compatibility, but it's considered legacy (it's synchronized and extends Vector). **Do not use Stack.**
> If you need a LIFO Stack, use `ArrayDeque` and call `push()` and `pop()`."

**Indepth:**
> **Speed**: `ArrayDeque` is faster than `Stack` and `LinkedList` because it uses a resizable array and doesn't allocate nodes for every item. It's cache-friendly.


---

**Q: PriorityQueue**
> "A standard Queue orders by arrival time. A **PriorityQueue** orders by **Priority**.
>
> When you call `poll()`, you don't get the oldest item; you get the 'smallest' item (according to `compareTo`).
> Internally, it uses a **Min-Heap**. Accessing the top element is O(1), but adding/removing is O(log n) because the heap has to re-balance."

**Indepth:**
> **Use Case**: This is perfect for task scheduling (high priority tasks first) or Dijkstra's shortest path algorithm (explore cheapest path first).


---

**Q: hashCode() and equals() contract**
> "The rule is simple but strict:
>
> 1.  If `a.equals(b)` is true, then `a.hashCode()` **MUST** be equal to `b.hashCode()`.
> 2.  If this is violated, HashMaps break. You will put an object in, and you won't be able to retrieve it because the map looks in the wrong bucket.
>
> Note: The reverse isn't true. Two different objects *can* describe the same hash code (Collision), and the map handles that."

**Indepth:**
> **Consistent Hashing**: If two objects are equal, they MUST hash to the same bucket. If they hashed to different buckets, the Map would check Bucket A, find nothing, and return null, effectively "losing" your object.


---

**Q: Fail-Fast vs Fail-Safe**
> "**Fail-Fast** iterators (like ArrayList, HashMap) throw `ConcurrentModificationException` immediately if they detect that someone else changed the collection while they were iterating. They'd rather crash than show you inconsistent data.
>
> "**Fail-Safe** iterators (like `ConcurrentHashMap`, `CopyOnWriteArrayList`) work on a snapshot or a weakly consistent view. They allow modifications during iteration and won't throw an exception, but they might not show you the very latest data."

**Indepth:**
> **COW**: `CopyOnWriteArrayList` is expensive for writes (it copies the entire array on every add!), but it's perfect for "Read-Mostly" lists like Event Listeners where iteration is frequent but modification is rare.

