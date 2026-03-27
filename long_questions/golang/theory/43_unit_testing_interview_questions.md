# 🟢 Golang Unit Testing Interview Questions

## 1. What is the difference between unit testing and integration testing in Go?

**Answer:**
Unit tests test individual functions/components in isolation with mocked dependencies. They're fast, deterministic, and don't touch external systems. Integration tests verify multiple components work together, often using real databases or external services. In Go, unit tests typically use interfaces and mocks, while integration tests might use testcontainers or real services.

**Code Example:**
```go
// Unit Test - Fast, isolated
func TestUserService_CreateUser_Unit(t *testing.T) {
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    
    mockRepo.On("Save", mock.Anything).Return(nil)
    
    user, err := service.CreateUser("john@example.com")
    assert.NoError(t, err)
    assert.Equal(t, "john@example.com", user.Email)
}

// Integration Test - Slower, real DB
//go:build integration
func TestUserService_CreateUser_Integration(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    repo := NewSQLRepository(db)
    service := NewUserService(repo)
    
    user, err := service.CreateUser("john@example.com")
    assert.NoError(t, err)
    
    saved, err := repo.FindByID(user.ID)
    assert.NoError(t, err)
    assert.Equal(t, user.Email, saved.Email)
}
```

---

## 2. How do you write effective table-driven tests in Go?

**Answer:**
Table-driven tests use a slice of structs to define test cases with inputs and expected outputs. This pattern eliminates code duplication and makes it easy to add new test cases. Each test case should have a descriptive name and test both positive and negative scenarios.

**Code Example:**
```go
func TestCalculateDiscount(t *testing.T) {
    tests := []struct {
        name           string
        customerType   string
        purchaseAmount float64
        expectedRate   float64
        expectError    bool
    }{
        {
            name:           "premium customer high value",
            customerType:   "premium",
            purchaseAmount: 1000.0,
            expectedRate:   0.15,
            expectError:    false,
        },
        {
            name:           "regular customer low value",
            customerType:   "regular",
            purchaseAmount: 50.0,
            expectedRate:   0.05,
            expectError:    false,
        },
        {
            name:           "invalid customer type",
            customerType:   "unknown",
            purchaseAmount: 100.0,
            expectedRate:   0,
            expectError:    true,
        },
        {
            name:           "negative amount",
            customerType:   "regular",
            purchaseAmount: -10.0,
            expectedRate:   0,
            expectError:    true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            calculator := NewDiscountCalculator()
            rate, err := calculator.Calculate(tt.customerType, tt.purchaseAmount)

            if tt.expectError {
                assert.Error(t, err)
                return
            }

            assert.NoError(t, err)
            assert.Equal(t, tt.expectedRate, rate)
        })
    }
}
```

---

## 3. What are the best practices for mocking in Go unit tests?

**Answer:**
Go encourages mocking through interfaces rather than frameworks. Best practices include: 1) Define small, focused interfaces, 2) Create manual mocks for simple cases, 3) Use testify/mock for complex expectations, 4) Keep mocks in test files, 5) Don't mock what you don't own.

**Code Example:**
```go
// Interface to mock
type EmailSender interface {
    SendEmail(to, subject, body string) error
}

// Manual mock (simple case)
type MockEmailSender struct {
    sentEmails []Email
    sendError  error
}

func (m *MockEmailSender) SendEmail(to, subject, body string) error {
    if m.sendError != nil {
        return m.sendError
    }
    m.sentEmails = append(m.sentEmails, Email{To: to, Subject: subject, Body: body})
    return nil
}

// Test using manual mock
func TestNotificationService_SendWelcomeEmail(t *testing.T) {
    mockSender := &MockEmailSender{}
    service := NewNotificationService(mockSender)
    
    err := service.SendWelcomeEmail("user@example.com", "John")
    
    assert.NoError(t, err)
    assert.Len(t, mockSender.sentEmails, 1)
    assert.Equal(t, "user@example.com", mockSender.sentEmails[0].To)
    assert.Contains(t, mockSender.sentEmails[0].Subject, "Welcome")
}

// Using testify/mock for complex cases
type MockDatabase struct {
    mock.Mock
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}

func TestUserService_GetUser(t *testing.T) {
    mockDB := new(MockDatabase)
    service := NewUserService(mockDB)
    
    expectedUser := &User{ID: 1, Name: "John"}
    mockDB.On("GetUser", 1).Return(expectedUser, nil)
    
    user, err := service.GetUser(1)
    
    assert.NoError(t, err)
    assert.Equal(t, expectedUser, user)
    mockDB.AssertExpectations(t)
}
```

