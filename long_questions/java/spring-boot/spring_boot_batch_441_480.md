## 🔹 Section 3: Testing in Spring Boot (441-460)

### Question 441: What is the difference between `@SpringBootTest` and `@WebMvcTest`?

**Answer:**
(See Q91/Q88). `SpringBootTest` is Full Context. `WebMvcTest` is Slice.

---

### Question 442: How do you test REST controllers in Spring Boot?

**Answer:**
(See Q90). Using `MockMvc`.

---

### Question 443: How do you mock services using `@MockBean`?

**Answer:**
(See Q89).

---

### Question 444: What is the use of `TestRestTemplate` in Spring Boot?

**Answer:**
(See Q59). Real HTTP Client for Integration tests.

---

### Question 445: How do you test JPA repositories effectively?

**Answer:**
(See Q80). `@DataJpaTest`.

---

### Question 446: How do you use in-memory databases for unit tests?

**Answer:**
(See Q80). H2 dependency + `@DataJpaTest`.

---

### Question 447: What is `@DataJpaTest` used for?

**Answer:**
(Duplicate of 445).

---

### Question 448: How do you perform integration testing with embedded containers?

**Answer:**
Using **TestContainers**.
(See Q202).

---

### Question 449: What is the purpose of `@TestConfiguration`?

**Answer:**
(See Q211). Register extra beans for tests.

---

### Question 450: How do you test exception scenarios in Spring Boot?

**Answer:**
(See Q206). `MockMvc` expectation of error message/status.

---

### Question 451: What is the difference between `@SpringBootTest` and `@WebMvcTest`?

**Answer:**
(Duplicate of 441).

---

### Question 452: How do you test REST controllers in Spring Boot?

**Answer:**
(Duplicate of 442).

---

### Question 453: How do you mock services using `@MockBean`?

**Answer:**
(Duplicate of 443).

---

### Question 454: What is the use of `TestRestTemplate` in Spring Boot?

**Answer:**
(Duplicate of 444).

---

### Question 455: How do you test JPA repositories effectively?

**Answer:**
(Duplicate of 445).

---

### Question 456: How do you use in-memory databases for unit tests?

**Answer:**
(Duplicate of 446).

---

### Question 457: What is `@DataJpaTest` used for?

**Answer:**
(Duplicate of 447).

---

### Question 458: How do you perform integration testing with embedded containers?

**Answer:**
(Duplicate of 448).

---

### Question 459: What is the purpose of `@TestConfiguration`?

**Answer:**
(Duplicate of 449).

---

### Question 460: How do you test exception scenarios in Spring Boot?

**Answer:**
(Duplicate of 450).

## 🔹 Section 4: Deployment, Profiles & Environment (461-480)

### Question 461: What are Spring Boot profiles and how do they work?

**Answer:**
(See Q20/Q29). Logical groups of configuration.

---

### Question 462: How do you activate multiple profiles simultaneously?

**Answer:**
`spring.profiles.active=dev,debug,cloud`.
Separated by comma.

---

### Question 463: How do you deploy a Spring Boot app as a WAR file?

**Answer:**
(See Q12). Extend `SpringBootServletInitializer`. Change packing to `war`.

---

### Question 464: What is the difference between embedded Tomcat and external Tomcat deployment?

**Answer:**
- **Embedded:** JAR, managed by Boot, isolated.
- **External:** WAR, managed by Ops, shared Tomcat instance (Standard Servlet Container).

---

### Question 465: How do you deploy Spring Boot apps on Heroku?

**Answer:**
Push code to Heroku Git or `heroku deploy:jar`.
Heroku detects `pom.xml`, builds app, and runs using `Procfile` (`web: java -jar ...`).

---

### Question 466: What is `spring.profiles.active` and where can you define it?

**Answer:**
Property to enable profiles.
Defined in `application.properties`, CLI Arg, or Env Var.

---

### Question 467: How do you implement zero-downtime deployment for Spring Boot apps?

**Answer:**
(See Q268). Rolling updates + Graceful Shutdown.

---

### Question 468: How do you enable graceful shutdown in Spring Boot?

**Answer:**
(See Q98). `server.shutdown=graceful`.

---

### Question 469: How do you deploy Spring Boot apps to AWS Lambda?

**Answer:**
Use **Spring Cloud Function**.
Adapter `SpringBootRequestHandler` or `FunctionInvoker`.
Wraps the function, providing cold-start optimization.

---

### Question 470: What is the role of `application.properties` vs `application.yml`?

**Answer:**
(See Q6). They serve same purpose. YAML is hierarchical.

---

### Question 471: What are Spring Boot profiles and how do they work?

**Answer:**
(Duplicate of 461).

---

### Question 472: How do you activate multiple profiles simultaneously?

**Answer:**
(Duplicate of 462).

---

### Question 473: How do you deploy a Spring Boot app as a WAR file?

**Answer:**
(Duplicate of 463).

---

### Question 474: What is the difference between embedded Tomcat and external Tomcat deployment?

**Answer:**
(Duplicate of 464).

---

### Question 475: How do you deploy Spring Boot apps on Heroku?

**Answer:**
(Duplicate of 465).

---

### Question 476: What is `spring.profiles.active` and where can you define it?

**Answer:**
(Duplicate of 466).

---

### Question 477: How do you implement zero-downtime deployment for Spring Boot apps?

**Answer:**
(Duplicate of 467).

---

### Question 478: How do you enable graceful shutdown in Spring Boot?

**Answer:**
(Duplicate of 468).

---

### Question 479: How do you deploy Spring Boot apps to AWS Lambda?

**Answer:**
(Duplicate of 469).

---

### Question 480: What is the role of `application.properties` vs `application.yml`?

**Answer:**
(Duplicate of 470).

---
