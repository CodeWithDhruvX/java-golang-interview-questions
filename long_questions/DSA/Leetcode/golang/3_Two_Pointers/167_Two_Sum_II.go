package main

import "fmt"

// 167. Two Sum II - Input Array Is Sorted
// Time: O(N), Space: O(1)
func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	
	for left < right {
		sum := numbers[left] + numbers[right]
		
		if sum == target {
			return []int{left + 1, right + 1} // 1-indexed
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	
	return []int{}
}

func main() {
	// Test cases
	testCases := []struct {
		numbers []int
		target  int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{2, 3, 4}, 6},
		{[]int{-1, 0}, -1},
		{[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8},
		{[]int{-10, -5, -3, 0, 1, 3, 5, 10}, 0},
		{[]int{1, 3, 5, 7, 9}, 12},
	}
	
	for i, tc := range testCases {
		result := twoSumII(tc.numbers, tc.target)
		fmt.Printf("Test Case %d: numbers=%v, target=%d -> Indices: %v\n", 
			i+1, tc.numbers, tc.target, result)
	}
}
