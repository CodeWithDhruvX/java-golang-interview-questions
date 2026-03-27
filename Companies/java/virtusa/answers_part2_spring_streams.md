# Virtusa Java Interview – Part 2: Spring Boot & Java Streams (Q&A with Code)

---

## Q1. What is the difference between Spring Framework and Spring Boot?

**Spoken Answer:**
"Spring Framework is the core container-based framework — it gives you dependency injection, AOP, data access abstractions. But configuring a Spring app is verbose: you define DispatcherServlet in `web.xml`, configure every bean, manage DataSource yourself.

Spring Boot is **opinionated setup built on top of Spring Framework**. It provides:
- **Auto-configuration** — detects what's on your classpath and configures it automatically
- **Embedded server** — Tomcat is bundled, no WAR deployment needed
- **Starter POMs** — `spring-boot-starter-web` brings in everything needed for a REST API
- **`application.properties`** — one file for all environment config

The key phrase I use: Spring Boot is **convention over configuration**."

```java
// Spring Boot: this is all you need to launch a web server!
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

@SpringBootApplication  // = @Configuration + @EnableAutoConfiguration + @ComponentScan
public class VirtusaApp {
    public static void main(String[] args) {
        SpringApplication.run(VirtusaApp.class, args);
        // Embedded Tomcat starts automatically on port 8080 — no web.xml, no WAR!
    }
}

@RestController
@RequestMapping("/api")
class HelloController {

    @GetMapping("/hello")
    public String hello(@RequestParam(defaultValue = "World") String name) {
        return "Hello, " + name + "!";
    }
}

// application.properties (all you need):
// server.port=8080
// spring.datasource.url=jdbc:h2:mem:testdb   ← auto-configures H2 DataSource!
// spring.jpa.hibernate.ddl-auto=update
```

---

## Q2. What is the difference between @Component, @Service, and @Repository?

**Spoken Answer:**
"All three are specializations of `@Component` — they all register the class as a Spring-managed bean. The difference is **semantic role and, for Repository, extra behavior**:

- `@Component` — generic bean, no specific layer implied
- `@Service` — marks business logic layer. No extra framework behavior but communicates intent clearly.
- `@Repository` — marks the data access layer. Spring adds **exception translation** — raw JDBC/JPA exceptions are converted into Spring's `DataAccessException` hierarchy automatically.

I always use the right annotation for the right layer — it improves readability and lets Spring provide the correct proxy behavior."

```java
import org.springframework.stereotype.*;
import org.springframework.beans.factory.annotation.*;
import java.util.Optional;

// Generic utility — no specific layer (e.g., email validation helper)
@Component
class EmailValidator {
    public boolean isValid(String email) {
        return email != null && email.contains("@");
    }
}

// Business logic layer
@Service
class EmployeeService {

    @Autowired
    private EmployeeRepository repo;

    @Autowired
    private EmailValidator validator;

    public Employee getEmployee(Long id) {
        return repo.findById(id)
               .orElseThrow(() -> new RuntimeException("Employee not found: " + id));
    }

    public void hireEmployee(Employee emp) {
        if (!validator.isValid(emp.email())) throw new IllegalArgumentException("Bad email");
        repo.save(emp);
    }
}

// Data access layer — gets Spring exception translation
@Repository
class EmployeeRepository {
    // If a DB constraint violation occurs, Spring translates:
    //   SQLIntegrityConstraintViolationException → DataIntegrityViolationException
    // so the service layer doesn't need to know about raw JDBC exceptions

    public Optional<Employee> findById(Long id) {
        // JPA / JDBC logic
        return Optional.empty();
    }
    public void save(Employee e) { /* persist */ }
}

record Employee(Long id, String name, String email) {}
```

---

## Q3. What is the difference between @Controller and @RestController?

**Spoken Answer:**
"`@Controller` is for Spring MVC — it returns a **view name** (like a Thymeleaf template). Spring resolves that name to an HTML page using a ViewResolver.

`@RestController` is `@Controller + @ResponseBody` combined. Every method automatically serializes the return value to **JSON or XML** and writes it directly to the HTTP response body. This is what I always use for REST APIs.

