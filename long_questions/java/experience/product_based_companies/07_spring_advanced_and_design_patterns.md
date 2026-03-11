# 🛡️ 07 — Spring Advanced & Design Patterns
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- Spring Security (JWT, OAuth2)
- Advanced AOP (custom annotations)
- Design patterns: Factory, Strategy, Observer, Decorator
- SOLID principles in Java
- Clean Architecture
- Event-driven with Spring Events

---

## ❓ Most Asked Questions

### Q1. How does Spring Security with JWT work?

```java
// JWT Filter — validates token on every request
@Component
public class JwtAuthFilter extends OncePerRequestFilter {

    private final JwtService jwtService;
    private final UserDetailsService userDetailsService;

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response,
                                    FilterChain chain) throws IOException, ServletException {
        String authHeader = request.getHeader("Authorization");
        if (authHeader == null || !authHeader.startsWith("Bearer ")) {
            chain.doFilter(request, response);
            return;
        }
        String token = authHeader.substring(7);
        String username = jwtService.extractUsername(token);

        if (username != null && SecurityContextHolder.getContext().getAuthentication() == null) {
            UserDetails userDetails = userDetailsService.loadUserByUsername(username);
            if (jwtService.isValid(token, userDetails)) {
                UsernamePasswordAuthenticationToken auth = new UsernamePasswordAuthenticationToken(
                    userDetails, null, userDetails.getAuthorities());
                auth.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));
                SecurityContextHolder.getContext().setAuthentication(auth);
            }
        }
        chain.doFilter(request, response);
    }
}

// Security config
@Configuration
@EnableWebSecurity
public class SecurityConfig {

    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http,
                                            JwtAuthFilter jwtFilter) throws Exception {
        return http
            .csrf(csrf -> csrf.disable())
            .sessionManagement(s -> s.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/api/auth/**").permitAll()
                .requestMatchers("/api/admin/**").hasRole("ADMIN")
                .anyRequest().authenticated())
            .addFilterBefore(jwtFilter, UsernamePasswordAuthenticationFilter.class)
            .build();
    }
}

// JWT Service
@Service
public class JwtService {
    private static final String SECRET = "your-256-bit-secret-key-here";

    public String generateToken(UserDetails userDetails) {
        return Jwts.builder()
            .subject(userDetails.getUsername())
            .claim("roles", userDetails.getAuthorities())
            .issuedAt(new Date())
            .expiration(new Date(System.currentTimeMillis() + 86_400_000))  // 24h
            .signWith(getSignKey())
            .compact();
    }

    public String extractUsername(String token) {
        return Jwts.parser().verifyWith(getSignKey()).build()
                   .parseSignedClaims(token).getPayload().getSubject();
    }

    public boolean isValid(String token, UserDetails userDetails) {
        return extractUsername(token).equals(userDetails.getUsername());
    }
}
```

---

### 🎯 How to Explain in Interview

"Spring Security with JWT creates a stateless authentication system. I implement a JWT filter that extracts tokens from the Authorization header, validates them, and sets the authentication in the SecurityContext. The filter runs before the main authentication filter. I generate JWTs with user details and roles, sign them with a secret key, and set an expiration time. The security configuration disables CSRF, sets session management to stateless, and protects endpoints based on roles. The beauty is that once authenticated, the user context is available throughout the application via SecurityContextHolder. This approach scales well for microservices since there's no server-side session state to manage."

---

### Q2. What is a custom Spring AOP annotation?

```java
// Custom annotation
@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
public @interface Audited {
    String action() default "";
}

// Aspect implementation
@Aspect
@Component
public class AuditAspect {

    private final AuditLogRepository auditRepo;

    @Around("@annotation(audited)")
    public Object audit(ProceedingJoinPoint pjp, Audited audited) throws Throwable {
        String username = SecurityContextHolder.getContext()
                                               .getAuthentication().getName();
        long start = System.currentTimeMillis();
        Object result = null;
        String status = "SUCCESS";
        try {
            result = pjp.proceed();
        } catch (Exception e) {
            status = "FAILED: " + e.getMessage();
            throw e;
        } finally {
            auditRepo.save(new AuditLog(
                username,
                audited.action().isBlank() ? pjp.getSignature().getName() : audited.action(),
                status,
                System.currentTimeMillis() - start,
                Instant.now()
            ));
        }
        return result;
    }
}

// Usage — just annotate methods!
@Service
public class PaymentService {

    @Audited(action = "PROCESS_PAYMENT")
    public PaymentResponse processPayment(PaymentRequest request) {
        // ... business logic
    }
}
```

