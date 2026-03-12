## 🔹 Section 6: Testing Advanced (201-220)

### Question 201: How to write parameterized tests in Spring Boot?

**Answer:**
Use JUnit 5 `@ParameterizedTest`.
Providers: `@ValueSource(strings = {"a", "b"})`, `@CsvSource`.
Useful for validation logic testing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to write parameterized tests in Spring Boot?
**Your Response:** "I write parameterized tests using JUnit 5's `@ParameterizedTest` annotation. I use providers like `@ValueSource(strings = {'a', 'b'})` for simple values or `@CsvSource` for multiple parameters. This allows me to run the same test method with different inputs, which is perfect for testing validation logic or edge cases. Instead of writing multiple similar test methods, I can write one parameterized test that covers all scenarios. This approach reduces code duplication and makes it easier to add new test cases by just adding new parameter values."

---

### Question 202: How do you test database interactions using TestContainers?

**Answer:**
Library that spins up **Real Docker Containers** (Postgres/Redis) for tests.
Add dependency.
Annotate test class `@Testcontainers`.
Container `PostgreSQLContainer<?> postgres = ...`.
Much better than using H2 for Integration Tests as it tests DB-specific syntax correctness.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test database interactions using TestContainers?
**Your Response:** "TestContainers is a library that spins up real Docker containers like PostgreSQL or Redis for integration tests. I add the dependency, annotate my test class with `@Testcontainers`, and define containers like `PostgreSQLContainer`. This is much better than using H2 because it tests against the actual database I'll use in production, catching DB-specific syntax issues early. The containers are automatically started before tests and cleaned up afterward, giving me isolated, realistic test environments without manual setup. This approach gives me confidence that my database code works in production-like conditions."

---

### Question 203: What is the purpose of `@Transactional` in Spring Boot tests?

**Answer:**
Test methods annotated with `@Transactional` are **rolled back** automatically after execution.
This keeps the DB clean for the next test.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `@Transactional` in Spring Boot tests?
**Your Response:** "When I annotate test methods with `@Transactional`, Spring automatically rolls back the transaction after the test completes. This keeps the database clean for the next test, preventing test data from interfering with each other. This is particularly useful for integration tests that create or modify data - I can test my repository logic without worrying about cleaning up the data afterward. The rollback happens automatically, so my tests remain isolated and repeatable. This approach is much cleaner than manual cleanup and ensures that each test starts with a clean slate."

---

### Question 204: How do you mock service layers in a controller test?

**Answer:**
Use `@MockBean`.
`@WebMvcTest(UserController.class)`
`@MockBean private UserService userService;`
`given(userService.find(1)).willReturn(user);`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you mock service layers in a controller test?
**Your Response:** "I mock service layers in controller tests using `@MockBean`. I use `@WebMvcTest(UserController.class)` to load only the web layer and then add `@MockBean private UserService userService` to replace the real service with a mock. I use Mockito's `given(userService.find(1)).willReturn(user)` to define what the mock should return. This allows me to test my controller logic in isolation without depending on the actual service implementation. The mocks are automatically injected into the controller, making the setup clean and the tests fast since they don't need to spin up the full application context."

---

### Question 205: What are embedded databases and how are they used for testing?

**Answer:**
H2, HSQLDB, Derby.
Fast, in-memory.
Spring Boot Auto-configuration detects them on test classpath and replaces the DataSource with them automatically (unless `@AutoConfigureTestDatabase(replace=NONE)` is used).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are embedded databases and how are they used for testing?
**Your Response:** "Embedded databases like H2, HSQLDB, and Derby are fast, in-memory databases that run within the application process. Spring Boot's auto-configuration automatically detects them on the test classpath and replaces the DataSource with an embedded database, unless I specify `@AutoConfigureTestDatabase(replace=NONE)`. This makes testing incredibly fast since there's no external database setup required. The embedded database starts with the application and stops with it, providing a clean, isolated test environment. While H2 is great for speed, I use TestContainers when I need to test against the actual production database."

---

### Question 206: How do you test exception handling in controllers?

**Answer:**
MockMvc:
`.andExpect(status().isNotFound())`
`.andExpect(jsonPath("$.errorCode").value(404))`
Verifies that `@ControllerAdvice` effectively caught the exception and formatted the response.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test exception handling in controllers?
**Your Response:** "I test exception handling in controllers using MockMvc. I write tests that expect specific HTTP status codes like `.andExpect(status().isNotFound())` and verify the error response structure with `.andExpect(jsonPath('$.errorCode').value(404))`. This approach verifies that my `@ControllerAdvice` exception handler is working correctly and formatting error responses consistently. I can test different exception scenarios to ensure that all error cases are handled properly and that clients receive meaningful, structured error messages rather than stack traces."

