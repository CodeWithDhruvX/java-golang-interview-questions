# ⚒️ Go Theory Questions: 981–1000 Tooling, Maintenance & Real-world Scenarios

## 981. How do you refactor legacy Go code?

**Answer:**
**Test-Driven Refactoring**.
1.  **Pin it**: Write a high-level Integration Test (Golden File) capturing current output.
2.  **Lint**: Fix `golint` / `staticcheck` issues first.
3.  **Extract**: Move large functions to new packages (`internal/`).
4.  **Rename**: `gopls` rename to make code readable.
5.  **Verify**: Ensure Golden Test still passes.

---

## 982. How do you organize large-scale Go monorepos?

**Answer:**
We use **Go Workspaces** (`go.work`).
Structure:
```text
/
  go.work
  pkg/ (Public Shared)
  internal/ (Private Shared)
  services/
    payment/ (go.mod)
    auth/ (go.mod)
```
This keeps services decoupled (different `go.mod` deps) but allows atomic commits across the repo.

---

## 983. How do you distribute Go binaries securely?

**Answer:**
1.  **CheckSums**: Provide `SHA256SUMS.txt`.
2.  **Signatures**: Sign binary with GPG or **Cosign** (Sigstore).
3.  **HTTPS**: Serve only over TLS.
4.  **Reproducible Builds**: `go build -trimpath`. Users can rebuild from source and get the exact same byte-for-byte hash to verify no distinct backdoor was added by the compiler.

---

## 984. How do you maintain changelogs in Go projects?

**Answer:**
**Conventional Commits** (`feat:`, `fix:`).
Automation: `git-chglog` or `goreleaser` generates `CHANGELOG.md`.
It groups:
- Features
- Bug Fixes
- Breaking Changes
This avoids manual editing errors.

---

## 985. How do you rollback failed Go releases?

**Answer:**
We don't "rollback" binaries (cannot un-download).
We **Roll Forward** or Revert.
1.  `git revert <commit>`.
2.  Tag `v1.0.1` (Fix).
3.  Release.
In K8s: `kubectl rollout undo deployment/myapp`. This points the pods back to the previous Docker image instantly.

---

## 986. How do you add performance regression testing?

**Answer:**
**Benchmarks in CI**.
Store previous benchmark results (JSON).
Run current benchmarks.
`benchstat old.txt new.txt`.
If `delta` > 10% (slower), fail the build.
Tools like **Cobalt** or **PerfCheck** automate this comparison in GitHub Actions.

---

## 987. How do you build CLI-based installers in Go?

**Answer:**
1.  **Embed Assets**: `//go:embed config.yaml`.
2.  **OS detection**: `runtime.GOOS` to decide where to put files (`/etc` vs `%APPDATA%`).
3.  **Permissions**: `os.Chmod` to make scripts executable.
4.  **Self-Update**: Use `go-update` to replace the running binary with a newer downloaded version.

---

## 988. How do you generate dashboards from Go metrics?

**Answer:**
Go exports `/metrics` (Prometheus format).
Prometheus scrapes it.
**Grafana** visualizes it.
We use standard Go dashboards (Go Runtime Metrics, Garbage Collection, Goroutines) provided by the `grafana.com` marketplace (ID 6671).

---

## 989. How do you monitor file system changes in Go?

**Answer:**
`fsnotify/fsnotify`.
Wait loop:
```go
watcher, _ := fsnotify.NewWatcher()
watcher.Add("/var/log")
for {
    select {
    case event := <-watcher.Events:
        if event.Op&fsnotify.Write == fsnotify.Write {
            log.Println("modified file:", event.Name)
        }
    }
}
```
Used for hot-reloading Configs or watching for new Uploads.

---

## 990. How do you implement custom plugins in Go?

**Answer:**
1.  **Native Plugins**: `plugin` package. (Linux only, fragile).
2.  **HashiCorp Plugin**: RPC-based. The plugin is a separate process. Main app talks to it via gRPC over localhost. Secure, crash-safe, and language agnostic (plugin can be Python).

