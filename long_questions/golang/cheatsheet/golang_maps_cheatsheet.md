# Golang Maps & Hashing Cheatsheet

Everything you need to know about Go Maps for interviews, from internal mechanics to common patterns.

## 🟢 Map Internals (The "How It Works" Question)

### 1. The Structure (`hmap`)
Go maps are implemented as hash tables using **buckets**.
- **Header:** Contains metadata like `count` (size), `flags` (writing state), `B` (log_2 of buckets).
- **Buckets:** An array where data is stored. Each bucket holds up to **8 key-value pairs**.
- **Overflow Buckets:** If a bucket is full (hash collision), it links to an overflow bucket.

### 2. Time Complexity
- **Access/Insert/Delete:** **O(1)** on average.
- **Worst Case:** **O(N)** (if many collisions occur, though Go's hash seed makes this rare).

### 3. Hashing Mechanism
- Go uses a hash function (AES-based on standard hardware) to generate a hash for the key.
- **Low-Order Bits:** Used to select the **bucket**.
- **High-Order Bits:** Used to distinguish entries **within** the bucket (Top Hash).

### 4. Evacuation (Resizing)
- When the map grows too large (Load Factor > 6.5), Go allocates a new array of buckets (**doubles the size**).
- Entries are **incrementally moved** ("evacuated") to the new buckets as you access/write to the map, not all at once (avoids latency spikes).

---

## 🟡 Thread Safety & Race Conditions

### 1. ARE MAPS THREAD-SAFE?
❌ **NO.**
Concurrent reads are safe. **Concurrent read/write or write/write is NOT safe** and will cause a fatal runtime panic (`concurrent map writes`).

### 2. How to Handle Concurrency?
#### Option A: `sync.RWMutex` (Standard)
Best for most general use cases.
```go
type SafeMap struct {
    mu sync.RWMutex
    data map[string]int
}

func (s *SafeMap) Get(key string) int {
    s.mu.RLock()         // Read Lock
    defer s.mu.RUnlock()
    return s.data[key]
}

func (s *SafeMap) Set(key string, val int) {
    s.mu.Lock()          // Write Lock
    defer s.mu.Unlock()
    s.data[key] = val
}
```

#### Option B: `sync.Map` (Specialized)
Best for:
1. **Append-only caches** (keys written once, read many times).
2. **Disjoint sets** (multiple goroutines working on disjoint keys).
```go
var m sync.Map

// Store
m.Store("key", 42)

// Load
val, ok := m.Load("key")

// Delete
m.Delete("key")

// Range
m.Range(func(key, value interface{}) bool {
    fmt.Println(key, value)
    return true // continue iteration
})
```

---

## 🔵 Common Interview Patterns

### 1. Frequency Counter
Count occurrences of elements.
```go
func charCount(s string) map[rune]int {
    counts := make(map[rune]int)
    for _, char := range s {
        counts[char]++
    }
    return counts
}
// "hello" -> {'h':1, 'e':1, 'l':2, 'o':1}
```

### 2. Set Implementation
Go doesn't have a `set` type. Use `map[T]struct{}` (consumes 0 bytes for value).
```go
func uniqueElements(arr []int) []int {
    set := make(map[int]struct{})
    var result []int
    
    for _, v := range arr {
        if _, exists := set[v]; !exists {
            set[v] = struct{}{}
            result = append(result, v)
        }
    }
    return result
}
```

### 3. Grouping / Anagrams
Group strings by a sorted key or character count.
```go
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    
    for _, s := range strs {
        key := sortString(s) // hypothetical sorting helper
        groups[key] = append(groups[key], s)
    }
    
    var result [][]string
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}
```

### 4. Two Sum (Wait for value)
Check if complement exists in map while iterating.
```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, num := range nums {
        if idx, ok := seen[target-num]; ok {
            return []int{idx, i}
        }
        seen[num] = i
    }
    return nil
}
```

---

## 🟣 Gotchas & Tricks

### 1. Iteration Order is Random
Never rely on `range` order. It is deliberately randomized to prevent dependency on hash implementation detals.
**Fix:** Maintain a separate slice of keys if order matters.

### 2. "Nil" Map vs "Empty" Map
```go
var m1 map[string]int        // nil map
m2 := make(map[string]int)   // empty map

// Reading: SAFE for both (returns zero value)
fmt.Println(m1["foo"]) // 0
fmt.Println(m2["foo"]) // 0

// Writing: PANIC for nil map
m2["foo"] = 1 // OK
m1["foo"] = 1 // PANIC! "assignment to entry in nil map"
```

### 3. Map Values are Not Addressable
You cannot take the address of a map value.
```go
m := map[string]int{"a": 1}
p := &m["a"] // COMPILER ERROR
```
**Why?** Because map growth (evacuation) might move the value to a different bucket memory address.

### 4. Deleting from Map
`delete(map, key)` is safe even if the key doesn't exist or if the map is nil.

### 5. Clearing a Map (Go 1.21+)
```go
clear(m) // Removes all keys, keeps memory allocated for reuse
```
Before 1.21: `m = make(map[K]V)` (old map garbage collected).
