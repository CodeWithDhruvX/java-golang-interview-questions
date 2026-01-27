package main

import (
	"fmt"
)

// 91. Exception Handling Logic (Scenario)
// In Go, we use panic/recover for exceptions
func exceptionHandlingDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from exception:", r)
		}
	}()
	panic("Something went wrong!")
}

// 92. Method Overloading Logic (Not supported in Go directly, simulated)
// Go doesn't support method overloading. We use different names or interfaces.
func add(a, b int) int {
	return a + b
}
func addThree(a, b, c int) int {
	return a + b + c
}

// 93. Pass by Value vs Pass by Reference
func modifyValue(val int, ref *int) {
	val = 100
	*ref = 100
}

// 94. String Pool Logic (Go internals)
// Go strings are immutable and behave somewhat effectively, but no direct "String Pool" manipulation like Java.
func stringIdentity() {
	s1 := "Hello"
	s2 := "Hello"
	fmt.Println("s1 == s2:", s1 == s2) // True (Content comparison)
}

// 95. Switch Case Logic
func checkGrade(score int) {
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}
}

// 96. Static keyword logic (No static in Go, use package level vars)
var count int = 0

func counter() {
	count++
	fmt.Println("Count:", count)
}

// 97. Final/Const logic
func constantDemo() {
	const Pi = 3.14
	// Pi = 3.14159 // Error: cannot assign to Pi
	fmt.Println("Pi:", Pi)
}

// 98. Object Cloning (Struct copy)
type Person struct {
	Name string
	Age  int
}

func cloneDemo() {
	p1 := Person{"John", 30}
	p2 := p1 // Value copy (Shallow clone)
	p2.Name = "Doe"
	fmt.Println("p1:", p1) // John
	fmt.Println("p2:", p2) // Doe
}

// 99. Constructor Logic (Factory functions)
func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

// 100. Immutable Class Logic (Private fields with getters)
type ImmutablePoint struct {
	x, y int
}

func NewImmutablePoint(x, y int) ImmutablePoint {
	return ImmutablePoint{x: x, y: y}
}
func (p ImmutablePoint) X() int { return p.x }
func (p ImmutablePoint) Y() int { return p.y }

func main() {
	fmt.Println("91. Exception Handling:")
	exceptionHandlingDemo()

	fmt.Println("\n92. Overloading (Simulated):")
	fmt.Println("Add(1, 2):", add(1, 2))
	fmt.Println("AddThree(1, 2, 3):", addThree(1, 2, 3))

	fmt.Println("\n93. Pass by Value vs Ref:")
	v, r := 1, 1
	modifyValue(v, &r)
	fmt.Printf("Val: %d (Unchanged), Ref: %d (Changed)\n", v, r)

	fmt.Println("\n94. String Identity:")
	stringIdentity()

	fmt.Println("\n95. Switch Case:")
	checkGrade(85)

	fmt.Println("\n96. Static (Package Var):")
	counter()
	counter()

	fmt.Println("\n97. Constant:")
	constantDemo()

	fmt.Println("\n98. Cloning:")
	cloneDemo()

	fmt.Println("\n99. Constructor:")
	p := NewPerson("Alice", 25)
	fmt.Println(p)

	fmt.Println("\n100. Immutable Point:")
	ip := NewImmutablePoint(10, 20)
	fmt.Println(ip.X(), ip.Y())
}
