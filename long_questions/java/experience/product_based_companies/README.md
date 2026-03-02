# 🚀 Java Interview Questions — Product-Based Companies

> Companies like **Amazon, Flipkart, Paytm, Swiggy, Zomato, Razorpay, PhonePe, CRED, Meesho, Myntra**, etc.

Product-based companies ask **deep, scenario-based questions** on Java internals, distributed systems, concurrency patterns, and system design. Expect **whiteboard-style** coding, architecture discussions, debugging scenarios, and modern Java features.

---

## 📂 Category Index

| # | File | Topic | Difficulty |
|---|------|--------|------------|
| 01 | [Data Structures & Algorithms](./01_data_structures_algorithms.md) | DSA in Java — collections, custom structures, LeetCode patterns | 🟡 Medium–Hard |
| 02 | [Advanced Concurrency](./02_advanced_concurrency.md) | ReentrantLock, ForkJoin, CAS, CompletableFuture, Virtual Threads | 🔴 Hard |
| 03 | [System Design](./03_system_design.md) | Rate limiting, distributed caching, URL shortener, CQRS, CAP theorem | 🔴 Hard |
| 04 | [JVM Internals](./04_jvm_internals.md) | Memory model, GC algorithms, class loading, JIT, escape analysis | 🔴 Hard |
| 05 | [Performance Optimization](./05_performance_optimization.md) | JMH benchmarks, profiling, HikariCP tuning, allocation reduction | 🔴 Hard |
| 06 | [Microservices & Kafka Advanced](./06_microservices_and_kafka_advanced.md) | Kafka internals, Saga pattern, Event Sourcing, Kafka Streams, service mesh | 🟡 Medium–Hard |
| 07 | [Spring Advanced & Design Patterns](./07_spring_advanced_and_design_patterns.md) | Spring Security JWT, AOP, Strategy, Factory, SOLID, Spring Events | 🟡 Medium–Hard |
| 08 | [Advanced Java Topics](./08_advanced_topics.md) | Records, sealed classes, pattern matching, virtual threads, generics, reflection | 🔴 Hard |

---

## 🎯 Interview Focus Areas (Product Companies)

- ✅ **DSA**: Arrays, heaps, graphs, trees — implemented in Java idioms
- ✅ **Concurrency**: JMM, ReentrantLock, ForkJoin, CompletableFuture, Virtual Threads
- ✅ **JVM Internals**: GC algorithms, memory model, JIT, class loading, heap dumps
- ✅ **System Design**: Rate limiting, caching, distributed locks, CAP theorem, CQRS
- ✅ **Kafka Advanced**: Replication, Saga pattern, Event Sourcing, Kafka Streams
- ✅ **Spring Deep Dive**: AOP internals, JWT security, custom annotations, design patterns
- ✅ **Modern Java**: Records, sealed classes, pattern matching (Java 16–21)
- ✅ **Performance**: JMH, profiling, connection pool tuning, memory leak diagnosis

---

## 💡 Tips for Product Company Java Interviews

1. **Master DSA in Java** — use `ArrayDeque` for stack/queue, `PriorityQueue` for heaps, `TreeMap` for ordered maps
2. **Know JVM internals deeply** — GC algorithms, what triggers Full GC, how to tune G1GC/ZGC
3. **Concurrency is critical** — ReentrantLock vs synchronized, JMM happens-before, CompletableFuture chaining
4. **System design with Java** — rate limiting, distributed locking with Redis, caching layers
5. **Kafka Saga + Outbox** — explain Saga orchestration vs choreography, Outbox for atomicity
6. **Know modern Java** — Records, sealed classes, switch expressions are frequently asked
7. **SOLID + design patterns** — Strategy, Factory, Observer are must-know; link to Spring examples
8. **Performance mindset** — profiling tools (async-profiler, JFR), JMH for benchmarks, avoid premature optimization
