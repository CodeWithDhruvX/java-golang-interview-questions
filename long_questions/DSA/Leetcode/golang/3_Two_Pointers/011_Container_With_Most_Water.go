package main

import "fmt"

// 11. Container With Most Water
// Time: O(N), Space: O(1)
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0
	
	for left < right {
		// Calculate current area
		width := right - left
		height1 := height[left]
		height2 := height[right]
		currentHeight := height1
		if height2 < height1 {
			currentHeight = height2
		}
		
		currentArea := width * currentHeight
		if currentArea > maxWater {
			maxWater = currentArea
		}
		
		// Move the pointer with smaller height
		if height1 < height2 {
			left++
		} else {
			right--
		}
	}
	
	return maxWater
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
		{4, 3, 2, 1, 4},
		{1, 2, 1},
		{2, 3, 4, 5, 18, 17, 6},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 3, 2, 5, 25, 24, 5},
	}
	
	for i, height := range testCases {
		result := maxArea(height)
		fmt.Printf("Test Case %d: %v -> Max Area: %d\n", i+1, height, result)
	}
}
