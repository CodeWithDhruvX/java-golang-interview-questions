# 💻 Microservices — Live Coding: Saga & Outbox Patterns

> **Level:** 🔴 Senior
> **Asked at:** Uber, Razorpay, Flipkart, Swiggy, Zepto, PhonePe, CRED, Amazon

> **Interview Reality:** Product companies (especially fintech) will NOT just ask you to explain the Outbox/Saga patterns — they will share a screen and ask you to write it. This file prepares you for that.

---

## CODING Q1. Implement the Transactional Outbox Pattern in Java (Spring Boot + JPA)

**Prompt given by interviewer:**
> *"You have an `OrderService` that needs to place an order in the DB and publish an `OrderPlaced` event to Kafka. How do you guarantee both happen atomically? Write the code."*

### ✅ Step 1: Create the Outbox Table Entity

```java
// OutboxEvent.java
@Entity
@Table(name = "outbox_events")
public class OutboxEvent {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @Column(nullable = false)
    private String aggregateType; // e.g., "ORDER"

    @Column(nullable = false)
    private String aggregateId;   // e.g., "order-123"

    @Column(nullable = false)
    private String eventType;     // e.g., "ORDER_PLACED"

    @Column(nullable = false, columnDefinition = "TEXT")
    private String payload;       // JSON payload

    @Column(nullable = false)
    @Enumerated(EnumType.STRING)
    private OutboxStatus status = OutboxStatus.PENDING;

    @Column(nullable = false)
    private LocalDateTime createdAt = LocalDateTime.now();

    public enum OutboxStatus { PENDING, SENT, FAILED }

    // Getters, setters, constructors...
}
```

### ✅ Step 2: Write to DB + Outbox in ONE Transaction

```java
// OrderService.java
@Service
@Transactional
public class OrderService {

    @Autowired private OrderRepository orderRepository;
    @Autowired private OutboxEventRepository outboxEventRepository;
    @Autowired private ObjectMapper objectMapper;

    public Order placeOrder(CreateOrderRequest request) throws JsonProcessingException {
        // Step 1: Save the order
        Order order = new Order(request.getUserId(), request.getItems(), OrderStatus.PENDING);
        Order savedOrder = orderRepository.save(order);

        // Step 2: Write to Outbox in the SAME transaction
        // Both writes are in the same DB transaction — atomic!
        OutboxEvent outboxEvent = new OutboxEvent();
        outboxEvent.setAggregateType("ORDER");
        outboxEvent.setAggregateId(savedOrder.getId().toString());
        outboxEvent.setEventType("ORDER_PLACED");
        outboxEvent.setPayload(objectMapper.writeValueAsString(
            new OrderPlacedEvent(savedOrder.getId(), savedOrder.getUserId(), savedOrder.getTotal())
        ));
        outboxEventRepository.save(outboxEvent);

        return savedOrder;
        // When this method returns: BOTH order + outbox row committed together.
        // If Kafka is down, the order is still saved. The relay will retry.
    }
}
```

### ✅ Step 3: Outbox Relay — Polling Approach (Simple)

```java
// OutboxRelayService.java
@Service
public class OutboxRelayService {

    @Autowired private OutboxEventRepository outboxEventRepository;
    @Autowired private KafkaTemplate<String, String> kafkaTemplate;

    @Scheduled(fixedDelay = 1000) // Poll every 1 second
    @Transactional
    public void relay() {
        List<OutboxEvent> pendingEvents = outboxEventRepository
            .findTop100ByStatusOrderByCreatedAtAsc(OutboxStatus.PENDING);

        for (OutboxEvent event : pendingEvents) {
            try {
                kafkaTemplate.send(event.getAggregateType().toLowerCase() + "-events",
                                   event.getAggregateId(),
                                   event.getPayload()).get(); // .get() = synchronous, confirms send

                event.setStatus(OutboxStatus.SENT);
            } catch (Exception e) {
                event.setStatus(OutboxStatus.FAILED);
                // Production: add retry_count and exponential backoff
            }
            outboxEventRepository.save(event);
        }
    }
}
```

