## 🌐 REST APIs & gRPC Design (Questions 941-960)

### Question 941: How do you design versioned REST APIs in Go?

**Answer:**
(See Q562). URI Path versioning (`/v1/...`).

### Explanation
Versioned REST APIs in Go use URI path versioning with /v1/ prefixes. This approach maintains backward compatibility while allowing API evolution without breaking existing clients.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you design versioned REST APIs in Go?
**Your Response:** "I design versioned REST APIs using URI path versioning - I prefix endpoints with `/v1/` like `/api/v1/users`. This approach is clear and explicit about which version clients are using. When I need to make breaking changes, I create a `/v2/` version while keeping `/v1/` running for existing clients. I can run both versions concurrently and gradually migrate clients. I also document the differences between versions and provide migration guides. The key is planning for versioning from the start - designing clean APIs that can evolve. I might also use version-specific structs and handlers to keep the code organized. This approach gives me flexibility to improve the API while maintaining backward compatibility."

---

### Question 942: How do you add OpenAPI/Swagger support in Go?

**Answer:**
(See Q574). `swaggo/swag`.

### Explanation
OpenAPI/Swagger support in Go uses swaggo/swag library to automatically generate API documentation from code annotations, providing interactive API documentation and client generation capabilities.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you add OpenAPI/Swagger support in Go?
**Your Response:** "I add OpenAPI/Swagger support using the `swaggo/swag` library. I add annotations to my Go handlers that describe the endpoints, parameters, and responses. Then I run the swag tool to automatically generate the Swagger documentation. This creates an interactive API explorer that developers can use to test endpoints. The annotations include things like HTTP methods, request/response models, and authentication requirements. The generated documentation can be served from my application and stays in sync with the code. This approach is much better than maintaining separate documentation files. I can also generate client SDKs from the OpenAPI spec. The key is adding comprehensive annotations and keeping them updated as the API evolves."

---

### Question 943: How do you handle graceful shutdown of API servers?

**Answer:**
Capture `os.Signal` (INT/TERM).
Call `server.Shutdown(ctx)`.
This stops accepting new connections but allows active handlers to finish (up to ctx timeout).

### Explanation
Graceful shutdown of API servers captures os.Signal for INT/TERM, calls server.Shutdown with context, stops accepting new connections while allowing active handlers to finish within the timeout period.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle graceful shutdown of API servers?
**Your Response:** "I handle graceful shutdown by capturing operating system signals like SIGINT and SIGTERM. When I receive these signals, I call `server.Shutdown(ctx)` with a context that has a timeout. This tells the server to stop accepting new connections but gives active handlers time to finish their work. I set a reasonable timeout - maybe 30 seconds - to balance between completing requests and shutting down quickly. I also close database connections and other resources in the shutdown process. The key is ensuring in-flight requests complete successfully while stopping new work. This prevents dropped connections and data corruption during deployments or restarts. I test the graceful shutdown process to ensure it works reliably."

---

### Question 944: How do you write middleware for logging/auth?

**Answer:**
(See Q565/572). Decorator pattern on `http.Handler`.

### Explanation
Middleware for logging/auth in Go uses the decorator pattern on http.Handler. Middleware functions wrap handlers to add cross-cutting concerns like logging and authentication before passing control to the next handler.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write middleware for logging/auth?
**Your Response:** "I write middleware using the decorator pattern on `http.Handler`. A middleware function takes a handler and returns a new handler that adds behavior before and/or after calling the original handler. For logging, I might log the request details, time the request, and log the response. For authentication, I check credentials or tokens before passing the request to the next handler. I chain multiple middleware together to create a processing pipeline. The key is understanding that middleware wraps handlers and controls whether to call the next handler. I also use context to pass data between middleware and handlers. This pattern keeps my handlers focused on business logic while handling cross-cutting concerns in reusable middleware."

---

### Question 945: How do you secure REST APIs using JWT?

