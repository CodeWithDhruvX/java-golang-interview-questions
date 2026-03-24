# Spring Boot, REST & JPA Interview Questions (51-70)

## Spring Boot Core

### 51. What happens internally when a Spring Boot app starts?
"It all begins with the `@SpringBootApplication` annotation, which is a meta-annotation for `@Configuration`, `@EnableAutoConfiguration`, and `@ComponentScan`.

When you run `main()`, Spring Boot initializes the `SpringApplication` class. It detects the type of application (Servlet/Reactive) based on the classpath.

Then it creates the `ApplicationContext` (the IOC container). It scans your packages for components (`@Controller`, `@Service`), processes configuration classes, and triggers **Auto-Configuration**. This looks at your classpath—'Oh, I see `spring-boot-starter-web`? I’ll configure Tomcat and Spring MVC for you.'

Finally, it runs any `CommandLineRunner` beans and marks the application as started."

**Spoken Format:**
"Starting a Spring Boot app is like launching a rocket - there's a sophisticated sequence that happens automatically.

It all begins with the `@SpringBootApplication` annotation, which is like the mission control button that says 'This is a Spring Boot mission'.

When you run the main method, Spring Boot creates the `ApplicationContext` - this is like the mission control center that manages everything. It scans your code for special components (controllers, services, repositories) and automatically configures everything.

The magic is in the auto-configuration - Spring Boot looks at your classpath and says 'Oh, I see you have web dependencies, I'll set up a web server for you. I see you have database dependencies, I'll configure that too.'

It's like having a smart assistant who reads your room and sets up everything you need before you even ask for it.

Finally, it runs any startup tasks and says 'Ready for launch!' The whole process is like a well-orchestrated ceremony that gets your application from zero to running without you having to configure much manually."

### 52. Difference between `@Component`, `@Service`, `@Repository`?
"Technically, they are all the same. `@Service` and `@Repository` are just specializations of `@Component`.

`@Component` is the generic stereotype for any Spring-managed bean.

`@Service` is for the business layer. It doesn't add extra behavior *yet*, but it clarifies intent for developers reading the code.

`@Repository` is for the data layer (DAO). It *does* add behavior: it automatically translates database exceptions (like `SQLException`) into Spring’s consistent `DataAccessException` hierarchy. This is huge because it means your Service layer doesn't need to know if you're using Hibernate, JDBC, or Mongo—the exceptions are uniform."

**Spoken Format:**
"These annotations are like different uniforms for different jobs in your company - they all work for Spring, but each has a specific role.

`@Component` is like the standard employee uniform - anyone can wear it, it means 'I'm part of the Spring team'.

`@Service` is like the manager uniform - it tells everyone 'I handle business logic'. It doesn't give you special powers, but it makes your role clear to others.

`@Repository` is like the database administrator uniform - it tells everyone 'I handle data operations'. The special thing about this uniform is that it comes with built-in translation services.

If you're using Hibernate and get a `SQLException`, the repository uniform automatically translates it to Spring's `DataAccessException`. If you switch to MongoDB, it translates MongoDB exceptions to the same Spring exception.

This is like having a universal translator - no matter which database language your data people speak, you always get the error message in your company's standard language!

So while they're all technically the same (all are Spring components), each uniform serves a different purpose and makes the code organization clear to everyone."

### 53. What is auto-configuration?
"Auto-configuration is Spring Boot’s magic. It attempts to automatically configure your application based on the jar dependencies you have added.

For example, if `H2` database is on the classpath, and you haven't manually configured a `DataSource` bean, Spring Boot says, 'I guess you want an in-memory H2 database,' and configures one for you.

It works using `@Conditional` annotations—like `@ConditionalOnClass` or `@ConditionalOnMissingBean`. If you define your own `DataSource`, the auto-configuration backs off. It’s opinionated but overridable."

**Spoken Format:**
"Auto-configuration is like having a very smart assistant who sets up your entire workspace based on what tools you bring.

Imagine you walk into your office with a laptop. Spring Boot looks at your laptop and says 'Oh, I see you have PowerPoint installed, I'll set up the projector. I see you have Excel, I'll open spreadsheet software.'

The magic is in the conditional annotations - Spring Boot only sets up the projector if it detects PowerPoint. If you bring your own projector, it says 'I see you already have a projector, I won't set up another one.'

For databases, it's even smarter. If Spring Boot sees H2 database on your classpath, it says 'I guess you want an in-memory database for development' and sets one up automatically.

But if you bring your own MySQL configuration, Spring Boot respects that and uses yours instead.

