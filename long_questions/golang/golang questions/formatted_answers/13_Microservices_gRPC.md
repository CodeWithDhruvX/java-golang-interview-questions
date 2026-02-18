# 🔵 **241–260: Microservices, gRPC, and Communication**

### 241. What is gRPC and how is it used with Go?
"gRPC is a high-performance RPC framework from Google.
It uses **Protocol Buffers** (protobufs) as the IDL (Interface Definition Language).

In Go, I define my service structure in a `.proto` file (`service Greeter { rpc SayHello ... }`).
I run `protoc` to generate Go code.
Then I implement the interface. It’s strictly typed, uses HTTP/2 for transport (multiplexing), and is significantly faster than JSON REST because the payload is binary."

#### Indepth
Proto3 (the current version) removed "required" fields. All fields are optional by default. This seems weird but is essential for backward compatibility. If you add a required field, old clients crash. If you remove one, old servers crash. Making everything optional forces you to handle missing data gracefully.

---

### 242. How do you define Protobuf messages for Go?
"I write a `.proto` file.

`message User { string id = 1; string email = 2; }`
The numbers `1` and `2` are crucial—they replace the field names in the binary wire format.
When I generate the Go code using `protoc-gen-go`, this becomes a struct `User` with fields `Id` and `Email` (capitalized). It also generates methods for serialization (`Marshal` / `Unmarshal`)."

#### Indepth
Go generated code includes a `UnimplementedGreeterServer` struct. You *must* embed this struct in your implementation. This ensures forward compatibility: if you add a new RPC method to the `.proto` file but haven't implemented it in Go yet, the server will compile and return `Unimplemented` instead of panicking.

---

### 243. What are the benefits of gRPC over REST?
"1. **Contract First**: The `.proto` file is the single source of truth. Client and Server code are generated from it, guaranteeing they are always in sync.
2.  **Performance**: Binary serialization is smaller and faster to parse than JSON text.
3.  **Streaming**: Built-in support for bi-directional streaming, which is awkward in REST/HTTP1.1."

#### Indepth
The biggest downside of gRPC is **Browser Support**. Browsers don't support raw HTTP/2 frames exposed to JS. You need a proxy like **gRPC-Web** (Envoy) to translate JSON-over-HTTP1.1 from the frontend into gRPC-over-HTTP/2 for the backend. It adds operational complexity.

---

### 244. How do you implement unary and streaming RPC in Go?
"**Unary**: Traditional request/response.
`rpc GetUser(ID) returns (User)`.

**Streaming**:
*   **Server-side**: `rpc ListUsers(Request) returns (stream User)`. The client gets an iterator and reads users one by one.
*   **Bi-directional**: `rpc Chat(stream Message) returns (stream Message)`.
In Go, the specialized `stream` object has `Send()` and `Recv()` methods that block until data is available."

#### Indepth
Streaming is powerful but dangerous. A slow client can block the server if the TCP window fills up ("HOL blocking" at app level). Unlike Unary calls where the request fits in memory, streams can be infinite. Always enforce a timeout or a max-message-count limit on streams to prevent resource exhaustion.

---

### 245. What is the difference between gRPC and HTTP/2?
"gRPC is the **Application Layer**. HTTP/2 is the **Transport Layer**.

gRPC builds *on top of* HTTP/2.
It leverages HTTP/2 features like:
*   **Multiplexing**: Many RPC calls over a single TCP connection.
*   **Header Compression**: Efficient metadata exchange.
of HTTP/2 like: **Header Compression**."

#### Indepth
gRPC uses **HPACK** for header compression. If you send the same auth token (which is huge) in every request, HPACK compresses it to a few bytes after the first request. This saves massive bandwidth compared to REST, where headers are sent as plain text every single time.

---

### 246. How do you add authentication in gRPC services?
"I use **Interceptors** (Middleware).

Client side: I add credentials to the context (`metadata.AppendToOutgoingContext(ctx, "authorization", token)`).
Server side: My `UnaryInterceptor` extracts the metadata from the incoming context.
`md, ok := metadata.FromIncomingContext(ctx)`
I validate the JWT token. If valid, I call `handler(ctx, req)`. If not, I return `codes.Unauthenticated`."

#### Indepth
Interceptors are chained. The order matters! `Logger -> CrashRecovery -> Auth -> RateLimit`. If you put Auth before Logger, you won't log failed login attempts. If you put RateLimit before Auth, attackers can DDOS your expensive Auth verification logic.

---

### 247. How do you handle timeouts and retries in gRPC?
"Timeouts are mandatory.
On the client, I set a Deadline on the context.
`ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)`
If the server takes too long, the client kills the request with `DeadlineExceeded`.

