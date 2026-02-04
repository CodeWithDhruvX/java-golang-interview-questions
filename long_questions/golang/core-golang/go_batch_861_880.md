## 🧠 Go Compiler & Language Theory (Questions 861-880)

### Question 861: How do you build a custom Go compiler plugin?

**Answer:**
Go doesn't support "Compiler Plugins" like GCC/LLVM easily.
You can write a **Linter** (using `go/analysis`) or modify the source of the Go compiler (`cmd/compile`) directly.

---

### Question 862: What is SSA (Static Single Assignment) form in Go?

**Answer:**
Intermediate Representation (IR) used by the Go compiler backend.
Every variable is assigned exactly once.
Enables optimizations like Dead Code Elimination, Bounds Check Elimination, and Register Allocation.

---

### Question 863: How does Go handle type inference?

**Answer:**
Only inside functions (`:=`).
The compiler looks at the Right-Hand Side (RHS) type and assigns it to the Left.
For Generics, it uses **Unification** to deduce type parameters.

---

### Question 864: What is escape analysis in Go?

**Answer:**
(See Q526). Determines Stack vs Heap allocation.

---

### Question 865: How does inlining affect performance in Go?

**Answer:**
**Inlining:** Replacing function call with function body.
Pros: Removes call overhead, enables further optimization (DCE).
Cons: Increases binary size.
Check: `go build -gcflags "-m"`.

---

### Question 866: What are build constraints and how do they work?

**Answer:**
(See Q786). `//go:build tag`.

---

### Question 867: How does `defer` work at the bytecode level?

**Answer:**
Historically: Expensive (malloc).
Now (Go 1.14+): **Open-coded defer**.
The compiler injects the deferred code directly at every return point. Zero overhead for common cases.

---

### Question 868: What is the Go frontend written in?

**Answer:**
Go (Self-hosting since Go 1.5).
Parser, Type Checker, AST generator are all in `src/cmd/compile`.

---

### Question 869: How are interfaces implemented in memory?

**Answer:**
A tuple: `(type, data)`.
`type` points to the `itab` (Interface Table, containing method pointers).
`data` points to the concrete value.

---

### Question 870: What are method sets and how do they affect interfaces?

**Answer:**
- `T` has methods with receiver `T`.
- `*T` has methods with receiver `T` OR `*T`.
If interface requires `Write()`, and `Write` is defined on `*T`, you MUST pass `&T`. Passing `T` won't satisfy the interface.

---

### Question 871: How do you implement AST manipulation in Go?

**Answer:**
Use `go/parser`, `go/ast`, `go/printer`.
Parse source -> Edit AST nodes (rename var) -> Print back to source.
Used by `gofmt` and `rename` tools.

---

### Question 872: What is the Go toolchain pipeline from source to binary?

**Answer:**
1.  **Parse:** Source -> AST.
2.  **Type Check:** Validate types.
3.  **SSA Generation:** AST -> SSA.
4.  **Optimization:** SSA -> Optimized SSA.
5.  **Code Gen:** SSA -> Assembly (`.o`).
6.  **Link:** `.o` files -> Executable (ELF/PE/Mach-O).

---

### Question 873: How are function closures handled by the Go compiler?

**Answer:**
If the closure captures variables from the outer scope, those variables **escape to the heap**.
The closure becomes a struct `funcval { code_ptr, captured_var_ptr }`.

---

### Question 874: What is link-time optimization in Go?

**Answer:**
Go's linker removes **Dead Code** (DCE).
If function `Main` calls `A` but never `B`, `B` is not included in the binary (even if imported).

---

### Question 875: How does cgo interact with Go's runtime?

**Answer:**
CGO calls require switching stack (Go stack -> C stack/Thread stack).
This involves a heavy context switch (saving registers).
It also limits the scheduler (the M is locked to the C call).

---

### Question 876: What are zero-sized types and how are they used?

**Answer:**
`struct{}`. Size = 0.
Used for:
- Sets: `map[string]struct{}`.
- Signals: `chan struct{}`.
- Methods: Methods on `struct{}` to group logic without state.

---

### Question 877: How does type aliasing differ from type definition?

**Answer:**
- **Definition:** `type MyInt int`. New type. Cannot mix with `int`.
- **Alias:** `type MyInt = int`. Same type. Can mix. Used for refactoring (moving types between packages).

---

### Question 878: How does Go avoid null pointer dereferencing?

**Answer:**
It doesn't completely. `nil` exists.
However, it avoids *implicit* conversion.
You can't assign `nil` to `int`.
Some libraries using `Option[T]` types (Generics) try to enforce safety.

---

### Question 879: What’s the role of `go/types` package?

**Answer:**
The Type Checker.
It resolves identifiers (which validation function is this?), checks compatibility, and computes the type of every expression.

---

### Question 880: How does Go manage ABI stability?

**Answer:**
Go does **not** have a stable function ABI (Application Binary Interface) for plugins until recently (Go 1.17 register ABI).
Everything is recompiled from source (static linking).
This avoids "DLL Hell".

---
