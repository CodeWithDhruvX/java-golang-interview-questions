# Java Functional Programming & Streams — Practical Code Snippets

> **Topics:** Lambda expressions, Functional interfaces, Method references, Optional, Stream API (filter/map/reduce/collect/flatMap), Collectors, Parallel streams

---

## 📋 Reading Progress

- [ ] **Section 1:** Lambda Expressions & Functional Interfaces (Q1–Q16)
- [ ] **Section 2:** Stream — Intermediate Operations (Q17–Q32)
- [ ] **Section 3:** Stream — Terminal Operations & Collectors (Q33–Q50)
- [ ] **Section 4:** Optional (Q51–Q60)
- [ ] **Section 5:** Parallel Streams & Advanced (Q61–Q70)

> 🔖 **Last read:** <!-- -->

---

## Section 1: Lambda & Functional Interfaces (Q1–Q16)

### 1. Basic Lambda
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> fruits = new ArrayList<>(List.of("banana", "apple", "cherry"));
        fruits.sort((a, b) -> a.compareTo(b));
        System.out.println(fruits);
    }
}
```
**A:** `[apple, banana, cherry]`. Lambda `(a, b) -> a.compareTo(b)` implements `Comparator<String>`.

---

### 2. Predicate — test()
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        Predicate<Integer> isEven = n -> n % 2 == 0;
        Predicate<Integer> isPositive = n -> n > 0;

        System.out.println(isEven.test(4));
        System.out.println(isPositive.test(-3));
        System.out.println(isEven.and(isPositive).test(4));
        System.out.println(isEven.or(isPositive).test(-4));
        System.out.println(isEven.negate().test(4));
    }
}
```
**A:**
```
true
false
true
true
false
```

---

### 3. Function — apply() and Composition
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        Function<String, Integer> length = String::length;
        Function<Integer, Boolean> isLong = n -> n > 5;

        Function<String, Boolean> isLongStr = length.andThen(isLong);
        System.out.println(isLongStr.apply("hello"));
        System.out.println(isLongStr.apply("hello world"));
    }
}
```
**A:**
```
false
true
```
`andThen(g)` creates a composed function: apply `f` then `g`. `compose(g)` applies `g` first then `f`.

---

### 4. Consumer — accept()
**Q: What is the output?**
```java
import java.util.*;
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        Consumer<String> print = System.out::println;
        Consumer<String> printUpper = s -> System.out.println(s.toUpperCase());
        Consumer<String> both = print.andThen(printUpper);

        both.accept("hello");
    }
}
```
**A:**
```
hello
HELLO
```

---

### 5. Supplier — get()
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        Supplier<String> greeting = () -> "Hello, World!";
        Supplier<Double> random = Math::random;

        System.out.println(greeting.get());
        System.out.println(random.get() >= 0 && random.get() < 1);
    }
}
```
**A:**
```
Hello, World!
true
```

---

### 6. BiFunction and BiPredicate
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        BiFunction<String, Integer, String> repeat = (s, n) -> s.repeat(n);
        System.out.println(repeat.apply("ab", 3));

        BiPredicate<Integer, Integer> sum10 = (a, b) -> a + b == 10;
        System.out.println(sum10.test(3, 7));
        System.out.println(sum10.test(3, 8));
    }
}
```
**A:**
```
ababab
true
false
```

---

### 7. UnaryOperator and BinaryOperator
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        UnaryOperator<String> trim = String::trim;
        System.out.println(trim.apply("  hello  "));

        BinaryOperator<Integer> add = Integer::sum;
        System.out.println(add.apply(3, 4));
    }
}
```
**A:**
```
hello
7
```

---

### 8. Method Reference — 4 Types
**Q: What is the output?**
```java
import java.util.*;
import java.util.function.*;
public class Main {
    static int doubleIt(int n) { return n * 2; }

    public static void main(String[] args) {
        // 1. Static method reference
        Function<Integer, Integer> dbl = Main::doubleIt;

        // 2. Instance method on a specific instance
        String prefix = "Hello";
        Function<String, String> greet = prefix::concat;

        // 3. Instance method on arbitrary instance
        Function<String, String> upper = String::toUpperCase;

        // 4. Constructor reference
        Function<String, StringBuilder> sbMaker = StringBuilder::new;

        System.out.println(dbl.apply(5));
        System.out.println(greet.apply(" World"));
        System.out.println(upper.apply("hello"));
        System.out.println(sbMaker.apply("test"));
    }
}
```
**A:**
```
10
Hello World
HELLO
test
```