---

### 🎯 How to Explain in Interview

"Custom Spring AOP annotations let me add cross-cutting concerns without cluttering business logic. I create an annotation like @Audited, then implement an @Aspect that intercepts methods annotated with it. Using @Around advice, I can execute code before and after the method, measure execution time, and handle exceptions. The aspect has access to method parameters and can even modify return values. This is perfect for audit logging, performance monitoring, or security checks. The key benefit is that business code stays clean - I just add @Audited and the cross-cutting concern is automatically applied. This demonstrates AOP's power for separating concerns in enterprise applications."

---

### Q3. What is the Strategy Pattern?

```java
// Strategy — define a family of algorithms, make them interchangeable

// Strategy interface
public interface DiscountStrategy {
    BigDecimal calculateDiscount(Order order);
}

// Concrete strategies
@Component("seasonalDiscount")
public class SeasonalDiscountStrategy implements DiscountStrategy {
    @Override
    public BigDecimal calculateDiscount(Order order) {
        return order.getTotal().multiply(new BigDecimal("0.10"));  // 10% off
    }
}

@Component("loyaltyDiscount")
public class LoyaltyDiscountStrategy implements DiscountStrategy {
    @Override
    public BigDecimal calculateDiscount(Order order) {
        int points = order.getCustomer().getLoyaltyPoints();
        return points > 1000 ? order.getTotal().multiply(new BigDecimal("0.15"))
                             : order.getTotal().multiply(new BigDecimal("0.05"));
    }
}

@Component("noDiscount")
public class NoDiscountStrategy implements DiscountStrategy {
    @Override
    public BigDecimal calculateDiscount(Order order) { return BigDecimal.ZERO; }
}

// Context — uses strategy
@Service
public class OrderPricingService {

    private final Map<String, DiscountStrategy> strategies;  // Spring injects all strategies!

    public OrderPricingService(Map<String, DiscountStrategy> strategies) {
        this.strategies = strategies;
    }

    public BigDecimal calculateFinalPrice(Order order, String promotionType) {
        DiscountStrategy strategy = strategies.getOrDefault(promotionType + "Discount",
                                                             strategies.get("noDiscount"));
        return order.getTotal().subtract(strategy.calculateDiscount(order));
    }
}
```

---

### 🎯 How to Explain in Interview

"The Strategy pattern defines a family of algorithms and makes them interchangeable. I define a strategy interface like DiscountStrategy, then create multiple implementations like SeasonalDiscountStrategy and LoyaltyDiscountStrategy. Spring automatically injects all strategies into a Map where the key is the bean name. The context class selects the appropriate strategy based on conditions. This pattern is perfect for business rules that vary - like different discount calculations or payment processing methods. The beauty is that I can add new strategies without modifying existing code, following the Open/Closed principle. It's much cleaner than complex if-else chains and makes the system more maintainable."

---

### Q4. What is the Factory Pattern in Spring?

```java
// Abstract product
public interface NotificationSender {
    void send(String recipient, String message);
}

// Concrete products
@Component("email")
public class EmailSender implements NotificationSender {
    @Override public void send(String to, String msg) {
        System.out.println("Email→" + to + ": " + msg);
    }
}

@Component("sms")
public class SmsSender implements NotificationSender {
    @Override public void send(String to, String msg) {
        System.out.println("SMS→" + to + ": " + msg);
    }
}

@Component("push")
public class PushSender implements NotificationSender {
    @Override public void send(String to, String msg) {
        System.out.println("Push→" + to + ": " + msg);
    }
}

// Factory — Spring's ApplicationContext as factory
@Component
public class NotificationFactory {
    private final Map<String, NotificationSender> senders;  // auto-injected map

    public NotificationFactory(Map<String, NotificationSender> senders) {
        this.senders = senders;
    }

    public NotificationSender getSender(String channel) {
        return Optional.ofNullable(senders.get(channel))
            .orElseThrow(() -> new IllegalArgumentException("Unknown channel: " + channel));
    }
}

// Usage
notificationFactory.getSender("email").send("user@example.com", "Your order shipped!");
notificationFactory.getSender("sms").send("+919876543210", "OTP: 123456");
```