If I need to return both a page and JSON from the same controller, I can use `@Controller` with `@ResponseBody` on specific methods."

```java
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.ui.Model;
import java.util.*;

// ── For web pages (MVC) ──
@Controller
class WebController {

    @GetMapping("/home")
    public String homePage(Model model) {
        model.addAttribute("username", "Alice");
        return "home"; // → resolves to templates/home.html (Thymeleaf/JSP)
    }

    // Need JSON from a @Controller? Add @ResponseBody explicitly
    @GetMapping("/data")
    @ResponseBody
    public Map<String, String> getData() {
        return Map.of("key", "value"); // serialized to JSON
    }
}

// ── For REST APIs ──
@RestController  // = @Controller + @ResponseBody on all methods
@RequestMapping("/api/v1/employees")
class EmployeeApiController {

    // Returned object auto-serialized → JSON
    @GetMapping
    public List<Map<String, Object>> list() {
        return List.of(
            Map.of("id", 1, "name", "Alice", "salary", 70000),
            Map.of("id", 2, "name", "Bob",   "salary", 55000)
        );
    }

    // Request body auto-deserialized from JSON
    @PostMapping
    public Map<String, String> create(@RequestBody Map<String, String> body) {
        return Map.of("status", "created", "name", body.get("name"));
    }
}
```

---

## Q4. What is the difference between @RequestParam and @PathVariable?

**Spoken Answer:**
"`@PathVariable` extracts a value **from the URL path itself**: `GET /employees/42` — the 42 is part of the path structure. Use it when the value **identifies a resource**.

`@RequestParam` extracts from **query parameters**: `GET /employees?dept=IT&page=2` — key-value pairs after the `?`. Use it for **optional filters, pagination, search criteria**."

```java
@RestController
@RequestMapping("/employees")
class EmployeeController {

    // @PathVariable — embedded in URL structure
    // GET /employees/42
    @GetMapping("/{id}")
    public String getById(@PathVariable Long id) {
        return "Employee ID: " + id;
    }

    // GET /employees/dept/IT
    @GetMapping("/dept/{deptName}")
    public String getByDept(@PathVariable String deptName) {
        return "Department: " + deptName;
    }

    // @RequestParam — query string
    // GET /employees?dept=IT&minSalary=50000&page=0&size=10
    @GetMapping
    public String search(
            @RequestParam(required = false) String dept,
            @RequestParam(required = false) Double minSalary,
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "10") int size) {
        return "dept=" + dept + ", minSalary=" + minSalary + ", page=" + page;
    }

    // Combining both on one endpoint
    // GET /employees/42/projects?active=true
    @GetMapping("/{id}/projects")
    public String getProjects(
            @PathVariable Long id,
            @RequestParam(defaultValue = "true") boolean active) {
        return "Projects for employee " + id + ", active=" + active;
    }
}
```

---

## Q5. Java Streams coding — Print employee names with salary > 55,000. Use Streams to filter records with age > 30 AND salary > 1 lakh and concatenate their names.

**Spoken Answer:**
"Streams give us a declarative, pipeline-based way to process collections. I chain: `filter()` to select, `map()` to transform, `collect()` or `forEach()` as a terminal operation. Everything is lazy — nothing runs until the terminal operation is called."

