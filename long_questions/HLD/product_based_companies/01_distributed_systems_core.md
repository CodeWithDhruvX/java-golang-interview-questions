# High-Level Design (HLD): Distributed Systems Core (Product-Based Companies)

This section covers core distributed systems concepts frequently asked in HLD rounds at top product-based companies (FAANG, Uber, Airbnb, etc.). 

## 1. What are the key differences between stateful and stateless architectures?
**Answer:**
*   **Stateless Architecture:** Each request from a client contains all the information necessary to service the request. The server doesn't store any session context. 
    *   *Pros:* Highly scalable (just add more servers), easier to load balance, simpler recovery from failures.
    *   *Cons:* Can increase network overhead as each request must carry full state (e.g., tokens), harder to achieve low latency for session-heavy workflows.
*   **Stateful Architecture:** The server maintains the state of the client's session. Future requests depend on the stored state.
    *   *Pros:* Lower overhead per request, better for persistent connections like WebSockets or gaming.
    *   *Cons:* Harder to scale (requires sticky sessions or distributed states), complex failover mechanisms.

## 2. Explain the CAP Theorem and how it influences distributed database choices.
**Answer:**
The CAP Theorem states that a distributed system can only provide two of the following three guarantees simultaneously:
*   **Consistency (C):** Every read receives the most recent write or an error.
*   **Availability (A):** Every request receives a (non-error) response, without the guarantee that it contains the most recent write.
*   **Partition Tolerance (P):** The system continues to operate despite an arbitrary number of messages being dropped or delayed by the network between nodes.
*   *Application:* Since network partitions (P) are inevitable in distributed systems, architects must choose between Consistency and Availability.
    *   **CP Systems:** Choose Consistency over Availability in the event of a partition. Example: MongoDB, HBase, Redis (in certain configs), Zookeeper. If a partition occurs, the system might reject requests to ensure data isn't stale.
    *   **AP Systems:** Choose Availability over Consistency. Example: Cassandra, DynamoDB, CouchDB. The system will always return a response, but it might be stale. Eventual consistency is achieved later.

## 3. What is Eventual Consistency vs. Strong Consistency? Give use cases.
**Answer:**
*   **Strong Consistency:** After a write completes, any subsequent read (by any client) will return the updated value. It requires synchronous replication across nodes.
    *   *Use Cases:* Financial transactions, inventory management (e.g., booking the last seat on a flight), password updates.
*   **Eventual Consistency:** If no new updates are made to a given data item, eventually all accesses to that item will return the last updated value. Reads might return slightly stale data for a short period.
    *   *Use Cases:* Social media feeds (likes, comments), metrics dashboards, DNS propagation, Amazon cart (historically).

## 4. How does a Load Balancer work, and what are its standard routing algorithms?
**Answer:**
A load balancer distributes incoming network traffic across a group of backend servers to ensure no single server bears too much demand, improving responsiveness and availability.
*   **Layer 4 vs Layer 7:** L4 operates at the transport layer (TCP/UDP based on IP/Port). L7 operates at the application layer (HTTP/HTTPS, can route based on cookies, headers, URL paths).
*   **Algorithms:**
    *   *Round Robin:* Distributes requests sequentially across the pool.
    *   *Least Connections:* Sends traffic to the server with the fewest active connections.
    *   *IP Hash / Consistent Hashing:* Uses a hash of the client's IP address to map requests to the same server (good for sticky sessions handling and cache affinity).
    *   *Weighted Round Robin:* Assigns a weight to each server based on its capacity.

## 5. What is Consistent Hashing and why is it essential in distributed caches?
**Answer:**
In a traditional load balancer or cache node setup using `hash(key) % N` (where N is the number of servers), adding or removing a server changes the hash mapping for almost all keys, resulting in massive cache misses.
**Consistent Hashing:**
*   Maps both the servers and the data keys to the same abstract circle (hash ring).
*   A key is routed to the first server found by moving clockwise on the ring.
*   When a server is added or removed, only the keys mapped to that specific server (or its immediate neighbor) are reassigned. Most keys remain strictly on their original servers.
*   *Virtual Nodes:* To ensure an even distribution of keys (preventing hot spots), each physical server is represented by multiple "virtual nodes" scattered around the ring.

## 6. How would you design a Rate Limiter? (Architecture & Algorithms)
**Answer:**
A rate limiter controls the rate of traffic sent by a client or service, usually based on IP, user ID, or API key.
*   **Algorithms:**
    *   *Token Bucket:* Tokens are added to a bucket at a fixed rate. Each request consumes a token. Allows bursts.
    *   *Leaking Bucket:* Requests are added to a queue. Processed at a fixed rate. Smooths out traffic.
    *   *Fixed Window Counter:* Counters are maintained per window (e.g., 00:00 to 00:01). Can have bursts at the boundary.
    *   *Sliding Window Log:* Keeps a timestamp log of every request. Highly accurate but memory intensive.
    *   *Sliding Window Counter:* A hybrid approach taking a weighted average of the previous and current windows.
*   **Architecture:** Requires a fast in-memory store like Redis. The logic runs at the edge (API Gateway). Use Lua scripts in Redis to ensure rate-limiting check and decrement operations are atomic to avoid race conditions.

## 7. Explain Data Partitioning (Sharding). What are its benefits and challenges?
**Answer:**
Sharding is the process of horizontally splitting a large database into smaller, more easily managed pieces called shards, which are spread across multiple servers.
*   **Benefits:** High availability, faster query response (less data per shard), increased write throughput.
*   **Challenges/Disadvantages:**
    *   *Joins across shards:* Extremely expensive or impossible. Must denormalize data or perform joins at the application layer.
    *   *Rebalancing:* Moving data when a shard gets too large is complex and requires downtime or careful migration algorithms.
    *   *Hotspotting:* Depending on the shard key (e.g., user country where one country has 90% of users), one shard takes all the load.
*   **Sharding Strategies:** Range-based sharding (e.g., users A-C on shard 1), Hash-based sharding (hash of user_id), Directory-based (a lookup table maintains mapping).
