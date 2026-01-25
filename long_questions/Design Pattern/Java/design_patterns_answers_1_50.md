# 🎨 Design Pattern Interview Questions (Questions 1-50)

---

## 🏗️ Creational Patterns (Questions 1-15)

### Question 1: What is the Singleton Pattern?

**Answer:**
The **Singleton Pattern** ensures that a class has only one instance and provides a global point of access to it. It is commonly used for logging, driver objects, caching, and thread pools.

**Key Concepts:**
- **Single Instance:** Restricts class instantiation to one object.
- **Global Access:** Provides a static method to get the instance.
- **Private Constructor:** Prevents direct instantiation.

**Code:**
```java
public class Singleton {
    private static Singleton instance;
    private Singleton() {} // Private Constructor

    public static Singleton getInstance() {
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }
}
```

**Use Case:**
Managing a connection to a database or a configuration manager where multiple instances would cause inconsistency.

---

### Question 2: How do you make Singleton thread-safe?

**Answer:**
In a multithreaded environment, the standard implementation can create duplicate instances. You can fix this using **Double-Checked Locking**.

**Key Concepts:**
- **Synchronized Block:** Locking only when creating the instance.
- **Volatile Keyword:** Ensures visibility of the instance across threads.

**Code:**
```java
public class ThreadSafeSingleton {
    private static volatile ThreadSafeSingleton instance;

    private ThreadSafeSingleton() {}

    public static ThreadSafeSingleton getInstance() {
        if (instance == null) {
            synchronized (ThreadSafeSingleton.class) {
                if (instance == null) {
                    instance = new ThreadSafeSingleton();
                }
            }
        }
        return instance;
    }
}
```

**Use Case:**
High-concurrency applications (e.g., Web Servers) accessing a shared resource.

---

### Question 3: What is the Factory Method Pattern?

**Answer:**
The **Factory Method Pattern** defines an interface for creating an object but lets subclasses alter the type of objects that will be created.

**Key Concepts:**
- **Decoupling:** Client code doesn’t know the concrete classes.
- **Polymorphism:** Returns an interface or abstract class.

**Code:**
```java
abstract class Logistics {
    abstract Transport createTransport();
    
    void planDelivery() {
        Transport t = createTransport();
        t.deliver();
    }
}

class RoadLogistics extends Logistics {
    Transport createTransport() { return new Truck(); }
}
```

**Use Case:**
A framework needs to standardize the architectural model for a range of applications, but allow for individual applications to define their own domain objects.

---

### Question 4: What is the difference between Factory and Abstract Factory?

**Answer:**
Both create objects, but they differ in scope and complexity.

**Key Differences:**
- **Factory Method:** Creates **one** product. Uses inheritance (subclasses decide instantiation).
- **Abstract Factory:** Creates **families** of related products (e.g., Sofa + Chair). Uses composition (factory object is passed).

**Code (Concept):**
```java
// Factory Method
Transport t = logistics.createTransport();

// Abstract Factory
FurnitureFactory f = new ModernFurnitureFactory();
Chair c = f.createChair();
Sofa s = f.createSofa(); // Guaranteed to match style
```

**Use Case:**
Use **Factory** for a single object type. Use **Abstract Factory** when you need to ensure a set of objects work together (e.g., UI Theme: DarkButton + DarkWindow).

---

### Question 5: What is the Builder Pattern?

**Answer:**
The **Builder Pattern** separates the construction of a complex object from its representation. It allows you to create different representations using the same construction code.

**Key Concepts:**
- **Step-by-Step Construction:** methods like `setpartA()`, `setpartB()`.
- **Fluent Interface:** Method chaining (`.name().age().build()`).
- **Immutability:** Often used to build immutable objects.

**Code:**
```java
User user = new User.UserBuilder("John", "Doe")
                .age(30)
                .phone("123-456")
                .build();
```

**Use Case:**
Constructing complex objects like HTTP Requests, SQL Queries, or Configuration objects with many optional parameters.

---

### Question 6: What is the Prototype Pattern?

**Answer:**
The **Prototype Pattern** is used to create a duplicate object or clone of the current object to enhance performance.

