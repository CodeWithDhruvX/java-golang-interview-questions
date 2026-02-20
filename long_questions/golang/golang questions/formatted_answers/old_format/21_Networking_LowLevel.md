# 🔵 **401–420: Networking and Low-Level Programming**

### 401. How do you create a TCP server in Go?
"I use `net.Listen`.

`ln, _ := net.Listen("tcp", ":8080")`.
`for { conn, _ := ln.Accept(); go handle(conn) }`.
The `go handle(conn)` part is critical. It spawns a lightweight goroutine for every connection.
This allows me to handle 10,000 concurrent clients with minimal memory overhead, unlike the 'Thread-per-Client' model in older Java/C++ servers."

#### Indepth
Under the hood, Go uses the OS's non-blocking I/O (epoll on Linux, kqueue on BSD/macOS, IOCP on Windows). The Go runtime's "Netpoller" integrates with the scheduler. When a goroutine calls `conn.Read()`, if no data is available, it parks the goroutine and registers the file descriptor with epoll. When data arrives, the Netpoller wakes up the goroutine. This makes sync-looking code run asynchronously.

---

### 402. How do you create a UDP client in Go?
"I use `net.Dial("udp", "server:port")`.

Unlike TCP, there is no handshake. I just write bytes: `conn.Write([]byte("ping"))`.
Since UDP is unreliable, I must be prepared for packet loss. If I need reliability, I implement Application-Layer Acks and Retries, or I just switch to TCP if latency permits."

#### Indepth
UDP is "fire and forget", but be aware of **MTU (Maximum Transmission Unit)**. If you send a 5KB packet, IP fragmentation will slice it into ~1500-byte frames. If *one* frame is lost, the entire packet is dropped by the kernel. Keep UDP packets under 1472 bytes (Ethernet MTU - headers) to avoid fragmentation and increase reliability.

---

### 403. What is the difference between `net.Listen` and `net.Dial`?
"**Listen** opens a port and waits (Server).
**Dial** initiates a connection (Client).

Once established, both return a `net.Conn` object.
`net.Conn` implements `io.Reader` and `io.Writer`.
So, whether I am the server or the client, reading and writing bytes looks exactly the same."

