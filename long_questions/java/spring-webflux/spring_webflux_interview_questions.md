# Spring WebFlux Interview Questions

## Core Properties & Reactive Programming

### 1. What is Spring WebFlux?
Spring WebFlux is a parallel version of Spring MVC introduced in Spring 5 to support reactive programming. It is non-blocking, asynchronous, and event-driven. It is designed to handle high concurrency with a small number of threads.

**Explanation:** WebFlux uses reactive programming principles to handle I/O operations without blocking threads, making it ideal for applications that need to handle many concurrent connections with limited resources.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring WebFlux?
**Your Response:** Spring WebFlux is Spring's reactive web framework that runs on non-blocking servers. Unlike traditional Spring MVC which uses a thread-per-request model, WebFlux uses a small, fixed number of threads to handle many concurrent requests asynchronously. It's built on Project Reactor and uses Mono and Flux types to handle data streams. This makes it perfect for applications that need high scalability and low latency, like microservices, streaming applications, or APIs that handle many concurrent connections.

### 2. What is Reactive Programming?
It is a programming paradigm oriented around data streams and the propagation of change. It provides a way to handle asynchronous data streams with non-blocking backpressure.

**Explanation:** Reactive programming treats data as streams that can be observed and transformed, enabling efficient handling of asynchronous operations without blocking threads.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Reactive Programming?
**Your Response:** Reactive programming is a paradigm where I work with data streams and the propagation of changes. Instead of pulling data explicitly, I subscribe to streams that push data to me when it becomes available. This is perfect for handling asynchronous operations like database calls or web service requests. The key benefits are that it's non-blocking - threads aren't waiting for I/O operations to complete - and it supports backpressure, which prevents fast producers from overwhelming slow consumers. This makes applications more responsive and scalable.

### 3. Explain Mono vs. Flux.
These are the two core types in Project Reactor (the reactive library WebFlux is built on):
- **Mono<T>:** Represents a stream of **0 or 1** element. (Like `Optional<T>` but async).
- **Flux<T>:** Represents a stream of **0 to N** elements. (Like `List<T>` but async).

**Explanation:** Mono and Flux are the fundamental building blocks of reactive programming in Spring, representing different cardinalities of asynchronous data streams.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain Mono vs. Flux?
**Your Response:** Mono and Flux are the two core reactive types in Project Reactor. Mono represents a stream that can emit 0 or 1 items - it's like an asynchronous Optional. I use Mono when I expect a single result, like finding a user by ID or making a single database call. Flux represents a stream that can emit 0 to N items - it's like an asynchronous List. I use Flux when I'm dealing with multiple items, like streaming all users from a database or processing a sequence of events. Both types support operators like map, filter, and flatMap to transform the data streams.

### 4. What is Backpressure?
Backpressure is a mechanism that allows a consumer to signal to a producer how much data it can handle. This prevents the consumer from being overwhelmed by a fast producer. In Reactive Streams, the subscriber requests *n* items, and the publisher sends at most *n*.

**Explanation:** Backpressure ensures system stability by preventing fast data producers from overwhelming slower consumers, maintaining resource utilization without causing failures.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Backpressure?
**Your Response:** Backpressure is a crucial mechanism in reactive programming that prevents system overload. It allows the consumer to tell the producer how much data it can handle at any given time. For example, if I'm streaming data from a fast source like a message queue but my database writes are slower, backpressure prevents the queue from overwhelming my application. The consumer requests a specific number of items, processes them, then requests more. This ensures that the system remains stable and responsive even when there's a mismatch between production and consumption rates.

### 5. How does WebFlux differ from Spring MVC?
| Feature | Spring MVC | Spring WebFlux |
| :--- | :--- | :--- |
| **IO Model** | Blocking (Synchronous) | Non-Blocking (Asynchronous) |
| **Server** | Servlet Containers (Tomcat) | Netty, Undertow (Event Loop) |
| **Concurrency** | One Thread per Request | Event Loop (Few Threads) |
| **Best For** | CPU-bound, traditional apps | High IO, Streaming, Scalability |

**Explanation:** The choice between MVC and WebFlux depends on the application's characteristics - MVC for traditional blocking I/O and CPU-bound tasks, WebFlux for high-concurrency I/O-bound scenarios.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does WebFlux differ from Spring MVC?
**Your Response:** The main difference is in their concurrency models. Spring MVC uses a traditional blocking approach where each request gets its own thread, making it great for CPU-intensive applications. Spring WebFlux is non-blocking and uses an event loop model with a small number of threads to handle many concurrent requests. WebFlux runs best on servers like Netty that are designed for non-blocking I/O, while MVC runs on traditional servlet containers like Tomcat. I choose MVC for traditional applications and WebFlux when I need high scalability for I/O-bound operations like microservices or streaming applications.

