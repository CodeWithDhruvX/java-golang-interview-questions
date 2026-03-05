# Migration and Legacy Modernization (Product-Based Companies)

## 1. How would you migrate a massive legacy monolith to microservices without incurring downtime? Describe the Strangler Fig Pattern.

**Expected Answer:**
The "Big Bang" rewrite is almost always a failure. Migrations must be incremental. The industry standard is the **Strangler Fig Pattern**.

1.  **API Gateway/Routing Layer:** Introduce a smart proxy or API Gateway in front of the legacy monolith. Initially, 100% of traffic routes to the monolith.
2.  **Identify Boundary/Domain:** Choose a single, decoupled feature to extract first (e.g., "User Profile" or "Email Notifications"). Do not start with the core transaction engine.
3.  **Develop the Microservice:** Build the new service alongside the monolith.
4.  **Traffic Routing (The Strangulation):** Update the API Gateway to route specific endpoints (e.g., `/api/v2/users`) to the new microservice, while all other endpoints still go to the monolith.
5.  **Decommission:** Once the microservice is stable and data is migrated, delete the corresponding dead code in the monolith.
6.  **Repeat:** Continuously repeat this process until the monolith is entirely "strangled" and can be shut down.

## 2. When modernizing a system, how do you prevent the new system from adopting the poor data models or technical debt of the legacy system?

**Expected Answer:**
Use an **Anti-Corruption Layer (ACL)** (a concept from Domain-Driven Design).

*   **The Problem:** The new microservice needs to communicate with the legacy system, but the legacy system uses outdated paradigms (e.g., massive shared database tables, SOAP APIs, XML RPC). If the new service uses these same models, it becomes "corrupted" by the legacy design.
*   **The Solution:** Build a translation layer (the ACL) between the new service and the legacy system.
    *   The new service speaks its own modern, clean domain language (e.g., clean JSON, gRPC).
    *   The ACL intercepts requests, translates the modern domain model into the legacy format, makes the call to the legacy system, and translates the response back.
    *   *Result:* The new service remains pure. When the legacy system is eventually retired, you only have to discard the ACL; the new service's core logic remains untouched.

## 3. Database migrations are often the hardest part of modernizing. Explain how to migrate a zero-downtime database from a legacy RDBMS to a new NoSQL store.

**Expected Answer:**
Database migrations require careful synchronization to avoid data loss and downtime. The standard approach is the **Dual-Write / Multi-Phase Migration**.

*   **Phase 1: Dual Writes (Write to Both, Read from Old)**
    *   Modify the application to write to *both* the old RDBMS and the new NoSQL database synchronously.
    *   All reads still come from the old RDBMS (the source of truth).
*   **Phase 2: Backfill (Historical Data)**
    *   Run a background script to migrate all historical data from the old RDBMS to the NoSQL store.
    *   *Crucial:* Ensure the backfill script handles conflicts gracefully (if a record was updated via dual-write during the backfill, don't overwrite the newer NoSQL record with old RDBMS data).
*   **Phase 3: Verification / Shadow Reads**
    *   Application reads from *both* databases, but only returns the old RDBMS result to the user. Log any discrepancies between the two results to verify the NoSQL data integrity and migration logic.
*   **Phase 4: Cutover (Read from New)**
    *   Swap the primary read source to the NoSQL database. Leave dual writes on in case a rollback is needed.
*   **Phase 5: Cleanup**
    *   Once fully confident, stop dual writing. Decommission the old RDBMS.

## 4. What is Change Data Capture (CDC), and how does it help in system migration and event-driven architectures?

**Expected Answer:**
**Change Data Capture (CDC)** is a pattern to detect and capture changes made to a database and forward those changes to downstream systems in real-time.

*   **How it works:** Instead of polling the database with `SELECT * WHERE updated_at > ?` (which hurts DB performance and has high latency), CDC tools (like **Debezium**) read the database's internal transaction logs (e.g., MySQL `binlog`, PostgreSQL `WAL`).
*   **Role in Migration:** CDC silently streams data changes from a legacy database to a new database or message broker (Kafka) without modifying the legacy application code. This is a powerful alternative to Dual Writes (which require code changes in the legacy app).
*   **Role in Event-Driven Architecture:** CDC allows you to turn an antique legacy database into an event producer. When legacy App A writes a row to the DB, Debezium reads the WAL, publishes a "RowUpdated" event to Kafka, and modern Microservice B consumes that event to trigger its own logic.

## 5. How do you handle dark launching or canary releases during a major system migration?

**Expected Answer:**
Risk mitigation is paramount during migrations. It's unsafe to flip a switch for 100% of users.

*   **Feature Flags / Toggles:** Wrap the new migration logic in feature flags. Toggle the new behavior "on" dynamically without redeploying code.
*   **Dark Launching (Shadow Traffic):** Duplicate real user traffic at the network level (e.g., using Envoy or NGINX). Route the primary traffic to the legacy system (which replies to the user). Send the duplicated shadow traffic to the new system, ignore its response, but monitor its logs, error rates, and performance.
*   **Canary Deployments:**
    *   Route a small percentage (e.g., 1%) of real user traffic to the new modernized system.
    *   Actively monitor error rates, latency, and business metrics (e.g., checkout success rate).
    *   Gradually ramp up the percentage (5%, 20%, 50%, 100%) if metrics are healthy. Automatically rollback to 0% if error thresholds are breached.
