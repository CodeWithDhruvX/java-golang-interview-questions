# Collections Programs (76-85)

## 76. Iterate Slice
**Principle**: For-loop, Range, Iterator pattern.
**Question**: Iterate over a slice in different ways.
**Code**:
```go
package main

import "fmt"

func main() {
    slice := []string{"A", "B"}
    
    // 1. For loop with index
    for i := 0; i < len(slice); i++ {
        fmt.Print(slice[i])
    }
    fmt.Println()
    
    // 2. Range loop
    for _, s := range slice {
        fmt.Print(s)
    }
    fmt.Println()
    
    // 3. Iterator pattern (manual)
    iter := &SliceIterator{slice: slice, index: 0}
    for iter.HasNext() {
        fmt.Print(iter.Next())
    }
    fmt.Println()
}

type SliceIterator struct {
    slice []string
    index int
}

func (s *SliceIterator) HasNext() bool {
    return s.index < len(s.slice)
}

func (s *SliceIterator) Next() string {
    if s.HasNext() {
        val := s.slice[s.index]
        s.index++
        return val
    }
    return ""
}
```

## 77. Slice vs Linked List
**Principle**: Demonstration of types.
**Code**:
```go
package main

// Slice (Fast Access)
type SliceList []string

// Linked List (Fast Insert/Delete)
type Node struct {
    Value string
    Next  *Node
}

type LinkedList struct {
    Head *Node
}

func main() {
    slice := make(SliceList, 0)
    linkedList := &LinkedList{}
    
    _ = slice
    _ = linkedList
}
```

## 78. Convert Array to Slice
**Principle**: Use slice conversion.
**Question**: Convert Array to Slice.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := [2]string{"A", "B"}
    slice := arr[:]
    fmt.Printf("%T %v\n", slice, slice)
}
```

## 79. Iterate Map
**Principle**: Use range loop.
**Question**: Iterate over a Map.
**Code**:
```go
package main

import "fmt"

func main() {
    m := make(map[int]string)
    m[1] = "A"
    
    for key, value := range m {
        fmt.Printf("%d=%s\n", key, value)
    }
}
```

## 80. Sort Map by Value
**Principle**: Convert to slice of key-value pairs, sort.
**Question**: Sort a Map by its values.
**Code**:
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    m := map[string]int{
        "A": 3,
        "B": 1,
    }
    
    // Convert to slice of key-value pairs
    type KV struct {
        Key   string
        Value int
    }
    
    var kvs []KV
    for k, v := range m {
        kvs = append(kvs, KV{k, v})
    }
    
    // Sort by value
    sort.Slice(kvs, func(i, j int) bool {
        return kvs[i].Value < kvs[j].Value
    })
    
    for _, kv := range kvs {
        fmt.Printf("%s=%d\n", kv.Key, kv.Value)
    }
}
```

## 81. Remove Element while Iterating
**Principle**: Create new slice or use index manipulation.
**Question**: Remove elements safely during iteration.
**Code**:
```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    
    // Method 1: Create new slice
    result := []int{}
    for _, val := range slice {
        if val != 2 {
            result = append(result, val)
        }
    }
    fmt.Println(result)
    
    // Method 2: In-place removal
    slice = []int{1, 2, 3}
    for i := 0; i < len(slice); i++ {
        if slice[i] == 2 {
            slice = append(slice[:i], slice[i+1:]...)
            i-- // Adjust index
        }
    }
    fmt.Println(slice)
}
```

## 82. Merge Two Slices
**Principle**: Use append.
**Question**: Merge two slices.
**Code**:
```go
package main

import "fmt"

func main() {
    slice1 := []string{"A"}
    slice2 := []string{"B"}
    
    merged := append(slice1, slice2...)
    fmt.Println(merged)
}
```

## 83. Find Intersection of Two Slices
**Principle**: Use map for lookup.
**Question**: Find common elements in two slices.
**Code**:
```go
package main

import "fmt"

func main() {
    slice1 := []int{1, 2, 3}
    slice2 := []int{2, 3, 4}
    
    // Create map for fast lookup
    lookup := make(map[int]bool)
    for _, val := range slice2 {
        lookup[val] = true
    }
    
    // Find intersection
    var intersection []int
    for _, val := range slice1 {
        if lookup[val] {
            intersection = append(intersection, val)
        }
    }
    
    fmt.Println(intersection) // [2, 3]
}
```

## 84. Thread-Safe Slice
**Principle**: Use mutex or sync package.
**Question**: How to satisfy thread-safety for Slice.
**Code**:
```go
package main

import (
    "sync"
)

type SafeSlice struct {
    mu    sync.RWMutex
    slice []string
}

func (s *SafeSlice) Append(item string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.slice = append(s.slice, item)
}

func (s *SafeSlice) Get(index int) string {
    s.mu.RLock()
    defer s.mu.RUnlock()
    if index >= 0 && index < len(s.slice) {
        return s.slice[index]
    }
    return ""
}

func main() {
    safeSlice := &SafeSlice{slice: []string{}}
    safeSlice.Append("test")
    _ = safeSlice.Get(0)
}
```

## 85. Reverse a Slice
**Principle**: Manual reversal or use algorithm.
**Question**: Reverse a Slice.
**Code**:
```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    
    // Method 1: Manual reversal
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
    fmt.Println(slice)
    
    // Method 2: Create new reversed slice
    slice = []int{1, 2, 3}
    reversed := make([]int, len(slice))
    for i, val := range slice {
        reversed[len(slice)-1-i] = val
    }
    fmt.Println(reversed)
}
```
