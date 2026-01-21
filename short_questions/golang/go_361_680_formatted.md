# Go Programming - Questions 361-680 (Summary Version)

> **Quick reference guide with concise explanations for Go interview questions 361-680**

---

## 🟣 Cloud-Native and Distributed Systems in Go (Questions 361-380)

**Q361: How do you build a cloud-agnostic app in Go?**
Abstract cloud services behind interfaces. Use common SDK patterns. Configuration-driven provider selection. Avoid vendor-specific code in core logic.

**Q362: How do you use Go SDKs with AWS (S3, Lambda)?**
Import aws-sdk-go-v2. Configure credentials. Create service clients. Call operations with context.

**Q363: How do you upload a file to S3 using Go?**
Create S3 client with session. Use PutObject with bucket, key, and reader. Handle multipart for large files.

**Q364: How do you create a Pub/Sub system using Go and GCP?**
Use cloud.google.com/go/pubsub package. Create topics and subscriptions. Publish messages, receive with callbacks.

**Q365: How would you implement cloud-native config loading?**
Environment variables as primary source. Config files for complex settings. Remote config (etcd, Consul). Secrets from vault services.

**Q366: What is the role of service meshes with Go apps?**
Service-to-service communication, load balancing, observability, security (mTLS). Examples: Istio, Linkerd. Sidecar proxies.

**Q367: How do you secure service-to-service communication in Go?**
Mutual TLS (mTLS). Token-based auth (JWT). Service mesh. API keys. Encrypt in transit.

**Q368: How do you implement service registration and discovery?**
Register with service registry (Consul, etcd). Health checks. Query registry for service instances. Client-side load balancing.

**Q369: How do you manage retries and circuit breakers in Go?**
Use libraries like go-resilience, hystrix-go. Exponential backoff for retries. Circuit breaker pattern: open/half-open/closed states.

**Q370: How would you use etcd/Consul with Go for KV storage?**
Client library for connection. Put/Get operations. Watch for changes. Distributed configuration and coordination.

**Q371: What is leader election and how can you implement it in Go?**
Ensures single leader in distributed system. Use etcd or Consul for coordination. Lock acquisition, maintain lease.

**Q372: How do you build a distributed lock in Go?**
Use Redis SETNX, etcd, or Consul. Acquire lock with TTL. Release explicitly. Handle expiration.

**Q373: How would you implement a distributed queue in Go?**
Message broker (Kafka, RabbitMQ). Or build with Redis lists. Persistent storage. Consumer groups.

**Q374: How do you handle consistency in distributed Go systems?**
Choose consistency model (strong, eventual). Use distributed transactions carefully. Idempotency. Versioning. Conflict resolution.

**Q375: How do you monitor and trace distributed Go systems?**
Distributed tracing (Jaeger, Zipkin, OpenTelemetry). Correlation IDs. Centralized metrics and logging. Service dependency mapping.

**Q376: How do you implement eventual consistency in Go?**
Async propagation of changes. Event sourcing. Message queues. Reconciliation processes. Accept temporary inconsistencies.

**Q377: How do you replicate state in distributed Go apps?**
Leader-follower replication. Raft consensus. Multi-master with conflict resolution. State sync protocols.

**Q378: How do you detect and handle split-brain scenarios?**
Quorum-based decisions. Fencing tokens. Consensus algorithms (Raft, Paxos). Automated recovery or manual intervention.

**Q379: How do you implement quorum reads/writes in Go?**
Require majority of nodes to acknowledge. Consistency over availability. Track successful operations. Fail if quorum not reached.

**Q380: How would you build a simple distributed cache in Go?**
Consistent hashing for key distribution. Multiple cache nodes. Replication factor. Client library routing. Invalidation strategy.

---

## 🔴 Go in Real-World Projects & Architecture (Questions 381-400)

**Q381: How do you handle config versioning in Go projects?**
Version config files alongside code. Migration scripts for breaking changes. Backward compatibility. Feature flags for gradual rollout.

**Q382: How do you organize API versioning in Go apps?**
URL path versioning (/v1/, /v2/). Separate handler packages. Shared business logic. Deprecation notices.

**Q383: How do you validate struct fields with custom rules?**
Use validator package with struct tags. Custom validation functions. Error messages. Validate on input binding.

**Q384: How do you cache API responses in Go?**
In-memory cache (map with mutex, or sync.Map). Redis for distributed. Cache-Control headers. TTL and invalidation.

**Q385: How do you serve files over HTTP with conditional GET?**
Check If-Modified-Since, If-None-Match headers. Return 304 Not Modified. Set ETag, Last-Modified headers.

**Q386: How do you apply SOLID principles in Go?**
Single Responsibility: focused types. Open/Closed: interfaces for extension. Liskov: interface implementations. Interface Segregation: small interfaces. Dependency Inversion: depend on interfaces.

**Q387: How do you prevent breaking changes in shared Go modules?**
Semantic versioning. Deprecation warnings. New major version for breaking changes. Maintain backward compatibility within major version.

**Q388: What is the difference between horizontal and vertical scaling in Go services?**
Horizontal: add more instances (Go excels here). Vertical: increase resources per instance. Horizontal better for Go's concurrency model.

**Q389: How do you support internationalization in Go?**
golang.org/x/text package. Message catalogs. Locale detection. Format numbers, dates per locale. Translation strings.

**Q390: How do you write a Go SDK for third-party APIs?**
Client struct with HTTP client. Methods for API endpoints. Request/response types. Error handling. Context support. Documentation and examples.

**Q391: How do you manage request IDs and trace IDs?**
Generate at entry point. Propagate via context. Include in logs. Pass in headers to downstream services.

**Q392: How do you implement audit logging in Go?**
Log all important actions. Include who, what, when, where. Structured logging. Immutable audit trail. Compliance requirements.

**Q393: How would you version a binary CLI in Go?**
Embed version at build time (-ldflags). Version command. Include in user-agent, logs. Semantic versioning.

**Q394: How do you ensure backward compatibility in Go libraries?**
Don't remove/change exported APIs. Add new functions instead. Deprecation period. Follow semantic versioning strictly.

**Q395: How do you handle soft deletes in Go models?**
Add DeletedAt timestamp field. Filter out deleted records in queries. Nullable time or pointer. Periodic hard deletion.

**Q396: How do you refactor a large legacy Go codebase?**
Incremental changes. Add tests first. Extract interfaces. Create new structure alongside old. Migrate piece by piece. Feature flags.