**Key Concepts:**
- **Cloning:** Uses `clone()` method or copy constructor.
- **Performance:** Avoids expensive database calls or initializations.

**Code:**
```java
// Logic inside clone() method
Circle clone = (Circle) existingCircle.clone();
```

**Use Case:**
When object creation is expensive (e.g., parsing a huge XML file), create one instance and clone it for subsequent uses.

---

### Question 7: What is the difference between Shallow Copy and Deep Copy?

**Answer:**
This is critical when implementing Prototype.
- **Shallow Copy:** Copies the object's fields. If a field is a reference to an object, it copies the *reference*, not the object (both point to same memory).
- **Deep Copy:** Copies the object and recursively copies all objects referenced by it (creates completely independent instances).

**Use Case:**
Use Deep Copy if the prototype has mutable fields (List, Date) that shouldn't be shared.

---

### Question 8: Can Enum be used as Singleton?

**Answer:**
Yes, and it is the preferred way in Java.

**Key Concepts:**
- **Serialization Safe:** Enums handle serialization automatically.
- **Reflection Safe:** You cannot call an Enum constructor via reflection.
- **Thread Safe:** Creation is thread-safe by default.

**Code:**
```java
public enum Config {
    INSTANCE;
    public void doSomething() { ... }
}
```

---

### Question 9: What is the Telescoping Constructor anti-pattern?

**Answer:**
It occurs when a class has multiple constructors with increasing numbers of parameters: `User(name)`, `User(name, age)`, `User(name, age, phone)`.

**Problem:**
Hard to read and maintain.

**Solution:**
Use the **Builder Pattern**.

---

### Question 10: When would you use Static Factory Method over Constructor?

**Answer:**
A **Static Factory Method** is a `public static` method that returns an instance of the class.

**Benefits:**
- **Names:** Constructors are always named `ClassName`. Factory methods can have descriptive names: `Complex.fromPolar(r, theta)`.
- **Caching:** Can return the same instance (Singleton/Flyweight) instead of always creating a new one.
- **Subtypes:** Can return an object of any subclass (e.g., `EnumSet.of()` returns `RegularEnumSet` or `JumboEnumSet`).

---

### Question 11: What is the Object Pool Pattern?

**Answer:**
A pattern where a set of initialized objects is kept ready to use used, rather than allocating and destroying them on demand.

**Key Concepts:**
- **Acquire:** Client gets an object from the pool.
- **Release:** Client returns the object to the pool.

**Use Case:**
Database Connection Pools (HikariCP), Thread Pools. Creating connections/threads is expensive; reusing them is cheap.

---

### Question 12: What is the difference between Eager and Lazy Loading?

**Answer:**
- **Eager:** The object is created as soon as the class loads (static field). Fast runtime access, but slows down startup and wastes memory if unused.
- **Lazy:** The object is created only when requested (`if null then create`). Faster startup, saves memory, but requires thread-safety checks.

---

### Question 13: What is "Dependency Injection"?

**Answer:**
It is a form of Inversion of Control (IoC). A class does not create its dependencies; they are provided ("injected") to it.

**Key Concepts:**
- **Constructor Injection:** Pass dependencies in constructor (Recommended).
- **Setter Injection:** Pass dependencies via setters.

**Code:**
```java
// Without DI
class Car { Engine e = new GasEngine(); }

// With DI
class Car {
    Engine e;
    Car(Engine e) { this.e = e; } 
}
```

---

### Question 14: What is the Monostate Pattern?

**Answer:**
A weird variation o Singleton. The class has a normal constructor, but all its fields are `static`.

**Result:**
You can create 10 instances of `Monostate`, but they all share the same data. It behaves like a Singleton but looks like a normal class.

**Use Case:**
Rarely used. Singleton is generally preferred for clarity.

---

### Question 15: What is the Registry of Singletons?

**Answer:**
A pattern used when you want a flexible way to manage singletons. Instead of a hardcoded class, you utilize a Map (Registry) to store singleton instances by name.

**Code:**
```java
Map<String, Object> registry = new HashMap<>();
registry.put("Logger", new Logger());
registry.put("Config", new Config());
```

