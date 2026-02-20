# Expert Level Java Interview Questions

## From 11 Concurrency Practice
# 11. Concurrency (Practice)

**Q: Difference between Thread and Runnable?**
> "In Java, you have two main ways to create a thread.
>
> One is to **extend the Thread class**. You write `class MyThread extends Thread`. The problem is, Java only allows single inheritance. If you extend `Thread`, you can't extend anything else.
>
> The better way is to **implement the Runnable interface**. You write `class MyTask implements Runnable`. This leaves your class free to extend another class if needed. You then pass your `Runnable` instance to a `Thread` object to run it.
>
> So, rule of thumb: Always implement `Runnable` (or `Callable`) unless you are actually modifying the behavior of the Thread class itself."

**Indepth:**
> **Virtual Threads**: In Java 21+, `Thread.ofVirtual().start(runnable)` creates lightweight threads mapped to OS threads by the JVM. Implementing `Runnable` makes upgrading to Virtual Threads trivial.


---

**Q: What is a deadlock? How do you create one?**
> "A deadlock is a standoff. It happens when two threads are stuck forever, each waiting for the other to release a lock.
>
> Imagine Thread A holds 'Key 1' and wants 'Key 2'.
> At the same time, Thread B holds 'Key 2' and wants 'Key 1'.
> neither will give up what they have, so they wait forever.
>
> Creating one is easy: Just have two threads and two locks. Make Thread A lock Lock1 then wait for Lock2. Make Thread B lock Lock2 then wait for Lock1. Run them together, and your app freezes."

**Indepth:**
> **Prevention**: Avoid nested locks. If you must use multiple locks, always acquire them in the *same order* (e.g., always lock Resource A before Resource B) to prevent cycles.


---

**Q: What is ExecutorService?**
> "Managing raw threads (creating them, starting them) is expensive and error-prone.
>
> **ExecutorService** is a higher-level framework introduced in Java 5 to handle this. It manages a **pool of threads** for you.
>
> Instead of saying `new Thread(task).start()`, you say `executor.submit(task)`.
> It recycles threads (saving memory), handles scheduling, and gives you `Future` objects to track progress. You should almost always use this instead of raw Threads."

**Indepth:**
> **Shutdown**: Don't forget to shut it down! `executor.shutdown()` stops accepting new tasks. If you don't call it, your app might never exit because the pool threads are still alive.


---

**Q: Difference between Callable and Runnable?**
> "They are both tasks meant for threads, but they have key differences.
>
> **Runnable** is the old one (Java 1.0). Its method is `public void run()`. It doesn't return anything, and it *cannot* throw a checked exception to the caller.
>
> **Callable** is the new one (Java 5). Its method is `public T call()`. It *returns a result* (T), and it *can* throw an exception.
>
> If you need a value back from your thread (like a calculation result), use `Callable`."

**Indepth:**
> **Future**: `submit(callable)` returns a `Future<T>`. Calling `future.get()` blocks the current thread until the result is ready (or throws an exception).


---

**Q: What are atomic classes (AtomicInteger)?**
> "In multi-threading, simple operations like `count++` are not safe. It actually involves three steps: read, increment, write. If two threads do it at the same time, you lose data.
>
> **Atomic Classes** (like `AtomicInteger`) provide a way to perform these operations safely *without* using heavy locks (`synchronized`).
>
> They use low-level CPU instructions (CAS - Compare-And-Swap) to ensure that `incrementAndGet()` happens atomically. It's much faster than using synchronization for simple counters."

**Indepth:**
> **Non-blocking**: Synchronized blocks put threads to sleep (blocking). Atomics spin or use hardware instructions (non-blocking), which is much more scalable under high contention.


---

**Q: Thread-safe Singleton (Double-Checked Locking)**
> "If you are creating a Singleton lazily (creating it only when asked), you have to be careful.
>
> If two threads call `getInstance()` at the same time, and the instance is null, both might create a new object!
>
> To fix this efficiently, we use **Double-Checked Locking**:
> 1.  Check if instance is null (no locking, fast).
> 2.  If null, enter a `synchronized` block.
> 3.  Check if instance is null *again* (just in case another thread beat us to it while we were waiting to enter the block).
> 4.  Create the instance.
>
> Also, the instance variable must be marked `volatile` to prevent instruction reordering issues."

