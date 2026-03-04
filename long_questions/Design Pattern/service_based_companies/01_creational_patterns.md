# Creational Patterns - Service Based Companies

## 1. What are Creational Design Patterns? Give some examples.

**Answer:**
Creational design patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code. They help in making a system independent of how its objects are created, composed, and represented.

**Examples include:**
*   **Singleton:** Ensures a class has only one instance and provides a global point of access to it.
*   **Factory Method:** Defines an interface for creating an object, but lets subclasses decide which class to instantiate.
*   **Abstract Factory:** Provides an interface for creating families of related or dependent objects without specifying their concrete classes.
*   **Builder:** Separates the construction of a complex object from its representation so that the same construction process can create different representations.
*   **Prototype:** Specifies the kinds of objects to create using a prototypical instance, and creates new objects by copying this prototype.

---

## 2. Explain the Singleton Design Pattern and how to implement it in Java.

**Answer:**
The **Singleton pattern** restricts the instantiation of a class and ensures that only one instance of the class exists in the Java Virtual Machine (JVM).

**Key components:**
1.  **Private constructor:** To restrict instantiation of the class from other classes.
2.  **Private static variable:** Of the same class that is the only instance of the class.
3.  **Public static method:** That returns the instance of the class. This is the global access point.

**Implementation (Double-Checked Locking - Thread Safe):**

```java
public class MySingleton {
    // volatile ensures that multiple threads handle the singleton instance correctly
    private static volatile MySingleton instance;

    // Private constructor prevents instantiation from other classes
    private MySingleton() {
        // Optional: protect against reflection
        if (instance != null) {
            throw new RuntimeException("Use getInstance() method to get the single instance of this class.");
        }
    }

    public static MySingleton getInstance() {
        if (instance == null) { // Single checked
            synchronized (MySingleton.class) {
                if (instance == null) { // Double checked
                    instance = new MySingleton();
                }
            }
        }
        return instance;
    }
}
```

---

## 3. What is the difference between Singleton and Static Class?

**Answer:**

| Feature | Singleton Pattern | Static Class |
| :--- | :--- | :--- |
| **Instantiation** | Creates a single instance (object). | Cannot be instantiated at all. |
| **Object-Oriented** | Fully supports OOP principles like Inheritance and Polymorphism. | Cannot inherit state or behavior (no polymorphism). |
| **State Management** | Good for maintaining state; can have instance variables. | Only has static variables, which are global state (harder to manage). |
| **Interface Implementation**| A Singleton class can implement interfaces. | A Static class cannot implement interfaces. |
| **Lazy Loading** | Can be lazily initialized (created only when needed). | Initialized when the class is loaded by the ClassLoader. |
| **Serialization** | Can be serialized and deserialized. | Cannot be serialized in the standard way (since there's no object). |

---

## 4. What is the Factory Design Pattern? When do we use it?

**Answer:**
The **Factory Design Pattern** is a creational pattern that provides an interface or abstract class for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. It defines a method (the "factory" method) that is responsible for object creation, abstracting the `new` keyword away from the client code.

**When to use it:**
*   When a class doesn't know what sub-classes will be required to create.
*   When a class wants its sub-classes to specify the objects to be created.
*   When the creation logic involves complex logic (e.g., getting configuration from a file or environment, depending on input parameters).
*   To centralize and encapsulate object creation code, making the system loosely coupled and easier to extend.

**Example Scenario:**
A `NotificationFactory` that returns an `EmailNotification`, `SMSNotification`, or `PushNotification` object based on an input string without the client needing to know the specific implementation details of each notification type.

---

## 5. What is the difference between Factory and Abstract Factory patterns?

**Answer:**

**Factory Method:**
*   **Purpose:** Creates one specific type of object.
*   **Abstraction Level:** Single level. The factory defines a method to create an object, and subclasses implement that method to create specific instances.
*   **Focus:** Focuses on creating a single product.
*   **Example:** A `CarFactory` that returns a `Sedan` or an `SUV`.

**Abstract Factory:**
*   **Purpose:** Creates families of related or dependent objects. This is often called a "Factory of Factories."
*   **Abstraction Level:** Multi-level. An Abstract Factory interface declares a set of methods to create various related products. Concrete factories implement this interface to create a specific family of products.
*   **Focus:** Focuses on creating multiple related products that belong to a single family.
*   **Example:** A `VehicleFactory` that has methods `createCar()` and `createBike()`. A `HondaFactory` implements it returning `HondaCar` and `HondaBike`. A `ToyotaFactory` implements it returning `ToyotaCar` and `ToyotaBike`.

---

## 6. What is the Builder Design Pattern?

**Answer:**
The **Builder pattern** is a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code.

**Problem it solves:**
Imagine an object with many optional parameters in its constructor (`Telescoping Constructor Anti-Pattern`). This makes the constructor difficult to read and use.

**Solution:**
The Builder pattern extracts the object construction code out of its own class and moves it to separate objects called *builders*. The builder provides a set of methods to configure the object piece by piece. Finally, it provides a `build()` method that returns the fully constructed object.

**Use Case:**
Building an `HttpRequest` object where URL is mandatory, but headers, body, query parameters, and timeout are optional.
