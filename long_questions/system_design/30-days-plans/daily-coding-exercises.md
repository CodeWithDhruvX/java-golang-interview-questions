# Daily Coding Exercises for System Design

> **Hands-on coding exercises for each day of the 30-day learning path**
> 
> **Languages:** Go, Java, Python | **Focus:** Practical implementation of system design concepts

---

## **Phase 1: Introduction (Day 1-4)**

### **Day 1: System Design Basics**
**Exercise: Simple Blog System**
```go
// Go - Basic blog structure
type Blog struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Author    string    `json:"author"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type BlogSystem struct {
    blogs map[string]*Blog
    mu    sync.RWMutex
}

func (bs *BlogSystem) CreateBlog(title, content, author string) *Blog {
    bs.mu.Lock()
    defer bs.mu.Unlock()
    
    blog := &Blog{
        ID:        generateID(),
        Title:     title,
        Content:   content,
        Author:    author,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    
    bs.blogs[blog.ID] = blog
    return blog
}
```

**Task:** Extend this basic blog system with:
1. Add comment functionality
2. Implement basic caching for popular posts
3. Draw system diagram showing components

---

### **Day 2: Interview Framework**
**Exercise: Requirements Gathering**
```java
// Java - Requirements analysis framework
public class SystemRequirements {
    private long dailyActiveUsers;
    private long requestsPerSecond;
    private double readWriteRatio;
    private Map<String, Object> functionalRequirements;
    private Map<String, Object> nonFunctionalRequirements;
    
    public SystemRequirements clarifyRequirements(String problem) {
        System.out.println("Clarifying requirements for: " + problem);
        
        // Ask clarifying questions
        Scanner scanner = new Scanner(System.in);
        
        System.out.print("Expected DAU? ");
        this.dailyActiveUsers = scanner.nextLong();
        
        System.out.print("Expected QPS? ");
        this.requestsPerSecond = scanner.nextLong();
        
        System.out.print("Read/Write ratio (e.g., 80:20)? ");
        String ratio = scanner.next();
        // Parse ratio...
        
        return this;
    }
    
    public void estimateScale() {
        long storageNeeded = dailyActiveUsers * 10; // 10KB per user
        long bandwidthNeeded = requestsPerSecond * 1024; // 1KB per request
        
        System.out.println("Storage needed: " + storageNeeded + " KB");
        System.out.println("Bandwidth needed: " + bandwidthNeeded + " KB/s");
    }
}
```

**Task:** Use this framework to analyze:
1. URL shortener requirements
2. Chat application requirements
3. Practice explaining the 4-step framework

---

### **Day 3: HLD vs LLD**
**Exercise: Chat Application Design**
```python
# Python - HLD components
class ChatSystemHLD:
    def __init__(self):
        self.components = {
            "API Gateway": "Handles all incoming requests",
            "User Service": "Manages user profiles and authentication",
            "Message Service": "Handles message storage and retrieval",
            "WebSocket Gateway": "Manages real-time connections",
            "Notification Service": "Sends push notifications",
            "Database": "PostgreSQL for user data, MongoDB for messages",
            "Cache": "Redis for session management",
            "Message Queue": "Kafka for async processing"
        }
    
    def draw_architecture(self):
        print("=== High-Level Design ===")
        for component, description in self.components.items():
            print(f"{component}: {description}")

# LLD - Message entity design
from dataclasses import dataclass
from datetime import datetime
from typing import Optional

@dataclass
class Message:
    id: str
    sender_id: str
    receiver_id: str
    content: str
    message_type: str  # text, image, file
    timestamp: datetime
    status: str  # sent, delivered, read
    reply_to: Optional[str] = None
    
class MessageService:
    def __init__(self):
        self.messages = {}  # In-memory storage for demo
    
    def send_message(self, sender_id: str, receiver_id: str, content: str) -> Message:
        message = Message(
            id=self.generate_id(),
            sender_id=sender_id,
            receiver_id=receiver_id,
            content=content,
            message_type="text",
            timestamp=datetime.now(),
            status="sent"
        )
        
        self.messages[message.id] = message
        return message
    
    def get_conversation(self, user1_id: str, user2_id: str) -> list[Message]:
        conversation = []
        for message in self.messages.values():
            if (message.sender_id == user1_id and message.receiver_id == user2_id) or \
               (message.sender_id == user2_id and message.receiver_id == user1_id):
                conversation.append(message)
        
        return sorted(conversation, key=lambda x: x.timestamp)
```

**Task:** 
1. Create HLD diagram for chat system
2. Design LLD for user entity
3. Explain when to use HLD vs LLD

---

### **Day 4: Trade-offs Analysis**
**Exercise: Caching Strategies**
```go
// Go - Different caching strategies
type CacheStrategy interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{})
    Delete(key string)
}

// Write-through cache
type WriteThroughCache struct {
    cache     map[string]interface{}
    database  map[string]interface{}
    mu        sync.RWMutex
}

func (wtc *WriteThroughCache) Set(key string, value interface{}) {
    wtc.mu.Lock()
    defer wtc.mu.Unlock()
    
    // Write to cache first
    wtc.cache[key] = value
    // Then write to database
    wtc.database[key] = value
}

// Write-behind cache
type WriteBehindCache struct {
    cache      map[string]interface{}
    database   map[string]interface{}
    pending    map[string]interface{}
    mu         sync.RWMutex
}

