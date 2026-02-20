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
