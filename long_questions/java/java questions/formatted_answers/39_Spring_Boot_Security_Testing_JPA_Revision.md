# 39. Spring Boot (Security, Testing & JPA Deep Dive)

**Q: Basic Auth vs JWT**
> "**Basic Auth**: The browser sends the username/password with *every single request*. It’s simple but risky (if not using HTTPS) and requires the server to validate it every time.
>
> "**JWT (JSON Web Token)**: The user logs in *once*. The server gives them a signed 'Badge' (Token).
> For future requests, they just show the Badge. The server sees the signature is valid and lets them in. No database check needed."

**Indepth:**
> **Revocation**: The main downside of JWT is that you cannot "force logout" a user instantly properly without a blacklist (Redis). Basic Auth / Sessions are easier to invalidate server-side (just delete the session).


---

**Q: @WithMockUser (Testing Security)**
> "Testing guarded endpoints is annoying because you have to 'login' in your test code.
>
> `@WithMockUser(username="admin", roles={"ADMIN"})`
>
> This annotation automatically creates a fake 'Logged In' context for that test method. You don't need to actually send an Authorization header. It just assumes the user is already authenticated."

**Indepth:**
> **Custom User**: `@WithUserDetails` is better if your app relies on custom fields in your `User` object (like `tenantId`). It actually loads the user from your Test `UserDetailsService` rather than creating a fake mock.


---

**Q: Unit vs Integration vs E2E Tests**
> "**Unit Test**: Testing **one class** in isolation. Mock everything else. Fast (milliseconds).
>
> "**Integration Test**: Testing how **two things talk** (e.g., Service + DB, or Controller + Service). Slower (seconds).
>
> "**E2E (End-to-End)**: Testing the **entire flow**. Spin up the whole app, hit the API, check the DB. Slowest (minutes)."

**Indepth:**
> **Testcontainers**: For Integration/E2E tests, *always* use Testcontainers. Using H2 for integration tests is dangerous because H2 behaves differently than Postgres (syntax, locking, constraints).


---

**Q: @SpringBootTest**
> "This is the 'Big Gun' of testing.
> It starts your **entire** application. It finds your main class, reads `application.properties`, connects to the (test) database, and initializes all beans.
>
> Use this when you want to be 100% sure the application wires together correctly, but use it sparingly because it's slow."

**Indepth:**
> **Random Port**: `RANDOM_PORT`. Always use `webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT`. This allows running tests in parallel without port conflicts.


---

**Q: Spring Boot Admin**
> "Managing 50 microservices is a nightmare.
> **Spring Boot Admin** is a UI Dashboard.
>
> All your microservices (Clients) register with the Admin Server.
> The UI shows you a traffic light view: 'Service A is DOWN', 'Service B is running out of memory'. You can also view logs and change log levels directly from the browser."

**Indepth:**
> **Security**: The Admin server itself must be secured (login), otherwise anyone can view your env vars and heap dumps. The Client must also authenticate to register itself.


---

**Q: Entity Lifecycle States (JPA)**
> "An Entity can be in 4 states:
> 1.  **Transient**: Just created (`new User()`). Hibernate doesn't know about it.
> 2.  **Persistent**: Currently managed by Hibernate (e.g., just saved or fetched). Changes are auto-saved.
> 3.  **Detached**: Session is closed. Java object exists, but changes won't be saved to DB.
> 4.  **Removed**: Scheduled for deletion."

**Indepth:**
> **Merge**: `merge()` vs `persist()`. `persist` takes a transient instance. `merge` takes a detached instance and copies its state to a managed instance. It returns the *new* managed instance; it doesn't make the passed object managed.


---

**Q: Dirty Checking**
> "In Hibernate, you rarely call `save()` for updates.
> You just fetch an object, change a field (`user.setName("New Name")`), and walk away.
>
> When the Transaction commits, Hibernate compares the object with its original snapshot. If it sees a change, it **automatically** generates specific `UPDATE` SQL. This is Dirty Checking."

**Indepth:**
> **Optimization**: If you mark a transaction `@Transactional(readOnly = true)`, Spring disables dirty checking for performance. Hibernate won't snapshot the objects, saving memory and CPU.


---

**Q: L1 vs L2 Cache**
> "**L1 Cache (Session)**: Enabled by default. If you ask for User #1 twice in the *same* transaction, Hibernate returns the same object instantly without hitting the DB twice.
>
> "**L2 Cache (SessionFactory)**: Shared across *all* users. Requires a provider (EhCache/Redis). If User A fetches 'Country List', User B gets it from the cache. rarely used in Microservices (we prefer Redis at the app layer)."

**Indepth:**
> **Stale Data**: L2 Cache is tricky in distributed systems. If Instance A updates a user, Instance B's cache might be stale. You need a distributed cache invalidation mechanism (like Redis Pub/Sub or Hibernate Clustered Cache).


---

**Q: JPA Projections**
> "Don't fetch the whole Entity if you only need the name.
>
> Define an Interface:
> ```java
> interface UserNameOnly {
>     String getFirstName();
>     String getLastName();
> }
> ```
> `List<UserNameOnly> findByAge(int age);`
>
> Spring generates an efficient SQL query: `SELECT first_name, last_name FROM...` instead of `SELECT *`."

**Indepth:**
> **Class-based**: You can also use a Java Record/Class with a constructor matching the fields. `Select new com.example.UserDto(u.name, u.age)...`. This is type-safe and cleaner than Interfaces.


---

**Q: Auditing (@CreatedDate)**
> "Stop manually setting `createdAt` and `updatedAt`.
> 1.  Add `@EnableJpaAuditing`.
> 2.  Add `@EntityListeners(AuditingEntityListener.class)` to your entity.
> 3.  Annotate fields: `@CreatedDate`, `@LastModifiedDate`, `@CreatedBy`.
>
> Spring automatically fills these timestamps whenever you save or update."

**Indepth:**
> **AuditorAware**: How does Spring know *who* logged in? You implement `AuditorAware<String>`. It fetches the current user from `SecurityContextHolder` and returns the username to be stamped on the record.

