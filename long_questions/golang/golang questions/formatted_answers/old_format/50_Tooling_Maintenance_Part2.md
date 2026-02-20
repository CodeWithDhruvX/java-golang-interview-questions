# ⚒️ **981–1000: Tooling, Maintenance & Real-world Scenarios (Part 2)**

### 981. How do you refactor legacy Go code?
"1.  **Cover with Tests**. (Characterization Test).
2.  **Simplify**: Break huge functions into smaller ones.
3.  **Decouple**: Inject interfaces instead of structs.
I follow the 'Boy Scout Rule': Leave the code cleaner than I found it."

#### Indepth
**Parallel Change**. How to replace a core function `Old()` with `New()` safely?
1. Add `New()`.
2. Update callsites to use `New()`, but keep `Old()` available.
3. Once all callsites are migrated, remove `Old()`.
This "Expand-Contract" pattern avoids big bang rewrites and keeps the build green at all times.

---

### 982. How do you organize large-scale Go monorepos?
"**Bazel** or **Buck** build systems.
Or just standard Go Modules with **Go Workspaces** (`go.work`).
I prefer one `go.mod` per service to avoid dependency hell ('Diamond Dependency Problem')."

#### Indepth
**Workspace Mode**. Before Go 1.18, multi-module repos were painful (replace directives in go.mod). Now, creating a `go.work` file in the root (`use ./pkg/a; use ./pkg/b`) allows your editor (VSCode/gopls) to see all modules as one workspace, enabling "Jump to Definition" across module boundaries seamlessly.

---

### 983. How do you distribute Go binaries securely?
"Sign them (Cosign/GPG).
Generate **SBOM** (Software Bill of Materials).
Distribute via secure channels (Artifactory, GitHub Releases).
Users verify the checksum (`shasum -c`) before running."

#### Indepth
**Transparency Log**. Cosign pushes the signature to a public immutable log (Rekor). This means users don't need your public key. They can verify against the OIDC identity (e.g., "Signed by github actions workflow X"). This is Keyless Signing, reducing the risk of compromised private keys.

---

### 984. How do you maintain changelogs in Go projects?
"I use an automated tool like `git-chglog` or `release-please`.
It scans Git History.
Extracts PR titles.
Categorizes them (Features, Bugs).
Updates `CHANGELOG.md`."

#### Indepth
**Semantic Versioning**. Bumping versions is hard. `release-please` automates it. If commit message contains `feat!:` (breaking change), it bumps MAJOR. `feat:` bumps MINOR. `fix:` bumps PATCH. This enforces SemVer strictly based on git history, removing human error from versioning.

---

### 985. How do you rollback failed Go releases?
"If on K8s: `kubectl rollout undo deployment/myapp`.
It reverts to the previous ReplicaSet.
Since Go binaries are self-contained, rolling back is just restarting the old image version."

#### Indepth
**Database Rollbacks**. Code is easy to rollback (stateless). Data is hard. **Schema Compatibility**. Always make DB changes "Forward Compatible". "Add Column" is safe. "Rename Column" is unsafe. To rename: Add new column, Dual Write to both, Backfill old data, Switch reads to new, Stop writing old, Remove old. This takes 5 deployments.

---

### 986. How do you add performance regression testing?
"I store benchmark results (`v1.0: 500ops/s`).
CI runs benchmarks on PR.
If `current < previous * 0.9` (10% drop), CI fails.
`cobalt` or `benchstat` are tools for this comparison."

#### Indepth
**Noise**. CI environments are noisy (noisy neighbors). A 10% drop might be random. Run benchmarks `count=10` times and compare the *Distribution* (p-value), not just the mean. `benchstat` does this automatically: "Delta: ~5% (p=0.40)" -> Insignificant. "Delta: -20% (p=0.01)" -> Real Regression.

---

### 987. How do you build CLI-based installers in Go?
"The Go binary *is* the installer.
It bundles the artifacts (using `embed`).
Run `./install`.
It copies itself to `/usr/local/bin`.
It writes systemd unit files.
It starts the service."