**Answer:**
Middleware extracts `Authorization: Bearer <token>`.
Verifies Signature.
Parses Claims (User ID).
Sets User in Request Context. `r.WithContext(ctx)`.

### Explanation
JWT security in REST APIs uses middleware to extract Authorization Bearer tokens, verify signatures, parse claims like User ID, and set user information in request context using r.WithContext(ctx) for downstream handlers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure REST APIs using JWT?
**Your Response:** "I secure REST APIs with JWT using middleware that extracts the token from the `Authorization: Bearer <token>` header. The middleware verifies the token signature using the secret key, parses the claims to get user information, and sets the user data in the request context using `r.WithContext(ctx)`. Downstream handlers can then retrieve the user from context. I handle token expiration, invalid signatures, and missing tokens appropriately. I also refresh tokens when needed and implement proper error responses. The key is keeping the security logic in middleware so it's consistent across all endpoints. I also use HTTPS to protect the token in transit. This approach provides stateless authentication that scales well."

---

### Question 946: How do you design a RESTful file upload service?

**Answer:**
(See Q573). `multipart/form-data`.
Stream to Disk/S3 so RAM doesn't spike.

### Explanation
RESTful file upload services use multipart/form-data for file transfers. Files are streamed directly to disk or S3 to prevent memory spikes, allowing large file uploads without consuming excessive RAM.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you design a RESTful file upload service?
**Your Response:** "I design file upload services using `multipart/form-data` which is the standard for file uploads. I stream files directly to disk or cloud storage like S3 instead of keeping them in memory to prevent RAM spikes. I handle multiple files, validate file types and sizes, and generate unique filenames. I might also implement progress tracking for large uploads. The key is streaming rather than buffering - reading chunks and writing them immediately to storage. I also handle errors gracefully and clean up partial uploads. For security, I validate file contents and scan for malware. The streaming approach allows uploads of any size without memory issues. I also provide metadata about the upload like file size and type in the response."

---

### Question 947: How do you handle CORS in a Go API?

**Answer:**
Middleware.
Set headers:
`Access-Control-Allow-Origin: *`
`Access-Control-Allow-Methods: GET, POST...`
Handle `OPTIONS` preflight requests by returning 200 OK immediately.
Library: `rs/cors`.

### Explanation
CORS handling in Go uses middleware to set appropriate headers like Access-Control-Allow-Origin and Access-Control-Allow-Methods. OPTIONS preflight requests are handled by returning 200 OK immediately. Libraries like rs/cors simplify this process.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle CORS in a Go API?
**Your Response:** "I handle CORS using middleware that sets the appropriate headers. I set `Access-Control-Allow-Origin` to specify which origins can access the API, and `Access-Control-Allow-Methods` to list allowed HTTP methods. For preflight OPTIONS requests, I immediately return 200 OK without processing the actual request. I can implement this manually or use libraries like `rs/cors` which handle the complexity. I configure CORS based on the environment - more permissive for development, more restrictive for production. The key is understanding that CORS is a browser security feature and the headers tell the browser which cross-origin requests are allowed. I also handle credentials and custom headers when needed."

---

### Question 948: How do you paginate API responses?

**Answer:**
(See Q566). Limit/Offset or Cursor-based (Base64 encoded ID of last item).

### Explanation
API response pagination uses limit/offset for simple cases or cursor-based pagination using Base64 encoded IDs of the last item for more efficient and stable pagination, especially in frequently changing datasets.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you paginate API responses?
**Your Response:** "I paginate API responses using either limit/offset or cursor-based pagination. Limit/offset is simpler - I use `LIMIT n OFFSET m` in SQL queries. For cursor-based pagination, I encode the ID of the last item as a Base64 cursor and use it to fetch the next page. Cursor pagination is more stable when data changes frequently because it doesn't skip or duplicate items. I include pagination metadata in the response like total count, current page, and links to next/previous pages. The key is choosing the right approach based on the data characteristics. For large datasets or real-time data, cursor pagination works better. For static data, limit/offset is simpler. I also handle edge cases like empty pages and invalid cursors."

