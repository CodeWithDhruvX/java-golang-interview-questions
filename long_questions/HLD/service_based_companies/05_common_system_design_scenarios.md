# High-Level Design (HLD): Common System Design Scenarios

Service companies often ask standard, well-documented architecture scenarios to test end-to-end understanding.

## 1. Design a simple E-Commerce checkout workflow. Ensure no race conditions on inventory.
**Answer:**
This is a standard transactional HLD question for retail clients.
*   **Components:** Web Server, Order Service, Inventory Service, Payment Service, Database.
*   **The Race Condition Problem:** Two users try to buy the last iPhone at the exact same millisecond. 
*   **Solution (Pessimistic Locking / DB Row Lock):**
    1.  User clicks "Checkout".
    2.  Order Service starts an ACID Transaction.
    3.  Issues a `SELECT ... FOR UPDATE` query on the iPhone inventory row. This locks the row in the DB.
    4.  Second user's request is blocked by DB and has to wait.
    5.  First user's transaction checks if stock > 0, decrements it, creates the order, and `COMMIT`s.
    6.  Row lock is released. Second user's request continues, sees stock is 0, and receives an "Out of Stock" message.
*   *Alternative Solution (Optimistic Locking):* Use a `version` column. `UPDATE inventory SET stock = stock - 1, version = 2 WHERE product_id = 'iPhone' AND stock > 0 AND version = 1`. If it returns 0 rows updated, someone else bought it first.

## 2. Design an Enterprise Log Aggregation System
**Answer:**
When an enterprise runs 50 microservices, logging into 50 different Linux servers to read `system.log` when an error occurs is impossible.
*   **Architecture Pattern (ELK/EFK Stack):**
    *   **Generation (Filebeat / Fluentd):** A lightweight agent runs on every single application server. It tails the local log files and ships the log lines over the network.
    *   **Buffering (Queue/Kafka):** High throughput servers generate massive logs. Push logs to a Kafka queue to prevent overwhelming the indexer.
    *   **Processing & Indexing (Logstash):** Pulls logs from the queue, parses them (transforms raw text into structured JSON), and forwards them to storage.
    *   **Storage & Search Engine (Elasticsearch):** A distributed, RESTful search and analytics engine based on Lucene. Extremely fast at querying massive amounts of unstructured text.
    *   **Visualization (Kibana):** The UI dashboard where developers run queries and create charts to monitor the cluster.

## 3. How would you design an Employee Management System (Monolithic to REST API)?
**Core Requirements:** CRUD operations on Employee, Department, Salary. Standard Enterprise integration.
*   **Framework Options:** Java Spring Boot, Node/Express, or .NET Core.
*   **Database Schema:**
    *   `Employee`: id (PK), name, email, dept_id (FK).
    *   `Department`: id (PK), name.
*   **Architecture Layers (Spring Boot Example):**
    1.  `@RestController`: Handles the HTTP requests. (e.g., `GET /employees/{id}`). Parses JSON.
    2.  `@Service`: Contains complex business logic (e.g., calculates bonus, formats data). The Controller calls the Service.
    3.  `@Repository (Spring Data JPA)`: Interfaces with the Relational Database. Performs CRUD.
*   **Security:** Standard JWT authentication filter intercepting requests before reaching the Controller.
*   **Error Handling:** Use global exception handlers (`@ControllerAdvice`) to gracefully intercept application exceptions and translate them into standardized HTTP structured error JSON responses (e.g., 404 for EmployeeNotFound).

## 4. Design a Job Scheduling System (e.g., Quartz/Cron replacement across multiple nodes)
**Scenario:** A bank needs to run a nightly batch job across multiple application instances but needs to guarantee a specific job runs only ONCE.
*   **The Problem:** If you configure a cron expression on the application server, and run 3 clustered instances of the application, the nightly job will trigger 3 times concurrently, causing massive data duplication.
*   **Solutions:**
    1.  **Database Locking (Quartz Scheduler approach):**
        *   Maintain state in a centralized SQL database.
        *   Before executing the task, the cron thread on an instance tries to acquire a lock on the `QRTZ_LOCKS` table.
        *   Only one node successfully acquires the lock, executes the job, and releases the lock.
    2.  **Distributed Lock (Redis):**
        *   Use Redis `SETNX` (Set if Not eXists). Node attempts `SETNX("nightly_batch_lock", "timeout_value")`.
        *   If it returns 1 (success), the node executes. If 0, another node has the lock.
    3.  **Use a dedicated orchestration tool:** Airflow or AWS EventBridge/CloudWatch Events triggering an AWS Lambda function.

## 5. How do you plan a migration from an On-Premise Monolith to AWS Cloud?
**Answer:**
This is standard for service-based consulting interviews.
*   **Phase 1: Assessment.** Analyze current infrastructure, software licenses, dependencies, and network architecture.
*   **Phase 2: Lift and Shift (Rehosting).** The fastest way. Move the application "as-is" to cloud VMs (EC2). Move the on-prem DB to a cloud VM or managed service (RDS).
*   **Phase 3: Re-platforming.** Minor optimization without changing core code. E.g., moving from a self-managed WebSphere on EC2 to fully managed AWS Elastic Beanstalk.
*   **Phase 4: Refactoring/Rearchitecting.** Breaking the Monolith into Microservices over time. Replacing old relational schemas with NoSQL like DynamoDB where appropriate, implementing serverless architectures (Lambda) to cut costs.
