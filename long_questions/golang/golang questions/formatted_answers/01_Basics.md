# Basic Level Golang Interview Questions

## From 01 Basics

# 🟢 **1–20: Basics**

### 1. What is a function literal (anonymous function)?
"A **function literal** is a function defined without a name. It’s a value that I can assign to a variable or pass as an argument to another function.

Technically, these are **closures**. They capture variables from their surrounding scope by reference. If the outer function returns, the variables used by the closure survive on the heap as long as the closure exists.

I use them constantly for **middleware** in web servers or for `defer` logic. For example, `defer func() { file.Close() }()` is a common pattern to ensure cleanup logic runs with the latest context."

#### Indepth
Go's specific implementation of closures relies on **escape analysis**. The compiler determines if variables captured by the closure need to survive the function call. If they do, they are moved from the stack to the heap. This ensures safety but incurs a small GC pressure.

---

### 2. How does the `net/http` package work?
"The `net/http` package is Go's battle-tested standard library for building HTTP clients and servers.

Its core abstraction is the **Handler interface**, which has a single method: `ServeHTTP`. When a request comes in, the server spawns a new **goroutine** for it and calls this method. It handles keep-alives, timeouts, and protocol parsing (HTTP/1.1 and HTTP/2) automatically.

I love it because it’s production-ready out of the box. Unlike other languages where I need a heavy framework like Express or Spring just to say 'Hello World', Go’s standard library is powerful enough for 90% of my use cases."

#### Indepth
The `DefaultServeMux` used by `http.Handle` is a global variable. In production, avoid using it because any imported package can register a handler and expose endpoints you didn't intend. Always create your own `http.ServeMux`. Also, the default `http.Server` has no timeouts, which makes it vulnerable to Slowloris attacks.

---

### 3. What is Go and who developed it?
"Go is an open-source, statically typed language developed at **Google** by Robert Griesemer, Rob Pike, and Ken Thompson (the creators of Unix and C).

They designed it to solve Google-scale problems: slow compilation times, unreadable codebases, and the difficulty of writing concurrent software. It combines the performance of **C++** with the readability of **Python**.

For me, it’s the 'boring' language in the best way possible. It doesn't have a million features, which means any developer can join a project and understand the code in a day."

#### Indepth
Go maintains the Plan 9 OS philosophy of "everything is a file" and UTF-8 handling (Ken Thompson invented UTF-8). It also strictly enforces **SemVer** and backward compatibility; Go 1 code written in 2012 still compiles today.

---

### 4. What are the key features of Go?
"Go focuses on simplicity and efficiency. Its headline features are **Goroutines** for concurrency, **Channels** for communication, and a **fast Garbage Collector**.

It is statically typed but feels dynamic because of **type inference**. It compiles to a single static binary, making deployment trivial (no 'DLL hell').

I particularly appreciate the tooling. `gofmt`, `go test`, and `go mod` are standard tools that everyone uses, eliminating the 'configuration wars' I see in the JavaScript or Java ecosystems."

#### Indepth
Go's concurrency model is based on **CSP (Communicating Sequential Processes)**, a formal language described by Tony Hoare. Unlike thread-based concurrency where memory is shared, Go encourages "Sharing memory by communicating".

---

### 5. How do you declare a variable in Go?
"I have two main ways: the `var` keyword and the **short declaration** operator `:=`.

`var x int` declares a variable with its zero value. I use this at the package level or when I want to be explicit about the type. `x := 10` declares and initializes it in one step, letting the compiler infer the type.

I use `:=` for 99% of local variables because it keeps the code concise. However, if I need to declare a variable but initialize it later (like inside an `if` block), I stick to `var`."

#### Indepth
Be careful with **variable shadowing**. A `:=` declaration inside an `if` block creates a *new* variable that masks the outer one. This is a common source of bugs `if x, err := func(); ...` where the outer `x` is not updated.

---

### 6. What are the data types in Go?
"Go uses a small, standard set of types.

We have **Basic types** (bool, string, int, float64), **Aggregate types** (arrays, structs), and **Reference types** (slices, maps, channels, pointers, functions).

There is no 'Class' type. We build complex data structures using **structs**. I find this refreshing because it forces me to separate data (structs) from behavior (methods), avoiding the deep inheritance hierarchies of OOP."

#### Indepth
Go structs have flattened memory layouts. Fields are laid out sequentially in memory, with padding added for alignment. This contrasts with Java or Python where objects are references scattered in the heap. This data locality is a major factor in Go's performance.

---

### 7. What is the zero value in Go?
"The **zero value** is the default value assigned to variables that are declared but not initialized. Go never leaves memory uninitialized.

For `int` it is `0`, `bool` is `false`, `string` is `""`. For reference types like pointers, slices, and maps, it is `nil`.

I rely on this heavily. For example, a `sync.Mutex` is ready to use immediately because its zero value is an 'unlocked mutex'. This 'make the zero value useful' philosophy is a key part of idiomatic Go API design."

#### Indepth
While standard types are safe, be careful with `nil` Maps. You can read from a `nil` map (it returns zero values), but writing to a `nil` map causes a panic. Slices are safer; you can append to a `nil` slice perfectly fine.

---

### 8. How do you define a constant in Go?
"I use the `const` keyword.

Constants are evaluated at **compile time**. They can be typed (`const X int = 10`) or untyped (`const X = 10`). Untyped constants are powerful because they have arbitrary precision and can be used in any numeric context without casting.

I often use the `iota` identifier inside a `const` block to create enums. `const ( A = iota; B )` assigns 0 to A and 1 to B automatically."

#### Indepth
Go constants are represented with at least 256 bits of precision during compilation. This allows you to perform high-precision math with constants (like `math.Pi`) and only round to a float/int at the final assignment step.

---

### 9. Explain the difference between `var`, `:=`, and `const`.
"`var` declares a variable. It works at both package and function levels.
`:=` is short variable declaration. It *only* works inside functions and infers the type.
`const` declares a value that cannot change and is computed at compile time.

I use `const` for magic numbers and config strings. I use `:=` for almost all local logic. I reserve `var` for package-level globals (which I try to avoid) or zero-value initialization."

#### Indepth
You can use `:=` to redeclare a variable *if* it appears in a multi-variable declaration where at least one variable is new. `err := fn1(); res, err := fn2()` is valid because `res` is new, and `err` is just assigned to.

---

### 10. What is the purpose of `init()` function in Go?
"`init()` is a special function that runs automatically before `main()` starts.

It is used to set up package-level state, like parsing configuration flags or registering database drivers (like `postgres`). A package can have multiple `init` functions, and they run in the order they appear.

I use it sparingly because strict ordering dependencies can be hard to debug. If possible, I prefer explicit `New()` or `Setup()` functions so the caller has control over *when* initialization happens."

#### Indepth
Package initialization order is determined by dependency graph. Imports are initialized first. Within a package, variables are initialized before `init()` functions. This deterministic order is crucial for system reliability.

---

### 11. How do you write a for loop in Go?
"Go only has one looping keyword: `for`.

It can be a C-style loop: `for i:=0; i<10; i++`.
It can be a while loop: `for x < 100`.
It can be an infinite loop: `for { }`.

My favorite is the **range** clause: `for k, v := range items`. It iterates over slices, maps, strings, and channels. It handles the index/key and value extraction cleanly, avoiding off-by-one errors."

#### Indepth
Prior to Go 1.22, the loop variable was reused across iterations, referencing the same memory address. This caused bugs when capturing the variable in a goroutine. In modern Go (1.22+), loop variables have per-iteration scope, fixing this "most common Go mistake".

---

### 12. What is the difference between `break`, `continue`, and `goto`?
"`break` exits the innermost loop or switch statement immediately.
`continue` skips the rest of the current iteration and jumps to the next one.
`goto` jumps execution to a labelled line.

I use `break` and `continue` constantly. I almost *never* use `goto`, except perhaps to break out of a deeply nested loop (`break Label` does this better) or in generated code."

#### Indepth
Labels are block-scoped and must be defined in the same function. They are most useful when you have a `switch` inside a `for` loop and want to break the *loop*, not the switch. calling `break` alone would only exit the switch.

---

### 13. What is a `defer` statement?
"`defer` schedules a function call to run immediately *after* the surrounding function returns.

