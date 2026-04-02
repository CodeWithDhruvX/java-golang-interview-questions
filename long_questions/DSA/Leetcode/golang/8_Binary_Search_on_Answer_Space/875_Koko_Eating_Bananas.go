package main

import (
	"fmt"
	"math"
)

// 875. Koko Eating Bananas
// Time: O(N log M), Space: O(1) where M is max(piles)
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, maxPiles(piles)
	result := right
	
	for left <= right {
		mid := left + (right-left)/2
		
		if canEatAll(piles, h, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return result
}

func canEatAll(piles []int, h, speed int) bool {
	hours := 0
	for _, pile := range piles {
		hours += int(math.Ceil(float64(pile) / float64(speed)))
		if hours > h {
			return false
		}
	}
	return hours <= h
}

func maxPiles(piles []int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	return max
}

func main() {
	// Test cases
	testCases := []struct {
		piles []int
		h     int
	}{
		{[]int{3, 6, 7, 11}, 8},
		{[]int{30, 11, 23, 4, 20}, 5},
		{[]int{30, 11, 23, 4, 20}, 6},
		{[]int{1}, 1},
		{[]int{100}, 1},
		{[]int{100}, 100},
		{[]int{312884470}, 312884469},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 10},
		{[]int{5, 10, 15, 20, 25}, 15},
	}
	
	for i, tc := range testCases {
		result := minEatingSpeed(tc.piles, tc.h)
		fmt.Printf("Test Case %d: piles=%v, h=%d -> Min speed: %d\n", 
			i+1, tc.piles, tc.h, result)
	}
}
