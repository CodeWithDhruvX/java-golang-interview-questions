# 📡 **701–720: Network & Protocol-Level Programming**

### 701. How do you parse HTTP headers manually in Go?
"I use `textproto.MIMEHeader`.
Read line by line until `\r\n\r\n`.
Split on `:`.
`key := canonicalMIMEHeaderKey(parts[0])`.
But really, I use `http.ReadRequest`. Manual parsing is useful only for custom non-compliant HTTP-like protocols (e.g., SIP or RTSP)."

#### Indepth
**Zero-Allocation Parsing**. `fasthttp` parses headers without allocating strings. It returns raw byte slices `[]byte` pointing to the original buffer. This saves millions of allocations under load but makes usage dangerous (slices are invalid after the request handler returns).

---

### 702. How do you handle fragmented UDP packets in Go?
"UDP doesn't handle fragmentation; IP does.
But if I mean *application* fragmentation (message > MTU):
I implement it myself.
Header: `[MsgID][SeqNum][TotalParts]`.
Receiver buffers parts and reassembles when all Sequence Numbers arrive. This allows sending 64KB images over UDP."

#### Indepth
**IP Fragmentation is Evil**. Relying on IP-level fragmentation is bad because if *one* fragment is lost, the entire packet is dropped by the kernel. Firewalls also drop fragments. Always fragment at the Application Layer (keep UDP payloads < 1400 bytes to stay under standard MTU).

---

### 703. How do you implement a custom binary protocol in Go?
"**Framing** is key.
My format: `[Length: 4b][MsgType: 1b][Payload: N]`.
Reader:
`binary.Read(conn, BigEndian, &length)`.
`payload := make([]byte, length)`.
`io.ReadFull(conn, payload)`.
This loop ensures I read exactly one message, handling TCP's streaming nature correctly."

#### Indepth
**Delimiters**. An alternative to Length-Prefix is a Delimiter (like `\n` in text protocols). `bufio.Scanner` is perfect for this. However, you must handle "escaping" if the delimiter appears in the payload. Length-Prefix is generally robust and simpler for binary data.

---

### 704. How do you parse and encode protobufs manually?
"I use `google.golang.org/protobuf/encoding/protowire`.
It exposes low-level types: `Varint`, `Fixed32`, `Bytes`.
`for len(buf) > 0 { num, typ, n := protowire.ConsumeTag(buf); ... }`.
It allows me to scan a protobuf message without the `.pb.go` struct definition, useful for generic tools like 'proto-dump'."

#### Indepth
**ZigZag Encoding**. Protobuf uses ZigZag encoding for signed integers. Normal int `-1` is `0xFFFFFFFF` (huge varint). ZigZag maps `-1` to `1`, `1` to `2`, `-2` to `3`... keeping small negative numbers small on the wire. `protowire` handles this decoding for you.

---

### 705. How do you build a TCP proxy in Go?
"I pipe two connections.
`go io.Copy(dest, src)`
`go io.Copy(src, dest)`
I wait for one to finish, then close both.
I verified to handle timeouts and half-closures (`CloseWrite`).
This is the core of HAProxy or Nginx TCP streams, and Go does it in 10 lines."

#### Indepth
**Splice**. On Linux, `io.Copy` between two TCP connections uses the `splice` syscall. This moves data from one socket buffer to another entirely within the kernel, zero user-space copies. This allows Go proxies to push 10Gbps+ with minimal CPU usage.

---

### 706. How do you implement a reverse proxy in Go?
"I use `httputil.NewSingleHostReverseProxy(target)`.
I customize the `Director` to modify the request (change Host header).
I customize `ModifyResponse` to rewrite headers (e.g., CORS).
It handles connection pooling and error reporting automatically."

#### Indepth
**Buffering Issues**. `ReverseProxy` buffers the response body from the backend before sending it to the client? No, it streams by default. But it *does* buffer the request if you use retries. Be careful with large request bodies + retries, as it might buffer the upload in RAM.

---

### 707. How does Go handle IPv4 vs IPv6 dual-stack?
"`net.Listen("tcp", ":80")` listens on both (if OS supports it).
Go uses **dual-stack** socket by default.
To force IPv4: `net.Listen("tcp4", ...)`.
To force IPv6: `net.Listen("tcp6", ...)`.
I check `conn.RemoteAddr()` to see if the client connected via `[::1]` or `127.0.0.1`."

