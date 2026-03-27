# 🏗️ Design Patterns & System Design Basics (Rapid-Fire)

> 🔑 **Master Keyword:** **"CBFS-OO"** → Creational/Behavioral/Factory/Singleton-Observer-Oops

---

## 🏭 Section 1: Creational Patterns

### Q1: Singleton Pattern — All Strategies?
🔑 **Keyword: "ELSDBE"** → Eager/Lazy/Synchronized/Double-checked/Bill-Pugh/Enum

```java
// 1. EAGER — created at class load (thread-safe, but no lazy)
class Singleton {
    private static final Singleton INSTANCE = new Singleton();
    private Singleton() {}
    public static Singleton getInstance() { return INSTANCE; }
}

// 2. LAZY — created on first call (NOT thread-safe)
class Singleton {
    private static Singleton instance;
    public static Singleton getInstance() {
        if (instance == null) instance = new Singleton();  // race condition!
        return instance;
    }
}

// 3. DOUBLE-CHECKED LOCKING — lazy + thread-safe (volatile required!)
class Singleton {
    private static volatile Singleton instance;
    public static Singleton getInstance() {
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) instance = new Singleton();
            }
        }
        return instance;
    }
}

// 4. BILL PUGH — best without enum (lazy + thread-safe via class loader)
class Singleton {
    private Singleton() {}
    private static class Holder { static final Singleton INSTANCE = new Singleton(); }
    public static Singleton getInstance() { return Holder.INSTANCE; }
}

// 5. ENUM SINGLETON — BEST! (Safe from Reflection + Serialization)
enum Singleton {
    INSTANCE;
    public void doSomething() { }
}
Singleton.INSTANCE.doSomething();
```

---

### Q2: Factory Pattern?
🔑 **Keyword: "FDC"** → Factory=Decouple-Creation from client

```java
// Interface
interface Shape { void draw(); }

// Implementations
class Circle implements Shape { public void draw() { System.out.println("Circle"); } }
class Square implements Shape { public void draw() { System.out.println("Square"); } }

// Factory — client doesn't know about Circle/Square
class ShapeFactory {
    public static Shape getShape(String type) {
        return switch(type.toUpperCase()) {
            case "CIRCLE" -> new Circle();
            case "SQUARE" -> new Square();
            default -> throw new IllegalArgumentException("Unknown: " + type);
        };
    }
}

// Usage
Shape s = ShapeFactory.getShape("CIRCLE"); // decoupled!
```

---

### Q3: Abstract Factory Pattern?
🔑 **Keyword: "FAF"** → Factory-of-Factories (family of related objects)

- Creates **families** of related objects without specifying concrete classes
- Example: `UIFactory` → `WindowsFactory` (WindowsButton + WindowsCheckbox) vs `MacFactory`

---

### Q4: Builder Pattern?
🔑 **Keyword: "BSC"** → Builder=Step-by-step-Construction

```java
// Builder solves "telescoping constructor" (too many params)
Person person = Person.builder()
    .name("Alice")
    .age(30)
    .email("alice@example.com")
    .build();

// With Lombok
@Builder
@Data
public class Person {
    private String name;
    private int age;
    private String email;
}
```

---

### Q5: Prototype Pattern?
🔑 **Keyword: "PC-Clone"** → Prototype=Clone existing object cheaply

```java
class DatabaseConnection implements Cloneable {
    private String config; // expensive to set up
    @Override
    public DatabaseConnection clone() throws CloneNotSupportedException {
        return (DatabaseConnection) super.clone(); // clone instead of recreate
    }
}
```

---

## 🏛️ Section 2: Structural Patterns

### Q6: Adapter Pattern?
🔑 **Keyword: "AWC"** → Adapter=Wrapper to make incompatible Compatible

```java
// Legacy interface
class OldPaymentSystem { void makeOldPayment(int amount) { } }

// New interface
interface ModernPayment { void pay(double amount); }

// Adapter — bridges the gap
class PaymentAdapter implements ModernPayment {
    private OldPaymentSystem old;
    public PaymentAdapter(OldPaymentSystem old) { this.old = old; }
    @Override
    public void pay(double amount) { old.makeOldPayment((int) amount); }
}
```

