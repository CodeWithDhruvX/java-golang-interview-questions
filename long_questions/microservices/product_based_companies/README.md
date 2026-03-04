# 🏗️ Microservices — Product-Based Companies

> **Track:** Senior Engineering Roles
> **Companies:** Amazon, Google, Flipkart, Uber, Swiggy, Zomato, Razorpay, PhonePe, CRED, Zepto, Groww

---

## 📁 Files in This Folder

| File | Topics Covered | Level |
|------|---------------|-------|
| [01_microservices_deep_dive.md](./01_microservices_deep_dive.md) | Sagas, dual-writes, outbox pattern, resilience, sharding, CQRS, contracts, service mesh | 🔴 Senior |

---

## 🎯 Interview Strategy for Product Companies

### What They Care About
Product companies do not want definitions of "What is a microservice?". They want to know you understand the **fallacies of distributed computing** and how to architect around them at scale.

1. **Distributed Data:** How do you join data across 4 services? (CQRS, Materialized Views, API Composition)
2. **Distributed Transactions:** How do you roll back step 2 if step 3 fails? (Saga, Compensation, 2PC)
3. **Eventual Consistency:** Can you tolerate dirty reads? (Idempotency, Outbox Pattern)
4. **Resilience & Cascading Failures:** How do you stop a traffic spike from taking down the cluster? (Timeouts, Deadlines, Bulkheads, Circuit Breakers)
5. **Testing Architecture:** Why E2E tests fail at scale and what to use instead. (Contract Testing, Chaos Engineering)

### Key Patterns to Know Cold
- **Transactional Outbox Pattern** (Crucial for Kafka integration)
- **Saga Pattern** (Orchestration vs Choreography)
- **CQRS** (Command Query Responsibility Segregation)
- **Idempotency Keys** (Critical for payment/fintech)
- **Service Mesh** (Istio/Envoy - why and when)

---

## 🔥 Most Frequently Asked Systems / Concepts

- "We have a massive monolith. Walk me through how you would strangle it to microservices." (Strangler Fig Pattern)
- "Design a distributed rate limiter that works across 50 API Gateway nodes." (Redis + Lua, Token Bucket)
- "How do you guarantee a message is processed exactly once in Kafka?" (Idempotence, Consumer Offsets)
- "An order request spans 6 microservices. How do you trace where it failed?" (OpenTelemetry, correlation IDs)

---

## 📖 Recommended Preparation Order

1. Review `theory/03_Distributed_Data_Transactions.md`
2. Review `theory/04_Messaging_Systems.md`
3. Study `product_based_companies/01_microservices_deep_dive.md`
4. Dive into `theory/10_Real_Production_Scenarios.md` and `theory/12_Architect_Deep_Dive.md`
