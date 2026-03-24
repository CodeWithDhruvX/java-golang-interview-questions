# Java Collections Framework — High-Speed Retrieval Cheatsheet

> **Purpose:** Rapid interview reference with categorized snippets, exception alerts, performance tables, and real-world analogies.

---

## 🚀 The 'Engine' Map — Data Structure Categories

### 🔧 Array-Based Collections
**Core Concept:** Contiguous memory allocation, O(1) random access, O(n) insertions/deletions.

```java
// ArrayList — Dynamic Array
List<Integer> list = new ArrayList<>(List.of(1, 2, 3));
list.add(4);        // O(1) amortized
list.get(1);        // O(1) — direct index access
list.remove(0);     // O(n) — shifts all elements
```
**Interview Logic:** ArrayList is like theater seats — everyone has a specific numbered position, easy to find anyone but moving people requires shifting the whole row.

```java
// Arrays.asList() — Fixed-Size Array View
List<String> fixed = Arrays.asList("a", "b", "c");
fixed.set(0, "x");  // OK — can modify existing
fixed.add("d");     // ❌ UnsupportedOperationException
```
**Interview Logic:** Arrays.asList is like a photo frame — you can change what's inside but can't add or remove frames.

---

### ⛓️ Link-Based Collections  
**Core Concept:** Node-based with pointers, O(1) insertions/deletions, O(n) random access.

```java
// LinkedList — Doubly-Linked Nodes
List<Integer> linked = new LinkedList<>(List.of(1, 2, 3));
linked.addFirst(0);  // O(1) — pointer update
linked.remove(0);   // O(1) — pointer update  
linked.get(2);      // O(n) — traverse from head
```
**Interview Logic:** LinkedList is like a treasure hunt — each clue points to the next, easy to insert new clues but finding the 10th clue requires following the path.

```java
// LinkedList as Deque
Deque<Integer> deque = new LinkedList<>();
deque.addFirst(1);   // O(1)
deque.addLast(2);   // O(1)
deque.pollFirst();  // O(1)
```
**Interview Logic:** Deque is like a subway turnstile — people can enter/exit from both ends efficiently.

---

### 🔑 Hash-Based Collections
**Core Concept:** Hash table with buckets, O(1) average operations, unordered iteration.

```java
// HashSet — Hash Table Storage
Set<Integer> set = new HashSet<>(List.of(3, 1, 4, 1, 5));
set.add(2);         // O(1) average
set.contains(3);    // O(1) average
// Order unpredictable
```
**Interview Logic:** HashSet is like a parking garage — cars stored by license plate hash, instant retrieval but parking spots aren't in any particular order.

```java
// HashMap — Key-Value Hash Table
Map<String, Integer> map = new HashMap<>();
map.put("key", 1);   // O(1) average
map.get("key");      // O(1) average
map.computeIfAbsent("new", k -> new ArrayList<>()); // Lazy init
```
**Interview Logic:** HashMap is like a dictionary — instant word lookup by spelling, but the words aren't arranged alphabetically.

```java
// LinkedHashMap — Insertion Order
Map<String, Integer> ordered = new LinkedHashMap<>();
ordered.put("first", 1);
ordered.put("second", 2);
// Iteration preserves insertion order
```
**Interview Logic:** LinkedHashMap is like a queue ticket system — first-come, first-served order maintained with fast access.

---

### 🌳 Tree-Based Collections
**Core Concept:** Balanced binary search tree, O(log n) operations, sorted iteration.

```java
// TreeSet — Sorted Set
Set<Integer> sorted = new TreeSet<>(List.of(5, 3, 1, 4, 2));
sorted.add(6);       // O(log n)
sorted.first();      // O(log n) — smallest element
// Always sorted: [1, 2, 3, 4, 5, 6]
```
**Interview Logic:** TreeSet is like a library card catalog — always organized alphabetically, finding any book takes logarithmic time.

