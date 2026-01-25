# Strategy Pattern

## 🟢 What is it?
The **Strategy Pattern** lets you define a family of algorithms and make them interchangeable.

Think of it like **Navigation**: Car, Bus, or Taxi.

---

## 🎯 Strategy to Implement

1.  **Strategy Interface**: Declare the common interface (e.g., `PaymentStrategy`).
2.  **Concrete Strategies**: Implement the algorithm.
3.  **Context**: Maintains reference to Strategy.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Strategy Interface
type PaymentStrategy interface {
    Pay(amount int)
}

// 2. Concrete Strategies
type CreditCardPayment struct {
    CardNumber string
}

func (c *CreditCardPayment) Pay(amount int) {
    fmt.Printf("Paid $%d using Credit Card ending in %s\n", amount, c.CardNumber[len(c.CardNumber)-4:])
}

type PayPalPayment struct {
    Email string
}

func (p *PayPalPayment) Pay(amount int) {
    fmt.Printf("Paid $%d using PayPal (%s)\n", amount, p.Email)
}

// 3. Context
type ShoppingCart struct {
    strategy PaymentStrategy
}

func (s *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
    s.strategy = strategy
}

func (s *ShoppingCart) Checkout(amount int) {
    if s.strategy == nil {
        fmt.Println("Please select a payment method.")
    } else {
        s.strategy.Pay(amount)
    }
}

func main() {
    cart := &ShoppingCart{}
    
    // User chooses Credit Card
    cart.SetPaymentStrategy(&CreditCardPayment{CardNumber: "1234567890123456"})
    cart.Checkout(100)

    // User changes mind to PayPal
    cart.SetPaymentStrategy(&PayPalPayment{Email: "user@example.com"})
    cart.Checkout(100)
}
```

---

## ✅ When to use?

*   **Runtime Selection**: Switching algorithms.
*   **Eliminate Conditionals**: Replacing `if type == 'A'` logic.
