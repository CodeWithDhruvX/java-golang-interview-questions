1. What is a function literal (anonymous function)?
2. How does the `net/http` package work?

### 🟢 **1–20: Basics**
3. What is Go and who developed it?
4. What are the key features of Go?
5. How do you declare a variable in Go?
6. What are the data types in Go?
7. What is the zero value in Go?
8. How do you define a constant in Go?
9. Explain the difference between `var`, `:=`, and `const`.
10. What is the purpose of `init()` function in Go?
11. How do you write a for loop in Go?
12. What is the difference between `break`, `continue`, and `goto`?
13. What is a `defer` statement?
14. How does `defer` work with return values?
15. What are named return values?
16. What are variadic functions?
17. What is a type alias?
18. What is the difference between `new()` and `make()`?
19. How do you handle errors in Go?
20. What is panic and recover in Go?
21. What are blank identifiers in Go?

### 🟡 **21–40: Arrays, Slices, and Maps**
22. What is the difference between an array and a slice?
23. How do you append to a slice?
24. What happens when a slice is appended beyond its capacity?
25. How do you copy slices?
26. What is the difference between len() and cap()?
27. How do you create a multi-dimensional slice?
28. How are slices passed to functions (by value or reference)?
29. What are maps in Go?
30. How do you check if a key exists in a map?
31. Can maps be compared directly?
32. What happens if you delete a key from a map that doesn’t exist?
33. Can slices be used as map keys?
34. How do you iterate over a map?
35. How do you sort a map by key or value?
36. What are struct types in Go?
37. How do you define and use struct tags?
38. How to embed one struct into another?
39. How do you compare two structs?
40. What is the difference between shallow and deep copy in structs?
41. How do you convert a struct to JSON?

### 🔵 **41–60: Pointers, Interfaces, and Methods**
42. What are pointers in Go?
43. How do you declare and use pointers?
44. What is the difference between pointer and value receivers?
45. What are methods in Go?
46. How to define an interface?
47. What is the empty interface in Go?
48. How do you perform type assertion?
49. How to check if a type implements an interface?
50. Can interfaces be embedded?
51. What is polymorphism in Go?
52. How to use interfaces to write mockable code?
53. What is the difference between `interface{}` and `any`?
54. What is duck typing?
55. Can you create an interface with no methods?
56. Can structs implement multiple interfaces?
57. What is the difference between concrete type and interface type?
58. How to handle nil interfaces?
59. What are method sets?
60. Can a pointer implement an interface?
61. What is the use of `reflect` package?

### 🟣 **61–80: Concurrency and Goroutines**
62. What are goroutines?
63. How do you start a goroutine?
64. What is a channel in Go?
65. What is the difference between buffered and unbuffered channels?
66. How do you close a channel?
67. What happens when you send to a closed channel?
68. How to detect a closed channel while receiving?
69. What is the `select` statement in Go?
70. How do you implement timeouts with `select`?
71. What is a `sync.WaitGroup`?
72. How does `sync.Mutex` work?
73. What is `sync.Once`?
74. How do you avoid race conditions?
75. What is the Go memory model?
76. How do you use `context.Context` for cancellation?
77. How to pass data between goroutines?
78. What is the `runtime.GOMAXPROCS()` function?
79. How do you detect deadlocks in Go?
80. What are worker pools and how do you implement them?
81. How to write concurrent-safe data structures?

### 🔴 **81–100: Advanced & Best Practices**
82. How does Go handle memory management?
83. What is garbage collection in Go?
84. How do you profile CPU and memory in Go?
85. What is the difference between compile-time and runtime errors?
86. How to use `go test` for unit testing?
87. What is table-driven testing in Go?
88. How to benchmark code in Go?
89. What is `go mod` and how does it work?
90. What is vendoring in Go modules?
91. How to handle versioning in modules?
92. How do you structure a Go project?
93. What is the idiomatic way to name Go packages?
94. What is the purpose of the `internal` package?
95. How do you handle logging in Go?
96. What is the difference between `log.Fatal`, `log.Panic`, and `log.Println`?
97. What are build tags in Go?
98. What are cgo and its use cases?
99. What are some common Go anti-patterns?
100. What are Go code quality tools (lint, vet, staticcheck)?
101. What are the best practices for writing idiomatic Go code?

### 🟢 **101–120: Project Structure & Design Patterns**
102. How do you organize a large-scale Go project?
103. What is the standard Go project layout?
104. What is the `cmd` directory used for in Go?
105. How do you structure code for reusable packages?
106. What are Go's most used design patterns?
107. Explain the Factory Pattern in Go.
108. How to implement Singleton Pattern in Go?
109. What is Dependency Injection in Go?
110. What is the difference between composition and inheritance in Go?
111. What are Go generics and how do you use them?
112. How to implement a generic function with constraints?
113. What are type parameters?
114. Can you implement the Strategy pattern using interfaces?
115. What is middleware in Go web apps?
116. How do you structure code using the Clean Architecture?
117. What are service and repository layers?
118. How would you separate concerns in a RESTful Go app?
119. What is the importance of interfaces in layered design?
120. How would you implement a plugin system in Go?
121. How do you avoid circular dependencies in Go packages?

### 🟡 **121–140: Generics, Type System, and Advanced Types**
122. What is type inference in Go?
123. How do you use generics with struct types?
124. Can you restrict generic types using constraints?
125. How to create reusable generic containers (e.g., Stack)?
126. What is the difference between `any` and interface{}?
127. Can you have multiple constraints in a generic function?
128. Can interfaces be used in generics?
129. What is type embedding and how does it differ from inheritance?
130. How does Go perform type conversion vs. type assertion?
131. What are tagged unions and how can you simulate them in Go?
132. What is the use of `iota` in Go?
133. How are custom types different from type aliases?
134. What are type sets in Go 1.18+?
135. Can generic types implement interfaces?
136. How do you handle constraints with operations like +, -, *?
137. What is structural typing?
138. Explain the difference between concrete and abstract types.
139. What are phantom types and are they used in Go?
140. How would you implement an enum pattern in Go?
141. How can you implement optional values in Go idiomatically?

