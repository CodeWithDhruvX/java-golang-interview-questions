# Java Collections Framework — Practical Code Snippets

> **Topics:** List (ArrayList, LinkedList), Set (HashSet, TreeSet, LinkedHashSet), Map (HashMap, TreeMap, LinkedHashMap), Iterator, Comparator, Collections utility class, Queue, Deque, PriorityQueue

---

## 📋 Reading Progress

- [ ] **Section 1:** List — ArrayList & LinkedList (Q1–Q18)
- [ ] **Section 2:** Set — HashSet, TreeSet, LinkedHashSet (Q19–Q32)
- [ ] **Section 3:** Map — HashMap, TreeMap, LinkedHashMap (Q33–Q52)
- [ ] **Section 4:** Queue, Deque & PriorityQueue (Q53–Q62)
- [ ] **Section 5:** Collections Utility & Iterators (Q63–Q75)

> 🔖 **Last read:** <!-- -->

---

## Section 1: List (Q1–Q18)

### 1. ArrayList vs LinkedList — remove() Performance
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> arr = new ArrayList<>(List.of(1, 2, 3, 4, 5));
        arr.remove(0); // O(n) — shifts elements
        System.out.println(arr);

        List<Integer> linked = new LinkedList<>(List.of(1, 2, 3, 4, 5));
        linked.remove(0); // O(1) — pointer update
        System.out.println(linked);
    }
}
```
**A:**
```
[2, 3, 4, 5]
[2, 3, 4, 5]
```
Same result but different performance. `ArrayList.remove(0)` is O(n) (shifts). `LinkedList.remove(0)` is O(1) (pointer update). However, LinkedList has higher memory overhead (node objects).

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the performance difference between ArrayList and LinkedList?"

**Your Response:** "The output shows both lists as '[2, 3, 4, 5]' after removing the first element, but the performance characteristics are completely different. ArrayList's `remove(0)` is O(n) because it has to shift all remaining elements left by one position using `System.arraycopy()`. LinkedList's `remove(0)` is O(1) because it just updates a couple of pointers - the first node reference. However, LinkedList has higher memory overhead since each element is wrapped in a node object with prev/next pointers. For frequent insertions/deletions at the beginning, LinkedList is better. For random access and most operations, ArrayList is usually preferred due to better cache locality."

**Code Snippet Internal Behavior:**
- ArrayList uses dynamic array `Object[] elementData` internally
- `remove(0)` triggers `System.arraycopy()` to shift all elements left by 1 position
- LinkedList uses doubly-linked nodes with `Node<E> first`, `Node<E> last`
- Each node contains `E item`, `Node<E> next`, `Node<E> prev`
- `remove(0)` just updates `first = first.next` and adjusts pointers
- ArrayList memory: contiguous array, cache-friendly
- LinkedList memory: separate node objects, more GC pressure

---

### 2. List.remove(int) vs List.remove(Object)
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>(List.of(10, 20, 30));
        list.remove(1);                   // remove by INDEX
        System.out.println(list);

        list.remove(Integer.valueOf(10)); // remove by VALUE
        System.out.println(list);
    }
}
```
**A:**
```
[10, 30]
[30]
```
`remove(int)` removes by index. `remove(Object)`/`remove(Integer.valueOf(...))` removes by value. Autoboxing pitfall!

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why do the remove calls behave differently?"

**Your Response:** "The output is '[10, 30]' then '[30]'. This demonstrates a classic Java autoboxing gotcha. When we call `list.remove(1)`, Java calls the `remove(int index)` method, removing the element at index 1 (which is 20). When we call `list.remove(Integer.valueOf(10))`, Java calls the `remove(Object o)` method, removing the element with value 10. The difference is that primitive `int` binds to the index-based method, while the `Integer` object binds to the value-based method. This is a common source of bugs - developers might expect `list.remove(1)` to remove the value 1, but it actually removes the element at index 1."

**Code Snippet Internal Behavior:**
- ArrayList has two overloaded `remove()` methods: `remove(int index)` and `remove(Object o)`
- `remove(1)` calls `remove(int index)` - removes element at index 1 (value 20)
- `remove(Integer.valueOf(10))` calls `remove(Object o)` due to wrapper type
- Compiler chooses method based on compile-time type, not runtime value
- Primitive `int` binds to `remove(int index)`
- `Integer` object binds to `remove(Object o)`
- Common bug: `list.remove(1)` vs `list.remove(Integer.valueOf(1))` behave differently

---

### 3. ConcurrentModificationException
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>(List.of("a", "b", "c"));
        for (String s : list) {
            if (s.equals("a")) list.remove(s); // modifying while iterating!
        }
    }
}
```
**A:** **ConcurrentModificationException at runtime.** You cannot modify a collection while iterating with a for-each loop. Use `Iterator.remove()` or `list.removeIf()` instead.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What happens and why does this exception occur?"

**Your Response:** "This throws a ConcurrentModificationException at runtime. The issue is that we're modifying the list while iterating over it using a for-each loop. The for-each loop internally uses an iterator that tracks modifications. When we call `list.remove("a")`, the list's modification count changes, but the iterator still expects the old count. On the next iteration, the iterator detects this mismatch and throws the exception. This is Java's 'fail-fast' behavior - it detects concurrent modification early rather than allowing undefined behavior. The fix is to use either `Iterator.remove()` or the more modern `list.removeIf()` method."

**Code Snippet Internal Behavior:**
- For-each loop uses `Iterator<String> it = list.iterator()` internally
- ArrayList iterator maintains `int expectedModCount = modCount`
- `modCount` increments on every structural modification (add/remove)
- Iterator checks `expectedModCount == actualModCount` before each `next()` call
- `list.remove("a")` increments `modCount` but `expectedModCount` stays unchanged
- Next `hasNext()`/`next()` call detects mismatch → throws ConcurrentModificationException
- This is "fail-fast" behavior - detects concurrent modification early

---

### 4. Correct Way — removeIf
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>(List.of("apple", "banana", "cherry"));
        list.removeIf(s -> s.startsWith("b"));
        System.out.println(list);
    }
}
```
**A:** `[apple, cherry]`. `removeIf` is safe and concise.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does removeIf solve the ConcurrentModificationException?"

**Your Response:** "The output is '[apple, cherry]'. The `removeIf()` method is the modern, safe way to remove elements from a collection while iterating. It uses an internal iterator that properly handles modification tracking, so it doesn't throw ConcurrentModificationException. The method takes a Predicate that returns true for elements to remove. Internally, it uses efficient batch removal rather than removing elements one by one. This is much cleaner than writing manual iterator code and is the preferred approach in Java 8+ for conditional removal. It's also more readable and less error-prone."

**Code Snippet Internal Behavior:**
- `removeIf()` uses internal iterator with proper modification tracking
- Iterator's `remove()` method updates both `modCount` and `expectedModCount`
- Implementation creates temporary array of elements to remove
- Applies batch removal to avoid ConcurrentModificationException
- More efficient than manual iterator removal for multiple elements
- Uses Predicate<T> functional interface for condition evaluation
- Internally: `for (int i=0; i<size; i++) if (filter.test(element)) remove(i)`

---

### 5. Correct Way — Iterator.remove()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>(List.of(1, 2, 3, 4, 5));
        Iterator<Integer> it = list.iterator();
        while (it.hasNext()) {
            if (it.next() % 2 == 0) it.remove(); // safe removal
        }
        System.out.println(list);
    }
}
```
**A:** `[1, 3, 5]`

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Iterator.remove() work safely?"

**Your Response:** "The output is '[1, 3, 5]'. This shows the traditional way to safely remove elements while iterating using an Iterator. The key is that we call `it.next()` to get the element, check the condition, and then call `it.remove()` to remove the last returned element. The iterator keeps track of modifications internally, so it doesn't throw ConcurrentModificationException. The rule is you must call `next()` before `remove()` - otherwise it throws IllegalStateException. While this works, `removeIf()` is often preferred in modern Java for its simplicity."

**Code Snippet Internal Behavior:**
- Iterator maintains cursor position and expectedModCount
- `it.next()` advances cursor and returns current element
- `it.remove()` removes last returned element and updates expectedModCount
- Must call `next()` before `remove()` - throws IllegalStateException otherwise
- Single iterator cannot be used concurrently by multiple threads
- Iterator removal is O(1) - just shifts elements from removeIndex+1 to end
- After removal, cursor points to element after removed one

---

### 6. List.of() — Immutable
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = List.of("a", "b", "c");
        list.add("d"); // try to add to immutable list
    }
}
```
**A:** **UnsupportedOperationException at runtime.** `List.of()` (Java 9+) returns an immutable list. Use `new ArrayList<>(List.of(...))` to get a mutable copy.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What happens and why does List.of() throw an exception?"

**Your Response:** "This throws an UnsupportedOperationException at runtime. The issue is that `List.of()` creates an immutable list - you cannot add, remove, or modify elements after creation. This is by design for creating constant collections that won't change. If you need a mutable list, you should use `new ArrayList<>(List.of(...))` which creates a mutable copy. Immutable collections are great for returning from methods, creating constants, or ensuring thread safety without copying. They're also more memory-efficient since they don't need to support modifications."

