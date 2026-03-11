# Java Intermediate — Databases, JDBC & JPA Deep Dive

> **Topics:** Raw JDBC (`Connection`, `PreparedStatement`, `ResultSet`), SQL injection, JPA entity lifecycle, `EntityManager`, N+1 problem, `JOIN FETCH`, `EntityGraph`, transaction propagation, isolation levels, Spring Data JPA, `@Query`, optimistic/pessimistic locking, Hibernate caching

---

## 📋 Reading Progress

- [ ] **Section 1:** Raw JDBC — The Foundation (Q1–Q15)
- [ ] **Section 2:** JPA Entity Lifecycle & EntityManager (Q16–Q28)
- [ ] **Section 3:** Relationships & N+1 Problem (Q29–Q38)
- [ ] **Section 4:** Transactions — Propagation & Isolation (Q39–Q47)
- [ ] **Section 5:** Spring Data JPA — Queries & Locking (Q48–Q55)

> 🔖 **Last read:** <!-- -->

---

## Section 1: Raw JDBC — The Foundation (Q1–Q15)

### 1. JDBC Connection — Basic Flow
**Q: What is the correct order of operations?**
```java
import java.sql.*;
public class Main {
    public static void main(String[] args) throws SQLException {
        // 1. Get connection (from DriverManager or DataSource pool)
        Connection conn = DriverManager.getConnection(
            "jdbc:mysql://localhost:3306/mydb", "root", "password");

        // 2. Create statement
        PreparedStatement ps = conn.prepareStatement("SELECT id, name FROM users WHERE id = ?");
        ps.setLong(1, 42L);

        // 3. Execute query
        ResultSet rs = ps.executeQuery();

        // 4. Iterate results
        while (rs.next()) {
            System.out.println(rs.getLong("id") + ": " + rs.getString("name"));
        }

        // 5. Close in reverse order
        rs.close(); ps.close(); conn.close();
    }
}
```
**A:** Order: Connection → Statement → ResultSet → iterate → close in reverse. In modern code, use try-with-resources. Never use direct `DriverManager` in production — use a `DataSource` with connection pooling (HikariCP).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the correct order of operations here?
**Your Response:** The correct order is: first establish a connection, then create a prepared statement, execute the query to get a result set, iterate through the results, and finally close the resources in reverse order - result set first, then statement, then connection. In modern applications, we'd use try-with-resources to handle closing automatically, and we'd never use DriverManager directly in production - we'd use a connection pool like HikariCP for better performance.

---

### 2. SQL Injection — Statement vs PreparedStatement
**Q: Which is vulnerable?**
```java
import java.sql.*;
public class Main {
    // VULNERABLE: user input concatenated directly
    static void unsafe(Connection conn, String userId) throws SQLException {
        Statement st = conn.createStatement();
        ResultSet rs = st.executeQuery("SELECT * FROM users WHERE id = " + userId);
        // userId = "1 OR 1=1" → dumps entire table!
    }

    // SAFE: parameterized query
    static void safe(Connection conn, String userId) throws SQLException {
        PreparedStatement ps = conn.prepareStatement("SELECT * FROM users WHERE id = ?");
        ps.setString(1, userId); // input is treated as data, never SQL
        ResultSet rs = ps.executeQuery();
    }
}
```
**A:** `unsafe()` is vulnerable to SQL injection. `safe()` uses a parameterized query — the `?` placeholder is sent to the DB server separately from the data. **Always use `PreparedStatement`.**

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which method is vulnerable to SQL injection and why?
**Your Response:** The `unsafe()` method is definitely vulnerable because it concatenates user input directly into the SQL string. If someone passes `"1 OR 1=1"` as the userId, the query becomes `SELECT * FROM users WHERE id = 1 OR 1=1` which returns all users. The `safe()` method uses a parameterized query where the input is sent to the database separately from the SQL command, so it's treated as data, not as executable SQL. This is why we should always use PreparedStatement with parameter markers.

---

### 3. ResultSet — Column Access Patterns
**Q: What is the bug?**
```java
import java.sql.*;
public class Main {
    public static void main(String[] args) throws SQLException {
        // Assume: SELECT name, age FROM users
        ResultSet rs = /* ... */null;
        while (rs.next()) {
            // By column index (1-based, not 0-based!) ← common mistake
            String name = rs.getString(1); // correct
            int age     = rs.getInt(2);   // correct

            // By column name (preferred — immune to column order changes)
            String name2 = rs.getString("name");
            int age2     = rs.getInt("age");
        }
    }
}
```
**A:** JDBC column indexes are **1-based** (not 0-based like arrays). Index-based access is fragile — prefer column name access for maintainability.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the bug in this ResultSet code?
**Your Response:** The main issue here is that JDBC column indexes are 1-based, not 0-based like Java arrays. So `rs.getString(1)` gets the first column, not the second. But more importantly, using column indexes is fragile because if someone changes the SQL query column order, the code breaks. It's much better to use column names like `rs.getString("name")` - this is more readable and immune to column order changes in the SQL query.

---

### 4. Null Handling in ResultSet
**Q: What is the gotcha with rs.getInt() for a NULL column?**
```java
import java.sql.*;
public class Main {
    public static void main(String[] args) throws SQLException {
        ResultSet rs = /* ... */null;
        int age = rs.getInt("age");            // returns 0 if SQL NULL!
        boolean wasNull = rs.wasNull();        // check AFTER the getter call

        Integer ageBoxed = rs.getObject("age", Integer.class); // null-safe alternative (JDBC 4.1+)
        System.out.println("age=" + ageBoxed + " wasNull=" + wasNull);
    }
}
```
**A:** `rs.getInt()` returns `0` for SQL `NULL` — no exception. Call `rs.wasNull()` immediately after to detect `NULL`. Prefer `rs.getObject("col", Integer.class)` which returns a proper Java `null`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the gotcha with rs.getInt() for a NULL column?
**Your Response:** This is a classic JDBC gotcha! When you call rs.getInt() on a column that contains SQL NULL, it doesn't throw an exception - it returns 0. This can be dangerous because 0 might be a valid value in your domain. You need to call rs.wasNull() immediately after the getter to check if the original value was NULL. A better approach is to use rs.getObject() with the Integer class, which returns a proper Java null that you can handle explicitly.

---

### 5. Batch Updates — Performance
**Q: Why is batch insert faster?**
```java
import java.sql.*;
public class Main {
    static void batchInsert(Connection conn, List<String> names) throws SQLException {
        conn.setAutoCommit(false); // disable auto-commit for batch
        PreparedStatement ps = conn.prepareStatement("INSERT INTO names(name) VALUES(?)");

        for (String name : names) {
            ps.setString(1, name);
            ps.addBatch(); // buffer the insert
        }

        int[] counts = ps.executeBatch(); // send all at once
        conn.commit();
        System.out.println("inserted: " + counts.length + " rows");
    }
}
```
**A:** Batch insert sends multiple INSERTs in one network round-trip. Without batch, each insert = one round-trip. For 10,000 rows, batch can be 10–100x faster depending on network latency.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why is batch insert so much faster than individual inserts?
**Your Response:** Batch operations are dramatically faster because they reduce network round-trips. Without batching, each INSERT statement requires a separate network call to the database. With batching, we buffer multiple inserts and send them all in one network round-trip. For 10,000 rows, that's 10,000 round-trips vs just 1. The performance gain is especially noticeable with high network latency. We also disable auto-commit during batch operations to avoid committing after each individual insert, which adds overhead.

---

