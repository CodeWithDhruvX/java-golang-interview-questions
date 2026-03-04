# 🔄 Reliability & Resilience Patterns — Questions 1–10

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Amazon, Google, Flipkart, Swiggy, Razorpay — any company that cares about uptime

---

### 1. What is reliability and how is it measured?

"Reliability is the probability that a system **performs its required functions correctly over a specified time period** under specified conditions.

In SRE terms, reliability is measured by availability: the percentage of time the system is operational. Four nines (99.99%) means ~52 minutes of downtime per year. Five nines (99.999%) means ~5 minutes per year.

The uncomfortable truth: 100% availability is impossible. Infrastructure fails. Code has bugs. Dependencies are unreliable. Reliability engineering is the practice of designing systems that **fail gracefully** — fail in expected, contained, recoverable ways instead of catastrophically."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** All companies with production systems

#### Indepth
Availability tiers:
| Availability | Annual Downtime | Monthly Downtime |
|-------------|----------------|-----------------|
| 99% ("two nines") | 87.6 hours | 7.3 hours |
| 99.9% ("three nines") | 8.76 hours | 43.8 minutes |
| 99.99% ("four nines") | 52.6 minutes | 4.4 minutes |
| 99.999% ("five nines") | 5.26 minutes | 26 seconds |

**Mean Time Between Failures (MTBF):** How often failures occur. Long MTBF = reliable system.
**Mean Time To Recover (MTTR):** How long to recover from a failure. Short MTTR = resilient system.

The SRE insight: Improving MTTR is often more valuable than improving MTBF. You can't prevent all failures, but you can detect and recover from them faster. Netflix's Chaos Engineering deliberately injects failures to improve MTTR — the team is always trained to handle failures.

---

### 2. What is fault tolerance vs high availability?

"**Fault tolerance** is the ability to **continue operating correctly even when some components fail** — the system degraded but didn't stop. Full fault tolerance means zero downtime even during failures. Space shuttle computers are fully fault-tolerant: 5 computers vote on every decision; 4 out of 5 can fail and the shuttle still operates.

**High availability (HA)** is a slightly weaker guarantee: the system is **available most of the time** with minimal unplanned downtime. HA systems may have brief interruptions during failover (seconds to minutes) but recover quickly. Active-passive database replication is HA: primary fails, replica is promoted (30 seconds of downtime), then fully operational again.

Fault tolerance is more expensive — it requires duplicate active components that can transparently absorb failures. HA allows brief recovery windows. Most production systems target HA (99.99%) rather than full fault tolerance."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Companies designing distributed systems

#### Indepth
HA design patterns:
1. **Active-Passive:** Primary handles all traffic. Standby takes over on failure. Simple but has failover delay.
2. **Active-Active:** Multiple nodes handle traffic simultaneously. No failover delay, but requires conflict resolution for writes.
3. **N+1 redundancy:** N components needed, N+1 running. One failure is absorbed transparently.
4. **Geographic redundancy:** Multiple data centers in different regions. Protects against regional failures.

**Recovery objectives:**
- **RTO (Recovery Time Objective):** Maximum acceptable time for recovery. "We must be back in 4 hours."
- **RPO (Recovery Point Objective):** Maximum acceptable data loss. "We can lose at most 5 minutes of data."

Higher availability requires lower RTO and RPO, which costs more. The business defines the acceptable thresholds — the architecture is designed to meet them.

---

### 3. What is the bulkhead pattern?

"Bulkheads are a ship design pattern — a ship is divided into watertight compartments so a breach in one compartment doesn't sink the ship. In software, the bulkhead pattern **isolates system resources so a failure in one area doesn't cascade to others**.

The problem: A single thread pool serves all request types. A slow external service causes threads to block waiting for timeouts. Thread pool exhausts — all requests starve, including fast ones that have nothing to do with the slow service.

Bulkhead solution: Separate thread pools (or connection pools or semaphores) for each external dependency. The payment service API has its own thread pool (max 20 threads). If payment is slow and all 20 are exhausted, requests to the delivery service (separate pool, max 10 threads) are unaffected."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart — resilience pattern discussions

