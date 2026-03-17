# 🗄️ Data Architecture — Questions 1–10

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Amazon, Flipkart, Google, Groww, Zepto — data-intensive systems and senior design rounds

---

### 1. What is polyglot persistence?

"Polyglot persistence is the idea of **using different database technologies for different data needs** within the same application, rather than fitting all data into a single database type.

Classic example: A Swiggy-scale application might use PostgreSQL for structured transactional data (orders, users), Redis for sessions and caching, Elasticsearch for restaurant search (full-text), Cassandra for driver location history (time-series writes), and S3 for dish photos (blob storage). Each database is chosen because it excels at the specific access pattern.

The tradeoff: operational complexity goes up (more DBs to manage, monitor, backup). This is why smaller teams start with one good relational DB and add specialized stores only when they hit actual limits."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Swiggy, Zomato, Flipkart, any data-heavy company

#### Indepth
Database type → use case mapping:
| Database Type | Example | Best For |
|---------------|---------|----------|
| Relational (RDBMS) | PostgreSQL, MySQL | Structured data, transactions, complex queries, joins |
| Key-Value | Redis, DynamoDB | Sessions, caching, rate limiting, shopping carts |
| Document | MongoDB, Firestore | Semi-structured data, user profiles, catalogs |
| Wide-Column | Cassandra, HBase | Time-series, write-heavy, large scale (IoT, logs) |
| Graph | Neo4j, Amazon Neptune | Relationship traversal (social networks, fraud detection) |
| Search | Elasticsearch, Solr | Full-text search, log analytics, autocomplete |
| Time-Series | InfluxDB, TimescaleDB | Metrics, monitoring, financial tick data |
| Blob/Object | S3, GCS | Files, images, videos, ML models |
| NewSQL | CockroachDB, Spanner | Globally distributed SQL with strong consistency |

Choosing the right DB is one of the most consequential architectural decisions. PostgreSQL is often the right starting point — it's battle-tested, ACID-compliant, and can be extended with plugins for JSON, geospatial, and time-series.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is polyglot persistence?
**Your Response:** "Polyglot persistence is the practice of using multiple database technologies within the same application, each chosen for the specific access pattern it handles best. Instead of trying to force every piece of data into a single SQL database—which can lead to poor performance or overly complex schemas—we pick the 'right tool for the job.'

A classic example for a large e-commerce platform would be using **PostgreSQL** for transactional data like orders and user profiles, **Redis** for sub-millisecond session storage and rate limiting, and **Elasticsearch** for high-performance full-text search across the product catalog. The main trade-off is **operational complexity**—you now have more systems to monitor, back up, and secure—so I generally advise starting with a solid relational DB like Postgres and only adding specialized stores when you hit a clear technical bottleneck."

---

### 2. What is database replication and what are the types?

"Database replication is the process of **copying data from one database server (primary) to one or more others (replica)** to improve availability, read performance, and fault tolerance.

Three main types: **Leader-Follower (Master-Replica)** — all writes go to leader, reads can be distributed across replicas. Most common. Used by MySQL, PostgreSQL. **Leader-Leader (Multi-Master)** — writes accepted on any node, complex conflict resolution required. Used by CockroachDB, Cassandra. **Leaderless** — any node accepts writes (like Cassandra, DynamoDB) using quorum-based consistency.

I use leader-follower replication for read-heavy systems. The rule of thumb: if you have 80% reads and 20% writes, adding a read replica can double your read throughput without touching the application code."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Flipkart, Swiggy — any scaling discussion

#### Indepth
Replication strategies (how data is copied):
1. **Synchronous replication:** Primary waits for at least one replica to confirm write before acknowledging the client. Stronger consistency, higher latency.
2. **Asynchronous replication:** Primary acknowledges write immediately, propagates to replicas in background. Lower latency, risk of data loss on primary failure.
3. **Semi-synchronous:** Wait for one replica (RPO = near-zero data loss), rest are async. MySQL's default recommended mode.

