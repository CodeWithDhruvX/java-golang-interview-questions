# Apache ZooKeeper Interview Questions

## 1. Core Concepts & Basics

### Q1. What is Apache ZooKeeper and why is it used in distributed systems?
**Answer:**
Apache ZooKeeper is a centralized service for maintaining configuration information, naming, providing distributed synchronization, and offering group services. In distributed systems, managing coordination between nodes is complex and prone to race conditions and deadlocks. ZooKeeper abstracts this complexity and provides a simple set of primitives that applications can use to build higher-level synchronization features (like distributed locks, leader election, and queueing).

Key use cases include:
- **Configuration Management**: Storing and broadcasting configurations to all nodes.
- **Naming Service**: Mapping logical names to physical addresses.
- **Distributed Synchronization**: Locks, barriers, and leader election.
- **Cluster Management**: Node join/leave notifications and health monitoring.

### Q2. Explain the typical architecture of a ZooKeeper Ensemble.
**Answer:**
A ZooKeeper service is typically replicated over a set of hosts called an **Ensemble**. The architecture follows a client-server model where servers form a cluster.
- **Leader**: One node is elected as the leader. All write requests from clients are forwarded to the leader. The leader ensures that writes are serialized and replicated to followers.
- **Followers**: The other nodes are followers. They serve read requests locally to clients. If they receive a write request, they forward it to the leader.
- **Observer**: A non-voting follower. They hear the results of votes and serve read requests, used to scale read performance without impacting the quorum of write consensus.

### Q3. What are znodes? Explain the different types of znodes.
**Answer:**
ZooKeeper organizes its data in a hierarchical namespace, similar to a standard file system. Every node in this namespace is called a **znode**. Znodes store data and have metadata (version numbers, ACLs, timestamps).

Types of znodes:
1.  **Persistent znode**: Remains in ZooKeeper until explicitly deleted. Useful for storing configuration data.
2.  **Ephemeral znode**: Exists only as long as the session of the client that created it is active. If the client disconnects or crashes, the znode is automatically deleted. Useful for tracking active nodes in a cluster.
3.  **Sequential znode**: When created, ZooKeeper appends a monotonically increasing counter to the znode's name. Can be persistent or ephemeral.
4.  **Container znode** (Added in 3.5+): Special znodes used for leader election or locking. When the last child of a container znode is deleted, the container becomes a candidate for deletion by the server.
5.  **TTL znode** (Added in 3.5+): A persistent znode with a given Time-To-Live. If the znode has no children and hasn't been modified within the TTL, it is deleted.

## 2. Data Consistency & Protocols

### Q4. Describe the ZAB (ZooKeeper Atomic Broadcast) protocol.
**Answer:**
ZAB is a crash-recovery atomic broadcast protocol specific to ZooKeeper. It guarantees that updates are delivered in the same order they were sent. It works in two main phases:
1.  **Leader Election & Recovery (Discovery Phase)**: When the service starts or a leader fails, servers enter a state to elect a new leader. The nodes synchronize their state with the new leader to ensure everyone has the most up-to-date history.
2.  **Atomic Broadcast (Broadcast Phase)**: Once the leader is synchronized with a quorum of followers, it starts processing write requests.
    - Leader creates a proposal for a write request.
    - Sends proposal to followers.
    - When a quorum of followers acknowledges the proposal, the leader sends a COMMIT message.
    - The write is then applied to the state machine database.

### Q5. What is a "Quorum" in ZooKeeper and why is an odd number of nodes recommended?
**Answer:**
A Quorum is the minimum number of nodes that must be operational and agree on a transaction for it to be committed. In ZooKeeper, a quorum is a strict majority of nodes in the ensemble: `(N/2) + 1`.

An ensemble is typically deployed with an odd number of nodes (3, 5, 7) because:
- A 3-node cluster needs 2 nodes to form a quorum, so it can tolerate 1 failure.
- A 4-node cluster needs 3 nodes to form a quorum, so it can *also* only tolerate 1 failure. There is no availability advantage to a 4-node cluster over a 3-node cluster, but it adds network overhead.
- An odd number strictly prevents **Split-Brain** scenarios, ensuring only one partition can ever hold a majority.

### Q6. What consistency guarantees does ZooKeeper provide?
**Answer:**
ZooKeeper does *not* provide strict linearizable consistency for reads across all nodes. It provides:
1.  **Sequential Consistency**: Updates from a client will be applied in the order that they were sent.
2.  **Atomicity**: Updates either succeed or fail entirely.
3.  **Single System Image**: A client will see the same view of the service regardless of the server it connects to.
4.  **Reliability**: Once an update is applied, it persists until a client overwrites it.
5.  **Timeliness**: The clients' view of the system is guaranteed to be up-to-date within a certain time bound.

