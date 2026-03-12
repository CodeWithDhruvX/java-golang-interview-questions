# Spring Data JPA Interview Questions

## Repositories & Querying

### 1. What is Spring Data JPA?
Spring Data JPA is part of the Spring Data project that makes it easier to implement JPA-based repositories. It reduces the amount of boilerplate code required to access data layers. You just define interfaces (like `JpaRepository`), and Spring Data automatically provides the implementation at runtime.

**Explanation:** Spring Data JPA acts as an abstraction layer on top of JPA providers like Hibernate. Instead of writing DAO implementations with EntityManager and boilerplate CRUD operations, you simply extend repository interfaces and Spring generates the implementation at runtime using proxy patterns.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Data JPA?
**Your Response:** Spring Data JPA is a framework that simplifies database access by eliminating the need to write boilerplate DAO code. Instead of manually implementing repository classes with EntityManager and SQL queries, I just extend interfaces like JpaRepository and Spring automatically provides the implementation. It reduces development time significantly and follows the convention over configuration principle.

### 2. What is `JpaRepository`?
It is a JPA-specific extension of `Repository`. It contains the full API of `CrudRepository` and `PagingAndSortingRepository`. It contains API for basic CRUD operations and also API for pagination and sorting.

**Explanation:** JpaRepository inherits from PagingAndSortingRepository, which itself inherits from CrudRepository. This hierarchy provides a complete set of database operations - from basic CRUD to advanced features like batch operations, flushing, and entity-specific queries.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is JpaRepository?
**Your Response:** JpaRepository is a Spring Data interface that provides comprehensive database operations. It extends PagingAndSortingRepository and CrudRepository, giving us everything from basic create, read, update, delete operations to pagination, sorting, and batch operations. By extending this interface, I get all these methods without writing any implementation code.

### 3. How do you define custom queries in Spring Data JPA?
1.  **Derived Query Methods:** Naming the method correctly allows Spring to derive the query.
    - `findByLastname(String lastname)`
    - `inputUserByEmailAndActiveTrue(String email)`
2.  **@Query Annotation:** Writing JPQL or Native SQL directly.
    ```java
    @Query("SELECT u FROM User u WHERE u.status = ?1")
    List<User> findActiveUsers(Integer status);
    ```

**Explanation:** Derived queries use method naming conventions where Spring automatically generates the JPQL query based on method names. The @Query annotation gives more control for complex queries or when you need to optimize performance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define custom queries in Spring Data JPA?
**Your Response:** There are two main approaches. First, I can use derived query methods where I just name my method following Spring's conventions like findByUsername or findByEmailAndActive, and Spring automatically generates the query. Second, for more complex queries, I can use the @Query annotation to write custom JPQL or native SQL directly. This gives me full control when the naming conventions aren't sufficient.

### 4. What is the difference between `getOne()` (now `getReferenceById()`) and `findById()`?
- **`findById(id)`:** Actually hits the database and returns an `Optional<T>`. It performs an eager fetch.
- **`getReferenceById(id)` (formerly `getOne`):** Returns a proxy (lazy fetch). It assumes the entity exists and doesn't hit the DB until you access a property of the entity. Useful for setting foreign keys without a DB roundtrip.

**Explanation:** The key difference is when the database query happens. findById executes immediately, while getReferenceById uses a proxy pattern to defer the query until you actually access the entity's properties.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between getOne and findById?
**Your Response:** The main difference is when the database access happens. findById immediately queries the database and returns an Optional with the entity or empty. getReferenceById, which replaced the deprecated getOne method, returns a proxy object and doesn't hit the database until I actually access the entity's properties. This is useful when I just need to set up a relationship, like assigning a foreign key, without the overhead of loading the full entity.

### 5. What is the N+1 problem and how do you solve it in Hibernate/JPA?
The N+1 problem occurs when you fetch a list of N entities, and for each entity, you perform an additional query to fetch related data (e.g., retrieving a list of Users, then querying the Address for each User separately).
- **Solution:** Use **Join Fetch**.
    - in JPQL: `SELECT u FROM User u JOIN FETCH u.address`
    - Use `@EntityGraph` in Spring Data to specify fetch plans.

