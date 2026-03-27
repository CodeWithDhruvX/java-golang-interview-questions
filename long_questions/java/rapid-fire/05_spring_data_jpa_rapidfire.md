# 🗄️ Spring Data JPA & Hibernate (Rapid-Fire)

> 🔑 **Master Keyword:** **"RQELT"** → Repositories, Queries, Entity-lifecycle, Lazy/Eager, Transactions

---

## 📦 Section 1: JPA / Spring Data Fundamentals

### Q1: What is Spring Data JPA?
🔑 **Keyword: "PAR"** → Proxy-Auto-Repository (Spring generates implementation)

- Abstraction on top of JPA (Hibernate as default provider)
- Just extend `JpaRepository<Entity, ID>` → get CRUD for free
- No need to write DAO implementation — Spring generates it via **Proxy**

```java
// All CRUD just by extending!
public interface UserRepository extends JpaRepository<User, Long> {
    List<User> findByEmail(String email);  // derived query method
}
```

---

### Q2: Repository Hierarchy?
🔑 **Keyword: "CPSJ"** → CrudRepo → PagingSorting → JpaRepo

```
Repository (marker)
└── CrudRepository (save/find/delete)
    └── PagingAndSortingRepository (pagination + sorting)
        └── JpaRepository (batch ops + flush + deleteInBatch)
```

---

### Q3: Derived Query Methods?
🔑 **Keyword: "FBK"** → Find+By+Keywords → auto-SQL

```java
List<User> findByEmail(String email);
List<User> findByAgeGreaterThan(int age);
Optional<User> findByUsernameAndActiveTrue(String username);
long countByDepartment(String dept);
boolean existsByEmail(String email);
void deleteByEmail(String email);
List<User> findByNameOrderByCreatedAtDesc(String name);
List<User> findTop5ByOrderBySalaryDesc(); // top N results
```

---

### Q4: `@Query` annotation?
🔑 **Keyword: "JNP"** → JPQL/Native-query/Params

```java
// JPQL query (entity names, not table names)
@Query("SELECT u FROM User u WHERE u.status = :status")
List<User> findByStatus(@Param("status") String status);

// Native SQL query
@Query(value = "SELECT * FROM users WHERE age > ?1", nativeQuery = true)
List<User> findByAgeNative(int age);

// Update/Delete — needs @Modifying + @Transactional
@Modifying
@Transactional
@Query("UPDATE User u SET u.active = false WHERE u.lastLogin < :date")
int deactivateOldUsers(@Param("date") LocalDate date);
```

---

### Q5: `findById()` vs `getReferenceById()`?
🔑 **Keyword: "ELP"** → Eager=findById, Lazy-Proxy=getReferenceById

| Feature | `findById()` | `getReferenceById()` (replaces `getOne`) |
|---|---|---|
| DB hit | Immediate | Deferred (proxy, lazy) |
| Returns | `Optional<T>` | Proxy object |
| Not found | Returns `Optional.empty()` | `EntityNotFoundException` at access time |
| Best for | When you need entity now | Setting FK without loading entity |

---

## 🔁 Section 2: Entity Lifecycle

### Q6: JPA Entity States?
🔑 **Keyword: "TPDR"** → Transient→Persistent→Detached→Removed

```
new User() → [TRANSIENT — no ID, no JPA tracking]
    → entityManager.persist() / repo.save()
    → [PERSISTENT — JPA tracks, changes auto-synced to DB]
    → session closes / entityManager.detach()
    → [DETACHED — has ID, but changes not tracked]
    → entityManager.merge() → back to PERSISTENT
    → entityManager.remove()
    → [REMOVED — scheduled for deletion on commit]
```

---

### Q7: What is Dirty Checking?
🔑 **Keyword: "SDU"** → Snapshot-Diff→auto-UPDATE

- When entity is **persistent**, Hibernate takes a snapshot of its state
- On transaction commit → compares current state with snapshot
- If changed → auto-generates `UPDATE` SQL
- No need to call `save()` for already-managed entities!

```java
@Transactional
public void updateAge(Long id) {
    User user = userRepository.findById(id).get(); // persistent
    user.setAge(25); // no need to call save()!
    // Hibernate auto-updates on commit (dirty checking)
}
```

---

## ⚡ Section 3: Fetch Strategies

### Q8: Lazy vs Eager Loading?
🔑 **Keyword: "LD-FN"** → Lazy=Demand, Fetch when Needed; Eager=Fetch-Now

| | Lazy | Eager |
|---|---|---|
| Default for | `@OneToMany`, `@ManyToMany` | `@ManyToOne`, `@OneToOne` |
| DB query | On access | Immediately with parent |
| Performance | Better (no extra data) | Can be wasteful |

```java
@OneToMany(fetch = FetchType.LAZY) // explicitly set
private List<Order> orders;

@ManyToOne(fetch = FetchType.EAGER) // explicitly set
private User user;
```

---

### Q9: N+1 Problem?
🔑 **Keyword: "NP-JF"** → N+1→Join-Fetch fix

**Problem:** Fetch 100 users → 1 query. Load their orders → 100 extra queries = **101 queries!**

```java
// BAD — Causes N+1
@OneToMany(fetch = FetchType.LAZY)
private List<Order> orders;

// FIX 1: JOIN FETCH in JPQL
@Query("SELECT u FROM User u JOIN FETCH u.orders")
List<User> findAllWithOrders();

// FIX 2: @EntityGraph
@EntityGraph(attributePaths = {"orders"})
List<User> findAll();
```

