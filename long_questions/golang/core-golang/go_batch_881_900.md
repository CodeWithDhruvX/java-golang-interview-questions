## 🧰 Refactoring, CLI, WebAssembly & Design (Questions 881-900)

### Question 881: How do you refactor large Go codebases safely?

**Answer:**
1.  **Tests First:** Ensure high coverage.
2.  **Type Types:** Use `type Alias = old.Type` to move types gradually.
3.  **Interfaces:** Introduce interface to decouple boundaries.
4.  **Tools:** `gopls` (Rename symbol), `gofmt -r` (Rewrite rules).

### Explanation
Safe refactoring of large Go codebases requires ensuring high test coverage first, using type aliases for gradual type migration, introducing interfaces to decouple boundaries, and using tools like gopls for symbol renaming and gofmt -r for rewrite rules.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you refactor large Go codebases safely?
**Your Response:** "I refactor large Go codebases safely by following a systematic approach. First, I ensure high test coverage - tests are my safety net during refactoring. Second, I use type aliases like `type Alias = old.Type` to move types gradually without breaking existing code. Third, I introduce interfaces at boundaries to decouple components and make refactoring easier. Fourth, I use tools like `gopls` for safe symbol renaming and `gofmt -r` for automated rewrite rules. I refactor in small increments, running tests after each change. The key is to maintain the ability to compile and test at every step. I also use feature flags or conditional compilation when making larger changes. This incremental approach minimizes risk and allows me to catch issues early."

---

### Question 882: How do you break a monolith Go app into microservices?

**Answer:**
Identify bounded contexts.
1.  Isolate `package billing`.
2.  Define inputs/outputs strictly.
3.  Replace direct function calls with Interface calls.
4.  Swap Interface impl with gRPC client.
5.  Move code to new service.

### Explanation
Breaking monolith Go apps into microservices involves identifying bounded contexts, isolating packages like billing, defining strict inputs/outputs, replacing direct function calls with interface calls, swapping interface implementations with gRPC clients, and moving code to new services.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you break a monolith Go app into microservices?
**Your Response:** "I break monoliths into microservices by identifying bounded contexts first - finding natural boundaries like billing, users, or orders. Then I isolate a package like `package billing` and define its inputs and outputs strictly. I replace direct function calls with interface calls to create a clear contract. Next, I swap the interface implementation with a gRPC client while keeping the same interface. Finally, I move the actual implementation to a new service. This incremental approach allows me to extract services one at a time without breaking the entire system. The key is maintaining the same interface contract while changing the underlying implementation. I can test each step before moving to the next, which minimizes risk. This pattern works well for extracting any bounded context from a monolith."

---

### Question 883: How do you improve code readability in Go?

**Answer:**
- **Short Variable Names:** `i`, `ctx` (local) vs `TimeoutSeconds` (exported).
- **Happy Path Left:** Return errors early (Guard Clauses). Avoid `else`.
- **Small Interfaces:** `Reader` (1 method) > `BigInterface` (20 methods).

### Explanation
Code readability in Go improves with short variable names for local scope vs descriptive names for exported fields, keeping happy path left by returning errors early with guard clauses, and preferring small interfaces with single methods over large interfaces with many methods.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you improve code readability in Go?
**Your Response:** "I improve Go code readability through three main principles. First, I use short variable names like `i` or `ctx` for local scope, but descriptive names like `TimeoutSeconds` for exported fields. Second, I keep the happy path on the left by returning errors early using guard clauses - this avoids deep nesting and makes the main logic clearer. Third, I prefer small interfaces with single methods like `Reader` over big interfaces with twenty methods. Small interfaces are easier to understand, test, and implement. I also follow Go conventions like organizing imports, using consistent formatting, and keeping functions focused on a single responsibility. The goal is code that's easy to read and understand at a glance. These patterns make the code self-documenting and reduce the cognitive load for anyone reading it."

---

### Question 884: How do you organize domain-driven projects in Go?

**Answer:**
`internal/domain` (Pure Go code, entities).
`internal/service` (Business logic).
`internal/repository` (Data access).
`internal/handler` (HTTP/gRPC).

