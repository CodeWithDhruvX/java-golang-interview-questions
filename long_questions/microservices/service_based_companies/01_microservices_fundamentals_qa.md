# 🏗️ Microservices — Service-Based Companies Q&A

> **Level:** 🟢 Junior – 🟡 Mid
> **Asked at:** TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini, HCL, Tech Mahindra

---

## Q1. What is the difference between monolithic and microservices architecture?

"A monolithic architecture means the entire application—user interface, business logic, and database access—is bundled together into a single deployable unit (like a single `.war` file in Java deployments).
- **Pros:** Easy to develop initially, simple to test end-to-end, and simple to deploy.
- **Cons:** Any small change requires deploying the entire application. It's hard to scale specific parts (if the reporting module needs more CPU, you have to scale the whole app), and it becomes harder for large teams to work together without conflicts.

Microservices architecture breaks down the application into small, independent services based on business domains (e.g., an Order Service, a Payment Service, a User Service). Every service runs as its own process and communicates through APIs (usually HTTP/REST).
- **Pros:** Each service can be developed in a different technology stack. Teams can deploy updates to the Order Service without touching Payment. Services can scale independently.
- **Cons:** It's complex to manage network communication, testing is harder since you have to test across services, and managing distributed databases (each service having its own DB) is difficult."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** Almost every entry-level and mid-level enterprise interview (TCS, Wipro, Infosys).

#### Indepth
**When to choose what:** Monoliths are often best for starting brand-new projects (startups) where the domain boundaries aren't clear yet. 'Microservices first' is often an anti-pattern that leads to 'Distributed Monoliths'—services that are deployed separately but are so tightly coupled they have to be released together anyway.

---

## Q2. How do microservices communicate with each other? What is synchronous vs. asynchronous communication?

"Microservices communicate over the network. There are two primary ways:

**1. Synchronous Communication:**
Service A sends a request to Service B and *waits* for the response before continuing its work. The most common implementation is **REST over HTTP** or, increasingly, **gRPC** for internal service-to-service calls.
*Use Case:* When the user needs an immediate response, like checking account balance. If the Account service is down, the user gets an error immediately.

**2. Asynchronous Communication:**
Service A sends a message to an intermediary (a message broker like **Apache Kafka** or **RabbitMQ**) and immediately continues its work without waiting for Service B to reply. Service B picks up the message when it's ready.
*Use Case:* Triggering an email notification. The Order service publishes an 'OrderCreated' event. The Email service consumes it and sends the email later. If the Email service is down, the message stays in the queue until it comes back up, so the Order process never fails because of an email issue."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** Standard architecture theory question at all service companies. 

#### Indepth
**Tight Coupling Risk:** Designing synchronous REST calls across a chain of 4 services (A -> B -> C -> D) is incredibly fragile. If Service D is down, the entire chain fails. This is called 'distributed synchronous coupling' and is why Asynchronous (Event-Driven) architectures are preferred for most background business processes.

---

## Q3. What is an API Gateway and why do we need it in Microservices?

"Instead of having frontend applications (Mobile apps, Web browsers) call 10 different internal microservices directly, we place an **API Gateway** between the client and the microservices. It acts as the single entry point to the system.

Why we need it:
1. **Routing:** It takes a request (e.g., `/api/orders`) and routes it to the specific internal Docker container running the Order Service.
2. **Security & Authentication:** Instead of every microservice validating JWT tokens, the Gateway validates the token once and passes the user details down to the internal services.
3. **Cross-Cutting Concerns:** It handles Rate Limiting (preventing DDOS), CORS configuration, and SSL termination.
4. **Aggregation:** (Sometimes) The Gateway can call the User service and the Order service simultaneously, merge the JSON responses, and return a single response to the Mobile app to save network trips."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS, Cognizant, Accenture (commonly asked when migrating legacy systems).

#### Indepth
**Common Gateway Solutions:** Netflix Zuul (older), Spring Cloud Gateway (modern Java ecosystem standard), Kong (Nginx-based), AWS API Gateway (Managed).

