# Go Programming - Interview Questions and Answers

---

## 🟢 Basics (Questions 1-20)

### Question 1: What is Go and who developed it?

**Answer:**
Go is a programming language that helps developers build software. It was created by Google in 2007 by three engineers: Robert Griesemer, Rob Pike, and Ken Thompson. The language was made public in 2009. Go is designed to be simple, fast, and easy to use for building modern applications.

---

### Question 2: What are the key features of Go?

**Answer:**
Go has several important features:
- **Simple and easy to learn** - The language has fewer keywords and is easy to understand
- **Fast compilation** - Your code compiles very quickly into programs
- **Built-in concurrency** - Easy to run multiple tasks at the same time using goroutines
- **Garbage collection** - Automatic memory management, so you don't need to manually free memory
- **Static typing** - Types are checked at compile time, which helps catch errors early
- **Good standard library** - Comes with many useful built-in tools and packages

---

### Question 3: How do you declare a variable in Go?

**Answer:**
There are three main ways to declare variables in Go:

1. **Using var keyword:**
   ```go
   var name string = "John"
   var age int = 25
   ```

2. **Short declaration (inside functions only):**
   ```go
   name := "John"
   age := 25
   ```

3. **Without initial value:**
   ```go
   var count int  // automatically gets zero value (0)
   ```

---

### Question 4: What are the data types in Go?

**Answer:**
Go has several basic data types:

1. **Numbers:**
   - `int`, `int8`, `int16`, `int32`, `int64` - whole numbers
   - `uint`, `uint8`, `uint16`, `uint32`, `uint64` - positive whole numbers only
   - `float32`, `float64` - decimal numbers
   
2. **Text:**
   - `string` - for text like "hello"
   - `rune` - for single characters

3. **Boolean:**
   - `bool` - can be either true or false

4. **Others:**
   - Arrays, slices, maps, structs, pointers, interfaces, channels

---

### Question 5: What is the zero value in Go?

**Answer:**
In Go, when you create a variable without giving it a value, it automatically gets a "zero value" instead of being undefined:

- Numbers (int, float): `0`
- Strings: `""` (empty string)
- Boolean: `false`
- Pointers, slices, maps, channels, interfaces: `nil`

Example:
```go
var count int      // count is 0
var name string    // name is ""
var isActive bool  // isActive is false
```

---

### Question 6: How do you define a constant in Go?

**Answer:**
You use the `const` keyword to define values that never change:

```go
const Pi = 3.14159
const CompanyName = "Google"
const MaxUsers = 100
```

Constants cannot be changed after they are defined. You cannot use `:=` with constants.

---

### Question 7: Explain the difference between var, :=, and const.

**Answer:**

1. **var** - Declares a variable that can be changed later:
   ```go
   var age int = 25
   age = 26  // can change
   ```

2. **:=** - Short way to declare and assign (only inside functions):
   ```go
   name := "John"  // shorter way
   ```

3. **const** - Declares a value that never changes:
   ```go
   const Pi = 3.14
   // Pi = 3.15  // ERROR - cannot change
   ```

---

### Question 8: What is the purpose of init() function in Go?

**Answer:**
The `init()` function is a special function that runs automatically before your main program starts. It's used to set up or initialize things:

```go
func init() {
    // This runs automatically before main()
    fmt.Println("Setting up...")
}

func main() {
    fmt.Println("Program starts")
}
```

You can have multiple `init()` functions in your program, and they all run before `main()`.

---

### Question 9: How do you write a for loop in Go?

**Answer:**
Go only has `for` loops (no while or do-while). There are several ways to use it:

1. **Traditional loop:**
   ```go
   for i := 0; i < 5; i++ {
       fmt.Println(i)
   }
   ```

2. **Like a while loop:**
   ```go
   i := 0
   for i < 5 {
       fmt.Println(i)
       i++
   }
   ```

3. **Infinite loop:**
   ```go
   for {
       // runs forever until you use break
   }
   ```

---

### Question 10: What is the difference between break, continue, and goto?

**Answer:**

1. **break** - Exits the loop completely:
   ```go
   for i := 0; i < 10; i++ {
       if i == 5 {
           break  // stops loop when i is 5
       }
   }
   ```

2. **continue** - Skips current iteration and goes to next:
   ```go
   for i := 0; i < 10; i++ {
       if i == 5 {
           continue  // skips when i is 5, but loop continues
       }
       fmt.Println(i)
   }
   ```

3. **goto** - Jumps to a labeled line (rarely used):
   ```go
   goto myLabel
   myLabel:
       fmt.Println("Jumped here")
   ```

---

### Question 11: What is a defer statement?

**Answer:**
`defer` makes a function run after the current function finishes, no matter how it ends:

```go
func example() {
    defer fmt.Println("This runs last")
    fmt.Println("This runs first")
    fmt.Println("This runs second")
}
```

Output:
```
This runs first
This runs second
This runs last
```

It's commonly used to close files or clean up resources.

---

### Question 12: How does defer work with return values?

**Answer:**
Deferred functions run after the return statement, but they can still modify named return values:

```go
func example() (result int) {
    defer func() {
        result = result + 10  // modifies return value
    }()
    return 5  // returns 5, but defer changes it to 15
}
```

This function returns 15, not 5.

---

### Question 13: What are named return values?

**Answer:**
You can give names to return values in the function signature:

```go
func divide(a, b int) (result int, err error) {
    if b == 0 {
        err = errors.New("cannot divide by zero")
        return  // returns result=0, err=error
    }
    result = a / b
    return  // returns the named variables
}
```

Named return values are automatically initialized to zero values and can be returned without specifying them.

---

### Question 14: What are variadic functions?

**Answer:**
Variadic functions can accept any number of arguments of the same type using `...`:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Can call with different number of arguments:
sum(1, 2)           // returns 3
sum(1, 2, 3, 4, 5)  // returns 15
```

---

### Question 15: What is a type alias?

**Answer:**
A type alias creates a new name for an existing type:

```go
type Miles int
type Kilometers int

var distance Miles = 100
```

This helps make your code more readable and type-safe. `Miles` and `Kilometers` are both `int`, but using different names makes the code clearer.

---

### Question 16: What is the difference between new() and make()?

**Answer:**

1. **new()** - Allocates memory and returns a pointer to zero value:
   ```go
   p := new(int)  // p is *int, points to 0
   ```

2. **make()** - Creates and initializes slices, maps, or channels:
   ```go
   s := make([]int, 5)     // creates a slice with 5 elements
   m := make(map[string]int) // creates an initialized map
   c := make(chan int)     // creates a channel
   ```

Use `make()` only for slices, maps, and channels. Use `new()` for other types (though rarely needed).

---

### Question 17: How do you handle errors in Go?

**Answer:**
Go uses explicit error checking. Functions return an error as the last return value:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Using the function:
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

Always check if `err != nil` after calling a function that can return an error.

---

### Question 18: What is panic and recover in Go?

**Answer:**

1. **panic** - Stops the program immediately when something goes very wrong:
   ```go
   func example() {
       panic("Something terrible happened!")
   }
   ```

2. **recover** - Catches a panic and prevents the program from crashing:
   ```go
   func safe() {
       defer func() {
           if r := recover(); r != nil {
               fmt.Println("Recovered from:", r)
           }
       }()
       panic("Oh no!")
   }
   ```

Use panic only for serious errors. For normal errors, use error values.

---

### Question 19: What are blank identifiers in Go?

**Answer:**
The blank identifier `_` is used to ignore values you don't need:

```go
// Ignore error:
result, _ := divide(10, 2)

// Ignore index in loop:
for _, value := range mySlice {
    fmt.Println(value)
}

// Import package only for side effects:
import _ "github.com/lib/pq"
```

---

## 🟡 Arrays, Slices, and Maps (Questions 21-40)

### Question 20: What is the difference between an array and a slice?

**Answer:**

**Array** - Fixed size, cannot grow or shrink:
```go
var arr [5]int  // array of exactly 5 integers
arr[0] = 10
```

**Slice** - Dynamic size, can grow and shrink:
```go
var slice []int  // can have any number of integers
slice = append(slice, 10)  // can add more elements
```

Slices are more flexible and commonly used.

---

### Question 21: How do you append to a slice?

**Answer:**
Use the built-in `append()` function:

```go
numbers := []int{1, 2, 3}
numbers = append(numbers, 4)        // adds one element
numbers = append(numbers, 5, 6, 7)  // adds multiple elements

// numbers is now [1, 2, 3, 4, 5, 6, 7]
```

**Important:** Always assign the result back to the variable because `append()` may create a new slice if there's not enough space.

---

### Question 22: What happens when a slice is appended beyond its capacity?

**Answer:**
When a slice doesn't have enough space, Go automatically:
1. Creates a new, bigger array (usually double the size)
2. Copies all existing elements to the new array
3. Adds the new element
4. Returns a slice pointing to the new array

```go
slice := make([]int, 2, 3)  // length=2, capacity=3
slice = append(slice, 1)     // fits in capacity
slice = append(slice, 2)     // needs to grow, new array created
```

---

### Question 23: How do you copy slices?

**Answer:**
Use the built-in `copy()` function:

```go
original := []int{1, 2, 3}
duplicate := make([]int, len(original))
copy(duplicate, original)

// Now duplicate has its own copy: [1, 2, 3]
```

**Important:** If you just assign `duplicate := original`, both variables point to the same data.

---

### Question 24: What is the difference between len() and cap()?

**Answer:**

- **len()** - Returns the current number of elements in the slice:
  ```go
  slice := []int{1, 2, 3}
  len(slice)  // returns 3
  ```

- **cap()** - Returns the maximum capacity before needing to allocate new memory:
  ```go
  slice := make([]int, 3, 5)  // length 3, capacity 5
  len(slice)  // returns 3
  cap(slice)  // returns 5
  ```

---

### Question 25: How do you create a multi-dimensional slice?

**Answer:**
Create a slice of slices:

```go
// 2D slice (like a matrix):
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Access elements:
fmt.Println(matrix[0][1])  // prints 2

// Create empty 2D slice:
rows := 3
cols := 4
matrix := make([][]int, rows)
for i := range matrix {
    matrix[i] = make([]int, cols)
}
```

---

### Question 26: How are slices passed to functions (by value or reference)?

**Answer:**
Slices are passed by value, but the slice contains a reference to the underlying array, so it behaves like passing by reference:

```go
func modify(s []int) {
    s[0] = 100  // This WILL change the original
}

func main() {
    numbers := []int{1, 2, 3}
    modify(numbers)
    fmt.Println(numbers)  // prints [100, 2, 3]
}
```

However, appending inside the function won't affect the original unless you return the slice.

---

### Question 27: What are maps in Go?

**Answer:**
Maps are key-value pairs, like a dictionary:

```go
// Create a map:
ages := make(map[string]int)

// Add values:
ages["John"] = 25
ages["Jane"] = 30

// Create with initial values:
scores := map[string]int{
    "Math": 95,
    "English": 87,
}
```

---

### Question 28: How do you check if a key exists in a map?

**Answer:**
Use the two-value assignment:

```go
ages := map[string]int{"John": 25}

age, exists := ages["John"]
if exists {
    fmt.Println("Age:", age)
} else {
    fmt.Println("Not found")
}

// Short form:
if age, ok := ages["John"]; ok {
    fmt.Println("Found:", age)
}
```

---

### Question 29: Can maps be compared directly?

**Answer:**
No, you cannot compare maps using `==` or `!=`:

```go
map1 := map[string]int{"a": 1}
map2 := map[string]int{"a": 1}

// map1 == map2  // ERROR: cannot compare

// You can only compare to nil:
if map1 == nil {
    fmt.Println("map is nil")
}
```

To compare maps, you must write a loop to check each key-value pair.

---

### Question 30: What happens if you delete a key from a map that doesn't exist?

**Answer:**
Nothing! Go doesn't throw an error:

```go
ages := map[string]int{"John": 25}
delete(ages, "Jane")  // no error, even though "Jane" doesn't exist
```

The `delete()` function safely handles non-existent keys.

---

### Question 31: Can slices be used as map keys?

**Answer:**
No, slices cannot be map keys because they cannot be compared:

```go
// This will NOT work:
// map[[]int]string  // ERROR

// Arrays CAN be used as keys (they have fixed size):
map[[3]int]string  // OK
```

Map keys must be comparable types (strings, numbers, arrays, structs without slices/maps).

---

### Question 32: How do you iterate over a map?

**Answer:**
Use a `for range` loop:

```go
ages := map[string]int{
    "John": 25,
    "Jane": 30,
}

// Get both key and value:
for name, age := range ages {
    fmt.Println(name, "is", age, "years old")
}

// Get only keys:
for name := range ages {
    fmt.Println(name)
}
```

**Note:** Maps are unordered, so iteration order is random.

---

### Question 33: How do you sort a map by key or value?

**Answer:**
Maps are unordered, so you need to extract keys/values and sort them:

```go
ages := map[string]int{"John": 25, "Alice": 30, "Bob": 20}

// Get keys and sort:
keys := make([]string, 0, len(ages))
for k := range ages {
    keys = append(keys, k)
}
sort.Strings(keys)

// Print in sorted order:
for _, k := range keys {
    fmt.Println(k, ages[k])
}
```

---

### Question 34: What are struct types in Go?

**Answer:**
Structs are custom data types that group related data together:

```go
type Person struct {
    Name string
    Age  int
}

// Create and use:
p := Person{
    Name: "John",
    Age:  25,
}

fmt.Println(p.Name)  // Access fields with dot
```

---

### Question 35: How do you define and use struct tags?

**Answer:**
Struct tags add metadata to fields, commonly used for JSON:

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{Name: "John", Age: 25}

// Convert to JSON:
data, _ := json.Marshal(p)
fmt.Println(string(data))  // {"name":"John","age":25}
```

---

### Question 36: How to embed one struct into another?

**Answer:**
You can embed structs to create composition:

```go
type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Address  // embedded
}

p := Person{
    Name: "John",
    Address: Address{
        City:    "New York",
        Country: "USA",
    },
}

fmt.Println(p.City)  // Can access directly
```

---

### Question 37: How do you compare two structs?

**Answer:**
Structs can be compared with `==` if all their fields are comparable:

```go
type Point struct {
    X, Y int
}

p1 := Point{1, 2}
p2 := Point{1, 2}

if p1 == p2 {
    fmt.Println("Equal")  // This works
}

// Structs with slices/maps cannot be compared:
type Container struct {
    Items []int
}
// c1 == c2  // ERROR
```

---

### Question 38: What is the difference between shallow and deep copy in structs?

**Answer:**

**Shallow copy** - Copies the struct but shares underlying data:
```go
type Person struct {
    Name    string
    Hobbies []string
}

p1 := Person{Name: "John", Hobbies: []string{"reading"}}
p2 := p1  // shallow copy

p2.Hobbies[0] = "gaming"  // This affects p1 too!
```

**Deep copy** - Creates completely independent copy:
```go
p2 := p1
p2.Hobbies = make([]string, len(p1.Hobbies))
copy(p2.Hobbies, p1.Hobbies)  // Now independent
```

---

### Question 39: How do you convert a struct to JSON?

**Answer:**
Use the `json.Marshal()` function:

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{Name: "John", Age: 25}
jsonData, err := json.Marshal(p)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(jsonData))  // {"name":"John","age":25}
```

---

## 🔵 Pointers, Interfaces, and Methods (Questions 41-60)

### Question 40: What are pointers in Go?

**Answer:**
A pointer stores the memory address of a value, not the value itself:

```go
var x int = 10
var p *int = &x  // p points to x

fmt.Println(x)   // prints 10
fmt.Println(p)   // prints memory address
fmt.Println(*p)  // prints 10 (value at the address)

*p = 20  // changes x to 20
```

- `&` gets the address
- `*` gets the value at that address

---

### Question 41: How do you declare and use pointers?

**Answer:**

```go
// Declare a pointer:
var p *int

