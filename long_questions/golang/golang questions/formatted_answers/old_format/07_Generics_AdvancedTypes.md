# 🟡 **121–140: Generics, Type System, and Advanced Types**

### 122. What is type inference in Go?
"Type inference is when the compiler figures out the type of a variable for me.

When I write `x := 42`, Go infers that `x` is an `int`. I don't need to write `var x int = 42`.
It works for complex types too. If a function returns a `map[string]User`, I can just capture it with `m := fn()`. It keeps the code concise without losing the safety of static typing."

#### Indepth
Go's type inference is strictly "local". It only works within a function body. You cannot use `:=` for top-level package variables, nor can you infer function parameter types (like TypeScript or Haskell might). This design decision keeps parsing fast and readable—you always know the signature of a function just by looking at it.

---

### 123. How do you use generics with struct types?
"I define the struct with type parameters in square brackets.

`type Stack[T any] struct { items []T }`.
When I instantiate it, I specify the type: `s := Stack[int]{}`.
Now, usage is type-safe. `s.items` is strictly a slice of `int`. If I try to push a string, the compiler stops me. Before generics, I had to use `interface{}`, which was slow and unsafe."

#### Indepth
Generics in Go 1.18+ are instantiated at compile-time (monomorphization). This means `Stack[int]` and `Stack[string]` are compiled as two completely separate structs in the binary. This avoids the runtime overhead of boxing/unboxing found in Java's Type Erasure, but can slightly increase binary size.

---

### 124. Can you restrict generic types using constraints?
"Yes, I use **interfaces** to restrict what types are allowed.

If I write `[T Stringer]`, where `Stringer` is `interface{ String() string }`, then `T` *must* have a `.String()` method.
I can also use **type sets** in interfaces: `interface{ int | float64 }`. This restricts `T` to be one of those specific types, allowing me to use operators like `+` or `*` inside the generic function."

#### Indepth
An interface behaves as a **type set** only when used as a constraint. You cannot declare a variable of type `interface{ int | float64 }` because the runtime wouldn't know how much memory to allocate (int might be 64-bit, float might be 64-bit, but struct unions would differ). This dual nature of interfaces (Method Sets vs Type Sets) is key to understanding Go generics.

---

### 125. How to create reusable generic containers (e.g., Stack)?
"I implement them as a struct with a type parameter `[T any]`.

`func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }`
`func (s *Stack[T]) Pop() T { ... }`

This is a game changer for data structures. I can write a `Set[T]`, a `Queue[T]`, or a `LinkedList[T]` once, and use it for `int`, `string`, or `User` structs with zero code duplication and zero runtime overhead."

#### Indepth
Be careful not to over-use generics. If you are just copying data around (like a `Buffer` or `Cache`), `[]byte` or `any` might still be simpler. Generics shine when you need to manipulate the *contents* of the data in a type-safe way (like checking `item > max` in a priority queue).

---

### 126. What is the difference between `any` and interface{}?
"There is absolutely **no difference**. `any` is simply an alias for `interface{}`.

It was introduced in Go 1.18 alongside generics because typing `interface{}` repeatedly in generic constraints (`[T interface{}, U interface{}]`) was verbose and ugly.
I use `any` in all new code because it clearly communicates 'this can be anything'."

#### Indepth
Since `any` is just an alias, `go fmt` doesn't rewrite `interface{}` to `any` automatically for backwards compatibility. However, most teams execute a one-time `gofmt -r 'interface{} -> any'` rewrite rule to modernize their codebase.

---

### 127. Can you have multiple constraints in a generic function?
"Yes, I can define multiple type parameters, each with its own constraint.

`func Map[K comparable, V any](m map[K]V)`.
Here, `K` must be `comparable` (so it can be a map key), and `V` can be anything. I separate them with commas. It allows me to express complex relationships between input types."

#### Indepth
You can also reference one type parameter in another's constraint. `func Copy[S ~[]E, E any](s S) S`. Here `S` is constrained to be a slice of `E`. This is useful when you want to return the *exact same type* (including named types like `type MySlice []int`) rather than just `[]int`.

---

### 128. Can interfaces be used in generics?
"Yes, extensive interfaces *are* the constraints.

But standard interfaces (like `fmt.Stringer`) can also be used as type arguments.
`type Printer[T fmt.Stringer] struct { val T }`.
However, I have to be careful: interfaces with **type sets** (like `int | float`) can **only** be used as constraints; they cannot be used as variable types because the compiler doesn't know the memory layout of 'int OR float'."