#### Indepth
**Happy Eyeballs**. When *dialing*, Go uses RFC 6555 (Happy Eyeballs). It resolves both IPv4 and IPv6. It sends a SYN to IPv6. If it doesn't reply in 300ms, it sends a SYN to IPv4. Whichever connects first wins. This prevents broken IPv6 setups from hanging your app.

---

### 708. How do you implement a SOCKS5 server in Go?
"Handshake:
1.  Read `VER`, `NMETHODS`.
2.  Reply.
3.  Read Request: `CMD=Connect`, `ADDR`.
4.  `target := net.Dial(...)`.
5.  Pipe `client <-> target`.
I use it to bypass firewalls or implement VPN-like functionality purely in app-space."

#### Indepth
**UDP Associate**. SOCKS5 also supports UDP. The client asks "Please forward UDP for me". The server binds a UDP port and tells the client "Send packets here". Implementing this is tricky because you have to manage a dynamic mapping of UDP flows and timeouts.

---

### 709. How do you perform raw socket programming in Go?
"I use `syscall` or `golang.org/x/net/ipv4`.
`c, _ := ipv4.NewPacketConn(conn)`.
`c.SetControlMessage(ipv4.FlagTTL, true)`.
This allows me to read/write IP headers directly. Use case: implementing `ping` (ICMP) or `traceroute` without calling the external binary."

#### Indepth
**BPF Filters**. When reading raw sockets, you see *everything*. To filter, use BPF (Berkeley Packet Filter). `ipv4.PacketConn` allows attaching a BPF program to the socket in the kernel. The kernel then discards unwanted packets before they even wake up your Go program.

---

### 710. How do you handle TCP Half-Close in Go?
"TCP allows closing the *write* side while keeping *read* open.
`conn.(*net.TCPConn).CloseWrite()`.
This sends `FIN`.
The remote sees EOF on read but can still send data back.
I use this for pipeline protocols where I send a stream of data, close write to signal 'done', and then wait for the result."

#### Indepth
**HTTP/1.0**. This was the default mode in HTTP/1.0. Client sends Request, CloseWrite. Server sends Response, Close. `CloseWrite` is distinct from `Close`. `Close` tears down the socket entirely. `CloseWrite` sends a FIN packet but keeps the socket available for reading incoming data.

---

### 711. How do you implement Zero-Copy networking in Go?
"I use `syscall.Sendfile` (via `io.Copy` on `*os.File` -> `*net.TCPConn`).
Go optimizes `io.Copy` to use `sendfile` automatically if possible.
This copies data from Disk Cache to NIC buffer without passing through User Space RAM, saving massive CPU on file servers."

#### Indepth
**Limitations**. `Sendfile` only works if the source is a file (mmap-able) and dest is a socket. It doesn't work if you need to *encrypt* (TLS) or *compress* (GZIP) the data, because the CPU needs to see the data to transform it. For TLS, Look into Kernel TLS (kTLS) support in newer Linux/Go versions.

---

### 712. How do you tune TCP buffer sizes in Go?
"`conn.SetReadBuffer(bytes)`.
`conn.SetWriteBuffer(bytes)`.
I increase this for high-latency, high-bandwidth links (Long Fat Networks) to fill the BDP (Bandwidth-Delay Product).
Default is usually 4KB-16KB, which bottlenecks a 10Gbps link."

#### Indepth
**Autotuning**. Modern Linux kernels (4.x+) autotune TCP buffers dynamically (`tcp_moderate_rcvbuf`). Manually setting `SetReadBuffer` disables this autotuning. Only touch this if you are smarter than the Kernel (e.g., highly specific satellite link scenarios). Usually, default is best.

---

### 713. How to create a broadcast UDP server?
"I dial the broadcast address.
`conn, _ := net.Dial("udp", "255.255.255.255:9999")`.
Or Multicast: `net.ListenMulticastUDP(...)`.
I verify the network allows it (most cloud VPCs block broadcast/multicast). It works great on LAN."

#### Indepth
**Interfaces**. When multicasting, you must specify *which* interface to use. `net.ListenMulticastUDP("udp", iface, addr)`. The OS routing table decides where to send valid unicast, but for multicast, you often need to be explicit (`eth0` vs `wlan0`).

---