```java
import java.util.*;
import java.util.stream.*;

public class StreamsDemo {

    record Employee(String name, String dept, int age, double salary) {}

    public static void main(String[] args) {

        List<Employee> employees = List.of(
            new Employee("Alice",   "IT",      32, 120000),
            new Employee("Bob",     "HR",       28,  45000),
            new Employee("Charlie", "IT",       35,  75000),
            new Employee("Diana",   "Finance",  40, 110000),
            new Employee("Eve",     "IT",       25,  80000),
            new Employee("Frank",   "IT",       33,  56000)
        );

        // ─── Task 1: Print names with salary > 55,000 ───
        System.out.println("── Salary > 55,000 ──");
        employees.stream()
                 .filter(e -> e.salary() > 55_000)
                 .map(Employee::name)
                 .forEach(System.out::println); // Alice, Charlie, Diana, Eve, Frank

        // ─── Task 2: age > 30 AND salary > 1 lakh, concat names ───
        System.out.println("\n── age > 30 AND salary > 1L ──");
        String result = employees.stream()
                .filter(e -> e.age() > 30 && e.salary() > 100_000)
                .map(Employee::name)
                .collect(Collectors.joining(", "));
        System.out.println(result); // Alice, Diana

        // ─── Bonus: Convert List to Map (id/name → Employee) ───
        Map<String, Double> salaryMap = employees.stream()
                .collect(Collectors.toMap(Employee::name, Employee::salary));
        System.out.println("\nSalary map: " + salaryMap);

        // ─── Bonus: Group by department ───
        Map<String, List<Employee>> byDept = employees.stream()
                .collect(Collectors.groupingBy(Employee::dept));
        byDept.forEach((dept, emps) ->
            System.out.println(dept + " → " + emps.stream().map(Employee::name).toList()));

        // ─── Bonus: Statistics ───
        DoubleSummaryStatistics stats = employees.stream()
                .mapToDouble(Employee::salary).summaryStatistics();
        System.out.printf("%nAvg: %.0f | Max: %.0f | Min: %.0f%n",
                stats.getAverage(), stats.getMax(), stats.getMin());
    }
}
```

---

## Q6. Write a program to convert a List into a Map in Java.

**Spoken Answer:**
"I use `Collectors.toMap()` from the Streams API. I pass a key mapper lambda and a value mapper lambda. If there's a chance of duplicate keys I also pass a merge function — otherwise it throws `IllegalStateException` on collision."

```java
import java.util.*;
import java.util.stream.*;

public class ListToMap {

    record Person(int id, String name, String city) {}

    public static void main(String[] args) {

        List<Person> people = List.of(
            new Person(1, "Alice", "Mumbai"),
            new Person(2, "Bob",   "Delhi"),
            new Person(3, "Charlie", "Bangalore")
        );

        // id → name
        Map<Integer, String> idToName = people.stream()
                .collect(Collectors.toMap(Person::id, Person::name));
        System.out.println("id→name: " + idToName);

        // id → full Person object
        Map<Integer, Person> idToPerson = people.stream()
                .collect(Collectors.toMap(Person::id, p -> p));
        System.out.println("id→person: " + idToPerson);

        // city → list of names (grouping)
        Map<String, List<String>> cityToNames = people.stream()
                .collect(Collectors.groupingBy(
                        Person::city,
                        Collectors.mapping(Person::name, Collectors.toList())));
        System.out.println("city→names: " + cityToNames);

        // Handling duplicate keys (keep first)
        List<Person> withDup = List.of(
            new Person(1, "Alice",  "Mumbai"),
            new Person(1, "Alice2", "Pune")   // duplicate id=1
        );
        Map<Integer, String> safeMap = withDup.stream()
                .collect(Collectors.toMap(
                        Person::id, Person::name,
                        (existing, newer) -> existing)); // keep existing on collision
        System.out.println("dedup: " + safeMap);
    }
}
```

---

## Q7. What is the difference between Lambda expressions and Functional Interfaces? Can we handle exceptions inside a Lambda expression?

**Spoken Answer:**
"A **Functional Interface** is an interface with exactly one abstract method — `@FunctionalInterface` annotation makes this explicit. Examples: `Runnable`, `Comparator`, `Predicate`, `Function`, `Consumer`, `Supplier`.

A **Lambda expression** is the concrete implementation of that one abstract method — a concise syntax that replaces anonymous inner classes.

Can we handle exceptions in lambdas? Yes, inline with try-catch. But checked exceptions are tricky — you can't declare `throws` in a lambda body targeting a `Function<T,R>`. The workaround is to wrap the checked call, or create a custom functional interface that declares the checked exception."

