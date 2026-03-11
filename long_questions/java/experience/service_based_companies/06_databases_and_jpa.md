# 🗄️ 06 — Databases & JPA/Hibernate
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- JPA entities and mappings
- Spring Data JPA repositories
- JPQL and native queries
- Lazy vs eager fetching
- N+1 problem and solutions
- Transaction management
- Database migrations (Flyway/Liquibase)

---

## ❓ Most Asked Questions

### Q1. How do you define a JPA entity?

```java
@Entity
@Table(name = "users",
    uniqueConstraints = @UniqueConstraint(columnNames = "email"),
    indexes = { @Index(name = "idx_email", columnList = "email") })
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)  // auto-increment
    private Long id;

    @Column(name = "first_name", nullable = false, length = 50)
    private String firstName;

    @Column(unique = true, nullable = false)
    private String email;

    @Enumerated(EnumType.STRING)   // store as "ACTIVE"/"INACTIVE" not 0/1
    private UserStatus status;

    @CreationTimestamp
    @Column(updatable = false)
    private LocalDateTime createdAt;

    @UpdateTimestamp
    private LocalDateTime updatedAt;

    @Version           // optimistic locking — prevent lost updates
    private Long version;
}
```

---

### 🎯 How to Explain in Interview

"A JPA entity is a Java class that maps to a database table. I start with @Entity to mark the class as an entity, and @Table to specify the table name and constraints. For the primary key, I use @Id with @GeneratedValue for auto-increment. For regular columns, @Column lets me control constraints like nullable, length, and uniqueness. I use @Enumerated to store enums as readable strings instead of numbers. For timestamps, @CreationTimestamp and @UpdateTimestamp automatically track when records are created and modified. The @Version field is crucial for optimistic locking - it prevents concurrent modification issues by checking that the version hasn't changed since I read the record. This entity mapping approach lets me work with objects while JPA handles the SQL generation."

---

### Q2. What are JPA relationship mappings?

```java
// @OneToMany / @ManyToOne — most common (Order → OrderItems)
@Entity
public class Order {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToMany(mappedBy = "order",    // "order" = field name in OrderItem
               cascade = CascadeType.ALL,    // operations propagate to children
               orphanRemoval = true,         // remove orphaned items
               fetch = FetchType.LAZY)       // load items on demand (DEFAULT for *ToMany)
    private List<OrderItem> items = new ArrayList<>();
}

@Entity
public class OrderItem {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "order_id", nullable = false)  // FK column
    private Order order;

    private int quantity;
    private BigDecimal price;
}

// @ManyToMany — User ↔ Role
@Entity
public class User {
    @ManyToMany(fetch = FetchType.LAZY)
    @JoinTable(name = "user_roles",
        joinColumns = @JoinColumn(name = "user_id"),
        inverseJoinColumns = @JoinColumn(name = "role_id"))
    private Set<Role> roles = new HashSet<>();
}

// @OneToOne — User ↔ UserProfile
@Entity
public class UserProfile {
    @Id
    private Long id;

    @OneToOne
    @MapsId   // shares PK with User
    private User user;
}
```

---

### 🎯 How to Explain in Interview

"JPA relationship mappings define how entities relate to each other. The most common is @OneToMany and @ManyToOne - like an Order with many OrderItems. I use mappedBy on the parent side to avoid creating a join table, and cascade to propagate operations to children. For many-to-many relationships like Users and Roles, I use @JoinTable to create the junction table. For one-to-one, I can use @MapsId to share the primary key. The key decision is fetch type - LAZY for collections to avoid loading too much data, EAGER for single-valued references. I also think about cascade operations - do I want to delete children when I delete the parent? These mappings let me model complex relationships while JPA handles the underlying foreign keys and join tables."

---

### Q3. What is the N+1 problem and how to fix it?

