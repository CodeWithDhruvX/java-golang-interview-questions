## 🟢 Evolution, Resilience, Privacy (Questions 951-1000)

### Question 951: How do you use codemods to upgrade Node codebases safely?

**Answer:**
**Codemods** are scripts (using AST) that refactor code programmatically.
Use `jscodeshift`.
Example: Converting `var` to `const`, or updating old library imports (`require` -> `import`).
Ensures consistency across 1000s of files.

---

### Question 952: How do you automatically refactor deprecated API usages?

**Answer:**
Write a jscodeshift transform.
Identify deprecated pattern (e.g., `fs.exists`).
Replace with new pattern (`fs.stat` / `fs.access`).
Run script.

---

### Question 953: How do you enforce upgrade paths for Node major versions?

**Answer:**
CI check `engines` field.
Use `nvm` `.nvmrc` to pin version.
Upgrade process:
1.  Read Changelog.
2.  Run Deprecation handling (`--trace-warnings`).
3.  Update CI.

---

### Question 954: How do you evolve data models across live running Node2 Node nodes?

**Answer:**
(Microservices).
**tolerant reader** pattern.
Service A sends New Field. Service B (Old) ignores it.
Service B stops sending Old Field. Service A handles missing field (default value).

---

### Question 955: How do you guarantee schema consistency during rolling migrations?

**Answer:**
Expand-Contract strategy.
1.  Add new column (nullable).
2.  Code writes to both.
3.  Backfill.
4.  Code reads new.
5.  Delete old.

---

### Question 956: How do you freeze time in Node to debug scheduled tasks?

**Answer:**
`sinon.useFakeTimers()`.
`clock.tick(24 * 60 * 60 * 1000)`. // Fast forward 1 day.

---

### Question 957: How do you simulate network partitions for testing resilience?

**Answer:**
Use specialized proxies (`Toxiproxy`).
Configure it to cut connection between App and DB.
Verify App reconnects logic.

---

### Question 958: How do you replay historical production traffic in test environments?

**Answer:**
Traffic Shadowing (Mirroring).
Nginx / Envoy copies incoming request. Sends to Prod (Real) and Staging (Shadow).
Compare responses (Diffy).

---

### Question 959: How do you checkpoint state in-memory to rewind runtime state?

**Answer:**
State Pattern / Memento.
Serialize state object to JSON. Push to stack.
Undo: Pop stack, deserialize.

---

### Question 960: How do you emulate partial shearing or thread preemption manually?

**Answer:**
Inject `await delay(0)` at random points to force event loop yields.
Reveals race conditions rooted in assumption of atomicity.

---

### Question 961: How do you define and manage infra via a Node app (e.g. Pulumi program)?

**Answer:**
Pulumi allows defining resources in TS.
`const group = new aws.ec2.SecurityGroup(...)`.
Run `pulumi up` to deploy.

---

### Question 962: How do you orchestrate blue/green deployment within a Node process?

**Answer:**
Internal load balancing.
Start new logic.
Route 10% calls to new logic function.
Monitor.
Switch 100%.

---

### Question 963: How do you dynamically wire services at runtime via a declarative model?

**Answer:**
Dependency Injection Container (InversifyJS).
Read config.
Bind `IService` to `MockService` or `RealService`.

---

### Question 964: How do you trigger infrastructure changes via business‑logic code in Node?

**Answer:**
Call Cloud SDK (AWS SDK).
`ec2.launchInstances(...)` or `lambda.updateFunctionCode(...)`.
Potentially dangerous (Security risk).

---

### Question 965: How do you rollback infra as easily as application code?

**Answer:**
If Infra is Code (Terraform/Pulumi), `git revert` the commit.
Run pipeline.

---

### Question 966: How do you run on-device ML inference in Node within constrained environments?

**Answer:**
Quantization (int8 instead of float32).
Reduces model size and CPU usage.
Use `TensorFlow Lite`.

---

### Question 967: How do you offload AI compute to specialized edge hardware?

**Answer:**
(Coral TPU / Jetson).
Use specialized Node bindings for that hardware (`edgetpu-node`).

---

### Question 968: How do you manage model updates and versions across devices?

**Answer:**
IoT Management.
Device checks manifest.
Downloads new `.tflite` file.
Hot-swaps model in memory.

---

### Question 969: How do you securely fetch new model binaries at runtime?

**Answer:**
Signed URLs (S3).
Verify Checksum/Signature of downloaded model before loading.

---

### Question 970: How do you certify and audit inference outputs across deployments?

**Answer:**
Log Input Hash + Output + Model Version.
Sampling audit.

---

### Question 971: How do you use Temporal or Cadence with Node clients for workflows?

**Answer:**
**Temporal** provides a Node SDK.
Write workflow as code function.
Temporal server handles state persistence, retries, sleep (durability).

---

### Question 972: How do you coordinate multi-step saga patterns in Node services?

**Answer:**
Orchestrator approach.
Calling `Step1`. If success, `Step2`.
If `Step2` fail, call `Compensate1` (Undo Step1).

---

### Question 973: How do you simulate event replays and compensating logic?

**Answer:**
Store events in Event Store.
Reset "Read Pointer" to 0.
Reprocess events.
Ensure logic is **Idempotent**.

---

### Question 974: How do you safely store workflow states across crashes?

