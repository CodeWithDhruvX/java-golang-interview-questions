# Java Advanced — Testing Deep Dive

> **Topics:** JUnit 5 (parameterized tests, extensions, `@Nested`), Mockito (`@Mock`, `@InjectMocks`, `verify`, `ArgumentCaptor`, spy), Spring Boot Testing (`@WebMvcTest`, `@DataJpaTest`, `@SpringBootTest`, `MockMvc`), Testcontainers, test design principles (test doubles, mocking strategy)

---

## 📋 Reading Progress

- [ ] **Section 1:** JUnit 5 — Annotations & Assertions (Q1–Q15)
- [ ] **Section 2:** Mockito — Mocking, Stubbing, Verification (Q16–Q30)
- [ ] **Section 3:** Spring Boot Testing — Slices & Integration (Q31–Q42)
- [ ] **Section 4:** Testcontainers — Real Infrastructure in Tests (Q43–Q48)
- [ ] **Section 5:** Test Design — Patterns & Advanced Topics (Q49–Q55)

> 🔖 **Last read:** <!-- -->

---

## Section 1: JUnit 5 — Annotations & Assertions (Q1–Q15)

### 1. JUnit 5 vs JUnit 4 — Key Differences
**Q: What changed?**
```java
// JUnit 4
import org.junit.*;
public class OldTest {
    @Before public void setUp() {}      // before each test
    @After  public void tearDown() {}   // after each test
    @Test(expected = Exception.class)
    public void testException() {}

    @Test(timeout = 1000)
    public void testPerformance() {}
}

// JUnit 5
import org.junit.jupiter.api.*;
class NewTest {
    @BeforeEach void setUp() {}         // before each test
    @AfterEach  void tearDown() {}      // after each test
    @Test
    void testException() {
        assertThrows(Exception.class, () -> risky());
    }
    @Test @Timeout(1)                   // in seconds
    void testPerformance() {}
}
```
**A:** JUnit 5 = JUnit Platform + JUnit Jupiter + JUnit Vintage. Key changes: `@Before/@After` → `@BeforeEach/AfterEach`, no `public` required, `assertThrows()` replaces `expected=`, `@Timeout` replaces `timeout=`. All test methods can be package-private.

---

### 2. @Test Basic Assertions
**Q: What is the difference between assertEquals and assertSame?**
```java
import org.junit.jupiter.api.*;
import static org.junit.jupiter.api.Assertions.*;

class AssertionsTest {
    @Test
    void examples() {
        String a = new String("hello");
        String b = new String("hello");

        assertEquals(a, b);    // ✅ same content (equals())
        assertNotSame(a, b);   // ✅ different references
        assertSame(a, a);      // ✅ same reference

        assertTrue("hello".startsWith("he"));
        assertFalse("hello".isEmpty());
        assertNull(null);
        assertNotNull("not null");

        assertEquals(0.1 + 0.2, 0.3, 1e-10); // floating point: use delta!
    }
}
```
**A:** `assertEquals` uses `.equals()`. `assertSame` checks reference equality (`==`). Always use delta for floating-point comparisons — `0.1 + 0.2 != 0.3` in IEEE 754.

---

### 3. assertAll — Grouped Assertions
**Q: What is the output when multiple assertions fail?**
```java
@Test
void withAssertAll() {
    User user = new User(null, "", -1);

    // Without assertAll: first failure stops all others
    // assertNotNull(user.getId());     // fails → test stops
    // assertFalse(user.getName().isEmpty()); // never reached!

    // With assertAll: all assertions run, all failures reported together
    assertAll("user validation",
        () -> assertNotNull(user.getId()),
        () -> assertFalse(user.getName().isEmpty()),
        () -> assertTrue(user.getAge() >= 0)
    );
    // Reports ALL 3 failures in one run
}
```
**A:** `assertAll` is a "soft assertions" equivalent — collects all failures and reports them together. Without it, you'd need to run the test n times to see n failures. Use for validating multiple properties of one object.

---

### 4. assertThrows — Exception Verification
**Q: What is the output?**
```java
@Test
void exceptionTests() {
    // Verify exception type
    Exception ex = assertThrows(IllegalArgumentException.class,
        () -> new User(null, null, -1));
    assertEquals("age must be non-negative", ex.getMessage());

    // Verify no exception
    assertDoesNotThrow(() -> new User(1L, "Alice", 25));

    // Verify exact subtype
    assertThrows(NumberFormatException.class, () -> Integer.parseInt("abc"));
    // assertTrue(assertThrows(RuntimeException.class, () -> Integer.parseInt("abc"))
    //     instanceof NumberFormatException); // also valid
}
```
**A:** `assertThrows` verifies that a specific exception type is thrown and returns the exception for further inspection. `assertDoesNotThrow` verifies the opposite. Prefer `assertThrows` over `@Test(expected=)` — it gives you the exception object.

---

### 5. @ParameterizedTest — Data-Driven Tests
**Q: How many tests run?**
```java
import org.junit.jupiter.params.*;
import org.junit.jupiter.params.provider.*;

class ParameterizedTests {
    @ParameterizedTest
    @ValueSource(strings = {"  ", "", "\t", "\n"})
    void blankStrings(String input) {
        assertTrue(input.isBlank()); // runs 4 times
    }

    @ParameterizedTest
    @CsvSource({
        "2,  3, 5",
        "10, 20, 30",
        "-5, 5,  0"
    })
    void add(int a, int b, int expected) {
        assertEquals(expected, a + b); // runs 3 times
    }

    @ParameterizedTest
    @EnumSource(value = DayOfWeek.class, names = {"SATURDAY", "SUNDAY"})
    void weekends(DayOfWeek day) {
        assertTrue(isWeekend(day)); // runs 2 times
    }
}
```
**A:** 4 + 3 + 2 = **9 tests** total. JUnit generates distinct test names for each parameter set. Parameterized tests eliminate test duplication — one test method covers many input combinations.

---

### 6. @MethodSource — Complex Test Data
**Q: What does this generate?**
```java
import java.util.stream.*;

class MethodSourceTest {
    @ParameterizedTest
    @MethodSource("userProvider")
    void validateUser(String name, int age, boolean valid) {
        assertEquals(valid, userValidator.isValid(name, age));
    }

    static Stream<Arguments> userProvider() {
        return Stream.of(
            Arguments.of("Alice", 25, true),
            Arguments.of("",      25, false),  // blank name
            Arguments.of("Bob",   -1, false),  // negative age
            Arguments.of("Carol", 0,  false),  // zero age
            Arguments.of("Dave",  18, true)    // boundary
        );
    }
}
```
**A:** Generates **5 parameterized test cases** from the Stream. `@MethodSource` is preferred over `@CsvSource` for complex objects or when test data requires construction logic.

