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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference?
**Your Response:** Both are non-nil, empty contexts that never cancel. Background is the root context used at the top level of your application - in main, init, or as the base of a request tree. TODO is a placeholder you use when you know you need a context but haven't decided which one yet - it helps linters catch places where you forgot to pass a context properly.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints work stopped: context canceled. When we call cancel(), it closes the context's Done channel. The worker goroutine is waiting on ctx.Done() in a select, so it immediately receives the close signal and exits. Always call cancel() in a defer right after creating the context to ensure resources are cleaned up.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the resource leak if cancel is not called?
**Your Response:** If you don't call cancel(), Go leaks internal resources. The context package spawns internal goroutines and timers to handle cancellation. When you discard the cancel function, these resources can't be cleaned up. Go vet and staticcheck will flag this as an error. Always call cancel() in defer immediately after creating the context.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints context deadline exceeded. The fetch function simulates work that takes 500ms, but we create a context with a 200ms timeout. When the context expires, it cancels the Done channel, and fetch returns the context's error. This is the standard way to implement timeouts in Go operations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference?
**Your Response:** Both create contexts that expire after 5 seconds, but they're functionally equivalent here. WithTimeout takes a duration relative to now, while WithDeadline takes an absolute time. WithTimeout is just shorthand for WithDeadline with time.Now().Add(duration). Use WithDeadline when you have an absolute deadline from an incoming request or external system.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints deadline in: 10s. The Deadline() method returns the absolute deadline time and a boolean indicating if a deadline is set. This is useful when you need to know how much time is left or pass a reduced timeout to downstream calls. You can calculate the remaining time with time.Until(deadline).

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why must context be the first parameter by convention?
**Your Response:** This prints context canceled. Context flows downward through the call chain - each function receives it as the first parameter and passes it to the next. When we cancel the context before calling step1, every function in the chain can detect the cancellation via ctx.Err() and exit early. This convention makes APIs consistent and ensures cancellation propagates properly through all layers.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints request ID: abc-123. We use context.WithValue to store request-scoped data. The key is an unexported custom type (key) with a constant value - this prevents collisions with other packages that might also use string keys. Always use unexported types as context keys to avoid namespace conflicts.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints <nil> then true. When you call Value() on a key that doesn't exist in the context, it returns nil. Always check for nil before type-asserting the value, otherwise you'll panic if the key is missing. This is why context.Value returns an interface{} - you need to handle the nil case.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what is the output?
**Your Response:** This compiles and prints from-A then from-B. Even though both keys have the same underlying string value "id", they're different Go types (pkgAKey vs pkgBKey). Context keys are compared by their exact type and value, not just the underlying string. This is why you must use custom types as keys - it prevents collisions between different packages.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the standard pattern to propagate context in HTTP handlers?
**Your Response:** The standard pattern is to get the context from the incoming request with r.Context(). This context carries cancellation signals from the HTTP client - if the client disconnects, the context is cancelled. Always derive child contexts from r.Context(), not from Background(), to preserve this cancellation behavior throughout the request lifecycle.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** All 3 workers print exiting: context canceled, though the order may vary. When we call cancel(), it closes the context's Done channel. All goroutines waiting on ctx.Done() unblock simultaneously and receive the cancellation error. This is how you broadcast cancellation to multiple goroutines at once - a single cancel call stops all listeners.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints context canceled. When we cancel the parent context, the cancellation automatically propagates to all child contexts, even if the child has its own timeout. The child's deadline is effectively overridden by the parent's cancellation. This hierarchical cancellation ensures that cancelling at any level stops all downstream operations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints context canceled then shutdown requested. WithCancelCause is new in Go 1.20 and lets you attach a specific error when cancelling. The context's Err() still returns the generic context.Canceled, but context.Cause() returns the specific error you passed to cancel(). This provides more detailed information about why cancellation occurred.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is wrong with this?
**Your Response:** Storing context in a struct is an anti-pattern. Context should be request-scoped and passed as the first parameter to each function. If you store it in a struct, the same context persists across all operations, preventing per-request cancellation. This violates Go's context propagation guidelines and makes your code harder to test and reason about.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints true then <nil>. The Background context never cancels, so its Done() method returns nil. In a select statement, receiving from a nil channel blocks forever. This means code that checks ctx.Done() will always wait when using Background(), which is correct since Background() is meant for operations that should never be cancelled.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern for database calls with context?
**Your Response:** Always use the *Context variants of database methods like QueryRowContext. These accept a context and will cancel the query at the driver level if the context expires. This prevents long-running queries from continuing after a timeout or cancellation. In production, always pass the context through to database operations to ensure proper cancellation behavior.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What should NOT be stored in context?
**Your Response:** Context should only store optional, request-scoped cross-cutting concerns like trace IDs, auth tokens, or deadlines. Never use context to pass required function parameters like userID or database connections. This makes APIs implicit and hard to test since the dependencies aren't visible in the function signature. Required parameters should always be explicit arguments.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you detect client disconnect?
**Your Response:** When an HTTP client disconnects mid-stream, the request's context is automatically cancelled. By checking ctx.Done() in your streaming loop, you can detect when the client has disconnected and stop processing immediately. This prevents wasting resources generating data that nobody will receive and properly cleans up resources.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which case is checked first in select?
**Your Response:** Neither case has priority in Go's select - when multiple cases are ready, it picks one randomly. However, by convention we place ctx.Done() first as a style signal to readers that cancellation is important. If you need actual priority, use a nested select - first check ctx.Done() in a non-blocking select, then read from the channel if not cancelled.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this compile to?
**Your Response:** This prints hello. Interface composition lets you embed multiple interfaces into one. ReadWriter embeds Reader and Writer, so any type that satisfies both automatically satisfies ReadWriter. The Buffer type implements Read() and Write() methods, so it satisfies all three interfaces without any explicit declaration.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints copied 10 bytes: Hello, Go!. This demonstrates the power of Go interfaces - strings.Reader and bytes.Buffer satisfy io.Reader and io.Writer interfaces without explicitly declaring them. The copyData function works with any types that implement these interfaces, making it highly reusable and flexible.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** This shows how embedding smaller interfaces creates composable APIs. io.ReadCloser embeds both Reader and Closer interfaces, requiring implementations to provide both capabilities. http.Response.Body is a ReadCloser because you need to both read the body and close it to free resources. This pattern ensures all required behaviors are present.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which function signature is better?
**Your Response:** The io.Reader version is much better because it accepts any type that can be read, not just files. This follows the Go principle: accept interfaces, return concrete types. With io.Reader, the function works with files, HTTP bodies, strings, network connections, and any custom type that implements Read(). This makes your code much more flexible and testable.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints woof then Rex. The Dog type implements both Sound() and Name() methods, so it automatically satisfies both Animal and Namer interfaces without explicitly declaring them. Go's structural typing means any type with the required methods satisfies the interface, enabling clean composition without inheritance.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints data then data [flushed]. The maybeFlush function accepts any type with a Write method. It uses a type assertion to check if the type also implements Flusher. If it does, it calls Flush(). This optional interface pattern is used extensively in net/http for capabilities like HTTP flushing or connection hijacking that not all response writers support.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints [fig kiwi peach apple]. The ByLength type implements sort.Interface by providing Len(), Less(), and Swap() methods. This gives you full control over how elements are sorted - in this case, by string length rather than alphabetically. The sort package works with any type implementing these three methods.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the complete definition of the error interface?
**Your Response:** The error interface has exactly one method: Error() string. It's deliberately minimal - any type that implements this method is an error. This simplicity makes it easy to create custom error types with additional fields and methods while still satisfying the interface. The AppError example shows how to add structured data to errors.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens?
**Your Response:** This shows how to embed an interface in a struct to partially implement it. The LoggingReadWriter embeds io.ReadWriter, which provides Read() and Write() methods. We override Write() to add logging while delegating to the embedded ReadWriter. This is a common pattern for decorators and middleware where you want to modify specific behavior.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the Go guideline?
**Your Response:** The guideline is: accept interfaces, return concrete types. Return concrete types so callers have full access to the type's API and can type-assert if needed. Only return interfaces when you're intentionally hiding the implementation and there are multiple possible implementations. This gives callers maximum flexibility while still allowing abstraction when appropriate.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints the type and value of each argument. 'any' is a type alias for interface{} introduced in Go 1.18 for better readability. It can hold any type, making it useful for generic programming or when you need to handle values of unknown types. The %T verb prints the type, and %v prints the value.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference?
**Your Response:** This shows the difference between Stringer and GoStringer interfaces. Stringer's String() method is called with %v and %s for human-readable output. GoStringer's GoString() method is called with %#v for Go syntax representation, useful in debugging and logging where you want to see the struct field names.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints getUser: DB error on users, then true, then users. The fmt.Errorf with %w wraps the error while preserving the original. errors.As unwraps the entire chain to find a *DBError, even through multiple layers of wrapping. This allows you to extract and inspect specific error types while maintaining a full error context chain.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints db connection failed, cache timeout, then true twice. errors.Join is new in Go 1.20 and combines multiple errors into one. When printed, it shows all errors. errors.Is can check if the multi-error contains a specific error, returning true for each component error. This is useful for operations that might fail for multiple reasons.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints the full error message then status code: 404. The HTTPError struct implements the error interface with a custom Error() method. Using errors.As, we can extract the concrete error type and access its fields like StatusCode. This pattern allows structured error data while maintaining the simple error interface.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which should be used?
**Your Response:** The rule is: panic for programming errors and invariant violations - things that should never happen in correct code. Use errors for expected failures like network issues, IO problems, or invalid user input. Libraries should almost never panic - they should return errors. Panics are for unrecoverable bugs where continuing would be dangerous.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the most important sentinel errors to know?
**Your Response:** These are the most important sentinel errors in Go's standard library. io.EOF indicates end of stream - it's expected, not really an error. os.ErrNotExist and os.ErrPermission are for file operations. context.DeadlineExceeded and context.Canceled are for context cancellations. Always use errors.Is to check for these errors since they might be wrapped.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints error: task 2 failed. The pattern uses a buffered channel to collect errors from multiple goroutines. Each goroutine sends its error (or nil) to the channel. The main function receives all errors and processes them. This is an alternative to errgroup when you need custom error handling or want to process all errors, not just the first one.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints layer 1: base error then true. errors.Unwrap removes one level of wrapping, revealing the immediate wrapped error. errors.Is traverses the entire chain to check if any error in the chain matches the target. This shows the difference between examining just the next level versus searching the full chain.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this pattern do?
**Your Response:** This prints true then false. Must* functions like regexp.MustCompile panic if the input is invalid. This pattern is used for package-level initialization where the input is known at compile time. If the regex is invalid, it's a programming error, so panicking is appropriate. This avoids error handling for values that should never fail.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What should every error wrap include?
**Your Response:** Every wrapped error should include the package and function name, relevant parameters like IDs or inputs, and use %w to preserve the error chain. This creates a readable trace showing where the error occurred and what data was involved. The pattern package.Function: details for operation: %w gives context while maintaining the ability to inspect underlying errors.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference?
**Your Response:** errors.Is checks for identity - whether the error is exactly a specific value like io.EOF or a custom sentinel error. errors.As checks for type and extracts the value - whether the error is a *PathError and gives you access to its fields. Both traverse the full error chain, unwrapping as needed. Use Is for sentinel errors, As for structured error types.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints [2 4 6] then [1 2 3]. The Map function is generic with two type parameters T and U. It transforms a slice of T values to U values using a function. First it doubles integers, then converts integers to strings. Generics eliminate code duplication - one Map function works for any type transformation.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output and why do we need `comparable`?
**Your Response:** This prints true then false. The comparable constraint is built-in and allows types that support == and != operations. Without this constraint, the comparison v == target would be a compile error because not all types in Go are comparable - slices, maps, and functions can't be compared. The comparable constraint ensures only types that can be compared are accepted.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 15 then 6.6000000000000005. The Number interface uses a union of types with ~ prefix, meaning it accepts any type whose underlying type is int, int64, or float64. This includes named types like type MyInt int. The ~ operator makes the constraint more flexible by including underlying types, not just exact types.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 3. The Stack is a generic type parameterized by T. It can hold any type while maintaining type safety. The Push method takes a T value, and Pop returns a T value. This eliminates code duplication - one Stack implementation works for all types, unlike pre-generics where you'd need separate implementations or use interface{} with type assertions.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints [2 4 6]. The Filter function is generic and works with any type T. It takes a slice of T values and a predicate function that returns a boolean. It returns a new slice containing only the values that satisfy the predicate. This shows how generics can create reusable algorithms that work with any type.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 15 then go rocks. Reduce is a generic function with two type parameters - T for the slice elements and U for the accumulator. It accumulates values using a binary function. First it sums integers, then it concatenates strings. This demonstrates how generics can implement common functional programming patterns in a type-safe way.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 3. The Keys function is generic over any map with a comparable key type K and any value type V. It extracts just the keys from the map. Before generics, you'd need to write separate functions for each map type. Now one generic function works for all maps, making code much more reusable.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this require explicit type arguments?
**Your Response:** This prints 10 true then <nil> false. Go's type inference automatically determines the type parameter from the function arguments - First[int] is inferred from the []int slice. When passing nil, inference isn't possible, so we need explicit type arguments like First[string]. Type inference makes generic functions feel natural when the types can be determined from context.

---

*End of Part 2 — Context, Interfaces, Error Handling & Generics (50 questions)*

---

## Summary — Full 30% Coverage for Mid-Tier Product Companies

| File | Topics | Questions |
|---|---|---|
| `go_intermediate_concurrency.md` | Mutex, RWMutex, atomic, Worker Pool, Pipeline, Fan-In/Out, Goroutine lifecycle | 50 |
| `go_intermediate_context_interfaces.md` | Context, Interface composition, Advanced errors, Generics | 50 |
| **Total** | | **100** |
