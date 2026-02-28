# Low-Level System Design (LLD) / Machine Coding in Java

Product-based companies (Amazon, Swiggy, Flipkart, Uber) usually have a dedicated 90-120 minute "Machine Coding" or "Low-Level Design" round. You are expected to write working, compilable code that demonstrates OOD (Object-Oriented Design) principles, Design Patterns, and clean code practices.

---

## 🛠️ 1. Core Principles (The Evaluators)

Interviewers don't just look for a working solution; they evaluate *how* you write it.

### SOLID Principles
You must explicitly mention and use these during the interview.
1.  **Single Responsibility Principle (SRP):** A class should have one, and only one, reason to change (e.g., a `PaymentProcessor` shouldn't format email receipts).
2.  **Open/Closed Principle (OCP):** Software entities should be open for extension but closed for modification (e.g., use an `interface PaymentStrategy` instead of a massive `switch` statement for CreditCard, UPI, PayPal).
3.  **Liskov Substitution Principle (LSP):** Subtypes must be substitutable for their base types without altering the correctness of the program (e.g., if a `Penguin` extends `Bird`, but `Bird` has a `fly()` method, you violate LSP).
4.  **Interface Segregation Principle (ISP):** Clients should not be forced to depend on interfaces they do not use (e.g., instead of one massive `IMachine` interface with `print()`, `scan()`, `fax()`, split them into `IPrinter`, `IScanner`).
5.  **Dependency Inversion Principle (DIP):** High-level modules should not depend on low-level modules; both should depend on abstractions (e.g., inject a `DatabaseConnection` interface into a `UserService`, not a concrete `MySQLConnection`).

### Key Design Patterns (Must-Know for LLD)
*   **Strategy Pattern:** For interchangeable algorithms (e.g., different pricing strategies, payment methods).
*   **Factory Pattern:** For abstracting object creation (e.g., `VehicleFactory.create("CAR")`).
*   **Observer Pattern:** For publish-subscribe mechanisms (e.g., notifying users when a product is back in stock).
*   **Singleton Pattern:** For centralized resources (e.g., Database Connection Pool, Configuration Manager). Understand thread-safe implementation (Double-Checked Locking).
*   **Decorator Pattern:** For dynamically adding responsibilities to an object (e.g., adding extra cheese/toppings to a `BasePizza`).
*   **Builder Pattern:** For constructing complex objects step-by-step (e.g., building a `HttpRequest` with headers, body, params).

---

## 🚀 2. Top 10 LLD Interview Questions (Machine Coding)

Here are the classic LLD problems. You must be able to design the classes, interfaces, enums, and relationships for these scenarios.

### 1. Design a Parking Lot
*   **Focus:** Managing state, calculating fees, finding available spots.
*   **Key Entities:** `ParkingLot` (Singleton), `ParkingFloor`, `ParkingSpot` (Enum: COMPACT, LARGE, MOTORBIKE), `Vehicle` (Car, Truck, Bike), `Ticket`, `PaymentStrategy`.
*   **Challenges:** Concurrency (what if two cars enter simultaneously aiming for the last spot?), multi-level parking efficiency.

### 2. Design Amazon / E-commerce System
*   **Focus:** Inventory management, shopping cart logic, order fulfillment, payment processing.
*   **Key Entities:** `User` (Customer, Admin), `Product`, `Category`, `Inventory`, `Cart`, `Order`, `OrderLog`, `Payment`.
*   **Challenges:** Handling race conditions during checkout (Optimistic Locking/Pessimistic Locking).

### 3. Design Swiggy / Zomato / Food Delivery App
*   **Focus:** Matching customers, restaurants, and delivery executives.
*   **Key Entities:** `Restaurant`, `Menu`, `MenuItem`, `Order`, `DeliveryPartner`, `Customer`.
*   **Challenges:** Updating order status in real-time (Observer Pattern), strategy for assigning the nearest delivery partner.

### 4. Design BookMyShow / Movie Ticket Booking
*   **Focus:** Managing theater seating and preventing double-booking.
*   **Key Entities:** `Cinema`, `Screen`, `Show`, `Seat` (Enum: SILVER, GOLD, PLATINUM), `Booking`.
*   **Challenges:** The critical section: `lockSeat()` mechanism involving Database Transactions with Isolation Level Serializable or row-level `FOR UPDATE` locks.

### 5. Design a Vending Machine
*   **Focus:** State management.
*   **Key Entities:** `VendingMachine` (Context), `State` Interface (HasMoneyState, DispensingState, NoMoneyState), `Item`, `Coin`, `Inventory`.
*   **Challenges:** Implementing the **State Design Pattern** perfectly to handle transitions (e.g., what happens if you cancel after inserting money?).

### 6. Design an Elevator System
*   **Focus:** Algorithms for efficient dispatching.
*   **Key Entities:** `ElevatorController`, `ElevatorCar`, `Button` (Inside, Outside), `Display`, `Direction` (Enum: UP, DOWN, IDLE).
*   **Challenges:** Request queueing, optimizing wait times (SCAN algorithm / Elevator algorithm).

### 7. Design a Cache System (In-Memory)
*   **Focus:** Eviction policies and concurrency.
*   **Key Entities:** `CacheManager`, `EvictionStrategy` (LRU, LFU, FIFO), `Storage` (HashMap + Doubly Linked List for LRU).
*   **Challenges:** Ensuring O(1) time complexity for `get()` and `put()` operations, handling thread safety using `ConcurrentHashMap` or `ReentrantReadWriteLock`.

### 8. Design a Rate Limiter
*   **Focus:** Regulating API traffic limits.
*   **Key Entities:** `RateLimiterManager`, `RuleEngine`, `AlgorithmStrategy` (TokenBucket, SlidingWindow).
*   **Challenges:** Keeping track of timestamps efficiently, handling high concurrency without adding significant latency.

### 9. Design Tic-Tac-Toe / Chess Board Game
*   **Focus:** Game loops, rule validation, varying player types.
*   **Key Entities:** `Board`, `Cell`, `Player` (Human, AI), `Piece` (X, O or King, Queen, Pawn), `Move`, `GameStatus`.
*   **Challenges:** Checking win conditions efficiently (O(1) approach for Tic-Tac-Toe), validating complex moves (Chess).

### 10. Design an ATM System
*   **Focus:** Hardware interaction abstraction, bank API integration, transaction atomicity.
*   **Key Entities:** `ATM`, `CardReader`, `Keypad`, `CashDispenser`, `Transaction` (Withdrawal, Deposit, BalanceInquiry), `Account`, `BankService`.
*   **Challenges:** Handling exceptions (e.g., insufficient cash in machine, network failure during transaction deduction) ensuring no money is lost.

---

## 📝 3. Execution Strategy for LLD Interviews (90 mins)
1.  **Requirements Gathering (10 mins):** Restrict the scope. You cannot build the entire Amazon in 90 minutes. Agree on 3 core functionalities.
2.  **Class Diagrams / API Design (15 mins):** List out the primary classes, interfaces, and their relationships (Aggregation vs. Composition). Get buy-in from the interviewer.
3.  **Core Implementation (45 mins):** Start coding the core models and the service layer. Focus on SOLID principles. Use Mock repositories instead of actual DB connections.
4.  **Driver Code / Testing (15 mins):** Write a `Main` class to demonstrate the flow. E.g., `ParkingLot.getInstance().park(new Car())`.
5.  **Refactoring & Edge Cases (5 mins):** Discuss how you would handle thread safety or scaling the application.