---

### 🎯 How to Explain in Interview

"The Factory pattern in Spring is beautifully implemented through dependency injection. I define an interface like NotificationSender and create multiple implementations with different @Component names. Spring automatically injects all implementations into a Map where the key is the bean name. The factory class simply looks up the appropriate implementation from this map. This is much cleaner than manual factory classes with if-else statements. The beauty is that adding new notification types is as simple as creating a new @Component class - no factory code changes needed. This leverages Spring's IoC container to implement the Factory pattern naturally."

---

### Q5. What are SOLID principles?

```java
// S — Single Responsibility: one reason to change
// BAD:
class UserService {
    void registerUser(User u) { /* save user } /* send email */ /* log event */ }
}
// GOOD:
class UserService     { void registerUser(User u) { userRepo.save(u); eventPublisher.publish(new UserRegistered(u)); } }
class EmailService    { void sendWelcome(User u) { /* send email */ } }
class UserEventListener { @EventListener void on(UserRegistered e) { emailService.sendWelcome(e.user()); } }

// O — Open/Closed: open for extension, closed for modification
// Use Strategy/Decorator instead of adding if-else in existing classes

// L — Liskov Substitution: subclasses should be substitutable for parent
// BAD: Square extends Rectangle — setWidth breaks invariant

// I — Interface Segregation: many small interfaces > one fat interface
interface Readable  { String read(); }
interface Writable  { void write(String data); }
interface ReadWrite extends Readable, Writable {}  // client uses only what they need

// D — Dependency Inversion: depend on abstractions, not concretions
// BAD:  private MySQLRepository db = new MySQLRepository();
// GOOD: private final UserRepository userRepo;  // injected via interface
```

---

### 🎯 How to Explain in Interview

"SOLID principles are fundamental guidelines for writing maintainable object-oriented code. Single Responsibility means each class has one reason to change - I separate user registration from email sending. Open/Closed means I design for extension without modification - I use Strategy pattern instead of adding if-else. Liskov Substitution ensures subclasses can replace their parents without breaking functionality. Interface Segregation means I create focused interfaces rather than fat ones - clients only depend on methods they actually use. Dependency Inversion means I depend on abstractions, not concretions - I inject interfaces through constructors. These principles help me create flexible, maintainable code that's easy to test and extend."

---

---

### Q6. What is the Observer Pattern with Spring Events?

```java
// Event — simple POJO (or extend ApplicationEvent for pre-Java 4.2 compatibility)
public record OrderShippedEvent(Long orderId, String trackingNumber, Instant shippedAt) {}

// Publisher — publishes events
@Service
public class ShippingService {

    private final ApplicationEventPublisher eventPublisher;

    @Transactional
    public TrackingInfo shipOrder(Long orderId) {
        TrackingInfo tracking = courier.createShipment(orderId);

        // Publish async event — listeners run after transaction commits
        eventPublisher.publishEvent(new OrderShippedEvent(
            orderId, tracking.getNumber(), Instant.now()));
        return tracking;
    }
}

// Synchronous listener (same thread, same transaction)
@EventListener
public void onOrderShipped(OrderShippedEvent event) {
    notificationService.sendShippingConfirmation(event.orderId(), event.trackingNumber());
}

// Asynchronous listener (different thread)
@EventListener
@Async("notificationPool")
public void sendPushNotification(OrderShippedEvent event) {
    pushService.notify(event.orderId(), "Your order is on the way! 🚚");
}

// Transactional listener — runs AFTER commit (safe for external systems)
@TransactionalEventListener(phase = TransactionPhase.AFTER_COMMIT)
public void publishToKafka(OrderShippedEvent event) {
    kafkaTemplate.send("order-shipped", event.orderId().toString(), event);
}
```

---

### 🎯 How to Explain in Interview

"Spring's event system implements the Observer pattern beautifully. I create simple POJO events and publish them using ApplicationEventPublisher. Listeners use @EventListener to handle events. I can control when listeners run - synchronous listeners run in the same thread, @Async listeners run in different threads, and @TransactionalEventListener runs after transaction commits. This is perfect for decoupling components - the shipping service doesn't need to know about notification or Kafka services. It just publishes an event and interested parties react. This pattern is great for audit logging, notifications, or updating read models in CQRS architectures. Spring handles all the complexity of event routing and listener management."

---
