# 🏗️ Architecture — Product-Based Companies Deep Dive

> **Level:** 🔴 Senior
> **Asked at:** Amazon, Google, Flipkart, Uber, Swiggy, Razorpay, PhonePe, Zepto, CRED, Groww

---

## Q1. Design the architecture for an e-commerce order management system (like Flipkart/Amazon). Walk me through your approach.

"I'd start by clarifying requirements: DAU ~10M users, peak 100K orders/hour (Big Billion Day), 99.99% availability SLA, < 500ms order confirmation latency, support for 5M SKUs.

**HLD Components:**

1. **Client Layer:** React/Native apps → API Gateway (Kong/AWS ALB)
2. **Core Services:**
   - Order Service (owns the order lifecycle)
   - Inventory Service (stock management)
   - Payment Service (integration with Razorpay/payment gateways)
   - Delivery Service (logistics assignment)
   - Notification Service (email, SMS, push)
3. **Data Layer:**
   - PostgreSQL for orders (ACID transactions, audit trail)
   - Redis for inventory counters (atomic decrements, fast reads)
   - Kafka for order events (decoupled downstream processing)
   - Elasticsearch for order search
4. **Asynchronous flow via Sagas:**
   - OrderPlaced → ReserveInventory → ProcessPayment → AssignDelivery
   - Compensating transactions for each step on failure

**Critical design decisions:**
- Inventory counter in Redis with atomic DECR to prevent overselling
- Idempotency keys on the payment API to prevent double charges
- Outbox pattern to guarantee Kafka event delivery after DB commit
- Circuit breaker on Payment service with fallback to queue
- Read model for order tracking (separate from write model)"

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Amazon SDE-2/3, Flipkart Senior, Meesho Senior Backend

#### Deep Dive
**Preventing overselling (Race condition problem):**
```
Naive approach: READ stock (5000), CHECK stock > 0, DECREMENT stock
Race condition: 1000 concurrent requests all read stock=5000, all decrement → stock = 4000 (not -999995!)

Solution: Redis DECRBY with check:
local count = redis.call('DECR', 'inventory:' .. itemId)
if count < 0 then
    redis.call('INCR', 'inventory:' .. itemId)
    return -1  -- Out of stock
end
return count  -- Still in stock
-- This is atomic at the Redis level — no race condition possible
```

**Payment idempotency:**
- Client generates `idempotency_key = UUID-v4` per payment attempt
- Sends in request header
- Server stores result in `payment_results(key, status, amount, response)` table
- On retry: return stored result

**Order Event Flow:**
```
OrderService:
  BEGIN TRANSACTION
    INSERT INTO orders (id, status, ...) VALUES (...)
    INSERT INTO outbox (event_type='OrderPlaced', payload=...) VALUES (...)
  COMMIT
  
Debezium reads outbox → publishes to Kafka `orders.placed` topic

InventoryService: consumes `orders.placed` → DECRBY in Redis → publishes `inventory.reserved`
PaymentService: consumes `inventory.reserved` → charges payment → publishes `payment.completed`
DeliveryService: consumes `payment.completed` → assigns logistics
NotificationService: consumes all events → sends updates to user
```

---

## Q2. How would you design a real-time ride-matching system (like Uber/Ola)?

"Key requirements first: Match rider to nearest available driver in < 2 seconds, handle 1M active rides in peak, location updates from 500K active drivers every 5 seconds, strong consistency on ride assignment (no double-matching), multi-city support.

**Architecture:**

1. **Location Service:** Drivers send heartbeat with GPS coordinates every 5 seconds. Stored in Redis Geo (GEOADD). Not PostgreSQL — lat/lng updates are write-heavy, Redis is orders of magnitude faster.

2. **Matching Service:** On ride request, `GEORADIUS` query in Redis finds drivers within 2km radius. Applies business filters (driver rating, vehicle type, surge zone). Sorts by proximity.

3. **Assignment Service (Critical):** Distributed lock using Redis SET NX to atomically claim a driver. Prevents two riders from matching with the same driver.

4. **Trip Service:** Stores trip state in PostgreSQL. Manages trip lifecycle (requested → driver_assigned → started → completed).

5. **Real-time Communication:** WebSockets via a Socket.io gateway for driver location streaming to rider app."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Uber, Ola, Rapido, Swiggy (for delivery partner matching)