```java
import java.util.*;
import java.util.function.*;

public class LambdaFunctionalDemo {

    // Custom Functional Interface (with checked exception)
    @FunctionalInterface
    interface CheckedFunction<T, R> {
        R apply(T t) throws Exception;
    }

    // Wrapper to convert CheckedFunction → Function (rethrows as unchecked)
    static <T, R> Function<T, R> wrap(CheckedFunction<T, R> fn) {
        return t -> {
            try { return fn.apply(t); }
            catch (Exception e) { throw new RuntimeException(e); }
        };
    }

    public static void main(String[] args) {

        // ─── Standard Functional Interfaces ───
        Predicate<String> isLong     = s -> s.length() > 5;
        Function<String, Integer> len = String::length;   // method reference
        Consumer<String> print        = System.out::println;
        Supplier<List<String>> listOf  = ArrayList::new;

        System.out.println(isLong.test("HelloWorld")); // true
        System.out.println(len.apply("Java"));          // 4
        print.accept("Lambda Consumer");
        listOf.get().add("item");

        // ─── Without lambda (verbose) ───
        Comparator<String> old = new Comparator<String>() {
            @Override public int compare(String a, String b) { return a.compareTo(b); }
        };
        // With lambda:
        Comparator<String> modern = (a, b) -> a.compareTo(b);
        // Even shorter:
        Comparator<String> ref = String::compareTo;

        // ─── Exception handling inside lambda ───
        List<String> items = List.of("1", "abc", "3", "xyz");

        // Inline try-catch for RuntimeException
        items.forEach(s -> {
            try {
                System.out.println("Parsed: " + Integer.parseInt(s));
            } catch (NumberFormatException e) {
                System.out.println("Skipped non-numeric: " + s);
            }
        });

        // Checked exception via wrapper
        List<String> paths = List.of("file1.txt", "file2.txt");
        paths.stream()
             .map(wrap(path -> java.nio.file.Files.readString(java.nio.file.Path.of(path))))
             .forEach(System.out::println);
    }
}
```

---

## Q8. What are the SOLID Principles?

**Spoken Answer:**
"SOLID is five OOP design principles:
- **S — Single Responsibility**: A class should have only one reason to change.
- **O — Open/Closed**: Open for extension, closed for modification. Use interfaces.
- **L — Liskov Substitution**: A subclass must be replaceable for its parent without breaking behavior.
- **I — Interface Segregation**: Don't force a class to implement methods it doesn't use. Keep interfaces small.
- **D — Dependency Inversion**: High-level modules should depend on **abstractions**, not concrete implementations."

```java
// ─── D: Dependency Inversion (most commonly tested) ───
interface NotificationService {
    void send(String message, String recipient);
}

class EmailService implements NotificationService {
    public void send(String msg, String to) {
        System.out.println("Email → " + to + ": " + msg);
    }
}

class SmsService implements NotificationService {
    public void send(String msg, String to) {
        System.out.println("SMS → " + to + ": " + msg);
    }
}

// OrderService depends on abstraction — NOT on EmailService or SmsService
class OrderService {
    private final NotificationService notifier; // ← depends on interface

    OrderService(NotificationService notifier) { this.notifier = notifier; }

    void placeOrder(String item, String customer) {
        System.out.println("Order placed: " + item);
        notifier.send("Your order for " + item + " is confirmed!", customer);
    }
}

// ─── I: Interface Segregation ───
interface Printable  { void print(); }
interface Scannable  { void scan(); }
interface Faxable    { void fax(); }

// A basic printer only implements what it supports — not forced to implement fax
class BasicPrinter implements Printable {
    public void print() { System.out.println("Printing..."); }
}

class MultiFunctionDevice implements Printable, Scannable, Faxable {
    public void print() { System.out.println("Printing..."); }
    public void scan()  { System.out.println("Scanning..."); }
    public void fax()   { System.out.println("Faxing..."); }
}

class SOLIDDemo {
    public static void main(String[] args) {
        new OrderService(new EmailService()).placeOrder("Laptop", "alice@mail.com");
        new OrderService(new SmsService()).placeOrder("Phone",  "+91-99999");
        // OrderService unchanged — just swap the injected implementation
    }
}
```

---

## Q9. How do you handle global exceptions in Spring Boot?

