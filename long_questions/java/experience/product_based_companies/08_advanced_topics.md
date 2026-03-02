# 🔮 08 — Advanced Java Topics (Java 14–21)
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Records (Java 16)
- Sealed classes (Java 17)
- Pattern matching for switch (Java 21)
- Virtual threads (Java 21 — Project Loom)
- Generics, wildcards, and type bounds
- Reflection and dynamic proxy
- Java modules (JPMS)

---

## ❓ Most Asked Questions

### Q1. What are Java Records? (Java 16+)

```java
// Record — immutable data carrier, auto-generates: constructor, getters, equals, hashCode, toString

// Traditional DTO — lots of boilerplate
public class UserDTO {
    private final Long id;
    private final String name;
    private final String email;

    public UserDTO(Long id, String name, String email) {
        this.id = id; this.name = name; this.email = email;
    }
    public Long getId() { return id; }
    public String getName() { return name; }
    // equals, hashCode, toString...
}

// Record — equivalent, much less code
public record UserDTO(Long id, String name, String email) {}

// Usage — same as DTO
UserDTO user = new UserDTO(1L, "Alice", "alice@example.com");
System.out.println(user.name());    // accessor (no "get" prefix)
System.out.println(user);          // UserDTO[id=1, name=Alice, email=alice@example.com]

// Custom compact constructor — add validation
public record Product(String name, BigDecimal price, int stock) {
    public Product {  // compact constructor — no parameter list
        Objects.requireNonNull(name, "name required");
        if (price.compareTo(BigDecimal.ZERO) <= 0) throw new IllegalArgumentException("Price must be positive");
        if (stock < 0) throw new IllegalArgumentException("Stock cannot be negative");
        name = name.trim();  // can normalize
    }

    // Can add custom methods
    public boolean inStock() { return stock > 0; }
    public BigDecimal totalValue() { return price.multiply(BigDecimal.valueOf(stock)); }
}

// Records with Spring — perfect for DTOs, API responses
@RestController
public class ProductController {
    @GetMapping("/{id}")
    public ProductDTO getProduct(@PathVariable Long id) {
        return new ProductDTO(1L, "Phone", new BigDecimal("999.99"));  // no setters needed!
    }
    record ProductDTO(Long id, String name, BigDecimal price) {}
}
```

---

### Q2. What are Sealed Classes? (Java 17+)

```java
// Sealed class — restricts which classes can extend it
// Useful for: domain modeling, exhaustive pattern matching

public sealed interface Payment
    permits CreditCardPayment, UpiPayment, WalletPayment {

    BigDecimal amount();
    String currency();
}

public record CreditCardPayment(
    BigDecimal amount, String currency, String last4Digits, String network
) implements Payment {}

public record UpiPayment(
    BigDecimal amount, String currency, String upiId
) implements Payment {}

public record WalletPayment(
    BigDecimal amount, String currency, String walletProvider
) implements Payment {}

// Exhaustive switch — compiler checks all cases! (Java 21)
public String processPayment(Payment payment) {
    return switch (payment) {
        case CreditCardPayment cc -> "Charging card ending " + cc.last4Digits();
        case UpiPayment upi       -> "Processing UPI to " + upi.upiId();
        case WalletPayment wallet -> "Debiting " + wallet.walletProvider() + " wallet";
        // No default needed — compiler knows all subtypes!
    };
}

// Adding a new subtype (e.g., BankTransferPayment) forces updating all switch statements
// → Exhaustive type-safe modeling
```

---

### Q3. What is Pattern Matching for switch? (Java 21)

```java
// Old style — messy instanceof chain
static String format(Object obj) {
    if (obj instanceof Integer i) return "int: " + i;
    else if (obj instanceof Double d) return "double: " + d;
    else if (obj instanceof String s && !s.isBlank()) return "string: " + s.toUpperCase();
    else return "unknown";
}

// Java 21 — clean pattern switch
static String format(Object obj) {
    return switch (obj) {
        case Integer i       -> "int: " + i;
        case Double d        -> "double: " + d;
        case String s when !s.isBlank() -> "string: " + s.toUpperCase();  // guard
        case null            -> "null";
        default              -> "unknown: " + obj.getClass().getSimpleName();
    };
}

// With sealed hierarchy — no default needed
sealed interface Shape permits Circle, Rectangle, Triangle {}
record Circle(double radius) implements Shape {}
record Rectangle(double w, double h) implements Shape {}
record Triangle(double base, double height) implements Shape {}

double area(Shape shape) {
    return switch (shape) {
        case Circle c    -> Math.PI * c.radius() * c.radius();
        case Rectangle r -> r.w() * r.h();
        case Triangle t  -> 0.5 * t.base() * t.height();
    };  // exhaustive — no default needed
}
```

---

### Q4. Generics, Wildcards, and Bounded Types