**Q397: How do you maintain a mono-repo with multiple Go modules?**
Separate go.mod per module. Workspace feature (go.work). Consistent tooling. Shared dependencies management. CI/CD coordination.

**Q398: How would you go about building a plugin system in Go?**
Plugin package (limited, same Go version). Or RPC/gRPC based plugins. Interface definitions. Dynamic loading or separate processes.

**Q399: How do you document Go APIs automatically?**
godoc for package docs. Swagger/OpenAPI generation (swaggo). Comments on exported items. Examples in tests.

**Q400: How do you track tech debt and enforce code quality in large Go teams?**
Code reviews. Linters in CI. Document architectural decisions. Regular refactoring time. Metrics tracking. Static analysis tools.

---

## 🔵 Networking and Low-Level Programming (Questions 401-420)

**Q401: How do you create a TCP server in Go?**
net.Listen("tcp", address). Accept connections in loop. Handle each in goroutine. Read/write to connection.

**Q402: How do you create a UDP client in Go?**
net.Dial("udp", address). Or net.DialUDP for more control. Write/read packets. Handle connectionless nature.

**Q403: What is the difference between net.Listen and net.Dial?**
Listen: creates server, accepts incoming connections. Dial: creates client, connects to server.

**Q404: How do you manage TCP connection pools?**
Reuse connections. http.Transport does this automatically. Custom pool with channel of connections. Set limits.

**Q405: How would you implement a custom HTTP transport?**
Implement http.RoundTripper interface. RoundTrip method handles request/response. Custom connection logic, retry, logging.

**Q406: How do you read raw packets using gopacket?**
pcap library for packet capture. Parse layers. Filter packets. Analyze network traffic.

**Q407: What is a connection hijack in net/http and how is it done?**
Take over HTTP connection. Hijacker interface. Get underlying net.Conn. Use for WebSockets, custom protocols.

**Q408: How to implement a proxy server in Go?**
httputil.ReverseProxy for HTTP. Or custom logic: accept request, forward to backend, return response. Handle headers, modify if needed.

**Q409: How would you create an HTTP2 server from scratch in Go?**
http.Server with TLSConfig. HTTP/2 enabled by default with TLS. golang.org/x/net/http2 for low-level control.

**Q410: How does Go handle connection reuse (keep-alive)?**
HTTP/1.1 keep-alive enabled by default. Connection pooling in Transport. MaxIdleConns settings.

**Q411: How do you set timeouts on sockets in Go?**
SetDeadline, SetReadDeadline, SetWriteDeadline on net.Conn. Or context timeouts for higher-level operations.

**Q412: What is the difference between net/http and fasthttp?**
net/http: standard, stable, feature-complete. fasthttp: optimized for speed, less memory, limited features, no HTTP/2.

**Q413: How do you throttle network traffic in Go?**
Rate limiter (golang.org/x/time/rate). Token bucket. Limit requests per time window. Buffered channels as semaphores.

**Q414: How would you analyze network latency in Go?**
Measure request duration. Histogram metrics. pprof for profiling. Distributed tracing for cross-service latency.

**Q415: How would you implement WebRTC or peer-to-peer comms?**
Use pion/webrtc library. Signaling server. ICE, STUN, TURN. Media streaming or data channels.

**Q416: How do you simulate a slow network in integration tests?**
Custom RoundTripper adding delays. Network emulation tools (tc, toxiproxy). Inject latency in test environment.

**Q417: What's the difference between connection pooling and multiplexing?**
Pooling: reuse connections sequentially. Multiplexing: multiple requests on single connection (HTTP/2). Both reduce overhead.

**Q418: How do you verify DNS lookups in Go?**
net.LookupHost, net.LookupIP. Mock DNS resolver in tests. DNS-over-HTTPS for secure lookups.

**Q419: How do you use HTTP pipelining in Go?**
Multiple requests on connection without waiting for responses. Limited support in net/http. HTTP/2 multiplexing preferred.

**Q420: How do you implement NAT traversal in Go?**
STUN for reflexive address discovery. TURN for relaying. ICE protocol. Libraries like pion for WebRTC NAT traversal.

---

## 🟣 Error Handling & Observability (Questions 421-440)

**Q421: How do you create custom error types in Go?**
Struct implementing error interface. Embed additional fields. Error() method returns string message.

**Q422: How does Go 1.20+ errors.Join and errors.Is work?**
Join: combine multiple errors. Is: check if target error in chain. Supports wrapped errors.

**Q423: How do you implement error wrapping and unwrapping?**
fmt.Errorf with %w verb wraps errors. errors.Unwrap extracts wrapped error. Errors.Is/As work with wrapped errors.

**Q424: What are best practices for error categorization?**
Define error types by category (validation, network, business). Sentinel errors for known cases. Wrap with context.

**Q425: How do you handle critical vs recoverable errors?**
Critical: panic or fatal exit. Recoverable: return error, retry, fallback. Log differently. Alert on critical.

**Q426: How do you recover from panics in goroutines?**
defer with recover() in each goroutine. Can't recover from other goroutine's panic. Handle and log.

**Q427: How to capture stack traces on error?**
errors package with stack. Third-party: pkg/errors. runtime.Stack for manual capture. Include in logs.

**Q428: How do you notify Sentry/Bugsnag from Go?**
SDK integration. Initialize client. Capture errors and panics. Add context. Configure release tracking.

**Q429: How do you do structured error reporting in Go?**
Include error details as structured fields. Machine-readable format. Error codes. Context information.

**Q430: How do you correlate logs, errors, and traces together?**
Common trace/request ID. Include in all logs and errors. Propagate through context. Link in observability platform.

**Q431: How would you add distributed tracing to an existing Go service?**
OpenTelemetry SDK. Instrument HTTP handlers, DB calls. Create spans for operations. Export to backend (Jaeger).

**Q432: What are tags, attributes, and spans in tracing?**
Span: unit of work. Tags/attributes: metadata on spans. Tags describe operation. Group and filter traces.

**Q433: What is a traceparent header?**
W3C trace context standard. Propagates trace information. Version, trace ID, parent ID, flags.

**Q434: How do you send custom metrics to Prometheus?**
prometheus/client_golang library. Define metrics (counter, gauge, histogram). Expose /metrics endpoint. Scrape configuration.

**Q435: What is RED metrics model and how do you apply it?**
Rate (requests/sec), Errors (failure rate), Duration (latency). Essential service metrics. Track for all services.

**Q436: How do you expose application health and readiness probes?**
/health endpoint for liveness. /ready endpoint for readiness. Return 200 when healthy. Check dependencies in readiness.

