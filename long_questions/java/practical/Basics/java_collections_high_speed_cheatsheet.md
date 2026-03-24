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

### Frequency Count (Word Count) - Enhanced Explanation
```java
String[] words = {"apple", "banana", "apple", "cherry", "banana", "apple"};
Map<String, Integer> count = new HashMap<>();
for (String w : words) {
    count.merge(w, 1, Integer::sum);
}
```

**Interview Logic:** `Map.merge()` is like a smart accountant that automatically creates new accounts and updates existing ones in a single operation.

**Enhanced Explanation:**
The `merge()` method elegantly handles three scenarios in one call:

1. **First occurrence:** Key "apple" doesn't exist → inserts `apple=1`
2. **Subsequent occurrences:** Key exists with value → applies `Integer::sum` function
3. **Conditional removal:** If function returns null → entry removed

**Step-by-step execution:**
- First "apple": `merge("apple", 1, sum)` → key absent → stores `apple=1`
- Second "apple": `merge("apple", 1, sum)` → key exists with value 1 → calls `sum(1, 1)` → stores `apple=2`
- Third "apple": `merge("apple", 1, sum)` → key exists with value 2 → calls `sum(2, 1)` → stores `apple=3`

**Why it's better than alternatives:**
- **Traditional:** `count.put(w, count.getOrDefault(w, 0) + 1)` → 2 hash lookups
- **merge():** Single hash lookup + atomic operation
- **Thread-safe advantage:** Prevents race conditions in concurrent scenarios

**Real-world analogy:** Like a voting counter that automatically creates new candidate entries and increments existing votes without manual checking.

### Lazy Initialization (Multimap Pattern)
```java
Map<String, List<Integer>> multiMap = new HashMap<>();
multiMap.computeIfAbsent("key", k -> new ArrayList<>()).add(42);
```
**Interview Logic:** `computeIfAbsent()` is like a lazy landlord who only builds apartments when tenants actually move in.

### LRU Cache Pattern (LinkedHashMap Access Order)
```java
// LRU Cache foundation
Map<Integer, String> lru = new LinkedHashMap<>(16, 0.75f, true);
lru.put(1, "one"); lru.put(2, "two"); lru.put(3, "three");
lru.get(1); // Access key 1 — moves to end (most recently used)
System.out.println(lru.keySet()); // [2, 3, 1] - LRU order
```
**Interview Logic:** LinkedHashMap with access order is like a smart coffee shop that moves frequent customers to the front of the line while keeping track of who hasn't visited recently.

**Enhanced Explanation:**
The `accessOrder=true` parameter transforms LinkedHashMap from insertion-order to access-order mode, making it the perfect foundation for LRU (Least Recently Used) cache implementations.

**Internal Mechanism:**
1. **Initial State:** After `put(1, "one")`, `put(2, "two")`, `put(3, "three")` → `[1, 2, 3]`
2. **Access Operation:** `get(1)` triggers internal `afterNodeAccess()` method
3. **Reordering:** Entry 1 moves from head to tail → `[2, 3, 1]`
4. **LRU Meaning:** Head = least recently used, Tail = most recently used

**Complete LRU Cache Implementation:**
```java
class LRUCache<K, V> extends LinkedHashMap<K, V> {
    private final int maxCapacity;
    
    public LRUCache(int capacity) {
        super(capacity, 0.75f, true); // access-order mode
        this.maxCapacity = capacity;
    }
    
    @Override
    protected boolean removeEldestEntry(Map.Entry<K, V> eldest) {
        return size() > maxCapacity; // Auto-evict LRU when full
    }
}
```

**Why it's interview-worthy:**
- **O(1) access and reordering:** Single hash lookup + list manipulation
- **Automatic eviction:** `removeEldestEntry()` callback handles cache size limits
- **Memory-efficient:** No separate timestamp tracking needed
- **Thread-safety considerations:** Can be wrapped with `Collections.synchronizedMap()`

**Real-world applications:**
- **Browser cache:** Recently visited websites stay in memory
- **Database query cache:** Frequently accessed queries remain cached
- **Image thumbnail cache:** Recently viewed images load instantly

**Performance characteristics:**
- **get/put operations:** O(1) average case
- **Memory overhead:** 2 extra pointers per entry (prev/next)
- **Cache eviction:** O(1) - just remove head of linked list

**Common interview follow-up:** "How would you make this thread-safe?" → Use `ConcurrentHashMap` + manual ordering or wrap with synchronization.

### Balanced Parentheses (Stack Pattern)
```java
Deque<Character> stack = new ArrayDeque<>();
for (char c : expression.toCharArray()) {
    if (c == '(') stack.push(c);
    else if (c == ')' && (stack.isEmpty() || stack.pop() != '(')) return false;
}
return stack.isEmpty();
```
**Interview Logic:** Stack validation is like matching socks — every opening sock needs a matching closing sock in the right order.

### Custom Sorting
```java
// Multi-level comparator
students.sort(Comparator.comparingInt(Student::grade).reversed()
        .thenComparing(Student::name));
```
**Interview Logic:** Chained comparators are like tournament rankings — first by score, then alphabetically for ties.

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

## 💡 Advanced Map Operations Deep Dive

### Map.merge() Variations
```java
// Word counting with merge
Map<String, Integer> wordCount = new HashMap<>();
wordCount.merge("hello", 1, Integer::sum);

// Removing entries conditionally
Map<String, Integer> scores = new HashMap<>(Map.of("alice", 50, "bob", 80));
scores.merge("alice", 0, (old, val) -> old < 60 ? null : old); // Remove if < 60

// Combining maps
Map<String, Integer> map1 = Map.of("a", 1, "b", 2);
Map<String, Integer> map2 = Map.of("b", 3, "c", 4);
map2.forEach((k, v) -> map1.merge(k, v, Integer::sum));
```

### computeIfAbsent vs compute vs merge
```java
// computeIfAbsent: Only when key missing
Map<String, List<String>> groups = new HashMap<>();
groups.computeIfAbsent("team1", k -> new ArrayList<>()).add("member");

// compute: Always recomputes
Map<String, Integer> cache = new HashMap<>();
cache.compute("expensive", (k, v) -> v == null ? calculateExpensive() : v);

// merge: Combine existing with new
Map<String, Integer> totals = new HashMap<>();
totals.merge("sales", 100, Integer::sum); // Add to existing or create new
```

---

**💡 Interview Tip:** Always mention the underlying data structure and its impact on performance. Real-world analogies make complex concepts memorable!
