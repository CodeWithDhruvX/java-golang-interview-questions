# Azure Interview Questions & Answers (Summary)

> **Quick reference guide with concise explanations for Azure interview questions**

---

## 🔹 Azure Fundamentals (Questions 1-8)

**Q1: What is Microsoft Azure and why is it used?**
Microsoft Azure is a cloud computing platform providing a wide range of services (compute, storage, networking, AI, etc.) to build, deploy, and manage applications globally. It is used for scalability, cost-efficiency (pay-as-you-go), high availability, and global reach.

**Q2: Explain IaaS, PaaS, and SaaS with real examples.**
**IaaS (Infrastructure as a Service):** Rent IT infrastructure (VMs, Storage, Networks). You manage OS + Apps. Example: Azure VM.
**PaaS (Platform as a Service):** Rent platform for developing/deploying apps. Cloud manages OS/Runtime. You manage App + Data. Example: Azure App Service, Azure SQL.
**SaaS (Software as a Service):** Rent finished software. Cloud manages everything. Example: Microsoft 365, Outlook.

**Q3: What is an Azure subscription and resource group?**
**Subscription:** A logical container for billing and access control. It links user accounts to resources.
**Resource Group (RG):** A logical container that holds related resources for an Azure solution. Resources in an RG share the same lifecycle (deployed, updated, and deleted together).

**Q4: Difference between Azure region and availability zone?**
**Region:** A geographical area containing one or more datacenters connected by a low-latency network.
**Availability Zone (AZ):** Physically separate datacenters *within* a region. Each AZ has independent power, cooling, and networking to prevent failure of the entire region.

**Q5: What is Azure Resource Manager (ARM)?**
The deployment and management service for Azure. It provides a management layer that enables you to create, update, and delete resources in your Azure account using templates, CLI, or Portal.

**Q6: What are the benefits of cloud over on-premises?**
Scalability (elasticity), OpEx (pay-as-you-go) instead of CapEx (upfront cost), High Availability (global datacenters), Security (managed by MS), and Reduced Maintenance (no hardware management).

**Q7: What is SLA in Azure?**
Service Level Agreement (SLA) is a financial guarantee from Microsoft regarding uptime and connectivity (e.g., 99.9% availability). If Azure fails to meet it, you get service credits.

**Q8: What is Azure Portal, CLI, and PowerShell?**
**Portal:** Web-based GUI for managing resources.
**CLI:** Cross-platform command-line tool (Bash/Cmd).
**PowerShell:** Windows-based command-line shell/scripting language for automation.

---

## 🔹 Compute Services (Questions 9-15)

**Q9: What is an Azure Virtual Machine (VM)?**
An on-demand, scalable computing resource (IaaS) that emulates a physical computer. You control the OS, software, and configuration. Used for hosting apps, databases, or dev environments.

**Q10: Difference between VM and App Service?**
**VM (IaaS):** You manage OS, patching, and middleware. Full control.
**App Service (PaaS):** Managed web hosting platform. Microsoft patches OS/middleware. You just deploy code. Better for web apps/APIs.

**Q11: What is Azure App Service Plan?**
It defines the compute resources (vCPU, RAM, Region, Pricing Tier) for an App Service. Multiple apps can share one plan to save costs.

**Q12: What is Azure Functions and use cases?**
A serverless compute service (FaaS) that runs code in response to events (triggers). You pay only when code runs. Use cases: File processing, API backends, scheduled tasks, event-driven workflows.

**Q13: What is AKS (Azure Kubernetes Service)?**
A managed Kubernetes service for orchestrating containerized applications. Azure manages the K8s control plane (free), you manage and pay for the worker nodes.

**Q14: What is autoscaling?**
Automatically adjusting resources (VMs, App Service instances) based on demand (CPU usage, queue length). **Scale Out** adds instances, **Scale In** removes them.

**Q15: Difference between horizontal and vertical scaling?**
**Horizontal (Scale Out):** Adding *more* machines/instances (e.g., 1 VM -> 5 VMs). Unlimited capability.
**Vertical (Scale Up):** Adding *power* to existing machine (e.g., 8GB RAM -> 16GB RAM). Limited by hardware.

---

## 🔹 Networking (Questions 16-24)

**Q16: What is a VNet (Virtual Network)?**
Your private network in the cloud. It enables Azure resources (VMs) to securely communicate with each other, the internet, and on-premises networks.

**Q17: What is a Subnet?**
A specific range of IP addresses within a VNet. It segments the network for organization and security (e.g., Frontend Subnet, Backend Subnet).

**Q18: What is an NSG (Network Security Group) and how does it work?**
A firewall for VNets. It contains security rules allowing or denying inbound/outbound traffic based on IP, port, and protocol (e.g., Allow port 80, Deny port 22).

**Q19: Difference between NSG and Azure Firewall?**
**NSG:** Filters traffic at the subnet/NIC level (Layer 3/4). Basic, free.
**Azure Firewall:** Managed, stateful firewall service (Layer 3-7). Centralized, supports Threat Intelligence, FQDN filtering. Paid.

**Q20: What is Azure Load Balancer?**
A Layer 4 (Transport) load balancer. Distributes inbound TCP/UDP traffic across healthy backend VMs based on hash-based routing.

**Q21: Difference between Load Balancer and Application Gateway?**
**Load Balancer (L4):** Distributes traffic based on IP/Port. Simple, fast.
**App Gateway (L7):** HTTP/HTTPS load balancer. Supports SSL termination, URL-based routing, Cookie affinity, and WAF (Web Application Firewall).

