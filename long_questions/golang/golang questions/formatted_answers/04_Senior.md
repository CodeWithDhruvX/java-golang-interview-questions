# Senior Level Golang Interview Questions

## From 13 Microservices gRPC

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


## From 17 DevOps Containers

# 🟡 **321–340: DevOps, Docker, and Deployment**

### 321. How do you containerize a Go application?
"I write a **Dockerfile** that leverages Go's static compilation.

I almost always use a **multi-stage build**.
Stage 1 (Builder): Uses `golang:alpine`. I copy `go.mod`, download dependencies, and build the binary using `go build -o app`.
Stage 2 (Runner): Uses `gcr.io/distroless/static` or `scratch`. I copy *only* the compiled binary from Stage 1.

This results in a production image that is incredibly small (often <15MB) and secure because it lacks a shell or package manager."

#### Indepth
Always use `ENTRYPOINT` instead of `CMD` for the binary. `ENTRYPOINT ["/app"]` makes the container executable. `CMD` provides default arguments. This allows users to run `docker run my-image --help` and have the `--help` flag passed directly to your Go binary.

---

### 322. What is a multi-stage Docker build and how does it help with Go?
"A multi-stage build allows me to use a heavy image for compilation and a lightweight image for execution, all in one Dockerfile.

For Go, this is critical. The `golang` image (with GCC, Git, etc.) is ~800MB. My binary is ~10MB.
If I shipped the builder image, I’d be wasting 790MB of space and exposing a massive attack surface. Multi-stage builds strip all that away, leaving just the executable."

#### Indepth
You can reuse the Go Module Cache in multi-stage builds. Mount a cache volume: `RUN --mount=type=cache,target=/go/pkg/mod go build ...`. This speeds up subsequent builds by 10x because it doesn't re-download the internet every time you change a single line of code.

---

### 323. How do you reduce the size of a Go Docker image?
"I target the **scratch** base image, which is an empty file system.

To make this work, I compile with `CGO_ENABLED=0` to ensure a statically linked binary (no dependency on `glibc`).
I also use linker flags `-ldflags="-s -w"` to strip debug information, which shrinks the binary size by about 25%.

The final result is an image that is literally the size of the binary itself. You can't get smaller than that."

#### Indepth
For extreme size reduction, you can use **UPX** to compress the binary. It can shrink a 10MB Go binary to 3MB. However, it increases startup time (decompression overhead) and can trigger false positives in antivirus software. Generally, the standard strip (`-s -w`) is sufficient.

---

### 324. How do you handle secrets in Go apps deployed via Docker?
"I strictly follow the rule: **Never bake secrets into the image.**

I inject them at runtime as **Environment Variables**.
In Docker Compose, I use an `.env` file. In Kubernetes, I use `Secret` objects mapped to env vars.

My Go app reads them using `os.Getenv("DB_PASS")`. This separation of config and code allows the same image to run in Dev, Staging, and Prod safely."

#### Indepth
Avoid passing secrets as **Build Args** (`docker build --build-arg SECRET=...`). They persist in the Docker image history and can be recovered by anyone who checks `docker history`. Use **Docker Secrets** or runtime environment variables only.

---

### 325. How do you use environment variables in Go?
"I use the standard library `os.Getenv(key)`.

However, raw `Getenv` is limited because it returns an empty string if the key is missing.
For production apps, I use a configuration library like **kelseyhightower/envconfig** or **Viper**.
These libraries automatically map environment variables (like `APP_PORT=8080`) to struct fields (`Config.Port`), handling type conversion and default values for me."

#### Indepth
When using env vars, define a precedent order. **Flag > Env > Config File > Default**. This allows you to override a specific setting (like `LOG_LEVEL=debug`) quickly when debugging a container without rebuilding the image or changing the deployment implementation.

---

### 326. How do you compile a static Go binary for Alpine Linux?
"I must explicitly disable CGO.

`CGO_ENABLED=0 GOOS=linux go build`.
Alpine Linux uses `musl` instead of the standard `glibc`. If I rely on Cgo, the binary will crash on Alpine with 'no such file or directory'. By disabling Cgo, I force the compiler to use pure Go implementations for DNS and User lookups, making the binary truly portable across any Linux distribution."

#### Indepth
Alpine's `musl` libc has some quirks compared to `glibc`, especially with DNS resolution. Startups sometimes fail in Alpine due to these differences. If you hit weird DNS issues, switch to `gcr.io/distroless/static-debian11`, which is slightly larger but uses standard glibc compatibility.

---

### 327. What is `scratch` image in Docker and why is it used with Go?
"`scratch` is a special, empty Docker image. It contains absolutely nothing—no shell, no `/bin/ls`, no libraries.

It is the gold standard for Go deployments because Go binaries don't need an OS to run.
The security benefit is massive: even if an attacker finds a vulnerability in my app, they cannot 'shell out' to run commands because there makes no shell. The only caveat is needing to manually copy SSL certificates (`/etc/ssl/certs`) if I make HTTPS calls."

#### Indepth
Another missing piece in `scratch` is Timezone data. If your app relies on `time.LoadLocation("America/New_York")`, it will panic. You must manually copy `/usr/share/zoneinfo` from the builder stage to the scratch image.

---

### 328. How do you manage config files in Go across environments?
"I follow the **12-Factor App** principles, so valid config comes from the environment, not files.

However, for complex nested config, I use **Viper**.
It allows a layered approach:
1.  Read `config.yaml` (Base defaults).
2.  Override with Environment Variables (`VIPER_DB_HOST`).
3.  Override with Command Line flags.

This gives me the convenience of a file for local dev and the flexibility of env vars for Kubernetes."

#### Indepth
If you use Kubernetes ConfigMaps mounted as files, Viper allows **Hot Reloading** (`viper.WatchConfig()`). It detects when the file changes and updates the struct. This allows you to change the log level to DEBUG in production without restarting the pod!

---

### 329. How do you build Go binaries for different OS/arch?
"Cross-compilation is built right into the `go` tool.

I just set the `GOOS` and `GOARCH` environment variables.
To build for a standard server: `GOOS=linux GOARCH=amd64 go build`.
To build for a new MacBook: `GOOS=darwin GOARCH=arm64 go build`.

I don't need to install any extra toolchains or compilers. It’s one of Go’s superpowers given how painful this is in C++."

#### Indepth
Be careful with **CGO**. If you import `net` (which uses CGO for DNS by default on some OSs) or `sqlite3`, simple cross-compilation breaks. You typically need a designated "Cross-Compiler Docker Image" provided by projects like `goreleaser/goreleaser-cross` to handle the C dependencies.

---

### 330. How do you use GoReleaser?
"**GoReleaser** is my automation tool for shipping binaries.

I define a `.goreleaser.yaml` file. When I push a new Git tag (e.g., `v1.0.0`), GitHub Actions triggers GoReleaser.
It automatically builds binaries for every OS/Arch combination, zips them, generates checksums, creates a GitHub release, and even creates Homebrew recipes or pushes Docker images. It turns a manual 2-hour process into a background job."

#### Indepth
GoReleaser also supports **Snapshot Releases**. You can run `goreleaser release --snapshot --rm-dist` locally to build all artifacts and verify your release process works *without* actually pushing anything to GitHub. This is crucial for testing your build pipeline.

---

### 331. What is a Docker healthcheck for a Go app?
"It’s a command that runs inside the container to verify the app is working.

In my `Dockerfile`, I add:
`HEALTHCHECK --interval=30s CMD curl -f http://localhost:8080/health || exit 1`.

My Go app exposes a simple `/health` endpoint that returns 200 OK. If the database connection drops, I might return 500, causing Docker (or Kubernetes) to mark the container as unhealthy and restart it."

#### Indepth
Since `scratch` images don't have `curl`, you cannot use `CMD curl ...`. You must compile a tiny standalone binary (like `grpc-health-probe` or your own 50-line Go client) and copy it into the image to act as the healthheck executable.

---

### 332. How do you log container stdout/stderr from Go?
"I simply write to `os.Stdout` (for info) and `os.Stderr` (for errors).

In containerized environments, we **do not** write to log files.
The container runtime (Docker Engine) captures these streams. I then use a log shipper like **Fluentd** or **Filebeat** to read the Docker logs and forward them to a central system like Elasticsearch. This decouples the application from the logging infrastructure."

#### Indepth
Multi-line logs (like Stack Traces) are a pain in containers because Docker treats each line as a separate log event. Use structured logging (JSON) so that the entire stack trace is encapsulated in a single `"error": "..."` field, keeping the log entry atomic.

---

### 333. How do you set up autoscaling for Go services?
"I use the **Horizontal Pod Autoscaler (HPA)** in Kubernetes.

I define a CPU target, say 60%.
Because Go uses specialized goroutines, CPU usage is a very reliable proxy for load. When traffic spikes, CPU rises, and the HPA spins up new Pods.
Go apps start in milliseconds, so this scaling is highly responsive, unlike Java apps which might struggle with cold-start times during a surge."

#### Indepth
Don't just scale on CPU. Scale on **Custom Metrics** like "RabbitMQ Queue Depth". If the queue has 10,000 pending jobs, but CPU is low (because the workers are just waiting on IO), CPU scaling won't trigger. KEDA (Kubernetes Event-driven Autoscaling) is the standard tool for this.

---

### 334. How would you containerize a gRPC Go service?
"It’s similar to HTTP, but I expose the gRPC port (e.g., 50051).

The tricky part is the health check. standard `curl` doesn't speak `gRPC`.
I install **grpc_health_probe** in the container image.
My Go app implements the standard gRPC Health Checking Protocol. The probe calls this service, ensuring that not only is the TCP port open, but the gRPC server is ready to accept requests."

#### Indepth
When using `grpc_health_probe`, make sure to handle the **Shutdown Gracefully**. When a Pod terminates, it should set its health status to `SERVING_STATUS_NOT_SERVING` immediately, so the Load Balancer stops sending traffic *while* the existing requests finish (Grace Period).

---

### 335. How to deploy Go microservices in Kubernetes?
"I define three core Kubernetes objects:
1.  **Deployment**: Manages the Pods (replicas, image version, resources).
2.  **Service**: Provides a stable ClusterIP and DNS name for internal discovery.
3.  **ConfigMap/Secret**: Injects configuration via environment variables.

For zero-downtime deployments, I use a R ollingUpdate strategy, which Kubernetes handles natively because Go apps shut down gracefully on `SIGTERM`."

#### Indepth
Configure **Pod Disruption Budgets (PDB)**. This ensures that during a cluster upgrade (node draining), Kubernetes won't kill all your replicas at once. `minAvailable: 1` guarantees at least one pod is always up, preserving 100% availability.

---

### 336. How do you write Helm charts for a Go app?
"**Helm** allows me to template my Kubernetes manifests.

Instead of hardcoding `replicas: 3`, I write `replicas: {{ .Values.replicaCount }}`.
This lets me use the exact same templates for Dev (1 replica) and Prod (10 replicas) just by supplying a different `values.yaml` file. It packages my Go app deployment logic into a reusable artifact."

#### Indepth
Helm Charts can be complex. Keep them simple. Use **Library Charts** or **Subcharts** for shared logic (like standard probes, common labels, sidecars) to avoid copy-pasting code across 50 different microservice charts.

---

### 337. How do you monitor a Go service in production?
"I focus on the **RED** method: **Rate** (RPM), **Errors** (5xx), and **Duration** (Latency).

I instrument the code using **Prometheus**. I use middleware that wraps every HTTP handler, measuring how long it takes and recording the status code.
Prometheus scrapes my service’s `/metrics` endpoint every 15 seconds, and I visualize the RED metrics on Grafana dashboards."

#### Indepth
Be careful with **Cardinality Explosion**. If you add a label `path="/users/123"`, you create a new metric series for *every user*. This will crash Prometheus. Always normalize paths: `path="/users/:id"`.

---

### 338. How do you use Prometheus with a Go app?
"I use the `prometheus/client_golang` library.

I define metrics like `var httpRequests = promauto.NewCounter(...)`.
In my code, I call `httpRequests.Inc()`.
I then expose a standard HTTP handler `promhttp.Handler()` on a dedicated port (e.g., `:9090`).
This 'pull model' is robust because if my app is under heavy load, monitoring doesn't block critical request processing."

#### Indepth
For short-lived Batch Jobs (cron jobs), the Pull Model fails (the job dies before Prometheus scrapes it). Use the **Pushgateway**. The Go job pushes metrics to the Gateway just before exiting, and Prometheus scrapes the Gateway.

---

### 339. How do you enable structured logging in production?
"I use **Zap** or **slog** (Go 1.21+) with a JSON handler.

`logger := slog.New(slog.NewJSONHandler(os.Stdout))`
This outputs log lines as JSON objects: `{"level":"info", "msg":"user login", "user_id":42}`.
Structure is non-negotiable in production because it allows me to index and query logs by specific fields in tools like Splunk or Datadog, rather than grepping text."

#### Indepth
Include the **Trace ID** and **Span ID** in every log line automatically. This allows you to correlate "Error in DB" with "User Request X". In `slog`, you can use `slog.With("trace_id", ctx.TraceID())` to inject this context into the logger.

---

