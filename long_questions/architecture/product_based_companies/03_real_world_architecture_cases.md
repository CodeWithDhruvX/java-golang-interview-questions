# 🌐 Real-World Architecture Cases — Product-Based Companies

> **Level:** 🔴 Senior
> **Asked at:** Amazon, Google, Flipkart, Uber, Swiggy, Razorpay, PhonePe

---

## Q1. Design a URL shortener (like bit.ly or TinyURL).

"I'll start by clarifying scope: 100M URLs shortened per day, 10B redirects per day (reads are 100x writes), shortened URL must be 7 characters, must be globally available, links must not expire unless configured.

**Core flow:** User submits long URL → we generate a short code → store the mapping → redirect requests hit the short URL → we look up and 301/302 redirect.

**HLD:**
1. **Write path:** Client → API Gateway → URL Shortener Service → Write to PostgreSQL + async cache warm to Redis
2. **Read path (critical, 10x more):** Client → Global CDN (cache redirects) → Redis cluster → PostgreSQL fallback

The 301/302 redirect choice matters: 301 (permanent) lets browsers cache the redirect locally, reducing server load but making analytics impossible. 302 (temporary) forces every redirect through our server, enabling click analytics but increasing load. Build.ly uses 302 for analytics."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart — standard HLD interview question

#### Deep Dive
**Short code generation — the core algorithm problem:**

```
Option 1: Hash-based
  MD5("https://google.com") → 128-bit hash → take first 7 chars
  Problem: Collision — two different URLs may have same 7-char prefix
  Fix: Check DB, if collision regenerate. Acceptable at low volume, fails at scale.

Option 2: Base62 encoding of auto-increment ID (best for most cases)
  DB auto-increment ID: 12345678
  Base62 chars: [0-9a-zA-Z] → 62 characters
  12345678 in base 62: "VR7Xa"
  7 chars of base62 = 62^7 = 3.5 trillion unique URLs — more than enough
  
  No collision possible — each ID is unique by design.
  Bloom filter to check duplicates before DB write (for same-URL deduplication).

Option 3: Pre-generated keys (like bit.ly approach)
  A Key Generation Service (KGS) pre-generates millions of 7-char codes.
  Stores them in a "keys_available" table.
  When a short URL is requested, mark one key as used.
  KGS can cache a batch of keys in memory → extremely fast.
  Trade-off: KGS is a dependency; requires distributed lock for concurrent access.
```

**Database schema:**
```sql
CREATE TABLE url_mappings (
    short_code VARCHAR(7) PRIMARY KEY,    -- "abc1234"
    long_url   TEXT NOT NULL,             -- "https://very-long-url.com/..."
    user_id    BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP,                 -- NULL = never expires
    click_count BIGINT DEFAULT 0,
    is_active  BOOLEAN DEFAULT TRUE
);
CREATE INDEX idx_long_url ON url_mappings(MD5(long_url));
-- MD5 index for deduplication check: "has this long URL been shortened before?"
```

**Read path at 10B redirects/day (115K RPS):**
```
CDN caches: Short code → long URL with TTL.
             ~80% of redirects served by CDN, never hits origin.
Redis cluster: Remainder served from in-memory cache.
               Key: "url:abc1234", Value: "https://..."
               TTL: 24 hours.
PostgreSQL: Only on cache miss (cold URLs, rarely accessed).

Read throughput: CDN handles 92,000 RPS, Redis 10,000 RPS, DB 3,000 RPS.
```

**Analytics architecture:**
```
Every redirect event → Kafka topic "url.redirected"
Fields: {short_code, timestamp, ip, user_agent, country}
Kafka consumer → ClickHouse (columnar analytics DB)
Dashboard queries ClickHouse for:
  - clicks per hour per URL
  - geographic distribution
  - referrer breakdown
```

---

## Q2. Design a notification system that sends email, SMS, and push notifications.

"A notification system is deceptively complex. Simple version: call an API. Production version: must handle 100M notifications/day, across 3 channels, with delivery guarantees, retries, rate limiting (don't spam a user), user preferences (user opted out of SMS), and observability.

