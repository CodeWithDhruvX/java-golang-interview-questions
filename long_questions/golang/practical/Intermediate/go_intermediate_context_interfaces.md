# Go Intermediate — Context, Interfaces, Advanced Errors & Generics

> **30% Gap Coverage — Part 2 of 2**
> Topics: `context` package · Interface composition · Advanced error patterns · Generics (Go 1.18+) · Real-world patterns

---

## Section 1: context Package (Q1–Q20)

### 1. context.Background vs context.TODO
**Q: What is the difference?**
```go
package main
import (
    "context"
    "fmt"
)

func main() {
    ctx1 := context.Background() // top-level, never cancelled
    ctx2 := context.TODO()       // placeholder: you plan to add context later

    fmt.Println(ctx1)
    fmt.Println(ctx2)
}
```
**A:** Both are non-nil, empty contexts that are never cancelled. `Background` is used at `main`, `init`, and as the top of a request tree. `TODO` is a lint-friendly placeholder signaling "I know I need a context here but haven't decided which yet."

---

### 2. context.WithCancel — Manual Cancellation
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func work(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("work stopped:", ctx.Err())
            return
        default:
            time.Sleep(20 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go work(ctx)
    time.Sleep(100 * time.Millisecond)
    cancel() // signal done
    time.Sleep(50 * time.Millisecond)
}
```
**A:** `work stopped: context canceled`. `cancel()` closes `ctx.Done()` channel. The goroutine detects it via `select` and exits. **Always call `cancel()` — defer it immediately after creation.**

---

### 3. defer cancel() Is Mandatory
**Q: What is the resource leak if cancel is not called?**
```go
// WRONG — leaked context resources
func badHandler() {
    ctx, _ := context.WithCancel(context.Background()) // cancel discarded!
    doWork(ctx)
}

// CORRECT
func goodHandler() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // always call, even if context is also cancelled by parent
    doWork(ctx)
}
```
**A:** Without `cancel()`, Go cannot release internal context resources (goroutine checking for cancellation, timer). `go vet` and linters (`staticcheck`) flag discarded cancel functions.

---

### 4. context.WithTimeout
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func fetch(ctx context.Context) error {
    select {
    case <-time.After(500 * time.Millisecond): // simulate slow work
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
    defer cancel()
    err := fetch(ctx)
    fmt.Println(err)
}
```
**A:** `context deadline exceeded`. Task takes 500ms but context expires at 200ms.

---

### 5. WithTimeout vs WithDeadline
**Q: What is the difference?**
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func main() {
    // WithTimeout: relative duration from now
    ctx1, c1 := context.WithTimeout(context.Background(), 5*time.Second)
    defer c1()

    // WithDeadline: absolute point in time
    deadline := time.Now().Add(5 * time.Second)
    ctx2, c2 := context.WithDeadline(context.Background(), deadline)
    defer c2()

    fmt.Println(ctx1.Err(), ctx2.Err()) // both nil initially
}
```
**A:** `<nil> <nil>`. Functionally equivalent here. `WithTimeout(ctx, d)` is shorthand for `WithDeadline(ctx, time.Now().Add(d))`. Use `WithDeadline` when you have an absolute time (e.g., from an incoming request).

---

### 6. ctx.Deadline() — Inspect Deadline
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    deadline, ok := ctx.Deadline()
    if ok {
        fmt.Printf("deadline in: %.0fs\n", time.Until(deadline).Seconds())
    }
}
```
**A:** `deadline in: 10s`. `ctx.Deadline()` returns the deadline time and whether one is set. Useful to pass remaining time budget to downstream calls.

---

### 7. Context Propagation Through Call Chain
**Q: Why must context be the first parameter by convention?**
```go
package main
import (
    "context"
    "fmt"
)

func step3(ctx context.Context) error {
    if ctx.Err() != nil {
        return ctx.Err()
    }
    fmt.Println("step3 done")
    return nil
}

func step2(ctx context.Context) error { return step3(ctx) }
func step1(ctx context.Context) error { return step2(ctx) }

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    cancel() // cancelled before use
    fmt.Println(step1(ctx))
}
```
**A:** `context canceled` (step3 detects cancellation). Context flows **downward** through every function — if cancelled at any layer, all downstream calls can check and short-circuit.

