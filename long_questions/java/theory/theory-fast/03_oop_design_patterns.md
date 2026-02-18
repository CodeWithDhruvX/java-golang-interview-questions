# OOP & Design Patterns Interview Questions (36-50)

## Object-Oriented Programming & SOLID

### 36. Explain all SOLID principles with examples.
"This is my favorite topic. SOLID helps us write maintainable code.

**S - Single Responsibility Principle**: A class should do one thing. Instead of a `User` class handling data *and* saving itself to the DB *and* sending emails, I split it: `User` (data), `UserRepository` (DB), and `EmailService`.

**O - Open/Closed Principle**: Open for extension, closed for modification. If I have a `NotificationService`, I shouldn't modify its code to add SMS support. I should just be able to create a new `SmsNotification` class that implements a common `Notification` interface.

**L - Liskov Substitution Principle**: Subtypes must be substitutable for their base types. If I have a method that takes a `List`, it should work perfectly whether I pass an `ArrayList` or a `LinkedList`. If a subclass throws an exception where the parent didn't, it violates LSP.

**I - Interface Segregation**: Clients shouldn't be forced to depend on interfaces they don't use. Instead of one giant `Animal` interface with `fly()`, `swim()`, and `bark()`, I break them into smaller interfaces like `Flyable` and `Swimmable`.

**D - Dependency Inversion**: High-level modules shouldn't depend on low-level modules; both should depend on abstractions. My `UserService` doesn't depend on `MySQLDatabase` class directly; it depends on a `Database` interface. This lets me swap MySQL for PostgreSQL easily."

### 37. Difference between abstraction and encapsulation?
"They often get confused, but the intent is different.

**Encapsulation** is about **hiding data**. It's packaging data and methods into a single unit (class) and restricting access using `private`. It protects the internal state from being corrupted by the outside world.

**Abstraction** is about **hiding complexity**. It shows only the essential features of an object. When I use a `Car`, I just use the steering wheel and pedals (Interface). I don't need to know how the combustion engine works internally (Implementation). That's abstraction."

### 38. What is composition vs inheritance?
"Inheritance is an **IS-A** relationship (e.g., `Dog` IS-A `Animal`). It allows code reuse but creates tight coupling. If the parent class changes, all children break. Plus, you can't change the parent at runtime.

Composition is a **HAS-A** relationship (e.g., `Car` HAS-A `Engine`). You construct complex objects by combining simpler ones. It’s much more flexible because you can swap components at runtime.

The industry standard now is **Composition over Inheritance**. I almost always prefer composition unless there is a very strict hierarchical relationship."

### 39. What is dependency injection?
"Dependency Injection (DI) is basically the 'D' in SOLID. Instead of a class creating its own dependencies (using `new`), it asks for them from the outside (via constructor or setter).

So instead of `OrderService` saying `new EmailService()`, it says, 'Hey, whoever creates me, please pass me an EmailService.'

This makes testing incredibly easy because I can pass a mock `EmailService` during unit tests. In Spring, the container does all this wiring for us automatically."

### 40. How does SOLID help in real projects?
"In my last project, we had a payment module. Initially, it was hardcoded for PayPal.

When the business asked to add Stripe, if we hadn't used SOLID, we would have had to rewrite the entire `PaymentService`, risking bugs in the existing PayPal flow (violating Open/Closed).

But because we used dependency injection and interfaces, we just created a `StripeAdapter` implementing `PaymentProcessor`, and configured the app to use it. No existing code had to change. That’s the real value: it makes future changes safe and cheap."

## Design Patterns

### 41. Singleton pattern – problems and solutions?
"Singleton ensures a class has only one instance. It’s useful for things like a Configuration Manager or a Connection Pool.

The problem is it introduces global state, which makes unit testing specific components really hard because you can't easily mock the singleton. Also, in a multithreaded environment, if you don't implement it carefully, you can end up with multiple instances.

We solve the testing issue by using Dependency Injection frameworks (like Spring) which manage the singleton scope for us, rather than writing `getInstance()` everywhere."

