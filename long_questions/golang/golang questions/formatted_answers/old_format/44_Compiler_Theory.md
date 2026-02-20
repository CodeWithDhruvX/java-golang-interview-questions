# 🧠 **861–880: Go Compiler & Language Theory**

### 861. How do you build a custom Go compiler plugin?
"Go doesn't support compiler plugins easily.
I have to modify the compiler source (`src/cmd/compile`) and rebuild the toolchain.
Or I use `go/analysis` (Linter) framework to inject checks.
To change *codegen*, I'd need to fork the compiler."

#### Indepth
**RPC Plugins**. Go's native "Plugin" system (`.so` files) is notoriously fragile (requires exact same compiler version, stdlib, dependency versions). The industry standard (HashiCorp, VSCode) is **RPC Plugins**. The plugin is a separate binary. It talks to the main app over gRPC/stdout. This is robust, cross-platform, and language-agnostic.

---

### 862. What is SSA (Static Single Assignment) form in Go?
"It’s the intermediate representation used by the compiler backend.
Every variable is assigned exactly once.
`x = 1; x = 2` becomes `x1 = 1; x2 = 2`.
This simplifies optimization (Dead Code Elimination, Register Allocation) because the compiler knows exactly where every value comes from."

#### Indepth
**Passes**. The Go compiler runs 40+ optimization passes on the SSA. You can see them with `GOSSAFUNC=myFunc go build`. It generates an `ssa.html` file showing how your code evolves from "Source" -> "AST" -> "Lowered SSA" -> "Assembly". This is the ultimate tool for understanding *why* your code is slow.

---

### 863. How does Go handle type inference?
"It’s unidirectional.
`var x = 1` implies `int`.
It doesn't do complex Hindley-Milner bidirectional inference.
It infers from right-hand side to left-hand side.
Inside functions, `:=` works. At package level, must use `var =`."

#### Indepth
**Generic Constraints**. With Generics (Go 1.18), inference got smarter. `func Add[T Number](a, b T)`. Calling `Add(1, 2)` infers `T=int`. Calling `Add(1.0, 2)` fails because `2` (int) doesn't match `1.0` (float). You must explicitly cast `Add(1.0, float64(2))` or rely on the default type of untyped constants.

---

### 864. What is escape analysis in Go?
"The compiler determines if a variable's lifetime exceeds its stack frame.
If yes (returned pointer, passed to interface) -> **Heap**.
If no -> **Stack**.
Stack allocation is O(1). Heap is expensive (GC).
`go build -gcflags '-m'` shows the decisions."

#### Indepth
**Stack Growth**. Go stacks start small (2KB). If escape analysis says "Stack", but the variable is huge (10MB array), it might cause a "Stack Overflow" if the stack couldn't grow. But Go stacks *are* dynamic. They grow (copy themselves to larger memory) automatically. However, copying a stack is expensive. Large objects should usually be on Heap anyway.

---

### 865. How does inlining affect performance in Go?
"It replaces the function call with the function body.
Removes function call overhead (jumps).
Enables further optimizations (Constant Folding).
Cost: Binary size increases.
Go inlines small, leaf functions."

#### Indepth
**Mid-stack Inlining**. Historically, Go only inlined leaf functions (functions that call no one). Now it supports mid-stack inlining (inlining a function that calls another). This dramatically increases the scope of optimization. The "budget" for inlining is roughly 80 AST nodes. `//go:noinline` prevents it.

---

### 866. What are build constraints and how do they work?
"They tell `go build` which files to include.
`//go:build linux && amd64`.
Evaluated at the start of the build.
If false, the file is ignored.
Used for OS-specific code (syscalls) or features (integration tests)."

#### Indepth
**File Suffixes**. You don't always need `//go:build`. The compiler automatically recognizes `foo_linux.go`, `foo_windows_amd64.go`. This is "Implicit Constraints". It's cleaner than adding comment tags to every file. Prefer suffixes for OS/Arch separation, and Tags (`//go:build integration`) for feature flags.

---

### 867. How does `defer` work at the bytecode level?
"Historically: expensive closure allocation on heap.
Now: **Open-coded defer**.
The compiler injects the deferred code directly at the return points.
No allocation. Cost is almost zero.
Unless inside a loop—then it falls back to heap-allocated linked list."

