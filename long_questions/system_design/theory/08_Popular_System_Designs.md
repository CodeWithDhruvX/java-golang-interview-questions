# 🔴 Popular System Designs — Questions 71–80

> **Level:** 🔴 Senior — The classic hallmark FAANG-level system design interview questions
> **Asked at:** Google, Amazon, Meta, Uber — system design flagship questions; also asked at Flipkart, Swiggy, Hotstar for senior roles

---

### 71. Design YouTube
"I'd start by clarifying scale: 2 billion logged-in users per month, 500 hours of video uploaded every minute, billions of views per day.

The core flows are **upload** and **stream**. For upload: user uploads raw video to object storage (S3 GCS) via a presigned URL → an upload completion event triggers a Transcoding pipeline → the video is encoded into multiple formats and resolutions (360p, 720p, 1080p, 4K) and stored in CDN-backed object storage. For streaming: user requests a video → CDN serves the chunks (HLS/DASH segments) from the nearest edge node.

The metadata DB (video title, description, tags, likes, views) goes to a relational DB (MySQL). Search is powered by Elasticsearch. The recommendation engine is a separate ML system."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Google (YouTube design is their gold standard interview), Netflix, Amazon Prime Video, Hotstar

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design YouTube
**Your Response:** I'd start by clarifying the scale: 2 billion logged-in users per month, 500 hours of video uploaded every minute, and billions of views per day. The core flows are upload and stream. For upload: the user uploads raw video to object storage like S3 via a presigned URL. An upload completion event triggers a transcoding pipeline that encodes the video into multiple formats and resolutions - 360p, 720p, 1080p, 4K - and stores them in CDN-backed object storage. For streaming: when a user requests a video, the CDN serves the chunks using HLS or DASH from the nearest edge node. The metadata database with video title, description, tags, likes, and views goes to a relational database like MySQL. Search is powered by Elasticsearch, and the recommendation engine is a separate ML system.

#### Indepth
YouTube system deep dive:
```
User → Upload Service → Raw Video (S3)
                ↓ event
            Transcoding Workers (Zookeeper-coordinated)
                ↓ produces
         [360p, 720p, 1080p, 4K] segments → S3 → CDN

User → API Server → Metadata DB (MySQL/Spanner)
                 → CDN URL for video chunks
User → CDN Edge → Adaptive Bitrate Streaming (HLS/DASH)
```

Key design decisions:
- **Transcoding at scale:** YouTube uses a distributed ffmpeg-based pipeline. Video is chunked (1-minute segments) and transcoded in parallel across hundreds of workers. Zookeeper coordinates work assignment. Output: multiple resolution variants as HLS `.m3u8` manifest + `.ts` segment files.
- **Adaptive Bitrate Streaming (ABR):** Player monitors network bandwidth and switches quality in real-time. Poor network → 360p. Fast WiFi → 1080p. This is why YouTube rarely buffers.
- **View Count at scale:** Views counter can't do an increment per view on a DB row — at billions/day, that's a hot row problem. Solution: counts are stored in Redis (fast in-memory INCR), batch-flushed to DB every few minutes.
- **Storage cost optimization:** Store originals in cold storage (S3 Glacier), hot serving via CDN. Most videos are long-tail (rarely watched) → CDN doesn't cache them → cheap cold storage serves them on cache miss.

---

### 72. Design WhatsApp
"WhatsApp serves 2 billion users sending 100 billion messages/day. The core MVP: send a message from User A to User B, deliver it, and show delivery/read receipts.

The key architectural decision is the connection model. WhatsApp maintains a **persistent WebSocket connection** from each client to a server. Messages are sent over this connection — no polling. When User A sends to User B: A sends message to its WebSocket server → server looks up which server B is connected to (via a routing table in Cassandra) → forwards message to B's server → B's server pushes to B.

If B is offline: the message is stored in a Cassandra mailbox. When B comes online and connects, the server delivers queued messages."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Meta (they own WhatsApp!), Amazon, Google — messaging system design

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design WhatsApp
**Your Response:** WhatsApp serves 2 billion users sending 100 billion messages per day. The core MVP is sending a message from User A to User B, delivering it, and showing delivery and read receipts. The key architectural decision is the connection model - WhatsApp maintains a persistent WebSocket connection from each client to a server. Messages are sent over this connection with no polling. When User A sends to User B, A sends the message to its WebSocket server, which looks up which server B is connected to via a routing table in Cassandra, forwards the message to B's server, and B's server pushes it to B. If B is offline, the message is stored in a Cassandra mailbox and delivered when B comes online.

