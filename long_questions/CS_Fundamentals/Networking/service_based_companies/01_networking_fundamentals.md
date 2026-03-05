# 🌐 Networking Fundamentals — Interview Questions (Service-Based Companies)

This document covers networking concepts commonly tested at service-based companies like TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, and Tech Mahindra. Targeted at 1–5 years of experience rounds.

---

### Q1: What is the OSI model? Name all 7 layers and their responsibilities.

**Answer:**
The **OSI (Open Systems Interconnection)** model is a conceptual framework that standardizes how different network systems communicate.

| Layer | Name | Responsibility | Examples |
|---|---|---|---|
| 7 | **Application** | User-facing protocols, data formats | HTTP, FTP, SMTP, DNS |
| 6 | **Presentation** | Encryption, encoding, compression | TLS/SSL, JPEG, ASCII |
| 5 | **Session** | Session management (open/close/resume) | NetBIOS, RPC |
| 4 | **Transport** | End-to-end delivery, flow/error control | TCP, UDP |
| 3 | **Network** | Logical addressing, routing | IP, ICMP, routers |
| 2 | **Data Link** | Physical addressing (MAC), error detection | Ethernet, Wi-Fi, switches |
| 1 | **Physical** | Raw bits over physical medium | Cables, hubs, fiber |

**Memory tip:** "**A**ll **P**eople **S**eem **T**o **N**eed **D**ata **P**rocessing" (top-down)

**Interview shortcut:** For most backend questions, focus on Layers 3 (IP routing), 4 (TCP/UDP), and 7 (HTTP/application protocols).

---

### Q2: What is the difference between TCP and UDP? When do you use each?

**Answer:**

| Feature | TCP | UDP |
|---|---|---|
| Full form | Transmission Control Protocol | User Datagram Protocol |
| Connection | Connection-oriented (3-way handshake) | Connectionless |
| Reliability | Guaranteed delivery, retransmits on loss | No guarantee, fire-and-forget |
| Ordering | Data arrives in order | No ordering guarantee |
| Flow control | Yes (sliding window) | No |
| Speed | Slower (overhead of reliability) | Faster |
| Header size | 20 bytes minimum | 8 bytes |

**Use TCP when:**
- Data integrity is critical: HTTP/HTTPS, file transfer (FTP/SFTP), email (SMTP), database connections
- Example: Uploading a file — every byte must arrive correctly

**Use UDP when:**
- Speed matters more than reliability: Video streaming, online gaming, VoIP, DNS queries
- Example: A video call — a dropped frame is better than freezing to wait for retransmit

**3-Way Handshake (TCP):**
```
Client → Server: SYN (I want to connect)
Server → Client: SYN-ACK (OK, ready)
Client → Server: ACK (Got it, connected!)
```

---

### Q3: What is the difference between HTTP/1.1, HTTP/2, and HTTP/3?

**Answer:**

| Feature | HTTP/1.1 | HTTP/2 | HTTP/3 |
|---|---|---|---|
| Year | 1997 | 2015 | 2022 |
| Transport | TCP | TCP | QUIC (UDP-based) |
| Multiplexing | No (head-of-line blocking) | Yes (multiple streams) | Yes (no head-of-line blocking) |
| Header compression | No | Yes (HPACK) | Yes (QPACK) |
| Server Push | No | Yes | Yes |
| TLS | Optional | Required (in practice) | Built-in (always encrypted) |

**HTTP/1.1 problem:** Only one request per TCP connection at a time (or pipelining, which is buggy). Browsers open 6 parallel connections as workaround.

**HTTP/2 improvement:** Multiplexes multiple requests over a single TCP connection. But TCP head-of-line blocking still exists — a lost packet stalls ALL streams.

**HTTP/3 solution:** Uses QUIC (over UDP), which manages streams independently. A lost packet only blocks its own stream, not others.

---

### Q4: What is a subnet and what is CIDR notation?

**Answer:**
A **subnet** (subnetwork) divides a larger network into smaller, more manageable pieces. It improves security and reduces network congestion.

