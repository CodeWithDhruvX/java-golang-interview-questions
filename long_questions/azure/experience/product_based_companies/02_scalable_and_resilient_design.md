# 📘 02 — Scalable & Resilient Design
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Advanced

---

## 🔑 Must-Know Topics
- High Availability (HA) vs Disaster Recovery (DR)
- RTO (Recovery Time Objective) and RPO (Recovery Point Objective)
- Multi-Region Active-Active vs Active-Passive deployments
- Azure Front Door and Traffic Manager
- Asynchronous patterns (CQRS, Event Sourcing)

---

## ❓ Most Asked Questions

### Q1. Differentiate High Availability (HA) and Disaster Recovery (DR).
- **High Availability (HA):** Designing a system to handle component failures *without* downtime within a single region. (e.g., Load balancing across 3 VMs in different Availability Zones). Focuses on fault tolerance.
- **Disaster Recovery (DR):** The process of recovering systems and data after a catastrophic failure (e.g., an entire Azure Region goes down). Focuses on cross-region failover and backups.

---

### Q2. Explain RTO and RPO. How do they affect architectural choices?
- **RTO (Recovery Time Objective):** The maximum acceptable amount of time the application can be offline. (How fast must we recover? e.g., 2 hours).
- **RPO (Recovery Point Objective):** The maximum acceptable amount of data loss, measured in time. (How much data can we lose? e.g., 15 minutes of transactions).

**Architectural impact:**
- Near-zero RPO/RTO requires Multi-Region Active-Active deployments (expensive).
- High RPO/RTO can survive with weekly backups and manual redeployments (cheap).

---

### Q3. How would you design a Multi-Region Active-Active Web Application in Azure?
An Active-Active design means the application is running and answering requests in multiple regions simultaneously.

**Architecture:**
1. **Global Routing:** Use **Azure Front Door** or **Traffic Manager** to route users to the closest/fastest region.
2. **Compute:** Deploy **Azure App Services** or **AKS** in both Region A (e.g., East US) and Region B (e.g., West Europe).
3. **Data:** Use **Azure Cosmos DB** with Multi-Region Writes enabled (data is synchronized globally). If using SQL, you'll have to use Read Replicas and route writes to a primary region (making it Active-Passive for the DB).
4. **State:** Use **Azure Cache for Redis** (Geo-replicated) to ensure session state isn't lost if a user is routed to a different region.

---

### Q4. Compare Azure Front Door vs Azure Traffic Manager for Global Routing.

| Feature | Azure Traffic Manager | Azure Front Door |
|---------|-----------------------|------------------|
| **Routing Level** | DNS-based (Global) | Layer 7 (HTTP/HTTPS) Anycast |
| **Traffic Type** | Any protocol (TCP, UDP, HTTP) | Web/HTTP(S) only |
| **Caching/WAF** | No | Yes (Integrated CDN and Web App Firewall) |
| **Failover Speed**| Relies on DNS TTL (can be slow if clients cache DNS) | Extremely fast (proxies traffic over Microsoft backbone) |

> **Use Case:** Use Front Door for web applications and APIs. Use Traffic Manager for non-HTTP traffic or hybrid cloud scenarios.

---

### Q5. What is the CQRS (Command and Query Responsibility Segregation) pattern? Why use it?
CQRS separates the data mutation operations (Commands) from the data read operations (Queries).

**Implementation in Azure:**
- **Commands (Writes):** Sent to an Azure Service Bus, processed by an Azure Function, and written to Azure SQL Database.
- **Queries (Reads):** Read rapidly from an Azure Cache for Redis or Cosmos DB materialized view.

**Why use it?**
In highly scalable systems, the read load is usually vastly different from the write load. CQRS allows you to scale the read infrastructure independently from the write infrastructure.

---

### Q6. What is the Circuit Breaker pattern?
When a microservice relies on a remote service (like a third-party API or a database under heavy load), continuous retries can cause escalating failures.
A Circuit Breaker acts as a proxy:
- **Closed State:** Requests flow normally.
- **Open State:** If the failure rate exceeds a threshold, the circuit "opens" and immediately returns an error. It stops hammering the failing system.
- **Half-Open State:** Periodically allows a single test request through. If it succeeds, the circuit closes.

> **Azure implementation:** Often implemented in code using libraries like Polly (in .NET) or via a Service Mesh (like Istio in AKS).
