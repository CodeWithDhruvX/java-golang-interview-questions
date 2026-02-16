package main

import (
	"container/heap"
)

// 1. Rotate Array by K Places (Without Extra Space - O(1))
// Time: O(N), Space: O(1)
// Note: Already covered in Basic Questions but reiterating the Reversal Algorithm as it's the standard O(1) space approach.
// Logic: Reverse(0, n-1) -> Reverse(0, k-1) -> Reverse(k, n-1) for Right Rotation.
// For Left Rotation: Reverse(0, k-1) -> Reverse(k, n-1) -> Reverse(0, n-1).
func RotateArray(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}
	k %= n
	// Right Rotation
	reverse(arr, 0, n-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// 2. Find K-th Smallest Element (QuickSelect)
// Time: O(N) average, O(N^2) worst case.
// Space: O(1) (iterative) or O(log N) (recursive stack)
func FindKthSmallest(arr []int, k int) int {
	if k < 1 || k > len(arr) {
		return -1 // Error
	}
	return quickSelect(arr, 0, len(arr)-1, k-1)
}

func quickSelect(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}
	pivotIndex := partition(arr, left, right)
	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelect(arr, left, pivotIndex-1, k)
	} else {
		return quickSelect(arr, pivotIndex+1, right, k)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	pIndex := left
	for i := left; i < right; i++ {
		if arr[i] <= pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++
		}
	}
	arr[pIndex], arr[right] = arr[right], arr[pIndex]
	return pIndex
}

// 3. Find Median of Stream of Numbers (Two Heaps)
// Time: O(log N) insert, O(1) find
// Space: O(N)
type MedianFinder struct {
	low  *MaxHeap
	high *MinHeap
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		low:  &MaxHeap{},
		high: &MinHeap{},
	}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.low, num)
	heap.Push(mf.high, heap.Pop(mf.low))

	if mf.low.Len() < mf.high.Len() {
		heap.Push(mf.low, heap.Pop(mf.high))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.low.Len() > mf.high.Len() {
		return float64(mf.low.Peek())
	}
	return (float64(mf.low.Peek()) + float64(mf.high.Peek())) * 0.5
}

// Standard Heap Implementations
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MinHeap) Peek() int { return (*h)[0] }

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // Max Heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MaxHeap) Peek() int { return (*h)[0] }

// 4. Partition Array Into 3 Parts With Equal Sum
// Time: O(N), Space: O(1)
func CanThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	target := sum / 3
	parts := 0
	currSum := 0

	for _, v := range arr {
		currSum += v
		if currSum == target {
			parts++
			currSum = 0
		}
	}
	// We need at least 3 parts.
	// If sum is 0, we can have more than 3 parts (e.g. 0,0,0,0), return true if parts >= 3
	// If sum != 0, parts should be exactly 3? No, logic is simply finding 3 segments.
	// If we found 3 segments with target sum, remainder must be 0 sum?
	// Actually typical implementation:
	// Find i where sum[0..i] == target
	// Find j where sum[i+1..j] == target
	// Check if remaining has sum target.
	// The loop above counts how many times we see cumulative sum = target.
	// If parts >= 3, return true.
	return parts >= 3
}

// 5. Minimum Jumps to Reach End (Greedy - Jump Game II)
// Time: O(N), Space: O(1)
func Jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == currentEnd {
			jumps++
			currentEnd = farthest
			if currentEnd >= n-1 {
				break
			}
		}
	}
	return jumps
}

// 6. Check If Array Pairs Are Divisible by k
// Time: O(N), Space: O(N)
func CanArrange(arr []int, k int) bool {
	freq := make(map[int]int)
	for _, v := range arr {
		rem := v % k
		if rem < 0 {
			rem += k
		}
		freq[rem]++
	}

	for rem, count := range freq {
		if rem == 0 {
			if count%2 != 0 {
				return false
			}
		} else if 2*rem == k { // Special case for k/2
			if count%2 != 0 {
				return false
			}
		} else {
			if freq[k-rem] != count {
				return false
			}
		}
	}
	return true
}
