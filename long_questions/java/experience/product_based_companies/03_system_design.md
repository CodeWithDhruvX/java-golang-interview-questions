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

### 🎯 How to Explain in Interview

"Rate limiting is crucial for protecting APIs from abuse and ensuring fair usage. I implement it using different algorithms - the Token Bucket pattern allows bursting while maintaining average rate, and works well for APIs. For single-instance applications, I can use Guava's RateLimiter which implements token bucket internally. For distributed systems, I use Redis with atomic increments to track request counts per time window. I implement this as a Spring filter that checks the rate limit before processing requests. The key is choosing the right algorithm and scope - per-client, per-IP, or global rate limits. Redis-based distributed rate limiting ensures consistency across multiple service instances, which is essential for microservices architectures."

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

### 🎯 How to Explain in Interview

"For in-memory caching in production, I use Caffeine cache which is the modern high-performance alternative to Guava. I configure it with maximum size limits, time-based eviction policies, and metrics recording. Spring's @Cacheable annotations make caching declarative - I annotate methods that should cache their results. I use @Cacheable for reads, @CacheEvict for deletions, and @CachePut for updates to keep the cache consistent. The unless parameter prevents caching null results. This approach dramatically improves performance for frequently accessed data while keeping the code clean. For distributed caching, I'd use Redis, but for local caching, Caffeine provides excellent performance with minimal overhead."

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

### 🎯 How to Explain in Interview

"Designing a URL shortener involves several key components. I need a service to generate short codes, a database to store mappings, and a cache for hot URLs. For the short code generation, I use base62 encoding of a counter or Twitter Snowflake for distributed ID generation. The database schema stores the short code as primary key with the long URL, metadata, and analytics. I cache frequently accessed URLs in Redis to avoid database hits. The system needs to handle high redirect volumes, so I use a CDN or edge caching for the redirect service. For scalability, I can shard the database by short code ranges. The key challenges are generating collision-free short codes and handling redirect performance at scale."

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

### 🎯 How to Explain in Interview

"CQRS is about separating the read and write models of an application. The command side handles writes with a normalized, transactional database optimized for consistency. The query side uses denormalized data structures optimized for fast reads. When a command creates or updates data, it publishes an event that updates the read model. This separation allows me to optimize each side independently - the write side for ACID compliance, the read side for performance. I might use a relational database for writes and Elasticsearch or a document database for reads. This pattern is especially valuable in complex domains with different read and write requirements, though it adds complexity. The key benefit is scalability - I can scale reads and writes independently and use different technologies for each."

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

### 🎯 How to Explain in Interview

"The CAP theorem states that in distributed systems, I can only guarantee two out of three properties: Consistency, Availability, and Partition Tolerance. Since network partitions are inevitable in real distributed systems, I must choose between consistency and availability during partitions. CP systems prioritize consistency - they may become unavailable to avoid serving stale data, like MySQL with replication or Kafka. AP systems prioritize availability - they may serve stale data to remain responsive, like Cassandra. Most systems need to make this trade-off based on their requirements. For financial systems, I'd choose CP for data safety. For social media feeds, I'd choose AP for user experience. Understanding this trade-off is crucial for designing distributed systems that meet specific business requirements."

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

---

### How to Explain in Interview

"Distributed locks are essential for coordinating operations across multiple service instances. I implement this using Redis with SET NX EX commands - the atomic set-if-not-exists with expiration. Each lock attempt generates a unique token, and I only release the lock if I still own the token, verified through a Lua script to ensure atomicity. The TTL prevents deadlocks if a service crashes. I use this for critical sections like processing orders to prevent concurrent modifications. The key challenges are ensuring the lock is truly distributed and handling failures gracefully. Redis distributed locks are widely used because they're fast, reliable, and work across different programming languages and frameworks."
