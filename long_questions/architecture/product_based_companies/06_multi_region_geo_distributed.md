# 🌍 Multi-Region & Geo-Distributed Architecture — Product-Based Companies

> **Level:** 🔴 Senior / Staff / Principal
> **Asked at:** Google, Amazon, Uber, PhonePe, Razorpay — senior engineering and principal/staff roles

---

## Q1. Why would you build a multi-region architecture? What are the trade-offs?

"A multi-region architecture deploys your application and data across multiple geographic regions (e.g., AWS Mumbai + AWS Singapore + AWS Frankfurt). The reasons: **(1) Latency** — serve users from nearby regions for faster responses. **(2) Availability** — if an entire AWS region goes down (rare but happens), another region serves traffic. **(3) Data residency** — regulatory requirements may mandate user data stays in-country.

The trade-offs are significant: **data replication lag** (eventual consistency between regions), **cross-region write conflicts** (multi-leader replication complexity), **cost** (double or triple infrastructure), and **operational complexity** (traffic routing, failover automation, consistency models per data type)."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Google, Amazon, any company serving global traffic

#### Deep Dive
**Active-active vs active-passive:**
```
Active-passive (simpler, more common):
  Primary region (Mumbai): handles ALL traffic
  Secondary region (Singapore): replication target only, serves traffic only if Mumbai fails
  
  Failover: Route53 health checks detect Mumbai failure → switch DNS to Singapore
  Recovery time: 30 seconds to 2 minutes (DNS TTL + health check interval)
  Data loss risk: Replication lag at time of failure (typically < 1 second for synchronous)

Active-active (complex, high availability):
  Both Mumbai and Singapore handle traffic simultaneously
  Writes in each region replicated to the other
  Read traffic served from nearest region
  
  Challenge: User in Mumbai writes data. User in Singapore reads it 50ms later.
  If replication hasn't caught up → stale read (eventual consistency).
  
  Conflict resolution needed: What if same record updated in both regions simultaneously?
  Solution: Per-record ownership "user X's data always writes to Mumbai, routed there from anywhere"
```

**Latency reduction — the case for multi-region:**
```
Global CDN (CloudFront/Fastly): Static content served from nearest edge (100+ PoPs worldwide)
  → HTML/CSS/JS/images: user in Chennai hits a Chennai CDN node, not Mumbai origin

API calls: Cannot be cached by CDN (dynamic, personalized)
  → Round-trip time: Chennai to Mumbai: ~20ms, Chennai to US East: ~200ms
  → 10x latency difference for APIs → multi-region reduces this to 20ms everywhere

Dynamic API geo-routing:
  User in South India → Route to Mumbai region (20ms)
  User in Singapore   → Route to Singapore region (5ms)
  User in Europe      → Route to Frankfurt region (10ms)

AWS Global Accelerator: Routes to nearest AWS region by anycast routing.
  (Standard DNS routing has latency; anycast skips it.)
```

**Data that's easy vs hard to geo-distribute:**
```
EASY to multi-region:
  Static assets: CDN + S3 replication
  Read-heavy data with acceptable staleness: Product catalog, public content
  User sessions: Store in nearest Redis, replicate async
  
HARD to multi-region:
  Financial transactions: Require strong consistency (no double spends, no lost credits)
  Inventory counts: Race conditions in multi-region distributed decrements
  Order placement: Must be globally consistent before confirming to user
  Authentication tokens: Revocation must propagate before access granted

Strategy: Keep "hard" data in one primary region with read replicas.
          Keep "easy" data fully geo-distributed.
```

---

## Q2. How does AWS Route53 enable region failover? Explain health checks and failover routing.

"Route53 is AWS's DNS service with programmable routing. For multi-region failover, Route53 continuously health-checks your endpoints. When the primary region fails health checks, Route53 automatically updates DNS responses to point to the backup region — without any human intervention.

The key parameters: health check interval (10 or 30 seconds), failure threshold (how many consecutive failures before marking unhealthy), and DNS TTL (how quickly clients pick up the DNS change)."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Amazon, companies running AWS-based multi-region setups

