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

**Spoken Format:**

"Structural design patterns are like architects' blueprints for organizing your code structure. They help you put objects and classes together to form larger, more complex structures while keeping everything flexible and efficient. Think of them as different ways to connect and organize your code components - some patterns help incompatible interfaces work together like Adapter, while others simplify complex systems like Facade. The main goal is to make sure your code structure remains clean and maintainable even as your system grows."

"In enterprise Java projects, these patterns are essential for maintaining clean architecture. When I worked on a banking modernization project at Infosys, we had to integrate a new Spring Boot service with a 15-year-old legacy system that exposed SOAP services. The Adapter pattern was crucial here - we created a SOAPToRESTAdapter that made the old SOAP services look like modern REST endpoints to our new application. This allowed us to gradually migrate without rewriting the legacy system."

"The Facade pattern is something I use in almost every enterprise project. When building a customer management system for a retail client, we had multiple services - CustomerService, OrderService, PaymentService, NotificationService. The UI team didn't want to deal with all these services individually, so we created a CustomerManagementFacade that provided simple methods like 'createCustomerWithOrder()' which internally coordinated with all the backend services. This made the frontend code much cleaner and reduced the number of API calls they had to make."

---

## 2. Explain the Adapter Design Pattern. Give a real-life example.

**Answer:**
The **Adapter Pattern** works as a bridge between two incompatible interfaces. This type of design pattern comes under structural pattern as this pattern combines the capability of two independent interfaces.

It involves a single class called Adapter which joins functionalities of independent or incompatible interfaces.

**Real-life Example:**
Consider a card reader which acts as an adapter between a memory card and a laptop. You plugin the memory card into card reader and card reader into the laptop so that memory card can be read via laptop.

**Software Example:**
If an existing system expects an interface `XMLDataFormat`, but you want to integrate a 3rd party library that returns `JSONDataFormat`, you can create an `XMLToJSONAdapter` that implements `XMLDataFormat` but translates operations for the `JSONDataFormat` library.

**Spoken Format:**

"The Adapter pattern is like a universal adapter plug that lets you connect devices with incompatible plugs. In software, when you have two systems that can't talk to each other because they have different interfaces, the Adapter pattern acts as a translator between them. Think of a real-world example - a card reader that acts as an adapter between your memory card and laptop. The memory card speaks one language, the laptop speaks another, and the card reader translates between them. In code, if you have a system that expects XML data but you want to use a library that provides JSON data, you create an adapter that implements the XML interface but internally uses the JSON library and converts the data as needed."

"I had a great experience with the Adapter pattern on a project for a government client. They had a legacy mainframe system that exposed data through a proprietary protocol, but we were building a modern web application that needed JSON APIs. We created a MainframeAdapter that implemented our modern REST interfaces but internally communicated with the mainframe using its proprietary format. The beauty was that our web application didn't need to know anything about the mainframe - it just called standard REST endpoints, and the adapter handled all the complexity."

"What's really useful about Adapter is that it lets you integrate new systems with existing ones without changing the existing code. This is crucial in enterprise environments where you can't just rewrite everything. The adapter acts as a bridge, allowing gradual migration and coexistence of old and new systems. This is something service companies like TCS and Wipro deal with all the time - they're often building new systems that need to work with clients' existing infrastructure."

---

## 3. What is the Decorator Pattern? How is it different from inheritance?

**Answer:**
The **Decorator Pattern** allows behavior to be added to an individual object, dynamically or statically, without affecting the behavior of other objects from the same class. It wrap objects within special wrapper objects that contain the new behaviors.

**Difference from Inheritance:**
*   **Static vs Dynamic:** Inheritance extends behavior statically at compile-time. All instances of the derived class get the new behavior. Decorator extends behavior dynamically at runtime, allowing you to wrap an individual object with multiple decorators without defining entirely new subclasses.
*   **Flexibility:** Inheritance can lead to an explosion of subclasses for every combination of features (e.g., `WindowWithBorderAndScrollbar`, `WindowWithScrollbarAndMenu`). Decorators allow creating combinations gracefully by wrapping an object dynamically.
*   **Is-A vs Has-A:** Inheritance creates an "is-a" relationship (A `Dog` is an `Animal`). Decorator relies on a "has-a" or "wraps" relationship. The decorator implements the same interface as the wrapped object but *contains* an instance of that wrapped object.