```java
// TreeMap — Sorted Map
Map<String, Integer> treeMap = new TreeMap<>();
treeMap.put("zebra", 1);
treeMap.put("apple", 2);
// Keys always sorted: [apple, zebra]
```
**Interview Logic:** TreeMap is like a phone book — contacts always in alphabetical order, finding "Smith" means opening to the middle and narrowing down.

---

## 🚨 Red-Flag Alert — Exception Scenarios

### ConcurrentModificationException
**When it throws:** Modifying collection during for-each iteration.

```java
// ❌ THROWS ConcurrentModificationException
List<String> list = new ArrayList<>(List.of("a", "b", "c"));
for (String s : list) {
    if (s.equals("a")) list.remove(s); // Forbidden!
}
```

**Safe Alternatives:**
```java
// ✅ Iterator.remove()
Iterator<String> it = list.iterator();
while (it.hasNext()) {
    if (it.next().equals("a")) it.remove();
}

// ✅ removeIf() (Java 8+)
list.removeIf(s -> s.equals("a"));

// ✅ Create copy and iterate
for (String s : new ArrayList<>(list)) {
    if (s.equals("a")) list.remove(s);
}
```

### UnsupportedOperationException  
**When it throws:** Attempting to modify immutable collections.

```java
// ❌ THROWS UnsupportedOperationException
List<String> immutable = List.of("a", "b", "c");
immutable.add("d");        // Can't modify

Set<String> immutableSet = Set.of("x", "y");
immutableSet.add("z");      // Can't modify

Map<String, Integer> immutableMap = Map.of("a", 1);
immutableMap.put("b", 2);  // Can't modify
```

**Mutable Alternatives:**
```java
// ✅ Create mutable copies
List<String> mutable = new ArrayList<>(List.of("a", "b", "c"));
mutable.add("d");          // Works

Set<String> mutableSet = new HashSet<>(Set.of("x", "y"));
mutableSet.add("z");       // Works
```

### NullPointerException
**When it throws:** Null operations in collections that don't allow null.

```java
// ❌ THROWS NullPointerException  
Map<String, Integer> map = Map.of("a", null); // Map.of() disallows null
Set<String> set = Set.of("a", null);         // Set.of() disallows null

TreeMap<String, Integer> treeMap = new TreeMap<>();
treeMap.put(null, 1);                         // TreeMap disallows null keys
```

**Null-Safe Alternatives:**
```java
// ✅ HashMap allows null keys/values
Map<String, Integer> hashMap = new HashMap<>();
hashMap.put(null, null);  // Works fine
```

---

## ⚡ Big-O Performance Table

| Collection | Add | Remove | Get/Contains | Order | Memory |
|------------|-----|--------|-------------|-------|---------|
| **ArrayList** | O(1)* | O(n) | O(1) | Insertion | Medium |
| **LinkedList** | O(1) | O(1) | O(n) | Insertion | High |
| **HashSet** | O(1) | O(1) | O(1) | Unordered | Medium |
| **TreeSet** | O(log n) | O(log n) | O(log n) | Sorted | High |
| **HashMap** | O(1) | O(1) | O(1) | Unordered | Medium |
| **TreeMap** | O(log n) | O(log n) | O(log n) | Sorted | High |

*Amortized O(1), O(n) when resizing needed

---

## 🛡️ Safe vs. Aggressive — Queue/Deque Methods

| Operation | Throws Exception | Returns Null/Special |
|-----------|------------------|---------------------|
| **Add to Queue** | `add()` | `offer()` |
| **Remove from Queue** | `remove()` | `poll()` |
| **View Queue Head** | `element()` | `peek()` |
| **Add to Front** | `addFirst()` | `offerFirst()` |
| **Add to Back** | `addLast()` | `offerLast()` |
| **Remove from Front** | `removeFirst()` | `pollFirst()` |
| **Remove from Back** | `removeLast()` | `pollLast()` |
| **View Front** | `getFirst()` | `peekFirst()` |
| **View Back** | `getLast()` | `peekLast()` |

