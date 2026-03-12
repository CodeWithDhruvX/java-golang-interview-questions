## 🔹 Section 2: Annotations & Dependency Injection (21–40)

### Question 21: What is the use of `@Component`, `@Service`, and `@Repository`?

**Answer:**
They are **Stereotype Annotations** marking classes as Spring Beans.
*   **`@Component`**: Generic stereotype for any Spring-managed component.
*   **`@Repository`**: For Data Access Layer (DAO). Adds automatic Persistence Exception Translation (converts SQL errors to Spring DataAccessExceptions).
*   **`@Service`**: For Service Layer (Business Logic).
*   **`@Controller`**: For Web Layer.
Functionally, they all register beans, but the specialization adds semantic meaning and specific capabilities (like exception translation).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `@Component`, `@Service`, and `@Repository`?
**Your Response:** "These are stereotype annotations that mark classes as Spring beans, but they have different semantic meanings. `@Component` is the generic stereotype for any Spring-managed component. `@Repository` is specifically for the data access layer - it not only marks the class as a bean but also adds automatic exception translation, converting SQL exceptions to Spring's DataAccessExceptions. `@Service` is for the service layer containing business logic, and `@Controller` is for the web layer. While they all functionally register beans, the specialization makes my code more readable and enables framework-specific features like exception translation."

---

### Question 22: How does Spring Boot handle dependency injection?

**Answer:**
Using the **IoC (Inversion of Control) Container**.
1.  **Scan:** It scans for classes annotated with `@Component`/`@Service`/etc.
2.  **Instantiate:** It creates instances of these classes (Beans).
3.  **Inject:** It looks for dependencies (Constructor Args or `@Autowired` fields) and injects the matching Bean instance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot handle dependency injection?
**Your Response:** "Spring Boot uses the Inversion of Control container to handle dependency injection automatically. First, it scans for classes annotated with `@Component`, `@Service`, `@Repository`, and other stereotypes. Then it creates instances of these classes as beans. Finally, it looks for dependencies - either through constructor arguments or `@Autowired` fields - and injects the appropriate bean instances. This means I don't have to manually create objects or wire them together; Spring handles all the wiring automatically based on the annotations I use."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@Autowired` and `@Qualifier`?
**Your Response:** "`@Autowired` and `@Qualifier` work together to solve different problems. `@Autowired` injects a bean by type, which works perfectly when there's only one bean of that type. But if I have multiple implementations of the same interface, Spring throws an exception because it doesn't know which one to inject. That's where `@Qualifier` comes in - I use it along with `@Autowired` to specify the exact bean name I want. For example, if I have multiple payment services, I can use `@Autowired` with `@Qualifier('paymentServicePayPal')` to inject specifically the PayPal implementation."

---

### Question 24: What is the difference between `@Value` and `@ConfigurationProperties`?

**Answer:**
*   **`@Value("${property.name}")`**: Injects a single property value into a field. Good for one-off values.
*   **`@ConfigurationProperties(prefix="app")`**: Binds a group of properties to a POJO automatically using setters. Type-safe and supports Lists/Maps.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@Value` and `@ConfigurationProperties`?
**Your Response:** "`@Value` and `@ConfigurationProperties` serve different purposes for configuration management. `@Value` is perfect for injecting single property values into fields - I use it for one-off configurations like `@Value('${api.key}')`. But `@ConfigurationProperties` is more powerful - it binds a group of related properties to a POJO automatically using setters, making it type-safe and supporting complex structures like lists and maps. For example, I can create a class with `@ConfigurationProperties(prefix='app')` and it will automatically map all properties starting with 'app' to the class fields. This is much cleaner for managing related configuration as a cohesive unit."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create custom annotations in Spring Boot?
**Your Response:** "I can create custom annotations by defining an interface with the `@interface` keyword and then meta-annotating it with Spring annotations. For example, I can create a `@CustomService` annotation and annotate it with `@Service`. This makes my custom annotation behave exactly like `@Service` - Spring will recognize it and treat classes annotated with `@CustomService` as service beans. This is useful for creating domain-specific annotations that make my code more readable and expressive while still leveraging Spring's component scanning and dependency injection capabilities."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `@Bean` and how is it different from `@Component`?
**Your Response:** "`@Bean` and `@Component` both register beans with Spring, but they're used in different scenarios. `@Component` is a class-level annotation that I use on classes I write myself - Spring automatically detects these through component scanning. `@Bean` is a method-level annotation that I use inside `@Configuration` classes, primarily for third-party libraries where I can't modify the source code to add `@Component`. For example, I use `@Bean` to configure a `RestTemplate` or a `DataSource` from a library I didn't write. The method return value becomes the bean instance that Spring manages."