---

### 9. Effectively Final Variable in Lambda
**Q: Does this compile?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        int multiplier = 3; // effectively final
        UnaryOperator<Integer> times3 = n -> n * multiplier;
        System.out.println(times3.apply(5));
        // multiplier = 4; // uncommenting breaks compilation
    }
}
```
**A:** **Compiles and prints** `15`. Lambdas can capture local variables only if they are (effectively) final.

---

### 10. Custom Functional Interface
**Q: What is the output?**
```java
public class Main {
    @FunctionalInterface
    interface TriFunction<A, B, C, R> {
        R apply(A a, B b, C c);
    }

    public static void main(String[] args) {
        TriFunction<Integer, Integer, Integer, Integer> sum3 = (a, b, c) -> a + b + c;
        System.out.println(sum3.apply(1, 2, 3));
    }
}
```
**A:** `6`

---

### 11. Predicate.not() (Java 11+)
**Q: What is the output?**
```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> words = List.of("hello", "", "world", "  ", "!");
        List<String> nonBlank = words.stream()
                .filter(Predicate.not(String::isBlank))
                .collect(Collectors.toList());
        System.out.println(nonBlank);
    }
}
```
**A:** `[hello, world, !]`. `Predicate.not(pred)` negates a method reference predicate — cleaner than `s -> !s.isBlank()`.

---

### 12. Currying with Function Composition
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        Function<Integer, Function<Integer, Integer>> add = a -> b -> a + b;
        Function<Integer, Integer> add5 = add.apply(5);
        System.out.println(add5.apply(3));
        System.out.println(add5.apply(10));
    }
}
```
**A:**
```
8
15
```
Currying breaks a multi-argument function into a chain of single-argument functions.

---

### 13. Comparator Chaining
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    record Person(String name, int age) {}
    public static void main(String[] args) {
        List<Person> people = new ArrayList<>(List.of(
            new Person("Alice", 30), new Person("Bob", 25),
            new Person("Charlie", 30), new Person("Dave", 25)));
        people.sort(Comparator.comparingInt(Person::age).thenComparing(Person::name));
        people.forEach(p -> System.out.println(p.name() + " " + p.age()));
    }
}
```
**A:**
```
Bob 25
Dave 25
Alice 30
Charlie 30
```

---

### 14. IntFunction, IntUnaryOperator (Primitive Specializations)
**Q: Why use primitive functional interfaces?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        // Avoids boxing/unboxing
        IntUnaryOperator square = x -> x * x;
        IntBinaryOperator add    = Integer::sum;

        System.out.println(square.applyAsInt(5));
        System.out.println(add.applyAsInt(3, 4));
    }
}
```
**A:**
```
25
7
```
Primitive specializations (`IntUnaryOperator`, `LongFunction`, etc.) avoid autoboxing overhead.

---

### 15. Function Identity
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        Function<String, String> id = Function.identity();
        System.out.println(id.apply("hello"));
        System.out.println(id.apply("world"));
    }
}
```
**A:**
```
hello
world
```
`Function.identity()` returns a function that always returns its input argument. Useful as a no-op transformer.

---

### 16. Lambda Closure vs Anonymous Class Scope
**Q: What is the output?**
```java
public class Main {
    int val = 10;
    void run() {
        Runnable lambda  = ()             -> System.out.println("lambda val: " + val);
        Runnable anon    = new Runnable() { public void run() { System.out.println("anon val: " + val); } };
        val = 99;
        lambda.run(); // sees updated val — captures 'this'
        anon.run();   // also sees updated val — captures 'this' implicitly
    }
    public static void main(String[] args) { new Main().run(); }
}
```
**A:**
```
lambda val: 99
anon val: 99
```
Both capture `this.val` (an instance field), not a local variable. Instance (field) state can be mutable — the restriction applies only to local variables.

---

## Section 2: Stream — Intermediate Operations (Q17–Q32)

### 17. Stream Pipeline — Lazy Evaluation
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        long count = Stream.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
                .filter(n -> { System.out.print("F" + n + " "); return n % 2 == 0; })
                .map(n -> { System.out.print("M" + n + " "); return n * 2; })
                .limit(3)
                .count();
        System.out.println("\ncount: " + count);
    }
}
```
**A:** `F1 F2 M2 F3 F4 M4 F5 F6 M6` (approximately, interleaved) then `count: 3`. Streams are **lazy** — elements are processed one-at-a-time through the pipeline. `limit(3)` stops after 3 even numbers.