```java
// Bounded type parameters
public <T extends Comparable<T>> T findMax(List<T> list) {
    return list.stream().max(Comparator.naturalOrder()).orElseThrow();
}

// Upper bounded wildcard — read from collection of T or subtypes
public double sumAreas(List<? extends Shape> shapes) {
    return shapes.stream().mapToDouble(Shape::area).sum();
}
// ✅ Can READ shapes — List<Circle> and List<Rectangle> are both accepted
// ❌ Cannot ADD to shapes — type unsafe

// Lower bounded wildcard — write into collection of T or supertypes (PECS rule)
public void addNumbers(List<? super Integer> list) {
    list.add(1); list.add(2); list.add(3);  // can add Integer (or subtypes)
}
// ✅ Can ADD integers — List<Integer> and List<Number> are both accepted
// ❌ Cannot reliably READ as specific type

// PECS: Producer Extends, Consumer Super
// Reading from collection (producing elements) → use ? extends T
// Writing to collection (consuming elements) → use ? super T

// Type token pattern — capture generic type at runtime
public class TypedParser<T> {
    private final Class<T> type;
    public TypedParser(Class<T> type) { this.type = type; }

    public T parse(String json) {
        return objectMapper.readValue(json, type);
    }
}

new TypedParser<>(UserDTO.class).parse(json);
```

---

### Q5. Reflection and Dynamic Proxy

```java
// Reflection — inspect/invoke at runtime
Class<?> clazz = Class.forName("com.example.UserService");
Object instance = clazz.getDeclaredConstructor().newInstance();

// Get and invoke method
Method method = clazz.getDeclaredMethod("findById", Long.class);
method.setAccessible(true);  // bypass access control
Object result = method.invoke(instance, 42L);

// Create dynamic proxy — AOP-like behavior at runtime
LoggingHandler handler = new LoggingHandler(realService);
OrderService proxy = (OrderService) Proxy.newProxyInstance(
    OrderService.class.getClassLoader(),
    new Class[]{OrderService.class},
    handler);

// InvocationHandler — runs for every method call on proxy
class LoggingHandler implements InvocationHandler {
    private final Object target;
    LoggingHandler(Object target) { this.target = target; }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        System.out.println("Calling: " + method.getName());
        long start = System.currentTimeMillis();
        Object result = method.invoke(target, args);
        System.out.println("Done in " + (System.currentTimeMillis() - start) + "ms");
        return result;
    }
}
// Spring AOP uses JDK dynamic proxies (for interfaces) or CGLIB bytecode (for classes)
```

---

### Q6. What is the Java Module System (JPMS)?

```java
// module-info.java — declare module dependencies and exports
// Placed in: src/main/java/module-info.java

module com.example.orderservice {
    requires java.net.http;                    // JDK module
    requires spring.context;                   // Spring module
    requires com.example.shared;               // another custom module

    exports com.example.orderservice.api;      // public API — visible to dependents
    exports com.example.orderservice.dto;

    // Internal packages NOT exported — encapsulated!
    // com.example.orderservice.internal is invisible to external modules

    opens com.example.orderservice.entity to  // allow reflection (e.g., for Jackson)
        com.fasterxml.jackson.databind;

    provides com.example.shared.spi.PaymentProcessor
        with com.example.orderservice.PaymentProcessorImpl;  // ServiceLoader SPI

    uses com.example.shared.spi.NotificationSender;  // consume SPI
}
```

| Concept | Description |
|---------|-------------|
| `requires` | Declares dependency on another module |
| `exports` | Makes package accessible to other modules |
| `opens` | Allows deep reflection into package |
| `provides ... with` | Registers a service implementation |
| `uses` | Declares consumption of a service |

---

### Q7. Functional Interfaces and Advanced Streams

```java
// Common functional interfaces
Function<String, Integer>    parse    = Integer::parseInt;         // T → R
Predicate<String>            isEmpty  = String::isEmpty;           // T → boolean
Consumer<String>             print    = System.out::println;       // T → void
Supplier<List<String>>       newList  = ArrayList::new;            // () → T
BiFunction<Integer, Integer, Integer> add = Integer::sum;          // (T, U) → R
UnaryOperator<String>        upper    = String::toUpperCase;       // T → T
BinaryOperator<Integer>      max      = Integer::max;              // (T, T) → T

// Advanced Streams
// Collectors
Map<String, List<Order>> byCustomer = orders.stream()
    .collect(Collectors.groupingBy(Order::getCustomerId));

Map<String, Long> countByStatus = orders.stream()
    .collect(Collectors.groupingBy(o -> o.getStatus().name(), Collectors.counting()));

// Partitioning
Map<Boolean, List<Order>> highValue = orders.stream()
    .collect(Collectors.partitioningBy(o -> o.getTotal().compareTo(new BigDecimal("1000")) > 0));

// FlatMap — flatten nested collections
List<String> allTags = products.stream()
    .flatMap(p -> p.getTags().stream())
    .distinct()
    .sorted()
    .toList();

// Teeing collector (Java 12+) — two collectors in one pass
var stats = numbers.stream().collect(
    Collectors.teeing(
        Collectors.summingInt(n -> n),     // sum
        Collectors.counting(),              // count
        (sum, count) -> new Stats(sum, count)
    ));

// Gather (Java 22 preview) — custom stateful intermediate ops — future-proof knowledge
```