// Get address of a variable:
x := 42
p = &x

// Access value through pointer:
fmt.Println(*p)  // prints 42

// Modify value through pointer:
*p = 100
fmt.Println(x)  // x is now 100
```

---

### Question 42: What is the difference between pointer and value receivers?

**Answer:**

**Value receiver** - Works on a copy:
```go
func (p Person) ChangeName(name string) {
    p.Name = name  // changes only the copy
}
```

**Pointer receiver** - Works on the original:
```go
func (p *Person) ChangeName(name string) {
    p.Name = name  // changes the original
}
```

Use pointer receivers when you want to:
- Modify the original value
- Avoid copying large structs

---

### Question 43: What are methods in Go?

**Answer:**
Methods are functions attached to a type:

```go
type Rectangle struct {
    Width, Height int
}

// Method for Rectangle:
func (r Rectangle) Area() int {
    return r.Width * r.Height
}

// Usage:
rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()  // calls the method
```

---

### Question 44: How to define an interface?

**Answer:**
An interface defines behavior (methods) without implementation:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Any type with these methods implements Shape:
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

// Circle automatically implements Shape
```

---

### Question 45: What is the empty interface in Go?

**Answer:**
The empty interface `interface{}` (or `any` in Go 1.18+) can hold any type:

```go
var anything interface{}

anything = 42
anything = "hello"
anything = []int{1, 2, 3}

// To use the value, you need type assertion:
str := anything.(string)
```

---

### Question 46: How do you perform type assertion?

**Answer:**
Type assertion extracts the concrete value from an interface:

```go
var i interface{} = "hello"

// Unsafe way (panics if wrong):
s := i.(string)

// Safe way (returns ok boolean):
s, ok := i.(string)
if ok {
    fmt.Println("It's a string:", s)
} else {
    fmt.Println("Not a string")
}
```

---

### Question 47: How to check if a type implements an interface?

**Answer:**
Use a compile-time check:

```go
type Writer interface {
    Write([]byte) error
}

type MyWriter struct{}

func (m MyWriter) Write(data []byte) error {
    return nil
}

// Check at compile time:
var _ Writer = MyWriter{}  // Compiles only if MyWriter implements Writer
```

---

### Question 48: Can interfaces be embedded?

**Answer:**
Yes, interfaces can embed other interfaces:

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

// Combined interface:
type ReadWriter interface {
    Reader
    Writer
}

// Types implementing ReadWriter must have both Read and Write methods
```

---

### Question 49: What is polymorphism in Go?

**Answer:**
Polymorphism means different types can be used through the same interface:

```go
type Animal interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

// Polymorphism in action:
animals := []Animal{Dog{}, Cat{}}
for _, animal := range animals {
    fmt.Println(animal.Speak())
}
```

---

### Question 50: How to use interfaces to write mockable code?

**Answer:**
Define interfaces for dependencies to make testing easier:

```go
// Interface:
type Database interface {
    Save(data string) error
}

// Real implementation:
type PostgresDB struct{}
func (p PostgresDB) Save(data string) error {
    // actual database code
    return nil
}

// Mock for testing:
type MockDB struct{}
func (m MockDB) Save(data string) error {
    // fake implementation for tests
    return nil
}

// Your code uses the interface:
func ProcessData(db Database, data string) {
    db.Save(data)  // works with both real and mock
}
```

---

### Question 51: What is the difference between interface{} and any?

**Answer:**
They are exactly the same! `any` is just a cleaner alias for `interface{}` introduced in Go 1.18:

```go
var x interface{}  // old way
var y any          // new way (Go 1.18+)

// Both can hold any type
```

Use `any` in new code - it's more readable.

---

### Question 52: What is duck typing?

**Answer:**
"If it walks like a duck and quacks like a duck, it's a duck."

In Go, if a type has the right methods, it automatically implements the interface:

```go
type Speaker interface {
    Speak() string
}

type Person struct{}
func (p Person) Speak() string { return "Hello" }

// Person implements Speaker automatically - no need to declare it!
```

---

### Question 53: Can you create an interface with no methods?

**Answer:**
Yes! That's the empty interface:

```go
type Empty interface{}  // or use 'any'

// Can hold any value:
var x Empty = 42
var y Empty = "hello"
```

---

### Question 54: Can structs implement multiple interfaces?

**Answer:**
Yes, a struct can implement as many interfaces as it wants:

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type File struct{}

func (f File) Read() string { return "data" }
func (f File) Write(s string) {}

// File implements both Reader and Writer
var r Reader = File{}
var w Writer = File{}
```

---

### Question 55: What is the difference between concrete type and interface type?

**Answer:**

**Concrete type** - An actual type with data:
```go
type Person struct {
    Name string
}
```

**Interface type** - Defines behavior only:
```go
type Speaker interface {
    Speak() string
}
```

Concrete types implement interfaces.

---

### Question 56: How to handle nil interfaces?

**Answer:**
An interface is nil only when both its type and value are nil:

```go
var i interface{}  // nil interface

var p *Person = nil
var i interface{} = p  // NOT nil! (has type information)

if i == nil {
    // This won't run
}

// To check for nil pointer inside interface:
if i != nil && reflect.ValueOf(i).IsNil() {
    fmt.Println("Nil pointer inside interface")
}
```

---

### Question 57: What are method sets?

**Answer:**
Method set is all the methods a type can call:

- **Value type** `T` can call methods with receiver `T`
- **Pointer type** `*T` can call methods with both `T` and `*T` receivers

```go
type Person struct{}

func (p Person) ValueMethod() {}
func (p *Person) PointerMethod() {}

var p1 Person
p1.ValueMethod()     // OK
p1.PointerMethod()   // OK (Go automatically converts)

var p2 *Person
p2.ValueMethod()     // OK
p2.PointerMethod()   // OK
```

---

### Question 58: Can a pointer implement an interface?

**Answer:**
Both pointers and values can implement interfaces, depending on the receiver type:

```go
type Writer interface {
    Write()
}

type Logger struct{}

func (l *Logger) Write() {}  // pointer receiver

// Now only *Logger implements Writer, not Logger:
var w Writer = &Logger{}  // OK
// var w Writer = Logger{}  // ERROR
```

---

### Question 59: What is the use of reflect package?

**Answer:**
The `reflect` package lets you inspect and manipulate types at runtime:

```go
import "reflect"

x := 42
t := reflect.TypeOf(x)
v := reflect.ValueOf(x)

fmt.Println(t)  // int
fmt.Println(v)  // 42

// Use cases:
// - Building generic libraries
// - JSON/XML encoding
// - ORMs
// - Validation frameworks
```

**Note:** Avoid reflection when possible - it's slower and less type-safe.

---

## 🟣 Concurrency and Goroutines (Questions 61-80)

### Question 60: What are goroutines?

**Answer:**
Goroutines are lightweight threads managed by Go. They let you run functions concurrently:

```go
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()  // runs in background
    
    // Without this, program may exit before sayHello finishes:
    time.Sleep(time.Second)
}
```

Goroutines use very little memory (a few KB) so you can run thousands of them.

---

### Question 61: How do you start a goroutine?

**Answer:**
Use the `go` keyword before a function call:

```go
// Regular function call (blocks):
doWork()

// Goroutine (runs in background):
go doWork()

// With anonymous function:
go func() {
    fmt.Println("Running in goroutine")
}()
```

---

### Question 62: What is a channel in Go?

**Answer:**
Channels are pipes that goroutines use to communicate:

```go
// Create a channel:
ch := make(chan int)

// Send value (blocks until received):
ch <- 42

// Receive value (blocks until sent):
value := <-ch
```

Channels help goroutines share data safely.

---

### Question 63: What is the difference between buffered and unbuffered channels?

**Answer:**

**Unbuffered** - Send blocks until receive happens:
```go
ch := make(chan int)  // unbuffered

ch <- 1  // blocks forever if no one receives
```

**Buffered** - Send blocks only when buffer is full:
```go
ch := make(chan int, 3)  // buffer size 3

ch <- 1  // doesn't block
ch <- 2  // doesn't block
ch <- 3  // doesn't block
ch <- 4  // blocks - buffer full
```

---

### Question 64: How do you close a channel?

**Answer:**
Use the `close()` function:

```go
ch := make(chan int)

go func() {
    ch <- 1
    ch <- 2
    close(ch)  // close when done sending
}()

// Receiver can detect closure:
for value := range ch {
    fmt.Println(value)
}
```

**Important:** Only the sender should close channels, never the receiver.

---

### Question 65: What happens when you send to a closed channel?

**Answer:**
Sending to a closed channel causes a **panic**:

```go
ch := make(chan int)
close(ch)

ch <- 1  // PANIC: send on closed channel
```

Always make sure not to send to closed channels.

---

### Question 66: How to detect a closed channel while receiving?

**Answer:**
Use the two-value receive:

```go
ch := make(chan int)
close(ch)

// Check if channel is closed:
value, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}

// Or use range (stops automatically when closed):
for value := range ch {
    fmt.Println(value)
}
```

---

### Question 67: What is the select statement in Go?

**Answer:**
`select` lets you wait on multiple channel operations:

```go
ch1 := make(chan int)
ch2 := make(chan int)

select {
case v := <-ch1:
    fmt.Println("Received from ch1:", v)
case v := <-ch2:
    fmt.Println("Received from ch2:", v)
default:
    fmt.Println("No channel ready")
}
```

---

### Question 68: How do you implement timeouts with select?

**Answer:**
Use `time.After()` in a select statement:

```go
ch := make(chan int)

select {
case v := <-ch:
    fmt.Println("Received:", v)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout! No response in 5 seconds")
}
```

---

### Question 69: What is a sync.WaitGroup?

**Answer:**
`WaitGroup` waits for multiple goroutines to finish:

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)  // increment counter
    
    go func(n int) {
        defer wg.Done()  // decrement when done
        fmt.Println(n)
    }(i)
}

wg.Wait()  // wait for all goroutines to finish
fmt.Println("All done!")
```

---

### Question 70: How does sync.Mutex work?

**Answer:**
`Mutex` (mutual exclusion) locks shared data so only one goroutine can access it:

```go
var counter int
var mu sync.Mutex

func increment() {
    mu.Lock()         // lock
    counter++         // safe to modify
    mu.Unlock()       // unlock
}

// Now multiple goroutines can safely call increment()
```

---

### Question 71: What is sync.Once?

**Answer:**
`sync.Once` ensures a function runs exactly once, even with multiple goroutines:

```go
var once sync.Once

func setup() {
    fmt.Println("Setting up...")
}

func doWork() {
    once.Do(setup)  // setup runs only once, no matter how many goroutines call this
}
```

Used for initialization that should happen only once.

---

### Question 72: How do you avoid race conditions?

**Answer:**
Race conditions happen when multiple goroutines access shared data. Avoid them by:

1. **Using mutexes:**
   ```go
   var mu sync.Mutex
   mu.Lock()
   // access shared data
   mu.Unlock()
   ```

2. **Using channels:**
   ```go
   ch <- data  // send
   data := <-ch  // receive
   ```

3. **Using atomic operations:**
   ```go
   atomic.AddInt64(&counter, 1)
   ```

---

### Question 73: What is the Go memory model?

**Answer:**
The Go memory model defines when changes made by one goroutine are visible to another. Key points:

- Changes in one goroutine may not be immediately visible to others
- Use synchronization (mutexes, channels) to coordinate
- Channels and mutexes provide "happens-before" guarantees

Don't rely on execution order - always synchronize!

---

### Question 74: How do you use context.Context for cancellation?

**Answer:**
`Context` helps cancel or timeout operations:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

go func() {
    select {
    case <-time.After(10 * time.Second):
        fmt.Println("Work done")
    case <-ctx.Done():
        fmt.Println("Cancelled:", ctx.Err())
    }
}()
```

---

### Question 75: How to pass data between goroutines?

**Answer:**
Use channels:

```go
// Pass data from producer to consumer:
dataCh := make(chan string)

// Producer:
go func() {
    dataCh <- "Hello from goroutine"
}()

// Consumer:
message := <-dataCh
fmt.Println(message)
```

---

### Question 76: What is the runtime.GOMAXPROCS() function?

**Answer:**
`GOMAXPROCS` sets the maximum number of CPU cores that can run Go code simultaneously:

```go
import "runtime"

// Use 4 CPU cores:
runtime.GOMAXPROCS(4)

// Get current setting:
cores := runtime.GOMAXPROCS(0)
```

By default, Go uses all available CPU cores.

---

### Question 77: How do you detect deadlocks in Go?

**Answer:**
Go's runtime automatically detects deadlocks and panics:

```go
func main() {
    ch := make(chan int)
    <-ch  // deadlock! waiting for data that never comes
}

// Output: fatal error: all goroutines are asleep - deadlock!
```

Use the race detector during development:
```bash
go run -race myprogram.go
```

---

### Question 78: What are worker pools and how do you implement them?

**Answer:**
Worker pools limit the number of concurrent workers:

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Create 3 workers:
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send 9 jobs:
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results:
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```

---

### Question 79: How to write concurrent-safe data structures?

**Answer:**
Use mutexes to protect shared data:

```go
type SafeCounter struct {
    mu    sync.Mutex
    count map[string]int
}

func (c *SafeCounter) Inc(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count[key]++
}

func (c *SafeCounter) Value(key string) int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count[key]
}
```

---

## 🔴 Advanced & Best Practices (Questions 81-100)

### Question 80: How does Go handle memory management?

**Answer:**
Go uses automatic memory management:

1. **Stack** - For local variables in functions (fast, automatic cleanup)
2. **Heap** - For data that needs to live longer (garbage collected)
3. **Garbage Collector** - Automatically frees unused memory

You don't need to manually allocate or free memory like in C/C++.

---

### Question 81: What is garbage collection in Go?

**Answer:**
The garbage collector (GC) automatically removes unused memory:

- Runs in the background
- Uses a "concurrent mark-and-sweep" algorithm
- Pauses your program briefly (usually < 1ms)
- You can tune it if needed, but defaults work well

```go
import "runtime"

// Manually trigger GC (rarely needed):
runtime.GC()
```

---

### Question 82: How do you profile CPU and memory in Go?

**Answer:**
Use the `pprof` tool:

```go
import (
    "net/http"
    _ "net/http/pprof"
)

func main() {
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Your program...
}
```

Then visit `http://localhost:6060/debug/pprof/` in your browser or use:
```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

---

### Question 83: What is the difference between compile-time and runtime errors?

**Answer:**

**Compile-time errors** - Caught when building the code:
```go
var x int = "hello"  // ERROR: type mismatch
```

**Runtime errors** - Happen when the program is running:
```go
var x []int
x[0] = 10  // Panic: index out of range
```

Go catches many errors at compile time, making programs safer.

---

### Question 84: How to use go test for unit testing?

**Answer:**
Create a file ending in `_test.go`:

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

Run tests:
```bash
go test
```

---

### Question 85: What is table-driven testing in Go?

**Answer:**
Test multiple cases with one test function:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {10, 5, 15},
    }
    
    for _, tt := range tests {
        got := Add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d", 
                tt.a, tt.b, got, tt.want)
        }
    }
}
```

---

### Question 86: How to benchmark code in Go?

**Answer:**
Use benchmark functions:

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

Run benchmarks:
```bash
go test -bench=.
```

---

### Question 87: What is go mod and how does it work?

**Answer:**
`go mod` manages dependencies (external packages):

```bash
# Initialize module:
go mod init github.com/myname/myproject

# Add dependencies (automatically):
go get github.com/gin-gonic/gin

# Update dependencies:
go get -u

# Remove unused dependencies:
go mod tidy
```

Dependencies are listed in `go.mod` file.

---

### Question 88: What is vendoring in Go modules?

**Answer:**
Vendoring copies dependencies into your project:

```bash
go mod vendor
```

Creates a `vendor/` folder with all dependencies. Useful for:
- Ensuring exact versions
- Working offline
- Faster CI builds

