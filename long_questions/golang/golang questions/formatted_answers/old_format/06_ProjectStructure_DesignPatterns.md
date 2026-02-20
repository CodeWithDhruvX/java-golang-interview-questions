# 🟢 **101–120: Project Structure & Design Patterns**

### 102. How do you organize a large-scale Go project?
"I almost always adhere to the **Standard Go Project Layout**. It’s not official, but it’s the community standard.

**`cmd/`**: Contains the main applications. Each folder here is a binary (e.g., `cmd/server`, `cmd/worker`).
**`internal/`**: Contains the private application code. Go’s compiler actively prevents other projects from importing packages inside `internal`. This is where 90% of my codebase lives.
**`pkg/`**: Contains library code that is ok for external projects to use (like a generic validation library).

I also use `api/` for OpenAPI/Protobuf definitions and `build/` for Dockerfiles. This structure keeps the root directory clean and makes it obvious where the entry points are."

#### Indepth
This layout is often called "Standard Go Project Layout" (from the `golang-standards` repo), but it's not official. The official Go team suggestion is often flatter. However, for 10k+ LOC projects, `internal/` is non-negotiable to prevent tight coupling between microservices sharing a monorepo.

---

### 103. What is the standard Go project layout?
"It’s a convention popularized by the community (specifically the `golang-standards` repo) to organize code in a way that scales.

It separates 'executable' code (`cmd`) from 'library' code (`pkg`) and 'private business logic' (`internal`).
Before I adopted this, my projects were flat folders that became unmanageable. Now, even if I jump into a new codebase, if I see `internal/auth`, I know exactly what it does and that it's private to this project."

#### Indepth
The Go compiler enforces `internal` privacy. If `project-b` tries to import `github.com/project-a/internal/foo`, the build fails. This is the *only* way to enforce module boundaries in Go. It allows you to break APIs in `internal` without worrying about breaking downstream users.

---

### 104. What is the `cmd` directory used for in Go?
"The `cmd` directory holds the entry points (the `main` packages) for the project.

I usually have multiple subdirectories: `cmd/api`, `cmd/cron`, `cmd/admin-cli`.
The `main.go` files inside these folders should be **small**. They should effectively just parse flags, initialize configuration, set up the dependency injection, and call a `Run()` function located in `internal`. If my `main` function is 500 lines long, I know I've done something wrong."

#### Indepth
Keep `main` boring. It should only be "configuration wiring". It reads env vars, instantiates struct dependencies (SQL DB, Redis), and injects them into the `server` struct. This makes your application logic testable because the `server` doesn't know *how* `db` was created, only that it satisfies the interface.

---

### 105. How do you structure code for reusable packages?
"If I want a package to be reusable by *other* projects, I put it in `pkg/`.

For example, if I write a structured logger wrapper that wraps Zap, I put it in `pkg/logger`.
However, I’m careful not to create a 'utils' dumping ground. Go packages should be domain-specific. Instead of `pkg/utils`, I create `pkg/stringutil` or `pkg/sliceutil`. Small, focused packages are easier to test and reuse."

#### Indepth
Avoid the `common` or `util` package anti-pattern. These become garbage dumps where unrelated functions (string checks, basic math, HTTP helpers) live together. This causes dependency bloat: importing `common` for a specific string function might pull in `net/http` and `sql` dependencies unnecessarily.

---

### 106. What are Go's most used design patterns?
"Go favors composition over inheritance, so the patterns look different than in Java.

**Functional Options**: For configuring complex objects (`NewServer(WithPort(8080))`).
**Middleware**: For wrapping HTTP handlers (Decorator pattern).
**Adapter**: Using interfaces to make incompatible types work together.
**Worker Pool**: For concurrency.
**Singleton**: Using `sync.Once` for lazy initialization.

I rarely use complex patterns like Abstract Factory or Flyweight because Go’s simplicity often makes them unnecessary."

#### Indepth
Many GoF patterns (like Decorator, Observer, Strategy) are built into the language via **Interfaces** and **Higher-Order Functions**. You don't need a `Strategy` class; you just pass a `func(Context) error`. You don't need a `Singleton` class; you check `sync.Once`. Use the language features instead of implementing Java patterns.

---

### 107. Explain the Factory Pattern in Go.
"In Go, we don't have classes, so a 'Factory' is just a simple function that returns an initialized struct.

Conventionally, we name them `New` or `New[Type]`.
`func NewService(db *sql.DB, logger *Logger) *Service`.
This function handles the setup logic: validating config, initializing maps (to avoid nil pointer panics), and dependency wiring. I always return a pointer from a factory to avoid copying large structs."

#### Indepth
Factories in Go often return `(Struct, error)` to handle initialization failures (like invalid config). A common pattern is **Must**: `func MustNewService(...) *Service` which panics on error. This is useful for top-level global variables or `init()` blocks where errors are unrecoverable anyway.

---

