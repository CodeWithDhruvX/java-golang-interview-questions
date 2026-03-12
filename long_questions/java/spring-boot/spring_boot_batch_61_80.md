## 🔹 Section 4: JPA & Database (61–80)

### Question 61: What is Spring Data JPA?

**Answer:**
It is a layer on top of JPA (Java Persistence API).
It dramatically simplifies data access by eliminating boilerplate code. You define an **Interface** extending `JpaRepository`, and Spring automatically generates the implementation at runtime for common CRUD operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Data JPA?
**Your Response:** "Spring Data JPA is a powerful abstraction layer on top of JPA that dramatically simplifies database access. Instead of writing boilerplate DAO code, I simply define an interface that extends `JpaRepository`, and Spring automatically generates the implementation at runtime. This gives me all the standard CRUD operations - save, findById, findAll, delete - without writing any implementation code. Spring Data also provides query derivation, where I can create queries just by defining method names like `findByEmail()` or `findByUsernameAndPassword()`. It eliminates most of the repetitive data access code while still giving me full control when I need custom queries."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define an entity in Spring Boot?
**Your Response:** "Defining a JPA entity in Spring Boot is straightforward. I create a simple POJO and annotate it with `@Entity` to mark it as a database entity. I define the primary key field with `@Id`, and typically use `@GeneratedValue` for auto-generated IDs. Spring Boot with Spring Data JPA automatically detects these entities and maps them to database tables. By default, the class name maps to the table name and field names map to column names, though I can customize this with additional annotations. The entity becomes a first-class citizen in my persistence layer that I can work with through repositories."

---

### Question 63: What is the use of `@Entity`, `@Table`, and `@Id`?

**Answer:**
*   **`@Entity`**: Marks the class as a JPA entity mapped to a DB table.
*   **`@Table`**: Optional. customizing the table name/schema. (`@Table(name="products")`).
*   **`@Id`**: Marks the Primary Key field.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `@Entity`, `@Table`, and `@Id`?
**Your Response:** "These annotations are fundamental for defining JPA entities. `@Entity` marks the class as a JPA entity that maps to a database table. `@Table` is optional and lets me customize the table name or schema, like `@Table(name='products', schema='inventory')`. `@Id` is crucial - it marks the primary key field that uniquely identifies each row. Without `@Id`, JPA won't know how to uniquely identify and track entities. These annotations form the basic mapping between my Java objects and database tables, allowing JPA to handle the object-relational mapping automatically."

---

### Question 64: What is the difference between `CrudRepository`, `JpaRepository`, and `PagingAndSortingRepository`?

**Answer:**
*   **`CrudRepository`**: Basic CRUD (save, findById, delete).
*   **`PagingAndSortingRepository`**: Adds `findAll(Pageable)` and `findAll(Sort)`.
*   **`JpaRepository`**: Extends both. Adds JPA specifics (flush, batch delete). The most commonly used one.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `CrudRepository`, `JpaRepository`, and `PagingAndSortingRepository`?
**Your Response:** "These are different levels of repository interfaces in Spring Data. `CrudRepository` provides the basic CRUD operations - save, findById, findAll, delete, and count. `PagingAndSortingRepository` extends CrudRepository and adds pagination and sorting capabilities with methods like `findAll(Pageable)` and `findAll(Sort)`. `JpaRepository` is the most comprehensive - it extends both previous interfaces and adds JPA-specific features like batch operations, flushing, and entity management. In practice, I almost always use `JpaRepository` because it gives me all the functionality I need while still being simple to use."

---

### Question 65: How do you write custom queries using `@Query` annotation?

**Answer:**
When method name derivation (`findByEmail`) isn't enough.
Uses JPQL (Java Persistence Query Language) by default.