---

### Question 89: How to handle versioning in modules?

**Answer:**
Use semantic versioning (v1.2.3):

```go
// go.mod
module example.com/myproject

go 1.20

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/lib/pq v1.10.7
)
```

Get specific version:
```bash
go get github.com/gin-gonic/gin@v1.8.0
```

---

### Question 90: How do you structure a Go project?

**Answer:**
Common structure:

```
myproject/
├── cmd/              # Main applications
│   └── server/
│       └── main.go
├── internal/         # Private code
│   ├── handlers/
│   ├── models/
│   └── database/
├── pkg/              # Public libraries
├── api/              # API definitions
├── web/              # Static files
├── go.mod
└── README.md
```

---

### Question 91: What is the idiomatic way to name Go packages?

**Answer:**
Package naming conventions:

- Use **short, lowercase** names: `http`, `json`, `user`
- **No underscores or camelCase**: use `httputil` not `http_util` or `httpUtil`
- **Be descriptive**: `auth` is better than `a`
- **Don't repeat parent name**: `user.User` not `user.UserModel`

---

### Question 92: What is the purpose of the internal package?

**Answer:**
Code in `internal/` can only be imported by nearby code:

```
myproject/
├── internal/
│   └── auth/        # Only myproject can import this
│       └── token.go
└── api/
    └── handler.go   # Can import myproject/internal/auth
```

External projects cannot import your `internal/` packages.

---

### Question 93: How do you handle logging in Go?

**Answer:**
Use the `log` package or structured logging libraries:

```go
import "log"

// Basic logging:
log.Println("Starting server...")
log.Printf("Port: %d", 8080)

// With custom logger:
logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
logger.Println("Server started")

// Or use libraries like "logrus" or "zap" for structured logging
```

---

### Question 94: What is the difference between log.Fatal, log.Panic, and log.Println?

**Answer:**

1. **log.Println** - Just logs the message:
   ```go
   log.Println("Something happened")
   // Program continues
   ```

2. **log.Fatal** - Logs and exits program (calls os.Exit(1)):
   ```go
   log.Fatal("Critical error!")
   // Program stops immediately
   ```

3. **log.Panic** - Logs and panics (can be recovered):
   ```go
   log.Panic("Error!")
   // Can be caught with recover()
   ```

---

### Question 95: What are build tags in Go?

**Answer:**
Build tags let you compile different code for different situations:

```go
// +build linux

package main

import "fmt"

func platformMessage() {
    fmt.Println("Running on Linux")
}
```

```go
// +build windows

package main

import "fmt"

func platformMessage() {
    fmt.Println("Running on Windows")
}
```

Build with:
```bash
go build -tags linux
```

---

### Question 96: What are cgo and its use cases?

**Answer:**
`cgo` lets Go code call C code:

```go
/*
#include <stdio.h>
void hello() {
    printf("Hello from C!\n");
}
*/
import "C"

func main() {
    C.hello()
}
```

Use cases:
- Using existing C libraries
- Performance-critical code
- System-level programming

**Note:** cgo makes builds slower and less portable.

---

### Question 97: What are some common Go anti-patterns?

**Answer:**

1. **Ignoring errors:**
   ```go
   result, _ := doSomething()  // BAD
   ```

2. **Not closing resources:**
   ```go
   file, _ := os.Open("file.txt")
   // forgot to defer file.Close()
   ```

3. **Premature optimization:**
   - Optimize only after profiling

4. **Using panic for normal errors:**
   ```go
   if err != nil {
       panic(err)  // BAD - use error returns
   }
   ```

---

### Question 98: What are Go code quality tools (lint, vet, staticcheck)?

**Answer:**

1. **go fmt** - Formats code:
   ```bash
   go fmt ./...
   ```

2. **go vet** - Finds suspicious code:
   ```bash
   go vet ./...
   ```

3. **golint** - Style checker:
   ```bash
   golint ./...
   ```

4. **staticcheck** - Advanced static analysis:
   ```bash
   staticcheck ./...
   ```

5. **golangci-lint** - Runs multiple linters:
   ```bash
   golangci-lint run
   ```

---

### Question 99: What are the best practices for writing idiomatic Go code?

**Answer:**

1. **Handle errors explicitly** - Don't ignore them
2. **Use short variable names** - `i`, `err`, `buf` instead of `index`, `errorValue`
3. **Return early** - Avoid deep nesting
4. **Accept interfaces, return structs**
5. **Keep functions small** - Do one thing well
6. **Use defer for cleanup**
7. **Prefer composition over inheritance**
8. **Make zero values useful**
9. **Use gofmt** - Let the tool format your code
10. **Write tests** - Test coverage matters

---

*This document contains Questions 1-99. Continue reading for Questions 100-433...*

---

## 🟢 Project Structure & Design Patterns (Questions 101-120)

### Question 100: How do you organize a large-scale Go project?

**Answer:**
For large projects, use this structure:

```
project/
├── cmd/                   # Entry points
│   ├── api/
│   │   └── main.go       # API server
│   └── worker/
│       └── main.go       # Background worker
├── internal/              # Private application code
│   ├── domain/           # Business logic
│   ├── handler/          # HTTP handlers
│   ├── repository/       # Data access
│   └── service/          # Business services
├── pkg/                   # Public libraries
│   └── utils/
├── api/                   # API specs (OpenAPI, protobuf)
├── configs/               # Configuration files
├── migrations/            # Database migrations
├── scripts/               # Build/deploy scripts
└── go.mod
```

**Key principles:**
- Separate business logic from infrastructure
- Use `internal/` for private code
- Keep `cmd/` lightweight - just initialization
- Group by feature, not by type

---

### Question 101: What is the standard Go project layout?

**Answer:**
The standard Go project layout is a community-recommended structure:

```
myproject/
├── cmd/                  # Main applications
│   └── myapp/
│       └── main.go       # Application entry point
├── internal/             # Private application code
├── pkg/                  # Public libraries
├── api/                  # API definitions (OpenAPI, protobuf)
├── web/                  # Web assets
├── configs/              # Configuration files
├── scripts/              # Build and install scripts
├── build/                # Packaging and CI
├── deployments/          # Docker, k8s configs
├── test/                 # Additional test data
└── docs/                 # Documentation
```

This is not enforced by the compiler but widely adopted for consistency across Go projects.

---

### Question 102: What is the cmd directory used for in Go?

**Answer:**
The `cmd/` directory contains the entry points for your application:

```
cmd/
├── api/
│   └── main.go         # API server
├── worker/
│   └── main.go         # Background worker
└── cli/
    └── main.go         # CLI tool
```

Each subdirectory represents a separate executable. The main.go files should be minimal - just initialization and wiring, with actual logic in `internal/` or `pkg/`.

**Example:**
```go
// cmd/api/main.go
package main

import (
    "myapp/internal/server"
)

func main() {
    srv := server.New()
    srv.Start()
}
```

---

### Question 103: How do you structure code for reusable packages?

**Answer:**
Put reusable code in the `pkg/` directory:

```
pkg/
├── logger/              # Logging utilities
├── database/            # Database helpers
└── validator/           # Validation logic
```

**Guidelines:**
- Make packages focused on one responsibility
- Use clear, descriptive names
- Minimize dependencies
- Document exported functions

**Example:**
```go
// pkg/logger/logger.go
package logger

import "log"

func Info(msg string) {
    log.Printf("[INFO] %s", msg)
}

func Error(msg string) {
    log.Printf("[ERROR] %s", msg)
}
```

---

### Question 104: What are Go's most used design patterns?

**Answer:**
Common design patterns in Go:

1. **Factory Pattern** - Create objects without specifying exact class
2. **Builder Pattern** - Construct complex objects step by step
3. **Singleton Pattern** - Ensure only one instance exists
4. **Strategy Pattern** - Define family of algorithms using interfaces
5. **Decorator Pattern** - Add behavior to objects dynamically
6. **Repository Pattern** - Abstract data access layer
7. **Adapter Pattern** - Make incompatible interfaces work together
8. **Observer Pattern** - Notify multiple objects of state changes

Go favors composition over inheritance, so patterns look different than in OOP languages.

---

### Question 105: Explain the Factory Pattern in Go.

**Answer:**
Factory pattern creates objects without exposing creation logic:

```go
type Database interface {
    Connect() error
}

type PostgresDB struct{}
func (p PostgresDB) Connect() error {
    return nil
}

type MongoDB struct{}
func (m MongoDB) Connect() error {
    return nil
}

// Factory function:
func NewDatabase(dbType string) Database {
    switch dbType {
    case "postgres":
        return PostgresDB{}
    case "mongo":
        return MongoDB{}
    default:
        return nil
    }
}

// Usage:
db := NewDatabase("postgres")
db.Connect()
```

---

### Question 106: How to implement Singleton Pattern in Go?

**Answer:**
Use `sync.Once` to ensure initialization happens only once:

```go
package config

import "sync"

type Config struct {
    DatabaseURL string
    Port        int
}

var (
    instance *Config
    once     sync.Once
)

func GetConfig() *Config {
    once.Do(func() {
        instance = &Config{
            DatabaseURL: "localhost:5432",
            Port:        8080,
        }
    })
    return instance
}

// Usage:
cfg := config.GetConfig()  // First call creates instance
cfg2 := config.GetConfig() // Returns same instance
```

---

### Question 107: What is Dependency Injection in Go?

**Answer:**
Dependency Injection means passing dependencies to a struct instead of creating them inside:

```go
// Without DI (bad):
type UserService struct {
    db *sql.DB
}

func NewUserService() *UserService {
    db, _ := sql.Open("postgres", "...")  // Hardcoded
    return &UserService{db: db}
}

// With DI (good):
type UserService struct {
    db Database  // Interface
}

func NewUserService(db Database) *UserService {
    return &UserService{db: db}
}

// Now you can inject different implementations:
realDB := &PostgresDB{}
service := NewUserService(realDB)

// Or inject a mock for testing:
mockDB := &MockDB{}
testService := NewUserService(mockDB)
```

---

### Question 108: What is the difference between composition and inheritance in Go?

**Answer:**
Go doesn't have inheritance - it uses composition:

**Composition** - Build complex types from simpler ones:
```go
type Engine struct {
    Power int
}

func (e Engine) Start() {
    fmt.Println("Engine started")
}

type Car struct {
    Engine  // Embedded - Car "has-a" Engine
    Model string
}

car := Car{
    Engine: Engine{Power: 200},
    Model:  "Tesla",
}
car.Start()  // Can call Engine's methods
```

Benefits:
- More flexible than inheritance
- Easier to test
- Clearer relationships between types

---

### Question 109: What are Go generics and how do you use them?

**Answer:**
Generics (added in Go 1.18) let you write functions and types that work with any type:

```go
// Before generics - need separate functions:
func SumInts(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// With generics - one function for all number types:
func Sum[T int | float64](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

// Usage:
Sum([]int{1, 2, 3})          // Works with int
Sum([]float64{1.5, 2.5})     // Works with float64
```

---

### Question 110: How to implement a generic function with constraints?

**Answer:**
Use type constraints to limit which types can be used:

```go
// Define a constraint:
type Number interface {
    int | int64 | float64
}

// Generic function with constraint:
func Max[T Number](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Usage:
Max(10, 20)           // Works
Max(3.14, 2.71)       // Works
// Max("a", "b")      // ERROR - string not in Number
```

---

### Question 111: What are type parameters?

**Answer:**
Type parameters are placeholders for actual types in generic code:

```go
// [T any] is a type parameter
func Print[T any](value T) {
    fmt.Println(value)
}

// T gets replaced with actual type:
Print[int](42)        // T is int
Print[string]("hi")   // T is string
Print(true)           // T inferred as bool

// Multiple type parameters:
func Pair[K comparable, V any](key K, value V) map[K]V {
    return map[K]V{key: value}
}

m := Pair("age", 25)  // K=string, V=int
```

---

### Question 112: Can you implement the Strategy pattern using interfaces?

**Answer:**
Yes! Strategy pattern defines a family of algorithms through interfaces:

```go
type PaymentStrategy interface {
    Pay(amount float64) error
}

type CreditCard struct{}
func (c CreditCard) Pay(amount float64) error {
    fmt.Printf("Paid $%.2f with credit card\n", amount)
    return nil
}

type PayPal struct{}
func (p PayPal) Pay(amount float64) error {
    fmt.Printf("Paid $%.2f with PayPal\n", amount)
    return nil
}

type Checkout struct {
    strategy PaymentStrategy
}

func (c *Checkout) Process(amount float64) {
    c.strategy.Pay(amount)
}

// Usage:
checkout := &Checkout{strategy: CreditCard{}}
checkout.Process(100)  // Uses credit card

checkout.strategy = PayPal{}
checkout.Process(50)   // Switches to PayPal
```

---

### Question 113: What is middleware in Go web apps?

**Answer:**
Middleware is code that runs before/after your main handler:

```go
// Middleware function:
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)  // Call next handler
    })
}

// Another middleware:
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// Chain middlewares:
handler := LoggingMiddleware(AuthMiddleware(http.HandlerFunc(myHandler)))
http.ListenAndServe(":8080", handler)
```

---

### Question 114: How do you structure code using the Clean Architecture?

**Answer:**
Clean Architecture separates code into layers:

```
internal/
├── domain/           # Business entities (no dependencies)
│   └── user.go
├── usecase/          # Business logic
│   └── user_usecase.go
├── repository/       # Data access interfaces
│   └── user_repository.go
├── handler/          # HTTP handlers
│   └── user_handler.go
└── infrastructure/   # External concerns (DB, API)
    └── postgres/
        └── user_repo_impl.go
```

**Example:**
```go
// domain/user.go
type User struct {
    ID   int
    Name string
}

// repository/user_repository.go
type UserRepository interface {
    Save(user User) error
    FindByID(id int) (User, error)
}

// usecase/user_usecase.go
type UserUseCase struct {
    repo repository.UserRepository
}

func (uc *UserUseCase) CreateUser(name string) error {
    user := domain.User{Name: name}
    return uc.repo.Save(user)
}
```

---

### Question 115: What are service and repository layers?

**Answer:**

**Repository Layer** - Handles data access:
```go
type UserRepository interface {
    Create(user User) error
    GetByID(id int) (User, error)
    Update(user User) error
    Delete(id int) error
}

type PostgresUserRepo struct {
    db *sql.DB
}

func (r *PostgresUserRepo) Create(user User) error {
    // SQL insert logic
    return nil
}
```

**Service Layer** - Contains business logic:
```go
type UserService struct {
    repo UserRepository
}

func (s *UserService) RegisterUser(name, email string) error {
    // Business logic:
    if name == "" {
        return errors.New("name required")
    }
    
    user := User{Name: name, Email: email}
    return s.repo.Create(user)
}
```

---

### Question 116: How would you separate concerns in a RESTful Go app?

**Answer:**
Separate into layers with clear responsibilities:

```go
// 1. Handler Layer - HTTP concerns
type UserHandler struct {
    service *UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    json.NewDecoder(r.Body).Decode(&req)
    
    err := h.service.CreateUser(req.Name, req.Email)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    
    w.WriteHeader(201)
}

// 2. Service Layer - Business logic
type UserService struct {
    repo UserRepository
}

func (s *UserService) CreateUser(name, email string) error {
    // Validation and business rules
    if !isValidEmail(email) {
        return errors.New("invalid email")
    }
    return s.repo.Create(User{Name: name, Email: email})
}

// 3. Repository Layer - Data access
type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) Create(user User) error {
    // Database operations
    return nil
}
```

---

### Question 117: What is the importance of interfaces in layered design?

**Answer:**
Interfaces enable loose coupling and testability:

```go
// Define interface in service layer:
type EmailSender interface {
    Send(to, subject, body string) error
}

type UserService struct {
    emailSender EmailSender  // Depends on interface, not implementation
}

// Production implementation:
type SMTPSender struct{}
func (s SMTPSender) Send(to, subject, body string) error {
    // Real email sending
    return nil
}

// Test implementation:
type MockSender struct{}
func (m MockSender) Send(to, subject, body string) error {
    // Just log, don't actually send
    return nil
}

// Easy to swap implementations:
prodService := UserService{emailSender: SMTPSender{}}
testService := UserService{emailSender: MockSender{}}
```

---

### Question 118: How would you implement a plugin system in Go?

**Answer:**
Use Go plugins or interface-based plugins:

**Using Interfaces (recommended):**
```go
// Define plugin interface:
type Plugin interface {
    Name() string
    Execute() error
}

// Plugin implementations:
type LoggerPlugin struct{}
func (l LoggerPlugin) Name() string { return "logger" }
func (l LoggerPlugin) Execute() error {
    fmt.Println("Logging...")
    return nil
}

type CachePlugin struct{}
func (c CachePlugin) Name() string { return "cache" }
func (c CachePlugin) Execute() error {
    fmt.Println("Caching...")
    return nil
}

// Plugin manager:
type PluginManager struct {
    plugins []Plugin
}

func (pm *PluginManager) Register(p Plugin) {
    pm.plugins = append(pm.plugins, p)
}

func (pm *PluginManager) RunAll() {
    for _, p := range pm.plugins {
        fmt.Printf("Running %s\n", p.Name())
        p.Execute()
    }
}
```

---

### Question 119: How do you avoid circular dependencies in Go packages?

**Answer:**
Circular dependencies occur when package A imports B and B imports A. Solutions:

**1. Extract common code to a new package:**
```
// Before (circular):
pkg/user/    imports pkg/auth/
pkg/auth/    imports pkg/user/   ❌

// After (fixed):
pkg/user/    imports pkg/common/
pkg/auth/    imports pkg/common/  ✅
pkg/common/  (shared types)
```

**2. Use interfaces:**
```go
// auth/auth.go
type UserGetter interface {  // Define interface instead of importing user package
    GetUser(id int) User
}

type AuthService struct {
    userGetter UserGetter
}
```

**3. Dependency Inversion:**
- High-level modules shouldn't depend on low-level modules
- Both should depend on abstractions (interfaces)

---

## 🟡 Generics, Type System, and Advanced Types (Questions 120-140)

### Question 120: What is type inference in Go?

**Answer:**
Type inference means Go automatically determines the type:

```go
// Explicit type:
var x int = 42

// Type inference:
x := 42           // Go infers int
name := "John"    // Go infers string
pi := 3.14        // Go infers float64

// Works with generics too:
func Print[T any](v T) {
    fmt.Println(v)
}

Print(42)      // T inferred as int
Print("hi")    // T inferred as string
```

---

### Question 121: How do you use generics with struct types?

**Answer:**
Define generic structs using type parameters:

```go
// Generic stack:
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage:
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)

stringStack := Stack[string]{}
stringStack.Push("hello")
```

---

### Question 122: Can you restrict generic types using constraints?

**Answer:**
Yes, use interfaces as constraints:

```go
// Built-in constraint:
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Custom constraint:
type Numeric interface {
    int | int64 | float32 | float64
}

func Sum[T Numeric](values []T) T {
    var total T
    for _, v := range values {
        total += v
    }
    return total
}

// Multiple constraints:
type ComparableNumber interface {
    comparable           // Can use == and !=
    int | float64       // Must be one of these types
}
```

---

### Question 123: How to create reusable generic containers (e.g., Stack)?

**Answer:**
Create generic data structures:

```go
// Generic Stack:
type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{items: []T{}}
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if s.IsEmpty() {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if s.IsEmpty() {
        var zero T
        return zero, false
    }
    return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

// Usage:
stack := NewStack[int]()
stack.Push(1)
stack.Push(2)
val, ok := stack.Pop()  // val=2, ok=true
```

---

### Question 124: What is the difference between any and interface{}?

**Answer:**
They are **exactly the same** - `any` is an alias for `interface{}`:

```go
// These are identical:
func Print1(v interface{}) {
    fmt.Println(v)
}

func Print2(v any) {
    fmt.Println(v)
}

// In Go 1.18+, prefer 'any' - it's clearer:
func GenericFunc[T any](value T) T {
    return value
}
```

---

### Question 125: Can you have multiple constraints in a generic function?

**Answer:**
Yes, combine constraints using interfaces:

```go
// Method constraint:
type Stringer interface {
    String() string
}

// Combined constraint:
type StringableNumber interface {
    Stringer
    int | float64
}

// Multiple type parameters with different constraints:
func Process[K comparable, V any](key K, value V) map[K]V {
    return map[K]V{key: value}
}

// Union of constraints:
type NumberOrString interface {
    int | float64 | string
}
```

---

### Question 126: Can interfaces be used in generics?

**Answer:**
Yes, interfaces work as constraints:

```go
type Writer interface {
    Write() error
}

// Generic function accepting any Writer:
func Save[T Writer](item T) error {
    return item.Write()
}

type File struct{}
func (f File) Write() error {
    fmt.Println("Writing file...")
    return nil
}

type Database struct{}
func (d Database) Write() error {
    fmt.Println("Writing to DB...")
    return nil
}

// Both work:
Save(File{})
Save(Database{})
```

---

### Question 127: What is type embedding and how does it differ from inheritance?

**Answer:**

**Type embedding** - Composition of types:
```go
type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello"
}

type Employee struct {
    Person  // Embedded
    ID int
}

emp := Employee{
    Person: Person{Name: "John"},
    ID:     123,
}

emp.Speak()  // Can call Person's methods
fmt.Println(emp.Name)  // Direct access to Person's fields
```

**Key differences from inheritance:**
- No "is-a" relationship
- Can embed multiple types
- No method overriding (can shadow methods)
- More explicit and flexible

---

### Question 128: How does Go perform type conversion vs. type assertion?

**Answer:**

**Type Conversion** - Between compatible types:
```go
var x int = 42
var y float64 = float64(x)  // Convert int to float64

var s string = "123"
// num := int(s)  // ERROR - incompatible types
```

**Type Assertion** - Extract value from interface:
```go
var i interface{} = "hello"

// Unsafe (panics if wrong):
s := i.(string)

// Safe (returns ok boolean):
s, ok := i.(string)
if ok {
    fmt.Println("It's a string:", s)
}

// Type switch:
switch v := i.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Int:", v)
}
```

---

### Question 129: What are tagged unions and how can you simulate them in Go?

**Answer:**
Tagged unions (sum types) represent "one of several types". Simulate with interfaces:

```go
// Define a marker interface:
type Shape interface {
    isShape()  // Marker method
}

type Circle struct {
    Radius float64
}
func (Circle) isShape() {}

type Rectangle struct {
    Width, Height float64
}
func (Rectangle) isShape() {}

// Function accepting any Shape:
func Area(s Shape) float64 {
    switch shape := s.(type) {
    case Circle:
        return 3.14 * shape.Radius * shape.Radius
    case Rectangle:
        return shape.Width * shape.Height
    default:
        return 0
    }
}

// Usage:
Area(Circle{Radius: 5})
Area(Rectangle{Width: 10, Height: 5})
```

---

### Question 130: What is the use of iota in Go?

**Answer:**
`iota` creates auto-incrementing constants:

```go
type Weekday int

const (
    Sunday    Weekday = iota  // 0
    Monday                     // 1
    Tuesday                    // 2
    Wednesday                  // 3
    Thursday                   // 4
    Friday                     // 5
    Saturday                   // 6
)

// Skip values:
const (
    _ = iota  // 0 (ignored)
    KB = 1 << (10 * iota)  // 1024
    MB                      // 1048576
    GB                      // 1073741824
)

// Custom starting value:
const (
    First = iota + 1  // 1
    Second            // 2
    Third             // 3
)
```

---

### Question 131: How are custom types different from type aliases?

**Answer:**

**Custom Type** - Creates a new, distinct type:
```go
type Celsius float64
type Fahrenheit float64

var temp1 Celsius = 25
var temp2 Fahrenheit = 77

// temp1 = temp2  // ERROR - different types
temp1 = Celsius(temp2)  // Need explicit conversion
```

**Type Alias** - Just another name for existing type:
```go
type MyInt = int  // Notice the '='

var x MyInt = 42
var y int = x  // OK - they're the same type
```

---

### Question 132: What are type sets in Go 1.18+?

**Answer:**
Type sets define which types satisfy an interface:

```go
// Old style - method set:
type Writer interface {
    Write() error
}

// New style - type set (Go 1.18+):
type Integer interface {
    int | int32 | int64  // Set of types
}

type Signed interface {
    ~int | ~int8 | ~int16  // ~ means "underlying type"
}

// Combining methods and type sets:
type StringableInt interface {
    ~int
    String() string
}
```

---

### Question 133: Can generic types implement interfaces?

**Answer:**
Yes, generic types can implement interfaces:

```go
type Stringer interface {
    String() string
}

// Generic type:
type Pair[T any] struct {
    First, Second T
}

// Implement interface:
func (p Pair[T]) String() string {
    return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

// Now Pair implements Stringer:
var s Stringer = Pair[int]{First: 1, Second: 2}
fmt.Println(s.String())  // (1, 2)
```

---

### Question 134: How do you handle constraints with operations like +, -, *?

**Answer:**
Use type constraints that support those operations:

```go
import "golang.org/x/exp/constraints"

// For +, -, *, /:
func Add[T constraints.Integer | constraints.Float](a, b T) T {
    return a + b
}

// For comparison (<, >, <=, >=):
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Custom numeric constraint:
type Number interface {
    ~int | ~int64 | ~float32 | ~float64
}

func Multiply[T Number](a, b T) T {
    return a * b
}
```

---

### Question 135: What is structural typing?

**Answer:**
Structural typing means types are compatible if they have the same structure:

```go
// Go uses structural typing for interfaces:
type Reader interface {
    Read() string
}

type File struct{}
func (f File) Read() string { return "file data" }

type Network struct{}
func (n Network) Read() string { return "network data" }

// Both implement Reader automatically (no explicit declaration needed)
func Process(r Reader) {
    fmt.Println(r.Read())
}

Process(File{})     // Works
Process(Network{})  // Works
```

This is different from nominal typing where you must explicitly declare that a type implements an interface.

---

### Question 136: Explain the difference between concrete and abstract types.

**Answer:**

**Concrete Type** - Has actual data and implementation:
```go
type User struct {
    ID   int
    Name string
}

var u User  // Can create instances
```

**Abstract Type** - Defines behavior only (interfaces):
```go
type Database interface {
    Save(data string) error
    Load(id int) (string, error)
}

// var db Database  // Can't instantiate directly
// Need a concrete type that implements it:
var db Database = PostgresDB{}
```

---

### Question 137: What are phantom types and are they used in Go?

**Answer:**
Phantom types carry compile-time information without runtime data. Go doesn't have true phantom types, but you can simulate them:

```go
// Phantom type for compile-time safety:
type Validated struct{}
type Unvalidated struct{}

type Email[T any] struct {
    value string
    _     T  // Phantom type (not used at runtime)
}

func NewEmail(value string) Email[Unvalidated] {
    return Email[Unvalidated]{value: value}
}

func (e Email[Unvalidated]) Validate() (Email[Validated], error) {
    if !strings.Contains(e.value, "@") {
        return Email[Validated]{}, errors.New("invalid email")
    }
    return Email[Validated]{value: e.value}, nil
}

func SendEmail(e Email[Validated]) {
    // Can only send validated emails
}

// SendEmail(NewEmail("test"))  // ERROR - can't send unvalidated
validEmail, _ := NewEmail("test@example.com").Validate()
SendEmail(validEmail)  // OK
```

---

### Question 138: How would you implement an enum pattern in Go?

**Answer:**
Use constants with iota and custom types:

```go
type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)

// Add string representation:
func (s Status) String() string {
    switch s {
    case Pending:
        return "Pending"
    case Approved:
        return "Approved"
    case Rejected:
        return "Rejected"
    default:
        return "Unknown"
    }
}

// Add validation:
func (s Status) IsValid() bool {
    return s >= Pending && s <= Rejected
}

// Usage:
var orderStatus Status = Pending
fmt.Println(orderStatus)  // Prints "Pending"

if orderStatus == Approved {
    // Handle approved
}
```

---

### Question 139: How can you implement optional values in Go idiomatically?

**Answer:**
Go doesn't have Option/Maybe types, but you can use these patterns:

**1. Pointer (most common):**
```go
type User struct {
    Name  string
    Email *string  // Optional
}

user := User{Name: "John"}  // Email is nil
if user.Email != nil {
    fmt.Println(*user.Email)
}
```

**2. Two-return pattern:**
```go
func FindUser(id int) (User, bool) {
    // Return value and "found" boolean
    if user, exists := cache[id]; exists {
        return user, true
    }
    return User{}, false
}

user, found := FindUser(123)
if found {
    // Use user
}
```

**3. Generic Option type (Go 1.18+):**
```go
type Option[T any] struct {
    value T
    valid bool
}

func Some[T any](value T) Option[T] {
    return Option[T]{value: value, valid: true}
}

func None[T any]() Option[T] {
    return Option[T]{valid: false}
}

func (o Option[T]) Get() (T, bool) {
    return o.value, o.valid
}

// Usage:
opt := Some(42)
if val, ok := opt.Get(); ok {
    fmt.Println(val)
}
```

---
# Go Interview Questions (Continuation 140-433)

## ðŸ”µ Networking, APIs, and Web Dev (Questions 140-170)

### Question 140: How to build a REST API in Go?

**Answer:**
Build RESTful APIs using the `net/http` package or frameworks:

```go
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    users := []User{
        {ID: 1, Name: "John"},
        {ID: 2, Name: "Jane"},
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    
    // Save user to database...
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func main() {
    http.HandleFunc("/users", getUsers)
    http.HandleFunc("/users/create", createUser)
    http.ListenAndServe(":8080", nil)
}
```

---

### Question 141: How to parse JSON and XML in Go?

**Answer:**

**JSON Parsing:**
```go
import "encoding/json"

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Parse JSON to struct:
jsonData := []byte(`{"name":"John","age":30}`)
var person Person
json.Unmarshal(jsonData, &person)

// Convert struct to JSON:
jsonBytes, _ := json.Marshal(person)
fmt.Println(string(jsonBytes))
```

**XML Parsing:**
```go
import "encoding/xml"

type Book struct {
    Title  string `xml:"title"`
    Author string `xml:"author"`
}

// Parse XML:
xmlData := []byte(`<book><title>Go Programming</title><author>John</author></book>`)
var book Book
xml.Unmarshal(xmlData, &book)

// Convert to XML:
xmlBytes, _ := xml.Marshal(book)
```

---

### Question 142: What is the use of http.Handler and http.HandlerFunc?

**Answer:**

**http.Handler** - Interface with ServeHTTP method:
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

// Custom handler:
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

http.Handle("/", MyHandler{})
```

**http.HandlerFunc** - Function type that implements Handler:
```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

http.HandleFunc("/", myHandler)
// Or: http.Handle("/", http.HandlerFunc(myHandler))
```

---

###  Question 143: How do you implement middleware manually in Go?

**Answer:**
Create functions that wrap handlers:

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// Chain middlewares:
func main() {
    finalHandler := http.HandlerFunc(homeHandler)
    withAuth := authMiddleware(finalHandler)
    withLogging := loggingMiddleware(withAuth)
    
    http.ListenAndServe(":8080", withLogging)
}
```

---

### Question 144: How do you serve static files in Go?

**Answer:**
Use `http.FileServer`:

```go
// Serve files from ./static directory:
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))

// Or serve a specific file:
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
})

http.ListenAndServe(":8080", nil)
```

Project structure:
```
myproject/
â”œâ”€â”€ static/
â”‚   â”œâ”€â”€ css/
â”‚   â”œâ”€â”€ js/
â”‚   â””â”€â”€ images/
â””â”€â”€ main.go
```

