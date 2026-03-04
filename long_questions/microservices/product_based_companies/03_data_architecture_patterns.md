# 🗄️ Microservices — Data Architecture & Communication Patterns (Product Companies)

> **Level:** 🔴 Senior – Principal
> **Asked at:** Amazon, Google, Flipkart, Uber, CRED, Groww, Razorpay

---

## Q1. How do you migrate a massive monolith to microservices? Walk me through the Strangler Fig Pattern.

"The most dangerous approach to microservices is the **'Big Bang'** rewrite — throw away the monolith, rewrite everything from scratch, and switch over on a go-live date. This almost always fails. It takes years, the monolith keeps evolving during the rewrite, and you're not learning from real production traffic.

The **Strangler Fig Pattern** (named after a vine that slowly grows around a host tree) is the industry-standard approach for incremental migration:

**Steps:**
1. **Identify the right seam to cut first:** Don't start with the most complex, coupled domain. Start with a domain that is:
   - Well-defined with clear business boundaries
   - Infrequently changed by the legacy team
   - High business value (justifies the engineering investment)
   - *Example: Extract the Notification Service first — it's obviously separate and can easily send its own events.*

2. **Place a Facade/Proxy in front of the monolith:** An API Gateway or reverse proxy sits between clients and the system. Initially, 100% of traffic goes to the monolith.

3. **Build the new microservice in parallel:** Build the Notification Service as a standalone service with its own database. Do not touch the monolith DB.

4. **Migrate traffic incrementally:** Route 10% of `/api/notifications/*` traffic to the new service. Monitor. Route 50%. Monitor. Route 100%.

5. **Delete the old code from the monolith:** Once the new service is stable at 100% and backfilled with all historical data, delete the notification code from the monolith.

6. **Repeat for the next domain.**

**Key challenges:**
- **Data migration:** The new service needs its own authoritative copy of the data. This often requires a dual-write phase where both the monolith and new service write to their respective stores, then a cutover.
- **Shared DB:** During migration, new and old code might share a database. Managing schema changes across both codebases simultaneously is the hardest part."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Scale-ups undergoing modernization (CRED, Zepto), and in system design interviews at Flipkart and Amazon for senior roles where engineering judgment and not just technical knowledge is tested.

#### Indepth
**Anti-Corruption Layer (ACL):** When the new microservice must read entities defined in the monolith's model (different naming conventions, schemas), an ACL is an adapter layer that translates between the two models. It prevents the 'corruption' of the new service's clean domain model with the legacy monolith's legacy concepts.

---

## Q2. When would you choose gRPC over REST for internal microservice communication?

"Most teams default to REST for microservice communication because it's familiar. But REST (JSON over HTTP/1.1) has real performance bottlenecks at scale that gRPC addresses:

**Why gRPC is superior for internal communication:**

| Property | REST (JSON/HTTP1.1) | gRPC (Protobuf/HTTP2) |
|---|---|---|
| **Serialization** | JSON: text-based, human-readable, slow to parse | Protobuf: binary, ~10x smaller, ~5x faster to serialize |
| **Connection** | One request per TCP connection (or connection pooling overhead) | HTTP/2 multiplexing: many requests over one connection |
| **API Contract** | No enforced contract (OpenAPI is optional) | `.proto` file is a strict, versioned contract; code generated automatically |
| **Streaming** | Not natively supported | Native bi-directional streaming (server-side push) |
| **Browser support** | 100% native | Requires gRPC-Web proxy (not native in browsers) |

**When to choose gRPC:**
- High-frequency, low-latency internal service-to-service calls (e.g., 100K+ calls/sec between Order and Inventory)
- Polyglot services where you want compile-time type safety across Go, Java, Python, Node
- Streaming use cases: real-time sensor data, streaming search results, server-sent push notifications
- When you have strict API contracts and need backward-compatibility versioning managed via Protobuf field numbering rules

**When to stick with REST:**
- Public-facing APIs that browsers and mobile apps call directly
- When simplicity and developer familiarity outweighs the performance gains
- When your team is small and tooling overhead isn't worth it"

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber (Go-based microservices fleet), Google, and any company building high-throughput internal platforms where the overhead of JSON serialization and HTTP/1.1 actually shows up on latency dashboards.

#### Indepth
**Protobuf Field Numbering for Backward Compatibility:** In Protobuf, each field has a unique number (not just a name). If you want to add a new field to a message definition, you add a new field with a new number. Old services ignore unknown field numbers silently, and new services treat missing fields as zero-values. This means you can evolve your API without breaking old callers — something REST/JSON doesn't enforce natively.

---

## Q3. How do you handle API versioning in a large microservices ecosystem?

"API versioning is unavoidable. As services evolve, clients (frontend, mobile apps, other services) cannot always update simultaneously. You need a strategy to handle breaking changes gracefully.

**Primary approaches:**

**1. URI Path Versioning (Most Common):**
```
/api/v1/orders
/api/v2/orders
```
- *Pros:* Explicit, easy to understand, easy to route at the API Gateway level.
- *Cons:* URL semantically represents a resource, not a version. Old URIs must be maintained indefinitely.

**2. HTTP Header Versioning:**
```http
GET /api/orders
Accept: application/vnd.company.v2+json
```
- *Pros:* Clean URLs. The resource URI is stable.
- *Cons:* Less visible. Harder to test via browser. Caching can be tricky.