---

### 8. context.WithValue — Storing Request-Scoped Data
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

type key string

const requestIDKey key = "requestID"

func handler(ctx context.Context) {
    id := ctx.Value(requestIDKey)
    fmt.Println("request ID:", id)
}

func main() {
    ctx := context.WithValue(context.Background(), requestIDKey, "abc-123")
    handler(ctx)
}
```
**A:** `request ID: abc-123`. Use an **unexported custom type** as key (not `string`) to prevent key collisions across packages.

---

### 9. context.Value Returns nil for Missing Key
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

type myKey struct{}

func main() {
    ctx := context.Background()
    val := ctx.Value(myKey{})
    fmt.Println(val)
    fmt.Println(val == nil)
}
```
**A:**
```
<nil>
true
```
Always check for nil before type-asserting context values.

---

### 10. context.WithValue Key Collision Protection
**Q: Does this compile and what is the output?**
```go
package main
import (
    "context"
    "fmt"
)

type pkgAKey string
type pkgBKey string

func main() {
    ctx := context.Background()
    ctx = context.WithValue(ctx, pkgAKey("id"), "from-A")
    ctx = context.WithValue(ctx, pkgBKey("id"), "from-B")

    fmt.Println(ctx.Value(pkgAKey("id")))
    fmt.Println(ctx.Value(pkgBKey("id")))
}
```
**A:**
```
from-A
from-B
```
Despite both keys having the same underlying string `"id"`, they are different Go types — no collision. This is why context keys must be typed.

---

### 11. HTTP Request Context — Standard Pattern
**Q: What is the standard pattern to propagate context in HTTP handlers?**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context() // context from the incoming request

    // With a timeout for downstream calls
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    result, err := fetchFromDB(ctx)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            http.Error(w, "timeout", http.StatusGatewayTimeout)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}
```
**A:** `r.Context()` carries cancellation from the HTTP client (e.g., if client disconnects). Always derive child contexts from `r.Context()`, never from `context.Background()` in handlers.

---

### 12. Context Cancels Downstream Goroutine
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
    "sync"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    <-ctx.Done()
    fmt.Printf("worker %d exiting: %v\n", id, ctx.Err())
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(ctx, i, &wg)
    }

    cancel()
    wg.Wait()
}
```
**A:** All 3 workers print `exiting: context canceled` (order varies). One `cancel()` call stops all listeners simultaneously.

---

### 13. Chained Context — Parent Cancels Child
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func main() {
    parent, pcancel := context.WithCancel(context.Background())
    defer pcancel()

    child, ccancel := context.WithTimeout(parent, 10*time.Second)
    defer ccancel()

    pcancel() // cancel parent early
    time.Sleep(10 * time.Millisecond)
    fmt.Println(child.Err())
}
```
**A:** `context canceled`. Cancelling a parent **propagates automatically** to all derived children, regardless of their own deadlines.

---

### 14. context.Cause (Go 1.20+)
**Q: What is the output?**
```go
package main
import (
    "context"
    "errors"
    "fmt"
)

var ErrShutdown = errors.New("shutdown requested")

func main() {
    ctx, cancel := context.WithCancelCause(context.Background())
    cancel(ErrShutdown)

    fmt.Println(ctx.Err())
    fmt.Println(context.Cause(ctx)) // returns the specific cause
}
```
**A:**
```
context canceled
shutdown requested
```
`WithCancelCause` (Go 1.20+) lets you attach a specific error to cancellation. `context.Cause` retrieves it — more informative than just `context.Canceled`.

---

### 15. Anti-Pattern: Storing Context in Struct
**Q: What is wrong with this?**
```go
// WRONG
type Server struct {
    ctx context.Context // DON'T store context in struct
}

// CORRECT
type Server struct{}

func (s *Server) HandleRequest(ctx context.Context, req Request) error {
    // ctx is passed per-call, not stored
    return processWithContext(ctx, req)
}
```
**A:** Storing context in a struct means it cannot be cancelled per-request — the same context persists across all operations. **Context must be passed as the first function parameter** per the Go guidelines.

---

### 16. ctx.Done() Returns nil for Background
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

func main() {
    ctx := context.Background()
    fmt.Println(ctx.Done() == nil) // Background never cancels
    fmt.Println(ctx.Err())
}
```
**A:**
```
true
<nil>
```
`context.Background().Done()` returns `nil`. A nil channel blocks forever in `select` — so code that selects on `ctx.Done()` is correctly always waiting with a Background context.