func (wbc *WriteBehindCache) Set(key string, value interface{}) {
    wbc.mu.Lock()
    defer wbc.mu.Unlock()
    
    // Write to cache immediately
    wbc.cache[key] = value
    // Mark for database write
    wbc.pending[key] = value
    
    // Simulate async write to database
    go func() {
        time.Sleep(100 * time.Millisecond)
        wbc.mu.Lock()
        wbc.database[key] = value
        delete(wbc.pending, key)
        wbc.mu.Unlock()
    }()
}

func analyzeTradeoffs() {
    fmt.Println("=== Cache Strategy Trade-offs ===")
    fmt.Println("Write-through:")
    fmt.Println("  Pros: Data consistency, immediate persistence")
    fmt.Println("  Cons: Higher latency, writes are blocking")
    fmt.Println()
    fmt.Println("Write-behind:")
    fmt.Println("  Pros: Low latency, high throughput")
    fmt.Println("  Cons: Risk of data loss, eventual consistency")
}
```

**Task:**
1. Implement both caching strategies
2. Test performance differences
3. Document trade-offs for different scenarios

---

## **Phase 2: Fundamentals (Day 5-13)**

### **Day 5: Scalability Exercise**
**Exercise: Scaling Calculator**
```python
# Python - Scaling calculations
class ScalingCalculator:
    def __init__(self):
        self.current_users = 1000
        self.current_qps = 100
        self.target_users = 1000000
        
    def calculate_requirements(self):
        # Calculate growth factor
        growth_factor = self.target_users / self.current_users
        
        # Estimate new QPS (assuming same usage patterns)
        target_qps = self.current_qps * growth_factor
        
        # Calculate storage needs (10KB per user)
        storage_needed = self.target_users * 10 / (1024 * 1024)  # MB
        
        # Calculate bandwidth needs (1KB per request)
        bandwidth_needed = target_qps * 1024 / (1024 * 1024)  # MB/s
        
        return {
            "growth_factor": growth_factor,
            "target_qps": target_qps,
            "storage_mb": storage_needed,
            "bandwidth_mbps": bandwidth_needed
        }
    
    def suggest_scaling_strategy(self, requirements):
        print("=== Scaling Strategy ===")
        if requirements["growth_factor"] > 100:
            print("Recommendation: Horizontal scaling with microservices")
            print("- Use load balancers")
            print("- Implement database sharding")
            print("- Add caching layer")
        elif requirements["growth_factor"] > 10:
            print("Recommendation: Vertical scaling + read replicas")
            print("- Upgrade server resources")
            print("- Add read replicas")
            print("- Implement caching")
        else:
            print("Recommendation: Vertical scaling")
            print("- Upgrade server resources")
            print("- Optimize database queries")
```

**Task:**
1. Calculate scaling requirements for different scenarios
2. Design scaling strategy for e-commerce platform
3. Create capacity planning document

---

### **Day 6: Client-Server Architecture**
**Exercise: Protocol Selection**
```java
// Java - Protocol comparison
public class ProtocolSelector {
    
    public enum Protocol {
        HTTP, HTTPS, WEBSOCKET, TCP, UDP
    }
    
    public static Protocol selectProtocol(String useCase) {
        switch (useCase.toLowerCase()) {
            case "web_api":
                return Protocol.HTTPS;  // Secure web communication
            case "real_time_chat":
                return Protocol.WEBSOCKET;  // Bidirectional communication
            case "file_transfer":
                return Protocol.TCP;  // Reliable data transfer
            case "video_streaming":
                return Protocol.UDP;  // Fast, loss-tolerant
            case "mobile_app_api":
                return Protocol.HTTPS;  // Secure mobile communication
            default:
                return Protocol.HTTP;
        }
    }
    
    public static void explainChoice(String useCase, Protocol protocol) {
        System.out.println("Use case: " + useCase);
        System.out.println("Selected protocol: " + protocol);
        System.out.println("Reasoning: " + getReasoning(protocol));
    }
    
    private static String getReasoning(Protocol protocol) {
        switch (protocol) {
            case HTTPS:
                return "Secure, reliable, widely supported for web APIs";
            case WEBSOCKET:
                return "Real-time, bidirectional, low latency";
            case TCP:
                return "Reliable, ordered delivery, connection-oriented";
            case UDP:
                return "Fast, lightweight, suitable for streaming";
            default:
                return "Standard web protocol";
        }
    }
}
```

**Task:**
1. Implement protocol selector for different use cases
2. Create API design for user management system
3. Document protocol selection criteria

---

### **Day 7: API Design Exercise**
**Exercise: REST API Design**
```go
// Go - REST API structure
type UserAPI struct {
    userService *UserService
    router      *mux.Router
}

func (api *UserAPI) setupRoutes() {
    // User CRUD operations
    api.router.HandleFunc("/api/v1/users", api.getUsers).Methods("GET")
    api.router.HandleFunc("/api/v1/users", api.createUser).Methods("POST")
    api.router.HandleFunc("/api/v1/users/{id}", api.getUser).Methods("GET")
    api.router.HandleFunc("/api/v1/users/{id}", api.updateUser).Methods("PUT")
    api.router.HandleFunc("/api/v1/users/{id}", api.deleteUser).Methods("DELETE")
    
    // Additional endpoints
    api.router.HandleFunc("/api/v1/users/{id}/posts", api.getUserPosts).Methods("GET")
    api.router.HandleFunc("/api/v1/users/search", api.searchUsers).Methods("GET")
}

