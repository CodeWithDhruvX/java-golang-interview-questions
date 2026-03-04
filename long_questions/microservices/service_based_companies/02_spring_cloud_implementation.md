# 🌱 Microservices — Spring Cloud Implementation (Service-Based Companies)

> **Level:** 🟡 Mid
> **Asked at:** TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini, HCL, Tech Mahindra

---

## Q1. How do you make HTTP calls between microservices in Spring Boot? What is Feign Client?

"There are three primary ways to make HTTP calls from one Spring Boot microservice to another:

**1. RestTemplate (Legacy — still common in enterprise projects):**
```java
@Autowired
private RestTemplate restTemplate;

public UserDTO getUserById(Long userId) {
    return restTemplate.getForObject(
        "http://user-service/api/users/" + userId,
        UserDTO.class
    );
}
```

**2. WebClient (Modern — Reactive, non-blocking):**
```java
@Autowired
private WebClient webClient;

public Mono<UserDTO> getUserById(Long userId) {
    return webClient.get()
        .uri("http://user-service/api/users/{id}", userId)
        .retrieve()
        .bodyToMono(UserDTO.class);
}
```
Use `WebClient` for reactive stacks and high-concurrency scenarios.

**3. Feign Client (Declarative — Most popular for service companies):**
```java
@FeignClient(name = "user-service")
public interface UserServiceClient {

    @GetMapping("/api/users/{id}")
    UserDTO getUserById(@PathVariable Long id);
}
```
Feign Client is a declarative way to define inter-service HTTP calls. Instead of writing boilerplate `RestTemplate` code, you define an interface. The Spring Cloud framework generates the implementation automatically. The `name` attribute resolves the service URL via **Eureka Service Registry**.

**Why Feign is preferred in service companies:** Clean code, easy to test with mocks, integrates naturally with Eureka for load-balanced calls, and supports Circuit Breaker (Resilience4j) without additional configuration."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Infosys, Cognizant, HCL — standard Spring Cloud ecosystem question. If your project uses Feign, you'll be asked about it in every interview.

#### Indepth
**Load Balancing with Feign:** When multiple instances of `user-service` are registered in Eureka, Feign uses **Spring Cloud LoadBalancer** (or the older Ribbon) to automatically distribute requests across instances using Round Robin. You get client-side load balancing for free with no additional configuration.

---

## Q2. How do you add distributed logging and tracing in a Spring Boot microservices project?

"Without proper logging, debugging a user complaint in a microservices system is nearly impossible. A single user request may touch 5 different services, each writing to its own separate log file. There's no way to connect the dots without a `TraceId`.

**Step 1: Add Spring Cloud Sleuth (for auto-instrumentation)**
```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-sleuth</artifactId>
</dependency>
```
Once added, Sleuth automatically:
- Generates a `TraceId` (unique per user request) and a `SpanId` (unique per service hop)
- Injects them into `SLF4J` log MDC, so every log line automatically prints: `[service-name, traceId, spanId]`
- Propagates the `TraceId` in outgoing HTTP headers (via Feign/RestTemplate) and incoming headers

**Step 2: Push traces to Zipkin (for visualization)**
```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-sleuth-zipkin</artifactId>
</dependency>
```
```yaml
# application.yml
spring:
  zipkin:
    base-url: http://zipkin-server:9411
  sleuth:
    sampler:
      probability: 1.0  # 100% of requests traced (use 0.1 = 10% in prod)
```
In the Zipkin UI, you search by `TraceId` and see a timeline of every service call, with exact start times, durations, and any errors along the path.

**Step 3: Centralize logs with ELK Stack**
Each service forwards logs to **Elasticsearch** via **Logstash** (or directly). **Kibana** provides a UI to search across all logs from all services simultaneously using the shared `TraceId`."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Wipro, Accenture, Capgemini — common DevOps + Microservices integration question when the interviewer probes how you debug production issues across services.

#### Indepth
**Note on Spring Boot 3.x:** In Spring Boot 3+, Micrometer Tracing has replaced Spring Cloud Sleuth. The dependency is now `micrometer-tracing-bridge-otel` with an OTel exporter. The concepts are identical; only the library names changed.

---

## Q3. How do you secure inter-service communication in a Spring Boot microservices setup? How does JWT flow?

"Security must be applied at every layer in a microservices architecture. For most service-company projects, the standard approach involves JWT tokens validated at the API Gateway, then propagated downstream.

**The full JWT flow:**

```
1. User logs into Auth Service with username/password
2. Auth Service validates credentials → issues JWT token (signed with private key)
3. Client stores JWT (in memory or HttpOnly cookie) and sends it in every request:
   Authorization: Bearer eyJhbGc...

4. API Gateway intercepts the request
5. Gateway validates the JWT token (verify signature, expiry)
6. If valid: Extract user claims (userId, roles) and forward them as request headers:
   X-User-Id: 12345
   X-User-Role: ADMIN
7. Downstream microservices trust these headers (they only accept calls from inside the cluster)
   They DO NOT re-validate the JWT — the Gateway already did it.
```