It's like having a assistant who makes smart assumptions but always defers to your explicit choices. The assistant sets everything up automatically, but you're always in control and can override anything you want!"

### 54. What is Spring Bean lifecycle?
"Simply put: Use -> Create -> Configure -> Destroy.

1.  **Instantiation**: The container creates the bean instance (constructor).
2.  **Populate Properties**: Dependencies are injected (`@Autowired`).
3.  **Initialization**: If the bean implements `InitializingBean` or has a `@PostConstruct` method, that runs now. This is where I put logic that needs dependencies to be ready, like loading a cache.
4.  **Use**: The bean is ready and used by the application.
5. **Destruction**: When the context closes, `@PreDestroy` methods run to clean up resources like open sockets."

**Spoken Format:**
"Spring Bean lifecycle is like the life cycle of an employee - from hiring to retirement.

**1. Instantiation**: This is like hiring a new employee. Spring creates the bean instance using the constructor. The employee is ready to start working.

**2. Populate Properties**: This is like giving the employee their tools and equipment. Spring injects all the dependencies they need to do their job.

**3. Initialization**: This is like orientation day. If the employee needs to set up their workspace or load initial data, the `@PostConstruct` method runs. They're fully ready to start work.

**4. Use**: This is the employee's productive years. They're using their skills to help the application run.

**5. Destruction**: This is like retirement. When the application shuts down, the `@PreDestroy` method runs. It's the employee's chance to clean up their desk, close open files, save their work.

Spring manages this entire lifecycle automatically - you just focus on what the employee should do, not on the logistics of hiring and firing!"

### 55. Difference between `@Autowired` and constructor injection?
"`@Autowired` on fields (Field Injection) is easier to write but generally considered bad practice now. It makes testing hard because you can't easily instantiate the class without reflection, and it hides dependencies.

**Constructor Injection** is the standard. You declare `final` fields and a constructor.
1.  It ensures immutability (fields can be final).
2.  It ensures the bean is never in an invalid state (it *must* have its dependencies to be created).
3. It makes unit testing trivial—just `new MyService(mockRepo)`."

**Spoken Format:**
"This is about how you give your employees their tools - there are two main approaches.

**Field Injection** is like leaving tools on employee's desk and hoping they pick the right ones. You write `@Autowired private EmailService emailService;`. The problem is you can't easily test this employee because you can't give them different tools easily.

**Constructor Injection** is like handing employees their tools when they walk in the door. You write `private final EmailService emailService;` and require it in the constructor.

The benefits are huge:

1. **Immutability** - once you give employees their tools, they can't lose or change them. The final fields ensure the object is always in a valid state.

2. **Never invalid** - an employee can't start work without their essential tools. The constructor guarantees all dependencies are available.

3. **Easy testing** - for testing, you just give the employee fake tools (mocks) to see if they can do their job. No reflection needed!

Constructor injection is like being a good manager - you make sure your team has everything they need before they start working, not after!"

### 56. What is `@ConfigurationProperties`?
"It’s a way to bind a group of properties from `application.yml` directly to a Java POJO.

Instead of injecting individual values with `@Value("${app.mail.host}")` everywhere, I create a `MailConfig` class with fields like `host`, `port`, `username`. I annotate it with `@ConfigurationProperties(prefix="app.mail")`.

This offers type safety, validation (I can use JSR-303 annotations like `@NotNull`), and auto-completion in IDEs. It keeps configuration structured and centralized."

**Spoken Format:**
"`@ConfigurationProperties` is like having a personalized settings panel that automatically syncs with your configuration files.

Instead of writing code like `@Value("${app.mail.host}") private String host;` everywhere, you create a dedicated `MailConfig` class:

```
@ConfigurationProperties(prefix="app.mail")
public class MailConfig {
    private String host;
    private int port;
    private String username;
    // getters and setters
}
```

The magic is that Spring Boot automatically reads your `application.yml` file and maps values like `app.mail.host`, `app.mail.port`, `app.mail.username` to these fields.

The benefits are:

1. **Type Safety** - you get autocomplete in your IDE and compile-time checking instead of string-based errors.

2. **Validation** - you can add `@NotNull` or `@Min` annotations to validate values.

3. **Centralization** - all mail-related settings are in one place, not scattered across your code.

4. **Documentation** - IDEs can show you all available settings for the mail module.

It's like having a remote control for your TV settings - all the buttons are organized and labeled, instead of having to remember which button does what!"

## REST APIs

### 57. Difference between PUT and POST?
"Semantically, **POST** is for creating a *new* resource where the server decides the ID (e.g., `/users`). It’s not idempotent—calling it twice creates two users.

