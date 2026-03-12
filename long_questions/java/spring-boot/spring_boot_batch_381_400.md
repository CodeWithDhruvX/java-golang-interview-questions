## 🔹 Section 9: Microservices & Spring Cloud (381-390)

### Question 381: What is Spring Cloud Sleuth and how does it enhance traceability?

**Answer:**
(See Q229). Note: In Boot 3, replaced by Micrometer Tracing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Cloud Sleuth and how does it enhance traceability?
**Your Response:** "Spring Cloud Sleuth enhances traceability by adding trace and span IDs to log messages, allowing me to track requests across multiple microservices. It automatically generates unique trace IDs for each request and propagates them through HTTP headers. In Spring Boot 3, Sleuth has been replaced by Micrometer Tracing, which provides similar functionality. This distributed tracing helps me debug issues by following a single request through multiple services, making it essential for microservice architectures."

---

### Question 382: How does Spring Cloud Config work internally?

**Answer:**
Client app requests config from Config Server (`http://config-server/app/profile`).
Server fetches from Git/Vault.
Returns JSON.
Client `ConfigServicePropertySourceLocator` adds it to Environment.
Done during specific "Bootstrap" phase (before main context).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Cloud Config work internally?
**Your Response:** "Spring Cloud Config works through a client-server architecture. The client app requests configuration from the Config Server using an endpoint like `http://config-server/app/profile`. The server fetches configuration from Git or Vault and returns it as JSON. The client's `ConfigServicePropertySourceLocator` adds this to the Environment during the bootstrap phase, which happens before the main application context is created. This ensures configuration is available early in the startup process."

---

### Question 383: What is `spring.cloud.bootstrap.enabled=true` used for?

**Answer:**
In Spring Cloud 2020+, Bootstrap context (`bootstrap.properties`) was deprecated in favor of `spring.config.import`.
Setting this (with dependency `spring-cloud-starter-bootstrap`) re-enables the old behavior.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `spring.cloud.bootstrap.enabled=true` used for?
**Your Response:** "In Spring Cloud 2020+, the bootstrap context using `bootstrap.properties` was deprecated in favor of `spring.config.import`. However, if I have legacy applications that depend on the bootstrap mechanism, I can re-enable it by setting `spring.cloud.bootstrap.enabled=true` and adding the `spring-cloud-starter-bootstrap` dependency. This restores the old behavior where bootstrap configuration is loaded before the main application context, helping with migration from older Spring Cloud versions."

---

### Question 384: How do you use Feign clients in Spring Boot?

**Answer:**
Declarative REST Client.
1.  `@EnableFeignClients`.
2.  Interface `@FeignClient(name="user-service")`.
3.  Method `@GetMapping("/users/{id}") User getById(...)`.
Spring creates a proxy (using Ribbon/LoadBalancer) to call the remote service.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use Feign clients in Spring Boot?
**Your Response:** "I use Feign clients as declarative REST clients. First, I enable them with `@EnableFeignClients`. Then I create an interface annotated with `@FeignClient(name='user-service')` and define methods like `@GetMapping('/users/{id}') User getById(...)`. Spring creates a proxy that uses the load balancer to call the remote service. This approach eliminates the boilerplate of writing HTTP client code - I just define an interface and Spring handles the rest, including load balancing and error handling."

---

### Question 385: What is Spring Cloud Bus and how does it work with Kafka or RabbitMQ?

**Answer:**
Links nodes of a distributed system usually for management instructions.
Common use: **Broadcast Configuration Change**.
POST `/actuator/bus-refresh` on one node -> Message on Kafka -> All nodes receive -> All nodes refresh config.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Cloud Bus and how does it work with Kafka or RabbitMQ?
**Your Response:** "Spring Cloud Bus links nodes of a distributed system for management operations. The most common use is broadcasting configuration changes. When I POST to `/actuator/bus-refresh` on one node, it publishes a message on Kafka or RabbitMQ. All nodes receive this message and refresh their configuration. This eliminates the need to restart services for configuration changes. It's particularly useful in microservice environments where I need to propagate configuration updates across multiple instances simultaneously."

---

### Question 386: What is service discovery, and how does Eureka support it?

**Answer:**
Instead of hardcoding URLs (`http://192.168.1.5:8080`), services register with Eureka (`app-name`, `ip`, `port`).
Client asks Eureka: "Where is `app-name`?".
Eureka returns list of IPs. Client load balances.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is service discovery, and how does Eureka support it?
**Your Response:** "Service discovery eliminates hardcoded URLs in microservices. Instead of hardcoding `http://192.168.1.5:8080`, services register with Eureka using their app name, IP, and port. When a client needs to call a service, it asks Eureka 'Where is user-service?' and gets back a list of available IPs. The client then load balances across these instances. This dynamic discovery handles scaling, failover, and deployment changes without code changes."

