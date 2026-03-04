# NodeJS Interview Questions - Product Based Companies

This section contains Node.js interview questions frequently asked by Product-Based companies (e.g., Google, Amazon, Microsoft, Uber, Netflix, Stripe). Questions in this category are deeply technical, heavily focused on complex architecture, internals, scaling, security, and low-level performance tuning.

## 📂 File Structure

| Topic | Description | Difficulty |
|-------|-------------|------------|
| [01 — Node.js Internals & V8 Engine](./01_nodejs_internals_and_v8.md) | Libuv, V8 compilation pipeline, Garbage Collection internals, Memory memory profiling. | ⚙️ Hard |
| [02 — Advanced Async & Streams](./02_advanced_async_and_streams.md) | Backpressure, Streams pipelines, Worker Threads for CPU tasks. | 🌊 Hard |
| [03 — System Design & Architecture](./03_system_design_and_architecture.md) | Outbox pattern, Pub/Sub, API Gateways, and high-concurrency WebSocket scaling. | 🏛️ Hard |
| [04 — Scaling & Microservices](./04_scaling_and_microservices.md) | Sagas, Service Discovery, gRPC vs GraphQL, Circuit Breaker patterns. | 🛠️ Hard |
| [05 — Advanced Security & Auth](./05_advanced_security_and_auth.md) | Token blacklisting, CSRF/SSRF prevention, robust session management. | 🛡️ Hard |
| [06 — Machine Coding Rounds](./06_machine_coding_rounds.md) | Implementation of rate limiters, task queues, and custom event emitters. | 💻 Hard |

## 🎯 Interview Focus
For product-based companies, demonstrating that you can use a library (like Express or Mongoose) is not enough. You must show how it works under the hood. You should be able to:
- Explain what happens inside V8 when a variable is instantiated.
- Describe how Libuv offloads blocking calls.
- Design highly available distributed real-time systems.
- Explain trade-offs between different communication protocols (gRPC, REST, PubSub).
- Write clean, asynchronous code from scratch under time constraints in machine coding rounds.
