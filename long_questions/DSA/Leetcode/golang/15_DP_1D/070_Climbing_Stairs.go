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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Fibonacci-style DP
- **Linear DP**: Each state depends on previous two states
- **Bottom-up Iteration**: Build solution from base cases
- **Space Optimization**: Use constant space instead of DP array
- **Recurrence Relation**: dp[i] = dp[i-1] + dp[i-2]

## 2. PROBLEM CHARACTERISTICS
- **Linear Sequence**: Ways to reach step n depends on previous steps
- **Fibonacci Pattern**: Same recurrence as Fibonacci numbers
- **Base Cases**: Step 1 = 1 way, Step 2 = 2 ways
- **Unbounded Growth**: Number of ways grows exponentially

## 3. SIMILAR PROBLEMS
- Fibonacci Numbers (Classic) - Same recurrence relation
- House Robber (LeetCode 198) - Similar DP pattern
- Decode Ways (LeetCode 91) - Similar recurrence
- Tribonacci (Extension) - dp[i] = dp[i-1] + dp[i-2] + dp[i-3]

## 4. KEY OBSERVATIONS
- **Two-step constraint**: Can climb 1 or 2 steps at a time
- **Last step analysis**: Last step could be from n-1 or n-2
- **Fibonacci relationship**: f(n) = f(n-1) + f(n-2)
- **Space optimization**: Only need last two values

## 5. VARIATIONS & EXTENSIONS
- **Different step sizes**: Can climb 1, 2, 3 steps
- **Staircase with obstacles**: Some steps cannot be used
- **Cost per step**: Minimize cost to reach top
- **Count paths modulo**: Return result modulo large number

## 6. INTERVIEW INSIGHTS
- Always clarify: "What are the step options? Are there constraints?"
- Edge cases: n=0 (0 ways), n=1 (1 way), n=2 (2 ways)
- Time complexity: O(N) for iterative, O(N) for DP array
- Space complexity: O(1) optimized, O(N) for DP array

## 7. COMMON MISTAKES
- Not handling base cases properly (n=1, n=2)
- Using O(N) space when O(1) is sufficient
- Off-by-one errors in iteration bounds
- Not considering n=0 case
- Using recursion without memoization (exponential time)