### 🔵 **141–160: Networking, APIs, and Web Dev**
142. How to build a REST API in Go?
143. How to parse JSON and XML in Go?
144. What is the use of `http.Handler` and `http.HandlerFunc`?
145. How do you implement middleware manually in Go?
146. How do you serve static files in Go?
147. How do you handle CORS in Go?
148. What are context-based timeouts in HTTP servers?
149. How do you make HTTP requests in Go?
150. How do you manage connection pooling in Go?
151. What is an HTTP client timeout?
152. How do you upload and download files via HTTP?
153. What is graceful shutdown and how do you implement it?
154. How to work with multipart/form-data in Go?
155. How do you implement rate limiting in Go?
156. What is Gorilla Mux and how does it compare with net/http?
157. What are Go frameworks for web APIs (Gin, Echo)?
158. What are the trade-offs between using `http.ServeMux` and third-party routers?
159. How would you implement authentication in a Go API?
160. How do you implement file streaming in Go?

### 🟣 **161–180: Databases and ORMs**
161. How do you connect to a PostgreSQL database in Go?
162. What is the difference between `database/sql` and GORM?
163. How do you handle SQL injections in Go?
164. How do you manage connection pools in `database/sql`?
165. What are prepared statements in Go?
166. How do you map SQL rows to structs?
167. What are transactions and how are they implemented in Go?
168. How do you handle database migrations in Go?
169. What is the use of `sqlx` in Go?
170. What are the pros and cons of using an ORM in Go?
171. How would you implement pagination in SQL queries?
172. How do you log SQL queries in Go?
173. What is the N+1 problem in ORMs and how to avoid it?
174. How do you implement caching for DB queries in Go?
175. How do you write custom SQL queries using GORM?
176. How do you handle one-to-many and many-to-many relationships in GORM?
177. How would you structure your database layer in a Go project?
178. What is context propagation in database calls?
179. How do you handle long-running queries or timeouts?
180. How do you write unit tests for code that interacts with the DB?

### 🔴 **181–200: Tools, Testing, CI/CD, Ecosystem**
181. What is `go vet` and what does it catch?
182. How does `go fmt` help maintain code quality?
183. What is `golangci-lint`?
184. What is the difference between `go run`, `go build`, and `go install`?
185. How does `go generate` work?
186. What is a build constraint?
187. How do you write tests in Go?
188. How do you test for expected panics?
189. What are mocks and how do you use them in Go?
190. How do you use the `testing` and `testify` packages?
191. How do you structure test files in Go?
192. What is a benchmark test?
193. How do you measure test coverage in Go?
194. How do you test concurrent functions?
195. What is a race detector and how do you use it?
196. What is `go.mod` and `go.sum`?
197. How does semantic versioning work in Go modules?
198. How to build and deploy a Go binary to production?
199. What tools are used for Dockerizing Go apps?
200. How do you set up a CI/CD pipeline for a Go project?

### 🟢 **201–220: Performance & Optimization**
201. How do you optimize memory usage in Go?
202. What is memory escape analysis in Go?
203. How to reduce allocations in tight loops?
204. How do you profile a Go application?
205. What is the use of `pprof` in Go?
206. How do you benchmark against memory allocations?
207. How can you avoid unnecessary heap allocations?
208. What is inlining and how does the Go compiler handle it?
209. How do you debug GC pauses?
210. What are some common performance bottlenecks in Go apps?
211. How to detect and fix memory leaks?
212. How do you find goroutine leaks?
213. How do you tune GC parameters in production?
214. How to avoid blocking operations in hot paths?
215. What are the trade-offs of pooling in Go?
216. How do you measure latency and throughput in Go APIs?
217. What is backpressure and how do you handle it?
218. When should you prefer sync.Pool?
219. How do you manage high concurrency with low resource usage?
220. How do you monitor a Go application in production?

### 🟡 **221–240: Files, OS, and System Programming**
221. How do you read a file line by line in Go?
222. How do you write large files efficiently?
223. How do you watch file system changes in Go?
224. How to get file metadata like size, mod time?
225. How do you work with CSV files in Go?
226. How do you compress and decompress files in Go?
227. How do you execute shell commands from Go?
228. What is the `os/exec` package used for?
229. How do you set environment variables in Go?
230. How to create and manage temp files/directories?
231. How do you handle signals like SIGINT in Go?
232. How do you gracefully shut down a CLI app?
233. What are file descriptors and how does Go manage them?
234. How to handle large file uploads and streaming?
235. How do you access OS-specific syscalls in Go?
236. How do you implement a simple CLI tool in Go?
237. How do you build cross-platform binaries in Go?
238. What is syscall vs os vs exec package difference?
239. How do you write to logs with rotation?
240. What is the use of `ioutil` and its deprecation?

### 🔵 **241–260: Microservices, gRPC, and Communication**
241. What is gRPC and how is it used with Go?
242. How do you define Protobuf messages for Go?
243. What are the benefits of gRPC over REST?
244. How do you implement unary and streaming RPC in Go?
245. What is the difference between gRPC and HTTP/2?
246. How do you add authentication in gRPC services?
247. How do you handle timeouts and retries in gRPC?
248. How do you secure gRPC communication?
249. How do microservices communicate securely in Go?
250. What are message queues and how to use them in Go?
251. How to use NATS or Kafka in Go?
252. What are sagas and how would you implement them in Go?
253. How would you trace requests across services?
254. What is service discovery and how do you handle it?
255. How do you implement rate limiting across services?
256. What is the role of API gateway in microservices?
257. How do you use OpenTelemetry with Go?
258. How do you log correlation IDs between services?
259. How would you handle distributed transactions in Go?
260. How to deal with partial failures in distributed systems?

### 🟣 **261–280: Security and Best Practices**
261. How do you prevent injection attacks in Go?
262. What are Go's common security vulnerabilities?
263. How do you hash passwords securely in Go?
264. How to use `bcrypt` in Go?
265. How do you validate input in Go APIs?
266. How do you implement JWT authentication?
267. How do you prevent race conditions in Go?
268. What is CSRF and how to mitigate it?
269. How to use HTTPS in Go servers?
270. How do you sign and verify data in Go?
271. What are best practices for handling secrets in Go?
272. How do you handle OAuth2 flows in Go?
273. How do you restrict file uploads (size/type)?
274. How do you set up CORS properly in Go?
275. How do you scan Go code for vulnerabilities?
276. What is the Go ecosystem for SAST tools?
277. How to handle brute force protection in APIs?
278. How to secure communication between microservices?
279. What is the use of `context.Context` in secure APIs?
280. What is certificate pinning and can it be used in Go?

