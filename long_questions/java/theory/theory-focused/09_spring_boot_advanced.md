# Spring Boot Advanced - Interview Answers

> 🎯 **Focus:** These answers show you know how to run Spring Boot in production, not just on localhost.

### 1. What is Spring Boot Actuator?
"It’s a sub-project that gives you production-ready features to monitor your app.
It exposes endpoints like `/actuator/health` to check if the app is up, `/actuator/metrics` for memory/CPU usage, and `/actuator/env` to see properties.

In production, I always secure these endpoints using Spring Security so unauthorized users can't shutdown the app or see my config variables."

---

### 2. How do Profiles work (`@Profile`)?
"Profiles let us segregate configuration for different environments.
I usually have `application-dev.yml` for my local H2 database and `application-prod.yml` for the real MySQL database.

I activate them by setting `spring.profiles.active=prod` in an environment variable or command line argument when starting the JAR.
It ensures I never accidentally connect to the production database while testing locally."

---

### 3. How to dockerize a Spring Boot application?
"The standard way is to create a `Dockerfile`.
I start with a base image like `openjdk:17-alpine`.
Then I copy my compiled JAR file: `COPY target/myapp.jar app.jar`.
And finally, the entry point: `ENTRYPOINT ["java", "-jar", "app.jar"]`.

Recently, I've started using **Cloud Native Buildpacks** (via `mvn spring-boot:build-image`) because it builds an optimized Docker image without me needing to write a Dockerfile manually."

---

### 4. `CommandLineRunner` vs `ApplicationRunner`?
"Both are interfaces used to run code *once* right after the application starts—like for seeding a database or loading a cache.

The difference is how they receive arguments.
`CommandLineRunner` gets raw `String[] args`.
`ApplicationRunner` gets a parsed `ApplicationArguments` object, which makes it easier to handle flags like `--server.port=8080`.
I prefer `ApplicationRunner` purely for that convenience."

---

### 5. How does Spring Boot handle external configuration?
"It has a strict priority order for reading properties.
It checks:
1. Command Line Arguments (Highest priority)
2. Java System Properties (`-Dkey=value`)
3. OS Environment Variables
4. `application-prod.yml` (Profile specific)
5. `application.yml` (Default)

This hierarchy allows DevOps teams to override settings (like DB passwords) using Environment Variables in Kubernetes without rebuilding the JAR."

---

### 6. What is Spring Boot DevTools?
"It’s a developer productivity dependency.
Its main feature is **Automatic Restart**. When I change a class in my IDE and recompile, DevTools detects it and hot-reloads the application in a split second.
It also disables caching for Thymeleaf templates so I can see UI changes instantly.
It automatically disables itself when packaged as a pure JAR, so it's safe to leave in the `pom.xml`."

---

### 7. How do you handle Async methods (`@Async`)?
"By default, everything in Spring runs on the request thread.
If I annotate a method with `@Async` and enable it with `@EnableAsync`, Spring submits that method to a separate Thread Pool.

I use this for tasks like sending emails or generating reports—things that shouldn't block the user's response.
Note: It only works if you call the method from *outside* the class, because of how Spring proxies work."

---

### 8. What is `TestContainers`?
"It’s a library we use for Integration Testing.
Instead of mocking the database (which isn't 100% realistic) or using H2 (which behaves differently than Postgres), `TestContainers` spins up a *real* Postgres Docker container purely for the test execution.

It ensures our tests run against the exact same database engine we use in production, catching SQL syntax differences early."

---

### 9. How do you implement Scheduling (`@Scheduled`)?
"I use `@Scheduled` for cron jobs.
I just verify `@EnableScheduling` is on the main class, then add `@Scheduled(cron = '0 0 12 * * ?')` to a method.

However, by default, Spring uses a single thread for all schedulers. So if one job hangs, others wait. So I always configure a custom `TaskScheduler` pool to run them in parallel."

---

### 10. `application.properties` vs `application.yml`?
"Functionally, they are identical.
`properties` is flat: `spring.datasource.url=...`
`yml` is hierarchical:
```yaml
spring:
  datasource:
    url: ...
```
I prefer **YAML** because it's much more readable, especially when you have deeply nested configurations. But you have to be careful with indentation!"