**Code Snippet Internal Behavior:**
- `List.of()` returns `ImmutableCollections.ListN` or `ImmutableCollections.SubList`
- These are internal classes in `java.util.ImmutableCollections` package
- All modification methods throw `UnsupportedOperationException`
- Immutable lists store elements in final array `final E[] elements`
- No defensive copying needed - elements array is internal and never exposed
- Null elements not allowed - throws `NullPointerException` during creation
- Optimized for small lists: `List.of()` (empty), `List.of(e1)`, `List.of(e1, e2)` etc.

---

### 7. subList() — Backed by Original
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>(List.of(1, 2, 3, 4, 5));
        List<Integer> sub = list.subList(1, 4); // [2, 3, 4]
        sub.clear();
        System.out.println(list);
    }
}
```
**A:** `[1, 5]`. `subList()` returns a **view** backed by the original list. Modifying the subList modifies the original.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about subList()?"

**Your Response:** "The output is '[1, 5]'. This shows that `subList()` returns a view backed by the original list, not a copy. When we call `sub.clear()` on the subList, it actually removes elements from the original list. The subList is just a window into a portion of the original list. This is memory-efficient since it doesn't copy elements, but it means modifications to either list affect the other. This is useful for working with portions of large lists, but you need to be aware that structural changes to the original list will invalidate the subList."

**Code Snippet Internal Behavior:**
- `subList()` returns `ArrayList.SubList` inner class instance
- SubList holds reference to parent ArrayList and offset/size
- SubList operations delegate to parent with index translation: `parent.get(offset + index)`
- `subList.clear()` calls `parent.removeRange(fromIndex, toIndex)`
- No data copying - just index mapping for performance
- Structural changes to parent invalidate subList (ConcurrentModificationException)
- Memory efficient: shares underlying element array

---

### 8. Collections.sort() vs List.sort()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>(List.of("banana", "apple", "cherry"));
        Collections.sort(list);
        System.out.println(list);

        list.sort(Comparator.reverseOrder());
        System.out.println(list);
    }
}
```
**A:**
```
[apple, banana, cherry]
[cherry, banana, apple]
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between Collections.sort() and List.sort()?"

**Your Response:** "The output shows the list sorted naturally, then in reverse order. `Collections.sort(list)` is the older way to sort, while `list.sort(comparator)` was added in Java 8. Under the hood, `Collections.sort()` actually calls `list.sort()` anyway. The `List.sort()` method is more flexible since you can pass any comparator directly. Both use the efficient TimSort algorithm. The modern approach is to use `list.sort()` - it's more object-oriented and readable. Both methods sort the list in-place, they don't create a new list."

**Code Snippet Internal Behavior:**
- `Collections.sort(list)` uses `list.sort(null)` internally (since Java 8)
- Calls `Arrays.sort(elementArray)` - uses optimized TimSort algorithm
- `list.sort(comparator)` directly calls `Arrays.sort(elementArray, comparator)`
- TimSort is hybrid of merge sort and insertion sort - O(n log n) worst case
- For ArrayList: sorts backing array directly, then updates `modCount`
- Comparator.reverseOrder() returns `Collections.ReverseComparator`
- Reverse comparator flips comparison: `-(c1.compareTo(c2))`

---

### 9. Collections.unmodifiableList()
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> mutable = new ArrayList<>(List.of("a", "b"));
        List<String> immutable = Collections.unmodifiableList(mutable);
        mutable.add("c"); // modifies underlying list!
        System.out.println(immutable); // reflects the change!
        immutable.add("d"); // throws exception
    }
}
```
**A:** `[a, b, c]` then **UnsupportedOperationException**. `unmodifiableList` wraps the list — mutations via the original reference still show through. It only prevents direct mutation of the wrapper.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does unmodifiableList work?"

**Your Response:** "The output is '[a, b, c]' then an UnsupportedOperationException. This demonstrates that `Collections.unmodifiableList()` creates a wrapper that prevents modification through the wrapper reference, but the original list is still mutable. When we add 'c' to the original list, it shows up in the unmodifiable view because they share the same backing data. But when we try to modify through the unmodifiable wrapper, it throws an exception. This is the wrapper pattern - it provides read-only access without making a defensive copy. It's useful for API design when you want to prevent external modification."

**Code Snippet Internal Behavior:**
- `Collections.unmodifiableList()` returns `UnmodifiableList` wrapper
- Wrapper holds reference to original list: `final List<? extends E> list`
- All modification methods throw `UnsupportedOperationException`
- Read operations delegate to wrapped list: `return list.get(index)`
- Changes to original list are visible through wrapper (no defensive copy)
- Wrapper pattern - provides read-only view of mutable collection
- Useful for API design - prevent external modification but allow internal changes

---

### 10. ArrayList Initial Capacity vs size
**Q: What is the output?**
```java
import java.util.*;
import java.lang.reflect.*;
public class Main {
    public static void main(String[] args) {
        ArrayList<Integer> list = new ArrayList<>(100); // initial capacity 100
        System.out.println(list.size());
    }
}
```
**A:** `0`. Initial capacity is an internal optimization hint — it pre-allocates the backing array but doesn't affect `size()`. `size()` returns the number of actual elements.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between initial capacity and size?"

**Your Response:** "The output is '0'. This shows that initial capacity is different from size. When we create `new ArrayList<>(100)`, we're telling Java to pre-allocate space for 100 elements, but we haven't actually added any elements yet. The `size()` method returns the actual number of elements, which is 0. Initial capacity is just an optimization to avoid frequent array resizing. The default capacity is 10, and it grows by 1.5x when needed. Setting initial capacity is useful when you know roughly how many elements you'll add - it prevents multiple array copies."

**Code Snippet Internal Behavior:**
- `new ArrayList(100)` allocates `Object[] elementData = new Object[100]`
- `size` field remains 0 - tracks actual element count, not capacity
- Capacity is internal optimization to avoid frequent array resizing
- Default capacity is 10, grows by `oldCapacity + (oldCapacity >> 1)` (1.5x)
- When `size >= elementData.length`, `grow()` method allocates new array
- `Arrays.copyOf(elementData, newCapacity)` copies elements to larger array
- Initial capacity reduces array copies but increases initial memory usage

---

### 11. List.copyOf() — Immutable Deep Copy
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> mutable = new ArrayList<>(List.of("x", "y"));
        List<String> copy = List.copyOf(mutable);
        mutable.add("z");
        System.out.println(copy.size()); // unaffected by mutable change
    }
}
```
**A:** `2`. `List.copyOf()` creates an independent immutable snapshot.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does List.copyOf() differ from unmodifiableList?"

**Your Response:** "The output is '2'. This shows that `List.copyOf()` creates a completely independent immutable copy. When we modify the original list by adding 'z', the copy remains unchanged with size 2. Unlike `Collections.unmodifiableList()` which creates a wrapper around the original list, `List.copyOf()` actually copies the elements to a new immutable collection. This means changes to the original don't affect the copy at all. It's useful when you need to guarantee immutability even if the original collection changes."

**Code Snippet Internal Behavior:**
- `List.copyOf()` returns `ImmutableCollections.ListN` or `ImmutableCollections.SubList`
- Creates defensive copy of source collection: `Arrays.copyOf(collection.toArray(), n)`
- If source is already immutable, may return same instance (optimization)
- Null elements not allowed - throws `NullPointerException` if any element is null
- Copy is completely independent - changes to original don't affect copy
- Uses `collection.toArray()` then `Arrays.copyOf()` for type safety
- More efficient than `new ArrayList<>(collection).asList()` for immutable result

---

### 12. indexOf and lastIndexOf
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = List.of("a", "b", "a", "c", "a");
        System.out.println(list.indexOf("a"));
        System.out.println(list.lastIndexOf("a"));
        System.out.println(list.indexOf("z"));
    }
}
```
**A:**
```
0
4
-1
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do indexOf and lastIndexOf work?"

**Your Response:** "The output shows '0', '4', then '-1'. `indexOf('a')` returns the first occurrence index, which is 0. `lastIndexOf('a')` returns the last occurrence index, which is 4. `indexOf('z')` returns -1 because 'z' doesn't exist in the list. Both methods perform linear searches - O(n) time complexity. They use `Objects.equals()` for comparison, so they handle null values safely. For large lists with frequent lookups, you'd want to use a HashSet for O(1) performance instead."

**Code Snippet Internal Behavior:**
- `indexOf()` performs linear search: `for (int i=0; i<size; i++) if (Objects.equals(o, elementData[i])) return i;`
- `lastIndexOf()` searches backwards: `for (int i=size-1; i>=0; i--) if (Objects.equals(o, elementData[i])) return i;`
- Both use `Objects.equals()` for null-safe comparison
- Time complexity: O(n) - must scan potentially entire list
- Returns -1 if element not found (convention for "not found" in Java)
- For large lists with frequent lookups, consider `HashSet` for O(1) lookup
- Uses `==` comparison first, then `equals()` for performance

---

### 13. Sorting with Chained Comparators
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    record Student(String name, int grade) {}
    public static void main(String[] args) {
        List<Student> students = new ArrayList<>(List.of(
            new Student("Alice", 90), new Student("Bob", 85),
            new Student("Charlie", 90), new Student("Dave", 85)));
        students.sort(Comparator.comparingInt(Student::grade).reversed()
                .thenComparing(Student::name));
        students.forEach(s -> System.out.println(s.name() + " " + s.grade()));
    }
}
```
**A:**
```
Alice 90
Charlie 90
Bob 85
Dave 85
```
Sort by grade descending, then by name ascending.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do chained comparators work?"

