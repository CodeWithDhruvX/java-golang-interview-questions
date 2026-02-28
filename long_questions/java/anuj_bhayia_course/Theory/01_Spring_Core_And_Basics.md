# Spring Core and Basics - Interview Questions and Answers

## 1. What is the Spring Framework and what are its core features?
**Answer:**
Spring is an open-source, lightweight, Java-based framework that provides comprehensive infrastructure support for developing robust Java applications. It allows developers to focus on application-level business logic rather than dealing with the complexities of enterprise Java.

**Core Features:**
1. **Inversion of Control (IoC):** The core container manages the complete lifecycle of objects, from creation to destruction. Instead of objects creating their dependencies, they declare them, and the Spring IoC container injects them. This promotes loose coupling.
2. **Dependency Injection (DI):** It is a specific implementation of IoC where the framework injects the dependencies into objects. This can be done via constructor injection, setter injection, or field injection (using `@Autowired`).
3. **Aspect-Oriented Programming (AOP):** Spring supports separating cross-cutting concerns (like logging, security, and transaction management) from business logic using aspects.
4. **Spring MVC:** A powerful Model-View-Controller framework for building web applications and RESTful APIs.
5. **Transaction Management:** Spring provides a consistent programming model for managing transactions (both programmatic and declarative) across various APIs (JDBC, Hibernate, JPA, JDO).
6. **Data Access/Integration:** Spring simplifies database access by providing templates (e.g., `JdbcTemplate`, `JpaTemplate`) that handle boilerplate code, resource management, and exception translation.

## 2. Explain the concept of Inversion of Control (IoC) and Dependency Injection (DI) in Spring.
**Answer:**
**Inversion of Control (IoC):**
IoC is a design principle where the control of object creation, configuration, and lifecycle is handed over to a container or framework rather than the objects themselves. In traditional programming, an object creates its dependencies using the `new` keyword. In IoC, the Spring Container is responsible for instantiating objects (beans), wiring them together, and managing their lifecycles.

**Dependency Injection (DI):**
DI is the design pattern used to implement IoC. Instead of a class explicitly fetching or creating its dependencies, the dependencies are injected into the class by the Spring container at runtime.

**Types of Dependency Injection in Spring:**
1. **Constructor Injection:** Dependencies are provided through the class constructor. This is the recommended approach for mandatory dependencies because it ensures the object is in a valid state upon creation and makes the class easier to test.
2. **Setter Injection:** Dependencies are provided through setter methods. This is useful for optional dependencies or when dependencies need to be changed after initialization.
3. **Field Injection:** Dependencies are injected directly into fields using the `@Autowired` annotation. While convenient, it is generally discouraged because it makes the class tightly coupled to the Spring container and harder to unit test without reflection.

## 3. What is a Spring Bean and how is its lifecycle managed?
**Answer:**
A **Spring Bean** is a Java object that is instantiated, assembled, and managed by the Spring IoC container. Beans are the backbone of a Spring application and are typically defined via XML configuration or Java annotations (e.g., `@Component`, `@Service`, `@Repository`, `@Controller`, `@Bean`).

**Spring Bean Lifecycle:**
The lifecycle of a Spring Bean is managed entirely by the IoC container. The key phases are:
1. **Instantiation:** The Spring container instantiates the bean object (usually via its constructor).
2. **Populate Properties (Dependency Injection):** The container injects the required dependencies and sets the bean properties.
3. **BeanNameAware / BeanFactoryAware / ApplicationContextAware:** If the bean implements any of these `Aware` interfaces, the container sets the respective properties (e.g., setting the bean name or application context).
4. **Pre-Initialization (BeanPostProcessor):** The container calls the `postProcessBeforeInitialization()` method of any registered `BeanPostProcessor`s. `@PostConstruct` annotated methods are called here.
5. **Initialization:** If the bean implements `InitializingBean`, its `afterPropertiesSet()` method is called. If a custom `init-method` is declared in the configuration, it is invoked.
6. **Post-Initialization (BeanPostProcessor):** The container calls the `postProcessAfterInitialization()` method of any registered `BeanPostProcessor`s. This is often where AOP proxies are created.
7. **Ready for Use:** The bean is now fully initialized and ready to be used by the application.
8. **Destruction:** When the application context is closed, the container manages destruction. If the bean implements `DisposableBean`, its `destroy()` method is called. `@PreDestroy` annotated methods and custom `destroy-method`s are also invoked.