#### Deep Dive
**Route53 failover routing policy:**
```
DNS record: api.myapp.com

Primary record:
  Type: A
  Routing: FAILOVER → PRIMARY
  Value: 15.207.XXX.XXX  (Mumbai ALB IP)
  Health check: HTTP GET https://api.myapp.com/health every 10 seconds
  TTL: 60 seconds

Secondary record:
  Type: A
  Routing: FAILOVER → SECONDARY
  Value: 13.215.XXX.XXX  (Singapore ALB IP)
  Health check: HTTP GET https://api.myapp.com/health every 10 seconds
  TTL: 60 seconds

Behavior:
  Mumbai health check passes → Route53 returns Mumbai IP → all traffic to Mumbai
  Mumbai health check fails 3 times → Route53 marks Mumbai unhealthy
  → Returns Singapore IP → all traffic to Singapore
  
  Failover time: 3 × 10 seconds (health check) + 60 seconds (DNS TTL propagation)
  = ~90 seconds to full failover
```

**Latency-based routing (for active-active):**
```
Route53 latency-based routing:
  Record for Mumbai: 15.207.XXX.XXX → serves users whose lowest latency is to Mumbai
  Record for Singapore: 13.215.XXX.XXX → serves users whose lowest latency is to Singapore
  Record for Frankfurt: 52.XXX.XXX.XXX → serves users whose lowest latency is to Frankfurt

Route53 measures latency from user's IP to each region using AWS's latency data.
→ User in Chennai: Mumbai record (lowest latency)
→ User in Jakarta: Singapore record
→ User in Germany: Frankfurt record

Combined with health checks:
  If Mumbai becomes unhealthy → Route53 removes it from latency responses
  → Chennai users fall over to Singapore until Mumbai recovers
```

**Health check endpoint (what to check):**
```go
// Health check endpoint — must check actual dependencies, not just "server is running"
func healthHandler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    
    // Check DB connectivity
    if err := db.PingContext(r.Context()); err != nil {
        http.Error(w, "db_unavailable", 503)
        return
    }
    
    // Check Redis connectivity  
    if err := redisClient.Ping(r.Context()).Err(); err != nil {
        http.Error(w, "cache_unavailable", 503)
        return
    }
    
    // Check critical downstream services
    if !paymentServiceReachable() {
        http.Error(w, "payment_service_unavailable", 503)
        return
    }
    
    // All healthy
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status": "healthy",
        "region": os.Getenv("AWS_REGION"),
        "latency_ms": time.Since(start).Milliseconds(),
    })
}
// Route53 health check: GET /health → 200 = healthy, 503 = unhealthy
```

---

## Q3. How do you handle data replication across regions? What consistency guarantees can you offer?

"Cross-region replication introduces latency — the speed of light from Mumbai to Singapore is ~40ms round-trip. Any synchronous replication adds at least this latency to every write. So you must choose: accept writes being 40ms slower for strong consistency, or allow asynchronous replication for lower latency at the cost of eventual consistency.

The practical rule: **financial data uses synchronous replication to one backup region** (strong consistency, some write latency penalty). **User-generated content and behavioral data uses asynchronous replication** (eventual consistency, no write latency penalty)."

#### Company Context & Level
**Level:** 🔴 Staff | **Asked at:** PhonePe, Razorpay, Amazon, Uber — data architecture at scale

#### Deep Dive
**RDS Multi-AZ vs Multi-Region:**
```
Multi-AZ (within region, different Availability Zones):
  Primary DB in AZ-a, Synchronous replica in AZ-b
  Failover: <60 seconds → Multi-AZ promoted automatically
  RPO: 0 (synchronous — no data loss)
  RTO: ~60 seconds
  Cross-AZ latency: <1ms → synchronous is cheap
  
Multi-Region Read Replica (async):
  Primary DB in Mumbai, Async replica in Singapore
  Replication lag: typically 50-500ms (depends on write rate and cross-region bandwidth)
  Use case: Read traffic in Singapore served locally; Singapore cannot accept writes
  Failover: Manual promotion → now Singapore becomes primary (RPO: replication lag at failure time)
  
Multi-Region Active-Active (hard to achieve with RDS):
  → Use CockroachDB or YugabyteDB (distributed SQL with native multi-region support)
  → Or Aurora Global Database: primary region, up to 5 secondary regions, <1 second replication lag
```

