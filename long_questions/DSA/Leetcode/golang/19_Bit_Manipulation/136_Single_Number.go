package main

import "fmt"

// 136. Single Number
// Time: O(N), Space: O(1)
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// Alternative approach using hash map (O(N) space)
func singleNumberHash(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	
	for num, freq := range count {
		if freq == 1 {
			return num
		}
	}
	
	return -1 // Should not reach here for valid input
}

// XOR properties explanation
func singleNumberWithExplanation(nums []int) int {
	fmt.Printf("XOR Properties:\n")
	fmt.Printf("1. a ^ a = 0 (XOR of same numbers is 0)\n")
	fmt.Printf("2. a ^ 0 = a (XOR with 0 is the number itself)\n")
	fmt.Printf("3. XOR is commutative and associative\n\n")
	
	fmt.Printf("Processing: ")
	result := 0
	for i, num := range nums {
		if i > 0 {
			fmt.Printf(" ^ ")
		}
		fmt.Printf("%d", num)
		result ^= num
	}
	fmt.Printf(" = %d\n\n", result)
	
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
		{2, 2, 3, 3, 4},
		{0, 1, 1},
		{-1, -1, -2},
		{99, 99, 100},
		{5, 5, 6, 6, 7, 7, 8},
		{10, 10, 10, 10, 15},
	}
	
	for i, nums := range testCases {
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		
		if i == 0 {
			// Show detailed explanation for first test case
			result := singleNumberWithExplanation(nums)
			fmt.Printf("Single number: %d\n\n", result)
		} else {
			result1 := singleNumber(nums)
			result2 := singleNumberHash(nums)
			fmt.Printf("XOR: %d, Hash: %d\n\n", result1, result2)
		}
	}
}
