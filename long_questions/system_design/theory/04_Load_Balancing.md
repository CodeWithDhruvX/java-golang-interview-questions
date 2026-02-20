# 🟡 Load Balancing — Questions 31–40

> **Level:** 🟡 Mid to 🔴 Senior
> **Asked at:** Amazon, Netflix, Cloudflare, Akamai, Flipkart, LinkedIn

---

### 31. What is a load balancer? (Deep Dive)
"A load balancer distributes incoming network traffic across a pool of backend servers to **maximize throughput, minimize latency, and ensure no single server is a bottleneck or single point of failure**.

It's not just about splitting traffic evenly. A good LB does health checks (removes unhealthy servers automatically), supports session stickiness (routing the same user to the same server), performs SSL termination (offloading TLS handshake from backend servers), and supports multiple routing algorithms.

I think of a load balancer as the 'front door' of any scalable service. Netflix, for example, has multiple layers: Zuul (their API gateway/LB), AWS ALBs per region, and then Ribbon (client-side LB) at the inter-service level."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** Any system design interview — foundational concept

#### Indepth
Load balancer implementation modes:
- **Hardware LB (F5, Citrix ADC):** Dedicated appliance, very high performance, expensive. Used by large enterprises and telcos. Not cloud-native.
- **Software LB (Nginx, HAProxy, Envoy):** Runs on commodity servers. Flexible, cheap, cloud-friendly.
- **Cloud-managed LB (AWS ALB/NLB, GCP Load Balancing):** Fully managed, auto-scaling, deeply integrated with cloud ecosystem.
- **Client-side LB (Netflix Ribbon, gRPC):** The client itself has a list of server IPs and load-balances its own requests. Zero additional network hop. Used heavily in microservices within a data center.

---

### 32. Types of load balancing strategies.
"Load balancing strategies determine *how* traffic is distributed. Choosing the right algorithm for your access pattern is crucial.

**Round Robin** cycles through servers sequentially — great for stateless services with uniform request complexity. **Least Connections** routes to the server with fewest active connections — better for long-lived connections like video streaming or WebSocket. **Weighted** variants assign more traffic to powerful servers. **IP Hash** creates pseudo-sticky sessions by always routing the same client IP to the same server.

For Uber's driver-matching service, they use consistent hashing so a driver's location updates always route to the same server — not for performance but for simplicity of state management."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Netflix, Cloudflare

#### Indepth
Detailed breakdown:
- **Round Robin:** Simple, equal distribution. Works only if all requests have similar processing time. If one request takes 10s and the next takes 1ms, the slow request holds a connection and a server appears "used".
- **Weighted Round Robin:** Server A gets 70% traffic, Server B gets 30%. Used when servers have different capacities (e.g., a bigger EC2 instance).
- **Least Connections:** Tracks active connection count per server. Routes new connection to minimally loaded server. Much better than Round Robin for variable-length requests (APIs, DB queries).
- **Least Response Time:** Combines connection count with average response time. Routes to the server with best performance. Used in HAProxy's `leastconn` with latency awareness.
- **IP Hash:** `server = hash(client_IP) % num_servers`. Deterministic — same client always hits same server. Breaks if a server is added or removed (consistent hashing solves this).
- **Resource Based (Adaptive):** LB queries backend health endpoints for CPU/memory, routes to least loaded. Requires health endpoint per server.
- **Random:** Surprisingly effective at scale with many servers — random selection approaches uniform distribution by the law of large numbers.

---

### 33. How do you implement sticky sessions?
"Sticky sessions (Session Affinity) ensure a particular user's requests always go to the same backend server. This is needed when the server stores session state locally (in-memory) rather than in a shared store.

The cleanest implementation is **cookie-based**: on the first request, the LB injects a cookie like `SERVERID=server-3` into the response. On subsequent requests, the LB reads this cookie and routes to server-3.

But here's my strong opinion: sticky sessions are a design smell. They create an implicit dependency between a user and a server instance. If server-3 dies, that user's session is lost unless you also replicate it. I prefer **stateless services** where sessions are stored in Redis — then any server can handle any request, and sticky sessions become unnecessary."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Legacy enterprise stacks (Infosys, Accenture clients) and modern systems (Razorpay, game platforms)

