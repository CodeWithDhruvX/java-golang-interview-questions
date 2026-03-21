# SOLID Principles and Low-Level Design Impact

## Question
What are the SOLID principles and how would they impact in the low-level design with each approach?

## Answer

### SOLID Principles Overview
SOLID principles are the foundation of good Object-Oriented Design that guide developers in creating maintainable, scalable, and robust software systems.

### 1. Single Responsibility Principle (SRP)
**Definition**: A class should have one, and only one, reason to change.

**Low-Level Design Impact**:
- **Class Structure**: Create focused, single-purpose classes
- **Example**: Instead of a `OrderManager` that handles validation, processing, payment, and notification, create separate classes:
  - `OrderValidator` - validates order data
  - `OrderProcessor` - processes orders
  - `PaymentProcessor` - handles payments
  - `NotificationService` - sends notifications

**LLD Benefits**: Easier testing, reduced coupling, better maintainability

### 2. Open/Closed Principle (OCP)
**Definition**: Software entities should be open for extension but closed for modification.

**Low-Level Design Impact**:
- **Interface Design**: Use interfaces and abstract classes to allow extension
- **Example**: Payment system design
  ```java
  interface PaymentStrategy {
      void processPayment(double amount);
  }
  
  class CreditCardPayment implements PaymentStrategy { }
  class UPIPayment implements PaymentStrategy { }
  class PayPalPayment implements PaymentStrategy { }
  ```

**LLD Benefits**: New payment methods can be added without modifying existing code

### 3. Liskov Substitution Principle (LSP)
**Definition**: Subtypes must be substitutable for their base types without altering program correctness.

**Low-Level Design Impact**:
- **Inheritance Design**: Ensure subclasses honor base class contracts
- **Example**: Rectangle/Square problem
  ```java
  // Bad: Square cannot substitute Rectangle
  class Rectangle {
      void setWidth(double width) { this.width = width; }
      void setHeight(double height) { this.height = height; }
  }
  
  // Good: Use composition instead
  interface Shape {
      double getArea();
  }
  ```

**LLD Benefits**: Reliable polymorphism, predictable behavior

### 4. Interface Segregation Principle (ISP)
**Definition**: Clients should not be forced to depend on interfaces they do not use.

**Low-Level Design Impact**:
- **Interface Design**: Create small, focused interfaces
- **Example**: Instead of one large `Worker` interface:
  ```java
  interface Worker { void work(); }
  interface Eater { void eat(); }
  interface Sleeper { void sleep(); }
  
  class Human implements Worker, Eater, Sleeper { }
  class Robot implements Worker { }
  ```

**LLD Benefits**: Reduced coupling, easier implementation, better testability

### 5. Dependency Inversion Principle (DIP)
**Definition**: High-level modules should not depend on low-level modules; both should depend on abstractions.

**Low-Level Design Impact**:
- **Dependency Management**: Use dependency injection and interfaces
- **Example**: Service layer design
  ```java
  interface Database {
      void save(Object data);
  }
  
  class UserService {
      private final Database database;
      
      UserService(Database database) {
          this.database = database;
      }
  }
  ```

**LLD Benefits**: Easy unit testing, flexible implementations, reduced coupling

## Practical LLD Interview Example

### Parking Lot Design with SOLID Principles

**Without SOLID**:
```java
class ParkingLot {
    // Handles everything: parking, payment, validation, reporting
}
```

**With SOLID**:
```java
// SRP: Separate concerns
interface ParkingSpotManager { void parkVehicle(Vehicle v); }
interface PaymentProcessor { void processPayment(Payment p); }
interface TicketValidator { boolean validateTicket(Ticket t); }

// OCP: Extensible parking strategies
interface ParkingStrategy {
    ParkingSpot findSpot(List<ParkingSpot> spots);
}

// DIP: Depend on abstractions
class ParkingController {
    private final ParkingSpotManager spotManager;
    private final PaymentProcessor paymentProcessor;
    
    // Constructor injection
}
```

## Key Takeaways for LLD Interviews

1. **Mention SOLID explicitly** during requirements gathering
2. **Design classes with single responsibilities**
3. **Use interfaces for extensibility**
4. **Ensure proper inheritance hierarchies**
5. **Apply dependency injection for testability**

These principles directly impact your class diagrams, interface designs, and implementation approach in low-level design rounds.
