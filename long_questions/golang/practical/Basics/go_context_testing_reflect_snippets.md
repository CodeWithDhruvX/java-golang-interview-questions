# Go — Context, Goroutine Leaks, Strings Builder, Testing & Reflect Snippets

> **Format**: Each question is "predict the output / spot the bug / does it compile?" style.
> **Topics**: `context.Context` · Goroutine leaks · `strings.Builder` / `bytes.Buffer` · Table-driven tests · `testing.T` · Benchmarks · `reflect` basics · Pipeline pattern

---

## 📋 Reading Progress

> Mark each section `[x]` when done. Use `🔖` to note where you left off.

- [ ] **Section 1:** context.Context (Q1–Q22)
- [ ] **Section 2:** Goroutine Leaks (Q23–Q32)
- [ ] **Section 3:** strings.Builder & bytes.Buffer (Q33–Q44)
- [ ] **Section 4:** Table-Driven Tests & testing.T (Q45–Q60)
- [ ] **Section 5:** reflect Basics (Q61–Q72)

> 🔖 **Last read:** <!-- e.g. Q22 · Section 1 done -->

---

## Section 1: context.Context (Q1–Q22)

### 1. context.Background vs context.TODO
**Q: What is the difference?**
```go
package main
import (
    "context"
    "fmt"
)

func main() {
    bg  := context.Background()
    todo := context.TODO()
    fmt.Println(bg != nil, todo != nil)
}
```
**A:** `true true`. Both are non-nil empty contexts that are never cancelled. `Background` is the root context for production code; `TODO` signals "we plan to add a real context later" — a placeholder used during refactoring.

---

### 2. context.WithCancel — Basic Usage
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        <-ctx.Done()
        fmt.Println("goroutine done:", ctx.Err())
    }()

    cancel()
    // give goroutine time to print (in real code use sync)
    fmt.Println("main cancelled")
}
```
**A:** (order may vary)
```
main cancelled
goroutine done: context canceled
```
`cancel()` closes `ctx.Done()`. `ctx.Err()` returns `context.Canceled`.

---

### 3. ctx.Done() Is a Channel
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

func main() {
    ctx := context.Background()
    fmt.Println(ctx.Done() == nil)

    ctx2, cancel := context.WithCancel(ctx)
    defer cancel()
    fmt.Println(ctx2.Done() == nil)
}
```
**A:**
```
true
false
```
`context.Background()` returns a context whose `Done()` is nil (never cancelled). `WithCancel` returns one with a real channel.

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

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel()

    select {
    case <-time.After(200 * time.Millisecond):
        fmt.Println("operation done")
    case <-ctx.Done():
        fmt.Println("timed out:", ctx.Err())
    }
}
```
**A:** `timed out: context deadline exceeded`. The 50ms timeout fires before the 200ms timer.

---

### 5. context.WithDeadline
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func main() {
    deadline := time.Now().Add(30 * time.Millisecond)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()

    <-ctx.Done()
    fmt.Println(ctx.Err())
    fmt.Println(ctx.Err() == context.DeadlineExceeded)
}
```
**A:**
```
context deadline exceeded
true
```

---

### 6. Always Call cancel()
**Q: What is the bug?**
```go
func doWork() {
    ctx, _ := context.WithCancel(context.Background()) // BUG: cancel discarded
    // ... use ctx
    _ = ctx
}
```
**A:** **Resource leak.** The cancel function must always be called (via `defer cancel()`) to release resources associated with the context, even if the context times out. The linter `go vet` and `staticcheck` flag this pattern.

---

### 7. Passing Context to Functions — Convention
**Q: Does this compile and what is the convention?**
```go
package main
import (
    "context"
    "fmt"
)

func fetchUser(ctx context.Context, id int) (string, error) {
    select {
    case <-ctx.Done():
        return "", ctx.Err()
    default:
        return fmt.Sprintf("user-%d", id), nil
    }
}

func main() {
    ctx := context.Background()
    name, err := fetchUser(ctx, 42)
    fmt.Println(name, err)
}
```
**A:** `user-42 <nil>`. Convention: `context.Context` is always the **first parameter**, named `ctx`, never stored in a struct.

