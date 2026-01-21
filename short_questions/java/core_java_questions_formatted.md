# Core Java Interview Questions & Answers (Summary Version)

> **Quick reference guide with concise explanations for Core Java interview questions 1-170**

---

## 🧠 Core Java (Questions 1-25)

**Q1: What are the key features introduced in Java 8?**
Lambda expressions, Stream API, Optional class, Default methods in interfaces, Functional interfaces, Date/Time API (java.time), Method references.

**Q2: How do lambda expressions work internally?**
They don't create anonymous inner classes at compile time. Instead, they usage `invokedynamic` opcode to bind the method handle at runtime, often creating a private static method for the lambda body.

**Q3: What is the difference between `map`, `flatMap`, and `filter` in streams?**
`map`: Transforms each element 1-to-1.
`flatMap`: Transforms each element into a stream and flattens them (1-to-Many).
`filter`: Selects elements based on a boolean predicate.

**Q4: When would you use `Optional`, and when should you avoid it?**
Use as a return type to signal "value might be missing". Avoid using as parameters, class fields, or wrapping collections (return empty collection instead).

**Q5: Difference between `==` and `equals()`?**
`==` compares object references (memory address) or primitive values. `equals()` compares object content (logical equality) if overridden; otherwise defaults to `==`.

**Q6: Why must `hashCode()` be overridden when `equals()` is overridden?**
To maintain the contract: Equal objects must have equal hash codes. If not, hash-based collections (HashMap, HashSet) won't work correctly (can't find stored objects).

**Q7: Difference between `String`, `StringBuilder`, and `StringBuffer`?**
`String`: Immutable.
`StringBuilder`: Mutable, not thread-safe, faster.
`StringBuffer`: Mutable, thread-safe (synchronized), slower.

**Q8: What are immutable objects? Why are they important?**
Objects whose state cannot change after creation. Important for thread safety (no synchronization needed), caching, and safe use as map keys.

**Q9: What happens when you make a class immutable?**
Declared `final`, all fields `private final`. No setters. Mutable fields are defensively copied in constructor and getters.

**Q10: Difference between `ArrayList` and `LinkedList`?**
`ArrayList`: Dynamic array, fast random access (O(1)), slow insertion/deletion (O(n)).
`LinkedList`: Doubly linked list, slow access (O(n)), fast insertion/deletion (O(1)).

**Q11: Difference between `HashMap` and `ConcurrentHashMap`?**
`HashMap`: Not thread-safe, fast.
`ConcurrentHashMap`: Thread-safe, uses bucket-level locking (CAS + synchronized) instead of locking the whole map.

**Q12: How does `HashMap` work internally?**
Uses an array of buckets. `hash(key)` calculates index. Collisions are handled via Linked List (or Red-Black Tree if bucket size > 8).

**Q13: What happens when two keys have the same hashcode?**
Collision occurs. Both keys are stored in the same bucket. `equals()` is used to distinguish between them during retrieval.

**Q14: Difference between `Comparable` and `Comparator`?**
`Comparable`: Defines "natural" order (`compareTo`), implemented by the class itself.
`Comparator`: Defines "custom" order (`compare`), implemented by separate classes.

**Q15: Why are generics invariant in Java?**
To ensure type safety. `List<String>` is NOT a subtype of `List<Object>`. Prevents inserting wrong types into collections effectively.

**Q16: What are wildcards (`? extends`, `? super`)?**
`? extends T`: Upper bound (Producer), safe to read T.
`? super T`: Lower bound (Consumer), safe to write T.

**Q17: Difference between `Set`, `List`, and `Map`?**
`List`: Ordered, allows duplicates (ArrayList).
`Set`: Unique elements, unordered (HashSet) or sorted (TreeSet).
`Map`: Key-Value pairs, unique keys.

**Q18: When would you use `EnumMap` or `EnumSet`?**
When keys/elements are Enums. Internally uses arrays, highly optimized, memory efficient, faster than Hash variants.

**Q19: Difference between `Thread` and `Runnable`?**
`Thread`: Class representing a thread of execution. Limited (single inheritance).
`Runnable`: Functional interface representing a task. Better separation of task vs execution, allows extending other classes.

