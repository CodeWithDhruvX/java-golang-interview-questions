## 🔹 Section 1: Core Concepts & Annotations (301-310)

### Question 301: What is the use of `@ConditionalOnMissingBean` in Spring Boot?

**Answer:**
It tells Spring to register a bean **only if** a bean of that type (or name) does not already exist in the context.
This is the core implementation of "overridable defaults".
If the user defines their own `DataSource` bean, the default one annotated with `@ConditionalOnMissingBean` is skipped.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `@ConditionalOnMissingBean` in Spring Boot?
**Your Response:** "`@ConditionalOnMissingBean` tells Spring to register a bean only if a bean of that type doesn't already exist in the context. This is the core implementation of 'overridable defaults' - Spring Boot's auto-configuration uses this annotation extensively. For example, if I define my own `DataSource` bean, the default auto-configured one annotated with `@ConditionalOnMissingBean` is skipped. This allows me to override Spring Boot's defaults simply by defining my own beans, giving me complete control while still providing sensible defaults out of the box."

---

### Question 302: What does `@SpringBootConfiguration` do under the hood?

**Answer:**
It is simply a specialized `@Configuration` annotation.
It is used to mark the main configuration class.
Mainly used for tests (`@SpringBootTest`) to automatically find the primary configuration of the application.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does `@SpringBootConfiguration` do under the hood?
**Your Response:** "`@SpringBootConfiguration` is essentially a specialized `@Configuration` annotation. It's used to mark the main configuration class and is primarily used by tests like `@SpringBootTest` to automatically find the primary configuration of the application. While it functions like `@Configuration`, the specific annotation allows Spring Boot to identify the main configuration class more reliably, especially in complex applications with multiple configuration classes. It's a small but important detail for Spring Boot's auto-configuration and testing mechanisms."

---

### Question 303: How does Spring Boot handle circular dependency issues?

**Answer:**
By default in recent versions (2.6+), Circular References are **disabled** and throw an exception.
To fix:
1.  Use `@Lazy` on one of the constructor fields.
2.  Enable legacy circular handling: `spring.main.allow-circular-references=true`.
3.  Refactor architecture.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot handle circular dependency issues?
**Your Response:** "In recent Spring Boot versions (2.6+), circular references are disabled by default and throw an exception to prevent potential issues. To fix circular dependencies, I have several options: use `@Lazy` on one of the constructor fields to break the cycle, enable legacy circular handling with `spring.main.allow-circular-references=true`, or ideally refactor the architecture to eliminate the circular dependency. I prefer refactoring as it's the cleanest solution, but `@Lazy` is a practical workaround when refactoring isn't immediately possible."

---

### Question 304: What is `@EnableAutoConfiguration`, and how does it work with `@SpringBootApplication`?

**Answer:**
(See Q3/Q28). It triggers the auto-configuration logic.
Included inside `@SpringBootApplication`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `@EnableAutoConfiguration`, and how does it work with `@SpringBootApplication`?
**Your Response:** "`@EnableAutoConfiguration` triggers Spring Boot's auto-configuration logic, which automatically configures beans based on the classpath and other conditions. It's included inside the `@SpringBootApplication` annotation, so I don't need to add it separately when using the main application annotation. This annotation is what enables Spring Boot's convention-over-configuration approach - it scans for auto-configuration classes and applies them as needed, setting up everything from datasources to web servers automatically."

---

### Question 305: How do you use `@Lazy` initialization in Spring Boot?

**Answer:**
(See Q34/Q104).
Annotate a Bean with `@Lazy`.
Or set `spring.main.lazy-initialization=true` to make ALL beans lazy.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `@Lazy` initialization in Spring Boot?
**Your Response:** "I use `@Lazy` initialization in two ways: I can annotate individual beans with `@Lazy` to defer their creation until they're actually needed, or I can set `spring.main.lazy-initialization=true` to make all beans lazy by default. Lazy initialization improves startup time but can delay discovery of configuration errors. I typically use it for development to speed up restarts, but disable it in production to catch issues early. For specific beans that are expensive to create and might not always be used, `@Lazy` is perfect for optimizing resource usage."

---

### Question 306: What is the difference between `@Import` and `@ComponentScan`?