### Explanation
Domain-driven Go projects organize with internal/domain for pure Go entities, internal/service for business logic, internal/repository for data access, and internal/handler for HTTP/gRPC endpoints, following clean architecture principles.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you organize domain-driven projects in Go?
**Your Response:** "I organize domain-driven Go projects using clean architecture principles. I put pure Go entities and business rules in `internal/domain` - this contains the core business logic without external dependencies. Business logic and use cases go in `internal/service` - this orchestrates the domain entities. Data access code lives in `internal/repository` - this handles database operations and external data sources. HTTP and gRPC handlers go in `internal/handler` - this manages the API layer. This separation keeps the core business logic independent of infrastructure concerns. The `internal` directory ensures these packages aren't imported from outside the module. I might also have `internal/infrastructure` for external services and `pkg` for shared utilities. This structure makes the code testable, maintainable, and follows dependency inversion principles."

---

### Question 885: How do you handle circular dependencies?

**Answer:**
Go forbids A -> B -> A.
**Fix:**
1.  Move shared interface to `package common` or `package types`.
2.  Dependency Inversion: A defines interface, B implements it, Main wires them.

### Explanation
Circular dependencies in Go are forbidden (A imports B which imports A). Solutions include moving shared interfaces to a common/types package, or using dependency inversion where A defines the interface, B implements it, and main wires them together.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle circular dependencies?
**Your Response:** "Go prevents circular dependencies at compile time, which is actually a good thing for clean architecture. When I encounter circular imports like A importing B which imports A, I have two main solutions. First, I move shared interfaces or types to a separate `package common` or `package types` that both packages can import without circularity. Second, I use dependency inversion - package A defines the interface, package B implements it, and the main package wires them together. This follows the dependency inversion principle where high-level modules don't depend on low-level modules. Sometimes I also extract the shared functionality into a third package. The key is organizing packages by dependency direction - dependencies should flow one way. This constraint forces better design and prevents tangled dependencies that are hard to maintain."

---

### Question 886: How do you structure reusable Go modules?

**Answer:**
- **Pkg Layout:** Root contains library code. `cmd` contains binaries (if any).
- **SemVer:** Tag `v1.0.0`.
- **Minimal Deps:** Don't import heavy framework if you just need a string util.

### Explanation
Reusable Go modules structure with library code in root, binaries in cmd directory, use Semantic Versioning tags like v1.0.0, and minimize dependencies by avoiding heavy frameworks for simple utilities like string operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you structure reusable Go modules?
**Your Response:** "I structure reusable Go modules following Go conventions. The root package contains the main library code, and any command-line tools go in the `cmd` directory. I use semantic versioning with tags like `v1.0.0` to communicate breaking changes. Most importantly, I keep dependencies minimal - if I'm building a string utility, I won't import a heavy framework. This makes the module lightweight and easy to use. I also include a comprehensive README with examples, clear documentation in the code, and a proper go.mod file with a clean module path. For public modules, I ensure the API is stable and well-tested. The goal is to create something that's easy to import and use without pulling in unnecessary dependencies. This approach follows Go's philosophy of simplicity and composability."

---

### Question 887: How do you build CLI apps with Cobra?

**Answer:**
(See Q442/Q445).
`rootCmd.AddCommand(childCmd)`.
Use `init()` to register flags.

### Explanation
Cobra CLI building uses rootCmd.AddCommand(childCmd) to create command hierarchies, and init() functions to register flags. Cobra provides a framework for building powerful command-line applications with subcommands and flags.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build CLI apps with Cobra?
**Your Response:** "I build CLI apps with Cobra by creating a root command and adding child commands using `rootCmd.AddCommand(childCmd)`. Each command is a struct with methods for `Run`, `PreRun`, and so on. I use `init()` functions to register flags for each command - Cobra handles all the flag parsing and help generation automatically. The structure is hierarchical - I can have commands like `app server start` where `server` is a child of `app` and `start` is a child of `server`. Cobra also provides automatic help generation, shell completion, and validation. I organize each command in its own file to keep the code clean. The framework makes it easy to build professional CLI tools with consistent behavior. This pattern scales well from simple single-command tools to complex multi-command applications."

---

### Question 888: How do you add auto-completion to CLI tools?

**Answer:**
(See Q799). Cobra's `GenBashCompletion`.

### Explanation
CLI auto-completion uses Cobra's GenBashCompletion function to generate shell completion scripts. This provides tab completion for commands, flags, and arguments in bash, zsh, fish, and other shells.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you add auto-completion to CLI tools?
**Your Response:** "I add auto-completion to CLI tools using Cobra's built-in completion system. Cobra provides `GenBashCompletion` and similar functions for different shells. I call these functions to generate completion scripts that users can source in their shell configuration. The completion automatically handles commands, subcommands, flags, and even dynamic completion for things like file names or custom options. For example, if I have a command that takes a filename as an argument, I can provide a completion function that suggests files. Users just need to source the completion script once, and they get tab completion for free. This makes CLI tools much more user-friendly and professional. Cobra handles all the complexity of shell completion syntax, so I just need to provide the completion logic for my specific use cases."