**Q20: What is the Java Memory Model?**
Defines how threads interact through memory. Key concepts: Atomicity, Visibility (volatile), Ordering (happens-before), Main Memory vs Local Cache.

**Q21: Difference between `volatile` and `synchronized`?**
`volatile`: Ensures visibility (reads/writes go to main memory), prevents reordering. No atomicity.
`synchronized`: Ensures visibility AND atomicity (mutual exclusion).

**Q22: What is a deadlock? How do you prevent it?**
Two threads wait forever for each other's locks. Prevention: Lock ordering, lock timeout (tryLock), avoid nested locks.

**Q23: What is `ExecutorService`?**
Framework for asynchronous task execution. Manages thread pool lifecycle. Decouples task submission from execution mechanics.

**Q24: Difference between `Callable` and `Runnable`?**
`Runnable`: `run()` returns void, cannot throw checked exceptions.
`Callable`: `call()` returns a value (`Future`), can throw exceptions.

**Q25: What are atomic classes (`AtomicInteger`)?**
Wrapper classes providing thread-safe operations on single variables without locks. Uses low-level CPU Compare-And-Swap (CAS) instructions.

---

## ⚙️ JVM, Memory & GC (Questions 26-35)

**Q26: What are the JVM memory areas?**
Heap (Objects), Stack (Methods/Primitives), Metaspace (Class metadata), PC Register, Native Method Stack.

**Q27: Difference between Heap and Stack?**
Heap: Shared globally, stores objects, huge.
Stack: Per-thread, stores local vars/frames, small, fast access, LIFO.

**Q28: What is garbage collection?**
Automatic memory management process that identifies and deletes unused objects to free up heap space.

**Q29: Difference between minor GC and major GC?**
Minor GC: Cleans Young Generation (Eden/Survivor). Fast, happens often.
Major GC (Full GC): Cleans Old Generation (and usually Young). Slow, pauses application.

**Q30: What causes `OutOfMemoryError`?**
Heap space full (memory leak or too much data), Metaspace full, or excessive thread creation.

**Q31: Difference between `OutOfMemoryError` and `StackOverflowError`?**
OOM: Heap full (no space for new objects).
StackOverflow: Stack full (infinite recursion or too deep calls).

**Q32: What are GC roots?**
Objects always accessible: Local variables, Static variables, Active threads, JNI references. GC starts tracing from here.

**Q33: What is stop-the-world?**
Scenario where JVM pauses all application threads to perform Garbage Collection safely.

**Q34: How do you analyze memory leaks?**
Heap dump analysis (VisualVM, Eclipse MAT). Look for objects with unexpected references preventing GC.

**Q35: What JVM options have you used?**
`-Xms` (min heap), `-Xmx` (max heap), `-XX:+UseG1GC` (GC algorith), `-XX:MaxMetaspaceSize`, `-XX:+HeapDumpOnOutOfMemoryError`.

---

## 🧩 OOP, SOLID & Design Patterns (Questions 36-50)

**Q36: Explain all SOLID principles with examples.**
S (SRP): One reason to change.
O (OCP): Open extension, closed modification.
L (LSP): Subtypes substitutable for base types.
I (ISP): Specific interfaces > General ones.
D (DIP): Depend on abstractions, not concretions.

**Q37: Difference between abstraction and encapsulation?**
Abstraction: Hiding complexity (interfaces). "What it does".
Encapsulation: Hiding state/data (private fields + methods). "How it does it".

**Q38: What is composition vs inheritance?**
Inheritance: "Is-a" relationship (Code reuse via hierarchy). Fragile.
Composition: "Has-a" relationship (Building complex objects from parts). Flexible, looser coupling.

**Q39: What is dependency injection?**
Inverting control. Dependencies are provided to a class (constructor/setter) rather than created inside it. Improves testability.

**Q40: How does SOLID help in real projects?**
Reduces coupling, increases cohesion. Makes code easier to maintain, test, extend, and refactor without breaking unrelated parts.