```java
// THE N+1 PROBLEM:
// Load N orders → each lazily loads items → N extra queries!
List<Order> orders = orderRepository.findAll();
for (Order o : orders) {
    System.out.println(o.getItems().size()); // 1 query per order → N+1 total!
}

// FIX 1: JPQL JOIN FETCH — single query with JOIN
@Query("SELECT o FROM Order o JOIN FETCH o.items WHERE o.status = :status")
List<Order> findWithItemsByStatus(@Param("status") OrderStatus status);

// FIX 2: @EntityGraph — declarative eager loading
@EntityGraph(attributePaths = {"items", "items.product"})
List<Order> findByCustomerId(Long customerId);

// FIX 3: @BatchSize — batches lazy loads (N+1 → ceil(N/batchSize)+1 queries)
@BatchSize(size = 25)
@OneToMany(mappedBy = "order", fetch = FetchType.LAZY)
private List<OrderItem> items;

// FIX 4: DTO projection — select only what you need
@Query("SELECT new com.example.dto.OrderSummary(o.id, o.total, COUNT(i)) " +
       "FROM Order o JOIN o.items i GROUP BY o.id, o.total")
List<OrderSummary> findOrderSummaries();
```

---

### 🎯 How to Explain in Interview

"The N+1 problem is a classic performance issue in JPA. It happens when I load a list of entities and then access their lazy-loaded associations - JPA runs one query to load the main entities, plus one additional query for each entity's association. For 100 orders, that's 101 queries! The solutions are: JOIN FETCH in JPQL to load everything in one query, @EntityGraph for declarative eager loading, @BatchSize to batch the lazy loads into fewer queries, or DTO projections to select only the data I need. I choose the solution based on the use case - JOIN FETCH for related data I always need, DTO projections for reporting, and @BatchSize as a general optimization. Understanding and fixing N+1 is crucial for performant JPA applications."

---

### Q4. What is Spring Data JPA and custom queries?

```java
// Repository with built-in CRUD + paging
public interface UserRepository extends JpaRepository<User, Long> {

    // Derived query — Spring generates SQL from method name
    List<User> findByEmailAndStatus(String email, UserStatus status);
    List<User> findByFirstNameContainingIgnoreCase(String name);
    Optional<User> findTopByOrderByCreatedAtDesc();
    long countByStatus(UserStatus status);
    boolean existsByEmail(String email);

    // JPQL query
    @Query("SELECT u FROM User u WHERE u.department.id = :deptId AND u.salary > :minSalary")
    List<User> findByDeptAndMinSalary(@Param("deptId") Long deptId,
                                       @Param("minSalary") BigDecimal minSalary);

    // Native SQL query
    @Query(value = "SELECT * FROM users WHERE YEAR(created_at) = :year",
           nativeQuery = true)
    List<User> findByCreationYear(@Param("year") int year);

    // Custom update (requires @Modifying + @Transactional)
    @Modifying
    @Transactional
    @Query("UPDATE User u SET u.status = :status WHERE u.lastLoginAt < :cutoff")
    int deactivateInactiveUsers(@Param("status") UserStatus status,
                                 @Param("cutoff") LocalDateTime cutoff);
}
```

---

### 🎯 How to Explain in Interview

"Spring Data JPA makes database operations incredibly simple. I extend JpaRepository to get full CRUD functionality, pagination, and sorting out of the box. The magic happens with derived queries - Spring generates SQL from method names like findByEmailAndStatus. For complex queries, I can use @Query with JPQL for database-agnostic queries, or nativeQuery=true for database-specific SQL. When I need to modify data, I add @Modifying and @Transactional. Spring Data also supports specifications for dynamic queries and QueryDSL for type-safe queries. This approach eliminates most boilerplate JDBC code while still giving me full control when needed. I can focus on business logic rather than SQL plumbing."

---

### Q5. What is lazy vs eager fetching?

```java
// FetchType.LAZY — loads association when accessed (DEFAULT for @OneToMany, @ManyToMany)
@OneToMany(fetch = FetchType.LAZY, mappedBy = "department")
private List<Employee> employees;
// employees are NOT loaded with Department — loaded only when getEmployees() is called

// FetchType.EAGER — loads association immediately (DEFAULT for @ManyToOne, @OneToOne)
@ManyToOne(fetch = FetchType.EAGER)
@JoinColumn(name = "dept_id")
private Department department;
// department IS loaded with every Employee — always

// Problem with LAZY outside of session (LazyInitializationException)
@Service
public class DeptService {
    @Transactional        // ← needed to keep session open
    public List<String> getEmployeeNames(Long deptId) {
        Department dept = deptRepo.findById(deptId).orElseThrow();
        return dept.getEmployees().stream()  // lazy load happens here — session still open
                   .map(Employee::getName)
                   .toList();
    }
}
// If @Transactional is missing → LazyInitializationException when accessing employees!
```

---

