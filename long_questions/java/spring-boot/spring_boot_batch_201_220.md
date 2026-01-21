## 🔹 Section 6: Testing Advanced (201-220)

### Question 201: How to write parameterized tests in Spring Boot?

**Answer:**
Use JUnit 5 `@ParameterizedTest`.
Providers: `@ValueSource(strings = {"a", "b"})`, `@CsvSource`.
Useful for validation logic testing.

---

### Question 202: How do you test database interactions using TestContainers?

**Answer:**
Library that spins up **Real Docker Containers** (Postgres/Redis) for tests.
Add dependency.
Annotate test class `@Testcontainers`.
Container `PostgreSQLContainer<?> postgres = ...`.
Much better than using H2 for Integration Tests as it tests DB-specific syntax correctness.

---

### Question 203: What is the purpose of `@Transactional` in Spring Boot tests?

**Answer:**
Test methods annotated with `@Transactional` are **rolled back** automatically after execution.
This keeps the DB clean for the next test.

---

### Question 204: How do you mock service layers in a controller test?

**Answer:**
Use `@MockBean`.
`@WebMvcTest(UserController.class)`
`@MockBean private UserService userService;`
`given(userService.find(1)).willReturn(user);`

---

### Question 205: What are embedded databases and how are they used for testing?

**Answer:**
H2, HSQLDB, Derby.
Fast, in-memory.
Spring Boot Auto-configuration detects them on test classpath and replaces the DataSource with them automatically (unless `@AutoConfigureTestDatabase(replace=NONE)` is used).

---

### Question 206: How do you test exception handling in controllers?

**Answer:**
MockMvc:
`.andExpect(status().isNotFound())`
`.andExpect(jsonPath("$.errorCode").value(404))`
Verifies that `@ControllerAdvice` effectively caught the exception and formatted the response.

---

### Question 207: How do you test configuration properties in Spring Boot?

**Answer:**
Use `@TestPropertySource` or `@SpringBootTest(properties = "app.val=test")`.
Inject the `@ConfigurationProperties` bean and assert values.

---

### Question 208: What is the difference between `@MockBean` and `@SpyBean`?

**Answer:**
- **`@MockBean`:** Complete mock. Calls do nothing unless stubbed.
- **`@SpyBean`:** Wraps real bean. Calls real methods unless specific methods are stubbed.

---

### Question 209: How to isolate integration tests using test profiles?

**Answer:**
`@ActiveProfiles("test")`.
Creates `application-test.properties`.
Configure specific DB URL or disable async processing there.

---

### Question 210: How to use Postman/Newman for Spring Boot API testing?

**Answer:**
External testing.
Export Postman Collection.
Run in CI using `newman run collection.json`.
Asserts HTTP Status codes and Body schema from the "outside".

---

### Question 211: What is the role of `@TestConfiguration` in Spring Boot?

**Answer:**
Used to define extra beans for tests only.
Unlike standard `@Configuration`, it is NOT picked up by scanning automatically.
You must import it: `@Import(MyTestConfig.class)`.

---

### Question 212: How to run integration tests with a specific profile?

**Answer:**
(Duplicate of 209).

---

### Question 213: How do you use `MockMvc` vs `WebTestClient`?

**Answer:**
(See Q60 / Q112). `WebTestClient` allows testing asynchronous/streaming endpoints.

---

### Question 214: What is the difference between `@WebMvcTest` and `@SpringBootTest`?

**Answer:**
(See Q91/Q88). Slice vs Full Context.

---

### Question 215: How do you mock configuration properties in tests?

**Answer:**
(See Q207).

---

### Question 216: How to write integration tests using H2 in-memory DB?

**Answer:**
(See Q80). `@DataJpaTest`.

---

### Question 217: How to test scheduled jobs in Spring Boot?

**Answer:**
Wait libraries like Awaitility.
Or better: Extract logic to a Service. Test Service.
Integration: Set cron to run every second in test profile, use `@SpyBean` to verify invocation count.

---

### Question 218: How do you test Spring Events in Spring Boot?

**Answer:**
`@RecordApplicationEvents`.
Inject `ApplicationEvents`.
Check `events.stream(MyEvent.class).count()`.

---

### Question 219: How do you write parameterized integration tests?

**Answer:**
Combine `@SpringBootTest` with JUnit5 `@ParameterizedTest`.

---

### Question 220: What are TestContainers and how do you use them with Spring Boot?

**Answer:**
(See Q202).

---
