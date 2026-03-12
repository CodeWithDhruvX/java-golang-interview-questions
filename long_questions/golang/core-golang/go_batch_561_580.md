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

### Explanation
RESTful API design in Go using frameworks like Gin involves creating a router instance, organizing routes into logical groups (typically by version), and attaching handler functions to specific HTTP methods and paths. The router handles incoming requests and routes them to the appropriate handler based on the HTTP method and URL pattern.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define a RESTful API in Go using Gin or Echo?
**Your Response:** "I define RESTful APIs in Go using frameworks like Gin by creating a router instance with `gin.Default()`, then organizing routes into logical groups using `r.Group()` - typically by API version like '/api/v1'. Within each group, I attach handler functions to specific HTTP methods and paths using `v1.GET()`, `v1.POST()`, etc. Each handler function receives a context parameter that provides access to request data and response writing capabilities. Finally, I start the server with `r.Run()` on the desired port. This approach provides clean organization, versioning support, and follows REST conventions for resource-based URLs and HTTP methods."

---

### Question 562: How do you version a REST API?

**Answer:**
1.  **URL Path:** `/v1/users`, `/v2/users`. Most common/explicit.
2.  **Header:** `Accept: application/vnd.myapi.v1+json`. Cleaner URLs, harder to test in browser.
3.  **Code Structure:** Isolate logic in packages `handler/v1` and `handler/v2` to support backward compatibility.

### Explanation
REST API versioning can be done through URL paths, headers, or query parameters. URL path versioning is most common and explicit. Header-based versioning keeps URLs cleaner but is harder to test in browsers. Code structure should isolate version-specific logic in separate packages to support backward compatibility and clean separation of concerns.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you version a REST API?
**Your Response:** "I version REST APIs using three main approaches. URL path versioning like '/v1/users' and '/v2/users' is the most common and explicit method that's easy to understand and test. Header-based versioning using 'Accept: application/vnd.myapi.v1+json' keeps URLs cleaner but is harder to test in browsers. Regardless of the versioning approach, I organize my code structure by isolating version-specific logic in separate packages like 'handler/v1' and 'handler/v2' to support backward compatibility. This allows me to maintain multiple API versions simultaneously while keeping the codebase clean and manageable. The choice depends on the specific requirements - URL versioning for simplicity and browser compatibility, or header versioning for cleaner URLs."

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

### Explanation
API payload validation in Go frameworks like Gin is handled through struct tags that define validation rules. The go-playground/validator library provides validation tags like 'required', 'email', and 'gte'. The binding process automatically validates incoming JSON against these rules and returns detailed error messages when validation fails.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle validation of API payloads?
**Your Response:** "I handle API payload validation by binding incoming JSON to structs with validation tags. Gin uses the go-playground/validator library by default, which provides comprehensive validation rules. I define a struct with tags like `binding:'required,email'` for email validation and `binding:'gte=18'` for minimum age requirements. When I call `c.ShouldBindJSON(&req)`, Gin automatically validates the request against these rules. If validation fails, I return a 400 status with the error details. This approach provides automatic validation, clear error messages, and keeps my validation logic close to the data structure definition, making the code more maintainable and self-documenting."

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

### Explanation
Proper HTTP status codes should be returned from handlers using constants from the net/http package. Different scenarios require different status codes: 200 for successful GET operations, 201 for resource creation, 204 for successful operations with no response body, 400 for validation failures, 404 for missing resources, and 500 for server errors.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you return proper status codes from handlers?
**Your Response:** "I return proper HTTP status codes using constants from the net/http package. For different scenarios, I use specific status codes: 200 StatusOK for successful GET requests, 201 StatusCreated when a new resource is created, 204 StatusNoContent for successful operations like DELETE that don't return a response body, 400 StatusBadRequest for validation failures, 404 StatusNotFound when a resource doesn't exist, and 500 StatusInternalServerError for server errors or database failures. Using these constants makes the code more readable and ensures I'm following HTTP standards. This approach helps clients understand the outcome of their requests and handle different scenarios appropriately."

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