**Answer:**
(Duplicate of Q117).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@Import` and `@ComponentScan`?
**Your Response:** "`@ComponentScan` automatically scans packages for stereotype annotations like `@Component`, `@Service`, etc. - it's implicit and broad. `@Import` explicitly registers specific configuration classes or components regardless of their package location. I use `@ComponentScan` for general component discovery and `@Import` when I need precise control over which configuration classes to load. `@Import` is particularly useful for library integration or when I want to modularize my configuration by explicitly importing specific configuration classes."

---

### Question 307: How can you override auto-configured beans in Spring Boot?

**Answer:**
Just define a Bean of the same type in your configuration.
Because Auto-configurations usually use `@ConditionalOnMissingBean`, your bean (registered first) will cause the auto-config bean to back off.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can you override auto-configured beans in Spring Boot?
**Your Response:** "I override auto-configured beans simply by defining a bean of the same type in my configuration. Spring Boot's auto-configurations typically use `@ConditionalOnMissingBean`, so when I define my own bean, it gets registered first and causes the auto-configured bean to back off. This elegant mechanism allows me to customize any part of Spring Boot's configuration just by providing my own implementation. The ordering works because my configuration is processed before auto-configuration, giving me complete control while still benefiting from Spring Boot's defaults for everything I don't customize."

---

### Question 308: What is the purpose of `@DependsOn` annotation?

**Answer:**
Forces specific beans to be initialized **before** the annotated bean.
Useful when there is an implicit dependency not visible via Constructor/Autowired (e.g., Static initialization, DB Schema creation).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `@DependsOn` annotation?
**Your Response:** "`@DependsOn` forces specific beans to be initialized before the annotated bean. I use this when there's an implicit dependency that isn't visible through constructor injection or `@Autowired`. For example, if a bean relies on static initialization or database schema creation in another bean, I use `@DependsOn` to ensure the dependency is ready. This annotation is particularly useful for complex initialization scenarios where Spring can't automatically determine the correct bean initialization order."

---

### Question 309: What are factory beans in Spring Boot and when to use them?

**Answer:**
Beans that implement `FactoryBean<T>`.
Used to create complex objects (like `Proxy` objects or objects from 3rd party non-Spring libraries).
Spring injects the result of `getObject()` rather than the FactoryBean itself.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are factory beans in Spring Boot and when to use them?
**Your Response:** "Factory beans are beans that implement `FactoryBean<T>` and are used to create complex objects. I use them when I need to create objects that require special construction logic, like proxy objects, objects from third-party libraries, or objects that need complex configuration. Spring injects the result of `getObject()` rather than the FactoryBean itself, so other components don't need to know about the factory pattern. This is particularly useful for integrating with legacy systems or creating objects that can't be created through standard dependency injection."

---

### Question 310: What are some common bean scope types in Spring Boot?

**Answer:**
(Duplicate of Q32). Singleton, Prototype, Request, Session.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are some common bean scope types in Spring Boot?
**Your Response:** "Spring Boot supports several bean scopes. Singleton is the default - one instance per application context. Prototype creates a new instance each time it's injected. Request scope creates one instance per HTTP request in web applications. Session scope creates one instance per user session. I choose the scope based on the bean's purpose - Singleton for stateless services, Prototype for stateful objects that need fresh instances, Request for request-specific data, and Session for user-specific data that needs to persist across requests."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you bind lists and maps from YAML to `@ConfigurationProperties`?
**Your Response:** "I bind lists and maps from YAML using standard collection types. For lists, I define a `List<String> servers` in my properties class and YAML with array syntax. For maps, I use `Map<String, Integer> users` and YAML with key-value syntax. Spring Boot automatically maps the YAML structure to these Java collections. This makes complex configuration structures clean and type-safe. I can nest lists and maps arbitrarily deep, allowing me to represent sophisticated configuration hierarchies as natural Java object structures."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the order of precedence in Spring Boot property resolution?
**Your Response:** "Spring Boot has a well-defined property precedence hierarchy from highest to lowest priority. Command line arguments take highest priority, followed by JVM system properties, then environment variables. After that come profile-specific configuration files, then standard application properties, and finally defaults in `@ConfigurationProperties` classes. This order means I can override any configuration at runtime using command line arguments, which is perfect for containerized deployments. The hierarchy ensures that production settings can override development defaults without code changes."

---

### Question 313: How do you inject system environment variables into your Spring Boot config?

**Answer:**
Reference them in `application.properties`:
`app.db.password=${DB_PASSWORD_ENV_VAR}`.
Or map directly if names match (Relaxed Binding).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you inject system environment variables into your Spring Boot config?
**Your Response:** "I inject system environment variables in two ways. I can reference them directly in `application.properties` using placeholders like `app.db.password=${DB_PASSWORD_ENV_VAR}`. Alternatively, if the environment variable name matches the property name (considering relaxed binding), Spring Boot maps them automatically. The relaxed binding means `DB_PASSWORD` maps to `db.password`. This approach allows me to externalize configuration to the environment without changing application code, which is essential for containerized deployments."

---

### Question 314: How do you provide a fallback configuration in Spring Boot?

**Answer:**
Use the `:` default operator in SpEL or placeholders.
`@Value("${app.timeout:5000}")`.
If `app.timeout` is missing, use `5000`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you provide a fallback configuration in Spring Boot?
**Your Response:** "I provide fallback configuration using the colon default operator in SpEL or property placeholders. For example, `@Value('${app.timeout:5000}')` uses 5000 as the default value if `app.timeout` is not defined. I can also use this in YAML with `timeout: 5000` when the property might be missing. This approach ensures that my application has sensible defaults even when specific configuration properties are not provided, making the application more robust and reducing the need for mandatory configuration."

---

### Question 315: How can you encrypt configuration properties in Spring Boot?

**Answer:**
(See Q124). Using Jasypt.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can you encrypt configuration properties in Spring Boot?
**Your Response:** "I encrypt configuration properties using Jasypt. I add the jasypt-spring-boot-starter dependency and then encrypt sensitive values with `ENC(encrypted_string)` in my configuration files. Jasypt automatically decrypts these values at runtime using a password or key. This approach allows me to commit encrypted configuration to version control while keeping the actual values secure. I typically provide the decryption password through environment variables or system properties in production, ensuring that sensitive data is never stored in plain text."

---

### Question 316: What is the role of `PropertySourcesPlaceholderConfigurer`?

**Answer:**
Bean Post Processor that resolves `${...}` placeholders in bean definitions against the current Spring Environment properties.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `PropertySourcesPlaceholderConfigurer`?
**Your Response:** "`PropertySourcesPlaceholderConfigurer` is a bean post-processor that resolves `${...}` placeholders in bean definitions against the current Spring Environment properties. It's the mechanism that allows me to use placeholders like `${app.name}` in my configuration and have them replaced with actual property values. While modern Spring Boot applications typically use `@Value` and `@ConfigurationProperties`, this placeholder configurator is the underlying mechanism that makes property substitution work throughout the Spring context."

---

### Question 317: How do you resolve placeholders in Spring Boot programmatically?

**Answer:**
Inject `Environment`.
`env.getProperty("app.name")`.
Or `env.resolvePlaceholders("${app.name}")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you resolve placeholders in Spring Boot programmatically?
**Your Response:** "I resolve placeholders programmatically by injecting the `Environment` object and using its methods. I can use `env.getProperty('app.name')` to get a specific property value, or `env.resolvePlaceholders('${app.name}')` to resolve placeholder strings. This approach is useful when I need to access configuration dynamically or when I'm building custom configuration processing logic. The Environment object gives me access to all property sources and their values, allowing me to work with configuration programmatically."

