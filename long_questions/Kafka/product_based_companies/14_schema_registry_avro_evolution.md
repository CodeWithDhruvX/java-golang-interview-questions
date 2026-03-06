# 📐 Kafka — Schema Registry & Avro Schema Evolution (Deep Dive)

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Razorpay, PhonePe, Swiggy, Netflix, Confluent, LinkedIn

---

## Q1. What are Schema Registry compatibility modes?

"The Schema Registry enforces a **compatibility check** every time a new schema version is registered to prevent producers and consumers from silently breaking each other.

| Mode | Rule | Safe Change |
|---|---|---|
| `BACKWARD` | New schema can read messages written with old schema | Add optional field with default |
| `FORWARD` | Old schema can read messages written with new schema | Remove optional field |
| `FULL` | Both BACKWARD + FORWARD | Add/remove optional fields with defaults only |
| `BACKWARD_TRANSITIVE` | Compatible with ALL previous versions (not just last) | Additive changes only |
| `FULL_TRANSITIVE` | All-version FULL — enterprise standard | Strictly additive |
| `NONE` | No enforcement — dangerous in production | Anything |

**My production recommendation:** Use `FULL_TRANSITIVE` for shared topics. Use `BACKWARD` for single-team topics with coordinated deployments."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Kafka integrates seamlessly with Confluent Schema Registry using `KafkaAvroSerializer`. The compatibility checks happen automatically during the serialization phase before a message is sent to the broker.
* **Golang:** The `confluent-kafka-go` library provides an `sr` (Schema Registry) package. You must manually instantiate a `schemaregistry.Client` to register schemas and wrap your producer/consumer with an `AvroSerializer`/`AvroDeserializer`.

#### Indepth
`BACKWARD_TRANSITIVE` vs `BACKWARD`: `BACKWARD` only checks v3 vs v2. A consumer still on v1 (skipped v2) reading v3 messages may fail. `BACKWARD_TRANSITIVE` guarantees v3 is readable by consumers on ANY older version — critical in environments where consumers and producers don't upgrade together.

---

## Q2. Show a complete Avro schema evolution example.

"**Version 1 — Initial Schema:**
```json
{
  "type": "record", "name": "Order", "namespace": "com.mycompany.events",
  "fields": [
    {"name": "orderId", "type": "string"},
    {"name": "userId",  "type": "string"},
    {"name": "amount",  "type": "double"},
    {"name": "status",  "type": "string"}
  ]
}
```

**Version 2 — Add optional `currency` field (BACKWARD COMPATIBLE):**
```json
{
  "type": "record", "name": "Order", "namespace": "com.mycompany.events",
  "fields": [
    {"name": "orderId",  "type": "string"},
    {"name": "userId",   "type": "string"},
    {"name": "amount",   "type": "double"},
    {"name": "status",   "type": "string"},
    {
      "name": "currency",
      "type": ["null", "string"],   // union with null = optional
      "default": null               // MUST have default for backward compat
    }
  ]
}
```

**Why v2 is backward compatible:**
- Old consumers (v1) reading a v2 message → Avro **ignores unknown fields** ✅
- New consumers (v2) reading a v1 message → `currency` missing → Avro uses `default: null` ✅

**Version 3 — Add nested `deliveryAddress` (still BACKWARD COMPATIBLE):**
```json
{
  "name": "deliveryAddress",
  "type": ["null", {
    "type": "record", "name": "Address",
    "fields": [
      {"name": "street",  "type": "string"},
      {"name": "city",    "type": "string"},
      {"name": "pincode", "type": "string"}
    ]
  }],
  "default": null
}
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Avro schemas (`.avsc` files) are typically compiled into Java POJOs using the `avro-maven-plugin` during the build phase. This ensures strict type safety at compile time.
* **Golang:** Go doesn't inherently support Avro POJOs like Java does. Third-party libraries like `hamba/avro` or `actgardner/gogen-avro` are used to generate Go structs from `.avsc` files, providing type-safe serialization.

#### Indepth
**Common mistake:** Adding `{"name": "discountCode", "type": "string"}` with no default and no null union. This is NOT backward compatible. Old consumers reading a new message have no default for `discountCode` and will throw a deserialization error. The Schema Registry with `BACKWARD` mode will **reject** this registration — which is exactly the safety net it provides.

---

## Q3. Show complete Java producer/consumer code with Schema Registry.

"**Dependencies:**
```xml
<dependency>
    <groupId>io.confluent</groupId>
    <artifactId>kafka-avro-serializer</artifactId>
    <version>7.6.0</version>
