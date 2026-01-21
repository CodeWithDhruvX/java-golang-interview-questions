## 🔹 Section 2: Advanced Configuration & Profiles (121-140)

### Question 121: What is the order of precedence for property sources in Spring Boot?

**Answer:**
(Most Important):
1.  **Command Line Args**.
2.  **JVM System Properties** (`-D`).
3.  **OS Environment Variables**.
4.  **Profile-specific Config** (`application-prod.properties`).
5.  **Standard Config** (`application.properties`).
6.  **Defaults** in `@ConfigurationProperties`.

---

### Question 122: How do you use the `EnvironmentPostProcessor` interface?

**Answer:**
It allows you to manipulate the `Environment` (Properties) *before* the application context is created.
Useful for decrypting passwords from Env Vars or fetching config from a remote server manually before Spring starts.
Must be registered in `spring.factories`.

---

### Question 123: What is a PropertySource and how do you load custom ones?

**Answer:**
A source of key-value pairs (File, Map, Env).
Use `@PropertySource("classpath:custom.properties")` on a configuration class to load non-standard files.

---

### Question 124: How can you externalize secrets (like passwords) securely in Spring Boot?

**Answer:**
Never commit them to git.
1.  **Environment Variables:** `DB_PASSWORD=...`.
2.  **Vault:** Use Spring Cloud Vault to fetch secrets at runtime.
3.  **K8s Secrets:** Mount as file or Env Var.
4.  **Jasypt:** Encrypt string in properties file `ENC(encrypted_string)`.

---

### Question 125: How to implement hierarchical configuration using YAML in Spring Boot?

**Answer:**
YAML supports structure naturally.
```yaml
app:
  menu:
    title: Home
    items:
      - name: A
      - name: B
```
Spring flattens this to keys: `app.menu.title`, `app.menu.items[0].name`.

---

### Question 126: How do you resolve configuration conflicts in multi-environment setups?

**Answer:**
Spring "Overlays" configuration.
Common properties go in `application.yml`.
Environment specific overrides go in `application-production.yml`.
When active profile is `production`, values in the specific file OVERRIDE values in the common file.

---

### Question 127: What is relaxed binding in Spring Boot?

**Answer:**
Spring Boot is lenient about property name formats.
Property: `my.first-name` matches:
- `@ConfigurationProperties` field `firstName` (CamelCase).
- Env Var `MY_FIRST_NAME` (Snake Case Upper).
- Property `my.first_name`.

---

### Question 128: How do you override application properties programmatically?

**Answer:**
Usually via tests.
`@SpringBootTest(properties = "server.port=0")`.
Or `System.setProperty()`.

---

### Question 129: Can you use SpEL in Spring Boot configuration files?

**Answer:**
In `.properties`: No.
in `.xml`: Yes.
Inside `@Value` annotation: Yes. `@Value("#{systemProperties['user.region']}")`.
Inside `application.yml`: No, YAML is static.

---

### Question 130: How to enable configuration reloading without restarting the app?

**Answer:**
1.  **Spring Cloud Config:** Supports `/actuator/refresh`.
    Annotate bean with `@RefreshScope`.
    Call POST `/actuator/refresh` -> Bean is re-instantiated with new properties.
2.  **DevTools:** Triggers restart, not hot reload.

---

### Question 131: How do you inject nested configuration properties using `@ConfigurationProperties`?

**Answer:**
Create nested static classes inside the properties class.
```java
@ConfigurationProperties("app")
public class AppProps {
   private final Database db = new Database();
   // getter/setter
   public static class Database { private String url; ... }
}
```

---

### Question 132: How do you configure multiple datasources in Spring Boot?

**Answer:**
1.  Define two sets of properties (`spring.datasource.one`, `spring.datasource.two`).
2.  Create two `@Configuration` classes / beans of type `DataSource`.
3.  Mark one as `@Primary`.
4.  Configure `LocalContainerEntityManagerFactoryBean` for each if using JPA.

---

### Question 133: What is the role of `@ConfigurationPropertiesScan`?

**Answer:**
Instead of adding `@EnableConfigurationProperties(MyProps.class)` on the main class, you can add `@ConfigurationPropertiesScan`.
It scans the classpath for any class annotated with `@ConfigurationProperties` and registers them as beans automatically.

---

### Question 134: How do you validate fields in a `@ConfigurationProperties` class?

**Answer:**
Add `@Validated` on the class.
Use JSR-303 annotations (`@NotNull`, `@Min`).
If validation fails at startup, Spring Boot throws an exception and stops (Fail Fast).

```java
@ConfigurationProperties("app")
@Validated
public class AppConfig {
    @NotNull
    private String name;
}
```

---

### Question 135: What are custom configuration namespaces in Spring Boot?

**Answer:**
Simply the "prefix" you choose.
`spring.*` is reserved.
You can use `acme.service.*`. Defining a configuration properties class effectively claims that namespace.

---

### Question 136: How to use `spring.profiles.group` for grouped profile activation?

**Answer:**
In `application.yml`, you can define an alias.
```yaml
spring:
  profiles:
    group:
      production: "proddb,prodmq,prodmetrics"
```
Activating `production` automatically activates `proddb`, `prodmq`, and `prodmetrics`.

---

### Question 137: How do you create a fail-fast configuration mechanism?

**Answer:**
The Validation mechanism (Q134) is fail-fast by default.
Or implement `InitializingBean` and throw exceptions manually if config is invalid.

---

### Question 138: How to define system-wide default properties in Spring Boot?

**Answer:**
`SpringApplication.setDefaultProperties(Map)`.
Sets default values that are used if no other property source defines them. lowest precedence.

---

### Question 139: How do you merge configuration files in Spring Boot?

**Answer:**
Spring Boot automatically merges `application.properties` and profile specific files.
Also supports `spring.config.import` (Boot 2.4+) to explicitly import other files.
`spring.config.import=optional:file:./custom.yml`.

---

### Question 140: What is profile-specific YAML merging behavior?

**Answer:**
Lists are **overridden**, not merged, by default.
If `common.yml` has list `[A, B]` and `dev.yml` has `[C]`. Using dev profile results in `[C]`.
Maps are merged (keys added/overwritten).

---