---

### Question 207: How do you test configuration properties in Spring Boot?

**Answer:**
Use `@TestPropertySource` or `@SpringBootTest(properties = "app.val=test")`.
Inject the `@ConfigurationProperties` bean and assert values.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test configuration properties in Spring Boot?
**Your Response:** "I test configuration properties using `@TestPropertySource` or `@SpringBootTest(properties = 'app.val=test')` to override properties for the test. I inject the `@ConfigurationProperties` bean and assert that the values are correctly bound from the test properties. This approach allows me to test my configuration binding logic and validation without affecting the actual application.properties. I can test different property combinations and ensure that my configuration classes work correctly with various inputs, including edge cases and invalid values."

---

### Question 208: What is the difference between `@MockBean` and `@SpyBean`?

**Answer:**
- **`@MockBean`:** Complete mock. Calls do nothing unless stubbed.
- **`@SpyBean`:** Wraps real bean. Calls real methods unless specific methods are stubbed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@MockBean` and `@SpyBean`?
**Your Response:** "`@MockBean` creates a complete mock where all method calls do nothing unless I explicitly stub them. `@SpyBean` wraps the real bean and calls the actual methods by default, only overriding specific methods when I stub them. I use `@MockBean` when I want to completely isolate a component, and `@SpyBean` when I want to test the real behavior but override certain methods. `@SpyBean` is useful for testing partial functionality while keeping most of the original implementation intact. The choice depends on whether I want complete isolation or partial real behavior."

---

### Question 209: How to isolate integration tests using test profiles?

**Answer:**
`@ActiveProfiles("test")`.
Creates `application-test.properties`.
Configure specific DB URL or disable async processing there.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to isolate integration tests using test profiles?
**Your Response:** "I isolate integration tests using test profiles by annotating my test class with `@ActiveProfiles('test')`. This loads `application-test.properties` where I can configure test-specific settings like a different database URL or disable asynchronous processing. Test profiles allow me to create a completely different configuration for testing without affecting the main application. This is essential for integration tests where I need specific test data, external service mocks, or different behavior than production. The test profile ensures that tests run in a controlled, predictable environment."

---

### Question 210: How to use Postman/Newman for Spring Boot API testing?

**Answer:**
External testing.
Export Postman Collection.
Run in CI using `newman run collection.json`.
Asserts HTTP Status codes and Body schema from the "outside".

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to use Postman/Newman for Spring Boot API testing?
**Your Response:** "I use Postman for external API testing and Newman for automation in CI pipelines. I create comprehensive Postman collections that test all my API endpoints with different scenarios. Then I export the collection and run `newman run collection.json` in my CI pipeline to automate API testing from the outside. This approach validates HTTP status codes, response schemas, and actual behavior as a client would experience it. It complements my unit and integration tests by testing the complete request-response cycle and ensuring the API works as documented."

---

### Question 211: What is the role of `@TestConfiguration` in Spring Boot?

**Answer:**
Used to define extra beans for tests only.
Unlike standard `@Configuration`, it is NOT picked up by scanning automatically.
You must import it: `@Import(MyTestConfig.class)`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `@TestConfiguration` in Spring Boot?
**Your Response:** "`@TestConfiguration` allows me to define extra beans specifically for tests. Unlike standard `@Configuration`, it's not picked up by component scanning automatically - I must explicitly import it with `@Import(MyTestConfig.class)`. This is perfect for defining test-specific beans like mocks, test data loaders, or alternative implementations that I only want in tests. It keeps my test configuration separate from production configuration while still allowing me to customize the test context. This approach ensures that test-specific beans don't accidentally end up in production."

---

### Question 212: How to run integration tests with a specific profile?

**Answer:**
(Duplicate of 209).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to run integration tests with a specific profile?
**Your Response:** "I run integration tests with specific profiles using `@ActiveProfiles('test')` to load test-specific configuration. This allows me to use `application-test.properties` with settings tailored for integration testing, like using an in-memory database, disabling certain features, or configuring test-specific beans. The test profile ensures that my integration tests run in a controlled environment that's isolated from production settings. I can create multiple test profiles for different testing scenarios - like `integration-test` vs `performance-test` - each with their own configuration."

---

### Question 213: How do you use `MockMvc` vs `WebTestClient`?

**Answer:**
(See Q60 / Q112). `WebTestClient` allows testing asynchronous/streaming endpoints.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `MockMvc` vs `WebTestClient`?
**Your Response:** "I use `MockMvc` for testing traditional Spring MVC controllers with synchronous requests. For testing reactive applications with WebFlux, I use `WebTestClient` which can handle asynchronous and streaming endpoints. `WebTestClient` works with both MVC and WebFlux, while `MockMvc` only works with MVC. `WebTestClient` provides a more modern, fluent API and can test reactive streams, server-sent events, and other async patterns. I choose `MockMvc` for traditional servlet-based apps and `WebTestClient` when working with reactive applications or when I need to test streaming responses."

---

### Question 214: What is the difference between `@WebMvcTest` and `@SpringBootTest`?

**Answer:**
(See Q91/Q88). Slice vs Full Context.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@WebMvcTest` and `@SpringBootTest`?
**Your Response:** "`@WebMvcTest` is a slice test that only loads the web layer - controllers, filters, and related components - making it fast for testing controller logic in isolation. `@SpringBootTest` loads the full application context, which is slower but tests the complete integration. I use `@WebMvcTest` when I want to test just the web layer with mocked dependencies, and `@SpringBootTest` when I need to test the complete application including database integration, messaging, and other components. The choice depends on whether I want fast, isolated tests or comprehensive integration tests."

