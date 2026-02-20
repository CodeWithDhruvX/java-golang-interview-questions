# 🟢 **21–40: Communication**

### 21. Synchronous vs asynchronous communication?
"In synchronous communication, the client sends a request and blocks (waits) for a response from the server. An example is a standard HTTP REST call. If the Payment service is slow, the Order service is forced to wait, potentially causing a timeout.

In asynchronous communication, the client sends a message and immediately moves on without waiting for a reply. It usually involves a message broker like Kafka or RabbitMQ. The Order service drops an 'OrderCreated' event into a queue and immediately returns a 'Success' to the user. The Payment service picks it up whenever it's ready.

I prefer async communication between microservices because it drastically improves system resilience. If the Payment service goes down, the Order service can still accept new orders."

#### Indepth
Synchronous introduces temporal coupling—both services must be alive at the exact same moment. Asynchronous introduces temporal decoupling. However, asynchronous adds complexity: you need to handle message duplication, out-of-order delivery, and eventual consistency for the user (e.g., 'Payment Pending' UI states).

---

### 22. REST vs gRPC?
"REST is built on HTTP/1.1 (usually), using JSON as the data format. It uses standard HTTP verbs (GET, POST) and is universally understood. However, JSON serialization is relatively slow, and HTTP/1.1 has large headers.

gRPC is developed by Google and built on HTTP/2. It uses Protocol Buffers (Protobuf) instead of JSON. Protobuf is a strongly-typed, binary format. Because it's binary and HTTP/2 supports multiplexing, gRPC is insanely fast and has a much smaller payload size.

I use REST for public-facing APIs or external integrations because every language and browser understands it. But for internal, high-throughput communication between my backend microservices, I strongly prefer gRPC."

#### Indepth
gRPC relies on contract-first design. You write a `.proto` file defining the service and its messages, then compile it into native client/server stubs in Java, Go, Python, etc. This eliminates manual DTO mapping errors but makes debugging harder because you cannot easily read binary payloads in a network tab like you can with JSON.

---

### 23. What is idempotency?
"Idempotency is a property of an operation where applying it multiple times yields the same result as applying it just once.

In microservices, network connections drop frequently. If a client sends a 'Charge Credit Card' request and doesn't receive a response due to a brief network blip, they will automatically retry the request. If the endpoint is not idempotent, the user gets charged twice.

I always ensure critical endpoints (especially POST/PUT requests handling money or state changes) are idempotent."

#### Indepth
HTTP verbs have standard idempotency rules. GET, PUT, and DELETE are inherently idempotent by definition (deleting user ID 5 ten times results in user ID 5 remaining deleted). POST is notoriously non-idempotent because creating the same resource ten times creates ten distinct resources unless specifically constrained.

---

### 24. How to design idempotent APIs?
"To make a POST API idempotent, I require the client to generate a unique 'Idempotency Key' (usually a UUID) and send it in the request header (`Idempotency-Key: 12345`).

When my service receives the request, it first checks the database or a Redis cache: 'Have I seen this key before?'
If yes, I immediately return the cached response from the first successful attempt without re-executing the business logic.
If no, I process the request, save the result against the key in Redis, and return the response.

This guarantees that a user clicking 'Checkout' ten times rapidly only results in one actual purchase."

#### Indepth
Idempotency keys should ideally have a Time-To-Live (TTL). You don't need to store an idempotency key forever in Redis—usually 24 hours is enough to cover any immediate network retry windows. Stripe popularized this specific header-based approach to API idempotency.

---

### 25. What is request-response pattern?
"The request-response pattern is the most fundamental communication paradigm where a client sends a message to a server and waits for a specific reply. 

In a microservices architecture, this is typically implemented using HTTP REST calls or gRPC. Point A talks directly to Point B. 

I use this pattern when the client absolutely needs an immediate answer to proceed. For example, validating a user's login credentials must strictly be a synchronous request-response."

