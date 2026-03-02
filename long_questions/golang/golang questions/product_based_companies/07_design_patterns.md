# 🧩 07 — Design Patterns in Go
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- Creational: Factory, Builder, Singleton
- Structural: Decorator/Middleware, Adapter
- Behavioral: Strategy, Observer, Command
- Go-specific: Functional Options, Repository, Clean Architecture

---

## ❓ Most Asked Questions

### Q1. Factory Pattern

```go
type Notification interface {
    Send(message string) error
}

type EmailNotification struct{ Address string }
type SMSNotification struct{ Phone string }
type PushNotification struct{ DeviceToken string }

func (e *EmailNotification) Send(msg string) error {
    fmt.Printf("Email to %s: %s\n", e.Address, msg); return nil
}
func (s *SMSNotification) Send(msg string) error {
    fmt.Printf("SMS to %s: %s\n", s.Phone, msg); return nil
}
func (p *PushNotification) Send(msg string) error {
    fmt.Printf("Push to %s: %s\n", p.DeviceToken, msg); return nil
}

// Factory function
func NewNotification(notifType, target string) (Notification, error) {
    switch notifType {
    case "email": return &EmailNotification{Address: target}, nil
    case "sms":   return &SMSNotification{Phone: target}, nil
    case "push":  return &PushNotification{DeviceToken: target}, nil
    default:      return nil, fmt.Errorf("unknown notification type: %s", notifType)
    }
}

// Usage
n, _ := NewNotification("email", "user@example.com")
n.Send("Welcome!")
```

---

### Q2. Functional Options Pattern (Go-specific)

```go
// Common Go pattern for optional configuration
type Server struct {
    host    string
    port    int
    timeout time.Duration
    maxConn int
    tls     bool
}

type Option func(*Server)

func WithHost(host string) Option    { return func(s *Server) { s.host = host } }
func WithPort(port int) Option       { return func(s *Server) { s.port = port } }
func WithTimeout(d time.Duration) Option { return func(s *Server) { s.timeout = d } }
func WithTLS() Option                { return func(s *Server) { s.tls = true } }
func WithMaxConns(n int) Option      { return func(s *Server) { s.maxConn = n } }

func NewServer(opts ...Option) *Server {
    // Default values
    s := &Server{
        host:    "localhost",
        port:    8080,
        timeout: 30 * time.Second,
        maxConn: 100,
    }
    for _, opt := range opts { opt(s) }  // apply each option
    return s
}

// Clean, extensible usage
srv := NewServer(
    WithHost("0.0.0.0"),
    WithPort(9090),
    WithTLS(),
    WithTimeout(60*time.Second),
)
```

---

### Q3. Strategy Pattern

```go
// Sort strategy
type SortStrategy interface {
    Sort([]int) []int
}

type BubbleSort struct{}
type QuickSort  struct{}

func (b BubbleSort) Sort(nums []int) []int {
    n := len(nums)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if nums[j] > nums[j+1] { nums[j], nums[j+1] = nums[j+1], nums[j] }
        }
    }
    return nums
}

func (q QuickSort) Sort(nums []int) []int {
    sort.Ints(nums); return nums  // using stdlib for brevity
}

// Context that uses a strategy
type Sorter struct{ strategy SortStrategy }
func (s *Sorter) SetStrategy(strategy SortStrategy) { s.strategy = strategy }
func (s *Sorter) Sort(nums []int) []int              { return s.strategy.Sort(nums) }

// Usage — swap strategy at runtime
sorter := &Sorter{strategy: BubbleSort{}}
sorter.Sort([]int{5, 2, 8, 1})
sorter.SetStrategy(QuickSort{})  // change strategy
sorter.Sort([]int{5, 2, 8, 1})
```

---

### Q4. Observer Pattern using Channels

```go
type EventType string
const (
    OrderCreated EventType = "order.created"
    OrderShipped EventType = "order.shipped"
)

type Event struct {
    Type    EventType
    Payload interface{}
}

type EventBus struct {
    mu          sync.RWMutex
    subscribers map[EventType][]chan Event
}

func NewEventBus() *EventBus {
    return &EventBus{subscribers: make(map[EventType][]chan Event)}
}

func (eb *EventBus) Subscribe(t EventType) <-chan Event {
    ch := make(chan Event, 10)
    eb.mu.Lock()
    eb.subscribers[t] = append(eb.subscribers[t], ch)
    eb.mu.Unlock()
    return ch
}

func (eb *EventBus) Publish(e Event) {
    eb.mu.RLock()
    defer eb.mu.RUnlock()
    for _, ch := range eb.subscribers[e.Type] {
        select { case ch <- e: default: }
    }
}

// Usage
bus := NewEventBus()
ch := bus.Subscribe(OrderCreated)
go func() { for e := range ch { fmt.Println("Order event:", e.Payload) } }()
bus.Publish(Event{Type: OrderCreated, Payload: "order#42"})
```

---

### Q5. Repository Pattern (Clean Architecture)

```go
// Domain entity
type User struct {
    ID    int
    Name  string
    Email string
}

// Repository interface — domain layer
type UserRepository interface {
    GetByID(ctx context.Context, id int) (*User, error)
    Save(ctx context.Context, u *User) error
    Delete(ctx context.Context, id int) error
}

// Concrete implementation — infrastructure layer
type PostgresUserRepo struct{ db *sql.DB }

func (r *PostgresUserRepo) GetByID(ctx context.Context, id int) (*User, error) {
    row := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id=$1", id)
    var u User
    if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil { return nil, err }
    return &u, nil
}

// Service layer — uses interface, testable with mock
type UserService struct{ repo UserRepository }

func NewUserService(repo UserRepository) *UserService { return &UserService{repo: repo} }

func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
    return s.repo.GetByID(ctx, id)
}
```

---

### Q6. Builder Pattern

```go
type QueryBuilder struct {
    table      string
    conditions []string
    orderBy    string
    limit      int
    args       []interface{}
}

func NewQuery(table string) *QueryBuilder { return &QueryBuilder{table: table} }

func (q *QueryBuilder) Where(cond string, args ...interface{}) *QueryBuilder {
    q.conditions = append(q.conditions, cond)
    q.args = append(q.args, args...)
    return q
}

func (q *QueryBuilder) OrderBy(field string) *QueryBuilder { q.orderBy = field; return q }
func (q *QueryBuilder) Limit(n int) *QueryBuilder          { q.limit = n; return q }

func (q *QueryBuilder) Build() (string, []interface{}) {
    sql := "SELECT * FROM " + q.table
    if len(q.conditions) > 0 {
        sql += " WHERE " + strings.Join(q.conditions, " AND ")
    }
    if q.orderBy != "" { sql += " ORDER BY " + q.orderBy }
    if q.limit > 0 { sql += fmt.Sprintf(" LIMIT %d", q.limit) }
    return sql, q.args
}

// Usage
query, args := NewQuery("users").
    Where("age > ?", 18).
    Where("active = ?", true).
    OrderBy("name").
    Limit(10).
    Build()
```
