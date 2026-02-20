# 28. Spring Boot 3.0 & Advanced Concepts

**Q: Major Changes in Spring Boot 3.0**
> "The biggest shift is the **baseline upgrade**.
> 1.  **Java 17 is mandatory**. You cannot run Spring Boot 3 on Java 8 or 11.
> 2.  **Jakarta EE 9 APIs**: They renamed `javax.*` packages to `jakarta.*`. This breaks old libraries (like Hibernate 5). You must upgrade to Hibernate 6 and Jakarta servlet containers (Tomcat 10).
> 3.  **Native Image Support**: Official GraalVM support is now built-in."

**Indepth:**
> **Observability**: Boot 3 also standardizes Observability with Micrometer. Tracing (formerly Sleuth) and Metrics are now unified APIs, making it easier to plug in generic monitoring tools without custom vendor code.


---

**Q: Javax vs Jakarta Migration**
> "It's purely legal. Oracle donated Java EE to the Eclipse Foundation, but they kept the rights to the name 'javax' and 'Java'.
> So, Eclipse renamed everything to 'Jakarta'.
>
> **Impact**: If you upgrade a project to Boot 3, you have to Find & Replace `import javax.servlet` with `import jakarta.servlet`, `javax.persistence` with `jakarta.persistence`, etc."

**Indepth:**
> **Automation**: Use the OpenRewrite tool (specifically the Spring Boot 3 migration recipe) to automatically refactor your codebase. It scans your source files and updates the imports for you safely.


---

**Q: AOT Compilation & Native Images**
> "**AOT (Ahead-of-Time)** compilation means converting your Java byte-code into a native binary (like an .exe file) *at build time*, not runtime.
>
> **Native Image**: The result is a standalone executable.
> *   **Startup Time**: Instant (milliseconds vs seconds).
> *   **Memory**: Tiny footprint.
> *   **JIT**: No JIT optimization at runtime. What you build is what you run."

**Indepth:**
> **Limitations**: Native images do NOT support dynamic class loading or reflection easily. You must provide "Hints" (configuration files) telling GraalVM exactly which classes need reflection. Spring Boot 3 does 90% of this for you automatically.


---

**Q: Distributed Tracing (Micrometer vs Sleuth)**
> "Spring Cloud Sleuth is **removed** in Boot 3.
> It has been replaced by **Micrometer Tracing**.
>
> If you used Sleuth for trace IDs + Zipkin, you now need to add `micrometer-tracing-bridge-brave` or `micrometer-tracing-bridge-otel` (OpenTelemetry). The logic is similar, but the underlying library has standardized on Micrometer."

**Indepth:**
> **W3C Standard**: The biggest change in Micrometer Tracing is that it enforces the W3C Trace Context standard (traceparent header) by default, replacing the old proprietary B3 headers used by Zipkin.


---

**Q: EnvironmentPostProcessor**
> "This is a power-user interface. It runs **way before** the ApplicationContext is created.
> It lets you manipulate the `Environment` (properties, profiles) very early in the boot process.
>
> **Use Case**: Decrypting database passwords from a file and injecting them as system properties before Spring starts loading beans."

**Indepth:**
> **Registration**: You must register this class in `META-INF/spring.factories` (or the new `imports` file in Boot 3) because it runs before component scanning even starts.


---

**Q: FailureAnalyzer**
> "You know those nice 'Application Failed to Start' error messages with a big 'ACTION:' section?
> A `FailureAnalyzer` creates those.
> If your library throws a specific exception, you can write an analyzer to intercept it and print a human-readable diagnosis instead of a raw stack trace."

**Indepth:**
> **UX**: A good FailureAnalyzer is the difference between a developer staring at a 500-line stack trace for an hour versus fixing a "Port 8080 already in use" error in 5 seconds.


---

**Q: @Configuration proxyBeanMethods=false**
> "By default (`true`), Spring creates a CGLIB proxy for your `@Configuration` class.
> This ensures that if you call `beanA()` inside `beanB()`, you get the **same shared instance** (Singleton).
>
> Setting `false` (Lite Mode) disables this proxying. Method calls became standard Java calls (new instance every time).
> **Why do it?** It's faster and saves memory. Use it if your beans don't depend on each other directly."

**Indepth:**
> **Inter-bean Dependencies**: Be careful. If `false`, calling `beanA()` from `beanB()` creates a *new* instance of A. If A implies a database connection, you might accidentally create multiple connection pools.


---

**Q: Lazy Initialization**
> "In Spring Boot 2.2+, you can enable global lazy initialization: `spring.main.lazy-initialization=true`.
>
> *   **Normal**: All beans are created at startup. Fast first request, slow startup.
> *   **Lazy**: Beans are created only when needed. Fast startup, potentially slow first request.
> Use it for dev environments to iterate faster, but be careful in production (you might miss configuration errors until a user hits that specific endpoint)."

**Indepth:**
> **Memory**: Lazy init saves specific heap memory during startup, but it shifts the CPU spike to runtime. In Kubernetes (where liveness probes check startup), it helps the pod start faster and pass health checks sooner.


---

**Q: Conditional Annotations (@ConditionalOnClass)**
> "This is the magic behind Spring Boot's 'Auto Configuration'.
>
> `@ConditionalOnClass(name = "com.mysql.jdbc.Driver")`
>
> Spring checks: 'Is the MySQL driver on the classpath?'
> *   **Yes**: It automatically configures a DataSource bean.
> *   **No**: It skips the configuration.
> This allows Spring Boot to adapt intelligently to whatever jars you add to your pom.xml."

**Indepth:**
> **Outcome**: You can also use `@ConditionalOnProperty` to enable features via config (`app.feature.enabled=true`) or `@ConditionalOnMissingBean` to provide a default bean only if the user hasn't defined their own.


---

**Q: Reactive vs Servlet (Web Application Type)**
> "Spring Boot detects the stack:
>
> *   **Servlet (Default)**: Uses Tomcat/Jetty. Blocking I/O. Uses `DispatcherServlet`. Standard Spring MVC.
> *   **Reactive**: Uses Netty. Non-blocking I/O. Uses `DispatcherHandler`. Spring WebFlux.
>
> You can force a specific type using `spring.main.web-application-type=reactive` if you have both libraries present."

**Indepth:**
> **Thread Model**: Servlet uses one thread per request (Thread-per-Request). Reactive uses a small number of threads (Event Loop) to handle thousands of concurrent requests. Reactive is harder to debug but scales better for high concurrency.


---

**Q: META-INF/spring.factories**
> "In Boot 2.x, this file was used to register Auto-configuration classes.
>
> **Breaking Change in 2.7/3.0**:
> It is now recommended to use `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports` instead.
> The old mechanism still works for some things, but Auto-configurations have moved to the new file format."

**Indepth:**
> **Splitting**: This change was made because the single `spring.factories` file was becoming a dumping ground for everything (Listeners, EnvironmentPostProcessors, AutoConfiguration). The new system separates them by functional interface.


---

**Q: DevTools Restart Strategy**
> "DevTools splits your classpath into two:
> 1.  **Base Classloader**: Libraries (JARs) that don't change.
> 2.  **Restart Classloader**: Your project code (classes) which change often.
>
> When you save a file, it only throws away the 'Restart Classloader' and keeps the Base one. This makes 'restarting' incredibly fast compared to a full cold start."

**Indepth:**
> **LiveReload**: DevTools also includes a LiveReload server that triggers a browser refresh when resources (CSS/JS) change. It's a massive productivity booster for full-stack developers.

