# Testing Interview Questions (134-140)

## Testing Strategies

### 134. Difference between unit, integration, and end-to-end tests?
"**Unit Tests** test a single class or method in isolation. We mock all external dependencies. They are fast (milliseconds) and give precise feedback. 'The logic in `calculateTax()` is wrong.'

**Integration Tests** check how components work together. 'Does the Service correctly fetch data from the Database?' This requires a real (or in-memory) DB but mocks external APIs.

**End-to-End (E2E) Tests** test the full flow from user perspective. 'Open browser -> Login -> Add to Cart -> Checkout'. They are slow, flaky, and hard to debug, but they prove the system actually works."

### 135. What is mocking and when should you avoid it?
"Mocking is creating a fake version of a dependency (using Mockito) to isolate the code under test.

I use it heavily in **Unit Tests**. If I test `UserService`, I mock `UserRepository` because I don't want to hit the DB.

However, I avoid mocking in **Integration Tests**. If I mock the database there, I’m not really testing the integration. I’m testing my assumption of how the DB works. Over-mocking leads to tests that pass but a system that fails in production."

### 136. How do you test REST controllers?
"I use `@WebMvcTest`. It spins up a slice of the Spring context containing just the Controller layer. It mocks the Service layer.

I use `MockMvc` to perform requests: `.perform(get("/users/1"))`.
Then I assert the response: `.andExpect(status().isOk())` and `.andExpect(jsonPath("$.name").value("John"))`.

This tests the routing, JSON serialization, and validation logic without starting the whole server."

### 137. What is `@SpringBootTest` vs `@WebMvcTest`?
"`@SpringBootTest` loads the **entire** application context. It scans every bean, connects to the DB (or embedded DB), and starts the server. It’s heavy and slow but necessary for full integration testing.

`@WebMvcTest` is a **slice test**. It only loads beans related to the web layer (Controllers, ControllerAdvice, JsonComponents). It does *not* load Services or Repositories. It’s much faster and focused on testing the HTTP interface."

### 138. How do you test database interactions?
"I use `@DataJpaTest`. This is another slice test.

It configures an in-memory database (H2), scans for Entities and Repositories, and configures Hibernate.

By default, it is **transactional and rolls back** at the end of each test. This means I can save a user, assert it exists, and the test finishes without leaving junk data for the next test.

For more realistic testing, I use **Testcontainers** to spin up a *real* Dockerized PostgreSQL instead of H2, avoiding 'works in H2, fails in Postgres' issues."

### 139. What is test coverage and why it can be misleading?
"Test coverage measures the percentage of code lines executed during tests.

It’s misleading because executing a line doesn't mean testing the logic. I can write a test that calls every method but asserts nothing. `userService.doEverything(); // 100% coverage!`

High coverage is necessary but not sufficient. I focus on **branch coverage** (testing all if/else paths) and meaningful assertions over raw line percentage."

### 140. What makes a good test?
"A good test is **Deterministic**: It always passes or always fails. No flakiness.
It’s **Isolated**: It doesn't depend on the order or state left by other tests.
It’s **Readable**: It acts as documentation. 'shouldThrowExceptionWhenUserFound'.
And crucially, it tests **behavior, not implementation**. If I refactor the private methods inside a class, the test should still pass."
