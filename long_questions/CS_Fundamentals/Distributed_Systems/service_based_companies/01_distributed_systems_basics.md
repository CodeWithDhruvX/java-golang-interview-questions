# 📡 Distributed Systems Basics — Interview Questions (Service-Based Companies)

This document covers distributed systems concepts commonly tested at service-based companies like TCS, Infosys, Wipro, Capgemini, HCL. Targeted at 2–5 years of experience rounds.

---

### Q1: What is the CAP theorem? Explain with examples.

**Answer:**
The **CAP theorem** states that a distributed system can guarantee at most **2 out of 3** of the following properties simultaneously:

- **C — Consistency**: Every read receives the most recent write or an error (all nodes see the same data at the same time).
- **A — Availability**: Every request receives a response (not an error), though it might not be the most recent data.
- **P — Partition Tolerance**: The system continues to operate even when network partitions (message loss/delays) occur between nodes.

**Key insight:** In real-world distributed systems, **network partitions ALWAYS happen** (cables fail, switches reboot, etc.). Therefore, you MUST accept P. The real choice is between **CP** and **AP**.

**Examples:**

| System | CAP Choice | Reasoning |
|---|---|---|
| **Zookeeper, etcd** | CP | Returns error on partition — prefers consistency over availability |
| **Cassandra, DynamoDB** | AP | Returns potentially stale data — prefers availability |
| **HBase** | CP | Single master — consistency but may be unavailable if master fails |
| **MongoDB (default)** | CP | Linearizable reads from primary |
| **Couchbase** | AP (configurable) | Tunable consistency |

---

### Q2: What is eventual consistency? How does it differ from strong consistency?

**Answer:**

**Strong (Linearizable) Consistency:**
- After a write completes, all subsequent reads from ANY node return the updated value.
- Requires synchronization between nodes → higher latency.
- Example: Bank account balance — must always be accurate.

**Eventual Consistency:**
- After a write, replicas will EVENTUALLY converge to the same value, but there's no guarantee when.
- During the convergence window, different nodes may return different (stale) values.
- Higher availability, lower latency.
- Example: Amazon product review count — OK if slightly delayed.

**Consistency levels (Cassandra example):**
```
Write Consistency:
  ONE    → Write to 1 replica (fastest, can lose data)
  QUORUM → Write to majority of replicas (balanced)
  ALL    → Write to all replicas (slowest, strongest)

Read Consistency:
  ONE    → Read from 1 replica (can return stale)
  QUORUM → Read from majority (consistent if write was QUORUM)
  ALL    → Read from all (strongest)
```

**Strong consistency:** Read_CL + Write_CL > Replication_Factor (e.g., QUORUM + QUORUM with RF=3: 2+2=4>3 ✓)

---

### Q3: What is service discovery and why is it needed in microservices?

**Answer:**
In microservices, services are dynamic — they start/stop, scale up/down, and IP addresses change constantly. **Service discovery** is the mechanism by which services find each other without hardcoded addresses.

**Why needed:**
- Containers get new IPs at every restart.
- Multiple instances of a service run for load balancing.
- Services are deployed across different nodes/regions.

**Patterns:**

**1. Client-Side Discovery:**
- Client queries a Service Registry (like Consul, Eureka) to get a list of instances.
- Client does its own load balancing logic.
```
Client → Registry: "Where are payment-service instances?"
Registry → Client: ["10.0.1.5:8080", "10.0.1.7:8080"]
Client → load balance → 10.0.1.7:8080
```

**2. Server-Side Discovery:**
- Client sends request to a Load Balancer/API Gateway.
- Load Balancer queries Service Registry internally.
- Client doesn't handle discovery logic.
```
Client → Load Balancer/Gateway → Registry (internal) → Backend
```

**Tools:**
- **Consul**: Service registry + health checking + key-value store.
- **Eureka** (Netflix): Java-centric, used in Spring Cloud.
- **Kubernetes**: Built-in service DNS (`payment-service.default.svc.cluster.local`).

---

### Q4: What is the difference between synchronous and asynchronous communication in distributed systems?

**Answer:**

| Feature | Synchronous | Asynchronous |
|---|---|---|
| Coupling | Tight (caller waits) | Loose (caller continues) |
| Availability | Chain failure — if B fails, A fails | Producer continues even if consumer down |
| Latency | Added by each hop | Decoupled latency |
| Complexity | Simple (request-response) | Complex (message tracking, retry logic) |
| Use case | Simple reads, real-time queries | Order processing, notifications, event streaming |
| Examples | REST API, gRPC | Kafka, RabbitMQ, SQS |

**Synchronous example:**
```
Client → Payment Service → (waits) → Inventory Service
       ← payment result ←────────────┘
```
If Inventory Service is slow/down → Client times out.

**Asynchronous example:**
```
Client → Payment Service → Kafka (order_placed event)
       ← 202 Accepted (immediately)

Inventory Service → (later) → consumes order_placed → reserves stock
```

**Saga Pattern**: Breaking a distributed transaction into a sequence of local transactions coordinated by events (async). Each step publishes an event. If a step fails, compensating events undo previous steps.

---

### Q5: What is a health check? What is the difference between liveness and readiness probes?

**Answer:**
**Health checks** allow load balancers, container orchestrators (Kubernetes, ECS), and service registries to know if an instance is healthy and should receive traffic.

**Liveness Probe:**
- "Is this service alive (not deadlocked/crashed)?"
- If **fails** → container is **restarted**.
- Example: `/health/live` returns 200 if the process is running and not deadlocked.
- Should be a simple check — avoid complex DB calls.

**Readiness Probe:**
- "Is this service READY to serve traffic?"
- If **fails** → container is **removed from load balancer** (but NOT restarted).
- Allows graceful handling during: startup initialization, loading data, DB migrations completing.
- Example: `/health/ready` checks if DB connection pool is initialized.

**Startup Probe (Kubernetes):**
- "Has the application started yet?"
- Gives slow-starting applications time to initialize before liveness kicks in.

**Example Kubernetes configuration:**
```yaml
livenessProbe:
  httpGet:
    path: /health/live
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 10
  failureThreshold: 3

readinessProbe:
  httpGet:
    path: /health/ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
```

---

*Prepared for technical screening at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, Tech Mahindra).*
