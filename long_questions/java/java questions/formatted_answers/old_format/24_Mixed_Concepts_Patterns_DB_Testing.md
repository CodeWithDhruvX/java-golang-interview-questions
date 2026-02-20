# 24. Mixed Concepts (Advanced SQL, Design Patterns, JDBC, Testing)

**Q: SQL: Handle Nulls and Unions**
> "**Handling Nulls**: Use `COALESCE(column_name, 'Default Value')` (standard SQL) or `IFNULL()` (MySQL). It checks if the value is null and creates a fallback.
>
> **Union vs Union All**:
> *   `UNION` combines results from two queries and **removes duplicates**. It's slower because it has to sort/check to find those duplicates.
> *   `UNION ALL` just blind appends everything. It's much faster. Use it if you know your data sets are unique or if you *want* duplicates."

**Indepth:**
> **Performance**: `UNION ALL` is significantly faster than `UNION` because it involves simple concatenation, whereas `UNION` requires sorting/hashing the entire result set to identify duplicates. Always prefer `UNION ALL` unless you specifically need deduplication.


---

**Q: SQL: Self Join and Ranking**
> "**Self Join**: Joining a table to itself.
> Useful for hierarchy (Managers/Employees) or finding pairs (Finding 2 employees with the same salary: `JOIN on e1.salary = e2.salary AND e1.id <> e2.id`).
>
> **Ranking**:
> Use Window Functions: `RANK()`, `DENSE_RANK()`, or `ROW_NUMBER()`.
> *   `ROW_NUMBER()`: Unique ID for every row (1, 2, 3, 4).
> *   `RANK()`: Skips numbers for ties (1, 1, 3).
> *   `DENSE_RANK()`: No skipping (1, 1, 2). Usually the best one for 'Nth highest' questions."

**Indepth:**
> **Optimization**: Window functions like `RANK()` are generally much more performant than using self-joins or correlated subqueries for ranking problems because the database engine can optimize the window operation in a single pass.


---

**Q: Factory vs Builder Pattern**
> "**Factory Pattern**: You have a method that creates objects for you. You don't say `new Dog()`, you say `AnimalFactory.create("Dog")`. It hides the complex logic of *instantiation*.
>
> **Builder Pattern**: Used for complex objects with many parameters (some optional). Instead of a constructor with 10 arguments (`new Pizza("Thin", null, "Pepperoni", true...)`), you chain methods:
> `new PizzaBuilder().crust("Thin").topping("Pepperoni").build()`. It's readable and flexible."

**Indepth:**
> **Fluency**: The Builder pattern often uses a static inner class `Builder` to ensure thread safety during construction and immutability of the final object. It solves the "Telescoping Constructor" anti-pattern.


---

**Q: Observer and Strategy Patterns**
> "**Observer**: Push notifications. An object (Subject) has a list of 'Subscribers'. When something changes, it loops through the list and calls `notify()`. Used in Event Listeners and Chat apps.
>
> **Strategy**: Pluggable behavior. You define a family of algorithms (SortFast, SortSlow, SortReverse) and make them interchangeable. You can pass the specific 'Strategy' into the method at runtime using an interface."

**Indepth:**
> **Decoupling**: Strategy allows you to change the guts of an object (how it does something) without changing the object itself. Observer allows you to react to changes without polling. Both heavily rely on Interfaces.


---

**Q: JDBC Steps (The 5 Steps)**
> "1.  **Load Driver**: `Class.forName(...)` (Though modern JDBC drivers auto-load).
> 2.  **Get Connection**: `DriverManager.getConnection(url, user, pass)`.
> 3.  **Create Statement**: `con.prepareStatement(sql)`.
> 4.  **Execute**: `stmt.executeQuery()` (for SELECT) or `stmt.executeUpdate()` (for INSERT/UPDATE).
> 5.  **Close**: Always close `ResultSet`, `Statement`, and `Connection` in a `try-with-resources` block to avoid leaks."

**Indepth:**
> **Modern JDBC**: In production, nobody writes raw JDBC like this anymore. We use Connection Pools (HikariCP) to reuse connections and Frameworks (Spring JDBC, Hibernate) to handle the boilerplate and resource cleanup.


---

**Q: PreparedStatement vs Statement**
> "Always use **PreparedStatement**.
> 1.  **Security**: It prevents SQL Injection by escaping inputs automatically.
> 2.  **Performance**: The database can compile and cache the query plan effectively because the structure (`SELECT * FROM users WHERE id = ?`) doesn't change, only the parameter changes."

**Indepth:**
> **Caching**: The DB parses the SQL template `SELECT * FROM users WHERE id = ?` once and caches the execution plan. If you run it 1000 times with different IDs, it reuses that plan. With `Statement`, every query string is different (`id=1`, `id=2`), forcing re-compilation.


---

**Q: ExecutorService & CompletableFuture**
> "**ExecutorService**: The manager for threads. You give it tasks (`Runnable` or `Callable`), and it assigns them to a pool of worker threads.
> *   `Callable`: Like Runnable, but it can return a result (`Future`) and throw Exceptions.
>
> **CompletableFuture**: The modern, non-blocking way to do async.
> It lets you chain tasks: 'Do task A, **then** do task B, **then** handle the error'. It keeps your main thread free."

**Indepth:**
> **Composition**: The real power of `CompletableFuture` is composition. `thenCompose()` allows you to chain dependent async operations (result of A is input to B), while `thenCombine()` runs independent operations in parallel and merges results.


---

**Q: Unit Testing (JUnit Basics)**
> "A Unit Test checks a small, isolated piece of logic.
>
> **Key Annotations**:
> *   `@Test`: Marks a method as a test.
> *   `@BeforeEach`: Runs before *every* test (good for resetting data).
> *   `@BeforeAll`: Runs once before *anything* (good for expensive setup like DB connections).
>
> Use `Assertions` to verify: `assertEquals(expected, actual)`.
> Mock dependencies (using Mockito) so you are testing *only* your class, not the database or network."

**Indepth:**
> **Isolation**: A unit test should never touch the file system, network, or database. If it does, it's an Integration Test. Unit tests must be fast (milliseconds) and deterministic. Use Mocking to fake external dependencies.