#### Indepth
**Self-Update**. A CLI tool should be able to update itself. `myapp update`. It calls the GitHub Releases API, downloads the new binary (for correct OS/Arch), verifies hash, and replaces `os.Executable()`. Libraries like `go-update` or `equinox` handle the tricky parts (replacing a running binary on Windows).

---

### 988. How do you generate dashboards from Go metrics?
"Go -> Prometheus -> Grafana.
I import the 'Go Processes' dashboard (ID: 6671).
It shows Heap, Goroutines, GC Pauses out of the box.
Then I add custom panels for my business metrics (`orders_total`)."

#### Indepth
**SLO Dashboards**. Don't just graph "CPU Usage". Graph "User Happiness". SLO (Service Level Objective): "99.9% of requests < 200ms". Graph the **Error Budget Burn Rate**. "At this rate, we will violate our SLA in 4 hours". This alerts you to *real* problems, not just random CPU spikes.

---

### 989. How do you monitor file system changes in Go?
"I use `fsnotify/fsnotify`.
It wraps `inotify` (Linux), `FSEvents` (Mac).
`watcher.Add("/path/to/watch")`.
`case event := <-watcher.Events: ...`.
Essential for 'Hot Reload' tools or processing uploaded files."

#### Indepth
**Debouncing**. File systems are noisy. "Save File" might trigger 3 events (Create, Write, Chmod). If you rebuild on every event, you waste CPU. Implement a **Debouncer**: Wait 100ms after the first event. If no new events come, *then* trigger the action. This coalesces the burst into a single build.

---

### 990. How do you implement custom plugins in Go?
"HashiCorp Plugin System (gRPC).
My App starts the Plugin (subprocess).
Talks via gRPC.
This is safer than native `plugin` package and works on Windows."

#### Indepth
**WASM Plugins**. The future of plugins is **WebAssembly (Wasm)**. `wazero` lets you run Wasm modules *inside* your Go app with near-native speed. It's sandboxed (safe), cross-platform, and supports multiple languages (Rust/C++ plugins for Go app). It avoids the overhead of gRPC/Process per plugin.

---

### 991. How do you keep Go dependencies up to date?
"**Dependabot** or **Renovate**.
They scan `go.mod`.
Open PRs: 'Bump github.com/gin-gonic/gin from 1.7 to 1.8'.
I verify the Changelog and tests before merging."

#### Indepth
**Indirect Dependencies**. `go.mod` lists direct deps. `go.sum` lists everything. Vulnerabilities often appear in indirect deps. `go mod graph` shows the tree. Use `go mod why -m <pkg>` to find *who* is importing the vulnerable package so you can update the parent.

---

### 992. How do you audit Go packages for security issues?
"`govulncheck ./...`.
It connects to the Go Vulnerability Database.
It tells me if I'm *calling* vulnerable code, reducing false positives."

#### Indepth
**Symbol Tracing**. Standard scanners (Trivy, Snyk) just check version numbers. "You use lib v1.0, it has bug". `govulncheck` is smarter. It parses the AST/Binary. "You use lib v1.0, but you never call the `VulnerableFunc()`. You are Safe." This drastically reduces "Upgrade Fatigue" for developers.

---

### 993. How do you migrate Go modules across repos?
"Move the code.
Update `go.mod` module path.
Use `gofmt -w -r 'old/pkg -> new/pkg'` to rewrite imports in all source files.
It’s a bit manual but `gopls` can help."

#### Indepth
**Gomvpkg**. The `golang.org/x/tools/cmd/gomvpkg` tool automates this. `gomvpkg -from github.com/old/pkg -to github.com/new/pkg`. It moves the source files AND updates all import paths in the entire project, including `_test.go` files and internal references.

---

