# Advanced Java Features: Practical Programs

**Goal**: Master modern Java 8+ features including Streams, Lambda, Optional, Date/Time API, and Generics.

## 1. Streams API & Lambda Expressions

### Example: Employee Data Processing

```java
import java.util.*;
import java.util.stream.*;
import java.util.function.*;

class Employee {
    private int id;
    private String name;
    private String department;
    private double salary;
    private int age;
    
    public Employee(int id, String name, String department, double salary, int age) {
        this.id = id;
        this.name = name;
        this.department = department;
        this.salary = salary;
        this.age = age;
    }
    
    // Getters
    public int getId() { return id; }
    public String getName() { return name; }
    public String getDepartment() { return department; }
    public double getSalary() { return salary; }
    public int getAge() { return age; }
    
    @Override
    public String toString() {
        return String.format("ID: %d, Name: %s, Dept: %s, Salary: %.2f, Age: %d", 
                           id, name, department, salary, age);
    }
}

public class StreamLambdaDemo {
    public static void main(String[] args) {
        List<Employee> employees = Arrays.asList(
            new Employee(1, "Alice", "IT", 75000, 28),
            new Employee(2, "Bob", "HR", 65000, 32),
            new Employee(3, "Charlie", "IT", 85000, 35),
            new Employee(4, "Diana", "Finance", 90000, 30),
            new Employee(5, "Eve", "IT", 70000, 26),
            new Employee(6, "Frank", "HR", 60000, 40)
        );
        
        // 1. Filter and Map with Lambda
        System.out.println("=== IT Employees with salary > 70000 ===");
        employees.stream()
                .filter(e -> e.getDepartment().equals("IT") && e.getSalary() > 70000)
                .map(e -> e.getName() + " ($" + e.getSalary() + ")")
                .forEach(System.out::println);
        
        // 2. Grouping and Counting
        System.out.println("\n=== Employees by Department ===");
        Map<String, Long> deptCount = employees.stream()
                .collect(Collectors.groupingBy(
                    Employee::getDepartment, 
                    Collectors.counting()
                ));
        deptCount.forEach((dept, count) -> 
            System.out.println(dept + ": " + count + " employees"));
        
        // 3. Average Salary by Department
        System.out.println("\n=== Average Salary by Department ===");
        Map<String, Double> avgSalary = employees.stream()
                .collect(Collectors.groupingBy(
                    Employee::getDepartment,
                    Collectors.averagingDouble(Employee::getSalary)
                ));
        avgSalary.forEach((dept, avg) -> 
            System.out.printf("%s: $%.2f\n", dept, avg));
        
        // 4. Custom Predicate and Function
        Predicate<Employee> isSenior = e -> e.getAge() > 30;
        Function<Employee, String> employeeInfo = e -> 
            String.format("%s (Age: %d, Dept: %s)", e.getName(), e.getAge(), e.getDepartment());
        
        System.out.println("\n=== Senior Employees Info ===");
        employees.stream()
                .filter(isSenior)
                .map(employeeInfo)
                .forEach(System.out::println);
        
        // 5. Parallel Stream for large datasets
        System.out.println("\n=== Total Payroll (Parallel Processing) ===");
        double totalPayroll = employees.parallelStream()
                .mapToDouble(Employee::getSalary)
                .sum();
        System.out.printf("Total Payroll: $%.2f\n", totalPayroll);
        
        // 6. Find and Optional operations
        System.out.println("\n=== Highest Paid Employee ===");
        Optional<Employee> highestPaid = employees.stream()
                .max(Comparator.comparingDouble(Employee::getSalary));
        highestPaid.ifPresent(emp -> System.out.println(emp.getName() + " ($" + emp.getSalary() + ")"));
        
        // 7. Custom Collector
        System.out.println("\n=== Employee Summary Statistics ===");
        IntSummaryStatistics ageStats = employees.stream()
                .mapToInt(Employee::getAge)
                .summaryStatistics();
        System.out.println("Age Statistics: " + ageStats);
    }
}
```

## 2. Functional Interfaces

