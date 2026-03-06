# 🗣️ Theory — Data Structures & Algorithms in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How do you implement a Stack in Go? What's the idiomatic approach?"

> *"Go doesn't have a built-in stack, so you implement it yourself using a slice. A slice is a perfect backing structure because `append` gives you push in O(1) amortized, and taking the last element then re-slicing to exclude it gives you pop in O(1). Since Go 1.18 added generics, the cleanest approach is a generic stack: `type Stack[T any] struct { items []T }`. That way you get a type-safe stack that works with integers, strings, or any type. Before generics, you'd use `interface{}` and type-assert on pop, which was less clean."*

---

## Q: "How does sorting work in Go?"

> *"Go's `sort` package has built-in functions for common types: `sort.Ints`, `sort.Strings`, `sort.Float64s`. For custom sorting, you use `sort.Slice` which takes a comparison function — `sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })`. For stable sort — where equal elements keep their original order — use `sort.SliceStable`. Under the hood, Go uses pattern-defeating quicksort — a hybrid of quicksort, heapsort, and insertion sort that's O(n log n) worst case. For generic sorting in Go 1.21+, `slices.Sort` and `slices.SortFunc` are the new idiomatic way."*

---

## Q: "How do you use maps for algorithmic problems in Go?"

> *"Maps are Go's hash table, and they're central to many algorithm solutions. Key patterns: frequency counting — iterate over a slice and increment `freq[val]++`; two-pointer with tracking — store the index of each value in a map; memoization — cache results of expensive recursive calls. Important things to know for interviews: Go maps have no guaranteed iteration order; reading a missing key returns the zero value, not an error; you check existence with `v, ok := m[key]`; and maps passed to functions are reference types — the caller sees your mutations."*

---

## Q: "What is binary search and how do you implement it in Go?"

> *"Binary search works on a sorted slice by repeatedly halving the search space. You have a left and right pointer. Calculate the midpoint — `mid := left + (right-left)/2`, note the `(right-left)/2` form to avoid integer overflow. If the target matches the mid element, return its index. If target is greater, search the right half by moving left to mid+1. If target is smaller, search the left half by moving right to mid-1. O(log n) time. Go's standard library has `sort.SearchInts` for sorted int slices. A common interview trick: binary search can be applied anywhere you have a monotone predicate, not just on sorted arrays."*

---

## Q: "How do you implement a min-heap in Go?"

> *"Go's `container/heap` package provides a heap that works generically via an interface. You implement five methods on your type: `Len`, `Less`, `Swap` from `sort.Interface`, plus `Push` and `Pop`. The `Less` function determines whether it's a min-heap or max-heap — if `h[i] < h[j]` you get a min-heap. Then you call `heap.Init(&h)`, `heap.Push(&h, val)`, and `heap.Pop(&h)`. The heap package is a bit verbose but it's flexible — you can have heaps of any type including structs, making it useful for Dijkstra's, priority queues, and k-th largest element problems."*

---

## Q: "What is the sliding window technique? How have you used it?"

> *"Sliding window is a technique for problems on contiguous subarrays or substrings where you maintain a window of elements and slide it through the array. Instead of recalculating from scratch for each window — which would be O(n*k) — you maintain a running value and update it incrementally by adding the new element entering the window and removing the element leaving it, giving O(n). Classic examples: maximum sum subarray of size k, longest substring without repeating characters, minimum window substring. For variable-size windows, you use two pointers — left expands the window by moving right and shrinks it when a condition is violated."*

---

## Q: "How do you work with linked lists in Go interviews?"

> *"Linked list problems in Go use a `ListNode` struct with a `Val` and a `Next` pointer — same pattern as LeetCode. The key patterns to know: reversal — you maintain prev, curr, and next pointers and rewire them; cycle detection — Floyd's tortoise and hare algorithm with slow and fast pointers; finding the middle — slow/fast pointer where fast moves two steps, slow moves one; merging sorted lists — compare heads and recurse or iterate. Go's lack of pointer arithmetic actually makes linked list code cleaner than C — you just have `*ListNode` and Go handles the memory."*