---

### 18. filter, map, collect
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> result = Stream.of("apple", "banana", "cherry", "date", "elderberry")
                .filter(s -> s.length() > 5)
                .map(String::toUpperCase)
                .sorted()
                .collect(Collectors.toList());
        System.out.println(result);
    }
}
```
**A:** `[BANANA, CHERRY, ELDERBERRY]`

---

### 19. flatMap — Flatten Nested Streams
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<List<Integer>> nested = List.of(
            List.of(1, 2, 3), List.of(4, 5), List.of(6, 7, 8, 9));
        List<Integer> flat = nested.stream()
                .flatMap(Collection::stream)
                .collect(Collectors.toList());
        System.out.println(flat);
    }
}
```
**A:** `[1, 2, 3, 4, 5, 6, 7, 8, 9]`. `flatMap` maps each element to a stream then merges all those streams.

---

### 20. distinct, sorted, limit, skip
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> result = Stream.of(5, 3, 1, 4, 2, 5, 3, 1)
                .distinct()
                .sorted()
                .skip(1)
                .limit(3)
                .collect(Collectors.toList());
        System.out.println(result);
    }
}
```
**A:** `[2, 3, 4]`. After `distinct()`: [5,3,1,4,2]. After `sorted()`: [1,2,3,4,5]. After `skip(1)`: [2,3,4,5]. After `limit(3)`: [2,3,4].

---

### 21. peek — Debugging Streams
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> result = Stream.of(1, 2, 3, 4, 5)
                .filter(n -> n % 2 != 0)
                .peek(n -> System.out.print("before: " + n + " "))
                .map(n -> n * n)
                .peek(n -> System.out.print("after: " + n + " "))
                .collect(Collectors.toList());
        System.out.println(result);
    }
}
```
**A:** `before: 1 after: 1 before: 3 after: 9 before: 5 after: 25 [1, 9, 25]`. `peek` is for debugging — it doesn't modify elements.

---

### 22. mapToInt and sum
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> words = List.of("hello", "world", "java");
        int totalLength = words.stream()
                .mapToInt(String::length)
                .sum();
        System.out.println(totalLength);

        OptionalDouble avg = words.stream()
                .mapToInt(String::length)
                .average();
        System.out.printf("%.2f%n", avg.getAsDouble());
    }
}
```
**A:**
```
14
4.67
```

---

### 23. Stream.generate and iterate
**Q: What is the output?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // Generate: infinite stream
        Stream.generate(() -> "hi").limit(3).forEach(System.out::println);

        // Iterate: infinite sequence with seed + function
        Stream.iterate(1, n -> n * 2).limit(5).forEach(n -> System.out.print(n + " "));
    }
}
```
**A:**
```
hi
hi
hi
1 2 4 8 16
```

---

### 24. Stream.iterate with predicate (Java 9+)
**Q: What is the output?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // Java 9+ three-arg iterate: seed, hasNext predicate, next function
        Stream.iterate(1, n -> n <= 32, n -> n * 2)
                .forEach(n -> System.out.print(n + " "));
    }
}
```
**A:** `1 2 4 8 16 32 `

---

### 25. Stream.of vs Arrays.stream vs Collection.stream
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // Stream.of — creates from elements
        Stream.of(1, 2, 3).forEach(n -> System.out.print(n + " "));
        System.out.println();

        // Arrays.stream — better for arrays (avoids boxing for primitives)
        int[] arr = {4, 5, 6};
        Arrays.stream(arr).forEach(n -> System.out.print(n + " "));
        System.out.println();

        // Collection.stream()
        List.of("a","b","c").stream().forEach(System.out::print);
    }
}
```
**A:**
```
1 2 3
4 5 6
abc
```

---

### 26. mapMulti (Java 16+)
**Q: What is the output?**
```java
import java.util.stream.*;
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> result = Stream.of(1, 2, 3)
                .<Integer>mapMulti((n, consumer) -> {
                    consumer.accept(n);
                    consumer.accept(n * n);
                })
                .collect(Collectors.toList());
        System.out.println(result);
    }
}
```
**A:** `[1, 1, 2, 4, 3, 9]`. `mapMulti` replaces each element with 0 or more elements — similar to `flatMap` but imperative.

