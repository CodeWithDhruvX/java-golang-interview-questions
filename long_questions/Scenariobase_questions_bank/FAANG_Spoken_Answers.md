# FAANG-Level Scenario Answers (71-100) - Perfect Spoken Format

## 🔴 Distributed Systems & Scale Scenarios

### Question 71: A cache stampede brings down your database. How do you prevent it?
**What Perfect Spoken Format Looks Like:**
> "If a cache stampede is actively taking down the database, my immediate action would be to aggressively rate limit or shed traffic at the API Gateway level to give the database breathing room.
>
> To discover what is triggering it, I would monitor the cache hit rate closely to identify the specific keys or patterns experiencing mass expirations simultaneously while under heavy traffic.
>
> To permanently resolve this, I would implement a locking mechanism like a distributed mutex or a 'singleflight' pattern where only the first request fetching an expired key is allowed to query the database, while the rest are blocked until the cache is populated.
>
> For long-term prevention, I would use 'Probabilistic Early Expiration', where a background worker refreshes popular keys slightly before their TTL expires, completely decoupling the datastore query from user requests and preventing stampedes entirely."

### Question 72: A leader node crashes during job processing. What happens next?
**What Perfect Spoken Format Looks Like:**
> "If a leader crashes, my immediate assumption is that the system will experience a brief pause in write operations. I would monitor our alerting rules to ensure a new leader election is triggered within our expected timeout bounds.
>
> To investigate the state of the disrupted job, I would check the follower nodes to see if the crashed leader’s last lease heartbeat expired successfully and whether the job state had been persisted completely before the crash.
>
> To resolve the disruption, the system’s consensus algorithm, like Raft or Paxos, will naturally elect a new leader. The interrupted job must be resumed by the new leader reading the last known checkpoint from a Write-Ahead Log to process any remaining events.
>
> To prevent any data corruption, especially if the old leader experiences a 'zombie' revival, I would ensure we use fencing tokens. The old leader’s token would be invalidated so that any lingering write attempts it tries to make are rejected by the data store."

### Question 73: Network partition splits your cluster. How does your system behave?
**What Perfect Spoken Format Looks Like:**
> "In the event of a network partition, my immediate priority is understanding our system’s CAP theorem stance. I would quickly verify whether the system is configured to prioritize Consistency (CP) or Availability (AP) so I know what behavior to expect.
>
> I would investigate the cluster health metrics to identify which nodes reside in the minority partition and which are in the majority quorum.
>
> The resolution depends entirely on the architecture. If we are running a CP system like ZooKeeper, the minority partition will automatically stop accepting writes to avoid Split-Brain. If we run an AP system like Cassandra, both sides will continue accepting writes to remain available.
>
> To prevent data inconsistencies once the network heals in an AP system, I would implement robust conflict resolution mechanisms, such as Vector Clocks or Last-Write-Wins timestamps, and ensure read-repairs are properly configured."

### Question 74: Hot keys overload a single cache node. How do you fix it?
**What Perfect Spoken Format Looks Like:**
> "If a single cache node starts getting overloaded or even failing due to a hot key, my immediate action would be to restart the node if it crashed and rely on circuit breakers to momentarily redirect traffic.
>
> To investigate, I would look at cache access logs and metrics to identify the exact hot key or group of keys skewing the traffic distribution so heavily.
>
> To fix the issue, I would add a local, in-memory 'L1' cache on the application servers themselves for those highly requested items. If that isn’t feasible, I would implement key splitting strategies where we replicate the data across multiple keys (like key_1, key_2, key_3) and have the application randomly read from any of them to distribute the load across the cluster.
>
> To prevent this structurally, I would build automated hot key detection into the application layer to dynamically shift popular items into a secondary local cache before they can ever threaten a single centralized caching node."

