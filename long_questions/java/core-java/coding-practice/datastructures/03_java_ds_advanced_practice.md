# 🔴 Java Data Structures: Level 3 (Advanced) Practice
Contains runnable code examples for Questions 86-115.

## Question 86: `Arrays.mismatch()` (Java 9+).

### Answer
Finds index of difference.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class MismatchDemo {
    public static void main(String[] args) {
        int[] a = {1, 2, 3};
        int[] b = {1, 2, 4};
        System.out.println(Arrays.mismatch(a, b)); // 2
    }
}
```

---

## Question 87: `Arrays.parallelSort()` (Java 8+).

### Answer
ForkJoin Sort.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class ParallelSort {
    public static void main(String[] args) {
        int[] arr = {5, 2, 8, 1};
        Arrays.parallelSort(arr);
        System.out.println(Arrays.toString(arr));
    }
}
```

---

## Question 88: `String.repeat(int count)` (Java 11+).

### Answer
Repeat string.

### Runnable Code
```java
package datastructures;

public class StringRepeat {
    public static void main(String[] args) {
        System.out.println("Na".repeat(3)); // NaNaNa
    }
}
```

---

## Question 89: `String.isBlank()` (Java 11+).

### Answer
Checks whitespace.

### Runnable Code
```java
package datastructures;

public class CheckBlank {
    public static void main(String[] args) {
        System.out.println("  ".isBlank()); // true
    }
}
```

---

## Question 90: `Map.ofEntries()` (Java 9+).

### Answer
Large Immutable maps.

### Runnable Code
```java
package datastructures;

import java.util.Map;

public class MapEntries {
    public static void main(String[] args) {
        Map<String, Integer> map = Map.ofEntries(
            Map.entry("Key1", 1),
            Map.entry("Key2", 2)
        );
        System.out.println(map);
    }
}
```

---

## Question 91: `Set.of()` (Java 9+).

### Answer
Immutable Set.

### Runnable Code
```java
package datastructures;

import java.util.Set;

public class SetOf {
    public static void main(String[] args) {
        System.out.println(Set.of("A", "B"));
    }
}
```

---

## Question 92: `List.of()` (Java 9+).

### Answer
Immutable List.

### Runnable Code
```java
package datastructures;

public class ListOf {
    public static void main(String[] args) {
        // List.of("A").add("B"); // Crash
    }
}
```

---

## Question 93: `Collections.emptyList()`, `Collections.emptyMap()`.

### Answer
Empty Constants.

### Runnable Code
```java
package datastructures;

import java.util.Collections;
import java.util.List;

public class EmptyList {
    public static void main(String[] args) {
        List<String> list = Collections.emptyList();
        System.out.println(list.size());
    }
}
```

---

## Question 94: `Collectors.partitioningBy()` and `Collectors.groupingBy()`.

### Answer
Stream Groups.

### Runnable Code
```java
package datastructures;

import java.util.*;
import java.util.stream.Collectors;

public class Grouping {
    public static void main(String[] args) {
        List<String> words = List.of("a", "bb", "ccc", "d");
        
        Map<Integer, List<String>> byLength = words.stream()
            .collect(Collectors.groupingBy(String::length));
            
        System.out.println(byLength); // {1=[a, d], 2=[bb], 3=[ccc]}
    }
}
```

---

## Question 95: `Collectors.counting()`.

### Answer
Count groups.

### Runnable Code
```java
package datastructures;

import java.util.*;
import java.util.stream.Collectors;

public class CountGroups {
    public static void main(String[] args) {
        List<String> words = List.of("a", "b", "a");
        Map<String, Long> count = words.stream()
            .collect(Collectors.groupingBy(w -> w, Collectors.counting()));
            
        System.out.println(count); // {a=2, b=1}
    }
}
```

---

## Question 96: `Collectors.joining()`.

### Answer
Join Strings.

### Runnable Code
```java
package datastructures;

import java.util.List;
import java.util.stream.Collectors;

public class JoinStrings {
    public static void main(String[] args) {
        String res = List.of("A", "B").stream().collect(Collectors.joining(","));
        System.out.println(res);
    }
}
```

---

## Question 97: Stream `peek()` — when and why to use carefully.

### Answer
Debug.

