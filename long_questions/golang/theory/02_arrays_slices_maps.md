# 🟡 Go Theory Questions: 21–40 Arrays, Slices, and Maps

## 21. What is the difference between an Array and a Slice?

**Answer:**
Fundamentally, an **Array** is a fixed-size block of memory where the size is part of the type—`[5]int` and `[10]int` are completely different, incompatible types. A **Slice**, on the other hand, is a dynamic window or view over an underlying array.

Mechanically, an array is a "value type"—if you pass it to a function, Go copies the entire block of memory. A slice is a lightweight "reference type" (a header of just 24 bytes) that points to the array.

In practice, you rarely use arrays directly unless you need a fixed buffer (like a sha256 hash). You almost always use slices because they can grow, shrink, and be passed around efficiently without copying massive amounts of data.

---

## 22. How do you append to a slice?

**Answer:**
We use the built-in `append()` function, which adds elements to the end of a slice and returns the updated slice.

Under the hood, `append` checks if the underlying array has enough **capacity**. If it does, it simply puts the value in the next slot. If it doesn't, it allocates a completely new, larger array (usually double the size), copies all existing data over, and then adds the new element. This is why you must always reassign the result: `slice = append(slice, elem)`.

Realistically, while `append` is convenient, relying on it inside tight loops can trigger frequent memory allocations. If you know the size beforehand, it's always better to pre-allocate using `make` to avoid that reallocation cost.

---

## 23. What happens when a slice is appended beyond its capacity?

**Answer:**
When you exceed capacity, Go triggers a reallocation event. The runtime allocates a new backing array that is larger than the old one—typically doubling the size for small slices, or growing by about 25% for very large ones to be memory efficient.

It effectively pauses to **copy** every single element from the old location to the new memory address. The old array is then abandoned and eventually garbage collected if nothing else references it.

This "grow-and-copy" strategy is generally fast enough (`O(1)` amortized), but in high-performance hot paths, it creates CPU spikes and GC pressure. That's why pre-allocating is a best practice.

---

## 24. How do you copy slices?

**Answer:**
You use the built-in `copy(dst, src)` function, which copies data from a source slice to a destination slice.

The critical thing to remember is that `copy` only moves as many elements as the **smaller** of the two slices can hold. It does not allocate memory for you. If your destination slice has a length of 0, `copy` does absolutely nothing, no matter how much data is in the source.

We use this often when we want to take a snapshot of data—say, to process it on a background thread without race conditions—or when we want to extract a small piece of a huge array so the garbage collector can free the rest of the memory.

---

## 25. What is the difference between `len()` and `cap()`?

**Answer:**
`len()` is the number of elements *currently* in the slice that you can access. `cap()` is the total space available in the underlying backing array starting from the slice's pointer.

Think of it like a window. `len` is how wide the window is open right now. `cap` is how wide the window frame allows it to open before you need to install a completely new frame (reallocate).

This distinction matters heavily for performance. You can extend a slice up to its capacity using reslicing syntax `s[:cap(s)]` without allocating new memory. We use this trick constantly in buffer pooling to reuse memory and reduce GC overhead.

---

## 26. How do you create a multi-dimensional slice?

**Answer:**
Go doesn't have true multi-dimensional arrays in the C sense. Instead, we create a **"slice of slices"**—like `[][]int`.

This means strictly speaking, you have a "Jagged Array." The outer slice holds pointers to inner slices, and each inner slice can technically have a different length and live in a completely different part of memory.

While flexible, this structure isn't cache-friendly because iterating through it involves chasing pointers (pointer indirection) rather than scanning a contiguous block. For high-performance matrix math, we usually flatten the data into a single 1D slice and calculate indices manually (`y * width + x`).

---

## 27. How are slices passed to functions (by value or reference)?

**Answer:**
Technically, everything in Go is passed by **value**. However, a slice is just a tiny "Header" struct containing a pointer, length, and capacity.

So when you pass a slice, you are copying that header, but the **Pointer** inside it still points to the same original underlying array. This creates a "reference-like" behavior. If the function modifies an index (`s[0] = 1`), the caller sees that change.

The catch is `append`. If the function appends to the slice and triggers a reallocation, the function gets a *new* pointer to a *new* array, but the caller is still holding the *old* pointer. That's why you often see functions returning the updated slice.

---

## 28. What are maps in Go?

**Answer:**
A map is Go’s built-in hash table implementation, providing key-value storage with `O(1)` average lookups.

Internally, it uses a system of **buckets**. When you assign a key, Go hashes it to find the right bucket and stores the entry there. A unique feature of Go maps is that iteration order is randomized by design—the runtime actually picks a random starting bucket every time you loop over it.

We use maps everywhere: as caches, as sets (using `map[string]bool`), or for quick lookups. The main trade-off is they are **not thread-safe**; writing to a map from multiple goroutines without a lock will crash your entire program instantly.

---

## 29. How do you check if a key exists in a map?

**Answer:**
We use the **"Comma OK" idiom**.

When you access a map, you can request two return values: `val, ok := myMap[key]`. If `ok` is true, the key exists. If `false`, the key is missing, and `val` is just the zero value for that type.

This is safer than checking for nil or 0, because sometimes 0 is a valid value (like a temperature of 0 degrees). Without the `ok` check, you can't distinguish between "stored 0" and "missing". It’s a standard pattern you’ll see in almost every Go codebase.

---

## 30. Can maps be compared directly?

