# Microservices Architecture and Spring Cloud - Interview Questions and Answers

## 1. What are Microservices? Compare Monolithic vs. Microservice Architecture.
**Answer:**
A Microservices architecture is designing a large application as a suite of small, independently deployable services, built around specific business domains. Each service runs in its own process and communicates through lightweight mechanisms, often an HTTP resource API.

**Monolithic Architecture:**
- **Structure:** All components (UI, business logic, data access) are packaged into a single deployable unit (e.g., one huge WAR file) and usually share a single massive database.
- **Pros:** simple to develop initially, simple to test, simple to deploy early on.
- **Cons:** Becomes too large for any single developer to understand fully; a small change requires rebuilding and redeploying the entire application; scaling is inefficient (must scale the whole app even if only one module is under load); tightly coupled technologies (locked into one tech stack).

**Microservices Architecture:**
- **Structure:** Divided into many small services (e.g., User Service, Product Service, Order Service), each with its own dedicated database.
- **Pros:** Scalability (scale only the services under load); technological freedom (different services can use different languages or DBs); fault isolation (if the Order service fails, the User service stays up); smaller, focused teams.
- **Cons:** High operational complexity (managing 50 services instead of 1); complex distributed transactions; potential network latency; complex debugging and tracing across services.

## 2. Explain inter-service communication in Microservices (RestTemplate vs. WebClient vs. OpenFeign).
**Answer:**
Services must communicate with each other over the network, usually synchronously via HTTP/REST.

1. **`RestTemplate`:** The traditional Spring tool for making synchronous HTTP requests. It is blocking, meaning the thread waits until a response is received. It requires boilerplate code to construct URLs, set headers, and map responses.
2. **`WebClient`:** Part of Spring WebFlux. It is a non-blocking, reactive client. It uses a fluent builder API and is much more resource-efficient under high concurrency because it doesn't tie up threads waiting for network I/O.
3. **`Spring Cloud OpenFeign`:** A declarative web service client developed by Netflix and wrapped by Spring. It is the most robust and easiest way to implement synchronous communication between Spring Boot services.
    - **How it works:** Instead of manually writing HTTP calls, you simply define a Java interface and annotate it with `@FeignClient("service-name")` and Spring MVC annotations (`@GetMapping`, `@PathVariable`).
    - At runtime, Spring generates the proxy implementation to execute the HTTP requests, load balance them, and map the JSON response back to Java DTOs seamlessly.

## 3. What is Service Registry and Discovery, and why is it needed? Explain Eureka.
**Answer:**
In a microservices ecosystem, instances of services spin up and down dynamically due to auto-scaling or failures. Their IP addresses and ports change constantly. Hardcoding IPs in configuration properties is impossible.

**Service Registry (Netflix Eureka):**
- It acts as a dynamic "phone book" or address directory for microservices.
- **Registration:** When a service instance (e.g., an Order Service instance on port 8081) starts, it registers its logical name (`ORDER-SERVICE`), IP address, and port with the Eureka Server. It sends periodic "heartbeats" to prove it's still alive.
- **Discovery:** When the Client Service (e.g., API Gateway) needs to call the Order Service, it queries the Eureka Server for the name `ORDER-SERVICE`. Eureka responds with a list of all currently active instances (IP/Ports).
- **Client-Side Load Balancing:** Tools like OpenFeign use the list returned by Eureka to distribute requests evenly among available instances, bypassing the need for a traditional hardware load balancer.

## 4. What is an API Gateway, and what role does Spring Cloud Gateway play?
**Answer:**
An API Gateway is a central entry point for all client requests (frontends, mobile apps) into the microservices ecosystem. It prevents clients from having to know the addresses of individual microservices and talking to them directly.

**Spring Cloud Gateway:**
Built on Spring WebFlux (reactive, non-blocking), it provides a simple, effective way to route traffic.

**Key Responsibilities:**
- **Routing:** Directing traffic to specific microservices based on URL paths or headers. Examples: `/api/orders/**` goes to the Order Service.
- **Cross-Cutting Concerns:** Centralizing authentication (validating JWTs once at the gateway rather than in every single service), SSL termination, rate limiting, logging, and CORS handling.
- **Load Balancing:** It integrates tightly with Eureka to load balance incoming client requests across multiple instances of backend services.

## 5. How do you implement resilience in microservices using Resilience4j (Circuit Breaker Pattern)?
**Answer:**
If Service A calls Service B, and Service B is slow or unresponsive down, Service A's threads will block waiting. This can cascade, taking down the entire system. Resilience patterns prevent this.