---

### Question 889: How do you handle subcommands in CLI tools?

**Answer:**
Cobra handles tree traversal.
Parsing logic automatically routes `app server start` to the `start` handler function.

### Explanation
Cobra handles CLI subcommands through tree traversal, automatically parsing and routing commands like `app server start` to the appropriate handler function based on the command hierarchy.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle subcommands in CLI tools?
**Your Response:** "Cobra handles subcommands automatically through tree traversal. When I define a command hierarchy like `app server start`, Cobra builds a tree where `app` is the root, `server` is a child, and `start` is a grandchild. When a user types `app server start`, Cobra parses the arguments and walks the tree to find the right handler function. I don't need to write any parsing logic - Cobra handles all the routing automatically. Each command can have its own flags and arguments, and Cobra ensures they're validated and routed correctly. This makes building complex CLI tools with multiple subcommands straightforward. The tree structure also scales well - I can nest commands as deeply as needed. Cobra's approach means I can focus on the actual command logic rather than argument parsing and routing."

---

### Question 890: How do you package Go binaries for multiple platforms?

**Answer:**
`GOOS=windows GOARCH=amd64 go build`.
Zip the result.
Use **GoReleaser** to automate zipping + checksums + upload.

### Explanation
Multi-platform Go binary packaging uses GOOS and GOARCH environment variables for cross-compilation, zipping the results, and GoReleaser for automation of zipping, checksums, and uploads to distribution platforms.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you package Go binaries for multiple platforms?
**Your Response:** "I package Go binaries for multiple platforms using Go's cross-compilation capabilities. I set environment variables like `GOOS=windows GOARCH=amd64 go build` to compile for different operating systems and architectures. Go can compile for dozens of platforms from a single machine. After building, I zip the binaries with any necessary assets. For automation, I use GoReleaser which handles the entire process - it builds for multiple platforms, creates zip archives, generates checksums for security, and can even upload to GitHub releases or other distribution platforms. I configure GoReleaser with a YAML file that specifies the build matrix, archive formats, and release targets. This makes releasing new versions with binaries for all platforms a single command. The approach is much simpler than setting up separate build environments for each platform."

---

### Question 891: How do you write a Wasm frontend in Go?

**Answer:**
Compile with `GOOS=js GOARCH=wasm`.
Use `syscall/js`.
Manipulate DOM: `js.Global().Get("document").Call("getElementById", "app")`.

### Explanation
Go WebAssembly frontend development compiles with GOOS=js GOARCH=wasm, uses syscall/js package for JavaScript interop, and manipulates DOM through js.Global() calls to access browser APIs like document.getElementById.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write a Wasm frontend in Go?
**Your Response:** "I write WebAssembly frontends in Go by compiling with `GOOS=js GOARCH=wasm` which targets the WebAssembly architecture. I use the `syscall/js` package to interact with JavaScript and browser APIs. To manipulate the DOM, I use `js.Global().Get("document").Call("getElementById", "app")` to access browser objects and call their methods. The process feels similar to writing regular Go code, but instead of calling Go libraries, I'm calling JavaScript APIs through the js bridge. I can handle events, update UI elements, and manage application state all from Go. The compilation produces a .wasm file that I load in the browser along with the wasm_exec.js loader. This approach lets me write web frontends using Go's type safety and familiar syntax while still running in the browser."

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

### Explanation
Exposing Go functions to JavaScript in WebAssembly uses js.Global().Set() to attach functions to the global scope, js.FuncOf() to wrap Go functions as JavaScript-callable functions, and select{} to prevent the Go program from exiting.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you expose Go functions to JS using Wasm?
**Your Response:** "I expose Go functions to JavaScript using the `syscall/js` package. I wrap a Go function with `js.FuncOf()` which makes it callable from JavaScript, then attach it to the global scope with `js.Global().Set()`. The function receives JavaScript values as arguments and can return Go values that get converted to JavaScript. I use `select{}` at the end to prevent the Go program from exiting immediately. From JavaScript, I can then call `window.myGoFunc()` just like any regular JavaScript function. This bridge works both ways - I can also call JavaScript functions from Go. The conversion between Go and JavaScript types happens automatically for basic types. This approach lets me expose Go business logic to the frontend while keeping the implementation in Go."

