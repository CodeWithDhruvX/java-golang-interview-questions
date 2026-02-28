# Spring Boot & Microservices Practical Solutions

This folder contains the accepted solutions and pseudo-code implementations for the practical coding questions found in the parent directory.

## How to Set Up the Practice Project

To practice these questions locally and verify your code, the best approach is to generate a single Spring Boot application and add features as you go.

### 1. Generate the Project using Spring Initializr
1. Go to **[start.spring.io](https://start.spring.io/)**
2. **Project:** Maven
3. **Language:** Java
4. **Spring Boot:** 3.2.x (or latest stable)
5. **Project Metadata:**
   - Group: `com.interview`
   - Artifact: `practical`
   - Packaging: `Jar`
   - Java: `17` or `21`
6. **Add Dependencies:**
   - **Web:** Spring Web
   - **Data:** Spring Data JPA, MySQL Driver
   - **Validation:** Validation (Hibernate Validator)
   - **Lombok:** Lombok (to reduce boilerplate)
   - **Security:** Spring Security
   - **Actuator:** Spring Boot Actuator
   - **Messaging:** Spring for Apache Kafka
   - **Cache:** Spring Cache Abstraction, Spring Data Redis
   - **AI:** Spring AI (OpenAI for paid, Ollama for free local models)

### 2. Import into your IDE
1. Click **GENERATE** to download the `.zip` file.
2. Unzip it and open the folder in **IntelliJ IDEA**, **Eclipse**, or **VS Code**.
3. Allow the IDE to download the Maven dependencies.

### 3. Application Properties Setup
In your `src/main/resources/application.properties` (or `application.yml`), add the following basic configurations to get started using the MySQL database:

```properties
# Server
server.port=8080

# MySQL Database Configuration
spring.datasource.url=jdbc:mysql://localhost:3306/interviewdb?useSSL=false&serverTimezone=UTC&allowPublicKeyRetrieval=true
spring.datasource.driverClassName=com.mysql.cj.jdbc.Driver
spring.datasource.username=root
spring.datasource.password=password
spring.jpa.database-platform=org.hibernate.dialect.MySQLDialect

# Hibernate
spring.jpa.hibernate.ddl-auto=update
spring.jpa.show-sql=true
```

### 4. How to Use These Solutions
- Treat the problems in the parent folder as an actual interview assessment.
- Attempt to write the classes, annotations, and logic *before* looking at the solution files.
- Compare your implementation with the solutions provided here. Note that in Spring Boot, there are often multiple ways to achieve the same result; the solutions provided here follow modern enterprise best practices.
