# 🟢 **1–20: Core Microservices Fundamentals**

### 1. What is Microservices architecture?
"Microservices architecture is an approach where a single application is built as a suite of small, independent services. Each service runs in its own process and communicates using lightweight mechanisms, typically HTTP REST APIs or messaging queues.

Instead of having all features tangled in one massive codebase, I can build and deploy individual components independently. For example, an e-commerce app might have separate services for Users, Orders, and Payments.

What I love about it is the isolation—if the recommendation service crashes or needs a massive scale-up, it doesn't bring the core checkout process down."

#### Indepth
A true microservice embodies the Single Responsibility Principle at an architectural level. It must own its own data (database-per-service paradigm), avoiding direct database sharing to ensure loose coupling. Communication happens strictly via well-defined API contracts.

---

### 2. Why microservices over monolith?
"Microservices solve scaling bottlenecks—both organizational and technical. 

Technically, a monolith scales entirely; if my image processing module is CPU heavy, I have to deploy 10 copies of the *entire* application. With microservices, I only scale that specific module. 

Organizationally, it stops developers from stepping on each other's toes. In a monolith, 50 developers mean constant merge conflicts and slow deployment queues. Microservices allow small, autonomous squads to own a service end-to-end and release multiple times a day."

#### Indepth
Microservices also enable "Polyglot Programming" and "Polyglot Persistence." One team can use Go and Redis for a high-concurrency rate limiter, while another team uses Java and PostgreSQL for transactional billing, selecting the best tool for the specific job.

---

### 3. When should you NOT use microservices?
"I avoid microservices when starting a brand new project ('greenfield') where the domain is poorly understood. 

If we don't know the business boundaries yet, we will inevitably draw the wrong microservice boundaries, leading to a 'distributed monolith'—where every service talks to every other service synchronously. 

I also avoid them if the engineering team is small or lacks DevOps maturity. The operational overhead of deploying, monitoring, and securing 20 services requires robust CI/CD and Kubernetes expertise that small startups often don't have."

#### Indepth
Martin Fowler advocates for the "Monolith First" approach. Build a well-modularized monolith first, figure out the domain logic, and then carve out microservices only when specific modules require independent scaling or deployment lifecycles.

---

### 4. What are the characteristics of microservices?
"There are a few defining characteristics I look for. First, they are **independently deployable**. I should be able to push an update to Service A without touching Service B.

Second, they are **organized around business capabilities**, not technical layers (like having a 'UI team' and a 'DB team'). A microservice team owns the UI, logic, and database for their specific feature. 

Finally, they exhibit **decentralized governance and data management**, relying on smart endpoints and dumb pipes (like simple REST over HTTP) rather than heavy enterprise service buses (ESB)."

#### Indepth
According to the Reactive Manifesto, a well-designed microservice ecosystem should be Responsive, Resilient, Elastic, and Message Driven. Resilience is achieved through patterns like Circuit Breakers, Bulkheads, and Fallbacks to prevent cascading failures.

---

### 5. What are common microservice anti-patterns?
"The most dangerous anti-pattern is the **Distributed Monolith**—building services that are tightly coupled. If one goes down, the whole system halts. 

Another is the **Shared Database** anti-pattern. If Service A and Service B read/write to the same database tables, a schema change in A will unexpectedly break B. 

I also frequently see **Hardcoded IPs/URLs** instead of using Service Discovery, and **Synchronous Chains** (Service A calls B, which calls C, which calls D). A failure or delay in D cascades all the way back to the user."

#### Indepth
The "Mega-Service" anti-pattern occurs when a service grows too large and takes on too many responsibilities, failing the Single Responsibility Principle. Conversely, "Nano-services" happen when services are split too finely, ending up with excessive network overhead just to fulfill a basic business transaction.

---

### 6. What is loose coupling?
"Loose coupling means that services interact with each other without needing to know the internal workings of one another.

If I change the internal code, database schema, or even the programming language of the Payment service, the Order service shouldn't care or break, as long as the API contract remains unchanged.

I achieve this by using versioned APIs, asynchronous messaging (like Kafka), and ensuring services never share a database."

#### Indepth
Loose coupling minimizes dependencies. In eventual consistency architectures, loose coupling is maximized. For instance, instead of the Order service commanding the Inventory service to deduct stock synchronously, it drops an "OrderPlaced" event. The Inventory service listens and acts independently, decoupling their temporal availability.

---

### 7. What is high cohesion?
"High cohesion means that related logic and data belong together in the same service. 

If a business rule changes and I have to deploy updates to five different microservices simultaneously just to release that one feature, that's low cohesion. It tells me my service boundaries are wrong.

I strive for high cohesion because a service should represent a single, focused business capability. 'The code that changes together, stays together.' This drastically reduces cross-service network chatter."