**Aurora Global Database — practical for most companies:**
```
Architecture:
  Primary cluster: Mumbai
    → All reads + writes
    → Writes committed locally, then replicated
    
  Secondary clusters: Singapore, Frankfurt (read-only)
    → Read traffic served locally (no latency penalty for reads)
    → Replication lag: typically 10-100ms
    → If Mumbai fails: Managed Failover promotes Singapore to primary in < 1 minute

Use cases at companies like Razorpay:
  Indian users: Read from Mumbai primary (local reads) + write to Mumbai primary
  Singapore users: Read from Singapore replica (fast local reads) + write to Mumbai (50ms penalty)
  → Payments are rare but must be consistent → acceptable to route writes to primary
  → Product catalog reads happen constantly → serve from local replica
```

**CRDT for conflict-free multi-region writes (advanced):**
```
Problem: Shopping cart updated simultaneously in Mumbai (add shoes) and Singapore (add shirt).
  After replication: conflict — which version wins?

CRDT (Conflict-Free Replicated Data Types) — mathematically merge without conflicts:
  OR-Set (Observed-Remove Set) for shopping carts:
    Mumbai adds: {shoes, v1_mumbai}
    Singapore adds: {shirt, v1_singapore}
    After merge: {shoes, shirt} — both applied, no conflict
    
    Remove semantics: "remove shoe" only removes shoe if you've seen the "add shoe" operation
    → Prevents accidentally removing items added concurrently

Used by: Riak (DynamoDB predecessor), Redis CRDT modules, collaborative tools (Google Docs uses OT/CRDT).

When to use: Only for data types where merge is semantically meaningful.
  Counters, sets, shopping carts → good CRDT candidates.
  Account balances, inventory → NOT good candidates (CRDTs would allow negative balances).
```

---

## Q4. How would you design an architecture for a geo-distributed payment system (like PhonePe serving India and Southeast Asia)?

"A payment system has uniquely strict requirements: financial-grade consistency (no double charges, no money loss), regulatory compliance per country, and low latency for user experience.

**Key principle:** Financial records must be strongly consistent. A user's wallet balance cannot be eventually consistent — if two concurrent reads return different balances and both result in deductions, you've lost money.

