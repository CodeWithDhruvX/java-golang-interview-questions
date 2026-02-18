# Spring WebFlux & Reactive - Interview Answers

> 🎯 **Focus:** This is a niche but high-value skill. Focus on the "Why"—scaling with fewer threads.

### 1. Spring MVC vs Spring WebFlux?
"**Spring MVC** is blocking (Servlet API). One request = One Thread. Ideally, if you have 200 threads, you can handle 200 concurrent requests. It uses Tomcat.

**Spring WebFlux** is non-blocking (Reactive Streams). It uses an Event Loop (Netty). A small number of threads (equal to CPU cores) can handle thousands of concurrent requests by never waiting for I/O.
I choose WebFlux for high-throughput streaming apps or gateways; for standard CRUD, MVC is easier to debug."

---

### 2. Mono vs Flux?
"These are the two core publishers in Project Reactor.
**Mono<T>**: Emits 0 or 1 item. Like a standard `Future` or `Optional`. Use it for `findById`.
**Flux<T>**: Emits 0 to N items. It’s a stream. Use it for `findAll`.

The key is that nothing happens until you **subscribe**. They are lazy descriptions of a data pipeline."

---

### 3. What is Backpressure?
"It’s a flow control mechanism.
In a traditional push model, a fast producer can overwhelm a slow consumer (causing OutOfMemory).
In Reactive Streams, the Consumer (Subscriber) tells the Producer *exactly* how much data it can handle: `request(10)`.
The Producer respects this and only pushes what was requested. It stabilizes the system under load."

---

### 4. `flatMap` vs `map` in Reactive?
"**map** transforms items synchronously. Input `String` -> Output `Integer`. One-to-one.

**flatMap** transforms items asynchronously. Input `String` -> Output `Mono<Integer>`.
It flattens the inner publisher.
If I need to make a DB call or API call for each item in a stream, I must use `flatMap`. If I used `map`, I’d end up with a `Flux<Mono<User>>`, which is awkward."

---

### 5. `WebClient` vs `RestTemplate`?
"**RestTemplate** is deprecated (in maintenance mode). It’s blocking.
**WebClient** is the modern reactive client.

The cool thing is `WebClient` can be used even in synchronous MVC apps. It supports both sync (`.block()`) and async calls.
It’s fluent, functional, and efficient. I use it for all new HTTP integrations."

---

### 6. Cold vs Hot Publishers?
"**Cold Publisher**: restarts the data stream for every subscriber. Like a Netflix movie—if I start watching, it starts from the beginning. If you start watching, it starts from the beginning for you too. Most calls (Mono/Flux) are cold by default.

**Hot Publisher**: broadcasts data to all subscribers in real-time. Like a Live TV channel. If you create a Hot Flux, late subscribers miss the data that was already emitted."

---

### 7. Reactive Error Handling?
"You can't use `try-catch` blocks because the code runs on different threads at different times.
Instead, we use operators:
`.onErrorReturn(defaultValue)`: Fallback to a default value.
`.onErrorResume(e -> fallbackFlux)`: Switch to a different path logic.
`.retry(3)`: Automatically retry the operation 3 times before failing."

---

### 8. `StepVerifier`?
"It’s the testing tool for Reactor.
Since streams are async, you can't just assert value X.
You define a script of expectations:
```java
StepVerifier.create(myFlux)
    .expectNext("A")
    .expectNext("B")
    .verifyComplete();
```
It subscribes to your flux and verifies the events strictly match your expectations."

---

### 9. Threading in WebFlux (`Schedulers`)?
"Reactor is concurrency-agnostic by default.
Operators like `zip` or `map` run on the current thread.

If I have a blocking operation (like a legacy JDBC call), I must offload it to a separate thread pool to avoid blocking the Event Loop.
`Mono.fromCallable(blockingTask).subscribeOn(Schedulers.boundedElastic())`.
This is critical—blocking the Event Loop kills the entire app."

---

### 10. Database Access in Reactive?
"You can't use JDBC. JDBC is blocking.
You need **R2DBC** (Reactive Relational Database Connectivity).
It allows non-blocking SQL access given you have a compatible driver (Postgres, MySQL, MSSQL all support it now).
Or you use a naturally reactive store like MongoDB or Cassandra which have had reactive drivers for years."