**Q41: Singleton pattern – problems and solutions?**
Problems: Global state, hard to test, hides dependencies. Solution: Use Enum Singleton or Dependency Injection frameworks (Spring).

**Q42: How do you implement thread-safe Singleton?**
Double-Checked Locking, Static Inner Helper Class (Bill Pugh), or simply use Enum (best).

**Q43: Factory vs Abstract Factory?**
Factory Method: Creates one type of object.
Abstract Factory: Creates families of related objects without specifying concrete classes.

**Q44: When would you use Builder pattern?**
When constructing complex objects with many optional parameters. Makes creation readable (`.setX().setY().build()`).

**Q45: Strategy pattern real-world use case?**
Payment processing (CreditCard, PayPal, Bitcoin strategies). Client chooses strategy at runtime.

**Q46: Observer pattern example?**
Event listeners (UI), Pub/Sub systems. Object (Subject) maintains list of dependents (Observers) and notifies them of state changes.

**Q47: Anti-patterns you have seen?**
God Class, Spaghetti Code, Hardcoded values, Swallowing exceptions, Copy-Paste programming.

**Q48: What is clean code?**
Readable, simple, focused (SRP), testable. Naming matters. "Code that looks like it was written by someone who cares."

**Q49: How do you reduce tight coupling?**
Interfaces, Dependency Injection, Event-driven architecture, Favor composition over inheritance.

**Q50: How do you design extensible code?**
Use Design Patterns (Strategy, Decorator), SOLID principles (Open/Closed), avoid hardcoding logic.

---

## 🌐 Spring Boot, REST & JPA (Questions 51-70)

**Q51: What happens internally when a Spring Boot app starts?**
Bootstrap context, component scan, auto-configuration (checks classpath), creates Beans, starts embedded server (Tomcat).

**Q52: Difference between `@Component`, `@Service`, `@Repository`?**
Stereotypes. `@Component`: Generic. `@Service`: Business logic. `@Repository`: DB access (translates SQL exceptions).

**Q53: What is auto-configuration?**
Spring Boot automatically configures beans based on classpath dependencies (e.g., sees H2 -> configures DataSource).

**Q54: What is Spring Bean lifecycle?**
Instantiate -> Populate Properties -> setBeanName -> PostProcessBeforeInit -> @PostConstruct/init-method -> PostProcessAfterInit -> Ready -> @PreDestroy.

**Q55: Difference between `@Autowired` and constructor injection?**
Constructor: Mandatory dependencies, easier testing, immutability.
Autowired: Field/Setter injection. Hidden dependencies, harder to unit test. Constructor is preferred.

**Q56: What is `@ConfigurationProperties`?**
Binds external properties (application.yml) to a POJO. Type-safe configuration management.

**Q57: Difference between PUT and POST?**
PUT: Idempotent. Update/Replace resource at specific URI.
POST: Not idempotent. Create new resource (server decides URI).

**Q58: What are idempotent APIs?**
Multiple identical requests have same effect as single request (GET, PUT, DELETE). Important for retry logic.

**Q59: How do you handle validation in REST APIs?**
`@Valid` / `@Validated` on DTOs using Jakarta Validation (Hibernate Validator). Return 400 Bad Request on error.

**Q60: What are HTTP status codes you use frequently?**
200 OK, 201 Created, 204 No Content, 400 Bad Req, 401 Unauth, 403 Forbidden, 404 Not Found, 500 Server Error.

**Q61: How do you handle global exception handling?**
`@ControllerAdvice` + `@ExceptionHandler`. Catches exceptions universally and returns consistent JSON error response.

**Q62: What is HATEOAS?**
Hypermedia As The Engine Of Application State. API returns links (next, self, edit) to navigate state transitions dynamically.

**Q63: REST vs SOAP?**
REST: Architectural style, uses HTTP, JSON/XML, lightweight, stateless.
SOAP: Protocol, XML only, strict standards (WS-Security), heavy.

**Q64: Difference between `findById()` and `getOne()`?**
`findById`: Returns `Optional<Entity>` (eager/real DB call).
`getOne` (now `getReferenceById`): Returns Proxy (lazy). Throws exception if access fails.

