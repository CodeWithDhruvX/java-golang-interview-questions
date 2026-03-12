## 🔹 Section 3: Testing in Spring Boot (441-460)

### Question 441: What is the difference between `@SpringBootTest` and `@WebMvcTest`?

**Answer:**
(See Q91/Q88). `SpringBootTest` is Full Context. `WebMvcTest` is Slice.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@SpringBootTest` and `@WebMvcTest`?
**Your Response:** "`@SpringBootTest` loads the full application context, making it suitable for integration tests that need the complete application setup. `@WebMvcTest` is a slice test that only loads the web layer - controllers, filters, and related components - making it much faster for testing just the MVC layer. I use `@WebMvcTest` when I want to test my controllers in isolation without loading the entire application, which significantly reduces test execution time and focuses the test on specific functionality."

---

### Question 442: How do you test REST controllers in Spring Boot?

**Answer:**
(See Q90). Using `MockMvc`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test REST controllers in Spring Boot?
**Your Response:** "I test REST controllers using `MockMvc`, which provides a powerful testing framework for Spring MVC applications. I can perform HTTP requests and assert responses without actually running a server. For example, I use `mockMvc.perform(get('/api/users')).andExpect(status().isOk())` to test endpoints. `MockMvc` allows me to test request handling, response status, headers, and JSON content in a fast, reliable way. It's perfect for unit testing controllers with realistic HTTP interactions."

---

### Question 443: How do you mock services using `@MockBean`?

**Answer:**
(See Q89).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you mock services using `@MockBean`?
**Your Response:** "I mock services using `@MockBean` which creates a Mockito mock and replaces the actual bean in the Spring context. This allows me to test components in isolation by controlling the behavior of their dependencies. For example, I can mock a UserService to return specific test data when testing a UserController. The mock is automatically injected into any class that depends on it, making it easy to test scenarios without needing real implementations of external services or databases."

---

### Question 444: What is the use of `TestRestTemplate` in Spring Boot?

**Answer:**
(See Q59). Real HTTP Client for Integration tests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `TestRestTemplate` in Spring Boot?
**Your Response:** "`TestRestTemplate` is a real HTTP client designed for integration tests. Unlike `MockMvc`, it makes actual HTTP requests to a running application, making it perfect for end-to-end testing. I use it to test the complete request-response cycle, including JSON serialization, validation, and error handling. It's particularly useful for testing that the application works correctly when deployed, as it tests the actual HTTP layer rather than mocking it."

---

### Question 445: How do you test JPA repositories effectively?

**Answer:**
(See Q80). `@DataJpaTest`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test JPA repositories effectively?
**Your Response:** "I test JPA repositories effectively using `@DataJpaTest`, which is a slice test that only loads the JPA configuration and repository beans. It uses an in-memory database by default, making tests fast and isolated. I can test repository methods, custom queries, and entity relationships without loading the entire application. This focused testing approach ensures my data access layer works correctly while keeping test execution fast and reliable."

---

### Question 446: How do you use in-memory databases for unit tests?

**Answer:**
(See Q80). H2 dependency + `@DataJpaTest`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use in-memory databases for unit tests?
**Your Response:** "I use in-memory databases like H2 for unit tests by adding the H2 dependency and using `@DataJpaTest`. Spring Boot automatically configures H2 as the test database, allowing me to test JPA entities and repositories without needing an external database. The in-memory database starts fresh for each test, ensuring test isolation and fast execution. This approach is perfect for testing data access logic without the overhead of managing a real database during testing."

---

### Question 447: What is `@DataJpaTest` used for?

**Answer:**
(Duplicate of 445).

---

### Question 448: How do you perform integration testing with embedded containers?

**Answer:**
Using **TestContainers**.
(See Q202).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you perform integration testing with embedded containers?
**Your Response:** "I perform integration testing with embedded containers using TestContainers. TestContainers provides lightweight, throwaway instances of databases, message brokers, and other services that run inside Docker containers during tests. This allows me to test against real services rather than mocks, ensuring my application works correctly with actual external dependencies. TestContainers manages the container lifecycle automatically, making integration tests more realistic and reliable while keeping them portable across different environments."

---

### Question 449: What is the purpose of `@TestConfiguration`?

**Answer:**
(See Q211). Register extra beans for tests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `@TestConfiguration`?
**Your Response:** "I use `@TestConfiguration` to register extra beans specifically for test scenarios. Unlike regular `@Configuration` classes, `@TestConfiguration` is only loaded when explicitly imported in tests. This allows me to define test-specific beans like mocks, test data loaders, or alternative configurations without affecting the main application configuration. It's perfect for setting up test environments that need slight variations from the production configuration."

---

### Question 450: How do you test exception scenarios in Spring Boot?

**Answer:**
(See Q206). `MockMvc` expectation of error message/status.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test exception scenarios in Spring Boot?
**Your Response:** "I test exception scenarios using `MockMvc` to perform requests that should trigger exceptions and then assert the expected error responses. I can use `.andExpect(status().isBadRequest())` to check HTTP status codes and `.andExpect(jsonPath('$.message').value('Expected error'))` to verify error messages. This approach ensures that my error handling works correctly and clients receive appropriate error information when things go wrong."

---

### Question 451: What is the difference between `@SpringBootTest` and `@WebMvcTest`?

**Answer:**
(Duplicate of 441).

---

### Question 452: How do you test REST controllers in Spring Boot?

**Answer:**
(Duplicate of 442).

---

### Question 453: How do you mock services using `@MockBean`?

**Answer:**
(Duplicate of 443).

---

### Question 454: What is the use of `TestRestTemplate` in Spring Boot?

**Answer:**
(Duplicate of 444).

---

### Question 455: How do you test JPA repositories effectively?

**Answer:**
(Duplicate of 445).

---

### Question 456: How do you use in-memory databases for unit tests?

**Answer:**
(Duplicate of 446).

---

### Question 457: What is `@DataJpaTest` used for?

**Answer:**
(Duplicate of 447).

---

### Question 458: How do you perform integration testing with embedded containers?

**Answer:**
(Duplicate of 448).

---

### Question 459: What is the purpose of `@TestConfiguration`?

**Answer:**
(Duplicate of 449).

---

### Question 460: How do you test exception scenarios in Spring Boot?

**Answer:**
(Duplicate of 450).

## 🔹 Section 4: Deployment, Profiles & Environment (461-480)

### Question 461: What are Spring Boot profiles and how do they work?

**Answer:**
(See Q20/Q29). Logical groups of configuration.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are Spring Boot profiles and how do they work?
**Your Response:** "Spring Boot profiles are logical groups of configuration that allow me to customize application behavior for different environments. I can define profiles like 'dev', 'test', and 'production', each with their own configuration properties. When I activate a profile, Spring Boot loads the corresponding configuration files and beans. This mechanism enables me to use the same application code across different environments while customizing database connections, logging levels, and other settings for each environment."

---

### Question 462: How do you activate multiple profiles simultaneously?

**Answer:**
`spring.profiles.active=dev,debug,cloud`.
Separated by comma.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you activate multiple profiles simultaneously?
**Your Response:** "I activate multiple profiles simultaneously by separating them with commas in the `spring.profiles.active` property. For example, `spring.profiles.active=dev,debug,cloud` activates the dev, debug, and cloud profiles together. This allows me to combine configurations - for instance, having a base environment profile plus additional profiles for specific features like debugging or cloud deployment. The profiles are merged, with later profiles potentially overriding earlier ones if there are conflicting properties."

---

### Question 463: How do you deploy a Spring Boot app as a WAR file?

**Answer:**
(See Q12). Extend `SpringBootServletInitializer`. Change packing to `war`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you deploy a Spring Boot app as a WAR file?
**Your Response:** "I deploy a Spring Boot app as a WAR file by extending `SpringBootServletInitializer` in my main class and changing the packaging to `war` in Maven. This makes the application compatible with traditional servlet containers like Tomcat or JBoss. The initializer allows the external container to manage the application lifecycle instead of Spring Boot's embedded server. This approach is useful when I need to deploy to existing infrastructure that requires WAR files rather than standalone JARs."

---

### Question 464: What is the difference between embedded Tomcat and external Tomcat deployment?

**Answer:**
- **Embedded:** JAR, managed by Boot, isolated.
- **External:** WAR, managed by Ops, shared Tomcat instance (Standard Servlet Container).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between embedded Tomcat and external Tomcat deployment?
**Your Response:** "Embedded Tomcat deployment packages the server inside the JAR, making it self-contained and isolated. External Tomcat deployment packages the application as a WAR file that runs in a shared Tomcat instance managed by operations. Embedded deployment is simpler and more portable, while external deployment allows multiple applications to share resources and be managed centrally. I choose embedded for microservices and cloud deployment, and external for traditional enterprise environments where operations teams manage the infrastructure."

---

### Question 465: How do you deploy Spring Boot apps on Heroku?

**Answer:**
Push code to Heroku Git or `heroku deploy:jar`.
Heroku detects `pom.xml`, builds app, and runs using `Procfile` (`web: java -jar ...`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you deploy Spring Boot apps on Heroku?
**Your Response:** "I deploy Spring Boot apps to Heroku by pushing code to Heroku Git or using `heroku deploy:jar`. Heroku automatically detects the `pom.xml`, builds the application using Maven, and runs it using a `Procfile` that specifies `web: java -jar ...`. Heroku handles the infrastructure, scaling, and process management. This platform-as-a-service approach simplifies deployment significantly - I just push my code, and Heroku takes care of the rest."

---

### Question 466: What is `spring.profiles.active` and where can you define it?

**Answer:**
Property to enable profiles.
Defined in `application.properties`, CLI Arg, or Env Var.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `spring.profiles.active` and where can you define it?
**Your Response:** "`spring.profiles.active` is the property I use to enable specific Spring Boot profiles. I can define it in multiple places: in `application.properties` for default configuration, as a command-line argument like `-Dspring.profiles.active=prod`, or as an environment variable `SPRING_PROFILES_ACTIVE=prod`. The precedence follows Spring Boot's property hierarchy, so command-line arguments override environment variables, which override configuration files. This flexibility allows me to control profile activation in different deployment scenarios."

---

### Question 467: How do you implement zero-downtime deployment for Spring Boot apps?

**Answer:**
(See Q268). Rolling updates + Graceful Shutdown.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement zero-downtime deployment for Spring Boot apps?
**Your Response:** "I implement zero-downtime deployment using rolling updates combined with graceful shutdown. In Kubernetes, I configure the deployment to update pods one at a time, waiting for each new pod to pass its readiness probe before terminating the old one. I enable graceful shutdown in Spring Boot so the application finishes processing current requests before shutting down. This approach ensures that users never experience service interruptions during deployments, as there's always at least one healthy instance serving traffic."

---

### Question 468: How do you enable graceful shutdown in Spring Boot?

**Answer:**
(See Q98). `server.shutdown=graceful`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you enable graceful shutdown in Spring Boot?
**Your Response:** "I enable graceful shutdown by setting `server.shutdown=graceful` in configuration. This tells Spring Boot to wait for ongoing requests to complete before shutting down. I also configure `spring.lifecycle.timeout-per-shutdown-phase` to control how long to wait. During shutdown, Spring Boot stops accepting new requests but continues processing existing ones. This prevents data loss and provides a better user experience during deployments or restarts, ensuring in-flight operations complete gracefully."

---

### Question 469: How do you deploy Spring Boot apps to AWS Lambda?

**Answer:**
Use **Spring Cloud Function**.
Adapter `SpringBootRequestHandler` or `FunctionInvoker`.
Wraps the function, providing cold-start optimization.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you deploy Spring Boot apps to AWS Lambda?
**Your Response:** "I deploy Spring Boot apps to AWS Lambda using Spring Cloud Function. I use adapters like `SpringBootRequestHandler` or `FunctionInvoker` that wrap the Spring Boot application for Lambda execution. Spring Cloud Function provides cold-start optimization by initializing the application context efficiently. This approach allows me to run Spring Boot applications in a serverless environment, benefiting from Lambda's scalability and pay-per-use model while maintaining Spring's familiar programming model."

---

### Question 470: What is the role of `application.properties` vs `application.yml`?

**Answer:**
(See Q6). They serve same purpose. YAML is hierarchical.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `application.properties` vs `application.yml`?
**Your Response:** "Both `application.properties` and `application.yml` serve the same purpose - configuring Spring Boot applications. The key difference is that YAML provides hierarchical structure which makes complex configurations more readable and organized. I can nest related properties and use lists naturally in YAML. Properties files use flat key-value pairs. I choose YAML for complex configurations with nested structures, and properties for simpler configurations. Spring Boot loads both automatically, so I can use either or even both together."

---

### Question 471: What are Spring Boot profiles and how do they work?

**Answer:**
(Duplicate of 461).

---

### Question 472: How do you activate multiple profiles simultaneously?

**Answer:**
(Duplicate of 462).

---

### Question 473: How do you deploy a Spring Boot app as a WAR file?

**Answer:**
(Duplicate of 463).

---

### Question 474: What is the difference between embedded Tomcat and external Tomcat deployment?

**Answer:**
(Duplicate of 464).

---

### Question 475: How do you deploy Spring Boot apps on Heroku?

**Answer:**
(Duplicate of 465).

---

### Question 476: What is `spring.profiles.active` and where can you define it?

**Answer:**
(Duplicate of 466).

---

### Question 477: How do you implement zero-downtime deployment for Spring Boot apps?

**Answer:**
(Duplicate of 467).

---

### Question 478: How do you enable graceful shutdown in Spring Boot?

**Answer:**
(Duplicate of 468).

---

### Question 479: How do you deploy Spring Boot apps to AWS Lambda?

**Answer:**
(Duplicate of 469).

---

### Question 480: What is the role of `application.properties` vs `application.yml`?

**Answer:**
(Duplicate of 470).

---
