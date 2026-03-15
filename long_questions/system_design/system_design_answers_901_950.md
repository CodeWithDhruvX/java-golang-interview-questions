## 🔸 Serverless & Event-Driven Patterns (Questions 901-910)

### Question 901: Design a serverless backend for image resizing.

**Answer:**
*   **Trigger Mechanism:** S3 `ObjectCreated` event invokes Lambda function
*   **Processing Pipeline:**
    *   **Lambda Function:** Resizes image using ImageMagick/Sharp library
    *   **Layer Strategy:** Pre-packaged ImageMagick layer for faster cold starts
    *   **Output Storage:** Write resized images to `bucket-resized/` prefix
*   **Infinite Loop Prevention:**
    *   **Prefix Isolation:** Different buckets for input/output
    *   **Tag Filtering:** Only process images without `resized` tag
    *   **Metadata Check:** Skip if image already processed
*   **Performance Optimizations:**
    *   **Memory Allocation:** 1024MB RAM for large images
    *   **Concurrent Execution:** 1000 concurrent Lambda limit
    *   **Batch Processing:** Process multiple sizes in single invocation

### Question 902: How do you handle "Cold Start" in FaaS?

**Answer:**
*   **Keep-Alive Strategies:**
    *   **Scheduled Pings:** CloudWatch event triggers Lambda every 5 minutes
    *   **Provisioned Concurrency:** AWS feature keeps N instances warm (costly)
    *   **Dummy Payloads:** Send lightweight requests to maintain warmth
*   **Optimization Techniques:**
    *   **Language Choice:** Go/Rust (fast startup) vs Java/Spring (JVM overhead)
    *   **Package Size:** Minimize dependencies, use Lambda layers
    *   **Initialization Code:** Move expensive operations outside handler
*   **Architecture Patterns:**
    *   **Warm Standby:** Always-on instances for critical functions
    *   **Pre-warming:** Trigger functions before expected load
    *   **Graceful Degradation:** Serve cached responses during cold starts

### Question 903: Build a stateful workflow using serverless functions.

**Answer:**
*   **Orchestration Framework:** AWS Step Functions / Azure Logic Apps / Temporal Cloud
*   **State Management Strategy:**
    *   **External State:** DynamoDB table keyed by WorkflowID
    *   **Immutable State:** Each step writes new state, never modifies
    *   **State Passing:** JSON output of Step A becomes input for Step B
*   **Implementation Pattern:**
    ```json
    {
      "Comment": "Order processing workflow",
      "StartAt": "ValidateOrder",
      "States": {
        "ValidateOrder": {
          "Type": "Task",
          "Resource": "arn:aws:lambda:...:ValidateOrder",
          "Next": "ProcessPayment"
        }
      }
    }
    ```
*   **Error Handling:**
    *   **Retry Policies:** Configurable per state
    *   **Catch Blocks:** Fallback states for error recovery
    *   **Compensation:** Undo steps for saga pattern

### Question 904: Design a serverless API with rate limiting per tenant.

**Answer:**
*   **Gateway Strategy:**
    *   **API Gateway Usage Plans:** Built-in rate limiting and quotas
    *   **Custom Authorizer:** Lambda function for fine-grained control
*   **Rate Limiting Implementation:**
    *   **Redis Storage:** `Rate:{TenantID}:{Window}` -> RequestCount
    *   **Sliding Window:** More accurate than fixed windows
    *   **Token Bucket:** Allow bursts within limits
*   **Authorization Flow:**
    ```python
    def lambda_handler(event):
        tenant_id = extract_tenant(event['headers'])
        key = f"rate:{tenant_id}:{get_minute_window()}"
        
        if redis.incr(key) > LIMIT:
            return {'statusCode': 429}
        
        return generate_policy('Allow', event)
    ```
*   **Multi-Tier Limits:**
    *   **Free Tier:** 100 requests/hour
    *   **Basic Tier:** 1000 requests/hour
    *   **Enterprise Tier:** 10000 requests/hour

### Question 905: How to coordinate distributed transactions in serverless?

