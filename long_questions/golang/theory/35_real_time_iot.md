# 🟢 Go Theory Questions: 681–700 Real-Time Systems, IoT, and Edge Computing

## 681. How do you build a real-time chat server in Go?

**Answer:**
We use **WebSockets** (`gorilla/websocket` or `nhooyr.io/websocket`).
Architecture:
1.  **Hub**: Maintains a map of `Register`, `Unregister`, and `Broadcast` channels.
2.  **Client**: Each connection spawns 2 goroutines: `ReadPump` (from WS to Hub) and `WritePump` (from Hub to WS).
The Hub acts as the central router. When a message comes in, it loops over registered clients and pushes to their `send` channel.

---

## 682. How do you implement WebSockets in Go?

**Answer:**
Standard HTTP Handlers.
```go
func wsHandler(w, r) {
    conn, _ := upgrader.Upgrade(w, r, nil) // Hijacks the TCP connection
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil { break }
        // handle msg
    }
}
```
We must handle the Ping/Pong control frames to keep the connection alive through load balancers (which often have 60s timeouts).

---

## 683. How do you ensure order of events in real-time systems?

**Answer:**
Time is relative. We use **Sequence Numbers** or **Lamport Timestamps**.
Server assigns a monotonic ID (101, 102) to events.
Client receives 101 and 103. It knows 102 is missing. It buffers 103 and waits for 102 (or requests retransmit).
TCP guarantees order on the wire, but application logic (retries, reconnects) can rearrange things.

---

## 684. How do you handle high concurrency in WebSocket servers?

**Answer:**
(See Q 346).
Go handles 1M websockets easily (1M goroutines = 2-4GB RAM).
**Optimization**: Epoll / Kqueue (using `gnet` or `gobwas/ws`).
Instead of 1 goroutine per connection (blocking Read), we use the OS Event Loop (Netpoller) to notify us which connection has data. This saves the stack memory overhead for idle connections.

---

## 685. How do you implement presence tracking in Go (like online users)?

**Answer:**
**Heartbeats** + **Redis**.
Client sends "I'm here" every 30s.
Server: `SET user:123:online "1" EX 45`.
If heartbeat stops, key expires.
To count online: `KEYS user:*:online` (Scan).
For clusters, Redis Pub/Sub distributes the "User Connected" event so all nodes know regarding the global state.

---

## 686. How do you reduce latency in real-time systems?

**Answer:**
1.  **Protocol**: Use binary (Protobuf) over Text (JSON) to reduce payload size.
2.  **Serialization**: Use fast libraries (`msgpack`, `gogo/protobuf`).
3.  **NoGC**: Reuse buffers (`sync.Pool`) to avoid GC pauses.
4.  **Edge**: Deploy Go servers geographically closer to users.

---

## 687. How do you build a real-time dashboard backend in Go?

**Answer:**
We don't query DB every second.
We use **Push**.
ETL pipeline (Kafka) -> Aggregator (Go) -> Redis (Current Stats).
Go Web Server subscribes to Redis/Kafka.
When stats change, Go pushes the delta to the frontend via Server-Sent Events (SSE) or WebSockets. SSE is preferred for dashboards (Uni-directional, easier than WS).

---

## 688. How do you handle message fan-out for WebSocket clients?

**Answer:**
Direct loop `for client := range clients { client.Write() }` is dangerous. One slow client blocks the loop.
**Solution**: Buffered Channels for each client.
The loop just pushes to channel (Non-blocking).
Each client has its own writer goroutine draining that channel.
If the channel fills, we disconnect the client ("Slow Consumer").

---

## 689. How do you design a real-time bidding system in Go?

**Answer:**
Latency Budget: 100ms.
1.  **In-Memory**: Everything must be in RAM (Structs/Maps). DB is too slow.
2.  **Concurrency**: Hash the AuctionID to a Shard (Channel).
3.  **Serializer**: Single-thread per auction ensures no race conditions on the "Highest Bid".
4.  Persistence: Asynchronous write-behind (WAL) to disk, but the read/bid path never touches IO.

---

## 690. How do you throttle real-time updates?