**Your Response:** "The output shows the students sorted by grade descending, then by name ascending. The key is the chained comparator: `Comparator.comparingInt(Student::grade).reversed().thenComparing(Student::name)`. First, it sorts by grade in descending order using `reversed()`. For students with the same grade (Alice and Charlie both have 90), it uses the secondary comparator to sort by name alphabetically. Chained comparators use short-circuit evaluation - they only move to the next comparator if the previous one returns 0 (equal). This is perfect for multi-level sorting requirements."

**Code Snippet Internal Behavior:**
- `Comparator.comparingInt(Student::grade)` creates key extractor comparator
- `.reversed()` inverts comparison order: `c2 - c1` instead of `c1 - c2`
- `.thenComparing(Student::name)` chains secondary comparator
- Chained comparators use short-circuit evaluation
- `Student::grade` and `Student::name` are method references to record accessors
- Comparator returns negative if first < second, zero if equal, positive if first > second
- `Collections.sort()` uses TimSort algorithm internally

---

### 14. LinkedList as Deque
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Deque<Integer> deque = new LinkedList<>();
        deque.addFirst(1);
        deque.addLast(2);
        deque.addFirst(0);
        System.out.println(deque);
        System.out.println(deque.peekFirst());
        System.out.println(deque.pollLast());
        System.out.println(deque);
    }
}
```
**A:**
```
[0, 1, 2]
0
2
[0, 1]
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does LinkedList work as a Deque?"

**Your Response:** "The output shows the deque operations in action. LinkedList implements the Deque interface, so it can be used as a double-ended queue. `addFirst(0)` adds to the front, `addLast(2)` adds to the back, `addFirst(0)` adds another to front. So we get '[0, 1, 2]'. `peekFirst()` looks at the front element without removing it (0). `pollLast()` removes and returns the last element (2). The final list is '[0, 1]'. All these operations are O(1) for LinkedList since it just updates pointers, unlike ArrayList which might need to shift elements."

**Code Snippet Internal Behavior:**
- `LinkedList` implements `Deque` interface with doubly-linked nodes
- `addFirst(1)` creates new node and updates head: `Node<E> newNode = new Node<>(null, item, first)`
- `addLast(2)` updates tail: `Node<E> newNode = new Node<>(last, item, null)`
- Each node has `E item`, `Node<E> prev`, `Node<E> next`
- `peekFirst()` returns `first.item` without removing node
- `pollLast()` removes tail node and updates `last = last.prev`
- All operations are O(1) - just pointer updates, no array copying
- Memory overhead: 3 references + object header per element

---

### 15. Collections.frequency()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = List.of("a", "b", "a", "c", "a");
        System.out.println(Collections.frequency(list, "a"));
        System.out.println(Collections.frequency(list, "z"));
    }
}
```
**A:**
```
3
0
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does Collections.frequency() do?"

**Your Response:** "The output shows '3' then '0'. `Collections.frequency(list, 'a')` counts how many times 'a' appears in the list, which is 3 times. `Collections.frequency(list, 'z')` returns 0 because 'z' doesn't exist. This is a simple utility method that iterates through the collection and counts matches using `Objects.equals()`. The time complexity is O(n) since it has to scan the entire collection. For large datasets with frequent frequency queries, you'd be better off using a Map to store counts for O(1) lookup."

**Code Snippet Internal Behavior:**
- `Collections.frequency()` iterates through collection: `int count = 0; for (E e : c) if (Objects.equals(o, e)) count++;`
- Uses `Objects.equals()` for null-safe comparison
- Time complexity: O(n) - linear scan of entire collection
- For large collections with frequent frequency queries, consider `Map<T, Integer>`
- Alternative: `list.stream().filter(x -> Objects.equals(x, target)).count()` (Java 8+)
- No built-in optimization - always scans entire collection
- Works with any `Collection` implementation, not just lists

---

### 16. Collections.nCopies()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = Collections.nCopies(5, "hello");
        System.out.println(list);
        System.out.println(list.size());
    }
}
```
**A:**
```
[hello, hello, hello, hello, hello]
5
```
`Collections.nCopies()` creates an immutable list with n copies of the specified object.

**Code Snippet Internal Behavior:**
- Returns `ImmutableCollections.ListN` with single object reference
- Doesn't actually create n copies - stores one reference and returns it for each index
- `get(int index)` always returns the same object reference: `return element`
- Memory efficient: O(1) space regardless of n (stores count + single reference)
- All operations O(1) except `contains()` which is O(n)
- Immutable - all modification methods throw `UnsupportedOperationException`
- Useful for creating repeated elements or padding collections

---

### 17. Collections.shuffle() and reverse()
**Q: Does this compile?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>(List.of(1, 2, 3, 4, 5));
        Collections.reverse(list);
        System.out.println(list);

        Collections.shuffle(list, new Random(42));
        System.out.println(list.size()); // still same size
    }
}
```
**A:**
```
[5, 4, 3, 2, 1]
5
```

**Code Snippet Internal Behavior:**
- `Collections.reverse(list)` uses two-pointer swap algorithm
- Implementation: `for (int i=0, mid=list.size()/2, j=list.size()-1; i<mid; i++, j--) swap(list, i, j)`
- `Collections.shuffle(list, random)` uses Fisher-Yates shuffle algorithm
- Fisher-Yates: `for (int i=list.size()-1; i>0; i--) swap(list, i, random.nextInt(i+1))`
- Both methods modify list in-place - no new list created
- Time complexity: O(n) for both operations
- `Random(42)` provides reproducible shuffle for testing
- Uses `list.set(i, list.set(j, list.get(i)))` for atomic swapping

---

### 18. Arrays.asList() — Fixed Size but Mutable
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = Arrays.asList("a", "b", "c");
        list.set(0, "x"); // OK — can modify existing elements
        System.out.println(list);
        list.add("d");    // UnsupportedOperationException — can't resize!
    }
}
```
**A:** `[x, b, c]` then **UnsupportedOperationException**. `Arrays.asList` returns a fixed-size list backed by an array — you can set elements but not add/remove.

**Code Snippet Internal Behavior:**
- `Arrays.asList()` returns `java.util.Arrays.ArrayList` (inner class, not `java.util.ArrayList`)
- This inner class stores reference to original array: `private final E[] a`
- `get(index)` returns `a[index]`, `set(index, element)` assigns `a[index] = element`
- `add()` and `remove()` throw `UnsupportedOperationException` - array size fixed
- Size always equals array length - cannot grow or shrink
- Changes to list are reflected in original array (same backing storage)
- Useful for converting arrays to lists for API compatibility
- Different from `new ArrayList<>(Arrays.asList(array))` which copies elements

---

## Section 2: Set (Q19–Q32)

### 19. HashSet — No Duplicates, Unordered
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<Integer> set = new HashSet<>(List.of(3, 1, 4, 1, 5, 9, 2, 6));
        System.out.println(set.size());
        // order is unpredictable
    }
}
```
**A:** `7`. HashSet removes duplicates (one `1` dropped). Order is unspecified — do not rely on it.

**Code Snippet Internal Behavior:**
- `HashSet` internally uses `HashMap<E, Object>` for storage
- Each element stored as key in backing map with dummy `Object` value: `PRESENT = new Object()`
- `add(e)` calls `map.put(e, PRESENT)` and returns `map.put()` result
- Duplicate detection: `HashMap` checks `hashCode()` first, then `equals()` if hash matches
- Initial capacity 16, load factor 0.75 by default
- When size > capacity * loadFactor, triggers rehashing: doubles capacity and rehashes all entries
- Order depends on hash values and internal table structure - not guaranteed
- `size()` returns `map.size()` - actual number of unique elements

---

### 20. TreeSet — Sorted, Natural Ordering
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<Integer> set = new TreeSet<>(List.of(5, 3, 1, 4, 2));
        System.out.println(set);
        System.out.println(((TreeSet<Integer>)set).first());
        System.out.println(((TreeSet<Integer>)set).last());
    }
}
```
**A:**
```
[1, 2, 3, 4, 5]
1
5
```

**Code Snippet Internal Behavior:**
- `TreeSet` uses `TreeMap<E, Object>` internally with dummy value `PRESENT`
- Implements `NavigableMap` interface for ordered operations
- Elements must be `Comparable` or provide `Comparator` at construction
- `first()` returns `map.firstKey()`, `last()` returns `map.lastKey()`
- Uses Red-Black Tree data structure - self-balancing binary search tree
- All operations O(log n): add, remove, contains, first, last
- Natural ordering: `e1.compareTo(e2)` for `Comparable` elements
- Tree maintains invariants: no red node has red child, every path to leaf has same black count

---

### 21. LinkedHashSet — Insertion Order Preserved
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<String> set = new LinkedHashSet<>(List.of("banana", "apple", "cherry", "apple"));
        System.out.println(set);
    }
}
```
**A:** `[banana, apple, cherry]`. `LinkedHashSet` maintains insertion order and removes duplicates.