**Use Case:**
Spring's ApplicationContext (BeanFactory) is essentially a giant Registry of Singletons.

---

## 🧱 Structural Patterns (Questions 16-30)

### Question 16: What is the Adapter Pattern?

**Answer:**
It allows incompatible interfaces to work together. It wraps an existing class with a new interface.

**Key Concepts:**
- **Target:** The interface the client expects.
- **Adaptee:** The existing class with incompatible interface.
- **Adapter:** The bridge class.

**Use Case:**
Integrating a 3rd party library that uses `xml` input into your system that uses `json`.

---

### Question 17: What is the Decorator Pattern?

**Answer:**
Attaches additional responsibilities to an object dynamically. It provides a flexible alternative to subclassing for extending functionality.

**Key Concepts:**
- **Wrapper:** It wraps the original object.
- **Transparency:** It implements the same interface as the wrapped object.

**Code:**
```java
Coffee c = new SimpleCoffee();
c = new Milk(c); // Decorator
c = new Sugar(c); // Decorator
System.out.println(c.cost()); // Cost of Coffee + Milk + Sugar
```

**Use Case:**
Java I/O Classes (`BufferedInputStream` wraps `FileInputStream`).

---

### Question 18: What is the Facade Pattern?

**Answer:**
Provides a simplified interface to a library, a framework, or any other complex set of classes.

**Key Concepts:**
- **Masking Complexity:** Hides the complex logic of subsystems.
- **Loose Coupling:** Client only talks to Facade.

**Code:**
```java
homeTheater.watchMovie("Matrix"); 
// Internally calls: lights.dim(), projector.on(), amp.on(), dvd.play()
```

**Use Case:**
A "Service Layer" in a web app often acts as a Facade for DAOs and external APIs.

---

### Question 19: Adapter vs Facade?

