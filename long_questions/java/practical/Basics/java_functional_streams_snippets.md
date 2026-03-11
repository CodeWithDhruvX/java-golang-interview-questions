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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this lambda work?"

**Your Response:** "The output is '[apple, banana, cherry]'. This demonstrates a basic lambda expression implementing a Comparator. The lambda `(a, b) -> a.compareTo(b)` takes two string parameters and returns the result of comparing them. This is equivalent to writing an anonymous class but much more concise. The sort method expects a Comparator, and the lambda provides that implementation. This is the essence of functional programming in Java - passing behavior as data. We could also write this as `Comparator.naturalOrder()` or `String::compareTo` using method references."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what do these Predicate operations do?"

**Your Response:** "The output shows 'true', 'false', 'true', 'true', 'false'. This demonstrates the Predicate functional interface and its composition methods. `isEven.test(4)` returns true because 4 is even. `isPositive.test(-3)` returns false because -3 is not positive. The interesting part is the composition: `isEven.and(isPositive).test(4)` returns true only if both conditions are met (4 is both even and positive). `isEven.or(isPositive).test(-4)` returns true because -4 is even (OR logic). `isEven.negate().test(4)` returns false because it's the opposite of isEven. Predicate composition is powerful for building complex conditions."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does function composition work?"

**Your Response:** "The output is 'false' then 'true'. This shows function composition using the Function interface. We have two functions: `length` converts a String to its length, and `isLong` checks if an Integer is greater than 5. Using `andThen()`, we compose them into `isLongStr` which first applies the length function, then applies the isLong function to the result. So 'hello' has length 5, which is not > 5, so false. 'hello world' has length 11, which is > 5, so true. This is function chaining - the output of one function becomes the input of the next."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does the Consumer work?"

**Your Response:** "The output shows 'hello' then 'HELLO'. This demonstrates the Consumer functional interface, which represents an operation that takes an input and returns no result. We have two consumers: `print` uses a method reference to `System.out.println`, and `printUpper` prints the uppercase version. Using `andThen()`, we chain them together so `both.accept("hello")` first prints 'hello', then prints 'HELLO'. Consumers are perfect for side-effect operations like printing, logging, or modifying objects. The andThen method executes the first consumer, then the second."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the purpose of Supplier?"

**Your Response:** "The output shows 'Hello, World!' then 'true'. This demonstrates the Supplier functional interface, which represents a supplier of results - it takes no arguments but returns a value. The `greeting` supplier always returns the same string, while the `random` supplier uses a method reference to `Math.random` to generate random numbers. The second line is true because `Math.random()` returns a value between 0.0 (inclusive) and 1.0 (exclusive). Suppliers are great for factory methods, lazy initialization, or generating test data."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do BiFunction and BiPredicate work?"

**Your Response:** "The output shows 'ababab', 'true', then 'false'. This demonstrates binary functional interfaces that take two parameters. The `BiFunction<String, Integer, String>` repeats a string n times, so 'ab' repeated 3 times gives 'ababab'. The `BiPredicate<Integer, Integer>` checks if two integers sum to 10, so 3+7=10 is true, but 3+8=11 is false. These are useful for operations that need two inputs - like mathematical operations, comparisons, or validations. They're the two-parameter versions of Function and Predicate."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what are UnaryOperator and BinaryOperator?"

**Your Response:** "The output shows 'hello' then '7'. This demonstrates specialized functional interfaces. `UnaryOperator<String>` is a Function that takes and returns the same type - here it trims whitespace from strings. `BinaryOperator<Integer>` is a BiFunction where both inputs and output are the same type - here it adds two integers using the method reference `Integer::sum`. These are more type-safe and expressive than using Function/BiFunction when the types are the same. They're essentially syntactic sugar for common cases, making the code more readable."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what are the four types of method references?"

**Your Response:** "The output shows '10', 'Hello World', 'HELLO', then 'test'. This demonstrates all four types of method references in Java: 1) Static method reference `Main::doubleIt` calls a static method, 2) Instance method on specific object `prefix::concat` calls concat on the prefix string, 3) Instance method on arbitrary instance `String::toUpperCase` calls toUpperCase on whatever string is passed, and 4) Constructor reference `StringBuilder::new` creates new StringBuilder instances. Method references are syntactic sugar for lambdas - they make code more concise and readable when the lambda just calls an existing method."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the effectively final rule?"

**Your Response:** "Yes, this compiles and prints '15'. This demonstrates the 'effectively final' rule for lambda expressions. The variable `multiplier` can be used inside the lambda because it's effectively final - meaning it's never reassigned after initialization. Before Java 8, variables had to be explicitly declared `final`. Java 8 relaxed this to 'effectively final'. If we uncommented the line that changes multiplier to 4, the code wouldn't compile because the lambda would be capturing a non-final variable. This restriction exists because the lambda instance might outlive the method call."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's a custom functional interface?"