---

## 4. How do you test error handling in Go unit tests?

**Answer:**
Test both error conditions and success paths. Use testify/assert for assertions, test for specific error types using errors.Is/As, and ensure error messages contain useful information. Test error wrapping and unwrapping behavior.

**Code Example:**
```go
func TestValidateUser_Input(t *testing.T) {
    tests := []struct {
        name    string
        user    User
        wantErr bool
        errType error
    }{
        {
            name:    "valid user",
            user:    User{Name: "John", Email: "john@example.com"},
            wantErr: false,
        },
        {
            name:    "empty name",
            user:    User{Name: "", Email: "john@example.com"},
            wantErr: true,
            errType: ErrInvalidName,
        },
        {
            name:    "invalid email",
            user:    User{Name: "John", Email: "invalid"},
            wantErr: true,
            errType: ErrInvalidEmail,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateUser(tt.user)

            if tt.wantErr {
                assert.Error(t, err)
                assert.True(t, errors.Is(err, tt.errType), 
                    "expected error %v, got %v", tt.errType, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}

// Test error wrapping
func TestUserService_CreateUser_WrappedError(t *testing.T) {
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    
    dbErr := errors.New("database connection failed")
    mockRepo.On("Save", mock.Anything).Return(
        fmt.Errorf("failed to save user: %w", dbErr))
    
    _, err := service.CreateUser("test@example.com")
    
    assert.Error(t, err)
    assert.True(t, errors.Is(err, dbErr))
    assert.Contains(t, err.Error(), "failed to save user")
}
```

---

## 5. How do you test concurrent code in Go unit tests?

**Answer:**
Use the race detector (`go test -race`), synchronize test execution with WaitGroups or channels, test for data races, and use stress testing with many goroutines. Ensure tests are deterministic and don't rely on timing.

**Code Example:**
```go
func TestCounter_ConcurrentIncrement(t *testing.T) {
    counter := NewThreadSafeCounter()
    
    const numGoroutines = 100
    const incrementsPerGoroutine = 1000
    
    var wg sync.WaitGroup
    wg.Add(numGoroutines)
    
    for i := 0; i < numGoroutines; i++ {
        go func() {
            defer wg.Done()
            for j := 0; j < incrementsPerGoroutine; j++ {
                counter.Increment()
            }
        }()
    }
    
    wg.Wait()
    
    expected := numGoroutines * incrementsPerGoroutine
    assert.Equal(t, expected, counter.Value(), 
        "counter should be %d after concurrent increments", expected)
}

func TestCache_ConcurrentAccess(t *testing.T) {
    cache := NewConcurrentCache()
    
    // Test concurrent reads and writes
    var wg sync.WaitGroup
    wg.Add(20)
    
    // 10 writers
    for i := 0; i < 10; i++ {
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                key := fmt.Sprintf("key-%d-%d", id, j)
                value := fmt.Sprintf("value-%d-%d", id, j)
                cache.Set(key, value)
            }
        }(i)
    }
    
    // 10 readers
    for i := 0; i < 10; i++ {
        go func() {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                cache.Get(fmt.Sprintf("key-%d", j%10))
            }
        }()
    }
    
    wg.Wait()
    
    // Verify cache integrity
    assert.True(t, cache.Size() > 0)
}
```

---

## 6. How do you use test helpers and setup/teardown in Go?

**Answer:**
Use helper functions with `t.Helper()` to improve test readability, use `TestMain(m *testing.M)` for global setup/teardown, and use `defer` for cleanup within tests. Organize common test utilities in separate files.