#### Deep Dive
**Driver location storage evolution:**
- Level 1: PostgreSQL with PostGIS extension. Works to ~100K active drivers.
- Level 2: Redis Geo. `GEOADD drivers {lat} {lng} {driverId}`. `GEORADIUS` returns nearby drivers. O(N+log(M)) where N is returned results.
- Level 3: S2 Geometry cells (Google's approach). Divide world into hierarchical cells. Driver's cell determines their index. Query is: "give me all drivers in cells overlapping the search radius." Scales to millions of active drivers.

**Race-free driver assignment (Redis atomic operations):**
```
// When assigning Driver D to Rider R:
SET driver:{driverId}:status "assigned" NX EX 30
// NX = only set if key doesn't exist (atomic check-and-set)
// EX 30 = auto-expire in 30 seconds (release lock if client dies)
// Returns OK if assignment succeeded, nil if driver already assigned
```

**Surge pricing architecture:** Event-driven: location updates → lambda to compute supply count by geohash → demand events → surge calculator compares supply/demand ratio per zone → surge factor stored in Redis → pricing service reads surge factor in real-time.

---

## Q3. Design a payment gateway (like Razorpay). What are the core architectural challenges?

"Core challenges: financial-grade consistency (no money lost, no double charges), regulatory compliance (PCI-DSS), third-party integrations (banks, card networks, UPI), extremely high availability, and 100% audit trail.

**Architecture:**

1. **Payment Ingestion Layer:** Handles incoming payment requests. Stateless, scales horizontally. Validates request, generates idempotency key, enqueues to payment processing queue.

2. **Payment Processing Engine:** Single-threaded, event-sourced processor per payment. Maintains the canonical state machine: initiated → bank_request_sent → bank_response_received → authorized → captured/failed.

3. **Bank Integration Layer:** Fans out to bank APIs (Axis, HDFC, ICICI) or UPI switch. Handles timeouts, retries, response parsing.

4. **Ledger Service:** Double-entry accounting ledger. Every rupee movement is a ledger entry. Immutable, append-only. Source of financial truth.

5. **Settlement Service:** Reconciles with bank files. Batch process that runs daily. Detects discrepancies between our records and bank records."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Razorpay, PhonePe, PayU, CRED, senior fintech roles

#### Deep Dive
**Dual-entry ledger:**
```sql
-- Transfer 100 INR from merchant M1 to system escrow
-- DEBIT merchant M1 wallet, CREDIT escrow account
INSERT INTO ledger_entries VALUES
  (txn_id, 'DEBIT',  account='merchant_M1',  amount=100, currency='INR'),
  (txn_id, 'CREDIT', account='system_escrow', amount=100, currency='INR');
-- Immutable: no UPDATE or DELETE on ledger_entries, ever
-- SUM(CREDIT) - SUM(DEBIT) across ALL accounts = 0 (always — mathematical guarantee)
```

**Payment state machine:** Each payment moves through states: `CREATED → PROCESSING → AUTHORIZED → CAPTURED / FAILED / REFUNDED`. State transitions are stored as events (event sourcing). Current state = replay of events. Any historical state queryable.

**Bank reconciliation:** Banks send settlement files (CSV/FTP) daily. Reconciliation service reads bank file, matches to our payment records, flags discrepancies:
- Our record shows `CAPTURED`, bank record shows `FAILED` → bank-side failure, reverse our charge
- No bank record for our payment → likely a late-settling transaction

**PCI-DSS compliance:** Card data (PAN, CVV) never touches our servers. Cards are tokenized at the card network level (Visa/Mastercard tokenization). We store tokens, not raw card data. Reduces our PCI scope dramatically.

---

## Q4. How would you design a feed architecture for a social platform (like Twitter/LinkedIn)?

"The classic fan-out problem: when Sachin Tendulkar (100M followers) posts a tweet, how do you deliver it to 100M feeds without making them wait?

Two approaches: **Fan-out on write** (push) and **fan-out on read** (pull).

**Fan-out on write:** When Sachin posts, enqueue a job to write the tweet ID to the feed of every follower. At 100M followers, this is 100M Redis writes — takes minutes. But reading the feed is instant (pre-computed).

**Fan-out on read:** Don't pre-populate feeds. When a user opens their feed, query the latest posts from everyone they follow, merge and sort. Simple write path, but reading is a cross-user query that doesn't scale.

**Twitter's hybrid approach (what I'd use):** Fan-out on write for users with < 1M followers (99.9% of accounts). Fan-out on read for celebrities (Sachin, Virat, Modi). Merge at read time: user's pre-computed feed + real-time query for celebrity posts."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Twitter, LinkedIn, Instagram-competing companies

