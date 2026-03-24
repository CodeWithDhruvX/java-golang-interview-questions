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

**Spoken Format:**

"Creational design patterns are basically blueprints that help us create objects in a better way. They make our code more flexible and reusable by separating the object creation logic from the actual business logic. Think of them as different ways to 'birth' objects - some patterns ensure only one object exists, like Singleton, while others help create families of related objects, like Abstract Factory. The main goal is to make our system independent of how objects are created, composed, and represented."

"In my experience working on enterprise Java projects at TCS, we use these patterns all the time. For a banking application, we used Singleton for the database connection manager - you don't want 50 database connections when you only need one. When building an e-commerce platform for a retail client, we used Factory Method to create different payment processors based on user selection - CreditCardFactory, PayPalFactory, UPIFactory. The Abstract Factory came in handy when we needed to create complete theme families - a DarkThemeFactory would create dark buttons, dark panels, dark text, all working together."

"The Builder pattern was a lifesaver when we had complex objects with many optional parameters. Instead of having constructors with 10 parameters, we used builders like 'new OrderBuilder().setCustomerId().setItems().setShippingAddress().build()'. This made the code much more readable and maintainable, especially when new developers joined the project."

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

**Spoken Format:**

"The Singleton pattern is all about ensuring that only one instance of a class exists throughout the entire application. Think of it like having only one database connection manager or one logging service in your whole system. To implement this, we make the constructor private so no one can create new instances, we create a private static variable to hold the single instance, and we provide a public static method called getInstance() that returns this single instance. For thread safety, we use double-checked locking with the volatile keyword to ensure that multiple threads can't create multiple instances at the same time."

"In a real enterprise project I worked on for a banking client, we used Singleton for the configuration manager. The application needed to read configuration properties from a file and make them available throughout the application. Using Singleton ensured that the configuration was loaded only once and cached properly. We also used it for the logging framework - you don't want multiple loggers writing to the same file simultaneously, that would create a mess."

"The key thing I've learned is to be careful with Singleton in distributed environments. If you're running multiple instances of your application behind a load balancer, each instance will have its own Singleton. For true singleton behavior across multiple JVMs, you need to use external stores like Redis or database locks. This is something that often comes up in service company interviews when they ask about scalability."

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

**Spoken Format:**

"The key difference between Singleton and Static class is that Singleton is actually an object that follows object-oriented principles, while a static class is just a collection of static methods. With Singleton, you get all the benefits of OOP - you can implement interfaces, extend other classes, and maintain state properly. A Singleton can be created lazily when needed, but a static class is loaded immediately when the class loader loads it. Think of Singleton as a smart, flexible object that happens to have only one instance, while a static class is more like a utility toolbox with static methods."

"In my experience, this distinction becomes crucial in enterprise applications. I worked on a project where we needed to implement a caching service that had to follow a specific interface defined by the client's architecture team. We couldn't use a static class because static classes can't implement interfaces in Java. We had to use Singleton so it could implement the CacheService interface and be injected using Spring's dependency injection."

"Another practical difference I've encountered is with testing. Static classes are hard to mock in unit tests, but Singletons can be designed to be testable. You can create a test double or mock implementation of the Singleton interface. This is why in modern enterprise applications, we often prefer dependency injection over traditional Singleton - it gives us the singleton behavior but with better testability and flexibility."

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

**Spoken Format:**

"The Factory Design Pattern is like having a smart factory that knows how to create different types of objects based on your requirements. Instead of using 'new' keyword everywhere in your code, you ask the factory to create the object for you. This is super useful when you don't know which exact type of object you'll need at runtime, or when the creation logic is complex. For example, in a notification system, you might need to send an email, SMS, or push notification based on user preference. Instead of writing if-else statements everywhere, you create a NotificationFactory that returns the right type of notification object. This makes your code cleaner, more maintainable, and easier to extend with new notification types."

"In a project for an insurance client, we implemented a document generator that could create PDFs, Excel files, or Word documents based on the user's selection. The Factory pattern was perfect here - the UI would call DocumentFactory.createDocument('PDF') and get back a PDF document object. When the client later asked for HTML reports, we just added a new HTMLDocument class and updated the factory - no changes needed in the existing code. This follows the Open/Closed Principle that enterprise architects love to talk about."

"The real power comes when you combine Factory with dependency injection. In Spring Boot applications, we often create factory beans that can produce different implementations based on configuration. For example, based on a property file, the factory might return a MySQL implementation or an Oracle implementation of a repository interface. This makes it easy to switch between development and production environments without code changes."

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

**Spoken Format:**

"The main difference between Factory and Abstract Factory is that Factory Method creates just one type of object, while Abstract Factory creates families of related objects. Think of Factory Method as a single-purpose factory - like a car factory that only makes cars. But Abstract Factory is like a complete vehicle manufacturing plant that can make cars, bikes, and trucks - all belonging to the same brand family. So if you have a Honda factory, it can create Honda cars AND Honda bikes, ensuring they work well together. The key is that Abstract Factory maintains consistency - all objects created by the same factory belong to the same family and are designed to work together."

"I worked on a banking application that needed different UI themes for different banks. We used Abstract Factory for this. The SBIThemeFactory would create SBI buttons, SBI panels, SBI colors - all following SBI's branding guidelines. The ICICIThemeFactory would create ICICI buttons, ICICI panels, ICICI colors. The beauty was that all components from the same factory worked together perfectly - you couldn't accidentally mix an SBI button with ICICI colors because the factory ensured consistency."

"This pattern really shines in enterprise applications when you have multiple deployment environments. For example, a CloudFactory might create AWS services like S3Client, SQSClient, and DynamoDBClient for AWS deployment. The AzureFactory would create the corresponding Azure services like BlobClient, ServiceBusClient, and CosmosClient. The application code remains the same - only the factory implementation changes based on configuration. This makes it easy to deploy the same application to different cloud providers."

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

**Spoken Format:**

"The Builder pattern is perfect when you need to create complex objects with many optional parameters. Instead of having multiple constructors with different parameter combinations - which gets really messy - you use a builder that lets you construct the object step by step. Think of it like building a custom pizza - you start with the base (mandatory parts), then add toppings one by one (optional parts), and finally bake it. In code, instead of having a constructor with 10 parameters where some are null, you have a clean, readable chain of method calls like 'builder.setUrl().addHeader().setBody().build()'. This makes the code much more readable and maintainable, especially when dealing with objects that have many configuration options."

"In a microservices project for a retail client, we had to create complex HTTP requests to integrate with multiple third-party APIs. Each request had different requirements - some needed authentication headers, some needed query parameters, some needed request bodies. The Builder pattern was a game-changer here. We could write code like 'HttpRequestBuilder.url("https://api.vendor.com").header("Authorization", token).body(jsonPayload).timeout(5000).build()'. The code was self-documenting and new team members could understand it immediately."

"What I love about Builder is that it prevents the telescoping constructor problem. I once maintained a legacy system that had constructors like 'new User(String name, String email, String phone, String address, String city, String state, String country, String zip)' - it was a nightmare to use and maintain. We refactored it to use Builder, and the code became so much cleaner. Plus, Builder makes it easy to add new optional parameters later without breaking existing code - this is crucial in enterprise applications where requirements change frequently."
