# Redis Architecture

## 1. Why is Redis so fast?
Redis is an **In-Memory** Key-Value store. But the key to its speed is that it is **Single Threaded**.
*   **No Context Switching**: CPU doesn't waste time switching threads.
*   **No Locks**: No need for complex Mutexes on shared data (because only 1 thread accesses it).
*   **I/O Multiplexing**: Uses `epoll` (Linux) or `kqueue` (BSD) to handle thousands of concurrent connections in a non-blocking way.

## 2. The Event Loop
1.  Redis waits for a socket to become readable (Client sends command).
2.  Reads the command.
3.  Parses the command.
4.  Executes the command on the in-memory data.
5.  Sends the reply to the socket.
6.  Repeats.
*Note*: Since it's single-threaded, one slow command (like `KEYS *`) blocks the entire server for everyone.

## 3. Persistence (Durability)
Since RAM is volatile, Redis offers two ways to save data to disk.

### A. RDB (Redis Database File) - Snapshots
*   **Mechanism**: Saves a point-in-time snapshot of the dataset at specified intervals (e.g., every 5 mins).
*   **Pros**: Compact file (fast recovery/backup). Minimal impact on performance (forks a child process).
*   **Cons**: **Data Loss**. If Redis crashes, you lose the last 5 mins of data.

### B. AOF (Append Only File) - Logs
*   **Mechanism**: Logs every write operation (SET, INCR) received by the server.
*   **Sync Policies**:
    *   `always`: Slowest, safest.
    *   `everysec`: Default. Logs every second.
*   **Pros**: **High Durability**. Max 1 second data loss.
*   **Cons**: Larger file size. Slower recovery (needs to replay all commands).

## 4. Replication & High Availability
*   **Master-Slave**: Master handles Writes. Slaves handle Reads. Async replication.
*   **Sentinel**: A monitoring system. If Master dies, Sentinel promotes a Slave to Master (Automatic Failover).
*   **Redis Cluster**: Sharded solution. Splits data across multiple nodes (16384 hash slots).

## 5. Interview Questions
1.  **Why is Redis Single Threaded?**
    *   *Ans*: Because CPU is not the bottleneck (RAM and Network are). Single-threading makes the code simpler (no deadlocks, no race conditions) and avoids context-switch overhead.
2.  **How to delete millions of keys without blocking?**
    *   *Ans*: Do NOT use `DEL`. Use `UNLINK` (non-blocking delete). It removes the key from keyspace in O(1) and reclaims memory in a background thread.
3.  **Memcached vs Redis?**
    *   *Ans*:
        *   **Memcached**: Simple Key-Value string store. Multi-threaded. No persistence.
        *   **Redis**: Advanced structures (Lists, Sets, Sorted Sets). Persistence. Replication.