### Question 75: A global service shows higher latency in one region. How do you debug?
**What Perfect Spoken Format Looks Like:**
> "My immediate response to a regional latency spike would be to route global traffic temporarily away from the impacted region, falling back to a healthy adjacent region using Geo-DNS if our SLAs dictate strict latency requirements.
>
> I would begin triangulating by checking if our users are being correctly routed or if there’s a misconfigured trans-oceanic cable routing causing packet loss. I would then check the local dependencies in that region to see if a database replica or localized API is struggling.
>
> If the problem stems from a specific dependency, I would scale it up or roll back any regional canary deployments that could have introduced the bottleneck. If it's pure network latency, I might reach out to the ISP or CDN provider.
>
> To completely prevent this type of unseen degradation, I would implement rigorous regional synthetic monitoring 'Blackbox' tests, which mimic real user requests from around the world to immediately alert us of route or regional degradation long before full traffic drops."

### Question 76: Event consumers fall behind producers. What do you do?
**What Perfect Spoken Format Looks Like:**
> "If consumers are severely lagging behind producers, my immediate action would be to dynamically scale up the number of consumer instances to aggressively chew through the backlog, up to the maximum limit of our partitions.
>
> To investigate the cause of the lag, I would check consumer metrics to see if the consumers themselves are bottlenecked on CPU, memory, or, as is most common, blocking I/O calls to a slow database or external API.
>
> If scaling up consumers doesn’t outpace the ingestion, I would apply backpressure on the producers to temporarily rate-limit or throttle the ingestion of non-critical events. For critical events, I might split the workload by moving oversized or slow-to-process messages into a separate 'Slow Lane' topic.
>
> To prevent future lag, I would continually audit our Kafka partition counts and consumer group sizes, ensuring we always have enough granular partitions to horizontally scale up parallel consumers on demand before the queue ever builds."

### Question 77: Duplicate messages appear in a queue. How do you handle idempotency?
**What Perfect Spoken Format Looks Like:**
> "If duplicate messages are discovered causing multiple executions, my immediate action is to pause the downstream workers if the operations are highly destructive, such as duplicated financial transactions.
>
> I would investigate by tracing the message lifecycle. In distributed systems, retries from producers due to network timeouts are almost guaranteed to create 'At-Least-Once' delivery duplicates.
>
> To resolve this gracefully, I would ensure every single message generated includes a universally unique Idempotency Key (like a UUID). The consumer must check a fast store like Redis or the database to see if that key has already been successfully processed before taking any action.
>
> For bulletproof prevention, I would wrap the deduplication check and the actual state change in a single atomic database transaction using mechanisms like conditional `INSERT IGNORE` or checking the status before updating, fully transitioning the system to process idempotently by default."

### Question 78: Exactly-once processing is required. How do you design for failures?
**What Perfect Spoken Format Looks Like:**
> "First, I'd acknowledge that theoretical 'Exactly-Once delivery' over an unreliable network is impossible due to the Two Generals problem. My immediate action is to align with the product owner that we will simulate exactly-once semantics through idempotent processing instead.
>
> In evaluating the architecture, I'd trace where an event risks being dropped or duplicated—specifically between the producer acknowledging an event and the consumer reacting and committing it.
>
> The safest resolution is the Transactional Outbox pattern. We use a single, atomic database transaction to update our application's state and insert the event payload into a local 'outbox' table. A background publisher then safely reads the outbox table to guarantee the message is pushed to the queue.
>
> To prevent duplication on the consumer side, I would combine this with strict idempotent consumers checking UUIDs, or I would natively utilize Kafka’s transactional producer capabilities with idempotency enabled (`enable.idempotence=true`) if we are deeply embedded in the Kafka ecosystem."

### Question 79: Redis goes down in a critical path. How does your system recover?
**What Perfect Spoken Format Looks Like:**
> "If Redis completely drops out of a critical path, my immediate physical action is to stop sending traffic to the dead nodes. The system must immediately trigger its circuit breakers to halt all outbound requests to Redis to prevent our application threads from hanging and exhausting our pools.
>
> I would check our Redis Sentinel or Cluster dashboards to determine if this is an isolated node failure awaiting automatic replica promotion or a widespread cluster outage.
>
> To resolve user impact instantly, the application should automatically fall back. Depending on the load, we either bypass Redis and query the primary database—if it’s sized to handle the traffic—or we serve stale data from a local, in-memory fallback cache like Guava or Caffeine.
>
> To prevent systemic failure in the future, I must ensure that caching is always treated as an ephemeral optimization, never a strict dependency, gracefully degrading functionality rather than taking down the core service."

