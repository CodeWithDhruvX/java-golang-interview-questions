# Go Interview Questions (Continuation 140-433)

## 🔵 Networking, APIs, and Web Dev (Questions 140-170)

### Question 140: How to build a REST API in Go?

**Answer:**
Build RESTful APIs using the `net/http` package or frameworks:

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

func getUsers(w http.ResponseWriter, r *http.Request) {
    users := []User{
        {ID: 1, Name: "John"},
        {ID: 2, Name: "Jane"},
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    
    // Save user to database...
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func main() {
    http.HandleFunc("/users", getUsers)
    http.HandleFunc("/users/create", createUser)
    http.ListenAndServe(":8080", nil)
}
```

---

### Question 141: How to parse JSON and XML in Go?

**Answer:**

**JSON Parsing:**
```go
import "encoding/json"

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Parse JSON to struct:
jsonData := []byte(`{"name":"John","age":30}`)
var person Person
json.Unmarshal(jsonData, &person)

// Convert struct to JSON:
jsonBytes, _ := json.Marshal(person)
fmt.Println(string(jsonBytes))
```

**XML Parsing:**
```go
import "encoding/xml"

type Book struct {
    Title  string `xml:"title"`
    Author string `xml:"author"`
}

// Parse XML:
xmlData := []byte(`<book><title>Go Programming</title><author>John</author></book>`)
var book Book
xml.Unmarshal(xmlData, &book)

// Convert to XML:
xmlBytes, _ := xml.Marshal(book)
```

---

### Question 142: What is the use of http.Handler and http.HandlerFunc?

**Answer:**

**http.Handler** - Interface with ServeHTTP method:
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

// Custom handler:
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

http.Handle("/", MyHandler{})
```

**http.HandlerFunc** - Function type that implements Handler:
```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

http.HandleFunc("/", myHandler)
// Or: http.Handle("/", http.HandlerFunc(myHandler))
```

---

###  Question 143: How do you implement middleware manually in Go?

**Answer:**
Create functions that wrap handlers:

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// Chain middlewares:
func main() {
    finalHandler := http.HandlerFunc(homeHandler)
    withAuth := authMiddleware(finalHandler)
    withLogging := loggingMiddleware(withAuth)
    
    http.ListenAndServe(":8080", withLogging)
}
```

---

### Question 144: How do you serve static files in Go?

**Answer:**
Use `http.FileServer`:

```go
// Serve files from ./static directory:
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))

// Or serve a specific file:
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
})

http.ListenAndServe(":8080", nil)
```

Project structure:
```
myproject/
├── static/
│   ├── css/
│   ├── js/
│   └── images/
└── main.go
```

---

### Question 145: How do you handle CORS in Go?

**Answer:**
Set CORS headers manually or use a library:

**Manual approach:**
```go
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

handler := enableCORS(http.HandlerFunc(myHandler))
```

**Using gorilla/handlers:**
```go
import "github.com/gorilla/handlers"

handler := handlers.CORS(
    handlers.AllowedOrigins([]string{"*"}),
    handlers.AllowedMethods([]string{"GET", "POST"}),
)(myHandler)
```

---

### Question 146: What are context-based timeouts in HTTP servers?

**Answer:**
Use `context.Context` to enforce timeouts:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Create timeout context:
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    
    // Make slow operation:
    result := make(chan string)
    go func() {
        // Simulate slow work:
        time.Sleep(3 * time.Second)
        result <- "Done"
    }()
    
    select {
    case res := <-result:
        fmt.Fprintf(w, res)
    case <-ctx.Done():
        http.Error(w, "Request timeout", 504)
    }
}
```

---

### Question 147: How do you make HTTP requests in Go?

**Answer:**
Use the `net/http` package:

```go
// GET request:
resp, err := http.Get("https://api.example.com/users")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body))

// POST request with JSON:
user := User{Name: "John", Age: 30}
jsonData, _ := json.Marshal(user)

resp, err = http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(jsonData),
)

// Custom request:
req, _ := http.NewRequest("PUT", "https://api.example.com/users/1", body)
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer token")

client := &http.Client{}
resp, err = client.Do(req)
```