### Custom Functional Interfaces

```java
@FunctionalInterface
interface MathOperation {
    double operate(double a, double b);
    
    // Default method
    default void printResult(double a, double b) {
        System.out.printf("%.2f %s %.2f = %.2f\n", a, getSymbol(), b, operate(a, b));
    }
    
    // Helper method
    private String getSymbol() {
        return "+"; // Default, can be overridden
    }
}

@FunctionalInterface
interface StringProcessor {
    String process(String input);
    
    // Static method
    static String reverse(String input) {
        return new StringBuilder(input).reverse().toString();
    }
}

@FunctionalInterface
interface Predicate<T> {
    boolean test(T t);
    
    // Default methods for chaining
    default Predicate<T> and(Predicate<T> other) {
        return t -> this.test(t) && other.test(t);
    }
    
    default Predicate<T> or(Predicate<T> other) {
        return t -> this.test(t) || other.test(t);
    }
}

public class FunctionalInterfaceDemo {
    public static void main(String[] args) {
        // Lambda expressions for different operations
        MathOperation addition = (a, b) -> a + b;
        MathOperation multiplication = (a, b) -> a * b;
        MathOperation division = (a, b) -> {
            if (b == 0) throw new IllegalArgumentException("Division by zero");
            return a / b;
        };
        
        System.out.println("=== Math Operations ===");
        addition.printResult(10, 5);
        multiplication.printResult(10, 5);
        division.printResult(10, 5);
        
        // Method references
        StringProcessor toUpperCase = String::toUpperCase;
        StringProcessor toLowerCase = String::toLowerCase;
        StringProcessor reverseString = StringProcessor::reverse;
        
        System.out.println("\n=== String Processing ===");
        String text = "Hello World";
        System.out.println("Original: " + text);
        System.out.println("Upper: " + toUpperCase.process(text));
        System.out.println("Lower: " + toLowerCase.process(text));
        System.out.println("Reverse: " + reverseString.process(text));
        
        // Custom Predicate
        Predicate<String> isLong = s -> s.length() > 5;
        Predicate<String> containsJava = s -> s.toLowerCase().contains("java");
        
        System.out.println("\n=== Predicate Chaining ===");
        String[] testStrings = {"Java", "JavaScript", "Python", "C++"};
        for (String str : testStrings) {
            boolean result = isLong.and(containsJava).test(str);
            System.out.printf("%s: Long and contains Java? %b\n", str, result);
        }
        
        // BiFunction example
        BiFunction<String, Integer, String> repeatString = (s, n) -> s.repeat(n);
        System.out.println("\n=== BiFunction Example ===");
        System.out.println(repeatString.apply("Hi! ", 3));
        
        // Consumer example
        Consumer<String> printer = str -> System.out.println("Printing: " + str);
        Consumer<String> logger = str -> System.err.println("Log: " + str);
        
        System.out.println("\n=== Consumer Chaining ===");
        Consumer<String> printAndLog = printer.andThen(logger);
        printAndLog.accept("Test Message");
    }
}
```

## 3. Optional Class

### Comprehensive Optional Usage

