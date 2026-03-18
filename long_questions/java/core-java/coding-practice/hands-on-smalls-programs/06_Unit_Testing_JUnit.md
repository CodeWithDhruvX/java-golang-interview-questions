# Specialized Topic 3: Unit Testing (JUnit)

**Goal**: Learn to write test cases to verify code correctness.

**Prerequisites**: Add `junit-jupiter-api` and `junit-jupiter-engine` dependencies (Maven/Gradle) or add JARs to classpath.

## 1. Class to Test (Calculator)

```java
public class Calculator {
    public int add(int a, int b) {
        return a + b;
    }
    
    public int divide(int a, int b) {
        if (b == 0) throw new ArithmeticException("Zero Division");
        return a / b;
    }
}
```

## 2. JUnit Test Class

```java
import org.junit.jupiter.api.*;
import static org.junit.jupiter.api.Assertions.*;

public class CalculatorTest {

    Calculator calc;

    // Runs before EACH test
    @BeforeEach
    void setup() {
        calc = new Calculator();
    }

    @Test
    void testAdd() {
        int result = calc.add(5, 3);
        // assert(Expected, Actual)
        assertEquals(8, result, "5 + 3 should be 8");
    }

    @Test
    void testDivide() {
        assertEquals(2, calc.divide(10, 5));
    }

    @Test
    void testDivideByZero() {
        // Assert that an exception is thrown
        assertThrows(ArithmeticException.class, () -> {
            calc.divide(10, 0);
        });
    }
    
    @AfterEach
    void tearDown() {
        // Clean up resources if any
    }
}
```

## Key Annotations
*   `@Test`: Marks a method as a test case.
*   `@BeforeEach`: Runs before every test method (setup).
*   `@AfterEach`: Runs after every test method (cleanup).
*   `@BeforeAll` / `@AfterAll`: Runs once per class (static).

## Assertions
*   `assertEquals(expected, actual)`: Checks equality.
*   `assertTrue(condition)` / `assertFalse(condition)`.
*   `assertThrows(Exception.class, executable)`: Verifies exception handling.
*   `assertNotNull(object)`.

---

## 📋 Comprehensive Interview Questions

### **JUnit Fundamentals & Annotations Questions**

**Q1: Explain the JUnit test lifecycle and execution order.**
**A**: "JUnit follows a specific lifecycle: For each test method, it creates a new test class instance, then runs @BeforeAll methods (once per class), @BeforeEach methods (before each test), the @Test method itself, @AfterEach methods (after each test), and finally @AfterAll methods (once per class). This isolation ensures tests don't interfere with each other. The @BeforeAll/@AfterAll methods must be static since they run before any instance exists. This lifecycle provides clean setup/teardown for both individual tests and the entire test suite."

**Q2: What's the difference between @BeforeEach and @BeforeAll?**
**A**: "@BeforeEach runs before every single test method, ensuring each test starts with a fresh, clean state. It's perfect for creating new objects, resetting shared resources, or any setup that needs to be repeated for each test. @BeforeAll runs only once before all tests in the class, making it ideal for expensive operations like database connections, starting servers, or loading large datasets. The key difference is frequency - per-test vs per-class setup."

**Q3: How does JUnit 5 differ from JUnit 4?**
**A**: "JUnit 5 brought significant improvements: It uses Java 8+ features like lambdas and streams, has a more modular architecture with separate APIs for different concerns, uses @Test instead of @Test annotation from org.junit, @DisplayName for readable test names, parameterized tests are built-in, and it supports dynamic tests. JUnit 5 also provides better exception testing with assertThrows, improved assumptions, and more flexible extensions instead of runners. The migration is worthwhile for modern Java applications."

### **Assertions & Testing Strategies Questions**

**Q4: What are the different types of assertions in JUnit?**
**A**: "JUnit provides various assertions: `assertEquals()` for equality checks with optional delta for floating-point numbers, `assertTrue()`/`assertFalse()` for boolean conditions, `assertNull()`/`assertNotNull()` for null checks, `assertSame()`/`assertNotSame()` for object identity vs equality, `assertThrows()` for exception testing, `assertTimeout()` for performance testing, and `assertAll()` for grouped assertions. Each assertion can include a custom message that displays when the test fails, making debugging easier."