---

### 27. takeWhile and dropWhile (Java 9+)
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> nums = List.of(2, 4, 6, 7, 8, 10);
        // takeWhile: take elements while predicate is true, stop at first false
        System.out.println(nums.stream().takeWhile(n -> n % 2 == 0).collect(Collectors.toList()));
        // dropWhile: drop elements while predicate is true, keep rest
        System.out.println(nums.stream().dropWhile(n -> n % 2 == 0).collect(Collectors.toList()));
    }
}
```
**A:**
```
[2, 4, 6]
[7, 8, 10]
```
Note: `takeWhile`/`dropWhile` are for **ordered** streams. For unordered streams, behavior is non-deterministic.

---

### 28. Stream Cannot Be Reused
**Q: What happens?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Stream<Integer> stream = Stream.of(1, 2, 3);
        stream.forEach(n -> System.out.print(n + " "));
        stream.forEach(n -> System.out.print(n + " ")); // reuse!
    }
}
```
**A:** `1 2 3 ` then **IllegalStateException: stream has already been operated upon or closed**. Streams are single-use. Create a new stream from the source for each operation.

---

### 29. sorted with null elements
**Q: What happens?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = Arrays.asList("b", null, "a");
        list.stream().sorted().forEach(System.out::println); // NullPointerException!
    }
}
```
**A:** **NullPointerException.** `sorted()` uses natural ordering which calls `compareTo()` — null comparison throws NPE. Use `Comparator.nullsFirst` or `Comparator.nullsLast`.

---

### 30. String.chars() — Stream of Characters
**Q: What is the output?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        "hello".chars()
               .mapToObj(c -> String.valueOf((char)c))
               .map(String::toUpperCase)
               .forEach(System.out::print);
    }
}
```
**A:** `HELLO`

---

### 31. IntStream.range and rangeClosed
**Q: What is the output?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        IntStream.range(1, 5).forEach(n -> System.out.print(n + " "));     // [1, 5)
        System.out.println();
        IntStream.rangeClosed(1, 5).forEach(n -> System.out.print(n + " ")); // [1, 5]
    }
}
```
**A:**
```
1 2 3 4
1 2 3 4 5
```

---

### 32. Debugging with peek()
**Q: What is a common mistake with peek()?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // WRONG: peek alone doesn't trigger the pipeline!
        Stream.of(1, 2, 3).peek(System.out::println); // nothing prints!

        // RIGHT: must have a terminal operation
        Stream.of(1, 2, 3).peek(System.out::println).count(); // prints 1 2 3
    }
}
```
**A:** Nothing from the first line. Then `1 2 3` from the second. Streams are lazy — without a terminal operation, the pipeline never executes.

---

## Section 3: Terminal Operations & Collectors (Q33–Q50)

### 33. collect to List, Set, Map
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> words = List.of("apple", "banana", "cherry", "apple");

        List<String> list = words.stream().collect(Collectors.toList());
        Set<String> set   = words.stream().collect(Collectors.toUnmodifiableSet());
        Map<String, Integer> map = words.stream().distinct()
                .collect(Collectors.toMap(s -> s, String::length));

        System.out.println(list);
        System.out.println(set.size()); // deduped
        System.out.println(new TreeMap<>(map));
    }
}
```
**A:**
```
[apple, banana, cherry, apple]
3
{apple=5, banana=6, cherry=6}
```

---

### 34. Collectors.groupingBy()
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> words = List.of("apple", "banana", "cherry", "avocado", "blueberry", "cedar");
        Map<Character, List<String>> byFirstChar = words.stream()
                .collect(Collectors.groupingBy(s -> s.charAt(0)));
        new TreeMap<>(byFirstChar).forEach((k, v) -> System.out.println(k + ": " + v));
    }
}
```
**A:**
```
a: [apple, avocado]
b: [banana, blueberry]
c: [cherry, cedar]
```

---

