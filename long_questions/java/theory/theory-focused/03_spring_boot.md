# Spring Boot & Frameworks - Interview Answers

> 🎯 **Focus:** These answers assume you are a developer using the framework daily, focusing on "how it helps me" vs "internal theory".

### 1. Why Spring Boot over legacy Spring?
"Spring Boot is basically 'Opinionated Spring.' It prefers convention over configuration.

Legacy Spring required massive XML files and manual setup for everything—Datasources, ViewResolvers, TransactionManagers. It took days to just get a 'Hello World' running.

**Spring Boot** changes that with **Auto-Configuration**. It looks at my classpath—sees that I have the MySQL driver and Hibernate—and automatically configures the database beans for me. Plus, it has an **Embedded Server** (Tomcat), so I can just run a generic JAR file instead of managing a complex external WebLogic server. It lets me focus on business logic immediately."

---

### 2. How does Auto-Configuration work?
"It’s driven by the `@EnableAutoConfiguration` annotation, which is part of the main `@SpringBootApplication` annotation.

On startup, Spring scans a file called `spring.factories` (or imports files in Boot 3). It loads configuration classes based on **Conditions**—like `@ConditionalOnClass`.

It’s essentially saying: 'Is the Jackson library on the classpath? Yes. Is there already an `ObjectMapper` bean defined? No. Okay, I'll create a default one for you.'

It’s magic until it breaks, but when it does, I just use the `debug=true` property to see the report and find out exactly why a bean was or wasn't loaded."

---

### 3. @Component vs @Service vs @Repository?
"Functionally, they are almost effectively the same—they are all stereotypes of `@Component` that tell Spring 'manage this bean.'

But semantically, they define the architecture:
**@Service** marks the business logic layer.
**@Controller** marks the web layer.
**@Repository** marks the data access layer.

The key technical difference is with `@Repository`. It enables **Persistence Exception Translation**, which catches DB-specific exceptions (like a clear Postgres error) and translates them into a cleaner, unchecked Spring `DataAccessException`. That's why I always use strict annotations instead of just generic `@Component`."

---

### 4. @RestController vs @Controller?
"`@RestController` is just a convenience composition of `@Controller` plus `@ResponseBody`.

**@Controller** is the legacy MVC way. It typically returns a String which resolves to a view template (like `index.html`).

**@RestController** assumes you are building a REST API. It implies that the return value of every method should be serialized directly into the HTTP response body (usually as JSON), rather than looking for a view. Since I mostly build backends for React/Angular apps, I almost exclusively use `RestController`."

---

### 5. What is Dependency Injection (DI)?
"It’s a design pattern where objects don't create their own dependencies; they ask for them. This is **Inversion of Control**.

Instead of my `UserService` saying `new UserRepository()`, I just add `UserRepository` to the constructor. The Spring container sees this, finds the correct bean, and injects it for me.

This makes code incredibly **testable**. In my unit tests, I can easily pass in a `MockRepository` to the constructor. If I had hard-coded `new UserRepository()`, testing would be a nightmare because I'd be stuck with the real database connection."

---

### 6. Bean Scopes in Spring?
"Scope defines the lifecycle of a bean.

The default is **Singleton**. Spring creates one instance per application context and reuses it everywhere. This is what I use 99% of the time for stateless services.

**Prototype** creates a new instance every time you ask for it. I rarely use this, maybe for a stateful parser that isn't thread-safe.

Then there are web scopes like **Request** (one per HTTP request) and **Session** (one per User Session), which are useful if I need to store user-specific state like a shopping cart, though stateless JWTs are more common nowadays."

---

### 7. How does JPA/Hibernate work?
"JPA is the specification, and Hibernate is the implementation. It’s an ORM that maps Java Objects to Relational Tables.

I define an entity like `@Entity User`. Hibernate reads this and handles the SQL for me. When I call `save(user)`, it generates the `INSERT` statement.

It also manages a **Session** (First Level Cache). If I fetch an object and modify it, Hibernate's 'Dirty Checking' detects the change and automatically flushes an update to the Database at the end of the transaction. It saves a boilerplate, though I have to watch out for the N+1 select problem."

---

### 8. N+1 Problem in Hibernate?
"This is a classic performance killer. It happens when you fetch a parent entity, and then iterate to access its children, causing extra queries for every single parent.

For example, if I fetch 1000 `Orders` and then loop through them to get the `Customer` name for each, Hibernate might fire 1 initial query for Orders, plus 1000 separate queries for each Customer. That’s 1001 queries!

I usually solve this by using `JOIN FETCH` in my JPQL query. This forces Hibernate to bring back all the related data in a single optimized join, reducing it to just one query."

---

### 9. @Transactional annotation?
"It ensures that a method runs within a database transaction. It essentially means **all or nothing**.

If I have a method that saves an Order and then updates Inventory, I mark it `@Transactional`. If the 'Update Inventory' part fails and throws a runtime exception, Spring will automatically **rollback** the 'Save Order' part, ensuring the database stays consistent.

One thing to remember is that it works via AOP proxies. So if I call a transactional method from *within* the same class, the proxy is bypassed and the transaction won't work. I always put this on my Service layer."

---

### 10. Spring Security Authentication vs Authorization?
"**Authentication** is 'Who are you?'
**Authorization** is 'What are you allowed to do?'

For **Authentication**, Spring uses the `AuthenticationManager` to verify credentials—like checking a username/password or validating a JWT signature.

For **Authorization**, once we know who you are, we check your Roles. I usually use method-level security like `@PreAuthorize('hasRole("ADMIN")')` to lock down sensitive endpoints.

So, the Login endpoint handles Authentication; the Filter Chain handles Authorization."

---

### 11. Application properties vs YAML?
"They do the exact same thing, but I prefer **YAML** because it supports hierarchy.

In `application.properties`, you have to repeat prefixes constantly (`spring.datasource.url=...`, `spring.datasource.username=...`).
In **YAML**, you can nest them nicely, which makes the configuration much easier to read—especially with Spring Boot's deeply nested structures.

It also lets me define multiple profiles in a single file using the `---` separator, which is handy for local dev."

---

### 12. PathVariable vs RequestParam?
"Both extract values from the URL, but the intent is different.

**@PathVariable** is for identifying a specific resource. It’s part of the URI itself.
`GET /users/101` -> Here `101` is the Path Variable. It’s mandatory because without it, the URI points to a different resource.

**@RequestParam** is for filtering or sorting. It comes after the `?`.
`GET /users?role=admin` -> Here `admin` is a Request Param. It’s usually optional.

So I use PathVariable for IDs, and RequestParam for options."