---

### Question 145: How do you handle CORS in Go?

**Answer:**
Set CORS headers manually or use a library:

**Manual approach:**
```go
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

handler := enableCORS(http.HandlerFunc(myHandler))
```

**Using gorilla/handlers:**
```go
import "github.com/gorilla/handlers"

handler := handlers.CORS(
    handlers.AllowedOrigins([]string{"*"}),
    handlers.AllowedMethods([]string{"GET", "POST"}),
)(myHandler)
```

---

### Question 146: What are context-based timeouts in HTTP servers?

**Answer:**
Use `context.Context` to enforce timeouts:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Create timeout context:
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    
    // Make slow operation:
    result := make(chan string)
    go func() {
        // Simulate slow work:
        time.Sleep(3 * time.Second)
        result <- "Done"
    }()
    
    select {
    case res := <-result:
        fmt.Fprintf(w, res)
    case <-ctx.Done():
        http.Error(w, "Request timeout", 504)
    }
}
```

---

### Question 147: How do you make HTTP requests in Go?

**Answer:**
Use the `net/http` package:

```go
// GET request:
resp, err := http.Get("https://api.example.com/users")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body))

// POST request with JSON:
user := User{Name: "John", Age: 30}
jsonData, _ := json.Marshal(user)

resp, err = http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(jsonData),
)

// Custom request:
req, _ := http.NewRequest("PUT", "https://api.example.com/users/1", body)
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer token")

client := &http.Client{}
resp, err = client.Do(req)
```

---

### Question 148: How do you manage connection pooling in Go?

**Answer:**
Configure the HTTP client transport:

```go
transport := &http.Transport{
    MaxIdleConns:        100,              // Max idle connections
    MaxIdleConnsPerHost: 10,               // Max idle per host
    IdleConnTimeout:     90 * time.Second, // Idle timeout
}

client := &http.Client{
    Transport: transport,
    Timeout:   10 * time.Second,
}

// Reuse the client for multiple requests:
resp, err := client.Get("https://api.example.com")
```

Connection pooling happens automatically - connections are reused when possible.

---

### Question 149: What is an HTTP client timeout?

**Answer:**
Set timeout to prevent hanging requests:

```go
// Client-level timeout:
client := &http.Client{
    Timeout: 10 * time.Second,  // Total timeout for request
}

resp, err := client.Get("https://slow-api.example.com")

// Request-level timeout using context:
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
resp, err := client.Do(req)

// Different timeouts for different stages:
transport := &http.Transport{
    DialContext: (&net.Dialer{
        Timeout:   5 * time.Second,  // Connection timeout
        KeepAlive: 30 * time.Second,
    }).DialContext,
    TLSHandshakeTimeout:   10 * time.Second,  // TLS timeout
    ResponseHeaderTimeout: 10 * time.Second,   // Header timeout
}
```

---

### Question 150: How do you upload and download files via HTTP?

**Answer:**

**Upload file:**
```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Parse multipart form (32MB max):
    r.ParseMultipartForm(32 << 20)
    
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    defer file.Close()
    
    // Create file on server:
    dst, _ := os.Create("./uploads/" + handler.Filename)
    defer dst.Close()
    
    io.Copy(dst, file)
    fmt.Fprintf(w, "File uploaded successfully")
}
```

**Download file:**
```go
func downloadHandler(w http.ResponseWriter, r *http.Request) {
    filename := "document.pdf"
    
    w.Header().Set("Content-Disposition", "attachment; filename="+filename)
    w.Header().Set("Content-Type", "application/pdf")
    
    http.ServeFile(w, r, "./files/"+filename)
}
```

---

### Question 151: What is graceful shutdown and how do you implement it?

**Answer:**
Graceful shutdown finishes active requests before stopping:

```go
func main() {
    srv := &http.Server{
        Addr: ":8080",
        Handler: myHandler,
    }
    
    // Start server in goroutine:
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()
    
    // Wait for interrupt signal:
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    log.Println("Shutting down server...")
    
    // Graceful shutdown with 5-second timeout:
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }
    
    log.Println("Server exited")
}
```

---

### Question 152: How to work with multipart/form-data in Go?

**Answer:**
Handle file uploads and form fields:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Parse multipart form:
    if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB max
        http.Error(w, err.Error(), 400)
        return
    }
    
    // Get regular form fields:
    name := r.FormValue("name")
    description := r.FormValue("description")
    
    // Get uploaded files:
    files := r.MultipartForm.File["files"]
    
    for _, fileHeader := range files {
        file, _ := fileHeader.Open()
        defer file.Close()
        
        // Save file:
        dst, _ := os.Create("./uploads/" + fileHeader.Filename)
        defer dst.Close()
        
        io.Copy(dst, file)
    }
    
    fmt.Fprintf(w, "Uploaded %d files", len(files))
}
```

---

### Question 153: How do you implement rate limiting in Go?

**Answer:**
Use golang.org/x/time/rate or custom implementation:

**Using rate limiter:**
```go
import "golang.org/x/time/rate"

// Create limiter: 5 requests per second, burst of 10:
limiter := rate.NewLimiter(5, 10)

func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Rate limit exceeded", 429)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

**Custom token bucket:**
```go
type RateLimiter struct {
    tokens    int
    capacity  int
    refillRate time.Duration
    mu        sync.Mutex
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    if rl.tokens > 0 {
        rl.tokens--
        return true
    }
    return false
}
```

---

*[Due to length constraints, I'm creating this as a separate file. The user can merge this with the main file.]*
### Question 154: What is Gorilla Mux and how does it compare with net/http?

**Answer:**
Gorilla Mux is a powerful router with more features than net/http:

**net/http (basic):**
```go
http.HandleFunc("/users", usersHandler)
// Limited pattern matching
```

**Gorilla Mux (advanced):**
```go
import "github.com/gorilla/mux"

r := mux.NewRouter()

// URL parameters:
r.HandleFunc("/users/{id}", getUserHandler)

// Query parameters, methods, headers:
r.HandleFunc("/api/users", getUsers).Methods("GET")
r.HandleFunc("/api/users", createUser).Methods("POST")

// Host matching:
r.Host("api.example.com").HandleFunc("/", apiHandler)

// Subrouters:
api := r.PathPrefix("/api").Subrouter()
api.HandleFunc("/users", usersHandler)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]  // Get URL parameter
    fmt.Fprintf(w, "User ID: %s", id)
}
```

**Benefits of Gorilla Mux:**
- URL path variables
- Method-based routing
- Regex in routes
- Middleware support
- Subrouters

---

### Question 155: What are Go frameworks for web APIs (Gin, Echo)?

**Answer:**
Popular Go web frameworks:

**1. Gin - Fast HTTP framework:**
```go
import "github.com/gin-gonic/gin"

r := gin.Default()  // With logger and recovery

r.GET("/users/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"id": id})
})

r.POST("/users", func(c *gin.Context) {
    var user User
    c.BindJSON(&user)
    c.JSON(201, user)
})

r.Run(":8080")
```

**2. Echo - High performance framework:**
```go
import "github.com/labstack/echo/v4"

e := echo.New()

e.GET("/users/:id", func(c echo.Context) error {
    id := c.Param("id")
    return c.JSON(200, map[string]string{"id": id})
})

e.Start(":8080")
```

**3. Fiber - Express-inspired:**
```go
import "github.com/gofiber/fiber/v2"

app := fiber.New()

app.Get("/users/:id", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"id": c.Params("id")})
})

app.Listen(":8080")
```

---

### Question 156: What are the trade-offs between using http.ServeMux and third-party routers?

**Answer:**

**http.ServeMux (Standard library):**

**Pros:**
- No external dependencies
- Simple and lightweight
- Sufficient for basic routing
- Well-tested and stable

**Cons:**
- No URL parameters
- No regex matching
- Limited pattern matching
- No middleware support built-in

**Third-party routers (Gorilla Mux, Chi, etc.):**

**Pros:**
- URL parameters: `/users/{id}`
- Method-based routing
- Middleware support
- Advanced features
- Better developer experience

**Cons:**
- External dependency
- Slightly more overhead
- Learning curve

**When to use each:**
- **Use ServeMux** for simple APIs, microservices with few routes
- **Use third-party** for complex routing needs, RESTful APIs

---

### Question 157: How would you implement authentication in a Go API?

**Answer:**
Implement JWT-based authentication:

```go
import "github.com/golang-jwt/jwt/v5"

type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}

var jwtSecret = []byte("your-secret-key")

// Generate JWT token:
func generateToken(userID int) (string, error) {
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// Verify token middleware:
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        
        // Remove "Bearer " prefix:
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", 401)
            return
        }
        
        // Add user ID to context:
        ctx := context.WithValue(r.Context(), "userID", claims.UserID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Login handler:
func login(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&creds)
    
    // Verify credentials (simplified):
    if creds.Username == "admin" && creds.Password == "password" {
        token, _ := generateToken(1)
        json.NewEncoder(w).Encode(map[string]string{"token": token})
    } else {
        http.Error(w, "Invalid credentials", 401)
    }
}
```

---

### Question 158: How do you implement file streaming in Go?

**Answer:**
Stream large files efficiently:

```go
func streamFile(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open("large-video.mp4")
    if err != nil {
        http.Error(w, "File not found", 404)
        return
    }
    defer file.Close()
    
    // Get file info:
    stat, _ := file.Stat()
    
    // Set headers:
    w.Header().Set("Content-Type", "video/mp4")
    w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
    
    // Stream file:
    io.Copy(w, file)  // Efficiently streams in chunks
}

// Stream with buffer control:
func streamWithBuffer(w http.ResponseWriter, r *http.Request) {
    file, _ := os.Open("large-file.dat")
    defer file.Close()
    
    buffer := make([]byte, 4096)  // 4KB chunks
    
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        w.Write(buffer[:n])
        
        // Flush to client:
        if f, ok := w.(http.Flusher); ok {
            f.Flush()
        }
    }
}
```

---

### Question 159: What is the HTTP/2 Server Push and how to use it?

**Answer:**
HTTP/2 Server Push sends resources before the client requests them:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Check if pusher is supported:
    pusher, ok := w.(http.Pusher)
    if ok {
        // Push CSS before client requests it:
        if err := pusher.Push("/style.css", nil); err != nil {
            log.Printf("Failed to push: %v", err)
        }
    }
    
    // Serve main HTML:
    http.ServeFile(w, r, "index.html")
}

func main() {
    http.HandleFunc("/", handler)
    
    // HTTP/2 requires HTTPS:
    log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
```

---

## ðŸŸ£ Databases and ORMs (Questions 160-180)

### Question 160: How do you connect to a PostgreSQL database in Go?

**Answer:**
Use database/sql with a PostgreSQL driver:

```go
import (
    "database/sql"
    _ "github.com/lib/pq"  // PostgreSQL driver
)

func main() {
    connStr := "host=localhost port=5432 user=postgres password=secret dbname=mydb sslmode=disable"
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Test connection:
    if err := db.Ping(); err != nil {
        log.Fatal("Cannot connect:", err)
    }
    
    fmt.Println("Connected to PostgreSQL!")
}
```

---

### Question 161: What is the difference between database/sql and GORM?

**Answer:**

**database/sql (Standard library):**
- Low-level SQL operations
- Manual query writing
- Full control
- No magic/hidden queries
- More verbose

```go
rows, err := db.Query("SELECT id, name FROM users WHERE age > $1", 18)
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
}
```

**GORM (ORM):**
- High-level object mapping
- Auto migrations
- Associations/relations
- Less boilerplate
- Query builder

```go
type User struct {
    ID   uint
    Name string
    Age  int
}

db.AutoMigrate(&User{})

var users []User
db.Where("age > ?", 18).Find(&users)
```

**When to use:**
- **database/sql**: Performance-critical, complex queries, full control
- **GORM**: Rapid development, CRUD operations, less SQL knowledge needed

---

### Question 162: How do you handle SQL injections in Go?

**Answer:**
Always use parameterized queries:

**âŒ BAD (Vulnerable to SQL injection):**
```go
username := r.FormValue("username")
query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username)
rows, _ := db.Query(query)  // DANGEROUS!
```

**âœ… GOOD (Safe from SQL injection):**
```go
username := r.FormValue("username")
rows, _ := db.Query("SELECT * FROM users WHERE username = $1", username)
```

**Why it works:**
- Database driver escapes parameters
- SQL and data are sent separately
- No string concatenation

**Examples:**
```go
// INSERT:
stmt, _ := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)")
stmt.Exec(name, email)

// UPDATE:
db.Exec("UPDATE users SET name = $1 WHERE id = $2", newName, userID)

// DELETE:
db.Exec("DELETE FROM users WHERE id = $1", userID)

// Multiple parameters:
db.Query("SELECT * FROM orders WHERE user_id = $1 AND status = $2", userID, status)
```

---

### Question 163: How do you manage connection pools in database/sql?

**Answer:**
Configure connection pool settings:

```go
db, err := sql.Open("postgres", connStr)

// Maximum number of open connections:
db.SetMaxOpenConns(25)

// Maximum number of idle connections:
db.SetMaxIdleConns(5)

// Maximum lifetime of a connection:
db.SetConnMaxLifetime(5 * time.Minute)

// Maximum idle time:
db.SetConnMaxIdleTime(10 * time.Minute)

// Get connection stats:
stats := db.Stats()
fmt.Printf("Open connections: %d\n", stats.OpenConnections)
fmt.Printf("In use: %d\n", stats.InUse)
fmt.Printf("Idle: %d\n", stats.Idle)
```

**Best practices:**
- Set MaxOpenConns to match database server capacity
- Keep MaxIdleConns reasonable (5-10)
- Set ConnMaxLifetime to prevent stale connections
- Monitor stats in production

---

### Question 164: What are prepared statements in Go?

**Answer:**
Prepared statements compile SQL once and reuse it:

```go
// Prepare statement:
stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// Execute multiple times:
for _, user := range users {
    _, err := stmt.Exec(user.Name, user.Email)
    if err != nil {
        log.Printf("Error inserting %s: %v", user.Name, err)
    }
}

// SELECT with prepared statement:
stmt, _ = db.Prepare("SELECT name, email FROM users WHERE id = $1")
defer stmt.Close()

var name, email string
err = stmt.QueryRow(123).Scan(&name, &email)
```

**Benefits:**
- Better performance (compiled once)
- Protection against SQL injection
- Reusable across multiple executions

---

### Question 165: How do you map SQL rows to structs?

**Answer:**
Use Scan to map rows to struct fields:

```go
type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

// Single row:
func getUser(db *sql.DB, id int) (*User, error) {
    user := &User{}
    
    err := db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name, &user.Email, &user.Age)
    
    if err == sql.ErrNoRows {
        return nil, errors.New("user not found")
    }
    
    return user, err
}

// Multiple rows:
func getUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, name, email, age FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    
    return users, rows.Err()
}

// Using sqlx for automatic mapping:
import "github.com/jmoiron/sqlx"

var users []User
db.Select(&users, "SELECT * FROM users")  // Automatic mapping!
```

---

### Question 166: What are transactions and how are they implemented in Go?

**Answer:**
Transactions ensure atomic operations:

```go
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
    // Begin transaction:
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Defer rollback (will be skipped if committed):
    defer tx.Rollback()
    
    // Deduct from sender:
    _, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err  // Transaction will be rolled back
    }
    
    // Add to receiver:
    _, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err  // Transaction will be rolled back
    }
    
    // Commit transaction:
    return tx.Commit()
}

// Using context for timeout:
func transferWithTimeout(db *sql.DB, fromID, toID int, amount float64) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    tx, _ := db.BeginTx(ctx, nil)
    defer tx.Rollback()
    
    _, err := tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err
    }
    
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err
    }
    
    return tx.Commit()
}
```

---

### Question 167: How do you handle database migrations in Go?

**Answer:**
Use migration tools like golang-migrate:

**Using golang-migrate:**
```bash
# Install:
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration:
migrate create -ext sql -dir migrations -seq create_users_table
```

**Migration files:**
```sql
-- migrations/000001_create_users_table.up.sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

-- migrations/000001_create_users_table.down.sql
DROP TABLE users;
```

**Run migrations in code:**
```go
import "github.com/golang-migrate/migrate/v4"
import _ "github.com/golang-migrate/migrate/v4/database/postgres"
import _ "github.com/golang-migrate/migrate/v4/source/file"

func runMigrations(dbURL string) error {
    m, err := migrate.New(
        "file://migrations",
        dbURL,
    )
    if err != nil {
        return err
    }
    
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }
    
    return nil
}
```

---

### Question 168: What is the use of sqlx in Go?

**Answer:**
sqlx extends database/sql with convenient features:

```go
import "github.com/jmoiron/sqlx"

db, _ := sqlx.Connect("postgres", connStr)

type User struct {
    ID    int    `db:"id"`
    Name  string `db:"name"`
    Email string `db:"email"`
}

// Get single row:
var user User
err := db.Get(&user, "SELECT * FROM users WHERE id = $1", 123)

// Get multiple rows:
var users []User
err = db.Select(&users, "SELECT * FROM users WHERE age > $1", 18)

// Named queries:
query := `INSERT INTO users (name, email) VALUES (:name, :email)`
_, err = db.NamedExec(query, map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
})

// Struct binding:
newUser := User{Name: "Jane", Email: "jane@example.com"}
_, err = db.NamedExec(`INSERT INTO users (name, email) VALUES (:name, :email)`, newUser)

// In queries:
query, args, _ := sqlx.In("SELECT * FROM users WHERE id IN (?)", []int{1, 2, 3})
query = db.Rebind(query)  // Convert ? to $1, $2 for PostgreSQL
db.Select(&users, query, args...)
```

---

### Question 169: What are the pros and cons of using an ORM in Go?

**Answer:**

**Pros:**
1. **Faster development** - Less boilerplate code
2. **Auto migrations** - Schema management
3. **Type safety** - Compile-time checks
4. **Relationship handling** - Automatic joins
5. **Database agnostic** - Easy to switch databases
6. **Query builder** - Readable, chainable queries

**Cons:**
1. **Performance overhead** - Generated queries not always optimal
2. **Learning curve** - Need to learn ORM API
3. **Hidden complexity** - Hard to debug generated SQL
4. **N+1 problem** - Can generate many queries
5. **Complex queries** - Sometimes easier in raw SQL
6. **Black box** - Less control over exact SQL

**Example:**
```go
// GORM (ORM):
var users []User
db.Where("age > ?", 18).Find(&users)

// vs database/sql (raw):
rows, _ := db.Query("SELECT * FROM users WHERE age > $1", 18)
defer rows.Close()
for rows.Next() {
    var u User
    rows.Scan(&u.ID, &u.Name, &u.Age)
    users = append(users, u)
}
```

**Recommendation:**
- **Use ORM** for standard CRUD, rapid prototyping
- **Use raw SQL** for complex queries, performance-critical code
- Mix both as needed

---

### Question 170: How would you implement pagination in SQL queries?

**Answer:**
Use LIMIT and OFFSET:

```go
type PaginationParams struct {
    Page     int
    PageSize int
}

func getUsers(db *sql.DB, params PaginationParams) ([]User, error) {
    // Calculate offset:
    offset := (params.Page - 1) * params.PageSize
    
    query := `
        SELECT id, name, email
        FROM users
        ORDER BY id
        LIMIT $1 OFFSET $2
    `
    
    rows, err := db.Query(query, params.PageSize, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    
    return users, nil
}

// With total count:
type PaginatedResponse struct {
    Data       []User `json:"data"`
    Page       int    `json:"page"`
    PageSize   int    `json:"page_size"`
    TotalCount int    `json:"total_count"`
    TotalPages int    `json:"total_pages"`
}

func getUsersPaginated(db *sql.DB, params PaginationParams) (*PaginatedResponse, error) {
    // Get total count:
    var totalCount int
    db.QueryRow("SELECT COUNT(*) FROM users").Scan(&totalCount)
    
    // Get paginated data:
    offset := (params.Page - 1) * params.PageSize
    rows, _ := db.Query("SELECT id, name, email FROM users LIMIT $1 OFFSET $2", 
        params.PageSize, offset)
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    
    return &PaginatedResponse{
        Data:       users,
        Page:       params.Page,
        PageSize:   params.PageSize,
        TotalCount: totalCount,
        TotalPages: (totalCount + params.PageSize - 1) / params.PageSize,
    }, nil
}

// Cursor-based pagination (better for large datasets):
func getUsersCursorBased(db *sql.DB, cursor int, limit int) ([]User, error) {
    query := `
        SELECT id, name, email
        FROM users
        WHERE id > $1
        ORDER BY id
        LIMIT $2
    `
    
    rows, _ := db.Query(query, cursor, limit)
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    
    return users, nil
}
```

---

### Question 171: How do you log SQL queries in Go?

**Answer:**
Implement query logging:

**Method 1: Custom logger wrapper:**
```go
type LoggedDB struct {
    *sql.DB
}

func (db *LoggedDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
    start := time.Now()
    
    log.Printf("[SQL Query] %s | Args: %v", query, args)
    
    rows, err := db.DB.Query(query, args...)
    
    duration := time.Since(start)
    if err != nil {
        log.Printf("[SQL Error] %v | Duration: %v", err, duration)
    } else {
        log.Printf("[SQL Success] Duration: %v", duration)
    }
    
    return rows, err
}

func (db *LoggedDB) Exec(query string, args ...interface{}) (sql.Result, error) {
    start := time.Now()
    log.Printf("[SQL Exec] %s | Args: %v", query, args)
    
    result, err := db.DB.Exec(query, args...)
    
    duration := time.Since(start)
    log.Printf("[SQL] Duration: %v | Error: %v", duration, err)
    
    return result, err
}
```

**Method 2: GORM logging:**
```go
import "gorm.io/gorm/logger"

newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags),
    logger.Config{
        SlowThreshold: time.Second,   // Slow SQL threshold
        LogLevel:      logger.Info,   // Log level
        Colorful:      true,          // Color output
    },
)

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: newLogger,
})
```

---

### Question 172: What is the N+1 problem in ORMs and how to avoid it?

**Answer:**
N+1 problem occurs when fetching related data in a loop:

**âŒ BAD (N+1 Problem):**
```go
// 1 query to get users:
var users []User
db.Find(&users)  // SELECT * FROM users

// N queries for each user's posts:
for i := range users {
    db.Model(&users[i]).Association("Posts").Find(&users[i].Posts)
    // SELECT * FROM posts WHERE user_id = 1
    // SELECT * FROM posts WHERE user_id = 2
    // SELECT * FROM posts WHERE user_id = 3
    // ... N more queries!
}
```

**âœ… GOOD (Using Preload):**
```go
// Single query with JOIN:
var users []User
db.Preload("Posts").Find(&users)
// SELECT * FROM users
// SELECT * FROM posts WHERE user_id IN (1,2,3,...)
```

**Other solutions:**
```go
// Eager loading multiple associations:
db.Preload("Posts").Preload("Comments").Find(&users)

// Nested preloading:
db.Preload("Posts.Comments").Find(&users)

// Joins (single query):
db.Joins("Posts").Find(&users)

// Raw SQL with proper joins:
db.Raw(`
    SELECT users.*, posts.*
    FROM users
    LEFT JOIN posts ON posts.user_id = users.id
`).Scan(&results)
```

---

### Question 173: How do you implement caching for DB queries in Go?

**Answer:**
Use Redis or in-memory cache:

**Using Redis:**
```go
import (
    "github.com/go-redis/redis/v8"
    "encoding/json"
)

var rdb *redis.Client

func getUser(db *sql.DB, id int) (*User, error) {
    ctx := context.Background()
    cacheKey := fmt.Sprintf("user:%d", id)
    
    // Try cache first:
    cached, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        var user User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }
    
    // Cache miss - query database:
    var user User
    err = db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name, &user.Email)
    
    if err != nil {
        return nil, err
    }
    
    // Store in cache:
    userData, _ := json.Marshal(user)
    rdb.Set(ctx, cacheKey, userData, 10*time.Minute)
    
    return &user, nil
}

// Invalidate cache on update:
func updateUser(db *sql.DB, user *User) error {
    _, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3",
        user.Name, user.Email, user.ID)
    
    if err == nil {
        // Invalidate cache:
        cacheKey := fmt.Sprintf("user:%d", user.ID)
        rdb.Del(context.Background(), cacheKey)
    }
    
    return err
}
```

**Using in-memory cache (sync.Map):**
```go
var cache sync.Map

type CacheItem struct {
    Data      interface{}
    ExpiresAt time.Time
}

func getCached(key string) (interface{}, bool) {
    val, ok := cache.Load(key)
    if !ok {
        return nil, false
    }
    
    item := val.(CacheItem)
    if time.Now().After(item.ExpiresAt) {
        cache.Delete(key)
        return nil, false
    }
    
    return item.Data, true
}

func setCache(key string, data interface{}, ttl time.Duration) {
    cache.Store(key, CacheItem{
        Data:      data,
        ExpiresAt: time.Now().Add(ttl),
    })
}
```

---

### Question 174: How do you write custom SQL queries using GORM?

**Answer:**
GORM provides multiple ways to execute raw SQL:

```go
// Raw SQL query:
var users []User
db.Raw("SELECT * FROM users WHERE age > ?", 18).Scan(&users)

// With named parameters:
db.Raw("SELECT * FROM users WHERE name = @name", 
    sql.Named("name", "John")).Scan(&users)

// Execute SQL:
db.Exec("UPDATE users SET active = ? WHERE last_login < ?", true, time.Now())

// Complex query:
var result []map[string]interface{}
db.Raw(`
    SELECT u.name, COUNT(p.id) as post_count
    FROM users u
    LEFT JOIN posts p ON p.user_id = u.id
    GROUP BY u.name
    HAVING COUNT(p.id) > 5
`).Scan(&result)

// Mix raw SQL with GORM:
db.Where("age > ?", 18).
    Where("created_at > ?", time.Now().AddDate(0, -1, 0)).
    Find(&users)

// Combine with Scan:
type UserStats struct {
    UserID     int
    PostCount  int
    CommentCount int
}

var stats []UserStats
db.Raw(`
    SELECT 
        u.id as user_id,
        COUNT(DISTINCT p.id) as post_count,
        COUNT(DISTINCT c.id) as comment_count
    FROM users u
    LEFT JOIN posts p ON p.user_id = u.id
    LEFT JOIN comments c ON c.user_id = u.id
    GROUP BY u.id
`).Scan(&stats)
```

---

### Question 175: How do you handle one-to-many and many-to-many relationships in GORM?

**Answer:**

**One-to-Many:**
```go
type User struct {
    ID    uint
    Name  string
    Posts []Post  // Has many posts
}

type Post struct {
    ID     uint
    Title  string
    UserID uint    // Foreign key
    User   User    // Belongs to user
}

// Create with associations:
user := User{
    Name: "John",
    Posts: []Post{
        {Title: "First Post"},
        {Title: "Second Post"},
    },
}
db.Create(&user)

// Query with associations:
var user User
db.Preload("Posts").First(&user, 1)

// Add post to existing user:
post := Post{Title: "New Post", UserID: user.ID}
db.Create(&post)
```

**Many-to-Many:**
```go
type User struct {
    ID    uint
    Name  string
    Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
    ID    uint
    Name  string
    Users []User `gorm:"many2many:user_roles;"`
}

// GORM creates join table automatically:
// CREATE TABLE user_roles (user_id INT, role_id INT)

// Create with associations:
user := User{
    Name: "John",
    Roles: []Role{
        {Name: "Admin"},
        {Name: "Editor"},
    },
}
db.Create(&user)

// Query:
var user User
db.Preload("Roles").First(&user, 1)

// Add role to user:
var admin Role
db.First(&admin, "name = ?", "Admin")
db.Model(&user).Association("Roles").Append(&admin)

// Remove role:
db.Model(&user).Association("Roles").Delete(&admin)

// Replace all roles:
db.Model(&user).Association("Roles").Replace(&newRoles)

// Clear all:
db.Model(&user).Association("Roles").Clear()

// Count:
count := db.Model(&user).Association("Roles").Count()
```

---

### Question 176: How would you structure your database layer in a Go project?

**Answer:**
Use repository pattern for clean architecture:

```
internal/
â”œâ”€â”€ domain/
â”‚   â””â”€â”€ user.go           # Entity definitions
â”œâ”€â”€ repository/
â”‚   â”œâ”€â”€ interface.go      # Repository interfaces
â”‚   â””â”€â”€ postgres/
â”‚       â””â”€â”€ user_repo.go  # PostgreSQL implementation
â””â”€â”€ service/
    â””â”€â”€ user_service.go   # Business logic
```

**domain/user.go:**
```go
package domain

type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}
```

**repository/interface.go:**
```go
package repository

type UserRepository interface {
    Create(user *domain.User) error
    GetByID(id int) (*domain.User, error)
    GetAll() ([]domain.User, error)
    Update(user *domain.User) error
    Delete(id int) error
}
```

**repository/postgres/user_repo.go:**
```go
package postgres

type PostgresUserRepo struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
    return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Create(user *domain.User) error {
    query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
    return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func (r *PostgresUserRepo) GetByID(id int) (*domain.User, error) {
    user := &domain.User{}
    query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
    err := r.db.QueryRow(query, id).Scan(
        &user.ID, &user.Name, &user.Email, &user.CreatedAt,
    )
    return user, err
}

func (r *PostgresUserRepo) GetAll() ([]domain.User, error) {
    rows, err := r.db.Query(`SELECT id, name, email FROM users`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []domain.User
    for rows.Next() {
        var u domain.User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    return users, nil
}
```

**service/user_service.go:**
```go
package service

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) error {
    // Business logic validation:
    if !isValidEmail(email) {
        return errors.New("invalid email")
    }
    
    user := &domain.User{Name: name, Email: email}
    return s.repo.Create(user)
}
```

---

### Question 177: What is context propagation in database calls?

**Answer:**
Pass context through database operations for timeout/cancellation:

```go
// Create context with timeout:
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use context in query:
rows, err := db.QueryContext(ctx, "SELECT * FROM large_table")
if err != nil {
    if err == context.DeadlineExceeded {
        log.Println("Query timed out!")
    }
    return err
}

// In transactions:
func transferMoney(ctx context.Context, db *sql.DB, from, to int, amount float64) error {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, from)
    if err != nil {
        return err
    }
    
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, to)
    if err != nil {
        return err
    }
    
    return tx.Commit()
}

// With HTTP request context:
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()  // Get request context
    
    var users []User
    rows, err := db.QueryContext(ctx, "SELECT * FROM users")
    // Query automatically cancelled if client disconnects
}

// Repository with context:
type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*User, error) {
    user := &User{}
    err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name)
    return user, err
}
```

---

### Question 178: How do you handle long-running queries or timeouts?

**Answer:**
Use context timeouts and query optimization:

```go
// Set query timeout:
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, `
    SELECT * FROM orders 
    WHERE created_at > $1
    ORDER BY id DESC