**Code Snippet Internal Behavior:**
- `LinkedHashSet` extends `HashSet` and uses `LinkedHashMap<E, Object>` internally
- Maintains doubly-linked list of entries in addition to hash table
- Each entry has `before` and `after` pointers for order maintenance
- `add(e)` appends to end of linked list if not already present
- Iteration follows linked list order, not hash table order
- Slightly more memory overhead than `HashSet` (2 extra pointers per entry)
- Still O(1) average for add, remove, contains operations
- Useful when iteration order matters but need Set semantics
- Rehashing preserves insertion order by rebuilding linked list

---

### 22. Custom Object in HashSet — hashCode Required
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static class Point {
        int x, y;
        Point(int x, int y) { this.x = x; this.y = y; }
        // equals() and hashCode() NOT overridden
    }
    public static void main(String[] args) {
        Set<Point> set = new HashSet<>();
        set.add(new Point(1, 2));
        set.add(new Point(1, 2)); // different object, same data
        System.out.println(set.size());
    }
}
```
**A:** `2`. Without overriding `hashCode()` and `equals()`, the two `Point` objects are treated as different (default identity-based comparison).

**Code Snippet Internal Behavior:**
- `HashSet` uses `hashCode()` to find bucket, then `equals()` to check for duplicates
- Default `Object.hashCode()` returns unique memory address for each object instance
- Default `Object.equals()` uses `==` reference comparison
- Two `new Point(1, 2)` objects have different hash codes and different references
- `add(point1)` stores in bucket based on `point1.hashCode()`
- `add(point2)` goes to different bucket (different hash) → no collision check
- Fix: Override `hashCode()` to return consistent hash for equal coordinates
- Fix: Override `equals()` to compare x and y fields instead of references
- Contract: if `equals()` true, `hashCode()` must return same value

---

### 23. Set.of() — No Duplicates Allowed
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<String> set = Set.of("a", "b", "a"); // duplicate!
    }
}
```
**A:** **IllegalArgumentException at runtime.** `Set.of()` throws an exception if duplicate elements are provided.

**Code Snippet Internal Behavior:**
- `Set.of()` performs duplicate detection during creation
- Uses `Object.equals()` to check for duplicates in varargs array
- Throws `IllegalArgumentException` with message "duplicate element: a"
- Returns `ImmutableCollections.SetN` or `ImmutableCollections.Set12` for small sets
- Null elements not allowed - throws `NullPointerException`
- Immutable set - all modification methods throw `UnsupportedOperationException`
- Optimized implementations: `Set.of()` (empty), `Set.of(e1)`, `Set.of(e1, e2)` etc.
- Uses `SetN` for >2 elements, `Set1`/`Set2` for 1-2 elements

---

### 24. Set addAll() — Union
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<Integer> a = new HashSet<>(Set.of(1, 2, 3));
        Set<Integer> b = new HashSet<>(Set.of(3, 4, 5));
        a.addAll(b); // union
        System.out.println(new TreeSet<>(a));
    }
}
```
**A:** `[1, 2, 3, 4, 5]`

**Code Snippet Internal Behavior:**
- `addAll()` calls `map.putAll()` internally (HashSet uses HashMap)
- For each element in collection: `add(e)` which calls `map.put(e, PRESENT)`
- Duplicate elements automatically ignored by HashMap put logic
- Time complexity: O(n) where n is size of collection being added
- Hash table may resize if adding many elements beyond load factor threshold
- Returns `true` if set changed (new elements added), `false` if all duplicates
- Union operation: result contains all unique elements from both sets
- More efficient than manual iteration and individual `add()` calls

---

### 25. Set retainAll() — Intersection
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<Integer> a = new HashSet<>(Set.of(1, 2, 3, 4));
        Set<Integer> b = new HashSet<>(Set.of(3, 4, 5, 6));
        a.retainAll(b); // intersection
        System.out.println(new TreeSet<>(a));
    }
}
```
**A:** `[3, 4]`

**Code Snippet Internal Behavior:**
- `retainAll()` uses iterator to remove elements not in specified collection
- Implementation: `Iterator<E> it = iterator(); while (it.hasNext()) if (!c.contains(it.next())) it.remove();`
- For each element, checks if it exists in collection `b` using `contains()`
- Removes elements from set `a` that are not present in set `b`
- Time complexity: O(n * m) where n=size(a), m=size(b) for contains check
- Uses fail-fast iterator - throws ConcurrentModificationException if set modified during operation
- Returns `true` if set changed, `false` if no elements removed
- Intersection operation: result contains only elements present in both sets

---

### 26. Set removeAll() — Difference
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<Integer> a = new HashSet<>(Set.of(1, 2, 3, 4));
        Set<Integer> b = new HashSet<>(Set.of(3, 4, 5, 6));
        a.removeAll(b); // difference
        System.out.println(new TreeSet<>(a));
    }
}
```
**A:** `[1, 2]`

**Code Snippet Internal Behavior:**
- `removeAll()` removes all elements that are also in specified collection
- Implementation: `Iterator<E> it = iterator(); while (it.hasNext()) if (c.contains(it.next())) it.remove();`
- For each element in set `a`, checks if it exists in set `b`
- Removes element from `a` if found in `b`
- Time complexity: O(n * m) where n=size(a), m=size(b)
- Uses iterator with safe removal to avoid ConcurrentModificationException
- Returns `true` if set changed (elements removed), `false` if no overlap
- Set difference operation: result contains elements in `a` but not in `b`
- More efficient than creating new set and manually filtering

---

### 27. TreeSet with Custom Comparator
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        // sort strings by length, then alphabetically
        Set<String> set = new TreeSet<>(Comparator.comparingInt(String::length).thenComparing(Comparator.naturalOrder()));
        set.addAll(List.of("banana", "fig", "apple", "kiwi", "plum"));
        System.out.println(set);
    }
}
```
**A:** `[fig, kiwi, plum, apple, banana]`

**Code Snippet Internal Behavior:**
- `TreeSet` constructor accepts custom `Comparator<String>`
- `Comparator.comparingInt(String::length)` creates comparator based on string length
- `.thenComparing(Comparator.naturalOrder())` chains secondary alphabetical comparison
- Comparator returns: negative if first < second, zero if equal, positive if first > second
- Red-Black tree uses comparator for all ordering decisions (insertion, search)
- Elements with same length ordered alphabetically as tie-breaker
- All tree operations O(log n) due to balanced tree property
- Comparator stored internally and used consistently throughout tree lifecycle
- No need for elements to implement `Comparable` when custom comparator provided

---

### 28. EnumSet — Efficient Set for Enums
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    enum Day { MON, TUE, WED, THU, FRI, SAT, SUN }
    public static void main(String[] args) {
        EnumSet<Day> weekend = EnumSet.of(Day.SAT, Day.SUN);
        EnumSet<Day> workweek = EnumSet.complementOf(weekend);
        System.out.println(weekend);
        System.out.println(workweek.size());
    }
}
```
**A:**
```
[SAT, SUN]
5
```

**Code Snippet Internal Behavior:**
- `EnumSet` is specialized Set implementation for enum types
- Uses bit vector internally - each enum constant maps to a bit position
- Extremely memory efficient: one long (64 bits) can store up to 64 enum values
- `EnumSet.of(Day.SAT, Day.SUN)` sets corresponding bits: `bits |= (1L << SAT.ordinal())`
- `complementOf()` flips all bits: `bits = ~bits & universeMask`
- Operations are O(1) - just bit manipulation, no hashing or comparisons
- Iterator returns elements in natural enum order (ordinal order)
- Fails fast if you try to add non-enum elements (compile-time type safety)
- More efficient than `HashSet` for enum keys due to bit operations

---

### 29. contains() on a List vs Set — Performance
**Q: Which is faster and why?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>();
        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < 1_000_000; i++) { list.add(i); set.add(i); }

        // Which is faster?
        System.out.println(list.contains(999_999)); // O(n) scan
        System.out.println(set.contains(999_999));  // O(1) hash lookup
    }
}
```
**A:** Both print `true`. But `set.contains()` is **O(1)** average. `list.contains()` is **O(n)** — scans every element. Always use `HashSet` for frequent membership checks.

**Code Snippet Internal Behavior:**
- `ArrayList.contains()` performs linear search: `for (int i=0; i<size; i++) if (Objects.equals(o, elementData[i])) return true;`
- `HashSet.contains()` computes hash: `int hash = Objects.hashCode(o); int index = (n-1) & hash; return getNode(hash, o) != null`
- ArrayList: worst-case O(n) comparisons, best-case O(1) if element at index 0
- HashSet: average O(1) hash lookup, worst-case O(n) if all elements in same bucket
- HashSet uses `hashCode()` to find bucket, then `equals()` for collision resolution
- Performance difference grows dramatically with collection size
- Memory tradeoff: HashSet uses more memory but faster lookups
- For 1M elements: ArrayList may need up to 1M comparisons, HashSet typically 1-3