**Indepth:**
> **Volatile**: Without `volatile`, the "Create instance" step (allocate memory, init variables, assign reference) can be reordered. Another thread might see a non-null but *partially initialized* object and crash.


## From 12 Extra Concepts Practice
# 12. Extra Concepts (Practice)

**Q: When would you use Optional, and when should you avoid it?**
> "**Optional** is a container object that might (or might not) contain a value. It was introduced in Java 8 to avoid `NullPointerException`.
>
> You **should** use it as a *return type* for methods that might not find a result. Like `findUserById()`. It forces the caller to think: 'What if the user isn't found?' and handle it gracefully using methods like `.orElse()` or `.ifPresent()`.
>
> You **should generally avoid** using it for:
> 1.  Field variables (it's not serializable).
> 2.  Method parameters (it just makes calling the method annoying).
> 3.  Collections (never put Optional in a List, just have an empty List)."

**Indepth:**
> **Performance**: `Optional` is an object. Creating it adds overhead. Using it deeply in tight loops or for every single field in a massive data structure will hurt performance and GC.


---

**Q: Why are generics invariant in Java?**
> "This is a tricky one. Invariance means `List<String>` is **not** a subtype of `List<Object>`.
>
> Why? Because of type safety.
> If Java allowed you to treat a `List<String>` as a `List<Object>`, you could add an `Integer` to it!
>
> ```java
> List<String> strings = new ArrayList<>();
> List<Object> objects = strings; // If this were allowed...
> objects.add(10); // You just put an int into a list of strings!
> ```
> When you try to read that 'int' back as a String, your program would crash. So Java prevents this at compile time by making generics invariant."

**Indepth:**
> **Covariance**: Generics *can* be covariant using wildcards (`List<? extends Number>`). This allows reading (you know everything inside is at least a Number) but prevents writing (you don't know if it's meant to hold Integers or Doubles).


---

**Q: Strategy Pattern real-world use case?**
> "The **Strategy Pattern** is about swapping algorithms at runtime.
>
> Think of a Payment System on an e-commerce site. You have a `pay()` method.
> But the user might want to pay with **Credit Card**, **PayPal**, or **Bitcoin**.
>
> Instead of writing one giant `if-else` block inside the `pay()` method, you define a `PaymentStrategy` interface. Then you create classes `CreditCardStrategy`, `PayPalStrategy`, etc.
>
> You pass the chosen strategy to the payment processor. This makes it super easy to add a new payment method later (like Apple Pay) without touching the existing code."

**Indepth:**
> **Open/Closed Principle**: This is the textbook example of OCP. Classes should be open for extension (adding new Strategies) but closed for modification (not touching the `pay()` method).


---

**Q: Abstract Factory vs Factory Method?**
> "They both create objects, but strictly speaking:
>
> **Factory Method** uses *inheritance*. You have a method `createAnimal()` in a base class, and subclasses override it to return a `Dog` or `Cat`. It creates *one* product.
>
> **Abstract Factory** uses *composition*. It's a factory *of factories*. It creates *families* of related products.
> Like a `GUIFactory` that creates a `Button`, `Checkbox`, and `Scrollbar`. You might have a `WindowsFactory` that returns Windows-style buttons and checkboxes, and a `MacFactory` that returns Mac-style ones. You ensure that all components match the same theme."

**Indepth:**
> **Dependency Inversion**: Abstract Factory allows the client code to be completely decoupled from concrete classes. It only knows about the interfaces (`Button`, `Window`). This makes cross-platform UI toolkits possible.


---

**Q: StackOverflowError Simulation**
> "A `StackOverflowError` happens when the call stack gets too deep, usually due to **infinite recursion**.
>
> To simulate it, just write a method that calls itself without a breaking condition:
>
> ```java
> public void recursive() {
>     recursive();
> }
> ```
> Run that, and boom—StackOverflow."

**Indepth:**
> **Tail Call Optimization**: Java does *not* support tail call optimization (yet). So even if the recursive call is the very last thing, it still consumes a stack frame.


---

**Q: OutOfMemoryError Simulation**
> "An `OutOfMemoryError` (OOM) happens when the **Heap** is full.
>
> To simulate it, just keep creating objects and holding onto them so the Garbage Collector can't delete them.
>
> ```java
> List<byte[]> list = new ArrayList<>();
> while (true) {
>     list.add(new byte[1024 * 1024]); // Add 1MB chunks continuously
> }
> ```
> Eventually, the heap fills up, and you crash."

**Indepth:**
> **Analysis**: When OOM happens, you need a Heap Dump. Tools like Eclipse MAT or VisualVM can analyze this dump to find the "Leak Suspects" (which objects are consuming the most RAM).


## From 13 Advanced Concurrency JVM Practice
# 13. Advanced Concurrency and JVM (Practice)

**Q: How to use CompletableFuture for asynchronous programming?**
> "**CompletableFuture** (Java 8) is the modern way to write non-blocking async code. It allows you to chain tasks together like a pipeline.
>
> You start a task with `CompletableFuture.supplyAsync(() -> performTask())`.
> You can then chain what happens next: `.thenApply(result -> process(result))` or `.thenAccept(final -> print(final))`.
>
> The main thread doesn't wait. It continues execution. You can combine multiple futures, waiting for all of them to finish (`allOf`) or just the fastest one (`anyOf`). It's much more powerful than the old `Future` interface."

**Indepth:**
> **Exception Handling**: Standard `Future.get()` throws messy checked exceptions. `CompletableFuture` handles exceptions gracefully inside the pipeline using `.exceptionally()` or `.handle()`, keeping the flow clean.


---

**Q: Difference between synchronized and ReentrantLock?**
> "Both handle locking, but **ReentrantLock** has more features.
>
> **synchronized** is a keyword. It's implicit. You enter the block, you get the lock; you leave, you release it. It's cleaner but rigid. If a thread waits for a synchronized lock, it waits forever.
>
> **ReentrantLock** is a class. You manually call `lock()` and `unlock()`. This gives you power:
> 1.  **tryLock()**: 'Attempt to get the lock, but if it's busy, give up immediately' (or wait for a timeout).
> 2.  **Fairness**: You can set it to 'fair' mode, granting locks in First-Come-First-Served order (preventing starvation).
> 3.  **Interruptible**: A waiting thread can be interrupted."

**Indepth:**
> **Condition**: `ReentrantLock` allows multiple `Condition` objects (like `wait()`/`notify()` but with multiple waiting rooms). You can signal *specific* threads waiting for "Not Empty" vs "Not Full", reducing useless wake-ups.


---

**Q: What is ThreadLocal and when to use it?**
> "**ThreadLocal** allows you to create variables that can only be read and written by the *same thread*. It's like a private pocket for each thread.
>
> Imagine a web server handling 100 requests. Each request is a thread. You want a transaction ID for strictly that request/thread.
> If you save it in a static variable, all threads would overwrite each other.
> If you use `ThreadLocal`, each thread sees its own unique value.
>
> Use it sparingly though—specifically for things like user sessions or transaction contexts. And always `remove()` it when done to prevent memory leaks in thread pools."

**Indepth:**
> **Internals**: It's implemented as a `ThreadLocalMap` inside the `Thread` class itself. The key is the `ThreadLocal` object (weak reference), and the value is your data.


---

**Q: How does a ClassLoader work? (Delegation Model)**
> "When you ask Java to load a class (`MyClass`), it doesn't just look in one place. It follows the **Delegation Hierarchy**:
>
> 1.  It asks the **Application ClassLoader** (your classpath).
> 2.  That delegates to the **Extension/Platform ClassLoader** (lib/ext).
> 3.  That delegates to the **Bootstrap ClassLoader** (core Java libs like `String`, `Object`).
>
> The loading actually happens in reverse order. The Bootstrap tries first. If it can't find it, it goes back down the chain.
> This ensures security: you can't trick Java into loading your own hacked version of `java.lang.String` because the Bootstrap loader will always find the real one first."

**Indepth:**
> **Class Uniqueness**: A class is identified by its **Fully Qualified Name + ClassLoader**. You can have two classes with the precise same name `com.myapp.User` loaded by two different ClassLoaders, and JVM treats them as completely different types!


---

**Q: What are Virtual Threads (Java 21+)?**
> "This is Project Loom. It's revolutionary.
>
> Traditionally, one Java thread `==` one OS thread. OS threads are heavy (consume 1MB RAM) and limited (you can only have a few thousand).
>
> **Virtual Threads** are lightweight threads managed by the **JVM**, not the OS. You can create *millions* of them.
> When a virtual thread waits for I/O (like a database call), the JVM unmounts it and puts another virtual thread on the CPU.
> This allows high-throughput server applications written in a simple, synchronous style without needing complex async callbacks."

**Indepth:**
> **Adoption**: Virtual Threads are designed to work with existing synchronous code (blocking I/O). They do *not* improve CPU-bound tasks (number crunching), only I/O-bound tasks (web servers).


---

**Q: What is the ReadWriteLock?**
> "A standard lock is exclusive—only one thread can enter, period.
>
> A **ReadWriteLock** is smarter. It says: 'Multiple threads can read at the same time, as long as nobody is writing.'
>
> It maintains a pair of locks: a **Read Lock** (shared) and a **Write Lock** (exclusive).
> If you have a system with lots of reads but few writes (like a cache), this drastically improves performance compared to a standard synchronized block."

**Indepth:**
> **Starvation**: A risk with ReadWrite locks is "Writer Starvation". If there are constant readers, the writer might never get the lock. Modern implementations (stamped locks) try to mitigate this.


---

**Q: SOLID - Single Responsibility Principle (SRP) Example**
> "**SRP** states: 'A class should have only one reason to change.'
>
> **Violating SRP**:
> A `User` class that handles user data AND saves the user to the database AND sends a welcome email.
> If you change the database, you touch the User class. If you change the email provider, you touch the User class. Bad.
>
> **Following SRP**:
> Split it into multiple classes:
> 1.  `User`: Just holds data (POJO).
> 2.  `UserRepository`: Handles database operations.
> 3.  `EmailService`: Handles sending emails.
>
> Now, each class focuses on one job. It’s easier to test and maintain."

**Indepth:**
> **Microservices**: SRP applied at the architecture level leads to Microservices. Each service does one thing well (User Service, Email Service, Payment Service).


## From 32 Spring Core Monitoring WebFlux
# 32. Spring Core, Monitoring & WebFlux

**Q: @Component vs @Service vs @Repository**
> "Technically, they are all the same. `@Service` and `@Repository` are just aliases for `@Component`.
>
> However, we use them for **Semantics** and **Exception Translation**:
> *   `@Repository`: Tells Spring 'This class talks to a DB'. It also automatically translates low-level SQL exceptions into cleaner Spring DataAccessExceptions.
> *   `@Service`: Tells devs 'Business Logic lives here'.
> *   `@Component`: Generic utility classes."

**Indepth:**
> **AOP**: `@Service` and `@Repository` can be targeted by Aspect-Oriented Programming pointcuts. For example, `@Transactional` attributes usually apply to the Service layer, while Exception Translation (translating SQL errors to Spring ones) only happens on `@Repository`.


---

**Q: @Autowired vs @Qualifier**
> "**@Autowired** by default looks for a bean by **Type**.
> If you have an interface `PaymentService` and two implementations (`CreditCardService`, `PayPalService`), Spring throws `NoUniqueBeanDefinitionException`.
>
> **@Qualifier** tells Spring to look by **Name**.
> `@Autowired @Qualifier("payPalService")` resolves the ambiguity."

**Indepth:**
> **Primary**: Alternatively, you can simplify injection by using `@Primary` on one of the implementations. This tells Spring "if there's ambiguity and no qualifier is present, use this bean by default."


---

**Q: @Value vs @ConfigurationProperties**
> "**@Value** is for injecting single values.
> `@Value("${app.timeout}") int timeout;`
> Good for quick, one-off properties.
>
> "**@ConfigurationProperties** is for grouping properties.
> It maps a hierarchical structure (`server.port`, `server.address`) to a Java POJO.
> It is type-safe, validates fields, and supports loose binding (kebab-case to camelCase). Always prefer this for complex configs."

**Indepth:**
> **Relaxed Binding**: `@ConfigurationProperties` supports relaxed binding. `my.app-name` in properties matches `myAppName`, `my_app_name`, and `my.app.name` in Java. `@Value` requires exact string matches.


---

**Q: Constructor Injection**
> "Stop using Field Injection (`@Autowired private Repo repo`). It makes testing hard (you have to use Reflection to set the repo).
>
> **Constructor Injection** is the standard.
> ```java
> private final Repo repo;
> public Service(Repo repo) { this.repo = repo; }
> ```
> It forces dependencies to be provided, ensures immutability (`final`), and makes Unit Testing trivial (just pass a mock in the constructor)."

**Indepth:**
> **Circular Dependencies**: Constructor injection prevents circular dependencies at compile-time/start-time (Bean A needs B, B needs A). Spring throws `BeanCurrentlyInCreationException` immediately, forcing you to refactor your bad design.


---

**Q: Bean Scopes (Singleton vs Prototype)**
> "**Singleton** (Default): Spring creates **one** instance of the bean per container. Shared by everyone. Stateless services should be Singletons.
>
> "**Prototype**: Spring creates a **new** instance every time you ask for it (`context.getBean()`). Stateful beans (like a 'ShoppingCart' or 'UserSession') might use this, though `SessionScope` is usually better for web apps."

**Indepth:**
> **Proxy**: If you inject a Prototype bean into a Singleton bean, the Prototype is created *only once* (when the Singleton is created). To get a new Prototype every time, you must use `ObjectFactory<MyPrototype>` or `Lookup` method injection.


---

**Q: Actuator Health Endpoint**
> "`/actuator/health` provides a status check.
> By default, it just says `{"status": "UP"}`.
>
> If you enable details (`management.endpoint.health.show-details=always`), it checks:
> *   Disk Space
> *   Database Connection
> *   Message Broker Connectivity
> If **any** of these down, the overall status becomes `DOWN` (503 Service Unavailable)."

**Indepth:**
> **Custom**: You can write your own `HealthIndicator`. For example, checking if a critical 3rd party API is reachable. Implement the interface and return `Health.up()` or `Health.down().withDetail("reason", "timeout")`.


---

**Q: Micrometer & Prometheus**
> "**Micrometer** is like SLF4J but for metrics.
> It's a facade. You write code against Micrometer (`Counter.increment()`), and it translates that to whatever backend you use (Prometheus, Datadog, NewRelic).
>
> **Prometheus** scrapes these metrics. You expose `/actuator/prometheus`, and Prometheus comes and 'pulls' the data every 15 seconds."

**Indepth:**
> **Dimensionality**: Micrometer supports tags (dimensions). Instead of just `http_requests_total`, you track `http_requests_total{method="GET", status="200"}`. This allows powerful querying like "Show me only 500 errors on POST requests".


---

**Q: Application Monitoring (Memory/CPU)**
> "Use the `/actuator/metrics` endpoint.
> *   `jvm.memory.used`
> *   `system.cpu.usage`
> *   `hikaricp.connections.active`
>
> You don't usually read the JSON manually. You connect Grafana to visualize 'CPU Spikes' or 'Memory Leaks' over time."

**Indepth:**
> **Alerting**: Monitoring is useless without alerting. Set up Prometheus/Grafana alerts for "High Memory Usage (> 85%)" or "High Error Rate (> 1%)". Don't wait for users to complain.


---

**Q: Spring WebFlux vs Spring MVC**
> "**Spring MVC**: Thread-per-request.
> 1 Request = 1 Thread. If the thread waits for DB, it sits idle (Blocked).
> Good for standard CRUD apps.
>
> "**Spring WebFlux**: Event-Loop based (like Node.js).
> Small number of threads handle thousands of concurrent requests. If waiting for DB, the thread goes to work on another request.
> Returns **Mono** (0-1 item) or **Flux** (0-N items).
> Good for High-Scale Streaming apps."

**Indepth:**
> **Backpressure**: WebFlux supports **Backpressure**. If the client (Consumer) is slow, it tells the Server (Producer) to slow down ("I can only handle 5 items right now"). Spring MVC just overwhelms the client.


---

**Q: Mono vs Flux**
> "In Reactive Programming, we don't return `List<User>`.
>
> *   **Mono<T>**: A wrapper for zero or one item. 'I promise to give you a User (or error) in the future'.
> *   **Flux<T>**: A wrapper for zero to N items. 'I will stream Users to you as they arrive'.
>
> You subscribe to them to get the data."

**Indepth:**
> **Cold vs Hot**: `Flux` is "Cold" by default. Nothing happens until you subscribe. If you have a DB call in a Flux but nobody subscribes, the DB call never executes.


## From 33 Messaging Kafka Docker Kubernetes
# 33. Messaging (Kafka) & Containerization (Docker/Kubernetes)

**Q: Server-Sent Events (SSE)**
> "SSE is a one-way communication channel from Server to Client.
> Unlike WebSockets (which are bidirectional), SSE is simpler.
>
> In Spring Boot:
> Return `Flux<ServerSentEvent<String>>`.
> The browser keeps the connection open, and you can push stock updates or notifications in real-time."

**Indepth:**
> **Reconnection**: Standard HTTP requests don't auto-reconnect. SSE has built-in reconnection logic. If the connection drops, the browser automatically tries to reconnect, sending the "Last-Event-ID" so the server can resume from where it left off.


---

**Q: Kafka Producer/Consumer (Spring Boot)**
> "It's all about `KafkaTemplate` and `@KafkaListener`.
> 1.  **Publishing**: `kafkaTemplate.send("topic_name", "message")`.
> 2.  **Consuming**:
>     ```java
>     @KafkaListener(topics = "topic_name", groupId = "my-group")
>     public void listen(String message) {
>         // Process message
>     }
>     ```"

**Indepth:**
> **Serialization**: Spring Boot uses `StringSerializer` by default for keys and values. In production, you'll likely switch the Value serializer to `JsonSerializer` (Jackson) to send complex objects easily.


---

**Q: Kafka Error Handling (Retries/DLQ)**
> "What if processing a message fails?
> 1.  **Retry**: Configure a `DefaultErrorHandler` with a `FixedBackOff`. It retries 3 times.
> 2.  **Dead Letter Queue (DLQ)**: If it still fails, send the message to a separate topic (`orders-dlt`). You can inspect these later manually."

**Indepth:**
> **Non-Blocking**: By default, retries might block the consumer thread, stopping it from processing *other* messages. **Non-Blocking Retries** (using `@RetryableTopic`) publish the failed message to a delay-queue topic, freeing up the consumer immediately.


---

**Q: WebClient vs RestTemplate**
> "**RestTemplate** is blocking. It waits for the response. Deprecated (in maintenance mode).
>
> "**WebClient** is non-blocking (Reactive).
> It uses Netty. It allows you to make parallel calls easily:
> `Mono.zip(callA(), callB())`.
> Even if you use blocking Spring MVC, you should start using WebClient for external API calls."

**Indepth:**
> **Resources**: `RestTemplate` creates a new Thread for every request. If you call 100 external APIs, you potentially block 100 threads. `WebClient` can handle 100 requests with just 1 thread using Non-Blocking IO.


---

**Q: Dockerizing Spring Boot**
> "The simplest way:
> 1.  Build the jar: `mvn clean package`.
> 2.  Write a `Dockerfile`:
>     ```dockerfile
>     FROM openjdk:17-alpine
>     COPY target/app.jar app.jar
>     ENTRYPOINT ["java", "-jar", "app.jar"]
>     ```
> 3.  `docker build -t my-app .`"

**Indepth:**
> **Multi-Stage**: Use Multi-Stage Docker builds to optimize image size. Stage 1 (Maven) builds the jar (requires 500MB of deps). Stage 2 (JRE) only copies the final jar (requires 50MB). The final image is tiny.


---

**Q: Layered JARs (Optimization)**
> "A standard Spring Boot JAR is huge (App Code + 50MB of Libraries).
> If you change one line of code, Docker has to re-push the whole 50MB layer.
>
> **Layered JARs** separate them:
> Layer 1: Dependencies (rarely change).
> Layer 2: Your Code (changes often).
> Docker reuses Layer 1 from cache and only pushes Layer 2. Faster builds, faster deployments."

**Indepth:**
> **Cache**: The `spring-boot-maven-plugin` has a `layers` configuration. When enabled, it splits `dependencies`, `spring-boot-loader`, `snapshot-dependencies`, and `application` classes into separate folders in the docker image specifically for caching.


---

**Q: Jib (Google Tool)**
> "Jib allows you to build Docker images **without** a Docker daemon and **without** a Dockerfile.
>
> You just add the `jib-maven-plugin` plugin to your pom.xml.
> Run `mvn jib:build`.
> It analyzes your project, intelligently layers it, and pushes it directly to a registry (like Docker Hub)."

**Indepth:**
> **Reproducibility**: Jib separates the application from the OS. It doesn't use a Dockerfile, so "It works on my machine" issues related to different base OS installations are minimized.


---

**Q: Spring Boot on Kubernetes (K8s)**
> "Spring Boot runs naturally on K8s.
>
> **Configuration**: Use `ConfigMaps` and `Secrets` mapped to environment variables.
> **Health**: K8s uses 'Probes' to check if your app is alive. Map them to Actuator:
> *   Liveness Probe -> `/actuator/health/liveness`
> *   Readiness Probe -> `/actuator/health/readiness`"

**Indepth:**
> **Graceful Shutdown**: Configure `server.shutdown=graceful`. When K8s kills a pod, Spring Boot will stop accepting new requests but will wait (e.g., 30s) for existing requests to finish processing before shutting down the JVM.


---

**Q: Rolling Updates**
> "K8s handles this. You don't do it in Spring.
> You tell K8s: 'Update to version 2.0'.
> K8s spins up a v2 pod. Waits for the **Readiness Probe** to pass. Then kills a v1 pod. Then repeats.
> Zero downtime."

**Indepth:**
> **Deployment Strategies**: Beyond basic rolling updates, K8s supports "Blue-Green" (spin up full v2 parallel to v1, then switch traffic) and "Canary" (send 5% of traffic to v2 to test it) deployments.


---

**Q: Idempotency in Consumers**
> "In Kafka, you might receive the same message twice (At-Least-Once delivery).
> Your consumer **must** be idempotent.
>
> Strategy:
> 1.  Use a unique `message_id`.
> 2.  Maintain a 'Processed IDs' table in your DB.
> 3.  Check: `if (repo.exists(id)) return;` before processing."

**Indepth:**
> **Transactionality**: For exactly-once semantics inside the Kafka ecosystem, you can use Kafka Transactions (`producer.send` + `consumer.commit` are atomic). But for external side-effects (DB writes), Idempotency keys are safer.


---

**Q: Avro/Protobuf (Schema Registry)**
> "Sending raw JSON is wasteful and error-prone.
> **Avro** is a binary format. It's smaller and faster.
>
> You use a **Schema Registry**. The Producer checks the schema ID, sends binary data. The Consumer downloads the schema and deserializes it. It ensures structural compatibility (Contract Testing) automatically."

**Indepth:**
> **Evolution**: Schema Registry allows Schema Evolution. You can add a nullable field to your user object, and old consumers (that don't know about the field) will simply ignore it, ensuring backward compatibility without breaking the system.


## From 36 Spring Boot NoSQL Integration Cloud
# 36. Spring Boot (NoSQL, Integration & Cloud)

**Q: Redis with Spring Boot**
> "Redis is typically used for caching, but you can use it as a primary store.
> You use `RedisTemplate`.
>
> It’s a Key-Value store. So you treat it like a giant, persistent `HashMap` in the cloud.
> `template.opsForValue().set("user:1", jsonString);`
> It's incredibly fast (sub-millisecond) but data must fit in memory."

**Indepth:**
> **Serializers**: `JdkSerializationRedisSerializer` is default but bad (binary blobs). Use `Jackson2JsonRedisSerializer` so data is readable JSON in Redis CLI. `StringRedisTemplate` is a pre-configured template just for String keys/values.


---

**Q: Redis Key Expiration**
> "One of the best features of Redis is that data can self-destruct.
> When you save a key, you set a TTL (Time To Live).
>
> `template.opsForValue().set("otp:12345", "8732", Duration.ofMinutes(5));`
>
> After 5 minutes, Redis automatically deletes it. This is perfect for OTPs, User Sessions, and temporary Cache entries."

**Indepth:**
> **Eviction**: What happens when Redis is full? It deletes keys. You configure the eviction policy. `allkeys-lru` deletes any key. `volatile-lru` only deletes keys with an expiry set.


---

**Q: Spring Integration DSL**
> "Spring Integration is about connecting systems (Files, FTP, Queues).
> The DSL (Domain Specific Language) allows you to define these 'Pipelines' in Java code so it reads like a story:
>
> ```java
> return IntegrationFlow.from("inputChannel")
>     .filter("payload.amount > 100")
>     .transform(Transformers.toJson())
>     .handle(Amqp.outboundAdapter(rabbitTemplate))
>     .get();
> ```
> It says: Take input -> Filter high amounts -> Convert to JSON -> Send to RabbitMQ."

**Indepth:**
> **Channels**: Channels are the pipes. `DirectChannel` is a method call (synchronous, same thread). `QueueChannel` is a buffer (asynchronous, different thread). The DSL hides this complexity.


---

**Q: File Polling (Spring Integration)**
> "If you need to watch a folder for new PDF files and process them.
> You define an `InboundFileAdapter`.
>
> It polls the directory every 5 seconds. If it finds a new file, it locks it, passes it to your processing method, and then moves it to a 'processed' folder automatically. No manual `While(true)` loops needed."

**Indepth:**
> **Idempotency**: `AcceptOnceFileListFilter`. How do you prevent processing the same file twice? You need a filter. Be careful: standard filters keep state in memory. If you restart the app, it might process old files again unless you use a persistent usage store (MetadataStore).


---

**Q: Spring Cloud Sleuth & Zipkin**
> "**Sleuth** adds a unique ID (Trace ID) to your logs.
> When Service A calls Service B, Sleuth passes that ID in the HTTP Headers.
>
> **Zipkin** is the UI. It collects all these logs and draws a timeline: 'Request blocked for 200ms in Service A, then took 50ms in Service B'. It's essential for debugging microservices latency."

**Indepth:**
> **Sampling**: `probability`. You don't want to trace 100% of requests in production (performance overhead). You set `spring.sleuth.sampler.probability=0.1` (10%). But for errors, you always want the trace.


---

**Q: Service Discovery (Eureka)**
> "In the cloud, IP addresses change all the time. You can't hardcode `http://192.168.1.50`.
>
> **Eureka** is a phonebook.
> 1.  Service A starts up and says: 'I am Service A, my IP is X'. (Registration)
> 2.  Service B asks Eureka: 'Where is Service A?'.
> 3.  Eureka replies: 'It's at IP X'.
> Service B then calls Service A directly."

**Indepth:**
> **Self Preservation**: If Eureka stops receiving heartbeats from *many* instances at once (e.g., network partition), it stops expiring them. It assumes the network is down, not the instances. This prevents mass accidental shutdowns.


---

**Q: Feign Client**
> "Stop using `RestTemplate` for calling other microservices. It's verbose.
>
> **Feign** is declarative. You just write an interface:
> ```java
> @FeignClient(name = "inventory-service")
> public interface Inventory {
>     @GetMapping("/items/{id}")
>     Item getItem(@PathVariable String id);
> }
> ```
> Spring generates the implementation at runtime. If you call `getItem("5")`, it automatically calls `http://inventory-service/items/5` via Eureka."

**Indepth:**
> **Error Decoding**: Feign throws `FeignException` by default. You implement a custom `ErrorDecoder` to translate 404s/500s from the remote service into your own domain exceptions (`InventoryNotFoundException`).


---

**Q: Circuit Breaker (Resilience4j)**
> "If the 'Inventory Service' goes down, you don't want the 'Order Service' to hang and crash too (Cascading Failure).
>
> A **Circuit Breaker** wraps the call.
> *   **Closed**: Normal operation.
> *   **Open**: Too many failures (result > 50%). Stick fails immediately (Fast Fail) without waiting for timeout.
> *   **Half-Open**: Let one request through to see if it's fixed.
>
> It keeps your system responsive even when dependencies fail."

**Indepth:**
> **Bulkhead**: A Circuit Breaker stops all calls when the failure rate is high. A **Bulkhead** limits concurrency. "Max 10 concurrent calls to Inventory Service". If the 11th comes, it's rejected immediately. This prevents one slow service from exhausting all your Tomcat threads.


---

**Q: Buildpacks (Cloud Native)**
> "You don't need a Dockerfile anymore.
> `mvn spring-boot:build-image`
>
> This uses **Cloud Native Buildpacks**. It detects 'Oh, this is a Java 17 app'. It automatically downloads the best JRE image, optimizes memory settings, layers the JAR, and gives you a production-ready Docker image. It's magic."

**Indepth:**
> **Rebase**: The coolest feature. If there is a security patch in the underlying OS (Ubuntu SSL), you don't need to rebuild your app. You just "rebase" the image layers. The app layer stays the same; the OS layer is swapped underneath it instantly.


---

**Q: Blue-Green Deployment**
> "You have Version 1 running (Blue).
> You deploy Version 2 (Green) alongside it.
> You run tests on Green.
>
> Then, you switch the Load Balancer router: 100% traffic goes to Green.
> If Green crashes, you instantly switch back to Blue.
> Spring Boot doesn't do this itself, but it provides the **Metrics** and **Health Probes** that tools like Kubernetes or AWS use to orchestrate this switch safely."

**Indepth:**
> **Database**: The DB is shared. Version 2 app cannot rename a column that Version 1 app is still using. You must perform "Expand and Contract" migrations (Add new column, copy data, unrelated changes) to ensure backward compatibility.

