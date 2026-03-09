# 🏛️ Kafka — Advanced Architecture Patterns

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Netflix, Amazon, Uber, LinkedIn, Meta

---

## Q1. How would you implement Event Sourcing with Kafka?

"Event Sourcing is a pattern where all state changes are stored as a sequence of events. Kafka is perfect for this due to its immutable log nature.

**Core concepts:**
1. **Event Store:** Kafka topic serves as the immutable event log
2. **Aggregates:** Business entities reconstructed from events
3. **Snapshots:** Periodic state checkpoints for performance
4. **Event Versioning:** Schema evolution for compatibility

**Implementation architecture:**
```
Command → Event Store (Kafka) → Event Processors → Read Models
```

**Event design:**
```json
{
  "eventId": "uuid",
  "aggregateId": "order-123",
  "eventType": "OrderCreated",
  "eventData": {...},
  "eventVersion": "1.0",
  "timestamp": "2024-01-01T10:00:00Z"
}
```

**Producer implementation:**
```java
@Component
public class OrderEventProducer {
    @Autowired
    private KafkaTemplate<String, OrderEvent> kafkaTemplate;
    
    public void publishOrderEvent(OrderEvent event) {
        kafkaTemplate.send("order-events", event.getAggregateId(), event);
    }
}
```

**Consumer for state reconstruction:**
```java
@KafkaListener(topics = "order-events", groupId = "order-rebuilder")
public void rebuildOrderState(OrderEvent event) {
    switch(event.getEventType()) {
        case "OrderCreated":
            createNewOrder(event);
            break;
        case "OrderUpdated":
            updateOrder(event);
            break;
        case "OrderCancelled":
            cancelOrder(event);
            break;
    }
}
```

**Snapshot strategy:**
- Create snapshots every 100 events or hourly
- Store snapshots in fast-access database
- Rebuild from snapshot + events since snapshot

**Benefits:**
- Complete audit trail
- Temporal queries (state at any point in time)
- Event replay for debugging
- Natural integration with event-driven architecture

**Challenges:**
- Event schema evolution
- Snapshot management
- Event ordering guarantees
- Storage growth management"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use Spring Data JDBC for snapshot storage. Implement event versioning with Jackson polymorphic deserialization.
* **Golang:** Use Go structs for event definitions. Implement snapshot storage with Redis or PostgreSQL. Use Go's channels for event processing pipelines.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Financial services, e-commerce — where complete audit trails and state reconstruction are regulatory requirements.

#### Indepth
**Compaction Topics:** Use log compaction for event store topics to keep only latest state per aggregate while maintaining full event history in separate archive topics.

---

## Q2. How do you implement CQRS (Command Query Responsibility Segregation) with Kafka?

"CQRS separates read and write models, optimizing each for their specific use case. Kafka provides the perfect event bus between them.

**Architecture overview:**
```
Commands → Write Model → Events (Kafka) → Read Models → Queries
```

**Write Model (Command side):**
```java
@Component
public class OrderCommandHandler {
    @Autowired
    private OrderRepository orderRepository;
    @Autowired
    private KafkaTemplate<String, OrderEvent> kafkaTemplate;
    
    public Order createOrder(CreateOrderCommand command) {
        // Validate command
        Order order = new Order(command);
        orderRepository.save(order);
        
        // Publish event
        OrderEvent event = new OrderCreatedEvent(order);
        kafkaTemplate.send("order-events", order.getId(), event);
        
        return order;
    }
}
```

**Read Model (Query side):**
```java
@KafkaListener(topics = "order-events", groupId = "order-read-model")
public void updateReadModel(OrderEvent event) {
    switch(event.getEventType()) {
        case "OrderCreated":
            orderViewRepository.save(new OrderView(event));
            break;
        case "OrderUpdated":
            updateOrderView(event);
            break;
    }
}

@Service
public class OrderQueryService {
    public OrderView getOrder(String orderId) {
        return orderViewRepository.findById(orderId);
    }
    
    public List<OrderView> getOrdersByCustomer(String customerId) {
        return orderViewRepository.findByCustomerId(customerId);
    }
}
```

**Eventual consistency handling:**
- Use correlation IDs to track command-to-event flow
- Implement read model versioning
- Handle out-of-order events with sequence numbers
- Monitor lag between write and read models

