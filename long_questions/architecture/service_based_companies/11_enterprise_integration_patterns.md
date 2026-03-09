# 🏢 Enterprise Integration Patterns

> **Focus:** Service-Based Companies (TCS, Infosys, Wipro, Cognizant, Accenture)
> **Level:** 🟡 Mid – 🔴 Senior

---

## 📋 Table of Contents

1. [Integration Fundamentals](#1-integration-fundamentals)
2. [Messaging Patterns](#2-messaging-patterns)
3. [Enterprise Service Bus (ESB)](#3-enterprise-service-bus-esb)
4. [API Gateway & Management](#4-api-gateway--management)
5. [Data Integration Patterns](#5-data-integration-patterns)
6. [Legacy Modernization](#6-legacy-modernization)
7. [Enterprise Security Integration](#7-enterprise-security-integration)
8. [Common Interview Questions](#8-common-interview-questions)

---

## 1. Integration Fundamentals

### Q1: What is enterprise integration and why is it needed?

**Answer:**
Enterprise integration connects disparate systems, applications, and data sources across an organization to enable seamless business processes and data flow.

**Key Drivers:**
- **Mergers & Acquisitions:** Integrating acquired company systems
- **Digital Transformation:** Modernizing legacy applications
- **Business Process Automation:** End-to-end workflow automation
- **Data Consistency:** Single source of truth across systems
- **Customer Experience:** Unified view across touchpoints

**Integration Challenges:**
- Heterogeneous technologies (mainframe, cloud, on-premise)
- Different data formats and protocols
- Varying reliability and performance requirements
- Security and compliance constraints
- Legacy system limitations

---

## 2. Messaging Patterns

### Q2: Explain Point-to-Point vs Publish-Subscribe messaging

**Answer:**

**Point-to-Point (P2P):**
```
Producer → Queue → Consumer
```
- **One-to-one** communication
- **Load balancing:** Multiple consumers compete for messages
- **Reliability:** Messages persist until consumed
- **Use Cases:** Order processing, task distribution

**Publish-Subscribe (Pub/Sub):**
```
Publisher → Topic → Multiple Subscribers
```
- **One-to-many** communication
- **Broadcasting:** Same message to multiple consumers
- **Decoupling:** Publishers don't know subscribers
- **Use Cases:** Event notifications, data synchronization

**Enterprise Examples:**
- **IBM MQ:** P2P for transaction processing
- **Apache Kafka:** Pub/Sub for event streaming
- **RabbitMQ:** Both patterns supported

### Q3: What are Message Transformation Patterns?

**Answer:**

**Content Enrichment:**
```java
// Enrich order with customer details
public Order enrichOrder(Order order, Customer customer) {
    order.setCustomerName(customer.getName());
    order.setCustomerTier(customer.getTier());
    return order;
}
```

**Content Filter:**
```java
// Filter sensitive data for external systems
public Order filterSensitiveData(Order order) {
    Order filtered = new Order();
    filtered.setId(order.getId());
    filtered.setAmount(order.getAmount());
    // Remove internal fields
    return filtered;
}
```

**Message Translator:**
```xml
<!-- XSLT transformation for XML -->
<xsl:transform version="1.0">
    <xsl:template match="/Order">
        <PurchaseOrder>
            <xsl:copy-of select="Amount"/>
            <xsl:copy-of select="Customer"/>
        </PurchaseOrder>
    </xsl:template>
</xsl:transform>
```

**Canonical Data Model:**
- **Standard format** for enterprise data exchange
- **Reduces transformation complexity**
- **Example:** FpML for financial data, HL7 for healthcare

---

## 3. Enterprise Service Bus (ESB)

### Q4: What is an ESB and when should you use it?

**Answer:**

**ESB Definition:**
Enterprise Service Bus is a middleware architecture that enables integration of disparate applications through a centralized communication bus.

**Key Capabilities:**
- **Message Routing:** Dynamic routing based on content/rules
- **Protocol Transformation:** SOAP ↔ REST ↔ JMS ↔ FTP
- **Message Orchestration:** Complex business process coordination
- **Security:** Centralized authentication and authorization
- **Monitoring:** Logging and tracking of message flows

**When to Use ESB:**
- **Multiple heterogeneous systems** need integration
- **Complex routing** and transformation requirements
- **Centralized governance** needed
- **Legacy modernization** projects

**Popular ESB Solutions:**
- **MuleSoft:** Cloud-based integration platform
- **IBM Integration Bus:** Enterprise-grade ESB
- **Apache ServiceMix:** Open-source ESB
- **WSO2 EI:** Open-source enterprise integrator

### Q5: Explain Service Orchestration vs Choreography

**Answer:**

**Service Orchestration:**
```
[Central Controller] → [Service A] → [Service B] → [Service C]
```
- **Centralized control:** One service orchestrates others
- **Explicit coordination:** Orchestrator defines the flow
- **Easier to manage:** Single point of control
- **Use Cases:** Business process workflows

**Service Choreography:**
```
[Service A] ↔ [Service B] ↔ [Service C] (Decentralized)
```
- **Decentralized control:** Each service knows its role
- **Event-driven:** Services react to events
- **More resilient:** No single point of failure
- **Use Cases:** Microservices, event-driven systems

**Implementation Example:**
```java
// Orchestration Pattern
@Service
public class OrderOrchestrator {
    
    public OrderResult processOrder(Order order) {
        // Centralized control
        InventoryService.reserve(order);
        PaymentService.charge(order);
        ShippingService.schedule(order);
        return new OrderResult("SUCCESS");
    }
}

// Choreography Pattern
@Service
public class OrderService {
    
    @EventHandler
    public void handleOrderCreated(OrderCreated event) {
        // React to events
        inventoryService.reserve(event.getOrder());
        // Emit next event
        eventBus.publish(new OrderReserved(event.getOrder()));
    }
}
```

---

## 4. API Gateway & Management

### Q6: What is the role of an API Gateway in enterprise integration?

**Answer:**

**API Gateway Responsibilities:**
- **Request Routing:** Route to appropriate backend services
- **Protocol Translation:** REST ↔ SOAP ↔ gRPC
- **Authentication & Authorization:** OAuth, JWT, API keys
- **Rate Limiting:** Prevent abuse and ensure fair usage
- **Caching:** Improve performance and reduce backend load
- **Logging & Monitoring:** Track API usage and performance

**Enterprise API Gateway Solutions:**
- **Apigee (Google):** Full API management platform
- **AWS API Gateway:** Cloud-native gateway
- **Kong:** Open-source API gateway
- **IBM API Connect:** Enterprise API management

### Q7: Explain API Versioning Strategies for Enterprise

**Answer:**

**URI Versioning:**
```
/api/v1/customers
/api/v2/customers
```
- **Pros:** Clear and explicit
- **Cons:** URL changes affect clients

**Header Versioning:**
```
GET /customers
Accept: application/vnd.company.v1+json
```
- **Pros:** Clean URLs, backward compatible
- **Cons:** Less discoverable

**Query Parameter Versioning:**
```
/customers?version=1
```
- **Pros:** Easy to implement
- **Cons:** Can be messy with many parameters

**Enterprise Best Practices:**
```java
@RestController
@RequestMapping("/api/v1")
public class CustomerControllerV1 {
    
    @GetMapping("/customers/{id}")
    public CustomerV1 getCustomer(@PathVariable Long id) {
        // V1 implementation
        return customerService.getCustomerV1(id);
    }
}

@RestController
@RequestMapping("/api/v2")
public class CustomerControllerV2 {
    
    @GetMapping("/customers/{id}")
    public CustomerV2 getCustomer(@PathVariable Long id) {
        // V2 implementation with additional fields
        return customerService.getCustomerV2(id);
    }
}
```

---

## 5. Data Integration Patterns

### Q8: How do you handle data synchronization between systems?

**Answer:**

**Batch Integration:**
```java
@Scheduled(cron = "0 0 2 * * ?") // Daily at 2 AM
public void syncCustomerData() {
    List<Customer> customers = crmService.getAllCustomers();
    erpService.bulkUpdate(customers);
}
```
- **Scheduled processing:** Nightly, hourly batches
- **High volume:** Efficient for large datasets
- **Eventual consistency:** Data may be temporarily out of sync

**Real-time Integration:**
```java
@EventListener
public void handleCustomerUpdate(CustomerUpdated event) {
    // Immediate propagation
    erpService.updateCustomer(event.getCustomer());
    analyticsService.trackCustomerChange(event);
}
```
- **Event-driven:** Immediate updates
- **Low latency:** Near real-time consistency
- **Higher complexity:** More infrastructure needed

**Change Data Capture (CDC):**
```java
@KafkaListener(topics = "database-changes")
public void handleDatabaseChange(DatabaseChangeEvent event) {
    if (event.getTable().equals("customers")) {
        syncCustomerToSystems(event.getData());
    }
}
```
- **Database-level:** Capture all changes automatically
- **Reliable:** No missed updates
- **Decoupled:** Source systems unaware of consumers

---

## 6. Legacy Modernization

### Q9: What are common legacy modernization patterns?

**Answer:**

**Strangler Fig Pattern:**
```java
// Gradually replace legacy functionality
@RestController
@RequestMapping("/api")
public class ModernizedController {
    
    @GetMapping("/orders/{id}")
    public Order getOrder(@PathVariable Long id) {
        // Route to new service if available
        if (modernizedOrderService.supports(id)) {
            return modernizedOrderService.getOrder(id);
        }
        // Fallback to legacy system
        return legacyOrderService.getOrder(id);
    }
}
```

**Anti-Corruption Layer:**
```java
@Service
public class LegacyAdapter {
    
    public ModernOrder getModernOrder(LegacyOrder legacy) {
        // Transform legacy format to modern
        ModernOrder modern = new ModernOrder();
        modern.setId(legacy.getOrderId());
        modern.setAmount(convertCurrency(legacy.getAmount()));
        return modern;
    }
}
```

**Lift and Shift:**
- **Rehost:** Move to cloud without changes
- **Replatform:** Minor changes for cloud optimization
- **Refactor:** Restructure for cloud-native patterns

### Q10: How do you ensure data consistency during modernization?

**Answer:**

**Dual Write Pattern:**
```java
@Service
public class OrderService {
    
    @Transactional
    public void createOrder(Order order) {
        // Write to both systems
        modernRepository.save(order);
        legacyRepository.save(convertToLegacy(order));
        
        // Verify consistency
        if (!validateConsistency(order)) {
            throw new ConsistencyException("Data mismatch");
        }
    }
}
```

**Event Sourcing for Migration:**
```java
@EventHandler
public void handleOrderEvent(OrderEvent event) {
    // Apply to both systems
    modernSystem.apply(event);
    legacySystem.apply(event);
    
    // Track migration progress
    migrationTracker.recordEvent(event);
}
```

**Reconciliation Jobs:**
```java
@Scheduled(fixedRate = 3600000) // Hourly
public void reconcileData() {
    List<String> discrepancies = findDiscrepancies();
    if (!discrepancies.isEmpty()) {
        alertService.notify(discrepancies);
        autoFixService.fix(discrepancies);
    }
}
```

---

## 7. Enterprise Security Integration

### Q11: How do you handle authentication across enterprise systems?

**Answer:**

**Single Sign-On (SSO) with SAML:**
```java
@RestController
public class SSOController {
    
    @GetMapping("/saml/login")
    public ResponseEntity<?> initiateSSO() {
        // Redirect to identity provider
        String samlRequest = samlService.buildAuthRequest();
        return ResponseEntity.redirect("https://idp.company.com/sso?" + samlRequest);
    }
    
    @PostMapping("/saml/acs")
    public ResponseEntity<?> handleSAMLResponse(@RequestBody String samlResponse) {
        // Process SAML response
        SAMLAssertion assertion = samlService.parseResponse(samlResponse);
        User user = userService.authenticate(assertion);
        return ResponseEntity.ok(createToken(user));
    }
}
```

**OAuth 2.0 for API Integration:**
```java
@Configuration
@EnableResourceServer
public class ResourceServerConfig extends ResourceServerConfigurerAdapter {
    
    @Override
    public void configure(HttpSecurity http) throws Exception {
        http.authorizeRequests()
            .antMatchers("/api/public/**").permitAll()
            .antMatchers("/api/enterprise/**").hasRole("ENTERPRISE_USER")
            .anyRequest().authenticated();
    }
}
```

**LDAP/Active Directory Integration:**
```java
@Service
public class LdapAuthenticationService {
    
    public Authentication authenticate(String username, String password) {
        DirContext ctx = null;
        try {
            ctx = ldapContextSource.getContext(username + "@company.com", password);
            // Authentication successful
            return new UsernamePasswordAuthenticationToken(username, password, 
                AuthorityUtils.createAuthorityList("ROLE_ENTERPRISE_USER"));
        } catch (Exception e) {
            throw new BadCredentialsException("LDAP authentication failed", e);
        } finally {
            LdapUtils.closeContext(ctx);
        }
    }
}
```

---

## 8. Common Interview Questions

### Q12: How would you integrate a legacy mainframe system with a modern web application?

**Answer:**
"I would use a **phased approach**:

1. **Analysis Phase:** Understand mainframe APIs, data formats, and business rules
2. **Adapter Layer:** Create a modern API wrapper around mainframe services
3. **Data Synchronization:** Implement CDC or batch processes for data consistency
4. **Gradual Migration:** Use Strangler Fig pattern to replace functionality incrementally
5. **Testing:** Comprehensive integration testing with both systems

**Technical Approach:**
- **IBM MQ or JCA** for mainframe connectivity
- **REST API gateway** to modernize the interface
- **Data transformation layer** for format conversion
- **Caching strategy** to reduce mainframe load"

### Q13: What considerations are important when integrating with third-party SaaS applications?

**Answer:**
"Key considerations include:

**API Limitations:**
- Rate limits and quotas
- Authentication methods (OAuth, API keys)
- Data format and pagination

**Reliability & Resilience:**
- Retry policies for failed requests
- Circuit breakers for service degradation
- Local caching for offline scenarios

**Security & Compliance:**
- Data residency requirements
- GDPR/privacy compliance
- Audit logging and monitoring

**Cost Management:**
- API call pricing models
- Data transfer costs
- Licensing considerations

**Implementation Example:**
```java
@Service
public class SaaSIntegrationService {
    
    @Retryable(maxAttempts = 3, backoff = @Backoff(delay = 1000))
    @CircuitBreaker(failureRateThreshold = 50)
    public Customer syncCustomerToSaaS(Customer customer) {
        try {
            return saasClient.createCustomer(customer);
        } catch (RateLimitException e) {
            // Implement exponential backoff
            throw new ServiceUnavailableException("Rate limit exceeded", e);
        }
    }
}
```"

### Q14: How do you handle message ordering in distributed enterprise systems?

**Answer:**
"Several approaches for message ordering:

**Single Partition Queue:**
```java
// Kafka single partition for ordering
@KafkaListener(topics = "ordered-events", groupId = "processor")
public void handleOrderedEvent(Event event) {
    // Guaranteed order within partition
    processInOrder(event);
}
```

**Sequence Numbers:**
```java
public class OrderedMessageProcessor {
    
    private final Map<String, Long> lastSequence = new ConcurrentHashMap<>();
    
    public void processMessage(String key, Message message) {
        long expectedSeq = lastSequence.getOrDefault(key, 0L) + 1;
        
        if (message.getSequence() == expectedSeq) {
            // Process in order
            handleMessage(message);
            lastSequence.put(key, expectedSeq);
        } else {
            // Buffer out-of-order messages
            bufferMessage(key, message);
        }
    }
}
```

**Database-Based Ordering:**
```sql
-- Use database sequence for ordering
CREATE TABLE message_queue (
    id BIGSERIAL PRIMARY KEY,
    payload JSONB,
    status VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Process in order
SELECT * FROM message_queue 
WHERE status = 'PENDING' 
ORDER BY id LIMIT 1 FOR UPDATE SKIP LOCKED;
```"

### Q15: What monitoring and observability practices do you implement for enterprise integrations?

**Answer:**
"Comprehensive observability includes:

**Logging Strategy:**
```java
@Component
public class IntegrationLogger {
    
    private static final Logger logger = LoggerFactory.getLogger(IntegrationLogger.class);
    
    public void logIntegration(Message message, String system, String operation) {
        logger.info("INTEGRATION: {} {} - MessageId: {}, Size: {} bytes", 
            system, operation, message.getId(), message.getSize());
    }
    
    public void logError(Exception e, Message message, String system) {
        logger.error("INTEGRATION_ERROR: {} - MessageId: {}, Error: {}", 
            system, message.getId(), e.getMessage(), e);
    }
}
```

**Metrics Collection:**
```java
@Component
public class IntegrationMetrics {
    
    private final MeterRegistry meterRegistry;
    
    public void recordMessageProcessed(String system, String status) {
        Counter.builder("integration.messages.processed")
            .tag("system", system)
            .tag("status", status)
            .register(meterRegistry)
            .increment();
    }
    
    public void recordProcessingTime(String system, Duration duration) {
        Timer.builder("integration.processing.time")
            .tag("system", system)
            .register(meterRegistry)
            .record(duration);
    }
}
```

**Distributed Tracing:**
```java
@RestController
public class IntegrationController {
    
    @NewSpan("process-integration")
    public ResponseEntity<?> processIntegration(@RequestBody Request request) {
        Span span = tracer.nextSpan().name("integration-flow");
        try (Tracer.SpanInScope ws = tracer.withSpanInScope(span)) {
            // Integration logic
            return ResponseEntity.ok(processRequest(request));
        } finally {
            span.end();
        }
    }
}
```

**Health Checks:**
```java
@Component
public class IntegrationHealthCheck implements HealthIndicator {
    
    @Override
    public Health health() {
        boolean allSystemsHealthy = checkAllSystems();
        
        if (allSystemsHealthy) {
            return Health.up()
                .withDetail("systems", getSystemStatus())
                .build();
        } else {
            return Health.down()
                .withDetail("failed_systems", getFailedSystems())
                .build();
        }
    }
}
```"

---

## 🎯 Quick Reference

### Essential Patterns for Service Companies
- **Point-to-Point:** Task distribution
- **Publish-Subscribe:** Event notification
- **Content Enrichment:** Data augmentation
- **Canonical Data Model:** Standardized format
- **API Gateway:** Centralized access
- **Strangler Fig:** Gradual modernization

### Common Technologies
- **IBM MQ:** Enterprise messaging
- **MuleSoft:** Integration platform
- **Apache Kafka:** Event streaming
- **RabbitMQ:** Message broker
- **Apigee:** API management
- **Spring Integration:** Framework-based integration

### Interview Focus Areas
- **Pattern recognition:** Identify appropriate patterns
- **Trade-off analysis:** When to use which approach
- **Enterprise context:** Understand business requirements
- **Security considerations:** Authentication, authorization
- **Monitoring:** Observability and troubleshooting
