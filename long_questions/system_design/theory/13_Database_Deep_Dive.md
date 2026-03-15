# 🔴 Database Deep Dive — Questions 121–130

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Amazon (DynamoDB, Aurora), Google (Spanner, Bigtable), Flipkart, Razorpay, PhonePe — any system with significant data workloads

---

### 121. How to choose between SQL and NoSQL?
"Choosing between SQL and NoSQL is one of the most consequential design decisions in a system. My decision framework starts with the data model and access patterns.

**Choose SQL (PostgreSQL, MySQL)** when: the data is highly relational (users → orders → products), you need ACID transactions (financial data, inventory), you need complex queries and ad-hoc analysis (JOINs, GROUP BY, subqueries), or your schema is relatively stable.

**Choose NoSQL** when: you need to scale to petabytes of data (Cassandra, DynamoDB), your access patterns are simple key lookups or by a primary key range (Redis, DynamoDB), you have rapidly evolving schemas (MongoDB's document model), or you need sub-millisecond reads at massive scale (Redis cache)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Every system design interview — this is the foundational DB choice

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to choose between SQL and NoSQL?
**Your Response:** Choosing between SQL and NoSQL is one of the most consequential design decisions. My framework starts with the data model and access patterns. I choose SQL like PostgreSQL or MySQL when the data is highly relational like users to orders to products, when I need ACID transactions for financial data or inventory, when I need complex queries and ad-hoc analysis with JOINs and GROUP BY, or when my schema is relatively stable. I choose NoSQL when I need to scale to petabytes of data with Cassandra or DynamoDB, when my access patterns are simple key lookups, when I have rapidly evolving schemas with MongoDB's document model, or when I need sub-millisecond reads at massive scale with Redis.

#### Indepth
Decision matrix:

| Criteria | SQL (PostgreSQL) | NoSQL (Cassandra/DynamoDB) |
|---|---|---|
| ACID transactions | ✅ Yes — multi-row, multi-table | ❌ No / limited (Dynamo: single item) |
| Horizontal scaling | ⚠️ Harder — sharding required | ✅ Designed for it |
| Query flexibility | ✅ Full SQL: JOINs, aggregates | ❌ Limited — data model per access pattern |
| Schema flexibility | ❌ Fixed schema / migrations | ✅ Flexible document/column structure |
| Consistency | ✅ Strong by default | ⚠️ Eventual by default (tunable) |
| Latency at scale | ⚠️ Degrades with massive data | ✅ Consistent low latency |
| Operational complexity | ✅ Well-understood | ⚠️ Different mental model |

**Hybrid approach (common for large companies):** PostgreSQL for the transactional core (user accounts, orders, payments), Redis for caching, Cassandra/DynamoDB for write-heavy event/activity data, Elasticsearch for search. Each tool used for what it's best at.

**Pitfall: Using MongoDB as a relational DB.** Many teams choose MongoDB for its flexibility, then implement complex `$lookup` (JOIN equivalent) operations in every query. This performs much worse than SQL. If you're doing JOINs in MongoDB, you should be using PostgreSQL.

---

### 122. What is database indexing?
"An index is a **data structure that allows the database engine to find rows matching a query criterion without scanning every row in the table**.

Without an index, `SELECT * FROM orders WHERE customer_id=123` requires scanning all 500 million rows. With an index on `customer_id`, the DB tree-walks the B-tree index in O(log n) and goes directly to the matching rows.

Indexes have a cost: every INSERT, UPDATE, DELETE must also update the index. An over-indexed table (20+ indexes) has very fast reads but painfully slow writes. I always add indexes based on measured query slow logs, not gut feeling."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Every backend engineering interview — fundamental DB performance knowledge

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is database indexing?
**Your Response:** An index is a data structure that allows the database engine to find rows matching a query criterion without scanning every row in the table. Without an index, a query like SELECT * FROM orders WHERE customer_id=123 requires scanning all 500 million rows. With an index on customer_id, the database tree-walks the B-tree index in O(log n) time and goes directly to the matching rows. Indexes have a cost - every INSERT, UPDATE, DELETE must also update the index. An over-indexed table with 20+ indexes has very fast reads but painfully slow writes. I always add indexes based on measured query slow logs, not gut feeling.

#### Indepth
Index types:
- **B-tree index (default):** Balanced binary search tree. Good for equality (`=`), ranges (`>`, `<`, `BETWEEN`), sorting (`ORDER BY`). The default index type in PostgreSQL and MySQL.
- **Hash index:** Only for equality. O(1) lookup. Cannot be used for range queries or sorts. PostgreSQL supports it; MySQL's Memory engine uses it.
- **GiST / GIN (PostgreSQL):** Generalized Search Tree — for geometric data, full-text search, arrays. `CREATE INDEX ON documents USING gin(to_tsvector('english', content))` enables fast full-text search.
- **Covering index:** An index that includes all columns needed by a query — DB can answer entirely from the index without touching the table. `CREATE INDEX idx_orders_cover ON orders(customer_id) INCLUDE (order_date, total_amount)`.
- **Partial index:** Index only a subset of rows. `CREATE INDEX idx_active_users ON users(email) WHERE active = true`. Smaller index, faster for queries filtering on the indexed condition.
- **Composite index:** Index on multiple columns `(customer_id, order_date)`. The leftmost prefix rule: query must use the leftmost column(s) to benefit from the index. A query `WHERE order_date=X` won't use this index; `WHERE customer_id=X AND order_date>Y` will.

`EXPLAIN (ANALYZE, BUFFERS)` in PostgreSQL shows index usage, estimated vs actual row counts, buffer hits. The starting point for any query optimization.

---

### 123. What is query optimization?
"Query optimization is the process of improving the performance of database queries — making them run faster and consume fewer resources.

Three-step process: (1) **Identify** slow queries via slow query log or `pg_stat_statements`. (2) **Analyze** execution plan with `EXPLAIN ANALYZE`. (3) **Fix** by adding indexes, rewriting the query, partitioning the table, or denormalizing data.

Common quick wins: add an index for the column in the WHERE clause, avoid `SELECT *` (fetch only needed columns), rewrite correlated subqueries as JOINs, use `LIMIT` for pagination instead of fetching all rows."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any data-intensive backend role — Flipkart (catalog queries), Razorpay (transaction queries), Hotstar (user queries)

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is query optimization?
**Your Response:** Query optimization is the process of improving database query performance - making them run faster and consume fewer resources. My three-step process is: first identify slow queries via the slow query log or pg_stat_statements, then analyze the execution plan with EXPLAIN ANALYZE, and finally fix by adding indexes, rewriting the query, partitioning the table, or denormalizing data. Common quick wins are adding an index for columns in the WHERE clause, avoiding SELECT * to fetch only needed columns, rewriting correlated subqueries as JOINs, and using LIMIT for pagination instead of fetching all rows.

#### Indepth
Common anti-patterns and fixes:
1. **N+1 Query Problem:** Fetch 100 orders, then for each order fetch the customer in a separate query → 101 queries. Fix: JOIN or `IN (...)` to fetch all in one query, or DataLoader batching in GraphQL.
2. **`SELECT *`:** Fetches unused columns, wastes network/memory. Fix: specify exact columns.
3. **OFFSET-based pagination at large offsets:** `SELECT * FROM orders ORDER BY id LIMIT 10 OFFSET 1000000` — DB scans and discards 1M rows. Fix: keyset (cursor) pagination: `WHERE id > :last_seen_id LIMIT 10`. O(1) regardless of offset.
4. **Function in WHERE clause:** `WHERE YEAR(created_at) = 2024` — prevents index use on `created_at`. Fix: `WHERE created_at BETWEEN '2024-01-01' AND '2024-12-31'`.
5. **Implicit type cast:** `WHERE user_id = '123'` when `user_id` is INT — forces type conversion, may skip index. Fix: use correct type in query.
6. **Missing JOIN index:** Join on `orders.customer_id` without index on `customer_id` → nested loop scanning orders for each customer. Fix: index on the JOIN column.

PostgreSQL-specific tools:
- `pg_stat_statements`: Find the 10 most time-consuming queries across all executions
- `auto_explain`: Automatically log explain plans for slow queries
- `pg_trgm`: Trigram index for LIKE queries — `WHERE name LIKE '%partial%'` can use GIN index with pg_trgm extension

---

### 124. What is a time-series database?
"A time-series database (TSDB) is **optimized for storing and querying time-stamped data** — data where every point has a timestamp and the primary access pattern is by time range.

Examples: IoT sensor readings (temperature every second), server metrics (CPU every 15 seconds), financial tick data (stock price every millisecond), application metrics (request counts per minute).

Traditional SQL databases can store time-series data but perform poorly at scale — they're not optimized for time-range scans with high write throughput. TSDBs like InfluxDB, TimescaleDB, and Victoria Metrics are designed specifically for this: high-speed appends, efficient time-range queries, automatic compression of historical data, and downsampling (aggregate old data into hourly/daily summaries to save space)."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Monitoring/observability systems, IoT companies, fintech (trading data), any Prometheus/Grafana setup discussion

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a time-series database?
**Your Response:** A time-series database is optimized for storing and querying time-stamped data where every point has a timestamp and the primary access pattern is by time range. Examples include IoT sensor readings, server metrics, financial tick data, and application metrics. Traditional SQL databases can store time-series data but perform poorly at scale. TSDBs like InfluxDB, TimescaleDB, and Victoria Metrics are designed specifically for high-speed appends, efficient time-range queries, automatic compression of historical data, and downsampling to aggregate old data into hourly or daily summaries to save space.

#### Indepth
Time-series DB characteristics:
- **Write pattern:** Append-only high throughput. Millions of data points per second. Data arrives roughly in time-order (exceptions for out-of-order writes).
- **Query pattern:** `SELECT metric WHERE time BETWEEN T1 AND T2`. Range scans by time. Aggregations (avg, max, p99) over time windows.
- **Compression:** Time-series data is highly compressible. Adjacent timestamps have small deltas → delta-encode timestamps. Values for metrics like CPU often change slowly → Gorilla compression (Facebook's algorithm, ~12x compression). InfluxDB, Prometheus, VictoriaMetrics all use variants of this.
- **Downsampling:** After 30 days, aggregate per-minute data into hourly averages. After 1 year, further aggregate to daily averages. Reduces storage dramatically while preserving trend visibility.

Popular options:
- **InfluxDB:** Purpose-built TSDB. InfluxQL and Flux query languages. Good for IoT and metrics.
- **TimescaleDB:** PostgreSQL extension — adds TSDB optimizations on top of Postgres. Hypertables (auto-partitioned by time). Full SQL.
- **Prometheus TSDB:** Prometheus's local storage engine. Optimized for pull-based metrics. Limited retention (weeks).
- **VictoriaMetrics:** High-performance, Prometheus-compatible. Better compression, better performance. Good Prometheus drop-in.

---

### 125. What is full-text search and how to implement it?
"Full-text search allows users to **search natural language text content** — not exact matches, but semantic matching. A search for 'running shoes' should match documents containing 'shoe for jogging' — understanding language, not just character matching.

The foundation is an **inverted index**: for every word in the corpus, store a list of documents containing that word. `running → [doc1, doc5, doc9]`. To search for 'running shoes', intersect both lists. Elasticsearch, Solr, and PostgreSQL's built-in FTS all use inverted indexes.

For basic search (Flipkart product search, Swiggy restaurant search), I use **Elasticsearch** — it handles tokenization, stemming, relevance scoring (BM25), and filtering on structured fields simultaneously."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Flipkart, Amazon, Swiggy, Zomato, any product with search — search is pervasively important

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is full-text search and how to implement it?
**Your Response:** Full-text search allows users to search natural language text content - not exact matches, but semantic matching. A search for 'running shoes' should match documents containing 'shoe for jogging' by understanding language, not just character matching. The foundation is an inverted index - for every word in the corpus, store a list of documents containing that word. To search for 'running shoes', intersect both lists. For basic search like product search, I use Elasticsearch which handles tokenization, stemming, relevance scoring with BM25, and filtering on structured fields simultaneously.

#### Indepth
Full-text search processing pipeline:
1. **Tokenization:** "running shoes for beginners" → ["running", "shoes", "for", "beginners"]
2. **Normalization:** Lowercase, remove punctuation
3. **Stop word removal:** "for" → removed
4. **Stemming/Lemmatization:** "running" → "run", "shoes" → "shoe"
5. **Inverted index:**
   - "run" → [doc1(pos:3), doc5(pos:8)]
   - "shoe" → [doc1(pos:7), doc3(pos:2), doc5(pos:1)]

Elasticsearch query pipeline:
```json
{
  "query": {
    "bool": {
      "must": [
        {"match": {"name": "running shoes"}}
      ],
      "filter": [
        {"term": {"category": "sports"}},
        {"range": {"price": {"gte": 500, "lte": 5000}}}
      ]
    }
  },
  "sort": [
    {"_score": "desc"},
    {"popularity": "desc"}
  ]
}
```

Relevance scoring: BM25 (Best Match 25) is the default algorithm. Higher score for: term in many positions in document (TF — term frequency), term rare across documents (IDF — inverse document frequency), shorter documents (length normalization).

**Typo tolerance:** Elasticsearch's fuzzy query uses Levenshtein edit distance. `{"fuzzy": {"name": {"value": "shooes", "fuzziness": 1}}}` matches "shoes" (edit distance 1 — one deletion).

---

### 126. What is Database Partitioning?
"Partitioning divides large tables into smaller, physically distinct pieces — **partitions** — that can be queried, maintained, and backed up independently. Unlike sharding (across multiple servers), partitioning keeps all parts on one DB server.

The most common type: **range partitioning by time**. An `orders` table partitioned by month: all October 2024 orders in `orders_2024_10` partition, November in `orders_2024_11`, etc. A query for 'orders in October 2024' only scans the October partition (partition pruning) — not all 5 years of data.

This is the single most impactful optimization for time-based data — query time goes from 'scan 2 billion rows' to 'scan 50 million rows in this month's partition'."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Razorpay (transaction history), Flipkart (order history), PhonePe, any system with large time-series operational data

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Database Partitioning?
**Your Response:** Database partitioning divides large tables into smaller, physically distinct pieces called partitions that can be queried, maintained, and backed up independently. Unlike sharding across multiple servers, partitioning keeps all parts on one database server. The most common type is range partitioning by time - an orders table partitioned by month where all October orders are in one partition, November in another, etc. A query for orders in October only scans the October partition through partition pruning, not all 5 years of data. This is the single most impactful optimization for time-based data.

#### Indepth
Partitioning strategies:
- **Range partitioning:** On a column with ranges (date, ID ranges). `orders_2024_Q1`, `orders_2024_Q2`. Most common for time-series data.
- **List partitioning:** On a specific list of values. `orders_region_north`, `orders_region_south`. Good for geographic or categorical distribution.
- **Hash partitioning:** `partition_num = hash(user_id) % N`. Even distribution, unpredictable access by value. Good for general distribution when no natural range or list.
- **Composite partitioning:** Combine strategies. Range by year then hash within year.

PostgreSQL declarative partitioning:
```sql
CREATE TABLE orders (
  order_id BIGINT,
  customer_id BIGINT,
  order_date DATE,
  amount DECIMAL(15,2)
) PARTITION BY RANGE (order_date);

CREATE TABLE orders_2024_q1 PARTITION OF orders
  FOR VALUES FROM ('2024-01-01') TO ('2024-04-01');

CREATE TABLE orders_2024_q2 PARTITION OF orders
  FOR VALUES FROM ('2024-04-01') TO ('2024-07-01');
```

Query with partition pruning:
```sql
SELECT * FROM orders WHERE order_date BETWEEN '2024-01-01' AND '2024-03-31';
-- EXPLAIN shows: Seq Scan on orders_2024_q1 (only this partition scanned)
```

Benefits: parallel queries across partitions, drop partition (instant) vs DELETE (slow), partition-level vacuuming in PostgreSQL.

---

### 127. What is the N+1 query problem?
"The N+1 problem occurs when fetching a list of N items requires N additional database queries to fetch associated data for each item — resulting in N+1 total queries instead of 1 or 2.

The classic ORM example in Python (Django, SQLAlchemy) or Java (Hibernate): fetch N posts, then for each post access `post.author` — the ORM fires a separate SELECT for each author. 100 posts → 1 SELECT for posts + 100 SELECT for authors = 101 queries.

The fix: use **eager loading** (`select_related` in Django, `JOIN FETCH` in JPA) to fetch all data in one query:
`SELECT posts.*, users.* FROM posts JOIN users ON posts.author_id = users.id`"

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Backend interviews everywhere — one of the most common performance bugs in web applications

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the N+1 query problem?
**Your Response:** The N+1 problem occurs when fetching a list of N items requires N additional database queries to fetch associated data for each item, resulting in N+1 total queries instead of 1 or 2. The classic example is fetching N posts and then for each post accessing the author - the ORM fires a separate SELECT for each author. 100 posts becomes 1 SELECT for posts plus 100 SELECT for authors = 101 queries. The fix is eager loading using SQL JOINs to fetch all data in one query, or using DataLoader patterns in GraphQL to batch all IDs and fetch in one IN query.

#### Indepth
N+1 fixes by technology:
- **Django:** `Post.objects.select_related('author')` — uses SQL JOIN for one-to-one/many-to-one. `Post.objects.prefetch_related('comments')` — uses separate IN query for one-to-many.
- **Hibernate/JPA:** `@ManyToOne(fetch = FetchType.EAGER)` or `JOIN FETCH p.author` in JPQL.
- **SQLAlchemy:** `session.query(Post).options(joinedload(Post.author)).all()`
- **GraphQL:** DataLoader pattern — batch all IDs, fetch in one IN query. Essential for GraphQL resolvers.

How to detect N+1: log all queries in development. Django Debug Toolbar is excellent — shows "101 queries | 234ms" warnings. In production, detect via `pg_stat_statements` — a query pattern like `SELECT * FROM users WHERE id = $1` executed 5000 times in one request is an N+1 signature.

**N+1 at Swiggy example:** Listing 50 restaurants → for each restaurant, lazy-load its cuisines, ratings, distance → 50 × 3 = 150 extra queries. Fix: JOIN or subquery to fetch all restaurant metadata in 1-3 queries total.

---

### 128. What is connection pooling?
"Every database connection requires a TCP handshake, authentication, and session setup — taking 20-50ms. If a 100ms API handler opens a new DB connection for every request, connection setup is 30% of latency — wasteful.

**Connection pooling** reuses existing connections. A pool of N persistent connections is maintained — requests borrow a connection from the pool, use it, and return it. The next request reuses the same connection without reconnecting.

Tools: **PgBouncer** (PostgreSQL connection pooler), **HikariCP** (Java — fastest JDBC pool), Go's `database/sql` package has built-in pooling. Pool sizing rule of thumb: `pool_size = number_of_CPU_cores * 2 + number_of_disk_spindles` (Hikari's recommendation)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company with a production database — Razorpay, Flipkart, PhonePe, Swiggy

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is connection pooling?
**Your Response:** Every database connection requires a TCP handshake, authentication, and session setup taking 20-50ms. If a 100ms API handler opens a new DB connection for every request, connection setup is 30% of latency which is wasteful. Connection pooling reuses existing connections - a pool of persistent connections is maintained, requests borrow a connection, use it, and return it. The next request reuses the same connection without reconnecting. Tools include PgBouncer for PostgreSQL, HikariCP for Java, and Go's database/sql package has built-in pooling.

#### Indepth
PgBouncer modes:
- **Session mode:** One DB connection per client session. Client gets a dedicated connection for its entire session. Equivalent to no pooling for most web apps (HTTP is stateless, each request is a new session).
- **Transaction mode:** Client gets a DB connection only during a transaction. Between transactions, the connection is returned to the pool. Most efficient for web apps — recommended mode.
- **Statement mode:** Pool returns connection after each SQL statement. Most aggressive reuse, but transactions spanning multiple statements don't work. Rarely suitable.

Serverless + DB connection explosion:
- AWS Lambda: each cold start opens new DB connections. At 10K concurrent lambdas × 5 connections = 50K connections → PostgreSQL default max_connections=100 → CRASH.
- Fix: **RDS Proxy** (AWS managed PgBouncer/connection pooler for RDS). Lambda connects to RDS Proxy, which multiplexes thousands of Lambda connections into a few hundred actual DB connections.

Go's `database/sql` pool config:
```go
db, _ := sql.Open("postgres", dsn)
db.SetMaxOpenConns(25)   // max connections to DB
db.SetMaxIdleConns(10)   // keep 10 connections idle (ready to use)
db.SetConnMaxLifetime(5 * time.Minute) // rotate connections to prevent stale
```

---

### 129. Difference between OLTP and OLAP.
"OLTP (Online Transaction Processing) and OLAP (Online Analytical Processing) represent two fundamentally different database workloads.

**OLTP** is the operational database for running the business: handling user-facing transactions — place order, process payment, update profile. Queries are simple, affect few rows, require sub-millisecond latency. Schema is highly normalized. PostgreSQL, MySQL, Aurora are OLTP systems.

**OLAP** is for business intelligence and analytics: aggregate billions of rows across months of history to answer questions like 'what was our GMV by city last quarter?'. Queries are complex, scan millions of rows, take seconds/minutes. Schema is denormalized (star schema, dimensional model). Redshift, BigQuery, Snowflake, ClickHouse are OLAP systems."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon (Redshift), Google (BigQuery), Flipkart/Swiggy data engineering roles, any company with analytics/BI function

### How to Explain in Interview (Spoken style format)
**Interviewer:** Difference between OLTP and OLAP.
**Your Response:** OLTP and OLAP represent two fundamentally different database workloads. OLTP is the operational database for running the business - handling user-facing transactions like placing orders, processing payments, or updating profiles. Queries are simple, affect few rows, and require sub-millisecond latency. The schema is highly normalized. Examples are PostgreSQL, MySQL, and Aurora. OLAP is for business intelligence and analytics - aggregating billions of rows across months of history to answer questions like 'what was our GMV by city last quarter?' Queries are complex, scan millions of rows, and take seconds to minutes. The schema is denormalized with star schemas. Examples are Redshift, BigQuery, and Snowflake.

#### Indepth
| Feature | OLTP | OLAP |
|---|---|---|
| Purpose | Operational (run the business) | Analytical (understand the business) |
| Query type | Simple: INSERT/UPDATE/SELECT by PK | Complex: GROUP BY, aggregations, JOINs |
| Data volume | GBs to TBs | TBs to PBs |
| Latency requirement | Milliseconds | Seconds to minutes |
| Update frequency | Continuous, real-time | Batch loads (hourly/daily ETL) |
| Schema | Normalized (3NF) | Denormalized (star/snowflake schema) |
| Indexes | Many (fast point lookups) | Column store (no row-level indexes) |
| Examples | PostgreSQL, MySQL, Oracle | BigQuery, Redshift, Snowflake, ClickHouse |
| Users | Application code (millions) | Data analysts, BI tools (hundreds) |

**Why separate them?** Running analytical queries (full table scan of 10B rows for a daily report) on the OLTP database would compete with transaction processing — slowing down user-facing operations. Companies replicate OLTP data into a data warehouse (via ETL/CDC) for analytics, keeping production OLTP fast.

**Real-time analytics:** Druid, ClickHouse, Apache Pinot enable sub-second queries on billions of events. Used for live dashboards (Swiggy's real-time order heatmap, Flipkart's live sale dashboard). These update in seconds/minutes vs hours for traditional batch ETL.

---

### 130. What is change data capture (CDC)?
"Change Data Capture is the process of **identifying and capturing changes (inserts, updates, deletes) made to a database and delivering them to downstream systems** in real-time.

Instead of batch ETL (copy entire table every night), CDC streams every individual change the moment it happens. This enables: real-time data warehouse updates, cache invalidation (update Redis when DB changes), replicating data to a different DB technology, and auditing.

The canonical implementation: **Debezium** reads PostgreSQL's Write-Ahead Log (WAL) or MySQL's binary log — the same log the DB uses for replication. Debezium decodes these log events into structured change records and publishes them to Kafka. Downstream consumers subscribe and react to changes."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Companies with real-time data needs — Flipkart (warehouse sync), Swiggy (real-time analytics), PhonePe (audit trail), Razorpay

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is change data capture (CDC)?
**Your Response:** Change Data Capture is the process of identifying and capturing changes made to a database and delivering them to downstream systems in real-time. Instead of batch ETL copying entire tables every night, CDC streams every individual change the moment it happens. This enables real-time data warehouse updates, cache invalidation, replicating data to different database technologies, and auditing. The canonical implementation is Debezium which reads PostgreSQL's Write-Ahead Log or MySQL's binary log, decodes these log events into structured change records, and publishes them to Kafka for downstream consumers to process.

#### Indepth
CDC implementation approaches:
1. **Log-based CDC (Debezium + Kafka):** Reads DB replication log. Zero impact on source DB performance (log is already written). Captures ALL changes including those from DB migrations and direct SQL. Most reliable and low-overhead. Gold standard.
   ```
   PostgreSQL WAL → Debezium Connector → Kafka topic (product-changes)
                                          ↓
                                    Elasticsearch consumer (search index update)
                                    Redis consumer (cache invalidation)
                                    ClickHouse consumer (analytics update)
   ```
2. **Trigger-based CDC:** DB triggers write to a change log table on every DML. Simpler setup. Doubles write I/O (every INSERT → trigger → another INSERT). Misses direct SQL changes (bypasses application).
3. **Timestamp-based polling:** App/ETL queries `WHERE updated_at > :last_poll_time`. Misses deletes entirely. Requires `updated_at` on every table. Latency = poll interval (minutes).

**Outbox pattern + CDC:** Service writes business event to an `outbox` table in the same local transaction as the business data change. Debezium reads the outbox table via CDC and publishes to Kafka. Guarantees event is published if and only if the DB write succeeded. The most reliable at-least-once event publishing pattern in microservices.
