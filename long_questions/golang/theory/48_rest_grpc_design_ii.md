# 🌐 Go Theory Questions: 941–960 REST APIs & gRPC Design

## 941. How do you design versioned REST APIs in Go?

**Answer:**
Strategy: **URL Path Versioning** (Most common).
`/v1/users` and `/v2/users`.
In Go Code:
```go
mux.Handle("/v1/", v1Handler)
mux.Handle("/v2/", v2Handler)
```
We keep separate DTO structs (`UserV1`, `UserV2`) to prevent breaking changes in V1 when V2 evolves.

---

## 942. How do you add OpenAPI/Swagger support in Go?

**Answer:**
We use **swaggo/swag**.
We add comments to handlers:
```go
// @Summary Create User
// @Router /users [post]
func CreateUser(...)
```
Run `swag init`. It generates `docs/docs.go`.
Serve it with `http-swagger`.
This keeps documentation close to code, ensuring it doesn't drift.

---

## 943. How do you handle graceful shutdown of API servers?

**Answer:**
(See Q 657).
1.  Listen for `SIGINT/SIGTERM`.
2.  `ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)`.
3.  `server.Shutdown(ctx)`.
This stops accepting new connections but allows active handlers to finish their work (up to 5s).

---

## 944. How do you write middleware for logging/auth?

**Answer:**
Middleware Pattern: `func(http.Handler) http.Handler`.
```go
func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w, r) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL, time.Since(start))
    })
}
```
We chain them: `Logging(Auth(gzip(handler)))`.

---

## 945. How do you secure REST APIs using JWT?

**Answer:**
Middleware extracts `Authorization: Bearer <token>`.
Validates signature (HMAC/RSA).
If valid, places `Claims` into `r.Context()`.
`ctx := context.WithValue(r.Context(), "user", claims)`.
Handlers retrieve user info from Context.
If invalid, return 401 immediately.

---

## 946. How do you design a RESTful file upload service?

**Answer:**
**Multipart/Form-Data**.
`r.ParseMultipartForm(10 << 20)` (10MB RAM limit).
`file, header, _ := r.FormFile("file")`.
We stream `io.Copy(s3Client, file)`.
Returing JSON: `{"url": "https://..."}`.
For huge files (>1GB), we use **Presigned URLs** so the client uploads directly to S3, bypassing our Go server entirely.

---

## 947. How do you handle CORS in a Go API?

**Answer:**
Middleware `rs/cors`.
Browser sends `OPTIONS` (Preflight).
Server must reply:
`Access-Control-Allow-Origin: https://frontend.com`
`Access-Control-Allow-Methods: GET, POST, PUT`
If we strictly set Origin (not `*`), we prevent malicious sites from calling our API from their frontend.

---

## 948. How do you paginate API responses?

**Answer:**
Cursor-based (preferred) or Offset-based.
Request: `GET /items?limit=10&cursor=TokenXYZ`.
DB: `WHERE id > TokenXYZ LIMIT 10`.
Response:
```json
{
  "data": [...],
  "next_cursor": "TokenABC"
}
```
Offset (`OFFSET 10000`) is slow on large DBs. Cursor is O(1).

---

## 949. How do you implement rate-limiting on APIs?

**Answer:**
**Token Bucket** in Redis.
Key: `rate:ip:1.2.3.4`.
Value: Count.
TTL: 1 minute.
Middleware checks Redis. If count > 100, return `429 Too Many Requests`.
Headers: `X-RateLimit-Remaining: 5`.
Go libraries like `tollbooth` handle this locally (in-memory) for single-instance apps.

---

## 950. How do you handle multipart/form-data in Go?

**Answer:**
(See Q 946).
Critical: `defer file.Close()`.
The `mime/multipart` package creates temporary files on disk if the upload exceeds the memory limit. If we don't close the file handle, we run out of file descriptors (Too Many Open Files).

---

## 951. How do you expose metrics from a Go API?

**Answer:**
**Prometheus**.
Middleware wraps `ServeHTTP`.
Records `request_duration_seconds` (Histogram) and `requests_total` (Counter).
Labels: `path`, `method`, `status`.
We expose `/metrics` for the scraper.
Be careful not to include high-cardinality data (like `userID` or `rawURL`) in labels, or memory explodes.

---

## 952. How do you mock gRPC services in tests?

**Answer:**
`gomock` + `protoc-gen-mock`.
Or simply implement the interface.
`type MockGreeter struct { grpc.ServerStream }`.
In test:
`client := NewGreeterClient(conn)`
`resp, err := client.SayHello(ctx, req)`
The `conn` can be a `bufconn` (in-memory net.Conn) connected to a Mock Server implementation.

---

## 953. How do you set up gRPC with reflection?

**Answer:**
`reflection.Register(s)`.
This allows CLI tools like `grpcurl` or UI tools like Postman to query the server for its schema/proto definition at runtime.
`grpcurl -plaintext localhost:8080 list`.
We enable this in Dev/Staging but disable in Production to hide API surface area.

---

## 954. How do you stream data over gRPC?

**Answer:**
**Server Streaming**: `rpc List(Req) returns (stream Item)`.
Go: `func (s *S) List(req, stream) error`.
Loop: `stream.Send(&Item{...})`.
**Client Streaming**: `rpc Upload(stream Chunk) returns (Status)`.
Loop: `stream.Recv()`.
This allows sending 1GB of data without loading it all into RAM.

---

## 955. How do you version gRPC APIs?

**Answer:**
Package Versioning.
`package my.api.v1;` in `.proto`.
Go package: `github.com/my/api/v1`.
If we change fields destructively, we create `v2`.
Both v1 and v2 servers can run in the same binary (registered on the same gRPC server) to support old and new clients simultaneously during migration.

---

## 956. How do you enforce contracts with protobuf validators?

**Answer:**
`protoc-gen-validate` (PGV).
In `.proto`:
`string email = 1 [(validate.rules).string.email = true];`
The generator creates a `Validate()` method on the generated Go struct.
Interceptor calls `req.Validate()` before business logic.

---

## 957. How do you convert REST to gRPC clients?

**Answer:**
**gRPC-Gateway**.
It generates a Reverse Proxy (Go HTTP Handler) from your `.proto` definition.
It accepts JSON HTTP requests, marshals them to Protobuf, and calls the gRPC service (even in the same process).
This allows you to expose `One logic` as both gRPC (for microservices) and REST (for frontend).

---

## 958. How do you monitor gRPC health checks?

**Answer:**
Standard Health Protocol.
`google.golang.org/grpc/health`.
`healthServer.SetServingStatus("myservice", healthpb.HealthCheckResponse_SERVING)`.
K8s probes call this rpc. If it returns NOT_SERVING (after 3 failures), K8s restarts the pod.

---

## 959. How do you build a gRPC gateway in Go?

**Answer:**
(See Q 957).
We run the gRPC server on port 9090.
We run the Gateway HTTP proxy on port 8080.
Gateway dials 9090.
Or, we use `cmux` to serve both HTTP/1.1 and gRPC (HTTP/2) on the **same port** (multiplexing based on Content-Type).

---

## 960. How do you throttle gRPC traffic in Go?

**Answer:**
**Tap Handle** or **Interceptor**.
`ratelimit` interceptor.
Also `MaxConcurrentStreams` option in `grpc.NewServer`.
This limits the number of active HTTP/2 streams. If exceeded, client receives `RESOURCE_EXHAUSTED` (gRPC status 8).