### Question 80: Distributed lock causes throughput drop. How do you redesign?
**What Perfect Spoken Format Looks Like:**
> "If a distributed lock is causing massive contention and dropping our throughput, my immediate action would be to evaluate if we can temporarily disable the feature causing the contention if it’s non-critical, or aggressively reduce the lock acquisition timeouts.
>
> I would investigate the locking granularity. Often, services mistakenly lock entire tables or broadly defined logical objects when multiple threads are actually trying to write to completely independent rows.
>
> To resolve this, I would migrate us away from pessimistic locking via Redis or Zookeeper and switch to Optimistic Concurrency Control. By using a 'version' column in the database, we allow multiple threads to attempt updates simultaneously—with `UPDATE table SET val=new_val, version=version+1 WHERE id=x AND version=old_version`. Only conflicting updates are rejected and forced to retry.
>
> For long-term prevention, I rigorously review architectures to ensure we use fine-grained, row-level locks or optimistic locking loops, avoiding centralized lock managers entirely unless absolutely necessary to scale throughput effectively."

### Question 81: Schema change breaks older services. How do you deploy safely?
**What Perfect Spoken Format Looks Like:**
> "If a schema change actively breaks an older rolling deployment, my immediate action is to halt the deployment and roll back the database schema if—and only if—it is safe and we won’t lose newly written user data. If rollback isn’t safe, I have to fast-forward a patch to the older services.
>
> I'd investigate our deployment logs to determine where the backward compatibility contract was violated.
>
> To ensure this never happens again, I strictly enforce the 'Expand and Contract' pattern for all schema evolutions. Phase 1 expands the schema by adding the new column while the code writes to both old and new. Phase 2 backfills historical data. Phase 3 switches all reads to the new column.
>
> To prevent regressions permanently, Phase 4 finally contracts the database by safely dropping the old column long after the old service instances are entirely offline. This pattern safely decouples the database migration from the application deployment lifecycle."

### Question 82: Thundering herd during cache warm-up. How do you solve it?
**What Perfect Spoken Format Looks Like:**
> "If reviving a cold cache instantly triggers a thundering herd that threatens our database, my immediate action is to dramatically throttle incoming traffic at the edge to buy the database time to recover.
>
> I’d analyze the cache miss logs to confirm that thousands of concurrent requests are simultaneously seeing cache misses and bombarding the database for the exact same set of resources.
>
> To resolve the active situation and build a resilient pipeline, I would introduce request coalescing. If 10,000 requests all want the same missing key, only the first request is allowed to fetch from the DB, while the others pause and wait for the result.
>
> For bulletproof prevention moving forward, I would require explicit 'Cache Pre-warming' scripts to run as part of our deployment pipeline. You seamlessly populate the cache layer prior to routing any live traffic to the newly spun-up regions or shards."

### Question 83: A background reprocessing job overloads production. How do you control it?
**What Perfect Spoken Format Looks Like:**
> "If a heavy background job is thrashing the production database and impacting live users, my immediate action is to pause or kill the background job to restore primary service SLAs.
>
> I would check database metrics to confirm if the job is monopolizing disk I/O, locking tables excessively, or exhausting the connection pools.
>
> To resolve this, I would decouple the job from the primary read/write nodes. If the job only needs to read data, I would reroute its connection string entirely to an Async Read Replica database.
>
> To prevent this permanently, I would enforce strict Rate Limiting and Quality of Service prioritization on all background workers. Background tasks should dynamically sense the latency of the primary database and implement backpressure—automatically pausing or slowing down their own throughput when real user traffic spikes."