**Spoken Answer:**
"Instead of surrounding every controller method with try-catch, I create a `@RestControllerAdvice` class with `@ExceptionHandler` methods. This acts as a central exception interceptor for all controllers. I return a consistent ErrorResponse JSON shape so the client always knows what to expect."

```java
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;

// Custom exceptions
class ResourceNotFoundException  extends RuntimeException {
    ResourceNotFoundException(String msg)  { super(msg); }
}
class BadRequestException extends RuntimeException {
    BadRequestException(String msg) { super(msg); }
}

// Consistent error response DTO
record ErrorResponse(int status, String error, String message, String path) {}

// ─── Single class handles ALL controller exceptions ───
@RestControllerAdvice
class GlobalExceptionHandler {

    @ExceptionHandler(ResourceNotFoundException.class)
    ResponseEntity<ErrorResponse> handleNotFound(
            ResourceNotFoundException ex,
            jakarta.servlet.http.HttpServletRequest req) {
        return ResponseEntity.status(HttpStatus.NOT_FOUND)
                .body(new ErrorResponse(404, "Not Found", ex.getMessage(), req.getRequestURI()));
    }

    @ExceptionHandler(BadRequestException.class)
    ResponseEntity<ErrorResponse> handleBadRequest(
            BadRequestException ex,
            jakarta.servlet.http.HttpServletRequest req) {
        return ResponseEntity.badRequest()
                .body(new ErrorResponse(400, "Bad Request", ex.getMessage(), req.getRequestURI()));
    }

    @ExceptionHandler(Exception.class) // catch-all
    ResponseEntity<ErrorResponse> handleGeneral(
            Exception ex,
            jakarta.servlet.http.HttpServletRequest req) {
        return ResponseEntity.internalServerError()
                .body(new ErrorResponse(500, "Internal Server Error",
                        "Something went wrong", req.getRequestURI()));
    }
}
```

---

## Q10. What is normalization in SQL? Difference between DELETE, TRUNCATE, and DROP?

**Spoken Answer:**
"**Normalization** is the process of organizing a relational database to eliminate redundancy and ensure data integrity. The normal forms (1NF, 2NF, 3NF, BCNF) achieve this progressively.

For DELETE vs TRUNCATE vs DROP:
- `DELETE` — removes specific rows based on a `WHERE` clause. It's DML, can be rolled back, fires triggers.
- `TRUNCATE` — removes ALL rows from a table instantly. DDL in most DBs, cannot have WHERE, faster than DELETE because it doesn't log each row.
- `DROP` — removes the entire table structure + data permanently from the database. Cannot be rolled back."

```sql
-- DELETE: removes specific rows, can rollback, fires triggers
DELETE FROM employees WHERE dept = 'HR';
-- DELETE FROM employees; -- removes all rows but table structure stays

-- TRUNCATE: removes all rows instantly, faster, no WHERE clause allowed
TRUNCATE TABLE employees;  -- table structure stays, identity resets

-- DROP: removes everything — structure, data, indexes, constraints
DROP TABLE employees;  -- table is gone completely

-- Normalization example:
-- ❌ Before normalization (1NF violation - repeating groups):
-- OrderID | CustomerName | Items
-- 1       | Alice        | Laptop, Mouse
-- 2       | Bob          | Keyboard

-- ✅ After 1NF (atomic values):
-- OrderID | CustomerName | Item
-- 1       | Alice        | Laptop
-- 1       | Alice        | Mouse
-- 2       | Bob          | Keyboard

-- ✅ After 3NF (no transitive dependency): split into Orders + Customers tables
-- Customers: CustomerID, CustomerName
-- Orders:    OrderID, CustomerID (FK), Item

-- SQL: Top 3rd highest marks
SELECT name, marks FROM (
    SELECT name, marks,
           DENSE_RANK() OVER (ORDER BY marks DESC) AS rnk
    FROM student
) ranked
WHERE rnk = 3;

-- SQL: Max salary per department
SELECT d.dept_name, MAX(e.salary) AS max_salary
FROM employee e
JOIN department d ON e.dept_id = d.id
GROUP BY d.dept_name;
```