### 🔴 **281–300: Testing Strategy, CI/CD, Observability**
281. What are test doubles and how are they used in Go?
282. How do you structure unit vs integration tests?
283. What are flaky tests and how do you identify them?
284. How do you write deterministic tests for concurrency?
285. How do you test RESTful APIs in Go?
286. How do you mock HTTP calls?
287. What is Golden Testing in Go?
288. How do you run tests in parallel?
289. How do you mock time-dependent code?
290. How do you simulate DB failures in tests?
291. How do you use GitHub Actions to test Go apps?
292. What is the structure of a Makefile for Go?
293. How to build and test Go code in Docker?
294. What CI tools are commonly used for Go projects?
295. What are the benefits of go:embed for test fixtures?
296. How do you generate coverage reports in HTML?
297. How to collect logs and metrics from Go services?
298. What is structured logging in Go?
299. What are common logging libraries in Go?
300. How do you aggregate and search logs across services?

### 🟢 **301–320: Go Internals and Runtime**
301. How does the Go scheduler work?
302. What is M:N scheduling in Golang?
303. How does the Go garbage collector work?
304. What are STW (stop-the-world) events in GC?
305. How are goroutines implemented under the hood?
306. How does stack growth work in Go?
307. What is the difference between blocking and non-blocking channels internally?
308. What is a GOMAXPROCS and how does it affect execution?
309. How does Go manage memory fragmentation?
310. How are maps implemented internally in Go?
311. How does slice backing array reallocation work?
312. What is the zero value concept in Go?
313. How does Go avoid data races with its memory model?
314. What is escape analysis and how can you visualize it?
315. How are method sets determined in Go?
316. What is the difference between pointer receiver and value receiver at runtime?
317. How does Go handle panics internally?
318. How is reflection implemented in Go?
319. What is type identity in Go?
320. How are interface values represented in memory?

### 🟡 **321–340: DevOps, Docker, and Deployment**
321. How do you containerize a Go application?
322. What is a multi-stage Docker build and how does it help with Go?
323. How do you reduce the size of a Go Docker image?
324. How do you handle secrets in Go apps deployed via Docker?
325. How do you use environment variables in Go?
326. How do you compile a static Go binary for Alpine Linux?
327. What is `scratch` image in Docker and why is it used with Go?
328. How do you manage config files in Go across environments?
329. How do you build Go binaries for different OS/arch?
330. How do you use GoReleaser?
331. What is a Docker healthcheck for a Go app?
332. How do you log container stdout/stderr from Go?
333. How do you set up autoscaling for Go services?
334. How would you containerize a gRPC Go service?
335. How to deploy Go microservices in Kubernetes?
336. How do you write Helm charts for a Go app?
337. How do you monitor a Go service in production?
338. How do you use Prometheus with a Go app?
339. How do you enable structured logging in production?
340. How do you handle log rotation in containerized Go apps?

### 🔵 **341–360: Streaming, Messaging, and Asynchronous Processing**
341. How do you consume messages from Kafka in Go?
342. How do you publish messages to a RabbitMQ topic?
343. What is the idiomatic way to implement a message handler in Go?
344. How would you implement a worker pool pattern?
345. How do you use the `context` package for cancellation in streaming apps?
346. How do you retry failed messages in Go?
347. What is dead-letter queue and how do you use it?
348. How do you handle idempotency in message consumers?
349. How do you implement exponential backoff in Go?
350. How do you stream logs to a file/socket in real-time?
351. How do you work with WebSockets in Go?
352. How do you handle bi-directional streaming in gRPC?
353. What is Server-Sent Events and how is it done in Go?
354. How do you manage fan-in/fan-out channel patterns?
355. How would you implement throttling on async tasks?
356. How do you avoid data races when consuming messages?
357. How would you implement a message queue from scratch in Go?
358. How do you implement ordered message processing in Go?
359. How do you handle large stream ingestion (100K+ msgs/sec)?
360. How do you persist in-flight streaming data?

### 🟣 **361–380: Cloud-Native and Distributed Systems in Go**
361. How do you build a cloud-agnostic app in Go?
362. How do you use Go SDKs with AWS (S3, Lambda)?
363. How do you upload a file to S3 using Go?
364. How do you create a Pub/Sub system using Go and GCP?
365. How would you implement cloud-native config loading?
366. What is the role of service meshes with Go apps?
367. How do you secure service-to-service communication in Go?
368. How do you implement service registration and discovery?
369. How do you manage retries and circuit breakers in Go?
370. How would you use etcd/Consul with Go for KV storage?
371. What is leader election and how can you implement it in Go?
372. How do you build a distributed lock in Go?
373. How would you implement a distributed queue in Go?
374. How do you handle consistency in distributed Go systems?
375. How do you monitor and trace distributed Go systems?
376. How do you implement eventual consistency in Go?
377. How do you replicate state in distributed Go apps?
378. How do you detect and handle split-brain scenarios?
379. How do you implement quorum reads/writes in Go?
380. How would you build a simple distributed cache in Go?

### 🔴 **381–400: Go in Real-World Projects & Architecture**
381. How do you handle config versioning in Go projects?
382. How do you organize API versioning in Go apps?
383. How do you validate struct fields with custom rules?
384. How do you cache API responses in Go?
385. How do you serve files over HTTP with conditional GET?
386. How do you apply SOLID principles in Go?
387. How do you prevent breaking changes in shared Go modules?
388. What is the difference between horizontal and vertical scaling in Go services?
389. How do you support internationalization in Go?
390. How do you write a Go SDK for third-party APIs?
391. How do you manage request IDs and trace IDs?
392. How do you implement audit logging in Go?
393. How would you version a binary CLI in Go?
394. How do you ensure backward compatibility in Go libraries?
395. How do you handle soft deletes in Go models?
396. How do you refactor a large legacy Go codebase?
397. How do you maintain a mono-repo with multiple Go modules?
398. How would you go about building a plugin system in Go?
399. How do you document Go APIs automatically?
400. How do you track tech debt and enforce code quality in large Go teams?

### 🔵 **401–420: Networking and Low-Level Programming**
401. How do you create a TCP server in Go?
402. How do you create a UDP client in Go?
403. What is the difference between `net.Listen` and `net.Dial`?
404. How do you manage TCP connection pools?
405. How would you implement a custom HTTP transport?
406. How do you read raw packets using `gopacket`?
407. What is a connection hijack in `net/http` and how is it done?
408. How to implement a proxy server in Go?
409. How would you create an HTTP2 server from scratch in Go?
410. How does Go handle connection reuse (keep-alive)?
411. How do you set timeouts on sockets in Go?
412. What is the difference between `net/http` and `fasthttp`?
413. How do you throttle network traffic in Go?
414. How would you analyze network latency in Go?
415. How would you implement WebRTC or peer-to-peer comms?
416. How do you simulate a slow network in integration tests?
417. What’s the difference between connection pooling and multiplexing?
418. How do you verify DNS lookups in Go?
419. How do you use HTTP pipelining in Go?
420. How do you implement NAT traversal in Go?

