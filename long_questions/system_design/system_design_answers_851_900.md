## 🔸 Graphs, Relationships & Social Features (Questions 851-860)

### Question 851: Design a system to recommend mutual connections.

**Answer:**
*   **Graph Model:** `User A -> Follows -> User B` directed edges.
*   **Algorithm:** Triangle Counting - If A follows B, and C follows B, suggest A to C.
*   **Implementation:** 
    *   Use adjacency lists in graph database (Neo4j/Cassandra)
    *   Query: `GetFriendsOfFriends(UserA)` 
    *   Rank by number of mutual friends and interaction frequency
*   **Scoring:** Weight = (mutual_friends × 0.6) + (interaction_strength × 0.4)
*   **Storage:** Redis for quick lookups, batch processing for recommendations

### Question 852: Build a social graph service with depth-based queries.

**Answer:**
*   **Database Choice:** Neo4j for native graph storage or Dgraph for distributed scale.
*   **Query Example:** "Find all users within 3 hops of Alice who work at Google."
*   **Architecture:**
    *   **Partitioning:** Shard by UserID hash to distribute load
    *   **Storage:** Adjacency Lists in RocksDB for fast neighbor traversal
    *   **Index:** Secondary indexes on attributes (company, location)
*   **Performance:** 
    *   Cache frequently accessed subgraphs in Redis
    *   Use bidirectional indices for reverse lookups
    *   Implement query timeouts to prevent deep traversal abuse

### Question 853: How to detect influencer clusters in a network?

**Answer:**
*   **Algorithms:** 
    *   **PageRank:** Identify influential nodes based on incoming link quality
    *   **Community Detection:** Louvain Method for modularity optimization
    *   **Betweenness Centrality:** Find bridge nodes between communities
*   **Cluster Definition:** High density of internal edges, low external connectivity
*   **Implementation:**
    *   Run nightly batch jobs on graph snapshots
    *   Use Spark GraphX for distributed processing
    *   Store cluster assignments in Cassandra
*   **Applications:** Targeted marketing, content seeding, viral campaign optimization

### Question 854: Design a system for follow/unfollow with eventual consistency.

**Answer:**
*   **Write Path:** Append `FollowEvent` to Kafka topic with UserID, TargetID, timestamp, action.
*   **Read Path Optimization:**
    *   **FollowingCount:** Increment counter in Redis for immediate UI updates
    *   **FollowersList:** Append to Cassandra row partitioned by TargetID
    *   **Feed Generation:** Fan-out on write for active users, pull-on-read for inactive
*   **Consistency Model:**
    *   Strong consistency for follow/unfollow actions (Kafka guarantees)
    *   Eventual consistency for follower counts (periodic reconciliation)
*   **Conflict Resolution:** Use event timestamps and vector clocks for ordering
*   **Scaling:** Shard by both follower and followee for balanced load

### Question 855: Build a friend suggestion engine using graph traversal.

**Answer:**
*   **Core Algorithm:** (See Q851) Mutual connections detection
*   **Advanced Techniques:**
    *   **Random Walk with Restart:** Simulate random traversal, return to source node with probability p
    *   **Personalized PageRank:** Bias random walks towards user's interests
    *   **Edge Weighting:** Weight by interaction frequency, recency, and response rates
*   **Scoring Pipeline:**
    ```python
    score = mutual_friends * 0.4 + 
            interaction_strength * 0.3 + 
            profile_similarity * 0.2 + 
            network_proximity * 0.1
    ```
*   **Performance:** Pre-compute for active users, real-time for new users
*   **Storage:** Materialized view updated daily, served from Redis

### Question 856: Design a graph-based spam detection platform.

**Answer:**
*   **Spam Patterns:** 
    *   **Tight Cliques:** Groups of accounts that only follow each other
    *   **Star Pattern:** One bot follows thousands of users, minimal followers back
    *   **Chain Reaction:** Sequential account creation with similar naming patterns
