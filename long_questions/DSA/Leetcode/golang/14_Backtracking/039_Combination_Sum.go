package main

import (
	"fmt"
	"sort"
)

// 39. Combination Sum
// Time: O(2^T/M), Space: O(T/M) where T is target, M is minimum candidate
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sort to help with pruning
	var result [][]int
	current := make([]int, 0, len(candidates))
	
	backtrackCombinationSum(candidates, target, 0, current, &result)
	return result
}

func backtrackCombinationSum(candidates []int, target, start int, current []int, result *[][]int) {
	if target == 0 {
		// Found a valid combination
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	
	if target < 0 {
		return // Exceeded target
	}
	
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break // No need to continue as candidates are sorted
		}
		
		// Include candidates[i]
		current = append(current, candidates[i])
		backtrackCombinationSum(candidates, target-candidates[i], i, current, result) // i (not i+1) because we can reuse
		current = current[:len(current)-1] // Backtrack
	}
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Backtracking with Pruning
- **Recursive Exploration**: Try all combinations recursively
- **Target Sum Tracking**: Maintain remaining target to reach
- **Pruning Strategy**: Skip candidates that exceed remaining target
- **Unlimited Reuse**: Candidates can be used multiple times

## 2. PROBLEM CHARACTERISTICS
- **Combination Problem**: Find all combinations that sum to target
- **Unlimited Usage**: Each candidate can be used unlimited times
- **No Duplicates**: Input candidates are unique
- **Order Independence**: [2,3] and [3,2] are same combination

## 3. SIMILAR PROBLEMS
- Combination Sum II (LeetCode 40) - Each candidate used once
- Combination Sum III (LeetCode 216) - Fixed size combinations
- Combination Sum IV (LeetCode 377) - Count combinations
- Partition Equal Subset Sum (LeetCode 416)

## 4. KEY OBSERVATIONS
- **Backtracking natural fit**: Need to explore all combinations
- **Sorting helps**: Enables early pruning when candidate > remaining target
- **Reuse parameter**: Pass same index (not i+1) to allow unlimited reuse
- **Target reduction**: Reduce target by candidate value during recursion

## 5. VARIATIONS & EXTENSIONS
- **Limited Usage**: Each candidate used at most once
- **Fixed Size**: Combinations of exactly k numbers
- **Count Only**: Return count instead of actual combinations
- **Multiple Targets**: Find combinations for multiple target values

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can numbers be reused? Are there duplicates?"
- Edge cases: empty candidates, target 0, no solution
- Time complexity: O(2^(T/M)) where T is target, M is minimum candidate
- Space complexity: O(T/M) for recursion depth

## 7. COMMON MISTAKES
- Not sorting candidates (misses pruning optimization)
- Using i+1 instead of i (prevents unlimited reuse)
- Not handling target < 0 base case properly
- Making shallow copies of current combination
- Not pruning when candidate > remaining target

## 8. OPTIMIZATION STRATEGIES
- **Sorting candidates**: Enables early pruning
- **Pruning condition**: Break when candidate > remaining target
- **Early exit**: Stop when target reaches 0
- **Memory optimization**: Reuse slices instead of creating new ones

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like making change with unlimited coins:**
- You have coins of different denominations (candidates)
- You need to make exact change for a target amount
- You can use each coin type unlimited times
- Find all possible ways to make the exact amount
- Order doesn't matter (using 2 then 3 is same as 3 then 2)

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of candidate numbers, target sum
2. **Goal**: Find all unique combinations that sum to target
3. **Output**: List of combinations, each combination sums to target
4. **Constraint**: Each candidate can be used unlimited times