**Q437: What's the difference between logs, metrics, and traces?**
Logs: discrete events, detailed. Metrics: aggregated measurements, compact. Traces: request flow, distributed. Three pillars of observability.

**Q438: How do you benchmark error impact on performance?**
Benchmark with and without errors. Measure allocation, CPU. Error wrapping overhead. Compare error handling strategies.

**Q439: What's the tradeoff between verbose and silent error handling?**
Verbose: better debugging, noise. Silent: cleaner code, harder to debug. Balance with appropriate logging levels.

**Q440: How would you enforce observability in a Go microservice?**
Middleware for automatic instrumentation. Required trace/log context. Code review standards. Observability as requirement.

---

## 🟢 CLI Tools, Automation, and Scripting (Questions 441-460)

**Q441: How do you build an interactive CLI in Go?**
Use survey or promptui libraries. Handle user input. Validation. Progress bars. Terminal UI libraries (tview, termui).

**Q442: What libraries do you use for command-line tools in Go?**
cobra: commands and subcommands. viper: configuration. pflag: POSIX flags. urfave/cli: simple CLI builder.

**Q443: How do you parse flags and config in CLI?**
flag package for basic flags. cobra+viper for advanced. Environment variable binding. Config file loading.

**Q444: How do you implement bash autocompletion for Go CLI?**
cobra.GenBashCompletion. Generate completion scripts. Install to shell completion directory. Support for multiple shells.

**Q445: How would you use cobra to build a nested command CLI?**
Root command with subcommands. AddCommand to build hierarchy. Flags at different levels. Examples: kubectl, docker.

**Q446: How do you manage color and styling in terminal output?**
fatih/color or charmbracelet/lipgloss. ANSI color codes. Detect terminal capabilities. Disable colors in non-TTY.

**Q447: How would you stream CLI output like tail -f?**
bufio.Scanner for reading. Print as lines arrive. Watch file for changes (fsnotify). Flush output.

**Q448: How do you handle secrets securely in a CLI?**
Prompt for input (hide typing). Environment variables. Config files with restricted permissions. External secret stores.

**Q449: How do you bundle a CLI as a standalone binary?**
go build produces static binary. Embed resources with go:embed. Cross-compile for platforms. Single executable.

**Q450: How would you version and release CLI with GitHub Actions?**
goreleaser for automated releases. Tag-based versioning. Build for multiple platforms. GitHub releases with artifacts.

**Q451: How do you schedule a Go CLI tool with cron?**
Cron expression for scheduling. Ensure idempotency. Proper logging. Lock file to prevent concurrent runs.

**Q452: How do you use Go as a scripting language?**
go run for quick scripts. Package main with simple logic. Better than bash for complex logic. Cross-platform.

**Q453: How do you embed templates in your Go CLI tool?**
go:embed directive. template/text or template/html. Execute templates with data. Bundle in binary.

**Q454: How would you create a system daemon in Go?**
Run as background service. systemd service file. Handle signals for graceful shutdown. Logging to syslog or files.

**Q455: What are good patterns for CLI testing?**
Test command execution. Capture stdout/stderr. Mock filesystem. Test flags and validation. Table-driven tests.

**Q456: How do you store and manage CLI state/config files?**
User home directory or XDG dirs. YAML/JSON config files. Config validation. Migration for version changes.

**Q457: How do you secure a CLI for local system access?**
File permissions. Encrypt sensitive data. Avoid storing secrets. Use OS keychain. Validate all inputs.

**Q458: How do you test CLI tools across multiple OS in CI?**
Matrix builds in CI (GitHub Actions). Test on Linux, macOS, Windows. Platform-specific test cases. Build tags.

**Q459: How do you expose analytics and usage for a CLI?**
Optional telemetry with opt-in. Anonymous usage stats. Error reporting. Respect user privacy.

**Q460: How would you build a CLI wrapper for REST APIs?**
HTTP client for API calls. Commands map to endpoints. Handle auth (tokens). Format output (JSON, table).

---

## 🔴 AI, Machine Learning & Data Processing in Go (Questions 461-480)

**Q461: How do you use TensorFlow or ONNX models in Go?**
TensorFlow Go bindings. ONNX Runtime Go. Load model, create session. Input tensors, run inference, get output.

**Q462: What is gorgonia and when would you use it?**
Machine learning library for Go. Define computational graphs. Automatic differentiation. For custom ML models in Go.

**Q463: How do you implement cosine similarity in Go?**
Dot product of normalized vectors. math operations on slices. Used for similarity comparison, embeddings.

**Q464: How would you stream CSV → transform → JSON using pipelines?**
csv.Reader → goroutine pipeline → json.Encoder. Process in stages. Channels between stages. Concurrent transformation.

**Q465: How do you process large datasets using goroutines?**
Worker pool pattern. Partition data. Process chunks concurrently. Aggregate results. Memory-efficient streaming.

**Q466: How do you implement TF-IDF in Go?**
Term frequency calculation. Inverse document frequency. Vector representation. Used for text analysis, search relevance.

**Q467: How do you parse and tokenize text in Go?**
strings.Fields for simple split. regexp for complex patterns. Third-party NLP libraries. Unicode support.

**Q468: How would you embed a local LLM into a Go app?**
C bindings to llama.cpp. ONNX Runtime. gRPC to Python service running model. Manage memory and performance.

**Q469: How do you integrate OpenAI API in Go?**
HTTP client for API requests. Authorization headers. Marshal/unmarshal JSON. Handle rate limits, retries.

**Q470: How do you do prompt engineering for AI from Go?**
Template system for prompts. Variable substitution. Few-shot examples. Chain prompts for complex tasks.

**Q471: How do you use a local vector database with Go?**
Clients for vector DBs (Pinecone, Weaviate, Milvus). Store embeddings. Similarity search. Integration with AI workflows.

**Q472: How would you implement semantic search using Go?**
Generate text embeddings (API or local). Store in vector DB. Query with embedding. Cosine similarity ranking.

**Q473: How would you extract entities using regex or AI?**
regexp for pattern-based extraction. API calls to NER services. Parse response. Structured data extraction.

**Q474: How do you manage model input/output formats in Go?**
Struct definitions matching model schema. JSON/protobuf serialization. Validation. Type conversion.

**Q475: How would you create a chatbot backend with Go?**
API endpoints for messages. Session management. Integration with LLM. Context window management. Streaming responses.

**Q476: How do you build a recommendation engine with Go?**
Collaborative filtering. Matrix operations. Score calculations. Caching. Real-time or batch inference.