#### Deep Dive
**Feed storage:** Redis SortedSet per user, scored by timestamp.
```
ZADD feed:{userId} {timestamp} {tweetId}
ZREVRANGE feed:{userId} 0 19  → get latest 20 tweets from user's feed
```

**Fan-out service:** Kafka consumer reads `PostCreated` events. For each post, queries follower list in batches of 1000. Issues Redis ZADD commands for each follower's feed. Horizontally scalable — multiple consumer instances partition by userId.

**Timeline cache eviction:** Feed only stores the latest N tweets (e.g., 800). On cache miss (user hasn't opened app in 30 days), rebuild feed from storage on demand.

**Facebook's approach (Ranked feed vs Chronological):** Raw feed is timeline sorted. Ranked feed applies an ML model to score each post and re-rank. The ML scoring service reads the raw feed, applies scores, returns top K posts. This is an additional layer on top of the timeline architecture.

---

## Q5. Walk me through designing a distributed rate limiter that works across multiple service instances.

"A single-instance rate limiter (in-memory) doesn't work when you have 50 instances of the API service — each instance has its own counter, so a user could make 50x the allowed requests by round-robining across instances.

**Solution: Centralized counter in Redis with Lua scripts for atomicity.**

Token bucket algorithm in Redis: Each user has a key `rate:{userId}`. The value is the current token count. A Lua script atomically: checks current tokens, refills tokens based on elapsed time, decrements if available, returns allowed/denied.

Using Lua for atomicity is critical — a Redis WATCH/MULTI/EXEC transaction is slower and more error-prone than a single atomic Lua execution."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Razorpay, Amazon (API Gateway), Zepto, any API-first company

#### Deep Dive
**Redis-based sliding window rate limiter:**
```lua
-- KEYS[1] = rate:{userId}, ARGV[1] = now_ms, ARGV[2] = window_ms, ARGV[3] = limit
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local limit = tonumber(ARGV[3])

-- Remove entries outside the sliding window
redis.call('ZREMRANGEBYSCORE', KEYS[1], 0, now - window)

-- Count entries in window
local count = redis.call('ZCARD', KEYS[1])

if count < limit then
    -- Add current request
    redis.call('ZADD', KEYS[1], now, now .. math.random())
    redis.call('EXPIRE', KEYS[1], math.ceil(window/1000))
    return 1  -- Allowed
else
    return 0  -- Rate limited
end
```

**Redis cluster for rate limiter scale:** For millions of users, shard the Redis cluster by userId. A consistent hash routes `rate:{userId}` to a specific shard. Each shard handles its users independently. Adding shards (consistent hashing) minimizes redistribution.

**Fallback on Redis failure:** If Redis is down, rate limiter fails open (allow requests) rather than fail closed (block all requests). The alternative — failing closed — creates a service outage when the rate limiter is down. Log that rate limiting is disabled and alert. Don't let rate limiter be a SPOF.

---

## Q6. How do you achieve zero-downtime database migrations in production?

"The core challenge: your application code is being updated (rolling deployment) while your DB schema is being changed. For a period, old code and new code run simultaneously against the same database. Your schema change must be backward compatible with both the old and new application.

**Expand-contract for DB migrations:**

Step 1 — **Expand:** Add new column `full_name` (nullable, no default required by old code). Run migration. Old code runs fine (ignores new column). New code writes to both `user_name` and `full_name`.

Step 2 — **Backfill:** Update all existing rows to populate `full_name = user_name`. Do this in batches to avoid locking the table.

Step 3 — **Switch reads:** Deploy app version that reads from `full_name`. Verify.

Step 4 — **Contract:** Remove `user_name` column (now unused). Run migration. Zero downtime."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Any company with continuous deployment

#### Deep Dive
**Critical: Never do these in production:**
- ADD NOT NULL COLUMN without default: Fails on any row that exists before the migration runs — table is locked
- DROP COLUMN with running code that uses it: Old code instances crash
- RENAME column: Old code reads old column name (fails), new code reads new column name — both can't coexist

**Large table migrations (avoid table locks):**
```
ALTER TABLE orders ADD COLUMN delivery_date TIMESTAMP;
-- On a 100M row table, this locks the table for minutes on MySQL

-- Alternative: gh-ost (GitHub's Online Schema Change)
-- Creates a ghost table, asynchronously copies data, swaps atomically
gh-ost --table=orders --alter="ADD COLUMN delivery_date TIMESTAMP" --execute
```

**Tools:** `gh-ost` (GitHub), `pt-online-schema-change` (Percona) for MySQL. PostgreSQL 12+ ADD COLUMN with default is instant (stores default in catalog, not per row). Always check if your DB version makes the operation safe.