func (api *UserAPI) createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    createdUser := api.userService.Create(user)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdUser)
}

// GraphQL alternative
type GraphQLResolver struct {
    userService *UserService
}

func (r *GraphQLResolver) CreateUser(ctx context.Context, input UserInput) (*User, error) {
    user := User{
        Name:  input.Name,
        Email: input.Email,
    }
    return r.userService.Create(user), nil
}

func (r *GraphQLResolver) GetUser(ctx context.Context, id string) (*User, error) {
    return r.userService.GetByID(id)
}
```

**Task:**
1. Design complete REST API for blogging platform
2. Create GraphQL schema equivalent
3. Compare REST vs GraphQL for mobile app use case

---

### **Day 8: Load Balancing Exercise**
**Exercise: Load Balancer Algorithms**
```python
# Python - Load balancing algorithms
import random
import hashlib
from typing import List, Dict

class Server:
    def __init__(self, id: str, weight: int = 1):
        self.id = id
        self.weight = weight
        self.connections = 0
        self.healthy = True
    
    def add_connection(self):
        self.connections += 1
    
    def remove_connection(self):
        self.connections = max(0, self.connections - 1)

class LoadBalancer:
    def __init__(self, servers: List[Server]):
        self.servers = servers
        self.current_index = 0
    
    def round_robin(self) -> Server:
        healthy_servers = [s for s in self.servers if s.healthy]
        if not healthy_servers:
            raise Exception("No healthy servers available")
        
        server = healthy_servers[self.current_index % len(healthy_servers)]
        self.current_index += 1
        server.add_connection()
        return server
    
    def weighted_round_robin(self) -> Server:
        healthy_servers = [s for s in self.servers if s.healthy]
        if not healthy_servers:
            raise Exception("No healthy servers available")
        
        # Create weighted list
        weighted_servers = []
        for server in healthy_servers:
            weighted_servers.extend([server] * server.weight)
        
        server = weighted_servers[self.current_index % len(weighted_servers)]
        self.current_index += 1
        server.add_connection()
        return server
    
    def least_connections(self) -> Server:
        healthy_servers = [s for s in self.servers if s.healthy]
        if not healthy_servers:
            raise Exception("No healthy servers available")
        
        server = min(healthy_servers, key=lambda s: s.connections)
        server.add_connection()
        return server
    
    def consistent_hash(self, key: str) -> Server:
        healthy_servers = [s for s in self.servers if s.healthy]
        if not healthy_servers:
            raise Exception("No healthy servers available")
        
        # Simple hash ring implementation
        hash_value = int(hashlib.md5(key.encode()).hexdigest(), 16)
        server_index = hash_value % len(healthy_servers)
        server = healthy_servers[server_index]
        server.add_connection()
        return server

# Health check simulation
def health_check_simulation():
    servers = [
        Server("server1", weight=1),
        Server("server2", weight=2),
        Server("server3", weight=3),
    ]
    
    lb = LoadBalancer(servers)
    
    # Simulate server failure
    servers[1].healthy = False
    
    print("=== Load Balancing with Server Failure ===")
    for i in range(10):
        try:
            server = lb.round_robin()
            print(f"Request {i+1}: {server.id}")
        except Exception as e:
            print(f"Request {i+1}: {e}")
```

**Task:**
1. Implement all load balancing algorithms
2. Test with different server configurations
3. Simulate server failures and failover

---

### **Day 9: Caching Exercise**
**Exercise: Cache Implementation**
```java
// Java - LRU Cache implementation
public class LRUCache<K, V> {
    private final int capacity;
    private final Map<K, Node<K, V>> cache;
    private final DoublyLinkedList<K, V> usageList;
    
    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.cache = new HashMap<>();
        this.usageList = new DoublyLinkedList<>();
    }
    
    public V get(K key) {
        Node<K, V> node = cache.get(key);
        if (node == null) {
            return null;
        }
        
        // Move to front (most recently used)
        usageList.moveToFront(node);
        return node.value;
    }
    
    public void put(K key, V value) {
        Node<K, V> existing = cache.get(key);
        
        if (existing != null) {
            // Update existing node
            existing.value = value;
            usageList.moveToFront(existing);
        } else {
            // Add new node
            Node<K, V> newNode = new Node<>(key, value);
            
            if (cache.size() >= capacity) {
                // Remove least recently used
                Node<K, V> lru = usageList.removeTail();
                cache.remove(lru.key);
            }
            
            cache.put(key, newNode);
            usageList.addToFront(newNode);
        }
    }
    
    public void display() {
        System.out.print("Cache: ");
        Node<K, V> current = usageList.head;
        while (current != null) {
            System.out.print(current.key + "=" + current.value + " ");
            current = current.next;
        }
        System.out.println();
    }
}

class Node<K, V> {
    K key;
    V value;
    Node<K, V> prev;
    Node<K, V> next;
    
    public Node(K key, V value) {
        this.key = key;
        this.value = value;
    }
}

class DoublyLinkedList<K, V> {
    Node<K, V> head;
    Node<K, V> tail;
    
    public void addToFront(Node<K, V> node) {
        node.prev = null;
        node.next = head;
        
        if (head != null) {
            head.prev = node;
        } else {
            tail = node;
        }
        
        head = node;
    }
    
    public void moveToFront(Node<K, V> node) {
        if (node == head) return;
        
        removeNode(node);
        addToFront(node);
    }
    