### ✅ Step 4: Debezium CDC Approach (Production-grade — mention this!)

```
Instead of polling, use Debezium (a CDC tool) to watch the outbox_events table 
MySQL/Postgres binary log and automatically publish to Kafka.
Zero polling overhead. Near real-time (< 50ms latency).

Config: debezium-connector connects to DB → reads WAL → publishes to Kafka topic
Producer: The outbox_events table IS the Kafka producer. No application polling needed.
```

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, PhonePe, Zepto — all fintech/e-commerce companies where order-payment atomicity is critical.

#### Interview Follow-up Questions to Expect:
- *"What happens if your relay crashes midway? Could an event be sent twice?"*
  → Yes, at-least-once delivery. The consumer must be idempotent (check `aggregateId` before processing).
- *"How does Debezium handle this better than polling?"*
  → CDC reads the WAL directly; no polling overhead, near real-time, guaranteed order.

---

## CODING Q2. Implement Saga Orchestration Pattern in Java

**Prompt given by interviewer:**
> *"Design a distributed order checkout flow: OrderService → PaymentService → InventoryService. If Payment fails, the order must be cancelled. If Inventory is out of stock after payment, refund must happen. Write the Saga Orchestrator."*

### ✅ Step 1: Define Saga Steps

```java
// SagaStep.java — represents one step in the saga
public interface SagaStep<T> {
    void execute(T context) throws Exception;
    void compensate(T context);  // The rollback action
}
```

### ✅ Step 2: Saga Context

```java
// OrderSagaContext.java
public class OrderSagaContext {
    private String orderId;
    private String userId;
    private BigDecimal amount;
    private String paymentTransactionId;  // Set after payment step
    private List<String> reservedItemIds; // Set after inventory step
    private String failureReason;

    // Getters, setters, builder...
}
```

### ✅ Step 3: Implement Each Step

```java
// CreateOrderStep.java
@Component
public class CreateOrderStep implements SagaStep<OrderSagaContext> {

    @Autowired private OrderRepository orderRepository;

    @Override
    public void execute(OrderSagaContext ctx) {
        Order order = new Order(ctx.getUserId(), OrderStatus.PENDING);
        order.setId(ctx.getOrderId());
        orderRepository.save(order);
    }

    @Override
    public void compensate(OrderSagaContext ctx) {
        // Mark order as CANCELLED
        orderRepository.updateStatus(ctx.getOrderId(), OrderStatus.CANCELLED);
    }
}

// PaymentStep.java
@Component
public class PaymentStep implements SagaStep<OrderSagaContext> {

    @Autowired private PaymentServiceClient paymentClient; // Feign client

    @Override
    public void execute(OrderSagaContext ctx) throws Exception {
        PaymentResponse response = paymentClient.charge(ctx.getUserId(), ctx.getAmount());
        if (!response.isSuccess()) throw new PaymentFailedException(response.getReason());
        ctx.setPaymentTransactionId(response.getTransactionId());
    }

    @Override
    public void compensate(OrderSagaContext ctx) {
        // REFUND the payment
        if (ctx.getPaymentTransactionId() != null) {
            paymentClient.refund(ctx.getPaymentTransactionId());
        }
    }
}

// InventoryStep.java
@Component
public class InventoryStep implements SagaStep<OrderSagaContext> {

    @Autowired private InventoryServiceClient inventoryClient;

    @Override
    public void execute(OrderSagaContext ctx) throws Exception {
        ReservationResponse res = inventoryClient.reserve(ctx.getOrderId());
        if (!res.isSuccess()) throw new OutOfStockException();
        ctx.setReservedItemIds(res.getReservedItemIds());
    }

    @Override
    public void compensate(OrderSagaContext ctx) {
        if (ctx.getReservedItemIds() != null) {
            inventoryClient.release(ctx.getReservedItemIds());
        }
    }
}
```

