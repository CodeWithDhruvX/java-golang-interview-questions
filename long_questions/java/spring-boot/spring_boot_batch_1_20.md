## 🔹 Section 1: Basics of Spring Boot (1–20)

### Question 1: What is Spring Boot and how is it different from Spring?

**Answer:**
**Spring Framework** is a vast Java framework for building enterprise applications, handling dependency injection, transaction management, etc. However, it requires significant XML or Java configuration setup.

**Spring Boot** is an extension of the Spring Framework that simplifies the setup and development of Spring applications.
*   **Spring:** Requires manual configuration of DispatcherServlet, ComponentScan, ViewResolver, etc.
*   **Spring Boot:** Provides "Opinionated Defaults" and Auto-Configuration. It embeds a web server (Tomcat/Jetty) so you don't need to deploy WARs.

---

### Question 2: What are the main features of Spring Boot?

**Answer:**
1.  **Auto-Configuration:** Automatically configures beans based on classpath dependencies.
2.  **Starter Dependencies:** Simplifies Maven/Gradle config (e.g., `spring-boot-starter-web` brings in Spring MVC, Tomcat, and Jackson).
3.  **Embedded Servers:** Runs as a standalone JAR with embedded Tomcat, Jetty, or Undertow.
4.  **Actuator:** Production-ready metrics and health checks.
5.  **Externalized Configuration:** Supports properties, YAML, and environment variables.

---

### Question 3: What is the purpose of the `@SpringBootApplication` annotation?

**Answer:**
It is a convenience annotation that combines three others:
1.  **`@Configuration`**: Marks the class as a source of bean definitions.
2.  **`@EnableAutoConfiguration`**: Tells Boot to start adding beans based on classpath settings.
3.  **`@ComponentScan`**: Tells Spring to look for other components, configurations, and services in the current package.

```java
@SpringBootApplication
public class MyApp {
    public static void main(String[] args) {
        SpringApplication.run(MyApp.class, args);
    }
}
```

---

### Question 4: How does Spring Boot auto-configuration work?

**Answer:**
It uses `@EnableAutoConfiguration`.
Spring Boot scans `META-INF/spring.factories` (or `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports` in newer versions) for a list of `AutoConfiguration` classes.
It applies conditions (`@ConditionalOnClass`, `@ConditionalOnMissingBean`) to decide whether to register a bean.
*Example:* If `H2` is on the classpath and no `DataSource` bean is defined, Boot creates an in-memory H2 connection automatically.

---

### Question 5: What is `spring-boot-starter`? Name a few starters.

**Answer:**
Starters are set of convenient dependency descriptors to include in your application.
Examples:
*   `spring-boot-starter-web`: For building RESTful web services (includes Spring MVC + Jackson + Tomcat).
*   `spring-boot-starter-data-jpa`: For using Spring Data JPA + Hibernate.
*   `spring-boot-starter-test`: Includes JUnit, Mockito, Spring Test.
*   `spring-boot-starter-security`: For Spring Security.

---

### Question 6: What is the use of `application.properties` or `application.yml`?

**Answer:**
They are used for **Externalized Configuration**. You can configure port numbers, database credentials, log levels, etc., without changing code.
`application.yml` is preferred for hierarchical data.

```properties
server.port=8081
spring.datasource.url=jdbc:mysql://localhost:3306/db
```

---

### Question 7: How do you create a Spring Boot application from scratch?

**Answer:**
1.  **Spring Initializr:** Go to start.spring.io, select dependencies, generate ZIP.
2.  **IDE:** Use IntelliJ IDEA or Eclipse (Spring Tools Suite) wizard.
3.  **CLI:** `spring init --dependencies=web my-app`.
4.  **Manually:** Create a Maven/Gradle project and add the `spring-boot-starter-parent`.

---

### Question 8: How does Spring Boot reduce boilerplate code?