*   **Key Features:**
    *   `OutDegree >> InDegree` ratio > 100:1
    *   `AccountAge < 24 hours` with high activity
    *   `Content Similarity` > 90% across posts
*   **Detection Pipeline:**
    *   Real-time graph analysis using Neo4j
    *   Machine learning model on graph features
    *   Human review for borderline cases
*   **Actions:** Shadowban, account suspension, network-wide flagging

### Question 857: Build a group and subgroup system with scoped permissions.

**Answer:**
*   **Data Structure:** DAG (Directed Acyclic Graph) representing group hierarchy
*   **Permission Model:**
    *   `User IN Group A` AND `Group A IS_CHILD_OF Group B`
    *   Inherited permissions flow down the hierarchy
    *   Explicit denies override allows (deny priority)
*   **Implementation:**
    *   **Recursive CTE** in PostgreSQL for permission traversal
    *   **Materialized Path** for fast ancestry checks (`/1/4/7`)
    *   **Closure Table** for transitive relationships
*   **Caching Strategy:**
    *   Cache user permissions in Redis with TTL
    *   Invalidate on group membership changes
    *   Pre-compute common permission sets

### Question 858: Design a real-time "who viewed your profile" system.

**Answer:**
*   **Event Stream:** `ViewEvent(ViewerID, TargetID, Timestamp, Context)`
*   **Processing Pipeline:**
    *   **Deduplication:** One view per viewer-target pair per 24 hours
    *   **Privacy Rules:** Exclude blocked users, private profiles
    *   **Ranking:** Prioritize recent views, mutual connections, profile completeness
*   **Storage Design:**
    *   **Redis List:** `ProfileViews:{TargetID}` -> Circular buffer of last 50 viewers
    *   **TTL:** 30 days for GDPR compliance
    *   **Backup:** Persistent storage in Cassandra for analytics
*   **API:** `GET /profile/{id}/views?limit=20&offset=0`
*   **Performance:** Sub-50ms response, 99.9% availability

### Question 859: Build a common connections insight engine.

**Answer:**
*   **Core Algorithm:** Set intersection `FriendsA ∩ FriendsB`
*   **Optimization Strategies:**
    *   **Bloom Filter:** Send compressed BloomFilter(FriendsA) to User B's shard
    *   **MinHash:** Estimate Jaccard similarity for large sets
    *   **Partitioning:** Shard by user ID for parallel processing
*   **Implementation Details:**
    ```python
    # Bloom filter approach
    filter_b = create_bloom_filter(friends_b)
    common = [f for f in friends_a if filter_b.might_contain(f)]
    # Verify actual intersection
    ```
*   **Caching:** Pre-compute for celebrity/high-profile accounts
*   **API Response:** Mutual count, connection details, introduction suggestions

### Question 860: Design a relationship recommendation system using vector similarity.

**Answer:**
*   **Graph Embedding:** Node2Vec algorithm converts graph nodes to dense vectors
    *   **Parameters:** p=1.0 (return parameter), q=0.5 (in-out parameter)
    *   **Dimensions:** 128-256 dimensions for balance of quality and performance
*   **Similarity Calculation:** Cosine similarity between user vectors
    ```python
    similarity = cosine_similarity(user_a_vector, user_b_vector)
    ```
*   **Advantages over Graph Traversal:**
    *   Captures implicit structural similarity (similar roles, communities)
    *   Works for cold-start users with few connections
    *   Enables fast approximate nearest neighbor search
*   **Architecture:**
    *   **Training:** Weekly batch job on GPU cluster
    *   **Serving:** FAISS index for real-time similarity search
    *   **Hybrid:** Combine with collaborative filtering for better results

---

## 🔸 Automation, Agents & Worker Systems (Questions 861-870)

### Question 861: Design a cron-as-a-service platform.