---

### Question 949: How do you implement rate-limiting on APIs?

**Answer:**
(See Q514). `x/time/rate` per User IP.

### Explanation
API rate limiting uses x/time/rate package per user IP to control request frequency. This prevents abuse and ensures fair resource allocation across different clients accessing the API.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement rate-limiting on APIs?
**Your Response:** "I implement rate limiting using Go's `x/time/rate` package, typically per user IP. I create a rate limiter for each IP address that allows a certain number of requests per second. When a request comes in, I check if the limiter allows it - if not, I return a 429 Too Many Requests status. I might implement different limits for different endpoints or user tiers. I also use sliding window algorithms for more sophisticated limiting. The key is protecting the API from abuse while not blocking legitimate users. I store limiters in a map with expiration to manage memory. I also provide rate limit headers in responses so clients know their limits. This approach prevents DoS attacks and ensures fair usage."

---

### Question 950: How do you handle multipart/form-data in Go?

**Answer:**
`r.ParseMultipartForm(32 << 20)` (32MB RAM buffer, rest on disk).
Iterate `r.MultipartForm.File`.

### Explanation
Multipart/form-data handling in Go uses r.ParseMultipartForm with memory buffer size (32MB) - files larger than this are stored on disk. The r.MultipartForm.File map is then iterated to access uploaded files.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle multipart/form-data in Go?
**Your Response:** "I handle multipart/form-data by calling `r.ParseMultipartForm(32 << 20)` which sets a 32MB memory buffer. Files smaller than this stay in memory, larger ones are written to temporary disk files. Then I iterate through `r.MultipartForm.File` to access each uploaded file. I can get file headers, open the files, and process them - either saving to permanent storage or processing in memory. The key is understanding the memory vs disk tradeoff and setting an appropriate buffer size. I also handle multiple files in one request and validate file types and sizes. I make sure to clean up temporary files when done. This approach handles both small and large uploads efficiently."

---

### Question 951: How do you expose metrics from a Go API?

**Answer:**
Prometheus Middleware (`promhttp`).
Wrap `ServeHTTP`. Measures Latency Histogram and Request Count.

### Explanation
Metrics exposure from Go APIs uses Prometheus middleware with promhttp. It wraps ServeHTTP to measure latency histograms and request counts, providing observability into API performance and usage patterns.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you expose metrics from a Go API?
**Your Response:** "I expose metrics using Prometheus middleware. I wrap the `ServeHTTP` method to automatically collect metrics like request latency histograms and request counters. I use the `promhttp` handler to expose these metrics on a `/metrics` endpoint that Prometheus can scrape. I also create custom metrics for business-specific metrics like active users or database connections. The middleware automatically tracks HTTP status codes, request durations, and response sizes. The key is having comprehensive metrics to monitor the API's health and performance. I also instrument critical business operations to track their performance. This approach gives me deep visibility into how the API is performing and being used."

---

### Question 952: How do you mock gRPC services in tests?

**Answer:**
(See Q839).
Or generate mocks using `mockgen` for the `ServiceClient` interface.

### Explanation
gRPC service mocking uses bufconn for in-memory testing or generates mocks using mockgen for the ServiceClient interface. Both approaches enable unit testing of gRPC client code without actual server dependencies.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you mock gRPC services in tests?
**Your Response:** "I mock gRPC services in two ways. First, I use `bufconn` to create an in-memory gRPC server that my tests can connect to without network overhead. Second, I generate mocks using `mockgen` for the `ServiceClient` interface, which lets me create mock clients that return predefined responses. The bufconn approach is good for integration tests, while mockgen is better for unit tests. I can set up expected responses, error conditions, and verify that my client code handles different scenarios correctly. The key is isolating the code under test from the actual gRPC server. I also test error handling, timeouts, and retry logic. This approach makes my tests fast and reliable while still testing the gRPC interaction logic."

---

### Question 953: How do you set up gRPC with reflection?

