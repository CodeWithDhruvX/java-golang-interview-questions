# 🏗️ Architecture Fundamentals — Questions 1–15

> **Level:** 🟢 Junior – 🟡 Mid
> **Asked at:** Service companies (TCS, Infosys, Wipro) as core theory, product companies (Amazon, Flipkart, Swiggy) in HLD rounds.

---

### 1. What is software architecture?

"Software architecture is the **high-level structure of a software system** — the set of decisions about how components are organized, how they communicate, and what constraints guide the system's evolution over time.

When I'm asked to 'architect a payment system', I'm not writing code. I'm deciding: should this be a monolith or microservices? What database? How do services talk? What happens when one fails?

I think of architecture in three layers: **structure** (how is the system physically divided?), **behaviour** (how do parts collaborate at runtime?), and **cross-cutting concerns** (security, logging, observability). Getting these wrong early costs 10x to fix later."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys, Wipro (as theory) | Amazon, Flipkart (as intro to HLD rounds)

#### Indepth
Architecture is not just about technology choices. It encompasses:
- **Non-functional requirements (NFRs):** Performance, scalability, availability, security, maintainability
- **Architecture drivers:** Business requirements, team size, deployment constraints, tech ecosystem
- **Trade-off analysis:** Every architectural decision trades one quality for another. REST is simpler than gRPC but slower. Microservices are more scalable than monoliths but operationally complex.

The IEEE definition: *"Software architecture is the fundamental organization of a system embodied in its components, their relationships to each other and to the environment, and the principles guiding its design and evolution."*

#### 🗣️ How to Explain in Interview
**Interviewer:** What is software architecture?
**Your Response:** "Software architecture is the collection of **significant, hard-to-change structural decisions** that define a system's evolution. It's not just about drawing boxes and arrows; it's about defining how components are organized and how they interact under specific constraints.

I view architecture through three primary lenses: **structure** (the physical decomposition), **behavior** (runtime collaborations), and **cross-cutting concerns** like security and observability. The ultimate goal is to minimize the cost of change over the system's lifetime by making the right 'load-bearing' decisions early on, such as choosing between a monolith and microservices based on business drivers."

---

### 2. What is a monolithic architecture?

"A monolith is an application where **all components are packaged and deployed as a single unit**. The UI, business logic, and database layer are all bundled together, and you deploy one artifact.

Early-stage companies almost always start monolithic — it's simpler to develop, test, and deploy. No network hops between components. The original Instagram was a Django monolith serving millions of users before they broke it apart.

The problem comes at scale: a change to the notification code requires redeploying the entire payment system. One buggy module can take down everything. The team size that can efficiently work on one codebase has a hard ceiling."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys, Wipro | Swiggy, Zomato (in context of 'when did you migrate from monolith?')

#### Indepth
Types of monoliths:
- **Single-process monolith:** One process, one deployable. Classic.
- **Modular monolith:** Single deployment but with strict internal module boundaries. A good middle ground — Netflix calls this 'microservices discipline inside a monolith'.
- **Distributed monolith:** Multiple services but tightly coupled (shared database, synchronous chains). This is the **worst of both worlds** — complexity of microservices, brittleness of monolith.

Monolith strengths: Simple to develop locally, easy end-to-end testing, low operational overhead, direct function calls (no serialization).

When a monolith becomes painful: **Conway's Law** — organization structure starts reflecting the code structure. When 10+ teams work on one codebase, you have merge conflicts, slow CI/CD, and deployment coupling.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a monolithic architecture?
**Your Response:** "A Monolith is an architectural style where all components—UI, business logic, and data access—are packaged and deployed as a **single unit or artifact**. It’s the industry standard for early-stage products because it's simpler to develop, test, and deploy without the overhead of network communication.

