package main

import "fmt"

// 496. Next Greater Element I
// Time: O(N), Space: O(N)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// First pass: find next greater for all elements in nums2
	nextGreater := make(map[int]int)
	stack := []int{}
	
	for _, num := range nums2 {
		// While stack is not empty and current element is greater than stack top
		for len(stack) > 0 && num > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[top] = num
		}
		stack = append(stack, num)
	}
	
	// Elements remaining in stack have no next greater element
	for _, num := range stack {
		nextGreater[num] = -1
	}
	
	// Second pass: build result for nums1
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}
	
	return result
}

func main() {
	// Test cases
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
		{[]int{2, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 3, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{2}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1}, []int{1}},
	}
	
	for i, tc := range testCases {
		result := nextGreaterElement(tc.nums1, tc.nums2)
		fmt.Printf("Test Case %d: nums1=%v, nums2=%v -> Next greater: %v\n", 
			i+1, tc.nums1, tc.nums2, result)
	}
}