The arguments are evaluated *when the defer is declared*, but the function execution happens at the end. It follows a LIFO (Last-In-First-Out) order.

This is my favorite error-handling feature. It ensures that `file.Close()` happens whether the function returns normally or panics. It puts the cleanup logic right next to the allocation logic, making leaks much less likely."

#### Indepth
Deferred calls are pushed onto a stack. When the function exits, they are popped and executed. The overhead of `defer` used to be significant (~50ns), but recent Go versions (1.14+) use "open-coded defers" which compile them inline, making them nearly zero-cost.

---

### 14. How does `defer` work with return values?
"Deferred functions run **after** the `return` statement has set the return values, but **before** the function actually exits.

This means a deferred function can inspect and even **modify** named return values.

I use this pattern to handle panics in a unified way. `defer func() { if r := recover(); r != nil { err = fmt.Errorf("panic: %v", r) } }()` allows me to turn a crash into a structured error return."

#### Indepth
This pattern is the only way to modify a return value *after* the `return` statement has executed. It is widely used for **error wrapping**, where you want to add context to an error returned by any point in the function.

---

### 15. What are named return values?
"Go allows me to give names to return parameters in the function signature: `func div(a, b int) (res int, err error)`.

These variables are initialized to their zero values. I can return them simply by typing `return` (called a 'naked return').

I use them for documentation clarity. `func (lat, long float64)` is clearer than `func (float64, float64)`. However, I avoid naked returns in long functions because they hurt readability—I prefer explicit `return res, err`."

#### Indepth
Behind the scenes, named return values are just local variables declared at the top of the function. **Naked returns** should be used sparingly; if the function is longer than a screen, explicit returns improve readability significantly.

---

### 16. What are variadic functions?
"A variadic function accepts a variable number of arguments. `fmt.Println` is the most famous example.

I define it using `...Type` as the last parameter, like `func sum(nums ...int)`. Inside the function, `nums` is just a slice `[]int`.

I use this often for 'Option' patterns in constructors, like `NewServer(opts ...Option)`. It allows the caller to provide zero, one, or ten configuration options without changing the function signature."

#### Indepth
When you pass a slice to a variadic function like `fn(slice...)`, Go does *not* create a new slice. It passes the existing slice header. This means modifying elements in the variadic function will modify the original slice's backing array.

---

### 17. What is a type alias?
"A **type alias** (`type MyInt = int`) creates an alternative name for the *same* type. They are interchangeable.

This is distinct from a **type definition** (`type MyInt int`), which creates a *new*, distinct type that requires casting.

I rarely use aliases restrictedly. Their primary use case is during **refactoring**: moving a type from Package A to Package B, but leaving an alias in Package A so I don't break existing code that imports it."

#### Indepth
Type aliases were introduced efficiently in Go 1.9 to aid strictly in **gradual code repair**. They allow large codebases to move a type between packages (e.g., `context.Context` moving from `golang.org/x/net/context` to `std`) without breaking consumers.

---

### 18. What is the difference between `new()` and `make()`?
"`new(T)` allocates zeroed memory for a type and returns a **pointer** (`*T`). It works for any type (int, struct, etc.).

`make(T)` is strictly for **slices, maps, and channels**. It initializes the internal data structure (like the hash bucket headers) and returns the **value** (`T`), not a pointer.

If I tried `new(map[string]int)`, I'd get a pointer to a nil map. Writing to it would panic. So I always use `make()` for reference types and `new()` (or `&Struct{}`) for value types."

#### Indepth
The difference exists because slices, maps, and channels are complex structures that must be initialized before use (e.g., creating the hash map buckets). `new()` only zeroes memory, which isn't enough for these types. `make()` performs that necessary setup.

---

### 19. How do you handle errors in Go?
"In Go, errors are **values**, usually returned as the last value of a function. There are no exceptions.

I check `if err != nil` immediately after a call.

This explicit handling forces me to think about failure modes at every step. I can't just 'hope' a higher-level catch block handles it. It makes Go code verbose but incredibly robust and readable—I can see exactly where the control flow goes."

#### Indepth
Go 1.13 introduced error wrapping. Use `errors.Is(err, target)` instead of `err == target` to check if an error wraps a specific sentinel value. Use `errors.As(err, &target)` to check if it wraps a specific type.

---

### 20. What is panic and recover in Go?
"**Panic** ceases normal execution. It runs any deferred functions and then crashes the program. It should only be used for unrecoverable errors (like a corrupted config at startup).

**Recover** is a built-in function that regains control of a panicking goroutine. It *only* works inside a `defer` function.

I treat `panic/recover` like an assertion mechanism, not for control flow. The only place I use `recover` is in the top-level HTTP middleware to ensure a single bad request implies a 500 error rather than crashing the entire server."

#### Indepth
Panic unwinds the stack, running deferred functions. If not recovered, it crashes the process with a non-zero exit code. `recover()` returns `nil` if called normally (not during a panic), so it effectively does nothing outside a defer.

---

### 21. What are blank identifiers in Go?
"The **blank identifier** (`_`) is a write-only variable placeholder.

Go refuses to compile unused variables. If a function returns `(val, err)` and I only care about the error, I must write `_, err := fn()`.

I also use it for **import side-effects**: `import _ "image/png"` registers the PNG decoder without exposing the package name. It’s a clean way to satisfy the compiler while ignoring data I don't need."

#### Indepth
Identified by `_`, this is a language mechanism to discard values. It is effectively a "black hole" for data. Since Go requires every declared variable to be used, `_` is necessary for ignoring return values you don't care about.


## From 02 Arrays Slices Maps

# 🟡 **21–40: Arrays, Slices, and Maps**

### 22. What is the difference between an array and a slice?
"An **array** is a fixed-length sequence of items. When I define `[5]int`, its size is part of its type. It can never grow, and if I pass it to a function, Go copies the entire block of memory (value type).

A **slice**, on the other hand, is a lightweight, dynamic view over an array. It can grow and shrink.

I use slices 99% of the time because they are flexible. Under the hood, a slice is just a tiny struct with a pointer to the array, a length, and a capacity. Passing a slice to a function is cheap because I'm just copying that tiny header, not the data itself."

#### Indepth
The slice header is defined in `reflect.SliceHeader` (deprecated) or `unsafe.Slice` in newer Go versions. It contains `Data uintptr`, `Len int`, and `Cap int`. Because it's a struct passed by value, the `Data` pointer is copied, so the function can modify the underlying array elements, but cannot change the caller's slice length/capacity (unless returned).

---

### 23. How do you append to a slice?
"I use the built-in `append()` function. It takes the slice and the new elements, and returns a **new slice**.

It’s crucial to prevent bugs: `append` might return a pointer to a *different* underlying array if the original one ran out of capacity.

Because of this, I strictly follow the pattern `s = append(s, val)`. If I ignore the return value, I’m referring to the old slice header which doesn’t know about the new elements, leading to data loss."

#### Indepth
`append` uses a sophisticated growth strategy. For small slices (<1024 elements), it doubles capacity. For larger slices, it grows by ~1.25x to avoid wasting memory. It also aligns memory blocks to system page sizes for allocator efficiency.

---

### 24. What happens when a slice is appended beyond its capacity?
"Go performs a **reallocation**. It allocates a new, larger array (usually double the size for smaller slices) in memory.

Then, it copies all the existing elements from the old array to the new one, and updates the slice's pointer to refer to this new block.

This 'grow and copy' operation is why `append` is usually fast (O(1)) but occasionally slow (O(N)). If I know the size beforehand, I always use `make([]int, 0, capacity)` to pre-allocate memory and avoid these expensive resize operations."

#### Indepth
When re-slicing a large array (e.g., `largeArr[:10]`), the new slice keeps the *entire* backing array in memory, preventing GC. This is a common **memory leak**. To fix it, copy the small slice to a new, minimal slice via `copy()` so the large array can be garbage collected.

---

### 25. How do you copy slices?
"I use the built-in `copy(dest, src)` function.

The most important thing to remember is that `copy` only moves the minimum number of elements common to both slices. It doesn't allocate memory for me.

A common mistake I’ve made is trying to copy into an empty nil slice. `copy(nil, src)` does nothing. I must explicitly `make` the destination slice with `len(src)` before calling copy."