**Answer:**
*   **API Design:** `POST /schedule { cron: "0 * * * *", url: "...", headers: {...} }`
*   **Core Architecture:**
    *   **Min-Heap:** Store jobs ordered by `NextRunTime` for O(1) retrieval
    *   **Dispatcher Service:** Continuously polls heap, fires webhooks when `Time >= NextRunTime`
    *   **Persistence:** PostgreSQL for job definitions, Redis for runtime state
*   **Scaling Strategy:**
    *   **Partitioning:** Shard jobs by execution time windows
    *   **Workers:** Multiple dispatcher instances with leader election
    *   **Backpressure:** Rate limit webhook firing to prevent target overload
*   **Features:**
    *   **Retry Logic:** Exponential backoff with jitter
    *   **Monitoring:** Success/failure metrics, alert on consecutive failures
    *   **Timezone Support:** Convert all times to UTC for consistency

### Question 862: Build a task queue system with retry, delay, and dependencies.

**Answer:**
*   **Core Design:** (See Q562) Enhanced with dependencies
*   **Delayed Execution:**
    *   **Redis ZSET:** `Score = ExecuteAt`, `Member = TaskID`
    *   **Poller:** Scan for tasks where `Score <= Now`
*   **Dependency Management:**
    *   **DAG Representation:** Tasks as nodes, dependencies as edges
    *   **Topological Sort:** Determine execution order
    *   **Blocking Queue:** Wait for all parents to complete before scheduling child
*   **Retry Strategy:**
    *   **Exponential Backoff:** `delay = base_delay * 2^attempt + random_jitter`
    *   **Dead Letter Queue:** Move tasks after max retries
    *   **Circuit Breaker:** Stop retrying failing task types
*   **Monitoring:** Real-time metrics on queue depth, processing latency, error rates

### Question 863: Design a system for prioritizing queued tasks by SLA.

**Answer:**
*   **Queue Hierarchy:**
    *   **Gold Queue:** Critical tasks (payment processing, emergency alerts)
    *   **Silver Queue:** Important tasks (email notifications, reports)
    *   **Bronze Queue:** Background tasks (analytics, cleanup)
*   **Resource Allocation:**
    *   **Dynamic:** 50% capacity Gold, 30% Silver, 20% Bronze
    *   **Auto-scaling:** If Gold queue depth > threshold, steal from Bronze
    *   **Starvation Prevention:** Process lower queues if higher queues empty
*   **SLA Monitoring:**
    *   **Metrics:** Time-in-queue, processing time, success rate
    *   **Alerting:** SLA breach warnings, queue depth alerts
    *   **Auto-escalation:** Promote tasks approaching SLA deadline
*   **Implementation:**
    *   **Priority Queues:** Redis streams with consumer groups
    *   **Worker Pools:** Separate pools per priority with configurable sizes

### Question 864: Build a distributed worker pool manager.

**Answer:**
*   **Heartbeat Mechanism:**
    *   **Frequency:** Worker pings Manager every 5 seconds with status
    *   **Payload:** `WorkerID, CPU_Load, Memory_Usage, Active_Tasks`
    *   **Timeout Detection:** Mark worker as failed after 3 missed heartbeats
*   **Task Assignment Strategy:**
    *   **Push Model:** Manager pushes `TaskID` to `Worker:{ID}:Queue`
    *   **Load Balancing:** Assign to least loaded worker with required capabilities
    *   **Affinity:** Prefer same datacenter/region for data locality
*   **Failure Recovery:**
    *   **Task Reassignment:** Move in-progress tasks to healthy workers
    *   **Checkpointing:** Workers save progress periodically for resume capability
    *   **Circuit Breaking:** Temporarily exclude repeatedly failing workers
*   **Scaling:** Auto-scale workers based on queue depth and processing rate

### Question 865: Design a workflow engine with failure recovery.

**Answer:**
*   **Framework Choice:** (See Q561) Cadence / Temporal for durability
*   **Event Sourcing Architecture:**
    *   **Event History:** Immutable log of all workflow events (start, complete, fail)
    *   **Replay Capability:** Resume exactly where crashed by replaying history
    *   **State Snapshots:** Periodic checkpoints to speed up replay
