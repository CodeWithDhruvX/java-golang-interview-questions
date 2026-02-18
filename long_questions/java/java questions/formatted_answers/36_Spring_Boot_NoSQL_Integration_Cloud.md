# 36. Spring Boot (NoSQL, Integration & Cloud)

**Q: Redis with Spring Boot**
> "Redis is typically used for caching, but you can use it as a primary store.
> You use `RedisTemplate`.
>
> It’s a Key-Value store. So you treat it like a giant, persistent `HashMap` in the cloud.
> `template.opsForValue().set("user:1", jsonString);`
> It's incredibly fast (sub-millisecond) but data must fit in memory."

**Indepth:**
> **Serializers**: `JdkSerializationRedisSerializer` is default but bad (binary blobs). Use `Jackson2JsonRedisSerializer` so data is readable JSON in Redis CLI. `StringRedisTemplate` is a pre-configured template just for String keys/values.


---

**Q: Redis Key Expiration**
> "One of the best features of Redis is that data can self-destruct.
> When you save a key, you set a TTL (Time To Live).
>
> `template.opsForValue().set("otp:12345", "8732", Duration.ofMinutes(5));`
>
> After 5 minutes, Redis automatically deletes it. This is perfect for OTPs, User Sessions, and temporary Cache entries."

**Indepth:**
> **Eviction**: What happens when Redis is full? It deletes keys. You configure the eviction policy. `allkeys-lru` deletes any key. `volatile-lru` only deletes keys with an expiry set.


---

**Q: Spring Integration DSL**
> "Spring Integration is about connecting systems (Files, FTP, Queues).
> The DSL (Domain Specific Language) allows you to define these 'Pipelines' in Java code so it reads like a story:
>
> ```java
> return IntegrationFlow.from("inputChannel")
>     .filter("payload.amount > 100")
>     .transform(Transformers.toJson())
>     .handle(Amqp.outboundAdapter(rabbitTemplate))
>     .get();
> ```
> It says: Take input -> Filter high amounts -> Convert to JSON -> Send to RabbitMQ."

**Indepth:**
> **Channels**: Channels are the pipes. `DirectChannel` is a method call (synchronous, same thread). `QueueChannel` is a buffer (asynchronous, different thread). The DSL hides this complexity.


---

**Q: File Polling (Spring Integration)**
> "If you need to watch a folder for new PDF files and process them.
> You define an `InboundFileAdapter`.
>
> It polls the directory every 5 seconds. If it finds a new file, it locks it, passes it to your processing method, and then moves it to a 'processed' folder automatically. No manual `While(true)` loops needed."

**Indepth:**
> **Idempotency**: `AcceptOnceFileListFilter`. How do you prevent processing the same file twice? You need a filter. Be careful: standard filters keep state in memory. If you restart the app, it might process old files again unless you use a persistent usage store (MetadataStore).


---

**Q: Spring Cloud Sleuth & Zipkin**
> "**Sleuth** adds a unique ID (Trace ID) to your logs.
> When Service A calls Service B, Sleuth passes that ID in the HTTP Headers.
>
> **Zipkin** is the UI. It collects all these logs and draws a timeline: 'Request blocked for 200ms in Service A, then took 50ms in Service B'. It's essential for debugging microservices latency."

**Indepth:**
> **Sampling**: `probability`. You don't want to trace 100% of requests in production (performance overhead). You set `spring.sleuth.sampler.probability=0.1` (10%). But for errors, you always want the trace.


---

**Q: Service Discovery (Eureka)**
> "In the cloud, IP addresses change all the time. You can't hardcode `http://192.168.1.50`.
>
> **Eureka** is a phonebook.
> 1.  Service A starts up and says: 'I am Service A, my IP is X'. (Registration)
> 2.  Service B asks Eureka: 'Where is Service A?'.
> 3.  Eureka replies: 'It's at IP X'.
> Service B then calls Service A directly."

**Indepth:**
> **Self Preservation**: If Eureka stops receiving heartbeats from *many* instances at once (e.g., network partition), it stops expiring them. It assumes the network is down, not the instances. This prevents mass accidental shutdowns.


---

**Q: Feign Client**
> "Stop using `RestTemplate` for calling other microservices. It's verbose.
>
> **Feign** is declarative. You just write an interface:
> ```java
> @FeignClient(name = "inventory-service")
> public interface Inventory {
>     @GetMapping("/items/{id}")
>     Item getItem(@PathVariable String id);
> }
> ```
> Spring generates the implementation at runtime. If you call `getItem("5")`, it automatically calls `http://inventory-service/items/5` via Eureka."

**Indepth:**
> **Error Decoding**: Feign throws `FeignException` by default. You implement a custom `ErrorDecoder` to translate 404s/500s from the remote service into your own domain exceptions (`InventoryNotFoundException`).


---

**Q: Circuit Breaker (Resilience4j)**
> "If the 'Inventory Service' goes down, you don't want the 'Order Service' to hang and crash too (Cascading Failure).
>
> A **Circuit Breaker** wraps the call.
> *   **Closed**: Normal operation.
> *   **Open**: Too many failures (result > 50%). Stick fails immediately (Fast Fail) without waiting for timeout.
> *   **Half-Open**: Let one request through to see if it's fixed.
>
> It keeps your system responsive even when dependencies fail."

**Indepth:**
> **Bulkhead**: A Circuit Breaker stops all calls when the failure rate is high. A **Bulkhead** limits concurrency. "Max 10 concurrent calls to Inventory Service". If the 11th comes, it's rejected immediately. This prevents one slow service from exhausting all your Tomcat threads.


---

**Q: Buildpacks (Cloud Native)**
> "You don't need a Dockerfile anymore.
> `mvn spring-boot:build-image`
>
> This uses **Cloud Native Buildpacks**. It detects 'Oh, this is a Java 17 app'. It automatically downloads the best JRE image, optimizes memory settings, layers the JAR, and gives you a production-ready Docker image. It's magic."

**Indepth:**
> **Rebase**: The coolest feature. If there is a security patch in the underlying OS (Ubuntu SSL), you don't need to rebuild your app. You just "rebase" the image layers. The app layer stays the same; the OS layer is swapped underneath it instantly.


---

**Q: Blue-Green Deployment**
> "You have Version 1 running (Blue).
> You deploy Version 2 (Green) alongside it.
> You run tests on Green.
>
> Then, you switch the Load Balancer router: 100% traffic goes to Green.
> If Green crashes, you instantly switch back to Blue.
> Spring Boot doesn't do this itself, but it provides the **Metrics** and **Health Probes** that tools like Kubernetes or AWS use to orchestrate this switch safely."

**Indepth:**
> **Database**: The DB is shared. Version 2 app cannot rename a column that Version 1 app is still using. You must perform "Expand and Contract" migrations (Add new column, copy data, unrelated changes) to ensure backward compatibility.

