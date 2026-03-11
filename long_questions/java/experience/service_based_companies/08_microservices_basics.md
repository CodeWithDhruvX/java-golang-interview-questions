# 🔗 08 — Microservices Basics
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Microservices vs Monolith
- Spring Cloud components
- Service discovery (Eureka)
- API Gateway (Spring Cloud Gateway)
- Config Server
- Circuit Breaker (Resilience4j)
- Inter-service communication (Feign, RestTemplate)

---

## ❓ Most Asked Questions

### Q1. What are microservices and when to use them?

```text
MONOLITH                    MICROSERVICES
┌─────────────────┐         ┌──────────┐  ┌──────────┐  ┌──────────┐
│                 │         │  Order   │  │  User    │  │ Payment  │
│  Order  User    │         │ Service  │  │ Service  │  │ Service  │
│  Payment Catalog│   ──►   │  :8081   │  │  :8082   │  │  :8083   │
│  Notification   │         └──────────┘  └──────────┘  └──────────┘
│                 │              ↕              ↕              ↕
└─────────────────┘         ┌──────────────────────────────────────┐
      1 deployment           │         Message Bus (Kafka)          │
      1 DB                   └──────────────────────────────────────┘
```

| Aspect | Monolith | Microservices |
|--------|----------|---------------|
| Deployment | Single unit | Independent per service |
| Scaling | Scale all or nothing | Scale individual services |
| Technology | One stack | Polyglot (different languages OK) |
| Team structure | One large team | Small teams per service |
| Complexity | Low/Middle | High (distributed systems) |
| Fault isolation | One failure = all down | Failure isolated to service |

---

### 🎯 How to Explain in Interview

"Microservices is an architectural pattern where I break down a monolithic application into small, independent services. Each service has its own database, can be deployed independently, and communicates with other services through APIs or messaging. The key benefits are independent scaling - I can scale just the order service if it's under load, technology diversity - different services can use different tech stacks, and fault isolation - if one service fails, others continue working. The trade-offs are increased complexity - I need to handle service discovery, distributed transactions, and monitoring. I choose microservices when I have a large application with clear domain boundaries, different scaling requirements per component, and multiple teams working on different parts. For small applications or single teams, a monolith is often simpler and more efficient."

---

### Q2. What is Service Discovery with Eureka?

```yaml
# Eureka Server — pom.xml: spring-cloud-starter-netflix-eureka-server
# application.yml (Eureka Server)
server.port: 8761
eureka:
  client:
    register-with-eureka: false  # server doesn't register itself
    fetch-registry: false
```

```java
// Eureka Server
@SpringBootApplication
@EnableEurekaServer
public class DiscoveryServerApplication { }

// Eureka Client (each microservice)
// application.yml
// spring.application.name: order-service
// eureka.client.service-url.defaultZone: http://localhost:8761/eureka/
@SpringBootApplication
@EnableDiscoveryClient   // or @EnableEurekaClient
public class OrderServiceApplication { }

// Load-balanced RestTemplate — uses service name instead of URL
@Configuration
public class RestConfig {
    @Bean
    @LoadBalanced   // intercepts and resolves "PAYMENT-SERVICE" to actual host
    public RestTemplate restTemplate() { return new RestTemplate(); }
}

// Usage — no hardcoded URLs!
PaymentResponse response = restTemplate.postForObject(
    "http://PAYMENT-SERVICE/api/payments",  // Eureka resolves this
    paymentRequest, PaymentResponse.class);
```

---

### 🎯 How to Explain in Interview

"Service discovery is crucial in microservices because service instances change dynamically - they can scale up/down, fail, or move to different servers. Eureka solves this by acting as a service registry. Each microservice registers itself with Eureka on startup, providing its hostname and port. When one service needs to call another, it asks Eureka for the current instances instead of using hardcoded URLs. I use @EnableEurekaServer for the registry and @EnableDiscoveryClient for services. With @LoadBalanced RestTemplate, I can call services by name like 'http://PAYMENT-SERVICE' and Spring automatically resolves it to an actual instance. This makes my system resilient to changes and enables load balancing across multiple instances of the same service."

---

### Q3. What is an API Gateway?

```yaml
# Spring Cloud Gateway — routes + filters
spring:
  cloud:
    gateway:
      routes:
        - id: order-service
          uri: lb://ORDER-SERVICE       # lb:// = load-balanced via Eureka
          predicates:
            - Path=/api/orders/**
          filters:
            - StripPrefix=1             # remove /api prefix before forwarding

        - id: user-service
          uri: lb://USER-SERVICE
          predicates:
            - Path=/api/users/**
            - Method=GET,POST

        - id: rate-limit
          uri: lb://PRODUCT-SERVICE
          predicates:
            - Path=/api/products/**
          filters:
            - name: RequestRateLimiter
              args:
                redis-rate-limiter.replenishRate: 10      # 10 requests/sec
                redis-rate-limiter.burstCapacity: 20
```