### 🟣 **421–440: Error Handling & Observability**
421. How do you create custom error types in Go?
422. How does Go 1.20+ `errors.Join` and `errors.Is` work?
423. How do you implement error wrapping and unwrapping?
424. What are best practices for error categorization?
425. How do you handle critical vs recoverable errors?
426. How do you recover from panics in goroutines?
427. How to capture stack traces on error?
428. How do you notify Sentry/Bugsnag from Go?
429. How do you do structured error reporting in Go?
430. How do you correlate logs, errors, and traces together?
431. How would you add distributed tracing to an existing Go service?
432. What are tags, attributes, and spans in tracing?
433. What is a traceparent header?
434. How do you send custom metrics to Prometheus?
435. What is RED metrics model and how do you apply it?
436. How do you expose application health and readiness probes?
437. What’s the difference between logs, metrics, and traces?
438. How do you benchmark error impact on performance?
439. What’s the tradeoff between verbose and silent error handling?
440. How would you enforce observability in a Go microservice?

### 🟢 **441–460: CLI Tools, Automation, and Scripting**
441. How do you build an interactive CLI in Go?
442. What libraries do you use for command-line tools in Go?
443. How do you parse flags and config in CLI?
444. How do you implement bash autocompletion for Go CLI?
445. How would you use `cobra` to build a nested command CLI?
446. How do you manage color and styling in terminal output?
447. How would you stream CLI output like `tail -f`?
448. How do you handle secrets securely in a CLI?
449. How do you bundle a CLI as a standalone binary?
450. How would you version and release CLI with GitHub Actions?
451. How do you schedule a Go CLI tool with cron?
452. How do you use Go as a scripting language?
453. How do you embed templates in your Go CLI tool?
454. How would you create a system daemon in Go?
455. What are good patterns for CLI testing?
456. How do you store and manage CLI state/config files?
457. How do you secure a CLI for local system access?
458. How do you test CLI tools across multiple OS in CI?
459. How do you expose analytics and usage for a CLI?
460. How would you build a CLI wrapper for REST APIs?

### 🔴 **461–480: AI, Machine Learning & Data Processing in Go**
461. How do you use TensorFlow or ONNX models in Go?
462. What is `gorgonia` and when would you use it?
463. How do you implement cosine similarity in Go?
464. How would you stream CSV → transform → JSON using pipelines?
465. How do you process large datasets using goroutines?
466. How do you implement TF-IDF in Go?
467. How do you parse and tokenize text in Go?
468. How would you embed a local LLM into a Go app?
469. How do you integrate OpenAI API in Go?
470. How do you do prompt engineering for AI from Go?
471. How do you use a local vector database with Go?
472. How would you implement semantic search using Go?
473. How would you extract entities using regex or AI?
474. How do you manage model input/output formats in Go?
475. How would you create a chatbot backend with Go?
476. How do you build a recommendation engine with Go?
477. How would you integrate LangChain-like logic in Go?
478. How would you cache AI model outputs in Go?
479. What is the role of concurrency in AI inference in Go?
480. How do you monitor and scale AI pipelines in Go?

### 🟡 **481–500: WebAssembly, Blockchain, and Experimental Go**
481. What is WebAssembly and how can Go compile to WASM?
482. How do you share memory between JS and Go in WASM?
483. What is TinyGo and what are its limitations?
484. How do you write a smart contract simulator in Go?
485. What is Tendermint and how does Go power it?
486. How do you use `go-ethereum` to interact with smart contracts?
487. How do you parse blockchain data using Go?
488. How do you generate and verify ECDSA signatures in Go?
489. What is the role of Go in decentralized storage (IPFS)?
490. How would you implement a Merkle Tree in Go?
491. How do you handle base58 and hex encoding/decoding?
492. How do you write a deterministic VM interpreter in Go?
493. How do you simulate a P2P network in Go?
494. How do you create a lightweight Go runtime for edge computing?
495. How would you handle offline-first apps in Go?
496. What is the future of `Generics` in Go (beyond v1.22)?
497. What is fuzz testing and how do you use it in Go?
498. What is the `any` type in Go and how is it different from `interface{}`?
499. What is the latest experimental feature in Go and why is it important?
500. How do you contribute to the Go open-source project?

### 🔐 **501–520: Security in Golang**
501. How do you prevent SQL injection in Go?
502. How do you securely store user passwords in Go?
503. How do you implement OAuth 2.0 in Go?
504. What is CSRF and how do you prevent it in Go web apps?
505. How do you use JWT securely in a Go backend?
506. How do you validate and sanitize user input in Go?
507. How do you set secure cookies in Go?
508. How do you avoid path traversal vulnerabilities?
509. How do you prevent XSS in Go HTML templates?
510. How would you encrypt sensitive fields before storing in DB?
511. How do you securely generate random strings or tokens?
512. How do you verify digital signatures in Go?
513. What are best practices for TLS config in Go HTTP servers?
514. How do you implement rate limiting in Go to avoid DDoS?
515. How do you handle secrets in Go apps (Vault, env, etc.)?
516. How do you perform mutual TLS authentication in Go?
517. What is the difference between `crypto/rand` and `math/rand`?
518. How do you prevent replay attacks using Go?
519. How do you build a secure authentication system in Go?
520. How do you scan Go projects for vulnerabilities?

### 🚀 **521–540: Performance Optimization**
521. How do you benchmark Go code using `testing.B`?
522. What tools can you use to profile a Go application?
523. How does memory allocation affect Go performance?
524. How do you detect and fix memory leaks in Go?
525. How do you avoid unnecessary allocations in hot paths?
526. What is escape analysis and how does it impact performance?
527. How do you use `pprof` to trace CPU usage?
528. How do you optimize slice operations for speed?
529. What is object pooling and how is it implemented in Go?
530. How does GC tuning affect latency in Go services?
531. How do you measure and reduce goroutine contention?
532. What is lock contention and how to identify it in Go?
533. How do you batch DB operations for better throughput?
534. How would you profile goroutine leaks?
535. What are the downsides of excessive goroutines?
536. How would you measure and fix cold starts in Go Lambdas?
537. How do you decide between a map vs slice for performance?
538. How would you write a memory-efficient parser in Go?
539. How do you use channels efficiently under heavy load?
540. When should you use sync.Pool?

