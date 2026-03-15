# 🟡 Networking & Communication Protocols — Questions 101–110

> **Level:** 🟢 Junior – 🔴 Senior (varies by depth)
> **Asked at:** Cloudflare, Amazon, Google, any backend/infra role

---

### 101. What is HTTP vs HTTPS?
"HTTP (Hypertext Transfer Protocol) is the foundation of data communication on the web — it defines how messages are formatted and transmitted between browsers and servers. HTTPS is simply HTTP with **TLS encryption** layered on top.

Without HTTPS, all data transmitted — including passwords, credit card numbers, session tokens — is in plaintext and can be read by anyone who can intercept the network traffic (your ISP, someone on the same WiFi, any middleman). This is called a man-in-the-middle attack.

For any real application in 2025, HTTPS is mandatory. Let's Encrypt provides free automated TLS certificates. There is no excuse to serve production traffic over plain HTTP."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** Any company — foundational web development knowledge

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is HTTP vs HTTPS?
**Your Response:** HTTP is the Hypertext Transfer Protocol that defines how messages are formatted and transmitted between browsers and servers. HTTPS is simply HTTP with TLS encryption layered on top. Without HTTPS, all data transmitted including passwords, credit card numbers, and session tokens is in plaintext and can be read by anyone who can intercept the network traffic - your ISP, someone on the same WiFi, or any middleman. This is called a man-in-the-middle attack. For any real application in 2025, HTTPS is mandatory. Let's Encrypt provides free automated TLS certificates, so there's no excuse to serve production traffic over plain HTTP.

#### Indepth
HTTPS/TLS key points:
- **Certificate validation:** Browser checks: is the certificate issued by a trusted CA? Is the domain name in the certificate? Is the certificate within its validity period? Has it been revoked (OCSP)?
- **HTTP/2 + TLS:** HTTP/2 is multiplexed (multiple requests over one connection) and is significantly faster than HTTP/1.1. Most browsers require HTTPS for HTTP/2.
- **HSTS (HTTP Strict Transport Security):** Response header `Strict-Transport-Security: max-age=31536000; includeSubDomains; preload` tells the browser to *always* use HTTPS for this domain, even if the user types `http://`. Browser maintains an HSTS list and upgrades connections preemptively.
- **Certificate transparency:** Publicly auditable log of all issued certificates. Protects against rogue CAs issuing unauthorized certificates for your domain.

---

### 102. What is REST vs GraphQL vs gRPC?
"REST, GraphQL, and gRPC are three different API paradigms, each with distinct trade-offs.

**REST** uses HTTP methods (GET/POST/PUT/DELETE) and URLs to define operations on resources. It's the most widely understood and supported — browser, mobile, and any language can call a REST API with zero tooling. The downside is over-fetching (endpoint returns more data than needed) and under-fetching (multiple round trips needed to assemble a complete response).

**GraphQL** allows clients to specify exactly what data they want in a query. One request can fetch nested data from multiple 'tables'. Perfect for complex client data needs — a mobile app and a web app can call the same endpoint and get different shaped responses. Overhead: complex query parsing, N+1 problem with naive resolvers.

**gRPC** uses protobuf binary serialization over HTTP/2. 5-10x smaller payload than JSON, strongly typed via protobuf schemas, and supports streaming. Best for internal service-to-service communication where performance and contract enforcement matter."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** API-design interviews at Razorpay, Swiggy, Amazon, Uber — choosing the right protocol for the right use case

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is REST vs GraphQL vs gRPC?
**Your Response:** REST, GraphQL, and gRPC are three different API paradigms with distinct trade-offs. REST uses HTTP methods and URLs to define operations on resources. It's the most widely understood and supported - any language can call a REST API with zero tooling. The downside is over-fetching where endpoints return more data than needed, and under-fetching where multiple round trips are needed. GraphQL allows clients to specify exactly what data they want in a query - one request can fetch nested data from multiple sources. It's perfect for complex client data needs. gRPC uses protobuf binary serialization over HTTP/2, giving 5-10x smaller payloads than JSON, strong typing via schemas, and supports streaming. It's best for internal service-to-service communication.

