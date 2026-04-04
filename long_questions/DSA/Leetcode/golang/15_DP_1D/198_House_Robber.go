package main

import "fmt"

// 198. House Robber
// Time: O(N), Space: O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	
	prev2, prev1 := nums[0], max(nums[0], nums[1])
	
	for i := 2; i < len(nums); i++ {
		current := max(prev1, prev2+nums[i])
		prev2, prev1 = prev1, current
	}
	
	return prev1
}

// DP array approach
func robDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	
	return dp[len(nums)-1]
}

// Recursive with memoization
func robMemo(nums []int) int {
	memo := make(map[int]int)
	return robHelper(nums, len(nums)-1, memo)
}

func robHelper(nums []int, index int, memo map[int]int) int {
	if index < 0 {
		return 0
	}
	if index == 0 {
		return nums[0]
	}
	
	if val, exists := memo[index]; exists {
		return val
	}
	
	result := max(robHelper(nums, index-1, memo), robHelper(nums, index-2, memo)+nums[index])
	memo[index] = result
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: House Robber DP
- **Choice-based DP**: At each house, choose to rob or skip
- **Linear DP**: Each state depends on previous two states
- **Space Optimization**: Use constant space instead of DP array
- **Recurrence Relation**: dp[i] = max(dp[i-1], dp[i-2] + nums[i])

## 2. PROBLEM CHARACTERISTICS
- **Adjacency Constraint**: Cannot rob adjacent houses
- **Maximum Sum**: Find maximum amount that can be robbed
- **Linear Arrangement**: Houses are in a straight line
- **Independent Choices**: Each house decision affects next

## 3. SIMILAR PROBLEMS
- House Robber II (LeetCode 213) - Circular houses
- House Robber III (LeetCode 337) - Binary tree houses
- Maximum Sum of Non-adjacent Elements (Classic DP)
- Delete and Earn (LeetCode 740) - Similar pattern

## 4. KEY OBSERVATIONS
- **Local decision**: For each house, rob or skip
- **Global optimum**: Local decisions lead to global optimum
- **Two-state dependency**: Current depends on previous two states
- **Space optimization**: Only need last two values

## 5. VARIATIONS & EXTENSIONS
- **Circular houses**: First and last houses are adjacent
- **Tree houses**: Houses arranged in binary tree
- **Multiple thieves**: More complex constraints
- **Time windows**: Houses available at specific times

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are houses in a line or circle? What about empty input?"
- Edge cases: empty array (0), single house (value), two houses (max of two)
- Time complexity: O(N) for iterative, O(N) for DP array
- Space complexity: O(1) optimized, O(N) for DP array

## 7. COMMON MISTAKES
- Not handling empty array case
- Using O(N) space when O(1) is sufficient
- Off-by-one errors in DP array indexing
- Not considering single house case
- Using greedy approach (fails for many cases)

## 8. OPTIMIZATION STRATEGIES
- **Space optimization**: Use only two variables instead of array
- **Early termination**: Not applicable (need to process all houses)
- **Input validation**: Handle edge cases efficiently
- **Memory optimization**: Reuse variables when possible

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like planning a heist with constraints:**
- You're a thief planning to rob houses on a street
- You can't rob two adjacent houses (alarms will trigger)
- Each house has a different amount of money
- You want to maximize your total haul
- For each house, you decide: rob it (and skip previous) or skip it

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of house values
2. **Goal**: Find maximum amount that can be robbed
3. **Constraint**: Cannot rob adjacent houses
4. **Output**: Maximum amount (integer)

#### Phase 2: Key Insight Recognition
- **"Local decision"** → For each house: rob or skip
- **"Two-state dependency"** → Current decision depends on previous two
- **"DP natural fit"** → Optimal substructure and overlapping subproblems
- **"Space optimization"** → Only need last two values

#### Phase 3: Strategy Development
```
Human thought process:
"I need to maximize money from non-adjacent houses.
For each house, I have two choices:
1. Rob this house: then I can't rob the previous one
2. Skip this house: then I take whatever I got from previous
So dp[i] = max(rob current, skip current)
= max(dp[i-2] + nums[i], dp[i-1])
I can build this bottom-up using DP."
```

#### Phase 4: Edge Case Handling
- **Empty array**: 0 money
- **One house**: Rob that house
- **Two houses**: Rob the richer one
- **Large arrays**: Handle efficiently with O(1) space

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Houses: [2, 7, 9, 3, 1]

Human thinking:
"I'll go house by house:

House 0 (value 2):
- Must rob it (only option)
- Max so far: 2

House 1 (value 7):
- Rob house 1: 7 (skip house 0)
- Skip house 1: 2 (keep previous)
- Choose max: 7
- Max so far: 7

House 2 (value 9):
- Rob house 2: 9 + dp[0] = 9 + 2 = 11
- Skip house 2: dp[1] = 7
- Choose max: 11
- Max so far: 11

House 3 (value 3):
- Rob house 3: 3 + dp[1] = 3 + 7 = 10
- Skip house 3: dp[2] = 11
- Choose max: 11
- Max so far: 11

House 4 (value 1):
- Rob house 4: 1 + dp[2] = 1 + 11 = 12
- Skip house 4: dp[3] = 11
- Choose max: 12
- Final answer: 12

Rob houses 0, 2, 4: 2 + 9 + 1 = 12"
```

#### Phase 6: Intuition Validation
- **Why DP works**: Optimal substructure and overlapping subproblems
- **Why two-state dependency**: Robbing current affects previous choice
- **Why O(N) time**: Need to process each house once
- **Why O(1) space**: Only need last two maximums

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use greedy?"** → Greedy fails (e.g., [2,1,1,2])
2. **"Should I use recursion?"** → Can, but iterative is more efficient
3. **"What about circular houses?"** → Different problem (House Robber II)
4. **"Can I track which houses?"** → Yes, but requires additional space

### Real-World Analogy
**Like scheduling tasks with conflicts:**
- You have tasks with different values/profits
- Some tasks conflict and can't be scheduled together
- You want to maximize total profit
- For each task, decide to take it (and skip conflicting) or skip it
- This creates the same DP pattern

### Human-Readable Pseudocode
```
function rob(nums):
    if len(nums) == 0:
        return 0
    if len(nums) == 1:
        return nums[0]
    
    prev2 = nums[0]           // dp[0]
    prev1 = max(nums[0], nums[1])  // dp[1]
    
    for i from 2 to len(nums)-1:
        current = max(prev1, prev2 + nums[i])
        prev2 = prev1
        prev1 = current
    
    return prev1
```

### Execution Visualization

### Example: nums = [2, 7, 9, 3, 1]
```
DP Evolution:
dp[0] = 2
dp[1] = max(2, 7) = 7
dp[2] = max(dp[1], dp[0] + 9) = max(7, 2 + 9) = 11
dp[3] = max(dp[2], dp[1] + 3) = max(11, 7 + 3) = 11
dp[4] = max(dp[3], dp[2] + 1) = max(11, 11 + 1) = 12

Space-optimized evolution:
prev2=2, prev1=7 (dp[0], dp[1])
i=2: current=11, prev2=7, prev1=11
i=3: current=11, prev2=11, prev1=11
i=4: current=12, prev2=11, prev1=12

Result: 12
```

### Key Visualization Points:
- **Choice at each step**: Rob current or skip current
- **Two-state dependency**: Current depends on previous two states
- **Space optimization**: Only track last two maximums
- **Bottom-up building**: Start from base cases and build up

### Memory Layout Visualization:
```
DP Array Approach:
[2, 7, 11, 11, 12]
 ^  ^   ^   ^   ^
 |  |   |   |   |
0  1   2   3   4 (house indices)

Space-Optimized Approach:
House 0: prev2=2, prev1=7
House 1: prev2=7, prev1=11
House 2: prev2=11, prev1=11
House 3: prev2=11, prev1=12

Decision at each house:
House 0: Rob (only option)
House 1: Rob (7 > 2)
House 2: Rob (11 > 7)
House 3: Skip (11 > 10)
House 4: Rob (12 > 11)
```

### Time Complexity Breakdown:
- **DP array**: O(N) time, O(N) space
- **Space optimized**: O(N) time, O(1) space
- **Recursive without memo**: O(2^N) time, O(N) space
- **Recursive with memo**: O(N) time, O(N) space

### Alternative Approaches:

#### 1. Recursive with Memoization (O(N) time, O(N) space)
```go
func robMemo(nums []int) int {
    memo := make(map[int]int)
    return robHelper(nums, len(nums)-1, memo)
}

func robHelper(nums []int, index int, memo map[int]int) int {
    if index < 0 {
        return 0
    }
    if index == 0 {
        return nums[0]
    }
    
    if val, exists := memo[index]; exists {
        return val
    }
    
    result := max(robHelper(nums, index-1, memo), 
                 robHelper(nums, index-2, memo)+nums[index])
    memo[index] = result
    return result
}
```
- **Pros**: Intuitive recursive approach
- **Cons**: More memory than iterative approach

#### 2. Track Houses (O(N) time, O(N) space)
```go
func robWithHouses(nums []int) (int, []int) {
    if len(nums) == 0 {
        return 0, []int{}
    }
    if len(nums) == 1 {
        return nums[0], []int{0}
    }
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])
    
    for i := 2; i < len(nums); i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
    }
    
    // Reconstruct which houses to rob
    houses := []int{}
    i := len(nums) - 1
    for i >= 0 {
        if i == 0 || dp[i] != dp[i-1] {
            houses = append([]int{i}, houses...)
            i -= 2
        } else {
            i -= 1
        }
    }
    
    return dp[len(nums)-1], houses
}
```
- **Pros**: Returns actual houses to rob
- **Cons**: Uses O(N) space for reconstruction

#### 3. Greedy with Two Pointers (Incorrect - for demonstration)
```go
func robGreedy(nums []int) int {
    // This is INCORRECT but shows why greedy fails
    evenSum, oddSum := 0, 0
    
    for i, val := range nums {
        if i%2 == 0 {
            evenSum += val
        } else {
            oddSum += val
        }
    }
    
    return max(evenSum, oddSum)
}
```
- **Pros**: Simple to implement
- **Cons**: Incorrect for many cases (e.g., [2,1,1,2])

### Extensions for Interviews:
- **House Robber II**: Circular houses (first and last adjacent)
- **House Robber III**: Houses in binary tree structure
- **Multiple Thieves**: More complex constraints
- **Time Windows**: Houses available at specific times
- **Path Reconstruction**: Also return which houses to rob
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
		{2, 1, 1, 2},
		{1},
		{1, 2},
		{5, 5, 10, 100, 10, 5},
		{2, 7, 9, 3, 1, 5, 6, 8},
		{100, 1, 1, 100},
		{4, 1, 2, 7, 5, 3, 1},
	}
	
	for i, nums := range testCases {
		result1 := rob(nums)
		result2 := robDP(nums)
		result3 := robMemo(nums)
		
		fmt.Printf("Test Case %d: %v -> Iterative: %d, DP: %d, Memo: %d\n", 
			i+1, nums, result1, result2, result3)
	}
}
