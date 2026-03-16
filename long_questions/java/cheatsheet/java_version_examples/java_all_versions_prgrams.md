It sounds like you want **one single Java program that demonstrates many features from different Java versions in one file** (instead of many small snippets).
Below is a **single consolidated demo program** showing major concepts like **Generics, Lambda, Streams, Optional, Records, Pattern Matching, Virtual Threads, etc.**.

> ⚠️ This program requires **Java 21** or later because it includes modern features.

---

# ☕ Single Java Program Demonstrating Major Java Features

```java
import java.util.*;
import java.util.concurrent.*;
import java.util.stream.*;
import java.time.*;

// Record (Java 14+)
record Person(String name, int age) {}

// Sealed Interface (Java 17+)
sealed interface Shape permits Circle, Rectangle {}

record Circle(double radius) implements Shape {}
record Rectangle(double width, double height) implements Shape {}

public class JavaFeaturesDemo {

    // Generics + Varargs (Java 5)
    public static <T> void printItems(T... items) {
        for (T item : items) {
            System.out.println(item);
        }
    }

    // Pattern Matching for switch (Java 21)
    public static double area(Shape shape) {
        return switch (shape) {
            case Circle(double r) -> Math.PI * r * r;
            case Rectangle(double w, double h) -> w * h;
        };
    }

    public static void main(String[] args) throws Exception {

        System.out.println("=== Java Features Demo ===");

        // Collections + Generics
        List<String> names = new ArrayList<>();
        names.add("Alice");
        names.add("Bob");
        names.add("Charlie");

        // Lambda + Stream API (Java 8)
        List<String> filtered =
                names.stream()
                        .filter(n -> n.startsWith("A"))
                        .map(String::toUpperCase)
                        .collect(Collectors.toList());

        System.out.println("Stream Result: " + filtered);

        // Optional (Java 8)
        Optional<String> optional = Optional.ofNullable("Hello");
        optional.ifPresent(System.out::println);

        // Date Time API (Java 8)
        LocalDate today = LocalDate.now();
        System.out.println("Today: " + today);

        // Record (Java 14)
        Person p = new Person("John", 25);
        System.out.println("Person: " + p);

        // Pattern Matching
        Object obj = "Java Programming";
        if (obj instanceof String str) {
            System.out.println("String length: " + str.length());
        }

        // Sealed Classes + Pattern Switch
        Shape circle = new Circle(5);
        Shape rectangle = new Rectangle(4, 6);

        System.out.println("Circle Area: " + area(circle));
        System.out.println("Rectangle Area: " + area(rectangle));

        // Virtual Threads (Java 21)
        try (ExecutorService executor = Executors.newVirtualThreadPerTaskExecutor()) {
            Future<String> future = executor.submit(() -> {
                Thread.sleep(1000);
                return "Task executed in virtual thread";
            });

            System.out.println(future.get());
        }

        // Var (Java 10)
        var message = "Using var keyword";
        System.out.println(message);

        // Text Blocks (Java 15)
        String json = """
                {
                  "name": "Java",
                  "version": "21"
                }
                """;
        System.out.println(json);

        // Sequenced Collections (Java 21)
        List<String> sequenced = new ArrayList<>(List.of("Early", "Mid", "Latest"));
        System.out.println("First element: " + sequenced.getFirst());
        System.out.println("Last element: " + sequenced.getLast());

        // Switch Expressions (Java 14)
        int score = 95;
        String grade = switch (score / 10) {
            case 9, 10 -> "A";
            case 8 -> "B";
            default -> "C";
        };
        System.out.println("Grade from Switch Expression: " + grade);

        // Varargs method
        printItems("Java", "Python", "Go");

        System.out.println("=== Program Finished ===");
    }
}
```

---

# 🔎 Features Demonstrated in This Single Program

| Feature                     | Java Version |
| --------------------------- | ------------ |
| Generics                    | Java 5       |
| Varargs                     | Java 5       |
| Collections Framework       | Java 2       |
| Lambda Expressions          | Java 8       |
| Stream API                  | Java 8       |
| Optional                    | Java 8       |
| Date Time API               | Java 8       |
| Local Variable `var`        | Java 10      |
| Switch Expressions          | Java 14      |
| Records                     | Java 14      |
| Text Blocks                 | Java 15      |
| Sealed Classes              | Java 17      |
| Pattern Matching            | Java 17      |
| Sequenced Collections       | Java 21      |
| Pattern Matching for switch | Java 21      |
| Virtual Threads             | Java 21      |

---

✅ **This one program can be used to explain Java evolution in interviews.**

You can say:

> “This single program demonstrates the evolution of Java from Java 5 (Generics) to Java 21 (Virtual Threads and Pattern Matching).”

---

If you want, I can also give:

* **1 Ultimate Java Program covering 25+ features** (very powerful for interviews)
* **A single program explaining all Java OOP concepts**
* **A single program showing Java 8 features only** (most asked in interviews)
* **A visual Java version roadmap for interviews**

Just tell me. 🚀
