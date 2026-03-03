# 📘 01 — Azure Fundamentals & Core Services
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy to Medium

---

## 🔑 Must-Know Topics
- Azure Geographies, Regions, and Availability Zones
- Resource Groups and Subscriptions
- Azure Resource Manager (ARM) vs Classic
- Cloud Models (IaaS, PaaS, SaaS)
- Management Tools (Portal, CLI, PowerShell)

---

## ❓ Most Asked Questions

### Q1. What is the difference between an Azure Region and an Availability Zone?

| Feature | Azure Region | Availability Zone (AZ) |
|---------|--------------|------------------------|
| **Definition** | A geographical area containing one or more datacenters. | Unique physical locations within a region. |
| **Purpose** | Data residency, compliance, and latency reduction. | High availability and fault tolerance within a region. |
| **Datacenters** | Multiple datacenters (usually close together). | Each AZ is made of 1+ datacenters with independent power/cooling. |
| **Latency** | Milliseconds within the region. | < 2ms latency between AZs. |
| **Availability** | Regional outages affect all AZs. | Zonal outages only affect the impacted zone. |

> **Note:** Not all Azure Regions support Availability Zones.

---

### Q2. What is a Resource Group in Azure?
A Resource Group (RG) is a logical container into which Azure resources (like web apps, databases, and storage accounts) are deployed and managed.

**Key Concepts:**
1. **Lifecycle Management:** If you delete a Resource Group, all resources inside it are deleted.
2. **Location:** A Resource Group has a location (e.g., East US), which specifies where the metadata for the resources is stored. The resources themselves can be in different regions.
3. **RBAC:** Role-Based Access Control can be applied at the RG level, granting users access to all resources within it.

---

### Q3. Explain Azure Resource Manager (ARM). What are its benefits?
ARM is the deployment and management service for Azure. It provides a management layer that enables you to create, update, and delete resources in your Azure account.

**Benefits of ARM:**
- **Deploy Multiple Resources Together (Templates):** You can deploy infrastructure as code using ARM templates (JSON/Bicep).
- **Declarative Syntax:** You define *what* you want, rather than writing scripts for *how* to deploy it.
- **RBAC:** ARM integrates natively with Azure RBAC.
- **Tagging:** Apply tags to resources for billing and logical organization.

---

### Q4. What is the difference between default/implicit limits and Azure Quotas?
- **Default Limits:** Pre-set limits on resources per subscription (e.g., max number of VMs). This is physical or designed.
- **Quotas (Soft Limits):** Adjustable limits designed to prevent runaway spending and ensure resources are available for all customers. You can request a quota increase through Azure Support.

---

### Q5. What is the difference between IaaS, PaaS, and SaaS in Azure?

| Service Type | Definition | Example in Azure | User Responsibility |
|--------------|------------|------------------|---------------------|
| **IaaS** | Infrastructure as a Service. You manage the OS, runtime, data, apps. | Azure Virtual Machines | High (OS patching, networking) |
| **PaaS** | Platform as a Service. You manage the application and data. | Azure App Services, Azure SQL | Medium (App logic, data) |
| **SaaS** | Software as a Service. You consume the software; cloud provider manages everything. | Microsoft 365, Dynamics 365 | Low (Access, user management) |

---

### Q6. Can a resource be moved from one Resource Group to another? From one Subscription to another?
**Yes, to both, but with limitations.**
Most resources can be moved between Resource Groups and Subscriptions. However, there are exceptions:
- **Downtime:** Moving usually doesn't cause downtime, but it locks the resource during the move.
- **Dependencies:** Resources with dependent resources must move together (e.g., a VM and its disks and network interfaces).
- **Service Support:** Some services do not support moving (e.g., certain App Service Plans across resource group boundaries depending on creation layout).

---

### Q7. What are Azure Tags and why use them?
Tags are name/value pairs assigned to resources or resource groups.

**Use cases:**
- **Billing:** Track costs by environment (e.g., `Env: Production`) or department (e.g., `Dept: HR`).
- **Management:** Organize and find resources quickly via the portal or scripts.
- **Automation:** Run scripts conditionally based on tags (e.g., shutdown VMs with `Env: Dev` at night).

> **Note:** Tags applied to a Resource Group are **not** automatically inherited by the resources within it, although Azure Policy can enforce inheritance.

---

### Q8. What is the difference between Azure CLI and Azure PowerShell?
Both are command-line tools to manage Azure resources.

| Feature | Azure CLI | Azure PowerShell |
|---------|-----------|------------------|
| **Base Language** | Python | PowerShell |
| **Best For** | Linux/macOS users, shell scripting | Windows users, object pipelining |
| **Syntax Style** | `az group create --name RG1` | `New-AzResourceGroup -Name RG1` |

> **Tip:** Neither is fundamentally better; it depends on the administrator's background and OS preference.