**Code Example:**
```go
// Helper function
func setupTestDB(t *testing.T) *sql.DB {
    t.Helper()
    
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("failed to open test database: %v", err)
    }
    
    // Create tables
    _, err = db.Exec(`
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL
        )
    `)
    if err != nil {
        t.Fatalf("failed to create tables: %v", err)
    }
    
    return db
}

func cleanupTestDB(t *testing.T, db *sql.DB) {
    t.Helper()
    if err := db.Close(); err != nil {
        t.Logf("warning: failed to close test database: %v", err)
    }
}

// Test using helpers
func TestUserRepository_Create(t *testing.T) {
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    repo := NewUserRepository(db)
    user := &User{Name: "John", Email: "john@example.com"}
    
    err := repo.Create(user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
}

// Global setup/teardown
func TestMain(m *testing.M) {
    // Global setup
    testDB := setupGlobalTestDB()
    defer cleanupGlobalTestDB(testDB)
    
    // Run tests
    code := m.Run()
    
    os.Exit(code)
}
```

---

## 7. How do you achieve good test coverage without targeting 100%?

**Answer:**
Focus on testing critical business logic, error paths, and edge cases. Use coverage reports to identify untested code paths. Aim for 80-90% coverage but prioritize quality over quantity. Use integration tests for external dependencies and unit tests for business logic.

**Code Example:**
```go
// Focus on critical business logic
func TestLoanCalculator_CalculateMonthlyPayment_EdgeCases(t *testing.T) {
    calculator := NewLoanCalculator()
    
    tests := []struct {
        name        string
        principal   float64
        rate        float64
        months      int
        expected    float64
        expectError bool
    }{
        {
            name:        "zero principal",
            principal:   0,
            rate:        0.05,
            months:      12,
            expected:    0,
            expectError: false,
        },
        {
            name:        "negative principal",
            principal:   -1000,
            rate:        0.05,
            months:      12,
            expected:    0,
            expectError: true,
        },
        {
            name:        "zero interest rate",
            principal:   10000,
            rate:        0,
            months:      12,
            expected:    833.33,
            expectError: false,
        },
        {
            name:        "very high interest rate",
            principal:   1000,
            rate:        0.5, // 50% annual rate
            months:      12,
            expected:    95.62,
            expectError: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            payment, err := calculator.CalculateMonthlyPayment(tt.principal, tt.rate, tt.months)
            
            if tt.expectError {
                assert.Error(t, err)
                return
            }
            
            assert.NoError(t, err)
            assert.InDelta(t, tt.expected, payment, 0.01)
        })
    }
}
```

---

## 8. How do you test HTTP handlers and middleware in Go?

**Answer:**
Use `net/http/httptest` for testing HTTP handlers without starting a real server. Test request parsing, response writing, status codes, headers, and middleware behavior. Use table-driven tests for different request scenarios.

