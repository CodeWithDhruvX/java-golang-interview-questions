## 🔥 Performance & Optimization (Questions 201-220) - CONTINUATION

### Question 208: What is inlining and how does the Go compiler handle it?

**Answer:**
Inlining replaces function calls with the function body for performance:

**Automatic inlining:**
```go
// Small functions are automatically inlined:
func add(a, b int) int {
    return a + b  // Will be inlined
}

// Usage:
result := add(2, 3)
// Becomes: result := 2 + 3
```

**Control inlining:**
```go
// Prevent inlining:
//go:noinline
func expensive() {
    // Complex logic
}

// Check if function is inlined:
// go build -gcflags="-m" main.go
```

**Benefits:**
- Eliminates function call overhead
- Enables further optimizations
- Reduces stack operations

**Limitations:**
- Only small functions (< 80 nodes)
- Functions with loops, defer, recover not inlined
- Recursive functions not inlined

---

### Question 209: How do you debug GC pauses?

**Answer:**
Monitor and tune garbage collection:

**Enable GC logging:**
```bash
GODEBUG=gctrace=1 ./myapp
```

**Output:**
```
gc 1 @0.004s 0%: 0.018+1.2+0.003 ms clock, 0.14+0/0.59/1.3+0.030 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
```

**Analyze GC stats:**
```go
import "runtime"

func printGCStats() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Alloc = %v MB", m.Alloc/1024/1024)
    fmt.Printf("TotalAlloc = %v MB", m.TotalAlloc/1024/1024)
    fmt.Printf("Sys = %v MB", m.Sys/1024/1024)
    fmt.Printf("NumGC = %v\n", m.NumGC)
    fmt.Printf("PauseTotal = %v ms\n", m.PauseTotalNs/1000000)
}
```

**Tune GC:**
```bash
# Set GC percentage (default 100):
GOGC=200 ./myapp  # Less frequent GC

# Disable GC (not recommended):
GOGC=off ./myapp
```

**Reduce GC pressure:**
```go
// 1. Reuse objects with sync.Pool
// 2. Pre-allocate slices
// 3. Avoid unnecessary pointers
// 4. Use value types when possible
```

---

### Question 210: What are some common performance bottlenecks in Go apps?

**Answer:**

**1. Excessive allocations:**
```go
// BAD:
for {
    item := &Item{}  // Allocation per iteration
    process(item)
}

// GOOD:
pool := &sync.Pool{New: func() interface{} { return &Item{} }}
for {
    item := pool.Get().(*Item)
    process(item)
    pool.Put(item)
}
```

**2. Using `+` for string concatenation:**
```go
// BAD:
s := ""
for i := 0; i < 1000; i++ {
    s += "item"  // Creates 1000 strings
}

// GOOD:
var b strings.Builder
for i := 0; i < 1000; i++ {
    b.WriteString("item")
}
s := b.String()
```

**3. Not buffering I/O:**
```go
// BAD:
for _, line := range lines {
    file.WriteString(line)  // Unbuffered
}

// GOOD:
writer := bufio.NewWriter(file)
for _, line := range lines {
    writer.WriteString(line)
}
writer.Flush()
```

**4. Goroutine leaks:**
```go
// BAD:
for {
    go processRequest()  // Never-ending goroutines
}

// GOOD:
workerPool := make(chan struct{}, 100)
for {
    workerPool <- struct{}{}
    go func() {
        processRequest()
        <-workerPool
    }()
}
```

**5. Lock contention:**
```go
// BAD - Single mutex for everything:
var mu sync.Mutex
var data map[string]int

// GOOD - Sharded locks:
type ShardedMap struct {
    shards [256]struct {
        sync.RWMutex
        data map[string]int
    }
}
```

---

### Question 211: How to detect and fix memory leaks?

**Answer:**

**Detect leaks:**
```bash
# Memory profile over time:
curl http://localhost:6060/debug/pprof/heap > heap1.prof
# ... wait ...
curl http://localhost:6060/debug/pprof/heap > heap2.prof

# Compare:
go tool pprof -base heap1.prof heap2.prof
```

**Common causes:**

**1. Goroutine leaks:**
```go
// BAD:
func leak() {
    ch := make(chan int)
    go func() {
        <-ch  // Blocks forever
    }()
}

// FIX:
func noLeak() {
    ch := make(chan int)
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    go func() {
        select {
        case <-ch:
        case <-ctx.Done():
            return
        }
    }()
}
```

**2. Global variable accumulation:**
```go
// BAD:
var cache = make(map[string][]byte)

func store(key string, data []byte) {
    cache[key] = data  // Never cleared
}

// FIX - Use LRU cache or TTL:
type Cache struct {
    items sync.Map
   ttl   time.Duration
}

func (c *Cache) Set(key string, value []byte) {
    c.items.Store(key, &item{
        value:     value,
        expiresAt: time.Now().Add(c.ttl),
    })
}
```

