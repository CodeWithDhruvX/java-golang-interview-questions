# Java Functional Programming & Streams — High-Speed Retrieval Cheatsheet

> **Interview-Ready Framework**: 5-Step Mental Model for Instant Recall

---

## 🔧 The 'Underlying Engine' Map

### **Category 1: Functional Interfaces Core**
**Logic Pattern**: Input → Process → Output
- **Predicate**: Boolean test (`test()`) → `filter()`, validation
- **Function**: Transform (`apply()`) → `map()`, data conversion  
- **Consumer**: Action (`accept()`) → `forEach()`, side effects
- **Supplier**: Generate (`get()`) → factory methods, lazy values
- **Binary variants**: BiFunction, BiPredicate → two-input operations

### **Category 2: Stream Pipeline Operations**
**Logic Pattern**: Data Flow → Transformation → Collection
- **Intermediate**: `filter()`, `map()`, `flatMap()`, `sorted()`, `distinct()`, `limit()`, `skip()`
- **Terminal**: `collect()`, `reduce()`, `forEach()`, `count()`, `min()`, `max()`
- **Short-circuiting**: `findFirst()`, `findAny()`, `anyMatch()`, `allMatch()`, `noneMatch()`

### **Category 3: Data Aggregation & Collection**
**Logic Pattern**: Group → Aggregate → Transform
- **Grouping**: `groupingBy()`, `partitioningBy()` → categorization
- **Numeric**: `summingInt()`, `averagingInt()`, `summarizingInt()` → statistics
- **Collection**: `toList()`, `toSet()`, `toMap()`, `joining()` → output formatting

### **Category 4: Optional & Error Handling**
**Logic Pattern**: Maybe Value → Safe Extraction → Fallback
- **Creation**: `of()`, `ofNullable()`, `empty()` → value containment
- **Transformation**: `map()`, `flatMap()`, `filter()` → value processing
- **Extraction**: `orElse()`, `orElseGet()`, `orElseThrow()` → safe access

---

## 🚨 The 'Red-Flag' Failure Section

### **Critical Runtime Errors**

| **Error Type** | **Trigger** | **Example** | **Fix** |
|---|---|---|---|
| **NullPointerException** | `Optional.of(null)` | `Optional.of(null)` | Use `Optional.ofNullable()` |
| **NoSuchElementException** | `Optional.get()` on empty | `empty.get()` | Use `orElse()`, `orElseGet()` |
| **IllegalStateException** | Stream reuse | `stream.forEach(); stream.forEach()` | Create new stream |
| **UnsupportedOperationException** | Modify immutable collection | `toUnmodifiableList().add()` | Use mutable collector |
| **NullPointerException** | Sorting with nulls | `list.stream().sorted()` with nulls | Use `nullsFirst()`, `nullsLast()` |

### **Logic Failure Patterns**

- **Infinite Streams**: `Stream.generate()`/`iterate()` without `limit()`
- **Empty Results**: `filter()` removing all elements → empty `Optional`
- **Performance Kill**: Small datasets + `parallel()` → overhead > benefit
- **Stateful Operations**: Shared mutable state in parallel streams → race conditions
- **Lazy Evaluation Trap**: `peek()` without terminal operation → no execution

---

## ⚡ The 'Performance & Complexity' Table

| **Operation** | **Time Complexity** | **Memory Usage** | **Best Use Case** |
|---|---|---|---|
| **filter()** | O(n) | O(1) | Removing unwanted elements |
| **map()** | O(n) | O(n) | Transforming each element |
| **flatMap()** | O(n+m) | O(n+m) | Flattening nested structures |
| **sorted()** | O(n log n) | O(n) | Ordering data |
| **distinct()** | O(n) | O(n) | Removing duplicates |
| **groupingBy()** | O(n) | O(n) | Categorization |
| **reduce()** | O(n) | O(1) | Aggregation |
| **parallel()** | O(n/p) | O(p) | Large CPU-bound tasks |

### **Performance Rules**
- **Primitive streams** (`IntStream`) vs `Stream<Integer>`: 3-5x faster, less GC
- **Sequential vs Parallel**: Use parallel only for >10,000 elements, CPU-bound
- **Lazy Evaluation**: Intermediate ops are free until terminal op
- **Short-circuiting**: `findFirst()`, `anyMatch()` stop early

---

## 🛡️ The 'Safe vs. Risky' Comparison

### **Standard/Safe Methods**
```java
// ✅ SAFE: Modern, recommended approaches
Optional.ofNullable(value).orElse(defaultValue);
stream.filter(predicate).collect(Collectors.toList());
IntStream.range(1, 100).sum();
Optional.empty().orElseGet(() -> expensiveDefault());
```

