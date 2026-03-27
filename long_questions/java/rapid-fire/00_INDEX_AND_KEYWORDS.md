# ☕ Java Interview Rapid-Fire Study Guide — INDEX

> **Complete Java + Spring Framework interview preparation in 7 files.**
> Each concept has a 🔑 **Memorable Keyword/Mnemonic** so you never blank out in interviews!

---

## 📚 Files in this Rapid-Fire Folder

| # | File | Topics Covered | Keywords to Remember |
|---|---|---|---|
| 1 | `01_core_java_oop_rapidfire.md` | OOP Pillars, Generics, Reflection, Enums, Modern Java | **"EPIC-CC"** |
| 2 | `02_exceptions_io_collections_rapidfire.md` | Exceptions, Streams, Collections, Functional Programming | **"ETICS"** |
| 3 | `03_concurrency_multithreading_rapidfire.md` | Threads, Locks, Executor, CompletableFuture, Atomic | **"TELF-SDA"** |
| 4 | `04_spring_boot_mvc_rapidfire.md` | Spring Boot, Auto-config, REST, DI, Bean Scopes | **"SADIE"** |
| 5 | `05_spring_data_jpa_rapidfire.md` | JPA Repositories, Queries, Transactions, Caching | **"RQELT"** |
| 6 | `06_spring_security_rapidfire.md` | JWT, OAuth2, CORS/CSRF, Method Security, BCrypt | **"AAJCP"** |
| 7 | `07_design_patterns_rapidfire.md` | Singleton, Factory, Builder, Observer, SOLID, Caching | **"CBFS-OO"** |
| 8 | `08_spring_webflux_reactive_rapidfire.md` | Reactive, Mono/Flux, Backpressure, WebClient, Schedulers, R2DBC | **"RMBFS"** |

---

## 🧠 Master Keywords Quick Reference

### File 1 — Core Java OOP

| Topic | 🔑 Keyword | Meaning |
|---|---|---|
| 4 OOP Pillars | **PEIA** | Polymorphism, Encapsulation, Inheritance, Abstraction |
| Abstract vs Interface | **SCIV** | State/Constructor/Is-A vs Capability |
| Polymorphism types | **OL-OR** | Overload=Load-time, Override=Runtime |
| Composition vs Inheritance | **HALO** | Has-A=Loose, Is-A=tightly cOupled |
| `static` keyword | **SVBC** | Static=Variable/Block/Class |
| `volatile` | **VMR** | Visibility+Memory+no-Reorder |
| Wrapper & Autoboxing | **WAN** | Wrapper-Autobox-NPE |
| Generics Type Erasure | **CER** | Compile→Erase→Runtime-raw |
| PECS Rule | **PECS** | Producer-Extends, Consumer-Super |
| Reflection | **RISC** | Runtime-Inspect-Self-Code |
| Immutable Class | **FNSD** | Final-class, No-setters, private-Final, Defensive-copy |
| Singleton Strategies | **ELSDBE** | Eager/Lazy/Sync/Double-check/Bill-Pugh/Enum |

---

### File 2 — Exceptions, IO & Collections

| Topic | 🔑 Keyword | Meaning |
|---|---|---|
| Exception Hierarchy | **T-E-C-U** | Throwable→Error/Exception→Checked/Unchecked |
| `throw` vs `throws` | **II-MS** | throw=Inside-method, throws=Method-Signature |
| Try-with-resources | **ARCA** | AutoCloseable→Resource→Close→Automatic |
| Serialization | **OBS** | Object→Byte-Stream |
| `transient` | **Skip-sensitive** | Skip sensitive fields during serialization |
| ArrayList vs LinkedList | **ARML** | Array=Random-access, LinkedList=Modifications |
| HashMap internals | **HABT** | Hash→Array→Bucket→Tree (Java 8+) |
| Comparable vs Comparator | **NE-CE** | Natural=Comparable, Custom=Comparator |
| Fail-fast vs Fail-safe | **FFFE** | FailFast=original-CME, FailSafe=copy |
| Stream pipeline | **SIT** | Source→Intermediate→Terminal |
| Functional Interfaces | **PCSFB** | Predicate/Consumer/Supplier/Function/BiFunction |
| Stream collectors | **LGMJ** | List/Grouping/Map/Joining |

---

### File 3 — Concurrency