**Answer:**
*   **Saga Pattern Implementation:** (See Q329) Enhanced for serverless
*   **Step Functions Approach:**
    ```json
    {
      "StartAt": "ProcessPayment",
      "States": {
        "ProcessPayment": {
          "Type": "Task",
          "Resource": "arn:aws:lambda:...:ProcessPayment",
          "Catch": [{"ErrorEquals": ["States.ALL"], "Next": "RefundPayment"}],
          "Next": "ReserveInventory"
        },
        "RefundPayment": {
          "Type": "Task",
          "Resource": "arn:aws:lambda:...:RefundPayment",
          "End": true
        }
      }
    }
    ```
*   **Compensation Logic:**
    *   **Automatic:** Step Functions trigger compensation on failure
    *   **Manual:** Human approval for large transactions
    *   **Timeout:** Auto-rollback if steps take too long
*   **State Persistence:**
    *   **Step Functions:** Built-in state tracking
    *   **DynamoDB:** Transaction log for audit trail
    *   **Dead Letter Queue:** Handle failed compensations

### Question 906: Build a real-time chat using serverless websockets.

**Answer:**
*   **WebSocket Management:**
    *   **API Gateway WebSocket API:** Maintains persistent connections
    *   **Connection Routes:** `$connect`, `$disconnect`, `$default` for message handling
    *   **Connection Storage:** DynamoDB table `ConnectionID -> UserID, RoomID, LastSeen`
*   **Message Broadcasting:**
    *   **Fan-out Pattern:** SNS topic for each room
    *   **Direct Push:** `POST https://.../@connections/{conn_id}` for targeted messages
    *   **Message Persistence:** DynamoDB for chat history
*   **Implementation Details:**
    ```python
    def on_connect(event):
        conn_id = event['requestContext']['connectionId']
        user_id = extract_user_from_token(event)
        dynamodb.put_item(Item={'ConnectionID': conn_id, 'UserID': user_id})
    
    def on_message(event):
        message = json.loads(event['body'])
        room_id = message['room']
        # Get all connections in room
        connections = get_room_connections(room_id)
        # Broadcast to all
        for conn in connections:
            apigateway_management.post_to_connection(
                ConnectionId=conn, Data=message['content']
            )
    ```
*   **Scaling Considerations:**
    *   **Connection Limits:** 100,000 concurrent connections per region
    *   **Message Throttling:** Rate limit per user to prevent spam
    *   **Reconnection Logic:** Automatic reconnection with message replay

### Question 907: Design a serverless map-reduce job.

**Answer:**
*   **Map Phase Implementation:**
    *   **S3 Batch Operations:** Process 1M objects in parallel
    *   **Lambda Mapper:** Each invocation processes chunk of data
    *   **Intermediate Storage:** Write results to DynamoDB/S3
*   **Shuffle Phase:**
    *   **DynamoDB Streams:** Trigger on intermediate data writes
    *   **Key-based Partitioning:** Ensure same keys go to same reducer
    *   **Aggregation Buffer:** Accumulate values before reduction
*   **Reduce Phase:**
    *   **Reducer Lambda:** Triggered by DynamoDB Streams
    *   **Final Output:** Write aggregated results to S3
*   **Architecture Flow:**
    ```
    S3 Input -> Lambda Map -> DynamoDB Intermediate -> 
    DynamoDB Streams -> Lambda Reduce -> S3 Output
    ```
*   **Optimization Strategies:**
    *   **Batch Size:** Optimize Lambda payload for cost/performance
    *   **Parallelism:** Control concurrency with reserved capacity
    *   **Error Handling:** Dead letter queues for failed processing

### Question 908: How to secret management in serverless?

**Answer:**
*   **Storage Strategies:**
    *   **Environment Variables:** Encrypted at rest, easy access
    *   **AWS Secrets Manager:** Centralized secret storage with rotation
    *   **Parameter Store:** Hierarchical configuration with IAM policies
*   **Access Patterns:**
    *   **Runtime Fetch:** Load secrets at Lambda cold start
    *   **Caching:** Store in `/tmp` filesystem for warm invocations
    *   **Lazy Loading:** Fetch only when needed to reduce cost
*   **Implementation Example:**
    ```python
    import json
    
    # Global variable for caching
    _secrets = None
    
    def get_secrets():
        global _secrets
        if _secrets is None:
            client = boto3.client('secretsmanager')
            response = client.get_secret_value(SecretId='my-app-secrets')
            _secrets = json.loads(response['SecretString'])
            # Cache to /tmp for faster access
            with open('/tmp/secrets.json', 'w') as f:
                json.dump(_secrets, f)
        return _secrets
    ```
