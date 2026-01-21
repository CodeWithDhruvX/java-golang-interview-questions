## 🔹 Section 10: Spring Boot Internals & Bootstrapping (281-300)

### Question 281: What is the role of `SpringApplicationBuilder` in Spring Boot?

**Answer:**
A Fluent API builder for `SpringApplication`.
Allows chaining configuration:
`new SpringApplicationBuilder(Parent.class).child(Child.class).profiles("dev").run(args)`.
Useful for complex hierarchies or modifying context before run.

---

### Question 282: How to conditionally load beans based on environment variables?

**Answer:**
`@ConditionalOnProperty(name="MY_ENV_VAR", havingValue="true")`.

---

### Question 283: What is the role of `spring-boot-loader` in executable JARs?

**Answer:**
The magic that makes `java -jar` work.
Since default Java ClassLoaders cannot load JARs inside JARs.
Boot adds a custom `JarLauncher` main class and handling logic to read nested dependencies.

---

### Question 284: How to load external JARs at runtime in Spring Boot?

**Answer:**
Use `PropertiesLauncher`.
Allow passing `-Dloader.path=lib/` to add external folders to the classpath via the custom loader.

---

### Question 285: How do you implement custom logging initialization in Spring Boot?

**Answer:**
Add `ApplicationListener<ApplicationEnvironmentPreparedEvent>`.
Since Logging initializes BEFORE ApplicationContext, you cannot use Beans.
Or simply put `logback-spring.xml`.

---

### Question 286: What is a non-web Spring Boot application, and how do you build one?

**Answer:**
(See Q11). `web-application-type=none`. Implements `CommandLineRunner`.

---

### Question 287: How does Spring Boot optimize startup performance?

**Answer:**
1.  Index (`spring-context-indexer`).
2.  AOT (Ahead-of-Time) in Boot 3.
3.  Lazy Initialization.
4.  Parallelizing bean creation (experimental).

---

### Question 288: What is the difference between `CommandLineRunner` and `ApplicationRunner`?

**Answer:**
- **CommandLineRunner:** Receives raw `String[] args`.
- **ApplicationRunner:** Receives `ApplicationArguments` object (Parses `--foo=bar` into option names and values).

---

### Question 289: How to extend Spring Boot's application lifecycle?

**Answer:**
Implement `SpringApplicationRunListener`.
Has hooks for `starting()`, `environmentPrepared()`, `contextLoaded()`, etc.
Must be registered in `spring.factories`.

---

### Question 290: What is `SpringBootExceptionReporter` used for?

**Answer:**
Callback interface used to support custom reporting of failure analysis.
If startup fails, Boot asks reporters to display the error nicely.

---

### Question 291: How to inject nested configuration properties using `@ConfigurationProperties`?

**Answer:**
(See Q131).

---

### Question 292: How do you configure multiple datasources in Spring Boot?

**Answer:**
(See Q132).

---

### Question 293: What is the role of `@ConfigurationPropertiesScan`?

**Answer:**
(See Q133).

---

### Question 294: How do you validate fields in a `@ConfigurationProperties` class?

**Answer:**
(See Q134).

---

### Question 295: What are custom configuration namespaces in Spring Boot?

**Answer:**
(See Q135).

---

### Question 296: How to use `spring.profiles.group` for grouped profile activation?

**Answer:**
(See Q136).

---

### Question 297: How do you create a fail-fast configuration mechanism?

**Answer:**
(See Q137).

---

### Question 298: How to define system-wide default properties in Spring Boot?

**Answer:**
(See Q138).

---

### Question 299: How do you merge configuration files in Spring Boot?

**Answer:**
(See Q139).

---

### Question 300: What is profile-specific YAML merging behavior?

**Answer:**
(See Q140).

---
