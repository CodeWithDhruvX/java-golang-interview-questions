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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Unbounded Knapsack DP
- **Bottom-up DP**: Build solution from smallest amount to target
- **Coin Combination**: Each coin can be used unlimited times
- **Minimum Optimization**: Find minimum number of coins
- **Unbounded Knapsack**: Classic unbounded knapsack variant

## 2. PROBLEM CHARACTERISTICS
- **Unlimited Coins**: Each coin denomination can be used multiple times
- **Minimum Count**: Find minimum number of coins, not combinations
- **Exact Amount**: Must make exact target amount
- **No Solution Case**: Return -1 if amount cannot be made

## 3. SIMILAR PROBLEMS
- Coin Change II (LeetCode 518) - Count combinations
- Combination Sum (LeetCode 39) - Find actual combinations
- Minimum Cost to Fill Bag (Classic) - Similar optimization
- Partition Equal Subset Sum (LeetCode 416) - Subset sum variant

## 4. KEY OBSERVATIONS
- **DP natural fit**: Optimal substructure and overlapping subproblems
- **Bottom-up approach**: Build from 0 to target amount
- **Unlimited usage**: Each coin can be used multiple times
- **Minimization**: Track minimum coins for each amount

## 5. VARIATIONS & EXTENSIONS
- **Count Combinations**: Number of ways to make amount
- **Limited Coins**: Each coin has limited quantity
- **Coin Tracking**: Also return which coins used
- **Multiple Queries**: Answer queries for multiple amounts

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are coins unlimited? What about no solution case?"
- Edge cases: amount=0 (0 coins), no solution (-1)
- Time complexity: O(S × N) where S is amount, N is number of coins
- Space complexity: O(S) for DP array

## 7. COMMON MISTAKES
- Not initializing DP array with infinity/large value
- Not handling no solution case properly
- Using wrong DP order (coins vs amount)
- Not optimizing coin order for early termination
- Forgetting amount=0 base case

## 8. OPTIMIZATION STRATEGIES
- **Coin ordering**: Sort coins for potential early termination
- **Early pruning**: Skip impossible amounts
- **Space optimization**: Use 1D array instead of 2D
- **Input validation**: Handle edge cases efficiently

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like making change with unlimited coins:**
- You have unlimited coins of different denominations
- You need to make exact change for a target amount
- You want to use the minimum number of coins possible
- For each amount, you try each coin and see if it helps
- Build up from small amounts to the target amount

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of coin denominations, target amount
2. **Goal**: Find minimum number of coins to make exact amount
3. **Constraint**: Unlimited coins of each denomination
4. **Output**: Minimum coins count, or -1 if impossible

#### Phase 2: Key Insight Recognition
- **"Bottom-up DP"** → Build solution from 0 to target amount
- **"Unbounded knapsack"** → Each coin can be used multiple times
- **"Minimization"** → Track minimum coins for each amount
- **"Subproblem optimal"** → Optimal solution for amount depends on smaller amounts

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find minimum coins to make the target amount.
For each amount from 0 to target, I'll try each coin.
If I can use coin c, then dp[amount] = min(dp[amount], dp[amount-c] + 1).
I'll build this bottom-up from 0 to target.
If dp[target] is still infinity, no solution exists."
```

#### Phase 4: Edge Case Handling
- **Amount = 0**: 0 coins needed
- **No solution**: Return -1 if dp[target] remains infinity
- **Single coin**: If target equals coin value, return 1
- **Large amounts**: Handle efficiently with DP

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Coins: [1, 2, 5], Amount: 11

Human thinking:
"I'll build DP from 0 to 11:

dp[0] = 0 (base case)
dp[1..11] = infinity initially

Amount 1:
- Try coin 1: dp[1] = min(inf, dp[0] + 1) = 1
- Try coin 2: too big
- Try coin 5: too big
dp[1] = 1

Amount 2:
- Try coin 1: dp[2] = min(inf, dp[1] + 1) = 2
- Try coin 2: dp[2] = min(2, dp[0] + 1) = 1
- Try coin 5: too big
dp[2] = 1

Amount 3:
- Try coin 1: dp[3] = min(inf, dp[2] + 1) = 2
- Try coin 2: dp[3] = min(2, dp[1] + 1) = 2
- Try coin 5: too big
dp[3] = 2

Continue this process...
Amount 11:
- Try coin 1: dp[11] = min(inf, dp[10] + 1) = 4
- Try coin 2: dp[11] = min(4, dp[9] + 1) = 4
- Try coin 5: dp[11] = min(4, dp[6] + 1) = 3

Final answer: 3 coins (5 + 5 + 1)"
```

#### Phase 6: Intuition Validation
- **Why DP works**: Optimal substructure and overlapping subproblems
- **Why bottom-up**: Need results from smaller amounts first
- **Why O(S × N)**: S amounts, N coins per amount
- **Why unbounded**: Can reuse same coin multiple times

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use greedy?"** → Greedy fails for many coin systems
2. **"Should I use recursion?"** → Can, but iterative is more efficient
3. **"What about coin order?"** → Order doesn't matter for correctness
4. **"Can I track coins used?"** → Yes, but requires additional space