### 🧪 **541–560: Testing in Go**
541. How do you write table-driven tests in Go?
542. What is the difference between `t.Fatal` and `t.Errorf`?
543. How do you use `go test -cover` to check coverage?
544. How do you mock a database in Go tests?
545. How do you unit test HTTP handlers?
546. What is testable design and how does Go encourage it?
547. How do you use interfaces to improve testability?
548. How do you write tests for concurrent code in Go?
549. What is the `httptest` package and how is it used?
550. How do you mock time in tests?
551. How do you perform integration testing in Go?
552. How do you use `testify/mock` for mocking dependencies?
553. How do you run subtests and benchmarks?
554. How do you test panic recovery?
555. How do you generate test data using faker or random data?
556. What is golden file testing and when is it useful?
557. How do you automate test workflows with `go generate`?
558. How do you test CLI apps built with Cobra?
559. What is fuzz testing and how do you do it in Go?
560. How do you organize test files and test suites?

### 🧩 **561–580: API Design, REST/gRPC & Data Models**
561. How do you define a RESTful API in Go using Gin or Echo?
562. How do you version a REST API?
563. How do you handle validation of API payloads?
564. How do you return proper status codes from handlers?
565. How do you implement middleware in a Go web API?
566. How do you handle pagination in Go APIs?
567. What’s the difference between `json.Unmarshal` vs `Decode`?
568. How do you define a gRPC service in Go?
569. How do you handle gRPC errors and return codes?
570. How do you secure a gRPC service in Go?
571. How do you do field-level validation in proto definitions?
572. How do you log incoming requests/responses in a Go API?
573. How do you handle file uploads/downloads in APIs?
574. What is OpenAPI/Swagger and how do you generate docs in Go?
575. How do you serve static files securely in Go?
576. How do you implement a proxy API gateway in Go?
577. How do you generate Go code from `.proto` files?
578. How do you integrate gRPC with REST (gRPC-Gateway)?
579. How do you implement idempotency in APIs?
580. What is a contract-first API development approach?

### 🧠 **581–600: Design Patterns, Architecture & Real-World Scenarios**
581. How do you implement the Factory pattern in Go?
582. How do you use the Strategy pattern in Go?
583. What is the Singleton pattern and how is it safely used in Go?
584. How do you write a middleware chain in Go?
585. How do you use interfaces to decouple layers?
586. How do you implement the Observer pattern using channels?
587. What is the repository pattern and when do you use it?
588. How would you create a CQRS architecture in Go?
589. How do you design a plug-in architecture in Go?
590. What is a “clean architecture” in Go projects?
591. How do you structure a multi-module Go project?
592. How do you decouple business logic from transport layers?
593. How would you implement retryable jobs in Go?
594. How would you design a billing system in Go?
595. How would you scale a notification system written in Go?
596. How do you build a real-time leaderboard in Go?
597. How would you implement transactional emails in Go?
598. How do you model money and currencies in Go?
599. How do you do dependency injection in Go?
600. How do you create a rule engine in Go?

### 🔸 **601–620: Advanced Concurrency Patterns**
601. How do you implement a fan-in pattern in Go?
602. How do you implement a fan-out pattern in Go?
603. How do you prevent goroutine leaks in producer-consumer patterns?
604. How would you create a semaphore in Go?
605. What’s the difference between sync.WaitGroup and sync.Cond?
606. How do you implement a pub-sub model in Go?
607. How do you use a context to timeout multiple goroutines?
608. How do you build a rate-limiting queue with channels?
609. What is a worker pool, and how do you implement it?
610. How do you handle backpressure in channel-based designs?
611. How do you gracefully shut down workers?
612. How do you use sync.Cond for event signaling?
613. How do you prioritize tasks in concurrent processing?
614. How do you avoid starvation in goroutines?
615. How do you detect race conditions without `-race` flag?
616. How do you trace execution flow in concurrent systems?
617. How do you implement exponential backoff with retries in goroutines?
618. How do you structure long-running daemons with concurrency?
619. How would you implement circuit breakers in Go?
620. How do you handle concurrent map access with minimal locking?

### 🟤 **621–640: Event-Driven, Pub/Sub & Messaging**
621. How do you publish and consume events using NATS in Go?
622. How do you use Apache Kafka in Go with `sarama`?
623. What are the trade-offs between RabbitMQ and Kafka in Go apps?
624. How do you manage message acknowledgements in Go consumers?
625. How do you handle message deduplication in Go?
626. How do you implement a retry queue for failed messages?
627. How do you batch message processing efficiently in Go?
628. How do you use Google Pub/Sub with Go?
629. How do you persist event logs for replay in Go?
630. How do you ensure exactly-once delivery in Go message systems?
631. How do you create a lightweight in-memory pub-sub system?
632. How do you handle DLQs (Dead Letter Queues) in Go?
633. How do you create idempotent message consumers in Go?
634. How do you enforce ordering of messages?
635. How do you use channels as message queues?
636. How do you handle push vs pull consumers?
637. How do you deal with large payloads in a messaging system?
638. How do you build an event sourcing system in Go?
639. How would you test message-driven systems?
640. What’s the role of event schemas in Go-based systems?

### 🟢 **641–660: Go for DevOps & Infrastructure**
641. How do you create a custom Kubernetes operator in Go?
642. How do you write a Helm plugin in Go?
643. How do you use Go for infrastructure automation?
644. How do you write a CLI for managing AWS/GCP resources?
645. How do you use Go to write Terraform providers?
646. How do you build a dynamic inventory script in Go for Ansible?
647. How do you parse and generate YAML in Go?
648. How do you interact with Docker API in Go?
649. How do you manage Kubernetes CRDs in Go?
650. How do you write Go code to scale deployments in K8s?
651. How do you tail logs from containers using Go?
652. How do you manage service discovery in Go apps?
653. How do you build a Kubernetes admission controller in Go?
654. How do you build a metrics exporter for Prometheus in Go?
655. How do you set up health checks for a Go microservice?
656. How do you build a custom load balancer in Go?
657. How do you implement graceful shutdown with Kubernetes SIGTERM?
658. How do you use Go with Envoy/Consul service mesh?
659. How do you configure Go apps for 12-factor principles?
660. How do you use Go for cloud automation scripts?

### 🟣 **661–680: Caching & Storage Systems**
661. How do you cache database query results in Go?
662. How do you use Redis with Go for distributed caching?
663. How do you implement LRU cache in Go?
664. How do you ensure cache invalidation on data update?
665. How do you handle stale reads in Go apps with caching?
666. How do you implement a write-through cache in Go?
667. How do you handle concurrency in in-memory caches?
668. How do you use bloom filters in Go?
669. How do you build a TTL-based memory cache?
670. How do you use memcached in Go?
671. How do you store large binary blobs in Go?
672. How do you build an append-only log file storage in Go?
673. How do you use BoltDB or BadgerDB in Go?
674. How do you structure a file-based key-value store in Go?
675. How do you handle distributed caching with Go?
676. How do you monitor cache hit/miss ratios in Go?
677. How do you use consistent hashing for distributed caching?
678. How do you build a cache warming strategy in Go?
679. How do you use S3-compatible storage APIs in Go?
680. How do you implement local persistent disk caching?