**Your Response:** "The output is '6'. This shows how to create custom functional interfaces. The `@FunctionalInterface` annotation marks `TriFunction` as a functional interface - it has exactly one abstract method `apply()` that takes three parameters and returns a result. This allows us to use lambda expressions with three parameters, which isn't possible with the built-in Function interface that only takes one parameter. Custom functional interfaces are useful when you need specific function signatures that aren't provided by Java's built-in ones. The annotation is optional but provides compile-time checking."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about Predicate.not()?"

**Your Response:** "The output is '[hello, world, !]'. This shows `Predicate.not()` which was added in Java 11. It's a static method that takes a predicate and returns its negation. Here we use `Predicate.not(String::isBlank)` to filter out blank strings. This is much cleaner and more readable than writing `s -> !s.isBlank()`. It's particularly useful with method references where you want to negate the predicate. This is part of Java's continued improvement of functional programming APIs, making code more expressive and less verbose."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what is currying?"

**Your Response:** "The output shows '8' then '15'. This demonstrates currying - a functional programming concept where we transform a function that takes multiple arguments into a chain of functions that each take a single argument. Here, `add` is a function that takes an integer and returns another function that takes an integer and returns their sum. We create `add5` by partially applying the first argument (5), then call it with different second arguments. Currying enables function composition and partial application, which are powerful techniques in functional programming."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does comparator chaining work?"

**Your Response:** "The output shows people sorted by age ascending, then by name alphabetically. The key is the chained comparator: `Comparator.comparingInt(Person::age).thenComparing(Person::name)`. First, it sorts by age using the specialized `comparingInt()` method for primitive types. For people with the same age (Bob and Dave are both 25, Alice and Charlie are both 30), it uses the secondary comparator to sort by name. Comparator chaining uses short-circuit evaluation - it only moves to the next comparator if the previous one returns 0 (equal). This is perfect for multi-level sorting requirements."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why use primitive functional interfaces?"

**Your Response:** "The output shows '25' then '7'. This demonstrates primitive functional interfaces that avoid autoboxing overhead. `IntUnaryOperator` works with primitive int values directly, while `IntBinaryOperator` takes two int parameters. Using these instead of `Function<Integer, Integer>` avoids the cost of boxing primitives into objects and unboxing them back. This performance difference matters in high-frequency operations like stream processing or numerical computations. Java provides primitive specializations for int, long, and double types. They're essentially the same as their generic counterparts but work directly with primitives."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's Function.identity() used for?"

**Your Response:** "The output shows 'hello' then 'world'. `Function.identity()` returns a function that always returns its input unchanged - it's essentially the identity function. This might seem trivial, but it's very useful in functional programming and stream operations. For example, when you need to group by the element itself in a stream, you'd use `Collectors.groupingBy(Function.identity())`. It's also used as a default transformer or when you need a Function but don't want to modify the input. It's the functional equivalent of doing nothing, which is sometimes exactly what you need."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does lambda closure work with instance variables?"

**Your Response:** "The output shows 'lambda val: 99' then 'anon val: 99'. This demonstrates that both lambdas and anonymous classes capture instance fields differently from local variables. The 'val' here is an instance field, not a local variable, so both the lambda and anonymous class capture a reference to 'this' and access the current value of the field when they execute. The effectively final restriction only applies to local variables, not instance fields. This is why both see the updated value 99 instead of 10. Instance fields are mutable and accessible to both lambdas and anonymous classes through the implicit 'this' reference."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does lazy evaluation mean in streams?"

**Your Response:** "The output shows the interleaved processing 'F1 F2 M2 F3 F4 M4 F5 F6 M6' then 'count: 3'. This demonstrates stream lazy evaluation. The filter and map operations don't process all elements at once - they process elements one by one as needed. When `limit(3)` is called, the stream only processes enough elements to find 3 matching items. You can see the interleaved 'F' (filter) and 'M' (map) operations as each element flows through the pipeline. This lazy evaluation makes streams efficient - they don't do unnecessary work. The count is 3 because we limited to 3 even numbers."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do these stream operations work together?"

**Your Response:** "The output is '[BANANA, CHERRY, ELDERBERRY]'. This shows a classic stream pipeline. First, `filter()` keeps only strings longer than 5 characters, removing 'apple' and 'date'. Then `map()` transforms each remaining string to uppercase using a method reference. Next, `sorted()` orders the results alphabetically. Finally, `collect(Collectors.toList())` gathers the results into a list. This demonstrates the declarative nature of streams - we describe what we want, not how to do it. Each operation is lazy and composable, making the code readable and efficient."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does flatMap work?"