---

## Q4. Explain Service Discovery in Microservices. What is Eureka?

"In a monolith, Service A knows where Service B is because they are compiled together. In microservices, services are deployed in containers whose IP addresses change dynamically (especially in Kubernetes or auto-scaling environments). Hardcoding IP addresses like `http://192.168.1.5:8080/users` will fail instantly.

**Service Discovery** solves this. It has two parts:
1. **Service Registry (like Netflix Eureka):** A central database containing the network locations (IP and Port) of all available service instances.
2. **Discovery Process:** When the 'Order instance' starts up, it registers its IP with Eureka. When the 'Payment instance' needs to call 'Order', it asks Eureka: 'Give me the IP for the Order Service'. Eureka returns the IP, and Payment makes the call.

*Note: In modern deployments using Kubernetes, Netflix Eureka is often skipped because Kubernetes has built-in DNS-based service discovery (e.g., calling `http://order-service:8080` internally).* "

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Infosys, Tech Mahindra (specifically focusing on the Spring Cloud Netflix SOS ecosystem).

#### Indepth
**Client-Side vs. Server-Side Discovery:** 
Eureka is 'Client-Side Discovery'; the caller service queries the registry and makes the HTTP call itself (usually using a Ribbon or Spring Cloud LoadBalancer to pick between multiple IPs). 
Kubernetes provides 'Server-Side Discovery'; the caller just sends traffic to a stable virtual IP (a K8s 'Service'), and K8s load balances it to the backing pods.

---

## Q5. What is the role of a Circuit Breaker? How does it help in a Microservices ecosystem?

"When one microservice depends on another over the network, that network call can fail or hang indefinitely. If Service A waits for a slow Service B, Service A's threads will block. If enough requests come in, Service A crashes.

A **Circuit Breaker** (like Resilience4j or Netflix Hystrix) wraps the outgoing network call. It acts like an electrical circuit breaker:

1. **Closed State:** Everything is normal. Requests pass through to Service B.
2. **Open State:** If Service B fails consecutively (or gets too slow) beyond a threshold (e.g., 50% failures in a 10s window), the breaker 'trips' open. Further calls from Service A to Service B are failed *immediately* without attempting the network call. This saves Service A from crashing and gives Service B time to recover.
3. **Half-Open State:** After a timeout, the breaker lets a few test requests through. If they succeed, it closes again. If they fail, it stays open.

It prevents 'cascading failures' across the architecture."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS, Capgemini, Accenture (essential microservice resilience question).

#### Indepth
**Fallback Method:** When the circuit is 'Open', the framework allows you to execute a fallback method locally. For example, if the live 'Recommendation Service' is down, the fallback method might return a hardcoded list of globally popular products instead of throwing a 500 error to the UI.

---

## Q6. How do you manage configuration in a microservices application? What is Spring Cloud Config?

"When you have 50 microservices moving through Dev, QA, UAT, and Prod environments, managing database passwords, API keys, and feature flags inside individual `application.properties` files becomes an operational nightmare. You'd have to rebuild and redeploy services just to change a property.

To solve this, we use **Centralized Configuration Management**.

**Spring Cloud Config** is a popular tool for this. Instead of reading local files, a microservice reaches out to the central 'Config Server' on startup. The Config Server reads the configuration properties from a central, secured Git repository (or Vault) based on the environment (e.g., `application-dev.yml`).

If a property changes in Git, we can update the live microservices without restarting them (using tools like Spring Boot `@RefreshScope` and Spring Cloud Bus to broadcast the change)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Wipro, Infosys (DevOps implementation in Java projects).

#### Indepth
**Kubernetes Alternatives:** If your client uses Kubernetes heavily, Spring Cloud Config is often replaced by Kubernetes `ConfigMaps` and `Secrets`. The container orchestration layer handles mounting these properties into the container as files or environment variables at runtime, which is more language-agnostic than the Spring-specific ecosystem.
