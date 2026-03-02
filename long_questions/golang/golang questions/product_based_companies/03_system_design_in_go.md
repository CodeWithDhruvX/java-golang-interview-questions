# 🏗️ 03 — System Design in Go
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Rate limiting (Token Bucket, Leaky Bucket)
- Distributed cache (Redis-like)
- Load balancing strategies
- Circuit breaker pattern
- Consistent hashing
- Distributed locking

---

## ❓ Most Asked Questions

### Q1. Implement Token Bucket Rate Limiter

```go
type TokenBucket struct {
    mu       sync.Mutex
    tokens   float64
    maxTokens float64
    refillRate float64  // tokens per second
    lastRefill time.Time
}

func NewTokenBucket(maxTokens, refillRate float64) *TokenBucket {
    return &TokenBucket{
        tokens:     maxTokens,
        maxTokens:  maxTokens,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }
}

func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()

    now := time.Now()
    elapsed := now.Sub(tb.lastRefill).Seconds()
    tb.tokens = math.Min(tb.maxTokens, tb.tokens+elapsed*tb.refillRate)
    tb.lastRefill = now

    if tb.tokens >= 1 {
        tb.tokens--
        return true  // request allowed
    }
    return false  // rate limited
}

// Usage with middleware
limiter := NewTokenBucket(10, 2)  // 10 tokens max, 2 tokens/sec
func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

---

### Q2. Implement an In-Memory LRU Cache

```go
import "container/list"

type LRUCache struct {
    mu       sync.Mutex
    capacity int
    cache    map[string]*list.Element
    order    *list.List
}

type entry struct{ key, value string }

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        cache:    make(map[string]*list.Element),
        order:    list.New(),
    }
}

func (c *LRUCache) Get(key string) (string, bool) {
    c.mu.Lock(); defer c.mu.Unlock()
    if elem, ok := c.cache[key]; ok {
        c.order.MoveToFront(elem)
        return elem.Value.(*entry).value, true
    }
    return "", false
}

func (c *LRUCache) Put(key, value string) {
    c.mu.Lock(); defer c.mu.Unlock()
    if elem, ok := c.cache[key]; ok {
        c.order.MoveToFront(elem)
        elem.Value.(*entry).value = value
        return
    }
    if c.order.Len() == c.capacity {
        // evict least recently used
        oldest := c.order.Back()
        c.order.Remove(oldest)
        delete(c.cache, oldest.Value.(*entry).key)
    }
    elem := c.order.PushFront(&entry{key, value})
    c.cache[key] = elem
}
```

---

### Q3. Implement Circuit Breaker Pattern

```go
type CircuitBreakerState int
const (
    StateClosed    CircuitBreakerState = iota  // normal: requests pass through
    StateOpen                                   // tripped: requests fail fast
    StateHalfOpen                              // trial: one request allowed
)

type CircuitBreaker struct {
    mu           sync.Mutex
    state        CircuitBreakerState
    failureCount int
    threshold    int
    timeout      time.Duration
    lastFailure  time.Time
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    switch cb.state {
    case StateOpen:
        if time.Since(cb.lastFailure) > cb.timeout {
            cb.state = StateHalfOpen
        } else {
            cb.mu.Unlock()
            return errors.New("circuit breaker open")
        }
    }
    cb.mu.Unlock()

    err := fn()

    cb.mu.Lock()
    defer cb.mu.Unlock()
    if err != nil {
        cb.failureCount++
        cb.lastFailure = time.Now()
        if cb.failureCount >= cb.threshold { cb.state = StateOpen }
        return err
    }
    cb.failureCount = 0
    cb.state = StateClosed
    return nil
}
```

---

### Q4. How do you implement distributed locking with Redis in Go?

```go
import "github.com/go-redis/redis/v8"

// Acquire distributed lock using SET NX PX
func acquireLock(rdb *redis.Client, key, value string, ttl time.Duration) (bool, error) {
    ctx := context.Background()
    result, err := rdb.SetNX(ctx, key, value, ttl).Result()
    return result, err
}

// Release lock (only if we own it)
var releaseScript = redis.NewScript(`
    if redis.call("get", KEYS[1]) == ARGV[1] then
        return redis.call("del", KEYS[1])
    else
        return 0
    end
`)

func releaseLock(rdb *redis.Client, key, value string) error {
    ctx := context.Background()
    return releaseScript.Run(ctx, rdb, []string{key}, value).Err()
}

// Usage
lockID := uuid.New().String()
acquired, err := acquireLock(rdb, "resource:lock", lockID, 30*time.Second)
if acquired {
    defer releaseLock(rdb, "resource:lock", lockID)
    // critical section
}
```

---

### Q5. How do you design a URL shortener in Go?

```go
// Key: base62 encoding of a counter or hash
const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func toBase62(num int64) string {
    if num == 0 { return "0" }
    var result []byte
    for num > 0 {
        result = append([]byte{base62Chars[num%62]}, result...)
        num /= 62
    }
    return string(result)
}

type URLShortener struct {
    mu      sync.Mutex
    store   map[string]string  // short → long
    counter int64
}

func (s *URLShortener) Shorten(longURL string) string {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.counter++
    short := toBase62(s.counter)
    s.store[short] = longURL
    return "https://short.ly/" + short
}

func (s *URLShortener) Resolve(short string) (string, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    long, ok := s.store[short]
    return long, ok
}
```

---

### Q6. Consistent Hashing for distributed systems

```go
import (
    "crypto/md5"
    "fmt"
    "sort"
)

type ConsistentHash struct {
    replicas int
    ring     map[uint32]string
    keys     []uint32  // sorted
}

func NewConsistentHash(replicas int) *ConsistentHash {
    return &ConsistentHash{replicas: replicas, ring: make(map[uint32]string)}
}

func hash(key string) uint32 {
    h := md5.Sum([]byte(key))
    return uint32(h[0])<<24 | uint32(h[1])<<16 | uint32(h[2])<<8 | uint32(h[3])
}

func (c *ConsistentHash) AddNode(node string) {
    for i := 0; i < c.replicas; i++ {
        key := hash(fmt.Sprintf("%s:%d", node, i))
        c.ring[key] = node
        c.keys = append(c.keys, key)
    }
    sort.Slice(c.keys, func(i, j int) bool { return c.keys[i] < c.keys[j] })
}

func (c *ConsistentHash) GetNode(key string) string {
    h := hash(key)
    idx := sort.Search(len(c.keys), func(i int) bool { return c.keys[i] >= h })
    if idx == len(c.keys) { idx = 0 }
    return c.ring[c.keys[idx]]
}
```
