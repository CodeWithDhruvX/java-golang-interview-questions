# Core Java Comprehensive Cheatsheet

A quick reference guide for Java syntax, OOP concepts, collections, and modern features.

---

## 🟢 Basics

### Variables & Data Types
```java
// Primitives
int age = 30;
double price = 19.99;
boolean isActive = true;
char grade = 'A';

// Reference Types (Wrappers & String)
Integer count = 100;      // Wrapper for int
String name = "Java";     // Immutable string
```

### Type Casting
```java
// Widening (Automatic)
int myInt = 9;
double myDouble = myInt; // 9.0

// Narrowing (Manual)
double pi = 3.14;
int p = (int) pi;        // 3
```

### String Operations
```java
String s1 = "Hello";
String s2 = "World";

int len = s1.length();                // Length
String s3 = s1 + " " + s2;            // Concatenation
boolean eq = s1.equals("Hello");      // Comparison (Never use == for content)
String sub = s1.substring(0, 2);      // "He"
```

---

## 🟡 OOP Concepts

### Class Structure
```java
public class User {
    // Fields (Encapsulation)
    private String name;
    private int age;

    // Constructor
    public User(String name, int age) {
        this.name = name;
        this.age = age;
    }

    // Methods (Getters/Setters)
    public String getName() {
        return name;
    }
}
```

### Inheritance (`extends`)
```java
class Animal {
    void eat() { System.out.println("Eating..."); }
}

class Dog extends Animal {
    @Override
    void eat() { System.out.println("Dog eating..."); }
}
```

### Interfaces (`implements`)
Defines a contract.
```java
interface Printable {
    void print();
}

class Document implements Printable {
    public void print() {
        System.out.println("Printing doc...");
    }
}
```

### Abstract Classes
Can have both abstract (empty) and concrete methods.
```java
abstract class Shape {
    abstract void draw(); // Must be implemented
    void message() { System.out.println("drawing..."); }
}
```

---

## 🔵 Collections Framework

### List
Ordered collection, allows duplicates.
```java
List<String> list = new ArrayList<>();

// Operations
list.add("Apple");          // Add
list.add("Banana");
String fruit = list.get(0); // Access: "Apple"
list.remove(0);             // Remove

// Iterate
for (String item : list) {
    System.out.println(item);
}
```

### Set
Unordered, unique elements.
```java
Set<Integer> set = new HashSet<>();

set.add(10);
set.add(10); // Duplicate ignored
set.add(20);

// Check existence
boolean hasTen = set.contains(10); // true
```

### Map
Key-Value pairs.
```java
Map<String, Integer> scores = new HashMap<>();

// Operations
scores.put("Alice", 90);    // Add/Update
scores.put("Bob", 85);

int aliceScore = scores.get("Alice"); // 90
scores.remove("Bob");       // Remove

// Iterate
for (Map.Entry<String, Integer> entry : scores.entrySet()) {
    System.out.println(entry.getKey() + ": " + entry.getValue());
}
```

---

## 🟣 Flow Control

### Loops
```java
// Enhanced For-Loop
int[] nums = {1, 2, 3};
for (int n : nums) {
    System.out.println(n);
}

// Standard For-Loop
for (int i = 0; i < 5; i++) {
    System.out.println(i);
}

// While Loop
while (condition) { /*...*/ }
```

### Switch (Modern)
```java
String day = "MONDAY";

// Java 14+ specific syntax (if enabled) or standard:
switch (day) {
    case "MONDAY":
        System.out.println("Start of week");
        break;
    default:
        System.out.println("Other day");
}
```

---

## 🟠 Exception Handling

### Try-Catch-Finally
```java
try {
    int result = 10 / 0;
} catch (ArithmeticException e) {
    System.out.println("Cannot divide by zero: " + e.getMessage());
} finally {
    System.out.println("Always runs (cleanup)");
}
```

### Throw & Throws
```java
// Method declares it *might* throw an exception
public void readFile(String path) throws IOException {
    if (path == null) {
        // Actually throw exception
        throw new IOException("Path missing");
    }
}
```

---

## 🔴 Java 8+ Features

### Lambdas
Concise way to represent functional interfaces.
```java
List<String> names = Arrays.asList("John", "Jane");

// Sort using lambda
Collections.sort(names, (a, b) -> a.compareTo(b));

// ForEach
names.forEach(name -> System.out.println(name));
```

### Streams API
Functional-style operations on collections.
```java
List<Integer> numbers = Arrays.asList(1, 2, 3, 4, 5, 6);

List<Integer> evens = numbers.stream()
    .filter(n -> n % 2 == 0)       // Filter
    .map(n -> n * n)               // Map (Transform)
    .collect(Collectors.toList()); // Collect

System.out.println(evens); // [4, 16, 36]
```

### Optional
Avoid `NullPointerException`.
```java
Optional<String> opt = Optional.ofNullable(null);

// If present, print; else default
System.out.println(opt.orElse("Default Value"));

// Execute if present
opt.ifPresent(val -> System.out.println("Value: " + val));
```

---

## 🟤 Concurrency

### Threads
```java
// Extend Thread
class MyThread extends Thread {
    public void run() { System.out.println("Running"); }
}
new MyThread().start();

// Implement Runnable (Preferred)
Runnable task = () -> System.out.println("Task running");
new Thread(task).start();
```

### Executor Service (Thread Pool)
Modern way to manage threads.
```java
ExecutorService executor = Executors.newFixedThreadPool(2);

executor.submit(() -> System.out.println("Task 1"));
executor.submit(() -> System.out.println("Task 2"));

executor.shutdown();
```

---

## ⚡ Common Utilities

### StringBuilder
Efficient string manipulation (mutable).
```java
StringBuilder sb = new StringBuilder();
sb.append("Java");
sb.append(" ");
sb.append("Programming");
String result = sb.toString();
```

### Arrays Class
```java
int[] arr = {3, 1, 4};
Arrays.sort(arr);                     // Sort
String s = Arrays.toString(arr);      // Readable string "[1, 3, 4]"
List<Integer> list = Arrays.asList(1, 2, 3); // Convert to List
```
