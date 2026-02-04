## 🔴 Tools, Testing, CI/CD, Ecosystem (Questions 181-230)

### Question 181: What is go vet and what does it catch?

**Answer:**
`go vet` examines Go source code and reports suspicious constructs:

```bash
go vet ./...
```

**What it catches:**
1. **Incorrect Printf formats:**
   ```go
   fmt.Printf("%d", "string")  // Caught: format %d expects int
   ```

2. **Unreachable code:**
   ```go
   return
   fmt.Println("This will never run")  // Caught
   ```

3. **Incorrect struct tags:**
   ```go
   type User struct {
       Name string `json:"name"extra"`  // Caught: malformed tag
   }
   ```

4. **Copying locks:**
   ```go
   var mu sync.Mutex
   mu2 := mu  // Caught: copying lock value
   ```

5. **Invalid method signatures:**
   ```go
   func (t *T) String(x int) string {  // Caught: should be String() string
       return ""
   }
   ```

**Usage in CI/CD:**
```yaml
# .github/workflows/test.yml
- name: Run go vet
  run: go vet ./...
```

---

### Question 182: How does go fmt help maintain code quality?

**Answer:**
`go fmt` automatically formats Go code to follow standard style:

```bash
# Format all files in current directory:
go fmt ./...

# Format specific file:
go fmt main.go

# Check what would be formatted (dry run):
gofmt -l .
```

**What it does:**
- Fixes indentation
- Aligns code blocks  
- Standardizes spacing
- Organizes imports

**Example:**
```go
// Before:
func main( ){
if   true{
fmt.Println( "hello" )
}
}

// After go fmt:
func main() {
    if true {
        fmt.Println("hello")
    }
}
```

**Benefits:**
- Eliminates style debates
- Makes code reviews easier
- Consistent codebase
- No configuration needed

**IDE Integration:**
```json
// VS Code settings.json
{
    "go.formatTool": "gofmt",
    "[go]": {
        "editor.formatOnSave": true
    }
}
```

---

### Question 183: What is golangci-lint?

**Answer:**
golangci-lint runs multiple linters in parallel:

```bash
# Install:
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run:
golangci-lint run

# Run specific linters:
golangci-lint run --enable=golint,errcheck,staticcheck
```

**Configuration (.golangci.yml):**
```yaml
linters:
  enable:
    - gofmt
    - govet
    - errcheck
    - staticcheck
    - unused
    - gosimple
    - structcheck
    - varcheck
    - ineffassign
    - deadcode
    
linters-settings:
  govet:
    check-shadowing: true
  gocyclo:
    min-complexity: 15
    
run:
  timeout: 5m
  tests: true
```

**Built-in linters:**
- **errcheck** - Checks for unchecked errors
- **staticcheck** - Advanced static analysis
- **unused** - Finds unused code
- **gosimple** - Simplification suggestions
- **govet** - Go vet checks
- **ineffassign** - Ineffectual assignments

---

### Question 184: What is the difference between go run, go build, and go install?

**Answer:**

**go run** - Compiles and runs immediately (no binary saved):
```bash
go run main.go
# Useful for: Quick testing, scripts
```

**go build** - Compiles and creates binary in current directory:
```bash
go build -o myapp
# Creates: ./myapp
# Useful for: Local testing, deployment packages
```

**go install** - Compiles and installs binary to $GOPATH/bin:
```bash
go install
# Creates: $GOPATH/bin/myapp
# Useful for: CLI tools, globally available commands
```

**Comparison:**
```go
// main.go
package main
import "fmt"
func main() {
    fmt.Println("Hello")
}
```

```bash
# go run - runs immediately:
go run main.go
# Output: Hello
# No binary created

# go build - creates binary:
go build -o hello
./hello
# Output: Hello
# Binary: ./hello

# go install - installs globally:
go install
$GOPATH/bin/hello
# Output: Hello
# Binary: $GOPATH/bin/hello
```

**Build flags:**
```bash
# Optimize for size:
go build -ldflags="-s -w" -o app

# Cross-compile:
GOOS=linux GOARCH=amd64 go build -o app-linux

# Build with race detector:
go build -race -o app
```

---

### Question 185: How does go generate work?

**Answer:**
`go generate` runs commands specified in source files:

**Usage:**
```go
// math.go
package math

//go:generate stringer -type=Operation
type Operation int

const (
    Add Operation = iota
    Subtract
    Multiply
    Divide
)
```

```bash
# Run code generation:
go generate ./...

