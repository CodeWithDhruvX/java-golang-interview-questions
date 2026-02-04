# 🔴 Java Data Structures: Level 3 (Advanced)
*Topics: Java 9+ Features, Edge Cases, Custom Implementations, Concurrency*

## 7️⃣ Gaps & Modern Java Utils

### Question 86: `Arrays.mismatch()` (Java 9+).

**Answer:**
Found in `java.util.Arrays`.
- Returns the **index** of the first element that differs between two arrays.
- Returns `-1` if both arrays are equal.
- Works on primitive and object arrays.
- **UseCase:** Efficiently finding *where* two large arrays diverge without writing a manual loop.

---

### Question 87: `Arrays.parallelSort()` (Java 8+).

**Answer:**
- Uses the **Fork/Join framework** to sort the array using multiple threads.
- **Threshold:** Generally switches to sequential sort for small arrays, but for large datasets, it is significantly faster than `Arrays.sort()` (which is sequential Dual-Pivot Quicksort/TimSort).
- **Time Complexity:** O(n log n) but with better wall-clock time on multi-core CPUs.

---

### Question 88: `String.repeat(int count)` (Java 11+).

**Answer:**
- Returns a new string consisting of the original string repeated `count` times.
- **Example:** `"abc".repeat(3)` -> `"abcabcabc"`.
- **Implementation:** Internally initializes a byte array of exact size and copies data efficiently (faster than loop + StringBuilder for simple repeats).

---

### Question 89: `String.isBlank()` (Java 11+).

**Answer:**
- Returns `true` if the string is empty **OR** contains only **whitespace** (spaces, tabs, newlines).
- **Difference from `isEmpty()`:**
  - `"".isEmpty()` -> true
  - `"  ".isEmpty()` -> false (length is 2)
  - `"  ".isBlank()` -> **true** (visually empty)

---

### Question 90: `Map.ofEntries()` (Java 9+).

**Answer:**
- Used to create **immutable maps** with any number of entries.
- Unlike `Map.of(k1, v1, k2, v2...)` (which handles up to 10 pairs), `Map.ofEntries` takes varargs of `Map.Entry`.
- **Code:**
  ```java
  Map<String, Integer> map = Map.ofEntries(
      Map.entry("a", 1),
      Map.entry("b", 2)
  );
  ```

---

### Question 91: `Set.of()` (Java 9+).

**Answer:**
- Creates an **immutable Set**.
- **Characteristics:**
  - No `null` elements allowed.
  - No duplicate elements allowed (throws IllegalArgumentException at runtime).
  - Iteration order is randomized (not guaranteed).

---

### Question 92: `List.of()` (Java 9+).

**Answer:**
- Creates an **unmodifiable List**.
- **Characteristics:**
  - No `null` elements allowed.
  - Immutable (cannot `add`, `remove`, `set` -> throws `UnsupportedOperationException`).
  - More memory efficient than `Arrays.asList()`.

---

### Question 93: `Collections.emptyList()`, `Collections.emptyMap()`.

**Answer:**
- Returns a **single, shared, immutable** empty instance.
- **Why use it?**
  - **Memory Optimization:** Avoids creating a new `ArrayList` or `HashMap` instance just to return an empty result.
  - **Type Safety:** Generic type inference adapts to the caller.
  - **Safety:** Prevents clients from modifying the returned list (throws exception on write).

---

### Question 94: `Collectors.partitioningBy()` and `Collectors.groupingBy()`.

**Answer:**
- **`groupingBy(Function classifier)`:**
  - Groups elements by a key. Returns `Map<Key, List<T>>`.
  - E.g., Group strings by length.
- **`partitioningBy(Predicate predicate)`:**
  - Special case of grouping. Returns `Map<Boolean, List<T>>`.
  - Keys are always `true` and `false`.
  - E.g., Split numbers into Even vs Odd.

---

### Question 95: `Collectors.counting()`.

**Answer:**
- A downstream collector used inside `groupingBy`.
- Counts the number of elements in each group.
- **Example:**
  ```java
  Map<String, Long> counts = items.stream()
      .collect(Collectors.groupingBy(Item::getType, Collectors.counting()));
  ```

---

### Question 96: `Collectors.joining()`.

**Answer:**
- Concatenates stream elements (Strings) into a single String.
- **Overloads:**
  1. `joining()` -> simple concat.
  2. `joining(delimiter)` -> "a, b, c".
  3. `joining(delimiter, prefix, suffix)` -> "[a, b, c]".
- Uses `StringBuilder` internally.

---

### Question 97: Stream `peek()` — when and why to use carefully.

