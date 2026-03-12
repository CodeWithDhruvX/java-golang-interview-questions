## 🧪 Testing in Go (Questions 541-560)

### Question 541: How do you write table-driven tests in Go?

**Answer:**
Define a struct for test cases (input/expected), slice of that struct, and loop.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {1, 2, 3},
        {-1, 1, 0},
    }
    for _, tt := range tests {
        if got := Add(tt.a, tt.b); got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}
```

### Explanation
Table-driven tests in Go use a struct to define test case structure, create a slice of test cases, and iterate through them. This pattern eliminates code duplication and makes it easy to add new test cases. Each test case contains inputs and expected outputs, and the loop runs the same test logic against all cases, providing clear error messages when tests fail.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write table-driven tests in Go?
**Your Response:** "I write table-driven tests in Go by defining a struct that represents my test case structure with fields for inputs and expected outputs. I create a slice of these structs containing all my test cases, then loop through the slice running the same test logic against each case. This approach eliminates code duplication and makes it easy to add new test cases just by adding another struct to the slice. When a test fails, I get a clear error message showing exactly which inputs failed and what was expected versus what was actually returned. This pattern is especially useful for testing functions with multiple input combinations and edge cases, making the tests more maintainable and readable."

---

### Question 542: What is the difference between `t.Fatal` and `t.Errorf`?

**Answer:**
- **`t.Errorf`:** Log error and **continue** executing the test. Use when testing multiple independent assertions in a single test function.
- **`t.Fatal` (or `t.Fatalf`):** Log error and **stop** the test immediately. Use when a setup step fails (e.g., db connection) making further assertions impossible.

### Explanation
The difference between t.Errorf and t.Fatal is test continuation behavior. t.Errorf logs an error but allows the test to continue executing, useful for testing multiple independent assertions. t.Fatal logs an error and immediately stops the test, preventing any further assertions from running. t.Fatal is typically used for setup failures where continuing the test would be meaningless.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `t.Fatal` and `t.Errorf`?
**Your Response:** "The key difference is whether the test continues after an error. `t.Errorf` logs the error but continues executing the test, which I use when testing multiple independent assertions in a single test function - this way I can see all the failures at once. `t.Fatal` or `t.Fatalf` logs the error and immediately stops the test execution, preventing any further assertions from running. I use `t.Fatal` when a setup step fails, like a database connection that can't be established, because continuing the test would be meaningless without that setup. The choice depends on whether I want to see all possible failures or stop immediately when a critical failure occurs."

---

### Question 543: How do you use `go test -cover` to check coverage?

**Answer:**
- **Basic:** `go test -cover` gives a % summary.
- **Profile:** `go test -coverprofile=c.out` saves detailed data.
- **View:** `go tool cover -html=c.out` opens a web browser to visually see which lines are red (uncovered) vs green.

### Explanation
Go's test coverage tools provide different levels of detail. The basic `-cover` flag shows a percentage summary of code coverage. The `-coverprofile` flag saves detailed coverage data to a file. The `go tool cover -html` command generates an HTML visualization showing which lines are covered (green) and uncovered (red), making it easy to identify gaps in test coverage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `go test -cover` to check coverage?
**Your Response:** "I check test coverage in Go using several approaches. For a quick overview, I use `go test -cover` which gives me a percentage summary of how much code is covered by tests. For more detailed analysis, I use `go test -coverprofile=c.out` which saves detailed coverage data to a file. Then I run `go tool cover -html=c.out` which generates an HTML visualization that opens in my browser and shows exactly which lines are covered in green and which are uncovered in red. This visual approach makes it easy to identify specific areas of code that need more test coverage. I can also use the coverage profile to integrate with CI/CD pipelines and set coverage thresholds for code quality gates."

---

### Question 544: How do you mock a database in Go tests?

**Answer:**
1.  **DATA-DOG/go-sqlmock:** Mocks the `sql/driver` layer. Tests your SQL query syntax/arguments without a real DB.
2.  **Docker (Testcontainers):** Spin up a real Postgres for Integration Tests.
3.  **Interface:** Mock the `Repository` interface (application layer) rather than the database driver.

### Explanation
Database mocking in Go can be done at different levels. go-sqlmock mocks the SQL driver layer, allowing testing of SQL query syntax and arguments without a real database. Testcontainers spins up real databases in Docker for integration testing. Interface-based mocking focuses on the application repository layer rather than the database driver, providing better isolation and testability.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you mock a database in Go tests?
**Your Response:** "I mock databases in Go tests using three main approaches depending on what I need to test. For unit testing SQL queries, I use go-sqlmock which mocks the sql/driver layer and lets me verify query syntax and arguments without needing a real database. For integration testing, I use Testcontainers to spin up real databases like Postgres in Docker containers. For application-level testing, I prefer mocking the Repository interface rather than the database driver itself, which gives me better isolation and focuses on testing my business logic. The choice depends on whether I need to test SQL specifics, integration behavior, or application logic. Interface mocking is usually the cleanest approach for unit tests, while Testcontainers is great for more comprehensive integration testing."

---

### Question 545: How do you unit test HTTP handlers?

**Answer:**
Use `httptest.NewRecorder()` to capture the response.

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    
    MyHandler(w, req)
    
    resp := w.Result()
    if resp.StatusCode != 200 { t.Fail() }
}
```