#### Indepth
**Defers in Loops**. A common performance trap is `for { defer f() }`. This allocates a `defer` struct on the heap *every iteration*. It won't be freed until the function returns (which might be "never" in a server loop), causing a memory leak. Wrap the loop body in a `func()` closure to ensure defers run per iteration.

---

### 868. What is the Go frontend written in?
"It’s written in Go.
Originally C (Plan 9 C compiler).
Transpiled to Go in Go 1.5 (The Great Self-Hosting).
Now the entire toolchain is pure Go."

#### Indepth
**Bootstrap**. How do you compile the Go compiler if it's written in Go? You need Go to compile Go. This is the **Bootstrap** problem. To build Go 1.20, you need Go 1.17 installed. The chain goes back to Go 1.4 (the last C version). `dist` is the tool that orchestrates this bootstrap process.

---

### 869. How are interfaces implemented in memory?
"A pair of words `(type, data)`.
`type`: Pointer to `itab` (Interface Table), containing method pointers.
`data`: Pointer to the concrete struct.
Method call: `tab->fun[0](data)`.
This indirection explains why interface calls are slightly slower than direct calls."

#### Indepth
**Itab Caching**. Computing the method table (`itab`) for a dynamic pair `(ConcreteType, InterfaceType)` is expensive (O(N)). Go computes it lazily and caches it in a global hash table. The first time you cast `MyStruct` to `Reader`, it pays the cost. Subsequent times are just a hash lookup.

---

### 870. What are method sets and how do they affect interfaces?
"A type `T` has methods with receiver `(t T)`.
A type `*T` has methods with `(t T)` AND `(t *T)`.
If I fulfill an interface with a pointer method, I *must* pass a pointer.
`var i Iface = MyStruct{}` fails if `MyStruct` only has `func (m *MyStruct) Foo()`."

#### Indepth
**Addressability**. You can call a pointer method on a value: `v.Foo()` (auto-referenced to `(&v).Foo()`). BUT only if `v` is **Addressable**. `MyStruct{}.Foo()` fails because a temporary struct literal is not addressable (it has no memory location yet). You must do `(&MyStruct{}).Foo()`.

---

### 871. How do you implement AST manipulation in Go?
"I use `go/parser`, `go/ast`, `go/token`.
1.  Parse source to AST.
2.  Walk the tree (`ast.Inspect`).
3.  Modify nodes (rename variable).
4.  Print back to source.
This is how tools like `gomvpkg` or `gorename` work."

#### Indepth
**dst vs ast**. The standard `go/ast` destroys comments when modifying nodes (because comments are floating, not attached to nodes). Use `dave/dst` (Decorated Syntax Tree) for rigorous refactoring tools. It attaches comments *to* the nodes as "Decorations" (StartDecs, EndDecs), allowing faithful reproduction of the source code.

---

### 872. What is the Go toolchain pipeline from source to binary?
"1.  **Parsing**: Source -> AST.
2.  **Type Checking**: AST -> Typed AST.
3.  **SSA Generation**: IR generation.
4.  **Optimization**: SSA -> Optimized SSA.
5.  **Code Gen**: SSA -> Assembly (`.o`).
6.  **Linking**: Combine `.o` -> Executable."

#### Indepth
**Object Format**. Go doesn't use standard ELF `.o` files for intermediate objects because they are too slow to parse. It uses a custom highly-optimized object format. The Linker (`cmd/link`) is also custom and optimized for incremental linking, though it is slower than the compiler itself for large projects.

---

### 873. How are function closures handled by the Go compiler?
"If the closure captures variables from the outer scope, those variables must survive.
The compiler moves them to the Heap.
The closure struct contains a pointer to the function code and a pointer to the captured variables (Context)."

#### Indepth
**Function Values**. In Go, `func` is a value. It's technically a pointer to a struct. `type funcVal struct { fn uintptr; args ... }`. When you pass a function `f` to `sort.Slice`, you are passing this pointer. It allows the runtime to distinguish between the same function `Add` called with different captured environments.

---

