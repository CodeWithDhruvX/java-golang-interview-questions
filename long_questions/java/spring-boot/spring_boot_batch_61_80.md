## 🔹 Section 4: JPA & Database (61–80)

### Question 61: What is Spring Data JPA?

**Answer:**
It is a layer on top of JPA (Java Persistence API).
It dramatically simplifies data access by eliminating boilerplate code. You define an **Interface** extending `JpaRepository`, and Spring automatically generates the implementation at runtime for common CRUD operations.

---

### Question 62: How do you define an entity in Spring Boot?

**Answer:**
Create a simple Java class (POJO).
Annotate it with `@Entity`.
Define a primary key with `@Id`.

```java
@Entity
public class Product {
    @Id @GeneratedValue
    private Long id;
    private String name;
}
```

---

### Question 63: What is the use of `@Entity`, `@Table`, and `@Id`?

**Answer:**
*   **`@Entity`**: Marks the class as a JPA entity mapped to a DB table.
*   **`@Table`**: Optional. customizing the table name/schema. (`@Table(name="products")`).
*   **`@Id`**: Marks the Primary Key field.

---

### Question 64: What is the difference between `CrudRepository`, `JpaRepository`, and `PagingAndSortingRepository`?

**Answer:**
*   **`CrudRepository`**: Basic CRUD (save, findById, delete).
*   **`PagingAndSortingRepository`**: Adds `findAll(Pageable)` and `findAll(Sort)`.
*   **`JpaRepository`**: Extends both. Adds JPA specifics (flush, batch delete). The most commonly used one.

---

### Question 65: How do you write custom queries using `@Query` annotation?

**Answer:**
When method name derivation (`findByEmail`) isn't enough.
Uses JPQL (Java Persistence Query Language) by default.

```java
@Query("SELECT u FROM User u WHERE u.active = true")
List<User> findAllActiveUsers();
```

---

### Question 66: How does Spring Boot handle transactions?

**Answer:**
Uses the **PlatformTransactionManager** (e.g., `JpaTransactionManager` for Hibernate).
It abstracts the specific transaction handling code (Begin/Commit/Rollback).
Spring Boot auto-configures this when `spring-boot-starter-data-jpa` is present.

---

### Question 67: What is the use of `@Transactional` annotation?

**Answer:**
Wraps a method (or class) in a database transaction.
*   **Success:** Automatically commits.
*   **RuntimeException:** Automatically rolls back.
*   **Checked Exception:** Does NOT rollback by default (configurable via `rollbackFor`).

---

### Question 68: How to enable lazy and eager loading in Spring Boot JPA?

**Answer:**
Defined in the relationship annotation.
*   **Eager:** Fetches association immediately via Join. (`fetch = FetchType.EAGER`).
*   **Lazy (Default for Collections):** Fetches only when accessed. (`fetch = FetchType.LAZY`). Optimal for performance.

---

### Question 69: How to use `application.properties` to configure the database?

**Answer:**
Standard properties:
```properties
spring.datasource.url=jdbc:mysql://localhost:3306/mydb
spring.datasource.username=root
spring.datasource.password=secret
spring.jpa.hibernate.ddl-auto=update
```

---

### Question 70: How to switch between in-memory DB (H2) and MySQL in Spring Boot?

**Answer:**
Simply change the `pom.xml` dependency (remove `h2`, add `mysql-connector-j`) and update `spring.datasource.*` properties.
Spring Boot auto-configuration detects the driver on the classpath and sets up the dialect/connection.

---

### Question 71: What is the role of `EntityManager`?

**Answer:**
It is the core interface of JPA. It manages the persistence context (Session).
`JpaRepository` uses it internally.
You can inject it directly (`@PersistenceContext`) to perform complex operations not supported by Repository methods (e.g., Detaching objects, Clear cache).

---

### Question 72: What is optimistic vs pessimistic locking in JPA?

**Answer:**
*   **Optimistic:** Uses a `@Version` field. Checks version on Update. If changed by another, throws `OptimisticLockException`. No DB locks.
*   **Pessimistic:** Locks the DB row (`SELECT ... FOR UPDATE`). Prevents others from reading/writing until transaction ends.

---

### Question 73: What are native queries in Spring Boot JPA?

**Answer:**
Raw SQL queries. Used for DB-specific features not supported by JPQL.
```java
@Query(value = "SELECT * FROM users WHERE email = ?1", nativeQuery = true)
User findByEmailNative(String email);
```

---

### Question 74: How to implement OneToOne, OneToMany, ManyToOne relationships?

**Answer:**
Use relationship annotations.
*   `@OneToMany(mappedBy="user")`: Parent side (List).
*   `@ManyToOne`: Child side (Foreign Key). owner of relationship.
*   `@OneToOne`: Single link.

---

### Question 75: How to fetch child entities with parent using joins in JPA?

**Answer:**
1.  **JPQL Fetch Join:** `@Query("SELECT u FROM User u JOIN FETCH u.orders")`. Solves N+1 problem.
2.  **Entity Graph:** `@EntityGraph(attributePaths = {"orders"})`.

---

### Question 76: What is `CascadeType` in JPA?

**Answer:**
Propagates operations from Parent to Child.
*   `PERSIST`: Save parent -> Save child.
*   `remove`: Delete parent -> Delete child.
*   `ALL`: All operations.

---

### Question 77: How do you handle schema generation in Spring Boot?

**Answer:**
Property: `spring.jpa.hibernate.ddl-auto`.
*   `create`: Drop and create table on startup.
*   `update`: Update schema (add columns). Safe-ish for dev.
*   `validate`: Verify DB matches Entity.
*   `none`: Do nothing (Production).

---

### Question 78: What is Flyway or Liquibase and how do you use it with Spring Boot?

**Answer:**
**Database Migration Tools.** Essential for production.
Versioning for DB Schema (like Git for SQL).
Boot auto-detects them.
Create files like `V1__init.sql` in `src/main/resources/db/migration`. Boot runs them on startup.

---

### Question 79: How to log SQL queries in Spring Boot?

**Answer:**
```properties
spring.jpa.show-sql=true
# Use logger for formatted output
logging.level.org.hibernate.SQL=DEBUG
logging.level.org.hibernate.type.descriptor.sql.BasicBinder=TRACE # Show parameters
```

---

### Question 80: How to test a Spring Boot repository with `@DataJpaTest`?

**Answer:**
Slice Test annotation.
*   Configures an **in-memory** DB (H2).
*   Scans only `@Entity` and `@Repository`.
*   Rolls back transactions after each test (Clean state).
Fast isolation testing for DAL.

---