For retries, I use the `grpc.WithDefaultServiceConfig` option to enable automatic retries for specific error codes (like `Unavailable` or `ResourceExhausted`)."

#### Indepth
Retries **must** have jitter. If a microservice crashes and 1000 clients retry exactly 1 second later, they will re-crash the service (Thundering Herd). `Backoff = Base * 2^Attempt + Random(0-100ms)`. This spreads the load.

---

### 248. How do you secure gRPC communication?
"I strictly use **TLS**.

Server: `creds := credentials.NewServerTLSFromFile("cert.pem", "key.pem")`.
`s := grpc.NewServer(grpc.Creds(creds))`.

Client: `creds := credentials.NewClientTLSFromFile("cert.pem", "")`.
`conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(creds))`.
This ensures the entire channel is encrypted, preventing eavesdropping."

#### Indepth
With Let's Encrypt and modern protocols, we typically use **ALPN** (Application-Layer Protocol Negotiation). When the TLS handshake happens, the client sends "I speak h2" (HTTP/2). The server agrees. This allows running gRPC and standard HTTPS on the *same port* (443).

---

### 249. How do microservices communicate securely in Go?
"I enforce **mTLS** (Mutual TLS).

This means the Server verifies the Client's certificate *and* the Client verifies the Server's certificate.
Nobody can talk to my API unless they possess a valid certificate signed by my internal CA.
I configure `tls.Config{ClientAuth: tls.RequireAndVerifyClientCert}`. It creates a Zero Trust network."

#### Indepth
Managing certificates (rotation, revocation) is hell. This is why **Service Meshes** (Linkerd, Istio) exist. They inject a sidecar proxy that handles mTLS automatically. The Go app speaks plain text to localhost, and the proxy encrypts traffic to the destination. It offloads complexity from the code.

---

### 250. What are message queues and how to use them in Go?
"A Message Queue (RabbitMQ, Kafka) decouples services.

Instead of Service A calling Service B synchronously (tight coupling), A publishes an event: `OrderCreated`.
Service B consumes it when it's ready.
In Go, I run the consumers in background goroutines. This allows Service A to respond to the user immediate response, improving perceived performance and availability."

#### Indepth
Queues provide **Load Smoothing**. If traffic spikes to 5x, the API accepts it instantly, filling the queue. The consumers churn through it at their constant max speed. The system latency increases, but it *doesn't crash*. This is the primary reason to use queues over HTTP.

---

### 251. How to use NATS or Kafka in Go?
"For **NATS**, I use `nats.go`. It’s incredibly simple and fast.
`nc.Publish("updates", data)`.
`nc.Subscribe("updates", func(m *Msg) { ... })`.

For **Kafka**, I use `sarama` or `segmentio/kafka-go`.
It’s more complex due to consumer groups and offset management. I implement a loop that reads batches of messages, processes them, and commits the offsets."

#### Indepth
NATS Core is "At Most Once" (fire and forget). NATS JetStream is "At Least Once" (durable). Kafka is "At Least Once". In Go, always design your consumers to be **Idempotent**. If you process the same "PaymentDeducted" message twice, the second one should be a no-op, not a double charge.

---

### 252. What are sagas and how would you implement them in Go?
"Sagas manage distributed transactions without two-phase commit.

If I have a chain of actions (Book Hotel -> Book Flight -> Charge Card), and 'Charge Card' fails, I must undo previous steps.
I trigger **Compensating Transactions**: 'Cancel Flight' -> 'Cancel Hotel'.
I implement this via an Orchestrator service (using a state machine) or Choreography (events) where each service listens for failure events."

#### Indepth
The "Saga Pattern" is complex to debug. If a compensation fails (e.g., "Cancel Flight" fails because the airline API is down), you are stuck in an inconsistent state. You need manual intervention (Human in the Loop) or infinite retries. Always log "Zombie Sagas" heavily.

---

### 253. How would you trace requests across services?
"I use **OpenTelemetry**.

I start a Span at the edge (API Gateway).
I create a `TraceID`.
I inject this ID into the HTTP headers (`traceparent`) or gRPC metadata.
Every downstream service extracts this ID and creates a Child Span.
Finally, I visualize the entire waterfall in Jaeger or Grafana Tempo to find the slow service."

#### Indepth
Distributed tracing adds overhead (serializing context, sending spans). Use **Sampling**. For high-traffic services, sample 0.1% of requests (`TraceID % 1000 == 0`). You still get the statistical picture without burning CPU on headers for every health check.

---

### 254. What is service discovery and how do you handle it?
"It’s how Service A finds the IP of Service B.

