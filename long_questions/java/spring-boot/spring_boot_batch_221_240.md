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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the `/actuator/health` endpoint and how to customize it?
**Your Response:** "The `/actuator/health` endpoint returns the application's health status as `UP` or `DOWN`. I customize it by creating beans that implement the `HealthIndicator` interface. Each indicator contributes to the overall health status. For example, I can create a database health indicator that checks database connectivity, or an external service indicator that checks API availability. I implement the `health()` method to return `Health.up()` or `Health.down()` with additional details. Spring Boot aggregates all indicators to determine the overall application health, giving me comprehensive monitoring of all critical dependencies."

---

### Question 222: How do you secure actuator endpoints?

**Answer:**
(See Q188). Use Spring Security `requestMatchers`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure actuator endpoints?
**Your Response:** "I secure actuator endpoints using Spring Security's `SecurityFilterChain` configuration with `requestMatchers`. I use `http.requestMatchers(EndpointRequest.toAnyEndpoint()).hasRole('ADMIN')` to restrict all actuator endpoints to administrative users. I can also secure specific endpoints individually using `EndpointRequest.to(HealthEndpoint.class)` for fine-grained control. Since actuator endpoints expose sensitive application information and management capabilities, I always secure them in production, typically restricting access to only authorized administrators or monitoring systems."

---

### Question 223: How do you integrate Prometheus and Grafana with Spring Boot?

**Answer:**
1.  Add `micrometer-registry-prometheus`.
2.  Enable endpoint: `management.endpoints.web.exposure.include=prometheus`.
3.  Configure Prometheus server to scrape `localhost:8080/actuator/prometheus`.
4.  Point Grafana data source to Prometheus.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you integrate Prometheus and Grafana with Spring Boot?
**Your Response:** "I integrate Prometheus and Grafana by first adding the `micrometer-registry-prometheus` dependency to expose metrics in Prometheus format. I enable the Prometheus endpoint with `management.endpoints.web.exposure.include=prometheus`. Then I configure Prometheus to scrape metrics from `localhost:8080/actuator/prometheus`. Finally, I point Grafana to Prometheus as a data source and create dashboards to visualize the metrics. This setup gives me comprehensive monitoring with Prometheus collecting metrics and Grafana providing beautiful visualizations and alerts."

---

### Question 224: What is Micrometer in Spring Boot?

**Answer:**
**SLF4J for Metrics.**
A facade over monitoring systems.
You instrument code once (`Timer.sample()`), and Micrometer exports it to Prometheus, Datadog, New Relic, etc., based on the classpath dependency.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Micrometer in Spring Boot?
**Your Response:** "Micrometer is like SLF4J but for metrics - it's a facade over various monitoring systems. I instrument my code once using Micrometer's `Timer.sample()` or `Counter.increment()`, and Micrometer automatically exports the metrics to whatever monitoring system I have on the classpath - Prometheus, Datadog, New Relic, etc. This vendor-neutral approach means I can switch monitoring systems without changing my instrumentation code. Micrometer handles the translation to each system's specific format, making my application monitoring portable and flexible."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define custom metrics using Micrometer?
**Your Response:** "I define custom metrics by injecting the `MeterRegistry` and using its builder methods. For counters, I use `Counter.builder('my.custom.counter').tag('type', 'order').register(registry).increment()` to create and increment a counter with tags. For timers, I use `Timer.builder('api.response.time').register(registry)` to measure execution time. I can create gauges, distribution summaries, and other metric types. These custom metrics appear alongside the built-in metrics, giving me complete visibility into my application's specific business metrics and performance characteristics."

---

### Question 226: How to add tags and dimensions to metrics?

**Answer:**
Tags allow filtering in Grafana (e.g., `region=us-east`).
Pass them in the builder:
`.tags("region", "us-east", "status", "failed")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to add tags and dimensions to metrics?
**Your Response:** "I add tags to metrics using the builder's `.tags()` method. Tags allow me to filter and group metrics in monitoring tools like Grafana. For example, I can add tags like `.tags('region', 'us-east', 'status', 'failed')` to categorize metrics by geographical region and operation status. Tags are key-value pairs that become dimensions in my monitoring system, enabling powerful filtering and aggregation. I can then create Grafana dashboards that show metrics by region, compare success rates across regions, or track error rates by status. This dimensional approach makes monitoring much more insightful."

---

### Question 227: How to monitor application memory and CPU usage using Actuator?

**Answer:**
Automatically exposed in `/actuator/metrics`.
JVM metrics (`jvm.memory.used`, `system.cpu.usage`) are bound by default if Micrometer is present.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to monitor application memory and CPU usage using Actuator?
**Your Response:** "Spring Boot automatically exposes JVM and system metrics through the `/actuator/metrics` endpoint when Micrometer is present. I get metrics like `jvm.memory.used` for heap usage, `jvm.gc.pause` for garbage collection, and `system.cpu.usage` for CPU utilization. These metrics are collected automatically without any additional code. I can access them via the metrics endpoint or configure monitoring systems to scrape them. This built-in monitoring gives me instant visibility into my application's resource usage and performance characteristics without requiring any custom instrumentation."

---

### Question 228: What is the difference between `@Timed`, `@Metered`, and `@Gauge` annotations?

**Answer:**
*   **`@Timed` (Micrometer):** Measures execution time and count of a method.
*   **`@Metered` / `@Gauge`:** Legacy (Dropwizard metrics). Gauge measures a value at a point in time (e.g., List size).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@Timed`, `@Metered`, and `@Gauge` annotations?
**Your Response:** "`@Timed` is the modern Micrometer annotation that measures method execution time and call counts. `@Metered` and `@Gauge` are legacy annotations from Dropwizard metrics. A gauge measures a value at a specific point in time, like the size of a collection or current queue length. I prefer `@Timed` for measuring method performance because it's part of the modern Micrometer ecosystem. The gauge approach is useful when I want to monitor the current state of something rather than measuring execution time or counting occurrences."

