# High-Level Design (HLD): Advanced System Design Case Studies

This section covers complex design problems frequently asked in senior engineering rounds (SDE-2, SDE-3, Staff) at top product companies.

## 1. Design a Distributed Rate Limiter
**Core Requirements:** Limit the number of requests a user/IP can make across multiple global gateway servers.
*   **Challenges:** If Server A and Server B both keep a local counter for "User X", User X could bypass the limit by hitting both servers.
*   **Design:**
    *   **Centralized Storage:** Fast in-memory store like Redis.
    *   **Architecture:** API Gateways check Redis before fulfilling requests.
    *   **Concurrency Issues:** If two concurrent requests from the same user hit two different API gateways, both might read `counter=4`, both increment to `5`, which bypasses the limit if the bucket size is 5.
    *   **Solution for Concurrency:**
        1.  **Redis Lua Scripts:** Lua scripts run atomically in Redis. The script checks the limit and increments it in one atomic step.
        2.  **Redis Sorted Sets (Sliding Window Log):** Store timestamps as the score in a ZSET. `ZREMRANGEBYSCORE` removes timestamps older than the window. `ZCARD` gets the current count.

## 2. Design a Web Crawler (e.g., Google Search Bot)
**Core Requirements:** Discover and download web pages, extract URLs, avoid infinite loops, respect `robots.txt`, highly concurrent.
*   **Architecture Components:**
    *   **Seed URLs:** Initial list of domains.
    *   **URL Frontier (Queue):** Stores URLs to be downloaded. Prioritizes URLs and ensures politeness (not hitting the same domain too fast).
    *   **HTML Fetcher:** Multi-threaded workers pulling from the Frontier.
    *   **DNS Resolver:** Needs heavy caching; otherwise, DNS lookups become the biggest bottleneck.
    *   **Content Parser & Extractor:** Extracts new links.
    *   **Duplicate Detection (Important):** 
        *   *URL Dedup:* Use a Bloom Filter to quickly check if a URL has already been visited (highly memory efficient). If Bloom Filter says "maybe", fall back to querying the database/Redis.
    *   **Storage:** Store HTML content in an Object Store (Amazon S3) or a big data store like HBase/Bigtable.

## 3. Design Google Drive / Dropbox (Cloud Storage)
**Core Requirements:** Upload/download files, sync files across multiple devices for the same user, handle offline edits, highly available.
*   **Metadata vs. Block Storage:**
    *   Never store the entire 5GB movie file in a relational DB.
    *   **Block Servers:** Split large files into smaller chunks (e.g., 4MB blocks). Hash each block (SHA-256). Store chunks in Amazon S3.
    *   **Metadata DB:** Store the file tree structure (`user_id, file_path, list_of_block_hashes`) in a SQL DB or MongoDB.
*   **Sync Mechanism (Delta Sync):** If a user changes 1 line in a 50MB text file, do NOT re-upload 50MB. Calculate the hash of the blocks. Only the block that changed will receive a new hash. Upload *only* that new 4MB block.
*   **Conflict Resolution:** If User edits on phone offline, and edits on laptop online, when the phone comes online, you have a conflict. Use vector clocks or standard "last write wins" with a separate "conflicted copy" saved for the user to resolve manually.

## 4. Design a Notification System (Push, SMS, Email)
**Core Requirements:** Send millions of timely push notifications, handle third-party API failures, track delivery.
*   **Architecture:**
    *   **Event Producers:** Microservices publish events like "Order Shipped" to a Message Broker (Kafka).
    *   **Notification Servers:** Consume events from Kafka. Determine *how* the user wants to be notified (Preferences DB) and look up their device tokens.
    *   **Priority Queues:** Push high-priority OTPs to a separate queue than low-priority promotional emails.
    *   **Workers & 3rd Party Integrations:** Workers pull from queues and call APNS (Apple), FCM (Firebase), Twilio (SMS), SendGrid (Email).
*   **Handling Failures:** If Twilio is down, workers must catch the HTTP 500, push the message to a "Retry Queue" with exponential backoff.
*   **Deduping:** Use a Redis cache with a short TTL. Before sending, check if `exists(notification_id)`. If yes, skip.

## 5. Design a Location-Based Service (e.g., Yelp / Nearby Friends)
**Core Requirements:** Given a user's latitude and longitude, find top venues within a 5km radius.
*   **The Problem with Standard SQL:** An index on `(lat, long)` requires doing a massive bounding-box query encompassing millions of rows, and then doing expensive math `sqrt((lat1-lat2)^2 + (long1-long2)^2)` to filter the circle. Too slow.
*   **Geospatial Indexing Solutions:**
    *   **Geohash:** Divides the world into an alternating grid of smaller and smaller rectangles, represented as a string (e.g., `9q8yy`). Two places close to each other usually share the same Geohash prefix. A query becomes an index-prefix match `SELECT * WHERE geohash LIKE '9q8yy%'`.
    *   **QuadTrees:** An in-memory tree data structure where each node represents a bounding box and has 4 children (splitting the box into 4 quadrants). Searching is very fast O(log N).
*   **Data Store:** Redis supports GeoHashes natively (`GEOADD`, `GEORADIUS`). PostgreSQL with PostGIS extension is the industry standard for persistent geospatial data.