Replication lag: In async replication, replicas may lag behind the primary by milliseconds to seconds. Reading from a replica immediately after a write may return stale data. Solutions: Route reads to primary for the session immediately after a write (read-your-own-writes), wait for replica to catch up, or use synchronous replication for critical reads.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is database replication and what are the types?
**Your Response:** "Database replication is the process of keeping multiple copies of the same data across different servers to ensure high availability and improve read performance. The standard pattern I use is **Leader-Follower (Master-Replica)**, where all writes are directed to a primary node and then propagated to one or more replicas. 

We can choose between **Synchronous replication**, which ensures zero data loss but adds a latency penalty to every write, or **Asynchronous replication**, which is much faster but introduces a small 'replication lag.' In a read-heavy system, adding replicas is a strategic way to scale—you can offload all the heavy analytics or read traffic to the followers, keeping the primary leader free to handle the critical business transactions."

---

### 3. What is database sharding and how do you choose a shard key?

"Sharding is **horizontal partitioning** — distributing data across multiple database nodes so each node holds a subset. When a single DB can't handle the write volume or storage, sharding is the solution.

Choosing the shard key is the most critical decision. The shard key determines which node a record lives on. A bad shard key creates **hotspots** — one shard gets all the traffic while others sit idle. A good shard key distributes data and traffic uniformly.

My criteria for a shard key: (1) High cardinality (enough distinct values to distribute evenly), (2) Uniform distribution (not 90% of orders from Mumbai), (3) Query alignment (most queries include the shard key, avoiding cross-shard scatter-gather), (4) Immutable (once assigned, never changes)."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, PhonePe, Twitter — any system at extreme scale

#### Indepth
Shard key examples:
- **userId:** Good for social/commerce apps where most queries are user-scoped
- **Geographic region:** Good for location-based services, but risk of regional hotspots
- **Hash(userId):** Even distribution, but range queries are impossible
- **Composite key (date + userId):** Good for time-ordered user data; prevents temporal hotspots

Sharding pitfalls:
1. **Cross-shard joins:** `JOIN` across shards requires scatter-gather (ask all shards, merge results). Very expensive. Solution: avoid joins across shard boundaries by design.
2. **Cross-shard transactions:** ACID transactions can't span shards without 2PC or saga patterns. Accept eventual consistency.
3. **Resharding:** When you add a new shard, you need to redistribute data. Hash-based sharding requires re-hashing all data. **Consistent hashing** minimizes data movement — only 1/N of data moves on adding 1 shard to an N-shard cluster.
4. **Hotspot on sequential IDs:** If shard key is an auto-increment ID, all writes go to the shard with the highest IDs. Solution: use ULIDs or hash-based IDs.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is database sharding and how do you choose a shard key?
**Your Response:** "Sharding is the process of horizontally partitioning your database when a single server can no longer handle the write volume or the total data size. The most critical part of this architecture is choosing the **Shard Key**. 

A good shard key needs to have **high cardinality** and an even distribution—for example, a `User_ID` or a `Tenant_ID`. If you pick a bad key, like 'Country,' you might end up with a 'hotspot' where one shard handles 90% of the traffic while the others are idle. Sharding is a high-complexity move because it makes joins and transactions across shards very difficult, so I always treat it as a last resort once we've exhausted other options like vertical scaling or read replicas."

---

### 4. What is data lake vs data warehouse vs data mart?

"Three different data storage paradigms for analytical workloads:

**Data Warehouse:** Structured, schema-on-write, optimized for SQL analytics. Data is cleaned, transformed, and loaded (ETL) before storage. Queries are fast because schema is known. Amazon Redshift, Google BigQuery, Snowflake. Good for: business intelligence dashboards, financial reports.

**Data Lake:** Raw, schema-on-read, stores everything in original format. Data is loaded as-is (CSV, JSON, Parquet, images, logs). Schema is applied when querying. Amazon S3 + Athena, Azure Data Lake, GCS. Good for: data science, ML model training, exploratory analysis.

