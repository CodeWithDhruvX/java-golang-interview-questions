# 📒 05 — REST API & Web Development in Go
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- `net/http` package (Handler, ServeMux)
- JSON encoding/decoding
- Gin framework basics
- Middleware pattern
- CORS handling
- Request validation
- Status codes and response patterns

---

## ❓ Most Asked Questions

### Q1. How do you build a basic REST API with `net/http`?

```go
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    user := User{ID: 1, Name: "Alice"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/user", getUser)
    http.ListenAndServe(":8080", mux)
}
```

---

### Q2. How do you build a REST API using Gin?

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

var products = []Product{
    {1, "Laptop", 999.99},
    {2, "Phone", 499.99},
}

func main() {
    r := gin.Default()

    // GET all products
    r.GET("/products", func(c *gin.Context) {
        c.JSON(http.StatusOK, products)
    })

    // GET product by ID
    r.GET("/products/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{"id": id})
    })

    // POST create product
    r.POST("/products", func(c *gin.Context) {
        var p Product
        if err := c.ShouldBindJSON(&p); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        products = append(products, p)
        c.JSON(http.StatusCreated, p)
    })

    r.Run(":8080")
}
```

---

### Q3. How does JSON marshaling/unmarshaling work?

```go
import "encoding/json"

type Order struct {
    ID       int       `json:"id"`
    Product  string    `json:"product"`
    Quantity int       `json:"quantity"`
    Price    float64   `json:"price"`
}

// Struct → JSON (Marshal)
order := Order{ID: 1, Product: "Laptop", Quantity: 2, Price: 999.99}
data, err := json.Marshal(order)
if err != nil { panic(err) }
fmt.Println(string(data))
// {"id":1,"product":"Laptop","quantity":2,"price":999.99}

// JSON → Struct (Unmarshal)
jsonStr := `{"id":2,"product":"Phone","quantity":1,"price":499.99}`
var o2 Order
json.Unmarshal([]byte(jsonStr), &o2)
fmt.Println(o2.Product)  // Phone

// Using Decoder (for http.Request.Body)
func handleOrder(w http.ResponseWriter, r *http.Request) {
    var order Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // process order
}
```

---

### Q4. What is `http.Handler` and `http.HandlerFunc`?

```go
// http.Handler is an interface
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

// Custom handler by implementing the interface
type HomeHandler struct{}
func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome Home!")
}
http.Handle("/", HomeHandler{})

// http.HandlerFunc converts a function to a Handler
func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
}
http.HandleFunc("/hello", hello)
// Equivalent to:
http.Handle("/hello", http.HandlerFunc(hello))
```

---

### Q5. How do you implement middleware in Go?

```go
// Middleware is a function that wraps a handler
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)  // call next handler
        fmt.Printf("Completed in %v\n", time.Since(start))
    })
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// Chain middleware
mux := http.NewServeMux()
mux.HandleFunc("/api/data", dataHandler)
http.ListenAndServe(":8080", LoggingMiddleware(AuthMiddleware(mux)))

// Gin middleware
r := gin.New()
r.Use(gin.Logger())
r.Use(gin.Recovery())
```

---

### Q6. How do you handle query parameters and path parameters?

```go
// net/http — query params
func search(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    name  := query.Get("name")    // ?name=Alice
    page  := query.Get("page")    // ?page=2
    fmt.Fprintf(w, "name=%s, page=%s", name, page)
}

// Gin — path params + query params
r.GET("/users/:id", func(c *gin.Context) {
    id   := c.Param("id")           // /users/42
    role := c.Query("role")         // ?role=admin
    page := c.DefaultQuery("page", "1")  // ?page=2, default "1"
    c.JSON(200, gin.H{"id": id, "role": role, "page": page})
})
```

---

### Q7. How do you handle CORS in Go?

```go
// Manual CORS middleware
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// With Gin + cors package
import "github.com/gin-contrib/cors"
r := gin.Default()
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://myapp.com"},
    AllowMethods:     []string{"GET", "POST", "PUT"},
    AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
}))
```

---

### Q8. How do you implement graceful shutdown?

```go
func main() {
    srv := &http.Server{
        Addr:    ":8080",
        Handler: http.DefaultServeMux,
    }

    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe error: %v", err)
        }
    }()

    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit
    fmt.Println("Shutting down server...")

    // Give 5 seconds for ongoing requests to complete
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }
    fmt.Println("Server exited")
}
```

---

### Q9. How do you validate request input?

```go
// Using go-playground/validator
import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
    Name     string `json:"name" validate:"required,min=2,max=50"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"required,gte=18,lte=120"`
    Password string `json:"password" validate:"required,min=8"`
}

var validate = validator.New()

func createUser(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := validate.Struct(req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // process...
    c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
```

---

### Q10. What are common HTTP status codes you use in REST APIs?

| Code | Constant | Meaning |
|------|---------|---------|
| 200 | `StatusOK` | Success |
| 201 | `StatusCreated` | Resource created |
| 204 | `StatusNoContent` | Success, no body |
| 400 | `StatusBadRequest` | Invalid request data |
| 401 | `StatusUnauthorized` | Not authenticated |
| 403 | `StatusForbidden` | Not authorized |
| 404 | `StatusNotFound` | Resource not found |
| 409 | `StatusConflict` | Duplicate resource |
| 422 | `StatusUnprocessableEntity` | Validation failed |
| 429 | `StatusTooManyRequests` | Rate limit exceeded |
| 500 | `StatusInternalServerError` | Server error |
