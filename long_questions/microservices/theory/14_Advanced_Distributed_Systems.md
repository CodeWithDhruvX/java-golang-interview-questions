# 🟣 **196–210: Advanced Distributed Systems**

### 196. What is the CAP theorem in practical terms?
"The CAP theorem states that a distributed system can only guarantee two out of three characteristics: Consistency (all nodes see the same data at the same time), Availability (every request receives a response, even if it's node failure), and Partition Tolerance (the system continues to operate despite network drops between nodes).

In reality, network partitions (P) are unavoidable in distributed systems. A router will eventually fail. Therefore, the real choice is always between Consistency (CP) and Availability (AP) during a network failure.

If a network drops between Node A and Node B, and a write goes to Node A:
- A **CP system** will refuse the write (sacrificing Availability) because it cannot sync with Node B to guarantee Consistency. Use case: Banking balances.
- An **AP system** will accept the write (sacrificing Consistency). Node B will serve stale data until the network heals and they sync. Use case: Social media likes."

#### Indepth
Modern systems try to circumvent strict CAP limits using 'PACELC'. It states that in case of Network Partition (P), choose Availability (A) or Consistency (C). Else (E) - meaning normal operation - choose Latency (L) or Consistency (C). Databases like Cassandra allow tuning this per-query via Quorums (Read/Write Consistency Levels).

---

### 197. What is Eventual Consistency vs Strong Consistency?
"**Strong Consistency** guarantees that once a write completes, any subsequent read will return the updated value. It requires synchronous replication and distributed locking. It's safe but slow and less available during outages. Relational databases default to this.

**Eventual Consistency** guarantees that if no new updates are made, all nodes will *eventually* return the last updated value. Writes are accepted quickly on one node and replicated asynchronously in the background. It's highly available and fast, but clients might temporarily read stale data.

In microservices, we heavily favor Eventual Consistency (via async messaging/Kafka) because forcing Strong Consistency across service boundaries (like using Two-Phase Commit) destroys both performance and availability."

#### Indepth
There are intermediate consistency models. 'Read-Your-Own-Writes' consistency guarantees that if a user updates their profile and refreshes the page, they see the new data, even if other users globally still see the old data for a few seconds. This is often achieved by routing a user's reads to the leader node for a short period after a write.

---

### 198. What is Vector Clock terminology and why is it used?
"In deeply distributed databases like DynamoDB or Cassandra, when there is a network partition, two different nodes might accept conflicting writes for the same piece of data (e.g., Node A updates a cart to 'Apple', Node B updates it to 'Banana'). 

When the network heals, how does the system know which update happened 'last'? Because server physical clocks are notoriously out of sync, relying on simple timestamps is dangerous.

A **Vector Clock** is an array of counters (logical clocks) kept per node, e.g., `[NodeA:2, NodeB:1]`. It tracks the *causality* of events. By comparing vector clocks, the system can mathematically determine if one event genuinely happened before another, or if they happened concurrently (a conflict). If it's a conflict, the system asks the client application to resolve it (e.g., merge 'Apple' and 'Banana' into one cart)."

