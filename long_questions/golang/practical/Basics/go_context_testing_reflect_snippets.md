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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between Background and TODO?
**Your Response:** Both return non-nil contexts that are never cancelled, so both print `true`. The difference is semantic: `context.Background()` is the root context you use in production code as the starting point. `context.TODO()` is a placeholder you use when you're refactoring code to add context support but haven't propagated it yet. It signals "I intend to add a real context here later."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does WithCancel work?
**Your Response:** This prints both messages (order may vary). `context.WithCancel` creates a cancellable context and a cancel function. When we call `cancel()`, it closes the `ctx.Done()` channel. The goroutine waiting on `ctx.Done()` wakes up and `ctx.Err()` returns `context.Canceled`. This is the pattern for cancelling operations in Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `true` then `false`. `context.Background()` returns a context that can never be cancelled, so its `Done()` method returns `nil`. When you create a cancellable context with `context.WithCancel`, it returns a context with a real channel that will be closed when cancelled, so `Done()` is not `nil`.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `timed out: context deadline exceeded`. `context.WithTimeout` creates a context that automatically cancels after 50ms. The select statement waits for either the 200ms timer or the context to be cancelled. Since the timeout happens first (50ms < 200ms), the context case wins and we get the timeout error.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `context deadline exceeded` then `true`. `context.WithDeadline` creates a context that cancels at a specific time. We set the deadline 30ms in the future, then wait for `ctx.Done()`. When the deadline passes, the context is cancelled and `ctx.Err()` returns `context.DeadlineExceeded`. The comparison with `context.DeadlineExceeded` returns `true`.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the bug here?
**Your Response:** This is a resource leak. When you call `context.WithCancel`, it creates resources (like a timer and channel) that must be cleaned up by calling the cancel function. By discarding the cancel function with `_`, we never clean up these resources. Even if the context times out naturally, you should still call `defer cancel()` to ensure immediate cleanup. Tools like `go vet` will flag this pattern.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what's the convention?
**Your Response:** Yes, this compiles and prints `user-42 <nil>`. The convention in Go is that `context.Context` should always be the first parameter in function signatures, typically named `ctx`. It should flow through your call chain like a parameter, never be stored in a struct field. This makes the context's lifetime and propagation clear and explicit.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `101` then `<nil>`. `context.WithValue` stores a key-value pair in the context chain. When we retrieve with `ctx.Value(key("userID"))`, we get `101`. When we try to retrieve a key that doesn't exist (`key("other")`), it returns `nil` rather than panicking. This is how you check if a value exists in the context.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why is this the correct pattern?
**Your Response:** Using a plain string key for context values is dangerous because any package can use the same string key and accidentally overwrite your value. The best practice is to define an unexported key type (like `type ctxKey string`) in your package. Since the type is unexported, other packages can't create keys of that type, preventing collisions. This makes your context values safe from accidental conflicts.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** All 3 workers print "cancelled" (order may vary). This shows how cancellation propagates through goroutines. When we call `cancel()`, it closes the `ctx.Done()` channel. All goroutines waiting on that channel wake up simultaneously. This is the idiomatic way to coordinate cancellation across multiple goroutines - one cancel call can signal many workers to stop.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `child done: context canceled`. When you cancel a parent context, it automatically cancels all its child contexts. This creates a cancellation tree - cancelling the root propagates down through all derived contexts. This is useful for cancelling an entire operation hierarchy with a single call.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `parent still active`. Cancellation only flows downward in Go - when you cancel a child context, it doesn't affect its parent. This makes sense because the parent might have other children that need to continue running. The cancellation hierarchy is one-directional.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the best practice here?
**Your Response:** Even with `context.WithTimeout`, you should always call `defer cancel()`. If your operation finishes before the timeout expires, calling cancel() immediately releases the timer resources instead of waiting for the deadline. This prevents resource leaks and is the idiomatic pattern.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `<nil>` then `context canceled`. Before cancellation, `ctx.Err()` returns `nil` indicating the context is still active. After calling `cancel()`, `ctx.Err()` returns the `context.Canceled` error. This is how you check if a context has been cancelled without blocking on it.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this do?
**Your Response:** This creates an HTTP request that will automatically be cancelled if it takes longer than 3 seconds. `http.NewRequestWithContext` attaches the context to the request, so when the context times out, the underlying HTTP connection is closed. This is the modern way to add timeouts to HTTP requests in Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `<nil>` then `true`. When you call `ctx.Value()` with a key that doesn't exist in the context chain, it returns `nil`. The comparison with `nil` returns `true`. This is the safe way to check if a context value exists - it never panics, just returns nil for missing keys.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints the zero time with `false`, then `true`. `context.Background()` has no deadline, so `Deadline()` returns the zero time and `false`. A context created with `WithDeadline` has a real deadline, so the second call returns `true`. The boolean tells you whether the context actually has a deadline.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** This is the non-blocking context check pattern. The `select` with a `default` case allows you to check if the context is cancelled without blocking. If the context is cancelled, the first case runs and you return the error. If not, the `default` case runs and you continue processing. This is perfect for long-running loops that need to be cancellable.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `task failed`. The `errgroup` package coordinates multiple goroutines with a shared context. When any goroutine returns an error, the group cancels the context (signalling other goroutines to stop) and `g.Wait()` returns that first error. This is perfect for running multiple operations where you want to stop everything if one fails.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When should you use Background vs TODO?
**Your Response:** Use `context.Background()` at the top level of your application - in main(), server entry points, or tests. Use `context.TODO()` when you're refactoring code to add context support but haven't propagated it through the call chain yet - it signals "I intend to add a real context here later". Never use TODO in production code.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `context canceled` then `disk full`. `context.WithCancelCause` (Go 1.20+) lets you attach a specific reason when cancelling. `ctx.Err()` still returns the standard cancellation error, but `context.Cause(ctx)` gives you the actual reason you passed to `cancel()`. This is useful for providing more detailed error information.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which is the correct pattern?
**Your Response:** Never store context in a struct field. Always pass it explicitly as the first parameter to functions that need it. This makes the context's lifetime clear and explicit, prevents coupling between the context and object lifecycle, and follows Go's official recommendation. Context should flow through your program like a parameter, not be embedded in objects.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This is a classic goroutine leak. The goroutine tries to send `42` on an unbuffered channel, but nobody is reading from it. Since `leaky()` returns immediately, the channel is never read from, so the goroutine blocks forever waiting for a receiver. This leaks a goroutine that will never be garbage collected. You can detect this with `runtime.NumGoroutine()` or the `goleak` package.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this fix the leak?
**Your Response:** Yes, this fixes the leak. By making the channel buffered with capacity 1, the send operation `ch <- 42` doesn't block - it immediately places the value in the buffer and the goroutine exits. When the function returns, the channel goes out of scope and gets garbage collected, so no resources are leaked.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This is another goroutine leak. The `subscribe()` function creates a channel but never sends anything to it or closes it. The goroutine in `main()` tries to read from this channel with `v := <-ch` and blocks forever waiting for a value that will never come. Every receiver needs either a sender or someone to close the channel.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints both messages (order may vary). This shows the idiomatic way to fix goroutine leaks using context. We create a cancellable context and pass it to the worker. The worker runs a loop with a select statement that checks `ctx.Done()`. When we call `cancel()`, it closes the `ctx.Done()` channel, causing the worker to exit cleanly instead of blocking forever.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This leaks both a TCP connection and a goroutine. When you make an HTTP request, Go's HTTP client manages a goroutine that reads the response body. If you don't close `resp.Body`, that goroutine never gets the signal to stop, and the TCP connection stays open. Always call `defer resp.Body.Close()` right after checking for errors to ensure proper cleanup.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this pattern do?
**Your Response:** This is a manual goroutine leak detection pattern for tests. It counts goroutines before and after running the code under test. If the count increases, it means goroutines were leaked. The `runtime.Gosched()` call gives the scheduler a chance to clean up. For more reliable detection, use the `go.uber.org/goleak` package which automatically checks for leaks after each test.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This is a goroutine leak in an HTTP handler. If the HTTP client disconnects (closes the browser, loses connection), the handler returns immediately, but the goroutine doing heavy work is still running and tries to send its result to a channel that nobody is reading from anymore. This goroutine blocks forever. The fix is to pass `r.Context()` to the heavy work function so it gets cancelled when the client disconnects.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This is a goroutine leak caused by not closing a channel. The `range` over a channel in `main()` blocks until the channel is closed, but the goroutine that sends values never calls `close(ch)`. So after sending 5 values, the goroutine exits but the channel remains open, causing `main()` to hang forever waiting for more values. Always close channels when you're done sending to them.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This goroutine leaks because it has no way to exit when there's no more work. The `select` statement only has one case for receiving from `ch`. If the caller stops sending values and never closes the channel, this goroutine will block forever waiting for a value that never comes. The fix is to add a `done` channel parameter and a `case <-done: return` to provide a clean exit path.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this do?
**Your Response:** This sets up automatic goroutine leak detection for an entire test package. The `goleak.VerifyTestMain` wraps your test suite and checks after each test that no goroutines were left running. If it finds any unexpected goroutines, it fails the test suite with a detailed report. This is much more reliable than manual goroutine counting and catches leaks you might otherwise miss.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the complexity issue?
**Your Response:** This is O(n²) time complexity because strings in Go are immutable. Each `+=` operation creates a completely new string by copying all the previous characters plus the new one. For n iterations, you're copying 1+2+3+...+n characters, which is O(n²). This also creates a lot of garbage for the GC. Use `strings.Builder` for O(n) performance.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `01234` then `5`. `strings.Builder` efficiently builds strings by writing to an internal buffer that grows as needed. It doesn't copy the existing content when appending - it just writes the new data. `b.String()` returns the final string without any additional allocation. `b.Len()` gives the current length of the builder.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `hello` then `world`. The `Reset()` method clears the builder's internal buffer, allowing you to reuse the same `strings.Builder` instance. This is useful in loops or when building multiple strings to avoid allocating new builders each time. After `Reset()`, the builder is empty and ready to build a new string.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This panics at runtime because `strings.Builder` has internal copy protection. When you copy a builder (`b2 := b`), both copies point to the same underlying buffer. If you then write to one, it could corrupt the other. To prevent this, `strings.Builder` has a copy check that panics if you try to use a copied builder. Never copy builders - pass them by pointer if you need to share them.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hello, Go!` then `10`. `bytes.Buffer` is similar to `strings.Builder` but works with bytes instead of strings. It accumulates writes in an internal buffer. `String()` converts the buffer contents to a string, and `Len()` returns the number of bytes currently stored. Unlike `strings.Builder`, `bytes.Buffer` can also be used for reading.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the key difference?
**Your Response:** The key difference is their purpose. `strings.Builder` is specifically optimized for building strings - it's write-only until you call `String()`. `bytes.Buffer` is more general-purpose - it implements both `io.Reader` and `io.Writer`, so you can read from it and write to it. Use `strings.Builder` when you're building a string, use `bytes.Buffer` when you need a general in-memory byte buffer.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hi!`. `strings.Builder` provides type-specific methods for writing single characters. `WriteByte` writes a single byte, `WriteRune` writes a Unicode rune (which can be multiple bytes for non-ASCII characters), and `WriteString` writes a full string. These methods are more efficient than the general `Write` method when you know the type.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `name: Alice, age: 30`. Since `bytes.Buffer` implements the `io.Writer` interface, you can use it directly with `fmt.Fprintf`. This is very convenient for building formatted strings in memory. The formatted output goes directly into the buffer instead of stdout, and you can later get the result with `buf.String()`.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hello` then `5`. `strings.NewReader` creates an `io.Reader` from a string. When we call `r.Read(buf)`, it reads up to 5 bytes (the buffer size) from the string into `buf`. The return value `n` tells us how many bytes were actually read (5). `r.Len()` returns the number of unread bytes remaining in the reader (11 - 5 = 6, but the string is only 11 bytes total, so it shows 5 remaining after reading 5).

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. Strings in Go are immutable - you cannot modify individual characters. The compiler gives an error saying strings are not addressable. If you need to modify a string, you must convert it to a `[]byte` slice, make your modifications, and then convert it back to a string. This creates a new string rather than modifying the original.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `hello` then `Hello`. When you convert a string to `[]byte`, Go creates a copy of the string's bytes in a new mutable slice. When we modify `b[0]` to 'H', we're changing the copy, not the original string. The original `s` remains unchanged. This demonstrates string immutability - to "modify" a string, you actually create a new one.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** This pattern pre-allocates memory in `strings.Builder` to avoid repeated reallocations. We first calculate the total size needed by summing the lengths of all parts. Then `b.Grow(total)` tells the builder to allocate enough memory for the final result upfront. This prevents the builder from having to grow its internal buffer multiple times as we append, giving better performance for large strings.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output when go test runs this?
**Your Response:** This prints `ok mymath` indicating all tests passed. This is the table-driven test pattern - we define test cases in a slice of structs, then loop through them running each as a subtest with `t.Run`. Each subtest gets a descriptive name, and you can run individual subtests with `go test -run TestAdd/testname`. This makes tests organized and easy to maintain.