### 714. How do you implement a custom DNS server in Go?
"I use `miekg/dns` (the gold standard).
`dns.HandleFunc("example.com.", handleRequest)`.
In handler: construct `dns.Msg` response with `A` record.
CoreDNS (Kubernetes DNS) is built entirely on this library."

#### Indepth
**EDNS0**. Modern DNS requires EDNS0 to support larger packets (> 512 bytes) and security extensions (DNSSEC). `miekg/dns` handles `SetEdns0` automatically. If you write a raw UDP DNS server, packets larger than 512 bytes will be truncated, breaking many modern lookups.

---

### 715. How do you parse URL parameters manually without `net/url`?
"I shouldn't. `net/url` handles encoding (`%20`), escaping, and edge cases.
But if I must: `strings.Split(url, "?")`, then `strings.Split(params, "&")`.
Security risk: I'll likely forget to unescape values, leading to bugs. Always use standard lib."

#### Indepth
**Query vs Path**. parameters in path (`/users/12%2F34`) are tricky. Is it user `12/34` or user `12` sub-resource `34`? `net/url` parses RawPath to preserve the encoding differentiation. Manual string splitting destroys this distinction and opens up Path Traversal vulnerabilities.

---

### 716. What is the standard `io.Reader` and `io.Writer` interface?
"`Read(p []byte) (n int, err error)`.
`Write(p []byte) (n int, err error)`.
These are the universal abstractions.
My HTTP Handler takes `Body` (Reader). My Compressor takes `Writer`.
This composability allows me to chain `File -> Gzip -> Encrypt -> Net` just by wrapping interfaces, Unix-pipe style."

#### Indepth
**ReaderAt**. `io.Reader` is a stream (forward only). `io.ReaderAt` allows random access (`ReadAt(p, offset)`). This is required for downloading multiple chunks of a file in parallel (HTTP Range requests). `os.File` implements both. `net.Conn` only implements Reader.

---

### 717. How do you implement a QUIC server in Go?
"I use `quic-go`.
`listener, _ := quic.ListenAddr(addr, tlsConfig, nil)`.
It runs over UDP.
It provides streams: `stream, _ := sess.AcceptStream()`.
It feels like TCP (reliable streams) but without Head-of-Line blocking. HTTP/3 is built on top of this."

#### Indepth
**UDP Tuning**. QUIC runs in user-space on UDP. The kernel's UDP buffer defaults are often too small (200KB) for high-speed QUIC transfers. You almost always need to utilize `sysctl -w net.core.rmem_max=2500000` (2.5MB) to get decent HTTP/3 performance.

---

### 718. How do you handle endianness in network protocols?
"Network Byte Order is **Big Endian**.
Go (on x86) is Little Endian.
I use `binary.BigEndian.PutUint32(buf, val)`.
Never cast a struct pointer to send it over wire (`unsafe`), because the receiver might represent integers differently."

#### Indepth
**Network Byte Order**. Historically, "Network Order" is Big Endian (Motorola style). The web (IP, TCP headers) is Big Endian. Most modern CPUs (x86, ARM) are Little Endian. You *always* swap bytes at the network boundary. `binary.Read` does this. Casting memory (`*int32`) does not.

---

### 719. How do you detect port availability in Go?
"I try to listen on it.
`ln, err := net.Listen("tcp", ":8080")`.
If `err != nil` (bind: address already in use), it's taken.
Ensure to `ln.Close()` immediately if I was just checking.
Race condition: It might be taken 1ms after I check. Always handle bind errors gracefully."

#### Indepth
**Port 0**. Asking the OS for a free port is best done by binding to port 0. `net.Listen("tcp", ":0")`. The OS picks an ephemeral port. You can retrieve it via `ln.Addr()`. This is standard practice for running parallel integration tests that start real servers.

---

### 720. How do you inspect TLS handshake details?
"I use `tls.Config{ VerifyConnection: func(cs tls.ConnectionState) error { ... } }`.
Or examine `conn.ConnectionState()` after handshake.
I can see the negotiated:
*   Cipher Suite (e.g., TLS_AES_128_GCM_SHA256).
*   Protocol Version (TLS 1.3).
*   Peer Certificates."

#### Indepth
**JA3 Fingerprinting**. TLS Client Hello packets have a unique signature (Order of ciphers, extensions, etc.). You can use `github.com/refraction-networking/utls` to "impersonate" Chrome/Firefox fingerprints. This is used by scrapers to avoid bot detection, or by firewalls to detect suspicious clients.
