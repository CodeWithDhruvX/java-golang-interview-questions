# 📡 Go Theory Questions: 721–740 Network & Protocol-Level Programming

## 721. How do you parse HTTP headers manually in Go?

**Answer:**
Raw TCP reading (`net.Conn`).
1.  Read until `\r\n\r\n` (End of Header).
2.  Split by `\r\n` (Lines).
3.  First line: `GET / HTTP/1.1`.
4.  Subsequent lines: `Key: Value`. Split by first `:`.
We typically use `textproto.Reader`, which handles the MIME header normalization (multi-line headers) automatically.

---

## 722. How do you handle fragmented UDP packets in Go?

**Answer:**
UDP has no "fragments" at the app level. You get a datagram.
If the sender sent a packet larger than MTU (1500 bytes), IP fragmented it. The OS reassembles it.
If the OS fails (packet loss), we get nothing.
If we implement a protocol on top of UDP (like QUIC), we implement fragmentation manually:
Header: `[SeqID][FragID]`.
We read small packets into a buffer `map[SeqID][]byte` and reassemble when we satisfy the length.

---

## 723. How do you implement a custom binary protocol in Go?

**Answer:**
We use `encoding/binary`.
Fixed Header approach:
`[Length (4 bytes)] [MsgType (1 byte)] [Payload (Length bytes)]`.

Reader:
```go
var length uint32
binary.Read(conn, binary.BigEndian, &length)
payload := make([]byte, length)
io.ReadFull(conn, payload)
```
This framing allows TCP (a stream) to carry discrete messages reliably.

---

## 724. How do you parse and encode protobufs manually?

**Answer:**
Protobuf wire format is Key-Value pairs.
Key = FieldNumber << 3 | WireType.
Varint encoding (base 128).

To manual parse:
Read byte. Check MSB (Most Significant Bit). If 1, read next byte. Combine 7-bit chunks. This gives the Tag.
If WireType=2 (Length Delimited), read length varint, then read bytes.
We rarely do this manual work; we use `google.golang.org/protobuf/proto`.

---

## 725. How do you build a TCP proxy in Go?

**Answer:**
It's a byte pipe.
1.  Accept Conn `src`.
2.  Dial Backend `dst`.
3.  Spawn two goroutines:
    a. `io.Copy(dst, src)`
    b. `io.Copy(src, dst)`
4.  Wait for one to close, then close the other.
We must handle the `CloseWrite` (TCP Half-Close) correctly to allow graceful termination of streams.

---

## 726. How do you implement a reverse proxy in Go?

**Answer:**
The standard library `httputil.ReverseProxy` is robust.

```go
target, _ := url.Parse("http://backend:8080")
proxy := httputil.NewSingleHostReverseProxy(target)
proxy.ServeHTTP(w, r)
```
Be careful with the `Director` function: it modifies the request. You must ensure you clear `RequestURI` (which is client specific) so the client lib re-generates it for the backend target.

---

## 727. How do you sniff packets using Go?

**Answer:**
We use **GoPacket** (wrapper around `libpcap`).
Requires CGO and `npcap` installed.

```go
handle, _ := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)
packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
for packet := range packetSource.Packets() {
    fmt.Println(packet.String())
}
```
This captures raw frames (Ethernet/IP/TCP) allowing deep network analysis or building firewalls.

---

## 728. How do you build a SOCKS5 proxy in Go?

**Answer:**
SOCKS5 Handshake:
1.  Client: `[VER 05] [NMETHODS] [METHODS...]`.
2.  Server: `[VER 05] [METHOD 00]`.
3.  Client: `[VER 05] [CMD 01] [RSV] [ATYP] [DST.ADDR] [DST.PORT]`.
4.  Server: Dials DST. Returns Success.
5.  Pipe Start (See Q 725).
It differs from HTTP Proxy because it operates at formatting TCP level, allowing any traffic (SSH, DB) to pass through.

---

## 729. How do you write a raw socket listener in Go?

**Answer:**
Requires `syscall` or `golang.org/x/net/ipv4`.
`conn, _ := net.ListenPacket("ip4:tcp", "0.0.0.0")`
This gives access to the IP payload (the TCP Header).
Note: You need `CAP_NET_RAW` (Root privileges) to open raw sockets. We use this for tools like `ping` (ICMP) or custom scanners.

---

## 730. How do you implement an HTTP client with timeout handling?

**Answer:**
**Never** use generic `http.Get`. It has no timeout.
Always use `http.Client`.

```go
c := &http.Client{
    Timeout: 10 * time.Second,
}
```
This timeout includes Dial, TLS Handshake, Sending Body, and Reading Headers.
For granular control (e.g., "Connect in 1s, but allow 10s for body"), we use `net.Dialer` with timeouts inside a `Transport`.

