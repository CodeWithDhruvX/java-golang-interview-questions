# ⚡ Spring WebFlux & Reactive Programming (Rapid-Fire)

> 🔑 **Master Keyword:** **"RMBFS"** → Reactive/Mono-Flux/Backpressure/Functional/Scheduler

---

## 🔁 Section 1: Reactive Programming Fundamentals

### Q1: What is Reactive Programming?
🔑 **Keyword: "SANDP"** → Streams/Async/Non-blocking/Data-push/backPressure

- Programming paradigm around **data streams** and propagation of change
- Instead of **pulling** data → subscribe to streams that **push** data when ready
- Key benefits: non-blocking, backpressure support, high scalability

> Think of it like a **notification system**: you don't keep calling to check for news, you **subscribe** and get notified when news arrives.

---

### Q2: Reactive Streams Specification — 4 Interfaces?
🔑 **Keyword: "PPSS"** → Publisher/Processor/Subscriber/Subscription

```
Publisher  →  emits data  →  Subscriber
                               ↑
                          Subscription (controls flow via request(n))
Processor  =  Publisher + Subscriber (transforms in the middle)
```

| Interface | Role |
|---|---|
| `Publisher<T>` | Produces data, accepts subscriber |
| `Subscriber<T>` | Consumes data (`onNext`, `onError`, `onComplete`) |
| `Subscription` | Contract between publisher/subscriber (`request(n)`, `cancel()`) |
| `Processor<T,R>` | Both publisher and subscriber (transforms) |

---

### Q3: `Mono<T>` vs `Flux<T>`?
🔑 **Keyword: "M01-FN"** → Mono=0-or-1, Flux=0-to-N

| Type | Emits | Analogy | Use-case |
|---|---|---|---|
| `Mono<T>` | 0 or 1 element | Async `Optional<T>` | findById, single HTTP call |
| `Flux<T>` | 0 to N elements | Async `List<T>` | findAll, event stream |

```java
// Mono — single async result
Mono<User> user = userRepository.findById(1L);

// Flux — multiple async results
Flux<User> users = userRepository.findAll();

// Creation shortcuts
Mono<String> m1 = Mono.just("Hello");
Mono<String> m2 = Mono.empty();                     // 0 elements
Mono<String> m3 = Mono.error(new RuntimeException());

Flux<Integer> f1 = Flux.just(1, 2, 3, 4, 5);
Flux<Integer> f2 = Flux.range(1, 10);               // 1 to 10
Flux<String> f3 = Flux.fromIterable(List.of("a", "b"));
```

---

### Q4: What is Backpressure?
🔑 **Keyword: "CPOF"** → Consumer-tells-Producer-Only-send-N (flow control)

- Consumer signals to producer: **"send me only N items"** — no overwhelming!
- Subscriber calls `subscription.request(n)` to pull items at its own pace
- Prevents fast producer from overwhelming slow consumer → avoids `OutOfMemoryError`

```
Producer (fast) → [ Backpressure: request(10) ] → Consumer (slow)
                         ↑ consumer controls the rate!
```

---

## 🌐 Section 2: Spring WebFlux vs Spring MVC

### Q5: WebFlux vs Spring MVC?
🔑 **Keyword: "BNE-IO"** → Blocking/New-thread=MVC, Event-loop/IO-bound=WebFlux

| Feature | Spring MVC | Spring WebFlux |
|---|---|---|
| IO Model | Blocking (synchronous) | Non-blocking (async) |
| Threading | 1 thread per request | Event loop (few threads, many requests) |
| Server | Tomcat (servlet) | **Netty** (default), Undertow |
| Return types | `String`, `Object`, `ResponseEntity` | `Mono<T>`, `Flux<T>` |
| Best for | CPU-bound, traditional CRUD | High-concurrency I/O-bound, streaming |
| DB access | JDBC (blocking) | R2DBC (reactive, non-blocking) |

---

### Q6: Thread Model — Event Loop vs Thread-per-Request?
🔑 **Keyword: "ETMH"** → Event-loop=Tiny-threads-Many-Handles

