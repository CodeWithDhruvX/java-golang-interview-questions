# 📘 03 — System Design on Azure
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Advanced

---

## 🔑 Must-Know Topics
- Microservices architecture mapping to Azure services
- Designing high-throughput data ingestion pipelines
- Scalable E-commerce backends
- Distributed Caching (Redis)

---

## ❓ Most Asked Questions

### Q1. How would you design a scalable Netflix-like video streaming service on Azure?

**Component Breakdown:**
1. **Global Entry:** **Azure Front Door** with globally configured WAF to handle routing and SSL termination.
2. **Static Content (UI/Images):** Hosted in **Azure Storage (Static Web Apps)** and distributed via **Azure CDN**.
3. **Video Content:** Stored in **Azure Blob Storage**. Use **Azure Media Services** for encoding the video into multiple bitrates (HLS/DASH).
4. **Metadata/Search:** Store movie metadata in **Cosmos DB** for fast global reads. Use **Azure AI Search** for complex, typo-tolerant searching.
5. **Microservices (Billing, User Profile, Recommendations):** Hosted on **Azure Kubernetes Service (AKS)**.
6. **Streaming Delivery:** Use **Azure CDN** heavily caching the video chunks at edge nodes closest to the user.

---

### Q2. How would you handle a sudden massive spike in traffic (e.g., ticket sales for a Taylor Swift concert)?
Spiky traffic will overwhelm databases quickly. The key is **asynchronous processing (Load Leveling)**.

1. **Edge Caching:** Cache the "Concerts list" page in **Azure Redis Cache** or **CDN**. Prevent database queries for static data.
2. **Queueing the Purchase:** When a user clicks "Buy", do *not* write directly to the database. Instead:
   - Accept the HTTP request via an **Azure Function** or **API Management**.
   - Immediately place a message representing the purchase intent onto an **Azure Service Bus Queue**.
   - Return an HTTP 202 (Accepted) immediately to the user, saying "Processing your order, please hold."
3. **Background Processing:** Background worker services pull from the Service Bus at a controlled rate (e.g., 500/sec) that the back-end Azure SQL Database can handle safely.
4. **Real-time Feedback:** Use **Azure SignalR Service** to push a notification to the user's browser once the background worker finishes the transaction.

---

### Q3. Explain how Azure Cache for Redis is used for scalability.
Redis is an in-memory data store used to drastically reduce database latency and load.

**Common Patterns:**
- **Cache-Aside:** The application tries to read from Redis. If it's a "Miss", it reads from SQL, writes the data to Redis with a TTL, and returns it.
- **Session State:** In a stateless web farm (multiple App Service instances), session state is stored in Redis so the user doesn't lose their cart if routed to a different server.
- **Pub/Sub:** Used for lightweight messaging between microservices.

---

### Q4. Describe an architecture for ingesting and processing millions of IoT telemetry messages per second.

**The "Lambda" or "Kappa" Architecture on Azure:**
1. **Ingestion:** IoT devices send JSON payloads to **Azure IoT Hub** or **Azure Event Hubs**. These services are designed to ingest millions of events per second.
2. **Stream Processing (Hot Path):** Use **Azure Stream Analytics** to analyze the data continuously in real-time (e.g., "Alert if temperature > 100°C over a 5-minute tumbling window"). Outputs alerts to a Service Bus Queue.
3. **Batch Processing (Cold Path):** Use Event Hubs Capture to save the raw telemetry as Avro/Parquet files in **Azure Data Lake Storage Gen2**.
4. **Analytics:** Use **Azure Databricks** or **Synapse Analytics** to train ML models or run daily batch reports on the historical data.

---

### Q5. What are the benefits of using Azure API Management (APIM) in a microservices system design?
If you have 50 microservices on AKS, exposing them directly to the internet is a security nightmare and violates the "API Gateway" pattern.
APIM acts as the central front-end for all APIs:
- **Routing:** Hides internal IP addresses and routes `/users` to the User Service and `/orders` to the Order Service.
- **Authentication:** Validates JWT tokens centrally. Internal microservices trust APIM, offloading auth checks.
- **Rate Limiting:** Prevents DDoS and limits partners to 1000 calls/hour.
- **Caching:** Caches frequent identical requests directly at the gateway to reduce downstream load.
