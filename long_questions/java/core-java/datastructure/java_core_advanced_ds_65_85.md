## 5️⃣ Other Important Data Structures

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

### Hashing

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

### Tree & Graph Basics

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

## 6️⃣ Stream + Functional Methods

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

## ✅ Bonus: Common Coding Interview Patterns

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
