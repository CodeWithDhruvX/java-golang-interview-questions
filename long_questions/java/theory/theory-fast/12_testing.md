# Testing Interview Questions (134-140)

## Testing Strategies

### 134. Difference between unit, integration, and end-to-end tests?
"**Unit Tests** test a single class or method in isolation. We mock all external dependencies. They are fast (milliseconds) and give precise feedback. 'The logic in `calculateTax()` is wrong.'

**Integration Tests** check how components work together. 'Does the Service correctly fetch data from the Database?' This requires a real (or in-memory) DB but mocks external APIs.

**End-to-End (E2E) Tests** test the full flow from user perspective. 'Open browser -> Login -> Add to Cart -> Checkout'. They are slow, flaky, and hard to debug, but they prove that system actually works."

**Spoken Format:**
"Testing strategies are like having different types of quality control checkpoints.

**Unit Tests** are like testing individual car parts in isolation - you test the engine, the brakes, the steering wheel separately. If the engine fails, you know exactly which part is broken.

**Integration Tests** are like testing how parts work together - you test if the engine and transmission work together properly. This catches issues that individual tests might miss.

**End-to-End Tests** are like test-driving the entire car - you simulate the complete user experience from starting the car to reaching the destination.

Each type serves different purposes:
- Unit tests find specific bugs quickly
- Integration tests catch interaction issues
- E2E tests verify the whole system works as expected

You need all three to ensure your car (application) is safe and reliable!"

### 135. What is mocking and when should you avoid it?
"Mocking is creating a fake version of a dependency (using Mockito) to isolate the code under test.

I use it heavily in **Unit Tests**. If I test `UserService`, I mock `UserRepository` because I don't want to hit the DB.

However, I avoid mocking in **Integration Tests**. If I mock the database there, I'm not really testing the integration. I'm testing my assumption of how the DB works. Over-mocking leads to tests that pass but a system that fails in production."

**Spoken Format:**
"Mocking is like using a stunt double in a movie instead of the real actor.

In **Unit Tests**, mocking is perfect - you're testing if the director (your code) can work with the stunt double (mocked dependency). You don't need the real actor for this test.

In **Integration Tests**, mocking is dangerous - you're testing if the director can work with a stunt double, but the audience expects the real actor. If the real actor behaves differently, your test passes but the movie fails.

The key insight: Integration tests should use real dependencies to test actual interactions, not your assumptions about how they work.

Mock in unit tests to isolate code, but use real components in integration tests to catch real-world issues!"

### 136. How do you test REST controllers?
"I use `@WebMvcTest`. It spins up a slice of the Spring context containing just the Controller layer. It mocks the Service layer.

I use `MockMvc` to perform requests: `.perform(get("/users/1"))`.
Then I assert the response: `.andExpect(status().isOk())` and `.andExpect(jsonPath("$.name").value("John"))`.

This tests routing, JSON serialization, and validation logic without starting the whole server."

**Spoken Format:**
"Testing REST controllers is like testing just the front desk of a hotel.

`@WebMvcTest` is like setting up a mock front desk environment - you have the receptionist (controller) and phone system (MockMvc), but you don't need the entire hotel running.

You test specific interactions:
- When guest asks for room key (HTTP GET request)
- When guest checks in (HTTP POST request)
- When guest asks for bill (HTTP response format)

This is much faster than starting the entire hotel because you're only testing the front desk operations, not housekeeping, restaurant, or maintenance.

The beauty is that you can test the user-facing interface without the overhead of full application startup."

### 137. What is `@SpringBootTest` vs `@WebMvcTest`?
"`@SpringBootTest` loads the **entire** application context. It scans every bean, connects to the DB (or embedded DB), and starts the server. It’s heavy and slow but necessary for full integration testing.

`@WebMvcTest` is a **slice test**. It only loads beans related to the web layer (Controllers, ControllerAdvice, JsonComponents). It does *not* load Services or Repositories. It’s much faster and focused on testing the HTTP interface."

**Spoken Format:**
"`@SpringBootTest` is like testing the entire hotel - you're checking that every department works together seamlessly.

`@WebMvcTest` is like testing just the front desk - you're focusing on the user-facing interface and ensuring it works correctly.

Both are necessary:
- `@SpringBootTest` ensures the whole system works together
- `@WebMvcTest` ensures the user interface is correct and fast

Use `@SpringBootTest` for full integration testing and `@WebMvcTest` for focused web layer testing."

### 138. How do you test database interactions?
"I use `@DataJpaTest`. This is another slice test.

It configures an in-memory database (H2), scans for Entities and Repositories, and configures Hibernate.

By default, it is **transactional and rolls back** at the end of each test. This means I can save a user, assert it exists, and the test finishes without leaving junk data for the next test.

For more realistic testing, I use **Testcontainers** to spin up a *real* Dockerized PostgreSQL instead of H2, avoiding 'works in H2, fails in Postgres' issues."

**Spoken Format:**
"Testing database interactions is like making sure your recipe works with real ingredients.

`@DataJpaTest` is like having a test kitchen with a miniature stove and ingredients:

- **H2 Database** is like using toy ingredients - they work, but they're not exactly the same as real ingredients
- **Testcontainers** is like having a real stove and actual ingredients from your pantry

The problem with H2 is that it's a different database than production. Your recipe might work perfectly with toy ingredients but fail with real ones.

Testcontainers solves this by:
- Spinning up actual PostgreSQL in Docker
- Using real database schema and constraints
- Testing with production-like environment

This catches issues that H2 would miss, like SQL syntax differences or constraint violations. It's like testing your recipe with the same ingredients you'll use in production!"

### 139. What is test coverage and why it can be misleading?
"Test coverage measures the percentage of code lines executed during tests.

It’s misleading because executing a line doesn't mean testing the logic. I can write a test that calls every method but asserts nothing. `userService.doEverything(); // 100% coverage!`

High coverage is necessary but not sufficient. I focus on **branch coverage** (testing all if/else paths) and meaningful assertions over raw line percentage."

**Spoken Format:**
"Test coverage is like having a map that shows which roads you've driven on, but it doesn't tell you if you drove safely.

High line coverage can be misleading because:
- You can call every method but not test the actual logic
- You can test every branch but not test error conditions
- You can have 100% coverage but still miss critical bugs

**Meaningful testing** is like actually testing if the brakes work, not just that you pressed the brake pedal.

I focus on:
- **Branch coverage** - testing all possible paths through your code
- **Edge cases** - testing what happens with empty inputs, null values, etc.
- **Business logic** - testing the actual rules and requirements

The goal isn't just to drive every road, but to drive safely and reach the correct destination!"

### 140. What makes a good test?
"A good test is **Deterministic**: It always passes or always fails. No flakiness.
It’s **Isolated**: It doesn't depend on the order or state left by other tests.
It’s **Readable**: It acts as documentation. 'shouldThrowExceptionWhenUserFound'.
And crucially, it tests **behavior, not implementation**. If I refactor to private methods inside a class, the test should still pass."

**Spoken Format:**
"A good test is like having a reliable quality control inspector.

**Deterministic** means the inspector always gives the same result for the same product - no random failures or flaky behavior.

**Isolated** means each product is tested independently - one product's failure doesn't affect the test of another product.

**Readable** means the test acts as documentation - someone can read the test and understand exactly what the product should do.

**Testing behavior, not implementation** is like testing that a car stops when you press the brakes, not testing that the brakes use hydraulic fluid. If you change from hydraulic to electric brakes, the test should still pass because the behavior (stopping) is the same.

This makes your tests robust and maintainable - they focus on what users care about (what the system does) rather than how it does it internally!"