#### Indepth
`copy` is a built-in for valid slices, but it relies on `memmove` under the hood. It handles overlapping slices correctly (e.g., `copy(s[1:], s[0:])` to shift elements). It is faster than a `for` loop because the compiler optimizes it to block memory operations.

---

### 26. What is the difference between len() and cap()?
"**len()** is the length—the number of elements I can validly access right now (indices 0 to len-1).

**cap()** is the capacity—the total size of the underlying backing array. It tells me how many elements I can append before Go needs to allocate a new array.

I check `cap()` when optimizing loops. If `len` is 5 but `cap` is 1000, I might want to re-slice or copy the data to a smaller slice to allow the garbage collector to free that giant backing array."

#### Indepth
You can modify the capacity of a slice by reslicing *up to* the capacity: `s = s[:cap(s)]`. This "recovers" hidden elements. However, you can never extend a slice beyond its capacity; doing so causes a runtime panic.

---

### 27. How do you create a multi-dimensional slice?
"Go doesn't have C-style multi-dimensional arrays (contiguous blocks). Instead, we have **slices of slices** (`[][]int`).

I have to initialize them in two steps: first `make` the outer slice, then loop through it to `make` each inner slice.

This structure allows 'jagged' arrays where each row has a different valid length. However, it’s not memory-contiguous, which can cause cache misses in high-performance numerical code. For matrix math, I usually flatten it into a single `[]float64` and calculate indices manually."

#### Indepth
With Go 1.21+, we have the `slices` package. `slices.Clone` or `slices.Concat` simplifies many of these operations. `Multi-dimensional` slices add pointer indirection overhead. For high-performance numerical computing, a flat 1D slice with stride arithmetic (`index = y * width + x`) is significantly faster and cache-friendly.

---

### 28. How are slices passed to functions (by value or reference)?
"Everything in Go is passed by **value**.

However, a slice variable is just a **header** (Pointer, Length, Capacity). When I pass a slice, I am copying this header solely.

The *pointer* inside the copy still points to the same underlying array. So, if the function modifies an index (`s[0] = 1`), the caller sees the change. But if the function calls `append` and triggers a resize, the caller *won't* see the new array unless I return the modified slice."

#### Indepth
This behavior highlights why `append` returns a value. If `append` caused a reallocation, the new slice points to a *different* memory address. The caller's slice header would still point to the old (now stale) array if you didn't return and reassign the new slice header.

---

### 29. What are maps in Go?
"A **map** is Go's built-in hash table implementation. It provides unordered key-value pairs with O(1) average lookup time.

I define it like `map[string]User`. Like slices, they behave like reference types.

One trap is that a zero-value map is `nil`. I can read from it (getting zeroes), but writing to it causes a **panic**. I always initialize maps using `make(map[string]int)` or a literal `{}` before writing."

#### Indepth
Go maps are implemented as **hash maps** with buckets. Each bucket holds up to 8 key/value pairs. When a bucket overflows, it chains to an overflow bucket. This structure keeps the map compact in memory and cache-friendly compared to linked-list chaining.

---

### 30. How do you check if a key exists in a map?
"I use the **comma-ok idiom**.

`value, ok := myMap["key"]`.
If `ok` is true, the key exists.
If `ok` is false, the key is missing, and `value` is the zero-value for that type.

This distinction is vital. If I just wrote `val := myMap["key"]` and got `0`, I wouldn't know if the user's balance is actually 0 or if the user doesn't exist."

#### Indepth
Accessing a map is not an atomic operation. Concurrent read/write to a map without synchronization causes a fatal runtime error: `concurrent map iteration and map write`. Unlike standard panics, this one *cannot* be recovered from. It forces the program to crash to prevent data corruption.

---

### 31. Can maps be compared directly?
"No, the `==` operator is not defined for maps (except comparison to `nil`).

If I try `mapA == mapB`, the compiler stops me.

To check equality, I must loop through both maps and compare every key and value manually. In tests, I use `reflect.DeepEqual` or `cmp.Diff`, but in production code, I avoid this because it's slow (O(N))."

#### Indepth
Use `maps.Equal` (from `golang.org/x/exp/maps` or standard `maps` in Go 1.21+) for equality checks. It handles `NaN` values correctly (where `NaN != NaN`). Beware that `reflect.DeepEqual` is recursive and slow, using it in a hot loop is a performance killer.

---

### 32. What happens if you delete a key from a map that doesn’t exist?
"Nothing. It is a **no-op**.

`delete(myMap, "missing_key")` does not panic or return an error.

I appreciate this design choice because it simplifies code—I don't need to wrap every delete in an `if _, ok := m[k]; ok` check. I just command 'delete it', and Go ensures it's gone."

#### Indepth
While `delete` removes the key, it typically does *not* shrink the allocated memory of the map. If you fill a map with 1 million items and delete them all, the map will still consume a large amount of RAM (buckets remain). To reclaim memory, you must recreate the map.

---

### 33. Can slices be used as map keys?
"No, because slices are not **comparable** (they don't support `==`).

A map key must be a type that is valid for equality checks (like ints, strings, pointers, structs of simple types).

If I need a composite key (like a coordinate `[x, y]`), I use an **array** `[2]int` (which *is* comparable) or a struct `struct{X, Y int}` as the key instead."

#### Indepth
The specific requirement for a map key is that the type implementation must define `equality`. Structs are comparable if all their fields are comparable. Formatting a slice to a string key (e.g., `fmt.Sprint(slice)`) is a common workaround but is slow and collision-prone.

---

### 34. How do you iterate over a map?
"I use the `for range` loop: `for k, v := range m`.

The critical thing to remember is that **iteration order is random**. It changes every time I run the program. This is intentional to prevent developers from relying on hash ordering.

If I need deterministic output (like in a JSON API or a test), I collect the keys into a slice, sort them, and then iterate over the map using the sorted keys."

#### Indepth
The randomization of map iteration is achieved by the runtime picking a random "start bucket" offset. This avoids hash flooding attacks (DoS) where an attacker could predict iteration order to slow down the server. It effectively forces developers to not rely on implementation details.

---

### 35. How do you sort a map by key or value?
"Maps themselves cannot be sorted.

I have to extract the data. Usually, I pull all the keys into a slice: `keys := make([]string, 0, len(m))`. Then I sort the slice: `sort.Strings(keys)`.

Finally, I iterate the sorted slice and lookup values: `m[key]`. It’s verbose, but it’s the standard pattern in Go."

#### Indepth
With Go 1.21, the `slices.Sort` function makes this easier. But for huge maps, extracting keys and sorting them is expensive (O(N) allocations + O(N log N) sort). If sorted order is critical, consider using a B-Tree or Skip List implementation instead of a standard map.

---

### 36. What are struct types in Go?
"A **struct** is a typed collection of fields. It is the foundation of data modeling in Go, similar to a Class in Java but without the inheritance baggage.

I define it like `type User struct { Name string; Age int }`.

It purely holds data. I can define methods *on* it, but the struct definition itself is just the memory layout. This separation of state (struct) and behavior (methods) is core to Go's design."

#### Indepth
Empty structs `struct{}` consume zero bytes of storage. `struct{}{}` is effectively free. This is heavily used in `map[string]struct{}` to simulate a **Set** data structure, where we only care about keys and want 0 memory overhead for values.

---

### 37. How do you define and use struct tags?
"Struct tags are string annotations like `` `json:"id"` `` appearing after a field.

By themselves, they do nothing. They are accessed via **Reflection**.

Libraries like `encoding/json` or database ORMs read these tags at runtime to know how to map a field (e.g., 'serialize `UserID` as `user_id`'). It’s a powerful way to add declarative metadata to my static types."

#### Indepth
Tags are often conventionally space-separated key-value pairs: `key:"value" key2:"value2"`. If you need multiple options for a key (like JSON), use commas: `json:"name,omitempty"`. The `reflect` package parses these strings. It's not magic; it's just a string, so typoes (like `jso:"id"`) are silently ignored!

---

### 38. How to embed one struct into another?
"I use **anonymous embedding**. I declare `type Admin struct { User; Level int }`.

This isn't inheritance. `Admin` *has a* `User`. However, Go promotes the fields of `User` so I can access `admin.Name` directly instead of `admin.User.Name`.

I use this for composition—building complex objects from small, reusable pieces. But I’m careful, because `Admin` is *not* a `User` type (polymorphism rules differ from OOP)."

