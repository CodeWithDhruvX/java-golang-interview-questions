## 🔹 Section 4: JPA, Queries & Data Layer (161-180)

### Question 161: How to handle null-safe queries in Spring Data JPA?

**Answer:**
Use **QueryDSL** or **Specifications**.
Standard JPQL string concatenation is prone to errors if parameters are null.
Specifications allow composing predicates: `if (name != null) spec.and(hasName(name))`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle null-safe queries in Spring Data JPA?
**Your Response:** "For null-safe dynamic queries, I use QueryDSL or Specifications rather than string concatenation which is error-prone. Specifications allow me to build queries programmatically by composing predicates. For example, I can write `if (name != null) spec.and(hasName(name))` to conditionally add criteria only when parameters are not null. This approach is type-safe, prevents SQL injection, and handles null values elegantly. Specifications are composable, so I can build complex queries by combining multiple conditions, and Spring Data translates them to proper SQL while handling null values safely."

---

### Question 162: What is Specification API and how is it used in Spring Boot?

**Answer:**
Allows building dynamic queries programmatically.
Repository must extend `JpaSpecificationExecutor`.
`findAll((root, query, cb) -> cb.equal(root.get("status"), "ACTIVE"))`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Specification API and how is it used in Spring Boot?
**Your Response:** "The Specification API allows me to build dynamic queries programmatically in a type-safe way. I extend my repository with `JpaSpecificationExecutor` to get access to methods like `findAll(Specification)`. I create specifications using lambda expressions that work with the Criteria API - for example, `findAll((root, query, cb) -> cb.equal(root.get('status'), 'ACTIVE'))`. This approach is perfect for building complex search functionality where users can filter by multiple optional criteria. The specifications are reusable and composable, making my query logic clean and maintainable."

---

### Question 163: How to implement soft deletes using JPA?

**Answer:**
1.  Add field `deleted` (boolean) to Entity.
2.  Add `@SQLDelete(sql = "UPDATE table SET deleted = true WHERE id = ?")`.
3.  Add `@Where(clause = "deleted = false")`.
Now `repo.delete()` runs the Update, and `repo.findAll()` only sees active rows.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to implement soft deletes using JPA?
**Your Response:** "I implement soft deletes by adding a boolean `deleted` field to my entity and using Hibernate-specific annotations. I add `@SQLDelete(sql = 'UPDATE table SET deleted = true WHERE id = ?')` to override the delete operation with an update instead. I also add `@Where(clause = 'deleted = false')` to automatically filter out deleted records from queries. With this setup, when I call `repo.delete()`, it performs a soft delete by setting the flag, and `repo.findAll()` only returns active records. This preserves data history while making deleted records invisible to the application."

---

### Question 164: How do you audit entity changes in Spring Boot?

**Answer:**
Use **Spring Data Envers** or **Auditing**.
1.  Enable: `@EnableJpaAuditing`.
2.  Entity Fields: `@CreatedDate`, `@LastModifiedDate`, `@CreatedBy`.
3.  Spring automatically populates timestamps on save/update.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you audit entity changes in Spring Boot?
**Your Response:** "I use Spring Data Envers or the built-in Auditing feature for entity auditing. I enable auditing with `@EnableJpaAuditing` and add annotations like `@CreatedDate`, `@LastModifiedDate`, and `@CreatedBy` to my entity fields. Spring automatically populates these fields when entities are saved or updated. For more comprehensive auditing, Spring Data Envers tracks complete change history with revision information. This automatic auditing is invaluable for compliance, debugging, and understanding who changed what and when in the system."

---

### Question 165: What are entity graphs and how do you use them in JPA?

