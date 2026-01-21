## 📡 Network & Protocol-Level Programming (Questions 721-740)

### Question 721: How do you create a custom TCP server in Go?

**Answer:**
Using `net.Listen`. (See Question 401 for code).
Important concepts:
- `for { Accept() }` loop.
- `go handle(conn)` per connection.
- Define a protocol (e.g., Line-based `\n` or Length-prefixed) to read messages from the stream.

---

### Question 722: How do you parse HTTP headers manually in Go?

**Answer:**
If you can't use `net/http` and possess a raw `net.Conn`:
Use `textproto.NewReader(bufio.NewReader(conn))`.

```go
tp := textproto.NewReader(bufio.NewReader(conn))
line, _ := tp.ReadLine() // "GET / HTTP/1.1"
headers, _ := tp.ReadMIMEHeader()
fmt.Println(headers.Get("Host"))
```

---

### Question 723: How do you handle fragmented UDP packets in Go?

**Answer:**
UDP has a max size (MTU, often 1500 bytes). If sending large data, IP fragmentation occurs (bad) or you must implement app-level fragmentation.
**App-Level:**
1.  Split data into chunks (e.g., 1KB).
2.  Add header: `[SeqID, TotalChunks, ChunkIndex, Data]`.
3.  Receiver buffers chunks and reassembles when all indices arrive.

---

### Question 724: How do you implement a custom binary protocol in Go?

**Answer:**
Use `encoding/binary`.
Fixed Header (4 bytes Length + 1 byte Type) + Body.

```go
// Header
binary.Write(buf, binary.BigEndian, uint32(len(data)))
binary.Write(buf, binary.BigEndian, uint8(msgType))
// Body
buf.Write(data)
```
Reader reads 5 bytes first, determines length, then reads Body.

---

### Question 725: How do you parse and encode protobufs manually?

**Answer:**
Without `protoc` generated code (hard), you use `google.golang.org/protobuf/encoding/protowire`.
It provides low-level functions to parse tag-length-value (TLV) pairs from the byte stream.
Generally not recommended over generated code.

---

### Question 726: How do you build a TCP proxy in Go?

**Answer:**
Pipe two connections together using `io.Copy`.

```go
func handle(src net.Conn) {
    dst, _ := net.Dial("tcp", "backend:80")
    defer dst.Close()
    
    go io.Copy(dst, src) // Src -> Dst
    io.Copy(src, dst)    // Dst -> Src
}
```

---

### Question 727: How do you implement a reverse proxy in Go?

**Answer:**
Use `httputil.NewSingleHostReverseProxy`.
It handles Header forwarding (`X-Forwarded-For`), buffering, and connection pooling automatically.
(See Question 408).

---

### Question 728: How do you sniff packets using Go?

**Answer:**
Use `gopacket`.
It binds to the network interface in promiscuous mode (Admin required).
Useful for debugging network protocols or building IDS (Intrusion Detection Systems).

---

### Question 729: How do you build a SOCKS5 proxy in Go?

**Answer:**
SOCKS5 performs a handshake before relaying.
1.  Handshake: Client says "I support Auth methods...".
2.  Request: "Connect to google.com:80".
3.  Server dials google.com.
4.  Relay: Copy bytes between Client and Google.
Library: `github.com/armon/go-socks5`.

---

### Question 730: How do you write a raw socket listener in Go?

**Answer:**
Use `syscall.Socket`, `syscall.Bind`, `syscall.Listen`.
Or `net.ListenPacket("ip4:icmp", ...)` for ICMP (Ping).
Raw sockets require root privileges and allow crafting custom IP headers.

---

### Question 731: How do you implement an HTTP client with timeout handling?

**Answer:**
**Crucial:** Never use default `http.Get`. It has no timeout.

```go
client := &http.Client{
    Timeout: 10 * time.Second,
}
resp, err := client.Get(url)
```
This timeout covers connect + request + reading headers + reading body.

---

### Question 732: How do you use netpoll in high-performance Go networking?

**Answer:**
The Go runtime uses `netpoll` (epoll/kqueue) internally.
You don't call it directly.
You write standard blocking code (`conn.Read()`). The runtime registers the file descriptor with epoll and parks the goroutine. When data arrives, epoll wakes the runtime, which wakes the goroutine.
This gives sync-looking code the performance of async I/O.

---

### Question 733: How do you build a DNS resolver in Go?

**Answer:**
Use `github.com/miekg/dns`.
It is a heavy-duty DNS library (used by CoreDNS).
Can act as a Client (Query A record) or a Server (Answer queries).

---

### Question 734: How do you manage connection pooling in network services?

**Answer:**
Maintain a valid `Pool` of "Ready" connections.
On `Acquire()`: Check if Pool is empty. If yes, Dial new. If no, pop one. Verify it's still alive (SetReadDeadline + 1 byte peek, or ping).
On `Release()`: Push back to pool.

---

### Question 735: How do you detect dropped connections in TCP?

**Answer:**
Trying to write to it returns `syscall.EPIPE` (Broken Pipe).
Reading determines EOF.
For handling "Half-Open" connections (cable pulled), use **TCP Keepalive** (`conn.SetKeepAlive(true)`).

---

### Question 736: What’s the difference between persistent and non-persistent HTTP in Go?

**Answer:**
- **Persistent (Keep-Alive):** Default. Reuses TCP connection for multiple requests. Good for perf.
- **Non-Persistent:** `req.Close = true`. Connection closes after response. Good for infrequent requests or to free fd resources on servers.

---

### Question 737: How do you write a TLS server in Go from scratch?

**Answer:**
```go
cert, _ := tls.LoadX509KeyPair("cert.pem", "key.pem")
cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
ln, _ := tls.Listen("tcp", ":443", cfg)

for {
    conn, _ := ln.Accept()
    // conn is automatically a TLS-wrapped connection
    go handle(conn)
}
```

---

### Question 738: How do you implement rate limiting per IP in a TCP server?

**Answer:**
Map `IP -> *RateLimiter`.
Use `sync.Map` for concurrency safety.
In `Accept()`:
```go
ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
limiter := getLimiter(ip)
if !limiter.Allow() {
    conn.Close()
    return
}
```
Clean up old entries periodically to prevent memory leaks.

---

### Question 739: How do you use Go to test API latency?

**Answer:**
`net/http/httptrace`. (See Question 414).
Measure `ConnectStart` to `ConnectDone` (TCP Handshake time), and `GotFirstResponseByte` (TTFB).

---

### Question 740: How do you monitor and log TCP connections?

**Answer:**
Wrap the `net.Listener`.

```go
type LogListener struct { net.Listener }
func (l LogListener) Accept() (net.Conn, error) {
    c, err := l.Listener.Accept()
    log.Println("New conn from", c.RemoteAddr())
    return c, err
}
```

---
