# 🌐 Networking — Deep Dive Interview Questions (Product-Based Companies)

This document covers advanced networking concepts for product-based company interviews (Google, Meta, Amazon, Flipkart, Razorpay, Swiggy, CRED, Zepto, Groww). Targeted at 3–10 years of experience rounds.

---

### Q1: How does TCP's congestion control work? Explain Slow Start, Congestion Avoidance, and Fast Retransmit.

**Answer:**
TCP's congestion control prevents a sender from overwhelming the network. It uses a **Congestion Window (cwnd)** to limit how much unacknowledged data can be in-flight.

**Phases:**

**1. Slow Start:**
- cwnd starts at 1 MSS (Max Segment Size, ~1460 bytes).
- For every ACK received, cwnd doubles (exponential growth).
- Continues until cwnd reaches **ssthresh** (slow start threshold).

**2. Congestion Avoidance:**
- After ssthresh, cwnd grows linearly (increases by 1 MSS per RTT).
- Prevents runaway growth.

**3. Fast Retransmit:**
- If sender receives **3 duplicate ACKs** (not a timeout), it retransmits the lost segment immediately — doesn't wait for timeout.
- Timeout is expensive (RTO can be seconds); 3 dup-ACK detection is fast.

**4. Upon Loss Detection:**
- **Timeout**: ssthresh = cwnd/2, cwnd resets to 1 (Slow Start from scratch). Tahoe algorithm.
- **3 Dup-ACKs**: ssthresh = cwnd/2, cwnd = ssthresh (skip Slow Start). Reno algorithm.

**Modern algorithms:**
- **CUBIC** (Linux default): More aggressive, better for high-bandwidth links.
- **BBR (Bottleneck Bandwidth and RTT)**: Google's algorithm — models network bandwidth directly, achieves higher throughput with lower latency.

```
cwnd
|  /
|  / ← Slow Start (exponential)
| /
|/_____ ← ssthresh
|       \ ← Congestion Avoidance (linear)
|        \
|         ← Packet loss event, reduce cwnd
```

---

### Q2: Explain the TIME_WAIT state in TCP. Why is it important and what problems can it cause?

**Answer:**
After a TCP connection is fully closed, the initiating side (active closer) waits in **TIME_WAIT** for **2 × MSL (Maximum Segment Lifetime)** — typically 60–120 seconds before the socket is fully released.

**Why TIME_WAIT exists:**
1. **Ensures reliable connection close**: The final ACK for the server's FIN might get lost. If it does, the server retransmits FIN. Client must still be alive to respond.
2. **Prevents stale packets**: Old delayed packets from a previous connection could accidentally be accepted by a new connection on the same 4-tuple (src IP:port + dst IP:port). TIME_WAIT guarantees old packets expire first.

**4-Way Close:**
```
Client → Server: FIN
Server → Client: ACK
Server → Client: FIN
Client → Server: ACK  ← Client enters TIME_WAIT here
```

**Problems at scale:**
- High-traffic servers (Nginx, load balancers) can accumulate tens of thousands of TIME_WAIT sockets.
- Each TIME_WAIT socket holds kernel memory.
- Port exhaustion: Ephemeral ports (typically 32768–60999) can all be in TIME_WAIT.

**Solutions:**
```bash
# Enable TCP socket reuse for TIME_WAIT sockets
net.ipv4.tcp_tw_reuse = 1

# Reduce TIME_WAIT duration (not recommended for internet-facing servers)
net.ipv4.tcp_fin_timeout = 15

# Increase ephemeral port range
net.ipv4.ip_local_port_range = 1024 65535
```
- **SO_REUSEADDR** / **SO_REUSEPORT** socket options allow reuse.
- HTTP keep-alive reduces connection teardown frequency.

---

### Q3: What is Head-of-Line (HoL) blocking? How does HTTP/2 solve it partially and how does HTTP/3 solve it fully?

**Answer:**
**HoL blocking** occurs when a "heavy" request/packet blocks subsequent requests/packets from being processed.

**HTTP/1.1 HoL:**
- Only one outstanding request per connection.
- A slow response blocks all subsequent requests.
- Browsers workaround: Open 6 parallel TCP connections per domain.

**HTTP/2 and TCP HoL:**
- HTTP/2 multiplexes multiple streams over ONE TCP connection.
- If a TCP packet is lost, TCP must wait for retransmit before passing ANY data to the application — even if it belongs to a different HTTP/2 stream.
- This is **TCP-level HoL blocking** — HTTP/2 can't solve it because TCP is ordered.

**HTTP/3 and QUIC solution:**
- QUIC is built on UDP, which has no ordering guarantee at the transport layer.
- QUIC independently manages each stream.
- A lost UDP packet only stalls the stream it belongs to.
- Other streams proceed without interruption.

```
HTTP/2 over TCP:
Stream 1: [P1][P2][P3]   ← P2 lost
Stream 2: [Q1][Q2]       ← BLOCKED waiting for P2 retransmit

HTTP/3 over QUIC:
Stream 1: [P1][?][P3]    ← P2 lost, retransmitting
Stream 2: [Q1][Q2]       ← Continues immediately, not blocked
```

---

### Q4: How does a CDN (Content Delivery Network) work? What is anycast routing?

**Answer:**
A **CDN** is a globally distributed network of servers (PoPs — Points of Presence) that caches content closer to users, reducing latency.