*   **Security Best Practices:**
    *   **Least Privilege:** IAM roles restrict access to specific secrets
    *   **Rotation:** Automatic secret rotation every 90 days
    *   **Audit Logging:** CloudTrail tracks all secret access

### Question 909: Design a fan-out notification system using serverless.

**Answer:**
*   **Event Distribution Architecture:**
    *   **SNS Topic:** `NewPost` as central event bus
    *   **Multi-Protocol Support:** Email, Push, SMS, Webhook
    *   **Parallel Processing:** Thousands of concurrent Lambda invocations
*   **Subscriber Implementation:**
    ```yaml
    # CloudFormation Template
    Resources:
      NewPostTopic:
        Type: AWS::SNS::Topic
        Properties:
          TopicName: NewPost
          
      EmailSubscriber:
        Type: AWS::SNS::Subscription
        Properties:
          TopicArn: !Ref NewPostTopic
          Protocol: lambda
          Endpoint: !GetAtt EmailFunction.Arn
          
      PushSubscriber:
        Type: AWS::SNS::Subscription
        Properties:
          TopicArn: !Ref NewPostTopic
          Protocol: lambda
          Endpoint: !GetAtt PushFunction.Arn
    ```
*   **Scaling Characteristics:**
    *   **Burst Capacity:** 3000 concurrent Lambda executions
    *   **Message Filtering:** SNS attribute-based filtering
    *   **Dead Letter Queue:** Failed notifications for retry
*   **Performance Optimizations:**
    *   **Batch Processing:** Process multiple notifications per invocation
    *   **Async Patterns:** Fire-and-forget for non-critical notifications
    *   **Priority Queues:** Separate topics for urgent vs batch notifications

### Question 910: Build a serverless URL shortener.

**Answer:**
*   **Write Path:**
    *   **API Gateway:** `POST /shorten` endpoint
    *   **Lambda Function:** Generate hash, store in DynamoDB
    *   **Storage:** DynamoDB table `Hash -> {LongURL, CreatedAt, Clicks}`
*   **Read Path:**
    *   **API Gateway:** `GET /{hash}` endpoint
    *   **Lambda Function:** Lookup hash, return 301 redirect
    *   **Analytics:** Increment click counter asynchronously
*   **Caching Strategy:**
    *   **CloudFront CDN:** Cache redirects at edge locations
    *   **TTL:** 1 hour for popular URLs, 5 minutes for others
    *   **Cache Invalidation:** Update on URL changes
*   **Hash Generation:**
    ```python
    import base62, hashlib
    
    def generate_short_code(long_url):
        hash = hashlib.sha256(long_url.encode()).hexdigest()[:6]
        return base62.encode(int(hash, 16))
    ```
*   **Performance Metrics:**
    *   **Latency:** < 50ms for cached redirects
    *   **Hit Ratio:** 99% cache hit rate for popular URLs
    *   **Cost Reduction:** 99% fewer Lambda invocations due to CDN

---

## 🔸 Big Data & Analytics Deep Dive (Questions 911-920)

### Question 911: Design a Data Lake architecture on cloud.

**Answer:**
*   **Zone-Based Architecture:**
    *   **Raw Zone:** Immutable original data (JSON, CSV, logs) with versioning
    *   **Curated Zone:** Cleaned, validated data in columnar formats (Parquet/Avro)
    *   **Consumer Zone:** Aggregated tables optimized for BI and analytics
*   **Metadata Management:**
    *   **Data Catalog:** AWS Glue / Hive Metastore for schema discovery
    *   **Partitioning Strategy:** Date-based partitioning for efficient queries
    *   **Schema Registry:** Track schema evolution and compatibility
*   **Data Processing Pipeline:**
    ```
    Raw Data -> Validation -> Cleaning -> Transformation -> Curated Data
        |                                            |
        |                                            v
        +----> Archive (S3 Glacier) <--- Analytics <---+
    ```
*   **Security & Governance:**
    *   **Access Control:** IAM policies per zone and dataset
    *   **Data Encryption:** At rest (S3 KMS) and in transit (TLS)
    *   **Audit Logging:** CloudTrail for data access tracking