**3. Unclosed resources:**
```go
// BAD:
func readFile() {
    f, _ := os.Open("file.txt")
    // Never closed
}

// FIX:
func readFile() {
    f, _ := os.Open("file.txt")
    defer f.Close()
}
```

---

### Question 212: How do you find goroutine leaks?

**Answer:**

**Check goroutine count:**
```go
import "runtime"

func printGoroutines() {
    fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
}
```

**Get goroutine dump:**
```bash
# From pprof endpoint:
curl http://localhost:6060/debug/pprof/goroutine?debug=2

# Or send SIGQUIT:
kill -QUIT <pid>
```

**Use goleak for tests:**
```go
import "go.uber.org/goleak"

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}

func TestFunction(t *testing.T) {
    defer goleak.VerifyNone(t)
    
    // Test code that might leak
}
```

**Common leak patterns:**
```go
// 1. Channel without receiver:
ch := make(chan int)
go func() {
    ch <- 1  // Blocks forever
}()

// 2. Context not canceled:
ctx := context.Background()
go worker(ctx)  // Never stops

// 3. Infinite loop without exit:
go func() {
    for {
        time.Sleep(time.Second)
        // No way to stop
    }
}()
```

**Fixes:**
```go
// Always provide a way to stop goroutines:
done := make(chan struct{})
go func() {
    for {
        select {
        case <-done:
            return
        default:
            work()
        }
    }
}()

// Later:
close(done)
```

---

### Question 213: How do you tune GC parameters in production?

**Answer:**

**GOGC environment variable:**
```bash
# Default is 100 (GC when heap doubles):
GOGC=100 ./app

# Less frequent GC (more memory):
GOGC=200 ./app

# More frequent GC (less memory):
GOGC=50 ./app
```

**Set programmatically:**
```go
import "runtime/debug"

func init() {
    // Set GC percentage:
    debug.SetGCPercent(200)
    
    // Set memory limit (Go 1.19+):
    debug.SetMemoryLimit(8 * 1024 * 1024 * 1024) // 8GB
}
```

**Monitor GC:**
```go
var m runtime.MemStats

ticker := time.NewTicker(10 * time.Second)
go func() {
    for range ticker.C {
        runtime.ReadMemStats(&m)
        log.Printf("GC runs: %d, Pause: %v ms", 
            m.NumGC, 
            m.PauseTotalNs/1000000)
    }
}()
```

**Best practices:**
- Start with default (GOGC=100)
- Monitor pause times and memory
- Increase GOGC if memory is available
- Decrease if memory is constrained

---

### Question 214: How to avoid blocking operations in hot paths?

**Answer:**

**1. Use buffered channels:**
```go
// BAD - May block:
ch := make(chan int)
ch <- value

// GOOD - Won't block if buffer not full:
ch := make(chan int, 100)
ch <- value
```

**2. Non-blocking send:**
```go
select {
case ch <- value:
    // Sent successfully
default:
    // Would block, handle differently
    log.Println("Channel full, dropping")
}
```

**3. Use sync.Pool instead of new allocations:**
```go
var pool = sync.Pool{
    New: func() interface{} {
        return &Object{}
    },
}

// In hot path:
obj := pool.Get().(*Object)
defer pool.Put(obj)
```

**4. Avoid mutex in read-heavy scenarios:**
```go
// BAD:
var (
    mu   sync.Mutex
    data map[string]int
)

func get(key string) int {
    mu.Lock()
    defer mu.Unlock()
    return data[key]
}

// GOOD - Use RWMutex:
var (
    mu   sync.RWMutex
    data map[string]int
)

func get(key string) int {
    mu.RLock()
    defer mu.RUnlock()
    return data[key]
}

// BETTER - Use sync.Map for concurrent access:
var data sync.Map

func get(key string) int {
    val, _ := data.Load(key)
    return val.(int)
}
```

---

### Question 215: What are the trade-offs of pooling in Go?

**Answer:**

**Pros:**
- Reduces allocations
- Better GC performance
- Faster for frequently created objects

**Cons:**
- Memory not freed immediately
- Need to reset pooled objects
- Complex to use correctly

**sync.Pool example:**
```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func process(data []byte) string {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    
    buf.Reset()  // IMPORTANT - Clear old data
    buf.Write(data)
    return buf.String()
}
```

**When to use:**
✅ High-frequency allocations
✅ Uniform object sizes
✅ Short-lived objects in hot paths

**When NOT to use:**
❌ Long-lived objects
❌ Objects with complex state
❌ Rare allocations

