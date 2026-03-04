# 📘 07 — Event-Driven Microservices Architecture
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Advanced

---

## 🔑 Must-Know Topics
- Choreography vs Orchestration
- Dapr (Distributed Application Runtime)
- Azure Container Apps (ACA)
- Saga Pattern (Distributed Transactions)

---

## ❓ Most Asked Questions

### Q1. Microservices pattern: Orchestration vs Choreography?

| Feature | Orchestration | Choreography |
|---------|---------------|--------------|
| **Concept** | A central "controller" service commands other services (like a conductor). | Each service listens to events and acts independently (like dancers). |
| **Coupling**| Higher coupling. The orchestrator must know about all services. | Decoupled. Services only know about the Event Bus. |
| **Azure Tools**| **Azure Durable Functions**, Logic Apps. | **Azure Event Grid**, **Service Bus Topics**. |
| **Pros/Cons**| Easier to track the workflow/transaction status. | Highly scalable, but harder to monitor "Where is this business transaction stuck?" |

---

### Q2. How do you handle distributed transactions across microservices? (The Saga Pattern)
In a microservices architecture, you cannot use ACID SQL transactions across two different databases (e.g., Order Database and Inventory Database).

Use the **Saga Pattern**:
1. A transaction is broken into consecutive local transactions.
2. After the Order service updates its DB, it publishes an event: `OrderCreated`.
3. The Inventory service listens, updates its DB, and publishes `InventoryReserved`.
4. **Compensation:** If the Payment service subsequently fails, it publishes a `PaymentFailed` event. The previous services must listen for this and execute compensating logic (e.g., Inventory service adds the stock back).

---

### Q3. What is Azure Container Apps (ACA)? How does it relate to AKS?
ACA is a fully managed serverless container service built *top of* AKS, but it hides the Kubernetes API and complexity from the developer.

- **Best for:** Event-driven microservices.
- **Built-in open source:** Comes natively integrated with **KEDA** (for autoscaling to zero based on events) and **Dapr** (Distributed Application Runtime).
- **Difference from AKS:** You don't have access to the underlying Kubernetes API, node pools, or advanced networking configurations. It is PaaS for containers.

---

### Q4. What is Dapr (Distributed Application Runtime)?
Dapr is an open-source framework used extensively in Azure Container Apps and AKS to make microservice development easier. It uses a **Sidecar Pattern**.

**Why use it?**
Instead of writing SDK code to connect to Azure Service Bus, Redis, or Azure Key Vault, your application makes simple HTTP/gRPC calls to the local Dapr sidecar (`localhost:3500`).
The sidecar handles the complex connections to the Azure services. If you decide to switch from Azure Service Bus to RabbitMQ, you simply change a Dapr YAML config file. **Zero code changes are required in your application.**
- Handles state management.
- Handles pub/sub messaging.
- Handles service-to-service invocation (with mTLS automatically).

---

### Q5. Explain the Sidecar Pattern in Kubernetes.
A sidecar is a secondary container that runs alongside the main application container within the same Kubernetes Pod. They share the same network namespace and disk volumes.

**Use cases:**
- **Logging:** A Filebeat sidecar reading logs from the app container's volume and forwarding them to elasticsearch.
- **Proxy/Service Mesh:** Envoy proxy routing traffic to the app, handling SSL termination, and emitting metrics.
- **Dapr:** Handling distributed architecture concerns as described above.