| Topic | 🔑 Keyword | Meaning |
|---|---|---|
| Thread creation | **TRC** | Thread/Runnable/Callable |
| `start()` vs `run()` | **SNR** | Start=New-thread, Run=same-thread |
| `wait()` vs `sleep()` | **WRL-SK** | Wait=Releases-Lock, Sleep=Keeps-lock |
| Daemon thread | **DJVM** | JVM exits when only daemons left |
| Thread pools | **FCSS** | Fixed/Cached/Single/Scheduled |
| `submit()` vs `execute()` | **SFE** | Submit=Future, Execute=fire-forget |
| CompletableFuture | **CFP** | CompletableFuture=Promise-like-Chaining |
| Virtual threads | **VJML** | Virtual=JVM-managed-Lightweight (Java 21) |
| Deadlock | **MHCN-4** | 4 conditions: Mutual-exclusion/Hold+wait/Circular/No-preemption |
| `volatile` | **VMR** | Visibility+Memory+no-Reorder |
| AtomicInteger | **ACas** | Atomic=CompareAndSwap hardware |
| CountDownLatch vs CyclicBarrier | **CDO-CRR** | CountDown=One-time, Cyclic=Reusable |

---

### File 4 — Spring Boot & MVC

| Topic | 🔑 Keyword | Meaning |
|---|---|---|
| `@SpringBootApplication` | **CEC** | Configuration+EnableAutoConfig+ComponentScan |
| Auto-configuration | **CCA** | Classpath→Conditional→AutoConfig |
| Spring Starters | **SBB** | Starter=Bundle-of-Blessed-dependencies |
| Config priority | **CESO** | Command-line>Env>Specific-profile>base |
| Spring Profiles | **ADPT** | Activate Different Properties per Target |
| Spring Actuator | **HMEI** | Health/Metrics/Env/Info endpoints |
| MVC request flow | **DHCVR** | DispatcherServlet→Handler→Controller→View→Response |
| `@Controller` vs `@RestController` | **VJ** | View-return vs JSON-return |
| Filters vs Interceptors | **FS-SI** | Filter=Servlet, Interceptor=Spring-context |
| IoC vs DI | **IoC-invert, DI-inject** | Framework creates and injects objects |
| DI types | **CFS** | Constructor/Field/Setter |
| Bean Scopes | **SPRRS** | Singleton/Prototype/Request/Response/Session |
| `@Bean` vs `@Component` | **BC-TM** | @Bean=third-party-manual, @Component=your-class-auto |

---

### File 5 — Spring Data JPA

| Topic | 🔑 Keyword | Meaning |
|---|---|---|
| Spring Data JPA | **PAR** | Proxy-Auto-Repository |
| Repository hierarchy | **CPSJ** | Crud→PagingSort→Jpa |
| Derived query methods | **FBK** | Find+By+Keywords |
| `@Query` annotation | **JNP** | JPQL/Native/Params |
| Entity states | **TPDR** | Transient→Persistent→Detached→Removed |
| Dirty checking | **SDU** | Snapshot-Diff→auto-UPDATE |
| Lazy vs Eager | **LD-FN** | Lazy=Demand, Eager=Fetch-Now |
| N+1 problem | **NP-JF** | N+1→Join-Fetch fix |
| `@Transactional` | **ACID-AOP** | ACID via AOP proxy |
| Propagation | **REQ-NEW** | Required=join-existing, RequiresNew=always-new |
| Locking | **OV-PLock** | Optimistic=Version, Pessimistic=Lock-database |
| Hibernate Cache | **L1S-L2G** | L1=Session-local, L2=Global |
| Pagination | **PCR** | Pageable→Content→Result |

---

### File 6 — Spring Security

| Topic | 🔑 Keyword | Meaning |
|---|---|---|
| Auth vs AuthZ | **WHO-WHAT** | AuthN=WHO, AuthZ=WHAT |
| Filter chain | **DFC** | DelegatingFilterProxy→FilterChainProxy→SecurityFilterChain |
| SecurityContextHolder | **CTH** | Context=Auth-storage, Holder=ThreadLocal-accessor |
| `UserDetailsService` | **LUN** | Load-User-byUsername |
| JWT structure | **HPS** | Header.Payload.Signature |
| JWT flow | **LVFS** | Login→Validate→Filter→SecurityContext |
| JWT vs Session | **JSS** | JWT=Stateless, Session=Server-state |
| CSRF | **CSRF-C-disable** | Cookie-based attack, disable for stateless JWT |
| CORS | **CORS-Allow** | Configure to allow trusted origins |
| Method security | **SPA** | Secured/PreAuthorize/Annotations |
| BCrypt | **BSW** | BCrypt=Salted+Work-factor |
| OAuth2 | **UCAST** | User→Client→AuthServer→Token→ResourceServer |
| Role Hierarchy | **RHIA** | Inherit-below-roles-Automatically |

---

### File 7 — Design Patterns

