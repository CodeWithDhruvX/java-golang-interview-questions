# 🟢 Go Theory Questions: 781–800 Go Tooling, CI/CD & Developer Experience

## 781. How do you build a multi-binary Go project?

**Answer:**
We organize `main` packages in the `cmd/` directory.
`cmd/server/main.go`
`cmd/worker/main.go`
`cmd/cli/main.go`

To build all:
`go build ./cmd/...`
This produces three binaries: `server`, `worker`, and `cli` in the current directory (or `$GOPATH/bin` if `go install` is used). This is the standard Go monorepo structure.

---

## 782. How do you configure GoReleaser for automated builds?

**Answer:**
**GoReleaser** automates cross-compilation, packing, and publishing (GitHub Releases, Docker Hub).
We create `.goreleaser.yaml`.
It defines:
- **Builds**: `GOOS` (linux, windows, darwin) and `GOARCH` (amd64, arm64).
- **Archives**: `.tar.gz` format.
- **Release**: GitHub token.
Running `goreleaser release` generates all binaries and uploads them with a changelog automatically.

---

## 783. How do you sign binaries in Go before release?

**Answer:**
We use **Cosign** (part of Sigstore) or GPG.
In GoReleaser:
```yaml
signs:
  - cmd: cosign
    args: ["sign-blob", "--key=cosign.key", "${artifact}"]
```
This generates a `.sig` file next to the binary. Users can verify the signature to ensure the binary hasn't been tampered with (Supply Chain Security).

---

## 784. How do you use `go vet` to detect issues?

**Answer:**
`go vet` uses heuristics to find logical bugs that compile but are likely wrong.
Common catches:
- `Printf` format mismatches (`%d` for a string).
- Unreachable code.
- Mutex locks copied by value.
We run it in CI: `go vet ./...`. It is also built into `go test` (automatically runs a subset of checks).

---

## 785. How do you manage environment-specific builds in Go?

**Answer:**
We use **Build Tags** (`//go:build ...`) or `GOOS/GOARCH`.
Example: `file_storage_aws.go` has `//go:build aws`.
`file_storage_local.go` has `//go:build !aws`.

When running `go build -tags aws`, only the AWS file is compiled. This allows us to swap entire implementations (Stub vs Real, Pro vs Free) at compile time without runtime overhead.

---

## 786. How do you use `build tags` in Go?

**Answer:**
Place a comment at the very top of the file:
`//go:build integration`
This file is ignored by default.
To run it: `go test -tags integration ./...`.
We use this for:
1.  **Integration Tests** (Slow).
2.  **OS-specific code** (Windows syscalls).
3.  **Feature Flags** (Enterprise features).

---

## 787. How do you profile CPU/memory usage in CI pipelines?

**Answer:**
We run a benchmark and compare it to the `master` branch.
Tool: **benchstat**.
CI Step:
1.  `go test -run=^$ -bench=. > new.txt`
2.  `git checkout master && go test ... > old.txt`
3.  `benchstat old.txt new.txt`
If performance drops > 10%, fail the build. This prevents regression creep.

---

## 788. How do you automate `go test` and coverage in GitHub Actions?

**Answer:**
Standard Action: `setup-go`.
```yaml
- name: Test
  run: go test -v -race -coverprofile=coverage.out ./...
- name: Upload Coverage
  uses: codecov/codecov-action@v3
```
We use `-race` to catch concurrency bugs. We use `-coverprofile` to generate the report and upload it to Codecov/SonarCloud to track % coverage trends over time.

---

## 789. How do you write a custom Go linter?

**Answer:**
We use `golang.org/x/tools/go/analysis`.
Framework: define an `Analyzer`.
`Run: func(pass *analysis.Pass) (interface{}, error)`
Inside: traverse the **AST** (Abstract Syntax Tree). Look for specific patterns (e.g., "Function name should not start with 'Test' unless in _test.go").
We integrate this into `golangci-lint` as a private plugin.

---

## 790. How do you automate versioning and changelogs in Go projects?

**Answer:**
We use **Semantic Versioning** (v1.2.3).
Tools: **Conventional Commits** + **GoReleaser**.
Commits like `feat: add login` or `fix: crash` are parsed.
GoReleaser uses this to bump the version (Minor vs Patch) and generate the changelog text in the GitHub Release payload.