---

### 17. Timeout Propagation to DB Query
**Q: What is the pattern for database calls with context?**
```go
func getUserByID(ctx context.Context, db *sql.DB, id int) (*User, error) {
    // DB respects context cancellation — query is killed if ctx expires
    row := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = $1", id)
    
    var u User
    if err := row.Scan(&u.Name); err != nil {
        return nil, fmt.Errorf("getUserByID %d: %w", id, err)
    }
    return &u, nil
}
```
**A:** `QueryRowContext` (and `ExecContext`, `QueryContext`) accept a context. If the context times out, the query is cancelled at the database driver level. Always use `*Context` variants in production.

---

### 18. Context Value Anti-Pattern
**Q: What should NOT be stored in context?**
```go
// BAD: passing required function params via context
ctx = context.WithValue(ctx, "userID", 42)
ctx = context.WithValue(ctx, "db", dbConn)

// GOOD: context for optional cross-cutting concerns only
ctx = context.WithValue(ctx, requestIDKey, "req-abc")  // tracing
ctx = context.WithValue(ctx, authTokenKey, token)       // middleware-injected auth
```
**A:** Context values should only carry **request-scoped cross-cutting data** (trace IDs, auth tokens, deadlines). Never use context to smuggle required function parameters — that makes APIs implicit and hard to test.

---

### 19. Detecting Client Disconnect in HTTP
**Q: How do you detect client disconnect?**
```go
func streamHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    for {
        select {
        case <-ctx.Done():
            fmt.Println("client disconnected:", ctx.Err())
            return
        default:
            fmt.Fprintln(w, "data chunk")
            w.(http.Flusher).Flush()
            time.Sleep(time.Second)
        }
    }
}
```
**A:** When an HTTP client disconnects mid-stream, `r.Context()` is cancelled. Checking `ctx.Done()` in the streaming loop prevents wasted work and resource leaks.

---

### 20. select Priority — Context Done First
**Q: Which case is checked first in select?**
```go
func drainWithCancel(ctx context.Context, ch <-chan int) {
    for {
        select {
        case <-ctx.Done():
            return // always honour cancellation first
        case v := <-ch:
            process(v)
        }
    }
}
```
**A:** Neither — Go `select` has **no built-in priority**. However, by convention, placing `ctx.Done()` first is a style signal to readers. To enforce priority, use a nested select: check `ctx.Done()` in a non-blocking select before reading `ch`.

---

## Section 2: Interface Composition (Q21–Q32)

### 21. Composing Interfaces
**Q: What does this compile to?**
```go
package main
import "fmt"

type Reader interface { Read() string }
type Writer interface { Write(s string) }

// Composed interface
type ReadWriter interface {
    Reader
    Writer
}

type Buffer struct{ data string }
func (b *Buffer) Read() string    { return b.data }
func (b *Buffer) Write(s string)  { b.data = s }

func process(rw ReadWriter) {
    rw.Write("hello")
    fmt.Println(rw.Read())
}

func main() {
    process(&Buffer{})
}
```
**A:** `hello`. Interface composition embeds multiple interfaces. `*Buffer` satisfies all three interfaces simultaneously.

---

### 22. io.Reader and io.Writer — Foundation Interfaces
**Q: What is the output?**
```go
package main
import (
    "bytes"
    "fmt"
    "strings"
    "io"
)

func copyData(dst io.Writer, src io.Reader) (int64, error) {
    return io.Copy(dst, src)
}

func main() {
    src := strings.NewReader("Hello, Go!")
    dst := &bytes.Buffer{}
    n, _ := copyData(dst, src)
    fmt.Printf("copied %d bytes: %s\n", n, dst.String())
}
```
**A:** `copied 10 bytes: Hello, Go!`. The power of Go interfaces: `strings.Reader` and `bytes.Buffer` satisfy `io.Reader`/`io.Writer` without any explicit declaration.

