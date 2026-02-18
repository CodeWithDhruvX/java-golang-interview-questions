# Design Patterns & Architecture - Interview Answers

> 🎯 **Focus:** Patterns are tools, not rules. Focus on *where* you've used them.

### 1. What is the Singleton Pattern?
"It ensures a class has only one instance and provides a global access point to it.

I use this for expensive resources like a Database Connection Pool or a Configuration Manager. You don't want to parse the `application.yml` file every time you need a property; you want to parse it once and share that instance.

In modern Spring, all beans are Singletons by default, so the framework handles this for us. But if I implement it manually, I typically use an `Enum` because it's inherently thread-safe and prevents serialization attacks."

---

### 2. Factory vs Abstract Factory?
"Both are creational patterns to hide the complexity of creating objects.

**Factory Method** relies on inheritance. You have a method `createAnimal()` and subclasses decide whether to return a Dog or a Cat. It creates *one* product.

**Abstract Factory** is a step up—it's a factory of factories. It creates *families* of related products. Like a UI theme factory: one implementation gives you MacButtons and MacWindows; another gives you WinButtons and WinWindows.

To be honest, I rarely write Abstract Factories manually; dependency injection usually solves the problem of swapping implementations."

---

### 3. What is the Strategy Pattern?
"It lets you define a family of algorithms and make them interchangeable at runtime.

For example, imagine a Payment Processing system. I can have a `PaymentStrategy` interface with a `pay()` method.
I then implement `CreditCardStrategy`, `PayPalStrategy`, and `BitcoinStrategy`.

In my code, I don't write `if (type == PAYPAL)`. I just inject the correct Strategy and call `strategy.pay()`. This makes the code open for extension (Open/Closed Principle)—I can add ApplePay later without touching the core logic."

---

### 4. What is the Observer Pattern?
"It defines a subscription mechanism. When one object (Subject) changes state, all its dependents (Observers) are notified automatically.

A classic real-world example is an Event Listener. When a user clicks a button, the button 'fires' an event, and all the registered listeners run their code.

In distributed systems, this has evolved into the **Pub/Sub** model (like Kafka or RabbitMQ). Service A publishes an event 'UserCreated', and Service B and C listen to it and react."

---

### 5. What are SOLID principles?
"They are five rules for writing clean, maintainable object-oriented code.

**S - Single Responsibility:** A class should have one reason to change. Don't mix User Validation logic with Database logic.
**O - Open/Closed:** Open for extension, closed for modification. Use interfaces so you can add new behavior without rewriting existing code.
**L - Liskov Substitution:** A subclass should be swappable for its parent without breaking the app.
**I - Interface Segregation:** Better to have tiny specific interfaces (`Readable`, `Writable`) than one huge general one.
**D - Dependency Inversion:** Depend on abstractions, not concrete classes. This is what Spring DI is all about."

---

### 6. What is the Builder Pattern?
"It’s used to construct complex objects step-by-step.

Instead of a constructor with 10 parameters—which is hard to read and error-prone—you chain method calls like `.setName("John").setAge(30).build()`.

I use this religiously for DTOs and test data setup. I typically use the **Lombok** `@Builder` annotation so I don't have to write the boilerplate code manually."

---

### 7. Composition vs Inheritance?
"**Composition** means 'Has-A' relationship. A Car *has an* Engine.
**Inheritance** means 'Is-A' relationship. A Car *is a* Vehicle.

I strongly prefer **Composition over Inheritance**. Inheritance is rigid; you are married to the parent class. If the parent changes, it breaks the child. Composition is flexible—I can swap out the Engine component for a different one easily. It leads to looser coupling."

---

### 8. What is the Proxy Pattern?
"A Proxy acts as a placeholder or gatekeeper for another object.

The best example is Spring's `@Transactional`. When I call a method, I'm not calling my actual class; I'm calling a Spring Proxy. This proxy starts the database transaction, *then* calls my actual method, and then commits the transaction.

It allows you to wrap behavior (like security or logging) around the real logic transparently."

---

### 9. Adapter Pattern?
"It allows incompatible interfaces to work together. It’s like a travel adapter that lets you plug a US laptop into a UK socket.

In code, I often use this when integrating with a 3rd party legacy library. They might expect data in XML, but my app uses JSON. I write an 'Adapter' class that converts my JSON objects into their XML format so the two systems can talk without changing their core code."

---

### 10. Template Method Pattern?
"It defines the skeleton of an algorithm in a base class but lets subclasses override specific steps.

For example, a `DataParser` class might have a fixed flow: `openFile()`, `parseData()`, `closeFile()`.
The base class enforces that structure. But `parseData()` is abstract. One subclass implements CSV parsing, another implements JSON parsing.

It ensures consistency in the process while allowing flexibility in the details."
