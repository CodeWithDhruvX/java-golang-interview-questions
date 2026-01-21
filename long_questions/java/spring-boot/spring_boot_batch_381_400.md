## 🔹 Section 9: Microservices & Spring Cloud (381-390)

### Question 381: What is Spring Cloud Sleuth and how does it enhance traceability?

**Answer:**
(See Q229). Note: In Boot 3, replaced by Micrometer Tracing.

---

### Question 382: How does Spring Cloud Config work internally?

**Answer:**
Client app requests config from Config Server (`http://config-server/app/profile`).
Server fetches from Git/Vault.
Returns JSON.
Client `ConfigServicePropertySourceLocator` adds it to Environment.
Done during specific "Bootstrap" phase (before main context).

---

### Question 383: What is `spring.cloud.bootstrap.enabled=true` used for?

**Answer:**
In Spring Cloud 2020+, Bootstrap context (`bootstrap.properties`) was deprecated in favor of `spring.config.import`.
Setting this (with dependency `spring-cloud-starter-bootstrap`) re-enables the old behavior.

---

### Question 384: How do you use Feign clients in Spring Boot?

**Answer:**
Declarative REST Client.
1.  `@EnableFeignClients`.
2.  Interface `@FeignClient(name="user-service")`.
3.  Method `@GetMapping("/users/{id}") User getById(...)`.
Spring creates a proxy (using Ribbon/LoadBalancer) to call the remote service.

---

### Question 385: What is Spring Cloud Bus and how does it work with Kafka or RabbitMQ?

**Answer:**
Links nodes of a distributed system usually for management instructions.
Common use: **Broadcast Configuration Change**.
POST `/actuator/bus-refresh` on one node -> Message on Kafka -> All nodes receive -> All nodes refresh config.

---

### Question 386: What is service discovery, and how does Eureka support it?

**Answer:**
Instead of hardcoding URLs (`http://192.168.1.5:8080`), services register with Eureka (`app-name`, `ip`, `port`).
Client asks Eureka: "Where is `app-name`?".
Eureka returns list of IPs. Client load balances.

---

### Question 387: What are ribbon clients, and how are they configured?

**Answer:**
(Legacy). Client-side Load Balancer.
Replaced by **Spring Cloud LoadBalancer**.
Config: `service-name.ribbon.listOfServers=...`.

---

### Question 388: How do you implement distributed tracing with Zipkin and Sleuth?

**Answer:**
Sleuth generates Spans.
Config: `spring.zipkin.baseUrl=http://zipkin:9411`.
Logs are pushed (HTTP/Kafka) to Zipkin server which UI visualizes latencies.

---

### Question 389: How do you configure load balancing across microservices?

**Answer:**
Using `@LoadBalanced` on `RestTemplate` or `WebClient.Builder`.
Or Feign (Auto-integrated).
Resolves service names (`http://user-service/`) to IPs using Discovery Client.

---

### Question 390: What is the purpose of Circuit Breakers in Spring Boot microservices?

**Answer:**
(Resilience4j).
Prevents cascading failures.
If Service A calls Service B and B is slow/down, the Circuit "Open" after N failures.
Subsequent calls fail fast (without waiting) until B recovers.

## 🔹 Section 10: Advanced DevOps & CI/CD (391-400)

### Question 391: How do you build Spring Boot apps as OCI-compliant containers using buildpacks?

**Answer:**
(See Q261). `mvn spring-boot:build-image`.

---

### Question 392: How do you configure a multistage Dockerfile for a Spring Boot app?

**Answer:**
(See Q264).
Stage 1 (Builder): Maven image, run `mvn package`.
Stage 2 (Runner): JRE image, copy JAR from Stage 1.

---

### Question 393: How do you push Spring Boot metrics to Prometheus?

**Answer:**
(See Q223).

---

### Question 394: How to automate release versioning using Maven or Gradle plugins?

**Answer:**
`maven-release-plugin`.
Updates `pom.xml` version (SNAPSHOT -> Release), tags git, deploys to Nexus, increments to next SNAPSHOT.

---

### Question 395: What is the role of `spring-boot-maven-plugin` and `spring-boot-gradle-plugin`?

**Answer:**
1.  **Repackage:** Creates the "Fat JAR" (executable with deps).
2.  **Run:** `mvn spring-boot:run`.
3.  **Build Image:** OCI image creation.

---

### Question 396: How to detect and fail builds on deprecated Spring APIs?

**Answer:**
Compiler argument `-Xlint:deprecation`.
Configure Maven Compiler Plugin to `failOnWarning` (strict).
Or use `OpenRewrite` to auto-migrate.

---

### Question 397: How to dynamically reload configuration in Kubernetes with Spring Boot?

**Answer:**
Use `spring-cloud-starter-kubernetes-config`.
Changes in K8s `ConfigMap` trigger a refresh event in the app (App must have Watched enabled or use a Poller).

---

### Question 398: What is the use of Spring Boot Admin in DevOps pipelines?

**Answer:**
(See Q94). Visualization.

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

---