    public Node<K, V> removeTail() {
        if (tail == null) return null;
        
        Node<K, V> removed = tail;
        tail = tail.prev;
        
        if (tail != null) {
            tail.next = null;
        } else {
            head = null;
        }
        
        return removed;
    }
    
    private void removeNode(Node<K, V> node) {
        if (node.prev != null) {
            node.prev.next = node.next;
        } else {
            head = node.next;
        }
        
        if (node.next != null) {
            node.next.prev = node.prev;
        } else {
            tail = node.prev;
        }
    }
}
```

**Task:**
1. Implement LRU cache
2. Add TTL (Time To Live) functionality
3. Implement cache invalidation strategies

---

### **Day 10: Database Selection Exercise**
**Exercise: Database Comparison**
```python
# Python - Database selection framework
class DatabaseSelector:
    def __init__(self):
        self.criteria = {
            "consistency": ["strong", "eventual"],
            "scalability": ["vertical", "horizontal"],
            "data_structure": ["structured", "semi_structured", "unstructured"],
            "query_pattern": ["simple", "complex_joins", "graph", "time_series"],
            "transaction_need": ["acid", "base"]
        }
    
    def recommend_database(self, requirements):
        recommendations = []
        
        # SQL databases for structured data with ACID
        if (requirements["data_structure"] == "structured" and 
            requirements["transaction_need"] == "acid"):
            recommendations.append("PostgreSQL")
            recommendations.append("MySQL")
        
        # NoSQL for horizontal scaling
        if requirements["scalability"] == "horizontal":
            if requirements["data_structure"] == "semi_structured":
                recommendations.append("MongoDB")
            elif requirements["query_pattern"] == "time_series":
                recommendations.append("InfluxDB")
            elif requirements["query_pattern"] == "graph":
                recommendations.append("Neo4j")
            else:
                recommendations.append("Cassandra")
        
        # Redis for caching
        if requirements["consistency"] == "eventual" and requirements.get("latency_critical"):
            recommendations.append("Redis")
        
        return recommendations
    
    def analyze_tradeoffs(self, db1, db2):
        tradeoffs = {
            "PostgreSQL vs MongoDB": {
                "PostgreSQL": {
                    "pros": ["ACID transactions", "Complex queries", "Data integrity"],
                    "cons": ["Vertical scaling limits", "Rigid schema"]
                },
                "MongoDB": {
                    "pros": ["Horizontal scaling", "Flexible schema", "Document-oriented"],
                    "cons": ["Eventual consistency", "Limited transactions"]
                }
            }
        }
        
        return tradeoffs.get(f"{db1} vs {db2}", {})

# CAP theorem demonstration
class CAPTheorem:
    def __init__(self):
        self.properties = {
            "C": "Consistency - All nodes see same data",
            "A": "Availability - System always responds",
            "P": "Partition Tolerance - System works despite network splits"
        }
    
    def analyze_database(self, db_name):
        cap_properties = {
            "PostgreSQL": "CA",  # Sacrifices partition tolerance
            "MongoDB": "CP",     # Sacrifices availability
            "Cassandra": "AP",   # Sacrifices consistency
            "Redis": "AP"        # Sacrifices consistency
        }
        
        cap = cap_properties.get(db_name, "Unknown")
        print(f"{db_name}: {cap}")
        
        for prop in cap:
            print(f"  {self.properties[prop]}")
        
        # What's sacrificed
        sacrificed = {"C", "A", "P"} - set(cap)
        for prop in sacrificed:
            print(f"  Sacrificed: {self.properties[prop]}")
```

**Task:**
1. Use database selector for different scenarios
2. Analyze CAP theorem for various databases
3. Create database decision matrix

---

## **Phase 3: Core Components (Day 14-17)**

### **Day 14: Message Queue Exercise**
**Exercise: Simple Message Queue**
```go
// Go - Simple message queue implementation
type Message struct {
    ID        string        `json:"id"`
    Topic     string        `json:"topic"`
    Payload   interface{}   `json:"payload"`
    Timestamp time.Time     `json:"timestamp"`
    Attempts  int           `json:"attempts"`
}

type Subscriber struct {
    ID       string
    Callback func(Message) error
}

type MessageQueue struct {
    queues    map[string][]Message
    subscribers map[string][]*Subscriber
    mu        sync.RWMutex
    processed chan Message
    failed    chan Message
}

func NewMessageQueue() *MessageQueue {
    mq := &MessageQueue{
        queues:      make(map[string][]Message),
        subscribers: make(map[string][]*Subscriber),
        processed:   make(chan Message, 100),
        failed:      make(chan Message, 100),
    }
    
    go mq.processMessages()
    return mq
}

func (mq *MessageQueue) Publish(topic string, payload interface{}) error {
    mq.mu.Lock()
    defer mq.mu.Unlock()
    
    message := Message{
        ID:        generateID(),
        Topic:     topic,
        Payload:   payload,
        Timestamp: time.Now(),
        Attempts:  0,
    }
    
    mq.queues[topic] = append(mq.queues[topic], message)
    return nil
}

func (mq *MessageQueue) Subscribe(topic string, callback func(Message) error) string {
    mq.mu.Lock()
    defer mq.mu.Unlock()
    
    subscriberID := generateID()
    subscriber := &Subscriber{
        ID:       subscriberID,
        Callback: callback,
    }
    
    mq.subscribers[topic] = append(mq.subscribers[topic], subscriber)
    return subscriberID
}