**Answer:**
No, `map1 == map2` is a compile error in Go. You can only check if a map is `nil`.

The reason is that map equality is ill-defined. Should we compare insertion order? Should we handle cycles? Since comparison could be extremely expensive (O(N)), Go forces you to write your own loop to compare keys and values if you really need it.

For unit tests, we usually rely on `reflect.DeepEqual()` or helper libraries like `testify/assert`, but in production code, we generally avoid comparing maps directly due to the performance cost.

---

## 31. What happens if you delete a key from a map that doesn’t exist?

**Answer:**
Nothing happens. It’s a **no-op**.

The `delete(m, key)` function is safe to call even if the key is missing. Go won't panic or throw an error.

This is convenient for cleanup logic—you can blindly delete session IDs or cache entries without wasting CPU cycles checking if they exist first. It simplifies the code, though it does mean you don't get feedback if your logic was trying to delete something that should have been there but wasn't.

---

## 32. Can slices be used as map keys?

**Answer:**
No, because slices are not **comparable**.

Map keys must support the `==` operator so the runtime can verify hash collisions. Since slices don't support `==` (because they are mutable references), they can't be keys. Arrays, however, *can* be keys if their element types are comparable.

If you desperately need a list as a key—for example, caching query results based on a list of IDs—you typically convert the slice to a string or use a helper struct to hash the values manually before using them in the map.

---

## 33. How do you iterate over a map?

**Answer:**
You use the `for range` loop: `for k, v := range myMap`.

As I mentioned earlier, the iteration order is **random**. Run the loop twice, and you’ll likely get the keys in a different sequence. This is a deliberate design choice by the Go team to prevent developers from relying on the internal memory layout of the hash table.

If you need a consistent order (like for a JSON API response), you must extract the keys into a slice, sort that slice, and then iterate over the sorted keys to look up the values.

---

## 34. How do you sort a map by key or value?

**Answer:**
Maps are unsorted by definition, so you can't sort the map itself. You have to sort a **projection** of the map.

The standard approach is to grab all the keys, put them in a slice, and use `sort.Strings()` or `slices.Sort()`. Then you iterate over that sorted slice and pull values from the map.

It’s an O(N log N) operation. For small maps, it's negligible, but if you're building a system that needs usually-sorted data, a map might be the wrong data structure—you might want a B-Tree or a sorted slice instead.

---

## 35. What are struct types in Go?

**Answer:**
A struct is a custom data type that groups related fields together—it's the Go equivalent of a Class in Java or Python, but without the inheritance magic.

Mechanically, a struct is a contiguous block of memory. If you have a struct with two `int64` fields, it takes up exactly 16 bytes. This makes them extremely cache-efficient compared to objects in dynamic languages.

We use them for everything: representing database rows, JSON payloads, or configuration objects. They are the fundamental building block of domain modeling in Go.

---

## 36. How do you define and use struct tags?

**Answer:**
Struct tags are string metadata attached to field declarations, like `` `json:"user_id"` ``.

The compiler ignores them, but they are accessible at runtime via **Reflection**. Libraries like standard `encoding/json` or ORMs read these tags to map your struct fields to JSON keys or database columns.

It’s a powerful declarative system, but it’s essentially "stringly typed." If you typo a tag, the compiler won't warn you; your code will just silently fail to unmarshal that field. So, we often rely on linters (`go vet`) to catch syntax errors in tags.

---

## 37. How to embed one struct into another?

**Answer:**
Embedding is Go’s simplified version of inheritance, strictly based on **Composition**.

You define a struct as a field without a name: `type Manager struct { User }`. This "promotes" the fields and methods of `User` up to `Manager`. You can call `manager.Name` directly, even though `Name` actually belongs to the inner `User`.

It’s great for code reuse, but it’s not polymorphism. A `Manager` is *not* a `User`. You cannot pass a `Manager` to a function expecting a `User`. It’s just syntactic sugar to avoid typing `manager.User.Name` everywhere.

---

## 38. How do you compare two structs?

**Answer:**
You can use the standard `==` operator, but only if **every field** in the struct is comparable.

If your struct contains only simple types like ints and strings, `s1 == s2` compares them value-by-value. But if your struct contains a slice, map, or function pointer, trying to use `==` will cause a **compiler error**.

In those complex cases, you have to use `reflect.DeepEqual()`, which is slow, or—better yet—write a custom `.Equal()` method that checks the fields relevant to your business logic.

---

## 39. What is the difference between shallow and deep copy in structs?

**Answer:**
Assignable struct copying in Go is **shallow** by default.

When you say `b := a`, Go copies the struct’s memory bits. If the struct contains a pointer or a slice header, `b` gets a copy of that pointer, pointing to the **same** underlying data as `a`. Changing data via `b`'s pointer will affect `a`.

To get a **Deep Copy**—where you duplicate the underlying data too—you have to write it yourself. You’d create a `Clone()` method that allocates new memory for all the slices and pointers and copies the values over. We typically avoid deep copies unless necessary for concurrency safety to avoid race conditions.

---

## 40. How do you convert a struct to JSON?

**Answer:**
We use the `encoding/json` package, specifically `json.Marshal(struct)`.

It uses **reflection** to inspect your struct, reads the exported fields (starting with Uppercase letters), respects any struct tags you defined, and constructs a JSON byte slice.

It’s widely used, but be aware it has a runtime cost. For high-performance systems processing millions of events, we often switch to code-generation tools like `easyjson` which generate static marshalling code, bypassing reflection completely for speed.
