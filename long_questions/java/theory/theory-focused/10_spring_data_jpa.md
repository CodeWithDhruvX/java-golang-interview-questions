# Spring Data JPA - Interview Answers

> 🎯 **Focus:** Database interaction is where most performance issues happen. These answers prove you know how to write efficient queries.

### 1. `JpaRepository` vs `CrudRepository`?
"`CrudRepository` is the base interface. It gives you standard methods like `save`, `findById`, `delete`.
`JpaRepository` extends it and adds JPA-specific features like `flush()` (to force SQL execution) and batch deletion.

I usually just extend `JpaRepository` because it covers everything, including Paging and Sorting features."

---

### 2. What is the N+1 Problem? How to solve it?
"This is a classic performance killer.
It happens when you fetch a list of entities (say, 10 `Orders`), and for *each* order, Hibernate fires a *separate* query to fetch its child (e.g., `Customer`).
So you get 1 query for the list + N queries for children. 100 orders = 101 queries.

To solve it, I use **Fetch Joins**.
In JPQL: `SELECT o FROM Order o JOIN FETCH o.customer`.
This forces Hibernate to load everything in a *single* SQL query using a JOIN."

---

### 3. Difference between `getOne` (getReference) vs `findById`?
"`findById` actually hits the database immediately. It returns an `Optional` containing the real entity.

`getOne` (now called `getReferenceById`) is lazy. It returns a **Proxy** object with just the ID. It doesn't hit the DB until you try to access a property other than the ID.
I use this when I just need to set a foreign key relationship (like saving an Order for a Customer) without wasting a query to fetch the full Customer object."

---

### 4. `@Transactional` propagation types?
"Propagation defines how transactions relate to each other.
**REQUIRED** (Default): If a transaction exists, join it. If not, create a new one. I use this 90% of the time.
**REQUIRES_NEW**: Suspends the current transaction and starts a brand new independent one. I use this for audit logging—even if the main business logic fails and rolls back, I still want to save the audit log."

---

### 5. Optimistic vs Pessimistic Locking?
"**Optimistic Locking** assumes conflicts are rare. It uses a `@Version` column. If two users try to update the same record, the second one fails with an Exception because the version doesn't match. It’s better for performance.

**Pessimistic Locking** locks the database row (`SELECT ... FOR UPDATE`). No one else can read/write that row until the transaction finishes. It’s safe but kills performance. I avoid it unless absolutely necessary (like handling money transfers)."

---

### 6. Cascade Types? (`ALL`, `PERSIST`, `REMOVE`)
"Cascading spreads operations from parent to child.
If I have an `Order` with `LineItems`:
`CascadeType.PERSIST`: If I save the Order, it automatically saves the new LineItems.
`CascadeType.REMOVE`: If I delete the Order, it deletes the LineItems.

I have to be careful with `CascadeType.ALL` (especially `REMOVE`) to avoid accidentally deleting data I didn't mean to."

---

### 7. Derived Query Methods?
"Spring Data is smart enough to generate SQL based on method names.
If I write `findByEmailAndActiveTrue(String email)`, it automatically creates:
`SELECT * FROM users WHERE email = ? AND active = true`.

It’s great for simple queries. But for complex ones with 3+ conditions, the method name gets ridiculously long, so I switch to `@Query` with JPQL."

---

### 8. `@Query` vs Native Query?
"**JPQL (@Query)** operates on *Entities*. It’s database-agnostic. `SELECT u FROM User u`.
**Native Query** operates on *Tables*. It’s raw SQL. `SELECT * FROM users_table`.

I prefer JPQL because it's portable. But if I need to use a database-specific feature (like Postgres JSONB operators or Window Functions), I have to use Native Queries."

---

### 9. How does `OpenSessionInView` work?
"It keeps the Hibernate Session open for the entire HTTP request, even during View rendering.
This allows you to lazy-load collections in the Controller or JSP/Thymeleaf.

While convenient, it’s considered an anti-pattern because it keeps the DB connection held for too long. In high-performance apps, I disable it (`spring.jpa.open-in-view=false`) and ensure I fetch all necessary data in the Service layer."

---

### 10. `First Level` vs `Second Level` Cache?
"**Level 1** is the Session cache. It’s on by default and lives only for one transaction. If I request ID=1 twice in the same transaction, Hibernate only queries the DB once.

**Level 2** is global (SessionFactory level). It spans across transactions. You have to enable it explicitly (using EhCache or Caffeine). I use L2 caching only for reference data that rarely changes (like Country lists) to avoid stale data issues."