### Question 84: Partial failures cause cascading outages. How do you stop them?
**What Perfect Spoken Format Looks Like:**
> "If a partial, localized failure is actively bringing down healthy services in a cascading pattern, my immediate action is to physically open circuit breakers to the failing dependency or sever the network route immediately to protect the larger ecosystem.
>
> I would investigate where the resource exhaustion is happening—cascading failures usually mean healthy services are hanging indefinitely on synchronous network calls, chewing up all available threads or memory waiting for a response that will never come.
>
> To resolve the bottleneck, I would ensure strict, aggressive timeouts are enforced on every single cross-service network call in the architecture, allowing threads to fail fast and return rather than block.
>
> To prevent this systemically, I would implement the Bulkhead Pattern. By isolating thread pools—giving Service A a strict maximum of 10 threads out of 100 to communicate with Service B—if Service B crashes, Service A only loses those 10 threads, isolating the blast radius and preventing a total system collapse."

### Question 85: A single shard becomes overloaded. How do you rebalance?
**What Perfect Spoken Format Looks Like:**
> "If a single database shard is overheating under massive request volume, my immediate action is to scale up the compute and memory resources vertically on that specific node if the cloud provider allows it quickly. 
>
> I would investigate the query distribution. I need to know if this is a genuinely unbalanced hash distribution across thousands of users, or if one specific 'whale' tenant is monopolizing the traffic.
>
> If it's a structural imbalance, my resolution is to utilize Consistent Hashing with virtual nodes uniformly mapping keys to shards, and gracefully triggering a migration script to split the hot shard in two and rebalance the data to a new instance.
>
> To prevent this long-term, especially in multi-tenant systems, if the issue is a single whale tenant, I would utilize 'Tenant Isolation', moving that massive customer to their own dedicated hardware entirely to protect the noisy neighbor ecosystem of smaller tenants."

### Question 86: You need to debug a bug that happens once a day at scale. How?
**What Perfect Spoken Format Looks Like:**
> "If there is an obscure bug happening only once a day across millions of requests, my immediate priority is to stop attempting standard manual reproduction and instead heavily bias towards supreme observability.
>
> I'd first investigate the metrics to triangulate patterns. Does this align perfectly with a midnight cron job? Does it align with a specific payload size, a specific GC pause, or a specific geographic region? 
>
> To dynamically capture it natively, I would massively increase our distributed tracing sampling rate specifically around the suspected time window, and ensure all logs carry universally unique Correlation IDs connecting the edge proxy all the way down to the database.
>
> For future prevention and real-time capture, if the standard logs aren't enough, I would isolate a small percentage of anomalous traffic—using a canary via a load balancer—and route it to a single dedicated node running with heightened APM profilers attached, ready to automatically capture the thread dump the second the anomaly hits."

### Question 87: A rolling deployment causes inconsistent reads. Why?
**What Perfect Spoken Format Looks Like:**
> "If users are experiencing inconsistent reads or bizarre UI states mid-deployment, my immediate action is to pause the rolling wave to halt the bleeding, or accelerate the completion if pausing leaves us totally broken.
>
> I’d heavily investigate the payload contracts. This happens when a user’s initial request is served by a 'V2' node writing a newly formatted object to the cache, but their subsequent rapid request hits a 'V1' node that doesn't understand the cache structure and fat-fingers the response.
>
> To resolve it for the active sessions, if it's severe, we must invalidate the poisoned cache entries or temporarily enable Sticky Sessions on the load balancer to bind a user to a specific node version.
>
> To strictly prevent this in the future, I enforce strict Backward and Forward Compatibility rules. Any V1 service must be coded to gracefully ignore unknown JSON fields generated by a V2 service, ensuring cache contracts are completely decoupled from deployment waves."

### Question 88: Clock skew causes ordering issues. How do you fix?
**What Perfect Spoken Format Looks Like:**
> "If clock skew across our distributed nodes is scrambling the causal order of our events, my immediate action is to verify that Network Time Protocol (NTP) daemons are running and actively syncing time across all hosts. 
>
> I'd investigate the specific timestamps in conflict. Relying on physical wall clocks in distributed systems guarantees race conditions because 'Time' is relativistic between two independent physical CPU crystals.
>
> To resolve this permanently, we must transition the architecture. We cannot rely on physical time for event ordering. We must use Logical Clocks—specifically Lamport Timestamps or Vector Clocks—which increment a sequential counter strictly based on causal dependencies between services.
>
> For supreme prevention in a globally distributed database, I would utilize technologies built on true bounded-time APIs—like Google Spanner's TrueTime, which uses atomic clocks to calculate explicit, guaranteed upper-bound uncertainties before committing writes."

