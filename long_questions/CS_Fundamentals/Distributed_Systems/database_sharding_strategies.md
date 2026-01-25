# Database Sharding Strategies

## 1. What is Sharding?
Sharding is a method of splitting and storing a single logical dataset in multiple databases. It is a form of **Horizontal Scaling** (Scaling Out).
*   **Vertical Scaling (Scaling Up)**: Buying a bigger server (more RAM, CPU). Has a limit.
*   **Horizontal Scaling (Sharding)**: Adding more cheap servers. Unlimited scale.

## 2. Sharding Strategies

### A. Key Based Sharding (Hash Based)
*   **Algorithm**: `ShardID = Hash(Key) % NumberOfShards`
*   **Mechanism**: A hash function determines which shard holds the data.
*   **Pros**: Even distribution of data. Prevents "Hot Spots" if the hash function is uniform.
*   **Cons**:
    *   **Resharding is expensive**: If you add a new shard, the Modulo changes. Almost all keys need to be moved to different shards.
    *   **Solution**: **Consistent Hashing** (avoids massive data movement).

### B. Range Based Sharding
*   **Algorithm**: Data is split based on ranges of a key.
    *   Shard 1: UserIDs 1-100,000
    *   Shard 2: UserIDs 100,001-200,000
*   **Pros**: Easy to implement. Good for range queries (e.g., "Give me users created in Jan 2023").
*   **Cons**: **Hot Spots**. If Shard 1 holds "active" users and Shard 2 holds "old" users, Shard 1 will take all the load.

### C. Directory Based Sharding
*   **Algorithm**: A "Lookup Service" maintains a mapping table of `Key -> Shard ID`.
*   **Pros**: Extreme flexibility. You can move individual keys without changing algorithms.
*   **Cons**: **Single Point of Failure**. The Lookup Service becomes a bottleneck. Every query needs a lookup.

## 3. Challenges of Sharding

### 1. Joins across Shards
*   Running a SQL JOIN on data that sits on two different physical servers is impossible (or extremely slow).
*   **Solution**: Denormalization (duplicate data) or handling joins in Application Logic.

### 2. Distributed Transactions
*   How to ensure atomicity if you update Shard A and Shard B?
*   **Solution**: **Two-Phase Commit (2PC)**. Very slow. Avoid cross-shard transactions if possible.

### 3. Resharding / Rebalancing
*   Moving data while the system is live is complex and risky.

## 4. Interview Questions
1.  **When should you shard?**
    *   *Ans*: Only when you hit the limits of a single node (Wait until you have > 1TB data or huge write throughput). Don't shard prematurely.
2.  **How to handle unique IDs in a sharded DB?**
    *   *Ans*: You can't use auto-increment (both shards will generate ID 1). Use UUIDs (large) or **Snowflake IDs** (Twitter's approach: Timestamp + MachineID + Sequence).