### ✅ Step 4: The Saga Orchestrator Engine

```java
// SagaOrchestrator.java
@Service
public class SagaOrchestrator {

    public <T> void execute(List<SagaStep<T>> steps, T context) {
        List<SagaStep<T>> executedSteps = new ArrayList<>();

        for (SagaStep<T> step : steps) {
            try {
                step.execute(context);
                executedSteps.add(step);
            } catch (Exception e) {
                // Step failed — run compensations in REVERSE order
                System.err.println("Step failed: " + step.getClass().getSimpleName() + " — " + e.getMessage());
                rollback(executedSteps, context);
                throw new SagaExecutionException("Saga failed at step: " + step.getClass().getSimpleName(), e);
            }
        }
    }

    private <T> void rollback(List<SagaStep<T>> executedSteps, T context) {
        // Reverse order! Last executed step compensates first
        for (int i = executedSteps.size() - 1; i >= 0; i--) {
            try {
                executedSteps.get(i).compensate(context);
            } catch (Exception e) {
                // Log compensation failure — alert ops team
                // In production: save to a compensation_failures table for manual intervention
                System.err.println("COMPENSATION FAILED for " + executedSteps.get(i).getClass().getSimpleName());
            }
        }
    }
}
```

### ✅ Step 5: Wire Them Together

```java
// OrderCheckoutService.java
@Service
public class OrderCheckoutService {

    @Autowired private SagaOrchestrator orchestrator;
    @Autowired private CreateOrderStep createOrderStep;
    @Autowired private PaymentStep paymentStep;
    @Autowired private InventoryStep inventoryStep;

    public void checkout(CheckoutRequest request) {
        OrderSagaContext context = new OrderSagaContext();
        context.setOrderId(UUID.randomUUID().toString());
        context.setUserId(request.getUserId());
        context.setAmount(request.getAmount());

        List<SagaStep<OrderSagaContext>> steps = List.of(
            createOrderStep,
            paymentStep,
            inventoryStep
        );

        orchestrator.execute(steps, context);
    }
}
```

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber (Go version), Flipkart, Swiggy, Amazon India. The interviewer will watch your compensation logic carefully — *"what if the compensation itself fails?"* is the classic follow-up.

#### ✅ Key Points to Always Say:
- "Compensating transactions are NOT rollbacks — they are new forward-moving actions (e.g., a refund is a new credit transaction, not an undo)"
- "Idempotency is critical: if the InventoryStep compensation is called twice, it should not double-release"
- "In production, I'd use a persistent saga state machine (Axon Framework or custom DB table) so a crashed orchestrator can resume"

---

## CODING Q3. Implement Outbox Pattern in Go

**Prompt:** *"Same problem, but in Go with Postgres."*