**Answer:**
- **Purpose:** Exists mainly to support **debugging** (inspecting values as they flow pipeline).
- **Caveat:** It is an *intermediate* operation. If there is no terminal operation, `peek` code will **never run**.
- **Warning:** Do not use `peek` to modify state (side-effects) in a way that affects the result, as stream execution order can be unpredictable (especially in parallel).

---

### Question 98: Sliding window for arrays with variable window size.

**Answer:**
- Unlike fixed-size (k), here the window expands/shrinks based on a condition efficiently.
- **Pattern:**
  1. Expand `right`.
  2. Add `arr[right]` to current state (sum/count).
  3. While (condition broken, e.g., sum > Target):
     - Remove `arr[left]` from state.
     - Increment `left`.
  4. Update result (e.g., min length).

---

### Question 99: Rolling hash / Rabin-Karp style substring search.

**Answer:**
- Used to search for a pattern in text in O(N).
- **Concept:** Compute hash of the pattern. Then compute hash of the first window of text.
- **Rolling:** When moving window one step right, update hash in O(1) by removing the leading character's term and adding the new character's term, instead of rehashing the whole string.
- Crucial for plagiarism detection or DNA sequencing interviews.

---

## 8️⃣ Custom Data Structure Implementations & Advanced Concepts

### Question 100: How to implement an LRU Cache in Java?

**Answer:**
LRU (Least Recently Used) Cache evicts the least recently accessed item when capacity is reached.
**Implementation:**
Use `LinkedHashMap` with `accessOrder = true`.
Override `removeEldestEntry`.

```java
import java.util.LinkedHashMap;
import java.util.Map;

class LRUCache<K, V> extends LinkedHashMap<K, V> {
    private final int capacity;

    public LRUCache(int capacity) {
        super(capacity, 0.75f, true); // true = access order
        this.capacity = capacity;
    }

    @Override
    protected boolean removeEldestEntry(Map.Entry<K, V> eldest) {
        return size() > capacity; // Remove if size exceeds capacity
    }
}
```
**Time Complexity:** O(1) for `get` and `put`.

---

### Question 101: How to implement a Trie (Prefix Tree) in Java?

**Answer:**
Used for Autocomplete, Spell Checker, Strings.
**Structure:** Node contains `Map<Character, Node>` or `Node[26]` and `isEndOfWord`.
```java
class TrieNode {
    TrieNode[] children = new TrieNode[26];
    boolean isEndOfWord;
}

class Trie {
    TrieNode root = new TrieNode();

    public void insert(String word) {
        TrieNode node = root;
        for (char c : word.toCharArray()) {
            if (node.children[c - 'a'] == null) {
                node.children[c - 'a'] = new TrieNode();
            }
            node = node.children[c - 'a'];
        }
        node.isEndOfWord = true;
    }

    public boolean search(String word) {
        TrieNode node = root;
        for (char c : word.toCharArray()) {
            if (node.children[c - 'a'] == null) return false;
            node = node.children[c - 'a'];
        }
        return node.isEndOfWord;
    }
}
```

---

### Question 102: How to implement a Graph using Adjacency List?

**Answer:**
`Map<Integer, List<Integer>>` or `List<List<Integer>>`.
```java
class Graph {
    private Map<Integer, List<Integer>> adjList = new HashMap<>();

    public void addEdge(int src, int dest, boolean bidirectional) {
        adjList.computeIfAbsent(src, k -> new ArrayList<>()).add(dest);
        if (bidirectional) {
            adjList.computeIfAbsent(dest, k -> new ArrayList<>()).add(src);
        }
    }
    
    // Traversal...
}
```
Used for sparse graphs (most real-world graphs). Space O(V + E).

---

### Question 103: BFS vs DFS implementation in Java?

**Answer:**
- **BFS (Breadth-First Search):** Uses **Queue**. Level-order traversal. Shortest path in unweighted graph.
  ```java
  Queue<Integer> q = new LinkedList<>();
  q.offer(start);
  seen.add(start);
  while(!q.isEmpty()) { int curr = q.poll(); ... }
  ```
- **DFS (Depth-First Search):** Uses **Stack** or **Recursion**. exploring path to depth.
  ```java
  void dfs(int node, Set<Integer> visited) {
      if(visited.contains(node)) return;
      visited.add(node);
      for(int neighbor : adj.get(node)) dfs(neighbor, visited);
  }
  ```

---

### Question 104: How to implement a Min/Max Heap?

**Answer:**
Although `PriorityQueue` exists, knowing array implementation is key.
**Min Heap Property:** Parent <= Children.
- **Parent(i):** `(i-1)/2`
- **Left(i):** `2*i + 1`
- **Right(i):** `2*i + 2`
**Operations:**
- `insert`: Add to end, **heapifyUp** (swap with parent until rule satisfied).
- `removeMin`: Swap root with last, remove last, **heapifyDown** (swap with smaller child).