---

### Question 148: How do you manage connection pooling in Go?

**Answer:**
Configure the HTTP client transport:

```go
transport := &http.Transport{
    MaxIdleConns:        100,              // Max idle connections
    MaxIdleConnsPerHost: 10,               // Max idle per host
    IdleConnTimeout:     90 * time.Second, // Idle timeout
}

client := &http.Client{
    Transport: transport,
    Timeout:   10 * time.Second,
}

// Reuse the client for multiple requests:
resp, err := client.Get("https://api.example.com")
```

Connection pooling happens automatically - connections are reused when possible.

---

### Question 149: What is an HTTP client timeout?

**Answer:**
Set timeout to prevent hanging requests:

```go
// Client-level timeout:
client := &http.Client{
    Timeout: 10 * time.Second,  // Total timeout for request
}

resp, err := client.Get("https://slow-api.example.com")

// Request-level timeout using context:
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
resp, err := client.Do(req)

// Different timeouts for different stages:
transport := &http.Transport{
    DialContext: (&net.Dialer{
        Timeout:   5 * time.Second,  // Connection timeout
        KeepAlive: 30 * time.Second,
    }).DialContext,
    TLSHandshakeTimeout:   10 * time.Second,  // TLS timeout
    ResponseHeaderTimeout: 10 * time.Second,   // Header timeout
}
```

---

### Question 150: How do you upload and download files via HTTP?

**Answer:**

**Upload file:**
```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Parse multipart form (32MB max):
    r.ParseMultipartForm(32 << 20)
    
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    defer file.Close()
    
    // Create file on server:
    dst, _ := os.Create("./uploads/" + handler.Filename)
    defer dst.Close()
    
    io.Copy(dst, file)
    fmt.Fprintf(w, "File uploaded successfully")
}
```

**Download file:**
```go
func downloadHandler(w http.ResponseWriter, r *http.Request) {
    filename := "document.pdf"
    
    w.Header().Set("Content-Disposition", "attachment; filename="+filename)
    w.Header().Set("Content-Type", "application/pdf")
    
    http.ServeFile(w, r, "./files/"+filename)
}
```

---

### Question 151: What is graceful shutdown and how do you implement it?

**Answer:**
Graceful shutdown finishes active requests before stopping:

```go
func main() {
    srv := &http.Server{
        Addr: ":8080",
        Handler: myHandler,
    }
    
    // Start server in goroutine:
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()
    
    // Wait for interrupt signal:
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    log.Println("Shutting down server...")
    
    // Graceful shutdown with 5-second timeout:
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }
    
    log.Println("Server exited")
}
```

---

### Question 152: How to work with multipart/form-data in Go?

**Answer:**
Handle file uploads and form fields:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Parse multipart form:
    if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB max
        http.Error(w, err.Error(), 400)
        return
    }
    
    // Get regular form fields:
    name := r.FormValue("name")
    description := r.FormValue("description")
    
    // Get uploaded files:
    files := r.MultipartForm.File["files"]
    
    for _, fileHeader := range files {
        file, _ := fileHeader.Open()
        defer file.Close()
        
        // Save file:
        dst, _ := os.Create("./uploads/" + fileHeader.Filename)
        defer dst.Close()
        
        io.Copy(dst, file)
    }
    
    fmt.Fprintf(w, "Uploaded %d files", len(files))
}
```

---

### Question 153: How do you implement rate limiting in Go?

**Answer:**
Use golang.org/x/time/rate or custom implementation:

**Using rate limiter:**
```go
import "golang.org/x/time/rate"

// Create limiter: 5 requests per second, burst of 10:
limiter := rate.NewLimiter(5, 10)

func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Rate limit exceeded", 429)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

**Custom token bucket:**
```go
type RateLimiter struct {
    tokens    int
    capacity  int
    refillRate time.Duration
    mu        sync.Mutex
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    if rl.tokens > 0 {
        rl.tokens--
        return true
    }
    return false
}
```

---

*[Due to length constraints, I'm creating this as a separate file. The user can merge this with the main file.]*
