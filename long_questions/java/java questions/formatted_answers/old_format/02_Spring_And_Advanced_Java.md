# 02. Spring, Database, and Advanced Java

**Q: @Component vs @Service vs @Repository**
> "Technically, they are all the same! `@Service` and `@Repository` are just specializations of `@Component`. They all mark a class as a Spring Bean.
>
> But we distinguish them for semantic clarity:
> *   **@Component** is a generic bean.
> *   **@Service** marks the business logic layer.
> *   **@Repository** marks the data access layer (DAO) and enables automatic exception translation for database errors.
>
> Using the specific annotations helps you (and Spring) understand the architecture of your app."

**Indepth:**
> **Aspect-Oriented Programming (AOP)**: These annotations act as pointcuts. For `@Repository`, Spring automatically registers a `PersistenceExceptionTranslationPostProcessor` which intercepts platform-specific database exceptions (SQL/Hibernate) and re-throws them as Spring’s unified, unchecked `DataAccessException` hierarchy.
>
> **Layering**: While functional the same, `@Service` defines a service layer, often acting as the transaction boundary (`@Transactional`) for business logic unit of work.


---

**Q: Why Spring Boot?**
> "Spring Boot makes Spring easier. In the old days, setting up a Spring project meant hours of XML configuration and dependency hell.
>
> **Spring Boot** solves this with three main things:
> 1.  **Auto Configuration**: It scans your classpath and automatically sets up things for you. If it sees a database driver, it configures a DataSource.
> 2.  **Starter Dependencies**: Instead of adding 10 different libraries for a web app, you just add `spring-boot-starter-web`.
> 3.  **Embedded Server**: It comes with Tomcat built-in, so you can run your app as a simple JAR file without needing to install a separate server."

**Indepth:**
> **Opinionated Defaults**: Spring Boot uses "Convention over Configuration". It assumes standard defaults (e.g., if Hibernate is on classpath, it configures a DataSource) so you don't have to manualy configure beans.
>
> **Dependency Management**: It relies on the "Bill of Materials" (BOM) pattern in `spring-boot-dependencies` to manage library versions, ensuring compatibility between Spring Core, Logging, JSON parsers, etc., preventing "JAR hell".


---

**Q: application.properties vs yml**
> "They do the exact same thing—configure your app. The difference is just syntax.
>
> **Properties** files are flat keys: `server.port=8080`.
> **YAML** files are hierarchical:
> ```yaml
> server:
>   port: 8080
> ```
> YAML is generally preferred for complex configurations because it's cleaner and less repetitive. But watch out for indentation errors!"

**Indepth:**
> **Type Safety**: Neither is type-safe by default, but `@ConfigurationProperties` binds them to Java beans, providing compile-time validation of properties.
>
> **Profiles**: YAML supports multi-document syntax (using `---`) to define multiple profiles (dev, prod) in a single file. Properties files requires separate files (e.g., `application-dev.properties`) for each profile.


---

**Q: @Controller vs @RestController**
> "Basically, **@RestController** is a convenience annotation. It combines `@Controller` and `@ResponseBody`.
>
> If you use **@Controller**, you typically return a String that maps to a View (like an HTML page or JSP).
> If you use **@RestController**, the return value is automatically serialized to JSON (or XML) and sent directly to the HTTP response body.
>
> So, use **@Controller** for traditional web apps with pages, and **@RestController** for REST APIs."

**Indepth:**
> **Under the Hood**: `@RestController` is annotated with `@Controller` and `@ResponseBody`.
>
> **Message Converters**: When returning objects from a `@RestController`, Spring uses `HttpMessageConverter`s (like Jackson for JSON) to negotiate the content type based on the HTTP `Accept` header and serialize the Java object accordingly.


---

**Q: @RequestBody vs @RequestParam vs @PathVariable**
> "These are all ways to get data into your controller.
>
> **@PathVariable** pulls data right out of the URL path itself. Like `/users/{id}`—the `123` in `/users/123` is a path variable. Use this for identifying resources.
>
> **@RequestParam** pulls data from the query string after the `?`. Like `/search?query=java`. Use this for filtering or sorting.
>
> **@RequestBody** grabs the entire body of the HTTP request (usually JSON) and converts it into a Java Object. Use this for 'Create' or 'Update' operations where you're sending a lot of data."

