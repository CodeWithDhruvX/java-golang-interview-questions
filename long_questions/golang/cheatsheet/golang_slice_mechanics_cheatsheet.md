# Golang Slice Mechanics Cheatsheet

A deep dive into how Slices work under the hood. Crucial for performance tuning and advanced interviews.

## 🟢 Internal Structure (The "Slice Header")

A slice is a lightweight descriptor (struct) containing 3 fields:

```go
type SliceHeader struct {
    Data uintptr // Pointer to the underlying array
    Len  int     // Number of elements in the slice
    Cap  int     // Capacity (elements in array starting from Data)
}
```

### Key Concept: Slices share storage
Multiple slices can point to the **same underlying array**.
- Modifying elements in one slice **modifies them in others** if they share the same memory range.
- `append` **may or may not** allocate a new array (see below).

---

## 🟡 The "Append" Mechanics

### 1. Growth Strategy (Re-allocation)
When you `append` to a full slice (`len == cap`):
1. Go allocates a **new, larger array**.
2. Copies existing data to the new array.
3. Appends the new value.
4. Returns a slice pointing to the **new array**.

**Growth Rate:**
- **< 256 elements:** Capacity **doubles (2x)**.
- **> 256 elements:** Capacity grows by factor of **~1.25x** (smoother transition).
*(Note: Exact threshold changed in Go 1.18, previously 1024)*

### 2. Gotcha: Disconnected Slices
```go
a := make([]int, 0, 3) 
b := append(a, 1)      // b underlying array = [1, _, _]
c := append(a, 2)      // c overwrites indices! underlying = [2, _, _]

// Since a, b, c share capacity 3, they fight for the same slots.
```

---

## 🔵 Nil vs. Empty Slices

Both have `len` = 0 and `cap` = 0, but they are semantically different.

| Feature | `var s []int` (Nil) | `s := []int{}` (Empty) |
| :--- | :--- | :--- |
| **Pointer** | `nil` | Non-nil (zerobase) |
| **JSON Marshal** | `null` | `[]` |
| **Equality** | `s == nil` is TRUE | `s == nil` is FALSE |
| **Usage** | Best for "no result" | Best for "0 results found" |

**Functional equivalence:** You can `append`, `len`, `cap`, and `range` over a `nil` slice safely.

---

## 🟣 Common Pitfalls & Performance

### 1. Memory Leaks (The "Sub-slice" Problem)
If you slice a small chunk from a **huge** array, the huge array stays in memory because the small slice references it.

**Problem:**
```go
func getHeader() []byte {
    raw := loadHugeFile() // 100MB array
    return raw[:10]       // Returns tiny slice, but keeps 100MB allocated!
}
```

**Fix (Copy):**
```go
func getHeader() []byte {
    raw := loadHugeFile()
    res := make([]byte, 10)
    copy(res, raw[:10])   // Allocates tiny new array
    return res            // 100MB array can now be GC'd
}
```

### 2. Pre-allocating Performance
Always pre-allocate if you know the size to avoid multiple resizes `O(N)`.

```go
// BAD: re-allocates multiple times
s := []int{}
for i := 0; i < 10000; i++ { s = append(s, i) }

// GOOD: no re-allocation
s := make([]int, 0, 10000)
for i := 0; i < 10000; i++ { s = append(s, i) }
```

### 3. Range Copy Gotcha
The value in a `for range` loop is a **copy**. Capturing its address works on the *loop variable*, not the slice element.

```go
// BAD
refs := []*int{}
for _, v := range nums {
    refs = append(refs, &v) // All point to the same address!
}

// GOOD
for i := range nums {
    refs = append(refs, &nums[i])
}
```

---

## 🔦 2D Slices

Go does not have true multi-dimensional arrays (like C matrix). It has "slices of slices".

```go
// Create
rows, cols := 5, 5
matrix := make([][]int, rows)
for i := range matrix {
    matrix[i] = make([]int, cols) // allocate each row
}
// Rows are technically independent arrays in memory.
```

---

## 🔁 Copy/Cloning

### Deep Copy (Slice)
```go
original := []int{1, 2, 3}
clone := make([]int, len(original))
copy(clone, original)
```

### Shallow Copy
```go
s1 := []int{1, 2, 3}
s2 := s1 // s2 points to same array
```
