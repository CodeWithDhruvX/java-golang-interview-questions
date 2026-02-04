## 🧰 Refactoring, CLI, WebAssembly & Design (Questions 881-900)

### Question 881: How do you refactor large Go codebases safely?

**Answer:**
1.  **Tests First:** Ensure high coverage.
2.  **Type Types:** Use `type Alias = old.Type` to move types gradually.
3.  **Interfaces:** Introduce interface to decouple boundaries.
4.  **Tools:** `gopls` (Rename symbol), `gofmt -r` (Rewrite rules).

---

### Question 882: How do you break a monolith Go app into microservices?

**Answer:**
Identify bounded contexts.
1.  Isolate `package billing`.
2.  Define inputs/outputs strictly.
3.  Replace direct function calls with Interface calls.
4.  Swap Interface impl with gRPC client.
5.  Move code to new service.

---

### Question 883: How do you improve code readability in Go?

**Answer:**
- **Short Variable Names:** `i`, `ctx` (local) vs `TimeoutSeconds` (exported).
- **Happy Path Left:** Return errors early (Guard Clauses). Avoid `else`.
- **Small Interfaces:** `Reader` (1 method) > `BigInterface` (20 methods).

---

### Question 884: How do you organize domain-driven projects in Go?

**Answer:**
`internal/domain` (Pure Go code, entities).
`internal/service` (Business logic).
`internal/repository` (Data access).
`internal/handler` (HTTP/gRPC).

---

### Question 885: How do you handle circular dependencies?

**Answer:**
Go forbids A -> B -> A.
**Fix:**
1.  Move shared interface to `package common` or `package types`.
2.  Dependency Inversion: A defines interface, B implements it, Main wires them.

---

### Question 886: How do you structure reusable Go modules?

**Answer:**
- **Pkg Layout:** Root contains library code. `cmd` contains binaries (if any).
- **SemVer:** Tag `v1.0.0`.
- **Minimal Deps:** Don't import heavy framework if you just need a string util.

---

### Question 887: How do you build CLI apps with Cobra?

**Answer:**
(See Q442/Q445).
`rootCmd.AddCommand(childCmd)`.
Use `init()` to register flags.

---

### Question 888: How do you add auto-completion to CLI tools?

**Answer:**
(See Q799). Cobra's `GenBashCompletion`.

---

### Question 889: How do you handle subcommands in CLI tools?

**Answer:**
Cobra handles tree traversal.
Parsing logic automatically routes `app server start` to the `start` handler function.

---

### Question 890: How do you package Go binaries for multiple platforms?

**Answer:**
`GOOS=windows GOARCH=amd64 go build`.
Zip the result.
Use **GoReleaser** to automate zipping + checksums + upload.

---

### Question 891: How do you write a Wasm frontend in Go?

**Answer:**
Compile with `GOOS=js GOARCH=wasm`.
Use `syscall/js`.
Manipulate DOM: `js.Global().Get("document").Call("getElementById", "app")`.

---

### Question 892: How do you expose Go functions to JS using Wasm?

**Answer:**
```go
js.Global().Set("myGoFunc", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    return "Hello from Go"
}))
// Prevent exit
select {}
```
In JS: `window.myGoFunc()`.

---

### Question 893: How do you reduce Wasm binary size?

**Answer:**
(See Q483). Use **TinyGo**.
Or gzip/brotli compression on the server (Go WASM reduces from 2MB to 500KB compressed).

---

### Question 894: How do you interact with DOM from Go Wasm?

**Answer:**
It's slow (Bridge crossing overhead).
Batch DOM updates if possible.
Use frameworks like **Vecty** or **Go-App** for React-like components in Go.

---

### Question 895: How do you debug Go WebAssembly apps?

**Answer:**
Console.log (`fmt.Println` redirects to browser console).
Newer Chrome DevTools allow debugging C++/Rust/Go via Source Maps (Dwarf), but support is experimental.

---

### Question 896: How do you build a WebAssembly module loader?

**Answer:**
Use the provided `wasm_exec.js`.
```js
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
```

---

### Question 897: How do you manage state in Go WebAssembly apps?

**Answer:**
Just like a normal Go app (Structs in memory).
State persists as long as the tab doesn't reload.
To persist: `js.Global().Get("localStorage").Call("setItem", ...)`

---

### Question 898: How do you integrate Go Wasm with JS promises?

**Answer:**
Go function cannot block JS event loop.
Return a Promise from Go? Hard.
Usually: Go calls a JS callback function passed as an argument when work is done.

---

### Question 899: How do you decide between Go CLI and REST tool?

**Answer:**
- **CLI:** Interactive, Scriptable, Local Files access.
- **REST:** Remote access, integrations.
**Start with Core Logic (Library).** Then wrap it in BOTH a CLI `cmd/cli` and Server `cmd/server`.

---

### Question 900: How do you document CLI help and usage info?

**Answer:**
Cobra auto-generates help.
Short/Long descriptions.
Example usage strings.
`app --help` shows flags and subcommands.

---