---

### 30. TreeSet headSet, tailSet, subSet
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        TreeSet<Integer> set = new TreeSet<>(List.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10));
        System.out.println(set.headSet(5));    // < 5
        System.out.println(set.tailSet(7));    // >= 7
        System.out.println(set.subSet(3, 7)); // [3, 7)
    }
}
```
**A:**
```
[1, 2, 3, 4]
[7, 8, 9, 10]
[3, 4, 5, 6]
```

**Code Snippet Internal Behavior:**
- `TreeSet` extends `SortedSet` and provides range view operations
- `headSet(5)` returns view of elements < 5 (exclusive upper bound)
- `tailSet(7)` returns view of elements ≥ 7 (inclusive lower bound)
- `subSet(3, 7)` returns view of elements [3, 7) - inclusive start, exclusive end
- Returns `TreeSet.SubSet` which is a view, not a copy
- Changes to original set reflected in sub-sets and vice versa
- Uses Red-Black tree structure for efficient range operations O(log n)
- Sub-sets share underlying tree - no data duplication
- Useful for pagination, data partitioning, and range queries
- All range operations maintain sorted order automatically

---

### 31. Set Iteration Order
**Q: What guarantees does each Set type provide?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        // HashSet: no order guarantee
        // LinkedHashSet: insertion order preserved
        // TreeSet: sorted (natural or custom Comparator)
        Set<String> linked = new LinkedHashSet<>(List.of("c", "a", "b"));
        Set<String> tree   = new TreeSet<>(List.of("c", "a", "b"));
        System.out.println(linked);
        System.out.println(tree);
    }
}
```
**A:**
```
[c, a, b]
[a, b, c]
```

**Code Snippet Internal Behavior:**
- `LinkedHashSet` iteration follows doubly-linked list order (insertion order)
- `TreeSet` iteration follows in-order traversal of Red-Black tree (sorted order)
- `HashSet` iteration order depends on hash table internal structure and bucket order
- LinkedHashSet: maintains `head` and `tail` pointers, iteration follows `after` pointers
- TreeSet: in-order traversal visits left subtree, then node, then right subtree
- HashSet: iteration order can change between runs and JVM versions
- Order guarantees affect performance: LinkedHashSet/TreeSet have extra overhead
- Choose Set type based on whether iteration order matters for your use case

---

### 32. Set.copyOf() — Immutable
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<String> mutable = new HashSet<>(Set.of("x", "y"));
        Set<String> copy = Set.copyOf(mutable);
        copy.add("z");
    }
}
```
**A:** **UnsupportedOperationException.** `Set.copyOf()` returns an immutable set.

**Code Snippet Internal Behavior:**
- `Set.copyOf()` returns `ImmutableCollections.SetN` or optimized variants
- Creates defensive copy of source collection using `collection.toArray()`
- For small sets (0-2 elements), uses specialized `Set0`, `Set1`, `Set2` classes
- For larger sets, uses `SetN` which stores elements in array
- All modification methods throw `UnsupportedOperationException`
- Null elements not allowed - throws `NullPointerException` if source contains null
- Copy is completely independent - changes to original don't affect copy
- More memory efficient than copying to `HashSet` then wrapping with `unmodifiableSet`
- Optimized for immutable use case with minimal overhead

---

## Section 3: Map (Q33–Q52)

### 33. HashMap — Basic Operations
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>();
        map.put("a", 1);
        map.put("b", 2);
        map.put("a", 10); // overwrite existing key
        System.out.println(map.get("a"));
        System.out.println(map.get("c")); // missing key
        System.out.println(map.getOrDefault("c", 0));
    }
}
```
**A:**
```
10
null
0
```
Duplicate key overwrites the value. `get()` returns `null` for missing keys; use `getOrDefault()` to avoid NPE.

**Code Snippet Internal Behavior:**
- `HashMap` uses array of `Node<K,V>[] table` for storage
- `put(key, value)` computes hash: `hash = Objects.hashCode(key) ^ (hash >>> 16)`
- Index calculation: `index = (n-1) & hash` where n is table length (power of 2)
- `put("a", 1)` creates new node at computed index
- `put("a", 10)` finds existing node by hash/key and overwrites value
- `get("a")` computes hash, finds bucket, traverses linked list/red-black tree
- `get("c")` returns null - bucket empty or key not found in chain
- `getOrDefault("c", 0)` returns default value instead of null
- Initial capacity 16, load factor 0.75 - triggers resize at 12 elements

---

### 34. Map.putIfAbsent()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>(Map.of("a", 1));
        map.putIfAbsent("a", 99);  // key exists — NOT overwritten
        map.putIfAbsent("b", 99);  // key absent — inserted
        System.out.println(map.get("a"));
        System.out.println(map.get("b"));
    }
}
```
**A:**
```
1
99
```

**Code Snippet Internal Behavior:**
- `putIfAbsent(key, value)` checks existence before insertion
- Implementation: `V v = get(key); if (v == null) return put(key, value); else return v;`
- For existing key "a": finds value 1, returns it without overwriting
- For missing key "b": value is null, inserts new entry with value 99
- Atomic operation - prevents race conditions in concurrent scenarios
- More efficient than separate `containsKey()` + `put()` calls
- Returns previous value if key existed, null if key was absent
- Useful for initializing maps with default values only when needed
- Internally uses same hash calculation and bucket logic as regular `put()`

---

### 35. Map.computeIfAbsent() — Lazy Initialization
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, List<Integer>> map = new HashMap<>();
        map.computeIfAbsent("evens", k -> new ArrayList<>()).add(2);
        map.computeIfAbsent("evens", k -> new ArrayList<>()).add(4);
        System.out.println(map.get("evens"));
    }
}
```
**A:** `[2, 4]`. `computeIfAbsent()` creates the value only if the key is missing, then returns it. Perfect for grouping/multimaps.

**Code Snippet Internal Behavior:**
- `computeIfAbsent(key, function)` lazily computes value only when key absent
- First call: key "evens" missing, calls function `k -> new ArrayList<>()`
- Creates new ArrayList, stores in map, returns reference to list
- `add(2)` modifies the returned list (now stored in map)
- Second call: key "evens" exists, returns existing list without calling function
- `add(4)` modifies same list instance
- Function called at most once per key - memoization pattern
- Thread-safe only if the computed value is immutable or map is synchronized
- More efficient than `get()` + `putIfAbsent()` pattern
- Common pattern for implementing multimaps: `Map<K, List<V>>`

---

### 36. Map.merge() — Word Count Pattern
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        String[] words = {"apple", "banana", "apple", "cherry", "banana", "apple"};
        Map<String, Integer> count = new HashMap<>();
        for (String w : words) {
            count.merge(w, 1, Integer::sum);
        }
        System.out.println(new TreeMap<>(count));
    }
}
```
**A:** `{apple=3, banana=2, cherry=1}`. `merge(key, 1, sum)` inserts 1 if absent, or applies the function (sum) if present.

**Code Snippet Internal Behavior:**
- `merge(key, value, remappingFunction)` combines existing and new values
- First "apple": key absent, inserts `apple=1`
- Second "apple": key exists with value 1, calls `Integer::sum` with (1, 1) → stores 2
- Third "apple": key exists with value 2, calls `Integer::sum` with (2, 1) → stores 3
- `Integer::sum` is method reference equivalent to `(old, newVal) -> old + newVal`
- Function receives existing value and new value, returns combined result
- If function returns null, entry is removed from map
- Atomic operation - useful for counters, accumulators, and aggregations
- More concise than `map.put(key, map.getOrDefault(key, 0) + 1)` pattern
- Handles all cases: insertion, update, and removal in one method call

---

### 37. LinkedHashMap — Insertion Order
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new LinkedHashMap<>();
        map.put("banana", 2);
        map.put("apple", 1);
        map.put("cherry", 3);
        System.out.println(map.keySet());
    }
}
```
**A:** `[banana, apple, cherry]`. `LinkedHashMap` preserves insertion order, unlike `HashMap`.

**Code Snippet Internal Behavior:**
- `LinkedHashMap` extends `HashMap` and adds doubly-linked list to maintain order
- Each entry has `before` and `after` pointers in addition to hash table entry
- `put()` method appends new entry to end of linked list after hash table insertion
- `keySet()` iteration follows linked list order, not hash table bucket order
- Maintains `head` and `tail` pointers for efficient order maintenance
- Rehashing preserves insertion order by rebuilding linked list
- Slightly more memory overhead: 2 extra references per entry
- All operations still O(1) average case like HashMap
- Useful when iteration order matters but need Map performance
- Order preserved during serialization/deserialization

---