---

## 731. How do you use netpoll in high-performance Go networking?

**Answer:**
You use it implicitly by using `net.Conn`.
When you call `conn.Read()`, if no data is there, Go parks the goroutine and registers the file descriptor (FD) with Epoll/Kqueue.
When OS signals "Data Ready", the runtime wakes the goroutine.
Libraries like **gnet** or **gev** expose the event loop directly (Callback based) to avoid the overhead of Goroutines/Stacks for C10M (10 Million Connections) scenarios.

---

## 732. How do you build a DNS resolver in Go?

**Answer:**
DNS is a UDP packet protocol.
ID, Flags, Questions, Answers.
We can send port 53 UDP query using `miekg/dns` (the standard Go DNS lib).
`m := new(dns.Msg); m.SetQuestion("google.com.", dns.TypeA)`.
`c := new(dns.Client); in, _, _ := c.Exchange(m, "8.8.8.8:53")`.
Writing a Server is just mapping `dns.HandleFunc` to response logic.

---

## 733. How do you manage connection pooling in network services?

**Answer:**
In HTTP, `http.Transport` does it automatically.
`MaxIdleConns`, `MaxIdleConnsPerHost`.
For custom protocols (like DBs), we implement a pool:
Channel of connections `chan net.Conn`.
Get: `select { case c := <-pool: return c; default: return dial() }`.
Put: `select { case pool <- c: return; default: c.Close() }` (If pool is full, discard).

---

## 734. How do you detect dropped connections in TCP?

**Answer:**
TCP is silent. If the cable is cut, `Read` blocks forever.
1.  **Read Deadline**: `conn.SetReadDeadline(time.Now().Add(time.Minute))`. If no heartbeat, error.
2.  **KeepAlive**: `net.Dialer{KeepAlive: 30*time.Second}`. OS sends empty ACKs.
3.  **App Level Heartbeat**: Send `Ping`, expect `Pong` within 5s. This is the most reliable method.

---

## 735. How do you implement Zero-Copy networking in Go?

**Answer:**
Use `syscall.Sendfile` (via `io.Copy` on `TCPConn` and `File`).
Typical Copy: Disk -> Kernel -> User -> Kernel -> NIC.
Zero Copy: Disk -> Kernel -> NIC.
In Go, `TCPConn.ReadFrom` detects if the source is an `os.File` and automatically uses `sendfile` syscall, bypassing userspace memory entirely efficiently.

---

## 736. What is the Nagle Algorithm and when to disable it in Go?

**Answer:**
Nagle buffers small writes to make full packets.
Good for Bandwidth, Bad for Latency (Wait 200ms).
In Go: `conn.(*net.TCPConn).SetNoDelay(true)`.
This disables Nagle (sets `TCP_NODELAY`).
We do this for Real-Time games, SSH sessions, or request/response RPCs where we want the `Send` to go out on the wire *immediately*.

---

## 737. How do you handle TLS handshake manually in Go?

**Answer:**
`tls.Client(conn, config)`.
This wraps a raw TCP connection.
The Handshake itself usually happens on the first `Read` or `Write`.
To force it (checking for errors early): `tlsConn.Handshake()`.
We can inspect `tlsConn.ConnectionState()` to see the CipherSuite negotiated, PeerCertificates, and SNI info.

---

## 738. How do you parse and generate JSON over a socket efficiently?

**Answer:**
`json.Encoder` / `json.Decoder`.
They stream.
**Delimiting**: Multiple JSONs on one socket need delimiters.
Common: Newline Delimited JSON (NDJSON).
`dec := json.NewDecoder(conn)`
`for dec.More() { dec.Decode(&msg) }`.
The decoder is smart enough to stop parsing after a valid JSON object is closed `}` (usually), but adding a newline is safer for framing.

---

## 739. How do you implement a minimal HTTP server without `net/http`?

**Answer:**
1.  `net.Listen("tcp", ":80")`.
2.  `conn.Accept()`.
3.  Read line: `GET / HTTP/1.1`.
4.  Write:
```text
HTTP/1.1 200 OK\r\n
Content-Length: 13\r\n
\r\n
Hello World
```
This is useful for understanding the protocol, but in production, we always use `net/http` because it covers edge cases (Chunked encoding, Keep-Alives, Headers) that are painful to write from scratch.

---

## 740. How do you implement port scanning concurrently?

**Answer:**
Target: `192.168.1.1` Ports 1-65535.
Common Go pattern:
1.  `jobs` channel (Port numbers).
2.  `results` channel.
3.  Spawn 1000 workers.
4.  Worker: `net.DialTimeout("tcp", fmt.Sprintf(ip, port), 1*time.Second)`.
5.  If success, send to results.
This scans the entire port range in seconds due to Go's ability to handle high concurrency IO.
