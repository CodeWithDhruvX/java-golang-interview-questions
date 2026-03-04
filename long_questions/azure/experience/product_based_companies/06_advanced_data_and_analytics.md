# 📘 06 — Advanced Data & Analytics
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Advanced

---

## 🔑 Must-Know Topics
- Azure Synapse Analytics vs Databricks
- Data Lakes (ADLS Gen2)
- Azure Data Factory Pipeline Design
- Data Governance (Microsoft Purview)

---

## ❓ Most Asked Questions

### Q1. What is the difference between Azure Databricks and Azure Synapse Analytics?
Both are big data and analytics platforms, but they come from different ecosystems:

- **Azure Databricks:** Based on **Apache Spark**. Best for Data Engineering (ETL/ELT) using Python/Scala/R, Machine Learning, and streaming data architectures. Highly collaborative notebook environment.
- **Azure Synapse Analytics:** The evolution of Azure SQL Data Warehouse. It bridges the gap between Big Data (Spark) and Enterprise Data Warehousing (SQL). Best if the team is heavily invested in T-SQL and needs a unified workspace for Pipelines, Spark, and SQL pools.

---

### Q2. Explain Azure Data Lake Storage (ADLS) Gen2 vs Blob Storage.
ADLS Gen2 is built *on top* of Azure Blob Storage, but adds a **Hierarchical Namespace**.

- **Blob Storage:** Flat namespace. A "folder" is an illusion created by the object name (e.g., `image/2023/photo.jpg`). Renaming a directory requires renaming every single blob inside it (terrible performance).
- **ADLS Gen2:** True directory structure. Renaming a directory is an atomic, $O(1)$ operation. Crucial for Big Data tools like Spark and Hadoop to interact efficiently with the file system. ALDS Gen2 supports POSIX ACLs.

---

### Q3. How do you handle failure in an Azure Data Factory (ADF) pipeline?
Robust ETL pipelines require excellent error handling.
1. **Retry Policies:** Configure activities (like Copy Data) to retry automatically $X$ times with a specific interval if a transient error occurs (e.g., DB connection drop).
2. **Activity Dependency / Failure Paths:** Connect a second activity using a "Failure" dependency (red arrow). If the first activity fails, trigger an alert (Web Hook to Logic App -> Email/Slack) or write a row to an Error Log table.
3. **Tumbling Window Triggers:** If a pipeline for a specific hour fails, you can re-run just that specific time window without affecting real-time pipelines.

---

### Q4. What is ELT vs ETL in cloud data architectures?
- **ETL (Extract, Transform, Load):** Data is extracted, transformed heavily on an intermediate compute server, and then loaded into the data warehouse. (Traditional model, compute bottleneck).
- **ELT (Extract, Load, Transform):** Given the massive power of cloud data warehouses (Synapse), data is extracted and loaded *raw* into the Data Lake/Warehouse. Transformations happen *inside* the warehouse using native T-SQL or Spark. (Modern Cloud model).

---

### Q5. What is Microsoft Purview?
As companies amass petabytes of data across SQL, Cosmos, AWS, and on-prem, finding and governing data becomes impossible.
**Purview** is a unified data governance service.
- **Data Discovery:** Automates scanning of data sources to build a holistic map.
- **Data Catalog:** Provides a search engine for data scientists to find what data exists.
- **Data Lineage:** Shows how data flows (e.g., "This Power BI report comes from this Synapse table, which came from this ADF pipeline, which pulled from Salesforce").
- **Classification:** Automatically labels PII (Credit Cards, SSNs) across the entire data estate.