However, as the organization grows, the monolith often becomes a 'Big Ball of Mud.' A change in a minor module like 'Logging' can force a redeployment of the entire 'Payments' engine, and a single memory leak in one module can take down the whole application. While it offers speed at the start, it eventually leads to **deployment coupling** and shared-resource bottlenecks that limit team velocity."

---

### 3. What is a microservices architecture?

"Microservices is an architectural style where an application is built as a **collection of small, independently deployable services**, each owning its own data and communicating over a network.

The key word is **independently deployable**. If I can't deploy Service A without coordinating with Service B, they're not truly independent microservices — they're a distributed monolith in disguise.

At Uber, there are 2,000+ microservices. The Rides service doesn't know how Payments works internally. They communicate over APIs. If Payments goes down, Rides can degrade gracefully rather than crashing entirely. That isolation is the core value."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Flipkart, Swiggy, Zomato, PhonePe, Razorpay

#### Indepth
Microservices characteristics (from Sam Newman's definition):
1. **Modeled around business domains** — one service per bounded context (not per tech layer)
2. **Smart endpoints, dumb pipes** — business logic in services, not in the network (no ESB)
3. **Culture of automation** — CI/CD pipeline per service is mandatory
4. **Decentralized governance** — teams choose their own tech stack
5. **Decentralized data management** — each service owns its own database
6. **Failure by design** — design for failure from day one; use circuit breakers, retries

Cost of microservices: distributed tracing complexity, network failures, eventual consistency, operational overhead (service discovery, health checks, config management). Never migrate to microservices because it's trendy — migrate because you have a team scaling or deployment coupling problem.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a microservices architecture?
**Your Response:** "Microservices is an architectural style where an application is built as a suite of small, **independently deployable services**, each modeled around a specific business Bounded Context. Unlike a monolith, each microservice owns its own data and communicates over lightweight protocols like HTTP or gRPC.

The most critical factor is 'Independent Deployability'—if you have to coordinate releases, you've likely built a distributed monolith. Companies like Uber and Netflix use this to achieve extreme **fault isolation** and team autonomy. However, it introduces 'distributed system taxes' like eventual consistency and the need for robust distributed tracing and service discovery."

---

### 4. Monolith vs Microservices — when to use which?

"My default answer: **start with a monolith, migrate when you feel the pain**. Premature microservices is a form of over-engineering.

Use a monolith when: you're a startup validating a product idea, you have a small team (< 5 engineers), the domain is not yet well-understood, or you need fast iteration.

Migrate to microservices when: different parts of the system have radically different scaling needs (the image processing service needs 10x more CPU than the auth service), different teams need to deploy independently, or the monolith has become a deployment bottleneck.

The signal: when your CI pipeline takes 45 minutes and 15 teams are blocked on each other's deployments — that's when microservices pays off."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Nearly every company that has gone through a scaling journey

#### Indepth
| Factor | Monolith | Microservices |
|--------|----------|---------------|
| Team size | Small (< 10 devs) | Large (many teams) |
| Domain clarity | Unknown/evolving | Well-understood |
| Deployment frequency | Low | High, per service |
| Operational maturity | Low | High (k8s, service mesh) |
| Data consistency | Strong (single DB) | Eventual (distributed) |
| Network latency | None (in-process) | Yes (100ms+ budget) |
| Dev experience | Simple | Complex (local env hell) |

**Martin Fowler's rule:** "Don't start with microservices. Almost all the successful microservice stories I know have started with a monolith that got too big and got broken up. Those that started directly with microservices have had a harder time."

#### 🗣️ How to Explain in Interview
**Interviewer:** Monolith vs Microservices — when to use which?
**Your Response:** "My heuristic is to always **start with a Monolith and only migrate once you feel the pain of scaling**. Premature microservices can kill a startup's iteration speed by forcing engineers to solve infrastructure problems before they even have a working product.

You should move to microservices when you encounter two things: **Scaling Mismatch** (e.g., your image processing needs 10x more CPU than your auth) or **Team Contention** (e.g., 20 teams are blocking each other's deployments). When your CI/CD pipeline starts taking 45 minutes and deployment failures become frequent due to merge conflicts, that's your signal that the complexity budget of microservices is finally worth the investment."

---

### 5. What is Service-Oriented Architecture (SOA)?

"SOA is an architectural style from the 2000s where applications expose functionality as **reusable services** that communicate through a central messaging layer, typically an **Enterprise Service Bus (ESB)**.

SOA was the enterprise world's answer to siloed, tightly coupled systems. The idea was right — decompose into services — but the implementation got bloated. The ESB became a 'God bus' that knew too much business logic, creating a central point of failure and chokepoint.

SOA is to microservices what a Soviet-era five-year plan is to a startup culture — heavy governance, centralization, slow iteration. Microservices is SOA done right, with smart endpoints, lightweight protocols (HTTP/gRPC), and no ESB."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Service-based companies (TCS, Infosys) transitioning to cloud

#### Indepth
| Aspect | SOA | Microservices |
|--------|-----|---------------|
| Communication | ESB (heavy, centralized) | Lightweight (REST, gRPC, Kafka) |
| Service size | Large, coarse-grained | Small, fine-grained |
| Data | Often shared DB | Each service owns its DB |
| Governance | Centralized, heavy | Decentralized, teams own |
| Deployment | Often shared deployment | Independent per service |
| Reuse focus | Enterprise reuse | Bounded context isolation |

SOA's ESB (IBM MQ, TIBCO, Oracle Service Bus) vs modern microservices message brokers (Kafka, RabbitMQ): The difference is that the ESB contained *routing logic and transformations*, making it a bottleneck. Kafka is just a dumb pipe — the transformation logic lives in the consumer.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is Service-Oriented Architecture (SOA)?
**Your Response:** "Service-Oriented Architecture (SOA) was the enterprise precursor to microservices in the early 2000s. It focused on making services reusable across a company, typically using a heavy **Enterprise Service Bus (ESB)** for routing and transformations.

The main issue was that the ESB became a 'God Object'—it was too smart and centralized, creating a massive single point of failure and coordination bottleneck. Microservices is essentially **'SOA done right'** by favoring 'Smart Endpoints and Dumb Pipes.' We strip the logic out of the network—using lightweight brokers like Kafka or gRPC—and put it back into the services to allow for faster, decentralized updates."

---

### 6. What are architectural patterns? Name the main ones.

"Architectural patterns are **reusable solutions to recurring architecture problems** — the same way design patterns solve recurring code problems.

The ones I use most: **Layered** (presentation → business → data — the classic MVC stack), **Event-Driven** (producers, events, consumers — used in Kafka-based systems), **Microservices** (domain services communicating over APIs), **CQRS** (separate read and write paths), **Event Sourcing** (store events not state), and **Space-Based** (for ultra-high-throughput like trading systems).

The key is matching the pattern to the problem. A CQRS/Event Sourcing setup for a simple CRUD blog is massive overkill."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Razorpay, PhonePe, Flipkart

#### Indepth
Core patterns:
1. **Layered (N-Tier):** Presentation → Service → Repository → Database. Simple, widely understood. Problem: changes often cut across all layers. Good for: internal tools, admin dashboards.
2. **Event-Driven:** Publisher → Event Bus → Subscriber. Highly decoupled, async. Problem: hard to trace flow, eventual consistency. Good for: order processing, notifications.
3. **Microkernel (Plugin):** Core system + pluggable extensions. Good for: IDEs, CMS systems that need extensibility.
4. **Space-Based:** Shared memory grid with no central DB. Good for: low-latency, high-throughput (trading platforms, gaming).
5. **Pipeline:** Data flows through a series of processing stages. Good for: ETL, build systems, ML pipelines.
6. **CQRS:** Commands modify state, Queries read state — separate models. Good for: high-read systems where query optimization conflicts with write optimization.

#### 🗣️ How to Explain in Interview
**Interviewer:** What are architectural patterns? Name the main ones.
**Your Response:** "Architectural patterns are high-level blueprints that provide reusable solutions to recurring design problems. For simple systems, the **Layered (N-Tier)** pattern is the baseline. For high-scale, decoupled systems, we look at **Event-Driven** architectures using Kafka or RabbitMQ.

Other critical patterns include **CQRS** (segregating read and write models), **Event Sourcing** (storing changes, not current state), and **Hexagonal (Ports and Adapters)** for domain isolation. The goal isn't to pick the 'trendiest' pattern but the one that aligns with your NFRs—using a space-based architecture for a simple CRUD app is overkill, just as a plain layered monolith won't survive a global trading platform's load."

---

### 7. What is the difference between architecture and design?

"Architecture is **what you can't easily change** — the structural decisions that are expensive to reverse. Design is **how you implement** within those constraints.

If I decide to use a microservices architecture with REST APIs and PostgreSQL, that's architecture. How I structure my Go packages, what naming conventions I use, which design patterns I apply within a service — that's design.

Architecture = load-bearing decisions. Design = implementation details within those decisions. A wrong design is a bug. A wrong architecture is a rewrite."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Senior roles at Amazon, Google, Uber

#### Indepth
The **Grady Booch** distinction: Architecture represents the significant decisions — significant by the cost of change.

Architecture decisions examples:
- "We will use microservices" — expensive to change later
- "We will use PostgreSQL as primary DB" — significant migration cost
- "Services communicate via async events over Kafka" — changing to sync would require touching every service

Design decisions examples:
- "I'll use the Repository pattern to abstract the DB layer"
- "I'll use a factory method to create payment processors"
- "I'll use an LRU cache in the service layer"

**Architecture fitness functions** (from *Building Evolutionary Architectures* by Ford, Parsons, Kua): Automated checks that verify the architecture's integrity over time — e.g., "no service may share its database with another service" can be enforced by a test that scans connection strings.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the difference between architecture and design?
**Your Response:** "I define architecture as the set of **'load-bearing' decisions** that are expensive and difficult to reverse once implemented. Decisions like 'Microservices vs Monolith' or choosing between 'PostgreSQL vs DynamoDB' are architectural because they dictate the system's long-term constraints.

Design, conversely, is about how you implement logic *within* those boundaries. Choosing a 'Factory' pattern to create objects or using 'Dependency Injection' within a Go service are design choices. A 'wrong' design is a bug you can refactor in a day; a 'wrong' architecture is often a multi-month rewrite or a complete project failure."

---

### 8. What is the CAP theorem?

"CAP theorem states that a distributed system can only guarantee **two out of three** properties: **Consistency, Availability, and Partition Tolerance**.

Since network partitions *always happen* in real distributed systems, you're really choosing between CA (impossible in practice), CP (consistent but may be unavailable during partition), or AP (available but may return stale data during partition).

Real example: If the Mumbai and Chennai nodes get disconnected from each other (partition), and a write hits Mumbai, Chennai can either: refuse to serve reads until partition heals (CP — consistent, less available), or serve its potentially stale data (AP — available, not consistent). DynamoDB defaults to AP. ZooKeeper is CP."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Amazon, Google, Uber, Flipkart, PhonePe

#### Indepth
Modern databases and their CAP positioning:
| Database | CAP Type | Why |
|----------|----------|-----|
| Zookeeper | CP | Refuses reads during partition to maintain consistency |
| etcd | CP | Used for leader election — must be consistent |
| Cassandra | AP | Always available, eventual consistency, tunable |
| DynamoDB | AP (tunable) | Default eventual, but strongly consistent reads available |
| MongoDB (replica) | CP | Reads from primary only to ensure consistency |
| CockroachDB | CP | Distributed SQL with serializable transactions |

The **PACELC extension** by Daniel Abadi: Even without partitions, there's a trade-off between Latency and Consistency. DynamoDB, for example, can be low-latency eventual consistent or have higher latency strongly consistent — that's the PACELC trade. More useful for real systems because partitions are rare but latency is always relevant.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the CAP theorem?
**Your Response:** "The CAP theorem states that a distributed system can only provide two out of three: Consistency, Availability, and Partition Tolerance. Because network partitions (P) are an unavoidable reality of distributed computing, the real choice is between **CP (Consistency/Partition Tolerance)** or **AP (Availability/Partition Tolerance)**.

ZooKeeper is a classic **CP** system—it will refuse to serve requests if it can't guarantee consistency during a partition. Cassandra, however, is a typical **AP** system—it prioritizes being available to the user even if some nodes are out of sync, meaning it might return slightly stale data. As an architect, you choose the model based on the business cost of staleness versus the cost of downtime."

---

### 9. What is the difference between consistency and eventual consistency?

"**Strong consistency** means after a write completes, every subsequent read — from any node — will return the new value. Like a synchronized clock.

**Eventual consistency** means after a write, reads might return the old value for a while, but eventually — once all nodes have synced — every read will return the new value.

The classic example: your Twitter feed. When you post a tweet, your followers might not see it for a couple of seconds. That's eventual consistency — Twitter chose availability (always show *something*) over strong consistency (wait until all nodes agree before showing anything). For a bank balance, you'd want strong consistency."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Flipkart, Swiggy, Razorpay, PhonePe

#### Indepth
Consistency models (from weakest to strongest):
1. **Eventual consistency:** Writes propagate asynchronously. Reads may be stale. Cassandra, DynamoDB default.
2. **Read-your-writes consistency:** You always see your own writes. Others may not. DynamoDB with read-your-own-writes option.
3. **Monotonic read consistency:** If you read value X, you'll never read an older value. Prevents time-travel confusion.
4. **Causal consistency:** Reads that causally follow a write see the write. Used in session guarantees.
5. **Sequential consistency:** All nodes see operations in same order, but not necessarily real-time.
6. **Linearizability (strong consistency):** All operations appear instantaneous and in real-time order. Most expensive. Used in ZooKeeper, CockroachDB.

**The practical rule:** Use eventual consistency for high-read, low-stakes data (product catalog, social feeds). Use strong consistency for financial transactions, inventory counts, leader election.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the difference between consistency and eventual consistency?
**Your Response:** "Strong consistency guarantees that after a write confirms, every reader—no matter where they are—sees that new value immediately. It requires expensive distributed coordination, like a 2-Phase Commit or Paxos protocol.

**Eventual consistency** is more relaxed: after a write, some users might see the old value for a brief window, but 'eventually' every node catches up. We use this for high-throughput systems like Twitter or Instagram where user experience (Availability) is more important than absolute real-time accuracy. However, for a bank balance transfer, strong consistency is a non-negotiable architectural requirement."

---

### 10. What is the SOLID principle in the context of architecture?

"SOLID principles apply at the **component and service level**, not just at the class level.

The most important for architecture: **Single Responsibility** — each service should do one thing and own one domain (the microservice decomposition principle). **Open/Closed** — add new functionality by creating new services rather than modifying existing ones. **Dependency Inversion** — high-level services should depend on abstractions (APIs), not on the specific implementation of downstream services.

When I violate SRP at the service level — one service handling orders, payments, and notifications — I get the distributed monolith problem: one change requires coordinating across all three concerns."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Flipkart, Amazon, Razorpay, all product companies

#### Indepth
SOLID at the architecture level:
- **S (SRP):** A service should have one reason to change. Violation: An Order service that also sends emails. If email template changes, I shouldn't need to touch the Order service.
- **O (OCP):** System should be open for extension, closed for modification. Implementation: Use event-driven architecture. New subscribers can be added without touching the publisher.
- **L (LSP):** Services should be substitutable. Violation: A mock payment service that behaves differently from the real one — this breaks contract testing.
- **I (ISP):** Don't force consumers to depend on interfaces they don't need. Implementation: BFF (Backend for Frontend) pattern — each client gets a tailored API.
- **D (DIP):** Depend on abstractions. Implementation: Define service contracts via OpenAPI specs or protobuf. Services don't depend on each other's internals, only on the contract.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the SOLID principle in the context of architecture?
**Your Response:** "SOLID principles provide a roadmap for avoiding a 'distributed monolith.' **Single Responsibility (SRP)** at the service level means each service should own exactly one business Bounded Context, such as 'Payments' or 'Inventory.' 

**Open/Closed (OCP)** is best achieved through Event-Driven architectures, where you can extend system behavior by adding new subscribers without modifying the publisher's code. **Dependency Inversion (DIP)** is implemented by services depending on stable API contracts (OpenAPI/Protobuf) rather than their internal implementations. Following macro-level SOLID ensures that your services remain decoupled and independently evolvable."

---

### 11. What is Conway's Law and how does it affect architecture?

"Conway's Law states: **'Organizations which design systems are constrained to produce designs which are copies of the communication structures of those organizations.'**

In practice: if you have separate frontend, backend, and DBA teams, you'll end up with a classic 3-tier architecture even if microservices would serve you better — because that's how the teams communicate. The team structure becomes the architecture.

The Amazon two-pizza team rule directly addresses this. Jeff Bezos mandated that every team must be small enough to be fed by two pizzas, and every team owns one service end-to-end. Result: microservices architecture naturally emerged from the org structure."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Google, Amazon, Senior Engineering Manager roles

#### Indepth
The **Inverse Conway Maneuver**: Instead of letting Conway's Law happen to you, deliberately design your team structure to produce the architecture you want.

Example: If you want to split the monolith into a Checkout service and a Product service, first create two separate teams — the Checkout team and the Product team. The architecture change will follow naturally because teams will communicate via API contracts rather than shared code.

**Platform teams vs stream teams** (Team Topologies book): Stream-aligned teams own end-to-end business flow (checkout, search). Platform teams provide internal platforms (logging, CI/CD, observability). Enabling teams help other teams adopt new tech. This organizational structure maps directly to the architecture pattern of microservices + platform.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is Conway's Law and how does it affect architecture?
**Your Response:** "Conway's Law states that a system's architecture will inevitably mirror the communication patterns of the organization that built it. If you have four separate teams, you'll likely end up with four main modules or services.

As architects, we use the **'Inverse Conway Maneuver'** to our advantage. If we want a decoupled, microservices-based system, we deliberately structure our teams into small, cross-functional units (like Amazon's **'Two-Pizza Teams'**) that each own a single service 'soup to nuts.' By changing the organizational boundaries first, we set the stage for the architectural boundaries we want to see emerge in the code."

---

### 12. What is the twelve-factor app methodology?

"The Twelve-Factor App is a methodology for building **cloud-native, scalable, maintainable web applications**. It was published by Adam Wiggins from Heroku and is the foundational checklist for modern SaaS architecture.

The three most important factors to me: **Config stored in environment** (not hardcoded or in files — this enables the same code to run in dev, staging, and prod), **Processes are stateless** (any state goes in a backing store like Redis — this enables horizontal scaling), and **Logs are event streams** (write to stdout, let the platform aggregate — this enables observability at scale).

When I review a service's architecture, I use the 12 factors as a checklist."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Swiggy, Razorpay, Zepto, Groww, all cloud-first companies

#### Indepth
All twelve factors:
1. **Codebase:** One codebase tracked in VCS, many deploys
2. **Dependencies:** Explicitly declare and isolate dependencies (`go.mod`, `package.json`)
3. **Config:** Store config in environment variables (not code)
4. **Backing Services:** Treat DB, cache, queue as attached resources (swap without code change)
5. **Build, Release, Run:** Strict separation — build once, release to any env
6. **Processes:** Execute app as stateless processes
7. **Port Binding:** Export services via port binding (no app server dependency)
8. **Concurrency:** Scale out via the process model
9. **Disposability:** Fast startup, graceful shutdown; support instant scale up/down
10. **Dev/Prod Parity:** Keep development, staging, production as similar as possible
11. **Logs:** Treat logs as event streams (stdout only)
12. **Admin Processes:** Run admin/management tasks as one-off processes

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the twelve-factor app methodology?
**Your Response:** "The Twelve-Factor App is the gold standard for building cloud-native SaaS applications. It’s a methodology designed to maximize portability and scalability while minimizing the 'works on my machine' syndrome.

The pillars for me are: **Config in the Environment** (decoupling code from its environment), **Stateless Processes** (storing any state in backing services like Redis to enable horizontal scaling), and **Logs as Event Streams** (writing to stdout/stderr and letting the platform handle aggregation). Adhering to these 12 factors ensures that your application is resilient, easy to automate, and ready for deployment on platforms like Kubernetes or AWS."

---

### 13. What are non-functional requirements (NFRs) and why do they matter in architecture?

"Non-functional requirements define **how the system performs**, not what it does. They're the '-ilities': scalability, availability, reliability, security, maintainability, observability.

NFRs drive architectural decisions more than functional requirements do. 'Show a product page' is a functional requirement — almost any architecture can handle it. 'Show a product page in < 200ms for 5 million concurrent users with 99.99% uptime' — that's where architecture gets hard.

I always start an architecture discussion by surfacing the NFRs. If the requirement is 100K RPM with < 50ms p99, that's a very different architecture than 100 RPM with < 500ms p99."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Amazon, Google, Flipkart, senior design rounds

#### Indepth
Key NFRs and their architectural implications:
| NFR | Implication |
|-----|-------------|
| **Availability (99.99%)** | Active-active multi-region, health checks, failover |
| **Latency (p99 < 100ms)** | CDN, caching, async processing, co-location |
| **Throughput (1M RPM)** | Horizontal scaling, load balancing, DB read replicas |
| **Data consistency** | Sync vs async communication, ACID vs BASE |
| **Security** | Zero-trust, encryption at rest/transit, RBAC |
| **Maintainability** | Service isolation, well-defined contracts, ADRs |
| **Observability** | Distributed tracing, centralized logging, metrics |
| **Scalability** | Stateless services, sharding, autoscaling |

**SLA, SLO, SLI relationship:** SLI (indicator: actual measurement, e.g. "p99 latency = 95ms") → SLO (objective: internal target, e.g. "p99 latency < 100ms, 99.9% of the time") → SLA (agreement: contractual commitment with penalties, e.g. "p99 < 200ms, 99.5% of the time"). Architecture decisions are made to meet SLOs.

#### 🗣️ How to Explain in Interview
**Interviewer:** What are non-functional requirements (NFRs) and why do they matter in architecture?
**Your Response:** "Functional requirements tell you what a system does, but **Non-Functional Requirements (NFRs)**—like scalability, availability, and low latency—tell you how it must perform. In reality, NFRs are what actually drive 80% of our architectural decisions.

While any architecture can 'place an order,' placing 10,000 orders per second with **99.99% availability** and under **200ms p99 latency** requires complex structural choices like sharding, active-active replication, and asynchronous processing. I always start every design session by quantifying these 'ilities' because building for 1,000 users is a trivial task, but building for 100 million users is an architectural challenge."

---

### 14. What is a layered (N-tier) architecture?

"Layered architecture organizes the system into **horizontal layers**, each with a specific responsibility, and each layer only communicates with the layer directly below it.

The classic 3-tier: **Presentation layer** (React/Angular — renders the UI), **Application layer** (Spring Boot/Go — business logic), **Data layer** (PostgreSQL/MongoDB — persistence). A request flows top-to-bottom, response flows bottom-to-top.

It's the architecture I use for internal tools, admin dashboards, and CRUD applications. It's familiar, well-understood, and easy to onboard new developers onto."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys, Wipro, Cognizant

#### Indepth
Layer variants:
- **3-tier:** Presentation → Business → Data (most common)
- **4-tier:** Presentation → API Gateway → Business → Data
- **Hexagonal (Ports & Adapters):** Business logic at the center, external concerns (DB, UI, queues) plug in via adapters. Better than classic layered because the business core doesn't depend on frameworks.
- **Onion architecture:** Domain model at center, surrounded by domain services, application services, infrastructure. Dependencies always point inward.

Problems with classic layered:
- **Sinkhole anti-pattern:** Request passes through all layers but only one layer does real work.
- **Cross-cutting concerns:** Features that span multiple layers (logging, auth) get duplicated.
- **Coupling to framework:** Business logic often gets contaminated by Spring annotations or ORM entities.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a layered (N-tier) architecture?
**Your Response:** "Layered architecture is the 'de facto' way to organize most applications by grouping code into horizontal levels like Presentation, Business, and Data. It follows the **'Principle of Separation of Concerns,'** making it very intuitive for new developers.

It's my go-to for standard business apps, internal tools, or CRUD services where the logic isn't hyper-complex. However, you have to be careful about the **'Sinkhole Anti-pattern,'** where requests pass through all four layers but only one does real work, and avoid 'Tight Coupling' where your business logic becomes accidentally dependent on a specific database framework."

---

### 15. What is hexagonal architecture (Ports and Adapters)?

"Hexagonal architecture puts the **business logic at the center**, completely isolated from the outside world. Everything external — HTTP requests, database calls, message queue publishing — connects through well-defined **ports** (interfaces) via **adapters** (implementations).

In code: the `Order` domain object doesn't import `database/sql` or `gin`. It only knows about its own rules and interfaces. A `PostgresOrderRepository` adapter implements the `OrderRepository` interface to connect the domain to the actual DB.

The huge benefit: I can test the entire business logic with no database, no HTTP server, no external dependencies. Just inject mock adapters. I can also swap PostgreSQL for MongoDB without touching a line of business logic."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Zepto, Groww, Razorpay, Cred — companies that care about clean code

#### Indepth
Core concepts:
- **Port:** An interface that represents how the application interacts with the outside world. Two types:
  - *Primary (driving) port:* How the outside world triggers the application (HTTP endpoint, CLI, message consumer)
  - *Secondary (driven) port:* How the application triggers the outside world (DB, email service, payment gateway)
- **Adapter:** The implementation of a port. Primary adapter: REST controller, gRPC handler. Secondary adapter: SQL repository, Stripe HTTP client.

```
[HTTP Adapter] → |Primary Port| → [Application Core] → |Secondary Port| → [DB Adapter]
[gRPC Adapter] → |Primary Port| →         |           → |Secondary Port| → [Redis Adapter]
```

The business core is **framework-agnostic**: it doesn't know if it's being called via HTTP or a CLI. This is why teams like Netflix can run the same business logic in a request-response context AND in a batch processing context without modification.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is hexagonal architecture (Ports and Adapters)?
**Your Response:** "Hexagonal architecture, or Ports and Adapters, aims to put the **Business Logic at the absolute center**, shielding it from external frameworks like databases or HTTP servers. You define 'Ports' as interfaces and 'Adapters' as their concrete implementations.

For instance, your domain logic talks to an `OrderRepository` interface (the Port) without knowing if the actual storage is PostgreSQL, MongoDB, or a mock in-memory store (the Adapter). This makes the system **completely testable in isolation** and incredibly resilient to technical changes. It's the gold standard for building clean, maintainable microservices that can outlast any specific database or cloud provider."