**Spoken Format:**

"The Decorator pattern is like adding accessories to your outfit - you can dynamically add different features without changing the base outfit. Unlike inheritance where you permanently extend functionality at compile time, decorators let you add behaviors at runtime. Think of it this way - with inheritance, you'd have to create separate classes like 'WindowWithBorder', 'WindowWithScrollbar', 'WindowWithBorderAndScrollbar' for every combination. But with decorators, you can start with a basic window and then wrap it with a border decorator, then wrap that with a scrollbar decorator, creating any combination you want dynamically. It's more flexible and avoids the class explosion problem."

"In an e-commerce project for a retail client, we used the Decorator pattern extensively for product features. We had a base Product class, and then decorators for different features - WarrantyDecorator, GiftWrapDecorator, ExpressShippingDecorator, DiscountDecorator. At checkout, based on what the customer selected, we would dynamically wrap the product with the appropriate decorators. The final price calculation would automatically include all the features without us having to create hundreds of product subclasses."

"The pattern really shines in Java I/O classes, which is a perfect example to mention in interviews. The BufferedInputStream is a decorator that adds buffering to any InputStream. You can wrap it with a DataInputStream to add data type reading functionality, and then wrap that with a GZIPInputStream to add compression. Each layer adds functionality without changing the underlying stream. This is exactly how enterprise applications work - you start with basic functionality and keep adding features through decorators as requirements evolve."

---

## 4. Give an example or use case for the Facade Pattern.

**Answer:**
The **Facade Pattern** hides the complexities of the system and provides an interface to the client using which the client can access the system. It adds a higher-level interface that makes the subsystem easier to use.

**Use Case:**
Imagine you are writing a multimedia application. To start playing a video, your code might need to initialize a video decoder, an audio decoder, read the file format, allocate memory buffers, and synchronize audio/video playback. Exposing all these complex classes to an end client (e.g., standard UI layer) creates tight coupling.

A `VideoPlayerFacade` class can provide a simple `playFile(String filename)` method to the client. The facade handles the complexity of interacting with the various decoder and buffer classes internally. The client doesn't need to know the intricate initialization sequence.

**Spoken Format:**

"The Facade pattern is like having a remote control for your home entertainment system - instead of dealing with multiple remotes for TV, sound system, and Blu-ray player, you have one simple interface. In software, when you have a complex subsystem with many classes and interactions, the Facade pattern provides a simple, high-level interface that hides all the complexity. Think of a video player - to play a video, you need to initialize video decoders, audio decoders, manage buffers, and synchronize everything. Instead of exposing all this complexity to the client, you create a facade with a simple 'playFile()' method that handles everything internally. The client gets a clean, easy-to-use interface while the facade manages all the complex interactions behind the scenes."

"When I was working on a banking application at Wipro, we had to integrate with multiple external services - credit bureau APIs, payment gateways, SMS services, email services. Each service had its own authentication, error handling, and retry logic. The frontend team was struggling with all this complexity, so we created a BankingServicesFacade that provided simple methods like 'checkCreditScore()' or 'sendPaymentNotification()'. The facade handled all the authentication, error handling, and retries internally. This made the frontend code much cleaner and reduced the number of support tickets we got for integration issues."

"What's great about Facade is that it doesn't force you to use the simplified interface - advanced clients can still access the underlying subsystem directly if they need to. This gives you the best of both worlds - simplicity for most use cases, but flexibility when you need it. In enterprise applications, this is especially valuable because different teams have different needs - the UI team wants simple APIs, while the integration team might need direct access to individual services for custom workflows."

---

## 5. What is the Proxy Design Pattern? What are its common use cases?

**Answer:**
The **Proxy Pattern** dictates that a class represents functionality of another class. It provides a surrogate or placeholder for another object to control access to it.

