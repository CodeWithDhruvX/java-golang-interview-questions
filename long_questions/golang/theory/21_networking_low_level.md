# 🟢 Go Theory Questions: 401–420 Networking and Low-Level Programming

## 401. How do you create a TCP server in Go?

**Answer:**
We use the `net` package.

1.  Listen: `ln, _ := net.Listen("tcp", ":8080")`.
2.  Accept Loop: `for { conn, _ := ln.Accept(); go handle(conn) }`.
3.  Handle: `conn.Read()` and `conn.Write()`.

The magic is the `go handle(conn)`. Because goroutines are cheap, we spawn one per connection. This allows us to handle 10,000 concurrent TCP connections with simple blocking I/O code, whereas in C/Node.js you'd need complex event loops/callbacks.

---

## 402. How do you create a UDP client in Go?

**Answer:**
UDP is connectionless, but we still "Dial" to set a default destination.

`conn, _ := net.Dial("udp", "127.0.0.1:8080")`
`fmt.Fprintf(conn, "Message")`

Unlike TCP, `Dial` doesn't actually send packets (no handshake). It just sets the socket address. If the server is down, `Write` still "succeeds" because UDP is fire-and-forget. We use UDP for metrics (StatsD) or real-time gaming where packet loss is acceptable.

---

## 403. What is the difference between `net.Listen` and `net.Dial`?

**Answer:**
`net.Listen` creates a **Server Socket**. It binds to a local port and waits for incoming connections (Passive Open).

`net.Dial` creates a **Client Socket**. It initiates a connection to a remote address (Active Open).

In TCP three-way handshake terms: `Dial` sends the SYN. `Listen` (via Accept) sends the SYN-ACK.

---

## 404. How do you manage TCP connection pools?

**Answer:**
The standard library `sql.DB` and `http.Client` handle pooling automatically.

For custom protocols, we implement a pool using a **Buffered Channel** of connections.
`pool := make(chan net.Conn, 10)`
Get: `conn := <-pool`.
Put: `pool <- conn`.

If the channel is empty, we `Dial` a new one. If full, we close the connection (or block). We must also handle "Health Checks"—if a connection sits in the pool for an hour, the firewall might silently drop it, so we check readiness before returning it to the caller.

---

## 405. How would you implement a custom HTTP transport?

**Answer:**
We wrap the default `http.Transport`.

It allows us to intercept the request *at the socket level*.
Common use case: **Logging** or **Authentication**.
`t := &http.Transport{ DialContext: func(...) { ... } }`.
We can inject logic to route traffic through a SOCKS5 proxy, or force a specific DNS resolution logic, or enable HTTP/2 only.

---

## 406. How do you read raw packets using `gopacket`?

**Answer:**
`gopacket` (from Google) uses `libpcap` to capture raw ethernet frames (like Wireshark).

We open a handle: `pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)`.
We process the packet source: `packet := source.NextPacket()`.

We can inspect layers: `ipLayer := packet.Layer(layers.LayerTypeIPv4)`.
This is used for building firewalls, network analyzers, or deep packet inspection tools directly in Go.

---

## 407. What is a connection hijack in `net/http` and how is it done?

**Answer:**
Hijacking lets a handler take over the underlying TCP connection from the HTTP server.

Useful for **WebSockets**.
`hj, ok := w.(http.Hijacker)`
`conn, buf, _ := hj.Hijack()`

Once hijacked, the standard HTTP server stops managing that connection. You are now responsible for reading/writing raw bytes to the TCP socket and closing it when done. If you forget to close it, you leak a file descriptor.

---

## 408. How to implement a proxy server in Go?

**Answer:**
The `net/http/httputil` package provides `ReverseProxy`.

`proxy := httputil.NewSingleHostReverseProxy(targetURL)`
`http.ListenAndServe(":8080", proxy)`

It automatically handles forwarding headers, copying the body, and connection pooling.
For a forward proxy (CONNECT method), we manually `Hijack` the connection, Dial the target, and shovel bytes between the two connections using `io.Copy(dest, src)` in two goroutines.

---

## 409. How would you create an HTTP2 server from scratch in Go?

**Answer:**
You don't need to do anything "from scratch." Go's `net/http` supports HTTP/2 automatically if TLS is enabled.

`http.ListenAndServeTLS(...)`.
During the TLS handshake (ALPN), if the client supports `h2`, Go transparently switches to the HTTP/2 framer.
If you *must* force H2C (HTTP/2 Cleartext, no TLS), you need `golang.org/x/net/http2` and specifically configure the `H2C` server handler, as browsers generally don't support H2C.

---

## 410. How does Go handle connection reuse (keep-alive)?

**Answer:**
Go's `http.Client` enables Keep-Alive by default.

