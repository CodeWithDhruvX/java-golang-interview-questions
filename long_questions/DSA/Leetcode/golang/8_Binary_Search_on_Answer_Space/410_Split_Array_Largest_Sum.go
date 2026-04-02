package main

import "fmt"

// 410. Split Array Largest Sum
// Time: O(N log S), Space: O(1) where S is sum(nums)
func splitArray(nums []int, k int) int {
	left, right := maxNum(nums), sumNums(nums)
	result := right
	
	for left <= right {
		mid := left + (right-left)/2
		
		if canSplit(nums, k, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return result
}

func canSplit(nums []int, k, maxSum int) bool {
	subarrays := 1
	currentSum := 0
	
	for _, num := range nums {
		if currentSum+num <= maxSum {
			currentSum += num
		} else {
			subarrays++
			currentSum = num
			if subarrays > k {
				return false
			}
		}
	}
	
	return subarrays <= k
}

func maxNum(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func sumNums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{7, 2, 5, 10, 8}, 2},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{1, 4, 4}, 3},
		{[]int{10, 5, 2, 7, 8, 9}, 3},
		{[]int{2, 3, 1, 2, 4, 3}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{[]int{100}, 1},
		{[]int{1, 1, 1, 1, 1}, 5},
		{[]int{5, 5, 5, 5}, 2},
	}
	
	for i, tc := range testCases {
		result := splitArray(tc.nums, tc.k)
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Min largest sum: %d\n", 
			i+1, tc.nums, tc.k, result)
	}
}
