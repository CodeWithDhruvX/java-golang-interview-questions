# Coding: Streams & Functional - Interview Answers

> 🎯 **Focus:** Interviewers love minimal code. Show off `filter`, `map`, `collect`.

### 1. Filter Even Numbers from a List
"I’ll use `stream()` followed by a `filter` predicate. Then I collect it back into a List."

```java
List<Integer> numbers = Arrays.asList(1, 2, 3, 4, 5, 6);

List<Integer> evens = numbers.stream()
                             .filter(n -> n % 2 == 0)
                             .collect(Collectors.toList());
```

---

### 2. Convert a List of Strings to Uppercase
"This is a classic mapping operation. I use `.map(String::toUpperCase)` to transform each element."

```java
List<String> names = Arrays.asList("alice", "bob", "charlie");

List<String> upperNames = names.stream()
                               .map(String::toUpperCase)
                               .collect(Collectors.toList());
```

---

### 3. Find the First Non-Repeated Character
"I need to count frequencies first. I'll use `LinkedHashMap` in the collector to preserve the insertion order, so I can find the *first* one easily."

```java
String input = "swiss";

Character result = input.chars() // IntStream
    .mapToObj(c -> (char) c)
    .collect(Collectors.groupingBy(Function.identity(), LinkedHashMap::new, Collectors.counting()))
    .entrySet()
    .stream()
    .filter(entry -> entry.getValue() == 1L)
    .map(Map.Entry::getKey)
    .findFirst()
    .orElse(null);
```

---

### 4. Sort a List of Employees by Salary
"I’ll use `Comparator.comparing`. It reads like English. I can also chain `.reversed()` if I want descending order."

```java
List<Employee> employees = getEmployees();

List<Employee> sorted = employees.stream()
    .sorted(Comparator.comparingDouble(Employee::getSalary).reversed())
    .collect(Collectors.toList());
```

---

### 5. Check if List contains any name starting with 'A'
"I should use short-circuiting logic here. `anyMatch` stops processing as soon as it finds one match, so it's efficient."

```java
boolean hasA = names.stream()
                    .anyMatch(name -> name.startsWith("A"));
```

---

### 6. Sum of all numbers in a List
"I can use `reduce` or simply `mapToInt().sum()`. The primitive stream `IntStream` is more efficient because it avoids boxing overhead."

```java
List<Integer> nums = Arrays.asList(1, 2, 3, 4, 5);

int sum = nums.stream()
              .mapToInt(Integer::intValue)
              .sum();
```

---

### 7. Join a List of Strings with a delimiter
"There’s a dedicated collector for this: `Collectors.joining`. It’s much cleaner than a loop."

```java
List<String> langs = Arrays.asList("Java", "Go", "Python");

String result = langs.stream()
                     .collect(Collectors.joining(", "));
// Output: "Java, Go, Python"
```

---

### 8. Find Max and Min Numbers
"I can use `IntSummaryStatistics` to get min, max, average, and count all in one pass. It’s very handy."

```java
List<Integer> numbers = Arrays.asList(5, 1, 9, 3, 7);

IntSummaryStatistics stats = numbers.stream()
                                    .mapToInt(x -> x)
                                    .summaryStatistics();

System.out.println("Max: " + stats.getMax());
System.out.println("Min: " + stats.getMin());
```