*   **Failure Handling:**
    *   **Automatic Retry:** Configurable retry policies per activity type
    *   **Manual Intervention:** Human tasks for exception handling
    *   **Compensation:** Undo operations for saga pattern
*   **Implementation Details:**
    *   **Workflow Definition:** JSON/YAML describing steps and dependencies
    *   **Activity Workers:** Stateless services executing individual tasks
    *   **Decision Points:** Conditional logic based on previous outcomes
*   **Monitoring:** Real-time workflow status, execution metrics, SLA tracking

### Question 866: Build a job monitoring and alerting platform.

**Answer:**
*   **Agent Architecture:**
    *   **Wrapper Library:** Wraps job execution with instrumentation
    *   **Event Emission:** `Start`, `Progress`, `End`, `Error` events with timestamps
    *   **Metadata Collection:** Job type, resource usage, custom metrics
*   **Monitoring Engine:**
    ```python
    class JobMonitor:
        def check_job_health(self, job_id):
            job = self.get_job(job_id)
            duration = now() - job.start_time
            
            if duration > job.expected_duration * 2:
                self.alert("Stuck Job", job_id, duration)
            elif job.status == "ERROR":
                self.alert("Job Failed", job_id, job.error_message)
            
            # Resource usage alerts
            if job.cpu_usage > 90%:
                self.alert("High CPU", job_id, job.cpu_usage)
    ```
*   **Alerting System:**
    *   **Channels:** Email, Slack, PagerDuty, webhook
    *   **Escalation:** Tiered alerting based on severity and duration
    *   **Suppression:** Alert grouping to prevent spam
*   **Dashboard Features:**
    *   **Real-time View:** Active jobs with progress bars
    *   **Historical Analysis:** Job success rates, performance trends
    *   **Resource Utilization:** Cluster-wide resource monitoring

### Question 867: Design a dynamic workload auto-scaling system.

**Answer:**
*   **Predictive Analytics:**
    *   **Time Series Models:** ARIMA, Prophet, LSTM for load forecasting
    *   **Feature Engineering:** Hour of day, day of week, seasonality, events
    *   **Model Training:** Retrain weekly on historical usage patterns
*   **Pre-scaling Strategy:**
    ```python
    def predict_and_scale():
        # Predict load for next hour
        predicted_load = arima_model.predict(horizon=60)
        
        # Calculate required resources
        required_nodes = ceil(predicted_load / node_capacity)
        current_nodes = get_active_node_count()
        
        # Scale 10 minutes before predicted spike
        if predicted_load > current_threshold:
            scale_up(required_nodes - current_nodes, delay=600)
    ```
*   **Reactive Scaling:**
    *   **Metrics Monitoring:** CPU, memory, queue depth, response time
    *   **Scaling Policies:** Target utilization 70%, scale up at 80%, scale down at 30%
    *   **Cooldown Periods:** Prevent thrashing (5 min up, 15 min down)
*   **Optimization Features:**
    *   **Cost Awareness:** Prefer spot instances for non-critical workloads
    *   **Resource Rightsizing:** Adjust instance types based on actual usage
    *   **Multi-dimensional Scaling:** Scale CPU and memory independently

### Question 868: Build an orchestrator for data processing pipelines.

**Answer:**
*   **Pipeline Definition:**
    *   **DAG Specification:** YAML/JSON defining tasks and dependencies
    *   **Task Types:** Extract, Transform, Load, ML training, validation
    *   **Dependency Management:** Explicit dependencies and conditional execution
*   **Execution Engine:**
    ```yaml
    # Example pipeline definition
    pipeline:
      name: "daily_etl"
      schedule: "0 2 * * *"
      tasks:
        - name: extract_data
          type: kubernetes_job
          image: etl/extractor:latest
          resources: {cpu: 2, memory: 4Gi}
        - name: transform_data
          type: kubernetes_job
          depends_on: [extract_data]
          image: etl/transformer:latest
          volume: shared-data
    ```