#### Indepth
This distinction is mostly about **runtime capability**. A variable at runtime must have a single, known layout. A constraint is a **compile-time** rule. Since `int | float` isn't a single layout, it can't be a variable type. Basic interfaces (methods only) describe a "fat pointer" layout (data + itable) so they can be valid variable types.

---

### 129. What is type embedding and how does it differ from inheritance?
"Type embedding is **composition**.

If I embed `User` inside `Admin`: `type Admin struct { User }`.
`Admin` gets all of `User`'s methods promoted to it.
However, `Admin` is **not** a `User`. I cannot pass `Admin` to a function expecting `User`. This avoids the 'is-a' relationship trap of inheritance and encourages explicit interfaces."

#### Indepth
This is often called **Promoted Methods**. It applies to fields too. `admin.ID` works even if `ID` belongs to the embedded `User`. However, serialization (JSON) treats them as nested unless you use the `inline` tag in some libraries, though standard `encoding/json` flattens embedded structs automatically.

---

### 130. How does Go perform type conversion vs. type assertion?
"**Conversion** transforms data: `float64(myInt)`. It changes the bits from integer format to float format.
**Assertion** checks types: `val.(int)`. It unwraps an interface to reveal the concrete data inside.

Conversion works for compatible types (int to float). Assertion works only on interfaces. If I try to assert `val.(int)` and `val` actually holds a string, it panics (unless I use the comma-ok idiom)."

#### Indepth
Type assertions are performed at runtime by checking the `itable` (interface table). It’s a very fast pointer comparison. Conversion (`T(x)`) is a compile-time (or runtime) operation that actually changes the underlying bits (e.g., float to int truncates the decimal). They are fundamentally different operations.

---

### 131. What are tagged unions and how can you simulate them in Go?
"A tagged union (or sum type) holds one of several fixed types. Go doesn't have them natively.

I simulate them using an interface with a **sealed method**.
`type Result interface { isResult() }`
`type Success struct { Val int }; func (Success) isResult() {}`
`type Failure struct { Err error }; func (Failure) isResult() {}`
Then I use a type switch to handle each case exhaustively. It’s verbose but safe."

#### Indepth
This pattern is heavily used in the standard library (e.g., `ast.Node` implementations). With generics, some libraries introduce `Either[L, R]` types, but the interface-based "sum type" remains the most idiomatic way to handle "one of N types" scenarios in Go.

---

### 132. What is the use of `iota` in Go?
"`iota` is a built-in counter for `const` blocks.

`const ( A = iota; B; C )` assigns 0, 1, 2.
I use it for **enums**. It auto-increments.
It’s also powerful for bitmasks: `1 << iota` gives me 1, 2, 4, 8. It saves me from manually typing out increasing values and avoiding typos."

#### Indepth
`iota` resets to 0 whenever the keyword `const` appears. It's scoped to the block. If you have two const blocks, `iota` restarts. A common trick is to use `_ = iota` to skip the zero value if '0' is not a valid enum state (like `Unknown`).

---

### 133. How are custom types different from type aliases?
"A **custom type** (`type MyInt int`) creates a distinct type.
It loses all methods of the underlying type. I cannot assign `int` to `MyInt` without a cast. I use this to attach methods to primitives (e.g., `func (m MyInt) IsPositive() bool`).

A **type alias** (`type MyInt = int`) is just a rename. It’s identical to `int`. I only use aliases for refactoring (moving a type between packages)."

#### Indepth
Type aliases were introduced primarily to help with **gradual code repair**. If you move `oldpkg.Thing` to `newpkg.Thing`, you can leave `type Thing = newpkg.Thing` in `oldpkg` so that existing clients don't break. It's a "soft link" for types.

---

### 134. What are type sets in Go 1.18+?
"Type sets simplify interface definitions by allowing me to list concrete types.

`type Number interface { int | float64 | float32 }`.
This interface matches any of those types. This is the foundation of **generics constraints**. It allows me to write a `Min[T Number](a, b T)` function that works on all numeric types without casting."

#### Indepth
The `~` tilde token is often used in type sets: `~int`. This means "any type whose *underlying* type is int". So `type MyInt int` satisfies `~int`, but does not satisfy `int`. You almost always want `~T` in library code to support custom types.

---

### 135. Can generic types implement interfaces?
"Yes. A generic struct `Stack[T]` can implement `fmt.Stringer`.

