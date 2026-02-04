# Spring WebFlux Interview Questions

## Core Properties & Reactive Programming

### 1. What is Spring WebFlux?
Spring WebFlux is a parallel version of Spring MVC introduced in Spring 5 to support reactive programming. It is non-blocking, asynchronous, and event-driven. It is designed to handle high concurrency with a small number of threads.

### 2. What is Reactive Programming?
It is a programming paradigm oriented around data streams and the propagation of change. It provides a way to handle asynchronous data streams with non-blocking backpressure.

### 3. Explain Mono vs. Flux.
These are the two core types in Project Reactor (the reactive library WebFlux is built on):
- **Mono<T>:** Represents a stream of **0 or 1** element. (Like `Optional<T>` but async).
- **Flux<T>:** Represents a stream of **0 to N** elements. (Like `List<T>` but async).

### 4. What is Backpressure?
Backpressure is a mechanism that allows a consumer to signal to a producer how much data it can handle. This prevents the consumer from being overwhelmed by a fast producer. In Reactive Streams, the subscriber requests *n* items, and the publisher sends at most *n*.

### 5. How does WebFlux differ from Spring MVC?
| Feature | Spring MVC | Spring WebFlux |
| :--- | :--- | :--- |
| **IO Model** | Blocking (Synchronous) | Non-Blocking (Asynchronous) |
| **Server** | Servlet Containers (Tomcat) | Netty, Undertow (Event Loop) |
| **Concurrency** | One Thread per Request | Event Loop (Few Threads) |
| **Best For** | CPU-bound, traditional apps | High IO, Streaming, Scalability |

## Implementation & Testing

### 6. Controller Programming Model (Annotation-based).
WebFlux supports the same `@Controller`, `@RequestMapping`, `@GetMapping` annotations as Spring MVC. The difference is the return type (`Mono` or `Flux`).
```java
@GetMapping("/users")
public Flux<User> getAllUsers() {
    return userRepository.findAll();
}
```

### 7. Functional Endpoints (Router Functions).
WebFlux introduces a functional programming model as an alternative to annotatons.
- **RouterFunction:** Routes requests to handler functions (like `@RequestMapping`).
- **HandlerFunction:** Handles the request and returns a response (like the Controller method body).

### 8. What is `WebClient`?
`WebClient` is a non-blocking, reactive client for performing HTTP requests. It is the modern alternative to `RestTemplate`. It supports both synchronous (blocking) and asynchronous scenarios.

### 9. How do you handle errors in WebFlux?
Since traditional try-catch blocks don't work well in async streams, we use operators:
- `onErrorReturn()`: Return a default value.
- `onErrorResume()`: Switch to a fallback sequence.
- `onErrorMap()`: Transform the exception into another.

### 10. Which server is the default for Spring WebFlux?
**Netty**. It is an asynchronous event-driven network application framework. However, WebFlux can also run on Tomcat, Jetty, or Undertow as long as they support Servlet 3.1+ (Non-blocking IO).

## Advanced WebFlux Topics

### 11. Cold vs. Hot Streams.
- **Cold:** Does not generate data until a subscriber subscribes. Each subscriber gets its own fresh stream of data (re-started). Most Reactor publishers (`Flux.just`, database calls) are cold.
- **Hot:** Broadcasts data regardless of subscribers. Late subscribers miss previously emitted data. (e.g., Mouse events, stock ticker).

### 12. How do you test Reactive Streams?
Use **`StepVerifier`**. It allows you to define expectations on the sequence of data signals.
```java
StepVerifier.create(flux)
    .expectNext("a")
    .expectNext("b")
    .verifyComplete();
```

### 13. What is R2DBC?
**Reactive Relational Database Connectivity**. It is an API specification for reactive programming with SQL databases. Unlike JDBC (which is blocking), R2DBC allows fully non-blocking database access, making it suitable for WebFlux applications.

### 14. Explain Schedulers in Reactor.
Schedulers allow switching the execution context (threading model).
- `subscribeOn()`: Affects the context of the source emission (where subscription happens).
- `publishOn()`: Affects the execution context of subsequent operators in the chain.
- Common Schedulers: `Schedulers.boundedElastic()` (for blocking IO), `Schedulers.parallel()` (for CPU tasks).