**Your Response:** "The output is '[1, 2, 3, 4, 5, 6, 7, 8, 9]'. This demonstrates `flatMap`, which is one of the most powerful stream operations. We start with a list of lists (nested structure), and `flatMap` transforms it into a single flat list. The `Collection::stream` method reference converts each inner list to a stream, and `flatMap` merges all those streams into one. This is incredibly useful for working with nested data structures, like lists of lists, or when you want to flatten one-to-many relationships. It's the stream equivalent of nested loops but much more declarative."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do these stream operations work together?"

**Your Response:** "The output is '[2, 3, 4]'. This shows how stream operations chain together. Starting with `[5,3,1,4,2,5,3,1]`, `distinct()` removes duplicates leaving `[5,3,1,4,2]`. `sorted()` orders them to `[1,2,3,4,5]`. `skip(1)` skips the first element, giving `[2,3,4,5]`. Finally `limit(3)` takes only the first 3 elements, resulting in `[2,3,4]`. These operations are composable and lazy, so they work together efficiently. This pattern is common for pagination - skip for offset, limit for page size."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the purpose of peek()?"

**Your Response:** "The output shows 'before: 1 after: 1 before: 3 after: 9 before: 5 after: 25' then '[1, 9, 25]'. This demonstrates `peek()`, which is primarily used for debugging streams. `peek()` allows you to perform an action on each element as it flows through the pipeline without modifying the element. Here we use it to see the values before and after the map operation. The stream filters odd numbers (1, 3, 5), then we peek to see the original values, map them to squares, then peek again to see the transformed values. `peek()` is great for debugging complex stream pipelines."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why use mapToInt()?"

**Your Response:** "The output shows '14' then '4.67'. This demonstrates `mapToInt()`, which converts a stream of objects to a primitive IntStream. Here we map each string to its length, then use `sum()` and `average()` which are only available on primitive streams. The advantage is avoiding autoboxing overhead - IntStream works with primitive int values directly. The total length is 5+5+4=14, and the average is 14/3=4.67. Primitive streams have specialized terminal operations like sum(), average(), min(), max() that aren't available on regular streams."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do generate() and iterate() work?"

**Your Response:** "The output shows 'hi' printed 3 times, then '1 2 4 8 16'. This demonstrates two ways to create infinite streams. `Stream.generate(() -> "hi")` creates an infinite stream by repeatedly calling a supplier - here it always returns "hi". We limit it to 3 elements. `Stream.iterate(1, n -> n * 2)` creates an infinite sequence starting with 1, where each element is generated by applying the function to the previous element. We limit it to 5 elements, giving us powers of 2. These are useful for generating test data, mathematical sequences, or when you need streams that don't come from existing collections."

---

### 24. Stream.iterate with predicate (Java 9+)
**Q: What is the output?**
```java
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        // Java 9+ three-arg iterate: seed, hasNext predicate, next function
        Stream.iterate(1, n -> n <= 10, n -> n * 2)
              .forEach(n -> System.out.print(n + " "));
    }
}
```
**A:** `1 2 4 8`. The three-arg `iterate` stops when the predicate fails.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does the three-arg iterate work?"

**Your Response:** "The output is '1 2 4 8'. This shows the Java 9+ version of `Stream.iterate()` that takes three arguments: a seed value, a predicate to determine when to stop, and a function to generate the next value. Unlike the two-arg version which creates infinite streams, this version stops when the predicate (`n -> n <= 10`) returns false. It starts with 1, doubles each time (1, 2, 4, 8, 16...), but stops at 8 because 16 would exceed 10. This is safer than the infinite version as it prevents unintentional infinite streams."
### 25. Creating Streams (Stream.of, Arrays.stream, Collection.stream)
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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what are the different ways to create streams?"

**Your Response:** "The output shows '1 2 3', '4 5 6', then 'abc'. This demonstrates three common ways to create streams. `Stream.of()` creates a stream from individual elements. `Arrays.stream()` creates a stream from an array - it's more efficient for primitive arrays because it avoids boxing. `Collection.stream()` creates a stream from any collection. Each method has its use case: Stream.of() for a few known elements, Arrays.stream() for existing arrays, and Collection.stream() for collections. Choosing the right method depends on your data source and performance needs."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does mapMulti work?"

**Your Response:** "The output is '[1, 1, 2, 4, 3, 9]'. This demonstrates `mapMulti`, introduced in Java 16, which is an alternative to `flatMap` for one-to-many transformations. For each input element, we can emit 0, 1, or more output elements using the provided consumer. Here, for each number n, we emit both n and n². So for 1 we emit [1,1], for 2 we emit [2,4], and for 3 we emit [3,9]. Unlike `flatMap` which requires returning a stream, `mapMulti` uses an imperative style with a consumer, which can be more efficient and readable for complex transformations."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do takeWhile and dropWhile work?"