### 340. How do you handle log rotation in containerized Go apps?
"**I do not handle it in the app.**

The application should effectively write infinite logs to Stdout.
Log rotation is the responsibility of the execution environment. In Docker, I configure the `json-file` logging driver with `max-size` and `max-file` options.
If I tried to handle rotation inside the Go app, I’d run into concurrency issues and risk losing logs during rotation."

#### Indepth
In legacy environments (VMs) where you *must* log to a file, use an external tool like `logrotate` (standard Linux utility) combined with the `SIGHUP` signal. When `logrotate` moves the file, it sends `SIGHUP` to your app, telling it to reopen the file handle.


## From 18 Streaming Async

# 🔵 **341–360: Streaming, Messaging, and Asynchronous Processing**

### 341. How do you consume messages from Kafka in Go?
"I use the `sarama` library or `segmentio/kafka-go`.

I always implement a **Consumer Group**.
This allows me to scale horizontally. I start 10 replicas of my service, and Kafka automatically assigns a subset of partitions to each.
In the code, I loop over `claim.Messages()`. Crucially, I only mark the message (`session.MarkMessage`) *after* I have successfully processed it. This ensures I never lose data if my pod crashes."

#### Indepth
Sarama's default configuration is unsafe for high reliability. You must set `Producer.RequiredAcks = WaitForAll` to ensure the broker actually wrote the message to disk. On the consumer side, enable `Rebalance.Strategy = Sticky` to reduce the "stop-the-world" pause when a new consumer joins the group.

---

### 342. How do you publish messages to a RabbitMQ topic?
"I use the `streadway/amqp` library (or the new `rabbitmq/amqp091-go`).

1.  Connect and open a **Channel**.
2.  Declare an **Exchange** (Topic).
3.  Publish with a **Routing Key** (e.g., `order.created.us`).
RabbitMQ uses this key to route the message to the correct queues.
Since AMQP connections are stateful and fragile, I always wrap the publisher in a struct that handles **automatic reconnection** seamlessly."

#### Indepth
RabbitMQ has "Publisher Confirms". When you publish, the broker sends back an ACK. If you don't wait for this ACK, you might lose messages if the broker crashes instantly (or if the TCP packet drops). Use `channel.NotifyPublish()` to listen for these confirmations for 100% durability.

---

### 343. What is the idiomatic way to implement a message handler in Go?
"I define a simple interface.
`type Handler interface { Handle(ctx context.Context, msg []byte) error }`.

My consumer code is generic: it reads bytes and calls `handler.Handle`.
If `Handle` returns an error, the consumer decides the retry strategy (Nack, Requeue, or DLQ).
This separation means I can unit test my business logic (`Handle`) without needing to mock complex Kafka/RabbitMQ structs."

#### Indepth
Decorate your handler! `LoggingMiddleware(RetryMiddleware(Handle))`. This creates a chain of responsibility. Just like HTTP middleware, this allows you to standardise observability, panic recovery, and tracing across all your async workers without polluting the business logic.

---

### 344. How would you implement a worker pool pattern?
"I use **Buffered Channels**.

`jobs := make(chan Job, 100)`.
I start `N` goroutines (workers) that range over this channel.
`for job := range jobs { process(job) }`.
When I want to process work, I push to `jobs`. If the buffer fills up, the sender blocks. This provides natural **backpressure**, preventing memory exhaustion if the workers fall behind."

#### Indepth
Worker pools are great, but `errgroup` is often simpler for "scatter-gather" workflows. If you need to process 10 files in parallel and fail if *any* of them fail, `errgroup.WithContext` manages the goroutines and cancellation propagation for you automatically.

---

### 345. How do you use the `context` package for cancellation in streaming apps?
"Context is the **Kill Switch**.

When I start a stream consumer, I pass it a context.
In the inner loop (e.g., reading from Kafka), I always `select` on `ctx.Done()`.
`select { case msg := <-stream: process(msg); case <-ctx.Done(): return }`.
This allows me to shut down the stream instantly (stop waiting for new messages) when the application receives a `SIGTERM`."

#### Indepth
For network calls inside the loop, use `context.WithTimeout`. `req, _ := http.NewRequestWithContext(ctx...)`. If the main parent `ctx` is canceled (shutdown), the HTTP request aborts immediately. This ensures your consumers don't hang for 60s finishing a request during deployment.

---

### 346. How do you retry failed messages in Go?
"I use **Exponential Backoff**.

If a message fails, I don't retry immediately. I wait: `1s, 2s, 4s, 8s`.
I use a library like `cenkalti/backoff` for this.
However, I can't block the valid messages behind the bad one forever. After 3-5 attempts, I move the failed message to a **Dead Letter Queue (DLQ)** and move on. This keeps the pipeline flowing."

#### Indepth
Don't just log retry failures—increment a metric `job_failures_total`. Integrate your Circuit Breaker here. If the destination service is down, waiting 8 seconds (backoff) 100 times is wasteful. Trip the breaker and fail fast to the DLQ immediately.

---

### 347. What is dead-letter queue and how do you use it?
"A **DLQ** is a 'parking lot' for bad messages.

If a message causes a crash or fails validation repeatedly (Poison Pill), I assume it's unprocessable.
I publish it to a separate topic `project-dlq`.
I have alarms on this queue. Later, a human (or a script) can inspect these payloads, fix the bug, and re-inject them into the main queue."

#### Indepth
Automate the DLQ "Replay". Write a small CLI tool that reads from the DLQ scope and publishes back to Main Topic. Often, the failure was transient (DB down), and simply replaying 2 hours later fixes everything without code changes.

---

### 348. How do you handle idempotency in message consumers?
"Idempotency means 'processing the same message twice has the same effect as once'.
Message queues often deliver duplicates (At-Least-Once).

I handle this by assigning a **Unique ID** to every message at source.
In the consumer, I check a store (Redis/Postgres) using `SETNX key_id`. If the key exists, I skip processing. This simple check makes the system robust against duplicate deliveries."

#### Indepth
For "Exactly-Once" processing without Redis, use **Bloom Filters** for a fast, probabilistic check before hitting the DB. Or better, design your DB schema to handle duplicates naturally: `INSERT ... ON CONFLICT DO NOTHING`. Relying on the DB constraint is the most robust deduplication.

---

### 349. How do you implement exponential backoff in Go?
"I typically use a loop with a `time.Sleep`.

`delay := 100 * time.Millisecond`
`for i := 0; i < retries; i++ { err := work(); if err == nil { return }; time.Sleep(delay); delay *= 2 }`.

Crucially, I add **Jitter** (randomness). Instead of exactly 200ms, I sleep `200ms + rand(50ms)`. This prevents the 'Thundering Herd' problem where 1000 failing instances all hit the database again at the exact same millisecond."

#### Indepth
Use the "Decorrelated Jitter" algorithm. Instead of just adding random noise, the formula `sleep = min(cap, random(base, sleep * 3))` allows the backoff to fluctuate wildly, which statistically spreads out the load much better than simple "Equal Jitter".

---

### 350. How do you stream logs to a file/socket in real-time?
"I implement the `io.Writer` interface.

If I'm writing to a slow socket (like Logstash), I wrap it in an **Async Writer**.
The logger writes to a channel (fast, non-blocking).
A background goroutine reads the channel and writes to the socket.
This ensures that a network glitch in the logging infrastructure effectively typically drops logs rather than freezing the main application logic."

#### Indepth
Use a **Ring Buffer** for the async writer. If the buffer is full (Logstash is down), the writer should *overwrite old logs* or drop new ones (Non-Blocking mode). Never let logging block the application. It is better to lose logs than to take down the service (`os.Stderr` is usually safe though).

---

### 351. How do you work with WebSockets in Go?
"I use `gorilla/websocket` (or `nhooyr.io/websocket`).

I upgrade the HTTP request to a WebSocket connection.
Then I enter a `for { conn.ReadMessage() }` loop.
**Critical detail**: `conn.WriteMessage` is **not thread-safe**. If I have multiple goroutines trying to send data to the same user, I must protect the write with a Mutex or standardizing on a single 'writer goroutine' fed by a channel."

#### Indepth
WebSockets need **Keep-Alives**. Intermediate load balancers (nginx) kill idle TCP connections after 60s. You must implement a "Ping/Pong" loop. The server sends a Ping every 30s. If the client doesn't respond with Pong within 10s, close the connection.

---

### 352. How do you handle bi-directional streaming in gRPC?
"I define the RPC: `rpc Chat(stream Note) returns (stream Note)`.

In the handler, I get a `stream` object.
I need concurrency here.
1.  I spawn a goroutine to `stream.Recv()` in a loop (handle incoming).
2.  I use the main thread (or another goroutine) to `stream.Send()` (push outgoing).
The connection stays open until one side calls `Close()`."

#### Indepth
Flow Control is automatic in gRPC (HTTP/2). If the receiver reads slower than the sender, the window fills up, and the sender blocks. However, you can check `ctx.Done()` in the sender loop to detect if the client disconnected, preventing "ghost" streams from leaking resources.

---

### 353. What is Server-Sent Events and how is it done in Go?
"**SSE** is a one-way channel from Server to Browser over standard HTTP.

I set the header `Content-Type: text/event-stream`.
Then I cast the `http.ResponseWriter` to `http.Flusher`.
I write data `fmt.Fprintf(w, "data: %s\n\n", msg)` and immediately call `flusher.Flush()`.
If I don't flush, Go buffers the response, and the client sees nothing until the buffer fills up."

#### Indepth
Browser limits! Browsers (Chrome) limit concurrent connections to the same domain (max 6). If you open 6 SSE tabs, the 7th will hang. HTTP/2 solves this (multiplexing), but if you are on HTTP/1.1, you must shard domains (`s1.api.com`, `s2.api.com`) to bypass this limit.

---

### 354. How do you manage fan-in/fan-out channel patterns?
"**Fan-Out**: I launch multiple worker goroutines reading from the **same** channel. The runtime distributes the tasks.

**Fan-In**: I have multiple result channels merging into one.
I launch a goroutine for each input channel that forwards to the output.
I use `sync.WaitGroup` to track when all inputs are closed, so I can safely close the output channel."

#### Indepth
Error handling in Fan-In is tricky. If one worker fails, do you stop everything? Use `errgroup` (again). It cancels the Context for all other workers immediately upon the first error, ensuring you fail fast and don't waste CPU on a doomed batch.

---

### 355. How would you implement throttling on async tasks?
"I use the **Semaphore** pattern with a buffered channel.

`sem := make(chan struct{}, 10)`.
Before starting a goroutine: `sem <- struct{}{}`.
Inside the goroutine (defer): `<-sem`.
This limits concurrent execution to 10. It’s much lighter than a full worker pool if the tasks are just quick computations."

#### Indepth
The `golang.org/x/sync/semaphore` package provides a Weighted Semaphore. This allows advanced limiting: "Heavy tasks take 5 slots, light tasks take 1 slot". This gives you fine-grained control over resource consumption compared to a simple channel of structs.

#### Indepth
The `golang.org/x/sync/semaphore` package provides a Weighted Semaphore. This allows advanced limiting: "Heavy tasks take 5 slots, light tasks take 1 slot". This gives you fine-grained control over resource consumption compared to a simple channel of structs.

---

### 356. How do you avoid data races when consuming messages?
"**Isolation**.
Each worker goroutine should work on its own local data.
I never share a map or slice between workers.
If they need to aggregate results (e.g., 'count total errors'), I use `atomic.AddInt64` or, better yet, send the result to a dedicated 'Aggregator' goroutine via a channel."

#### Indepth
Run the race detector (`-race`) in your Integration Tests too! Many race conditions only appear when real network latency and IO are involved. It slows down tests, but finding a data race in a payment processing pipeline is worth the CPU cycles.

---

### 357. How would you implement a message queue from scratch in Go?
"In memory: Use a buffered channel.
`queue := make(chan Job, 1000)`.

For durability: **Write-Ahead Log (WAL)**.
Before pushing to the channel, I append the job to a file on disk.
On startup, I read the file to repopulate the channel. This is basically how Kafka works fundamentally."

#### Indepth
For the "Ring Buffer" implementation, `github.com/dgraph-io/ristretto` (cache) uses a high-performance ring buffer. Study its source code. It uses `atomic` counters and power-of-two sizing (`mask = size - 1`) to avoid expensive modulo operators (`%`) during the hot path.

---

### 358. How do you implement ordered message processing in Go?
"Parallelism breaks ordering. To allow parallelism *and* ordering, I use **Sharding**.

I hash the 'Entity ID' (e.g., specific `UserID`).
`shardID := hash(userID) % numWorkers`.
Each worker has its own channel. Messages for User A always go to Worker 1. Messages for User B go to Worker 2.
This guarantees User A's events are processed in order, while allowing User A and B to run in parallel."

#### Indepth
This is the **Partition Key** strategy. Be careful of "Hot Partitions". If Justin Bieber joins your app, and all his events go to Shard 1, Shard 1 will lag while Shard 2-10 are idle. You need a "Virtual Node" consistent hashing strategy to rebalance hot keys.

---

### 359. How do you handle large stream ingestion (100K+ msgs/sec)?
"I focus on **Batching** and **Allocation reduction**.