**Q477: How would you integrate LangChain-like logic in Go?**
Chain of operations. Prompt templates. Tool/agent pattern. Composition of AI components.

**Q478: How would you cache AI model outputs in Go?**
Hash input to key. Store in Redis/memory. Check cache before API call. TTL for freshness.

**Q479: What is the role of concurrency in AI inference in Go?**
Parallel batch processing. Multiple model requests. Pipeline stages. Reduces latency. Goroutines for async inference.

**Q480: How do you monitor and scale AI pipelines in Go?**
Metrics: latency, throughput, errors. Queue depth monitoring. Horizontal scaling. Model warm-up. Resource utilization.

---

## 🟡 WebAssembly, Blockchain, and Experimental Go (Questions 481-500)

**Q481: What is WebAssembly and how can Go compile to WASM?**
Binary format for web. GOOS=js GOARCH=wasm. Run Go in browser. Interop with JavaScript.

**Q482: How do you share memory between JS and Go in WASM?**
syscall/js package. Access JS objects. Callbacks between Go and JS. Typed arrays for data.

**Q483: What is TinyGo and what are its limitations?**
Go compiler for small environments. Smaller binaries, WASM support. Limited stdlib. Not all Go features.

**Q484: How do you write a smart contract simulator in Go?**
State machine implementation. Transaction processing. Gas metering. Event emission. Testing blockchain logic.

**Q485: What is Tendermint and how does Go power it?**
Byzantine fault tolerant consensus. Cosmos SDK foundation. Go implementation. Powers many blockchain networks.

**Q486: How do you use go-ethereum to interact with smart contracts?**
abigen generates Go bindings from ABI. Connect to Ethereum node. Call contract methods. Send transactions.

**Q487: How do you parse blockchain data using Go?**
RPC client to node. Decode blocks, transactions. Event logs parsing. Indexing for queries.

**Q488: How do you generate and verify ECDSA signatures in Go?**
crypto/ecdsa package. Generate keys. Sign hash. Verify signature. Used in crypto, blockchain.

**Q489: What is the role of Go in decentralized storage (IPFS)?**
go-ipfs implementation. Content addressing. DHT for routing. P2P networking.

**Q490: How would you implement a Merkle Tree in Go?**
Hash leaf nodes. Pair and hash up tree. Root hash for verification. Proofs of inclusion.

**Q491: How do you handle base58 and hex encoding/decoding?**
encoding/hex for hex. Third-party for base58 (btcd). Cryptocurrency addresses. Binary to string.

**Q492: How do you write a deterministic VM interpreter in Go?**
Bytecode execution. OpCodes with handlers. Stack machine. Gas metering. Deterministic execution.

**Q493: How do you simulate a P2P network in Go?**
libp2p for networking. Discovery protocols. Gossip for message propagation. DHT for routing.

**Q494: How do you create a lightweight Go runtime for edge computing?**
Minimal binary size. TinyGo for small footprint. Strip debug info. Static linking. Optimize for resource constraints.

**Q495: How would you handle offline-first apps in Go?**
Local storage (SQLite). Sync when connected. Conflict resolution. Queue for pending operations.

**Q496: What is the future of Generics in Go (beyond v1.22)?**
More constraint types. Generic methods. Performance optimizations. Community evolution.

**Q497: What is fuzz testing and how do you use it in Go?**
Automated input generation. go test -fuzz. Find edge cases. Crash detection. Built-in since Go 1.18.

**Q498: What is the any type in Go and how is it different from interface{}?**
any is alias for interface{}. More readable. Identical functionality. Introduced Go 1.18.

**Q499: What is the latest experimental feature in Go and why is it important?**
Check release notes. Range over functions, iterators. Performance improvements. Language evolution.

**Q500: How do you contribute to the Go open-source project?**
GitHub golang/go. Proposal process. CLA signing. Code contributions. Issue reporting and discussion.

---

## 🔐 Security in Golang (Questions 501-520)

**Q501: How do you prevent SQL injection in Go?**
Parameterized queries. Never concatenate SQL strings. Use placeholders. PreparedStatements. Validate input.

**Q502: How do you securely store user passwords in Go?**
bcrypt, argon2, or scrypt. Never plain text. Salt automatically. Set appropriate cost factor.

**Q503: How do you implement OAuth 2.0 in Go?**
golang.org/x/oauth2 package. Authorization code flow. Token exchange. Refresh tokens. Store securely.

**Q504: What is CSRF and how do you prevent it in Go web apps?**
Cross-Site Request Forgery. Generate tokens per session. Validate on state-changing requests. SameSite cookies.

**Q505: How do you use JWT securely in a Go backend?**
Sign with strong secret. Validate signature. Check expiration. Refresh tokens. HTTPS only. Store securely client-side.

**Q506: How do you validate and sanitize user input in Go?**
Type validation. Regular expressions. Whitelist allowed values. Escape for output context. Length limits.

**Q507: How do you set secure cookies in Go?**
HttpOnly flag. Secure flag (HTTPS). SameSite attribute. Limited scope. Expiration.

**Q508: How do you avoid path traversal vulnerabilities?**
Validate file paths. Filepath.Clean. Restrict to allowed directories. No user input in paths directly.

**Q509: How do you prevent XSS in Go HTML templates?**
Auto-escaping in html/template. Context-aware escaping. Validate input. Content Security Policy.

**Q510: How would you encrypt sensitive fields before storing in DB?**
crypto/aes for encryption. Key management. Encrypt before save, decrypt on load. Field-level encryption.

**Q511: How do you securely generate random strings or tokens?**
crypto/rand, not math/rand. Sufficient entropy. Right length. Base64/hex encoding.

**Q512: How do you verify digital signatures in Go?**
crypto packages (rsa, ecdsa). Public key verification. Hash message. Verify signature against hash.

**Q513: What are best practices for TLS config in Go HTTP servers?**
TLS 1.2+ minimum. Strong cipher suites. Cert verification. HSTS header. Regular cert renewal.

**Q514: How do you implement rate limiting in Go to avoid DDoS?**
Token bucket or leaky bucket. Per IP limits. golang.org/x/time/rate. Return 429. Progressive delays.

**Q515: How do you handle secrets in Go apps (Vault, env, etc.)?**
Environment variables. HashiCorp Vault. AWS Secrets Manager. Never in code. Rotation support.

**Q516: How do you perform mutual TLS authentication in Go?**
Client and server certificates. TLS config with ClientAuth. Verify client cert. Two-way authentication.