#### Indepth
| Feature | REST | GraphQL | gRPC |
|---|---|---|---|
| Transport | HTTP/1.1 or HTTP/2 | HTTP/1.1 or HTTP/2 | HTTP/2 only |
| Data Format | JSON | JSON | Protobuf (binary) |
| Schema | Optional (OpenAPI) | Mandatory (SDL) | Mandatory (.proto) |
| Versioning | URL versioning (`/v2/`) | Schema evolution | Proto field numbering |
| Type Safety | Weak (JSON is dynamic) | Strong (SDL types) | Strong (proto types) |
| Browser Support | Native | Needs client library | Needs gRPC-Web proxy |
| Streaming | Limited (SSE, WS separately) | Subscriptions | Native bi-directional streaming |
| Best For | Public APIs, CRUD services | Complex data fetching, BFF | Internal microservices |

**N+1 problem in GraphQL:** 100 posts each with an author field → 100 separate `SELECT * FROM users WHERE id=?` queries. Solution: **DataLoader** batches all user IDs and fetches in one `SELECT * FROM users WHERE id IN (...)` query. This is a mandatory pattern for any production GraphQL implementation.

---

### 103. What is WebSocket?
"WebSocket is a **full-duplex, persistent communication channel** over a single TCP connection. Unlike HTTP (request-response), WebSocket allows both client and server to send data at any time, in any direction.

The classic use cases: real-time chat (WhatsApp Web), live dashboards, multiplayer gaming, collaborative editing (Google Docs), financial price feeds. All of these require the server to push updates to the client without the client asking.

The connection starts as an HTTP handshake with an `Upgrade: websocket` header. Once upgraded, it's no longer HTTP — it's the WebSocket protocol. The connection stays open, and both sides can send messages (frames) at any time."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Swiggy/Zomato (live order tracking), Zerodha (live market prices), gaming companies, Slack/Discord (chat)

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is WebSocket?
**Your Response:** WebSocket is a full-duplex, persistent communication channel over a single TCP connection. Unlike HTTP which is request-response, WebSocket allows both client and server to send data at any time in any direction. The classic use cases are real-time chat like WhatsApp Web, live dashboards, multiplayer gaming, collaborative editing like Google Docs, and financial price feeds. All of these require the server to push updates to the client without the client asking. The connection starts as an HTTP handshake with an Upgrade: websocket header, and once upgraded, it's no longer HTTP - it's the WebSocket protocol where both sides can send messages at any time.

#### Indepth
WebSocket vs alternatives:

- **Short Polling:** Client repeatedly calls `GET /updates` every 1 second. Wasteful — most responses are 'no new data'. 1000 clients → 1000 requests/second just for polling.
- **Long Polling:** Client calls `GET /updates`. Server holds the connection open until there's new data (or timeout). Better than short polling. But still re-establishes a new HTTP connection after each response. Works through proxies and firewalls.
- **Server-Sent Events (SSE):** Server-to-client one-way streaming over HTTP. Client subscribes, server pushes events. Simpler than WebSocket for push-only use cases. Built on HTTP — works through load balancers without special config. Used for progress bars, live notification feeds.
- **WebSocket:** True bi-directional. Client and server both can initiate. Persistent connection. Requires special LB config (sticky sessions or L4 LB) since connection must stay on same server.

Scaling WebSocket servers: WebSocket connections are stateful — each user's connection is on a specific server. Horizontal scaling requires routing the same user's messages to their server. Solution: Redis Pub/Sub as message bus between WebSocket servers. A message for User A is published to Redis; the server where A is connected subscribes and delivers to A.

---

### 104. What is long polling?
"Long polling is a web communication technique where the **client makes an HTTP request and the server holds the connection open until new data is available**, then responds. The client immediately sends a new request after receiving a response, maintaining a continuous 'loop'.

It simulates server push without WebSocket. The user experience: events appear in near-real-time without the user refreshing the page. The server holds the connection for up to N seconds (typically 30-90 seconds). If no event occurs, it responds with an empty/timeout response.