**Code Example:**
```go
func TestUserHandler_GetUser(t *testing.T) {
    tests := []struct {
        name           string
        userID         string
        mockUser       *User
        mockError      error
        expectedStatus int
        expectedBody   string
    }{
        {
            name:           "valid user",
            userID:         "1",
            mockUser:       &User{ID: 1, Name: "John", Email: "john@example.com"},
            mockError:      nil,
            expectedStatus: http.StatusOK,
            expectedBody:   `{"id":1,"name":"John","email":"john@example.com"}`,
        },
        {
            name:           "user not found",
            userID:         "999",
            mockUser:       nil,
            mockError:      ErrUserNotFound,
            expectedStatus: http.StatusNotFound,
            expectedBody:   `{"error":"user not found"}`,
        },
        {
            name:           "invalid user ID",
            userID:         "invalid",
            mockUser:       nil,
            mockError:      nil,
            expectedStatus: http.StatusBadRequest,
            expectedBody:   `{"error":"invalid user ID"}`,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockService := &MockUserService{}
            handler := NewUserHandler(mockService)
            
            if tt.userID != "invalid" && tt.mockError == nil {
                mockService.On("GetUser", 1).Return(tt.mockUser, tt.mockError)
            } else if tt.mockError != nil {
                mockService.On("GetUser", mock.AnythingOfType("int")).Return(nil, tt.mockError)
            }
            
            req := httptest.NewRequest("GET", "/users/"+tt.userID, nil)
            req = mux.SetURLVars(req, map[string]string{"id": tt.userID})
            
            w := httptest.NewRecorder()
            handler.GetUser(w, req)
            
            resp := w.Result()
            defer resp.Body.Close()
            
            assert.Equal(t, tt.expectedStatus, resp.StatusCode)
            
            body, _ := io.ReadAll(resp.Body)
            assert.JSONEq(t, tt.expectedBody, string(body))
            
            mockService.AssertExpectations(t)
        })
    }
}

func TestAuthMiddleware_RequireAuth(t *testing.T) {
    middleware := AuthMiddleware{}
    
    tests := []struct {
        name           string
        authHeader     string
        expectedStatus int
    }{
        {
            name:           "valid token",
            authHeader:     "Bearer valid-token",
            expectedStatus: http.StatusOK,
        },
        {
            name:           "missing token",
            authHeader:     "",
            expectedStatus: http.StatusUnauthorized,
        },
        {
            name:           "invalid token format",
            authHeader:     "InvalidFormat token",
            expectedStatus: http.StatusUnauthorized,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest("GET", "/protected", nil)
            if tt.authHeader != "" {
                req.Header.Set("Authorization", tt.authHeader)
            }
            
            w := httptest.NewRecorder()
            
            // Mock handler that sets status to OK if middleware allows
            nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(http.StatusOK)
            })
            
            middleware.RequireAuth(nextHandler).ServeHTTP(w, req)
            
            assert.Equal(t, tt.expectedStatus, w.Code)
        })
    }
}
```

---

## 9. How do you test time-dependent functionality in Go?

**Answer:**
Inject a time interface, use a mock clock for deterministic tests, test time-based logic without actual delays, and verify timeout behavior. This makes tests fast and reliable.

**Code Example:**
```go
// Time interface for dependency injection
type Clock interface {
    Now() time.Time
    Sleep(duration time.Duration)
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }
func (RealClock) Sleep(d time.Duration) { time.Sleep(d) }

// Mock clock for testing
type MockClock struct {
    currentTime time.Time
}

func (m *MockClock) Now() time.Time { return m.currentTime }
func (m *MockClock) Sleep(duration time.Duration) { 
    m.currentTime = m.currentTime.Add(duration)
}

func (m *MockClock) SetTime(t time.Time) { m.currentTime = t }
func (m *MockClock) Advance(duration time.Duration) { 
    m.currentTime = m.currentTime.Add(duration)
}

// Service that uses time
type TokenService struct {
    clock  Clock
    tokens map[string]time.Time // token -> expiry
}

func (s *TokenService) IsExpired(token string) bool {
    expiry, exists := s.tokens[token]
    if !exists {
        return true
    }
    return s.clock.Now().After(expiry)
}

// Test with mock clock
func TestTokenService_IsExpired(t *testing.T) {
    mockClock := &MockClock{currentTime: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)}
    service := &TokenService{clock: mockClock, tokens: make(map[string]time.Time)}
    
    // Set token to expire in 1 hour
    expiry := mockClock.Now().Add(time.Hour)
    service.tokens["valid-token"] = expiry
    
    // Token should not be expired
    assert.False(t, service.IsExpired("valid-token"))
    
    // Advance time by 2 hours
    mockClock.Advance(2 * time.Hour)
    
    // Token should now be expired
    assert.True(t, service.IsExpired("valid-token"))
    
    // Non-existent token should be expired
    assert.True(t, service.IsExpired("non-existent"))
}
```

---

## 10. How do you organize and structure Go test files effectively?

**Answer:**
Follow Go conventions: `xxx_test.go` files, separate unit and integration tests with build tags, organize test utilities in `testutils` package, use descriptive test names, and group related tests together.