`, time.Now().AddDate(0, -1, 0))

if err != nil {
    if err == context.DeadlineExceeded {
        log.Println("Query took too long, consider optimizing or adding index")
    }
    return err
}

// For very long queries, use channels:
func longRunningQuery(db *sql.DB) <-chan QueryResult {
    results := make(chan QueryResult)
    
    go func() {
        defer close(results)
        
        rows, err := db.Query("SELECT * FROM massive_table")
        if err != nil {
            results <- QueryResult{Error: err}
            return
        }
        defer rows.Close()
        
        for rows.Next() {
            var data Data
            rows.Scan(&data.ID, &data.Value)
            results <- QueryResult{Data: data}
        }
    }()
    
    return results
}

// Use with timeout:
func processWithTimeout() {
    resultsChan := longRunningQuery(db)
    timeout := time.After(1 * time.Minute)
    
    for {
        select {
        case result, ok := <-resultsChan:
            if !ok {
                return  // Channel closed
            }
            processResult(result)
        case <-timeout:
            log.Println("Query timeout!")
            return
        }
    }
}

// Pagination for large datasets:
func processLargeDataset(db *sql.DB) error {
    pageSize := 1000
    offset := 0
    
    for {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        
        rows, err := db.QueryContext(ctx, 
            "SELECT * FROM large_table LIMIT $1 OFFSET $2",
            pageSize, offset)
        
        cancel()
        
        if err != nil {
            return err
        }
        
        count := 0
        for rows.Next() {
            // Process row
            count++
        }
        rows.Close()
        
        if count < pageSize {
            break  // Last page
        }
        
        offset += pageSize
    }
    
    return nil
}
```

---

### Question 179: How do you write unit tests for code that interacts with the DB?

**Answer:**
Use mocks, interfaces, and test databases:

**Method 1: Interface mocking:**
```go
// Define interface:
type UserRepository interface {
    GetUser(id int) (*User, error)
    CreateUser(user *User) error
}

// Real implementation:
type SQLUserRepo struct {
    db *sql.DB
}

func (r *SQLUserRepo) GetUser(id int) (*User, error) {
    // Real database code
}

// Mock implementation for tests:
type MockUserRepo struct {
    users map[int]*User
}

func (m *MockUserRepo) GetUser(id int) (*User, error) {
    user, ok := m.users[id]
    if !ok {
        return nil, errors.New("not found")
    }
    return user, nil
}

func (m *MockUserRepo) CreateUser(user *User) error {
    m.users[user.ID] = user
    return nil
}

// Test:
func TestUserService(t *testing.T) {
    mockRepo := &MockUserRepo{
        users: map[int]*User{
            1: {ID: 1, Name: "John"},
        },
    }
    
    service := NewUserService(mockRepo)
    
    user, err := service.GetUser(1)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    
    if user.Name != "John" {
        t.Errorf("Expected John, got %s", user.Name)
    }
}
```

**Method 2: Using sqlmock:**
```go
import "github.com/DATA-DOG/go-sqlmock"

func TestGetUser(t *testing.T) {
    // Create mock database:
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock: %v", err)
    }
    defer db.Close()
    
    // Set expectations:
    rows := sqlmock.NewRows([]string{"id", "name", "email"}).
        AddRow(1, "John", "john@example.com")
    
    mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
        WithArgs(1).
        WillReturnRows(rows)
    
    // Test:
    repo := NewUserRepository(db)
    user, err := repo.GetUser(1)
    
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    if user.Name != "John" {
        t.Errorf("Expected John, got %s", user.Name)
    }
    
    // Verify expectations:
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Unfulfilled expectations: %v", err)
    }
}
```

**Method 3: Test database:**
```go
func setupTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("postgres", "postgresql://localhost/test_db")
    if err != nil {
        t.Fatal(err)
    }
    
    // Run migrations:
    db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100)
    )`)
    
    t.Cleanup(func() {
        db.Exec("TRUNCATE TABLE users")
        db.Close()
    })
    
    return db
}

func TestUserRepositoryIntegration(t *testing.T) {
    db := setupTestDB(t)
    repo := NewUserRepository(db)
    
    user := &User{Name: "John", Email: "john@example.com"}
    err := repo.CreateUser(user)
    
    if err != nil {
        t.Fatalf("Error creating user: %v", err)
    }
    
    retrieved, err := repo.GetUser(user.ID)
    if retrieved.Name != "John" {
        t.Errorf("Expected John, got %s", retrieved.Name)
    }
}
```

---

### Question 180: What is connection string format for different databases?

**Answer:**

**PostgreSQL:**
```go
// Standard format:
connStr := "host=localhost port=5432 user=postgres password=secret dbname=mydb sslmode=disable"

// URL format:
connStr := "postgresql://postgres:secret@localhost:5432/mydb?sslmode=disable"

// With connection pool settings:
db, _ := sql.Open("postgres", connStr)
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
```

**MySQL:**
```go
// Format: user:password@tcp(host:port)/dbname
connStr := "user:password@tcp(localhost:3306)/mydb?parseTime=true"

// With charset:
connStr := "user:password@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=true"
```

**SQLite:**
```go
// File-based:
db, _ := sql.Open("sqlite3", "./mydb.db")

// In-memory:
db, _ := sql.Open("sqlite3", ":memory:")

// With options:
db, _ := sql.Open("sqlite3", "file:mydb.db?cache=shared&mode=rwc")
```

**MongoDB:**
```go
import "go.mongodb.org/mongo-driver/mongo"

// Standard:
client, _ := mongo.Connect(ctx, options.Client().
    ApplyURI("mongodb://localhost:27017"))

// With auth:
uri := "mongodb://username:password@localhost:27017/mydb?authSource=admin"
client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))

// Replica set:
uri := "mongodb://host1:27017,host2:27017,host3:27017/mydb?replicaSet=rs0"
```

**Redis:**
```go
import "github.com/go-redis/redis/v8"

rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",  // no password
    DB:       0,   // default DB
})

// With password:
rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "secret",
    DB:       0,
})

// Cluster:
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002"},
})
```

---

*[Questions 181-240 will be added in the next batch to manage file size and ensure quality]*
## ðŸ”´ Tools, Testing, CI/CD, Ecosystem (Questions 181-230)

### Question 181: What is go vet and what does it catch?

**Answer:**
`go vet` examines Go source code and reports suspicious constructs:

```bash
go vet ./...
```

**What it catches:**
1. **Incorrect Printf formats:**
   ```go
   fmt.Printf("%d", "string")  // Caught: format %d expects int
   ```

2. **Unreachable code:**
   ```go
   return
   fmt.Println("This will never run")  // Caught
   ```

3. **Incorrect struct tags:**
   ```go
   type User struct {
       Name string `json:"name"extra"`  // Caught: malformed tag
   }
   ```

4. **Copying locks:**
   ```go
   var mu sync.Mutex
   mu2 := mu  // Caught: copying lock value
   ```

5. **Invalid method signatures:**
   ```go
   func (t *T) String(x int) string {  // Caught: should be String() string
       return ""
   }
   ```

**Usage in CI/CD:**
```yaml
# .github/workflows/test.yml
- name: Run go vet
  run: go vet ./...
```

---

### Question 182: How does go fmt help maintain code quality?

**Answer:**
`go fmt` automatically formats Go code to follow standard style:

```bash
# Format all files in current directory:
go fmt ./...

# Format specific file:
go fmt main.go

# Check what would be formatted (dry run):
gofmt -l .
```

**What it does:**
- Fixes indentation
- Aligns code blocks  
- Standardizes spacing
- Organizes imports

**Example:**
```go
// Before:
func main( ){
if   true{
fmt.Println( "hello" )
}
}

// After go fmt:
func main() {
    if true {
        fmt.Println("hello")
    }
}
```

**Benefits:**
- Eliminates style debates
- Makes code reviews easier
- Consistent codebase
- No configuration needed

**IDE Integration:**
```json
// VS Code settings.json
{
    "go.formatTool": "gofmt",
    "[go]": {
        "editor.formatOnSave": true
    }
}
```

---

### Question 183: What is golangci-lint?

**Answer:**
golangci-lint runs multiple linters in parallel:

```bash
# Install:
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run:
golangci-lint run

# Run specific linters:
golangci-lint run --enable=golint,errcheck,staticcheck
```

**Configuration (.golangci.yml):**
```yaml
linters:
  enable:
    - gofmt
    - govet
    - errcheck
    - staticcheck
    - unused
    - gosimple
    - structcheck
    - varcheck
    - ineffassign
    - deadcode
    
linters-settings:
  govet:
    check-shadowing: true
  gocyclo:
    min-complexity: 15
    
run:
  timeout: 5m
  tests: true
```

**Built-in linters:**
- **errcheck** - Checks for unchecked errors
- **staticcheck** - Advanced static analysis
- **unused** - Finds unused code
- **gosimple** - Simplification suggestions
- **govet** - Go vet checks
- **ineffassign** - Ineffectual assignments

---

### Question 184: What is the difference between go run, go build, and go install?

**Answer:**

**go run** - Compiles and runs immediately (no binary saved):
```bash
go run main.go
# Useful for: Quick testing, scripts
```

**go build** - Compiles and creates binary in current directory:
```bash
go build -o myapp
# Creates: ./myapp
# Useful for: Local testing, deployment packages
```

**go install** - Compiles and installs binary to $GOPATH/bin:
```bash
go install
# Creates: $GOPATH/bin/myapp
# Useful for: CLI tools, globally available commands
```

**Comparison:**
```go
// main.go
package main
import "fmt"
func main() {
    fmt.Println("Hello")
}
```

```bash
# go run - runs immediately:
go run main.go
# Output: Hello
# No binary created

# go build - creates binary:
go build -o hello
./hello
# Output: Hello
# Binary: ./hello

# go install - installs globally:
go install
$GOPATH/bin/hello
# Output: Hello
# Binary: $GOPATH/bin/hello
```

**Build flags:**
```bash
# Optimize for size:
go build -ldflags="-s -w" -o app

# Cross-compile:
GOOS=linux GOARCH=amd64 go build -o app-linux

# Build with race detector:
go build -race -o app
```

---

### Question 185: How does go generate work?

**Answer:**
`go generate` runs commands specified in source files:

**Usage:**
```go
// math.go
package math

//go:generate stringer -type=Operation
type Operation int

const (
    Add Operation = iota
    Subtract
    Multiply
    Divide
)
```

```bash
# Run code generation:
go generate ./...

# This executes:
# stringer -type=Operation
# Which generates: operation_string.go
```

**Common use cases:**

**1. Generate mocks:**
```go
//go:generate mockgen -source=interface.go -destination=mocks/mock_interface.go
type UserService interface {
    GetUser(id int) (*User, error)
}
```

**2. Generate protobuf:**
```go
//go:generate protoc --go_out=. --go_grpc_out=. user.proto
```

**3. Embed files:**
```go  
//go:generate go run assets/generate.go

//go:embed templates/*.html
var templates embed.FS
```

**4. Generate enums:**
```go
//go:generate enumer -type=Status -json
type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)
```

---

### Question 186: What is a build constraint?

**Answer:**
Build constraints (build tags) control which files are compiled:

**Syntax:**
```go
//go:build linux
// +build linux

package main

func platformSpecific() {
    // Linux-only code
}
```

**Multiple tags:**
```go
//go:build linux && amd64
// Linux AMD64 only

//go:build darwin || linux  
// macOS or Linux

//go:build !windows
// Everything except Windows
```

**Use cases:**

**1. Platform-specific code:**
```go
// file_unix.go
//go:build unix

func openFile() {
    // Unix implementation
}

// file_windows.go
//go:build windows

func openFile() {
    // Windows implementation
}
```

**2. Feature flags:**
```go
//go:build debug

func enableDebugLogging() {
    // Debug-only code
}
```

```bash
# Build with tag:
go build -tags debug
```

**3. Integration tests:**
```go
//go:build integration

func TestDatabaseIntegration(t *testing.T) {
    // Only runs with: go test -tags integration
}
```

---

###Question 187: How do you write tests in Go?

**Answer:**
Create `_test.go` files with Test functions:

**Basic test:**
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go  
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

**Table-driven tests:**
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive numbers", 2, 3, 5},
        {"with zero", 0, 5, 5},
        {"negative numbers", -2, -3, -5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

**Subtests:**
```go
func TestUser(t *testing.T) {
    t.Run("Creation", func(t *testing.T) {
        user := NewUser("John")
        if user.Name != "John" {
            t.Error("Name not set correctly")
        }
    })
    
    t.Run("Validation", func(t *testing.T) {
        user := NewUser("")
        if user.IsValid() {
            t.Error("Empty name should be invalid")
        }
    })
}
```

**Run tests:**
```bash
go test                    # Current package
go test ./...              # All packages
go test -v                 # Verbose
go test -run TestAdd       # Specific test
go test -cover             # With coverage
```

---

### Question 188: How do you test for expected panics?

**Answer:**
Use `defer` with `recover`:

```go
func TestPanicRecovery(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("Expected panic, but didn't panic")
        }
    }()
    
    // Code that should panic:
    causePanic()
}

// Better approach with helper:
func assertPanic(t *testing.T, f func()) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("Expected panic")
        }
    }()
    f()
}

func TestDivideByZero(t *testing.T) {
    assertPanic(t, func() {
        Divide(10, 0)  // Should panic
    })
}

// Using testify package:
import "github.com/stretchr/testify/assert"

func TestPanic(t *testing.T) {
    assert.Panics(t, func() {
        causePanic()
    }, "Function should panic")
    
    assert.NotPanics(t, func() {
        safeFunction()
    }, "Function should not panic")
}
```

---

### Question 189: What are mocks and how do you use them in Go?

**Answer:**
Mocks simulate dependencies for testing:

**Manual mocks:**
```go
// Interface to mock:
type EmailSender interface {
    Send(to, subject, body string) error
}

// Mock implementation:
type MockEmailSender struct {
    SendCalled bool
    SendError  error
}

func (m *MockEmailSender) Send(to, subject, body string) error {
    m.SendCalled = true
    return m.SendError
}

// Test using mock:
func TestUserService(t *testing.T) {
    mockEmail := &MockEmailSender{}
    service := NewUserService(mockEmail)
    
    service.RegisterUser("john@example.com")
    
    if !mockEmail.SendCalled {
        t.Error("Expected Email.Send to be called")
    }
}
```

**Using testify/mock:**
```go
import "github.com/stretchr/testify/mock"

type MockEmailSender struct {
    mock.Mock
}

func (m *MockEmailSender) Send(to, subject, body string) error {
    args := m.Called(to, subject, body)
    return args.Error(0)
}

func TestWithTestify(t *testing.T) {
    mockEmail := new(MockEmailSender)
    
    // Set expectations:
    mockEmail.On("Send", "john@example.com", "Welcome", mock.Anything).
        Return(nil)
    
    service := NewUserService(mockEmail)
    service.RegisterUser("john@example.com")
    
    // Assert expectations were met:
    mockEmail.AssertExpectations(t)
}
```

**Using mockgen (gomock):**
```bash
go install github.com/golang/mock/mockgen@latest

mockgen -source=interface.go -destination=mocks/mock_interface.go
```

```go
//go:generate mockgen -source=user_service.go -destination=mocks/mock_user_service.go
type UserService interface {
    GetUser(id int) (*User, error)
}

// In tests:
func TestWithGomock(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockService := mocks.NewMockUserService(ctrl)
    mockService.EXPECT().
        GetUser(1).
        Return(&User{ID: 1, Name: "John"}, nil)
    
    user, _ := mockService.GetUser(1)
    if user.Name != "John" {
        t.Error("Expected John")
    }
}
```

---

### Question 190: How do you use the testing and testify packages?

**Answer:**

**Standard testing package:**
```go
import "testing"

func TestBasic(t *testing.T) {
    t.Log("Starting test")
    t.Error("This test failed")
    t.Skip("Skipping this test")
    t.Fatal("Stop execution")
}
```

**testify/assert - Convenient assertions:**
```go
import "github.com/stretchr/testify/assert"

func TestWithAssert(t *testing.T) {
    // Equality:
    assert.Equal(t, 5, result)
    assert.NotEqual(t, 0, count)
    
    // Nil checks:
    assert.Nil(t, err)
    assert.NotNil(t, user)
    
    // Boolean:
    assert.True(t, condition)
    assert.False(t, flag)
    
    // Strings:
    assert.Contains(t, "hello world", "world")
    assert.Empty(t, "")
    
    // Collections:
    assert.Len(t, list, 5)
    assert.ElementsMatch(t, []int{1, 2, 3}, result)
    
    // Errors:
    assert.NoError(t, err)
    assert.Error(t, err)
}
```

**testify/require - Fail fast:  **
```go
import "github.com/stretchr/testify/require"

