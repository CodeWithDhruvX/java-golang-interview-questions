package main

import "fmt"

// 84. Largest Rectangle in Histogram
// Time: O(N), Space: O(N)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []int{-1} // Initialize with -1 to handle edge case
	maxArea := 0
	
	for i := 0; i < n; i++ {
		// While current height is less than height at stack top
		for stack[len(stack)-1] != -1 && heights[i] < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	
	// Process remaining bars in stack
	for stack[len(stack)-1] != -1 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		width := n - stack[len(stack)-1] - 1
		area := height * width
		if area > maxArea {
			maxArea = area
		}
	}
	
	return maxArea
}

func main() {
	// Test cases
	testCases := [][]int{
		{2, 1, 5, 6, 2, 3},
		{2, 4},
		{1, 1, 1, 1},
		{4, 2, 0, 3, 2, 5},
		{6, 5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6},
		{2, 1, 2, 3, 1},
		{0},
		{},
		{1},
		{1000, 1000, 1000},
	}
	
	for i, heights := range testCases {
		result := largestRectangleArea(heights)
		fmt.Printf("Test Case %d: %v -> Largest rectangle area: %d\n", 
			i+1, heights, result)
	}
}