---

### 8. context.WithValue — Storing and Retrieving
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

type key string

func main() {
    ctx := context.WithValue(context.Background(), key("userID"), 101)
    fmt.Println(ctx.Value(key("userID")))
    fmt.Println(ctx.Value(key("other")))
}
```
**A:**
```
101
<nil>
```
`Value` returns `nil` for keys not in the context chain.

---

### 9. context.WithValue — Use Unexported Key Type (Best Practice)
**Q: Why is this the correct pattern?**
```go
package mypackage

// BAD: uses a plain string key — risk of collision across packages
ctx = context.WithValue(ctx, "userID", 42)

// GOOD: uses a package-private type as key — no collision possible
type ctxKey string
const userKey ctxKey = "userID"
ctx = context.WithValue(ctx, userKey, 42)
```
**A:** Using a plain `string` key means any package could accidentally read or overwrite it. A package-private type makes the key unique to the package — other packages cannot construct the same key type.

---

### 10. Propagating Cancellation Through Goroutines
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
    select {
    case <-ctx.Done():
        fmt.Printf("worker %d cancelled\n", id)
    }
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
**A:** All 3 workers print "cancelled" (order varies). A single `cancel()` call propagates to all goroutines watching `ctx.Done()`.

---

### 11. Child Context Cancelled When Parent is Cancelled
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

func main() {
    parent, cancelParent := context.WithCancel(context.Background())
    child, cancelChild := context.WithCancel(parent)
    defer cancelChild()

    cancelParent()

    <-child.Done()
    fmt.Println("child done:", child.Err())
}
```
**A:** `child done: context canceled`. Cancelling a parent automatically cancels all derived child contexts.

---

### 12. Cancelling Child Does NOT Cancel Parent
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

func main() {
    parent, cancelParent := context.WithCancel(context.Background())
    defer cancelParent()

    child, cancelChild := context.WithCancel(parent)
    cancelChild()

    select {
    case <-parent.Done():
        fmt.Println("parent cancelled")
    default:
        fmt.Println("parent still active")
    }
}
```
**A:** `parent still active`. Cancellation flows downward (parent → child), never upward.

---

### 13. context.WithTimeout — cancel Must Still Be Called
**Q: What is the best practice?**
```go
// CORRECT
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel() // releases resources if operation finishes before timeout

// WRONG — leaks resources if the function returns before the deadline
ctx, _ = context.WithTimeout(parent, 5*time.Second)
```
**A:** Even when using `WithTimeout`, always call `defer cancel()`. If the function completes before the timeout, the cancel frees the timer resources immediately instead of waiting for the deadline.

---

### 14. Checking ctx.Err() Without Blocking
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    fmt.Println(ctx.Err())
    cancel()
    fmt.Println(ctx.Err())
}
```
**A:**
```
<nil>
context canceled
```
`ctx.Err()` is nil while active; returns the cancellation error after cancellation.

---

### 15. Passing Context to http.Request
**Q: What does this do?**
```go
package main
import (
    "context"
    "net/http"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        // err wraps context.DeadlineExceeded if timeout hit
        return
    }
    defer resp.Body.Close()
}
```
**A:** The HTTP request is automatically cancelled if it takes longer than 3 seconds. `http.NewRequestWithContext` is the idiomatic way to attach a context to an outgoing request.

---

### 16. context.WithValue — Value Not Found Returns nil (No Panic)
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

---

### 17. Deadline() Method
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx1 := context.Background()
    d1, ok1 := ctx1.Deadline()
    fmt.Println(d1, ok1)

    ctx2, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))
    defer cancel()
    _, ok2 := ctx2.Deadline()
    fmt.Println(ok2)
}
```
**A:**
```
0001-01-01 00:00:00 +0000 UTC false
true
```
`Background()` has no deadline (`ok=false`). A `WithDeadline` context has one (`ok=true`).

---

### 18. Long-Running Loop With Context Check
**Q: What is the pattern?**
```go
package main
import (
    "context"
    "fmt"
)