#### Indepth
Wait, embedding also promotes **methods**. If `User` has `Login()`, `Admin` automatically has `Login()`. This allows `Admin` to satisfy interfaces that `User` satisfies. However, the method receiver inside `Login` is still `User`, not `Admin`. It doesn't know it's embedded.

---

### 39. How do you compare two structs?
"If the struct contains only **comparable fields** (ints, strings, arrays), I can use `==`.

If it contains **slices**, **maps**, or **functions**, it is not comparable, and `==` will cause a compile error.

In those cases, I have to write a custom `.Equal()` method or use `reflect.DeepEqual`. I always prefer the custom method for performance-critical hot paths."

#### Indepth
Struct comparison stops at the first mismatch. It does a memory comparison for simple types. If a struct contains a `func` field, it makes the entire struct non-comparable. You can make a struct non-comparable *intentionally* by adding a `_ [0]func()` field, forcing users to use your API for equality.

---

### 40. What is the difference between shallow and deep copy in structs?
"A standard assignment `b := a` performs a **shallow copy**.

Go copies the struct's values. If the struct contains a pointer or a slice, the copy contains the *same pointer address*. Modifying the data via that pointer in `b` will affect `a`.

To get a **deep copy**, I must manually allocate new memory for the referenced data (cloning the slice or map). Go doesn't have a built-in `clone()` method."

#### Indepth
Be very careful with structs containing `sync.Mutex`. Copying the struct (by value) copies the Mutex in its current state (locked or unlocked). The copy is a *separate* mutex, leading to subtle race conditions or deadlocks. **Never copy a struct that contains a Mutex.** Pass it by pointer.

---

### 41. How do you convert a struct to JSON?
"I use `json.Marshal(myStruct)`.

It returns a byte slice. The catch is **visibility**: only **Exported fields** (starting with a Capital Letter) are serialized.

I’ve spent hours debugging why my JSON was empty, only to realize I named my field `password` instead of `Password`."

#### Indepth
To marshal private fields, you must implement the `Marshaler` interface (`MarshalJSON()`). This allows you to control the output format completely, like outputting a computed field: `func (u User) MarshalJSON() ([]byte, error) { return json.Marshal(struct{ FullName string }{ u.First + " " + u.Last }) }`.


## From 03 Pointers Interfaces Methods

# 🔵 **41–60: Pointers, Interfaces, and Methods**

### 42. What are pointers in Go?
"A **pointer** is simply a variable that holds the memory address of another variable, rather than the value itself.

In Go, pointers are strictly used for **sharing** data. If I pass a large struct to a function by value, Go copies every byte. If I pass a pointer, I only copy the 8-byte address. This is critical for performance.

Unlike C, Go pointers are safe. I can't do pointer arithmetic (like `ptr + 1`) by default, which eliminates a whole class of buffer overflow bugs."

#### Indepth
You *can* do pointer arithmetic using the `unsafe` package (`unsafe.Pointer` to `uintptr`), but you opt out of Go's type safety and GC guarantees. This is rarely needed outside of low-level system calls or extreme optimization (like serialization libraries).

---

### 43. How do you declare and use pointers?
"I use the `*` operator to declare a pointer type (e.g., `*int`). I use the `&` operator to get the address of a variable.

`x := 10; p := &x`. `p` now holds the address of `x`. To read the value back, I **dereference** it with `*p`.

I frequently use this pattern when I need a function to modify the input variable, like `func increment(x *int) { *x++ }`. Without the pointer, the function would just increment a local copy."

#### Indepth
Pointers in Go automatically handle **dereferencing** for struct fields. If `p` is `*User`, you can write `p.Name` instead of `(*p).Name`. This syntactic sugar makes working with pointers feel very similar to working with values.

---

### 44. What is the difference between pointer and value receivers?
"It comes down to **semantics vs mechanics**.

Mechanically, a **value receiver** (`func (s MyStruct)`) gets a copy. Modifying `s` inside the method touches the copy, not the original. A **pointer receiver** (`func (s *MyStruct)`) gets the address, so modifications affect the caller.

Semantically, if I consider the struct to be an 'entity' with identity (like a User or Database Connection), I always use pointer receivers. If it's a 'value object' (like a Time or Point), I often use value receivers."

#### Indepth
If a type allows `Mutating` methods (pointer receivers), you should generally *only* use pointer receivers for all methods to maintain consistency. Mixing them can lead to subtle bugs where you think you're modifying state but are actually modifying a copy.

---

### 45. What are methods in Go?
"A **method** is just a function with a special 'receiver' argument placed before the function name.

It allows me to attach behavior to *any* user-defined type, not just structs. I can define `func (m MyInt) IsPositive() bool`.

This is powerful because it keeps logic close to data. Instead of passing an object around to utility functions `Process(obj)`, I can just call `obj.Process()`, which reads better."

#### Indepth
Methods can be attached to *any* named type in the same package, except for pointer types or interface types. You can't define methods on `*int`, but you can on `type IntPointer *int`. This flexibility allows extending primitive types with domain-specific logic.

---

### 46. How to define an interface?
"An **interface** defines a contract behavior. It lists a set of method signatures.

`type Shape interface { Area() float64 }`.

The magic of Go is that implementation is **implicit**. I don't write `implements Shape`. If my `Circle` struct has an `Area()` method, it *automatically* satisfies the interface. This prevents the 'header file' maintenance burden seen in other languages."

#### Indepth
This is called **Structural Typing**. It allows for "Consumer-Defined Interfaces". I don't need to import `Shape` to implement it. I just need to have the method. This decouples packages: the low-level library doesn't need to know about the high-level interface.

---

### 47. What is the empty interface in Go?
"The empty interface `interface{}` (or `any` in modern Go) specifies zero methods.

Since every type has at least zero methods, **every type satisfies the empty interface**. `int`, `string`, `struct`, even functions.

I use it when I need to handle data of unknown structure, like decoding arbitrary JSON or implementing a generic container. However, I use it sparingly because it bypasses compile-time type safety—I have to cast it back to a real type to do anything useful."

#### Indepth
Internally, an empty interface is represented by `eface`, a struct containing two pointers: one to the type information (`_type`) and one to the data. This means assigning an `int` to `interface{}` triggers an allocation (boxing) if the value escapes to the heap.

---

### 48. How do you perform type assertion?
"I use the syntax `val.(Type)` to retrieve the dynamic value stored inside an interface.

I almost always use the 'comma-ok' variant: `s, ok := val.(string)`.

If `ok` is true, `s` is the string. If `ok` is false, it failed, but safely. If I skip the `ok` check and the type is wrong, my program crashes with a panic. It’s my runtime safety check."

#### Indepth
Type assertions are fast but not free. The runtime checks the `itab` (interface table) to verify the type matches. If you find yourself doing a "type switch" `switch v := val.(type)` with many cases, consider if polmorphism (adding a method to the interface) would be cleaner.

---

### 49. How to check if a type implements an interface?
"I typically let the compiler check it by assigning the value.

But if I need to check at runtime without crashing, I use a type assertion: `if _, ok := val.(Writer); ok { ... }`.

Occasionally, I add a compile-time guard to my code: `var _ MyInterface = (*MyType)(nil)`. If `MyType` stops implementing the interface (e.g., I renamed a method), the compiler yells at me immediately."

#### Indepth
This pattern `var _ Interface = (*Type)(nil)` is a standard idiom. It has zero runtime cost because it's evaluated at compile time. It acts as a constraint mechanism, similar to `implements` keywords in other languages, but optional.

---

### 50. Can interfaces be embedded?
"Yes, this is how we compose behavior.

The standard library does this with `io.ReadWriter`. It simply embeds `io.Reader` and `io.Writer`.

`type ReadWriter interface { Reader; Writer }`.

This means anything implementable `ReadWriter` must implement both `Read` and `Write`. I use this to create small, reusable interfaces (Lego blocks) and build larger ones only when necessary."

#### Indepth
Interface embedding avoids the "Diamond Problem" of multiple inheritance because interfaces have no state. If two embedded interfaces define the same method `Close()`, the resulting interface simply requires `Close()`. It merges cleanly.

---

### 51. What is polymorphism in Go?
"Polymorphism is achieved entirely through **interfaces**.