### Explanation
Middleware in Go web APIs follows the Decorator pattern, wrapping handlers to add cross-cutting concerns like authentication, logging, or rate limiting. Middleware functions take an http.Handler and return a new http.Handler that performs additional processing before or after calling the next handler in the chain.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement middleware in a Go web API?
**Your Response:** "I implement middleware in Go web APIs using the Decorator pattern, where middleware functions take an http.Handler and return a new http.Handler. The middleware can perform processing before and after calling the next handler. For example, an authentication middleware checks for an Authorization header, and if it's missing, returns a 401 error. If authentication passes, it calls `next.ServeHTTP()` to continue the request chain. This approach allows me to add cross-cutting concerns like authentication, logging, rate limiting, or CORS handling without modifying individual handlers. I can chain multiple middleware together to create a processing pipeline that handles all aspects of request processing in a clean, reusable way."

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

### Explanation
Pagination in Go APIs is implemented by accepting page and page_size (or limit/offset) query parameters. The response should include both the data and metadata containing pagination information like current page, total pages, and total records. SQL queries use LIMIT and OFFSET clauses, and default limits should be enforced to prevent returning entire databases.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle pagination in Go APIs?
**Your Response:** "I handle pagination by accepting query parameters like page and page_size, or limit and offset. I return the paginated data wrapped in a response that includes both the actual data and metadata about the pagination. The metadata contains information like the current page number, total pages available, and total record count. In SQL, I implement this using LIMIT and OFFSET clauses. I always enforce default limits to prevent queries that could return the entire database, which would cause performance issues. This approach gives clients control over how much data they receive while providing the context they need to implement proper pagination controls in their user interfaces."

---

### Question 567: What’s the difference between `json.Unmarshal` vs `Decode`?

**Answer:**
- **`json.Unmarshal`:** Takes `[]byte` (requires full JSON in memory). Fast/Simple for small payloads.
- **`json.NewDecoder(r).Decode(&v)`:** Reads from `io.Reader` (Stream). Better for large payloads or reading directly from HTTP Body to avoid copying data to a byte slice first.

### Explanation
The difference between json.Unmarshal and json.Decode is their input source and memory usage. Unmarshal requires the complete JSON as a byte slice in memory, making it fast and simple for small payloads. Decode reads from an io.Reader as a stream, making it more memory-efficient for large payloads and avoiding the need to copy data to a byte slice first.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between `json.Unmarshal` vs `Decode`?
**Your Response:** "The key difference is their input source and memory usage. `json.Unmarshal` takes a byte slice, so it requires the entire JSON document to be loaded into memory first. This makes it fast and simple for small payloads. `json.NewDecoder(r).Decode(&v)` reads from an io.Reader as a stream, which is more memory-efficient for large payloads. I use Decode when reading directly from an HTTP request body to avoid copying the data to a byte slice first. For small configuration files or API responses, Unmarshal is usually fine. For large files or streaming data, Decode is the better choice to avoid high memory usage. The choice depends on the size of the JSON payload and whether I need to process it as a stream."

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

### Explanation
gRPC services in Go are defined using Protocol Buffers in .proto files. The service definition specifies RPC methods with their request and response message types. The protoc compiler generates Go code from these definitions, including client interfaces, server implementations, and message structs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define a gRPC service in Go?
**Your Response:** "I define gRPC services in Go using Protocol Buffers in .proto files. I create a service definition with RPC methods, each specifying request and response message types. For example, I might define a UserService with a GetUser RPC that takes a UserRequest and returns a UserResponse. The messages contain the data structures with field numbers. After defining the service, I run the protoc compiler with Go plugins to generate the actual Go code, including client interfaces, server stubs, and message structs. This generated code handles all the serialization, deserialization, and network communication, allowing me to focus on implementing the business logic."

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