**Answer:**
*   **No XML Configuration:** Uses annotation-based config.
*   **Auto Config:** Eliminates manual setup of `DataSource`, `EntityManager`, `DispatcherServlet`.
*   **Starters:** Reduces the need to hunt for compatible library versions (BOM manages versions).

---

### Question 9: What is embedded Tomcat in Spring Boot?

**Answer:**
Traditionally, you produce a WAR file and deploy it to an external Tomcat server.
Spring Boot embeds the Tomcat **library** directly into the generated JAR. When you run `java -jar app.jar`, it starts Tomcat programmatically and deploys the servlet context.

---

### Question 10: What are Spring Initializr and its advantages?

**Answer:**
A web tool (start.spring.io) to bootstrap Spring Boot projects.
**Advantages:**
*   Generates correct project structure.
*   Provides `pom.xml`/`build.gradle` with selected dependencies.
*   Ensures version compatibility via the "Bill of Materials" (BOM).

---

### Question 11: Can you run Spring Boot without a web server?

**Answer:**
Yes. Use `spring-boot-starter` (without `web`) or set `spring.main.web-application-type=none`.
Useful for console applications, cron jobs, or batch processing.
Implement `CommandLineRunner` to execute logic on startup.

---

### Question 12: How do you package a Spring Boot application (JAR vs WAR)?

**Answer:**
*   **JAR (Default):** Self-contained, executable. Run with `java -jar`. Best for Microservices/Containers.
*   **WAR:** For deploying to a standalone container (like generic Tomcat/Wildfly). requires extending `SpringBootServletInitializer`.

---

### Question 13: What is the default port of Spring Boot web application?

**Answer:**
Port **8080**.

---

### Question 14: How do you change the default port in Spring Boot?

**Answer:**
1.  **Properties:** `server.port=9090` in `application.properties`.
2.  **Command Line:** `java -jar app.jar --server.port=9090`.
3.  **OS Env Var:** `SERVER_PORT=9090`.

---

### Question 15: What is a `CommandLineRunner` in Spring Boot?

**Answer:**
An interface with a `run(String... args)` method.
Beans implementing this interface are executed **after** the Spring context is fully loaded but **before** the application startup completes.
Useful for data seeding or one-time tasks.

```java
@Bean
public CommandLineRunner init(UserRepository repo) {
    return args -> repo.save(new User("Admin"));
}
```

---

### Question 16: What are actuators in Spring Boot?

**Answer:**
A feature (via `spring-boot-starter-actuator`) to monitor and manage the application in production.
Exposes endpoints like `/actuator/health`, `/actuator/info`, `/actuator/metrics`, `/actuator/env`.

---

### Question 17: How do you enable and use Spring Boot Actuator endpoints?

**Answer:**
1.  Add dependency: `spring-boot-starter-actuator`.
2.  By default, only `/health` and `/info` are exposed over HTTP.
3.  Expose all: `management.endpoints.web.exposure.include=*`.

---

### Question 18: What is DevTools in Spring Boot and how does it help?

**Answer:**
Dependency: `spring-boot-devtools`.
**Features:**
*   **Automatic Restart:** Restarts the app whenever files on classpath change (fast restart).
*   **LiveReload:** Triggers browser refresh when resources change.
*   **Property Defaults:** Disables caching for templates/static files during dev.

---

### Question 19: How does Spring Boot support externalized configuration?

**Answer:**
It allows you to run the same code in different environments just by changing config.
Priority Order (High to Low):
1.  Command Line Args.
2.  `JAVA_OPTS` / OS Env Vars.
3.  `application-prod.properties`.
4.  `application.properties`.

---

### Question 20: How can you run Spring Boot in different environments (dev, test, prod)?

**Answer:**
Use **Profiles**.
1.  Create `application-dev.properties` and `application-prod.properties`.
2.  Activate: `spring.profiles.active=dev` in `application.properties`.
3.  At runtime: `java -jar app.jar --spring.profiles.active=prod`.

---
