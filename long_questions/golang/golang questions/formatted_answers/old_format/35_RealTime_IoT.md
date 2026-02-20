# 🔴 **681–700: Real-Time Systems, IoT, and Edge Computing**

### 681. How do you build a real-time chat server in Go?
"I use **Gorilla WebSocket** or **Melody**.
I maintain a `Client` struct per connection.
I have a `Hub` that manages `register`, `unregister`, and `broadcast` channels.
When a user types, the handler sends to `broadcast`. The Hub loop iterates all active clients and writes the message down their websocket connections. Go handles 10k concurrent chats easily."

#### Indepth
**Horizontal Scaling**. The Hub is local to one server. If User A is on Server 1 and User B is on Server 2, they can't chat. You MUST use a **Pub/Sub backend** (Redis/NATS). Server 1 publishes `chat_msg` to Redis. Server 2 subscribes and pushes it to User B's websocket.

---

### 682. How do you implement WebSockets in Go?
"Standard library doesn't support it directly.
`upgrader := websocket.Upgrader{}`.
HTTP Handler:
`conn, _ := upgrader.Upgrade(w, r, nil)`.
Now I have a TCP-like connection.
`for { _, msg, _ := conn.ReadMessage(); handle(msg) }`.
I must handle pings/pongs to keep the connection alive through load balancers."

#### Indepth
**Compression**. Text-based JSON is bloated. `gorilla/websocket` supports Per-Message Compression Extensions (PMCE). Enabling `EnableCompression: true` can reduce bandwidth usage by 70% for JSON data, at the cost of slight CPU overhead for zipping/unzipping frames.

---

### 683. How do you ensure order of events in real-time systems?
"I use a **Sequence Number**.
Server appends `seq: 1`, `seq: 2`.
Client receives `1`, then `3`.
Client knows `2` is missing, so it buffers `3` and waits (or requests retransmission) for `2`.
TCP guarantees order on the wire, but if I utilize multiple backend servers or reconnections, app-level sequencing is mandatory."

#### Indepth
**Vector Clocks**. In distributed systems with no central counter, "Sequence 1, 2, 3" is hard. Vector Clocks (`{NodeA: 1, NodeB: 5}`) help detect causal relationships and merge conflicts. If strict total ordering is required, you need a centralized Sequencer (like Apache Kafka) which is a single point of serialization.

---