**Architecture:**
1. **Notification API:** Receives requests from any internal service. Validates and enqueues.
2. **Channel Routers:** Separate queues per channel (email, SMS, push). Each with independent scaling.
3. **Channel Workers:** Pull from queues, call third-party providers (SendGrid for email, Twilio for SMS, FCM/APNs for push).
4. **Delivery Tracker:** Stores delivery status, handles retries on failure.
5. **Preference Service:** Before sending, checks if user has opted out of this channel/notification type."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Every mid-to-large product company — Swiggy, Zomato, Amazon, Flipkart

#### Deep Dive
**The notification flow with retry logic:**
```
Source Service: "Order OD123 placed for user U456"
    ↓ POST /notifications
Notification API:
    - Resolve user preferences: U456 → email: YES, SMS: YES, push: YES
    - Resolve template: "order_placed" → "Your order {{order_id}} has been placed!"
    - Publish 3 events:
        Kafka: email.notifications   → {to: "user@example.com", template: "order_placed", vars: {...}}
        Kafka: sms.notifications     → {to: "+919876543210", template: "order_placed_sms", vars: {...}}
        Kafka: push.notifications    → {device_ids: ["token1","token2"], payload: {...}}
    ↓
Email Worker (3 instances):          SMS Worker (3 instances):
    Consumes email.notifications         Consumes sms.notifications
    Calls SendGrid API                   Calls Twilio API
    On success → mark delivered          On failure → retry with backoff
    On failure → re-enqueue (DLQ)        After 3 retries → Dead Letter Queue
```

**Rate limiting — don't spam users:**
```
Per-user limits stored in Redis:
  key: "notif_rate:{user_id}:{channel}:{hour}"
  value: count of notifications sent this hour

Before sending:
  redis INCR + EXPIRE
  if count > limit → discard or schedule for next window

Global limits per channel (protect third-party API quotas):
  SendGrid: 100K emails/hour limit → throttle queue consumption accordingly
```

**Priority queues:**
```
HIGH priority: OTP, payment failure, order cancellation
NORMAL priority: Order placed, delivery update
LOW priority: Promotional, marketing

Separate Kafka topics per priority.
High-priority workers process immediately.
Low-priority workers have rate limits applied.
```

**Template management:**
```
Templates stored in DB with variable placeholders:
  "order_placed_email": "Hi {{first_name}}, your order #{{order_id}} worth ₹{{amount}} has been placed!"
  
Template service renders at send time using variable map.
A/B testing: different users get template_v1 vs template_v2 based on user_id hash.
```

**Delivery status tracking:**
```sql
CREATE TABLE notification_logs (
    id          UUID PRIMARY KEY,
    user_id     BIGINT NOT NULL,
    channel     ENUM('email', 'sms', 'push'),
    template    VARCHAR(100),
    status      ENUM('queued', 'sent', 'delivered', 'failed', 'bounced'),
    provider_id VARCHAR(255),     -- SendGrid message ID, Twilio SID
    sent_at     TIMESTAMP,
    delivered_at TIMESTAMP,
    error_code  VARCHAR(50)
);
```

---

## Q3. Design a search autocomplete system (like Google's or Amazon's search bar).

"When a user types 'iph' in Amazon's search bar, they should see ['iphone 15', 'iphone case', 'iphone charger'] within 100ms. This requires pre-computed candidates, not real-time search. At Amazon's scale, you can't query the product database on every keystroke.

**Architecture:**
1. **Offline pipeline:** Aggregate recent search queries. Rank by frequency, filter spam/offensive terms, compute top-K suggestions per prefix.
2. **Trie service:** Stores precomputed top suggestions for each prefix in a Trie (or Redis sorted sets).
3. **Low-latency API:** On each keystroke, query the trie/cache with the prefix. Return top 5-10 results.
4. **Personalization layer (optional):** Blend global suggestions with user's own search history."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Google, Amazon, Flipkart, Swiggy — classic HLD question

#### Deep Dive
**Data structure options:**

```
Option 1: Trie (Prefix Tree)
  "iph" → node → children: {"iphone", "iphoto", ...}
  Each node stores: top_k_suggestions = ["iphone 15", "iphone 14", "iphone case"]
  
  Storage: 26+10 children per node, massive memory for large vocabularies
  Lookup: O(length of prefix) — very fast

Option 2: Redis Sorted Set (more practical at scale)
  Key: "suggest:iph"
  Members: {"iphone 15", "iphone case", "iphoto"}
  Scores: search_frequency (higher = better suggestion)
  
  ZREVRANGE suggest:iph 0 9 → top 10 suggestions for "iph"
  
  On new popular search: ZINCRBY suggest:iphon 1 "iphone 15"
                         ZINCRBY suggest:iph 1 "iphone 15"
                         ZINCRBY suggest:ip 1 "iphone 15"
                         (update ALL prefix keys — write amplification)
```

