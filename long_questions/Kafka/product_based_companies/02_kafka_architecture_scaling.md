# 🏗️ Kafka — Product-Based Companies Architecture & Scaling

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Uber, Netflix, LinkedIn, Hotstar

---

## Q1. How would you design a Kafka architecture to support 100k+ Transactions Per Second (TPS) across multiple regions?

"To handle such immense scale with geographical resilience, the design requires **Active-Active multi-region replication** alongside proper cluster sizing.

**1. Sizing and Partitioning:**
For 100k+ TPS, a single broker or partition will hit disk I/O and network limits. We must heavily partition the topics. A general rule of thumb is a single partition on a standard SSD can handle 5-10MB/s of sequential writes. We would provision numerous brokers (e.g., 6-12) and partition the topic into hundreds of partitions to parallelize the load.

**2. Multi-Region Replication:**
Kafka natively replicates within a single datacenter (AZ), but for multi-region, we use **MirrorMaker 2** (or Confluent Cluster Linking). 
We deploy independent Kafka clusters in US-East and US-West. Applications in US-East deal exclusively with the US-East cluster for low latency. MirrorMaker 2 continuously streams messages asynchronously from US-East to US-West and vice-versa. 

To avoid infinite replication loops, MirrorMaker 2 isolates topics by prefixing them (e.g., messages produced in US-East appear in US-West under the topic `us-east.transactions`)."

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Uber, Netflix — testing ability to handle global data localization and cross-region disaster recovery.

#### Indepth
**Handling Split-Brain:** In an active-active setup, preventing data conflicts relies heavily on how consumers aggregate. A global downstream service would consume from a regex pattern `.*\.transactions` to merge streams from all regions logically while retaining exact origin timestamps to handle out-of-order anomalies.

---

## Q2. Explain how Kafka fits into the CQRS (Command Query Responsibility Segregation) pattern.

"CQRS separates the data mutation operations (Commands) from the data retrieval operations (Queries) to allow independent scaling and distinct data models. Kafka is the perfect bridge for this architecture.

When an application receives a 'Command' (e.g., `UpdateUserProfile`), it doesn't update a central monolithic database. Instead, it validates the request and publishes a `ProfileUpdatedEvent` into a highly durable Kafka topic. This acts as the single source of truth (Event Sourcing).

Various independent 'Query' services act as consumers of this topic. One query service might be a search engine (like Elasticsearch) that indexes the profile for fast text searches. Another query service might update a Redis cache for instant user lookups. They read the event from Kafka and update their own local read-optimized databases. If one database blows up, we can just reset its Kafka offset to 0 and perfectly reconstruct it from the immutable log."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Flipkart, Hotstar — evaluating architectural chops extending beyond simple CRUD apps.

#### Indepth
**Event Sourcing vs. Kafka Retention:** If relying entirely on Event Sourcing, the Kafka topic's `retention.ms` constraint must be set to `-1` (infinite), effectively turning Kafka into the permanent system of record instead of just an ephemeral transit layer.

---

## Q3. When should you fundamentally AVOID using Kafka?

"While Kafka is incredibly robust, it is an anti-pattern to use it in certain scenarios:

**1. Point-to-Point Task Queuing:** If you just need a standard queue to farm out worker tasks (like sending emails) where messages need to be deleted upon read and complex routing is required, a standard message broker like **RabbitMQ** or AWS SQS is significantly more appropriate and easier to maintain.

**2. Strictly Synchronous Request-Reply:** If Service A needs an immediate answer from Service B to return a 200 OK to the frontend, forcing Kafka in the middle creates unnecessary complexity. Use gRPC or REST.

**3. Low Volume / Low Budget:** Operating a resilient Kafka cluster (even managed) is complex. You need ZooKeeper/KRaft quorum, minimum 3 brokers for replication, and dedicated DevOps. For trivial throughput, simple databases (like Postgres Listen/Notify) or Redis Pub/Sub suffice."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Google — Product companies love asking candidates *not* to use the tool to test for 'resume-driven development'.
---