### 108. How to implement Singleton Pattern in Go?
"The idiomatic way is to use the `sync.Once` primitive.

`var once sync.Once`
`var instance *Singleton`

`func GetInstance() *Singleton { once.Do(func() { instance = &Singleton{} }); return instance }`

`sync.Once` guarantees that the initialization function runs **exactly once**, even if 100 goroutines call `GetInstance` at the same nanosecond. It’s safer and faster than manually managing mutexes."

#### Indepth
Under the hood, `sync.Once` uses an atomic variable `done`. If `done == 0`, it locks a mutex and runs the function. If `done == 1`, it returns immediately. This "fast path" means calling `GetInstance()` in a tight loop is extremely cheap (atomic load overhead, approx 0.5ns).

---

### 109. What is Dependency Injection in Go?
"I rarely use DI frameworks (like generic containers). I prefer **Constructor Injection**.

My structs define their dependencies as interfaces in their fields.
`type Server struct { db Database; mail Mailer }`
Then I pass them in the factory: `func NewServer(db Database, mail Mailer) *Server`.

This makes the dependency graph explicit. If I add a dependency, the compiler forces me to update the `main.go` wiring, so I can't forget it."

#### Indepth
If the dependency graph gets huge (nested dependencies), consider using **Wire** (from Google). It's a compile-time code generation tool for DI. It analyzes your providers and generates the initialization code. Avoid run-time reflection-based DI frameworks (like Uber's Dig or Facebook's Inject) unless necessary, as they hide errors until runtime.

---

### 110. What is the difference between composition and inheritance in Go?
"Go does not have class-based inheritance. We cannot say `class Dog extends Animal`.

Instead, we use **Composition** via struct embedding.
`type Dog struct { Animal }`.
`Dog` now has all the methods of `Animal`, but it **is not** an `Animal` in the type system (unless they share an interface). I prefer composition because it prevents the 'fragile base class' problem where changing a parent breaks all children. It keeps types loosely coupled."

#### Indepth
Embedded fields can be accessed directly (`dog.Eat()`) but are also accessible by their type name (`dog.Animal.Eat()`). Method promotion only works if the outer struct doesn't override the method. This allows "default implementations" via embedding, which can be overridden by the child struct.

---

### 111. What are Go generics and how do you use them?
"Generics (introduced in Go 1.18) allow me to write functions and data structures that work with **any** type.

For example, `func Reverse[T any](s []T) []T`.
Before generics, I had to either write `ReverseInt`, `ReverseString`, etc., or use `interface{}` and lose type safety.
I use them mostly for utility libraries (sets, queues, linked lists, map transformations). I avoid them in business logic where interfaces usually suffice."

#### Indepth
Go's generics implementation uses **GC Shape Stenciling**. It generates different machine code for different underlying types (int vs string), but shares code for types with the same memory layout (pointers). This balances binary size code bloat with execution performance.

---

### 112. How to implement a generic function with constraints?
"I use an interface to constrain what types are allowed.

`func Add[T Number](a, b T) T { return a + b }`
Here, `Number` is an interface that includes `int | float64`.
Without this constraint, the compiler would complain because it doesn't know if `T` supports the `+` operator. Constraints tell the compiler exactly what capabilities the type argument must have."

#### Indepth
The `constraints` package (currently in `golang.org/x/exp/constraints`) provides standard constraints like `Ordered` (supports <, >) and `Signed/Unsigned`. You can create custom constraints using union elements `type MyConstraint interface { int | string }`.

---

### 113. What are type parameters?
"Type parameters are the variables inside the square brackets `[...]`.

In `func Map[K comparable, V any](m map[K]V)`, `K` and `V` are the type parameters.
They are placeholders that get replaced by concrete types (like `string` and `int`) at compile time. This process is called **monomorphization**—the compiler actually generates a version of the function for `string` and another for `int`, so there is zero runtime performance cost."

#### Indepth
Because of monomorphization, generic code is usually as fast as hand-written code for each type. However, it increases compile time and binary size slightly. It is distinct from Java Generics (Type Erasure) where `List<Integer>` is just `List<Object>` at runtime involving boxing/unboxing overhead.

---

### 114. Can you implement the Strategy pattern using interfaces?
"Yes, interfaces are the perfect tool for the Strategy pattern.

I define a behavior, like `PaymentStrategy` with a method `Pay(amount)`.
Then I implement it with `CreditCard` and `PayPal` structs.
My context (e.g., `Checkout`) accepts a `PaymentStrategy`. This allows me to swap the payment method at runtime without changing the Checkout logic. It adheres perfectly to the defaults."

#### Indepth
This pattern is so common in Go that it's often just a function field. `type Server struct { notify func(msg string) }`. The strategy is just a function. You can plug in `fmt.Println` or `smtp.SendMail` directly without defining a formal `Notifier` interface.

---

### 115. What is middleware in Go web apps?
"Middleware is code that wraps an `http.Handler` to perform logic **before** or **after** the request is processed.