**Answer:**
Solves N+1 loading problem selectively.
Define `@NamedEntityGraph` on Entity.
In Repository: `@EntityGraph(value = "graph.User.orders")`.
Overrides the default Lazy loading fetch plan for that specific query.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are entity graphs and how do you use them in JPA?
**Your Response:** "Entity graphs solve the N+1 loading problem by allowing me to selectively override the default lazy loading fetch plan. I define a `@NamedEntityGraph` on my entity specifying which relationships to fetch eagerly. Then in my repository, I use `@EntityGraph(value = 'graph.User.orders')` to apply that graph to specific queries. This gives me fine-grained control - I can fetch related data in one query only when needed, rather than always using eager loading or suffering from N+1 problems. Entity graphs are more declarative and reusable than fetch joins in JPQL."

---

### Question 166: How do you use stored procedures in Spring Boot JPA?

**Answer:**
`@Procedure` annotation in Repository.
```java
@Procedure("calculatetax")
int calculateTax(int userId);
```
Or `@NamedStoredProcedureQuery` on Entity.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use stored procedures in Spring Boot JPA?
**Your Response:** "I can call stored procedures using the `@Procedure` annotation directly in my repository methods. For example, `@Procedure('calculatetax') int calculateTax(int userId)` maps to a stored procedure. Alternatively, I can use `@NamedStoredProcedureQuery` on the entity class for more complex procedures. This approach allows me to leverage database-specific logic while maintaining a clean API in my Spring Boot application. The stored procedure calls are type-safe and integrate seamlessly with Spring Data's repository pattern, making them easy to test and maintain."

---

### Question 167: What is the role of `JpaSpecificationExecutor`?