**Q517: What is the difference between crypto/rand and math/rand?**
crypto/rand: cryptographically secure, unpredictable. math/rand: pseudorandom, fast, predictable. Use crypto for security.

**Q518: How do you prevent replay attacks using Go?**
Nonces (number used once). Timestamps with tolerance. Request signing. Store used nonces temporarily.

**Q519: How do you build a secure authentication system in Go?**
Password hashing (bcrypt). Session management. MFA support. Rate limiting. Secure password reset. HTTPS.

**Q520: How do you scan Go projects for vulnerabilities?**
govulncheck for dependencies. gosec for code. Regular updates. CI integration. Dependency audits.

---

## 🚀 Performance Optimization (Questions 521-540)

**Q521: How do you benchmark Go code using testing.B?**
Benchmark* functions. b.N loop iterations. go test -bench. Reports ns/op. Useful for comparisons.

**Q522: What tools can you use to profile a Go application?**
pprof for CPU, memory, goroutines. trace for execution. go test -cpuprofile. Runtime profiling endpoints.

**Q523: How does memory allocation affect Go performance?**
Allocations trigger GC. Heap slower than stack. Escape analysis. Reduce allocations for performance.

**Q524: How do you detect and fix memory leaks in Go?**
Memory profiling over time. Look for growing heap. Find unreleased goroutines. pprof heap analysis.

**Q525: How do you avoid unnecessary allocations in hot paths?**
Pre-allocate slices. Reuse buffers (sync.Pool). Avoid string concatenation in loops. Avoid interface boxing.

**Q526: What is escape analysis and how does it impact performance?**
Compiler decides stack vs heap. Stack is faster. -gcflags '-m' to view. Keep values on stack for speed.

**Q527: How do you use pprof to trace CPU usage?**
Import net/http/pprof. Serve debug endpoint. Generate CPU profile. Analyze with go tool pprof. Find hotspots.

**Q528: How do you optimize slice operations for speed?**
Pre-allocate capacity. Avoid repeated appends. Use index assignment when possible. Copy efficiently.

**Q529: What is object pooling and how is it implemented in Go?**
sync.Pool for object reuse. Put/Get methods. Reduces GC pressure. GC can clear pool. For temporary objects.

**Q530: How does GC tuning affect latency in Go services?**
GOGC controls GC frequency. Higher = less GC, more memory. Lower = more GC, less memory. Tune for latency tolerance.

**Q531: How do you measure and reduce goroutine contention?**
Block profiling. Mutex profiling. Identify locks. Reduce critical section size. Lock-free algorithms.

**Q532: What is lock contention and how to identify it in Go?**
Multiple goroutines waiting for lock. Mutex profiling. pprof block profile. Redesign to reduce locking.

**Q533: How do you batch DB operations for better throughput?**
Collect multiple operations. Execute in transaction. Bulk inserts. Prepared statement reuse. Reduces round trips.

**Q534: How would you profile goroutine leaks?**
runtime.NumGoroutine() over time. Goroutine profile. Find stuck goroutines. Check for unclosed channels, missing context cancellation.

**Q535: What are the downsides of excessive goroutines?**
Memory overhead (stack). Scheduler overhead. Context switching. Diminishing returns. Use worker pools.

**Q536: How would you measure and fix cold starts in Go Lambdas?**
Reduce binary size. Lazy initialization. Keep functions warm. Provisioned concurrency. Optimize imports.

**Q537: How do you decide between a map vs slice for performance?**
Slice: sequential access, O(n) search, contiguous memory. Map: key access, O(1) lookup, hash overhead. Context-dependent.

**Q538: How would you write a memory-efficient parser in Go?**
Stream processing. bufio.Scanner. Avoid loading entire input. Reuse buffers. Incremental parsing.

**Q539: How do you use channels efficiently under heavy load?**
Buffer appropriately. Avoid blocking. Select with default. Close properly. Worker pool pattern.

**Q540: When should you use sync.Pool?**
Frequently allocated temporary objects. Benchmarked improvement. Reduce GC pressure. Not for long-lived objects.

---

## 🧪 Testing in Go (Questions 541-560)

**Q541: How do you write table-driven tests in Go?**
Slice of test cases. Loop executing each. t.Run for subtests. Readability and maintenance.

**Q542: What is the difference between t.Fatal and t.Errorf?**
Fatal: stops test immediately. Errorf: reports error, continues. Use Fatal when further testing pointless.

**Q543: How do you use go test -cover to check coverage?**
-cover flag shows percentage. -coverprofile saves data. go tool cover for visualization.

**Q544: How do you mock a database in Go tests?**
sqlmock package. Interface abstraction. Test database. Repository pattern aids testing.

**Q545: How do you unit test HTTP handlers?**
httptest.NewRecorder. httptest.NewRequest. Test handler function. Verify status, body, headers.

**Q546: What is testable design and how does Go encourage it?**
Interfaces enable mocking. Dependency injection. Small, focused functions. Composition. Explicit dependencies.

**Q547: How do you use interfaces to improve testability?**
Depend on interfaces not concrete types. Inject dependencies. Mock interface implementations in tests.

**Q548: How do you write tests for concurrent code in Go?**
-race flag. Multiple goroutines in test. Verify thread safety. Use sync primitives. Check for deadlocks.

**Q549: What is the httptest package and how is it used?**
Testing HTTP clients and servers. NewServer, NewRecorder, NewRequest. Simulate HTTP without network.

**Q550: How do you mock time in tests?**
Interface for time operations. Inject clock. Mock returns controlled time. Or time.Now variable for simple cases.

**Q551: How do you perform integration testing in Go?**
Test multiple components. Real dependencies or testcontainers. Setup/teardown. Slower than unit tests.

**Q552: How do you use testify/mock for mocking dependencies?**
Generate mocks. Define expectations. Verify calls. AssertExpectations. More powerful than manual mocks.

**Q553: How do you run subtests and benchmarks?**
t.Run for subtests. Hierarchical naming. Parallel execution. Filter with -run. -bench for benchmarks.

**Q554: How do you test panic recovery?**
defer/recover in test. Verify panic occurred. Check panic value. Or testify require.Panics.

**Q555: How do you generate test data using faker or random data?**
gofakeit library. Realistic test data. Randomized testing. Avoid hardcoded values.

**Q556: What is golden file testing and when is it useful?**
Compare output to saved "golden" file. Update flag to regenerate. For complex output (HTML, formatted text).

**Q557: How do you automate test workflows with go generate?**
//go:generate comments. Run code generation. Mock generation. Keep generated code current.

