# Spring Boot, REST & JPA Interview Questions (51-70)

## Spring Boot Core

### 51. What happens internally when a Spring Boot app starts?
"It all begins with the `@SpringBootApplication` annotation, which is a meta-annotation for `@Configuration`, `@EnableAutoConfiguration`, and `@ComponentScan`.

When you run `main()`, Spring Boot initializes the `SpringApplication` class. It detects the type of application (Servlet/Reactive) based on the classpath.

Then it creates the `ApplicationContext` (the IOC container). It scans your packages for components (`@Controller`, `@Service`), processes configuration classes, and triggers **Auto-Configuration**. This looks at your classpath—'Oh, I see `spring-boot-starter-web`? I’ll configure Tomcat and Spring MVC for you.'

Finally, it runs any `CommandLineRunner` beans and marks the application as started."

### 52. Difference between `@Component`, `@Service`, `@Repository`?
"Technically, they are all the same. `@Service` and `@Repository` are just specializations of `@Component`.

`@Component` is the generic stereotype for any Spring-managed bean.

`@Service` is for the business layer. It doesn't add extra behavior *yet*, but it clarifies intent for developers reading the code.

`@Repository` is for the data layer (DAO). It *does* add behavior: it automatically translates database exceptions (like `SQLException`) into Spring’s consistent `DataAccessException` hierarchy. This is huge because it means your Service layer doesn't need to know if you're using Hibernate, JDBC, or Mongo—the exceptions are uniform."

### 53. What is auto-configuration?
"Auto-configuration is Spring Boot’s magic. It attempts to automatically configure your application based on the jar dependencies you have added.

For example, if `H2` database is on the classpath, and you haven't manually configured a `DataSource` bean, Spring Boot says, 'I guess you want an in-memory H2 database,' and configures one for you.

It works using `@Conditional` annotations—like `@ConditionalOnClass` or `@ConditionalOnMissingBean`. If you define your own `DataSource`, the auto-configuration backs off. It’s opinionated but overridable."

### 54. What is Spring Bean lifecycle?
"Simply put: Use -> Create -> Configure -> Destroy.

1.  **Instantiation**: The container creates the bean instance (constructor).
2.  **Populate Properties**: Dependencies are injected (`@Autowired`).
3.  **Initialization**: If the bean implements `InitializingBean` or has a `@PostConstruct` method, that runs now. This is where I put logic that needs dependencies to be ready, like loading a cache.
4.  **Use**: The bean is ready and used by the application.
5.  **Destruction**: When the context closes, `@PreDestroy` methods run to clean up resources like open sockets."

### 55. Difference between `@Autowired` and constructor injection?
"`@Autowired` on fields (Field Injection) is easier to write but generally considered bad practice now. It makes testing hard because you can't easily instantiate the class without reflection, and it hides dependencies.

**Constructor Injection** is the standard. You declare `final` fields and a constructor.
1.  It ensures immutability (fields can be final).
2.  It ensures the bean is never in an invalid state (it *must* have its dependencies to be created).
3.  It makes unit testing trivial—just `new MyService(mockRepo)`."

### 56. What is `@ConfigurationProperties`?
"It’s a way to bind a group of properties from `application.yml` directly to a Java POJO.

Instead of injecting individual values with `@Value("${app.mail.host}")` everywhere, I create a `MailConfig` class with fields like `host`, `port`, `username`. I annotate it with `@ConfigurationProperties(prefix="app.mail")`.

This offers type safety, validation (I can use JSR-303 annotations like `@NotNull`), and auto-completion in IDEs. It keeps configuration structured and centralized."

## REST APIs

### 57. Difference between PUT and POST?
"Semantically, **POST** is for creating a *new* resource where the server decides the ID (e.g., `/users`). It’s not idempotent—calling it twice creates two users.

**PUT** is for updating or creating a resource at a *specific* URI (e.g., `/users/123`). It is **idempotent**—calling it multiple times results in the same state (User 123 serves as the identifier).

In practice, I use POST for creation and PUT for full updates. For partial updates (like just changing a password), I prefer **PATCH**, though people often abuse PUT for that too."

### 58. What are idempotent APIs?
"Idempotency means that making multiple identical requests has the same effect as making a single request.

`GET`, `PUT`, and `DELETE` should be idempotent. If I `DELETE /users/123` five times, the result is the same: the user is gone. The first one might return 200, the rest 404, but the *server state* is the same.

`POST` is generally NOT idempotent. If I retry a payment request because of a network timeout, I might accidentally charge the customer twice. That’s why we need idempotency keys for payment APIs."

### 59. How do you handle validation in REST APIs?
"I use the standard Bean Validation API (`javax.validation` / `jakarta.validation`).

I annotate my DTO fields with `@NotNull`, `@Size`, `@Email`. Then in the Controller, I add `@Valid` before the request body parameter.

If validation fails, Spring throws a `MethodArgumentNotValidException`. I catch this in a global `@ControllerAdvice` and return a clean 400 Bad Request JSON with the list of field errors, rather than letting the stack trace leak to the client."

### 60. What are HTTP status codes you use frequently?
"Beyond the obvious **200 OK** and **500 Internal Server Error**:

**201 Created**: For successful POST requests.
**204 No Content**: For successful DELETEs or PUTs that don’t return a body.
**400 Bad Request**: Validation errors.
**401 Unauthorized**: Authentication missing/invalid.
**403 Forbidden**: Authenticated, but waiting permission (Authorization).
**404 Not Found**: Resource doesn't exist.
**409 Conflict**: Trying to create a user that already exists.
**429 Too Many Requests**: Rate limiting."

### 61. How do you handle global exception handling?
"I use `@ControllerAdvice` combined with `@ExceptionHandler`.

I create a `GlobalExceptionHandler` class. Inside, I have methods for specific exceptions like `ResourceNotFoundException` (returns 404) or `BadRequestException` (returns 400).

This separates error handling logic from the business logic handling in controllers. It ensures that no matter where an exception occurs, the API returns a consistent JSON error structure (e.g., with `timestamp`, `errorCode`, `message`)."

### 62. What is HATEOAS?
"Hypermedia as the Engine of Application State. It’s a constraint of REST that says API responses should include links to related actions.

If I get a User resource, the response shouldn't just be data; it should include links like `_self`, `orders`, `update-profile`.

Ideally, the client interacts with the API entirely through these links, like browsing a website. In reality? It adds a lot of complexity, and very few projects I’ve seen actually use it fully. Most just stick to simple JSON."

### 63. REST vs SOAP?
"SOAP is a protocol; REST is an architectural style.

SOAP relies on XML and strict contracts (WSDL). It’s heavy, verbose, but has built-in standards for security (WS-Security) and transactions. It’s still used in legacy banking/telecom systems.

REST is lightweight, usually JSON-based, and uses standard HTTP methods. It’s flexible, scalable, and browser-friendly. 99% of modern web services are REST (or gRPC/GraphQL)."

## JPA / Hibernate

### 64. Difference between `findById()` and `getOne()`?
"`findById()` (CrudRepository) usually returns an `Optional<T>`. It actually hits the database immediately (eagerly) with a SELECT query.

`getOne()` (JpaRepository, now `getReferenceById` in newer versions) returns a **proxy**. It assumes the entity exists and doesn't hit the DB immediately. It only runs the query when you access a property of that proxy. It’s useful when you just need a reference to set a foreign key relationship (e.g., `order.setCustomer(customerProxy)`) without fetching the whole Customer object."

### 65. What is the N+1 problem?
"This is the most common performance killer in ORMs like Hibernate.

It happens when you fetch a list of N entities (1 query), and then iterate over them to access a related lazy-loaded entity (N queries).

Example: fetching 10 Users. Then for each User, printing their Address. Hibernate runs 1 query for Users + 10 queries for Addresses = 11 queries.

To fix it, we use `JOIN FETCH` in JPQL to retrieve the related data in the initial query (1 query total) or use `@EntityGraph`."

### 66. Difference between `EAGER` and `LAZY` fetching?
"This defines when related entities are loaded.

`EAGER` means load immediately. If I load a User, it loads their Orders right waway. This is often wasteful if I don't need the Orders.

`LAZY` means load on demand. The Orders are only fetched when I call `user.getOrders()`. This is the default for `@OneToMany` and is generally preferred for performance.

However, `LAZY` can cause `LazyInitializationException` if you try to access the data *after* the Hibernate Session/Transaction has closed."

### 67. What is a transactional boundary?
"A transactional boundary defines the scope where database operations are atomic—either all succeed or all fail.

In Spring, we define this with `@Transactional`. Usually, we put it at the **Service layer**.

This ensures that if a method saves an Order, updates Inventory, and charges a Card, they all happen in one transaction. If the Card charge fails, the Order and Inventory updates are rolled back automatically."

### 68. Difference between `save()` and `saveAndFlush()`?
"`save()` persists the entity to the Persistence Context (First Level Cache). Hibernate might not write the SQL INSERT/UPDATE immediately; it waits until 'flush time' (usually commit) to batch updates for performance.

`saveAndFlush()` forces Hibernate to write the SQL immediately.

I rarely use `saveAndFlush()` unless I need to catch a constraint violation exception (like duplicate unique key) right there in the try-catch block, or if subsequent code relies on a trigger/DB-side calculation."

### 69. What is dirty checking?
"Dirty checking is how Hibernate updates records without us writing explicit `save()` or `update()` calls.

When you load an entity inside a `@Transactional` method, Hibernate keeps a copy of it in the persistence context. If you modify any field (e.g., `user.setName("New Name")`), Hibernate detects the difference between the object and the snapshot at the end of the transaction.

It then automatically generates and executes the SQL UPDATE statement. It’s convenient but can lead to accidental updates if you aren't careful."

### 70. What are entity states?
"JPA entities have four lifecycle states:

1.  **Transient**: Just created with `new`, not associated with any session. No ID yet.
2.  **Persistent**: Associated with a session (managed) and has a DB identity. Changes are tracked (dirty checking).
3.  **Detached**: Was persistent, but the session closed. Changes are *not* tracked anymore.
4.  **Removed**: Scheduled for deletion.

Understanding Detached vs Persistent is key to fixing obscure Hibernate bugs."
