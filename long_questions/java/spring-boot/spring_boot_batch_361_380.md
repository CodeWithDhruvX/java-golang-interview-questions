## 🔹 Section 7: Spring Boot with NoSQL (361-380)

### Question 361: How do you configure MongoDB with Spring Boot?

**Answer:**
`spring-boot-starter-data-mongodb`.
Props: `spring.data.mongodb.uri=mongodb://localhost:27017/db`.
Spring auto-configures `MongoTemplate` and `MongoRepository`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure MongoDB with Spring Boot?
**Your Response:** "I configure MongoDB with Spring Boot by adding the `spring-boot-starter-data-mongodb` dependency. I set the connection URI in properties like `spring.data.mongodb.uri=mongodb://localhost:27017/db`. Spring Boot automatically configures `MongoTemplate` for programmatic access and `MongoRepository` for data access. The auto-configuration handles the connection factory, database initialization, and template setup. I can also configure individual properties like host, port, and database name separately if needed."

---

### Question 362: How does Spring Data MongoDB support query methods?

**Answer:**
Like JPA: `findByUsername(String user)`.
Start with `findBy...`, `deleteBy...`.
Supports JSON-based queries: `@Query("{ 'name': ?0 }")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Data MongoDB support query methods?
**Your Response:** "Spring Data MongoDB supports query methods similar to JPA. I can define methods like `findByUsername(String user)` and Spring automatically generates the query. The methods start with `findBy`, `deleteBy`, or `countBy`. For complex queries, I can use JSON-based queries with `@Query('{ 'name': ?0 }')` where I can write native MongoDB queries. This gives me both the convenience of derived queries and the flexibility of native MongoDB queries when needed."

---

### Question 363: How do you define compound indexes in MongoDB using Spring annotations?

**Answer:**
Annotate Class:
`@CompoundIndex(def = "{'lastname': 1, 'age': -1}")`.
Spring Data ensures this index exists on startup (`auto-index-creation=true`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define compound indexes in MongoDB using Spring annotations?
**Your Response:** "I define compound indexes using the `@CompoundIndex` annotation on my entity class. For example, `@CompoundIndex(def = "{'lastname': 1, 'age': -1}")` creates a compound index on lastname and age. Spring Data ensures this index exists on startup when `auto-index-creation=true`. The index definition uses MongoDB's index syntax, allowing me to specify ascending (1) or descending (-1) order. This automatic index management ensures optimal query performance without manual database administration."

---

### Question 364: What is the role of `ReactiveMongoRepository`?

**Answer:**
(See Q255). Non-blocking driver usage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `ReactiveMongoRepository`?
**Your Response:** "`ReactiveMongoRepository` provides reactive data access for MongoDB using non-blocking drivers. Instead of returning entities directly, methods return `Mono<T>` for single results or `Flux<T>` for multiple results. This allows me to build fully reactive applications that can handle high concurrency with fewer threads. I use it when I need to integrate MongoDB into a reactive stack with WebFlux, ensuring the entire data pipeline remains non-blocking from the web layer to the database layer."

---

### Question 365: How do you work with Cassandra in Spring Boot?

**Answer:**
`spring-boot-starter-data-cassandra`.
`@Table`, `@PrimaryKey`.
Uses `CassandraRepository`.
Config: `spring.cassandra.contact-points`, `keyspace-name`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you work with Cassandra in Spring Boot?
**Your Response:** "I work with Cassandra by adding the `spring-boot-starter-data-cassandra` dependency. I annotate entities with `@Table` and `@PrimaryKey` to map them to Cassandra tables. I use `CassandraRepository` for data access operations. I configure the connection using `spring.cassandra.contact-points` for the cluster nodes and `keyspace-name` for the keyspace. Spring Boot auto-configures the session and mapping context, making Cassandra integration straightforward while still allowing full control over the database configuration."

---

### Question 366: How do you implement auditing in MongoDB?

**Answer:**
Enable `@EnableMongoAuditing`.
(Same as JPA Q164).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement auditing in MongoDB?
**Your Response:** "I implement auditing in MongoDB by enabling `@EnableMongoAuditing` on my configuration class. This automatically tracks created and modified dates for my entities. I annotate fields with `@CreatedDate`, `@LastModifiedDate`, `@CreatedBy`, and `@LastModifiedBy`. Spring Data MongoDB automatically populates these fields when entities are created or updated. This provides audit trails without manual coding, which is essential for compliance and debugging. The auditing works similarly to JPA auditing but is specifically optimized for MongoDB's document structure."

---

### Question 367: How to use Redis as a primary data store with Spring Boot?

**Answer:**
Usually Redis is Cache.
For primary: Use `RedisTemplate` or `RedisRepository` (Spring Data Redis).
Map Objects to Hashes (`HMSET`).
`@RedisHash("users")` on Entity.
Note: Redis is in-memory (Persistence optional), risk of data loss on restart if not configured.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to use Redis as a primary data store with Spring Boot?
**Your Response:** "While Redis is typically used as a cache, I can use it as a primary data store with `RedisTemplate` or `RedisRepository`. I map objects to Redis hashes using `@RedisHash('users')` on entities. However, I need to be aware that Redis is primarily in-memory with optional persistence, so there's a risk of data loss on restart if not properly configured. I use this approach for use cases where performance is critical and occasional data loss is acceptable, like session management or temporary data storage."

---

### Question 368: How do you handle key expiration in Redis?

**Answer:**
`@RedisHash(value = "users", timeToLive = 600)`. (10 mins).
Or `redisTemplate.expire(key, 10, TimeUnit.MINUTES)`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle key expiration in Redis?
**Your Response:** "I handle key expiration in Redis using the `timeToLive` attribute in `@RedisHash(value = 'users', timeToLive = 600)` which sets a 10-minute TTL on all keys. For more granular control, I use `redisTemplate.expire(key, 10, TimeUnit.MINUTES)` to set expiration programmatically. This automatic expiration is perfect for session data, temporary caches, or any data that should automatically disappear after a certain time. Redis efficiently manages the expiration, removing keys automatically when they expire."

---

### Question 369: How do you store objects in Redis using Spring Boot?

**Answer:**
Use `RedisSerializer` (JSON or String).
`redisTemplate.opsForValue().set("key", myObject)`.
Requires Object to be Serializable or use Jackson serializer.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you store objects in Redis using Spring Boot?
**Your Response:** "I store objects in Redis using `RedisSerializer` for JSON or String serialization. I use `redisTemplate.opsForValue().set('key', myObject)` to store objects. The object needs to be Serializable, or I configure a Jackson serializer to automatically convert objects to JSON. Spring Boot provides default serializers, but I can customize them for specific serialization needs. This approach allows me to store complex Java objects in Redis while maintaining type safety and automatic deserialization."

---

### Question 370: How do you implement caching using MongoDB or Redis?

**Answer:**
`@Cacheable("items")`.
If Redis Starter present -> RedisCacheManager.
If Mongo -> No default CacheManager, must implement custom one or use Mongo as DB and Redis as Cache separately.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement caching using MongoDB or Redis?
**Your Response:** "I implement caching using `@Cacheable('items')` annotations. If the Redis starter is present, Spring Boot automatically configures a `RedisCacheManager`. For MongoDB, there's no default CacheManager, so I'd need to implement a custom one or use MongoDB as the primary database while using Redis specifically for caching. I prefer using Redis for caching due to its in-memory performance and built-in expiration features, while using MongoDB for persistent data storage. This separation gives me the best of both worlds."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Integration DSL?
**Your Response:** "Spring Integration DSL provides a Java-based Domain Specific Language for defining integration flows. Instead of XML configuration, I can write fluent Java code like `IntegrationFlows.from('inputChannel').filter('payload > 10').transform('payload * 2').handle(System.out::println).get()` to build message processing pipelines. This DSL makes integration patterns more readable and type-safe while maintaining the powerful Enterprise Integration Patterns. It's particularly useful for complex routing, transformation, and integration scenarios where I need programmatic control."

---

### Question 372: How do you build a file polling system using Spring Integration?

**Answer:**
`Files.inboundAdapter(new File("/inbox"))`.
Poller with fixed delay.
Messages sent to channel when file appears.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a file polling system using Spring Integration?
**Your Response:** "I build a file polling system using Spring Integration's `Files.inboundAdapter(new File('/inbox'))`. I configure a poller with a fixed delay to check the directory periodically. When files appear, messages are sent to a channel for processing. This approach is perfect for batch processing systems where I need to monitor directories for incoming files. The adapter handles file watching, message creation, and error handling, allowing me to focus on the business logic of processing the files."

---

### Question 373: How do you integrate with SOAP services using Spring Boot?

**Answer:**
`spring-boot-starter-web-services`.
Use `WebServiceTemplate`.
Marshal/Unmarshal using JAXB.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you integrate with SOAP services using Spring Boot?
**Your Response:** "I integrate with SOAP services using the `spring-boot-starter-web-services` dependency. I use `WebServiceTemplate` for making SOAP calls and marshal/unmarshal XML using JAXB. Spring Boot simplifies SOAP integration by auto-configuring the template and providing support for WSDL generation. I can create client-side proxies or server-side endpoints with minimal configuration. This approach is essential when working with legacy systems that use SOAP protocols, providing a clean integration with modern Spring Boot applications."

---

### Question 374: How do you implement FTP file transfers with Spring Integration?

**Answer:**
`Ftp.inboundAdapter(sessionFactory)`.
Connects to FTP server, downloads files to local dir, processes them.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement FTP file transfers with Spring Integration?
**Your Response:** "I implement FTP file transfers using Spring Integration's `Ftp.inboundAdapter(sessionFactory)`. The adapter connects to an FTP server, downloads files to a local directory, and processes them. I configure the session factory with connection details and set up polling to check for new files periodically. This approach is perfect for integrating with external systems that use FTP for data exchange. Spring Integration handles the FTP connection, file transfer, and error handling, making the integration robust and maintainable."

---

### Question 375: How to trigger workflows using Spring Events?

**Answer:**
(See Q350). Publish event -> Listener triggers next step.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to trigger workflows using Spring Events?
**Your Response:** "I trigger workflows using Spring's event system. I publish events from one component and have listeners that trigger the next step in the workflow. For example, after processing an order, I publish an `OrderProcessedEvent`, and different listeners handle inventory updates, notification sending, and billing. This decoupled approach allows me to add new workflow steps without modifying existing code. Spring's event system provides a clean way to implement event-driven architectures within a single application."

---

### Question 376: What is an `IntegrationFlow` in Spring Integration?

**Answer:**
Represents a pipeline of processing steps (Pipes and Filters).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is an `IntegrationFlow` in Spring Integration?
**Your Response:** "An `IntegrationFlow` represents a pipeline of processing steps following the Pipes and Filters pattern. Each component in the flow performs a specific operation like filtering, transforming, or routing messages. I define flows using the DSL to create readable, maintainable integration logic. The flow represents the entire message processing pipeline from input to output, handling all the intermediate transformations and routing. This abstraction makes complex integration scenarios manageable and testable."

---

### Question 377: What are channels in Spring Integration, and how are they used?

**Answer:**
Pipes connecting components.
- **DirectChannel:** Point-to-point, in-thread.
- **QueueChannel:** Buffered, pollable (can be async).
- **PublishSubscribeChannel:** One-to-many.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are channels in Spring Integration, and how are they used?
**Your Response:** "Channels in Spring Integration are pipes that connect components. I use different channel types based on needs: `DirectChannel` for point-to-point, in-thread communication; `QueueChannel` for buffered, pollable communication that can be asynchronous; and `PublishSubscribeChannel` for one-to-many broadcasting. Channels decouple message producers from consumers, allowing me to change the processing strategy without affecting the components. This flexibility is key to building robust integration systems."

---

### Question 378: What is a gateway in Spring Integration?

**Answer:**
Hides the messaging system from application code.
Interface `MyGateway`.
Calling `gateway.process(data)` sends message to channel and waits for reply.
Code looks like normal method call.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a gateway in Spring Integration?
**Your Response:** "A gateway hides the messaging system from application code. I define an interface like `MyGateway` and Spring Integration creates a proxy. When I call `gateway.process(data)`, it sends a message to a channel and waits for a reply. The code looks like a normal method call, but behind the scenes, it's using messaging. This abstraction allows me to integrate messaging without exposing the complexity to the business logic, making the code cleaner and easier to test."

---

### Question 379: How do you implement message transformation pipelines?

**Answer:**
`.transform(Transformers.toJson())`.
`.transform(MyBean, "methodName")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement message transformation pipelines?
**Your Response:** "I implement message transformation using the `.transform()` method in Spring Integration. I can use built-in transformers like `Transformers.toJson()` for JSON conversion, or call methods on beans with `.transform(MyBean, 'methodName')`. These transformations can be chained together to create complex processing pipelines. The transformation step receives a message, processes it, and passes it to the next step. This approach is perfect for data enrichment, format conversion, or any message processing that needs to happen between systems."

---

### Question 380: How do you integrate with Apache Camel from Spring Boot?

**Answer:**
Alternative to Spring Integration.
`camel-spring-boot-starter`.
Define `RouteBuilder` bean.
`from("file:inbox").to("jms:queue:orders")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you integrate with Apache Camel from Spring Boot?
**Your Response:** "I integrate with Apache Camel as an alternative to Spring Integration using the `camel-spring-boot-starter`. I define routes using `RouteBuilder` beans with DSL like `from('file:inbox').to('jms:queue:orders')`. Camel provides extensive connectivity options and enterprise integration patterns. I choose Camel when I need its extensive component library or when I'm working in an environment that already uses Camel. The Spring Boot starter makes Camel integration seamless while maintaining Spring's dependency injection and configuration."

---