---

### Question 229: How to enable distributed tracing using Spring Cloud Sleuth?

**Answer:**
(Renamed to **Micrometer Tracing** in Boot 3).
Dependency: `micrometer-tracing-bridge-brave`.
Automatically adds `TraceID` and `SpanID` to logs using MDC.
Propagates headers (`b3`, `traceparent`) to downstream HTTP/Messaging calls.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to enable distributed tracing using Spring Cloud Sleuth?
**Your Response:** "In Spring Boot 3, distributed tracing is now called Micrometer Tracing. I add the `micrometer-tracing-bridge-brave` dependency to enable tracing. It automatically adds TraceID and SpanID to my logs using MDC, making it easy to trace requests across multiple services. It also propagates tracing headers like `b3` or `traceparent` to downstream HTTP calls and messaging, ensuring the trace continues across service boundaries. This gives me end-to-end visibility into request flows through my microservices architecture without requiring manual instrumentation."

---

### Question 230: How to monitor thread pool metrics in Spring Boot?

**Answer:**
Exposed via `executor` metrics if you use Spring's `ThreadPoolTaskExecutor`.
Metrics: `executor.active`, `executor.queued`, `executor.pool.size`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to monitor thread pool metrics in Spring Boot?
**Your Response:** "Spring Boot automatically exposes thread pool metrics when I use Spring's `ThreadPoolTaskExecutor`. I get metrics like `executor.active` for currently active threads, `executor.queued` for queued tasks, and `executor.pool.size` for the pool size. These metrics help me monitor the health and performance of my async operations. I can track whether my thread pools are saturated, if tasks are backing up, or if I have too many idle threads. This visibility is crucial for tuning thread pool sizes and identifying performance bottlenecks in asynchronous processing."

---

### Question 231: How do you add custom health indicators to Spring Boot Actuator?

**Answer:**
(See Q221).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you add custom health indicators to Spring Boot Actuator?
**Your Response:** "I add custom health indicators by creating beans that implement the `HealthIndicator` interface. I implement the `health()` method to check the health of specific components like databases, external APIs, or other dependencies. Each indicator contributes to the overall application health. For example, I can create a `DatabaseHealthIndicator` that checks database connectivity, or an `ExternalServiceHealthIndicator` that pings critical APIs. Spring Boot automatically aggregates all health indicators, giving me a comprehensive view of my application's health status through the `/actuator/health` endpoint."

---

### Question 232: What is the role of `InfoContributor` in Spring Boot Actuator?

**Answer:**
Populates `/actuator/info`.
Implement `InfoContributor`.
Often acts as a place to expose Build Version (`git.properties`), Build Time, or OS details.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `InfoContributor` in Spring Boot Actuator?
**Your Response:** "`InfoContributor` is used to populate the `/actuator/info` endpoint with application information. I implement this interface to expose details like build version from `git.properties`, build time, or OS information. This endpoint is perfect for providing deployment information, version tracking, or environment details that help with operations and debugging. Unlike health indicators which show runtime status, info contributors provide static or semi-static information about the application itself. This makes it easy to identify which version is deployed and when it was built."

---

### Question 233: How do you expose actuator endpoints over JMX?

**Answer:**
Enabled by default.
Property `spring.jmx.enabled=true`.
Use **JConsole** or **VisualVM** to connect to the running process and invoke managed operations (like shutdown or logger level change).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you expose actuator endpoints over JMX?
**Your Response:** "Spring Boot enables JMX exposure of actuator endpoints by default when `spring.jmx.enabled=true`. I can then use tools like JConsole or VisualVM to connect to the running application process and invoke managed operations. Through JMX, I can perform operations like shutting down the application, changing logger levels, or triggering health checks without using HTTP endpoints. This is particularly useful in enterprise environments where JMX is the standard management protocol, or when I need programmatic access to management operations."

---

### Question 234: What is the purpose of the `metrics` endpoint?