---

### Question 387: What are ribbon clients, and how are they configured?

**Answer:**
(Legacy). Client-side Load Balancer.
Replaced by **Spring Cloud LoadBalancer**.
Config: `service-name.ribbon.listOfServers=...`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are ribbon clients, and how are they configured?
**Your Response:** "Ribbon was a client-side load balancer in Spring Cloud, but it's now considered legacy and has been replaced by Spring Cloud LoadBalancer. With Ribbon, I configured load balancing using properties like `service-name.ribbon.listOfServers=...` to specify server lists. Ribbon handled client-side load balancing logic, retry policies, and failure detection. While Ribbon is still available for legacy applications, new projects should use Spring Cloud LoadBalancer which provides similar functionality with better maintenance and integration."

---

### Question 388: How do you implement distributed tracing with Zipkin and Sleuth?

**Answer:**
Sleuth generates Spans.
Config: `spring.zipkin.baseUrl=http://zipkin:9411`.
Logs are pushed (HTTP/Kafka) to Zipkin server which UI visualizes latencies.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement distributed tracing with Zipkin and Sleuth?
**Your Response:** "I implement distributed tracing by having Sleuth generate spans for each operation. I configure the Zipkin base URL with `spring.zipkin.baseUrl=http://zipkin:9411`. Sleuth automatically sends trace data to Zipkin via HTTP or Kafka. The Zipkin server provides a UI that visualizes the trace data, showing latencies and dependencies between services. This helps me identify performance bottlenecks and understand how requests flow through my microservice architecture."

---

### Question 389: How do you configure load balancing across microservices?

**Answer:**
Using `@LoadBalanced` on `RestTemplate` or `WebClient.Builder`.
Or Feign (Auto-integrated).
Resolves service names (`http://user-service/`) to IPs using Discovery Client.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure load balancing across microservices?
**Your Response:** "I configure load balancing across microservices by adding `@LoadBalanced` to my `RestTemplate` or `WebClient.Builder` beans. When I make requests to service names like `http://user-service/`, Spring Cloud resolves these to actual IP addresses using the Discovery Client. Feign clients have load balancing integrated automatically. This approach enables client-side load balancing without hardcoding server addresses, making the system resilient to individual server failures and supporting horizontal scaling."

---

### Question 390: What is the purpose of Circuit Breakers in Spring Boot microservices?

**Answer:**
(Resilience4j).
Prevents cascading failures.
If Service A calls Service B and B is slow/down, the Circuit "Open" after N failures.
Subsequent calls fail fast (without waiting) until B recovers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of Circuit Breakers in Spring Boot microservices?
**Your Response:** "Circuit breakers prevent cascading failures in microservice architectures. Using Resilience4j, when Service A calls Service B and B becomes slow or unresponsive, the circuit opens after N failures. Subsequent calls fail fast immediately without waiting, preventing Service A from being affected by B's problems. After a timeout, the circuit moves to a half-open state to test if B has recovered. This pattern improves system resilience and prevents a single failing service from bringing down the entire application."

## 🔹 Section 10: Advanced DevOps & CI/CD (391-400)

### Question 391: How do you build Spring Boot apps as OCI-compliant containers using buildpacks?

**Answer:**
(See Q261). `mvn spring-boot:build-image`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build Spring Boot apps as OCI-compliant containers using buildpacks?
**Your Response:** "I build OCI-compliant containers using Spring Boot's buildpack support with `mvn spring-boot:build-image`. This uses Cloud Native Buildpacks to create optimized, standards-compliant container images without needing a Dockerfile. Buildpacks automatically handle layering, security scanning, and optimization for the specific application. The resulting images follow OCI standards and can run on any compliant container runtime. This approach simplifies the containerization process while ensuring production-ready, secure images."

---

### Question 392: How do you configure a multistage Dockerfile for a Spring Boot app?

**Answer:**
(See Q264).
Stage 1 (Builder): Maven image, run `mvn package`.
Stage 2 (Runner): JRE image, copy JAR from Stage 1.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure a multistage Dockerfile for a Spring Boot app?
**Your Response:** "I configure a multistage Dockerfile to optimize image size. Stage 1 uses a Maven image to run `mvn package` and build the application. Stage 2 uses a minimal JRE image and copies only the JAR file from Stage 1. This approach separates the build environment from the runtime environment, resulting in smaller, more secure production images. The final image contains only what's needed to run the application, not the entire build toolchain, reducing the attack surface and image size."

---

### Question 393: How do you push Spring Boot metrics to Prometheus?