func (mq *MessageQueue) processMessages() {
    for {
        mq.mu.RLock()
        
        for topic, messages := range mq.queues {
            if len(messages) > 0 {
                message := messages[0]
                subscribers := mq.subscribers[topic]
                
                mq.mu.RUnlock()
                
                success := true
                for _, sub := range subscribers {
                    if err := sub.Callback(message); err != nil {
                        success = false
                        break
                    }
                }
                
                mq.mu.Lock()
                if success {
                    // Remove processed message
                    mq.queues[topic] = messages[1:]
                    mq.processed <- message
                } else {
                    // Retry logic
                    message.Attempts++
                    if message.Attempts < 3 {
                        // Put back at end of queue
                        mq.queues[topic] = append(messages[1:], message)
                    } else {
                        // Move to failed
                        mq.queues[topic] = messages[1:]
                        mq.failed <- message
                    }
                }
                mq.mu.Unlock()
                
                continue
            }
        }
        
        mq.mu.RUnlock()
        time.Sleep(100 * time.Millisecond)
    }
}
```

**Task:**
1. Implement message queue with pub/sub
2. Add message durability (persistence)
3. Implement at-least-once delivery

---

### **Day 15: Microservices Exercise**
**Exercise: Service Discovery**
```python
# Python - Service discovery implementation
import json
import time
from typing import Dict, List
import requests

class ServiceRegistry:
    def __init__(self):
        self.services: Dict[str, List[Dict]] = {}
        self.health_check_interval = 30
    
    def register_service(self, service_name: str, service_url: str, health_check_url: str):
        service_info = {
            "url": service_url,
            "health_check_url": health_check_url,
            "registered_at": time.time(),
            "last_health_check": time.time(),
            "healthy": True
        }
        
        if service_name not in self.services:
            self.services[service_name] = []
        
        self.services[service_name].append(service_info)
        print(f"Service {service_name} registered at {service_url}")
    
    def discover_service(self, service_name: str) -> str:
        if service_name not in self.services:
            raise Exception(f"Service {service_name} not found")
        
        # Return first healthy service
        for service in self.services[service_name]:
            if service["healthy"]:
                return service["url"]
        
        raise Exception(f"No healthy instances of {service_name}")
    
    def health_check(self):
        while True:
            for service_name, instances in self.services.items():
                for instance in instances:
                    try:
                        response = requests.get(instance["health_check_url"], timeout=5)
                        if response.status_code == 200:
                            instance["healthy"] = True
                            instance["last_health_check"] = time.time()
                        else:
                            instance["healthy"] = False
                    except:
                        instance["healthy"] = False
            
            time.sleep(self.health_check_interval)

class APIGateway:
    def __init__(self, service_registry: ServiceRegistry):
        self.service_registry = service_registry
        self.routes = {
            "/api/users": "user-service",
            "/api/orders": "order-service",
            "/api/products": "product-service"
        }
    
    def route_request(self, path: str, method: str, data: dict = None):
        service_name = self.routes.get(path)
        if not service_name:
            return {"error": "Route not found"}, 404
        
        try:
            service_url = self.service_registry.discover_service(service_name)
            full_url = f"{service_url}{path}"
            
            if method == "GET":
                response = requests.get(full_url)
            elif method == "POST":
                response = requests.post(full_url, json=data)
            else:
                return {"error": "Method not supported"}, 405
            
            return response.json(), response.status_code
            
        except Exception as e:
            return {"error": str(e)}, 503

# Usage example
def setup_microservices():
    registry = ServiceRegistry()
    
    # Register services
    registry.register_service(
        "user-service", 
        "http://localhost:8001",
        "http://localhost:8001/health"
    )
    
    registry.register_service(
        "order-service", 
        "http://localhost:8002",
        "http://localhost:8002/health"
    )
    
    # Start health checking
    import threading
    health_thread = threading.Thread(target=registry.health_check, daemon=True)
    health_thread.start()
    
    # Setup API Gateway
    gateway = APIGateway(registry)
    
    return gateway
```

**Task:**
1. Implement service registry
2. Add load balancing for multiple instances
3. Implement circuit breaker pattern

---

### **Day 16: Rate Limiting Exercise**
**Exercise: Token Bucket Algorithm**
```java
// Java - Token bucket rate limiter
public class TokenBucketRateLimiter {
    private final int capacity;
    private final double refillRate;
    private double tokens;
    private long lastRefillTimestamp;
    
    public TokenBucketRateLimiter(int capacity, double refillRatePerSecond) {
        this.capacity = capacity;
        this.refillRate = refillRatePerSecond;
        this.tokens = capacity;
        this.lastRefillTimestamp = System.currentTimeMillis();
    }
    
    public synchronized boolean allowRequest() {
        refillTokens();
        
        if (tokens >= 1) {
            tokens--;
            return true;
        }
        
        return false;
    }
    
    private void refillTokens() {
        long now = System.currentTimeMillis();
        long timePassed = now - lastRefillTimestamp;
        
        // Calculate tokens to add based on time passed
        double tokensToAdd = (timePassed / 1000.0) * refillRate;
        
        tokens = Math.min(capacity, tokens + tokensToAdd);
        lastRefillTimestamp = now;
    }
    
    public synchronized double getAvailableTokens() {
        refillTokens();
        return tokens;
    }
}

