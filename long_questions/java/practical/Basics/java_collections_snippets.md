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