---

### 7. @Nested — Hierarchical Test Organization
**Q: How does nesting improve test organization?**
```java
class UserServiceTest {
    @Nested
    class WhenUserExists {
        User user;
        @BeforeEach void setUp() { user = new User(1L, "Alice"); }

        @Test void getUser_returnsUser() { /* ... */ }
        @Test void updateUser_updatesFields() { /* ... */ }
        @Test void deleteUser_removesUser() { /* ... */ }
    }

    @Nested
    class WhenUserNotFound {
        @Test void getUser_throwsNotFoundException() { /* ... */ }
        @Test void updateUser_throwsNotFoundException() { /* ... */ }
    }

    @Nested @DisplayName("Email validation")
    class EmailValidation {
        @Test @DisplayName("accepts valid email")   void validEmail() {}
        @Test @DisplayName("rejects blank email")   void blankEmail() {}
        @Test @DisplayName("rejects invalid format") void invalidFormat() {}
    }
}
```
**A:** `@Nested` groups related tests hierarchically and allows `@BeforeEach` at each level. Test runner shows a tree view. Each nested class has its own lifecycle — `@BeforeEach` in `WhenUserExists` only runs for tests in that class.

---

### 8. @ExtendWith — JUnit 5 Extension Model
**Q: What does @ExtendWith do?**
```java
// Register Mockito extension
@ExtendWith(MockitoExtension.class)
class MockitoTest {
    @Mock UserRepository repo;
    @InjectMocks UserService service;
    // No @RunWith(MockitoJUnitRunner.class) needed (that's JUnit 4 style)
}

// Register Spring extension
@ExtendWith(SpringExtension.class) // included in @SpringBootTest etc.
class SpringTest { }

// Custom extension:
class TimingExtension implements BeforeEachCallback, AfterEachCallback {
    public void beforeEach(ExtensionContext ctx) { /* record start time */ }
    public void afterEach(ExtensionContext ctx)  { /* print elapsed */ }
}
@ExtendWith(TimingExtension.class)
class TimedTest { }
```
**A:** `@ExtendWith` replaces JUnit 4's `@RunWith`. Extensions hook into the test lifecycle via callbacks. Multiple extensions can be combined. `@MockitoExtension` initializes mocks and verifies strictness.

---

### 9. @BeforeAll and @AfterAll — Class-Level Setup
**Q: What constraint applies to @BeforeAll methods?**
```java
class DatabaseTest {
    static Connection conn;

    @BeforeAll
    static void connect() throws SQLException { // MUST be static (or @TestInstance(PER_CLASS))
        conn = DriverManager.getConnection("jdbc:h2:mem:test");
        System.out.println("Connected once for all tests");
    }

    @Test void test1() { /* use conn */ }
    @Test void test2() { /* use conn */ }

    @AfterAll
    static void disconnect() throws SQLException {
        if (conn != null) conn.close();
        System.out.println("Disconnected after all tests");
    }
}
```
**A:** `@BeforeAll/@AfterAll` methods must be `static` (one instance per class by default). Use for expensive setup (DB connection, embedded server). Alternative: `@TestInstance(PER_CLASS)` — allows non-static, one instance for all tests.

---

### 10. @TestInstance(PER_CLASS)
**Q: What changes with PER_CLASS?**
```java
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class StatefulTest {
    private List<String> results = new ArrayList<>(); // shared mutable state across tests

    @BeforeAll
    void setup() { results.add("setup"); } // non-static now OK

    @Test @Order(1) void first()  { results.add("first"); }
    @Test @Order(2) void second() { results.add("second"); }

    @AfterAll
    void verify() {
        assertEquals(List.of("setup", "first", "second"), results);
    }
}
```
**A:** `PER_CLASS` creates one test instance for all test methods. State is shared. `@BeforeAll/@AfterAll` can be non-static. Default `PER_METHOD` creates a new instance per test (better isolation). Use `PER_CLASS` only when shared state is intentional.

---

### 11. @TempDir — Temporary File System
**Q: What does @TempDir provide?**
```java
import java.nio.file.*;

class FileProcessorTest {
    @TempDir Path tempDir; // injected by JUnit — auto-deleted after test

    @Test
    void processFile() throws IOException {
        Path input  = tempDir.resolve("input.txt");
        Path output = tempDir.resolve("output.txt");

        Files.writeString(input, "hello world");
        fileProcessor.process(input, output);

        assertThat(Files.readString(output)).isEqualTo("HELLO WORLD");
    }
    // tempDir and all its contents are deleted automatically
}
```
**A:** `@TempDir` injects a fresh temporary directory for each test (or class with `PER_CLASS`). Automatically cleaned up after the test. Replaces manual `File.createTempDir()` + `deleteOnExit()`.

---

### 12. Assumptions — Conditional Test Execution
**Q: What happens when an assumption fails?**
```java
import static org.junit.jupiter.api.Assumptions.*;

class ConditionalTest {
    @Test
    void onlyOnLinux() {
        assumeTrue("Linux".equals(System.getProperty("os.name")));
        // test continues only on Linux; skipped otherwise (not failed!)
    }

    @Test
    void onlyWithDatabase() {
        assumingThat(dbAvailable(),
            () -> {
                // this block runs only if DB is available
                assertThat(repo.count()).isGreaterThan(0);
            }
        );
        // test continues here regardless
    }
}
```
**A:** A failed assumption causes the test to be **skipped** (not failed). Use assumptions for environment-dependent tests (OS-specific, CI/CD-specific). `assumingThat` runs a block conditionally but doesn't stop the test if the condition is false.

---

### 13. @Disabled and @EnabledOnOs
**Q: When do these activate?**
```java
import org.junit.jupiter.api.condition.*;

class ConditionalTests {
    @Test @Disabled("known bug JIRA-123")
    void skippedTest() { } // never runs

    @Test @EnabledOnOs(OS.LINUX)
    void linuxOnly() { }

    @Test @EnabledOnOs({OS.WINDOWS, OS.MAC})
    void notLinux() { }

    @Test @EnabledIfSystemProperty(named = "test.env", matches = "integration")
    void integrationOnly() { } // run with -Dtest.env=integration

    @Test @EnabledIfEnvironmentVariable(named = "CI", matches = "true")
    void ciOnly() { }
}
```
**A:** Conditional annotations on JUnit 5 tests control execution environment. `@Disabled` always skips with a note. `@EnabledOnOs` skips on other OSes. More composable than `assumeTrue`. All disabled tests show as skipped in the test report.