#### Indepth
Implementation options:
1. **Thread pool isolation:** Separate executor for each external call. Hystrix's default approach.
```java
// Separate thread pools for each dependency
@HystrixCommand(threadPoolKey = "paymentServicePool",
                threadPoolProperties = {
                    @HystrixProperty(name = "coreSize", value = "10"),
                    @HystrixProperty(name = "maxQueueSize", value = "5")
                })
public PaymentResult chargeCard(PaymentRequest req) { ... }
```

2. **Semaphore isolation:** Limit concurrent callers instead of pool threads. Lighter weight but blocking.

3. **Connection pool isolation:** Separate DB connection pools per service feature area. The reports feature can't starve the order feature of DB connections.

4. **Kubernetes resource quotas:** Each service gets CPU and memory limits. A runaway service can't consume all node resources and starve neighbors.

---

### 4. What is retry with exponential backoff and jitter?

"Retry is the simple act of trying again after a failure. But naive retry (retry immediately, retry constantly) can make failures worse — imagine 1000 services all retry the same failed endpoint simultaneously after a 1-second failure: you get a **retry storm** that prevents the endpoint from recovering.

**Exponential backoff**: Wait 1s after first failure, 2s after second, 4s after third, 8s after fourth. Dramatically reduces load on a recovering service.

**Jitter** (the critical addition): Add randomness to the backoff interval. Instead of all clients waiting exactly 4 seconds, they wait 3.2–4.8 seconds (random). This prevents the **thundering herd problem** where all retried clients hit the service at exactly the same moment."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company designing distributed systems

#### Indepth
Retry algorithm:
```go
const maxRetries = 5
const baseDelay = 100 * time.Millisecond

func retryWithBackoff(fn func() error) error {
    for attempt := 0; attempt < maxRetries; attempt++ {
        err := fn()
        if err == nil { return nil }
        if isNonRetryable(err) { return err }  // Don't retry 4xx errors
        
        // Exponential backoff with full jitter
        delay := baseDelay * (1 << attempt)  // 100ms, 200ms, 400ms, 800ms, 1600ms
        jitter := time.Duration(rand.Int63n(int64(delay)))
        time.Sleep(jitter)
    }
    return lastErr
}
```

What NOT to retry:
- **4xx client errors:** 400 Bad Request, 404 Not Found — retrying won't help (bad input)
- **402 Payment Required, 429 Too Many Requests (with Retry-After):** Handle specifically
- **Non-idempotent operations:** Retrying `POST /payments` can cause double charges — use idempotency keys

**Circuit breaker + retry combo:** Use retry for transient failures (network hiccup). Use circuit breaker to stop retrying when the downstream is fundamentally down. Together: retry handles momentary blips; circuit breaker prevents retry storm on sustained failure.

---

### 5. What is a fallback strategy?

"A fallback is a **predefined response** returned when the primary path fails — typically when a circuit breaker is open or a timeout occurs. Instead of propagating an error to the user, the system returns something: cached data, a default value, a degraded response, or a friendly error message.

Fallback strategies (from best to worst user experience): (1) Return cached data from a recent successful response, (2) Return default/generic data (popular products, empty list), (3) Return a degraded feature (show order list but not order details), (4) Return a human-readable error (much better than a 500 crash).

At Netflix, if the recommendation service is down, the fallback shows the most popular movies globally. Not personalized, but users see something instead of a broken page."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Companies designing resilient user-facing systems

#### Indepth
Fallback implementation options:
1. **Cached response:** Store the last successful API response in Redis. If Service B is down, return the cached response (may be stale, but acceptable).
```java
@HystrixCommand(fallbackMethod = "getFromCache")
public Product getProduct(String id) { /* calls Product Service */ }

public Product getFromCache(String id) {
    return redis.get("product:" + id);  // Return stale cache
}
```

2. **Stub/default response:** Return empty list, default settings, zero counts.
3. **Degraded mode:** Skip non-critical features. Show order list without payment details if payment service is down.
4. **Fail fast with user message:** "Recommendations are temporarily unavailable. Try again in a few minutes."

**Chaos engineering:** Netflix's Chaos Monkey randomly terminates production instances to ensure fallbacks actually work. Fallback code is often forgotten, untested, and broken. The only way to know it works is to test it in production.

---

### 6. What is a health check and how do you implement it?

"A health check is an **endpoint that reports whether a service is functioning correctly**. It's how load balancers and orchestration platforms know whether to send traffic to a service instance.

Two types: **Liveness check** (is the process alive? should it be restarted?) and **Readiness check** (is it ready to receive traffic? should the LB send requests to it?).

