package main

import "fmt"

// 55. Jump Game
// Time: O(N), Space: O(1)
func canJump(nums []int) bool {
	maxReach := 0
	
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	
	return true
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

## 1. ALGORITHM PATTERN: Greedy with Reachable Range Tracking
- **Greedy Strategy**: Always make the locally optimal jump
- **Reachable Range**: Track the farthest position that can be reached
- **Single Pass**: Process array from left to right
- **Early Termination**: Stop if we can reach the end

## 2. PROBLEM CHARACTERISTICS
- **Array Navigation**: Jump from position i to i+nums[i]
- **Goal Achievement**: Reach the last index
- **Greedy Validity**: Local optimal choices lead to global optimum
- **Range Expansion**: Each jump expands the reachable range

## 3. SIMILAR PROBLEMS
- Jump Game II (LeetCode 45) - Minimum number of jumps
- Gas Station (LeetCode 134) - Circular route feasibility
- Candy (LeetCode 135) - Distribution with constraints
- Partition Labels (LeetCode 763) - String partitioning

## 4. KEY OBSERVATIONS
- **Reachable Range**: If you can reach position i, you can reach any position ≤ i
- **Greedy Choice**: Always jump to the position that maximizes reachability
- **Early Success**: If maxReach ≥ last index, return true
- **Failure Condition**: If current position exceeds maxReach, return false

## 5. VARIATIONS & EXTENSIONS
- **Minimum Jumps**: Find minimum number of jumps to reach end
- **Jump Variations**: Different jump rules or constraints
- **Multiple Queries**: Answer many jump feasibility queries
- **Obstacle Avoidance**: Some positions are blocked

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can we jump backwards? What about negative values?"
- Edge cases: empty array, single element, all zeros
- Time complexity: O(N) time, O(1) space
- Greedy works because larger jumps are always better
- Key insight: track farthest reachable position

## 7. COMMON MISTAKES
- Not updating maxReach correctly
- Wrong termination condition
- Not handling edge cases (empty array, single element)
- Using complex DP when greedy works
- Off-by-one errors in array indexing

## 8. OPTIMIZATION STRATEGIES
- **Greedy Approach**: O(N) time, O(1) space - optimal
- **Early Termination**: Stop as soon as end is reachable
- **Single Pass**: Process array once from left to right
- **Range Tracking**: Maintain maximum reachable position

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like planning a journey with the best possible moves:**
- You're at position 0 and can jump forward
- Each position tells you how far you can jump from there
- You want to know if you can reach the destination
- The key insight: if you can reach position i, you can reach any position ≤ i
- Always make the jump that gives you the farthest reach

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array nums where nums[i] is max jump from position i
2. **Goal**: Determine if you can reach the last index
3. **Constraint**: Start at position 0, can only jump forward
4. **Output**: Boolean indicating reachability

#### Phase 2: Key Insight Recognition
- **"Greedy natural fit"** → Always make the jump that maximizes reachability
- **"Range expansion"** → Track farthest reachable position
- **"Early termination"** → Stop if end is reachable
- **"Local optimal = global optimal"** → Larger jumps are always better

#### Phase 3: Strategy Development
```
Human thought process:
"I need to determine if I can reach the end.
I'll track the farthest position I can reach so far.
At each position, I can jump nums[i] positions forward.
So from position i, I can reach up to i + nums[i].
I'll keep updating my maximum reachable position.
If at any point my maxReach ≥ last index, I can succeed!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Cannot reach anything
- **Single element**: Can reach end if nums[0] > 0
- **All zeros**: Stuck at starting position
- **Large jumps**: Early termination when end is reachable

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
nums = [2,3,1,1,4]

Human thinking:
"I start at position 0, maxReach = 0.

Position 0: nums[0] = 2
- I can jump to position 0+2 = 2
- Update maxReach = max(0, 2) = 2
- Can I reach end? No (position 4)

Position 1: nums[1] = 3  
- I can jump to position 1+3 = 4
- Update maxReach = max(2, 4) = 4
- Can I reach end? Yes! maxReach = 4 ≥ 4
- Return true ✓"

Let me trace another example:
nums = [3,2,1,0,4]

Position 0: nums[0] = 3
- Can reach position 3, maxReach = 3

Position 1: nums[1] = 2
- Can reach position 3, maxReach = 3

Position 2: nums[2] = 1
- Can reach position 3, maxReach = 3

Position 3: nums[3] = 0
- Can reach position 3, maxReach = 3
- Current position (3) > maxReach (3)? No, position equals maxReach

Position 4: nums[4] = 4
- Current position (4) > maxReach (3)? Yes!
- Cannot proceed further, return false ✗"
```

#### Phase 6: Intuition Validation
- **Why greedy works**: Larger jumps always increase or maintain reachability
- **Why range tracking**: If you can reach i, you can reach any position ≤ i
- **Why O(N)**: Single pass through array
- **Why O(1) space**: Only need to track maxReach

### Common Human Pitfalls & How to Avoid Them
1. **"Why not DP?"** → Greedy is sufficient and more efficient
2. **"Should I try all paths?"** → No need, greedy guarantees optimality
3. **"What about negative jumps?"** → Clarify problem constraints
4. **"Can I optimize further?"** → Greedy is already optimal

### Real-World Analogy
**Like planning a road trip with gas stations:**
- You're at starting point with various possible routes
- Each station tells you how far you can travel from there
- You want to know if you can reach your destination
- If you can reach a certain point, you can reach any point before it
- Always choose the route that gets you farthest

### Human-Readable Pseudocode
```
function canJump(nums):
    if nums is empty:
        return false
    
    maxReach = 0
    
    for i from 0 to len(nums)-1:
        if i > maxReach:
            return false  // Cannot reach this position
        
        maxReach = max(maxReach, i + nums[i])
        
        if maxReach >= len(nums) - 1:
            return true  // Can reach the end
    
    return maxReach >= len(nums) - 1
```

### Execution Visualization

### Example: nums = [2,3,1,1,4]
```
Array Traversal:
Position: 0, nums[0] = 2, maxReach = max(0, 0+2) = 2
Position: 1, nums[1] = 3, maxReach = max(2, 1+3) = 4
Position: 2, nums[2] = 1, maxReach = max(4, 2+1) = 4
Position: 3, nums[3] = 1, maxReach = max(4, 3+1) = 4

Key moments:
- At position 1: maxReach becomes 4, which equals last index
- Since maxReach ≥ 4, we can reach the end ✓
```

### Key Visualization Points:
- **Reachable Range**: [0, maxReach] is always reachable
- **Greedy Choice**: Always maximize reachability
- **Early Success**: Terminate as soon as end is reachable
- **Failure Condition**: Current position exceeds maxReach

### Memory Layout Visualization:
```
Reachability Expansion:
nums = [2,3,1,1,4]

Step 0: maxReach = 0
After position 0: maxReach = max(0, 0+2) = 2
Reachable range: [0,1,2]

Step 1: maxReach = 2  
After position 1: maxReach = max(2, 1+3) = 4
Reachable range: [0,1,2,3,4] ✓ End reached!
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) time complexity
- **Constant Space**: O(1) additional space
- **Early Termination**: Can stop early when end is reachable
- **Optimal**: Greedy approach is provably optimal

### Alternative Approaches:

#### 1. Dynamic Programming (O(N²) time, O(N) space)
```go
func canJumpDP(nums []int) bool {
    n := len(nums)
    dp := make([]bool, n)
    dp[0] = true
    
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if dp[j] && j+nums[j] >= i {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[n-1]
}
```
- **Pros**: More general approach
- **Cons**: O(N²) time, unnecessary complexity

#### 2. BFS/Queue Approach (O(N²) time, O(N) space)
```go
func canJumpBFS(nums []int) bool {
    n := len(nums)
    visited := make([]bool, n)
    queue := []int{0}
    visited[0] = true
    
    for len(queue) > 0 {
        pos := queue[0]
        queue = queue[1:]
        
        if pos == n-1 {
            return true
        }
        
        for jump := 1; jump <= nums[pos]; jump++ {
            next := pos + jump
            if next < n && !visited[next] {
                visited[next] = true
                queue = append(queue, next)
            }
        }
    }
    
    return false
}
```
- **Pros**: Guarantees correctness
- **Cons**: O(N²) time, unnecessary complexity

#### 3. Recursive with Memoization (O(N²) time, O(N) space)
```go
func canJumpRecursive(nums []int) bool {
    memo := make(map[int]bool)
    return canJumpHelper(nums, 0, memo)
}

func canJumpHelper(nums []int, pos int, memo map[int]bool) bool {
    if pos >= len(nums)-1 {
        return true
    }
    if pos >= len(nums) {
        return false
    }
    
    if val, exists := memo[pos]; exists {
        return val
    }
    
    for jump := 1; jump <= nums[pos]; jump++ {
        if canJumpHelper(nums, pos+jump, memo) {
            memo[pos] = true
            return true
        }
    }
    
    memo[pos] = false
    return false
}
```
- **Pros**: Intuitive approach
- **Cons**: O(N²) time, unnecessary complexity

### Extensions for Interviews:
- **Minimum Jumps**: Find minimum number of jumps to reach end
- **Jump Variations**: Different jump rules or constraints
- **Multiple Queries**: Answer many jump feasibility queries
- **Obstacle Avoidance**: Some positions are blocked
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
		{0},
		{1},
		{2, 0, 0},
		{1, 1, 1, 1, 1},
		{3, 2, 1, 0, 4, 5},
		{2, 0, 0, 0, 1},
		{1, 2, 3},
		{100, 0, 0, 0, 0},
	}
	
	for i, nums := range testCases {
		result := canJump(nums)
		fmt.Printf("Test Case %d: %v -> Can jump: %t\n", i+1, nums, result)
	}
}