**Common Use Cases:**
*   **Virtual Proxy:** Used for lazy initialization. Creates expensive objects only when they are first needed. For example, loading a large high-resolution image only when it actually needs to be rendered on the screen.
*   **Protection Proxy:** Used for authorization control. Controls access to the original object based on the client's credentials or permissions. E.g., An admin proxy that allows access to sensitive data only to admin users.
*   **Remote Proxy:** Provides a local representation for an object that is present in a different address space (Remote JVM/Machine). Example: The stub in RMI or client-side proxies for web services.
*   **Caching Proxy:** Provides temporary storage for results of expensive operations related to the subject.

**Spoken Format:**

"The Proxy pattern is like having a personal assistant who controls access to you. The proxy acts as a stand-in or placeholder for another object, controlling who can access it and when. Think of different scenarios - a virtual proxy is like lazy loading, where you only create expensive objects when they're actually needed. A protection proxy is like a security guard that checks permissions before letting someone access sensitive data. A remote proxy is like having a local representative for someone in another country - it handles communication with the actual object that might be in a different server or system. The key is that the proxy sits between the client and the real object, adding some extra functionality like security, caching, or lazy loading."

"In a financial services project, we used the Proxy pattern extensively for security and performance. We had sensitive customer data that could only be accessed by authorized users. We created a SecureCustomerDataProxy that implemented the same interface as the real CustomerData service, but before forwarding any request, it would check the user's permissions using Spring Security. If the user wasn't authorized, it would throw an exception. The client code didn't need to know about this security check - it just called the proxy methods like it would call the real service."

"We also used virtual proxies for performance optimization. In a reporting application, some reports required loading millions of records from the database, which was expensive. We created a ReportProxy that initially returned a basic report structure quickly, and only loaded the full data when the user actually requested it. This made the application feel much faster. The proxy pattern is essential in enterprise applications where you need to add cross-cutting concerns like security, caching, or lazy loading without changing the core business logic."

---

## 6. Compare Proxy vs Decorator patterns.

**Answer:**
Both patterns have a similar structure: they wrap an object and forward requests to it.

| Feature | Proxy Pattern | Decorator Pattern |
| :--- | :--- | :--- |
| **Intent / Purpose** | *Controls access* to an object. (Lazy loading, security, remote execution). | *Adds behaviors or responsibilities* to an object dynamically. |
| **Instantiation** | The Proxy often instantiates the real subject itself (e.g., Lazy initialization) or its lifecycle is tied to the Proxy. | The Decorator is typically instantiated with an existing object passed into its constructor. The client decides what to decorate. |
| **Visibility** | The client might not even know it's interacting with a proxy; it thinks it's the real object. | The client explicitly creates the Decorator and wrapped object combination. |

**Spoken Format:**

"While Proxy and Decorator patterns look similar structurally - both wrap an object - they serve very different purposes. Think of a proxy as a gatekeeper that controls access to the real object, while a decorator is like adding accessories to enhance the object's capabilities. A proxy's main job is to manage access - whether for security, lazy loading, or remote communication. The client might not even know they're talking to a proxy. But a decorator's job is to add new behaviors - the client specifically chooses which decorators to apply to enhance the object's functionality. It's like the difference between a security guard (proxy) who controls who enters a building, versus an interior designer (decorator) who adds decorations to make the rooms more functional."

"This distinction became really clear to me when I was working on an enterprise application that needed both patterns. We had a sensitive document service that required both security and additional functionality. We created a SecurityProxy that checked user permissions before allowing access to documents. Then we created decorators like AuditDecorator (to log access), CacheDecorator (to improve performance), and FormatDecorator (to convert documents to different formats). The proxy was always the outermost layer - security first, then we could add any combination of decorators based on what the client needed."

"In interviews, I always explain it this way: if you're controlling access, it's a proxy. If you're adding behavior, it's a decorator. A proxy protects and manages, a decorator enhances and extends. Both are valuable in enterprise applications - proxies for cross-cutting concerns like security and caching, decorators for extending functionality without modifying existing code. Understanding when to use each pattern shows that you think about the intent behind the pattern, not just the structure."