#### Indepth
This pattern creates tight temporal coupling. If the responding service is down or slow, the requesting service suffers equally. In high-load systems, synchronous request-response chains (A -> B -> C -> D) are major anti-patterns because latency compounds and a single failure cascades through the entire chain.

---

### 26. What is publish-subscribe pattern?
"The Publish-Subscribe (Pub/Sub) pattern relies on an asynchronous broker, like Apache Kafka or AWS SNS. 

A service 'publishes' a message to a topic (e.g., 'OrderCreated'), without knowing who is listening. Other services 'subscribe' to that topic. The broker broadcasts the message to all subscribers. 

I rely on Pub/Sub to decouple my services. The Order service doesn't need to know that the Email service, the Shipping service, and the Analytics service all need to be notified about a new order. It just fires the event and forgets."

#### Indepth
Pub/Sub architectures are highly scalable because adding a new subscriber (e.g., a new 'Fraud Detection' service) requires zero code changes to the publishing service. It completely reverses the dependency tree.

---

### 27. What is event-driven architecture?
"Event-Driven Architecture (EDA) is a design paradigm where microservices communicate primarily by producing and consuming 'events'—notifications that state has changed.

Instead of making synchronous REST API calls to command another service to do something ('Command-driven'), a service broadcasts an event ('Order Completed'). Any service that cares about an order completing listens to that event and reacts accordingly.

I use EDA to build highly scalable, resilient systems. Because services don't talk directly to each other, a massive traffic spike simply queues up events in the broker (like Kafka) instead of crashing backend services."

