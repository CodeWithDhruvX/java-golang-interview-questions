# 🏗️ Architecture — Service-Based Companies Q&A

> **Level:** 🟢 Junior – 🟡 Mid
> **Asked at:** TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini, HCL, Tech Mahindra

---

## Q1. What is a 3-tier architecture? Explain with an example.

"A 3-tier architecture separates an application into **three logical layers**: Presentation (UI), Application (Business Logic), and Data (Database). Each tier has a specific role and communicates only with adjacent tiers.

Example: A Java Spring Boot banking application.
- **Presentation tier:** The Angular frontend that the customer uses — login page, account dashboard, transfer form.
- **Application tier:** The Spring Boot REST API — validates the transfer request, checks balance, applies business rules, authenticates the user.
- **Data tier:** The MySQL database — stores accounts, transactions, user credentials.

When the customer transfers ₹500: Angular sends POST `/transfer` → Spring Boot validates and processes → MySQL records the debit and credit. No layer is skipped — the frontend never directly queries the database."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys, Wipro (fresher/0-2 year rounds)

#### Indepth
Benefits of 3-tier:
- **Maintainability:** UI changes don't touch business logic. DB schema changes are hidden behind the API layer.
- **Security:** Database is not directly accessible from the internet — only via the application tier.
- **Scalability:** Each tier can scale independently — add more app servers behind a load balancer without adding DB servers.

In practice, a 3-tier Spring Boot app:
```
[Browser: Angular] → [Load Balancer: Nginx] → [App: Spring Boot x3] → [DB: MySQL Primary + Replica]
```

The presentation tier in modern apps is often served from a CDN, not the same server. The app tier is stateless (session in Redis), enabling horizontal scaling. The data tier is the primary scaling bottleneck — addressed with caching (Redis) and read replicas.

---

## Q2. What is RESTful web service and how is it different from SOAP?

"A RESTful web service is an API that follows REST architectural constraints — **stateless, uniform interface, resource-based URIs, standard HTTP verbs (GET, POST, PUT, DELETE)**. Responses are typically JSON.

SOAP (Simple Object Access Protocol) is an older XML-based protocol for web services. It uses a strict message format (XML envelope + header + body), has a formal contract (WSDL file), and supports advanced features like WS-Security and WS-ReliableMessaging.

Modern API development almost exclusively uses REST. SOAP is found in enterprise/banking systems from the 2000s that haven't been migrated, and in some government and BFSI integrations where its WS-Security (XML-level signing and encryption) is mandated."

#### 💬 **How to Explain in Interviews (Spoken Format)**

*"RESTful services are lightweight APIs that use standard HTTP methods like GET, POST, PUT, DELETE to work with resources. Think of it like browsing the web - you use URLs to access resources and HTTP verbs to tell the server what you want to do with them. The responses are usually in JSON format, which is easy for both humans and machines to read."*

*"SOAP, on the other hand, is more formal and heavyweight. It uses XML for everything - the request, the response, even the envelope that wraps it all. It comes with a strict contract called WSDL that defines exactly what the service can do. SOAP has built-in security features at the message level, which is why banks and government systems still use it."*

*"In my experience, when building modern applications for web and mobile, we almost always choose REST because it's simpler, faster, and easier to work with. But when I had to integrate with a banking client's legacy system, we had to use SOAP because that's what their mainframe exposed - we had to generate client code from their WSDL and handle SOAP faults."*

*"The key differences I've noticed: REST is stateless and cacheable, SOAP can maintain state and isn't easily cached. REST uses HTTP status codes for errors, SOAP has its own fault elements. For most new projects, REST is the way to go, but in enterprise environments, you'll still encounter SOAP systems that you need to integrate with."*

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys as standard theory, enterprise Java projects

#### Indepth
| Aspect | REST | SOAP |
|--------|------|------|
| Format | JSON (or XML) | XML only |
| Transport | HTTP | HTTP, SMTP, TCP |
| Contract | OpenAPI spec (optional) | WSDL (required, strict) |
| State | Stateless | Can be stateful |
| Caching | HTTP cache headers | Not cacheable by default |
| Overhead | Low | High (verbose XML) |
| Error handling | HTTP status codes | SOAP Fault element |
| Security | HTTPS + JWT | WS-Security (message-level) |
| Use case | Mobile/web APIs, microservices | Enterprise integrations, banking |