---

### 14. @RepeatedTest — Stability Tests
**Q: What is the output?**
```java
class StabilityTest {
    @RepeatedTest(value = 5, name = "{displayName} #{currentRepetition}/{totalRepetitions}")
    void isConsistent(RepetitionInfo info) {
        int result = randomSeed.compute(); // check for non-determinism
        assertEquals(42, result, "Failed on repetition " + info.getCurrentRepetition());
    }
}
// Output: "isConsistent #1/5", "isConsistent #2/5", etc.
```
**A:** `@RepeatedTest` runs the test `n` times. Useful for testing non-deterministic code (Thread.sleep, random number generation, race conditions). `RepetitionInfo` provides current and total repetition counts.

---

### 15. AssertJ — Fluent Assertions
**Q: Why use AssertJ over JUnit assertions?**
```java
import static org.assertj.core.api.Assertions.*;

class AssertJTest {
    @Test
    void fluent() {
        List<String> names = List.of("Alice", "Bob", "Charlie");

        assertThat(names)
            .hasSize(3)
            .contains("Alice", "Bob")
            .doesNotContain("Dave")
            .allMatch(n -> n.length() > 2);

        assertThat("hello world")
            .startsWith("hello")
            .endsWith("world")
            .containsIgnoringCase("WORLD");

        assertThatThrownBy(() -> Integer.parseInt("abc"))
            .isInstanceOf(NumberFormatException.class)
            .hasMessageContaining("abc");
    }
}
```
**A:** AssertJ provides: (1) IDE-friendly fluent API with auto-complete, (2) better failure messages, (3) no need to import many assert methods, (4) supports complex object assertions. Preferred over raw JUnit assertions in modern Java.

---

## Section 2: Mockito — Mocking, Stubbing, Verification (Q16–Q30)

### 16. @Mock vs @Spy vs @InjectMocks
**Q: What is the difference?**
```java
@ExtendWith(MockitoExtension.class)
class ServiceTest {
    @Mock   UserRepository repo;   // full mock: all methods return defaults (null, 0, false)
    @Spy    EmailValidator  validator; // spy: real object, unless stubbed
    @InjectMocks UserService service; // inject both @Mock and @Spy into this

    @Test
    void test() {
        when(repo.findById(1L)).thenReturn(Optional.of(new User(1L, "Alice")));
        // validator.isValid() will call the REAL implementation
        // unless: doReturn(true).when(validator).isValid(any());

        User result = service.getUser(1L);
        assertEquals("Alice", result.getName());
    }
}
```
**A:** `@Mock`: completely fake, all calls return empty/default. `@Spy`: wraps a real object — unstubbed calls go to the real implementation. `@InjectMocks`: creates the class under test and injects `@Mock`/`@Spy` fields.

---

### 17. when().thenReturn() — Stubbing
**Q: What is the output of each call?**
```java
@Test
void stubbing() {
    UserRepository mock = Mockito.mock(UserRepository.class);

    when(mock.findById(1L)).thenReturn(Optional.of(new User(1L, "Alice")));
    when(mock.findById(2L)).thenReturn(Optional.empty());
    when(mock.findById(anyLong())).thenReturn(Optional.of(new User(99L, "default")));

    System.out.println(mock.findById(1L).get().getName()); // Alice (specific stub wins)
    System.out.println(mock.findById(2L).isPresent());     // false
    System.out.println(mock.findById(3L).get().getName()); // default (any() matches)
    System.out.println(mock.findById(0L).get().getName()); // Alice? No — default (specific stubs override any())
}
```
**A:** Specific stubs (`findById(1L)`) override general stubs (`anyLong()`). More specific stubs registered later take priority. `any()` matches any argument including null.

---

### 18. thenThrow — Stubbing Exceptions
**Q: What is the output?**
```java
@Test
void stubbingExceptions() {
    UserRepository mock = Mockito.mock(UserRepository.class);

    // Stub to throw
    when(mock.findById(-1L)).thenThrow(new IllegalArgumentException("invalid id"));

    // Stub void method to throw:
    doThrow(new RuntimeException("db error")).when(mock).delete(any());

    assertThrows(IllegalArgumentException.class, () -> mock.findById(-1L));
    assertThrows(RuntimeException.class, () -> mock.delete(new User()));

    // Consecutive stubs:
    when(mock.findById(99L))
        .thenReturn(Optional.of(new User(99L, "temp")))
        .thenThrow(new RuntimeException("gone"))
        .thenReturn(Optional.empty());
}
```
**A:** First call to `mock.findById(99L)` → returns User. Second call → throws RuntimeException. Third+ → returns empty. Consecutive stubs model state changes in sequences.

---

### 19. Argument Matchers
**Q: What is the rule when mixing matchers with exact values?**
```java
@Test
void argumentMatchers() {
    UserRepository mock = Mockito.mock(UserRepository.class);

    // All arguments must be matchers if any is a matcher:
    // WRONG: when(mock.save(any(), 1)).thenReturn(user); // compile error if second param is primitive
    // RIGHT:
    when(mock.findByNameAndAge(anyString(), anyInt())).thenReturn(List.of());
    when(mock.findByNameAndAge(eq("Alice"), eq(25))).thenReturn(List.of(alice));

    // Custom matcher:
    when(mock.findByName(argThat(name -> name.startsWith("A"))))
        .thenReturn(List.of(alice));

    verify(mock).findByName(argThat(n -> n.length() > 3));
}
```
**A:** If any argument uses a matcher, **all** arguments must use matchers. Mix `eq()` for exact values with matchers. `argThat()` accepts a lambda predicate.

---

### 20. verify — Interaction Verification
**Q: What do these verify statements check?**
```java
@Test
void verificationExamples() {
    UserRepository mock = Mockito.mock(UserRepository.class);
    service.process(1L);

    verify(mock).findById(1L);                              // called exactly once
    verify(mock, times(2)).save(any());                     // called exactly 2 times
    verify(mock, never()).delete(any());                     // never called
    verify(mock, atLeastOnce()).findById(anyLong());         // >= 1 times
    verify(mock, atMost(3)).findByName(anyString());        // <= 3 times
    verify(mock, timeout(1000)).asyncOperation();           // called within 1s

    verifyNoMoreInteractions(mock); // no other methods called
    verifyNoInteractions(otherMock); // zero interactions
}
```
**A:** `verify` checks that a method was called with specific arguments. `times(n)` = exactly n times. Always verify interactions explicitly — don't rely on no exception being thrown as proof of correct behavior.

