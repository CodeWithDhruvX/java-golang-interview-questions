## 🔹 Section 5: Security, Testing, and Deployment (81–100)

### Question 81: How to secure a Spring Boot application with Spring Security?

**Answer:**
1.  Add `spring-boot-starter-security`.
2.  By default, it secures all endpoints with Basic Auth (user/generated-password).
3.  Create a `SecurityFilterChain` bean to customize rules (e.g., permit `/public`, require auth for `/api/**`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to secure a Spring Boot application with Spring Security?
**Your Response:** "To secure a Spring Boot application, I first add the `spring-boot-starter-security` dependency. By default, Spring Security automatically secures all endpoints with Basic Authentication and generates a random password that's printed in the console. For real-world scenarios, I create a `SecurityFilterChain` bean to customize the security rules - for example, I can make certain endpoints like `/public/**` accessible without authentication while securing API endpoints like `/api/**` that require authentication. This gives me fine-grained control over which parts of my application need protection."

---

### Question 82: How to configure Basic Authentication in Spring Boot?

**Answer:**
Usage of `httpBasic()` in the Security Config.

```java
@Bean
public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
    http.authorizeHttpRequests(auth -> auth.anyRequest().authenticated())
        .httpBasic(Customizer.withDefaults()); // Enables Basic Auth
    return http.build();
}
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to configure Basic Authentication in Spring Boot?
**Your Response:** "For Basic Authentication, I configure it in my security configuration class by creating a `SecurityFilterChain` bean. Inside the configuration, I use the `httpBasic()` method with `Customizer.withDefaults()` to enable Basic Authentication. I also ensure that all requests require authentication using `authorizeHttpRequests().anyRequest().authenticated()`. This setup will prompt users with a browser login dialog for credential input, and Spring Security will handle the authentication process automatically."

---

### Question 83: How to implement JWT-based authentication in Spring Boot?

**Answer:**
Spring Security doesn't provide a default JWT implementation. You build it:
1.  **Auth Filter:** Intercept requests, extract `Bearer` token header.
2.  **Validate:** Verify signature/expiration.
3.  **Context:** Set `SecurityContextHolder.getContext().setAuthentication(user)`.
4.  **Login Endpoint:** Generate and return Token on successful credentials.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to implement JWT-based authentication in Spring Boot?
**Your Response:** "JWT authentication requires a custom implementation since Spring Security doesn't provide it out of the box. I create a custom authentication filter that intercepts incoming requests and extracts the JWT token from the `Authorization` header using the `Bearer` scheme. The filter validates the token's signature and expiration date. If valid, I set the authentication in the `SecurityContextHolder`. I also create a login endpoint that authenticates user credentials and generates a JWT token upon successful authentication, which clients then use for subsequent requests."

---

### Question 84: What is CSRF and how do you disable/enable it?

**Answer:**
**Cross-Site Request Forgery.** Attacks using cookies to perform actions on behalf of a logged-in user.
Enabled by default.
For **Stateless APIs (JWT)**, it is safe (and recommended) to **Disable** it, because browsers don't auto-send JWTs like they do Cookies.
`http.csrf(csrf -> csrf.disable())`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is CSRF and how do you disable/enable it?
**Your Response:** "CSRF stands for Cross-Site Request Forgery, which is an attack where malicious websites can perform actions on behalf of a logged-in user using their cookies. Spring Security enables CSRF protection by default. However, for stateless APIs using JWT tokens, it's safe and actually recommended to disable CSRF protection because browsers don't automatically send JWT tokens like they do with cookies. I disable it using `http.csrf(csrf -> csrf.disable())` in my security configuration. For traditional web applications that use session-based authentication, I keep CSRF protection enabled to prevent these types of attacks."

---

### Question 85: How to test Spring Security configurations?

**Answer:**
Use `spring-security-test` dependency.
Use `MockMvc` with `SecurityMockMvcRequestPostProcessors`.
`.with(user("admin").roles("ADMIN"))` allows simulating granular permissions in tests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to test Spring Security configurations?
**Your Response:** "To test Spring Security configurations, I use the `spring-security-test` dependency along with `MockMvc`. This allows me to test security rules by simulating authenticated requests. I use `SecurityMockMvcRequestPostProcessors` to mock different user scenarios - for example, I can test that admin endpoints are properly secured by making requests with `.with(user("admin").roles("ADMIN"))`. This approach lets me verify that my security configuration is working correctly without needing to set up full authentication, ensuring that endpoints are properly protected based on user roles and permissions."

---

### Question 86: How to use `@WithMockUser` in unit tests?

**Answer:**
Annotation for Test Methods.
Simulates a logged-in user in the SecurityContext for that test.
`@WithMockUser(username="user", roles={"USER"})`
Bypasses the actual authentication logic.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to use `@WithMockUser` in unit tests?
**Your Response:** "The `@WithMockUser` annotation is a convenient way to simulate a logged-in user in my test methods. When I add this annotation to a test method, Spring Security automatically populates the `SecurityContext` with a mock user having the specified username and roles. For example, `@WithMockUser(username="user", roles={"USER"})` creates a mock user with those credentials. This is particularly useful for unit testing because it bypasses the actual authentication logic and lets me focus on testing the business logic that depends on user authentication, making my tests faster and more focused."

---

### Question 87: What are unit tests, integration tests, and end-to-end tests in Spring Boot?

**Answer:**
*   **Unit:** Tests isolated class (Service) with Mocks (`@Mock`). Fast.
*   **Integration:** Tests interaction between components (Controller + Service + DB). Uses `@SpringBootTest`. Slower.
*   **E2E:** Tests entire flow (Browser -> App -> DB). Selenium/Cypress.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are unit tests, integration tests, and end-to-end tests in Spring Boot?
**Your Response:** "I categorize tests into three levels. Unit tests focus on testing individual classes in isolation using mocks for dependencies - they're fast and help catch bugs early. Integration tests verify how multiple components work together, like testing a controller that calls a service which interacts with the database - I use `@SpringBootTest` for these. End-to-end tests simulate complete user journeys from the browser through the entire application stack to the database, typically using tools like Selenium or Cypress. Each level provides different confidence and has different trade-offs in terms of speed and complexity."

---

### Question 88: What is `@SpringBootTest`?

**Answer:**
Annotation that boots up the **entire** Spring ApplicationContext.
It is used for Integration Tests.
Usually combined with `@AutoConfigureMockMvc` for API testing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `@SpringBootTest`?
**Your Response:** "`@SpringBootTest` is an annotation that boots up the entire Spring ApplicationContext for integration testing. This means it loads all beans, configurations, and dependencies just like in a real application. I typically use this annotation when I want to test how multiple components work together. It's often combined with `@AutoConfigureMockMvc` when I need to test REST APIs, which provides a fully configured MockMvc instance that I can use to simulate HTTP requests and verify responses in a realistic environment."

---

### Question 89: How to mock beans using Mockito and MockBean?

**Answer:**
*   **Mockito (Unit):** `PaymentService svc = mock(PaymentService.class);`.
*   **`@MockBean` (Spring Test):** Replaces a specific Bean in the Spring Context with a Mock. Useful in Integration tests where you want the real DB but a fake Payment Gateway.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to mock beans using Mockito and MockBean?
**Your Response:** "For mocking, I use both Mockito and Spring's `@MockBean`. Mockito is great for pure unit tests where I create mocks manually using `mock(PaymentService.class)`. But when I'm doing integration tests with Spring, I use `@MockBean` which replaces a real bean in the Spring context with a mock. This is particularly useful when I want to test most of the application realistically but need to fake external dependencies like payment gateways or third-party APIs. The key difference is that `@MockBean` integrates with the Spring context while Mockito works independently."

---

### Question 90: How to test REST APIs with MockMvc or WebTestClient?

**Answer:**
*   **MockMvc:** Servlet-based.
    `mockMvc.perform(get("/api")).andExpect(status().isOk());`
*   **WebTestClient:** Reactive (but works for Servlet too). Can be used to test running servers over HTTP.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to test REST APIs with MockMvc or WebTestClient?
**Your Response:** "For testing REST APIs, I use either MockMvc or WebTestClient depending on the scenario. MockMvc is servlet-based and perfect for traditional Spring MVC applications - I can write tests like `mockMvc.perform(get("/api")).andExpect(status().isOk())`. WebTestClient was originally designed for reactive applications but works for servlet apps too, and has the advantage of being able to test against running servers over HTTP. I choose MockMvc for most unit and integration tests because it's lightweight, but use WebTestClient when I need to test the actual HTTP layer or when working with reactive endpoints."

---

### Question 91: What is the purpose of `@DataJpaTest` and `@WebMvcTest`?

**Answer:**
**Test Slices.** They load only partial context for speed.
*   `@DataJpaTest`: Loads Repositories + EntityManager (No Controllers).
*   `@WebMvcTest`: Loads Controllers + Filters (No Repositories/Services). (Need to use `@MockBean` for dependencies).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `@DataJpaTest` and `@WebMvcTest`?
**Your Response:** "These are test slice annotations that load only specific parts of the Spring context to make tests faster. `@DataJpaTest` loads just the JPA components - repositories, entity manager, and database configuration - which is perfect for testing database operations without loading the entire web layer. `@WebMvcTest` loads only the web layer - controllers, filters, and related components - which is ideal for testing REST endpoints without the database. When using `@WebMvcTest`, I need to mock the service dependencies with `@MockBean` since those aren't loaded. These slice tests are much faster than `@SpringBootTest` because they avoid loading unnecessary components."

---

### Question 92: How to deploy a Spring Boot app on a cloud (AWS/GCP/Azure)?

**Answer:**
1.  **PaaS (Elastic Beanstalk / App Engine):** Upload the JAR. The platform handles the JVM.
2.  **IaaS (EC2):** Install Java, copy JAR, run `java -jar`.
3.  **Container (K8s/ECS):** Dockerize the app, push image, deploy pod.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to deploy a Spring Boot app on a cloud (AWS/GCP/Azure)?
**Your Response:** "I have several options for cloud deployment. The simplest is using PaaS services like AWS Elastic Beanstalk or Google App Engine, where I just upload the JAR file and the platform handles the JVM and infrastructure setup. For more control, I can use IaaS like AWS EC2, where I install Java manually and run the application with `java -jar`. For modern microservices architecture, I prefer containerization - I dockerize the Spring Boot app, push the image to a container registry, and deploy it using Kubernetes or ECS. Each approach offers different levels of control and managed services based on the project requirements."

---

### Question 93: How do you containerize a Spring Boot application using Docker?

**Answer:**
Create a `Dockerfile`:
```dockerfile
FROM eclipse-temurin:17-jdk-alpine
COPY target/app.jar app.jar
ENTRYPOINT ["java", "-jar", "/app.jar"]
```
Run: `docker build -t my-app .`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you containerize a Spring Boot application using Docker?
**Your Response:** "To containerize a Spring Boot application, I create a Dockerfile that starts with a base Java image like `eclipse-temurin:17-jdk-alpine`. I copy the built JAR file from the target directory into the container and set the entry point to run the application with `java -jar`. The multi-stage build approach is particularly efficient - I use a Maven image to build the application, then copy only the JAR file to a minimal runtime image. This keeps the final image small and secure. After building with `docker build -t my-app .`, I can run the containerized application anywhere Docker is available."

---

### Question 94: What is Spring Boot Admin?

**Answer:**
A community project (UI Dashboard) to manage/monitor Spring Boot applications.
It connects to the **Actuator** endpoints of registered clients.
Visualizes Health, Metrics, Logs, Heap Dumps.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Boot Admin?
**Your Response:** "Spring Boot Admin is a community project that provides a web-based dashboard for monitoring and managing Spring Boot applications. It works by having client applications register with the admin server and expose their Actuator endpoints. Through the admin UI, I can view health status, metrics, logs, and even take heap dumps of running applications. It's particularly useful in microservices environments where I need to monitor multiple services from a central location. The admin server aggregates data from all registered applications, making it easy to spot issues and manage the overall system health."

---

### Question 95: How to monitor and health-check a Spring Boot app in production?

**Answer:**
Use **Actuator** (`/actuator/health`).
Connect to Monitoring Tools:
*   **Prometheus:** Scrapes `/actuator/prometheus`.
*   **Grafana:** Visualizes Prometheus data.
*   **New Relic / Datadog:** Use Java Agents.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to monitor and health-check a Spring Boot app in production?
**Your Response:** "For production monitoring, I start with Spring Boot Actuator which provides health check endpoints like `/actuator/health`. I integrate this with monitoring tools - Prometheus can scrape metrics from `/actuator/prometheus`, and Grafana can visualize those metrics in dashboards. For comprehensive application performance monitoring, I use commercial tools like New Relic or Datadog by adding their Java agents to the application. This gives me insights into response times, error rates, and system resource usage. The combination of Actuator for basic health checks and these monitoring tools gives me complete visibility into the application's performance in production."

---

### Question 96: How to enable logging and tracing in Spring Boot?

**Answer:**
*   **Logging:** Uses Logback by default. Configured via `logback-spring.xml`.
*   **Tracing:** Use **Micrometer Tracing** (formerly Spring Cloud Sleuth). Adds TraceID/SpanID to logs so you can follow a request across microservices.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to enable logging and tracing in Spring Boot?
**Your Response:** "Spring Boot uses Logback as the default logging framework, which I configure through `logback-spring.xml` for custom log formats and appenders. For distributed tracing across microservices, I use Micrometer Tracing, which was formerly known as Spring Cloud Sleuth. It automatically adds unique TraceID and SpanID to log entries, allowing me to trace a single request as it flows through multiple services. This is invaluable for debugging issues in distributed systems - I can search logs by TraceID to see the complete journey of a request across all the services it touched."

---

### Question 97: How to handle application startup failures gracefully?

**Answer:**
implement `FailureAnalyzer`.
Spring Boot calls this when an exception occurs during startup.
You can catch specific errors (e.g., Port In Use) and print a user-friendly ASCII description instead of a massive stack trace.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle application startup failures gracefully?
**Your Response:** "I implement custom `FailureAnalyzer` classes to handle startup failures gracefully. When Spring Boot encounters an exception during startup, it calls registered failure analyzers. I can create analyzers for specific common issues - like port conflicts, database connection failures, or missing configuration properties. Instead of showing users a massive stack trace, my analyzer can detect the specific problem and display a clear, actionable error message with suggestions for resolution. This makes troubleshooting much easier for developers and operations teams, especially in production environments where quick diagnosis is critical."

---

### Question 98: How to perform graceful shutdown in Spring Boot?

**Answer:**
Property: `server.shutdown=graceful`.
When SIGTERM is received:
1.  Stops accepting new requests.
2.  Waits for active requests to complete (timeout configurable).
3.  Shuts down.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to perform graceful shutdown in Spring Boot?
**Your Response:** "I enable graceful shutdown by setting the property `server.shutdown=graceful`. When the application receives a SIGTERM signal - like during a deployment or container restart - it stops accepting new requests immediately but continues processing any requests that are already in progress. I can configure a timeout period to give ongoing requests time to complete. After the timeout expires or all active requests finish, the application shuts down gracefully. This prevents cutting off users mid-request and ensures a smooth deployment experience, especially important for APIs that handle long-running operations."

---

### Question 99: How does Spring Boot integrate with message brokers (Kafka, RabbitMQ)?

**Answer:**
*   **Kafka:** `spring-kafka`. use `@KafkaListener` to consume. `KafkaTemplate` to produce.
*   **RabbitMQ:** `spring-boot-starter-amqp`. `@RabbitListener` / `RabbitTemplate`.
Configuration via `application.properties` makes connection setup trivial.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot integrate with message brokers (Kafka, RabbitMQ)?
**Your Response:** "Spring Boot provides excellent integration with message brokers through auto-configuration. For Kafka, I add the `spring-kafka` dependency and use `@KafkaListener` annotations to create consumers, while `KafkaTemplate` handles message production. For RabbitMQ, I use `spring-boot-starter-amqp` with `@RabbitListener` and `RabbitTemplate`. The beauty of Spring Boot's auto-configuration is that I only need to define the connection details in `application.properties` - Spring automatically creates the connection factories, templates, and listener containers. This makes it incredibly simple to implement messaging patterns without dealing with the boilerplate setup code."

---

### Question 100: What is reactive programming in Spring Boot and how does it differ from traditional MVC?

**Answer:**
Uses **Spring WebFlux** (Netty server).
*   **MVC:** One Thread per Request. Blocking I/O.
*   **Reactive:** Event Loop model (few threads). Non-Blocking I/O.
Handles massive concurrency with low memory footprint. Uses `Mono` (0-1) and `Flux` (0-N) streams.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is reactive programming in Spring Boot and how does it differ from traditional MVC?
**Your Response:** "Reactive programming in Spring Boot uses Spring WebFlux with a non-blocking, event-driven architecture, unlike traditional MVC which uses a thread-per-request blocking model. In MVC, each request ties up a thread until completion, limiting scalability. WebFlux uses an event loop with few threads handling many concurrent requests through non-blocking I/O. This allows handling massive concurrent connections with minimal memory usage. The programming model is also different - instead of returning simple objects, I work with reactive streams using `Mono` for 0-1 items and `Flux` for 0-N items. This approach is ideal for I/O-intensive applications that need to handle high concurrency efficiently."

---