**Q22: What is VPN Gateway?**
Sends encrypted traffic between an Azure VNet and an on-premises location over the public internet (Site-to-Site VPN) or between VNets (VNet-to-VNet).

**Q23: What is ExpressRoute?**
A private, dedicated connection between your on-premises datacenters and Azure. Faster, more reliable, and secure than VPN, as it doesn't traverse the public internet.

**Q24: Difference between Public IP and Private IP?**
**Public IP:** Reachable from the internet.
**Private IP:** Reachable only within the VNet or connected networks.

---

## 🔹 Storage Services (Questions 25-31)

**Q25: What is an Azure Storage Account?**
A unique namespace for your Azure storage data. It contains Blob, File, Queue, and Table storage services.

**Q26: Types of Azure Storage services?**
**Blob:** Unstructured object storage (images, logs).
**Files:** Managed file shares (SMB/NFS).
**Queues:** Messaging for decoupling apps.
**Tables:** NoSQL key-value store.

**Q27: What is Blob Storage?**
Object storage optimized for storing massive amounts of unstructured data (text, binary data). Types: Block Blob (files), Append Blob (logs), Page Blob (disks).

**Q28: Difference between Blob Storage and File Storage?**
**Blob:** Accessed via REST API or SDK. Flat structure. Good for app data.
**Files:** Accessed via SMB protocol (like a network drive). Hierarchical folders. Good for legacy apps requiring shared disk.

**Q29: What is LRS, GRS, ZRS?**
**LRS (Locally Redundant):** 3 copies in 1 datacenter (cheap, lowest durability).
**ZRS (Zone Redundant):** 3 copies across 3 AZs in one region (survives DC failure).
**GRS (Geo-Redundant):** 6 copies (3 in primary region, 3 in secondary paired region). Survives region failure.

**Q30: What is Azure Backup?**
managed, secure service to back up data from Azure VMs, SQL DBs, and file shares to a Recovery Services Vault.

**Q31: What is Azure Site Recovery (ASR)?**
Disaster Recovery (DR) service. Replicates workloads (VMs) from a primary site to a secondary location. Failing over ensures business continuity during outages.

---

## 🔹 Security & Identity (Questions 32-40)

**Q32: What is Azure Active Directory (AAD) / Entra ID?**
Microsoft's cloud-based identity and access management service. Handles authentication (Login) and authorization (Access) to resources.

**Q33: Difference between Azure AD and Windows AD?**
**Windows AD:** On-premises, uses Kerberos/NTLM, hierarchical (OU, Forests). Designed for devices/servers.
**Azure AD:** Cloud, uses HTTP/REST (OIDC, OAuth, SAML), flat structure. Designed for web apps/SaaS.

**Q34: What is RBAC (Role-Based Access Control)?**
Authorization system to manage who has access to Azure resources, what they can do, and where (scope).

**Q35: What are Azure roles?**
Built-in definitions of permissions.
**Owner:** Full access + can manage access.
**Contributor:** Create/manage resources, cannot manage access.
**Reader:** View only.

**Q36: What is MFA (Multi-Factor Authentication)?**
Requires more than one verification method to sign in (Password + Phone/App). drastically increases security.

**Q37: What is Managed Identity?**
An automatically managed identity in Azure AD for applications. Eliminates storing credentials (like client secrets) in code. Apps use it to get tokens for Azure services (Key Vault, SQL).

**Q38: What is Azure Key Vault?**
Securely stores and manages secrets (passwords, connection strings), keys (encryption), and certificates.

**Q39: What is Defender for Cloud?**
Cloud security posture management (CSPM) and threat protection (CWP) tool. Scans resources for vulnerabilities and recommends security fixes.

**Q40: What is Conditional Access?**
Policies that enforce rules like "If user is not in corporate network, require MFA" or "Block risky sign-ins".

---

## 🔹 Database Services (Questions 41-46)

**Q41: What is Azure SQL Database?**
A fully managed (PaaS) relational database engine based on SQL Server. Handles upgrades, patching, backups, and monitoring automatically.

**Q42: Difference between Azure SQL and SQL Server?**
**SQL Server:** Product you install on a VM/On-prem (IaaS/Physical). You manage everything.
**Azure SQL:** Managed service (PaaS). MS manages hardware/software.

**Q43: What is Cosmos DB?**
Globally distributed, multi-model NoSQL database. Supports APIs for SQL, MongoDB, Cassandra, Gremlin, and Table. Low latency, high availability.

**Q44: What APIs are supported in Cosmos DB?**
Core (SQL), MongoDB, Cassandra, Gremlin (Graph), Table (Key-Value), and PostgreSQL.

**Q45: What is backup and restore in Azure DB?**
Automatic backups (PITR - Point In Time Restore). SQL DB keeps backups for 7-35 days by default.

**Q46: What is Geo-replication?**
Creating readable secondary replicas of your database in different Azure regions. Used for DR and reading scaling.

---

## 🔹 Monitoring & DevOps (Questions 47-60)

**Q47: What is Azure Monitor?**
Comprehensive solution for collecting, analyzing, and acting on telemetry from cloud and on-premises environments (Metrics, Logs).

**Q48: What is Log Analytics Workspace?**
The container where Azure Monitor logs (and other logs) are stored. You use KQL (Kusto Query Language) to query this data.

**Q49: What is Application Insights?**
APM (Application Performance Management) service for developers. Detects performance anomalies, request failures, and dependencies in web apps.