---

### 21. ArgumentCaptor — Capturing Arguments
**Q: What does ArgumentCaptor capture?**
```java
@ExtendWith(MockitoExtension.class)
class EmailServiceTest {
    @Mock SmtpClient smtpClient;
    @InjectMocks EmailService emailService;

    @Captor ArgumentCaptor<EmailMessage> emailCaptor;

    @Test
    void sendWelcomeEmail() {
        emailService.sendWelcome(new User(1L, "alice@example.com"));

        verify(smtpClient).send(emailCaptor.capture());
        EmailMessage sent = emailCaptor.getValue();

        assertEquals("alice@example.com", sent.getTo());
        assertTrue(sent.getSubject().contains("Welcome"));
        assertFalse(sent.getBody().isBlank());
    }
}
```
**A:** `ArgumentCaptor` captures the exact argument passed to a mocked method — lets you inspect it without adding `equals()` methods. `getValue()` = last captured value. `getAllValues()` for multiple captures.

---

### 22. BDDMockito — Given/When/Then Style
**Q: How does BDDMockito differ from Mockito?**
```java
import static org.mockito.BDDMockito.*;

@Test
void bddStyle() {
    // GIVEN
    given(userRepo.findById(1L)).willReturn(Optional.of(new User(1L, "Alice")));
    given(orderRepo.findByUserId(1L)).willReturn(List.of(order1, order2));

    // WHEN
    UserProfile profile = userService.getProfile(1L);

    // THEN
    then(userRepo).should().findById(1L);
    then(orderRepo).should(times(1)).findByUserId(1L);
    assertThat(profile.getOrders()).hasSize(2);
}
```
**A:** `BDDMockito` is a Mockito wrapper that maps to BDD vocabulary: `given()` = `when()`, `then()` = `verify()`. Identical behavior, just reads more naturally with Given/When/Then test structure.

---

### 23. Spy — Partial Mocking
**Q: What is the output?**
```java
@Test
void spyExample() {
    List<String> realList = new ArrayList<>();
    List<String> spy = Mockito.spy(realList);

    spy.add("first");      // calls REAL add()
    spy.add("second");     // calls REAL add()

    System.out.println(spy.size()); // 2 — real size

    when(spy.size()).thenReturn(100); // stub size()
    System.out.println(spy.size()); // 100 — stubbed

    // GOTCHA: use doReturn() for spy to avoid calling real method during stubbing:
    // when(spy.get(0)).thenReturn("mock"); // calls real get(0) first!
    doReturn("mock").when(spy).get(0);
}
```
**A:**
```
2
100
```
Spy wraps a real object. Real methods run unless stubbed. Use `doReturn().when()` (not `when().thenReturn()`) with spies to avoid calling the real method during stub setup.

---

### 24. Mockito.reset() — Anti-Pattern
**Q: Why is reset() usually a code smell?**
```java
@Test
void antiPattern() {
    when(mock.findById(1L)).thenReturn(Optional.of(user1));
    service.first();

    Mockito.reset(mock); // clears all stubs and interactions — why?

    when(mock.findById(2L)).thenReturn(Optional.of(user2));
    service.second();
}
// BETTER: separate test methods, each with fresh @Mock setup via @BeforeEach
```
**A:** `reset()` is a code smell — usually indicates the test does too much. Split into separate test methods with clean setup per test. Mockito's `@ExtendWith(MockitoExtension.class)` resets all mocks between tests automatically.

---

### 25. Strict Stubbing — UnnecessaryStubbingException
**Q: What does Mockito's strict mode catch?**
```java
@ExtendWith(MockitoExtension.class) // strict mocking enabled by default
class StrictTest {
    @Mock UserRepository repo;
    @InjectMocks UserService service;

    @Test
    void test() {
        when(repo.findById(1L)).thenReturn(Optional.of(user)); // stubbed
        when(repo.findById(2L)).thenReturn(Optional.empty());  // stubbed but never used!

        service.getUser(1L);
        // MockitoException: UnnecessaryStubbingException — stub for findById(2L) is never used
    }
}
```
**A:** Strict stubbing (default in JUnit 5 with `MockitoExtension`) fails the test if a stub is set up but never used — catches stale/wrong stubs. Disable with `@MockitoSettings(strictness = Strictness.LENIENT)` when needed.

---

### 26. Mockito Answer — Custom Return Logic
**Q: When do you use Answer?**
```java
@Test
void customAnswer() {
    when(userRepo.save(any(User.class))).thenAnswer(invocation -> {
        User arg = invocation.getArgument(0);
        arg.setId(ThreadLocalRandom.current().nextLong(1, 1000)); // assign random ID
        return arg; // return same object with ID set
    });

    User saved = userService.register(new User(null, "alice@test.com"));
    assertNotNull(saved.getId());
    assertEquals("alice@test.com", saved.getEmail());
}
```
**A:** `thenAnswer` provides dynamic stub behavior — the return value depends on the argument. Alternatives: `RETURNS_DEEP_STUBS` (for chained calls), `RETURNS_SELF` (builder pattern), `CALLS_REAL_METHODS`.

---

### 27. InOrder — Verifying Call Sequence
**Q: What does InOrder verify?**
```java
@Test
void verifyOrder() {
    UserRepository repo = mock(UserRepository.class);
    AuditLog auditLog = mock(AuditLog.class);

    service.createUser(new User("Alice"));

    InOrder inOrder = inOrder(repo, auditLog);
    inOrder.verify(repo).save(any(User.class));     // must happen first
    inOrder.verify(auditLog).log(anyString());      // must happen second
}
```
**A:** `InOrder` verifies that methods were called in a specific sequence across mocks. Without `InOrder`, `verify()` checks that calls happened at some point — regardless of order.

---

### 28. Mockito with Static Methods (Mockito 3.4+)
**Q: How do you mock static methods?**
```java
import org.mockito.*;

@Test
void staticMock() {
    try (MockedStatic<UUID> mockedUUID = Mockito.mockStatic(UUID.class)) {
        UUID fixedId = UUID.fromString("00000000-0000-0000-0000-000000000001");
        mockedUUID.when(UUID::randomUUID).thenReturn(fixedId);

        String id = orderService.generateOrderId(); // internally calls UUID.randomUUID()
        assertEquals("00000000-0000-0000-0000-000000000001", id);
    }
    // After try: UUID.randomUUID() behaves normally again
}
```
**A:** `MockedStatic` mocks static methods within a try-with-resources scope. The mock is automatically closed and restored. Requires `mockito-inline` (included in `mockito-core` 5.x).

