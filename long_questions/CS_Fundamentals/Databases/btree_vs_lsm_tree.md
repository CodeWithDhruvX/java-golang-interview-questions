# B-Tree vs LSM Tree

## 1. Overview
The two most common data structures used for the storage engine of a database.
*   **B-Tree (Balance Tree)**: Optimized for **Reads**. (PostgreSQL, MySQL/InnoDB).
*   **LSM Tree (Log-Structured Merge Tree)**: Optimized for **Writes**. (Cassandra, LevelDB, RocksDB, Kafka).

## 2. B-Tree (Read Optimized)
Standard standard tree structure where every node is a disk page (usually 4KB).

### How it works
*   **In-Place Updates**: To modify data, we find the page on disk, load it, modify it, and write it back.
*   **Structure**: Wide and short (High branching factor).
*   **Pros**:
    *   **Fast Reads**: O(log N). Can find any key with very few disk seeks.
    *   **Consistent Performance**: Read/Write times are stable.
*   **Cons**:
    *   **Slow Random Writes**: Inserting a random key requires loading a random page from disk (Disk Seek), potentially causing "Write Amplification" if the page splits.

## 3. LSM Tree (Write Optimized)
Append-only structure.

### How it works
1.  **MemTable**: Writes go to an in-memory buffer (Red-Black Tree or Skip List). Very fast (O(1) disk I/O).
2.  **WAL (Write Ahead Log)**: Writes also appended to a sequential log file for durability.
3.  **SSTable (Sorted String Table)**: When MemTable is full, it is flushed to disk as an immutable sorted file (SSTable).
4.  **Compaction**: In background, the system merges multiple SSTables, discards old/deleted keys, and creates a larger sorted file.

### Pros
*   **Extremely Fast Writes**: Every write is just an append. No random disk seeks.
*   **Compression**: Immutable files are easy to compress.

### Cons
*   **Slower Reads**: To find a key, you check MemTable -> SSTable 1 -> SSTable 2 -> ... -> SSTable N.
    *   **Mitigation**: **Bloom Filters** are used to quickly check if a key exists in an SSTable before scanning it.
*   **Compaction Overhead**: Background compaction takes CPU/IO.

## 4. Comparison Table

| Feature | B-Tree | LSM Tree |
| :--- | :--- | :--- |
| **Primary Use** | RDBMS (SQL) | NoSQL / Big Data |
| **Write Pattern** | Random I/O (Update in place) | Sequential I/O (Append only) |
| **Read Speed** | Very Fast | Slower (requires checking multiple files) |
| **Write Speed** | Slower (Disk Seeks) | Very Fast |
| **Space** | Fragmentation issues | Compacted (good compression) |
| **Examples** | MySQL, Postgres, Oracle | Cassandra, RocksDB, HBase |

## 5. Interview Questions
1.  **Why is LSM faster for writes?**
    *   *Ans*: It turns random writes into sequential writes. Hard Drives (HDD) and SSDs are significantly faster at sequential writes than random writes.
2.  **How do we make LSM reads faster?**
    *   *Ans*: Bloom Filters. A probabilistic data structure that tells us "Key is definitely NOT in this file" or "Key MIGHT be in this file". Saves unnecessary disk checks.