func process(ctx context.Context, items []int) error {
    for _, item := range items {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        fmt.Println("processing:", item)
    }
    return nil
}
```
**A:** The `select`/`default` idiom is a non-blocking context check inside a loop — allows the loop to be cancelled between iterations without blocking.

---

### 19. errgroup for Concurrent Tasks with Cancellation
**Q: What is the output (assuming golang.org/x/sync)?**
```go
// import "golang.org/x/sync/errgroup"
g, ctx := errgroup.WithContext(context.Background())

g.Go(func() error {
    <-ctx.Done()
    return ctx.Err()
})

g.Go(func() error {
    return fmt.Errorf("task failed")
})

err := g.Wait()
fmt.Println(err)
```
**A:** `task failed`. When any goroutine in the group returns an error, the group's context is cancelled, signalling all other goroutines. `g.Wait()` returns the first non-nil error.

---

### 20. context.TODO vs context.Background — When to Use Which
**Q: Match the usage:**
```
A. Top-level main() or server entry point
B. Called from a function that doesn't yet have a context parameter
C. Tests with no real context
```
**A:**
- A → `context.Background()`
- B → `context.TODO()` (signals intent to propagate context later)
- C → `context.Background()` (or `t.Context()` in Go 1.21+)

---

### 21. context.WithCancelCause (Go 1.20+)
**Q: What is the output?**
```go
package main
import (
    "context"
    "errors"
    "fmt"
)

func main() {
    ctx, cancel := context.WithCancelCause(context.Background())
    cancel(errors.New("disk full"))

    fmt.Println(ctx.Err())
    fmt.Println(context.Cause(ctx))
}
```
**A:**
```
context canceled
disk full
```
`WithCancelCause` lets you attach a reason to the cancellation. `context.Cause` retrieves it; `ctx.Err()` still returns the standard `context.Canceled`.

---

### 22. Never Store context in a Struct
**Q: Which is the correct pattern?**
```go
// BAD: storing context in struct couples lifetime and makes it hard to trace
type Server struct {
    ctx context.Context
}

// GOOD: pass context explicitly per-method call
type Server struct{}
func (s *Server) HandleRequest(ctx context.Context, req Request) error {
    // use ctx here
    return nil
}
```
**A:** The Go team explicitly states: *"Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it."* — `context` package docs.

---

## Section 2: Goroutine Leaks (Q23–Q32)

### 23. Classic Goroutine Leak — Unbuffered Channel, No Reader
**Q: What is the bug?**
```go
package main
import "fmt"

func leaky() {
    ch := make(chan int)
    go func() {
        ch <- 42 // blocks forever — nobody reads
    }()
    fmt.Println("leaky returned")
}

func main() { leaky() }
```
**A:** The goroutine blocks on `ch <- 42` forever because `leaky()` returns without reading from `ch`. This is a **goroutine leak** — detectable with `runtime.NumGoroutine()` or `goleak`.

---

### 24. Fix: Use Buffered Channel
**Q: Does this fix the leak?**
```go
func notLeaky() {
    ch := make(chan int, 1) // buffered — send doesn't block
    go func() {
        ch <- 42
    }()
}
```
**A:** **Yes.** The goroutine sends and exits immediately. The value in the channel will be GC'd when `ch` goes out of scope.

---

### 25. Leak — Goroutine Waiting on Channel That's Never Sent To
**Q: What is the bug?**
```go
func subscribe() <-chan int {
    ch := make(chan int)
    // forgot to send anything or close ch
    return ch
}

func main() {
    ch := subscribe()
    go func() {
        v := <-ch // blocks forever
        _ = v
    }()
}
```
**A:** The goroutine blocks forever waiting on `ch`. Always ensure every receiver has a corresponding send or `close`.

---

### 26. Fix Using done Channel / Context
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
)

func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("worker exited")
            return
        default:
            // do work
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go worker(ctx)
    cancel()
    // give goroutine time to exit
    fmt.Println("done")
}
```
**A:** `done` and `worker exited` (order may vary). Passing a cancellable context is the idiomatic fix for goroutine leaks.

