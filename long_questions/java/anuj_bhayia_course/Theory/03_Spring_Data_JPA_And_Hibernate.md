# Spring Data JPA and Hibernate - Interview Questions and Answers

## 1. What is JPA and what is Hibernate? Relation between them?
**Answer:**
- **JPA (Jakarta Persistence API, formerly Java Persistence API):** It is a Java specification for accessing, persisting, and managing data between Java objects (POJOs) and a relational database. It is only a specification; it defines the interfaces (like `EntityManager`, `Query`) and annotations (like `@Entity`, `@Id`, `@OneToMany`) but does not provide the actual implementation.
- **Hibernate ORM:** It is a mature, open-source Object-Relational Mapping (ORM) framework that maps Java objects to relational database tables and Java data types to SQL data types.
- **Relation:** Hibernate is the most popular, default implementation of the JPA specification. While you program against the JPA interfaces and use JPA annotations, it is Hibernate working under the hood in a Spring Boot application to generate and execute the actual SQL queries against your database. You can also use Hibernate-specific features (features not defined in JPA), but sticking purely to JPA ensures you can swap out Hibernate for another provider (like EclipseLink) in the future if needed.

## 2. What is Spring Data JPA?
**Answer:**
Spring Data JPA is a part of the larger Spring Data family. It is not an ORM itself; rather, it provides a layer of abstraction on top of JPA providers (like Hibernate). It significantly reduces the amount of boilerplate code required to implement data access layers (DAOs).

Instead of writing custom `JpaTemplate` or managing the `EntityManager` to perform simple CRUD operations or queries, Spring Data JPA allows developers to simply define repository interfaces extending predefined Spring interfaces (like `JpaRepository` or `CrudRepository`). Spring automatically creates proxies for these interfaces at runtime and implements standard database operations dynamically.

## 3. What are the key annotations used in JPA Entity mapping?
**Answer:**
- **`@Entity`:** Marks a plain Java class as a JPA entity, meaning its instances will be mapped to a database table.
- **`@Table`:** Optional. Used to specify the exact table name in the database if it differs from the class name, or to define unique constraints.
- **`@Id`:** Marks a field as the primary key of the entity. Every entity must have an `@Id`.
- **`@GeneratedValue`:** Specifies how the primary key should be automatically generated (e.g., `GenerationType.IDENTITY` for auto-incrementing columns in MySQL/PostgreSQL, `SEQUENCE`, or `AUTO`).
- **`@Column`:** Optional. Used to map a field to a specific database column name, or to define constraints like `nullable=false`, `length=255`, `unique=true`.
- **`@Transient`:** Marks a field that should NOT be persisted to the database. It exists only in the Java object state.

## 4. Explain the `@OneToOne`, `@OneToMany`, `@ManyToOne`, and `@ManyToMany` mappings.
**Answer:**
These annotations define relationships between database tables (and their corresponding Java objects).

- **`@OneToOne`:** A single record in Table A is associated with a single record in Table B (e.g., one User has one Profile). One side is usually the owning side containing the `@JoinColumn` (foreign key), and the other is the inverse side using `mappedBy`.
- **`@ManyToOne` / `@OneToMany`:** The most common relationship. Multiple records in Table A are associated with one record in Table B (e.g., many Employees belong to one Department).
    - `@ManyToOne` is placed on the side containing the foreign key (the "Many" side, e.g., Employee).
    - `@OneToMany` is placed on the "One" side (e.g., Department) and highly relies on the `mappedBy` attribute referencing the field in the owning entity.
- **`@ManyToMany`:** Multiple records in Table A are associated with multiple records in Table B (e.g., Students and Courses). This requires a separate "join table" in the database to link the primary keys of both tables. It uses `@JoinTable` to define the mapping.

