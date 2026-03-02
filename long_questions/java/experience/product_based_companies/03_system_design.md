# 🏗️ 03 — System Design with Java
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Rate limiting patterns (Token Bucket, Leaky Bucket)
- Distributed caching (Redis + Spring Cache)
- HLD for common systems (URL shortener, notification service)
- Scalability patterns (CQRS, Event Sourcing)
- Database sharding and replication
- CAP theorem and eventual consistency

---

## ❓ Most Asked Questions

### Q1. How do you design a Rate Limiter?

```java
// Token Bucket — allows bursting, common in APIs
@Component
public class TokenBucketRateLimiter {

    // Using Guava's RateLimiter (token bucket internally)
    private final Map<String, RateLimiter> limiters = new ConcurrentHashMap<>();

    public boolean tryAcquire(String clientId, double requestsPerSecond) {
        RateLimiter limiter = limiters.computeIfAbsent(clientId,
            k -> RateLimiter.create(requestsPerSecond));
        return limiter.tryAcquire();  // non-blocking: true if token available
    }
}

// Redis-based distributed rate limiter (works across multiple instances)
@Component
public class RedisRateLimiter {

    private final StringRedisTemplate redis;

    public boolean isAllowed(String key, int limit, int windowSeconds) {
        String redisKey = "rate:" + key;
        Long count = redis.opsForValue().increment(redisKey);
        if (count == 1) {
            redis.expire(redisKey, Duration.ofSeconds(windowSeconds));
        }
        return count <= limit;
    }
}

// Spring filter for rate limiting
@Component
public class RateLimitFilter extends OncePerRequestFilter {

    @Override
    protected void doFilterInternal(HttpServletRequest req, HttpServletResponse res,
                                    FilterChain chain) throws IOException, ServletException {
        String clientId = req.getHeader("X-API-Key");
        if (clientId != null && !rateLimiter.tryAcquire(clientId, 10.0)) {
            res.setStatus(429);  // Too Many Requests
            res.getWriter().write("{\"error\": \"Rate limit exceeded\"}");
            return;
        }
        chain.doFilter(req, res);
    }
}
```

---

### Q2. How do you implement an in-memory cache? (LRU + TTL)

```java
// Production: Use Caffeine cache (Java alternative to Guava)
@Configuration
public class CacheConfig {

    @Bean
    public CacheManager cacheManager() {
        CaffeineCacheManager manager = new CaffeineCacheManager();
        manager.setCaffeine(Caffeine.newBuilder()
            .maximumSize(10_000)                // evict when > 10K entries
            .expireAfterWrite(10, TimeUnit.MINUTES)  // TTL
            .expireAfterAccess(5, TimeUnit.MINUTES)  // evict if not accessed for 5m
            .recordStats());                    // expose hit/miss metrics
        return manager;
    }
}

// Usage via @Cacheable annotations
@Service
public class ProductService {

    @Cacheable(value = "products", key = "#id",
               unless = "#result == null")      // don't cache null results
    public ProductDTO getProduct(Long id) {
        return productRepository.findById(id)
                                .map(productMapper::toDTO)
                                .orElseThrow();
    }

    @CacheEvict(value = "products", key = "#id")   // remove from cache
    public void deleteProduct(Long id) {
        productRepository.deleteById(id);
    }

    @CachePut(value = "products", key = "#result.id")  // update cache
    public ProductDTO updateProduct(Long id, ProductDTO dto) {
        return productMapper.toDTO(productRepository.save(productMapper.toEntity(dto)));
    }
}
```

---

### Q3. How do you design a URL Shortener?

```text
SYSTEM DESIGN — URL Shortener (bit.ly)

Requirements:
- POST /shorten  { longUrl }  → shortUrl (e.g., sht.ly/abc123)
- GET  /{code}               → 301 redirect to longUrl

High-Level Design:
Client → API Gateway → URL Service → Database (MySQL for metadata)
                                   → Redis (hot URL cache)
                                   → ID Generator (Twitter Snowflake)

Database Schema:
┌──────────────────────────────────────────────┐
│              url_mappings                     │
├────────────┬─────────────────────────────────┤
│ short_code │ VARCHAR(10) PRIMARY KEY          │
│ long_url   │ TEXT NOT NULL                    │
│ user_id    │ BIGINT (nullable)                │
│ created_at │ TIMESTAMP                        │
│ expires_at │ TIMESTAMP (nullable)             │
│ clicks     │ BIGINT DEFAULT 0                 │
└──────────────────────────────────────────────┘
```

