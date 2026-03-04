# 📘 01 — Advanced Cloud Architecture
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Advanced

---

## 🔑 Must-Know Topics
- Azure Well-Architected Framework
- Cloud Adoption Framework (CAF)
- CAPEX vs OPEX and TCO (Total Cost of Ownership)
- Multi-tenant architecture design patterns
- Infrastructure as Code (IaC) maturity

---

## ❓ Most Asked Questions

### Q1. What are the 5 pillars of the Azure Well-Architected Framework?
The Well-Architected Framework provides guiding principles to improve the quality of a workload.
1. **Reliability:** Ability of a system to recover from failures and continue to function.
2. **Security:** Protecting applications and data from threats.
3. **Cost Optimization:** Managing costs to maximize the value delivered.
4. **Operational Excellence:** Operations processes that keep a system running in production (DevOps, monitoring).
5. **Performance Efficiency:** The ability of a system to adapt to changes in load (Autoscaling, caching).

---

### Q2. How do you design a Multi-Tenant application on Azure?
Designing an architecture that serves multiple customers (tenants) securely and efficiently is core to SaaS products.
There are three main models:
1. **Database-per-tenant (Silo model):** Each tenant has their own database or even their own isolated compute. High isolation, easy billing, but very expensive and hard to manage at scale.
2. **Shared compute, isolated data (Bridge model):** Tenants share the web/API servers, but each has their own database (e.g., Azure SQL Elastic Pools). Balances cost and data isolation.
3. **Shared everything (Pool model):** Tenants share compute and the database (differentiated by `TenantId`). Very cost-effective, but high risk of "Noisy Neighbor" problems and complex data partitioning.

> **Azure Services for Multi-Tenancy:** Azure AD B2C (Identity), Azure SQL Elastic Pools, Event Hubs (with partitioning).

---

### Q3. Explain the "Noisy Neighbor" problem in the cloud. How do you mitigate it?
The "noisy neighbor" problem occurs when a single tenant uses a large amount of shared resources (CPU, Memory, IOPS), causing performance degradation for other tenants.

**Mitigation strategies:**
1. **Throttling/Rate Limiting:** Implement API Management (APIM) to throttle requests per tenant.
2. **Queueing (Load Leveling):** Put incoming heavy requests into an Azure Service Bus queue and process them asynchronously.
3. **Sharding/Data Partitioning:** Move the noisy neighbor to their own dedicated infrastructure (Database-per-tenant).
4. **Cosmos DB RUs:** Provision Requests Units (RUs) per container rather than at the database level to ensure one container doesn't consume all RUs.

---

### Q4. What is the Cloud Adoption Framework (CAF)?
CAF is a collection of documentation, implementation guidance, and best practices that help organizations align business and technical strategies to succeed in the cloud.
Phases include:
- **Strategy:** Define business justification and expected outcomes.
- **Plan:** Align actionable adoption plans to business outcomes.
- **Ready:** Prepare the cloud environment (Landing Zones).
- **Adopt:** Migrate existing apps or innovate new ones.
- **Govern & Manage:** Implement Azure Policy, Cost Management, and monitoring.

---

### Q5. What is an Azure Landing Zone?
An Azure Landing Zone is the output of a multi-subscription Azure environment that accounts for scale, security governance, networking, and identity. It is the architectural *foundation* of the "Ready" phase in CAF.
It ensures that when a new product team needs a subscription, they get an environment that already has RBAC, networking peering to the hub, and security policies pre-configured.