**Answer:**
`reflection.Register(s)`.
Allows tools like **Postman** or **grpcurl** to inspect the schema at runtime and make requests without having the `.proto` file locally.

### Explanation
gRPC reflection uses reflection.Register(s) to enable runtime schema inspection. Tools like Postman or grpcurl can discover services and methods without requiring local .proto files, enabling dynamic API exploration.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you set up gRPC with reflection?
**Your Response:** "I set up gRPC reflection by calling `reflection.Register(s)` on my gRPC server. This enables server reflection which allows tools like Postman or grpcurl to discover the services and methods at runtime. The reflection service exposes the protobuf schema so clients can inspect available APIs without having the .proto files locally. This is incredibly useful for debugging and testing - I can use grpcurl to make ad-hoc requests to my service. It also helps with API exploration and documentation. The key is that reflection makes the gRPC API self-describing. I typically enable it in development and staging environments but might disable it in production for security. This feature makes gRPC much more approachable for developers."

---

### Question 954: How do you stream data over gRPC?

**Answer:**
Define `stream` keyword in proto.
`rpc Download(Request) returns (stream Chunk);`
Go handler uses `stream.Send(chunk)` loop.

### Explanation
gRPC streaming uses the stream keyword in proto definitions like rpc Download(Request) returns (stream Chunk). The Go handler uses stream.Send(chunk) in a loop to send multiple responses to the client.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you stream data over gRPC?
**Your Response:** "I stream data over gRPC by defining the `stream` keyword in my proto file. For server streaming, I define something like `rpc Download(Request) returns (stream Chunk)`. In the Go handler, I use a loop with `stream.Send(chunk)` to send multiple chunks of data to the client. For client streaming or bidirectional streaming, I use similar patterns but with `stream.Recv()` instead. The key is understanding that streaming allows sending multiple messages over a single connection, which is much more efficient than multiple RPC calls. I handle errors and connection closure properly. I use streaming for things like file downloads, real-time updates, or processing large datasets. The gRPC runtime handles all the complexity of maintaining the stream connection."

---

### Question 955: How do you version gRPC APIs?

**Answer:**
Package names in Proto: `package myapi.v1;`
In Go, it becomes `myapi_v1`.
You run v1 and v2 servers concurrently on same port (different service names).

### Explanation
gRPC API versioning uses package names in proto files like package myapi.v1, which becomes myapi_v1 in Go. Multiple versions can run concurrently on the same port with different service names for backward compatibility.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you version gRPC APIs?
**Your Response:** "I version gRPC APIs by using package names in the proto file like `package myapi.v1`. In Go, this becomes `myapi_v1`. I can run v1 and v2 servers concurrently on the same port as long as they have different service names. This allows me to maintain backward compatibility while introducing new features. Clients can continue using the v1 service while new clients use v2. I might also implement a compatibility layer that forwards v1 requests to v2 logic. The key is planning the versioning from the start and designing clean service boundaries. I can gradually deprecate old versions as clients migrate. This approach gives me flexibility to evolve the API without breaking existing clients."

---

### Question 956: How do you enforce contracts with protobuf validators?

**Answer:**
(See Q571). PGV.

### Explanation
Protobuf contract enforcement uses PGV (Protobuf Validation) which adds validation rules to protobuf definitions, ensuring data contracts are enforced at the serialization level for type safety and validation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you enforce contracts with protobuf validators?
**Your Response:** "I enforce contracts using PGV - Protobuf Validation. I add validation rules directly in my proto files using annotations that define constraints like field requirements, string patterns, or numeric ranges. PGV generates validation code that checks these rules when messages are created or serialized. This ensures data contracts are enforced at the protobuf level, providing type safety and validation automatically. I can define things like required fields, email formats, or minimum values. The validation happens both client-side and server-side, ensuring data integrity throughout the system. This approach moves validation logic into the contract definition rather than having separate validation code. The key is defining comprehensive validation rules that match the business requirements."

---

### Question 957: How do you convert REST to gRPC clients?

