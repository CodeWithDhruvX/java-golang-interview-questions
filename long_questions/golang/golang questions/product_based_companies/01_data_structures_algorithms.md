# 🧮 01 — Data Structures & Algorithms in Go
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

> Product companies like **Google, Meta, Uber** heavily focus on DSA implemented in Go idioms.

---

## 🔑 Must-Know Topics
- Arrays, slices, maps as DSA structures
- Linked lists (single, double)
- Stacks and queues using slices/channels
- Binary trees and BST
- Sorting and searching algorithms
- Common LeetCode-pattern problems in Go

---

## ❓ Most Asked Questions

### Q1. Implement a Stack in Go

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.items) == 0 { return zero, false }
    n := len(s.items)
    item := s.items[n-1]
    s.items = s.items[:n-1]
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    var zero T
    if len(s.items) == 0 { return zero, false }
    return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool { return len(s.items) == 0 }
func (s *Stack[T]) Size() int     { return len(s.items) }

// Usage
s := &Stack[int]{}
s.Push(1); s.Push(2); s.Push(3)
val, _ := s.Pop()    // 3
val, _ = s.Peek()    // 2
```

---

### Q2. Implement a Queue in Go

```go
type Queue[T any] struct {
    items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    var zero T
    if len(q.items) == 0 { return zero, false }
    item := q.items[0]
    q.items = q.items[1:]  // O(n) — for O(1) use two stacks or linked list
    return item, true
}

func (q *Queue[T]) IsEmpty() bool { return len(q.items) == 0 }

// Channel-based concurrent queue
ch := make(chan int, 100)  // buffered channel as queue
ch <- 1; ch <- 2          // enqueue
v := <-ch                 // dequeue — 1
```

---

### Q3. Implement a Linked List in Go

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

type LinkedList struct {
    Head *ListNode
    size int
}

func (ll *LinkedList) Push(val int) {
    ll.Head = &ListNode{Val: val, Next: ll.Head}
    ll.size++
}

func (ll *LinkedList) Append(val int) {
    node := &ListNode{Val: val}
    if ll.Head == nil { ll.Head = node; ll.size++; return }
    curr := ll.Head
    for curr.Next != nil { curr = curr.Next }
    curr.Next = node
    ll.size++
}

func (ll *LinkedList) Print() {
    curr := ll.Head
    for curr != nil {
        fmt.Printf("%d -> ", curr.Val)
        curr = curr.Next
    }
    fmt.Println("nil")
}

// Reverse a linked list — a classic interview question
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
```

---

### Q4. How do you implement binary search in Go?

```go
// Iterative (preferred)
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2  // avoids overflow
        if nums[mid] == target { return mid }
        if nums[mid] < target { left = mid + 1 } else { right = mid - 1 }
    }
    return -1
}

// Use built-in sort.SearchInts
import "sort"
nums := []int{1, 3, 5, 7, 9, 11}
idx := sort.SearchInts(nums, 7)  // returns index 3
```

---

### Q5. How do you sort in Go?

```go
import "sort"

// Sort ints/strings/floats
nums := []int{5, 2, 8, 1, 9}
sort.Ints(nums)      // [1 2 5 8 9]
strs := []string{"banana", "apple", "cherry"}
sort.Strings(strs)   // [apple banana cherry]

// Custom sort with sort.Slice
type Person struct{ Name string; Age int }
people := []Person{{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}}
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age  // sort by age ascending
})

// Stable sort (preserves order of equal elements)
sort.SliceStable(people, func(i, j int) bool {
    return people[i].Name < people[j].Name
})
```

---

### Q6. Implement a Binary Search Tree (BST)

```go
type BSTNode struct {
    Val         int
    Left, Right *BSTNode
}

func insert(root *BSTNode, val int) *BSTNode {
    if root == nil { return &BSTNode{Val: val} }
    if val < root.Val { root.Left = insert(root.Left, val) } else
    if val > root.Val { root.Right = insert(root.Right, val) }
    return root
}

func search(root *BSTNode, val int) bool {
    if root == nil { return false }
    if root.Val == val { return true }
    if val < root.Val { return search(root.Left, val) }
    return search(root.Right, val)
}

// Inorder traversal (sorted output)
func inorder(root *BSTNode) {
    if root == nil { return }
    inorder(root.Left)
    fmt.Printf("%d ", root.Val)
    inorder(root.Right)
}
```

---

### Q7. How do you use maps for common algorithmic problems?

```go
// Frequency counter
func charFrequency(s string) map[rune]int {
    freq := make(map[rune]int)
    for _, c := range s { freq[c]++ }
    return freq
}

// Two Sum — O(n) with map
func twoSum(nums []int, target int) (int, int) {
    seen := make(map[int]int)  // value → index
    for i, n := range nums {
        complement := target - n
        if j, ok := seen[complement]; ok {
            return j, i
        }
        seen[n] = i
    }
    return -1, -1
}

// Check anagram
func isAnagram(s, t string) bool {
    if len(s) != len(t) { return false }
    freq := make(map[rune]int)
    for _, c := range s { freq[c]++ }
    for _, c := range t {
        freq[c]--
        if freq[c] < 0 { return false }
    }
    return true
}
```

---

### Q8. Implement a Min-Heap in Go

```go
import "container/heap"

// IntMinHeap implements heap.Interface
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() interface{} {
    old := *h; n := len(old)
    x := old[n-1]; *h = old[:n-1]
    return x
}

// Usage
h := &IntMinHeap{5, 1, 3, 2, 4}
heap.Init(h)
heap.Push(h, 0)
fmt.Println(heap.Pop(h))  // 0 (min element)
```

---

### Q9. Common sliding window pattern in Go

```go
// Maximum sum subarray of size k
func maxSumSubarray(nums []int, k int) int {
    if len(nums) < k { return 0 }

    windowSum := 0
    for i := 0; i < k; i++ { windowSum += nums[i] }
    maxSum := windowSum

    for i := k; i < len(nums); i++ {
        windowSum += nums[i] - nums[i-k]
        if windowSum > maxSum { maxSum = windowSum }
    }
    return maxSum
}

// Longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    maxLen, left := 0, 0
    for right := 0; right < len(s); right++ {
        if idx, ok := charIndex[s[right]]; ok && idx >= left {
            left = idx + 1
        }
        charIndex[s[right]] = right
        if right-left+1 > maxLen { maxLen = right - left + 1 }
    }
    return maxLen
}
```