**PUT** is for updating or creating a resource at a *specific* URI (e.g., `/users/123`). It is **idempotent**—calling it multiple times results in the same state (User 123 serves as the identifier).

In practice, I use POST for creation and PUT for full updates. For partial updates (like just changing a password), I prefer **PATCH**, though people often abuse PUT for that too."

**Spoken Format:**
"Think of these like different types of delivery instructions for a package.

**POST** is like asking someone to create something new - 'Please create a new user account'. The server decides the account number (like user ID 123). If you ask again, you get another account (User ID 124). It's not idempotent.

**PUT** is like asking someone to update a specific account - 'Please update account number 456 with this information'. You're telling exactly which account to update. If you ask multiple times, you get the same result - it's idempotent.

**PATCH** is like asking someone to make specific changes to an account - 'Please just change the email address for account 456'. It's for partial updates.

The key insight: POST creates new resources, PUT updates entire resources, PATCH updates parts of resources.

Most people use POST for creation and PUT for updates, but PATCH is technically correct for partial updates. However, many systems abuse PUT for partial updates because it's simpler to implement!"

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

This separates error handling logic from business logic handling in controllers. It ensures that no matter where an exception occurs, the API returns a consistent JSON error structure (e.g., with `timestamp`, `errorCode`, `message`)."

**Spoken Format:**
"Global exception handling is like having a customer service department that handles every complaint the same way.

Instead of each employee handling complaints differently (and maybe giving inconsistent responses), you create a central `@ControllerAdvice` class.

When any exception happens anywhere in your application, Spring routes it to this central handler. It's like saying 'No matter what went wrong, send all complaints to the customer service team.'

Inside this handler, you have specific methods for different types of problems:
- `ResourceNotFoundException` → returns 404 Not Found
- `BadRequestException` → returns 400 Bad Request
- `UnauthorizedException` → returns 401 Unauthorized

The beauty is that your API always returns consistent error responses in JSON format, with fields like `timestamp`, `errorCode`, `message`, and `path`.

This way, whether a user tries to access a non-existent user or sends invalid data, they get the same professional error response structure. It's like having a standard complaint form for all customer service issues!"

### 62. What is HATEOAS?
"Hypermedia as the Engine of Application State. It’s a constraint of REST that says API responses should include links to related actions.

If I get a User resource, the response shouldn't just be data; it should include links like `_self`, `orders`, `update-profile`.

Ideally, client interacts with API entirely through these links, like browsing a website. In reality? It adds a lot of complexity, and very few projects I've seen actually use it fully. Most just stick to simple JSON."

**Spoken Format:**
"HATEOAS is like giving someone a map with your API responses instead of just the data they asked for.

When you request a user, instead of just returning the user data, you also include links like:
- `_self`: link to the user itself
- `orders`: link to the user's orders
- `profile`: link to the user's profile

The idea is that the client can navigate your entire API using these links, like browsing a website by clicking links instead of manually typing URLs.

For example, after getting a user, the client can follow the `orders` link to see all their orders, then the `profile` link to see their details, all without knowing the specific URLs.

The challenge is that it adds complexity - you need to generate and maintain these links. But in theory, it makes APIs more discoverable and self-documenting.

In practice, most projects skip the complexity and just return simple JSON. It's like giving someone a business card instead of a full map to your company!"

### 63. REST vs SOAP?
"SOAP is a protocol; REST is an architectural style.

SOAP relies on XML and strict contracts (WSDL). It’s heavy, verbose, but has built-in standards for security (WS-Security) and transactions. It’s still used in legacy banking/telecom systems.

REST is lightweight, usually JSON-based, and uses standard HTTP methods. It's flexible, scalable, and browser-friendly. 99% of modern web services are REST (or gRPC/GraphQL)."

**Spoken Format:**
"This is like comparing two different ways of sending messages.

**SOAP** is like sending a formal business letter - it's very structured, uses specific format (XML), and comes with a detailed instruction manual (WSDL). It has built-in features for security (WS-Security) and transactions. It's reliable but heavy and complex.

**REST** is like sending a quick text message - it's lightweight, uses common format (JSON), and works with simple HTTP methods that everyone understands. It's flexible and easy to work with.

The key differences:
- SOAP is strict and formal - good for enterprise banking where reliability and standards matter
- REST is flexible and simple - good for modern web apps where speed and scalability matter

SOAP is like wearing a suit to a business meeting - very formal but restrictive. REST is like casual Friday dress code - comfortable and flexible.

Today, almost everyone uses REST (or newer alternatives like gRPC and GraphQL) because it's simpler, faster, and works better with modern web browsers and mobile apps."

