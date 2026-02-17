Here’s a **complete overview of Java Collections keywords, interfaces, classes, and commonly used methods** in the **Java Collections Framework**.

---

# 1️⃣ Important Packages

```java
import java.util.*;
import java.util.stream.*;
```

---

# 2️⃣ Core Interfaces (Hierarchy)

### 🔹 Iterable (Root Interface)

* `iterator()`
* `forEach()`
* `spliterator()`

---

### 🔹 Collection (Extends Iterable)

Common methods:

* `add(E e)`
* `addAll(Collection c)`
* `remove(Object o)`
* `removeAll(Collection c)`
* `retainAll(Collection c)`
* `clear()`
* `size()`
* `isEmpty()`
* `contains(Object o)`
* `containsAll(Collection c)`
* `iterator()`
* `toArray()`
* `stream()`
* `parallelStream()`

---

## 📚 3️⃣ List Interface (Ordered, Allows Duplicates)

### Common Methods:

* `add(int index, E element)`
* `get(int index)`
* `set(int index, E element)`
* `remove(int index)`
* `indexOf(Object o)`
* `lastIndexOf(Object o)`
* `listIterator()`
* `subList(int from, int to)`
* `sort(Comparator c)`
* `replaceAll(UnaryOperator op)`

### 🔹 Implementations:

* ArrayList
* LinkedList
* Vector
* Stack

---

## 📚 4️⃣ Set Interface (No Duplicates)

### Common Methods:

Same as Collection +

* `spliterator()`

### 🔹 Implementations:

* HashSet
* LinkedHashSet
* TreeSet

---

## 📚 5️⃣ Queue Interface (FIFO)

### Common Methods:

* `add(E e)`
* `offer(E e)`
* `remove()`
* `poll()`
* `element()`
* `peek()`

### 🔹 Implementations:

* PriorityQueue
* ArrayDeque
* LinkedList

---

## 📚 6️⃣ Deque Interface (Double-ended Queue)

### Methods:

* `addFirst()`
* `addLast()`
* `offerFirst()`
* `offerLast()`
* `removeFirst()`
* `removeLast()`
* `peekFirst()`
* `peekLast()`
* `push()`
* `pop()`

### Implementation:

* ArrayDeque
* LinkedList

---

# 7️⃣ Map Interface (Key-Value Pair)

⚠ Map does NOT extend Collection.

### Common Methods:

* `put(K key, V value)`
* `putAll(Map m)`
* `get(Object key)`
* `remove(Object key)`
* `replace(K key, V value)`
* `containsKey(Object key)`
* `containsValue(Object value)`
* `keySet()`
* `values()`
* `entrySet()`
* `getOrDefault()`
* `compute()`
* `computeIfAbsent()`
* `computeIfPresent()`
* `merge()`
* `size()`
* `isEmpty()`
* `clear()`

### 🔹 Implementations:

* HashMap
* LinkedHashMap
* TreeMap
* Hashtable
* ConcurrentHashMap

---

# 8️⃣ Iterator Interface

Methods:

* `hasNext()`
* `next()`
* `remove()`
* `forEachRemaining()`

---

# 9️⃣ ListIterator Interface (For List Only)

Methods:

* `hasPrevious()`
* `previous()`
* `nextIndex()`
* `previousIndex()`
* `set(E e)`
* `add(E e)`

---

# 🔟 Collections Utility Class

### Collections methods:

* `sort()`
* `reverse()`
* `shuffle()`
* `binarySearch()`
* `min()`
* `max()`
* `frequency()`
* `copy()`
* `fill()`
* `unmodifiableList()`
* `synchronizedList()`
* `emptyList()`
* `singleton()`

---

# 1️⃣1️⃣ Arrays Utility Class

### Arrays methods:

* `sort()`
* `parallelSort()`
* `binarySearch()`
* `equals()`
* `fill()`
* `copyOf()`
* `asList()`
* `toString()`
* `deepToString()`

---

# 1️⃣2️⃣ Comparable & Comparator

### Comparable

* `compareTo(T o)`

### Comparator

* `compare(T o1, T o2)`
* `reversed()`
* `thenComparing()`
* `comparing()`

---

# 1️⃣3️⃣ Stream Methods (Java 8+)

From Collection:

* `stream()`
* `parallelStream()`

Common Stream methods:

* `filter()`
* `map()`
* `flatMap()`
* `sorted()`
* `distinct()`
* `limit()`
* `skip()`
* `forEach()`
* `collect()`
* `reduce()`
* `count()`
* `anyMatch()`
* `allMatch()`
* `noneMatch()`
* `findFirst()`
* `findAny()`

---

# ✅ Summary of Main Interfaces

```
Iterable
   └── Collection
         ├── List
         ├── Set
         └── Queue
Map (separate hierarchy)
```

---

If you want, I can also provide:

* ✅ Java Collection hierarchy diagram
* ✅ Interview important methods only
* ✅ Difference between List, Set, Map
* ✅ Time complexity chart of all collections
* ✅ Practice programs

Tell me what you need 😊