### Runnable Code
```java
package datastructures;

import java.util.stream.Stream;

public class PeekDemo {
    public static void main(String[] args) {
        Stream.of(1, 2)
            .peek(n -> System.out.println("Processing: " + n))
            .forEach(System.out::println);
    }
}
```

---

## Question 98: Sliding window for arrays with variable window size.

### Answer
Condition based.

### Runnable Code
```java
package datastructures;

public class VariableWindow {
    // Smallest Subarray Length with Sum >= S
    public static void main(String[] args) {
        int[] arr = {2, 1, 5, 2, 3, 2};
        int S = 7;
        
        int minLen = Integer.MAX_VALUE;
        int windowSum = 0;
        int left = 0;
        
        for (int right = 0; right < arr.length; right++) {
            windowSum += arr[right];
            
            while (windowSum >= S) {
                minLen = Math.min(minLen, right - left + 1);
                windowSum -= arr[left];
                left++;
            }
        }
        System.out.println("Min Len: " + (minLen == Integer.MAX_VALUE ? 0 : minLen)); // 2 (5, 2)
    }
}
```

---

## Question 99: Rolling hash / Rabin-Karp style substring search.

### Answer
Hash optimization.

### Runnable Code
*(Advanced Algo)*

---

## Question 100: How to implement an LRU Cache in Java?

### Answer
LinkedHashMap.

### Runnable Code
```java
package datastructures;

import java.util.LinkedHashMap;
import java.util.Map;

class LRUCache<K, V> extends LinkedHashMap<K, V> {
    private final int capacity;
    LRUCache(int cap) { 
        super(cap, 0.75f, true); 
        this.capacity = cap;
    }
    @Override
    protected boolean removeEldestEntry(Map.Entry<K, V> eldest) {
        return size() > capacity;
    }
}

public class LRUDemo {
    public static void main(String[] args) {
        Map<Integer, String> lru = new LRUCache<>(2);
        lru.put(1, "A");
        lru.put(2, "B");
        lru.get(1); // Access 1
        lru.put(3, "C"); // Evicts 2 (Eldest unused)
        
        System.out.println(lru.keySet()); // [1, 3]
    }
}
```

---

## Question 101: How to implement a Trie (Prefix Tree) in Java?

### Answer
Tree of Maps.

### Runnable Code
```java
package datastructures;

class TrieNode {
    TrieNode[] children = new TrieNode[26];
    boolean isEnd;
}

class Trie {
    TrieNode root = new TrieNode();
    
    void insert(String word) {
        TrieNode curr = root;
        for (char c : word.toCharArray()) {
            int idx = c - 'a';
            if (curr.children[idx] == null) curr.children[idx] = new TrieNode();
            curr = curr.children[idx];
        }
        curr.isEnd = true;
    }
    
    boolean search(String word) {
        TrieNode curr = root;
        for (char c : word.toCharArray()) {
            if (curr.children[c - 'a'] == null) return false;
            curr = curr.children[c - 'a'];
        }
        return curr.isEnd;
    }
}

public class TrieDemo {
    public static void main(String[] args) {
        Trie t = new Trie();
        t.insert("apple");
        System.out.println(t.search("apple")); // true
        System.out.println(t.search("app"));   // false
    }
}
```

---

## Question 102: How to implement a Graph using Adjacency List?

### Answer
Map<Int, List>.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class GraphDemo {
    public static void main(String[] args) {
        Map<Integer, List<Integer>> adj = new HashMap<>();
        adj.computeIfAbsent(1, k -> new ArrayList<>()).add(2);
        adj.computeIfAbsent(2, k -> new ArrayList<>()).add(1);
        System.out.println(adj);
    }
}
```

---

## Question 103: BFS vs DFS implementation in Java?

### Answer
Queue vs Stack.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class GraphTraversal {
    public static void main(String[] args) {
        // Graph: 1 -> 2, 1 -> 3
        Map<Integer, List<Integer>> adj = Map.of(1, List.of(2, 3));
        
        // BFS
        Queue<Integer> q = new LinkedList<>();
        q.add(1);
        while(!q.isEmpty()) {
            System.out.print(q.poll() + " ");
        }
    }
}
```

---

## Question 104: How to implement a Min/Max Heap?

### Answer
Array logic.

### Runnable Code
*(See PriorityQueue Q64)*

---