### 6. Transactions — commit() and rollback()
**Q: What is the output if exception occurs mid-transaction?**
```java
import java.sql.*;
public class Main {
    static void transfer(Connection conn, long fromId, long toId, double amt) throws SQLException {
        conn.setAutoCommit(false);
        try {
            PreparedStatement debit = conn.prepareStatement("UPDATE accounts SET balance = balance - ? WHERE id = ?");
            debit.setDouble(1, amt); debit.setLong(2, fromId); debit.executeUpdate();

            // Simulate error after debit but before credit
            if (true) throw new SQLException("bank error");

            PreparedStatement credit = conn.prepareStatement("UPDATE accounts SET balance = balance + ? WHERE id = ?");
            credit.setDouble(1, amt); credit.setLong(2, toId); credit.executeUpdate();

            conn.commit();
        } catch (SQLException e) {
            conn.rollback(); // undo debit
            System.out.println("rolled back: " + e.getMessage());
        }
    }
}
```
**A:** `rolled back: bank error`. The debit is undone. Without explicit `rollback()`, the partial transaction is committed when the connection is returned to the pool (`setAutoCommit(true)`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens when an exception occurs in this transaction?
**Your Response:** When the exception occurs, the catch block calls rollback(), which undoes the debit that was already executed. This ensures atomicity - either both the debit and credit happen, or neither happens. Without explicit rollback(), when the connection returns to the pool, the pool might call setAutoCommit(true), which could commit the partial transaction. This is why it's critical to always handle rollback in catch blocks when working with manual transaction management.

---

### 7. Connection Pool — HikariCP
**Q: Why use a connection pool?**
```java
import com.zaxxer.hikari.*;
public class Main {
    public static void main(String[] args) {
        HikariConfig config = new HikariConfig();
        config.setJdbcUrl("jdbc:mysql://localhost:3306/mydb");
        config.setUsername("root"); config.setPassword("pass");
        config.setMaximumPoolSize(10);    // max 10 simultaneous connections
        config.setMinimumIdle(2);         // keep 2 connections warm
        config.setConnectionTimeout(3000); // 3s to get a connection

        HikariDataSource ds = new HikariDataSource(config);
        // Spring Boot auto-configures HikariCP — it's the default

        System.out.println("pool ready");
    }
}
```
**A:** Creating DB connections is expensive (~100ms). A pool reuses existing connections. HikariCP is the fastest Java connection pool — Spring Boot auto-configures it. Never call `DriverManager.getConnection()` in production.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why is connection pooling so important in production applications?
**Your Response:** Database connections are expensive to create - they involve network handshakes, authentication, and server-side resource allocation, often taking 100ms or more. Connection pooling maintains a set of established connections that can be reused, eliminating this overhead for each request. HikariCP is the fastest Java connection pool and is Spring Boot's default. In production, you should never call DriverManager.getConnection() directly - always use a connection pool. This dramatically improves application throughput and reduces database load.

---

### 8. DataSource in Spring Boot
**Q: What does Spring Boot auto-configure?**
```yaml
# application.yml
spring:
  datasource:
    url: jdbc:postgresql://localhost:5432/mydb
    username: myapp
    password: secret
    hikari:
      maximum-pool-size: 20
      idle-timeout: 30000
```
```java
@Repository
class UserDao {
    @Autowired DataSource ds; // injected automatically

    public User find(Long id) throws SQLException {
        try (Connection c = ds.getConnection();
             PreparedStatement ps = c.prepareStatement("SELECT * FROM users WHERE id=?")) {
            ps.setLong(1, id);
            ResultSet rs = ps.executeQuery();
            // ...
        }
    }
}
```
**A:** Spring Boot reads `spring.datasource.*` and auto-creates a `HikariDataSource` bean. JdbcTemplate, JPA, and Spring Data all use this `DataSource` automatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot handle database configuration?
**Your Response:** Spring Boot provides excellent auto-configuration for databases. When you include a JDBC driver on the classpath and configure spring.datasource properties, Spring Boot automatically creates a HikariDataSource connection pool. This DataSource is then injected and used by all Spring data access components like JdbcTemplate, JPA EntityManagerFactory, and Spring Data repositories. This convention-over-configuration approach eliminates most boilerplate - you just need to provide the database URL, credentials, and optionally pool settings in your application.properties or yml file.

---

### 9. JdbcTemplate — Simplified JDBC
**Q: What is the output?**
```java
import org.springframework.jdbc.core.*;
import org.springframework.stereotype.*;

@Repository
class UserDao {
    @Autowired JdbcTemplate jdbc;

    public List<String> findAllNames() {
        return jdbc.queryForList("SELECT name FROM users", String.class);
    }

    public User findById(Long id) {
        return jdbc.queryForObject(
            "SELECT id, name FROM users WHERE id = ?",
            (rs, rowNum) -> new User(rs.getLong("id"), rs.getString("name")),
            id
        );
    }

    public int insert(String name) {
        return jdbc.update("INSERT INTO users(name) VALUES(?)", name);
    }
}
```
**A:** `JdbcTemplate` eliminates boilerplate: no manual `Connection`, `PreparedStatement`, `ResultSet` lifecycle management. Exceptions are translated to Spring's `DataAccessException` hierarchy.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the benefits of using JdbcTemplate?
**Your Response:** JdbcTemplate eliminates all the boilerplate code of raw JDBC - no more manual Connection, PreparedStatement, and ResultSet lifecycle management. It handles resource cleanup automatically with try-with-resources patterns and translates SQLExceptions into Spring's DataAccessException hierarchy, which is unchecked and provides better error information. This makes database code much cleaner and less error-prone while still giving you full control over the SQL being executed.

---

### 10. RowMapper vs ResultSetExtractor
**Q: When do you use each?**
```java
// RowMapper: maps one row to one object (most common)
RowMapper<User> mapper = (rs, rowNum) -> new User(rs.getLong("id"), rs.getString("name"));
List<User> users = jdbc.query("SELECT * FROM users", mapper);

// ResultSetExtractor: processes entire ResultSet (e.g., multi-row aggregation)
Map<Long, List<String>> userOrders = jdbc.query(
    "SELECT u.id, o.product FROM users u JOIN orders o ON u.id = o.user_id",
    rs -> {
        Map<Long, List<String>> map = new LinkedHashMap<>();
        while (rs.next()) {
            map.computeIfAbsent(rs.getLong("id"), k -> new ArrayList<>()).add(rs.getString("product"));
        }
        return map;
    }
);
```
**A:** `RowMapper` maps row-by-row (Spring calls it per row). `ResultSetExtractor` receives the whole `ResultSet` — use for complex aggregations like one-to-many result set processing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use RowMapper vs ResultSetExtractor?
**Your Response:** RowMapper is the common choice - it maps one row to one object, and Spring calls it for each row in the result set. Use RowMapper for simple object mapping. ResultSetExtractor is more powerful - it receives the entire ResultSet at once, allowing you to implement complex aggregations or build nested object structures. It's perfect for one-to-many relationships where you need to process multiple rows to build a single object hierarchy. The trade-off is that ResultSetExtractor requires more manual iteration code.

---

### 11. @Transactional with JdbcTemplate
**Q: Are both inserts in the same transaction?**
```java
@Service
class AccountService {
    @Autowired JdbcTemplate jdbc;

    @Transactional
    public void transfer(long from, long to, int amount) {
        jdbc.update("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, from);
        jdbc.update("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, to);
        // both updates share the same connection/transaction
    }
}
```
**A:** Yes — Spring's `@Transactional` binds the connection to the current thread. `JdbcTemplate` uses `DataSourceUtils.getConnection()` which retrieves the bound connection. Both updates commit or rollback together.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring ensure both JdbcTemplate operations are in the same transaction?
**Your Response:** Spring's @Transactional works by binding a single database connection to the current thread for the duration of the transaction. When JdbcTemplate needs a connection, it calls DataSourceUtils.getConnection(), which first checks if there's already a connection bound to the current thread. If there is, it reuses that connection instead of getting a new one from the pool. This ensures that all database operations within the transactional method use the same connection, so they either all commit together or all rollback together.

---

### 12. Named Parameters — NamedParameterJdbcTemplate
**Q: What is the benefit?**
```java
@Repository
class SearchDao {
    @Autowired NamedParameterJdbcTemplate jdbc;

    public List<User> search(String name, int minAge) {
        Map<String, Object> params = Map.of("name", "%" + name + "%", "minAge", minAge);
        return jdbc.query(
            "SELECT * FROM users WHERE name LIKE :name AND age >= :minAge",
            params,
            (rs, n) -> new User(rs.getLong("id"), rs.getString("name"))
        );
    }
}
```
**A:** Named parameters (`:name`) are more readable than positional (`?`) especially with many params. Prevents accidental parameter swaps. Backed by `PreparedStatement` — still safe from SQL injection.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the advantage of using NamedParameterJdbcTemplate?
**Your Response:** NamedParameterJdbcTemplate makes SQL much more readable and maintainable by using named parameters like :name instead of question marks. This is especially valuable when you have many parameters or complex queries, as it prevents accidental parameter swaps that can happen with positional parameters. Under the hood, it still uses PreparedStatement for SQL injection protection, so you get both readability and security. It's particularly useful for dynamic query building where you might conditionally include different parameters.

---

### 13. SimpleJdbcInsert — Auto-Generated Keys
**Q: How do you get the auto-generated ID after insert?**
```java
@Repository
class ProductDao {
    private final SimpleJdbcInsert insert;

    ProductDao(DataSource ds) {
        this.insert = new SimpleJdbcInsert(ds)
            .withTableName("products")
            .usingGeneratedKeyColumns("id");
    }

    public long save(String name, double price) {
        Map<String, Object> params = Map.of("name", name, "price", price);
        Number key = insert.executeAndReturnKey(params);
        return key.longValue();
    }
}
```
**A:** `SimpleJdbcInsert.executeAndReturnKey()` returns the auto-generated primary key. This avoids a `SELECT LAST_INSERT_ID()` call.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you retrieve auto-generated keys with Spring JDBC?
**Your Response:** SimpleJdbcInsert provides a clean way to retrieve auto-generated keys without writing additional SQL. You configure it with the table name and which columns are auto-generated, then call executeAndReturnKey() which returns the generated key as a Number. This is much cleaner than writing a separate SELECT LAST_INSERT_ID() call, which can be database-specific and error-prone. SimpleJdbcInsert handles the differences between databases and gives you a consistent API across all supported databases.

---

### 14. StoredProcedure — Calling Stored Procedures
**Q: How do you call a stored procedure with JdbcTemplate?**
```java
@Repository
class ReportDao {
    @Autowired JdbcTemplate jdbc;

    public Map<String, Object> generateReport(int year) {
        return jdbc.call(
            conn -> {
                CallableStatement cs = conn.prepareCall("{call generate_report(?, ?)}");
                cs.setInt(1, year);
                cs.registerOutParameter(2, Types.INTEGER); // OUT param
                return cs;
            },
            List.of(
                SqlParameter.make("year_in", Types.INTEGER),
                new SqlOutParameter("record_count", Types.INTEGER)
            )
        );
    }
}
```
**A:** Use `jdbc.call()` for stored procedures. `SqlOutParameter` declares output parameters. The returned `Map` contains output parameter values by name.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you call stored procedures with Spring JDBC?
**Your Response:** Spring JDBC provides the call() method for stored procedures. You provide a CallableStatementCreator that sets up the stored procedure call with input parameters, and declare the output parameters using SqlParameter objects. The method returns a Map containing the output parameter values indexed by name. This approach gives you type-safe access to both input and output parameters while handling the CallableStatement lifecycle automatically. It's much cleaner than working with CallableStatement directly.

---

### 15. JDBC Metadata — DatabaseMetaData
**Q: What can you get from DatabaseMetaData?**
```java
import java.sql.*;
public class Main {
    public static void main(String[] args) throws SQLException {
        Connection conn = /* ... */null;
        DatabaseMetaData meta = conn.getMetaData();
        System.out.println(meta.getDatabaseProductName()); // MySQL, PostgreSQL, etc.
        System.out.println(meta.getDatabaseProductVersion());
        System.out.println(meta.supportsTransactions()); // true
        System.out.println(meta.getMaxConnections()); // 0 = unlimited
        // List tables, columns, primary keys, foreign keys...
    }
}
```
**A:** `DatabaseMetaData` provides DB discovery — useful for migration tools, schema validators, and dynamic query builders. Liquibase and Flyway use it to detect DB type and version.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is DatabaseMetaData used for?
**Your Response:** DatabaseMetaData is JDBC's database discovery API that lets you inspect the database structure and capabilities at runtime. You can get information like the database product name and version, supported features, maximum limits, and even query the schema for tables, columns, primary keys, and foreign keys. This is incredibly useful for building database-agnostic tools like migration frameworks (Liquibase, Flyway), schema validators, or dynamic query generators that need to adapt to different database capabilities.

---

## Section 2: JPA Entity Lifecycle & EntityManager (Q16–Q28)

### 16. Entity States — New, Managed, Detached, Removed
**Q: What state is the entity in each case?**
```java
import jakarta.persistence.*;
public class Main {
    public static void main(String[] args) {
        EntityManagerFactory emf = Persistence.createEntityManagerFactory("myUnit");
        EntityManager em = emf.createEntityManager();

        em.getTransaction().begin();

        User u = new User("alice"); // NEW — no persistence context
        em.persist(u);             // MANAGED — tracked by persistence context
        System.out.println("id after persist: " + u.getId()); // may be null until flush

        em.flush();
        System.out.println("id after flush: " + u.getId()); // DB assigned ID

        em.getTransaction().commit();

        em.detach(u);              // DETACHED — no longer tracked
        u.setName("bob");          // change NOT tracked
        em.merge(u);               // re-attaches and merges changes

        em.remove(em.merge(u));    // REMOVED — queued for DELETE
        em.getTransaction().begin();
        em.getTransaction().commit(); // DELETE executed
    }
}
```
**A:** NEW → `persist()` → MANAGED → `commit()`/`detach()` → DETACHED → `merge()` → MANAGED → `remove()` → REMOVED → commit → gone from DB.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain the entity lifecycle states in JPA?
**Your Response:** JPA entities have four main states. NEW is when you create a new entity but haven't persisted it yet - it's not tracked by the persistence context. When you call persist(), it becomes MANAGED, meaning JPA tracks all changes to it. When the transaction ends or you explicitly call detach(), it becomes DETACHED - changes are no longer tracked. If you call remove(), it enters the REMOVED state and will be deleted from the database on commit. You can bring a detached entity back to managed state using merge(), which creates a managed copy and synchronizes the changes.

---

### 17. EntityManager.find() vs getReference()
**Q: What is the difference?**
```java
EntityManager em = /* ... */null;

User user = em.find(User.class, 1L);
// Immediately hits DB, returns null if not found

User proxy = em.getReference(User.class, 1L);
// Returns a Hibernate proxy immediately (no DB hit)
// DB hit occurs only when a field is accessed
// Throws EntityNotFoundException if accessed and entity doesn't exist
System.out.println(proxy.getId()); // no DB hit — ID is in proxy
System.out.println(proxy.getName()); // DB hit here if entity exists
```
**A:** `find()` eagerly loads from DB; returns `null` if missing. `getReference()` returns a lazy proxy; throws `EntityNotFoundException` on first field access if missing. Use `getReference()` for associations where you only need the foreign key.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between find() and getReference() in EntityManager?
**Your Response:** The key difference is when the database hit occurs. find() immediately queries the database and returns the actual entity or null if not found. getReference() returns a proxy object right away without hitting the database - it only queries the DB when you actually access a field on the proxy. If the entity doesn't exist, getReference() will throw EntityNotFoundException when you first try to access a field. Use getReference() when you only need the entity for setting a relationship or when you know the entity exists and want to avoid the immediate database hit.

---

### 18. EntityManager.merge() — Copy Semantics
**Q: What does merge() return? Are they the same object?**
```java
import jakarta.persistence.*;
public class Main {
    public static void main(String[] args) {
        EntityManager em = /* ... */null;
        em.getTransaction().begin();

        User detached = new User(1L, "alice"); // detached (has ID, not in context)
        User managed  = em.merge(detached);   // DIFFERENT object — a managed copy

        System.out.println(detached == managed); // false
        System.out.println(managed == em.find(User.class, 1L)); // true
        em.getTransaction().commit();
    }
}
```
**A:** `false`, `true`. `merge()` creates a new managed entity (or finds an existing one in the context) and copies state from the detached instance. The original detached object is NOT tracked.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does merge() return and is it the same object you passed in?
**Your Response:** merge() always returns a different object than the one you passed in. It creates a new managed entity instance or finds an existing one in the persistence context, then copies the state from your detached entity to this managed instance. Your original detached object remains untracked. This is why you should always use the returned object from merge() for further operations within the same transaction. The returned object is the one that's actually managed by the persistence context and will be synchronized with the database.

---

### 19. Dirty Checking — Auto-Update Without persist()
**Q: Is an UPDATE executed?**
```java
import jakarta.persistence.*;
public class Main {
    public static void main(String[] args) {
        EntityManager em = /* ... */null;
        em.getTransaction().begin();

        User u = em.find(User.class, 1L); // managed
        u.setName("charlie"); // just set a field — no explicit persist/update call

        em.getTransaction().commit(); // dirty check → generates UPDATE SQL
    }
}
```
**A:** Yes. JPA tracks changes to managed entities (dirty checking). On flush/commit, Hibernate compares the current state with the snapshot taken at load time and generates `UPDATE` for changed fields.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Will an UPDATE be executed if I just modify a field on a managed entity without calling persist()?
**Your Response:** Yes, absolutely! That's the beauty of JPA's dirty checking feature. When you load an entity, JPA takes a snapshot of its state. As long as the entity is managed, JPA automatically tracks any changes you make to its fields. When the transaction commits or when flush() is called, JPA compares the current state with the original snapshot and automatically generates UPDATE statements for any fields that changed. You don't need to call persist() or update() - just modify the object and JPA handles the rest.

---

### 20. Persistence Context — Scope
**Q: When is the persistence context closed?**
```java
@Service
class UserService {
    @PersistenceContext EntityManager em;

    @Transactional
    public User getAndModify(Long id) {
        User u = em.find(User.class, id); // managed
        u.setName("modified");
        return u; // returned to controller...
    } // transaction ends → persistence context closes → entity becomes DETACHED
}

@RestController
class UserController {
    @Autowired UserService service;

    @GetMapping("/users/{id}")
    public User get(@PathVariable Long id) {
        User u = service.getAndModify(id);
        u.getName(); // OK — field already loaded
        u.getOrders().size(); // LazyInitializationException if not fetched!
    }
}
```
**A:** With `@Transactional`, the persistence context lives for the duration of the transaction. After the method returns, the entity is DETACHED. Accessing lazy associations → `LazyInitializationException`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** When does the persistence context close and what happens to entities after?
**Your Response:** In Spring with @Transactional, the persistence context is bound to the transaction lifecycle. It opens when the transaction begins and closes when the transaction commits or rolls back. Once the persistence context closes, all managed entities become detached. This means if you try to access a lazy association outside the transaction, you'll get a LazyInitializationException because there's no active session to fetch the data. That's why you need to fetch all the data you need within the transaction or use techniques like JOIN FETCH to eagerly load relationships.

---

### 21. @GeneratedValue Strategies
**Q: What are the differences?**
```java
import jakarta.persistence.*;

@Entity class A {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    // DB auto-increment (MySQL: AUTO_INCREMENT, PostgreSQL: SERIAL)
    // ❌ Prevents JDBC batch insert (each row needs round-trip for ID)
    Long id;
}

@Entity class B {
    @Id @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "seq")
    @SequenceGenerator(name = "seq", sequenceName = "my_seq", allocationSize = 50)
    // Pre-fetches 50 IDs at a time → efficient with batch inserts
    Long id;
}

@Entity class C {
    @Id @GeneratedValue(strategy = GenerationType.UUID) // Java 17+ / Hibernate 6
    java.util.UUID id;
}
```
**A:** `IDENTITY`: simplest but disables batch inserts. `SEQUENCE`: best performance (pre-allocation). `UUID`: globally unique, no coordination needed. Spring Boot with Hibernate uses `SEQUENCE` by default for PostgreSQL.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the differences between IDENTITY, SEQUENCE, and UUID generation strategies?
**Your Response:** IDENTITY relies on database auto-increment columns - it's simple but prevents JDBC batch inserts because each insert needs to return to get the generated ID. SEQUENCE uses database sequences and can pre-allocate IDs in batches, making it much more efficient for bulk operations. UUID generates universally unique identifiers without requiring database coordination, which is great for distributed systems but takes more storage space. Spring Boot typically defaults to SEQUENCE for databases like PostgreSQL because it provides the best balance of performance and features.

---

### 22. @Embeddable — Value Objects
**Q: What does @Embeddable do?**
```java
import jakarta.persistence.*;

@Embeddable
class Address {
    String street, city, country;
}

@Entity
class Customer {
    @Id @GeneratedValue Long id;
    String name;

    @Embedded
    @AttributeOverrides({
        @AttributeOverride(name = "street", column = @Column(name = "billing_street"))
    })
    Address billingAddress;

    @Embedded Address shippingAddress; // maps to shipping_street, shipping_city, etc.
}
```
**A:** `@Embeddable` maps a value object to columns in the owner's table (no separate table). `@AttributeOverrides` renames columns when embedding the same type twice.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does @Embeddable do and when would you use it?
**Your Response:** @Embeddable allows you to map value objects - objects that don't have their own identity but are part of another entity. Instead of creating a separate table, the embeddable's fields are mapped directly to columns in the owning entity's table. This is perfect for things like addresses, monetary amounts, or coordinate pairs that logically belong together but don't need their own lifecycle. When you need to embed the same type multiple times in one entity, you can use @AttributeOverrides to customize the column names to avoid conflicts.

---

### 23. @MappedSuperclass — Shared Base Class
**Q: What does @MappedSuperclass do?**
```java
import jakarta.persistence.*;
import java.time.*;

@MappedSuperclass
abstract class BaseEntity {
    @Id @GeneratedValue Long id;
    @Column(updatable = false) LocalDateTime createdAt;
    LocalDateTime updatedAt;

    @PrePersist void prePersist() { createdAt = updatedAt = LocalDateTime.now(); }
    @PreUpdate  void preUpdate()  { updatedAt = LocalDateTime.now(); }
}

@Entity class Product extends BaseEntity { String name; }
@Entity class Order   extends BaseEntity { BigDecimal total; }
```
**A:** `@MappedSuperclass` shares mapping configuration (fields, lifecycle callbacks) without creating a separate table. Unlike `@Inheritance`, each subclass has its own table with the inherited columns duplicated.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between @MappedSuperclass and @Inheritance?
**Your Response:** @MappedSuperclass is for sharing common fields and behavior across entities without creating a separate table or inheritance hierarchy. Each subclass gets its own table with all the inherited fields duplicated. It's perfect for audit fields like createdAt and updatedAt that you want on every entity. @Inheritance creates a true inheritance mapping strategy where subclasses might share a table or have related tables, and it supports polymorphic queries. Use @MappedSuperclass for code reuse, and @Inheritance when you need actual object-oriented inheritance in your domain model.

---

### 24. Hibernate Second-Level Cache
**Q: When does the second-level cache help?**
```java
import jakarta.persistence.*;
import org.hibernate.annotations.*;

@Entity
@Cache(usage = CacheConcurrencyStrategy.READ_WRITE) // enable 2L cache
class Product {
    @Id Long id;
    String name;
    // ...
}

// First request: DB query → result cached in L2 cache
// Subsequent requests in DIFFERENT transactions: served from L2 cache
// Unlike L1 cache (per EntityManager), L2 cache spans EntityManagers
```
**A:** L1 cache = per EntityManager/transaction. L2 cache = shared across EntityManagers (process-level). Configure with Caffeine, Ehcache, or Redis. Invalidated on write. Use for mostly-read entities.

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use Hibernate's second-level cache?
**Your Response:** The second-level cache is perfect for entities that are read frequently but updated rarely - things like reference data, lookup tables, or configuration data. The first-level cache is per EntityManager and only lasts for the duration of a transaction, but the second-level cache is shared across all EntityManagers in the application. This means data cached in L2 can serve multiple transactions without hitting the database. You should enable it carefully though - only for data that changes infrequently, because the cache gets invalidated on writes. Popular implementations include Caffeine for in-memory caching or Redis for distributed caching.

---

### 25. Criteria API — Type-Safe Queries
**Q: What does this do?**
```java
import jakarta.persistence.criteria.*;

CriteriaBuilder cb = em.getCriteriaBuilder();
CriteriaQuery<User> q = cb.createQuery(User.class);
Root<User> user = q.from(User.class);

q.select(user)
 .where(cb.and(
     cb.greaterThan(user.get("age"), 18),
     cb.like(user.get("email"), "%@example.com")
 ))
 .orderBy(cb.asc(user.get("name")));

List<User> results = em.createQuery(q).getResultList();
```
**A:** Criteria API constructs JPQL programmatically with compile-time type safety. Verbose but avoids string concatenation. Used by dynamic query builders. Simpler alternative: Querydsl.

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use the Criteria API instead of JPQL?
**Your Response:** I'd use the Criteria API when I need to build dynamic queries at runtime - things like search screens where users can combine multiple optional filters. It provides compile-time type safety, so if I rename a field, the compiler will catch it rather than failing at runtime. It's more verbose than JPQL strings, but it prevents SQL injection errors and makes refactoring safer. For simple static queries, I prefer JPQL because it's more readable. For complex dynamic queries, many developers prefer Querydsl as a cleaner alternative to the Criteria API.

---

### 26. JPQL vs Native SQL
**Q: When do you use each?**
```java
// JPQL — works across databases, uses entity class names
List<User> users = em.createQuery(
    "SELECT u FROM User u WHERE u.age > :age", User.class)
    .setParameter("age", 18)
    .getResultList();

// Native SQL — database-specific, needed for DB features JPA doesn't support
List<Object[]> rows = em.createNativeQuery(
    "SELECT id, name FROM users WHERE age > ? LIMIT 10")
    .setParameter(1, 18)
    .getResultList();
```
**A:** JPQL is portable (translated to SQL by JPA provider). Native SQL is for DB-specific features (window functions, full-text search, hints). Use `@SqlResultSetMapping` or `Tuple` to map native query results.

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use JPQL versus native SQL?
**Your Response:** I use JPQL for most queries because it's database-agnostic - JPA translates it to the appropriate SQL dialect for the underlying database. JPQL works with entity names and relationships, making it more object-oriented. I switch to native SQL when I need to use database-specific features that JPQL doesn't support, like window functions, recursive queries, database hints, or when I need to optimize a critical query with vendor-specific syntax. When using native SQL, I need to manually map the results using @SqlResultSetMapping or extract values from Tuple objects.

---

### 27. EntityManager in Spring — @PersistenceContext vs @Autowired
**Q: Which is correct?**
```java
@Repository
class UserRepo {
    @PersistenceContext // CORRECT — Spring injects a transaction-scoped proxy
    EntityManager em;

    // @Autowired EntityManager em; // WRONG — injects a shared instance (not thread-safe)
}
```
**A:** `@PersistenceContext` injects a thread-safe proxy that delegates to the current transaction's `EntityManager`. `@Autowired EntityManager` injects the raw shared instance — breaks in multi-threaded environments.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why should you use @PersistenceContext instead of @Autowired for EntityManager?
**Your Response:** You must use @PersistenceContext because it injects a thread-safe proxy that's aware of the current transaction context. This proxy automatically delegates to the correct EntityManager for the current transaction. If you use @Autowired, you'd get the raw shared EntityManagerFactory, which isn't thread-safe and doesn't participate properly in Spring's transaction management. The @PersistenceContext proxy ensures that each transaction gets its own EntityManager instance while still allowing for dependency injection.

---

### 28. EntityManagerFactory vs EntityManager
**Q: What is the lifecycle?**
```java
// EntityManagerFactory: heavyweight, created once per app
EntityManagerFactory emf = Persistence.createEntityManagerFactory("myPU");
// In Spring Boot: auto-created from DataSource + JPA config

// EntityManager: lightweight, created per transaction/request
EntityManager em = emf.createEntityManager();
em.getTransaction().begin();
// do work
em.getTransaction().commit();
em.close(); // must close manually (Spring manages this with @Transactional)
```
**A:** `EntityManagerFactory` is thread-safe and expensive — one per application. `EntityManager` is not thread-safe — one per transaction/thread. Spring's `@Transactional` manages `EntityManager` lifecycle automatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between EntityManagerFactory and EntityManager?
**Your Response:** EntityManagerFactory is a heavyweight, thread-safe object that you create once per application - it's like a connection pool factory. It's expensive to create because it parses all the mapping metadata and sets up the database connections. EntityManager is lightweight and not thread-safe - you create one per transaction or operation. In Spring, you don't manage either manually - Spring Boot auto-creates the EntityManagerFactory, and @Transactional automatically manages the EntityManager lifecycle, ensuring each transaction gets its own EntityManager instance.

---

## Section 3: Relationships & N+1 Problem (Q29–Q38)

### 29. @OneToMany + @ManyToOne — Bidirectional
**Q: What is the owning side?**
```java
@Entity class Post {
    @Id @GeneratedValue Long id;
    String title;

    @OneToMany(mappedBy = "post", cascade = CascadeType.ALL)
    List<Comment> comments = new ArrayList<>();
}

@Entity class Comment {
    @Id @GeneratedValue Long id;
    String text;

    @ManyToOne @JoinColumn(name = "post_id") // OWNING SIDE — has FK column
    Post post;
}
```
**A:** The owning side is where `@JoinColumn` is — it controls the foreign key. `mappedBy` tells JPA the other end is the mirror. Always maintain both sides: `post.getComments().add(comment); comment.setPost(post);`

### How to Explain in Interview (Spoken style format)
**Interviewer:** In a bidirectional relationship, which side is the owning side and why does it matter?
**Your Response:** The owning side is the entity that has the @JoinColumn annotation - it's the one that actually controls the foreign key column in the database. The other side uses mappedBy to indicate it's the inverse or mirror side. This matters because JPA only looks at the owning side when deciding what SQL to generate. That's why you must maintain both sides programmatically - add the comment to the post's comments collection AND set the post on the comment. If you only set one side, you'll get inconsistent behavior depending on when the entity gets flushed to the database.

---

### 30. The N+1 Problem
**Q: How many SQL queries run?**
```java
@Repository
class OrderRepo extends JpaRepository<Order, Long> {}

// In service:
List<Order> orders = orderRepo.findAll(); // 1 query: SELECT * FROM orders
for (Order o : orders) {
    System.out.println(o.getCustomer().getName()); // N queries: SELECT * FROM customers WHERE id=?
}
// Total: 1 + N queries — N+1 problem!
```
**A:** If there are 100 orders, 101 SQL queries run. The N+1 problem is the most common JPA performance issue. **Fix:** use `JOIN FETCH`, `@EntityGraph`, or `@BatchSize`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the N+1 problem and how do you solve it?
**Your Response:** The N+1 problem is when you execute one query to fetch a list of entities, then execute N additional queries to fetch some lazy association for each entity. For example, you query 100 orders, then access the customer for each order, resulting in 101 total queries. This kills performance. The solution is to eagerly fetch the associations you need using JOIN FETCH in JPQL, @EntityGraph to define a fetch graph, or @BatchSize to load collections in batches. JOIN FETCH is most efficient as it uses a single SQL join.

---

### 31. JOIN FETCH — Solving N+1
**Q: How many queries now?**
```java
@Repository
interface OrderRepo extends JpaRepository<Order, Long> {
    @Query("SELECT o FROM Order o JOIN FETCH o.customer")
    List<Order> findAllWithCustomer();
}

// Service:
List<Order> orders = orderRepo.findAllWithCustomer(); // 1 query with JOIN
for (Order o : orders) {
    System.out.println(o.getCustomer().getName()); // no extra query — already loaded!
}
```
**A:** **1 query** with a SQL JOIN — eliminates N+1. Caveat: `DISTINCT` may be needed to avoid duplicate `Order` objects when joining `@OneToMany` collections. Use `@Query("SELECT DISTINCT o FROM Order o JOIN FETCH o.items")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does JOIN FETCH solve the N+1 problem?
**Your Response:** JOIN FETCH solves the N+1 problem by using a single SQL query with a JOIN to eagerly fetch the associated entities along with the main entities. Instead of 1 query for orders plus N queries for customers, you get one query that joins orders and customers and returns all the data at once. One gotcha is that when you JOIN FETCH a collection, you might get duplicate parent objects in the result, so you need to use DISTINCT in JPQL. Also, you can't use JOIN FETCH with pagination on collections - you'll need a different approach for that.

---

### 32. @EntityGraph — Load Graph
**Q: What does @EntityGraph do?**
```java
@Repository
interface UserRepo extends JpaRepository<User, Long> {
    @EntityGraph(attributePaths = {"orders", "orders.items"})
    List<User> findByActiveTrue();
}
// Generates: SELECT ... FROM users u
//            LEFT JOIN FETCH orders o ON o.user_id = u.id
//            LEFT JOIN FETCH order_items i ON i.order_id = o.id
```
**A:** `@EntityGraph` specifies which lazy associations to eagerly fetch for that query. Cleaner than writing `JOIN FETCH` manually. Produces a `LEFT JOIN FETCH` so users without orders are still returned.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does @EntityGraph do and when would you use it?
**Your Response:** @EntityGraph is a declarative way to specify which associations should be eagerly fetched for a particular query method. Instead of writing JOIN FETCH in your JPQL string, you annotate the method with @EntityGraph and list the attribute paths to fetch. It's cleaner and separates the fetching strategy from the query logic. It generates LEFT JOIN FETCH queries, so entities without the associated data are still included in the results. I use it when I need consistent fetching behavior across multiple query methods.

---

### 33. @BatchSize — Partial N+1 Fix
**Q: How does @BatchSize change the query count?**
```java
@Entity
class User {
    @OneToMany(mappedBy = "user")
    @BatchSize(size = 20) // load 20 collections at a time
    List<Order> orders;
}

// Loading 100 users:
// Without @BatchSize: 100 SELECT queries for orders
// With @BatchSize(20): ceil(100/20) = 5 SELECT queries (IN clause with 20 IDs)
```
**A:** `@BatchSize` reduces N+1 to `ceil(N/batchSize)` queries using SQL `IN` clauses. Not as efficient as `JOIN FETCH` but easier to apply globally. Configure globally: `hibernate.default_batch_fetch_size=20`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does @BatchSize help with the N+1 problem?
**Your Response:** @BatchSize is a partial solution to N+1 that reduces the number of queries by batching them. Instead of loading one collection per entity, it loads collections for multiple entities at once using SQL IN clauses. For example, with @BatchSize(20), loading 100 users with their orders would result in 5 queries instead of 100, each loading 20 collections at a time. It's not as efficient as JOIN FETCH because it still requires multiple round-trips, but it's easier to apply globally and works well with pagination. You can also set it globally as a Hibernate property.

---

### 34. @ManyToMany — Join Table
**Q: What table is created?**
```java
@Entity class Student {
    @Id @GeneratedValue Long id;
    String name;

    @ManyToMany
    @JoinTable(
        name = "student_course",
        joinColumns = @JoinColumn(name = "student_id"),
        inverseJoinColumns = @JoinColumn(name = "course_id")
    )
    Set<Course> courses = new HashSet<>();
}

@Entity class Course {
    @Id @GeneratedValue Long id;
    String name;

    @ManyToMany(mappedBy = "courses") // mirror side
    Set<Student> students = new HashSet<>();
}
```
**A:** Creates a `student_course` join table with columns `student_id` and `course_id`. Use `Set` instead of `List` to avoid Hibernate's duplicate-row multiply issue with `JOIN FETCH`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What table structure does @ManyToMany create and why should you use Set instead of List?
**Your Response:** @ManyToMany creates a join table that contains foreign keys to both entities - in this case, student_course table with student_id and course_id columns. The join table manages the many-to-many relationship. You should use Set instead of List because when you JOIN FETCH a many-to-many collection, Hibernate can return duplicate parent entities due to the join producing multiple rows. Using Set automatically eliminates duplicates, while List would contain the same entity multiple times. This is especially important when you're fetching entities that have many-to-many relationships.

---

### 35. Orphan Removal
**Q: What happens when a comment is removed from the list?**
```java
@Entity
class Post {
    @OneToMany(mappedBy = "post", cascade = CascadeType.ALL, orphanRemoval = true)
    List<Comment> comments = new ArrayList<>();
}

// In service:
@Transactional
public void deleteFirstComment(Long postId) {
    Post post = em.find(Post.class, postId);
    post.getComments().remove(0); // removes from list
    // orphanRemoval = true → DELETE FROM comments WHERE id = ? is generated
}
```
**A:** The removed `Comment` becomes an "orphan" (no parent). With `orphanRemoval = true`, JPA deletes it from the DB. Without `orphanRemoval`, it's just removed from the list — the DB row remains (orphaned foreign key).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does orphanRemoval = true do?
**Your Response:** Orphan removal automatically deletes child entities when they're removed from their parent's collection. When you remove a comment from the post's comments list, that comment becomes an "orphan" - it no longer has a parent reference. With orphanRemoval = true, JPA will automatically issue a DELETE statement to remove that comment row from the database. Without orphanRemoval, the comment would just be removed from the collection but the database row would remain, leaving an orphaned record with a null or invalid foreign key. This is great for cascade delete scenarios.

---

### 36. Fetch Join with Pagination — The Trap
**Q: What is the warning?**
```java
@Query("SELECT u FROM User u JOIN FETCH u.orders")
Page<User> findUsersWithOrders(Pageable page);
// WARNING in log: HHH90003004: firstResult/maxResults specified with collection fetch;
// applying in memory! (inefficient for large datasets)
```
**A:** Hibernate cannot apply SQL-level `LIMIT`/`OFFSET` when using `JOIN FETCH` with a collection — it fetches ALL rows into memory and paginates in Java. **Fix:** use a two-query approach: first query for IDs with pagination, then `findAllById(ids)` with `JOIN FETCH`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why can't you use JOIN FETCH with pagination on collections?
**Your Response:** This is a classic Hibernate limitation. When you JOIN FETCH a collection, the SQL result contains duplicate parent rows - one for each child in the collection. Hibernate needs to fetch all these rows and then deduplicate them in memory to figure out the actual pagination. This means your database might return 10,000 rows even though you only want 20 unique entities. The solution is to use a two-query approach: first paginate the parent IDs only, then fetch those specific parents with their collections using JOIN FETCH. This way you get proper database-level pagination without the memory overhead.

---

### 37. @OneToOne — Shared Primary Key
**Q: What does this mapping do?**
```java
@Entity class User {
    @Id @GeneratedValue Long id;
    String email;

    @OneToOne(mappedBy = "user", cascade = CascadeType.ALL, fetch = FetchType.LAZY)
    UserProfile profile;
}

@Entity class UserProfile {
    @Id Long id; // same PK as User

    @OneToOne
    @MapsId // user.id == userProfile.id
    @JoinColumn(name = "id")
    User user;

    String bio;
}
```
**A:** `@MapsId` makes `UserProfile.id` share the value with `User.id`. No separate FK column — the primary key is also the foreign key. Efficient: no extra column, guaranteed uniqueness.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does @MapsId do in a @OneToOne relationship?
**Your Response:** @MapsId creates a shared primary key relationship where the child entity uses the same primary key as the parent entity. The child's primary key column serves as both the primary key and foreign key to the parent. This is the most efficient way to model one-to-one relationships because you don't need a separate foreign key column. It also guarantees one-to-one cardinality at the database level since the child can't exist without the parent - they share the same identifier. It's perfect for things like User and UserProfile where the profile logically can't exist without the user.

---

### 38. Soft Delete — @SQLRestriction (Hibernate 6+)
**Q: What does this do?**
```java
import org.hibernate.annotations.*;
import jakarta.persistence.*;

@Entity
@SQLRestriction("deleted = false") // all queries auto-add WHERE deleted = false
class Product {
    @Id @GeneratedValue Long id;
    String name;
    boolean deleted = false;
}

// productRepo.findAll() → SELECT * FROM product WHERE deleted = false
// em.remove(p) → UPDATE product SET deleted = true WHERE id = ?
// Use @SQLDelete for custom delete SQL:
// @SQLDelete(sql = "UPDATE product SET deleted = true WHERE id = ?")
```
**A:** Soft delete keeps rows in the DB but hides them. `@SQLRestriction` adds a filter to all queries automatically. `@SQLDelete` overrides the DELETE SQL with an UPDATE. Old annotation: `@Where(clause = "deleted = false")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement soft delete in JPA?
**Your Response:** Soft delete is when you mark records as deleted instead of actually removing them. In Hibernate 6+, you can use @SQLRestriction to automatically add a WHERE clause to all queries for that entity, filtering out the deleted records. When you call em.remove(), you can use @SQLDelete to override the DELETE statement with an UPDATE that sets a deleted flag. This way, deleted data is preserved for auditing or recovery, but is hidden from normal application queries. Older Hibernate versions used @Where annotation for the same purpose. This approach is great for data retention requirements.

---

## Section 4: Transactions — Propagation & Isolation (Q39–Q47)

### 39. REQUIRED vs REQUIRES_NEW — Full Example
**Q: What is rolled back when method B fails?**
```java
@Service class A {
    @Autowired B b;
    @Transactional
    public void doA() {
        saveA();       // in transaction T1
        b.doB();       // REQUIRED: joins T1
        b.doBNew();    // REQUIRES_NEW: suspends T1, starts T2
    }
}

@Service class B {
    @Transactional(propagation = Propagation.REQUIRED)
    public void doB() { saveB(); throw new RuntimeException("fail"); } // T1 rollback

    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void doBNew() { saveC(); throw new RuntimeException("fail"); } // T2 rollback, T1 continues
}
```
**A:** `REQUIRED` — any exception rolls back the whole transaction (A + B). `REQUIRES_NEW` — T2 fails → only T2 rolls back; T1 can continue unless the exception propagates to A.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between REQUIRED and REQUIRES_NEW propagation?
**Your Response:** REQUIRED is the default - it joins the existing transaction if one exists, or creates a new one if not. If any method in the transaction fails, the entire transaction rolls back. REQUIRES_NEW always suspends the current transaction and starts a new one. This means if the REQUIRES_NEW method fails, only that new transaction rolls back - the original transaction can continue independently. This is useful for logging or audit operations that should commit even if the main transaction fails, or for independent operations that shouldn't be affected by the main transaction's outcome.

---

### 40. NESTED — Savepoints
**Q: What is Propagation.NESTED?**
```java
@Transactional
public void outer() {
    saveOuter();
    try {
        inner(); // NESTED saves a savepoint here
    } catch (RuntimeException e) {
        // rollback to savepoint — saveOuter() is preserved
        log.warn("inner failed but outer continues");
    }
    saveFinal(); // committed along with saveOuter()
}

@Transactional(propagation = Propagation.NESTED)
public void inner() {
    saveInner();
    throw new RuntimeException("partial failure");
}
```
**A:** `NESTED` creates a JDBC savepoint. On exception, rolls back to the savepoint only. The outer transaction is unaffected if the exception is caught. Not supported by all databases/JPA providers (works with Hibernate + JDBC).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does Propagation.NESTED do and when would you use it?
**Your Response:** NESTED creates a savepoint within the current transaction, allowing partial rollbacks. If the nested method fails, it rolls back only to the savepoint - the work done before the savepoint in the outer transaction is preserved. This is different from REQUIRES_NEW which creates a completely separate transaction. NESTED is useful when you want to attempt an operation that might fail, but you don't want it to undo all the work done so far. However, it's not universally supported - it works with Hibernate and JDBC but not with all JPA providers or databases.

---

### 41. Isolation Levels — Dirty Read
**Q: What does READ_UNCOMMITTED allow?**
```java
@Transactional(isolation = Isolation.READ_UNCOMMITTED)
public int getBalance(Long accountId) {
    // Can read uncommitted changes from other transactions
    // "Dirty read" = reading data that may be rolled back
    return jdbc.queryForObject("SELECT balance FROM accounts WHERE id=?", Integer.class, accountId);
}

//    Thread A: UPDATE balance = 1000 (uncommitted)
//    Thread B (READ_UNCOMMITTED): reads 1000
//    Thread A: ROLLBACK (balance back to 500)
//    Thread B now has stale data!
```
**A:** `READ_UNCOMMITTED` allows dirty reads — fastest but provides no consistency guarantees. Almost never used in practice. Only safe for approximate counts/statistics where accuracy isn't critical.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is READ_UNCOMMITTED isolation level and when would you use it?
**Your Response:** READ_UNCOMMITTED is the lowest isolation level that allows dirty reads - you can see uncommitted changes from other transactions, even if they might be rolled back. This is extremely fast but provides almost no consistency guarantees. In practice, it's almost never used because the risk of reading invalid data is too high. The only scenario where it might make sense is for approximate analytics or monitoring queries where exact accuracy isn't critical and you want maximum performance.

---

### 42. Isolation Levels — Non-Repeatable Read
**Q: What does READ_COMMITTED prevent?**
```java
// READ_COMMITTED (default in PostgreSQL, Oracle):
@Transactional(isolation = Isolation.READ_COMMITTED)
public void check() {
    int balance1 = getBalance(1L); // reads committed data: 500
    // Another transaction commits: balance changed to 1000
    int balance2 = getBalance(1L); // reads 1000 — non-repeatable read!
    System.out.println(balance1 == balance2); // false — different!
}

// REPEATABLE_READ: balance2 == balance1 = 500 (snapshot at start of transaction)
```
**A:** `READ_COMMITTED` prevents dirty reads but allows non-repeatable reads (same row reads different values within a transaction). `REPEATABLE_READ` prevents this (MySQL InnoDB default).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between READ_COMMITTED and REPEATABLE_READ isolation levels?
**Your Response:** READ_COMMITTED prevents dirty reads by only showing you committed data from other transactions, but it allows non-repeatable reads - if you read the same row twice in a transaction and another transaction modifies and commits it in between, you'll see different values. REPEATABLE_READ prevents this by taking a snapshot at the start of your transaction and ensuring all reads see that same consistent view. Most databases like PostgreSQL default to READ_COMMITTED, while MySQL InnoDB defaults to REPEATABLE_READ for better consistency.

---

### 43. Isolation Levels — Phantom Read
**Q: What is a phantom read?**
```java
// REPEATABLE_READ doesn't prevent phantom reads:
@Transactional(isolation = Isolation.REPEATABLE_READ)
public void check() {
    List<User> users1 = findUsersOlderThan(18); // 10 users
    // Another transaction inserts a new user age=20 and commits
    List<User> users2 = findUsersOlderThan(18); // 11 users — phantom!
}

// SERIALIZABLE prevents phantom reads:
@Transactional(isolation = Isolation.SERIALIZABLE)
public void safeCheck() {
    // No other transaction can insert into the range while this runs
}
```
**A:** Phantom reads = new rows appearing in a range query within the same transaction. `SERIALIZABLE` prevents all anomalies but is the slowest (table-level locking or predicate locking).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a phantom read and how does SERIALIZABLE prevent it?
**Your Response:** A phantom read occurs when new rows appear in a range query within the same transaction. For example, you query all users older than 18 and get 10 results, then another transaction inserts a new 20-year-old user, and when you re-run the same query you get 11 results - the new row is a phantom. SERIALIZABLE isolation prevents this by locking the entire range of rows, preventing other transactions from inserting new rows that would appear in your query. This provides complete isolation but at a significant performance cost due to extensive locking.

---

### 44. Optimistic Locking Implementation
**Q: How do you implement optimistic locking correctly?**
```java
@Entity class Account {
    @Id Long id;
    long balance;
    @Version int version;
}

@Service class AccountService {
    @Transactional
    public void transfer(Long id, long amount) {
        Account a = repo.findById(id).orElseThrow();
        a.setBalance(a.getBalance() - amount);
        // Hibernate adds: UPDATE account SET balance=?, version=version+1
        //                 WHERE id=? AND version=<old_version>
        // If 0 rows updated → OptimisticLockException
    }

    public void transferWithRetry(Long id, long amount) {
        int retries = 3;
        while (retries-- > 0) {
            try { transfer(id, amount); return; }
            catch (OptimisticLockException e) { /* retry */ }
        }
        throw new RuntimeException("too many conflicts");
    }
}
```
**A:** Optimistic locking is suitable for low-contention scenarios. Always implement retry logic. JPA throws `OptimisticLockException` → Spring wraps it in `ObjectOptimisticLockingFailureException`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement optimistic locking in JPA?
**Your Response:** Optimistic locking uses a version field that JPA increments on each update. When you try to update an entity, JPA includes the version in the WHERE clause. If another transaction updated the entity first, the version won't match and no rows will be updated, causing an OptimisticLockException. This approach assumes conflicts are rare and is more efficient than pessimistic locking. You should always implement retry logic - catch the exception and retry the operation a few times before giving up. Spring wraps the JPA exception in ObjectOptimisticLockingFailureException for consistency.

---

### 45. Pessimistic Locking — LockModeType
**Q: What does PESSIMISTIC_WRITE do?**
```java
@Repository
interface AccountRepo extends JpaRepository<Account, Long> {
    @Lock(LockModeType.PESSIMISTIC_WRITE) // SELECT ... FOR UPDATE
    @Query("SELECT a FROM Account a WHERE a.id = :id")
    Optional<Account> findByIdForUpdate(@Param("id") Long id);
}

// In service:
@Transactional
public void process(Long id) {
    Account acc = accountRepo.findByIdForUpdate(id).orElseThrow();
    // Other transactions trying to lock same row will BLOCK until this commits
    acc.setBalance(acc.getBalance() - 100);
}
```
**A:** `PESSIMISTIC_WRITE` → `SELECT FOR UPDATE`. Other transactions wait. Use for high-contention operations (inventory decrement, seat booking). `PESSIMISTIC_READ` → `SELECT FOR SHARE` (allows other reads, blocks writes).

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use pessimistic locking?
**Your Response:** Pessimistic locking is for high-contention scenarios where conflicts are likely and you want to avoid retry logic. PESSIMISTIC_WRITE locks the row for updates using SELECT FOR UPDATE, preventing other transactions from modifying or even reading the row in some databases. PESSIMISTIC_READ uses SELECT FOR SHARE which allows other transactions to read but blocks writes. This is perfect for things like inventory management where you need to decrement stock, seat booking systems, or any scenario where multiple users might compete for the same resource simultaneously.

---

### 46. @Transactional readOnly = true
**Q: What is the benefit?**
```java
@Service
class ReportService {
    @Transactional(readOnly = true)
    public List<SalesData> getMonthlySales() {
        return repo.findAll(); // read-only hint
    }
}
```
**A:** `readOnly = true`: (1) Hibernate skips dirty checking (faster). (2) Database driver may route to a read replica. (3) Flushes are skipped. Use for all read-only methods. Does NOT prevent writes — it's an optimization hint.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does readOnly = true do in @Transactional?
**Your Response:** The readOnly flag is an optimization hint that tells Spring and Hibernate this transaction won't modify data. This provides several optimizations: Hibernate skips dirty checking since no changes are expected, the flush operation is skipped, and some database drivers might route the query to a read replica instead of the primary. It's important to note that this doesn't actually prevent writes - it's purely an optimization. You should still use it for all read-only operations like queries and reports to get better performance.

---

### 47. TransactionSynchronizationManager — Post-Commit Hook
**Q: How do you run code after a transaction commits?**
```java
@Service class EventPublisher {
    @Autowired ApplicationEventPublisher publisher;

    @Transactional
    public void save(User user) {
        repo.save(user);
        // Don't publish event immediately — DB not yet committed!
        TransactionSynchronizationManager.registerSynchronization(
            new TransactionSynchronizationAdapter() {
                public void afterCommit() {
                    publisher.publishEvent(new UserCreatedEvent(user));
                }
            }
        );
    }
}
```
**A:** Without the post-commit hook, the event fires before the DB commit — consumers may not find the data yet. `afterCommit()` ensures the event is published only after the transaction is durably committed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you run code after a transaction successfully commits?
**Your Response:** You use TransactionSynchronizationManager to register a synchronization callback. The key insight is that if you publish an event directly within a @Transactional method, the event fires immediately but the database transaction hasn't committed yet. Event listeners might try to read the data and find it doesn't exist. By registering an afterCommit callback, you ensure the event only fires after the transaction has successfully committed to the database. This guarantees that event consumers can reliably find the data they're supposed to process.

---

## Section 5: Spring Data JPA — Queries & Locking (Q48–Q55)

### 48. Derived Query Methods — Naming Rules
**Q: What SQL does each method generate?**
```java
interface UserRepo extends JpaRepository<User, Long> {
    List<User> findByFirstNameAndLastName(String first, String last);
    // WHERE first_name = ? AND last_name = ?

    List<User> findTop5ByAgeGreaterThanOrderByNameAsc(int age);
    // WHERE age > ? ORDER BY name ASC LIMIT 5

    boolean existsByEmail(String email);
    // SELECT COUNT(*) > 0 FROM users WHERE email = ?

    long countByActiveTrue();
    // SELECT COUNT(*) FROM users WHERE active = true

    void deleteByCreatedAtBefore(LocalDate date);
    // DELETE FROM users WHERE created_at < ?
}
```
**A:** Spring Data parses method names at startup. Keywords: `findBy`, `existsBy`, `countBy`, `deleteBy`, `Top/First`, `And/Or`, `GreaterThan`, `Like`, `In`, `IsNull`, `OrderBy`. If parsing fails, Spring throws an error at startup — not at runtime.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do Spring Data derived query methods work?
**Your Response:** Spring Data parses the method name at startup and automatically generates the JPQL query based on naming conventions. It looks for keywords like findBy, existsBy, countBy, deleteBy followed by property names and operators like And, Or, GreaterThan, Like, In, IsNull, and OrderBy. The great thing is that if you make a mistake in the method name, Spring fails at startup rather than at runtime, so you catch errors early. This eliminates most boilerplate query code while maintaining type safety.

---

### 49. @Query — JPQL and SpEL
**Q: What does #{#entityName} do?**
```java
@Repository
interface ProductRepo extends JpaRepository<Product, Long> {
    @Query("SELECT p FROM #{#entityName} p WHERE p.active = true AND p.price < :maxPrice")
    List<Product> findActive(@Param("maxPrice") BigDecimal max);

    @Query(value = "SELECT * FROM products WHERE stock > 0", nativeQuery = true)
    List<Product> findInStock();

    @Query("SELECT p FROM Product p WHERE p.name IN :names")
    List<Product> findByNames(@Param("names") Collection<String> names);
}
```
**A:** `#{#entityName}` is replaced with the entity name — useful in generic repositories. `IN :names` works with any `Collection`. `nativeQuery = true` sends raw SQL; must use positional `?1` or named `:param` parameters.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are some advanced features of @Query in Spring Data JPA?
**Your Response:** @Query supports several advanced features. The #{#entityName} SpEL expression gets replaced with the actual entity name, which is great for generic repositories where you don't know the entity type at compile time. You can use IN clauses with any Collection parameter - Spring automatically expands it to the appropriate JPQL. When you set nativeQuery = true, you can write database-specific SQL for features JPQL doesn't support, but then you need to use positional parameters like ?1 or named parameters like :param instead of the JPA parameter binding.

---

### 50. Projections — DTOs Without Full Entity Load
**Q: Why use projections?**
```java
// Interface projection (lazy — only specified fields fetched)
interface UserSummary {
    Long getId();
    String getEmail();
}

// Class projection (DTO constructor)
record UserDto(Long id, String email) {}

@Repository
interface UserRepo extends JpaRepository<User, Long> {
    List<UserSummary> findAllByActiveTrue(); // interface projection
    @Query("SELECT new com.example.UserDto(u.id, u.email) FROM User u WHERE u.active = true")
    List<UserDto> findActiveDtos(); // class projection via JPQL constructor expression
}
```
**A:** Projections fetch only needed columns (smaller SQL result) and avoid loading entity state into the persistence context (no dirty checking overhead). Use for read-only responses, especially in list endpoints.

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use Spring Data JPA projections?
**Your Response:** Projections are perfect for read-only operations where you only need a subset of entity fields. Instead of loading the entire entity with all its columns and putting it in the persistence context, projections fetch only the specific columns you need. This reduces memory usage and network traffic, and avoids dirty checking overhead since the projected objects aren't managed entities. I use them heavily in REST API endpoints and reports where I'm returning data that won't be modified, like user summaries or product listings.

---

### 51. Pagination — Pageable
**Q: What does this return?**
```java
@GetMapping("/users")
public Page<User> list(@RequestParam(defaultValue = "0") int page,
                       @RequestParam(defaultValue = "20") int size) {
    Pageable pageable = PageRequest.of(page, size, Sort.by("name").ascending());
    Page<User> result = userRepo.findAll(pageable);

    System.out.println("total: "   + result.getTotalElements());
    System.out.println("pages: "   + result.getTotalPages());
    System.out.println("current: " + result.getContent().size());
    return result;
}
```
**A:** Returns a `Page<User>` containing the content + metadata (total elements, total pages, current page). Spring Data executes a `SELECT ... LIMIT/OFFSET` + a `SELECT COUNT(*)` query.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does pagination work in Spring Data JPA?
**Your Response:** Spring Data makes pagination easy with the Pageable parameter and Page return type. When you pass a Pageable object with page number, size, and sort information, Spring Data automatically generates the appropriate SQL with LIMIT and OFFSET for the data query, plus a separate COUNT query to get the total number of elements. The returned Page object contains both the actual content for the current page and metadata like total elements, total pages, and whether there are more pages available. This gives you everything you need to build pagination controls in the UI.

---

### 52. Specifications — Dynamic Queries
**Q: How do you build dynamic WHERE clauses?**
```java
import org.springframework.data.jpa.domain.Specification;

class UserSpec {
    static Specification<User> hasEmail(String email) {
        return (root, query, cb) -> email == null ? null : cb.equal(root.get("email"), email);
    }
    static Specification<User> isActive() {
        return (root, query, cb) -> cb.isTrue(root.get("active"));
    }
}

@Repository
interface UserRepo extends JpaRepository<User, Long>, JpaSpecificationExecutor<User> {}

// Usage:
List<User> users = userRepo.findAll(
    Specification.where(UserSpec.isActive()).and(UserSpec.hasEmail(filter)));
```
**A:** `Specification` builds type-safe, composable WHERE clauses. Return `null` from a Specification to skip that predicate. Combine with `and()`, `or()`, `not()`. Good for search filters with optional parameters.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle dynamic queries with optional filters in Spring Data JPA?
**Your Response:** I use the Specification API for dynamic queries where filters are optional. Each Specification represents a single WHERE clause condition that you can build programmatically with type safety. You return null from a Specification to indicate that particular filter shouldn't be applied. Then you combine multiple specifications using and(), or(), and not() operators. This approach is much cleaner than building JPQL strings and provides compile-time checking. It's perfect for search screens where users can combine multiple optional criteria.

---

### 53. Auditing — @CreatedDate, @LastModifiedDate
**Q: How does Spring Data auditing work?**
```java
@Configuration
@EnableJpaAuditing
class JpaConfig {}

@MappedSuperclass
@EntityListeners(AuditingEntityListener.class)
abstract class AuditableEntity {
    @CreatedBy     String createdBy;
    @LastModifiedBy String modifiedBy;
    @CreatedDate   @Column(updatable = false) LocalDateTime createdAt;
    @LastModifiedDate LocalDateTime updatedAt;
}

// Implement AuditorAware to supply the current user:
@Component
class SecurityAuditorAware implements AuditorAware<String> {
    public Optional<String> getCurrentAuditor() {
        return Optional.ofNullable(SecurityContextHolder.getContext())
                       .map(ctx -> ctx.getAuthentication().getName());
    }
}
```
**A:** Spring Data auditing auto-populates audit fields on `@PrePersist`/`@PreUpdate`. No manual `setCreatedAt()` calls needed. `AuditorAware` provides the current user from the security context.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Data JPA auditing work?
**Your Response:** Spring Data auditing automatically manages audit fields like createdAt, updatedAt, createdBy, and modifiedBy. You annotate your entity with @EntityListeners and use @CreatedDate, @LastModifiedDate, @CreatedBy, and @LastModifiedBy annotations. Spring then automatically populates these fields before persisting and updating entities. For the user fields, you implement an AuditorAware bean that extracts the current user from the security context. This eliminates boilerplate code and ensures consistent audit tracking across all entities without manual intervention.

---

### 54. @Modifying with clearAutomatically
**Q: Why is clearAutomatically important?**
```java
@Modifying(clearAutomatically = true, flushAutomatically = true)
@Transactional
@Query("UPDATE User u SET u.active = false WHERE u.lastLogin < :cutoff")
int deactivateOld(@Param("cutoff") LocalDate cutoff);
```
**A:** After a bulk UPDATE/DELETE, the first-level cache (persistence context) is stale — it still has the old entity state. `clearAutomatically = true` evicts the cache after the query. `flushAutomatically = true` flushes pending changes first so the bulk query sees them.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why do you need clearAutomatically and flushAutomatically in @Modifying queries?
**Your Response:** When you execute bulk UPDATE or DELETE operations with @Modifying, they bypass the persistence context and execute directly against the database. This creates a mismatch - the entities still cached in your persistence context have the old values, but the database has been updated. clearAutomatically = true clears the cache after the bulk operation to prevent this inconsistency. flushAutomatically = true ensures any pending changes are written to the database before the bulk operation runs, so the bulk query sees the current state. Without these, you can get stale data or lost updates.

---

### 55. Custom Repository Implementation
**Q: How do you add non-standard methods to a JPA repository?**
```java
// Step 1: define custom interface
interface UserRepoCustom {
    List<User> searchByComplexCriteria(String keyword, Map<String, Object> filters);
}

// Step 2: implement it
class UserRepoCustomImpl implements UserRepoCustom {
    @PersistenceContext EntityManager em;

    public List<User> searchByComplexCriteria(String keyword, Map<String, Object> filters) {
        // use Criteria API or JPQL dynamically
        CriteriaBuilder cb = em.getCriteriaBuilder();
        // ... build complex query
        return em.createQuery(/* ... */).getResultList();
    }
}

// Step 3: extend both
@Repository
interface UserRepo extends JpaRepository<User, Long>, UserRepoCustom {}
```
**A:** Spring Data detects `UserRepoCustomImpl` (naming convention: `repoInterfaceName + "Impl"`) and delegates custom method calls to it. This pattern keeps JPA repositories extensible without losing Spring Data's auto-generated methods.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you add custom methods to Spring Data JPA repositories?
**Your Response:** You use the custom repository pattern when you need methods that Spring Data can't generate automatically. First, create a separate interface with your custom methods. Then implement that interface in a class named RepositoryNameImpl - Spring Data finds this implementation by naming convention. Finally, make your main repository interface extend both JpaRepository and your custom interface. This way, you get all the auto-generated CRUD methods plus your custom methods, and Spring automatically wires everything together.

---

> 🔖 **Last read:** <!-- update here -->