## 5. What is the difference between `FetchType.LAZY` and `FetchType.EAGER`? Which is default?
**Answer:**
- **`FetchType.EAGER`:** When an entity is loaded from the database, the associated related entities (mappings) are immediately fetched along with it in a single query (using JOINs) or a secondary query.
- **`FetchType.LAZY`:** When an entity is loaded, the associated related entities are NOT fetched immediately. Instead, Hibernate puts a proxy object in place. The related entities are only fetched from the database when their getter methods are explicitly called for the first time.

**Defaults in JPA:**
- `@OneToOne` and `@ManyToOne`: **EAGER** by default (fetching a single entity is relatively cheap).
- `@OneToMany` and `@ManyToMany`: **LAZY** by default (fetching a large collection could cause severe performance issues and memory overhead).

**Best Practice:** Prefer `FetchType.LAZY` for almost all associations. Fetching too much unneeded data (eager fetching) is a leading cause of performance bottlenecks. If you need the data eagerly for a specific use case, use specific query strategies to fetch it dynamically.

## 6. What is the "N+1 Select Problem", and how do you solve it in Hibernate?
**Answer:**
The N+1 problem occurs when an application executes one primary query to fetch a list of *N* entities (e.g., 100 Authors), and then for each of those *N* entities, executes an additional *1* query to fetch their lazily-loaded associations (e.g., their Books). This results in 1 + 100 = 101 separate database queries, severely degrading performance.

**Solutions:**
1. **`JOIN FETCH` in JPQL / HQL:** Write a custom `@Query` using the `JOIN FETCH` clause. This tells Hibernate to eagerly load the associated collection in the same initial query using an SQL `INNER JOIN` or `LEFT OUTER JOIN`.
    ```java
    @Query("SELECT a FROM Author a JOIN FETCH a.books")
    List<Author> findAllWithBooks();
    ```
2. **EntityGraphs (JPA 2.1+):** An `@EntityGraph` defines a template of which attributes (and associations) should be fetched eagerly when a specific repository method is called, overriding the default lazy fetching dynamically without writing custom JPQL.
3. **Hibernate `@BatchSize`:** Setting `@BatchSize(size = 10)` on a collection entity instructs Hibernate to fetch the associations in batches (e.g., fetch books for 10 authors at once using `IN (...)` clauses), reducing 100 queries to 10.

## 7. How does Spring Data JPA simplify writing SQL queries? (Query Methods)
**Answer:**
Spring Data JPA offers multiple ways to query the database, significantly reducing the need to write raw SQL:

1. **Derived Query Methods (Method Name Parsing):** Spring Data automatically parses the method name in a repository interface and constructs the corresponding JPQL query.
    - Example: `List<User> findByEmailAndStatus(String email, String status);` -> Generates `SELECT u FROM User u WHERE u.email = ? AND u.status = ?`.
    - Supported keywords include `And`, `Or`, `Between`, `LessThan`, `GreaterThan`, `Like`, `OrderBy`, etc.
2. **`@Query` Annotation:** For complex queries that cannot be easily expressed via method names, you can use the `@Query` annotation.
    - **JPQL (Java Persistence Query Language):** Queries the JPA entities rather than raw tables. Very portable.
      `@Query("SELECT u FROM User u WHERE u.firstName = :name")`
    - **Native SQL Queries:** You can set `nativeQuery = true` to write raw SQL specific to your database (e.g., PostgreSQL JSON functions).
      `@Query(value = "SELECT * FROM users WHERE active = true", nativeQuery = true)`

## 8. Database configuration: How do you configure a connection to MySQL or PostgreSQL in Spring Boot?
**Answer:**
Connection pooling and data source configuration are heavily abstracted by Spring Boot auto-configuration.