**Q50: How do you monitor Azure resources?**
Use Azure Monitor for metrics (CPU %), Log Analytics for deep logs, and Alerts to notify admins of issues.

**Q51: How do you reduce Azure cost?**
Use Reserved Instances (1-3 yr commit), Spot VMs, Rightsizing (downsizing underutilized resources), Delete unused resources, Azure Advisor recommendations, Budgets.

**Q52: What is Azure Advisor?**
Personalized cloud consultant that analyzes your configuration and recommends best practices for Cost, Security, Reliability, Performance, and Excellence.

**Q53: What is Tagging and why is it used?**
Metadata (Key-Value pairs) assigned to resources. Used for billing allocation (Cost Center), resource organization (Env: Prod), and automation.

**Q54: What is Azure DevOps?**
Set of development tools: Boards (Agile), Repos (Git), Pipelines (CI/CD), Test Plans, Artifacts.

**Q55: What is CI/CD?**
**CI (Continuous Integration):** Auto-build and test code on commit.
**CD (Continuous Deployment):** Auto-release to production environments.

**Q56: What is Azure Pipeline?**
Cloud service to build and test code projects and push them to any target (Azure, AWS, On-prem). Defined via YAML.

**Q57: What is ARM Template?**
Infrastructure as Code (IaC) JSON files defining Azure resources. Ensures consistent deployments.

**Q58: Difference between ARM and Terraform?**
**ARM/Bicep:** Native to Azure, JSON-based (ARM) or DSL (Bicep).
**Terraform:** Open-source, multi-cloud, uses HCL (HashiCorp Configuration Language), manages state file.

**Q59: What is Infrastructure as Code (IaC)?**
Managing infrastructure (VMs, networks) using code/definition files rather than manual configuration. Enables version control and repeatability.

**Q60: What is Azure CLI?**
Command-line interface to manage Azure resources. Scriptable (Bash) and cross-platform.

---

## 🔹 Scenario Based Questions (Questions 61-70)

**Q61: How do you design a highly available application in Azure?**
Use Load Balancers/App Gateway, deploy across multiple Availability Zones (to survive DC failure) and Regions (to survive region failure) using Traffic Manager/Front Door. Use PaaS with built-in patterns.

**Q62: How do you secure Azure resources?**
Use RBAC (Least Privilege), Manage Identities, NSGs/Firewalls (Network security), Encryption (At rest/In transit), Defender for Cloud, and Private Endpoints.

**Q63: How do you migrate on-premises to Azure?**
Assess (Azure Migrate), Migrate (ASR for VMs, DMS for DBs), Optimize (Right-size). Strategies: Rehost (Lift & Shift), Refactor (PaaS), Rearchitect.

**Q64: How do you implement backup & disaster recovery?**
**Backup:** Use Azure Backup vaults for VMs/SQL to protect against corruption/deletion.
**DR:** Use Azure Site Recovery (ASR) to replicate VMs to another region for failover.

**Q65: How do you control access for users?**
Use Azure AD Groups assigned to RBAC roles at the Subscription or Resource Group scope. Enforce MFA and Conditional Access Policies.

**Q66: How do you monitor performance issues?**
Start with Azure Monitor Metrics (is CPU high?). Check Application Insights for code bottlenecks/slow dependencies. Query Log Analytics for error patterns.

**Q67: How do you design a scalable architecture?**
Use stateless apps, VM Scale Sets/AKS for compute, Queues for decoupling, Caching (Redis) for read loads, and CDN for static content.

**Q68: How do you handle production outages?**
Check Service Health (Azure side issue?). Check App Insights/Logs. Rollback recent changes. Scale out if load-related. Use DR failover as last resort. Communication is key.

**Q69: How do you optimize cost and performance?**
Autoscaling (performance + cost savings), Reservced Instances, selecting correct tiers (Standard vs Premium), cleaning up orphan disks/IPs.

**Q70: How do you deploy applications automatically?**
Use Azure DevOps Pipelines or GitHub Actions. CI triggers on commit -> Build Artifact -> CD triggers release -> Deploy to Staging -> Approval -> Deploy to Prod.

---

## 🔹 Real-Time Azure Administration (Questions 71-80)

**Q71: How do you structure subscriptions and resource groups in a real project?**
**Subscriptions:** Split by Department (HR, IT), Environment (Prod, Non-Prod), or Cost Center.
**Resource Groups:** Group by Lifecycle (App1-Web-RG, App1-DB-RG) or Environment (Dev-RG, Prod-RG).

**Q72: How do you move resources between resource groups or subscriptions?**
Use the **Azure Portal (Move button)**, **PowerShell (`Move-AzResource`)**, or **CLI**. Both source and destination must be active, and the resource type must support moving.

**Q73: What naming conventions do you follow in Azure?**
Standard: `Resource-App-Env-Region-Instance`.
Example: `vm-webapp-prod-eastus-001`. Consistent naming (Cloud Adoption Framework) enables easy identification and automation.

**Q74: How do you manage multiple environments (Dev / QA / Prod)?**
Use **Subscriptions** (Prod Sub, Non-Prod Sub) to isolate billing/access. Use **DevOps Pipelines** with variable groups to deploy the same code/infrastructure to different environments.

**Q75: How do you handle resource locks?**
Apply **CanNotDelete** (prevents deletion, allows modification) or **ReadOnly** (prevent any change) locks to critical resources (Prod DBs, VNets) to prevent accidental loss.

