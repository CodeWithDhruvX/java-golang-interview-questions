## 🔴 Real-Time Systems, IoT, and Edge Computing (Questions 681-700)

### Question 681: How do you build a real-time chat server in Go?

**Answer:**
1.  **Transport:** WebSockets (`gorilla/websocket`).
2.  **Hub:** A central struct managing `Register`, `Unregister`, and `Broadcast` channels.
3.  **Client:** Goroutine reading form WS, Goroutine writing to WS.

```go
type Hub struct {
    clients map[*Client]bool
    broadcast chan []byte
}

func (h *Hub) Run() {
    for {
        select {
        case msg := <-h.broadcast:
            for client := range h.clients {
                client.send <- msg
            }
        }
    }
}
```

---

### Question 682: How do you implement WebSockets in Go?

**Answer:**
Use `github.com/gorilla/websocket`.
Upgrade logic:
```go
var upgrader = websocket.Upgrader{}
func handler(w http.ResponseWriter, r *http.Request) {
    conn, _ := upgrader.Upgrade(w, r, nil)
    defer conn.Close()
    for {
        _, msg, _ := conn.ReadMessage()
        conn.WriteMessage(websocket.TextMessage, msg)
    }
}
```

---

### Question 683: How do you ensure order of events in real-time systems?

**Answer:**
WebSockets over TCP guarantee order **per connection**.
For global ordering (multiple users):
- Use a central event bus (Redis/Kafka) with a sequential ID or Lamport Timestamp.
- Client re-sorts buffer based on Sequence ID.

---

### Question 684: How do you handle high concurrency in WebSocket servers?

**Answer:**
**epoll/kqueue (Advanced):**
Standard Go spins 2 goroutines per WS (Read/Write). For 1M connections = 2M goroutines (4GB+ RAM).
Use `github.com/gobwas/ws` or `gnet` (event-loop based) to handle IO with fewer goroutines using non-blocking syscalls (Epoll).

---

### Question 685: How do you implement presence tracking in Go (like online users)?

**Answer:**
**Heartbeat.**
- Client sends "Ping" every 30s.
- Server updates Redis `SET user:123 "online" EX 40` (TTL 40s).
- Total Online: `SCAN` Redis for keys.

---

### Question 686: How do you reduce latency in real-time systems?

**Answer:**
1.  **Protocol:** Use Binary (Protobuf) over Text (JSON) on WS.
2.  **Geo-Distribution:** Deploy Go servers to Edge (AWS Global Accelerator).
3.  **GC:** Tune `GOGC` or use arena allocation to avoid GC pauses during broadcast.

---

### Question 687: How do you build a real-time dashboard backend in Go?

**Answer:**
**Server-Sent Events (SSE).**
Simpler than WebSockets for Read-Only dashboards.
```go
w.Header().Set("Content-Type", "text/event-stream")
for {
    fmt.Fprintf(w, "data: %s\n\n", getStats())
    w.Flush()
    time.Sleep(1 * time.Second)
}
```

---

### Question 688: How do you handle message fan-out for WebSocket clients?

**Answer:**
Broadcast 1 message to 100k users.
Looping sequentially is too slow.
**Sharding:**
Split clients into 100 Shards (Hubs).
Spin up 100 Goroutines.
Each goroutine iterates its own list and writes messages.

---

### Question 689: How do you design a real-time bidding system in Go?

**Answer:**
Latency is critical (<50ms).
- Keep auction state in **Memory** (Go `map` protected by `RWMutex`), not Redis.
- Use Channels to linearize bids: `BidsChan <- bid`.
- Single processor goroutine determines winner.
- Log to disk async (WAL) for durability.

---

### Question 690: How do you throttle real-time updates?

**Answer:**
**Debounce/Conflate.**
If stock price changes 100 times/sec, don't send 100 JSONs.
Send 1 JSON every 200ms with the *latest* price.
Use `time.Ticker` inside the client writer loop.

---

### Question 691: How do you buffer real-time data safely?

**Answer:**
Ring Buffer (Circular Queue).
Fixed size (e.g., last 100 messages).
No GC overhead (reuse memory).
If consumer is slow, overwrite old data (Drop oldest) to prevent OOM.

---

### Question 692: How do you build a publish-subscribe engine for WebSockets?

**Answer:**
Use **Redis Pub/Sub** as the backend.
1.  Go Server A, B, C subscribe to Redis channel `chat`.
2.  User on Server A sends msg.
3.  Server A publishes to Redis.
4.  Server B & C receive msg from Redis -> Broadcast to their local WS connections.

---

### Question 693: How do you sync real-time state between browser and Go backend?

**Answer:**
**CRDTs (Conflict-free Replicated Data Types).**
Send operations ("Add char 'a' at pos 5") instead of full state.
Libraries: `Yjs` (JS) + Go port.

---

### Question 694: How do you implement real-time location tracking?

**Answer:**
**Geospatial Indexing.**
- Redis Geo (`GEOADD`, `GEORADIUS`).
- In-Memory: QuadTree or R-Tree implemented in Go.
- Update positions, Query "Drivers near User" efficiently.

---

### Question 695: How do you use Go in resource-constrained IoT devices?

**Answer:**
**TinyGo.** (See Q483).
Compiles to ARM thumb instructions, runs on bare metal (no OS) or RTOS.
Control GPIO pins (`machine.Pin`).

---

### Question 696: How do you collect telemetry data from IoT devices?

**Answer:**
MQTT Protocol.
Use `eclipse/paho.mqtt.golang`.
Go server acts as Subscriber to MQTT Broker (Mosquitto), creates efficient pipes to InfluxDB/TimescaleDB.

---

### Question 697: How do you compress and transmit data from edge devices in Go?

**Answer:**
Use **MessagePack** or **Protobuf** (Binary).
Gzip compression might be too heavy for CPU.
Snappy or Zstd is better.

---

### Question 698: How do you implement OTA (over-the-air) updates using Go?

**Answer:**
1.  Device polls Server `GET /version`.
2.  If new: Download binary (Go `http.Get`), compute checksum (`sha256`).
3.  **A/B Partitioning:** Write to Partition B. Verify. Switch Boot flag. Reboot. (Library: `mender`).

---

### Question 699: How do you design protocols for edge-device communication?

**Answer:**
Use **CoAP** (Constrained Application Protocol) or **MQTT-SN**.
Go has libraries (`go-coap`).
UDP-based, low overhead, suitable for unstable networks (LoRaWAN / Cellular).

---

### Question 700: How do you build secure, low-latency edge APIs in Go?

**Answer:**
- **TLS 1.3 0-RTT:** Resumption for faster handshake.
- **QUIC / HTTP/3:** Use `lucas-clemente/quic-go`. Runs over UDP, solves Head-of-Line blocking, better for mobile networks switching between WiFi/4G.

---