**Q5: How do you test exceptions in JUnit 5?**
**A**: "In JUnit 5, I use `assertThrows(ExpectedExceptionType, executable)` which is much cleaner than JUnit 4's @Test(expected=...). I pass a lambda expression containing the code that should throw the exception. The method returns the actual exception, allowing me to further assert on the exception message or cause. For example: `Exception exception = assertThrows(IllegalArgumentException.class, () -> calculator.divide(1, 0)); assertEquals("Cannot divide by zero", exception.getMessage());`"

**Q6: What is the difference between assertEquals and assertSame?**
**A**: "`assertEquals()` uses the equals() method to check if two objects have the same value or content. `assertSame()` uses == to check if two references point to the exact same object in memory. For primitives, both work the same way. For objects, assertEquals checks logical equality while assertSame checks identity. I use assertSame when testing singletons, caching, or when I need to verify that the same object instance is returned, and assertEquals for value comparisons."

### **Advanced Testing Concepts Questions**

**Q7: What are parameterized tests and when would you use them?**
**A**: "Parameterized tests allow me to run the same test logic with multiple different inputs. I use @ParameterizedTest along with sources like @ValueSource, @EnumSource, @MethodSource, or @CsvSource. They're perfect for testing edge cases, boundary conditions, or when I have a method that should behave consistently across different inputs. Instead of writing multiple similar test methods, I can write one parameterized test that covers many scenarios, making tests more maintainable and comprehensive."

**Q8: How do you test private methods in Java?**
**A**: "Generally, I don't test private methods directly - I test them indirectly through public methods that use them. This follows the principle of testing the class's public contract. However, if I must test private methods, I have several options: 1) Use reflection to make the method accessible, 2) Change the method visibility to package-private and test in the same package, 3) Use nested test classes that can access private members, or 4) Extract the private logic to a separate class and test that class. The best approach is usually to refactor rather than break encapsulation."

**Q9: What is test-driven development (TDD) and how does it work?**
**A**: "TDD is a development approach where I write tests before writing the production code. The cycle is: Red (write a failing test), Green (write minimal code to make the test pass), Refactor (improve the code while keeping tests green). This approach ensures comprehensive test coverage, drives better design, and provides immediate feedback. TDD helps me think about requirements and edge cases before implementation, leading to cleaner, more testable code. It's a discipline that improves both code quality and developer confidence."

### **Mocking & Test Isolation Questions**

**Q10: What is mocking and when do you use it?**
**A**: "Mocking creates fake implementations of dependencies to isolate the class under test. I use mocking when dependencies are expensive, slow, unreliable, or have side effects - like database connections, web services, or file systems. Mocks allow me to simulate different scenarios, verify interactions, and control the behavior of dependencies. Popular frameworks include Mockito and EasyMock. The key is to mock only external dependencies, not the class I'm actually testing."

**Q11: Explain the difference between mock, stub, and spy.**
**A**: "A stub is a simple fake that returns predefined responses - it's stateless and doesn't verify interactions. A mock is more sophisticated - it can verify method calls, arguments, and call counts. A spy is a partial mock that wraps a real object, allowing me to override specific methods while keeping the original behavior for others. I use stubs for simple scenarios, mocks when I need to verify interactions, and spies when I want to test most behavior but override specific parts."

**Q12: How do you handle database testing?**
**A**: "For database testing, I use several strategies: 1) In-memory databases like H2 for fast, isolated tests, 2) Testcontainers for real database instances in Docker, 3) Database migrations with Flyway/Liquibase to ensure consistent schema, 4) Transactional tests that roll back changes after each test, 5) Separate test database to avoid polluting production data. I also use repository pattern to make database code easier to mock for unit tests, and integration tests to verify actual database interactions."

### **Test Organization & Best Practices Questions**

**Q13: What makes a good unit test?**
**A**: "A good unit test should be FAST - runs quickly to encourage frequent execution, ISOLATED - independent of other tests and external dependencies, REPEATABLE - produces same results every time, SELF-VALIDATING - has automatic assertions, and TIMELY - written at the right time. I also follow the FIRST principles: Fast, Independent, Repeatable, Self-validating, and Thorough. Good tests have clear names, test one specific behavior, use descriptive assertions, and are easy to understand and maintain."