```java
@Query("SELECT u FROM User u WHERE u.active = true")
List<User> findAllActiveUsers();
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write custom queries using `@Query` annotation?
**Your Response:** "When method name derivation isn't sufficient for complex queries, I use the `@Query` annotation. I write JPQL queries, which are similar to SQL but operate on my entity objects rather than database tables. For example, I can write `@Query('SELECT u FROM User u WHERE u.active = true')` to find all active users. JPQL is object-oriented, so I reference entity names and field names rather than table and column names. If I need to use database-specific features or raw SQL, I can set `nativeQuery = true` and write native SQL. This gives me the full power of SQL while keeping my queries organized and type-safe within my repository interfaces."

---

### Question 66: How does Spring Boot handle transactions?

**Answer:**
Uses the **PlatformTransactionManager** (e.g., `JpaTransactionManager` for Hibernate).
It abstracts the specific transaction handling code (Begin/Commit/Rollback).
Spring Boot auto-configures this when `spring-boot-starter-data-jpa` is present.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot handle transactions?
**Your Response:** "Spring Boot handles transactions through the `PlatformTransactionManager`, which abstracts the specific transaction handling code. For JPA with Hibernate, it uses `JpaTransactionManager`. The transaction manager handles the low-level details of beginning, committing, and rolling back transactions. The beauty is that Spring Boot auto-configures this automatically when I include `spring-boot-starter-data-jpa`. I don't need to manually set up transaction management - Spring Boot detects my JPA setup and configures the appropriate transaction manager. I just need to use `@Transactional` annotations, and Spring handles all the transaction boundaries automatically."

---

### Question 67: What is the use of `@Transactional` annotation?

**Answer:**
Wraps a method (or class) in a database transaction.
*   **Success:** Automatically commits.
*   **RuntimeException:** Automatically rolls back.
*   **Checked Exception:** Does NOT rollback by default (configurable via `rollbackFor`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `@Transactional` annotation?
**Your Response:** "The `@Transactional` annotation is the cornerstone of transaction management in Spring Boot. When I annotate a method or class with `@Transactional`, Spring wraps it in a database transaction. If the method completes successfully, Spring automatically commits the transaction. If a `RuntimeException` is thrown, it automatically rolls back. One important detail is that checked exceptions don't trigger rollback by default, though I can configure this with the `rollbackFor` attribute. This annotation makes transaction management declarative rather than programmatic - I just annotate my service methods, and Spring handles all the transaction plumbing behind the scenes."

---

### Question 68: How to enable lazy and eager loading in Spring Boot JPA?

**Answer:**
Defined in the relationship annotation.
*   **Eager:** Fetches association immediately via Join. (`fetch = FetchType.EAGER`).
*   **Lazy (Default for Collections):** Fetches only when accessed. (`fetch = FetchType.LAZY`). Optimal for performance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to enable lazy and eager loading in Spring Boot JPA?
**Your Response:** "JPA provides two fetching strategies that control when related entities are loaded. Eager loading, specified with `fetch = FetchType.EAGER`, fetches related entities immediately using a JOIN query. Lazy loading, the default for collections, fetches related entities only when I actually access them. Lazy loading is generally better for performance because it avoids loading unnecessary data and prevents N+1 query problems. I can configure the fetch type directly in relationship annotations like `@OneToMany(fetch = FetchType.LAZY)`. The key is understanding that lazy loading defers database access until the data is actually needed, which is crucial for efficient database operations."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to use `application.properties` to configure the database?
**Your Response:** "Configuring databases in Spring Boot is done through standard properties in `application.properties`. I specify the JDBC URL with `spring.datasource.url`, credentials with `spring.datasource.username` and `spring.datasource.password`, and JPA-specific settings like `spring.jpa.hibernate.ddl-auto` to control schema generation. Spring Boot's auto-configuration uses these properties to set up the datasource and JPA EntityManager. The nice thing is that different databases use the same property structure - I just change the URL and driver class, and Spring Boot handles the rest. This makes it easy to switch between databases or use different configurations for different environments."

---

### Question 70: How to switch between in-memory DB (H2) and MySQL in Spring Boot?

**Answer:**
Simply change the `pom.xml` dependency (remove `h2`, add `mysql-connector-j`) and update `spring.datasource.*` properties.
Spring Boot auto-configuration detects the driver on the classpath and sets up the dialect/connection.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to switch between in-memory DB (H2) and MySQL in Spring Boot?
**Your Response:** "Switching databases in Spring Boot is surprisingly simple. I just need to change the dependency in my `pom.xml` - remove the H2 dependency and add the MySQL connector dependency. Then I update the `spring.datasource.*` properties to point to the MySQL database. Spring Boot's auto-configuration automatically detects the database driver on the classpath and configures the appropriate Hibernate dialect and connection settings. This makes it easy to use H2 for development and testing, then switch to MySQL or PostgreSQL for production without changing any application code - just the dependencies and configuration properties."

---

### Question 71: What is the role of `EntityManager`?

**Answer:**
It is the core interface of JPA. It manages the persistence context (Session).
`JpaRepository` uses it internally.
You can inject it directly (`@PersistenceContext`) to perform complex operations not supported by Repository methods (e.g., Detaching objects, Clear cache).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `EntityManager`?
**Your Response:** "The `EntityManager` is the core JPA interface that manages the persistence context - essentially the session that tracks entity state. While `JpaRepository` uses the EntityManager internally, I can inject it directly using `@PersistenceContext` for advanced operations that repository methods don't support. The EntityManager handles entity lifecycle operations like persist, merge, and remove, and manages the first-level cache. I use it directly when I need to perform complex queries, detach entities manually, or clear the persistence cache. It gives me fine-grained control over entity management when the higher-level repository abstractions aren't sufficient."

---

### Question 72: What is optimistic vs pessimistic locking in JPA?

**Answer:**
*   **Optimistic:** Uses a `@Version` field. Checks version on Update. If changed by another, throws `OptimisticLockException`. No DB locks.
*   **Pessimistic:** Locks the DB row (`SELECT ... FOR UPDATE`). Prevents others from reading/writing until transaction ends.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is optimistic vs pessimistic locking in JPA?
**Your Response:** "These are two strategies for handling concurrent access to data. Optimistic locking assumes conflicts are rare, so it doesn't lock database rows. Instead, I add a `@Version` field to my entity, and JPA checks this version on updates. If another transaction modified the data, JPA throws an `OptimisticLockException`. Pessimistic locking locks the database row using `SELECT ... FOR UPDATE`, preventing other transactions from reading or writing until the transaction ends. Optimistic locking is better for high-concurrency applications because it doesn't block other users, while pessimistic locking is simpler but can cause performance bottlenecks."

---

### Question 73: What are native queries in Spring Boot JPA?

**Answer:**
Raw SQL queries. Used for DB-specific features not supported by JPQL.
```java
@Query(value = "SELECT * FROM users WHERE email = ?1", nativeQuery = true)
User findByEmailNative(String email);
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are native queries in Spring Boot JPA?
**Your Response:** "Native queries are raw SQL queries that I can use when JPQL isn't sufficient or when I need to use database-specific features. I create them with the `@Query` annotation and set `nativeQuery = true`. This tells JPA to execute the SQL directly on the database rather than translating JPQL. Native queries are useful for complex operations, database-specific functions, or when I need to optimize performance with hand-tuned SQL. The trade-off is that I lose some of JPQL's database portability, but I gain access to the full power of my specific database. I still get the result mapping to entity objects, which maintains the object-oriented benefits."