1.  **Read Batch**: Read 100 messages at a time from Kafka.
2.  **Process Batch**: Use `sync.Pool` to reuse message objects/buffers.
3.  **Write Batch**: Insert 100 rows into Postgres in one transaction.
Processing 1 message at a time is the death of performance. Batching reduces overhead by orders of magnitude."

#### Indepth
Allocations in the hot loop kill throughput. Use `sync.Pool` to reuse the structs used for unmarshalling JSON. Also, use `json.Scanner` or `easyjson` to avoid reflection overhead. Profiling usually shows `mallocgc` as the bottleneck in stream processors.

---

### 360. How do you persist in-flight streaming data?
"I treat the Stream (Kafka) as the **Source of Truth**.

1.  Read message (offset 100).
2.  Process and write to DB.
3.  **Commit Offset 100**.
If I crash before step 3, I re-read message 100 on restart.
This is 'At-Least-Once' semantics. My DB write must be idempotent (e.g., `INSERT ON CONFLICT DO NOTHING`) to handle the duplicates."

#### Indepth
Avoid storing offset in Zookeeper/Kafka if possible. Store the offset **in the same DB transaction** as your data. `INSERT INTO users ...; UPDATE offsets SET val=101`. This effectively gives you **Exactly-Once** processing because the data and offset commit atomically.


## From 19 Architecture SystemDesign

# 🟢 **361–380: Architecture and System Design**

### 361. How do you design a scalable Go application?
"I adhere to the **Shared Nothing Architecture**.

Each instance of my Go service is stateless.
Any state (User Session, Shopping Cart) is pushed to a persistent store (Postgres) or a fast cache (Redis).
This allows me to scale horizontally: if traffic doubles, I just double the number of pods. I also heavily use **Message Queues** (Kafka) to decouple components so a load spike becomes a backlog, not an outage."

#### Indepth
Statelessness is a spectrum. Even "stateless" services often have local caches (for repeated DB lookups). Ensure you have a **Cache Invalidation Strategy** (e.g., TTLs or Pub/Sub events) so that one pod doesn't serve stale data while others serve new data.

---

### 362. What is Clean Architecture in Go?
"It’s a layered approach that enforces the **Dependency Rule**: dependencies only point inward.

*   **Entities (Core)**: Pure business objects (`User`). No SQL tags, no JSON tags.
*   **Use Cases**: Application logic (`CreateUser`). Depends only on Entities.
*   **Adapters**: HTTP handlers, SQL implementations. Depend on Use Cases.
*   **Infrastructure**: The DB driver, the Web framework.
This ensures I can swap Postgres for MongoDB without changing a single line of my core business logic."

#### Indepth
Clean Architecture comes with a cost: **Boilerplate**. You will map `UserDTO` (API) to `User` (Domain) to `UserModel` (DB). For simple CRUD apps, this is overkill. Don't be afraid to use a "Vertical Slice Architecture" (Transaction Script) for simple features, and upgrade to Clean Architecture only for complex domains.

---

### 363. How do you implement Domain-Driven Design (DDD) in Go?
"I organize code by **Domain**, not by Layer.

Instead of `packages/models` and `packages/controllers`, I have:
`packages/billing` (contains its own models, logic, and repo interfaces).
`packages/inventory`.
I define **Aggregates** (Transaction Boundaries) and enforce that changes to an aggregate only happen through its methods. This prevents 'anemic domains' where logic leaks into the service layer."

#### Indepth
A common DDD mistake in Go is creating "God Aggregates". If your `User` aggregate contains `Orders`, `Payments`, and `Reviews`, you seal the database row for `User` every time you add a Review. Keep aggregates small. `Order` should be its own aggregate and reference `UserID` by ID, not by embedding.

---

### 364. What is the Hexagonal Architecture (Ports and Adapters)?
"It’s about decoupling the app from the technology.

**Ports**: Interfaces defined by the Core.
*   Driving Port: `OrderService` interface (called by HTTP).
*   Driven Port: `OrderRepository` interface (implemented by SQL).

**Adapters**: The glue code.
*   `HttpHandler` adapter plugs into the Driving Port.
*   `PostgresRepo` adapter plugs into the Driven Port.
This allows me to run the entire application logic in a unit test with a `MemoryRepo` adapter, in milliseconds."

#### Indepth
Hexagonal architecture makes **Contract Testing** easier. You can write a test suite for the `Repository` interface and run it against both the `MemoryRepo` (for unit tests) and the `PostgresRepo` (for integration tests). This ensures your fake and real implementations actually behave identically.

---

### 365. How do you handle configuration management in distributed systems?
"I treat config as **Immutable**.

I use environment variables (12-Factor App) injected by Kubernetes ConfigMaps.
I read them on startup using `viper` or `kelseyhightower/envconfig`.
I avoid 'hot reloading' config because it leads to split-brain states (half the pods have new config, half have old). If I need to change config, I deploy a new revision of the pod."

#### Indepth
Feature Flags are a form of dynamic config. Use a system like **LaunchDarkly** or **Unleash** to toggle features on/off without redeploying. This decouples "Deployment" (moving code) from "Release" (enabling features), allowing for safer rollouts.

---

### 366. How do you design for failure in Go systems?
"I assume the network is hostile.

1.  **Timeouts**: Every `http.Client` and `sql.DB` call has a strict context timeout.
2.  **Retries**: I use exponential backoff with jitter for transient errors.
3.  **Circuit Breakers**: If the Payment Service is failing 50% of the time, I stop calling it for 30s to let it recover.
4.  **Bulkheads**: I isolate thread pools so a slow feature doesn't starve the whole app."

#### Indepth
"Fail Open" vs "Fail Closed". If your Fraud Check service is down, do you block the user (Fail Closed) or let them pass (Fail Open)? For high-value transactions, Fail Closed. For low-stakes features (Recommendations), Fail Open. Document this decision explicitly.

---

### 367. What is the CQRS pattern and when to use it?
"**Command Query Responsibility Segregation**.

I split my application into two parts:
*   **Write Side (Command)**: Handles `CreateOrder`. Sticktly consistent, normalized data.
*   **Read Side (Query)**: Handles `GetOrderHistory`. Eventually consistent, denormalized (e.g., Elasticsearch or a specific Read View table).
I use it when the read load is massive (1000:1 read:write ratio) or when the query patterns are too complex for the normalized write schema."

#### Indepth
CQRS enables **polyglot persistence**. The Write model can be a normalized PostgreSQL (3rd Normal Form) for data integrity. The Read model can be a denormalized MongoDB document, or a Redis Hash, optimized specifically for the "User Profile" screen. You project events from SQL to Mongo asynchronously.

---

### 368. How do you implement Event Sourcing in Go?
"Instead of storing the 'Current State' (Balance: $50), I store the **Sequence of Events** (Deposited $100, Withdrew $50).

To get the balance, I replay the events: $0 + $100 - $50 = $50.
In Go, I append these events to an immutable log (Kafka/EventStore).
It gives me a perfect audit trail and allows 'Time Travel' debugging, but it adds massive complexity (snapshots, schema evolution), so I only use it for critical financial/audit systems."

#### Indepth
Event Sourcing has a "snapshots" problem. To get the current balance, you can't replay 1 million events every time. You verify a Snapshot every 100 events. The generic `Aggregate.Fold()` logic in Go is typically: `state = Snapshot + apply(events_since_snapshot)`.

---

### 369. How do you handle database migrations in a microservices architecture?
"Each service **owns** its schema.

Service A effectively checks out Service A's DB. Service B cannot touch it.
I run migrations (using `golang-migrate`) as an **Init Container** in Kubernetes.
When a pod starts, it upgrades the schema *before* the app boots.
I strictly write **Backward Compatible** migrations (Add Column -> Deploy Code -> Backfill -> Remove Column) to support Zero Downtime deployments."

#### Indepth
Distributed locks are dangerous for migrations. If two pods try to run migrations simultaneously, they might corrupt the DB. Use `advisory_locks` (Postgres) or a dedicated K8s Job `helm install --wait`. Never let the application pod run migrations on startup in a multi-replica setup.

---

### 370. What is the Strangler Fig pattern?
"It’s how I kill a legacy Monolith.

1.  Put a Proxy (API Gateway) in front of the Monolith.
2.  Write a new Go Microservice for *one* feature (e.g., Search).
3.  Route `/search` traffic to the new Go service.
4.  Keep `/users`, `/billing` going to the Monolith.
5.  Repeat until the Monolith is empty.
This minimizes risk compared to a 'Big Bang' rewrite."

#### Indepth
The hardest part of Strangler Fig is **Data Sync**. If the monolith and microservice share the same DB, you have a "Distributed Monolith". If they split the DB, you need a sync mechanism (CDC or Dual Write) until the migration is complete. Prefer splitting the DB early if possible.

---

### 371. How do you design API rate limiting for high traffic?
"I use the **Token Bucket** algorithm backed by Redis.

Each request decrements a counter in Redis (`INCR key` is atomic).
Keys are sharded by UserID or IP.
If the count > Limit, I return **HTTP 429**.
I prefer Redis because it works across distributed instances. For ultra-high scale (DDOS protection), I rely on a specialized WAF (Cloudflare) at the edge rather than Go middleware."

#### Indepth
Token bucket is bursty. If you want smooth traffic, use **Leaky Bucket**. If you want a hard cap, use **Fixed Window**. If you want a hard cap without boundary issues (the "double spike" at minute boundaries), use **Sliding Window Log**, though it's memory expensive. **Sliding Window Counter** is the best memory/accuracy trade-off.

---

### 372. How do you manage distributed sessions?
"I go **Stateless** with JWTs.

The session data is signed and stored in the token on the client side. The server just verifies the signature.
If I absolutely need server-side state (e.g., to revoke a login instantly), I store the Session ID in **Redis**.
I never store sessions in application memory because Sticky Sessions break load balancing and complicate autoscaling."

#### Indepth
JWTs cannot be revoked easily. If you need revocation (e.g., "Sign out all devices"), you must keep a blacklist of revoked JTI (JWT IDs) in Redis, which checks on every request. This re-introduces state, negating some benefits of stateless JWTs, but it's a necessary trade-off for security.

---

### 373. What is the Outbox Pattern?
"It solves the 'Dual Write' problem (Write to DB + Publish to Kafka).

1.  Start SQL Transaction.
2.  Insert `User`.
3.  Insert `Msg` into an `outbox` table in the *same* DB.
4.  Commit Transaction.
5.  A background poller reads the `outbox` table and publishes to Kafka.
This guarantees that **if** the user is created, the event **will** be published, assuming eventually consistency."

#### Indepth
Debezium is a popular tool for this. It reads the Postgres Write-Ahead Log (WAL) and streams changes to Kafka. This assumes you don't even need an explicit "Outbox" table if you treat the domain table changes as events, but the Outbox table is safer for keeping internal schema private.

---

### 374. How do you handle idempotency in API design?
"I require an **Idempotency-Key** header (UUID) for critical POST requests.

On the server:
1.  Check Redis: 'Have I seen this Key?'
2.  If yes: Return the *cached response* from the previous success. Do not execute logic.
3.  If no: Execute logic, save response to Redis, return it.
This ensures that if a client retries a payment due to a network timeout, we don't charge them twice."

#### Indepth
Idempotency keys should expire (e.g., 24 hours). Also, the check must be atomic. Use `SET key val NX EX 86400` in Redis. If it returns false, the key exists—fetch the old response. If true, proceed. This avoids the race condition of "Check -> Logic -> Set".

---

### 375. How do you design for eventual consistency?
"I embrace the lag.

In the UI, I use **Optimistic Updates** (show 'Done' immediately implies success locally).
In the backend, I use **Sagas**.
If Service A updates, it emits an event. Service B listens and updates.
If Service B fails, it emits a 'Failed' event, and Service A executes a compensation logic (undo).
I monitor the 'Replication Lag' to ensure 'Eventually' doesn't mean 'Tomorrow'."

#### Indepth
A common pattern is "Read Your Own Writes". When a user updates their profile, pin them to the master DB (or the updated replica) for a few seconds via a cookie. This ensures they don't see their old profile name immediately after changing it, which builds user trust.

---

### 376. What is the Sidecar pattern?
"It’s deploying a helper container in the same Pod as my Go app.

Example: **Envoy Proxy**.
My Go app talks to `localhost:80`. Envoy intercepts it, handles mTLS, retries, and tracing, then forwards it to the destination.
This creates a **Service Mesh**. It keeps my Go code clean of infrastructure concerns like cert rotation or circuit breaking."

#### Indepth
Sidecars have resource overhead. If your Go app uses 10MB RAM but the Envoy sidecar uses 100MB, you are wasting money. For simple setups, "Proxyless Service Mesh" (gRPC library directly talking to the control plane) is emerging as a lighter alternative.

---

### 377. How to implement a circuit breaker in Go?
"I use `gobreaker` or `hystrix-go`.

I wrap unreliable calls:
`cb.Execute(func() error { return http.Get(...) })`.
If failure rate > 50%, the breaker **trips** (Open).
Subsequent calls fail immediately (Fast Failure) without touching the network.
After a sleep window, it lets one request through (Half-Open). If it succeeds, it resets (Closed)."

