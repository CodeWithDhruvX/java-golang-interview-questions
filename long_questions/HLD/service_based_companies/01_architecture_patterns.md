# High-Level Design (HLD): Architecture Patterns (Service-Based Companies)

In service-based companies (TCS, Infosys, Cognizant, Wipro, Accenture), HLD questions often focus on standard architectural patterns, system integration, and enterprise software design.

## 1. What is the N-Tier Architecture? Explain the typical 3-Tier setup.
**Answer:**
N-Tier (or Multi-Tier) architecture separates an application into multiple layers logically and physically. This provides separation of concerns, scalability, and security.
*   **Presentation Tier (Web/UI):** The top-most level. Can be a React/Angular SPA, or traditional MVC views (e.g., Spring MVC, JSP). Displays data and interacts with users.
*   **Logic / Application Tier (Business Logic):** The middle tier. Contains business logic, calculations, data processing. Usually built with REST APIs (Spring Boot, Node.js).
*   **Data Tier (Database):** Stores and manages the data. Could be an RDBMS (Oracle, SQL Server) or NoSQL database.
*   *Why use 3-Tier?* Each tier can be deployed on separate servers, scaled independently, and updated without affecting the others.

## 2. Monolithic Architecture vs. Service-Oriented Architecture (SOA)
**Answer:**
*   **Monolithic Architecture:** All components (UI, Business Logic, DB Access) are tightly coupled in a single codebase and deployed as one massive unit (e.g., a huge WAR file in Java).
    *   *Pros:* Easy to test, simple deployment (initially), simpler debugging.
    *   *Cons:* Very hard to scale specific features, technology lock-in, slow startup time, difficult for multiple teams to work concurrently.
*   **SOA (Service-Oriented Architecture):** An enterprise pattern where application components provide services to other components via a communications protocol (often Enterprise Service Bus - ESB, or SOAP/WSDL). They are loosely coupled and highly reusable across the enterprise. 
    *   *Note on Microservices:* Microservices are a modern, specific implementation of SOA principles using lightweight protocols (REST) and decentralized data.

## 3. What is MVC (Model-View-Controller) pattern?
**Answer:**
MVC is primarily an architectural pattern for user interfaces that separates application logic into three interconnected elements.
*   **Model:** The central component. It manages the data, logic, and rules of the application. Notifies the View of state changes.
*   **View:** The visual representation of the Model. The UI component users interact with. (e.g., HTML pages).
*   **Controller:** Accepts input from the user (via the View), translates it into commands for the Model or View.
*   *Enterprise Example:* In Spring Web MVC, the `DispatcherServlet` acts as a front controller, routing requests to specific `@Controller` classes.

## 4. Difference between Broker and Hub-and-Spoke integration patterns.
**Answer:**
When integrating many internal systems within an enterprise:
*   **Hub-and-Spoke (Centralized):** All communication flows through a central Hub (like an Enterprise Service Bus - ESB). The Hub handles routing, transformation, and delivery.
    *   *Pros:* Centralized management, easy adding of new systems.
    *   *Cons:* Single Point of Failure (SPOF), the Hub can become a bottleneck.
*   **Broker Pattern (Decentralized/Event-Driven):** Systems communicate via an event broker/message queue (e.g., ActiveMQ, Kafka). The broker simply routes messages based on topics/queues. Systems subscribe to messages they care about.
    *   *Pros:* Highly scalable, decentralized decoupling of producers and consumers.

## 5. What are the pros and cons of using an ORM (Object-Relational Mapping)?
**Answer:**
ORMs (like Hibernate/JPA in Java, Entity Framework in C#, Sequelize in Node) bridge the gap between Object-Oriented application code and Relational Databases.
*   **Pros:**
    *   Speeds up development (no need to write boiler-plate SQL/JDBC/ADO.NET code).
    *   Database Agnostic (switching from MySQL to PostgreSQL is often just a configuration change).
    *   Provides out-of-the-box features like caching and transaction management.
*   **Cons:**
    *   **N+1 Query Problem:** Automatically lazy-loading collections can result in hundreds of hidden SQL queries, killing performance.
    *   Performance Overhead: Generated SQL is often not as optimal as hand-written SQL for complex joins or reporting queries.
    *   Steep learning curve for advanced features.