**Offline pipeline for computing top-K per prefix:**
```
Input: Search query logs from last 7 days
  → MapReduce / Spark job:
      1. Group by query string
      2. Count frequency
      3. Filter: length < 3 chars, contains profanity, suspicious burst
      4. For each query, generate all prefixes: "iphone" → ["i", "ip", "iph", "ipho", "iphon", "iphone"]
      5. For each prefix, maintain top-K queries by frequency (min-heap, K=10)

Output: prefix → [top-10 queries with scores]
  Load into Redis cluster (or Trie service)
  
Schedule: Run daily at 2 AM, update Redis with new suggestions
```

**Personalization:**
```
Global suggestions: "iphone 15" (bought by 10M users)
User history: user_456 always searches for "samsung" products

Blend at query time:
  global_top_10 = query Redis for prefix
  user_history = query user's recent search history (last 30 days)
  
  merged = merge(global_top_10, user_history, alpha=0.7)
  -- 70% global ranking, 30% personal relevance
```

**Latency targets:**
```
Keystroke to suggestion visible: <100ms end-to-end
  Network (client to server): ~20ms (CDN PoP nearby)
  API processing: <5ms (Redis ZREVRANGE is O(log N + K))
  Rendering: ~10ms

Redis cluster handles 1M+ QPS for sorted set reads.
Deploy Redis cluster in same region as API servers to minimize latency.
```

---

## Q4. Design WhatsApp (or similar messaging system).

"WhatsApp at scale: 2B users, 100B messages per day. Core requirements: message delivery to offline users, end-to-end encryption, group messaging (up to 1024 members), message ordering per conversation, read receipts.

**HLD:**
1. **Chat Service:** Manages active WebSocket connections. Routes messages between online users.
2. **Message Store:** Persists all messages for offline delivery. Cassandra is ideal: high write throughput, time-ordered within a partition, easy TTL for message expiry.
3. **User Presence Service:** Tracks online/offline status. Redis pub/sub for real-time updates.
4. **Push Notification Service:** Delivers messages when user is offline (via APNs/FCM).
5. **Media Service:** Images/video → upload to S3, store S3 key in message, client downloads directly."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Uber, Amazon, PhonePe, Razorpay — complex HLD question

#### Deep Dive
**Message delivery states (the key concern in interviews):**
```
Sender sends message:
  1. Message reaches Chat Service → single tick (✓)
  2. Message stored in DB → recipient's queue populated
  3. Message delivered to recipient's device → double tick (✓✓)
  4. Recipient opens the conversation → blue ticks (✓✓ blue)

If recipient is offline:
  Chat Service → stores in Cassandra
  Push Notification Service → sends silent push to recipient's device
  When device comes online → fetches pending messages from Cassandra
  → marks as delivered
```

**Database design (Cassandra — why not relational DB):**
```sql
-- Cassandra table design
CREATE TABLE messages (
    conversation_id UUID,           -- Shard key: all messages per conversation on same node
    message_id      TIMEUUID,       -- Natural ordering by time (Cassandra TIMEUUID)
    sender_id       UUID,
    content         TEXT,           -- Encrypted ciphertext (E2EE)
    message_type    TEXT,           -- 'text', 'image', 'video'
    media_url       TEXT,           -- S3 URL if image/video
    created_at      TIMESTAMP,
    PRIMARY KEY (conversation_id, message_id)
) WITH CLUSTERING ORDER BY (message_id DESC)  -- Latest messages first
  AND default_time_to_live = 7776000;          -- 90-day TTL (WhatsApp keeps 3 months)

-- Query: get last 50 messages in a conversation
SELECT * FROM messages WHERE conversation_id = ? LIMIT 50;
-- O(1) — Cassandra partition lookup, then read top 50 from SSTable
```