---

### 23. io.ReadCloser Composition
**Q: What is the pattern?**
```go
import "io"

// io.ReadCloser is defined as:
type ReadCloser interface {
    Reader
    Closer
}

// http.Response.Body is an io.ReadCloser
func readBody(body io.ReadCloser) string {
    defer body.Close() // Closer
    data, _ := io.ReadAll(body) // Reader
    return string(data)
}
```
**A:** Embedding smaller interfaces builds focused, composable APIs. `http.Response.Body` must be both read AND closed — `ReadCloser` enforces both.

---

### 24. Interface Segregation — Accept Minimal Interface
**Q: Which function signature is better?**
```go
// BAD: accepts more than needed
func logContent(f *os.File) {
    data, _ := io.ReadAll(f)
    fmt.Println(string(data))
}

// GOOD: accepts minimal interface
func logContent(r io.Reader) {
    data, _ := io.ReadAll(r)
    fmt.Println(string(data))
}
```
**A:** The `io.Reader` version is better. It works with files, HTTP bodies, strings, byte buffers, network connections — anything that can be read. This is the **Go interface principle**: accept interfaces, return concrete types.

---

### 25. Satisfying Multiple Interfaces
**Q: What is the output?**
```go
package main
import "fmt"

type Animal interface{ Sound() string }
type Namer  interface{ Name() string }

type Dog struct{ name string }
func (d Dog) Sound() string { return "woof" }
func (d Dog) Name()  string { return d.name }

func makeSound(a Animal) { fmt.Println(a.Sound()) }
func getName(n Namer)    { fmt.Println(n.Name()) }

func main() {
    d := Dog{name: "Rex"}
    makeSound(d)
    getName(d)
}
```
**A:**
```
woof
Rex
```
`Dog` satisfies both `Animal` and `Namer` without any explicit declaration.

---

### 26. Interface Upgrade via Type Assertion
**Q: What is the output?**
```go
package main
import "fmt"

type Flusher interface{ Flush() }

type BasicWriter struct{}
func (b *BasicWriter) Write(s string) { fmt.Print(s) }

type FlushWriter struct{ BasicWriter }
func (f *FlushWriter) Flush() { fmt.Println("[flushed]") }

func maybeFlush(w interface{ Write(string) }) {
    w.Write("data ")
    if f, ok := w.(Flusher); ok {
        f.Flush()
    }
}

func main() {
    maybeFlush(&BasicWriter{})
    maybeFlush(&FlushWriter{})
}
```
**A:**
```
data 
data [flushed]
```
Type asserting to a richer interface is the **optional interface pattern** — used extensively in `net/http` (e.g., `http.Flusher`, `http.Hijacker`).

---

### 27. Implementing sort.Interface
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

type ByLength []string

