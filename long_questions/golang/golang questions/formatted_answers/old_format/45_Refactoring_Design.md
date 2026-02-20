# 🧰 **881–900: Refactoring, CLI, WebAssembly & Design**

### 881. How do you refactor large Go codebases safely?
"Incrementally.
1.  **Tests First**: Ensure high coverage.
2.  **Type Aliases**: Move types to new packages while keeping aliases in old ones to avoid breaking importers.
3.  **Interface Abstraction**: Replace concrete structs with interfaces to decouple components.
I use `gopls` rename capabilities extensively."

#### Indepth
**Atomic Commits**. Refactoring breaks things. Don't mix "Refactor" and "Feature" in one PR. Use **Atomic Refactoring**: 1. Introduce new Interface (Commit 1). 2. Make Old Code use Interface (Commit 2). 3. Swap Implementation (Commit 3). If Commit 3 breaks prod, you revert only Commit 3, leaving the clean interface structure in place.

---

### 882. How do you break a monolith Go app into microservices?
"Identify **Bounded Contexts**.
Isolate a module (e.g., 'Billing').
Define its Interface.
Move it to a separate folder.
Replace the direct function call with a Network Client (gRPC/HTTP).
Deploy it as a separate binary.
This creates a 'distributable monolith' initially, then evolves."

#### Indepth
**Strangler Fig Pattern**. Don't rewrite the whole monolith. Put a Proxy (Nginx/Envoy) in front. Route `/api/v1/billing` to the New Go Microservice. Route everything else to the Legacy Monolith. Gradually "strangle" the monolith by moving routes one by one until nothing is left. This manages risk.

---

### 883. How do you improve code readability in Go?
"**Flat is better than nested**.
I return early (`if err != nil { return }`).
I avoid `else`.
I name variables based on scope: short name `i` for small loop, descriptive `customerBalance` for package global.
I verify to document *Why*, not *What*."

#### Indepth
**Cyclomatic Complexity**. Tools like `gocyclo` measure how many `if/for/switch` paths a function has. If specific metric > 15, the function is too hard to test. Refactor by extracting the body of the `if` block into a named function. `if checksPassed() { process() }` is better than 50 lines of nested logic.

---

### 884. How do you organize domain-driven projects in Go?
"Root: `domain/` (Entities, Interfaces). Pure Go.
`app/` (Use Cases). Depends on `domain`.
`infra/` (Postgres, HTTP). Depends on `domain`.
`cmd/` (Main). Wires them up.
This is the **Hexagonal Architecture**. It keeps the core logic pristine."

#### Indepth
**Ports and Adapters**. `domain` defines "Ports" (Interfaces: `UserRepository`). `infra` defines "Adapters" (`PostgresRepository`). The `app` layer injects the Adapter into the Port. This means `domain` depends on *nothing*. You can swap Postgres for a MockInMemoryRepo without touching one line of business logic.

---

### 885. How do you handle circular dependencies?
"Go forbids them.
It usually means my design is coupled.
Fix 1: **Interface**. Define the interface in Package A, implement in Package B.
Fix 2: **Third Package**. Move the shared type to Package C.
Fix 3: **Merge**. Packages A and B are actually one logical component."

#### Indepth
**DIP**. Dependency Inversion Principle. High-level modules shouldn't depend on low-level modules. Both should depend on Abstractions. If `Auth` (High) needs `DB` (Low), `Auth` should define `type UserStore interface`. `DB` implements it. Now `Auth` doesn't import `DB`. `DB` doesn't import `Auth`. `Main` wires them. Loop broken.

---

### 886. How do you structure reusable Go modules?
"I put code in the root if small.
`github.com/me/mylib`.
I avoid `src/` or `pkg/` folders for libraries.
I use `internal/` for code I don't want consumers to import.
I commit `go.mod`."

#### Indepth
**v2+ Directories**. If you release v2.0.0, Go Modules requires the import path to end in `/v2`. `github.com/me/mylib/v2`. You can do this by handling the `v2` folder *inside* the repo, or branching. The "subdirectory strategy" (putting `go.mod` in `v2/` folder) is the most compatible with the ecosystem.