**Q65: What is the N+1 problem?**
Fetching parent (1 query) then fetching children for *each* parent (N queries). Solved with `JOIN FETCH`.

**Q66: Difference between `EAGER` and `LAZY` fetching?**
EAGER: Load relation immediately with parent.
LAZY: Load relation only when accessed. Default for `@OneToMany`.

**Q67: What is a transactional boundary?**
Scope where transaction is active (e.g., `@Transactional` method). Commit on success, Rollback on unchecked exception.

**Q68: Difference between `save()` and `saveAndFlush()`?**
`save()`: Persist to context (might happen later).
`saveAndFlush()`: Persist and immediately flush changes to DB (sync).

**Q69: What is dirty checking?**
Hibernate feature. Automatically detects changes in managed entities within transaction and updates DB without explicit `save()`.

**Q70: What are entity states?**
Transient (new), Persistent (managed by Session), Detached (session closed), Removed (deleted).

---

## 🏗️ System Design (Questions 71-80)

**Q71: How do you design a scalable REST service?**
Statelessness, Caching, Load Balancing, Asynchronous processing, Horizontal scaling, Database optimizations.

**Q72: How does load balancing work?**
Distributes traffic across multiple servers. Algorithms: Round Robin, Least Connections, IP Hash. Layer 4 (Transport) or Layer 7 (App).

**Q73: What is horizontal vs vertical scaling?**
Horizontal: Adding more machines (Scale Out). Unlimited potential.
Vertical: Adding power to existing machine (Scale Up). Hardware limits.

**Q74: How would you design a URL shortener?**
Hash function (Base62) mapping ID to short string. DB: ID -> Long URL. High read throughput logic.

**Q75: How would you design a rate limiter?**
Token Bucket or Leaky Bucket algorithm. Redis to store counters/timestamps. Middleware to intercept/reject requests.

**Q76: Where would you use caching?**
Browser, CDN, API Gateway, App memory (local), Distributed Cache (Redis), Database buffer pool.

**Q77: Redis vs in-memory cache?**
Redis: Distributed, persistent, shared across instances.
In-memory (Caffeine): Local to JVM, fastest, data lost on restart, memory limited.

**Q78: Database vs cache consistency?**
Cache-Aside (Lazy loading), Write-Through, Write-Back. Hard problem. TTL (Time To Live) is simple solution.

**Q79: Microservices pros and cons?**
Pros: Independent scaling, technology agnostic, fault isolation.
Cons: Complexity, network latency, distributed data/transactions, harder debugging.

**Q80: Monolith vs microservices — when to choose what?**
Monolith: Small team, simple domain, speed to market.
Microservices: Large team, complex domain, need independent scaling/deployment.

---

## 🗄️ SQL & Database (Questions 81-90)

**Q81: Difference between INNER JOIN and LEFT JOIN?**
INNER: Returns rows only when match in *both* tables.
LEFT: Returns all rows from left table, matched (or NULL) from right.

**Q82: What is indexing and how does it work?**
Data structure (B-Tree) to speed up retrieval. Trades write speed/storage for read speed.

**Q83: When can indexes hurt performance?**
On Write/Update/Delete operations (index needs updating). Too many indexes consume storage and slow writes.

**Q84: What is normalization?**
Organizing data to reduce redundancy and improve integrity (1NF, 2NF, 3NF).

**Q85: What is a transaction?**
Logical unit of work. All operations succeed or all fail (Atomic).

**Q86: Explain ACID properties.**
Atomicity, Consistency, Isolation, Durability. Guarantees for reliable transactions.

**Q87: Isolation levels and problems they solve?**
Read Uncommitted, Read Committed (Dirty Read), Repeatable Read (Non-repeatable read), Serializable (Phantom read).

**Q88: Difference between `WHERE` and `HAVING`?**
`WHERE`: Filters rows *before* grouping.
`HAVING`: Filters groups *after* aggregation (GROUP BY).

**Q89: How do you optimize a slow query?**
Explain Plan, add Indexes, avoid `SELECT *`, optimize joins, denormalize if needed, caching.