func (s ByLength) Len() int           { return len(s) }
func (s ByLength) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func (s ByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
    fruits := []string{"peach", "kiwi", "apple", "fig"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
```
**A:** `[fig kiwi peach apple]`. Sorted by string length. Implementing `sort.Interface` (3 methods) gives full custom sort control.

---

### 28. error Interface
**Q: What is the complete definition of the error interface?**
```go
// The entire built-in error interface:
type error interface {
    Error() string
}

// Any type with Error() string satisfies it:
type AppError struct {
    Code    int
    Message string
}
func (e AppError) Error() string {
    return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}
```
**A:** The `error` interface has exactly **one method**: `Error() string`. It's deliberately minimal — any struct with that method is an error.

---

### 29. Embedding Interface in Concrete Type — Partial Implementation
**Q: What happens?**
```go
package main
import "io"

// Struct that partially satisfies io.ReadWriter by embedding
type LoggingReadWriter struct {
    io.ReadWriter       // provides Read() and Write() from embedded
    prefix string
}

// Override Write to add logging:
func (l *LoggingReadWriter) Write(p []byte) (n int, err error) {
    println(l.prefix, string(p))
    return l.ReadWriter.Write(p)
}
```
**A:** Embedding an interface in a struct lets you override only the methods you care about while delegating the rest. Common pattern for decorators/middleware.

---

### 30. Interface vs Concrete Type in Return
**Q: What is the Go guideline?**
```go
// GOOD: return concrete type (callers can use full API)
func newBuffer() *bytes.Buffer {
    return &bytes.Buffer{}
}

// BAD: unnecessarily return interface (hides concrete type)
func newBuffer() io.Writer {
    return &bytes.Buffer{}
}

// Exception: return interface when multiple implementations possible
func newStore(useRedis bool) Cache {
    if useRedis {
        return &RedisCache{}
    }
    return &MemCache{}
}
```
**A:** **Accept interfaces, return concrete types.** Return interfaces only when the function intentionally abstracts over multiple implementations. This preserves the caller's ability to type-assert or access concrete methods.

---

### 31. Empty Interface (any) — Go 1.18+
**Q: What is the output?**
```go
package main
import "fmt"

func printType(v any) { // 'any' is alias for interface{}
    fmt.Printf("%T: %v\n", v, v)
}

func main() {
    printType(42)
    printType("hello")
    printType([]int{1, 2, 3})
    printType(nil)
}
```
**A:**
```
int: 42
string: hello
[]int: [1 2 3]
<nil>: <nil>
```
`any` is a type alias for `interface{}` introduced in Go 1.18 for readability.

---

### 32. fmt.Stringer vs fmt.GoStringer
**Q: What is the difference?**
```go
package main
import "fmt"

type Point struct{ X, Y int }

func (p Point) String() string { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }
func (p Point) GoString() string { return fmt.Sprintf("Point{X:%d, Y:%d}", p.X, p.Y) }

func main() {
    p := Point{3, 4}
    fmt.Printf("%v\n", p)   // calls String()
    fmt.Printf("%s\n", p)   // calls String()
    fmt.Printf("%#v\n", p)  // calls GoString()
}
```
**A:**
```
(3,4)
(3,4)
Point{X:3, Y:4}
```
`fmt.Stringer` (`%v`/`%s`) is for human-readable output. `fmt.GoStringer` (`%#v`) is for Go-syntax representation, useful in debugging.

---

## Section 3: Advanced Error Handling (Q33–Q42)

### 33. Wrapping Errors — Full Chain
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

type DBError struct{ table string }
func (e *DBError) Error() string { return "DB error on " + e.table }

func queryDB() error {
    return &DBError{table: "users"}
}

func getUser() error {
    if err := queryDB(); err != nil {
        return fmt.Errorf("getUser: %w", err)
    }
    return nil
}

func main() {
    err := getUser()
    fmt.Println(err)

    var dbErr *DBError
    fmt.Println(errors.As(err, &dbErr))
    fmt.Println(dbErr.table)
}
```
**A:**
```
getUser: DB error on users
true
users
```
`errors.As` unwraps the chain to find a `*DBError`, even if wrapped multiple times.

---

### 34. errors.Join (Go 1.20+)
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

func main() {
    err1 := errors.New("db connection failed")
    err2 := errors.New("cache timeout")
    
    combined := errors.Join(err1, err2)
    fmt.Println(combined)
    fmt.Println(errors.Is(combined, err1))
    fmt.Println(errors.Is(combined, err2))
}
```
**A:**
```
db connection failed
cache timeout
true
true
```
`errors.Join` (Go 1.20+) creates a multi-error that `errors.Is` and `errors.As` can unwrap individually.

---

### 35. Custom Error with Multiple Fields
**Q: What is the output?**
```go
package main
import "fmt"

type HTTPError struct {
    StatusCode int
    Status     string
    URL        string
}

func (e *HTTPError) Error() string {
    return fmt.Sprintf("HTTP %d %s: %s", e.StatusCode, e.Status, e.URL)
}

func fetch(url string) error {
    return &HTTPError{404, "Not Found", url}
}

func main() {
    err := fetch("https://example.com/data")
    fmt.Println(err)

    var he *HTTPError
    if errors.As(err, &he) { // uses errors package
        fmt.Println("status code:", he.StatusCode)
    }
}
```
**A:**
```
HTTP 404 Not Found: https://example.com/data
status code: 404
```

---

### 36. Panic vs Error — Decision Rule
**Q: Which should be used?**
```go
// Use panic for programming errors (should never happen in correct code)
func mustPositive(n int) int {
    if n <= 0 {
        panic(fmt.Sprintf("mustPositive: got %d", n))
    }
    return n
}

// Use error for expected failures (network, IO, user input)
func readFile(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("readFile %q: %w", path, err)
    }
    return data, nil
}
```
**A:** **Rule:** panic for invariant violations (bugs in YOUR code); error for recoverable external failures. Libraries almost never panic — they return errors. CLI tools and `init` code may panic defensively.

---

### 37. Sentinel Errors — Common stdlib Examples
**Q: What are the most important sentinel errors to know?**
```go
import (
    "errors"
    "io"
    "os"
)

// io.EOF — end of stream (expected, not really an error)
if err == io.EOF { ... }

// os.ErrNotExist — file not found
if errors.Is(err, os.ErrNotExist) { ... }

// os.ErrPermission — permission denied
if errors.Is(err, os.ErrPermission) { ... }

// context.DeadlineExceeded / context.Canceled
if errors.Is(err, context.DeadlineExceeded) { ... }
if errors.Is(err, context.Canceled) { ... }
```
**A:** These are the sentinel errors you'll handle daily. Always use `errors.Is` (not `==`) for wrapped errors.

---

### 38. Error Handling in Goroutines — Channel Pattern
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "errors"
)

func riskyWork(id int) error {
    if id == 2 {
        return errors.New("task 2 failed")
    }
    return nil
}

func main() {
    errc := make(chan error, 3)
    for i := 1; i <= 3; i++ {
        go func(id int) {
            errc <- riskyWork(id)
        }(i)
    }

    for i := 0; i < 3; i++ {
        if err := <-errc; err != nil {
            fmt.Println("error:", err)
        }
    }
}
```
**A:** `error: task 2 failed` (other goroutines send nil). Channel-based error collection from goroutines — alternative to `errgroup`.

---

### 39. Unwrap Chain with errors.Unwrap
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

func main() {
    base := errors.New("base error")
    wrapped := fmt.Errorf("layer 1: %w", base)
    wrapped2 := fmt.Errorf("layer 2: %w", wrapped)

    fmt.Println(errors.Unwrap(wrapped2)) // one level down
    fmt.Println(errors.Is(wrapped2, base)) // traverses full chain
}
```
**A:**
```
layer 1: base error
true
```

---

### 40. Must* Pattern for Initialization
**Q: What does this pattern do?**
```go
package main
import (
    "fmt"
    "regexp"
)

// Must panics on error — used for package-level init only
var emailRegex = regexp.MustCompile(`^[a-z0-9]+@[a-z]+\.[a-z]{2,}$`)

func main() {
    fmt.Println(emailRegex.MatchString("test@example.com"))
    fmt.Println(emailRegex.MatchString("invalid"))
}
```
**A:**
```
true
false
```
`Must*` functions (e.g., `regexp.MustCompile`, `template.Must`) panic if initialization fails — appropriate for compile-time-known values where failure is a programming error, not runtime condition.

---

### 41. Error Wrapping Best Practice
**Q: What should every error wrap include?**
```go
// BAD — no context, can't trace origin
return err

// BAD — duplicates error message
return fmt.Errorf("error: %v", err)

// GOOD — adds context, preserves chain with %w
return fmt.Errorf("userService.Create(name=%q): %w", name, err)

// GOOD — explicit wrapping
return fmt.Errorf("package.Function: operation failed for id=%d: %w", id, err)
```
**A:** Every wrapped error should include: package/function name + relevant parameters + `%w` to preserve chain. This creates readable stack-trace-like error messages without an external library.

---

### 42. errors.Is vs errors.As — When to Use Which
**Q: What is the difference?**
```go
var ErrNotFound = errors.New("not found")

// errors.Is: checks if error IS a specific value/sentinel
if errors.Is(err, ErrNotFound) { ... }         // identity check

// errors.As: checks if error IS a specific type, extracts it
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println(pathErr.Path) // access typed fields
}
```
**A:** `errors.Is` → check **identity** (is this specific error value?). `errors.As` → check **type** (is this a *PathError?) and extract typed fields. Both unwrap the full chain.

---

## Section 4: Generics — Go 1.18+ (Q43–Q50)

### 43. Generic Function — Basic Syntax
**Q: What is the output?**
```go
package main
import "fmt"

func Map[T, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

func main() {
    doubled := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
    fmt.Println(doubled)

    strs := Map([]int{1, 2, 3}, func(n int) string { return fmt.Sprintf("%d", n) })
    fmt.Println(strs)
}
```
**A:**
```
[2 4 6]
[1 2 3]
```

---

### 44. comparable Constraint
**Q: What is the output and why do we need `comparable`?**
```go
package main
import "fmt"

func Contains[T comparable](s []T, target T) bool {
    for _, v := range s {
        if v == target {
            return true
        }
    }
    return false
}

func main() {
    fmt.Println(Contains([]int{1, 2, 3}, 2))
    fmt.Println(Contains([]string{"a", "b"}, "c"))
}
```
**A:**
```
true
false
```
`comparable` is a built-in constraint allowing `==` and `!=`. Without it, `v == target` would be a compile error since not all types are comparable (e.g., slices, maps).

---

### 45. Type Constraints with Union
**Q: What is the output?**
```go
package main
import "fmt"

type Number interface {
    ~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
    fmt.Println(Sum([]float64{1.1, 2.2, 3.3}))
}
```
**A:**
```
15
6.6000000000000005
```
The `~` prefix means "any type whose underlying type is int" — includes named types like `type MyInt int`.

---

### 46. Generic Stack
**Q: What is the output?**
```go
package main
import "fmt"

type Stack[T any] struct {
    data []T
}

func (s *Stack[T]) Push(v T)    { s.data = append(s.data, v) }
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.data) == 0 {
        var zero T
        return zero, false
    }
    top := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return top, true
}

func main() {
    s := Stack[int]{}
    s.Push(1); s.Push(2); s.Push(3)
    v, _ := s.Pop()
    fmt.Println(v)
}
```
**A:** `3`. Generic types allow type-safe data structures without code duplication.

---

### 47. Filter Generic Function
**Q: What is the output?**
```go
package main
import "fmt"

func Filter[T any](s []T, pred func(T) bool) []T {
    var result []T
    for _, v := range s {
        if pred(v) {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    evens := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println(evens)
}
```
**A:** `[2 4 6]`

---

### 48. Reduce Generic Function
**Q: What is the output?**
```go
package main
import "fmt"

func Reduce[T, U any](s []T, init U, f func(U, T) U) U {
    acc := init
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}

func main() {
    sum := Reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, n int) int { return acc + n })
    fmt.Println(sum)

    concat := Reduce([]string{"go", " ", "rocks"}, "", func(acc, s string) string { return acc + s })
    fmt.Println(concat)
}
```
**A:**
```
15
go rocks
```

---

### 49. Generic Keys Extraction from Map
**Q: What is the output?**
```go
package main
import "fmt"

func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func main() {
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    fmt.Println(len(Keys(m)))
}
```
**A:** `3`. Before generics, you'd have to write this for each map type. Now one function works for all.

---

### 50. Type Inference in Generic Calls
**Q: Does this require explicit type arguments?**
```go
package main
import "fmt"

func First[T any](s []T) (T, bool) {
    if len(s) == 0 {
        var zero T
        return zero, false
    }
    return s[0], true
}

func main() {
    // Go infers T=int from argument
    v, ok := First([]int{10, 20, 30})
    fmt.Println(v, ok)

    // Explicit: First[string](...)
    v2, ok2 := First[string](nil)
    fmt.Println(v2, ok2)
}
```
**A:**
```
10 true
 false
```
Go infers the type parameter from the argument —`First[int]` is inferred automatically. Explicit type args (`First[string]`) are needed when inference isn't possible (e.g., nil).

---

*End of Part 2 — Context, Interfaces, Error Handling & Generics (50 questions)*

---

## Summary — Full 30% Coverage for Mid-Tier Product Companies

| File | Topics | Questions |
|---|---|---|
| `go_intermediate_concurrency.md` | Mutex, RWMutex, atomic, Worker Pool, Pipeline, Fan-In/Out, Goroutine lifecycle | 50 |
| `go_intermediate_context_interfaces.md` | Context, Interface composition, Advanced errors, Generics | 50 |
| **Total** | | **100** |