---

### 29. Mockito Constructor Mocking
**Q: How do you control object construction?**
```java
@Test
void constructorMock() {
    File mockFile = mock(File.class);
    when(mockFile.exists()).thenReturn(false);

    try (MockedConstruction<File> mc = Mockito.mockConstruction(File.class,
        (mock, ctx) -> when(mock.exists()).thenReturn(true))) {

        // Any new File("...") inside this try block returns a mock
        fileService.checkAndCreate("/tmp/test.txt");
        verify(mc.constructed().get(0)).exists();
    }
}
```
**A:** `MockedConstruction` intercepts `new File(...)` calls and returns a mock. Useful when testing code that constructs objects internally and you can't inject them. Use sparingly — usually a sign the code needs refactoring.

---

### 30. Mockito Verification Timeout — Async Tests
**Q: How do you verify async interactions?**
```java
@Test
void asyncVerification() throws InterruptedException {
    // Service that processes in a background thread
    asyncService.process(new OrderEvent(1L));

    // Verify the repo was called within 2 seconds
    verify(orderRepo, timeout(2000)).save(any(Order.class));
    verify(emailSender, timeout(2000).times(1)).sendConfirmation(anyLong());
}
```
**A:** `timeout(ms)` retries verification until the call is observed or the timeout expires. Essential for testing `@Async` methods or background processing. `timeout(ms).times(n)` combines timing with invocation count.

---

## Section 3: Spring Boot Testing — Slices & Integration (Q31–Q42)

### 31. @SpringBootTest — Full Context
**Q: What are the different webEnvironment options?**
```java
// Full app, no server — use MockMvc for HTTP tests
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.MOCK)
class MockMvcTest { @Autowired MockMvc mockMvc; }

// Full app, random port — use TestRestTemplate for real HTTP
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class RealServerTest {
    @Autowired TestRestTemplate restTemplate;
    @LocalServerPort int port;
}

// Full app, fixed port 8080
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.DEFINED_PORT)
class FixedPortTest { }

// No web context at all (service/repo tests)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.NONE)
class NoWebTest { }
```
**A:** `MOCK` (default) for fast controller tests. `RANDOM_PORT` for full stack integration tests. `NONE` for service-layer tests that don't need the web tier. `@LocalServerPort` injects the actual port number.

---

### 32. @WebMvcTest — Controller Slice
**Q: What beans are loaded?**
```java
@WebMvcTest(UserController.class) // only MVC components
class UserControllerTest {
    @Autowired MockMvc mockMvc;
    @Autowired ObjectMapper objectMapper;

    @MockBean UserService userService; // service is mocked
    // UserRepository is NOT in context — @WebMvcTest doesn't load it

    @Test
    void getUser_ok() throws Exception {
        when(userService.findById(1L)).thenReturn(new User(1L, "Alice"));

        mockMvc.perform(get("/api/users/1"))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.name").value("Alice"))
            .andDo(print()); // print request/response to console
    }
}
```
**A:** `@WebMvcTest` loads: controllers, filters, `@ControllerAdvice`, security config, `Jackson`, `Validator`. Does NOT load: `@Service`, `@Repository`, `@Component`. Much faster than `@SpringBootTest`.

---

### 33. MockMvc — POST with JSON Body
**Q: Does this test pass?**
```java
@Test
void createUser() throws Exception {
    User req = new User(null, "bob@test.com");
    User saved = new User(2L, "bob@test.com");

    when(userService.create(any())).thenReturn(saved);

    mockMvc.perform(post("/api/users")
            .contentType(MediaType.APPLICATION_JSON)
            .content(objectMapper.writeValueAsString(req)))
        .andExpect(status().isCreated())
        .andExpect(jsonPath("$.id").value(2))
        .andExpect(jsonPath("$.email").value("bob@test.com"))
        .andExpect(header().string("Location", endsWith("/users/2")));
}
```
**A:** Yes, if the controller returns `201 Created` with `Location` header. `objectMapper.writeValueAsString()` serializes the request body. `jsonPath("$.id")` uses JSONPath syntax to navigate the response JSON.

---

### 34. MockMvc — Testing Security
**Q: What does @WithMockUser do?**
```java
@Test
@WithMockUser(username = "alice", roles = {"USER"})
void authenticatedUser() throws Exception {
    mockMvc.perform(get("/api/profile"))
        .andExpect(status().isOk()); // 200 with USER role
}

@Test
@WithMockUser(username = "admin", roles = {"ADMIN"})
void adminEndpoint() throws Exception {
    mockMvc.perform(delete("/api/admin/users/1"))
        .andExpect(status().isNoContent());
}

@Test
void unauthenticated() throws Exception {
    mockMvc.perform(get("/api/profile"))
        .andExpect(status().isUnauthorized()); // 401
}
```
**A:** `@WithMockUser` simulates an authenticated user in the Spring Security context. Avoids needing a real JWT or login flow in controller tests. `roles = "ADMIN"` automatically adds `ROLE_ADMIN` prefix.

---

### 35. @DataJpaTest — Repository Slice
**Q: Does this use the real database?**
```java
@DataJpaTest
// @AutoConfigureTestDatabase(replace = Replace.NONE) // use real DB instead of H2
class UserRepositoryTest {
    @Autowired TestEntityManager em;  // simplified EntityManager for tests
    @Autowired UserRepository repo;

    @Test
    @Transactional
    void findByEmail() {
        User user = new User(null, "test@test.com");
        em.persistAndFlush(user); // save and flush to DB

        Optional<User> found = repo.findByEmail("test@test.com");
        assertThat(found).isPresent();
        assertThat(found.get().getEmail()).isEqualTo("test@test.com");
    }
    // Transaction is rolled back after each test
}
```
**A:** By default, `@DataJpaTest` uses an embedded H2 database (not the real DB). `TestEntityManager` wraps `EntityManager` with test-friendly methods. Add `@AutoConfigureTestDatabase(replace = Replace.NONE)` to use the real database.

---

### 36. @DataJpaTest — Testing Named Queries and @Query
**Q: Does Spring Data validate @Query at test time?**
```java
@DataJpaTest
class OrderRepositoryTest {
    @Autowired TestEntityManager em;
    @Autowired OrderRepository repo;

    @Test
    void customQuery_returnsMatchingOrders() {
        Order o1 = em.persist(new Order("pending", LocalDate.now()));
        Order o2 = em.persist(new Order("completed", LocalDate.now()));
        em.flush();

        List<Order> pending = repo.findByStatus("pending");
        assertThat(pending).hasSize(1);
        assertThat(pending.get(0).getStatus()).isEqualTo("pending");
    }
    // @Query JPQL is validated at context startup — syntax errors caught here
}
```
**A:** Yes — `@DataJpaTest` starts a JPA context that validates all `@NamedQuery` and `@Query` annotations. JPQL syntax errors are caught at test time, not production runtime.