**Your Response:** "The output shows '[2, 4, 6]' then '[7, 8, 10]'. This demonstrates Java 9's `takeWhile` and `dropWhile` operations. `takeWhile` takes elements from the beginning while the predicate is true, stopping at the first false. It takes [2, 4, 6] but stops when it encounters 7 (odd). `dropWhile` does the opposite - it drops elements from the beginning while the predicate is true, then keeps the rest. It drops [2, 4, 6] and keeps [7, 8, 10]. These operations are different from filter - they're short-circuiting and depend on the order of elements. They're particularly useful for processing sorted or ordered data streams."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What happens and why can't streams be reused?"

**Your Response:** "The output shows '1 2 3' then an IllegalStateException. This demonstrates that streams are single-use - once you call a terminal operation like forEach(), the stream is closed and cannot be used again. The reason is that streams are designed to be lazy and efficient, so they maintain internal state about where they are in processing. After a terminal operation, that state is consumed. If you need to process the same data multiple times, either create a new stream from the source or collect the results to a collection for reuse. This is a fundamental design principle of the Stream API."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What happens and how do you handle nulls in sorting?"

**Your Response:** "This throws a NullPointerException because `sorted()` uses natural ordering which calls `compareTo()` on each element. When it encounters a null value, it tries to call `compareTo()` on null, which throws NPE. The solution is to use null-safe comparators: `Comparator.nullsFirst()` puts nulls at the beginning, or `Comparator.nullsLast()` puts nulls at the end. This is a common issue when working with data that might contain nulls. These null-safe comparators make your code more robust and prevent runtime exceptions when sorting mixed data."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does String.chars() work?"

**Your Response:** "The output is 'HELLO'. This demonstrates `String.chars()` which returns an IntStream of the character values (Unicode code points) in the string. We then use `mapToObj()` to convert each int back to a String character, `map()` to convert to uppercase, and `forEach()` to print. The key insight is that `chars()` returns an IntStream of int values, not a Stream<Character>, so we need to convert them back to characters. This is useful for character-level processing of strings, like counting specific characters, filtering, or transforming each character individually."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between range and rangeClosed?"

**Your Response:** "The output shows '1 2 3 4' then '1 2 3 4 5'. This demonstrates the difference between `IntStream.range()` and `IntStream.rangeClosed()`. `range(1, 5)` creates a stream from 1 (inclusive) to 5 (exclusive), so it includes [1, 2, 3, 4]. `rangeClosed(1, 5)` creates a stream from 1 (inclusive) to 5 (inclusive), so it includes [1, 2, 3, 4, 5]. The difference is whether the upper bound is inclusive or exclusive. This is similar to Python's range function or traditional for-loop semantics. `rangeClosed` is useful when you want to include the end point, while `range` follows the more common exclusive upper bound pattern."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's a common mistake with peek() and what happens here?"

**Your Response:** "The first line prints nothing, while the second line prints '1 2 3'. This demonstrates a common mistake with `peek()` - it's an intermediate operation, not a terminal operation. Streams are lazy, meaning they only execute when a terminal operation like `forEach()`, `collect()`, or `count()` is called. The first call to `peek()` alone doesn't trigger the pipeline, so nothing happens. The second call includes `count()`, which is a terminal operation that triggers the entire pipeline, including the peek. This is why `peek()` is mainly for debugging - it lets you see values flowing through the pipeline without consuming them."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do these basic collectors work?"

**Your Response:** "The output shows the list with duplicates, then 3 (unique count), then a map of words to lengths. This demonstrates three basic collectors. `toList()` collects to a List, preserving duplicates and order. `toUnmodifiableSet()` creates an immutable Set, automatically removing duplicates. `toMap()` creates a Map where we specify the key and value mapping functions. Here we use `distinct()` before toMap to avoid duplicate key conflicts. Each collector serves different purposes: List for preserving order and duplicates, Set for uniqueness, and Map for key-value transformations."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does groupingBy work?"

**Your Response:** "The output shows words grouped by their first character. This demonstrates `Collectors.groupingBy()` which is one of the most powerful collectors. It takes a classification function and groups elements by the result of that function. Here we group by `s.charAt(0)`, so all words starting with 'a' go into one list, 'b' in another, etc. We use `TreeMap` to sort the output by key. `groupingBy` is extremely useful for categorizing data, creating frequency distributions, or any kind of bucket-based analysis. It's the stream equivalent of SQL's GROUP BY clause."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's a downstream collector?"

**Your Response:** "The output shows word lengths mapped to their counts. This demonstrates `groupingBy` with a downstream collector - a powerful pattern for multi-level aggregations. The first parameter groups by length, and the second parameter `Collectors.counting()` is applied to each group to count elements. So words of length 5 (apple) appear once, length 6 (banana, cherry) appear twice, etc. Downstream collectors let you perform additional operations on each group, like summing, averaging, or even nested grouping. This is much more efficient than collecting to intermediate collections and processing them separately."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does partitioningBy differ from groupingBy?"