**Code Example:**
```go
// File: user_service_test.go
package users

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

// Unit tests - same package
func TestUserService_CreateUser_Unit(t *testing.T) {
    // ... unit test implementation
}

// Test utilities
func createTestUser(t *testing.T) *User {
    t.Helper()
    return &User{
        Name:  "Test User",
        Email: "test@example.com",
    }
}

// File: user_service_integration_test.go
//go:build integration
// +build integration

package users

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestUserService_CreateUser_Integration(t *testing.T) {
    // ... integration test implementation
}

// File: testutils/users_test.go
package testutils

import (
    "testing"
    "time"
    "yourproject/users"
)

func NewTestUserService(t *testing.T) *users.UserService {
    t.Helper()
    
    mockRepo := &MockUserRepository{}
    return users.NewUserService(mockRepo)
}

func AssertUserEqual(t *testing.T, expected, actual *users.User) {
    t.Helper()
    
    assert.Equal(t, expected.ID, actual.ID)
    assert.Equal(t, expected.Name, actual.Name)
    assert.Equal(t, expected.Email, actual.Email)
    assert.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
}
```

---

## 11. How do you use testify framework effectively in Go unit tests?

**Answer:**
Use testify/assert for assertions, testify/require for fatal assertions, testify/suite for complex test setup, and testify/mock for mocking. Leverage assertion helpers for better error messages and cleaner test code.

**Code Example:**
```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "github.com/stretchr/testify/suite"
)

// Using assert for non-fatal assertions
func TestCalculator_Add(t *testing.T) {
    calculator := NewCalculator()
    
    result := calculator.Add(2, 3)
    assert.Equal(t, 5, result, "2 + 3 should equal 5")
    assert.NotZero(t, result, "result should not be zero")
    
    // Test continues even if assertion fails
    result2 := calculator.Add(-1, 1)
    assert.Equal(t, 0, result2, "-1 + 1 should equal 0")
}

// Using require for fatal assertions (stops test on failure)
func TestUserService_CreateUser_Required(t *testing.T) {
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    
    // Setup - if this fails, test should stop
    mockRepo.On("Save", mock.Anything).Return(nil)
    
    user, err := service.CreateUser("test@example.com")
    require.NoError(t, err, "user creation should not fail")
    require.NotNil(t, user, "user should not be nil")
    
    // Continue testing with valid user
    assert.Equal(t, "test@example.com", user.Email)
}

// Using testify/suite for complex setup
type UserServiceTestSuite struct {
    suite.Suite
    service *UserService
    mockRepo *MockUserRepository
}

func (suite *UserServiceTestSuite) SetupTest() {
    suite.mockRepo = &MockUserRepository{}
    suite.service = NewUserService(suite.mockRepo)
}

func (suite *UserServiceTestSuite) TearDownTest() {
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserServiceTestSuite) TestCreateUser_Success() {
    suite.mockRepo.On("Save", mock.Anything).Return(nil)
    
    user, err := suite.service.CreateUser("test@example.com")
    
    suite.NoError(err)
    suite.NotNil(user)
    suite.Equal("test@example.com", user.Email)
}

func (suite *UserServiceTestSuite) TestCreateUser_DuplicateEmail() {
    suite.mockRepo.On("Save", mock.Anything).Return(ErrDuplicateEmail)
    
    user, err := suite.service.CreateUser("duplicate@example.com")
    
    suite.Error(err)
    suite.Nil(user)
    suite.True(errors.Is(err, ErrDuplicateEmail))
}

func TestUserServiceTestSuite(t *testing.T) {
    suite.Run(t, new(UserServiceTestSuite))
}
```

---

## 12. How do you test file system operations in Go?

**Answer:**
Use `testing/fstest` for in-memory file systems, create temporary files/dirs with `os.CreateTemp`, use `testify/fs` for file assertions, and clean up resources with `defer`. Avoid testing against the real file system when possible.