---

### Question 215: How do you mock configuration properties in tests?

**Answer:**
(See Q207).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you mock configuration properties in tests?
**Your Response:** "I mock configuration properties in tests using `@TestPropertySource` to provide test-specific property files or `@SpringBootTest(properties = 'app.val=test')` to override individual properties. This allows me to test how my application behaves with different configuration values without modifying the actual application.properties. I can test edge cases, default values, and validation logic by providing controlled property values in my tests. This approach ensures that my configuration binding and validation work correctly across different scenarios."

---

### Question 216: How to write integration tests using H2 in-memory DB?

**Answer:**
(See Q80). `@DataJpaTest`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to write integration tests using H2 in-memory DB?
**Your Response:** "I write integration tests with H2 using `@DataJpaTest`, which automatically configures an in-memory H2 database for testing JPA repositories. This annotation loads only the data access layer - entities, repositories, and database configuration - making the tests fast and focused. H2 provides a clean, isolated database for each test, and Spring Boot automatically rolls back transactions after each test to keep the data clean. This approach is perfect for testing repository logic, queries, and entity relationships without needing an external database setup."

---

### Question 217: How to test scheduled jobs in Spring Boot?

**Answer:**
Wait libraries like Awaitility.
Or better: Extract logic to a Service. Test Service.
Integration: Set cron to run every second in test profile, use `@SpyBean` to verify invocation count.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to test scheduled jobs in Spring Boot?
**Your Response:** "Testing scheduled jobs can be tricky. I prefer extracting the job logic to a service and testing the service directly. For integration testing, I configure the cron expression to run every second in my test profile and use `@SpyBean` to verify that the method was called the expected number of times. I can also use wait libraries like Awaitility to wait for the scheduled execution. The key is to separate the scheduling concern from the business logic so I can test the logic independently of the scheduling mechanism."

---

### Question 218: How do you test Spring Events in Spring Boot?

**Answer:**
`@RecordApplicationEvents`.
Inject `ApplicationEvents`.
Check `events.stream(MyEvent.class).count()`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test Spring Events in Spring Boot?
**Your Response:** "I test Spring Events using `@RecordApplicationEvents` which records all application events during the test. I inject the `ApplicationEvents` object and then verify that specific events were published using `events.stream(MyEvent.class).count()`. This approach allows me to test that my event publishing logic works correctly and that the expected events are fired under the right conditions. I can also inspect the event contents to ensure they contain the correct data. This is much cleaner than trying to capture events manually and provides reliable event testing."

---

### Question 219: How do you write parameterized integration tests?

**Answer:**
Combine `@SpringBootTest` with JUnit5 `@ParameterizedTest`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write parameterized integration tests?
**Your Response:** "I write parameterized integration tests by combining `@SpringBootTest` with JUnit 5's `@ParameterizedTest`. This allows me to run full integration tests with different input parameters. I use `@CsvSource` or `@MethodSource` to provide the test data, and the test runs with the complete Spring context for each parameter set. This is useful for testing integration scenarios with different data variations, like testing how my API handles different input values or how my database operations work with different entities. It combines the realism of integration tests with the efficiency of parameterized testing."

---

### Question 220: What are TestContainers and how do you use them with Spring Boot?

**Answer:**
(See Q202).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are TestContainers and how do you use them with Spring Boot?
**Your Response:** "TestContainers is a library that provides lightweight, disposable Docker containers for testing. I use it with Spring Boot to spin up real dependencies like PostgreSQL, Redis, or Kafka for my integration tests. I add the `@Testcontainers` annotation and define the containers I need. Spring Boot automatically configures itself to use these containers instead of mock or in-memory alternatives. This gives me realistic integration tests that run against actual services, catching issues that might only appear with the real dependencies. The containers are automatically managed, so I don't need to worry about setup or cleanup."

---
