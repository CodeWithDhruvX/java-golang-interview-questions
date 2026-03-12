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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the internals of Spring Boot's startup process?
**Your Response:** "Spring Boot's startup process follows a well-defined sequence. First, it initializes the `SpringApplication` instance and runs generic listeners. Then it prepares the Environment by loading properties and profiles. Next, it creates the appropriate ApplicationContext - either Servlet or Reactive based on the classpath. After that, it refreshes the context which triggers bean creation and dependency injection. Finally, it calls any `CommandLineRunner` or `ApplicationRunner` beans. This orchestrated process ensures everything is properly initialized before the application starts serving requests."

---

### Question 402: How does Spring Boot reduce boilerplate configuration?

**Answer:**
1.  **Opinionated Defaults:** standard locations for views, static files.
2.  **Auto-Configuration:** Configures beans based on classpath.
3.  **Starters:** Curated dependencies.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot reduce boilerplate configuration?
**Your Response:** "Spring Boot reduces boilerplate through three key mechanisms. First, it provides opinionated defaults - standard locations for views, static files, and configurations that work out of the box. Second, it uses auto-configuration to automatically set up beans based on the classpath - if I include a database starter, it configures the datasource automatically. Third, it provides starter dependencies that aggregate related libraries with compatible versions. This combination eliminates most of the manual configuration typically required in Spring applications."

---

### Question 403: What is a starter dependency and how does it work?

**Answer:**
It is an empty JAR (Maven POM) that aggregates related dependencies using `<dependencies>`.
It ensures all libraries needed for a feature (e.g., Data JPA) are present and version-compatible.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a starter dependency and how does it work?
**Your Response:** "A starter dependency is essentially an empty JAR that aggregates related dependencies using Maven's dependency management. For example, `spring-boot-starter-data-jpa` includes all the libraries needed for JPA - Hibernate, Spring Data JPA, and the appropriate database drivers. The starter ensures all these libraries are present and version-compatible. This simplifies dependency management - I just add one starter instead of managing multiple individual dependencies and their versions."

---

### Question 404: How does Spring Boot decide which auto-configurations to apply?

**Answer:**
It evaluates `@Conditional...` annotations on the AutoConfiguration classes found in `spring.factories`.
Checks for Classes, Beans, Properties, and Resources.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot decide which auto-configurations to apply?
**Your Response:** "Spring Boot decides which auto-configurations to apply by evaluating conditional annotations on the AutoConfiguration classes found in `spring.factories`. It checks for the presence of specific classes, beans, properties, or resources. For example, it only configures a datasource if it finds a DataSource class on the classpath and no existing DataSource bean. This conditional approach ensures that auto-configuration only applies when appropriate, avoiding unnecessary bean creation and potential conflicts."

---

### Question 405: What is the significance of `spring.factories` in Spring Boot?

**Answer:**
(See Q112). It is the mechanism to discover AutoConfiguration classes from JAR files on the classpath.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the significance of `spring.factories` in Spring Boot?
**Your Response:** "`spring.factories` is the discovery mechanism that allows Spring Boot to find AutoConfiguration classes from JAR files on the classpath. Each starter can include a `META-INF/spring.factories` file that lists its auto-configuration classes. Spring Boot scans these files during startup to find all available auto-configurations. This enables the modular design of Spring Boot where each starter can contribute its own auto-configuration without requiring explicit registration."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `@EnableConfigurationProperties`?
**Your Response:** "`@EnableConfigurationProperties` enables support for `@ConfigurationProperties` annotated beans. I typically use it on a `@Configuration` class to explicitly register property POJOs like `@EnableConfigurationProperties(MyProps.class)`. While Spring Boot can often auto-detect these classes, explicitly enabling them provides better control and clearer documentation of which configuration classes are being used. This annotation ensures that the configuration properties are properly validated and bound to the Environment."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you enable and customize access logs in Spring Boot?
**Your Response:** "I enable access logs by setting `server.tomcat.accesslog.enabled=true` in `application.properties`. I can customize the pattern with `server.tomcat.accesslog.pattern=%t %a '%r' %s (%D ms)` to control what information is logged. This logs every HTTP request hitting the embedded Tomcat server, which is invaluable for debugging, monitoring, and security auditing. The access logs show timestamps, client IPs, request lines, status codes, and response times, helping me understand how my application is being used."

