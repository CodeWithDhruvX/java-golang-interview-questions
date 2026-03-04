# High-Level Design (HLD): Capacity Estimation (Back-of-the-Envelope)

In senior engineering interviews (especially at FAANG), you will be asked to estimate the scale of the system before designing it. This ensures you choose the right data stores and architecture.

## 1. Why do we do Capacity Estimation?
**Answer:**
To determine the fundamental limits of the system we must build.
*   Are we designing for 100 Requests Per Second (RPS) where a single MySQL instance is fine?
*   Or are we designing for 100,000 RPS where we absolutely need Redis caching, DynamoDB, and heavy horizontal sharding?
*   It proves to the interviewer you understand the mathematical relationship between user behavior and hardware requirements.

## 2. Key Numbers Every Architect Should Memorize
**Answer:**
You don't need exact numbers, just orders of magnitude.
*   **Time:**
    *   L1 Cache reference: ~1 ns
    *   Mutex lock/unlock: ~25 ns
    *   Main Memory (RAM) reference: ~100 ns
    *   Read 1MB sequentially from SSD: ~1 ms
    *   Read 1MB sequentially from Disk: ~20 ms
    *   Send packet within same datacenter: ~0.5 ms
    *   Send packet CA to Netherlands to CA: ~150 ms
*   **Data Sizes:**
    *   Char: 1-2 bytes
    *   Integer/Timestamp: 4-8 bytes
    *   Short String (Name/Tweet): 50-200 bytes
    *   Image: 100KB - 2MB
    *   Video: 50MB - 1GB
*   **Math shortcuts:**
    *   1 Million requests / day = ~12 Requests Per Second (RPS).
    *   100 Million requests / day = ~1,200 RPS.

## 3. Example Execution: Estimate Capacity for Twitter/X
**Prompt:** Design Twitter. 300 Million Monthly Active Users (MAU). 50% use it daily. On average, a user posts 2 tweets a day and views the feed 5 times a day.

**Step 1: Traffic/RPS Estimation**
*   **Daily Active Users (DAU):** 300M * 50% = 150 Million.
*   **Write/Tweet RPS:** 
    *   150M DAU * 2 tweets/day = 300 Million tweets/day.
    *   300M / 24 hours / 3600 seconds ≈ 300M / 100,000 = **3,500 Write RPS**.
*   **Read/Feed View RPS:** 
    *   150M DAU * 5 views/day = 750 Million views/day.
    *   750M / 100,000 = **8,500 Read RPS**. (Assuming 1 API call per feed load).
*   *Conclusion:* Heavily read-oriented. Single DB can barely handle writes, needs sharding.

**Step 2: Storage Estimation**
*   **Assume:** A single tweet object (text, user_id, timestamp) averages 200 Bytes.
*   **Assume:** 10% of tweets contain a 100KB image. 1% contain a 10MB video.
*   **Daily Media Storage:**
    *   Text: 300M tweets * 200 Bytes = 60 GB / day.
    *   Images: 30M tweets * 100 KB = 3 TB / day.
    *   Videos: 3M tweets * 10 MB = 30 TB / day.
    *   *Total Daily Storage:* ~33 Terabytes / day.
*   **5-Year Storage Capacity Plan:**
    *   33 TB * 365 days * 5 years ≈ **60 Petabytes (PB)**.
*   *Conclusion:* Text data can fit in a moderately sharded RDBMS/NoSQL. Media strictly requires high-capacity object storage (S3).

**Step 3: Bandwidth Estimation**
*   **Ingress (Incoming to server):**
    *   33 TB per day uploaded.
    *   33 TB / 100,000 seconds = **~330 MB/sec**.
*   **Egress (Outgoing from server):**
    *   Assuming feed loads 20 tweets at once. 10% have media.
    *   Egress bandwidth will be significantly higher than Ingress. If 8,500 RPS feed views fetch ~2MB of data each, that's **~17 GB/sec** outbound data.
*   *Conclusion:* Absolute necessity for CDNs to handle the massive 17 GB/s egress media bandwidth.
