# 📡 06 — Microservices & gRPC in Go
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- gRPC basics and protobuf
- Unary vs streaming RPC
- gRPC interceptors (middleware)
- Service-to-service communication
- Message queues (Kafka/NATS)
- Service discovery and health checks

---

## ❓ Most Asked Questions

### Q1. What is gRPC and how does it differ from REST?

| | gRPC | REST |
|--|------|------|
| Protocol | HTTP/2 | HTTP/1.1 or HTTP/2 |
| Data format | Protobuf (binary) | JSON (text) |
| Contract | Strict `.proto` schema | Loose (OpenAPI optional) |
| Streaming | Bidirectional | Limited (SSE, WebSockets) |
| Performance | 2-10x faster | Slower |
| Browser support | Limited (gRPC-Web) | Native |
| Code gen | Yes (automatic) | Manual |

---

### Q2. How do you define and use protobuf in Go?

```protobuf
// user.proto
syntax = "proto3";
package user;
option go_package = "./userpb";

message User {
    int32 id = 1;
    string name = 2;
    string email = 3;
}

message GetUserRequest  { int32 id = 1; }
message GetUserResponse { User user = 1; }
message CreateUserRequest {
    string name = 1;
    string email = 2;
}

service UserService {
    rpc GetUser(GetUserRequest) returns (GetUserResponse);
    rpc CreateUser(CreateUserRequest) returns (User);
    rpc ListUsers(stream GetUserRequest) returns (stream User);
}
```

```bash
# Generate Go code from proto
protoc --go_out=. --go-grpc_out=. user.proto
```

---

### Q3. How do you implement a gRPC server and client in Go?

```go
// server.go
type userServer struct {
    userpb.UnimplementedUserServiceServer
    db map[int32]*userpb.User
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
    user, ok := s.db[req.Id]
    if !ok {
        return nil, status.Errorf(codes.NotFound, "user %d not found", req.Id)
    }
    return &userpb.GetUserResponse{User: user}, nil
}

func main() {
    lis, _ := net.Listen("tcp", ":50051")
    srv := grpc.NewServer()
    userpb.RegisterUserServiceServer(srv, &userServer{db: make(map[int32]*userpb.User)})
    reflection.Register(srv)  // enables grpcurl
    srv.Serve(lis)
}

// client.go
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
defer conn.Close()
client := userpb.NewUserServiceClient(conn)

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
resp, err := client.GetUser(ctx, &userpb.GetUserRequest{Id: 1})
```

---

### Q4. How do you implement gRPC interceptors (middleware)?

```go
// Unary interceptor — like middleware for individual calls
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    start := time.Now()
    resp, err := handler(ctx, req)
    log.Printf("RPC: %s | Duration: %v | Error: %v", info.FullMethod, time.Since(start), err)
    return resp, err
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok { return nil, status.Error(codes.Unauthenticated, "missing metadata") }
    tokens := md.Get("authorization")
    if len(tokens) == 0 { return nil, status.Error(codes.Unauthenticated, "missing token") }
    // validate token...
    return handler(ctx, req)
}

// Chain multiple interceptors
srv := grpc.NewServer(
    grpc.ChainUnaryInterceptor(authInterceptor, loggingInterceptor),
)
```

---

### Q5. How does server-side streaming work in gRPC?

```go
// proto: rpc ListUsers(Empty) returns (stream User);

// Server implementation
func (s *userServer) ListUsers(req *emptypb.Empty, stream userpb.UserService_ListUsersServer) error {
    users := s.getAllUsers()
    for _, user := range users {
        // Check if client cancelled
        if err := stream.Context().Err(); err != nil { return err }
        if err := stream.Send(user); err != nil { return err }
        time.Sleep(100 * time.Millisecond)  // simulate slow data
    }
    return nil
}

// Client implementation
stream, err := client.ListUsers(ctx, &emptypb.Empty{})
if err != nil { log.Fatal(err) }
for {
    user, err := stream.Recv()
    if err == io.EOF { break }  // stream ended
    if err != nil { log.Fatal(err) }
    fmt.Printf("Received: %s\n", user.Name)
}
```

---

### Q6. How do you integrate Kafka with Go microservices?

```go
import "github.com/IBM/sarama"

// Producer
func newKafkaProducer(brokers []string) (sarama.SyncProducer, error) {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true
    return sarama.NewSyncProducer(brokers, config)
}

func publishEvent(producer sarama.SyncProducer, topic string, key, value []byte) error {
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Key:   sarama.ByteEncoder(key),
        Value: sarama.ByteEncoder(value),
    }
    _, _, err := producer.SendMessage(msg)
    return err
}

// Consumer
func consumeEvents(brokers []string, topic, groupID string, handler func([]byte)) {
    config := sarama.NewConfig()
    config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

    consumer, _ := sarama.NewConsumerGroup(brokers, groupID, config)
    handler := &ConsumerGroupHandler{msgHandler: handler}
    for {
        consumer.Consume(context.Background(), []string{topic}, handler)
    }
}
```

---

### Q7. How do you implement service health checks?

```go
// HTTP health check endpoints
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func readinessHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := db.Ping(); err != nil {
            http.Error(w, "DB not ready", http.StatusServiceUnavailable)
            return
        }
        w.WriteHeader(http.StatusOK)
    }
}

// gRPC health check (standard grpc_health_v1)
import "google.golang.org/grpc/health/grpc_health_v1"
healthSrv := health.NewServer()
grpc_health_v1.RegisterHealthServer(grpcServer, healthSrv)
healthSrv.SetServingStatus("UserService", grpc_health_v1.HealthCheckResponse_SERVING)
```
