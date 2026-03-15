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

### Explanation
Creating a TCP server in Go involves using the `net` package to listen on a specific port and accept incoming connections. The server runs in an infinite loop, accepting connections and handling each one in a separate goroutine to enable concurrent processing. Each connection is handled independently, allowing the server to serve multiple clients simultaneously.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create a TCP server in Go?
**Your Response:** "I create a TCP server in Go using the `net` package. First, I use `net.Listen()` to bind to a specific port like ':8080'. Then I run an infinite loop calling `listener.Accept()` to wait for incoming connections. When a connection arrives, I handle it in a separate goroutine to enable concurrent processing. Each connection gets its own `net.Conn` object that I can read from and write to. I make sure to defer `conn.Close()` to properly clean up resources. This pattern allows the server to handle multiple clients simultaneously without blocking on individual connections."

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

### Explanation
UDP is a connectionless protocol, so the implementation differs from TCP. For clients, `net.Dial("udp", address)` creates a connection that can send packets to the specified address. For servers, `net.ListenPacket` creates a packet conn that can receive UDP datagrams from multiple clients without establishing persistent connections.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create a UDP client in Go?
**Your Response:** "I create a UDP client in Go using `net.Dial("udp", address)` which establishes a connection to a UDP server. Unlike TCP, UDP is connectionless, so this just sets up the default destination for my packets. I can then use `conn.Write()` to send data to the server. For a UDP server, I'd use `net.ListenPacket("udp", ":8080")` to create a connection that can receive packets from any client. UDP doesn't have the concept of persistent connections like TCP, so each packet is independent and can be lost without affecting the protocol."

---

### Question 403: What is the difference between `net.Listen` and `net.Dial`?

**Answer:**
- **`net.Listen` (Server Side):** Binds to a local port and waits for incoming connections. Returns a `Listener`.
- **`net.Dial` (Client Side):** Initiates a connection to a remote (or local) address and port. Returns a `Conn`.

### Explanation
`net.Listen` and `net.Dial` serve opposite purposes in network programming. `net.Listen` is used by servers to bind to a port and wait for incoming connections, returning a Listener object that can accept connections. `net.Dial` is used by clients to actively connect to a server, returning a Connection object that can be used for communication.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `net.Listen` and `net.Dial`?
**Your Response:** "`net.Listen` and `net.Dial` serve opposite purposes in Go networking. `net.Listen` is used on the server side to bind to a local port and wait for incoming connections. It returns a Listener object that I can call `Accept()` on to get individual connections. `net.Dial` is used on the client side to actively initiate a connection to a remote server. It returns a Connection object that's ready for reading and writing. So servers listen and accept, while clients dial and connect. It's like the difference between waiting for a phone call versus making one."

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

### Explanation
Connection pooling reuses existing connections instead of creating new ones for each request, reducing overhead. Standard libraries like `database/sql` and `net/http` have built-in connection pooling. For custom TCP connections, you can implement a pool using a buffered channel as a queue, with Get() retrieving connections and Put() returning them, creating new connections only when the pool is empty.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage TCP connection pools?
**Your Response:** "I manage TCP connection pools in two ways. For standard libraries like `database/sql` and `net/http`, connection pooling is built-in and I configure parameters like `MaxIdleConns` and `MaxOpenConns`. For custom TCP connections, I implement my own pool using a buffered channel. I create a Pool struct with a channel of connections, where Get() retrieves from the channel or creates a new connection if empty, and Put() returns connections to the pool or closes them if the pool is full. This reduces the overhead of establishing new connections for each request and improves performance by reusing existing connections."

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