**3. Query Parameter Versioning:**
```
/api/orders?version=2
```
- Easiest to implement but considered poor REST design.

**My recommended strategy at scale:**
1. **Never break existing API contracts.** Additive changes (new fields) are always safe. Prefer adding new optional fields to existing versions rather than creating new versions.
2. **Use Consumer-Driven Contract Testing (Pact)** so any breaking change to a Provider API fails the CI build before it reaches production.
3. When a breaking change is truly necessary: **run both versions in parallel for a deprecation window** (e.g., 6 months). The API Gateway routes `v1` traffic to the old service or an adapter, while `v2` traffic goes to the new service implementation.
4. **Sunset headers:** Add `Deprecation` and `Sunset` HTTP headers to v1 responses to notify API consumers that v1 will be removed on a specific date."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Companies with large external API ecosystems (Razorpay, payment gateways), and platform engineering teams managing APIs consumed by mobile apps that can't force-upgrade all users simultaneously.

#### Indepth
**API Gateway as a Versioning Adapter:** A powerful pattern: instead of maintaining two separate running versions of a service (cost), the API Gateway can contain transformation logic or route to a lightweight 'adapter service' that maps v1 requests into v2 format and passes them to the single v2 service. The old service code is retired, but v1 clients still work.

---

## Q4. How do you design a microservices system for multi-region deployment? What are the key data challenges?

"Multi-region deployment is primarily a **data problem**, not a compute problem. Deploying containerized services to multiple AWS regions (`us-east-1`, `eu-west-1`, `ap-south-1`) is relatively straightforward with Kubernetes. Making the **data consistent** across those regions is the hard part.

**The core tension: CAP Theorem in action**
When a network partition occurs between regions, you must choose:
- **Consistency:** Reject writes in the secondary region until the primary acknowledges them (high availability sacrifice, no stale reads).
- **Availability:** Accept writes in both regions even if they can't sync, and resolve conflicts later (eventual consistency, potential data conflicts).

**Common architectures:**
1. **Active-Passive (Leader-Follower):**
  - All writes go to one primary region (e.g., `us-east-1`).
  - Secondary regions (`eu-west-1`) receive replicated read replicas.
  - *Pros:* Simple consistency model — no write conflicts.
  - *Cons:* Write latency is high for users in Europe. If the primary region goes down, a manual or automated failover is required (RTO matters here).

2. **Active-Active (Multi-Master):**
  - Users in each region can write to their local database.
  - Writes are asynchronously replicated across regions.
  - *Pros:* Low write latency globally. High availability.
  - *Cons:* Write conflicts are possible if two users in different regions modify the same record simultaneously. Conflict resolution logic (e.g., Last-Write-Wins, CRDTs) is complex.

**Data Residency / GDPR compliance:**
EU user data must stay within EU regions. This requires data partitioning by user geography at the application level — routing EU users to EU clusters and preventing their PII from being replicated to US regions."

#### 🏢 Company Context
**Level:** 🔴 Principal | **Asked at:** Amazon, Google, and globally-operating fintech companies. Multi-region architecture is typically a Principal Engineer / System Design interview topic reserved for 7+ year engineers.

#### Indepth
**CRDTs (Conflict-free Replicated Data Types):** For certain data structures (counters, sets, maps), mathematical structures called CRDTs guarantee that concurrent writes in different regions can always be merged deterministically without conflicts. Used in systems like Cassandra (eventual consistency) and Riak. They don't work for all data types (e.g., arbitrary financial balances), but are powerful for collaborative or counter-based data.

---

## Q5. How do you approach zero-downtime database schema migrations in a running microservices system?

"Database schema migrations are one of the most dangerous operations in microservices because the database is a shared artifact that the old and new versions of your service must both be able to use simultaneously during a rolling deployment.

**The core problem:**
If your rolling deployment runs old and new pod versions simultaneously (which it always does), and your new version expects a new column `user_email_verified BOOLEAN`, but old pods don't know about this column — you need the schema change to be compatible with BOTH versions during the transition.

**The Expand-Contract Pattern (the only safe approach):**

**Phase 1: Expand (Deploy backward-compatible schema change)**
- Add the new column as `NULLABLE` with no default. Run the DB migration before deploying new code.
- The old pods see the new column but don't write to it (they only write the columns they know about). New pods start writing to the new column.

**Phase 2: Backfill**
- Run a background job to populate the new column for existing rows.

**Phase 3: Contract (Clean up legacy)**
- Once 100% of pods are on the new version and the old column is no longer read/written, drop the old column or add the NOT NULL constraint.

**Tooling:**
- **Flyway / Liquibase:** Manages versioned SQL migration scripts. Locks the DB during migration and tracks which scripts have been applied. Use `V1__Add_nullable_email_verified.sql` (Expand), and later `V2__Backfill_and_add_constraint.sql` (Contract).
- Never rename or drop a column in a single deployment."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Any company running multiple microservice replicas with rolling deployments. A classic trap question — developers new to distributed systems often propose a risky single-step migration.

#### Indepth
**Online Schema Change Tools:** For large tables (100M+ rows), running a standard `ALTER TABLE` locks the table for minutes or hours, causing an outage. Tools like **gh-ost** (GitHub's Online Schema Change) and **Percona Online Schema Change** create a shadow table, replay writes to it, and do an atomic rename at the end — allowing schema migrations with near-zero downtime on running production databases.