---

### Question 318: How do you externalize secrets in Spring Boot in a secure way?

**Answer:**
(Duplicate of Q124).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you externalize secrets in Spring Boot in a secure way?
**Your Response:** "I externalize secrets securely using several approaches. Environment variables are great for containerized deployments. Spring Cloud Vault can fetch secrets from a vault service at runtime. Kubernetes secrets can be mounted as files or environment variables. For applications that need encrypted configuration in files, I use Jasypt to encrypt values in configuration files. The key principle is never committing secrets to version control and always loading them from external sources at runtime. This ensures that sensitive information remains secure even if the code is compromised."

---

### Question 319: How do you define immutable configuration beans?

**Answer:**
Use `@ConfigurationProperties` with `@ConstructorBinding` (Spring Boot 2.x) or just Record classes / final fields with 1 constructor (Spring Boot 3.x).
No Setters.
```java
@ConfigurationProperties("app")
public record AppConfig(String name, int timeout) {}
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define immutable configuration beans?
**Your Response:** "I define immutable configuration beans using modern Java features. In Spring Boot 3.x, I can simply use record classes with final fields and a single constructor, annotated with `@ConfigurationProperties`. In Spring Boot 2.x, I use `@ConstructorBinding` with final fields and no setters. Records are particularly elegant - `public record AppConfig(String name, int timeout) {}` creates an immutable configuration bean. This approach ensures configuration values can't be changed after initialization, making the configuration thread-safe and predictable."

---

### Question 320: How can you log all loaded configuration properties during startup?

**Answer:**
(See Q298).
Or loop through `(AbstractEnvironment) env).getPropertySources()`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can you log all loaded configuration properties during startup?
**Your Response:** "I can log all loaded configuration properties by iterating through the property sources in the Environment. I cast the Environment to `AbstractEnvironment` and access `getPropertySources()` to see all loaded properties. I can also use system-wide default properties to log configuration during startup. This is particularly useful for debugging configuration issues or when I need to verify that the correct configuration is being loaded. However, I'm careful not to log sensitive information like passwords or API keys."

---
