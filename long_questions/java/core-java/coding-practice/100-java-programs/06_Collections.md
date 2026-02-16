# Collections Programs (76-85)

## 76. Iterate ArrayList
**Principle**: For-loop, Enhanced For, Iterator, Stream.
**Question**: Iterate over an ArrayList in different ways.
**Code**:
```java
import java.util.*;
public class IterateList {
    public static void main(String[] args) {
        List<String> list = Arrays.asList("A", "B");
        // 1. ForEach
        for(String s : list) System.out.print(s);
        // 2. Iterator
        Iterator<String> it = list.iterator();
        while(it.hasNext()) System.out.print(it.next());
    }
}
```

## 77. ArrayList vs LinkedList
**Principle**: Demonstration of types.
**Code**:
```java
List<String> al = new ArrayList<>(); // Fast Access
List<String> ll = new LinkedList<>(); // Fast Insert/Delete
```

## 78. Convert Array to ArrayList
**Principle**: `Arrays.asList()` (fixed size) or `new ArrayList(Arrays.asList())`.
**Question**: Convert Array to List.
**Code**:
```java
String[] arr = {"A", "B"};
List<String> list = new ArrayList<>(Arrays.asList(arr));
```

## 79. Iterate HashMap
**Principle**: `entrySet()`, `keySet()`.
**Question**: Iterate over a HashMap.
**Code**:
```java
import java.util.*;
public class MapIterate {
    public static void main(String[] args) {
        Map<Integer, String> map = new HashMap<>(); 
        map.put(1, "A");
        for(Map.Entry<Integer,String> entry : map.entrySet()) {
            System.out.println(entry.getKey() + "=" + entry.getValue());
        }
    }
}
```

## 80. Sort HashMap by Value
**Principle**: Convert entrySet to List, sort using Comparator.
**Question**: Sort a Map by its values.
**Code**:
```java
import java.util.*;
public class SortMap {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>();
        map.put("A", 3); map.put("B", 1);
        List<Map.Entry<String, Integer>> list = new ArrayList<>(map.entrySet());
        list.sort(Map.Entry.comparingByValue());
        
        list.forEach(System.out::println); // B=1, A=3
    }
}
```

## 81. Remove Element while Iterating
**Principle**: Use `Iterator.remove()` to avoid `ConcurrentModificationException`.
**Question**: Remove elements safely during iteration.
**Code**:
```java
import java.util.*;
public class SafeRemove {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>(Arrays.asList(1, 2, 3));
        Iterator<Integer> it = list.iterator();
        while(it.hasNext()) {
            if(it.next() == 2) it.remove();
        }
        System.out.println(list);
    }
}
```

## 82. Merge Two Lists
**Principle**: `addAll`.
**Question**: Merge two lists.
**Code**:
```java
import java.util.*;
public class MergeLists {
    public static void main(String[] args) {
        List<String> l1 = new ArrayList<>(Arrays.asList("A"));
        List<String> l2 = Arrays.asList("B");
        l1.addAll(l2);
        System.out.println(l1);
    }
}
```

## 83. Find Intersection of Two Lists
**Principle**: `retainAll`.
**Question**: Find common elements in two lists.
**Code**:
```java
import java.util.*;
public class Intersection {
    public static void main(String[] args) {
        List<Integer> l1 = new ArrayList<>(Arrays.asList(1, 2, 3));
        List<Integer> l2 = Arrays.asList(2, 3, 4);
        l1.retainAll(l2);
        System.out.println(l1); // [2, 3]
    }
}
```

## 84. Synchronized ArrayList
**Principle**: `Collections.synchronizedList` or `CopyOnWriteArrayList`.
**Question**: How to satisfy thread-safety for List.
**Code**:
```java
List<String> syncList = Collections.synchronizedList(new ArrayList<>());
```

## 85. Reverse a List
**Principle**: `Collections.reverse()`.
**Question**: Reverse a List.
**Code**:
```java
List<Integer> list = Arrays.asList(1, 2, 3);
Collections.reverse(list);
System.out.println(list);
```