---

### 46. t.Error vs t.Fatal
**Q: What is the difference?**
```go
t.Error("something went wrong") // marks test as failed; continues execution
t.Fatal("critical failure")     // marks test as failed; stops this test immediately
```
**A:** `t.Error` / `t.Errorf` log the failure and continue. `t.Fatal` / `t.Fatalf` call `t.FailNow()` which stops the current test function (via `runtime.Goexit()`). Use `Fatal` when further execution is meaningless after the failure.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference?
**Your Response:** `t.Error` logs the failure but continues running the test - useful when you want to see all failures. `t.Fatal` logs the failure and immediately stops the test - use when the test can't continue after this failure (like when setup fails). `t.Fatal` calls `runtime.Goexit()` which stops just the test goroutine, not the whole program.

---

### 47. t.Run Subtests — Run a Single Subtest
**Q: What command runs only the `negative` subtest?**
```
go test -run TestAdd/negative
```
**A:** Runs only the `negative` subtest. The pattern is `TestFunctionName/subtest_name`. Spaces in subtest names become underscores in the `-run` pattern.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What command runs only the negative subtest?
**Your Response:** You use `go test -run TestAdd/negative`. The pattern is `TestFunctionName/subtestName`. This lets you run specific subtests without running the entire test suite, which is great for debugging failing tests. Note that spaces in subtest names become underscores when using the -run flag.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does t.Parallel() do?
**Your Response:** `t.Parallel()` marks a subtest to run in parallel with other parallel subtests. This can speed up your test suite significantly. The test framework waits for all parallel subtests to complete. Before Go 1.22, you needed `tt := tt` to capture the range variable, but Go 1.22 fixed this issue.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of TestMain?
**Your Response:** `TestMain` is the entry point for all tests in a package. It lets you run global setup before any tests run (like starting a test database) and teardown after all tests complete (like cleaning up resources). You must call `m.Run()` to actually execute the tests, and then use `os.Exit()` with the result to preserve the exit code.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does t.Helper() do?
**Your Response:** `t.Helper()` marks a function as a test helper. When a test fails, it shows the line in the actual test function, not inside the helper. This makes test failures much easier to debug because you see the relevant test code that failed, not some generic helper function. It's a small but very useful feature for writing clean test utilities.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the structure of a benchmark?
**Your Response:** A benchmark function runs the code to be measured in a loop where `b.N` is set by the testing framework. The framework automatically adjusts `b.N` until the measurement is stable (usually running for at least 1 second). Run benchmarks with `go test -bench=BenchmarkName -benchmem` to see both timing and allocation metrics.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** `b.ResetTimer()` excludes setup time from benchmark measurements. When you have expensive setup before the actual benchmark loop, call `b.ResetTimer()` to reset the timing counters. This ensures you're only measuring the performance of the code you care about, not the setup overhead.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `test body` then `cleanup called`. `t.Cleanup` registers cleanup functions that run when the test finishes, even if it fails. This is better than `defer` in table-driven tests because cleanup runs for each subtest individually, not when the outer function returns. Multiple cleanup functions run in reverse order (LIFO).

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the advantage?
**Your Response:** `t.Setenv` is safer than manual environment variable handling. It sets the env var and automatically restores the original value when the test ends, even if the test panics. It also prevents the test from running in parallel since modifying environment variables affects the whole process. This prevents test isolation issues.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which produces clearer output?
**Your Response:** `testify/assert` generally produces clearer output with nice diffs and has many convenience methods. The standard library approach requires no dependencies. Be careful with `assert.Equal` - it uses the opposite argument order than the stdlib pattern: it's `assert.Equal(t, want, got)` not `got, want`. This trips up many developers!

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** The test passes. When using wrapped errors with `fmt.Errorf("db: %w", ErrNotFound)`, you can't compare directly with `==`. Use `errors.Is` to unwrap and check if the error contains or is `ErrNotFound`. This works even through multiple layers of error wrapping, which is the modern way to handle errors in Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does t.Skip do?
**Your Response:** `t.Skip` marks a test as skipped rather than failed. Use this for tests that shouldn't run under certain conditions, like when external services aren't available or when environment variables aren't set. The test output shows `--- SKIP` making it clear the test was intentionally skipped, not failed.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** The bug is that `defer db.Close()` runs when the outer `TestSuite` function returns, not when each individual subtest finishes. This means resources aren't cleaned up between subtests. Use `t.Cleanup(db.Close)` inside each subtest to ensure cleanup happens when that specific subtest completes.