**Q76: Difference between soft delete and hard delete?**
**Soft Delete:** Data is marked for deletion but retained for a period (e.g., 7-90 days) and can be recovered (e.g., Key Vault, Storage Blobs).
**Hard Delete:** Permanent removal. Cannot be recovered.

**Q77: How do you manage Azure using RBAC in real projects?**
Assign roles to **Groups** (not users). Use built-in roles (`Contributor`, `Reader`) where possible. Use Custom Roles if specific granularity is needed. Apply at the highest necessary scope (RG or Subscription).

**Q78: What is Azure Policy and how have you used it?**
A governance tool to enforce rules and compliance. Used to: enforce specific regions, require tags (e.g., "CostCenter"), or restrict specific VM SKUs to control costs.

**Q79: How do you restrict users from creating costly resources?**
Use **Azure Policy** to deny creation of expensive VM SKUs (e.g., G-series) or restrict resource creation to specific allowable regions.

**Q80: How do you track who deleted a resource?**
Check the **Activity Log**. Filter by `Operation name = Delete` and look at the `Event initiated by` field to see the user identity.

---

## 🔹 Networking - Advanced (Questions 81-90)

**Q81: How do you design a secure hub-and-spoke VNet architecture?**
**Hub VNet:** Central point for shared services (Firewall, VPN Gateway, Bastion).
**Spoke VNets:** Workload specific (App1, App2).
**Peering:** Connect Spokes to Hub. Traffic flows through Hub (Firewall) for inspection using UDRs.

**Q82: How do you connect on-premises to Azure?**
**VPN Gateway (S2S):** Encrypted tunnel over internet. Good for small/medium workloads.
**ExpressRoute:** Private, dedicated MPLS connection. High speed, reliability, and security for enterprise.

**Q83: Difference between VPN Gateway and ExpressRoute?**
**VPN:** Public internet, lower bandwidth (up to 10Gbps aggregate), higher latency, cheaper.
**ExpressRoute:** Private network, high bandwidth (up to 100Gbps), low latency, SLA guaranteed, expensive.

**Q84: How does NSG traffic flow work?**
Rules are processed by **Priority** (lower number = higher priority). The first rule that matches traffic (Allow/Deny) wins. If no rule matches, default deny applies.

**Q85: What is User Defined Route (UDR)?**
Custom route table assigned to a subnet to override default Azure routing. Used to force traffic through a firewall (NVA) instead of going directly to the internet (0.0.0.0/0 -> Firewall IP).

**Q86: What is a Private Endpoint?**
A network interface in your VNet that connects privately to a PaaS service (Storage, SQL) using a private IP. Disables public internet access to that resource.

**Q87: Difference between Service Endpoint and Private Endpoint?**
**Service Endpoint:** Traffic stays on Azure backbone, but resource still has a public IP endpoint.
**Private Endpoint:** Resource gets a real private IP in your VNet. More secure.

**Q88: How do you expose a private app securely?**
Use **Application Gateway** with WAF (Public IP) -> Backend VNet (Private IP). Or use **Azure Front Door** with Private Link.

**Q89: How do you troubleshoot VNet connectivity issues?**
Use **Network Watcher**. Tools: `IP Flow Verify` (Check NSG), `Next Hop` (Check Routing), `Connection Troubleshoot` (End-to-end check).

**Q90: What happens if two VNets have overlapping IP ranges?**
They **cannot** be peered. You must re-address one VNet or use a NAT Gateway/NVA to translate addresses.

---

## 🔹 Compute & Scaling - Hands-On (Questions 91-100)

**Q91: How do you choose between VM, App Service, AKS, Functions?**
**VM:** Legacy apps, OS control needed.
**App Service:** Web apps, APIs, no OS mgmt.
**AKS:** Microservices, complex orchestration, container portability.
**Functions:** Event-driven, sporadic workloads, short tasks.

**Q92: What is VM Scale Set (VMSS)?**
Group of identical, load-balanced VMs. Supports auto-scaling (add/remove VMs) and auto-healing (replace unhealthy VMs).

**Q93: How do you perform VM patching?**
Enable **Automatic Guest Patching** for critical updates. Use **Azure Update Manager** to schedule and orchestrate patches across fleet of VMs.

**Q94: What is Azure Automation Account?**
Service to automate frequent, time-consuming, and error-prone cloud management tasks. Used for **Runbooks** (PowerShell/Python scripts) and Configuration Management (DSC).

**Q95: How do you handle auto-scaling in production?**
Configure specialized **Autoscale Rules** based on metrics (e.g., CPU > 70% -> Scale Out). Set minimum and maximum instance limits to prevent runaway costs or downtime.

**Q96: How do you deploy apps with zero downtime?**
Use **Deployment Slots** (App Service) or **Rolling Updates** (AKS/VMSS). Swap "Staging" slot to "Production" instance instantly after warming up.

**Q97: How do you manage custom VM images?**
Use **Azure Compute Gallery** (formerly Shared Image Gallery). It stores, versions, and replicates custom images across regions for consistent VM creation.

**Q98: What is Azure Bastion?**
PaaS service that provides secure RDP/SSH access to VMs directly from the Azure Portal over SSL (HTML5). No public IP needed on the VM.

**Q99: How do you secure VM access?**
Disable Public IP. Use **Bastion**. Use **JIT (Just-In-Time) VM Access** (open ports only when needed). Enforce MFA for RDP/SSH users.