---

### Question 74: How to implement OneToOne, OneToMany, ManyToOne relationships?

**Answer:**
Use relationship annotations.
*   `@OneToMany(mappedBy="user")`: Parent side (List).
*   `@ManyToOne`: Child side (Foreign Key). owner of relationship.
*   `@OneToOne`: Single link.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to implement OneToOne, OneToMany, ManyToOne relationships?
**Your Response:** "I implement JPA relationships using specific annotations. For a one-to-many relationship, I use `@OneToMany(mappedBy='user')` on the parent side, where `mappedBy` indicates that the child entity owns the relationship. For the child side, I use `@ManyToOne` which defines the foreign key column and makes it the relationship owner. For one-to-one relationships, I use `@OneToOne` on both sides. The key concept is relationship ownership - the side without `mappedBy` owns the relationship and defines the foreign key. This maps directly to database foreign key relationships and lets JPA handle the joins and cascading operations automatically."

---

### Question 75: How to fetch child entities with parent using joins in JPA?

**Answer:**
1.  **JPQL Fetch Join:** `@Query("SELECT u FROM User u JOIN FETCH u.orders")`. Solves N+1 problem.
2.  **Entity Graph:** `@EntityGraph(attributePaths = {"orders"})`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to fetch child entities with parent using joins in JPA?
**Your Response:** "To solve the N+1 query problem when fetching related entities, I have two main approaches. The first is JPQL fetch joins - I write queries like `SELECT u FROM User u JOIN FETCH u.orders` which tells JPA to fetch the orders in the same query using a JOIN. The second approach is using Entity Graphs with `@EntityGraph(attributePaths = {'orders'})` which defines a template of what relationships to fetch. Both approaches prevent the N+1 problem by loading all needed data in a single query rather than separate queries for each relationship. Fetch joins are more flexible while Entity Graphs are more declarative and reusable."

---

