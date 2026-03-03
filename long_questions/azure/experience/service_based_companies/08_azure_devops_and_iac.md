# 📘 08 — Azure DevOps & IaC
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- CI/CD Concepts
- Azure DevOps Services (Repos, Pipelines, Boards)
- ARM Templates vs Bicep vs Terraform
- Infrastructure as Code (IaC)

---

## ❓ Most Asked Questions

### Q1. What are the core components of Azure DevOps?
Azure DevOps provides developer services for supporting teams to plan work, collaborate on code development, and build/deploy applications.
- **Azure Boards:** Agile planning tools, Kanban boards, work item tracking.
- **Azure Repos:** Unlimited cloud-hosted private Git repositories.
- **Azure Pipelines:** CI/CD that works with any language, platform, and cloud.
- **Azure Test Plans:** Rich, powerful tools for manual and exploratory testing.
- **Azure Artifacts:** Universal package repository for NPM, NuGet, Maven, and Python packages.

---

### Q2. What is Continuous Integration (CI) and Continuous Deployment (CD)?
- **CI (Build):** The practice of merging all developers' working copies to a shared mainline several times a day. Triggered by a PR or Commit. It automatically compiles the code, resolves dependencies, and runs unit tests, producing a deployable artifact.
- **CD (Release):** Takes the artifact produced by the CI pipeline and automatically deploys it to various environments (Dev, QA, UAT, Production). It handles environment-specific variables and approval gates.

---

### Q3. Compare ARM Templates, Bicep, and Terraform for Azure.

| Feature | ARM Templates (JSON) | Azure Bicep | HashiCorp Terraform |
|---------|-----------------------|-------------|---------------------|
| **Language** | JSON | Domain Specific Language (DSL), similar to HCL | HashiCorp Configuration Language (HCL) |
| **Provider** | Microsoft Native | Microsoft Native (compiles to ARM) | Third-party (HashiCorp) |
| **State File** | No state file (Azure itself is the state). | No state file. | Uses `.tfstate` files (must be managed securely). |
| **Multi-Cloud** | Azure Only | Azure Only | Any Cloud (AWS, GCP, Azure, etc.) |
| **Readability** | Very verbose, hard to read/author manually. | Very clean, easy to read, minimal syntax. | Easy to read, highly modular. |

---

### Q4. What are the advantages of using Infrastructure as Code (IaC)?
- **Consistency:** Eliminates environment drift. Dev, QA, and Prod are created identically.
- **Speed:** Provisioning complex multi-tier architecture takes minutes instead of days of manual clicking.
- **Version Control:** Infrastructure definitions are committed to Git. You can track who changed what (e.g., who opened port 22?), and rollback easily.
- **Reusability:** Code modules can be reused across different projects.

---

### Q5. How does an Azure Pipeline authenticate to Azure to deploy resources?
Through an **Azure Resource Manager Service Connection** inside Azure DevOps.
When you set this up, Azure DevOps creates a **Service Principal** in your target Azure Entra ID (Azure AD). That Service Principal is then granted an RBAC role (e.g., `Contributor`) on the specific Azure Subscription or Resource Group. The pipeline uses this identity to authenticate and execute deployments.

---

### Q6. What are Approval Gates in Azure Pipelines?
Pre-deployment or post-deployment conditions that must be met before a release can proceed to the next stage (e.g., from QA to Production).
- **Manual Approvals:** An authorized user (e.g., QA Lead or Change Manager) must click "Approve".
- **Automated Gates:** Querying Azure Monitor alerts (make sure no active alerts exist), checking security API results, or checking a ServiceNow ticket status before deploying.
