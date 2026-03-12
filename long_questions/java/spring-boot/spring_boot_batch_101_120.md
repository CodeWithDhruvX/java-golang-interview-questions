## 🔹 Section 1: Memory & Core Internals (101-120)

### Question 101: How does Spring Boot handle backward compatibility across versions?

**Answer:**
Spring Boot follows strict semantic versioning. 
- The team maintains a **Compatibility Matrix**.
- **Deprecation Cycle:** Features are marked `@Deprecated` in minor releases (e.g., 2.6) and removed in major ones (e.g., 3.0).
- **Configuration Migrator:** The `spring-boot-properties-migrator` module logs warnings at startup if you use keys that have been renamed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot handle backward compatibility across versions?
**Your Response:** "Spring Boot takes backward compatibility very seriously through strict semantic versioning. The team maintains a comprehensive compatibility matrix to track what works across versions. They follow a deprecation cycle where features are marked with `@Deprecated` in minor releases like 2.6, giving developers time to update, and then removed in major releases like 3.0. There's also a configuration migrator module that logs warnings at startup if I'm using property keys that have been renamed. This gradual approach gives me plenty of time to migrate and avoids breaking changes in production."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens internally when a Spring Boot application starts?
**Your Response:** "Spring Boot startup is a well-orchestrated process. It starts with a StopWatch to measure startup time, then fires various listener events like ApplicationStartingEvent. The environment is prepared with profiles and properties, the banner is printed, and then the ApplicationContext is created - for web apps, it's typically AnnotationConfigServletWebServerApplicationContext. The critical step is the context refresh where all beans are created and dependency injection happens. Finally, any CommandLineRunner or ApplicationRunner beans are executed. This whole process is highly optimized and customizable, which is why Spring Boot starts so quickly compared to traditional Spring applications."

---

### Question 103: What are banner files in Spring Boot, and how do you customize them?

**Answer:**
The ASCII art shown at startup.
Place a `banner.txt` file in `src/main/resources`.
You can use variables:
- `${spring-boot.version}`
- `${application.title}`
- ANSI colors `${AnsiColor.RED}`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are banner files in Spring Boot, and how do you customize them?
**Your Response:** "The banner is the ASCII art that displays when Spring Boot starts up. I can customize it by creating a `banner.txt` file in `src/main/resources`. What's cool is that I can use variables like `${spring-boot.version}` to show the Spring Boot version, `${application.title}` for my app name, and even ANSI color codes like `${AnsiColor.RED}` to add colors. This makes the startup experience more personalized and branded. It's a small touch but really helpful for identifying which application is starting, especially when running multiple services in development."

---

### Question 104: What is lazy initialization in Spring Boot 2.2+?

**Answer:**
A global property to defer bean creation.
`spring.main.lazy-initialization=true`.
**Pros:** Faster startup (Good for dev loop).
**Cons:** If a bean configuration is broken, you won't know until the first HTTP request hits that bean (Runtime failure instead of Startup failure).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is lazy initialization in Spring Boot 2.2+?
**Your Response:** "Lazy initialization is a global property I can set with `spring.main.lazy-initialization=true` that defers bean creation until they're actually needed. The main benefit is faster startup time, which is great for development loops where I'm frequently restarting the application. However, there's a trade-off - if a bean has configuration issues, I won't find out until the first time that bean is accessed, which means runtime failures instead of startup failures. This can make debugging harder, so I typically use it for development but disable it for production to catch configuration errors early."

---

### Question 105: How do you disable specific auto-configurations in Spring Boot?

**Answer:**
If you want to use MongoDB manually but keep the dependency:
```java
@SpringBootApplication(exclude = { MongoAutoConfiguration.class })
public class App { ... }
```
Or via properties: `spring.autoconfigure.exclude=...`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you disable specific auto-configurations in Spring Boot?
**Your Response:** "Sometimes I need to disable specific auto-configurations when I want to handle things manually. I can do this either through the `@SpringBootApplication` annotation using the exclude parameter, like `@SpringBootApplication(exclude = { MongoAutoConfiguration.class })`, or through properties with `spring.autoconfigure.exclude=...`. This is useful when I have the dependency on the classpath but want to configure the component myself rather than using Spring Boot's defaults. For example, I might want to configure MongoDB manually while still using the spring-data-mongodb dependency for other features."

---

### Question 106: How do conditional annotations like `@ConditionalOnClass` work?