---

### Question 893: How do you reduce Wasm binary size?

**Answer:**
(See Q483). Use **TinyGo**.
Or gzip/brotli compression on the server (Go WASM reduces from 2MB to 500KB compressed).

### Explanation
WebAssembly binary size reduction uses TinyGo for smaller binaries or server-side compression with gzip/brotli. Go WebAssembly reduces from 2MB to 500KB compressed, making it more practical for web deployment.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you reduce Wasm binary size?
**Your Response:** "I reduce WebAssembly binary size in two main ways. First, I can use TinyGo instead of the standard Go compiler - TinyGo is specifically designed for embedded systems and WebAssembly, producing much smaller binaries. Second, I use server-side compression with gzip or brotli - a Go WebAssembly binary that starts at 2MB compresses down to about 500KB. Most web servers handle compression automatically, so users get the compressed version. I also optimize by removing unused imports and keeping the code minimal. The combination of TinyGo compilation and compression makes Go WebAssembly much more practical for web applications. While still larger than equivalent JavaScript, the size becomes manageable for many use cases, especially when considering the benefits of Go's type safety and performance."

---

### Question 894: How do you interact with DOM from Go Wasm?

**Answer:**
It's slow (Bridge crossing overhead).
Batch DOM updates if possible.
Use frameworks like **Vecty** or **Go-App** for React-like components in Go.

### Explanation
DOM interaction from Go WebAssembly is slow due to bridge crossing overhead between Go and JavaScript. Best practices include batching DOM updates and using frameworks like Vecty or Go-App for React-like component patterns in Go.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you interact with DOM from Go Wasm?
**Your Response:** "DOM interaction from Go WebAssembly has performance overhead due to the bridge between Go and JavaScript. Each DOM operation requires crossing this bridge, which is relatively expensive. I mitigate this by batching DOM updates - making multiple changes in a single operation when possible. For more complex applications, I use frameworks like Vecty or Go-App that provide React-like component patterns in Go. These frameworks optimize DOM updates and handle the bridge crossing more efficiently. They also provide virtual DOM diffing and other optimizations. For simple applications, direct DOM manipulation through `syscall/js` works fine, but for anything complex, a framework helps manage the performance overhead. The key is to minimize the number of bridge crossings and batch operations when possible."

---

### Question 895: How do you debug Go WebAssembly apps?

**Answer:**
Console.log (`fmt.Println` redirects to browser console).
Newer Chrome DevTools allow debugging C++/Rust/Go via Source Maps (Dwarf), but support is experimental.

### Explanation
Go WebAssembly debugging uses console.log (fmt.Println redirects to browser console) and newer Chrome DevTools support debugging via Source Maps with Dwarf debug information, though support is still experimental.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you debug Go WebAssembly apps?
**Your Response:** "I debug Go WebAssembly apps primarily using console logging - `fmt.Println` automatically redirects to the browser console, so I can see Go output alongside JavaScript logs. For more advanced debugging, newer Chrome DevTools support debugging WebAssembly through Source Maps with Dwarf debug information, though this support is still experimental. I can set breakpoints and step through Go code in the browser debugger. I also use browser developer tools to inspect the WebAssembly memory and performance. The debugging experience is improving but still not as seamless as regular Go development. I often combine logging with careful testing to catch issues. Some developers also use unit tests that run in Node.js with WebAssembly support for easier debugging. The ecosystem is still evolving, but basic debugging is definitely possible."

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

### Explanation
WebAssembly module loading uses the provided wasm_exec.js file which creates a Go runtime instance, fetches the .wasm file, and runs it using WebAssembly.instantiateStreaming with the Go import object.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a WebAssembly module loader?
**Your Response:** "I build WebAssembly module loaders using the provided `wasm_exec.js` file that comes with Go. This script sets up the Go runtime environment in the browser. I create a new Go instance, fetch the WebAssembly file using `WebAssembly.instantiateStreaming`, and then run it with `go.run()`. The loader handles all the complexity of setting up the Go runtime, memory management, and the bridge between Go and JavaScript. I typically wrap this in an async function and handle any loading errors. The loader also provides functions for calling Go functions from JavaScript. This approach works consistently across browsers and handles all the WebAssembly initialization details. I can customize the loading process, add progress indicators, or handle different WebAssembly modules as needed."

---

### Question 897: How do you manage state in Go WebAssembly apps?

