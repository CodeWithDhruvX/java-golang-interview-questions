# Golang Functional Programming — Questions & Answers

> **Topics:** Real-world scenarios using functions, goroutines, channels, slices, and functional patterns in business applications

---

## Section 1: Data Processing & Analytics (Q1–Q8)

### 1. Sales Analytics — Calculate Revenue Metrics
**Q: Analyze sales data to calculate key metrics using Go slices and functions. What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

type Sale struct {
    Product string
    Amount  float64
    Region  string
    Date    time.Time
}

func main() {
    sales := []Sale{
        {"Laptop", 1200.0, "North", time.Now().AddDate(0, 0, -10)},
        {"Mouse", 25.0, "North", time.Now().AddDate(0, 0, -8)},
        {"Keyboard", 75.0, "South", time.Now().AddDate(0, 0, -5)},
        {"Laptop", 999.0, "South", time.Now().AddDate(0, 0, -3)},
        {"Monitor", 300.0, "North", time.Now().AddDate(0, 0, -1)},
    }

    type Stats struct {
        Total float64
        Sum   float64
        Count int
    }
    regionStats := make(map[string]*Stats)

    for _, s := range sales {
        if _, ok := regionStats[s.Region]; !ok {
            regionStats[s.Region] = &Stats{}
        }
        regionStats[s.Region].Sum += s.Amount
        regionStats[s.Region].Count++
    }

    fmt.Println("Revenue by Region:")
    for region, stats := range regionStats {
        fmt.Printf("%s: $%.2f (avg: $%.2f, orders: %d)\n",
            region, stats.Sum, stats.Sum/float64(stats.Count), stats.Count)
    }
}
```
**A:**
```
Revenue by Region:
North: $1525.00 (avg: $508.33, orders: 3)
South: $1074.00 (avg: $537.00, orders: 2)
```

---

### 2. Customer Segmentation — Behavioral Analysis
**Q: Segment customers based on purchase behavior using Go functions and maps. What is the output?**
```go
package main
import "fmt"

type Customer struct {
    ID string; Name string; Orders int; TotalSpent float64
}

func main() {
    customers := []Customer{
        {"C001", "Alice", 15, 2500.0},
        {"C002", "Bob", 3, 180.0},
        {"C003", "Charlie", 8, 1200.0},
        {"C004", "Diana", 25, 4500.0},
        {"C005", "Eve", 1, 50.0},
    }

    isVIP := func(c Customer) bool { return c.Orders >= 10 && c.TotalSpent >= 1000 }
    isRegular := func(c Customer) bool { return c.Orders >= 3 && c.TotalSpent >= 100 }

    segments := make(map[string][]Customer)
    for _, c := range customers {
        label := "New"
        if isVIP(c) { label = "VIP" } else if isRegular(c) { label = "Regular" }
        segments[label] = append(segments[label], c)
    }

    for label, list := range segments {
        fmt.Printf("%s Customers (%d):\n", label, len(list))
        for _, c := range list {
            fmt.Printf("  %s: %d orders, $%.2f\n", c.Name, c.Orders, c.TotalSpent)
        }
    }
}
```
**A:**
```
VIP Customers (2):
  Alice: 15 orders, $2500.00
  Diana: 25 orders, $4500.00
Regular Customers (1):
  Charlie: 8 orders, $1200.00
New Customers (2):
  Bob: 3 orders, $180.00
  Eve: 1 orders, $50.00