### Explanation
HTTP handler unit testing in Go uses httptest.NewRecorder() to create a response recorder that captures the HTTP response. This allows testing handlers without starting a real HTTP server. The recorder captures status codes, headers, and body content, enabling assertions on the handler's behavior in isolation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you unit test HTTP handlers?
**Your Response:** "I unit test HTTP handlers in Go using the httptest package. I create a test request with httptest.NewRequest() and a response recorder with httptest.NewRecorder(). Then I call my handler directly with these test objects, which captures the HTTP response without needing a real server. After the handler executes, I can examine the response using w.Result() to check status codes, headers, and body content. This approach allows me to test handlers in isolation, making the tests fast and reliable. I can verify that my handlers return the correct status codes, set appropriate headers, and generate the expected response bodies for different request scenarios."

---

### Question 546: What is testable design and how does Go encourage it?

**Answer:**
Testable design favors **Dependency Injection** via Interfaces.
Go encourages this by having implicit interfaces. You define `type Reader interface` where you use it, making it easy to swap a real `os.File` with an in-memory `bytes.Buffer` for testing.

### Explanation
Testable design in Go emphasizes dependency injection through interfaces. Go's implicit interfaces allow defining interfaces where they're used rather than where they're implemented. This makes it easy to swap real implementations with test doubles, such as replacing os.File with bytes.Buffer for testing, without modifying the production code.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is testable design and how does Go encourage it?
**Your Response:** "Testable design in Go focuses on dependency injection through interfaces. What makes Go particularly good for this is its implicit interfaces - I can define an interface right where I use it, without needing the implementation to explicitly declare it. This makes it incredibly easy to swap real implementations with test doubles. For example, I can define a Reader interface where I need it, then use a real os.File in production but swap it for an in-memory bytes.Buffer in tests. Go's approach means I don't need complex dependency injection frameworks - the language itself provides the mechanisms for writing testable code through interfaces. This design philosophy encourages loose coupling and makes testing much more straightforward."

---

### Question 547: How do you use interfaces to improve testability?

**Answer:**
Instead of accepting a concrete type `*Service`, accept an interface `ServiceAPI`.

```go
type Downloader interface { Download(url string) ([]byte, error) }

func Process(d Downloader) { ... }

// Test
type MockDownloader struct{}
func (m MockDownloader) Download(u string) ([]byte, error) { return []byte("fake"), nil }
// Now you can test Process() without internet.
```

