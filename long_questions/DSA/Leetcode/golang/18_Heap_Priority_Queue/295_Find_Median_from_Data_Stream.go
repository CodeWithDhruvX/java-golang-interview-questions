package main

import (
	"container/heap"
	"fmt"
)

// 295. Find Median from Data Stream
// Time: O(log N) for addNum, O(1) for findMedian, Space: O(N)
type MedianFinder struct {
	maxHeap *MaxHeap // For smaller half
	minHeap *MinHeap // For larger half
}

// MaxHeap for smaller half
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MinHeap for larger half
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

/** initialize your data structure here. */
func Constructor() MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	
	return MedianFinder{
		maxHeap: maxHeap,
		minHeap: minHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// Add to maxHeap first
	heap.Push(this.maxHeap, num)
	
	// Move the largest element from maxHeap to minHeap
	if this.maxHeap.Len() > 0 {
		max := heap.Pop(this.maxHeap).(int)
		heap.Push(this.minHeap, max)
	}
	
	// Balance the heaps
	if this.minHeap.Len() > this.maxHeap.Len() {
		min := heap.Pop(this.minHeap).(int)
		heap.Push(this.maxHeap, min)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64((*this.maxHeap)[0])
	}
	
	return (float64((*this.maxHeap)[0]) + float64((*this.minHeap)[0])) / 2.0
}

func main() {
	// Test cases
	testCases := []struct {
		operations []string
		values    [][]int
	}{
		{
			[]string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {1}, {2}, {}, {1}, {}},
		},
		{
			[]string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {5}, {}, {15}, {}, {1}, {}},
		},
		{
			[]string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {2}, {3}, {4}, {}, {1}, {}},
		},
		{
			[]string{"MedianFinder", "addNum", "findMedian"},
			[][]int{{}, {-1}, {}},
		},
		{
			[]string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "addNum", "findMedian"},
			[][]int{{}, {0}, {0}, {}, {0}, {0}, {}},
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		
		var mf MedianFinder
		for j, op := range tc.operations {
			switch op {
			case "MedianFinder":
				mf = Constructor()
				fmt.Printf("  Created MedianFinder\n")
			case "addNum":
				mf.AddNum(tc.values[j][0])
				fmt.Printf("  Added: %d\n", tc.values[j][0])
			case "findMedian":
				median := mf.FindMedian()
				fmt.Printf("  Median: %.1f\n", median)
			}
		}
		fmt.Println()
	}
}
