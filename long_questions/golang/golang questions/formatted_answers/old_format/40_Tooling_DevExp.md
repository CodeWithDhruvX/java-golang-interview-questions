# 🧪 **781–800: Go Tooling, CI/CD & Developer Experience**

### 781. How do you create custom `go generate` commands?
"I create a tool (e.g., `gen-assets.go`).
In my source: `//go:generate go run gen-assets.go`.
Running `go generate ./...` executes it.
I use it for generating Mocks, Protobufs, or Embedding static assets. It keeps the build reproducible without external Makefiles."

#### Indepth
**Versioning Tools**. `go:generate` depends on tools installed in your `$GOPATH/bin`. If Developer A has `mockgen v1.6` and Developer B has `v1.7`, `go generate` produces different files. Use a `tools.go` file with `_ "github.com/golang/mock/mockgen"` imports to lock tool versions in `go.mod`, and run `go run github.com/golang/mock/mockgen ...` instead of relying on global binaries.

---

### 782. How do you build a multi-binary Go project?
"I structure `cmd/` folders.
`cmd/server/main.go`.
`cmd/worker/main.go`.
Build:
`go build -o bin/server ./cmd/server`.
`go build -o bin/worker ./cmd/worker`.
They share code from `internal/` and `pkg/` but produce distinct binaries."

#### Indepth
**Magefile**. Using `go build` commands manually is error-prone. Instead of Makefiles (which require `make` installed), use **Mage** (Make-in-Go). You write build scripts in Go (`mage.go`). This is cross-platform (works on Windows without WSL) and allows you to use the full power of Go for build logic.

---

### 783. How do you configure GoReleaser for automated builds?
"`.goreleaser.yaml`.
I define `builds` (linux/amd64, darwin/arm64) and `archives`.
On `git tag v1.0.0`, GitHub Action runs `goreleaser release`.
It builds, signs, and uploads everything to GitHub Releases automatically."

#### Indepth
**Snapshots**. For local testing, running a full release is slow. Use `goreleaser release --snapshot --rm-dist`. This builds everything locally without uploading to GitHub. It allows you to verify that the Docker images and binaries are generated correctly before tagging a real release.

---

### 784. How do you sign binaries in Go before release?
"I use **cosign** or standard GPG.
GoReleaser supports `sign` hooks.
It runs `cosign sign-blob` on the generated binary.
Users can verify the signature to ensure the binary wasn't tampered with. This is crucial for supply chain security."

#### Indepth
**SBOM**. Supply Chain attacks are real. GoReleaser helps generate an **SBOM** (Software Bill of Materials) in CycloneDX or SPDX format. This lists every dependency and version in your binary. Security scanners use this to quickly find if you are affected by a new vulnerability (like Log4j) without reverse engineering your binary.

---

### 785. How do you use `go vet` to detect issues?
"It runs automatically with `go test`.
It catches:
*   `Printf` format mismatches.
*   Unreachable code.
*   Struct tag typos.
I run it explicitly in CI: `go vet ./...` to block bad merges."

#### Indepth
**Shadowing**. One common bug `go vet` catches is **Variable Shadowing**, but it's not enabled by default. `err := foo(); if err != nil { err := bar() }`. The inner `err` shadows the outer one. Use `go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest` and run `go vet -vettool=$(which shadow) ./...` to catch this.

---

### 786. How do you manage environment-specific builds in Go?
"**Build Tags**.
`//go:build pro` vs `//go:build dev`.
`func config() string { return "production" }` (in `config_prod.go`).
`func config() string { return "localhost" }` (in `config_dev.go`).
`go build -tags prod`.
This excludes debug code from the production binary completely."

#### Indepth
**Integration Tests**. You don't want to run slow integration tests on every "Save". Add `//go:build integration` to `main_test.go`. Now `go test ./...` skips them (fast). Run `go test -tags=integration ./...` in CI to run them. This keeps the inner dev loop tight.

---

### 787. How do you use `build tags` in Go?
"To support multiple OS features.
`file_windows.go`: `//go:build windows`.
`file_unix.go`: `//go:build !windows`.
Go automatically picks the right file.
I also use it for integration tests: `//go:build integration` so they don't run during normal unit tests."