#### Indepth
Cohesion and Coupling are two sides of the same coin. The goal of microservices design is "High Cohesion, Loose Coupling". Finding the right level of cohesion relies heavily on identifying correct "Bounded Contexts" during Domain-Driven Design exercises.

---

### 8. What is bounded context?
"Bounded Context is a core concept from Domain-Driven Design (DDD). It defines the explicit boundary within which a particular domain model is valid.

For example, the term 'User' means something entirely different in a Billing context (where it needs credit card details) versus an Authentication context (where it only needs password hashes and roles). 

Instead of building one massive 'User' table that serves everyone, I create two separate microservices, each with its own tailored representation of a User. This prevents massive, confusing data models."

#### Indepth
In DDD, crossing a bounded context boundary means data must go through an translation map (Anti-Corruption Layer) or adhere strictly to public APIs. A single microservice should ideally encapsulate exactly one Bounded Context.

---

### 9. What is domain-driven design (DDD)?
"Domain-Driven Design is a software design approach focused on modeling software to match a domain according to input from that domain's experts.

It provides a framework for breaking a large, complex business into logical components. In the microservices world, I use DDD as the primary tool to decide where to draw my service boundaries.

By sitting with domain experts (like the shipping team), we identify Entities, Value Objects, and Aggregates, and group them into Bounded Contexts, which seamlessly map directly to my microservices."

#### Indepth
DDD differentiates between the Problem Space (Subdomains like Core, Generic, Supporting) and the Solution Space (Bounded Contexts). Properly mapping the Subdomains to Bounded Contexts ensures that the architecture is driven by business needs, rather than technical convenience.

---

### 10. What is database-per-service pattern?
"The database-per-service pattern is the golden rule of microservices data architecture: each microservice tightly owns its own data store, and no other service can access it directly.

If the Order service needs customer information, it cannot run a SQL query against the Customer database. It must make an HTTP or gRPC call to the Customer service's API.

I use this pattern to ensure true loose coupling. It allows me to change a service's database schema or even migrate from PostgreSQL to MongoDB without asking permission from other teams."

#### Indepth
While excellent for isolation, this pattern introduces the massive challenge of Distributed Transactions. Implementing operations that span multiple services (like creating an order and deducting inventory) requires complex patterns like Sagas, as traditional ACID transactions using Two-Phase Commit (2PC) are not viable over HTTP.

---

### 11. Why is shared database bad in microservices?
"A shared database violently breaks encapsulation. 

If five microservices all read and write to the same `orders` table, and I decide to rename a column or change a data type to optimize my service, I immediately crash the other four services. 

It also restricts technology choices—everyone is forced to use the same relational DB even if a graph or document DB would suit their specific service better. Finally, it creates a massive single point of failure and scaling bottleneck."

#### Indepth
The only exception (which is still generally frowned upon) is when breaking apart a legacy monolith, where a shared database might be used temporarily during the transition phase. This is sometimes paired with the Strangler Fig Pattern until the data can be decoupled safely.

---

### 12. What is polyglot persistence?
"Polyglot persistence is the practice of using different database technologies within the same system to handle different types of data storage needs.

Because microservices enforce the database-per-service rule, polyglot persistence is naturally enabled. 

For instance, I might use Neo4j (a graph DB) for my Social Recommendations service, Redis for caching user sessions, ElasticSearch for product search, and PostgreSQL for financial transactions. I can pick the absolute best tool for each specific job."

#### Indepth
While technologically empowering, polyglot persistence drastically increases operational complexity. The Operations/DevOps team now has to learn how to deploy, back up, monitor, and scale four entirely different database paradigms instead of just one standard RDBMS.

---

### 13. What is API composition?
"API composition is a pattern used to retrieve data that spans multiple microservices. 

Since I can't write a SQL `JOIN` across different databases, I have to perform a 'network join'. An API Composer (usually an API Gateway or a dedicated aggregator service) queries the required microservices and stitches the results together in memory.

For example, to display an Order History page, the composer calls the Order Service, invokes the Product Service to get product names, invokes the Delivery Service for tracking status, and merges it into one big JSON response for the client."

#### Indepth
API Composition works well for simple queries but becomes terribly inefficient for complex, multi-service aggregate queries (like "find all users who spent over $100 on electronics last month"). For such queries, the CQRS (Command Query Responsibility Segregation) pattern with materialized views is preferred.

---

### 14. What is backend for frontend (BFF)?
"The BFF pattern involves creating a dedicated API Gateway tailored specifically for a single type of client interface.

A mobile app and a complex desktop web portal need entirely different data formats and payload sizes. Instead of bloat-fitting a single API Gateway to serve both, I create a 'Mobile BFF' and a 'Web BFF'.