**Answer:**
They use the `Condition` interface (matches method).
During startup, Spring checks the classpath.
If the bytecode for the specified Class is NOT found (via `Class.forName()`), the condition fails, and the associated Bean/Configuration is skipped. This is the magic behind "Starters".

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do conditional annotations like `@ConditionalOnClass` work?
**Your Response:** "Conditional annotations are the magic behind Spring Boot's intelligent auto-configuration. They work by implementing the `Condition` interface with a matches method. During startup, Spring checks the classpath for specific classes using techniques like `Class.forName()`. For `@ConditionalOnClass`, if the specified class bytecode is not found on the classpath, the condition fails and the associated bean or configuration is skipped. This is how Spring Boot starters work - they only configure components when the required dependencies are present, making the application lightweight and adaptable to different environments."

---

### Question 107: What is the difference between `spring.main.web-application-type=reactive` vs `servlet`?

**Answer:**
Determines the `ApplicationContext` implementation.
- **Servlet:** Uses `Tomcat`/`Jetty`. Blocking I/O. Standard Spring MVC.
- **Reactive:** Uses `Netty`. Non-blocking. Spring WebFlux.
- **None:** No web server started (Batch jobs).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `spring.main.web-application-type=reactive` vs `servlet`?
**Your Response:** "This property determines which ApplicationContext implementation Spring Boot uses. The servlet type uses traditional Tomcat or Jetty servers with blocking I/O and standard Spring MVC. The reactive type uses Netty server with non-blocking I/O for Spring WebFlux applications. There's also a 'none' option for applications that don't need a web server at all, like batch processing jobs. The choice affects everything from the underlying server technology to the programming model - servlet uses traditional request-per-thread while reactive uses event-loop architecture. I choose based on whether I need high concurrency with reactive programming or prefer traditional MVC."

---

### Question 108: How to exclude a Spring Boot starter dependency?

**Answer:**
In `pom.xml` (Maven), use `<exclusion>`.
Example: Using `starter-web` but want `Jetty` instead of `Tomcat`.
1.  Exclude `spring-boot-starter-tomcat` from `spring-boot-starter-web`.
2.  Add `spring-boot-starter-jetty`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to exclude a Spring Boot starter dependency?
**Your Response:** "Sometimes I need to replace default dependencies in Spring Boot starters. For example, if I want to use Jetty instead of Tomcat, I can exclude the default and include my preferred option. In Maven's pom.xml, I use the `<exclusion>` tag to exclude `spring-boot-starter-tomcat` from `spring-boot-starter-web`, then add `spring-boot-starter-jetty` as a separate dependency. This gives me the flexibility to customize the technology stack while still benefiting from the starter's convenience. The same approach works for other default dependencies that I might want to replace with alternatives."

---

### Question 109: Can you create a multi-module Spring Boot project?

**Answer:**
Yes. Standard Maven/Gradle multi-module structure.
- **Parent Pom:** Manages versions.
- **Core Module:** Domain/Service logic.
- **Web Module:** Controllers +Main Class (depends on Core).
Only the Web Module acts as the Spring Boot entry point.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you create a multi-module Spring Boot project?
**Your Response:** "Yes, I frequently create multi-module Spring Boot projects using standard Maven or Gradle structures. I typically have a parent POM that manages versions and common dependencies, a core module containing domain and service logic, and a web module with controllers and the main class that depends on the core module. Only the web module acts as the Spring Boot entry point - it contains the main method and the `@SpringBootApplication` annotation. This separation allows me to reuse core business logic across different applications, like a REST API and a batch processor, while keeping the web-specific concerns isolated."

---

### Question 110: How does Spring Boot handle memory management?

**Answer:**
Spring Boot itself is just Java.
However, it provides **metadata** for native images (GraalVM) to reduce footprint.
Configuration properties like `server.tomcat.max-threads` impact heap usage.
Actuator helps monitor Heap/Non-Heap usage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot handle memory management?
**Your Response:** "Spring Boot itself runs on the JVM, so memory management follows standard Java practices. However, Spring Boot provides features to help optimize memory usage. It generates metadata for GraalVM native images, which can significantly reduce memory footprint. Configuration properties like `server.tomcat.max-threads` directly impact heap usage by controlling thread pools. Spring Boot Actuator provides endpoints to monitor heap and non-heap memory usage in production. While Spring Boot doesn't manage memory directly, it gives me the tools and configurations to optimize memory usage for different deployment scenarios."

