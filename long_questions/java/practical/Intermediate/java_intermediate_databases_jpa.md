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

---

> 🔖 **Last read:** <!-- update here -->
