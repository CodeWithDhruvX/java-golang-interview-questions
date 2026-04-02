package main

import "fmt"

// 70. Climbing Stairs
// Time: O(N), Space: O(1)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	
	prev2, prev1 := 1, 2
	
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2, prev1 = prev1, current
	}
	
	return prev1
}

// DP array approach
func climbStairsDP(n int) int {
	if n <= 2 {
		return n
	}
	
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	
	return dp[n]
}

// Recursive with memoization
func climbStairsMemo(n int) int {
	memo := make(map[int]int)
	return climbStairsHelper(n, memo)
}

func climbStairsHelper(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}
	
	if val, exists := memo[n]; exists {
		return val
	}
	
	result := climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
	memo[n] = result
	return result
}

func main() {
	// Test cases
	testCases := []int{
		1, 2, 3, 4, 5, 10, 20, 30, 45, 50,
	}
	
	for i, n := range testCases {
		result1 := climbStairs(n)
		result2 := climbStairsDP(n)
		result3 := climbStairsMemo(n)
		
		fmt.Printf("Test Case %d: n=%d -> Iterative: %d, DP: %d, Memo: %d\n", 
			i+1, n, result1, result2, result3)
	}
}
