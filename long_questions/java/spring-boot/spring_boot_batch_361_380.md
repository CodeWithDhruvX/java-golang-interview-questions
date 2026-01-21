## 🔹 Section 7: Spring Boot with NoSQL (361-380)

### Question 361: How do you configure MongoDB with Spring Boot?

**Answer:**
`spring-boot-starter-data-mongodb`.
Props: `spring.data.mongodb.uri=mongodb://localhost:27017/db`.
Spring auto-configures `MongoTemplate` and `MongoRepository`.

---

### Question 362: How does Spring Data MongoDB support query methods?

**Answer:**
Like JPA: `findByUsername(String user)`.
Start with `findBy...`, `deleteBy...`.
Supports JSON-based queries: `@Query("{ 'name': ?0 }")`.

---

### Question 363: How do you define compound indexes in MongoDB using Spring annotations?

**Answer:**
Annotate Class:
`@CompoundIndex(def = "{'lastname': 1, 'age': -1}")`.
Spring Data ensures this index exists on startup (`auto-index-creation=true`).

---

### Question 364: What is the role of `ReactiveMongoRepository`?

**Answer:**
(See Q255). Non-blocking driver usage.

---

### Question 365: How do you work with Cassandra in Spring Boot?

**Answer:**
`spring-boot-starter-data-cassandra`.
`@Table`, `@PrimaryKey`.
Uses `CassandraRepository`.
Config: `spring.cassandra.contact-points`, `keyspace-name`.

---

### Question 366: How do you implement auditing in MongoDB?

**Answer:**
Enable `@EnableMongoAuditing`.
(Same as JPA Q164).

---

### Question 367: How to use Redis as a primary data store with Spring Boot?

**Answer:**
Usually Redis is Cache.
For primary: Use `RedisTemplate` or `RedisRepository` (Spring Data Redis).
Map Objects to Hashes (`HMSET`).
`@RedisHash("users")` on Entity.
Note: Redis is in-memory (Persistence optional), risk of data loss on restart if not configured.

---

### Question 368: How do you handle key expiration in Redis?

**Answer:**
`@RedisHash(value = "users", timeToLive = 600)`. (10 mins).
Or `redisTemplate.expire(key, 10, TimeUnit.MINUTES)`.

---

### Question 369: How do you store objects in Redis using Spring Boot?

**Answer:**
Use `RedisSerializer` (JSON or String).
`redisTemplate.opsForValue().set("key", myObject)`.
Requires Object to be Serializable or use Jackson serializer.

---

### Question 370: How do you implement caching using MongoDB or Redis?

**Answer:**
`@Cacheable("items")`.
If Redis Starter present -> RedisCacheManager.
If Mongo -> No default CacheManager, must implement custom one or use Mongo as DB and Redis as Cache separately.

## 🔹 Section 8: Integration Patterns (371-380)

### Question 371: What is Spring Integration DSL?

**Answer:**
Java DSL to define integration flows.
```java
IntegrationFlows.from("inputChannel")
    .filter("payload > 10")
    .transform("payload * 2")
    .handle(System.out::println)
    .get();
```

---

### Question 372: How do you build a file polling system using Spring Integration?

**Answer:**
`Files.inboundAdapter(new File("/inbox"))`.
Poller with fixed delay.
Messages sent to channel when file appears.

---

### Question 373: How do you integrate with SOAP services using Spring Boot?

**Answer:**
`spring-boot-starter-web-services`.
Use `WebServiceTemplate`.
Marshal/Unmarshal using JAXB.

---

### Question 374: How do you implement FTP file transfers with Spring Integration?

**Answer:**
`Ftp.inboundAdapter(sessionFactory)`.
Connects to FTP server, downloads files to local dir, processes them.

---

### Question 375: How to trigger workflows using Spring Events?

**Answer:**
(See Q350). Publish event -> Listener triggers next step.

---

### Question 376: What is an `IntegrationFlow` in Spring Integration?

**Answer:**
Represents a pipeline of processing steps (Pipes and Filters).

---

### Question 377: What are channels in Spring Integration, and how are they used?

**Answer:**
Pipes connecting components.
- **DirectChannel:** Point-to-point, in-thread.
- **QueueChannel:** Buffered, pollable (can be async).
- **PublishSubscribeChannel:** One-to-many.

---

### Question 378: What is a gateway in Spring Integration?

**Answer:**
Hides the messaging system from application code.
Interface `MyGateway`.
Calling `gateway.process(data)` sends message to channel and waits for reply.
Code looks like normal method call.

---

### Question 379: How do you implement message transformation pipelines?

**Answer:**
`.transform(Transformers.toJson())`.
`.transform(MyBean, "methodName")`.

---

### Question 380: How do you integrate with Apache Camel from Spring Boot?

**Answer:**
Alternative to Spring Integration.
`camel-spring-boot-starter`.
Define `RouteBuilder` bean.
`from("file:inbox").to("jms:queue:orders")`.

---