#### Indepth
Causality means "Event A caused Event B". If the vector clocks prove concurrency rather than causality, systems fall back to Last-Write-Wins (LWW) resolution (which drops data based on inaccurate server timestamps), or they return both versions to the client applications (like Amazon's Dynamo shopping cart implementation) to do a semantic merge.

---

### 199. How do you implement Distributed Caching at scale?
"At small scale, a single Redis instance works. At massive scale, a single node will run out of memory or CPU bandwidth. To scale distributed caching, we use **Consistent Hashing**.

Instead of assigning cache keys to nodes using a simple modulo `Hash(key) % N` (which causes a massive reshuffle of 99% of keys if a node dies, leading to a thundering herd that crashes the database), Consistent Hashing maps both the servers and the cache keys onto a conceptual 'ring' of values.

A key is assigned to the first server it encounters moving clockwise on the ring. If a server dies, only the keys mapped to that specific server are remapped to the next server. 90% of the cache remains entirely intact, saving the database from an avalanche of cache misses."

#### Indepth
To prevent the remaining servers from being overwhelmed when one dies, Virtual Nodes (vnodes) are used. Each physical Redis server maps to 100 random spots on the ring. When a server dies, its load is distributed evenly across all other remaining servers, rather than dumping all the traffic onto a single 'next' neighbor.

---

### 200. What is a Distributed Lock and how do you implement it?
"When multiple instances of a microservice need to execute a critical piece of logic that must only happen once globally (like running a daily billing batch job, or altering a shared resource), they need a Distributed Lock. Local synchronized blocks in Java don't work across different Docker containers.

I implement distributed locks using **Redis** (specifically the Redlock algorithm) or **Zookeeper/Etcd**. 

In Redis, a service instance requests a lock by setting a unique key with an expiration time (`SET resource_name my_instance_id NX PX 30000`). `NX` means 'Only set if it doesn't exist'. If successful, this service holds the lock. Operations happen, and then it deletes the key to release the lock. The expiration (`PX 30000`) prevents deadlocks if the service crashes while holding the lock."

#### Indepth
Releasing a lock must be done via a Lua script in Redis to ensure atomicity. The script checks: "Is the value of this lock still my Instance ID? If yes, delete it." If you don't do this check, you might accidentally delete a lock that expired and was grabbed by a *different* instance, leading to massive data corruption.

---

### 201. What is the difference between Leader-Follower and Leaderless replication?
"**Leader-Follower (Master-Slave):** One node is the Leader. All writes MUST go to the Leader. The Leader then replicates the data to Followers. Reads can go to any node. It guarantees there are no write conflicts. *Databases: PostgreSQL, MySQL, MongoDB.*

**Leaderless (Peer-to-Peer):** There is no designated leader. A client can send a write to ANY node. That node coordinates replicating to the others. It provides massive write availability and performance, but allows conflicting writes which must be resolved. *Databases: Cassandra, DynamoDB.*"

#### Indepth
In Leaderless replication, consistency is managed via Quorums. W (Write nodes) + R (Read nodes) > N (Total replicas). If N=3, and I require W=2 and R=2, I am guaranteed strong consistency because any read quorum of 2 nodes must contain at least one node that received the latest write quorum of 2.

---

### 202. What is a Split-Brain scenario?
"Split-brain occurs in high-availability clusters when the network connection between active nodes fails, but the nodes themselves are still running. 

Because they can't see each other, Node A thinks Node B is dead, and Node B thinks Node A is dead. If left unchecked, both nodes might promote themselves to be the 'Primary / Leader' and start accepting writes independently. This leads to massive, unresolvable data corruption and completely un-syncable divergence.

To prevent this, clusters use **Quorum and Fencing mechanisms**. A node is only allowed to become the leader if it can get a majority vote (e.g., 3 out of 5 nodes). In a split, only one side of the network partition will have a majority. The minority side gracefully demotes itself or shuts down."

#### Indepth
Fencing involves completely cutting off the "dead" node from shared resources. "STONITH" (Shoot The Other Node In The Head) is a brutal but effective mechanism where the primary server literally sends a command to the networked power strip to turn off the power to the unresponsive secondary server, guaranteeing it cannot corrupt data.

---

### 203. How do you design an Idempotent API?
"An idempotent API is one where making the same exact request multiple times has the same effect as making it once. It is mandatory for distributed systems because network retries are inevitable. If a 'Charge Credit Card' API drops the response packet, the client *will* retry. If it's not idempotent, the user is charged twice.

To design this, the client generates a unique `Idempotency-Key` (a UUID) and sends it in the HTTP header with the POST request. 

The Server checks its database: 'Have I seen this Idempotency-Key before?'
- If No: Process the payment, save the result and the Key to the DB.
- If Yes: DO NOT process the payment. Just return the cached successful response from the database.

The client can retry 100 times safely."

#### Indepth
The database lookup and the business action must be wrapped in a transaction or handled carefully to prevent race conditions (two identical requests arriving at the exact same millisecond). A common approach is trying to insert the Idempotency-Key into a unique constraint column first; if it violates the constraint, the request is a duplicate.

---

### 204. What is the Thundering Herd problem and how do you mitigate it?
"The Thundering Herd problem occurs when a highly accessed piece of data in a cache expires (TTL ends). 

Suddenly, 5,000 concurrent requests for 'Dashboard Stats' check the cache, see a cache miss at the exact same millisecond, and all 5,000 requests query the underlying Database to calculate the stats simultaneously. The database's CPU spikes to 100% and it crashes.

I mitigate this using **Cache Stamping/Mutex Locks**. When the cache misses, the first thread to notice acquires a distributed lock for that specific key. Only that one thread is allowed to query the database and update the cache. The other 4,999 threads wait a few milliseconds, check the cache again (which is now populated by the first thread), and return gracefully."

#### Indepth
Another mitigation is "Probabilistic Early Expiration". Instead of letting the cache naturally expire, threads randomly check if the cache is *about* to expire (e.g., within the next 30 seconds). One lucky thread mathematically decides it's time to recalculate the value in the background asynchronously, ensuring the cache is never truly empty for the main traffic.

---

### 205. How do you handle clock drift in distributed systems?
"You handle it by assuming physical clocks are completely unreliable. In a distributed system, Server A might read 10:00:00 AM while Server B reads 10:00:05 AM. NTP (Network Time Protocol) syncs clocks, but always has milliseconds of drift.

If you rely on `System.currentTimeMillis()` to order events or resolve database write conflicts across different nodes, you will corrupt data. 

Therefore, you must use **Logical Clocks** (like Lamport timestamps or Vector Clocks). These don't measure 'time of day'; they measure 'sequence of events' (Event A happened before Event B). 

If you *must* use physical time for ordering (like Google Spanner does), you have to use specialized atomic clock hardware and GPS receivers in every data center (TrueTime API) to guarantee drift is under 7 milliseconds, and then mathematically wait out the uncertainty window during transactions."

#### Indepth
Snowflake IDs (pioneered by Twitter) are an excellent way to generate globally unique, roughly time-sorted IDs without centralized coordination. They embed a 41-bit timestamp, a 10-bit machine ID, and a 12-bit auto-increment sequence into a single 64-bit integer. Because the machine ID ensures uniqueness, slight clock drifts between machines don't cause ID collisions.