### Question 76: What is `CascadeType` in JPA?

**Answer:**
Propagates operations from Parent to Child.
*   `PERSIST`: Save parent -> Save child.
*   `remove`: Delete parent -> Delete child.
*   `ALL`: All operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `CascadeType` in JPA?
**Your Response:** "CascadeType controls how operations propagate from parent entities to child entities. When I specify `CascadeType.PERSIST`, saving a parent automatically saves its children. `CascadeType.REMOVE` means deleting a parent also deletes its children. `CascadeType.ALL` includes all operations including persist, merge, remove, and refresh. This is useful for parent-child relationships where I want to treat them as a single unit - when I work with the parent, JPA automatically handles the children. The key is understanding that cascading means the same operation applied to the parent is also applied to the children, which simplifies entity management but requires careful consideration to avoid unintended side effects."

---

### Question 77: How do you handle schema generation in Spring Boot?

**Answer:**
Property: `spring.jpa.hibernate.ddl-auto`.
*   `create`: Drop and create table on startup.
*   `update`: Update schema (add columns). Safe-ish for dev.
*   `validate`: Verify DB matches Entity.
*   `none`: Do nothing (Production).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle schema generation in Spring Boot?
**Your Response:** "I control schema generation through the `spring.jpa.hibernate.ddl-auto` property. For development, I often use `update` which automatically updates the schema by adding new columns without dropping existing data. `create` drops and recreates tables on each startup, which is good for testing but destructive. `validate` checks that the database schema matches my entity definitions and fails if there are mismatches. For production, I always set it to `none` and use database migration tools like Flyway or Liquibase instead. The key is using different strategies for different environments - automatic schema generation for development, and controlled migrations for production."

---

### Question 78: What is Flyway or Liquibase and how do you use it with Spring Boot?

**Answer:**
**Database Migration Tools.** Essential for production.
Versioning for DB Schema (like Git for SQL).
Boot auto-detects them.
Create files like `V1__init.sql` in `src/main/resources/db/migration`. Boot runs them on startup.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Flyway or Liquibase and how do you use it with Spring Boot?
**Your Response:** "Flyway and Liquibase are database migration tools that provide version control for database schema changes - essentially Git for my database schema. This is essential for production environments where I need to manage schema changes reliably. Spring Boot auto-detects these tools and integrates them seamlessly. I create SQL migration files with version numbers like `V1__init.sql` in the `src/main/resources/db/migration` directory, and Spring Boot automatically runs them in order on startup. This ensures that every environment has the exact same database schema, and I can track changes over time. It's the professional way to manage database evolution in production applications."

---

### Question 79: How to log SQL queries in Spring Boot?

**Answer:**
```properties
spring.jpa.show-sql=true
# Use logger for formatted output
logging.level.org.hibernate.SQL=DEBUG
logging.level.org.hibernate.type.descriptor.sql.BasicBinder=TRACE # Show parameters
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to log SQL queries in Spring Boot?
**Your Response:** "I can enable SQL logging in Spring Boot through simple properties. Setting `spring.jpa.show-sql=true` makes Hibernate print SQL statements to the console. For more detailed logging with proper formatting, I configure the logging levels: `logging.level.org.hibernate.SQL=DEBUG` shows the SQL statements, and `logging.level.org.hibernate.type.descriptor.sql.BasicBinder=TRACE` shows the parameter values. This is incredibly useful for debugging - I can see exactly what SQL Hibernate is generating and what values are being bound. In production, I'd typically disable this for performance and security, but it's invaluable during development and troubleshooting."

---

### Question 80: How to test a Spring Boot repository with `@DataJpaTest`?

**Answer:**
Slice Test annotation.
*   Configures an **in-memory** DB (H2).
*   Scans only `@Entity` and `@Repository`.
*   Rolls back transactions after each test (Clean state).
Fast isolation testing for DAL.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to test a Spring Boot repository with `@DataJpaTest`?
**Your Response:** "`@DataJpaTest` is a specialized test annotation for testing the data access layer. It configures an in-memory database like H2 automatically, scans only `@Entity` and `@Repository` beans, and rolls back transactions after each test to maintain a clean state. This makes my repository tests fast and isolated since they don't load the entire Spring application context. The in-memory database means tests run quickly without external dependencies, and the automatic rollback ensures tests don't interfere with each other. It's the perfect way to test my JPA repositories and queries in isolation without the overhead of a full integration test."

---