*(Note: Reads are eventually consistent because you might read from a follower that hasn't yet processed a recent commit. Clients can use the `sync` command to force a follower to catch up with the leader before reading).*

## 3. Advanced Features & Implementation

### Q7. Explain the concept of "Watches" in ZooKeeper.
**Answer:**
A Watch is a one-time trigger that a client can set on a znode. When the data or state of the znode changes (data changed, child created/deleted), ZooKeeper sends a notification to the client.

Important characteristics:
- **One-time trigger**: Once a watch is triggered, it is removed. To get future notifications, the client must set a new watch.
- **Sent to the client**: Watch events are asynchronously sent to the client.
- **Ordering guarantee**: A client is guaranteed to see a watch event *before* it sees the new data that triggered the watch.

### Q8. How do you implement Leader Election using ZooKeeper?
**Answer:**
Leader election can be implemented using **Ephemeral Sequential Znodes**:
1. All client nodes participating in the election create a znode with the `EPHEMERAL_SEQUENTIAL` flag under a common parent znode, e.g., `/election/node-`
2. ZooKeeper appends a sequence number to the name (e.g., `node-00001`, `node-00002`).
3. Each client fetches the list of children under `/election/`.
4. The client whose znode has the **lowest** sequence number becomes the Leader.
5. All other nodes become Followers.
6. To detect leader failure: Each follower sets a watch on the znode that has the sequence number immediately *preceding* its own. (e.g., node 3 watches node 2). This avoids the "herd effect" where all nodes watch the leader and get notified simultaneously upon failure. When a node's watched znode disappears, it re-checks the list of children holding the lock.

### Q9. How do you implement a Distributed Lock in ZooKeeper?
**Answer:**
The logic is almost identical to Leader Election:
1. Create an `EPHEMERAL_SEQUENTIAL` znode under a `/locks/my-resource/` parent.
2. Get all children of the parent.
3. If the created znode has the lowest sequence number, the client has acquired the lock.
4. If not, the client sets a watch on the znode with the next lowest sequence number.
5. When notified of deletion, the client checks again if it holds the lowest sequence number.

Because they are ephemeral, if a client holding the lock crashes, the znode is deleted, and the next node in line is notified to acquire the lock.

## 4. Troubleshooting & Ecosystem

### Q10. How does ZooKeeper prevent the Split-Brain problem?
**Answer:**
Split-Brain occurs when a cluster partition splits into two isolated subgroups, and each subgroup thinks it should be the active cluster, leading to data corruption.

ZooKeeper prevents this via the **Quorum** requirement. To successfully elect a leader or commit a write transaction, a strict majority (e.g., 3 out of 5 nodes) must agree. If a network partition splits a 5-node cluster into a group of 3 and a group of 2, only the group of 3 can form a quorum. The group of 2 will halt operations (stop taking writes and block reads) until connectivity is restored, completely eliminating the possibility of two separate subgroups acting as leaders.

### Q11. What is a zxid?
**Answer:**
`zxid` stands for ZooKeeper Transaction ID. Every change to the ZooKeeper state receives a unique `zxid` formatted as a 64-bit number.

It consists of two parts:
- **Epoch (32 bits)**: Identifies the leader tenure. Incremented every time a new leader is elected.
- **Counter (32 bits)**: Incremented for each unique transaction processed by the current leader.

`zxid` exposes the total ordering of all changes. If `zxid(A) < zxid(B)`, then transaction A happened before transaction B.

### Q12. Explain the relationship between Apache Kafka and ZooKeeper. Why is Kafka deprecating it (KRaft)?
**Answer:**
Historically, Kafka heavily relied on ZooKeeper to:
- Elect the Controller node.
- Store cluster topology (which brokers are alive).
- Topic configuration and partition assignments.
- Track in-sync replica (ISR) lists.

**Why Kafka moved to KRaft (Kafka Raft Metadata mode):**
1. **Scalability Limitations**: ZooKeeper became a bottleneck when scaling to millions of partitions because metadata operations took too long.
2. **Operational Complexity**: Administrators had to manage, secure, and monitor two separate distributed systems (Kafka and ZooKeeper).
3. **Recovery Time**: When a Kafka controller failed, fetching all metadata from ZooKeeper to elect a new one could take significant time, causing partition unavailability.

KRaft integrates the metadata quorum directly into Kafka using a Raft-based consensus protocol, removing the ZooKeeper dependency entirely as of Kafka 3.3+.
