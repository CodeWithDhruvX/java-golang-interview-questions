# 🏗️ Architecture — Product-Based Companies

> **Track:** Senior Engineering Roles
> **Companies:** Amazon, Google, Flipkart, Uber, Swiggy, Zomato, Razorpay, PhonePe, CRED, Zepto, Groww

---

## 📁 Files in This Folder

| File | Topics Covered | Level |
|------|---------------|-------|
| [01_architecture_deep_dive.md](./01_architecture_deep_dive.md) | E-commerce orders, ride-matching, payment gateways, feeds, rate limiters, DB migrations | 🔴 Senior |

---

## 🎯 Interview Strategy for Product Companies

### What They Care About
1. **Trade-offs:** Never one correct answer — always compare approaches and justify
2. **Scale:** "Design for 10M users" is not the same as "Design for 100"
3. **Failure modes:** What happens when each component fails?
4. **NFRs first:** Define availability, consistency, and latency requirements before designing
5. **Depth on demand:** Be ready to go deep on any component

### Framework for Design Questions
1. **Clarify requirements** (3-5 min): Scale, SLA, consistency needs, client types
2. **Define core entities** (2 min): What are the main data objects?
3. **HLD + data flow** (10 min): Draw the system, walk through request flow
4. **DB design** (5 min): Schema, choice of database type
5. **Scale + Bottlenecks** (5 min): What breaks first, how to handle it
6. **Failure handling** (3 min): Circuit breakers, retries, fallbacks

### Common Patterns to Know Cold
- **Outbox pattern:** DB + event publish atomicity
- **Saga pattern:** Distributed transactions
- **CQRS:** Separate read/write models
- **Circuit breaker:** Failure isolation
- **Consistent hashing:** Distributed caching and sharding
- **Token bucket:** Rate limiting
- **Fan-out on write/read:** Feed architectures

---

## 🔥 Most Frequently Asked Systems

- Design a payment gateway (Razorpay, PhonePe)
- Design an order management system (Flipkart, Amazon)
- Design a ride-matching system (Uber, Ola)
- Design a social media feed (Instagram, Twitter)
- Design a notification system (every company)
- Design a URL shortener (standard interview question)
- Design a distributed rate limiter (API companies)
- Design a cache system (Redis internals)

---

## 📖 Recommended Preparation Order

1. Start with `theory/01_Architecture_Fundamentals.md`
2. Read `theory/02_Microservices_Architecture.md`
3. Study `theory/03_Event_Driven_Architecture.md`
4. Understand `theory/06_Data_Architecture.md`
5. Practice with `product_based_companies/01_architecture_deep_dive.md`
6. Then go deep: `theory/04_Domain_Driven_Design.md` + `theory/09_Reliability_Resilience.md`