### Explanation
A custom HTTP transport implements the `http.RoundTripper` interface to intercept and modify HTTP requests before they're sent. This is useful for adding cross-cutting concerns like logging, authentication, caching, or retry logic. The transport wraps another transport (usually `http.DefaultTransport`) and delegates the actual request execution after applying its custom behavior.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you implement a custom HTTP transport?
**Your Response:** "I implement a custom HTTP transport by creating a struct that implements the `http.RoundTripper` interface, specifically the `RoundTrip` method. This method receives an HTTP request, allows me to modify or log it, then delegates to another transport - usually `http.DefaultTransport`. I use this pattern for adding cross-cutting concerns like request logging, authentication headers, caching, or retry logic. For example, I might create a LoggingTransport that prints request URLs before delegating to the underlying transport. I then configure an `http.Client` to use my custom transport instead of the default one."

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

### Explanation
Reading raw packets requires low-level network access typically provided by libpcap on Unix systems. The `gopacket` library provides a Go wrapper around libpcap, allowing you to capture and analyze network packets at the packet level. This requires elevated privileges since you're accessing the network interface directly.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you read raw packets using gopacket?
**Your Response:** "I read raw packets using the `github.com/google/gopacket` library, which is a Go wrapper around libpcap. First, I open a live handle to a network interface using `pcap.OpenLive()` with the interface name, snapshot length, and blocking mode. Then I create a packet source from this handle and iterate over the packets channel. Each packet contains the raw bytes and parsed layers for different protocols. This approach requires root or admin privileges since I'm accessing the network interface directly. I use this for network monitoring, packet analysis, or building custom network tools that need to inspect traffic at the packet level."

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

### Explanation
Connection hijacking in HTTP allows a handler to take complete control of the underlying TCP connection, bypassing the HTTP server's connection management. This is essential for protocols like WebSockets that need to switch from HTTP to a different communication protocol. The hijacker interface provides access to the raw connection and buffered reader/writer.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a connection hijack in `net/http` and how is it done?
**Your Response:** "Connection hijacking in Go's HTTP package allows me to take over the underlying TCP connection from the HTTP server. This is necessary when implementing protocols like WebSockets that start as HTTP requests but then switch to a different communication protocol. I check if the ResponseWriter implements the `http.Hijacker` interface, then call `Hijack()` to get the raw connection and buffered reader/writer. Once hijacked, the HTTP server no longer manages this connection - I'm responsible for reading, writing, and closing it. This gives me full control to implement custom protocols or long-lived connections beyond HTTP's request-response model."

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

### Explanation
A reverse proxy sits in front of backend services and forwards requests to them. Go's `httputil.ReverseProxy` provides a built-in implementation that handles request forwarding, response copying, and error handling. For forward proxies (where clients explicitly route through the proxy), you'd need to implement custom logic to parse requests and dial the target servers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to implement a proxy server in Go?
**Your Response:** "I implement a reverse proxy server in Go using the `httputil.ReverseProxy` from the standard library. I parse the target service URL, create a reverse proxy with `NewSingleHostReverseProxy()`, and then have it listen on a port. The proxy automatically handles forwarding requests to the backend service and returning responses. This is useful for load balancing, SSL termination, or routing requests to microservices. For a forward proxy where clients explicitly route through it, I'd need to implement custom logic to handle the request parsing and dial the target servers myself. The reverse proxy approach is simpler and more common for microservice architectures."

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

### Explanation
Go's HTTP server automatically enables HTTP/2 when using TLS, as the protocol requires TLS for standard deployments. The server and client negotiate HTTP/2 through ALPN during the TLS handshake. For cleartext HTTP/2 (H2C), you need to use the `h2c` package which provides a handler that upgrades HTTP/1.1 connections to HTTP/2 without TLS.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you create an HTTP2 server from scratch in Go?
**Your Response:** "Go's HTTP server actually enables HTTP/2 automatically when I use TLS with `ListenAndServeTLS`. The server and client negotiate HTTP/2 through ALPN during the TLS handshake. For most cases, I just need to serve over HTTPS and HTTP/2 works automatically. If I need cleartext HTTP/2 (H2C) for development or specific use cases, I use the `golang.org/x/net/http2/h2c` package. I create an HTTP/2 server and wrap my handler with `h2c.NewHandler()`. This allows HTTP/2 without TLS, though it's not recommended for production. The beauty is that Go handles most of the complexity automatically - I just focus on writing my HTTP handlers."

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

