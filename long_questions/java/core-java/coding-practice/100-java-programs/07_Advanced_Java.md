# Advanced Java Programs (86-100)

## 86. Functional Interface & Lambda
**Principle**: Interface with one method.
**Question**: Use a Lambda expression.
**Code**:
```java
interface Calc { int op(int a, int b); }
public class Main {
    public static void main(String[] args) {
        Calc add = (a, b) -> a + b;
        System.out.println(add.op(5, 3));
    }
}
```

## 87. Stream API - Filter & Map
**Principle**: `stream().filter().map()`.
**Question**: Filter even numbers and square them.
**Code**:
```java
import java.util.*;
import java.util.stream.Collectors;
public class StreamDemo {
    public static void main(String[] args) {
        List<Integer> list = Arrays.asList(1, 2, 3, 4);
        List<Integer> res = list.stream()
                .filter(n -> n % 2 == 0)
                .map(n -> n * n)
                .collect(Collectors.toList());
        System.out.println(res);
    }
}
```

## 88. Count Elements using Stream
**Principle**: `count()`.
**Question**: Count strings starting with 'A'.
**Code**:
```java
long count = Stream.of("Apple", "Banana", "Apricot")
                   .filter(s -> s.startsWith("A"))
                   .count();
```

## 89. Max/Min using Stream
**Principle**: `max(Comparator)`.
**Question**: Find max element.
**Code**:
```java
int max = Arrays.asList(1, 5, 3).stream().max(Integer::compare).get();
```

## 90. Find Duplicate Numbers using Stream
**Principle**: `Collections.frequency` or `Set` add.
**Question**: Find duplicates.
**Code**:
```java
List<Integer> list = Arrays.asList(1, 2, 1, 3);
Set<Integer> set = new HashSet<>();
list.stream().filter(n -> !set.add(n)).forEach(System.out::println);
```

## 91. GroupingBy in Stream
**Principle**: `Collectors.groupingBy`.
**Question**: Group strings by length.
**Code**:
```java
Map<Integer, List<String>> map = Stream.of("A", "BB", "C")
    .collect(Collectors.groupingBy(String::length));
System.out.println(map);
```

## 92. Producer-Consumer Problem
**Principle**: Wait/Notify.
**Question**: Implement Producer-Consumer.
**Code**:
```java
import java.util.*;
class Q {
    int n; boolean set = false;
    synchronized int get() {
        while(!set) try { wait(); } catch(Exception e){}
        System.out.println("Got: " + n);
        set = false; notify();
        return n;
    }
    synchronized void put(int n) {
        while(set) try { wait(); } catch(Exception e){}
        this.n = n;
        System.out.println("Put: " + n);
        set = true; notify();
    }
}
```

## 93. Deadlock Scenario
**Principle**: Two threads holding lock waiting for each other.
**Question**: Create a deadlock.
**Code**:
```java
// Thread 1 locks A then B
// Thread 2 locks B then A
```

## 94. Thread Creation
**Principle**: `extends Thread` or `implements Runnable`.
**Question**: Create and start a thread.
**Code**:
```java
new Thread(() -> System.out.println("Running")).start();
```

## 95. Try-with-resources
**Principle**: Auto-close logical.
**Question**: Read file using try-with-resources.
**Code**:
```java
try (BufferedReader br = new BufferedReader(new FileReader("file.txt"))) {
    System.out.println(br.readLine());
} catch(IOException e) {}
```

## 96. Optional Class
**Principle**: Avoid NullPointerException.
**Question**: Use Optional to handle null.
**Code**:
```java
Optional<String> opt = Optional.ofNullable(null);
System.out.println(opt.orElse("Default"));
```

## 97. Date and Time API (Java 8)
**Principle**: `LocalDate`, `LocalTime`.
**Question**: Get current date.
**Code**:
```java
System.out.println(LocalDate.now());
```

## 98. Check specific Exception
**Principle**: Multiple catch blocks.
**Question**: Handle ArithmeticException.
**Code**:
```java
try { int a = 1/0; } catch (ArithmeticException e) { System.out.println("Zero Div"); }
```

## 99. Execute Service (ThreadPool)
**Principle**: `Executors.newFixedThreadPool`.
**Question**: Create thread pool.
**Code**:
```java
ExecutorService exec = Executors.newFixedThreadPool(2);
exec.submit(() -> System.out.println("Task"));
exec.shutdown();
```

## 100. Reflection API
**Principle**: Inspect class at runtime.
**Question**: Get class name via reflection.
**Code**:
```java
Class<?> clazz = String.class;
System.out.println(clazz.getName());
```
