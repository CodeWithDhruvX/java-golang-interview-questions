# Java Full Stack Developer Interview Answers

## 1. Java Full Stack Developer (General)

### Core Java
- **Difference between JDK, JRE, JVM**:
  - **JDK (Java Development Kit)**: Includes JRE + development tools (compiler, debugger, javadoc).
  - **JRE (Java Runtime Environment)**: Includes JVM + core libraries needed to run Java applications.
  - **JVM (Java Virtual Machine)**: Executes bytecode and provides the runtime environment.

- **OOP concepts**:
  - **Encapsulation**: Wrapping data and methods into a single unit (class) and restricting access (private fields).
  - **Inheritance**: Acquiring properties from a parent class (code reuse).
  - **Polymorphism**: One interface, many forms (Method Overloading & Overriding).
  - **Abstraction**: Hiding implementation details and showing only functionality (Abstract classes & Interfaces).

- **Abstract class vs Interface**:
  - **Abstract Class**: Can have constructors, instance variables, concrete methods. Supports "is-a" relationship. Only single inheritance.
  - **Interface**: By default methods are public abstract (Java 8 added default/static methods). Supports "can-do" capability. Supports multiple inheritance.

- **String vs StringBuilder vs StringBuffer**:
  - **String**: Immutable. String operations create new objects. Thread-safe (due to immutability).
  - **StringBuilder**: Mutable. Not thread-safe (faster).
  - **StringBuffer**: Mutable. Thread-safe (synchronized methods, slower).

- **Immutability**: Object whose state cannot be modified after creation. Achieved by making class `final`, fields `private final`, and no setters. (e.g., String, Wrapper classes).

- **equals() vs ==**:
  - **==**: Compares references (memory address) for objects; compares values for primitives.
  - **equals()**: Method to compare object content (must be overridden for custom equality logic).

- **hashCode() contract**: If two objects are equal according to `equals()`, they must have the same `hashCode()`. However, same hash code doesn't guarantee equality (collision).

- **Exception Hierachy**: `Throwable` is the root.
  - **Error**: Serious problems (OutOfMemoryError), usually not recoverable.
  - **Exception**: Conditions a program might want to catch.
    - **Checked**: Compile-time (IOException, SQLException). Must be caught or declared.
    - **Unchecked**: Runtime (NullPointerException, ArrayIndexOutOfBoundsException).

- **final vs finally vs finalize**:
  - **final**: Keyword to restrict modification (variables), inheritance (classes), or overriding (methods).
  - **finally**: Block in try-catch to execute cleanup code (always runs unless JVM exits).
  - **finalize**: Method called by GC before object reclamation (deprecated in Java 9).

- **Collections vs Arrays**:
  - **Array**: Fixed size, holds primitives or objects.
  - **Collection**: Dynamic size, holds only objects. extensive utility methods.

### Java Collections Framework
- **List vs Set vs Map**:
  - **List**: Ordered, allows duplicates.
  - **Set**: Unordered, unique elements only.
  - **Map**: Key-Value pairs, unique keys.

- **ArrayList vs LinkedList**:
  - **ArrayList**: Backed by dynamic array. Fast random access (O(1)). Slow insertion/deletion (shifting).
  - **LinkedList**: Doubly linked list. Fast insertion/deletion (O(1)). Slow random access (O(n)).

- **HashMap vs TreeMap vs LinkedHashMap**:
  - **HashMap**: No order. O(1) avg access.
  - **TreeMap**: Sorted order (Red-Black tree). O(log n).
  - **LinkedHashMap**: Insertion order preserved. Slighty slower than HashMap.

- **HashMap Internal Working**: Uses hashing. `hashCode()` determines bucket index. If collision, uses linked list (or Red-Black tree after threshold 8 in Java 8+) to store multiple entries. `equals()` checks key uniqueness.

- **ConcurrentHashMap**: Thread-safe map efficient for high concurrency. Uses segment locking (Java 7) or CAS + synchronized on node heads (Java 8+), allowing concurrent reads/writes without locking the entire map.