**Q558: How do you test CLI apps built with Cobra?**
Execute command with args. Capture stdout/stderr. Verify exit codes. Test flag parsing.

**Q559: What is fuzz testing and how do you do it in Go?**
Automated input generation. F test functions. go test -fuzz. Finds crashes and edge cases. Built-in Go 1.18+.

**Q560: How do you organize test files and test suites?**
*_test.go files. Same or _test package. TestMain for setup. Group related tests. Subtests for organization.

---

## 🧩 API Design, REST/gRPC & Data Models (Questions 561-580)

**Q561: How do you define a RESTful API in Go using Gin or Echo?**
Router setup. Define routes with handlers. HTTP methods (GET, POST, etc). JSON binding and response.

**Q562: How do you version a REST API?**
URL path (/v1/, /v2/). Header-based. Query parameter. Maintain compatibility within version.

**Q563: How do you handle validation of API payloads?**
Struct tags with validator. Binding validation. Return 400 with errors. Custom validation functions.

**Q564: How do you return proper status codes from handlers?**
200 OK, 201 Created, 400 Bad Request, 401 Unauthorized, 404 Not Found, 500 Internal Server Error. Context-appropriate.

**Q565: How do you implement middleware in a Go web API?**
Function wrapping handler. Chain middleware. Execute before/after handler. Logging, auth, recovery.

**Q566: How do you handle pagination in Go APIs?**
Query params: page, limit, offset. Return total count. Cursor-based for large datasets. HATEOAS links.

**Q567: What's the difference between json.Unmarshal vs Decode?**
Unmarshal: from byte slice. Decode: from io.Reader (streaming). Decode better for HTTP bodies.

**Q568: How do you define a gRPC service in Go?**
.proto file with service and message definitions. protoc generates Go code. Implement service interface.

**Q569: How do you handle gRPC errors and return codes?**
status package. Error codes (NotFound, InvalidArgument, etc). WithDetails for context. Client parses status.

**Q570: How do you secure a gRPC service in Go?**
TLS for transport. Token auth via metadata. Interceptors for auth. mTLS for mutual authentication.

**Q571: How do you do field-level validation in proto definitions?**
buf for breaking change detection. protoc plugins. validate proto extension. Custom validation in code.

**Q572: How do you log incoming requests/responses in a Go API?**
Middleware logging. Structured logs. Include: method, path, status, duration, error. Request ID correlation.

**Q573: How do you handle file uploads/downloads in APIs?**
Upload: multipart/form-data, ParseMultipartForm. Download: ServeFile, stream with io.Copy. Size limits.

**Q574: What is OpenAPI/Swagger and how do you generate docs in Go?**
API specification format. swaggo annotations in code. Generate swagger.json. Serve with swagger-ui.

**Q575: How do you serve static files securely in Go?**
http.FileServer. StripPrefix for routing. Restrict to specific directories. Prevent directory traversal.

**Q576: How do you implement a proxy API gateway in Go?**
httputil.ReverseProxy. Route to upstream services. Add auth, rate limiting. Aggregate responses if needed.

**Q577: How do you generate Go code from .proto files?**
protoc compiler. protoc-gen-go plugin. Generate message and service code. Include in build process.

**Q578: How do you integrate gRPC with REST (gRPC-Gateway)?**
grpc-gateway plugin. Generate reverse proxy. Serve both gRPC and REST. Same proto definitions.

**Q579: How do you implement idempotency in APIs?**
Idempotency keys. Store request IDs. Detect duplicate requests. Return same response. Important for payments.

**Q580: What is a contract-first API development approach?**
Define API specification first (OpenAPI, proto). Generate server stubs. Ensures client-server agreement. Documentation upfront.

---

## 🧠 Design Patterns, Architecture & Real-World Scenarios (Questions 581-600)

**Q581: How do you implement the Factory pattern in Go?**
Function returning interface. Hide concrete type creation. Select implementation based on parameters.

**Q582: How do you use the Strategy pattern in Go?**
Interface defining strategy. Multiple implementations. Inject strategy into context. Swap behaviors at runtime.

**Q583: What is the Singleton pattern and how is it safely used in Go?**
Single instance across application. sync.Once ensures one initialization. Useful for config, connections.

**Q584: How do you write a middleware chain in Go?**
Functions wrapping handlers. Return modified handler. Chain with composition. Execute in order.

**Q585: How do you use interfaces to decouple layers?**
Define interfaces in consumer package. Implementations in provider. Dependency inversion. Testable boundaries.

**Q586: How do you implement the Observer pattern using channels?**
Channel for events. Subscribers listen on channel. Publisher sends to channel. Fan-out for multiple observers.

**Q587: What is the repository pattern and when do you use it?**
Abstract data access. Interface for CRUD operations. Implementation handles DB specifics. Decouples business logic from storage.

**Q588: How would you create a CQRS architecture in Go?**
Separate read and write models. Commands modify state. Queries read state. Different optimization strategies.

**Q589: How do you design a plug-in architecture in Go?**
Plugin interfaces. Load implementations. plugin package or RPC-based. Extend functionality without recompiling.

**Q590: What is a "clean architecture" in Go projects?**
Layers: entities, use cases, interfaces, frameworks. Dependency rule: inner layers independent. Business logic isolated.

**Q591: How do you structure a multi-module Go project?**
Separate go.mod per module. Workspace (go.work). Clear boundaries. Versioning independence.

**Q592: How do you decouple business logic from transport layers?**
Core logic in service layer. Handlers translate HTTP/gRPC to service calls. Interfaces between layers.

**Q593: How would you implement retryable jobs in Go?**
Job queue. Retry count in job. Exponential backoff. Dead letter queue for failures. Worker pool.

**Q594: How would you design a billing system in Go?**
Event-driven. Audit trail. Idempotent operations. Transaction safety. Decimal for money. Reconciliation.

**Q595: How would you scale a notification system written in Go?**
Message queue for async processing. Worker pools. Multiple instances. Deduplication. Rate limiting per user.

**Q596: How do you build a real-time leaderboard in Go?**
Redis sorted sets. WebSocket for updates. Rank calculations. Pagination. Caching frequent queries.

**Q597: How would you implement transactional emails in Go?**
Queue-based. SMTP lib or service (SendGrid). Templates. Retry logic. Tracking. Unsubscribe handling.

**Q598: How do you model money and currencies in Go?**
Avoid float. Use integer cents or decimal package (shopspring/decimal). Include currency. Rounding rules.

**Q599: How do you do dependency injection in Go?**
Constructor injection. Interfaces. wire or fx for complex DI. Manual wiring often sufficient.