---

## 🔄 Section 4: Transactions

### Q10: `@Transactional` explained?
🔑 **Keyword: "ACID-AOP"** → ACID guarantees via AOP proxy

```java
@Service
public class OrderService {
    @Transactional  // Spring wraps with begin + commit/rollback
    public void placeOrder(Order order) {
        orderRepo.save(order);
        paymentService.charge(order); // if this fails → rollback order save!
    }
}
```

- Default: rolls back on **RuntimeException** (unchecked)
- Override: `@Transactional(rollbackFor = Exception.class)` for checked exceptions

---

### Q11: Transaction Propagation?
🔑 **Keyword: "REQ-NEW"** → Required=join-existing, Requires_New=always-new

| Propagation | Behavior |
|---|---|
| `REQUIRED` (default) | Join existing or create new |
| `REQUIRES_NEW` | Always create new (suspends existing) |
| `SUPPORTS` | Join if exists, else no transaction |
| `MANDATORY` | Must have existing, else exception |
| `NEVER` | Must NOT have transaction |

---

### Q12: Transaction Isolation Levels?
🔑 **Keyword: "DRRS"** → Dirty-Read/Read-Committed/Repeatable/Serializable

| Level | Dirty Read | Non-Repeatable Read | Phantom Read |
|---|---|---|---|
| `READ_UNCOMMITTED` | Yes | Yes | Yes |
| `READ_COMMITTED` (Default) | No | Yes | Yes |
| `REPEATABLE_READ` | No | No | Yes |
| `SERIALIZABLE` | No | No | No |

High isolation → more consistency but less performance (more locking).

---

## 🔗 Section 5: Relationships & Cascading

### Q13: JPA Relationship Annotations?
🔑 **Keyword: "OMMO"** → OneToMany/ManyToOne/ManyToMany/OneToOne

```java
@Entity
public class Department {
    @OneToMany(mappedBy = "department", cascade = CascadeType.ALL)
    private List<Employee> employees;
}

@Entity
public class Employee {
    @ManyToOne
    @JoinColumn(name = "dept_id")
    private Department department;
}
```

---

### Q14: Cascade Types?
🔑 **Keyword: "PAMRRD"** → Persist/All/Merge/Remove/Refresh/Detach

| CascadeType | Meaning |
|---|---|
| `PERSIST` | Save child when saving parent |
| `MERGE` | Update child when merging parent |
| `REMOVE` | Delete child when deleting parent |
| `ALL` | All above combined |
| `REFRESH` | Reload child data with parent |

```java
@OneToMany(cascade = CascadeType.ALL, orphanRemoval = true)
private List<Address> addresses;
```
`orphanRemoval = true` → delete child if removed from parent collection.

---

### Q15: Optimistic vs Pessimistic Locking?
🔑 **Keyword: "OV-PLock"** → Optimistic=Version, Pessimistic=Lock-database-row

```java
// Optimistic Locking — version column
@Entity
public class Account {
    @Version
    private Long version; // auto-incremented, OptimisticLockException on conflict
}

// Pessimistic Locking — database-level row lock
@Lock(LockModeType.PESSIMISTIC_WRITE)
Optional<Account> findById(Long id); // SELECT ... FOR UPDATE
```

---

## 📊 Section 6: Advanced JPA

### Q16: JPA Projections?
🔑 **Keyword: "ICD"** → Interface/Class/DTO projections

```java
// Interface-based projection — only specific fields
public interface UserSummary {
    String getName();
    String getEmail();
}
List<UserSummary> findBy(); // Spring generates query for only those fields

// Class-based DTO
@Query("SELECT new com.example.UserDTO(u.name, u.email) FROM User u")
List<UserDTO> findAllDTOs();
```

---

### Q17: Hibernate L1 vs L2 Cache?
🔑 **Keyword: "L1S-L2G"** → L1=Session-local, L2=Global-across-sessions

| Cache | Scope | Enabled by default? | Provider |
|---|---|---|---|
| L1 (Session cache) | Within one session/transaction | ✅ YES | Built-in |
| L2 (Application cache) | Across sessions | ❌ NO (configure) | EhCache, Redis |

```java
// Enable L2 for an entity
@Entity
@Cache(usage = CacheConcurrencyStrategy.READ_WRITE)
public class Country { ... }
```

---

### Q18: Spring Data Auditing?
🔑 **Keyword: "ECAL"** → EnableAuditing+CreatedDate+AuditListener

```java
// Configuration
@EnableJpaAuditing
@Configuration
public class JpaConfig { }

// Entity
@EntityListeners(AuditingEntityListener.class)
@Entity
public class Product {
    @CreatedDate
    private LocalDateTime createdAt;

    @LastModifiedDate
    private LocalDateTime updatedAt;

    @CreatedBy
    private String createdBy;
}
```

---

### Q19: Pagination in Spring Data?
🔑 **Keyword: "PCR"** → Pageable→Content→Result

```java
// Repository
Page<User> findByActive(boolean active, Pageable pageable);

// Service
Pageable pageable = PageRequest.of(0, 10, Sort.by("name").ascending());
Page<User> page = userRepository.findByActive(true, pageable);

// Access result
List<User> users = page.getContent();
long total = page.getTotalElements();
int totalPages = page.getTotalPages();
boolean hasMore = page.hasNext();
```

---

*End of File — Spring Data JPA & Hibernate*