```
Thread-per-Request (MVC):
Request1 → Thread1 (BLOCKED waiting for DB)
Request2 → Thread2 (BLOCKED waiting for network)
Request3 → Thread3 (BLOCKED...)
→ 1000 requests = 1000 threads 💥 (thread starvation)

Event Loop (WebFlux):
Request1 →  Event Loop Thread → dispatches → async DB call
Request2 →  Event Loop Thread → dispatches → async HTTP call
Request3 →  Event Loop Thread → dispatches → ...
→ 1000 requests = still ~8 threads ✅ (non-blocking callbacks)
```

---

## 📝 Section 3: Programming Models

### Q7: Annotation-based Controller (same as MVC!)?
🔑 **Keyword: "SAMR"** → Same-Annotations, Mono/Flux-Return

```java
@RestController
@RequestMapping("/users")
public class UserController {

    @GetMapping("/{id}")
    public Mono<User> getUser(@PathVariable Long id) {           // Mono for single
        return userRepository.findById(id);
    }

    @GetMapping
    public Flux<User> getAllUsers() {                            // Flux for many
        return userRepository.findAll();
    }

    @PostMapping
    public Mono<ResponseEntity<User>> createUser(@RequestBody User user) {
        return userRepository.save(user)
            .map(saved -> ResponseEntity.status(201).body(saved));
    }
}
```

---

### Q8: Functional Endpoints (RouterFunction)?
🔑 **Keyword: "RHB"** → RouterFunction+HandlerFunction=no-reflection-based

```java
// Handler (like controller method body)
@Component
public class UserHandler {
    private final UserRepository repo;
    
    public Mono<ServerResponse> findAll(ServerRequest request) {
        return ServerResponse.ok()
            .contentType(MediaType.APPLICATION_JSON)
            .body(repo.findAll(), User.class);
    }

    public Mono<ServerResponse> findById(ServerRequest request) {
        Long id = Long.parseLong(request.pathVariable("id"));
        return repo.findById(id)
            .flatMap(user -> ServerResponse.ok().bodyValue(user))
            .switchIfEmpty(ServerResponse.notFound().build());
    }
}

// Router (like @RequestMapping)
@Configuration
public class UserRouter {
    @Bean
    public RouterFunction<ServerResponse> routes(UserHandler handler) {
        return RouterFunctions
            .route(GET("/users"), handler::findAll)
            .andRoute(GET("/users/{id}"), handler::findById);
    }
}
```

---

## 🔗 Section 4: Key Operators

### Q9: Essential Reactor Operators?
🔑 **Keyword: "MFFZ"** → Map/FlatMap/Filter/Zip core operators

```java
Flux<String> names = Flux.just("alice", "bob", "charlie");

// map — synchronous transformation (1-to-1)
Flux<String> upper = names.map(String::toUpperCase);

// flatMap — async transformation (1-to-Mono/Flux), flattened
Flux<User> users = userIds.flatMap(id -> userRepo.findById(id)); // DB call per id

// filter — keep elements matching predicate
Flux<String> longNames = names.filter(n -> n.length() > 3);

// take / skip
Flux<String> first2 = names.take(2);     // take first 2
Flux<String> skip1 = names.skip(1);      // skip first 1

// zip — combine two streams element-by-element
Flux<String> combined = Flux.zip(namesFlux, agesFlux, (n, a) -> n + ":" + a);

// merge — combine streams as they arrive (no order)
Flux<String> merged = Flux.merge(flux1, flux2);

// collectList — aggregate Flux into Mono<List>
Mono<List<String>> list = names.collectList();
```

---

### Q10: `flatMap` vs `concatMap`?
🔑 **Keyword: "FC-OI"** → FlatMap=Concurrent/Out-of-order, ConcatMap=sequential/In-order

| Operator | Execution | Order |
|---|---|---|
| `flatMap` | **Concurrent** (all at once) | Not guaranteed |
| `concatMap` | **Sequential** (one after another) | Maintained |
| `switchMap` | Cancels previous, switches to new | Latest only |

