package main

import "fmt"

// 303. Range Sum Query - Immutable
// Time: O(1) query, Space: O(N) for preprocessing
type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		queries []struct {
			left  int
			right int
		}
	}{
		{
			[]int{-2, 0, 3, -5, 2, -1},
			[]struct {
				left  int
				right int
			}{{0, 2}, {2, 5}, {0, 5}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]struct {
				left  int
				right int
			}{{0, 0}, {0, 4}, {1, 3}, {2, 2}},
		},
		{
			[]int{},
			[]struct {
				left  int
				right int
			}{},
		},
		{
			[]int{-1, -1, -1, -1},
			[]struct {
				left  int
				right int
			}{{0, 3}, {1, 2}, {0, 1}},
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: nums=%v\n", i+1, tc.nums)
		numArray := Constructor(tc.nums)
		
		for j, query := range tc.queries {
			result := numArray.SumRange(query.left, query.right)
			fmt.Printf("  Query %d: sumRange(%d, %d) = %d\n", 
				j+1, query.left, query.right, result)
		}
		fmt.Println()
	}
}
