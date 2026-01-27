package main

import (
	"fmt"
)

// --- OOPS CONCEPTS (Simulated/Explained) ---

// 81. Method Overloading
// Go does not support method overloading. Functions must have unique names.
// Example: func PrintInt(i int) {}, func PrintString(s string) {}

// 82. Method Overriding
// Go uses embedding for composition, not inheritance. Methods can be shadowed.
type Animal struct{}

func (a Animal) Sound() { fmt.Println("Generic Sound") }

type Dog struct{ Animal }

func (d Dog) Sound() { fmt.Println("Bark") }

// 83. Encapsulation
// In Go, capitalization determines visibility.
// unexported field (private) vs Exported field (public)
type BankAccount struct {
	balance float64 // lowercase = private to package
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

// 84. Inheritance
// Go uses Struct Embedding.
// type Person struct { Name string }
// type Employee struct { Person; ID int }

// 85. Polymorphism
// Go uses Interfaces.
type Shape interface {
	Area() float64
}
type Circle struct{ Radius float64 }

func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

// 86. Abstract Class vs Interface
// Go only has Interfaces (Method signatures). No Abstract Classes.

// 87. Final
// Go uses 'const' for values. No final methods/classes.

// 88. Exception Handling
// Go uses panic, defer, recover.

// 89. Singleton
// Typically using sync.Once
/*
var instance *Singleton
var once sync.Once
func GetInstance() *Singleton {
	once.Do(func() { instance = &Singleton{} })
	return instance
}
*/

// 90. Immutable
// Export fields via getters only, no setters.

// --- SQL & BASIC CS (Queries as String constants) ---

const (
	Q91_SecondHighestSalary = "SELECT MAX(Salary) FROM Employee WHERE Salary < (SELECT MAX(Salary) FROM Employee);"
	Q92_DeleteVsTruncate    = "Delete: DML (Row by row, slow, rollback). Truncate: DDL (All pages, fast, no rollback)."
	Q93_Joins               = "Inner: Matching rows. Left: All left + matching right."
	Q94_DuplicateRows       = "SELECT col, COUNT(*) FROM table GROUP BY col HAVING COUNT(*) > 1;"
	Q95_PKvsFK              = "PK: Unique Identifier. FK: Reference to PK in another table."
	Q96_Normalization       = "Process to minimize redundancy (1NF, 2NF, 3NF)."
	Q97_Index               = "Data structure (B-Tree) to speed up retrieval."
	Q98_ACID                = "Atomicity, Consistency, Isolation, Durability."
	Q99_Deadlock            = "Situation where two processes wait for each other to release resources."
	Q100_GetVsPost          = "GET: Request data (Idempotent). POST: Submit data (Not idempotent)."
)

func main() {
	fmt.Println("--- OOPS CONCEPTS ---")
	d := Dog{}
	d.Sound() // Bark (Overriding simulated)

	acc := BankAccount{}
	acc.Deposit(100)
	fmt.Println("Balance (Encapsulation):", acc.GetBalance())

	c := Circle{Radius: 5}
	var s Shape = c
	fmt.Println("Polymorphism Area:", s.Area())

	fmt.Println("\n--- SQL & CS QUESTIONS ---")
	fmt.Println("91. Second Highest Salary Query:", Q91_SecondHighestSalary)
	fmt.Println("92. Delete vs Truncate:", Q92_DeleteVsTruncate)
	fmt.Println("93. Joins:", Q93_Joins)
	fmt.Println("94. Duplicate Rows Query:", Q94_DuplicateRows)
	fmt.Println("95. PK vs FK:", Q95_PKvsFK)
	fmt.Println("100. GET vs POST:", Q100_GetVsPost)
}