**Q100: How do you troubleshoot VM performance issues?**
Check **VM Insights** (Azure Monitor). Look for high CPU, Memory, or Disk IOPS throttling. Check Boot Diagnostics (Screenshot/Serial Log).

---

## 🔹 Storage & Data - Advanced (Questions 101-110)

**Q101: How do you secure Azure Storage accounts?**
Disable "Allow Blob Public Access". Use Firewalls and Virtual Networks (limit to specific VNets). Use Private Endpoints. Rotate Access Keys periodically or use Managed Identity (preferred).

**Q102: What is SAS token and when do you use it?**
**Shared Access Signature:** A URI that grants restricted access rights (Time, Permissions, IP) to Azure Storage resources. Used to give temporary access to external clients without sharing the master key.

**Q103: What is Azure File Sync?**
Syncs on-premises file servers with Azure Files cloud shares. Keeps hot data local (caching) and tiers cold data to the cloud.

**Q104: How do you implement storage lifecycle management?**
Use **Lifecycle Management Policies** (JSON rules) to automatically move blobs to Cool/Archive tiers or delete them after X days of inactivity.

**Q105: Difference between Hot, Cool, and Archive tiers?**
**Hot:** Frequent access, higher storage cost, lowest access cost.
**Cool:** Infrequent access (>30 days), lower storage cost, higher access cost.
**Archive:** Rare access (>180 days), lowest storage cost, highest rehydration cost/time (hours).

**Q106: How do you restrict public access to Blob storage?**
Set `AllowBlobPublicAccess` property to **False** on the Storage Account. Configure container ACLs to Private.

**Q107: How do you recover deleted storage data?**
Enable **Soft Delete** for Blobs, Containers, and File Shares. It allows undeleting data within the retention period (e.g., 7 days).

**Q108: What is immutability (WORM)?**
**Write Once, Read Many:** Policy that prevents data from being modified or deleted for a specified period. Used for legal/compliance (e.g., financial logs).

**Q109: How do you encrypt data at rest and in transit?**
**At Rest:** Azure Storage Encryption (SSE) is enabled by default using Microsoft-managed keys (can use Customer-managed keys in Key Vault).
**In Transit:** Enforce `Secure Transfer Required` (HTTPS/TLS) on the storage account.

**Q110: What is Azure Backup vs Site Recovery?**
**Backup:** Copies data to restore *historical state* (User deleted a file).
**Site Recovery:** Replicates entire workloads (VMs) to a secondary region to restore *availability* during a disaster (Datacenter down).

---

## 🔹 Security - Highly Focused (Questions 111-120)

**Q111: How do you implement least privilege access?**
Audit current permissions. remove broad roles like "Owner" or "Contributor". Assign specific roles (e.g., "Virtual Machine Contributor") only to the specific resources users need. Use PIM.

**Q112: What is PIM (Privileged Identity Management)?**
Service in entra ID to manage, control, and monitor access. Users request "Just-In-Time" access to high-privilege roles (Admin) for a limited duration, often requiring approval/MFA.

**Q113: How do you rotate secrets in Key Vault?**
Enable **Key Rotation Policies** (auto-rotate keys). For secrets, use logic apps/functions or native integration (e.g., Azure SQL) to update the credential and the Key Vault version.

**Q114: What is Managed Identity vs Service Principal?**
**Managed Identity:** Wrapper around SP. Automatically managed by Azure. No password to handle/rotate. Linked to resource.
**Service Principal:** Account for non-human apps. You must manage/rotate the Client Secret manually.

**Q115: How do you secure public-facing applications?**
Use **DDoS Protection**. Put behind **WAF** (App Gateway/Front Door). Enforce HTTPS. Use Identity (Azure AD). Sanitize inputs.

**Q116: What is DDoS Protection?**
**Basic:** Free, enabled for everyone (protects Azure infra).
**Standard:** Paid, tuned to your app traffic, provides logging, alerting, and cost protection during attacks.

**Q117: How do you audit security compliance?**
Use **Defender for Cloud**. It compares your environment against standards (NIST, PCI-DSS, ISO) and gives a compliance score + remediation steps.

**Q118: What happens when a user leaves the organization?**
Disable account in AD (synced to Azure AD). Revoke active sessions. Their access tokens expire. Since RBAC is usually group-based, removing them from AD groups removes Azure access.

**Q119: How do you protect data from accidental deletion?**
Enable **Soft Delete** (Storage, Key Vault, VM Backup). Use **Resource Locks** (CanNotDelete). Implement **Azure Backup**.

**Q120: What is Zero Trust in real architecture?**
"Never Trust, Always Verify". Verify explicitly (MFA everywhere), Use Least Privilege, Assume Breach. Don't trust traffic just because it's inside the firewall (use Micro-segmentation).

---

## 🔹 Monitoring, Logs & Troubleshooting (Questions 121-130)

**Q121: How do you use Azure Monitor in production?**
Enable **Diagnostic Settings** on all resources to send logs to a Log Analytics Workspace. Use **Application Insights** for code-level monitoring and **Alerts** to notify teams of critical thresholds (e.g., CPU > 90%).

**Q122: Difference between metrics and logs?**
**Metrics:** Numerical values at a point in time (e.g., CPU usage = 80%). Lightweight, real-time, good for alerting.
**Logs:** Record of events (e.g., "User logged in", "Error 500 occurred"). detailed, good for debugging and analysis.

**Q123: How do you create alerts?**
Go to **Azure Monitor > Alerts**. Create an **Alert Rule** (Condition: Failed Requests > 5), define an **Action Group** (Email DevOps Team, trigger Webhook), and link them.