**Scaling considerations:**
- Scale write model vertically (strong consistency)
- Scale read model horizontally (query performance)
- Use multiple read models for different query patterns
- Implement materialized views for complex queries

**Benefits:**
- Optimized read/write performance
- Independent scaling of models
- Better separation of concerns
- Natural fit for microservices

**Trade-offs:**
- Eventual consistency complexity
- Increased system complexity
- Debugging across models
- Data synchronization challenges"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use Spring Data for different databases per model (PostgreSQL for write, MongoDB for read). Implement event correlation with Spring Cloud Sleuth.
* **Golang:** Use different database drivers per model (SQL for write, NoSQL for read). Implement correlation tracking with Go contexts.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** E-commerce platforms, social media — where read-heavy workloads require different optimization than write operations.

#### Indepth
**Saga Pattern:** Combine CQRS with Saga pattern for distributed transactions. Use Kafka events to coordinate compensating actions across multiple services.

---

## Q3. How would you design a multi-region active-active Kafka architecture?

"Active-active multi-region architecture provides low latency and high availability across geographic regions.

**Architecture components:**
1. **Regional Kafka clusters:** Independent clusters per region
2. **Cross-region replication:** MirrorMaker 2 or Cluster Linking
3. **Traffic routing:** GeoDNS or API Gateway
4. **Conflict resolution:** Last-write-wins or application-specific

**Replication topology:**
```
US-East Cluster ←→ Europe Cluster ←→ Asia-Pacific Cluster
```

**MirrorMaker 2 configuration:**
```properties
# mm2.properties
clusters=us-east, europe, asia-pacific
us-east.bootstrap.servers=us-east-kafka:9092
europe.bootstrap.servers=europe-kafka:9092
asia-pacific.bootstrap.servers=apac-kafka:9092

# Replication policies
replication.policy.class=org.apache.kafka.connect.mirror.DefaultReplicationPolicy
sync.topic.acls.enabled=true
sync.topic.configs.enabled=true

# Heartbeat and checkpoint
heartbeats.topic.replication.factor=3
offset-syncs.topic.replication.factor=3
```

**Conflict resolution strategies:**
1. **Timestamp-based:** Use event timestamps
2. **Region priority:** Predefined region hierarchy
3. **Application-level:** Business logic for conflicts
4. **Vector clocks:** Causal ordering detection

**Producer configuration for active-active:**
```java
@Configuration
public class MultiRegionProducerConfig {
    
    @Bean
    public ProducerFactory<String, Object> producerFactory() {
        Map<String, Object> props = new HashMap<>();
        props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, 
                 getRegionalBootstrapServers());
        props.put(ProducerConfig.ENABLE_IDEMPOTENCE_CONFIG, true);
        props.put(ProducerConfig.ACKS_CONFIG, "all");
        return new DefaultKafkaProducerFactory<>(props);
    }
}
```

**Consumer considerations:**
- Use regional consumer groups for local processing
- Implement duplicate detection across regions
- Handle out-of-order events from different regions
- Monitor cross-region replication lag

**Failure scenarios:**
1. **Region isolation:**
   - Local cluster continues operating
   - Queued events replicate when connectivity restored
   - Manual conflict resolution may be needed

2. **Network partition:**
   - Each region operates independently
   - Event ordering may be affected
   - Reconciliation process after recovery

**Monitoring:**
- Cross-region replication lag metrics
- Regional throughput and latency
- Conflict resolution statistics
- Cluster health per region

**Cost optimization:**
- Tiered storage for cold data
- Compression for cross-region traffic
- Selective replication (not all topics)
- Regional instance sizing based on load"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use Spring Cloud Gateway for regional routing. Implement circuit breakers for cross-region calls.
* **Golang:** Use Go's built-in HTTP client with regional endpoints. Implement connection pooling for cross-region Kafka connections.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Global platforms (Netflix, Amazon) — where users are distributed worldwide and latency is critical.

#### Indepth
**Cluster Linking:** Confluent's Cluster Linking provides more sophisticated active-active capabilities than MirrorMaker 2, including bidirectional replication and conflict resolution.

---

## Q4. How do you implement a real-time analytics pipeline using Kafka and stream processing?