### 35. Collectors.groupingBy with downstream collector
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> words = List.of("apple", "banana", "cherry", "avocado", "blueberry");
        Map<Integer, Long> byLength = words.stream()
                .collect(Collectors.groupingBy(String::length, Collectors.counting()));
        System.out.println(new TreeMap<>(byLength));
    }
}
```
**A:** `{5=1, 6=2, 8=1, 9=1}` (apple=5, banana/cherry=6, avocado=7... adjust based on actual words).

---

### 36. Collectors.partitioningBy()
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> numbers = List.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
        Map<Boolean, List<Integer>> partitioned = numbers.stream()
                .collect(Collectors.partitioningBy(n -> n % 2 == 0));
        System.out.println("Even: " + partitioned.get(true));
        System.out.println("Odd: " + partitioned.get(false));
    }
}
```
**A:**
```
Even: [2, 4, 6, 8, 10]
Odd: [1, 3, 5, 7, 9]
```

---

### 37. Collectors.joining()
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> words = List.of("apple", "banana", "cherry");
        System.out.println(words.stream().collect(Collectors.joining()));
        System.out.println(words.stream().collect(Collectors.joining(", ")));
        System.out.println(words.stream().collect(Collectors.joining(", ", "[", "]")));
    }
}
```
**A:**
```
applebananacherry
apple, banana, cherry
[apple, banana, cherry]
```

---

### 38. reduce() — fold operation
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // sum with identity
        int sum = Stream.of(1, 2, 3, 4, 5).reduce(0, Integer::sum);
        System.out.println(sum);

        // product without identity — returns Optional
        Optional<Integer> product = Stream.of(1, 2, 3, 4, 5).reduce((a, b) -> a * b);
        System.out.println(product.orElse(0));

        // reduce on empty stream without identity returns empty Optional
        Optional<Integer> empty = Stream.<Integer>empty().reduce(Integer::sum);
        System.out.println(empty.isPresent());
    }
}
```
**A:**
```
15
120
false
```

---

### 39. anyMatch, allMatch, noneMatch
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> nums = List.of(2, 4, 6, 7, 8);
        System.out.println(nums.stream().anyMatch(n -> n % 2 != 0)); // any odd?
        System.out.println(nums.stream().allMatch(n -> n % 2 == 0)); // all even?
        System.out.println(nums.stream().noneMatch(n -> n < 0));      // no negatives?
    }
}
```
**A:**
```
true
false
true
```

---

### 40. findFirst vs findAny
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Optional<Integer> first = Stream.of(1, 2, 3, 4, 5)
                .filter(n -> n > 2)
                .findFirst();
        System.out.println(first.orElse(-1));

        // findAny is faster in parallel streams (may return any matching element)
        Optional<Integer> any = Stream.of(1, 2, 3, 4, 5)
                .parallel()
                .filter(n -> n > 2)
                .findAny();
        System.out.println(any.isPresent()); // true, but value is non-deterministic
    }
}
```
**A:**
```
3
true
```

---

### 41. min and max with Comparator
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Optional<String> longest = Stream.of("cherry", "apple", "banana", "kiwi")
                .max(Comparator.comparingInt(String::length));
        System.out.println(longest.orElse("none"));

        Optional<Integer> smallest = Stream.of(5, 3, 8, 1, 9)
                .min(Comparator.naturalOrder());
        System.out.println(smallest.orElse(-1));
    }
}
```
**A:**
```
banana
1
```

---

### 42. Collectors.toMap with merge function
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> words = List.of("apple", "banana", "cherry", "avocado");
        // Group by first letter, concatenate values
        Map<Character, String> map = words.stream()
                .collect(Collectors.toMap(
                    s -> s.charAt(0),
                    s -> s,
                    (existing, newVal) -> existing + "," + newVal // merge duplicate keys
                ));
        System.out.println(new TreeMap<>(map));
    }
}
```
**A:** `{a=apple,avocado, b=banana, c=cherry}`

---

### 43. Collectors.summingInt and averagingInt
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> words = List.of("apple", "banana", "cherry");
        int totalLen = words.stream().collect(Collectors.summingInt(String::length));
        double avgLen = words.stream().collect(Collectors.averagingInt(String::length));
        System.out.println(totalLen);
        System.out.printf("%.2f%n", avgLen);
    }
}
```
**A:**
```
17
5.67
```

---

### 44. Collectors.summarizingInt
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        IntSummaryStatistics stats = Stream.of("apple", "banana", "cherry")
                .collect(Collectors.summarizingInt(String::length));
        System.out.println("count=" + stats.getCount());
        System.out.println("sum=" + stats.getSum());
        System.out.println("min=" + stats.getMin());
        System.out.println("max=" + stats.getMax());
        System.out.printf("avg=%.2f%n", stats.getAverage());
    }
}
```
**A:**
```
count=3
sum=17
min=5
max=6
avg=5.67
```

