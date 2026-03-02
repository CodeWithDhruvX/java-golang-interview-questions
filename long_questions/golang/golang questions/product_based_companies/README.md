# 🚀 Go Interview Questions — Product-Based Companies

> Companies like **Google, Meta, Uber, Netflix, Stripe, Cloudflare, GitHub, HashiCorp, CockroachDB, Grafana, Docker, Kubernetes/CNCF projects**, etc.

Product-based companies ask **deep, scenario-based questions** on Go internals, system design, concurrency patterns, and performance optimization. Expect **whiteboard-style** coding, architecture discussions, and debugging scenarios.

---

## 📂 Category Index

| # | File | Topic | Difficulty |
|---|------|--------|------------|
| 01 | [Data Structures & Algorithms](./01_data_structures_algorithms.md) | DSA in Go — arrays, maps, trees, graphs | 🟡 Medium–Hard |
| 02 | [Concurrency & Goroutines](./02_concurrency_goroutines.md) | Advanced goroutines, channels, sync primitives | 🔴 Hard |
| 03 | [System Design in Go](./03_system_design_in_go.md) | Distributed systems, scalability, HLD/LLD | 🔴 Hard |
| 04 | [Go Internals & Runtime](./04_go_internals_runtime.md) | Scheduler, GC, memory model, escape analysis | 🔴 Hard |
| 05 | [Performance Optimization](./05_performance_optimization.md) | pprof, benchmarks, allocations, GC tuning | 🔴 Hard |
| 06 | [Microservices & gRPC](./06_microservices_grpc.md) | gRPC, protobuf, service mesh, Kafka | 🟡 Medium–Hard |
| 07 | [Design Patterns](./07_design_patterns.md) | Factory, Strategy, Observer, CQRS, Clean Arch | 🟡 Medium–Hard |
| 08 | [Advanced Topics](./08_advanced_topics.md) | Generics, reflection, unsafe, WebAssembly | 🔴 Hard |

---

## 🎯 Interview Focus Areas (Product Companies)

- ✅ DSA: arrays, linked lists, trees, graphs, heaps in Go
- ✅ Advanced concurrency: race conditions, lock-free structures, patterns
- ✅ Go runtime internals: GC, goroutine scheduler (GMP model), escape analysis
- ✅ System design: distributed caching, rate limiting, event sourcing
- ✅ Performance: profiling with pprof, memory optimization
- ✅ gRPC + protobuf + microservices architecture
- ✅ Design patterns applied to Go idioms
- ✅ Generics, reflection, unsafe package

---

## 💡 Tips for Product Company Interviews

1. **Master concurrency** — race conditions, WaitGroup, atomic operations, channels
2. **Study Go internals** — GMP scheduler, GC phases, escape analysis
3. **Practice DSA in Go** — implement common algorithms using Go idioms
4. **System design with Go** — rate limiting, pub-sub, distributed lock
5. **Profile your code** — use `pprof`, benchmark tests, `-race` flag
6. **Know generics** — type constraints, generic data structures
7. **Understand trade-offs** — when to use goroutines vs sync.Pool vs mutex