**Your Response:** "The output shows numbers partitioned into even and odd groups. This demonstrates `Collectors.partitioningBy()` which is a special case of `groupingBy` that always creates exactly two groups: true and false. Unlike `groupingBy` which can create any number of groups based on the classification function, `partitioningBy` is optimized for boolean predicates. It always returns a Map<Boolean, List<T>> with exactly two entries. This is perfect for yes/no categorizations like pass/fail, valid/invalid, or in this case even/odd. It's more efficient and semantically clearer than using groupingBy for binary classifications."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does the joining collector work?"

**Your Response:** "The output shows three different ways to join strings. This demonstrates `Collectors.joining()` which is perfect for concatenating stream elements into a single string. The first version uses no delimiter, just concatenating everything. The second adds a comma and space between elements. The third adds both delimiter and prefix/suffix - like creating JSON array format. `joining()` is much more efficient than manually building strings with StringBuilder or using string concatenation in loops. It's particularly useful for creating CSV output, JSON arrays, or human-readable lists from stream data."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does reduce work?"

**Your Response:** "The output shows '15', '120', then 'false'. This demonstrates the `reduce()` operation, which is a fundamental fold operation in functional programming. The first version uses an identity value (0) and accumulates the sum, always returning a result. The second version has no identity, so it returns an Optional to handle the case of empty streams. The third shows that reducing an empty stream without identity returns an empty Optional. `reduce()` is powerful because it can implement many operations - sum, product, min, max, or any custom aggregation. It's the building block for many terminal operations."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between these matching operations?"

**Your Response:** "The output shows 'true', 'false', then 'true'. This demonstrates the three matching operations: `anyMatch`, `allMatch`, and `noneMatch`. `anyMatch` returns true if at least one element matches the predicate - here we have an odd number (7), so true. `allMatch` returns true only if ALL elements match - not all are even, so false. `noneMatch` returns true if NO elements match - there are no negative numbers, so true. These operations are short-circuiting: `anyMatch` stops as soon as it finds a match, `allMatch` and `noneMatch` stop as soon as they find a counterexample. This makes them efficient for large streams."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between findFirst and findAny?"

**Your Response:** "The output shows '3' then 'true'. This demonstrates the difference between `findFirst()` and `findAny()`. `findFirst()` always returns the first matching element in encounter order - here it's 3, the first number greater than 2. `findAny()` can return any matching element, which in sequential streams is usually the first, but in parallel streams can be any. The key difference is performance: `findAny()` is more efficient in parallel streams because it doesn't have to wait for the first element. We use `parallel()` to demonstrate this. Both return Optional to handle the case where no elements match. Use `findFirst()` when order matters, `findAny()` when you just need any match."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do min and max work with comparators?"

**Your Response:** "The output shows 'banana' then '1'. This demonstrates `min()` and `max()` terminal operations which work with comparators. For `max()`, we use `Comparator.comparingInt(String::length)` to find the longest string - 'banana' with 6 characters. For `min()`, we use `Comparator.naturalOrder()` to find the smallest integer - 1. Both return Optional because the stream might be empty. These operations are more flexible than using `sorted()` and taking the first element because they don't require sorting the entire stream - they just need to track the minimum or maximum seen so far, making them more efficient for large streams."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the merge function in toMap?"

**Your Response:** "The output shows a map where duplicate keys are merged. This demonstrates `Collectors.toMap()` with a merge function to handle duplicate keys. When multiple elements have the same key (like 'apple' and 'avocado' both starting with 'a'), the merge function combines their values. Here we concatenate them with a comma: `existing + "," + newVal`. Without the merge function, this would throw an IllegalStateException for duplicate keys. This pattern is useful for grouping related data or when you want to aggregate values with the same key instead of just keeping the last one. The merge function gives you full control over how to resolve conflicts."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what do these numeric collectors do?"

**Your Response:** "The output shows '17' then '5.67'. This demonstrates specialized numeric collectors. `summingInt()` calculates the sum of integer values - here the total length of all words (5+6+6=17). `averagingInt()` calculates the average - 17/3=5.67. These are more efficient than using `mapToInt().sum()` or manually calculating averages because they work directly with primitive values and avoid boxing. They're particularly useful when you need statistical operations on stream data. Java provides similar collectors for `summingLong`, `summingDouble`, etc., making numeric aggregations straightforward and performant."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about summarizingInt?"