## Implementation & Testing

### 6. Controller Programming Model (Annotation-based).
WebFlux supports the same `@Controller`, `@RequestMapping`, `@GetMapping` annotations as Spring MVC. The difference is the return type (`Mono` or `Flux`).
```java
@GetMapping("/users")
public Flux<User> getAllUsers() {
    return userRepository.findAll();
}
```

**Explanation:** The annotation-based programming model provides a familiar development experience while leveraging reactive types for non-blocking operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the controller programming model work in WebFlux?
**Your Response:** WebFlux supports the same annotation-based programming model as Spring MVC, so I can still use @Controller, @GetMapping, and other familiar annotations. The key difference is that my controller methods return reactive types - Mono for single results or Flux for multiple results. For example, instead of returning List<User>, I return Flux<User> from a @GetMapping method. Spring WebFlux handles the subscription and streaming automatically. This makes the transition from MVC to WebFlux much easier since I can keep the same controller structure while making my endpoints non-blocking.

### 7. Functional Endpoints (Router Functions).
WebFlux introduces a functional programming model as an alternative to annotatons.
- **RouterFunction:** Routes requests to handler functions (like `@RequestMapping`).
- **HandlerFunction:** Handles the request and returns a response (like the Controller method body).

**Explanation:** Functional endpoints provide a more explicit and type-safe way to define routing without reflection, making the configuration more testable and maintainable.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are Functional Endpoints in WebFlux?
**Your Response:** Functional endpoints are an alternative to the annotation-based programming model in WebFlux. Instead of using @Controller and @RequestMapping, I define RouterFunction and HandlerFunction beans. The RouterFunction maps incoming requests to appropriate handler functions, similar to how @RequestMapping works. The HandlerFunction contains the actual request handling logic. This approach is more functional and type-safe, avoids reflection, and can be easier to test. It's particularly useful for building APIs where I want more explicit control over the routing configuration.

### 8. What is `WebClient`?
`WebClient` is a non-blocking, reactive client for performing HTTP requests. It is the modern alternative to `RestTemplate`. It supports both synchronous (blocking) and asynchronous scenarios.

**Explanation:** WebClient provides a reactive HTTP client that can handle streaming responses and backpressure, making it ideal for microservice communication in reactive applications.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is WebClient?
**Your Response:** WebClient is Spring's modern, reactive HTTP client that replaces the older RestTemplate. It's non-blocking and built on Project Reactor, so it returns Mono or Flux types. This makes it perfect for calling other services in a reactive application. WebClient supports both synchronous and asynchronous usage, can handle streaming responses, and includes features like retry, timeout, and backpressure handling. I use it when I need to make HTTP calls from my WebFlux application to other microservices, ensuring that my entire request pipeline remains non-blocking.

### 9. How do you handle errors in WebFlux?
Since traditional try-catch blocks don't work well in async streams, we use operators:
- `onErrorReturn()`: Return a default value.
- `onErrorResume()`: Switch to a fallback sequence.
- `onErrorMap()`: Transform the exception into another.

**Explanation:** Reactive error handling uses declarative operators to manage exceptions in asynchronous streams, maintaining the non-blocking nature of the reactive pipeline.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle errors in WebFlux?
**Your Response:** Error handling in reactive programming is different from traditional try-catch blocks because the operations are asynchronous. Instead, I use reactive operators to handle errors declaratively. onErrorReturn lets me provide a fallback value when an error occurs. onErrorResume allows me to switch to an alternative reactive sequence. onErrorMap lets me transform one type of exception into another. These operators keep the reactive flow going even when errors happen, which is essential for building resilient reactive applications. I can also use the doOnError operator for logging errors without changing the stream behavior.

### 10. Which server is the default for Spring WebFlux?
**Netty**. It is an asynchronous event-driven network application framework. However, WebFlux can also run on Tomcat, Jetty, or Undertow as long as they support Servlet 3.1+ (Non-blocking IO).