### Explanation
gRPC error handling requires using the status and codes packages rather than returning normal Go errors. The status.Error function creates proper gRPC errors with specific codes that map to HTTP status codes when using gRPC-Gateway. This ensures consistent error handling across gRPC services and proper HTTP status code mapping.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle gRPC errors and return codes?
**Your Response:** "I handle gRPC errors using the google.golang.org/grpc/status and codes packages instead of returning normal Go errors. I use status.Error() to create proper gRPC errors with specific codes like codes.NotFound or codes.InvalidArgument. When I use gRPC-Gateway, these gRPC codes automatically map to appropriate HTTP status codes - for example, codes.NotFound maps to HTTP 404. This approach ensures consistent error handling across my gRPC services and provides proper HTTP status code mapping when clients access the service through REST APIs. The key is to always use the gRPC status package rather than plain Go errors to maintain the gRPC error model."

---

### Question 570: How do you secure a gRPC service in Go?

**Answer:**
1.  **TLS/SSL:** Use credentials when creating the server (`credentials.NewServerTLSFromFile`).
2.  **Interceptors:** Use Unary/Stream Interceptors (middleware) to validate `metadata` (headers) containing JWT tokens for Auth.

### Explanation
Securing gRPC services involves implementing TLS/SSL encryption and authentication. TLS credentials are used when creating the server to encrypt all communication. Interceptors act as middleware to validate metadata (headers) containing authentication tokens like JWT, providing authentication and authorization for gRPC calls.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure a gRPC service in Go?
**Your Response:** "I secure gRPC services using two main approaches. First, I implement TLS/SSL encryption by using credentials when creating the server with `credentials.NewServerTLSFromFile()`, which encrypts all communication between clients and the server. Second, I use interceptors as middleware to validate metadata containing authentication tokens. I create unary and stream interceptors that check for JWT tokens in the metadata headers and validate them before allowing the request to proceed. This combination provides both transport-level security through TLS and application-level security through authentication. The interceptors can handle both authentication and authorization logic, ensuring only authorized clients can access the gRPC services."

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

### Explanation
Protocol Buffers don't include built-in validation enforcement. Field-level validation can be added using protoc-gen-validate or buf.build/validate. These tools allow adding validation rules as annotations in the proto file, and the generated Go code includes Validate() methods that enforce these rules at runtime.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you do field-level validation in proto definitions?
**Your Response:** "Since Protocol Buffers don't include built-in validation, I use tools like protoc-gen-validate or buf.build/validate to add field-level validation. I annotate the proto file with validation rules like `[(validate.rules).string.email = true]` for email validation. When I generate the Go code, it includes a `.Validate()` method that enforces these rules at runtime. This approach allows me to define validation constraints directly in my schema definitions, ensuring consistent validation across all services that use the proto definitions. The generated validation code handles complex rules like string patterns, numeric ranges, and required fields, making it easy to maintain data integrity throughout the system."

---

### Question 572: How do you log incoming requests/responses in a Go API?

**Answer:**
Use Middleware.
- **REST:** Wrap `ServeHTTP`.
- **gRPC:** Use a logging Interceptor.
Log: Method, Path, Duration (Latency), Status Code, and Client IP.
Avoid logging Body directly (PII risk), unless debug mode is on.

### Explanation
Logging incoming requests and responses in Go APIs is implemented using middleware. For REST APIs, this involves wrapping the ServeHTTP method. For gRPC, logging interceptors are used. Important information to log includes the method, path, duration, status code, and client IP. Request/response bodies should be avoided due to PII risks unless in debug mode.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you log incoming requests/responses in a Go API?
**Your Response:** "I log incoming requests and responses using middleware. For REST APIs, I wrap the ServeHTTP method to intercept all requests and responses. For gRPC services, I use logging interceptors that wrap each RPC call. I log key information like the HTTP method or RPC method, request path, duration for latency tracking, status codes, and client IP addresses. I'm careful about logging request or response bodies directly due to PII risks - I only log bodies in debug mode when explicitly enabled. This approach provides comprehensive observability without exposing sensitive user data. The middleware approach ensures consistent logging across all endpoints while keeping the logging logic centralized and maintainable."

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