---

### 887. How do you build CLI apps with Cobra?
"`cobra init myapp`.
`cobra add serve`.
I define flags in `init()`: `serveCmd.Flags().Int(...)`.
I implement `Run: func(cmd, args) { ... }`.
Cobra handles help generation, flag parsing, and subcommands hierarchy automatically."

#### Indepth
**Viper**. Cobra works best with Viper (Config). Bind flags to keys: `viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))`. Now `viper.GetInt("port")` returns the Flag value if set, OR the Env Var, OR the Config File value. This "Cascading Configuration" is standard for robust CLIs.

---

### 888. How do you add auto-completion to CLI tools?
"Cobra does it for me.
`rootCmd.GenBashCompletion(os.Stdout)`.
For dynamic completion (e.g., completing Hostnames from SSH config), I use `RegisterFlagCompletionFunc`.
It allows the user to hit TAB and see real-time suggestions."

#### Indepth
**Hidden Command**. Cobra generates a hidden `__complete` command that the shell calls. `myapp __complete serve --po[TAB]`. The binary runs, sees the incomplete flag, calculates suggestions, and prints them to stdout. The shell script captures this and displays it to the user. Fast and magic.

---

### 889. How do you handle subcommands in CLI tools?
"`myapp user create`.
`userCmd` is added to `rootCmd`.
`createCmd` is added to `userCmd`.
Each command has its own `Run` function.
I use persistent flags (`Parent()`) if I want `--verbose` to apply to all subcommands."

#### Indepth
**Traversal**. `cmd.Execute()` finds the leaf command. If you run `app user create`, it traverses `root` -> `user` -> `create`. Middleware (PreRun/PostRun) runs at each level. Use `PersistentPreRun` on Root to set up logging/config for *every* subcommand globally.

---

### 890. How do you package Go binaries for multiple platforms?
"**GoReleaser**.
It runs cross-compilation loops.
It zips them.
It creates checksums.
It uploads them to GitHub.
It minimizes manual error in the release process."

#### Indepth
**NFPM**. GoReleaser uses **nFPM** to generate `.deb` (Debian/Ubuntu) and `.rpm` (RedHat) packages from the binary. It adds systemd unit files (`/etc/systemd/system/myapp.service`) automatically. This allows users to `apt-get install myapp` instead of manually moving binaries to `/usr/local/bin`.

---

### 891. How do you write a Wasm frontend in Go?
"I write a `main()` function.
`//go:build js && wasm`.
I use `syscall/js`.
`dom := js.Global().Get("document")`.
`dom.Call("getElementById", "app").Set("innerHTML", "Hello Wasm")`.
I compile with `GOOS=js GOARCH=wasm go build`."

#### Indepth
**DOM Cost**. Calling JavaScript from WebAssembly is expensive (overhead of crossing the boundary). Don't do `for i < 1000 { setPixel() }`. It will freeze the browser. Build a buffer in Go memory, pass the pointer to JS once, and let JS update the canvas in bulk.

---

### 892. How do you expose Go functions to JS using Wasm?
"`js.Global().Set("myGoFunc", js.FuncOf(func(this js.Value, args []js.Value) any { ... }))`.
This attaches the Go function to the browser's `window` object.
I must keep the Go program running (select{}) so the callback remains active."

#### Indepth
**KeepAlive**. If `main()` exits, the Wasm instance dies. All callbacks (`js.FuncOf`) become invalid and throw errors if called by JS. You must hold the Go runtime open. `<-make(chan struct{})` blocks forever. Also, `js.Func` creates resources. You must `func.Release()` them when done to avoid memory leaks.

---

### 893. How do you reduce Wasm binary size?
"Standard Go Wasm is ~2MB.
1.  Use **TinyGo**. Compiles to ~10KB.
2.  Use `gzip` / `brotli` on the web server.
3.  Strip debug symbols (`-ldflags="-s -w"`)."