// Sliding window rate limiter
public class SlidingWindowRateLimiter {
    private final int windowSizeInSeconds;
    private final int maxRequests;
    private final Queue<Long> requestTimestamps;
    
    public SlidingWindowRateLimiter(int windowSizeInSeconds, int maxRequests) {
        this.windowSizeInSeconds = windowSizeInSeconds;
        this.maxRequests = maxRequests;
        this.requestTimestamps = new LinkedList<>();
    }
    
    public synchronized boolean allowRequest() {
        long now = System.currentTimeMillis();
        long windowStart = now - (windowSizeInSeconds * 1000);
        
        // Remove old requests outside the window
        while (!requestTimestamps.isEmpty() && 
               requestTimestamps.peek() < windowStart) {
            requestTimestamps.poll();
        }
        
        // Check if we can accept this request
        if (requestTimestamps.size() < maxRequests) {
            requestTimestamps.offer(now);
            return true;
        }
        
        return false;
    }
}

// Distributed rate limiter using Redis (pseudo-code)
public class DistributedRateLimiter {
    private final JedisPool redisPool;
    private final String keyPrefix;
    
    public DistributedRateLimiter(JedisPool redisPool, String keyPrefix) {
        this.redisPool = redisPool;
        this.keyPrefix = keyPrefix;
    }
    
    public boolean allowRequest(String userId, int limit, int windowSeconds) {
        try (Jedis jedis = redisPool.getResource()) {
            String key = keyPrefix + userId;
            long now = System.currentTimeMillis();
            long windowStart = now - (windowSeconds * 1000);
            
            // Remove old entries
            jedis.zremrangeByScore(key, 0, windowStart);
            
            // Count current requests
            long currentCount = jedis.zcard(key);
            
            if (currentCount < limit) {
                // Add this request
                jedis.zadd(key, now, String.valueOf(now));
                jedis.expire(key, windowSeconds);
                return true;
            }
            
            return false;
        }
    }
}
```

**Task:**
1. Implement token bucket rate limiter
2. Add sliding window rate limiter
3. Design distributed rate limiting with Redis

---

### **Day 17: Monitoring Exercise**
**Exercise: Metrics Collection**
```python
# Python - Simple metrics collection
import time
import threading
from collections import defaultdict, deque
from typing import Dict, List

class MetricsCollector:
    def __init__(self):
        self.counters = defaultdict(int)
        self.gauges = defaultdict(float)
        self.histograms = defaultdict(list)
        self.timers = defaultdict(deque)
        self.lock = threading.Lock()
    
    def increment(self, metric_name: str, value: int = 1):
        with self.lock:
            self.counters[metric_name] += value
    
    def set_gauge(self, metric_name: str, value: float):
        with self.lock:
            self.gauges[metric_name] = value
    
    def record_histogram(self, metric_name: str, value: float):
        with self.lock:
            self.histograms[metric_name].append(value)
            # Keep only last 1000 values
            if len(self.histograms[metric_name]) > 1000:
                self.histograms[metric_name] = self.histograms[metric_name][-1000:]
    
    def record_timer(self, metric_name: str, duration: float):
        with self.lock:
            self.timers[metric_name].append(duration)
            # Keep only last 1000 values
            if len(self.timers[metric_name]) > 1000:
                self.timers[metric_name] = self.timers[metric_name][-1000:]
    
    def get_metrics(self) -> Dict:
        with self.lock:
            metrics = {
                "counters": dict(self.counters),
                "gauges": dict(self.gauges),
                "histograms": {},
                "timers": {}
            }
            
            # Calculate histogram statistics
            for name, values in self.histograms.items():
                if values:
                    metrics["histograms"][name] = {
                        "count": len(values),
                        "sum": sum(values),
                        "avg": sum(values) / len(values),
                        "min": min(values),
                        "max": max(values),
                        "p50": self.percentile(values, 50),
                        "p95": self.percentile(values, 95),
                        "p99": self.percentile(values, 99)
                    }
            
            # Calculate timer statistics
            for name, values in self.timers.items():
                if values:
                    metrics["timers"][name] = {
                        "count": len(values),
                        "avg": sum(values) / len(values),
                        "min": min(values),
                        "max": max(values),
                        "p50": self.percentile(values, 50),
                        "p95": self.percentile(values, 95),
                        "p99": self.percentile(values, 99)
                    }
            
            return metrics
    
    def percentile(self, values: List[float], p: int) -> float:
        if not values:
            return 0
        
        sorted_values = sorted(values)
        index = int((p / 100) * len(sorted_values))
        return sorted_values[min(index, len(sorted_values) - 1)]

# Decorator for timing functions
def timer(metric_name: str):
    def decorator(func):
        def wrapper(*args, **kwargs):
            start_time = time.time()
            try:
                result = func(*args, **kwargs)
                return result
            finally:
                duration = time.time() - start_time
                metrics.record_timer(metric_name, duration)
        return wrapper
    return decorator

# Counter decorator
def counter(metric_name: str):
    def decorator(func):
        def wrapper(*args, **kwargs):
            try:
                result = func(*args, **kwargs)
                metrics.increment(metric_name + "_success")
                return result
            except Exception as e:
                metrics.increment(metric_name + "_error")
                raise
        return wrapper
    return decorator

# Global metrics instance
metrics = MetricsCollector()

# Usage example
@timer("api_request_duration")
@counter("api_requests")
def api_request():
    time.sleep(0.1)  # Simulate work
    return "success"