```java
@Service
public class UrlShortenerService {

    @Autowired private UrlRepository urlRepository;
    @Autowired private RedisTemplate<String, String> redis;

    private static final String ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    private final AtomicLong counter = new AtomicLong(1_000_000);  // or distributed Snowflake

    public String shorten(String longUrl) {
        // Check if already shortened
        String existingCode = urlRepository.findByLongUrl(longUrl)
            .map(UrlMapping::getShortCode).orElse(null);
        if (existingCode != null) return "https://sht.ly/" + existingCode;

        String code = toBase62(counter.incrementAndGet());
        urlRepository.save(new UrlMapping(code, longUrl));
        redis.opsForValue().set("url:" + code, longUrl, 24, TimeUnit.HOURS);  // cache
        return "https://sht.ly/" + code;
    }

    public String resolve(String code) {
        String cached = redis.opsForValue().get("url:" + code);
        if (cached != null) return cached;
        return urlRepository.findByShortCode(code)
                            .orElseThrow(() -> new ResourceNotFoundException("URL not found"))
                            .getLongUrl();
    }

    private String toBase62(long num) {
        StringBuilder sb = new StringBuilder();
        while (num > 0) {
            sb.append(ALPHABET.charAt((int)(num % 62)));
            num /= 62;
        }
        return sb.reverse().toString();
    }
}
```

---

### Q4. What is CQRS (Command Query Responsibility Segregation)?

```java
// CQRS — separate READ models from WRITE models

// Write side (Command) — transactional, normalized DB
@Service
public class OrderCommandService {

    @Transactional
    public OrderId createOrder(CreateOrderCommand cmd) {
        Order order = new Order(cmd.getCustomerId(), cmd.getItems());
        order = orderRepository.save(order);

        // Publish event for read side to consume
        eventPublisher.publish(new OrderCreatedEvent(order.getId(), order.getTotal()));
        return order.getId();
    }
}

// Read side (Query) — denormalized, optimized for reads (different DB/table)
@Service
public class OrderQueryService {

    // Reads from pre-built projection (e.g., Elasticsearch or denormalized view)
    public OrderSummaryPage getOrdersByCustomer(Long customerId, Pageable pageable) {
        return orderProjectionRepository.findByCustomerId(customerId, pageable);
    }
}

// Event handler updates the read model
@EventListener
public void on(OrderCreatedEvent event) {
    orderProjectionRepository.save(
        new OrderProjection(event.getOrderId(), event.getTotal(), "PENDING"));
}

// REST controller routes to correct service
@PostMapping("/orders")
public ResponseEntity<OrderId> createOrder(@RequestBody CreateOrderCommand cmd) {
    return ResponseEntity.ok(orderCommandService.createOrder(cmd));  // COMMAND
}

@GetMapping("/orders")
public Page<OrderSummary> getOrders(Long customerId, Pageable p) {
    return orderQueryService.getOrdersByCustomer(customerId, p);     // QUERY
}
```

---

### Q5. What is CAP Theorem?

```text
CAP THEOREM: A distributed system can guarantee at most 2 of 3:

C — Consistency:    All nodes see the same data at the same time
A — Availability:  Every request gets a response (success or failure)
P — Partition Tol: System continues despite network partition

During a network partition (P is always needed in real world):
↓
Choose: C (consistent but may be unavailable) OR A (available but may be stale)

REAL SYSTEMS:
System            | Choice | Reason
MySQL (single)    | CA     | No partition tolerance (not distributed)
MySQL (InnoDB+rep)| CP     | Prioritizes consistency, replica may lag
MongoDB (default) | CP     | Strong consistency with majority write concern
MongoDB (w:0)     | AP     | Fire-and-forget writes — may lose data
Apache Cassandra  | AP     | High availability, eventual consistency
Redis Cluster     | CP     | Prefers data safety over availability
Kafka             | CP     | Replication with leader election

BASE (alternative to ACID for distributed):
B - Basically Available
S - Soft state (data may change even without input)
E - Eventually consistent (will become consistent over time)
```

---

### Q6. Design a Distributed Lock with Redis

```java
// Distributed lock — only one service instance holds lock at a time
@Component
public class RedisDistributedLock {

    private final StringRedisTemplate redis;
    private static final String LOCK_PREFIX = "lock:";
    private static final Duration DEFAULT_TTL = Duration.ofSeconds(30);

    // Returns lock token if acquired, empty if not
    public Optional<String> tryLock(String resource, Duration ttl) {
        String lockKey = LOCK_PREFIX + resource;
        String token = UUID.randomUUID().toString();
        Boolean acquired = redis.opsForValue()
            .setIfAbsent(lockKey, token, ttl != null ? ttl : DEFAULT_TTL);
        return Boolean.TRUE.equals(acquired) ? Optional.of(token) : Optional.empty();
    }

    // Release only if we still own the lock (atomic via Lua script)
    public boolean release(String resource, String token) {
        String script =
            "if redis.call('get', KEYS[1]) == ARGV[1] then " +
            "  return redis.call('del', KEYS[1]) " +
            "else return 0 end";
        Long result = redis.execute(
            new DefaultRedisScript<>(script, Long.class),
            List.of(LOCK_PREFIX + resource), token);
        return Long.valueOf(1).equals(result);
    }
}

// Usage — critical section
public void processOrder(Long orderId) {
    Optional<String> token = lock.tryLock("order:" + orderId, Duration.ofSeconds(10));
    if (token.isEmpty()) throw new ConcurrentModificationException("Order locked");
    try {
        doWork(orderId);
    } finally {
        lock.release("order:" + orderId, token.get());
    }
}
```