**Answer:**
Interface definition that adds methods accepting `Specification<T>` to your repository (`findOne(Spec)`, `findAll(Spec)`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `JpaSpecificationExecutor`?
**Your Response:** "`JpaSpecificationExecutor` is an interface that extends my repository with methods for executing dynamic queries using Specifications. When I extend this interface, I get methods like `findOne(Specification)` and `findAll(Specification)` that allow me to build complex, type-safe queries programmatically. This is particularly useful for search functionality where users can filter by multiple optional criteria. The interface bridges Spring Data repositories with the JPA Criteria API, giving me the power of dynamic queries while maintaining the clean repository abstraction."

---

### Question 168: How to manage complex joins using Criteria API?

**Answer:**
Criteria API is verbose but typesafe.
Using `CriteriaBuilder`, you create a `Join` object from Root.
`Join<User, Order> orders = root.join("orders");`
`query.where(cb.gt(orders.get("total"), 100))`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to manage complex joins using Criteria API?
**Your Response:** "The Criteria API provides a type-safe way to build complex queries with joins. I use `CriteriaBuilder` to create `Join` objects from the root entity, like `Join<User, Order> orders = root.join('orders')`. Then I can apply conditions to the joined entities using `query.where(cb.gt(orders.get('total'), 100))`. While the Criteria API is verbose, it gives me compile-time safety and works well for dynamic queries where the structure isn't known until runtime. It's especially useful for building search APIs with optional filters and complex relationships."

---

### Question 169: What is the default transaction propagation behavior in Spring Boot?

**Answer:**
`Propagation.REQUIRED`.
If a transaction exists, join it.
If not, create a new one.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the default transaction propagation behavior in Spring Boot?
**Your Response:** "The default transaction propagation is `Propagation.REQUIRED`. This means if a transaction already exists when a transactional method is called, the method joins the existing transaction. If no transaction exists, Spring creates a new one. This 'join existing or create new' behavior works well for most scenarios where I want multiple service methods to participate in the same transaction. It ensures data consistency across method boundaries while allowing individual methods to work independently when called outside of a transactional context."

---

### Question 170: How can you implement multi-tenancy with Spring Boot JPA?

**Answer:**
1.  **Discriminator Column:** Filter all queries by `tenant_id`. (Hibernate Filters).
2.  **Schema per Tenant:** Switch JDBC connection `currentSchema` property based on Request Context.
3.  **DB per Tenant:** Use `AbstractRoutingDataSource` to pick DataSource dynamically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can you implement multi-tenancy with Spring Boot JPA?
**Your Response:** "I have several approaches for multi-tenancy. The discriminator column approach uses a `tenant_id` column and Hibernate filters to automatically filter queries by tenant. Schema-per-tenant switches the database schema based on the current tenant context. Database-per-tenant uses `AbstractRoutingDataSource` to dynamically route to different databases based on the tenant. Each approach has different trade-offs - discriminator is simplest but shares database resources, schema-per-tenant provides better isolation, and database-per-tenant provides complete isolation but requires more infrastructure. I choose based on security requirements and scalability needs."

---

### Question 171: How do you use `@Query` with native SQL in Spring Data JPA?

**Answer:**
(See Q73). `nativeQuery=true`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `@Query` with native SQL in Spring Data JPA?
**Your Response:** "When I need to use database-specific features or optimize performance, I use native SQL queries with `@Query`. I set `nativeQuery=true` and write raw SQL instead of JPQL. This gives me access to database-specific functions, hints, or optimizations that JPQL doesn't support. The trade-off is losing database portability, but I gain the full power of my specific database. Spring Data still maps the results to my entities, so I maintain the object-oriented benefits while using optimized SQL when needed."

---

### Question 172: What is the difference between `EntityManager.merge()` and `save()`?

**Answer:**
- **`save()`:** (Spring Data). Checks if ID is new. Calls `persist` (Insert) or `merge` (Update).
- **`merge()`:** (JPA). Copies state of a detached object into the persistence context. Returns a new managed instance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `EntityManager.merge()` and `save()`?
**Your Response:** "`save()` is Spring Data's convenience method that checks if the entity's ID is null to determine whether to persist (insert) or merge (update). `merge()` is the standard JPA method that copies the state of a detached object into the persistence context and returns a managed instance. The key difference is that `merge()` always returns a managed instance, while `save()` may return the original instance. I use `save()` for simplicity in most cases, but use `merge()` when I need to work with detached entities and want to ensure I'm working with the managed context."

---

### Question 173: How do you implement cascading deletes in JPA?

**Answer:**
`@OneToMany(cascade = CascadeType.REMOVE, orphanRemoval = true)`.
Deleting Parent deletes Children.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement cascading deletes in JPA?
**Your Response:** "I implement cascading deletes using `@OneToMany(cascade = CascadeType.REMOVE, orphanRemoval = true)`. When I delete a parent entity, this automatically deletes all child entities. The `orphanRemoval = true` part handles cases where I remove a child from the parent's collection - it automatically deletes the orphaned child from the database. This cascading behavior maintains referential integrity and simplifies my code - I only need to call delete on the parent, and JPA handles the cleanup of related entities automatically."

---

### Question 174: How do you use database views in Spring Boot JPA?

**Answer:**
Treat the View exactly like a Table.
Create an `@Entity` mapped to `@Table(name="my_view")`.
Should be marked `@Immutable` so Hibernate doesn't try to Update it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use database views in Spring Boot JPA?
**Your Response:** "I treat database views exactly like tables in JPA. I create an `@Entity` class mapped to the view using `@Table(name='my_view')`. Since views are read-only, I mark the entity with `@Immutable` to prevent Hibernate from trying to update it. This allows me to query complex database views using Spring Data repositories while maintaining the same programming model as regular entities. The entity becomes a read-only projection of the view data, which is perfect for reporting or read-only operations that involve complex joins or aggregations."

---

### Question 175: How do you map stored procedures using Spring Boot JPA?

**Answer:**
(See 166).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you map stored procedures using Spring Boot JPA?
**Your Response:** "I map stored procedures using either the `@Procedure` annotation in repository methods or `@NamedStoredProcedureQuery` on entities. The `@Procedure` approach is simpler for basic procedures - I just annotate a repository method with the procedure name and parameters. For more complex procedures with multiple parameters or result sets, I use `@NamedStoredProcedureQuery` on the entity class. Both approaches allow me to call database procedures while maintaining the clean Spring Data repository abstraction, making stored procedures feel like regular repository methods."

---

### Question 176: How do you implement pagination using `Slice` and `Page`?

**Answer:**
- **`Page<T>`:** Returns List + Total Count (runs an extra `COUNT(*)` query). Expensive.
- **`Slice<T>`:** Returns List + `hasNext()` flag. (One query with limit+1). Cheaper. Use for infinite scroll.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement pagination using `Slice` and `Page`?
**Your Response:** "`Page<T>` and `Slice<T>` serve different pagination needs. `Page<T>` returns the data list plus total count information, but it requires an extra `COUNT(*)` query which can be expensive on large datasets. `Slice<T>` is more efficient - it returns the data list plus a `hasNext()` flag using only one query with limit+1. I use `Slice<T>` for infinite scroll scenarios where I just need to know if there's more data, and `Page<T>` when I need total page counts for pagination controls. The choice depends on whether I need total count information or just next-page detection."

---

### Question 177: How do you create and use custom repository implementations in Spring Data?

**Answer:**
1.  Define Interface `UserCustomRepo`.
2.  Implement it `UserCustomRepoImpl` (must end with `Impl`).
3.  Inject `EntityManager` and write custom logic.
4.  Make standard `UserRepository` extend `UserCustomRepo`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create and use custom repository implementations in Spring Data?
**Your Response:** "When Spring Data's repository methods aren't sufficient, I create custom repository implementations. I define a separate interface like `UserCustomRepo`, implement it in a class named `UserCustomRepoImpl` (the `Impl` suffix is required), and inject `EntityManager` to write custom logic. Then I make my standard `UserRepository` extend the custom interface. Spring Data automatically detects and uses the custom implementation. This gives me the best of both worlds - the convenience of Spring Data for standard operations and complete flexibility for custom queries when needed."

---

### Question 178: What’s the role of `JpaSpecificationExecutor` in Spring Data JPA?

**Answer:**
(See 167).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the role of `JpaSpecificationExecutor` in Spring Data JPA?
**Your Response:** "`JpaSpecificationExecutor` extends my repository interface with methods for executing dynamic, type-safe queries using the JPA Criteria API. When I extend this interface, I get access to methods like `findAll(Specification)` and `findOne(Specification)` that allow me to build complex queries programmatically. This is particularly useful for search functionality where query criteria are determined at runtime based on user input. The interface bridges the gap between Spring Data's simple repository methods and the full power of JPA's Criteria API, giving me dynamic querying capabilities while maintaining the repository pattern."

---

### Question 179: How do you use `@SqlResultSetMapping` for native queries?

**Answer:**
Used to map Native Query results to complex Non-Entity DTOs.
Defined on an Entity class.
Pass the mapping name to `.createNativeQuery(sql, "MappingName")` in EntityManager.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `@SqlResultSetMapping` for native queries?
**Your Response:** "I use `@SqlResultSetMapping` to map native query results to complex non-entity DTOs. I define the mapping on an entity class using annotations that specify how to map columns to DTO fields. Then in my EntityManager, I call `createNativeQuery(sql, 'MappingName')` to execute the native query with the mapping. This is perfect when I need to return complex projections that don't map directly to entities, like reporting queries with calculated fields or joins across multiple tables. The mapping gives me type-safe results while leveraging the database's full query capabilities."

---

### Question 180: How do you use `@Converter` to transform custom types in JPA entities?

**Answer:**
Implement `AttributeConverter<MyObj, String>`.
Annotate with `@Converter(autoApply=true)`.
Useful for encrypting fields or converting Enums/JSON to String for DB storage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `@Converter` to transform custom types in JPA entities?
**Your Response:** "I implement `AttributeConverter<MyObj, String>` and annotate it with `@Converter(autoApply=true)` to create custom type converters. This allows me to transform complex objects to database-compatible types. For example, I can encrypt sensitive fields before storing them, convert enums to their string representations, or serialize JSON objects to strings for database storage. The `autoApply=true` makes JPA automatically use this converter for fields of the specified type throughout the application. This is a powerful way to handle custom data types while keeping the entity code clean and the database schema simple."

---