## JPA / Hibernate

### 64. Difference between `findById()` and `getOne()`?
"`findById()` (CrudRepository) usually returns an `Optional<T>`. It actually hits the database immediately (eagerly) with a SELECT query.

`getOne()` (JpaRepository, now `getReferenceById` in newer versions) returns a **proxy**. It assumes the entity exists and doesn't hit the DB immediately. It only runs the query when you access a property of that proxy. It's useful when you just need a reference to set a foreign key relationship (e.g., `order.setCustomer(customerProxy)`) without fetching the whole Customer object."

**Spoken Format:**
"This is about when Hibernate actually goes to the database to get data.

`findById()` is like sending someone to the store with a shopping list - they go to the database immediately and get everything you asked for. It returns the actual product (entity), not just a promise to get it later.

`getReferenceById()` is like giving someone a note that says 'Product #123 exists at this store'. You don't actually go get the product yet, but you have a reference to it.

The beauty is performance - you don't hit the database until you actually need the product details.

This is perfect for setting relationships. If you have an order and want to set the customer, you can use `getReferenceById()` to get a customer reference without fetching all customer data.

The gotcha with LAZY is the `LazyInitializationException`. This happens when you try to access lazy-loaded data after the database connection is closed. It's like trying to get pizza toppings after the restaurant has closed for the night.

The rule is: use LAZY by default for performance, but be careful about when and where you access the lazy data!"

### 65. What is the N+1 problem?
"This is the most common performance killer in ORMs like Hibernate.

It happens when you fetch a list of N entities (1 query), and then iterate over them to access a related lazy-loaded entity (N queries).

Example: fetching 10 Users. Then for each User, printing their Address. Hibernate runs 1 query for Users + 10 queries for Addresses = 11 queries.

To fix it, we use `JOIN FETCH` in JPQL to retrieve related data in initial query (1 query total) or use `@EntityGraph`."

**Spoken Format:**
"The N+1 problem is like a performance trap that's easy to fall into.

Imagine you're fetching 10 users from the database. Then for each user, you want to display their address. Without lazy loading optimization, this becomes:

1 query to get all 10 users
10 queries to get address for each user
Total: 11 queries

This is the N+1 problem - you run N+1 queries instead of just 1 or 2.

The solutions are:

1. **JOIN FETCH** in JPQL - like asking the database 'Give me users AND their addresses in one trip'. This becomes 1 query total.

2. **EntityGraph** - like giving the database a map of what you'll need ahead of time. You tell it 'Whenever you give me a user, also give me their address'.

Both solutions turn 11 queries into just 1 or 2 queries. The performance improvement is huge, especially as the number of related entities grows.

It's like going to the grocery store once with a complete shopping list versus going back and forth for each item. One trip is always more efficient!"

### 66. Difference between `EAGER` and `LAZY` fetching?
"This defines when related entities are loaded.

`EAGER` means load immediately. If I load a User, it loads their Orders right waway. This is often wasteful if I don't need the Orders.

`LAZY` means load on demand. The Orders are only fetched when I call `user.getOrders()`. This is the default for `@OneToMany` and is generally preferred for performance.

However, `LAZY` can cause `LazyInitializationException` if you try to access data *after* the Hibernate Session/Transaction has closed."

**Spoken Format:**
"This is about when Hibernate loads related data from the database.

**EAGER loading** is like when you order a pizza and automatically get all the toppings with it. Whether you want them or not, you're paying for them and carrying them around. It's convenient but wasteful if you don't need the toppings.

**LAZY loading** is like ordering a pizza and only getting the toppings when you actually ask for them. The pizza comes plain, and when you want to see the toppings, you make a separate request.

The performance difference is huge - with EAGER loading, getting one user might load their entire order history. With LAZY, you get just the user, and only load order history when you click 'View Orders'.

The gotcha with LAZY is the `LazyInitializationException`. This happens when you try to access lazy-loaded data after the database connection is closed. It's like trying to get pizza toppings after the restaurant has closed for the night.

The rule is: use LAZY by default for performance, but be careful about when and where you access the lazy data!"

### 67. What is a transactional boundary?
"A transactional boundary defines the scope where database operations are atomic—either all succeed or all fail.

In Spring, we define this with `@Transactional`. Usually, we put it at the **Service layer**.

This ensures that if a method saves an Order, updates Inventory, and charges a Card, they all happen in one transaction. If the Card charge fails, the Order and Inventory updates are rolled back automatically."

**Spoken Format:**
"Transactional boundaries are like safety nets for database operations - they ensure that either everything succeeds or everything fails.

