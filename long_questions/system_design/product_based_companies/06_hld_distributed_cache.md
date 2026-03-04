# System Design (HLD) - Distributed Cache (Redis/Memcached)

## Problem Statement
Design an in-memory Distributed Key-Value cache similar to Redis or Memcached. The system should sit between back-end application servers and the primary database to reduce database load and improve read latency.

## 1. Requirements Clarification
### functional Requirements
*   `put(key, value)`
*   `get(key)`
*   `delete(key)`
*   Eviction Policy (e.g., LRU).

### Non-Functional Requirements
*   **Fast:** Sub-millisecond read/write latency.
*   **Highly Scalable:** Ability to add more cache nodes seamlessly.
*   **High Availability:** Single node failures shouldn't bring down the caching layer.

## 2. The Core Problems of a Distributed Cache
A single machine's RAM is limited. To cache 10TB of data, we need 100 machines each with 100GB of RAM. The fundamental architectural challenge is: **How do we know which machine holds which key?**

## 3. Data Partitioning / Routing (Consistent Hashing)
If we have 4 cache servers (Node 0, 1, 2, 3) and a naive hashing algorithm: `hash(key) % 4`, then mapping works.
*   **The Problem:** If Node 1 crashes, the formula becomes `hash(key) % 3`. Suddenly, every single key hashes to a different node. The entire cache is invalidated instantly (Cache Miss Storm), crushing the main database.

*   **The Solution: Consistent Hashing.**
    We map both the Nodes and the Keys onto a hash ring (e.g., $0$ to $2^{32}-1$).
    1. Hash the `Node ID` to place the node on the ring.
    2. Hash the `Key` to place it on the ring.
    3. To find a key's node, move clockwise on the ring until you encounter a Node.
    *   **Advantage:** If a node crashes, only the keys belonging to that node are remapped to the next adjacent node. The rest of the cache remains entirely intact.
    *   **Virtual Nodes:** To ensure data is distributed evenly across all servers, we hash each physical node multiple times (e.g., `Node1_v1`, `Node1_v2`) to essentially dot the ring with thousands of virtual nodes.

## 4. High Availability & Replication
Cache servers run in memory. If a node restarts, all data is lost.
*   **Replication Architecture:** Master-Slave architecture. Each primary hash ring node (Master) has one or more replica nodes (Slaves).
*   **Writes:** Go to the Master node, then sync asynchronously to Replicas.
*   **Reads:** Can be served by Master or Replicas (improving read throughput).
*   **Failover:** If the Master dies, a Zookeeper/Consul cluster detects the heartbeat failure and promotes a Replica to Master.

## 5. Eviction Policies
When RAM is 100% full, new `put()` calls require deleting old data.
*   **LRU (Least Recently Used):** Eject keys that haven't been read in the longest time. (Implemented locally on each node via a Doubly Linked List + HashMap structure).
*   **TTL (Time-To-Live):** Eject keys that have expired. Background threads periodically sweep and delete expired keys to free up memory without waiting for a `get()` call to trigger the deletion.

## 6. Bottlenecks and Considerations: Cache Penetration & Stampede
*   **Cache Penetration:** A malicious user requests a `key` that doesn't exist in the Cache AND doesn't exist in the Database. It constantly misses the cache and blasts the DB.
    *   *Solution:* Store "Empty" or "Null" responses in the cache with a short TTL, or use a **Bloom Filter** in front of the DB.
*   **Cache Avalanche / Stampede:** A highly popular key (e.g., front-page news) expires. Simultaneously, 10,000 threads miss the cache and all 10,000 threads query the DB directly to recreate the object.
    *   *Solution:* Use a mutex lock. The first thread acquires the lock, goes to the DB, and sets the Cache. The other 9,999 threads wait a few milliseconds and read it successfully from the cache.

## 7. Follow-up Questions for Candidate
1.  Compare Redis and Memcached. When would you use one over the other? (Redis supports complex data types like Lists, Sets, Hashes, and persistence; Memcached is strictly pure strings and multi-threaded for simple raw performance).
2.  If the cache is updated, how do you keep the database in sync? (Write-Through vs Write-Around vs Write-Behind caching strategies).