---

### Q7: Decorator Pattern?
🔑 **Keyword: "DWS"** → Decorator=Wrapper-Stack (add behavior dynamically)

```java
// Real-world: Java IO Streams!
BufferedReader br = new BufferedReader(  // add buffering
    new InputStreamReader(               // add encoding
        new FileInputStream("file.txt")  // base
    )
);
// Each layer adds behavior without modifying original!
```

---

### Q8: Proxy Pattern?
🔑 **Keyword: "PCC"** → Proxy=Control-access (like Spring AOP / Hibernate lazy)

```java
// Spring AOP = Proxy! Intercepts method calls to add behavior
@Transactional  // Spring creates a proxy around your service
public class OrderService { ... }

// Hibernate lazy loading = Proxy! getReferenceById returns proxy
User user = userRepo.getReferenceById(1L); // proxy, no DB hit
user.getName(); // now hits DB!
```

---

### Q9: Facade Pattern?
🔑 **Keyword: "FS"** → Facade=Simple-interface to complex-system

```java
// Complex subsystems
class VideoEncoder { void encode(String file) { } }
class AudioMixer { void mix(String file) { } }
class VideoUploader { void upload(String file) { } }

// Facade — single simple interface
class VideoProcessingFacade {
    void processAndUpload(String file) {
        new VideoEncoder().encode(file);
        new AudioMixer().mix(file);
        new VideoUploader().upload(file);
    }
}
// Client only talks to Facade!
```

---

## 🎭 Section 3: Behavioral Patterns

### Q10: Observer Pattern?
🔑 **Keyword: "PNS"** → Publisher→Notify→Subscribers (one-to-many)

```java
// Publisher (Subject)
class EventBus {
    private List<EventListener> listeners = new ArrayList<>();
    void subscribe(EventListener l) { listeners.add(l); }
    void publish(String event) { listeners.forEach(l -> l.onEvent(event)); }
}

// Subscribers
class EmailService implements EventListener {
    public void onEvent(String event) { System.out.println("Email: " + event); }
}

// Usage — used in Spring Events, RxJava, etc.
EventBus bus = new EventBus();
bus.subscribe(new EmailService());
bus.publish("ORDER_PLACED"); // all subscribers notified
```

---

### Q11: Strategy Pattern?
🔑 **Keyword: "SR"** → Strategy=Runtime-behavior-Switch

```java
interface SortingStrategy { void sort(int[] arr); }
class BubbleSort implements SortingStrategy { ... }
class QuickSort implements SortingStrategy { ... }

class Sorter {
    private SortingStrategy strategy;
    public Sorter(SortingStrategy strategy) { this.strategy = strategy; }
    public void sort(int[] arr) { strategy.sort(arr); } // delegate to strategy
}

// Switch behavior at runtime!
Sorter sorter = new Sorter(new QuickSort());
sorter = new Sorter(new BubbleSort()); // change strategy
```

---

### Q12: Command Pattern?
🔑 **Keyword: "EOR"** → Encapsulate-Operation for undo/Redo/Queue

```java
interface Command { void execute(); void undo(); }
class TextEditor {
    private String text = "";
    // Command
    class AddTextCommand implements Command {
        String addedText;
        AddTextCommand(String text) { this.addedText = text; }
        public void execute() { TextEditor.this.text += addedText; }
        public void undo() { text = text.substring(0, text.length() - addedText.length()); }
    }
}
```

---

### Q13: Template Method Pattern?
🔑 **Keyword: "ST-OH"** → Skeleton-Template, child Overrides-Hook-steps

```java
abstract class DataMiner {
    final void mine() {        // template method (final — not overridable)
        readData();            // fixed step
        parseData();           // abstract — child implements
        analyzeData();         // abstract — child implements
        sendReport();          // fixed step
    }
    abstract void parseData();
    abstract void analyzeData();
    void readData() { /* read from source */ }
    void sendReport() { /* send email */ }
}
```

---

## 🏛️ Section 4: SOLID Principles

