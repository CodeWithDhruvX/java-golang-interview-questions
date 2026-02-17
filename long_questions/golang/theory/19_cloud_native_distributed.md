# 🟢 Go Theory Questions: 361–380 Cloud-Native and Distributed Systems in Go

## 361. How do you build a cloud-agnostic app in Go?

**Answer:**
We follow the **Hexagonal Architecture** (Code to Interfaces).

We define a `BlobStorage` interface with methods like `Upload` and `Download`.
We implement `S3Storage` (AWS) and `GCSStorage` (Google Cloud).
In `main.go`, we check an env var `CLOUD_PROVIDER` to decide which implementation to inject.

We also avoid proprietary services where standard ones exist (e.g., use Postgres instead of DynamoDB if portability is key). This ensures our core business logic knows nothing about Amazon or Google, only about "Storage."

---

## 362. How do you use Go SDKs with AWS (S3, Lambda)?

**Answer:**
We use the official `aws-sdk-go-v2`.

It is modular. We don't import the whole world; we import `service/s3` or `service/dynamodb`.
Key pattern: **Configuration**.
`cfg, err := config.LoadDefaultConfig(context.TODO())`.
This automatically looks for credentials in Env Vars, `~/.aws/credentials`, or IAM Roles (if running on EC2/Lambda). We never hardcode keys.

---

## 363. How do you upload a file to S3 using Go?

**Answer:**
We use the `s3.Client.PutObject` method.

For large files, we use the **S3 Manager** (`manager.NewUploader`).
It automatically splits the file into pieces and uploads them in parallel (Multipart Upload).
This is crucial for reliability. If a 1GB upload fails at 99%, a simple `PutObject` fails completely. The Uploader allows resuming or just retrying the failed chunks.

---

## 364. How do you create a Pub/Sub system using Go and GCP?

**Answer:**
We use the Google Cloud Pub/Sub client library.

**Publisher**: `topic.Publish(ctx, &Message{Data: []byte("foo")})`.
**Subscriber**: `sub.Receive(ctx, func(ctx, msg) { ... msg.Ack() })`.

Important design interaction: GCP Pub/Sub pushes messages to your callback concurrently. You don't write a loop; you just handle the callback. If you need to limit concurrency (to protect your DB), you must configure `ReceiveSettings.MaxOutstandingMessages` or use a local semaphore.

---

## 365. How would you implement cloud-native config loading?

**Answer:**
In Kubernetes, config is often a **ConfigMap** mounted as a file.

However, for dynamic reloading, we can use a "Watcher".
We use `fsnotify` to watch the config file. When K8s updates the ConfigMap, the file changes inside the container.
Our Go app detects the `Write` event, re-reads the JSON/YAML, and updates the global config struct safely (using a `RWMutex` to prevent races with readers). This allows "Hot Reloading" without restarting the pod.

---

## 366. What is the role of service meshes with Go apps?

**Answer:**
A Service Mesh (Istio/Linkerd) moves network logic **out of the application** and into a sidecar proxy (Envoy).

Without a mesh, our Go app needs code for Circuit Breaking, Retries, mTLS, and Tracing.
With a mesh, our Go app just makes a simple HTTP call to `http://billing-service`. The sidecar intercepts it, adds mTLS, handles retries, and emits metrics. This keeps our Go code "Business Logic Only," shrinking the codebase significantly.

---

## 367. How do you secure service-to-service communication in Go?

**Answer:**
If not using a Service Mesh, we must implement **mTLS** (Mutual TLS) manually.

Server: `Use `tls.RequireAndVerifyClientCert`.
Client: Load a Client Certificate signed by the internal CA.

Additionally, we use **JWTs** (Service Accounts). Service A signs a JWT with its private key (or gets one from the IDP/OIDC provider). Service B validates the JWT to verify "This request really came from Service A" before allowing access to the data.

---

## 368. How do you implement service registration and discovery?

**Answer:**
In modern K8s, we rely on **DNS**. We call `http://my-service`. K8s DNS resolves this to the ClusterIP.

In legacy or bare-metal setups, we use **Consul** or **Etcd**.
On startup, the Go app registers itself: "I am Service A, IP 10.0.0.1, Port 8080." and sends a heartbeat (TTL).
Clients query Consul: "Give me healthy IPs for Service A." Client-side load balancing (in gRPC) then picks one IP to verify connectivity.

---

## 369. How do you manage retries and circuit breakers in Go?

**Answer:**
We use middleware logic.

**Retry**: For GET requests (idempotent), we retry 3 times with backoff. We *never* retry POSTs blindly (could duplicate payments).
**Circuit Breaker**: Use libraries like `gobreaker`. We wrap the `http.Client`.
If the error rate > 50%, the breaker trips (Open). Future calls return "Circuit Open" error immediately without hitting the network. This protects the downstream service from being hammered while it's trying to recover.

---

## 370. How would you use etcd/Consul with Go for KV storage?

**Answer:**
Etcd is a strongly consistent Distributed Key-Value store (used by K8s itself).

We use the `clientv3` library.
Use cases: **Dynamic Config** and **Flag Toggles**.
We store `flags/new-ui-enabled = true`.
All Go services watch this key (`client.Watch`). When the value changes, they get an event notification and flip the feature flag instantly across the entire fleet without a redeploy.

