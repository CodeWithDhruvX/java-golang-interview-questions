# Structural Patterns - Service Based Companies

## 1. What are Structural Design Patterns? Give some examples.

**Answer:**
Structural design patterns explain how to assemble objects and classes into larger structures, while keeping these structures flexible and efficient. They are concerned with how classes and objects are composed to form larger structures.

**Examples include:**
*   **Adapter:** Allows objects with incompatible interfaces to collaborate.
*   **Bridge:** Lets you split a large class or a set of closely related classes into two separate hierarchies—abstraction and implementation—which can be developed independently.
*   **Composite:** Lets you compose objects into tree structures and then work with these structures as if they were individual objects.
*   **Decorator:** Lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.
*   **Facade:** Provides a simplified interface to a library, a framework, or any other complex set of classes.
*   **Proxy:** Provides a substitute or placeholder for another object. A proxy controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.

---

## 2. Explain the Adapter Design Pattern. Give a real-life example.

**Answer:**
The **Adapter Pattern** works as a bridge between two incompatible interfaces. This type of design pattern comes under structural pattern as this pattern combines the capability of two independent interfaces.

It involves a single class called Adapter which joins functionalities of independent or incompatible interfaces.

**Real-life Example:**
Consider a card reader which acts as an adapter between a memory card and a laptop. You plugin the memory card into card reader and card reader into the laptop so that memory card can be read via laptop.

**Software Example:**
If an existing system expects an interface `XMLDataFormat`, but you want to integrate a 3rd party library that returns `JSONDataFormat`, you can create an `XMLToJSONAdapter` that implements `XMLDataFormat` but translates operations for the `JSONDataFormat` library.

---

## 3. What is the Decorator Pattern? How is it different from inheritance?

**Answer:**
The **Decorator Pattern** allows behavior to be added to an individual object, dynamically or statically, without affecting the behavior of other objects from the same class. It wrap objects within special wrapper objects that contain the new behaviors.

**Difference from Inheritance:**
*   **Static vs Dynamic:** Inheritance extends behavior statically at compile-time. All instances of the derived class get the new behavior. Decorator extends behavior dynamically at runtime, allowing you to wrap an individual object with multiple decorators without defining entirely new subclasses.
*   **Flexibility:** Inheritance can lead to an explosion of subclasses for every combination of features (e.g., `WindowWithBorderAndScrollbar`, `WindowWithScrollbarAndMenu`). Decorators allow creating combinations gracefully by wrapping an object dynamically.
*   **Is-A vs Has-A:** Inheritance creates an "is-a" relationship (A `Dog` is an `Animal`). Decorator relies on a "has-a" or "wraps" relationship. The decorator implements the same interface as the wrapped object but *contains* an instance of that wrapped object.

---

## 4. Give an example or use case for the Facade Pattern.

**Answer:**
The **Facade Pattern** hides the complexities of the system and provides an interface to the client using which the client can access the system. It adds a higher-level interface that makes the subsystem easier to use.

**Use Case:**
Imagine you are writing a multimedia application. To start playing a video, your code might need to initialize a video decoder, an audio decoder, read the file format, allocate memory buffers, and synchronize audio/video playback. Exposing all these complex classes to an end client (e.g., standard UI layer) creates tight coupling.

A `VideoPlayerFacade` class can provide a simple `playFile(String filename)` method to the client. The facade handles the complexity of interacting with the various decoder and buffer classes internally. The client doesn't need to know the intricate initialization sequence.

---

## 5. What is the Proxy Design Pattern? What are its common use cases?

**Answer:**
The **Proxy Pattern** dictates that a class represents functionality of another class. It provides a surrogate or placeholder for another object to control access to it.

**Common Use Cases:**
*   **Virtual Proxy:** Used for lazy initialization. Creates expensive objects only when they are first needed. For example, loading a large high-resolution image only when it actually needs to be rendered on the screen.
*   **Protection Proxy:** Used for authorization control. Controls access to the original object based on the client's credentials or permissions. E.g., An admin proxy that allows access to sensitive data only to admin users.
*   **Remote Proxy:** Provides a local representation for an object that is present in a different address space (Remote JVM/Machine). Example: The stub in RMI or client-side proxies for web services.
*   **Caching Proxy:** Provides temporary storage for results of expensive operations related to the subject.

---

## 6. Compare Proxy vs Decorator patterns.

**Answer:**
Both patterns have a similar structure: they wrap an object and forward requests to it.

| Feature | Proxy Pattern | Decorator Pattern |
| :--- | :--- | :--- |
| **Intent / Purpose** | *Controls access* to an object. (Lazy loading, security, remote execution). | *Adds behaviors or responsibilities* to an object dynamically. |
| **Instantiation** | The Proxy often instantiates the real subject itself (e.g., Lazy initialization) or its lifecycle is tied to the Proxy. | The Decorator is typically instantiated with an existing object passed into its constructor. The client decides what to decorate. |
| **Visibility** | The client might not even know it's interacting with a proxy; it thinks it's the real object. | The client explicitly creates the Decorator and wrapped object combination. |
