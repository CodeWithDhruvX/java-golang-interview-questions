# High-Level Design (HLD): Enterprise Integration and Legacy Modernization

Service-based companies handle massive digital transformation projects, moving 30-year-old on-premise systems to the cloud.

## 1. What are Enterprise Integration Patterns (EIP)? Give examples.
**Answer:**
EIP is a collection of technology-agnostic design patterns for integrating enterprise applications.
*   **Message Router:** A component that consumes a message from a channel and, based on a set of conditions, routes it to a specific output channel (e.g., routing European orders to a European DB, and US orders to a US DB).
*   **Message Translator:** If System A outputs XML and System B expects JSON, a translator sits between them, converting the payload.
*   **Aggregator:** Receives multiple related messages (e.g., flight price, hotel price, rental car price from 3 different APIs) and combines them into a single comprehensive message to present to the user interface.
*   **Scatter-Gather:** A message is broadcast (scattered) to multiple recipients, and the responses are aggregated (gathered) back into a single response. Useful for getting quotes from 5 different vendors simultaneously.

## 2. Compare API Gateway vs. Enterprise Service Bus (ESB) vs. Service Mesh.
**Answer:**
*   **ESB (Enterprise Service Bus):** A centralized architectural model where all communication between applications goes through a massive, heavy "hub." The hub handles complex translations, orchestration, and routing. 
    *   *Drawback:* Becomes a monolithic bottleneck and single point of failure. Modern architectures are moving away from it.
*   **API Gateway:** Sits at the edge of the network. It's the entry point for *external* clients (mobile phones, SPA web apps) to access internal microservices. Handles auth, rate limiting, and basic routing.
*   **Service Mesh (e.g., Istio, Linkerd):** Sits *inside* the network. It manages service-to-service (internal) communication. It handles complex traffic routing (canary deployments), internal mTLS encryption, retry logic, and observability *without* changing the application code (usually deployed as a "Sidecar" proxy next to each container).

## 3. What is the Strangler Fig Pattern?
**Answer:**
The industry-standard approach for safely migrating from a huge legacy Monolith to Microservices.
*   **The Problem:** Rewriting a 10-year-old application entirely from scratch (a "Big Bang" rewrite) usually fails spectacularly.
*   **The Strangler Fig Solution:**
    1.  Put an API Gateway or Reverse Proxy (like NGINX) in front of the legacy monolith. Route 100% of traffic to the monolith.
    2.  Identify a small, distinct piece of functionality (e.g., the User Profile module).
    3.  Build a modern microservice just for that functionality in the cloud.
    4.  Update the API Gateway to route `/profile` traffic to the new microservice, while leaving all other traffic routed to the legacy monolith.
    5.  Slowly, over months or years, repeat this process. The new microservices slowly "strangle" the monolith until the legacy system has no traffic and can be shut off.

## 4. What is the Anti-Corruption Layer (ACL) pattern?
**Answer:**
Also used heavily during legacy modernization.
*   **The Problem:** You have a clean, modern microservice built around clear Domain-Driven Design (DDD). It needs to coordinate with an older, messy legacy system with terrible, confusing database tables. If you let your new service understand the legacy schema, your new code becomes polluted and messy.
*   **The Solution (ACL):** You build an intermediary translation layer (a façade/adapter) between the modern service and the legacy system.
    *   The modern system only speaks to the ACL using its clean, modern data models.
    *   The ACL translates those models into the messy, confusing calls required by the legacy system.
    *   *Benefit:* When the legacy system is finally fully retired, you just delete the ACL. Your modern service's core code remains untouched and clean.

## 5. What are the 6 R's of Cloud Migration Strategy?
**Answer:**
When consulting on moving a client to AWS/Azure, architects classify applications into 6 buckets:
1.  **Rehosting ("Lift and Shift"):** Moving the application exactly as-is from on-premise servers to cloud Virtual Machines (EC2). Fast, but misses out on native cloud benefits.
2.  **Replatforming ("Lift, Tinker, and Shift"):** Making minor optimizations without changing the core architecture. E.g., moving from a self-managed Oracle DB to AWS RDS managed database.
3.  **Repurchasing:** Throwing away the custom-built legacy app and buying a modern SaaS product (moving from a custom HR system to Workday).
4.  **Refactoring / Rearchitecting:** Completely rewriting the application to be cloud-native (breaking a monolith into Lambda functions and DynamoDB). Most expensive, highest ROI.
5.  **Retain:** Keeping the application on-premise because it's too complex or due to strict regulatory compliance.
6.  **Retire:** Discovering the application is no longer used by anyone and permanently shutting it down to save costs.