### Real-World Analogy
**Like making change at a cash register:**
- You have bills/coins of different denominations
- Customer gives you an amount, you need to give exact change
- You want to use the fewest bills/coins possible
- For each amount, you try different combinations
- You build up from smaller amounts to larger ones

### Human-Readable Pseudocode
```
function coinChange(coins, amount):
    if amount == 0:
        return 0
    
    dp = array of size amount+1, filled with infinity
    dp[0] = 0
    
    for i from 1 to amount:
        for coin in coins:
            if coin <= i:
                dp[i] = min(dp[i], dp[i-coin] + 1)
    
    if dp[amount] == infinity:
        return -1
    else:
        return dp[amount]
```

### Execution Visualization

### Example: coins = [1, 2, 5], amount = 11
```
DP Array Evolution:
dp[0] = 0
dp[1..11] = [∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞]

Amount 1:
- coin 1: dp[1] = min(∞, dp[0] + 1) = 1
dp = [0, 1, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞]

Amount 2:
- coin 1: dp[2] = min(∞, dp[1] + 1) = 2
- coin 2: dp[2] = min(2, dp[0] + 1) = 1
dp = [0, 1, 1, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞]

Amount 3:
- coin 1: dp[3] = min(∞, dp[2] + 1) = 2
- coin 2: dp[3] = min(2, dp[1] + 1) = 2
dp = [0, 1, 1, 2, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞]

Continue until amount 11:
Final dp = [0, 1, 1, 2, 2, 1, 2, 2, 3, 3, 4, 3]

Answer: dp[11] = 3
```

### Key Visualization Points:
- **Bottom-up building**: Start from 0 and build to target
- **Minimization**: Always keep minimum coins for each amount
- **Unlimited usage**: Can reuse same coin multiple times
- **Infinity handling**: Use large value for impossible amounts

### Memory Layout Visualization:
```
DP Table Structure:
Amount:   0  1  2  3  4  5  6  7  8  9  10  11
Coins:   1  2  5
dp:      0  1  1  2  2  1  2  2  3  3  4   3

Optimal Solutions:
0: []
1: [1]
2: [2]
3: [2,1]
4: [2,2]
5: [5]
6: [5,1]
7: [5,2]
8: [5,2,1]
9: [5,2,2]
10: [5,5]
11: [5,5,1]
```

### Time Complexity Breakdown:
- **Standard DP**: O(S × N) time, O(S) space
- **Optimized DP**: O(S × N) time, O(S) space
- **Recursive with memo**: O(S × N) time, O(S) space
- **Greedy**: O(N) time but incorrect for many cases

### Alternative Approaches:

#### 1. Recursive with Memoization (O(S × N) time, O(S) space)
```go
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
```
- **Pros**: Intuitive recursive approach
- **Cons**: More overhead than iterative DP

#### 2. BFS Approach (O(S × N) time, O(S) space)
```go
func coinChangeBFS(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    
    queue := []int{0}
    visited := make(map[int]bool)
    visited[0] = true
    level := 0
    
    for len(queue) > 0 {
        levelSize := len(queue)
        
        for i := 0; i < levelSize; i++ {
            current := queue[0]
            queue = queue[1:]
            
            for _, coin := range coins {
                next := current + coin
                if next == amount {
                    return level + 1
                }
                if next < amount && !visited[next] {
                    visited[next] = true
                    queue = append(queue, next)
                }
            }
        }
        level++
    }
    
    return -1
}
```
- **Pros**: Finds shortest path in terms of coins
- **Cons**: More memory for visited set

#### 3. Track Coins Used (O(S × N) time, O(S) space)
```go
func coinChangeWithCoins(coins []int, amount int) (int, []int) {
    if amount == 0 {
        return 0, []int{}
    }
    
    dp := make([]int, amount+1)
    parent := make([]int, amount+1)
    
    for i := 1; i <= amount; i++ {
        dp[i] = math.MaxInt32
        parent[i] = -1
    }
    
    for i := 1; i <= amount; i++ {
        for _, coin := range coins {
            if coin <= i && dp[i-coin] != math.MaxInt32 {
                if dp[i-coin]+1 < dp[i] {
                    dp[i] = dp[i-coin] + 1
                    parent[i] = i - coin
                }
            }
        }
    }
    
    if dp[amount] == math.MaxInt32 {
        return -1, []int{}
    }
    
    // Reconstruct coins used
    coinsUsed := []int{}
    current := amount
    for current > 0 {
        coinsUsed = append(coinsUsed, current-parent[current])
        current = parent[current]
    }
    
    return dp[amount], coinsUsed
}
```
- **Pros**: Returns actual coins used
- **Cons**: Additional space for reconstruction

### Extensions for Interviews:
- **Coin Change II**: Count number of ways to make amount
- **Limited Coins**: Each coin has limited quantity
- **Multiple Queries**: Answer multiple amount queries efficiently
- **Coin Tracking**: Also return which coins were used
- **Large Amounts**: Handle very large amounts efficiently
*/
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
