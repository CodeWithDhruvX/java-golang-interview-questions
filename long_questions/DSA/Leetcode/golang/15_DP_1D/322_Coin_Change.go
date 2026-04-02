package main

import (
	"fmt"
	"math"
)

// 322. Coin Change
// Time: O(S*N), Space: O(S) where S is amount, N is number of coins
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	
	// dp[i] = minimum number of coins needed to make amount i
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	
	// Build dp table
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	
	return dp[amount]
}

// Bottom-up DP with early optimization
func coinChangeOptimized(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	
	// Sort coins for potential early termination
	for i := 0; i < len(coins); i++ {
		for j := i + 1; j < len(coins); j++ {
			if coins[i] > coins[j] {
				coins[i], coins[j] = coins[j], coins[i]
			}
		}
	}
	
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	
	return dp[amount]
}

// Recursive with memoization
func coinChangeMemo(coins []int, amount int) int {
	memo := make(map[int]int)
	return coinChangeHelper(coins, amount, memo)
}

func coinChangeHelper(coins []int, amount int, memo map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return math.MaxInt32
	}
	
	if val, exists := memo[amount]; exists {
		return val
	}
	
	minCoins := math.MaxInt32
	for _, coin := range coins {
		subResult := coinChangeHelper(coins, amount-coin, memo)
		if subResult != math.MaxInt32 {
			minCoins = min(minCoins, subResult+1)
		}
	}
	
	memo[amount] = minCoins
	return minCoins
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 2, 5}, 11},
		{[]int{2}, 3},
		{[]int{1}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 5}, 100},
		{[]int{2, 5, 10, 1}, 27},
		{[]int{186, 419, 83, 408}, 6249},
		{[]int{1, 3, 4, 5}, 7},
		{[]int{2, 5, 10, 1}, 0},
	}
	
	for i, tc := range testCases {
		result1 := coinChange(tc.coins, tc.amount)
		result2 := coinChangeOptimized(tc.coins, tc.amount)
		result3 := coinChangeMemo(tc.coins, tc.amount)
		
		fmt.Printf("Test Case %d: coins=%v, amount=%d\n", i+1, tc.coins, tc.amount)
		fmt.Printf("  Standard: %d, Optimized: %d, Memo: %d\n\n", result1, result2, result3)
	}
}