**Q90: What is an execution plan?**
Roadmap showing how DB engine executes query (scan type, join method, cost). Used for tuning.

---

## 🛠️ Git, Linux, Debugging & Dev Practices (Questions 91-100)

**Q91: Difference between `merge` and `rebase`?**
Merge: Preserves history, creates merge commit.
Rebase: Rewrites history (linear), cleaner log, dangerous on shared branches.

**Q92: What is a pull request?**
Process to propose changes, review code, discuss, and merge into target branch.

**Q93: How do you resolve merge conflicts?**
Identify conflicting files, manually select desired code, remove markers, commit. Standard dev workflow.

**Q94: What is CI/CD?**
Continuous Integration (Auto build/test on commit). Continuous Deployment (Auto release to prod).

**Q95: Common Linux commands you use?**
`ls`, `cd`, `grep`, `curl`, `tail -f`, `ps aux`, `chmod`, `ssh`, `netstat`.

**Q96: How do you check if a port is open?**
`telnet ip port`, `nc -zv ip port`, or `netstat -tuln`.

**Q97: Difference between HTTP and HTTPS?**
HTTPS is HTTP over SSL/TLS. Encrypted communication, secure. Port 80 vs 443.

**Q98: What is TCP vs UDP?**
TCP: Reliable, connection-oriented, ordered (Web, Email).
UDP: Unreliable, connectionless, fast (Streaming, DNS).

**Q99: How do you debug production issues?**
Analyze logs (Splunk/ELK), check metrics (Grafana), reproduce locally, check recent changes/deployments.

**Q100: How do you read and analyze stack traces?**
Find the first line referring to *your* code ("Caused by"). Read exception type and message. Trace execution flow.

---

## 🔐 Security & Authentication (Questions 101-110)

**Q101: What is authentication vs authorization?**
AuthN: "Who are you?" (Login).
AuthZ: "What can you do?" (Permissions/Roles).

**Q102: How does JWT work internally?**
Header (Algo) + Payload (Claims/Data) + Signature (Hash). Base64Encoded. Stateless auth token.

**Q103: Where should JWT be stored on the client and why?**
HttpOnly Cookie (prevents XSS). LocalStorage is vulnerable to XSS.

**Q104: What are common security vulnerabilities in REST APIs?**
Injection (SQL), Broken Auth, Excessive Data Exposure, Lack of Rate Limiting, Weak Logging.

**Q105: What is CORS and how does it work?**
Browser mechanism restricting cross-origin requests. Server must send `Access-Control-Allow-Origin` header.

**Q106: Difference between OAuth2 and JWT?**
OAuth2: Protocol/Framework for authorization (delegated access).
JWT: Token format often used *within* OAuth2 flow.

**Q107: How does Spring Security filter chain work?**
Chain of filters (AuthenticationFilter, AuthorizationFilter, etc.) processing request. Request passes/fails before reaching Controller.

**Q108: What is CSRF and how do you prevent it?**
Cross-Site Request Forgery. Prevent with Anti-CSRF Tokens, SameSite Cookies.

**Q109: How do you secure internal microservice communication?**
mTLS (Mutual TLS), JWT passed internally, Private networking (VPC).

**Q110: What is password hashing and salting?**
Hash: One-way transformation. Salt: Random string added before hashing to prevent Rainbow Table attacks.

---

## 📦 API Versioning, Documentation & Contracts (Questions 111-117)

**Q111: What are different API versioning strategies?**
URI Path (`/v1/users`), Query Param (`?version=1`), Header (`Accept-Version: v1`).

**Q112: Pros and cons of URL vs header-based versioning?**
URL: Easy to test/cache, hurts semantic purity.
Header: Clean URLs, harder to test via browser, caching complexity.

**Q113: What is OpenAPI / Swagger used for?**
Standard for describing REST APIs. Generates documentation (Swagger UI) and client/server code.

**Q114: What is backward compatibility?**
New changes don't break existing clients. Adding fields is safe; removing/renaming is unsafe.

**Q115: What is a contract-first approach?**
Define API spec (OpenAPI/Proto) *before* writing code. Ensures alignment between frontend/backend.