## 4. What are the different Scopes of a Spring Bean?
**Answer:**
Spring provides several bean scopes, allowing developers to define the lifecycle and visibility of bean instances:
1. **Singleton (Default):** The Spring container creates exactly one instance of the bean per Spring IoC container. The same instance is returned for every request. It is thread-safe only if the bean is stateless.
2. **Prototype:** A new instance of the bean is created each time it is requested from the container. The container does not manage the full lifecycle of prototype beans; it instantiates, configures, and hands them over, but does not track them for destruction.
3. **Request:** A single instance is created per HTTP request. Only valid in web-aware Spring ApplicationContexts.
4. **Session:** A single instance is created per HTTP session. Only valid in web-aware Spring ApplicationContexts.
5. **Application:** A single instance is created for the lifecycle of a `ServletContext`. Only valid in web-aware Spring ApplicationContexts.
6. **WebSocket:** A single instance is created for the lifecycle of a WebSocket.

## 5. What is the difference between `@Component`, `@Service`, `@Repository`, and `@Controller` annotations?
**Answer:**
These are stereotype annotations used to mark classes as Spring-managed components. They enable auto-detection during component scanning, but each serves a semantic purpose:

- **`@Component`:** The most generic stereotype annotation. It marks a Java class as a bean so that component scanning can find it and add it to the application context.
- **`@Service`:** A specialization of `@Component`. It is used to mark a class that holds the business logic of the application. Currently, it behaves exactly like `@Component` but adds semantic meaning to the class for developers.
- **`@Repository`:** A specialization of `@Component` used for Data Access Objects (DAOs). It provides an additional feature: it catches persistence-specific exceptions (like `SQLException`) and rethrows them as Spring's unified unchecked data access exceptions (`DataAccessException`). This handles the exception translation mechanism.
- **`@Controller`:** A specialization of `@Component` used in Spring MVC to mark a class as a web controller that handles HTTP requests. It is often used with `@RequestMapping` to map URLs to handler methods. (`@RestController` is a convenience annotation that combines `@Controller` and `@ResponseBody`).

## 6. How does Spring Auto-Configuration work in Spring Boot?
**Answer:**
**Auto-Configuration** is a key feature of Spring Boot that attempts to automatically configure the Spring application based on the jar dependencies present on the classpath. It aims to eliminate boilerplate configuration.

**How it works:**
1. When a Spring Boot application starts, the `@EnableAutoConfiguration` (typically inherited from `@SpringBootApplication`) annotation is processed.
2. Spring Boot looks for the `META-INF/spring.factories` file (or `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports` in newer versions) in all auto-configure jars on the classpath.
3. These files contain lists of Auto-Configuration classes (e.g., `DataSourceAutoConfiguration`, `WebMvcAutoConfiguration`).
4. Spring Boot evaluates **Conditional Annotations** placed on these configuration classes to decide whether to activate them. Key conditional annotations include:
    - `@ConditionalOnClass`: Activates the configuration only if a specific class is present on the classpath (e.g., configure a database pool only if HikariCP is present).
    - `@ConditionalOnMissingBean`: Activates the configuration only if a bean of a specific type does NOT already exist in the container. This allows developers to easily override auto-configuration by providing their own custom beans.
    - `@ConditionalOnProperty`: Activates the configuration only if a certain application property is set (e.g., `spring.datasource.url`).

This mechanism allows Spring Boot to provide sensible defaults that back off gracefully when the user defines custom configurations.