```java
// flatMap — concurrent DB calls (faster but unordered)
userIds.flatMap(id -> repo.findById(id))    // all IDs queried simultaneously

// concatMap — sequential (ordered but slower)
userIds.concatMap(id -> repo.findById(id))  // wait for each before next
```

---

## 🌐 Section 5: WebClient

### Q11: `WebClient` — Replaces `RestTemplate`?
🔑 **Keyword: "WNRB"** → WebClient=Non-blocking-Reactive-Builder

```java
// Create WebClient
WebClient client = WebClient.builder()
    .baseUrl("https://api.example.com")
    .defaultHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
    .build();

// GET — returns Mono
Mono<User> user = client.get()
    .uri("/users/{id}", 1L)
    .retrieve()
    .bodyToMono(User.class);

// GET — returns Flux (streaming)
Flux<User> users = client.get()
    .uri("/users")
    .retrieve()
    .bodyToFlux(User.class);

// POST
Mono<User> created = client.post()
    .uri("/users")
    .bodyValue(new User("Alice"))
    .retrieve()
    .bodyToMono(User.class);

// Error handling
Mono<User> withFallback = client.get()
    .uri("/users/99")
    .retrieve()
    .onStatus(HttpStatus::is4xxClientError, resp -> Mono.error(new NotFoundException()))
    .bodyToMono(User.class);
```

---

### Q12: `WebClient` vs `RestTemplate`?
🔑 **Keyword: "WNBS"** → WebClient=Non-Blocking-Streams, RestTemplate=Synchronous-legacy

| Feature | `RestTemplate` | `WebClient` |
|---|---|---|
| Blocking | ✅ Blocks thread | ❌ Non-blocking |
| Reactive | ❌ | ✅ Mono/Flux |
| Streaming | ❌ | ✅ |
| Status | **Deprecated** (Spring 5+) | **Recommended** |

---

## ❌ Section 6: Error Handling

### Q13: Error Handling Operators?
🔑 **Keyword: "ORM"** → onErrorReturn/onErrorResume/onErrorMap

```java
Mono<User> user = userRepo.findById(id)
    // Return default value on error
    .onErrorReturn(new User("default"))

    // Switch to another Mono/Flux on error
    .onErrorResume(e -> Mono.just(cachedUser))

    // Transform exception type
    .onErrorMap(DatabaseException.class, e -> new ServiceException("DB error: " + e.getMessage()))

    // Side-effect for logging (doesn't change stream)
    .doOnError(e -> log.error("Error fetching user", e));
```

> Traditional `try-catch` does **not** work for async reactive chains — always use these operators!

---

## 🧵 Section 7: Schedulers

### Q14: Schedulers — Thread Control?
🔑 **Keyword: "PBSE"** → publishOn/subscribeOn→Schedulers(Bounded/Parallel/Single/Elastic)

```java
// subscribeOn — controls where the source subscribes (upstream)
Flux.fromIterable(largeList)
    .subscribeOn(Schedulers.boundedElastic()) // run source on elastic thread pool

// publishOn — controls where SUBSEQUENT operators run
Flux.just(1, 2, 3)
    .publishOn(Schedulers.parallel())        // switch to parallel thread pool here
    .map(this::cpuIntensiveWork)
```

| Scheduler | Thread Pool | Best for |
|---|---|---|
| `Schedulers.parallel()` | Fixed (CPU count) | CPU-intensive tasks |
| `Schedulers.boundedElastic()` | Bounded elastic pool | **Blocking I/O** (JDBC, file IO) |
| `Schedulers.single()` | Single thread | Sequential operations |
| `Schedulers.immediate()` | Current thread | No switching |

---

## 🔥 Section 8: Hot vs Cold Streams

### Q15: Cold vs Hot Streams?
🔑 **Keyword: "CS-HS"** → Cold=per-Subscriber-fresh, Hot=Shared-broadcast