**Answer:**
`/actuator/metrics` lists all available metric names.
`/actuator/metrics/jvm.memory.used` shows the value of that specific metric.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of the `metrics` endpoint?
**Your Response:** "The `/actuator/metrics` endpoint serves two purposes. First, `/actuator/metrics` lists all available metric names, giving me an overview of what's being monitored. Second, `/actuator/metrics/{metric.name}` like `/actuator/metrics/jvm.memory.used` shows the detailed value and metadata for a specific metric. This endpoint is the foundation for monitoring - it exposes all the raw metrics that monitoring systems like Prometheus can scrape. I can explore available metrics and check their current values through this endpoint, making it essential for debugging and monitoring setup."

---

### Question 235: How do you publish custom application metrics?

**Answer:**
(See 225).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you publish custom application metrics?
**Your Response:** "I publish custom application metrics using Micrometer's `MeterRegistry`. I inject the registry and create metrics like counters for business events, timers for operation performance, or gauges for current state. For example, I can create a counter for user registrations, a timer for API response times, or a gauge for current active sessions. These custom metrics appear alongside the built-in metrics, giving me complete visibility into both technical and business aspects of my application. This approach makes monitoring more meaningful by tracking what matters to my specific application."

---

### Question 236: How do you log long GC pauses with Spring Boot?

**Answer:**
Micrometer binds `jvm.gc.pause` metrics.
To LOG them: Enable GC Logging flags in JVM (`-Xlog:gc*`) or register a `JvmGcMetrics` listener that logs duration > threshold.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you log long GC pauses with Spring Boot?
**Your Response:** "Micrometer automatically binds `jvm.gc.pause` metrics for garbage collection pauses. To actually log long pauses, I have two options. I can enable JVM GC logging with flags like `-Xlog:gc*` to get detailed GC information. Alternatively, I can register a custom listener that monitors the `JvmGcMetrics` and logs when pause durations exceed a threshold. This helps me identify performance issues caused by garbage collection. While the metrics are great for monitoring, logging long pauses provides immediate visibility into GC problems that might be affecting application performance."

---

### Question 237: What is the difference between `health`, `info`, and `metrics` endpoints?

**Answer:**
- **Health:** Status (Up/Down) for Liveness/Readiness probes.
- **Info:** Static text (Version, Description).
- **Metrics:** Numeric time-series data (CPU, Request Count).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `health`, `info`, and `metrics` endpoints?
**Your Response:** "These endpoints serve different monitoring purposes. The `health` endpoint shows liveness and readiness status - whether the application is running and ready to serve traffic. The `info` endpoint provides static information like version, build time, or environment details. The `metrics` endpoint exposes numeric time-series data like CPU usage, request counts, and response times. Health is for orchestration systems, info is for identification, and metrics are for performance monitoring and alerting. Together they provide comprehensive observability of the application's state and behavior."

---

### Question 238: How to use actuator with Spring Cloud Config?

**Answer:**
`/actuator/refresh` endpoint allows reloading `@ConfigurationProperties` from the Config Server without restarting the app.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to use actuator with Spring Cloud Config?
**Your Response:** "I use the `/actuator/refresh` endpoint to reload configuration properties from the Config Server without restarting the application. When I make changes to configuration in the Config Server, I can call POST `/actuator/refresh` and Spring Boot reloads `@ConfigurationProperties` beans with the new values. This is particularly useful for updating configuration in production without downtime. The refresh endpoint only affects configuration properties annotated with `@RefreshScope`, giving me control over which configurations can be dynamically updated."

---

### Question 239: How to restrict access to sensitive actuator endpoints?

**Answer:**
(See Q188). Firewall/Security Config is mandatory.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to restrict access to sensitive actuator endpoints?
**Your Response:** "I restrict access to sensitive actuator endpoints using a combination of Spring Security configuration and network firewalls. In my security configuration, I use `requestMatchers(EndpointRequest.to(ShutdownEndpoint.class, EnvEndpoint.class)).hasRole('ADMIN')` to restrict sensitive endpoints to administrators only. I also configure network firewalls to only allow access from trusted IP ranges or monitoring systems. This defense-in-depth approach ensures that critical management endpoints are protected both at the application level and network level, preventing unauthorized access to sensitive operations."

---

### Question 240: What metrics are collected by default in Spring Boot?

**Answer:**
JVM (Memory, GC, Threads).
Tomcat (Sessions, Requests).
System (CPU, Uptime).
Logback (Log events count).
HttpClient/DataSource usage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What metrics are collected by default in Spring Boot?
**Your Response:** "Spring Boot automatically collects a comprehensive set of metrics when Micrometer is on the classpath. This includes JVM metrics like memory usage, garbage collection, and thread counts. For web applications, it collects Tomcat metrics like sessions, requests, and response times. System metrics include CPU usage and uptime. It also tracks logback events, HTTP client usage, and database connection pool usage. These built-in metrics give me immediate visibility into the application's performance and resource utilization without any custom instrumentation, forming a solid foundation for application monitoring."

---