### 684. How do you handle high concurrency in WebSocket servers?
"Standard Go (goroutine per conn) works up to ~100k.
For 1M+, I use **library specific optimizations** (`gnet` / `gobwas/ws`) to minimize memory per connection (Epoll).
I avoid storing heavy state in the connection struct.
I optimize buffer sizes (don't alloc 4KB buffer for every idle client)."

#### Indepth
`gnet` / `nbio` use an **Event Loop** (Non-blocking I/O) instead of Goroutine-per-connection. This mimics Node.js/Netty. It allows handling 1M+ idle connections with very little RAM. However, your business logic *must not block* the loop, or you freeze the entire server.

---

### 685. How do you implement presence tracking in Go (like online users)?
"I need a shared store: Redis.
User connects: `SET user:123:online 1 EX 60` (Heartbeat every 30s).
User disconnects: `DEL user:123:online`.
To count: Scan keys (slow) or use HyperLogLog.
To show friends: `MGET user:A:online user:B:online`.
The WebSocket server just pings Redis; it doesn't hold the global truth."

#### Indepth
**Bitmaps**. If user IDs are integers, Redis Bitmaps are ultra-efficient. `SETBIT online_users 123 1`. To count online users: `BITCOUNT online_users`. This takes 1 bit per user. 1 million users = ~125KB of RAM. This is the fastest way to track "Who is online" at scale.

---

### 686. How do you reduce latency in real-time systems?
"1.  **Protocol**: Use generic WebSocket or QUIC (HTTP/3) to avoid handshake overhead.
2.  **Serialization**: Use Protobuf, not JSON.
3.  **Geo-Distribution**: Run Go edge servers near the user (Fly.io).
4.  **No GC**: Optimize tight loops to avoid GC pauses buffering packets."

#### Indepth
**Zero-Copy Networking**. The standard `io.Copy(conn, file)` copies data from Kernel -> User Space -> Kernel. Use `syscall.Sendfile` (or `io.Copy` which optimizes for it) to copy directly from Disk Cache -> Network Card, bypassing Go's memory entirely for static file serving.

---

### 687. How do you build a real-time dashboard backend in Go?
"I use **Server-Sent Events (SSE)**.
It's simpler than WebSockets (uni-directional).
`w.Header().Set("Content-Type", "text/event-stream")`.
`for { data := <-updates; fmt.Fprintf(w, "data: %s\n\n", data); w.Flush() }`.
The browser automatically reconnects if dropped. Perfect for stock tickers."

#### Indepth
**HTTP/2**. Legacy SSE used 1 TCP connection per tab. Multi-tab users hit the browser limit (6 connections/domain). HTTP/2 multiplexes all SSE streams over a single TCP connection. Ensure your Go server supports HTTP/2 (`http.Server` does by default with TLS) to fix the connection limit issue.

---

### 688. How do you handle message fan-out for WebSocket clients?
"If I have 10k users in a 'Lobby', a single `for` loop to write to 10k connections takes too long.
I **Shard** the hub.
10 sub-hubs, each managing 1k users.
I broadcast in parallel (10 goroutines).
Or I use a tiered architecture: Backend -> Nats -> Edge Nodes -> Users."

#### Indepth
**Tree Broadcast**. Instead of 1 Loop sending to 10k users, split it. Root sends to 10 Workers. Each Worker sends to 1k users. This parallelizes the syscalls (`write`). For global scale, propagate message to regional Edge Servers first, then fan-out locally to users in that region.

---

### 689. How do you design a real-time bidding system in Go?
"Latency is critical (< 100ms).
I keep the auction state **In-Memory** (Go struct), sharded by AuctionID.
I use a mutex per auction.
Requests come in, lock, update bid, unlock.
Persistence is async (write to WAL/Kafka).
I strictly avoid DB queries on the bidding path."

#### Indepth
**Lock Contention**. If 1000 bids arrive for Item X, `mutex.Lock()` becomes the bottleneck. Use **Atomic Instructions** (`atomic.CompareAndSwapInt64`) specifically for the "Current Price" field. This is lock-free and much faster. Only lock the full struct when settling the auction.

---

### 690. How do you throttle real-time updates?
"I use **Conflation**.
If a stock price changes 100 times/sec, the human eye handles 60fps.
I overwrite the pending update in the buffer.
`select { case out <- msg: (sent) default: (buffer full, drop old, insert new) }`.
The client only gets the *latest* state when it's ready to read, skipping intermediate values."

#### Indepth
**Debounce vs Throttle**. Throttle = "At most 1 update per 100ms" (Stock tick). Debounce = "Wait until silence for 100ms, then send" (Search autocomplete). In real-time data, Throttling/Conflation is usually what you want to prevent overwhelming the client's CPU.

---

### 691. How do you buffer real-time data safely?
"I use a **Ring Buffer**.
Fixed size.
It prevents OOM if the consumer is slow.
If the ring fills, I overwrite the oldest data (for metric streams) or close the connection (for critical streams). Infinite buffering is the root cause of server crashes."

#### Indepth
**Ring Implementation**. A simple array `buf [1024]T` plus two integers `head` and `tail`. Avoid slices (resizing allocations). Calculate position with `seq % 1024`. This is CPU cache-friendly and generates zero garbage, essential for high-throughput packet processing.

---

### 692. How do you build a publish-subscribe engine for WebSockets?
"I map `Topic -> Set<Connection>`.
`Subscribe("sports", conn)`.
When `Publish` happens, I iterate the Set.
Challenge: locking. `RWMutex` on the map.
For distributed support, I back this with Redis Pub/Sub: receiving a msg from Redis triggers the local fan-out to WebSockets."

#### Indepth
**Wildcard Matching**. Triemap (Prefix Tree) is efficient for `sports.*` style subscriptions. When `sports.tennis` is published, walk the tree. If node `sports` has children `*`, include them. This is faster than iterating all topics and running `regex.Match`.

---

### 693. How do you sync real-time state between browser and Go backend?
"**Operational Transforms (OT)** or **CRDTs** (Conflict-free Replicated Data Types).
CRDT (like Yjs) is simpler.
I implement a Go CRDT library.
The browser sends updates (`merge(A, B)`). Go server applies merge.
Both sides eventually converge to the same state without conflicts."

#### Indepth
**LWW-Element-Set**. A simple CRDT. You keep an "Add Set" and a "Remove Set" with timestamps. To check if Item X exists: Is it in Add Set? Is it in Remove Set? If in both, implies "Add" timestamp > "Remove" timestamp. This resolves the "I deleted it while you edited it" conflict mathematically.

---

### 694. How do you implement real-time location tracking?
"I use **Geo-Hashing** (e.g., Redis `GEOADD`).
Driver sends (Lat, Lon) every 5s.
Go server updates Redis.
Rider subscribes to `GEOSEARCH radius=1km`.
Go server polls Redis and pushes updates to the Rider via WebSocket."

#### Indepth
**Geohash Precision**. Geohash is a string (`u4pruydqqv`). The longer the string, the smaller the box. "Near" = "Share same prefix". You can find nearby drivers by searching `u4pru*` (broad) or `u4pruy*` (narrow). This turns 2D spatial search into a 1D string prefix search, which is O(1) in generic KV stores.

---

### 695. How do you use Go in resource-constrained IoT devices?
"I use **TinyGo**.
I compile to an ARM binary or WASM.
I avoid `fmt` and `reflect` if possible.
I use `sync.Pool` aggressively to avoid GC.
For communication, I use **MQTT** (`paho.mqtt.golang`) over TCP, which is lighter than HTTP."

#### Indepth
**GC Tuning**. On 64KB RAM devices, Go's GC is heavy. In TinyGo, the GC is simpler (conservative). You can often disable it (`GOGC=off`) and statically allocate all memory at startup (buffers, structs) to ensure deterministic timing, which is critical for controlling hardware pins.

---

### 696. How do you collect telemetry data from IoT devices?
"Device publishes to MQTT Topic `telemetry/device-1`.
Go Service subscribes to `telemetry/#`.
It decodes the binary payload.
It batches points into Time-Series DB (InfluxDB).
I maintain no state on the ingestion service, allowing me to scale it purely based on CPU load."

#### Indepth
**High Cardinality**. Don't create a new measurement for every device ID. `device_cpu_usage` is the measurement. `device_id=123` is a "Tag". Time-Series DBs index Tags. If you include `error_message` (random string) as a Tag, the index explodes. Store random data as "Fields" (unindexed values).

---

### 697. How do you compress and transmit data from edge devices in Go?
"I use **Protobuf** (small payload).
I enable GZIP/ZSTD on the transport.
For extreme cases, I use **Delta Encoding**: only send the *change* (`temp +0.1`) rather than full value (`25.1`)."

#### Indepth
**Gorilla Compression (XOR)**. Time series data often changes slightly (`timestamp` + 10s, `value` + 0.01). XOR the current value with the previous one. The result is mostly zeros. Run Run-Length Encoding on the zeros. This compresses float streams by 10x-20x, used by Prometheus internally.

---

### 698. How do you implement OTA (over-the-air) updates using Go?
"Device polls `GET /updates?version=1.0`.
Server returns URL to binary `v2.0` + Checksum.
Device downloads, verifies hash.
Go binary uses `equinox` or `rupdate` library to apply the binary patch (replacing itself) and restart."

#### Indepth
**A/B Partitioning**. Embedded Linux standard. Device has partition A (active) and B (idle). You flash the new update to B. Reboot. Bootloader tries B. If it fails (panics/watchdog timeout), it automatically falls back to A. This prevents "bricking" the device remotely.

---

### 699. How do you design protocols for edge-device communication?
"I prefer **Binary, Header-Length-Value (TLV)** protocols over JSON.
Header: `[Type: 1b][Len: 2b]`.
Value: `[Payload]`.
It’s trivial to parse in Go (`binary.Read`) and saves bandwidth.
JSON parsing is too CPU-intensive for very small battery-powered microcontrollers."

#### Indepth
**Endianness**. When communicating between Go (Cloud/amd64 - Little Endian) and a Microcontroller (Big Endian), `binary.Read` handles the conversion if you specify `binary.BigEndian`. Ignoring this flips your `uint16` values (Value `1` becomes `256`), causing subtle data corruption.

---

### 700. How do you build secure, low-latency edge APIs in Go?
"I use **mTLS** (Mutual TLS).
Every device has a burned-in Certificate.
Go server validates the cert chain.
This avoids the latency of an OAuth Handshake (Roundtrip) on every connection. Auth is done during the TCP handshake itself."

#### Indepth
**Secure Enclave**. Store the Private Key in a TPM or Secure Enclave (Hardware). The Go app asks the TPM to "Sign this data". The Private Key never enters RAM. If the device is hacked, the attacker cannot steal the key, only use it while they have access.
