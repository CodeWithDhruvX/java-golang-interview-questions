# High-Performance Order Processor (Persistent Systems Demo)

This project demonstrates a production-grade **Golang Microservice** architecture tailored for the Persistent Systems interview.

## **Key Features Implemented**
1.  **Concurrency Patterns**: Uses a **Worker Pool** of Goroutines consuming from a **Buffered Channel**.
2.  **Backpressure Handling**: Returns `503 Service Unavailable` immediately if the queue is full, preventing system overload.
3.  **Graceful Shutdown**: Uses `context.WithTimeout` and `sync.WaitGroup` to ensure running jobs finish before the server exits.
4.  **Observability**:
    *   Exposes Prometheus-compatible metrics at `/metrics`.
    *   Tracks `orders_processed` (Counter), `orders_failed` (Counter), and `queue_depth` (Gauge).

## **Additional Components (Persistent Systems Add-ons)**

### **1. Kubernetes Deployment**
Production-ready manifests are in the `k8s/` directory.
```bash
kubectl apply -f k8s/deploy.yaml
```

### **2. gRPC Interface (Proto)**
The protocol buffer definition is in `proto/order.proto`.
The server implementation logic is demonstrated in `grpc_server_example.go`.
*Note: Requires `protoc` to regenerate code.*

### **3. BDD Testing (Ginkgo)**
Behavior Driven Development tests are in `main_test.go`.
To run them:
```bash
go test -v ./...
```
*Note: I have already set up the dependencies. You might need to run `go mod tidy` if you make changes.*

## **How to Run**

### **1. Local Go**
```bash
go run main.go
```

### **2. Docker**
```bash
docker build -t order-processor .
docker run -p 8080:8080 order-processor
```

## **Testing the Endpoints**

**1. Submit an Order:**
```bash
# PowerShell
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/submit" -Body '{"id": "ord-001", "value": 100.50}' -ContentType "application/json"

# Curl
curl -X POST -d '{"id": "ord-001", "value": 100.50}' http://localhost:8080/submit
```

**2. Check Prometheus Metrics:**
Open [http://localhost:8080/metrics](http://localhost:8080/metrics) in your browser.
You will see:
```text
# HELP orders_processed Total valid orders processed
# TYPE orders_processed counter
orders_processed 1
...
queue_depth 0
```

**3. Test Concurrency & Backpressure:**
Run a load test (e.g., using `hey` or a simple loop) to fill the queue and see the `queue_depth` metric rise and eventual 503 errors.