**Answer:**
**Debounce/Conflation**.
If a stock price changes 100 times a second, human eyes can't see it.
We send update max once per 100ms.
Go:
```go
select {
case val := <-updates:
    current = val // Overwrite, don't send yet
case <-ticker.C:
    send(current) // Send latest snapshot
}
```
This conflates intermediate updates into a single network packet.

---

## 691. How do you buffer real-time data safely?

**Answer:**
**Ring Buffer** (Circular Queue).
Fixed size array.
Write pointer wraps around.
If write catches read (full), we either Drop Oldest (Metrics) or Drop Newest (commands).
This avoids uncontrolled memory growth (Slice append) during traffic spikes.

---

## 692. How do you build a publish-subscribe engine for WebSockets?

**Answer:**
We match Topics.
`Map[Topic] -> Map[ConnID] -> Conn`.
When a client subscribes "chat:room1", we add them to the "chat:room1" bucket.
When a message arrives for "chat:room1", we only iterate that bucket.
We use `sync.RWMutex` to protect the map, as subscriptions/unsubscriptions happen concurrently with broadcasting.

---

## 693. How do you sync real-time state between browser and Go backend?

**Answer:**
**Operational Transformation (OT)** or **CRDTs** (Conflict-free Replicated Data Types).
Simpler: **Version Vectors**.
Server sends `State v5`.
Client sends `Update based on v5`.
If Server is now `v6`, it detects conflict, rejects update, sends `v6` to Client. Client rebases its changes on top of `v6`.

---

## 694. How do you implement real-time location tracking?

**Answer:**
**Geospatial Index** (QuadTree or S2).
Redis `GEOADD keys lat lon`.
Redis `GEORADIUS key ...` to find drivers nearby.
Go acts as the ingest pipe. It batches updates (Driver moves every 1s), updates Redis, and queries for nearby users to push "Car is coming" updates.

---

## 695. How do you use Go in resource-constrained IoT devices?

**Answer:**
**TinyGo**.
It compiles Go to very small binaries (flashable to Arduino, ESP32, WASM).
It relies on LLVM.
We avoid `reflection`, `fmt` (too heavy), and heavy GC usage.
We use static buffers. This allows writing safe, concurrent driver code for sensors using goroutines (simulated on single-core microcontrollers).

---

## 696. How do you collect telemetry data from IoT devices?

**Answer:**
Protocol: **MQTT** (Lightweight, Pub/Sub).
Go Broker or Client `paho.mqtt.golang`.
Devices publish to `sensors/temp`.
Go service subscribes, batches 1000 readings, writes to **Time Series DB** (InfluxDB/Prometheus).
Go's efficiency is vital here to handle 100k connected IoT devices on a small server.

---

## 697. How do you handle network partitions in distributed IoT systems?

**Answer:**
IoT devices are often offline.
Strategy: **Store and Forward**.
Device stores data locally (Flash/SQLite).
When Go server accepts connection, Device sends backlog.
Go server must handle out-of-order timestamps (trust the device timestamp, not the receive timestamp).

---

## 698. How do you implement OTA (Over-The-Air) updates in Go?

**Answer:**
Go binaries are self-contained.
1.  Device polls Server: "Check Update".
2.  Server says "New Ver Available (URL, Hash)".
3.  Device downloads binary to separate partition.
4.  Verifies SHA256 / Signature (Security critical).
5.  Reboots into new partition.
We use libraries like `go-update` to handle the binary replacement (on Linux/Gateway devices).

---

## 699. What are the security challenges in Go IoT apps?

**Answer:**
1.  **mTLS**: Go `tls.Config` to require client certs. Each device has a unique cert.
2.  **Hardcoded Credentials**: Avoid. Use TPM (Trusted Platform Module) access via Go syscalls.
3.  **Binary Stripping**: Strip symbol tables to make reverse engineering harder.

---

## 700. How do you use Go for edge computing (AWS Lambda @ Edge / Cloudflare)?

**Answer:**
We compile to **WASM** (Cloudflare Workers) or standard **Linux Binary** (AWS Lambda).
Latency is king.
We minimize `init()` logic.
We use pure Go dependencies (avoid CGO).
We design for "Shared Nothing" architecture, assuming the instance can disappear immediately after the request.