---

### 37. @RestClientTest — Testing HTTP Clients
**Q: What does @RestClientTest configure?**
```java
@RestClientTest(ExternalApiClient.class) // test the HTTP client bean
class ExternalApiClientTest {
    @Autowired ExternalApiClient client;
    @Autowired MockRestServiceServer server; // intercepts RestTemplate calls

    @Test
    void getProduct() {
        server.expect(requestTo("/products/1"))
            .andExpect(method(HttpMethod.GET))
            .andRespond(withSuccess("{\"id\":1,\"name\":\"Phone\"}", MediaType.APPLICATION_JSON));

        Product p = client.getProduct(1L);
        assertEquals("Phone", p.getName());
        server.verify(); // assert all expected requests were made
    }
}
```
**A:** `@RestClientTest` loads only `RestTemplate` + Jackson. `MockRestServiceServer` intercepts HTTP calls and returns canned responses — no real network. Use `MockWebServer` (from OkHttp) for `WebClient` tests.

---

### 38. @SpringBootTest — Application Properties Override
**Q: How do you override properties in tests?**
```java
@SpringBootTest(properties = {
    "spring.datasource.url=jdbc:h2:mem:testdb",
    "feature.experimental=true",
    "kafka.bootstrap-servers=localhost:19092"
})
class ConfiguredTest { }

// Or use a test-specific file:
@SpringBootTest
@TestPropertySource(locations = "classpath:application-test.yml")
class PropertyFileTest { }

// Or use @ActiveProfiles:
@SpringBootTest
@ActiveProfiles("test")
class ProfileTest { } // loads application-test.yml automatically
```
**A:** Property override methods ranked by priority: `@SpringBootTest(properties=...)` > `@TestPropertySource` > Profile-based overrides. Use `@ActiveProfiles("test")` as the cleanest solution — keep all test config in `application-test.yml`.

---

### 39. @MockBean — Replace Beans in Context
**Q: What is the difference between @MockBean and @Mock?**
```java
@SpringBootTest
class ServiceIntegrationTest {
    @MockBean PaymentGateway gateway; // replaces real bean in Spring context
    @Autowired OrderService service;  // gets real service with mocked gateway

    @Test
    void createOrder_chargesGateway() {
        when(gateway.charge(any())).thenReturn(PaymentResult.success("txn-123"));

        Order order = service.createOrder(request);

        verify(gateway).charge(argThat(req -> req.getAmount().equals(order.getTotal())));
        assertEquals("txn-123", order.getTransactionId());
    }
}
```
**A:** `@MockBean` registers a Mockito mock into the Spring Application Context — replaces the real `PaymentGateway` bean. All beans that depend on `PaymentGateway` receive the mock. `@Mock` is only for non-Spring unit tests.

---

### 40. @SpyBean — Partial Mocking in Context
**Q: When do you use @SpyBean?**
```java
@SpringBootTest
class AuditTest {
    @SpyBean AuditService auditService; // real, but we can verify calls
    @Autowired OrderService orderService; // uses real auditService

    @Test
    void orderCreation_isAudited() {
        orderService.create(request);

        // Verify the real auditService was called correctly
        verify(auditService).log(eq("ORDER_CREATED"), any(AuditEvent.class));
    }
}
```
**A:** `@SpyBean` wraps the real Spring bean with a Mockito spy. Real methods run but interactions are recorded. Use when you want to verify a real bean was called — without fully mocking it.

---

### 41. @Sql — Database State Setup
**Q: What does @Sql do?**
```java
@SpringBootTest
@Transactional
class IntegrationTest {
    @Autowired UserRepository repo;

    @Sql("/sql/insert-test-users.sql") // runs before this test method
    @Test
    void findActiveUsers() {
        List<User> users = repo.findByActiveTrue();
        assertThat(users).hasSize(3); // 3 from SQL file
    }

    @Sql(scripts = "/sql/cleanup.sql", executionPhase = Sql.ExecutionPhase.AFTER_TEST_METHOD)
    @Test
    void cleanupAfterTest() { }
}
```
**A:** `@Sql` executes SQL scripts before/after test methods. `BEFORE_TEST_METHOD` (default) sets up test data. `AFTER_TEST_METHOD` cleans up. Can be placed at class level to apply to all tests. Combined with `@Transactional`, provides clean state per test.

---

### 42. Test Slices — Reference Table
**Q: Which slice for which layer?**
```
@WebMvcTest         → Controllers, Filters, Security, Jackson
@DataJpaTest        → Repositories, JPA, Validators, H2
@DataMongoTest      → Mongo repositories, embedded MongoDB
@RestClientTest     → RestTemplate clients, Jackson
@DataRedisTest      → Redis repositories
@WebFluxTest        → WebFlux controllers (reactive)
@JsonTest           → ObjectMapper JSON serialization
@DataJdbcTest       → Spring Data JDBC repositories
@DataR2dbcTest      → R2DBC reactive repositories
@SpringBootTest     → Full context (use as last resort — slowest)

Rule: Use the narrowest slice possible.
  Fast unit test (ms): @ExtendWith(MockitoExtension)
  Medium slice test (s): @WebMvcTest, @DataJpaTest
  Slow integration test (10s+): @SpringBootTest
```
**A:** Slices dramatically reduce test startup time by loading only the relevant layer. Design your architecture so each layer is independently testable. Use `@SpringBootTest` only for true end-to-end smoke tests.

---

## Section 4: Testcontainers — Real Infrastructure in Tests (Q43–Q48)

### 43. Testcontainers — PostgreSQL
**Q: What does Testcontainers start?**
```java
import org.testcontainers.junit.jupiter.*;
import org.testcontainers.containers.*;

@Testcontainers
@DataJpaTest
@AutoConfigureTestDatabase(replace = Replace.NONE)
class RealDatabaseTest {
    @Container
    static PostgreSQLContainer<?> postgres = new PostgreSQLContainer<>("postgres:15-alpine")
        .withDatabaseName("testdb")
        .withUsername("test")
        .withPassword("test");

    @DynamicPropertySource
    static void properties(DynamicPropertyRegistry reg) {
        reg.add("spring.datasource.url",      postgres::getJdbcUrl);
        reg.add("spring.datasource.username", postgres::getUsername);
        reg.add("spring.datasource.password", postgres::getPassword);
    }

    @Autowired UserRepository repo;

    @Test
    void saveAndFind() {
        User u = repo.save(new User("test@test.com"));
        assertThat(repo.findById(u.getId())).isPresent();
    }
}
```
**A:** Testcontainers starts a real PostgreSQL Docker container for each test run. `@DynamicPropertySource` injects the container's URL/credentials into Spring's config. Container is shared across tests in the class when `static`.