**Explanation:** This performance issue happens when lazy loading causes multiple additional queries instead of one efficient query with joins. Entity graphs provide a declarative way to define fetch plans at the repository level.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the N+1 problem and how do you solve it?
**Your Response:** The N+1 problem is a common performance issue where I execute one query to fetch a list of entities, then N additional queries to fetch related data for each entity. For example, fetching 100 users and then separately querying each user's address results in 101 database queries instead of one. I solve this using JOIN FETCH in JPQL or @EntityGraph annotations to load related data in a single query, significantly improving performance.

## Entities & Mapping

### 6. Explain Entity Lifecycle States.
1.  **Transient:** The object is created (`new User()`) but not yet associated with an EntityManager. It has no ID.
2.  **Persistent:** The object is associated with an EntityManager and has an ID (e.g., after `save()`). Changes to it are tracked and synced to DB.
3.  **Detached:** The object has an ID but the EntityManager session is closed. Changes are not synced.
4.  **Removed:** The object is scheduled for deletion.

**Explanation:** These states represent how JPA manages entity objects within the persistence context. The EntityManager tracks state transitions and ensures data consistency.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain the JPA entity lifecycle states?
**Your Response:** JPA entities go through four states. Transient is when I create a new object but haven't saved it yet - it has no database identity. Persistent is when the entity is managed by the EntityManager and any changes are automatically tracked. Detached happens when the persistence context ends but the object still has an ID. Removed means the entity is marked for deletion. Understanding these states helps me manage entity lifecycle effectively.

### 7. What is `@Transactional`?
It defines the scope of a single database transaction.
- If a method is annotated with `@Transactional`, Spring ensures a transaction is started before the method begins and committed after it ends.
- If an unchecked exception (Runtime Exception) is thrown, the transaction is automatically rolled back. Checked exceptions do not trigger rollback by default (unless configured).

**Explanation:** Spring's transaction management uses AOP proxies to wrap methods with transactional behavior, ensuring ACID properties and handling rollbacks automatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is @Transactional?
**Your Response:** @Transactional is Spring's annotation for managing database transactions. When I annotate a method, Spring automatically handles the transaction lifecycle - it starts a transaction before the method executes and commits it after completion. If any runtime exception occurs, Spring automatically rolls back the transaction. I can configure it for different isolation levels, propagation behaviors, and specify which exceptions should trigger rollback.

### 8. What is Optimistic vs. Pessimistic Locking?
- **Optimistic Locking:** Assumes conflict is rare. Uses a version column (`@Version`). If the version in the DB is different from the one being updated, an `OptimisticLockException` is thrown.
- **Pessimistic Locking:** Locks the record in the database (`SELECT ... FOR UPDATE`). No other transaction can modify it until the lock is released.

**Explanation:** Optimistic locking is better for high-concurrency read scenarios, while pessimistic locking prevents conflicts but can reduce concurrency and performance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between optimistic and pessimistic locking?
**Your Response:** Optimistic locking assumes that database conflicts are rare, so it doesn't lock records upfront. Instead, it uses a version field to detect if someone else modified the data since I read it. If the version changed, I get an exception. Pessimistic locking takes the opposite approach - it locks the record immediately when I read it, preventing others from modifying it until I'm done. I use optimistic locking for most applications as it provides better concurrency, and pessimistic locking only when conflicts are very likely and costly.

### 9. What is Cascading?
It allows operations performed on a parent entity to be propagated to child entities.
- `CascadeType.ALL`: Propagates all operations.
- `CascadeType.PERSIST`: Saving the parent saves the child.
- `CascadeType.REMOVE`: Deleting the parent deletes the child.

**Explanation:** Cascading defines how entity operations flow through object relationships, reducing the need to manually manage related entity operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is cascading in JPA?
**Your Response:** Cascading allows me to automatically propagate operations from parent entities to their children. For example, if I have an Order entity with OrderItem entities, I can configure cascading so that when I save the Order, all its OrderItems are automatically saved too. I can specify which operations cascade - persist, merge, remove, or all operations. This eliminates the need to manually manage related entities and ensures data consistency.