#### Indepth
State persistence is tricky. If you have 10 pods, each has its own local circuit breaker. If the DB is down, all 10 pods must individually trip their breakers (letting some requests fail). Distributed circuit breakers (using Redis) exist but add latency; usually, local breakers are fine.

---

### 378. How do you handle cascading failures?
"I use **Load Shedding** and **Timeouts**.

If my CPU is > 90%, I start rejecting requests immediately (HTTP 503) to save the remaining capacity for in-flight requests.
I also use **Timeouts** everywhere. If the DB is slow, my API shouldn't hang forever; it should timeout fast and free up the thread.
This prevents a failure in a low-priority service from taking down the critical path."

#### Indepth
Implement **Priority Queues**. If the system is overloaded, drop "Background Sync" traffic but keep "User Checkout" traffic alive. This requires checking the "Task Priority" header at the ingress or middleware layer and maintaining separate semaphores for each tier.

---

### 379. What is API Gateway versus Service Mesh?
"**API Gateway** (North-South): The front door. Handles external clients, AuthN, Rate Limiting, and Routing to microservices.
**Service Mesh** (East-West): The internal wiring. Handles mTLS, Retries, and Tracing *between* microservices.
I use a Gateway (Kong) to let users in, and a Mesh (Istio) to let services talk to each other safely."

#### Indepth
API Gateways can handle **AuthN** (Who are you? - validating the JWT signature) but should delegate **AuthZ** (Are you allowed to do this? - checking scopes/roles) to the service or an external policy engine like OPA (Open Policy Agent). The Gateway is too coarse-grained for complex business rules.

---