In **Kubernetes**, it’s built-in via DNS (`http://payment-service`). I just call that hostname.
Outside K8s, I use **Consul** or **Etcd**.
My Go app queries Consul: 'Give me healthy IPs for Payment Service'. I use client-side load balancing to pick one and connect."

#### Indepth
Client-side load balancing (in gRPC) is superior to L4 (TCP) load balancing. L4 balancers see long-lived persistent connections and can't distribute requests evenly. Client-side LB sees *individual RPCs* and can do Round Robin or Least Loaded distribution per request.

---

### 255. How do you implement rate limiting across services?
"I use a **Distributed Rate Limiter** backed by Redis.

Libraries like `go-redis/redis_rate` implement the Leaky Bucket algorithm.
Service A checks Redis: `Allow("user:123", 10 requests/min)`.
If Redis says 'No', I return HTTP 429.
Redis is atomic, so it works perfectly even if I have 50 replicas of Service A running."

#### Indepth
If Redis becomes a bottleneck (hot key), switch to a **Two-Layer Limiter**. Use a local in-memory token bucket (e.g., 50 req/s) which is fast but approximate, and sync it with Redis asynchronously. This sacrifices strict global accuracy for massive performance.

---

### 256. What is the role of API gateway in microservices?
"It is the single Entry Point.

It handles cross-cutting concerns:
*   **Auth**: Verifies JWTs.
*   **Rate Limiting**: Protects backends.
*   **Routing**: `/api/v1/users` -> User Service.
I often use standard tools like **Kong** or **Traefik**, but for custom logic, I write a Go gateway using `httputil.ReverseProxy`."

#### Indepth
The "Backends for Frontends" (BFF) pattern is a variation where you have specific Gateways for specific clients (Mobile Gateway vs Web Gateway). This allows the Mobile Gateway to strip unused fields to save bandwidth, while the Web Gateway serves fuller data. Go is excellent for writing these lightweight aggregators.

---

### 257. How do you use OpenTelemetry with Go?
"I initialize a `TracerProvider` that points to my collector (e.g., OTLP/Jaeger).

In my code:
`ctx, span := tracer.Start(ctx, "CalculateTax")`
`defer span.End()`
I add attributes `span.SetAttributes(attribute.String("user_id", id))`.
If an error occurs, `span.RecordError(err)`.
This gives me rich, structured data about every operation's latency."

#### Indepth
OpenTelemetry (OTel) is the industry standard now, replacing proprietary agents (New Relic, Datadog agents). It decouples your code from the vendor. You instrument with OTel SDK, and the OTel Collector exports to Datadog/Jaeger/Sentry/whatever. You can switch vendors by changing YAML, not Go code.

---

### 258. How do you log correlation IDs between services?
"I generate a unique `Correlation-ID` (UUID) at the ingress.

I pass it in the `context`.
I write a custom Logger wrapper that automatically pulls the ID from the context and adds it to the log entry.
`logger.Info(ctx, "processing payment")`.
Output: `{"msg": "processing payment", "correlation_id": "a1-b2-c3"}`.
This ties all logs from all services together."

#### Indepth
Propagate headers! When Service A calls Service B, it must copy `X-Correlation-ID` from the incoming request to the outgoing request. Middleware is the rightful place for this. If you miss one link in the chain, the trace is broken and debugging becomes a nightmare.

---

### 259. How would you handle distributed transactions in Go?
"I act as if they don't exist. Two-Phase Commit (2PC) is a scalability killer.

I prefer **Eventual Consistency**.
I write to my local DB and publish an event to Kafka.
If I need strong guarantees, I use the **Outbox Pattern**:
1.  Insert data into `users` table.
2.  Insert event into `outbox` table.
3.  Commit transaction.
4.  Background worker reads `outbox` and publishes to Kafka."

#### Indepth
The Outbox Pattern solves the "Dual Write Problem" (write to DB + publish to Kafka). Without it, if you write to DB and then crash before publishing, your system is inconsistent. By writing the event to the *same* DB transaction, you guarantee atomicity. The delivery to Kafka happens later (eventually).

---

### 260. How to deal with partial failures in distributed systems?
"I assume things will fail.

1.  **Timeouts**: Never wait forever.
2.  **Retries**: Exponential backoff (but cap it to avoid storms).
3.  **Circuit Breaker**: If Payment Service fails 50% of the time, I stop calling it for 30s to let it recover.
4.  **Graceful Degradation**: If Recommendations are down, show a static 'Popular Items' list instead of a 500 error."

#### Indepth
Circuit Breakers have three states: Closed (Normal), Open (Fail Fast), and Half-Open (Testing). In Half-Open state, let 1 request through. If it succeeds, close the breaker (recover). If it fails, open it again. This prevents a recovering service from getting instantly hammered back into oblivion.