### 🔴 **681–700: Real-Time Systems, IoT, and Edge Computing**
681. How do you build a real-time chat server in Go?
682. How do you implement WebSockets in Go?
683. How do you ensure order of events in real-time systems?
684. How do you handle high concurrency in WebSocket servers?
685. How do you implement presence tracking in Go (like online users)?
686. How do you reduce latency in real-time systems?
687. How do you build a real-time dashboard backend in Go?
688. How do you handle message fan-out for WebSocket clients?
689. How do you design a real-time bidding system in Go?
690. How do you throttle real-time updates?
691. How do you buffer real-time data safely?
692. How do you build a publish-subscribe engine for WebSockets?
693. How do you sync real-time state between browser and Go backend?
694. How do you implement real-time location tracking?
695. How do you use Go in resource-constrained IoT devices?
696. How do you collect telemetry data from IoT devices?
697. How do you compress and transmit data from edge devices in Go?
698. How do you implement OTA (over-the-air) updates using Go?
699. How do you design protocols for edge-device communication?
700. How do you build secure, low-latency edge APIs in Go?

### ⚙️ **701–720: Go Internals & Runtime**
701. How does the Go scheduler work internally?
702. What are GOMAXPROCS and how does it affect performance?
703. What’s the internal structure of a goroutine?
704. How does garbage collection work in Go’s runtime?
705. What are safepoints in the Go runtime?
706. What is cooperative scheduling in Go?
707. What are the stages of Go's garbage collector?
708. What is the role of the `runtime` package?
709. How does Go handle stack traces?
710. How does the Go runtime manage memory allocation?
711. What’s the difference between a green thread and a goroutine?
712. What are finalizers in Go and how do they work?
713. What is the role of `goexit` internally?
714. How does Go avoid stop-the-world pauses?
715. How does memory fragmentation affect Go programs?
716. What’s the meaning of “non-preemptible” code in Go?
717. What are M:N scheduling models and how does Go implement it?
718. How does Go detect deadlocks at runtime?
719. What are the internal states of a goroutine?

### 📡 **721–740: Network & Protocol-Level Programming**
720. How do you create a custom TCP server in Go?
721. How do you parse HTTP headers manually in Go?
722. How do you handle fragmented UDP packets in Go?
723. How do you implement a custom binary protocol in Go?
724. How do you parse and encode protobufs manually?
725. How do you build a TCP proxy in Go?
726. How do you implement a reverse proxy in Go?
727. How do you sniff packets using Go?
728. How do you build a SOCKS5 proxy in Go?
729. How do you write a raw socket listener in Go?
730. How do you implement an HTTP client with timeout handling?
731. How do you use netpoll in high-performance Go networking?
732. How do you build a DNS resolver in Go?
733. How do you manage connection pooling in network services?
734. How do you detect dropped connections in TCP?
735. What’s the difference between persistent and non-persistent HTTP in Go?
736. How do you write a TLS server in Go from scratch?
737. How do you implement rate limiting per IP in a TCP server?
738. How do you use Go to test API latency?
739. How do you monitor and log TCP connections?

### 📦 **741–760: Error Handling & Observability**
740. How do you implement a custom error type in Go?
741. How do you wrap errors in Go?
742. What is `errors.Is()` and `errors.As()` used for?
743. How do you categorize errors in large Go applications?
744. How do you log structured errors in Go?
745. How do you use sentry/bugsnag with Go?
746. How do you implement centralized error logging?
747. What is the role of stack traces in debugging Go apps?
748. How do you implement panic recovery with context?
749. How do you differentiate retryable vs fatal errors?
750. How do you expose Prometheus metrics in Go?
751. How do you set up OpenTelemetry in Go?
752. How do you trace gRPC requests in Go?
753. How do you record and export application traces?
754. How do you handle slow endpoints in production Go apps?
755. How do you add custom labels/tags to logs?
756. How do you redact sensitive data in logs?
757. How do you detect memory leaks using Go tools?
758. How do you instrument performance counters in Go?
759. How do you implement a tracing middleware?

### 🔄 **761–780: Streaming, Batching & Data Pipelines**
760. How do you process large CSV files using streaming?
761. How do you implement backpressure in a data stream?
762. How do you connect Go with Apache Kafka for streaming?
763. How do you build an ETL pipeline in Go?
764. How do you handle JSONL (JSON Lines) in real-time streams?
765. How do you split and parallelize stream processing?
766. How do you deal with schema evolution in streaming data?
767. How do you throttle input data rate?
768. How do you aggregate streaming metrics?
769. How do you implement checkpointing in Go pipelines?
770. How do you persist intermediate results in streams?
771. How do you implement a rolling window average?
772. How do you batch messages for optimized DB writes?
773. How do you stream process financial transactions in Go?
774. How do you integrate with Apache Pulsar in Go?
775. How do you compress/decompress streaming data?
776. How do you handle late data in streaming?
777. How do you fan-out a stream to multiple destinations?
778. How do you filter events in a stream dynamically?
779. How do you manage ordered processing in Kafka consumers?

### 🧪 **781–800: Go Tooling, CI/CD & Developer Experience**
780. How do you create custom `go generate` commands?
781. How do you build a multi-binary Go project?
782. How do you configure GoReleaser for automated builds?
783. How do you sign binaries in Go before release?
784. How do you use `go vet` to detect issues?
785. How do you manage environment-specific builds in Go?
786. How do you use `build tags` in Go?
787. How do you profile CPU/memory usage in CI pipelines?
788. How do you automate `go test` and coverage in GitHub Actions?
789. How do you write a custom Go linter?
790. How do you automate versioning and changelogs in Go projects?
791. How do you use `go:embed` for bundling files?
792. How do you validate Go module versions in a monorepo?
793. How do you containerize a Go application for fast startup?
794. How do you enable live reloading for Go dev servers?
795. How do you run multiple Go services locally with Docker Compose?
796. How do you handle secrets securely in Go CI pipelines?
797. How do you cross-compile Go binaries for ARM and Linux?
798. How do you build Go CLIs that auto-complete in Bash and Zsh?
799. How do you keep your Go codebase idiomatic and consistent?

