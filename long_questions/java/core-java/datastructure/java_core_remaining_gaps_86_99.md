## 🔹 Tiny Remaining Gaps (Java 9+, Edge Cases, Rare Algos)

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