**Explanation:** Netty's event-driven architecture is optimized for non-blocking I/O operations, making it the ideal server choice for reactive applications, though WebFlux maintains compatibility with traditional servlet containers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which server is the default for Spring WebFlux?
**Your Response:** Netty is the default server for Spring WebFlux because it's specifically designed for non-blocking, event-driven applications. Netty's architecture with its event loop model is perfect for handling many concurrent connections efficiently. However, WebFlux is flexible - it can also run on traditional servlet containers like Tomcat, Jetty, or Undertow as long as they support Servlet 3.1+ with non-blocking I/O. I typically use Netty for pure reactive applications, but might choose Tomcat if I need to integrate with existing servlet-based infrastructure or have specific operational requirements.

## Advanced WebFlux Topics

### 11. Cold vs. Hot Streams.
- **Cold:** Does not generate data until a subscriber subscribes. Each subscriber gets its own fresh stream of data (re-started). Most Reactor publishers (`Flux.just`, database calls) are cold.
- **Hot:** Broadcasts data regardless of subscribers. Late subscribers miss previously emitted data. (e.g., Mouse events, stock ticker).

**Explanation:** Cold streams provide isolation between subscribers while hot streams share data, making them suitable for different use cases like on-demand data versus live events.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between cold and hot streams?
**Your Response:** Cold and hot streams behave differently regarding data sharing. Cold streams don't emit data until someone subscribes, and each subscriber gets its own independent stream. For example, a database query wrapped in Flux.just is cold - each subscription triggers a new query. Hot streams emit data regardless of whether anyone is subscribed, like a stock ticker or mouse events. If I subscribe late to a hot stream, I miss the data that was already emitted. I use cold streams for on-demand data fetching and hot streams for live events or when I want multiple subscribers to share the same data source.

### 12. How do you test Reactive Streams?
Use **`StepVerifier`**. It allows you to define expectations on the sequence of data signals.
```java
StepVerifier.create(flux)
    .expectNext("a")
    .expectNext("b")
    .verifyComplete();
```

**Explanation:** StepVerifier provides a declarative testing approach for reactive streams, allowing precise verification of emission timing, values, and error conditions.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test Reactive Streams?
**Your Response:** I test reactive streams using StepVerifier, which is Spring's testing utility for Project Reactor. StepVerifier allows me to define expectations about what should happen in a reactive stream. I can expect specific values to be emitted in order, verify that the stream completes successfully, or assert that certain errors occur. For example, I can test that a Flux emits 'a' then 'b' then completes, or that a Mono emits a specific value. StepVerifier also supports testing time-based operators and virtual time for testing time-sensitive operations without actually waiting.

### 13. What is R2DBC?
**Reactive Relational Database Connectivity**. It is an API specification for reactive programming with SQL databases. Unlike JDBC (which is blocking), R2DBC allows fully non-blocking database access, making it suitable for WebFlux applications.

**Explanation:** R2DBC brings reactive programming to database access, eliminating the blocking nature of traditional JDBC and maintaining the non-blocking benefits throughout the entire application stack.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is R2DBC?
**Your Response:** R2DBC stands for Reactive Relational Database Connectivity. It's a specification for reactive database access that's the reactive equivalent of JDBC. While traditional JDBC is blocking and would defeat the purpose of WebFlux, R2DBC allows me to interact with relational databases in a fully non-blocking way. This means I can build end-to-end reactive applications where even the database layer doesn't block threads. R2DBC drivers are available for major databases like PostgreSQL, MySQL, and SQL Server, allowing me to maintain the reactive benefits throughout my entire application stack.

### 14. Explain Schedulers in Reactor.
Schedulers allow switching the execution context (threading model).
- `subscribeOn()`: Affects the context of the source emission (where subscription happens).
- `publishOn()`: Affects the execution context of subsequent operators in the chain.
- Common Schedulers: `Schedulers.boundedElastic()` (for blocking IO), `Schedulers.parallel()` (for CPU tasks).

**Explanation:** Schedulers provide control over thread execution in reactive streams, enabling optimal resource utilization by separating subscription, computation, and I/O operations across different thread pools.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain Schedulers in Reactor?
**Your Response:** Schedulers in Reactor control where reactive operations run on different threads. subscribeOn determines which thread the subscription and initial data emission happens on, while publishOn affects where subsequent operators in the chain execute. I use different schedulers for different tasks - Schedulers.parallel() for CPU-intensive work, Schedulers.boundedElastic() for blocking I/O operations, and Schedulers.single() for sequential work. This allows me to optimize performance by ensuring that blocking operations don't block the event loop threads, keeping my reactive application responsive and scalable.