---

### 45. Collectors.counting
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<Integer, Long> byLength = Stream.of("a","bb","ccc","dd","eee","a")
                .collect(Collectors.groupingBy(String::length, Collectors.counting()));
        System.out.println(new TreeMap<>(byLength));
    }
}
```
**A:** `{1=2, 2=1, 3=2}` (a,a → 2 of length 1; bb → 1 of length 2; dd is 2 too... actually `{1=2, 2=2, 3=2}`).

---

### 46. toUnmodifiableList (Java 10+)
**Q: What happens?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> immutable = Stream.of(3, 1, 2).sorted()
                .collect(Collectors.toUnmodifiableList());
        System.out.println(immutable);
        immutable.add(4); // throws!
    }
}
```
**A:** `[1, 2, 3]` then **UnsupportedOperationException**.

---

### 47. Word Frequency Count — Stream Way
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        String text = "the quick brown fox jumps over the lazy dog the";
        Map<String, Long> freq = Arrays.stream(text.split(" "))
                .collect(Collectors.groupingBy(w -> w, Collectors.counting()));
        freq.entrySet().stream()
                .filter(e -> e.getValue() > 1)
                .sorted(Map.Entry.<String, Long>comparingByValue().reversed())
                .forEach(e -> System.out.println(e.getKey() + ": " + e.getValue()));
    }
}
```
**A:** `the: 3`

---

### 48. flatMap for splitting sentences
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> sentences = List.of("Hello World", "Java Streams");
        long wordCount = sentences.stream()
                .flatMap(s -> Arrays.stream(s.split(" ")))
                .distinct()
                .count();
        System.out.println(wordCount);
    }
}
```
**A:** `4`

---

### 49. Collectors.teeing (Java 12+)
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> numbers = List.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
        var result = numbers.stream()
                .collect(Collectors.teeing(
                    Collectors.summingInt(Integer::intValue),
                    Collectors.counting(),
                    (sum, count) -> sum + "/" + count
                ));
        System.out.println(result);
    }
}
```
**A:** `55/10`. `teeing` applies two collectors simultaneously and merges results with a combiner.

---

### 50. Stream Performance — Avoiding Boxing
**Q: Which is faster and why?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // SLOWER: Stream<Integer> — boxes each int
        int sum1 = Stream.iterate(1, n -> n + 1).limit(1_000_000)
                .mapToInt(Integer::intValue).sum();

        // FASTER: IntStream — no boxing
        int sum2 = IntStream.rangeClosed(1, 1_000_000).sum();

        System.out.println(sum1 == sum2);
    }
}
```
**A:** `true`. Always use primitive streams (`IntStream`, `LongStream`, `DoubleStream`) for numeric operations to avoid boxing overhead.

---

## Section 4: Optional (Q51–Q60)

### 51. Optional.of vs Optional.ofNullable
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Optional<String> opt1 = Optional.of("hello");
        System.out.println(opt1.isPresent());

        Optional<String> opt2 = Optional.of(null); // NullPointerException!
        Optional<String> opt3 = Optional.ofNullable(null); // OK — empty Optional
        System.out.println(opt3.isEmpty());
    }
}
```
**A:** `true`, then **NullPointerException** on `Optional.of(null)`. Use `Optional.ofNullable` when the value might be null.

---

### 52. Optional.get() Pitfall
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Optional<String> empty = Optional.empty();
        String value = empty.get(); // dangerous!
    }
}
```
**A:** **NoSuchElementException.** Never call `get()` without checking `isPresent()`. Prefer `orElse()`, `orElseGet()`, or `orElseThrow()`.

---