```java
import java.util.*;
import java.util.concurrent.ThreadLocalRandom;

class User {
    private String name;
    private String email;
    private Optional<String> phone;
    private Optional<Address> address;
    
    public User(String name, String email, String phone, Address address) {
        this.name = name;
        this.email = email;
        this.phone = Optional.ofNullable(phone);
        this.address = Optional.ofNullable(address);
    }
    
    public String getName() { return name; }
    public String getEmail() { return email; }
    public Optional<String> getPhone() { return phone; }
    public Optional<Address> getAddress() { return address; }
}

class Address {
    private String street;
    private String city;
    private Optional<String> zipCode;
    
    public Address(String street, String city, String zipCode) {
        this.street = street;
        this.city = city;
        this.zipCode = Optional.ofNullable(zipCode);
    }
    
    public String getStreet() { return street; }
    public String getCity() { return city; }
    public Optional<String> getZipCode() { return zipCode; }
}

public class OptionalDemo {
    public static void main(String[] args) {
        // Create test data
        Address addr1 = new Address("123 Main St", "New York", "10001");
        Address addr2 = new Address("456 Oak Ave", "Los Angeles", null);
        
        List<User> users = Arrays.asList(
            new User("Alice", "alice@email.com", "555-1234", addr1),
            new User("Bob", "bob@email.com", null, addr2),
            new User("Charlie", "charlie@email.com", "555-5678", null)
        );
        
        // 1. Basic Optional operations
        System.out.println("=== Basic Optional Operations ===");
        for (User user : users) {
            System.out.println("\nUser: " + user.getName());
            
            // ifPresent
            user.getPhone().ifPresent(phone -> 
                System.out.println("Phone: " + phone));
            
            // orElse
            String phoneDisplay = user.getPhone().orElse("Not provided");
            System.out.println("Phone (orElse): " + phoneDisplay);
            
            // orElseGet
            String phoneLazy = user.getPhone().orElseGet(() -> 
                "No phone available for " + user.getName());
            System.out.println("Phone (orElseGet): " + phoneLazy);
        }
        
        // 2. Nested Optional operations
        System.out.println("\n=== Nested Optional Operations ===");
        for (User user : users) {
            String city = user.getAddress()
                    .map(Address::getCity)
                    .orElse("No address");
            System.out.println(user.getName() + " lives in: " + city);
            
            // Nested mapping
            String zipCode = user.getAddress()
                    .flatMap(Address::getZipCode)
                    .orElse("No ZIP code");
            System.out.println(user.getName() + " ZIP: " + zipCode);
        }
        
        // 3. Optional with streams
        System.out.println("\n=== Optional with Streams ===");
        List<String> phoneNumbers = users.stream()
                .map(User::getPhone)
                .filter(Optional::isPresent)
                .map(Optional::get)
                .toList();
        System.out.println("All phone numbers: " + phoneNumbers);
        
        // 4. Optional in method return
        System.out.println("\n=== Method Returning Optional ===");
        Optional<User> foundUser = findUserByEmail(users, "bob@email.com");
        foundUser.ifPresent(user -> 
            System.out.println("Found user: " + user.getName()));
        
        // 5. Optional for exception handling
        System.out.println("\n=== Optional for Exception Handling ===");
        Optional<Integer> result = safeDivide(10, 2);
        result.ifPresent(res -> System.out.println("10 / 2 = " + res));
        
        Optional<Integer> invalidResult = safeDivide(10, 0);
        System.out.println("10 / 0 = " + invalidResult.orElse(-1));
        
        // 6. Custom Optional creation
        System.out.println("\n=== Custom Optional Creation ===");
        Optional<String> empty = Optional.empty();
        Optional<String> nonNull = Optional.of("Hello");
        Optional<String> nullable = Optional.ofNullable(null);
        
        System.out.println("Empty present: " + empty.isPresent());
        System.out.println("Non-null present: " + nonNull.isPresent());
        System.out.println("Nullable present: " + nullable.isPresent());
    }
    
    private static Optional<User> findUserByEmail(List<User> users, String email) {
        return users.stream()
                .filter(user -> user.getEmail().equals(email))
                .findFirst();
    }
    
    private static Optional<Integer> safeDivide(int a, int b) {
        try {
            return Optional.of(a / b);
        } catch (ArithmeticException e) {
            return Optional.empty();
        }
    }
}
```

## 4. Date/Time API (Java 8+)

### Modern Date/Time Operations