### Question 89: A metrics system lies during outages. How do you validate data?
**What Perfect Spoken Format Looks Like:**
> "If our APM dashboard shows 100% success but users are reporting an outage on Twitter, my immediate action is to bypass the aggregated dashboards completely and tail the raw load balancer and edge proxy logs to see exactly what HTTP status codes users are actually receiving.
>
> I would investigate why the metrics fell out of sync. This usually occurs because internally tracked application metrics fail when the application itself is too dead or deadlocked to emit the failure metrics, or the metrics aggregator itself drops events during extreme traffic spikes.
>
> To resolve the visibility gap instantly, I cross-check internal latency against black-box external probes—like Pingdom or Route53 health checks—to see what the outside world truthfully experiences.
>
> To prevent this permanently, our alerting strategy must be fundamentally built on Edge Metrics and External SLIs. We measure reliability at the Load Balancer level—because it will faithfully log a 502 Bad Gateway even if the application servers are completely turned off."

### Question 90: Distributed tracing adds overhead. How do you balance visibility vs performance?
**What Perfect Spoken Format Looks Like:**
> "If implementing distributed tracing is crippling our application's CPU or network bandwidth, my immediate action is to drastically slash the sampling rate dynamically via our control plane down to 0.1% or 1% of total traffic.
>
> I’d investigate our APM configuration to see if we are eagerly transmitting gigantic payload bodies in our spans or if we are emitting spans too deep within extremely tight 'hot' computational loops.
>
> To resolve this gracefully, I would switch our strategy from 'Head-based sampling'—where we randomly decide to trace a request at the API gateway—to 'Tail-based sampling'. 
>
> For long-term prevention, Tail-based sampling allows the application to buffer traces in memory extremely cheaply, only committing to the expensive network serialization step and sending the trace to the collector if the transaction ultimately fails or violates a latency threshold. This guarantees we always capture 100% of errors without paying the full cost of tracing successes at scale."

### Question 91: Leader election flaps frequently. What are the consequences?
**What Perfect Spoken Format Looks Like:**
> "If a cluster is rapidly flipping its leadership role, my immediate action is to manually intervene, potentially hard-pinning a leader or relaxing the election timeout rules to restore immediate cluster stability.
>
> I'd investigate the health metrics and networking logs between the nodes. The primary consequence of leader flapping is intense write-unavailability, because every time a leader dies, the cluster pauses writes to hold an election, severely impacting our availability SLA.
>
> To resolve this, I would check if the network connection between our consensus nodes (like Zookeeper or Etcd) is unstable, or if the current leader is suffering heavy garbage collection (GC) pauses causing it to miss consecutive heartbeat acknowledgments.
>
> For strict prevention, I would dynamically tune the heartbeat intervals and election timeouts to be slightly more tolerant of transient network jitter or minor GC pauses, ensuring a leader is only deposed if it is genuinely dead, not just momentarily sluggish."

### Question 92: A service must degrade gracefully under overload. How?
**What Perfect Spoken Format Looks Like:**
> "If a service is actively crumbling under an enormous spike of real user traffic, my immediate priority is to activate Load Shedding at the inbound API Gateway, mercilessly dropping low-priority requests like background analytics to preserve core functionality.
>
> My triangulation involves looking at telemetry to determine exactly which subsystem is choking—is it the DB, CPU, or an external API pool—so I know which parts of the service must be turned down.
>
> To resolve the load dynamically, I would rely on heavily pre-configured Feature Toggles. During extreme load, we automatically toggle off heavy, computationally expensive components—like 'Recommended Products' on an e-commerce site—to guarantee the checkout workflow survives.
>
> To prevent system death long-term, our architecture must be designed to serve stale data. If the primary database cannot handle read queries during an extreme event, the application automatically bypasses the DB and happily serves mildly stale, highly cached data from Redis or a CDN, fulfilling the read request while the system recovers."

