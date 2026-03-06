# 🗣️ Theory — Error Handling in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does Go handle errors? Why doesn't it use exceptions?"

> *"Go treats errors as values, not as exceptional events. Instead of throwing and catching exceptions, you return an error as one of the return values — typically the last one. So your function signature looks like `func doSomething() (Result, error)`, and the caller immediately checks `if err != nil`. The Go designers chose this because exceptions add invisible control flow — you can't look at a function call and know whether it might throw. With Go's approach, every possible failure point is explicit in the code. It's more verbose, but it's also more honest and predictable."*

---

## Q: "How do you create errors in Go? What are the different approaches?"

> *"There are three main ways. The simplest is `errors.New('something went wrong')` — good when you just need a static error message. Then there's `fmt.Errorf('user %d not found, %w', id, err)` — use this when you want to format a dynamic message or, importantly, wrap another error using the `%w` verb. And the third approach is creating a custom struct that implements the `error` interface — the interface only requires one method: `Error() string`. Custom error types let you attach extra data, like an HTTP status code, a field name, or an error category, that the caller can inspect."*

---

## Q: "What is error wrapping? What is the `%w` verb?"

> *"Error wrapping lets you add context to an error without losing the original cause. Using `fmt.Errorf('database layer: %w', originalErr)`, you create a new error that contains the original error wrapped inside it. You can chain these wrappings through multiple layers. Then to check the chain, you use `errors.Is(err, targetError)` — which checks if any error in the wrapped chain matches your target. Or `errors.As(err, &myErr)` — which unwraps the chain and tries to extract an error of a specific type. Before wrapping was added in Go 1.13, people would lose the original error. Now you can layer context all the way up and still introspect the root cause."*

---

## Q: "What are sentinel errors?"

> *"Sentinel errors are package-level errors you define once and compare against — like `var ErrNotFound = errors.New('not found')`. The name 'sentinel' means they're special signal values. They're common in the standard library — `io.EOF` is the most famous one. You check them with `errors.Is(err, io.EOF)`. The advantage is that you can match on the exact error identity rather than string-matching. The downside is they become part of your package's public API — once you export a sentinel error, changing it is a breaking change. For richer errors with attached data, use a custom error type instead."*

---

## Q: "What is panic and recover? When should you use them?"

> *"Panic is Go's version of an unrecoverable error — it stops the normal execution of a goroutine, unwinds the call stack, and runs all deferred functions. If not recovered from, it crashes the program with a stack trace. Recover is a function you can call inside a deferred function to catch a panic and resume normal execution. The important rule is: use panic only for programming errors — like invalid arguments or impossible states — never for expected runtime errors like user input failures or network issues. Those should be returned as errors. In production web servers, frameworks often use recover in middleware to convert panics into 500 errors instead of crashing the whole server."*

---

## Q: "What is `errors.Join`? (Go 1.20+)"

> *"Before Go 1.20, if you had multiple errors from concurrent operations, you had to pick one, concatenate their strings, or use a third-party library. `errors.Join` lets you combine multiple errors into a single error value in a standard way. The resulting error's `Error()` string shows all the messages, and importantly, `errors.Is` and `errors.As` work through a joined error — they'll check all the wrapped errors. It's particularly useful with things like form validation where you want to report all validation failures at once."*

---

## Q: "What is the best practice for error handling in Go?"

> *"There are a few key patterns. First, always handle your errors — don't assign to blank identifier and hope for the best. Second, add context when propagating: use `fmt.Errorf('operation context: %w', err)` so errors have meaningful context by the time they reach the top. Third, return early on error — don't nest deep if-else trees, check the error at each step and return it. Fourth, use `errors.Is` and `errors.As` for comparison, never compare error strings. And fifth, handle errors at the right level — don't log and return, choose one or the other. Either log and stop, or return for the caller to handle."*