**Q124: How do you investigate high CPU or memory usage?**
Check **Metrics** to identify the time of spike. Use **Log Analytics** to see if load increased (Request count). For VMs, check **VM Insights** processes. For App Service, check **Diagnose and solve problems** tool or Profiler.

**Q125: How do you monitor application performance?**
Use **Application Insights**. Check the **Application Map** to see dependencies (DB, Redis) latency. Check **Failures** tab for exceptions and **Performance** tab for slow API calls.

**Q126: What is Log Analytics KQL?**
**Kusto Query Language.** A powerful, SQL-like language used to query logs in Azure.
Example: `AppRequests | where Success == false | summarize count() by OperationName`

**Q127: How do you track failed deployments?**
Check **Azure DevOps Pipeline logs** for build/release errors. Check **Activity Log** in Azure for "Write/Create" failures. Check **Deployment History** at the Resource Group level.

**Q128: What is Service Health vs Resource Health?**
**Service Health:** Status of Azure datacenters/services globally (e.g., "East US is down").
**Resource Health:** Status of your specific resource (e.g., "Your VM is rebooting").

**Q129: How do you handle production incidents?**
Detect (Alerts) -> Acknowledge -> Triage (Severity) -> Mitigation (Restore service) -> Root Cause Analysis (RCA) -> Prevention (Fix bug/Add alert).

**Q130: How do you do root cause analysis (RCA)?**
The "5 Whys" technique. Correlate Logs, Metrics, and Change Management (Deployments) timestamps to find the trigger.

---

## 🔹 Cost Management (Questions 131-140)

**Q131: How do you estimate Azure cost before deployment?**
Use the **Azure Pricing Calculator**. configuring specific services, tiers, regions, and usage hours to get an estimated monthly bill.

**Q132: How do you reduce cloud costs?**
(See Q51). Additionally: Shutdown dev resources at night, use **Spot Instances** for fault-tolerant workloads, and identify/delete **Orphaned Disks** and unattached Public IPs.

**Q133: What is Azure Budget and Alerts?**
A cost management feature to set a spending limit (e.g., $1000/month). It triggers alerts (Email) when usage hits 50%, 80%, 100% of the budget. It can trigger action groups (Runbooks) to shut down resources.

**Q134: What is Reserved Instance (RI)?**
Commitment to use a VM/DB for 1 or 3 years in exchange for up to **72% discount** compared to pay-as-you-go prices.

**Q135: Difference between Pay-as-you-go and Reserved VM?**
**Pay-as-you-go:** Flexible, no commitment, highest hourly rate. Good for short-term/spiky tests.
**Reserved:** Long-term commitment, lower hourly rate. Good for 24/7 production workloads.

**Q136: How do you identify unused resources?**
Use **Azure Advisor**. It scans for idle Virtual Network Gateways, unassociated IPs, and stopped VMs. Also use **Cost Analysis** views.

**Q137: What is Azure Advisor recommendation?**
(See Q52). It provides actionable recommendations categorized by High/Medium/Low impact, often with a "One-click fix" option.

**Q138: How do you control cost at scale?**
Use **Management Groups** to enforce Policy (restrict expensive SKUs) and Budgets across all subscriptions. Use **Tags** for chargeback/showback.

**Q139: How do you bill multiple teams?**
Enforce a compulsory "CostCenter" or "Department" **Tag**. Use **Cost Analysis** to group costs by this Tag and generate reports for each team.

**Q140: How do you optimize storage cost?**
Use **Data Lifecycle Management** policies to automatically move older files to Cool/Archive tiers. Enable **Compression** on backups.

---

## 🔹 DevOps & Automation - Advanced (Questions 141-150)

**Q141: How do you automate Azure resource creation?**
Use IaC (Bicep/Terraform). Write code defining the infra, commit to Git, run a Pipeline (`terraform apply`) to deploy it. Avoid clicking in the Portal.

**Q142: What is ARM template structure?**
JSON file containing: `$schema`, `contentVersion`, `parameters` (inputs), `variables` (internal values), `resources` (what to build), and `outputs` (results).

**Q143: Difference between ARM and Terraform (real usage)?**
Terraform is often preferred for **multi-cloud** strategy and its **State File** (knowing existing state). ARM/Bicep is preferred for **day-zero** support of new Azure features and no state file management.

**Q144: What is CI/CD pipeline flow?**
Code Commit -> Trigger Build -> Unit Tests -> Build Docker Image/Artifact -> Push to Registry -> Trigger Release -> Deploy to Dev -> Integration Tests -> Deploy to Prod.

**Q145: How do you manage secrets in pipelines?**
Never hardcode them. Use **Azure Key Vault**. The pipeline uses a Service Connection (SP) to fetch the secret at runtime (`$(secret-variable)`).

**Q146: What is Blue-Green deployment?**
Two identical environments (Blue=Live, Green=Staging). Deploy to Green, test it. Switch traffic (Load Balancer/DNS) to Green. Blue becomes idle/backup. Zero downtime, easy rollback.

**Q147: What is Canary deployment?**
Release new version to a small subset of users (e.g., 5%). Monitor health. If good, gradually increase traffic to 100%. If bad, route traffic back to old version.

**Q148: How do you roll back a failed deployment?**
**App Service:** Swap slots back.
**K8s:** `kubectl rollout undo`.
**Terraform:** Revert code commit and re-apply.
**Database:** Restore from backup (Point-in-time) if data corrupted.

