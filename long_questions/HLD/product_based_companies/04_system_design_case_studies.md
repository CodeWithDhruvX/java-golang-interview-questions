# High-Level Design (HLD): System Design Case Studies

During product-based company interviews, candidates are asked to design large-scale systems end-to-end. This document outlines key patterns for standard case studies.

## 1. Design a URL Shortener (e.g., TinyURL)
**Core Requirements:** Generate short alias for long URLs, redirect upon hitting short URL, highly available, track click metrics.
*   **Traffic/Data estimation:** Heavy read-to-write ratio (100:1).
*   **DB choice:** NoSQL (e.g., DynamoDB or MongoDB) is suitable. Schema: `short_hash -> {long_url, user_id, creation_date}`.
*   **Base62 Encoding:** A 7-character Base62 string (a-z, A-Z, 0-9) yields `62^7 = 3.5 trillion` combinations.
*   **Generating the Hash:**
    *   *Approach 1 (MD5):* Hash the long URL, take first 7 chars. Problem: Collisions.
    *   *Approach 2 (Global Counter):* Use Zookeeper or an RDBMS sequence generator to generate a unique integer ID, then convert to Base62.
    *   *Approach 3 (Token Key Generation Service - KGS):* Pre-generate a list of random 7-character strings and store them in a DB. Allocate them instantly when requested. Highly scalable.
*   **Caching:** Redis in front of the DB mapping `short_url -> long_url` handles the read-heavy load. (LRU eviction).

## 2. Design a Streaming Service (e.g., Netflix / YouTube)
**Core Requirements:** Upload video, process video into various resolutions, stream video to millions of clients globally without buffering.
*   **Storage:** 
    *   Videos stored in Object Storage (Amazon S3 / Azure Blob). 
    *   Metadata stored in RDBMS or MongoDB.
*   **Video Processing Pipeline (Async):** When a user uploads a video, push an event to a Message Queue (Kafka/RabbitMQ). Worker nodes pick up the task and encode the video into different formats (1080p, 720p) and chunk the video into segments using formats like HLS or DASH.
*   **Content Delivery Network (CDN):** Essential. The video chunks are distributed to Edge Servers globally. Users stream directly from the nearest CDN, drastically reducing latency and backbone bandwidth.
*   **Client Player:** Uses Adaptive Bitrate Streaming. Downloads chunks dynamically, switching resolution based on the user's current internet speed.

## 3. Design a Chat Application (e.g., WhatsApp / Facebook Messenger)
**Core Requirements:** 1-on-1 chat, group chat, online presence, low latency.
*   **Connection protocol:** WebSockets for bidirectional persistent connections between client and server.
*   **Chat Servers & Routing:** Millions of connections require multiple stateless Chat Servers. 
*   **Session Management:** Use Redis to maintain exactly which user is connected to which specific Chat Server instance.
*   **Sending a message:** User A sends a message to Chat Server A. Server A checks Redis for User B's location. Redis says User B is on Chat Server C. Server A forwards message via an internal RPC/Queue to Server C. Server C pushes it down the WebSocket to User B.
*   **Message Ordering:** Use a counter or Cassandra to store messages. Cassandra is excellent here due to fast writes and sequential reads.

## 4. Design Twitter (Newsfeed System)
**Core Requirements:** Post tweets, view feed of people you follow, celebrity accounts (millions of followers). Follows a highly read-heavy pattern.
*   **Classic Approach (Pull-based):** When User A loads feed, query DB: `SELECT * FROM tweets WHERE user_id IN (list of followings) ORDER BY time`. Expensive.
*   **Push Approach (Fan-out on Write):** Every user has pre-computed feed in Redis. When User A posts a tweet, background workers asynchronously append the tweet to the Redis feed of *every* follower. Feed loading is O(1).
*   **The Celebrity Problem (Hybrid Fan-out):** Pushing a Justin Bieber tweet to 100M active followers overloads the system.
    *   *Solution:* Pull for celebrities, Push for normal users. Viewers' feeds are composed of pushed tweets from normal friends, combined at read-time with pulled tweets from followed celebrities.

## 5. Design a Ride-Sharing App (e.g., Uber / Lyft)
**Core Requirements:** Match riders with drivers, real-time location tracking, dynamic pricing, handle millions of concurrent rides.
*   **Geospatial Indexing:** Use S2 Geometry Library or Geohash for efficient location-based queries. Store driver locations in Redis Geo for O(1) proximity searches.
*   **Matching Algorithm:** Multi-factor scoring considering distance, driver rating, availability, and ETA. Use Hungarian algorithm for optimal assignment in small areas.
*   **Real-time Tracking:** MQTT/WebSocket for bidirectional communication. Driver location updates every 2-5 seconds.
*   **Dynamic Pricing:** Surge pricing based on demand/supply ratio, weather, events, time of day.
*   **Payment Integration:** Multiple payment methods including UPI, cash, cards, wallets.

## 6. Design a Food Delivery System (e.g., Swiggy/Zomato)
**Core Requirements:** Restaurant discovery, order management, real-time delivery tracking, multiple payment options.
*   **Microservices Architecture:** User, Restaurant, Order, Delivery, Payment, Search services.
*   **Search & Discovery:** Elasticsearch with geo-based indexing, cuisine filtering, price range, ratings.
*   **Delivery Matching:** Geospatial algorithm for driver assignment, considering proximity, capacity, rating.
*   **Real-time Updates:** WebSocket for order status, location tracking, ETA calculations.
*   **Payment Processing:** UPI integration, COD handling, wallet payments, daily settlement.

## 7. Design a Payment Gateway (UPI-focused)
**Core Requirements:** Process millions of transactions, integrate with banks/NPCI, ensure security and compliance.
*   **Integration Layer:** Adapter pattern for multiple bank APIs, NPCI UPI protocol implementation.
*   **Security:** PCI DSS compliance, end-to-end encryption, tokenization, fraud detection.
*   **Transaction Processing:** Idempotency handling, retry logic, circuit breakers for bank failures.
*   **Settlement Engine:** Daily T+1 settlement, reconciliation, fee calculation.
*   **Risk Management:** Real-time fraud detection, ML models for anomaly detection.

## 8. Design a Banking Core System
**Core Requirements:** Account management, transaction processing, loan management, regulatory compliance.
*   **Core Banking Modules:** Customer, Account, Transaction, Loan, Card, Risk, Compliance services.
*   **Payment Processing:** NEFT/RTGS/IMPS integration, UPI support, card network connectivity.
*   **Database Design:** ACID compliance for financial transactions, multi-database strategy.
*   **Regulatory Compliance:** RBI guidelines, KYC/AML, audit trails, reporting requirements.
*   **High Availability:** Disaster recovery, data replication, 99.99% uptime guarantee.
**Core Requirements:** Match rider with nearest driver, track location in real-time, estimate ETA.
*   **Location Tracking:** Drivers hit the server every 4 seconds with GPS coordinates. Do NOT write this straight to RDBMS.
*   **Spatial Database:** Use heavily optimized systems like Redis GeoHash mapping or QuadTrees to quickly calculate objects within a specific radius.
*   **Architecture:**
    *   Gateway routes driver updates to Location Service.
    *   Location Service updates In-Memory QuadTree.
    *   Rider requests a cab. Request hits Dispatch Service.
    *   Dispatch Service queries Location Service for nearest 5 drivers.
    *   Dispatch Service sends push notifications to drivers (via Websockets) sequentially until accepted.
*   **Analytics/ETA:** Asynchronous pipeline via Kafka pushes historical route data to Spark/Hadoop for Machine Learning predictive models.