def test_monitoring():
    # Simulate some API requests
    for i in range(100):
        api_request()
    
    # Get metrics
    current_metrics = metrics.get_metrics()
    
    print("=== Metrics Summary ===")
    print(f"API Requests: {current_metrics['counters']['api_requests_success']}")
    print(f"API Errors: {current_metrics['counters'].get('api_requests_error', 0)}")
    print(f"Average Duration: {current_metrics['timers']['api_request_duration']['avg']:.3f}s")
    print(f"95th Percentile: {current_metrics['timers']['api_request_duration']['p95']:.3f}s")
```

**Task:**
1. Implement metrics collection system
2. Add alerting based on thresholds
3. Create simple dashboard visualization

---

## **Phase 4: Mini Projects (Day 18-23)**

### **Day 18: TinyURL Project**
**Exercise: URL Shortener Implementation**
```go
// Go - URL shortener with different hashing strategies
type URLShortener struct {
    urlMap       map[string]string
    reverseMap   map[string]string
    counter      int64
    mu           sync.RWMutex
    base62Chars  []byte
}

func NewURLShortener() *URLShortener {
    return &URLShortener{
        urlMap:      make(map[string]string),
        reverseMap:  make(map[string]string),
        base62Chars: []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
    }
}

// Base62 encoding
func (us *URLShortener) toBase62(num int64) string {
    if num == 0 {
        return string(us.base62Chars[0])
    }
    
    var encoded []byte
    for num > 0 {
        remainder := num % 62
        encoded = append([]byte{us.base62Chars[remainder]}, encoded...)
        num = num / 62
    }
    
    return string(encoded)
}

// MD5 hash approach
func (us *URLShortener) shortenWithHash(longURL string) string {
    us.mu.Lock()
    defer us.mu.Unlock()
    
    // Check if URL already exists
    if shortCode, exists := us.reverseMap[longURL]; exists {
        return shortCode
    }
    
    // Generate hash
    hash := md5.Sum([]byte(longURL))
    hashStr := hex.EncodeToString(hash[:])
    
    // Use first 6 characters
    shortCode := hashStr[:6]
    
    // Handle collision
    for i := 0; i < 10; i++ {
        if _, exists := us.urlMap[shortCode]; !exists {
            us.urlMap[shortCode] = longURL
            us.reverseMap[longURL] = shortCode
            return shortCode
        }
        // Try different part of hash
        shortCode = hashStr[i*6 : (i+1)*6]
    }
    
    // Fallback to counter
    return us.shortenWithCounter(longURL)
}

// Counter approach
func (us *URLShortener) shortenWithCounter(longURL string) string {
    us.mu.Lock()
    defer us.mu.Unlock()
    
    // Check if URL already exists
    if shortCode, exists := us.reverseMap[longURL]; exists {
        return shortCode
    }
    
    us.counter++
    shortCode := us.toBase62(us.counter)
    
    us.urlMap[shortCode] = longURL
    us.reverseMap[longURL] = shortCode
    
    return shortCode
}

func (us *URLShortener) expandURL(shortCode string) (string, error) {
    us.mu.RLock()
    defer us.mu.RUnlock()
    
    longURL, exists := us.urlMap[shortCode]
    if !exists {
        return "", errors.New("URL not found")
    }
    
    return longURL, nil
}

// Add analytics
type URLAnalytics struct {
    ShortCode    string    `json:"short_code"`
    LongURL      string    `json:"long_url"`
    ClickCount   int64     `json:"click_count"`
    CreatedAt    time.Time `json:"created_at"`
    LastAccessed time.Time `json:"last_accessed"`
}

type AnalyticsTracker struct {
    analytics map[string]*URLAnalytics
    mu        sync.RWMutex
}

func (at *AnalyticsTracker) trackClick(shortCode string) {
    at.mu.Lock()
    defer at.mu.Unlock()
    
    if analytics, exists := at.analytics[shortCode]; exists {
        analytics.ClickCount++
        analytics.LastAccessed = time.Now()
    }
}

func (at *AnalyticsTracker) getAnalytics(shortCode string) *URLAnalytics {
    at.mu.RLock()
    defer at.mu.RUnlock()
    
    return at.analytics[shortCode]
}
```

**Task:**
1. Implement complete URL shortener
2. Add expiration functionality
3. Implement analytics tracking
4. Add rate limiting for abuse prevention

---

### **Day 19: Rate Limiter Project**
**Exercise: Advanced Rate Limiting**
```python
# Python - Distributed rate limiter with Redis
import redis
import time
import json
from typing import Dict, Optional

class RedisRateLimiter:
    def __init__(self, redis_client: redis.Redis):
        self.redis = redis_client
        self.lua_script = """
        local key = KEYS[1]
        local limit = tonumber(ARGV[1])
        local window = tonumber(ARGV[2])
        local current_time = tonumber(ARGV[3])
        
        -- Remove old entries
        redis.call('ZREMRANGEBYSCORE', key, 0, current_time - window * 1000)
        
        -- Count current requests
        local current_requests = redis.call('ZCARD', key)
        
        if current_requests < limit then
            -- Add this request
            redis.call('ZADD', key, current_time, current_time)
            redis.call('EXPIRE', key, window)
            return 1
        else
            return 0
        end
        """
        self.script_sha = self.redis.script_load(self.lua_script)
    
    def is_allowed(self, key: str, limit: int, window_seconds: int) -> bool:
        current_time = int(time.time() * 1000)  # milliseconds
        
        result = self.redis.evalsha(
            self.script_sha,
            1,  # number of keys
            key,
            limit,
            window_seconds,
            current_time
        )
        
        return bool(result)

