# 🟡 Java Data Structures: Level 2 (Intermediate) Practice
Contains runnable code examples for Questions 43-85.

## Question 43: Add, remove, get, set methods.

### Answer
List operations.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class ListOps {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        list.add("A");
        list.add(0, "B"); // [B, A]
        
        list.set(1, "C"); // [B, C]
        
        System.out.println("Get: " + list.get(0));
        
        list.remove("C");
        System.out.println(list);
    }
}
```

---

## Question 44: `addAll()`, `removeAll()`, `retainAll()`.

### Answer
Union, Difference, Intersection.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class BulkOps {
    public static void main(String[] args) {
        List<Integer> l1 = new ArrayList<>(Arrays.asList(1, 2, 3));
        List<Integer> l2 = Arrays.asList(2, 3, 4);
        
        // Intersection
        l1.retainAll(l2); 
        System.out.println("Intersection: " + l1); // [2, 3]
    }
}
```

---

## Question 45: Difference between `ArrayList` and `LinkedList`.

### Answer
Array speed vs Insertion speed.

### Runnable Code
*(Implementation difference shown in previous answers)*

---

## Question 46: `subList()` — what happens if you modify the original list?

### Answer
ConcurrentModificationException.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class SubListDemo {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>(Arrays.asList("A", "B", "C"));
        List<String> sub = list.subList(0, 2); // [A, B]
        
        list.add("D"); // Modifying struct of original
        
        try {
            System.out.println(sub); // Throws exception
        } catch (ConcurrentModificationException e) {
            System.out.println("Sublist broken by structural change!");
        }
    }
}
```

---

## Question 47: `List.of()` vs `Arrays.asList()` vs mutable lists.

### Answer
Immutable vs Fixed vs Mutable.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class ListFactories {
    public static void main(String[] args) {
        List<String> immutable = List.of("A");
        // immutable.add("B"); // UnsupportedOperationException
        
        List<String> fixed = Arrays.asList("A");
        fixed.set(0, "B"); // OK
        // fixed.add("C"); // UnsupportedOperationException
    }
}
```

---

## Question 48: Iterating lists: for, iterator, for-each, streams.

### Answer
Ways to loop.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class IterateList {
    public static void main(String[] args) {
        List<String> list = Arrays.asList("A", "B");
        
        // Iterator (Safe remove)
        Iterator<String> it = new ArrayList<>(list).iterator();
        while(it.hasNext()) {
            if (it.next().equals("A")) it.remove();
        }
    }
}
```

---

## Question 49: `add()`, `remove()`, `contains()` (Set).

### Answer
Set operations.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class SetOps {
    public static void main(String[] args) {
        Set<String> set = new HashSet<>();
        System.out.println(set.add("A")); // true
        System.out.println(set.add("A")); // false
    }
}
```

---

## Question 50: Difference between `HashSet`, `TreeSet`, `LinkedHashSet`.

### Answer
Order: None, Sorted, Insertion.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class SetTypes {
    public static void main(String[] args) {
        Set<Integer> hash = new HashSet<>();
        Set<Integer> linked = new LinkedHashSet<>();
        Set<Integer> tree = new TreeSet<>();
        
        for (int i : Arrays.asList(5, 1, 3)) {
            hash.add(i); linked.add(i); tree.add(i);
        }
        
        System.out.println("Hash (Random): " + hash);
        System.out.println("Linked (Order): " + linked); // [5, 1, 3]
        System.out.println("Tree (Sorted): " + tree);   // [1, 3, 5]
    }
}
```

---

## Question 51: How does `TreeSet` maintain order? What interface is required?

### Answer
`Comparable`.

### Runnable Code
```java
package datastructures;

import java.util.*;

class User implements Comparable<User> {
    int id;
    User(int id) { this.id = id; }
    public int compareTo(User o) { return this.id - o.id; }
    public String toString() { return ""+id; }
}

public class TreeSetSort {
    public static void main(String[] args) {
        Set<User> set = new TreeSet<>();
        set.add(new User(10));
        set.add(new User(5));
        System.out.println(set); // [5, 10]
    }
}
```

---

## Question 52: How to convert a List to Set and vice versa?

### Answer
Constructor.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class ConvertColl {
    public static void main(String[] args) {
        List<Integer> list = List.of(1, 1, 2);
        Set<Integer> set = new HashSet<>(list); // Removes dup
        List<Integer> back = new ArrayList<>(set);
        System.out.println(back); // [1, 2]
    }
}
```

---

## Question 53: `put()`, `get()`, `remove()`, `containsKey()`, `containsValue()` (Map).