---

## 371. What is leader election and how can you implement it in Go?

**Answer:**
Leader Election ensures only *one* instance of a service performs a task (e.g., a Cron Job scheduling emails) to avoid duplication.

We use **Kubernetes Leases** (`client-go/tools/leaderelection`) or a Redis/Etcd lock.
Concept: Everyone tries to create a key `lock:email-scheduler` with a TTL (Time To Live).
Only one succeeds. That one becomes Leader. It must continuously "renew" the lease (heartbeat). If it dies, the TTL expires, and another instance grabs the lock.

---

## 372. How do you build a distributed lock in Go?

**Answer:**
The standard algorithm is **Redlock** (Redis) or utilizing Etcd's strong consistency.

With Redis: `SET resource_name my_random_id NX PX 30000`.
NX = Only set if not exists. PX = Expires in 30s.

Critical part: **Releasing**. You must check "Is the value still `my_random_id`?" before deleting. This prevents you from deleting a lock that *expired* and was grabbed by someone else. We use a Lua script to make this Check-And-Delete atomic.

---

## 373. How would you implement a distributed queue in Go?

**Answer:**
Don't build one; use Kafka/SQS. But if asked to *design* one:

We need durability and ordering.
Backend: Append-only log files (sharded by partition).
Metadata: Zookeeper/Etcd to track "Who is consuming Partition 1?".
Go: High-performance TCP server. Uses `sendfile` syscall to stream log data from disk to network socket directly (Zero Copy), exactly how Kafka does it.

---

## 374. How do you handle consistency in distributed Go systems?

**Answer:**
We must choose: **Strong Consistency** (CP) vs **Eventual Consistency** (AP).

For Strong (Banking), we use Distributed Transactions (Two-Phase Commit) or stick to a single SQL DB.
For Eventual (Social Feed), we allow replicas to lag.
In Go, we handle this by designing idempotent consumers. If data arrives late or out of order, the system self-corrects. "Last Write Wins" is a common strategy, handled by passing timestamps with every update.

---

## 375. How do you monitor and trace distributed Go systems?

**Answer:**
**Distributed Tracing** (OpenTelemetry) is mandatory.

Every request gets a `TraceID` at the Ingress.
This ID is passed in HTTP Headers (`Traceparent`) or gRPC Metadata.
Every Go service extracts it, creates a new `Span` (child), does work, and closes the span.
We visualize this in Jaeger. We can see: "Total 500ms. 50ms in Handler, 20ms in Auth Service, 430ms in Database Query." This instantly identifies the bottleneck.

---

## 376. How do you implement eventual consistency in Go?

**Answer:**
We use the **Outbox Pattern** or **Saga Pattern**.

When a user buys an item:
1.  Update DB: `INSERT INTO orders ...` (Local Transaction)
2.  Publish Event: `kafka.Publish("OrderCreated")`

If step 2 fails, we are inconsistent.
Fix: Write the event to a `outbox` table in the *same* DB transaction as the order.
A background Go worker polls the `outbox` table and pushes to Kafka. If it fails, it retries. This guarantees "At Least Once" delivery, ensuring the downstream eventually gets the data.

---

## 377. How do you replicate state in distributed Go apps?

**Answer:**
We use **Raft** or **Paxos** consensus algorithms.

Hashicorp's `raft` library (written in Go) is the standard.
Each node has a Finite State Machine (FSM).
When a write comes in (Leader), it appends to its log and replicates to Followers. Once a Quorum (Majority) confirms, the write is committed and applied to the FSM.
This is how Etcd, Consul, and CockroachDB work internally.

---

## 378. How do you detect and handle split-brain scenarios?

**Answer:**
Split-brain happens when a network partition makes two clusters think *they* are the primary.

Mitigation: **Quorum**.
You need `N/2 + 1` nodes to accept a write.
If you have 5 nodes, you need 3. If the network splits into 2 and 3, the side with 2 cannot accept writes (becomes Read-Only). The side with 3 continues.
In Go code (using Raft), this logic is built-in. The leader steps down if it loses contact with the quorum.

---

## 379. How do you implement quorum reads/writes in Go?

**Answer:**
This is often manual in NoSQL (Cassandra/Dynamo).

**Write Quorum (W)**: Wait for ACKs from W nodes.
**Read Quorum (R)**: Read from R nodes, return the latest timestamp.
Rule: `R + W > N`. This guarantees strict consistency (the read set and write set must overlap by at least one node).

In Go, we use `sync.WaitGroup` to launch N parallel requests, wait for W success responses, and then return success to the client immediately (canceling the others).

---

## 380. How would you build a simple distributed cache in Go?

**Answer:**
1.  **Sharding**: Use Consistent Hashing to decide which node holds key "A".
2.  **Communication**: Each node knows the ring members (gossip protocol like `memberlist`).
3.  **Storage**: `sync.Map` or an LRU library in memory.

Client asks Node 1 for "Key A".
Node 1 hashes "Key A" -> Belongs to Node 3.
Node 1 forwards request to Node 3 (or tells Client "Redirect to Node 3").
This is essentially how **Groupcache** (written by Go team) works.