| Topic | 🔑 Keyword | Meaning |
|---|---|---|
| Singleton (all strategies) | **ELSDBE** | Eager/Lazy/Sync/Double-check/Bill-Pugh/Enum |
| Factory | **FDC** | Factory=Decouple-Creation |
| Builder | **BSC** | Builder=Step-by-step-Construction |
| Adapter | **AWC** | Adapter=Wrapper to make incompatible Compatible |
| Decorator | **DWS** | Decorator=Wrapper-Stack (add behavior dynamically) |
| Proxy | **PCC** | Proxy=Control-access (Spring AOP, Hibernate lazy) |
| Facade | **FS** | Facade=Simple-interface to complex-system |
| Observer | **PNS** | Publisher→Notify→Subscribers |
| Strategy | **SR** | Strategy=Runtime-behavior-Switch |
| Template Method | **ST-OH** | Skeleton-Template, child Overrides-Hook-steps |
| SOLID | **SOLID** | Single/Open/Liskov/Interface/Dependency |
| CAP Theorem | **CAP-2of3** | Pick 2: Consistency/Availability/Partition |

---

### File 8 — Spring WebFlux

| Topic | 🔑 Keyword | Meaning |
|---|---|---|
| Reactive Programming | **SANDP** | Streams/Async/Non-blocking/Data-push/Pressure |
| Reactive Streams interfaces | **PPSS** | Publisher/Processor/Subscriber/Subscription |
| Mono vs Flux | **M01-FN** | Mono=0-or-1, Flux=0-to-N |
| Backpressure | **CPOF** | Consumer-tells-Producer-Only-send-N |
| WebFlux vs MVC | **BNE-IO** | Blocking/New-thread=MVC, Event-loop/IO=WebFlux |
| flatMap vs concatMap | **FC-OI** | FlatMap=Concurrent/Out-of-order, ConcatMap=sequential/In-order |
| WebClient | **WNRB** | WebClient=Non-blocking-Reactive-Builder |
| Error handling | **ORM** | onErrorReturn/onErrorResume/onErrorMap |
| Schedulers | **PBSE** | publishOn/subscribeOn→Bounded/Parallel/Single |
| Cold vs Hot streams | **CS-HS** | Cold=per-Subscriber, Hot=Shared-broadcast |
| StepVerifier | **SVE** | StepVerifier=Expect-values-then-Verify |
| R2DBC | **RJNB** | Reactive-JDBC-Non-Blocking |
| When to use WebFlux | **HSIM** | High-concurrency/Streaming/IO-bound/Microservices |

---

## 🎯 High-Priority Interview Topics (Must Know)

### Absolute Must-Know (Asked in almost every interview):

1. ✅ **OOP Pillars** — PEIA mnemonic
2. ✅ **Abstract Class vs Interface** — table differences
3. ✅ **`HashMap` internals** — HABT (Hash→Array→Bucket→Tree)
4. ✅ **`ArrayList` vs `LinkedList`** — ARML (read vs write)
5. ✅ **`synchronized` vs `volatile`** — different guarantees
6. ✅ **`Comparable` vs `Comparator`** — NE-CE
7. ✅ **Java 8 Streams** — SIT (Source→Intermediate→Terminal)
8. ✅ **`@SpringBootApplication`** — CEC (Configuration+EnableAutoConfig+ComponentScan)
9. ✅ **Spring DI** — Constructor preferred over Field injection
10. ✅ **`@Transactional`** — ACID via AOP
11. ✅ **JWT flow** — LVFS (Login→Validate→Filter→SecurityContext)
12. ✅ **Singleton Pattern** — ELSDBE, best = Enum
13. ✅ **N+1 problem** — NP-JF, fix with JOIN FETCH
14. ✅ **Deadlock** — MHCN-4 (4 conditions)
15. ✅ **Exception hierarchy** — T-E-C-U

---

## 📖 Study Order Recommendation

**Day 1:** `01_core_java_oop_rapidfire.md` — Foundation
**Day 2:** `02_exceptions_io_collections_rapidfire.md` — Data structures
**Day 3:** `03_concurrency_multithreading_rapidfire.md` — Threading (hard, give more time)
**Day 4:** `04_spring_boot_mvc_rapidfire.md` — Framework basics
**Day 5:** `05_spring_data_jpa_rapidfire.md` — Database layer
**Day 6:** `06_spring_security_rapidfire.md` — Security
**Day 7:** `07_design_patterns_rapidfire.md` — Architecture
**Day 8:** `08_spring_webflux_reactive_rapidfire.md` — Reactive Programming
**Day 9:** Revise all keywords + mock interview

---

*Total questions covered: ~220+ across all files*
