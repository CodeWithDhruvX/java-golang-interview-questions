# 🐍 Python Interview Questions — Product-Based Companies

> Questions tailored for **Google, Microsoft, Amazon, Flipkart, Swiggy, Razorpay, Zepto** and similar product-based organizations.
> Difficulty: 🔴 Hard | Focus: Internals, advanced patterns, system design, machine coding

---

## 📂 Topics Covered

| # | File | Topics |
|---|------|--------|
| 01 | [Python Internals & CPython](./01_python_internals_and_cpython.md) | Bytecode, `dis`, reference counting, GC, object model, `__slots__`, descriptors |
| 02 | [Advanced OOP & Metaclasses](./02_advanced_oop_and_metaclasses.md) | Metaclasses, `type`, `__new__` vs `__init__`, Singleton, Registry, Mixins, ABCs, `dataclasses` |
| 03 | [Concurrency, GIL & Async](./03_concurrency_gil_and_async.md) | GIL deep dive, asyncio internals, tasks, semaphores, rate limiting, bridging sync/async |
| 04 | [Memory Management & Optimization](./04_memory_management_and_optimization.md) | `gc`, `tracemalloc`, `weakref`, object pools, generator pipelines, memory profiling |
| 05 | [Decorators, Generators & Context Managers](./05_decorators_generators_and_context_managers.md) | Class-based decorators, async decorators, bidirectional generators, `contextlib.ExitStack` |
| 06 | [Design Patterns in Python](./06_design_patterns_in_python.md) | Singleton, Factory, Builder (QueryBuilder), Observer, Strategy, Command, Adapter, Proxy |
| 07 | [Data Science & ML Python](./07_data_science_and_ml_python.md) | NumPy broadcasting, Pandas groupby/merge/pivot, serialization (JSON/Pickle/Parquet), Pydantic |
| 08 | [System Design with Python](./08_system_design_with_python.md) | Redis caching, rate limiter, Circuit Breaker, Celery task queues, FastAPI middleware |
| 09 | [Testing, Debugging & Profiling](./09_testing_debugging_and_profiling.md) | Advanced pytest, property-based testing (hypothesis), cProfile, line_profiler, pdb, logging |
| 10 | [Machine Coding Rounds](./10_machine_coding_rounds.md) | LRU Cache (thread-safe), Parking Lot System, Event-Driven Notification Bus |

---

## 🎯 Interview Focus Areas

### Expected in Phone Screens:
- CPython memory management (refcounting, GIL)
- asyncio: how coroutines, tasks, and the event loop interact
- Custom decorators + metaclasses

### Expected in Design Rounds:
- System design using Python: caching, rate limiting, circuit breakers
- Task queues and background workers (Celery, RQ)
- Microservice communication patterns

### Expected in Machine Coding:
- Thread-safe data structures (LRU Cache, Connection Pool)
- Complete OOP system design (Parking Lot, File System, Snake Game)
- Event-driven architectures

---

## 📝 Quick Tips
- Always think about thread safety when designing classes
- Use `asyncio.gather` for concurrent I/O, `ProcessPoolExecutor` for CPU-bound
- Demonstrate `@dataclass`, `__slots__`, and descriptors — they impress interviewers
- For system design, always mention: caching strategy, failure modes, rate limiting
- Write at least 3–5 unit tests for your machine coding solution