### Explanation
HTTP connection reuse (Keep-Alive) maintains TCP connections between requests to avoid the overhead of establishing new connections. Go's default HTTP transport implements connection pooling, keeping idle connections open for a configurable time period. This significantly improves performance for multiple requests to the same host by eliminating TCP handshake and TLS negotiation overhead.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Go handle connection reuse (keep-alive)?
**Your Response:** "Go's HTTP client handles connection reuse automatically through connection pooling in the default `http.Transport`. It keeps idle TCP connections open to avoid the overhead of establishing new connections for subsequent requests to the same host. This eliminates TCP handshakes and TLS negotiations, significantly improving performance. I can configure this behavior with parameters like `MaxIdleConns` to control the pool size and `IdleConnTimeout` to set how long connections remain open. If needed, I can disable keep-alives entirely by setting `DisableKeepAlives: true` in the transport, though this is rarely recommended as it would hurt performance."

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

### Explanation
Socket timeouts in Go are implemented using deadlines on connection objects. Deadlines specify absolute time points after which I/O operations will fail with a timeout error. You can set separate deadlines for read and write operations, or a combined deadline for both. This prevents goroutines from blocking indefinitely on slow or unresponsive network connections.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you set timeouts on sockets in Go?
**Your Response:** "I set timeouts on sockets in Go using the deadline methods on `net.Conn` objects. I can call `SetDeadline()` to set a combined deadline for both read and write operations, or use `SetReadDeadline()` and `SetWriteDeadline()` individually. These methods take absolute time values, so I calculate them relative to the current time. If an I/O operation doesn't complete before the deadline, it returns a timeout error. This prevents my goroutines from blocking indefinitely on slow or unresponsive connections. It's essential for building robust network applications that can handle network issues gracefully."

---

### Question 412: What is the difference between `net/http` and `fasthttp`?

**Answer:**
- **`net/http`:** Standard library. General purpose, HTTP/1.1 & HTTP/2. Interface-based (`io.Reader` body). Safe and idiomatic. Allocates new objects per request.
- **`fasthttp`:** Performance-oriented. Minimizes allocations (reuses objects). Faster for high-throughput plain HTTP/1.1. API is different (Request/Response context object). **Downsides:** No native HTTP/2, harder to use standard middleware.

### Explanation
`net/http` is Go's standard HTTP library focused on safety, compatibility, and ease of use. `fasthttp` is a third-party library optimized for maximum performance by minimizing allocations and reusing objects. While `fasthttp` can be significantly faster for high-throughput HTTP/1.1 scenarios, it sacrifices some compatibility and doesn't support HTTP/2. The choice depends on whether performance or compatibility is the priority.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `net/http` and `fasthttp`?
**Your Response:** "`net/http` is Go's standard HTTP library that prioritizes safety, compatibility, and ease of use. It supports both HTTP/1.1 and HTTP/2, uses idiomatic Go interfaces like `io.Reader`, and allocates new objects per request which is safer but less performant. `fasthttp` is a third-party library optimized for maximum performance - it minimizes allocations by reusing objects and can be significantly faster for high-throughput HTTP/1.1 scenarios. However, it has a different API that's less idiomatic, doesn't support HTTP/2, and can be harder to integrate with standard middleware. I choose `net/http` for most applications unless I have extreme performance requirements where `fasthttp`'s trade-offs make sense."

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

### Explanation
Network traffic throttling controls the rate of requests to prevent overwhelming services. The token bucket algorithm, implemented in `golang.org/x/time/rate`, allows bursts of requests while limiting the overall rate. The limiter maintains a bucket of tokens that refill at a specified rate, with each request consuming a token. If the bucket is empty, requests are rejected.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you throttle network traffic in Go?
**Your Response:** "I throttle network traffic in Go using the `golang.org/x/time/rate` package which implements the token bucket algorithm. I create a limiter with a specific rate, like 10 requests per second, and optionally a burst size. For each request, I call `limiter.Allow()` to check if the request should be processed. If there are available tokens, the request proceeds; if not, I return a 429 Too Many Requests error. This approach allows for controlled bursts while maintaining an overall rate limit, which is perfect for protecting APIs from abuse or managing resource usage. The token bucket algorithm is more flexible than simple rate limiting because it handles natural traffic patterns better."

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