## Question 105: Disjoint Set Union (DSU) / Union-Find.

### Answer
Union/Find path compression.

### Runnable Code
```java
package datastructures;

class UnionFind {
    int[] parent;
    UnionFind(int n) {
        parent = new int[n];
        for(int i=0; i<n; i++) parent[i] = i;
    }
    int find(int x) {
        if(parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }
    void union(int x, int y) {
        parent[find(x)] = find(y);
    }
}
```

---

## Question 106: ConcurrentHashMap vs Hashtable vs SynchronizedMap.

### Answer
Bucket locking vs Map locking.

### Runnable Code
```java
package datastructures;

import java.util.concurrent.ConcurrentHashMap;

public class ConcurrentMapDemo {
    public static void main(String[] args) {
        var map = new ConcurrentHashMap<String, Integer>();
        map.put("A", 1);
        // Safe to iterate while modifying
        for(String k : map.keySet()) {
            map.put("B", 2); 
        }
    }
}
```

---

## Question 107: BlockingQueue — Producer-Consumer Pattern.

### Answer
Thread safe queue.

### Runnable Code
```java
package datastructures;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

public class ProducerConsumer {
    public static void main(String[] args) throws InterruptedException {
        BlockingQueue<Integer> q = new ArrayBlockingQueue<>(1);
        q.put(1); // Waits if full
        System.out.println(q.take()); // Waits if empty
    }
}
```

---

## Question 108: CopyOnWriteArrayList.

### Answer
Safe iteration.

### Runnable Code
```java
package datastructures;

import java.util.concurrent.CopyOnWriteArrayList;

public class CWALDemo {
    public static void main(String[] args) {
        var list = new CopyOnWriteArrayList<>(java.util.List.of("A"));
        for (String s : list) {
            list.add("B"); // No ConcurrentModificationException
        }
        System.out.println(list); // [A, B]
    }
}
```

---

## Question 109: IdentityHashMap vs WeakHashMap.

### Answer
Ref Equal, GC Keys.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class SpecialMaps {
    public static void main(String[] args) {
        Map<String, String> identity = new IdentityHashMap<>();
        identity.put(new String("A"), "1");
        identity.put(new String("A"), "2");
        System.out.println("Identity Size: " + identity.size()); // 2
    }
}
```

---

## Question 110: EnumSet and EnumMap.

### Answer
Optimized for Enums.

### Runnable Code
```java
package datastructures;

import java.util.*;

enum Day { MON, TUE }

public class EnumCollections {
    public static void main(String[] args) {
        Set<Day> days = EnumSet.of(Day.MON);
        System.out.println(days);
    }
}
```

---

## Question 111: BitSet in Java.

### Answer
Bit array.

### Runnable Code
```java
package datastructures;

import java.util.BitSet;

public class BitSetDemo {
    public static void main(String[] args) {
        BitSet bs = new BitSet();
        bs.set(0);
        bs.set(2);
        System.out.println(bs); // {0, 2}
    }
}
```

---

## Question 112: How to detect a cycle in a LinkedList?

### Answer
Floyd's Cycle.

### Runnable Code
```java
package datastructures;

class ListNode { int val; ListNode next; ListNode(int x) { val = x; } }

public class DetectCycle {
    public static boolean hasCycle(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast) return true;
        }
        return false;
    }
}
```

---

## Question 113: How to find the middle of a LinkedList?

### Answer
Slow/Fast.

### Runnable Code
```java
package datastructures;

public class MiddleNode {
    // slow moves 1, fast moves 2. When fast end, slow is middle.
}
```

---

## Question 114: Flatten a nested List or Iterator.

### Answer
Stack of iterators.

### Runnable Code
*(Algorithm)*

---

## Question 115: Monotonic Stack (Next Greater Element).

### Answer
Stack decreasing.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class NextGreater {
    public static void main(String[] args) {
        int[] arr = {2, 1, 5};
        Stack<Integer> s = new Stack<>();
        int[] res = new int[arr.length];
        
        for(int i = arr.length-1; i >= 0; i--) {
            while(!s.isEmpty() && s.peek() <= arr[i]) s.pop();
            res[i] = s.isEmpty() ? -1 : s.peek();
            s.push(arr[i]);
        }
        System.out.println(Arrays.toString(res)); // [5, 5, -1]
    }
}
```