**Your Response:** "The output shows comprehensive statistics for word lengths. This demonstrates `Collectors.summarizingInt()` which returns an `IntSummaryStatistics` object containing count, sum, min, max, and average all in one operation. Instead of using multiple collectors to get each statistic separately, `summarizingInt()` computes everything in a single pass, making it very efficient. The object provides convenient getter methods for each statistic. This is perfect for generating quick data insights or reports. There are similar versions for `summarizingLong` and `summarizingDouble` for different numeric types. It's essentially a multi-purpose statistical collector in one."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does counting work as a downstream collector?"

**Your Response:** "The output shows a map of word lengths to their frequencies. This demonstrates `Collectors.counting()` used as a downstream collector with `groupingBy()`. We group words by their length, then count how many words are in each group. Two words have length 1 ('a', 'a'), one has length 2 ('bb'), and two have length 3 ('ccc', 'eee'). `counting()` returns a Long, representing the count of elements in each group. This pattern is extremely common for frequency analysis, histograms, or any kind of distribution analysis. It's much more efficient than manually counting with loops or nested operations."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What happens and why does this throw an exception?"

**Your Response:** "The output shows '[1, 2, 3]' then an UnsupportedOperationException. This demonstrates `Collectors.toUnmodifiableList()` introduced in Java 10, which creates an immutable collection. The list can be read and iterated, but any modification attempt like `add()`, `remove()`, or `set()` throws UnsupportedOperationException. This is useful for creating defensive copies, ensuring data integrity, or when you want to guarantee that a collection won't be modified accidentally. It's more efficient than `Collections.unmodifiableList(new ArrayList<>())` because it creates the immutable collection directly without the intermediate mutable list."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this word frequency count work?"

**Your Response:** "The output shows 'the: 3' because 'the' appears 3 times in the text. This demonstrates a classic word frequency count using streams. First, we split the text into words using `Arrays.stream(text.split(" "))`. Then we use `groupingBy()` with `counting()` to count occurrences of each word. Finally, we filter for words that appear more than once, sort by frequency in descending order, and print. This is a common pattern in text analysis, log processing, or data mining. The stream approach is much more concise and readable than traditional loop-based approaches, and it can be easily parallelized for large texts."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does flatMap work for text processing?"

**Your Response:** "The output is '4'. This demonstrates using `flatMap` for text processing. We start with a list of sentences, and `flatMap` transforms each sentence into a stream of words using `Arrays.stream(s.split(" "))`. The `flatMap` then merges all these word streams into a single stream of words. After removing duplicates with `distinct()`, we count the total unique words: 'Hello', 'World', 'Java', 'Streams'. This pattern is extremely useful for flattening nested structures, especially in text processing where you want to go from sentences to words, or from documents to paragraphs to sentences. It's much cleaner than nested loops."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about the teeing collector?"

**Your Response:** "The output is '55/10'. This demonstrates `Collectors.teeing()` introduced in Java 12, which allows you to apply two different collectors to the same stream simultaneously and then merge their results. Here we use `summingInt()` to get the sum (55) and `counting()` to get the count (10), then combine them with a simple string formatter. This is much more efficient than collecting the stream twice because it processes the elements only once. `teeing` is perfect for when you need multiple different aggregations from the same data, like getting both min and max, sum and average, or any combination of statistics."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Which is faster and why?"

**Your Response:** "Both return true (same result), but the second version is significantly faster. This demonstrates the performance impact of autoboxing. The first version uses `Stream<Integer>` which boxes each primitive int into an Integer object, causing memory overhead and garbage collection pressure. The second version uses `IntStream` which works directly with primitive int values, avoiding boxing entirely. For large datasets like 1 million elements, the performance difference can be substantial. Always prefer primitive streams for numeric operations - they have specialized terminal operations like `sum()`, `average()`, `min()`, `max()` and avoid the overhead of object creation."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What happens and what's the difference between of and ofNullable?"

**Your Response:** "The output shows 'true' then a NullPointerException. This demonstrates the critical difference between `Optional.of()` and `Optional.ofNullable()`. `Optional.of()` requires a non-null value and throws NPE if you pass null. `Optional.ofNullable()` accepts null values and returns an empty Optional instead. The key rule is: use `Optional.of()` when you're certain the value is non-null (for code clarity), and use `Optional.ofNullable()` when the value might be null. This is fundamental to Optional's purpose - representing values that may or may not be present without using null."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What happens and why is get() dangerous?"

**Your Response:** "This throws a NoSuchElementException. This demonstrates why `Optional.get()` is considered dangerous and is now deprecated in favor of safer alternatives. The `get()` method should only be called when you're absolutely certain the Optional contains a value, which defeats the purpose of Optional. Instead, use `orElse()` for simple defaults, `orElseGet()` for expensive default computation, or `orElseThrow()` for custom exceptions. Modern Java code should avoid `get()` entirely - it's essentially the same as dereferencing a null without the null check."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the performance difference?"

