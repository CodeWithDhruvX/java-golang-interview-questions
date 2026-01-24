package main

import (
	"container/heap"
	"fmt"
)

// Pattern: Heap / Priority Queue
// Difficulty: Medium
// Key Concept: Efficiently maintaining the "Top K" or "Smallest/Largest" element in a changing dataset.

/*
INTUITION:
"Kth Largest Element"
You have 1,000,000 numbers. You want the 5th largest.
Sorting takes O(N log N). Can we do faster?
Yes. We can maintain a "Hall of Fame" list of size 5.
As we scan numbers, if a number is bigger than the *smallest* person in our Hall of Fame, we kick the smallest person out and put the new one in.

Structure:
- We need efficient access to the "Smallest" in our Top K group.
- A **Min-Heap** of size K gives us the Smallest element in O(1)!
- Pushing/Popping takes O(log K).

Algorithm:
1. Initialize a Min-Heap.
2. Iterate through array.
3. Push number to Heap.
4. If Heap size > K: Pop the smallest (Root). (This removes the "worst" of the top candidates).
5. At the end, the heap contains the Top K Largest elements. The Root is the Kth largest.

Time: O(N log K). (Much better than O(N log N) if K is small).
*/

// --- Go Heap Boilerplate (Standard Library requires implementing interface) ---
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// --- Algorithm ---

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	// DRY RUN: [3,2,1,5,6,4], k=2
	//
	// i=0 (3): Push 3. Heap: [3]
	// i=1 (2): Push 2. Heap: [2, 3]
	// i=2 (1): Push 1. Heap: [1, 3, 2]. Size=3 (>2)! Pop 1. Heap: [2, 3].
	// i=3 (5): Push 5. Heap: [2, 3, 5]. Size=3! Pop 2. Heap: [3, 5].
	// i=4 (6): Push 6. Heap: [3, 5, 6]. Size=3! Pop 3. Heap: [5, 6].
	// i=5 (4): Push 4. Heap: [4, 6, 5]. Size=3! Pop 4. Heap: [5, 6].
	//
	// End. Root is 5.

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2

	fmt.Printf("Nums: %v, K: %d\n", nums, k)
	res := findKthLargest(nums, k)
	fmt.Printf("Kth Largest: %d\n", res) // Expected: 5
}
