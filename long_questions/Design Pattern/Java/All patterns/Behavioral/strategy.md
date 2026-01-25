# Strategy Pattern

## 🟢 What is it?
The **Strategy Pattern** is a behavioral pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.

Think of it like **Navigation**:
*   You want to go to the Airport.
*   You can choose a **Strategy**:
    *   **Car Strategy**: Drive yourself (Fast, Cost varies).
    *   **Bus Strategy**: Public transport (Slow, Cheap).
    *   **Taxi Strategy**: Cab (Fast, Expensive).
*   The "Navigator" (Context) stays the same, but the "Route Calculation" (Strategy) changes based on your choice.

---

## 🎯 Strategy to Implement

1.  **Strategy Interface**: Declare an interface common to all supported versions of the algorithm (e.g., `RouteStrategy`).
2.  **Concrete Strategies**: Implement the algorithm in different classes (e.g., `RoadStrategy`, `PublicTransportStrategy`).
3.  **Context Class**: Define a class that maintains a reference to a Strategy object. The Context doesn't know the concrete class of a strategy. It works with all strategies via the interface.
4.  **Client Code**: The client decides which strategy to use and passes it to the Context.

---

## 💻 Code Example

```java
// 1. Strategy Interface
interface PaymentStrategy {
    void pay(int amount);
}

// 2. Concrete Strategies
class CreditCardPayment implements PaymentStrategy {
    private String cardNumber;

    public CreditCardPayment(String cardNumber) {
        this.cardNumber = cardNumber;
    }

    @Override
    public void pay(int amount) {
        System.out.println("Paid $" + amount + " using Credit Card ending in " + cardNumber.substring(cardNumber.length() -4));
    }
}

class PayPalPayment implements PaymentStrategy {
    private String email;

    public PayPalPayment(String email) {
        this.email = email;
    }

    @Override
    public void pay(int amount) {
        System.out.println("Paid $" + amount + " using PayPal (" + email + ")");
    }
}

// 3. Context
class ShoppingCart {
    private PaymentStrategy paymentStrategy;

    // Strategy set at runtime
    public void setPaymentStrategy(PaymentStrategy strategy) {
        this.paymentStrategy = strategy;
    }

    public void checkout(int amount) {
        if (paymentStrategy == null) {
            System.out.println("Please select a payment method.");
        } else {
            paymentStrategy.pay(amount);
        }
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        ShoppingCart cart = new ShoppingCart();
        int billAmount = 100;

        // User chooses Credit Card
        cart.setPaymentStrategy(new CreditCardPayment("1234567890123456"));
        cart.checkout(billAmount);

        // User changes mind to PayPal
        cart.setPaymentStrategy(new PayPalPayment("user@example.com"));
        cart.checkout(billAmount);
    }
}
```

---

## ✅ When to use?

*   **Many Versions of Algorithm**: When you have diverse ways of solving a problem (sorting, compression, pathfinding) and you want to be able to switch between them.
*   **Runtime Selection**: When you want to choose the algorithm implementation at runtime based on user input or system configuration.
*   **Eliminate Conditionals**: When you have a massive `if...else if...else if` block that selects specific behavior. Moving each branch into a Strategy class cleans up the code significantly.
