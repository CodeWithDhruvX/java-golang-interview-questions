# Design Patterns Interview Questions & Answers

## 🔹 1. Creational Patterns (Questions 1-15)

**Q1: What is the Singleton Pattern?**
Ensures a class has only one instance and provides a global point of access to it (e.g., Database Connection, Logger).

**Q2: How do you implement a thread-safe Singleton?**
By using locks (mutex) or synchronization mechanisms during the creation of the instance to ensure two threads don't create it simultaneously.

**Q3: What is the Factory Design Pattern?**
A method that creates objects without specifying the exact class of object that will be created, allowing subclasses to decide.

**Q4: What is the difference between Factory and Abstract Factory?**
Factory Method creates *one* type of object. Abstract Factory creates *families* of related objects (e.g., a UI toolkit with Buttons, Checkboxes, and Windows).

**Q5: When should you use the Builder Pattern?**
When constructing complex objects with many optional parameters, or when you want to make the creation process clear and readable.

**Q6: What is the Prototype Pattern?**
Creates new objects by copying (cloning) an existing object instead of creating a new one from scratch.

**Q7: Usage of Prototype Pattern?**
Useful when creating an object is expensive (e.g., requires database calls) and you need a similar object quickly.

**Q8: What is Eager vs. Lazy Initialization?**
Eager: The instance is created as soon as the application starts. Lazy: The instance is created only when it is actually needed.

**Q9: Can a Singleton be broken?**
Yes, in many languages, advanced features like reflection, serialization, or cloning can accidentally create a second instance.

**Q10: What is the best way to implement Singleton safely?**
Using language-specific features that guarantee a single instance, such as Enums (in some languages) or module-level singletons.

**Q11: What is Dependency Injection (DI)?**
A technique where objects are "given" their dependencies from the outside rather than creating them internally, making code easier to test.

**Q12: What is the value of a Singleton Registry?**
A central place (like a map) to store named Singleton instances, allowing you to look them up by name.

**Q13: When should you avoid Singleton?**
When it behaves like a global variable, making testing difficult and hiding dependencies between parts of your code.

**Q14: What is the Static Factory Method?**
A public static method that returns an instance of the class, often used instead of constructors for better naming and control.

**Q15: What is the Object Pool Pattern?**
A creational pattern that uses a set of initialized objects kept ready to use (e.g., Database Connection Pool), rather than creating and destroying them repeatedly.

---

## 🔹 2. Structural Patterns (Questions 16-30)

**Q16: What is the Adapter Pattern?**
Makes incompatible interfaces work together. It acts like a real-world power adapter, converting one interface into another.

**Q17: What is the Decorator Pattern?**
Adds new functionality to an existing object dynamically without altering its structure (e.g., adding "ScrollBars" to a "Window" object).

**Q18: What is the Facade Pattern?**
Provides a simple, unified interface to a complex system. It hides the complexity of the underlying subsystems (e.g., a "Start Car" button handling engine, fuel, and battery).

**Q19: Difference between Adapter and Facade?**
Adapter fixes incompatibility between two existing interfaces. Facade simplifies a complex system by providing a new interface.

**Q20: What is the Proxy Pattern?**
A placeholder for another object to control access to it. Used for things like lazy loading, security checks, or remote access.

**Q21: What is the Composite Pattern?**
Treats individual objects and groups of objects the same way. Useful for tree structures like file systems (Files and Folders).

**Q22: What is the Bridge Pattern?**
Splits an object into two parts: the abstraction (what it does) and the implementation (how it does it), so they can change independently.

**Q23: What is the Flyweight Pattern?**
Saves memory by sharing common parts of state between many objects (e.g., sharing the font data for every letter 'A' in a document).

**Q24: Difference between Proxy and Decorator?**
Proxy controls *access* to an object. Decorator adds *behavior* to an object.

**Q25: Real-world example of Adapter Pattern?**
A memory card reader acts as an adapter, allowing a computer to read a memory card it naturally cannot connect to.

**Q26: Real-world example of Decorator Pattern?**
Adding toppings to a pizza. The base pizza is the object, and toppings (Cheese, Pepperoni) are decorators adding flavor/cost.

**Q27: Real-world example of Facade Pattern?**
A customer service department. You talk to one person (Facade) who handles issues with billing, shipping, or technical support internally.

**Q28: Types of Proxies?**
Remote Proxy (network access), Virtual Proxy (delaying expensive creation), Protection Proxy (access control).

**Q29: When to use Composite Pattern?**
When you have a hierarchy of objects (like menus and sub-menus) and want to treat a single item and a full menu uniformly.

**Q30: How does Bridge Pattern differ from Strategy Pattern?**
Bridge is about structure (separating interface from code). Strategy is about behavior (swapping algorithms).

---

## 🔹 3. Behavioral Patterns (Questions 31-50)

**Q31: What is the Observer Pattern?**
A subscription mechanism where objects (Observers) listen for changes in another object (Subject) and get notified automatically (e.g., YouTube channel notifications).

**Q32: What is the Strategy Pattern?**
Allows you to swap algorithms or behaviors at runtime. You can choose different "Strategies" (e.g., Sorting methods, Payment methods) without changing the code that uses them.

**Q33: What is the Command Pattern?**
Turns a request (like "Save" or "Copy") into a standalone object. This allows you to queue requests, undo them, or pass them around.

**Q34: What is the Template Method Pattern?**
Defines the skeleton of an algorithm in a base class but lets subclasses fill in the specific steps (e.g., a recipe with fixed steps but interchangeable ingredients).

