# 📘 05 — Azure Security & Identity
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Microsoft Entra ID (formerly Azure AD) vs Windows Server AD
- Role-Based Access Control (RBAC) vs Azure AD Roles
- Managed Identities (System-Assigned vs User-Assigned)
- Azure Key Vault
- Multi-Factor Authentication (MFA) and Conditional Access

---

## ❓ Most Asked Questions

### Q1. What is the difference between Microsoft Entra ID (Azure AD) and Windows Server AD?

| Feature | Microsoft Entra ID | Windows Server AD |
|---------|--------------------|-------------------|
| **Identity Type** | Cloud-based Identity and Access Management. | On-premises traditional identity management. |
| **Protocols** | Uses HTTP/HTTPS (SAML, OAuth 2.0, OpenID Connect). | Uses LDAP, Kerberos, NTLM. |
| **Hierarchy** | Flat structure (no OUs - Organizational Units). | Hierarchical structure (Forests, Domains, OUs). |
| **Devices** | Supports Azure AD Join / Intune MDM. | Group Policy Objects (GPO). |

---

### Q2. Explain Azure Role-Based Access Control (RBAC). What are its core components?
RBAC helps manage who has access to Azure *resources*, what they can do with those resources, and what areas they have access to.

**Three core concepts:**
1. **Security Principal:** WHO? (User, Group, Service Principal, or Managed Identity).
2. **Role Definition:** WHAT? A collection of permissions (e.g., `Owner` (full access), `Contributor` (can't grant access), `Reader` (read-only)).
3. **Scope:** WHERE? The level at which the access applies (Management Group -> Subscription -> Resource Group -> Resource).
   - *Example: Assigning `Virtual Machine Contributor` to `User A` at the `Resource Group 1` Scope.*

---

### Q3. What is the difference between Azure RBAC Roles and Microsoft Entra ID Roles?
- **Azure RBAC Roles:** Control access to **Azure Resources** (VMs, Storage, VNets). Managed in the Azure portal under "Access control (IAM)".
- **Microsoft Entra ID Roles:** Control access to **Tenant-level Identity and Admin features** (e.g., creating users, assigning licenses, resetting passwords). Examples: `Global Administrator`, `User Administrator`. Managed inside Microsoft Entra ID.

---

### Q4. What is a Managed Identity? Why is it better than Service Principals?
A Managed Identity is an identity automatically managed in Entra ID for applications to use when connecting to resources that support Entra ID authentication (like Key Vault, Azure SQL).

**Why it's better:**
- **No Credentials Management:** You don't have to manage credentials (secrets/passwords) in your code or config files. Azure handles rotating the underlying certificate automatically.
- **Cost:** Free of charge.

**Types:**
- **System-Assigned:** Tied to the lifecycle of the Azure resource. If the VM is deleted, the identity is deleted.
- **User-Assigned:** Created as a standalone Azure resource. Can be assigned to multiple VMs. Useful when multiple resources need the exact same access.

---

### Q5. What is Azure Key Vault?
Azure Key Vault is a cloud service for securely storing and accessing secrets.

**What does it store?**
- **Secrets:** API keys, database connection strings, passwords.
- **Keys:** Cryptographic keys used for encryption (e.g., Transparent Data Encryption).
- **Certificates:** Provision, manage, and deploy public and private SSL/TLS certificates.

**Best Practice:** Have your application authenticate to Key Vault via a Managed Identity to retrieve the database connection string at runtime, keeping all secrets out of source code.

---

### Q6. What are Conditional Access Policies?
Conditional Access is the tool used by Entra ID to bring signals together, to make decisions, and enforce organizational policies based on the concept of **Zero Trust**.

**It works functionally on `IF-THEN` statements.**
- **IF (Signal):** User is coming from a high-risk IP, or using an unmanaged device, or accessing a highly sensitive app.
- **THEN (Decision):** Require MFA (Multi-Factor Authentication), Block access, or Require device to be compliant.

---

### Q7. What is a Service Principal?
An application whose tokens can be used to authenticate and grant access to specific Azure resources from a user-app, service, or automation tool (like a CI/CD pipeline in GitHub Actions or Jenkins). It is essentially an "identity" for an application, whereas a standard user account is an identity for a person.
