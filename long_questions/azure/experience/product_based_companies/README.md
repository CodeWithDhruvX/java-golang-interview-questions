# Azure Interview Questions for Product-Based Companies

This repository contains an advanced collection of Azure-related interview questions and architectural design scenarios. These questions are specifically curated for product-based technology companies that build scalable, resilient, and multi-tenant SaaS products natively on the cloud.

The focus here is heavily on system design, microservices, containerization, and advanced scalability patterns, rather than simple administration.

## 📂 Content Index

1. [Advanced Cloud Architecture](01_advanced_cloud_architecture.md)
   - Covered concepts: Well-Architected Framework, Cloud Adoption Framework (CAF), Multi-Tenant Architectures (Silo vs Pool), Noisy Neighbor problem.
2. [Scalable & Resilient Design](02_scalable_and_resilient_design.md)
   - Covered concepts: HA vs DR, RTO/RPO, Multi-Region Active-Active deployments, CQRS, Circuit Breaker, Azure Front Door vs Traffic Manager.
3. [System Design on Azure](03_system_design_on_azure.md)
   - Covered concepts: Streaming service architectures, handling spiky traffic (Load Leveling), distributed caching, Lambda/Kappa architectures for IoT.
4. [Azure Kubernetes Service (AKS) Advanced](04_azure_kubernetes_service_aks_advanced.md)
   - Covered concepts: Kubenet vs Azure CNI, Pod Autoscaling (HPA, KEDA), Network Policies, CSI Drivers for Persistent Storage, App Gateway Ingress Controller (AGIC).
5. [Performance Optimization & Cost Management](05_performance_optimization_and_cost_management.md)
   - Covered concepts: FinOps, Cosmos DB Hot Partitions and RU tuning, HTTP 429 backoff strategies, App Service profiling, edge caching.
6. [Advanced Data & Analytics](06_advanced_data_and_analytics.md)
   - Covered concepts: Databricks vs Synapse Analytics, ADLS Gen2 Hierarchical Namespaces, ETL vs ELT, ADF error handling, Microsoft Purview.
7. [Event-Driven Microservices Architecture](07_event_driven_microservices_architecture.md)
   - Covered concepts: Orchestration vs Choreography, Saga Pattern for distributed transactions, Azure Container Apps (ACA), Dapr, Sidecar pattern.
8. [Advanced Security & Compliance](08_advanced_security_and_compliance.md)
   - Covered concepts: Zero Trust Architecture, Defender for Cloud vs Microsoft Sentinel, VNet Injection vs Private Endpoints, Customer-Managed Keys (BYOK).

## 🎯 Target Audience
These materials are tailored to senior technical roles such as:
- Senior Cloud Software Engineer
- Cloud Solutions Architect / Enterprise Architect
- Site Reliability Engineer (SRE)
- Lead Data Engineer

## 💡 How to use this guide
- **Think in Trade-offs:** Product companies rarely look for a single "right" answer. They want to hear you discuss the trade-offs (e.g., consistency vs availability, cost vs performance).
- **Focus on Scale:** Always consider how a solution behaves when traffic spikes 100x.
- **Draw it out:** When practicing these questions, try whiteboarding the architecture simultaneously.
