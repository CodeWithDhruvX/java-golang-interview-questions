## 🔹 Section 2: Annotations & Dependency Injection (21–40)

### Question 21: What is the use of `@Component`, `@Service`, and `@Repository`?

**Answer:**
They are **Stereotype Annotations** marking classes as Spring Beans.
*   **`@Component`**: Generic stereotype for any Spring-managed component.
*   **`@Repository`**: For Data Access Layer (DAO). Adds automatic Persistence Exception Translation (converts SQL errors to Spring DataAccessExceptions).
*   **`@Service`**: For Service Layer (Business Logic).
*   **`@Controller`**: For Web Layer.
Functionally, they all register beans, but the specialization adds semantic meaning and specific capabilities (like exception translation).

---

### Question 22: How does Spring Boot handle dependency injection?

**Answer:**
Using the **IoC (Inversion of Control) Container**.
1.  **Scan:** It scans for classes annotated with `@Component`/`@Service`/etc.
2.  **Instantiate:** It creates instances of these classes (Beans).
3.  **Inject:** It looks for dependencies (Constructor Args or `@Autowired` fields) and injects the matching Bean instance.

---

### Question 23: What is the difference between `@Autowired` and `@Qualifier`?

**Answer:**
*   **`@Autowired`**: Injects a bean by **Type**. If multiple beans of the same type exist, it throws `NoUniqueBeanDefinitionException`.
*   **`@Qualifier`**: Used *along with* `@Autowired` to resolve ambiguity by specifying the **Name** of the bean.

```java
@Autowired
@Qualifier("paymentServicePayPal")
private PaymentService paymentService;
```

---

### Question 24: What is the difference between `@Value` and `@ConfigurationProperties`?

**Answer:**
*   **`@Value("${property.name}")`**: Injects a single property value into a field. Good for one-off values.
*   **`@ConfigurationProperties(prefix="app")`**: Binds a group of properties to a POJO automatically using setters. Type-safe and supports Lists/Maps.

---

### Question 25: How do you create custom annotations in Spring Boot?

**Answer:**
Define an interface with `@interface`. You can meta-annotate it with Spring annotations.

```java
@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Service // Meta-annotation
public @interface CustomService { }
```
Now `@CustomService` acts just like `@Service`.

---

### Question 26: What is the use of `@Bean` and how is it different from `@Component`?

**Answer:**
*   **`@Component`**: Class-level. Used for classes *you* write (Auto-detection).
*   **`@Bean`**: Method-level. Used in `@Configuration` classes. Used for third-party libraries where you cannot modify the source code to add `@Component`.

```java
@Bean
public RestTemplate restTemplate() {
    return new RestTemplate();
}
```

---

### Question 27: What is the `@Configuration` annotation?

**Answer:**
Indicates that a class declares one or more `@Bean` methods.
It is a source of bean definitions.
Spring uses CGLIB proxying on `@Configuration` classes to ensure that calling `@Bean` methods internally returns the *same* singleton instance (inter-bean dependency).

---

### Question 28: Explain the use of `@EnableAutoConfiguration`.

**Answer:**
It tells Spring Boot to "guess" how you want to configure Spring, based on the jar dependencies that you have added.
For example, if `tomcat-embed-core.jar` is on your classpath, it creates a `TomcatServletWebServerFactory`.

---

### Question 29: What is the use of `@Profile` in Spring Boot?

**Answer:**
Restricts a Bean or Configuration to be loaded only when a specific Environment Profile is active.
```java
@Component
@Profile("dev")
public class DevDataSource implements DataSource { ... }
```

---

### Question 30: How to handle circular dependencies in Spring Boot?

**Answer:**
Occurs when Bean A needs Bean B, and Bean B needs Bean A.
**Solutions:**
1.  **Refactor:** Break the cycle (best).
2.  **`@Lazy`:** Inject one of them lazily (`@Autowired @Lazy`).
3.  **Setter/Field Injection:** Instead of Constructor Injection (allows partially constructed beans).

---

### Question 31: Can we use constructor injection in Spring Boot?

**Answer:**
Yes, and it is the **Recommended** way.
It ensures immutability and makes testing easier (you can pass mocks in the constructor).
If a class has only one constructor, `@Autowired` is optional (Spring 4.3+).

---

### Question 32: What is a `@Scope` annotation and its types?

**Answer:**
Defines the lifecycle/visibility of a Bean.
*   **Singleton (Default):** One instance per container.
*   **Prototype:** New instance every request.
*   **Request:** One instance per HTTP request (Web).
*   **Session:** One instance per HTTP session (Web).
*   **Application:** One instance per ServletContext.

---

### Question 33: What are singleton and prototype beans?

**Answer:**
*   **Singleton:** Stateless (usually). Shared across the entire app. Created on startup.
*   **Prototype:** Stateful. Created each time `getBean()` is called or injected. Beware: Spring does *not* manage the destruction lifecycle of Prototypes.

---

### Question 34: What is lazy initialization and how do you enable it?

**Answer:**
By default, Spring creates all Singleton beans eagerly on startup.
**Lazy:** Stalls creation until the bean is requested.
*   `@Lazy` on specific bean.
*   `spring.main.lazy-initialization=true` property (Global).

---

### Question 35: How to define conditional beans? (`@ConditionalOnProperty`)

**Answer:**
Load a bean only if a certain condition is met.
```java
@Bean
@ConditionalOnProperty(name = "feature.x.enabled", havingValue = "true")
public MyService myService() { ... }
```
If property is false, this bean is skipped.

---

### Question 36: What is a proxy bean and when is it created?

**Answer:**
A wrapper object created by Spring (using CGLIB or JDK Dynamic Proxy) to add behavior (AOP) around the actual bean.
Created when you use features like `@Transactional`, `@Async`, `@Cacheable`. The caller calls the proxy, which handles the transaction, then calls the real method.

---

### Question 37: What are factory methods in Spring beans?

**Answer:**
Methods annotated with `@Bean`. They act as a factory for creating object instances.
Or static factory methods specified in XML (legacy).

---

### Question 38: What is the difference between `ApplicationContext` and `BeanFactory`?

**Answer:**
*   **BeanFactory:** The root interface. Provides basic DI support. Lazy loading.
*   **ApplicationContext:** Extends BeanFactory. Adds enterprise features: AOP, Event Publication, I18N, Actuator support. Eager loading. Spring Boot uses `ApplicationContext`.

---

### Question 39: How does Spring Boot manage bean lifecycle?

**Answer:**
1.  Instantiate.
2.  Populate Properties (DI).
3.  `BeanNameAware` / `BeanFactoryAware` callbacks.
4.  `BeanPostProcessor` (BeforeInitialization).
5.  `@PostConstruct` / `InitializingBean.afterPropertiesSet()`.
6.  `BeanPostProcessor` (AfterInitialization - Proxy creation here).
7.  Ready to use.
8.  `@PreDestroy` / `DisposableBean` (Shutdown).

---

### Question 40: What is dependency injection and inversion of control in Spring Boot?

**Answer:**
*   **IoC:** The principle where the control of object creation is transferred from the programmer ("new X()") to a container (Framework).
*   **DI:** The design pattern used to implement IoC. The container "injects" the dependencies an object needs (Service) into it (Controller).

---
