# System Design (HLD) - Product Based Companies

This section contains top-tier High-Level System Design (HLD) interview questions commonly asked by FAANG and other hyper-growth product-based companies (e.g., Uber, Netflix, ByteDance).

These interviews test your ability to take a vague, massive problem (e.g., "Design YouTube") and structure it into a scalable, highly available, low-latency architecture.

## Topics Covered:
1.  [URL Shortener (TinyURL)](01_hld_url_shortener.md)
2.  [Chat Application (WhatsApp)](02_hld_chat_application.md)
3.  [Video Streaming Platform (Netflix/YouTube)](03_hld_video_streaming.md)
4.  [Ride Sharing App (Uber/Lyft)](04_hld_ride_sharing.md)
5.  [Distributed API Rate Limiter](05_hld_rate_limiter.md)
6.  [Distributed Cache (Redis/Memcached)](06_hld_distributed_cache.md)
7.  [E-commerce Platform (Amazon)](07_hld_ecommerce_amazon.md)
8.  [Ticketmaster (High Concurrency Booking)](08_hld_ticketmaster.md)
9.  [Distributed Web Crawler](09_hld_web_crawler.md)

## Success Criteria for System Design Interviews
1.  **Requirement Clarification:** Never start designing right away. Ask questions. Define Functional vs Non-functional requirements (Scale, Latency vs Consistency).
2.  **Back-of-the-envelope Estimation:** Calculate read/write ratios, traffic (QPS), and storage limits for 5 years to justify your database choices.
3.  **High-Level Architecture:** Draw the boxes and arrows. API Gateways, Load Balancers, App Servers, Caches, Databases, Message Queues (Kafka).
4.  **Database Design:** Choosing SQL vs NoSQL. Do you need ACID compliance (E-commerce)? Or is horizontal scalability and high write throughput more important (Social Media feeds -> Cassandra)?
5.  **Identify Bottlenecks:** Discuss single points of failure, cache stampedes, database sharding strategies, and trade-offs globally (CAP Theorem).