### **Legacy/Dangerous Methods**
```java
// ❌ RISKY: Avoid in modern code
Optional.get(); // Throws NoSuchElementException
stream.forEach(); // Without terminal op - no execution
Stream<Integer> for numbers; // Use IntStream instead
Optional.of(value); // Throws NPE if value is null
```

### **Why Use X Over Y?**

| **Safe Choice** | **Risky Alternative** | **Reason** |
|---|---|---|
| `orElseGet(() -> value)` | `orElse(value)` | Lazy evaluation for expensive defaults |
| `Optional.ofNullable()` | `Optional.of()` | Handles null values gracefully |
| `IntStream` | `Stream<Integer>` | Avoids boxing overhead |
| `forEachOrdered()` | `forEach()` (parallel) | Preserves order when needed |
| `Collectors.toUnmodifiableList()` | `Collections.unmodifiableList()` | Direct creation, more efficient |

---

## 🎯 The 'Interview Logic' Column

### **Core Concepts with Analogies & Golden Rules**

| **Concept** | **Real-World Analogy** | **Golden Rule** |
|---|---|---|
| **Lambda Expression** | **Recipe Card**: Instructions without a chef | "Lambdas are behavior as data - pass actions like you pass values" |
| **Stream Pipeline** | **Assembly Line**: Raw materials → processing → finished product | "Streams are lazy - nothing happens until you pull the lever (terminal operation)" |
| **Optional** | **Gift Box**: Might contain something, might be empty | "Never open the box without checking - use orElse, not get" |
| **Functional Interface** | **USB Port**: Standard shape for different devices | "One abstract method = infinite implementations through lambdas" |
| **Parallel Stream** | **Multiple Workers**: Same task, divided among people | "Parallel helps with big jobs, hurts with small ones" |
| **Collector** | **Storage Container**: Different shapes for different needs | "Choose the right container - List for order, Set for uniqueness, Map for key-value" |
| **Method Reference** | **Speed Dial**: Direct call to known number | "Use method references when lambda just calls an existing method" |
| **Lazy Evaluation** | **On-Demand Cooking**: Prepare food only when order comes | "Intermediate operations are free until terminal operation" |

### **Quick Interview Decision Tree**

1. **Need to test condition?** → `Predicate` → `filter()`
2. **Need to transform data?** → `Function` → `map()`
3. **Need to group data?** → `groupingBy()` + downstream collector
4. **Need to handle nulls?** → `Optional` + `ofNullable()`
5. **Need performance with numbers?** → Primitive streams (`IntStream`)
6. **Need to process large data?** → Consider `parallel()` with benchmarks

---

## 📚 Mental Index Cards for Rapid Recall

### **Card 1: Functional Interfaces Quick Reference**
```
Predicate<T>  → boolean test(T)           → filter()
Function<T,R> → R apply(T)                → map()
Consumer<T>   → void accept(T)            → forEach()
Supplier<T>   → T get()                   → generate()
```

### **Card 2: Stream Pipeline Pattern**
```
source.stream()
    .intermediate1()    // lazy
    .intermediate2()    // lazy  
    .terminal()         // executes pipeline
```

### **Card 3: Optional Safety Pattern**
```
Optional.ofNullable(value)
    .map(transform)
    .filter(validate)
    .orElseGet(defaultSupplier)
```

### **Card 4: Collector Decision Matrix**
```
Need order?     → toList()
Need uniqueness? → toSet()
Need key-value? → toMap()
Need groups?    → groupingBy()
Need string?    → joining()
```

---

## 🔥 Top 10 Interview Patterns

1. **Filter-Map-Collect**: `stream.filter().map().collect()`
2. **Group-Count**: `groupingBy(key, counting())`
3. **Optional Chain**: `ofNullable().map().orElseGet()`
4. **FlatMap Flatten**: `flatMap(Collection::stream)`
5. **Primitive Stream**: `mapToInt().sum()`
6. **Word Frequency**: `groupingBy(word, counting())`
7. **Partition Data**: `partitioningBy(predicate)`
8. **Custom Collector**: `Collector.of(supplier, accumulator, combiner, finisher)`
9. **Parallel Processing**: `parallel().forEachOrdered()`
10. **Multi-Group**: `groupingBy(key1, groupingBy(key2))`

---

**Interview Strategy**: When asked about streams, start with the "assembly line" analogy, then pick the appropriate pattern from this cheatsheet. Always mention lazy evaluation and choose primitive streams for performance.
