# Object-Oriented Design (LLD) - Service Based Companies

This section contains Low-Level Design (LLD) or Object-Oriented Design (OOD) interview questions commonly asked by Service-Based companies (e.g., TCS, Infosys, Wipro, Cognizant, IBM).

While Product-Based companies might focus heavily on vast scalability and complex concurrency, Service-Based companies usually focus on **clean Object-Oriented principles, properly mapping real-world business requirements into classes, and using straightforward Design Patterns.**

## Topics Covered:
1.  [Library Management System](01_lld_library_management.md)
2.  [ATM System](02_lld_atm_machine.md)
3.  [E-commerce Shopping Cart](03_lld_shopping_cart.md)
4.  [Traffic Light Control System](04_lld_traffic_light.md)
5.  [Vehicle Rental System](05_lld_vehicle_rental.md)
6.  [Hotel Management System](06_lld_hotel_management.md)
7.  [Banking System (Core Banking)](07_lld_banking_system.md)

## Success Criteria for LLD Interviews
*   **Requirements Gathering:** Can you define the scope of a broad idea like an "ATM"?
*   **Core Entities:** Can you map out classes like `Account`, `Transaction`, `BookItem` accurately?
*   **Relationships:** Do you understand Aggregation vs Composition? (e.g., A `Library` *has* `Books`).
*   **Design Patterns:** Knowledge of core patterns:
    *   **Factory:** Creating objects.
    *   **Observer:** Event notifications.
    *   **State:** Changing behavior based on entity status (like Traffic Lights or ATM states).
    *   **Singleton:** Central controller objects.
    *   **Strategy:** Executing different algorithms (like varying search types).