**Q600: How do you create a rule engine in Go?**
Define rule interface. Evaluate conditions. Execute actions. Compose rules. DSL or data-driven rules.

---

## 🔸 Advanced Concurrency Patterns (Questions 601-620)

**Q601: How do you implement a fan-in pattern in Go?**
Multiple input channels, single output. Goroutine per input merging to output. WaitGroup for completion.

**Q602: How do you implement a fan-out pattern in Go?**
Single input channel, multiple workers. Distribute work. Workers process concurrently. Aggregate results.

**Q603: How do you prevent goroutine leaks in producer-consumer patterns?**
Close channels when done. Context cancellation. Ensure all goroutines can exit. WaitGroup tracking.

**Q604: How would you create a semaphore in Go?**
Buffered channel as semaphore. Acquire: send to channel. Release: receive from channel. Limits concurrency.

**Q605: What's the difference between sync.WaitGroup and sync.Cond?**
WaitGroup: wait for goroutines to complete. Cond: wait/signal for conditions. Different use cases.

**Q606: How do you implement a pub-sub model in Go?**
Channels for topics. Map of subscribers. Broadcast to all. Register/unregister subscriptions.

**Q607: How do you use a context to timeout multiple goroutines?**
WithTimeout context. Pass to all goroutines. Check ctx.Done(). Cancels all when timeout.

**Q608: How do you build a rate-limiting queue with channels?**
Token bucket with channel. time.Ticker for refill. Workers wait for tokens. Limits processing rate.

**Q609: What is a worker pool, and how do you implement it?**
Fixed number of worker goroutines. Jobs channel. Workers consume jobs. Results channel. Controlled concurrency.

**Q610: How do you handle backpressure in channel-based designs?**
Bounded channels. Block producers when full. Drop or queue. Monitor queue depth. Graceful degradation.

**Q611: How do you gracefully shut down workers?**
Signal channel or context cancellation. Workers check for signal. Finish current work. WaitGroup for completion.

**Q612: How do you use sync.Cond for event signaling?**
Mutex with condition variable. Wait for condition. Signal or Broadcast to wake waiters. Complex coordination.

**Q613: How do you prioritize tasks in concurrent processing?**
Multiple queues by priority. Workers check high-priority first. Or priority queue with heap.

**Q614: How do you avoid starvation in goroutines?**
Fair scheduling. Avoid long-running critical sections. Prioritize fairly. Use select with default.

**Q615: How do you detect race conditions without -race flag?**
Code review. Understanding memory model. Careful design. But -race is best tool.

**Q616: How do you trace execution flow in concurrent systems?**
Distributed tracing. Log with goroutine ID (caution). Trace IDs through context. OpenTelemetry.

**Q617: How do you implement exponential backoff with retries in goroutines?**
Time.Sleep with doubling delay. Max retries and max delay. Jitter to avoid thundering herd.

**Q618: How do you structure long-running daemons with concurrency?**
Main goroutine coordinates. Worker goroutines. Signal handling for shutdown. Context for cancellation.

**Q619: How would you implement circuit breakers in Go?**
Track failures. Open circuit after threshold. Half-open for testing. Close on success. Libraries: hystrix-go.

**Q620: How do you handle concurrent map access with minimal locking?**
sync.Map for specific patterns. Or partition map with multiple mutexes. RWMutex for read-heavy.

---

## 🟤 Event-Driven, Pub/Sub & Messaging (Questions 621-640)

**Q621: How do you publish and consume events using NATS in Go?**
nats-go client. Connect to server. Publish to subject. Subscribe with handler. JetStream for persistence.

**Q622: How do you use Apache Kafka in Go with sarama?**
sarama library. Producer sends messages. Consumer group for parallel consumption. Offset management.

**Q623: What are the trade-offs between RabbitMQ and Kafka in Go apps?**
RabbitMQ: flexible routing, per-message ack. Kafka: high throughput, log-based, replay. Choose by use case.

**Q624: How do you manage message acknowledgements in Go consumers?**
Manual ack after processing. Auto-ack for fire-and-forget. Nack for redelivery. At-least-once delivery.

**Q625: How do you handle message deduplication in Go?**
Unique message ID. Store processed IDs. Check before processing. Idempotent handlers.

**Q626: How do you implement a retry queue for failed messages?**
Send to retry topic/queue. Delay before reprocessing. Max retries before DLQ. Exponential backoff.

**Q627: How do you batch message processing efficiently in Go?**
Collect messages. Process batch together. DB bulk operations. Flush on count or timeout.

**Q628: How do you use Google Pub/Sub with Go?**
cloud.google.com/go/pubsub package. Create topics/subscriptions. Publish sync or async. Receive with callbacks.

**Q629: How do you persist event logs for replay in Go?**
Append-only log. Kafka topic. File-based log. Sequence numbers. Enable event sourcing.

**Q630: How do you ensure exactly-once delivery in Go message systems?**
Difficult to achieve. Idempotent processing easier. Transaction outbox pattern. Deduplication.

**Q631: How do you create a lightweight in-memory pub-sub system?**
Map of topics to subscriber channels. Publish sends to all subscribers. Register/unregister. Simple in-process.

**Q632: How do you handle DLQs (Dead Letter Queues) in Go?**
Move failed messages to DLQ. Monitor and alert. Manual inspection and retry. Log for debugging.

**Q633: How do you create idempotent message consumers in Go?**
Store processed message IDs. Idempotent operations (set vs increment). Database constraints.

**Q634: How do you enforce ordering of messages?**
Single consumer per partition. Sequential processing. Ordering key. Trade-off with parallelism.

**Q635: How do you use channels as message queues?**
Buffered channels store messages. Producers send, consumers receive. In-memory, ephemeral. Good for in-process.

**Q636: How do you handle push vs pull consumers?**
Push: server sends to consumer. Pull: consumer requests messages. Pull gives consumer control. Kafka pull model.

**Q637: How do you deal with large payloads in a messaging system?**
Reference pattern: ID in message, payload in storage. Compression. Split large messages. Stream processing.

**Q638: How do you build an event sourcing system in Go?**
Events as source of truth. Event store. Rebuild state from events. Projections. Snapshots for performance.

**Q639: How would you test message-driven systems?**
Mock message broker. Test handlers independently. Integration tests with testcontainers. Verify idempotency.

**Q640: What's the role of event schemas in Go-based systems?**
Define message structure. Protobuf, Avro, JSON Schema. Versioning. Schema registry. Contract enforcement.

---