</dependency>
```

**Producer:**
```java
Properties props = new Properties();
props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka:9092");
props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, KafkaAvroSerializer.class);
props.put(AbstractKafkaSchemaSerDeConfig.SCHEMA_REGISTRY_URL_CONFIG, "http://schema-registry:8081");
props.put(KafkaAvroSerializerConfig.AUTO_REGISTER_SCHEMAS, false); // PROD: disable auto-register

KafkaProducer<String, Order> producer = new KafkaProducer<>(props);

Order order = Order.newBuilder()
    .setOrderId("ORD-123456")
    .setUserId("USR-789")
    .setAmount(1299.99)
    .setStatus("PLACED")
    .setCurrency("INR")
    .setDeliveryAddress(null)  // null is fine (nullable union)
    .build();

producer.send(new ProducerRecord<>("order-events", order.getOrderId(), order));
```

**Consumer (v3 consumer reading v1 messages — evolution handled automatically):**
```java
Properties props = new Properties();
props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka:9092");
props.put(ConsumerConfig.GROUP_ID_CONFIG, "order-processor");
props.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
props.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, KafkaAvroDeserializer.class);
props.put(AbstractKafkaSchemaSerDeConfig.SCHEMA_REGISTRY_URL_CONFIG, "http://schema-registry:8081");
props.put(KafkaAvroDeserializerConfig.SPECIFIC_AVRO_READER_CONFIG, true);

KafkaConsumer<String, Order> consumer = new KafkaConsumer<>(props);
consumer.subscribe(List.of("order-events"));

while (true) {
    ConsumerRecords<String, Order> records = consumer.poll(Duration.ofMillis(100));
    for (ConsumerRecord<String, Order> record : records) {
        Order order = record.value();
        // Safely handle evolution — fields from v2/v3 may be null for old messages
        String currency = order.getCurrency() != null ? order.getCurrency().toString() : "INR";
        Address addr = order.getDeliveryAddress(); // null for v1/v2 messages
        processOrder(order, currency, addr);
    }
    consumer.commitSync();
}
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Configuration properties like `schema.registry.url` and `specific.avro.reader` are effortlessly mapped in `application.yml` and automatically injected into the `KafkaTemplate` and `@KafkaListener` configurations.
* **Golang:** The `AvroDeserializer` requires explicitly passing the deserialization target struct as a reference (`&MyEvent{}`) during the `DeserializeInto` call, managing the type mapping manually.

#### Indepth
`auto.register.schemas=false` in production is mandatory. If true, a developer pushing broken code could silently register an incompatible schema, corrupting the topic for all consumers. In production, schema registration happens through CI/CD (`maven-schema-registry-mojo`) after explicit compatibility checks, not at runtime.

---

## Q4. How do you handle a BREAKING schema change?

"**Scenario:** Changing `amount` from `double` to `decimal` — a breaking type change with no backward-compatible path.

**Strategy: Topic Migration with Dual-Write Period**

```text
Day 1: Deploy producer with dual-write to both old + new topic
Day 3: All consumer teams migrated to new topic
Day 7: No lag on any consumer group for new topic  
Day 8: Disable dual-write, set old topic retention.ms = 3600000 (1 hour)
Day 9: Old topic deleted
```

**Step 1 — Create new versioned topic + register schema:**
```bash
kafka-topics.sh --create --topic order-events-v2 --partitions 24 --bootstrap-server kafka:9092
curl -X POST http://schema-registry:8081/subjects/order-events-v2-value/versions \
  -H 'Content-Type: application/json' -d @order-new-schema.json
```

**Step 2 — Dual-write producer:**
```java
producer.send(new ProducerRecord<>("order-events",    orderId, oldOrder));   // v1 topic
producer.send(new ProducerRecord<>("order-events-v2", orderId, newOrder));   // v2 topic
```

