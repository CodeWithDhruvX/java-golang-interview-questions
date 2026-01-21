## 🔹 Section 1: Memory & Core Internals (101-120)

### Question 101: How does Spring Boot handle backward compatibility across versions?

**Answer:**
Spring Boot follows strict semantic versioning. 
- The team maintains a **Compatibility Matrix**.
- **Deprecation Cycle:** Features are marked `@Deprecated` in minor releases (e.g., 2.6) and removed in major ones (e.g., 3.0).
- **Configuration Migrator:** The `spring-boot-properties-migrator` module logs warnings at startup if you use keys that have been renamed.

---

### Question 102: What happens internally when a Spring Boot application starts?

**Answer:**
1.  **StopWatch** starts.
2.  **Listeners** (e.g., ApplicationStartingEvent) are fired.
3.  **Environment** is prepared (Profiles, Properties).
4.  **Banner** is printed.
5.  **ApplicationContext** is created (AnnotationConfigServletWebServerApplicationContext for web).
6.  **Refresh Context**: Beans created, DEPENDENCY INJECTION happens.
7.  **Runners**: CommandLineRunner/ApplicationRunner executred.

---

### Question 103: What are banner files in Spring Boot, and how do you customize them?

**Answer:**
The ASCII art shown at startup.
Place a `banner.txt` file in `src/main/resources`.
You can use variables:
- `${spring-boot.version}`
- `${application.title}`
- ANSI colors `${AnsiColor.RED}`.

---

### Question 104: What is lazy initialization in Spring Boot 2.2+?

**Answer:**
A global property to defer bean creation.
`spring.main.lazy-initialization=true`.
**Pros:** Faster startup (Good for dev loop).
**Cons:** If a bean configuration is broken, you won't know until the first HTTP request hits that bean (Runtime failure instead of Startup failure).

---

### Question 105: How do you disable specific auto-configurations in Spring Boot?

**Answer:**
If you want to use MongoDB manually but keep the dependency:
```java
@SpringBootApplication(exclude = { MongoAutoConfiguration.class })
public class App { ... }
```
Or via properties: `spring.autoconfigure.exclude=...`.

---

### Question 106: How do conditional annotations like `@ConditionalOnClass` work?

**Answer:**
They use the `Condition` interface (matches method).
During startup, Spring checks the classpath.
If the bytecode for the specified Class is NOT found (via `Class.forName()`), the condition fails, and the associated Bean/Configuration is skipped. This is the magic behind "Starters".

---

### Question 107: What is the difference between `spring.main.web-application-type=reactive` vs `servlet`?

**Answer:**
Determines the `ApplicationContext` implementation.
- **Servlet:** Uses `Tomcat`/`Jetty`. Blocking I/O. Standard Spring MVC.
- **Reactive:** Uses `Netty`. Non-blocking. Spring WebFlux.
- **None:** No web server started (Batch jobs).

---

### Question 108: How to exclude a Spring Boot starter dependency?

**Answer:**
In `pom.xml` (Maven), use `<exclusion>`.
Example: Using `starter-web` but want `Jetty` instead of `Tomcat`.
1.  Exclude `spring-boot-starter-tomcat` from `spring-boot-starter-web`.
2.  Add `spring-boot-starter-jetty`.

---

### Question 109: Can you create a multi-module Spring Boot project?

**Answer:**
Yes. Standard Maven/Gradle multi-module structure.
- **Parent Pom:** Manages versions.
- **Core Module:** Domain/Service logic.
- **Web Module:** Controllers +Main Class (depends on Core).
Only the Web Module acts as the Spring Boot entry point.

---

### Question 110: How does Spring Boot handle memory management?

**Answer:**
Spring Boot itself is just Java.
However, it provides **metadata** for native images (GraalVM) to reduce footprint.
Configuration properties like `server.tomcat.max-threads` impact heap usage.
Actuator helps monitor Heap/Non-Heap usage.

---

### Question 111: How to set JVM arguments in Spring Boot applications?

**Answer:**
They are arguments to the `java` command, NOT `application.properties`.
`java -Xmx512m -Dspring.profiles.active=dev -jar app.jar`.
In Docker, usually passed via `JAVA_OPTS` environment variable if the entrypoint script supports it.

---

### Question 112: What is the role of the `META-INF/spring.factories` file?

**Answer:**
(Legacy but still critical).
It is the SPI (Service Provider Interface) registry for Spring Boot.
It maps interfaces to implementation classes.
Crucially, it maps `EnableAutoConfiguration` key to the list of all AutoConfig classes (e.g., `DataSourceAutoConfiguration`).

---

### Question 113: How does Spring Boot support reactive programming?

**Answer:**
Via **Spring WebFlux** and **Project Reactor**.
- **Netty** Server (Event Loop).
- **Reactive Repositories** (R2DBC, Reactive Mongo).
- **WebClient** for async HTTP calls.
Allows handling thousands of concurrent requests with few threads.

---

### Question 114: How to enable debugging logs for specific packages?

**Answer:**
In `application.properties`:
`logging.level.com.mycompany.service=DEBUG`
`logging.level.org.springframework.web=TRACE` (to see request mapping details).

---

### Question 115: What are some common pitfalls when using Spring Boot in production?

**Answer:**
1.  **Defaults:** Leaving default passwords (Security) or default connection pool settings.
2.  **Actuator Security:** Exposing sensitive endpoints to public internet.
3.  **Memory:** Not setting Heap limits (`-Xmx`) in containerized environments.
4.  **Slow Startup:** Due to scanning too many packages or eager Hibernate validation.

---

### Question 116: What is the Spring Boot DevTools restart strategy?

**Answer:**
It uses two ClassLoaders.
1.  **Base ClassLoader:** Loads immutable dependencies (JARs).
2.  **Restart ClassLoader:** Loads your project code.
When code changes, only the Restart ClassLoader is thrown away and recreated. Much faster than a "Cold Start".

---

### Question 117: What is the difference between `@Import` and `@ComponentScan`?

**Answer:**
- **`@ComponentScan`:** Scans a package for *any* stereotype (`@Component`, etc). Implicit.
- **`@Import`:** Explicitly registers a specific class (Configuration or Component) as a bean, regardless of package location. Useful for library/module integration.

---

### Question 118: How to create a custom auto-configuration module?

**Answer:**
1.  Create a `@Configuration` class.
2.  Add conditions (`@ConditionalOnClass`).
3.  Register it in `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports`.
4.  Package it as a JAR.
Any app depending on this JAR gets the config automatically.

---

### Question 119: How does Spring Boot detect dependencies for conditional configuration?

**Answer:**
Using the `TypeFilter` and bytecode analysis (ASM) during scanning generally, but specifically for `@ConditionalOnClass`, it attempts to load the class using `ClassUtils.isPresent()`. If it returns false, the config is skipped.

---

### Question 120: What is the role of `spring.factories` vs `spring.components`?

**Answer:**
`spring.components` (via `spring-context-indexer`) is an index of all components generated at compile time.
If present, Spring avoids CLASSPATH SCANNING (slow) and just reads the index (fast).
Massively improves startup time for large apps.

---