**Spring Boot implementation:**
```java
// API Gateway or Security Config in each service
@Bean
public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
    http
        .authorizeHttpRequests(auth -> auth
            .requestMatchers("/api/public/**").permitAll()
            .anyRequest().authenticated()
        )
        .oauth2ResourceServer(oauth2 -> oauth2.jwt(Customizer.withDefaults()));
    return http.build();
}
```

**Role-based access:**
```java
@GetMapping("/admin/report")
@PreAuthorize("hasRole('ADMIN')")
public ResponseEntity<Report> getAdminReport() { ... }
```"

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS, Cognizant, Accenture — security is a standard interview topic. Interviewers want to know you understand the difference between stateless JWT and session-based authentication, and where validation happens in a microservices topology.

#### Indepth
**Service-to-Service Authentication:** The JWT flow above handles user-facing authentication. For internal service-to-service calls (where there's no human user), services use **OAuth2 Client Credentials flow** or **Mutual TLS (mTLS)**. With mTLS, each service has its own certificate, and calling services must present their certificate to be authenticated — common in Service Mesh (Istio) setups.

---

## Q4. How do you monitor Spring Boot microservices in production? What is Spring Boot Actuator?

"**Spring Boot Actuator** is a built-in module that exposes operational endpoints about your running application over HTTP, without you having to write any custom code.

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-actuator</artifactId>
</dependency>
```

```yaml
management:
  endpoints:
    web:
      exposure:
        include: health,info,metrics,prometheus,env,loggers
  endpoint:
    health:
      show-details: always
```

**Key Actuator endpoints:**
| Endpoint | Purpose |
|---|---|
| `/actuator/health` | Application health status (DB, Kafka, Redis connectivity checks built-in) |
| `/actuator/metrics` | JVM memory, thread counts, HTTP request rates, GC stats |
| `/actuator/prometheus` | Metrics in Prometheus scrape format — used by Grafana dashboards |
| `/actuator/loggers` | View and dynamically change log levels at runtime without restart |
| `/actuator/env` | View all active environment properties |
| `/actuator/info` | Custom app info (version, git commit hash) |

**Integration with Kubernetes health checks:**
```yaml
# K8s Deployment
livenessProbe:
  httpGet:
    path: /actuator/health/liveness
    port: 8080
readinessProbe:
  httpGet:
    path: /actuator/health/readiness
    port: 8080
```
Kubernetes will restart a pod if `/actuator/health/liveness` returns DOWN, and stop routing traffic to a pod if `/actuator/health/readiness` returns DOWN."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** This is a near-universal Spring Boot question at all service companies. If you've ever deployed a microservice to production, knowing Actuator is expected.

#### Indepth
**Custom Health Indicators:** You can write your own health check logic:
```java
@Component
public class ExternalApiHealthIndicator implements HealthIndicator {
    @Override
    public Health health() {
        boolean isApiUp = checkExternalPaymentGateway();
        return isApiUp ? Health.up().build() : Health.down().withDetail("reason", "Payment gateway unreachable").build();
    }
}
```
This appears automatically in `/actuator/health` output.

---

## Q5. How do you Dockerize and deploy a Spring Boot microservice to Kubernetes?

"This is the end-to-end deployment pipeline that most service company projects follow.

**Step 1: Write a Dockerfile**
```dockerfile
# Multi-stage build to keep the final image small
FROM eclipse-temurin:17-jdk-alpine AS build
WORKDIR /app
COPY . .
RUN ./mvnw package -DskipTests

FROM eclipse-temurin:17-jre-alpine
WORKDIR /app
COPY --from=build /app/target/order-service-0.0.1-SNAPSHOT.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
```

**Step 2: Build and push Docker image**
```bash
docker build -t myregistry.io/order-service:v1.2 .
docker push myregistry.io/order-service:v1.2
```

**Step 3: Write a Kubernetes Deployment and Service YAML**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: order-service
  template:
    metadata:
      labels:
        app: order-service
    spec:
      containers:
      - name: order-service
        image: myregistry.io/order-service:v1.2
        ports:
        - containerPort: 8080
        env:
        - name: SPRING_DATASOURCE_URL
          valueFrom:
            secretKeyRef:
              name: order-service-secrets
              key: db-url
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        readinessProbe:
          httpGet:
            path: /actuator/health/readiness
            port: 8080
          initialDelaySeconds: 15
---
apiVersion: v1
kind: Service
metadata:
  name: order-service
spec:
  selector:
    app: order-service
  ports:
  - port: 8080
    targetPort: 8080
```

**Step 4: Deploy**
```bash
kubectl apply -f deployment.yaml
kubectl rollout status deployment/order-service
```

**Sensitive configuration (DB passwords) must go in Kubernetes Secrets, not ConfigMaps.** Secrets are base64-encoded and can be encrypted at rest. Never hardcode credentials in Docker images or Git."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS, Infosys, Wipro — if a project says 'microservices', this deployment pipeline is what the interviewer expects you to have hands-on experience with.

#### Indepth
**ConfigMaps for non-sensitive config:** Use Kubernetes `ConfigMap` to inject non-sensitive application configuration (like `SPRING_PROFILES_ACTIVE=prod`, feature flags, connection pool sizes). Mount them as environment variables or as a volume (file). This is the Kubernetes-native alternative to Spring Cloud Config Server.