I can write a function `Render(s Shape)`. I can pass it a `Circle`, `Rectangle`, or `Triangle`. The function just calls `s.Area()`.

At runtime, Go looks up the concrete type's method implementation in a dispatch table (itab) and calls the correct version. This gives me flexibility to swap implementations (e.g., Real Database vs Mock Database) without changing the core logic."

#### Indepth
Go uses **itable** (interface table) dispatch. When you assign a concrete type to an interface, the runtime generates a table of function pointers matching the interface's methods. Calling a method via an interface is an indirect function call (slightly slower than direct calls but negligible for most apps).

---

### 52. How to use interfaces to write mockable code?
"This is the #1 reason I use interfaces.

Instead of my Service depending on a concrete struct `PostgresDB`, I define an interface `Repository` with methods `GetUser` and `SaveUser`.

In production, I inject the real Postgres struct. In `_test.go`, I define a `MockRepo` struct that also implements `Repository` but returns dummy data. This makes my unit tests fast and deterministic."

#### Indepth
For simple mocks, I manually write a struct. For complex interfaces, I use `go.uber.org/mock` (formerly `gomock`). It generates the mock implementation automatically, allowing me to set expectations like "Method X should be called exactly twice with argument Y".

---

### 53. What is the difference between `interface{}` and `any`?
"There is absolutely **no difference**. `any` is a built-in alias for `interface{}` added in Go 1.18.

It was introduced because typing `interface{}` everywhere in generic code (`[T any]`) looked messy.

I typically use `any` in new code because it reads better (`map[string]any`), but under the hood, it compiles to the exact same empty interface type."

#### Indepth
While `any` is clearer, legacy codebases are full of `interface{}`. Use `gofmt -r 'interface{} -> any'` to modernize code, but ensure your team is on Go 1.18+ before committing. It helps readability significantly in complex map signatures: `map[string]any`.

---

### 54. What is duck typing?
"Duck typing is the concept: 'If it walks like a duck and quacks like a duck, it is a duck.'

Go's implicit interfaces leverage this. I can declare an interface `Quacker` in my consumer package. Any type from any library that has a `Quack()` method satisfies it instantly, even if that library author never heard of my interface.

This decouples dependencies significantly compared to Java/C#."

#### Indepth
Duck typing facilitates the **Adapter Pattern**. If a library returns a `ConcreteLogger` but my code expects a `Logger` interface, I can create an adapter struct that translates my interface calls to the library's methods without modifying the library source code.

---

### 55. Can you create an interface with no methods?
"Yes, that is the `empty interface` (`interface{}` or `any`).

It serves as a universal container. While useful for generic code (like `fmt.Println` which takes `any`), relying on it too much leads to 'stringly typed' code where you have to constantly assert types at runtime.

I verify to replace it with **Generics** (`[T any]`) nowadays whenever possible, keeping type safety at compile time."

#### Indepth
Before generics, `interface{}` was the only way to write containers (List, Set). Now, you should almost always use `[T any]`. Only use `interface{}` when you literally need to store heterogenous data types (e.g., a JSON object with mixed strings/ints).

---

### 56. Can structs implement multiple interfaces?
"Yes, a single struct can implement as many interfaces as its methods satisfy.

A `File` struct often implements `Reader`, `Writer`, `Closer`, and `Seeker`.

This allows different consumers to view the same object through different 'lenses'. A function that just wants to `Close()` it takes a `Closer`; it doesn't care that it can also write bytes."

#### Indepth
This aligns with the **Interface Segregation Principle** (ISP). "Clients should not be forced to depend on methods they do not use." By defining small interfaces (`Reader`), you allow broad compatibility. If you asked for `File`, you'd be coupled to the OS filesystem implementation unnecessarily.

---

### 57. What is the difference between concrete type and interface type?
"A **concrete type** (`int`, `User`) describes the *memory layout* and *exact implementation*. It is what the object actually 'is'.

An **interface type** describes *behavior*. It is an abstract wrapper.

At runtime, an interface value is essentially a tuple: `(type, value)`. It holds a pointer to the concrete type info and a pointer to the actual data. This extra layer of indirection is the cost of polymorphism."

#### Indepth
It's important to know that an interface takes up 16 bytes on a 64-bit system (2 words). Passing parameters as `interface{}` is slightly more expensive than passing concrete types due to this multi-word structure and potential heap allocation for the data value.

---

### 58. How to handle nil interfaces?
"This is a classic Go trap. An interface is `nil` only if **both** the type and value are `nil`.

If I have a pointer `var p *int = nil` and assign it to an interface `var i any = p`, `i` is **not nil**. It holds `(*int, nil)`.

Checking `i == nil` returns `false`. But if I try to use it with reflection, I'll crash. I always verify to check if the underlying value is nil if I'm doing reflection or dealing with error interfaces."

#### Indepth
This behavior happens because an interface is a tuple `(T=Type, V=Value)`. A "nil pointer to int" results in `(T=*int, V=nil)`. The interface itself is only `nil` if `(T=nil, V=nil)`. **Best Practice**: Always return specifically `nil` (the literal), not a typed pointer variable that happens to be nil.

---

### 59. What are method sets?
"Method sets define which methods belong to a type, critical for interface satisfaction.

The rule is: `*T` (pointer) has all methods of `*T` AND `T`. But `T` (value) only has methods declared on `T`.

This means if an interface requires a pointer-receiver method `Modify()`, I **cannot** pass a value `T` to it. It *must* be addressable. This often bites beginners when they try to pass a struct value to a function expecting a pointer-receiver interface."

#### Indepth
This restriction exists because a value `T` might not be addressable (e.g., a temporary value returned by a function, or a map entry). The pointer receiver needs a stable memory address to modify. Since the runtime can't guarantee `T` is addressable, it forbids `T` from satisfying pointer-receiver interfaces.

---

### 60. Can a pointer implement an interface?
"Yes, and usually it should.

If my type needs to mutate state or is large (like a DB connection), I define methods on `*MyType`. Consequently, only `*MyType` implements the interface.

This forces me to pass pointers around, ensuring I'm sharing the single instance rather than copying it, which is exactly what I want for stateful objects."

#### Indepth
If a pointer implements the interface, checking for nil requires care. If you have a `var r Runner` effectively holding a nil pointer, calling `r.Run()` *will* invoke the method! It will only panic if `Run()` tries to dereference the nil receiver. Methods in Go can be called on nil receivers.

---

### 61. What is the use of `reflect` package?
"Reflection allows the program to inspect its own structure at runtime.

I can discover the type of a variable, iterate over struct fields, or call methods dynamically by name.

It is heavily used in libraries like `json`, `xml`, and `orm` to map data to structs. However, in application code, I avoid it. It’s slow, complex, and forfeits type safety. If I find myself using `reflect`, I usually step back and ask if an Interface could solve the problem simpler."

#### Indepth
Reflection is the only way to inspect struct tags. It works by converting an `interface{}` value into a `reflect.Value` and `reflect.Type`. It's powerful but fragile: refactoring field names breaks code using reflection to lookup fields by name strings.


## From 12 Files OS System

# 🟡 **221–240: Files, OS, and System Programming**

### 221. How do you read a file line by line in Go?
"I use `bufio.Scanner`.

`scanner := bufio.NewScanner(file)`
`for scanner.Scan() { line := scanner.Text() }`
This is memory efficient because it doesn't load whole file into RAM.
However, I always check `scanner.Err()` after the loop. If the line is too long (over 64KB default), the scanner might error out, so for massive lines, I switch to `bufio.Reader.ReadLine`."

#### Indepth
`bufio.Scanner` strips the newline character (`\n`) automatically. Be wary of this if you are rewriting the file—you must manually append `\n`. Also, the default buffer size is 64KB. If you are reading lines longer than that (e.g., minified JSON), use `scanner.Buffer()` to increase the limit, or use `bufio.Reader` which can grow indefinitely.

---

### 222. How do you write large files efficiently?
"I use `bufio.Writer`.

Writing 1 byte at a time to disk is a syscall, which is slow.
`bufio.NewWriter(file)` buffers the writes in memory (default 4KB) and flush them to disk in fewer, larger chunks."