The Web BFF might aggregate data from five microservices, while the Mobile BFF fetches a smaller subset of data from two services to save bandwidth. The frontend teams usually own their respective BFFs."

#### Indepth
Using GraphQL in a BFF is highly popular. The BFF acts as a GraphQL server, allowing the specific client to declare exactly what data it needs. The BFF then orchestrates the underlying REST/gRPC microservice calls to fetch precisely that data and nothing more.

---

### 15. What is strangler pattern?
"The Strangler pattern (or Strangler Fig) is a strategy for migrating a monolithic application to a microservices architecture gradually.

Instead of a risky 'big bang' rewrite, I put a proxy/API Gateway in front of the legacy monolith. I then take one feature (e.g., User Profile), rewrite it as a new microservice, and update the proxy routing to divert 'Profile' traffic to the new service.

Over months or years, the new microservices 'strangle' the old monolith until it handles no traffic and can be safely deleted."

#### Indepth
This pattern significantly de-risks migrations. It allows for continuous delivery of new business value while paying down technical debt. Importantly, if the new microservice fails or performs poorly, the gateway routing can be instantly reverted back to the legacy monolith as a fallback mechanism.

---

### 16. What is service registry?
"A Service Registry is the central database containing the network locations (IP addresses and ports) of all active microservice instances.

In dynamic environments like the cloud, IP addresses change constantly due to autoscaling or instance failures. I can't hardcode them. 

When a microservice boots up, it registers itself with the registry. It's essentially the 'phonebook' of the microservice ecosystem. Examples include Netflix Eureka, Consul, and Apache Zookeeper."

#### Indepth
Service registries must be highly available and strictly consistent. If the registry goes down, services cannot find each other. Most registries require services to send regular "heartbeats"; if a heartbeat is missed, the registry removes that instance from its database so traffic isn't sent to a dead node.

---

### 17. What is service discovery?
"Service Discovery is the process of a microservice or client querying the Service Registry to find the current IP address of a service it needs to communicate with.

If the Order service needs to call the Payment service, it asks the Service Discovery mechanism: 'Give me the IP of the Payment Service'. It receives the IPs, picks one, and makes the HTTP request.

I rely on this completely to ensure robust communication in a highly volatile cloud infrastructure."

#### Indepth
Kubernetes natively provides Service Discovery using its internal CoreDNS and ClusterIP abstractions. In a pure K8s environment, application-level tools like Eureka are often discarded; you simply make an HTTP call to the service name (e.g., `http://payment-service:8080`), and K8s DNS routes the traffic to a healthy pod.

---

### 18. Client-side vs server-side discovery?
"In **Client-side discovery**, the calling microservice directly queries the Service Registry, gets a list of IPs, and uses a client-side load balancer (like Spring Cloud LoadBalancer) to pick an instance and make the call. 

In **Server-side discovery**, the calling microservice sends the request to a central router or load balancer. The router queries the registry and forwards the request. 

I prefer server-side (like Kubernetes Services or AWS ELB) because it offloads the complex discovery and load-balancing logic away from the microservice code."

#### Indepth
Client-side discovery reduces network hops but ties your services to a specific programming language library (like Java's Eureka Client). Server-side discovery adds a network hop (the router) but is completely language agnostic, which is vital in polyglot microservice environments.

---

### 19. What is API Gateway?
"An API Gateway is a server that acts as the single entry point into the microservices landscape for all external clients.

Rather than a mobile client knowing about 20 different internal microservice IPs, it sends all requests to `api.myapp.com`. The gateway checks the URL path and routes it to the correct internal service.

It's a mandatory architectural layer for me, as it drastically simplifies the client interface and securely hides our internal network structure."

#### Indepth
Because all external traffic flows through it, the API Gateway is a single point of failure (SPOF) and a potential bottleneck. It must be highly available, clustered, and capable of extreme horizontal scaling. Popular implementations include Kong, Spring Cloud Gateway, and AWS API Gateway.

---

### 20. What are responsibilities of API Gateway?
"The primary responsibility is request routing and API composition. However, I use it as a powerful edge layer to handle cross-cutting concerns for all services.

It handles **Authentication & Authorization** (validating JWT tokens before hitting services). It handles **Rate Limiting** (blocking abusers). It provides **SSL Termination** (decrypting HTTPS so internal traffic can be faster plain HTTP). 

It also provides **Caching**, **CORS management**, and **Metrics collection**, ensuring my internal microservices can focus purely on business logic rather than security plumbing."

#### Indepth
By offloading JWT validation to the Gateway, internal microservices only need to trust the headers passed by the Gateway. However, one must be careful not to put business logic inside the Gateway (an anti-pattern known as "Smart Pipes, Dumb Endpoints"), as this recreates the monolith tightly coupled at the API layer.