### Explanation
Using interfaces for testability involves accepting interface types instead of concrete types. This allows swapping real implementations with mock implementations during testing. The example shows a Downloader interface that can be implemented by both real HTTP clients and mock implementations, enabling testing without external dependencies like network access.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use interfaces to improve testability?
**Your Response:** "I improve testability by accepting interface types instead of concrete types in my functions. Instead of taking a `*Service`, I define an interface like `ServiceAPI` and accept that. This allows me to create mock implementations for testing. For example, I can define a `Downloader` interface with a `Download` method, then my `Process` function accepts this interface. In production, I pass a real downloader that makes HTTP requests, but in tests, I pass a `MockDownloader` that returns fake data without needing internet access. This approach makes my code completely testable in isolation without external dependencies, and Go's implicit interfaces make this pattern very natural to implement."

---

### Question 548: How do you write tests for concurrent code in Go?

**Answer:**
1.  **Race Detector:** Always run with `go test -race`.
2.  **Stress Testing:** Run multiple iterations or use parallel subtests.
3.  **Synchronization:** Use `WaitGroup` or channels in the test to ensure all goroutines complete before checking assertions.

### Explanation
Testing concurrent code requires special approaches. The race detector (`go test -race`) identifies data races between goroutines. Stress testing with multiple iterations or parallel subtests helps uncover concurrency bugs. Synchronization primitives like WaitGroup or channels ensure all goroutines complete before assertions are checked, preventing false test failures.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write tests for concurrent code in Go?
**Your Response:** "I test concurrent code in Go using three key techniques. First, I always run tests with the race detector using `go test -race` which identifies data races between goroutines that could cause unpredictable behavior. Second, I perform stress testing by running multiple iterations or using parallel subtests to increase the likelihood of exposing concurrency bugs. Third, I use synchronization primitives like WaitGroup or channels in my tests to ensure all goroutines complete before I check my assertions. This prevents false failures where the test finishes before the concurrent code has had time to execute. The combination of race detection, stress testing, and proper synchronization helps me write reliable tests for concurrent code that catch real issues without producing flaky results."

---

### Question 549: What is the `httptest` package and how is it used?

**Answer:**
It provides utilities for HTTP testing.
- `ResponseRecorder`: Records writes to `ResponseWriter`.
- `Server`: Starts a real (local) HTTP server on a random port for end-to-end testing of HTTP Clients.
    ```go
    ts := httptest.NewServer(http.HandlerFunc(func(w, r) { fmt.Fprintln(w, "Hello") }))
    defer ts.Close()
    client.Get(ts.URL)
    ```

### Explanation
The httptest package provides utilities for HTTP testing. ResponseRecorder captures HTTP responses for handler testing without a real server. Server starts a local HTTP server on a random port for end-to-end testing of HTTP clients, allowing testing against real HTTP behavior while keeping tests isolated and fast.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the `httptest` package and how is it used?
**Your Response:** "The httptest package provides utilities for HTTP testing in Go. It has two main components: ResponseRecorder for testing HTTP handlers and Server for testing HTTP clients. ResponseRecorder lets me capture HTTP responses without starting a real server - I create a recorder, call my handler with it, and then examine the captured response. Server lets me start a real local HTTP server on a random port for end-to-end testing of HTTP clients. This is perfect for testing client code against real HTTP behavior while keeping tests isolated and fast. I use ResponseRecorder for unit testing handlers and Server for integration testing HTTP clients, giving me comprehensive coverage of HTTP functionality."

---

### Question 550: How do you mock time in tests?

**Answer:**
Avoid `time.Now()`. Inject a Clock interface.

```go
type Clock interface { Now() time.Time }
type App struct { Clock Clock }

func (a *App) IsWeekEnd() bool {
    return a.Clock.Now().Weekday() == time.Sunday
}
```
In tests, pass a `MockClock` that returns a fixed static time.

### Explanation
Mocking time in tests involves avoiding direct calls to time.Now() and instead injecting a Clock interface. This allows controlling time during tests by providing mock implementations that return fixed times, enabling deterministic testing of time-dependent behavior without waiting for actual time to pass.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you mock time in tests?
**Your Response:** "I mock time in tests by avoiding direct calls to `time.Now()` and instead injecting a Clock interface. I define a simple Clock interface with a Now() method, and my application struct takes this interface as a dependency. In production, I pass a real clock implementation that calls time.Now(), but in tests, I pass a MockClock that returns a fixed static time. This approach makes time-dependent code deterministic and testable - I can test weekend logic by setting the clock to return a Sunday, or test expiration logic by controlling exactly what time is returned. This eliminates flaky tests caused by actual time changes and makes testing time-sensitive behavior much more reliable and fast."