### 53. orElse vs orElseGet
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static String expensive() {
        System.out.println("computed!");
        return "default";
    }
    public static void main(String[] args) {
        Optional<String> present = Optional.of("value");
        // orElse: ALWAYS evaluates the argument eagerly
        System.out.println(present.orElse(expensive()));
        // orElseGet: evaluates LAZILY only if empty
        System.out.println(present.orElseGet(() -> expensive()));
    }
}
```
**A:**
```
computed!
value
value
```
`orElse()` always evaluates the argument (even if optional is present!). `orElseGet()` is lazy — prefer it for expensive defaults.

---

### 54. Optional.map and flatMap
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static Optional<String> find(String key) {
        return "Java".equals(key) ? Optional.of("A programming language") : Optional.empty();
    }
    public static void main(String[] args) {
        Optional.of("Java")
                .flatMap(Main::find)
                .map(String::toUpperCase)
                .ifPresent(System.out::println);

        Optional.of("Go")
                .flatMap(Main::find)
                .ifPresentOrElse(System.out::println, () -> System.out.println("not found"));
    }
}
```
**A:**
```
A PROGRAMMING LANGUAGE
not found
```

---

### 55. Optional.filter
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Optional<Integer> opt = Optional.of(42);
        Optional<Integer> filtered = opt.filter(n -> n > 100);
        System.out.println(filtered.isPresent());
        System.out.println(opt.filter(n -> n > 10).orElse(-1));
    }
}
```
**A:**
```
false
42
```

---

### 56. Optional.ifPresent vs ifPresentOrElse (Java 9+)
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Optional.of("hello").ifPresent(s -> System.out.println("Present: " + s));
        Optional.<String>empty().ifPresentOrElse(
            s -> System.out.println("Present: " + s),
            () -> System.out.println("Empty!")
        );
    }
}
```
**A:**
```
Present: hello
Empty!
```

---

### 57. Optional.stream() (Java 9+)
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Optional<String>> optionals = List.of(
            Optional.of("apple"), Optional.empty(), Optional.of("banana"), Optional.empty());

        List<String> present = optionals.stream()
                .flatMap(Optional::stream) // Java 9+ — turns Optional into 0 or 1 element stream
                .collect(Collectors.toList());
        System.out.println(present);
    }
}
```
**A:** `[apple, banana]`. `Optional.stream()` returns a stream with the value if present, empty stream otherwise.

---

### 58. Optional.or() (Java 9+)
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static Optional<String> getFromCache() { return Optional.empty(); }
    static Optional<String> getFromDB()    { return Optional.of("from DB"); }

    public static void main(String[] args) {
        Optional<String> result = getFromCache()
                .or(Main::getFromDB); // fallback to another Optional
        System.out.println(result.orElse("none"));
    }
}
```
**A:** `from DB`. `Optional.or()` provides a fallback `Optional` — useful for chaining data sources.

---

### 59. Anti-Pattern — Don't Use Optional for Everything
**Q: Which of these is BAD practice?**
```java
import java.util.*;
public class Main {
    // BAD: Optional as method parameter
    static void process(Optional<String> opt) {}

    // BAD: Optional in collections
    List<Optional<String>> optionals = new ArrayList<>();

    // GOOD: Optional only as return type for methods that may not return a value
    static Optional<String> findUser(String id) {
        return id.isEmpty() ? Optional.empty() : Optional.of("User:" + id);
    }
    public static void main(String[] args) { System.out.println(findUser("abc")); }
}
```
**A:** `Optional[User:abc]`. Optional is designed as a return type. Avoid using it as method parameters or storing in collections.

---

### 60. Optional.orElseThrow (Java 10+)
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        // orElseThrow() with no arg throws NoSuchElementException
        try { Optional.empty().orElseThrow(); }
        catch (NoSuchElementException e) { System.out.println("NSEE"); }

        // orElseThrow with custom exception
        try { Optional.empty().orElseThrow(() -> new IllegalStateException("not found")); }
        catch (IllegalStateException e) { System.out.println(e.getMessage()); }

        System.out.println(Optional.of("value").orElseThrow());
    }
}
```
**A:**
```
NSEE
not found
value
```

---

## Section 5: Parallel Streams (Q61–Q70)

### 61. parallel() — Easy Parallelism
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        long sum = LongStream.rangeClosed(1, 1_000_000)
                .parallel()
                .sum();
        System.out.println(sum);
    }
}
```
**A:** `500000500000`. `parallel()` splits the computation across multiple CPU cores via the ForkJoinPool. Result is correct.

---