---

### Question 27: What is the `@Configuration` annotation?

**Answer:**
Indicates that a class declares one or more `@Bean` methods.
It is a source of bean definitions.
Spring uses CGLIB proxying on `@Configuration` classes to ensure that calling `@Bean` methods internally returns the *same* singleton instance (inter-bean dependency).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the `@Configuration` annotation?
**Your Response:** "The `@Configuration` annotation indicates that a class declares one or more `@Bean` methods, making it a source of bean definitions. What's special about `@Configuration` classes is that Spring uses CGLIB proxying to ensure that when I call one `@Bean` method from within the same configuration class, I get the same singleton instance rather than creating a new one. This is important for managing dependencies between beans. For example, if one bean needs another bean defined in the same class, the proxying ensures I get the Spring-managed instance rather than a new object."

---

### Question 28: Explain the use of `@EnableAutoConfiguration`.

**Answer:**
It tells Spring Boot to "guess" how you want to configure Spring, based on the jar dependencies that you have added.
For example, if `tomcat-embed-core.jar` is on your classpath, it creates a `TomcatServletWebServerFactory`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the use of `@EnableAutoConfiguration`.
**Your Response:** "`@EnableAutoConfiguration` is the annotation that tells Spring Boot to automatically configure the application context based on the jar dependencies on the classpath. It's the magic behind Spring Boot's convention over configuration approach. For instance, if Spring Boot detects `tomcat-embed-core.jar` on my classpath, it automatically creates a `TomcatServletWebServerFactory` bean to configure the embedded Tomcat server. Similarly, if it sees a database driver, it configures a datasource. This intelligent configuration happens without me writing any explicit configuration code, though I can override it when needed."

---

### Question 29: What is the use of `@Profile` in Spring Boot?

