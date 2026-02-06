# Golang OOPs Concepts Cheatsheet

Go is not a traditional Object-Oriented language (no `class` keyword), but it supports OOP features through **Structs**, **Methods**, and **Interfaces**.

---

## 🟦 Structs (The "Class" Alternative)

In Go, a `struct` replaces the class. It defines the state (data fields).

### Definition & Instantiation
```go
// Define a struct (like a class definition)
type Employee struct {
    ID   int
    Name string
    Role string
}

// Instantiate
func main() {
    // 1. Literal style
    e1 := Employee{ID: 1, Name: "Alice", Role: "Dev"}

    // 2. Key-value style (Recommended)
    e2 := Employee{
        Name: "Bob",
        Role: "Manager", // ID is 0 (default)
    }

    // 3. New keyword (Returns pointer)
    e3 := new(Employee) // *Employee
    e3.Name = "Charlie"
}
```

### Constructor Pattern
Go doesn't have constructors. We use factory functions, typically named `New<Type>`.

```go
func NewEmployee(name string, role string) *Employee {
    return &Employee{
        ID:   rand.Int(),
        Name: name,
        Role: role,
    }
}

// Usage
emp := NewEmployee("Dave", "Lead")
```

---

## 🟧 Methods (Behavior)

Methods are functions attached to a specific type (receiver).

### Value vs. Pointer Receivers
* **Value Receiver**: Operates on a *copy*. Cannot modify original state.
* **Pointer Receiver**: Operates on the *actual* object. Can modify state.

```go
type Counter struct {
    Value int
}

// Value Receiver (Safe, Read-only logic)
func (c Counter) Get() int {
    // c.Value = 100 // This would only change the local copy
    return c.Value
}

// Pointer Receiver (Mutates state)
func (c *Counter) Increment() {
    c.Value++
}

func main() {
    c := Counter{}
    c.Increment()      // Go automatically converts &c
    fmt.Println(c.Get()) // Output: 1
}
```
**Rule of Thumb:**
- Use **Pointer Receiver** if you need to modify the struct OR if the struct is large (to avoid copying).
- Use **Value Receiver** for small, immutable types.

---

## 🔒 Encapsulation

Go uses **capitalization** to control visibility (Public/Private), not keywords like `public` or `private`.

| Naming          | Visibility               | Scope                                   |
| :-------------- | :----------------------- | :-------------------------------------- |
| **`TitleCase`** | **Public** (Exported)    | Accessible from other packages          |
| **`camelCase`** | **Private** (Unexported) | Accessible ONLY within the same package |
|                 |                          |                                         |

```go
package users

type Account struct {
    Owner    string // Exported (Public)
    balance  int    // Unexported (Private)
}

// Setter (Public method to access private field)
func (a *Account) Deposit(amount int) {
    if amount > 0 {
        a.balance += amount
    }
}

// Getter
func (a *Account) Balance() int {
    return a.balance
}
```

[^1]
## 🔗 Composition (Inheritance Alternative)

Go prefers **Composition over Inheritance**. It uses **Struct Embedding**.

### Type Embedding
By embedding a struct anonymously, the inner struct's fields and methods are **promoted** to the outer struct.

```go
type Animal struct {
    Species string
}

func (a *Animal) Move() {
    fmt.Println(a.Species, "is moving")
}

type Dog struct {
    Animal // Embedded Struct (Is-a relationship)
    Breed  string
}

func main() {
    d := Dog{
        Animal: Animal{Species: "Canine"},
        Breed:  "Labrador",
    }

    // Direct access to promoted fields/methods
    d.Move()           // Works! Prints "Canine is moving"
    fmt.Println(d.Species) // Works!
}
```

### Overriding Methods (Shadowing)
If `Dog` defines its own `Move()` method, it shadows the embedded `Animal.Move()`.

```go
func (d *Dog) Move() {
    fmt.Println("Dog runs fast!")
}

// d.Move() // Calls Dog's Move
// d.Animal.Move() // Explicitly calls embedded Move
```

---

## 🟣 Polymorphism (Interfaces)

Interfaces define **behavior** (method signatures). Types implement interfaces **implicitly** just by having the methods.

### Definition
```go
type Flyer interface {
    Fly(distance int)
}
```

### Implementation (Implicit)
```go
type Bird struct{ Name string }

// Implicitly implements Flyer because signature matches
func (b Bird) Fly(dist int) {
    fmt.Println(b.Name, "flew", dist, "miles")
}

type Plane struct{ Model string }

func (p Plane) Fly(dist int) {
    fmt.Println(p.Model, "flew", dist, "miles")
}
```

### Usage (Polymorphism)
```go
func Travel(f Flyer) {
    f.Fly(100)
}

func main() {
    b := Bird{"Eagle"}
    p := Plane{"Boeing 747"}

    // Both work because they implement Flyer
    Travel(b)
    Travel(p)
}
```

### Interface Composition
Interfaces can be embedded into other interfaces.

```go
type Reader interface { Read(b []byte) (n int, err error) }
type Writer interface { Write(b []byte) (n int, err error) }

// ReadWriter embeds both
type ReadWriter interface {
    Reader
    Writer
}
```

---

## ⚡ Type Assertions & Switches