#### Phase 2: Key Insight Recognition
- **"Backtracking natural fit"** → Need to explore all possible combinations
- **"Unlimited reuse"** → Can use same candidate multiple times in one combination
- **"Pruning opportunity"** → Stop exploring when sum exceeds target
- **"Order independence"** → Need to avoid duplicate combinations

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find all combinations that sum to the target.
I'll use backtracking to explore possibilities recursively.
For each candidate, I can either include it or skip it.
If I include it, I reduce the target and can use it again.
If I skip it, I move to the next candidate.
I'll sort candidates to enable early pruning."
```

#### Phase 4: Edge Case Handling
- **Empty candidates**: Return empty result
- **Target 0**: Return empty combination (valid solution)
- **No solution**: Return empty result after exploring all possibilities
- **Large target**: Handle recursion depth appropriately

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Candidates: [2,3,6,7], Target: 7

Human thinking:
"I'll sort candidates: [2,3,6,7]

Start with empty combination, target=7:

Try 2:
- Include 2: current=[2], target=5
  - Can use 2 again: include 2: current=[2,2], target=3
    - Can use 2 again: include 2: current=[2,2,2], target=1 (too small, skip)
    - Try 3: include 3: current=[2,2,3], target=0 → FOUND [2,2,3]
    - Try 6: too big, skip
    - Try 7: too big, skip
  - Try 3: include 3: current=[2,3], target=2
    - Try 2: include 2: current=[2,3,2], target=0 → FOUND [2,3,2] (same as [2,2,3])
  - Try 6: too big, skip
  - Try 7: include 7: current=[2,7], target=-2 (too big, backtrack)

Try 3:
- Include 3: current=[3], target=4
  - Try 3: include 3: current=[3,3], target=1 (too small)
  - Try 6: too big, skip
  - Try 7: too big, skip

Try 6:
- Include 6: current=[6], target=1 (too small)

Try 7:
- Include 7: current=[7], target=0 → FOUND [7]

Final result: [[2,2,3], [7]]"
```

#### Phase 6: Intuition Validation
- **Why backtracking works**: Need to explore all combination possibilities
- **Why unlimited reuse works**: Pass same index to allow reuse
- **Why pruning works**: Stop when candidate > remaining target
- **Why sorting helps**: Enables early pruning and consistent ordering

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use greedy?"** → Greedy fails for many cases (e.g., [2,3,6,7], target=7)
2. **"Should I use DP?"** → Could for counting, but need actual combinations
3. **"What about very large targets?"** → Pruning helps but complexity is exponential
4. **"Can I optimize further?"** → Sorting and pruning are already optimal

### Real-World Analogy
**Like finding all ways to make change for a dollar:**
- You have coins of different denominations (pennies, nickels, dimes, quarters)
- You need to make exactly one dollar
- You can use each coin type as many times as needed
- Find all possible combinations that make exactly one dollar
- Order doesn't matter (penny+nickel same as nickel+penny)

### Human-Readable Pseudocode
```
function combinationSum(candidates, target):
    sort(candidates)
    result = []
    current = []
    
    backtrack(candidates, target, 0, current, result)
    return result

function backtrack(candidates, target, start, current, result):
    if target == 0:
        add copy of current to result
        return
    
    if target < 0:
        return  // Exceeded target
    
    for i from start to length(candidates)-1:
        if candidates[i] > target:
            break  // Prune
        
        current.append(candidates[i])
        backtrack(candidates, target - candidates[i], i, current, result)  // i allows reuse
        current.pop()  // Backtrack
```

### Execution Visualization

### Example: candidates=[2,3,6,7], target=7
```
Recursion Tree:
                    [], target=7
                   /    |    \
                [2]     [3]   [6]   [7]
                t=5     t=4    t=1    t=0 ✓
               / | \    / |           
            [2,2] [2,3] [2,6] [3,3] [3,6]
            t=3   t=0 ✓  t=-1   t=1   t=-2
           / |                    
        [2,2,2] [2,2,3]
        t=1     t=0 ✓
        (prune) 

Valid combinations found: [2,2,3], [7]
```

### Key Visualization Points:
- **Recursive exploration**: Each level represents including a candidate
- **Target reduction**: Target decreases as we include candidates
- **Pruning**: Stop when candidate > remaining target
- **Reuse allowed**: Same index passed to recursive call
- **Backtracking**: Remove last candidate when exploring different path

### Memory Layout Visualization:
```
Call Stack Evolution:
backtrack([],7,0) → backtrack([2],5,0) → backtrack([2,2],3,0) → backtrack([2,2,3],0,0) ✓
                ↘ backtrack([2,3],2,0) → backtrack([2,3,2],0,0) ✓
                ↘ backtrack([3],4,0)
                ↘ backtrack([6],1,0) (prune)
                ↘ backtrack([7],0,0) ✓

Current combinations at each level:
[] → [2] → [2,2] → [2,2,3] ✓
          → [2,3] → [2,3,2] ✓
     → [3]
     → [6] (prune)
     → [7] ✓
```

