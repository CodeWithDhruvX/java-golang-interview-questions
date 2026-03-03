# 📘 05 — Performance Optimization & Cost Management
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Advanced

---

## 🔑 Must-Know Topics
- Azure FinOps Practices
- Advanced Autoscaling Patterns
- Cosmos DB performance tuning (RUs, Partitioning)
- Content Delivery Networks (CDN) and Edge Caching
- SQL Database Indexing and Profiling

---

## ❓ Most Asked Questions

### Q1. How do you approach cost optimization in a massive Azure environment? (FinOps)
Cost optimization isn't a one-time task; it's a continuous lifecycle.
- **Visibility:** Use **Azure Cost Management** to create dashboards. Implement a strict **Tagging Policy** (via Azure Policy) requiring a `CostCenter` or `Project` tag on all resources.
- **Rightsizing:** Use **Azure Advisor** to identify underutilized VMs, SQL databases, or App Services and scale them down.
- **Commitments:** Purchase **Reserved Instances (1 or 3 years)** for predictable workloads (saves up to 72%) or **Azure Savings Plans** for flexible compute consumption.
- **Licensing:** Use **Azure Hybrid Benefit** to apply existing Windows Server/SQL on-prem licenses to the cloud.
- **Dev/Test:** Enable Auto-Shutdown for dev VMs at night. Dev/Test environments should use the cheapest SKU (e.g., App Service Free/Basic tier) or **Spot Instances**.

---

### Q2. How do you resolve a "Hot Partition" issue in Cosmos DB?
A hot partition occurs when a large percentage of requests (Reads or Writes) hit the same logical partition key, consuming all the provisioned Request Units (RUs) for that physical partition and causing HTTP 429 (Too Many Requests) errors.

**Resolutions:**
1. **Choose a Better Partition Key:** (e.g., Don't use `Date` as an IoT partition key if thousands of messages arrive currently; instead, use `{DeviceId}-{Date}`). *Note: You must migrate data to a new container to change a partition key.*
2. **Hash/Append Suffix:** If appending random suffixes (e.g., `TenantId_1`, `TenantId_2`), it spreads the writes but complicates the reads (you must read from all suffixed partitions).
3. **Increase RUs:** A temporary, expensive fix is to increase the provisioned RU/s.

---

### Q3. Explain the "Throttling" pattern and how `HTTP 429` is handled.
Throttling means a service is rejecting your requests because you exceeded limits.

**Handling it properly:**
Do NOT immediately retry the request, as this makes the throttling worse. Implement the **Exponential Backoff with Jitter** pattern using libraries like Polly (.NET).
- Attempt 1: Fails (HTTP 429)
- Attempt 2: Wait 2 seconds (± random jitter)
- Attempt 3: Wait 4 seconds (± random jitter)
- Attempt 4: Wait 8 seconds...

---

### Q4. Describe caching strategies to improve Application Performance.
1. **Browser Caching:** Setting `Cache-Control` headers so the browser caches CSS/JS.
2. **CDN (Azure CDN/Front Door):** Caching static assets at the edge (PoPs globally). This vastly reduces latency for global users and offloads traffic from your App Service.
3. **In-Memory Caching (.NET MemoryCache):** Very fast, but limited to a single instance. If the server scales out, cache miss rates increase.
4. **Distributed Caching (Azure Cache for Redis):** State is kept in a central redis cluster. All web servers access it. Best for microservices and session state.

---

### Q5. How do you identify performance bottlenecks in an Azure App Service?
1. **Application Insights:** Check the "Performance" tab. Look at the **Dependency Tracker** to see if the delay is in the App Service code or an external call (e.g., Azure SQL taking 5 seconds to respond).
2. **Log Profiler:** Run the Application Insights Profiler to get code-level traces. It shows exactly which method in your code is consuming CPU time.
3. **Metrics:** Check CPU Percentage and Memory Percentage. If memory is consistently high (memory leak), take a Memory Dump using the **Kudu Console** (`.scm.azurewebsites.net`) and analyze it with a tool like dotMemory.