---

## 791. How do you use `go:embed` for bundling files?

**Answer:**
Go 1.16+ allows embedding static assets into the binary.
```go
//go:embed static/*
var content embed.FS
```
We can treat `content` as a filesystem (`content.Open("static/index.html")`).
This is perfect for Single Binary Deployments (Web Server + React Frontend + SQL Migrations all in one `.exe`).

---

## 792. How do you validate Go module versions in a monorepo?

**Answer:**
We ensure all internal modules use the same version of third-party deps.
Tool: **Dependabot** or **Renovate**.
In a Monorepo using Go Workspaces, we run `go work sync` to ensure the `go.sum` files are consistent. We also forbid `replace` directives in `go.mod` (except for local development) to ensure reproducibility.

---

## 793. How do you containerize a Go application for fast startup?

**Answer:**
**Multistage Build**.
Stage 1 (Builder): `FROM golang:1.22`. Run `go build`.
Stage 2 (Runner): `FROM gcr.io/distroless/static`.
Copy only the binary.
Result: 5MB Docker image. No OS, no Shell, no Package Manager. Starts in milliseconds.
Secure and fast.

---

## 794. How do you enable live reloading for Go dev servers?

**Answer:**
Go is compiled, so we restart the process on file change.
Tools: **Air** or **CompileDaemon**.
Config `air.toml`: Watch `**/*.go`, Exclude `tmp/`.
cmd: `go build -o ./tmp/main . && ./tmp/main`.
This gives a "Hot Reload" experience similar to Node.js for developer velocity.

---

## 795. How do you run multiple Go services locally with Docker Compose?

**Answer:**
`docker-compose.yml`.
Services: `auth`, `payment`, `frontend`.
Volume Mount: Map local code `.` to `/app` in container.
Command: Use `Air` (see Q 794) inside the container.
This allows running the full microservices stack locally while editing code on the host, with instant live reload inside the Docker network.

---

## 796. How do you handle secrets securely in Go CI pipelines?

**Answer:**
We never commit `.env`.
In GitHub Actions: Use `${{ secrets.DB_PASS }}`.
Pass to Go test: `DB_PASS=${{ ... }} go test`.
In Go: `os.Getenv("DB_PASS")`.
For Integration Tests requiring real credentials, we use short-lived tokens (OIDC) via "Configure AWS Credentials" action, so we don't store long-lived static keys in CI secrets.

---

## 797. How do you cross-compile Go binaries for ARM and Linux?

**Answer:**
Go makes this trivial.
Command: `GOOS=linux GOARCH=arm64 go build -o app-arm64`.
We can build a Linux binary from a Mac or Windows machine without Docker. This is why Go is the language of choice for CLI tools that need to run on Raspberry Pi (ARM), Cloud (Linux), and Developer Laptops (Mac/Win) from a single codebase.

---

## 798. How do you build Go CLIs that auto-complete in Bash and Zsh?

**Answer:**
We use **Cobra**.
`rootCmd.GenBashCompletion(os.Stdout)`.
`rootCmd.GenZshCompletion(os.Stdout)`.
Users source this output in their `.bashrc`.
Cobra handles the logic: if user types `myapp deploy --reg<TAB>`, it suggests `--region`.

---

## 799. How do you keep your Go codebase idiomatic and consistent?

**Answer:**
**Linters** and **Formatters**.
1.  **gofmt** / **goimports**: Enforces style (tabs, spacing).
2.  **golangci-lint**: Runs 50+ linters (errcheck, staticcheck, revive).
We enforce this in CI/CD (`golangci-lint run`). If it fails, the PR is blocked. This assumes the role of the "Style Police" so humans focus on logic reviews.

---

## 800. How do you implement HMAC-based authentication in Go?

**Answer:**
HMAC = Hash-based Message Authentication Code.
Used for webhooks (Stripe/Slack).
Sender signs payload: `mac := hmac.New(sha256.New, secret); mac.Write(body)`.
Receiver verifies:
`expected := mac.Sum(nil)`.
`if !hmac.Equal(receivedSig, expected) { return 401 }`.
It guarantees **Integrity** (body not changed) and **Authenticity** (sender has secret).