### Question 93: Data loss is reported in an eventually consistent system. How do you investigate?
**What Perfect Spoken Format Looks Like:**
> "If a user actively reports missing data in an eventually consistent cluster like Cassandra, my immediate action is to instruct support not to panic, as 'Data Loss' might just be severe 'Replication Delay'. I'd check the replication lag metrics.
>
> I would triangulate by validating the user's reads against the primary nodes. Did the user write to Node A, but instantly read from a lagging replica on Node B? Or did a genuine conflict resolution protocol silently overwrite their data?
>
> To resolve the confusion, I would check our consistency settings. If multiple nodes received concurrent writes for the same row, a naive 'Last Write Wins' timestamp policy might have discarded the user's data due to clock skew between instances.
>
> To prevent actual data loss and guarantee safety, we must rigorously audit our read/write Quorums. We ensure that our `Write_Consistency_Level` plus `Read_Consistency_Level` is greater than the total number of Replicas. By doing so, we guarantee strict overlap, meaning users will always get strong consistency when required, trading a fraction of availability for perfect read safety."

### Question 94: A retry storm worsens an outage. How do you design retries?
**What Perfect Spoken Format Looks Like:**
> "If we are actively caught in a retry storm, I would immediately open the circuit breakers manually or drop all incoming retried requests at the load balancer to protect the recovering downstream service.
>
> To investigate, I'd trace the logs to see if a massive block of parallel workers hit a minor timeout and subsequently hammered the target API in perfect unison a second later.
>
> To resolve the system design flaw natively, I would strictly adhere to three principles for any network call. First, I'd implement Exponential Backoff—instead of retrying instantly, the server waits 1 second, then 2, 4, and 8 seconds. Second, and most critical, I add Jitter—a randomized delay mathematical curve so the retries are desynchronized smoothly. 
>
> Finally, for robust prevention, I would implement a 'retry budget'. The service should track the ratio of retries to normal requests. If retries exceed 10% or 20% of our general traffic pool, the client realizes the backend is hopelessly broken and stops retrying entirely to fail fast."

### Question 95: Global rate limiting is needed. How do you implement it?
**What Perfect Spoken Format Looks Like:**
> "If we are being actively swamped and need to implement global rate limiting immediately, my first move is to deploy IP-based rate limiting rules aggressively at the edge using our WAF or Cloudflare.
>
> As I investigate a permanent architectural solution, I'd realize a purely in-memory rate limiter per application node fails because traffic splits unevenly across the cluster; we need a centralized state.
>
> For resolution, I implement a high-performance centralized store, primarily a clustered Redis instance, utilizing the Token Bucket algorithmic pattern utilizing incredibly fast Lua scripts that enforce atomic limits per generic identifier like UserID or API Key.
>
> To prevent Redis itself from becoming a latency bottleneck under immense scale, I would design a hybrid batching system. Local nodes calculate limits in memory temporarily and sync their consumed tokens to the central Redis cache asynchronously every few seconds, guaranteeing massive throughput with extreme global accuracy."

### Question 96: Backpressure is missing and causes crashes. How do you add it?
**What Perfect Spoken Format Looks Like:**
> "If an application is consuming unbounded data and crashing from Out-Of-Memory errors or thread exhaustion, my immediate action is to aggressively throttle the upstream ingestion endpoints or pause Kafka topic consumption.
>
> I’d investigate the thread pools and internal queues within our Java or Go application, generally finding unbounded asynchronous queues allowing memory to inflate continuously because producers outpace consumers.
>
> My primary resolution is implementing bounded queues, like Java’s `ArrayBlockingQueue`. Once the queue hits capacity, any new inbound requests are strictly rejected with a 429 'Too Many Requests' or 503 error, forcing the client to back off rather than crashing our JVM.
>
> For rigorous future prevention in streaming ecosystems, I would fundamentally build the application atop Reactive Streams principles using libraries like Project Reactor or RxJava, allowing consumers to explicitly signal to producers precisely how many items they can handle, baking mechanical backpressure straight into the protocol."

