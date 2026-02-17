# 🟢 Go Theory Questions: 381–400 Go in Real-World Projects & Architecture

## 381. How do you handle config versioning in Go projects?

**Answer:**
We decouple config Schema from Values.

The **Values** (e.g., DB host) live int Environment Variables or ConfigMaps, which are versioned by Infrastructure-as-Code (Terraform/Helm) in a separate repo.
The **Schema** (`struct Config`) lives in the Go code.
If we need to change the schema (rename `PORT` to `HTTP_PORT`), we support both for one release cycle (backward compatibility) to allow Ops to migrate the values without downtime.

---

## 382. How do you organize API versioning in Go apps?

**Answer:**
We prefer **URL Path Versioning**: `/api/v1/users`.

In code, we structure packages as:
`pkg/api/v1/handler.go`
`pkg/api/v2/handler.go`

This isolates the logic. `v2` might use a completely different struct than `v1`.
We avoid "Header Versioning" (Accept: application/vnd.company.v1+json) because it makes it harder to use standard tools like cURL or Swagger UI, and CDN caching becomes tricky.

---

## 383. How do you validate struct fields with custom rules?

**Answer:**
We use `go-playground/validator`.

We define a custom tag function.
`validate.RegisterValidation("is-isbn", func(fl FieldLevel) bool { ... })`.
Then in the struct: `ISBN string \`validate:"is-isbn"\``.

This keeps the validation logic centralized. We verify the struct *immediately* after unmarshalling JSON. If validation fails, we return a 400 Bad Request with a structured error: `{"field": "isbn", "error": "invalid format"}`.

---

## 384. How do you cache API responses in Go?

**Answer:**
We use **HTTP Middlewares**.

For modest scale, we use an in-memory LRU cache (like `hashicorp/golang-lru`) inside the middleware. Key = Request URI.
For distributed scale, we use **Redis**.
Middleware Logic:
1.  Check Redis for `GET /users/1`. If hit, write response immediately within 2ms.
2.  If miss, call `next.ServeHTTP()`. Capture the response.
3.  Write response to Redis with TTL (e.g., 60s).

---

## 385. How do you serve files over HTTP with conditional GET?

**Answer:**
We use `http.ServeContent`, not `w.Write`.

`ServeContent` handles **If-Modified-Since** headers automatically.
You must provide the file's `ModTime`.
If the browser sends `If-Modified-Since: <Time>`, and the file hasn't changed, Go automatically sends `304 Not Modified` with an empty body. This saves massive bandwidth and makes the site feel instant for returning users.

---

## 386. How do you apply SOLID principles in Go?

**Answer:**
Go's Interface system is the ultimate expression of **ISP (Interface Segregation Principle)**.

Instead of a giant `User` interface with 20 methods, we define tiny interfaces at the call site:
`type NameChanger interface { ChangeName(string) }`.
This adheres to **DIP (Dependency Inversion)**: High-level modules (Services) don't depend on low-level modules (Postgres structs); they depend on abstractions (Interfaces) that they define themselves.

---

## 387. How do you prevent breaking changes in shared Go modules?

**Answer:**
We strictly follow **Semantic Versioning** via Go Modules.

If we change a function signature, we **MUST** bump the MAJOR version (v1 -> v2).
In Go, v2 is effectively a different package: `github.com/my/lib/v2`.
This allows consumers to import *both* v1 and v2 simultaneously if needed during a migration, preventing the "Diamond Dependency" hell common in other languages.

---

## 388. What is the difference between horizontal and vertical scaling in Go services?

**Answer:**
**Vertical**: Buying a bigger server (more RAM/CPU).
Go scales vertically exceptionally well because of the efficient scheduler; one process can saturate 128 cores easily.

**Horizontal**: Adding more machines (Pods).
This is preferred for reliability. Go microservices are designed to be stateless. We use a Load Balancer to distribute traffic. Horizontal scaling allows Zero Downtime Deployments (Rolling Update), which Vertical scaling usually cannot (requires restart).

---

## 389. How do you support internationalization in Go?

**Answer:**
We use `golang.org/x/text`.

We don't hardcode strings: `fmt.Println("Hello")`.
We use a printer: `p.Printf("Hello")`.
The `p` is created from the user's `Accept-Language` header. We store translations in `catalog.go` files (often generated from JSON/PO files). The library handles the complex logic of Pluralization ("1 item" vs "2 items") which varies wildly across languages.

---

## 390. How do you write a Go SDK for third-party APIs?

**Answer:**
We design it to be **Testable** and **Configurable**.

