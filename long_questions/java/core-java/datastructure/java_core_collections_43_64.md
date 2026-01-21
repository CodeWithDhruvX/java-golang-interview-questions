## 3️⃣ Collections (List, Set, Map)

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

## 4️⃣ Queue, Stack, and Deque

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