## 8. OPTIMIZATION STRATEGIES
- **Space optimization**: Use only two variables instead of array
- **Matrix exponentiation**: O(log N) time for very large N
- **Closed-form solution**: Use Fibonacci formula (Binet's formula)
- **Modulo optimization**: Apply modulo during computation

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like climbing stairs with two options:**
- You're at the bottom of a staircase with N steps
- At each step, you can take either 1 step or 2 steps
- You want to know how many different ways to reach the top
- The last step you take could be from step N-1 (1-step) or N-2 (2-step)
- This creates a Fibonacci-like pattern

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Integer N (number of stairs)
2. **Goal**: Count distinct ways to climb to the top
3. **Output**: Number of ways (integer)
4. **Constraint**: Can take 1 or 2 steps at a time

#### Phase 2: Key Insight Recognition
- **"Last step analysis"** → Final step determines recurrence
- **"Fibonacci pattern"** → Same recurrence as Fibonacci
- **"DP natural fit"** → Optimal substructure and overlapping subproblems
- **"Space optimization"** → Only need last two values

#### Phase 3: Strategy Development
```
Human thought process:
"I need to count ways to climb N stairs.
The last step could be from N-1 (1-step) or N-2 (2-step).
So ways(N) = ways(N-1) + ways(N-2).
This is exactly the Fibonacci sequence!
I can build this bottom-up using DP.
I only need the last two values, so I can optimize space."
```

#### Phase 4: Edge Case Handling
- **N=0**: 0 ways (no stairs to climb)
- **N=1**: 1 way (single 1-step)
- **N=2**: 2 ways (1+1 or 2)
- **Large N**: Handle integer overflow if needed

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
N=4 stairs:

Human thinking:
"I'll build from bottom up:

Base cases:
ways(0) = 0 (no stairs)
ways(1) = 1 (just 1-step)
ways(2) = 2 (1+1 or 2)

Now build up:
ways(3) = ways(2) + ways(1) = 2 + 1 = 3
- Ways: 1+1+1, 1+2, 2+1

ways(4) = ways(3) + ways(2) = 3 + 2 = 5
- Ways: 1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2

So there are 5 ways to climb 4 stairs."
```

#### Phase 6: Intuition Validation
- **Why DP works**: Optimal substructure and overlapping subproblems
- **Why Fibonacci relationship**: Same recurrence as Fibonacci numbers
- **Why O(N) time**: Need to compute each value once
- **Why O(1) space**: Only need last two values

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Exponential time without memoization
2. **"Should I use DP array?"** → Can optimize to O(1) space
3. **"What about very large N?"** → Consider matrix exponentiation
4. **"Can I find a formula?"** → Yes, Binet's formula but has precision issues

### Real-World Analogy
**Like counting different ways to pay with coins:**
- You need to make N cents using 1-cent and 2-cent coins
- The last coin used could be 1-cent (leaving N-1 cents) or 2-cent (leaving N-2 cents)
- Count ways(N) = ways(N-1) + ways(N-2)
- This creates the same Fibonacci pattern

### Human-Readable Pseudocode
```
function climbStairs(n):
    if n <= 2:
        return n
    
    prev2 = 1  // ways(1)
    prev1 = 2  // ways(2)
    
    for i from 3 to n:
        current = prev1 + prev2
        prev2 = prev1
        prev1 = current
    
    return prev1
```

### Execution Visualization

### Example: N=5
```
DP Evolution:
ways[0] = 0
ways[1] = 1
ways[2] = 2
ways[3] = ways[2] + ways[1] = 2 + 1 = 3
ways[4] = ways[3] + ways[2] = 3 + 2 = 5
ways[5] = ways[4] + ways[3] = 5 + 3 = 8

Space-optimized evolution:
prev2=1, prev1=2 (ways[1], ways[2])
i=3: current=3, prev2=2, prev1=3
i=4: current=5, prev2=3, prev1=5
i=5: current=8, prev2=5, prev1=8

Result: 8 ways
```

### Key Visualization Points:
- **Fibonacci pattern**: Each value is sum of previous two
- **Space optimization**: Only track last two values
- **Bottom-up building**: Start from base cases and build up
- **Linear progression**: Each step depends only on immediate predecessors

### Memory Layout Visualization:
```
DP Array Approach:
[0, 1, 2, 3, 5, 8, ...]
 ^  ^  ^  ^  ^  ^
 |  |  |  |  |  |
0  1  2  3  4  5 (step numbers)

Space-Optimized Approach:
Step 1: prev2=1, prev1=2
Step 2: prev2=2, prev1=3
Step 3: prev2=3, prev1=5
Step 4: prev2=5, prev1=8
```

### Time Complexity Breakdown:
- **DP array**: O(N) time, O(N) space
- **Space optimized**: O(N) time, O(1) space
- **Recursive without memo**: O(2^N) time, O(N) space
- **Recursive with memo**: O(N) time, O(N) space

### Alternative Approaches:

#### 1. Matrix Exponentiation (O(log N) time, O(1) space)
```go
func climbStairsMatrix(n int) int {
    if n <= 2 {
        return n
    }
    
    // Using matrix exponentiation for Fibonacci
    // [f(n)]   [1 1]^(n-2) * [f(2)]
    // [f(n-1)] [1 0]         [f(1)]
    
    result := matrixPower([][]int{{1, 1}, {1, 0}}, n-2)
    return result[0][0]*2 + result[0][1]*1
}

func matrixPower(matrix [][]int, power int) [][]int {
    if power == 0 {
        return [][]int{{1, 0}, {0, 1}} // Identity matrix
    }
    
    if power == 1 {
        return matrix
    }
    
    half := matrixPower(matrix, power/2)
    result := multiplyMatrices(half, half)
    
    if power%2 == 1 {
        result = multiplyMatrices(result, matrix)
    }
    
    return result
}

func multiplyMatrices(a, b [][]int) [][]int {
    result := make([][]int, 2)
    for i := 0; i < 2; i++ {
        result[i] = make([]int, 2)
        for j := 0; j < 2; j++ {
            result[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
        }
    }
    return result
}
```
- **Pros**: O(log N) time for very large N
- **Cons**: Complex implementation, overhead for small N

#### 2. Closed-form Formula (O(1) time, O(1) space)
```go
import "math"

func climbStairsFormula(n int) int {
    // Binet's formula: F(n) = (φ^n - ψ^n) / √5
    // where φ = (1 + √5) / 2, ψ = (1 - √5) / 2
    // For climbing stairs: result = F(n+1)
    
    sqrt5 := math.Sqrt(5)
    phi := (1 + sqrt5) / 2
    psi := (1 - sqrt5) / 2
    
    // Round to nearest integer to handle floating point precision
    result := math.Round((math.Pow(phi, float64(n+1)) - math.Pow(psi, float64(n+1))) / sqrt5)
    return int(result)
}
```
- **Pros**: O(1) time
- **Cons**: Floating point precision issues for large N

#### 3. Recursive with Memoization (O(N) time, O(N) space)
```go
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
```
- **Pros**: Intuitive recursive approach
- **Cons**: More memory than iterative approach

### Extensions for Interviews:
- **Different step sizes**: Can climb 1, 2, 3 steps (tribonacci)
- **Staircase with obstacles**: Some steps cannot be used
- **Cost per step**: Minimize total cost to reach top
- **Modulo arithmetic**: Return result % 10^9+7 for large numbers
- **Path reconstruction**: Also return the actual paths
*/
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
