# Golang Interfaces Cheatsheet

Interfaces in Go are unique (implicit implementation). This cheatsheet covers the mechanics and the most famous "gotchas".

## 🟢 Basics

### 1. Implicit Implementation
No `implements` keyword. A type implements an interface by implementing its methods.

```go
type Stringer interface {
    String() string
}

type User struct {
    Name string
}

// User now implements Stringer automatically
func (u User) String() string {
    return u.Name
}
```

### 2. The Empty Interface (`interface{}` or `any`)
Holds values of **any type**.
- `interface{}` says nothing about the behavior of data.
- Go 1.18+ introduced `any` as an alias for `interface{}`.

```go
func PrintAny(v any) {
    fmt.Println(v)
}
```

---

## 🟡 Type Mechanics

### 1. Type Assertion
Extract the underlying concrete value.
**Syntax:** `value, ok := interfaceVar.(ConcreteType)`

```go
var i any = "hello"

// Unsafe (Panics if wrong type)
s := i.(string)

// Safe (Check boolean)
if s, ok := i.(string); ok {
    fmt.Println(s)
} else {
    fmt.Println("Not a string")
}
```

### 2. Type Switch
Handle multiple types.
```go
func do(i any) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %q\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

---

## 🔴 The "Billion Dollar Mistake" Trap

### ☠️ The "Nil Interface" vs "Interface Containing Nil"
An interface is a tuple `(type, value)`. It is `nil` ONLY if **both** are nil.

**The Trap:**
If you store a `nil` pointer inside an interface, the interface is **NOT nil**.

```go
type MyError struct{}
func (e *MyError) Error() string { return "oops" }

func fail() error {
    var err *MyError = nil
    return err // Returns an interface containing (MyError, nil) -> NOT NIL!
}

func main() {
    err := fail()
    if err != nil {
        fmt.Println("Error occurred!") // This PRINTS! Unexpectedly.
        // attempting to call err.Error() might panic or work depending on implementation
    }
}
```

**The Fix:**
Always return explicit `nil` for error interfaces, not a nil concrete pointer.
```go
func fail() error {
    return nil // Explicit nil interface
}
```

---

## 🔵 Advanced Concepts

### 1. Interface Embedding
Combine interfaces.
```go
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }

type ReadWriter interface {
    Reader
    Writer
}
```

### 2. Compile-Time Check
Verify a type implements an interface.
```go
// Fails to compile if *MyType doesn't implement MyInterface
var _ MyInterface = (*MyType)(nil)
```

### 3. Best Practice: "Accept Interfaces, Return Structs"
- **Accept Interfaces:** Makes your function flexible (easy to mock).
- **Return Structs:** forcing return of interfaces implies pre-emptive abstraction, which can be rigid.

---

## 🟣 Common Standard Library Interfaces

| Interface | Method | Used For |
| :--- | :--- | :--- |
| `fmt.Stringer` | `String() string` | Custom print formatting |
| `error` | `Error() string` | Error handling |
| `io.Reader` | `Read(p []byte) (n int, err error)` | Stream reading |
| `io.Writer` | `Write(p []byte) (n int, err error)` | Stream writing |
| `sort.Interface` | `Len()`, `Less(i, j)`, `Swap(i, j)` | Custom sorting |