### 42. How do you implement thread-safe Singleton?
"The classic way is **Double-Checked Locking**.

You check if the instance is null. If it is, you enter a `synchronized` block. Then inside that block, you check *again* if it's null (in case another thread beat you to it), and only then create the instance.

The variable must be declared `volatile` to prevent instruction reordering issues.

Or, simpler: use an `Enum`. Java guarantees enums are instantiated only once and are thread-safe. It’s the cleanest way."

### 43. Factory vs Abstract Factory?
"Factory Method is about creating one type of product. You have a `CarFactory` that creates `Sedan` or `SUV`. The client doesn't know the exact class, just the interface.

Abstract Factory is a factory of factories. It creates *families* of related products. Imagine a UI toolkit. You have a `ThemeFactory`.
If you use the `DarkThemeFactory`, it creates a `DarkButton` and a `DarkWindow`.
If you use `LightThemeFactory`, it creates `LightButton` and `LightWindow`.

It ensures that the Button and Window match each other."

### 44. When would you use Builder pattern?
"I use Builder whenever I have a complex object with many optional parameters.

Instead of a constructor with 10 arguments (which represents 'telescoping constructor' anti-pattern), or calling 10 setters (which leaves the object in an inconsistent state during creation), the Builder lets me chain methods: `.setName().setAge().build()`.

It makes the code readable and ensures the object is fully constructed and valid before it’s returned."

### 45. Strategy pattern real-world use case?
"We use Strategy heavily for things like Validation or Pricing.

For example, a shopping cart. We have a `DiscountStrategy` interface. We can have implementations like `holidayDiscount`, `membervipDiscount`, or `noDiscount`.

At runtime, based on the user or date, we inject the correct strategy into the Cart calculation method. This avoids massive `if-else` blocks and makes it easy to add new discount rules later."

### 46. Observer pattern example?
"This is the core of event-driven programming. It defines a subscription mechanism.

A real-world example is a Distributed Event Bus like Kafka or RabbitMQ, but in pure Java, it’s used in GUI listeners (button clicks).

In backend logic, think of an Order system. When an order is 'Placed', multiple services need to know: Inventory (to reserve stock), Notification (to email user), and Analytics. Instead of the OrderService calling all of them, it publishes an `OrderPlaced` event, and those services 'observe' and react to it independently."

### 47. Anti-patterns you have seen?
"The most common one is the **God Class**—a single class that does everything (Concepts + Logic + DB access). I’ve seen Services with 5000 lines of code.

Another is **Spaghetti Code**, usually caused by using exceptions for flow control or overly complex logic without abstraction.

I also see **Premature Optimization** a lot—people writing complex caching layers before checking if the database query is actually slow. It adds complexity for no proven gain."

### 48. What is clean code?
"Clean code is code that is easy to understand and easy to change.

To me, it means:
1.  **Meaningful Names**: Variables that explain *what* they are (`daysSinceCreation` vs `d`).
2.  **Small Functions**: Functions should do one thing and fit on a screen.
3.  **No Comments**: Ideally, the code should explain itself. If you need a comment to explain *what* the code does, refactor the code. Comments are for *why*.
4.  **Handling Errors**: Proper exception handling instead of returning error codes."

### 49. How do you reduce tight coupling?
"Interfaces are the key. I program to an interface, not an implementation.

I also use Dependency Injection to invert control.

And I use weirdly specific DTOs (Data Transfer Objects). I don’t pass my internal Database Entities to the frontend or other microservices. I map them to DTOs. This decouples my internal schema from the external contract."

### 50. How do you design extensible code?
"I focus on the Open/Closed Principle.

I try to identify what parts of the system are likely to change (like tax rules or report formats) and isolate them behind interfaces or strategy patterns.

I also avoid 'Magic Numbers' and hardcoded strings; I move them to config files or constants. This way, adding a new feature usually just means adding a new class, not rewriting existing logic."