| | Cold Stream | Hot Stream |
|---|---|---|
| Data generation | On subscription | Regardless of subscribers |
| Each subscriber | Gets fresh/independent stream | Shares same stream (may miss data) |
| Examples | `Flux.just()`, DB query | Mouse events, stock ticker, Kafka topic |
| Late subscriber | Sees full data | Misses already-emitted data |

```java
// Cold — each subscriber gets fresh DB query
Flux<User> cold = userRepository.findAll(); // new query per subscription

// Hot — shared, live source
Sinks.Many<String> sink = Sinks.many().multicast().onBackpressureBuffer();
Flux<String> hot = sink.asFlux();           // multiple subscribers share this
sink.tryEmitNext("event1");                 // emitted even if no subscriber yet
```

---

## 🧪 Section 9: Testing

### Q16: `StepVerifier` — Testing Reactive Streams?
🔑 **Keyword: "SVE"** → StepVerifier=Expect-values-then-Verify

```java
// Basic usage
StepVerifier.create(Flux.just("a", "b", "c"))
    .expectNext("a")
    .expectNext("b")
    .expectNext("c")
    .verifyComplete();

// Error assertions
StepVerifier.create(Mono.error(new RuntimeException("fail")))
    .expectErrorMessage("fail")
    .verify();

// Testing with time (virtual time — no actual wait!)
StepVerifier.withVirtualTime(() -> Flux.interval(Duration.ofSeconds(1)).take(3))
    .expectSubscription()
    .thenAwait(Duration.ofSeconds(3))  // virtually advance time
    .expectNextCount(3)
    .verifyComplete();
```

---

## 🗃️ Section 10: R2DBC (Reactive Database)

### Q17: What is R2DBC?
🔑 **Keyword: "RJNB"** → R2DBC=Reactive-JDBC-Non-Blocking

- **Reactive Relational Database Connectivity** — non-blocking counterpart to JDBC
- JDBC is blocking → defeats the purpose of WebFlux
- R2DBC maintains end-to-end reactive pipeline (no thread blocking anywhere!)

```java
// Spring Data R2DBC repository (same as JPA repo, but reactive!)
public interface UserRepository extends ReactiveCrudRepository<User, Long> {
    Flux<User> findByDepartment(String dept);   // returns Flux
    Mono<User> findByEmail(String email);       // returns Mono
}

// Usage in service
Flux<User> users = userRepository.findByDepartment("Engineering");
```

| Feature | JDBC | R2DBC |
|---|---|---|
| Blocking | ✅ Blocks thread | ❌ Non-blocking |
| Used with | Spring MVC | Spring WebFlux |
| Returns | `List<T>` | `Flux<T>` / `Mono<T>` |

---

## 🔑 Section 11: Quick-Reference Summary

### Q18: When to choose WebFlux over MVC?

🔑 **Keyword: "HSIM"** → High-concurrency/Streaming/IO-bound/Microservices

**Choose WebFlux when:**
- Need to handle **thousands of concurrent** connections (APIs, gateways)
- Working with **streaming data** (Server-Sent Events, WebSocket, Kafka)
- Application is **I/O-bound** (calling many external services)
- Building **reactive microservices** end-to-end (R2DBC + WebFlux + WebClient)

**Stick with MVC when:**
- Team is not familiar with reactive programming
- Application is **CPU-bound** or simple CRUD
- Using **JDBC** (blocking DB) — no benefit from WebFlux
- Need simpler debugging and testing

---

### Q19: Key Annotations & Classes Quick Reference?

| What | Class/Annotation |
|---|---|
| Reactive controller | `@RestController` + return `Mono`/`Flux` |
| Functional router | `RouterFunction`, `HandlerFunction` |
| HTTP client | `WebClient` |
| Test utility | `StepVerifier` |
| Default server | `Netty` |
| Reactive DB API | `R2DBC` + `ReactiveCrudRepository` |
| Cold stream creation | `Flux.just()`, `Flux.fromIterable()` |
| Hot stream | `Sinks.many().multicast()` |
| Thread switching | `publishOn()`, `subscribeOn()` |
| Error fallback | `onErrorReturn()`, `onErrorResume()` |

---

*End of File — Spring WebFlux & Reactive Programming*