**Answer:**
Restricts a Bean or Configuration to be loaded only when a specific Environment Profile is active.
```java
@Component
@Profile("dev")
public class DevDataSource implements DataSource { ... }
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `@Profile` in Spring Boot?
**Your Response:** "`@Profile` allows me to conditionally load beans or configurations based on the active environment profile. I can annotate a component or configuration class with `@Profile('dev')` and it will only be loaded when the 'dev' profile is active. This is perfect for environment-specific configurations - for example, I can have a `DevDataSource` bean that only loads in development and a `ProdDataSource` that loads in production. I activate profiles through properties or command line arguments, and Spring automatically includes only the beans matching the active profile, keeping my environment configurations cleanly separated."

---

### Question 30: How to handle circular dependencies in Spring Boot?

**Answer:**
Occurs when Bean A needs Bean B, and Bean B needs Bean A.
**Solutions:**
1.  **Refactor:** Break the cycle (best).
2.  **`@Lazy`:** Inject one of them lazily (`@Autowired @Lazy`).
3.  **Setter/Field Injection:** Instead of Constructor Injection (allows partially constructed beans).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle circular dependencies in Spring Boot?
**Your Response:** "Circular dependencies occur when Bean A needs Bean B, and Bean B needs Bean A. The best solution is to refactor the code to break the cycle, usually by extracting common functionality into a third bean. If refactoring isn't possible, I have a few options. I can use `@Lazy` annotation on one of the dependencies, which delays its initialization until it's actually needed, breaking the circular reference. Alternatively, I can switch from constructor injection to setter or field injection, which allows Spring to create partially constructed beans and then wire them together. However, I prefer refactoring as it's the cleanest long-term solution."

---

### Question 31: Can we use constructor injection in Spring Boot?

**Answer:**
Yes, and it is the **Recommended** way.
It ensures immutability and makes testing easier (you can pass mocks in the constructor).
If a class has only one constructor, `@Autowired` is optional (Spring 4.3+).

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can we use constructor injection in Spring Boot?
**Your Response:** "Yes, and it's actually the recommended approach for dependency injection. Constructor injection ensures immutability since dependencies are final, and it makes testing much easier because I can pass mocks directly through the constructor. Starting with Spring 4.3, if a class has only one constructor, I don't even need the `@Autowired` annotation - Spring automatically uses it for dependency injection. This makes my code cleaner and more testable. Constructor injection also fails fast if dependencies are missing, which helps catch configuration issues early during application startup."

---

### Question 32: What is a `@Scope` annotation and its types?

**Answer:**
Defines the lifecycle/visibility of a Bean.
*   **Singleton (Default):** One instance per container.
*   **Prototype:** New instance every request.
*   **Request:** One instance per HTTP request (Web).
*   **Session:** One instance per HTTP session (Web).
*   **Application:** One instance per ServletContext.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a `@Scope` annotation and its types?
**Your Response:** "The `@Scope` annotation defines the lifecycle and visibility of a Spring bean. By default, beans are singleton, meaning one instance per Spring container. But I can change this behavior. For web applications, I can use `@Scope('request')` to create a new instance for each HTTP request, or `@Scope('session')` for one instance per HTTP session. There's also `@Scope('prototype')` which creates a new instance every time the bean is requested, and `@Scope('application')` which creates one instance per ServletContext. Each scope serves different use cases - singleton for stateless services, prototype for stateful objects, and request/session for web-specific scenarios."

---

### Question 33: What are singleton and prototype beans?

**Answer:**
*   **Singleton:** Stateless (usually). Shared across the entire app. Created on startup.
*   **Prototype:** Stateful. Created each time `getBean()` is called or injected. Beware: Spring does *not* manage the destruction lifecycle of Prototypes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are singleton and prototype beans?
**Your Response:** "Singleton and prototype are the two most common bean scopes in Spring. Singleton beans, which is the default, create exactly one instance per Spring container - they're typically stateless and shared across the entire application, created when the container starts up. Prototype beans create a new instance every time they're requested or injected, making them suitable for stateful objects. However, I need to be careful with prototype beans because Spring doesn't manage their destruction lifecycle - once Spring creates and hands out a prototype bean, it doesn't track it afterward. So I'm responsible for cleanup if needed."

---

### Question 34: What is lazy initialization and how do you enable it?

**Answer:**
By default, Spring creates all Singleton beans eagerly on startup.
**Lazy:** Stalls creation until the bean is requested.
*   `@Lazy` on specific bean.
*   `spring.main.lazy-initialization=true` property (Global).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is lazy initialization and how do you enable it?
**Your Response:** "By default, Spring creates all singleton beans eagerly when the application starts up. Lazy initialization delays bean creation until the bean is actually requested. I can enable this in two ways: by adding `@Lazy` to specific beans that I want to initialize lazily, or globally by setting `spring.main.lazy-initialization=true`. Lazy initialization can speed up application startup, especially for applications with many beans that aren't used immediately. However, I need to be careful because it can hide configuration errors - problems with lazy beans only surface when they're first accessed, not during startup."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to define conditional beans? (`@ConditionalOnProperty`)
**Your Response:** "Conditional beans allow me to load beans only when certain conditions are met. `@ConditionalOnProperty` is particularly useful for feature toggling. I can annotate a bean method with `@ConditionalOnProperty(name='feature.x.enabled', havingValue='true')` and Spring will only create that bean if the specified property is set to true. This is perfect for optional features - if the feature is disabled, the bean is completely skipped, saving resources. I can also use other conditional annotations like `@ConditionalOnClass` (load only if a class is present) or `@ConditionalOnMissingBean` (load only if another bean isn't defined)."

---

### Question 36: What is a proxy bean and when is it created?

**Answer:**
A wrapper object created by Spring (using CGLIB or JDK Dynamic Proxy) to add behavior (AOP) around the actual bean.
Created when you use features like `@Transactional`, `@Async`, `@Cacheable`. The caller calls the proxy, which handles the transaction, then calls the real method.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a proxy bean and when is it created?
**Your Response:** "A proxy bean is a wrapper object that Spring creates to add cross-cutting concerns like transactions, caching, or security around my actual beans. Spring uses either CGLIB or JDK Dynamic Proxy to create these wrappers. When I use annotations like `@Transactional`, `@Async`, or `@Cacheable`, Spring automatically creates a proxy. When I call a method on the bean, I'm actually calling the proxy, which handles the cross-cutting concern first - like starting a transaction - and then delegates to the real method. This happens transparently, so my code stays clean while Spring handles the infrastructure concerns."

---

### Question 37: What are factory methods in Spring beans?

**Answer:**
Methods annotated with `@Bean`. They act as a factory for creating object instances.
Or static factory methods specified in XML (legacy).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are factory methods in Spring beans?
**Your Response:** "Factory methods in Spring are methods annotated with `@Bean` that act as factories for creating object instances. Instead of Spring creating objects through reflection, it calls my factory method to get the bean instance. This gives me complete control over how objects are created - I can perform custom initialization, apply configuration, or call specific constructors. Factory methods are especially useful for third-party classes where I can't add annotations, or when I need to configure objects in specific ways before they become Spring beans. In legacy XML configuration, this was done through factory-method attributes, but now `@Bean` is the modern approach."

---

### Question 38: What is the difference between `ApplicationContext` and `BeanFactory`?

**Answer:**
*   **BeanFactory:** The root interface. Provides basic DI support. Lazy loading.
*   **ApplicationContext:** Extends BeanFactory. Adds enterprise features: AOP, Event Publication, I18N, Actuator support. Eager loading. Spring Boot uses `ApplicationContext`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `ApplicationContext` and `BeanFactory`?
**Your Response:** "`BeanFactory` is the root interface of Spring's container system, providing basic dependency injection support with lazy loading. `ApplicationContext` extends `BeanFactory` and adds enterprise features like AOP support, event publication, internationalization, and integration with Spring Boot's Actuator. While `BeanFactory` loads beans lazily, `ApplicationContext` loads them eagerly on startup. In modern Spring Boot applications, I always work with `ApplicationContext` because it provides the full feature set I need for enterprise applications. The additional capabilities make it worth the slightly longer startup time."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot manage bean lifecycle?
**Your Response:** "Spring Boot manages bean lifecycle through a well-defined sequence of steps. First, Spring instantiates the bean. Then it populates properties through dependency injection. Next come the awareness callbacks like `BeanNameAware` and `BeanFactoryAware`. After that, `BeanPostProcessor` runs before initialization, followed by initialization callbacks like `@PostConstruct` or `InitializingBean`. Then another `BeanPostProcessor` runs after initialization - this is where proxy creation happens for AOP. Finally, the bean is ready to use. During shutdown, Spring calls `@PreDestroy` or `DisposableBean` methods for cleanup. This lifecycle management ensures beans are properly initialized and destroyed."

---

### Question 40: What is dependency injection and inversion of control in Spring Boot?

**Answer:**
*   **IoC:** The principle where the control of object creation is transferred from the programmer ("new X()") to a container (Framework).
*   **DI:** The design pattern used to implement IoC. The container "injects" the dependencies an object needs (Service) into it (Controller).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is dependency injection and inversion of control in Spring Boot?
**Your Response:** "Inversion of Control is a fundamental principle where control of object creation shifts from me manually creating objects with 'new' to the framework managing object creation. Dependency Injection is the design pattern that implements IoC. Instead of my controller creating its own service dependencies, the Spring container injects those dependencies into the controller. This decouples my code - the controller doesn't need to know how to create or find its dependencies. It just declares what it needs, and Spring provides it. This makes my code more testable, flexible, and easier to maintain because I can swap implementations without changing the consuming classes."

---
