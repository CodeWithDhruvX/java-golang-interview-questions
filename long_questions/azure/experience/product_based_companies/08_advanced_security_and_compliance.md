# 📘 08 — Advanced Security & Compliance
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Advanced

---

## 🔑 Must-Know Topics
- Zero Trust Architecture concepts
- Microsoft Defender for Cloud (Cloud Security Posture Management - CSPM)
- Microsoft Sentinel (SIEM/SOAR)
- VNet Injection and Private Link
- Customer-Managed Keys (BYOK)

---

## ❓ Most Asked Questions

### Q1. What is the "Zero Trust" model?
Zero trust is a security framework requiring all users, whether in or outside the organization's network, to be authenticated, authorized, and continuously validated.

**Three Guiding Principles:**
1. **Verify Explicitly:** Always authenticate and authorize based on all available data points (identity, location, device health, service, data classification). (Implementation: Azure AD Conditional Access).
2. **Use Least Privilege Access:** Limit user access with Just-In-Time (JIT) and Just-Enough-Access (JEA) policies, risk-based adaptive policies, and data protection.
3. **Assume Breach:** Minimize blast radius by segmenting access and verifying end-to-end encryption. Use analytics to get visibility, drive threat detection, and improve defenses.

---

### Q2. Detail the difference between Microsoft Defender for Cloud and Microsoft Sentinel.
- **Microsoft Defender for Cloud:** A Cloud Security Posture Management (CSPM) and Cloud Workload Protection Platform (CWPP). It assesses your Azure (and AWS/GCP) resources against security benchmarks, gives you a Secure Score, and alerts on active threats (e.g., "Someone is brute-forcing your VM's RDP port" or "Your storage account is public").
- **Microsoft Sentinel:** A cloud-native Security Information and Event Management (SIEM) and Security Orchestration, Automation, and Response (SOAR) solution. It ingests logs from *everywhere* (Office 365, Defender, Firewalls, On-Prem servers), analyzes them using AI/ML to detect complex attack chains, and automates responses via Playbooks (Logic Apps).

---

### Q3. What is VNet Injection vs Private Endpoints?
When securing a PaaS service (like Azure Databricks or API Management), you want to remove it from the public internet.

- **VNet Injection:** The service is physically injected into a dedicated subnet in your VNet. The compute nodes of the PaaS service get IPs from your subnet and abide by your NSG rules. (Supported by APIM, App Service Environment, Databricks).
- **Private Endpoints (Private Link):** You create a Network Interface (NIC) in your VNet that maps to the PaaS resource. Traffic to the PaaS resource goes through the Azure Backbone via that private IP. (Supported by Storage, SQL, Cosmos DB, Key Vault).

---

### Q4. Explain the concept of "Customer-Managed Keys" (CMK) / "Bring Your Own Key" (BYOK).
By default, data at rest in Azure is encrypted using Microsoft-Managed Keys (MMK). Microsoft owns and rotates these keys.
In highly regulated industries (Finance, Healthcare), compliance dictates that the customer must have complete control over the encryption keys.
With CMK, the customer generates a key, stores it in **Azure Key Vault**, and configures the Storage Account or Azure SQL Database to use that specific key for encryption.
If the company revokes access to the Key Vault, all data in the Storage Account becomes instantly unreadable by anyone, including Microsoft.

---

### Q5. How do you secure an Azure Function from the public internet, ensuring only an API Management (APIM) instance can call it?
1. **Network Level (VNet Integration):** Use VNet Integration on the Function App and restrict inbound traffic using Access Restrictions (Service Endpoints/Private Endpoints) to only allow IPs from the APIM subnet.
2. **Identity Level (Managed Identity):** Enable a System-Assigned Managed Identity on APIM. On the Azure Function, configure Azure AD Authentication to require incoming requests to have a valid token belonging to the APIM Managed Identity.
3. **Key Level:** Require the Function's `Host Key` to be passed in the header by APIM (less secure than Identity, but easier).