### Explanation
File uploads in Go APIs are handled using ParseMultipartForm and FormFile to process multipart form data. The uploaded file is read and copied to a destination. File downloads use http.ServeFile which automatically handles MIME type detection and range requests for partial content delivery and resume functionality.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle file uploads/downloads in APIs?
**Your Response:** "I handle file uploads using multipart form processing with `r.FormFile()` to extract uploaded files from the request. I create a destination file and copy the uploaded content using `io.Copy()`. For downloads, I use `http.ServeFile()` which is convenient because it automatically handles MIME type detection and range requests. Range requests are important for supporting features like download resumption and video streaming. The ServeFile function takes care of all the HTTP headers and content negotiation, making it much simpler than implementing file serving manually. This approach provides robust file handling with minimal code while supporting important features like partial content delivery."

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

### Explanation
OpenAPI/Swagger is the standard for describing REST APIs. The swaggo/swag tool generates documentation from code comments. Developers add annotations to handlers describing endpoints, responses, and parameters. The swag init command generates swagger.json and swagger.html files that can be served to provide interactive API documentation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is OpenAPI/Swagger and how do you generate docs in Go?
**Your Response:** "OpenAPI/Swagger is the industry standard for describing REST APIs. In Go, I use the swaggo/swag tool to generate documentation from code comments. I add annotations to my handler functions describing what each endpoint does, its parameters, and response formats. For example, I add comments like `@Summary Get User`, `@Success 200 {object} User`, and `@Router /users/{id} [get]`. Then I run `swag init` which generates swagger.json and interactive HTML documentation. I serve these generated files through a documentation endpoint, providing developers with interactive API documentation they can explore and test directly in their browser. This approach keeps the documentation synchronized with the code since it's generated from the actual implementation."

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

### Explanation
Static file serving in Go uses http.FileServer with http.Dir to create a file system handler. Security considerations include preventing directory traversal attacks and ensuring root directories aren't exposed. Go's http.Dir typically handles '..' safely, but custom implementations require careful path sanitization.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you serve static files securely in Go?
**Your Response:** "I serve static files securely using `http.FileServer` with `http.Dir`. I create a file server for a specific directory like './static' and mount it at a URL path using `http.Handle()` and `http.StripPrefix()`. For security, I ensure I don't expose root directories and carefully sanitize paths to prevent directory traversal attacks. Go's built-in `http.Dir` usually handles '..' sequences safely, but I'm extra careful with any custom file handling. I also make sure to scope the file server to specific subdirectories rather than serving from the file system root. This approach provides secure static file serving while protecting against common path traversal vulnerabilities that could expose sensitive files."

---

### Question 576: How do you implement a proxy API gateway in Go?

**Answer:**
Use `httputil.NewSingleHostReverseProxy`.
You can intercept the request in `Director` to add Auth headers or rewrite paths before forwarding to the upstream microservice.

### Explanation
API gateway proxy implementation in Go uses httputil.NewSingleHostReverseProxy to create a reverse proxy that forwards requests to upstream services. The Director function allows intercepting and modifying requests before forwarding, enabling features like adding authentication headers, path rewriting, or request transformation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a proxy API gateway in Go?
**Your Response:** "I implement API gateway proxies using `httputil.NewSingleHostReverseProxy` which creates a reverse proxy that forwards requests to upstream microservices. The key feature is the Director function, which allows me to intercept and modify requests before they're forwarded. I use this to add authentication headers, rewrite paths to match upstream service expectations, or transform request data. This approach enables building a centralized gateway that handles cross-cutting concerns like authentication, routing, and request transformation while forwarding to the appropriate backend services. It's a clean way to implement API gateway functionality without having to manually handle HTTP forwarding and response proxying."

---

### Question 577: How do you generate Go code from .proto files?

**Answer:**
Use the compiler `protoc` with Go plugins.

```bash
protoc --go_out=. --go-grpc_out=. user.proto
```
This generates `user.pb.go` (structs) and `user_grpc.pb.go` (client/server interfaces).

### Explanation
Go code generation from .proto files uses the protoc compiler with Go plugins. The --go-out flag generates message structs, while --go-grpc-out generates gRPC service interfaces. This produces separate files for data structures and service definitions, enabling both client and server implementations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you generate Go code from .proto files?
**Your Response:** "I generate Go code from .proto files using the protoc compiler with Go plugins. I run `protoc --go_out=. --go-grpc_out=. user.proto` which generates two files: `user.pb.go` containing the message structs and data types, and `user_grpc.pb.go` containing the gRPC client and server interfaces. The --go-out flag handles the Protocol Buffer message generation, while --go-grpc-out generates the actual gRPC service code. This gives me everything I need to implement both gRPC clients and servers in Go. The generated code handles all the serialization, deserialization, and network communication, allowing me to focus on implementing the business logic for my services."