### 🛡️ **801–820: Security & Authentication**
800. How do you implement HMAC-based authentication in Go?
801. How do you use JWT securely in Go APIs?
802. How do you manage CSRF protection in a Go web app?
803. How do you handle XSS prevention in Go templates?
804. How do you implement OAuth 2.0 flows in Go?
805. How do you encrypt/decrypt sensitive data in Go?
806. What’s the use of `crypto/rand` vs `math/rand`?
807. How do you manage TLS certs in Go servers?
808. How do you validate tokens in Go microservices?
809. How do you securely store API keys in Go apps?
810. How do you create and validate secure cookies?
811. How do you implement role-based access control in Go?
812. How do you generate a secure random token in Go?
813. How do you prevent replay attacks with Go?
814. How do you audit Go applications for security issues?
815. How do you apply security headers in Go HTTP servers?
816. How do you secure gRPC endpoints in Go?

### 🧪 **821–840: Testing & Quality**
817. How do you mock HTTP clients in Go tests?
818. How do you achieve high test coverage in Go?
819. How do you test race conditions in Go?
820. How do you benchmark functions in Go?
821. How do you structure tests for a large Go codebase?
822. How do you use interfaces for testability?
823. How do you test panics in Go?
824. How do you generate test data in Go?
825. How do you test concurrent code in Go?
826. How do you mock database interactions in Go?
827. How do you test middleware in a Go web app?
828. How do you use `httptest.Server`?
829. How do you run parallel tests in Go?
830. How do you test CLI apps in Go?
831. How do you perform fuzz testing in Go?
832. How do you simulate network failures in tests?
833. How do you write integration tests with Docker?
834. How do you test gRPC services in Go?
835. How do you set up CI pipelines for testing Go apps?

### 🏎️ **841–860: Performance Optimization**
836. How do you avoid unnecessary allocations?
837. How do you reduce GC pressure in Go apps?
838. How do you profile heap allocations?
839. How do you use escape analysis to optimize code?
840. How do you optimize JSON marshaling in Go?
841. How do you write cache-friendly code in Go?
842. How do you improve startup time of a Go app?
843. How do you reduce lock contention in Go?
844. How do you identify goroutine leaks?
845. How do you minimize context switches?
846. How do you use sync.Pool effectively?
847. How do you optimize string concatenation?
848. How do you use benchmarking to choose better algorithms?
849. How do you eliminate redundant computations?
850. How do you spot unnecessary interface conversions?
851. How do you improve performance of I/O-heavy apps?
852. How do you handle large slices without GC spikes?
853. How do you reduce reflection usage in Go?
854. How do you apply zero-copy techniques?

### 🧠 **861–880: Go Compiler & Language Theory**
855. How do you build a custom Go compiler plugin?
856. What is SSA (Static Single Assignment) form in Go?
857. How does Go handle type inference?
858. What is escape analysis in Go?
859. How does inlining affect performance in Go?
860. What are build constraints and how do they work?
861. How does `defer` work at the bytecode level?
862. What is the Go frontend written in?
863. How are interfaces implemented in memory?
864. What are method sets and how do they affect interfaces?
865. How do you implement AST manipulation in Go?
866. What is the Go toolchain pipeline from source to binary?
867. How are function closures handled by the Go compiler?
868. What is link-time optimization in Go?
869. How does cgo interact with Go's runtime?
870. What are zero-sized types and how are they used?
871. How does type aliasing differ from type definition?
872. How does Go avoid null pointer dereferencing?
873. What’s the role of go/types package?
874. How does Go manage ABI stability?

### 🧰 **881–900: Refactoring, CLI, WebAssembly & Design**
875. How do you refactor large Go codebases safely?
876. How do you break a monolith Go app into microservices?
877. How do you improve code readability in Go?
878. How do you organize domain-driven projects in Go?
879. How do you handle circular dependencies?
880. How do you structure reusable Go modules?
881. How do you build CLI apps with Cobra?
882. How do you add auto-completion to CLI tools?
883. How do you handle subcommands in CLI tools?
884. How do you package Go binaries for multiple platforms?
885. How do you write a Wasm frontend in Go?
886. How do you expose Go functions to JS using Wasm?
887. How do you reduce Wasm binary size?
888. How do you interact with DOM from Go Wasm?
889. How do you debug Go WebAssembly apps?
890. How do you build a WebAssembly module loader?
891. How do you manage state in Go WebAssembly apps?
892. How do you integrate Go Wasm with JS promises?
893. How do you decide between Go CLI and REST tool?
894. How do you document CLI help and usage info?

### 🧠 **901–920: AI, ML & Generative Use Cases in Go**
895. How do you call an OpenAI API using Go?
896. How do you stream ChatGPT responses in Go?
897. How do you build a Telegram AI bot in Go?
898. How do you integrate Go with HuggingFace models?
899. How do you use TensorFlow models in Go?
900. How do you build a Go app that uses image recognition?
901. How do you generate code snippets using LLMs in Go?
902. How do you do prompt templating in Go?
903. How do you build a LangChain-style pipeline in Go?
904. How do you fine-tune prompts using Go templates?
905. How do you handle concurrent API calls to LLMs?
906. How do you track token usage in LLM APIs from Go?
907. How do you stream generation results to a web frontend in Go?
908. How do you handle OpenAI rate limits in Go apps?
909. How do you generate embeddings and store in Go?
910. How do you integrate Pinecone or Weaviate with Go?
911. How do you manage vector searches using Go?
912. How do you build a question-answering bot using Go?
913. How do you evaluate AI responses using Go logic?
914. How do you serialize LLM chat history in Go?

### 💾 **921–940: Go + Databases (SQL, NoSQL, ORMs)**
915. How do you use database/sql in Go?
916. What are connection pools and how to manage them?
917. How do you write raw queries using `sqlx`?
918. How do you use GORM with PostgreSQL?
919. How do you handle transactions in Go?
920. How do you create database migrations in Go?
921. How do you use MongoDB with Go?
922. How do you store JSONB in PostgreSQL using Go?
923. How do you index and search in Elasticsearch using Go?
924. How do you use Redis with Go for caching?
925. How do you use prepared statements in Go?
926. How do you prevent N+1 queries using Go ORM?
927. How do you map complex nested objects from DB in Go?
928. How do you benchmark DB performance in Go?
929. How do you test DB queries with mocks?
930. How do you stream large query results in Go?
931. How do you use SQLite for embedded apps in Go?
932. How do you connect Go to Amazon RDS or Aurora?
933. How do you manage read replicas in Go?
934. How do you handle DB failovers in Go apps?

