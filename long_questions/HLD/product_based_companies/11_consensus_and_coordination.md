# High-Level Design (HLD): Consensus and Coordination

In heavily distributed environments, ensuring multiple independent nodes agree on a single source of truth—even when some nodes fail or networks partition—is a fundamental problem. 

## 1. What is the Split-Brain Problem?
**Answer:**
In a High Availability (HA) cluster with an Active node and a Standby node, they usually send heartbeat messages to each other. 
*   **The Problem:** If the network link between the two nodes breaks, but both nodes are still running, the Standby node assumes the Active node has died. It promotes itself to Active. Now you have TWO Active nodes (a "split brain"). Both accept writes, leading to catastrophic data corruption and conflicts.
*   **The Solution:** You need a "Quorum". Instead of 2 nodes, use 3 nodes (or 5, or 7). To make a decision or write data, a strict majority (e.g., 2 out of 3) must agree. If a network splits the nodes into a group of 2 and a group of 1, only the group of 2 (the quorum) can continue operating. The isolated node pauses itself.

## 2. What are Paxos and Raft?
**Answer:**
They are **Distributed Consensus Algorithms**. Their purpose is to allow a cluster of machines to agree on a sequence of values (like a transaction log) and guarantee that everyone has the same data in the same order, even if servers crash.
*   **Paxos:** The original, mathematically rigorous consensus algorithm. It is notoriously difficult to understand and implement correctly in software.
*   **Raft:** Designed explicitly to be understandable. It decomposes consensus into independent sub-problems: Leader Election, Log Replication, and Safety.
    *   **Leader Election:** Nodes start as Followers. If they don't hear a heartbeat, one becomes a Candidate and requests votes. With a majority vote, it becomes the Leader.
    *   **Log Replication:** All client write requests go *only* to the Leader. The Leader appends the command to its log, sends the append command to followers, waits for a majority to acknowledge it, commits the entry, and then executes it.

## 3. What is Apache Zookeeper and etcd? Why do we need them?
**Answer:**
These are distributed, highly reliable key-value stores used specifically for cluster coordination and configuration management. They implement consensus algorithms internally (Zookeeper uses ZAB - Zookeeper Atomic Broadcast, etcd uses Raft).
*   **Why use them?** You rarely write Raft from scratch. Instead, you use Zookeeper/etcd.
*   **Use Cases:**
    *   **Leader Election for your Microservices:** If you have 5 instances of a worker process but only want 1 to actively process a cron job, they can all try to create an ephemeral node in Zookeeper. The first one succeeds and becomes the leader. If that instance crashes, Zookeeper deletes the ephemeral node, and the remaining 4 compete to be the new leader.
    *   **Service Discovery:** Microservices register their IPs in etcd. Other services query etcd to find where to send traffic. (Used heavily in Kubernetes).
    *   **Distributed Locking:** Similar to leader election, used to ensure safe access to a shared resource across a cluster.

## 4. What is the Gossip Protocol?
**Answer:**
Consensus algorithms like Raft require strong consistency and become bottlenecks if you have thousands of nodes. The Gossip protocol is an eventually consistent, peer-to-peer communication mechanism for massive clusters.
*   **How it works:** Think of a rumor spreading in a crowd. Node A randomly selects a few peers (B and C) and sends them its state. B and C then randomly select distinct peers and forward the state. The information spreads exponentially fast across the cluster.
*   **Use Cases:** 
    *   **Cassandra and DynamoDB:** Use gossip to discover cluster membership, detect node failures, and figure out which nodes hold which segment of the partition ring.
    *   **Cryptocurrency Networks:** Bitcoin nodes use gossip to broadcast transactions and new blocks to the global network.