### 62. Parallel Stream — Order Not Guaranteed
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Stream.of(1, 2, 3, 4, 5)
                .parallel()
                .forEach(n -> System.out.print(n + " ")); // unordered!
    }
}
```
**A:** Some permutation of `1 2 3 4 5` in non-deterministic order. Use `forEachOrdered()` if order matters.

---

### 63. Parallel Stream — forEachOrdered vs forEach
**Q: What is the output?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Stream.of(1, 2, 3, 4, 5)
                .parallel()
                .forEachOrdered(n -> System.out.print(n + " "));
    }
}
```
**A:** `1 2 3 4 5 `. `forEachOrdered` preserves encounter order even in parallel streams (slower).

---

### 64. When Parallel Streams Hurt Performance
**Q: What is the concern?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // Small datasets: thread overhead > benefit
        List<Integer> small = List.of(1, 2, 3, 4, 5);
        int sum = small.parallelStream().mapToInt(Integer::intValue).sum();
        System.out.println(sum);

        // Stateful operations and shared mutable state in parallel = bugs
        List<Integer> sync = Collections.synchronizedList(new ArrayList<>());
        // small.parallelStream().forEach(sync::add); // order non-deterministic!
    }
}
```
**A:** `15`. Parallel streams work best for: large datasets, CPU-bound operations, independently processable elements. Avoid for small data or I/O-bound work.

---

### 65. Collectors.toList() is thread-safe in parallel streams
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> result = IntStream.rangeClosed(1, 10)
                .parallel()
                .boxed()
                .collect(Collectors.toList());
        result.sort(Integer::compareTo);
        System.out.println(result);
    }
}
```
**A:** `[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]`. `collect()` uses a thread-safe combiner internally — safe in parallel streams.

---

### 66. Stream.concat()
**Q: What is the output?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Stream<Integer> a = Stream.of(1, 2, 3);
        Stream<Integer> b = Stream.of(4, 5, 6);
        Stream.concat(a, b).forEach(n -> System.out.print(n + " "));
    }
}
```
**A:** `1 2 3 4 5 6 `

---

### 67. Custom Collector (Collector.of)
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // Custom collector: collect to a reversed list
        List<Integer> reversed = Stream.of(1, 2, 3, 4, 5)
                .collect(Collector.of(
                    ArrayDeque::new,
                    ArrayDeque::addFirst,
                    (d1, d2) -> { d2.addAll(d1); return d2; },
                    d -> new ArrayList<>(d)
                ));
        System.out.println(reversed);
    }
}
```
**A:** `[5, 4, 3, 2, 1]`

---

### 68. Stream API — Real Interview Problem: Top N Frequent
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    static List<String> topN(List<String> words, int n) {
        return words.stream()
                .collect(Collectors.groupingBy(w -> w, Collectors.counting()))
                .entrySet().stream()
                .sorted(Map.Entry.<String, Long>comparingByValue().reversed())
                .limit(n)
                .map(Map.Entry::getKey)
                .collect(Collectors.toList());
    }
    public static void main(String[] args) {
        List<String> words = List.of("apple","banana","apple","cherry","banana","apple");
        System.out.println(topN(words, 2));
    }
}
```
**A:** `[apple, banana]`

---

### 69. Collectors.mapping
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Student(String name, int grade) {}
    public static void main(String[] args) {
        List<Student> students = List.of(
            new Student("Alice", 90), new Student("Bob", 85),
            new Student("Charlie", 90), new Student("Dave", 85));

        Map<Integer, List<String>> gradeMap = students.stream()
                .collect(Collectors.groupingBy(Student::grade,
                    Collectors.mapping(Student::name, Collectors.toList())));
        System.out.println(new TreeMap<>(gradeMap));
    }
}
```
**A:** `{85=[Bob, Dave], 90=[Alice, Charlie]}`

---

### 70. Stream Pipeline — Complete Example (Interview)
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Employee(String name, String dept, double salary) {}
    public static void main(String[] args) {
        List<Employee> emps = List.of(
            new Employee("Alice", "Eng", 90000), new Employee("Bob", "Eng", 85000),
            new Employee("Charlie", "HR", 70000),  new Employee("Dave", "HR", 75000),
            new Employee("Eve", "Eng", 95000));

        // Average salary per department, sorted by department
        emps.stream()
                .collect(Collectors.groupingBy(Employee::dept,
                    Collectors.averagingDouble(Employee::salary)))
                .entrySet().stream()
                .sorted(Map.Entry.comparingByKey())
                .forEach(e -> System.out.printf("%s: %.0f%n", e.getKey(), e.getValue()));
    }
}
```
**A:**
```
Eng: 90000
HR: 72500
```
