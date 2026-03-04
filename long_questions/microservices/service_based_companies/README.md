# 🏗️ Microservices — Service-Based Companies

> **Track:** Mid-level Engineering Roles
> **Companies:** TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini, HCL, Tech Mahindra

---

## 📁 Files in This Folder

| File | Topics Covered | Level |
|------|---------------|-------|
| [01_microservices_fundamentals_qa.md](./01_microservices_fundamentals_qa.md) | Monolith vs Microservices, sync/async, API Gateway, Service Discovery, Circuit Breaker, Config Server | 🟢 Junior – 🟡 Mid |
| [02_spring_cloud_implementation.md](./02_spring_cloud_implementation.md) | Feign Client, Sleuth + Zipkin tracing, JWT security flow, Spring Boot Actuator, Docker + Kubernetes deployment | 🟡 Mid |

---

## 🎯 What Service Companies Ask

1. **Definitions & Differences:** "What is a microservice?" and "How is it different from SOA or Monolith?"
2. **Spring Cloud Architecture:** Instead of generic distributed systems questions, service companies want to know if you know the **Netflix OSS / Spring Cloud stack** (Eureka, Zuul/Gateway, Resilience4j, Config Server).
3. **API Gateway vs Load Balancer:** They test if you know the difference between routing business logic vs. purely spreading TCP/HTTP traffic.
4. **Resilience patterns:** Explain Circuit Breakers. What happens when service B goes down?
5. **Security:** How do you pass user identity between 5 microservices? (Usually JWT passed in headers).
6. **Deployments:** Docker and basic Kubernetes questions.

---

## 💡 Interview Tips for Service Companies

- **Focus on the framework implementation:** Service companies often look for hands-on knowledge. Instead of just explaining a "Circuit Breaker", explain *how* you implemented it using `@CircuitBreaker` annotation in Spring Boot with Resilience4j.
- **Reference existing projects:** "In my previous banking project, we broke down the core monolithic application into 5 microservices using Spring Boot..."
- **Don't overcomplicate:** Keep answers grounded in how Enterprise Java works today. You don't need to explain Raft consensus algorithms unless asked; focus on standard REST APIs and basic RabbitMQ/Kafka async flows.

---

## 📖 Recommended Study Order

1. Read `theory/01_Core_Fundamentals.md`
2. Read `theory/02_Communication.md`
3. Study `service_based_companies/01_microservices_fundamentals_qa.md`
4. Study `service_based_companies/02_spring_cloud_implementation.md`
5. Review `theory/08_Observability.md` (to understand Elk/Splunk/Zipkin for tracing)