- **Fail-fast vs Fail-safe**:
  - **Fail-fast**: Throws `ConcurrentModificationException` if collection is modified during iteration (e.g., ArrayList, HashMap).
  - **Fail-safe**: Iterates over a clone or snapshot, no exception on modification (e.g., CopyOnWriteArrayList, ConcurrentHashMap).

### Java 8 & Above
- **Lambda Expressions**: Concise way to implement functional interfaces (anonymous functions). Syntax: `(args) -> body`.
- **Functional Interface**: Interface with exactly one abstract method. Annotated with `@FunctionalInterface`. (e.g., Runnable, Callable, Comparator).
- **Stream API**: Sequence of elements supporting sequential/parallel aggregate operations (filter, map, reduce). Doesn't store data.
- **map() vs flatMap()**:
  - **map()**: One-to-one transformation. Input stream `Stream<T>` -> Output `Stream<R>`.
  - **flatMap()**: One-to-many transformation. Flattens nested structures. Input `Stream<List<T>>` -> Output `Stream<T>`.
- **Optional Class**: Container object which may or may not contain a non-null value. Avoids NullPointerException.
- **Default Methods**: Allows adding methods with implementation in interfaces without breaking implementing classes.

### Multithreading & Concurrency
- **Thread vs Runnable**:
  - **Thread**: Extend `Thread` class. Cannot extend other classes.
  - **Runnable**: Implement interface. Better for flexibility (can extend another class).

- **Thread Lifecycle**: New -> Runnable -> Running -> Blocked/Waiting -> Terminated.
- **synchronized**: Ensures only one thread executes a block/method at a time (Mutual Exclusion).
- **Deadlock**: Two threads waiting for each other significantly to release locks held by the other. Avoid by acquiring locks in consistent order or using timeouts.
- **volatile**: Variable value is always read from main memory, not thread cache. Ensures visibility but NOT atomicity.
- **Executor Framework**: Decouples task submission from execution. Manages thread pools (Fixed, Cached, Scheduled) efficiently.

### Spring Framework
- **Spring**: Enterprise Java framework providing comprehensive infrastructure support (IoC, AOP, Data Access, Web).
- **IoC (Inversion of Control)**: Transferring control of object creation and dependency management from application code to the container.
- **Dependency Injection (DI)**: Pattern to implement IoC. Dependencies are "injected" into objects (Constructor, Setter, Field).
- **Bean Lifecycle**: Instantiate -> Populate Props -> BeanNameAware -> BeanFactoryAware -> PostProcessBeforeInit -> Init (@PostConstruct/afterPropertiesSet) -> PostProcessAfterInit -> Destroy (@PreDestroy).
- **@Component vs @Service vs @Repository**:
  - **@Component**: Generic stereotype for any Spring-managed component.
  - **@Service**: Semantic specialization for Service layer (business logic).
  - **@Repository**: Persistence layer. Adds translation of unchecked exceptions.
- **@Autowired**: Marks a constructor, field, or setter for DI type-matching.
- **Why Spring Boot**: Simplifies Spring app development. Convention over configuration. Standalone (embedded server), production-ready defaults, no XML required.
- **Auto-configuration**: Spring Boot automatically configures beans based on classpath dependencies (e.g., adds H2 DB bean if H2 jar is present).
- **application.properties vs yml**: Key-value format vs Hierarchical format (YAML). YAML is more readable for complex configs.
- **@SpringBootApplication**: Meta-annotation combining `@Configuration`, `@EnableAutoConfiguration`, `@ComponentScan`.
- **@Controller vs @RestController**:
  - **@Controller**: Returns views (JSP/Thymeleaf).
  - **@RestController**: Returns data (JSON/XML). @Controller + @ResponseBody.
- **@RequestBody vs @RequestParam vs @PathVariable**:
  - **@RequestBody**: Maps HTTP body to Java object.
  - **@RequestParam**: Extract query parameters (`/url?id=1`).
  - **@PathVariable**: Extract URI template variables (`/url/{id}`).

