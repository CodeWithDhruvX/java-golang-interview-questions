package main

import "fmt"

// 134. Gas Station
// Time: O(N), Space: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	currentGas := 0
	start := 0
	
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentGas += gas[i] - cost[i]
		
		// If current gas is negative, we can't start from current start
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}
	
	// If total gas is less than total cost, no solution exists
	if totalGas < totalCost {
		return -1
	}
	
	return start
}

func main() {
	// Test cases
	testCases := []struct {
		gas  []int
		cost []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
		{[]int{2, 3, 4}, []int{3, 4, 3}},
		{[]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}},
		{[]int{3, 3, 4}, []int{3, 4, 4}},
		{[]int{1, 2, 3, 4, 5, 5, 70}, []int{2, 3, 4, 3, 4, 5, 50}},
		{[]int{2}, []int{2}},
		{[]int{1}, []int{2}},
		{[]int{5, 8, 2, 8}, []int{6, 5, 6, 6}},
		{[]int{4, 5, 3, 1, 4}, []int{5, 4, 3, 4, 2}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
	}
	
	for i, tc := range testCases {
		result := canCompleteCircuit(tc.gas, tc.cost)
		fmt.Printf("Test Case %d: gas=%v, cost=%v -> Start: %d\n", 
			i+1, tc.gas, tc.cost, result)
	}
}
