# 🔭 08 — Advanced Topics in Go
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Generics (type parameters, constraints)
- Reflection (`reflect` package)
- `unsafe` package
- Fuzz testing
- WebAssembly (WASM)
- Go 1.18–1.22 new features

---

## ❓ Most Asked Questions

### Q1. How do Go Generics work? (Go 1.18+)

```go
// Generic function with type parameter
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice { result[i] = fn(v) }
    return result
}

nums := []int{1, 2, 3, 4, 5}
doubled := Map(nums, func(n int) int { return n * 2 })       // [2 4 6 8 10]
strs   := Map(nums, func(n int) string { return strconv.Itoa(n) }) // ["1" "2" ...]

// Generic function with constraints
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b { return a }
    return b
}
Min(3, 5)        // 3 (int)
Min(3.14, 2.71)  // 2.71 (float64)
Min("a", "b")    // "a" (string)

// Custom constraint
type Number interface {
    ~int | ~int64 | ~float64 | ~float32
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums { total += n }
    return total
}
Sum([]int{1, 2, 3})         // 6
Sum([]float64{1.1, 2.2})    // 3.3
```

---

### Q2. Implement a generic data structure (e.g., ordered map)

```go
// Generic Stack
type Stack[T any] struct{ items []T }
func (s *Stack[T]) Push(v T)        { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool)  {
    var zero T
    if len(s.items) == 0 { return zero, false }
    n := len(s.items)
    v := s.items[n-1]
    s.items = s.items[:n-1]
    return v, true
}

intStack := &Stack[int]{}; intStack.Push(1)
strStack := &Stack[string]{}; strStack.Push("hello")

// Generic Result type (like Result<T, E> in Rust)
type Result[T any] struct {
    value T
    err   error
}

func Ok[T any](v T) Result[T]    { return Result[T]{value: v} }
func Err[T any](e error) Result[T] { return Result[T]{err: e} }
func (r Result[T]) Unwrap() (T, error) { return r.value, r.err }
```

---

### Q3. How does reflection work in Go?

```go
import "reflect"

// Get type and value info at runtime
x := 42
t := reflect.TypeOf(x)   // int
v := reflect.ValueOf(x)  // 42

// Get struct fields dynamically
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
u := User{ID: 1, Name: "Alice"}
rt := reflect.TypeOf(u)
rv := reflect.ValueOf(u)

for i := 0; i < rt.NumField(); i++ {
    field := rt.Field(i)
    value := rv.Field(i)
    tag   := field.Tag.Get("json")
    fmt.Printf("Field: %s, Tag: %s, Value: %v\n", field.Name, tag, value)
}

// Modify struct via reflection (must use pointer)
rv2 := reflect.ValueOf(&u).Elem()
rv2.FieldByName("Name").SetString("Bob")
fmt.Println(u.Name)  // Bob

// Call methods dynamically
m := rv2.MethodByName("SomeMethod")
if m.IsValid() { m.Call(nil) }
```

---

### Q4. What is the `unsafe` package and when should you use it?

```go
import "unsafe"

// unsafe.Sizeof — size of type in bytes (no allocation)
fmt.Println(unsafe.Sizeof(int64(0)))    // 8
fmt.Println(unsafe.Sizeof(struct{}{}))  // 0

// unsafe.Pointer — bypass Go's type system (use with extreme care)
x := int64(42)
p := unsafe.Pointer(&x)
// Convert to *float64 to reinterpret bits (dangerous!)
f := *(*float64)(p)

// Common legitimate use: memory layout optimization
type Aligned struct {
    a int8   // 1 byte
    b int64  // 8 bytes
}
type Optimized struct {
    b int64  // 8 bytes
    a int8   // 1 byte, now at end — smaller due to alignment
}
fmt.Println(unsafe.Sizeof(Aligned{}))   // 16 (padding added)
fmt.Println(unsafe.Sizeof(Optimized{})) // 16 still, but try int32+int8+int8

// Legitimate use: interacting with C code via cgo
// Or zero-copy conversions between []byte and string
func bytesToStr(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))  // no allocation!
}
```

---

### Q5. What is fuzz testing in Go? (Go 1.18+)

```go
// Fuzz test: Go generates random inputs to find crashes/panics
func FuzzParseDate(f *testing.F) {
    // Seed corpus (starting inputs)
    f.Add("2024-01-15")
    f.Add("invalid-date")
    f.Add("")

    // Fuzz target — called with generated inputs
    f.Fuzz(func(t *testing.T, input string) {
        // Should NOT panic on ANY input
        _, err := ParseDate(input)
        // We don't care if it errors, just that it doesn't PANIC
        _ = err
    })
}
```

```bash
# Run fuzz test (generates random inputs)
go test -fuzz=FuzzParseDate -fuzztime=30s ./...

# Run just the seed corpus (regular test mode)
go test -run=FuzzParseDate ./...
```

---

### Q6. What are the important Go 1.18–1.22 features?

| Version | Feature | Description |
|---------|---------|-------------|
| Go 1.18 | **Generics** | Type parameters, constraints |
| Go 1.18 | **Fuzz testing** | `go test -fuzz` built-in |
| Go 1.18 | `any` keyword | Alias for `interface{}` |
| Go 1.19 | `GOMEMLIMIT` | Soft memory limit for GC |
| Go 1.20 | `errors.Join` | Combine multiple errors |
| Go 1.20 | Comparable generics | `comparable` constraint improvements |
| Go 1.21 | `log/slog` | Structured logging built-in |
| Go 1.21 | `slices`, `maps` packages | Generic slice/map utilities |
| Go 1.21 | `min`, `max` builtins | Built-in min/max functions |
| Go 1.22 | `range` over integers | `for i := range 10 { }` |
| Go 1.22 | Improved routing | `net/http` method/path routing |

```go
// Go 1.21: structured logging
import "log/slog"
slog.Info("server started", "port", 8080, "env", "production")

// Go 1.21: slice utilities
import "slices"
slices.Sort(nums)
slices.Contains(strs, "go")

// Go 1.22: range over int
for i := range 5 { fmt.Println(i) }  // 0 1 2 3 4

// Go 1.22: method routing
mux := http.NewServeMux()
mux.HandleFunc("GET /users/{id}", getUserHandler)
mux.HandleFunc("POST /users", createUserHandler)
```

---

### Q7. How do you compile Go to WebAssembly?

```go
// main.go — compiled to WASM
//go:build js,wasm

package main

import (
    "fmt"
    "syscall/js"
)

func greet(this js.Value, args []js.Value) interface{} {
    name := args[0].String()
    return fmt.Sprintf("Hello, %s from Go WASM!", name)
}

func main() {
    // Expose Go function to JavaScript
    js.Global().Set("goGreet", js.FuncOf(greet))
    // Keep running
    select {}
}
```

```bash
# Compile to WASM
GOOS=js GOARCH=wasm go build -o main.wasm .

# Copy runtime support file
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
```

```html
<!-- Load in browser -->
<script src="wasm_exec.js"></script>
<script>
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
    .then(result => go.run(result.instance));
</script>
```