### Explanation
HTTP tracing in Go provides visibility into the low-level lifecycle of HTTP requests, including DNS resolution, connection establishment, TLS handshakes, and request/response timing. The `httptrace` package allows you to hook into specific events to measure latency and identify performance bottlenecks in the HTTP client stack.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you analyze network latency in Go?
**Your Response:** "I analyze network latency in Go using the `net/http/httptrace` package which provides hooks into the HTTP request lifecycle. I create a `ClientTrace` struct with callbacks for events like DNS resolution, connection establishment, TLS handshake, and request completion. Each callback receives timing information that I can log or aggregate to identify performance bottlenecks. For example, the `GotConn` callback tells me if a connection was reused and how long it was idle, while `DNSDone` provides DNS resolution timing. I attach this trace to the request context using `httptrace.WithClientTrace()`. This gives me detailed visibility into where time is being spent in HTTP requests, helping me optimize performance."

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

### Explanation
WebRTC enables peer-to-peer communication between browsers without requiring intermediary servers. In Go, the Pion WebRTC library provides a pure Go implementation of the WebRTC API. The process involves creating peer connections, exchanging session descriptions through a signaling server, gathering ICE candidates for NAT traversal, and establishing direct communication channels.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you implement WebRTC or peer-to-peer comms?
**Your Response:** "I implement WebRTC in Go using the Pion WebRTC library, which is a pure Go implementation of the WebRTC API. The process involves several steps: first, I create a `PeerConnection` object on each peer. Then I exchange SDP offers and answers through a signaling server - this could be WebSocket or HTTP. Next, I gather and exchange ICE candidates to handle NAT traversal. Once the peer-to-peer connection is established, I can send data through data channels or media through video/audio tracks. The signaling server is only needed for the initial setup - after that, communication happens directly between peers. This approach enables real-time communication with low latency without requiring media servers for every connection."

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

### Explanation
Simulating slow network conditions in integration tests helps verify application resilience and timeout handling. Approaches include implementing a custom HTTP transport that adds delays, using external tools like Toxiproxy that can manipulate network conditions, or adding middleware that introduces artificial latency to test timeout and retry logic.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you simulate a slow network in integration tests?
**Your Response:** "I simulate slow network conditions in integration tests using several approaches. The simplest is creating a custom HTTP transport that adds sleep delays in the `RoundTrip` method. For more sophisticated testing, I use Toxiproxy, a Go-based tool that can simulate various network conditions like latency, connection resets, and timeouts between services. I can also add middleware that introduces artificial delays to test timeout and retry logic. These approaches help me verify that my application handles network issues gracefully - that timeouts work correctly, retry mechanisms function as expected, and the application remains responsive even when downstream services are slow. This is crucial for building robust distributed systems."

---

### Question 417: What’s the difference between connection pooling and multiplexing?

**Answer:**
- **Pooling (HTTP/1.1):** Maintain multiple open TCP connections. Request 2 waits for Request 1 to finish on Connection A, or uses Connection B.
- **Multiplexing (HTTP/2):** Use **ONE** TCP connection. Divide connection into "streams". Request 1 and Request 2 are sent as interleaved frames simultaneously over the same wire. No Head-of-Line blocking (mostly).

### Explanation
Connection pooling and multiplexing are different approaches to handling multiple HTTP requests. Connection pooling (HTTP/1.1) maintains multiple TCP connections, each handling one request at a time. Multiplexing (HTTP/2) uses a single TCP connection divided into multiple streams, allowing requests to be sent concurrently without head-of-line blocking, where one slow request doesn't block others.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between connection pooling and multiplexing?
**Your Response:** "Connection pooling and multiplexing are two different approaches to handling multiple HTTP requests. Connection pooling, used in HTTP/1.1, maintains multiple open TCP connections and assigns requests to different connections. If one connection is busy with a slow request, other requests can use different connections. Multiplexing, used in HTTP/2, uses a single TCP connection but divides it into multiple streams. Requests are sent as interleaved frames simultaneously, so a slow request doesn't block others on the same connection. This eliminates head-of-line blocking and is more efficient since it uses just one TCP connection instead of multiple, reducing overhead while maintaining better concurrency."

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

