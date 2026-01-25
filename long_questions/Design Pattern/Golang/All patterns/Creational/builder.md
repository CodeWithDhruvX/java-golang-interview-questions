# Builder Pattern

## 🟢 What is it?
The **Builder Pattern** lets you construct complex objects step by step. It allows you to produce different types and representations of an object using the same construction code.

Think of it like **Subway (Sandwich Shop)**:
*   You don't just say "Give me a Sandwich" (Constructor with 10 nulls).
*   You say: "Start with Italian Bread" -> "Add Turkey" -> "Add Cheese" -> "No Onions" -> "Add Mayo".
*   The process is step-by-step, and the final result depends on your choices.

---

## 🎯 Strategy to Implement

1.  **Main Struct**: Define the struct you want to build.
2.  **Builder Struct**: Create a separate struct (or use the same one if simple) to hold the temporary state.
3.  **Chainable Methods**: Create methods for setting fields that return the Builder pointer (`*Builder`) to allow chaining.
4.  **Build Method**: Create a `Build()` method that returns the final fully constructed struct.

---

## 💻 Code Example

```go
package main

import "fmt"

type User struct {
    FirstName string // Required
    LastName  string // Required
    Age       int    // Optional
    Phone     string // Optional
    Address   string // Optional
}

type UserBuilder struct {
    user User
}

// Constructor with Required parameters
func NewUserBuilder(firstName, lastName string) *UserBuilder {
    return &UserBuilder{
        user: User{
            FirstName: firstName,
            LastName:  lastName,
        },
    }
}

// Chainable Setter Methods
func (b *UserBuilder) Age(age int) *UserBuilder {
    b.user.Age = age
    return b
}

func (b *UserBuilder) Phone(phone string) *UserBuilder {
    b.user.Phone = phone
    return b
}

func (b *UserBuilder) Address(address string) *UserBuilder {
    b.user.Address = address
    return b
}

// Build Method to return the final object
func (b *UserBuilder) Build() User {
    // Optional: Validate logic here
    return b.user
}

func main() {
    // Clean, readable, and flexible
    user := NewUserBuilder("John", "Doe").
        Age(30).
        Phone("123-456-7890").
        Build()

    // Another variation
    user2 := NewUserBuilder("Jane", "Smith").
        Address("123 Main St"). // No age or phone
        Build()

    fmt.Printf("%+v\n", user)
    fmt.Printf("%+v\n", user2)
}
```

---

## ✅ When to use?

*   **Telescoping Constructor Problem**: When your initialization function has 10 parameters, and you have to pass default values for 7 of them. `NewUser("John", "Doe", 0, "", "")` is ugly.
*   **Complex Creation**: When creating an object requires steps that might fail or validation logic.
*   **Immutability**: When you want "Immutable Objects" (though Go structs are mutable by default, you can restrict access by taking a `User` return value instead of `*User`).
