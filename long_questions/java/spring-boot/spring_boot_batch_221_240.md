## 🔹 Section 7: Actuator, Metrics & Monitoring (221-240)

### Question 221: What is the `/actuator/health` endpoint and how to customize it?

**Answer:**
Returns application health status (`UP`, `DOWN`).
Customize by creating a `HealthIndicator` bean.
```java
@Component
public class MyHealth implements HealthIndicator {
    public Health health() {
        return Health.up().withDetail("service", "ok").build();
    }
}
```

---

### Question 222: How do you secure actuator endpoints?

**Answer:**
(See Q188). Use Spring Security `requestMatchers`.

---

### Question 223: How do you integrate Prometheus and Grafana with Spring Boot?

**Answer:**
1.  Add `micrometer-registry-prometheus`.
2.  Enable endpoint: `management.endpoints.web.exposure.include=prometheus`.
3.  Configure Prometheus server to scrape `localhost:8080/actuator/prometheus`.
4.  Point Grafana data source to Prometheus.

---

### Question 224: What is Micrometer in Spring Boot?

**Answer:**
**SLF4J for Metrics.**
A facade over monitoring systems.
You instrument code once (`Timer.sample()`), and Micrometer exports it to Prometheus, Datadog, New Relic, etc., based on the classpath dependency.

---

### Question 225: How do you define custom metrics using Micrometer?

**Answer:**
Inject `MeterRegistry`.
```java
Counter.builder("my.custom.counter")
    .tag("type", "order")
    .register(registry)
    .increment();
```

---

### Question 226: How to add tags and dimensions to metrics?

**Answer:**
Tags allow filtering in Grafana (e.g., `region=us-east`).
Pass them in the builder:
`.tags("region", "us-east", "status", "failed")`.

---

### Question 227: How to monitor application memory and CPU usage using Actuator?

**Answer:**
Automatically exposed in `/actuator/metrics`.
JVM metrics (`jvm.memory.used`, `system.cpu.usage`) are bound by default if Micrometer is present.

---

### Question 228: What is the difference between `@Timed`, `@Metered`, and `@Gauge` annotations?

**Answer:**
*   **`@Timed` (Micrometer):** Measures execution time and count of a method.
*   **`@Metered` / `@Gauge`:** Legacy (Dropwizard metrics). Gauge measures a value at a point in time (e.g., List size).

---

### Question 229: How to enable distributed tracing using Spring Cloud Sleuth?

**Answer:**
(Renamed to **Micrometer Tracing** in Boot 3).
Dependency: `micrometer-tracing-bridge-brave`.
Automatically adds `TraceID` and `SpanID` to logs using MDC.
Propagates headers (`b3`, `traceparent`) to downstream HTTP/Messaging calls.

---

### Question 230: How to monitor thread pool metrics in Spring Boot?

**Answer:**
Exposed via `executor` metrics if you use Spring's `ThreadPoolTaskExecutor`.
Metrics: `executor.active`, `executor.queued`, `executor.pool.size`.

---

### Question 231: How do you add custom health indicators to Spring Boot Actuator?

**Answer:**
(See Q221).

---

### Question 232: What is the role of `InfoContributor` in Spring Boot Actuator?

**Answer:**
Populates `/actuator/info`.
Implement `InfoContributor`.
Often acts as a place to expose Build Version (`git.properties`), Build Time, or OS details.

---

### Question 233: How do you expose actuator endpoints over JMX?

**Answer:**
Enabled by default.
Property `spring.jmx.enabled=true`.
Use **JConsole** or **VisualVM** to connect to the running process and invoke managed operations (like shutdown or logger level change).

---

### Question 234: What is the purpose of the `metrics` endpoint?

**Answer:**
`/actuator/metrics` lists all available metric names.
`/actuator/metrics/jvm.memory.used` shows the value of that specific metric.

---

### Question 235: How do you publish custom application metrics?

**Answer:**
(See 225).

---

### Question 236: How do you log long GC pauses with Spring Boot?

**Answer:**
Micrometer binds `jvm.gc.pause` metrics.
To LOG them: Enable GC Logging flags in JVM (`-Xlog:gc*`) or register a `JvmGcMetrics` listener that logs duration > threshold.

---

### Question 237: What is the difference between `health`, `info`, and `metrics` endpoints?

**Answer:**
- **Health:** Status (Up/Down) for Liveness/Readiness probes.
- **Info:** Static text (Version, Description).
- **Metrics:** Numeric time-series data (CPU, Request Count).

---

### Question 238: How to use actuator with Spring Cloud Config?

**Answer:**
`/actuator/refresh` endpoint allows reloading `@ConfigurationProperties` from the Config Server without restarting the app.

---

### Question 239: How to restrict access to sensitive actuator endpoints?

**Answer:**
(See Q188). Firewall/Security Config is mandatory.

---

### Question 240: What metrics are collected by default in Spring Boot?

**Answer:**
JVM (Memory, GC, Threads).
Tomcat (Sessions, Requests).
System (CPU, Uptime).
Logback (Log events count).
HttpClient/DataSource usage.

---