*   **Kubernetes Integration:**
    *   **Job Management:** Create and monitor Kubernetes Jobs
    *   **Resource Management:** CPU/memory requests and limits
    *   **Persistent Volumes:** Shared storage for intermediate data
*   **Advanced Features:**
    *   **Parameter Passing:** Runtime configuration between tasks
    *   **Error Handling:** Retry policies, dead letter queues
    *   **Monitoring:** Pipeline status, task logs, performance metrics
    *   **Version Control:** GitOps for pipeline definitions

### Question 869: How to detect zombie or stuck background workers?

**Answer:**
*   **Heartbeat Mechanism:**
    *   **Regular Pings:** Workers update `LastSeen` timestamp in Redis every 5 seconds
    *   **Health Payload:** Include CPU usage, memory consumption, task queue depth
    *   **Timeout Detection:** Mark as failed after 3 consecutive missed heartbeats
*   **Reaper Service:**
    ```python
    class WorkerReaper:
        def run(self):
            while True:
                workers = redis.hgetall("workers:heartbeat")
                now = time.time()
                
                for worker_id, last_seen in workers.items():
                    if now - float(last_seen) > TIMEOUT:
                        self.handle_zombie_worker(worker_id)
                
                sleep(10)  # Check every 10 seconds
        
        def handle_zombie_worker(self, worker_id):
            # Kill process if local
            if self.is_local_worker(worker_id):
                os.kill(int(worker_id), signal.SIGKILL)
            
            # Reassign tasks
            tasks = self.get_worker_tasks(worker_id)
            for task in tasks:
                self.requeue_task(task)
            
            # Alert operations
            self.alert(f"Zombie worker detected: {worker_id}")
    ```
*   **Detection Strategies:**
    *   **Process Monitoring:** Check if process is still running
    *   **Task Progress:** Verify tasks are making progress
    *   **Resource Monitoring:** Detect hung processes with 0 CPU usage
*   **Recovery Actions:**
    *   **Automatic Restart:** Restart failed workers automatically
    *   **Task Reassignment:** Move in-progress tasks to healthy workers
    *   **Manual Intervention:** Alert operations for persistent issues

### Question 870: Design a backpressure-aware task dispatch system.

**Answer:**
*   **Feedback Loop Architecture:**
    *   **Consumer Metrics:** Report queue depth, CPU load, memory usage in acknowledgments
    *   **Producer Adaptation:** Adjust dispatch rate based on consumer feedback
    *   **Load Shedding:** Drop low-priority tasks when system overloaded
*   **Implementation Strategy:**
    ```python
    class BackpressureDispatcher:
        def dispatch_task(self, task):
            # Get consumer load metrics
            consumer_load = self.get_consumer_metrics()
            
            # Calculate dispatch rate
            if consumer_load.queue_depth > CRITICAL_THRESHOLD:
                self.dispatch_rate = MIN_RATE
            elif consumer_load.queue_depth > WARNING_THRESHOLD:
                self.dispatch_rate *= 0.8  # Reduce by 20%
            else:
                self.dispatch_rate = min(self.dispatch_rate * 1.1, MAX_RATE)
            
            # Apply rate limiting
            if self.can_dispatch():
                self.send_to_consumer(task)
            else:
                self.queue_task(task)  # Buffer for later
    ```
*   **Backpressure Signals:**
    *   **Explicit Signals:** Consumer sends "OVERLOAD" status
    *   **Implicit Signals:** Increased processing time, queue buildup
    *   **Protocol-Level:** HTTP 503, custom headers, TCP backpressure
*   **Adaptive Features:**
    *   **Priority Queuing:** Process high-priority tasks first during overload
    *   **Bounded Queues:** Drop tasks when queues exceed limits
    *   **Circuit Breaking:** Stop sending to repeatedly failing consumers
*   **Monitoring:**
    *   **Throughput Metrics:** End-to-end processing rate
    *   **Latency Tracking:** Time in queue vs processing time
    *   **System Health:** Overall load vs capacity ratio