### 10. What is Lazy Loading vs. Eager Loading?
- **Lazy:** Related data is fetched only when accessed. (Default for `@OneToMany`, `@ManyToMany`).
- **Eager:** Related data is fetched immediately with the parent. (Default for `@ManyToOne`, `@OneToOne`).

**Explanation:** Loading strategy impacts performance significantly. Lazy loading prevents unnecessary data retrieval but can cause N+1 problems, while eager loading ensures all data is available but may fetch more than needed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between lazy and eager loading?
**Your Response:** Lazy loading means related data is only fetched from the database when I actually access it, while eager loading fetches everything immediately when I load the main entity. By default, JPA uses lazy loading for collections like @OneToMany and eager loading for single relationships like @ManyToOne. I choose lazy loading most of the time for better performance, but I need to be careful about the N+1 problem and sometimes use eager loading or JOIN FETCH when I know I'll need the related data right away.

## Advanced JPA Topics

### 11. What are JPA Projections?
Projections allow fetching only a subset of attributes (columns) instead of the whole entity.
- **Interface-based:** Define an interface with getter methods matching property names.
    ```java
    public interface UserNamesOnly {
        String getFirstname();
        String getLastname();
    }
    ```
- **Class-based (DTO):** Use constructor expression in JPQL (`SELECT new com.example.UserDTO(...)`).

**Explanation:** Projections improve performance by reducing data transfer and memory usage, especially for read-heavy operations or when entities have many fields.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are JPA projections?
**Your Response:** JPA projections let me fetch only specific fields from an entity instead of loading the entire object. This is really useful for performance optimization, especially when entities have many fields but I only need a few. I can create interface-based projections where I define an interface with just the getter methods I need, or class-based DTOs using constructor expressions. This reduces database load and network traffic, making my applications more efficient.

### 12. Explain Hibernate Caching (L1 & L2).
- **Level 1 (Session) Cache:** Enabled by default. Caches objects within the current transaction/session.
- **Level 2 (SessionFactory) Cache:** Optional, global cache across transactions. Requires a provider (EhCache, Redis). Good for read-heavy data that changes rarely (e.g., Country list).

**Explanation:** L1 cache is built-in and provides automatic caching within a session, while L2 cache requires configuration and provides cross-session caching for frequently accessed data.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain Hibernate caching?
**Your Response:** Hibernate has two levels of caching. Level 1 cache is enabled by default and works within a single database session - it prevents multiple queries for the same entity within the same transaction. Level 2 cache is optional and provides application-wide caching across different sessions. I configure L2 cache for data that doesn't change often but is read frequently, like reference data or lookup tables. This significantly reduces database load and improves application performance.

### 13. What is Dirty Checking?
Hibernate feature where it automatically detects changes made to a persistent object during a transaction and synchronizes them with the database upon commit. You don't need to explicitly call `save()` for managed entities.

**Explanation:** Dirty checking uses snapshot comparison to detect changes, automatically generating UPDATE statements for modified entities without manual intervention.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is dirty checking in Hibernate?
**Your Response:** Dirty checking is Hibernate's automatic mechanism for tracking changes to entities. When I load an entity, Hibernate takes a snapshot of its state. At the end of the transaction, it compares the current state with the snapshot and automatically generates UPDATE statements for any fields that changed. This means I don't need to manually call save or update methods - Hibernate handles it automatically as long as the entity is in a persistent state within the transaction.

### 14. How do you implement Auditing in Spring Data?
- Use `@EnableJpaAuditing` configuration.
- Annotate entity with `@EntityListeners(AuditingEntityListener.class)`.
- Use `@CreatedDate`, `@LastModifiedDate`, `@CreatedBy`, `@LastModifiedBy` on fields.

**Explanation:** Spring Data JPA auditing automatically tracks entity metadata like creation/modification timestamps and users, reducing boilerplate code and ensuring consistent audit trails.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement auditing in Spring Data?
**Your Response:** I implement auditing using Spring Data JPA's built-in features. First, I enable JPA auditing with @EnableJpaAuditing in my configuration. Then I add @EntityListeners to my entities and annotate fields with @CreatedDate, @LastModifiedDate for timestamps, and @CreatedBy, @LastModifiedBy for user tracking. Spring automatically populates these fields when entities are created or modified, giving me a complete audit trail without writing any additional code.