# This executes:
# stringer -type=Operation
# Which generates: operation_string.go
```

**Common use cases:**

**1. Generate mocks:**
```go
//go:generate mockgen -source=interface.go -destination=mocks/mock_interface.go
type UserService interface {
    GetUser(id int) (*User, error)
}
```

**2. Generate protobuf:**
```go
//go:generate protoc --go_out=. --go_grpc_out=. user.proto
```

**3. Embed files:**
```go  
//go:generate go run assets/generate.go

//go:embed templates/*.html
var templates embed.FS
```

**4. Generate enums:**
```go
//go:generate enumer -type=Status -json
type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)
```

---

### Question 186: What is a build constraint?

**Answer:**
Build constraints (build tags) control which files are compiled:

**Syntax:**
```go
//go:build linux
// +build linux

package main

func platformSpecific() {
    // Linux-only code
}
```

**Multiple tags:**
```go
//go:build linux && amd64
// Linux AMD64 only

//go:build darwin || linux  
// macOS or Linux

//go:build !windows
// Everything except Windows
```

**Use cases:**

**1. Platform-specific code:**
```go
// file_unix.go
//go:build unix

func openFile() {
    // Unix implementation
}

// file_windows.go
//go:build windows

func openFile() {
    // Windows implementation
}
```

**2. Feature flags:**
```go
//go:build debug

func enableDebugLogging() {
    // Debug-only code
}
```

```bash
# Build with tag:
go build -tags debug
```

**3. Integration tests:**
```go
//go:build integration

func TestDatabaseIntegration(t *testing.T) {
    // Only runs with: go test -tags integration
}
```

---

###Question 187: How do you write tests in Go?

**Answer:**
Create `_test.go` files with Test functions:

**Basic test:**
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go  
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

**Table-driven tests:**
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive numbers", 2, 3, 5},
        {"with zero", 0, 5, 5},
        {"negative numbers", -2, -3, -5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

**Subtests:**
```go
func TestUser(t *testing.T) {
    t.Run("Creation", func(t *testing.T) {
        user := NewUser("John")
        if user.Name != "John" {
            t.Error("Name not set correctly")
        }
    })
    
    t.Run("Validation", func(t *testing.T) {
        user := NewUser("")
        if user.IsValid() {
            t.Error("Empty name should be invalid")
        }
    })
}
```

**Run tests:**
```bash
go test                    # Current package
go test ./...              # All packages
go test -v                 # Verbose
go test -run TestAdd       # Specific test
go test -cover             # With coverage
```

---

### Question 188: How do you test for expected panics?

**Answer:**
Use `defer` with `recover`:

```go
func TestPanicRecovery(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("Expected panic, but didn't panic")
        }
    }()
    
    // Code that should panic:
    causePanic()
}

// Better approach with helper:
func assertPanic(t *testing.T, f func()) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("Expected panic")
        }
    }()
    f()
}

func TestDivideByZero(t *testing.T) {
    assertPanic(t, func() {
        Divide(10, 0)  // Should panic
    })
}

// Using testify package:
import "github.com/stretchr/testify/assert"

func TestPanic(t *testing.T) {
    assert.Panics(t, func() {
        causePanic()
    }, "Function should panic")
    
    assert.NotPanics(t, func() {
        safeFunction()
    }, "Function should not panic")
}
```

---

### Question 189: What are mocks and how do you use them in Go?

**Answer:**
Mocks simulate dependencies for testing:

**Manual mocks:**
```go
// Interface to mock:
type EmailSender interface {
    Send(to, subject, body string) error
}

// Mock implementation:
type MockEmailSender struct {
    SendCalled bool
    SendError  error
}

func (m *MockEmailSender) Send(to, subject, body string) error {
    m.SendCalled = true
    return m.SendError
}

// Test using mock:
func TestUserService(t *testing.T) {
    mockEmail := &MockEmailSender{}
    service := NewUserService(mockEmail)
    
    service.RegisterUser("john@example.com")
    
    if !mockEmail.SendCalled {
        t.Error("Expected Email.Send to be called")
    }
}
```

**Using testify/mock:**
```go
import "github.com/stretchr/testify/mock"

type MockEmailSender struct {
    mock.Mock
}

func (m *MockEmailSender) Send(to, subject, body string) error {
    args := m.Called(to, subject, body)
    return args.Error(0)
}