```

---

### 3. Product Performance — Top Sellers Analysis
**Q: Find top-performing products using Go sort and slice operations. What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

type ProductSale struct {
    Product string; Quantity int; Revenue float64
}

func main() {
    sales := []ProductSale{
        {"Laptop", 45, 45000.0},
        {"Mouse", 120, 3000.0},
        {"Keyboard", 85, 6375.0},
        {"Monitor", 60, 18000.0},
        {"Webcam", 200, 8000.0},
    }

    sort.Slice(sales, func(i, j int) bool { return sales[i].Revenue > sales[j].Revenue })
    fmt.Println("Top 3 by Revenue:")
    for i := 0; i < 3; i++ {
        fmt.Printf("  %s: $%.2f (%d units)\n", sales[i].Product, sales[i].Revenue, sales[i].Quantity)
    }

    sort.Slice(sales, func(i, j int) bool { return sales[i].Quantity > sales[j].Quantity })
    fmt.Println("\nTop 3 by Quantity:")
    for i := 0; i < 3; i++ {
        fmt.Printf("  %s: %d units ($%.2f)\n", sales[i].Product, sales[i].Quantity, sales[i].Revenue)
    }
}
```
**A:**
```
Top 3 by Revenue:
  Laptop: $45000.00 (45 units)
  Monitor: $18000.00 (60 units)
  Webcam: $8000.00 (200 units)

Top 3 by Quantity:
  Webcam: 200 units ($8000.00)
  Mouse: 120 units ($3000.00)
  Keyboard: 85 units ($6375.00)
```

---

### 4. Time Series Analysis — Trend Detection
**Q: Analyze sales trends over time using Go time package and slices. What is the output?**
```go
package main
import "fmt"

type DailySale struct {
    Day int; Amount float64
}

func main() {
    sales := []DailySale{{1, 1000}, {2, 1200}, {3, 1150}, {4, 1400}, {5, 1600}, {6, 1550}, {7, 1800}}
    
    var changes []float64
    for i := 1; i < len(sales); i++ {
        changes = append(changes, sales[i].Amount - sales[i-1].Amount)
    }

    var totalChange float64
    upDays := 0
    for _, c := range changes {
        totalChange += c
        if c > 0 { upDays++ }
    }

    fmt.Printf("Average daily change: $%.2f\n", totalChange/float64(len(changes)))
    fmt.Printf("Up days: %d/%d (%.1f%%)\n", upDays, len(changes), float64(upDays)/float64(len(changes))*100)
}
```
**A:**
```
Average daily change: $133.33
Up days: 5/6 (83.3%)
```

---

### 5. Customer Lifetime Value — Predictive Analytics
**Q: Calculate customer lifetime value using Go functions and structs. What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

type Customer struct {
    ID string; AvgValue float64; Freq int; Months int
}

func main() {
    customers := []Customer{{"C001", 150, 2, 12}, {"C002", 75, 1, 6}, {"C003", 200, 3, 18}, {"C004", 50, 1, 3}}
    
    clvCalc := func(c Customer) float64 {
        return (c.AvgValue * float64(c.Freq)) * (float64(c.Months) * 1.5)
    }

    type Result struct { ID string; CLV float64 }
    var results []Result
    for _, c := range customers { results = append(results, Result{c.ID, clvCalc(c)}) }

    sort.Slice(results, func(i, j int) bool { return results[i].CLV > results[j].CLV })
    fmt.Println("Customer Lifetime Values:")
    for _, r := range results { fmt.Printf("%s: $%.2f\n", r.ID, r.CLV) }
}
```
**A:**
```
Customer Lifetime Values:
C003: $16200.00
C001: $5400.00
C002: $675.00
C004: $225.00
```

---

### 6. Market Basket Analysis — Product Associations
**Q: Find products frequently bought together using Go maps and channels. What is the output?**
```go
package main
import "fmt"

func main() {
    transactions := [][]string{
        {"Laptop", "Mouse", "Keyboard"},
        {"Laptop", "Monitor"},
        {"Mouse", "Keyboard", "Webcam"},
        {"Laptop", "Mouse"},
        {"Keyboard", "Monitor"},
    }

    target := "Laptop"
    counts := make(map[string]int)

    for _, tx := range transactions {
        containsTarget := false
        for _, p := range tx { if p == target { containsTarget = true; break } }
        if containsTarget {
            for _, p := range tx { if p != target { counts[p]++ } }
        }
    }

    fmt.Printf("Products frequently bought with %s:\n", target)
    for p, c := range counts { fmt.Printf("  %s: %d times\n", p, c) }
}
```
**A:**
```
Products frequently bought with Laptop:
  Mouse: 2 times
  Monitor: 1 times
  Keyboard: 1 times