```java
// Custom global filter — JWT validation, logging, request ID injection
@Component
public class AuthFilter implements GlobalFilter, Ordered {

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        String token = exchange.getRequest().getHeaders()
                               .getFirst(HttpHeaders.AUTHORIZATION);
        if (token == null || !validateToken(token)) {
            exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
            return exchange.getResponse().setComplete();
        }
        // Add user info to downstream header
        ServerHttpRequest mutated = exchange.getRequest().mutate()
            .header("X-User-Id", extractUserId(token))
            .build();
        return chain.filter(exchange.mutate().request(mutated).build());
    }

    @Override public int getOrder() { return -1; }  // run first
}
```

---

### 🎯 How to Explain in Interview

"An API Gateway is the front door to my microservices architecture. Instead of clients calling each service directly, they go through the gateway which handles routing, security, and cross-cutting concerns. Spring Cloud Gateway lets me define routes based on paths, methods, or headers, and apply filters for authentication, rate limiting, or request transformation. The gateway can route to load-balanced service instances using 'lb://SERVICE-NAME', and I can implement custom filters for JWT validation or adding correlation IDs. This centralizes concerns like security and monitoring, reduces complexity for clients, and provides a single entry point to my system. It's especially useful for exposing APIs to external consumers while keeping internal services protected."

---

### Q4. What is Circuit Breaker (Resilience4j)?

```java
// PROBLEM: Service A calls Service B; if B is slow/down, A threads pile up
// SOLUTION: Circuit breaker opens after N failures, fast-fails, tries again after timeout

@Service
public class PaymentServiceClient {

    @CircuitBreaker(name = "payment-service", fallbackMethod = "paymentFallback")
    @Retry(name = "payment-retry")
    @TimeLimiter(name = "payment-timeout")
    public CompletableFuture<PaymentResponse> processPayment(PaymentRequest request) {
        return CompletableFuture.supplyAsync(() ->
            restTemplate.postForObject(PAYMENT_URL, request, PaymentResponse.class));
    }

    // Fallback — called when circuit is open or all retries exhausted
    public CompletableFuture<PaymentResponse> paymentFallback(PaymentRequest request, Exception e) {
        log.warn("Payment service unavailable, using fallback: {}", e.getMessage());
        return CompletableFuture.completedFuture(PaymentResponse.queued(request.getOrderId()));
    }
}
```

```yaml
# application.yml — Resilience4j config
resilience4j:
  circuitbreaker:
    instances:
      payment-service:
        sliding-window-size: 10          # last 10 calls
        failure-rate-threshold: 50       # open if ≥50% fail
        wait-duration-in-open-state: 30s # wait before half-open
        permitted-calls-in-half-open-state: 3
  retry:
    instances:
      payment-retry:
        max-attempts: 3
        wait-duration: 500ms
        retry-exceptions:
          - java.io.IOException
          - java.util.concurrent.TimeoutException
  timelimiter:
    instances:
      payment-timeout:
        timeout-duration: 3s
```

---

### 🎯 How to Explain in Interview

"Circuit breakers are essential for building resilient microservices. The problem is that when one service becomes slow or fails, services that call it can get overwhelmed waiting for responses, causing a cascading failure. A circuit breaker monitors calls to external services - if too many fail, it 'opens' the circuit and immediately fails fast with a fallback response instead of waiting. After a timeout, it tries again in 'half-open' state. I use Resilience4j with annotations like @CircuitBreaker, @Retry, and @TimeLimiter. I can configure failure rate thresholds, wait times, and fallback methods. This pattern prevents cascading failures, provides graceful degradation, and allows the system to recover automatically when downstream services come back online."

---

### Q5. What is Feign Client?

