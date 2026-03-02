# 🔀 06 — Microservices & Kafka Advanced
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- Kafka internals (replication, ISR, leader election)
- Saga pattern (Choreography vs Orchestration)
- Event Sourcing
- CQRS with Kafka streams
- Service mesh concepts
- Kafka Streams and KTable

---

## ❓ Most Asked Questions

### Q1. Kafka Internals — Replication and Leader Election

```text
KAFKA REPLICATION:
Topic "orders" — 3 partitions, replication factor 3

Partition 0:  Leader=Broker1  Replicas=[Broker1, Broker2, Broker3]
              ISR (In-Sync Replicas)=[Broker1, Broker2, Broker3]

WRITE PATH:
1. Producer sends to Broker1 (Partition 0 Leader)
2. Broker1 writes to local log
3. Broker2 and Broker3 fetch and replicate
4. Acks=all → Broker1 responds ONLY after all ISR replicas acknowledge

READ PATH:
- Only Leader serves reads (by default)
- Kafka 2.4+: Followers can serve reads with --replica-fetch-max-bytes

ISR MECHANISM:
- Follower falls behind (replica.lag.time.max.ms = 10s) → removed from ISR
- If only Leader remains in ISR and Leader fails → unclean.leader.election.enable?
  - true (default off): elect non-ISR leader — potential data loss!
  - false: partition unavailable until ISR leader recovers

LEADER ELECTION:
- All partition metadata stored in ZooKeeper/KRaft
- Kafka Controller (one broker elected) handles leader elections
- On broker failure: Controller assigns new leader from ISR within seconds
```

---

### Q2. What is the Saga Pattern?

```text
PROBLEM: Distributed transactions across microservices
       (No 2PC in microservices — too coupled, too slow)

SOLUTION: Saga — series of local transactions, each publishes events

CHOREOGRAPHY SAGA (event-driven):
OrderService ──OrderCreated──► PaymentService ──PaymentCompleted──► InventoryService
                                     │                                     │
                              PaymentFailed                         StockReserved
                                     │                                     │
                              OrderCancelled                        OrderConfirmed

Pros: No central coordinator, simple
Cons: Hard to trace, complex rollback logic scattered across services

ORCHESTRATION SAGA (centralized):
                         ┌────────────────────────┐
                         │    Order Saga           │
                         │    Orchestrator         │
                         └──┬────────┬────────┬───┘
                    command │  command│ command│
                            ▼         ▼        ▼
                      Payment    Inventory   Shipping
                      Service    Service     Service
```

```java
// Orchestration Saga with Axon Framework (or manual state machine)
@Saga
public class CreateOrderSaga {
    @Autowired private transient CommandGateway commandGateway;

    @StartSaga
    @SagaEventHandler(associationProperty = "orderId")
    public void on(OrderCreatedEvent event) {
        SagaLifecycle.associateWith("orderId", event.getOrderId().toString());
        commandGateway.send(new ReserveStockCommand(event.getOrderId(), event.getItems()));
    }

    @SagaEventHandler(associationProperty = "orderId")
    public void on(StockReservedEvent event) {
        commandGateway.send(new ProcessPaymentCommand(event.getOrderId(), event.getTotal()));
    }

    @SagaEventHandler(associationProperty = "orderId")
    public void on(PaymentProcessedEvent event) {
        commandGateway.send(new ConfirmOrderCommand(event.getOrderId()));
        SagaLifecycle.end();   // saga complete
    }

    // Compensating transaction — rollback
    @SagaEventHandler(associationProperty = "orderId")
    public void on(PaymentFailedEvent event) {
        commandGateway.send(new ReleaseStockCommand(event.getOrderId()));
        commandGateway.send(new CancelOrderCommand(event.getOrderId()));
        SagaLifecycle.end();
    }
}
```

---

### Q3. What is Event Sourcing?

```java
// Store ALL events, not current state. Replay to reconstruct.

// Domain events
public record MoneyDepositedEvent(String accountId, BigDecimal amount, Instant at) {}
public record MoneyWithdrawnEvent(String accountId, BigDecimal amount, Instant at) {}

// Event store — append-only
@Entity @Table(name = "events")
public class DomainEvent {
    @Id private String id;
    private String aggregateId;
    private String eventType;              // "MoneyDeposited"
    private String payload;                // JSON
    private long sequenceNumber;           // ordering
    private Instant occurredAt;
}

// Aggregate rebuilds state by replaying events
public class BankAccount {
    private String id;
    private BigDecimal balance = BigDecimal.ZERO;
    private List<Object> pendingEvents = new ArrayList<>();

    public void deposit(BigDecimal amount) {
        // apply event to self (doesn't persist yet)
        apply(new MoneyDepositedEvent(id, amount, Instant.now()));
    }

    private void apply(MoneyDepositedEvent event) {
        this.balance = this.balance.add(event.amount());
        pendingEvents.add(event);   // collected for persistence
    }

    // Rebuild from history
    public static BankAccount fromHistory(List<Object> events) {
        BankAccount account = new BankAccount();
        for (Object e : events) {
            if (e instanceof MoneyDepositedEvent d) account.apply(d);
            if (e instanceof MoneyWithdrawnEvent w) account.applyWithdraw(w);
        }
        return account;
    }
}

// Snapshot optimization — don't replay 10K events every time
// Every N events, store a snapshot of current state
// On load: find latest snapshot + replay events after snapshot
```

