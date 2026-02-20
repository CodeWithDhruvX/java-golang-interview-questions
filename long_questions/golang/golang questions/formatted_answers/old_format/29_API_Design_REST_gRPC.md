# 🧩 **561–580: API Design, REST/gRPC & Data Models**

### 561. How do you define a RESTful API in Go using Gin or Echo?
"I define a router and group main resources.
`r := gin.Default()`.
`v1 := r.Group("/v1")`.
`v1.GET("/users", GetUsers)`.
I ensure my handlers accept interfaces (`Context`) so they are testable.
I stick to standard HTTP verbs: POST for create, GET for read, PUT/PATCH for update, DELETE for remove."

#### Indepth
Content Negotiation is often overlooked. Your API should respect the `Accept` header. If the client asks for `application/xml`, return XML or `406 Not Acceptable`. Gin/Echo have helpers like `c.Negotiate()` to handle this automatically, making your API more robust for legacy integrations.

---

### 562. How do you version a REST API?
"I prefer **URL Versioning**.
`/api/v1/users`.
New major breaking changes go to `/api/v2`.
This allows `v1` and `v2` to coexist. I package the code as `handler/v1` and `handler/v2`.
Header versioning is cleaner puristically but harder to debug with simple tools like `curl`."

#### Indepth
Deprecation Policy: When releasing `v2`, don't just kill `v1`. Add a `Warning` header to `v1` responses: `Warning: 299 - "This API is deprecated and will be removed on 2025-01-01"`. This gives clients a programmatic way to detect that they are running on borrowed time.

---

### 563. How do you handle validation of API payloads?
"I use **struct tags** and a library like `validator` (playground/validator).
`type User struct { Email string 'validate:"required,email"' }`.
In the handler: `if err := c.ShouldBindJSON(&user); err != nil { return 400 }`.
Then `validate.Struct(user)`.
This simplifies validation into declarative rules in the model definition, keeping the handler clean."

#### Indepth
Custom Validators: Default rules like `required` aren't enough. Register custom logic: `v.RegisterValidation("is-after-today", validateDate)`. This allows you to say `validate:"required,is-after-today"` in your struct tag, encapsulating complex domain rules directly in the DTO layer.

---

### 564. How do you return proper status codes from handlers?
"I map errors to codes using a helper.
`if errors.Is(err, ErrNotFound) { c.JSON(404, ...) }`.
`if errors.Is(err, ErrInvalidInput) { c.JSON(400, ...) }`.
`else { c.JSON(500, ...) }`.
I always return **201 Created** for POST, not 200. And **204 No Content** for DELETE."

#### Indepth
**401 vs 403**: Confusion is common. **401 Unauthorized** means "I don't know who you are" (Missing/Invalid Token). **403 Forbidden** means "I know who you are, but you can't do this" (Insufficent Role). Getting this wrong breaks frontend auth flows (e.g., redirecting to login when they are already logged in but just lack permission).

---

### 565. How do you implement middleware in a Go web API?
"Middleware is a function that takes a Handler and returns a Handler.
`func AuthMiddleware(next http.Handler) http.Handler`.
It executes logic *before* `next.ServeHTTP`.
Common uses: Logging, CORS, Auth, Rate Limiting.
In Gin, it's `func(c *gin.Context)`, but the concept is the same: do work, call `c.Next()`."

#### Indepth
Context Propagation! Middleware is where you populate the `context.Context`. Generate a `RequestID`, put it in the context. Extract the `UserID` from the JWT, put it in the context. This allows the inner handler (and the logger) to access these values without polluting function signatures.

---

### 566. How do you handle pagination in Go APIs?
"I accept `page` and `limit` (or `cursor`) query params.
Default `limit=10`. Max `limit=100`.
I return a wrapper struct:
`{ "data": [...], "meta": { "total": 100, "page": 1, "next_cursor": "abc" } }`.
For high performance, I use **Cursor Pagination** (`WHERE id > last_seen_id`) instead of Offset (`OFFSET 10000`)."

#### Indepth
Don't expose raw DB IDs in the cursor if possible. Encode the cursor (e.g. base64 of `last_created_at` + `id`). This opaque string (`?cursor=eyJ...`) prevents users from guessing sequences and allows you to change the underlying implementation (like switching from ID to Timestamp) without breaking clients.

---