"Real-time analytics pipelines combine Kafka's event streaming with stream processing for immediate insights.

**Pipeline architecture:**
```
Events → Kafka Topics → Stream Processing → Analytics Store → Dashboard
```

**Stream processing options:**
1. **Kafka Streams:** Native Kafka stream processing
2. **Apache Flink:** Advanced stream processing
3. **Apache Spark Structured Streaming:** Batch and stream unified
4. **ksqlDB:** SQL interface for stream processing

**Kafka Streams example:**
```java
@Component
public class OrderAnalyticsProcessor {
    
    @Autowired
    private StreamsBuilder streamsBuilder;
    
    public KStream<String, OrderEvent> buildAnalyticsPipeline() {
        KStream<String, OrderEvent> orderStream = streamsBuilder
            .stream("order-events", Consumed.with(Serdes.String(), 
                 new JsonSerde<>(OrderEvent.class)));
        
        // Calculate revenue per minute
        KTable<Windowed<String>, Double> revenueByMinute = orderStream
            .groupByKey()
            .windowedBy(TimeWindows.of(Duration.ofMinutes(1)))
            .aggregate(
                () -> 0.0,
                (key, order, revenue) -> revenue + order.getAmount(),
                Materialized.with(Serdes.String(), Serdes.Double())
            );
        
        revenueByMinute.toStream()
            .to("order-revenue-by-minute", 
                Produced.with(WindowedSerdes.timeWindowedSerdeFrom(
                    Serdes.String()), Serdes.Double()));
        
        return orderStream;
    }
}
```

**Flink integration:**
```java
public class FlinkOrderAnalytics {
    public void execute() throws Exception {
        StreamExecutionEnvironment env = 
            StreamExecutionEnvironment.getExecutionEnvironment();
        
        KafkaSource<OrderEvent> source = KafkaSource.<OrderEvent>builder()
            .setBootstrapServers("kafka:9092")
            .setTopics("order-events")
            .setDeserializer(new OrderEventDeserializer())
            .build();
        
        DataStream<OrderEvent> orderStream = env.fromSource(
            source, WatermarkStrategy.noWatermarks(), "order-source");
        
        // Real-time aggregations
        DataStream<Tuple2<String, Double>> revenueByCategory = orderStream
            .keyBy(OrderEvent::getCategory)
            .window(TumblingProcessingTimeWindows.of(Time.minutes(1)))
            .aggregate(new RevenueAggregator());
        
        revenueByCategory.print();
        env.execute("Order Analytics");
    }
}
```

**Analytics storage:**
1. **Redis:** Real-time counters and dashboards
2. **ClickHouse:** Columnar database for analytics
3. **Elasticsearch:** Full-text search and aggregations
4. **TimescaleDB:** Time-series data optimized

**Dashboard integration:**
```javascript
// WebSocket connection for real-time updates
const socket = new WebSocket('ws://analytics-service:8080/updates');

socket.onmessage = function(event) {
    const data = JSON.parse(event.data);
    updateDashboard(data);
};

function updateDashboard(data) {
    document.getElementById('revenue-today').textContent = 
        '$' + data.revenueToday.toLocaleString();
    document.getElementById('orders-count').textContent = 
        data.ordersToday.toLocaleString();
}
```

**Performance optimization:**
- Window sizing for aggregation granularity
- Backpressure handling for bursty traffic
- State management for stream processors
- Exactly-once processing guarantees

**Monitoring:**
- Stream processing latency
- End-to-end pipeline lag
- Throughput metrics per stage
- Error rates and retry patterns"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use Spring Cloud Stream for stream processing abstraction. Integrate with Micrometer for metrics collection.
* **Golang:** Use Go's concurrency for stream processing. Implement custom aggregators with Go's channels and goroutines.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Netflix — where real-time analytics drive business decisions and user experiences.

#### Indepth
**Window Types:** Tumbling windows (non-overlapping), sliding windows (overlapping), and session windows (activity-based) each serve different analytics use cases.

---

## Q5. How do you design a microservices communication pattern using Kafka?

"Kafka enables sophisticated microservices communication patterns beyond simple request-response.

**Communication patterns:**
1. **Event-driven architecture:** Services communicate through events
2. **Saga pattern:** Distributed transactions via events
3. **API composition:** Materialized views for complex queries
4. **Command/Query separation:** Different topics for different purposes

