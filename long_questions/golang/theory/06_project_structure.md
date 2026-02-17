# 🟢 Go Theory Questions: 101–120 Project Structure and Design

## 101. What is the Standard Go Project Layout?

**Answer:**
Go doesn't officially mandate a project layout, but the community has coalesced around a standard pattern, often called the "Standard Go Project Layout."

It uses specific top-level directories: `/cmd` for your application binaries (main.go files), `/internal` for private library code that shouldn't be imported by others, and `/pkg` for library code that _should_ be importable by the wider world.

We use this structure for almost all medium-to-large projects because it’s a shared language. A new developer can open the repo and immediately know where the entry point is (`/cmd`) and where the business logic lives (`/internal`), saving hours of onboarding time.

---

## 102. Why use `/cmd` directory?

**Answer:**
The `/cmd` directory is the home for your entry points.

If your project produces multiple binaries—say, a server, a CLI admin tool, and a background worker—you don't want to clutter the root directory. Instead, you create `/cmd/server`, `/cmd/cli`, and `/cmd/worker`.

This keeps your project clean. It also separates the "wiring" logic (in `main.go`) from your actual business logic. The `cmd` folder usually just parses flags, connects to the database, and then hands off control to your library packages.

---

## 103. What is the `internal` package?

**Answer:**
`internal` is a special directory name. The Go compiler enforces a rule: code inside `internal` can only be imported by packages rooted in the same parent directory.

It’s effectively a "Private" modifier for your entire module.

We use it extensively to hide implementation details. If I write a library, I might put my ugly helper functions in `/internal`. This ensures that nobody else can depend on them. It gives me the freedom to aggressively refactor or break that code later without worrying about breaking downstream users.

---

## 104. How to structure a Hexagonal Architecture in Go?

**Answer:**
Hexagonal Architecture (or Ports and Adapters) is about decoupling your core business logic from the outside world.

In Go, we define "Ports" as **Interfaces**. Your core logic says, "I need a UserStorer interface," not "I need a Postgres database."

Then, we write "Adapters" in separate packages—like a `postgres` package that satisfies the interface. We wire them together in `main()`. This allows us to swap out the database or the HTTP framework instantly without touching a single line of business logic, making the code incredibly testable.

---

## 105. What is strict dependency injection?

**Answer:**
Strict DI means you never create your dependencies inside your logic; you ask for them.

Instead of your Service saying `db = connectDB()` inside its code, you define a constructor `NewService(db *sql.DB)`.

It sounds simple, but it changes everything. It removes hidden global state. You can look at the function signature and know *exactly* what this service needs to run. Plus, it makes testing trivial—you can pass in a mock DB instead of a real one without any monkey-patching hacks.

---

## 106. How to handle configuration management?

**Answer:**
The standard way is to read from **Environment Variables**, following the 12-Factor App methodology.

We usually define a struct that mirrors our config—`type Config struct { DB_URL string; Port int }`. Then we use a library like `kelseyhightower/envconfig` or `viper` to parse the environment variables and populate that struct at startup.

This keeps configuration seemingly centralized but allows us to deploy the same binary to Dev, Staging, and Prod just by changing the environment variables in Kubernetes.

---

## 107. Should you use dot-imports `import . "fmt"`?

**Answer:**
Generally, no. Dot-imports dump all exported names from the package directly into your current namespace.

This means `fmt.Println` becomes just `Println`. While it saves typing, it destroys readability. A reader sees `Println` and wonders, "Is that a local function? Or from a package?"

The only exception is inside **Tests**. Frameworks like Ginkgo usually dot-import their matchers so you can write natural sentences like `Expect(result).To(Equal(5))` without typing `gomega.` every time.

---

## 108. What is `go.mod` file?

**Answer:**
The `go.mod` file is the manifest for your module.

It defines two things: the generic identity of your module (its name, usually a GitHub URL) and the exact requirement versions of your dependencies.

It’s the entry point for the Go toolchain. When you run `go build`, Go looks here to know which version of libraries to fetch to ensure your build is deterministic.

---

## 109. What is `vendor` directory?

**Answer:**
The `vendor` directory is a local folder where you store a copy of all your dependencies' source code.

You create it with `go mod vendor`. If you check this into source control, you verify that you have a self-contained compilation unit.

We use it in high-security enterprise environments where build servers don't have internet access to fetch modules from GitHub. It guarantees that even if GitHub goes down or a library author deletes their repo, your build will still work.

---

## 110. How to version your Go module?