### REST API & Microservices
- **REST**: Representational State Transfer. Architectural style for distributed hypermedia systems.
- **HTTP Methods**:
  - **GET**: Retrieve resource.
  - **POST**: Create resource.
  - **PUT**: Update/Replace resource.
  - **DELETE**: Remove resource.
  - **PATCH**: Partial update.
- **Status Codes**: 2xx (Success), 3xx (Redirection), 4xx (Client Error), 5xx (Server Error).
- **Idempotent**: Making multiple identical requests has the same effect as a single request (GET, PUT, DELETE).
- **Statelessness**: Server contains no client state. Each request must contain all info.
- **Monolith vs Microservices**:
  - **Monolith**: Single codebase/deployment unit. Tightly coupled. Harder to scale components.
  - **Microservices**: Small, independent services. Loosely coupled. Scalable, tech-agnostic. Challenge: complexity.
- **Service Discovery**: Mechanism for services to find each other (e.g., Eureka, Consul). Client-side or Server-side.
- **API Gateway**: Single entry point for clients. Handles routing, authentication, rate limiting (Zuul, Spring Cloud Gateway).
- **Circuit Breaker**: Prevents cascading failures when a service is down. Fails fast or provides fallback (Resilience4j, Hystrix).

### Hibernate / JPA
- **ORM**: Object-Relational Mapping. Maps Java objects to DB tables.
- **Hibernate vs JPA**: JPA is the specification (API); Hibernate is the implementation.
- **Entity Lifecycle**: Transient (new), Persistent (managed), Detached (session closed), Removed.
- **Fetch Types**:
  - **EAGER**: Load related entities immediately.
  - **LAZY**: Load on demand (when accessed). Better for performance.
- **N+1 Problem**: Executing 1 query to fetch parent, then N queries for N children. Fix: `JOIN FETCH` or `@BatchSize`.

### Database (SQL + Basics)
- **Joins**:
  - **INNER**: Matched rows in both tables.
  - **LEFT**: All rows from left, matched from right.
  - **RIGHT**: All rows from right, matched from left.
- **Primary vs Foreign Key**:
  - **PK**: Unique identifier for row. Not Null.
  - **FK**: Reference to PK in another table. Enforces referential integrity.
- **Indexes**: improve read speed. B-Tree structure. Slows down writes.
- **Normalization**: Organizing data to minimize redundancy (1NF, 2NF, 3NF).
- **ACID**: Atomicity (all or nothing), Consistency (valid state), Isolation (transactions independent), Durability (saved permanently).
- **DELETE vs TRUNCATE vs DROP**:
  - **DELETE**: DML. Deletes rows. Can rollback. Slow.
  - **TRUNCATE**: DDL. Resets table. Cannot rollback. Fast.
  - **DROP**: DDL. Deletes entire table structure.

### System Design
- **Stateless Services**: Services that don't store session state. Easier to scale horizontally.
- **Load Balancing**: Distributing traffic across multiple servers (Round Robin, Least Conn).
- **Caching**: Storing frequently accessed data in memory (Redis, Memcached) to reduce DB load.

## 2. Senior Developer (5+ Years - Deep Dive)

### Core Java & JVM
- **ClassLoader Hierarchy**: Bootstrap -> Platform/Extension -> App/System ClassLoader.
- **Garbage Collection (GC)**: Automatic memory management. Mark-and-Sweep.
  - **Generational**: Eden, Survivor, Old Gen.
  - **Algorithms**: Serial, Parallel, CMS (deprecated), G1 (default), ZGC.
- **Shallow vs Deep Copy**:
  - **Shallow**: Copies references. Changes to nested objects affect original.
  - **Deep**: Copies objects recursively. Completely independent clone.
- **Memory Management**: Stack (method frames, primitives, references) vs Heap (Objects).
- **Metaspace vs PermGen**:
  - **PermGen**: Fixed size, part of Heap (Java 7). OOM frequent.
  - **Metaspace**: Native memory, auto-resizable (Java 8+).

### Concurrency
- **Synchronized vs Lock**:
  - **synchronized**: Implicit lock, scoped to block/method. Automatic release.
  - **Lock (ReentrantLock)**: Explicit. `lock()`/`unlock()`. Flexible (tryLock, multiple conditions).