**Q116: How do you deprecate an API safely?**
Announce timelines, use `@Deprecated` header, monitor usage, provide migration path, eventually shut down.

**Q117: How do you handle breaking changes in APIs?**
Create new version (v2). Never break v1. Maintain both until v1 is deprecated.

---

## ⚠️ Exception Handling, Resilience & Stability (Questions 118-125)

**Q118: Difference between checked and unchecked exceptions in real systems?**
Checked: Compile-time forced handling.
Unchecked (Runtime): Logic errors. In modern web apps, prefer Unchecked for cleaner code.

**Q119: Why should you avoid catching `Exception`?**
Catches *everything* (even unexpected runtimes), hiding bugs. Catch specific exceptions instead.

**Q120: What is a circuit breaker pattern?**
Prevents cascading failures. If service fails repeatedly, stop calling it (Open state) for a while, then retry.

**Q121: What is retry vs timeout?**
Timeout: Stop waiting after X seconds.
Retry: Try request again if it failed. Idempotency required.

**Q122: What is bulkhead pattern?**
Isolate resources (thread pools) per service. Failure in Service A doesn't crash Service B.

**Q123: How do you design fault-tolerant services?**
Redundancy, Circuit Breakers, Fallbacks (default values), Rate Limiting, Retries.

**Q124: What is graceful shutdown?**
Stop accepting new requests -> Finish active requests -> Close connections -> Exit.

**Q125: How do you handle partial failures in microservices?**
Return partial response, use fallback content, or return specific error code. Don't crash entire UI.

---

## 📈 Observability, Monitoring & Production Readiness (Questions 126-133)

**Q126: What is observability?**
Measure system internal state by external outputs. "Why is it slow?" vs Monitoring "Is it slow?".

**Q127: Difference between logs, metrics, and traces?**
Logs: Events (error details).
Metrics: Aggregates (CPU, RPS).
Traces: Request flow across services.

**Q128: What are SLAs, SLOs, and SLIs?**
SLI: Indicator (Latency). SLO: Goal (99% < 200ms). SLA: Contract + Penalty.

**Q129: How do you monitor a Spring Boot application?**
Spring Boot Actuator (health, metrics). Prometheus + Grafana.

**Q130: What is distributed tracing?**
Tracking a request across microservices using a Trace ID (Zipkin, Jaeger, OpenTelemetry).

**Q131: What are health checks and readiness probes?**
Liveness: "Am I running?" (Restart if no).
Readiness: "Can I take traffic?" (Load balance if yes).

**Q132: How do you troubleshoot high latency in production?**
Check APM (NewRelic/Datadog), DB slow query logs, GC logs, thread dumps.

**Q133: What is log correlation?**
Tagging all logs for a single request with a unique `correlation-id` to trace flow.

---

## 🧪 Testing (Questions 134-140)

**Q134: Difference between unit, integration, and end-to-end tests?**
Unit: Single function/class (Mocked). Fast.
Integration: Multiple components/DB. Slower.
E2E: Full flow (Browser/API). Slowest.

**Q135: What is mocking and when should you avoid it?**
Simulating dependencies. Avoid mocking value objects or simple logic; mock external I/O (DB, APIs).

**Q136: How do you test REST controllers?**
`MockMvc` / `@WebMvcTest`. Simulate HTTP requests without starting full server.

**Q137: What is `@SpringBootTest` vs `@WebMvcTest`?**
`@SpringBootTest`: Loads full context (Integration).
`@WebMvcTest`: Loads only web layer (Unit-ish).

**Q138: How do you test database interactions?**
`@DataJpaTest` with H2 (in-memory) or TestContainers (Real DB in Docker).

**Q139: What is test coverage and why it can be misleading?**
% of code executed by tests. High coverage != High quality tests (can assert nothing).

**Q140: What makes a good test?**
Fast, Deterministic (no flake), Independent, Descriptive, tests Behavior not Implementation.

---

## ⚡ Performance & Scalability (Questions 141-148)

**Q141: How do you identify performance bottlenecks?**
Profiling (VisualVM), APM tools, Load testing (JMeter), DB analysis.