A service is live (don't kill it) but not ready (don't send traffic) during warm-up — while loading caches, establishing DB connections, or waiting for dependent services. Separating these prevents the 'restart loop from hell' where the process is killed and restarted repeatedly before it can warm up."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company using Kubernetes or microservices

#### Indepth
Kubernetes probe types:
```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 10  # Wait 10s before first check
  periodSeconds: 15        # Check every 15s
  failureThreshold: 3      # Kill after 3 consecutive failures

readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  periodSeconds: 5
  failureThreshold: 2      # Remove from LB after 2 failures
```

Health check endpoint implementation:
```go
// /healthz - liveness: "am I running?"
func healthz(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Write([]byte(`{"status":"ok"}`))
}

// /ready - readiness: "can I handle requests?"
func ready(w http.ResponseWriter, r *http.Request) {
    if !db.Ping() || !cache.Ping() {
        w.WriteHeader(503)
        w.Write([]byte(`{"status":"not ready", "error":"DB unavailable"}`))
        return
    }
    w.WriteHeader(200)
    w.Write([]byte(`{"status":"ready"}`))
}
```

Deep health check: Don't just return 200 — actually test critical dependencies (DB connection test, cache connectivity, downstream service reachability). A "healthy" service that can't reach its DB is not actually healthy.

---

### 7. What is graceful shutdown?

"Graceful shutdown is the process of **stopping a service cleanly** — completing in-flight requests, releasing resources, and deregistering from service discovery — before the process terminates.

Without graceful shutdown: Kubernetes sends SIGTERM to a pod during a rolling update. The OS immediately kills all goroutines/threads. 50 in-flight requests get `connection reset by peer`. Users see errors for no reason.

With graceful shutdown: On SIGTERM, (1) stop accepting new requests (close the listen socket), (2) wait for in-flight requests to complete (with a timeout), (3) close DB connections and Kafka producers cleanly, (4) exit. Now users experience the deployment seamlessly."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company doing continuous deployment

#### Indepth
Go graceful shutdown implementation:
```go
func main() {
    server := &http.Server{Addr: ":8080"}
    
    go server.ListenAndServe()
    
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
    <-quit  // Block until signal received
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    server.Shutdown(ctx)  // Stops accepting, waits for in-flight requests
    db.Close()            // Close DB connections
    producer.Flush(15000) // Flush Kafka producer
}
```

Kubernetes pre-stop hook: 
```yaml
lifecycle:
  preStop:
    exec:
      command: ["/bin/sleep", "5"]  # Wait 5s for LB to route traffic away before shutdown
```

The 5-second pre-stop sleep is important: Kubernetes removes the pod from the Service endpoints and tells the pod to shut down simultaneously. The iptables rules update takes ~2 seconds. During this gap, traffic can still reach the pod. The pre-stop sleep ensures no requests arrive after shutdown begins.

---

### 8. What is chaos engineering?

"Chaos engineering is the practice of **deliberately injecting failures into production systems** to discover weaknesses before they cause unplanned outages.

Pioneered by Netflix with 'Chaos Monkey' (2011) — a service that randomly terminates production EC2 instances during business hours. The reasoning: systems will fail eventually. If you routinely train for failure, your team and system are better equipped to handle real incidents.

The process: define the 'steady state' (normal behavior metrics), form a hypothesis ('the system will remain available even if Service X fails'), run the experiment (kill Service X), observe if steady state holds. If it doesn't — you've found a weakness in a controlled, intentional experiment rather than during a real incident."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Senior SRE roles, companies with mature reliability practices

#### Indepth
Chaos engineering tools:
- **Chaos Monkey (Netflix):** Terminates random EC2 instances in production
- **Chaos Gorilla:** Terminates entire AWS availability zones
- **Chaos Kong:** Terminates entire AWS regions (Netflix), forces traffic to work with one fewer region
- **Gremlin:** Commercial chaos engineering platform. Network latency injection, CPU spikes, disk fill, process kill.
- **Chaos Mesh / LitmusChaos:** Kubernetes-native chaos engineering. Inject pod failures, node failures, network partitions.
- **AWS Fault Injection Simulator:** Managed chaos engineering from AWS.