### Question 912: How to handle schema evolution in Parquet files?

**Answer:**
*   **Schema Discovery:**
    *   **Automated Crawlers:** AWS Glue crawlers detect new columns and data types
    *   **Schema Registry:** Central repository for schema versions and compatibility
    *   **Validation:** Ensure backward compatibility before schema changes
*   **Modern Table Formats:**
    *   **Apache Iceberg:** Metadata files track schema per data file
    *   **Delta Lake:** Transaction log with schema evolution support
    *   **Apache Hudi:** Timeline-based schema management
*   **Implementation Strategy:**
    ```sql
    -- Iceberg example
    ALTER TABLE my_table ADD COLUMNS (new_column string);
    -- Metadata tracks: Schema v1 applies to files_1-100, Schema v2 applies to files_101+
    ```
*   **Query-Time Resolution:**
    *   **Schema Merge:** Combine schemas from multiple files on read
    *   **Default Values:** Handle missing columns in older schemas
    *   **Type Promotion:** Safe type conversions (int -> bigint)

### Question 913: Design a real-time dashboard for log analytics.

**Answer:**
*   **Data Ingestion Pipeline:**
    *   **Log Collection:** Filebeat/Fluentd agents on servers
    *   **Buffer Layer:** Kafka topics for log aggregation and buffering
    *   **Processing:** Flink/Spark Streaming for enrichment and filtering
*   **Storage Architecture:**
    *   **Hot Tier:** Elasticsearch cluster for recent data (7 days)
    *   **Warm Tier:** Elasticsearch with slower disks for historical data (30 days)
    *   **Cold Tier:** S3 for archival data beyond 30 days
*   **Visualization Stack:**
    ```yaml
    # Kibana Dashboard Configuration
    dashboards:
      - name: "Error Rate Analysis"
        visualization: "line-chart"
        query: "level:ERROR AND timestamp:[now-1h TO now]"
        refresh: "1m"
      - name: "Top Error Sources"
        visualization: "bar-chart"
        aggregation: "terms:source_ip:10"
    ```
*   **Performance Optimization:**
    *   **Index Templates:** Optimize mappings for log data
    *   **Rollup Policies:** Aggregate old data to reduce storage
    *   **Caching:** Redis cache for frequent dashboard queries

### Question 914: Build a privacy-compliant data deletion pipeline (GDPR).

**Answer:**
*   **Technical Challenge:**
    *   **Columnar Storage:** Parquet files store data column-wise in large blocks (1GB+)
    *   **Immutable Files:** Cannot delete individual rows without rewriting entire file
    *   **Distributed Storage:** Data spread across multiple files and partitions
*   **Delta Lake Solution:**
    ```sql
    -- Logical deletion (immediate)
    DELETE FROM user_data WHERE user_id = '12345';
    
    -- Physical deletion (async)
    VACUUM user_data RETAIN 0 HOURS; -- Removes tombstoned data
    ```
*   **Implementation Pipeline:**
    *   **Tombstone Records:** Mark deleted rows with deletion timestamp
    *   **Compaction Jobs:** Periodically rewrite files excluding tombstoned data
    *   **Verification:** Audit logs to confirm data removal
*   **Multi-System Coordination:**
    *   **Search Indexes:** Delete from Elasticsearch/Solr
    *   **Cache Layers:** Clear Redis/Memcached entries
    *   **Backup Systems:** Exclude from future restores or encrypt backups
*   **Compliance Features:**
    *   **Deletion Reports:** Automated certificates of data destruction
    *   **Audit Trails:** Immutable logs of all deletion requests
    *   **Data Mapping:** Track data flow across systems for complete deletion

### Question 915: Design a high-cardinality metrics storage engine.

**Answer:**
*   **Time Series Database Architecture:** (Cortex/Thanos pattern)
    *   **Inverted Index:** Labels -> SeriesID mappings for fast lookups
    *   **Chunked Storage:** Time-series data in compressed chunks
    *   **Distributed Architecture:** Horizontal scaling across nodes
*   **Cardinality Challenge:**
    ```python
    # Problem: Too many unique label combinations
    metrics = {
        'request_count{pod="pod-1", ip="10.0.0.1", user="alice"}': 100,
        'request_count{pod="pod-2", ip="10.0.0.2", user="bob"}': 150,
        # Millions of unique combinations = index explosion
    }
    ```