func TestWithTestify(t *testing.T) {
    mockEmail := new(MockEmailSender)
    
    // Set expectations:
    mockEmail.On("Send", "john@example.com", "Welcome", mock.Anything).
        Return(nil)
    
    service := NewUserService(mockEmail)
    service.RegisterUser("john@example.com")
    
    // Assert expectations were met:
    mockEmail.AssertExpectations(t)
}
```

**Using mockgen (gomock):**
```bash
go install github.com/golang/mock/mockgen@latest

mockgen -source=interface.go -destination=mocks/mock_interface.go
```

```go
//go:generate mockgen -source=user_service.go -destination=mocks/mock_user_service.go
type UserService interface {
    GetUser(id int) (*User, error)
}

// In tests:
func TestWithGomock(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockService := mocks.NewMockUserService(ctrl)
    mockService.EXPECT().
        GetUser(1).
        Return(&User{ID: 1, Name: "John"}, nil)
    
    user, _ := mockService.GetUser(1)
    if user.Name != "John" {
        t.Error("Expected John")
    }
}
```

---

### Question 190: How do you use the testing and testify packages?

**Answer:**

**Standard testing package:**
```go
import "testing"

func TestBasic(t *testing.T) {
    t.Log("Starting test")
    t.Error("This test failed")
    t.Skip("Skipping this test")
    t.Fatal("Stop execution")
}
```

**testify/assert - Convenient assertions:**
```go
import "github.com/stretchr/testify/assert"

func TestWithAssert(t *testing.T) {
    // Equality:
    assert.Equal(t, 5, result)
    assert.NotEqual(t, 0, count)
    
    // Nil checks:
    assert.Nil(t, err)
    assert.NotNil(t, user)
    
    // Boolean:
    assert.True(t, condition)
    assert.False(t, flag)
    
    // Strings:
    assert.Contains(t, "hello world", "world")
    assert.Empty(t, "")
    
    // Collections:
    assert.Len(t, list, 5)
    assert.ElementsMatch(t, []int{1, 2, 3}, result)
    
    // Errors:
    assert.NoError(t, err)
    assert.Error(t, err)
}
```

**testify/require - Fail fast:  **
```go
import "github.com/stretchr/testify/require"

func TestWithRequire(t *testing.T) {
    user, err := GetUser(1)
    require.NoError(t, err)  // Stops if error
    require.NotNil(t, user)  // Stops if nil
    
    // Continue only if above passed:
    assert.Equal(t, "John", user.Name)
}
```

**testify/suite - Test suites:**
```go
import "github.com/stretchr/testify/suite"

type UserTestSuite struct {
    suite.Suite
    db *sql.DB
}

func (s *UserTestSuite) SetupSuite() {
    // Runs once before all tests
    s.db = setupDB()
}

func (s *UserTestSuite) TearDownSuite() {
    // Runs once after all tests
    s.db.Close()
}

func (s *UserTestSuite) SetupTest() {
    // Runs before each test
    s.db.Exec("TRUNCATE users")
}

func (s *UserTestSuite) TestCreateUser() {
    user := &User{Name: "John"}
    err := CreateUser(s.db, user)
    
    s.NoError(err)
    s.Greater(user.ID, 0)
}

func TestUserSuite(t *testing.T) {
    suite.Run(t, new(UserTestSuite))
}
```

---

### Question 191: How do you structure test files in Go?

**Answer:**

**Project structure:**
```
myproject/
├── user/
│   ├── user.go           # Source code
│   ├── user_test.go      # Unit tests (same package)
│   └── user_integration_test.go  # Integration tests
├── testdata/             # Test fixtures
│   ├── users.json
│   └── config.yaml
└── tests/
    └── e2e_test.go       # End-to-end tests (separate package)
```

**File naming:**
```
user.go              # Implementation
user_test.go         # Tests (package user)
user_internal_test.go  # Tests (package user_test)
```

**Same package vs separate:**
```go
// user_test.go (same package - can test private members)
package user

func TestInternalMethod(t *testing.T) {
    u := &User{}
    u.internalMethod()  // Can access private members
}

// user_external_test.go (separate package - only public API)
package user_test

import "myproject/user"

func TestPublicAPI(t *testing.T) {
    u := user.NewUser()  // Only public members
}
```

**Test organization:**
```go
// unit_test.go - Fast unit tests
func TestAdd(t *testing.T) { }
func TestSubtract(t *testing.T) { }

// integration_test.go - Slower integration tests
//go:build integration

func TestDatabaseIntegration(t *testing.T) { }

// benchmark_test.go - Benchmarks
func BenchmarkAdd(b *testing.B) { }
```

**Shared test helpers:**
```go
// testing.go
package user

