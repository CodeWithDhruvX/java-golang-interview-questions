# 34. Spring Boot Architecture & Advanced Config

**Q: SpringApplicationBuilder**
> "Usually you just call `SpringApplication.run(Main.class)`.
> But if you want to customize the startup **fluently**, use the Builder.
>
> `new SpringApplicationBuilder(Main.class).bannerMode(OFF).profiles("prod").run(args);`
> It allows you to chain configuration methods before the application context is even created."

**Indepth:**
> **Hierarchy**: `SpringApplicationBuilder` allows setting parent/child contexts (rarely used now but possible). It also lets you register listeners that need to run *before* the context is created, which `application.properties` cannot do (e.g., customizing the Environment).


---

**Q: Spring Boot Loader (Executable JARs)**
> "How does `java -jar app.jar` work if your JAR contains *other* JARs inside it (nested dependencies)? Standard Java doesn't support that.
>
> **Spring Boot Loader** is a special piece of code added to the top of your JAR.
> It reads the `BOOT-INF/lib` folder, creates a custom ClassLoader handling nested JARs, and then calls your `main()` method. It’s the magic glue."

**Indepth:**
> **Manifest**: The `Manifest.MF` file has a `Main-Class` pointing to `JarLauncher` (Spring's code), and a `Start-Class` pointing to *your* Main class. `JarLauncher` creates the custom classloader, reads `BOOT-INF/lib`, and invokes your `main`.


---

**Q: Custom Logging Initialization**
> "Spring Boot automatically configures Logback if it sees `logback-spring.xml`.
>
> But if you need to do something logic-based *before* logging starts (like fetching log levels from a remote server), you need a **LoggingSystem** implementation or an `ApplicationListener` listening for `ApplicationStartingEvent`. This runs before the ApplicationContext is up."

**Indepth:**
> **MDC**: Mapped Diagnostic Context. In distributed systems, you often want to add a "Trace ID" to every log line without passing it as an argument to every method. Logging frameworks + Spring Sleuth/Micrometer Tracing use MDC to attach these context variables automatically to the thread.


---

**Q: CommandLineRunner vs ApplicationRunner**
> "Both run **after** the application starts.
>
> 1.  `CommandLineRunner`: Gives you raw String arrays: `run(String... args)`. You have to parse flags like `--port=80` yourself.
> 2.  `ApplicationRunner`: Gives you a parsed `ApplicationArguments` object. `args.getOptionValues("port")`.
>
> Always prefer `ApplicationRunner`."

**Indepth:**
> **Fail-Fast**: You can use `@Order(1)` annotation to define distinct execution order if you have multiple runners. If an exception is thrown in a Runner, the application startup **fails** (unless caught), stopping the deployment.


---

**Q: SpringBootExceptionReporter**
> "This is an internal interface used to pretty-print startup errors.
> If your app fails to start because port 8080 is in use, you don't want a 50-line stack trace. You want a clear message:
> *'Port 8080 is already in use. Identify the process or change the port.'*
> The ExceptionReporter does this formatting."

**Indepth:**
> **Extension**: This is how Spring analyzes `PortInUseException`. It's an extension point suitable for libraries that want to provide friendly error pages or console messages for their own custom startup failures.


---

**Q: @ConditionalOnMissingBean**
> "This is the most important annotation for writing reusable libraries or Starters.
>
> ```java
> @Bean
> @ConditionalOnMissingBean
> public ObjectMapper objectMapper() { ... }
> ```
> It says: 'Create this bean **only if** the user hasn't defined their own version'.
> It allows users to override your auto-configuration defaults easily."

**Indepth:**
> **Ordering**: Auto-configurations run *last* (after user configs). This ensures Spring sees the user's `@Bean` first, so when `@ConditionalOnMissingBean` runs in the auto-config, it correctly sees "Oh, a bean exists, I'll back off".


---

**Q: Circular Dependencies**
> "Bean A needs B. Bean B needs A.
> In older Spring versions, this caused a crash at startup.
>
> In recent Spring Boot versions, this is **disabled by default**.
> Steps to fix:
> 1.  **Refactor**: This is a design smell. Extract the common logic into Bean C.
> 2.  **@Lazy**: Inject one side lazily (`@Autowired @Lazy ServiceA a`). This breaks the cycle during startup."

**Indepth:**
> **Smell**: A circular dependency usually means your components are too coupled. A common fix (besides `@Lazy`) is to use an Event-Driven architecture (ApplicationEvents) so A notifies B without holding a reference to B.


---

**Q: @DependsOn**
> "Spring usually figures out the creation order based on dependency injection.
>
> But sometimes, Bean A needs Bean B to be ready, but doesn't technically *inject* it (e.g., Bean B sets up a static database connection or system property).
> `@DependsOn("beanB")` forces Spring to ensure B is created before A."

**Indepth:**
> **Legacy**: This is often needed for "static" initialization legacy code or when a Bean (like `DBMigrationBean`) must finish its work (altering tables) before the `UserRepo` bean attempts to connect.


---

**Q: Property Resolution Order**
> "Config values can come from everywhere. The hierarchy (simplified):
> 1.  **Test properties** (`@TestPropertySource`).
> 2.  **Command Line Args**.
> 3.  **OS Environment Variables** (`SERVER_PORT`).
> 4.  **Profile-specific Config** (`application-prod.yml`).
> 5.  **Standard Config** (`application.yml`).
>
> If you are confused why a value isn't changing, check if an Environment Variable is overriding your file."

**Indepth:**
> **Random**: `RandomValuePropertySource`. Spring Boot can inject random values using `${random.int}` or `${random.uuid}`. This is useful for generating unique instance IDs or ephemeral secrets for tests.


---

**Q: Encrypting Properties (Jasypt)**
> "Never commit clear-text passwords to Git.
> Use `Jasypt` (Java Simplified Encryption).
>
> `spring.datasource.password=ENC(G6v7X...)`
>
> At runtime, you pass the decryption key: `-Djasypt.encryptor.password=SECRET`.
> Spring automatically decrypts the value before injecting it into your beans."

**Indepth:**
> **Bootstrap**: The encryption password itself is the "Bootstrap Problem". You usually inject it via an Environment Variable (`JASYPT_PASSWORD`) provided by the CI/CD pipeline or the container orchestrator (K8s Secrets).


---

**Q: Externalizing Secrets**
> "For production, don't even put encrypted passwords in `application.yml`.
> Use a **Vault** (HashiCorp Vault, AWS Parameter Store, Azure Key Vault).
> Spring Cloud has specialized starters for these.
> The app starts, authenticates with the Vault, fetches secrets into memory, and injects them. The secrets never touch the disk."

**Indepth:**
> **Config Server**: Spring Cloud Config Server allows you to store properties in a Git repo and serve them to microservices at runtime. It supports encryption/decryption on the fly and dynamic reloading via `/actuator/refresh`.

