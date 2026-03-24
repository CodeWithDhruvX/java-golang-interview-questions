# OOP & Design Patterns Interview Questions (36-50)

## Object-Oriented Programming & SOLID

### 36. Explain all SOLID principles with examples.
"This is my favorite topic. SOLID helps us write maintainable code.

**S - Single Responsibility Principle**: A class should do one thing. Instead of a `User` class handling data *and* saving itself to the DB *and* sending emails, I split it: `User` (data), `UserRepository` (DB), and `EmailService`.

**O - Open/Closed Principle**: Open for extension, closed for modification. If I have a `NotificationService`, I shouldn't modify its code to add SMS support. I should just be able to create a new `SmsNotification` class that implements a common `Notification` interface.

**L - Liskov Substitution Principle**: Subtypes must be substitutable for their base types. If I have a method that takes a `List`, it should work perfectly whether I pass an `ArrayList` or a `LinkedList`. If a subclass throws an exception where the parent didn't, it violates LSP.

**I - Interface Segregation**: Clients shouldn't be forced to depend on interfaces they don't use. Instead of one giant `Animal` interface with `fly()`, `swim()`, and `bark()`, I break them into smaller interfaces like `Flyable` and `Swimmable`.

**D - Dependency Inversion**: High-level modules shouldn't depend on low-level modules; both should depend on abstractions. My `UserService` doesn't depend on `MySQLDatabase` class directly; it depends on a `Database` interface. This lets me swap MySQL for PostgreSQL easily."

**Spoken Format:**
"SOLID is like having five golden rules for building LEGO castles that won't fall apart.

**Single Responsibility** is like giving each LEGO person only one job - one person builds walls, another builds roofs, another builds doors. If you ask the wall-builder to also do plumbing, they'll get confused and do both poorly.

**Open/Closed** is like having LEGO pieces that snap together but you can add new types without changing the existing ones. If you have a car base, you can add a helicopter piece without changing how the car wheels work.

**Liskov Substitution** is like saying any square LEGO piece should work wherever a rectangle piece is expected. If you have a spot for a rectangle, a square should fit perfectly without breaking anything.

**Interface Segregation** is like having small, specific instruction manuals instead of one giant book. If you only need to build walls, you get the wall-building manual, not the entire castle-building book.

**Dependency Inversion** is like building with connectors instead of gluing pieces directly. Your castle doesn't depend on a specific brand of LEGO pieces - it depends on standard connectors, so you can use any brand that fits the connectors.

These principles help you build code that doesn't break when you need to change something later!"

### 37. Difference between abstraction and encapsulation?
"They often get confused, but the intent is different.

**Encapsulation** is about **hiding data**. It's packaging data and methods into a single unit (class) and restricting access using `private`. It protects the internal state from being corrupted by the outside world.

**Abstraction** is about **hiding complexity**. It shows only the essential features of an object. When I use a `Car`, I just use the steering wheel and pedals (Interface). I don't need to know how the combustion engine works internally (Implementation). That's abstraction."

**Spoken Format:**
"These two concepts get mixed up all the time, but they solve different problems.

**Encapsulation** is like having a locked box - you put your data inside and lock it. Only the box itself knows what's inside and can change it. Others can only use the methods you provide (like a small slot to put coins in). This protects your data from being messed up by outsiders.

**Abstraction** is like having a TV remote - you see the volume button, channel button, power button. You don't need to know about the circuit boards, antennas, or how the signal gets processed. You just use the simple interface.

The key difference: Encapsulation hides data, Abstraction hides complexity. Encapsulation is about protection, Abstraction is about simplicity.

Think of it this way: Enculation puts your valuables in a safe, Abstraction gives you a simple remote to control complex things without understanding how they work!"

### 38. What is composition vs inheritance?
"Inheritance is an **IS-A** relationship (e.g., `Dog` IS-A `Animal`). It allows code reuse but creates tight coupling. If the parent class changes, all children break. Plus, you can't change the parent at runtime.

Composition is a **HAS-A** relationship (e.g., `Car` HAS-A `Engine`). You construct complex objects by combining simpler ones. It’s much more flexible because you can swap components at runtime.

The industry standard now is **Composition over Inheritance**. I almost always prefer composition unless there is a very strict hierarchical relationship."

