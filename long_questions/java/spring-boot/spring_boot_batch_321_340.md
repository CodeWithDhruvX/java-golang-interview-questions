## 🔹 Section 3: REST APIs & Web Layer (321-330)

### Question 321: How do you implement request rate limiting per IP in Spring Boot?

**Answer:**
(Duplicate of Q150). Bucket4j or Redis.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement request rate limiting per IP in Spring Boot?
**Your Response:** "I implement rate limiting per IP using libraries like Bucket4j or Redis. Bucket4j uses the token bucket algorithm where I configure how many requests an IP can make within a time window. For distributed systems, I use Redis to store rate limit counters that can be shared across multiple application instances. I implement this as a filter or aspect that checks the IP address against the current rate limit before processing the request. This prevents abuse and ensures fair resource allocation among clients."

---

### Question 322: What is the difference between `@ControllerAdvice` and `@ExceptionHandler`?

**Answer:**
(Duplicate of Q152).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@ControllerAdvice` and `@ExceptionHandler`?
**Your Response:** "`@ExceptionHandler` handles exceptions within a single controller class, while `@ControllerAdvice` handles exceptions globally across all controllers. I use `@ExceptionHandler` when I want controller-specific exception handling, and `@ControllerAdvice` for application-wide exception handling strategies. `@ControllerAdvice` is particularly useful for implementing consistent error responses across the entire application, like standardizing error formats or handling common exceptions like validation errors."

---

### Question 323: What is the role of `@CrossOrigin` in Spring Boot?

**Answer:**
(Duplicate of Q153).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `@CrossOrigin` in Spring Boot?
**Your Response:** "`@CrossOrigin` configures Cross-Origin Resource Sharing (CORS) for specific controller methods or entire controllers. It allows me to specify which origins can access my API, what HTTP methods are allowed, and whether credentials can be included. I can apply it at the method level for fine-grained control or at the controller level for broader access. For global CORS configuration, I prefer using a `CorsConfigurationSource` bean, but `@CrossOrigin` is perfect for quick, controller-specific CORS rules."

---

### Question 324: How do you implement role-based access control at controller level?

**Answer:**
(Duplicate of Q154).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement role-based access control at controller level?
**Your Response:** "I implement role-based access control using Spring Security annotations on my controller methods. I use `@PreAuthorize('hasRole('ADMIN')')` to restrict access to users with specific roles, or `@PreAuthorize('hasAuthority('READ_USERS')')` for fine-grained permissions. I can also access method parameters like `@PreAuthorize('#id == authentication.name')` for object-level security. These annotations are evaluated before the method executes, providing declarative security that's clean and easy to understand."

---

### Question 325: How do you deal with file streaming (PDF, ZIP) in Spring Boot REST?

**Answer:**
(Duplicate of Q155).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you deal with file streaming (PDF, ZIP) in Spring Boot REST?
**Your Response:** "I handle file streaming by returning `ResponseEntity<Resource>` from my controller methods. For PDFs, I create a `FileSystemResource` or `ByteArrayResource` and set the content type to `application/pdf`. For ZIP files, I use `application/zip`. I set headers like `Content-Disposition` to specify the filename and use `InputStreamResource` for large files to avoid loading everything into memory. This approach efficiently streams files to clients without excessive memory usage."

---

### Question 326: How do you implement ETag support in REST responses?

**Answer:**
(Duplicate of Q156).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement ETag support in REST responses?
**Your Response:** "I implement ETag support by adding an `ETag` header to responses based on the resource's content hash. Spring Boot's `ShallowEtagHeaderFilter` can automatically generate ETags for responses. On subsequent requests, clients send the `If-None-Match` header, and if the ETag matches, I return a 304 Not Modified response. This reduces bandwidth usage and improves performance by avoiding unnecessary data transfer when the content hasn't changed."

---

### Question 327: How to apply request/response compression in Spring Boot?

**Answer:**
(Duplicate of Q157).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to apply request/response compression in Spring Boot?
**Your Response:** "I enable compression in Spring Boot by setting `server.compression.enabled=true` in properties. I can configure which MIME types to compress and the minimum response size threshold. Spring Boot uses GZIP compression by default, which significantly reduces response sizes for text-based content like JSON or HTML. This is particularly valuable for APIs that return large payloads or for applications serving over slow networks. The compression is transparent to both the application code and clients."

---

### Question 328: What are the options to handle request timeout in Spring controllers?

**Answer:**
(Duplicate of Q158).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the options to handle request timeout in Spring controllers?
**Your Response:** "I handle request timeouts using several approaches. I can set `server.servlet.session.timeout` for session timeout, or use `@Transactional` with timeout for database operations. For asynchronous operations, I use `CompletableFuture` with `orTimeout()`. Spring WebFlux provides built-in timeout operators. For controller-level timeouts, I can implement interceptors that measure request duration and cancel long-running operations. The choice depends on whether I need HTTP-level, transaction-level, or business logic timeout control."

---

### Question 329: How do you enable HATEOAS in Spring Boot?

**Answer:**
(Duplicate of Q147).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you enable HATEOAS in Spring Boot?
**Your Response:** "I enable HATEOAS by adding the `spring-boot-starter-hateoas` dependency. I use `RepresentationModel` and `Link` objects to add hypermedia links to my responses. I can create `EntityModel` wrappers around my domain objects and add links using `linkTo()` and `methodOn()`. Spring HATEOAS automatically generates links for controller methods, making my API self-discoverable. This allows clients to navigate the API dynamically rather than hardcoding URLs, making the API more flexible and maintainable."

---

### Question 330: How do you use filters to intercept and modify incoming requests?

**Answer:**
(Duplicate of Q159).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use filters to intercept and modify incoming requests?
**Your Response:** "I create filters by implementing `Filter` and annotating with `@Component`. I override the `doFilter()` method to intercept requests and responses. I can modify request headers, add logging, implement authentication, or transform responses. Filters are chained, so I call `chain.doFilter()` to pass control to the next filter. Filters execute before controllers, making them perfect for cross-cutting concerns like security, logging, or request modification that applies to multiple endpoints."

## 🔹 Section 4: Data & JPA Advanced (331-340)

### Question 331: How do you use `@Query` with native SQL in Spring Data JPA?

**Answer:**
(Duplicate of Q73).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `@Query` with native SQL in Spring Data JPA?
**Your Response:** "I use `@Query` with `nativeQuery = true` to execute native SQL queries. I can write complex SQL that leverages database-specific features, then map results to entity classes or custom DTOs. For parameter binding, I use named parameters like `:name` and method parameters annotated with `@Param`. Native queries are perfect when I need performance optimizations, complex joins, or database-specific functions that JPQL can't express. However, I prefer JPQL when possible for database portability."

---

### Question 332: What is the difference between `EntityManager.merge()` and `save()`?

**Answer:**
(Duplicate of Q172).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `EntityManager.merge()` and `save()`?
**Your Response:** "`merge()` updates an existing entity or creates a new one if it doesn't exist, while `save()` is a Spring Data JPA method that delegates to `persist()` for new entities or `merge()` for existing ones. `merge()` returns the managed instance, while `persist()` doesn't return anything. I use `merge()` when I'm not sure if an entity exists, and `save()` for the simpler case when I know I'm creating a new entity. Spring Data's `save()` is more convenient for most CRUD operations."

---

### Question 333: How do you implement cascading deletes in JPA?

**Answer:**
(Duplicate of Q173).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement cascading deletes in JPA?
**Your Response:** "I implement cascading deletes using `@OneToMany(cascade = CascadeType.REMOVE)` or `@OneToOne(cascade = CascadeType.REMOVE)`. When I delete the parent entity, JPA automatically deletes the associated child entities. I can also use `orphanRemoval = true` to delete children when they're removed from the collection. I need to be careful with cascading operations as they can have performance implications. For complex delete scenarios, I might use custom delete queries to avoid loading entities into memory first."

---

### Question 334: How do you use database views in Spring Boot JPA?

**Answer:**
(Duplicate of Q174).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use database views in Spring Boot JPA?
**Your Response:** "I use database views by creating entities that map to the view instead of tables. I annotate the entity with `@Immutable` and `@Subselect` or simply reference the view name in `@Table(name = 'my_view')`. Since views are read-only, I mark the entity as immutable to prevent accidental updates. This is perfect for reporting queries or when I need to expose complex joins as simple entities. Views provide a clean way to encapsulate complex SQL while still using JPA's object mapping capabilities."

---

### Question 335: How do you map stored procedures using Spring Boot JPA?

**Answer:**
(Duplicate of Q166).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you map stored procedures using Spring Boot JPA?
**Your Response:** "I map stored procedures using `@NamedStoredProcedureQuery` on entities or `@Procedure` on repository methods. I specify the procedure name, parameters, and result mapping. For simple procedures, I can use `@Procedure(name = 'my_procedure')` directly in repository methods. Spring Data JPA automatically handles the parameter binding and result mapping. This approach allows me to leverage existing database logic while maintaining a clean Java API. Stored procedures are particularly useful for complex business logic that's better implemented in the database."

---

### Question 336: How do you implement pagination using `Slice` and `Page`?

**Answer:**
(Duplicate of Q176).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement pagination using `Slice` and `Page`?
**Your Response:** "I implement pagination using Spring Data's `Pageable` parameter in repository methods. The repository returns a `Page<T>` which contains the data slice plus total count information, or a `Slice<T>` which just contains the data and whether there's more content. I use `Page` when I need total count for pagination UI, and `Slice` when I only need to know if there's more data for infinite scrolling. Spring Data automatically generates the LIMIT/OFFSET queries, making pagination implementation clean and efficient."

---

### Question 337: How do you create and use custom repository implementations in Spring Data?

**Answer:**
(Duplicate of Q177).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create and use custom repository implementations in Spring Data?
**Your Response:** "I create custom repository implementations by defining an interface with methods I want to add, then creating an implementation class with the suffix 'Impl'. Spring Data automatically finds and uses this implementation alongside the generated repository. I use this for complex queries that can't be expressed with query methods or when I need to batch operations. This approach extends Spring Data repositories while maintaining the clean programming model. It's perfect for adding custom functionality without losing the benefits of Spring Data's automatic query generation."

---

### Question 338: What’s the role of `JpaSpecificationExecutor` in Spring Data JPA?

**Answer:**
(Duplicate of Q167).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the role of `JpaSpecificationExecutor` in Spring Data JPA?
**Your Response:** "`JpaSpecificationExecutor` provides type-safe criteria API queries for dynamic filtering. I extend my repository with this interface to get access to methods like `findAll(Specification)`. I create specifications using lambda expressions that work with the Criteria API, allowing me to build complex queries programmatically. This is perfect for search functionality where users can filter by multiple optional criteria. Specifications are composable, so I can build complex queries by combining multiple conditions."

---

### Question 339: How do you use `@SqlResultSetMapping` for native queries?

**Answer:**
(Duplicate of Q179).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `@SqlResultSetMapping` for native queries?
**Your Response:** "I use `@SqlResultSetMapping` to map the results of native SQL queries to entities or DTOs when the column names don't match the field names. I define the mapping with `@Entity` and `@FieldResult` annotations to specify how each column maps to each field. Then I reference this mapping in my `@NamedNativeQuery`. This is particularly useful when working with legacy databases or complex queries that return custom result sets. It gives me full control over how SQL results map to Java objects."

---

### Question 340: How do you use `@Converter` to transform custom types in JPA entities?

**Answer:**
(Duplicate of Q180).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `@Converter` to transform custom types in JPA entities?
**Your Response:** "I use `@Converter` to transform custom types between database representation and Java objects. I implement `AttributeConverter<X, Y>` where X is my entity type and Y is the database column type. I annotate the converter with `@Converter(autoApply = true)` to apply it automatically. This is perfect for enums, custom date formats, or any domain-specific types that need special handling. The converter handles both the conversion to database format and back to Java format, keeping my entity code clean."

---