**Indepth:**
> **Validation**: `@RequestBody` is often used with `@Valid` to trigger automatic bean validation (JSR-303) before the method body executes.
>
> **Encoding**: `@PathVariable` and `@RequestParam` values are usually URL-decoded by the container. `@RequestBody` reads the raw input stream and deserializes it, handling complex nested objects that URL parameters cannot easily represent.


---

**Q: Monolith vs Microservices**
> "In a **Monolith**, everything is in one big codebase and deployed as a single unit. It's simple to develop and test initially, but as it grows, it becomes hard to scale and maintain. If one part breaks, the whole thing might crash.
>
> **Microservices** break the app down into small, independent services that talk to each other (usually via REST). Each service can be developed, deployed, and scaled independently. It’s more complex to manage (infrastructure-wise), but it allows for better scalability and agility."

**Indepth:**
> **Distributed Complexity**: Microservices introduce the "Fallacies of Distributed Computing". Calls over the network introduce latency, timeouts, and partial failures, requiring resiliency patterns like Circuit Breakers (Resilience4j).
>
> **Data Consistency**: Monoliths use ACID transactions. Microservices often require eventual consistency patterns like Sagas or Event Sourcing (BASE), as distributed transactions (2PC) are hard to scale.


---

**Q: Hibernate vs JPA**
> "**JPA** (Java Persistence API) is a *specification*. It’s effectively a document that says 'This is how we should do ORM in Java.' It’s a set of interfaces.
>
> **Hibernate** is an *implementation* of that specification. It’s the actual code that does the work.
>
> Think of JPA as the 'Interface' and Hibernate as the 'Class'. You code against the JPA standard, but Hibernate runs underneath to talk to the database."

**Indepth:**
> **Provider Swapping**: Because you code to JPA interfaces (`EntityManager`), you can switch the underlying provider (Hibernate vs EclipseLink) with minimal code changes, making your app portable.
>
> **Caching**: Hibernate adds features beyond the JPA spec, such as the Second Level Cache (L2 Cache) which operates at the `SessionFactory` scope to cache data across transactions.


---

**Q: Primary vs Foreign Key**
> "A **Primary Key** uniquely identifies a row in a table. It cannot be null and must be unique. Think of it like your Social Security Number—it identifies *you*.
>
> A **Foreign Key** is a field that links to the Primary Key of *another* table. It creates a relationship between two tables. It ensures referential integrity—you can't have an 'Order' pointing to a 'Customer' that doesn't exist."

**Indepth:**
> **Indexing**: Primary Keys are automatically indexed (usually a Clustered Index), physically ordering data on disk. Foreign Keys are NOT automatically indexed in many DBs (like SQL Server/Postgres), often requiring manual index creation to prevent full table scans during joins.
>
> **Cascading**: Foreign keys enable `ON DELETE CASCADE` rules, allowing the database to automatically clean up orphaned child records when a parent is deleted.


---

**Q: DELETE vs TRUNCATE vs DROP**
> "**DELETE** is a DML command. It removes rows one by one, and you *can* rollback the transaction. It’s slower but safer for removing specific data.
>
> **TRUNCATE** is a DDL command. It nukes all the data in the table instantly by resetting the table structure. You typically *cannot* rollback easily. It’s faster.
>
> **DROP** deletes the entire table structure itself. The data is gone, and the table definition is gone. Poof."

**Indepth:**
> **Transaction Log**: `DELETE` logs every individual row removal, making it slow but recoverable. `TRUNCATE` logs only page deallocations, making it extremely fast but minimal logging means it's harder to recover without backups.
>
> **Identity Reset**: `TRUNCATE` usually resets auto-increment counters to the seed value; `DELETE` continues where it left off.


---