```java
// Declarative HTTP client — no boilerplate RestTemplate code
@FeignClient(name = "inventory-service", fallback = InventoryFallback.class)
public interface InventoryClient {

    @GetMapping("/api/inventory/{productId}")
    InventoryResponse checkStock(@PathVariable Long productId);

    @PostMapping("/api/inventory/reserve")
    ReservationResponse reserve(@RequestBody ReservationRequest request);

    @DeleteMapping("/api/inventory/reserve/{reservationId}")
    void cancelReservation(@PathVariable String reservationId);
}

// Fallback implementation
@Component
public class InventoryFallback implements InventoryClient {
    @Override
    public InventoryResponse checkStock(Long productId) {
        return InventoryResponse.unknown(productId);  // safe default
    }

    @Override
    public ReservationResponse reserve(ReservationRequest request) {
        throw new ServiceUnavailableException("Inventory service is down");
    }

    @Override
    public void cancelReservation(String reservationId) { /* silent fail */ }
}

// Enable Feign in main class
@SpringBootApplication
@EnableFeignClients
public class OrderServiceApplication { }
```

---

### 🎯 How to Explain in Interview

"Feign Client is a declarative HTTP client that makes calling other microservices much cleaner than using RestTemplate. Instead of writing boilerplate code with URLs and parameter mapping, I just define an interface with annotations like @GetMapping and @PostMapping. Spring automatically generates the implementation, handling service discovery, load balancing, and serialization. I can also integrate fallbacks for resilience - if the called service is down, Feign automatically calls my fallback implementation. The beauty is that my service calls look like regular method calls rather than HTTP requests, making the code much more readable and type-safe. It's especially useful when I have many inter-service communications in a complex microservices architecture."

---

### Q6. What is Spring Cloud Config Server?

```yaml
# Config Server — pom: spring-cloud-config-server
# application.yml (Config Server)
server.port: 8888
spring:
  cloud:
    config:
      server:
        git:
          uri: https://github.com/myorg/config-repo
          default-label: main
          search-paths: '{application}'   # looks in /order-service/ folder
```

```java
@SpringBootApplication
@EnableConfigServer
public class ConfigServerApplication { }

// Client services — bootstrap.yml (loaded before application.yml)
// spring.application.name: order-service
// spring.config.import: optional:configserver:http://localhost:8888

// Refresh config at runtime without restart
@RefreshScope   // beans with this annotation reload on /actuator/refresh
@RestController
public class OrderController {
    @Value("${order.max-items:100}")
    private int maxItems;   // refreshed on POST /actuator/refresh
}
```

---

### 🎯 How to Explain in Interview

"Spring Cloud Config Server centralizes configuration management across microservices. Instead of each service having its own properties files, I store all configuration in a Git repository and the Config Server serves it. Each microservice connects to the Config Server on startup and loads its configuration. The beauty is that I can update configuration without redeploying - I just push changes to Git and call the refresh endpoint. I use @RefreshScope on beans that should reload when configuration changes. This approach makes it easy to manage environment-specific configs, feature flags, and runtime configuration changes. It's especially valuable in production where I need to tweak settings without downtime."

---

### Q7. What are distributed tracing and correlation IDs?

```yaml
# Micrometer Tracing + Zipkin
management:
  tracing:
    sampling:
      probability: 1.0   # trace 100% of requests
  zipkin:
    tracing:
      endpoint: http://localhost:9411/api/v2/spans
```

```java
// Correlation ID propagation across services
@Component
public class CorrelationIdFilter implements Filter {

    private static final String CORRELATION_ID_HEADER = "X-Correlation-ID";

    @Override
    public void doFilter(ServletRequest req, ServletResponse res, FilterChain chain)
            throws IOException, ServletException {
        HttpServletRequest request = (HttpServletRequest) req;
        String correlationId = request.getHeader(CORRELATION_ID_HEADER);
        if (correlationId == null) {
            correlationId = UUID.randomUUID().toString();
        }
        MDC.put("correlationId", correlationId);  // add to log context
        ((HttpServletResponse) res).setHeader(CORRELATION_ID_HEADER, correlationId);
        try {
            chain.doFilter(req, res);
        } finally {
            MDC.remove("correlationId");
        }
    }
}

// logback-spring.xml — include correlationId in every log line
// %d{yyyy-MM-dd HH:mm:ss} [%X{correlationId}] [%thread] %-5level %logger - %msg%n
```

---

### 🎯 How to Explain in Interview

"Distributed tracing is essential for debugging requests that flow through multiple microservices. When a user request hits my API gateway, goes through order service, then payment service, and finally notification service, I need to track the entire journey. I use correlation IDs to tag each request with a unique identifier that gets passed between services. This lets me trace a single request across all log files. I also integrate with tools like Zipkin or Jaeger using Micrometer Tracing - these automatically collect timing data and visualize the call chain. The correlation ID gets added to MDC (Mapped Diagnostic Context) so it appears in every log message. This approach makes debugging distributed systems much easier - I can see exactly where a request failed and how long each service took."