## 🟢 Go for DevOps & Infrastructure (Questions 641-660)

**Q641: How do you create a custom Kubernetes operator in Go?**
kubebuilder or operator-sdk. CRDs for custom resources. Reconcile loop. Watch cluster state. Controllers.

**Q642: How do you write a Helm plugin in Go?**
Binary in plugin directory. Environment variables from Helm. Extend Helm functionality.

**Q643: How do you use Go for infrastructure automation?**
CLI tools. AWS/GCP/Azure SDKs. Terraform provider development. Scripting replacing bash. System-level programming.

**Q644: How do you write a CLI for managing AWS/GCP resources?**
Cloud SDK integration. Cobra for CLI structure. Auth handling. Output formatting. Resource CRUD operations.

**Q645: How do you use Go to write Terraform providers?**
Terraform plugin SDK. Define resources and data sources. CRUD operations. Schema definition. State management.

**Q646: How do you build a dynamic inventory script in Go for Ansible?**
Output JSON inventory. Query infrastructure. Group hosts. Variables for hosts. Dynamic sources.

**Q647: How do you parse and generate YAML in Go?**
gopkg.in/yaml.v3. Unmarshal to structs. Marshal from structs. Tags for field names.

**Q648: How do you interact with Docker API in Go?**
docker/client library. Create/start/stop containers. Image operations. Network management. Logs and stats.

**Q649: How do you manage Kubernetes CRDs in Go?**
client-go library. Dynamic client. Generated clients. CRUD on custom resources. Watch for changes.

**Q650: How do you write Go code to scale deployments in K8s?**
client-go. Update deployment replicas. Patch operations. HPA for automatic scaling.

**Q651: How do you tail logs from containers using Go?**
Docker/Kubernetes client. ContainerLogs with stream. Follow mode. Output to stdout or file.

**Q652: How do you manage service discovery in Go apps?**
Kubernetes DNS. Consul/etcd client. Service registry. Health checks. Load balancing.

**Q653: How do you build a Kubernetes admission controller in Go?**
Webhook server. Validate or mutate requests. Register webhook config. TLS required. Quick response.

**Q654: How do you build a metrics exporter for Prometheus in Go?**
prometheus/client_golang. Collector interface. Register collector. Expose /metrics. Scrape configuration.

**Q655: How do you set up health checks for a Go microservice?**
/health endpoint. Check dependencies. Liveness and readiness. Return appropriate status codes.

**Q656: How do you build a custom load balancer in Go?**
Listen for connections. Connection pools to backends. Round-robin, least-conn algorithms. Health checks.

**Q657: How do you implement graceful shutdown with Kubernetes SIGTERM?**
Signal handling. Server.Shutdown() with timeout. Finish in-flight requests. Kubernetes waits configured time.

**Q658: How do you use Go with Envoy/Consul service mesh?**
Config generation. API integration. Sidecar pattern. Service registration. Policy enforcement.

**Q659: How do you configure Go apps for 12-factor principles?**
Config from env. Stateless processes. Port binding. Concurrency via processes. Disposability. Logs to stdout.

**Q660: How do you use Go for cloud automation scripts?**
Cloud SDKs. CLI tools. Scheduled tasks. Infrastructure provisioning. Resource management automation.

---

## 🟣 Caching & Storage Systems (Questions 661-680)

**Q661: How do you cache database query results in Go?**
In-memory map or Redis. Key by query params. TTL for expiration. Invalidate on updates. Consider memory limits.

**Q662: How do you use Redis with Go for distributed caching?**
go-redis client. Set/Get operations. Expiration. Pub/Sub. Data structures (lists, sets). Pipeline for batch.

**Q663: How do you implement LRU cache in Go?**
container/list + map. Track access order. Evict least recently used. Capacity limit. O(1) operations.

**Q664: How do you ensure cache invalidation on data update?**
Invalidate on write. Versioning. TTL as safety. Cache-aside pattern. Event-driven invalidation.

**Q665: How do you handle stale reads in Go apps with caching?**
TTL. Refresh cache on access. Cache-through/write-through. Accept eventual consistency. Monitoring.

**Q666: How do you implement a write-through cache in Go?**
Write to cache and DB together. Consistency. Slower writes. Always fresh cache. Transaction coordination.

**Q667: How do you handle concurrency in in-memory caches?**
sync.RWMutex for map. sync.Map for high concurrency. Partitioned locks. Atomic operations where possible.

**Q668: How do you use bloom filters in Go?**
Probabilistic membership test. No false negatives. Low memory. Libraries available. Use before expensive check.

**Q669: How do you build a TTL-based memory cache?**
Store value with expiry time. Check on access. Background goroutine for cleanup. heap for efficient cleanup.

**Q670: How do you use memcached in Go?**
gomemcache client. Set/Get with expiration. Distributed cache. Simple key-value. CAS for concurrency.

**Q671: How do you store large binary blobs in Go?**
Filesystem or object storage (S3). Database for metadata. Stream large files. Chunked upload/download.

**Q672: How do you build an append-only log file storage in Go?**
os.OpenFile with O_APPEND. Write-ahead log. Sequential writes. Rotation. Indexing for reads.

**Q673: How do you use BoltDB or BadgerDB in Go?**
Embedded key-value store. Transactions. Buckets. ACID. Good for local storage. No separate server.

**Q674: How do you structure a file-based key-value store in Go?**
Index file + data file. Hash map in memory. Compaction. Crash recovery. Simple LSM-like.

**Q675: How do you handle distributed caching with Go?**
Consistent hashing. Multiple cache nodes. Client-side routing. Replication. Libraries: groupcache, redis cluster.

**Q676: How do you monitor cache hit/miss ratios in Go?**
Counters for hits and misses. Prometheus metrics. Calculate ratio. Alert on low hit rate. Optimize cache strategy.

**Q677: How do you use consistent hashing for distributed caching?**
Hash nodes and keys to ring. Key served by nearest node. Minimal redistribution on node add/remove. Virtual nodes for balance.

**Q678: How do you build a cache warming strategy in Go?**
Pre-populate cache at startup. Background refresh for popular items. Predictive loading. Avoid cold start.

**Q679: How do you use S3-compatible storage APIs in Go?**
MinIO client or AWS SDK. PutObject, GetObject. Bucket operations. Presigned URLs. Multipart upload.

**Q680: How do you implement local persistent disk caching?**
File-based cache. Directory structure for keys. LRU eviction. Metadata tracking. OS file caching benefit.

---

**END OF FORMATTED SUMMARY - Questions 361-680 Completed**