**Answer:**
(See Q223).

---

### Question 394: How to automate release versioning using Maven or Gradle plugins?

**Answer:**
`maven-release-plugin`.
Updates `pom.xml` version (SNAPSHOT -> Release), tags git, deploys to Nexus, increments to next SNAPSHOT.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to automate release versioning using Maven or Gradle plugins?
**Your Response:** "I automate release versioning using the `maven-release-plugin`. It updates the `pom.xml` version from SNAPSHOT to release, tags the commit in Git, deploys the artifact to Nexus, and increments to the next SNAPSHOT version. This automated process ensures consistent versioning across releases. The plugin handles all the Git operations and version updates, reducing manual errors and ensuring proper release management. Similar functionality is available in Gradle with plugins like the release plugin."

---

### Question 395: What is the role of `spring-boot-maven-plugin` and `spring-boot-gradle-plugin`?

**Answer:**
1.  **Repackage:** Creates the "Fat JAR" (executable with deps).
2.  **Run:** `mvn spring-boot:run`.
3.  **Build Image:** OCI image creation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `spring-boot-maven-plugin` and `spring-boot-gradle-plugin`?
**Your Response:** "The Spring Boot plugins are essential for building Spring Boot applications. They provide three main functions: Repackage creates the executable 'fat JAR' with all dependencies included; Run allows me to run the application directly with `mvn spring-boot:run` during development; and Build Image creates OCI-compliant container images using buildpacks. These plugins integrate with the build system to provide all the Spring Boot-specific build tasks needed to create executable applications."

---

### Question 396: How to detect and fail builds on deprecated Spring APIs?

**Answer:**
Compiler argument `-Xlint:deprecation`.
Configure Maven Compiler Plugin to `failOnWarning` (strict).
Or use `OpenRewrite` to auto-migrate.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to detect and fail builds on deprecated Spring APIs?
**Your Response:** "I detect deprecated Spring APIs by configuring the compiler argument `-Xlint:deprecation` and setting the Maven Compiler Plugin to `failOnWarning`. This makes the build fail if any deprecated APIs are used. For automated migration, I can use OpenRewrite to automatically update deprecated code to newer APIs. This proactive approach ensures I stay current with Spring Boot versions and avoid technical debt from deprecated APIs that might be removed in future releases."

---

### Question 397: How to dynamically reload configuration in Kubernetes with Spring Boot?

**Answer:**
Use `spring-cloud-starter-kubernetes-config`.
Changes in K8s `ConfigMap` trigger a refresh event in the app (App must have Watched enabled or use a Poller).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to dynamically reload configuration in Kubernetes with Spring Boot?
**Your Response:** "I enable dynamic configuration reload in Kubernetes using `spring-cloud-starter-kubernetes-config`. When I change a K8s ConfigMap, it triggers a refresh event in the application. The app must have either watch enabled or use a poller to detect changes. This allows configuration updates without restarting the application. The integration automatically maps Kubernetes ConfigMaps and Secrets to Spring Boot properties, providing seamless configuration management in containerized environments."

---

### Question 398: What is the use of Spring Boot Admin in DevOps pipelines?

**Answer:**
(See Q94). Visualization.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of Spring Boot Admin in DevOps pipelines?
**Your Response:** "Spring Boot Admin provides visualization and monitoring for Spring Boot applications in DevOps pipelines. It aggregates data from multiple application instances, showing health status, metrics, and configuration information in a centralized dashboard. I use it to monitor all my microservices from a single interface. It integrates with Actuator endpoints to provide comprehensive monitoring, making it easier to identify issues and track application health across the entire system. This visualization is crucial for operational visibility in production environments."

---

### Question 399: How do you expose Prometheus-compatible metrics with Micrometer?

**Answer:**
(See Q223).

---

### Question 400: How do you automate blue-green deployment strategies with Spring Boot?

**Answer:**
Not strictly a Spring Boot feature.
Platform (K8s/AWS) feature.
Load Balancer points to Blue (v1).
Deploy Green (v2). Wait for Health (`/actuator/health`).
Switch LB to Green.
If Metric (Error Rate) spikes, Switch back to Blue.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you automate blue-green deployment strategies with Spring Boot?
**Your Response:** "Blue-green deployment is primarily a platform feature rather than Spring Boot specific. I implement it by deploying the new version (Green) alongside the current version (Blue). The load balancer initially points to Blue. I wait for health checks on Green using `/actuator/health`, then switch the load balancer to Green. If metrics like error rate spike, I can quickly switch back to Blue. Spring Boot's health endpoints and graceful shutdown support make this strategy smoother, but the actual deployment orchestration is handled by the platform like Kubernetes or cloud load balancers."

---