**Architecture:** Financial transactions are always processed in a single authoritative region (the 'golden region' for that user's account), even if the user's request originates from another region. Read-heavy operations (transaction history, account dashboard) can be served from regional replicas."

#### Company Context & Level
**Level:** 🔴 Staff/Principal | **Asked at:** PhonePe, Razorpay, Juspay — fintech architecture roles

#### Deep Dive
**Account sharding by geography:**
```
User registration determines their "home region":
  Indian user    → account_home = "in-mumbai"
  Singapore user → account_home = "sg-singapore"
  European user  → account_home = "eu-frankfurt"

Account home region stored in a global directory service (lightweight, globally replicated):
  directory.lookup(user_id) → "in-mumbai"  (cached aggressively, changes rarely)

Payment request from Singapore user:
  1. Request hits Singapore API gateway
  2. API gateway queries global directory: user_456 → home = "in-mumbai"
  3. API gateway routes payment processing to Mumbai
  4. Mumbai processes transaction (strong consistency, single DB writes)
  5. Mumbai returns success, response sent back to Singapore API Gateway
  6. Singapore gateway returns success to user
  
Total added latency: Singapore ↔ Mumbai round trip (~80ms)
Acceptable for payments (users expect 1-3 seconds for payment processing)
```

**Data residency compliance (GDPR, India's PDPB):**
```
Legal requirements:
  India PDPB: Sensitive Indian user data cannot leave India
  GDPR (Europe): European user personal data cannot leave EU without safeguards
  
Implementation:
  Separate AWS accounts per regulatory jurisdiction
  VPC → no peering that would allow data to flow across jurisdictions
  Encryption keys stored in AWS KMS per region (key for Indian data in ap-south-1 only)
  
  Shared services that handle PII:
    User Service: sharded per region, no cross-region PII queries
    Analytics pipelines: anonymized/aggregated data before flowing cross-region
    
  Log aggregation: Must anonymize logs before sending to global SIEM
    PII scrubbing pipeline: replace field "user_email" → "REDACTED" in logs before cross-border
```

**Disaster recovery — the detailed runbook:**
```
RTO target: 15 minutes (how long the system can be "down" before resuming service)
RPO target: 0 (no financial data loss acceptable — synchronous replication to DR region)

DR setup:
  Production: Mumbai (ap-south-1) — Aurora Primary, all write traffic
  DR: Hyderabad (ap-south-2) — Aurora Global Read Replica, 10-100ms lag

Failover procedure (15-minute target):
  T+0:    Mumbai region health checks fail × 3
  T+1:    PagerDuty alert to on-call SRE
  T+3:    SRE confirms region failure (not a transient issue or false alarm)
  T+5:    SRE initiates Aurora Global Failover → Hyderabad promoted to primary
  T+8:    DNS updated (Route53 ALB DNS → Hyderabad ALB)
  T+10:   Application services in Hyderabad scaled out to handle full traffic
  T+12:   Traffic validation: synthetic transaction tests pass
  T+15:   Incident declared stable, communication sent to customers
  
  RPO: the ~10-100ms replication lag at time T+0. For financial systems, this is
       treated as "potentially lost transactions" requiring reconciliation after recovery.
```

---

## Q5. What is a CDN (Content Delivery Network) and how does it work? When should you use one?

"A CDN is a globally distributed network of edge servers that cache and serve content from the location nearest to the user. Instead of every request going to your origin servers in Mumbai, users in Chennai hit a CDN PoP in Chennai (~1ms), while users in London hit a CDN PoP in London (~2ms).

CDNs are essential for: static assets (JS/CSS/images), video streaming, large file downloads, and any globally distributed application. They're optional for highly personalized, dynamic API responses (though CDNs increasingly support this too via edge computing)."

#### Company Context & Level
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Swiggy, Zomato, any company with consumer mobile apps

#### Deep Dive
**CDN cache hierarchy:**
```
Origin server (Mumbai): The authoritative source. Only accessed on cache miss.
Regional PoP (Point of Presence): First level of cache. 100+ locations worldwide.
Edge PoP: Even closer to user. Major cities in each country.

Request flow:
  User in Chennai requests: GET https://cdn.swiggy.com/menu/restaurant_12345.json
  
  1. DNS resolves cdn.swiggy.com → nearest EdgePoP IP (anycast routing)
  2. Chennai EdgePoP: cache HIT → return directly (~1ms)
  3. Chennai EdgePoP: cache MISS → query Mumbai origin → cache response → return (~25ms)
  4. Next Chennai user: cache HIT → 1ms

Cache control headers (from origin, controls CDN behavior):
  Cache-Control: public, max-age=3600      → CDN caches for 1 hour, shared
  Cache-Control: private, no-store         → CDN must NOT cache (private user data)
  Cache-Control: public, s-maxage=300      → CDN caches 5 min, browser caches differently
  Surrogate-Control: max-age=86400         → CDN-specific header (Fastly, Varnish)
```

**Cache invalidation at CDN level:**
```
Problem: Restaurant menu changes → CDN caches old version for 1 hour → users see stale menu

Solutions:
  1. Short TTL (aggressive, simple):
     max-age=60 → users see stale data for max 60 seconds. Refresh happens naturally.
     Cost: 60x more origin traffic vs 1-hour TTL.

  2. Cache purge on update (surgical, complex):
     When menu updates → call CloudFront/Fastly API to purge specific key:
     cf.CreateInvalidation(paths=["/menu/restaurant_12345.json"])
     → CDN drops cached version → next request fetches fresh
     
     Cost: Purge APIs have rate limits (CloudFront: 1000 paths/month free, then paid)
     
  3. Cache busting via versioned URLs (best practice):
     /static/bundle-v3.2.4.js → URL includes content hash
     When file changes → hash changes → new URL → no cache conflict
     → Set max-age: 31536000 (1 year) — URL will never change if content doesn't change
     
  4. Surrogate keys / Cache tags (Fastly, Cloudflare):
     Tag cached responses: restaurant_12345_menu → tag invalidates all tagged entries
     → Purge by tag instead of individual URLs
```