**Your Response:** "The output shows 'computed!' then 'value' twice. This demonstrates a critical performance difference between `orElse()` and `orElseGet()`. `orElse()` eagerly evaluates the default value even when the Optional contains a value, so 'computed!' prints even though we have 'value'. `orElseGet()` is lazy - it only evaluates the supplier when the Optional is empty. For expensive operations like database calls or complex calculations, always use `orElseGet()` to avoid unnecessary work. This is a common performance pitfall that can have significant impact in production code."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do map and flatMap work with Optional?"

**Your Response:** "The output shows 'A PROGRAMMING LANGUAGE' then 'not found'. This demonstrates `Optional.map()` and `flatMap()` for chaining operations. `flatMap()` is used when the mapping function returns another Optional - it flattens the result to avoid nested Optionals. Here we look up 'Java' which returns an Optional, then `map()` transforms the value to uppercase. For 'Go', the lookup returns empty, so `ifPresentOrElse()` executes the empty action. The key difference: `map()` transforms the value inside Optional, `flatMap()` transforms and returns another Optional, preventing nesting. This pattern is perfect for chaining operations that might fail."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Optional.filter work?"

**Your Response:** "The output shows 'false' then '42'. This demonstrates `Optional.filter()` which applies a predicate to the Optional value. If the value matches the predicate, it returns the same Optional; otherwise, it returns an empty Optional. Here, 42 doesn't match `n > 100`, so the filtered Optional is empty (false). But 42 does match `n > 10`, so `orElse(-1)` returns 42. This is useful for validating values before using them - like checking if a number is in range, a string meets certain criteria, or an object is in a valid state. It's more concise than checking with if statements."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between ifPresent and ifPresentOrElse?"

**Your Response:** "The output shows 'Present: hello' then 'Empty!'. This demonstrates the difference between `ifPresent()` and `ifPresentOrElse()` (Java 9+). `ifPresent()` only executes an action if the Optional contains a value - it does nothing for empty Optionals. `ifPresentOrElse()` takes two actions: one for when the value is present, and another for when it's empty. This is perfect for handling both cases without needing separate `ifPresent()` and `isEmpty()` checks. It's more expressive and eliminates the need for conditional logic when you want to handle both the present and empty cases."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about Optional.stream()?"

**Your Response:** "The output is '[apple, banana]'. This demonstrates `Optional.stream()` (Java 9+) which converts an Optional to a stream - either a stream with one element (if present) or an empty stream (if empty). This is incredibly useful when working with collections of Optionals because you can use `flatMap(Optional::stream)` to filter out empty Optionals and flatten the present values into a single stream. It's much cleaner than filtering with `filter(Optional::isPresent)` and then mapping with `map(Optional::get)`. This method bridges the gap between Optional and Stream APIs, making it easier to work with mixed data structures."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Optional.or() work?"

**Your Response:** "The output is 'from DB'. This demonstrates `Optional.or()` (Java 9+) which provides a fallback Optional when the first one is empty. Unlike `orElse()` which returns a raw value, `or()` returns another Optional, allowing you to chain multiple Optional-returning operations. Here we try to get from cache (empty), then fall back to database (has value). This pattern is perfect for trying multiple data sources in order - cache, then database, then default value. It's more composable than nested if-else statements and maintains the Optional contract throughout the chain."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Which of these is BAD practice and why?"

**Your Response:** "Using Optional as method parameters or in collections is considered bad practice. Optional was designed primarily as a return type to indicate that a method may not return a value. Using it as parameters adds unnecessary complexity and makes method signatures harder to work with. Storing Optionals in collections creates nested complexity - you're better off filtering out nulls or empty values at the source. The proper use case is what we see in `findUser()` - returning Optional to signal that a lookup might not find a result. This maintains clear intent and avoids the overhead of wrapping already-present values."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about orElseThrow()?"

**Your Response:** "The output shows 'NSEE', 'not found', then 'value'. This demonstrates `Optional.orElseThrow()` which comes in two forms (Java 10+). The no-arg version throws NoSuchElementException if empty. The version with a supplier throws a custom exception. This is the preferred way to handle cases where an empty Optional represents an error condition. It's more explicit than calling `get()` and allows for meaningful exception messages. The key advantage over `get()` is that you can provide context about what went wrong, making debugging and error handling much more effective in production code."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does parallel() work?"

**Your Response:** "The output is '500000500000', which is the sum of numbers from 1 to 1,000,000. This demonstrates how `parallel()` enables easy parallelism. The stream splits the work across multiple CPU cores using the ForkJoinPool, computes partial sums in parallel, then combines them. The result is identical to sequential processing but potentially much faster for large datasets. Parallel streams are perfect for CPU-intensive operations on large collections where the work can be divided into independent chunks. Java handles all the thread management and result combination automatically."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why is the order not guaranteed?"