**Q: Shallow vs Deep Copy**
> "This is about copying objects.
>
> A **Shallow Copy** creates a new object, but inserts *references* to the objects found in the original. So if the original object points to a list, the copy points to that *same* list. Changing the list in the copy changes it in the original too!
>
> A **Deep Copy** creates a new object and recursively creates copies of everything inside it. The two objects are completely independent. Changing one does not affect the other."

**Indepth:**
> **Cloning**: The default `Object.clone()` performs a shallow copy. For deep copies, you often rely on Serialization/Deserialization (slow) or copy constructors.
>
> **Immutability**: Deep copying is unnecessary if objects are immutable, as shared references are thread-safe by definition.


---

**Q: Metaspace vs PermGen**
> "**PermGen** (Permanent Generation) was the memory area in Java 7 (and older) where class metadata was stored. It had a fixed size, so you often got `OutOfMemoryError: PermGen space`.
>
> **Metaspace** replaced PermGen in Java 8. The key difference is that Metaspace is part of native memory (OS memory), not the JVM heap. This means it can grow automatically as needed (up to the OS limit), greatly reducing those OutOfMemory errors."

**Indepth:**
> **Tuning**: In Java 8+, `MaxPermSize` is ignored. `MaxMetaspaceSize` defaults to unlimited (OS RAM), reducing startup OOMs. However, this can mask memory leaks if dynamic classes (Proxies) are generated and never unloaded.
>
> **Storage**: Metaspace stores class metadata. Static variables (which used to be in PermGen) were moved to the main Heap in Java 8.


---

**Q: Synchronized vs Lock**
> "**synchronized** is a keyword. It’s cleaner and easier to use—you just put it on a method or block, and the JVM handles the locking and unlocking automatically. But it’s rigid; you can't interrupt a thread waiting for a synchronized lock.
>
> **Lock** (like `ReentrantLock`) is a class. It gives you more control. You can try to acquire a lock with a timeout, interrupt a waiting thread, or lock/unlock in different scopes. But you *must* remember to unlock it in a `finally` block, or you'll cause a deadlock."

**Indepth:**
> **Fairness**: `ReentrantLock` allows a "fairness" policy (FCFS), which `synchronized` (always unfair) does not supports.
>
> **Condition Variables**: `Lock` supports multiple `Condition` queues (via `newCondition()`), allowing threads to wait for specific signals (not empty, not full), whereas `synchronized` monitors only support a single wait set.


---

**Q: ApplicationContext vs BeanFactory**
> "**BeanFactory** is the most basic container in Spring. It lazy-loads beans (creates them only when requested). It’s good for resource-constrained environments (like mobile/embedded), but rarely used today.
>
> **ApplicationContext** extends BeanFactory. It does everything BeanFactory does but adds enterprise features: it eagerly loads singleton beans on startup, supports internationalization (i18n), event propagation, and annotations.
>
> Always use **ApplicationContext** unless you have a crazy specific reason not to."

**Indepth:**
> **Event Handling**: `ApplicationContext` supports the Pub/Sub pattern via `ApplicationEvent` and `ApplicationListener`.
>
> **AOP Integration**: `ApplicationContext` is required for full AOP support and automatic `BeanPostProcessor` registration, which drives features like `@Autowired` and transaction management. ALWAYS use ApplicationContext in production.


---

**Q: Volatile vs Synchronized**
> "**synchronized** provides both *atomicity* and *visibility*. It ensures only one thread executes a block at a time, and changes are visible to others.
>
> **volatile** is a keyword for variables that strictly ensures *visibility*. It tells the JVM: 'Don't cache this variable in a CPU register; always read/write it from main memory.' It guarantees that if Thread A changes the value, Thread B sees it immediately. However, it does NOT guarantee atomicity (so `count++` is still not safe with just volatile)."

**Indepth:**
> **Happens-Before**: `volatile` establishes a "happens-before" relationship. A write to a volatile variable is guaranteed to be visible to any subsequent read.
>
> **Instruction Reordering**: It prevents the compiler/CPU from reordering instructions involving the volatile variable, ensuring execution order. However, it does not guarantee atomicity for compound operations (like `i++`).