For teams that can't use WebSockets (some corporate firewalls block WebSocket upgrades), long polling is a reliable fallback. Facebook's original chat system used long polling before upgrading to WebSocket."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any real-time communication design question — Swiggy live orders, notification systems

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is long polling?
**Your Response:** Long polling is a web communication technique where the client makes an HTTP request and the server holds the connection open until new data is available, then responds. The client immediately sends a new request after receiving a response, maintaining a continuous loop. It simulates server push without WebSocket. The user experience is that events appear in near-real-time without refreshing. The server holds the connection for up to 30-90 seconds, and if no event occurs, it responds with an empty or timeout response. For teams that can't use WebSockets due to corporate firewalls, long polling is a reliable fallback.

#### Indepth
Long polling implementation (Go server):
```go
func pollHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
    defer cancel()
    
    select {
    case event := <-eventChannel:  // New event available
        json.NewEncoder(w).Encode(event)
    case <-ctx.Done():  // 30s timeout
        json.NewEncoder(w).Encode(map[string]string{"status": "no_update"})
    case <-r.Context().Done():  // Client disconnected
        return
    }
}
```

Challenges:
- **Server resources:** Each open connection holds a goroutine/thread. At 100K concurrent users doing long polling → 100K concurrent connections. Go handles this gracefully (goroutines are cheap). Node.js (event loop) handles it well too. Java with blocking I/O struggles.
- **Load balancer timeout:** Default LB timeout is 60s. Long poll connections near 60s will be killed by the LB. Set LB idle timeout above your poll timeout.
- **Mobile networks:** Connections may drop silently. Client should reconnect on any error.

---

### 105. What is the difference between TCP and UDP?
"TCP and UDP are the two core transport layer protocols. They make opposite trade-offs between **reliability and speed**.

**TCP (Transmission Control Protocol)** is reliable: it guarantees delivery, ordering, and no duplication. It does this via handshaking, sequence numbers, acknowledgments, and retransmission of lost packets. This overhead makes it slightly slower but safe. Use TCP for everything where accuracy matters: APIs, databases, file transfer, email.

**UDP (User Datagram Protocol)** is 'fire and forget': it sends packets with no guarantee of delivery, order, or deduplication. But it's much faster — no connection setup, no acknowledgment overhead. Use UDP where speed matters more than perfection: live video streaming, online games, DNS, VoIP."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** Cloudflare, CDN/networking companies, gaming companies, any networking fundamentals discussion

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between TCP and UDP?
**Your Response:** TCP and UDP are the two core transport layer protocols that make opposite trade-offs between reliability and speed. TCP is reliable - it guarantees delivery, ordering, and no duplication through handshaking, sequence numbers, acknowledgments, and retransmission of lost packets. This overhead makes it slightly slower but safe. I use TCP for everything where accuracy matters like APIs, databases, file transfer, and email. UDP is 'fire and forget' - it sends packets with no guarantee of delivery, order, or deduplication. But it's much faster with no connection setup or acknowledgment overhead. I use UDP where speed matters more than perfection like live video streaming, online games, DNS, and VoIP.

#### Indepth
| Feature | TCP | UDP |
|---|---|---|
| Connection | Connection-oriented (3-way handshake) | Connectionless |
| Delivery | Guaranteed (retransmit on loss) | Not guaranteed |
| Ordering | Guaranteed (sequence numbers) | Not guaranteed |
| Duplicates | Eliminated | Possible |
| Overhead | Higher (headers, handshake, ACKs) | Lower (8-byte header vs 20+) |
| Speed | Slower | Faster |
| Use Cases | HTTP/S, SSH, DB, email | Video streaming, gaming, DNS, QUIC |