**Group messaging — the fan-out problem:**
```
Group of 1000 members, Alice posts a message.

Naive approach: Alice's Chat Service sends to all 999 members.
  Problem: Chat Service is blocking for seconds on each group message.

Solution: Async fan-out via Kafka
  Alice → Chat Service → Kafka topic "group.messages"
  Fan-out Workers (separate pool) consume Kafka:
    For each online member → push via WebSocket
    For each offline member → store in their personal message queue
    
Message queue per user (for offline delivery):
  Redis List: "pending:{user_id}" → list of message IDs
  When user comes online → batch fetch messages → clear queue
```

**End-to-end encryption (simplified):**
```
WhatsApp uses Signal Protocol (X3DH + Double Ratchet).
Simplified explanation for interviews:
  - Alice and Bob each have a public/private key pair (generated on device)
  - Public keys uploaded to WhatsApp's key server at registration
  - When Alice sends to Bob:
      1. Alice fetches Bob's public key
      2. Generates a session key (X3DH handshake)
      3. Encrypts message with session key → AES-256-CBC
      4. Sends ciphertext to WhatsApp server
      5. WhatsApp stores/forwards the ciphertext — cannot read it
  - WhatsApp's servers never see plaintext — true E2EE
```

---

## Q5. Design a distributed cache system (like Redis).

"Designing a distributed cache means: deciding what data to cache, how to handle cache misses, what eviction policy to use, how to handle cache invalidation, and how to scale the cache cluster.

**Core patterns:**
- **Cache-Aside (Lazy Loading):** Application checks cache first. On miss → fetch from DB → populate cache. Most common pattern.
- **Write-Through:** Write to cache AND DB simultaneously. Cache always up-to-date. Slower writes.
- **Write-Behind (Write-Back):** Write to cache only. Write to DB asynchronously in batch. Fast writes, risk of data loss.
- **Read-Through:** Cache fetches from DB on miss automatically (cache handles the DB call).

For most web applications, Cache-Aside is the right default."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Amazon, Razorpay, Zepto — Redis internals + caching strategy questions

#### Deep Dive
**Cache-Aside implementation (with the thundering herd problem):**
```
Standard cache-aside:
  result = redis.GET("product:12345")
  if result is nil:               -- Cache MISS
      result = db.query("SELECT * FROM products WHERE id=12345")
      redis.SET("product:12345", result, EX=3600)    -- Cache for 1 hour
  return result

PROBLEM: Thundering herd
  If 10,000 concurrent requests all miss the cache at once (cold start / after expiry)
  → 10,000 simultaneous DB queries → DB overloads

SOLUTION: Cache lock / probabilistic early expiration
  On cache miss, only ONE request fetches from DB.
  Others wait (short poll) or return stale data.
  
  Redis SETNX-based lock:
    lock = redis.SET("lock:product:12345", "1", NX=True, EX=5)
    if lock:
        # I own the lock — fetch from DB and populate cache
        result = db.query(...)
        redis.SET("product:12345", result, EX=3600)
        redis.DELETE("lock:product:12345")
    else:
        # Someone else is refreshing — wait and retry, or return stale
        sleep(0.1)
        return redis.GET("product:12345")
```

**Eviction policies (LRU vs LFU):**
```
LRU (Least Recently Used): Evict the item accessed least recently.
  Good for: Workloads with temporal locality (recently accessed = likely to be accessed again).
  Bad for: Scanning large datasets (scan pollutes cache with one-time data).

LFU (Least Frequently Used): Evict the item accessed least often.
  Good for: Items that are always popular (product catalog hot items).
  Bad for: New items — they start with frequency 0, evicted before they get popular.

Redis 4.0+: True LFU implementation (Morris counter — approximate frequency tracking).
Redis default: allkeys-lru (evict any key using LRU when memory is full).
```

**Cache invalidation — the hardest problem:**
```
Strategy 1: TTL-based (most common)
  Accept stale data for TTL duration.
  Set TTL based on acceptable staleness: product price (60s), user profile (5min), static content (24h).

Strategy 2: Event-driven invalidation (stronger consistency)
  When product updated in DB → publish "product.updated" event to Kafka
  Cache invalidation consumer: redis.DELETE("product:{id}")
  Next read → cache miss → fresh data from DB.
  
Strategy 3: Versioned keys (zero stale reads)
  Cache key includes a version: "product:12345:v3"
  On update → increment version in DB
  Cache key changes → natural miss → fresh fetch
  Old keys expire via TTL
```