---

### Q4. Kafka Streams — Stateful stream processing

```java
@Configuration
public class OrderProcessingTopology {

    @Bean
    public KStream<String, OrderEvent> buildTopology(StreamsBuilder builder) {
        // Read from topic
        KStream<String, OrderEvent> orders = builder.stream("orders");

        // Filter + transform
        KStream<String, OrderEvent> validOrders = orders
            .filter((key, event) -> event.getTotal().compareTo(BigDecimal.ZERO) > 0)
            .mapValues(event -> enrichOrder(event));

        // Branch — split stream
        Map<String, KStream<String, OrderEvent>> branches = validOrders
            .split(Named.as("prefix-"))
            .branch((k, v) -> v.getTotal().compareTo(new BigDecimal("1000")) > 0,
                    Branched.as("high-value"))
            .branch((k, v) -> true, Branched.as("regular"))
            .noDefaultBranch();

        // Aggregate — count orders per customer (windowed)
        branches.get("prefix-regular")
            .groupByKey()
            .windowedBy(TimeWindows.ofSizeWithNoGrace(Duration.ofHours(1)))
            .count(Materialized.as("order-count-store"))
            .toStream()
            .to("order-counts");

        // KTable — updatable view (latest value per key)
        KTable<String, Product> products = builder.table("products",
            Materialized.as("products-store"));

        // Join stream with table — enrich order with product info
        KStream<String, EnrichedOrder> enriched = validOrders.join(
            products,
            (order, product) -> new EnrichedOrder(order, product),
            Joined.with(Serdes.String(), orderSerde, productSerde));

        enriched.to("enriched-orders");
        return enriched.mapValues(e -> e.getOrder());
    }
}
```

---

### Q5. What is a Service Mesh?

```text
SERVICE MESH — infrastructure layer for service-to-service communication

WITHOUT service mesh:
  Service A → Service B → Service C
  (each service handles retries, mTLS, tracing, circuit-breaking)

WITH service mesh (e.g., Istio, Linkerd):
  Service A → [Sidecar Proxy] → [Sidecar Proxy] → Service B
              Envoy proxy        Envoy proxy

FEATURES handled by mesh (not application code):
├── Traffic management: routing, load balancing, canary releases
├── Security: automatic mTLS between all services, RBAC
├── Observability: distributed tracing, metrics, access logs
├── Resilience: retries, timeouts, circuit breaking
└── Service discovery: no Eureka needed

SPRING BOOT NATIVE vs SERVICE MESH:
Spring: Circuit breaker via Resilience4j in app code
Mesh:   Circuit breaker in Envoy sidecar — transparent, language-agnostic

VirtualService (Istio config for A/B testing):
spec:
  http:
  - match:
    - headers:
        x-user-group: {exact: "beta"}
    route:
    - destination:
        host: order-service
        subset: v2     # 100% beta users to v2
  - route:
    - destination:
        host: order-service
        subset: v1     # everyone else to v1
```

---

### Q6. What is a Dead Letter Queue pattern?

```java
// Multi-stage retry with escalating DLQ — production pattern

// Retry topic chain: orders → orders.retry.1 → orders.retry.2 → orders.DLQ

@Configuration
public class RetryTopicConfig {

    @Bean
    public RetryTopicConfiguration retryTopicConfig(KafkaTemplate<?, ?> template) {
        return RetryTopicConfigurationBuilder
            .newInstance()
            .maxAttempts(4)           // 1 original + 3 retries
            .exponentialBackoff(
                1_000,    // initial delay 1s
                2.0,      // multiplier
                60_000)   // max delay 60s
            .retryOn(TransientException.class, TimeoutException.class)
            .notRetryOn(ValidationException.class)  // don't retry business errors
            .dltSuffix(".DEAD")
            .dltProcessingFailureStrategy(DltStrategy.FAIL_ON_ERROR)
            .create(template);
    }
}

// DLQ monitoring & reprocessing service
@Service
public class DlqReprocessingService {

    @KafkaListener(topics = "orders.DEAD", groupId = "dlq-monitor")
    public void handleDeadLetter(ConsumerRecord<String, OrderEvent> record) {
        String errorMessage = Optional.ofNullable(
            record.headers().lastHeader(KafkaHeaders.EXCEPTION_MESSAGE))
            .map(h -> new String(h.value())).orElse("unknown");

        log.error("DLQ message: key={}, error={}", record.key(), errorMessage);
        alertingService.sendPagerDutyAlert("DLQ accumulation", record);
        auditRepository.saveDlqEntry(record.key(), record.value(), errorMessage);
    }

    // Manual reprocessing endpoint
    public void reprocess(List<String> dlqMessageIds) {
        dlqMessageIds.forEach(id -> {
            DlqEntry entry = auditRepository.findById(id).orElseThrow();
            kafkaTemplate.send("orders", entry.getKey(), entry.getPayload());
            auditRepository.markReprocessed(id);
        });
    }
}
```