---

### Question 551: How do you perform integration testing in Go?

**Answer:**
Integration tests interact with external systems (DB, Redis).
1.  Use build tags `// +build integration` to separate them from unit tests.
2.  Run with `go test -tags=integration ./...`.
3.  Use `testcontainers-go` to spin up dependencies in Docker automatically.

### Explanation
Integration testing in Go involves testing against external systems like databases and Redis. Build tags separate integration tests from unit tests, allowing selective execution. Testcontainers-go automates spinning up dependencies in Docker containers, providing isolated and reproducible integration testing environments.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you perform integration testing in Go?
**Your Response:** "I perform integration testing in Go by testing against real external systems like databases and Redis. I use build tags with `// +build integration` to separate integration tests from unit tests, allowing me to run them selectively with `go test -tags=integration ./...`. This keeps my unit tests fast and isolated while still having comprehensive integration coverage. For managing external dependencies, I use testcontainers-go which automatically spins up services like databases in Docker containers. This approach gives me real integration testing with actual databases while keeping the environment isolated and reproducible. The combination of build tags and testcontainers allows me to have the best of both worlds - fast unit tests for most logic and thorough integration tests for critical paths."

---

### Question 552: How do you use `testify/mock` for mocking dependencies?

**Answer:**
`testify` is a popular assertion/mocking library.

```go
type MyMock struct { mock.Mock }
func (m *MyMock) Do(n int) int {
    args := m.Called(n)
    return args.Int(0)
}

// Test
m := new(MyMock)
m.On("Do", 5).Return(10)
result := m.Do(5)
m.AssertExpectations(t)
```

### Explanation
testify provides assertion and mocking capabilities for Go testing. The mock.Mock type enables creating mock implementations with expected method calls and return values. The On() method sets up expectations for specific method calls with specific arguments, and AssertExpectations() verifies all expectations were met.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `testify/mock` for mocking dependencies?
**Your Response:** "I use testify/mock for creating mock implementations of dependencies. I create a mock struct that embeds mock.Mock and implements the interface I want to mock. In my test, I set up expectations using the On() method to specify what method should be called with what arguments and what it should return. Then I call my code under test, which uses the mock, and finally use AssertExpectations() to verify all expected calls were made. This approach gives me fine-grained control over mock behavior and allows me to test edge cases and error conditions that might be difficult to reproduce with real implementations. The library provides a clean, fluent API for setting up mocks and verifying interactions, making it much easier to write comprehensive unit tests."

---

### Question 553: How do you run subtests and benchmarks?

**Answer:**
Use `t.Run()` or `b.Run()`.

```go
func TestMath(t *testing.T) {
    t.Run("Add", func(t *testing.T) { ... })
    t.Run("Sub", func(t *testing.T) { ... })
}
```
Run specific subtest: `go test -run TestMath/Add`.

### Explanation
Subtests in Go are created using t.Run() for tests and b.Run() for benchmarks. This allows organizing related tests hierarchically and running specific subtests individually. Subtests provide better test organization and selective execution, making it easier to focus on specific functionality during development and debugging.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you run subtests and benchmarks?
**Your Response:** "I create subtests and benchmarks using `t.Run()` for tests and `b.Run()` for benchmarks. This allows me to organize related tests hierarchically within a single test function. For example, I can have a TestMath function with subtests for Add and Sub operations. The benefit is that I can run specific subtests individually using patterns like `go test -run TestMath/Add`, which is great for debugging specific issues. Subtests also provide better test organization and make test output more readable. Each subtest runs independently, so failures in one don't affect others, and I can even run them in parallel if needed. This approach is particularly useful for testing complex functions with multiple scenarios or when I want to group related tests together."

---

### Question 554: How do you test panic recovery?

**Answer:**
Use `defer` inside the test to catch the panic.