#### Indepth
WhatsApp architecture deep dive:
- **Connection Layer:** Millions of long-lived TCP connections. Erlang/BEAM VM is famous for handling millions of concurrent lightweight processes — this is why WhatsApp built on Erlang. Go is another choice.
- **Message routing:** A consistent-hash ring determines which server any user is "homed" to. Other servers forwards messages to the home server. Redis/ZooKeeper tracks which server each user is connected to.
- **Message storage:** Messages stored in user's device (local SQLite DB). Servers store messages only until delivered. After delivery, server deletes them (privacy by design). Media goes to BLOB storage (CDN).
- **End-to-End Encryption (E2E):** Signal Protocol. Uses X3DH (Extended Triple Diffie-Hellman) key agreement and Double Ratchet for forward secrecy. WhatsApp's server never has the decryption keys.
- **Delivery receipts:** One tick = sent to server. Two ticks = delivered to device. Blue ticks = read. Each acknowledgment is a protocol message flowing back through the WebSocket connection.
- **Group Messaging:** For groups, sender's server expands the recipient list and sends individual messages to each member's server. WhatsApp groups cap at 1024 members to control fan-out cost.

---

### 73. Design Twitter (X)
"Twitter serves 200M DAU, 500M tweets/day. The hard problem is the **News Feed / Timeline** — each user follows hundreds of accounts; when they open Twitter, they need to see a ranked feed of recent tweets from followed accounts.

Two approaches: **Pull (fan-out on read)** — when user opens Twitter, query tweets from all followed users + rank + serve. Too slow for users following 1000+ accounts. **Push (fan-out on write)** — when a tweet is posted, push it to all followers' pre-computed timeline caches. Near-instant feed load. Problem: celebrities (Justin Bieber, 100M followers) → every tweet triggers 100M cache writes. Twitter uses a **hybrid approach**: regular users get fan-out on write to Redis timeline caches. Celebrity tweets are NOT pushed; they're fetched on read and merged with the pre-computed cache."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Twitter/X (their actual system), Meta, LinkedIn, any social platform company

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design Twitter (X)
**Your Response:** Twitter serves 200 million daily active users and 500 million tweets per day. The hard problem is the News Feed or Timeline - each user follows hundreds of accounts, and when they open Twitter, they need to see a ranked feed of recent tweets from followed accounts. There are two approaches: Pull or fan-out on read, where when a user opens Twitter we query tweets from all followed users and rank them - this is too slow for users following 1000+ accounts. Push or fan-out on write means when a tweet is posted, we push it to all followers' pre-computed timeline caches for near-instant feed load. The problem is celebrities like Justin Bieber with 100 million followers - every tweet would trigger 100 million cache writes. Twitter uses a hybrid approach: regular users get fan-out on write to Redis timeline caches, while celebrity tweets are not pushed but fetched on read and merged with the pre-computed cache.

#### Indepth
Twitter timeline system:
```
Tweet Posted by @justinbieber (100M followers)
      ↓
  Fan-out service
      ↓
 [Regular followers (10K)]: write tweet ID to each user's Redis timeline list
 [Celebrity followers]: Skip — too expensive
      
User opens Twitter:
  1. Fetch their pre-computed Redis timeline (push-model tweets)
  2. Fetch latest tweets from followed celebrities (pull-model)
  3. Merge and rank (by time + engagement score)
  4. Return sorted feed
```

