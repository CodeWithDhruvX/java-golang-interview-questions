# Testing and CI/CD - Interview Questions and Answers

## 1. What is the difference between Unit Testing and Integration Testing in Spring Boot?
**Answer:**
- **Unit Testing:** Tests an individual "unit" of code (usually a single class or method) in isolation from its dependencies. Dependencies (like database repositories or other services) are mocked. It is extremely fast and pinpointed. A failure usually means a logic error in that specific method. Tools: JUnit, Mockito.
- **Integration Testing:** Tests how different parts of the application work *together*. It tests the integration between the application code and external components like the database, external APIs, or message brokers. These tests involve booting up the Spring Context (partially or fully). They are slower but ensure the system as a whole functions correctly. Tools: `@SpringBootTest`, Testcontainers, `MockMvc` (for API endpoint testing).

## 2. Explain JUnit 5 and common Assert methods.
**Answer:**
JUnit 5 (Jupiter API) is the de-facto standard testing framework for Java. It uses annotations to identify test methods.

**Common Annotations:**
- `@Test`: Marks a method as a test case.
- `@BeforeEach` / `@AfterEach`: Executed before/after every single test method in the class (useful for resetting mocks or setting up test data).
- `@BeforeAll` / `@AfterAll`: Executed once before/after all tests in the class (must be static; useful for heavy resource initialization).

**Common Assertions (from `org.junit.jupiter.api.Assertions`):**
Assertions verify that the actual outcome matches the expected outcome.
- `assertEquals(expected, actual)`: Checks if two values are equal.
- `assertTrue(condition)` / `assertFalse(condition)`: Checks boolean conditions.
- `assertNull(object)` / `assertNotNull(object)`: Checks for nullability.
- `assertThrows(ExpectedException.class, () -> executableCode())`: Verifies that a specific piece of code throws the expected exception. It returns the thrown exception, allowing you to further assert its message.

## 3. What is Mockito, and how is it used for unit testing a Spring Service?
**Answer:**
**Mockito** is a mocking framework that allows you to create dummy objects (mocks) of interfaces or classes. This allows you to test a class in complete isolation by defining the behavior of its dependencies.

**Usage in a Spring Boot Service Test:**
1. **`@ExtendWith(MockitoExtension.class)`:** Annotate the test class to initialize mocks automatically.
2. **`@Mock`:** Create a mock instance of the dependency (e.g., `UserRepository`). The real database is completely ignored.
3. **`@InjectMocks`:** Create an instance of the class under test (e.g., `UserService`) and inject the `@Mock`s into it automatically via constructor injection.
4. **Behavior Stubbing (The "When/Then" pattern):** Before calling the method under test, tell the mock how to behave when a specific method is called.
    ```java
    User mockUser = new User("John");
    // When the repository's findById is called with an argument of 1L, return an Optional of mockUser.
    when(userRepository.findById(1L)).thenReturn(Optional.of(mockUser));
    ```
5. **Execution:** Call the method on the `@InjectMocks` instance (`userService`).
6. **Assertion:** Assert the result (JUnit) and verify interactions (Mockito).
    ```java
    // Verify that the repository's findById method was actually called exactly once with argument 1L.
    verify(userRepository, times(1)).findById(1L);
    ```

## 4. How do you implement Integration Tests for REST APIs (including Reactive CRUD)?
**Answer:**
To test an entire API flow (from the HTTP request reaching the Controller, down to the Service, Repository, and back), you use Spring's integration testing features.

**1. Traditional MVC Integration Test (`@WebMvcTest` + `@MockBean`):**
- You annotate the test class with `@WebMvcTest(UserController.class)`. This starts a "sliced" Spring context that only loads web-layer components (Controllers, Security, Filters) but ignores Services and Repositories.
- You use `@MockBean` to mock the Service layer. This adds the mock directly to the Spring Context.
- You use `MockMvc` to perform an HTTP request (e.g., `mockMvc.perform(get("/api/users/1"))`) and chain `.andExpect(status().isOk())` and `.andExpect(jsonPath("$.name").value("John"))` to verify the JSON response structure.

**2. Full End-to-End Test (`@SpringBootTest` + `TestRestTemplate` / `WebTestClient`):**
- You annotate the test class with `@SpringBootTest(webEnvironment = WebEnvironment.RANDOM_PORT)`. This starts the *entire* application context, including embedded Tomcat and database connections (often an embedded H2 database or a Testcontainer like PostgreSQL).
- **Reactive APIs (WebFlux):** For testing reactive CRUD APIs built with Spring WebFlux, you use `WebTestClient`. It is a non-blocking, reactive client specifically designed for testing APIs.
    ```java
    webTestClient.get().uri("/api/reactive-users/1")
        .exchange() // Executes the request
        .expectStatus().isOk()
        .expectBody() // Parses the reactive JSON stream
        .jsonPath("$.name").isEqualTo("John");
    ```

## 5. What are CI/CD Pipelines, and why are they necessary?
**Answer:**
**CI/CD (Continuous Integration / Continuous Deployment or Delivery)** automation bridges the gap between development and operation activities.

- **Continuous Integration (CI):** The practice of automating the building, testing, and merging of code changes into a central repository multiple times a day. If a developer breaks the unit tests, the CI build fails, and the code isn't merged until fixed.
- **Continuous Deployment (CD):** The automated process of releasing the validated code from the repository directly to staging or production environments.

**Necessity:** Manual deployments are slow, error-prone, and painful ("Deployment Fridays"). CI/CD enables fast, reliable, reproducible, and frequent releases (sometimes hundreds of times a day in microservices architectures).

## 6. How do you implement CI/CD using AWS CodePipeline and CodeDeploy for a Spring Boot app?
**Answer:**
AWS provides a suite of managed services to automate the release process.

**1. AWS CodeCommit (or GitHub/Bitbucket):** The source code repository where developers push their code.
**2. AWS CodeBuild (The CI Phase):**
- It compiles the source code, runs JUnit tests, and produces deployment artifacts (like a massive JAR file or a Docker image).
- **`buildspec.yml`:** A YAML file placed in the root of your source code. It defines the build phases (install, pre_build, build, post_build).
    ```yaml
    version: 0.2
    phases:
      build:
        commands:
          - mvn clean package -DskipTests=false
      post_build:
        commands:
          - echo Build completed on `date`
    artifacts:
      files:
        - target/my-spring-app.jar
      discard-paths: yes
    ```
**3. AWS CodeDeploy (The CD Phase):**
- It automates the deployment of the artifact produced by CodeBuild to compute services like Amazon EC2, AWS Fargate, or **AWS Elastic Beanstalk**.
- You define an `appspec.yml` file to manage the deployment lifecycle hooks (e.g., run a script to stop the old Java process, install new dependencies, start the new Java process, validate the service).

**4. AWS CodePipeline (The Orchestrator):**
- A service that models and visualizes the entire release pipeline. You configure it to watch your source code repository. When a push happens to the `main` branch, CodePipeline automatically triggers CodeBuild, takes the resulting artifact, and passes it to CodeDeploy or Elastic Beanstalk, achieving full automation from commit to production.