**Q14: How do you organize test packages and classes?**
**A**: "I typically mirror the production package structure in test packages - so com.example.service would have corresponding tests in com.example.service.test. Test classes are usually named by appending 'Test' to the class being tested. I organize tests by the class they test, not by functionality. For integration tests, I might create separate test packages or use different naming conventions. The key is making tests easy to find and understand their relationship to production code."

**Q15: What is code coverage and what are its limitations?**
**A**: "Code coverage measures how much of the production code is executed by tests - usually line, branch, or path coverage. High coverage gives confidence but doesn't guarantee quality - tests could have 100% coverage but miss edge cases or have wrong assertions. Coverage is a useful metric for identifying untested code, but it's not a substitute for thoughtful test design. I aim for meaningful coverage of critical paths rather than chasing arbitrary percentage targets. Quality of tests matters more than quantity."

### **Integration & System Testing Questions**

**Q16: What's the difference between unit tests and integration tests?**
**A**: "Unit tests test individual components in isolation, using mocks for dependencies. They're fast, numerous, and focus on specific functionality. Integration tests test how components work together - they might use real databases, web services, or file systems. Integration tests are slower but catch issues that unit tests miss, like configuration problems or interface mismatches. A good test suite has both - fast unit tests for immediate feedback and integration tests for confidence in the overall system."

**Q17: How do you test REST APIs?**
**A**: "For REST API testing, I use tools like MockMvc for Spring applications or libraries like RestAssured. I test various aspects: HTTP status codes, response bodies, headers, content types, and error scenarios. I test both happy path and edge cases like invalid inputs, missing required fields, and authentication failures. For comprehensive testing, I also test API contracts, performance under load, and integration with dependent services. Tests should be independent and not rely on external services being available."

**Q18: What are end-to-end tests and when are they necessary?**
**A**: "End-to-end tests simulate real user scenarios by testing the entire application stack - UI, business logic, database, and external integrations. They're slower and more brittle than unit tests but catch integration issues that other tests miss. I use them for critical user journeys, smoke testing after deployments, and validating complex workflows. While valuable, I keep E2E tests minimal because they're expensive to maintain and can slow down the development cycle. The test pyramid suggests many unit tests, fewer integration tests, and very few E2E tests."

### **Testing Tools & Framework Questions**

**Q19: What is Testcontainers and how does it help testing?**
**A**: "Testcontainers provides lightweight, throwaway instances of databases, message brokers, web browsers, or other services in Docker containers for testing. It solves the problem of testing against real dependencies without the complexity of manual setup. I can spin up a PostgreSQL database, Redis cache, or Kafka broker for integration tests, ensuring tests run against realistic environments. Testcontainers makes tests more reliable and closer to production while keeping them fast and isolated."

**Q20: How do you test asynchronous code?**
**A**: "Testing async code requires careful handling of timing. I use several approaches: 1) CountDownLatch to wait for async operations to complete, 2) CompletableFuture with timeout for future-based code, 3) Mockito's verify with timeout for verifying async method calls, 4) Awaitility library for fluent waiting conditions, 5) Thread.sleep() as a last resort. The key is making tests deterministic rather than relying on arbitrary timing. I also test edge cases like timeouts, failures, and concurrent access."

**Q21: What are test doubles and when do you use each type?**
**A**: "Test doubles are objects that stand in for real dependencies during testing. Types include: Dummy objects that are passed but never used, Fake objects that have working implementations but aren't suitable for production, Stubs that provide canned responses, Spies that wrap real objects and record interactions, and Mocks that verify behavior. I choose based on what I need - Dummies for filling parameters, Fakes for lightweight alternatives, Stubs for simple responses, Spies for partial behavior verification, and Mocks for interaction testing."

**Q22: How do you implement continuous testing in your workflow?**
**A**: "Continuous testing involves running tests automatically as part of the development process. I implement this through: 1) IDE plugins that run tests on file changes, 2) Pre-commit hooks that run fast tests before commits, 3) CI/CD pipelines that run full test suites on every push, 4) Git hooks to prevent merging if tests fail, 5) Gradle/Maven daemon for faster test execution. The goal is immediate feedback to catch issues early. I prioritize fast unit tests for continuous feedback and run slower tests in CI pipelines."
