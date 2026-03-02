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
