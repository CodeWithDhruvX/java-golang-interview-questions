## 🔹 Section 4: JPA, Queries & Data Layer (161-180)

### Question 161: How to handle null-safe queries in Spring Data JPA?

**Answer:**
Use **QueryDSL** or **Specifications**.
Standard JPQL string concatenation is prone to errors if parameters are null.
Specifications allow composing predicates: `if (name != null) spec.and(hasName(name))`.

---

### Question 162: What is Specification API and how is it used in Spring Boot?

**Answer:**
Allows building dynamic queries programmatically.
Repository must extend `JpaSpecificationExecutor`.
`findAll((root, query, cb) -> cb.equal(root.get("status"), "ACTIVE"))`.

---

### Question 163: How to implement soft deletes using JPA?

**Answer:**
1.  Add field `deleted` (boolean) to Entity.
2.  Add `@SQLDelete(sql = "UPDATE table SET deleted = true WHERE id = ?")`.
3.  Add `@Where(clause = "deleted = false")`.
Now `repo.delete()` runs the Update, and `repo.findAll()` only sees active rows.

---

### Question 164: How do you audit entity changes in Spring Boot?

**Answer:**
Use **Spring Data Envers** or **Auditing**.
1.  Enable: `@EnableJpaAuditing`.
2.  Entity Fields: `@CreatedDate`, `@LastModifiedDate`, `@CreatedBy`.
3.  Spring automatically populates timestamps on save/update.

---

### Question 165: What are entity graphs and how do you use them in JPA?

**Answer:**
Solves N+1 loading problem selectively.
Define `@NamedEntityGraph` on Entity.
In Repository: `@EntityGraph(value = "graph.User.orders")`.
Overrides the default Lazy loading fetch plan for that specific query.

---

### Question 166: How do you use stored procedures in Spring Boot JPA?

**Answer:**
`@Procedure` annotation in Repository.
```java
@Procedure("calculatetax")
int calculateTax(int userId);
```
Or `@NamedStoredProcedureQuery` on Entity.

---

### Question 167: What is the role of `JpaSpecificationExecutor`?

**Answer:**
Interface definition that adds methods accepting `Specification<T>` to your repository (`findOne(Spec)`, `findAll(Spec)`).

---

### Question 168: How to manage complex joins using Criteria API?

**Answer:**
Criteria API is verbose but typesafe.
Using `CriteriaBuilder`, you create a `Join` object from Root.
`Join<User, Order> orders = root.join("orders");`
`query.where(cb.gt(orders.get("total"), 100))`.

---

### Question 169: What is the default transaction propagation behavior in Spring Boot?

**Answer:**
`Propagation.REQUIRED`.
If a transaction exists, join it.
If not, create a new one.

---

### Question 170: How can you implement multi-tenancy with Spring Boot JPA?

**Answer:**
1.  **Discriminator Column:** Filter all queries by `tenant_id`. (Hibernate Filters).
2.  **Schema per Tenant:** Switch JDBC connection `currentSchema` property based on Request Context.
3.  **DB per Tenant:** Use `AbstractRoutingDataSource` to pick DataSource dynamically.

---

### Question 171: How do you use `@Query` with native SQL in Spring Data JPA?

**Answer:**
(See Q73). `nativeQuery=true`.

---

### Question 172: What is the difference between `EntityManager.merge()` and `save()`?

**Answer:**
- **`save()`:** (Spring Data). Checks if ID is new. Calls `persist` (Insert) or `merge` (Update).
- **`merge()`:** (JPA). Copies state of a detached object into the persistence context. Returns a new managed instance.

---

### Question 173: How do you implement cascading deletes in JPA?

**Answer:**
`@OneToMany(cascade = CascadeType.REMOVE, orphanRemoval = true)`.
Deleting Parent deletes Children.

---

### Question 174: How do you use database views in Spring Boot JPA?

**Answer:**
Treat the View exactly like a Table.
Create an `@Entity` mapped to `@Table(name="my_view")`.
Should be marked `@Immutable` so Hibernate doesn't try to Update it.

---

### Question 175: How do you map stored procedures using Spring Boot JPA?

**Answer:**
(See 166).

---

### Question 176: How do you implement pagination using `Slice` and `Page`?

**Answer:**
- **`Page<T>`:** Returns List + Total Count (runs an extra `COUNT(*)` query). Expensive.
- **`Slice<T>`:** Returns List + `hasNext()` flag. (One query with limit+1). Cheaper. Use for infinite scroll.

---

### Question 177: How do you create and use custom repository implementations in Spring Data?

**Answer:**
1.  Define Interface `UserCustomRepo`.
2.  Implement it `UserCustomRepoImpl` (must end with `Impl`).
3.  Inject `EntityManager` and write custom logic.
4.  Make standard `UserRepository` extend `UserCustomRepo`.

---

### Question 178: What’s the role of `JpaSpecificationExecutor` in Spring Data JPA?

**Answer:**
(See 167).

---

### Question 179: How do you use `@SqlResultSetMapping` for native queries?

**Answer:**
Used to map Native Query results to complex Non-Entity DTOs.
Defined on an Entity class.
Pass the mapping name to `.createNativeQuery(sql, "MappingName")` in EntityManager.

---

### Question 180: How do you use `@Converter` to transform custom types in JPA entities?

**Answer:**
Implement `AttributeConverter<MyObj, String>`.
Annotate with `@Converter(autoApply=true)`.
Useful for encrypting fields or converting Enums/JSON to String for DB storage.

---
