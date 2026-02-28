# Practical Coding Questions: Microservices, Kafka, Redis, and Kubernetes

*These problems are designed to test your actual coding ability during a technical interview. Try to write out the code or pseudocode before looking at the suggested architectures.*

---

## 1. Microservices Communication (OpenFeign + Resilience4j)
**Problem Statement:**
You have an `Order-Service` and an `Inventory-Service`. The `Order-Service` must call the `Inventory-Service` synchronously to check if an item is in stock before placing an order.

1. Write a Spring Cloud `OpenFeign` client interface in the `Order-Service` that calls the `GET /api/inventory/{skuCode}` endpoint.
2. The `Inventory-Service` is sometimes slow or goes down. Implement a **Circuit Breaker** using Resilience4j on the feign client call.
3. Write a fallback method that physically executes if the circuit is Open or the call fails, returning a default `false` (item out of stock).

**Expected Focus Areas:**
- `@FeignClient(name = "inventory-service")`.
- Method mapping: `@GetMapping("/api/inventory/{skuCode}") boolean isInStock(@PathVariable("skuCode") String skuCode);`.
- Decorating the calling service method with `@CircuitBreaker(name = "inventory", fallbackMethod = "fallbackMethod")`.
- Implementing `public boolean fallbackMethod(String skuCode, Throwable throwable) { return false; }`.

---

## 2. Event-Driven Architecture (Apache Kafka)
**Problem Statement:**
To decouple your services, you decide that when an Order is placed, the `Order-Service` should NOT synchronously call the `Notification-Service`. Instead, it should publish an event to Kafka.

1. **Publisher (`Order-Service`):** Write a service class using `KafkaTemplate` to publish an `OrderPlacedEvent` object (serialized as JSON) to a Kafka topic named `order-events`.
2. **Consumer (`Notification-Service`):** Write a service class with a method annotated with `@KafkaListener` that listens to the `order-events` topic, deserializes the JSON back into an object, and prints "Sending email for order XYZ".

**Expected Focus Areas:**
- **Publisher:** `@Autowired KafkaTemplate<String, OrderPlacedEvent>`, calling `kafkaTemplate.send("order-events", event)`.
- **Consumer:** `@KafkaListener(topics = "order-events", groupId = "notificationGroup")`.
- Method signature matching the payload `public void handleOrderEvent(OrderPlacedEvent event)`.

---

## 3. High-Performance Caching (Redis)
**Problem Statement:**
You have a `Product-Service` with an endpoint `GET /api/products/{id}`. The database query takes 500ms, which is too slow for your high-traffic frontend.

1. Enable caching in your Spring Boot configuration.
2. Annotate the `ProductService.getProductById(Long id)` method to cache results in Redis.
3. When a product is updated via `putProduct(Product product)`, the cache must be explicitly updated or invalidated so the frontend doesn't see stale data. Write the code for this update method.

**Expected Focus Areas:**
- Class-level configuration: `@EnableCaching`.
- Fetch method: `@Cacheable(value = "products", key = "#id")`.
- Update method: `@CachePut(value = "products", key = "#product.id")` (updates cache) OR `@CacheEvict(value = "products", key = "#product.id")` (deletes cache entry, forcing a DB read on the next GET).

---

## 4. Containerization (Docker + Docker Compose)
**Problem Statement:**
You have a compiled Spring Boot application JAR: `target/user-service.jar`. It requires a PostgreSQL database to run.

1. Write a `Dockerfile` to containerize the Spring Boot application using Java 17.
2. Write a `docker-compose.yml` file that stands up:
   - A `postgres` container (setting the `POSTGRES_USER`, `POSTGRES_PASSWORD`, and `POSTGRES_DB` environment variables).
   - Your Spring Boot `user-service` container, built from the current directory, exposing port `8080`, and relying on the network to connect to the postgres container.

**Expected Focus Areas:**
- **Dockerfile:** `FROM openjdk:17`, `COPY target/user-service.jar app.jar`, `ENTRYPOINT ["java", "-jar", "/app.jar"]`.
- **Compose:** Using `depends_on: - db` to specify startup order. Passing environment variables like `SPRING_DATASOURCE_URL=jdbc:postgresql://db:5432/mydb` into the Java container so it resolves the DB container's hostname.

---

## 5. Kubernetes Orchestration (Deployments & Services)
**Problem Statement:**
Your `user-service` Docker image is now pushed to a registry: `mycompany/user-service:1.0`. You need to deploy it to a Kubernetes cluster for scale and high availability.

1. Write the YAML for a Kubernetes **Deployment** that runs exactly 3 replicas of the `user-service` container.
2. Write the YAML for a Kubernetes **Service** (type `ClusterIP`) that exposes these 3 pods on internal port `8080` so other microservices can communicate with it via a stable DNS name.

**Expected Focus Areas:**
- **Deployment:** `apiVersion: apps/v1`, `kind: Deployment`, `spec.replicas: 3`, using a `spec.selector.matchLabels` (e.g., `app: user-service`) and assigning that label to the `template.metadata.labels`. Defining the container `image`.
- **Service:** `apiVersion: v1`, `kind: Service`, `targetPort: 8080`, using a `spec.selector` that exactly matches the label (`app: user-service`) defined in the Deployment.
