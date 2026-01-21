## 🧪 Go Tooling, CI/CD & Developer Experience (Questions 781-800)

### Question 781: How do you create custom `go generate` commands?

**Answer:**
Write a small Go program (e.g., `gen.go`) that outputs a file.
Add `//go:generate go run gen.go` in your source.
Run `go generate ./...`.
Common for generating Enum strings, Mocks, or embedding version info.

---

### Question 782: How do you build a multi-binary Go project?

**Answer:**
Structure:
```
cmd/
  app1/main.go
  app2/main.go
internal/
pkg/
```
Build: `go build ./cmd/...` builds all binaries in `cmd`.

---

### Question 783: How do you configure GoReleaser for automated builds?

**Answer:**
Create `.goreleaser.yaml`.
Specify `builds` (os/arch targets), `archives` (zip/tar), and `release` (GitHub info).
It handles cross-compilation and packaging automatically.

---

### Question 784: How do you sign binaries in Go before release?

**Answer:**
Use `gpg` or **Cosign**.
GoReleaser supports `sign` hooks.
It runs the signing tool against the generated artifact (binary/shasum) so users can verify authenticity.

---

### Question 785: How do you use `go vet` to detect issues?

**Answer:**
Runs static analysis heuristics.
Detects: `Printf` format mismatches, Mutex copying, Unreachable code, Atomic misuse.
Run `go vet ./...` in CI.

---

### Question 786: How do you manage environment-specific builds in Go?

**Answer:**
**Build Tags.**
File `db_postgres.go` -> `//go:build !sqlite`
File `db_sqlite.go` -> `//go:build sqlite`
Build: `go build -tags sqlite`.

---

### Question 787: How do you use build tags in Go?

**Answer:**
Add comment at top of file: `//go:build linux || darwin`.
This file is ignored on Windows during compilation.

---

### Question 788: How do you profile CPU/memory usage in CI pipelines?

**Answer:**
Run benchmarks with `-cpuprofile` and store the artifact.
Use `benchstat` to compare current PR vs `main` branch performance.
Fail CI if performance drops > 10%.

---

### Question 789: How do you automate `go test` and coverage in GitHub Actions?

**Answer:**
```yaml
- run: go test -race -coverprofile=coverage.txt ./...
- uses: codecov/codecov-action@v3
  with:
    files: ./coverage.txt
```

---

### Question 790: How do you write a custom Go linter?

**Answer:**
Use `golang.org/x/tools/go/analysis`.
Define an `Analyzer`.
Walk the AST (Abstract Syntax Tree). Report nodes that violate your custom rule (e.g., "Struct field name must start with `X`").

---

### Question 791: How do you automate versioning and changelogs in Go projects?

**Answer:**
**Conventional Commits** + **GoReleaser** (or `semantic-release`).
Commits like `feat: add login` bump Minor version. `fix: bug` bump Patch.
Tools generate CHANGELOG.md based on commit history.

---

### Question 792: How do you use `go:embed` for bundling files?

**Answer:**
`//go:embed static/*`
`var staticFiles embed.FS`
Can serve a whole React frontend from a single Go binary using `http.FileServer(http.FS(staticFiles))`.

---

### Question 793: How do you validate Go module versions in a monorepo?

**Answer:**
Use `go mod tidy` in CI.
Ensure `go.work` (Workspaces) is configured if multiple modules inter-depend locally.

---

### Question 794: How do you containerize a Go application for fast startup?

**Answer:**
**Distroless** or **Scratch** image.
Build static binary (`CGO_ENABLED=0`).
Copy ONLY the binary to the image.
Size: 10-20MB. Startup: milliseconds.

---

### Question 795: How do you enable live reloading for Go dev servers?

**Answer:**
Use **Air** (`github.com/cosmtrek/air`) or **Reflex**.
These tools watch `.go` files. On Save -> Kill Process -> Recompile -> Restart.

---

### Question 796: How do you run multiple Go services locally with Docker Compose?

**Answer:**
`docker-compose.yml` defining `service-a`, `service-b`, `redis`.
Mount source code as Volume if you want live-reload inside container (slower), or run Go binary on Host and connect to Docker Redis (faster).

---

### Question 797: How do you handle secrets securely in Go CI pipelines?

**Answer:**
Inject as Environment Variables in the CI Runner settings (GitHub Secrets).
Never check `.env` into git.
Test code reads `os.Getenv()`.

---

### Question 798: How do you cross-compile Go binaries for ARM and Linux?

**Answer:**
`GOOS=linux GOARCH=arm64 go build`.
Pure Go makes this trivial (no cross-compiler toolchain needed unless CGO is involved).

---

### Question 799: How do you build Go CLIs that auto-complete in Bash and Zsh?

**Answer:**
If using **Cobra**: `rootCmd.GenBashCompletion(os.Stdout)`.
User sources this script in `.bashrc`.

---

### Question 800: How do you keep your Go codebase idiomatic and consistent?

**Answer:**
1.  **gofmt:** Enforces format.
2.  **golangci-lint:** Enforces style (errcheck, staticcheck, unused).
3.  **review:** Humans check architecture.

---