**Q149: What is Git branching strategy?**
Rules for managing code. Common: **Gitflow** (Main, Develop, Feature branches) or **Trunk-Based** (Small, frequent commits to Main with Feature Flags).

**Q150: How do you handle configuration drift?**
Occurs when manual changes are made in Portal, differing from IaC code. Fix: Re-run Terraform/Pipeline to overwrite manual changes, or use **Drift Detection** tools to alert on manual changes.

---

## 🔹 Migration & DR (Questions 151-160)

**Q151: How do you migrate VMs to Azure?**
**Azure Migrate** service. Steps: 1. Discovery (Collector appliance), 2. Assessment (Sizing/Cost), 3. Migration (Replication -> Test Migration -> Cutover).

**Q152: Which tools are used for migration?**
**Azure Migrate**, **Database Migration Service (DMS)**, **Data Box** (Physical shared storage for large data), **AzCopy**.

**Q153: How do you plan Disaster Recovery (DR)?**
Identify Business Critical apps. Define **RTO** (how fast to recover) and **RPO** (how much data loss allowed). Choose region pairs (e.g., East US + West US). Use ASR/Geo-redundant storage.

**Q154: What is RTO and RPO?**
**Recovery Time Objective (RTO):** Max downtime allowed (e.g., "Must be up in 4 hours").
**Recovery Point Objective (RPO):** Max data loss allowed (e.g., "Can lose up to 15 mins of data").

**Q155: How do you test DR?**
Perform a **Test Failover** in ASR. It creates VMs in the DR region isolated from prod network. Validate app functionality. Does not impact production.

**Q156: How do you design multi-region architecture?**
Deploy App/DB in Primary and Secondary regions. Use **Traffic Manager** or **Front Door** to route traffic. Use **SQL Geo-Replication**. Active-Active (both serve traffic) or Active-Passive (DR standby).

**Q157: How do you failover an application?**
**Manual:** DNS update or Front Door weight change.
**Automated:** Front Door detects health probe failure and routes to secondary. ASR triggers VM failover scripts.

**Q158: How do you handle data consistency?**
In async replication (Cross-region), slight delay exists. Design apps to handle **Eventual Consistency**. Use strong consistency for critical financial data if latency permits.

**Q159: What is Azure Site Recovery (ASR) workflow?**
Install Mobility Service on source VM -> Replicate data to Azure Cache Storage -> Traffic processed to Recovery Region Disks -> On Failover, create VMs from disks.

**Q160: What are DR best practices?**
Document the plan. **Test regularly** (Quarterly). Automate failover scripts. Ensure capacity in the secondary region (Quota). Backups are NOT DR (Backups = Data retention, DR = Business continuity).

---

## 🔹 Behavioral + Practical (Questions 161-170)

**Q161: Describe a real Azure project you worked on.**
"I worked on migrating a monolithic .NET app to Azure. Rehosted DB to SQL Managed Instance, refactored Web App to App Service. Implemented CI/CD with Azure DevOps and monitored using App Insights."

**Q162: Biggest production issue you faced and how you solved it?**
"App was slow. Checked App Insights, found high DB dependency latency. Found missing index on SQL Table. Added index, latency dropped 90%. Configured alerts to catch this earlier."

**Q163: How do you handle pressure during outages?**
Stay calm. Communicate clear status to stakeholders ("We are investigating"). Focus on restoration (Roleback/Restart) first, then Root Cause. Follow the Runbook.

**Q164: How do you communicate incidents to stakeholders?**
Send regular updates (every 30 mins). Format: Current Status, Impact, Next Steps, ETA. Avoid technical jargon.

**Q165: What security issue did you fix recently?**
"Found a Storage Account with Public Access enabled. Disabled it, enabled Private Endpoint, and set up an Azure Policy to prevent creating public storage accounts in the future."

**Q166: How do you keep Azure skills updated?**
MS Learn, Azure Friday videos, reading the Azure Blog for new features, and preparing for certifications (AZ-104, AZ-305).

**Q167: What mistakes should be avoided in Azure?**
Leaving ports open (RDP/SSH), not using Budgets, Clicking-Ops (Manual) instead of IaC, Oversizing VMs (Costs), Ignoring Health Alerts.

**Q168: How do you review architecture?**
Use the **Well-Architected Framework** pillars: Cost, Security, Reliability, Performance, Operational Excellence. Check against these benchmarks.

**Q169: How do you handle customer data?**
Encrypt at rest/transit. Use strict RBAC. complying with GDPR/HIPAA. retain only as long as needed. Mask PII in logs.

**Q170: Why should we hire you as an Azure engineer?**
"I have hands-on experience with core/advanced services, a mindset for automation (IaC) and security, and I can handle real-world troubleshooting/incidents effectively."

---

## 🔹 Governance & Enterprise-Scale (Questions 171-180)

**Q171: What is a Management Group?**
A scope above Subscriptions. Used to manage access, policy, and compliance for multiple subscriptions efficiently (e.g., "Contoso Corp" > "IT Dept" > "Prod Sub").

**Q172: What is an Azure Landing Zone?**
A pre-configured environment (subscriptions, networking, identity, policies) that allows you to migrate/deploy workloads securely and at scale, following Microsoft best practices.

**Q173: Difference between Azure Policy, RBAC, and Blueprint?**
**RBAC:** Controls *who* can do what (User Access).
**Policy:** Controls *what* resources can be created (Compliance rules).
**Blueprint:** A package of everything (Composes RBAC + Policy + ARM Templates) to set up a new environment.