---

### 44. Testcontainers — Kafka
**Q: How do you test against a real Kafka?**
```java
@Testcontainers
@SpringBootTest
class KafkaIntegrationTest {
    @Container
    static KafkaContainer kafka = new KafkaContainer(
        DockerImageName.parse("confluentinc/cp-kafka:7.5.0"));

    @DynamicPropertySource
    static void props(DynamicPropertyRegistry r) {
        r.add("spring.kafka.bootstrap-servers", kafka::getBootstrapServers);
    }

    @Autowired KafkaTemplate<String, OrderEvent> template;

    @Test
    void publishAndConsume() throws InterruptedException {
        CountDownLatch latch = new CountDownLatch(1);
        AtomicReference<OrderEvent> received = new AtomicReference<>();

        // Start consumer
        template.receive("orders", 0, 0L); // poll from beginning

        template.send("orders", new OrderEvent(1L, "NEW"));

        assertThat(latch.await(10, TimeUnit.SECONDS)).isTrue();
        assertEquals(1L, received.get().getOrderId());
    }
}
```
**A:** Testcontainers starts a real Kafka broker in Docker. No embedded Kafka needed. Tests catch issues that embedded Kafka hides (partitioning, rebalances, actual serialization).

---

### 45. Testcontainers — Reusable Containers
**Q: How do you speed up Testcontainers?**
```java
// Shared container across multiple test classes
abstract class BaseIntegrationTest {
    @ClassRule // JUnit 4 style, or:
    static PostgreSQLContainer<?> postgres;

    static {
        postgres = new PostgreSQLContainer<>("postgres:15-alpine")
            .withReuse(true); // reuse container across JVM restarts!
        postgres.start();
    }
}

// All test classes extending BaseIntegrationTest share one container
class UserRepoTest   extends BaseIntegrationTest { /* ... */ }
class OrderRepoTest  extends BaseIntegrationTest { /* ... */ }
class ProductRepoTest extends BaseIntegrationTest { /* ... */ }
```
**A:** `withReuse(true)` keeps the container alive between test runs (requires `~/.testcontainers.properties: testcontainers.reuse.enable=true`). Dramatically reduces test suite startup time — container starts once, not per class.

---

### 46. Testcontainers — Redis
**Q: What does this test?**
```java
@Testcontainers
@SpringBootTest
class CacheIntegrationTest {
    @Container
    static GenericContainer<?> redis = new GenericContainer<>("redis:7-alpine")
        .withExposedPorts(6379);

    @DynamicPropertySource
    static void props(DynamicPropertyRegistry r) {
        r.add("spring.data.redis.host", redis::getHost);
        r.add("spring.data.redis.port", () -> redis.getMappedPort(6379).toString());
    }

    @Autowired ProductService productService;

    @Test
    void caching_reducesDbCalls() {
        productService.getProduct(1L); // DB call
        productService.getProduct(1L); // cache hit
        productService.getProduct(1L); // cache hit

        verify(productRepo, times(1)).findById(1L); // DB called once
    }
}
```
**A:** Tests the real `@Cacheable` integration with Redis. Validates that the cache actually works — not possible with mocks. `GenericContainer` handles any Docker image without a dedicated Testcontainers module.

---

### 47. WireMock — Mocking External HTTP APIs
**Q: How do you test code that calls external REST APIs?**
```java
import com.github.tomakehurst.wiremock.junit5.*;

@ExtendWith(WireMockExtension.class)
@SpringBootTest
class ExternalApiTest {
    @InjectWireMock WireMockServer wireMock;

    @DynamicPropertySource
    static void props(DynamicPropertyRegistry r) {
        r.add("external.api.url", () -> "http://localhost:" + wireMock.getPort());
    }

    @Test
    void getShippingRate() {
        wireMock.stubFor(get(urlEqualTo("/rates/US"))
            .willReturn(aResponse()
                .withStatus(200)
                .withHeader("Content-Type", "application/json")
                .withBody("{\"rate\": 5.99}")));

        ShippingRate rate = shippingClient.getRate("US");
        assertEquals(5.99, rate.getAmount());

        wireMock.verify(getRequestedFor(urlEqualTo("/rates/US")));
    }
}
```
**A:** WireMock runs a real HTTP server that returns canned responses. Better than `MockRestServiceServer` — tests the full HTTP stack including headers, status codes, timeouts. Can simulate delays, partial responses, and network errors.

---

### 48. Testcontainers — Compose
**Q: How do you test with multiple services?**
```java
@Testcontainers
@SpringBootTest
class FullStackTest {
    @Container
    static DockerComposeContainer<?> compose = new DockerComposeContainer<>(
        new File("src/test/resources/docker-compose-test.yml"))
        .withExposedService("postgres", 5432)
        .withExposedService("redis", 6379)
        .withExposedService("kafka", 9092);

    // Spring properties point to compose services
    @DynamicPropertySource
    static void props(DynamicPropertyRegistry r) {
        r.add("spring.datasource.url", () ->
            "jdbc:postgresql://" + compose.getServiceHost("postgres", 5432) +
            ":" + compose.getServicePort("postgres", 5432) + "/testdb");
    }
}
```
**A:** `DockerComposeContainer` reads a `docker-compose.yml` file and starts all services together. Useful for testing service interactions (e.g., Kafka + Postgres + Redis) with proper networking between them.

---

## Section 5: Test Design — Patterns & Advanced Topics (Q49–Q55)