### Answer
Map Basic Ops.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class MapOps {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>();
        map.put("A", 1);
        System.out.println(map.containsKey("A")); // true
    }
}
```

---

## Question 54: Iterating over Map entries: `entrySet()`, `keySet()`, `values()`.

### Answer
Looping maps.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class MapIterate {
    public static void main(String[] args) {
        Map<String, Integer> map = Map.of("A", 1, "B", 2);
        
        for (Map.Entry<String, Integer> e : map.entrySet()) {
            System.out.println(e.getKey() + "=" + e.getValue());
        }
    }
}
```

---

## Question 55: Difference between `HashMap` and `TreeMap`?

### Answer
Sort order.

### Runnable Code
*(See SetTypes code logic)*

---

## Question 56: Difference between `computeIfAbsent()` and `putIfAbsent()`.

### Answer
Lazy vs Eager.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class MapCompute {
    public static void main(String[] args) {
        Map<String, List<String>> map = new HashMap<>();
        
        // Smart initialization of list
        map.computeIfAbsent("Key", k -> new ArrayList<>()).add("Value");
    }
}
```

---

## Question 57: What happens when two keys have the same hashcode?

### Answer
Collision (Chaining).

### Runnable Code
*(Internal workings demonstration not applicable via simple API)*

---

## Question 58: How to maintain insertion order in a Map?

### Answer
`LinkedHashMap`.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class OrderedMap {
    public static void main(String[] args) {
        Map<String, String> map = new LinkedHashMap<>();
        map.put("Z", "1");
        map.put("A", "2");
        System.out.println(map.keySet()); // [Z, A]
    }
}
```

---

## Question 59: Difference between `Queue` and `Deque`.

### Answer
One-ended vs Two-ended.

### Runnable Code
*(See next questions)*

---

## Question 60: Methods in `Queue`: `offer()`, `poll()`, `peek()`.

### Answer
Queue Ops.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class QueueDemo {
    public static void main(String[] args) {
        Queue<String> q = new LinkedList<>();
        q.offer("First");
        System.out.println(q.peek()); // First
        System.out.println(q.poll()); // First
    }
}
```

---

## Question 61: Methods in `Deque`: `addFirst()`, `addLast()`, `removeFirst()`, `removeLast()`.

### Answer
Stack/Queue hybrid.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class DequeDemo {
    public static void main(String[] args) {
        Deque<String> d = new ArrayDeque<>();
        d.addFirst("A");
        d.addLast("B");
        System.out.println(d); // [A, B]
    }
}
```

---

## Question 62: Difference between `Stack` class and `Deque` for stack operations.

### Answer
Use Deque (Stack is legacy).

### Runnable Code
```java
package datastructures;

import java.util.*;

public class StackDemo {
    public static void main(String[] args) {
        // Correct Stack
        Deque<Integer> stack = new ArrayDeque<>();
        stack.push(1);
        stack.push(2);
        System.out.println(stack.pop()); // 2
    }
}
```

---

## Question 63: How to implement a stack/queue using arrays or linked lists?

### Answer
Custom Implementation.

### Runnable Code
```java
package datastructures;

class MyStack {
    int[] arr = new int[10];
    int top = -1;
    void push(int x) { arr[++top] = x; }
    int pop() { return arr[top--]; }
}
```

---

## Question 64: Difference between `PriorityQueue` and `Queue`.

### Answer
Sorted Queue.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class PQDemo {
    public static void main(String[] args) {
        Queue<Integer> pq = new PriorityQueue<>();
        pq.offer(5);
        pq.offer(1);
        System.out.println(pq.poll()); // 1 (Smallest)
    }
}
```

---

## Question 65: Difference between singly and doubly linked list.

### Answer
Prev pointer.

### Runnable Code
*(Conceptual)*

---

## Question 66: Common methods: `addFirst()`, `addLast()`, `removeFirst()`, `removeLast()`.

### Answer
LinkedList Ops.

### Runnable Code
*(See Q61)*

---

## Question 67: Traversal and searching operations.

### Answer
ListIterator.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class ListIterDemo {
    public static void main(String[] args) {
        List<String> l = new LinkedList<>(List.of("A", "B"));
        ListIterator<String> it = l.listIterator(l.size());
        while(it.hasPrevious()) System.out.print(it.previous()); // BA
    }
}
```

---

## Question 68: How does `HashMap` compute the bucket for a key?

### Answer
`(n-1) & hash`.

### Runnable Code
*(Internal)*

---

## Question 69: Difference between `hashCode()` and `equals()`.

### Answer
Find Bucket vs Match Content.

### Runnable Code
*(See Q44)*

---

## Question 70: Why is `hashCode()` required for hash-based collections?

### Answer
Performance.