Why service-based companies still care about SOAP: Many government and banking systems (NPCI, banking core systems) expose SOAP interfaces. A TCS/Infosys developer integrating with these systems needs to understand WSDL files, generate client stubs (Apache CXF, JAX-WS), and handle SOAP faults.

---

## Q3. What is Spring Boot and what problem does it solve?

"Spring Boot is an opinionated framework built on top of Spring that solves the **configuration and boilerplate problem** of traditional Spring applications.

Traditional Spring required XML configuration files, web.xml deployment descriptors, and manual bean wiring. Setting up a new project took hours. Spring Boot replaced this with: **auto-configuration** (sensible defaults based on what's on the classpath), **starter dependencies** (spring-boot-starter-web pulls in everything you need for a web app), and **embedded server** (no separate Tomcat installation — it runs inside your JAR).

`spring-boot-starter-web` + one `@SpringBootApplication` annotation + `main()` method = a running HTTP server. What used to take an hour now takes 2 minutes."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** Every Java-based service company role

#### Indepth
Key Spring Boot features:
- **Auto-configuration:** `@EnableAutoConfiguration` scans the classpath and configures beans automatically. If Spring Data JPA is on classpath + a DB URL is configured → database connection is auto-configured.
- **Spring Initializr:** `start.spring.io` generates a project skeleton with selected dependencies in seconds.
- **Actuator:** Built-in endpoints for health, metrics, info, env — `/actuator/health`, `/actuator/metrics`. Critical for production monitoring.
- **Profiles:** `application-dev.properties`, `application-prod.properties` — environment-specific configuration.
- **Spring Security:** Add `spring-boot-starter-security` → all endpoints are secured by default. Configure roles, JWT, OAuth2 in Java (no XML).

For service-based company interviews: Know the Spring Boot auto-configuration mechanism (`@Conditional` annotations), understand how `@RestController` + `@Service` + `@Repository` (stereotype annotations) map to the 3-tier architecture, and be able to explain how the Spring IoC container works.

---

## Q4. What is microservices architecture and how does it differ from monolith? (Service company perspective)

"In a monolith, everything is in one deployable unit — one WAR file deployed to one Tomcat instance. All our modules (user management, order processing, reporting) are in one codebase. A change to reporting requires redeploying the whole application, including order processing.

In microservices, each business domain is a separate service with its own codebase, its own deployment, and ideally its own database. A change to reporting is deployed without touching order processing.

For a service company like TCS delivering a banking project: the legacy system is a monolith. The modernization project extracts each banking function (account management, loans, payments) into independent microservices. The work for us is: identifying the correct service boundaries, setting up the CI/CD pipelines for each service, and handling the integration testing between services."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS, Infosys, Wipro — especially modernization projects and banking clients

#### Indepth
How service companies work with microservices:
- Modernization projects: Converting a 10-year-old Java EE monolith to Spring Boot microservices
- New greenfield: Building a new fintech product for a banking client using microservices from scratch
- Integration work: Building middleware that connects multiple existing enterprise services

Key enterprise microservices patterns for service company projects:
1. **API Gateway (Spring Cloud Gateway):** Route external traffic to internal services
2. **Service Registry (Spring Cloud Eureka):** Services register and discover each other
3. **Config Server (Spring Cloud Config):** Centralized configuration from Git
4. **Distributed Tracing (Sleuth + Zipkin):** Track requests across services
5. **Circuit Breaker (Resilience4j):** Prevent cascading failures

---

## Q5. What is the difference between synchronous and asynchronous communication in microservices?

"**Synchronous** communication: Service A calls Service B and **waits for the response** before proceeding. REST APIs and gRPC are synchronous. Simple flow, immediate response, but tight coupling — if B is slow, A is slow. If B is down, A fails too.

**Asynchronous** communication: Service A sends a message to a queue/topic and **continues immediately**. Service B processes the message in its own time. RabbitMQ, Kafka, AWS SQS are async message brokers. Decoupled — A and B can fail independently. Harder to debug and trace.

Example: Sending an email after order placement. Synchronous approach: A crashes if the email service is slow. Asynchronous approach: A publishes 'order placed' event, email service picks it up and sends the email later. The user gets order confirmation immediately; the email may arrive 2 seconds later."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS, Infosys, Wipro — Java backend 3-5 year experience

#### Indepth
When to use each in enterprise projects:
- **Synchronous:** User authentication (must be immediate), payment confirmation (need to know success/failure now), inventory check (can't place order without knowing stock level)
- **Asynchronous:** Email/SMS notifications, audit log generation, reporting updates, data warehouse sync, PDF generation

RabbitMQ vs Kafka for service company projects:
- **RabbitMQ:** Simpler, well-documented, good for task queues. Most enterprise Spring Boot projects.
- **Kafka:** Higher throughput, event replay, better for data pipelines. Chosen for new greenfield projects or when event sourcing is needed.

Spring Boot integration:
- `spring-boot-starter-amqp` for RabbitMQ (with `@RabbitListener`)
- `spring-kafka` for Kafka (with `@KafkaListener`)

---

## Q6. What is Docker and why is it used in modern projects?

"Docker packages an application and all its dependencies into a **container** — a self-contained, portable unit that runs identically regardless of where it's deployed.

The classic problem it solves: 'It works on my machine.' With Docker, you define the entire environment (OS, Java version, libraries, config) in a `Dockerfile`. The image built from this file runs identically on a developer's laptop, the CI server, and the production AWS instance. No more environment configuration mismatches.

In a TCS/Infosys Java project: the Spring Boot application is packaged as a Docker image (jar + JRE in a container). Docker Compose runs multiple containers (app + MySQL + Redis) locally for development. In production, containers run in Kubernetes or ECS."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Service companies transitioning to DevOps practices

#### Indepth
Dockerfile for a Spring Boot app:
```dockerfile
FROM eclipse-temurin:17-jre-alpine
WORKDIR /app
COPY target/myapp.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
```

Docker Compose for local development:
```yaml
services:
  app:
    build: .
    ports: ["8080:8080"]
    environment:
      SPRING_DATASOURCE_URL: jdbc:mysql://db:3306/mydb
    depends_on: [db]
  db:
    image: mysql:8.0
    environment:
      MYSQL_DATABASE: mydb
      MYSQL_ROOT_PASSWORD: secret
```

Key Docker concepts for interviews:
- **Image:** Read-only template (built from Dockerfile)
- **Container:** Running instance of an image
- **Registry:** Docker Hub, AWS ECR — stores images
- **Volume:** Persistent storage that survives container restarts
- **Network:** Containers in the same network can communicate by service name

---

## Q7. What is the role of a load balancer in enterprise architecture?

"A load balancer **distributes incoming traffic across multiple server instances** so no single server is overwhelmed. It's the single entry point for client requests, hiding the complexity of multiple backend servers.

In a typical TCS enterprise Java project: the Spring Boot application is deployed across 3 servers for availability. A load balancer (Nginx or AWS ALB) sits in front. When 1000 users hit the application, requests are distributed across all 3 servers — each handles ~330 requests.

Benefits: **Availability** (if one server crashes, others continue handling traffic), **Scalability** (add more servers without clients knowing), **Zero-downtime deployment** (take one server offline for update, re-add it, repeat)."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Wipro, Infosys — standard infrastructure question

#### Indepth
Load balancing algorithms:
1. **Round Robin:** Each server gets requests in turn. Simple, equal distribution.
2. **Least Connections:** New request goes to the server with fewest active connections. Better for heterogeneous request types.
3. **IP Hash:** Client IP determines which server. Ensures same client always hits same server (session affinity). Required if sessions are not stored externally.
4. **Weighted:** Assign more requests to more powerful servers.

Nginx round-robin configuration:
```nginx
upstream myapp {
    server app1.example.com:8080;
    server app2.example.com:8080;
    server app3.example.com:8080;
}
server {
    listen 80;
    location / {
        proxy_pass http://myapp;
    }
}
```

**Session affinity vs stateless design:** If the Spring Boot app stores session state in memory (HttpSession), you need IP Hash load balancing (sticky sessions). Better design: make the app stateless, store sessions in Redis, and use any load balancing algorithm. This is the 12-factor app principle.

---

## Q8. What is CI/CD and how is it implemented in a service company project?

"CI/CD stands for Continuous Integration (CI) and Continuous Deployment/Delivery (CD). It's the practice of **automatically building, testing, and deploying code** every time a developer pushes a change.

CI: When a developer pushes to Git, an automated pipeline runs unit tests, integration tests, code quality checks (SonarQube), and builds the artifact (Docker image or JAR). If any step fails, the pipeline fails and the developer is notified.

CD (Delivery): The tested artifact is automatically deployed to a staging environment. The team can then manually trigger production deployment (for regulated environments like banking).

CD (Deployment): Full automation — if all tests pass, automatically deploy to production. Used by tech-first companies. Service company clients in banking/insurance typically require manual approval for production."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Service companies implementing DevOps for clients

#### Indepth
Typical service company CI/CD stack:
- **Version control:** Git (GitHub, Bitbucket, Azure DevOps)
- **CI server:** Jenkins (most common in enterprises), GitLab CI, Azure DevOps Pipelines
- **Code quality:** SonarQube (code smell, bug detection, coverage thresholds)
- **Artifact registry:** Nexus, JFrog Artifactory, Docker Hub
- **Deployment:** Jenkins pipelines deploying to Tomcat/Kubernetes; Ansible for configuration management

Jenkins pipeline (Jenkinsfile):
```groovy
pipeline {
    stages {
        stage('Build')  { steps { sh 'mvn package' } }
        stage('Test')   { steps { sh 'mvn test' }   }
        stage('Quality'){ steps { sh 'mvn sonar:sonar' } }
        stage('Docker') { steps { sh 'docker build -t myapp:${BUILD_NUMBER} .' } }
        stage('Deploy Staging') { steps { sh './deploy.sh staging' } }
        stage('Deploy Prod') {
            when { branch 'main' }
            input { message 'Deploy to production?' }
            steps { sh './deploy.sh prod' }
        }
    }
}
```

---

## Q9. What is API versioning and why is it important?

"API versioning allows you to **make changes to an API without breaking existing clients**. When you have 20 client applications using your API, you can't change the response format without telling every client team in advance — and they'll need time to update.

With versioning, you introduce `/v2/orders` with the new format while keeping `/v1/orders` running. Clients can migrate at their own pace. Once the migration windw has passed (6-12 months), you can retire v1.

This is especially important in service company projects where you're delivering APIs consumed by other teams within the same client organization — or consumed by the client's partners and third-party integrators who don't respond to urgent deployment timelines."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Service companies delivering API projects

#### Indepth
Versioning approaches:
1. **URL versioning** (most common in practice): `/api/v1/orders`, `/api/v2/orders`. Clear, easy to route, easy to test different versions.
2. **Header versioning**: `Accept: application/vnd.myapi.v2+json`. Cleaner URLs but harder to test in browser.
3. **Query param**: `/api/orders?version=2`. Simple but pollutes the query string.

Spring Boot URL versioning:
```java
@RestController
@RequestMapping("/api/v1/orders")
public class OrderControllerV1 { ... }

@RestController  
@RequestMapping("/api/v2/orders")
public class OrderControllerV2 { ... }  // New response format
```

What constitutes a breaking change in a service company context:
- Removing a JSON field (client's deserialization may fail)
- Renaming a field
- Changing a field type (string to integer)
- Making an optional field required
- Changing the HTTP status code of a response

Non-breaking: Adding a new optional field (clients should ignore unknown fields), adding a new endpoint, adding a new query parameter that's optional.

---

## Q10. What is the saga pattern and when would a service company implement it?

"The Saga pattern manages **multi-step business transactions across multiple services** when you can't use a single database transaction.

Service company context: You're building a banking onboarding system for an HDFC project. To onboard a new customer, you must: (1) Create a user record (User Service), (2) Open a savings account (Account Service), (3) Issue a debit card (Card Service), (4) Send a welcome SMS (Notification Service). These are 4 separate services with 4 separate databases.

If step 3 fails after steps 1 and 2 succeeded, you need to compensate — mark the account as 'onboarding failed', don't activate the user. The Saga defines the sequence and the compensating actions for each failure scenario."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Service companies working on banking/fintech clients with microservices

#### Indepth
Orchestration-based Saga (most practical for enterprise):
```
OnboardingOrchestrator:
  → sends CreateUser to UserService
     on success → sends OpenAccount to AccountService
       on success → sends IssueCard to CardService
         on success → sends SendWelcomeSMS to NotificationService (fire and forget)
         on failure → sends CloseAccount to AccountService (compensate)
                    → sends DeleteUser to UserService (compensate)
       on failure → sends DeleteUser to UserService (compensate)
```

**Axon Framework** (Java): Popular in service company Java projects. Provides Saga handling, event sourcing, and CQRS out of the box. Used in complex banking domain implementations.

**AWS Step Functions:** Serverless saga orchestration. Define the workflow as a state machine in JSON. Each state calls a Lambda function or an API. Built-in retry, error handling, and compensation logic. Used in AWS-centric enterprise projects.

The key complexity: ensuring compensating transactions are idempotent (they may be called multiple times). A `CloseAccount` compensation that runs twice must not fail on the second run — it should check if the account is already closed and succeed idempotently.