**Q142: Difference between throughput and latency?**
Throughput: Requests per second (Capacity).
Latency: Time per request (Speed).

**Q143: How do you handle high traffic spikes?**
Caching, Auto-scaling, Queue-based load leveling, Rate limiting, CDN.

**Q144: What is backpressure?**
Consumer signals Producer to slow down when overwhelmed (Reactive Streams).

**Q145: When would you use async processing?**
Long-running tasks (Email, Report generation). Enhances responsiveness.

**Q146: How do you tune JVM for performance?**
Heap size optimization, GC selection (G1GC/ZGC), Thread pool tuning.

**Q147: How do you design for read-heavy systems?**
Caching (Redis), Read Replicas (DB), CDN, Denormalization.

**Q148: What are hot keys in Redis?**
Keys accessed too frequently causing uneven load. specific caching or sharding required.

---

## 🧩 Configuration & Environment Management (Questions 149-153)

**Q149: How do you manage configs across environments?**
Profiles (`application-dev.yml`, `prod.yml`), Environment Variables, Config Server.

**Q150: What is 12-factor app methodology?**
Best practices for cloud-native apps (Config in env, Stateless, Disposability, Logs as streams).

**Q151: Difference between application.yml and bootstrap.yml?**
Bootstrap: Loaded *before* application context (Cloud Config). Application: Standard config.

**Q152: How do you handle secrets securely?**
Vault, AWS Secrets Manager, Kubernetes Secrets. Never in git code.

**Q153: What is feature flagging?**
Toggling functionality at runtime without deploying code. Enables Canary releases/Dark launches.

---

## 🧱 Data Consistency & Distributed Systems Basics (Questions 154-160)

**Q154: What is CAP theorem?**
Store can strictly provide only 2 of 3: Consistency, Availability, Partition Tolerance. CP or AP.

**Q155: Difference between strong and eventual consistency?**
Strong: Read sees latest write immediately.
Eventual: Read sees latest write *eventually*.

**Q156: What is idempotency in distributed systems?**
Operation applied multiple times has same result as applying once. Essential for messaging/retries.

**Q157: How do you design idempotent APIs?**
Client sends unique ID. Server checks if ID processed. If yes, return stored response.

**Q158: What is saga pattern?**
Managing distributed transactions. Sequence of local transactions. If one fails, execute compensating transactions meant to undo changes.

**Q159: Two-phase commit vs saga?**
2PC: Strong consistency (locking), slow. Saga: Eventual consistency (compensating actions), scalable.

**Q160: How do you handle duplicate messages?**
Idempotent consumer. Track processed Message IDs in DB/Redis.

---

## 🧠 Behavioral / Real-World Engineering Questions (Questions 161-170)

**Q161: Describe a production issue you handled end-to-end.**
(STAR Method) Situation: System slow. Action: Analyzed thread dump, found DB lock. Result: Fixed index, improved speed 50%.

**Q162: How do you prioritize bug fixes vs new features?**
Severity/Impact first. Critical bugs > Features > Minor bugs. Balance tech debt.

**Q163: How do you handle technical debt?**
Allocate % of sprint time (20%). Refactor continuously (Boy Scout Rule). Document it.

**Q164: How do you disagree with a design decision?**
Data/Facts over opinion. Propose alternatives with trade-offs. Commit to final decision.

**Q165: How do you estimate backend work?**
Break down tasks. Complexity points (Fibonacci). Add buffer for testing/unknowns.

**Q166: How do you do code reviews effectively?**
Check logic, not just syntax. Be constructive. Ask questions. Ensure tests exist.

**Q167: How do you handle on-call incidents?**
Acknowledge, Triage (Severity), Mitigate (Restart/Rollback), Investigate (Root Cause), Post-Mortem.

**Q168: What trade-offs have you made for delivery speed?**
Accepted technical debt (hardcoding) to meet deadline, with plan to refactor immediately after.

**Q169: How do you ensure code quality under deadlines?**
CI/CD pipelines, Critical tests must pass. Do not compromise on Core logic reviews.

**Q170: What would you improve in your last project?**
Better observability, more automated testing, or clearer documentation.