- **ConcurrentHashMap Internal**:
  - **Java 7**: Segment locking (16 segments).
  - **Java 8+**: Node-level locking using CAS (Compare-And-Swap) for insertions and `synchronized` only on the specific bin (bucket) head during collision.
- **ForkJoinPool**: Optimized for divide-and-conquer (recursive) tasks. Uses work-stealing algorithm.

### Spring Internals
- **Bean Scopes**: Singleton (default), Prototype, Request, Session, GlobalSession.
- **Circular Dependency**: Bean A needs B, B needs A. Spring solves via setter injection (creates A, caches unfinished A, creates B injecting A, finishes A). Constructor injection fails.
- **@Lazy**: Bean initialized only when requested, not at startup.
- **ApplicationContext vs BeanFactory**: AppContext extends BeanFactory. Adds AOP, i18n, Event publishing.

### Microservices Patterns
- **Saga Pattern**: Managing distributed queries spanning multiple services.
  - **Choreography**: Events trigger next step.
  - **Orchestration**: Central coordinator directs steps.
- **Event-Driven**: Services communicate via events (Async). Loose coupling.
- **CQRS**: Command Query Responsibility Segregation. Separate models for read and write.

### Database Performance
- **Isolation Levels**:
  - **Read Uncommitted**: Dirty reads allowed.
  - **Read Committed**: No dirty reads (default in many DBs).
  - **Repeatable Read**: No non-repeatable reads.
  - **Serializable**: Strict, no phantoms.
- **Dirty Checking**: Hibernate automatically detects changes to persistent objects and syncs to DB on flush.

## 3. Strong Senior (Separator Questions)

- **OutOfMemoryError Troubleshooting**:
  1. Analyze stack trace.
  2. Generate Heap Dump (`jmap`).
  3. Analyze dump with Eclipse MAT / VisualVM.
  4. Look for large objects, memory leaks (references causing retention).
  5. Check JVM arguments (`-Xmx`).

- **Memory Leak**: Objects are no longer needed but referenced, preventing GC. Symptoms: Gradual heap usage increase. Fix: Remove references (static fields, listener deregistration).

- **Stop-the-World**: GC pauses all application threads to mark/compact. High pause time affects latency.

- **Volatile vs Synchronized**: Volatile guarantees visibility and ordering (happens-before), but NOT atomicity. Synchronized guarantees visibility AND atomicity. Volatile is not a replacement for counters/mutex.

- **Parallel Streams Pitfall**:
  - Uses common ForkJoinPool (shared across app).
  - Blocking tasks in parallel stream can starve the pool, freezing the entire app.
  - Use purely for CPU-intensive, non-blocking tasks.

- **Spring Boot Auto-Config**: Uses `@Conditional` annotations (`@ConditionalOnClass`, `@ConditionalOnMissingBean`). Checks classpath and beans to decide configuration configurations.

- **N+1 Example**: fetching List of `Orders`. Loop through orders and call `order.getCustomer()`. If `Customer` is Lazy, 1 query for Orders + N queries for N Customers.

- **Idempotency Design**: Use idempotency keys (unique request ID) in header. Store key in DB/Cache with response. Return cached response for duplicate keys.

- **Distributed Transactions**: Avoid if possible (CAP theorem). Use Eventual Consistency (Saga). 2PC (Two-Phase Commit) is too slow/locking.

- **Kafka Delivery Guarantees**:
  - **At-most-once**: Fire and forget.
  - **At-least-once**: Retry until Ack. Duplicates possible (Idempotent consumer required).
  - **Exactly-once**: Transactional support (Kafka Streams).

- **JWT Internals**: Header (Algorithm), Payload (Claims, User Info, Exp), Signature (Base64UrlEncoded(header) + "." + Base64UrlEncoded(payload), Secret). Stateless auth.

- **Rate Limiter Design**: Token Bucket or Leaky Bucket algorithm. Redis to store counters/timestamps. Middleware filter implementation.