## 7. Explain Dependency Injection using Constructor vs. Setter vs. Field Injection. Which is preferred and why?
**Answer:**
1. **Field Injection (`@Autowired` on a field):**
   - **Pros:** Very concise, easy to read and write.
   - **Cons:** Tightly couples the class to the Spring container. The class cannot be easily instantiated outside of the container (e.g., for unit testing) without reflection or a mocking framework. Hides structural dependencies, making it easy to violate the Single Responsibility Principle by packing too many dependencies into one class. Cannot be used for `final` (immutable) fields.
2. **Setter Injection (`@Autowired` on a setter method):**
   - **Pros:** Useful for optional dependencies or properties that can be changed after object initialization. Can resolve circular dependencies.
   - **Cons:** Objects can be created in an incomplete state if the setter is not called. Doesn't support immutability (`final` fields).
3. **Constructor Injection (`@Autowired` on a constructor, or implicit in newer Spring versions if there's only one constructor):**
   - **Pros:** **This is the highly recommended approach.** It ensures that the object is in a fully initialized, valid state upon creation. Supports immutability (dependencies can be marked `final`). Makes dependencies explicit and enforces structural requirements. Makes unit testing very easy, as dependencies can be mocked and passed directly to the constructor in a test environment without needing Spring.
   - **Cons:** Can lead to constructors with too many parameters if the class violates the Single Responsibility Principle (though this is often a code smell pointing to a needed refactor rather than a flaw of constructor injection).

## 8. What is the difference between BeanFactory and ApplicationContext?
**Answer:**
Both `BeanFactory` and `ApplicationContext` are interfaces that act as the IoC container. However, `ApplicationContext` is a sub-interface of `BeanFactory` and provides advanced features.

**BeanFactory:**
- It is the simplest container representing basic IoC and DI features.
- It uses **lazy initialization** (beans are instantiated only when they are requested via `getBean()`).
- Does not support annotations based dependency injection (like @Autowired) by default without explicit registration of PostProcessors.
- Used in resource-constrained applications (e.g., mobile devices) where memory consumption is a strict concern.

**ApplicationContext:**
- It extends `BeanFactory` and adds enterprise-specific features.
- It uses **eager initialization** for singletons by default (beans are instantiated when the context starts up).
- Supports AOP integration, message resource handling (i18n), event publication (`ApplicationEventPublisher`), and application-layer specific contexts (like `WebApplicationContext`).
- Supports an annotation-based approach seamlessly.
- **Preferred choice** for almost all modern Spring applications.

## 9. What is Java-based configuration in Spring? How is it implemented?
**Answer:**
Java-based configuration allows defining Spring beans and their configurations using Java classes instead of XML files. This approach is type-safe, refactoring-friendly, and often easier to read.

It is implemented using primarily the `@Configuration` and `@Bean` annotations.

- **`@Configuration`:** Indicates that a class declares one or more `@Bean` methods and may be processed by the Spring container to generate bean definitions and service requests for those beans at runtime.
- **`@Bean`:** Indicates that a method produces a bean to be managed by the Spring container. The method name becomes the bean name by default, and the return value is registered as the bean instance.

**Example:**
```java
@Configuration
public class AppConfig {

    @Bean
    public MyService myService() {
        // Dependencies can be explicitly passed or resolved by calling other @Bean methods
        return new MyServiceImpl(myRepository());
    }

    @Bean
    public MyRepository myRepository() {
        return new MyRepositoryImpl();
    }
}
```

## 10. How does Spring resolve Circular Dependencies, and how can they be avoided?
**Answer:**
**Circular Dependency Definition:** This occurs when Bean A depends on Bean B, and Bean B depends on Bean A. If using constructor injection, the application will fail to start with a `BeanCurrentlyInCreationException` because neither bean can be fully instantiated before the other.

**How Spring handles it:**
Spring can automatically resolve circular dependencies if they involve **setter injection or field injection** for singleton beans. It does this by exposing a partially constructed object reference (an "early reference") before the bean is fully initialized and its dependencies injected.

**How to avoid/fix them:**
1. **Redesign (Best Approach):** A circular dependency often indicates a design flaw. The classes might be too tightly coupled and the responsibilities need to be separated or extracted into a third component.
2. **Use `@Lazy`:** Annotating one of the injected dependencies with `@Lazy`. This tells Spring to inject a proxy instead of the actual bean. The actual bean will only be created when one of its methods is called for the first time, breaking the startup cycle.
## 11. What is Maven, and what are its core Build Lifecycle goals?
**Answer:**
**Maven** is a popular build automation, dependency management, and project management tool primarily used for Java projects. It relies on a Project Object Model (`pom.xml`) file to define the project structure, dependencies, and build plugins.

**Core Build Lifecycle Goals (Phases):**
Maven operates on a standardized build lifecycle consisting of distinct phases. Executing a later phase automatically executes all preceding phases.
1. **`validate`:** Validates the project is correct and all necessary information is available.
2. **`compile`:** Compiles the source code of the project.
3. **`test`:** Tests the compiled source code using a suitable unit testing framework (e.g., JUnit). These tests should not require the code be packaged or deployed.
4. **`package`:** Takes the compiled code and packages it into its distributable format, such as a JAR or WAR file.
5. **`verify`:** Runs any checks on results of integration tests to ensure quality criteria are met.
6. **`install`:** Installs the package into the *local repository* (`~/.m2`), making it available as a dependency for other projects running locally on your machine.
7. **`deploy`:** Done in the build environment, copies the final package to a *remote repository* (like Nexus or Artifactory) for sharing with other developers and projects.
*Common Commands:* `mvn clean package`, `mvn clean install` (`clean` is a separate lifecycle that deletes the `target/` compilation directory before rebuilding).

## 12. Explain the internal startup flow of a Spring Boot application (`SpringApplication.run()`).
**Answer:**
When you execute the `main` method in a Spring Boot application, `SpringApplication.run(MyApp.class, args)` triggers a sequence of core events to bootstrap the application:

**The Startup Sequence:**
1. **Instantiate `SpringApplication`:** It determines the application type (Reactive, Servlet, or Non-Web).
2. **Load `ApplicationContextInitializer` & `ApplicationListener`:** It finds and loads these from `META-INF/spring.factories` to hook into the lifecycle events.
3. **Trigger `ApplicationStartingEvent`:** Listeners are notified that the app is starting.
4. **Prepare Environment:** It prepares the `Environment` (loading OS environment variables, application arguments, and properties from `application.properties`/`application.yml`).
5. **Trigger `ApplicationEnvironmentPreparedEvent`.**
6. **Print Banner:** The Spring ASCII art banner is printed to the console.
7. **Create `ApplicationContext`:** It creates the appropriate IoC container (e.g., `AnnotationConfigServletWebServerApplicationContext` for standard web apps).
8. **Register Singleton Beans:** All properties are bound, and core internal beans are registered.
9. **Load Sources (`@SpringBootApplication`):** The primary source class is loaded. The core `@EnableAutoConfiguration` annotation kicks in.
10. **Refresh Context (The Heavy Lifting):** The `refresh()` method of the `ApplicationContext` is called. The `BeanFactory` destroys old singletons, creates new singletons, performs Component Scanning (`@ComponentScan`), reads all `@Configuration` classes, processes Auto-Configuration conditions, and injects all `@Autowired` dependencies.
11. **Start Embedded Web Server:** If it's a web application, the embedded Tomcat/Jetty/Undertow server starts during the refresh phase.
12. **Trigger `ApplicationStartedEvent`:** The context is refreshed, but explicitly defined `CommandLineRunner` or `ApplicationRunner` beans haven't run yet.
13. **Execute Runners:** Any beans implementing `CommandLineRunner` or `ApplicationRunner` are executed. (Useful for custom database seeding or startup scripts).
14. **Trigger `ApplicationReadyEvent`:** The application has fully started and is ready to accept incoming HTTP traffic.