```go
// models.go
type OutboxEvent struct {
    ID            string    `db:"id"`
    AggregateType string    `db:"aggregate_type"`
    AggregateID   string    `db:"aggregate_id"`
    EventType     string    `db:"event_type"`
    Payload       string    `db:"payload"`
    Status        string    `db:"status"` // PENDING, SENT, FAILED
    CreatedAt     time.Time `db:"created_at"`
}

// order_service.go
func (s *OrderService) PlaceOrder(ctx context.Context, req CreateOrderRequest) (*Order, error) {
    tx, err := s.db.BeginTx(ctx, nil)
    if err != nil {
        return nil, fmt.Errorf("begin tx: %w", err)
    }
    defer tx.Rollback() // No-op if committed

    // Step 1: Insert order
    order := &Order{
        ID:     uuid.New().String(),
        UserID: req.UserID,
        Status: "PENDING",
    }
    _, err = tx.ExecContext(ctx,
        "INSERT INTO orders (id, user_id, status, created_at) VALUES ($1, $2, $3, $4)",
        order.ID, order.UserID, order.Status, time.Now(),
    )
    if err != nil {
        return nil, fmt.Errorf("insert order: %w", err)
    }

    // Step 2: Insert outbox event in SAME transaction
    payload, _ := json.Marshal(map[string]interface{}{
        "order_id": order.ID,
        "user_id":  order.UserID,
    })
    _, err = tx.ExecContext(ctx, `
        INSERT INTO outbox_events (id, aggregate_type, aggregate_id, event_type, payload, status, created_at)
        VALUES ($1, $2, $3, $4, $5, 'PENDING', $6)`,
        uuid.New().String(), "ORDER", order.ID, "ORDER_PLACED", string(payload), time.Now(),
    )
    if err != nil {
        return nil, fmt.Errorf("insert outbox: %w", err)
    }

    // Both succeed or both fail — atomically
    if err = tx.Commit(); err != nil {
        return nil, fmt.Errorf("commit tx: %w", err)
    }
    return order, nil
}

// outbox_relay.go — background goroutine
func (r *OutboxRelay) Start(ctx context.Context) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            r.relay(ctx)
        }
    }
}

func (r *OutboxRelay) relay(ctx context.Context) {
    rows, err := r.db.QueryContext(ctx,
        "SELECT id, aggregate_type, aggregate_id, event_type, payload FROM outbox_events WHERE status = 'PENDING' ORDER BY created_at LIMIT 100",
    )
    if err != nil {
        log.Printf("outbox query error: %v", err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var e OutboxEvent
        if err := rows.Scan(&e.ID, &e.AggregateType, &e.AggregateID, &e.EventType, &e.Payload); err != nil {
            continue
        }

        topic := strings.ToLower(e.AggregateType) + "-events"
        err = r.producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Key:            []byte(e.AggregateID),
            Value:          []byte(e.Payload),
        }, nil)

        status := "SENT"
        if err != nil {
            status = "FAILED"
            log.Printf("kafka produce error for event %s: %v", e.ID, err)
        }

        r.db.ExecContext(ctx,
            "UPDATE outbox_events SET status = $1 WHERE id = $2",
            status, e.ID,
        )
    }
}
```

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber (Go-first shop), Razorpay Go services. Go interviewers specifically watch for proper `defer tx.Rollback()` + `tx.Commit()` patterns and context propagation.

---

## CODING Q4. Idempotent Kafka Consumer (Java)

**Prompt:** *"Your Outbox relay uses at-least-once delivery. How do you make the consumer idempotent? Show the code."*

```java
// IdempotentOrderEventConsumer.java
@Service
public class OrderEventConsumer {

    @Autowired private ProcessedEventRepository processedEventRepo;
    @Autowired private InventoryService inventoryService;

    @KafkaListener(topics = "order-events", groupId = "inventory-service")
    @Transactional
    public void consume(ConsumerRecord<String, String> record) {
        String eventId = record.key(); // aggregateId used as idempotency key

        // Fast check: have we already processed this event?
        if (processedEventRepo.existsById(eventId)) {
            log.info("Duplicate event {} — skipping", eventId);
            return; // Idempotent: safe to skip
        }

        // Process the event
        OrderPlacedEvent event = objectMapper.readValue(record.value(), OrderPlacedEvent.class);
        inventoryService.reserveStock(event.getOrderId(), event.getItems());

        // Mark as processed IN THE SAME DB TRANSACTION
        processedEventRepo.save(new ProcessedEvent(eventId, LocalDateTime.now()));
    }
}

// ProcessedEvent.java — a simple deduplication table
@Entity
@Table(name = "processed_events")
public class ProcessedEvent {
    @Id
    private String eventId; // The Kafka message key
    private LocalDateTime processedAt;
}
```

#### ✅ Key Insight to Say:
*"The `processed_events` table check + business logic + insert — all in one transaction. Either all succeed (event is fully processed and marked) or all fail (the consumer re-reads and re-tries). This guarantees exactly-once business effect from an at-least-once transport."*
