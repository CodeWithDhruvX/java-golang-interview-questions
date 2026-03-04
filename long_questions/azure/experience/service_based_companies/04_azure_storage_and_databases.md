# 📘 04 — Azure Storage & Databases
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Azure Storage Account basics (Blob, File, Queue, Table)
- Storage redundancy options (LRS, ZRS, GRS, RA-GRS)
- Blob Storage Access Tiers (Hot, Cool, Cold, Archive)
- Azure SQL Database (DTU vs vCore) vs SQL Managed Instance
- Cosmos DB basics (Partitioning and Consistency)

---

## ❓ Most Asked Questions

### Q1. What are the four types of services in an Azure Storage Account?
1. **Azure Blob Storage:** Object storage for unstructured data (images, videos, backups, logs).
2. **Azure Files:** Fully managed file shares mapped via SMB or NFS protocols. Can be mounted by on-premises and Azure VMs simultaneously.
3. **Azure Queue Storage:** Asynchronous message queuing for communication between application components.
4. **Azure Table Storage:** NoSQL key-value store for massive datasets without complex relationships. (Cosmos DB Table API is the newer equivalent).

---

### Q2. Explain the different redundancy options in Azure Storage.
To ensure durability, Azure keeps multiple copies of your data:
- **LRS (Locally Redundant Storage):** 3 copies within a single datacenter. Cheapest, but data is lost if the datacenter goes down.
- **ZRS (Zone-Redundant Storage):** 3 copies spread across 3 Availability Zones in the primary region. Protects against datacenter failure.
- **GRS (Geo-Redundant Storage):** LRS in the primary region (3 copies) + asynchronously replicated to a secondary paired region (3 copies). Total 6 copies. Protects against regional disasters.
- **RA-GRS (Read-Access GRS):** Same as GRS, but you can read from the secondary region even if the primary region hasn't failed (useful for global read performance).

---

### Q3. What are Blob Storage access tiers?

| Tier | Usage | Storage Cost | Access/Transaction Cost |
|------|-------|--------------|-------------------------|
| **Hot** | Data accessed frequently. | Highest | Lowest |
| **Cool** | Data accessed infrequently (stored for at least 30 days). | Lower | Higher |
| **Cold** | Data rarely accessed (stored for at least 90 days). | Even Lower | Even Higher |
| **Archive**| Long-term backup/compliance (stored for at least 180 days). Offline tier. | Lowest | Highest (requires "rehydration" taking hours) |

---

### Q4. Shared Access Signatures (SAS) vs Access Keys for Storage?
- **Access Keys (Account Keys):** Root passwords to the entire storage account. Avoid sharing these. If compromised, the attacker has full control.
- **Shared Access Signature (SAS):** A URI that grants restricted access rights to Azure Storage resources (e.g., specific blob/container). You define permissions (Read/Write), IP restrictions, and an expiration time.

> **Best Practice:** Always use SAS URLs (or Entra ID Managed Identities) for application access, never Access Keys.

---

### Q5. Explain Azure SQL Database purchasing models: DTU vs vCore.

| Model | Pricing based on | Target Audience |
|-------|------------------|-----------------|
| **DTU Model** | Database Transaction Unit (Blended measure of CPU, Memory, I/O). | Simple budgets, predictable workloads, straightforward tier sizing (Basic, Standard, Premium). |
| **vCore Model** | Allows independent scaling of compute (vCores) and storage. | Enterprise migrations, offers Azure Hybrid Benefit (bring your own SQL Server license to save money). |

---

### Q6. Difference between Azure SQL Database and Azure SQL Managed Instance (MI)?
- **Azure SQL Database:** PaaS. Fully managed, single or pooled databases. Excellent for modern cloud applications but has some limitations regarding server-level features (e.g., no SQL Server Agent, no cross-database querying natively).
- **SQL Managed Instance (MI):** Near 100% compatibility with on-premises SQL Server Enterprise Edition. Includes features like SQL Server Agent, Database Mail, Linked Servers, and Service Broker. Designed for lift-and-shift migrations of classic on-premises SQL applications.

---

### Q7. What are the 5 Consistency Levels in Cosmos DB?
Cosmos DB offers well-defined trade-offs between consistency, availability, and latency:
1. **Strong:** Reads are guaranteed to return the most recent committed version (highest consistency, lowest availability/highest latency).
2. **Bounded Staleness:** Reads lag behind writes by at most 'K' versions or 'T' time.
3. **Session (Default):** Guarantees read-your-own-writes within a specific session. Best for user-centric apps.
4. **Consistent Prefix:** Updates return in order, never out of order, but might be delayed.
5. **Eventual:** Order is not guaranteed. Replicas will eventually converge (highest availability/lowest latency, weakest consistency).

---

### Q8. What is a Partition Key in Cosmos DB and why is it important?
A partition key (e.g., `UserId`, `TenantId`) is essential for scaling in Cosmos DB. It determines how data is distributed across underlying physical partitions.

**A good partition key should:**
- Have a high cardinality (lots of distinct values).
- Distribute storage evenly.
- Distribute request/compute (RU/s) evenly across the keys.
> **Warning:** You cannot change a partition key after a container is created. Choosing a "hot" partition key (e.g., grouping all active requests into one physical partition) causes rate-limiting bottlenecks.