**CIDR (Classless Inter-Domain Routing)** notation expresses an IP address with its subnet mask:
```
192.168.1.0/24
```
- `192.168.1.0` → Network address
- `/24` → 24 bits are the network portion, 8 bits are for hosts
- `/24` = 256 addresses (254 usable — first is network, last is broadcast)

**Common CIDR blocks:**

| CIDR | Subnet Mask | Hosts Available |
|---|---|---|
| /8 | 255.0.0.0 | ~16 million |
| /16 | 255.255.0.0 | ~65,000 |
| /24 | 255.255.255.0 | 254 |
| /28 | 255.255.255.240 | 14 |
| /32 | 255.255.255.255 | 1 (single host) |

**Real-world:** Cloud VPCs (Virtual Private Clouds) use CIDR blocks. E.g., AWS default VPC is `172.31.0.0/16`.

---

### Q5: What is the difference between a hub, switch, and router?

**Answer:**

| Device | OSI Layer | What it does |
|---|---|---|
| **Hub** | Layer 1 (Physical) | Broadcasts all data to ALL ports — dumb device, creates collisions |
| **Switch** | Layer 2 (Data Link) | Forwards data to the **specific port** using MAC address table |
| **Router** | Layer 3 (Network) | Routes data between **different networks** using IP addresses |

**Analogy:**
- **Hub**: Shout in a room — everyone hears (broadcasting)
- **Switch**: Whisper directly to one person — only they hear (MAC-based forwarding)
- **Router**: Post office — routes letters between different countries/cities (IP routing)

**Modern networks:** Hubs are obsolete. Switches handle local traffic; routers connect to internet/different subnets.

---

### Q6: What is NAT (Network Address Translation)?

**Answer:**
**NAT** allows multiple devices on a private network to share a single public IP address.

**How it works:**
1. Devices in your home/office have private IPs (e.g., `192.168.1.x`).
2. The router has one public IP (e.g., `203.0.113.5`).
3. When a device sends a request, the router **replaces** the private source IP with its public IP and stores the mapping.
4. When the response returns, the router **translates** back to the private IP and forwards to the device.

**Types:**
- **SNAT (Source NAT)**: Changes source IP — most common (home routers, cloud VMs)
- **DNAT (Destination NAT)**: Changes destination IP — used in port forwarding/load balancers

**Why NAT exists:** IPv4 addresses are limited (~4.3 billion). NAT lets millions of devices share a handful of public IPs.

---

### Q7: What is a firewall, and what is the difference between a stateful and stateless firewall?

**Answer:**
A **firewall** inspects network traffic and allows or denies it based on rules.

**Stateless Firewall:**
- Inspects each packet independently.
- Checks: Source IP, destination IP, port, protocol.
- Does NOT track connection state.
- Fast but less secure — cannot understand context.

**Stateful Firewall:**
- Tracks the **state of active connections** in a connection table.
- Knows if an incoming packet is part of an established, valid connection.
- Automatically allows return traffic for outgoing connections.
- Slower than stateless but far more secure — default in modern systems.

**Example:**
- You send a request to `google.com:443`.
- Stateful firewall remembers this outgoing connection.
- When Google's server responds, the firewall sees the reply is part of a tracked session → allows it automatically.

---

### Q8: What happens when you type `www.google.com` in your browser and press Enter?

**Answer:** (Classic interview question — walk through each step)

1. **URL Parsing**: Browser parses the URL — scheme (HTTPS), host (`www.google.com`), path (`/`).
2. **DNS Resolution**:
   - Browser cache → OS cache → ISP Recursive Resolver → Root DNS → TLD `.com` → Google's Authoritative DNS → Returns IP (e.g., `142.250.80.68`)
3. **TCP Connection**: Browser initiates 3-way TCP handshake with Google's server on port 443.
4. **TLS Handshake**: Browser and server negotiate TLS version, exchange certificates, derive session keys.
5. **HTTP Request**: Browser sends `GET / HTTP/2` with headers (Host, Accept, Cookie, etc.)
6. **Server Processing**: Google's load balancer receives the request, routes to backend server, which generates the response.
7. **HTTP Response**: Server sends back `200 OK` with HTML content.
8. **Browser Rendering**: Browser parses HTML → fetches CSS/JS/images → renders the page.

---

*Prepared for technical screening at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, Tech Mahindra).*