```

---

### 7. Sales Forecasting — Trend Prediction
**Q: Predict future sales using historical data and Go functions. What is the output?**
```go
package main
import "fmt"

func main() {
    history := []float64{10000, 10500, 11200, 11800, 12500, 13200}
    var growth []float64
    for i := 1; i < len(history); i++ {
        growth = append(growth, (history[i]-history[i-1])/history[i-1])
    }

    sumGrowth := 0.0
    for _, g := range growth { sumGrowth += g }
    avgGrowth := sumGrowth / float64(len(growth))

    prediction := history[len(history)-1] * (1 + avgGrowth)
    fmt.Printf("Average growth: %.2f%%\n", avgGrowth*100)
    fmt.Printf("Next month prediction: $%.2f\n", prediction)
}
```
**A:**
```
Average growth: 5.67%
Next month prediction: $13948.44
```

---

### 8. Customer Churn Analysis — Risk Assessment
**Q: Identify customers at risk of churning using Go functions and filters. What is the output?**
```go
package main
import "fmt"

type Profile struct { ID string; DaysSince int; LastMonthOrders int }

func main() {
    ps := []Profile{{"C001", 45, 2}, {"C002", 5, 8}, {"C003", 90, 0}, {"C004", 30, 1}, {"C005", 15, 4}}
    
    highRisk := func(p Profile) bool { return p.DaysSince > 60 || p.LastMonthOrders == 0 }
    
    fmt.Println("High Risk Customers:")
    for _, p := range ps {
        if highRisk(p) {
            fmt.Printf("  %s: %d days since last purchase\n", p.ID, p.DaysSince)
        }
    }
}
```
**A:**
```
High Risk Customers:
  C003: 90 days since last purchase
```

---

## Section 2: Business Logic & Decision Making (Q9–Q16)

### 9. Discount Engine — Complex Pricing Rules
**Q: Apply complex discount rules using Go function composition. What is the output?**
```go
package main
import "fmt"

type Order struct { ID string; Amount float64; Tier string; Count int }

func main() {
    orders := []Order{
        {"C001", 500, "VIP", 15},
        {"C005", 1200, "VIP", 8},
    }

    for _, o := range orders {
        discount := 0.0
        if o.Tier == "VIP" { discount += o.Amount * 0.15 }
        if o.Count > 10 { discount += o.Amount * 0.10 }
        if o.Amount > 1000 { discount += o.Amount * 0.05 }

        fmt.Printf("%s: $%.2f -> $%.2f (Disc: $%.2f)\n", o.ID, o.Amount, o.Amount-discount, discount)
    }
}
```
**A:**
```
C001: $500.00 -> $375.00 (Disc: $125.00)
C005: $1200.00 -> $960.00 (Disc: $240.00)
```

---

### 10. Lead Scoring — Qualification System
**Q: Score and qualify leads using Go functions and predicates. What is the output?**
```go
package main
import "fmt"

type Lead struct { Name string; Size int; Budget float64 }

func main() {
    leads := []Lead{{"TechCorp", 500, 50000}, {"SmallBiz", 20, 5000}}
    
    score := func(l Lead) int {
        s := 0
        if l.Size >= 200 { s += 30 }
        if l.Budget >= 20000 { s += 40 }
        return s
    }

    for _, l := range leads {
        fmt.Printf("%s: Score %d\n", l.Name, score(l))
    }
}
```
**A:**
```
TechCorp: Score 70
SmallBiz: Score 0
```

---

### 11. Inventory Optimization — Stock Level Decisions
**Q: Determine optimal stock levels using Go functions and analysis. What is the output?**
```go
package main
import "fmt"

type Product struct { Name string; Stock int; Demand int; LeadTime int }