### 49. Test Doubles — Types
**Q: What is the difference between Dummy, Stub, Mock, Spy, and Fake?**
```java
// DUMMY: passed but never used (just to satisfy parameter requirements)
UserValidator dummy = null; // or mock(UserValidator.class) without expectations

// STUB: returns predefined values (no behavior verification)
when(userRepo.findById(1L)).thenReturn(Optional.of(user));

// MOCK: stub + expectation (verifies interactions)
when(sender.send(any())).thenReturn(true);
verify(sender).send(any()); // must be called

// SPY: real object, partial override (verifies calls on real objects)
List<String> spy = Mockito.spy(new ArrayList<>());
spy.add("real"); // real behavior

// FAKE: real, working, simplified implementation
class FakeUserRepo implements UserRepository {
    Map<Long, User> data = new HashMap<>();
    public User save(User u) { data.put(u.getId(), u); return u; }
    // other methods implemented simply...
}
```
**A:** Gerard Meszaros's taxonomy. In practice: Stub = stubbed mock without verification. Mock = stubbed + verified. Fake = in-memory implementation. Prefer Fakes for complex dependencies — more realistic than mocks, but requires maintaining the fake.

---

### 50. Test Pyramid
**Q: What is the recommended ratio?**
```
    /\
   /  \    E2E Tests (Selenium, REST) — few, slow, expensive
  /    \
 /------\  Integration Tests (Spring Boot) — some, medium
/        \
----------  Unit Tests (JUnit + Mockito) — many, fast, cheap

Google ratio: 70% unit / 20% integration / 10% E2E

Speed comparison:
  Unit test:   1–10ms
  Slice test:  100ms–2s (Spring context startup)
  Full integration: 5–30s
  E2E: 30s–5min

Design code for testability:
  - Constructor injection (not field injection)
  - Depend on interfaces (not concrete classes)
  - Side effects isolated to boundaries
  - Pure functions where possible
```
**A:** Unit tests are fast and locate bugs precisely. E2E tests give confidence but are slow and brittle. Maintain the pyramid shape — don't let integration tests dominate (slow feedback cycle).

---

### 51. Test Data Builders — Object Mother Pattern
**Q: How do you create consistent test data?**
```java
// Object Mother / Test Builder pattern
class UserTestData {
    public static User.Builder aUser() {
        return User.builder()
            .id(1L)
            .email("alice@test.com")
            .name("Alice")
            .active(true)
            .age(30);
    }

    public static User defaultUser()   { return aUser().build(); }
    public static User inactiveUser()  { return aUser().active(false).build(); }
    public static User adminUser()     { return aUser().role("ADMIN").build(); }
}

// Usage in tests:
@Test void test() {
    User user = UserTestData.inactiveUser();
    // No "magic" new User(1L, "alice@test.com", true, ..., ...) inline
}
```
**A:** Object Mother provides named factory methods for variants. Test Builders use fluent API for customization. Both avoid brittle `new User(null, null, 0, false, ...)` constructor calls that break when fields are added.

---

### 52. Contract Testing — Consumer-Driven Contracts
**Q: What problem does contract testing solve?**
```java
// Pact — consumer-driven contract testing
// Consumer (frontend) defines what it expects:
@Pact(consumer = "order-ui", provider = "order-service")
public RequestResponsePact orderPact(PactDslWithProvider builder) {
    return builder
        .uponReceiving("get order by ID")
        .path("/api/orders/1")
        .method("GET")
        .willRespondWith()
        .status(200)
        .body(new PactDslJsonBody()
            .integerType("id")
            .stringType("status")
            .decimalType("total"))
        .toPact();
}

@Test @PactVerification
void verifyPact() {
    ResponseEntity<Order> res = restTemplate.getForEntity("/api/orders/1", Order.class);
    assertEquals(200, res.getStatusCodeValue());
}
// Provider team runs: verify their service against consumer's pact
```
**A:** Contract testing catches integration breaks between services without deploying everything. Consumer defines the contract; provider verifies it. Pact Broker stores contracts. Runs in CI — fast and without a real environment.

---

### 53. Property-Based Testing — jqwik
**Q: What does property-based testing verify?**
```java
import net.jqwik.api.*;
import net.jqwik.api.constraints.*;

class PropertyTests {
    @Property
    void sortedListIsOrdered(@ForAll List<@IntRange(min=0, max=1000) Integer> list) {
        List<Integer> sorted = mergeSort(list);
        for (int i = 0; i < sorted.size() - 1; i++) {
            assertTrue(sorted.get(i) <= sorted.get(i + 1),
                "Not sorted at index " + i + ": " + sorted);
        }
    }

    @Property
    void serializeDeserializeIdentity(@ForAll @AlphaChars @NotEmpty String s) {
        assertEquals(s, deserialize(serialize(s)));
    }
}
// jqwik generates hundreds of random inputs and reports the minimal failing case
```
**A:** Property-based testing generates random inputs to find edge cases unit tests miss. jqwik shrinks failing cases to the minimal reproducing example. Excellent for: algorithms, serialization, parsers, mathematical properties.

---

### 54. Mutation Testing — PITest
**Q: What does mutation testing measure?**
```
Code coverage tells you which lines ran.
Mutation testing tells you if your tests actually caught bugs.

PItest introduces mutations:
  - Change > to >= (boundary condition)
  - Return null instead of a value
  - Remove a conditional
  - Negate a boolean

Then it runs your tests:
  - Mutation KILLED = test caught the change ✅
  - Mutation SURVIVED = test missed it ❌ (test gap!)

mvn test-compile org.pitest:pitest-maven:mutationCoverage

Mutation score = killed / total mutations
Goal: >80% mutation score (not just 80% line coverage)
```
**A:** PITest shows that 100% line coverage doesn't mean 100% correctness. A test that calls a method but doesn't assert the result provides no protection. Mutation testing quantifies test effectiveness.

---

### 55. Testing Anti-Patterns
**Q: Identify the anti-patterns in these tests:**
```java
// ANTI-PATTERN 1: Testing implementation, not behavior
@Test void test() {
    service.process(req);
    verify(helper, times(2)).internalMethod(); // who cares about internal calls?
}

// ANTI-PATTERN 2: Testing the mock, not the code
@Test void test() {
    when(repo.findById(1L)).thenReturn(Optional.of(user));
    assertEquals(user, repo.findById(1L).get()); // testing Mockito, not your code!
}

// ANTI-PATTERN 3: Giant test with multiple concerns
@Test void doEverything() {
    // 50 lines: create user, place order, process payment, send email...
    // When it fails, which part broke?
}

// ANTI-PATTERN 4: Flaky async test
@Test void asyncTest() {
    service.processAsync(event);
    Thread.sleep(500); // hope it's done by now
    verify(mock).result(any());
}
```
**A:** (1) Test observable behavior, not internal calls. (2) Don't test the mock — test your production code. (3) One assertion per test (or one concept). (4) Use `timeout()` or `await().atMost()` from Awaitility — never `Thread.sleep()` in tests.

---

> 🔖 **Last read:** <!-- update here -->