**Data Mart:** A subset of a warehouse focused on a specific business unit (marketing data mart, finance data mart). Pre-aggregated for specific query patterns."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Data engineering roles, senior backend roles at analytics-heavy companies

#### Indepth
Modern evolution — **Data Lakehouse:**
Combines the raw storage of a data lake with the structured querying of a data warehouse. Delta Lake (Databricks), Apache Iceberg, Apache Hudi enable ACID transactions and schema enforcement on top of object storage (S3).

| Aspect | Data Warehouse | Data Lake | Data Lakehouse |
|--------|---------------|-----------|----------------|
| Schema | Schema on write | Schema on read | Both (flexible) |
| Data type | Structured | Any (raw) | Any |
| Query speed | Fast | Slower (scan raw files) | Fast (optimized with indexing) |
| Use case | BI, reports | Data science, ML | Both |
| Cost | Expensive | Cheap (S3 storage) | Medium |
| Tool | Redshift, BigQuery | S3+Athena, GCS | Delta Lake, Iceberg |

At Flipkart: Raw clickstream events go to S3 (data lake). Overnight Spark jobs transform and load into Redshift (data warehouse). Marketing analysts query Redshift for daily reports. Data scientists query the raw S3 lake for ML model training.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is data lake vs data warehouse vs data mart?
**Your Response:** "These paradigms serve different analytical needs. A **Data Warehouse**, like Redshift or BigQuery, is a highly structured environment built for fast building of business intelligence reports—the data is cleaned and transformed *before* it gets stored. A **Data Lake** is more of a 'schema-on-read' system where you dump raw data in its original format, like JSON or logs, which is perfect for data scientists who need the 'full history' for machine learning.

In my recent designs, I’ve leaned toward the **'Data Lakehouse'** approach. It allows us to keep the cost-effective storage of a data lake on S3 while adding a structured layer on top using technologies like Iceberg or Delta Lake. This gives us the best of both worlds: the flexibility of a lake and the query performance of a warehouse."

---

### 5. What is the ETL vs ELT pattern?

"**ETL (Extract-Transform-Load):** Extract data from source → Transform/clean it → Load into destination. The transformation happens before the load, using a separate processing layer. Traditional data warehouses use ETL because you had to clean data before putting it in the expensive warehouse.

**ELT (Extract-Load-Transform):** Extract data from source → Load raw into destination → Transform inside the destination. Modern cloud warehouses (BigQuery, Snowflake) are powerful enough to do the transformation in SQL after loading. Raw data is preserved, transformations are rerunnable, and the schema is flexible.

I prefer ELT for modern stacks: cheaper (raw storage on S3 is cheap), flexible (change transformation logic without re-ingesting), and auditable (raw data is always available for re-processing)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Data engineering interviews, backend roles at analytics companies

#### Indepth
Modern ETL/ELT tooling:
- **Ingestion:** Airbyte, Fivetran (200+ connectors), custom Kafka consumers
- **Transformation:** dbt (SQL-based transformation inside the warehouse), Apache Spark (large-scale distributed transform)
- **Orchestration:** Apache Airflow (DAG-based pipelines), Prefect, Dagster
- **Warehouse:** BigQuery, Snowflake, Redshift

**dbt (Data Build Tool):** Lets data analysts write SQL transforms as version-controlled models. `dbt run` compiles SQL and runs it in the warehouse. Tests are built-in (`not null`, `unique`, `referential integrity`). This is how modern analytics engineering works.

The pipeline: `Kafka → S3 (raw) → Airbyte → BigQuery (raw tables) → dbt → BigQuery (mart tables) → Looker (dashboards)`

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the ETL vs ELT pattern?
**Your Response:** "ETL is the traditional approach where you clean and transform your data *before* it ever reaches the database. It’s useful when security or data quality constraints are very strict. However, in the cloud era, I prefer **ELT (Extract-Load-Transform)**. 

