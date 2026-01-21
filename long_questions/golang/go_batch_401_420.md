## 🔵 Networking and Low-Level Programming (Questions 401-420)

### Question 401: How do you create a TCP server in Go?

**Answer:**
You use the `net` package. Specifically `net.Listen()` to bind to a port and `listener.Accept()` to handle incoming connections, typically handling each connection in a new goroutine.

```go
func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil { panic(err) }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil { continue }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    // Read/Write logic...
    conn.Write([]byte("Hello TCP\n"))
}
```

---

### Question 402: How do you create a UDP client in Go?

**Answer:**
UDP is connectionless. You can use `net.Dial("udp", address)` to send packets.

```go
func main() {
    conn, err := net.Dial("udp", "localhost:8080")
    if err != nil { panic(err) }
    defer conn.Close()

    conn.Write([]byte("Hello UDP"))
}
```
For a server, you would use `net.ListenPacket("udp", ":8080")`.

---

### Question 403: What is the difference between `net.Listen` and `net.Dial`?

**Answer:**
- **`net.Listen` (Server Side):** Binds to a local port and waits for incoming connections. Returns a `Listener`.
- **`net.Dial` (Client Side):** Initiates a connection to a remote (or local) address and port. Returns a `Conn`.

---

### Question 404: How do you manage TCP connection pools?

**Answer:**
1.  **DB/HTTP:** Libraries like `database/sql` and `net/http` manage pools automatically (e.g., `MaxIdleConns`, `MaxOpenConns`).
2.  **Custom TCP:** You can implement a pool using a buffered channel of `net.Conn`.

```go
type Pool struct {
    conns chan net.Conn
}

func (p *Pool) Get() net.Conn {
    select {
    case conn := <-p.conns:
        return conn
    default:
        return createNewConn()
    }
}

func (p *Pool) Put(conn net.Conn) {
    select {
    case p.conns <- conn:
        // returned to pool
    default:
        conn.Close() // pool full
    }
}
```

---

### Question 405: How would you implement a custom HTTP transport?

**Answer:**
Implement the `http.RoundTripper` interface. This is useful for logging, caching, or authenticating requests globally.

```go
type LoggingTransport struct {
    Transport http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    fmt.Println("Request:", req.URL)
    return t.Transport.RoundTrip(req)
}

func main() {
    client := &http.Client{
        Transport: &LoggingTransport{Transport: http.DefaultTransport},
    }
    client.Get("http://example.com")
}
```

---

### Question 406: How do you read raw packets using gopacket?

**Answer:**
Use the `github.com/google/gopacket` library (wrapper for libpcap).

```go
import "github.com/google/gopacket/pcap"

func main() {
    handle, _ := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    
    for packet := range packetSource.Packets() {
        fmt.Println(packet)
    }
}
```
*Note: Requires root/admin privileges.*

---

### Question 407: What is a connection hijack in `net/http` and how is it done?

**Answer:**
Hijacking allows a handler to take over the underlying TCP connection from the HTTP server, preventing the server from managing the connection (closing/writing to it) after the handler returns. Used for WebSockets.

```go
func funcHandler(w http.ResponseWriter, r *http.Request) {
    hj, ok := w.(http.Hijacker)
    if !ok { http.Error(w, "Hijacking not supported", 500); return }
    
    conn, bufrw, err := hj.Hijack()
    defer conn.Close()
    
    // Now you have the raw TCP connection
    bufrw.WriteString("HTTP/1.1 101 Switching Protocols\r\n\r\n")
}
```

---

### Question 408: How to implement a proxy server in Go?

**Answer:**
Use `httputil.ReverseProxy`.

```go
func main() {
    targetURL, _ := url.Parse("http://background-service:8080")
    proxy := httputil.NewSingleHostReverseProxy(targetURL)
    
    http.ListenAndServe(":80", proxy)
}
```
For a forward proxy, you'd handle the request and `Dial` the target yourself.

---

### Question 409: How would you create an HTTP2 server from scratch in Go?

**Answer:**
Go's `net/http` server enables HTTP/2 automatically if:
1.  You listen on **TLS** (`ListenAndServeTLS`).
2.  The client supports ALPN.

To force H2C (Cleartext HTTP/2):
Use `golang.org/x/net/http2/h2c`.

```go
h2s := &http2.Server{}
http.ListenAndServe(":8080", h2c.NewHandler(myHandler, h2s))
```

---

### Question 410: How does Go handle connection reuse (keep-alive)?

**Answer:**
The default `http.Transport` uses connection pooling (Keep-Alive) by default.
- It keeps idle TCP connections open to avoid the handshake overhead for subsequent requests.
- Managed via `Transport.MaxIdleConns` and `Transport.IdleConnTimeout`.

