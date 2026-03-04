# 📘 07 — Azure Messaging & Integration
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Logic Apps vs Azure Functions
- Service Bus (Queues vs Topics)
- Event Grid vs Event Hubs
- API Management (APIM)

---

## ❓ Most Asked Questions

### Q1. Compare Azure Logic Apps and Azure Functions.

| Feature | Azure Functions | Azure Logic Apps |
|---------|-----------------|------------------|
| **Core Concept** | Code-first (Write code in C#, Java, Python, Node.js). | Designer-first (Visual drag-and-drop workflow). |
| **Best For** | Complex custom logic, specialized algorithms, data processing. | Orchestrating integrations between SaaS apps (Office 365, Salesforce, SAP) with out-of-the-box connectors. |
| **State** | Stateless (unless using Durable Functions). | Stateful (each step in the workflow retains state visually). |
| **Pricing** | Consumption-based (per execution time/memory). | Consumption-based (per action/connector execution). |

---

### Q2. What is Azure Service Bus? Difference between Queues and Topics?
Service Bus is an enterprise message broker offering highly reliable cloud messaging between applications and services.

- **Queues (1:1 - Point-to-Point):** A sender sends a message to the queue, and a **single** receiver picks it up and processes it. Good for load-leveling and decoupling.
- **Topics (1:N - Publish/Subscribe):** A sender sends a message to a topic. **Multiple** subscribers can listen to that topic using Subscriptions. Each subscription receives a copy of the message (often filtered based on rules).

---

### Q3. Compare Azure Event Grid, Event Hubs, and Service Bus.

| Service | Type | Focus | Use Case |
|---------|------|-------|----------|
| **Event Grid** | Events | Reactive programming (Pub/Sub) | **"Something happened."** (e.g., A file was uploaded to Blob Storage, trigger a Function). High scale, lightweight routing. |
| **Event Hubs** | Events | Big Data streaming / Telemetry | **"Here is a continuous stream of data."** (e.g., Millions of IoT sensor readings per second). Needs partition keys and stream processing. |
| **Service Bus** | Messages | Enterprise Messaging | **"Do this specific task."** (e.g., Process payment order #1234). Needs high reliability, transaction support, decoupling, and FIFO ordering. |

> **Key Distinction:** A **Message** contains the raw data required to execute an action (e.g., the JSON of an order). An **Event** is merely a notification that a state changed (e.g., "Order 1234 was created").

---

### Q4. What is Azure API Management (APIM)?
APIM is a hybrid, multicloud management platform for APIs across all environments. It acts as a gateway/facade between backend services and frontend clients.

**Key Features:**
- **Security:** Protect backends via rate limiting, quotas, IP filtering, and validating OAuth/JWT tokens.
- **Transformation:** Convert XML to JSON on the fly, remove headers, or change URLs without modifying the backend.
- **Analytics:** Centralized logging of all API calls, latency tracking, and error monitoring.
- **Developer Portal:** Auto-generated documentation for third-party or internal developers to discover and test APIs.

---

### Q5. What is the Dead-Letter Queue (DLQ) in Azure Service Bus?
The DLQ is a secondary sub-queue created automatically by Service Bus. It holds messages that cannot be successfully delivered or processed by the receiver.

**Reasons messages go to the DLQ:**
- `MaxDeliveryCount` exceeded (the receiver tried to process it but threw an exception multiple times).
- Time To Live (TTL) expired before the message was processed.
- The receiver explicitly "dead-letters" the message (e.g., validation failed).
> Developers monitor the DLQ to investigate poison messages and either fix the bug and replay them, or discard them.
