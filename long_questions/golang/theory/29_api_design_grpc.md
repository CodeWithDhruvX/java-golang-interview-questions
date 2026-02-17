# 🟢 Go Theory Questions: 561–580 API Design, REST/gRPC & Data Models

## 561. How do you define a RESTful API in Go using Gin or Echo?

**Answer:**
We use a router library like **Gin**.

```go
r := gin.Default()
r.POST("/users", CreateUser)
r.GET("/users/:id", GetUser)
```
The handler receives `c *gin.Context`.
We bind input: `c.ShouldBindJSON(&user)`.
We return output: `c.JSON(200, user)`.
Gin handles the routing trie lookup, middleware chain, and standard error serialization efficiently.

---

## 562. How do you version a REST API?

**Answer:**
**URL Versioning** is standard (`/v1/...`).

In Go, we group routes:
```go
v1 := r.Group("/v1")
v1.GET("/items", ...)
```
This physically separates the routing logic. If v2 requires breaking changes (different JSON structure), we create a new struct `UserV2` and a new handler, keeping v1 operational for legacy clients.

---

## 563. How do you handle validation of API payloads?

**Answer:**
We use struct tags with `go-playground/validator` (builtin to Gin).

`type User struct { Email string \`json:"email" binding:"required,email"\` }`
If validation fails, Gin returns error.
We write a custom Error Middleware to intercept this error and format it into a nice response: `{"code": "INVALID_EMAIL", "field": "email"}` instead of returning the raw developer error string.

---

## 564. How do you return proper status codes from handlers?

**Answer:**
We follow HTTP semantics strictly.
*   **200 OK**: Synchronous success.
*   **201 Created**: Resource created (POST).
*   **202 Accepted**: Async processing started (Queued).
*   **204 No Content**: Delete success.
*   **400 Bad Request**: Validation failed.
*   **401 Unauthorized**: No/Bad token.
*   **403 Forbidden**: Valid token, but no permission.
*   **404 Not Found**: ID doesn't exist.
*   **500 Internal Error**: DB crashed.

---

## 565. How do you implement middleware in a Go web API?

**Answer:**
Middleware is a function that wraps a Handler.
Signature: `func(next http.Handler) http.Handler`.

Logic:
1.  **Pre-processing**: Read header, start timer.
2.  **Next**: Call `next.ServeHTTP(w, r)`.
3.  **Post-processing**: Log duration, intercept status code.
Global middleware (Logger, Recovery) runs for all routes. Per-route middleware (Auth) runs only for protected endpoints.

---

## 566. How do you handle pagination in Go APIs?

**Answer:**
We generally use **Cursor-based** or **Offset-based** pagination.

**Offset**: `?page=2&limit=10`. Easy (`LIMIT 10 OFFSET 10`). Performance kills DB at offset 1,000,000.
**Cursor**: `?after=LAST_ID&limit=10`. Fast (`WHERE id > LAST_ID LIMIT 10`).
In the Go response, we include a `next_page_token` (base64 encoded cursor) in the JSON envelope so the client can fetch the next batch.

---

## 567. What’s the difference between `json.Unmarshal` vs `Decode`?

**Answer:**
**Unmarshal**: Takes `[]byte` (all data in memory). Use this if you already have the data (e.g., from a database string).
**Decode**: Takes `io.Reader` (stream). Use this for **HTTP Requests**.

`json.NewDecoder(r.Body).Decode(&v)` reads from the network socket directly into the struct, avoiding the need to buffer the huge JSON body into a temporary byte slice. It is more memory efficient.

---

## 568. How do you define a gRPC service in Go?

**Answer:**
We define it in a `.proto` file (Protocol Buffers).

```protobuf
service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}
```
We run `protoc` to generate `user_grpc.pb.go`.
Our Go struct implements the generated interface `UnsafeUserServiceServer`.
We register it: `pb.RegisterUserServiceServer(grpcServer, &myServiceImpl{})`.

---

## 569. How do you handle gRPC errors and return codes?

**Answer:**
We do *not* return standard Go errors. We return **gRPC Status Errors**.

`return status.Errorf(codes.NotFound, "user %s not found", id)`
The codes (`NotFound`, `InvalidArgument`, `Unauthenticated`) map to standard HTTP codes if transcoded but are transport-agnostic.
We can also attach **Details** (metadata) to the error using `status.WithDetails()`, allowing rich error, like "RetryInfo" or "DebugInfo".

---

## 570. How do you secure a gRPC service in Go?

**Answer:**
1.  **Transport Security (TLS)**: Use credentials.NewServerTLSFromFile().
2.  **Per-RPC Credentials (Auth)**: We use Interceptors (Middleware).
The client sends `Authorization: Bearer <token>` in metadata.
The UnaryServerInterceptor extracts `ctx.Value`, validates the JWT, and rejects the call (`codes.Unauthenticated`) if invalid, before it reaches the handler.