After a request finishes, the TCP connection is not closed. It is placed in an "Idle Pool" inside the Transport.
The next request to the same host reuses this connection, skipping the expensive TCP/TLS handshake.
If you see "Too Many Open Files" errors, you might be forgetting to close `resp.Body`—Go only returns the connection to the pool after the body is fully read and closed.

---

## 411. How do you set timeouts on sockets in Go?

**Answer:**
We use `conn.SetDeadline(time.Now().Add(5 * time.Second))`.

This applies to both Read and Write.
If a Read blocks longer than 5s, it returns an "i/o timeout" error.
This is critical for preventing "Slowloris" attacks where an attacker opens 1000 connections and sends 1 byte per minute to keep them open, exhausting server resources.

---

## 412. What is the difference between `net/http` and `fasthttp`?

**Answer:**
`net/http` allocates a new goroutine for every request and is generous with memory allocations (user-friendly interfaces).

`fasthttp` is a zero-allocation library. It reuses goroutines (worker pool) and reuses byte buffers aggressively.
**Trade-off**: `fasthttp` implementation is complex and doesn't fully conform to standard HTTP/2. We only use it if we are building a gateway doing 1M+ RPS. For 99% of apps, `net/http` is fast enough and safer.

---

## 413. How do you throttle network traffic in Go?

**Answer:**
We wrap the `net.Conn`.

We implement a `ThrottledConn` struct that implements `io.Reader`.
In `Read()`, we check a **Token Bucket**. If we want 1MB/s, we allow reading 1KB every 1ms. If the bucket is empty, we sleep.
This is mostly used in file servers or backup tools to prevent saturating the entire uplink.

---

## 414. How would you analyze network latency in Go?

**Answer:**
We use `httptrace`.

`trace := &httptrace.ClientTrace{ GotConn: func(info) { ... }, DNSDone: func(info) { ... } }`
`ctx := httptrace.WithClientTrace(context.Background(), trace)`

This hooks into the HTTP client lifecycle. We can measure exactly how many milliseconds were spent on DNS lookup, TCP Handshake, TLS Handshake, and First Byte Received.

---

## 415. How would you implement WebRTC or peer-to-peer comms?

**Answer:**
We use the `pion/webrtc` library (Pure Go implementation).

Standard WebRTC requires:
1.  **Signaling**: (Go WebSocket server) to exchange SDP offers/answers.
2.  **STUN/TURN**: To punch through NATs.
3.  **Data Channels**: UDP-based streams.
Go's concurrency model excels here because managing thousands of UDP packet streams is CPU-efficient compared to C++ based solutions wrapped in Node.js.

---

## 416. How do you simulate a slow network in integration tests?

**Answer:**
We use a **Proxy Middleware** or `toxiproxy`.

In pure Go, we can wrap our `http.Handler`.
`func(w, r) { time.Sleep(2 * time.Second); realHandler(w, r) }`.
Or wrap the `net.Listener` to delay `Accept()`.
This proves that our client-side timeouts are correctly configured (e.g., ensuring the app cancels the context after 1s instead of hanging for 2s).

---

## 417. What’s the difference between connection pooling and multiplexing?

**Answer:**
**Pooling**: Reusing a connection for sequential requests. Request 1 finishes, then Request 2 starts. (HTTP/1.1).

**Multiplexing**: Sending multiple requests *concurrently* over the *same* connection. Request 2 can start before Request 1 finishes. (HTTP/2, gRPC).
Go supports both transparently. Multiplexing is vastly more efficient for high-latency links (solves Head-of-Line blocking).

---

## 418. How do you verify DNS lookups in Go?

**Answer:**
`net.LookupHost("example.com")` returns the IP addresses.
`net.LookupMX` returns mail servers.

By default, Go uses a pure Go resolver. However, on systems like Android or when using CGO, it might use `getaddrinfo` (OS libc).
You can force the Go resolver: `export GODEBUG=netdns=go`. This is safer for static binaries as it doesn't depend on `/etc/nsswitch.conf`.

---

## 419. How do you use HTTP pipelining in Go?

**Answer:**
**We don't.**

HTTP Pipelining (HTTP/1.1 feature) is disabled by default in most browsers and servers (including Go) because of Head-of-Line blocking issues.
Instead, we use **HTTP/2**, which solves the same problem correctly via Multiplexing. The Go client automatically upgrades to H2 if supported.

---

## 420. How do you implement NAT traversal in Go?

**Answer:**
We use **STUN** (Session Traversal Utilities for NAT).

Library: `pion/stun`.
We send a UDP packet to a public STUN server (like Google's). It replies with "Here is your Public IP and Port as seen from the internet."
We use this public address to tell peers "Connect to me here."
If STUN fails (Symmetric NAT), we fall back to **TURN** (Relaying traffic through a middleman server), which uses more bandwidth but guarantees connectivity.
