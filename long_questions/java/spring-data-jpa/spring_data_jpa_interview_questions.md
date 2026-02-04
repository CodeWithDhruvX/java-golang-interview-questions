# Spring Data JPA Interview Questions

## Repositories & Querying

### 1. What is Spring Data JPA?
Spring Data JPA is part of the Spring Data project that makes it easier to implement JPA-based repositories. It reduces the amount of boilerplate code required to access data layers. You just define interfaces (like `JpaRepository`), and Spring Data automatically provides the implementation at runtime.

### 2. What is `JpaRepository`?
It is a JPA-specific extension of `Repository`. It contains the full API of `CrudRepository` and `PagingAndSortingRepository`. It contains API for basic CRUD operations and also API for pagination and sorting.

### 3. How do you define custom queries in Spring Data JPA?
1.  **Derived Query Methods:** Naming the method correctly allows Spring to derive the query.
    - `findByLastname(String lastname)`
    - `inputUserByEmailAndActiveTrue(String email)`
2.  **@Query Annotation:** Writing JPQL or Native SQL directly.
    ```java
    @Query("SELECT u FROM User u WHERE u.status = ?1")
    List<User> findActiveUsers(Integer status);
    ```

### 4. What is the difference between `getOne()` (now `getReferenceById()`) and `findById()`?
- **`findById(id)`:** Actually hits the database and returns an `Optional<T>`. It performs an eager fetch.
- **`getReferenceById(id)` (formerly `getOne`):** Returns a proxy (lazy fetch). It assumes the entity exists and doesn't hit the DB until you access a property of the entity. Useful for setting foreign keys without a DB roundtrip.

### 5. What is the N+1 problem and how do you solve it in Hibernate/JPA?
The N+1 problem occurs when you fetch a list of N entities, and for each entity, you perform an additional query to fetch related data (e.g., retrieving a list of Users, then querying the Address for each User separately).
- **Solution:** Use **Join Fetch**.
    - in JPQL: `SELECT u FROM User u JOIN FETCH u.address`
    - Use `@EntityGraph` in Spring Data to specify fetch plans.

## Entities & Mapping

### 6. Explain Entity Lifecycle States.
1.  **Transient:** The object is created (`new User()`) but not yet associated with an EntityManager. It has no ID.
2.  **Persistent:** The object is associated with an EntityManager and has an ID (e.g., after `save()`). Changes to it are tracked and synced to DB.
3.  **Detached:** The object has an ID but the EntityManager session is closed. Changes are not synced.
4.  **Removed:** The object is scheduled for deletion.

### 7. What is `@Transactional`?
It defines the scope of a single database transaction.
- If a method is annotated with `@Transactional`, Spring ensures a transaction is started before the method begins and committed after it ends.
- If an unchecked exception (Runtime Exception) is thrown, the transaction is automatically rolled back. Checked exceptions do not trigger rollback by default (unless configured).

### 8. What is Optimistic vs. Pessimistic Locking?
- **Optimistic Locking:** Assumes conflict is rare. Uses a version column (`@Version`). If the version in the DB is different from the one being updated, an `OptimisticLockException` is thrown.
- **Pessimistic Locking:** Locks the record in the database (`SELECT ... FOR UPDATE`). No other transaction can modify it until the lock is released.

### 9. What is Cascading?
It allows operations performed on a parent entity to be propagated to child entities.
- `CascadeType.ALL`: Propagates all operations.
- `CascadeType.PERSIST`: Saving the parent saves the child.
- `CascadeType.REMOVE`: Deleting the parent deletes the child.

### 10. What is Lazy Loading vs. Eager Loading?
- **Lazy:** Related data is fetched only when accessed. (Default for `@OneToMany`, `@ManyToMany`).
- **Eager:** Related data is fetched immediately with the parent. (Default for `@ManyToOne`, `@OneToOne`).

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

### 12. Explain Hibernate Caching (L1 & L2).
- **Level 1 (Session) Cache:** Enabled by default. Caches objects within the current transaction/session.
- **Level 2 (SessionFactory) Cache:** Optional, global cache across transactions. Requires a provider (EhCache, Redis). Good for read-heavy data that changes rarely (e.g., Country list).

### 13. What is Dirty Checking?
Hibernate feature where it automatically detects changes made to a persistent object during a transaction and synchronizes them with the database upon commit. You don't need to explicitly call `save()` for managed entities.

### 14. How do you implement Auditing in Spring Data?
- Use `@EnableJpaAuditing` configuration.
- Annotate entity with `@EntityListeners(AuditingEntityListener.class)`.
- Use `@CreatedDate`, `@LastModifiedDate`, `@CreatedBy`, `@LastModifiedBy` on fields.