Storage:
- **Tweets:** MySQL (sharded by tweet_id). Tweet IDs are Snowflake IDs (Twitter's distributed ID generation: 41 bits timestamp + 10 bits machine + 12 bits sequence → unique, time-sortable, distributed).
- **Social graph (follows):** Cassandra (user_id → list of followed user_ids) — wide-column is perfect for this.
- **Timelines:** Redis lists — `LPUSH timeline:{user_id} tweet_id`. `LRANGE timeline:{user_id} 0 799` to fetch last 800 tweets for ranking.
- **Search:** Elasticsearch with inverted index on tweet text. Twitter's real-time search requires near-instant indexing (tweets appear in search in <15 seconds).

---

### 74. Design Uber
"Uber's core challenge: match a rider with the nearest available driver in real-time, across millions of concurrent users globally.

The key data: driver locations update every 5 seconds. With 1M active drivers, that's 200K location updates/second. These can't go to a relational DB directly — too many writes. They go to a **geo-spatial hash** system (S2 cells or Geohash) backed by Redis.

When a rider requests a ride: find all drivers within a radius → rank by ETA (not just distance) → send the request to the top-ranked driver → if they decline, try next. This matching must happen in <1 second."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Lyft, Ola, Porter, Rapido

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design Uber
**Your Response:** Uber's core challenge is matching a rider with the nearest available driver in real-time across millions of concurrent users globally. The key data is that driver locations update every 5 seconds. With 1 million active drivers, that's 200,000 location updates per second. These can't go to a relational database directly due to too many writes - they go to a geospatial hash system like S2 cells or Geohash backed by Redis. When a rider requests a ride, we find all drivers within a radius, rank them by ETA not just distance, send the request to the top-ranked driver, and if they decline, we try the next. This matching must happen in under 1 second.

#### Indepth
Uber's architecture layers:
1. **Location Service:** Drivers send GPS updates every 5s via WebSocket. Location writes go to Redis geo index (`GEOADD drivers_online lng lat driver_id`). Redis Geo commands give nearby drivers within radius efficiently.
2. **Geospatial Indexing:** Redis uses Sorted Sets with Geohash-based scores. `GEORADIUSBYMEMBER drivers_online current_driver 5 km ASC COUNT 10` returns nearest 10 drivers.
3. **Matching Service:** Scoring function: `score = f(distance, driver_rating, surge_multiplier, ETA)`. Uber H3 hexagonal grid (by Uber H3 library) divides the earth into hierarchical hexagonal cells. Queries are cell-based — "drivers in current cell + neighboring cells".
4. **Surge Pricing:** Computed dynamically by comparing supply (available drivers) vs demand (pending ride requests) per H3 cell. Published to all apps in real-time.
5. **Trip Lifecycle:** State machine — `REQUESTED → DRIVER_ASSIGNED → DRIVER_ARRIVED → TRIP_STARTED → TRIP_ENDED → PAYMENT_PROCESSED`. Each state transition published as Kafka event.

---

### 75. Design Instagram
"Instagram: photo/video sharing. 1 billion MAU (Monthly Active Users), hundreds of millions of photos uploaded daily.

The core flows: upload media → process (filter, resize, CDN) → post to followers' feeds. The bottleneck is the feed: generating a personalized, ranked feed for a billion users.

Infrastructure: media uploaded to S3 → Lambda triggers resize pipeline → multiple sizes (thumbnail, low-res, full) stored back to S3 → CloudFront CDN serves globally. Feed generation: Instagram uses a trained ranking ML model that scores candidate posts by predicted engagement. The candidate pool is pulled from followed accounts' recent posts."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Meta, Snap, Pinterest, Twitter

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design Instagram
**Your Response:** Instagram is a photo and video sharing platform with 1 billion monthly active users and hundreds of millions of photos uploaded daily. The core flows are uploading media, processing it with filters and resizing, posting to followers' feeds, and serving those feeds. The bottleneck is generating a personalized, ranked feed for a billion users. The infrastructure works like this: media is uploaded to S3, Lambda triggers a resize pipeline that creates multiple sizes - thumbnail, low-res, and full - stored back in S3 and served globally via CloudFront CDN. For feed generation, Instagram uses a trained ML ranking model that scores candidate posts by predicted engagement. The candidate pool is pulled from followed accounts' recent posts.

#### Indepth
Instagram infrastructure highlights:
- **Photos → Cassandra:** Photo metadata (photo_id, user_id, caption, timestamp, location) stored in Cassandra — write-heavy, no complex joins needed.
- **Follower graph → Cassandra:** `user_id → [follower_ids]` and `user_id → [following_ids]`. Wide rows perfect for Cassandra.
- **Feed (original implementation):** Pure fan-out on write as Instagram scaled. When User A posts, a Celery task pushes `photo_id` to each follower's Redis feed list. For Instagram at early scale (~millions of users), this was manageable.
- **Feed (current):** Hybrid approach with ML ranking. Ranker uses signals: recency, engagement prediction, relationship strength, content type. Hundreds of features processed by a neural network ranking model.
- **Stories:** 24-hour TTL. Stored in Cassandra with TTL column. No need for explicit deletion — Cassandra handles TTL natively.
- **Instagram's original LAMP stack → evolved to Python (Django) services + PostgreSQL + Redis + Cassandra** — classic startup scale-up story.

---

### 76. Design a URL Shortener (like bit.ly)
"A URL shortener: input `https://very-long-url.com/page?params=many` → output `https://sho.rt/abc123`. Requirements: 100M URLs stored, 10B redirects/day (high read, low write).

The core service: **encode** input URL → **store** mapping in DB → **redirect** short URL to original. The shortcode generation is the key design decision. I use **Base62 encoding** (a-z A-Z 0-9 = 62 chars). A 6-character Base62 code gives 62^6 = ~56 billion unique codes. Generate a new auto-incrementing numeric ID, convert to Base62.

The redirect is a `301 Redirect` (permanent, browser caches) or `302 Redirect` (temporary, server handles every hit — needed for analytics)."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Amazon, Google, Flipkart — classic beginner-friendly system design question with depth in scalability discussion

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design a URL Shortener (like bit.ly)
**Your Response:** A URL shortener takes a long URL like https://very-long-url.com/page?params=many and outputs a short URL like https://sho.rt/abc123. The requirements are storing 100 million URLs and handling 10 billion redirects per day - that's high read, low write. The core service has three parts: encode the input URL, store the mapping in a database, and redirect the short URL to the original. For shortcode generation, I use Base62 encoding with characters a-z, A-Z, and 0-9. A 6-character Base62 code gives us 62 to the power of 6, which is about 56 billion unique codes. I generate a new auto-incrementing numeric ID and convert it to Base62. For the redirect, I use either a 301 redirect for permanent browser caching or a 302 redirect for temporary when we need analytics on every hit.

#### Indepth
URL shortener design decisions:

**Encoding approaches:**
1. **Counter + Base62:** Monotonically increasing DB auto-increment ID → convert to Base62. Cons: predictable sequence, reveals approximate creation order.
2. **MD5/SHA256 hash + truncate:** Take first 6 chars of hash. Cons: collisions (probability low but non-zero), requires collision check.
3. **Random Base62:** Generate random 6-char string, check uniqueness in DB. Cons: DB check on every write.
4. **Pre-generated key pool:** Background service generates random keys in bulk, stores in a `keys_available` table. URL shortener service picks one from the pool. Avoids collision checking at write time. Scales cleanly with Zookeeper-based key distribution.

**Performance at 10B redirects/day = ~115K QPS:**
- Cache popular short URLs in Redis (`GET sho.rt/abc123 → original URL`). 90% of traffic is repeat redirects to popular URLs → cache hit rate very high.
- `301` vs `302`: Use `302` if you want to count every redirect (analytics), run A/B tests, or re-target the short URL to a different destination later. Use `301` if you want to reduce server load (browser caches the redirect).

---

### 77. Design a file storage system (like Dropbox or Google Drive)
"Dropbox/Google Drive: store files reliably, sync across devices, support sharing. Design for 1B users, exabytes of storage.

The key insight: **files are chunked**. Large files are split into 4MB chunks, hashed (SHA-256), and each chunk stored independently in a distributed object store (S3). Two files sharing a paragraph? That chunk is stored once — deduplication by content hash. Only changed chunks are uploaded on sync — delta sync.

Metadata (file names, folder structure, ownership, share permissions) goes to a relational DB (PostgreSQL + sharded MySQL). The sync service uses a vector clock or Lamport timestamp to detect conflicts when two devices modify the same file offline."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Dropbox (their actual system!), Google, Microsoft, Box

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design a file storage system (like Dropbox or Google Drive)
**Your Response:** Dropbox and Google Drive need to store files reliably, sync across devices, and support sharing for potentially 1 billion users with exabytes of storage. The key insight is that files are chunked - large files are split into 4MB chunks, hashed with SHA-256, and each chunk is stored independently in a distributed object store like S3. This enables deduplication - if two files share a paragraph, that chunk is stored only once. Only changed chunks are uploaded during sync, which is called delta sync. Metadata like file names, folder structure, ownership, and share permissions goes to a relational database. The sync service uses vector clocks or Lamport timestamps to detect conflicts when two devices modify the same file offline.

#### Indepth
Dropbox's architecture:
- **Chunking + Deduplication:** Files split into content-defined chunks (Rabin fingerprinting — chunk boundaries at natural content breaks, not fixed byte positions). Each chunk identified by SHA-256 hash. Same chunk across multiple files = stored once. Client-side deduplication: before uploading chunk, check `POST /chunk/exists/{hash}`. If exists, skip upload — just reference it.
- **Block Server:** Receives chunks. Stores in S3. Records `{hash, s3_key}` in a block metadata DB.
- **Sync:** File changes detected by Dropbox desktop client file system watcher (inotify on Linux, FSEvents on macOS). Changed files rechunked, deltas computed, only changed chunks uploaded.
- **Vector Clocks for conflict detection:** Each device maintains a version vector. On sync, server compares. If two devices both modified the same file → conflict → both versions preserved (`file.txt` and `file (John's conflicted copy).txt`).
- **Notification system (for real-time sync):** Long polling or WebSocket from clients. Server pushes "file changed" notifications. Client pulls the delta of changed metadata.

---

### 78. Design a news feed (like Facebook)
"Facebook's News Feed is one of the most complex personalization systems ever built. At its core: when you open Facebook, you see a ranked, personalized feed of posts from friends, pages, and groups you follow.

The scale: 2B+ DAU, each user follows hundreds of friends/pages. Naively pulling posts from everyone you follow at read time is O(n * posts_per_friend) — way too slow. Facebook uses a hybrid fan-out: for normal users, pre-compute the feed and cache in Redis. Celebrities' posts are fetched at read time and merged.

The ranking is the hard part — Facebook trains massive ML models that predict "if user X sees post Y, will they react, comment, or share?" The top-scored posts form the feed."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Meta (their core product), LinkedIn, Twitter, any social platform

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design a news feed (like Facebook)
**Your Response:** Facebook's News Feed is one of the most complex personalization systems ever built. When you open Facebook, you see a ranked, personalized feed of posts from friends, pages, and groups you follow. The scale is massive - 2 billion daily active users, each following hundreds of friends or pages. Naively pulling posts from everyone you follow at read time would be way too slow. Facebook uses a hybrid fan-out approach: for normal users, they pre-compute the feed and cache it in Redis. For celebrities' posts, they fetch them at read time and merge them. The really hard part is the ranking - Facebook trains massive ML models that predict if a user will engage with a post. The top-scored posts form the feed.

#### Indepth
Facebook news feed pipeline:
1. **Candidate Generation:** Retrieve N candidates from:
   - Friends' recent posts (via social graph → last 24h posts for each friend)
   - Followed pages/groups
   - Sponsored content (ads)
   Total: thousands of candidates per user

2. **Filtering:** Remove already-seen posts, blocked users, reported content.

3. **Ranking:** ML model scores each candidate on: `P(engagement | user, post, context)`. Features: user-post relationship, historical engagement patterns, post freshness, media type. Light model (GBDT) for rough ranking → heavy model (DNN) for final top-k re-ranking.

4. **Diversity + Fairness:** Avoid showing 10 consecutive posts from same friend. Interleave content types. Apply content policies.

5. **Delivery + Caching:**
   - Computed feed stored in Redis list per user (top 500 posts as post_ids)
   - Feed pagination: user scroll → load next batch → trigger feed refresh computation for next page

**Facebook's EdgeRank (original algorithm) evolved into Neural News Feed Ranking (NNF) — a hundreds-of-features neural model, retrained daily on engagement signals.**

---

### 79. Design a video streaming service (like Netflix)
"Netflix serves 200M subscribers streaming HD/4K content concurrently. The challenge: serve different resolutions to users with different bandwidth, minimize startup latency, and eliminate buffering.

Architecture: content is **transcoded** once into dozens of formats (AVC, HEVC, AV1 codecs × multiple resolutions × multiple bitrates per resolution) and served via **Adaptive Bitrate Streaming** through a massive CDN (Netflix's own Open Connect CDN with servers in ISPs and IXPs globally). The player dynamically selects the best bitrate segment every few seconds.

The recommendation system (80% of what users watch comes from recommendations) is a hybrid collaborative filtering + content-based neural model, one of the most sophisticated in the industry."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Hotstar, Amazon Prime Video, YouTube, Jio Cinema

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design a video streaming service (like Netflix)
**Your Response:** Netflix serves 200 million subscribers streaming HD and 4K content concurrently. The challenges are serving different resolutions to users with different bandwidth, minimizing startup latency, and eliminating buffering. The architecture works like this: content is transcoded once into dozens of formats - different codecs like AVC, HEVC, and AV1, multiple resolutions, and multiple bitrates per resolution. This is served through adaptive bitrate streaming via a massive CDN. Netflix has its own Open Connect CDN with servers in ISPs and IXPs globally. The player dynamically selects the best bitrate segment every few seconds. The recommendation system, which drives 80% of what users watch, is a hybrid collaborative filtering and content-based neural model.

#### Indepth
Netflix's technical innovations:
- **Open Connect CDN:** Netflix doesn't use public CDNs like Cloudflare or Akamai (at scale). They partner with ISPs and IXPs to place Open Connect Appliances (OCAs) — dedicated servers pre-seeded with popular content. ~90% of Netflix traffic served from OCAs within the ISP's network. Zero internet transit cost.
- **Content Popularity Prediction:** Netflix knows which shows will spike (new season drop). They pre-seed OCAs with content before release — content pushed overnight to thousands of appliances.
- **Chaos Engineering:** Netflix's Chaos Monkey / Chaos Kong: kills EC2 instances, entire AZs, even entire AWS regions. If streaming still works, the resilience is proven. This started with "Chaos Monkey" and evolved into a chaos engineering practice.
- **Microservices at extreme scale:** Netflix has hundreds of microservices. API gateway (Zuul) handles client requests. Each service (title service, rating service, recommendation service) independently scalable. Hystrix circuit breakers prevent cascading failures.
- **A/B Testing at scale:** Netflix runs 1000s of A/B tests simultaneously — testing different AI-generated thumbnail images, UI flows, recommendation algorithms. A/B testing infrastructure is a first-class product.

---

### 80. Design an e-commerce platform (like Amazon)
"Amazon.com: product catalog (hundreds of millions of SKUs), search, cart, checkout, payment. The scale challenge during peak (Amazon Prime Day): orders at 100K/second.

Core microservices: **Product Service** (catalog, inventory), **Search Service** (Elasticsearch), **Cart Service** (Redis — sessions), **Order Service** (Postgres), **Payment Service** (idempotent, heavily tested), **Notification Service** (async).

The hardest part: **inventory management** at checkout. Two users can't buy the last item simultaneously. I use **optimistic locking** (try to decrement inventory, check version, retry if conflict) or **database row lock** (`SELECT FOR UPDATE`) to prevent overselling during the critical purchase flow."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (they literally built this), Flipkart, Meesho, Myntra, Ajio

### How to Explain in Interview (Spoken style format)
**Interviewer:** Design an e-commerce platform (like Amazon)
**Your Response:** Amazon.com has a product catalog with hundreds of millions of SKUs, plus search, cart, checkout, and payment. During peak events like Amazon Prime Day, they handle orders at 100,000 per second. The architecture uses core microservices: Product Service for catalog and inventory, Search Service using Elasticsearch, Cart Service using Redis for sessions, Order Service with Postgres, Payment Service which is idempotent and heavily tested, and Notification Service for async operations. The hardest part is inventory management at checkout - two users can't buy the last item simultaneously. I use optimistic locking where I try to decrement inventory and check the version, retrying if there's a conflict, or database row locks with SELECT FOR UPDATE to prevent overselling during the critical purchase flow.

#### Indepth
E-commerce architecture deep dive:

**Product Catalog:**
- Structured data (category, price, brand) in MySQL/PostgreSQL
- Unstructured (descriptions, specifications) in DynamoDB or MongoDB
- Images in S3 + CloudFront CDN

**Search:**
- Elasticsearch with custom scoring: relevance × popularity × stock availability × personalization
- Typeahead: pre-computed suggestions in Redis sorted sets
- Faceted filtering: aggregation queries on ES

**Inventory at Scale:**
- Reservations pattern: request places a "soft hold" on inventory for 15 minutes during checkout. If payment succeeds, hold becomes a deduction. If payment fails or expires, hold is released.
- Flash sales (e.g., iPhone launches): overselling prevention with Redis atomic operations. `DECR inventory:product_id` — Redis is single-threaded, DECR is atomic. If result < 0 → `INCR` back and reject.

**Checkout flow (Saga):**
```
1. Validate cart items + prices (current snapshot)
2. Create order (PENDING) — local transaction
3. Reserve inventory → event: InventoryReserved
4. Charge payment → event: PaymentSuccess
5. Confirm order (CONFIRMED) → event: OrderConfirmed
6. Trigger fulfillment/shipping → event: FulfillmentStarted

If step 4 fails:
  Compensate step 3: ReleaseInventory
  Order status → PAYMENT_FAILED
```

**Why Amazon built DynamoDB:** During 2004 holiday season, Amazon's Oracle DB couldn't handle the load. Engineers built a proprietary system (Dynamo) optimized for Amazon's exact usage pattern. The 2007 Dynamo paper became one of the most influential distributed systems papers ever written.
