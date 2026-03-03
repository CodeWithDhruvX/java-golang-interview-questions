# 📘 06 — Azure Monitoring & Governance
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Azure Monitor vs Application Insights vs Log Analytics
- Azure Policy
- Azure Blueprints vs ARM Templates
- Cost Management and Budgets
- Azure Advisor

---

## ❓ Most Asked Questions

### Q1. What is the difference between Azure Monitor, Log Analytics, and Application Insights?
They are all part of the overarching **Azure Monitor** suite, but serve distinct purposes:

- **Azure Monitor:** The core hub for collecting, analyzing, and acting on telemetry from your cloud and on-premises environments. It collects **Metrics** (numerical values at a point in time) and **Logs** (timestamped records of events).
- **Application Insights:** An Application Performance Management (APM) feature of Azure Monitor. It is specifically aimed at developers to monitor live web applications (detects anomalies, tracks request rates, response times, failure rates, and dependencies).
- **Log Analytics:** The tool/workspace in the Azure portal used to edit and run **KQL (Kusto Query Language)** log queries against data collected by Azure Monitor Logs.

---

### Q2. What is Azure Policy and how does it differ from RBAC?

| Feature | Azure Policy | Azure RBAC |
|---------|--------------|------------|
| **Focus** | Focuses on **Resource Properties** during deployment and for already existing resources. | Focuses on **User Actions** (who can do what). |
| **Example** | "Only allow VMs to be created in the `East US` region." or "Require the `Environment` tag on all resource groups." | "User A can Restart this VM, but cannot Delete it." |
| **Enforcement** | Can automatically remediate non-compliant resources (DeployIfNotExists) or block creation. | Grants or denies access to APIs. |

> **Note:** RBAC evaluates *before* an action is performed. Azure Policy evaluates *after* RBAC authorizes the action, but before the resource is created.

---

### Q3. Explain Azure Management Groups.
Management Groups are containers that help you manage access, policy, and compliance across multiple **Subscriptions**.

**Hierarchy:**
`Root Management Group \-> Sub Management Groups \-> Subscriptions \-> Resource Groups \-> Resources`

Instead of applying an Azure Policy or RBAC role to 50 individual subscriptions, you apply it once to the Management Group housing those subscriptions, and they inherit the rules automatically.

---

### Q4. What is Azure Advisor?
A free, personalized cloud consultant that helps you follow best practices to optimize your Azure deployments.

**It provides recommendations across 5 categories:**
1. **Reliability:** (e.g., Add VMs to an availability set).
2. **Security:** (Integrates with Microsoft Defender for Cloud).
3. **Performance:** (e.g., Scale down an underutilized Database).
4. **Cost:** (e.g., Buy Reserved Instances, shutdown unused VMs).
5. **Operational Excellence:** (e.g., Create Azure Service Health alerts).

---

### Q5. How can you control costs in Azure?
1. **Budgets & Alerts:** Set spending limits at the Subscription/Resource Group level. Notify stakeholders via email/SMS when a threshold (e.g., 80%) is reached.
2. **Azure Policy:** Restrict the creation of expensive VM SKUs or force deployments only in cheaper regions.
3. **Reserved Instances & Savings Plans:** Commit to 1 or 3 years for deep discounts.
4. **Azure Hybrid Benefit:** Bring on-premises Windows Server or SQL Server licenses to Azure.
5. **Autoscaling:** Ensure PaaS services scale down during off-hours.
6. **Azure Advisor:** Act on cost-saving recommendations.

---

### Q6. What is the overarching purpose of Azure Blueprints? (Note: Being deprecated in favor of Template Specs)
Azure Blueprints enable cloud architects to define a repeatable set of Azure resources that implement and adhere to an organization's standards, patterns, and requirements.
A Blueprint consists of:
- Role Assignments (RBAC)
- Policy Assignments
- ARM Templates
- Resource Groups

**Difference from ARM Templates:** ARM Templates deploy resources once. Blueprints keep track of the deployment; if someone alters a resource created by a blueprint, the blueprint can lock it or revert it (enforcing compliance over time).