func TestWithRequire(t *testing.T) {
    user, err := GetUser(1)
    require.NoError(t, err)  // Stops if error
    require.NotNil(t, user)  // Stops if nil
    
    // Continue only if above passed:
    assert.Equal(t, "John", user.Name)
}
```

**testify/suite - Test suites:**
```go
import "github.com/stretchr/testify/suite"

type UserTestSuite struct {
    suite.Suite
    db *sql.DB
}

func (s *UserTestSuite) SetupSuite() {
    // Runs once before all tests
    s.db = setupDB()
}

func (s *UserTestSuite) TearDownSuite() {
    // Runs once after all tests
    s.db.Close()
}

func (s *UserTestSuite) SetupTest() {
    // Runs before each test
    s.db.Exec("TRUNCATE users")
}

func (s *UserTestSuite) TestCreateUser() {
    user := &User{Name: "John"}
    err := CreateUser(s.db, user)
    
    s.NoError(err)
    s.Greater(user.ID, 0)
}

func TestUserSuite(t *testing.T) {
    suite.Run(t, new(UserTestSuite))
}
```

---

### Question 191: How do you structure test files in Go?

**Answer:**

**Project structure:**
```
myproject/
â”œâ”€â”€ user/
â”‚   â”œâ”€â”€ user.go           # Source code
â”‚   â”œâ”€â”€ user_test.go      # Unit tests (same package)
â”‚   â””â”€â”€ user_integration_test.go  # Integration tests
â”œâ”€â”€ testdata/             # Test fixtures
â”‚   â”œâ”€â”€ users.json
â”‚   â””â”€â”€ config.yaml
â””â”€â”€ tests/
    â””â”€â”€ e2e_test.go       # End-to-end tests (separate package)
```

**File naming:**
```
user.go              # Implementation
user_test.go         # Tests (package user)
user_internal_test.go  # Tests (package user_test)
```

**Same package vs separate:**
```go
// user_test.go (same package - can test private members)
package user

func TestInternalMethod(t *testing.T) {
    u := &User{}
    u.internalMethod()  // Can access private members
}

// user_external_test.go (separate package - only public API)
package user_test

import "myproject/user"

func TestPublicAPI(t *testing.T) {
    u := user.NewUser()  // Only public members
}
```

**Test organization:**
```go
// unit_test.go - Fast unit tests
func TestAdd(t *testing.T) { }
func TestSubtract(t *testing.T) { }

// integration_test.go - Slower integration tests
//go:build integration

func TestDatabaseIntegration(t *testing.T) { }

// benchmark_test.go - Benchmarks
func BenchmarkAdd(b *testing.B) { }
```

**Shared test helpers:**
```go
// testing.go
package user

// Test helpers (not exported):
func createTestUser(t *testing.T) *User {
    t.Helper()  // Mark as helper
    return &User{Name: "Test User"}
}

// test_fixtures.go
func loadTestData(t *testing.T, filename string) []byte {
    data, err := os.ReadFile(filepath.Join("testdata", filename))
    if err != nil {
        t.Fatal(err)
    }
    return data
}
```

---

### Question 192: What is a benchmark test?

**Answer:**
Benchmark tests measure performance:

**Basic benchmark:**
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

**Run benchmarks:**
```bash
go test -bench=.                  # Run all benchmarks
go test -bench=Add               # Specific benchmark
go test -bench=. -benchmem       # With memory stats
go test -bench=. -cpuprofile=cpu.prof  # CPU profile
```

**More examples:**
```go
// Table-driven benchmarks:
func BenchmarkStringOperations(b *testing.B) {
    tests := []struct{
        name string
        fn   func(string) string
    }{
        {"concat", func(s string) string { return s + s }},
        {"builder", func(s string) string {
            var builder strings.Builder
            builder.WriteString(s)
            builder.WriteString(s)
            return builder.String()
        }},
    }
    
    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                tt.fn("hello")
            }
        })
    }
}

// Reset timer for setup:
func BenchmarkWithSetup(b *testing.B) {
    // Expensive setup:
    data := generateLargeDataset()
    
    b.ResetTimer()  // Don't count setup time
    
    for i := 0; i < b.N; i++ {
        process(data)
    }
}

// Parallel benchmarks:
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // Work that can run in parallel
            expensiveOperation()
        }
    })
}

// Memory allocations:
func BenchmarkAllocations(b *testing.B) {
    b.ReportAllocs()  // Report memory allocations
    
    for i := 0; i < b.N; i++ {
        _ = make([]int, 1000)
    }
}
```

**Output:**
```
BenchmarkAdd-8          1000000000     0.25 ns/op
BenchmarkConcat-8         5000000      250 ns/op    48 B/op   1 allocs/op
```

---

### Question 193: How do you measure test coverage in Go?

**Answer:**
Use built-in coverage tools:

```bash
# Run tests with coverage:
go test -cover

# Generate coverage profile:
go test -coverprofile=coverage.out

# View coverage in terminal:
go tool cover -func=coverage.out

# Generate HTML report:
go tool cover -html=coverage.out

# Coverage for specific packages:
go test ./... -coverprofile=coverage.out

# Coverage modes:
go test -covermode=set       # Boolean (covered or not)
go test -covermode=count     # Count how many times
go test -covermode=atomic    # Thread-safe count
```

**Example output:**
```
math/add.go:5:  Add     100.0%
math/sub.go:5:  Sub     80.0%
total:                  90.0%
```

**CI/CD integration:**
```yaml
# .github/workflows/test.yml
- name: Run tests with coverage
  run: go test -v -coverprofile=coverage.out ./...
 
- name: Upload coverage
  uses: codecov/codecov-action@v3
  with:
    files: ./coverage.out
```

**Makefile:**
```makefile
test-coverage:
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html
    @echo "Coverage report: coverage.html"
```

---

### Question 194: How do you test concurrent functions?

**Answer:**
Testing concurrency requires special care:

**Basic concurrent test:**
```go
func TestConcurrent Counter(t *testing.T) {
    counter := NewSafeCounter()
    var wg sync.WaitGroup
    
    // Start multiple goroutines:
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    
    if counter.Value() != 100 {
        t.Errorf("Expected 100, got %d", counter.Value())
    }
}
```

**Race detector:**
```bash
go test -race ./...
```

**Testing channels:**
```go
func TestChannel(t *testing.T) {
    ch := make(chan int, 1)
    
    go func() {
        ch <- 42
    }()
    
    select {
    case val := <-ch:
        assert.Equal(t, 42, val)
    case <-time.After(time.Second):
        t.Fatal("Timeout waiting for value")
    }
}
```

**Testing goroutine leaks:**
```go
import "go.uber.org/goleak"

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}

func TestNoGoroutineLeak(t *testing.T) {
    defer goleak.VerifyNone(t)
    
    // Code that should not leak goroutines
    startWorker()
}
```

**Stress testing:**
```go
func TestConcurrentWrites(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping stress test in short mode")
    }
    
    m := &sync.Map{}
    done := make(chan bool)
    
    // Multiple writers:
    for i := 0; i < 100; i++ {
        go func(id int) {
            for j := 0; j < 1000; j++ {
                m.Store(fmt.Sprintf("%d-%d", id, j), id)
            }
            done <- true
        }(i)
    }
    
    // Wait for all:
    for i := 0; i < 100; i++ {
        <-done
    }
}
```

---

### Question 195: What is a race detector and how do you use it?

**Answer:**
The race detector finds data races at runtime:

**Usage:**
```bash
# Run tests with race detector:
go test -race ./...

# Run program with race detector:
go run -race main.go

# Build with race detector:
go build -race -o myapp
```

**Example race condition:**
```go
// BAD - Race condition:
var counter int

func increment() {
    counter++  // RACE!
}

func TestRace(t *testing.T) {
    for i := 0; i < 1000; i++ {
        go increment()
    }
    time.Sleep(time.Second)
}
```

**Running with race detector:**
```bash
$ go test -race
==================
WARNING: DATA RACE
Write at 0x00c000018090 by goroutine 8:
  main.increment()
      /path/to/file.go:5 +0x44

Previous read at 0x00c000018090 by goroutine 7:
  main.increment()
      /path/to/file.go:5 +0x3a
==================
```

**Fixed version:**
```go
// GOOD - No race:
var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

// Or use atomic:
var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}
```

**Common race patterns:**
```go
// 1. Map races:
m := make(map[string]int)
go func() { m["key"] = 1 }()
go func() { m["key"] = 2 }()  // RACE!

// Fix: Use sync.Map or mutex

// 2. Slice races:
s := []int{}
go func() { s = append(s, 1) }()
go func() { s = append(s, 2) }()  // RACE!

// Fix: Use channels or mutex

// 3. Closure variable:
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // RACE with loop variable!
    }()
}

// Fix: Pass as parameter:
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println(n)
    }(i)
}
```

---

### Question 196: What is go.mod and go.sum?

**Answer:**

**go.mod** - Defines module and dependencies:
```go
module github.com/user/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/lib/pq v1.10.9
)

require (
    // Indirect dependencies (transitive)
    github.com/golang/protobuf v1.5.3 // indirect
)

replace github.com/old/package => github.com/new/package v1.0.0

exclude github.com/bad/package v1.2.3
```

**go.sum** - Checksums for dependency verification:
```
github.com/gin-gonic/gin v1.9.0 h1:abc123...
github.com/gin-gonic/gin v1.9.0/go.mod h1:def456...
```

**Commands:**
```bash
# Initialize module:
go mod init github.com/user/myproject

# Add missing dependencies:
go mod tidy

# Download dependencies:
go mod download

# Verify dependencies:
go mod verify

# Show dependency graph:
go mod graph

# Update dependency:
go get -u github.com/gin-gonic/gin

# Update all:
go get -u ./...

# Vendor dependencies:
go mod vendor
```

**Version selection:**
```bash
# Specific version:
go get github.com/gin-gonic/gin@v1.9.0

# Latest:
go get github.com/gin-gonic/gin@latest

# Specific commit:
go get github.com/gin-gonic/gin@abc1234

# Specific branch:
go get github.com/gin-gonic/gin@master
```

---

### Question 197: How does semantic versioning work in Go modules?

**Answer:**
Go uses semantic versioning (vX.Y.Z):

**Version format:**
- v1.2.3
  - Major (1): Breaking changes
  - Minor (2): New features (backward compatible)
  - Patch (3): Bug fixes

**Version ranges:**
```bash
# Latest v1:
go get github.com/pkg/errors@v1

# At least v1.2:
go get github.com/pkg/errors@>=v1.2.0

# Before v2:
go get github.com/pkg/errors@<v2.0.0
```

**Major version in import path:**
```go
// v0 and v1 - no version in path:
import "github.com/user/package"

// v2+ - version in path:
import "github.com/user/package/v2"
import "github.com/user/package/v3"
```

**Module path for v2+:**
```go
// go.mod
module github.com/user/package/v2

// Import in code:
import "github.com/user/package/v2"
```

**Minimal version selection:**
Go uses the minimum required version that satisfies all constraints:
```
A requires B v1.2
C requires B v1.3

Go selects: B v1.3 (minimum that satisfies both)
```

**Pre-release versions:**
```bash
v1.2.3-alpha.1
v1.2.3-beta.2
v1.2.3-rc.1
```

**Pseudo-versions (commits without tags):**
```
v0.0.0-20230101120000-abc1234567ab
```

---

### Question 198: How to build and deploy a Go binary to production?

**Answer:**

**Build for production:**
```bash
# Standard build:
go build -o myapp

# Optimized build (smaller binary):
go build -ldflags="-s -w" -o myapp
# -s: strip symbol table
# -w: strip DWARF debug info

# With version info:
VERSION=$(git describe --tags)
go build -ldflags="-X main.Version=$VERSION" -o myapp

# Cross-compile:
GOOS=linux GOARCH=amd64 go build -o myapp-linux
GOOS=windows GOARCH=amd64 go build -o myapp.exe
GOOS=darwin GOARCH=arm64 go build -o myapp-mac
```

**main.go with version:**
```go
package main

var (
    Version   = "dev"
    BuildTime = "unknown"
    GitCommit = "unknown"
)

func main() {
    fmt.Printf("Version: %s\n", Version)
    fmt.Printf("Built: %s\n", BuildTime)
    fmt.Printf("Commit: %s\n", GitCommit)
    
    // Your app code...
}
```

**Build script:**
```bash
#!/bin/bash
VERSION=$(git describe --tags --always)
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(git rev-parse HEAD)

go build -ldflags="\
    -X main.Version=$VERSION \
    -X main.BuildTime=$BUILD_TIME \
    -X main.GitCommit=$GIT_COMMIT \
    -s -w" \
    -o bin/myapp
```

**Makefile:**
```makefile
VERSION := $(shell git describe --tags --always)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -s -w"

.PHONY: build
build:
    go build $(LDFLAGS) -o bin/myapp

.PHONY: build-linux
build-linux:
    GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-linux

.PHONY: build-all
build-all:
    GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-linux
    GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp-mac
    GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/myapp.exe
```

**Docker deployment:**
```dockerfile
# Multi-stage build
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app

FROM scratch
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
```

---

### Question 199: What tools are used for Dockerizing Go apps?

**Answer:**

**Multi-stage Dockerfile:**
```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/main .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
```

**Using scratch (minimal image):**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app .

FROM scratch
COPY --from=builder /build/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT ["/app"]
```

**Docker Compose:**
```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://localhost/mydb
    depends_on:
      - db
  
  db:
    image: postgres:15-alpine
    environment:
      POSTGRES_PASSWORD: secret
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:
```

**.dockerignore:**
```
.git
.gitignore
README.md
Dockerfile
docker-compose.yml
*.md
.env
tmp/
```

**Build and run:**
```bash
# Build:
docker build -t myapp:latest .

# Run:
docker run -p 8080:8080 my app:latest

# With env vars:
docker run -p 8080:8080 -e DATABASE_URL=... myapp:latest

# Docker Compose:
docker-compose up --build
```

---

### Question 200: How do you set up a CI/CD pipeline for a Go project?

**Answer:**

**GitHub Actions (.github/workflows/ci.yml):**
```yaml
name: CI/CD

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Install dependencies
        run: go mod download
      
      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...
      
      - name: Run go vet
        run: go vet ./...
      
      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: latest
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          files: ./coverage.out
  
  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Build
        run: go build -v -o bin/myapp .
      
      - name: Upload artifact
        uses: actions/upload-artifact@v3
        with:
          name: myapp
          path: bin/myapp
  
  deploy:
    needs: build
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - name: Download artifact
        uses: actions/download-artifact@v3
        with:
          name: myapp
      
      - name: Deploy to server
        run: |
          # Your deployment script
          echo "Deploying..."
```

**GitLab CI (.gitlab-ci.yml):**
```yaml
stages:
  - test
  - build
  - deploy

test:
  stage: test
  image: golang:1.21
  script:
    - go mod download
    - go test -v -race -coverprofile=coverage.out ./...
    - go vet ./...
  coverage: '/coverage: \d+.\d+% of statements/'

build:
  stage: build
  image: golang:1.21
  script:
    - go build -o bin/myapp .
  artifacts:
    paths:
      - bin/myapp

deploy:
  stage: deploy
  only:
    - main
  script:
    - echo "Deploying to production..."
```

**Makefile for CI:**
```makefile
.PHONY: ci
ci: lint test build

.PHONY: lint
lint:
    golangci-lint run ./...

.PHONY: test
test:
    go test -v -race -coverprofile=coverage.out ./...

.PHONY: build
build:
    go build -o bin/myapp .

.PHONY: coverage
coverage:
    go tool cover -html=coverage.out
```

---

*[Questions 201-280 will be added in the next batch]*
