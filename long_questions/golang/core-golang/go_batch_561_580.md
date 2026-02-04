## 🧩 API Design, REST/gRPC & Data Models (Questions 561-580)

### Question 561: How do you define a RESTful API in Go using Gin or Echo?

**Answer:**
Using **Gin** as an example:
Define a router, group routes, and attach handlers.

```go
r := gin.Default()
v1 := r.Group("/api/v1")
{
    v1.GET("/users", GetUsers)
    v1.POST("/users", CreateUser)
}
r.Run(":8080")
```

---

### Question 562: How do you version a REST API?

**Answer:**
1.  **URL Path:** `/v1/users`, `/v2/users`. Most common/explicit.
2.  **Header:** `Accept: application/vnd.myapi.v1+json`. Cleaner URLs, harder to test in browser.
3.  **Code Structure:** Isolate logic in packages `handler/v1` and `handler/v2` to support backward compatibility.

---

### Question 563: How do you handle validation of API payloads?

**Answer:**
Bind the JSON to a struct with **validation tags**.
Gin uses `go-playground/validator` by default.

```go
type CreateRequest struct {
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=18"`
}

func CreateUser(c *gin.Context) {
    var req CreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
}
```

---

### Question 564: How do you return proper status codes from handlers?

**Answer:**
Use the constants in `net/http` package.
- 200 `StatusOK`: Success.
- 201 `StatusCreated`: Resource created.
- 204 `StatusNoContent`: Success but no body (Delete).
- 400 `StatusBadRequest`: Validation fail.
- 404 `StatusNotFound`: Resource missing.
- 500 `StatusInternalServerError`: Crashing/DB fail.

---

### Question 565: How do you implement middleware in a Go web API?

**Answer:**
Middleware mimics the Decorator pattern. It takes `http.Handler` and returns `http.Handler`.

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Authorization") == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

---

### Question 566: How do you handle pagination in Go APIs?

**Answer:**
Accept `page` and `page_size` (or `limit`/`offset`) as query parameters.
Return a metadata envelope.

```json
{
  "data": [...],
  "meta": {
    "page": 1,
    "total_pages": 5,
    "total_records": 50
  }
}
```
In SQL: `LIMIT ? OFFSET ?`. Use specific default limits to avoid dumping the whole DB.

---

### Question 567: What’s the difference between `json.Unmarshal` vs `Decode`?

**Answer:**
- **`json.Unmarshal`:** Takes `[]byte` (requires full JSON in memory). Fast/Simple for small payloads.
- **`json.NewDecoder(r).Decode(&v)`:** Reads from `io.Reader` (Stream). Better for large payloads or reading directly from HTTP Body to avoid copying data to a byte slice first.

---

### Question 568: How do you define a gRPC service in Go?

**Answer:**
Define in a `.proto` file (Protocol Buffers).

```protobuf
service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}
message UserRequest { string id = 1; }
message UserResponse { string name = 1; }
```
Then run `protoc` to generate Go code.

---

### Question 569: How do you handle gRPC errors and return codes?

**Answer:**
Use the `google.golang.org/grpc/status` and `codes` packages.
Do not return normal Go errors.

```go
if user == nil {
    return nil, status.Error(codes.NotFound, "user not found")
}
```
This maps to HTTP status codes (NotFound -> 404) if using gRPC-Gateway.

---

### Question 570: How do you secure a gRPC service in Go?

**Answer:**
1.  **TLS/SSL:** Use credentials when creating the server (`credentials.NewServerTLSFromFile`).
2.  **Interceptors:** Use Unary/Stream Interceptors (middleware) to validate `metadata` (headers) containing JWT tokens for Auth.

---

### Question 571: How do you do field-level validation in proto definitions?

**Answer:**
Protobuf itself doesn't enforce validation.
Use **protoc-gen-validate (PGV)** or **buf.build/validate**.
Annotate the proto:
```protobuf
string email = 1 [(validate.rules).string.email = true];
```
The generated Go code will have a `.Validate()` method.

---

### Question 572: How do you log incoming requests/responses in a Go API?

**Answer:**
Use Middleware.
- **REST:** Wrap `ServeHTTP`.
- **gRPC:** Use a logging Interceptor.
Log: Method, Path, Duration (Latency), Status Code, and Client IP.
Avoid logging Body directly (PII risk), unless debug mode is on.

---

### Question 573: How do you handle file uploads/downloads in APIs?

**Answer:**
- **Upload:** Use `ParseMultipartForm` and `FormFile`.
    ```go
    file, header, _ := r.FormFile("upload")
    dst, _ := os.Create(header.Filename)
    io.Copy(dst, file)
    ```
- **Download:** `http.ServeFile(w, r, "path/to/file")` handles MIME types and Range requests automatically.

---

### Question 574: What is OpenAPI/Swagger and how do you generate docs in Go?

**Answer:**
Standard for describing REST APIs.
Use **swaggo/swag**.
1.  Add comments to handlers:
    ```go
    // @Summary Get User
    // @Success 200 {object} User
    // @Router /users/{id} [get]
    func GetUser(...) {}
    ```
2.  Run `swag init`.
3.  Serve the generated `docs/swagger.json`.

---

### Question 575: How do you serve static files securely in Go?

**Answer:**
Use `http.FileServer`.
**Security:** Ensure you don't expose root directories.
Sanitize paths to prevent Directory Traversal (Go's `http.Dir` usually handles `..` safely, but be careful with custom handlers).

```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

---

### Question 576: How do you implement a proxy API gateway in Go?

**Answer:**
Use `httputil.NewSingleHostReverseProxy`.
You can intercept the request in `Director` to add Auth headers or rewrite paths before forwarding to the upstream microservice.

---

### Question 577: How do you generate Go code from .proto files?

**Answer:**
Use the compiler `protoc` with Go plugins.

```bash
protoc --go_out=. --go-grpc_out=. user.proto
```
This generates `user.pb.go` (structs) and `user_grpc.pb.go` (client/server interfaces).

---

### Question 578: How do you integrate gRPC with REST (gRPC-Gateway)?

**Answer:**
**gRPC-Gateway** reads Protobuf service definitions and generates a reverse-proxy server which translates a RESTful JSON API into gRPC.
Enables supporting both REST (for browser JS) and gRPC (for backend services) from one codebase.

---

### Question 579: How do you implement idempotency in APIs?

**Answer:**
Crucial for Payments.
1.  Client generates a unique ID (`Idempotency-Key` header).
2.  Server checks Middleware/DB/Redis: "Have I seen this Key?"
    - **Yes:** Return the *saved* response from the previous success (do not re-process).
    - **No:** Process -> Save Key+Response -> Return.

---

### Question 580: What is a contract-first API development approach?

**Answer:**
Define the specification **before** writing code.
- **REST:** Write OpenAPI (Swagger) YAML first. Generate Go stubs (`oapi-codegen`).
- **gRPC:** Write Proto files first.
**Benefits:** Client and Server teams can work in parallel; API is well-documented by definition.

---