1.  **Interfaces**: Don't return concrete structs if possible, or at least allow injecting a custom `HTTPClient`.
2.  **Options Pattern**: `NewClient(apiKey, WithTimeout(5s), WithBaseURL("..."))`.
3.  **Context**: Every method *must* accept `context.Context` to allow cancellation. `GetUsers(ctx, filters)`.
4.  **Typed Errors**: Define `ErrNotFound`, `ErrRateLimited` so users can check `errors.Is(err, sdk.ErrRateLimited)`.

---

## 391. How do you manage request IDs and trace IDs?

**Answer:**
We use a **Middleware** at the edge (Ingress/Gateway).

It looks for `X-Request-ID`. If missing, it generates a UUID.
It puts this ID into the `context`.
We wrap the logger to automatically extract this ID from the context and prepend it to every log line. This ensures that even if 1000 requests happen in parallel, we can filter logs for just *one* specific user journey.

---

## 392. How do you implement audit logging in Go?

**Answer:**
Audit logs are different from application logs; they are a **Legal Record**.

We use a distinct pipeline. We might write to a separate standard `slog` logger that targets a secure, append-only storage (like S3 Object Lock or a dedicated SQL table).
We log **Intent** ("User attempted to delete X") and **Outcome** ("Success/Forbidden"). We often implement this via a Decorator pattern around the Service layer methods to ensure no business action is missed.

---

## 393. How would you version a binary CLI in Go?

**Answer:**
We embed version info at **Link Time**.

In code: `var Version = "dev"`.
In Makefile: `go build -ldflags="-X main.Version=v1.0.2 -X main.CommitHash=abc1234"`.

The CLI command `myapp version` prints these variables.
This ensures the binary carries its own provenance info, which is critical when debugging a customer issue ("Which version are you running?").

---

## 394. How do you ensure backward compatibility in Go libraries?

**Answer:**
We use the **"Accept Interfaces, Return Structs"** rule.

If a function accepts an interface, the user can pass anything. If we add a method to that interface, we break the user's implementation. So we keep interfaces tiny.
If we return a struct, we can add new fields to it later safely (users just ignore new fields).
We also run `apidiff` tools in CI to detect accidental breaking changes to exported symbols before merging.

---

## 395. How do you handle soft deletes in Go models?

**Answer:**
In the database, we iterate a `deleted_at` timestamp column (nullable).

In Go structs: `DeletedAt *time.Time`.
In Repositories, we must remember to always add `WHERE deleted_at IS NULL`.
If using GORM, this is automatic (`gorm.Model`). However, soft deletes add complexity (unique indexes need to include `deleted_at`). For high-volume data, we prefer "Hard Deletes" coupled with an "Archive Table" to keep the main index small.

---

## 396. How do you refactor a large legacy Go codebase?

**Answer:**
We use the **Strangler Fig Pattern** inside the code.

We define a new Interface for the component we want to rewrite.
We write the new implementation in a new package.
We switch the `main.go` wiring to use the new implementation.
We use **Feature Toggles** to route 1% of traffic to the new code to verify correctness before doing a 100% rollout. We rely heavily on the compiler; renaming a field triggers compile errors everywhere, guiding the refactor.

---

## 397. How do you maintain a mono-repo with multiple Go modules?

**Answer:**
We look at `go.work` (Go Workspaces, introduced in Go 1.18).

It allows us to work on multiple modules simultaneously (`use ./mod-a`, `use ./mod-b`) and having the IDE (gopls) understand the cross-module references without publishing to a remote.
We use **Bazel** or **Turborepo** (rare for Go, but growing) to run tests only for modules that changed, speeding up CI.

---

## 398. How would you go about building a plugin system in Go?

**Answer:**
Native `plugin` package is notoriously difficult (Linux only, requires exact build environment match).

Instead, we use **HashiCorp/go-plugin**.
It runs plugins as *separate processes* (subprocesses) and communicates via **gRPC** over a local socket.
This isolates the crash. If a plugin panics, the main app survives. It also allows plugins to be written in other languages (Python/Java), not just Go.

---

## 399. How do you document Go APIs automatically?

**Answer:**
We use **Swagger/OpenAPI** via comments using `swaggo/swag`.

You write magic comments above your handler:
`// @Success 200 {object} UserResponse`
Run `swag init`. It parses the AST and generates the `swagger.json`.
We serve this JSON using Swagger UI middleware. This keeps the docs valid because they live right next to the code.

---

## 400. How do you track tech debt and enforce code quality in large Go teams?

**Answer:**
We use strict linter rules (`golangci-lint`) with a config committed to the repo.

Rule: **Zero Warnings**. If the linter complains, the build fails.
We also track **Cyclomatic Complexity** (`gocyclo`). If a function score > 15, you must refactor.
For tech debt, we use `// TODO` comments. We have scripts that grep these or import them into Jira. If a TODO stays for 6 months, we declare "Task Bankruptcy" and admit we aren't going to do it, or schedule a dedicated sprint.