---

### Question 578: How do you integrate gRPC with REST (gRPC-Gateway)?

**Answer:**
**gRPC-Gateway** reads Protobuf service definitions and generates a reverse-proxy server which translates a RESTful JSON API into gRPC.
Enables supporting both REST (for browser JS) and gRPC (for backend services) from one codebase.

### Explanation
gRPC-Gateway bridges REST and gRPC by generating a reverse proxy server from Protobuf service definitions. It translates RESTful JSON API calls into gRPC calls, enabling support for both REST clients (like browser JavaScript) and gRPC clients (like backend services) from a single codebase.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you integrate gRPC with REST (gRPC-Gateway)?
**Your Response:** "I integrate gRPC with REST using gRPC-Gateway, which reads my Protobuf service definitions and generates a reverse proxy server. This proxy translates RESTful JSON API calls into gRPC calls, allowing me to support both REST clients like browser JavaScript and gRPC clients like backend services from a single codebase. I define my services once in Protobuf, and gRPC-Gateway generates the REST endpoints automatically. This approach gives me the performance benefits of gRPC for internal service-to-service communication while maintaining REST API compatibility for external clients. It's the best of both worlds - I get type-safe, high-performance gRPC internally and standard REST APIs externally without maintaining separate implementations."

---

### Question 579: How do you implement idempotency in APIs?

**Answer:**
Crucial for Payments.
1.  Client generates a unique ID (`Idempotency-Key` header).
2.  Server checks Middleware/DB/Redis: "Have I seen this Key?"
    - **Yes:** Return the *saved* response from the previous success (do not re-process).
    - **No:** Process -> Save Key+Response -> Return.

### Explanation
Idempotency in APIs is crucial for operations like payments where retrying the same request shouldn't cause duplicate charges. The client generates a unique Idempotency-Key header. The server checks if it has seen this key before - if yes, it returns the saved response; if no, it processes the request, saves the key and response, then returns the result.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement idempotency in APIs?
**Your Response:** "I implement idempotency in APIs, which is crucial for payment operations, using an Idempotency-Key header approach. The client generates a unique ID and sends it in the Idempotency-Key header. On the server side, I use middleware to check if I've seen this key before using a database or Redis. If I have seen the key, I return the saved response from the previous successful request without reprocessing. If it's a new key, I process the request normally, save the key and response together, then return the result. This ensures that retrying the same request multiple times won't cause duplicate operations like charging a customer twice. It's essential for any API that handles critical operations where retries are common."

---

### Question 580: What is a contract-first API development approach?

**Answer:**
Define the specification **before** writing code.
- **REST:** Write OpenAPI (Swagger) YAML first. Generate Go stubs (`oapi-codegen`).
- **gRPC:** Write Proto files first.
**Benefits:** Client and Server teams can work in parallel; API is well-documented by definition.

### Explanation
Contract-first API development involves defining the API specification before writing implementation code. For REST APIs, this means writing OpenAPI/Swagger YAML first and generating Go stubs. For gRPC, it means writing Proto files first. This approach enables parallel development of client and server teams and ensures the API is well-documented by definition.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a contract-first API development approach?
**Your Response:** "Contract-first API development means I define the API specification before writing any implementation code. For REST APIs, I write the OpenAPI/Swagger YAML specification first, then use tools like `oapi-codegen` to generate Go stubs. For gRPC, I write the Proto files first. The main benefits are that client and server teams can work in parallel since they both agree on the contract, and the API is well-documented by definition rather than being an afterthought. This approach prevents implementation details from leaking into the API design and ensures the API serves the needs of all consumers. It's like having a blueprint before building a house - everyone knows what they're building towards, which leads to better design and fewer integration issues."

---