```java
import java.time.*;
import java.time.format.*;
import java.time.temporal.*;
import java.util.*;

public class DateTimeDemo {
    public static void main(String[] args) {
        // 1. Local Date, Time, and DateTime
        System.out.println("=== Basic Date/Time ===");
        LocalDate today = LocalDate.now();
        LocalTime now = LocalTime.now();
        LocalDateTime current = LocalDateTime.now();
        
        System.out.println("Today: " + today);
        System.out.println("Now: " + now);
        System.out.println("Current DateTime: " + current);
        
        // 2. Creating specific dates
        LocalDate birthday = LocalDate.of(1990, Month.JANUARY, 15);
        LocalTime meetingTime = LocalTime.of(14, 30);
        LocalDateTime appointment = LocalDateTime.of(2024, 12, 25, 10, 0);
        
        System.out.println("\nBirthday: " + birthday);
        System.out.println("Meeting time: " + meetingTime);
        System.out.println("Appointment: " + appointment);
        
        // 3. Date manipulation
        System.out.println("\n=== Date Manipulation ===");
        LocalDate nextWeek = today.plusWeeks(1);
        LocalDate lastMonth = today.minusMonths(1);
        LocalDate nextYearSameDay = today.plusYears(1);
        
        System.out.println("Today: " + today);
        System.out.println("Next week: " + nextWeek);
        System.out.println("Last month: " + lastMonth);
        System.out.println("Next year same day: " + nextYearSameDay);
        
        // 4. Date calculations
        System.out.println("\n=== Date Calculations ===");
        Period age = Period.between(birthday, today);
        System.out.printf("Age: %d years, %d months, %d days\n", 
                         age.getYears(), age.getMonths(), age.getDays());
        
        // Days until Christmas
        LocalDate christmas = LocalDate.of(today.getYear(), Month.DECEMBER, 25);
        long daysUntilChristmas = ChronoUnit.DAYS.between(today, christmas);
        System.out.println("Days until Christmas: " + daysUntilChristmas);
        
        // 5. Working with ZonedDateTime
        System.out.println("\n=== Time Zones ===");
        ZonedDateTime utc = ZonedDateTime.now(ZoneId.of("UTC"));
        ZonedDateTime newYork = ZonedDateTime.now(ZoneId.of("America/New_York"));
        ZonedDateTime tokyo = ZonedDateTime.now(ZoneId.of("Asia/Tokyo"));
        
        System.out.println("UTC: " + utc);
        System.out.println("New York: " + newYork);
        System.out.println("Tokyo: " + tokyo);
        
        // 6. Duration and Instant
        System.out.println("\n=== Duration and Instant ===");
        Instant start = Instant.now();
        
        // Simulate some work
        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        
        Instant end = Instant.now();
        Duration duration = Duration.between(start, end);
        System.out.println("Duration: " + duration.toMillis() + " milliseconds");
        
        // 7. Date formatting and parsing
        System.out.println("\n=== Date Formatting ===");
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
        DateTimeFormatter usFormatter = DateTimeFormatter.ofPattern("MM/dd/yyyy");
        DateTimeFormatter customFormatter = DateTimeFormatter.ofPattern("EEEE, MMMM d, yyyy");
        
        String formatted = current.format(formatter);
        String usFormatted = today.format(usFormatter);
        String customFormatted = today.format(customFormatter);
        
        System.out.println("Standard format: " + formatted);
        System.out.println("US format: " + usFormatted);
        System.out.println("Custom format: " + customFormatted);
        
        // Parsing dates
        LocalDate parsedDate = LocalDate.parse("2024-12-25");
        LocalDateTime parsedDateTime = LocalDateTime.parse("2024-12-25 10:00:00", formatter);
        
        System.out.println("Parsed date: " + parsedDate);
        System.out.println("Parsed datetime: " + parsedDateTime);
        
        // 8. Working with calendars and business logic
        System.out.println("\n=== Business Date Logic ===");
        LocalDate projectStart = LocalDate.of(2024, 1, 1);
        LocalDate projectEnd = projectStart.plusMonths(3);
        
        // Calculate working days (excluding weekends)
        long workingDays = calculateWorkingDays(projectStart, projectEnd);
        System.out.println("Working days between " + projectStart + " and " + projectEnd + ": " + workingDays);
        
        // Find next business day
        LocalDate nextBusinessDay = findNextBusinessDay(today);
        System.out.println("Next business day: " + nextBusinessDay);
    }
    
    private static long calculateWorkingDays(LocalDate start, LocalDate end) {
        return start.datesUntil(end.plusDays(1))
                   .filter(date -> date.getDayOfWeek() != DayOfWeek.SATURDAY 
                              && date.getDayOfWeek() != DayOfWeek.SUNDAY)
                   .count();
    }
    
    private static LocalDate findNextBusinessDay(LocalDate date) {
        LocalDate next = date.plusDays(1);
        while (next.getDayOfWeek() == DayOfWeek.SATURDAY || 
               next.getDayOfWeek() == DayOfWeek.SUNDAY) {
            next = next.plusDays(1);
        }
        return next;
    }
}
```

