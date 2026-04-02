package main

import "fmt"

// 1295. Find Numbers with Even Number of Digits
// Time: O(N), Space: O(1)
func findNumbers(nums []int) int {
	count := 0
	
	for _, num := range nums {
		if hasEvenDigits(num) {
			count++
		}
	}
	
	return count
}

// Helper function to check if a number has even number of digits
func hasEvenDigits(num int) bool {
	if num == 0 {
		return false
	}
	
	digitCount := 0
	for num > 0 {
		num /= 10
		digitCount++
	}
	
	return digitCount%2 == 0
}

func main() {
	// Test cases
	testCases := [][]int{
		{12, 345, 2, 6, 7896},
		{555, 901, 482, 1771},
		{1, 22, 333, 4444},
		{0, 10, 100, 1000},
		{},
	}
	
	for i, nums := range testCases {
		result := findNumbers(nums)
		fmt.Printf("Test Case %d: %v -> Numbers with even digits: %d\n", i+1, nums, result)
	}
}