**Code Example:**
```go
import (
    "os"
    "testing"
    "testing/fstest"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

// Test with in-memory file system
func TestConfigLoader_Load_MemFS(t *testing.T) {
    // Create in-memory file system
    memFS := fstest.MapFS{
        "config.yaml": &fstest.MapFile{
            Data: []byte(`
database:
  host: localhost
  port: 5432
  name: testdb
server:
  port: 8080
`),
        },
        "empty.yaml": &fstest.MapFile{Data: []byte{}},
    }
    
    loader := NewConfigLoader()
    
    // Test valid config
    config, err := loader.LoadFromFS(memFS, "config.yaml")
    require.NoError(t, err)
    assert.Equal(t, "localhost", config.Database.Host)
    assert.Equal(t, 5432, config.Database.Port)
    assert.Equal(t, 8080, config.Server.Port)
    
    // Test non-existent file
    _, err = loader.LoadFromFS(memFS, "nonexistent.yaml")
    assert.Error(t, err)
    
    // Test empty file
    _, err = loader.LoadFromFS(memFS, "empty.yaml")
    assert.Error(t, err)
}

// Test with temporary files
func TestFileProcessor_Process_TempFile(t *testing.T) {
    // Create temporary file
    tmpFile, err := os.CreateTemp("", "test_*.txt")
    require.NoError(t, err)
    defer os.Remove(tmpFile.Name()) // Cleanup
    defer tmpFile.Close()
    
    // Write test data
    testData := "line1\nline2\nline3\n"
    _, err = tmpFile.WriteString(testData)
    require.NoError(t, err)
    
    // Reset file pointer
    _, err = tmpFile.Seek(0, 0)
    require.NoError(t, err)
    
    // Test file processing
    processor := NewFileProcessor()
    lines, err := processor.Process(tmpFile.Name())
    
    require.NoError(t, err)
    assert.Equal(t, []string{"line1", "line2", "line3"}, lines)
}

// Test directory operations
func TestDirectoryWatcher_Watch_TempDir(t *testing.T) {
    // Create temporary directory
    tmpDir, err := os.MkdirTemp("", "watch_test_*")
    require.NoError(t, err)
    defer os.RemoveAll(tmpDir) // Cleanup
    
    watcher := NewDirectoryWatcher()
    events := make(chan FileEvent, 10)
    
    err = watcher.Watch(tmpDir, events)
    require.NoError(t, err)
    defer watcher.Stop()
    
    // Create test file
    testFile := filepath.Join(tmpDir, "test.txt")
    err = os.WriteFile(testFile, []byte("test content"), 0644)
    require.NoError(t, err)
    
    // Wait for event (with timeout)
    select {
    case event := <-events:
        assert.Equal(t, EventTypeCreated, event.Type)
        assert.Equal(t, "test.txt", filepath.Base(event.Path))
    case <-time.After(time.Second):
        t.Fatal("expected file event but got none")
    }
}
```

---

## 13. How do you test JSON marshaling/unmarshaling in Go?

**Answer:**
Test both valid and invalid JSON, test edge cases (null values, empty arrays), use table-driven tests for different scenarios, and verify error messages. Test custom JSON tags and omitempty behavior.