**Q174: What is a Policy Initiative?**
A collection of multiple Policy Definitions grouped together towards a specific goal (e.g., "ISO 27001 Compliance Initiative" containing 50+ policies).

**Q175: How do you design Azure for large enterprises?**
Use **Hub-and-Spoke** topology, **Management Groups** for hierarchy, **Identity** centralization (Azure AD), and **Landing Zones** for standardized subscription vending.

**Q176: What is Azure Front Door?**
Global Layer 7 Load Balancer and CDN. Routes traffic to the closest backend (region) via Microsoft's global edge network. Optimizes performance and provides WAF.

**Q177: Difference between Front Door vs Application Gateway?**
**Front Door:** Global (multi-region), routes/balances traffic between regions.
**App Gateway:** Regional, routes/balances traffic within a region (to VMs/Pods).

**Q178: What is Traffic Manager?**
DNS-based traffic load balancer. It resolves the user's DNS request to the IP of the closest healthy endpoint. Does not see/process the traffic payload itself (unlike Front Door).

**Q179: What is Private DNS Zone?**
Allows you to use your own domain names (e.g., `db.internal.corp`) within a VNet without exposing them to the internet or managing custom DNS servers.

**Q180: What is Azure Firewall Manager?**
A central management service to configure and manage firewall policies across multiple Azure Firewalls (Hubs) or Secured Virtual Hubs (Azure Virtual WAN).

---

## 🔹 Containers & Microservices (Questions 181-190)

**Q181: What is Azure Container Registry (ACR)?**
Managed private Docker registry to store and manage container images. Supports Geo-replication and scanning for vulnerabilities (Defender).

**Q182: Difference between AKS (Kubernetes) vs ACI (Container Instances)?**
**AKS:** Orchestrator for multiple containers/apps. Complex, scalable, prod-grade.
**ACI:** Serverless single container. Fast startup, simple, good for burst jobs or testing.

**Q183: How do you expose an AKS application securely?**
Use an **Ingress Controller** (Nginx/App Gateway) with a Public IP, handle TLS termination, and route to internal Pod Services. Apply **Network Policies** to restrict pod-to-pod traffic.

**Q184: What is Helm?**
Package manager for Kubernetes. Helps define, install, and upgrade complex K8s applications using "Charts" (templates).

**Q185: How do you monitor AKS?**
Enable **Container Insights** (Azure Monitor). It collects logs/metrics from controllers, nodes, and containers (stdout/stderr).

**Q186: What is Pod scaling vs Node scaling?**
**HPA (Horizontal Pod Autoscaler):** Adds more Pods when CPU usage is high.
**Cluster Autoscaler:** Adds more Nodes (VMs) when Pods cannot be scheduled due to lack of resources.

**Q187: What are Namespaces in AKS?**
Virtual clusters within a physical cluster. Used to isolate resources between teams or environments (e.g., `dev-ns`, `prod-ns`) within the same AKS.

**Q188: What is Azure Service Bus?**
Enterprise message broker with queues and publish-subscribe topics. Decouples applications and handles reliable async messaging.

**Q189: Difference between Service Bus vs Storage Queue?**
**Service Bus:** Advanced. Topics/Subs, Ordering (FIFO), Transactions, Dead-lettering.
**Storage Queue:** Simple. High volume (>80GB), no ordering guarantee, basic.

**Q190: What is Azure Event Grid?**
Event routing service. "Push-Push" model. Reacts to changes (e.g., "Blob Created") and pushes event to a handler (Function/Logic App). Near real-time.

---

## 🔹 Integration, Hybrid & Security (Questions 191-200)

**Q191: What is Azure Logic Apps?**
Serverless workflow integration platform (PaaS). Low-code/No-code designer to connect apps/data (e.g., "When tweet received, save to SQL").

**Q192: When do you use Logic Apps vs Functions?**
**Logic Apps:** Integration heavy (Connectors to Salesforce, Office365), visual workflow, declarative.
**Functions:** Compute heavy (Data processing, algorithmic), coding required (C#, Python).

**Q193: What is API Management (APIM)?**
Gateway to publish, secure, and analyze APIs. Features: Throttling, Caching, Auth translation, Developer Portal.

**Q194: What is Azure Arc?**
Bridge that extends the Azure platform to help you build applications and services with the flexibility to run across datacenters, edge, and multi-cloud environments (AWS/GCP), managed from Azure Portal.

**Q195: What is Microsoft Sentinel?**
Cloud-native SIEM (Security Information and Event Management) and SOAR solution. Collects security stats, detects threats using AI, and automates response.

**Q196: What is Azure App Gateway WAF?**
Web Application Firewall. Protects web apps from common exploits like SQL Injection and XSS (OWASP Top 10) at the edge.

**Q197: How do you handle SSL certificates in Azure?**
**Key Vault:** Store and manage certs centrally.
**App Service Managed Certs:** Free certs for custom domains.
**App Gateway:** Offload SSL termination.

**Q198: What is Data Residency?**
Requirement that data remains within a specific geographic boundary (e.g., "Data must stay in Germany"). Azure regions help comply with this.

**Q199: What is Azure Synapse Analytics?**
Limitless analytics service that brings together Big Data (Spark) and Enterprise Data Warehousing (SQL). Unified experience for specific data needs.

**Q200: What is Azure Databricks?**
Apache Spark-based analytics platform optimized for Azure. Collaborative workspace for data engineers and data scientists.

---
