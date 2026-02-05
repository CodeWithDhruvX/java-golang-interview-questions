# Golang Comprehensive Cheatsheet

A quick reference guide for Go syntax, data structures, and common patterns.

---

## 🟢 Basics

### Variables & Constants
```go
// Method 1: Var keyword
var x int = 10
var y = 20        // Type inferred
var z int         // Zero value (0)

// Method 2: Short declaration (func only)
name := "Golang"
isActive := true

// Constants
const Pi = 3.14159
const (
    StatusOpened = 0
    StatusClosed = 1
)
```

### Basic Types
- `bool`: true, false
- `string`: "hello"
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `byte` (alias for uint8)
- `float32`, `float64`
- `complex64`, `complex128`

### Type Conversion
```go
i := 42
f := float64(i)
u := uint(f)
```

---

## 🟡 Data Structures

### Arrays
Fixed size sequence of elements.
```go
// Define
var arr [5]int                  // [0, 0, 0, 0, 0]
nums := [3]int{1, 2, 3}         // initialized
mixed := [...]int{4, 5, 6}      // size inferred

// Access & Modify
fmt.Println(nums[0])            // Read
nums[1] = 10                    // Write

// Iterate
for i, v := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, v)
}
```

### Slices
Dynamic view of an array. Most commonly used.
```go
// Define
var s []int                     // nil slice
s1 := []int{1, 2, 3}            // literal
s2 := make([]int, 5)            // len=5, cap=5, zero-filled
s3 := make([]int, 3, 10)        // len=3, cap=10

// Operations
s1 = append(s1, 4)              // Add element
s1 = append(s1, 5, 6)           // Add multiple
sub := s1[1:3]                  // Slicing (index 1 to 2)

// Copy
dest := make([]int, len(s1))
copy(dest, s1)
```

### Maps
Key-value pairs (unordered).
```go
// Define
m := make(map[string]int)       // Empty map
scores := map[string]int{       // Initialized
    "Alice": 90,
    "Bob":   85,
}

// Operations
scores["Charlie"] = 95          // Add/Update
val := scores["Alice"]          // Retrieve
delete(scores, "Bob")           // Delete

// Check existence
val, ok := scores["Dave"]
if ok {
    fmt.Println("Exists:", val)
} else {
    fmt.Println("Does not exist")
}

// Iterate
for key, value := range scores {
    fmt.Printf("%s: %d\n", key, value)
}
```

### Structs
Custom data types.
```go
// Define
type User struct {
    Name  string
    Email string
    Age   int
}

// Initialize
u1 := User{Name: "John", Age: 30}
u2 := User{"Jane", "jane@test.com", 25} // Order matters

// Access
fmt.Println(u1.Name)
u1.Age = 31

// Anonymous Struct
config := struct {
    Port int
    Env  string
}{
    Port: 8080,
    Env:  "dev",
}
```

---

## 🔵 Control Structures

### If / Else
```go
if x > 10 {
    fmt.Println("Big")
} else if x == 10 {
    fmt.Println("Ten")
} else {
    fmt.Println("Small")
}

// With initialization
if err := process(); err != nil {
    return err
}
```

### For Loop
Go only has `for`.
```go
// Standard C-style
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-style
count := 0
for count < 5 {
    count++
}

// Infinite
for {
    // loop forever
    break // exit
}

// Range (Slices, Maps, Channels)
for i, v := range []int{10, 20, 30} {
    fmt.Println(i, v)
}
```

### Switch
```go
// Standard
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X")
case "linux":
    fmt.Println("Linux")
default:
    fmt.Println("Other")
}

// No condition (Tagless)
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Morning")
case t.Hour() < 17:
    fmt.Println("Afternoon")
default:
    fmt.Println("Evening")
}

// Type Switch
func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}
```

---

## 🟣 Functions & Methods

### Functions
```go
// Basic
func add(a int, b int) int {
    return a + b
}

// Multiple Returns
func swap(x, y string) (string, string) {
    return y, x
}

// Named Returns
func div(x, y int) (res int, rem int) {
    res = x / y
    rem = x % y
    return // Returns res, rem
}

// Variadic
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

### Methods
Functions attached to types.
```go
type Rect struct {
    Width, Height int
}

// Value Receiver (Does not modify original)
func (r Rect) Area() int {
    return r.Width * r.Height
}

// Pointer Receiver (Modifies original)
func (r *Rect) Scale(f int) {
    r.Width = r.Width * f
    r.Height = r.Height * f
}
```

---

## 🟠 Interfaces

### Definition & Implementation
Implicit implementation (duck typing).
```go
// Interface
type Shaper interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

// Implementation
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    var s Shaper
    s = Circle{Radius: 5}
    fmt.Println(s.Area())
}
```

### Empty Interface
Holds values of any type.
```go
var i interface{}
i = "hello"
i = 42

// Type Assertion
s, ok := i.(string) // Check if string
```

---

## 🔴 Concurrency

### Goroutines
Lightweight threads.
```go
go func() {
    fmt.Println("I'm running concurrently!")
}()
```

### Channels
Typed conduits for communication.
```go
// Unbuffered (Blocks until send/receive)
ch := make(chan int)

// Buffered (Non-blocking until full)
bufCh := make(chan int, 2)

// Send
ch <- 42

// Receive
val := <-ch

// Close
close(ch)
```

### Select
Wait on multiple channel operations.
```go
select {
case msg1 := <-c1:
    fmt.Println("Received from c1:", msg1)
case msg2 := <-c2:
    fmt.Println("Received from c2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No activity")
}
```

### WaitGroup
Wait for goroutines to finish.
```go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("Worker", id)
    }(i)
}

wg.Wait()
```

---

## 🟤 Error Handling

### Basics
```go
func method() (int, error) {
    return 0, errors.New("something went wrong")
}

if val, err := method(); err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Success:", val)
}
```

### Custom Errors
```go
type MyError struct {
    Code int
    Msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%d - %s", e.Code, e.Msg)
}
```

### Panic & Recover
Use sparingly for unrecoverable errors.
```go
func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("boom")
}
```

---

## ⚡ Pointers

```go
x := 10
p := &x         // Address of x
fmt.Println(*p) // Dereference (read value: 10)
*p = 20         // Change value at address
```