// Test helpers (not exported):
func createTestUser(t *testing.T) *User {
    t.Helper()  // Mark as helper
    return &User{Name: "Test User"}
}

// test_fixtures.go
func loadTestData(t *testing.T, filename string) []byte {
    data, err := os.ReadFile(filepath.Join("testdata", filename))
    if err != nil {
        t.Fatal(err)
    }
    return data
}
```

---

### Question 192: What is a benchmark test?

**Answer:**
Benchmark tests measure performance:

**Basic benchmark:**
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

**Run benchmarks:**
```bash
go test -bench=.                  # Run all benchmarks
go test -bench=Add               # Specific benchmark
go test -bench=. -benchmem       # With memory stats
go test -bench=. -cpuprofile=cpu.prof  # CPU profile
```

**More examples:**
```go
// Table-driven benchmarks:
func BenchmarkStringOperations(b *testing.B) {
    tests := []struct{
        name string
        fn   func(string) string
    }{
        {"concat", func(s string) string { return s + s }},
        {"builder", func(s string) string {
            var builder strings.Builder
            builder.WriteString(s)
            builder.WriteString(s)
            return builder.String()
        }},
    }
    
    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                tt.fn("hello")
            }
        })
    }
}

// Reset timer for setup:
func BenchmarkWithSetup(b *testing.B) {
    // Expensive setup:
    data := generateLargeDataset()
    
    b.ResetTimer()  // Don't count setup time
    
    for i := 0; i < b.N; i++ {
        process(data)
    }
}

// Parallel benchmarks:
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // Work that can run in parallel
            expensiveOperation()
        }
    })
}

// Memory allocations:
func BenchmarkAllocations(b *testing.B) {
    b.ReportAllocs()  // Report memory allocations
    
    for i := 0; i < b.N; i++ {
        _ = make([]int, 1000)
    }
}
```

**Output:**
```
BenchmarkAdd-8          1000000000     0.25 ns/op
BenchmarkConcat-8         5000000      250 ns/op    48 B/op   1 allocs/op
```

---

### Question 193: How do you measure test coverage in Go?

**Answer:**
Use built-in coverage tools:

```bash
# Run tests with coverage:
go test -cover

# Generate coverage profile:
go test -coverprofile=coverage.out

# View coverage in terminal:
go tool cover -func=coverage.out

# Generate HTML report:
go tool cover -html=coverage.out

# Coverage for specific packages:
go test ./... -coverprofile=coverage.out

# Coverage modes:
go test -covermode=set       # Boolean (covered or not)
go test -covermode=count     # Count how many times
go test -covermode=atomic    # Thread-safe count
```

**Example output:**
```
math/add.go:5:  Add     100.0%
math/sub.go:5:  Sub     80.0%
total:                  90.0%
```

**CI/CD integration:**
```yaml
# .github/workflows/test.yml
- name: Run tests with coverage
  run: go test -v -coverprofile=coverage.out ./...
 
- name: Upload coverage
  uses: codecov/codecov-action@v3
  with:
    files: ./coverage.out
```

**Makefile:**
```makefile
test-coverage:
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html
    @echo "Coverage report: coverage.html"
```

---

### Question 194: How do you test concurrent functions?

**Answer:**
Testing concurrency requires special care:

**Basic concurrent test:**
```go
func TestConcurrent Counter(t *testing.T) {
    counter := NewSafeCounter()
    var wg sync.WaitGroup
    
    // Start multiple goroutines:
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    
    if counter.Value() != 100 {
        t.Errorf("Expected 100, got %d", counter.Value())
    }
}
```

**Race detector:**
```bash
go test -race ./...
```

**Testing channels:**
```go
func TestChannel(t *testing.T) {
    ch := make(chan int, 1)
    
    go func() {
        ch <- 42
    }()
    
    select {
    case val := <-ch:
        assert.Equal(t, 42, val)
    case <-time.After(time.Second):
        t.Fatal("Timeout waiting for value")
    }
}
```

**Testing goroutine leaks:**
```go
import "go.uber.org/goleak"

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}