**Steps:**
1. Ensure the appropriate JDBC driver dependency is in `pom.xml` (`mysql-connector-j` or `postgresql`).
2. Add the `spring-boot-starter-data-jpa` dependency.
3. Configure the `application.properties` or `application.yml` file:
    ```properties
    # For PostgreSQL
    spring.datasource.url=jdbc:postgresql://localhost:5432/mydb
    spring.datasource.username=postgres
    spring.datasource.password=secret

    # Hibernate configuration
    # Update schema automatically (useful for dev, use Flyway/Liquibase for prod)
    spring.jpa.hibernate.ddl-auto=update
    spring.jpa.show-sql=true
    spring.jpa.properties.hibernate.dialect=org.hibernate.dialect.PostgreSQLDialect
    ```
By default, Spring Boot 2.x and 3.x use **HikariCP** as the high-performance connection pool implementation.

## 9. How do you implement Transactions in Spring Boot?
**Answer:**
Spring provides comprehensive transaction management using the `@Transactional` annotation.

- **Placement:** You apply `@Transactional` at the class or method level, typically on the Service layer. (Applying it on the Controller layer is discouraged as it ties transactions to the web UI logic).
- **How it works (AOP proxies):** When a method annotated with `@Transactional` is called from *another* class, Spring's AOP proxy intercepts the call. It begins a database transaction before entering the method.
    - If the method completes successfully, the proxy commits the transaction.
    - If the method throws a `RuntimeException` (unchecked exception), the proxy automatically rolls back the transaction, ensuring data integrity across multiple repository operations.
- **Rollback Rules:** By default, it rolls back on `RuntimeException` and `Error`, but NOT on checked exceptions (like `IOException`). To enforce rollback on checked exceptions, use `@Transactional(rollbackFor = Exception.class)`.

## 10. Explain SQL vs NoSQL databases in a Java Microservices context.
**Answer:**
- **SQL (Relational Databases - MySQL, PostgreSQL):**
    - **Characteristics:** Structured data, strict schemas, tables and rows, ACID compliant (Atomicity, Consistency, Isolation, Durability), uses Joins.
    - **Use Case in Microservices:** Excellent for entities with strong relationships, complex querying, and transactions (e.g., User Management, Order Management, Financial configurations). Spring Data JPA handles this.
- **NoSQL (Non-Relational Databases - MongoDB, Cassandra, Redis):**
    - **Characteristics:** Unstructured or semi-structured data, dynamic schemas (document, key-value, column-family, graph), highly scalable horizontally, trades strong consistency for high availability (CAP theorem). No standard SQL joins.
    - **Use Case in Microservices:** Excellent for high-volume, rapidly changing data where scale is priority over rigid relations.
        - **MongoDB (Document Base):** E-commerce product catalogs, content management. Accessed via Spring Data MongoDB.
        - **Redis (Key-Value):** Caching, fast session store. Accessed via Spring Data Redis.

## 11. What is JPQL, and what is a Named Query?
**Answer:**
**JPQL (Java Persistence Query Language):**
An object-oriented query language used to perform database operations against JPA entities. Instead of querying database tables and columns directly (like SQL), you query Java classes and their properties.
- *SQL:* `SELECT * FROM tbl_users WHERE is_active = 1`
- *JPQL:* `SELECT u FROM User u WHERE u.active = true`

**Named Query (`@NamedQuery`):**
A statically defined JPQL query with an unchangeable query string. Instead of writing the JPQL directly in your repository methods, you define it on the Entity class itself using annotations.
- **Benefit:** They are parsed and validated by Hibernate exactly *once* when the Application Context starts up, offering a slight performance benefit and failing fast if there's a typo.
```java
@Entity
@NamedQuery(name = "User.findByStatus", query = "SELECT u FROM User u WHERE u.status = ?1")
public class User { ... }
```

## 12. Explain Transaction Isolation Levels in Spring/Databases.
**Answer:**
Transaction isolation dictates how changes made by one transaction become visible to other concurrent transactions. You can set it using `@Transactional(isolation = Isolation.READ_COMMITTED)`.

