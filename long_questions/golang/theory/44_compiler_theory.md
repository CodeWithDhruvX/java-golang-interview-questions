# 🧠 Go Theory Questions: 861–880 Go Compiler & Language Theory

## 861. How does `defer` work at the bytecode level?

**Answer:**
Historically: `defer` allocated a struct on the heap (slow).
Go 1.14+: **Open Coded Defer**.
The compiler inlines the deferred call at every exit point of the function.
`defer f()` -> Compiler adds `f()` before every `return`.
If there's a loop or conditional defer, it falls back to the runtime stack-based implementation, but for standard top-level defers, cost is near zero.

---

## 862. What is the Go frontend written in?

**Answer:**
Go.
Since Go 1.5 (2015), the compiler was self-hosted (translated from C to Go).
The parser uses `cmd/compile/internal/syntax`.
The type checker uses `cmd/compile/internal/types2`.
This makes the compiler easier to hack on for Go developers, but introduced the "Bootstrap" problem (you need Go 1.4 to build Go 1.5).

---

## 863. How are interfaces implemented in memory?

**Answer:**
An interface is a 2-word struct `(Type, Value)`.
1.  **Tab**: Pointer to `itab` (Interface Table). Contains type metadata and function pointers for the specific methods required by the interface.
2.  **Data**: Pointer to the actual data (concrete struct).
Calling `i.Method()` works by fetching index K from the `itab` funcs array and passing `Data` as the receiver.

---

## 864. What are method sets and how do they affect interfaces?

**Answer:**
Rules:
- `T` has methods with receiver `(t T)`.
- `*T` has methods with receiver `(t T)` AND `(t *T)`.
**Implication**:
If you verify an interface with a pointer receiver method (`func (u *User) Save()`), you **cannot** assign a value `User{}` to that interface. You must assign `&User{}`.
Values are not addressable in all contexts, so Go forbids value-to-pointer-method promotion in interface assignment.

---

## 865. How do you implement AST manipulation in Go?

**Answer:**
`go/parser`, `go/ast`, `go/printer`.
1.  Parse: `fset := token.NewFileSet(); node, _ := parser.ParseFile(...)`.
2.  Walk: `ast.Inspect(node, func(n ast.Node) bool { ... })`.
3.  Modify: `ident.Name = "NewName"`.
4.  Print: `printer.Fprint(output, fset, node)`.
This is how `gofmt` and `gomodifytags` work.

---

## 866. What is the Go toolchain pipeline from source to binary?

**Answer:**
1.  **Parsing**: Source -> AST.
2.  **Type Checking**: AST -> Typed AST (IR).
3.  **SSA Generation**: IR -> SSA (Static Single Assignment).
4.  **Optimization**: Dead code, BCE, Devirtualization.
5.  **Lowering**: SSA -> Machine-specific Assembly (Plan9).
6.  **Assembler**: Assembly -> Object Files (`.o`).
7.  **Pack**: Object Files -> Archive (`.a`).
8.  **Link**: Archives -> Executable Binary.

---

## 867. How are function closures handled by the Go compiler?

**Answer:**
The compiler detects "Capturing".
If a variable inside a function is accessed by a closure returned/passed out, it **Escapes to Heap**.
The closure struct contains:
1.  Function Pointer.
2.  Pointer to the captured variables on the Heap.
This allows the closure to mutate the original variable even after the parent function returns.

---

## 868. What is link-time optimization in Go?

**Answer:**
Go's linker removes **Dead Code** (DCE).
It builds a "Reachability Graph" starting from `main.main`.
If a function/global is never reached, it is not included in the final binary.
This is why `Hello World` is 1MB instead of 50MB (size of whole standard library).
Go doesn't do "LTO" in the C++ sense (cross-module inlining), but it does aggressive dead-code stripping.

---

## 869. How does cgo interact with Go's runtime?

**Answer:**
It switches stacks.
Go runs on 2KB expandable stacks. C runs on large fixed OS stacks.
Calling C:
1.  Runtime pauses M.
2.  Switches to system stack (`g0`).
3.  Executes C code.
4.  Switches back.
This overhead (call switching) is expensive (~150ns). Reducing "Chatty Cgo" (many small calls) is crucial for performance.

---

## 870. What are zero-sized types and how are they used?

