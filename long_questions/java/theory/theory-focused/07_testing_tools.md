# Testing & Tools - Interview Answers

> 🎯 **Focus:** These answers demonstrate you are a professional who ships high-quality, verified code.

### 1. JUnit 4 vs JUnit 5?
"JUnit 5 is the modern standard. It’s modular and supports Java 8 features better.

Key changes:
Annotations changed: `@Before` became `@BeforeEach`, and `@BeforeClass` became `@BeforeAll`.
It supports **Nested Tests** (`@Nested`), which allows me to group related test cases and create nice hierarchical reports.
It also has built-in support for **Parameter tests** (`@ParameterizedTest`), so I can run the same test logic with 10 different inputs cleanly."

---

### 2. What is Mockito? Why use it?
"Mockito is a mocking framework. It lets us create dummy implementations of dependencies.

If I'm testing a `UserService`, it dependencies on a `UserRepository`. I don't want my unit test to actually hit the real database—that's slow and fragile.
So I use Mockito to say: `when(repo.findById(1)).thenReturn(user)`.

This isolates the `UserService` logic. I am testing only the service code, not the database connection. That's the definition of a Unit Test."

---

### 3. @Mock vs @InjectMocks?
"These are Mockito annotations used to bootstrap a test class.

**@Mock** creates a fake, dummy object (like the Repository).
**@InjectMocks** creates the real object we are testing (like the Service) and injects the mocks into it.

So, standard setup:
`@Mock UserRepository repo;`
`@InjectMocks UserService service;`
When the test starts, Mockito puts the mock repo inside the real service so I can run my methods."

---

### 4. Maven lifecycle phases?
"Maven follows a standard lifecycle sequence.

The main ones I use are:
**clean**: Removes the `target` folder.
**compile**: Compiles source code to `.class` files.
**test**: Runs unit tests.
**package**: Bundles the code into a JAR or WAR.
**install**: Puts that JAR into my local `.m2` repository so other local projects can use it.
**deploy**: Pushes the artifact to a remote repository like Nexus/Artifactory."

---

### 5. Git Merge vs Rebase?
"Both integrate changes from one branch to another, but the history looks different.

**Merge** creates a 'merge commit.' It preserves the history exactly as it happened, including the branch structure. It’s non-destructive.
**Rebase** rewrites history. It takes my commits and 'plays' them on top of the latest master. It creates a linear, clean history without extra merge commits.

I prefer **Merge** (Squash Merge) for Pull Requests to keep the main branch clean, but I use **Rebase** locally to keep my feature branch up to date with master."

---

### 6. What is TDD (Test Driven Development)?
"It’s a development cycle: **Red, Green, Refactor**.

1. **Red**: Write a failing test for a feature that doesn't exist yet.
2. **Green**: Write just enough code to make that test pass.
3. **Refactor**: Clean up the code while keeping the test passing.

Honestly, I don't practice strict TDD 100% of the time. But for complex logic—like a calculation engine or a parser—writing the test first is a huge help to clarify the requirements before I get lost in the code."

---

### 7. Unit Test vs Integration Test?
"**Unit Tests** test a single class or method in isolation. They mock out all external dependencies (DB, APIs). They are incredibly fast (milliseconds).

**Integration Tests** verify that components work *together*. In Spring Boot, I use `@SpringBootTest`. It loads the actual Spring Context and uses a real (or in-memory H2) database. They are slower but prove that my Service and Repository actually talk to each other correctly."

---

### 8. Dependency Injection vs Dependency Management?
"**Dependency Injection** is a code pattern (Java/Spring) where objects are given their dependencies at runtime.

**Dependency Management** is a build tool concern (Maven/Gradle). It’s about resolving valid versions of libraries. For example, making sure my `spring-web` jar is compatible with my `jackson` jar. Maven handles the management; Spring handles the injection."

---

### 9. What is SonarQube?
"It’s a tool for static code analysis. We integrate it into our CI/CD pipeline.

It scans the code for bugs, vulnerabilities, and code smells—like duplicate code, cognitive complexity, or lack of test coverage. It acts as a quality gate. If my code coverage drops below 80% or if I have critical vulnerabilities, the build fails."