```go
func TestPanic(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()
    CausePanic()
}
```
*(Or use `assert.Panics` from testify)*

### Explanation
Testing panic recovery in Go involves using defer with recover() to catch panics. The defer function checks if a panic occurred and fails the test if no panic was expected. testify provides a more convenient assert.Panics function for the same purpose. This approach allows testing that code panics under expected conditions.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test panic recovery?
**Your Response:** "I test panic recovery using a defer function with recover() to catch expected panics. I wrap the code that should panic in a defer function that checks if a panic occurred using recover(). If no panic is found when one was expected, I fail the test. Alternatively, I can use the more convenient assert.Panics function from the testify library which handles this pattern automatically. This approach allows me to verify that my code panics under the right conditions, like when invalid inputs are provided or critical errors occur. It's important for testing error paths and ensuring that fail-fast behavior works correctly in production code."

---

### Question 555: How do you generate test data using faker or random data?

**Answer:**
Use libraries like `github.com/go-faker/faker` or `github.com/brianvoe/gofakeit`.
This generates random specific data (Emails, Names, Addresses) to test edge cases or validation logic.

```go
var u User
faker.FakeData(&u) // Fills struct with random valid data
```

### Explanation
Generating test data in Go can be done using libraries like faker or gofakeit that create realistic random data. These libraries can generate emails, names, addresses, and other structured data to test edge cases and validation logic. The data can be used to fill structs directly, making it easy to create comprehensive test scenarios.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you generate test data using faker or random data?
**Your Response:** "I generate test data using libraries like `github.com/go-faker/faker` or `github.com/brianvoe/gofakeit` which create realistic random data for testing. These libraries can generate emails, names, addresses, phone numbers, and other structured data that looks real but is deterministic enough for testing. I use them to fill structs directly with `faker.FakeData(&user)` which populates all fields with valid random values. This is particularly useful for testing validation logic, edge cases, and ensuring my code handles various input formats correctly. Instead of manually creating test data, I can generate comprehensive test scenarios automatically, making my tests more thorough and easier to maintain."

---

### Question 556: What is golden file testing and when is it useful?

**Answer:**
Useful for comparing large outputs (JSON, HTML).
1.  Run test.
2.  If `update` flag is set, write output to standard file (`testdata/output.golden`).
3.  If not set, read `output.golden` and `bytes.Equal(got, expected)`.
Ensures complex output remains stable.

### Explanation
Golden file testing is useful for comparing large outputs like JSON or HTML. The approach involves running tests and either updating golden files with current output or comparing current output against existing golden files. This ensures complex output remains stable over time and makes it easy to detect unintended changes in generated content.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is golden file testing and when is it useful?
**Your Response:** "Golden file testing is a technique I use for comparing large outputs like JSON or HTML. The approach involves running tests and either updating golden files with current output when an update flag is set, or comparing the current output against existing golden files. I store the expected output in files like `testdata/output.golden` and use `bytes.Equal()` to compare actual results. This is particularly useful for testing complex generators, template rendering, or serialization where manually writing expected values would be cumbersome. When I make intentional changes, I run the test with the update flag to refresh the golden files. This ensures that complex output remains stable over time and makes it easy to detect unintended changes in generated content."

---

### Question 557: How do you automate test workflows with `go generate`?

**Answer:**
`go generate` runs commands embedded in comments.
Use it to generate mocks before testing.
`//go:generate mockery --name=UserService`
Run: `go generate ./...` -> `go test ./...`

### Explanation
go generate runs commands embedded in Go source file comments. This is commonly used to generate mocks before running tests. The mockery command can generate mock implementations from interfaces, and the process involves running go generate to create the mocks, then running tests that use those generated mocks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you automate test workflows with `go generate`?
**Your Response:** "I automate test workflows with `go generate` by embedding commands in Go source comments. I commonly use it to generate mocks before running tests. For example, I'll add `//go:generate mockery --name=UserService` in my code, which tells go generate to create mock implementations for my UserService interface. The workflow involves running `go generate ./...` to create all the generated files, then running `go test ./...` to execute tests that use those generated mocks. This approach automates the tedious process of maintaining mock implementations and ensures they stay in sync with interface changes. It's particularly useful in large codebases with many interfaces that need mocking for comprehensive testing."