I implemented it as a function chain:
`func Logger(next http.Handler) http.Handler`.
It returns a handler that logs the request, calls `next.ServeHTTP`, and then logs the response time.
I use middleware for cross-cutting concerns: Authentication, Logging, CORS, Gzip compression, and Panic Recovery."

#### Indepth
The `http.Handler` interface is composable. Libraries like `alice` or `negroni` help chain middleware: `chain := alice.New(Auth, Limit, Log).Then(FinalHandler)`. But typically, a simple loop `for i := len(m)-1; i >= 0; i-- { h = m[i](h) }` is enough to build the chain.

---

### 116. How do you structure code using the Clean Architecture?
"I separate code into concentric layers:
1.  **Domain (Core)**: Pure business logic and entities. No external dependencies.
2.  **Usecase (Service)**: Application logic that orchestrates the domain.
3.  **Interface Adapters (Delivery)**: HTTP Handlers, CLI commands.
4.  **Infrastructure**: Database implementations, External APIs.

The Golden Rule is: **Dependencies only point inward.** The Service knows about the Domain, but the Domain knows nothing about the Service (or the DB)."

#### Indepth
In Clean Architecture (or Hexagonal/Onion), the core logic has no imports from the outside world. It defines interfaces (`Port`) that external layers must implement (`Adapter`). This makes the core pure and 100% unit testable without mocks of HTTP or SQL drivers.

---

### 117. What are service and repository layers?
" The **Repository Layer** is responsible for 'Data Access'. It speaks SQL, Redis, or File I/O. Its job is to CRUD entities.
The **Service Layer** is responsible for 'Business Logic'. It calls the repository to get data, applies rules (e.g., 'User must be over 18'), and returns results.

Separating them is critical. It allows me to swap the database (e.g., Postgres to Mongo) just by rewriting the Repository, without touching a single line of the Service logic."

#### Indepth
Typically, the Service receives a Domain Entity from the Repository, but returns a DTO (Data Transfer Object) or the Entity itself to the Controller. Avoid leaking DB tags `sql:"..."` into the Service layer. Convert DB models to pure Domain models in the Repository.

---

### 118. How would you separate concerns in a RESTful Go app?
"I avoid putting logic in HTTP handlers.

**Handler**: Decodes the JSON request, validates required fields, and calls the Service.
**Service**: Performs the actual work (e.g., 'Register User').
**Repository**: Saves the user to the DB.

If I put SQL queries inside the HTTP handler, I can't test the logic without spinning up a server, and I can't reuse that logic for a CLI command or a background worker."

#### Indepth
This separation also helps with validaton. The Handler validates **Struct/Format** (is "email" a valid email string?). The Service validates **Business Rules** (is this email banned?). The Repository validates **Data Constraints** (is this email unique?).

---

### 119. What is the importance of interfaces in layered design?
"Interfaces are the **boundary** between layers.

My Service layer defines a `UserRepository` interface. The Infrastructure layer implements it.
This inversion of control is what makes the architecture 'clean'. The high-level policy (Service) doesn't depend on the low-level detail (SQL). It depends on an abstraction. This makes unit testing trivial because I can generate a mock for the interface."

#### Indepth
Define interfaces where you **use** them, not where you implement them. `Service` should define `UserRepository`, not the `db` package. This adheres to the Interface Segregation Principle—Service only defines the methods it actually needs (`Get`), ignoring methods it doesn't (`Delete`).

---

### 120. How would you implement a plugin system in Go?
"For a native approach, I can use the `plugin` package to load shared object files (`.so`) at runtime. But this is Linux-only and has many caveats (shared dependencies must match exactly).

For a robust system, I prefer **RPC-based plugins** (like HashiCorp's `go-plugin`). The plugin runs as a separate process and talks to the main app via gRPC. It’s safer—if the plugin crashes, my app stays up."

#### Indepth
Go's `plugin` package has strict requirements: the plugin and the main app must be built with the *exact* same Go version, build tags, and dependency versions. This fragility makes it unsuitable for most projects. RPC plugins (gRPC over Unix Socket) are the industry standard (used by Terraform, Vault, Caddy).

---

### 121. How do you avoid circular dependencies in Go packages?
"Go strictly forbids import cycles (`A -> B -> A`).

I avoid them by designing a strict hierarchy. `cmd` imports `internal`, `internal` imports `pkg`.
If I hit a cycle, it usually means two packages are too tightly coupled. I fix it by:
1.  **Merging** them into one package.
2.  **Extracting** the shared types (interfaces/structs) into a third, lower-level package (e.g., `types` or `common`)."

#### Indepth
Tools like `go-cyclic` or `madge` can visualize circular dependencies. Cycles often indicate a design flaw where two components have mixed responsibilities. Breaking the cycle usually improves the design by forcing you to define a clear contract (interface) in a separate package.
