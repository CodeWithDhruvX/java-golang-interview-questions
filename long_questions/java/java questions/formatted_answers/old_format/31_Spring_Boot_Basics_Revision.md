# 31. Spring Boot (Basics & Testing Revision)

**Q: Spring Boot vs Spring**
> "**Spring** is the Engine. It provides Dependency Injection, MVC, Transaction Management, etc. But it requires a lot of setup (XML or huge Java Config classes).
>
> **Spring Boot** is the Car. It wraps the Engine with 'Auto Configuration', an Embedded Server, and 'Starters'. It lets you just turn the key (Run the main method) and drive, without assembling the parts yourself."

**Indepth:**
> **Convention over Configuration**: This is the core philosophy. Spring Boot assumes reasonable defaults (convention) so you don't have to configure them, unless you *want* to different (configuration).


---

**Q: @SpringBootApplication**
> "It's a convenience annotation that combines three others:
> 1.  `@Configuration`: Defines this as a source of bean definitions.
> 2.  `@EnableAutoConfiguration`: Tells Boot to start adding beans based on classpath settings.
> 3.  `@ComponentScan`: Tells Boot to look for other components (`@Service`, `@Controller`) in the current package and sub-packages."

**Indepth:**
> **Entry Point**: This annotation acts as the blueprint for the application. It triggers the scanning process that finds all your Beans, Controllers, and Services to wire them together in the ApplicationContext.


---

**Q: Auto Configuration**
> "It's Spring Boot's 'Opinionated' logic.
> At startup, Boot checks your Classpath.
> *   Do you have `h2.jar`? -> I'll create an in-memory DataSource.
> *   Do you have `spring-webmvc.jar`? -> I'll configure a DispatcherServlet and Tomcat.
>
> You can override any of these 'opinions' by defining your own bean."

**Indepth:**
> **Backing Off**: Conditional annotations like `@ConditionalOnMissingBean` allow you to "back off" the default. "If the user defined their own `DataSource`, don't create my default embedded one."


---

**Q: Spring Boot Starter**
> "A Starter is a 'Bill of Materials'. It's a single dependency in your `pom.xml` that brings in all the necessary jars for a feature.
>
> Instead of adding `spring-web`, `jackson`, `tomcat-embed`, and `validation-api` separately, you just add **`spring-boot-starter-web`**, and it pulls them all in with compatible versions."

**Indepth:**
> **Version Management**: Starters also act as a parent pom, managing transitive dependencies. You don't need to specify version numbers for individual libraries (like logging or jackson); the Starter ensures they are compatible.


---

**Q: Embedded Tomcat**
> "In the old days, you installed a separate Tomcat server, built a `.war` file, and deployed it.
>
> In Spring Boot, Tomcat is just a library (a JAR) *inside* your application.
> When you run your App, it starts Tomcat programmatically. This means your app is self-contained and portable."

**Indepth:**
> **Switching**: Boot also supports Jetty and Undertow. You can exclude `spring-boot-starter-tomcat` and add `spring-boot-starter-jetty` if you prefer a different servlet container engine.


---

**Q: JAR vs WAR**
> "**JAR (Java Archive)**: The default for Boot. It includes the embedded server. You run it with `java -jar app.jar`. Best for Microservices and Containers.
>
> "**WAR (Web Archive)**: Legacy format. Used if you *must* deploy to an external, existing Tomcat/Wildfly server. Boot supports it, but it's less common now."

**Indepth:**
> **Cloud Native**: JAR is the standard for Docker and Kubernetes. It aligns with the "12 Factor App" methodology where configuration and runtime are bundled together.


---

**Q: TestContainers**
> "Stop mocking your database in integration tests. And stop using H2 if you use Postgres in production.
>
> **TestContainers** spins up a *real* Docker container (e.g., a real Postgres DB) for your test, runs the test against it, and shuts it down.
> It ensures your code works with the actual database technology you use in Prod."

**Indepth:**
> **Transient**: The containers are ephemeral. They start fresh for the test suite and are destroyed afterwards. No more "dirty database" issues causing flaky tests.


---

**Q: @MockBean vs @SpyBean**
> "**@MockBean**: Replaces the real bean with a hollow shell (Mockito mock).
> *   `userService.getUser()` returns `null` unless you define `when(...).thenReturn(...)`.
> *   Use this to isolate the component you are testing.
>
> "**@SpyBean**: Wraps the *real* bean.
> *   Methods run normally, but you can verify them (`verify(bean).called()`) or stub specific methods.
> *   Use this when you want integration testing but need to spy on internal calls."

**Indepth:**
> **Mockito**: Both of these integrate natively with Mockito. `@MockBean` is essentially `Mockito.mock()`, and `@SpyBean` is `Mockito.spy()`, but they are automatically injected into the Spring ApplicationContext.


---

**Q: @WebMvcTest vs @SpringBootTest**
> "**@SpringBootTest**: Starts the **whole application context**. Connecting to DB, loading all services. Slow. Good for full Integration Tests.
>
> "**@WebMvcTest**: Slices the context. Only loads the Controller layer. It does **not** load `@Service` or `@Repository` beans.
> Fast. Use it for testing Unit Testing controllers (URL mapping, JSON serialization)."

**Indepth:**
> **Performance**: `@WebMvcTest` is dramatically faster than `@SpringBootTest` because it doesn't initialize the database or business layer. Use it for checking HTTP status codes and JSON formatting.


---

**Q: MockMvc**
> "This allows you to test your Controllers without starting a real HTTP Server.
> It simulates incoming HTTP requests.
>
> `mockMvc.perform(get("/api/users")).andExpect(status().isOk())`
>
> It tests the *web layer logic* (routing, headers, cookies) but calls the Java methods directly, skipping the network stack."

**Indepth:**
> **Integration**: MockMvc is usually used with `@WebMvcTest`. It allows checking the *content* of the response (body, headers) using a fluent assertion API.

