# 🚀 Advanced Streaming & Real-time Processing

> **Focus:** Product-Based Companies (Uber, Netflix, Twitter, Robinhood, Stripe)
> **Level:** 🔴 Senior – 🟠 Staff

---

## 📋 Table of Contents

1. [Streaming Architecture Fundamentals](#1-streaming-architecture-fundamentals)
2. [Stream Processing Engines](#2-stream-processing-engines)
3. [Windowing & Time Semantics](#3-windowing--time-semantics)
4. [State Management in Streaming](#4-state-management-in-streaming)
5. [Stream-Table Joins](#5-stream-table-joins)
6. [Exactly-Once Processing](#6-exactly-once-processing)
7. [Backpressure & Flow Control](#7-backpressure--flow-control)
8. [Real-time Analytics](#8-real-time-analytics)
9. [Stream Processing Patterns](#9-stream-processing-patterns)
10. [Common Interview Questions](#10-common-interview-questions)

---

## 1. Streaming Architecture Fundamentals

### Q1: What are the key components of a streaming architecture?

**Answer:**

**Core Components:**
```
┌─────────────┐    ┌──────────────┐    ┌─────────────┐    ┌─────────────┐
│   Data      │───▶│   Message    │───▶│  Stream     │───▶│   Sink      │
│  Sources    │    │    Broker    │    │ Processor  │    │  Systems    │
│             │    │             │    │             │    │             │
│ • Kafka     │    │ • Kafka      │    │ • Flink     │    │ • Databases │
│ • Kinesis   │    │ • Pulsar     │    │ • Spark     │    │ • Storage   │
│ • Pulsar    │    │ • RabbitMQ   │    │ • Kafka     │    │ • Analytics │
│ • APIs      │    │ • NATS       │    │ • Kinesis   │    │ • Alerts    │
└─────────────┘    └──────────────┘    └─────────────┘    └─────────────┘
```

**Detailed Architecture:**
- **Data Sources:** Event generators, sensors, applications
- **Message Broker:** Reliable message transport and retention
- **Stream Processor:** Real-time computation and transformation
- **State Store:** Persistent state for windowing and joins
- **Sink Systems:** Storage, analytics, and downstream systems

### Q2: Explain event time vs processing time vs ingestion time

**Answer:**

**Time Semantics:**
```java
public class TimeSemantics {
    
    // Event Time: When the event actually occurred
    @Timestamp("event_timestamp")
    public class Event {
        private long eventTimestamp;  // When user clicked
        private String userId;
        private String action;
    }
    
    // Processing Time: When the event is processed
    public void processEvent(Event event) {
        long processingTime = System.currentTimeMillis();
        
        // Calculate lag
        long eventTimeLag = processingTime - event.getEventTimestamp();
        
        // Handle out-of-order events
        if (eventTimeLag > MAX_ALLOWED_LAG) {
            handleLateEvent(event);
        }
    }
    
    // Ingestion Time: When event entered the system
    @StreamTime
    public void handleIngestion(Event event) {
        long ingestionTime = System.currentTimeMillis();
        // Used when event time is unavailable
    }
}
```

**Watermarks for Late Data:**
```java
// Flink Watermark Strategy
WatermarkStrategy<Event> watermarkStrategy = WatermarkStrategy
    .<Event>forBoundedOutOfOrderness(Duration.ofSeconds(10))
    .withTimestampAssigner((event, timestamp) -> event.getEventTimestamp())
    .withIdleness(Duration.ofMinutes(5));

DataStream<Event> stream = env.addSource(kafkaSource)
    .assignTimestampsAndWatermarks(watermarkStrategy);
```

---

## 2. Stream Processing Engines

### Q3: Compare Flink vs Spark Structured Streaming vs Kafka Streams

**Answer:**

**Apache Flink:**
```java
// Flink Stream Processing
StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();

DataStream<Event> events = env.addSource(new FlinkKafkaConsumer<>(
    "events",
    new EventDeserializer(),
    kafkaProperties));

// Windowed aggregation
DataStream<Result> results = events
    .keyBy(Event::getUserId)
    .window(TumblingEventTimeWindows.of(Time.minutes(5)))
    .aggregate(new CountAggregate());

// Stateful processing
DataStream<Alert> alerts = events
    .keyBy(Event::getUserId)
    .process(new FraudDetectionProcessFunction());

env.execute("Real-time Processing");
```

**Characteristics:**
- **True stream processing:** Event-at-a-time processing
- **Advanced windowing:** Event time, session windows, custom triggers
- **State management:** Checkpoints, savepoints, exactly-once
- **Low latency:** Sub-millisecond processing possible

**Spark Structured Streaming:**
```scala
// Spark Structured Streaming
val spark = SparkSession.builder.appName("Streaming").getOrCreate()

val events = spark.readStream
  .format("kafka")
  .option("kafka.bootstrap.servers", "localhost:9092")
  .option("subscribe", "events")
  .load()

val results = events
  .selectExpr("CAST(value AS STRING) as json")
  .as[String]
  .map(parseEvent)
  .withWatermark("timestamp", "10 minutes")
  .groupBy(window(col("timestamp"), "5 minutes"), col("userId"))
  .count()

val query = results.writeStream
  .format("console")
  .outputMode("update")
  .start()
```

**Characteristics:**
- **Micro-batch processing:** Small batches for streaming
- **Unified API:** Same code for batch and streaming
- **SQL integration:** Complex queries with window functions
- **Exactly-once:** Through checkpointing

**Kafka Streams:**
```java
// Kafka Streams
Properties props = new Properties();
props.put(StreamsConfig.APPLICATION_ID_CONFIG, "fraud-detection");
props.put(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092");

KStream<String, Event> events = builder.stream("events");

KTable<String, Long> userCounts = events
    .groupBy((key, event) -> event.getUserId())
    .count(Materialized.as("user-counts-store"));

// Join with reference data
KTable<String, User> users = builder.table("users");
KStream<String, EnrichedEvent> enriched = events
    .leftJoin(users, (event, user) -> enrichEvent(event, user));

KafkaStreams streams = new KafkaStreams(builder.build(), props);
streams.start();
```

**Characteristics:**
- **Kafka-native:** Built on top of Kafka
- **State stores:** RocksDB for local state
- **Exactly-once:** Through Kafka transactions
- **Deployment:** Library, not a cluster

### Q4: How do you handle state management in streaming applications?

**Answer:**

**Keyed State in Flink:**
```java
public class FraudDetectionFunction extends KeyedProcessFunction<String, Event, Alert> {
    
    // Value state for user transaction history
    private ValueState<TransactionHistory> historyState;
    
    // List state for recent transactions
    private ListState<Transaction> recentTransactionsState;
    
    // Map state for merchant patterns
    private MapState<String, MerchantPattern> merchantPatternsState;
    
    @Override
    public void open(Configuration parameters) {
        // Initialize state descriptors
        ValueStateDescriptor<TransactionHistory> historyDesc = 
            new ValueStateDescriptor<>("history", TransactionHistory.class);
        historyState = getRuntimeContext().getState(historyDesc);
        
        ListStateDescriptor<Transaction> recentDesc = 
            new ListStateDescriptor<>("recent", Transaction.class);
        recentTransactionsState = getRuntimeContext().getListState(recentDesc);
        
        MapStateDescriptor<String, MerchantPattern> merchantDesc = 
            new MapStateDescriptor<>("merchants", String.class, MerchantPattern.class);
        merchantPatternsState = getRuntimeContext().getMapState(merchantDesc);
    }
    
    @Override
    public void processElement(Event event, Context context, Collector<Alert> out) {
        try {
            // Update transaction history
            TransactionHistory history = historyState.value();
            if (history == null) {
                history = new TransactionHistory();
            }
            history.addTransaction(event);
            historyState.update(history);
            
            // Update recent transactions (sliding window)
            recentTransactionsState.add(event.getTransaction());
            cleanupOldTransactions(context.timestamp());
            
            // Check for fraud patterns
            if (detectFraud(event, history)) {
                out.collect(new Alert(event.getUserId(), "Suspicious activity detected"));
            }
            
        } catch (Exception e) {
            // Handle state access errors
            throw new RuntimeException("State management error", e);
        }
    }
    
    private void cleanupOldTransactions(long currentTimestamp) {
        // Remove transactions older than 24 hours
        recentTransactionsState.forEach(transaction -> {
            if (currentTimestamp - transaction.getTimestamp() > TimeUnit.HOURS.toMillis(24)) {
                recentTransactionsState.remove(transaction);
            }
        });
    }
}
```

**Checkpointing Configuration:**
```java
// Enable exactly-once processing with checkpoints
env.enableCheckpointing(60000); // Checkpoint every minute
env.getCheckpointConfig().setCheckpointingMode(CheckpointingMode.EXACTLY_ONCE);
env.getCheckpointConfig().setMinPauseBetweenCheckpoints(30000);
env.getCheckpointConfig().setCheckpointTimeout(600000);
env.getCheckpointConfig().setMaxConcurrentCheckpoints(1);
env.getCheckpointConfig().enableExternalizedCheckpoints(
    CheckpointConfig.ExternalizedCheckpointCleanup.RETAIN_ON_CANCELLATION);

// State backend configuration
env.setStateBackend(new RocksDBStateBackend("hdfs://namenode:port/flink/checkpoints"));
```

---

## 3. Windowing & Time Semantics

### Q5: What are different windowing strategies and when to use them?

**Answer:**

**Tumbling Windows:**
```java
// Fixed-size, non-overlapping windows
DataStream<Event> events = env.addSource(kafkaSource);

DataStream<WindowedResult> results = events
    .keyBy(Event::getUserId)
    .window(TumblingEventTimeWindows.of(Time.minutes(5)))
    .aggregate(new CountAggregate());

// Use cases: Regular reporting, periodic aggregation
// Example: User activity per 5-minute interval
```

**Sliding Windows:**
```java
// Overlapping windows for smooth metrics
DataStream<Metric> metrics = events
    .keyBy(Event::getUserId)
    .window(SlidingEventTimeWindows.of(Time.minutes(10), Time.minutes(1)))
    .aggregate(new MovingAverageAggregate());

// Use cases: Real-time dashboards, trend analysis
// Example: 10-minute moving average updated every minute
```

**Session Windows:**
```java
// Dynamic windows based on activity gaps
DataStream<SessionResult> sessions = events
    .keyBy(Event::getUserId)
    .window(EventTimeSessionWindows.withGap(Time.minutes(30)))
    .process(new SessionProcessFunction());

// Use cases: User behavior analysis, session tracking
// Example: User session ends after 30 minutes of inactivity
```

**Custom Windows:**
```java
// Custom trigger for business logic
public class CustomWindowTrigger extends Trigger<Event, TimeWindow> {
    
    @Override
    public TriggerResult onElement(Event element, 
                                 long timestamp, 
                                 TimeWindow window, 
                                 TriggerContext ctx) {
        
        // Trigger on specific conditions
        if (element.isHighValue()) {
            return TriggerResult.FIRE_AND_PURGE;
        }
        
        // Regular time-based trigger
        return TriggerResult.CONTINUE;
    }
    
    @Override
    public TriggerResult onProcessingTime(long time, 
                                        TimeWindow window, 
                                        TriggerContext ctx) {
        return TriggerResult.FIRE;
    }
    
    @Override
    public TriggerResult onEventTime(long time, 
                                   TimeWindow window, 
                                   TriggerContext ctx) {
        return TriggerResult.FIRE;
    }
}

// Apply custom window
DataStream<CustomResult> customResults = events
    .keyBy(Event::getUserId)
    .window(GlobalWindows.create())
    .trigger(new CustomWindowTrigger())
    .process(new CustomProcessFunction());
```

### Q6: How do you handle late data in stream processing?

**Answer:**

**Watermark Strategy for Late Data:**
```java
// Allow up to 10 seconds of lateness
WatermarkStrategy<Event> watermarkStrategy = WatermarkStrategy
    .<Event>forBoundedOutOfOrderness(Duration.ofSeconds(10))
    .withTimestampAssigner((event, timestamp) -> event.getEventTimestamp())
    .withIdleness(Duration.ofMinutes(5));

// Side output for late events
OutputTag<Event> lateEventsTag = new OutputTag<Event>("late-events") {};

SingleOutputStreamOperator<MainResult> mainStream = events
    .keyBy(Event::getUserId)
    .window(TumblingEventTimeWindows.of(Time.minutes(1)))
    .sideOutputLateData(lateEventsTag)
    .aggregate(new MainAggregate());

// Process late events separately
DataStream<Event> lateEvents = mainStream.getSideOutput(lateEventsTag);
lateEvents.process(new LateEventProcessor());
```

**Allowed Lateness with Updates:**
```java
// Allow late data to update previous results
DataStream<WindowedResult> results = events
    .keyBy(Event::getUserId)
    .window(TumblingEventTimeWindows.of(Time.minutes(5)))
    .allowedLateness(Time.minutes(2))  // Allow 2 minutes of lateness
    .sideOutputLateData(lateEventsTag)
    .aggregate(new UpdatingAggregate());

// Aggregate that handles updates
public class UpdatingAggregate implements AggregateFunction<Event, 
                                                         Accumulator, 
                                                         WindowedResult> {
    
    @Override
    public Accumulator add(Event event, Accumulator accumulator) {
        // Handle both on-time and late events
        accumulator.addEvent(event);
        return accumulator;
    }
    
    @Override
    public WindowedResult getResult(Accumulator accumulator) {
        return accumulator.getResult();
    }
}
```

---

## 4. State Management in Streaming

### Q7: What are different types of state in stream processing?

**Answer:**

**Keyed State:**
```java
public class KeyedStateExample extends KeyedProcessFunction<String, Event, Result> {
    
    // Value State: Single value per key
    private ValueState<UserProfile> userProfileState;
    
    // List State: Collection of values per key
    private ListState<Event> recentEventsState;
    
    // Map State: Key-value pairs per key
    private MapState<String, Integer> categoryCountsState;
    
    // Reducing State: Aggregated value per key
    private ReducingState<Double> totalAmountState;
    
    @Override
    public void open(Configuration parameters) {
        // Initialize state descriptors
        ValueStateDescriptor<UserProfile> profileDesc = 
            new ValueStateDescriptor<>("userProfile", UserProfile.class);
        userProfileState = getRuntimeContext().getState(profileDesc);
        
        ListStateDescriptor<Event> eventsDesc = 
            new ListStateDescriptor<>("recentEvents", Event.class);
        recentEventsState = getRuntimeContext().getListState(eventsDesc);
        
        MapStateDescriptor<String, Integer> countsDesc = 
            new MapStateDescriptor<>("categoryCounts", String.class, Integer.class);
        categoryCountsState = getRuntimeContext().getMapState(countsDesc);
        
        ReducingStateDescriptor<Double> amountDesc = 
            new ReducingStateDescriptor<>("totalAmount", 
                new SumReduce(), Double.class);
        totalAmountState = getRuntimeContext().getReducingState(amountDesc);
    }
    
    @Override
    public void processElement(Event event, Context context, Collector<Result> out) {
        try {
            // Update user profile
            UserProfile profile = userProfileState.value();
            if (profile == null) {
                profile = new UserProfile(event.getUserId());
            }
            profile.updateFromEvent(event);
            userProfileState.update(profile);
            
            // Add to recent events
            recentEventsState.add(event);
            
            // Update category counts
            String category = event.getCategory();
            Integer count = categoryCountsState.get(category);
            categoryCountsState.put(category, count == null ? 1 : count + 1);
            
            // Update total amount
            totalAmountState.add(event.getAmount());
            
            // Generate result
            Result result = new Result(
                profile,
                getTotalRecentEvents(),
                categoryCountsState.entries(),
                totalAmountState.get()
            );
            
            out.collect(result);
            
        } catch (Exception e) {
            throw new RuntimeException("State access error", e);
        }
    }
}
```

**Operator State:**
```java
public class OperatorStateExample extends RichSinkFunction<Event> {
    
    // List State: Parallel instances each have part of the state
    private transient ListState<String> processedEventsState;
    
    // Union State: All instances have the same state
    private transient ListState<Configuration> configState;
    
    @Override
    public void open(Configuration parameters) throws Exception {
        super.open(parameters);
        
        // Initialize list state
        ListStateDescriptor<String> processedDesc = 
            new ListStateDescriptor<>("processedEvents", String.class);
        processedEventsState = getRuntimeContext().getOperatorState(
            processedDesc, true, ListStateDescriptor.OperatorStateMode.SPLIT_DISTRIBUTE);
        
        // Initialize union state
        ListStateDescriptor<Configuration> configDesc = 
            new ListStateDescriptor<>("config", Configuration.class);
        configState = getRuntimeContext().getOperatorState(
            configDesc, true, ListStateDescriptor.OperatorStateMode.UNION);
    }
    
    @Override
    public void invoke(Event event, Context context) throws Exception {
        String eventId = event.getId();
        
        // Check if already processed
        for (String processedId : processedEventsState.get()) {
            if (processedId.equals(eventId)) {
                return; // Skip duplicate
            }
        }
        
        // Process event
        processEvent(event);
        
        // Mark as processed
        processedEventsState.add(eventId);
    }
}
```

**State TTL (Time To Live):**
```java
// Configure state TTL
StateTtlConfig ttlConfig = StateTtlConfig
    .newBuilder(Time.hours(24))  // State expires after 24 hours
    .setUpdateType(StateTtlConfig.UpdateType.OnCreateAndWrite)  // Update on access
    .setStateVisibility(StateTtlConfig.StateVisibility.NeverReturnExpired)  // Never return expired
    .cleanupInRocksdbCompactFilter(1000)  // Cleanup during compaction
    .build();

ValueStateDescriptor<UserProfile> profileDesc = 
    new ValueStateDescriptor<>("userProfile", UserProfile.class);
profileDesc.enableTimeToLive(ttlConfig);

userProfileState = getRuntimeContext().getState(profileDesc);
```

---

## 5. Stream-Table Joins

### Q8: How do you implement stream-table joins in real-time processing?

**Answer:**

**Stream-Table Join with Flink:**
```java
// Stream of events
DataStream<Event> eventStream = env.addSource(kafkaSource);

// Table of reference data (users)
Table usersTable = tableEnvironment.sqlQuery(
    "SELECT userId, name, tier, segment FROM users");

// Join stream with table
Table resultTable = tableEnvironment
    .fromDataStream(eventStream, "userId, action, amount, eventTime")
    .join(usersTable)
    .where("eventStream.userId = users.userId")
    .select("eventStream.*, users.name, users.tier, users.segment");

// Convert back to stream
DataStream<EnrichedEvent> enrichedStream = 
    tableEnvironment.toDataStream(resultTable, EnrichedEvent.class);
```

**Temporal Table Join:**
```java
// For versioned reference data
Table userVersionsTable = tableEnvironment.fromValues(
    row(1L, "Alice", "GOLD", toTimestamp("2023-01-01 00:00:00")),
    row(1L, "Alice", "PLATINUM", toTimestamp("2023-06-01 00:00:00")),
    row(2L, "Bob", "SILVER", toTimestamp("2023-01-01 00:00:00"))
).as("userId, name, tier, versionTime");

// Define temporal table function
TemporalTableFunction userVersions = userVersionsTable
    .createTemporalTableFunction(
        $"versionTime",  // Time attribute
        $"userId"        // Primary key
    );

// Register temporal table function
tableEnvironment.createTemporaryFunction("UserVersions", userVersions);

// Join with temporal table
Table enrichedTable = tableEnvironment
    .fromDataStream(eventStream, "userId, action, amount, eventTime")
    .joinLateral(call("UserVersions", $"eventTime"))
    .select("userId, action, amount, tier");
```

**Kafka Streams Stream-Table Join:**
```java
// Stream of events
KStream<String, Event> eventStream = builder.stream("events");

// Table of user data
KTable<String, User> userTable = builder.table("users");

// Stream-table join
KStream<String, EnrichedEvent> enrichedStream = eventStream
    .leftJoin(userTable, (event, user) -> {
        if (user != null) {
            return new EnrichedEvent(event, user);
        } else {
            return new EnrichedEvent(event, new User(event.getUserId()));
        }
    });

// Output enriched events
enrichedStream.to("enriched-events");
```

**Dynamic Table Updates:**
```java
// Handle table updates in real-time
public class UserTableUpdateProcessor extends 
    RichCoProcessFunction<Event, User, EnrichedEvent> {
    
    private ValueState<User> userState;
    
    @Override
    public void open(Configuration parameters) {
        ValueStateDescriptor<User> userDesc = 
            new ValueStateDescriptor<>("user", User.class);
        userState = getRuntimeContext().getState(userDesc);
    }
    
    @Override
    public void processElement1(Event event, Context context, 
                               Collector<EnrichedEvent> out) throws Exception {
        User user = userState.value();
        if (user != null) {
            out.collect(new EnrichedEvent(event, user));
        }
    }
    
    @Override
    public void processElement2(User user, Context context, 
                               Collector<EnrichedEvent> out) throws Exception {
        // Update user state
        userState.update(user);
        
        // Register timer for cleanup if needed
        context.timerService().registerProcessingTimeTimer(
            System.currentTimeMillis() + TimeUnit.HOURS.toMillis(24));
    }
}
```

---

## 6. Exactly-Once Processing

### Q9: How do you achieve exactly-once processing in streaming applications?

**Answer:**

**Flink Exactly-Once with Checkpoints:**
```java
// Enable exactly-once processing
env.enableCheckpointing(60000); // Checkpoint every minute
env.getCheckpointConfig().setCheckpointingMode(CheckpointingMode.EXACTLY_ONCE);

// Two-Phase Commit Sink
public class ExactlyOnceKafkaSink extends TwoPhaseCommitSinkFunction<Event, 
                                                               Transaction, 
                                                               Void> {
    
    public ExactlyOnceKafkaSink() {
        super(new KafkaSerializer(), new KafkaTransactionCoordinator(), 
              new KafkaTransactionContext());
    }
    
    @Override
    protected void invoke(Transaction transaction, Event event, Context context) {
        // Write to transaction (not yet committed)
        transaction.send(event.getTopic(), event.getKey(), event.getValue());
    }
    
    @Override
    protected Transaction beginTransaction() throws Exception {
        // Begin Kafka transaction
        return kafkaProducer.beginTransaction();
    }
    
    @Override
    protected void preCommit(Transaction transaction) throws Exception {
        // Prepare transaction for commit
        transaction.flush();
    }
    
    @Override
    protected void commit(Transaction transaction) {
        // Commit transaction (makes data visible)
        transaction.commit();
    }
    
    @Override
    protected void abort(Transaction transaction) {
        // Abort transaction on failure
        transaction.abort();
    }
}
```

**Kafka Streams Exactly-Once:**
```java
// Configure exactly-once processing
Properties props = new Properties();
props.put(StreamsConfig.APPLICATION_ID_CONFIG, "exactly-once-app");
props.put(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092");
props.put(StreamsConfig.PROCESSING_GUARANTEE_CONFIG, 
          StreamsConfig.EXACTLY_ONCE_V2);
props.put(StreamsConfig.REPLICATION_FACTOR_CONFIG, 3);

// Enable idempotent producer
props.put(ProducerConfig.ENABLE_IDEMPOTENCE_CONFIG, "true");

// Build topology
KStream<String, Event> events = builder.stream("input-topic");
KStream<String, Result> results = events
    .mapValues(value -> process(value))
    .filter((key, value) -> value.isValid());

// Write with exactly-once guarantee
results.to("output-topic");

KafkaStreams streams = new KafkaStreams(builder.build(), props);
streams.start();
```

**Idempotent Processing:**
```java
public class IdempotentProcessor extends KeyedProcessFunction<String, Event, Result> {
    
    private ValueState<ProcessedEventIds> processedIdsState;
    
    @Override
    public void open(Configuration parameters) {
        ValueStateDescriptor<ProcessedEventIds> desc = 
            new ValueStateDescriptor<>("processedIds", ProcessedEventIds.class);
        processedIdsState = getRuntimeContext().getState(desc);
    }
    
    @Override
    public void processElement(Event event, Context context, 
                               Collector<Result> out) throws Exception {
        ProcessedEventIds processedIds = processedIdsState.value();
        if (processedIds == null) {
            processedIds = new ProcessedEventIds();
        }
        
        // Check if already processed
        if (processedIds.contains(event.getId())) {
            return; // Skip duplicate
        }
        
        // Process event
        Result result = processEvent(event);
        out.collect(result);
        
        // Mark as processed
        processedIds.add(event.getId());
        processedIdsState.update(processedIds);
    }
}
```

---

## 7. Backpressure & Flow Control

### Q10: How do you handle backpressure in streaming systems?

**Answer:**

**Reactive Streams with Backpressure:**
```java
// Project Reactor implementation
public class ReactiveStreamProcessor {
    
    public Flux<Result> processEvents(Flux<Event> eventFlux) {
        return eventFlux
            .onBackpressureBuffer(1000, BufferOverflowStrategy.DROP_OLDEST)
            .flatMap(this::processEvent, 10)  // Concurrency of 10
            .onBackpressureLatest()  // Keep only latest if overwhelmed
            .timeout(Duration.ofSeconds(30))  // Timeout for processing
            .retry(3)  // Retry on failure
            .doOnError(error -> log.error("Processing error", error))
            .onErrorResume(error -> Flux.empty());  // Continue on error
    }
    
    private Mono<Result> processEvent(Event event) {
        return Mono.fromCallable(() -> heavyProcessing(event))
            .subscribeOn(Schedulers.parallel());
    }
}
```

**Flink Backpressure Handling:**
```java
// Configure network buffers for backpressure
env.getConfig().setNetworkBuffersMemory(33554432); // 32MB
env.getConfig().setTaskNetworkTimeout(10000); // 10 seconds

// Custom operator with backpressure awareness
public class BackpressureAwareOperator extends AbstractStreamOperator<Event> {
    
    @Override
    public void processElement(StreamRecord<Event> element) throws Exception {
        // Check available buffers
        if (getRuntimeContext().getNumberOfInputChannels() > 0) {
            // Slow down if downstream is overwhelmed
            Thread.sleep(10);  // Simple throttling
        }
        
        // Process element
        Event event = element.getValue();
        Event processedEvent = transform(event);
        
        // Output with backpressure consideration
        output.collect(new StreamRecord<>(processedEvent));
    }
}
```

**Kafka Consumer Backpressure:**
```java
public class BackpressureAwareKafkaConsumer {
    
    private final KafkaConsumer<String, Event> consumer;
    private final BlockingQueue<Event> eventQueue;
    private final AtomicBoolean isPaused = new AtomicBoolean(false);
    
    public void startConsuming() {
        Thread consumerThread = new Thread(() -> {
            while (true) {
                if (eventQueue.size() > MAX_QUEUE_SIZE) {
                    // Pause consumption if queue is full
                    if (isPaused.compareAndSet(false, true)) {
                        consumer.pause(consumer.assignment());
                    }
                    
                    // Wait for queue to drain
                    Thread.sleep(1000);
                    continue;
                }
                
                // Resume consumption
                if (isPaused.compareAndSet(true, false)) {
                    consumer.resume(consumer.assignment());
                }
                
                // Poll for messages
                ConsumerRecords<String, Event> records = consumer.poll(Duration.ofMillis(100));
                
                for (ConsumerRecord<String, Event> record : records) {
                    eventQueue.offer(record.value());
                }
            }
        });
        
        consumerThread.start();
    }
}
```

---

## 8. Real-time Analytics

### Q11: How do you implement real-time analytics dashboards?

**Answer:**

**Real-time Metrics Aggregation:**
```java
public class RealTimeMetricsProcessor extends KeyedProcessFunction<String, Event, Metric> {
    
    private ValueState<MetricAccumulator> metricState;
    
    @Override
    public void open(Configuration parameters) {
        ValueStateDescriptor<MetricAccumulator> desc = 
            new ValueStateDescriptor<>("metrics", MetricAccumulator.class);
        metricState = getRuntimeContext().getState(desc);
    }
    
    @Override
    public void processElement(Event event, Context context, 
                               Collector<Metric> out) throws Exception {
        
        MetricAccumulator accumulator = metricState.value();
        if (accumulator == null) {
            accumulator = new MetricAccumulator();
        }
        
        // Update metrics
        accumulator.addEvent(event);
        
        // Emit metrics every second
        long nextEmitTime = (context.timestamp() / 1000) * 1000 + 1000;
        context.timerService().registerEventTimeTimer(nextEmitTime);
        
        metricState.update(accumulator);
    }
    
    @Override
    public void onTimer(long timestamp, OnTimerContext ctx, 
                        Collector<Metric> out) throws Exception {
        
        MetricAccumulator accumulator = metricState.value();
        if (accumulator != null) {
            // Calculate metrics for the time window
            Metric metric = new Metric(
                ctx.getCurrentKey(),
                timestamp,
                accumulator.getCount(),
                accumulator.getSum(),
                accumulator.getAverage(),
                accumulator.getMax(),
                accumulator.getMin()
            );
            
            out.collect(metric);
            
            // Reset accumulator for next window
            accumulator.reset();
            metricState.update(accumulator);
        }
    }
}
```

**WebSocket Dashboard Updates:**
```java
@RestController
public class RealTimeDashboardController {
    
    private final SimpMessagingTemplate messagingTemplate;
    private final Map<String, Subscription> subscriptions = new ConcurrentHashMap<>();
    
    @MessageMapping("/subscribe")
    public void subscribeToMetrics(SubscriptionRequest request, 
                                   SimpMessageHeaderAccessor headerAccessor) {
        String sessionId = headerAccessor.getSessionId();
        
        // Subscribe to real-time metrics
        Subscription subscription = new Subscription(
            sessionId, 
            request.getMetricType(),
            request.getFilters()
        );
        
        subscriptions.put(sessionId, subscription);
        
        // Send initial data
        sendInitialMetrics(sessionId, subscription);
    }
    
    public void broadcastMetricUpdate(Metric metric) {
        // Find matching subscriptions
        subscriptions.values().stream()
            .filter(sub -> sub.matchesMetric(metric))
            .forEach(sub -> {
                String destination = "/topic/metrics/" + sub.getSessionId();
                messagingTemplate.convertAndSend(destination, metric);
            });
    }
}
```

**Time Series Database Integration:**
```java
public class TimeSeriesSink extends RichSinkFunction<Metric> {
    
    private InfluxDB influxDB;
    
    @Override
    public void open(Configuration parameters) throws Exception {
        influxDB = InfluxDBFactory.connect("http://localhost:8086");
        influxDB.setDatabase("metrics");
    }
    
    @Override
    public void invoke(Metric metric, Context context) throws Exception {
        Point point = Point.measurement("real_time_metrics")
            .tag("key", metric.getKey())
            .tag("type", metric.getType())
            .time(metric.getTimestamp(), TimeUnit.MILLISECONDS)
            .addField("count", metric.getCount())
            .addField("sum", metric.getSum())
            .addField("average", metric.getAverage())
            .addField("max", metric.getMax())
            .addField("min", metric.getMin())
            .build();
        
        influxDB.write(point);
    }
    
    @Override
    public void close() throws Exception {
        if (influxDB != null) {
            influxDB.close();
        }
    }
}
```

---

## 9. Stream Processing Patterns

### Q12: What are common stream processing patterns?

**Answer:**

**Pattern 1: Filtering and Routing:**
```java
// Route events to different streams based on conditions
public class EventRouter extends KeyedProcessFunction<String, Event, RoutedEvent> {
    
    @Override
    public void processElement(Event event, Context context, 
                               Collector<RoutedEvent> out) throws Exception {
        
        String route = determineRoute(event);
        
        RoutedEvent routedEvent = new RoutedEvent(event, route);
        out.collect(routedEvent);
        
        // Also output to side streams for specific processing
        if (event.isHighPriority()) {
            context.output(highPriorityTag, routedEvent);
        }
        
        if (event.isSuspicious()) {
            context.output(fraudTag, routedEvent);
        }
    }
    
    private String determineRoute(Event event) {
        if (event.getAmount() > 10000) {
            return "HIGH_VALUE";
        } else if (event.isInternational()) {
            return "INTERNATIONAL";
        } else {
            return "STANDARD";
        }
    }
}
```

**Pattern 2: Enrichment:**
```java
// Enrich events with reference data
public class EventEnricher extends KeyedProcessFunction<String, Event, EnrichedEvent> {
    
    private ValueState<UserProfile> userProfileState;
    private ValueState<DeviceProfile> deviceProfileState;
    
    @Override
    public void processElement(Event event, Context context, 
                               Collector<EnrichedEvent> out) throws Exception {
        
        // Get user profile
        UserProfile userProfile = userProfileState.value();
        if (userProfile == null) {
            userProfile = fetchUserProfile(event.getUserId());
            userProfileState.update(userProfile);
        }
        
        // Get device profile
        DeviceProfile deviceProfile = deviceProfileState.value();
        if (deviceProfile == null) {
            deviceProfile = fetchDeviceProfile(event.getDeviceId());
            deviceProfileState.update(deviceProfile);
        }
        
        // Create enriched event
        EnrichedEvent enriched = new EnrichedEvent(event, userProfile, deviceProfile);
        out.collect(enriched);
    }
}
```

**Pattern 3: Aggregation and Windowing:**
```java
// Complex windowed aggregation
public class SessionAggregator extends ProcessWindowFunction<Event, SessionSummary, 
                                                           String, TimeWindow> {
    
    @Override
    public void process(String key, Context context, 
                       Iterable<Event> elements, 
                       Collector<SessionSummary> out) throws Exception {
        
        List<Event> events = new ArrayList<>();
        double totalAmount = 0.0;
        int eventCount = 0;
        
        for (Event event : elements) {
            events.add(event);
            totalAmount += event.getAmount();
            eventCount++;
        }
        
        // Calculate session metrics
        SessionSummary summary = new SessionSummary(
            key,
            context.window().getStart(),
            context.window().getEnd(),
            eventCount,
            totalAmount,
            calculateAverageEventInterval(events),
            detectAnomalies(events)
        );
        
        out.collect(summary);
    }
}
```

**Pattern 4: Anomaly Detection:**
```java
// Real-time anomaly detection
public class AnomalyDetector extends KeyedProcessFunction<String, Event, Anomaly> {
    
    private ValueState<StatisticalModel> modelState;
    private ListState<Event> recentEventsState;
    
    @Override
    public void processElement(Event event, Context context, 
                               Collector<Anomaly> out) throws Exception {
        
        StatisticalModel model = modelState.value();
        if (model == null) {
            model = new StatisticalModel();
            modelState.update(model);
        }
        
        // Update model with new event
        model.update(event);
        
        // Check for anomalies
        double anomalyScore = model.calculateAnomalyScore(event);
        
        if (anomalyScore > ANOMALY_THRESHOLD) {
            Anomaly anomaly = new Anomaly(
                event,
                anomalyScore,
                model.getExpectedValue(event),
                "Statistical anomaly detected"
            );
            out.collect(anomaly);
        }
        
        // Store recent events for pattern analysis
        recentEventsState.add(event);
        cleanupOldEvents(context.timestamp());
    }
}
```

---

## 10. Common Interview Questions

### Q13: Design a real-time fraud detection system for a payment processor

**Answer:**
"I'll design a multi-layered real-time fraud detection system:

**Architecture:**
```
Payment Events → Kafka → Stream Processor → Rule Engine → ML Model → Decision Engine → Alert System
       │              │            │            │           │              │
       ▼              ▼            ▼            ▼           ▼              ▼
   Transaction    Flink/Spark   Drools/      TensorFlow   Risk Scoring   Notification
   Data           Processing   Esper         Serving     Engine         Service
```

**Components:**

1. **Ingestion Layer:**
```java
// High-throughput transaction ingestion
public class TransactionIngestion {
    
    @KafkaListener(topics = "transactions", concurrency = "10")
    public void processTransaction(Transaction transaction) {
        // Validate transaction format
        validateTransaction(transaction);
        
        // Add processing timestamp
        transaction.setProcessingTimestamp(System.currentTimeMillis());
        
        // Route to processing pipeline
        streamProcessor.process(transaction);
    }
}
```

2. **Stream Processing with Flink:**
```java
public class FraudDetectionPipeline {
    
    public void setupPipeline() {
        DataStream<Transaction> transactions = env.addSource(kafkaSource);
        
        // Multiple detection strategies in parallel
        DataStream<FraudSignal> ruleBasedSignals = transactions
            .keyBy(Transaction::getUserId)
            .process(new RuleBasedDetection());
        
        DataStream<FraudSignal> mlSignals = transactions
            .keyBy(Transaction::getUserId)
            .process(new MLDetection());
        
        DataStream<FraudSignal> patternSignals = transactions
            .keyBy(Transaction::getMerchantId)
            .process(new PatternDetection());
        
        // Combine all signals
        DataStream<FraudDecision> decisions = ruleBasedSignals
            .union(mlSignals)
            .union(patternSignals)
            .keyBy(FraudSignal::getTransactionId)
            .window(TumblingProcessingTimeWindows.of(Time.seconds(5)))
            .process(new DecisionEngine());
        
        // Output decisions
        decisions.addSink(new DecisionSink());
    }
}
```

3. **Rule-Based Detection:**
```java
public class RuleBasedDetection extends KeyedProcessFunction<String, Transaction, FraudSignal> {
    
    private ValueState<TransactionHistory> historyState;
    private ValueState<RiskProfile> riskProfileState;
    
    @Override
    public void processElement(Transaction transaction, Context context, 
                               Collector<FraudSignal> out) throws Exception {
        
        List<FraudSignal> signals = new ArrayList<>();
        
        // Rule 1: Velocity check
        TransactionHistory history = historyState.value();
        if (history != null && history.getRecentTransactionCount(5) > 10) {
            signals.add(new FraudSignal(transaction, "HIGH_VELOCITY", 0.8));
        }
        
        // Rule 2: Amount anomaly
        RiskProfile riskProfile = riskProfileState.value();
        if (riskProfile != null && 
            transaction.getAmount() > riskProfile.getMaxAmount() * 3) {
            signals.add(new FraudSignal(transaction, "AMOUNT_ANOMALY", 0.7));
        }
        
        // Rule 3: Geographic anomaly
        if (isGeographicAnomaly(transaction, history)) {
            signals.add(new FraudSignal(transaction, "GEO_ANOMALY", 0.6));
        }
        
        // Output all signals
        signals.forEach(out::collect);
        
        // Update state
        updateHistory(transaction);
        updateRiskProfile(transaction);
    }
}
```

4. **ML Model Serving:**
```java
public class MLDetection extends KeyedProcessFunction<String, Transaction, FraudSignal> {
    
    private TensorFlowService mlService;
    private ValueState<UserFeatures> featuresState;
    
    @Override
    public void processElement(Transaction transaction, Context context, 
                               Collector<FraudSignal> out) throws Exception {
        
        // Extract features
        UserFeatures features = extractFeatures(transaction);
        
        // Get ML prediction
        FraudPrediction prediction = mlService.predict(features);
        
        if (prediction.getFraudProbability() > 0.5) {
            FraudSignal signal = new FraudSignal(
                transaction,
                "ML_DETECTION",
                prediction.getFraudProbability(),
                prediction.getExplanation()
            );
            out.collect(signal);
        }
        
        // Update features for next prediction
        featuresState.update(features);
    }
}
```

5. **Decision Engine:**
```java
public class DecisionEngine extends ProcessWindowFunction<FraudSignal, FraudDecision, 
                                                         String, TimeWindow> {
    
    @Override
    public void process(String transactionId, Context context, 
                       Iterable<FraudSignal> signals, 
                       Collector<FraudDecision> out) throws Exception {
        
        List<FraudSignal> signalList = new ArrayList<>();
        signals.forEach(signalList::add);
        
        // Calculate combined risk score
        double combinedScore = calculateCombinedRiskScore(signalList);
        
        // Make decision
        Decision decision = makeDecision(combinedScore, signalList);
        
        FraudDecision fraudDecision = new FraudDecision(
            transactionId,
            decision,
            combinedScore,
            signalList,
            context.window().getEnd()
        );
        
        out.collect(fraudDecision);
    }
    
    private Decision makeDecision(double score, List<FraudSignal> signals) {
        if (score > 0.8) {
            return Decision.BLOCK;
        } else if (score > 0.5) {
            return Decision.MANUAL_REVIEW;
        } else {
            return Decision.APPROVE;
        }
    }
}
```

**Scaling Considerations:**
- **Parallel Processing:** Multiple Flink task managers
- **State Management:** RocksDB for large state
- **Model Updates:** Hot-swapping ML models
- **Monitoring:** Real-time metrics and alerting"

### Q14: How would you handle out-of-order events in a real-time analytics system?

**Answer:**
"Handling out-of-order events requires careful time semantics and watermarks:

**Watermark Strategy:**
```java
// Allow up to 30 seconds of lateness
WatermarkStrategy<Event> watermarkStrategy = WatermarkStrategy
    .<Event>forBoundedOutOfOrderness(Duration.ofSeconds(30))
    .withTimestampAssigner((event, timestamp) -> event.getEventTimestamp())
    .withIdleness(Duration.ofMinutes(2));

DataStream<Event> stream = env.addSource(kafkaSource)
    .assignTimestampsAndWatermarks(watermarkStrategy);
```

**Late Data Handling:**
```java
// Side output for late events
OutputTag<Event> lateEventsTag = new OutputTag<Event>("late-events") {};

DataStream<WindowedResult> mainResults = stream
    .keyBy(Event::getUserId)
    .window(TumblingEventTimeWindows.of(Time.minutes(1)))
    .sideOutputLateData(lateEventsTag)
    .aggregate(new CountAggregate());

// Process late events separately
DataStream<Event> lateEvents = mainResults.getSideOutput(lateEventsTag);
lateEvents.process(new LateEventHandler());

// Allow late data to update results
DataStream<UpdatedResult> updatedResults = stream
    .keyBy(Event::getUserId)
    .window(TumblingEventTimeWindows.of(Time.minutes(1)))
    .allowedLateness(Time.minutes(5))  // Allow 5 minutes of lateness
    .aggregate(new UpdatingAggregate());
```

**Event Reconciliation:**
```java
public class EventReconciliation {
    
    private final Map<String, WindowedResult> windowResults = new ConcurrentHashMap<>();
    
    public void handleOnTimeEvent(WindowedResult result) {
        // Store initial result
        windowResults.put(result.getWindowId(), result);
        
        // Send to downstream
        sendDownstream(result);
    }
    
    public void handleLateEvent(Event event, String windowId) {
        WindowedResult existingResult = windowResults.get(windowId);
        if (existingResult != null) {
            // Recalculate result with late event
            WindowedResult updatedResult = recalculateResult(existingResult, event);
            
            // Update stored result
            windowResults.put(windowId, updatedResult);
            
            // Send correction downstream
            sendCorrection(updatedResult);
        }
    }
    
    private WindowedResult recalculateResult(WindowedResult original, Event lateEvent) {
        // Add late event to aggregation
        long newCount = original.getCount() + 1;
        double newSum = original.getSum() + lateEvent.getValue();
        double newAverage = newSum / newCount;
        
        return new WindowedResult(
            original.getWindowId(),
            newCount,
            newSum,
            newAverage,
            original.getMax(),
            original.getMin(),
            true  // Mark as updated
        );
    }
}
```

**Idempotent Processing:**
```java
public class IdempotentAggregator extends ProcessWindowFunction<Event, Result, 
                                                           String, TimeWindow> {
    
    @Override
    public void process(String key, Context context, 
                       Iterable<Event> events, 
                       Collector<Result> out) throws Exception {
        
        Set<String> processedEventIds = new HashSet<>();
        double sum = 0.0;
        long count = 0;
        
        for (Event event : events) {
            // Skip duplicates
            if (processedEventIds.contains(event.getId())) {
                continue;
            }
            
            processedEventIds.add(event.getId());
            sum += event.getValue();
            count++;
        }
        
        Result result = new Result(key, context.window(), count, sum, sum / count);
        out.collect(result);
    }
}
```"

### Q15: Design a real-time recommendation system for a streaming platform

**Answer:**
"I'll design a multi-stage real-time recommendation system:

**Architecture:**
```
User Events → Feature Store → Real-time Scoring → Candidate Generation → Ranking → Personalization
     │              │              │               │              │           │
     ▼              ▼              ▼               ▼              ▼           ▼
  Clicks,       Redis,         ML Models,      Content-based,  Neural      Business
  Views,        Feature        TensorFlow      Collaborative,  Networks,   Rules
  Likes         Store          Serving         Filtering,      Transformers Diversity
```

**Implementation:**

1. **Real-time Feature Processing:**
```java
public class RealTimeFeatureProcessor extends KeyedProcessFunction<String, UserEvent, UserFeatures> {
    
    private ValueState<UserProfile> profileState;
    private ValueState<InteractionHistory> historyState;
    
    @Override
    public void processElement(UserEvent event, Context context, 
                               Collector<UserFeatures> out) throws Exception {
        
        // Update user profile
        UserProfile profile = updateProfile(event);
        
        // Calculate real-time features
        UserFeatures features = UserFeatures.builder()
            .userId(event.getUserId())
            .recentGenres(getRecentGenres(historyState.value()))
            .timeOfDay(getTimeOfDay(event.getTimestamp()))
            .deviceType(event.getDeviceType())
            .sessionDuration(calculateSessionDuration(event))
            .engagementScore(calculateEngagementScore(profile))
            .build();
        
        out.collect(features);
        
        // Update state
        profileState.update(profile);
        updateHistory(event);
    }
}
```

2. **Candidate Generation:**
```java
public class CandidateGenerator extends KeyedProcessFunction<String, UserFeatures, Candidate> {
    
    private BroadcastService<ContentCatalog> contentCatalog;
    
    @Override
    public void processElement(UserFeatures features, Context context, 
                               Collector<Candidate> out) throws Exception {
        
        List<Candidate> candidates = new ArrayList<>();
        
        // Content-based filtering
        List<Candidate> contentCandidates = generateContentBasedCandidates(features);
        candidates.addAll(contentCandidates);
        
        // Collaborative filtering (real-time)
        List<Candidate> collaborativeCandidates = generateCollaborativeCandidates(features);
        candidates.addAll(collaborativeCandidates);
        
        // Trending content
        List<Candidate> trendingCandidates = getTrendingContent(features);
        candidates.addAll(trendingCandidates);
        
        // Remove duplicates and limit
        candidates.stream()
            .distinct()
            .limit(100)
            .forEach(out::collect);
    }
    
    private List<Candidate> generateContentBasedCandidates(UserFeatures features) {
        return contentCatalog.getBroadcastValue()
            .getContentByGenres(features.getRecentGenres())
            .stream()
            .map(content -> new Candidate(content, calculateContentScore(content, features)))
            .collect(Collectors.toList());
    }
}
```

3. **Real-time Ranking:**
```java
public class RealTimeRanker extends KeyedProcessFunction<String, Candidate, RankedContent> {
    
    private TensorFlowService rankingModel;
    
    @Override
    public void processElement(Candidate candidate, Context context, 
                               Collector<RankedContent> out) throws Exception {
        
        // Extract features for ranking
        RankingFeatures features = extractRankingFeatures(candidate);
        
        // Get model prediction
        RankingPrediction prediction = rankingModel.predict(features);
        
        // Create ranked content
        RankedContent ranked = new RankedContent(
            candidate.getContent(),
            prediction.getScore(),
            prediction.getConfidence(),
            prediction.getExplanation()
        );
        
        out.collect(ranked);
    }
    
    private RankingFeatures extractRankingFeatures(Candidate candidate) {
        return RankingFeatures.builder()
            .contentFeatures(candidate.getContent().getFeatures())
            .userFeatures(candidate.getUserFeatures())
            .contextFeatures(candidate.getContextFeatures())
            .interactionFeatures(candidate.getInteractionFeatures())
            .build();
    }
}
```

4. **Personalization and Diversity:**
```java
public class PersonalizationEngine extends ProcessWindowFunction<RankedContent, 
                                                              Recommendation, 
                                                              String, 
                                                              TimeWindow> {
    
    @Override
    public void process(String userId, Context context, 
                       Iterable<RankedContent> contents, 
                       Collector<Recommendation> out) throws Exception {
        
        List<RankedContent> contentList = new ArrayList<>();
        contents.forEach(contentList::add);
        
        // Sort by score
        contentList.sort((a, b) -> Double.compare(b.getScore(), a.getScore()));
        
        // Apply diversity constraints
        List<RankedContent> diversified = applyDiversity(contentList);
        
        // Apply business rules
        List<RankedContent> filtered = applyBusinessRules(diversified);
        
        // Create final recommendation
        Recommendation recommendation = new Recommendation(
            userId,
            filtered.stream().limit(20).collect(Collectors.toList()),
            context.window().getEnd(),
            generateExplanation(filtered)
        );
        
        out.collect(recommendation);
    }
    
    private List<RankedContent> applyDiversity(List<RankedContent> contents) {
        List<RankedContent> result = new ArrayList<>();
        Set<String> genres = new HashSet<>();
        
        for (RankedContent content : contents) {
            String genre = content.getContent().getGenre();
            
            // Ensure genre diversity
            if (!genres.contains(genre) || result.size() < 5) {
                result.add(content);
                genres.add(genre);
                
                if (result.size() >= 20) {
                    break;
                }
            }
        }
        
        return result;
    }
}
```

**Performance Optimization:**
- **Feature Pre-computation:** Background jobs for static features
- **Model Caching:** In-memory model serving for low latency
- **Parallel Processing:** Multiple ranking models in parallel
- **A/B Testing:** Multiple recommendation strategies"

---

## 🎯 Quick Reference

### Key Streaming Concepts
- **Time Semantics:** Event time, processing time, watermarks
- **Windowing:** Tumbling, sliding, session, custom windows
- **State Management:** Keyed state, operator state, TTL
- **Processing Guarantees:** At-least-once, exactly-once
- **Backpressure:** Flow control and rate limiting

### Common Technologies
- **Stream Processors:** Apache Flink, Spark Streaming, Kafka Streams
- **Message Brokers:** Apache Kafka, Apache Pulsar, AWS Kinesis
- **State Stores:** RocksDB, Redis, Cassandra
- **Databases:** InfluxDB, TimescaleDB, ClickHouse

### Interview Focus Areas
- **System design:** End-to-end streaming architectures
- **State management:** Checkpointing, recovery, scaling
- **Performance:** Latency optimization, backpressure handling
- **Reliability:** Exactly-once, fault tolerance, monitoring
- **Patterns:** Joins, aggregations, enrichment, routing