#### Indepth
Sticky session implementation methods:
1. **Source IP Affinity:** `hash(client_IP) % num_servers`. Simple, no cookie needed. Problem: multiple users behind the same NAT/corporate proxy have the same IP → all routed to same server → uneven load.
2. **Cookie-based (LB-injected):** LB adds `AWSALB` cookie (AWS ALB's name). Duration-based or request-based. Most flexible and accurate.
3. **Application-managed:** App generates a session ID, encodes server ID into it, and returns it in cookie. Application-level awareness.
4. **Consistent Hashing on UserID:** After authentication, use `hash(userId)` for routing. Requires auth to happen before LB routing decision.

**When sticky sessions are necessary:** WebSocket connections (must hit same server for the connection lifetime), server-side server-sent events (SSE) or polling, legacy stateful applications that can't be refactored.

---

### 34. What are health checks in load balancing?
"Health checks are the mechanism by which a load balancer ensures it only routes traffic to servers that are actually able to serve requests. Without them, a crashed server keeps receiving traffic and users get errors.

There are two types: **Active health checks** where the LB proactively sends a request (like `GET /health`) to every server every N seconds. If a server returns a non-2xx response or times out, it's marked unhealthy and removed from rotation. **Passive health checks** where the LB monitors real traffic — if a server returns consecutive 5xx errors, it marks it unhealthy.

For my services, I implement a `/health` endpoint that doesn't just return 200 — it actually checks downstream dependencies (DB connectivity, Redis connection) and returns health status of each subsystem."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any SRE, DevOps, or senior backend role — Amazon, Google, Netflix, Razorpay

#### Indepth
Health check design best practices:
- **Shallow vs Deep health checks:** 
  - Shallow: Just returns 200 if the process is alive. Fast, but doesn't detect downstream failures (DB is down but the process is still running).
  - Deep: Checks DB, Redis, downstream APIs. More accurate, but risk of cascading health check failures if a shared dependency is slow.
- **Separate LB health vs readiness vs liveness:** Kubernetes separates these: *liveness* = is the process alive (restart if no), *readiness* = is it ready to serve traffic (remove from LB if no). The LB should use the readiness probe.
- **Health check endpoint implementation:**
  ```json
  GET /health → 200 OK
  {
    "status": "ok",
    "database": "ok",
    "redis": "ok",
    "kafka": "degraded - 200ms latency",
    "version": "v2.1.3"
  }
  ```
- **Circuit breaker integration:** Health check failures can trigger circuit breakers in upstream services, preventing them from calling a degrading service.

---

### 35. Difference between Layer 4 and Layer 7 load balancers.
"Layer 4 and Layer 7 refer to OSI model layers. The key difference is how much of the network packet they inspect.

**Layer 4 LB (transport layer):** Works with TCP/UDP. It sees IP addresses and ports only — it doesn't open the packet. It's essentially a very fast packet forwarder. Lower latency because no parsing is needed. AWS NLB is Layer 4.

**Layer 7 LB (application layer):** Opens the HTTP packet and reads headers, URLs, cookies, request body. It can make intelligent routing decisions: route `/api/` to the API cluster, `/static/` to the CDN origin, and authenticated requests to the premium tier. It can also do SSL termination, request header injection, response compression, and more. AWS ALB is Layer 7. This is what most modern applications need."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon (AWS certification-level questions), Cloudflare, Netflix, Razorpay

#### Indepth
| Feature | Layer 4 (TCP/UDP) | Layer 7 (HTTP/HTTPS) |
|---|---|---|
| OSI Layer | Transport (4) | Application (7) |
| Visibility | Source/Dest IP + Port only | Full HTTP: URL, headers, body, cookies |
| SSL Termination | No (SSL passthrough to backend) | Yes — terminates TLS, backend gets HTTP |
| Content Routing | No — can't inspect URL | Yes — route by path, host, header |
| Latency | Lower (no packet inspection) | Slightly higher |
| Protocols | TCP, UDP, any | HTTP, HTTPS, WebSocket, gRPC |
| Example | AWS NLB, F5 | AWS ALB, Nginx, HAProxy, Envoy |
| Use Case | Gaming (UDP), DB-level LB, SMTP | REST APIs, microservices, web apps |

**Hybrid use:** AWS puts NLB in front of ALB for elastic IPs (NLB provides static IPs that can be whitelisted by enterprise firewalls). This gives you the IP stability of L4 with the routing flexibility of L7.

---

### 36. What is DNS load balancing?
"DNS load balancing distributes traffic by returning different IP addresses in response to DNS queries for the same hostname.

For example: `api.myapp.com` might resolve to `192.168.1.1` for one user and `192.168.1.2` for another. The simplest implementation is **Round Robin DNS** where the DNS server rotates through a list of IPs.

The limitation is DNS caching — browsers, OS resolvers, and ISPs all cache DNS responses for the TTL duration. If `192.168.1.1` goes down, users with cached DNS entries keep sending traffic there until their cached entry expires. This makes DNS LB poor for health-based failover. It works best as a **geographic routing** mechanism: route Indian users to the Mumbai region, US users to the Virginia region."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Cloudflare, Akamai, Amazon Route 53 team, CDN-heavy product companies

#### Indepth
Advanced DNS-based load balancing:
- **GSLB (Global Server Load Balancing):** DNS-based geo-routing with health awareness. AWS Route 53 routing policies: Latency-based, Geo, Weighted, Failover. Route 53 health checks mark endpoints unhealthy and removes them from DNS responses within 60-300 seconds.
- **Anycast:** Multiple servers around the world advertise the *same* IP via BGP. The internet's routing layer automatically directs the user's packet to the nearest server. Used by Cloudflare, Google's 8.8.8.8 DNS. This is different from DNS round-robin — it works at the network layer, not DNS layer.
- **Low TTL trade-off:** Short TTL (30 seconds) enables faster failover but increases DNS query load. Long TTL (3600 seconds) reduces DNS load but slow failover. Cloudflare uses 300s TTL as a balance.

---

### 37. Explain round-robin vs least connections algorithm.
"Round-robin and least-connections represent two different philosophies for routing.

**Round-robin** assumes all requests are equal — each server gets every N-th request in rotation. It's dead simple and works perfectly when requests are stateless and similarly complex (like serving a static file).

**Least Connections** is dynamic — route to the server currently handling the fewest connections. This is dramatically better for heterogeneous workloads. Imagine some API requests take 1ms and others take 30 seconds (long-polling). With Round Robin, all servers get equal requests but some will be backed up with slow requests. With Least Connections, the backed-up server (with 50 active connections) stops getting new requests until it drains."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Netflix (streaming — variable-length connections), Amazon, Cloudflare

#### Indepth
When each algorithm wins:
- **Round Robin:** Stateless HTTP APIs with uniform request complexity, file servers, CDN edge nodes with many ephemeral connections.
- **Least Connections:** WebSocket servers (long-lived connections cause uneven load), database proxy load balancers (query times vary wildly), video transcoding (long-running jobs per connection).
- **Least Response Time (HAProxy: `leastconn` + `server-template`):** The most adaptive — measures actual response time and routes to fastest server. Implements a self-adjusting feedback loop.

**Consistent Hashing** is a separate-but-related concept for *stateful* routing: route the same user/session to the same server (for caching or state locality), but when servers are added or removed, only a minimal fraction of keys are remapped. Used for cache servers (Memcached cluster), streaming servers, and Uber's driver routing.

---

### 38. How to handle load balancer failures?
"The load balancer itself must not be a Single Point of Failure (SPOF). If the one LB you have crashes, everything behind it goes dark — no matter how many healthy app servers you have.

The standard solution is **Active-Passive LB pair using VRRP/Keepalived**: two LB instances share a Virtual IP (VIP). The active one handles all traffic. Both continuously exchange heartbeat messages. If the active LB misses heartbeats, the passive instantaneously takes over the VIP — the failover is sub-second. Clients never know there were two LBs; they always connect to the VIP.

For cloud-native setups, managed LBs (AWS ALB, GCP Load Balancing) are deployed as distributed infrastructure by the cloud provider. They're inherently HA — the concept of 'LB failure' is abstracted away."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** On-prem architecture discussions — Infosys, TCS delivery, financial companies. Cloud discussions at Amazon, Google

#### Indepth
LB HA architectures:
- **Active-Passive with VRRP:** Keepalived manages VIP takeover. Passive LB sits idle (resource waste). Failover: ~1-2 seconds for IP reassignment + BGP propagation.
- **Active-Active (DNS):** Two LBs, each with their own IP. DNS round-robins between both. If one dies, some % of requests fail until DNS propagates the health update. DNS TTL determines recovery time.
- **BGP Anycast (Cloudflare model):** Each of 200+ PoPs announces the same IP prefix. Traffic is naturally routed to the nearest PoP. If a PoP fails, BGP reconverges (~seconds), traffic redirects to next nearest PoP. Effectively infinite redundancy.
- **AWS ALB HA model:** AWS ALB is a managed distributed system — it has dozens of internal nodes behind the scenes. AWS guarantees 99.99% SLA. ALB nodes are deployed across multiple AZs; it automatically routes around failed AZ nodes.

---

### 39. What is geo-load balancing?
"Geo-load balancing routes user requests to the nearest or most appropriate data center based on geographic location — reducing latency and enabling data residency compliance.

The mechanism: a user in Mumbai making a request to `api.myapp.com` gets an IP that points to the Mumbai data center (or the nearest AWS `ap-south-1` region), not a US region. The latency difference: 5ms vs 220ms. For real-time applications (chat, gaming, financial trading), this is the difference between a good and terrible user experience.

AWS Route 53 with Latency-based routing, Cloudflare Load Balancing, and Google Cloud's Global Load Balancing all offer geo-routing as a built-in feature. I use Route 53 health checks + failover routing — if the nearest region fails, Route 53 automatically routes to the next best region."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Global companies: Amazon (Prime Video worldwide), Netflix, Google, Razorpay (India + APAC expansion)

#### Indepth
Geo-load balancing strategies:
- **Geo-proximity routing (AWS Route 53):** Route based on geographic location, with bias adjustments. Good for co-located CDN origin + backend.
- **Latency-based routing (AWS Route 53):** Route to the region with lowest measured latency for the user's location. More accurate than geo-proximity (accounts for network topology).
- **GSLB (Global Server Load Balancing):** Enterprise-grade solution from F5, Citrix. Combines DNS routing with health monitoring across DCs.
- **BGP Anycast:** Network layer routing. Used by DNS providers, DDoS mitigation services, CDNs.

**Compliance use case:** GDPR requires EU user data to stay in EU. India's DPDP Act has similar requirements. Geo-load balancing enforces data residency: EU users are always routed to EU data centers. This is not just a performance feature — it's a legal requirement for many products.

---

### 40. How to design a multi-region load balancing setup?
"Multi-region load balancing is one of the most architecturally complex designs because you're dealing with geographic distribution, data replication, and failover at global scale.

My design has three layers: **Global (DNS/Anycast)** routes users to their nearest Region. **Regional LB** (AWS ALB) distributes traffic across availability zones within that region. **Internal service mesh** (Envoy/Istio) handles inter-service routing within the region.

For failover: if the entire Mumbai region fails, Route 53 health checks detect it within ~30 seconds and update DNS to route all Mumbai traffic to the Singapore region. The Singapore region must have synchronized data (via async replication) to serve those users. This is the trade-off: some users may see slightly stale data during the failover window."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Google, Amazon (AWS Well-Architected Framework discussions), Razorpay (global expansion), Hotstar (India + worldwide for cricket)

#### Indepth
Complete multi-region architecture:

```
                    ┌─── DNS (Route 53/Cloudflare)
                    │    Geo-routing + health checks
                    │
        ┌───────────┴────────────┐
        ▼                        ▼
  [Mumbai Region]          [Singapore Region]
  AWS ap-south-1           AWS ap-southeast-1
  │                        │
  ├─ ALB (multi-AZ)        ├─ ALB (multi-AZ)
  ├─ App Servers           ├─ App Servers
  ├─ Cache (Redis)         ├─ Cache (Redis)
  └─ DB (Primary)◄─────── └─ DB (Read Replica)
         ▲                      async replication
         │
   Writes flow here; reads distributed across regions
```

Key design decisions:
- **Active-Active vs Active-Passive:** Active-Active (both regions serve traffic) maximizes utilization but requires cross-region data synchronization. Active-Passive is simpler — one region is hot, the other is a warm standby. AWS recommends pilot light (minimal standby) or warm standby.
- **Data replication lag:** Async replication means the standby region may be N seconds behind. Business must accept this RPO (Recovery Point Objective).
- **Session management:** Cross-region sticky sessions don't work. Store sessions in a globally distributed store (Amazon DynamoDB Global Tables, Redis with global replication).
- **RTO and RPO:** Define clearly. RTO = how long to recover. RPO = how much data loss is acceptable. DNS failover takes 60-300 seconds (depending on TTL). Ensure this is within your RTO.