#### Indepth
EDA introduces significant complexity around tracing (how do you track a user's action across 5 asynchronous events?) and data consistency. Developers must design compensating transactions (Sagas) because you cannot simply wrap HTTP and Kafka calls in a single ACID database transaction.

---

### 28. What is eventual consistency?
"Eventual consistency is a theoretical guarantee that, given no new updates to a piece of data, all nodes (or microservices) will eventually return the last updated value.

In a microservice architecture using async messaging, data is usually not consistent instantly. If I place an order, the Order database updates immediately. But it might take 500 milliseconds for the event to reach the Inventory database to deduct stock. During that 500ms window, the system is 'inconsistent'.

I have to carefully design my UI to handle this—for example, showing 'Order Processing...' instead of an immediate definitive state."

#### Indepth
Eventual consistency trades real-time accuracy for massive horizontal scalability and high availability, leaning into the AP side of the CAP theorem. It assumes that momentary inconsistency is acceptable for business logic as long as state converges to accuracy quickly.

---

### 29. What is strong consistency?
"Strong consistency ensures that once a piece of data is updated, any subsequent read from any node will instantly reflect that update.

In a monolithic application with a single relational database, strong consistency is easy: wrapping queries in an ACID transaction guarantees it. In distributed microservices, achieving strong consistency requires complex protocols like Two-Phase Commit (2PC).

I actively avoid designing microservices that require strong consistency across service boundaries, because it forces synchronous blocking across the network, severely hurting performance and availability."

#### Indepth
If Service A and Service B must be strongly consistent, a failure in B means A cannot complete its transaction. This means overall system availability equals the multiplied availability of all participating services—a recipe for downtime. Financial ledger systems sometimes require it, but most domains can tolerate eventual consistency.

---

### 30. What are HTTP status codes and their use?
"HTTP status codes are standard responses given by web servers to indicate the result of a client's request. They are essential for REST APIs because they tell the caller exactly what happened without parsing a JSON body.

- **2xx (Success):** 200 OK (Generic success), 201 Created (After a successful POST).
- **4xx (Client Error):** 400 Bad Request (Invalid validation), 401 Unauthorized (Missing JWT), 403 Forbidden (Valid JWT, wrong role), 404 Not Found.
- **5xx (Server Error):** 500 Internal Server Error (Code crash), 503 Service Unavailable (Gateway timeout or offline).

I enforce strict adherence to these codes so client applications can handle routing logic efficiently."

#### Indepth
Proper use of status codes informs client-side retry logic. A client should never retry a 400 Bad Request (the payload is wrong, retrying won't fix it). A client should optionally retry a 503 HTTP status (the server might be restarting) or a 429 Too Many Requests (using an exponential backoff strategy).

---

### 31. What is correlation ID?
"A Correlation ID is a unique identifier (usually a UUID) attached to a request as it enters the system at the API Gateway.

Because a single user action might trigger a chain of calls across five different microservices, debugging an error based on standard logs is nearly impossible. I pass this Correlation ID in the HTTP headers (e.g., `X-Correlation-ID`) from service to service. 

When every microservice logs its actions, it includes this ID. I can then go into my centralized logging tool (like Kibana), search for that ID, and see the exact chronological flow of that specific request across the entire system."

#### Indepth
Spring Cloud Sleuth (or Micrometer Tracing in newer Spring versions) automates this entirely. It implements Distributed Tracing standards (like W3C Trace Context) by injecting and propagating `traceId` and `spanId` globally across MDC (Mapped Diagnostic Context) log formats automatically.

---

### 32. What is timeout?
"A timeout is a maximum time a service is willing to wait for a response from another service before abandoning the request.

If my Order service calls an external Payment API and doesn't specify a timeout, it might wait infinitely if the Payment server silently hangs. This ties up a thread. Eventually, all threads in the Order service get exhausted waiting, and the Order service crashes too.

I strictly configure connection timeouts and read timeouts on every `RestTemplate` or `WebClient` call to ensure my service fails fast and gracefully."

#### Indepth
Timeouts should be configured hierarchically. The API Gateway might have a 5-second timeout, internal service A might have a 3-second timeout, and the database call might have a 1-second timeout. If internal timeouts are larger than outer edge timeouts, resources are effectively wasted processing requests that the client has already abandoned.

---

### 33. What is retry mechanism?
"A retry mechanism automatically re-attempts a failed operation, anticipating that the failure was transient (like a momentary network blip or a brief database lock).

When a REST call fails due to a `503 Service Unavailable` or an `IOException`, I configure libraries like Spring Retry or Resilience4j to simply try the exact same call again after a short delay.

While it's a great quick-fix for network flakiness, I only configure retries for completely idempotent endpoints to avoid accidental duplicate data creation."

#### Indepth
Naive linear retries can cause self-inflicted Distributed Denial of Service (DDoS) attacks against your own infrastructure (the "Retry Storm" phenomenon). If a database struggles, and 500 microservice instances immediately retry at the same time, the database collapses completely.

---

### 34. What is backoff strategy?
"A backoff strategy dictates the delay between retry attempts. Instead of retrying immediately, it spaces out the requests.

The best approach is **Exponential Backoff with Jitter**. The first retry waits 1 second, the second waits 2 seconds, the third waits 4 seconds... reducing the pressure on the struggling downstream service. 'Jitter' adds a random millisecond variance (e.g., waiting 1.2s instead of exactly 1.0s).

I use Jitter to prevent 'thundering herds', where hundreds of retrying client instances accidentally synchronize their wait timers and hit the server at the exact same millisecond."

#### Indepth
Exponential Backoff with Jitter is the industry standard for cloud integrations. AWS SDKs, for example, implement this by default on all API calls.

---

### 35. What is circuit breaker pattern?
"The Circuit Breaker pattern prevents an application from repeatedly trying to execute an operation that's likely to fail. 

If my Inventory service goes completely offline, there's no point in the Order service sending 1,000 requests per second to it—they will all timeout and exhaust Order service threads. 

A library like Resilience4j tracks failures. If 50% of requests fail within 10 seconds, the circuit 'opens'. All subsequent calls immediately fail-fast, throwing an exception instantly without waiting for network timeouts. Periodically, the circuit lets one request 'half-open' to test if the Inventory service has recovered."

#### Indepth
Circuit breakers prevent catastrophic cascading failures across a microservice architecture. They also provide the failing downstream service essential "breathing room" to recover without being continuously bombarded with queued network requests.

---

### 36. What is bulkhead pattern?
"The Bulkhead pattern isolates resources so that a failure in one part of the system doesn't bring down the entire system, similar to waterproof compartments in a ship's hull.

If my microservice exposes both an intensive 'Report Generation' API and a fast 'Profile Checkout' API, I assign them separate thread pools. 

If the reporting database slows down and uses up all 50 threads in the reporting pool, the profile checkout pool remains completely unaffected. I love this pattern because it guarantees one faulty feature won't consume resources required by healthy features."

#### Indepth
Without bulkheads, standard application servers (like Tomcat) share a global pool of active worker threads (e.g., 200 max connections). A slow external API call in a minor, unimportant service can monopolize all 200 Tomcat threads, bringing a critical server to a complete halt.

---

### 37. What is rate limiting?
"Rate limiting restricts the number of requests a client can make to an API within a specific timeframe (e.g., 100 requests per minute).

I enforce this primarily at the API Gateway level to protect backend microservices from intentional brute-force attacks, DDoS attacks, or badly written client scripts looping endlessly.

If a client exceeds the limit, the gateway immediately returns an HTTP `429 Too Many Requests` status, shedding the load before it even touches the internal network."

#### Indepth
Standard algorithms include the Token Bucket, Leaky Bucket, and Fixed/Sliding Window counters. Redis is commonly utilized to store the distributed counters rapidly across multiple API Gateway instances, ensuring an attacker can't bypass the limit by hitting different load-balanced nodes.

---

### 38. What is throttling?
"Throttling is often used interchangeably with rate limiting, but generally, throttling refers to gracefully slowing down requests or reducing service quality dynamically, rather than outright blocking them.

For example, if my backend system detects extremely high CPU usage, I might proactively 'throttle' incoming traffic at the gateway, enforcing a lower queue rate or rejecting low-priority background API requests to keep the core user-facing flows alive.

It's a defensive posture implemented to ensure the survival of the critical system components under massive stress."

#### Indepth
APIs often send back custom headers like `X-RateLimit-Limit` and `X-RateLimit-Remaining` to warn clients when they are approaching throttling boundaries. Proper client-side implementations will read these headers and proactively slow down their request rates in response.

---

### 39. What is API versioning?
"API versioning is the practice of managing changes to an API contract so that existing clients don't break when new, incompatible changes are deployed.

In microservices, frontend applications (mobile, web) might upgrade at different paces. If I rename a field from `userName` to `fullName`, an older mobile app looking for `userName` will crash.

By creating version `v2` of the API, I allow older clients to comfortably use `v1` while new clients migrate to `v2`. I maintain both until `v1` traffic drops to zero."

#### Indepth
Semantic Versioning (MAJOR.MINOR.PATCH) is standard. Only changes that break backward compatibility (removing a field, changing a data type) require a MAJOR version bump. Adding a new optional field is backward-compatible and does not require a new major version.

---

### 40. URI vs header versioning?
"These are the two most common ways to denote an API version.

**URI Versioning** includes the version explicitly in the URL path: `GET /api/v1/users`. It is incredibly simple, clearly visible in browser address bars, and easy to route via the API Gateway. It is my preferred pragmatist approach.

**Header Versioning** (or Content Negotiation) leaves the URL clean (`GET /api/users`) and requires the client to pass the version via an HTTP header (e.g., `Accept-Version: v1` or `Accept: application/vnd.mycompany.v1+json`). It's theoretically more RESTful but makes caching and sharing links via the browser much harder."

#### Indepth
Most major tech companies (Twitter, Stripe, GitHub) have gravitated toward URI versioning simply because the developer experience is significantly easier. You can drop a URI-versioned endpoint into Postman or a browser and see the result instantly without fumbling with headers.