---

### 27. Leak — http.Response Body Not Closed
**Q: What is the bug?**
```go
resp, err := http.Get("https://example.com")
if err != nil {
    return err
}
// forgot resp.Body.Close() — goroutine + connection leak
data, _ := io.ReadAll(resp.Body)
_ = data
```
**A:** Not closing `resp.Body` leaks the underlying TCP connection and the goroutine managing it. Always `defer resp.Body.Close()` immediately after the nil-error check.

---

### 28. Detecting Goroutine Leaks — runtime.NumGoroutine
**Q: What does this pattern do?**
```go
func TestNoLeak(t *testing.T) {
    before := runtime.NumGoroutine()
    // run code under test
    doSomething()
    runtime.Gosched() // yield to let goroutines settle
    after := runtime.NumGoroutine()
    if after > before {
        t.Errorf("goroutine leak: before=%d after=%d", before, after)
    }
}
```
**A:** Measures running goroutines before and after a function. A library like `go.uber.org/goleak` automates this more reliably in tests.

---

### 29. Goroutine Leak in HTTP Handler
**Q: What is the bug?**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    ch := make(chan Result)
    go func() {
        ch <- doHeavyWork() // if client disconnects, nobody reads ch
    }()
    result := <-ch
    _ = result
}
```
**A:** If the client disconnects, `handler` returns, but the goroutine is stuck sending on `ch` with no reader. Fix: use `r.Context()` to cancel the work.

---

### 30. Goroutine Leak — Range Over Channel, Never Closed
**Q: What is the bug?**
```go
func stream() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        // forgot close(ch)
    }()
    return ch
}

func main() {
    for v := range stream() { // blocks after 5 values — never exits
        _ = v
    }
}
```
**A:** `range` over a channel blocks until the channel is closed. Without `close(ch)`, `main` hangs forever. Fix: add `close(ch)` after the loop inside the goroutine.

---

### 31. Goroutine Leak — select With No Default and No Cancel
**Q: What is the bug?**
```go
func monitor(ch <-chan int) {
    go func() {
        for {
            select {
            case v := <-ch:
                fmt.Println(v)
            }
        }
    }()
}
```
**A:** If `ch` is never closed and the caller stops sending, the goroutine blocks inside `select` forever. Fix: add a `done <-chan struct{}` parameter and a `case <-done: return`.

---

### 32. goleak — Idiomatic Leak Detection
**Q: What does this do?**
```go
import "go.uber.org/goleak"

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}
```
**A:** `goleak.VerifyTestMain` checks that no goroutines are leaked after every test in the package. It fails the test suite if any unexpected goroutines remain after tests complete.

---

## Section 3: strings.Builder & bytes.Buffer (Q33–Q44)

### 33. String Concatenation in a Loop — Why It's Slow
**Q: What is the complexity issue?**
```go
func buildSlow(n int) string {
    s := ""
    for i := 0; i < n; i++ {
        s += fmt.Sprintf("%d", i) // allocates a new string every iteration
    }
    return s
}
```
**A:** Each `+=` allocates a brand-new string (strings are immutable), copying all previous content. This is **O(n²)** in time and allocations. Use `strings.Builder` instead.

---

### 34. strings.Builder — Efficient Concatenation
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    var b strings.Builder
    for i := 0; i < 5; i++ {
        fmt.Fprintf(&b, "%d", i)
    }
    fmt.Println(b.String())
    fmt.Println(b.Len())
}
```
**A:**
```
01234
5
```
`strings.Builder` accumulates writes without copying. `String()` returns the final string with zero extra allocation.

---

