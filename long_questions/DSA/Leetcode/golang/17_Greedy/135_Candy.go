package main

import "fmt"

// 135. Candy
// Time: O(N), Space: O(N)
func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	
	n := len(ratings)
	candies := make([]int, n)
	
	// Initialize each child with 1 candy
	for i := 0; i < n; i++ {
		candies[i] = 1
	}
	
	// Left to right pass
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	
	// Right to left pass
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	
	// Calculate total candies
	total := 0
	for _, candy := range candies {
		total += candy
	}
	
	return total
}

// Optimized version with O(1) space (but more complex)
func candyOptimized(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	
	n := len(ratings)
	total := 1
	up := 1
	down := 0
	oldSlope := 0
	
	for i := 1; i < n; i++ {
		newSlope := 0
		if ratings[i] > ratings[i-1] {
			newSlope = 1
		} else if ratings[i] < ratings[i-1] {
			newSlope = -1
		}
		
		if oldSlope > 0 && newSlope == 0 || oldSlope < 0 && newSlope >= 0 {
			total += up * (up + 1) / 2
			up = 1
		}
		
		if newSlope < 0 {
			down++
		} else {
			total += down * (down + 1) / 2
			down = 0
		}
		
		if newSlope > 0 {
			up++
		}
		
		if newSlope == 0 {
			total++
		}
		
		oldSlope = newSlope
	}
	
	total += up * (up + 1) / 2
	if down > 0 {
		total += down * (down + 1) / 2
	}
	
	return total
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 0, 2},
		{1, 2, 2},
		{1, 3, 4, 5, 2},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 2, 2, 1},
		{2, 2, 2, 2},
		{1},
		{1, 2},
		{2, 1},
		{1, 3, 2, 2, 1},
		{1, 2, 87, 87, 87, 2, 1},
	}
	
	for i, ratings := range testCases {
		result1 := candy(ratings)
		result2 := candyOptimized(ratings)
		fmt.Printf("Test Case %d: %v -> Standard: %d, Optimized: %d\n", 
			i+1, ratings, result1, result2)
	}
}