With ELT, we land the raw data into a powerful cloud warehouse like Snowflake first and then use the warehouse's own massive compute power to transform it using SQL tools like **dbt**. The huge advantage of ELT is flexibility: if our business logic changes, we don't have to re-ingest all the data from the source; we just update our SQL models and rerun them on the raw data that’s already there. It's faster to build and much easier to maintain as requirements evolve."

---

### 6. What is a data mesh architecture?

"Data Mesh is a **decentralized approach to data architecture** where each business domain owns and serves its own data as a product, rather than having a central data team own all data.

The problem with centralized data platforms: the data engineering team becomes a bottleneck. Every team that needs data analytics depends on one team to build pipelines. Data quality is poor because the data team doesn't understand the domain. The platform is generic and hard to scale organizationally.

Data Mesh's answer: the Order team owns the Order data domain, provides a well-defined data product (clean, documented, SLA-backed), and is responsible for its quality. The Analytics team doesn't own pipelines — domain teams do. A centralized self-serve platform provides the infrastructure tools."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Data platform engineering, principal/staff engineer roles at large companies

#### Indepth
Data Mesh's four principles (Zhamak Dehghani):
1. **Domain ownership:** Domain teams own their data end-to-end — from production to analytics
2. **Data as a product:** Data is served as a product with SLAs, documentation, schemas, and quality guarantees
3. **Self-serve data platform:** Infrastructure team provides tools (storage, cataloging, querying) so domain teams can serve data without each building their own infrastructure
4. **Federated computational governance:** Global policies (compliance, security) are enforced centrally but applied locally via automation

Companies using Data Mesh: Zalando (one of the earliest adopters), Netflix (domain-oriented data products), ThoughtWorks.

vs. Data Lakehouse: A data mesh is an **organizational approach** (who owns what). A data lakehouse is a **technical approach** (how data is stored and queried). They're complementary — a data mesh can use a lakehouse as the underlying storage.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a data mesh architecture?
**Your Response:** "Data Mesh is more of an **architectural pattern for the organization** than just a specific piece of software. It challenges the traditional model where a single 'Central Data Team' owns all the pipelines and data quality. In a large company, that central team always becomes a bottleneck because they don't understand the specific business context of the data they're moving.

In a Data Mesh, we treat **data as a product**. The team that builds the 'Orders' microservice is also responsible for serving a clean, documented analytical version of that order data to the rest of the company. It decentralizes the ownership so that the people who *produce* the data are the ones who *govern* it. My role as an architect in this setup is to ensure we have a shared, self-serve platform so that these teams don't have to reinvent the wheel for storage or security."

---

### 7. What is Change Data Capture (CDC) and when do you use it?

"CDC (Change Data Capture) is a pattern that **captures every change (INSERT, UPDATE, DELETE) from a database** and makes those changes available as a stream of events for downstream consumers.

Most databases write every change to a binary log (MySQL binlog, PostgreSQL WAL). CDC tools (Debezium, AWS DMS) read this log and publish changes as events to Kafka — without any application code changes.

Use cases: (1) Syncing data to a read model / search index without two-phase writes, (2) Outbox pattern — Debezium reads the outbox table as changes flow in, (3) Cache invalidation — when DB record changes, invalidate the Redis cache, (4) Audit log generation, (5) Real-time data warehouse sync."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Data platform teams, fintech with audit requirements

#### Indepth
Debezium flow:
```
MySQL (binlog) → Debezium Connector → Kafka Topic `mysql.orders` → Consumers
                                        ↓                            ↓
                               Elasticsearch          Read Model DB
                               (search index)         (denormalized)
```

CDC vs triggers vs batch polling:
| Approach | Latency | Load on DB | Implementation |
|----------|---------|------------|----------------|
| DB Triggers | Near-zero | Medium | DB-level logic |
| Batch polling | Minutes | High (queries) | App-level cron |
| CDC (WAL) | Milliseconds | Low (reads existing log) | Debezium connector |

