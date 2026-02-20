# 35. Advanced REST, CLI & MongoDB

**Q: Rate Limiting per IP**
> "You need a 'Bucket' for every IP address.
> Using `Bucket4j`:
> 1.  Create a Map: `Map<String, Bucket> cache`. Key is the IP.
> 2.  In a Filter, extract IP: `request.getRemoteAddr()`.
> 3.  Get or Create the bucket for that IP.
> 4.  Call `bucket.tryConsume(1)`. If false, return 429."

**Indepth:**
> **Distributed**: Distributed Rate Limiting. In K8s with 10 pods, local memory rate limiting allows 10x the traffic. You must use a centralized store (Redis/Hazelcast) with Lua scripts to ensure atomic token consumption across the cluster.


---

**Q: @ControllerAdvice vs @ExceptionHandler**
> "**@ExceptionHandler** handles exceptions for **one specific Controller**.
>
> "**@ControllerAdvice** is global. It wraps **all Controllers**.
> Always use `@ControllerAdvice` so you have a single, central place for error handling logic (like converting `UserNotFoundException` to a 404 JSON response)."

**Indepth:**
> **Response Body**: `ResponseBodyAdvice`. You can also implement `ResponseBodyAdvice` in a `@ControllerAdvice` class to intercept and modify the *return body* of every controller (e.g., wrapping every response in a standardized `{ "data": ..., "status": "success" }` envelope).


---

**Q: Native SQL in Spring Data JPA**
> "Sometimes HQL/JPQL is too restrictive.
> Use `value` and `nativeQuery = true`.
>
> ```java
> @Query(value = "SELECT * FROM users WHERE email = ?1", nativeQuery = true)
> User findByEmail(String email);
> ```
> Use this sparingly because it ties you to a specific database (Postgres/MySQL) and breaks portability."

**Indepth:**
> **Projections**: You don't have to return the Entity. You can return an Interface (`public interface UserSummary { String getName(); }`). Spring Data JPA's Native Query mapping is smart enough to map the result set columns to the interface getters.


---

**Q: Pagination (Slice vs Page)**
> "**Page<T>** executes two queries:
> 1.  Select the data (`LIMIT 10`).
> 2.  Count total rows (`COUNT(*)`).
>
> "**Slice<T>** executes only **one** query (`LIMIT 11`).
> If it gets 11 rows, it knows there is a 'Next Page'. It doesn't calculate the total pages. This is much faster for large datasets where you only need 'Load More' buttons."

**Indepth:**
> **Keyset Pagination**: Keyset Pagination (Seek Method) is faster than Offset Pagination (`LIMIT 10 OFFSET 1000000`) for deep scrolling because it uses the index (`WHERE id > last_seen_id LIMIT 10`) instead of scanning and discarding rows.


---

**Q: Login Throttling (Brute Force Protection)**
> "Spring Security doesn't do this out of the box. You have to implement it.
>
> On `AuthenticationFailureEvent`:
> 1.  Increment a counter in Redis/DB for that Username/IP.
> 2.  If counter > 5, lock the account for 15 minutes.
>
> On `AuthenticationSuccessEvent`:
> 1.  Reset the counter."

**Indepth:**
> **Soft Lock**: Simply blocking IP is risky (CGNAT shares IPs). A better "Soft Lock" strategy is to require a ReCaptcha challenge after 3 failed attempts instead of a hard lockout.


---

**Q: Stateless vs Stateful Authentication**
> "**Stateful (Session)**: Server keeps a `SessionID` in memory map. Client sends `JSESSIONID` cookie.
> *   Pros: Easy logout (just delete session).
> *   Cons: Hard to scale horizontally (need Sticky Sessions or Redis Session Store).
>
> "**Stateless (JWT)**: Server keeps **nothing**. Client sends a signed `Token`.
> *   Pros: Instantly scalable. Server doesn't care which node handles the request.
> *   Cons: Hard to logout (cannot invalidate a token until it expires)."

**Indepth:**
> **Size**: JWTs grow linearly with claims. If you put too much data (permissions, user profile) in the token, you hit HTTP Header size limits (usually 8KB). Keep JWTs small (just UserID + Roles).


---

**Q: Live Reload (DevTools)**
> "DevTools runs a tiny LiveReload Server in your app.
> You install the **LiveReload Browser Extension**.
>
> When you compile your Java code, Spring restarts (fast).
> When you edit HTML/CSS, DevTools triggers the browser extension to refresh the page automatically. It saves you from hitting F5 a thousand times."

**Indepth:**
> **State**: `LiveReload` doesn't work well with shared state in static variables (because the classloader resets, but system classloader generic statics might not). It also consumes more memory in Dev mode.


---

**Q: Spring Boot CLI**
> "It's a command-line tool that lets you run Groovy scripts as Spring Boot apps.
>
> File `app.groovy`:
> ```groovy
> @RestController
> class App {
>     @GetMapping("/")
>     def home() { "Hello" }
> }
> ```
> Run: `spring run app.groovy`.
> It automatically imports dependencies and starts Tomcat. Great for quick prototyping."

**Indepth:**
> **POCs**: It's rarely used for production apps. It's primarily for quick Proof of Concepts (POCs) or scripting server-side tasks where you need the power of the Spring ecosystem without the boilerplate of Maven/Gradle.


---

**Q: MongoDB Query Methods**
> "Spring Data MongoDB works just like JPA.
>
> `interface UserRepo extends MongoRepository<User, String>`
>
> You can write:
> `List<User> findByLastNameAndAgeGreaterThan(String name, int age);`
>
> Spring automatically translates this into a MongoDB JSON query:
> `{ "lastName": name, "age": { "$gt": age } }`."

**Indepth:**
> **JSON Query**: You can write raw JSON queries: `@Query("{ 'age' : { $gt: ?0 } }")`. This gives you access to specific Mongo operators (`$elemMatch`, `$regex`) that method naming conventions can't express effortlessly.