```java
// Exception-Throwing (Aggressive)
Queue<Integer> queue = new LinkedList<>();
queue.add(1);        // OK
queue.remove();      // ❌ NoSuchElementException if empty

// Safe (Returns null/special)
queue.offer(1);      // OK
queue.poll();        // ✅ Returns null if empty
```

---

## 🎯 The 'Why' Column — Interview Logic Summary

### List Collections
- **ArrayList:** "Like theater seats — numbered positions, easy access but shifting people is expensive."
- **LinkedList:** "Like a treasure hunt — each clue points to the next, easy to insert but finding position 10 requires traversal."
- **Vector:** "Like a synchronized theater — thread-safe but slower due to locking overhead."

### Set Collections  
- **HashSet:** "Like a parking garage — instant retrieval by license plate hash, but spots aren't organized."
- **TreeSet:** "Like a library card catalog — always organized, finding any book takes logarithmic time."
- **LinkedHashSet:** "Like a queue ticket system — first-come order with fast access."

### Map Collections
- **HashMap:** "Like a dictionary — instant word lookup, but not alphabetically organized."
- **TreeMap:** "Like a phone book — always alphabetical, finding contacts by narrowing down."
- **LinkedHashMap:** "Like a restaurant waitlist — maintains order of arrival with fast lookup."

### Queue Collections
- **ArrayDeque:** "Like a checkout line — efficient adding/removing from both ends."
- **PriorityQueue:** "Like an emergency room — highest priority patients served first, not arrival order."
- **LinkedList as Queue:** "Like a subway turnstile — people can enter/exit from both ends."

---

## 🔥 Critical Interview Patterns

### Frequency Count (Word Count)
```java
Map<String, Integer> count = new HashMap<>();
for (String word : words) {
    count.merge(word, 1, Integer::sum);  // Modern approach
    // OR: count.put(word, count.getOrDefault(word, 0) + 1);
}
```

### Lazy Initialization (Multimap Pattern)
```java
Map<String, List<Integer>> multiMap = new HashMap<>();
multiMap.computeIfAbsent("key", k -> new ArrayList<>()).add(42);
```

### Balanced Parentheses (Stack Pattern)
```java
Deque<Character> stack = new ArrayDeque<>();
for (char c : expression.toCharArray()) {
    if (c == '(') stack.push(c);
    else if (c == ')' && (stack.isEmpty() || stack.pop() != '(')) return false;
}
return stack.isEmpty();
```

### Custom Sorting
```java
// Multi-level comparator
students.sort(Comparator.comparingInt(Student::grade).reversed()
        .thenComparing(Student::name));
```

---

## ⚠️ Common Pitfalls

1. **remove(int) vs remove(Object):** `list.remove(1)` removes index, not value
2. **getOrDefault vs computeIfAbsent:** First doesn't modify map, second does
3. **Iterator Order:** PriorityQueue iteration ≠ sorted order
4. **Null Handling:** HashMap allows null keys, TreeMap doesn't
5. **Immutable Collections:** List.of(), Set.of(), Map.of() are immutable
6. **SubList Views:** Changes to sublist affect original list
7. **containsKey vs containsValue:** O(1) vs O(n) performance

---

## 🎓 Quick Reference Commands

```java
// Create collections
List.of(1, 2, 3)              // Immutable list
new ArrayList<>(List.of())    // Mutable list
Set.of("a", "b")              // Immutable set  
new HashSet<>(Set.of())       // Mutable set
Map.of("k", "v")              // Immutable map
new HashMap<>()               // Mutable map

// Common operations
list.removeIf(predicate)     // Safe removal
map.computeIfAbsent(k, f)     // Lazy initialization
map.merge(k, v, f)           // Combine values
Collections.frequency(list, x)  // Count occurrences
Collections.binarySearch(sortedList, key)  // Binary search

// Iteration
map.forEach((k, v) -> System.out.println(k + "=" + v));
for (Map.Entry<K, V> entry : map.entrySet()) { /* efficient */ }
```

---

**💡 Interview Tip:** Always mention the underlying data structure and its impact on performance. Real-world analogies make complex concepts memorable!