### Question 97: Multi-tenant service sees noisy neighbor issues. How do you isolate?
**What Perfect Spoken Format Looks Like:**
> "If a specific set of users are seeing their latency degrade because of a noisy neighbor on a shared cluster, my immediate priority is to aggressively throttle or temporarily rate-limit the offending 'whale' tenant to restore the SLA for the majority of users.
>
> I would investigate our telemetry to identify exactly how the heavy tenant is monopolizing resources—are they hoarding database I/O, consuming all available worker threads, or thrashing the shared queues?
>
> Assuming pure code-level quotas aren’t tight enough, my primary resolution is Shuffle Sharding. Instead of all tenants sharing all hardware, I deterministically assign tenants to overlapping subsets of infrastructure nodes. If one node goes down because of a whale, it only statistically impacts a microscopic fraction of other tenants.
>
> For the ultimate long-term preventative measure against hyper-massive accounts, I utilize a 'Dedicated Pool' architectural pattern. Once a tenant breaches a certain tier of traffic, the system dynamically migrates them onto physically isolated, provisioned compute clusters, guaranteeing absolute hardware separation."

### Question 98: Blue-green deployment causes traffic imbalance. How do you fix?
**What Perfect Spoken Format Looks Like:**
> "If we cut over a Blue-Green deployment and surprisingly see traffic perfectly split across both fleets, my immediate action is to halt attempting to shut down the old environment to prevent massive CPU exhaustion, as it's clearly still processing heavy load.
>
> I’d investigate our network routing layer. Time and time again, if the swap was executed purely via DNS updates, the issue is DNS caching—clients or ISPs are heavily caching the old IP address and ignoring the new TTL.
>
> To resolve the active situation, especially for long-lived systems like WebSockets or gRPC streams, users remain glued to the old servers via Keepalive connections. I would gracefully force connection closures from the server side on the old fleet, forcing all clients to re-resolve DNS and reconnect to the new Green fleet.
>
> For rigorous prevention moving forward, we never rely on DNS switches for immediate traffic shifting. I would orchestrate the Blue-Green swap strictly utilizing Weighted Target Groups directly on the Application Load Balancer, ensuring an instantaneous, router-level state change immune to client-side caching."

### Question 99: An SLO is violated intermittently. How do you root-cause it?
**What Perfect Spoken Format Looks Like:**
> "Intermittent SLO violations are notoriously tricky. My immediate action isn't panic since the system is technically functional; rather, it’s preserving forensics. I ensure logs and metrics spanning the violation windows are retained warmly and aggressively audited.
>
> I investigate by utilizing visual heatmaps instead of simple averages. A standard dashboard showing a 50ms average hides the fact that 5% of requests spiked perfectly to 3 seconds. P99 and P99.9 metrics are critical here.
>
> To resolve it, I triangulate the exact timestamp against external variables. Is there a massive nightly snapshot backup locking tables? Does garbage collection sweep at that exact moment? Or is it a subset of complex API queries from just one large customer?
>
> For strict prevention, once isolated, we implement High-Cardinality distributed tracing. We configure our metrics architecture to automatically capture and retain the full call stack and database queries associated with the specific outliers breaching our SLO latency threshold, ensuring the anomaly is vividly documented the next time it fires."

### Question 100: You’re on call and multiple alerts fire at once. How do you prioritize?
**What Perfect Spoken Format Looks Like:**
> "If the alerts board lights up brightly with dozens of simultaneous pages, my absolute immediate action is to take a deep breath, remain perfectly calm, and officially log into the incident channel to declare myself the Incident Commander to centralize communication.
>
> My core investigation methodology is purely based on the Blast Radius and Customer Impact. I ruthlessly ignore CPU alerts or internal tool warnings and prioritize immediately triaging: Are primary databases down? Are customers seeing 500 errors? Are financial transactions failing?
>
> In resolving the cascade, my focus is immediately identifying the 'Source' versus the 'Symptom'. Fifty APIs throwing timeout alerts are usually just symptoms of the central PostgreSQL database being unresponsive. I zero in on the root dependency rapidly and mitigate it—likely via automated rollback, failover, or scaling rather than deep debugging code.
>
> To prevent future alert fatigue, post-incident, I perform an aggressive purge on our alerting thresholds. We must rely heavily on Symptom-Based Alerting, meaning pagers only wake engineers up for breached user SLAs, not simply because a worker node ran high on memory for ten seconds."
