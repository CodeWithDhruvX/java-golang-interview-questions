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