#### Indepth
**TinyGo Constraints**. TinyGo is great but it doesn't support the full Go stdlib (especially `reflect`, `net/http` server, `encoding/json` is limited). It uses a different memory allocator. If your code uses heavy reflection (like `fmt.Sprintf` complex structs), TinyGo might fail or panic. Test compatibility early.

---

### 894. How do you interact with DOM from Go Wasm?
"It’s verbose via `syscall/js`.
I use a wrapper library like `vecty` or `go-app`.
They provide a virtual DOM (React-like) experience in Go.
`return elem.Div(elem.Text("Click me"))`."

#### Indepth
**Batching**. Go-app and Vecty use a "Virtual DOM" (VDOM) in Go. They diff the tree and apply only the changes to the real JS DOM. This batches the expensive JS calls. It makes Go Wasm apps responsive enough for Single Page Applications (SPAs) despite the call overhead.

---

### 895. How do you debug Go WebAssembly apps?
"I can't use Delve in the browser easily.
I use `fmt.Println`, which goes to the Browser Console.
For logic testing, I write standard Go unit tests (running on host OS) before compiling to WASM."

#### Indepth
**Source Maps**. Go 1.24+ might improve this, but currently, debugging Wasm is hard. Chrome DevTools sees binary instructions. Some tools can generate Source Maps to map Wasm offsets back to lines in `main.go`. Without this, "panic at PC 0x123" is your only clue. Use extensive Logging.

---

### 896. How do you build a WebAssembly module loader?
"Go provides `wasm_exec.js`.
I include it in my HTML.
`const go = new Go();`.
`WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(...)`.
`go.run(result.instance)`.
This boots the Go runtime inside the browser."

#### Indepth
**Polyfills**. The `wasm_exec.js` provided by Go must match the Go version used to compile. If you upgrade Go 1.21 -> 1.22, you MUST update the `.js` file in your web server. Otherwise, the ABI mismatch will crash the Wasm start process immediately with cryptic errors.

---

### 897. How do you manage state in Go WebAssembly apps?
"Just like a backend app.
Global structs or State Managers.
Since Wasm is single-threaded (in the browser UI thread), I don't need mutexes for UI state."

#### Indepth
**LocalStorage**. Wasm memory is volatile (cleared on refresh). To persist state, use `localStorage`. Write a helper `Save(key, json)`. Call it on every state change. On startup, `Load(key)`. This gives a "Native App" feel where the user returns to exactly where they left off.

---

### 898. How do you integrate Go Wasm with JS promises?
"Go blocks, JS is async.
To await a JS Promise in Go:
Create a channel.
Pass a callback to the Promise `.then()`.
In callback, send to channel.
Go: `<-channel`.
This bridges the async gap."

#### Indepth
**Async/Await**. Go doesn't have `await`. But it has Channels. You can write a helper `Await(promise) (js.Value, error)` that blocks the goroutine until the promise resolves. This makes calling `fetch` API look synchronous in Go: `resp := Await(jsFetch(url))`. Much cleaner than callback hell.

---

### 899. How do you decide between Go CLI and REST tool?
"**CLI**: For operators, scripting.
**REST**: For machine-to-machine.
I often build **Both**.
The Core Logic is a library.
The REST API imports Core.
The CLI imports Core."

#### Indepth
**Layered Arch**. Separate `cmd/` (interface) from `internal/app` (logic). `cmd/web` calls `app.CreateUser`. `cmd/cli` calls `app.CreateUser`. If you put logic in `http.Handler`, the CLI can't reuse it. Logic should accept Go structs, not `http.Request`.

---

### 900. How do you document CLI help and usage info?
"In Cobra:
`Use: "copy [source] [dest]"`.
`Short: "Copies files"`.
`Long: "A robust copy tool..."`.
Cobra auto-generates the `--help` output from these strings."

#### Indepth
**Markdown Docs**. Cobra can generate full Markdown documentation trees. `doc.GenMarkdownTree(rootCmd, "./docs")`. This creates `app.md`, `app_serve.md` etc. You can upload these directly to GitHub Wiki or Docusaurus. It keeps your website documentation in sync with your binary's actual flags.