**Answer:**
- **Adapter:** Makes **incompatible** interfaces work together. (My plug doesn't fit).
- **Facade:** Makes **complex** interfaces easier to use. (The interface is fine, but there are too many buttons).

---

### Question 20: What is the Proxy Pattern?

**Answer:**
Provides a surrogate or placeholder for another object to control access to it.

**Types:**
- **Virtual Proxy:** Lazy loading (load image only when scroll into view).
- **Protection Proxy:** Access control (check admin rights).
- **Remote Proxy:** Represents object in a different address space (RMI).

**Use Case:**
Hibernate lazy loading (the entity you get is a Proxy, DB call happens only when you call `.getOrders()`).

---

### Question 21: Proxy vs Decorator?

**Answer:**
- **Proxy:** Controls **Access**. The relationship is usually 1-to-1 and fixed (You don't wrap a Proxy in a Proxy).
- **Decorator:** Adds **Behavior**. You often chain multiple decorators.

---

### Question 22: What is the Bridge Pattern?

**Answer:**
Decouples an abstraction from its implementation so that the two can vary independently.

**Key Concepts:**
- **Abstraction:** The high-level logic (e.g., `RemoteControl`).
- **Implementation:** The low-level logic (e.g., `TV`, `Radio`).
- **Bridge:** The field `protected Device device;` inside the Remote.

**Use Case:**
When you have a class hierarchy that is exploding in two dimensions (e.g., Shape x Color). Bridge splits it into `Shape` hierarchy and `Color` hierarchy.

---

### Question 23: What is the Composite Pattern?

**Answer:**
Composes objects into tree structures to represent part-whole hierarchies. It lets clients treat individual objects and compositions uniformly.

**Key Concepts:**
- **Leaf:** Basic element (File).
- **Composite:** Container (Folder).
- **Component:** Common Interface (FileSystemItem).

**Use Case:**
GUI definitions (Panel contains Buttons and other Panels).

---

### Question 24: What is the Flyweight Pattern?

**Answer:**
Used to minimize memory usage by sharing as much data as possible with other similar objects.

**Key Concepts:**
- **Intrinsic State:** Shared, read-only data (e.g., Shape of bullet).
- **Extrinsic State:** Unique data passed in arguments (e.g., x,y position).

**Use Case:**
Rendering text (don't create object for every letter 'a'), Games (particles, trees).

---

### Question 25: What is the Private Class Data Pattern?

**Answer:**
A pattern that secures the data within a class by blocking setter access. The data is encapsulated in a data class that is created once in the constructor of the main class.

**Use Case:**
Preventing manipulation of immutable logic configuration.

---

### Question 26: What is the Marker Interface Pattern?

**Answer:**
An interface with **no methods**. It is used to "mark" a class so that the runtime environment can treat it specially.

**Examples:**
- `Serializable`
- `Cloneable`
- `Remote`

*Note: Annotations (e.g., `@Cachable`) have largely replaced Marker Interfaces.*

---

### Question 27: What is the Transfer Object Pattern (DTO)?

**Answer:**
A simple POJO (Plain Old Java Object) used to transfer data between software layers or over the network. It has no behavior, just data.

**Goal:**
Reduce the number of remote calls. Instead of calling `user.getName()`, `user.getAge()`, etc., you get one `UserDTO` with all data.

---

### Question 28: What is the DAO Pattern?

**Answer:**
**Data Access Object**. It abstracts and encapsulates all access to the data source.

**Structure:**
- **DAO Interface:** Defines operations (`findById, save`).
- **DAO Implementation:** Valid SQL/Hibernate code.
- **Model:** The object being saved.

**Goal:**
Separation of Logic. The Business Service doesn't know *how* data is saved (SQL vs File vs Cloud).

---

### Question 29: What is the Front Controller Pattern?

**Answer:**
A centralized request handling mechanism. All requests go through a single handler (Controller) first.

**Use Case:**
Spring MVC `DispatcherServlet`. It handles security, internationalization, and routing before passing to specific controllers.

---

### Question 30: What is the Intercepting Filter Pattern?

**Answer:**
Used to preprocess/postprocess requests. Filters are defined and applied to requests before they reach the target application.

**Use Case:**
Authentication, Logging, Encoding, Compression (GZip).

---

## 🧠 Behavioral Patterns (Questions 31-50)

### Question 31: What is the Strategy Pattern?

**Answer:**
Defines a family of algorithms, encapsulates each one, and makes them interchangeable.

**Code:**
```java
context.setStrategy(new CreditCardStrategy());
context.pay();
```

**Use Case:**
Sorting (BubbleSort vs QuickSort), Payment Methods, Route Calculation.

---

### Question 32: What is the Observer Pattern?

**Answer:**
Defines a one-to-many dependency so that when one object changes state, all its dependents are notified automatically.

**Key Concepts:**
- **Subject:** The publisher.
- **Observer:** The subscriber.

**Use Case:**
Event listeners in UI, News Feeds, JMS (Topic).

---

### Question 33: What is the Command Pattern?

**Answer:**
Encapsulates a request as an object. This allows parameterizing clients with requests, queueing requests, and supporting undo.

**Key Concepts:**
- **Invoker:** Remote Control.
- **Command:** Object (`LightOnCommand`).
- **Receiver:** Light.

**Use Case:**
Task Queues, Undo/Redo operations.

---

### Question 34: What is the Chain of Responsibility Pattern?

**Answer:**
Passes a request along a chain of handlers. Each handler allows processing or passing to the next.

**Code:**
```java
logger.log(ERROR, "Fail"); 
// ConsoleLogger -> FileLogger -> EmailLogger
```

**Use Case:**
Exception Handling, Middleware chains in Web Servers (Express.js/Spring Security).

---

### Question 35: What is the Template Method Pattern?

**Answer:**
Defines the skeleton of an algorithm in a method, deferring some steps to subclasses.

**Key Concepts:**
- **Final Method:** The algorithm structure (cannot be changed).
- **Abstract Methods:** The steps to be implemented by child classes.

**Use Case:**
Frameworks! `super.onCreate()` used in Android. Life-cycle hooks.

---

### Question 36: Template Method vs Strategy?

**Answer:**
- **Template Method:** Uses **Inheritance**. Can only change *parts* of the algorithm.
- **Strategy:** Uses **Composition**. Can replace the *entire* algorithm.

---

### Question 37: What is the State Pattern?

**Answer:**
Allows an object to alter its behavior when its internal state changes.

**Key Concepts:**
- **Context:** The object (Vending Machine).
- **State:** Interface.
- **Concrete State:** Logic for that state (`NoCoinState`, `HasCoinState`).

**Use Case:**
Vending Machine, TCP Connection (Open, Closed, Listening), Game Character State (Jumping, Running).

---

### Question 38: What is the Iterator Pattern?

**Answer:**
Provides a way to access elements of a collection sequentially without exposing the underlying representation.

**Key Concepts:**
- Uniform interface (`hasNext()`, `next()`) for List, Set, Map, Tree.

**Use Case:**
`java.util.Iterator`. Makes `for-each` loops possible.

---

### Question 39: What is the Mediator Pattern?

**Answer:**
Reduces chaotic dependencies between objects. Instead of Object A talking to B, C, and D, it talks to the Mediator, which routes the message.

**Use Case:**
Chat Room (User sends msg to Room, Room sends to all), Air Traffic Control.

---

### Question 40: What is the Memento Pattern?

**Answer:**
Captures and externalizes an object's internal state so that it can be restored later.

**Key Concepts:**
- **Originator:** The object.
- **Memento:** The snapshot.
- **Caretaker:** The history keeper.

**Use Case:**
"Ctrl+Z" (Undo) functionality.

---

### Question 41: What is the Visitor Pattern?

**Answer:**
Lets you define a new operation without changing the classes of the elements on which it operates.

**Structure:**
- **Element:** `accept(Visitor v)`
- **Visitor:** `visit(ElementA e)`, `visit(ElementB e)`

**Use Case:**
You have a fixed structure (e.g., XML Tree) but want to perform different, new operations on it (Export to JSON, Validate, Count Tags).

---

### Question 42: What is the Null Object Pattern?

**Answer:**
Instead of returning `null` (which causes NPE), return a default object that does nothing.

**Code:**
```java
// Instead of null, return this:
class NullLogger implements Logger {
    public void log(String msg) { /* Do nothing */ }
}
```

**Benefit:**
Removes the need for `if (logger != null)` checks everywhere.

---

### Question 43: What is the Interpreter Pattern?

**Answer:**
Given a language, define a representation for its grammar and an interpreter that uses the representation to interpret sentences.

**Use Case:**
SQL Parsing, Symbol processing engines, Regular Expressions.

---

### Question 44: What is the Service Locator Pattern?

**Answer:**
An old pattern used to locate services (Database, Messaging) using a central registry.

**Note:**
Considered an **Anti-Pattern** in modern dev because it hides dependencies. Dependency Injection is preferred.

---

### Question 45: What is the MVC Pattern?

**Answer:**
**Model-View-Controller**. Separation of concerns.
- **Model:** Data/Business Logic.
- **View:** UI presentation.
- **Controller:** Handling input and updating Model/View.

---

### Question 46: What is MVVM?

**Answer:**
**Model-View-ViewModel**. Used in modern UIs (React, Angular, WPF).
- **ViewModel:** Exposes data streams relevant to the View.
- **Binding:** The View binds to the ViewModel. Updates are automatic (Two-way binding).

---

### Question 47: What is CQRS (Command Query Responsibility Segregation)?

**Answer:**
Separating the Read (Query) and Write (Command) models of a system.
- **Command:** Updates data (Complex validation).
- **Query:** Reads data (Optimized for speed, maybe different DB).

---

### Question 48: What is Repository Pattern?

**Answer:**
Mediates between the domain and data mapping layers using a collection-like interface for accessing domain objects. It hides the complexity of query logic.

---

### Question 49: What is Inversion of Control (IoC)?

**Answer:**
The principle that the control flow of a program is inverted: instead of the programmer controlling the flow, the framework calls the programmer's code. Dependency Injection is a specific type of IoC.

---

### Question 50: What is the difference between Cohesion and Coupling?

**Answer:**
- **Cohesion (Good):** How related the functions within a single module are. (High Cohesion = Single Responsibility).
- **Coupling (Bad):** How dependent modules are on each other. (Low Coupling = Easy to change one without breaking another).