### 🎯 How to Explain in Interview

"Lazy vs eager fetching is about when JPA loads related data. With lazy loading, JPA doesn't load the association until I actually access it - perfect for collections where I might not need the data. Eager loading loads everything immediately - good for small, always-needed relationships. The default is lazy for collections and eager for single references. The catch with lazy loading is the LazyInitializationException - it happens when I try to access lazy data outside the transaction session. That's why I keep @Transactional on service methods that might access lazy-loaded data. The key is choosing the right strategy: lazy for performance, eager for convenience, and always being aware of the session boundaries. Getting this right is crucial for both performance and avoiding runtime errors."

---

### Q6. How do you manage database migrations?

```java
// Flyway — SQL migration scripts in src/main/resources/db/migration/
// Naming: V{version}__{description}.sql
// V1__create_users_table.sql
// V2__add_email_index.sql
// V3__add_orders_table.sql
```

```sql
-- V1__create_users_table.sql
CREATE TABLE users (
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50)  NOT NULL,
    email      VARCHAR(100) NOT NULL UNIQUE,
    status     VARCHAR(20)  NOT NULL DEFAULT 'ACTIVE',
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version    BIGINT       NOT NULL DEFAULT 0
);

-- V2__add_orders_table.sql
CREATE TABLE orders (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id     BIGINT         NOT NULL,
    total       DECIMAL(12, 2) NOT NULL,
    status      VARCHAR(20)    NOT NULL,
    created_at  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

```yaml
# application.yml
spring:
  flyway:
    enabled: true
    locations: classpath:db/migration
    baseline-on-migrate: true
  jpa:
    hibernate:
      ddl-auto: validate   # validate schema matches entities — don't auto-create!
```

---

### 🎯 How to Explain in Interview

"Database migrations are essential for managing schema changes in production. I use Flyway to version-control my database schema. Each migration is a SQL file named like V1__create_users_table.sql, and Flyway tracks which versions have been applied. This means I can evolve my schema safely - adding tables, columns, or indexes - without breaking existing data. The key is setting ddl-auto to validate, not update, so Hibernate doesn't auto-change the schema. Instead, Flyway handles all changes through versioned scripts. This approach gives me reproducible database setups across environments and makes deployments predictable. I can see exactly what schema changes happened and when, and roll back if needed."

---

### Q7. How do you implement optimistic and pessimistic locking?

```java
// OPTIMISTIC LOCKING — uses @Version field, no DB lock held
@Entity
public class Inventory {
    @Id private Long id;
    private int quantity;

    @Version                   // JPA bumps version on every update
    private Long version;
}

// If two threads read version=1 and both try to update:
// Thread A updates → version becomes 2
// Thread B updates → WHERE version=1 fails → OptimisticLockException thrown!

@Transactional
public void reserveStock(Long productId, int qty) {
    try {
        Inventory inv = inventoryRepo.findById(productId).orElseThrow();
        if (inv.getQuantity() < qty) throw new InsufficientStockException();
        inv.setQuantity(inv.getQuantity() - qty);
        inventoryRepo.save(inv);    // may throw OptimisticLockException
    } catch (OptimisticLockingFailureException e) {
        throw new ConcurrentUpdateException("Retry your request");
    }
}

// PESSIMISTIC LOCKING — holds DB-level lock until transaction commits
@Query("SELECT i FROM Inventory i WHERE i.id = :id")
@Lock(LockModeType.PESSIMISTIC_WRITE)   // SELECT ... FOR UPDATE
Optional<Inventory> findByIdForUpdate(@Param("id") Long id);
// Use when conflicts are frequent and data consistency is critical
```

---

### 🎯 How to Explain in Interview

"Optimistic and pessimistic locking are two strategies for handling concurrent database updates. Optimistic locking assumes conflicts are rare - I use @Version to track a version number, and JPA checks that the version hasn't changed since I read the record. If two users try to update simultaneously, the second one gets an OptimisticLockException and can retry. It's lightweight and works well for low-conflict scenarios. Pessimistic locking assumes conflicts are likely - I use @Lock with PESSIMISTIC_WRITE to lock the database row until my transaction commits. This prevents conflicts but hurts concurrency. I choose optimistic locking for most applications like user profiles, and pessimistic locking for critical operations like inventory management where conflicts are common and data consistency is crucial. Both strategies help prevent lost updates and maintain data integrity."