*   **Optimization Strategies:**
    *   **Label Cardinality Limits:** Restrict high-cardinality labels (IP, user_id)
    *   **Rollup Policies:** Drop detailed dimensions after retention period
        *   Raw data: Keep all dimensions for 24h
        *   Hourly rollup: Drop IP dimension, keep pod/service
        *   Daily rollup: Keep only aggregate metrics
    *   **Pre-aggregation:** Store aggregates alongside raw data
*   **Storage Optimization:**
    *   **Chunk Encoding:** Gorilla compression for time series
    *   **Index Sharding:** Distribute index across multiple nodes
    *   **Downsampling:** Reduce data resolution for older data

### Question 916: How to deduplicate data in a Kappa architecture?

**Answer:**
*   **Kappa Architecture Principles:**
    *   **Stream-Only Processing:** No separate batch layer, everything through streams
    *   **Immutable Events:** All data as immutable event stream
    *   **Replay Capability:** Recompute state by replaying events
*   **Deduplication Strategy:**
    ```java
    // Flink deduplication example
    public class DeduplicationFunction extends KeyedProcessFunction<String, Event, Event> {
        private ValueState<Boolean> seenState;
        private ValueState<Long> timestampState;
        
        @Override
        public void processElement(Event event, Context ctx, Collector<Event> out) {
            if (seenState.value() == null) {
                // First time seeing this event
                seenState.update(true);
                timestampState.update(ctx.timestamp());
                out.collect(event);
                
                // Set timer to clean up after 7 days
                ctx.timerService().registerEventTimeTimer(
                    ctx.timestamp() + TimeUnit.DAYS.toMillis(7));
            }
        }
    }
    ```
*   **State Management:**
    *   **RocksDB State Store:** Persistent key-value store for deduplication state
    *   **Message Hash:** SHA-256 hash of event payload for uniqueness
    *   **TTL Management:** Automatic cleanup of old deduplication entries
*   **Scaling Considerations:**
    *   **Key Partitioning:** Hash-based partitioning for load distribution
    *   **State Size:** Monitor state store size to prevent memory issues
    *   **Checkpointing:** Regular checkpoints for fault tolerance

### Question 917: Design a system to join two infinite streams.

**Answer:**
*   **Stream Join Challenges:** (See Q557) Enhanced with implementation details
*   **Temporal Constraints:**
    *   **Window Definition:** "Join if Click happens within 10 minutes of Impression"
    *   **Event Time vs Processing Time:** Handle out-of-order events
    *   **Watermark Strategy:** Track event time progress for window boundaries
*   **Implementation Architecture:**
    ```python
    # Flink stream join example
    impressions = kafka_source.consume("impressions")
    clicks = kafka_source.consume("clicks")
    
    # Join streams with 10-minute window
    joined = impressions.join(
        clicks,
        where=lambda imp, clk: imp.user_id == clk.user_id,
        window=TumblingEventTimeWindows.of(Time.minutes(10))
    )
    ```
*   **State Management:**
    *   **Buffer Storage:** Keep impressions in state store for matching window
    *   **Memory Management:** Evict old events outside window boundaries
    *   **Backpressure Handling:** Slow down producers if consumers can't keep up
*   **Optimization Strategies:**
    *   **Pre-filtering:** Reduce data volume before join
    *   **Partitioning:** Co-partition streams on join keys
    *   **Window Types:** Tumbling, sliding, or session windows based on use case

### Question 918: Build a distributed crawler for the web.

**Answer:**
*   **Frontier Management:**
    *   **URL Queue:** Redis priority queue ordered by importance/recency
    *   **Domain-based Partitioning:** Distribute URLs by domain hash
    *   **Seen URLs:** Bloom filter to avoid duplicate crawling
*   **Politeness Policy:**
    ```python
    class PolitenessManager:
        def __init__(self):
            self.domain_last_access = {}  # Map<Domain, LastAccessTime>
            self.crawl_delay = defaultdict(lambda: 1.0)  # Default 1 second
        
        def can_crawl(self, url):
            domain = extract_domain(url)
            last_access = self.domain_last_access.get(domain, 0)
            
            if time.time() - last_access < self.crawl_delay[domain]:
                return False  # Too soon, respect robots.txt
            
            self.domain_last_access[domain] = time.time()
            return True
    ```
