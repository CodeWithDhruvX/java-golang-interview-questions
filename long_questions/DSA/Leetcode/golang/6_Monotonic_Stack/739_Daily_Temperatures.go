package main

import "fmt"

// 739. Daily Temperatures
// Time: O(N), Space: O(N)
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // Store indices
	
	for i := 0; i < n; i++ {
		// While current temperature is greater than temperature at stack top
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	
	// Remaining indices in stack have no warmer day (result is already 0)
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
		{30, 40, 50, 60},
		{30, 60, 90},
		{90, 80, 70, 60},
		{55, 60, 65, 70, 75},
		{65, 70, 65, 60, 65},
		{50},
		{},
		{73, 73, 73, 73},
		{30, 40, 30, 50, 30, 60, 30},
	}
	
	for i, temperatures := range testCases {
		result := dailyTemperatures(temperatures)
		fmt.Printf("Test Case %d: %v -> Days until warmer: %v\n", 
			i+1, temperatures, result)
	}
}