**Common mistakes:**
```go
// BAD - Pool pointer to slice:
pool := sync.Pool{
    New: func() interface{} {
        s := make([]int, 0, 100)
        return &s  // Pointer to local variable
    },
}

// GOOD:
pool := sync.Pool{
    New: func() interface{} {
        return &[]int{}  // Pointer to new slice
    },
}

// Or better:
pool := sync.Pool{
    New: func() interface{} {
        s := make([]int, 0, 100)
        return &s
    },
}
```

---

### Question 216: How do you measure latency and throughput in Go APIs?

**Answer:**

**Middleware for timing:**
```go
func timingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Wrap response writer to capture status:
        wrapped := &responseWriter{ResponseWriter: w}
        
        next.ServeHTTP(wrapped, r)
        
        duration := time.Since(start)
        
        log.Printf("%s %s - %d - %v",
            r.Method,
            r.URL.Path,
            wrapped.status,
            duration,
        )
    })
}

type responseWriter struct {
    http.ResponseWriter
    status int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.status = code
    rw.ResponseWriter.WriteHeader(code)
}
```

**Prometheus metrics:**
```go
import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    httpDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "HTTP request latency",
        },
        []string{"method", "path", "status"},
    )
    
    httpRequests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total HTTP requests",
        },
        []string{"method", "path", "status"},
    )
)

func init() {
    prometheus.MustRegister(httpDuration)
    prometheus.MustRegister(httpRequests)
}

func metricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        wrapped := &responseWriter{ResponseWriter: w, status: 200}
        
        next.ServeHTTP(wrapped, r)
        
        duration := time.Since(start).Seconds()
        status := fmt.Sprintf("%d", wrapped.status)
        
        httpDuration.WithLabelValues(r.Method, r.URL.Path, status).Observe(duration)
        httpRequests.WithLabelValues(r.Method, r.URL.Path, status).Inc()
    })
}

// Expose metrics:
http.Handle("/metrics", promhttp.Handler())
```

**Throughput testing:**
```go
func BenchmarkAPIEndpoint(b *testing.B) {
    handler := setupHandler()
    req := httptest.NewRequest("GET", "/api/users", nil)
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            w := httptest.NewRecorder()
            handler.ServeHTTP(w, req)
        }
    })
    
    b.ReportMetric(float64(b.N)/b.Elapsed().Seconds(), "req/s")
}
```

---

### Question 217: What is backpressure and how do you handle it?

**Answer:**
Backpressure prevents system overload when consumers can't keep up with producers:

**Problem:**
```go
// Producer overwhelms consumer:
for data := range fastProducer() {
    slowConsumer(data)  // Queue keeps growing
}
```

**Solution 1 - Buffered channels with limits:**
```go
ch := make(chan Request, 100)  // Limited buffer

// Producer with backpressure:
select {
case ch <- request:
    // Sent successfully
case <-time.After(100 * time.Millisecond):
    // Timeout - reject request
    http.Error(w, "System overloaded", 503)
}
```

**Solution 2 - Worker pool:**
```go
type WorkerPool struct {
    workers   int
    taskQueue chan Task
}

func NewWorkerPool(workers int) *WorkerPool {
    wp := &WorkerPool{
        workers:   workers,
        taskQueue: make(chan Task, workers*2),  // Bounded queue
    }
    
    for i := 0; i < workers; i++ {
        go wp.worker()
    }
    
    return wp
}

func (wp *WorkerPool) Submit(task Task) error {
    select {
    case wp.taskQueue <- task:
        return nil
    default:
        return errors.New("queue full")  // Backpressure signal
    }
}

func (wp *WorkerPool) worker() {
    for task := range wp.taskQueue {
        task.Execute()
    }
}
```

**Solution 3 - Rate limiting:**
```go
import "golang.org/x/time/rate"

limiter := rate.NewLimiter(100, 200)  // 100 req/s, burst 200

func handler(w http.ResponseWriter, r *http.Request) {
    if !limiter.Allow() {
        http.Error(w, "Too many requests", 429)
        return
    }
    
    // Process request
}
```

---

### Question 218: When should you prefer sync.Pool?

**Answer:**

**Use sync.Pool when:**

✅ **High-frequency allocations:**
```go
// Processing millions of requests:
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func handleRequest(data []byte) {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    
    buf.Reset()
    buf.Write(data)
    // Process...
}
```

✅ **Uniform object size:**
```go
// All objects are similar:
type Parser struct {
    buffer [4096]byte
}

var parserPool = sync.Pool{
    New: func() interface{} {
        return &Parser{}
    },
}
```

✅ **Short-lived objects:**
```go
// Object used briefly then returned:
obj := pool.Get()
defer pool.Put(obj)
// Brief processing
```