---

### Question 105: Disjoint Set Union (DSU) / Union-Find.

**Answer:**
Used for Cycle Detection, MST (Kruskal’s), Connected Components.
**Optimizations:** Path Compression + Union by Rank/Size.
```java
class DSU {
    int[] parent;
    public DSU(int n) {
        parent = new int[n];
        for(int i=0; i<n; i++) parent[i] = i;
    }
    public int find(int x) { // Path Compression
        if(parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }
    public void union(int x, int y) {
        int rootX = find(x), rootY = find(y);
        if(rootX != rootY) parent[rootX] = rootY; 
    }
}
```
Time: O(α(N)) ≈ O(1).

---

### Question 106: ConcurrentHashMap vs Hashtable vs SynchronizedMap.

**Answer:**
- **Hashtable:** Legacy. Locks **entire map** for every operation. Slow. No nulls.
- **Collections.synchronizedMap():** Wraps map in a mutex. Locks **entire map**.
- **ConcurrentHashMap:**
  - **Java 7:** Segment Locking (Array of Segments).
  - **Java 8+:** CAS (Compare-And-Swap) + `synchronized` only on **Head of Bucket**.
  - Allows concurrent reads without locking.
  - Locking only happens on write, and only on the specific bucket node being modified.
  - **No null keys/values**.

---

### Question 107: BlockingQueue — Producer-Consumer Pattern.

**Answer:**
Thread-safe queue that waits (blocks) when:
- Retrieving from empty queue.
- Adding to full queue.
**Implementations:** `ArrayBlockingQueue`, `LinkedBlockingQueue`.
**Methods:**
- `put()` (blocks if full).
- `take()` (blocks if empty).
Crucial for decoupling threads in Producer-Consumer systems.

---

### Question 108: CopyOnWriteArrayList.

**Answer:**
- Thread-safe variant of `ArrayList`.
- **Mechanism:** All mutative operations (`add`, `set`) make a fresh copy of the underlying array.
- **Use Case:** Read-heavy, Write-rare scenarios (e.g., Listeners list).
- **Iterator:** Snapshot iterator (doesn't reflect changes made after iterator creation, never throws `ConcurrentModificationException`).

---

### Question 109: IdentityHashMap vs WeakHashMap.

**Answer:**
- **IdentityHashMap:** Uses `==` (reference equality) instead of `equals()` for keys. Useful for serialization or topology preservation.
- **WeakHashMap:** Keys are `WeakReference`. If a key is no longer strongly referenced outside the map, the entry is garbage collected. Useful for Caches/Metadata to avoid memory leaks.

---

### Question 110: EnumSet and EnumMap.

**Answer:**
- **EnumSet:** Highly optimized Set for Enums. Uses a **bit vector** (long). Extremely fast (O(1) with low constant). space efficient.
- **EnumMap:** Map with Enum keys. Uses internal **Array**. Faster than HashMap (no hashing needed, just ordinal index).

---

### Question 111: BitSet in Java.

**Answer:**
Vector of bits that grows as needed.
- Uses `long[]` internally (64 bits per word).
- **Operations:** `.set(bitIndex)`, `.get(bitIndex)`, `.and()`, `.or()`, `.xor()`.
- **Use Case:** Compact storage of boolean flags, Bloom Filters.

---

### Question 112: How to detect a cycle in a LinkedList?

**Answer:**
**Floyd’s Cycle-Finding Algorithm (Tortoise and Hare).**
- Slow ptr: 1 step.
- Fast ptr: 2 steps.
- If they meet (`slow == fast`), there is a cycle.
- If `fast` reaches null, no cycle.
- **Start of Cycle:** Reset slow to head, move both 1 step. Meet point is start.

---

### Question 113: How to find the middle of a LinkedList?

**Answer:**
Use Fast/Slow pointers.
- Move `fast` 2 steps, `slow` 1 step.
- When `fast` reaches end, `slow` is at middle.
- If even length, `slow` is at `n/2`.

---

### Question 114: Flatten a nested List or Iterator.

**Answer:**
Often asked as "Flatten 2D Vector" or "Nested List Iterator".
- Keep a Stack of Iterators (if deeply nested).
- Or `flatMap` in Streams.
- **Interview Pattern:** Maintain pointer to current inner list. Advance when exhausted.

---

### Question 115: Monotonic Stack (Next Greater Element).

**Answer:**
- Find the next greater element for every element in O(N).
- Stack stores indices/values in **decreasing** order.
- When current value > top of stack: Pop (Current is NGE for popped). Push current.