From lowest strictly to highest:
1. **`READ_UNCOMMITTED`:** The lowest level. A transaction can read uncommitted changes made by *other* transactions (Dirty Reads). In Spring, PostgreSQL ignores this and uses Read Committed anyway.
2. **`READ_COMMITTED` (Default for most DBs):** Guarantees that any data read is committed at the moment it is read. Prevents "Dirty Reads". However, if you read the same row twice in one transaction, another transaction might have updated it in between, yielding different results (Non-Repeatable Reads).
3. **`REPEATABLE_READ`:** If a transaction reads a row, no other transaction can modify that row until the first transaction finishes. Prevents "Non-Repeatable Reads". However, another transaction could insert *new* rows that match your query conditions, causing "Phantom Reads" on subsequent identical queries.
4. **`SERIALIZABLE`:** The highest and strictest level. Transactions occur in a completely isolated fashion, almost as if they were executed serially (one after another). Prevents all concurrency side effects (Dirty, Non-Repeatable, and Phantom reads) but drastically reduces application performance and concurrency.

## 13. What is Transactional Propagation in Spring? Explain `REQUIRED` vs `REQUIRES_NEW`.
**Answer:**
Propagation determines what happens when a `@Transactional` method is called from *another* `@Transactional` method.

1. **`Propagation.REQUIRED` (The Default):** If a transaction currently exists, join it. If none exists, create a new one.
   *Example:* Service A (transactional) calls Service B (transactional). Both execute in the *same* physical transaction. If B fails, the entire transaction (including A's work) rolls back.
2. **`Propagation.REQUIRES_NEW`:** Always suspend the current transaction (if one exists) and start a brand new, independent transaction.
   *Example:* Service A calls Service B (`REQUIRES_NEW`). Service B commits or rolls back completely independently of A. If B fails, A can catch the exception and still successfully commit its own work. Commonly used for Audit Logging (the log must save even if the main operation rolls back).
3. **`Propagation.SUPPORTS`:** If a transaction exists, join it. If not, execute non-transactionally.
4. **`Propagation.MANDATORY`:** Must run within an existing transaction. Throws an exception if none exists.
5. **`Propagation.NOT_SUPPORTED`:** Suspend the current transaction and execute non-transactionally.
6. **`Propagation.NEVER`:** Must execute non-transactionally. Throws an exception if a transaction exists.
7. **`Propagation.NESTED`:** Executes within a nested transaction using database savepoints. If the nested transaction fails, it rolls back to the savepoint without affecting the outer transaction's prior work.

## 14. What are Database Locking Mechanisms (Optimistic vs. Pessimistic Locking) in JPA?
**Answer:**
Locking prevents two concurrent transactions from updating the same row simultaneously and overriding each other's changes (the Lost Update anomaly).

**1. Optimistic Locking (Recommended for most web apps):**
- **Mental Model:** "I assume conflicts are rare, so I won't lock the database row. I'll just check for conflicts right before I save."
- **Implementation:** You add a `@Version` annotated field (usually an `Integer` or `Timestamp`) to your JPA Entity.
- **How it works:** When Transaction 1 reads a user (version=1), it brings it into memory. Transaction 2 reads the same user (version=1) and updates it (saving it to the DB with version=2). When Transaction 1 tries to save its changes later, Hibernate executes: `UPDATE users SET ... WHERE id = 1 AND version = 1`. Because the version in the DB is now 2, 0 rows are updated. Hibernate detects this and throws an `OptimisticLockException`, which you can catch to tell the user "This record was modified by someone else."

**2. Pessimistic Locking:**
- **Mental Model:** "I know there will be conflicts, so I will aggressively lock the actual database row until I'm completely finished."
- **Implementation:** Using `@Lock(LockModeType.PESSIMISTIC_WRITE)` on a Spring Data JPA Repository method.
- **How it works:** Hibernate translates this into a `SELECT ... FOR UPDATE` SQL statement. The database physically locks the row. If Transaction 2 tries to read or update that row, it will be physically blocked (frozen) until Transaction 1 commits or rolls back (or times out). It guarantees data integrity but heavily reduces concurrency and can cause deadlocks.
