# Comprehensive Java Full Stack Interview Questions
*Based on user experience at Ascendion, KPMG, Elegant Microweb, Tatvasoft, and others.*

This document provides a deep dive into potential interview questions, ranging from high-level architectural decisions to low-level implementation details.

---

## 1. Core Java & Advanced Multithreading
*Context: "High-throughput file uploads", "Asynchronous handling", "Image processing optimization".*

### **Multithreading & Concurrency**
1.  **Async File Uploads:** You mentioned designing an asynchronous pipeline for file uploads.
    *   Did you use `CompletableFuture`, `@Async`, or a reactive approach (Spring WebFlux)?
    *   How did you handle thread pool configuration? If the upload queue fills up, what is your rejection policy (Abort, CallerRuns, etc.)?
    *   How did you manage memory consumption during large file streams to prevent `OutOfMemoryError`? (e.g., streaming directly to Blob Storage vs. loading into memory).
2.  **Image Processing Pipeline:**
    *   For the image processing optimization (60% storage reduction), did you use parallel streams?
    *   What specific Java libraries did you use (e.g., Java 2D, thumbnailator)?
    *   How did you handle CPU-intensive tasks without blocking the main application threads?

### **Memory Management**
3.  **Garbage Collection:**
    *   In your high-throughput applications, did you encounter any GC pauses? Which Garbage Collector were you using (G1, ZGC, CMS)?
    *   How would you diagnose a memory leak in a production Docker container? (Keywords: Heap dump, VisualVM, Prometheus metrics).

---

## 2. Spring Boot & Microservices Architecture
*Context: "Architected microservices ecosystem", "Spring Cloud", "Circuit Breaker", "Audit Platform".*

### **Microservices Patterns**
4.  **Service Communication:**
    *   How do your microservices communicate? (REST via Feign, or Messaging via Kafka/RabbitMQ?)
    *   If Service A calls Service B and Service B is down, how does the **Circuit Breaker** (Resilience4j/Hystrix) handle the fallback? What happens when the circuit is "Half-Open"?
5.  **Data Consistency (Saga/Transactions):**
    *   In the KPMG Audit Platform, how did you maintain data consistency across services? Did you use the **Saga Pattern** or **Two-Phase Commit**?
    *   How did you handle distributed transactions?
6.  **Configuration Management:**
    *   How did you manage configurations across different environments (Dev, QA, Prod)? (e.g., Spring Cloud Config Server, Azure App Configuration, Kubernetes ConfigMaps).

### **Spring Security & Identity**
7.  **RBAC & JWT:**
    *   You implemented strict RBAC. Did you use standard Spring Security Filter Chains?
    *   How did you handle **JWT** token invalidation or revocation (e.g., logout)? (Blacklisting, short expiry + refresh tokens).
    *   Did you implement any method-level security using `@PreAuthorize`?

---

## 3. Cloud Native, Kubernetes (AKS) & DevOps
*Context: "Kubernetes on AKS", "Azure DevOps", "Docker", "CI/CD".*

### **Kubernetes (K8s) & Docker**
8.  **Container Optimization:**
    *   How did you optimize your Docker images? Did you use **Multi-stage builds** to reduce image size?
    *   How do you handle secrets (DB passwords, API keys) in Kubernetes? (K8s Secrets, Azure Key Vault integration).
9.  **Scalability:**
    *   How did you configure **Horizontal Pod Autoscaling (HPA)**? Was it based on CPU/Memory or custom metrics (e.g., queue length)?
    *   Explain the lifecycle of a Pod. What is the difference between specific **Liveness** and **Readiness** probes?

### **CI/CD & Azure DevOps**
10. **Pipeline Automation:**
    *   Describe your CI/CD pipeline stages.
    *   How did you integrate **SonarLint/SonarQube**? Did you block PRs based on Quality Gates (e.g., < 80% coverage)?
    *   What was your deployment strategy? (Blue/Green, Canary, Rolling updates).

---

## 4. Database (SQL) & Performance Tuning
*Context: "Optimized RESTful API performance by 20%", "Better query design".*

### **Optimization Techniques**
11. **Query Optimization:**
    *   Can you give an example of a query you optimized? Did you solve an **N+1 problem**?
    *   How did you identify slow queries? (Hibernate statistics, slow query logs, APM tools).
12. **Indexing & Caching:**
    *   What types of indexes did you use?
    *   Did you implement any caching (Redis, Caffeine)? If so, what was your eviction policy (LRU, TTL) and how did you handle cache invalidation?

---

## 5. Frontend Integration (Angular/React)
*Context: "Integrated PowerBI", "React and Redux", "Lazy loading", "Angular Material".*

### **React & Redux**
13. **State Management:**
    *   In your CRM console, how did you structure your Redux store? Did you use **Redux Thunk** or **Saga** for async side effects?
    *   How did you prevent unnecessary re-renders in React? (e.g., `useMemo`, `useCallback`).

### **Angular & Integration**
14. **Performance:**
    *   You mentioned **Lazy Loading**. How did you split your modules?
    *   How did you integrate **PowerBI**? Was it via IFrame or the PowerBI JavaScript SDK? How did you handle authentication (Embed Tokens)?

---

## 6. System Design & Behavioral Scenarios
*Context: "Bridged the gap between product owners", "Reduction in downtime".*

### **System Design**
15. **Design a "Google Drive" clone (File Upload):**
    *   Based on your Azure Blob Storage experience, how would you design a system to upload and share massive files (10GB+)?
    *   How would you handle resumable uploads? (Chunking, Block IDs).

### **Behavioral**
16. **Conflict Resolution:**
    *   Tell me about a time you disagreed with a Product Owner about a technical requirement (e.g., Tatvasoft). How did you convince them?
17. **Failure Analysis:**
    *   Describe the biggest production bug you faced in the Contis Banking System. How did you debug it, fix it, and prevent it from happening again?

---

### **Bonus: "Cheat Sheet" Metrics to Memorize**
*   **15%** Ops Efficiency Boost (Ascendion)
*   **20%** API Performance Improvement (Elegant Microweb)
*   **25%** Dev Time Reduction (Audit Platform Reusable Lib)
*   **30%** Downtime Reduction (Wassel-UI Circuit Breaker)
*   **60%** Storage Reduction (Contis Image Processing)
