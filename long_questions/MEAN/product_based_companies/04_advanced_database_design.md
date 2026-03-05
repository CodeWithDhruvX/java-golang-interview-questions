# Advanced MongoDB & Database Design (Product-Based Companies)

Interviews at top product companies will test your ability to design databases for high read/write throughput, large datasets, complex queries, and data consistency.

## Database Architecture and Scaling

### 1. Explain Sharding in MongoDB. How do you choose a Shard Key?
Sharding is MongoDB's method for horizontal scaling across multiple servers. It distributes data across "shards" (replica sets).
*   **Shard Key**: A field (or fields) that determines which chunk of data goes to which shard. This is the most critical decision in a sharded cluster.
*   **Choosing a Good Shard Key**:
    *   **High Cardinality**: Many unique values so chunks can be small and distributed evenly (avoiding "jumbo chunks").
    *   **Non-Monotonically Increasing**: If you use a timestamp or `_id`, all new writes go to the *same* shard (the "hot shard"), defeating the purpose of scaling writes. Hashed shard keys are often used to distribute monotonically increasing values evenly.
    *   **Query Isolation**: Ideally, queries should include the shard key so the mongos router only queries *one* shard (targeted query) instead of broadcasting the query to *all* shards (scatter-gather).

### 2. What is a Replica Set? How does Election work during a failover?
A Replica Set is a group of mongod processes that maintain the same data set. It provides redundancy and high availability.
*   **Primary Node**: Receives all write operations. Replicates changes (oplog) to secondaries.
*   **Secondary Node(s)**: Asynchronously replicate the Primary's oplog and apply operations to their datasets. They can serve read queries if configured (Read Preference).
*   **Elections**: If the Primary becomes unavailable for >10 seconds, the remaining nodes hold an election.
    *   Nodes send heartbeats.
    *   An eligible Secondary calls for an election.
    *   Nodes vote based on priority and who has the most up-to-date oplog.
    *   The node with the majority of votes becomes the new Primary.

## Advanced Indexing and Querying

### 3. Explain how Indexing works in MongoDB and the different types of indexes.
Indexes are special data structures (B-trees) that store a small portion of the collection's data in an easy-to-traverse form. They dramatically improve read performance at the cost of slower writes and more RAM usage.
*   **Single Field**: Indexing one field (`{ "age": 1 }`).
*   **Compound Index**: Indexing multiple fields (`{ "age": 1, "username": -1 }`). Order matters tremendously here. If you query by `username` only, this index cannot be used.
*   **Multikey Index**: Indexing an array field.
*   **Text Index**: For full-text search.
*   **Geospatial Index**: For location-based queries ($near, $geoWithin).
*   **Sparse/Partial Indexes**: Only index documents that have a specific field or meet a filter condition, saving huge amounts of RAM.

### 4. What is the ESR (Equality, Sort, Range) Rule for complex queries?
When creating compound indexes to support a complex query, the order of the fields in the index follows the ESR rule:
1.  **Equality**: Fields that your query tests against exact values (`{ status: "active" }`) should come first.
2.  **Sort**: Fields used to sort the results should come next (`.sort({ createdAt: -1 })`).
3.  **Range**: Fields accessed using range operators (`$gt`, `$lt`, `$in`) should come last (`{ age: { $gt: 20 } }`).
Placing range fields before sort fields forces an expensive in-memory sort or "blocking sort".

## Advanced Operations and Consistency

### 5. Does MongoDB support ACID transactions? When should you use them in a Node.js API?
Historically, MongoDB only offered ACID properties at the *single-document* level.
*   Since v4.0, MongoDB supports **Multi-Document ACID Transactions** across a replica set (and since v4.2, across a sharded cluster).
*   **When to use**: *Sparingly*. Transactions severely limit horizontal scalability and performance due to locking overhead.
*   Always try to model data so that updates affect only one document (using embedded sub-documents) to avoid transactions. Only use them when absolutely necessary (e.g., transferring money between two separate User Account documents in a banking app).

### 6. Explain the `$lookup` aggregation stage. What are the performance implications?
`$lookup` is MongoDB's equivalent of an SQL `LEFT OUTER JOIN`. It brings data from another collection into the current pipeline.
*   **Performance Implications**: `$lookup` can be extremely slow on large datasets because it is essentially a nested loop. It compares every document in the local collection against the foreign collection.
*   **Optimization**: Ensure the `foreignField` (the target field in the other collection) is **indexed**. Without an index on the foreign field, `$lookup` performs a full collection scan for every document in the source pipeline. Consider data denormalization (embedding) instead of `$lookup` for frequently accessed, highly related data.
