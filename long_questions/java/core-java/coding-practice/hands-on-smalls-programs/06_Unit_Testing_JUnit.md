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