#### Indepth
`net.Conn` is an interface, but the underlying implementation is usually `*net.TCPConn`. You can type-assert it to access TCP-specific methods like `SetKeepAlive`, `SetNoDelay` (Nagle's Algorithm), or `SetReadBuffer`. Disabling Nagle (`SetNoDelay(true)`) is critical for low-latency apps like games or trading.

---

### 404. How do you manage TCP connection pools?
"Usually, I don't. `net/http` handles it for me.

But for raw TCP, I implement a **Pool** pattern.
I use a buffered channel of `net.Conn`.
Get: `select { case conn := <-pool: return conn; default: return net.Dial(...) }`.
Put: `select { case pool <- conn: return; default: conn.Close() }`.
I must also handle 'Zombie Connections' by checking errors on Read/Write and discarding bad connections."

#### Indepth
Implementing a robust pool is hard. You must handle "max idle connections", "max lifetime", and "health checks". If a connection sits idle for 1 hour, the firewall might silently drop it. Your pool must try to Read/Write (heartbeat) or discard it on usage to avoid "Broken Pipe" errors in the middle of a business transaction.

---

### 405. How would you implement a custom HTTP transport?
"I wrap `http.DefaultTransport`.

`t := &http.Transport{ Proxy: ..., DialContext: ... }`.
I can override `DialContext` to route traffic over a custom tunnel (like SSH or SOCKS5) or override `RoundTrip` to log every request's detailed timing (DNS, TLS, TTFB).
This is how Service Meshes inject specific sidecar logic into the Go client."

#### Indepth
`http.RoundTripper` is powerful but be careful: it must be safe for concurrent use. Also, if you implement `RoundTrip`, you are responsible for handling redirects (unless you wrap `DefaultTransport`). A common pattern is **Circuit Breaking** at the transport level, so *every* HTTP request in your app automatically respects failure thresholds.

---

### 406. How do you read raw packets using `gopacket`?
"I use the `google/gopacket` library.

`handle, _ := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)`.
`packetSource := gopacket.NewPacketSource(handle, handle.LinkType())`.
I iterate over `packetSource.Packets()`.
This gives me access to the raw Ethernet frames, IP headers, and TCP flags. It’s perfect for building custom firewalls or network analyzers."

#### Indepth
Packet capture requires elevated privileges (root/Admin). Also, typical `pcap` is slow for high-throughput (10Gbps+). For production network appliances in Go, consider using **AF_XDP** (Linux eXpress Data Path) or specialized zero-copy drivers (`pf_ring`) to bypass the kernel stack entirely.

---

### 407. What is a connection hijack in `net/http` and how is it done?
"It allows my handler to take over the underlying TCP socket.

`conn, buf, _ := w.(http.Hijacker).Hijack()`.
Once I hijack, the HTTP server forgets about this connection.
I now have raw Read/Write access.
This is exactly how **WebSockets** work: they start as an HTTP GET, handshake, and then hijack the socket to switch to a binary protocol."

#### Indepth
When you `Hijack()`, you are responsible for managing the connection *entirely*. The `net/http` server will not close it, will not log it, and will not handle timeouts. You effectively steal the file descriptor from the runtime's HTTP loop. Use this sparingly; it breaks standard middleware chains.

---

### 408. How to implement a proxy server in Go?
"I use `httputil.ReverseProxy`.

`proxy := httputil.NewSingleHostReverseProxy(targetURL)`.
`http.ListenAndServe(":8080", proxy)`.
It automatically handles forwarding headers, buffering bodies, and dealing with backend failures.
For a forward proxy (tunneling HTTPS), I handle the `CONNECT` method and simply shovel bytes between the client and the destination."

#### Indepth
`httputil.ReverseProxy` defaults are sometimes dangerous. It sets `FlushInterval` to 0 (flush immediately), which is good for latency but bad for compression. It also relies on the standard `http.Transport`. For high-load proxies, optimize the `BufferPool` to reuse `[]byte` buffers and prevent GC thrashing.

---

### 409. How would you create an HTTP2 server from scratch in Go?
"I don't. `net/http` supports it transparently.

`http.ListenAndServeTLS(":443", "cert", "key", handler)`.
During the TLS handshake (ALPN), if the client supports `h2`, Go automatically upgrades the connection to HTTP/2.
I verify this by checking `req.Proto` in my handler. Writing a compliant H2 server from scratch is a massive undertaking."

#### Indepth
Go's HTTP/2 implementation (bundled in `net/http`) creates a new goroutine for *every stream*. If a malicious client opens 10,000 streams on a single TCP connection, it can cause memory exhaustion. Go 1.18+ added `MaxConcurrentStreams` to `http.Server` to mitigate this 'Stream Flood' attack.

---

### 410. How does Go handle connection reuse (keep-alive)?
"By default, the HTTP Client maintains a pool of open connections.

**Crucial Trap**: I MUST read the response body fully and close it.
`io.Copy(io.Discard, resp.Body)`
`resp.Body.Close()`
If I don't do this, the connection is considered 'dirty' and cannot be reused, forcing a new TCP handshake for every request (killing performance)."

#### Indepth
The connection pool defaults are: `MaxIdleConns: 100` (across all hosts) and `MaxIdleConnsPerHost: 2`. This `2` is often a bottleneck for microservices talking to a high-throughput backend. Increase `MaxIdleConnsPerHost` to 100 or more if you are making many parallel requests to the *same* service.

---

### 411. How do you set timeouts on sockets in Go?
"I use `conn.SetDeadline(time)`.

`conn.SetReadDeadline(time.Now().Add(5*time.Second))`.
If no data arrives within 5 seconds, `Read()` returns an error.
This is essential to prevent **Slowloris** attacks, where a malicious client connects, sends one byte every hour, and exhausts my server's file descriptors."

#### Indepth
Deadlines are absolute time points (`time.Now().Add(...)`), not durations. You must reset them *before every read/write* if you want a "rolling timeout". For `net/http`, `http.Server{ReadTimeout: ...}` covers the entire request read, while `ReadHeaderTimeout` covers just the headers (critical for Slowloris protection).

---

### 412. What is the difference between `net/http` and `fasthttp`?
"**net/http**: Standard, robust, allocates 1 goroutine per request.
**fasthttp**: Optimized, zero-allocation, reuses goroutines and buffers aggressively.

`fasthttp` is 10x faster in synthetic benchmarks, but it has a quirky API (no `io.Reader`) and doesn't support HTTP/2. I stick to `net/http` unless I'm building a gateway handling 1 Million RPS."

#### Indepth
`fasthttp` gains speed by avoiding allocations. It recycles Request/Response objects. This means **you cannot pass these objects to another goroutine**. If you do `go func() { process(ctx) }`, the context will be reset/reused by the server while your goroutine is still running, leading to race conditions and data corruption.

---

### 413. How do you throttle network traffic in Go?
"I wrap the `net.Conn`.

I create a `ThrottledConn` struct that intercepts `Read` and `Write`.
It sleeps for `bytes / limit` duration before returning.
`time.Sleep(100 * time.Millisecond)`.
Or I use `golang.org/x/time/rate` to bucket-limit the throughput. This helps simulate 3G networks in integration tests."

#### Indepth
Token Buckets allow "bursts". A rate limit of "10 req/s" might allow 10 requests in the first millisecond, then silence for 999ms. If you need smooth pacing (e.g., sending packets to a hardware device), use a **Leaky Bucket** or strict spacer: `time.Sleep(perPacketDuration)`.

---

### 414. How would you analyze network latency in Go?
"I use `httptrace`.

`trace := &httptrace.ClientTrace{ GotConn: ..., DNSStart: ..., TLSHandshakeStart: ... }`.
I inject this into the context: `req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))`.
It breaks down the latency: 'DNS took 50ms, TCP took 20ms, Server took 500ms'. This proves if the network is the problem."

#### Indepth
`httptrace` is context-aware. Be careful: the hooks are called synchronously. If your `GotConn` hook does a slow DNS reverse-lookup or prints to a slow console, you are slowing down the actual HTTP request! Keep trace hooks extremely lightweight and non-blocking.

---

### 415. How would you implement WebRTC or peer-to-peer comms?
"I use **Pion** (`github.com/pion/webrtc`).

It’s a pure Go implementation of the WebRTC stack.
1.  **Signaling**: Exchange SDP via WebSocket.
2.  **ICE**: Punch holes through NAT/Firewalls.
3.  **DataChannel**: Stream binary data or video.
It allows me to build P2P apps (like Zoom or File Sharing) without a central relay server."

#### Indepth
Pion is huge not just for video, but for **unreliable data channels**. You can send game state updates via UDP-like SCTP data channels directly between browsers. It handles the complexity of DTLS (encryption) and ICE (NAT traversal) purely in Go, making it easier to deploy than a C++ WebRTC stack.

---

### 416. How do you simulate a slow network in integration tests?
"I use a **Middleware** or a Proxy like **Toxiproxy**.

In code:
`func SlowHandler(w, r) { time.Sleep(500 * time.Millisecond); next.ServeHTTP(w, r) }`.
This simulates latency.
To simulate packet loss or bandwidth limits, I use Toxiproxy sidecar containers in my CI environment."

#### Indepth
Simulating latency in middleware isn't perfect—it throttles the *handler*, not the *network*. It won't trigger TCP timeout behavior or retransmissions. For true fidelity, use OS-level tools like `tc` (Traffic Control) on Linux or `pf` on macOS to drop packets at the kernel interface level.

---

### 417. What’s the difference between connection pooling and multiplexing?
"**Pooling** (HTTP/1.1): I open 10 TCP connections. I send 1 request on each concurrently.
**Multiplexing** (HTTP/2): I open 1 TCP connection. I send 100 requests *interleaved* on it.

Multiplexing is superior because it saves OS resources (File Descriptors) and eliminates the TCP Slow Start penalty for new connections."

#### Indepth
HTTP/2 Multiplexing has a downside: **TCP Head-of-Line Blocking**. If one packet is lost, *all* streams on that connection pause until retransmission. HTTP/3 (QUIC) solves this by moving multiplexing to UDP, so packet loss only affects the specific stream it belongs to.

---

### 418. How do you verify DNS lookups in Go?
"I use `net.Resolver`.

`r := &net.Resolver{PreferGo: true}`.
`ips, _ := r.LookupHost(ctx, "google.com")`.
I can configure it to use a specific DNS server (like `8.8.8.8`) instead of the system default. This is useful for debugging Split-Horizon DNS issues inside Kubernetes."

#### Indepth
Go has two resolvers: `cgo` (uses OS `getaddrinfo`) and `go` (pure Go). By default, it uses `go` unless you need things it can't do (like macOS `.local` multicast DNS). You can force one with `export GODEBUG=netdns=cgo` or `go`. In static binaries (Alpine), it's always the pure Go resolver.

---

### 419. How do you use HTTP pipelining in Go?
"**You don't.**

Pipelining (sending req A, req B before res A arrives) was a feature of HTTP/1.1 that caused 'Head-of-Line Blocking' and was universally disabled.
HTTP/2 Multiplexing replaced it. Go's `net/http` does not support pipelining."

#### Indepth
Pipeline requests are processed in order. If Request 1 involves a DB query taking 10s, Request 2 (static file) must wait 10s, even if it arrived at the server. Multiplexing (HTTP/2) allows Request 2 to finish while Request 1 is still working. Pipelining was a flawed optimization.

---

### 420. How do you implement NAT traversal in Go?
"I use **STUN/TURN** protocols via Pion.

**STUN**: Tells me my public IP.
**TURN**: Relays traffic if I'm behind a strict Symmetric NAT.
I run a TURN server (coturn) and configure my Go client to authenticate with it. This allows my P2P app to work even in restrictive corporate networks."

#### Indepth
If STUN/TURN fails, standard fallback is **Relaying**. You route traffic through your own WebSocket server. It consumes your bandwidth but guarantees connectivity. In Pion, you can configure "ICE Candidates" to include your relay server as a last resort.