**Spring Cloud CircuitBreaker / Resilience4j:**
- **Closed State:** Normal operation. Requests flow freely. The circuit breaker monitors success/failure rates.
- **Open State:** If the failure rate (or slow call rate) exceeds a configured threshold (e.g., 50% failures), the circuit breaker "opens." All subsequent calls to Service B fail instantly without actually making the network request. This prevents overloading the failing service and failing fast.
- **Half-Open State:** After a configured wait duration, it allows a limited number of test requests through to see if Service B has recovered. If they succeed, it closes; if they fail, it opens again.

**Usage:** You annotate Feign client methods or service methods with `@CircuitBreaker(name="serviceB", fallbackMethod="getDefaultFallback")`. If the circuit is open or an exception occurs, the system immediately executes the `fallbackMethod` (e.g., returning cached data or a default value) instead of failing.

## 6. What is centralized configuration in Microservices? Explain Spring Cloud Config.
**Answer:**
In an architecture with 50 microservices running across multiple environments (Dev, QA, Prod), managing `application.properties` files inside individual codebases is an operational nightmare. Changing a database password would require rebuilding and redeploying 50 projects.

**Spring Cloud Config Server:**
- It provides server and client-side support for externalized, centralized configuration.
- **Server:** A standalone Spring Boot application that exposes configuration properties securely over HTTP. The backend storage for these properties is typically a remote Git repository (e.g., GitHub, GitLab).
- **Client:** Every other microservice acts as a Config Client. Upon startup, before loading its own context, the client contacts the Config Server, requests its specific configuration (based on application name and active profile), and loads those properties into its Spring Environment.
- **Benefit:** Operations teams can change configurations dynamically in the central Git repository without touching application code.

## 7. How do you handle configuration changes dynamically without restarting services (Spring Cloud Bus)?
**Answer:**
While Spring Cloud Config allows centralized properties, clients only fetch them upon application *startup*. If you modify a property in Git, running services won't see it until restarted.

**Spring Cloud Bus:**
- It links the nodes of a distributed system with a lightweight message broker (like RabbitMQ or Kafka).
- **Process:**
    1. Update the property in the Git repository.
    2. Send an HTTP POST request to a specific Webhook endpoint (e.g., `/actuator/bus-refresh`) on the Config Server or any participating node.
    3. The receiving node publishes a "Refresh Event" to the message broker.
    4. All microservices listening on the bus receive this event, realize properties have changed, and dynamically reload their configuration beans (specifically those annotated with `@RefreshScope`) without restarting the JVM process.

## 8. What is Distributed Tracing, and how do Spring Cloud Sleuth and Zipkin help?
**Answer:**
A single client request (e.g., "Checkout Cart") might trigger a complex chain of internal HTTP calls: API Gateway -> Cart Service -> Payment Service -> Inventory Service -> Database. If the "Checkout Cart" request fails or is alarmingly slow, it's incredibly difficult to determine *which* specific hop in the chain caused the issue looking at standard, isolated log files.

- **Spring Cloud Sleuth (Now part of Micrometer Tracing in Spring Boot 3):**
    - Automatically instruments incoming and outgoing HTTP requests within the application.
    - Adds an overall **Trace ID** (unique per client request) to track the entire journey.
    - Adds a **Span ID** to track specific network hops or logical operations within that trace.
    - It automatically injects these IDs into the SLF4J MDC (Mapped Diagnostic Context), meaning every single log generated by any service participating in the request contains the `[TraceID, SpanID]`. You can easily grep across multiple servers using the Trace ID.
- **Zipkin:**
    - A dedicated UI and storage backend system for distributed tracing. Spring Boot apps send timing and tracing metadata to Zipkin over HTTP or Kafka. Zipkin aggregates this data and provides a visual timeline (Gantt chart) to immediately identify latency bottlenecks (e.g., "The Payment Service took 3000ms out of the total 3200ms processing time").

## 9. How do you centralize logs using the ELK Stack (Elasticsearch, Logstash, Kibana)?
**Answer:**
In a distributed system, you cannot SSH into dozens of containers to read text log files. Logs must be aggregated into a central, searchable platform.

**The ELK Stack:**
- **Logstash (or Filebeat/Fluentd):** The shipper/pipeline. It installs alongside your microservices or runs centrally. It ingests log files, parses the unstructured text log lines into structured JSON fields (extracting timestamps, log levels, Trace IDs, and messages), and ships them off.
- **Elasticsearch:** A highly scalable, distributed NoSQL search and analytics engine. It stores the parsed JSON logs, indexing every single field to allow lightning-fast text searches across billions of log entries.
- **Kibana:** The visualization UI dashboard connected to Elasticsearch. Support teams use Kibana to query the central repository (e.g., searching `traceId: "xyz123"` across all service indices simultaneously) and build dashboards showing system health based on error logs.
