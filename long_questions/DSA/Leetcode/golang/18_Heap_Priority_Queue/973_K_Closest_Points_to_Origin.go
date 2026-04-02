package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 973. K Closest Points to Origin
// Time: O(N log K), Space: O(K)
func kClosest(points [][]int, k int) [][]int {
	// Use a max-heap of size k to keep k closest points
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	
	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		heap.Push(maxHeap, Point{point, distance})
		
		if maxHeap.Len() > k {
			heap.Pop(maxHeap)
		}
	}
	
	// Extract points from heap
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(maxHeap).(Point).coords
	}
	
	return result
}

type Point struct {
	coords   []int
	distance int
}

// MaxHeap implementation for Point
type MaxHeap []Point

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Alternative solution using QuickSelect (O(N) average case)
func kClosestQuickSelect(points [][]int, k int) [][]int {
	distances := make([]int, len(points))
	for i, point := range points {
		distances[i] = point[0]*point[0] + point[1]*point[1]
	}
	
	kth := quickSelect(distances, 0, len(distances)-1, k)
	
	result := make([][]int, 0, k)
	for i, point := range points {
		if distances[i] <= kth {
			result = append(result, point)
		}
	}
	
	return result
}

func quickSelect(distances []int, left, right, k int) int {
	if left == right {
		return distances[left]
	}
	
	pivotIndex := partition(distances, left, right)
	
	if k == pivotIndex {
		return distances[pivotIndex]
	} else if k < pivotIndex {
		return quickSelect(distances, left, pivotIndex-1, k)
	} else {
		return quickSelect(distances, pivotIndex+1, right, k)
	}
}

func partition(distances []int, left, right int) int {
	pivot := distances[right]
	i := left
	
	for j := left; j < right; j++ {
		if distances[j] <= pivot {
			distances[i], distances[j] = distances[j], distances[i]
			i++
		}
	}
	
	distances[i], distances[right] = distances[right], distances[i]
	return i
}

func main() {
	// Test cases
	testCases := []struct {
		points [][]int
		k      int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2},
		{[][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, 3},
		{[][]int{{1, 0}, {0, 1}, {1, 1}}, 2},
		{[][]int{{-5, -5}, {-4, -4}, {-3, -3}, {-2, -2}, {-1, -1}}, 3},
		{[][]int{{100, 100}, {-100, -100}, {100, -100}, {-100, 100}}, 2},
		{[][]int{{0, 1}, {1, 0}}, 1},
		{[][]int{{2, 2}, {2, 2}, {2, 2}, {2, 2}}, 3},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}, 4},
	}
	
	for i, tc := range testCases {
		// Make copies for both methods
		points1 := make([][]int, len(tc.points))
		copy(points1, tc.points)
		points2 := make([][]int, len(tc.points))
		copy(points2, tc.points)
		
		result1 := kClosest(points1, tc.k)
		result2 := kClosestQuickSelect(points2, tc.k)
		
		fmt.Printf("Test Case %d: points=%v, k=%d\n", i+1, tc.points, tc.k)
		fmt.Printf("  Heap: %v\n", result1)
		fmt.Printf("  QuickSelect: %v\n\n", result2)
	}
}