### 874. What is link-time optimization in Go?
"Go does **Dead Code Elimination** at link time.
If `FuncA` is never called, it’s not included in the binary.
However, if I use `reflect`, the linker gets scared and keeps everything.
Go doesn't do aggressive LTO like C++ LLVM yet."

#### Indepth
**DWARf**. A huge chunk of binary size is Debug Information (DWARF). The linker generates this to allow GDB/Delve to map `0x1234` back to `main.go:55`. Using `go build -ldflags="-s -w"` strips the Symbol Table (-s) and DWARF (-w), reducing binary size by 20-30% at the cost of debugging capability.

---

### 875. How does cgo interact with Go's runtime?
"It switches stacks.
Go stack (tiny) -> System Stack (large) -> C code.
It ensures the GC doesn't move pointers while C is using them (Pinning).
The overhead is high (~150ns) due to stack switching. I batch C calls."

#### Indepth
**Syscall Interaction**. C code can block forever. If a C function blocks, the Go runtime (Sysmon) sees the P (Processor) is stuck. It detaches the M (Machine thread) from the P and starts a *new* M for the P to keep running other goroutines. This is why CGO apps can spawn thousands of OS threads if you aren't careful.

---

### 876. What are zero-sized types and how are they used?
"`struct{}` takes 0 bytes.
`map[string]struct{}` is a Set.
The compiler optimizes them away.
Pointers to zero-sized variables might all point to the same address (`zerobase`)."

#### Indepth
**Malloc Optimization**. `malloc(0)` in C returns a unique pointer (usually). In Go, `new(struct{})` returns `runtime.zerobase`. This is a global variable. All 0-byte allocations share it. This means `a := struct{}{}; b := struct{}{}; &a == &b` might be true, optimization-permitting.

---

### 877. How does type aliasing differ from type definition?
"**Definition**: `type MyInt int`. New type. `MyInt(1) != int(1)`. Useful for adding methods.
**Alias**: `type MyInt = int`. Same type. Interchangable.
Used for Refactoring (moving a type between packages) without breaking compatibility."

#### Indepth
**Gradual Repair**. Large refactors use aliases. 
Step 1: Move `Type` from `pkgA` to `pkgB`.
Step 2: Add `type Type = pkgB.Type` in `pkgA`.
Step 3: Update consumers to import `pkgB` incrementally.
Step 4: Once all consumers are updated, delete the alias in `pkgA`. This allows zero-downtime refactoring.

---

### 878. How does Go avoid null pointer dereferencing?
"It doesn't completely. `*p` panics if p is nil.
But it avoids uninitialized memory. All variables are zero-valued (`0`, `""`, `nil`).
You never get 'garbage' memory values, only deterministic nils."

#### Indepth
**Option Types**. Go lacks `Option<T>` (Rust) or `Optional<T>` (Java). `*T` is the de-facto Option type. `nil` = None. This conflates "Pointer to data" with "Presence of data". Generics allow creating `type Option[T any] struct { Val T; Present bool }`, which is safer but less ergonomic than built-in pointers.

---

### 879. What’s the role of go/types package?
"It performs Type Checking.
It resolves identifiers (e.g., `fmt.Println`) to Objects.
It computes types of expressions.
Included in the standard library, allowing anyone to build IDE-like tools."

#### Indepth
**Importers**. `go/types` needs to read dependencies. It relies on `Importer`. The default importer reads compiled `.a` files (`pkg` folder). `gopls` uses a custom importer that reads from source, allowing it to inspect code that hasn't been compiled yet. This is why `gopls` works even if your code has build errors.

---

### 880. How does Go manage ABI stability?
"Go has an internal ABI (Register based since 1.17).
It’s not stable between versions for Go code.
However, `cgo` provides a stable ABI to C.
The standard library guarantees API compatibility, but the binary interface changes often."

#### Indepth
**Register ABI**. Before Go 1.17, Go passed arguments on the Stack (slow). Now `amd64` uses registers (RAX, RBX...) for the first ~9 arguments. This reduced CPU overhead by ~5%. This is why Assembly written for Go 1.16 broke in Go 1.17. The compiler auto-generates adapter wrappers for old assembly code.