**Q35: What is the Iterator Pattern?**
Provides a standard way to loop through a collection of items without needing to know how they are stored (Array, Linked List, Tree).

**Q36: What is the Chain of Responsibility Pattern?**
Passes a request along a chain of handlers. Each handler decides to either process the request or pass it to the next one (e.g., Customer Support tiers).

**Q37: What is the State Pattern?**
Allows an object to change its behavior when its internal state changes (e.g., a Phone behaves differently in "Silent", "Ring", or "Vibrate" mode).

**Q38: What is the Mediator Pattern?**
Restricts direct communication between objects and forces them to collaborate only through a "Mediator" object, reducing dependnecies (e.g., Air Traffic Control Tower).

**Q39: What is the Visitor Pattern?**
Lets you add new operations to existing objects without modifying them. You create a "Visitor" that moves through the objects and performs the action.

**Q40: What is the Memento Pattern?**
Saves the state of an object so it can be restored later (e.g., "Undo" feature in editors or "Save Game" in video games).

**Q41: Difference between Strategy and State Pattern?**
Strategy: The client typically chooses the strategy once. State: The object acts like it changes its class/type as its internal state changes automatically.

**Q42: Difference between Template Method and Strategy?**
Template Method uses inheritance (subclasses override steps). Strategy uses composition (plugging in a different object).

**Q43: Real-world example of Observer Pattern?**
Newsletters. You subscribe (observe), and when a new issue comes out (event), it gets sent to you automatically.

**Q44: Real-world example of Chain of Responsibility?**
ATM money dispensing. It checks if the $100 bin can pay, if not, tries $50, then $20, etc.

**Q45: When to use Command Pattern?**
When you need to support "Undo/Redo" functionality or schedule tasks for later execution.

**Q46: What is the Interpreter Pattern?**
Used to evaluate sentences in a language. Used in SQL parsers, or calculating mathematical expressions (Like `5 + 10 * 2`).

**Q47: Benefits of Iterator Pattern?**
You can loop through a list, a tree, or a graph using the same simple code loop, hiding the complex traversal logic.

**Q48: Drawback of Visitor Pattern?**
If you add a new type of element to your structure, you have to update every single Visitor to handle it.

**Q49: How does Mediator reduce coupling?**
Instead of every plane talking to every other plane (chaos), they all talk to the Tower (organized).

**Q50: Use case for Null Object Pattern?**
Instead of returning `null` and causing potential crashes, return a "fake" object that does nothing (e.g., a "silent" logger instead of `null` logger).

---

## 🔹 4. Architectural & General Patterns (Questions 51-65)

**Q51: What is MVC (Model-View-Controller)?**
Separates an app into 3 parts: Model (Data), View (User Interface), and Controller (Process Logic).

**Q52: What is the DAO Pattern (Data Access Object)?**
Separates the business logic from the database code. The DAO is responsible for all database operations (CRUD).

**Q53: What is the DTO Pattern (Data Transfer Object)?**
A simple object used only to pass data between different parts of a system (like from server to client) to reduce the number of calls.

**Q54: What is the Front Controller Pattern?**
A single handler (Controller) that receives all requests for a website and routes them to the correct place.

**Q55: What is the Repository Pattern?**
Similar to DAO, but acts like an in-memory collection of objects. It hides the details of how data is actually stored or retrieved.

**Q56: What is the Service Locator Pattern?**
A central registry where the application can ask for the services it needs ("Give me the Database Service").

**Q57: What is MVVM (Model-View-ViewModel)?**
An evolution of MVC used in modern UIs. The "ViewModel" holds the state of the View and updates it automatically via data binding.

**Q58: What is Inversion of Control (IoC)?**
Instead of your code controlling the flow of the program, a framework controls it and calls your code when needed ("Hollywood Principle").

**Q59: What is the difference between DI and Service Locator?**
DI pushes dependencies into the object (Passive). Service Locator requires the object to ask for them (Active).

**Q60: What is the Intercepting Filter Pattern?**
A chain of filters that process a request before it reaches the core application (e.g., Authentication check -> Logging -> Compression).

**Q61: What is a Value Object?**
An object defined by its attributes, not an ID. Two Value Objects are equal if their values are equal (e.g., a "Color" or "Money" object).

**Q62: What is the Unit of Work Pattern?**
Keeps track of everything you do during a business transaction and sends all changes to the database at once.

**Q63: What represents the View in an API?**
The data (JSON/XML) sent back to the client is the "View", as specific UI rendering happens on the client side.

**Q64: What is Microservices Architecture?**
Building an application as a collection of small, independent services that talk to each other, rather than one giant "Monolith".

**Q65: What is an API Gateway?**
A single entry point for all API requests. It handles routing, security, and rate limiting before passing requests to backend services.

---

## 🔹 5. SOLID & Design Principles (Questions 66-70)

**Q66: What is 'S' in SOLID?**
Single Responsibility Principle: A class or module should do one thing and do it well.

**Q67: What is 'O' in SOLID?**
Open/Closed Principle: Code should be open for extension (adding new features) but closed for modification (changing existing code).

**Q68: What is 'L' in SOLID?**
Liskov Substitution Principle: You should be able to replace a parent class with any of its child classes without breaking the app.

**Q69: What is 'I' in SOLID?**
Interface Segregation Principle: It is better to have many small, specific interfaces than one huge, general-purpose one.

**Q70: What is 'D' in SOLID?**
Dependency Inversion Principle: High-level modules should not depend on low-level details. Both should depend on abstractions (interfaces).