**Answer:**
Refactor Client logic.
Instead of `http.Post`, call `grpcClient.Create(...)`.
The business inputs stay same.

### Explanation
REST to gRPC client conversion refactors client logic to replace http.Post calls with grpcClient.Create calls. The business inputs and logic remain the same, only the transport layer changes from HTTP to gRPC.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you convert REST to gRPC clients?
**Your Response:** "I convert REST clients to gRPC by refactoring the client logic. I replace `http.Post` calls with `grpcClient.Create(...)` calls to the gRPC service. The business inputs and logic stay the same - I'm just changing the transport layer from HTTP/JSON to gRPC/protobuf. I generate the gRPC client code from the proto files and update the client to use the new interface. The key is that the business logic doesn't change, just how I communicate with the server. I might need to handle different error formats and response structures, but the core functionality remains identical. I also update connection handling and error management for gRPC. This approach allows gradual migration from REST to gRPC without rewriting the entire application."

---

### Question 958: How do you monitor gRPC health checks?

**Answer:**
(See Q655). Standard Health V1 service.
K8s `grpc-health-probe` binary runs in pod to query it.

### Explanation
gRPC health monitoring uses the standard Health V1 service to report service status. Kubernetes uses grpc-health-probe binary in pods to query the health endpoint for readiness and liveness checks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you monitor gRPC health checks?
**Your Response:** "I monitor gRPC health using the standard Health V1 service. I implement the health checking service that reports the status of different services in my application. In Kubernetes, I use the `grpc-health-probe` binary in my pods to query the health endpoint. This allows Kubernetes to determine if the pod is ready and alive. The health service can report on different subsystems - database connections, dependencies, or internal services. I implement checks that return SERVING, NOT_SERVING, or UNKNOWN based on the actual health. The key is having comprehensive health checks that accurately reflect the service's ability to handle requests. This integrates with Kubernetes orchestration for automatic restarts and load balancing."

---

### Question 959: How do you build a gRPC gateway in Go?

**Answer:**
(See Q578). `grpc-gateway` library.

### Explanation
gRPC gateway building uses grpc-gateway library to generate a reverse proxy that translates RESTful JSON API calls to gRPC calls, allowing both REST and gRPC clients to use the same backend services.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a gRPC gateway in Go?
**Your Response:** "I build gRPC gateways using the `grpc-gateway` library. This generates a reverse proxy that translates RESTful JSON API calls to gRPC calls. I annotate my proto files with HTTP options that define how REST endpoints map to gRPC methods. The gateway generates HTTP handlers that receive REST requests, convert them to gRPC calls, and convert the responses back to JSON. This allows me to support both REST and gRPC clients with the same backend implementation. The key is defining clear mappings between REST endpoints and gRPC services. I get the performance benefits of gRPC internally while providing a REST API for clients that can't use gRPC. This approach gives me the best of both worlds."

---

### Question 960: How do you throttle gRPC traffic in Go?

**Answer:**
**TapHandle** (In-built rate limiting in gRPC is limited).
Use Interceptor.
`if !limiter.Allow() { return status.ResourceExhausted }`.

### Explanation
gRPC traffic throttling uses interceptors since built-in rate limiting is limited. Interceptors check if a rate limiter allows the request, returning status.ResourceExhausted if the rate limit is exceeded.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you throttle gRPC traffic in Go?
**Your Response:** "I throttle gRPC traffic using interceptors since gRPC's built-in rate limiting is limited. I create an interceptor that checks a rate limiter before allowing requests to proceed. If the limiter doesn't allow the request, I return `status.ResourceExhausted` to the client. I can implement different throttling strategies - per-client, per-endpoint, or global. The interceptor runs for every request and can track usage patterns. I might use token bucket algorithms or sliding windows for more sophisticated throttling. The key is implementing the throttling logic in an interceptor so it applies consistently across all services. I also provide feedback to clients about rate limits through response headers or gRPC metadata. This approach protects the service from overload while ensuring fair access."

---