**Code Example:**
```go
func TestUser_MarshalJSON(t *testing.T) {
    tests := []struct {
        name     string
        user     User
        expected string
        wantErr  bool
    }{
        {
            name: "complete user",
            user: User{
                ID:    1,
                Name:  "John Doe",
                Email: "john@example.com",
                Role:  "admin",
            },
            expected: `{"id":1,"name":"John Doe","email":"john@example.com","role":"admin"}`,
            wantErr:  false,
        },
        {
            name: "user with empty optional fields",
            user: User{
                ID:   2,
                Name: "Jane",
            },
            expected: `{"id":2,"name":"Jane","email":"","role":""}`,
            wantErr:  false,
        },
        {
            name: "user with zero values",
            user: User{},
            expected: `{"id":0,"name":"","email":"","role":""}`,
            wantErr:  false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            data, err := json.Marshal(tt.user)
            
            if tt.wantErr {
                assert.Error(t, err)
                return
            }
            
            assert.NoError(t, err)
            assert.JSONEq(t, tt.expected, string(data))
        })
    }
}

func TestUser_UnmarshalJSON(t *testing.T) {
    tests := []struct {
        name      string
        data      string
        expected  User
        wantErr   bool
        errType   error
    }{
        {
            name: "valid JSON",
            data: `{"id":1,"name":"John","email":"john@example.com","role":"admin"}`,
            expected: User{
                ID:    1,
                Name:  "John",
                Email: "john@example.com",
                Role:  "admin",
            },
            wantErr: false,
        },
        {
            name: "JSON with missing fields",
            data: `{"id":1,"name":"John"}`,
            expected: User{
                ID:   1,
                Name: "John",
            },
            wantErr: false,
        },
        {
            name: "JSON with extra fields",
            data: `{"id":1,"name":"John","email":"john@example.com","role":"admin","extra":"value"}`,
            expected: User{
                ID:    1,
                Name:  "John",
                Email: "john@example.com",
                Role:  "admin",
            },
            wantErr: false,
        },
        {
            name:    "invalid JSON",
            data:    `{"id":1,"name":"John",}`,
            wantErr: true,
        },
        {
            name:    "invalid type for ID",
            data:    `{"id":"not_a_number","name":"John"}`,
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var user User
            err := json.Unmarshal([]byte(tt.data), &user)
            
            if tt.wantErr {
                assert.Error(t, err)
                if tt.errType != nil {
                    assert.True(t, errors.Is(err, tt.errType))
                }
                return
            }
            
            assert.NoError(t, err)
            assert.Equal(t, tt.expected, user)
        })
    }
}

func TestUser_ValidationInUnmarshal(t *testing.T) {
    // Test custom validation during unmarshaling
    data := `{"id":1,"name":"","email":"invalid-email","role":"admin"}`
    
    var user User
    err := json.Unmarshal([]byte(data), &user)
    
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "validation failed")
}
```

---

## 14. How do you benchmark Go code effectively?

**Answer:**
Use `testing.B` for benchmarks, run benchmarks with `-bench` flag, use `-benchmem` for memory allocation stats, reset timers before measured code, and avoid including setup time in measurements. Use benchstat for comparing results.

**Code Example:**
```go
func BenchmarkStringConcatenation(b *testing.B) {
    tests := []struct {
        name  string
        size  int
    }{
        {"small", 10},
        {"medium", 100},
        {"large", 1000},
    }
    
    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            // Setup
            strings := make([]string, tt.size)
            for i := 0; i < tt.size; i++ {
                strings[i] = fmt.Sprintf("string-%d", i)
            }
            
            // Reset timer to exclude setup time
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                var result string
                for _, s := range strings {
                    result += s
                }
            }
        })
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    tests := []struct {
        name  string
        size  int
    }{
        {"small", 10},
        {"medium", 100},
        {"large", 1000},
    }
    
    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            // Setup
            strings := make([]string, tt.size)
            for i := 0; i < tt.size; i++ {
                strings[i] = fmt.Sprintf("string-%d", i)
            }
            
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                var builder strings.Builder
                for _, s := range strings {
                    builder.WriteString(s)
                }
                _ = builder.String()
            }
        })
    }
}

func BenchmarkUserService_CreateUser(b *testing.B) {
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    
    mockRepo.On("Save", mock.Anything).Return(nil)
    defer mockRepo.AssertExpectations(&testing.T{})
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        _, err := service.CreateUser(fmt.Sprintf("user%d@example.com", i))
        if err != nil {
            b.Fatal(err)
        }
    }
}

// Parallel benchmark
func BenchmarkCache_ConcurrentReads(b *testing.B) {
    cache := NewConcurrentCache()
    
    // Pre-populate cache
    for i := 0; i < 1000; i++ {
        cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
    }
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            key := fmt.Sprintf("key%d", rand.Intn(1000))
            _ = cache.Get(key)
        }
    })
}
```

---

