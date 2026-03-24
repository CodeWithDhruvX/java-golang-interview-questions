# OOPs Concepts (66-75)

## 66. Singleton Pattern
**Principle**: Private constructor, static instance.
**Question**: Determine how to create a Singleton class.
**Code**:
```go
package main

import (
    "fmt"
    "sync"
)

type Singleton struct {
    name string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{name: "Singleton"}
    })
    return instance
}

func main() {
    s1 := GetInstance()
    s2 := GetInstance()
    fmt.Printf("Same instance? %t\n", s1 == s2) // true
}
```

## 67. Immutable Struct
**Principle**: Use unexported fields, only getters.
**Question**: Create an immutable struct having one field.
**Code**:
```go
package main

import "fmt"

type Immutable struct {
    name string
}

func NewImmutable(name string) *Immutable {
    return &Immutable{name: name}
}

func (i *Immutable) GetName() string {
    return i.name
}

func main() {
    imm := NewImmutable("test")
    fmt.Println(imm.GetName())
}
```

## 68. Function Overloading (Manual)
**Principle**: Use different function names or interfaces.
**Question**: Demonstrate method overloading equivalent.
**Code**:
```go
package main

import "fmt"

type MathUtil struct{}

func (m MathUtil) AddInt(a, b int) int {
    return a + b
}

func (m MathUtil) AddFloat(a, b float64) float64 {
    return a + b
}

func main() {
    util := MathUtil{}
    fmt.Println(util.AddInt(1, 2))
    fmt.Println(util.AddFloat(1.5, 2.5))
}
```

## 69. Method Overriding (Embedding)
**Principle**: Use struct embedding and method overriding.
**Question**: Demonstrate method overriding.
**Code**:
```go
package main

import "fmt"

type Parent struct{}

func (p Parent) Show() {
    fmt.Println("Parent")
}

type Child struct {
    Parent
}

func (c Child) Show() {
    fmt.Println("Child")
}

func main() {
    var p Parent = Child{Parent{}}
    p.Show() // Parent (calls Parent's method)
    
    c := Child{Parent{}}
    c.Show() // Child (calls Child's method)
}
```

## 70. Interface Implementation
**Principle**: Define contract, implement in struct.
**Question**: Implement an interface.
**Code**:
```go
package main

import "fmt"

type Animal interface {
    Sound()
}

type Dog struct{}

func (d Dog) Sound() {
    fmt.Println("Bark")
}

func main() {
    var animal Animal = Dog{}
    animal.Sound()
}
```

## 71. Abstract Behavior (Interface)
**Principle**: Use interfaces with unimplemented methods.
**Question**: Use an abstract class equivalent.
**Code**:
```go
package main

import "fmt"

type Vehicle interface {
    Drive()
}

type Car struct{}

func (c Car) Drive() {
    fmt.Println("Drive Car")
}

func main() {
    var vehicle Vehicle = Car{}
    vehicle.Drive()
}
```

## 72. Custom Error
**Principle**: Implement error interface.
**Question**: Create a custom error.
**Code**:
```go
package main

import "fmt"

type MyError struct {
    message string
}

func (e MyError) Error() string {
    return e.message
}

func main() {
    err := MyError{message: "Error"}
    fmt.Println(err.Error())
}
```

## 73. Deep Copy vs Shallow Copy
**Principle**: Shallow copies reference; Deep copies object graph.
**Question**: Demonstrate cloning (Shallow Copy).
**Code**:
```go
package main

import "fmt"

type Node struct {
    Val int
}

func (n Node) Clone() Node {
    return Node{Val: n.Val}
}

func main() {
    original := Node{Val: 10}
    cloned := original.Clone()
    fmt.Printf("Original: %d, Cloned: %d\n", original.Val, cloned.Val)
}
```

## 74. Init Function vs Constructor
**Principle**: Init runs on package load, Constructor on creation.
**Question**: Show execution order of init and constructor.
**Code**:
```go
package main

import "fmt"

type Test struct{}

func init() {
    fmt.Println("Init Function")
}

func NewTest() *Test {
    fmt.Println("Constructor")
    return &Test{}
}

func main() {
    t1 := NewTest()
    t2 := NewTest()
    _ = t1
    _ = t2
}
```

## 75. Custom Sorting
**Principle**: Implement sort.Interface for custom ordering.
**Question**: Sort slice using custom comparator.
**Code**:
```go
package main

import (
    "fmt"
    "sort"
)

type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }

func main() {
    list := StringSlice{"B", "A"}
    sort.Sort(list)
    fmt.Println(list)
}
```