#### Indepth
On modern SSDs, writing 4KB vs 16KB doesn't change much, but on networked filesystems (NFS, EFS) or rotating rust, buffering is the difference between minutes and seconds. If you are writing critical data (like a WAL), call `file.Sync()` after `Flush()` to force the OS to physically write headers to the platter/NAND.
**Crucial step**: ALWAYS call `writer.Flush()` before closing the file, or the last chunk of data will be lost forever."

---

### 223. How do you watch file system changes in Go?
"I use the **fsnotify** library (cross-platform).

1.  Create a watcher: `fsnotify.NewWatcher()`.
2.  Add directories: `watcher.Add("/config")`.
3.  Loop over events: `select { case event := <-watcher.Events: ... }`.
It hooks into operating system primitives (inotify on Linux, FSEvents on macOS) to detect file creation, modification, or deletion instantly."

#### Indepth
Watching files recursively is hard. On Linux (`inotify`), you must manually add a watch for *every* subdirectory. If a user does `mkdir -p a/b/c`, you need to catch the creation of `a`, add a watch to it, catch `b`, etc. Libraries like `notify` wrap this complexity but `fsnotify` is the low-level primitive.

---

### 224. How to get file metadata like size, mod time?
"I use `os.Stat(filename)`.

It returns a `FileInfo` interface.
`info.Size()` gives me bytes. `info.ModTime()` gives me the timestamp.
If `os.Stat` returns an error, I check `errors.Is(err, os.ErrNotExist)` to handle the 'file not found' case explicitly."

#### Indepth
`os.Stat` follows symlinks. If you want to check the symlink itself (where it points to), use `os.Lstat`. Also, be aware of TOCTOU (Time Of Check, Time Of Use) bugs: checking if a file exists and then opening it is race-prone. It's better to just open it and handle the error.

---

### 225. How do you work with CSV files in Go?
"I use the standard `encoding/csv` package.

To read: `reader := csv.NewReader(file)`. `record, err := reader.Read()`.
To write: `writer := csv.NewWriter(file)`. `writer.Write([]string{"col1", "col2"})`.
I always `defer writer.Flush()` and check `writer.Error()` to ensure all data was physically written to disk."

#### Indepth
`encoding/csv` handles edge cases like quoted fields containing newlines: `"Hello\nWorld"`. If you are processing massive CSVs (GBs), avoid `ReadAll`. Use the streaming `Read()` in a loop to keep memory usage low. You can also set `reader.ReuseRecord = true` to avoid allocating a new slice for every row (zero-allocation parsing).

---

### 226. How do you compress and decompress files in Go?
"I use `compress/gzip`.

It wraps any `io.Writer`.
`zw := gzip.NewWriter(file)`.
Now, anything I write to `zw` is compressed on the fly and sent to the file.
To read, I wrap with `gzip.NewReader(file)`. It acts just like opening a normal file, but the bytes are decompressed transparently as I read them."

#### Indepth
Gzip is not seekable. You cannot jump to the middle of a `gz` file. For random access compressed data, you need block-based compression (like Snappy or Zstd with framing). Also, `gzip.Writer` buffers heavily; ignoring the `Close()` error means you might silently truncate the file footer.

---

### 227. How do you execute shell commands from Go?
"I use `os/exec`.

`cmd := exec.Command("ls", "-la")`.
I capture the output: `out, err := cmd.CombinedOutput()`.
If I need to stream the output (e.g., for a long-running build script), I pipe `cmd.Stdout` to `os.Stdout`.
I am extremely careful **never** to pass user input directly to `exec.Command` to avoid injection attacks."

#### Indepth
`CombinedOutput` captures both stdout and stderr. If you need to process them separately (e.g., log stderr as errors but parse stdout as JSON), you must assign separate `bytes.Buffer`s to `cmd.Stdout` and `cmd.Stderr`. This gives you granular control over the subprocess output.

---

### 228. What is the `os/exec` package used for?
"It lets my Go program run *other* programs.

I use it heavily for automation: invoking git, running docker builds, or calling legacy binaries.
I usually use `CommandContext` so I can kill the process if it hangs:
`ctx, cancel := context.WithTimeout(ctx, 5*time.Second)`
`exec.CommandContext(ctx, "sleep", "10").Run()`."

#### Indepth
When `CommandContext` kills a process, it sends `SIGKILL` (on POSIX). This gives the child no chance to clean up (delete temp files, release locks). You can customize this by setting `cmd.Cancel` to send `SIGTERM` first, wait a bit, and then kill.

---

### 229. How do you set environment variables in Go?
"For the current process: `os.Setenv("PORT", "8080")`.

To read: `port := os.Getenv("PORT")`.
If I need to distinguish between 'empty' and 'unset', I use `val, ok := os.LookupEnv("PORT")`.
When running a *subprocess* (`exec.Cmd`), I set `cmd.Env` explicitly to pass variables to only that child command."

#### Indepth
Values set with `os.Setenv` are process-wide. This is not thread-safe if you have multiple tests running in parallel that expect different env vars. For testing, it's safer to pass config via checks or dependency injection rather than relying on global environment mutations.

---

### 230. How to create and manage temp files/directories?
"I use `os.CreateTemp` (modern replacement for `ioutil`).

`f, err := os.CreateTemp("", "example-*.txt")`.
The `*` is replaced by a random string to prevent collisions.
I immediately `defer os.Remove(f.Name())` to clean up the file when the function exits. If I need a directory, I use `os.MkdirTemp`."

#### Indepth
The operating system automatically cleans up `/tmp` (or `%TEMP%`) effectively on reboot, but long-running servers can fill up the disk with abandoned temp files. Always use `defer` to clean up. In containerized environments, `/tmp` might be an in-memory `tmpfs`, so writing large temp files could OOM the pod.

---

### 231. How do you handle signals like SIGINT in Go?
"I use `os/signal` to implement graceful shutdowns.

`c := make(chan os.Signal, 1)`
`signal.Notify(c, os.Interrupt, syscall.SIGTERM)`
I block on `<-c`. When the user hits Ctrl+C, I catch it, cancel my context, close my database connections, and exit cleanly. If I didn't effectively catch this, the OS would kill my process mid-write."

#### Indepth
If you capture `SIGINT` but don't exit, the user will be annoyed (they have to `kill -9`). The standard pattern is to listen for the signal *once* to trigger graceful shutdown, and if the user hits Ctrl+C *again* (force quit), the default handler takes over and kills the process immediately.

---

### 232. How do you gracefully shut down a CLI app?
"I bind the OS signal to a Context.
`ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)`.

My main loop looks like:
`select { case <-ctx.Done(): cleanUp(); return }`.
Using `NotifyContext` automatically propagates the cancellation to all my child goroutines (HTTP clients, DB queries), ensuring a hierarchical shutdown."

#### Indepth
Graceful shutdown usually involves a timeout. `Select { case <-ctx.Done(): stop; case <-time.After(30*time.Second): forceStop }`. You don't want your deployment to hang for 10 minutes because one goroutine refused to exit. Kubernetes has a `terminationGracePeriodSeconds` (default 30s) after which it sends SIGKILL.

---

### 233. What are file descriptors and how does Go manage them?
"A file descriptor (FD) is an integer handle for an open resource (file, socket, pipe).

The OS has a limit (ulimit -n), usually 1024 or 65535.
In Go, `os.File` wraps the FD.
If I forget to `f.Close()`, I leak the FD. Eventually, I hit 'too many open files' and the app crashes. Go's GC *might* close it with a Finalizer, but I never rely on that."

#### Indepth
`ulimit` is a common production outage cause. The default on many Linux distros is 1024. A high-throughput web server needs tens of thousands. Always check `ulimit -n` in your startup script or Dockerfile (`LimitNOFILE` in systemd). Network sockets count as FDs too!

---

### 234. How to handle large file uploads and streaming?
"I use `io.Pipe` or direct streaming.

I never load the whole file into RAM (`ioutil.ReadAll` is banned).
`io.Copy(dstFile, requestBody)`.
This streams data in 32KB chunks. I can upload a 50GB file using only a few MB of RAM. It’s strictly IO-bound, not Memory-bound."

#### Indepth
For `io.Pipe`, the writer blocks until the reader reads. This is synchronous. It's powerful for transforming streams (e.g., File -> Gzip -> Encrypt -> S3) without ever holding more than a few kilobytes in RAM. It’s the Unix Philosophy applied to Go interfaces.

---