### 380. How do you secure data in transit and at rest?
"**Transit**: TLS 1.3 everywhere.
External: HTTPS (Let's Encrypt).
Internal: mTLS (Mutual Auth) so services prove identity to each other.

**Rest**: Encryption at the storage layer (AWS EBS / RDS encryption).
For sensitive columns (SSN/PII), I use **Envelope Encryption**: I encrypt the data with a generated Data Key (DEK) and store the DEK encrypted by a Master Key (KMS) alongside the data."

#### Indepth
Key Rotation is essential. With envelope encryption, you only rotate the Master Key (KMS). You don't need to re-encrypt terabytes of database rows (which are encrypted with DEKs). You just re-encrypt the DEKs with the new Master Key if you want full rotation, or simply use the new Master Key for new data.


## From 20 Troubleshooting Debugging

# 🟡 **381–400: Troubleshooting and Debugging**

### 381. How do you debugging a deadlock in Go?
"I check the **Stack Traces** of all goroutines.

When a Go program deadlocks (all goroutines asleep), the runtime crashes with a full dump.
I look for:
1.  `semacquire`: Waiting for a Mutex.
2.  `chan receive`: Waiting for a message.
If Goroutine A holds Lock 1 and waits for Lock 2, while Goroutine B holds Lock 2 and waits for Lock 1, that is my deadlock. To fix it, I enforce a strict Lock Ordering."

#### Indepth
Deadlocks are often deterministic but hard to reproduce. Use `go run -race` to catch them if valid, but better yet, analyze **Lock Ordering**. If Goroutine A locks `mu1` then `mu2`, and Goroutine B locks `mu2` then `mu1`, you have a potential deadlock. Always acquire locks in a consistent global order.

---

### 382. How do you analyze a memory leak in production?
"I take a **Heap Profile** from the running application.

`go tool pprof http://localhost:8080/debug/pprof/heap`.
I look at `inuse_space`. It shows me exactly which function is holding the most memory *right now*.
If I suspect a slow leak, I take two profiles 1 hour apart and use `pprof -diff_base`. This highlights the *growth* (delta), filtering out stable memory usage."

#### Indepth
A subtle leak: **Subslices**. `original := make([]int, 1000000); small := original[:10]`. If you keep `small` in memory, the entire backing array of 1 million integers is kept in memory! Use `slices.Clone()` (Go 1.21) to copy the data and drop the large backing array.

---

### 383. How do you debugging high CPU usage?
"I take a **CPU Profile** for 30 seconds.

The **Flame Graph** is my primary tool.
Wide bars mean 'time spent on CPU'.
Common culprits:
1.  **Serialization**: `json.Marshal` taking 40% of CPU.
2.  **GC Thrashing**: Failing to allocate memory fast enough.
3.  **Busy Loops**: A `for {}` loop without a `time.Sleep` or blocking call.
4.  **Regex**: Recompiling `regexp.MustCompile` inside a hot loop."

#### Indepth
If CPU is high but pprof shows nothing (all time in `runtime`), check `GODEBUG=schedtrace=1000`. You might have **Starvation** or excessive **Context Switching**. If you spawn 1 million short-lived goroutines per second, the scheduler overhead dominates the CPU.

---

### 384. What tools do you use for distributed tracing?
"I use **OpenTelemetry** with **Jaeger** or **Tempo**.

I ensure every request has a `Trace-ID`.
When a user says 'My request failed', I search for that ID.
The waterfall view shows me:
'API took 500ms... 300ms was waiting for Redis, 180ms was waiting for User Service, 20ms was logic'.
It pinpoints the bottleneck instantly across service boundaries."

#### Indepth
Tracing must be propagated! Use `context.Context` everywhere. If you drop the context in a goroutine (`go func() { ... }()`), you break the trace. Pass the context or create a detached context with the valid parent span ID to maintain the "Trace Waterfall".

---

### 385. How do you debug a panic in a production service?
"I rely on the **Panic Stack Trace**.

I catch the panic using `recover()` in my middleware.
I log the full stack trace to my centralized logging system (ELK/Sentry).
Without the stack trace, I'm guessing. With it, I know exactly which line caused the nil pointer or index out of range.
I fix the bug by adding a nil check or validating input, never by just suppressing the panic."

#### Indepth
`recover()` only works in the **same goroutine** as the panic. It does not catch panics from child goroutines. If you spawn `go func() { panic("boom") }()`, your server crashes even if the parent has a `recover`. You must add `defer recover()` to *every* goroutine you spawn.

---

### 386. How to use `delve` for debugging?
"**Delve** (`dlv`) is the Go debugger.

I use it locally for complex logic bugs.
`dlv debug main.go`.
*   `break main.go:42`: Set breakpoint.
*   `continue`: Run to specific line.
*   `print varName`: Inspect internal state.
*   `goroutines`: specific specific status of all blocked goroutines.
It’s much faster than adding `fmt.Println` and recompiling 50 times."

#### Indepth
Delve can attach to a running process: `dlv attach <PID>`. This is a lifesaver in staging environments where you can't just restart the app to add logs. You can inject non-breaking breakpoints (tracepoints) to print variables without stopping execution.

---

### 387. How do you debug race conditions that only happen occasionally?
"I use the **Race Detector**.

`go test -race -count=1000 ./mypackage/...`.
The standard race detector instruments memory accesses at runtime. Even if the race doesn't cause a crash during the test, the detector spots the *unsynchronized access* and flags it.
I treat every race report as a P0 bug, because in production it could mean data corruption."

#### Indepth
The Race Detector catches races that *actually happen* during execution. It doesn't prove code correctness. If your test suite doesn't trigger the specific timing overlap, the race detector won't complain. Run tests with `-count=10` and `-race` to increase the odds of catching flakes.

---

### 388. How do you analyze goroutine leaks?
"I check the **Goroutine Count** metric.

If `runtime.NumGoroutine()` climbs steadily over days, I have a leak.
I use `pprof` to see *where* they are stuck.
`go tool pprof http://localhost:8080/debug/pprof/goroutine`.
Usually, it’s a goroutine waiting on a channel read where the sender has already exited (or vice versa). The fix is ensuring every goroutine has a `ctx.Done()` exit path."

#### Indepth
Use `runtime.NumGoroutine()` in your health check handler. If the number grows linearly over time (monitor this in Prometheus), you have a leak. Common culprit: a goroutine writing to an unbuffered channel that no one is reading from anymore.

---

### 389. How do you debug network timeouts in Go?
"I categorize the timeout.

1.  **Dial Timeout**: Firewall or DNS. The server is unreachable.
2.  **Handshake Timeout**: TLS issue.
3.  **Response Header Timeout**: Server accepted connection but is overloaded/slow.
I use `httptrace` (ClientTrace) to measure each phase. This tells me if the problem is the Network (Dial slow) or the Server App (Response Header slow)."

#### Indepth
Use `net/http/httptrace`. It gives hooks for `GotConn`, `DNSStart`, `ConnectStart`. You can see if the "Timeout" was 99% DNS resolution time (DNS infrastructure issue) vs 99% Wait time (Server overloaded). This granularity is critical for debugging "flaky networks".

---

### 390. What is core dump analysis?
"A core dump is a snapshot of process memory at crash time.

I enable it with `GOTRACEBACK=crash`.
When the app panics, it writes a `core` file to disk.
I open it with `dlv core ./app core.1234`.
I can inspect variables and stack frames just like a live session. It’s the last resort for 'impossible' bugs that happen once a month in production but never locally."

#### Indepth
**Security Warning**: Core dumps contain the entire memory of the process, including **Secrets** (DB Passwords, Private Keys). Never send core dumps to an external vendor or upload them to public storage without sanitizing or encrypting them.

---

### 391. How do you verify if the GC is the bottleneck?
"I configure the **GC Tracer**.
`GODEBUG=gctrace=1`.

I watch the `PauseNs` and `%CPU` usage.
If the GC is stealing > 25% of my CPU cycles, or if Stop-The-World pauses exceed 10ms frequently, it is a bottleneck.
My solution is always to **Reduce Allocations** (reuse buffers, avoid pointers) rather than tweaking GC knobs."

#### Indepth
If GC CPU usage is > 25%, you are creating too much garbage. Look for **Memory Ballast** (allocating a huge byte array on startup) to trick the GC into running less often (legacy trick), or better, use `GOMEMLIMIT` (Go 1.19) to set a target memory usage, allowing the GC to utilize available RAM fully before marking.

---

### 392. How do you debug a slow SQL query in Go?
"I use a **Slow Query Logger**.

In my DB driver configuration, I set a hook: 'Log any query taking > 100ms'.
When a log appears, I verify the SQL.
I run `EXPLAIN ANALYZE` on the database to check for missing indexes / full table scans.
From the Go side, I ensure I'm not fetching 10,000 rows when I only process 10 (`LIMIT` missing)."

#### Indepth
In Go `database/sql`, check `db.Stats()`. If `OpenConnections == MaxOpenConnections`, your app is waiting for a free connection from the pool. The "slow query" might actually be "fast query, but waited 500ms for a connection". Increase pool size or fix the query holding the connection too long.

---

### 393. How do you handle transient network errors?
"I assume the network is flaky.

I wrap my HTTP client with a **Retry Middleware**.
It catches `503 Service Unavailable` or `Connection Refused`.
It retries with **Exponential Backoff** (1s, 2s, 4s).
**Crucial**: I only retry *Idempotent* operations (GET, PUT). I never retry a POST (Charge Payment) blindly unless I can verify it wasn't processed."

#### Indepth
**Jitter** is mandatory. If a microservice has a hiccups and 1000 clients retry exactly 1.0s later, they will DDOS the service again (Thundering Herd). `time.Sleep(1s + rand.Intn(500ms))` smooths out the retry wave and allows the service to recover.

---

### 394. How do you debug context cancellation issues?
"I inspect the **Cancellation Cause**.

`ctx.Err()` just says 'Canceled'.
In Go 1.20+, `context.Cause(ctx)` returns the *error* that caused the cancellation.
'DeadlineExceeded'? 'Client Disconnected'? 'Explicit Cancel called'?
This tells me if the user gave up (client disconnect) or if my backend was too slow (timeout)."

#### Indepth
Go 1.20 added `context.WithCancelCause(parent)`. This allows you to set *why* the context was canceled (e.g., "Timeout" vs "UserAbort"). Check `context.Cause(ctx)` to log the specific reason, stopping the "Why did this request cancel?" guessing game.

---

### 395. How do you monitor thread exhaustion?
"The Go runtime spawns OS threads (M) for blocking syscalls.

If I make too many CGO calls or blocking File I/O, the runtime spawns thousands of threads.
Eventually, I hit the OS `ulimit` or `runtime: program exceeds 10000-thread limit`.
I monitor the `runtime_thread_count` metric. If it spikes, I know I'm blocking the runtime, and I need to offload that work or rate-limit it."

#### Indepth
Go runtime creates a new OS thread if a goroutine blocks in a System Call (syscall). If you have 10,000 goroutines doing blocking file IO (without async IO poller support), Go *will* spawn 10,000 threads, crashing the app. Use `debug.SetMaxThreads` to prevent this global exhaust.

---

### 396. How do you debug incorrect JSON unmarshaling?
"Common Pitfalls:
1.  **Case Sensitivity**: Struct field `Name`, JSON `name`. Go handles this, but `name` (private field) is ignored.
2.  **Type Mismatch**: JSON `id: "123"` (string), Struct `ID int`.
I use `decoder.DisallowUnknownFields()` to catch typos in the JSON payload.
I verify the error returned by `Unmarshal`—it usually points to the exact byte offset of the mismatch."

#### Indepth
JSON numbers are standardly float64, which loses precision for large IDs. Use `decoder.UseNumber()` to decode numbers as `json.Number` (string) instead of float64. This preserves the exact value of 64-bit integers or high-precision decimals during unmarshalling.

---

### 397. How do you profile lock contention?
"I use the **Mutex Profile**.

`go test -bench=. -mutexprofile=mutex.out`.
`go tool pprof mutex.out`.
It shows me exactly which `sync.Mutex` processes are fighting over.
If contention is high, I fix it by:
1.  **Sharding**: Use 10 locks (one per bucket) instead of 1 global lock.
2.  **RWMutex**: Allow concurrent reads.
3.  **Reducing Critical Section**: Do less work while holding the lock."

#### Indepth
Use `sync.Map` for read-heavy, append-only caches. Standard `sync.Mutex` acts as a bottleneck because even readers contend for the lock. `sync.Map` (or `RWMutex`) allows concurrent readers, drastically reducing contention profiling hot spots in these specific scenarios.

---

### 398. How do you investigate 502 Bad Gateway errors?
"502 means the Load Balancer (Nginx) couldn't talk to my Go app.

Causes:
1.  **Crash**: App died (Check `restart_count`).
2.  **Hang**: App is stuck (GC pause or Deadlock).
3.  **Timeout**: App accepted connection but didn't write headers in time.
4.  **Socket Exhaustion**: No ports left.
I correlate the LB logs ('upstream timed out') with my app metrics to find the root cause."

#### Indepth
502 Bad Gateway often means "Keep-Alive Mismatch". If Go closes the idle connection after 30s, but the Load Balancer (AWS ALB) thinks it's open for 60s, the ALB sends a request to a closed socket. Always set Go's `IdleTimeout` slightly *higher* than the LB's idle timeout.

---

### 399. How do you debug issues that only appear in Docker/K8s?
"I assume Environment Differences.

1.  **OOMKill**: Is the container hitting its memory limit? (Exit Code 137).
2.  **CPU Throttling**: If I set `limits.cpu=100m`, K8s throttles my app heavily, causing latency.
3.  **DNS**: K8s uses CoreDNS. `musl` (Alpine) resolves differently than `glibc`.
I use `kubectl exec -it pod -- sh` to get inside and run `curl`/`nslookup` to verify the environment."

#### Indepth
Use **Ephemeral Containers** (`kubectl debug`). This allows you to attach a generic troubleshooting container (with `curl`, `dig`, `vi`) to a crashed or distroless pod. It shares the process namespace, so you can see the files and processes of the target pod even if it has no shell.

---

### 400. How do you maintain a 'Runbook' for debugging?
"I write **Executable Runbooks**.

Instead of a Word doc saying 'Check the database', I write:
'If Alert X fires:
1.  Run `scripts/debug_db.sh`.
2.  If output > 50 conns, scale up pool.'
This saves mental energy during a 3 AM outage. The goal is to make debugging mechanical and repeatable."

#### Indepth
Treat Runbooks as Code. Store them in Markdown in the repo. Better yet, make them **Executable Runbooks** (Jupyter Notebooks for Ops). If a human has to copy-paste commands during an outage, they will make mistakes. One-click remediation scripts are safer.


## From 32 EventDriven Messaging

# 🟤 **621–640: Event-Driven, Pub/Sub & Messaging**

### 621. How do you publish and consume events using NATS in Go?
"I use the `nats.go` library.
Publish: `nc.Publish("orders.created", []byte("data"))`.
Subscribe: `nc.Subscribe("orders.*", func(m *nats.Msg) { ... })`.
It’s fire-and-forget. Extremely fast (millions of msgs/sec).
If I need persistence (at-least-once delivery), I use **JetStream**."

#### Indepth
**Request-Reply Pattern**. NATS excels here. `nc.Request("help", []byte("me"), 1*time.Second)`. It creates a temporary inbox (subscription), sends the request with the `Reply-To` header set to that inbox, waits for the response, and then cleans up. It makes a distributed system feel like a function call.

---

### 622. How do you use Apache Kafka in Go with `sarama`?
"I create a `sarama.ConsumerGroup`.
I implement the `ConsumerGroupHandler` interface (`Setup`, `Cleanup`, `ConsumeClaim`).
The loop reads from `claim.Messages()`.
I verify to mark offsets *after* processing (`session.MarkMessage(msg, "")`) to ensure I don't process the same message twice if the pod restarts."

#### Indepth
**Rebalance Storms**. If your processing logic takes too long (`> session.Timeout`), the broker thinks the consumer is dead and triggers a Rebalance (stop-the-world). To fix this, either optimize logic or increase `MaxProcessingTime` and use background context cancellation to abort processing when a rebalance starts.

---

### 623. What are the trade-offs between RabbitMQ and Kafka in Go apps?
"**RabbitMQ**: Smart broker. Good for complex routing (topics, fanout) and task queues. Push-based. Harder to replay old messages.
**Kafka**: Dumb broker, smart consumer. Good for high throughput stream processing. Pull-based. Retains history (log), allowing replay.
I choose Rabbit for background jobs, Kafka for data pipelines."

#### Indepth
**Ordering guarantees**. RabbitMQ guarantees order within a queue. Kafka guarantees order ONLY within a *partition*. If you need global order in Kafka, you are limited to 1 partition (no scaling). RabbitMQ is often "easier" for simpler work queues where strict ordering isn't paramount or scale is moderate.

---

### 624. How do you manage message acknowledgements in Go consumers?
"Explicit Ops.
`msg.Ack()` tells the broker 'I am done'.
If I crash before Ack, the broker redelivers.
I only Ack *after* the DB commit is successful.
If DB fails, I `Nack()` (Negative Ack) or let the ack timeout expire."

#### Indepth
Manual Acks are critical. Auto-Ack (default in many libs) is dangerous; if your app crashes after receiving the message but before processing it, the message is lost forever. Always verify `AutoAck: false` in production consumers.

---

### 625. How do you handle message deduplication in Go?
"The broker usually guarantees 'At-Least-Once'.
I make my consumer **idempotent**.
I use a table `processed_messages (message_id PRIMARY KEY)`.
Transaction:
1.  Insert message_id.
2.  If error (duplicate key) -> Return Success (already done).
3.  Else -> Process logic -> Commit."

#### Indepth
**Redis SETNX**. For a lighter check (non-critical deduplication), use Redis `SETNX key "1" EX 86400`. It's faster than a SQL insert but less durable (if Redis crashes without AOF). For financial transactions, always use the SQL Transaction approach (Idempotency Key pattern).

---

### 626. How do you implement a retry queue for failed messages?
"I use distinct topics: `main`, `retry-1m`, `dlq`.
If processing fails:
Publish to `retry-1m`. Ack original.
The `retry-1m` consumer reads, checks timestamp. If < 1m has passed, it sleeps.
After 3 tries, publish to `dlq` (Dead Letter Queue) for manual intervention."

#### Indepth
**Exponential Backoff Headers**. Don't create 10 queues (`retry-1s`, `retry-2s`...). Instead, attach a header `Next-Retry-Time: <timestamp>` and republish to the *same* retry topic. The consumer reads, checks the header, and if it's too early, does a `Nack` (re-queue) with a small sleep. This keeps the topology simple.

---

### 627. How do you batch message processing efficiently in Go?
"I read from the channel into a buffer slice.
`batch := make([]Msg, 0, 100)`.
`timer := time.NewTimer(1 * time.Second)`.
Loop: select case msg: append; if len==100 { flush() } case timer: flush().
`flush()` sends a bulk insert to DB and then Acks all 100 messages. This reduces DB load by 100x."

#### Indepth
**Micro-batching Latency**. The downside is latency. The first message waits up to 1 second. You must tune the parameters (Batch Size vs Flush Interval). `size=100, time=50ms` is often a better sweet spot for near-real-time systems.

---

### 628. How do you use Google Pub/Sub with Go?
"I use `cloud.google.com/go/pubsub`.
It manages the pulling logic for me.
`sub.Receive(ctx, func(ctx, msg) { process(msg); msg.Ack() })`.
It spawns goroutines automatically for concurrency. I configure `ReceiveSettings` (MaxOutstandingMessages) to control memory usage."

#### Indepth
**Flow Control**. If you don't limit `MaxOutstandingMessages`, the library might pull 10,000 messages into RAM if your processing is slow. This causes OOM kills. Always set this to a reasonable number (e.g., `NumCPU * 20`) to match your worker pool capacity.

---

### 629. How do you persist event logs for replay in Go?
"I treat the **Log as the Database**.
In Kafka, I set `retention.ms = -1` (Infinite) for critical topics.
To replay: I start a new Consumer Group from `offset=0`.
My Go app re-processes every historical event. This allows me to rebuild my read-model (e.g., ElasticSearch index) from scratch."

#### Indepth
**Snapshotting**. Replaying 10 years of events takes too long. Periodically (e.g., every 10k events), create a "Snapshot" (current state) and save it to S3. To rebuild, load the latest Snapshot + replay only events *after* that snapshot. This is optimizing the "Recovery Time Objective" (RTO).

---

### 630. How do you ensure exactly-once delivery in Go message systems?
"It's theoretically impossible across boundaries without 2PC.
But Kafka **Transactional Producer** (`initTransactions`) enables it within Kafka.
For side-effects (DB writes), I rely on **Idempotency**.
Combining an Idempotent Consumer + At-Least-Once Delivery = Effectively Exactly-Once Processing."

#### Indepth
Kafka Transactions are heavy and complex (Zombie fencing, Group Coordinators). For 99% of use cases, Idempotency (handling duplicates gracefully) is the correct and more robust engineering solution. "Exactly-Once" is mostly a marketing term for "Exactly-Once *Effect*".

---

### 631. How do you create a lightweight in-memory pub-sub system?
"I use channels and a map of subscribers.
`type PubSub struct { subs map[string][]chan any }`.
`func (ps *PubSub) Pub(topic, val)`: non-blocking send to all channels.
This is great for decoupling components inside a Monolith (e.g., WebSocket handler subscribes to 'chat updates' from the Login handler)."

#### Indepth
**Blocking Sends**. If `Pub` blocks until all subscribers read, one slow subscriber kills the system. Make the subscriber channels buffered, or use `select { case ch <- msg: default: log.Warn("dropped") }` (Non-blocking send). Dropping messages is usually better than deadlocking the Publisher.

---

### 632. How do you handle DLQs (Dead Letter Queues) in Go?
"I write a specific **DLQ Consumer** tool.
It reads the bad messages.
It logs them or shows them in a UI dashboard.
It has a 'Reprocess' button which publishes the message back to the `main` topic.
I investigate *why* it failed (bad JSON? bug?) before clicking Reprocess."

#### Indepth
**Automated Redrive**. Sometimes failures are transient (3rd party API down for 1 hour). You can have a "Redrive Policy" that automatically moves moves messages from DLQ back to Main after X hours. But be careful of infinite loops if the message is "poison" (permanently unprocessable).

---

### 633. How do you create idempotent message consumers in Go?
"I use the **Outbox Pattern** or a **State Table**.
Also, I design operations to be naturally idempotent:
`UPDATE balance SET amount = 100` (Idempotent).
`UPDATE balance SET amount = amount + 10` (Not Idempotent! Needs deduplication logic)."

#### Indepth
**Database Constraints**. Let the DB do the work. `INSERT INTO processed_events (id) VALUES ($1) ON CONFLICT DO NOTHING`. Check `RowsAffected()`. If 0, it was a duplicate. This atomic check-and-insert is the bedrock of idempotency.

---

### 634. How do you enforce ordering of messages?
"I use **Sharding / Partitioning**.
In Kafka, ordering is guaranteed *only within a partition*.
I ensure all events for `Order-123` go to Partition 5 by using `Order-123` as the message Key.
If I use a random key, `Created` might arrive after `Shipped`."

#### Indepth
**Key Skew**. If you partition by `CustomerID`, and one customer (Amazon) does 50% of your traffic, that one partition (and the 1 consumer processing it) will be overwhelmed while others verify idle. You might need "Compound Keys" or specialized sharding strategies for "Celeb" entities.

---

### 635. How do you use channels as message queues?
"Buffered channels are queues.
`q := make(chan Job, 100)`.
But they are **ephemeral**. If the app crashes, data is lost.
I use channels for *work distribution* between goroutines, not for *storage*. For reliable queuing, I always use Redis or RabbitMQ."

#### Indepth
**Persistence vs Speed**. Channels are purely RAM. Redis is RAM + Disk. Postgres is Disk. The trade-off is Durability. If you can afford to lose the job on restart (e.g. sending a "Welcome" email), channels work. If it's a "Payment", you need Disk (Postgres/Kafka).

---

### 636. How do you handle push vs pull consumers?
"**Push** (NATS): Low latency. The broker floods me. I need rate limiting to avoid OOM.
**Pull** (Kafka/SQS): I control the rate. I ask for 10 messages.
In Go, Pull is safer for avoiding backpressure issues.
I implement Pull loops: `batch := sqs.ReceiveMessage(10); process(batch)`."

#### Indepth
**Long Polling**. When pulling from SQS, use `WaitTimeSeconds=20`. This tells SQS: "If the queue is empty, hold my connection open for 20s until a message arrives." This reduces empty API calls (and cost) by 99% compared to a tight loop.

---

### 637. What is event sourcing?
"It means storing *state changes*, not state.
`OrderCreated`, `ItemAdded`, `OrderPaid`.
My Go app reads these events to calculate `CurrentBalance`.
It allows me to answer 'What was the balance last Tuesday?' by replaying events up to that timestamp. The complexity is high, so I use it sparingly."

#### Indepth
**CQRS Relationship**. Event Sourcing almost always requires CQRS. Writing to the Event Store is fast (Append Only). But *Reading* "All users with name like 'Bob'" is impossible. You need a separate process that reads events and updates a standard SQL/NoSQL table (The "Read Model") for queries.

---

### 638. How do you handle schema evolution in event-driven systems?
"I use **Schema Registry** (Avro/Protobuf).
I enable compatibility modes (Forward/Backward).
Producer cannot publish a message that breaks existing Consumers (e.g., removing a required field).
This contract enforcement prevents 'poison pill' messages."

#### Indepth
**Schema Registry**. The Producer shouldn't send the schema with every message (too big). Instead, it registers the schema (ID: 5) and sends `[ID:5][BinaryData]`. The Consumer downloads Schema 5 from the Registry to decode. Confluent Schema Registry is the standard for Kafka.

---

### 639. How do you implement a competing consumers pattern?
"I place multiple consumers (in a Group) on the same queue.
The broker delivers each message to **only one** consumer.
NATS Queue Groups or Kafka Consumer Groups handle this.
It allows horizontal scaling: if processing is slow, I add 5 more Go replicas, and the throughput increases linearly."

#### Indepth
**Partition Limit**. In Kafka, you can't have more consumers than partitions. If you have 10 partitions, the 11th consumer sits idle. This is a hard limit on horizontal scaling. Plan your partition count (e.g. 50 or 100) upfront based on expected future scale.

---

### 640. How do you monitor message lag?
"**Consumer Lag** is the metric.
`Lag = LatestOffset - CurrentOffset`.
If Lag is growing, my consumers are too slow.
I export this to Prometheus using `kafka-exporter`.
I alert if `Lag > 10000` or `Latency > 1 minute`."

#### Indepth
**Burrow**. Tools like LinkedIn's Burrow monitor lag *without* being part of the consumer group. They inspect offsets independently. This prevents "Observer Effect" where monitoring slows down the consumer. Lag is the single best metric for "Is my system healthy?".


## From 33 DevOps Infrastructure

# 🟢 **641–660: Go for DevOps & Infrastructure**

### 641. How do you create a custom Kubernetes operator in Go?
"I use the **Operator SDK** or **Kubebuilder**.
It scaffolds the project.
I define a CRD struct: `type MyDB struct { Spec MySpec }`.
I implement the `Reconcile` loop:
`func (r *Reconciler) Reconcile(ctx, req) (Result, error)`.
This loop reads desired state (CRD) vs actual state (Pods), and creates/updates K8s objects to match."

#### Indepth
**Level 1 vs Level 5**. Most operators start as "Basic Install" (Level 1). The goal is "Autopilot" (Level 5): handling upgrades, backups, failure recovery, and vertical scaling without human intervention. The SDK provides `Operator Lifecycle Manager (OLM)` capabilities to package these advanced features.

---

### 642. How do you write a Helm plugin in Go?
"A Helm plugin is just a binary executed by Helm.
I write a Go CLI (using Cobra).
I add a `plugin.yaml` describing the hooks (`helm myplugin install`).
My Go code accesses Helm's env vars (`HELM_BIN`, `HELM_NAMESPACE`) to interact with the cluster or decrypt secrets before the chart installs."

#### Indepth
**Downloader Plugins**. Helm supports custom protocol handlers (`s3://`, `git://`). You can write a Go plugin that registers itself as a downloader, allowing `helm install s3://my-bucket/chart.tgz` to work seamlessly. This is great for private chart repositories.

---

### 643. How do you use Go for infrastructure automation?
"I prefer Go over Bash for complex logic.
I use `os/exec` or native Cloud SDKs.
Go's strict error handling prevents the 'script continued after error' bugs common in Bash.
I compile it to a single binary `ops-tool` and distribute it to the team—no `pip install` or Ruby gems required."

#### Indepth
`Mage` is a popular Make-alternative written in Go. Instead of a `Makefile`, you write `magefile.go`. This gives you the full power of Go (imports, loops, type safety) for your build scripts, which is much more maintainable than complex Bash spaghetti.

---

### 644. How do you write a CLI for managing AWS/GCP resources?
"I use the official SDKs: `aws-sdk-go-v2` or `google-cloud-go`.
I organize commands by resource: `mycli ec2 list`.
I use `context` for timeouts—crucial for cloud APIs.
I use interfaces for the cloud client so I can mock AWS in tests, allowing me to verify my 'delete unused instances' logic safely."

#### Indepth
**Pagination**. Most Cloud APIs paginate results. A common bug is processing only the first page (e.g., first 50 buckets). The Go SDKs often proide "Paginators" (`NewListBucketsPaginator`). Always use them to interact with the collection, even if you think you'll only have 10 items.

---

### 645. How do you use Go to write Terraform providers?
"I use the **Terraform Plugin Framework**.
I define `Resources` (Create, Read, Update, Delete).
`func (r *resource) Create(ctx, req, resp)`.
My Go code maps the Terraform HCL state to API calls.
This allows me to manage my company's internal platform (which has an API but no official TF provider) via Terraform."

#### Indepth
**State Drift**. Your provider must implement `Read` correctly to detect drift. If someone manually changes a value in the UI, `terraform plan` should see the difference. The `Read` function essentially maps the *External API Response* back to the *Internal Terraform Schema*.

---

### 646. How do you build a dynamic inventory script in Go for Ansible?
"Ansible expects a JSON output.
I write a Go CLI that queries my Cloud Source (AWS/Consul).
I marshal the result to the specific JSON format: `{ "_meta": { "hostvars": ... } }`.
I compile it and point `ansible -i my-go-inventory` to the binary."

#### Indepth
**Dynamic Groups**. Your Go tool can contain logic: "If name starts with 'db-', add to 'databases' group". This centralizes grouping logic in code rather than in static INI files, making your Ansible playbooks automatically adapt to new instances without manual edits.

---

### 647. How do you parse and generate YAML in Go?
"I use `gopkg.in/yaml.v3`.
It supports comments and preserving order (unlike JSON).
`yaml.Unmarshal(data, &configStruct)`.
For K8s manifests, I use `sigs.k8s.io/yaml` which handles the specific JSON/YAML duality of Kubernetes.
I check for `KnownFields: true` to catch typos in config files."

#### Indepth
Use `yaml:",omitempty"` sparingly in Infrastructure-as-Code. If a user sets `replicas: 0`, and you use `omitempty`, the YAML generator omits the field, and K8s might default it to `1`. Explicitly serializing zero-values is often safer for declarative configs.

---

### 648. How do you interact with Docker API in Go?
"I use the official `docker/docker/client`.
`cli, _ := client.NewClientWithOpts(client.FromEnv)`.
`cli.ContainerList(ctx, options)`.
I can build custom sidecars that listen to Docker events (container died) and trigger alerts or cleanup tasks."

#### Indepth
**Stream Handling**. `ContainerAttach` or `ImagePull` returns a stream. You must handle this stream correctly (it often multiplexes Stdout/Stderr/Stdin). Failing to read the stream can cause the Docker daemon to block (backpressure) and hang the operation. use `stdcopy.StdCopy`.

---

### 649. How do you manage Kubernetes CRDs in Go?
"I define the struct with JSON tags.
`type MyResource struct { Spec MySpec ... }`.
I use `controller-gen` to generate the YAML CRD manifest from the Go struct tags.
In my code, I use the dynamic client (`client.Client`) to List/Watch these custom resources just like native Pods."

#### Indepth
**Validation Schemas**. CRDs allow you to define OpenAPI v3 schemas. This enforces types at the API server level (e.g., "replicas must be > 0"). Always define these schemas rigourously so your Go controller doesn't crash trying to parse invalid JSON from a user's CR.

---

### 650. How do you write Go code to scale deployments in K8s?
"I use the client-go `Scale` subresource.
`clientset.AppsV1().Deployments(ns).UpdateScale(...)`.
Or I calculate the desired replica count based on custom metrics (e.g., RabbitMQ queue depth) and patch the Deployment. This is how KEDA works."

#### Indepth
**Metrics Server**. To use the Horizontal Pod Autoscaler (HPA) with your custom metric, you must implement the `External Metrics API`. You write a Go adapter that translates "Queue Depth" into a format the K8s Metrics Server understands, allowing standard HPA resources to scale your app.

---

### 651. How do you tail logs from containers using Go?
"I use `cli.ContainerLogs(ctx, containerID, options)`.
It returns a multiplexed stream (stdout + stderr headers).
I must use `stdcopy.StdCopy(dstOut, dstErr, src)` to demultiplex the binary stream into readable text lines."

#### Indepth
**Follow Mode**. If you use `Follow: true`, the connection stays open. You must handle network interruptions. A robust log tailer loop monitors the error channel and *re-connects* if the stream breaks, possibly using the `Since` timestamp to resume from the last received log line.

---

### 652. How do you manage service discovery in Go apps?
"I avoid hardcoding IPs.
In K8s, I use DNS: `http://my-service.default.svc`.
For external service discovery (Consul), I use the Consul API client to watch for healthy nodes and update my internal load balancer dynamically."

#### Indepth
**Client-Side Load Balancing**. Instead of a VIP (Virtual IP) which is a bottleneck, the Go client knows all 50 backend IPs. It picks one directly. gRPC has this built-in. It reduces latency (1 less hop) but requires the client to be "smart" and aware of the topology.

---

### 653. How do you build a Kubernetes admission controller in Go?
"It’s a webhook (HTTP Server).
K8s POSTs the Pod definition to me *before* creating it.
I decode the `AdmissionReview`.
I inspect the Pod ('Does it have root privileges?').
I return `Allowed: false` or a JSON Patch to mutate it (e.g., inject a sidecar)."

#### Indepth
**Cert Management**. Webhooks require HTTPS. Managing these internal certs is painful. Use `cert-manager` with a `CAInjector` to automatically inject the CA bundle into your `ValidatingWebhookConfiguration`. This prevents the "x509: certificate signed by unknown authority" error during cluster upgrades.

---

### 654. How do you build a metrics exporter for Prometheus in Go?
"I write a 'Collector'.
`func (c *MyCollector) Collect(ch chan<- prometheus.Metric)`.
On every scrape, I query my target (e.g., a hardware device), convert values to Metrics, and send to channel.
Then I run `http.Handle("/metrics", promhttp.Handler())`."

#### Indepth
**Describe vs Collect**. Implementing `Describe` is optional but good practice. It sends metric metadata (Help string, Type) to Prometheus. `Collect` is the hot path. Ensure `Collect` doesn't block for too long, or the scraper will timeout. If data gathering is slow, do it asynchronously and return the last cached value in `Collect`.

---

### 655. How do you set up health checks for a Go microservice?
"I follow the K8s pattern.
`/healthz` (Liveness): Returns 200 if loop is running.
`/ready` (Readiness): Returns 200 if DB is reachable.
I use a library like `hellofresh/health-go` to check DB/Redis connectivity and aggregate the result."

#### Indepth
**Cascading Failures**. Be careful with Readiness checks. If your DB is slow, and all 100 pods return "Not Ready", K8s kills traffic to ALL of them. Now the DB recovers, but your pods are effectively offline or flapping. Sometimes returning "Healthy" even if the DB is down (Degraded Mode) is safer to keep the UI accessible.

---

### 656. How do you build a custom load balancer in Go?
"I use `httputil.ReverseProxy` with a custom `Director`.
I maintain a list of backends `[]*url.URL`.
The Director function picks the next backend (Round Robin) and updates `req.URL.Host`.
I handle errors: if the backend is down, I retry on another."

#### Indepth
**Active Health Checks**. Passive check (wait for error) is slow. A real LB runs a background loop `HEAD /health` on backends. If one fails, it removes it from the rotation *before* a user request hits it. The `Director` needs read-access to this dynamic "Healthy Backends" list.

---

### 657. How do you implement graceful shutdown with Kubernetes SIGTERM?
"I listen for signals.
`c := make(chan os.Signal, 1); signal.Notify(c, syscall.SIGTERM)`.
Block: `<-c`.
Start shutdown:
1.  Set `readiness` to 503.
2.  `server.Shutdown(ctx)` (waits for active requests).
3.  Wait for background jobs.
This ensures K8s stops sending traffic and existing requests complete before the pod dies."

#### Indepth
**PreStop Hook**. K8s updates endpoints *async*. Even after SIGTERM, network rules might route traffic to you for a few seconds. It's common to add a `preStop` hook: `sleep 5`. This ensures the pod stays alive long enough for K8s to propagate the "Remove me from Endpoints" command to all kube-proxies.

---

### 658. How do you use Go with Envoy/Consul service mesh?
"I usually don't interact with them directly; they are sidecars.
However, I can use the **xDS protocol** (Go control plane) to configure Envoy dynamically.
This allows me to tell Envoy 'Route 50% traffic to v2' programmatically."

#### Indepth
**Sidecar vs Library**. With gRPC, you can skip Envoy and stick the mesh logic *inside* the Go binary (proxyless service mesh). This reduces latency (no localhost hop) and complexity (no sidecar container), but requires all apps to be written in languages with rich SDKs (Go/Java/C++).

---

### 659. How do you configure Go apps for 12-factor principles?
"**Env Vars** via `os.Getenv` or `kelseyhightower/envconfig`.
**Logs** to Stdout (JSON).
**Backing Services** (DB) attached via connection string URL.
**Stateless**: No local files.
**Port Binding**: Listen on `$PORT`.
I stick to these rules so my app runs anywhere (Docker, Heroku, K8s) without changes."

#### Indepth
**Config Separation**. strict separation means the same Docker image is used for Dev, Staging, and Prod. The *only* difference is the Environment Variables. If you are baking `config.prod.json` into the image, you are violating 12-Factor and making rollbacks harder.

---

### 660. How do you use Go for cloud automation scripts?
"I replace Bash scripts with Go.
Benefit: **Parallelism**.
'Restart 100 VMs':
Bash: Loop 1..100 (slow).
Go: `for vm := range vms { go restart(vm) }` (fast).
I compile it to a static binary for CI/CD pipelines."

#### Indepth
**Cross Combination**. `GOOS=linux GOARCH=amd64 go build`. You can build the maintenance script on a Mac and ship it to a Linux server. This portability is why Go is the language of cloud infrastructure (Docker, K8s, Terraform, Vault are all Go).


## From 34 Caching Storage

# 🟣 **661–680: Caching & Storage Systems**

### 661. How do you cache database query results in Go?
"I use the **Cache-Aside** pattern.
1.  Check Redis: `val, err := rdb.Get(ctx, key)`.
2.  If found (Hit), unmarshal and return.
3.  If missing (Miss), Query DB.
4.  Write to Redis (SetEX with TTL).
5.  Return data.
I wrap this logic in a generic function `GetCached[T]` to avoid repetition."

#### Indepth
**Thundering Herd**. If 1000 users request the same key simultaneously (Cache Miss), 1000 queries hit the DB. Use `singleflight` (Chapter 29) *before* the DB query. This ensures only 1 query runs, and the result is shared with all 1000 waiters. This is mandatory for hot keys.

---

### 662. How do you use Redis with Go for distributed caching?
"I use the `go-redis/v9` library.
It supports connection pooling and cluster mode automatically.
`rdb := redis.NewClient(...)`.
I prefer storing data as **Protobuf** or **MsgPack** (binary) instead of JSON to save space and CPU.
I always use a TTL to prevent the cache from filling up with stale data forever."

#### Indepth
**Pipelining**. If you need to set 50 keys, don't do 50 Round Trips. Use `rdb.Pipeline()`. It batches commands into a single TCP packet. `pipe.Set(...)`; `pipe.Exec(ctx)`. This reduces network latency by 50x for bulk operations.

---

### 663. How do you implement LRU cache in Go?
"I use `hashicorp/golang-lru` or implement it with a Map + Doubly Linked List.
Map stores `Key -> NodePtr`.
List stores nodes in order of access.
**Get**: Move node to front of list.
**Set**: If full, remove last node (tail), delete from map, add new node to front.
This guarantees O(1) operations."

#### Indepth
**Generics Implementation**. `hashicorp/golang-lru` uses `interface{}` (heap allocation). In Go 1.18+, you can write a type-safe generic LRU `[K comparable, V any]` to avoid these allocations. Combine this with a `sync.Pool` for the nodes to achieve zero-alloc LRU cache.

---

### 664. How do you ensure cache invalidation on data update?
"It's the hardest problem in CS.
Strategies:
1.  **TTL**: Accept staleness for N minutes.
2.  **Write-Through**: Update DB and deletions Cache in same transaction.
3.  **CDC (Change Data Capture)**: Listen to DB binlog (Debezium) -> Kafka -> Go Worker -> `redis.Del(key)`. This is the most robust web-scale solution."

#### Indepth
**Broadcasting**. If using Local Cache (in-memory), clearing Redis isn't enough. You must broadcast a "Invalidate Key X" message to *all* app instances (via Redis Pub/Sub). Each instance hears the event and deletes Key X from its local RAM.

---

### 665. How do you handle stale reads in Go apps with caching?
"I use the **Probabilistic Early Recomputation** (X-Fetch) pattern.
Store `beta` value.
If `TTL - now < beta * log(rand)`, I recompute the value in the background *while returning the stale value*.
This prevents the 'Thundering Herd' where everyone misses the cache at the exact same second."

#### Indepth
**Grace Period**. Another trick is "Soft vs Hard TTL". Item expires at 5 mins (Soft), but is kept for 10 mins (Hard). If accessed between 5-10m, return the Stale item immediately, but kick off a background refresh. This is cleaner than probabilistic math.

---

### 666. How do you implement a write-through cache in Go?
"My application wraps the Cache and DB.
`func SaveUser(u User) { tx := db.Begin(); repo.Save(tx, u); cache.Set(u.ID, u); tx.Commit() }`.
The downside is latency (writing to two places).
If Cache write fails, I often log a warning and let the cache be stale (eventually inconsistent) rather than failing the user request."

#### Indepth
**Write-Behind**. Write to Cache *only*, and let a background worker flush to DB. This is extremely fast (RAM speed) but risky (data loss if server crashes). Only use for non-critical data (like 'Like Counts' or 'User Presence').

---

### 667. How do you handle concurrency in in-memory caches?
"If using a simple map, I need `sync.RWMutex`.
`mu.RLock()` for getting. `mu.Lock()` for setting.
For high contention, I use **Sharding**.
`dgraph-io/ristretto` is a high-performance Go cache. It uses tiny LFUs and sharded locks to handle millions of Ops/sec without contention bottlenecks."

#### Indepth
**BigCache**. If you have GBs of data and GC pauses are killing you, use `bigcache`. It bypasses the GC by allocating a massive `[]byte` arena and storing entries as serialized bytes (pointers are hidden from GC). It effectively manually manages memory in Go.

---

### 668. How do you use bloom filters in Go?
"I use `bits-and-blooms/bloom`.
It’s a probabilistic set.
If `Test` says No, the key is **definitely not** in the set.
If `Test` says Yes, it *might* be (False Positive).
I use it before querying DB/Disk: 'Do we have user X?'. If No, save the disk IO. If Yes, check disk to confirm."

#### Indepth
**Sizing**. A Bloom Filter needs size (N) and error rate (P) upfront. If you underestimate N, the filter fills up and the False Positive rate spikes to 100%, rendering it useless. Always overprovision or use Scalable Bloom Filters that grow automatically.

---

### 669. How do you build a TTL-based memory cache?
"I store items with `Expiration int64`.
I run a background goroutine (Cleaner).
`ticker := time.NewTicker(1 * time.Minute)`.
`for range ticker.C { mu.Lock(); deleteExpired(); mu.Unlock() }`.
This 'Stop the World' cleanup is bad for large caches. Better approach: randomized sampling (Redis style) or a heap-based priority queue."

#### Indepth
**Active Expiration**. In user-land Go, rely on `Get()` to check expiry. `val, ok := map[key]; if val.Expired() { delete; return nil }`. The background cleaner is just a safety net for keys that are *never* accessed again to prevent memory leaks.

---

### 670. How do you use memcached in Go?
"I use `bradfitz/gomemcache`.
It’s simpler than Redis (Key-Value only, no structures).
It uses Consistent Hashing on the client side to distribute keys across multiple servers.
It’s great for raw HTML fragment caching (`GetMulti`) where Redis features aren't needed."

#### Indepth
**Slab Allocation**. Memcached never fragments memory because it uses Slabs (fixed size chunks). Redis uses `malloc`, which can lead to fragmentation over years. For pure, dumb, high-throughput caching of uniform objects, Memcached is technically superior in memory efficiency.

---

### 671. How do you store large binary blobs in Go?
"Not in the DB!
I store metadata in Postgres (`file_url`, `size`).
I stream the blob to **S3 / MinIO**.
Client Upload -> Go (generates Presigned URL) -> Client uploads to S3 directly.
This saves my Go server bandwidth and CPU."

#### Indepth
**Range Requests**. Using `http.ServeContent` with an `io.ReadSeeker` allows clients to request bytes 0-100 (Range Header). S3 supports this natively. If you proxy S3 through Go, ensure you pass the `Range` header through so video players can seek efficiently.

---

### 672. How do you build an append-only log file storage in Go?
"I open file with `O_APPEND`.
Write: `f.Write(entry); f.Sync()`.
To read efficiently, I maintain an in-memory index: `OffsetMap[ID] -> FilePosition`.
`f.Seek(pos, 0); f.Read(...)`.
This is exactly how Kafka and Bitcask storage engines work."

#### Indepth
**Log Rotation**. You can't write to one file forever. Implement rotation: when file reaches 1GB, rename to `data.1`, open new `data.active`. The reading index must now track `FileID + Offset`. Old files can be compacted (garbage collected) by removing deleted keys.

---

### 673. How do you use BoltDB or BadgerDB in Go?
"They are **Embedded KV Stores** (pure Go).
No external server process. Ideal for single-node apps.
**BoltDB**: B+Tree, read-heavy.
**BadgerDB**: LSM Tree, write-heavy.
I use them when deployment simplicity is priority (just one binary, no docker-compose for Redis)."

#### Indepth
**Read/Write Amplification**. BoltDB (B+Tree) does 1 read per page, but random writes are slow (updating pages). Badger (LSM) writes strictly sequentially (fast), but reads might need to check multiple SSTables. Choose Bolt for Read-Heavy (Config, CMS), Badger for Write-Heavy (Logs, Metrics).

---

### 674. How do you structure a file-based key-value store in Go?
"Naive: JSON file. Read all, modify, write all. Slow.
Better: **Log Structured Merge (LSM)**.
Writes go to MemTable (RAM). When full, flush to SSTable (Disk).
Compaction process merges old SSTables.
This implementation is complex; I usually defer to `LevelDB` or `Badger` unless learning."

#### Indepth
**Wal (Write Ahead Log)**. Before adding to MemTable, write to a pure append-only WAL file to survive crashes. On startup, replay the WAL to reconstruct the MemTable. This guarantees Durability (D in ACID).

---

### 675. How do you handle distributed caching with Go?
"I use **Groupcache** (by Google).
It’s a library, not a server.
Peers talk to each other.
If Peer A needs Key X, and Peer B owns Key X, A asks B.
It eliminates the 'Hot Key' problem because the hot key is stored in memory on the owner node, and others request it. It has no eviction, designed for static content."

#### Indepth
Groupcache does NOT support value updates (Immutable). It fails if you need to `Update("user:1", new_data)`. It works best for content that never changes (like a Blob based on Content-Hash) or changes very rarely. It powers Google's "dl.google.com" downloads.

---

### 676. How do you monitor cache hit/miss ratios in Go?
"I wrap my cache client.
`func Get(k) { increment('cache_total'); val := cache.Get(k); if val { increment('cache_hit') } else { increment('cache_miss') } }`.
I graph `hit / total` in Grafana.
If Hit Ratio drops < 80%, I investigate (TTL too low? Key eviction? Bad access pattern?)."

#### Indepth
**Metrics Cardinality**. Do NOT tag metrics with the *Key* (`cache_miss{key="user:123"}`). This creates infinite metrics and kills Prometheus. Tag by "Type" (`cache_miss{type="users"}`). You want to know if the "User Cache" is healthy, not specific keys.

---

### 677. How do you use consistent hashing for distributed caching?
"I map keys to a Ring of servers (0-360 degrees).
`hash(key)` finds the point on the ring.
I walk clockwise to find the first Server.
This ensures that adding a new Cache Node only invalidates 1/N keys, not ALL keys.
I use `stathat/consistent` or `serialx/hashring`."

#### Indepth
**Virtual Nodes**. To avoid lopsided distribution (one server getting 80% keys by bad luck), mapping 1 server to 100 "Virtual Nodes" on the ring. This statistically smoothes out the distribution so keys are spread evenly even with a small number of physical servers.

---

### 678. How do you build a cache warming strategy in Go?
"On startup, my app is cold (slow).
I implement a **Warmer**.
It reads the 'Top 1000 Accessed Keys' (from yesterday's logs).
It proactively fetches them from DB and populates Redis *before* the app marks itself as `/ready`.
This prevents the deployment latency spike."

#### Indepth
**Startup Probes**. K8s `startupProbe` is perfect here. It can fail for 60s while the cache warms up. Only once the warmer finishes does the app switch to `readinessProbe`, allowing traffic. This ensures users never hit a cold cache.

---

### 679. How do you use S3-compatible storage APIs in Go?
"I use the **MinIO SDK** or AWS SDK.
MinIO SDK is cleaner.
`minioClient.PutObject(ctx, bucket, name, reader, size, opts)`.
I ensure I handle **Context Cancellation**: if the user closes the connection, the upload to S3 should abort to save bandwidth."

#### Indepth
**Multipart Uploads**. For files > 100MB, splits them into 5MB chunks and upload in parallel goroutines. `minio` SDK handles this automatically. It also allows resuming failed uploads. Never upload a 1GB file in a single PUT request; a network blip at 99% forces a full restart.

---

### 680. How do you implement local persistent disk caching?
"I use a directory structure: `cache/ab/cd/abcdef...` (sharded folders).
Check file existence. Check mod time.
If expired or missing, fetch and write file.
I verify to use atomic file writes (`ioutil.TempFile` + `os.Rename`) so I never serve a half-written file to a concurrent reader."

#### Indepth
**LRU Deletion**. You can't keep writing files forever. Run a background garbage collector that checks disk usage. If Usage > 80%, delete the oldest files (based on mtime/atime) until Usage < 70%. `syscall.Statfs` gives you disk free space info.


## From 48 API Microservices Part2

# 🌐 **941–960: REST APIs & gRPC Design (Part 2)**

### 941. How do you handle CORS in a Go API?
"I prefer `rs/cors` middleware.
`c := cors.New(cors.Options{ AllowedOrigins: []string{"*"} })`.
`handler = c.Handler(mux)`.
Manual way: `w.Header().Set("Access-Control-Allow-Origin", "*")`...
But handling Preflight (`OPTIONS`) correctly is tedious manually."

#### Indepth
**Security Headers**. CORS is just one part. Secure implementations also add:
*   `Strict-Transport-Security` (HSTS)
*   `Content-Security-Policy` (CSP)
*   `X-Content-Type-Options: nosniff`
Use `unrolled/secure` middleware to set these automatically. Leaving them out makes your API vulnerable to XSS and MIME sniffing attacks.

---

### 942. How do you paginate API responses?
"I use **Cursor-based Pagination** for feeds.
`?after=cursor_xyz`.
The cursor is `base64(timestamp)`.
Query: `SELECT * FROM posts WHERE created_at < decoded_cursor LIMIT 10`.
This avoids the 'shifting window' problem of Offset pagination when new items are added."

#### Indepth
**Opaque Cursors**. Never expose the raw implementation details in the cursor (e.g., `?after=2023-01-01`). Users will try to hack it. Always encrypt or sign the cursor string so it's opaque (`?after=E5A3C...`). This allows you to change the underlying implementation (Switching from Timestamp to ID) without breaking API clients.

---

### 943. How do you implement rate-limiting on APIs?
"I use `token bucket` algorithm per IP Key.
Store in Redis: `key=rate:ip:127.0.0.1`.
If I need strict enforcement: Middleware checks Redis before handling.
`tollbooth` library is good for single-instance limiting; Redis script is needed for distributed limiting."

#### Indepth
**GCRA**. The "Leaky Bucket" is simple but "Generic Cell Rate Algorithm" (GCRA) is smoother. It calculates the theoretical "Arrival Time" of the next request. If `Arrival < Now`, request is allowed. Redis `rejson` or Lua scripts can implement GCRA atomically, ensuring strict rate limits even with 50 concurrent Go pods.

---

### 944. How do you handle multipart/form-data in Go?
"`r.ParseMultipartForm(32 << 20)`. (32MB max memory).
`file, handler, _ := r.FormFile("upload")`.
`dst, _ := os.Create(handler.Filename)`.
`io.Copy(dst, file)`.
I always limit the max upload size (`http.MaxBytesReader`) to prevent disk filling attacks."

#### Indepth
**Direct Uploads**. For high-scale apps, *never* handle uploads in the Go API. It blocks a goroutine and consumes bandwidth. Use "Presigned URLs" (S3). Client requests upload URL -> Go returns S3 URL -> Client uploads directly to S3. Go just receives a webhook/callback when the file is ready. this keeps the API stateless and fast.

---

### 945. How do you expose metrics from a Go API?
"Prometheus Middleware.
It wraps the `ServeHTTP`.
It records `http_request_duration_seconds` (histogram) and `http_requests_total` (counter).
Labels: `path`, `method`, `status`.
Note: I verify to sanitize `path` (use `/users/:id` instead of `/users/123`) to avoid high cardinality metrics explosion."

#### Indepth
**Exemplars**. Modern Prometheus (with OpenMetrics) supports "Exemplars". It attaches a `TraceID` to the metric bucket. If you see a latency spike in the histogram, you can click the bucket and jump directly to the specific Trace in Jaeger that caused that delay. This bridges the gap between Metrics and Tracing.

---

### 946. How do you mock gRPC services in tests?
"Protoc generates an interface `UserServiceClient`.
`mockgen` generates `MockUserServiceClient`.
`mockClient.EXPECT().GetUser(gomock.Any(), req).Return(resp, nil)`.
I inject this mock into my API handler to test the logic without spinning up a real gRPC server."

#### Indepth
**Buf**. Google's `protoc` is hard to configure. Use **Buf**. It simplifies generation (`buf generate`), detects breaking changes (`buf breaking`), and manages dependencies (`buf.yaml`). It's significantly faster and cleaner than managing complex `protoc -I ...` flags manually.

---

### 947. How do you set up gRPC with reflection?
"`import "google.golang.org/grpc/reflection"`.
`s := grpc.NewServer()`.
`reflection.Register(s)`.
This allows tools like **grpcurl** or **Postman** to inspect the server (ls services, describe methods) and call RPCs without having the local `.proto` file."

#### Indepth
**Security Risk**. Reflection is great for Dev, but dangerous for Prod (Information Disclosure). An attacker can enumerate your entire API surface. Ensure Reflection is enabled *only* in Staging/Dev environments, or protected behind an Admin-only port/interceptor.

---

### 948. How do you stream data over gRPC?
"Server Streaming: `func (s *server) List(req, stream) error`.
`for item := range items { stream.Send(item) }`.
Client sees: `for { msg, err := stream.Recv(); if err == io.EOF { break } }`.
This keeps the connection open. Ideal for real-time tickers."

#### Indepth
**Flow Control**. gRPC handles flow control strictly. If the Client reads slowly, the Server's `Send()` will eventually block (once the window is full). Go's `Send()` doesn't just push bytes; it waits for window availability. This prevents a fast Server from overwhelming a slow Client with memory pressure.

---

### 949. How do you version gRPC APIs?
"Protobuf package versioning.
`package myapp.v1`.
If I make breaking changes, I create `package myapp.v2`.
The Go code will have `pbv1` and `pbv2` imports.
Typically, we maintain backward compatibility (add optional fields only) to avoid breaking v1."

#### Indepth
**Shadowing**. When migrating `v1` -> `v2`, run **Shadow Traffic**. The Gateway sends the request to `v1` (returns response to user) AND asynchronously to `v2` (discards response). Compare the results/errors. Only when `v2` matches `v1` 100% do you switch the user traffic. This guarantees a bug-free migration.

---

### 950. How do you enforce contracts with protobuf validators?
"**Protoc-Gen-Validate (PGV)**.
`string id = 1 [(validate.rules).string.uuid = true];`
The generated Go code has `.Validate()` method.
I wrap my server with a Unary Interceptor that calls `req.(Validator).Validate()`.
If it fails, I return `InvalidArgument` error automatically. No manual if-checks used."

#### Indepth
**CEL**. The newer standard is **Common Expression Language (CEL)**. Google is moving from custom `validate` rules to standard `cel` annotations. `option (buf.validate.field).cel = { id: "email", expression: "this.size() < 255" }`. CEL is powerful, safe, and portable across languages.

---

### 951. How do you convert REST to gRPC clients?
"I don't convert clients; I wrap the gRPC Service with `grpc-gateway` to expose REST.
If I *must* call gRPC from a legacy REST app:
I write an `Adapter`.
`Adapter.GetUser` calls `grpcClient.GetUser`.
The legacy app just interacts with the Adapter interface."

#### Indepth
**ConnectRPC**. An alternative to `grpc-gateway` is **ConnectRPC** (by Buf). It supports gRPC *and* standard HTTP/JSON on the same port without a proxy. You can `curl` a Connect service with standard JSON body. It removes the need for the complex sidecar/reverse-proxy architecture of `grpc-gateway`.

---

### 952. How do you monitor gRPC health checks?
"Standard Health Protocol.
`grpc_health_v1`.
`healthServer := health.NewServer()`.
`healthServer.SetServingStatus("myservice", SERVING)`.
K8s probes call this via `grpc_health_probe` binary.
It confirms the application logic is up, not just the TCP port."

#### Indepth
**Liveness vs Readiness**.
*   **Liveness**: "Am I crashed?" (Restart me).
*   **Readiness**: "Can I take traffic?" (Load Balancer check).
If DB is down, `Readiness` should fail (stop sending traffic), but `Liveness` should pass (don't restart, restarting won't fix the DB). Getting this wrong causes "Cascading Restart Loops".

---

### 953. How do you build a gRPC gateway in Go?
"**grpc-gateway**.
It reads the `google.api.http` annotations in my Proto.
It generates a reverse proxy.
REST Request -> JSON -> Proto -> gRPC Server.
It’s transparent and allows me to support standard HTTP clients (Frontend/Webhooks) with zero extra coding."

#### Indepth
**OpenAPI Gen**. `grpc-gateway` can also generate `swagger.json` (OpenAPI spec) from the Protobuf annotations. This means your gRPC Source of Truth automatically generates your REST API Documentation UI (Swagger UI), keeping documentation perfectly in sync with the code.

---

### 954. How do you throttle gRPC traffic in Go?
"TapHandle or Interceptors.
In Interceptor:
`if !limiter.Allow() { return status.Error(ResourceExhausted, "rate limit") }`.
gRPC has built-in support for Load Balancing policies, but Rate Limiting is usually an application concern."

#### Indepth
**Concurrency Limits**. Instead of generic RPS, use **Max Concurrent Requests**. Go's `net/http` or gRPC `TapHandle` can track `atomic.Add(&active, 1)`. If `active > 100`, reject. This protects against "Slowloris" attacks or degraded performance where requests pile up and consume all RAM.

---

### 955. How do you test gRPC streams in Go?
"I use `bufconn`.
It creates an in-memory listener.
`lis = bufconn.Listen(1024)`.
Dial it: `grpc.DialContext(ctx, "bufnet", WithContextDialer(bufDialer))`.
This allows me to test streaming logic end-to-end without opening a real TCP port."

#### Indepth
**Race Detector**. Accessing the same stream `Send/Recv` from multiple goroutines is a common panic source in gRPC. Always run stream tests with `-race`. A common pattern is `go func() { stream.Recv() }` and `stream.Send()` in main loop. If not coordinated, the stream state can get corrupted.

---

### 956. How do you handle huge payloads in gRPC (>4MB)?
"Default limit is 4MB.
Either increase `MaxRecvMsgSize`.
Or better: **Chunking**.
Stream the message in 64KB chunks.
Or send a pre-signed URL to S3, upload there, and just pass the URL in the gRPC message."

#### Indepth
**Compression**. Before chunking, enable compression. `grpc.UseCompressor(gzip.Name)`. Text/JSON payloads compress by 90%. However, do NOT compress encrypted data or random binary data (it wastes CPU). Use compression selectively based on the Content-Type.

---

### 957. How do you implement retry logic in gRPC clients?
"I verify to use the `grpc-retry` interceptor.
`grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...))`.
I configure it to retry on `Unavailable` (network blip) but not `InvalidArgument` (bug).
Exponential backoff is built-in."

#### Indepth
**Hedging**. For critical low-latency calls, use **Hedging**. Config allows: "If no response in 10ms, send a 2nd request to a different backend instance. Return whichever replies first". This tackles "Tail Latency" caused by a single slow node/GC pause, drastically improving p99 performance.

---

### 958. How do you secure service-to-service communication?
"**mTLS**.
Every service has a certificate signed by the internal CA (e.g., Vault or Cert-Manager).
The server validates the client's cert.
The client validates the server's cert.
This guarantees identity and encryption."

#### Indepth
**SPIFFE/SPIRE**. Managing certs manually is hell. Use **SPIRE**. It automatically rotates short-lived certificates (TTL 5 mins) for every workload in K8s. It identifies workloads by "Attestation" (Docker SHA), not just IP. Go has `spiffe-go` libraries to handle the mTLS handshake transparently.

---

### 959. How do you trace distributed gRPC calls?
"OpenTelemetry.
I add the `otelgrpc` interceptor to both Client and Server.
It injects the Trace Context into metadata headers.
Spans across services are linked automatically in Jaeger."

#### Indepth
**Baggage**. OpenTelemetry Context can carry "Baggage". Key-Value pairs that propagate through the entire call graph (Service A -> B -> C -> D). Useful for `TenantID` or `FeatureFlag`. `Baggage: "debug=true"`. Service D sees this and turns on verbose logging for that specific request, even though A started it.

---

### 960. How do you handle breaking proto changes?
"I don't.
I adhere to:
1.  Never remove fields.
2.  Never rename fields.
3.  Never change field IDs.
4.  Only add optional fields.
If I really must break, I create a new `v2` package/service."

#### Indepth
**Field Mask**. Sometimes you want to update only specific fields. `UpdateUser(User{Name: "Bob"})`. Does this mean "Set Email to empty" or "Don't touch Email"? Protobuf uses `FieldMask`. The request contains a list of paths `["name"]`. The server only updates fields listed in the mask. This allows safe partial updates.