**Spoken Format:**
"Think of building with LEGOs - there are two ways to do it.

**Inheritance** is like saying 'my new car IS-A vehicle'. You start with a vehicle base and add car-specific features. The problem is if you change something in the vehicle base, every car, truck, motorcycle that inherits from it might break. Plus, you're stuck with whatever vehicle features you started with.

**Composition** is like saying 'my car HAS-A engine, HAS-A wheels, HAS-A steering wheel'. You build your car by combining these independent parts. The beauty is you can swap the engine anytime without changing the car itself.

The modern approach is **Composition over Inheritance** because it's like building with interchangeable parts instead of being locked into a family hierarchy. You can change pieces easily, test them independently, and your code is much more flexible.

Inheritance is like being born into a family - you can't choose your relatives. Composition is like choosing your friends - you can pick and change them anytime!"

### 39. What is dependency injection?
"Dependency Injection (DI) is basically the 'D' in SOLID. Instead of a class creating its own dependencies (using `new`), it asks for them from the outside (via constructor or setter).

So instead of `OrderService` saying `new EmailService()`, it says, 'Hey, whoever creates me, please pass me an EmailService.'

This makes testing incredibly easy because I can pass a mock `EmailService` during unit tests. In Spring, the container does all this wiring for us automatically."

**Spoken Format:**
"Dependency Injection is like having a personal assistant who brings you what you need instead of you having to get everything yourself.

Without DI, your code is like a chef who has to grow their own vegetables, raise their own chickens, and build their own kitchen just to cook a meal. The chef is tightly coupled to all these dependencies.

With DI, the chef just says 'I need vegetables and chicken' and the restaurant manager brings them. The chef focuses only on cooking.

In code terms: instead of `OrderService` creating `new EmailService()`, it just says 'Hey, whoever creates me, please give me an EmailService'.

The magic is that testing becomes super easy - for testing, you just give the chef plastic vegetables and toy chicken (mock objects) to see if they can cook properly.

Spring is like having a restaurant manager who automatically knows what every chef needs and brings them the right ingredients. It's all about letting someone else handle the dependencies so you can focus on your main job!"

### 40. How does SOLID help in real projects?
"In my last project, we had a payment module. Initially, it was hardcoded for PayPal.

When the business asked to add Stripe, if we hadn't used SOLID, we would have had to rewrite the entire `PaymentService`, risking bugs in the existing PayPal flow (violating Open/Closed).

But because we used dependency injection and interfaces, we just created a `StripeAdapter` implementing `PaymentProcessor`, and configured the app to use it. No existing code had to change. That’s the real value: it makes future changes safe and cheap."

## Design Patterns

### 41. Singleton pattern – problems and solutions?
"Singleton ensures a class has only one instance. It’s useful for things like a Configuration Manager or a Connection Pool.

The problem is it introduces global state, which makes unit testing specific components really hard because you can't easily mock the singleton. Also, in a multithreaded environment, if you don't implement it carefully, you can end up with multiple instances.

We solve the testing issue by using Dependency Injection frameworks (like Spring) which manage the singleton scope for us, rather than writing `getInstance()` everywhere."

**Spoken Format:**
"Singleton is like having only one president for a country - no matter who asks, they get the same person. It's useful for things like configuration managers or database connection pools where you want exactly one instance.

The problems are interesting: first, it's like having global state - everyone can access the president, which makes testing other parts of government really hard because you can't easily replace the president with a test double.

Second, in a multithreaded environment (like multiple people asking for the president at the same time), if you're not careful, you might end up with two presidents - which would be chaos!

The modern solution is to use Dependency Injection frameworks like Spring. Instead of everyone going to the presidential palace and asking for the president (calling `getInstance()`), you just say 'I need the president' and the framework brings you the right one.

This way, the framework handles all the complexity of ensuring there's only one instance, and for testing, it can easily give you a fake president instead. It's like having a chief of staff who manages access to the president for you!"

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

**Spoken Format:**
"Strategy pattern is like having different payment methods in your wallet and choosing the right one for each situation.

Imagine you're building a shopping cart. You need to calculate discounts, but the rules keep changing - sometimes it's holiday discount, sometimes VIP member discount, sometimes no discount.