**QUIC protocol (HTTP/3's transport):** QUIC is built on UDP but adds reliability, ordering, and multiplexing at the QUIC layer. Key advantage over TCP: **no head-of-line blocking** — in TCP multiplexed streams, a lost packet blocks ALL streams until retransmitted. QUIC streams are independent — a lost packet only blocks its own stream. Also: zero-RTT connection resumption (reconnects are instant for known servers). QUIC is what makes HTTP/3 faster, especially on mobile with packet loss.

---

### 106. How does DNS work?
"DNS (Domain Name System) translates human-readable domain names (`google.com`) into IP addresses that computers use to route packets.

When you type `google.com` in your browser: (1) Your OS checks its **local cache** — if remembered from a recent visit, return immediately. (2) Queries the **recursive resolver** (usually your ISP's DNS or 8.8.8.8). (3) Recursive resolver asks the **root nameserver** (13 root servers globally) — returns the `.com` TLD nameserver IPs. (4) Queries the **TLD nameserver** for `.com` — returns Google's authoritative nameserver IPs. (5) Queries **Google's authoritative nameserver** — returns the actual IP for `google.com`. (6) Resolver caches the result (per TTL) and returns it to your browser.

This whole process takes 50-150ms the first time but is cached thereafter."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Cloudflare, Amazon Route 53 team, any infrastructure role

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does DNS work?
**Your Response:** DNS translates human-readable domain names like google.com into IP addresses that computers use to route packets. When you type google.com in your browser, first your OS checks its local cache - if remembered from a recent visit, it returns immediately. If not, it queries the recursive resolver (usually your ISP's DNS or 8.8.8.8). The resolver asks the root nameserver for the .com TLD nameserver, then queries that TLD nameserver for Google's authoritative nameserver, and finally queries Google's authoritative nameserver for the actual IP. The resolver caches the result per TTL and returns it to your browser. This whole process takes 50-150ms the first time but is cached thereafter.

#### Indepth
DNS record types:
- **A record:** Domain → IPv4 address. `google.com → 142.250.182.46`
- **AAAA record:** Domain → IPv6 address
- **CNAME:** Domain → another domain (alias). `www.example.com → example.com`. Note: CNAME can't coexist with other records at the same name (zone apex problem). ALIAS/ANAME records (Route 53) solve this.
- **MX record:** Mail exchange servers for email delivery. `example.com MX mail.example.com`
- **TXT record:** Arbitrary text. Used for domain verification, SPF, DKIM.
- **NS record:** Authoritative nameservers for the domain.

**DNS attack vectors:**
- **DNS spoofing / Cache poisoning:** Injecting fake DNS records into a resolver's cache. Mitigated by DNSSEC (signs all records with cryptographic signature).
- **DNS amplification DDoS:** Attacker sends small DNS queries with spoofed source IP (victim's IP). Resolver sends large responses to victim. 300x amplification factor. Mitigated by response rate limiting (RRL) and BCP38.
- **DNS hijacking:** ISP or malicious router intercepts DNS queries and returns their own responses. Mitigated by DNS-over-HTTPS (DoH) or DNS-over-TLS (DoT).

---

### 107. What is the OSI model?
"The OSI (Open Systems Interconnection) model is a **conceptual framework** that describes how different network protocol components interact when transferring data between systems. It has 7 layers, each with specific responsibilities.

I use it daily when debugging: 'Is it a physical problem (Layer 1)? A routing problem (Layer 3)? An application issue (Layer 7)?' Load balancers operate at L4 or L7. Firewalls inspect L3 or L4. Nginx works at L7."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** Network engineer roles, cloud/infra interviews, any question about where a component sits in the stack

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the OSI model?
**Your Response:** The OSI model is a conceptual framework that describes how different network protocol components interact when transferring data between systems. It has 7 layers, each with specific responsibilities. I use it daily when debugging - is it a physical problem at Layer 1, a routing problem at Layer 3, or an application issue at Layer 7? Load balancers operate at Layer 4 or 7, firewalls inspect Layer 3 or 4, and Nginx works at Layer 7. The layers from top to bottom are Application, Presentation, Session, Transport, Network, Data Link, and Physical.

#### Indepth
7 OSI Layers (top to bottom):
| Layer | Name | Purpose | Examples |
|---|---|---|---|
| 7 | Application | User-facing protocols | HTTP, FTP, SMTP, DNS, gRPC |
| 6 | Presentation | Encoding, encryption, compression | SSL/TLS, JPEG, ASCII |
| 5 | Session | Session management, authentication | NetBIOS, RPC, OAuth flow |
| 4 | Transport | End-to-end delivery, port-based | TCP, UDP, SCTP |
| 3 | Network | IP routing, addressing | IP, ICMP, BGP, OSPF |
| 2 | Data Link | Node-to-node delivery (MAC) | Ethernet, Wi-Fi (802.11), ARP |
| 1 | Physical | Bits over physical medium | Cables, fiber, radio waves, NICs |

Memory device: **"Please Do Not Throw Sausage Pizza Away"** (Physical, Data Link, Network, Transport, Session, Presentation, Application) — bottom to top.

Practical application: `curl https://api.myapp.com/data` involves:
- L7 (Application): HTTP GET request, DNS lookup
- L6 (Presentation): TLS encryption/decryption of HTTP
- L4 (Transport): TCP connection (SYN, SYN-ACK, ACK)
- L3 (Network): IP routing across the internet
- L2/L1: Ethernet frames, actual bits on cable/WiFi

---

### 108. Difference between HTTP/1.1, HTTP/2, and HTTP/3.
"Each major HTTP version fundamentally changed how web communication works.

**HTTP/1.1** (1997): One request per TCP connection (unless keep-alive). If a page needs 50 resources (CSS, JS, images), browsers open multiple parallel TCP connections (up to 6 per domain). Head-of-line blocking: a slow response blocks the next request on that connection.

**HTTP/2** (2015): Multiplexing — multiple requests and responses simultaneously over *one* TCP connection. No head-of-line blocking at the HTTP layer. Header compression (HPACK). Server push (proactively send resources before client asks). Requires HTTPS in practice (browsers only support HTTP/2 over TLS).

**HTTP/3** (2022): Same semantics as HTTP/2 but runs on **QUIC** (over UDP) instead of TCP. Eliminates TCP-level head-of-line blocking. Faster connection setup (0-RTT for known servers). Better on mobile (handles IP changes — walking from WiFi to cellular doesn't break the connection)."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Cloudflare, web performance discussions, frontend infrastructure — any company serving millions of users

### How to Explain in Interview (Spoken style format)
**Interviewer:** Difference between HTTP/1.1, HTTP/2, and HTTP/3.
**Your Response:** Each major HTTP version fundamentally changed web communication. HTTP/1.1 from 1997 uses one request per TCP connection unless keep-alive is used. If a page needs 50 resources, browsers open multiple parallel TCP connections up to 6 per domain, and there's head-of-line blocking where a slow response blocks the next request. HTTP/2 from 2015 introduced multiplexing - multiple requests and responses over one TCP connection, eliminating HTTP-level head-of-line blocking, plus header compression and server push. HTTP/3 from 2022 has the same semantics as HTTP/2 but runs on QUIC over UDP instead of TCP, eliminating TCP-level head-of-line blocking with faster connection setup and better performance on mobile.

#### Indepth
| Feature | HTTP/1.1 | HTTP/2 | HTTP/3 |
|---|---|---|---|
| Transport | TCP | TCP | QUIC (UDP) |
| Multiplexing | No (6 connections per domain) | Yes (one connection) | Yes |
| HoL Blocking | Yes (HTTP + TCP) | HTTP-level: No; TCP-level: Yes | No (QUIC streams independent) |
| Header Compression | None | HPACK | QPACK |
| Server Push | No | Yes (often disabled) | Yes |
| Encryption | Optional | Effectively required | Required |
| RTT to Connect | 1-2 RTT TCP + TLS | 1-2 RTT TCP + TLS | 1 RTT (0-RTT possible) |

**HTTP/3 adoption:** Already at ~30% of web traffic. Cloudflare, Google, Facebook were early adopters. All major CDNs support it. YouTube uses QUIC. Go's standard library got HTTP/3 support via the `quic-go` library.

---

### 109. What is API versioning?
"API versioning is the practice of allowing multiple versions of an API to coexist simultaneously, so existing clients don't break when you introduce changes.

There are four common strategies: **URL versioning** (`/api/v1/users`), **Header versioning** (`API-Version: 2025-01`), **Query parameter versioning** (`/api/users?version=2`), and **Content negotiation** (`Accept: application/vnd.myapp.v2+json`).

My preference: **URL versioning** for public APIs because it's explicit, easy to see in logs and curl commands, and simple for developers to understand. Header versioning is cleaner architecturally but less discoverable."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Razorpay (major version changes in their payment API), Stripe (famous for their stable API versioning), Amazon AWS API Gateway

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is API versioning?
**Your Response:** API versioning allows multiple versions of an API to coexist so existing clients don't break when you introduce changes. There are four common strategies: URL versioning like /api/v1/users, header versioning with API-Version header, query parameter versioning, and content negotiation. My preference is URL versioning for public APIs because it's explicit, easy to see in logs and curl commands, and simple for developers to understand. Header versioning is cleaner architecturally but less discoverable. The key is to never remove old versions without proper notice and migration support.

#### Indepth
Versioning strategy comparison:
| Strategy | Example | Pros | Cons |
|---|---|---|---|
| URL Path | `/v1/users` | Explicit, cacheable, easy LB routing | 'Ugly' URL, may violate REST purity |
| Query Param | `/users?v=2` | Simple | Easy to omit/forget |
| Header | `API-Version: 2` | Clean URL, flexible | Less discoverable, harder to test in browser |
| Media Type | `Accept: application/vnd.company.v2+json` | Most RESTful | Complex for clients to implement |

**Stripe's versioning model (industry gold standard):** Stripe dates all versions (`2024-04-10`). New features default to the account's pinned version. Developers explicitly upgrade their version, test the changes, then pin to the new date. Breaking changes NEVER affect existing customers' pinned version. Stripe maintains ALL previous API versions forever. This is the most client-friendly approach.

**Deprecation process:** Never remove an old API version without: (1) 6-12 months advance notice, (2) clear migration guide, (3) automated emails to affected API key holders, (4) a Sunset header in responses (`Sunset: Sat, 31 Dec 2025 23:59:59 GMT`).

---

### 110. What is a message broker?
"A message broker is a **middleware service that receives messages from producers, stores them temporarily, and delivers them to consumers**. It's the infrastructure that enables async, decoupled communication between services.

The key benefit: producers and consumers are independent. The Order service (producer) doesn't need to call the Email service (consumer) directly — it just drops a message into the broker. The Email service picks it up in its own time. If the Email service is down, messages accumulate in the broker — nothing is lost.

RabbitMQ and Kafka are the two most common. RabbitMQ is a traditional message queue: messages are consumed by one consumer and deleted. Kafka is a distributed log: messages are retained for days/weeks, multiple consumer groups can read the same messages independently."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Swiggy, Zomato, Razorpay, Amazon, Uber — event-driven microservices design

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a message broker?
**Your Response:** A message broker is a middleware service that receives messages from producers, stores them temporarily, and delivers them to consumers. It's the infrastructure that enables async, decoupled communication between services. The key benefit is that producers and consumers are independent. The Order service doesn't need to call the Email service directly - it just drops a message into the broker. The Email service picks it up in its own time. If the Email service is down, messages accumulate in the broker and nothing is lost. RabbitMQ and Kafka are the two most common - RabbitMQ is a traditional message queue where messages are consumed by one consumer and deleted, while Kafka is a distributed log where messages are retained for days or weeks and multiple consumer groups can read the same messages independently.

#### Indepth
RabbitMQ vs Kafka comparison:
| Feature | RabbitMQ | Kafka |
|---|---|---|
| Pattern | Queue (point-to-point or pub-sub) | Distributed commit log |
| Message retention | Deleted after ack | Retained for configured period (days/weeks) |
| Consumer model | Competing consumers (~load balanced) | Consumer groups (each group gets all messages) |
| Ordering | Per-queue ordering | Per-partition ordering |
| Throughput | Good (~50K msg/s) | Very high (millions/s) |
| Replayability | No | Yes — consumers can re-read from offset |
| Use case | Task queues, RPC | Event streaming, audit log, data pipelines |

**When to use Kafka over RabbitMQ:**
- Need multiple independent consumers to process the same events (audit service + email service both get the same order event)
- Need to replay events (bug in consumer → fix bug → re-process last 7 days of events)
- Very high throughput (>100K messages/second)
- Events as source of truth (event sourcing, CQRS)

**When to use RabbitMQ:** Simple task queue, complex routing (content-based routing, topic exchange), when you need strong per-message acknowledgment guarantees, RPC over messaging.
