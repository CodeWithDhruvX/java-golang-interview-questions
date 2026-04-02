package main

import (
	"container/heap"
	"fmt"
)

// 215. Kth Largest Element in an Array
// Time: O(N log K), Space: O(K)
func findKthLargest(nums []int, k int) int {
	// Use a min-heap of size k
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	
	for _, num := range nums {
		heap.Push(minHeap, num)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	
	return (*minHeap)[0]
}

// MinHeap implementation
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Alternative solution using QuickSelect (O(N) average case)
func findKthLargestQuickSelect(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, kthSmallest int) int {
	if left == right {
		return nums[left]
	}
	
	pivotIndex := partition(nums, left, right)
	
	if kthSmallest == pivotIndex {
		return nums[pivotIndex]
	} else if kthSmallest < pivotIndex {
		return quickSelect(nums, left, pivotIndex-1, kthSmallest)
	} else {
		return quickSelect(nums, pivotIndex+1, right, kthSmallest)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
		{[]int{1}, 1},
		{[]int{2, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 1},
		{[]int{7, 10, 4, 3, 20, 15}, 3},
		{[]int{-1, -2, -3, -4, -5}, 2},
		{[]int{100, 200, 300, 400, 500}, 4},
	}
	
	for i, tc := range testCases {
		// Make copies for both methods
		nums1 := make([]int, len(tc.nums))
		copy(nums1, tc.nums)
		nums2 := make([]int, len(tc.nums))
		copy(nums2, tc.nums)
		
		result1 := findKthLargest(nums1, tc.k)
		result2 := findKthLargestQuickSelect(nums2, tc.k)
		
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Heap: %d, QuickSelect: %d\n", 
			i+1, tc.nums, tc.k, result1, result2)
	}
}
