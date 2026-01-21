## 🔹 Section 1: Core Concepts & Annotations (301-310)

### Question 301: What is the use of `@ConditionalOnMissingBean` in Spring Boot?

**Answer:**
It tells Spring to register a bean **only if** a bean of that type (or name) does not already exist in the context.
This is the core implementation of "overridable defaults".
If the user defines their own `DataSource` bean, the default one annotated with `@ConditionalOnMissingBean` is skipped.

---

### Question 302: What does `@SpringBootConfiguration` do under the hood?

**Answer:**
It is simply a specialized `@Configuration` annotation.
It is used to mark the main configuration class.
Mainly used for tests (`@SpringBootTest`) to automatically find the primary configuration of the application.

---

### Question 303: How does Spring Boot handle circular dependency issues?

**Answer:**
By default in recent versions (2.6+), Circular References are **disabled** and throw an exception.
To fix:
1.  Use `@Lazy` on one of the constructor fields.
2.  Enable legacy circular handling: `spring.main.allow-circular-references=true`.
3.  Refactor architecture.

---

### Question 304: What is `@EnableAutoConfiguration`, and how does it work with `@SpringBootApplication`?

**Answer:**
(See Q3/Q28). It triggers the auto-configuration logic.
Included inside `@SpringBootApplication`.

---

### Question 305: How do you use `@Lazy` initialization in Spring Boot?

**Answer:**
(See Q34/Q104).
Annotate a Bean with `@Lazy`.
Or set `spring.main.lazy-initialization=true` to make ALL beans lazy.

---

### Question 306: What is the difference between `@Import` and `@ComponentScan`?

**Answer:**
(Duplicate of Q117).

---

### Question 307: How can you override auto-configured beans in Spring Boot?

**Answer:**
Just define a Bean of the same type in your configuration.
Because Auto-configurations usually use `@ConditionalOnMissingBean`, your bean (registered first) will cause the auto-config bean to back off.

---

### Question 308: What is the purpose of `@DependsOn` annotation?

**Answer:**
Forces specific beans to be initialized **before** the annotated bean.
Useful when there is an implicit dependency not visible via Constructor/Autowired (e.g., Static initialization, DB Schema creation).

---

### Question 309: What are factory beans in Spring Boot and when to use them?

**Answer:**
Beans that implement `FactoryBean<T>`.
Used to create complex objects (like `Proxy` objects or objects from 3rd party non-Spring libraries).
Spring injects the result of `getObject()` rather than the FactoryBean itself.

---

### Question 310: What are some common bean scope types in Spring Boot?

**Answer:**
(Duplicate of Q32). Singleton, Prototype, Request, Session.

## 🔹 Section 2: Advanced Configuration & Property Handling (311-320)

### Question 311: How do you bind lists and maps from YAML to `@ConfigurationProperties`?

**Answer:**
**List:**
```yaml
app:
  servers:
    - server1
    - server2
```
Java: `private List<String> servers;`

**Map:**
```yaml
app:
  users:
    admin: 123
    guest: 456
```
Java: `private Map<String, Integer> users;`

---

### Question 312: What is the order of precedence in Spring Boot property resolution?

**Answer:**
(Duplicate of Q121).

---

### Question 313: How do you inject system environment variables into your Spring Boot config?

**Answer:**
Reference them in `application.properties`:
`app.db.password=${DB_PASSWORD_ENV_VAR}`.
Or map directly if names match (Relaxed Binding).

---

### Question 314: How do you provide a fallback configuration in Spring Boot?

**Answer:**
Use the `:` default operator in SpEL or placeholders.
`@Value("${app.timeout:5000}")`.
If `app.timeout` is missing, use `5000`.

---

### Question 315: How can you encrypt configuration properties in Spring Boot?

**Answer:**
(See Q124). Using Jasypt.

---

### Question 316: What is the role of `PropertySourcesPlaceholderConfigurer`?

**Answer:**
Bean Post Processor that resolves `${...}` placeholders in bean definitions against the current Spring Environment properties.

---

### Question 317: How do you resolve placeholders in Spring Boot programmatically?

**Answer:**
Inject `Environment`.
`env.getProperty("app.name")`.
Or `env.resolvePlaceholders("${app.name}")`.

---

### Question 318: How do you externalize secrets in Spring Boot in a secure way?

**Answer:**
(Duplicate of Q124).

---

### Question 319: How do you define immutable configuration beans?

**Answer:**
Use `@ConfigurationProperties` with `@ConstructorBinding` (Spring Boot 2.x) or just Record classes / final fields with 1 constructor (Spring Boot 3.x).
No Setters.
```java
@ConfigurationProperties("app")
public record AppConfig(String name, int timeout) {}
```

---

### Question 320: How can you log all loaded configuration properties during startup?

**Answer:**
(See Q298).
Or loop through `(AbstractEnvironment) env).getPropertySources()`.

---
