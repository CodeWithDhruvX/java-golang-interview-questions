# 🏗️ Architecture — Product-Based Companies

> **Track:** Senior Engineering Roles
> **Companies:** Amazon, Google, Flipkart, Uber, Swiggy, Zomato, Razorpay, PhonePe, CRED, Zepto, Groww

---

## 📁 Files in This Folder

| File | Topics Covered | Level |
|------|---------------|-------|
| [01_architecture_deep_dive.md](./01_architecture_deep_dive.md) | E-commerce orders, ride-matching, payment gateways, feeds, rate limiters, DB migrations | 🔴 Senior |
| [02_distributed_systems_tradeoffs.md](./02_distributed_systems_tradeoffs.md) | Consistent hashing, Raft/Paxos, vector clocks, sharding | 🔴 Staff |
| [03_real_world_architecture_cases.md](./03_real_world_architecture_cases.md) | URL shortener, WhatsApp/chat, search autocomplete, cache, notifications | 🔴 Senior |
| [04_database_internals.md](./04_database_internals.md) | B-Tree, LSM-Tree, WAL, MVCC, query plan scan types | 🔴 Senior |
| [05_observability_distributed_systems.md](./05_observability_distributed_systems.md) | Logs, metrics, traces, SLI/SLO/SLA, OpenTelemetry, ELK stack | 🔴 Senior |
| [06_multi_region_geo_distributed.md](./06_multi_region_geo_distributed.md) | Active-active vs passive, Route53, cross-region replication, CDN | 🔴 Staff |
| [07_live_coding_architecture.md](./07_live_coding_architecture.md) | Rate limiter, Task scheduler, Event Broker (in-memory implementation) | 🔴 Senior |
| [17_ml_infrastructure_model_serving.md](./17_ml_infrastructure_model_serving.md) | Feature stores, model training, real-time inference, MLOps, A/B testing, monitoring | 🔴 Senior |
| [18_advanced_streaming_realtime_processing.md](./18_advanced_streaming_realtime_processing.md) | Stream processing engines, windowing, state management, exactly-once, backpressure | 🔴 Senior |

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
- **Feature store:** Centralized feature management
- **Model serving:** Real-time and batch inference
- **Stream processing:** Windowing, state management, exactly-once

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
- Design a real-time analytics dashboard (product companies)
- Design a fraud detection system (fintech companies)
- Design a recommendation system (content platforms)
- Design a ML model serving pipeline (AI companies)

---

## 📖 Recommended Preparation Order

1. Core Theory: `theory/01_Architecture_Fundamentals.md` and `theory/02_Microservices_Architecture.md`
2. Database Deep Dive: `theory/06_Data_Architecture.md` + `product_based_companies/04_database_internals.md`
3. Distributed Systems: `product_based_companies/02_distributed_systems_tradeoffs.md` + `product_based_companies/06_multi_region_geo_distributed.md`
4. Resilience & EDA: `theory/03_Event_Driven_Architecture.md` + `theory/09_Reliability_Resilience.md`
5. Observability: `product_based_companies/05_observability_distributed_systems.md`
6. Practice HLD Scenarios: `product_based_companies/01_architecture_deep_dive.md` and `product_based_companies/03_real_world_architecture_cases.md`
7. Prepare for Live Coding: `product_based_companies/07_live_coding_architecture.md`
8. ML Infrastructure: `product_based_companies/17_ml_infrastructure_model_serving.md`
9. Streaming & Real-time: `product_based_companies/18_advanced_streaming_realtime_processing.md`