### 235. How do you access OS-specific syscalls in Go?
"I use `golang.org/x/sys/unix`.

It provides low-level access to kernel calls unimplemented in the standard library.
Examples: `unix.Mlock` (prevent swapping), `unix.Mount`, or specific socket flags (`SO_REUSEPORT`).
I guard this code with build tags (`//go:build linux`) because these calls obviously won't work on Windows."

#### Indepth
`golang.org/x/sys/unix` is preferred over the frozen `syscall` package because it's actively maintained and updated with new kernel features (like `io_uring` or `bpf`). It uses generated code to match the exact kernel headers of the target OS.

---

### 236. How do you implement a simple CLI tool in Go?
"I start with the standard `flag` package.

`name := flag.String("name", "World", "name to greet")`
`flag.Parse()`
If the tool grows (needs subcommands like `git commit`), I switch to **Cobra** or **Urfave CLI**.
They handle help text generation, bash completion, and nested commands (`app server start`) automatically."

#### Indepth
When building CLIs, follow the **12-Factor App** principles. Allow configuration via *both* flags and environment variables. Libraries like `viper` facilitate this precedence: Flag > Env > Config File > Default. This makes your tool friendly for both humans (flags) and scripts/Kubernetes (env).

---

### 237. How do you build cross-platform binaries in Go?
"It’s trivial. I set `GOOS` and `GOARCH`.

Check: `GOOS=linux GOARCH=arm64 go build`.
This produces a binary for Raspberry Pi / AWS Graviton on my MacBook.
I don't need Docker or a VM. The Go compiler knows how to generate machine code for almost every architecture out there."

#### Indepth
Cross-compilation disables CGO by default (`CGO_ENABLED=0`). If your app depends on C libraries (like SQLite), cross-compiling becomes much harder—you need a C cross-compiler (like `zig cc` or `musl-gcc`). For pure Go apps, it's seamless.

---

### 238. What is syscall vs os vs exec package difference?
"**syscall**: Raw kernel interface (deprecated, use `x/sys`). Hard to use.
**os**: Platform-independent wrapper (Open, Read, Write). Use this 99% of the time.
**os/exec**: For running *other* programs.
I use `os` to manipulate files, and `exec` to run scripts."

#### Indepth
The `os` package is designed to be POSIX-like but Windows-compatible. However, file permissions (`chmod`) behave very differently on Windows. Go tries to map them (e.g., ReadOnly bit), but don't expect 0777 to mean the same thing on NTFS as it does on ext4.

---

### 239. How do you write to logs with rotation?
"Standard library `log` does **not** support rotation.

I use **Lumberjack** (`gopkg.in/natefinch/lumberjack.v2`).
It implements `io.Writer`.
`log.SetOutput(&lumberjack.Logger{Filename: "app.log", MaxSize: 100})`.
It automatically closes the file when it hits 100MB, renames it to `app-timestamp.log.gz`, and starts a fresh `app.log`."

#### Indepth
In containerized environments (Kubernetes/Docker), **logs should go to stdout**, not files. Log rotation is the responsibility of the infrastructure (Docker logging driver or Fluentd), not the application. File logging is mostly for legacy VM/bare-metal deployments.

---

### 240. What is the use of `ioutil` and its deprecation?
"`ioutil` was a grab-bag of utility functions (`ReadFile`, `WriteFile`).
It is now **deprecated** (since Go 1.16).

I now use:
*   `os.ReadFile` / `os.WriteFile`.
*   `io.ReadAll`.
*   `io.Discard`.
The old code still works, but modern Go code should use the new, more logical locations (`os` and `io`)."

#### Indepth
One big improvement in the migration: `os.ReadFile` returns the file contents. `ioutil.ReadFile` did the same but was just a wrapper. The new structure organizes functions by *what they act on* (`os` = filesystem, `io` = abstract streams), reducing cyclic dependencies in the standard library.


## From 51 Modern Go Features

# 🆕 **1001–1020: Modern Go Features (v1.22–v1.24)**

### 1001. What is the loop variable scope change in Go 1.22?
"Before 1.22, `for i := range 10` shared the `i` variable across iterations.
This caused the famous `go func() { print(i) }` bug (printing 9, 9, 9...).
In 1.22, `i` is **newly allocated** for each iteration.
We no longer need `i := i` shadowing trick. It just works as expected."

#### Indepth
**The Old Bug**. The classic trap: `for _, v := range items { go func() { process(v) }() }`. All goroutines captured the *same* `v` variable. By the time they ran, the loop had finished and `v` held the last item. The fix was `v := v` inside the loop. Go 1.22 eliminates this entire class of bug by making each iteration's variable a distinct allocation.

---

### 1002. How do you iterate over integers in Go 1.22 (`for i := range n`)?
"Syntactic sugar!
`for i := range 10 { fmt.Println(i) }`.
It prints 0 to 9.
Replaces the verbose `for i := 0; i < 10; i++`.
It’s cleaner and less prone to off-by-one errors."

#### Indepth
**Range over Func**. This is part of a broader "range-over-func" initiative in Go 1.23. The same `range` keyword now works over integers, slices, maps, channels, *and* custom iterator functions. This unification means you learn one syntax and it works everywhere, reducing the cognitive overhead of remembering different loop idioms.

---

### 1003. How does `net/http.ServeMux` support wildcards and methods in Go 1.22?
"Standard library finally got good routing!
`mux.HandleFunc("POST /items/{id}", handler)`.
It supports:
*   **Method Matching**: `POST`, `GET`.
*   **Path Values**: `id := req.PathValue("id")`.
*   **Wildcards**: `/files/{path...}`.
I can finally drop `gorilla/mux` or `chi` for simple projects."

#### Indepth
**Precedence Rules**. The new `ServeMux` has defined precedence: More specific patterns win over less specific ones. `GET /items/{id}` beats `GET /items/`. Method-specific patterns beat method-less ones. This is deterministic and documented, unlike the old `ServeMux` which had subtle ordering bugs when patterns overlapped.

---

### 1004. What is the new `math/rand/v2` package?
"It’s faster and safer.
Global source is now ChaCha8 (cryptographically secure seeding, though still PRNG).
No more `rand.Seed(time.Now())`. It auto-seeds.
Methods like `rand.N(max)` are generic-friendly.
`rand.Intn` is gone; use `rand.IntN`.
Crucially, it removes the global lock contention of v1."

#### Indepth
**ChaCha8 vs PCG**. `math/rand/v2` offers two sources: `rand.New(rand.NewPCG(...))` (fast, for simulations) and `rand.New(rand.NewChaCha8(...))` (cryptographically seeded, default global). The global source uses ChaCha8 to prevent "seed guessing" attacks, even though it's not a CSPRNG. For actual secrets, always use `crypto/rand`.

---

### 1005. What are Go Iterators (`range-over-func`) in Go 1.23?
"Standardized Iterators.
`func(yield func(T) bool)`.
I can use `for v := range mySeq { ... }`.
`mySeq` is just a function that calls `yield(value)` repeatedly.
This unifies iteration over DB rows, API pages, and custom collections without exposing internal state (like `Next()` methods)."

#### Indepth
**Pull vs Push**. Go iterators are "Push" style (the iterator calls `yield`). The alternative is "Pull" style (the consumer calls `Next()`). Go 1.23 also provides `iter.Pull()` to convert a push iterator to a pull iterator when you need to interleave two iterators or need more control. Both styles are now first-class.

---

### 1006. How do you use the `unique` package in Go 1.23?
"It implements **Interning**.
`h := unique.Make("hello")`.
If I have 1 million strings of value 'hello', `unique` stores the string data only once and gives me a lightweight handle (canonical pointer).
Comparing handles (`h1 == h2`) is O(1).
Great for parsers, compilers, or repetitive JSON keys."

#### Indepth
**String Interning**. Without `unique`, two `"hello"` string literals in different packages are separate allocations. With `unique.Make`, they share one. The key benefit is O(1) equality comparison: instead of `strings.Compare` (O(n)), you compare two pointers. This is a massive win in hot paths like symbol tables or HTTP header maps.

---

### 1007. What improvements were made to `time.Timer` garbage collection in Go 1.23?
"Timers are now instantly collectible when unreferenced.
Previously, `time.After(1h)` would leak memory until the hour passed, even if the select finished earlier.
Now, the runtime cleans up the timer immediately if the channel is unreachable.
I don't need `defer timer.Stop()` as religiously anymore (though still good practice)."

