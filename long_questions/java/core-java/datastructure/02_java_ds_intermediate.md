# 🟡 Java Data Structures: Level 2 (Intermediate)
*Topics: Collections, Lists, Sets, Maps, Trees, Graphs*

## 3️⃣ Collections Framework (List, Set, Queue)

### Question 43: Add, remove, get, set methods.

**Answer:**
- `add(E e)`: Appends to end.
- `add(int index, E e)`: Inserts at index.
- `remove(Object o)`: Removes first occurrence.
- `remove(int index)`: Removes element at index.
- `get(int index)`: Returns element (O(1) for ArrayList).
- `set(int index, E e)`: Replaces element at index.

---

### Question 44: `addAll()`, `removeAll()`, `retainAll()`.

**Answer:**
Bulk operations.
- `list1.addAll(list2)`: Appends list2 to list1 (Union).
- `list1.removeAll(list2)`: Removes elements found in list2 (Difference).
- `list1.retainAll(list2)`: Keeps ONLY elements found in list2 (Intersection).

---

### Question 45: Difference between `ArrayList` and `LinkedList`.

**Answer:**
- **`ArrayList`**: Dynamic Array.
  - Access: O(1).
  - Insert/Delete (Middle): O(N) (shifting needed).
  - Cache friendly.
- **`LinkedList`**: Doubly Linked List.
  - Access: O(N).
  - Insert/Delete (if Ref held): O(1).
  - More memory overhead (Node objects).

---

### Question 46: `subList()` — what happens if you modify the original list?

**Answer:**
`sub = list.subList(from, to)`.
It returns a **View**.
- Modifications to `sub` affect `original`.
- **Structural modifications** to `original` (add/remove) make `sub` undefined (ConcurrentModificationException).

---

### Question 47: `List.of()` vs `Arrays.asList()` vs mutable lists.