`func (s Stack[T]) String() string { ... }`.
When I instantiate `Stack[int]`, it gains the `String()` method. This allows me to pass `Stack[int]` to `fmt.Println` just like any non-generic type. It bridges the gap between the new generic world and the old interface-based world."

#### Indepth
This works because the compiler knows that for any valid `T`, the resulting struct `Stack[T]` will have the method. If `T` was used in the method *signature* (e.g. `Pop() T`), then `Stack[T]` generally cannot satisfy a non-generic interface unless that interface method matches exactly (which usually only happens if `T` is fixed or `any` and the interface uses `any`).

---

### 136. How do you handle constraints with operations like +, -, *?
"Go strictly forbids operator overloading.

To use `+` in a generic function, I must constrain `T` to types that support it.
I use the `golang.org/x/exp/constraints` package.
`func Add[T constraints.Ordered](a, b T) T`.
The `Ordered` constraint includes all integers, floats, and strings, guaranteeing to the compiler that `a + b` is a valid operation."

#### Indepth
Go does not support **Operator Overloading** for custom methods. You cannot define `+` for your `Matrix` struct. You must write `m.Add(m2)`. This design keeps code simple—when you see `+`, you know it's a cheap CPU instruction, not an arbitrary function call that might take 5 seconds.

---

### 137. What is structural typing?
"Structural typing (Duck Typing) means compatibility is determined by structure, not name.

If interface `I` requires method `Foo()`, and struct `S` has method `Foo()`, then `S` implements `I`.
I don't need to write `implements I`. This decoupling allows me to define interfaces in the **consuming** package, rather than the producer package, which is a key architectural advantage in Go."

#### Indepth
This reverses the dependency graph compared to Java. In Java, `File` implements `Reader`. In Go, the *consumer* defines `Reader`, and `File` implicitly satisfies it. This means I can define a `MyReader` interface for a library I don't own, without needing that library to change its code.

---

### 138. Explain the difference between concrete and abstract types.
"**Concrete types** (like `int`, `*os.File`) describe a value's exact memory layout and implementation.
**Abstract types** (like `io.Reader`, `fmt.Stringer`) describe behavior (methods).

I follow the rule: **Accept interfaces, return structs.**
I accept abstract types to make my functions flexible, but I return concrete types to let the caller use the full power of the object."

#### Indepth
Returning structs (concrete types) allows you to add new methods to the implementation later without breaking consumers. If you return an interface, adding a method to that interface breaks all existing implementations (because they now fail to satisfy the new signature).

---

### 139. What are phantom types and are they used in Go?
"A phantom type is a generic type where the type parameter isn't used in the struct fields.

`type ID[T any] string`.
The `T` doesn't exist in memory; it’s just a label.
I use it for **Type Safety**. `ID[User]` and `ID[Product]` are both strings at runtime, but the compiler treats them as different types. It prevents me from accidentally passing a ProductID to a function that deletes a User."

#### Indepth
Phantom types are a zero-cost abstraction. Since `ID[User]` is just a string at runtime, there is no extra memory usage. It’s purely a compile-time "label" that forces correctness. It's extremely useful for preventing "Primitive Obsession" bugs where everything is just `int` or `string`.

---

### 140. How would you implement an enum pattern in Go?
"Go doesn't have `enum` keywords.
I use a custom integer type and a `const` block with `iota`.

`type State int`
`const ( Pending State = iota; Active; Closed )`

To make it usable, I implement the `String()` method to print "Active" instead of "1". For stricter validation, I might add a `Validate()` method to ensure the integer is within the valid range."

#### Indepth
Since Go enums are just integers, nothing stops a user from passing `State(999)`. The compiler won't catch it. Heavily used enums should always have a `IsValid() bool` method or use a linter like `exhaustive` to check switch statements coverage.

---

### 141. How can you implement optional values in Go idiomatically?
"The classic way is using **pointers**: `*int`. If `nil`, it's missing.

With generics, some people use `Option[T]`, but I find it un-idiomatic.
For function arguments, I use **Functional Options** (`WithTimeout(5s)`).
For maps, I use the `val, ok := m[key]` idiom.
I stick to pointers for JSON structs (`json:"age,omitempty"`) because the standard library supports it seamlessly."

#### Indepth
Use pointers (`*int`) sparingly for optionality. `nil` *int means "missing", but it also means "GC pressure" (variables on heap). For high-performance code, use a struct `type NullInt struct { Val int; Valid bool }` (similar to `sql.NullInt64`) to keep data on the stack.