#### Indepth
**Boolean Logic**. You can do complex logic: `//go:build (linux || darwin) && !cgo`. This replaced the old comment syntax `// +build linux,darwin` which was confusing (comma = OR, newline = AND). Always use the new syntax (Go 1.17+).

---

### 788. How do you profile CPU/memory usage in CI pipelines?
"I write a Benchmark.
`go test -bench=. -benchmem -cpuprofile=cpu.out`.
I can compare the result against the `main` branch (**Benchstat**).
If performance degrades > 10%, I fail the build. It prevents regressions."

#### Indepth
**Noise**. CI environments are noisy (shared CPU). A benchmark might fail just because the runner was busy. Use `benchstat` to compare *multiple* runs (count=10). It uses statistical tests (Mann-Whitney U-test) to tell you if the difference is "Statistical Noise" or "Real Regression". Only fail on "Real".

---

### 789. How do you automate `go test` and coverage in GitHub Actions?
"Step 1: `actions/setup-go`.
Step 2: `go test -race -coverprofile=coverage.txt ./...`.
Step 3: Upload `coverage.txt` to Codecov.
It tracks trend lines: 'Coverage dropped 5%'. I enforce a hard floor (e.g., 80%) to pass PRs."

#### Indepth
**Goveralls**. If you don't use Codecov, use `mattn/goveralls` to send coverage to Coveralls.io. Note that "Line Coverage" is a weak metric. You can have 100% coverage and still have bugs. Use "Mutational Testing" (`test-mutation`) to see if your tests *actually* fail when the code is broken.

---

### 790. How do you write a custom Go linter?
"I use `golang.org/x/tools/go/analysis`.
I define an `Analyzer`.
`func run(pass *analysis.Pass)`.
I traverse the AST.
If I find a pattern (e.g., 'calling Log without Context'), I report `pass.Reportf`.
I plug this into `golangci-lint` as a custom plugin."

#### Indepth
**Ruleguard**. Writing AST analyzers is hard. Use `quasilyte/go-ruleguard`. It allows you to write linter rules in a Go-like DSL. `Match("fmt.Sprintf(\"%d\", $x)").Where(m["x"].Type.Is("int")).Report("Use strconv.Itoa($x) instead")`. This makes adding custom team conventions ("Don't use `log.Print`, use `zap`") trivial.

---

### 791. How do you automate versioning and changelogs in Go projects?
"I use **Conventional Commits** (`feat:`, `fix:`).
I use a tool like `release-please`.
It analyzes commit messages since the last tag.
It bumps the SemVer (Patch for fix, Minor for feat).
It generates `CHANGELOG.md` and creates the Release Tag."

#### Indepth
**SemVer Tricks**. If you are pre-v1 (`v0.2.3`), "Breaking Changes" don't require a Major bump (v1.0.0), just a Minor bump (`v0.3.0`). Once you hit `v1.0.0`, strict Semantic Versioning applies. `go mod` treats `v1` and `v2` as completely different packages (`github.com/foo/bar/v2`). Keeping `v0` for a long time gives you flexibility.

---

### 792. How do you use `go:embed` for bundling files?
"It embeds static assets into the Go binary.
`//go:embed static/*`
`var assets embed.FS`.
This turns my simple web server into a **single binary** deployment—no need to copy HTML/CSS files alongside the executable."

#### Indepth
**HTTP FileSystem**. `embed.FS` implements `fs.FS`. To serve it over HTTP: `http.FileServer(http.FS(assets))`. Warning: `embed` preserves the directory structure. If your file is in `static/index.html`, the user must visit `/static/index.html`. You typically need to `fs.Sub(assets, "static")` to "root" the server inside the folder.

---

### 793. How do you validate Go module versions in a monorepo?
"I use a workspace or a script.
I check that all modules use the same version of shared dependencies (e.g., `grpc v1.50`).
If `module-a` uses `v1.50` and `module-b` uses `v1.40`, `go build` might panic at runtime due to symbol mismatch. I enforce consistency in CI."

#### Indepth
**Go Workspaces**. Go 1.18 introduced `go.work`. It allows you to work on multiple modules locally without messy `replace` directives in `go.mod`. `go work use ./mod-a ./mod-b`. The editor (VSCode) sees them as one big project, allowing "Go to Definition" to jump across module boundaries.

