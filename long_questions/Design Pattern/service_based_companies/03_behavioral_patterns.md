# Behavioral Patterns - Service Based Companies

## 1. What are Behavioral Design Patterns? Give examples.

**Answer:**
Behavioral design patterns are concerned with algorithms and the assignment of responsibilities between objects. They don't just specify patterns of objects or classes but also the patterns of communication between them.

**Examples include:**
*   **Observer:** Lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they're observing.
*   **Strategy:** Lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.
*   **Command:** Turns a request into a stand-alone object that contains all information about the request.
*   **Iterator:** Lets you traverse elements of a collection without exposing its underlying representation (list, stack, tree, etc.).
*   **State:** Lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.
*   **Template Method:** Defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.

---

## 2. Explain the Observer Design Pattern with a real-world scenario.

**Answer:**
The **Observer Pattern** is used when there is a one-to-many relationship between objects such that if one object is modified, its dependent objects are to be notified automatically. It falls under behavioral pattern category.

**Key Components:**
*   **Subject:** The object that holds the state and manages the observers. It provides methods to attach and detach observers.
*   **Observer:** Interface defining the `update()` method that is called by the Subject when its state changes.

**Real-World Scenario:**
*   **News Agency (Subject):** A news agency publishes daily news.
*   **Subscribers (Observers):** Multiple subscribers (e.g., users using mobile apps, email receivers, RSS feeds) want to be notified whenever new news is published. The news agency notifies all registered subscribers simultaneously without knowing who they are specifically.

**Software Scenario:**
In GUI programming, an `EventListener` (Observer) registered to a `Button` (Subject). When the button is clicked, all registered listeners are notified.

---

## 3. What is the Strategy Design Pattern? How does it promote flexibility?

**Answer:**
The **Strategy pattern** defines a family of algorithms, encapsulates each one, and makes them interchangeable at runtime. The algorithm can vary independently from clients that use it.

**How it promotes flexibility:**
Instead of hardcoding various behaviors or using large `if...else` or `switch` statements to pick an algorithm, the behavior is delegated to a separate strategy object. The context class holds a reference to a Strategy interface and delegates the work to it.

**Use Case & Flexibility:**
Imagine an e-commerce platform with a `PaymentProcessor`. You could have strategies like `CreditCardStrategy`, `PayPalStrategy`, `CryptocurrencyStrategy`.
At runtime, the user selects a payment method. The `PaymentProcessor` doesn't need to change its code; it just receives the specific strategy object and calls `pay()` on it. Adding a new payment method (e.g., `ApplePayStrategy`) only requires creating a new class implementing the strategy interface, adhering strictly to the Open/Closed Principle.

---

## 4. Compare State Pattern vs. Strategy Pattern.

**Answer:**
Structurally, they are almost identical – both delegate work to encapsulated helper objects. However, their intent determines how they are used.

| Feature | State Pattern | Strategy Pattern |
| :--- | :--- | :--- |
| **Intent** | An object changes its behavior when its *internal state* changes. | A client provides different *algorithms* to solve a specific problem. |
| **Awareness** | The context acts as a state machine. State classes *can be aware* of each other and handle transitions from one state to another. | Strategies are usually *independent* and unaware of other strategies. |
| **Client Role** | The client might not even know states exist. The object handles its state transitions internally based on events. | The client explicitly chooses the strategy to use and sets it in the context at runtime. |
| **Example** | A Media Player (States: Playing, Paused, Stopped). | Sorting algorithms for a list (Strategies: QuickSort, MergeSort, BubbleSort). |

---

## 5. What is the Command Pattern? What problems does it solve?

**Answer:**
The **Command Pattern** encapsulates a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.

**Key Components:**
*   **Command:** An interface with an `execute()` method.
*   **ConcreteCommand:** Implements the command interface and binds a receiver to an action.
*   **Client:** Creates the `ConcreteCommand` and sets its receiver.
*   **Invoker:** Holds the command and calls `execute()` on it.
*   **Receiver:** The object that actually performs the action.

**Problems Solved / Use Cases:**
*   **Decoupling:** Decouples the object that invokes the operation from the one that knows how to perform it.
*   **Queuing and Scheduling Tasks:** Commands can be serialized, added to a queue, and processed asynchronously by worker threads.
*   **Undo/Redo Functionality:** The command object can store its previous state and implement an `undo()` method, making it easy to reverse operations (e.g., text editors, photo editors).
*   **Macro Commands:** Multiple commands can be grouped together to execute a sequence of actions.