---

### Question 111: How to set JVM arguments in Spring Boot applications?

**Answer:**
They are arguments to the `java` command, NOT `application.properties`.
`java -Xmx512m -Dspring.profiles.active=dev -jar app.jar`.
In Docker, usually passed via `JAVA_OPTS` environment variable if the entrypoint script supports it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to set JVM arguments in Spring Boot applications?
**Your Response:** "JVM arguments are passed to the `java` command itself, not through Spring Boot's application.properties. I run applications like `java -Xmx512m -Dspring.profiles.active=dev -jar app.jar`. The `-Xmx` sets maximum heap size, and `-D` sets system properties. In containerized environments like Docker, I typically pass these through the `JAVA_OPTS` environment variable, assuming the entrypoint script supports it. It's important to distinguish between JVM arguments and Spring Boot properties - JVM arguments configure the Java runtime, while Spring Boot properties configure the application itself."

---

### Question 112: What is the role of the `META-INF/spring.factories` file?

**Answer:**
(Legacy but still critical).
It is the SPI (Service Provider Interface) registry for Spring Boot.
It maps interfaces to implementation classes.
Crucially, it maps `EnableAutoConfiguration` key to the list of all AutoConfig classes (e.g., `DataSourceAutoConfiguration`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of the `META-INF/spring.factories` file?
**Your Response:** "The `META-INF/spring.factories` file is a legacy but still critical SPI registry for Spring Boot. It maps interfaces to implementation classes, but most importantly, it maps the `EnableAutoConfiguration` key to the list of all auto-configuration classes like `DataSourceAutoConfiguration`. This file is how Spring Boot discovers and loads auto-configuration classes from all the dependencies on the classpath. While newer Spring Boot versions use a different mechanism, understanding spring.factories is important for working with older libraries and understanding how Spring Boot's extensibility works under the hood."

---

### Question 113: How does Spring Boot support reactive programming?

**Answer:**
Via **Spring WebFlux** and **Project Reactor**.
- **Netty** Server (Event Loop).
- **Reactive Repositories** (R2DBC, Reactive Mongo).
- **WebClient** for async HTTP calls.
Allows handling thousands of concurrent requests with few threads.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot support reactive programming?
**Your Response:** "Spring Boot supports reactive programming through Spring WebFlux and Project Reactor. It uses Netty as the default server with an event-loop architecture instead of traditional thread-per-request. I can use reactive repositories with R2DBC or reactive MongoDB, and WebClient for async HTTP calls. This approach allows handling thousands of concurrent connections with just a few threads, making it ideal for I/O-bound applications that need high throughput. The programming model uses reactive types like Mono and Flux, which represent streams of 0-1 and 0-N items respectively, enabling non-blocking, backpressure-aware applications."

---

### Question 114: How to enable debugging logs for specific packages?

**Answer:**
In `application.properties`:
`logging.level.com.mycompany.service=DEBUG`
`logging.level.org.springframework.web=TRACE` (to see request mapping details).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to enable debugging logs for specific packages?
**Your Response:** "I can enable debug logging for specific packages in Spring Boot using the logging.level properties in application.properties. For example, `logging.level.com.mycompany.service=DEBUG` enables debug logging for my service classes, and `logging.level.org.springframework.web=TRACE` shows detailed request mapping information. This granular control lets me debug specific parts of my application without being overwhelmed by logs from the entire framework. I can set different levels for different packages - DEBUG for my code, INFO for Spring, and TRACE for specific components I'm troubleshooting."

---

### Question 115: What are some common pitfalls when using Spring Boot in production?

**Answer:**
1.  **Defaults:** Leaving default passwords (Security) or default connection pool settings.
2.  **Actuator Security:** Exposing sensitive endpoints to public internet.
3.  **Memory:** Not setting Heap limits (`-Xmx`) in containerized environments.
4.  **Slow Startup:** Due to scanning too many packages or eager Hibernate validation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are some common pitfalls when using Spring Boot in production?
**Your Response:** "I've seen several common production pitfalls. First, leaving default security settings like default passwords can create serious vulnerabilities. Second, exposing sensitive Actuator endpoints to the public internet without proper security. Third, not setting heap limits with `-Xmx` in containerized environments can cause memory issues. Fourth, slow startup due to scanning too many packages or eager Hibernate validation. I always recommend customizing security settings, securing Actuator endpoints, setting appropriate JVM memory limits, and optimizing startup configuration for production environments. These small steps prevent major production issues."

---

### Question 116: What is the Spring Boot DevTools restart strategy?

**Answer:**
It uses two ClassLoaders.
1.  **Base ClassLoader:** Loads immutable dependencies (JARs).
2.  **Restart ClassLoader:** Loads your project code.
When code changes, only the Restart ClassLoader is thrown away and recreated. Much faster than a "Cold Start".

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the Spring Boot DevTools restart strategy?
**Your Response:** "Spring Boot DevTools uses a clever two-classloader strategy for fast restarts. It has a Base ClassLoader that loads immutable dependencies like third-party JARs, and a Restart ClassLoader that loads my project code. When I make changes to my code, DevTools only discards and recreates the Restart ClassLoader, keeping the Base ClassLoader intact. This is much faster than a full application restart because it doesn't need to reload all the dependencies. The result is near-instant restarts during development while maintaining a clean classloader structure."

---

### Question 117: What is the difference between `@Import` and `@ComponentScan`?

**Answer:**
- **`@ComponentScan`:** Scans a package for *any* stereotype (`@Component`, etc). Implicit.
- **`@Import`:** Explicitly registers a specific class (Configuration or Component) as a bean, regardless of package location. Useful for library/module integration.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@Import` and `@ComponentScan`?
**Your Response:** "`@ComponentScan` and `@Import` serve different purposes for bean registration. `@ComponentScan` automatically scans a package for any stereotype annotations like `@Component`, `@Service`, etc. - it's implicit and broad. `@Import` explicitly registers specific classes as beans regardless of their package location. I use `@Import` when I need precise control over which configuration classes or components to load, which is particularly useful for library integration or module composition. `@ComponentScan` is for general component discovery, while `@Import` is for targeted, explicit bean registration."

---

### Question 118: How to create a custom auto-configuration module?

**Answer:**
1.  Create a `@Configuration` class.
2.  Add conditions (`@ConditionalOnClass`).
3.  Register it in `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports`.
4.  Package it as a JAR.
Any app depending on this JAR gets the config automatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to create a custom auto-configuration module?
**Your Response:** "Creating a custom auto-configuration module involves a few key steps. First, I create a `@Configuration` class with the beans I want to auto-configure. I add conditional annotations like `@ConditionalOnClass` so the configuration only applies when required dependencies are present. Then I register the configuration in the newer `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports` file. When I package this as a JAR, any application that depends on it gets the auto-configuration automatically, just like Spring Boot's built-in starters. This is how I can create reusable, opinionated configurations for my organization or for open-source libraries."

---

### Question 119: How does Spring Boot detect dependencies for conditional configuration?

**Answer:**
Using the `TypeFilter` and bytecode analysis (ASM) during scanning generally, but specifically for `@ConditionalOnClass`, it attempts to load the class using `ClassUtils.isPresent()`. If it returns false, the config is skipped.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot detect dependencies for conditional configuration?
**Your Response:** "Spring Boot uses sophisticated dependency detection for conditional configuration. For general scanning, it uses TypeFilter and bytecode analysis with ASM. But specifically for `@ConditionalOnClass`, it uses `ClassUtils.isPresent()` to attempt loading the class. If the class can't be loaded, the condition fails and the configuration is skipped. This approach is efficient because it doesn't require full classpath scanning - it just checks for the presence of specific classes. This is how Spring Boot can quickly determine which auto-configurations to apply based on what's actually available on the classpath, making the startup process fast and adaptive."

---

### Question 120: What is the role of `spring.factories` vs `spring.components`?

**Answer:**
`spring.components` (via `spring-context-indexer`) is an index of all components generated at compile time.
If present, Spring avoids CLASSPATH SCANNING (slow) and just reads the index (fast).
Massively improves startup time for large apps.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `spring.factories` vs `spring.components`?
**Your Response:** "`spring.components` represents a newer approach for optimizing Spring Boot startup. It's generated at compile time by the spring-context-indexer and contains an index of all components. When this index is present, Spring can avoid the expensive classpath scanning process and just read the pre-built index instead. This massively improves startup time for large applications with many components. While `spring.factories` is about auto-configuration discovery, `spring.components` is about component discovery optimization. The combination of both gives Spring Boot fast startup times even in complex applications with hundreds of beans."

---