CDC advantages: Zero application code change, millisecond latency, low DB overhead (reads log that's already written). Perfect complement to the outbox pattern.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is Change Data Capture (CDC) and when do you use it?
**Your Response:** "CDC is a pattern for tracking every change that happens in your database—like inserts, updates, or deletes—and turning them into a real-time event stream. We usually implement this with a tool like **Debezium**, which reads the database's internal transaction log (like the MySQL Binlog) without adding any extra load or queries to the app.

I find CDC essential for **data synchronization**. For example, when an order is updated in our main SQL database, Debezium can instantly push that change to Kafka, which then updates our Elasticsearch search index or clears a Redis cache. It creates a 'bridge' between different systems that ensures they stay in sync within milliseconds, all without writing any complex 'dual-write' logic in the application code."

---

### 8. What is database indexing strategy at the architecture level?

"Indexing strategy at the architecture level goes beyond 'add an index to speed up a query'. It's about **understanding the access patterns at design time** and designing the schema + indexes to support them.

The principle: indexes are a trade-off — they speed up reads but slow down writes (every write must update all relevant indexes) and consume storage. Over-indexing a write-heavy table is a performance killer.

My approach: Start with queries, not the schema. List the top 10 query patterns. Design the schema and indexes to support those patterns. Common pattern: a `users` table needs indexes on `email` (login lookup), `phone` (OTP lookup), and `created_at + status` (admin dashboard queries) — that's it. Not an index on every column."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any backend role with database-heavy systems

#### Indepth
Index types and use cases:
- **B-Tree (default):** Range queries, `<`, `>`, `=`, `LIKE 'abc%'`. Most queries.
- **Hash:** Only `=` lookups. No range. Faster for equality. PostgreSQL only for in-memory.
- **GIN (Generalized Inverted Index):** Arrays, JSONB, full-text search in PostgreSQL.
- **GiST:** Geometric data, geospatial (PostGIS), full-text search.
- **Partial index:** Index only rows matching a condition. `WHERE status = 'active'`. Smaller, faster. Great for sparse conditions.
- **Covering index (include):** Index contains all columns needed by a query — query is served from index alone (no table lookup). Fastest possible read.
- **Composite index:** Multi-column index. Column order matters — leftmost prefix rule.

Composite index ordering: `(status, created_at)` supports `WHERE status = ? AND created_at > ?` and `WHERE status = ?`. Does NOT support `WHERE created_at > ?` alone (skips the first column). Always put equality conditions first, range conditions last.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is database indexing strategy at the architecture level?
**Your Response:** "Architecture-level indexing is about designing your data for its **most frequent access patterns**. Every index is a trade-off: it makes your reads significantly faster, but it slows down your writes because every index has to be synchronously updated. 

When I design a system, I analyze the 'top 10' queries. I look for opportunities to use **'Covering Indexes'**, where all the data the query needs is contained within the index itself, avoiding a second trip to the main table. I also use **'Partial Indexes'** in PostgreSQL to index only a subset of data—like only 'active' users—which keeps the index small and fast. The goal is to maximize read performance while keeping the 'write amplification' and storage costs under control."

---

### 9. What is the N+1 query problem and how do you solve it?

"The N+1 problem is a performance anti-pattern where you make **1 query to get N records, then N additional queries** to fetch related data for each record.

Classic example: Fetch 100 orders → for each order, fetch the user details → result: 1 + 100 = 101 DB queries. This is catastrophically slow and the #1 database performance bug I see in production.

Solutions: **JOIN** (fetch everything in one query), **eager loading** (ORM's `.Preload()` in GORM, `.Include()` in Hibernate — loads related data in one additional query, not N queries), **DataLoader** (batching — collect all IDs, then one `SELECT * WHERE id IN (...)`)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Backend interviews at any company using an ORM

#### Indepth
Solution approaches:
1. **SQL JOIN:**
```sql
SELECT orders.*, users.name, users.email
FROM orders
JOIN users ON orders.user_id = users.id
WHERE orders.created_at > '2024-01-01'
```
One query. Most efficient. Best when you know the access pattern upfront.

2. **ORM Eager Loading (GORM):**
```go
db.Preload("User").Find(&orders)
// Executes: SELECT * FROM orders; SELECT * FROM users WHERE id IN (1,2,...N)
// 2 queries instead of N+1
```

3. **DataLoader Pattern (GraphQL):**
```
userLoader.load(userId) // batched automatically
// Collects all userIds requested in a tick, then:
// SELECT * FROM users WHERE id IN (all collected IDs)
```

Detection: Enable SQL query logging in development. Any repeated query with different parameter values (same query template, different ID) is a potential N+1. `EXPLAIN ANALYZE` in PostgreSQL shows query execution plans and costs.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the N+1 query problem and how do you solve it?
**Your Response:** "The N+1 problem is one of the most common causes of slow applications. It happens when you fetch a list of items (1 query), and then your code loops through them and makes *another* query for each item (N queries) to get related data—like fetching 100 orders and then doing 100 lookups for the customer names.

To solve this architecturally, I use **Eager Loading** or **Batching**. For example, in GORM, I use the `Preload` function to fetch the related data in a single joined query or a fast 'IN' clause. When I'm working with GraphQL, I always use the **DataLoader pattern**, which batches all the individual IDs and fetches them in a single database call. It's about reducing the 'chattiness' between your application and your database to keep latency low."

---

### 10. What is eventual consistency vs strong consistency from a data architecture perspective?

"In distributed data systems, **strong consistency** guarantees that all nodes see the same data at the same time — a write is only considered complete when all (or a quorum of) nodes have committed it. **Eventual consistency** guarantees that if no new updates are made, eventually all nodes will converge to the same value — but reads may return stale data in the interim.

The practical choice: use strong consistency for correctness-critical data (bank balances, inventory counts, booking availability). Use eventual consistency for high-throughput, high-availability data where slight staleness is acceptable (social feeds, product recommendations, analytics counters).

A payment system must be strongly consistent (double-spend prevention). A news feed can be eventually consistent (delayed post is better than service downtime)."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Amazon, Razorpay, Groww, Google

#### Indepth
Consistency levels (from most to least strict):
1. **Linearizability:** Operations appear instantaneous and in real-time order (strongest)
2. **Sequential consistency:** Operations appear in the same order on all nodes, but not necessarily real-time
3. **Causal consistency:** Causally related operations seen in order; concurrent operations may differ
4. **Eventual consistency:** Given no new writes, all nodes converge (weakest)

Cassandra tunable consistency: `QUORUM` (W + R > N, strong), `ONE` (fastest reads, stale risk), `ALL` (strongest, slowest). Choose per operation.

**The data consistency spectrum in practice:**
- Financial transactions → Linearizability (PostgreSQL with synchronous replication)
- Inventory availability → Sequential (or near-sequential with short lag)
- User profiles → Causal (see your own writes immediately)
- Analytics metrics → Eventual (slight staleness acceptable for dashboards)
- CDN cached content → Eventual (TTL-based, acceptable for static content)

#### 🗣️ How to Explain in Interview
**Interviewer:** What is eventual consistency vs strong consistency from a data architecture perspective?
**Your Response:** "This is the core trade-off of the **CAP Theorem**. **Strong consistency** means that every reader always sees the absolute latest write, which is vital for systems like finance or inventory. The cost is that the system might become slower or even unavailable if nodes can't communicate.

**Eventual consistency**, on the other hand, prioritizes availability and performance. It accepts that for a few milliseconds, different users might see slightly different versions of the data (like a Facebook 'Like' count), but it ensures the system stays fast and responsive. In modern distributed systems, my default is to use **Eventual Consistency** for non-critical social features to maximize scale, while reserving **Strong Consistency** for the small subset of features—like 'Wallet Balance'—where accuracy is absolutely non-negotiable."
