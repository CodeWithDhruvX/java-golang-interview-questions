package main

import "fmt"

// 42. Trapping Rain Water
// Time: O(N), Space: O(N)
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	
	// Calculate left max for each position
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	
	// Calculate right max for each position
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	
	// Calculate trapped water
	water := 0
	for i := 0; i < n; i++ {
		water += min(leftMax[i], rightMax[i]) - height[i]
	}
	
	return water
}

// Alternative solution using two pointers (O(1) space)
func trapTwoPointers(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0
	
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	
	return water
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{2, 0, 2},
		{3, 0, 0, 2, 0, 4},
		{0, 0, 0, 0},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 2, 2, 2},
		{1},
		{},
		{0, 2, 0},
		{4, 2, 0, 3, 2, 5, 2, 1, 5, 2},
	}
	
	for i, height := range testCases {
		result1 := trap(height)
		result2 := trapTwoPointers(height)
		fmt.Printf("Test Case %d: %v -> Water (O(N) space): %d, Water (O(1) space): %d\n", 
			i+1, height, result1, result2)
	}
}