Chaos experiment design:
1. **Hypothesis:** "If the Redis cache fails, the order API falls back to the DB and serves requests within 500ms SLA"
2. **Experiment:** Use Gremlin to terminate all Redis pods
3. **Observe:** Monitor p99 latency, error rate, fallback activation
4. **Conclude:** Did the hypothesis hold? If yes → confidence. If no → fix the gap.

**Game days:** Scheduled chaos exercises where the team simulates a major outage scenario (region failure, DB corruption) and practices their incident response playbook. Amazon runs game days before Prime Day.

---

### 9. What is the timeout pattern?

"Timeouts define the **maximum duration to wait for a response** from an external operation. Without timeouts, a slow or unresponsive dependency can cause threads to block indefinitely, eventually exhausting the thread pool and taking down the calling service.

Setting the right timeout is challenging: too short causes false failures on legitimate slow operations; too long allows cascading failures to propagate. My approach: (1) measure the p99 latency of the target operation under normal conditions, (2) set timeout at 2-3x the p99 latency (captures occasional slowness, not cascade), (3) combine with a circuit breaker (stop trying after consistent timeout breaches).

At every level: **connect timeout** (maximum time to establish TCP connection — usually 1-2s) and **read timeout** (maximum time to receive the response after connection — depends on operation, 5-30s typically)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any distributed systems discussion

#### Indepth
Timeout cascade problem:
```
User Request → Service A (timeout: 30s) → Service B (timeout: 30s) → Service C (timeout: 30s)
If C hangs: B waits 30s, then A waits 30s = user waits up to 60s+ 
```

Prevent with **deadline propagation** (context in Go):
```go
// At API entry point, set an overall request deadline
ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
defer cancel()

// Pass ctx through all layers - each downstream call respects it
result, err := serviceB.Call(ctx, input)
```

If the overall deadline is 5 seconds, all downstream calls share that budget. Service B's call to Service C will fail with "context deadline exceeded" at second 5, not after 30 extra seconds. This requires consistent context propagation throughout the call chain.

HTTP timeout settings (Go):
```go
client := &http.Client{
    Timeout: 5 * time.Second, // Total request timeout
    Transport: &http.Transport{
        DialContext:           (&net.Dialer{Timeout: 1 * time.Second}).DialContext,
        TLSHandshakeTimeout:   1 * time.Second,
        ResponseHeaderTimeout: 3 * time.Second,
    },
}
```

---

### 10. What is a runbook and incident response architecture?

"A runbook is a **documented procedure** for humans to follow during a specific operational scenario — particularly incidents. It translates monitoring alerts into actionable steps: 'Alert: OrderService error rate > 5%. Step 1: Check recent deployments. Step 2: Check DB connection count. Step 3: Check upstream dependencies. Step 4: If DB connections are saturated, run [command] to kill idle connections.'

Incident response architecture: PagerDuty alerts the on-call engineer → they follow the runbook → resolve or escalate. The architecture must support this: comprehensive observability (can you diagnose quickly?), runbooks linked from alerts (zero friction to find the procedure), and automated remediation for known scenarios.

The goal is to reduce **MTTR** — from alert to resolution. A well-designed observability + runbook system brings MTTR from 2 hours to 15 minutes."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** SRE roles, team leads at any company with production systems

#### Indepth
Incident response lifecycle (SRE):
1. **Detect:** Alert fires (Prometheus, Datadog) based on SLO breach
2. **Triage:** On-call engineer assesses severity (P1/P2/P3)
3. **Respond:** Execute runbook, loop in SMEs if needed
4. **Resolve:** Mitigate root cause or apply workaround
5. **Review:** Post-incident review (blameless postmortem) within 24-48 hours
6. **Prevent:** Implement long-term fix + detection improvement

**Blameless postmortem template:**
- What happened? (Timeline of events)
- What was the impact? (Users affected, revenue impact, duration)
- What was the root cause? (Not "human error" — what system made human error easy?)
- What went well?
- What went poorly?
- Action items with owners and deadlines

The "5 Whys" for root cause analysis: Keep asking "why" until you find the systemic issue:
1. "Why was the service down?" → DB connections exhausted
2. "Why?" → Connection pool size was too small
3. "Why?" → No load testing was done before deployment
4. "Why?" → No load testing stage in CI/CD pipeline
5. **Root cause:** Missing load testing in the deployment pipeline
**Fix:** Add mandatory load testing to CI/CD, not just "increase connection pool"