**Answer:**
`struct{}` or `[0]int`. Size is 0 bytes.
Internal optimization: `runtime.zerobase`.
All pointers to zero-sized objects point to the *same* address.
Use cases:
1.  **Set**: `map[string]struct{}` (saves memory vs `bool`).
2.  **Signal Channel**: `chan struct{}` (message has no data, just event).

---

## 871. How does type aliasing differ from type definition?

**Answer:**
**Definition**: `type MyInt int`.
`MyInt` is a **New Type**. Need cast to convert to `int`. It can have methods.

**Alias**: `type MyInt = int`.
`MyInt` *IS* `int`.
Interchangeable without cast. Used for refactoring (moving a type from package A to B without breaking clients).

---

## 872. How does Go avoid null pointer dereferencing?

**Answer:**
It doesn't completely. `nil` exists.
But:
1.  **No Pointer Arithmetic**: Can't accidentally access random memory.
2.  **Zero Values**: All variables initialized to standard default (0, "", nil). No uninitialized garbage.
3.  **Nil Receivers**: You *can* call methods on nil pointers if the method handles it: `func (n *Node) String() { if n == nil return "nil" }`. This avoids many common crashes.

---

## 873. What’s the role of go/types package?

**Answer:**
It implements the Go Type System.
Used by Linters (`staticcheck`) and LSP (`gopls`).
It resolves:
- Identifiers to Objects (Vars, Funcs).
- Expressions to Types.
- Computes Method Sets.
It allows tools to "understand" code semantics, not just syntax tokens.

---

## 874. How does Go manage ABI stability?

**Answer:**
Go has *internal* ABI stability (for assembly).
Go 1.17 switched to **Register-based ABI** (passing args in registers instead of stack).
This yielded 5-10% perf gain.
The `//go:linkname` directive lets standard lib breach boundaries, but user code relying on internal runtime details is fragile and often breaks between releases.

---

## 875. How do you refactor large Go codebases safely?

**Answer:**
1.  **gopls**: Rename Symbol (F2).
2.  **Type Warning**: `type ID string` instead of `int` to prevent mixing IDs.
3.  **Mechanical Refactor**: `gofmt -r 'pattern -> replacement'`.
4.  **Complier Help**: Rename a field, run build, fix all errors. The compiler is your best refactoring tool because it catches every usage.

---

## 876. How do you break a monolith Go app into microservices?

**Answer:**
**Strangler Fig Pattern**.
1.  Identify a Bounded Context (e.g., "Billing").
2.  Define Interface in Monolith for Billing.
3.  Create new Go Service "Billing".
4.  Implement Interface in Monolith to call new Service (gRPC).
5.  Switch flag.
6.  Delete old code.
Don't rewrite from scratch. Extract modules one by one.

---

## 877. How do you improve code readability in Go?

**Answer:**
**Idioms over Cleverness**.
1.  **Flat structure**: Avoid `src/java/com/org...`.
2.  **Short variables**: `i`, `ctx`, `r` (Request). Scope is small -> Name is short.
3.  **Early Return**: Reduce nesting. `if err != nil { return err }`.
4.  **Package Names**: `user.Get()` not `user.GetUser()`.

---

## 878. How do you organize domain-driven projects in Go?

**Answer:**
Layered (Clean) Architecture:
1.  **Domain** (Core): Structs, Repository Interfaces. No dependencies.
2.  **Service** (Logic): Use Cases. Depends on Domain.
3.  **Adapter/Port** (Infra): SQL implementation, HTTP Handlers. Depends on Service.
`cmd/main.go` wires them up (Dependency Injection).

---

## 879. How do you handle circular dependencies?

**Answer:**
Go forbids them.
Solution:
1.  **Interface Extraction**: A needs B, B needs A. Create package C with Interface I. A and B both depend on C.
2.  **Merge**: A and B are tightly coupled. Put them in the same package.
3.  **Third Package**: Move common structs to `types` or `model` package.

---

## 880. How do you structure reusable Go modules?

**Answer:**
1.  **Root**: `go.mod`, `README`, generic logic.
2.  **Internal**: Code you don't want others to import (`internal/`).
3.  **Examples**: `examples/` folder (runnable).
4.  **SemVer**: Tag releases `v1.0.0`.
5.  **Doc**: Good comments for `godoc` parsing.
Avoid `common` or `util` packages; they become junk drawers. Be specific (`strutil`, `netutil`).