*   **Storage Architecture:**
    *   **WARC Files:** Standard web archive format in S3
    *   **Metadata Store:** PostgreSQL for URL status, crawl history
    *   **Index Layer:** Elasticsearch for content search
*   **Scaling Components:**
    *   **DNS Resolver:** Caching DNS resolver for performance
    *   **User Agent Rotation:** Rotate user agents to avoid blocking
    *   **Proxy Pool:** Distributed proxy network for IP rotation
*   **Quality Control:**
    *   **Content Validation:** Check for spam, duplicates, low-quality content
    *   **Link Extraction:** Parse and queue new discovered URLs
    *   **Crawl Scheduling:** Prioritize important/frequently updated sites

### Question 919: How to optimize top-K queries on big data?

**Answer:**
*   **Approximation Algorithms:**
    *   **Count-Min Sketch:** Probabilistic data structure for frequency estimation
    *   **Space Efficiency:** O(k * log n) space vs O(n) for exact counting
    *   **Error Bounds:** Configurable error rate and confidence
*   **MapReduce Implementation:**
    ```python
    # Mapper phase
    def mapper(document):
        words = tokenize(document)
        local_counts = Counter(words)
        return local_counts.most_common(K)  # Local top-K
    
    # Reducer phase
    def reducer(local_top_k_lists):
        global_counts = Counter()
        for local_list in local_top_k_lists:
            global_counts.update(local_list)
        return global_counts.most_common(K)  # Global top-K
    ```
*   **Optimization Strategies:**
    *   **Combiner Optimization:** Reduce data transfer between mappers and reducers
    *   **Sampling-Based:** Use random sampling for approximate results
    *   **Hierarchical Aggregation:** Multi-level aggregation for massive datasets
*   **Real-Time Top-K:**
    *   **Streaming Algorithms:** Maintain top-K in single pass
    *   **Heavy Hitters:** Identify items above frequency threshold
    *   **Sliding Window:** Top-K over recent time windows
*   **Accuracy vs Performance Trade-offs:**
    *   **Exact:** Full sort, high accuracy, slow
    *   **Approximate:** Sketch-based, 95% accuracy, fast
    *   **Hybrid:** Combine exact for top items, approximate for tail

### Question 920: Design a data lineage tracking system.

**Answer:**
*   **Lineage Graph Model:**
    *   **Nodes:** Datasets, jobs, transformations, users
    *   **Edges:** Data flow relationships (reads, writes, transforms)
    *   **Metadata:** Timestamps, schema changes, quality metrics
*   **Ingestion Mechanisms:**
    ```sql
    -- SQL parser extracts lineage information
    -- Input: INSERT INTO target SELECT col1, col2 FROM source1 JOIN source2
    -- Output: source1 -> target, source2 -> target relationships
    
    CREATE TABLE lineage_edges (
        source_dataset VARCHAR,
        target_dataset VARCHAR,
        job_id VARCHAR,
        transformation_type VARCHAR,
        timestamp TIMESTAMP
    );
    ```
*   **Automated Extraction:**
    *   **SQL Parsers:** Analyze SELECT, INSERT, UPDATE statements
    *   **API Integration:** Extract from Spark, Airflow, dbt jobs
    *   **Log Analysis:** Parse job logs for data dependencies
*   **Visualization Features:**
    *   **Impact Analysis:** "If I change Table X, what downstream jobs are affected?"
    *   **Root Cause Analysis:** "Table Y has bad data, trace upstream to find source"
    *   **Data Quality Propagation:** Track quality metrics through pipeline
*   **Advanced Capabilities:**
    *   **Schema Lineage:** Track column-level transformations
    *   **PII Detection:** Identify personal data flow for compliance
    *   **Cost Attribution:** Assign storage/compute costs to data sources

---

## 🔸 AR/VR & Spatial Computing (Questions 921-930)

### Question 921: Design a Point Cloud streaming system for AR.

**Answer:**
*   **Format:** PLY / PCD.
*   **Octree:** Divide 3D space. Stream only "Visible" nodes based on Camera Frustum.
*   **LOD:** Stream low detail first -> refine.