To **disable**:
```go
tr := &http.Transport{ DisableKeepAlives: true }
client := &http.Client{ Transport: tr }
```

---

### Question 411: How do you set timeouts on sockets in Go?

**Answer:**
Use the `SetDeadline` methods on the `net.Conn` object.

```go
conn, _ := net.Dial("tcp", "host:port")

// Set both read and write deadline
conn.SetDeadline(time.Now().Add(5 * time.Second))

// Or individually
conn.SetReadDeadline(time.Now().Add(2 * time.Second))
```
This forces I/O operations to return an error if they block longer than the deadline.

---

### Question 412: What is the difference between `net/http` and `fasthttp`?

**Answer:**
- **`net/http`:** Standard library. General purpose, HTTP/1.1 & HTTP/2. Interface-based (`io.Reader` body). Safe and idiomatic. Allocates new objects per request.
- **`fasthttp`:** Performance-oriented. Minimizes allocations (reuses objects). Faster for high-throughput plain HTTP/1.1. API is different (Request/Response context object). **Downsides:** No native HTTP/2, harder to use standard middleware.

---

### Question 413: How do you throttle network traffic in Go?

**Answer:**
Use `golang.org/x/time/rate` (Token Bucket algorithm).

```go
limiter := rate.NewLimiter(10, 1) // 10 req/sec

func handler(w http.ResponseWriter, r *http.Request) {
    if !limiter.Allow() {
        http.Error(w, "Too Many Requests", 429)
        return
    }
    // serve
}
```

---

### Question 414: How would you analyze network latency in Go?

**Answer:**
Use `net/http/httptrace`. It allows hooking into low-level HTTP lifecycle events.

```go
trace := &httptrace.ClientTrace{
    GotConn: func(info httptrace.GotConnInfo) {
        fmt.Printf("Reused: %v, IdleTime: %v\n", info.Reused, info.IdleTime)
    },
    DNSDone: func(info httptrace.DNSDoneInfo) {
        fmt.Println("DNS Limit:", info.Coalesced)
    },
}
req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
```

---

### Question 415: How would you implement WebRTC or peer-to-peer comms?

**Answer:**
Use the library **Pion WebRTC** (`github.com/pion/webrtc`).
It is a pure Go implementation of the WebRTC API.

Steps:
1.  Create `PeerConnection`.
2.  Exchange SDP offers/answers via a signaling server (WebSocket/HTTP).
3.  Establish P2P connection (ICE candidates).
4.  Send media/data via data channels.

---

### Question 416: How do you simulate a slow network in integration tests?

**Answer:**
1.  **Custom Transport:** Sleep in `RoundTrip`.
2.  **Toxiproxy:** An external tool (Go-based) to simulate network conditions (latency, reset, timeout) between services.
3.  **Local Middleware:**

```go
func SlowMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(2 * time.Second) // Simulate lag
        next.ServeHTTP(w, r)
    })
}
```

---

### Question 417: What’s the difference between connection pooling and multiplexing?

**Answer:**
- **Pooling (HTTP/1.1):** Maintain multiple open TCP connections. Request 2 waits for Request 1 to finish on Connection A, or uses Connection B.
- **Multiplexing (HTTP/2):** Use **ONE** TCP connection. Divide connection into "streams". Request 1 and Request 2 are sent as interleaved frames simultaneously over the same wire. No Head-of-Line blocking (mostly).

---

### Question 418: How do you verify DNS lookups in Go?

**Answer:**
Use `net.LookupHost` or `net.Resolver` for more control.

```go
ips, err := net.LookupHost("google.com")
for _, ip := range ips {
    fmt.Println(ip)
}

// Custom DNS Server
r := &net.Resolver{
    PreferGo: true,
    Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
        return net.Dial("udp", "8.8.8.8:53")
    },
}
```

---

### Question 419: How do you use HTTP pipelining in Go?

**Answer:**
Go's `net/http` client does **not** support HTTP/1.1 pipelining (sending multiple requests without waiting for responses) because it is fragile and often disabled on proxies/browsers.
Instead, Go relies on **HTTP/2 multiplexing**, which solves the same problem (latency) much better without the ordering issues of pipelining.

---

### Question 420: How do you implement NAT traversal in Go?

**Answer:**
1.  **UPnP:** Use libraries like `goupnp` to ask the router to open ports.
2.  **STUN/TURN:** Use `pion/stun` or `pion/turn` to discover public IP or relay traffic if P2P fails.
3.  **Hole Punching:** Attempting to connect from both sides simultaneously (usually managed by ICE in WebRTC).

---