**Answer:**
Just like a normal Go app (Structs in memory).
State persists as long as the tab doesn't reload.
To persist: `js.Global().Get("localStorage").Call("setItem", ...)`

### Explanation
State management in Go WebAssembly works like normal Go apps with structs in memory. State persists as long as the tab remains open. For persistence across sessions, use localStorage through js.Global().Get("localStorage").Call("setItem").

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage state in Go WebAssembly apps?
**Your Response:** "I manage state in Go WebAssembly apps just like in regular Go applications - using structs and variables in memory. The state persists as long as the browser tab stays open. For persistence across sessions, I use browser localStorage through the JavaScript bridge with `js.Global().Get('localStorage').Call('setItem', key, value)`. I can create complex state management with structs, maps, and slices just like in server-side Go. The main difference is that all state lives in the browser's memory space. For more complex applications, I might implement state management patterns similar to React or Redux, but using Go structs and methods. I can also synchronize state with a backend server if needed. The key is understanding that WebAssembly runs in a sandboxed environment, so all state is contained within that context unless I explicitly persist it."`

---

### Question 898: How do you integrate Go Wasm with JS promises?

**Answer:**
Go function cannot block JS event loop.
Return a Promise from Go? Hard.
Usually: Go calls a JS callback function passed as an argument when work is done.

### Explanation
Go WebAssembly integration with JavaScript promises is challenging because Go functions cannot block the JavaScript event loop. Instead of returning promises from Go, the typical pattern is Go calling a JavaScript callback function passed as an argument when work completes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you integrate Go Wasm with JS promises?
**Your Response:** "Integrating Go WebAssembly with JavaScript promises is tricky because Go functions can't block the JavaScript event loop. Returning a Promise directly from Go is difficult. Instead, I use a callback pattern where JavaScript passes a callback function to Go, and Go calls that function when the work is done. For example, JavaScript might call a Go function with a callback, and Go does its work asynchronously and then calls the callback with the result. This avoids blocking the event loop while still providing asynchronous behavior. Some developers wrap this pattern to create promise-like APIs, but the underlying mechanism is callbacks. The key is understanding that WebAssembly runs on the same thread as JavaScript, so long-running Go operations need to be non-blocking. I might use goroutines for background work and then call the callback when complete."

---

### Question 899: How do you decide between Go CLI and REST tool?

**Answer:**
- **CLI:** Interactive, Scriptable, Local Files access.
- **REST:** Remote access, integrations.
**Start with Core Logic (Library).** Then wrap it in BOTH a CLI `cmd/cli` and Server `cmd/server`.

### Explanation
Choosing between CLI and REST tools: CLI for interactive use, scripting, and local file access; REST for remote access and integrations. Best practice is to start with core logic as a library, then wrap it in both CLI and server interfaces.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you decide between Go CLI and REST tool?
**Your Response:** "I decide between CLI and REST tools based on the use case. CLI tools are better for interactive use, scripting, and when I need direct access to local files. REST APIs are better for remote access and integrations with other services. The best approach is to start with the core logic as a library, then wrap it in both a CLI interface in `cmd/cli` and a REST server in `cmd/server`. This gives me the best of both worlds - users can interact directly via CLI or integrate via REST. The shared library ensures both interfaces have the same functionality and behavior. I might also add other interfaces like gRPC or WebUI later. This pattern separates concerns well and makes the code more reusable. The decision isn't always either/or - often both interfaces are valuable for different use cases."

---

### Question 900: How do you document CLI help and usage info?

**Answer:**
Cobra auto-generates help.
Short/Long descriptions.
Example usage strings.
`app --help` shows flags and subcommands.

### Explanation
CLI documentation in Cobra uses auto-generated help with short and long descriptions, example usage strings, and app --help to display flags and subcommands. Cobra automatically formats and displays comprehensive help information.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you document CLI help and usage info?
**Your Response:** "I document CLI help and usage using Cobra's built-in help system. I provide short and long descriptions for each command, add example usage strings to show common patterns, and Cobra automatically generates comprehensive help documentation. When users run `app --help`, they see all available flags, subcommands, and usage examples. I can also add custom help templates if needed. The help system automatically includes flag descriptions, default values, and even validates flag usage. This approach ensures consistent, professional-looking documentation across all commands. I make sure to write clear, concise descriptions and include practical examples. The auto-generated help means I don't have to maintain separate documentation files - the help is always in sync with the actual command implementation."

---
