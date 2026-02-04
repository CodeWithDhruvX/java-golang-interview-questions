## 🌐 REST APIs & gRPC Design (Questions 941-960)

### Question 941: How do you design versioned REST APIs in Go?

**Answer:**
(See Q562). URI Path versioning (`/v1/...`).

---

### Question 942: How do you add OpenAPI/Swagger support in Go?

**Answer:**
(See Q574). `swaggo/swag`.

---

### Question 943: How do you handle graceful shutdown of API servers?

**Answer:**
Capture `os.Signal` (INT/TERM).
Call `server.Shutdown(ctx)`.
This stops accepting new connections but allows active handlers to finish (up to ctx timeout).

---

### Question 944: How do you write middleware for logging/auth?

**Answer:**
(See Q565/572). Decorator pattern on `http.Handler`.

---

### Question 945: How do you secure REST APIs using JWT?

**Answer:**
Middleware extracts `Authorization: Bearer <token>`.
Verifies Signature.
Parses Claims (User ID).
Sets User in Request Context. `r.WithContext(ctx)`.

---

### Question 946: How do you design a RESTful file upload service?

**Answer:**
(See Q573). `multipart/form-data`.
Stream to Disk/S3 so RAM doesn't spike.

---

### Question 947: How do you handle CORS in a Go API?

**Answer:**
Middleware.
Set headers:
`Access-Control-Allow-Origin: *`
`Access-Control-Allow-Methods: GET, POST...`
Handle `OPTIONS` preflight requests by returning 200 OK immediately.
Library: `rs/cors`.

---

### Question 948: How do you paginate API responses?

**Answer:**
(See Q566). Limit/Offset or Cursor-based (Base64 encoded ID of last item).

---

### Question 949: How do you implement rate-limiting on APIs?

**Answer:**
(See Q514). `x/time/rate` per User IP.

---

### Question 950: How do you handle multipart/form-data in Go?

**Answer:**
`r.ParseMultipartForm(32 << 20)` (32MB RAM buffer, rest on disk).
Iterate `r.MultipartForm.File`.

---

### Question 951: How do you expose metrics from a Go API?

**Answer:**
Prometheus Middleware (`promhttp`).
Wrap `ServeHTTP`. Measures Latency Histogram and Request Count.

---

### Question 952: How do you mock gRPC services in tests?

**Answer:**
(See Q839).
Or generate mocks using `mockgen` for the `ServiceClient` interface.

---

### Question 953: How do you set up gRPC with reflection?

**Answer:**
`reflection.Register(s)`.
Allows tools like **Postman** or **grpcurl** to inspect the schema at runtime and make requests without having the `.proto` file locally.

---

### Question 954: How do you stream data over gRPC?

**Answer:**
Define `stream` keyword in proto.
`rpc Download(Request) returns (stream Chunk);`
Go handler uses `stream.Send(chunk)` loop.

---

### Question 955: How do you version gRPC APIs?

**Answer:**
Package names in Proto: `package myapi.v1;`
In Go, it becomes `myapi_v1`.
You run v1 and v2 servers concurrently on same port (different service names).

---

### Question 956: How do you enforce contracts with protobuf validators?

**Answer:**
(See Q571). PGV.

---

### Question 957: How do you convert REST to gRPC clients?

**Answer:**
Refactor Client logic.
Instead of `http.Post`, call `grpcClient.Create(...)`.
The business inputs stay same.

---

### Question 958: How do you monitor gRPC health checks?

**Answer:**
(See Q655). Standard Health V1 service.
K8s `grpc-health-probe` binary runs in pod to query it.

---

### Question 959: How do you build a gRPC gateway in Go?

**Answer:**
(See Q578). `grpc-gateway` library.

---

### Question 960: How do you throttle gRPC traffic in Go?

**Answer:**
**TapHandle** (In-built rate limiting in gRPC is limited).
Use Interceptor.
`if !limiter.Allow() { return status.ResourceExhausted }`.

---
