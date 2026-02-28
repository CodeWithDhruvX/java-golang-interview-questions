# Solution: Microservices, Kafka, Redis, and Kubernetes

## 1. Microservices Communication (OpenFeign + Resilience4j)

**Solution:**

```java
// 1. OpenFeign Client Interface (in Order-Service)
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

// Note: name acts as the service ID for Eureka discovery
@FeignClient(name = "inventory-service")
public interface InventoryClient {

    @GetMapping("/api/inventory/{skuCode}")
    boolean isInStock(@PathVariable("skuCode") String skuCode);
}

// 2 & 3. Calling Service with Circuit Breaker
import io.github.resilience4j.circuitbreaker.annotation.CircuitBreaker;
import org.springframework.stereotype.Service;

@Service
public class OrderService {

    private final InventoryClient inventoryClient;

    public OrderService(InventoryClient inventoryClient) {
        this.inventoryClient = inventoryClient;
    }

    // Name references the configuration in application.yml
    @CircuitBreaker(name = "inventory", fallbackMethod = "fallbackMethod")
    public String placeOrder(OrderRequest orderRequest) {
        boolean inStock = inventoryClient.isInStock(orderRequest.getSkuCode());

        if (inStock) {
            // Save order to database...
            return "Order Placed Successfully";
        } else {
            throw new IllegalArgumentException("Product is not in stock, please try again later.");
        }
    }

    // Fallback: MUST have the exact same return type and take the same parameters PLUS a Throwable argument.
    public String fallbackMethod(OrderRequest orderRequest, Throwable throwable) {
        return "Oops! Something went wrong, please order after some time! Reason: " + throwable.getMessage();
    }
}
```

---

## 2. Event-Driven Architecture (Apache Kafka)

**Solution:**

```java
// 1. Publisher Configuration (Order-Service)
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

// The DTO payload
@Data
@AllArgsConstructor
public class OrderPlacedEvent {
    private String orderNumber;
    private String customerEmail;
}

@Service
public class OrderProducer {
    
    // Spring Boot Auto-configures this template if spring-kafka is on the classpath
    private final KafkaTemplate<String, OrderPlacedEvent> kafkaTemplate;

    public OrderProducer(KafkaTemplate<String, OrderPlacedEvent> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    public void publishOrderEvent(OrderPlacedEvent event) {
        // Send to topic "order-events". The key can be null.
        kafkaTemplate.send("order-events", event);
        System.out.println("Message published successfully to Kafka");
    }
}

// 2. Consumer Configuration (Notification-Service)
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class NotificationConsumer {

    // Spring Kafka automatically deserializes JSON to the object if configured correctly in application.yml
    @KafkaListener(topics = "order-events", groupId = "notificationGroup")
    public void handleOrderEvent(OrderPlacedEvent event) {
        System.out.println("Processing notification... Sending email for order " 
                           + event.getOrderNumber() + " to " + event.getCustomerEmail());
    }
}
```

---

## 3. High-Performance Caching (Redis)

**Solution:**

```java
import org.springframework.cache.annotation.CacheEvict;
import org.springframework.cache.annotation.CachePut;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.stereotype.Service;

// 1. Enable Caching at the application or config level
@SpringBootApplication
@EnableCaching
public class ProductApplication {
    public static void main(String[] args) {
        SpringApplication.run(ProductApplication.class, args);
    }
}

// 2 & 3. Product Service
@Service
public class ProductService {

    private final ProductRepository productRepository;

    public ProductService(ProductRepository productRepository) {
        this.productRepository = productRepository;
    }

    // 2. On first call, executes DB query and saves returning obj to Redis (products::id).
    // On subsequent calls, fetches directly from Redis without hitting the method body.
    @Cacheable(value = "products", key = "#id")
    public Product getProductById(Long id) {
        System.out.println("Fetching product from Database..."); 
        return productRepository.findById(id).orElseThrow();
    }

    // 3. Updates the DB, and then OVERWRITES the existing item in the Redis cache.
    // Alternatively, use @CacheEvict to delete the key entirely.
    @CachePut(value = "products", key = "#product.id")
    public Product putProduct(Product product) {
        System.out.println("Updating product in Database and Cache...");
        return productRepository.save(product);
    }
    
    // Example of Evict
    @CacheEvict(value = "products", key = "#id")
    public void deleteProduct(Long id) {
        productRepository.deleteById(id);
    }
}
```

---

## 4. Containerization (Docker + Docker Compose)

**Solution:**

```dockerfile
# 1. Dockerfile
FROM eclipse-temurin:17-jre-alpine
VOLUME /tmp
ARG JAR_FILE=target/user-service.jar
COPY ${JAR_FILE} app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/app.jar"]
```

```yaml
# 2. docker-compose.yml
version: '3.8'

services:
  # The Database
  db:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_USER: admin
      MYSQL_PASSWORD: password
      MYSQL_DATABASE: userdb
    ports:
      - "3306:3306"
    volumes:
      - mysql-data:/var/lib/mysql

  # The Spring Boot Application
  user-service:
    build: . # Build from the Dockerfile in the current directory
    ports:
      - "8080:8080"
    environment:
      # Use the literal service name 'db' as the hostname
      SPRING_DATASOURCE_URL: jdbc:mysql://db:3306/userdb
      SPRING_DATASOURCE_USERNAME: admin
      SPRING_DATASOURCE_PASSWORD: password
    depends_on:
      - db

volumes:
  mysql-data:
```

---

## 5. Kubernetes Orchestration (Deployments & Services)

**Solution:**

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service-deployment
  labels:
    app: user-service # High-level label for the deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: user-service # How the Deployment knows which Pods it manages
  template:
    metadata:
      labels:
        app: user-service # The label assigned to the actual Pods
    spec:
      containers:
      - name: user-service
        image: mycompany/user-service:1.0 # Pulls from registry
        ports:
        - containerPort: 8080

---
# service.yaml
apiVersion: v1
kind: Service
metadata:
  name: user-service-cluster-ip # DNS name inside the cluster
spec:
  type: ClusterIP
  selector:
    app: user-service # Matches the Pod template label exactly
  ports:
    - protocol: TCP
      port: 8080        # Port exposed on the Service
      targetPort: 8080  # Port where the container listens
```