func TestNoGoroutineLeak(t *testing.T) {
    defer goleak.VerifyNone(t)
    
    // Code that should not leak goroutines
    startWorker()
}
```

**Stress testing:**
```go
func TestConcurrentWrites(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping stress test in short mode")
    }
    
    m := &sync.Map{}
    done := make(chan bool)
    
    // Multiple writers:
    for i := 0; i < 100; i++ {
        go func(id int) {
            for j := 0; j < 1000; j++ {
                m.Store(fmt.Sprintf("%d-%d", id, j), id)
            }
            done <- true
        }(i)
    }
    
    // Wait for all:
    for i := 0; i < 100; i++ {
        <-done
    }
}
```

---

### Question 195: What is a race detector and how do you use it?

**Answer:**
The race detector finds data races at runtime:

**Usage:**
```bash
# Run tests with race detector:
go test -race ./...

# Run program with race detector:
go run -race main.go

# Build with race detector:
go build -race -o myapp
```

**Example race condition:**
```go
// BAD - Race condition:
var counter int

func increment() {
    counter++  // RACE!
}

func TestRace(t *testing.T) {
    for i := 0; i < 1000; i++ {
        go increment()
    }
    time.Sleep(time.Second)
}
```

**Running with race detector:**
```bash
$ go test -race
==================
WARNING: DATA RACE
Write at 0x00c000018090 by goroutine 8:
  main.increment()
      /path/to/file.go:5 +0x44

Previous read at 0x00c000018090 by goroutine 7:
  main.increment()
      /path/to/file.go:5 +0x3a
==================
```

**Fixed version:**
```go
// GOOD - No race:
var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

// Or use atomic:
var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}
```

**Common race patterns:**
```go
// 1. Map races:
m := make(map[string]int)
go func() { m["key"] = 1 }()
go func() { m["key"] = 2 }()  // RACE!

// Fix: Use sync.Map or mutex

// 2. Slice races:
s := []int{}
go func() { s = append(s, 1) }()
go func() { s = append(s, 2) }()  // RACE!

// Fix: Use channels or mutex

// 3. Closure variable:
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // RACE with loop variable!
    }()
}

// Fix: Pass as parameter:
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println(n)
    }(i)
}
```

---

### Question 196: What is go.mod and go.sum?

**Answer:**

**go.mod** - Defines module and dependencies:
```go
module github.com/user/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/lib/pq v1.10.9
)

require (
    // Indirect dependencies (transitive)
    github.com/golang/protobuf v1.5.3 // indirect
)

replace github.com/old/package => github.com/new/package v1.0.0

exclude github.com/bad/package v1.2.3
```

**go.sum** - Checksums for dependency verification:
```
github.com/gin-gonic/gin v1.9.0 h1:abc123...
github.com/gin-gonic/gin v1.9.0/go.mod h1:def456...
```

**Commands:**
```bash
# Initialize module:
go mod init github.com/user/myproject

# Add missing dependencies:
go mod tidy

# Download dependencies:
go mod download

# Verify dependencies:
go mod verify

# Show dependency graph:
go mod graph

# Update dependency:
go get -u github.com/gin-gonic/gin

# Update all:
go get -u ./...

# Vendor dependencies:
go mod vendor
```

**Version selection:**
```bash
# Specific version:
go get github.com/gin-gonic/gin@v1.9.0

# Latest:
go get github.com/gin-gonic/gin@latest

# Specific commit:
go get github.com/gin-gonic/gin@abc1234

# Specific branch:
go get github.com/gin-gonic/gin@master
```

---

### Question 197: How does semantic versioning work in Go modules?

**Answer:**
Go uses semantic versioning (vX.Y.Z):

**Version format:**
- v1.2.3
  - Major (1): Breaking changes
  - Minor (2): New features (backward compatible)
  - Patch (3): Bug fixes

**Version ranges:**
```bash
# Latest v1:
go get github.com/pkg/errors@v1

# At least v1.2:
go get github.com/pkg/errors@>=v1.2.0

# Before v2:
go get github.com/pkg/errors@<v2.0.0
```

**Major version in import path:**
```go
// v0 and v1 - no version in path:
import "github.com/user/package"

// v2+ - version in path:
import "github.com/user/package/v2"
import "github.com/user/package/v3"
```

**Module path for v2+:**
```go
// go.mod
module github.com/user/package/v2

// Import in code:
import "github.com/user/package/v2"
```

**Minimal version selection:**
Go uses the minimum required version that satisfies all constraints:
```
A requires B v1.2
C requires B v1.3

Go selects: B v1.3 (minimum that satisfies both)
```

**Pre-release versions:**
```bash
v1.2.3-alpha.1
v1.2.3-beta.2
v1.2.3-rc.1
```

**Pseudo-versions (commits without tags):**
```
v0.0.0-20230101120000-abc1234567ab
```

---

### Question 198: How to build and deploy a Go binary to production?

**Answer:**

**Build for production:**
```bash
# Standard build:
go build -o myapp