### Explanation
DNS lookups in Go can be performed using simple functions like `net.LookupHost` for basic queries, or with more control using the `net.Resolver` struct. The resolver allows customization of DNS servers, timeout settings, and whether to use the pure Go resolver or the system's cgo-based resolver.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you verify DNS lookups in Go?
**Your Response:** "I perform DNS lookups in Go using either simple functions like `net.LookupHost` for basic queries, or the more flexible `net.Resolver` for advanced control. For simple cases, I call `net.LookupHost()` with a domain name to get all IP addresses. When I need more control, I create a custom `Resolver` with specific DNS servers, timeouts, and other settings. I can configure it to use specific DNS servers like Google's 8.8.8.8, or force it to use the pure Go resolver instead of the system's native resolver. This flexibility is useful when I need to test DNS configurations or work around network restrictions."

---

### Question 419: How do you use HTTP pipelining in Go?

**Answer:**
Go's `net/http` client does **not** support HTTP/1.1 pipelining (sending multiple requests without waiting for responses) because it is fragile and often disabled on proxies/browsers.
Instead, Go relies on **HTTP/2 multiplexing**, which solves the same problem (latency) much better without the ordering issues of pipelining.

### Explanation
HTTP pipelining allows sending multiple requests over a single connection without waiting for responses, but it's problematic because responses must be returned in the same order as requests (head-of-line blocking). Go's HTTP client doesn't implement pipelining due to these issues. Instead, Go uses HTTP/2 multiplexing which allows concurrent requests and responses on a single connection without ordering constraints.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use HTTP pipelining in Go?
**Your Response:** "Actually, Go's `net/http` client doesn't support HTTP/1.1 pipelining because it's fragile and often disabled on proxies and browsers. Pipelining has issues with head-of-line blocking where responses must come back in the same order as requests. Instead of implementing pipelining, Go relies on HTTP/2 multiplexing which solves the same latency problem much better. HTTP/2 allows multiple requests and responses to be sent concurrently over a single connection without ordering constraints. This approach is more robust and efficient than pipelining. If I need to reduce latency in Go, I focus on using HTTP/2 or connection pooling rather than trying to implement pipelining."

---

### Question 420: How do you implement NAT traversal in Go?

**Answer:**
1.  **UPnP:** Use libraries like `goupnp` to ask the router to open ports.
2.  **STUN/TURN:** Use `pion/stun` or `pion/turn` to discover public IP or relay traffic if P2P fails.
3.  **Hole Punching:** Attempting to connect from both sides simultaneously (usually managed by ICE in WebRTC).

### Explanation
NAT traversal enables direct peer-to-peer connections when devices are behind different NAT routers. Techniques include UPnP for automatic port forwarding, STUN for discovering public IP addresses, TURN for relaying traffic when direct connection fails, and hole punching for simultaneous connection attempts. WebRTC's ICE framework combines these techniques automatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement NAT traversal in Go?
**Your Response:** "I implement NAT traversal in Go using several techniques depending on the scenario. For simple cases, I use UPnP libraries like `goupnp` to automatically ask routers to open ports. For more complex scenarios, I use STUN to discover the public IP address and TURN servers to relay traffic when direct connections fail. The most powerful approach is hole punching, where both sides attempt to connect simultaneously, usually managed by WebRTC's ICE framework. I often use the Pion libraries which provide pure Go implementations of STUN, TURN, and WebRTC. The choice depends on the network environment - UPnP works well on home networks, while STUN/TURN/hole punching is more reliable for internet-scale peer-to-peer connections."

---