When working with interfaces (especially empty interface `interface{}`), you may need to retrieve the underlying concrete type.

### Type Assertion
```go
var i interface{} = "hello"

s, ok := i.(string) // Check if 'i' is string
if ok {
    fmt.Println("It's a string:", s)
}

r, ok := i.(float64) // False, r is 0
```

### Type Switch
```go
func checkType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    case Flyer:
        fmt.Println("Implements Flyer")
    default:
        fmt.Println("Unknown type")
    }
}
```

---

## 🔥 Advanced / Tricky OOP Patterns

### 1. Method Sets (Pointer vs Value)
Pointer receivers are NOT automatically satisfied by values (unless addressable).

```go
type Shape interface { Area() int }
type Square struct { L int }

func (s *Square) Area() int { return s.L * s.L } // Pointer receiver!

func main() {
    var s Shape
    sq := Square{5}
    // s = sq      // ❌ Compile Error! 'sq' is value, method requires ptr
    s = &sq       // ✅ OK!
}
```

### 2. The "Nil" Interface Trap
A nil pointer to a struct is **NOT** a nil interface.

```go
func getError() error {
    var err *MyError = nil // Pointer is nil
    return err             // Interface is NOT nil (contains {Type: *MyError, Value: nil})
}

func main() {
    err := getError()
    if err != nil {
        fmt.Println("Error exists!") // Prints this! (Trap)
    }
}
```
**Fix:** Always return explicit `nil`.

### 3. Functional Options Pattern
Go's alternative to Builder pattern / constructors with many args.

```go
type Server struct {
    Port int
    Host string
}

type Option func(*Server)

func WithPort(p int) Option {
    return func(s *Server) { s.Port = p }
}

func NewServer(opts ...Option) *Server {
    srv := &Server{Host: "localhost", Port: 8080} // Defaults
    for _, opt := range opts {
        opt(srv)
    }
    return srv
}

// Usage
s := NewServer(WithPort(9000))
```

### 4. Empty Struct (`struct{}`)
Zero memory allocation. Used for sets types or signal channels.

```go
// 1. Set (Map with no values)
set := make(map[string]struct{})
set["item"] = struct{}{}

// 2. Signal Channel
done := make(chan struct{})
close(done) // Broadcast signal
```

### 5. Type Alias vs Definition
```go
type MyInt int      // New Type (No methods of int)
type IntAlias = int // Alias (Same methods as int)
```

[1]: To turn your snippet into a full, working program, you need to organize it into two files. In Go, access control (public vs. private) is enforced at the **package level**.

	To see the "private" behavior in action, we’ll create a `users` package and a `main` package to interact with it.

	### 1. The Package File

	Create a folder named `users` and place this code inside a file named `account.go`.

	Go

	```
	package users

	type Account struct {
		Owner   string // Exported: Accessible from any package
		balance int    // Unexported: Only accessible within the 'users' package
	}

	// NewAccount is a constructor function to initialize the struct
	func NewAccount(owner string) *Account {
		return &Account{
			Owner:   owner,
			balance: 0,
		}
	}

	// Deposit is a public method (Setter)
	func (a *Account) Deposit(amount int) {
		if amount > 0 {
			a.balance += amount
		}
	}

	// Balance is a public method (Getter)
	func (a *Account) Balance() int {
		return a.balance
	}
	```

	---

	### 2. The Main Program

	Create a file named `main.go` in your **root directory** (outside the users folder). This is where we’ll use the code.

	Go

	```
	package main

	import (
		"fmt"
		"your-module-name/users" // Replace 'your-module-name' with your actual module name
	)

	func main() {
		// 1. Create a new account using the constructor
		myAcc := users.NewAccount("Alice")

		// 2. Access a public field
		fmt.Println("Account Owner:", myAcc.Owner)

		// 3. Use public methods to modify private data
		myAcc.Deposit(500)
		fmt.Printf("Initial Balance: $%d\n", myAcc.Balance())

		// 4. Try to deposit more
		myAcc.Deposit(250)
		fmt.Printf("New Balance: $%d\n", myAcc.Balance())

		/* NOTE: If you tried to run 'myAcc.balance = 1000' here, 
		   the compiler would throw an error because 'balance' is unexported.
		*/
	}
	```

	---

	### Key Concepts Applied

	- **Encapsulation:** By keeping `balance` lowercase, you prevent other developers from accidentally setting the balance to a negative number or bypassing your logic.
	    
	- **The "Exported" Rule:** In Go, if a name starts with a **Capital letter**, it is public. If it starts with a **lowercase letter**, it is private to its package.
	    
	- **Pointer Receivers:** We use `(a *Account)` in your methods so that we are modifying the actual account data, rather than a copy of it.
	    

	### How to run this

	1. Initialize your module: `go mod init my-project`
	    
	2. Ensure your folder structure looks like this:
	    
	    - `my-project/`
	        
	        - `go.mod`
	            
	        - `main.go`
	            
	        - `users/`
	            
	            - `account.go`
	                
	3. Run the program: `go run main.go`
	    

PS G:\My Drive\All Documents\java-golang-interview-questions\practice\golang\encapsulation> go run main.go
# command-line-arguments
.\main.go:11:18: acc.balanceCount undefined (cannot refer to unexported method balanceCount)