---

### Question 422: How do you configure JSON logging format in Spring Boot?

**Answer:**
Default is plain text.
To get JSON (Logstash format), use a library like `logstash-logback-encoder`.
Add dependency, create `logback-spring.xml`, and configure the `LogstashEncoder`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure JSON logging format in Spring Boot?
**Your Response:** "By default, Spring Boot logs in plain text format. To get JSON logging in Logstash format, I add the `logstash-logback-encoder` dependency and create a `logback-spring.xml` configuration file. In this file, I configure the `LogstashEncoder` to output structured JSON logs. This format is much easier to parse and analyze with log aggregation tools like ELK stack. JSON logs provide consistent structure that makes searching and filtering logs more efficient in production environments."

---

### Question 423: How can you implement centralized logging with ELK stack in Spring Boot?

**Answer:**
1.  **App:** Writes logs to file (or stdout) in JSON format.
2.  **Filebeat:** Reads log file, pushes to Logstash/Elasticsearch.
3.  **Kibana:** Visualizes logs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can you implement centralized logging with ELK stack in Spring Boot?
**Your Response:** "I implement centralized logging with the ELK stack in three steps. First, I configure the Spring Boot application to write logs in JSON format to a file or stdout. Second, I use Filebeat to read the log files and push them to Logstash or directly to Elasticsearch. Third, I use Kibana to visualize and search the logs. This centralized approach allows me to aggregate logs from multiple application instances, making it easier to monitor and troubleshoot distributed systems."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are some built-in health checks provided by Spring Boot Actuator?
**Your Response:** "Spring Boot Actuator provides several built-in health indicators out of the box. The `DataSourceHealthIndicator` checks database connectivity, `DiskSpaceHealthIndicator` monitors available disk space, `MongoHealthIndicator` verifies MongoDB connectivity, and `RabbitHealthIndicator` checks RabbitMQ connection status. These indicators automatically contribute to the overall health status at `/actuator/health`. I can also add custom health indicators to monitor application-specific components, giving me comprehensive visibility into system health."

---

### Question 427: How do you monitor thread usage in a running Spring Boot app?

**Answer:**
Actuator Metrics: `jvm.threads.live`, `jvm.threads.peak`.
Thread Dump Endpoint: `/actuator/threaddump`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you monitor thread usage in a running Spring Boot app?
**Your Response:** "I monitor thread usage through Actuator metrics and endpoints. The metrics endpoint provides `jvm.threads.live` and `jvm.threads.peak` counters that show current and peak thread counts. For detailed analysis, I use the `/actuator/threaddump` endpoint which generates a thread dump showing all running threads, their states, and stack traces. This combination helps me identify thread leaks, deadlocks, or performance issues related to thread contention in production applications."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you push metrics to a time-series database from Spring Boot?
**Your Response:** "I push metrics to a time-series database by adding the specific Micrometer registry dependency, like `micrometer-registry-influx` for InfluxDB. I configure the database URL and authentication in `application.properties`. Micrometer then automatically pushes metrics periodically to the configured database. This integration works with various time-series databases including Prometheus, InfluxDB, and Graphite, allowing me to choose the best monitoring solution for my infrastructure while using the same Micrometer metrics API."

---

### Question 430: How do you configure log levels dynamically at runtime?

**Answer:**
Endpoint: `/actuator/loggers/{package}`.
POST requested payload: `{"configuredLevel": "DEBUG"}`.
Changes level instantly without restart.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure log levels dynamically at runtime?
**Your Response:** "I configure log levels dynamically using the Actuator endpoint `/actuator/loggers/{package}`. I send a POST request with a payload like `{'configuredLevel': 'DEBUG'}` to change the log level for a specific package or logger. The change takes effect instantly without requiring an application restart. This is incredibly useful for debugging production issues - I can temporarily enable debug logging for a problematic component, collect the detailed logs, and then restore the normal level without any downtime."

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
