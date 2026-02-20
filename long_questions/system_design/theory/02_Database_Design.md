# 🟡 Database Design & Management — Questions 11–20

> **Level:** 🟡 Mid-level (2–5 yrs)
> **Asked at:** Flipkart, Swiggy, Zomato, Amazon (SDE-2), Razorpay, PhonePe, Meesho

---

### 11. SQL vs NoSQL – when to use what?
"This is one of the most important decisions in system design, and the answer is always 'it depends on your access patterns'.

I choose **SQL** (PostgreSQL, MySQL) when I need ACID transactions, structured relational data with JOINs, and strong consistency. A banking ledger or an e-commerce order system needs SQL — I can't have money disappear due to eventual consistency.

I choose **NoSQL** when I need horizontal scale at massive volume, flexible schema, or a specific access pattern. For example, Cassandra for time-series sensor data (write-heavy, no JOINs), MongoDB for a product catalog with varying attributes, Redis for session storage, or Neo4j for social network friendship graphs.

In real systems, I often use *both* — what they call **polyglot persistence**. Uber uses MySQL for trip data and Cassandra for real-time driver location."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Every product company — Amazon, Flipkart, Swiggy, Razorpay, Google

#### Indepth
| Use Case | Best Choice | Why |
|---|---|---|
| Financial transactions | PostgreSQL | ACID, JOINs, strong consistency |
| User sessions | Redis | Key-Value, TTL, nanosecond latency |
| Product catalog | MongoDB | Flexible schema, nested docs |
| Social graph | Neo4j | Graph traversal queries |
| Real-time analytics | ClickHouse/Druid | Columnar, OLAP queries |
| Chat messages | Cassandra | Write-heavy, wide-column, time-ordered |
| Search | Elasticsearch | Full-text inverted index |

**Trade-offs:** NoSQL sacrifices ACID. Most NoSQL DBs offer eventual consistency. This is acceptable for a user's post appearing with a slight delay on followers' feeds. It's **not** acceptable for a payment being marked as failed after money was deducted.

---

### 12. How do you scale a database?
"Scaling a DB is the most complex part of system design. I follow a hierarchy of increasing complexity.

First, **optimize queries and add indexes** — this is free and often solves 80% of performance issues. Then **add a caching layer** (Redis) in front of the DB — reduce read pressure dramatically. Then **add read replicas** — writes go to master, reads go to N replicas. This is what most mid-scale systems (Flipkart, Zomato level) do.

If that's not enough, **vertical scale the DB** — bigger EC2 instance. The last resort is **sharding** — split data across multiple DB nodes. This is what WhatsApp, Instagram, and Twitter have done at extreme scale."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Amazon, Flipkart, Google, PhonePe, Twitter/X