Imagine you're processing an online payment that involves three steps:
1. Save the order to database
2. Update inventory to reserve the product
3. Charge the customer's credit card

Without transactions, if step 3 fails (card declined), you might have saved the order but not updated inventory. Now you have an order for a product that's out of stock!

With `@Transactional`, Spring says 'All three steps must succeed or none at all'. If the card charge fails, Spring automatically undoes the order save and inventory update.

This is the ACID principle - Atomicity (all or nothing), Consistency (database stays valid), Isolation (transactions don't interfere), and Durability (committed changes persist).

It's like having a group project where either everyone completes their part successfully, or the whole project is cancelled. No in-between states!"

### 68. Difference between `save()` and `saveAndFlush()`?
"`save()` persists the entity to the Persistence Context (First Level Cache). Hibernate might not write the SQL INSERT/UPDATE immediately; it waits until 'flush time' (usually commit) to batch updates for performance.

`saveAndFlush()` forces Hibernate to write the SQL immediately.

I rarely use `saveAndFlush()` unless I need to catch a constraint violation exception (like duplicate unique key) right there in the try-catch block, or if subsequent code relies on a trigger/DB-side calculation."

**Spoken Format:**
"This is about when Hibernate actually writes changes to the database.

`save()` is like writing a check and putting it in the mail - you don't actually send it yet, it's just waiting to be sent with other checks.

Hibernate batches these checks together for performance - like collecting all the mail to send in one big trip to the post office.

`saveAndFlush()` is like sending the check immediately via express delivery - you want it to go out right now, not wait for batching.

The main time I use `saveAndFlush()` is when I need to catch specific database errors immediately. For example, if I try to save a user with a duplicate email, I want to catch that `ConstraintViolationException` right away.

Or if I have database triggers that should run after the save (like updating a timestamp), I need the save to happen immediately so the trigger runs.

Normally, I let Hibernate batch the saves for better performance, but `saveAndFlush()` is useful for these immediate feedback cases!"

### 69. What is dirty checking?
"Dirty checking is how Hibernate updates records without us writing explicit `save()` or `update()` calls.

When you load an entity inside a `@Transactional` method, Hibernate keeps a copy of it in the persistence context. If you modify any field (e.g., `user.setName("New Name")`), Hibernate detects the difference between the object and the snapshot at the end of the transaction.

It then automatically generates and executes an SQL UPDATE statement. It’s convenient but can lead to accidental updates if you aren’t careful."

**Spoken Format:**
"Dirty checking is like having a smart notebook that tracks your changes automatically.

When you load an entity inside a transaction, Hibernate takes a snapshot of how it looks initially - like taking a photo of the entity.

If you make any changes to the entity (like `user.setName("New Name")`), Hibernate compares the current state with the initial photo. If they're different, it automatically generates an UPDATE SQL statement.

This happens without you calling any save or update methods - it's completely automatic.

The convenience is huge - you can modify objects naturally and Hibernate handles the database updates.

The danger is also real - you might accidentally update something you didn't mean to change. Like if you modify a field in a loop you didn't realize was being tracked.

It's like having an over-enthusiastic assistant who writes down every change you make and immediately tells the database about it. Great for productivity, but you need to be careful what you say around them!"

### 70. What are entity states?
"JPA entities have four lifecycle states:

1.  **Transient**: Just created with `new`, not associated with any session. No ID yet.
2.  **Persistent**: Associated with a session (managed) and has a DB identity. Changes are tracked (dirty checking).
3.  **Detached**: Was persistent, but the session closed. Changes are *not* tracked anymore.
4.  **Removed**: Scheduled for deletion.

Understanding Detached vs Persistent is key to fixing obscure Hibernate bugs."

**Spoken Format:**
"Entity states are like the different relationship status an entity can have with Hibernate.

**Transient**: This is like a newborn baby - just created with `new`, not yet registered with the government (database). No ID yet, Hibernate doesn't even know it exists.

**Persistent**: Like a registered citizen - has an ID card, is in the government system (database session), and the government tracks all changes. If you change the citizen's name, Hibernate notes it down.

**Detached**: Like an emigrated citizen - was once registered but moved away (session closed). The government no longer tracks changes. If you find this citizen and change their name, the government database won't know about it.

**Removed**: Like a deceased citizen - scheduled for removal from the system.

The key insight: most bugs happen when developers don't realize they're working with a detached entity. They try to save changes, but Hibernate ignores them because it's not tracking that entity anymore.

Understanding these states is like knowing the difference between talking to someone who's actively listening versus someone who moved away - it explains why your changes sometimes don't stick!"