### 994. How do you conduct performance reviews for Go codebases?
"I look for:
1.  Unnecessary allocations in loops.
2.  Casting byte slices to strings repeatedly.
3.  Incorrect lock usage (holding locks during IO).
4.  Missing timeouts on Contexts.
I use the **profiler** to back up my intuition."

#### Indepth
**Branch Prediction**. Look for code that confuses the CPU branch predictor. Sorted data processes faster than unsorted data (if `data[i] > 128`). While micro-opt, avoiding "Branch-y" code in hot loops (using bitwise ops instead of `if`) is a hallmark of high-performance Go code.

---

### 995. How do you implement a concurrent token bucket rate limiter?
"I use a struct with `sync.Mutex`.
`tokens float64`.
`lastCheck time.Time`.
Refill: `now.Sub(lastCheck) * rate`.
Consume: `if tokens >= 1 { tokens--; ok }`.
This lazy refill is efficient and thread-safe."

#### Indepth
**Uber Ratelimit**. Implementing a *correct* lock-free rate limiter is hard. `uber-go/ratelimit` implements the "Leaky Bucket" as a virtual scheduler. It sleeps the caller to smooth out traffic `Take()`. This makes traffic "Equidistant" (spaced out) rather than bursty, which is friendlier to downstream services.

---

### 996. How do you implement the Saga Pattern for distributed transactions?
"See Q972.
I use Choreography (Events) or Orchestration (Central Coordinator).
Go is great for Orchestrators (Workflow Engines) due to concurrency."

#### Indepth
**Event Sourcing**. Sagas store state in DB. Event Sourcing stores state as a sequence of events. Go's strong typing (Structs) and Protobuf make it ideal for defining these events. Replaying the event stream restores the current state of the Saga. This provides perfect auditability of the distributed transaction.

---

### 997. What is the Go Memory Model and the "Happens-Before" relationship?
"It defines when a read `r` observes a write `w`.
'A send on a channel happens before the receive'.
'A lock unlock happens before the next lock'.
If these rules aren't met, there is **no guarantee** of visibility between goroutines (Race Condition)."

#### Indepth
**Benign Data Races**. There is no such thing. "I don't care if I read an old integer". In Go, a write to a multi-word value (interface, string, slice) is not atomic. A race can cause you to read the Pointer of the new slice but the Length of the old slice. This leads to segfaults and Arbitrary Code Execution. **All data races are bugs**.

---

### 998. How do you implement distributed tracing with context propagation?
"`otel.GetTextMapPropagator().Inject(ctx, carrier)`.
This puts `traceparent` header into the HTTP request.
The downstream service extracts it, continuing the trace."

#### Indepth
**W3C Headers**. The industry standard is now `Trace Context` (W3C). Header `traceparent: version-traceid-parentid-flags`. Go's `otel` library supports this by default. This ensures your Go microservice can participate in a trace that started in a Java frontend and ends in a Python AI worker.

---

### 999. How do you use the `slices` and `maps` packages (Go 1.21+)?
"Generic helper functions!
`slices.Contains(s, val)`.
`slices.Sort(s)`.
`maps.Clone(m)`.
No more writing `stringInSlice` helper functions. It’s a massive QoL improvement."

#### Indepth
**Custom Generics**. Don't stop at stdlib. Write your own generic data structures. `Set[T]`, `Tree[T]`, `Graph[T]`. Pre-1.18, these required `interface{}` and casting. Now you can implement a type-safe, zero-allocation `RingBuffer[T]` that outperforms the channel-based implementation for specific use cases.

---

### 1000. What is Profile Guided Optimization (PGO) and how do you use it?
"The compiler uses a real-world CPU profile to make optimization decisions (Inlining).
`go build -pgo=cpu.prof`.
It improves performance by ~5-10% with zero code changes."

#### Indepth
**AutoFDO**. Google uses "AutoFDO". Instead of manual instrumentation, they collect profiles from production continuously. The build system automatically fetches the "Last Week's Profile" for the service and builds the new binary with PGO. This creates a feedback loop where the code optimizes itself for its actual usage pattern over time.