#### Indepth
The full scaling ladder:
1. **Indexing:** Correct indexes reduce O(n) scans to O(log n). Always start here.
2. **Query optimization:** Avoid N+1 queries, use `EXPLAIN ANALYZE` to find slow queries.
3. **Connection pooling:** PgBouncer/PgPool — DBs have limited connections (~200 for Postgres). Pool reuses them.
4. **Caching:** Redis/Memcached sits in front. 99% cache hit rate = 1% of traffic reaches DB.
5. **Read replicas:** Master handles writes; replicas handle reads. Replication lag is a trade-off.
6. **Partitioning:** Split one big table into smaller physical pieces within the same DB (different from sharding across DBs).
7. **Sharding:** Last resort. Each shard is a separate DB instance. Requires shard-aware app code.
8. **Managed DB services:** AWS Aurora (auto-scaling), Spanner (Google's globally distributed SQL).

---

### 13. What is denormalization and why is it useful?
"Normalization is the process of organizing a DB to eliminate redundancy. Denormalization deliberately **re-introduces redundancy** to optimize read performance.

In a normalized schema, getting a user's order with product details requires 3 JOINs: Users → Orders → OrderItems → Products. At scale, these JOINs are expensive. Denormalization means storing a pre-computed, redundant copy of that data in a single collection/table — no joins at query time.

I use denormalization heavily with NoSQL. In MongoDB, I embed the product name and price directly into the order document. Reads are instant — at the cost of write complexity (when a product's name changes, I must update it in many places)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Flipkart, Amazon, Swiggy (DB design discussions)

#### Indepth
Denormalization strategies:
- **Embedding (NoSQL):** Nest related data directly in the document. Best for one-to-few relationships.
- **Pre-computed aggregations:** Store `total_order_count` in the user row instead of `COUNT(*)` each time.
- **Materialized views:** DB-supported pre-computed views (PostgreSQL supports these). Updated on schedule or on write.
- **CQRS (Command Query Responsibility Segregation):** Maintain two data models — normalized for writes, denormalized for reads. The read model is updated asynchronously via events.

**Warning:** Denormalization is a *trade-off agreement*. You accept: higher storage usage, more complex writes, potential for inconsistency if update logic is buggy. Always ensure business requirements allow for this trade-off before denormalizing.

---

### 14. CAP theorem – explain and give examples.
"CAP theorem says that in a **distributed system**, you can only guarantee two out of three properties: **Consistency**, **Availability**, and **Partition Tolerance**.

Network partitions (nodes can't talk to each other) are *inevitable* in distributed systems — so Partition Tolerance is not optional. That leaves the real choice: **CP** or **AP**.

**CP (Consistent and Partition-Tolerant):** The system stays consistent but might be unavailable during a partition. Example: HBase, MongoDB (with strong reads). Use this for banking — it's better to reject a transaction than to show stale data.

**AP (Available and Partition-Tolerant):** The system stays up but might return stale data. Example: Cassandra, DynamoDB. Use this for social media — it's okay if a tweet appears 500ms later for some users rather than the site going down."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Google, Amazon, Uber, Twitter/X, Flipkart (principal-level design rounds)

#### Indepth
CAP is often misunderstood. A few nuances:
- **You always tolerate partitions** in real distributed systems. The choice is what to do *when* a partition occurs — sacrifice Consistency (AP) or Availability (CP).
- **PACELC theorem** is the more complete model: Even without partitions (else), you have a trade-off between **Latency** and **Consistency**. A CP system with low latency typically requires strong coordinator, but that coordinator is a bottleneck.
- Cassandra lets you *tune* per query: `QUORUM` reads for consistency vs `ONE` reads for availability — best of both worlds at query-level granularity.
- **Write scenarios matter most:** During a partition, a CP system will reject writes to maintain consistency. An AP system will accept writes but resolve conflicts later (vector clocks, last-write-wins).

---

### 15. What is eventual consistency?
"Eventual consistency is a consistency model where the system **guarantees that all nodes will eventually converge to the same value** — not instantly, but eventually.

DNS is the classic example everyone knows. When you update an A record, it doesn't propagate to all DNS resolvers instantly. It takes minutes or hours. But eventually, everywhere in the world resolves to the new IP. This delay is acceptable.

Another example: Amazon's shopping cart. If you add an item from London and simultaneously remove it from Mumbai with a network partition, Amazon's Dynamo DB will show different states briefly. After the partition heals, it merges using conflict resolution. The 'add' wins because Dynamo uses last-write-wins or vector clocks."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Amazon (Dynamo is their system!), Flipkart, Swiggy, Google

#### Indepth
Eventual consistency has several sub-types:
- **Monotonic Read Consistency:** Once you read a value, all subsequent reads return the same or newer value. You won't see data going "backward".
- **Read-Your-Writes Consistency:** After a write, you immediately see your own update (but others might not yet). Used in social media — you should see your own post immediately.
- **Causal Consistency:** Causally related operations are seen in order by all nodes. "If A happened before B, all nodes see A before B." Stronger than eventual, weaker than strong.
- **Strong Eventual Consistency (SEC):** Used in CRDTs. Guarantees that if two nodes have received the same set of updates (in any order), they'll have the same state. Conflict-free by design.

---

### 16. How would you design a schema for a social network?
"I'd design it around the core entities and their relationships: Users, Posts, Friendships, Comments, and Likes.

For the core tables: `users` (user_id PK, username, email, profile_pic, created_at), `posts` (post_id PK, user_id FK, content, media_url, created_at), `friendships` (user1_id, user2_id, status, created_at — composite PK), `comments` (comment_id, post_id FK, user_id FK, content), `likes` (post_id, user_id — composite PK).

But here's where it gets interesting — a pure relational schema struggles at Facebook scale. The friendship graph is best represented in a **graph database** like Neo4j for 'friends of friends' queries, while post storage goes to a NoSQL store like Cassandra for write throughput."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Meta, Twitter, LinkedIn, Snap — companies that *are* social networks; also asked at Flipkart for their social/community features

#### Indepth
Key design considerations:
- **The Newsfeed Problem:** Generating a user's feed by querying posts from all their friends at read time (pull model) is O(n) — where n is number of friends. At scale, use a **push/fan-out-on-write** model: when a user posts, pre-compute and write to all their followers' feed caches. Twitter uses this for users with <1M followers. For celebrities (Justin Bieber, 100M+ followers), fan-out is too expensive — use a hybrid: celebrity's tweets are fetched on read, merged with pre-computed feed.
- **Indexes:** `posts.user_id` (for profile page), `comments.post_id` (for comment threads), `likes.post_id` (for like counts).
- **Sharding:** Shard `posts` by `created_at` (time-based) for efficient range queries on feeds.

---

### 17. Explain database partitioning.
"Database partitioning is dividing a large table into smaller, more manageable pieces that are still treated as a single logical table by the application.

Unlike sharding (which puts data on different *servers*), partitioning happens *within one DB server*. The DB engine manages which rows are in which partition.

I use **range partitioning** on date-based tables: a 10-year transaction table partitioned by month. Queries for 'transactions this month' only scan one partition — massive performance improvement. Old partitions can be archived or dropped cheaply without affecting the live partition."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Paytm, PhonePe, Amazon (for financial/log tables), Swiggy (orders)

#### Indepth
Partitioning types in PostgreSQL/MySQL:
- **Range Partitioning:** rows with `created_at` in 2024 go to one partition, 2025 to another. Best for time-series data. Enables efficient TTL (drop old partition = O(1) instead of O(n) DELETE).
- **List Partitioning:** rows with `region = 'india'` in one partition, `region = 'us'` in another. Good for multi-region data.
- **Hash Partitioning:** `hash(user_id) % 8` → 8 partitions. Even distribution, good for OLTP.
- **Composite Partitioning:** Range by year, then hash by user_id within the year. Combines benefits.

**Partition pruning:** The DB engine can skip irrelevant partitions entirely. A query like `WHERE created_at BETWEEN '2025-01-01' AND '2025-01-31'` only scans the January partition — this is called partition pruning and can give 100x speedups on large tables.

---

### 18. How do you handle schema migrations?
"Schema migrations are versioned, incremental changes to the DB schema. I treat them like code — versioned in Git, peer-reviewed, and applied automatically during deployment.

My key principle is **backward-compatible migrations**. Never drop a column in the same deployment where you remove code that uses it. Instead: 1) Add new column (optional, nullable), 2) Deploy code that writes to both old and new column, 3) Backfill old data, 4) Switch reads to new column, 5) In a separate later deployment, remove old column.