**Alternative for field renames only — use Avro `aliases`:**
```json
{
  "name": "totalAmount",
  "type": "double",
  "aliases": ["amount"]   // old field name → new field name, backward compatible
}
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Both ecosystems handle dual-write periods cleanly by instantiating two distinct producers or creating a router function that publishes the old struct to topic V1 and the new mapped struct to topic V2 simultaneously.

#### Indepth
The `aliases` feature avoids a full topic migration for simple renames. Avro automatically maps the old wire-format field `amount` to the new in-memory field `totalAmount` during deserialization. No dual-write, no migration window needed.

---

## Q5. How do you manage schemas in CI/CD?

"**Schema-as-Code with Maven plugin:**
```xml
<plugin>
    <groupId>io.confluent</groupId>
    <artifactId>kafka-schema-registry-maven-plugin</artifactId>
    <version>7.6.0</version>
    <configuration>
        <schemaRegistryUrls><param>http://schema-registry:8081</param></schemaRegistryUrls>
        <subjects>
            <order-events-value>src/main/avro/order-v3.avsc</order-events-value>
        </subjects>
        <compatibilityLevels>
            <order-events-value>FULL_TRANSITIVE</order-events-value>
        </compatibilityLevels>
    </configuration>
</plugin>
```

**GitHub Actions pipeline:**
```yaml
- name: Check Schema Compatibility (PR gate)
  run: mvn kafka-schema-registry:test-compatibility
  # Fails PR if new schema breaks compatibility with any registered version

- name: Register Schema (on merge to main only)
  if: github.ref == 'refs/heads/main'
  run: mvn kafka-schema-registry:register
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Schema validation is often baked into the CI/CD pipeline using the Confluent Maven/Gradle plugins.
* **Golang:** Schema validation is usually performed by invoking the Schema Registry REST API directly from CI scripts or using the Confluent CLI, as Go lacks an equivalent native build-tool plugin for this.

#### Indepth
Run `test-compatibility` against the **production** Schema Registry URL in CI — not just dev. Catching production compatibility breakage before deployment is the entire point. Use separate subject naming per environment: `order-events-dev-value`, `order-events-prod-value`.

---

## Q6. How do you handle Poison Pill messages?

"A **Poison Pill** is an undeserializable message — schema not found, bytes corrupted, or wrong format produced.

**Kafka Streams — built-in handler:**
```java
config.put(
    StreamsConfig.DEFAULT_DESERIALIZATION_EXCEPTION_HANDLER_CLASS_CONFIG,
    LogAndContinueExceptionHandler.class  // Log + skip bad message
);
```

**Plain KafkaConsumer — seek past the bad offset:**
```java
try {
    records = consumer.poll(Duration.ofMillis(100));
} catch (SerializationException e) {
    log.error("Poison pill detected", e);
    // Seek past the bad message
    for (TopicPartition tp : consumer.assignment()) {
        consumer.seek(tp, consumer.position(tp) + 1);
    }
    sendToDLQ(e);   // Route raw bytes to Dead Letter Queue
    continue;
}
```

**DLQ producer for poison pills:**
```java
// Write raw bytes + error metadata to DLQ topic for forensics
dlqProducer.send(new ProducerRecord<>(
    "order-events-dlq",
    "poison-pill-" + System.currentTimeMillis(),
    rawBytes
));
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Kafka provides robust `ErrorHandlingDeserializer` wrappers that catch `SerializationException`s, allowing you to gracefully route poison pills to a DLQ without crashing the consumer loop.
* **Golang:** Poison pills manifest as errors returned by `Deserialize` or `ReadMessage`. Go's explicit error handling makes it straightforward: `if err != nil { publishToDLQ(rawBytes, err) }`.

#### Indepth
**Schema Registry 404 as poison pill:** If a producer registers a schema on a registry the consumer can't reach (wrong environment URL), the consumer gets a `404` fetching the schema ID → treated as deserialization failure. Include a `GET /subjects` health check against the Schema Registry in your service readiness probe to catch misconfiguration at startup.