**Answer:**
- `new ArrayList<>(...)`: Fully mutable. accepts nulls.
- `Arrays.asList(...)`: Fixed size (Mutable elements, can't add/remove). allows nulls.
- `List.of(...)` (Java 9): **Immutable**. No nulls.

---

### Question 48: Iterating lists: for, iterator, for-each, streams.

**Answer:**
1.  **For-loop:** `for (int i=0; i<list.size(); i++)` (Bad for LinkedList).
2.  **Enhanced For:** `for (String s : list)`.
3.  **Iterator:** `while (it.hasNext())`. Allows safe removal `it.remove()`.
4.  **Stream:** `list.stream().forEach(...)`.

---

### Question 49: `add()`, `remove()`, `contains()` (Set).

**Answer:**
- `add()`: Returns `false` if element already exists.
- `contains()`: O(1) for HashSet.
- `remove()`: O(1) for HashSet.

---

### Question 50: Difference between `HashSet`, `TreeSet`, `LinkedHashSet`.

**Answer:**
- **`HashSet`**: Unordered. Fastest (O(1)). Uses HashMap.
- **`LinkedHashSet`**: Insertion Order preserved. Slightly slower.
- **`TreeSet`**: Sorted Order (Default or Comparator). O(log N). Uses Red-Black Tree.

---

### Question 51: How does `TreeSet` maintain order? What interface is required?

**Answer:**
Elements must implement `Comparable` (Natural order).
OR you pass a `Comparator` to constructor.
Without comparison logic, throws `ClassCastException`.

---

### Question 52: How to convert a List to Set and vice versa?

**Answer:**
- List -> Set (Remove duplicates): `new HashSet<>(list)`.
- Set -> List: `new ArrayList<>(set)`.

---

### Question 53: `put()`, `get()`, `remove()`, `containsKey()`, `containsValue()` (Map).

**Answer:**
- `put(K, V)`: Returns old value or null.
- `get(K)`: Returns value or null.
- `containsKey(K)`: O(1).
- `containsValue(V)`: O(N) (Iterates all values).

---

### Question 54: Iterating over Map entries: `entrySet()`, `keySet()`, `values()`.

**Answer:**
- **`entrySet()`**: Best (Key + Value). `for (Map.Entry e : map.entrySet())`.
- **`keySet()`**: Keys only.
- **`values()`**: Values only.

---

### Question 55: Difference between `HashMap` and `TreeMap`?

**Answer:**
- **`HashMap`**: Unordered. O(1). Allows 1 null key.
- **`TreeMap`**: Sorted by Key. O(log N). No null keys.

---

### Question 56: Difference between `computeIfAbsent()` and `putIfAbsent()`.

**Answer:**
- **`putIfAbsent(K, V)`**: Puts V if K missing. V is calculated **before** calling (Eager).
- **`computeIfAbsent(K, Function)`**: Calculates value **only if** K is missing (Lazy). Better performance for expensive creation.

---

### Question 57: What happens when two keys have the same hashcode?

**Answer:**
**Collision.**
Elements store in the same bucket (Linked List or Tree).
`equals()` is called to distiguish keys.
If `equals` returns true, value is overwritten. If false, entry is chained.

---

### Question 58: How to maintain insertion order in a Map?

**Answer:**
Use **`LinkedHashMap`**.
It maintains a Doubly Linked List across all entries.

---

### Question 59: Difference between `Queue` and `Deque`.

**Answer:**
- **`Queue`**: FIFO (First In First Out). Add Tail, Remove Head.
- **`Deque`**: Double Ended Queue. Add/Remove from both Head and Tail.

---

### Question 60: Methods in `Queue`: `offer()`, `poll()`, `peek()`.

**Answer:**
Safe methods (return false/null instead of Exception).
- `offer()`: Add (False if full).
- `poll()`: Remove & Return (Null if empty).
- `peek()`: Return Head (Null if empty).

---

### Question 61: Methods in `Deque`: `addFirst()`, `addLast()`, `removeFirst()`, `removeLast()`.

**Answer:**
Allow using Deque as both Stack (`push` = `addFirst`, `pop` = `removeFirst`) and Queue (`addLast`, `removeFirst`).

---

### Question 62: Difference between `Stack` class and `Deque` for stack operations.

**Answer:**
- **`Stack` (Class):** Legacy. Synchronized (Slow). Extends Vector.
- **`Deque` (Interface):** Recommended. Use `ArrayDeque`. standard non-sync stack implementation.

---

### Question 63: How to implement a stack/queue using arrays or linked lists?

**Answer:**
- **Array Stack:** Ptr `top`. `push`: arr[++top]. `pop`: arr[top--]. (Check overflow).
- **LinkedList Queue:** `head` (pop), `tail` (push). O(1) ops.

---

### Question 64: Difference between `PriorityQueue` and `Queue`.

**Answer:**
`Queue` is FIFO.
`PriorityQueue` orders elements by Priority (Comparator or Natural).
`poll()` returns the Smallest/Highest priority element, NOT insertion order.
Uses Binary Heap (O(log N)).

---

## 4️⃣ Advanced Data Structures (LinkedList, Hashing, Trees)

### Question 65: Difference between singly and doubly linked list.

**Answer:**
- **Singly:** Node has `data` + `next`. Forward traversal only. Less memory.
- **Doubly:** Node has `data` + `next` + `prev`. Bi-directional traversal. Deletion is O(1) if Node ref is known (can access prev). More memory.

---

### Question 66: Common methods: `addFirst()`, `addLast()`, `removeFirst()`, `removeLast()`.

**Answer:**
Available in `LinkedList` (implements `Deque`).
- `addFirst(E e)`: Adds to head -> O(1).
- `removeLast()`: Removes from tail -> O(1).
Supports efficient stack/queue behavior.

---

### Question 67: Traversal and searching operations.

**Answer:**
- **Access/Search:** Sequential scan. Time O(N).
- **Traversal:** `ListIterator` allows traversing backward (`hasPrevious()`).

---

### Question 68: How does `HashMap` compute the bucket for a key?

**Answer:**
1.  Compute `h = key.hashCode()`.
2.  Spread bits (XOR shift) to reduce collisions.
3.  Index = `(n - 1) & h` (where n is array size, power of 2).

---

### Question 69: Difference between `hashCode()` and `equals()`.

**Answer:**
- **`hashCode()`:** Returns integer summary. Used to find Bucket.
- **`equals()`:** Returns boolean. Used to verify identity/content match inside the bucket.
**Contract:** If equal, hashCodes MUST be same. If hashCodes same, may NOT be equal (Collision).

---

### Question 70: Why is `hashCode()` required for hash-based collections?

**Answer:**
Without it, the collection cannot know *where* to look for the object.
It would have to scan the entire array (O(N)) instead of O(1) lookup.

---

### Question 71: Tree traversal (`preorder`, `inorder`, `postorder`) — implement in Java.

**Answer:**
Recursive approach on `Node { val, left, right }`.
```java
void inorder(Node n) {
    if (n == null) return;
    inorder(n.left);
    visit(n.val);
    inorder(n.right);
}
```
- Pre: Root, Left, Right.
- In: Left, Root, Right.
- Post: Left, Right, Root.

---

### Question 72: Binary search tree insertion & search.

**Answer:**
- **Search:** `if (val < node.val) search(left) else search(right)`. O(H).
- **Insert:** Same logic to find leaf position, then attach new node.

---

### Question 73: Difference between BST and Heap (method/operation-level).

**Answer:**
- **BST:** Ordered (Left < Root < Right). Fast Search O(log N).
- **Heap (Min-Heap):** Complete Tree. Root is Minimum. No order between siblings. Fast Access Min O(1).

---

## 5️⃣ Functional Streams Basics

### Question 74: `map()`, `flatMap()`, `filter()`.

**Answer:**
- `filter(Predicate)`: Selecting subset (`x > 10`).
- `map(Function)`: Transforming 1-to-1 (`x -> x*2`).
- `flatMap(Function)`: Transforming 1-to-N (`nestedList -> stream`). Flattens structure.

---

### Question 75: `collect()`, `toList()`, `toSet()`.

**Answer:**
Terminal operations.
- `.collect(Collectors.toList())` (Java 8).
- `.toList()` (Java 16+, returns unmodifiable list).

---

### Question 76: `reduce()` — sum, max, concatenation.

**Answer:**
Aggregates stream to single value.
`list.stream().reduce(0, (a, b) -> a + b)` (Sum).

---

### Question 77: `forEach()` — usage.

**Answer:**
Iterates stream. Side-effects only (Printing, Modifying external state).
`stream.forEach(System.out::println)`.

---

### Question 78: `sorted()`, `distinct()`.

**Answer:**
- `sorted()`: Natural order.
- `sorted(Comparator)`: Custom order.
- `distinct()`: Removes duplicates (uses `equals()`). State-ful intermediate operations.

---

### Question 79: Parallel stream vs sequential stream.

**Answer:**
`list.parallelStream()`.
Divides task (ForkJoinPool). Good for CPU-intensive tasks on large datasets.
Bad for blocking I/O or dependent tasks.

---

### Question 80: Stream from arrays: `Arrays.stream(arr)`.

**Answer:**
Handles primitive arrays efficiently (`IntStream`, `mDoubleStream`) unlike `Stream.of(arr)` which might box them if not careful.

---

## 6️⃣ Coding Patterns (Basics)

### Question 81: Sliding window over array/string.

**Answer:**
Technique for subarrays/substrings.
Maintain `left` and `right` indices. Expand `right`. If condition violated, shrink `left`.
(e.g., Max sum subarray of size K).

---

### Question 82: Two-pointer technique.

**Answer:**
One pointer at start, one at end. Or fast/slow pointers.
Used for: Reversing, Palindrome, Remove Duplicates, 2-Sum problem (Sorted).

---

### Question 83: Hashing for arrays/strings.

**Answer:**
Using `HashMap` or `int[26]` frequency array.
Solves "Two Sum", "Anagrams", "First Unique Char" in O(N).

---

### Question 84: Prefix sum arrays.

**Answer:**
`pre[i] = pre[i-1] + arr[i]`.
Allows calculating sum of range `[L, R]` in O(1): `pre[R] - pre[L-1]`.

---

### Question 85: Frequency maps for counting characters or numbers.

**Answer:**
`Map<Integer, Integer> count`.
`map.put(num, map.getOrDefault(num, 0) + 1)`.
Core of Majority Element and Bucket Sort logic.