### 567. What’s the difference between `json.Unmarshal` vs `Decode`?
"**Unmarshal**: Takes a `[]byte`. Reads the whole input into memory first. Good for small payloads.
**Decode**: Takes an `io.Reader`. Streams the data. Good for large payloads or HTTP bodies.
`json.NewDecoder(r.Body).Decode(&v)`.
It’s generally safer to use Decoder for web handlers to avoid buffering huge requests in RAM."

#### Indepth
**DisallowUnknownFields**: By default, Go ignores extra fields in JSON. This hides bugs (client sends `user_nmae` instead of `user_name`). Use `decoder.DisallowUnknownFields()` to make such requests fail immediately, saving you hours of debugging why the name wasn't updating.

---

### 568. How do you define a gRPC service in Go?
"I write a `.proto` file.
`service UserService { rpc GetUser(UserReq) returns (UserResp); }`.
I run `protoc --go_out=. --go-grpc_out=.`.
Then I implement the interface:
`type Server struct { pb.UnimplementedUserServiceServer }`.
`func (s *Server) GetUser(...)`.
This contract-first approach guarantees my API matches the spec."

#### Indepth
Syntax matters. Always use `syntax = "proto3";`. It removes "Required" fields (everything is optional/default). This seems scary but makes backward compatibility easier. "Required" fields in proto2 caused massive outages when a client didn't send a field that the server thought was required but the business logic didn't actually need.

---

### 569. How do you handle gRPC errors and return codes?
"I don't return standard Go errors.
I return `status.Error(codes.NotFound, "user not found")`.
The client receives the specific gRPC status code (NOT_FOUND = 5).
If I need details (like 'invalid field X'), I attach `status.WithDetails(&errdetails.BadRequestFieldViolation{...})`."

#### Indepth
The `google.rpc.Status` message is richer than `error`. It can hold a list of `Any`. Use the standard error model provided by Google (`errdetails` package): `RetryInfo`, `DebugInfo`, `QuotaFailure`. Clients like mobile apps can use these strongly-typed details to show "Retry in 5s" popups automatically.

---

### 570. How do you secure a gRPC service in Go?
"**TLS** functionality is built-in.
`creds := credentials.NewServerTLSFromFile(cert, key)`.
`grpc.NewServer(grpc.Creds(creds))`.
For Auth, I use Interceptors.
`UnaryServerInterceptor` extracts the JWT from `metadata` (headers), validates it, and rejects the call if invalid."

#### Indepth
**ALTS** (Application Layer Transport Security). If running on GCP, you might not need manual TLS certificates. ALTS provides zero-config mTLS between services running on Google infrastructure. for generic setups, use `cert-manager` to rotate short-lived certificates automatically.

---

### 571. How do you do field-level validation in proto definitions?
"Standard Protobuf doesn't have validation logic.
I use **protoc-gen-validate (PGV)** or Buf.
`string email = 1 [(validate.rules).string.email = true];`.
The generated code includes a `Validate()` method. My interceptor calls this method automatically on every request."

#### Indepth
The ecosystem is moving to **CEL** (Common Expression Language). Newer `protoc-gen-validate` versions use CEL to allow complex rules like `message.created_at < message.updated_at`. This logic sits inside the proto definitions, making the API self-documenting and safe by design.

---

### 572. How do you log incoming requests/responses in a Go API?
"I use Middleware / Interceptors.
I prefer structured logging (`slog` / `zap`).
Log: `method`, `path`, `status`, `latency`, `client_ip`.
I verify *not* to log sensitive bodies (passwords).
For gRPC, `grpc_zap` middleware handles this out of the box."

#### Indepth
**Sampling**. In high QPS systems, logging *every* request is too expensive (IO/Storage). Use Dynamic Sampling: Log 100% of errors, but only 1% of successes. This reduces noise while guaranteeing that if things go wrong, you have the data.

---

### 573. How do you handle file uploads/downloads in APIs?
"**Upload**: `MultipartReader` for streaming.
`reader, err := r.MultipartReader()`.
I iterate parts and copy to a temp file or S3 stream.
**Download**: `http.ServeContent` (handles Range requests/resuming).
Or `io.Copy(w, file)`.
I verify to set `Content-Disposition` header so the browser knows the filename."

#### Indepth
Set limits! `r.ParseMultipartForm(10 << 20)` (10MB). If you don't enforce limits, a user can upload a 50GB file and fill your disk (DoS). For massive uploads, prefer **Pre-signed URLs** (S3). The client uploads directly to S3; your Go server just validates permissions and hands out the upload token.

---

