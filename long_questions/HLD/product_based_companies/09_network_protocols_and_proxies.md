# High-Level Design (HLD): Network Protocols and Proxies

Top-tier product companies expect senior candidates to understand the underlying networking protocols rather than just REST APIs, as high-scale systems often hit bottlenecks at the network layer.

## 1. TCP vs. UDP: When to use which?
**Answer:**
*   **Transmission Control Protocol (TCP):** Connection-oriented. Guarantees delivery, order, and error-checking. Uses a 3-way handshake (`SYN`, `SYN-ACK`, `ACK`) before sending data.
    *   *Pros:* Reliable. If a packet is lost, it is retransmitted.
    *   *Cons:* High overhead, higher latency. Head-of-line blocking (if one packet is lost, subsequent packets must wait even if they arrived successfully).
    *   *Use Cases:* Web browsing (HTTP), File Transfers (FTP), Email (SMTP), reliable database connections.
*   **User Datagram Protocol (UDP):** Connectionless. "Fire and forget." No guarantee of delivery, order, or error-checking.
    *   *Pros:* Extremely fast, low overhead, no connection setup time.
    *   *Cons:* Unreliable (data can be lost or arrive out of order).
    *   *Use Cases:* Live Video/Audio Streaming (Zoom, Twitch), Online Multiplayer Gaming, DNS resolution, VoIP.

## 2. Compare HTTP/1.1, HTTP/2, and HTTP/3 (QUIC).
**Answer:**
*   **HTTP/1.1:** Text-based. Sequential. Suffers from Head-of-Line (HoL) blocking at the application layer because multiple requests over a single TCP connection must be served one by one.
*   **HTTP/2:** Binary protocol. Introduces **Multiplexing** allowing multiple requests and responses to fly concurrently over a *single* TCP connection, solving application-layer HoL blocking. Also features Server Push and Header Compression (HPACK). However, it still uses TCP, so it suffers from TCP-level HoL blocking (packet loss stalls everything).
*   **HTTP/3 (QUIC):** Built on **UDP** instead of TCP. Solves the TCP-level HoL blocking problem because streams are independent; losing a packet for Stream A doesn't block Stream B. It incorporates TLS 1.3 natively data, meaning a 0-RTT (Zero Round Trip Time) connection resumption.

## 3. WebSockets vs. Server-Sent Events (SSE) vs. Long Polling
**Answer:**
When you need real-time communication (e.g., chat applications, live stock tickers):
*   **Long Polling:** Client requests data. Server holds the request open until new data is available, sends the response, and the connection closes. The client immediately opens a new one. (High overhead, resource-intensive).
*   **Server-Sent Events (SSE):** One-way communication from Server sequentially down to Client over a single long-lived HTTP connection.
    *   *Use Case:* Live stock tickers, news feed updates, live sports scores. The client doesn't need to talk back much.
*   **WebSockets:** Full-duplex bidirectional communication over a single, persistent TCP connection. Both client and server can send data at any time.
    *   *Use Case:* Chat applications (WhatsApp Web), multiplayer games, collaborative editing (Google Docs).

## 4. Forward Proxy vs. Reverse Proxy
**Answer:**
*   **Forward Proxy:** Sits in front of client machines. When a client makes a request to the internet, it goes through the proxy.
    *   *Purpose:* Used by corporations to block access to certain sites, anonymize the client's IP, or cache content for internal users. The internet thinks the request came from the proxy, not the client.
*   **Reverse Proxy:** Sits in front of the backend servers. When a client on the internet makes a request, it hits the reverse proxy, which then routes it to the appropriate internal server.
    *   *Purpose:* Load balancing, SSL termination, caching, protection from DDoS attacks, hiding internal network topology. Examples: Nginx, HAProxy.

## 5. What is an API Gateway? How does it differ from a Load Balancer?
**Answer:**
While a Load Balancer simply distributes incoming traffic across multiple servers (mostly at Layer 4 or basic Layer 7 routing), an API Gateway is the central entry point for a microservices architecture.
*   **API Gateway Responsibilities:**
    *   **Request Routing:** Routing `/users` to the User Service and `/orders` to the Order Service.
    *   **Authentication & Authorization:** Validating JWT tokens before hitting internal services.
    *   **Rate Limiting & Throttling:** Preventing abuse by limiting requests per user/IP.
    *   **API Composition/Aggregation:** Fetching user data from Service A and order data from Service B, combining them, and returning a single JSON response to the client.
    *   **Protocol Translation:** Accepting HTTP from the client but making a gRPC call to the internal backend service.