### Time Complexity Breakdown:
- **Worst case**: O(2^(T/M)) where T is target, M is minimum candidate
- **Pruning effect**: Reduces actual explored paths significantly
- **Space complexity**: O(T/M) for maximum recursion depth
- **Result storage**: O(K × L) where K is number of combinations, L is average length

### Alternative Approaches:

#### 1. Dynamic Programming (O(T × N) time, O(T) space)
```go
func combinationSumDP(candidates []int, target int) [][]int {
    // This approach is more complex for generating actual combinations
    // Better suited for counting combinations only
    dp := make([][]int, target+1)
    dp[0] = []int{}
    
    for _, candidate := range candidates {
        for t := candidate; t <= target; t++ {
            if len(dp[t-candidate]) > 0 {
                // Build combinations (complex implementation)
                // This is simplified - actual implementation much more complex
            }
        }
    }
    
    return dp[target]
}
```
- **Pros**: Polynomial time for counting
- **Cons**: Complex for generating actual combinations

#### 2. Iterative Approach (O(2^(T/M)) time, O(T/M) space)
```go
func combinationSumIterative(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    result := [][]int{}
    queue := []struct {
        combo []int
        remaining int
        start int
    }{{[]int{}, target, 0}}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        if current.remaining == 0 {
            result = append(result, current.combo)
            continue
        }
        
        for i := current.start; i < len(candidates); i++ {
            if candidates[i] > current.remaining {
                break
            }
            
            newCombo := make([]int, len(current.combo))
            copy(newCombo, current.combo)
            newCombo = append(newCombo, candidates[i])
            
            queue = append(queue, struct {
                combo []int
                remaining int
                start int
            }{newCombo, current.remaining - candidates[i], i})
        }
    }
    
    return result
}
```
- **Pros**: Iterative, avoids recursion stack
- **Cons**: More complex state management

#### 3. Memoization with Backtracking (O(T × N) time, O(T × N) space)
```go
func combinationSumMemo(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    memo := make(map[int][][]int)
    
    var dfs func(int) [][]int
    dfs = func(remaining int) [][]int {
        if val, exists := memo[remaining]; exists {
            return val
        }
        
        var result [][]int
        for i, candidate := range candidates {
            if candidate > remaining {
                break
            }
            
            if candidate == remaining {
                result = append(result, []int{candidate})
            } else {
                for _, combo := range dfs(remaining - candidate) {
                    if len(combo) > 0 && combo[0] >= candidate {
                        newCombo := make([]int, len(combo)+1)
                        newCombo[0] = candidate
                        copy(newCombo[1:], combo)
                        result = append(result, newCombo)
                    }
                }
            }
        }
        
        memo[remaining] = result
        return result
    }
    
    return dfs(target)
}
```
- **Pros**: Avoids redundant calculations
- **Cons**: Complex memoization logic, still exponential worst case

### Extensions for Interviews:
- **Combination Sum II**: Each candidate used at most once
- **Combination Sum III**: Exactly k numbers that sum to target
- **Count Only**: Return count instead of actual combinations
- **Multiple Targets**: Find combinations for multiple target values
- **Large Numbers**: Handle very large candidate values efficiently
*/
func main() {
	// Test cases
	testCases := []struct {
		candidates []int
		target    int
	}{
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
		{[]int{2}, 1},
		{[]int{1}, 1},
		{[]int{1}, 2},
		{[]int{2, 3, 6, 7, 8, 10}, 10},
		{[]int{4, 5, 6, 7, 8}, 11},
		{[]int{3, 5, 7}, 0},
		{[]int{2, 4, 6, 8}, 16},
		{[]int{1, 2, 3, 4, 5}, 7},
	}
	
	for i, tc := range testCases {
		result := combinationSum(tc.candidates, tc.target)
		fmt.Printf("Test Case %d: candidates=%v, target=%d\n", i+1, tc.candidates, tc.target)
		fmt.Printf("  Combinations: %v\n\n", result)
	}
}