### 574. What is OpenAPI/Swagger and how do you generate docs in Go?
"I use **Swag** (swaggo).
I add comments to my handler:
`// @Summary Get User`
`// @Param id path int true "User ID"`
`swag init` generates the `swagger.json`.
I serve it with `swagger-ui` middleware.
This keeps code and docs in sync."

#### Indepth
**Code-First vs Spec-First**. Swag is Code-First (Go -> YAML). This is easier for devs but can lead to "Implementation leakage". Spec-First (OpenAPI -> Go using `oapi-codegen`) is superior for large teams. You agree on the YAML contract first, then frontend and backend build in parallel against the mock.

---

### 575. How do you serve static files securely in Go?
"I use `http.FileServer` but wrap it.
I sanitize the path to prevent directory traversal (`..`).
I set correct MIME types.
I set Cache-Control headers (long cache for hashed assets, no-cache for index.html).
I prefer serving static assets via Nginx/CDN in production, keeping Go for API logic only."

#### Indepth
`embed.FS` (Go 1.16+) changes the game. `http.FS(content)`. You can ship a single binary with the React frontend inside it. But remember, `http.FileServer` uses `ModTime` to handle caching (`304 Not Modified`). When embedding, `ModTime` might be zero or build time, so handle ETags carefully.

---

### 576. How do you implement a proxy API gateway in Go?
"I use `httputil.ReverseProxy`.
I match paths (`/api/v1/users` -> `users-service:8080`).
I can modify the request (add headers) or response.
Tools like **KrakenD** or **Tyk** are written in Go and do exactly this. Writing one from scratch is a good exercise in `http.RoundTripper`."

#### Indepth
**Singleflight**. A proxy is vulnerable to the "Thundering Herd". If 1000 users ask for the same resource, don't make 1000 backend calls. Use `golang.org/x/sync/singleflight` to coalesce them into ONE backend call and share the result. This one line of code can save your backend during traffic spikes.

---

### 577. How do you generate Go code from `.proto` files?
"I use **Buf** (modern tool).
`buf generate`.
It uses `buf.gen.yaml` config.
It manages plugin versions and standardizes the output directory structures.
It ensures my team generates the exact same code on every machine, preventing 'works on my machine' diffs."

#### Indepth
Look at **Connect-Go** (by Buf). It's a modern replacement for `grpc-go` that works over HTTP/1.1 effortlessly (no special proxy needed). It generates cleaner, more idiomatic Go code and supports the standard library `http.Handler` unlike standard gRPC which needs its own server/listener.

---

### 578. How do you integrate gRPC with REST (gRPC-Gateway)?
"I add annotations to my `.proto`.
`option (google.api.http) = { get: "/v1/users/{id}" };`.
I run `protoc-gen-grpc-gateway`.
It generates a reverse proxy in Go that listens on HTTP JSON, translates to Protobuf, and calls my gRPC server.
This gives me the best of both worlds: gRPC for internal microservices, REST for public clients."

#### Indepth
The syntax `google.api.http` is powerful. You can map body fields: `body: "*"`. Or map path parameters: `/v1/{book_id}/shelves/{shelf_id}`. `grpc-gateway` handles the type conversion (string "123" in URL -> int64 123 in Proto) automatically, rejecting invalid types with 400 Bad Request.

---

### 579. How do you implement idempotency in APIs?
"I check the **Idempotency-Key** header.
I store the key + response in Redis.
If the key exists, I return the stored response.
If not, I process (using a lock to prevent concurrent processing of the same key).
This is critical for Payment APIs to avoid double charging on network retries."

#### Indepth
The store must be atomic. `SETNX key "processing" EX 30`. If result is 0, someone else is doing it. If 1, proceed. When done, `SET key response`. If the process crashes mid-way, the key expires (30s), allowing a retry. This is the **Distributed Lock** pattern applied to API requests.

---

### 580. What is a contract-first API development approach?
"I write the **Swagger/OpenAPI** or **Proto** spec *before* writing code.
I generate the server stubs and client libraries from the spec.
This forces me to think about the data model and endpoints cleanly without getting bogged down in implementation details. It allows frontend and backend to work in parallel."

#### Indepth
**Breaking Changes**. With Spec-First, you can catch breaking changes in CI using `buf breaking` or `openapi-diff`. It compares `api.v1.yaml` vs `api.v2.yaml` and fails the build if you removed a field or changed a type, enforcing semantic versioning rigor automatically.