---

### 59. go test -count=N
**Q: What does this command do?**
```
go test -count=3 ./...
```
**A:** Runs every test 3 times. Useful for detecting flaky tests (tests that pass sometimes and fail others, often due to race conditions or timing).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this command do?
**Your Response:** This runs every test 3 times. It's great for detecting flaky tests - tests that sometimes pass and sometimes fail, usually due to race conditions or timing issues. If a test fails on any of the runs, the entire command fails. This is more reliable than running tests manually multiple times.

---

### 60. go test -race
**Q: What does this detect?**
```
go test -race ./...
```
**A:** Enables the Go data race detector. It instruments all memory accesses and reports concurrent read/write conflicts at runtime. Every Go project should run this in CI. Small overhead (~2–20×) but extremely high value.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this detect?
**Your Response:** This enables the Go race detector which finds data races at runtime. It instruments all memory accesses and reports when multiple goroutines access the same memory concurrently without proper synchronization. Every Go project should run this in CI - the performance overhead (2-20x slower) is worth catching race conditions before they hit production.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `int` then `42`. `reflect.TypeOf(x)` returns the type of the variable as a `reflect.Type` object. `reflect.ValueOf(x)` returns the value as a `reflect.Value` object. When printed, `TypeOf` shows the type name and `ValueOf` shows the actual value. These are the two entry points to reflection in Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `main.MyInt` then `int`. `reflect.TypeOf(x)` gives you the exact type, including the package name for named types. `t.Kind()` gives you the underlying primitive kind - in this case, even though `MyInt` is a named type, its kind is still `int`. This distinction is important when you need to know both the specific type and its general category.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `false`, `true`, then `100`. `v.CanSet()` returns `false` because `reflect.ValueOf(x)` creates a copy of `x`, which isn't addressable. But `reflect.ValueOf(&x).Elem()` gives us a reflect.Value that points to the original `x`, so `CanSet()` returns `true`. Then `v2.SetInt(100)` modifies the original variable, which is why `x` becomes 100.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints the field names and values. `reflect.TypeOf(p)` gives us the struct type, and `t.NumField()` tells us how many fields it has. We iterate through each field, getting its name with `t.Field(i).Name` and its value with `v.Field(i)`. This is how you can dynamically inspect and work with struct fields at runtime.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `name` then `required`. `t.FieldByName("Name")` finds the struct field by name, and `f.Tag.Get("json")` retrieves the value for the `json` tag key. This is how libraries like JSON encoding read struct tags to know how to serialize fields. Struct tags are a powerful way to add metadata to struct fields.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `false`, `true`, `true`. You can't use `==` to compare slices in Go, which is why `a == nil` is false (it's checking if the slice itself is nil, not comparing contents). `reflect.DeepEqual` recursively compares values, making it perfect for comparing slices, maps, and nested structs. It's essential in tests for complex data structures.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `int`, `string`, then `[]int`. When you pass values to `printType`, they're boxed into an `interface{}`. `reflect.TypeOf` extracts the actual dynamic type of what's stored in the interface. This is how you can inspect the concrete type of interface values at runtime, which is useful for writing generic code or debugging.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `[1 2]`. `reflect.TypeOf([]int{})` gets the type of an int slice, then `reflect.MakeSlice` creates a new slice of that type. `reflect.Append` appends values to the slice, returning a new slice value. `s.Interface()` converts the reflect.Value back to a regular interface{} which we can print. This is how you manipulate slices when you don't know their type at compile time.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `*int` then recovers from a panic. `reflect.TypeOf(p)` where `p` is a nil pointer returns `*int` - the type is known even though the value is nil. But `reflect.TypeOf(nil)` panics because a nil interface value has no type information. The recover catches the panic and prints the message. This shows the difference between a nil pointer and a nil interface.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hello, Go`. `reflect.ValueOf(g)` gives us a reflect.Value for the struct. `v.MethodByName("Greet")` finds the method by name. `method.Call([]reflect.Value{reflect.ValueOf("Go")})` invokes the method, passing a slice of reflect.Value arguments. The method returns a slice of reflect.Value results, and we get the first result with `result[0].String()`.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `ptr`, `int`, then `42`. `v.Kind()` shows that `v` is a pointer. `reflect.Indirect(v)` dereferences the pointer, giving us the value it points to. The indirect value's kind is `int`, and we can get its integer value with `.Int()`. `reflect.Indirect` is a convenient way to dereference pointers safely - if the value isn't a pointer, it just returns it as-is.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which is preferred?
**Your Response:** The generics approach is strongly preferred. Reflection is 10-100x slower than direct calls, provides no compile-time type safety, and makes code harder to read and maintain. Only use reflection for truly dynamic scenarios where you don't know types at compile time, like JSON serialization, ORMs, or testing utilities. For most cases, Go 1.18+ generics provide a type-safe, performant alternative.

---