### Q14: SOLID Principles?
🔑 **Keyword: "SOLID"** → Same acronym, each letter

| Letter | Principle | One-liner |
|---|---|---|
| **S** | Single Responsibility | One class, one job |
| **O** | Open/Closed | Open to extend, closed to modify |
| **L** | Liskov Substitution | Child should be substitutable for parent |
| **I** | Interface Segregation | Don't force unused method implementations |
| **D** | Dependency Inversion | Depend on abstractions, not concretions |

```java
// D - Dependency Inversion example
// BAD:
class OrderService { MySQLRepository repo = new MySQLRepository(); }

// GOOD:
class OrderService {
    private final OrderRepository repo; // interface (abstraction)
    OrderService(OrderRepository repo) { this.repo = repo; } // inject
}
```

---

## 🌐 Section 5: System Design Basics

### Q15: Microservices vs Monolith?
🔑 **Keyword: "MSSD"** → Microservices=Small-Services-Deployed-independently

| | Monolith | Microservices |
|---|---|---|
| Deployment | Single unit | Independent services |
| Scaling | Scale everything | Scale specific service |
| Development | Simpler start | More complex (distributed) |
| Failure | One failure = all down | Failure isolated |
| Tech stack | Single | Each service can differ |

---

### Q16: REST API Design Principles?
🔑 **Keyword: "SUSC"** → Stateless/Uniform-interface/Client-server/Cacheable

```
GET    /users         → list all users
GET    /users/{id}    → get specific user
POST   /users         → create user
PUT    /users/{id}    → replace user
PATCH  /users/{id}    → partial update
DELETE /users/{id}    → delete user
```

HTTP Status Codes to remember:
- `200 OK`, `201 Created`, `204 No Content`
- `400 Bad Request`, `401 Unauthorized`, `403 Forbidden`, `404 Not Found`
- `500 Internal Server Error`

---

### Q17: Load Balancing Algorithms?
🔑 **Keyword: "RRLIP"** → Round-Robin, Least-connections, IP-hash, Predictive

| Algorithm | How |
|---|---|
| Round Robin | Requests cycled evenly |
| Least Connections | Send to server with fewest active conns |
| IP Hash | Same client → same server (sticky) |
| Weighted Ring | Distribute based on server capacity |

---

### Q18: Caching Strategies?
🔑 **Keyword: "WRAP"** → Write-through/Read-through/Aside-pattern/Policies

```java
// Cache-Aside (most common in Spring)
@Cacheable("users")                    // read from cache
public User getUser(Long id) { ... }

@CacheEvict("users")                   // remove from cache
public void deleteUser(Long id) { ... }

@CachePut(value = "users", key = "#user.id")  // update cache
public User updateUser(User user) { ... }
```

Cache eviction policies: `LRU` (Least Recently Used), `LFU` (Least Frequently Used), `FIFO`

---

### Q19: CAP Theorem?
🔑 **Keyword: "CAP-2of3"** → Consistency/Availability/Partition-tolerance — can only guarantee 2

| Guarantee 2 of: | Examples |
|---|---|
| **C + P** (consistent + partition tolerant) | HBase, MongoDB during partition |
| **A + P** (available + partition tolerant) | Cassandra, DynamoDB |
| **C + A** (consistent + available) | MySQL, PostgreSQL (single node) |

---

### Q20: Common Java Design Patterns Used in Frameworks?

🔑 **Keyword: "SDOPP-frameworks"** → Singleton/Decorator/Observer/Proxy/Pattern

| Pattern | Where used |
|---|---|
| **Singleton** | Spring beans (default scope) |
| **Proxy** | Spring AOP, Hibernate lazy loading |
| **Decorator** | Java IO (`BufferedReader(FileReader)`) |
| **Observer** | Spring Events, JMS listeners |
| **Factory** | `BeanFactory`, `SessionFactory` |
| **Template Method** | `JdbcTemplate`, `RestTemplate`, `KafkaTemplate` |
| **Command** | Spring Batch, CQRS |
| **Strategy** | Spring Security `AuthenticationProvider` |

---

*End of File — Design Patterns & System Design Basics*