### Question 922: Build a Spatial Anchor synchronization system.

**Answer:**
*   **Anchor:** `ID`, `Descriptor (Visual Features)`, `Pose (x,y,z)`.
*   **Look up:** Device sends Camera Features. Server matches Descriptor. Returns Pose relative to World.
*   **Cloud:** Azure Spatial Anchors / Google Cloud Anchors.

### Question 923: Design a backend for a massive multiplayer VR world (Metaverse).

**Answer:**
*   **Partitioning:** Spatial Hashing. Server A handles (0,0) to (100,100).
*   **Handoff:** User walks across boundary -> Transfer session A to B.
*   **Voice:** Proximity Chat (Volume decreases with distance).

### Question 924: How to handle physics sync in networked VR?

**Answer:**
*   **Authority:** Server authoritative.
*   **Prediction:** Client predicts physics locally. Packet arrives -> Correction (Lerp).
*   **Ownership:** "User holding the cup" has authority over Cup physics to ensure 0 latency interaction.

### Question 925: Design a collaborative 3D modeling platform.

**Answer:**
*   **CRDT:** Tree structure of Scene Graph (`Root -> Transform -> Cube`).
*   **Locking:** User selects Vertex. Lock vertex.

### Question 926: Build a low-latency 360-video streaming service.

**Answer:**
*   **Projection:** Equirectangular.
*   **Tiling:** Cut 4K video into 4x4 tiles.
*   **ABR:** Based on Head Orientation ("Look at"), download tiles in view at High Res, others at Low Res.

### Question 927: How to persistent AR content in the real world?

**Answer:**
*   **Localization:** GPS (Coarse) + VPS (Visual Positioning System using Street View).
*   **DB:** `GeoHash` -> `List[ContentID]`.

### Question 928: Design a gesture recognition pipeline.

**Answer:**
*   **Edge:** Hand Tracking model (MediaPipe) runs on device (30fps).
*   **Output:** Skeleton (21 points).
*   **Classifier:** "Index + Thumb touching" = Click.

### Question 929: Build a backend for digital avatars and inventory.

**Answer:**
(See Q764). Interoperable standard (VRM).

### Question 930: Design a spatial audio engine for VR.

**Answer:**
*   **Input:** `SourcePos`, `ListenerPos`, `RoomGeometry`.
*   **HRTF:** Head Related Transfer Function.
*   **Calc:** Done Client-Side (CPU heavy). Server only syncs positions.

---

## 🔸 Gaming & Real-Time Simulation (Questions 931-940)

### Question 931: Design a matchmaking system for ranked games.

**Answer:**
*   **Pool:** Users with `MMR` (Matchmaking Rating).
*   **Bucket:** Group by `MMR +/- 100` and `Ping < 50ms`.
*   **Expansion:** If no match in 30s, expand range to `+/- 200`.

### Question 932: Build a state replication system for FPS games.

**Answer:**
*   **Snapshot:** Send full state every 50ms? (Too much bandwidth).
*   **Delta:** Send only changes.
*   **Compression:** Quantize Floats (Send angle as Byte 0-255).

### Question 933: Design a leaderboards system with seasonal resets.

**Answer:**
*   **Redis:** `Rank:Season_1`.
*   **Archive:** At season end, rename key to `Rank:Season_1_Final`. Create empty `Rank:Season_2`.

### Question 934: How to detect aimbots and wallhacks?

**Answer:**
*   **Server-Side:** "Heuristics".
    *   Did cursor snap 180 degrees in 1 frame?
    *   Did bullet hit head through "Unbreakable Wall"?
*   **Client-Side:** Anti-Cheat kernel driver (e.g., Vanguard) checks memory injection.

### Question 935: Design a global game server scaling strategy.

**Answer:**
*   **Agones:** K8s Custom Resource for Game Servers.
*   **Allocation:** Request game server -> Returns IP:Port.
*   **Lifecycle:** Server State `Allocated` -> `Ready` -> `Shutdown`.

### Question 936: Build an in-game economy and trading system.

**Answer:**
*   **ACID:** DB Transaction essential.
*   **Audit:** Record `Transfer(From: A, To: B, Item: Sword)`.
*   **Binding:** Some items "Soulbound" (Cannot trade).