### 38. TreeMap — Sorted by Key
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new TreeMap<>(Map.of("banana", 2, "apple", 1, "cherry", 3));
        System.out.println(map.keySet());
        System.out.println(((TreeMap<String,Integer>)map).firstKey());
        System.out.println(((TreeMap<String,Integer>)map).lastKey());
    }
}
```
**A:**
```
[apple, banana, cherry]
apple
cherry
```

**Code Snippet Internal Behavior:**
- `TreeMap` uses Red-Black Tree data structure internally
- Keys must be `Comparable` or provide `Comparator` at construction
- `firstKey()` returns leftmost node in tree (minimum key)
- `lastKey()` returns rightmost node in tree (maximum key)
- All operations O(log n) due to self-balancing property
- Tree maintains invariants: no red node has red child, all paths have same black count
- Natural ordering uses `key1.compareTo(key2)` for `Comparable` keys
- `TreeMap` implements `NavigableMap` for additional range operations
- No null keys allowed (throws `NullPointerException`)
- More memory overhead than `HashMap` but provides sorted iteration

---

### 39. Map Iteration — entrySet, keySet, values
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new TreeMap<>(Map.of("a", 1, "b", 2, "c", 3));
        for (Map.Entry<String, Integer> e : map.entrySet()) {
            System.out.println(e.getKey() + "=" + e.getValue());
        }
    }
}
```
**A:**
```
a=1
b=2
c=3
```
Always prefer `entrySet()` iteration for maps — more efficient than `keySet()` + `get()`.

**Code Snippet Internal Behavior:**
- `entrySet()` returns `Set<Map.Entry<K,V>>` of all map entries
- Each `Map.Entry` has `getKey()` and `getValue()` methods
- Iterator traverses internal data structure directly (hash table or tree)
- `keySet()` + `get()` pattern requires two hash lookups per entry
- `entrySet()` requires only one traversal - reads key and value together
- For HashMap: iterates array of buckets, then linked lists/red-black trees
- For TreeMap: in-order traversal of Red-Black tree
- Iterator is fail-fast - throws `ConcurrentModificationException` on structural changes
- Memory efficient: no temporary collections created during iteration
- Performance difference significant for large maps

---

### 40. Map.forEach()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new TreeMap<>(Map.of("x", 10, "y", 20, "z", 30));
        map.forEach((k, v) -> System.out.println(k + " -> " + v));
    }
}
```
**A:**
```
x -> 10
y -> 20
z -> 30
```

**Code Snippet Internal Behavior:**
- `forEach(BiConsumer<? super K, ? super V> action)` method added in Java 8
- Internally iterates over `entrySet()` and calls `action.accept(entry.getKey(), entry.getValue())`
- Uses same iteration logic as `entrySet()` - efficient single traversal
- `BiConsumer` functional interface: `void accept(K key, V value)`
- Lambda `(k, v) -> System.out.println(k + " -> " + v)` implements BiConsumer
- More concise than traditional for-each loop with entrySet
- Can be used with method references: `map.forEach(System.out::println)` (prints entries)
- Iterator is fail-fast like other map iteration methods
- Performance equivalent to `entrySet()` iteration
- Supports functional programming style and method chaining

---

### 41. Map.of() — No Null Keys or Values
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = Map.of("a", null); // null value
    }
}
```
**A:** **NullPointerException at runtime.** `Map.of()` does not allow `null` keys or values. Use `HashMap` if you need null support.

**Code Snippet Internal Behavior:**
- `Map.of()` returns `ImmutableCollections.MapN` or optimized variants
- Performs null checks during creation: `Objects.requireNonNull(key)` and `Objects.requireNonNull(value)`
- Throws `NullPointerException` immediately, not during later access
- Returns specialized implementations: `Map0`, `Map1`, `Map2`, `MapN` for different sizes
- All modification methods throw `UnsupportedOperationException`
- Internally stores entries in parallel arrays: `final K[] keys`, `final V[] values`
- Duplicate key detection during creation throws `IllegalArgumentException`
- More memory efficient than `HashMap` for small, immutable maps
- Designed for factory methods and functional programming patterns

---

### 42. HashMap allows null Key
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>();
        map.put(null, 1);
        map.put("a", 2);
        System.out.println(map.get(null));
        System.out.println(map.size());
    }
}
```
**A:**
```
1
2
```
`HashMap` allows one `null` key. `TreeMap` does NOT allow `null` key (throws NullPointerException).

**Code Snippet Internal Behavior:**
- `HashMap` handles null key specially: `hash(key)` returns 0 for null
- Null key stored at index 0 of hash table array
- Only one null key allowed - subsequent `put(null, value)` overwrites existing
- `get(null)` directly checks index 0 without hash calculation
- `TreeMap` calls `key.compareTo()` for ordering - null causes `NullPointerException`
- HashMap null handling: `if (key == null) return 0; else return key.hashCode()`
- TreeMap requires all keys to be `Comparable` and non-null
- Design choice: HashMap optimized for flexibility, TreeMap for strict ordering
- `containsKey(null)` works in HashMap, throws exception in TreeMap

---

### 43. Map.getOrDefault vs computeIfAbsent
**Q: What is the difference?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, List<String>> map = new HashMap<>();

        // getOrDefault does NOT insert into map
        List<String> got = map.getOrDefault("key", new ArrayList<>());
        got.add("hello"); // modifying the returned list does NOT affect map
        System.out.println(map.containsKey("key")); // still false!

        // computeIfAbsent DOES insert
        map.computeIfAbsent("key2", k -> new ArrayList<>()).add("world");
        System.out.println(map.containsKey("key2")); // true
    }
}
```
**A:**
```
false
true
```

**Code Snippet Internal Behavior:**
- `getOrDefault(key, defaultValue)` returns existing value or default, does NOT modify map
- `computeIfAbsent(key, function)` computes and stores value if key missing
- First case: `getOrDefault("key", new ArrayList<>())` creates new ArrayList but doesn't store it
- `got.add("hello")` modifies temporary list, map remains unchanged
- Second case: `computeIfAbsent("key2", k -> new ArrayList<>())` stores new list in map
- `add("world")` modifies stored list, map now contains entry
- Key difference: `getOrDefault` is read-only, `computeIfAbsent` is read-write
- `getOrDefault` useful for safe access without side effects
- `computeIfAbsent` useful for lazy initialization and caching patterns
- Function in `computeIfAbsent` called at most once per key

---

### 44. Frequency Count — Traditional Way
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        int[] nums = {1, 2, 2, 3, 3, 3, 4};
        Map<Integer, Integer> freq = new HashMap<>();
        for (int n : nums) {
            freq.put(n, freq.getOrDefault(n, 0) + 1);
        }
        System.out.println(new TreeMap<>(freq));
    }
}
```
**A:** `{1=1, 2=2, 3=3, 4=1}`

**Code Snippet Internal Behavior:**
- `getOrDefault(key, 0)` safely handles missing keys without null checks
- For each number: checks if key exists, returns current count or 0, then adds 1
- `put(key, count)` stores updated count back in map
- HashMap handles hash collisions automatically via linked lists/red-black trees
- Time complexity: O(n) for all numbers, each operation O(1) average
- More verbose than `merge()` but clearer for beginners
- Alternative: `freq.merge(n, 1, Integer::sum)` (Java 8+)
- HashMap resizes when load factor exceeded (size > capacity * 0.75)
- Uses `Integer.hashCode()` and `Integer.equals()` for key comparison

---

### 45. Map.replaceAll()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>(Map.of("a", 1, "b", 2, "c", 3));
        map.replaceAll((k, v) -> v * 10);
        System.out.println(new TreeMap<>(map));
    }
}
```
**A:** `{a=10, b=20, c=30}`. `replaceAll` applies a function to each value in place.

**Code Snippet Internal Behavior:**
- `replaceAll(BiFunction<? super K, ? super V, ? extends V> function)` iterates over all entries
- For each entry: calls `function.apply(key, value)` and stores result as new value
- Implementation: `for (Map.Entry<K,V> e : entrySet()) e.setValue(function.apply(e.getKey(), e.getValue()))`
- Lambda `(k, v) -> v * 10` receives key and value, returns multiplied value
- `Map.Entry.setValue()` updates value in existing entry (no new entry created)
- Iterator fail-fast - throws `ConcurrentModificationException` if map modified during operation
- More efficient than creating new map with transformed values
- Function can access both key and value for complex transformations
- Atomic per entry, but not atomic for entire map operation

---

### 46. Map.compute()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>(Map.of("a", 5));
        map.compute("a", (k, v) -> v == null ? 1 : v + 1); // a exists → 5+1=6
        map.compute("b", (k, v) -> v == null ? 1 : v + 1); // b absent → 1
        System.out.println(new TreeMap<>(map));
    }
}
```
**A:** `{a=6, b=1}`

**Code Snippet Internal Behavior:**
- `compute(key, remappingFunction)` updates value based on existing value and computation
- First call: key "a" exists with value 5, function returns `5 + 1 = 6`
- Second call: key "b" absent (null), function returns `1` (null case handling)
- Function receives `(key, oldValue)` where oldValue can be null
- If function returns null, entry is removed from map
- Implementation: `V oldValue = get(key); V newValue = remappingFunction.apply(key, oldValue); if (newValue != null) put(key, newValue) else remove(key)`
- More flexible than `merge()` - can access key in computation
- Atomic operation - prevents race conditions in concurrent scenarios
- Useful for conditional updates and complex value transformations
- Can implement increment, decrement, or any custom logic

---

### 47. TreeMap floorKey and ceilingKey
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        TreeMap<Integer, String> map = new TreeMap<>();
        map.put(1, "one"); map.put(3, "three"); map.put(5, "five"); map.put(7, "seven");
        System.out.println(map.floorKey(4));    // largest key <= 4
        System.out.println(map.ceilingKey(4));  // smallest key >= 4
        System.out.println(map.lowerKey(3));    // largest key < 3
        System.out.println(map.higherKey(5));   // smallest key > 5
    }
}
```
**A:**
```
3
5
1
7
```

