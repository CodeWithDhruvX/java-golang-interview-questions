# 📘 02 — Azure Compute Services
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Virtual Machines (VMs) and VM Scale Sets (VMSS)
- Azure App Services (Web Apps)
- Azure Functions (Serverless concepts)
- Availability Sets vs Availability Zones for VMs
- Dedicated vs Spot instances

---

## ❓ Most Asked Questions

### Q1. What is an Azure Virtual Machine Scale Set (VMSS)?
VMSS lets you create and manage a group of load-balanced VMs. The number of VM instances can automatically increase or decrease in response to demand or a defined schedule.

**Benefits:**
- **High Availability:** Automatically distributes VMs across Update Domains (UD) and Fault Domains (FD), or Availability Zones.
- **Autoscaling:** Scale-out (add instances) during peak load; scale-in (remove instances) when load drops, saving money.
- **Central Management:** You manage the set, rather than individual VMs.

---

### Q2. What is an Availability Set? Fault Domains vs Update Domains?
An Availability Set is a logical grouping capability for isolating VM resources from each other when they're deployed, providing high availability in the *same* datacenter.

- **Fault Domain (FD):** A group of VMs that share a common power source and network switch. If the rack loses power, only VMs in that FD go down.
- **Update Domain (UD):** A group of VMs and underlying hardware that can be rebooted at the same time during planned Azure maintenance.

> **Note:** For 99.99% SLA, use Availability Zones (spread across datacenters). For 99.95% SLA, use Availability Sets.

---

### Q3. What is Azure App Service and its plans?
Azure App Service is an HTTP-based PaaS for hosting web applications, REST APIs, and mobile back ends. Developers can build in .NET, Java, Ruby, Node.js, PHP, or Python.

**App Service Plans (Pricing Tiers determine CPU/Memory/Features):**
- **Free/Shared:** Dev/Test only. Apps run on shared VMs. No custom domains/SSL.
- **Basic:** Dedicated VMs. Supports custom domains, standard SLA.
- **Standard:** Auto-scaling, daily backups, staging slots.
- **Premium:** Faster processors, SSD storage, increased scaling limits.
- **Isolated:** Runs in an App Service Environment (ASE) for VNet isolation and maximum security.

---

### Q4. What are App Service Deployment Slots?
Deployment Slots are live apps with their own hostnames. You can swap the content and configuration between two deployment slots (e.g., `Staging` and `Production`).

**Why use them?**
- **Zero-Downtime Deployments:** Warm up the staging slot, then swap VIPs (Virtual IPs) instantly.
- **A/B Testing:** Route a percentage of traffic to the new slot.
- **Easy Rollback:** If the new production code has issues, swap back immediately to the previous slot.

---

### Q5. What are Azure Functions? (Serverless Compute)
Azure Functions is an event-driven, serverless compute service. You write the code (the "function") and Azure manages the infrastructure to run it.

**Key Concepts:**
- **Event-Driven:** Triggered by HTTP requests, timers (cron jobs), blob uploads, message queues, etc.
- **Micro-billing:** Pay only for the time the code runs and memory used (Consumption Plan).
- **Stateless by default:** They don't retain memory between executions (Durable Functions can be used for stateful workflows).

---

### Q6. Spot VMs vs Reserved VM Instances?

| Feature | Spot VMs | Reserved Instances (RI) | Pay-As-You-Go |
|---------|----------|-------------------------|---------------|
| **Pricing** | Deepest discount (up to 90%) | Up to 72% discount | Most expensive |
| **Commitment**| None | 1 or 3 years upfront commitment | None |
| **Eviction** | Can be evicted anytime if Azure needs capacity | Never evicted, reserved capacity | Never evicted |
| **Best For** | Batch jobs, stateless dev/test workloads | Steady, predictable production workloads | Spiky, unpredictable short-term workloads |

---

### Q7. How do you troubleshoot an Azure Web App that is running slow or returning 500 errors?
1. **Azure App Service Diagnostics:** Go to "Diagnose and solve problems" in the portal. It provides out-of-the-box analysis.
2. **Log Stream:** Enable application logging to view real-time log output from the app.
3. **Application Insights:** Check for slow dependencies (database queries, external API calls), exception rates, and performance counters.
4. **Scale Up/Out:** Check Metrics (CPU/Memory). If exhausted, scale up (bigger instance) or scale out (more instances).
5. **Kudu Console (Advanced Tools):** Inspect the file system, process explorer, and run memory dumps.

---

### Q8. What is Azure Container Instances (ACI) vs Azure Kubernetes Service (AKS)?
- **ACI (Azure Container Instances):** PaaS offering to run a single container or a small group of containers without managing any VMs or orchestrators. Easiest/fastest way to run a container. Best for simple apps, task automation.
- **AKS (Azure Kubernetes Service):** Managed Kubernetes offering. Provides orchestration, scaling, automated rollouts/rollbacks, and advanced networking. Best for full-scale microservices architectures.