**Your Response:** "The output shows the numbers 1-5 in some unpredictable order. This demonstrates that parallel streams don't preserve encounter order when using `forEach()`. Each thread processes its chunk of data independently, and the results are printed as soon as each thread finishes. The order depends on thread scheduling and which chunk completes first. If order matters, use `forEachOrdered()` which ensures the original encounter order is maintained, though it's slower because it needs to coordinate between threads. This is a fundamental trade-off in parallel processing: performance vs predictability."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does forEachOrdered work?"

**Your Response:** "The output shows '1 2 3 4 5' in the original order. This demonstrates `forEachOrdered()` which preserves encounter order even in parallel streams. Unlike `forEach()` which processes elements as soon as they're available, `forEachOrdered()` buffers results and outputs them in the original sequence. This coordination between threads adds overhead, making it slower than regular `forEach()`, but it guarantees order. Use `forEachOrdered()` when the order of processing matters (like generating reports or maintaining dependencies), but be aware of the performance cost. It's a trade-off between correctness and performance."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the concern with parallel streams here?"

**Your Response:** "The output is '15', but this demonstrates when parallel streams hurt performance. For small datasets (5 elements), the thread coordination overhead exceeds any parallel processing benefit. The key concerns are: 1) Small datasets - thread startup and coordination cost more than the work itself, 2) Shared mutable state - the commented code would cause race conditions, 3) I/O-bound operations - parallel streams don't help with waiting for external resources. Parallel streams shine with large datasets (thousands+ elements), CPU-intensive operations, and when work can be divided independently. Always benchmark - parallel isn't always faster."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and is this thread-safe?"

**Your Response:** "The output shows numbers 1-10 in order. This demonstrates that standard collectors like `Collectors.toList()` are thread-safe in parallel streams. The collector uses a combiner function internally to merge partial results from different threads safely. Each thread processes its chunk and creates a local list, then the combiner merges these lists. This is why you don't need synchronized collections when using standard collectors. However, custom operations on shared mutable state (like adding to an external list) would still need synchronization. The built-in collectors handle parallelism correctly."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Stream.concat work?"

**Your Response:** "The output shows '1 2 3 4 5 6'. This demonstrates `Stream.concat()` which merges two streams into one. It takes two stream parameters and returns a new stream that contains all elements from the first stream followed by all elements from the second stream. This is useful when you need to combine data from different sources or when you want to append/prepend elements to an existing stream. Note that once consumed, the original streams cannot be reused. Also, `concat()` preserves the order of both streams, making it predictable for ordered operations."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this custom collector work?"

**Your Response:** "The output shows a reversed list. This demonstrates creating a custom collector using `Collector.of()`. The collector has four parts: 1) Supplier (`ArrayDeque::new`) creates the result container, 2) Accumulator (`ArrayDeque::addFirst`) adds elements to the front, building a reversed order, 3) Combiner merges partial results from parallel threads, and 4) Finisher converts the ArrayDeque to an ArrayList. Custom collectors are powerful when you need specialized collection behavior not provided by built-in collectors. Here we achieve reversal during collection, which is more efficient than collecting then reversing."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this solve the top N frequent problem?"

**Your Response:** "The output is '[apple, banana]'. This demonstrates a classic interview problem: finding the top N most frequent items. The solution uses a multi-step stream pipeline: 1) `groupingBy` with `counting()` to create a frequency map, 2) `entrySet().stream()` to work with the map entries, 3) `sorted()` with `comparingByValue().reversed()` to sort by frequency descending, 4) `limit(n)` to get the top N, 5) `map()` to extract just the keys. This pattern is extremely common in data analysis and shows the power of combining multiple stream operations to solve complex problems elegantly."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Collectors.mapping work?"

**Your Response:** "The output shows students grouped by grade with just their names. This demonstrates `Collectors.mapping()` as a downstream collector. We group by grade, then use `mapping()` to transform each Student object to just its name before collecting to a list. `mapping()` is perfect when you want to extract specific fields from objects during grouping. Without it, we'd get `Map<Integer, List<Student>>`, but with `mapping()`, we get `Map<Integer, List<String>>`. This is cleaner than post-processing the grouped results and more efficient than mapping before grouping."

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

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this complete stream pipeline work?"

**Your Response:** "The output shows average salary per department. This demonstrates a complete real-world stream pipeline that combines multiple concepts. First, we group employees by department using `groupingBy()` with `averagingDouble()` to calculate average salaries. This gives us a Map of department to average salary. Then we create a new stream from the map's entrySet, sort by department name, and format the output. This pattern is typical in business analytics - grouping data, calculating aggregations, and presenting results. It shows how streams can elegantly handle complex data transformations that would require many lines of imperative code."
