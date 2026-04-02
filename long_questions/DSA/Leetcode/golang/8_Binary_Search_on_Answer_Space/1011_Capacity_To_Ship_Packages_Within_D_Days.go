package main

import "fmt"

// 1011. Capacity To Ship Packages Within D Days
// Time: O(N log M), Space: O(1) where M is sum(weights)
func shipWithinDays(weights []int, days int) int {
	left, right := maxWeight(weights), sumWeights(weights)
	result := right
	
	for left <= right {
		mid := left + (right-left)/2
		
		if canShip(weights, days, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return result
}

func canShip(weights []int, days, capacity int) bool {
	daysNeeded := 1
	currentLoad := 0
	
	for _, weight := range weights {
		if currentLoad+weight <= capacity {
			currentLoad += weight
		} else {
			daysNeeded++
			currentLoad = weight
			if daysNeeded > days {
				return false
			}
		}
	}
	
	return daysNeeded <= days
}

func maxWeight(weights []int) int {
	max := 0
	for _, weight := range weights {
		if weight > max {
			max = weight
		}
	}
	return max
}

func sumWeights(weights []int) int {
	sum := 0
	for _, weight := range weights {
		sum += weight
	}
	return sum
}

func main() {
	// Test cases
	testCases := []struct {
		weights []int
		days    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{[]int{3, 2, 2, 4, 1, 4}, 3},
		{[]int{1, 2, 3, 1, 1}, 4},
		{[]int{10, 50, 100, 100, 50, 10}, 5},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{100}, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 9},
		{[]int{5, 5, 5, 5, 5}, 3},
	}
	
	for i, tc := range testCases {
		result := shipWithinDays(tc.weights, tc.days)
		fmt.Printf("Test Case %d: weights=%v, days=%d -> Min capacity: %d\n", 
			i+1, tc.weights, tc.days, result)
	}
}