func main() {
    p := Product{"Monitor", 15, 15, 10}
    reorderPoint := (p.Demand/30)*p.LeadTime + (p.Demand/30)*p.LeadTime // demand + safety
    
    status := "OK"
    if p.Stock <= reorderPoint { status = "REORDER" }
    
    fmt.Printf("%s: Stock %d, Reorder at %d -> %s\n", p.Name, p.Stock, reorderPoint, status)
}
```
**A:**
```
Monitor: Stock 15, Reorder at 10 -> OK
```

---

### 12. Risk Assessment — Credit Scoring
**Q: Assess credit risk using Go function composition. What is the output?**
```go
package main
import "fmt"

type App struct { Score int; DTI float64; Years int }

func main() {
    a := App{750, 0.25, 5}
    calc := func(a App) int {
        s := 0
        if a.Score >= 750 { s += 40 }
        if a.DTI <= 0.30 { s += 20 }
        if a.Years >= 5 { s += 30 }
        return s
    }
    fmt.Printf("Credit Score: %d\n", calc(a))
}
```
**A:**
```
Credit Score: 90
```

---

### 13. Pricing Strategy — Dynamic Pricing
**Q: Implement dynamic pricing based on multiple factors using Go. What is the output?**
```go
package main
import "fmt"

type PricePlan struct { Base float64; Demand int; Stock int }

func main() {
    p := PricePlan{100, 9, 10}
    adj := 1.0
    if p.Demand >= 8 { adj *= 1.2 }
    if p.Stock < 20 { adj *= 1.15 }
    
    fmt.Printf("Final Price: $%.2f\n", p.Base * adj)
}
```
**A:**
```
Final Price: $138.00
```

---

### 14. Customer Support — Priority Queue
**Q: Prioritize support tickets using Go interfaces and sorting. What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

type Ticket struct { ID string; Priority int }

func main() {
    ts := []Ticket{{"T1", 3}, {"T2", 1}, {"T3", 2}}
    sort.Slice(ts, func(i, j int) bool { return ts[i].Priority < ts[j].Priority })
    fmt.Println("Queue:", ts)
}
```
**A:**
```
Queue: [{T2 1} {T3 2} {T1 3}]
```

---

### 15. Quality Control — Defect Detection
**Q: Identify quality issues using Go functions and predicates. What is the output?**
```go
package main
import "fmt"

type Tool struct { ID string; Defects int; WeightVar float64 }

func main() {
    tools := []Tool{{"P1", 0, 0.02}, {"P2", 3, 0.15}}
    fail := func(t Tool) bool { return t.Defects > 0 || t.WeightVar > 0.1 }
    
    for _, t := range tools {
        res := "PASS"
        if fail(t) { res = "FAIL" }
        fmt.Printf("%s: %s\n", t.ID, res)
    }
}
```
**A:**
```
P1: PASS
P2: FAIL
```

---

### 16. Marketing Campaign — Target Audience Selection
**Q: Select target audience for marketing campaigns using Go. What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

type User struct { ID string; Interests string }

func main() {
    users := []User{{"U1", "tech,gaming"}, {"U2", "finance"}}
    target := "tech"
    for _, u := range users {
        if strings.Contains(u.Interests, target) {
            fmt.Println("Targeting:", u.ID)
        }
    }
}
```
**A:**
```
Targeting: U1
```

---

## Section 3: File Processing & Data Transformation (Q17–Q24)

### 17. Data Validation — Schema Enforcement
**Q: Validate data records against schema using Go patterns. What is the output?**
```go
package main
import "fmt"

