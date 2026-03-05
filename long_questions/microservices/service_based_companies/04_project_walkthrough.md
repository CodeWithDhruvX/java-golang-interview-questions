# 🗣️ Microservices — Practical Project Walk-Through Template

> **Level:** 🟢 Junior – 🟡 Mid – 🔴 Senior
> **Asked at:** ALL companies — TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini, HCL, Zepto, Swiggy, Razorpay

> **Why this is critical:** *"Tell me about the microservices project you worked on"* is asked in 90%+ of interviews. Without a structured, compelling answer, even technically strong candidates fail. This file gives you ready-to-use templates.

---

## THE SITUATION: What Every Interviewer Really Wants to Know

When they ask about your project, they want to verify:
1. Did you actually work on real microservices, or just read about them?
2. Do you understand the architectural decisions made and WHY?
3. Can you articulate problems and how you solved them?
4. What was your specific contribution (not the team's)?

---

## TEMPLATE 1: E-Commerce / Order Management System (Most Universal)

**Use this if:** Your project involved orders, payments, inventory, or general CRUD microservices.

---

### 🎯 The Perfect 3-Minute Answer (Read, Internalize, Personalize)

*"In my previous project at [Company Name], I worked on a [B2B/B2C] e-commerce platform built on a microservices architecture. I'll walk you through the architecture and my specific contributions.*

**System Overview:**  
The platform was composed of around 8 core services: User Service, Product Catalog Service, Order Service, Payment Service, Inventory Service, Notification Service, Shipping Service, and an API Gateway. Each service had its own database — we followed the Database-per-Service pattern strictly.

**Tech Stack:**  
The services were built in Spring Boot 3, deployed on Kubernetes (EKS), communicated synchronously via REST (Feign clients for inter-service calls) and asynchronously via Apache Kafka for events like order placement and payment confirmation. We used PostgreSQL for transactional services and Redis for caching product catalog data.

**A Specific Problem I Solved:**  
The most interesting challenge I worked on was a *[CHOOSE ONE from below]*:

---

**Problem Option A — Dual Write / Outbox (Fintech-style):**
*"The Order Service needed to create an order record in PostgreSQL AND publish an `OrderPlaced` event to Kafka. The problem was: if the DB write succeeded but Kafka publish failed (Kafka was down), we'd lose the event. To solve this, I implemented the **Transactional Outbox Pattern**. Instead of writing directly to Kafka, the OrderService writes to both the orders table and an outbox_events table in a single ACID transaction. A separate polling service (OutboxRelay) reads from the outbox table and publishes to Kafka, marked status as SENT. This gave us guaranteed at-least-once delivery without the dual-write problem."*

**Problem Option B — Performance Issue:**
*"We had a performance problem where the Order Detail API was taking 1.5 seconds due to N+1 database queries — it was fetching order items one by one. I identified this using Hibernate's batch fetch logging and resolved it by replacing the `@OneToMany` lazy loading with a JOIN FETCH query, then adding a Redis cache for Order Detail with a 5-minute TTL. This brought the P99 latency from 1.5s down to 120ms."*

**Problem Option C — Service Communication Failure:**
*"The Order Service called Payment Service synchronously via Feign, but payments sometimes took 10+ seconds (fraud checks). This was causing thread pool exhaustion in the Order Service. I changed the architecture: Order Service publishes a `ProcessPayment` event to Kafka, Payment Service consumes it, processes asynchronously, and publishes back a `PaymentCompleted` or `PaymentFailed` event. Order Service consumed these events to update order status. This decoupled the services completely and the Order Service thread pool exhaustion issue was resolved."*

---

**My Specific Contribution:**  
*"My specific responsibility was the Order Service and its integration with the Payment Service. I designed the Order state machine (PENDING → PAYMENT_PROCESSING → CONFIRMED / FAILED / CANCELLED), wrote the Kafka consumers and producers, implemented the Outbox pattern, and wrote the JUnit + Testcontainers integration tests."*

**Scale/Impact:**  
*"The system handled around [X] orders per day / [X] concurrent users. My work on the outbox pattern eliminated order event loss that was causing support tickets."*"

---

### 📋 Follow-Up Questions & Ready Answers

**Q: "Why did you choose Kafka over RabbitMQ?"**  
*"We chose Kafka because we needed event replay capability — if a downstream service (Inventory) was down for 2 hours and came back up, it needed to catch up on all missed events. Kafka retains events (we set 7-day retention). RabbitMQ is a pure message queue — once consumed, the message is gone. Also, Kafka's consumer groups allowed us to add new consumers (like an Analytics Service) without modifying the producer."*

**Q: "How did services authenticate with each other?"**  
*"Internally, services communicate through the Kubernetes network inside a private namespace — only the API Gateway is exposed externally. For inter-service authentication, we used JWT tokens signed by the Auth Service. The API Gateway validates the JWT, and internal services trust the gateway-passed user context header. We also had mutual TLS configured via Istio sidecar to encrypt all network traffic inside the cluster."*

**Q: "How did you handle service failures?"**  
*"We implemented the Circuit Breaker pattern using Resilience4j. If the Payment Service was failing more than 50% of requests in a 10-second window, the circuit opened and we returned a fast failure with a 'payment temporarily unavailable' message rather than letting threads wait. We also had timeouts configured — max 3 seconds for synchronous calls. For async Kafka consumers, we had a Dead Letter Topic — messages that failed processing 3 times were sent to a DLT for manual review."*

**Q: "How was your CI/CD pipeline set up?"**  
*"We used GitHub Actions for CI — on every PR, it ran unit tests, integration tests (Testcontainers), SonarQube code quality check, and built a Docker image tagged with the commit SHA. For CD, we used ArgoCD with GitOps — merging to main updated the image tag in the Helm chart values file, and ArgoCD automatically synced the change to our Kubernetes cluster on EKS. Staging was auto-deployed; prod required a manual promotion step."*

**Q: "How did you monitor the system?"**  
*"We used the observability stack: Prometheus for metrics (Spring Actuator + Micrometer), Grafana for dashboards showing order success rate, payment failure rate, and service latency. Loki for centralized logs with structured JSON logging using SLF4J + Logback. Jaeger for distributed tracing — every request carried a TraceId through all services so we could see the full journey in the Jaeger UI."*

---

## TEMPLATE 2: Banking / Fintech System

**Use this if:** Your project involved accounts, transactions, wallets, or compliance.

*"I worked on a digital banking platform's microservices backend. The system was divided into Account Service, Transaction Service, KYC Service, Notification Service, Statement Service, and Fraud Detection Service.*

The most significant challenge was ensuring financial consistency across services — specifically the transaction flow: debit → fraud check → credit → notify. If fraud check failed, we needed to automatically trigger a reversal debit.

I designed and implemented the **Saga Orchestration Pattern** for this flow. A central Saga orchestrator called each step in sequence, and on failure, called the compensating transaction in reverse order (e.g., if the credit step failed after the debit, the orchestrator triggered a reversal credit). The saga state was persisted in a database table, so even if the orchestrator pod crashed, on restart it resumed from the last saved state.

I specifically built the Transaction Saga Orchestrator and the Compensation Logic. The system processed around [X] transactions per day with zero unresolved saga failures over a 6-month period."*

---

## TEMPLATE 3: For Freshers / SDET / Minimal Production Experience

**Use this if:** You worked on a college project, training project, or only QA/testing.

*"In my final year project / training at [Institution/Company], I built a microservices-based food ordering system similar to Swiggy as a learning exercise.*

The system had 5 services: User, Restaurant, Menu, Order, and Notification. I deployed them on a local Kubernetes cluster using Minikube and used Docker Compose for local development. Services communicated via REST APIs with a Spring Cloud Gateway as the entry point.*

The most valuable learning was understanding why you CAN'T just do a multi-table JOIN across services — since each service owns its database, the Order Service can't directly join with the User Service's table. Instead, the Order Service stores userId as a foreign reference and fetches user details via API when needed. I also learned about the Circuit Breaker pattern using Resilience4j after simulating service failures.*

Though this was a project rather than production, it gave me a solid foundation in microservices principles, and I'm eager to apply and deepen this in a production environment."*

---

## 🏆 Golden Rules for Project Walk-Throughs

| Rule | Why |
|------|-----|
| **Always have 3 specific problems ready with solutions** | Distinguishes real experience from resume reading |
| **Know the numbers**: daily orders, concurrent users, latency SLAs | Shows engineering judgment, not just code writing |
| **Say "I specifically" not just "we"** | Interviewer needs your individual contribution |
| **Mention what DIDN'T work and how you fixed it** | Shows maturity — nobody's first design is perfect |
| **End with business impact** | Shows you understand why engineering decisions matter |
| **Never trash the architecture** | Even if it was bad, say "Given the constraints, the team chose X. In hindsight, I'd have done Y because…" |

---

## Common Traps to Avoid

❌ **Don't say:** *"We used microservices"* — say WHICH services and WHY they were split that way  
❌ **Don't say:** *"We used REST"* — say when REST and when Kafka, and why  
❌ **Don't say:** *"I handled authentication"* — say HOW (JWT, OAuth2, which library, how tokens were validated)  
❌ **Don't just describe the happy path** — interviewers care most about failure handling  
✅ **Do say:** *"The challenge was..."* — this structure signals engineering maturity