---

## 🔸 Real-Time & Event-Driven Systems (Questions 871-880)

### Question 871: Design a real-time leaderboard system.

**Answer:**
(See Q191).

### Question 872: Build a system for real-time fraud alerts via SMS/email.

**Answer:**
*   **Speed:** Flink/Spark Streaming.
*   **Priority:** SMS Gateway (Twilio) is expensive. Only send High Confidence alerts via SMS. Low confidence via Email.

### Question 873: How would you handle event ordering in distributed systems?

**Answer:**
*   **Partition Key:** Ensure events for same `EntityID` go to same Partition.
*   **Sequence Num:** Producer attaches ID. Consumer ensures `CurrentID = LastID + 1`.

### Question 874: Build a webhook processor with retries and exponential backoff.

**Answer:**
*   **DB:** `WebhookAttempts`. `Status`, `NextRetryTime`.
*   **Poller:** `SELECT * FROM WebhookAttempts WHERE Status='Pending' AND NextRetryTime < Now`.

### Question 875: Design a publish/subscribe system with guaranteed delivery.

**Answer:**
*   **Durable:** Write to Disk (WAL) before Acking publish.
*   **Offset:** Consumer tracks Offset. Only moves forward after successful process.

### Question 876: Build a system for synchronizing document edits in real-time.

**Answer:**
(See Q522). Operational Transformation.

### Question 877: Design a real-time feed with aggregation and deduplication.

**Answer:**
*   **Buffer:** 1 min window.
*   **Agg:** "User A liked 5 photos" instead of 5 separate notifications.
*   **UI:** "User A and 4 others..."

### Question 878: Build a real-time sentiment analysis system for news.

**Answer:**
*   **Ingest:** RSS Feeds / Twitter API.
*   **Compute:** BERT Model.
*   **Dashboard:** Kibana showing "Negative Sentiment Spike" on "company X".

### Question 879: Design a collaborative whiteboard backend.

**Answer:**
(See Q521).

### Question 880: Build a real-time alerting system for critical operations.

**Answer:**
(See Q636).

---

## 🔸 IoT, Sensors & Device Management (Questions 881-890)

### Question 881: Design a system to track millions of connected devices.

**Answer:**
*   **Registry:** DeviceID, Certs, Owner.
*   **Shadow:** Current State (Online/Offline, Temp, FW Version). DynamoDB.

### Question 882: Build an alert system for abnormal sensor behavior.

**Answer:**
*   **Threshold:** Static (`Temp > 50`).
*   **Anomaly:** Dynamic (`Temp` changes > 10 degrees in 1 min).

### Question 883: How do you handle firmware updates over-the-air?

**Answer:**
*   **Manifest:** JSON file URL, Checksum, Version.
*   **Rollout:** Canary (1%).
*   **A/B Partition:** Device has Partition A (Active) and B (Idle). Update B. Reboot to B. If fail, revert to A.

### Question 884: Design a telemetry ingestion pipeline for edge devices.

**Answer:**
*   **Protocol:** MQTT (Lightweight).
*   **Topic:** `devices/{id}/telemetry`.
*   **Bridge:** IoT Core Rule -> Kinesis -> Firehose -> S3.

### Question 885: Build a digital twin platform for connected devices.

**Answer:**
*   **Concept:** Virtual representation of physical device.
*   **Sync:** Device writes to Twin. App reads Twin. (Decoupled).
*   **Simulation:** Run "What If" scenarios on Twin (e.g., increased load) without risking hardware.

### Question 886: Design a rule engine to trigger actions from sensor data.

**Answer:**
(See Q861).
*   **IoT SQL:** `SELECT * FROM 'topic/#' WHERE temp > 50`.
*   **Action:** Lambda / SNS.

### Question 887: Build a configuration push system for IoT endpoints.

**Answer:**
*   **Desire:** App sets "Desired Config".
*   **Report:** Device reports "Reported Config".
*   **Delta:** Device calculates Delta and applying.