**DON'T use sync.Pool when:**

❌ **Long-lived objects:**
```go
// BAD - Object kept for long time:
conn := pool.Get().(*Connection)
// Keep connection for hours
```

❌ **Need guaranteed cleanup:**
```go
// BAD - Pool may be cleared by GC:
// Don't rely on destructors
```

❌ **Complex state to reset:**
```go
// BAD - Expensive reset operation:
obj := pool.Get().(*ComplexObject)
obj.ResetAllFields()  // If this is expensive, pool doesn't help
```

**Example - HTTP client pooling:**
```go
var clientPool = sync.Pool{
    New: func() interface{} {
        return &http.Client{
            Timeout: 10 * time.Second,
        }
    },
}

func makeRequest(url string) (*http.Response, error) {
    client := clientPool.Get().(*http.Client)
    defer clientPool.Put(client)
    
    return client.Get(url)
}
```

---

### Question 219: How do you manage high concurrency with low resource usage?

**Answer:**

**1. Worker pool pattern:**
```go
type Job struct {
    ID   int
    Data interface{}
}

func workerPool(numWorkers int, jobs <-chan Job) {
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                processJob(job)
            }
        }()
    }
    
    wg.Wait()
}

// Usage:
jobs := make(chan Job, 1000)
go workerPool(10, jobs)  // 10 workers, not 1000 goroutines

for i := 0; i < 1000; i++ {
    jobs <- Job{ID: i}
}
close(jobs)
```

**2. Semaphore pattern:**
```go
sem := make(chan struct{}, 100)  // Max 100 concurrent

for _, item := range items {
    sem <- struct{}{}  // Acquire
    
    go func(item Item) {
        defer func() { <-sem }()  // Release
        process(item)
    }(item)
}

// Wait for all:
for i := 0; i < cap(sem); i++ {
    sem <- struct{}{}
}
```

**3. Connection pooling:**
```go
type ConnectionPool struct {
    conns chan *Connection
}

func NewPool(size int) *ConnectionPool {
    p := &ConnectionPool{
        conns: make(chan *Connection, size),
    }
    
    for i := 0; i < size; i++ {
        p.conns <- newConnection()
    }
    
    return p
}

func (p *ConnectionPool) Get() *Connection {
    return <-p.conns
}

func (p *ConnectionPool) Put(conn *Connection) {
    p.conns <- conn
}
```

**4. Batch processing:**
```go
func batchProcessor(items <-chan Item, batchSize int) {
    batch := make([]Item, 0, batchSize)
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()
    
    for {
        select {
        case item := <-items:
            batch = append(batch, item)
            
            if len(batch) >= batchSize {
                processBatch(batch)
                batch = batch[:0]
            }
            
        case <-ticker.C:
            if len(batch) > 0 {
                processBatch(batch)
                batch = batch[:0]
            }
        }
    }
}
```

---

### Question 220: How do you monitor a Go application in production?

**Answer:**

**1. Runtime metrics:**
```go
import (
    "runtime"
    "time"
)

func monitorRuntime() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        
        log.Printf("Goroutines: %d", runtime.NumGoroutine())
        log.Printf("Memory: Alloc=%v MB, Sys=%v MB", 
            m.Alloc/1024/1024, 
            m.Sys/1024/1024)
        log.Printf("GC: NumGC=%v, PauseTotal=%v ms",
            m.NumGC,
            m.PauseTotalNs/1000000)
    }
}
```

**2. Prometheus integration:**
```go
import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
        Name: "myapp_processed_ops_total",
        Help: "Total processed operations",
    })
    
    requestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
        Name: "myapp_request_duration_seconds",
        Help: "Request duration",
    })
)

func recordMetrics() {
    go func() {
        for {
            opsProcessed.Inc()
            time.Sleep(2 * time.Second)
        }
    }()
}
```

**3. Health check endpoint:**
```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    health := struct {
        Status     string `json:"status"`
        Goroutines int    `json:"goroutines"`
        Memory     uint64 `json:"memory_mb"`
    }{
        Status:     "healthy",
        Goroutines: runtime.NumGoroutine(),
    }
    
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    health.Memory = m.Alloc / 1024 / 1024
    
    json.NewEncoder(w).Encode(health)
}
```

**4. Structured logging:**
```go
import "github.com/sirupsen/logrus"

log := logrus.New()
log.SetFormatter(&logrus.JSONFormatter{})

log.WithFields(logrus.Fields{
    "user_id": userId,
    "action":  "purchase",
    "amount":  amount,
}).Info("Purchase completed")
```

---

*[Questions 221-280 will follow covering Files/OS, Microservices, and Security]*