## 15. How do you handle test data management and factories?

**Answer:**
Use factory pattern for creating test data, use builders for complex objects, separate test data from test logic, and use faker libraries for realistic data. Ensure test data is deterministic and reproducible.

**Code Example:**
```go
// User factory
type UserFactory struct {
    idCounter int
}

func NewUserFactory() *UserFactory {
    return &UserFactory{idCounter: 1}
}

func (f *UserFactory) CreateUser() *User {
    user := &User{
        ID:        f.idCounter,
        Name:      fmt.Sprintf("User %d", f.idCounter),
        Email:     fmt.Sprintf("user%d@example.com", f.idCounter),
        CreatedAt: time.Now(),
    }
    f.idCounter++
    return user
}

func (f *UserFactory) CreateUserWithEmail(email string) *User {
    user := f.CreateUser()
    user.Email = email
    return user
}

func (f *UserFactory) CreateAdminUser() *User {
    user := f.CreateUser()
    user.Role = "admin"
    return user
}

// Builder pattern for complex objects
type UserBuilder struct {
    user *User
}

func NewUserBuilder() *UserBuilder {
    return &UserBuilder{
        user: &User{
            CreatedAt: time.Now(),
        },
    }
}

func (b *UserBuilder) WithID(id int) *UserBuilder {
    b.user.ID = id
    return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
    b.user.Name = name
    return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
    b.user.Email = email
    return b
}

func (b *UserBuilder) WithRole(role string) *UserBuilder {
    b.user.Role = role
    return b
}

func (b *UserBuilder) Build() *User {
    return b.user
}

// Using factory and builder in tests
func TestUserService_CreateUser_WithFactory(t *testing.T) {
    factory := NewUserFactory()
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    
    testUser := factory.CreateUserWithEmail("test@example.com")
    mockRepo.On("Save", testUser).Return(nil)
    
    createdUser, err := service.CreateUser("test@example.com")
    
    assert.NoError(t, err)
    assert.Equal(t, testUser.Email, createdUser.Email)
    mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateRole_WithBuilder(t *testing.T) {
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    
    existingUser := NewUserBuilder().
        WithID(1).
        WithName("John").
        WithEmail("john@example.com").
        WithRole("user").
        Build()
    
    mockRepo.On("GetByID", 1).Return(existingUser, nil)
    mockRepo.On("Update", mock.MatchedBy(func(u *User) bool {
        return u.Role == "admin"
    })).Return(nil)
    
    err := service.UpdateRole(1, "admin")
    
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

// Using faker for realistic data
func TestUserService_BulkCreate_WithFaker(t *testing.T) {
    faker := gofakeit.New(0)
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    
    // Create 100 realistic users
    var users []*User
    for i := 0; i < 100; i++ {
        user := &User{
            Name:  faker.Name(),
            Email: faker.Email(),
            Role:  faker.RandomString([]string{"user", "admin", "moderator"}),
        }
        users = append(users, user)
    }
    
    mockRepo.On("Save", mock.AnythingOfType("*users.User")).Return(nil)
    
    err := service.BulkCreate(users)
    
    assert.NoError(t, err)
    assert.Equal(t, 100, mockRepo.SaveCallCount)
}
```

---

## Summary

These unit testing interview questions cover:

1. **Fundamental Concepts**: Unit vs integration testing, table-driven tests
2. **Mocking Strategies**: Interface-based mocking, testify/mock
3. **Error Handling**: Testing error paths, wrapped errors
4. **Concurrency**: Race detection, stress testing
5. **Test Organization**: Helpers, setup/teardown, file structure
6. **HTTP Testing**: Handlers, middleware, httptest
7. **Time Testing**: Mock clocks, deterministic tests
8. **Advanced Topics**: File operations, JSON, benchmarks, data factories

Key takeaways for interviews:
- Focus on testability through interface design
- Use table-driven tests for comprehensive coverage
- Mock external dependencies, don't test them
- Test both happy paths and error conditions
- Use testify for cleaner, more readable tests
- Consider performance with benchmarks and race detection