Without Strategy, you'd have massive `if-else` blocks: `if (isHoliday) { discount = 20%; } else if (isVIP) { discount = 15%; } else { discount = 0%; }`

With Strategy, you create a `DiscountStrategy` interface. Then you have `HolidayDiscount`, `VIPDiscount`, and `NoDiscount` classes.

At runtime, based on the user or date, you inject the right strategy: `cart.setDiscountStrategy(new HolidayDiscount())`

The beauty is that adding a new discount rule (like student discount) just means creating a new `StudentDiscount` class. You don't touch the existing cart logic at all.

It's like having interchangeable tools in your toolbox - you pick the right tool for the job without changing how the toolbox itself works!"

### 46. Observer pattern example?
"This is the core of event-driven programming. It defines a subscription mechanism.

A real-world example is a Distributed Event Bus like Kafka or RabbitMQ, but in pure Java, it’s used in GUI listeners (button clicks).

In backend logic, think of an Order system. When an order is 'Placed', multiple services need to know: Inventory (to reserve stock), Notification (to email user), and Analytics. Instead of the OrderService calling all of them, it publishes an `OrderPlaced` event, and those services 'observe' and react to it independently."

**Spoken Format:**
"Observer pattern is like having a subscription service - when something interesting happens, everyone who subscribed gets notified automatically.

Think of a YouTube channel. When a creator uploads a new video (event), all subscribers get notified instantly. The creator doesn't need to know who subscribed or notify them individually.

In programming terms, this is perfect for event-driven systems.

A real example: an order processing system. When an order is placed, multiple things need to happen:
- Inventory needs to reserve the stock
- Notification service needs to email the customer
- Analytics needs to track the sale

Without Observer, the OrderService would need to call all these services directly. But with Observer, OrderService just publishes an `OrderPlaced` event.

The Inventory, Notification, and Analytics services all 'observe' this event and react independently. If you need to add a new service (like shipping), you just add another observer - you don't change the OrderService at all.

It's like having a town crier - when there's news, everyone in town who wants to hear it gets the message automatically!"

### 47. Anti-patterns you have seen?
"The most common one is the **God Class**—a single class that does everything (Concepts + Logic + DB access). I’ve seen Services with 5000 lines of code.

Another is **Spaghetti Code**, usually caused by using exceptions for flow control or overly complex logic without abstraction.

I also see **Premature Optimization** a lot—people writing complex caching layers before checking if the database query is actually slow. It adds complexity for no proven gain."

**Spoken Format:**
"Anti-patterns are like bad habits in coding - common mistakes that make code harder to work with.

The most famous is the **God Class** - imagine a single person who's the CEO, accountant, janitor, and receptionist all at once. In code, this is a class that handles database, business logic, UI, and everything else. I've seen classes with thousands of lines that do everything!

Another classic is **Spaghetti Code** - imagine trying to follow a single noodle through a plate of pasta. That's what happens when code has so many `if-else` statements and exceptions used for flow control that you can't follow the logic.

My favorite anti-pattern is **Premature Optimization** - this is like spending three days optimizing a piece of code that only runs once a year, while the main code that runs thousands of times per day is slow.

People add complex caching, thread pools, and fancy algorithms before even measuring if the original code was actually slow. They add all this complexity for no real benefit.

The rule is: make it work first, then make it fast. Don't optimize what doesn't need optimizing!"

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

**Spoken Format:**
"Designing extensible code is like building with LEGOs - you want to be able to add new pieces without taking apart what you already built.

My main focus is the **Open/Closed Principle** - be open for extension but closed for modification.

I try to identify what's likely to change. For example, in a tax system, the tax rules will definitely change every year. So I isolate tax logic behind a `TaxCalculator` interface.

When new tax rules come, I just create a new `TaxCalculator2025` class without touching existing code.

I also avoid 'Magic Numbers' - instead of `if (age > 65)`, I use `if (age > RETIREMENT_AGE)`. The constant is defined in one place, so if the retirement age changes, I only update it in one spot.

For hardcoded strings like database URLs or API keys, I move them to configuration files. This way, I can use different configs for development, staging, and production without changing code.

The goal is that adding a new feature should be like adding a new LEGO piece - it snaps into the existing structure without breaking what's already there!"