**Event-driven communication:**
```java
// Service A - Order Service
@EventListener
public void handleOrderCreated(OrderCreatedEvent event) {
    // Process order
    Order order = orderRepository.save(event.toOrder());
    
    // Publish events for other services
    eventPublisher.publishEvent(new PaymentRequiredEvent(order));
    eventPublisher.publishEvent(new InventoryReservedEvent(order));
}

// Service B - Payment Service
@KafkaListener(topics = "payment-events", groupId = "payment-service")
public void handlePaymentRequired(PaymentRequiredEvent event) {
    Payment payment = paymentProcessor.process(event.getOrderId());
    
    if (payment.isSuccessful()) {
        kafkaTemplate.send("payment-completed", 
            new PaymentCompletedEvent(event.getOrderId(), payment.getId()));
    } else {
        kafkaTemplate.send("payment-failed", 
            new PaymentFailedEvent(event.getOrderId(), payment.getReason()));
    }
}
```

**Saga orchestration:**
```java
@Component
public class OrderSagaOrchestrator {
    
    @KafkaListener(topics = "order-events", groupId = "order-saga")
    public void startSaga(OrderCreatedEvent event) {
        SagaManager saga = SagaManager.start(event.getOrderId());
        
        // Step 1: Reserve inventory
        saga.addStep("reserve-inventory", 
            () -> kafkaTemplate.send("inventory-commands", 
                new ReserveInventoryCommand(event.getOrderId())),
            () -> kafkaTemplate.send("inventory-compensation", 
                new ReleaseInventoryCommand(event.getOrderId())));
        
        // Step 2: Process payment
        saga.addStep("process-payment",
            () -> kafkaTemplate.send("payment-commands", 
                new ProcessPaymentCommand(event.getOrderId())),
            () -> kafkaTemplate.send("payment-compensation", 
                new RefundPaymentCommand(event.getOrderId())));
    }
    
    @KafkaListener(topics = "payment-completed", groupId = "order-saga")
    public void handlePaymentCompleted(PaymentCompletedEvent event) {
        SagaManager.completeStep(event.getOrderId(), "process-payment");
    }
}
```

**API composition with materialized views:**
```java
// Read model service
@KafkaListener(topics = {"order-events", "payment-events", "inventory-events"}, 
               groupId = "order-composite-view")
public void updateCompositeView(Object event) {
    if (event instanceof OrderCreatedEvent) {
        updateOrderView((OrderCreatedEvent) event);
    } else if (event instanceof PaymentCompletedEvent) {
        updatePaymentStatus((PaymentCompletedEvent) event);
    }
    // Update composite view in read-optimized database
}

@RestController
public class OrderCompositeController {
    @GetMapping("/orders/{id}/composite")
    public OrderCompositeView getOrderComposite(@PathVariable String id) {
        return compositeViewRepository.findById(id);
    }
}
```

**Topic design patterns:**
1. **Command topics:** `service-name-commands`
2. **Event topics:** `service-name-events`
3. **Compensation topics:** `service-name-compensation`
4. **Query topics:** `service-name-queries`

**Service discovery:**
```yaml
# Spring Cloud Stream configuration
spring:
  cloud:
    stream:
      kafka:
        binder:
          brokers: kafka-1:9092,kafka-2:9092,kafka-3:9092
      bindings:
        orderEvents:
          destination: order-events
          group: order-service
        paymentEvents:
          destination: payment-events
          group: payment-service
```

**Benefits:**
- Loose coupling between services
- Asynchronous communication
- Natural scalability
- Event replay capabilities
- Fault tolerance through retries

**Challenges:**
- Event ordering across services
- Distributed transaction complexity
- Debugging across service boundaries
- Schema evolution management"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use Spring Cloud Stream for Kafka integration. Implement circuit breakers with Resilience4j for service resilience.
* **Golang:** Use Go's microservices frameworks like go-micro. Implement service discovery with Consul or etcd.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Large enterprises, cloud platforms — where microservices architecture requires sophisticated communication patterns.

#### Indepth
**Schema Registry:** Use Confluent Schema Registry for managing event schemas across services. Implement compatibility strategies (backward, forward, full) for independent service evolution.