I use **Flyway** or **Liquibase** in Java/Spring projects, **golang-migrate** for Go projects. Every migration is a numbered SQL script. The migration tool tracks applied versions in a `schema_history` table."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Razorpay, PhonePe, Flipkart, Amazon (SDE-2 onwards)

#### Indepth
Zero-downtime migration for adding a column with a NOT NULL constraint (the hardest case):
1. Add column as `NULLABLE` — zero impact, very fast (metadata change)
2. Backfill data: `UPDATE table SET new_col = default WHERE new_col IS NULL` — do in batches to avoid lock contention
3. Add `NOT NULL` constraint with a DEFAULT — Postgres 11+ does this without rewriting the table
4. Swap application code to use new column
5. Remove old column in future sprint

**Dangerous patterns to avoid:**
- Adding `NOT NULL` without default to large tables — full table rewrite, hours of downtime
- Renaming a column directly — breaks live app code immediately
- Dropping a column before removing app references — immediate production crash

---

### 19. What is a write-ahead log?
"A Write-Ahead Log (WAL) is the mechanism databases use to ensure durability — the 'D' in ACID. The idea is simple: **before modifying any data page on disk, first write the change to a sequential log file**.

The reason is economics and physics. Sequential writes to a log are orders of magnitude faster than random writes to data pages. If the DB crashes, it can replay the WAL to reconstruct any changes that weren't yet written to data pages.

