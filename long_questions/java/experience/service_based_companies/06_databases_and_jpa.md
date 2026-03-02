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