---

### Question 558: How do you test CLI apps built with Cobra?

**Answer:**
Set custom input/output buffers on the Root Command.

```go
buf := new(bytes.Buffer)
rootCmd.SetOut(buf)
rootCmd.SetArgs([]string{"create", "--name", "test"})
rootCmd.Execute()

if !strings.Contains(buf.String(), "Created") { t.Fail() }
```

### Explanation
Testing CLI applications built with Cobra involves setting custom input/output buffers on the root command. This allows capturing command output for assertions without actually writing to stdout. The SetOut method redirects output to a buffer, SetArgs simulates command line arguments, and Execute runs the command with the test configuration.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test CLI apps built with Cobra?
**Your Response:** "I test CLI applications built with Cobra by setting custom input/output buffers on the root command. I create a bytes.Buffer to capture output, then use SetOut() to redirect the command's output to this buffer instead of stdout. I use SetArgs() to simulate command line arguments like `['create', '--name', 'test']`. Then I call Execute() to run the command with the test configuration. Finally, I can examine the buffer contents to verify the command produced the expected output. This approach allows me to test CLI functionality end-to-end without actually writing to the console or depending on external files, making the tests fast and reliable."

---

### Question 559: What is fuzz testing and how do you do it in Go?

**Answer:**
Go 1.18+ Fuzzing. Uses a genetic algorithm to mutate inputs to trigger crashes.

```go
func FuzzReverse(f *testing.F) {
    f.Fuzz(func(t *testing.T, s string) {
        // Assert property: Reverse(Reverse(s)) == s
        if Reverse(Reverse(s)) != s { t.Error("Failed") }
    })
}
```

### Explanation
Fuzz testing in Go (available since Go 1.18) uses genetic algorithms to mutate inputs automatically to find crashes and bugs. The fuzzer generates random variations of input data and tests them against the target code, looking for panics, assertions failures, or other unexpected behaviors. This is particularly effective for finding edge cases that manual testing might miss.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is fuzz testing and how do you do it in Go?
**Your Response:** "Fuzz testing in Go, available since version 1.18, is an automated testing technique that uses genetic algorithms to find bugs. I create fuzz tests using the FuzzReverse pattern where I define a fuzz function that takes random inputs and tests properties or assertions. The fuzzer automatically generates and mutates input data, looking for crashes, panics, or assertion failures. For example, I might test that reversing a string twice returns the original string. The fuzzer tries countless variations including edge cases I might never think of manually. This approach is incredibly effective at finding subtle bugs and edge cases in parsing, validation, or data processing code that traditional testing might miss."

---

### Question 560: How do you organize test files and test suites?

**Answer:**
- **Unit Tests:** Inside the same package (`package foo`), file `foo_test.go`. Access to private members.
- **Blackbox Tests:** Use `package foo_test`. Import `foo`. Tests only public API. (Preferred for integration/API tests).
- **Testdata:** Store fixtures in a `testdata/` directory (ignored by compiler).
- **Suite:** Use `testify/suite` for Setup/Teardown logic shared across tests.

### Explanation
Go test organization follows several conventions. Unit tests in the same package (`package foo`) can access private members and are placed in `foo_test.go` files. Blackbox tests use `package foo_test` and can only test public API. Test fixtures are stored in a `testdata/` directory. Testify suites provide setup/teardown functionality shared across multiple tests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you organize test files and test suites?
**Your Response:** "I organize Go tests using several conventions. For unit tests that need access to private members, I use `package foo` and place tests in `foo_test.go` files. For blackbox tests that only exercise the public API, I use `package foo_test` and import the package under test. I store test fixtures and sample data in a `testdata/` directory, which the Go compiler ignores. For shared setup and teardown logic across multiple tests, I use `testify/suite` which provides a structured way to organize test suites. This organization gives me flexibility - unit tests can be fast and have full access, while blackbox tests ensure the public API works correctly. The testdata directory keeps test assets organized and out of the main codebase."

---