#### Indepth
**The Old Leak**. Before 1.23, `select { case <-time.After(1*time.Minute): ... }` in a loop was a classic memory leak. Each iteration created a new timer that lived for a full minute. In a high-throughput loop, you'd accumulate thousands of live timers. The fix was `t := time.NewTimer(d); defer t.Stop()`. Go 1.23 makes the naive version safe.

---

### 1008. What are generic type aliases in Go 1.24?
"I can alias a generic type.
`type MyList[T] = []T`.
Previously, aliases (`=`) only worked on concrete types.
This allows refactoring generic code across packages without breaking the API."

#### Indepth
**Migration Use Case**. Imagine you have `pkg/a` with `type Stack[T any] struct{...}`. You want to move it to `pkg/collections`. Before 1.24, you'd have to update every import. Now: `package a; type Stack[T any] = collections.Stack[T]`. Old code still compiles. You can migrate callsites gradually.

---

### 1009. How do you use the `go tool` directive in `go.mod` (Go 1.24)?
"`go.mod`: `tool golang.org/x/tools/cmd/stringer`.
I can run `go tool stringer`.
It manages tool dependencies **versioning** inside `go.mod`.
No more `tools.go` hack! It ensures every dev uses the exact same linter/generator version."

---

### 1010. What is `os.Root` and how does it improve file system isolation (Go 1.24)?
"`root, _ := os.OpenRoot("/tmp/sandbox")`.
`root.Open("file.txt")`.
If I try `root.Open("../etc/passwd")`, it fails.
It guarantees operations are confined to the directory tree, preventing path traversal attacks safely at the OS level (using `openat` syscalls)."

#### Indepth
**Chroot Alternative**. `os.Root` is safer than `chroot`. `chroot` requires root privileges and can be escaped. `os.Root` uses `openat`/`RESOLVE_NO_ESCAPE` kernel flags, which are enforced by the kernel itself. It works in unprivileged containers and is the correct way to implement sandboxed file access in Go.

---

### 1011. How do you implement weak pointers in Go 1.24?
"`weak.Make(&obj)`.
It returns a pointer that doesn't prevent GC.
If GC runs and `obj` is only held by weak pointers, it collects `obj`.
`ptr := w.Pointer()`. It returns `nil` if collected.
Useful for Caches where I want entries to disappear if memory is tight, without manual eviction."

#### Indepth
**Finalizer Comparison**. `SetFinalizer` had a "resurrection" problem: if the finalizer stored `obj` somewhere, it prevented collection. `weak.Make` avoids this. The GC simply nullifies the weak pointer without running user code. This makes it safe to use in concurrent data structures like `sync.Map`-based caches.

---

### 1012. What is the `omitzero` struct tag option?
"`json:"field,omitzero"`.
Similar to `omitempty`, but checks if the value is the **Zero Value** for its type.
It avoids the confusion where `omitempty` hides `0` or `false` (valid values) when using pointers. `omitzero` is smarter and clearer."

#### Indepth
**The `omitempty` Trap**. `omitempty` omits `0`, `false`, `""`, and `nil`. This means you can't distinguish "user explicitly set score to 0" from "user didn't provide score". The workaround was `*int` (pointer). `omitzero` checks the `IsZero() bool` method if present, allowing types to define their own zero-value semantics.

---

### 1013. How do you use `testing.B.Loop` for benchmarks?
"`for b.Loop() { DoWork() }`.
Replaces `for i := 0; i < b.N; i++`.
It handles the looping logic internally.
It makes setup/teardown logic cleaner (everything outside the loop is setup) and avoids the easy mistake of using `i` incorrectly."

#### Indepth
**Timer Reset**. `b.Loop()` also correctly handles timer reset. With the old `b.N` loop, if you had setup code inside the loop, you'd call `b.ResetTimer()` to exclude it. `b.Loop()` implicitly excludes everything before the first `b.Loop()` call and after the last, making benchmark timing more accurate by default.

---

### 1014. How does Go 1.24 support FIPS 140-3 compliance?
"`GOFIPS=1`.
The standard `crypto` library transparently switches to a FIPS-verified backend (if built with the FIPS toolchain).
It allows government-compliant apps without rewriting code to use specific FIPS SDKs."

#### Indepth
**FIPS 140-3**. This is a US government standard for cryptographic modules. Required for federal contracts, healthcare (HIPAA), and finance. Previously, Go teams had to use BoringCrypto (a Google fork) or external C libraries. Native FIPS support means Go is now a first-class citizen for compliance-heavy industries.

---

### 1015. What are usage comparisons for `slices.Concat`?
"`slices.Concat(s1, s2, s3)`.
It’s optimized.
Allocates the exact final size once.
Memcopies efficiently.
Much faster and cleaner than `append(append(s1, s2...), s3...)`."  

#### Indepth
**Allocation Count**. `append(append(s1, s2...), s3...)` may allocate twice: once for `s1+s2`, once for `(s1+s2)+s3`. `slices.Concat` pre-calculates the total length `len(s1)+len(s2)+len(s3)`, allocates once, and copies. For large slices, this halves memory pressure and eliminates one GC cycle.

---

### 1016. How do you use `runtime.AddCleanup` vs `SetFinalizer`?
"`AddCleanup(ptr, cleanupFunc, arg)`.
It’s the modern, safer `SetFinalizer`.
It allows attaching a cleanup to an object without the resurrection issues of Finalizers.
It's deterministic enough for resource tracking (but still not for urgent resource release—use `defer` for that)."

#### Indepth
**Finalizer Ordering**. `SetFinalizer` had a critical flaw: if object A's finalizer referenced object B, and B was also being finalized, the order was undefined. `AddCleanup` is designed to avoid this by not allowing the cleanup function to access the object being cleaned up (it receives a separate `arg` instead).

---

### 1017. What are the new WASM export capabilities in Go?
"`//go:wasmexport MyFunc`.
Compiles to a WASM function that the host (JS/Rust) can call directly.
No more `js.Global().Set(...)` hackery.
It follows the WASM Component Model standards, making Go WASM modules compatible with other languages."

#### Indepth
**WASM Component Model**. This is the future of WASM interoperability. Instead of sharing raw memory (the old way, requiring manual pointer arithmetic), the Component Model defines high-level types (strings, records, lists). `//go:wasmexport` generates the correct ABI, allowing a Rust host to call a Go function with a `string` argument naturally.

---

### 1018. How do you debug using `go build -asan`?
"**Address Sanitizer** (from C/C++ world).
Detects use-after-free, buffer overflows (in unsafe code/cgo).
Run tests: `go test -asan`.
It adds instrumentation. Slower execution, but catches memory corruption bugs that `-race` might miss."

#### Indepth
**CGO Safety**. `-asan` is most valuable when using CGO. Pure Go is memory-safe by design. But CGO calls C code, which can have buffer overflows. `-asan` instruments both the Go and C sides of the boundary, catching bugs like "C function wrote 5 bytes into a 4-byte Go buffer" that would otherwise cause silent data corruption.

---

### 1019. How do you manage tool dependencies without a `tools.go` file now?
"See Q1009.
I use the `tool` directive in `go.mod`.
`go get -tool golang.org/x/tools/gopls`.
It separates 'application dependencies' (imports) from 'development tools' (binaries).
`go tool` runs them."

#### Indepth
**The Old Hack**. The `tools.go` pattern was: create a file with `//go:build ignore` and blank imports of tools. This forced `go mod tidy` to track them. It was a community workaround, not an official feature. The `go.mod` `tool` directive is the official, clean solution that finally makes tool versioning a first-class concern.

---

### 1020. What is the anticipated "Flight Recorder" feature?
"A lightweight, always-on tracer.
It records execution events to a circular buffer.
If the program crashes, I can dump the buffer.
It tells me what happened in the last 1s before the crash (like a Black Box).
It's designed to be cheap enough for production use."

#### Indepth
**Continuous Profiling**. The Flight Recorder is complementary to tools like Parca or Pyroscope (Continuous Profiling). Those sample CPU/Memory over time. The Flight Recorder captures the exact sequence of events leading to a crash. Together, they give you "What was the system doing overall?" (profiling) AND "What happened in the last second?" (flight recorder).

---