# Optimized build (smaller binary):
go build -ldflags="-s -w" -o myapp
# -s: strip symbol table
# -w: strip DWARF debug info

# With version info:
VERSION=$(git describe --tags)
go build -ldflags="-X main.Version=$VERSION" -o myapp

# Cross-compile:
GOOS=linux GOARCH=amd64 go build -o myapp-linux
GOOS=windows GOARCH=amd64 go build -o myapp.exe
GOOS=darwin GOARCH=arm64 go build -o myapp-mac
```

**main.go with version:**
```go
package main

var (
    Version   = "dev"
    BuildTime = "unknown"
    GitCommit = "unknown"
)

func main() {
    fmt.Printf("Version: %s\n", Version)
    fmt.Printf("Built: %s\n", BuildTime)
    fmt.Printf("Commit: %s\n", GitCommit)
    
    // Your app code...
}
```

**Build script:**
```bash
#!/bin/bash
VERSION=$(git describe --tags --always)
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(git rev-parse HEAD)

go build -ldflags="\
    -X main.Version=$VERSION \
    -X main.BuildTime=$BUILD_TIME \
    -X main.GitCommit=$GIT_COMMIT \
    -s -w" \
    -o bin/myapp
```

**Makefile:**
```makefile
VERSION := $(shell git describe --tags --always)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -s -w"

.PHONY: build
build:
    go build $(LDFLAGS) -o bin/myapp

.PHONY: build-linux
build-linux:
    GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-linux

.PHONY: build-all
build-all:
    GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-linux
    GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-mac
    GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp.exe
```

**Docker deployment:**
```dockerfile
# Multi-stage build
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app

FROM scratch
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
```

---

### Question 199: What tools are used for Dockerizing Go apps?

**Answer:**

**Multi-stage Dockerfile:**
```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/main .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
```

**Using scratch (minimal image):**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app .

FROM scratch
COPY --from=builder /build/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT ["/app"]
```

**Docker Compose:**
```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://localhost/mydb
    depends_on:
      - db
  
  db:
    image: postgres:15-alpine
    environment:
      POSTGRES_PASSWORD: secret
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:
```

**.dockerignore:**
```
.git
.gitignore
README.md
Dockerfile
docker-compose.yml
*.md
.env
tmp/
```

**Build and run:**
```bash
# Build:
docker build -t myapp:latest .

# Run:
docker run -p 8080:8080 my app:latest

# With env vars:
docker run -p 8080:8080 -e DATABASE_URL=... myapp:latest

# Docker Compose:
docker-compose up --build
```

---

### Question 200: How do you set up a CI/CD pipeline for a Go project?

**Answer:**

**GitHub Actions (.github/workflows/ci.yml):**
```yaml
name: CI/CD

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Install dependencies
        run: go mod download
      
      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...
      
      - name: Run go vet
        run: go vet ./...
      
      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: latest
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          files: ./coverage.out
  
  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Build
        run: go build -v -o bin/myapp .
      
      - name: Upload artifact
        uses: actions/upload-artifact@v3
        with:
          name: myapp
          path: bin/myapp
  
  deploy:
    needs: build
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - name: Download artifact
        uses: actions/download-artifact@v3
        with:
          name: myapp
      
      - name: Deploy to server
        run: |
          # Your deployment script
          echo "Deploying..."
```

**GitLab CI (.gitlab-ci.yml):**
```yaml
stages:
  - test
  - build
  - deploy

test:
  stage: test
  image: golang:1.21
  script:
    - go mod download
    - go test -v -race -coverprofile=coverage.out ./...
    - go vet ./...
  coverage: '/coverage: \d+.\d+% of statements/'

build:
  stage: build
  image: golang:1.21
  script:
    - go build -o bin/myapp .
  artifacts:
    paths:
      - bin/myapp

deploy:
  stage: deploy
  only:
    - main
  script:
    - echo "Deploying to production..."
```

**Makefile for CI:**
```makefile
.PHONY: ci
ci: lint test build

.PHONY: lint
lint:
    golangci-lint run ./...

.PHONY: test
test:
    go test -v -race -coverprofile=coverage.out ./...

.PHONY: build
build:
    go build -o bin/myapp .

.PHONY: coverage
coverage:
    go tool cover -html=coverage.out
```

---

*[Questions 201-280 will be added in the next batch]*