### Runnable Code
*(Conceptual)*

---

## Question 71: Tree traversal (`preorder`, `inorder`, `postorder`) — implement in Java.

### Answer
Recursion.

### Runnable Code
```java
package datastructures;

class Node { 
    int val; Node left, right;
    Node(int v) { val = v; }
}

public class Traversals {
    static void inOrder(Node n) {
        if (n == null) return;
        inOrder(n.left);
        System.out.print(n.val + " ");
        inOrder(n.right);
    }
    
    public static void main(String[] args) {
        Node root = new Node(1);
        root.left = new Node(2);
        root.right = new Node(3);
        inOrder(root); // 2 1 3
    }
}
```

---

## Question 72: Binary search tree insertion & search.

### Answer
Left < Root < Right.

### Runnable Code
```java
package datastructures;

public class BSTDemo {
    // See Trie/Graph for complex structures.
    // Logic: if val < cur.val go left else right
}
```

---

## Question 73: Difference between BST and Heap (method/operation-level).

### Answer
Order vs Priority.

### Runnable Code
*(Conceptual)*

---

## Question 74: `map()`, `flatMap()`, `filter()`.

### Answer
Stream Ops.

### Runnable Code
```java
package datastructures;

import java.util.*;
import java.util.stream.*;

public class StreamOps {
    public static void main(String[] args) {
        List<List<Integer>> nested = List.of(List.of(1), List.of(2, 3));
        
        List<Integer> flat = nested.stream()
            .flatMap(Collection::stream)
            .map(x -> x * 2)
            .collect(Collectors.toList());
            
        System.out.println(flat); // [2, 4, 6]
    }
}
```

---

## Question 75: `collect()`, `toList()`, `toSet()`.

### Answer
Terminals.

### Runnable Code
*(See Q74)*

---

## Question 76: `reduce()` — sum, max, concatenation.

### Answer
Reduction.

### Runnable Code
```java
package datastructures;

import java.util.stream.Stream;

public class ReduceDemo {
    public static void main(String[] args) {
        int sum = Stream.of(1, 2, 3).reduce(0, Integer::sum);
        System.out.println(sum);
    }
}
```

---

## Question 77: `forEach()` — usage.

### Answer
Action.

### Runnable Code
```java
package datastructures;

import java.util.stream.Stream;

public class ForEachDemo {
    public static void main(String[] args) {
        Stream.of(1, 2).forEach(System.out::print);
    }
}
```

---

## Question 78: `sorted()`, `distinct()`.

### Answer
Stateful ops.

### Runnable Code
```java
package datastructures;

import java.util.stream.Stream;

public class StatefulOps {
    public static void main(String[] args) {
        Stream.of(3, 1, 2, 1).distinct().sorted().forEach(System.out::print); // 123
    }
}
```

---

## Question 79: Parallel stream vs sequential stream.

### Answer
Multi-thread.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class ParallelDemo {
    public static void main(String[] args) {
        List.of(1, 2, 3).parallelStream().forEach(n -> 
            System.out.println(Thread.currentThread().getName() + " " + n));
    }
}
```

---

## Question 80: Stream from arrays: `Arrays.stream(arr)`.

### Answer
Primitive Streams.

### Runnable Code
*(See Q16)*

---

## Question 81: Sliding window over array/string.

### Answer
Subarray pattern.

### Runnable Code
```java
package datastructures;

public class SlidingWindow {
    // Max Sum subarray of size k
    public static void main(String[] args) {
        int[] arr = {1, 4, 2, 10, 2};
        int k = 3;
        int sum = 0, max = 0;
        
        for (int i=0; i<k; i++) sum += arr[i];
        max = sum;
        
        for (int i=k; i<arr.length; i++) {
            sum += arr[i] - arr[i-k];
            max = Math.max(max, sum);
        }
        System.out.println("Max Window: " + max); // 4+2+10=16
    }
}
```

---

## Question 82: Two-pointer technique.

### Answer
Start/End pointers.

### Runnable Code
*(See Q5, Q26)*

---

## Question 83: Hashing for arrays/strings.

### Answer
Freq Map.

### Runnable Code
*(See Q30)*

---

## Question 84: Prefix sum arrays.

### Answer
Range sum.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class PrefixSum {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3};
        int[] pre = new int[arr.length];
        
        pre[0] = arr[0];
        for(int i=1; i<arr.length; i++) pre[i] = pre[i-1] + arr[i];
        
        System.out.println("Prefix: " + Arrays.toString(pre)); // 1, 3, 6
    }
}
```

---

## Question 85: Frequency maps for counting characters or numbers.

### Answer
Counting.

### Runnable Code
*(See Q30)*