---

## 571. How do you do field-level validation in proto definitions?

**Answer:**
Standard Protobuf is just types, no validation.
We use **protoc-gen-validate (PGV)** or **Buf**.

Annotation in .proto:
`string email = 1 [(validate.rules).string.email = true];`
The generator creates a `.Validate()` method for the Go struct.
In the interceptor, we call `req.Validate()`. If it fails, we return `InvalidArgument` automatically. This enforces constraints at the schema level.

---

## 572. How do you log incoming requests/responses in a Go API?

**Answer:**
**REST**: Use middleware (Logrus/Zap wrapper). Logs Method, Path, Latency, Status.
**gRPC**: Use Interceptors (`grpc_zap`, `grpc_ctxtags`).
Careful: Do **not** log the Body/Payload by default. It might contain PII (passwords, emails). Only log metadata. Log bodies only in Debug mode or for specific non-sensitive routes.

---

## 573. How do you handle file uploads/downloads in APIs?

**Answer:**
**REST**: `multipart/form-data`.
Go: `r.FormFile("file")`. It streams into a temp file. We ignore this for large files and stream directly to S3.

**gRPC**: **Client-Side Streaming**.
The client sends chunks of bytes (`bytes content = 1`).
The server receives a stream. `for { chunk, err := stream.Recv() }`.
We append chunks to a file/S3. This allows uploading 10GB files without loading 10GB into RAM.

---

## 574. What is OpenAPI/Swagger and how do you generate docs in Go?

**Answer:**
OpenAPI is the specification (JSON/YAML). Swagger is the tool.
We use **Swaggo**.
We write comments on handlers.
`// @Param id path int true "User ID"`
`swag init` generates `docs/swagger.json`.
We serve `swagger-ui` middleware. This gives frontend devs a "Try it out" button.

---

## 575. How do you serve static files securely in Go?

**Answer:**
`http.FileServer(http.Dir("/public"))`.

**Security Risk**: Directory Listing.
We must wrap the FileSystem to disable list viewing (return 404/403 for directories).
**Performance**: Go's default static server is okay, but for production, we usually put **NGINX** or a **CDN** (Cloudflare) in front of Go. Go shouldn't spend CPU cycles serving `logo.png`.

---

## 576. How do you implement a proxy API gateway in Go?

**Answer:**
We use `httputil.ReverseProxy`.
Or a library like **KrakenD** or **Ocelot** (Go based).

The Gateway handles:
1.  **Auth termination** (Validates JWT).
2.  **Rate Limiting**.
3.  **Routing**: `/users` -> UserMicroservice:8080.
It effectively multiplexes the frontend traffic to backend services, handling the "Cross-Cutting Concerns" centrally.

---

## 577. How do you generate Go code from `.proto` files?

**Answer:**
We use `protoc` (Protocol Buffers Compiler) with Go plugins.

Command:
`protoc --go_out=. --go-grpc_out=. user.proto`
`go_out` generates the Structs.
`go-grpc_out` generates the Interface and Client/Server stubs.
We usually put this in a `Makefile` or use **Buf** (a modern replacement for protoc) which handles dependency management (`buf generate`).

---

## 578. How do you integrate gRPC with REST (gRPC-Gateway)?

**Answer:**
**gRPC-Gateway** is a plugin that reverse-proxies JSON REST requests to gRPC.

In `.proto`:
`option (google.api.http) = { get: "/v1/users/{id}" };`
The generator creates a REST server that listens on port 8081.
When you call `GET /v1/users/1`, it marshals the JSON to Protobuf, calls the gRPC method on port 8080, gets the response, and marshals it back to JSON. Result: **Dual-Stack** API for free.

---

## 579. How do you implement idempotency in APIs?

**Answer:**
The client sends a header `Idempotency-Key: UUID`.

Middleware checks Redis: `GET key`.
If found: Return the **Saved Response** immediately (do not process).
If not found:
1.  Lock the key (SETNX ... "PROCESSING").
2.  Process request.
3.  Save Response to Redis.
This prevents double-charging a credit card if the network times out during the ACK.

---

## 580. What is a contract-first API development approach?

**Answer:**
We write the **API Spec** (OpenAPI or Proto) **First**.
We agree on it with Frontend/Mobile teams.
Then we generate the Go interfaces/Server stubs.
Then we implement the logic.
This prevents "Drift". If code changes, it doesn't match the spec. It forces documentation to be the Source of Truth, ensuring parallel development (Frontend can mock against the spec while Backend builds).