**Code Snippet Internal Behavior:**
- `floorKey(key)` finds greatest key ≤ given key using tree navigation
- `ceilingKey(key)` finds smallest key ≥ given key
- `lowerKey(key)` finds greatest key < given key
- `higherKey(key)` finds smallest key > given key
- Implementation uses Red-Black tree traversal: starts at root, navigates left/right
- `floorKey(4)`: finds 3 (greatest key ≤ 4)
- `ceilingKey(4)`: finds 5 (smallest key ≥ 4)
- `lowerKey(3)`: finds 1 (greatest key < 3)
- `higherKey(5)`: finds 7 (smallest key > 5)
- All operations O(log n) due to balanced tree property
- Returns null if no matching key found (e.g., `floorKey(0)` would return null)
- Useful for range queries and finding nearest values

---

### 48. LinkedHashMap Access Order — LRU Cache Base
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        // accessOrder=true: iteration in access order (LRU pattern base)
        Map<Integer, String> lru = new LinkedHashMap<>(16, 0.75f, true);
        lru.put(1, "one"); lru.put(2, "two"); lru.put(3, "three");
        lru.get(1); // access key 1 — moves to end
        System.out.println(lru.keySet());
    }
}
```
**A:** `[2, 3, 1]`. With `accessOrder=true`, `LinkedHashMap` orders entries by most-recently accessed. Basis of LRU Cache implementation.

**Code Snippet Internal Behavior:**
- `LinkedHashMap(int initialCapacity, float loadFactor, boolean accessOrder)` constructor
- `accessOrder=true` enables access-order mode instead of insertion-order
- Each access (`get()`) moves entry to end of linked list
- `get(1)` finds entry by hash, then calls `afterNodeAccess()` to move it
- `afterNodeAccess()` removes entry from current position and appends to tail
- Linked list maintains order from least-recently-used (head) to most-recently-used (tail)
- `removeEldestEntry()` can be overridden to implement LRU eviction
- All operations still O(1) average case (hash lookup + list manipulation)
- Perfect foundation for LRU Cache with size-based eviction
- `iteration()` follows access order, not insertion order

---

### 49. EnumMap — Efficient Map for Enum Keys
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    enum Day { MON, TUE, WED, THU, FRI }
    public static void main(String[] args) {
        EnumMap<Day, String> schedule = new EnumMap<>(Day.class);
        schedule.put(Day.MON, "Meeting");
        schedule.put(Day.FRI, "Review");
        System.out.println(schedule);
    }
}
```
**A:** `{MON=Meeting, FRI=Review}`. `EnumMap` maintains natural enum order and is more efficient than `HashMap` for enum keys.

**Code Snippet Internal Behavior:**
- `EnumMap` uses array internally indexed by enum ordinal values
- `EnumMap<Day, String>` creates `Object[] vals = new Object[Day.values().length]`
- `put(Day.MON, "Meeting")` stores at index `Day.MON.ordinal()` (0)
- `put(Day.FRI, "Review")` stores at index `Day.FRI.ordinal()` (4)
- All operations O(1) - direct array access, no hashing
- Memory efficient: only stores values, no hash table overhead
- Iteration follows enum natural order (ordinal order)
- Type-safe at compile time - can't put wrong enum type
- More efficient than `HashMap` for enum keys: no hash calculation, no collisions
- `size()` tracks number of non-null entries in array
- Null values allowed, but null keys not (enum constants are never null)

---

### 50. IdentityHashMap — Reference Equality
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> identity = new IdentityHashMap<>();
        String s1 = new String("key");
        String s2 = new String("key");
        identity.put(s1, 1);
        identity.put(s2, 2); // different object — different entry!
        System.out.println(identity.size());
    }
}
```
**A:** `2`. `IdentityHashMap` uses `==` (reference equality) instead of `equals()`. Two `String` objects with the same content are treated as different keys.

**Code Snippet Internal Behavior:**
- `IdentityHashMap` uses `System.identityHashCode()` instead of `Object.hashCode()`
- Key comparison uses `==` reference equality, not `equals()` method
- `new String("key")` creates two different objects with same content
- `s1 == s2` is false (different objects), so treated as different keys
- Hash calculation: `System.identityHashCode(object)` returns unique hash per object instance
- Useful when object identity matters more than logical equality
- Common use case: tracking object instances, metadata storage
- Table size is power of 2, uses linear probing for collision resolution
- More memory efficient than `WeakHashMap` for identity-based tracking
- `containsKey()` and `get()` use reference comparison throughout

---

### 51. Map.containsKey vs containsValue
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = Map.of("a", 1, "b", 2, "c", 3);
        System.out.println(map.containsKey("b"));    // O(1) for HashMap
        System.out.println(map.containsValue(2));    // O(n) — scans all values
        System.out.println(map.containsValue(99));
    }
}
```
**A:**
```
true
true
false
```

**Code Snippet Internal Behavior:**
- `containsKey(key)` uses hash table lookup: O(1) average for HashMap
- Implementation: `getNode(hash(key), key) != null` - direct hash calculation and bucket search
- `containsValue(value)` requires full scan: O(n) for all map implementations
- Implementation: `for (Map.Entry<K,V> e : entrySet()) if (Objects.equals(value, e.getValue())) return true;`
- `containsKey()` optimized: one hash calculation, at most few comparisons in bucket
- `containsValue()` unoptimized: must check every entry in map
- Performance difference grows with map size - significant for large maps
- `containsValue()` cannot be optimized due to lack of value indexing
- For frequent value lookups, consider maintaining inverse map or using `BiMap`
- Both methods use `Objects.equals()` for null-safe comparison

---

### 52. WeakHashMap — Garbage Collectible Keys
**Q: What is the concept?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) throws Exception {
        Map<Object, String> weak = new WeakHashMap<>();
        Object key = new Object();
        weak.put(key, "value");
        System.out.println(weak.size()); // 1

        key = null; // key no longer strongly reachable
        System.gc();
        Thread.sleep(100);
        System.out.println(weak.size()); // may be 0 — GC may collect the key
    }
}
```
**A:** `1` then likely `0`. `WeakHashMap` holds **weak references** to keys — when a key is garbage collected, the entry is automatically removed. Used for caches.

**Code Snippet Internal Behavior:**
- `WeakHashMap` uses `WeakReference` to wrap keys: `WeakReference<K> ref = new WeakReference<>(key)`
- Keys stored as weak references - eligible for GC when no strong references exist
- `key = null` removes strong reference, making key eligible for garbage collection
- `System.gc()` suggests garbage collection (not guaranteed)
- Entries with GC'd keys automatically removed during map operations
- Internal `ReferenceQueue` tracks garbage-collected keys for cleanup
- `size()` may return stale count until cleanup occurs
- Useful for caches where entries should expire when keys no longer used
- Memory leak prevention: automatically removes entries with dead keys
- Performance overhead due to reference queue processing and cleanup

---

## Section 4: Queue, Deque & PriorityQueue (Q53–Q62)

### 53. Queue — FIFO with offer/poll
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Queue<Integer> queue = new LinkedList<>();
        queue.offer(1); queue.offer(2); queue.offer(3);
        System.out.println(queue.peek());   // view head, don't remove
        System.out.println(queue.poll());   // remove head
        System.out.println(queue.poll());
        System.out.println(queue);
    }
}
```
**A:**
```
1
1
2
[3]
```

---

### 54. Queue — poll vs remove
**Q: What is the difference?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Queue<Integer> queue = new LinkedList<>();
        System.out.println(queue.poll());   // returns null if empty
        try {
            queue.remove(); // throws NoSuchElementException if empty
        } catch (NoSuchElementException e) {
            System.out.println("NoSuchElementException");
        }
    }
}
```
**A:**
```
null
NoSuchElementException
```

---

### 55. PriorityQueue — Min Heap by Default
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        pq.offer(5); pq.offer(1); pq.offer(3); pq.offer(2); pq.offer(4);
        while (!pq.isEmpty()) {
            System.out.print(pq.poll() + " ");
        }
    }
}
```
**A:** `1 2 3 4 5 `. PriorityQueue by default is a min-heap — `poll()` always removes the smallest element.

---

