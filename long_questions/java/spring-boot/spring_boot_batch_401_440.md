## 🔹 Section 1: Core Spring Boot Features (401-420)

### Question 401: What are the internals of Spring Boot’s startup process?

**Answer:**
(See Q102).
1.  Initialize `SpringApplication`.
2.  Run generic listeners.
3.  Prepare Environment.
4.  Create ApplicationContext (Servlet/Reactive).
5.  Refresh Context (Bean Creation).
6.  Call Runners.

---

### Question 402: How does Spring Boot reduce boilerplate configuration?

**Answer:**
1.  **Opinionated Defaults:** standard locations for views, static files.
2.  **Auto-Configuration:** Configures beans based on classpath.
3.  **Starters:** Curated dependencies.

---

### Question 403: What is a starter dependency and how does it work?

**Answer:**
It is an empty JAR (Maven POM) that aggregates related dependencies using `<dependencies>`.
It ensures all libraries needed for a feature (e.g., Data JPA) are present and version-compatible.

---

### Question 404: How does Spring Boot decide which auto-configurations to apply?

**Answer:**
It evaluates `@Conditional...` annotations on the AutoConfiguration classes found in `spring.factories`.
Checks for Classes, Beans, Properties, and Resources.

---

### Question 405: What is the significance of `spring.factories` in Spring Boot?

**Answer:**
(See Q112). It is the mechanism to discover AutoConfiguration classes from JAR files on the classpath.

---

### Question 406: How do you programmatically disable specific auto-configurations?

**Answer:**
(See Q105).

---

### Question 407: How do you create a custom banner for your Spring Boot application?

**Answer:**
(See Q103). `banner.txt`.

---

### Question 408: What is the role of `@EnableConfigurationProperties`?

**Answer:**
It enables support for `@ConfigurationProperties` annotated beans.
Normally used on a `@Configuration` class to explicitly register property POJOs: `@EnableConfigurationProperties(MyProps.class)`.

---

### Question 409: How is `SpringApplication.run()` internally implemented?

**Answer:**
(Duplicate of 401/102).

---

### Question 410: What is the purpose of `SpringBootExceptionReporter`?

**Answer:**
(See Q290).

---

### Question 411: What are the internals of Spring Boot’s startup process?

**Answer:**
(Duplicate of 401).

---

### Question 412: How does Spring Boot reduce boilerplate configuration?

**Answer:**
(Duplicate of 402).

---

### Question 413: What is a starter dependency and how does it work?

**Answer:**
(Duplicate of 403).

---

### Question 414: How does Spring Boot decide which auto-configurations to apply?

**Answer:**
(Duplicate of 404).

---

### Question 415: What is the significance of `spring.factories` in Spring Boot?

**Answer:**
(Duplicate of 405).

---

### Question 416: How do you programmatically disable specific auto-configurations?

**Answer:**
(Duplicate of 406).

---

### Question 417: How do you create a custom banner for your Spring Boot application?

**Answer:**
(Duplicate of 407).

---

### Question 418: What is the role of `@EnableConfigurationProperties`?

**Answer:**
(Duplicate of 408).

---

### Question 419: How is `SpringApplication.run()` internally implemented?

**Answer:**
(Duplicate of 401).

---

### Question 420: What is the purpose of `SpringBootExceptionReporter`?

**Answer:**
(Duplicate of 410).

## 🔹 Section 2: Logging, Monitoring & Observability (421-440)

### Question 421: How do you enable and customize access logs in Spring Boot?

**Answer:**
In `application.properties`:
`server.tomcat.accesslog.enabled=true`.
`server.tomcat.accesslog.pattern=%t %a "%r" %s (%D ms)`.
Logs every HTTP request hitting the embedded Tomcat.

---

### Question 422: How do you configure JSON logging format in Spring Boot?

**Answer:**
Default is plain text.
To get JSON (Logstash format), use a library like `logstash-logback-encoder`.
Add dependency, create `logback-spring.xml`, and configure the `LogstashEncoder`.

---

### Question 423: How can you implement centralized logging with ELK stack in Spring Boot?

**Answer:**
1.  **App:** Writes logs to file (or stdout) in JSON format.
2.  **Filebeat:** Reads log file, pushes to Logstash/Elasticsearch.
3.  **Kibana:** Visualizes logs.

---

### Question 424: How do you create custom metrics using Micrometer?

**Answer:**
(See Q225).

---

### Question 425: How do you expose custom health indicators in Spring Boot?

**Answer:**
(See Q221).

---

### Question 426: What are some built-in health checks provided by Spring Boot Actuator?

**Answer:**
- `DataSourceHealthIndicator` (DB connectivity).
- `DiskSpaceHealthIndicator`.
- `MongoHealthIndicator`.
- `RabbitHealthIndicator`.

---

### Question 427: How do you monitor thread usage in a running Spring Boot app?

**Answer:**
Actuator Metrics: `jvm.threads.live`, `jvm.threads.peak`.
Thread Dump Endpoint: `/actuator/threaddump`.

---

### Question 428: What is the use of `/actuator/metrics` endpoint?

**Answer:**
(See Q234).

---

### Question 429: How do you push metrics to a time-series database from Spring Boot?

**Answer:**
Add the specific registry dependency (e.g., `micrometer-registry-influx`).
Configure URL/Auth in `application.properties`.
Micrometer pushes metrics periodically.

---

### Question 430: How do you configure log levels dynamically at runtime?

**Answer:**
Endpoint: `/actuator/loggers/{package}`.
POST requested payload: `{"configuredLevel": "DEBUG"}`.
Changes level instantly without restart.

---

### Question 431: How do you enable and customize access logs in Spring Boot?

**Answer:**
(Duplicate of 421).

---

### Question 432: How do you configure JSON logging format in Spring Boot?

**Answer:**
(Duplicate of 422).

---

### Question 433: How can you implement centralized logging with ELK stack in Spring Boot?

**Answer:**
(Duplicate of 423).

---

### Question 434: How do you create custom metrics using Micrometer?

**Answer:**
(Duplicate of 424).

---

### Question 435: How do you expose custom health indicators in Spring Boot?

**Answer:**
(Duplicate of 425).

---

### Question 436: What are some built-in health checks provided by Spring Boot Actuator?

**Answer:**
(Duplicate of 426).

---

### Question 437: How do you monitor thread usage in a running Spring Boot app?

**Answer:**
(Duplicate of 427).

---

### Question 438: What is the use of `/actuator/metrics` endpoint?

**Answer:**
(Duplicate of 428).

---

### Question 439: How do you push metrics to a time-series database from Spring Boot?

**Answer:**
(Duplicate of 429).

---

### Question 440: How do you configure log levels dynamically at runtime?

**Answer:**
(Duplicate of 430).

---