### Question 888: Design a geofencing service for IoT-enabled vehicles.

**Answer:**
(See Q846).

### Question 889: Build a device registration and provisioning platform.

**Answer:**
*   **Factory:** Flash unique certificate in factory.
*   **Claim:** User scans QR code.
*   **Verification:** Cloud verifies Cert signature against CA. Ownership Transfer to User.

### Question 890: Design a secure data sync protocol for unreliable networks.

**Answer:**
*   **Store-and-Forward:** Device stores msg in flash. Retries forever until Ack.
*   **Deduplication:** Cloud handles duplicate sends (Idempotency).

---

## 🔸 Edge, Scale & Chaos Engineering (Questions 891-900)

### Question 891: Design a system for edge caching with invalidation strategies.

**Answer:**
*   **Edge:** Nginx/Varnish at POPs.
*   **Invalidation:**
    *   **Purge:** Explicit API call `PURGE /file`. (Fast, expensive).
    *   **Tagging:** Purge by Tag `PURGE-TAG product-123`.
    *   **TTL:** Short TTL for dynamic (1 min).

### Question 892: How to architect a resilient service mesh?

**Answer:**
*   **Sidecar:** Envoy Proxy handles all traffic.
*   **Features:** Retry, Timeout, Circuit Breaker implemented in Sidecar (Code agnostic).
*   **State:** Control Plane (Istio) pushes config to Sidecars. Data Plane (Envoy) is robust.

### Question 893: Design a chaos engineering tool to inject latency and failure.

**Answer:**
*   **Network:** `tc` (Traffic Control) on Linux. Drop 5% packets. Delay 500ms.
*   **App:** Middleware / Library injection. `if Random() < 0.05 { sleep(5s) }`.
*   **Target:** Select random pods in K8s.

### Question 894: Design a distributed circuit breaker system.

**Answer:**
*   **Local:** Each instance tracks error rate. Open circuit locally.
*   **Distributed:** Aggregate error rates to Redis. If Global Error > Threshold, trigger "Panic Mode" (Open circuit on all instances).

### Question 895: How to handle thundering herd at the edge?

**Answer:**
*   **Request Collapsing:** If 1000 users request `video.mp4` simultaneously, Edge makes ONE request to Origin. Streams response to 1000 users.
*   **Jitter:** Add random delay to client retries.

### Question 896: Design a multi-CDN strategy manager.

**Answer:**
*   **DNS:** Traffic Manager.
*   **Probe:** Client measures latency to Cloudfront, Akamai, Fastly.
*   **Route:** DNS returns CNAME of fastest CDN for that user segment.

### Question 897: Build a system for graceful degradation during partial outages.

**Answer:**
*   **feature-flags:** `enable_reviews = false`.
*   **Fallback:** If `Recommendations` service fails, return `Top 10 Global Popular` (Static Cache).
*   **UX:** Show "Some features unavailable" banner.

### Question 898: Design a rate-limiting sidecar for microservices.

**Answer:**
*   **Deployment:** Container in same Pod.
*   **Logic:** Token Bucket in shared memory (SHM).
*   **Control:** Sidecar intercepts ingress/egress. Deny if limit exceeded. Reduces app code complexity.

### Question 899: How to implement backpressure in a mesh network?

**Answer:**
*   **Health:** Upstream service reports "Overloaded" (HTTP 503 / Custom Header).
*   **Downstream:** Detects signal. Reduces request rate / Retries with exponential backoff.
*   **Queue:** Bounded queues. Drop requests if full (Load Shedding).

### Question 900: Design a distributed tracing collector for edge nodes.

**Answer:**
*   **Sampling:** Head-based (1%). Tail-based (Keep 100% of errors).
*   **Buffer:** Circular buffer on Edge.
*   **Flush:** Async flush to Collector (Jaeger).
*   **Correlation:** Inject `x-trace-id` header at Edge Ingress.