### 56. PriorityQueue — Max Heap with Comparator
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Comparator.reverseOrder());
        maxHeap.offer(5); maxHeap.offer(1); maxHeap.offer(3);
        System.out.println(maxHeap.peek()); // max element
        System.out.println(maxHeap.poll()); // removes max
    }
}
```
**A:**
```
5
5
```

---

### 57. PriorityQueue Iterator Order Not Guaranteed
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(List.of(5, 1, 3, 2, 4));
        // Iterating via for-each does NOT give sorted order!
        for (int x : pq) System.out.print(x + " ");
        System.out.println();
        // Poll gives sorted order
        while (!pq.isEmpty()) System.out.print(pq.poll() + " ");
    }
}
```
**A:** Iteration order is heap-order (not sorted), but poll gives `1 2 3 4 5`. Always use `poll()` for sorted extraction.

---

### 58. ArrayDeque — Stack and Queue
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Deque<Integer> stack = new ArrayDeque<>();
        stack.push(1); stack.push(2); stack.push(3); // push = addFirst
        System.out.println(stack.pop()); // pop = removeFirst → LIFO
        System.out.println(stack.peek());

        Deque<Integer> queue = new ArrayDeque<>();
        queue.offer(1); queue.offer(2); queue.offer(3); // addLast
        System.out.println(queue.poll()); // removeFirst → FIFO
    }
}
```
**A:**
```
3
2
1
```

---

### 59. ArrayDeque vs Stack class
**Q: Why prefer ArrayDeque over Stack?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        // Stack extends Vector (synchronized, slower)
        Stack<Integer> oldStack = new Stack<>();
        oldStack.push(1); oldStack.push(2);

        // ArrayDeque is preferred — faster, not synchronized
        Deque<Integer> newStack = new ArrayDeque<>();
        newStack.push(1); newStack.push(2);

        System.out.println(oldStack.pop());
        System.out.println(newStack.pop());
    }
}
```
**A:**
```
2
2
```
`ArrayDeque` is preferred over `Stack` — it's not synchronized (faster for single-thread), no legacy baggage.

---

### 60. BlockingQueue Concept (java.util.concurrent)
**Q: What does ArrayBlockingQueue provide?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        BlockingQueue<Integer> bq = new ArrayBlockingQueue<>(3);
        bq.put(1); bq.put(2); bq.put(3);
        // bq.put(4); // would BLOCK if full (capacity=3)

        System.out.println(bq.take()); // blocks if empty
        System.out.println(bq.size());
    }
}
```
**A:**
```
1
2
```
`BlockingQueue` is thread-safe. `put()` blocks when full, `take()` blocks when empty. Used in producer-consumer patterns.

---

### 61. Deque as both Stack and Queue
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Deque<String> deque = new ArrayDeque<>();
        deque.addFirst("A");
        deque.addLast("B");
        deque.addFirst("C");
        System.out.println(deque); // front to back
        System.out.println(deque.removeFirst());
        System.out.println(deque.removeLast());
    }
}
```
**A:**
```
[C, A, B]
C
B
```

---

### 62. PriorityQueue with Custom Object
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    record Task(String name, int priority) {}
    public static void main(String[] args) {
        PriorityQueue<Task> pq = new PriorityQueue<>(Comparator.comparingInt(Task::priority));
        pq.offer(new Task("Low", 3));
        pq.offer(new Task("High", 1));
        pq.offer(new Task("Medium", 2));
        while (!pq.isEmpty()) System.out.print(pq.poll().name() + " ");
    }
}
```
**A:** `High Medium Low `. Min-heap by priority field → highest-priority (lowest number) first.

---

## Section 5: Collections Utility & Iterators (Q63–Q75)

### 63. Collections.binarySearch()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>(List.of(1, 2, 3, 4, 5, 6, 7));
        System.out.println(Collections.binarySearch(list, 4));  // found
        System.out.println(Collections.binarySearch(list, 10)); // not found → negative
    }
}
```
**A:**
```
3
-8
```
`binarySearch` returns the index if found, or `-(insertion point) - 1` if not found. **List must be sorted** for correct results.

---

### 64. Collections.disjoint()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> a = List.of(1, 2, 3);
        List<Integer> b = List.of(4, 5, 6);
        List<Integer> c = List.of(3, 7, 8);
        System.out.println(Collections.disjoint(a, b)); // no common elements
        System.out.println(Collections.disjoint(a, c)); // share 3
    }
}
```
**A:**
```
true
false
```

---

### 65. Fail-Fast vs Fail-Safe Iterators
**Q: What is the difference?**
```java
import java.util.*;
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) {
        // Fail-fast: throws ConcurrentModificationException
        List<Integer> failFast = new ArrayList<>(List.of(1, 2, 3));

        // Fail-safe (snapshot copy): no exception, may miss new elements
        List<Integer> failSafe = new CopyOnWriteArrayList<>(List.of(1, 2, 3));

        for (int x : failSafe) {
            failSafe.add(99); // safe — iterates over a snapshot
            System.out.print(x + " ");
            break; // avoid infinite loop
        }
    }
}
```
**A:** `1 `. `CopyOnWriteArrayList` uses a snapshot on iteration — safe to modify during iteration but expensive for writes.

---

### 66. ListIterator — Bidirectional
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>(List.of("a", "b", "c"));
        ListIterator<String> it = list.listIterator(list.size()); // start at end
        while (it.hasPrevious()) {
            System.out.print(it.previous() + " "); // traversal in reverse
        }
    }
}
```
**A:** `c b a `

---

### 67. Collections.min() and max()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> list = List.of(3, 1, 4, 1, 5, 9, 2, 6);
        System.out.println(Collections.min(list));
        System.out.println(Collections.max(list));
    }
}
```
**A:**
```
1
9
```

---

### 68. Collections.singletonList()
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> single = Collections.singletonList("only");
        System.out.println(single.size());
        System.out.println(single.get(0));
        single.add("another"); // attempt to add
    }
}
```
**A:** `1`, `only`, then **UnsupportedOperationException**. `singletonList` is immutable and always contains exactly one element.

---

### 69. Iterable vs Iterator
**Q: What is the relationship?**
```java
import java.util.*;
public class Main {
    static class Range implements Iterable<Integer> {
        int from, to;
        Range(int from, int to) { this.from = from; this.to = to; }

        @Override
        public Iterator<Integer> iterator() {
            return new Iterator<>() {
                int current = from;
                public boolean hasNext() { return current <= to; }
                public Integer next() { return current++; }
            };
        }
    }

    public static void main(String[] args) {
        for (int n : new Range(1, 5)) System.out.print(n + " ");
    }
}
```
**A:** `1 2 3 4 5 `. Implementing `Iterable<T>` allows your class to be used in for-each loops.

---

### 70. Collections.synchronizedList()
**Q: What is the concept?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> syncList = Collections.synchronizedList(new ArrayList<>());
        syncList.add(1); syncList.add(2);

        // Still need manual sync for ITERATION
        synchronized (syncList) {
            for (int x : syncList) System.out.print(x + " ");
        }
    }
}
```
**A:** `1 2 `. Individual operations are thread-safe, but **compound operations** (like iteration) still need external synchronization.

---

### 71. TreeMap descendingMap()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        TreeMap<Integer, String> map = new TreeMap<>(Map.of(1, "one", 2, "two", 3, "three"));
        System.out.println(map.descendingKeySet());
    }
}
```
**A:** `[3, 2, 1]` (reverse order of keys).

---

### 72. Map.Entry manipulation
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>(Map.of("a", 1, "b", 2, "c", 3));
        // Find entry with max value
        Map.Entry<String, Integer> maxEntry = Collections.max(
            map.entrySet(), Map.Entry.comparingByValue());
        System.out.println(maxEntry.getKey() + " = " + maxEntry.getValue());
    }
}
```
**A:** `c = 3` (or whatever entry has the max value).

---

### 73. Sorting Map by Value
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = Map.of("banana", 2, "apple", 5, "cherry", 1);
        map.entrySet().stream()
            .sorted(Map.Entry.comparingByValue())
            .forEach(e -> System.out.println(e.getKey() + "=" + e.getValue()));
    }
}
```
**A:**
```
cherry=1
banana=2
apple=5
```

---

### 74. Collections.emptyList(), emptySet(), emptyMap()
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> empty = Collections.emptyList();
        System.out.println(empty.size());
        System.out.println(empty.isEmpty());
        try { empty.add("x"); } catch (UnsupportedOperationException e) { System.out.println("immutable!"); }
    }
}
```
**A:**
```
0
true
immutable!
```

---

### 75. Stack-based Expression Evaluator (Interview Pattern)
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static boolean isBalanced(String s) {
        Deque<Character> stack = new ArrayDeque<>();
        for (char c : s.toCharArray()) {
            if (c == '(' || c == '[' || c == '{') stack.push(c);
            else if (c == ')' && (stack.isEmpty() || stack.pop() != '(')) return false;
            else if (c == ']' && (stack.isEmpty() || stack.pop() != '[')) return false;
            else if (c == '}' && (stack.isEmpty() || stack.pop() != '{')) return false;
        }
        return stack.isEmpty();
    }
    public static void main(String[] args) {
        System.out.println(isBalanced("({[]})"));
        System.out.println(isBalanced("({[})"));
    }
}
```
**A:**
```
true
false
```
Classic interview question using a stack to validate balanced parentheses. `ArrayDeque` is the preferred stack implementation.