**How CDN works:**
1. User requests `https://static.myapp.com/image.png`.
2. DNS resolves to the nearest CDN PoP (e.g., Akamai → Mumbai PoP for Indian users).
3. CDN checks its cache:
   - **Cache hit**: Serves directly from edge server (latency: 5–20ms).
   - **Cache miss**: CDN fetches from origin server, caches it, then serves.
4. Subsequent requests from India hit the cached copy.

**Anycast Routing:**
A single IP address is announced from multiple geographic locations simultaneously via BGP.
When a user sends a packet to that IP, the internet's routing protocol automatically directs it to the **nearest** PoP (by BGP hop count).

```
IP: 1.2.3.4 announced from Mumbai, Singapore, London
Users in India → Mumbai (nearest BGP distance)
Users in Europe → London
Users in SE Asia → Singapore
```

**Uses:** Cloudflare, Google's DNS (8.8.8.8), DDoS mitigation.

**CDN cache invalidation strategies:**
- **TTL-based**: Cache expires after N seconds.
- **Purge API**: Explicitly invalidate on deploy (`Cache-Control: no-store` + CDN purge API).
- **Versioned URLs**: `/static/app.v3.js` — changing the filename forces fresh fetch.

---

### Q5: What is the difference between long polling, Server-Sent Events (SSE), and WebSockets?

**Answer:**
These are three patterns for pushing real-time data from server to client over HTTP infrastructure.

| Feature | Long Polling | SSE | WebSocket |
|---|---|---|---|
| Protocol | HTTP | HTTP | WebSocket (ws://) |
| Direction | Simulated push | Server → Client only | Full duplex (bidirectional) |
| Connection | Re-established per message | Single persistent HTTP stream | Single persistent socket |
| Browser support | All browsers | Not IE, needs proxy fixes | All modern browsers |
| Firewall friendly | Yes | Yes | Sometimes blocked |
| Overhead | High (repeated reconnects) | Low | Lowest |

**Long Polling:**
- Client sends request, server holds it open until there's data (or timeout).
- After response, client immediately re-requests. Simulates push.
- Problem: High overhead, complex server implementation.

**SSE (Server-Sent Events):**
- Single HTTP connection, server streams `data: ...\n\n` events.
- Auto-reconnect built into browser.
- One-way only (server → client).
- Use case: Live sports scores, stock tickers, notifications.

**WebSocket:**
- HTTP upgrade handshake, then switches to binary framing protocol.
- True bidirectional — send and receive simultaneously.
- Use case: Chat applications, collaborative editing, online gaming.

---

### Q6: Explain BGP (Border Gateway Protocol). How does the internet route traffic?

**Answer:**
**BGP** is the routing protocol of the internet — it determines how data travels between different Autonomous Systems (AS).

**Key concepts:**
- **AS (Autonomous System)**: A collection of IP networks under a single administrative domain (e.g., Jio = AS55836, Google = AS15169).
- **AS Number (ASN)**: Unique identifier for each AS.
- **BGP Peers**: BGP routers exchange routing information.

**How BGP works:**
1. Each AS announces its IP prefixes (e.g., "I own 74.125.0.0/16") to peer ASes.
2. BGP routers build a table of paths to reach every prefix.
3. BGP selects the "best" path based on attributes (AS path length, local preference, MED, etc.).
4. Once a path is selected, it's installed in the routing table.

**Types:**
- **iBGP**: BGP between routers within the same AS.
- **eBGP**: BGP between different ASes — how the internet connects.

**BGP hijacking:** A malicious AS announces more specific prefixes for another AS's IPs — traffic gets redirected. Famous incidents: China Telecom hijacking routes (2010), Pakistan Telecom accidentally blackholing YouTube (2008).

**Mitigation:** RPKI (Resource Public Key Infrastructure) — cryptographically signs IP prefix-AS origin pairs.

---

### Q7: How do you design a low-latency network for a system that needs < 10ms P99 latency globally?

**Answer:**
This is a system design-oriented networking question.

**Strategies:**

**1. Geographic Distribution:**
- Deploy servers in multiple regions (multi-region architecture).
- Route users to nearest region using Anycast or GeoDNS.
- Use CDN edge for static assets.

**2. Connection Optimization:**
- Use HTTP/3 + QUIC for 0-RTT reconnects (saves one round trip).
- Implement TCP connection pooling (don't tear down connections between requests).
- Enable TCP Fast Open (TFO) for repeat connections.

**3. Protocol Choices:**
- Use gRPC over HTTP/2 for internal microservice calls (binary protocol, multiplexed).
- For real-time, use WebSockets or QUIC.
- Minimize roundtrips: combine API calls, use batching.

**4. Network Hardware/OS Tuning:**
```bash
# Increase TCP buffer sizes
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216
net.ipv4.tcp_rmem = 4096 87380 16777216

# Use SO_REUSEPORT for better load distribution across CPU cores
# Enable TCP_NODELAY to disable Nagle's algorithm (critical for low latency)
```

**5. Eliminate bottlenecks:**
- Use kernel bypass networking (DPDK, io_uring) for ultra-low latency.
- Co-locate services (same datacenter, same rack) to minimize physical RTT.
- Use private internet backbone (Google's B4, AWS Direct Connect) instead of public internet.

**Measurement:** Always profile with percentiles (P50, P95, P99, P999) — averages hide tail latencies.

---

*Prepared for technical rounds at product-based companies (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay, Groww).*