## 5. Generic Types and Wildcards

### Comprehensive Generics Usage

```java
import java.util.*;

// Generic Box class
class Box<T> {
    private T content;
    
    public void setContent(T content) {
        this.content = content;
    }
    
    public T getContent() {
        return content;
    }
    
    @Override
    public String toString() {
        return "Box contains: " + content;
    }
}

// Generic Pair class
class Pair<K, V> {
    private K key;
    private V value;
    
    public Pair(K key, V value) {
        this.key = key;
        this.value = value;
    }
    
    public K getKey() { return key; }
    public V getValue() { return value; }
    
    @Override
    public String toString() {
        return key + " = " + value;
    }
}

// Generic utility class
class CollectionUtils {
    // Generic method with bounded type parameter
    public static <T extends Comparable<T>> T max(List<T> list) {
        if (list.isEmpty()) {
            throw new IllegalArgumentException("List is empty");
        }
        
        T max = list.get(0);
        for (T item : list) {
            if (item.compareTo(max) > 0) {
                max = item;
            }
        }
        return max;
    }
    
    // Method with wildcard
    public static void printList(List<?> list) {
        for (Object item : list) {
            System.out.print(item + " ");
        }
        System.out.println();
    }
    
    // Upper bounded wildcard
    public static double sumOfNumbers(List<? extends Number> numbers) {
        double sum = 0.0;
        for (Number num : numbers) {
            sum += num.doubleValue();
        }
        return sum;
    }
    
    // Lower bounded wildcard
    public static void addNumbers(List<? super Integer> list) {
        for (int i = 1; i <= 5; i++) {
            list.add(i);
        }
    }
    
    // Generic method with multiple type parameters
    public static <T, U> void processPairs(List<Pair<T, U>> pairs) {
        for (Pair<T, U> pair : pairs) {
            System.out.println("Processing: " + pair);
        }
    }
}

// Generic interface
interface Repository<T> {
    void save(T entity);
    T findById(int id);
    List<T> findAll();
    void delete(int id);
}

// Generic implementation
class InMemoryRepository<T> implements Repository<T> {
    private Map<Integer, T> storage = new HashMap<>();
    private int nextId = 1;
    
    @Override
    public void save(T entity) {
        storage.put(nextId++, entity);
    }
    
    @Override
    public T findById(int id) {
        return storage.get(id);
    }
    
    @Override
    public List<T> findAll() {
        return new ArrayList<>(storage.values());
    }
    
    @Override
    public void delete(int id) {
        storage.remove(id);
    }
}

public class GenericsDemo {
    public static void main(String[] args) {
        // 1. Basic generic class usage
        System.out.println("=== Generic Box Usage ===");
        Box<String> stringBox = new Box<>();
        stringBox.setContent("Hello Generics");
        System.out.println(stringBox);
        
        Box<Integer> integerBox = new Box<>();
        integerBox.setContent(42);
        System.out.println(integerBox);
        
        Box<List<String>> listBox = new Box<>();
        listBox.setContent(Arrays.asList("A", "B", "C"));
        System.out.println(listBox);
        
        // 2. Generic Pair usage
        System.out.println("\n=== Generic Pair Usage ===");
        Pair<String, Integer> studentScore = new Pair<>("Alice", 95);
        System.out.println(studentScore);
        
        Pair<String, String> countryCapital = new Pair<>("France", "Paris");
        System.out.println(countryCapital);
        
        // 3. Generic utility methods
        System.out.println("\n=== Generic Utility Methods ===");
        List<Integer> numbers = Arrays.asList(3, 1, 4, 1, 5, 9, 2, 6);
        System.out.println("Numbers: " + numbers);
        System.out.println("Max number: " + CollectionUtils.max(numbers));
        
        List<String> words = Arrays.asList("Apple", "Banana", "Cherry", "Date");
        System.out.println("Words: " + words);
        System.out.println("Max word: " + CollectionUtils.max(words));
        
        // 4. Wildcard usage
        System.out.println("\n=== Wildcard Usage ===");
        List<String> stringList = Arrays.asList("Hello", "World", "Java");
        List<Integer> intList = Arrays.asList(1, 2, 3, 4, 5);
        
        System.out.print("String list: ");
        CollectionUtils.printList(stringList);
        
        System.out.print("Integer list: ");
        CollectionUtils.printList(intList);
        
        // 5. Upper bounded wildcard
        System.out.println("\n=== Upper Bounded Wildcard ===");
        List<Integer> ints = Arrays.asList(1, 2, 3, 4, 5);
        List<Double> doubles = Arrays.asList(1.1, 2.2, 3.3, 4.4, 5.5);
        List<Long> longs = Arrays.asList(1L, 2L, 3L, 4L, 5L);
        
        System.out.println("Sum of integers: " + CollectionUtils.sumOfNumbers(ints));
        System.out.println("Sum of doubles: " + CollectionUtils.sumOfNumbers(doubles));
        System.out.println("Sum of longs: " + CollectionUtils.sumOfNumbers(longs));
        
        // 6. Lower bounded wildcard
        System.out.println("\n=== Lower Bounded Wildcard ===");
        List<Number> numberList = new ArrayList<>();
        CollectionUtils.addNumbers(numberList);
        System.out.println("Number list after adding integers: " + numberList);
        
        List<Object> objectList = new ArrayList<>();
        CollectionUtils.addNumbers(objectList);
        System.out.println("Object list after adding integers: " + objectList);
        
        // 7. Generic repository
        System.out.println("\n=== Generic Repository ===");
        Repository<String> stringRepo = new InMemoryRepository<>();
        stringRepo.save("First Item");
        stringRepo.save("Second Item");
        stringRepo.save("Third Item");
        
        System.out.println("All strings: " + stringRepo.findAll());
        System.out.println("Item at ID 2: " + stringRepo.findById(2));
        
        Repository<Double> doubleRepo = new InMemoryRepository<>();
        doubleRepo.save(3.14);
        doubleRepo.save(2.71);
        doubleRepo.save(1.61);
        
        System.out.println("All doubles: " + doubleRepo.findAll());
        
        // 8. Generic method with multiple parameters
        System.out.println("\n=== Multiple Type Parameters ===");
        List<Pair<String, Integer>> pairs = Arrays.asList(
            new Pair<>("Alice", 25),
            new Pair<>("Bob", 30),
            new Pair<>("Charlie", 35)
        );
        CollectionUtils.processPairs(pairs);
        
        // 9. Type inference with diamond operator
        System.out.println("\n=== Type Inference ===");
        Map<String, List<Integer>> map = new HashMap<>(); // Diamond operator
        
        map.put("even", Arrays.asList(2, 4, 6, 8));
        map.put("odd", Arrays.asList(1, 3, 5, 7));
        
        map.forEach((category, numbers) -> {
            System.out.println(category + " numbers: " + numbers);
        });
    }
}
```

## Practice Exercises

1. **Streams & Lambda**: Create a product management system with filtering, sorting, and aggregation
2. **Functional Interfaces**: Implement a custom event handling system using functional interfaces
3. **Optional**: Refactor existing code to use Optional instead of null checks
4. **Date/Time**: Build a scheduling application with recurring events and time zone support
5. **Generics**: Create a generic caching system that can store any type of data

## Interview Questions

1. What's the difference between `map()` and `flatMap()` in streams?
2. When would you use `Optional` over `null`?
3. What are the advantages of lambda expressions over anonymous inner classes?
4. How does the new Date/Time API improve upon the old `java.util.Date`?
5. What's the difference between `? extends T` and `? super T` wildcards?