### 35. strings.Builder Reset
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    var b strings.Builder
    b.WriteString("hello")
    fmt.Println(b.String())
    b.Reset()
    b.WriteString("world")
    fmt.Println(b.String())
}
```
**A:**
```
hello
world
```
`Reset()` discards the buffer contents and resets the builder for reuse.

---

### 36. strings.Builder — Copying Is Forbidden
**Q: What is the bug?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    var b strings.Builder
    b.WriteString("hello")
    b2 := b // copy — FORBIDDEN
    b2.WriteString(" world")
    fmt.Println(b2.String())
}
```
**A:** **Runtime panic**: `strings.Builder.copyCheck` detects the copy and panics: `strings: illegal use of non-zero Builder copied by value`. Never copy a non-zero `Builder`.

---

### 37. bytes.Buffer — Write and Read Back
**Q: What is the output?**
```go
package main
import (
    "bytes"
    "fmt"
)

func main() {
    var buf bytes.Buffer
    buf.WriteString("Hello, ")
    buf.WriteString("Go!")
    fmt.Println(buf.String())
    fmt.Println(buf.Len())
}
```
**A:**
```
Hello, Go!
10
```

---

### 38. bytes.Buffer vs strings.Builder — Key Difference
**Q: What is the key difference?**
```go
// strings.Builder: write-only, produces a string at the end
var sb strings.Builder
sb.WriteString("data")
result := sb.String() // read once

// bytes.Buffer: read AND write; implements io.Reader and io.Writer
var buf bytes.Buffer
buf.WriteString("data")
b := make([]byte, 4)
buf.Read(b) // can read back
```
**A:** `strings.Builder` is optimised for building strings — write-only, no reads until `String()`. `bytes.Buffer` supports both reads and writes, making it suitable as an `io.Reader`/`io.Writer` for in-memory I/O.

---

### 39. strings.Builder with WriteRune / WriteByte
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    var b strings.Builder
    b.WriteByte('H')
    b.WriteRune('i')
    b.WriteString("!")
    fmt.Println(b.String())
}
```
**A:** `Hi!`

---

### 40. bytes.Buffer as io.Writer for fmt.Fprintf
**Q: What is the output?**
```go
package main
import (
    "bytes"
    "fmt"
)

func main() {
    var buf bytes.Buffer
    fmt.Fprintf(&buf, "name: %s, age: %d", "Alice", 30)
    fmt.Println(buf.String())
}
```
**A:** `name: Alice, age: 30`. `bytes.Buffer` satisfies `io.Writer`, so `fmt.Fprintf` writes directly to it.

---

### 41. strings.NewReader — String as io.Reader
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, Go!")
    buf := make([]byte, 5)
    n, _ := r.Read(buf)
    fmt.Println(string(buf[:n]))
    fmt.Println(r.Len())
}
```
**A:**
```
Hello
5
```
`strings.NewReader` wraps a string as an `io.Reader`. `Len()` returns remaining unread bytes.

---

### 42. String Immutability — Modifying a Byte
**Q: Does this compile?**
```go
package main

func main() {
    s := "hello"
    s[0] = 'H'
}
```
**A:** **Compile Error.** `cannot assign to s[0] (strings are not addressable)`. Strings in Go are immutable. Convert to `[]byte`, modify, then convert back.

---

### 43. []byte ↔ string Conversion
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "hello"
    b := []byte(s)
    b[0] = 'H'
    s2 := string(b)
    fmt.Println(s)
    fmt.Println(s2)
}
```
**A:**
```
hello
Hello
```
`[]byte(s)` copies the string data into a mutable byte slice. Modifying `b` does not affect the original `s`.

---

### 44. strings.Builder Grows as Needed — Preallocate with Grow
**Q: What is the pattern?**
```go
package main
import "strings"

func buildLarge(parts []string) string {
    var b strings.Builder
    total := 0
    for _, p := range parts { total += len(p) }
    b.Grow(total) // preallocate
    for _, p := range parts {
        b.WriteString(p)
    }
    return b.String()
}
```
**A:** `b.Grow(n)` hints the builder to pre-allocate at least `n` bytes, preventing repeated reallocations. A best-practice optimisation when the final size is known upfront.

---

## Section 4: Table-Driven Tests & testing.T (Q45–Q60)

### 45. Basic Table-Driven Test Structure
**Q: What is the output when `go test` runs this?**
```go
package mymath