### 🌐 **941–960: REST APIs & gRPC Design**
935. How do you design versioned REST APIs in Go?
936. How do you add OpenAPI/Swagger support in Go?
937. How do you handle graceful shutdown of API servers?
938. How do you write middleware for logging/auth?
939. How do you secure REST APIs using JWT?
940. How do you design a RESTful file upload service?
941. How do you handle CORS in a Go API?
942. How do you paginate API responses?
943. How do you implement rate-limiting on APIs?
944. How do you handle multipart/form-data in Go?
945. How do you expose metrics from a Go API?
946. How do you mock gRPC services in tests?
947. How do you set up gRPC with reflection?
948. How do you stream data over gRPC?
949. How do you version gRPC APIs?
950. How do you enforce contracts with protobuf validators?
951. How do you convert REST to gRPC clients?
952. How do you monitor gRPC health checks?
953. How do you build a gRPC gateway in Go?
954. How do you throttle gRPC traffic in Go?

### 🧵 **961–980: Concurrency Architecture & Design Patterns**
955. How do you architect a pub/sub system in Go?
956. How do you build a pipeline using goroutines?
957. What is the fan-in/fan-out pattern in Go?
958. How do you limit concurrency using semaphores?
959. How do you implement a worker pool?
960. How do you handle retries with backoff in goroutines?
961. What is the circuit breaker pattern in Go?
962. How do you implement message deduplication?
963. How do you synchronize shared state across goroutines?
964. How do you detect livelocks in Go?
965. How do you timeout long-running operations?
966. How do you use the actor model in Go?
967. How do you architect loosely coupled goroutines?
968. How do you design state machines in Go?
969. How do you throttle a job queue in Go?
970. How do you monitor goroutine health?
971. How do you track context propagation in goroutines?
972. How do you implement saga pattern in Go services?
973. How do you chain async jobs with error handling?
974. How do you log and trace concurrent tasks?

### ⚒️ **981–1000: Tooling, Maintenance & Real-world Scenarios**
975. How do you create internal packages in Go?
976. How do you enforce code standards using golangci-lint?
977. How do you write makefiles for Go projects?
978. How do you manage secrets using Vault in Go?
979. How do you deploy a Go app with Kubernetes?
980. How do you perform zero-downtime deployment in Go?
981. How do you refactor legacy Go code?
982. How do you organize large-scale Go monorepos?
983. How do you distribute Go binaries securely?
984. How do you maintain changelogs in Go projects?
985. How do you rollback failed Go releases?
986. How do you add performance regression testing?
987. How do you build CLI-based installers in Go?
988. How do you generate dashboards from Go metrics?
989. How do you monitor file system changes in Go?
990. How do you implement custom plugins in Go?
991. How do you keep Go dependencies up to date?
992. How do you audit Go packages for security issues?
993. How do you migrate Go modules across repos?
994. How do you conduct performance reviews for Go codebases?
995. How do you implement a concurrent token bucket rate limiter?
996. How do you implement the Saga Pattern for distributed transactions?
997. What is the Go Memory Model and the "Happens-Before" relationship?
998. How do you implement distributed tracing with context propagation?
999. How do you use the `slices` and `maps` packages (Go 1.21+)?
1000. What is Profile Guided Optimization (PGO) and how do you use it?

### 🆕 **1001–1020: Modern Go Features (v1.22–v1.24)**
1001. What is the loop variable scope change in Go 1.22?
1002. How do you iterate over integers in Go 1.22 (`for i := range n`)?
1003. How does `net/http.ServeMux` support wildcards and methods in Go 1.22?
1004. What is the new `math/rand/v2` package?
1005. What are Go Iterators (`range-over-func`) in Go 1.23?
1006. How do you use the `unique` package in Go 1.23?
1007. What improvements were made to `time.Timer` garbage collection in Go 1.23?
1008. What are generic type aliases in Go 1.24?
1009. How do you use the `go tool` directive in `go.mod` (Go 1.24)?
1010. What is `os.Root` and how does it improve file system isolation (Go 1.24)?
1011. How do you implement weak pointers in Go 1.24?
1012. What is the `omitzero` struct tag option?
1013. How do you use `testing.B.Loop` for benchmarks?
1014. How does Go 1.24 support FIPS 140-3 compliance?
1015. What are usage comparisons for `slices.Concat`?
1016. How do you use `runtime.AddCleanup` vs `SetFinalizer`?
1017. What are the new WASM export capabilities in Go?
1018. How do you debug using `go build -asan`?
1019. How do you manage tool dependencies without a `tools.go` file now?
1020. What is the anticipated "Flight Recorder" feature?

### 🧩 **1021–1045: Niche Patterns, Frameworks & Tricky Snippets**
1021. How do you implement the "Or-Done" channel pattern in Go?
1022. What is a "Tee-Channel" and how do you implement it?
1023. How do you implement a "Bridge-Channel" to consume a sequence of channels?
1024. What is the **Temporal.io** workflow engine and how does it use Go?
1025. How does **Temporal** ensure determinism in Go workflows?
1026. What is the **Ent** framework and how does it differ from GORM?
1027. How do you define graph-based schemas (Edges) in **Ent**?
1028. **Tricky Snippet**: What is the output of `fmt.Println(s)` if `s := []int{1,2,3}; append(s[:1], 4)` is called?
1029. **Tricky Snippet**: Why is `interface{}(*int(nil)) != nil` true?
1030. **Tricky Snippet**: What happens if you run `for k, v := range m` on the same map multiple times?
1031. **Tricky Snippet**: What happens when you close a nil channel?
1032. **Tricky Snippet**: Can you take the address of a map value (`&m["key"]`)? Why or why not?
1033. How do you use `golang.org/x/sync/errgroup` for error propagation?
1034. What is the "Function Options" pattern for constructor configuration?
1035. How do you use `singleflight` to prevent cache stampedes?
1036. What is `uber-go/automaxprocs` and why is it used in K8s?
1037. How do you implement "Circuit Breaker" using `sony/gobreaker` or similar?
1038. How do you use build tags to separate integration tests (`//go:build integration`)?
1039. What is the difference between `crypto/rand` and `math/rand/v2` in terms of security?
1040. How do you use `go-cmp` for comparing complex structs in tests?
1041. What is "Mutation Testing" and are there tools for it in Go?
1042. How do you handle "Dual-Writes" (DB + Message Queue) consistency?
1043. What is the "Outbox Pattern" and how to implement it in Go?
1044. How does Go's "semver" compatibility guarantee work for standard library?
1045. How do you use `gdb` or `delve` to debug a running process (attach)?

