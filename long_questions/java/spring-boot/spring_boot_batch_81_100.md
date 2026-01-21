## 🔹 Section 5: Security, Testing, and Deployment (81–100)

### Question 81: How to secure a Spring Boot application with Spring Security?

**Answer:**
1.  Add `spring-boot-starter-security`.
2.  By default, it secures all endpoints with Basic Auth (user/generated-password).
3.  Create a `SecurityFilterChain` bean to customize rules (e.g., permit `/public`, require auth for `/api/**`).

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

---

### Question 83: How to implement JWT-based authentication in Spring Boot?

**Answer:**
Spring Security doesn't provide a default JWT implementation. You build it:
1.  **Auth Filter:** Intercept requests, extract `Bearer` token header.
2.  **Validate:** Verify signature/expiration.
3.  **Context:** Set `SecurityContextHolder.getContext().setAuthentication(user)`.
4.  **Login Endpoint:** Generate and return Token on successful credentials.

---

### Question 84: What is CSRF and how do you disable/enable it?

**Answer:**
**Cross-Site Request Forgery.** Attacks using cookies to perform actions on behalf of a logged-in user.
Enabled by default.
For **Stateless APIs (JWT)**, it is safe (and recommended) to **Disable** it, because browsers don't auto-send JWTs like they do Cookies.
`http.csrf(csrf -> csrf.disable())`.

---

### Question 85: How to test Spring Security configurations?

**Answer:**
Use `spring-security-test` dependency.
Use `MockMvc` with `SecurityMockMvcRequestPostProcessors`.
`.with(user("admin").roles("ADMIN"))` allows simulating granular permissions in tests.

---

### Question 86: How to use `@WithMockUser` in unit tests?

**Answer:**
Annotation for Test Methods.
Simulates a logged-in user in the SecurityContext for that test.
`@WithMockUser(username="user", roles={"USER"})`
Bypasses the actual authentication logic.

---

### Question 87: What are unit tests, integration tests, and end-to-end tests in Spring Boot?

**Answer:**
*   **Unit:** Tests isolated class (Service) with Mocks (`@Mock`). Fast.
*   **Integration:** Tests interaction between components (Controller + Service + DB). Uses `@SpringBootTest`. Slower.
*   **E2E:** Tests entire flow (Browser -> App -> DB). Selenium/Cypress.

---

### Question 88: What is `@SpringBootTest`?

**Answer:**
Annotation that boots up the **entire** Spring ApplicationContext.
It is used for Integration Tests.
Usually combined with `@AutoConfigureMockMvc` for API testing.

---

### Question 89: How to mock beans using Mockito and MockBean?

**Answer:**
*   **Mockito (Unit):** `PaymentService svc = mock(PaymentService.class);`.
*   **`@MockBean` (Spring Test):** Replaces a specific Bean in the Spring Context with a Mock. Useful in Integration tests where you want the real DB but a fake Payment Gateway.

---

### Question 90: How to test REST APIs with MockMvc or WebTestClient?

**Answer:**
*   **MockMvc:** Servlet-based.
    `mockMvc.perform(get("/api")).andExpect(status().isOk());`
*   **WebTestClient:** Reactive (but works for Servlet too). Can be used to test running servers over HTTP.

---

### Question 91: What is the purpose of `@DataJpaTest` and `@WebMvcTest`?

**Answer:**
**Test Slices.** They load only partial context for speed.
*   `@DataJpaTest`: Loads Repositories + EntityManager (No Controllers).
*   `@WebMvcTest`: Loads Controllers + Filters (No Repositories/Services). (Need to use `@MockBean` for dependencies).

---

### Question 92: How to deploy a Spring Boot app on a cloud (AWS/GCP/Azure)?

**Answer:**
1.  **PaaS (Elastic Beanstalk / App Engine):** Upload the JAR. The platform handles the JVM.
2.  **IaaS (EC2):** Install Java, copy JAR, run `java -jar`.
3.  **Container (K8s/ECS):** Dockerize the app, push image, deploy pod.

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

---

### Question 94: What is Spring Boot Admin?

**Answer:**
A community project (UI Dashboard) to manage/monitor Spring Boot applications.
It connects to the **Actuator** endpoints of registered clients.
Visualizes Health, Metrics, Logs, Heap Dumps.

---

### Question 95: How to monitor and health-check a Spring Boot app in production?

**Answer:**
Use **Actuator** (`/actuator/health`).
Connect to Monitoring Tools:
*   **Prometheus:** Scrapes `/actuator/prometheus`.
*   **Grafana:** Visualizes Prometheus data.
*   **New Relic / Datadog:** Use Java Agents.

---

### Question 96: How to enable logging and tracing in Spring Boot?

**Answer:**
*   **Logging:** Uses Logback by default. Configured via `logback-spring.xml`.
*   **Tracing:** Use **Micrometer Tracing** (formerly Spring Cloud Sleuth). Adds TraceID/SpanID to logs so you can follow a request across microservices.

---

### Question 97: How to handle application startup failures gracefully?

**Answer:**
implement `FailureAnalyzer`.
Spring Boot calls this when an exception occurs during startup.
You can catch specific errors (e.g., Port In Use) and print a user-friendly ASCII description instead of a massive stack trace.

---

### Question 98: How to perform graceful shutdown in Spring Boot?

**Answer:**
Property: `server.shutdown=graceful`.
When SIGTERM is received:
1.  Stops accepting new requests.
2.  Waits for active requests to complete (timeout configurable).
3.  Shuts down.

---

### Question 99: How does Spring Boot integrate with message brokers (Kafka, RabbitMQ)?

**Answer:**
*   **Kafka:** `spring-kafka`. use `@KafkaListener` to consume. `KafkaTemplate` to produce.
*   **RabbitMQ:** `spring-boot-starter-amqp`. `@RabbitListener` / `RabbitTemplate`.
Configuration via `application.properties` makes connection setup trivial.

---

### Question 100: What is reactive programming in Spring Boot and how does it differ from traditional MVC?

**Answer:**
Uses **Spring WebFlux** (Netty server).
*   **MVC:** One Thread per Request. Blocking I/O.
*   **Reactive:** Event Loop model (few threads). Non-Blocking I/O.
Handles massive concurrency with low memory footprint. Uses `Mono` (0-1) and `Flux` (0-N) streams.

---