PostgreSQL calls it WAL. MySQL calls it the Redo Log. Kafka is essentially one giant WAL. This pattern is used everywhere durability matters."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Google, Amazon (infrastructure roles), PhonePe, Razorpay (for DB architecture discussions)

#### Indepth
WAL serves multiple purposes beyond crash recovery:
1. **Replication:** Postgres Streaming Replication ships WAL records to replicas. The replica replays them to stay in sync. This is how read replicas work.
2. **Point-in-Time Recovery (PITR):** By archiving WAL segments, you can restore the DB to *any* point in time — invaluable for "undo" a bad migration or accidental DELETE.
3. **Logical Replication / Change Data Capture (CDC):** Tools like Debezium read the WAL stream and publish DB changes as events to Kafka — enabling event-driven architectures without any application code changes.

**WAL in Kafka:** Kafka's core data structure is a commit log — essentially a WAL on disk. This is why it's so fast for writes and why it can replay events from any offset, just like a DB WAL supports PITR.

---

### 20. How to design a time-series database?
"Time-series data is append-only, timestamped data — like CPU metrics every 10 seconds, stock prices every millisecond, or IoT sensor readings. Standard relational DBs handle it poorly at scale.

The key design insight: time-series data is **write-heavy and almost never updated**. We only append new readings. Queries are always range-based: 'all CPU readings in the last hour'. This makes it perfect for sequential, compressed storage.

I'd design it with: an append-only storage engine (LSM tree), automatic data compression (delta encoding works perfectly since consecutive readings differ by small amounts), downsampling (storing raw data for recent periods, pre-aggregated summaries for older data), and automatic TTL (nobody needs CPU readings from 3 years ago)."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Google (SRE), Netflix, Amazon CloudWatch team, Hotstar (metrics platform)

#### Indepth
Purpose-built TSDB storage techniques:
- **Delta encoding:** Instead of storing `[100, 101, 102, 99, 101]`, store `[100, +1, +1, -3, +2]`. Massive compression.
- **XOR encoding + Gorilla compression:** Used by Facebook's Gorilla TSDB. XOR consecutive float64 timestamps — exploits the fact that consecutive metrics share many bits.
- **LSM Tree storage:** InfluxDB, LevelDB, Cassandra all use LSM trees — perfect for append-heavy workloads. Writes go to in-memory memtable, periodically flushed to SSTables on disk.
- **Rollups and downsampling:** Store raw 1-second data for 7 days, 1-minute aggregates for 30 days, 1-hour aggregates for 1 year. Queries on "last year" use pre-aggregated data — fast.
- **Retention policies:** Automatically delete data older than N days.

Technologies: **InfluxDB** (general purpose), **Prometheus** (pull-based metrics, PromQL), **TimescaleDB** (Postgres extension), **Apache Druid** (analytics at scale), **VictoriaMetrics** (high-performance, drop-in Prometheus replacement).