---

## 991. How do you keep Go dependencies up to date?

**Answer:**
**Renovate Bot** or **Dependabot**.
It scans `go.mod`.
Checks proxy.golang.org for new tags.
Opens PR: "Bump github.com/gin-gonic/gin from 1.7 to 1.8".
We rely on the test suite to merge it.

---

## 992. How do you audit Go packages for security issues?

**Answer:**
`govulncheck ./...`.
It queries the **Go Vulnerability Database**.
It checks if you actually *call* the vulnerable function. (Precise).
Dependabot only checks if you import the package version (Noise).
We run `govulncheck` in CI to block releases with known CVEs.

---

## 993. How do you migrate Go modules across repos?

**Answer:**
**Google/gomvpkg** (Move Package).
It moves the code AND rewrites all imports in the source tree.
If moving to a new repo:
1.  Copy code.
2.  Update `go.mod` in old repo to `replace` or alias the code (deprecated).
3.  Update all consumers to point to new module path.

---

## 994. How do you conduct performance reviews for Go codebases?

**Answer:**
Checklist:
1.  **Allocations**: Are we allocating in hot loops?
2.  **Locks**: Are we holding locks during I/O?
3.  **Database**: N+1 queries? Missing indexes?
4.  **Concurrency**: Unbounded goroutine spawning?
5.  **Benchmarks**: Do they exist for critical paths?

---

## 995. How do you implement a concurrent token bucket rate limiter?

**Answer:**
`golang.org/x/time/rate`.
`limiter := rate.NewLimiter(10, 5)` (10 req/s, burst 5).
Blocking: `limiter.Wait(ctx)`.
Non-blocking: `if !limiter.Allow() { return 429 }`.
It uses a mathematical approach (time elapsed), not a background ticker, so it's extremely CPU efficient (`mutex` lock only).

---

## 996. How do you implement the Saga Pattern for distributed transactions?

**Answer:**
(See Q 978 - Orchestrator vs Choreography).
**Choreography**:
Service A emits `EventA`.
Service B listens. Do B. Emits `EventB`.
Service C listens. Do C. Fail? Emit `EventCFail`.
Service B listens `EventCFail`. Undoes B.
Service A listens `EventCFail`. Undoes A.
Hard to debug (events flying everywhere).
**Orchestrator** (Temporal.io) is preferred in Go.

---

## 997. What is the Go Memory Model and the "Happens-Before" relationship?

**Answer:**
Official Spec defining visibility of variable writes across goroutines.
"A write to `v` in G1 is visible to G2 if..."
Rules:
1.  Channel Send happens-before Receive.
2.  Loop `go func` happens-before function body.
3.  Mutex Unlock happens-before Lock (on next turn).
If you don't use these primitives, writes are NOT guaranteed to be visible (compiler reordering / CPU cache).

---

## 998. How do you implement distributed tracing with context propagation?

**Answer:**
(See Q 752).
Headers: `traceparent` (W3C Standard).
Client: Inject `traceparent` into HTTP Header.
Server: Extract `traceparent`. Create Span.
Go Logic: `otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))`.

---

## 999. How do you use the `slices` and `maps` packages (Go 1.21+)?

**Answer:**
Generic helpers.
`slices.Sort(list)` (No need for `sort.Interface`).
`slices.Contains(list, item)`.
`maps.Clone(m)`.
`maps.DeleteFunc(m, func(k, v) bool { ... })`.
These replace common "Util" libraries with standard, optimized generic implementations.

---

## 1000. What is Profile Guided Optimization (PGO) and how do you use it?

**Answer:**
Go 1.20+.
1.  Build prod binary.
2.  Run in prod, collect CPU profile: `default.pgo`.
3.  Commit `default.pgo` to repo.
4.  Build again: `go build -pgo=auto`.
The compiler uses the profile to inline Hot Functions more aggressively and devirtualize specific interface calls.
Gain: 2-7% CPU performance reduction for free.