import "testing"

func Add(a, b int) int { return a + b }

func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
        {"zero",     0, 0,  0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.expected {
                t.Errorf("Add(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```
**A:** `ok mymath` — all subtests pass. `t.Run` creates named subtests; you can run a single one with `go test -run TestAdd/negative`.

---

### 46. t.Error vs t.Fatal
**Q: What is the difference?**
```go
t.Error("something went wrong") // marks test as failed; continues execution
t.Fatal("critical failure")     // marks test as failed; stops this test immediately
```
**A:** `t.Error` / `t.Errorf` log the failure and continue. `t.Fatal` / `t.Fatalf` call `t.FailNow()` which stops the current test function (via `runtime.Goexit()`). Use `Fatal` when further execution is meaningless after the failure.

---

### 47. t.Run Subtests — Run a Single Subtest
**Q: What command runs only the `negative` subtest?**
```
go test -run TestAdd/negative
```
**A:** Runs only the `negative` subtest. The pattern is `TestFunctionName/subtest_name`. Spaces in subtest names become underscores in the `-run` pattern.

---

### 48. t.Parallel — Parallel Subtests
**Q: What does t.Parallel() do?**
```go
func TestAll(t *testing.T) {
    tests := []struct{ name string }{ {"a"}, {"b"}, {"c"} }
    for _, tt := range tests {
        tt := tt // capture range variable (pre Go 1.22)
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            // test runs concurrently with other parallel subtests
        })
    }
}
```
**A:** `t.Parallel()` signals that this subtest can run concurrently with other parallel subtests. The outer test function waits for all parallel subtests to finish before returning. Note the `tt := tt` copy needed pre-Go 1.22.

---

### 49. TestMain — Setup and Teardown
**Q: What is the purpose of TestMain?**
```go
package mypackage

import (
    "os"
    "testing"
)

func TestMain(m *testing.M) {
    // global setup
    setup()
    code := m.Run() // run all tests
    // global teardown
    teardown()
    os.Exit(code)
}
```
**A:** `TestMain` is the entry point for the test binary of a package. It allows global setup/teardown (e.g., starting a test DB, mock server) that runs once for all tests in the package. `m.Run()` actually executes the tests.

---

### 50. Testing Helper Functions — t.Helper()
**Q: What does t.Helper() do?**
```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper() // marks this as a helper
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestSomething(t *testing.T) {
    assertEqual(t, 1+1, 2)
    assertEqual(t, 1+1, 3) // line reported here, not inside assertEqual
}
```
**A:** `t.Helper()` marks the function as a test helper. Failure messages report the **caller's** line number (in `TestSomething`), not the helper's line — making test output much easier to read.

---

### 51. Benchmark Function
**Q: What is the structure of a benchmark?**
```go
package mymath

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}
```
**A:** `b.N` is set by the testing framework to a value large enough to produce a stable measurement. Run with `go test -bench=BenchmarkAdd -benchmem`. `-benchmem` shows allocations per operation.

---

### 52. b.ResetTimer — Excluding Setup Time
**Q: What is the pattern?**
```go
func BenchmarkProcess(b *testing.B) {
    data := generateLargeData() // expensive setup
    b.ResetTimer()              // don't count setup in benchmark timing
    for i := 0; i < b.N; i++ {
        process(data)
    }
}
```
**A:** `b.ResetTimer()` resets the elapsed time and memory counters. Useful when setup is expensive and should not pollute the benchmark measurement.

---

### 53. t.Cleanup
**Q: What is the output?**
```go
func TestWithCleanup(t *testing.T) {
    t.Cleanup(func() {
        fmt.Println("cleanup called")
    })
    fmt.Println("test body")
}
```
**A:**
```
test body
cleanup called
```
`t.Cleanup` registers a function that runs when the test (or subtest) finishes — including on failure. Multiple `Cleanup` calls run in LIFO order. Prefer over deferred cleanup in table-driven tests.

---

### 54. t.Setenv — Safe Environment Variable in Tests
**Q: What is the advantage?**
```go
func TestConfig(t *testing.T) {
    t.Setenv("APP_ENV", "test")
    // automatically restored after test completes
}
```
**A:** `t.Setenv` sets an environment variable and automatically restores the original value when the test ends. It also marks the test as non-parallel (modifying env is not safe concurrently). Prefer over manual `os.Setenv` + `defer os.Unsetenv`.

---

### 55. testify/assert vs Standard Library
**Q: Which produces clearer output?**
```go
// stdlib
if got != want {
    t.Errorf("got %v, want %v", got, want)
}

// testify
assert.Equal(t, want, got)
```
**A:** Both work. `testify/assert` provides richer diff output and many convenience methods. The stdlib approach requires no external dependency — a common requirement in some organisations. Note the argument order for `assert.Equal`: **want first, got second**.

---

### 56. Testing an Error — errors.Is in Tests
**Q: What is the output?**
```go
package main
import (
    "errors"
    "testing"
)

var ErrNotFound = errors.New("not found")

func find(id int) error {
    return fmt.Errorf("db: %w", ErrNotFound)
}

func TestFind(t *testing.T) {
    err := find(0)
    if !errors.Is(err, ErrNotFound) {
        t.Errorf("expected ErrNotFound, got %v", err)
    }
}
```
**A:** Test passes. Use `errors.Is` in tests to check wrapped errors — not `err == ErrNotFound`.

---

### 57. Skipping a Test
**Q: What does t.Skip do?**
```go
func TestIntegration(t *testing.T) {
    if os.Getenv("INTEGRATION") == "" {
        t.Skip("skipping integration test; set INTEGRATION=1 to run")
    }
    // ...
}
```
**A:** `t.Skip` marks the test as skipped (not failed) and stops its execution. Useful for tests that depend on external services or flags. Output: `--- SKIP: TestIntegration`.

---

### 58. Subtests Sharing State — Use t.Cleanup Not defer
**Q: What is the bug?**
```go
func TestSuite(t *testing.T) {
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T) {
            db := openDB()
            defer db.Close() // runs when the outer loop iteration ends, NOT after subtest
            // ...
        })
    }
}
```
**A:** `defer` in the outer function body runs when the **enclosing function** returns. Use `t.Cleanup(db.Close)` inside the subtest to close `db` when that specific subtest finishes.

---

### 59. go test -count=N
**Q: What does this command do?**
```
go test -count=3 ./...
```
**A:** Runs every test 3 times. Useful for detecting flaky tests (tests that pass sometimes and fail others, often due to race conditions or timing).

---

### 60. go test -race
**Q: What does this detect?**
```
go test -race ./...
```
**A:** Enables the Go data race detector. It instruments all memory accesses and reports concurrent read/write conflicts at runtime. Every Go project should run this in CI. Small overhead (~2–20×) but extremely high value.

---

## Section 5: reflect Basics (Q61–Q72)

### 61. reflect.TypeOf and reflect.ValueOf
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    x := 42
    fmt.Println(reflect.TypeOf(x))
    fmt.Println(reflect.ValueOf(x))
}
```
**A:**
```
int
42
```

---

### 62. reflect.Kind vs reflect.Type
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

type MyInt int

func main() {
    var x MyInt = 5
    t := reflect.TypeOf(x)
    fmt.Println(t)
    fmt.Println(t.Kind())
}
```
**A:**
```
main.MyInt
int
```
`Type` is the specific named type; `Kind` is the underlying primitive category.

---

### 63. reflect.Value — CanSet
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    x := 42
    v := reflect.ValueOf(x)
    fmt.Println(v.CanSet())

    v2 := reflect.ValueOf(&x).Elem()
    fmt.Println(v2.CanSet())
    v2.SetInt(100)
    fmt.Println(x)
}
```
**A:**
```
false
true
100
```
You can only set a value through a pointer. `reflect.ValueOf(&x).Elem()` gives an addressable, settable value.

---

### 64. Iterating Struct Fields with reflect
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    t := reflect.TypeOf(p)
    v := reflect.ValueOf(p)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        fmt.Printf("%s: %v\n", field.Name, value)
    }
}
```
**A:**
```
Name: Alice
Age: 30
```

---

### 65. Reading Struct Tags via reflect
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

type User struct {
    Name string `json:"name" validate:"required"`
}

func main() {
    t := reflect.TypeOf(User{})
    f, _ := t.FieldByName("Name")
    fmt.Println(f.Tag.Get("json"))
    fmt.Println(f.Tag.Get("validate"))
}
```
**A:**
```
name
required
```
`StructTag.Get(key)` retrieves the value for a specific tag key.

---

### 66. reflect.DeepEqual
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    a := []int{1, 2, 3}
    b := []int{1, 2, 3}
    fmt.Println(a == nil)              // can't use == on slices
    fmt.Println(reflect.DeepEqual(a, b))

    m1 := map[string]int{"x": 1}
    m2 := map[string]int{"x": 1}
    fmt.Println(reflect.DeepEqual(m1, m2))
}
```
**A:**
```
false
true
true
```
`reflect.DeepEqual` performs deep recursive equality — essential for comparing slices, maps, and nested structs in tests.

---

### 67. reflect.TypeOf on Interface
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

func printType(i interface{}) {
    fmt.Println(reflect.TypeOf(i))
}

func main() {
    printType(42)
    printType("hello")
    printType([]int{1, 2})
}
```
**A:**
```
int
string
[]int
```
`reflect.TypeOf` extracts the dynamic type of an interface value.

---

### 68. reflect.Slice — Creating and Appending
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    sliceType := reflect.TypeOf([]int{})
    s := reflect.MakeSlice(sliceType, 0, 3)
    s = reflect.Append(s, reflect.ValueOf(1))
    s = reflect.Append(s, reflect.ValueOf(2))
    fmt.Println(s.Interface())
}
```
**A:** `[1 2]`. `reflect.MakeSlice` and `reflect.Append` allow slice manipulation without knowing the concrete type at compile time.

---

### 69. reflect.TypeOf nil — Panics
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    var p *int
    fmt.Println(reflect.TypeOf(p))
    fmt.Println(reflect.TypeOf(nil))
}
```
**A:**
```
*int
recovered: reflect: call of reflect.TypeOf on zero Value
```
`reflect.TypeOf(*int nil)` works fine (returns `*int`). But `reflect.TypeOf(nil)` panics — a nil interface has no type.

---

### 70. reflect to Call a Method Dynamically
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

type Greeter struct{}
func (g Greeter) Greet(name string) string { return "Hello, " + name }

func main() {
    g := Greeter{}
    v := reflect.ValueOf(g)
    method := v.MethodByName("Greet")
    result := method.Call([]reflect.Value{reflect.ValueOf("Go")})
    fmt.Println(result[0].String())
}
```
**A:** `Hello, Go`. `MethodByName` finds a method by name; `Call` invokes it with `reflect.Value` arguments.

---

### 71. reflect.Indirect — Dereferencing a Pointer
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    x := 42
    p := &x
    v := reflect.ValueOf(p)
    fmt.Println(v.Kind())
    fmt.Println(reflect.Indirect(v).Kind())
    fmt.Println(reflect.Indirect(v).Int())
}
```
**A:**
```
ptr
int
42
```
`reflect.Indirect(v)` dereferences a pointer value. If `v` is not a pointer, it returns `v` unchanged.

---

### 72. When NOT to Use reflect
**Q: Which is preferred?**
```go
// With reflect — flexible but slow, no compile-time safety
func setField(obj interface{}, name string, value interface{}) { ... }

// With generics (Go 1.18+) — fast, type-safe
func SetField[T any](obj *T, fn func(*T)) { fn(obj) }
```
**A:** Avoid `reflect` when generics or interfaces solve the problem — reflection is ~10–100× slower than direct calls, has no compile-time type safety, and is harder to read. Use reflect only for truly dynamic scenarios: JSON serialisation, ORMs, dependency injection frameworks, testing utilities like `DeepEqual`.

---