class MultiTierRateLimiter:
    def __init__(self, redis_client: redis.Redis):
        self.redis = redis_client
        self.limiter = RedisRateLimiter(redis_client)
        
        # Define rate limits for different tiers
        self.tiers = {
            "free": {
                "requests_per_minute": 10,
                "requests_per_hour": 100,
                "requests_per_day": 1000
            },
            "premium": {
                "requests_per_minute": 100,
                "requests_per_hour": 1000,
                "requests_per_day": 10000
            },
            "enterprise": {
                "requests_per_minute": 1000,
                "requests_per_hour": 10000,
                "requests_per_day": 100000
            }
        }
    
    def check_limits(self, user_id: str, tier: str) -> Dict[str, bool]:
        limits = self.tiers.get(tier, self.tiers["free"])
        
        results = {}
        
        # Check per-minute limit
        minute_key = f"rate_limit:{user_id}:minute"
        results["minute"] = self.limiter.is_allowed(
            minute_key, 
            limits["requests_per_minute"], 
            60
        )
        
        # Check per-hour limit
        hour_key = f"rate_limit:{user_id}:hour"
        results["hour"] = self.limiter.is_allowed(
            hour_key, 
            limits["requests_per_hour"], 
            3600
        )
        
        # Check per-day limit
        day_key = f"rate_limit:{user_id}:day"
        results["day"] = self.limiter.is_allowed(
            day_key, 
            limits["requests_per_day"], 
            86400
        )
        
        return results
    
    def is_allowed(self, user_id: str, tier: str) -> bool:
        results = self.check_limits(user_id, tier)
        
        # User is allowed if all limits are not exceeded
        return all(results.values())

class AdaptiveRateLimiter:
    def __init__(self, redis_client: redis.Redis):
        self.redis = redis_client
        self.base_limiter = RedisRateLimiter(redis_client)
        self.system_load_key = "system_load"
    
    def get_system_load(self) -> float:
        # Simple system load calculation
        # In production, this would use actual system metrics
        try:
            current_requests = self.redis.get("current_requests_per_second") or b"0"
            load = float(current_requests) / 1000.0  # Normalize to 0-1
            return min(load, 1.0)
        except:
            return 0.0
    
    def adjust_limit(self, base_limit: int) -> int:
        load = self.get_system_load()
        
        # Reduce limit based on system load
        if load > 0.9:
            return int(base_limit * 0.1)  # Reduce to 10%
        elif load > 0.7:
            return int(base_limit * 0.5)  # Reduce to 50%
        elif load > 0.5:
            return int(base_limit * 0.8)  # Reduce to 80%
        else:
            return base_limit  # No reduction
    
    def is_allowed_adaptive(self, key: str, base_limit: int, window_seconds: int) -> bool:
        adjusted_limit = self.adjust_limit(base_limit)
        return self.base_limiter.is_allowed(key, adjusted_limit, window_seconds)

# Usage example
def setup_rate_limiter():
    redis_client = redis.Redis(host='localhost', port=6379, db=0)
    
    # Multi-tier rate limiter
    multi_tier = MultiTierRateLimiter(redis_client)
    
    # Test different user tiers
    free_user = "user123"
    premium_user = "user456"
    
    print("=== Rate Limiting Test ===")
    
    # Test free user (should be limited)
    for i in range(15):  # Exceeds free tier limit
        allowed = multi_tier.is_allowed(free_user, "free")
        print(f"Free user request {i+1}: {'Allowed' if allowed else 'Blocked'}")
    
    # Test premium user
    for i in range(15):
        allowed = multi_tier.is_allowed(premium_user, "premium")
        print(f"Premium user request {i+1}: {'Allowed' if allowed else 'Blocked'}")
    
    # Adaptive rate limiting
    adaptive = AdaptiveRateLimiter(redis_client)
    
    for i in range(10):
        allowed = adaptive.is_allowed_adaptive("adaptive_test", 100, 60)
        print(f"Adaptive request {i+1}: {'Allowed' if allowed else 'Blocked'}")
```

**Task:**
1. Implement distributed rate limiting with Redis
2. Add multi-tier rate limiting
3. Implement adaptive rate limiting based on system load
4. Add rate limiting analytics

---

## **Continue with remaining days...**

*This is a comprehensive set of coding exercises for the first 19 days. The remaining days (20-30) would follow similar patterns with increasingly complex implementations for the advanced topics and interview projects.*

---

## **How to Use These Exercises**

### **Daily Practice:**
1. **Read the theory** first from the referenced materials
2. **Implement the exercise** in your preferred language
3. **Test with different scenarios**
4. **Document trade-offs** and design decisions
5. **Practice explaining** the solution out loud

### **Language Flexibility:**
- All exercises are provided in multiple languages (Go, Java, Python)
- Choose the language you're most comfortable with
- Focus on understanding concepts, not just syntax

### **Progress Tracking:**
- Complete each exercise before moving to the next day
- Keep a coding journal of what you learned
- Review previous exercises weekly

### **Interview Preparation:**
- Practice explaining your implementations
- Be ready to discuss trade-offs
- Prepare to whiteboard these solutions

---

**Remember:** The goal is to understand the concepts behind each exercise, not just copy the code. Focus on the "why" behind each design decision!
