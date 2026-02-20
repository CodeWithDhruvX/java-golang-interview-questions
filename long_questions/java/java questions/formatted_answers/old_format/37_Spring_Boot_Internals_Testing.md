# 37. Spring Boot Internals & Testing Strategies

**Q: Spring Boot Startup Process (Internals)**
> "When you call `SpringApplication.run()`:
> 1.  It starts a **StopWatch** to track time.
> 2.  It prepares the **Environment** (reads properties, profiles).
> 3.  It prints the **Banner**.
> 4.  It starts the **IoC Container** (ApplicationContext).
> 5.  It triggers **Auto-Configuration** (scanning classpath).
> 6.  It calls any `CommandLineRunners`.
> It’s a carefully choreographed sequence of events."

**Indepth:**
> **Ready Event**: `ApplicationReadyEvent`. The last step. If you want to engage the user (send an email "Server Started"), listen for this event. It means everything is fully up and running.


---

**Q: Auto-Configuration Mechanism**
> "How does Spring know to configure H2?
> It uses `@Conditional` annotations.
>
> It looks at `spring.factories` (or `imports` file in Boot 3).
> It finds `H2AutoConfiguration`.
> It checks: `@ConditionalOnClass(H2.class)`.
> If the H2 JAR is on the classpath, the condition passes, and the bean is created. If not, it's ignored."

**Indepth:**
> **Boot 3**: The "Starter" pattern relies on the `META-INF` files. In Boot 2.7+, `spring.factories` is deprecated for auto-config; use `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports`.


---

**Q: Disabling Specific Auto-Configuration**
> "Sometimes Spring tries to be too smart. For example, configuring a DataSource when you don't even want one yet.
>
> You can exclude it:
> `@SpringBootApplication(exclude = { DataSourceAutoConfiguration.class })`
>
> This tells Boot: 'I know you see the driver, but don't touch it. I'll handle it manually'."

**Indepth:**
> **Properties**: You can also exclude via properties: `spring.autoconfigure.exclude=org.spring...DataSourceAutoConfiguration`. This is useful for "testing" profiles or when debugging weird conflicts.


---

**Q: Custom Banner**
> "It's a fun feature.
> Just drop a `banner.txt` file in `src/main/resources`.
> You can use ASCII art generators. You can even use placeholders like `${spring-boot.version}` inside the text file to print the version number dynamically on startup."

**Indepth:**
> **CI/CD**: You can turn it off (`spring.main.banner-mode=off`) to speed up startup logs slightly and reduce noise in CI (Continuous Integration) pipelines.


---

**Q: Logging Configuration (Logback)**
> "Spring Boot uses **Logback** by default.
> You don't need a config file for simple changes.
> `logging.level.org.springframework=DEBUG` in `application.properties` is enough.
>
> For complex stuff (like rotating files daily, producing JSON logs for Splunk), create a `logback-spring.xml` in resources. Spring will pick it up automatically."

**Indepth:**
> **Profiles**: `<springProfile name="prod">`. Inside `logback-spring.xml`, you can nest config blocks. "If profile is dev, print to Console. If profile is prod, print to File and Logstash appender."


---

**Q: REST Controller vs Controller**
> "**@Controller** is for standard Spring MVC. It typically returns a String (view name like `index.jsp`).
>
> "**@RestController** is a convenience annotation. It combines `@Controller` and `@ResponseBody`.
> It tells Spring: 'Whatever I return from methods, write it directly to the HTTP Response body as JSON'. You don't need to annotate every method with `@ResponseBody`."

**Indepth:**
> **Under the Hood**: `@ResponseBody` works by using `HttpMessageConverters`. If the return type is String, it uses `StringHttpMessageConverter`. If Object, `MappingJackson2HttpMessageConverter`.


---

**Q: Content Negotiation**
> "One URL, multiple formats.
> The Client asks: `Accept: application/json` -> Server returns JSON.
> The Client asks: `Accept: application/xml` -> Server returns XML.
>
> You don't change code. You just add the `jackson-dataformat-xml` dependency, and Spring automatically supports both formats based on the header."

**Indepth:**
> **Parameters**: `format=json`. You can also configure Spring to look at a query param `?format=xml` instead of the Header. This is easier for testing in browsers.


---

**Q: Testing: @MockBean**
> "When unit testing a Service, you don't want to hit the real Database.
>
> `@MockBean` is used in Spring tests to replace a real bean with a Mockito mock.
>
> ```java
> @MockBean
> private UserRepository userRepo;
>
> // In test
> given(userRepo.findById(1)).willReturn(Optional.of(mockUser));
> ```
> This bypasses the actual database logic entirely."

**Indepth:**
> **Context Caching**: Spring Test caches the context. If you use `@MockBean`, it modifies the context (swaps a bean). This "dirties" the context, forcing Spring to reload a fresh context for the next test class. Too many `@MockBean`s slow down your test suite massively.


---

**Q: Testing: @DataJpaTest**
> "This is a 'Slice Test'.
> It creates a Spring Context, but **only** loads JPA components (Entities, Repositories). It does **not** load Controllers or Services.
>
> It also automatically configures an in-memory DB (H2).
> It's perfect for testing if your custom JPQL queries actually work without starting the whole heavy application."

**Indepth:**
> **Rollback**: By default, every test method is `@Transactional` and rolls back at the end. Your data is clean for the next test. If you want to see data in the DB for debugging, use `@Commit`.


---

**Q: Profile-Specific Testing**
> "You should never run tests against your Production database config.
>
> Annotate your test class:
> `@ActiveProfiles("test")`
>
> Then create `application-test.yml` with H2 or TestContainer settings. Spring will switch to this config during the test run, keeping your data safe."

**Indepth:**
> **Inline**: `@TestPropertySource(properties = "app.feature=false")` has higher precedence than `application-test.yml`. It's great for overriding one specific flag for one specific test case.