func main() {
    emails := []string{"a@b.com", "invalid"}
    validate := func(e string) bool { return len(e) > 5 && strings.Contains(e, "@") }
    // Loop checks...
}
```
**A:** (Logic matches Java version: validates length/format)

---

### 18. Data Transformation — ETL Pipeline
**Q: Transform data through ETL pipeline using Go operations. What is the output?**
```go
// ETL Logic: Map raw string salary to float64, normalize names...
```
**A:** (Go implementation uses `strconv.ParseFloat` and `strings.Title`)

---

### 19. Log Analysis — Error Pattern Detection
**Q: Analyze log files to detect error patterns using Go. What is the output?**
```go
// Filter logs where level == "ERROR", count occurrences...
```
**A:** (Result matches Java: finds frequency of message types)

---

### 20. Data Aggregation — Multi-level Grouping
**Q: Aggregate data across multiple dimensions using Go. What is the output?**
```go
// Iterative map-in-map approach for [Region][Category]Stats
```

---

### 21. Data Cleaning — Remove Duplicates and Normalize
**Q: Clean and normalize dataset using Go operations. What is the output?**
```go
// Use map[Key]bool to track uniqueness after normalization
```

---

### 22. Data Enrichment — Add Computed Fields
**Q: Enrich data with computed fields using Go functions. What is the output?**

---

### 23. Data Filtering — Complex Query Logic
**Q: Apply complex filtering logic to dataset using Go. What is the output?**

---

### 24. Data Export — Format Transformation
**Q: Transform data for export in different formats using Go. What is the output?**

---

## Section 4: Performance & Optimization (Q25–Q32)

### 25. Parallel Processing — Large Dataset Analysis
**Q: Compare sequential vs parallel processing using goroutines. What is the output?**
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

func main() {
    data := make([]int, 1000000)
    start := time.Now()
    sum := 0
    for _, v := range data { sum += v }
    fmt.Println("Seq Time:", time.Since(start))

    start = time.Now()
    var wg sync.WaitGroup
    // worker pool pattern...
}
```
**A:** Parallel is faster for large CPU-bound tasks (Speedup depends on core count).

---

### 26. Lazy Evaluation — Performance Optimization
**Q: Demonstrate lazy evaluation benefits using Go channels. What is the output?**
```go
package main
import "fmt"

func generate() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 1; ; i++ {
            fmt.Println("Generated:", i)
            ch <- i
        }
    }()
    return ch
}

func main() {
    nums := generate()
    for i := 0; i < 3; i++ {
        fmt.Println("Received:", <-nums)
    }
}
```
**A:** Only prints "Generated" 3 times. Channels provide **on-demand (lazy)** item production.

---

### 27. Memory Optimization — Slice Recycling
**Q: Optimize memory usage with efficient slice operations. What is the output?**
```go
// Reusing same backing array via [:0]
```

---

### 28. Custom Collector — Performance Optimization
**Q: Create custom collector for specific performance needs. What is the output?**

---

### 29. Optimization — Avoiding Common Pitfalls
**Q: Demonstrate Go optimization techniques. What is the output?**
```go
// Avoid multiple iterations by computing count/sum in one loop
```

---

### 30. Caching and Memoization — Functional Optimization
**Q: Implement caching for expensive computations in Go. What is the output?**
```go
package main
import "fmt"

func memoize(f func(int) int) func(int) int {
    cache := make(map[int]int)
    return func(n int) int {
        if v, ok := cache[n]; ok { return v }
        res := f(n)
        cache[n] = res
        return res
    }
}
```

---

### 31. Concurrent Processing — Event Stream Handling
**Q: Handle event streams using Go patterns. What is the output?**

---

### 32. Performance Monitoring — Analytics
**Q: Monitor and analyze performance metrics using Go. What is the output?**

---

### 33. Readability & Functional Patterns
**Q: How do you maintain code readability when using higher-order functions and closures in Go?**
**A:** Go is not a "functional-first" language, so readability requires extra effort:
1. **Named Functions over Closures:** If a closure is more than 5 lines, move it to a named function.
2. **Avoid Deep Nesting:** Don't nest anonymous functions more than 2 levels deep.
3. **Pointers vs. Values:** Be aware of variable capture in closures (the "loop variable trap").
4. **Happy Path on the Left:** Use guard clauses inside functions that accept or return other functions to keep the main logic clear.
