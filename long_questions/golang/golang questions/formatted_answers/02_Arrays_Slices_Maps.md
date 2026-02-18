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
