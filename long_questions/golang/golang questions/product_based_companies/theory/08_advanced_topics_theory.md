# 🗣️ Theory — Advanced Topics in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are generics in Go? Why did it take so long to add them?"

> *"Generics — officially called type parameters — were added in Go 1.18 and let you write functions and types that work with any type that satisfies a constraint, without repeating code. Before generics, you'd either write separate versions for each type, or use `interface{}` and lose type safety. The reason it took until 2022 — Go is 13 years old — is that the team was very careful about the design. They wanted generics that didn't complicate the language, didn't hurt compile times, and didn't sacrifice readability. The result is cleaner than Java or C++ generics: `func Map[T, U any](s []T, fn func(T) U) []U`. The constraint system — using interfaces as type bounds — is elegant and reuses existing interface concepts."*

---

## Q: "What are type constraints in generics?"

> *"A type constraint is an interface that limits what types can be used as a type parameter. The simplest constraint is `any` — which allows any type. The standard library provides `comparable` — which allows types that can be compared with `==`. For arithmetic, you'd use `constraints.Ordered` from `golang.org/x/exp/constraints` which allows types that support `<`, `>`, etc. You can write custom constraints using the union element syntax: `type Number interface { ~int | ~int64 | ~float64 }`. The `~T` syntax means 'any type whose underlying type is T' — so a custom type `type Celsius float64` would satisfy `~float64`."*

---

## Q: "What is the `reflect` package and when do you actually use it?"

> *"Reflection lets you inspect and manipulate types and values at runtime — when you don't know the concrete type at compile time. `reflect.TypeOf(x)` gives you type metadata, `reflect.ValueOf(x)` gives you the actual value. You can iterate over struct fields, call methods by name, and set field values. In practice, reflection is used by marshaling libraries — the `encoding/json` package uses it to serialize any struct. ORM libraries use it to map struct fields to database columns. Dependency injection frameworks use it to wire up types. The rule of thumb: don't use reflection in your own business logic. It's slow, bypasses type safety, and is hard to understand. It's the right tool for framework and library authors."*

---

## Q: "What is the `unsafe` package? When is it acceptable to use?"

> *"The `unsafe` package bypasses Go's type safety. Key functions: `unsafe.Sizeof(T)` returns the size of a type in bytes — no allocation, purely a compile-time constant. `unsafe.Pointer` is a special pointer type that can be converted between any pointer type — you can reinterpret the same memory as a different type. This is genuinely unsafe — the compiler can't check correctness. Legitimate uses are narrow: zero-copy string-to-bytes conversion, struct padding analysis for memory layout optimization, and interoperability with C via cgo. Libraries like `encoding/json` use it for performance. In application code, you almost never need it. If you think you do, question whether there's a safe alternative first."*

---

## Q: "What is fuzz testing? How is it different from unit testing?"

> *"Fuzz testing — added as a built-in Go feature in 1.18 — is a technique where the testing framework generates random inputs to find crashes, panics, or unexpected behavior. You define a `FuzzXxx` function with a fuzz target — a function that accepts generated values. The fuzzer starts from a seed corpus you provide, then mutates inputs to find edge cases. It's fundamentally different from unit testing: unit tests check specific known cases you thought of. Fuzz testing finds cases you didn't think of — like a parser that panics on a specific byte sequence, or a codec that crashes on certain input sizes. It's most valuable for parsing, encoding, and security-sensitive code."*

---

## Q: "What are the most important features added in recent Go versions?"

> *"I'd categorize the biggest additions: Go 1.18 brought generics and native fuzz testing — the two most significant additions since Go 1.0. Go 1.19 added `GOMEMLIMIT`, a soft memory ceiling that makes GC tuning much easier. Go 1.20 added `errors.Join` for combining multiple errors. Go 1.21 added `log/slog` for structured logging as a standard — now you don't need Zap or Logrus for structured logs. It also added the `slices` and `maps` packages for common generic operations. Built-in `min` and `max` functions were added. Go 1.22 brought range-over-integer — `for i := range 10 {}` — which removes the need for `for i := 0; i < 10; i++`. And improved HTTP routing in `net/http` with method and wildcard matching."*

---

## Q: "What is WebAssembly (WASM) and how does Go compile to it?"

> *"WebAssembly is a binary instruction format that runs in a sandboxed virtual machine — both in browsers and in server environments like WASI. Go can compile to WASM with `GOOS=js GOARCH=wasm go build`. The result is a `.wasm` file that you load in a browser with the Go-provided `wasm_exec.js` runtime. Go functions can be exposed to JavaScript using `js.FuncOf` from `syscall/js`. The use cases: running Go business logic in the browser without a server round-trip, running Go code in edge environments. The major limitation is binary size — a 'Hello World' Go WASM is several megabytes. TinyGo — a Go subset compiler — addresses this and produces much smaller binaries at the cost of not supporting some standard library packages."*