### Question 937: Design a replay system for strategy games.

**Answer:**
*   **Deterministic:** Save `RandomSeed` and `List[PlayerInputs]`.
*   **Playback:** Re-simulate game from Seed applying inputs at exact frames. Result is identical.

### Question 938: How to minimize lag compensation (Rewind)?

**Answer:**
*   **Hitbox:** Server stores history of hitboxes for last 1s.
*   **Check:** "At T=100, Player shot. At T=100, where was Enemy?" Rewind enemy to position at T=100. Check collision.

### Question 939: Design a system for seamless map loading (Open World).

**Answer:**
*   **Chunking:** Divide world into grids.
*   **Streaming:** Load Grid (0,1) when player approaches within 100m. Unload Grid (0,-1).

### Question 940: Build a cross-platform save system (PC/Console/Mobile).

**Answer:**
*   **Conflict:** Timestamp check.
*   **Schema:** Binary Blob or JSON?
*   **Storage:** Cloud Key-Value store `SaveSlot_1`.

---

## 🔸 Web3 & Decentralized Applications (Questions 941-950)

### Question 941: Design a decentralized file storage system (IPFS inspired).

**Answer:**
*   **Addressing:** CIA (Content Addressed). `Hash(Content)`.
*   **DHT:** Kademlia DHT to find "Who has this hash?".
*   **Swarm:** Download chunks from multiple peers.

### Question 942: Build an indexing service for blockchain events (The Graph).

**Answer:**
*   **Ingest:** Connect to ETH Node RPC.
*   **Scan:** Filter logs by `ContractAddress` and `Topic`.
*   **Map:** Apply mapping logic (`Transfer -> UserBalance`).
*   **Store:** Postgres.
*   **Serve:** GraphQL.

### Question 943: Design a cryptocurrency wallet backend (Custodial).

**Answer:**
*   **Hot Wallet:** Online. Holds 5% funds for withdrawals.
*   **Cold Wallet:** Offline/HSM. Holds 95%.
*   **Signing:** Key never leaves HSM. Transaction sent to HSM -> Signed Tx returned.

### Question 944: How to implement "Sign in with Ethereum"?

**Answer:**
*   **Challenge:** Server generates `Nonce`.
*   **Sign:** User signs message `Nonce + Domain` with Private Key.
*   **Verify:** Server runs `ecrecover(Sig)`. If Address matches User, Auth successful.

### Question 945: Design a decentralized oracle (Chainlink).

**Answer:**
*   **Aggregation:** 20 Nodes fetch "Price of ETH" from different APIs (Binance, Coinbase).
*   **Commit:** Nodes submit value on-chain.
*   **Consensus:** Smart Contract takes Median of values.
*   **Reward:** Nodes paid in link for honest data.

### Question 946: Build a Layer 2 rollup relayer.

**Answer:**
*   **Accumulate:** Collect 1000 txs off-chain.
*   **Proof:** Generate ZK-Proof (Validity) or Merkle Root (Optimistic).
*   **Post:** Submit Root + CallData to L1. Saves gas.

### Question 947: Design a Merkle Tree for whitelist verification.

**Answer:**
*   **Off-chain:** Hash all 10k whitelist addresses -> Root.
*   **On-chain:** Store only Root.
*   **Proof:** User sends `[Hash A, Hash B...]` (Merkle Path) to prove they are in the tree.

### Question 948: Build a decentralized voting system (DAO).

**Answer:**
*   **Snapshot:** Off-chain signing. Free.
*   **Weight:** `BalanceOf(Token)` at Block Height X.
*   **Storage:** Store votes on IPFS.

### Question 949: Design a bridge between two blockchains.

**Answer:**
*   **Lock:** User locks Token A on Chain 1.
*   **Witness:** Relayer detects Lock Event.
*   **Mint:** Relayer calls Contract on Chain 2 to Mint Wrapped Token A.
*   **Risk:** 51% attack on Bridge Validators (Multi-sig).

### Question 950: How to mitigate MEV (Miner Extractable Value) in block building?

**Answer:**
*   **Private Mempool:** Flashbots.
*   **Bundle:** Searchers submit "Bundle" of txs.
*   **Promise:** Miner guarantees bundle order and no front-running, or reverts bundle.