**Answer:**
Go strictly enforces **Semantic Versioning**.

If you release version `v1.0.0`, you promise stability. If you need to make a breaking change, you **must** bump the major version to `v2.0.0`.

Crucially, Go treats `v2` as a completely different library. You have to update your import paths to include `/v2` at the end. This allows old code to use `v1` and new code to use `v2` in the same binary without crashing, solving the "Diamond Dependency" problem.

---

## 111. What is the `testdata` directory?

**Answer:**
`testdata` is a directory name reserved by the Go toolchain.

The compiler explicitly ignores it. It won't try to compile any Go files inside it.

We use it to store "Golden Files"—like massive JSON samples, XML fixtures, or images that we need to load during our unit tests to verify parsing logic. It keeps your test assets separate from your source code.

---

## 112. Explain Package Oriented Design.

**Answer:**
Package Oriented Design is the philosophy that "structure follows domain."

Instead of creating generic buckets like `folder/models` and `folder/controllers`, we creates folders based on features: `folder/user`, `folder/billing`, `folder/cart`.

This groups related logic together. The `billing` package contains the billing model, the billing service, and the billing database logic. It keeps the distinct parts of your application decoupled and prevents the "Circular Dependency" hell that often plagues layered architectures.

---

## 113. How to handle cyclic dependencies?

**Answer:**
Go forbids A importing B if B imports A. It’s a hard compile error.

This feels restrictive, but it forces you to keep your architecture clean and directed (DAG).

When we hit a cycle, we usually solve it by introducing a **third package**. If A and B both need a `User` type, we move `User` into a common `types` package that both A and B can import without depending on each other.

---

## 114. What are build constraints?

**Answer:**
Build constraints (or build tags) allow you to tell the compiler which files to care about.

You can name a file `storage_postgres.go` and `storage_mysql.go`, and use build tags to decide which one gets compiled into the final binary.

This is mostly used for operating system differences, but we also use it to swap out implementations—like compiling a "stub" version of a service for local development vs the real version for production.

---

## 115. What is the `tools.go` pattern?

**Answer:**
This is a clever hack in the Go community to track developer tools.

We create a file named `tools.go` and add a build tag `//go:build tools` so it never actually compiles. Inside, we import tools we need—like `golangci-lint` or `protoc`.

This forces `go.mod` to record the exact version of the linter we are using. It ensures that every developer on the team is using the exact same version of the tools, preventing "it works on my machine" issues related to tooling.

---

## 116. How to reduce binary size?

**Answer:**
Go binaries are large because they are statically linked (including all libraries) and contain debug symbols.

To shrink them, we pass flags to the linker: `go build -ldflags="-s -w"`. This strips the symbol table and DWARF debug info, often reducing size by 30-50%.

For extreme optimization (like embedded devices), we might use `upx` to compress the binary further, though that comes with startup time costs.

---

## 117. What is graceful degradation?

**Answer:**
Graceful degradation is the ability of your system to keep working—even partially—when a dependency fails.

If your "Recommendations" service goes down, your main e-commerce page shouldn't crash. It should just hide the recommendations widget and show the rest of the page.

We implement this using **Circuit Breakers** and timeouts. If a service fails, we catch the error and fallback to a default value (or an empty list) rather than propagating the error up to the user.

---

## 118. Use of Makefile in Go?

**Answer:**
Even though Go toolchain is great, we still use Makefiles to standardize "verbs."

Commands like `go run cmd/server/main.go` or `go test -race ./...` are annoying to type. We wrap them in a Makefile so `make run` or `make test` does the right thing.

It acts as documentation. A new developer can look at the Makefile and instantly see how to build, test, and deploy the application without guessing the arguments.

---

## 119. How to organize API schema?

**Answer:**
We typically centralize API contracts—like Protocol Buffer definitions or OpenAPI specs—in a dedicated `/api` folder.

This makes it the "Single Source of Truth." We then use `go generate` to compile those specs into Go code (usually in `/pkg/pb` or similar).

Keeping them in one place allows frontend teams or other services to easily find the contract they need to strictly adhere to.

---

## 120. Designing a CLI tool in Go?

**Answer:**
Go is fantastic for CLI tools because it starts instantly and compiles to a single binary.

We usually use the **Cobra** library. It helps structure your commands like a tree (`git remote add`), handles flag parsing, and generates help text automatically.

We put the main entry point in `/cmd/mytool` and the logic in internal packages. This structure has become the gold standard for Go CLIs like `kubectl` and `docker`.