**Answer:**
Persist state to DB after every step.
Or use Temporal/Durable Functions which abstract this.

---

### Question 975: How do you manage versioning of workflows in Node runtime?

**Answer:**
If workflow code changes, existing running workflows might break (non-determinism).
**Patching:** Check `if (version < 2) runOldLogic() else runNewLogic()`.

---

### Question 976: How do you run edge-aware health checks for Node services?

**Answer:**
Check local dependencies.
But also allow "Degraded" state (e.g., DB is down, but Cache is serving). Don't mark completely Unhealthy (which kills pod).

---

### Question 977: How do you aggregate and sample logs from edge nodes?

**Answer:**
Local aggregation (Vector/Fluentd sidecar).
Compress.
Send upstream in batches.

---

### Question 978: How do you expose fine‑grained metrics (e.g. per‑CPU core) in Node?

**Answer:**
Prometheus Exporter.
Custom metric `cpu_usage{core="0"}`.

---

### Question 979: How do you dynamically adjust health TTLs based on workload?

**Answer:**
If load is high, increase health check interval/timeout to avoid false positives (killing busy nodes).

---

### Question 980: How do you handle partial network failures gracefully in health probes?

**Answer:**
Retries.
And Soft Failure (Mark as warning, don't remove from load balancer immediately).

---

### Question 981: How do you auto‑generate tests using GPT‑like models from spec?

**Answer:**
Feed OpenAPI spec/TypeDefs to LLM.
Prompt: "Generate Jest tests covering 100% of this schema".
Review manually.

---

### Question 982: How do you integrate mutation testing into Node pipelines?

**Answer:**
**Stryker** Mutator.
It modifies your code (changes `+` to `-`). runs tests.
If tests pass (survive), your tests are weak.
Goal: Kill all mutants.

---

### Question 983: What are the risks of auto‑generated code tests?

**Answer:**
False sense of security.
Tests might just assert the current buggy behavior (Snapshot testing bad habits).

---

### Question 984: How do you validate coverage from AI‑generated test cases?

**Answer:**
Run coverage tool (Istanbul/C8).
Ensure logical branches are hit.

---

### Question 985: How do you maintain human oversight in AI test suites?

**Answer:**
Code Review the tests.
Treat tests as first-class code.

---

### Question 986: How do you capture and respond to POSIX signals (`SIGUSR1`, etc.) in Node?

**Answer:**
`process.on('SIGUSR1', () => { ... })`.
Commonly used to reload config or toggle debug mode.

---

### Question 987: How do you implement custom signal handlers in a multi-threaded Node app?

**Answer:**
Signals are sent to the Main Process.
Main process catches signal and `postMessage` to workers telling them to act (e.g., shutdown).

---

### Question 988: How do you restore state safely after an interrupt-driven crash?

**Answer:**
Crash-only software design.
On restart, check Journal/Write-Ahead-Log.
Replay uncommitted ops.

---

### Question 989: How do you coordinate signal-handling across clustered processes?

**Answer:**
(See Q987). Master receives signal -> Forwards to workers.

---

### Question 990: How do you trigger runtime behaviors (e.g. dump stats) via signals?

**Answer:**
`process.on('SIGUSR2', dumpStats)`.
Admin runs `kill -SIGUSR2 <PID>`.

---

### Question 991: How do you test compliance with JSON-RPC protocol in Node?

**Answer:**
Use validation library handling JSON-RPC 2.0 spec (id, method, params, result/error).
Verify error codes (-32700 Parse error, etc.).

---

### Question 992: How do you validate GraphQL server behavior matching the spec?

**Answer:**
`graphql-js` reference implementation ensures spec compliance.
Corner cases: Null bubbling, introspection queries.

---

### Question 993: How do you run WebSocket conformance test suites in Node?

**Answer:**
**Autobahn|Testsuite**.
Connects to your WS server and runs 100s of fuzz tests (frame splitting, invalid utf8, ping/pong).

---

### Question 994: How do you check `Multipart/form-data` correctness under large file stress?

**Answer:**
Inspect boundary markers.
Check memory usage during upload (Stream vs RAM).

---

### Question 995: How do you ensure `HTTP/3` support works correctly in Node behind proxies?

**Answer:**
Check `Alt-Svc` headers.
Verify QUIC protocol handshake (UDP).

---

### Question 996: How do you enforce PII minimization in Node logs and analytics?

**Answer:**
Automated Redaction.
Type-safe logging where `User` object automatically calls `.toLog()` which excludes PII.

---

### Question 997: How do you support GDPR-level data access / erasure requests?

**Answer:**
Architecture: "User Sharding" or "Encryption Per User" (crypto-shredding).
To delete: Delete user's encryption key. Data becomes garbage.

---

### Question 998: How do you store and rotate consent metadata securely?

**Answer:**
Audit Log (Immutable).
`{ userId, version, consent: true, timestamp }`.

---

### Question 999: How do you design APIs that respect user tracking preferences?

**Answer:**
Read `DNT` (Do Not Track) header or Global Privacy Control.
If set, disable analytics pixels injection.

---

### Question 1000: How do you build audit trails in Node for privacy compliance?

**Answer:**
Middleware intercepts non-GET requests.
Writes: `Who (User), What (Action), When, Data Changed` to write-only ledger (DB/S3).
Immutable.