---

### 794. How do you containerize a Go application for fast startup?
"**Multi-stage build**.
Build stage: `golang:1.24`. Compile static binary (`CGO_ENABLED=0`).
Final stage: `scratch` or `distroless/static`.
Copy binary.
Result: 10MB image. Starts instantly. No OS overhead."

#### Indepth
**Ko**. `ko` is a tool (by Google) that builds Go container images *without Docker*. It compiles the binary locally and wraps it in a tarball layer directly. It's faster than `docker build` and doesn't require a Docker daemon. Great for CI/CD pipelines (Kaniko alternative).

---

### 795. How do you enable live reloading for Go dev servers?
"I use **Air**.
Config: `air.toml`.
It watches `.go` files.
On change: Kills old process, rebuilds, restarts.
It feels like Node.js development."

#### Indepth
**Proxying**. While `Air` is running, if you hit a syntax error, the server crashes. The browser sees "Connection Refused". Better setup: `Air` runs a temporary *Proxy* on port 8080. It forwards traffic to your app on 8081. If your app crashes, the Proxy holds the connection and waits for the rebuild, preventing the "Site can't be reached" error.

---

### 796. How do you run multiple Go services locally with Docker Compose?
"`docker-compose.yml`.
Service A depends on DB. Service B depends on A.
I mount the source code and use `Air` inside the container for hot-reloading *inside* Docker networking context."

#### Indepth
**Host Networking**. On Linux, you can use `network_mode: host` to let the container share the host's IP/Ports. This removes the need for port mapping. On Mac/Windows, this doesn't work (VM isolation). Use `host.docker.internal` DNS name to access the "Host localhost" from inside the container.

---

### 797. How do you handle secrets securely in Go CI pipelines?
"GitHub Secrets injected as Env Vars.
`go test` reads `os.Getenv("API_KEY")`.
I verify *never* to print these secrets to the console log."

#### Indepth
**OIDC**. Long-lived API keys (AWS_ACCESS_KEY) in GitHub Secrets are a security risk. Use **OpenID Connect (OIDC)**. GitHub Actions exchanges a temporary JWT token with AWS/GCP to get short-lived credentials for that specific job. No static keys to rotate or leak.

---

### 798. How do you cross-compile Go binaries for ARM and Linux?
"`GOOS=linux GOARCH=arm64 go build`.
That's it.
No cross-compiler toolchain needed (unlike C++).
This makes it trivial to build for Raspberry Pi or AWS Graviton instances from my MacBook."

#### Indepth
**CGO Cross-Compile**. If `CGO_ENABLED=1` (e.g., using SQLite), cross-compilation is hell. You need a C cross-compiler (`aarch64-linux-gnu-gcc`). Use **Zig**. `CC="zig cc -target aarch64-linux" CGO_ENABLED=1 go build`. Zig acts as a drop-in C compiler that supports every target out of the box.

---

### 799. How do you build Go CLIs that auto-complete in Bash and Zsh?
"Using **Cobra**.
`cmd.Root().GenBashCompletion(os.Stdout)`.
I instruct the user to `source <(my-cli completion bash)`.
Cobra handles the magic of proposing flags and subcommands."

#### Indepth
**Fig/Carapace**. Cobra's bash completion is okay. For modern, rich auto-completion (with icons and detailed descriptions), tools like **Carapace** or **Fig** integrate with Cobra. They introspect your Go binary and generate specs for Zsh/Fish/PowerShell that feel like a native GUI menu.

---

### 800. How do you keep your Go codebase idiomatic and consistent?
"**Machine Enforcement**.
1.  `gofmt` (Formatting).
2.  `golangci-lint` (Linting).
3.  `goimports` (Import sorting).
I run these on `pre-commit` hook and CI.
Code Review focuses on logic/design, not style."

#### Indepth
**Revive**. `golangci-lint` includes `revive`. Revive is a faster, configurable drop-in replacement for `golint`. It allows you to disable specific annoying rules (like "Comment the exported function") or enforce strict new ones. It gives you control over the "Idiomatic" definition.
