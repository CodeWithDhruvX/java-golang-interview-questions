import java.util.Arrays;

public class JumpGame {
    
    // 55. Jump Game
    // Time: O(N), Space: O(1)
    public static boolean canJump(int[] nums) {
        int maxReach = 0;
        
        for (int i = 0; i < nums.length; i++) {
            if (i > maxReach) {
                return false;
            }
            maxReach = Math.max(maxReach, i + nums[i]);
            if (maxReach >= nums.length - 1) {
                return true;
            }
        }
        
        return true;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {2, 3, 1, 1, 4},
            {3, 2, 1, 0, 4},
            {0},
            {1},
            {2, 0, 0},
            {1, 1, 1, 1, 1},
            {3, 2, 1, 0, 4, 5},
            {2, 0, 0, 0, 1},
            {1, 2, 3},
            {100, 0, 0, 0, 0}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            boolean result = canJump(testCases[i]);
            System.out.printf("Test Case %d: %s -> Can jump: %b\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Greedy Algorithm
- **Jump Game**: Always reach the farthest possible position
- **Greedy Choice**: At each position, choose the maximum reachable jump
- **Range Expansion**: Update reachable range as we progress
- **Linear Pass**: Single pass through array

## 2. PROBLEM CHARACTERISTICS
- **Array Jumping**: Can jump nums[i] steps from current position
- **Reachability Check**: Can we reach or surpass the last index
- **Greedy Validity**: Local optimal choices lead to global optimum
- **Monotonic**: Once we can reach a position, we can always reach further

## 3. SIMILAR PROBLEMS
- Jump Game II (minimum jumps)
- Gas Station Problem
- H-Index Problem
- Candy Distribution Problem

## 4. KEY OBSERVATIONS
- If we can reach position i, we can reach any position ≤ i + nums[i]
- Maintain current maximum reachable position
- If maxReach >= last index, return true
- Greedy works because larger jumps only help, never hurt

## 5. VARIATIONS & EXTENSIONS
- Minimum number of jumps
- Jump with different costs
- Backward jumping
- Multiple test cases

## 6. INTERVIEW INSIGHTS
- Clarify: "Can I jump 0 steps?" (usually no)
- Edge cases: empty array, single element, all zeros
- Time complexity: O(N) vs O(N²) DP approach
- Space complexity: O(1) vs O(N) DP

## 7. COMMON MISTAKES
- Using DP when greedy suffices
- Not handling edge case of empty array
- Integer overflow for large jump values
- Off-by-one errors in reachability check

## 8. OPTIMIZATION STRATEGIES
- Early termination when reachability is guaranteed
- Single pass algorithm is optimal
- No need for extra data structures
- Handle large integers properly

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like expanding your reach:**
- You're at position 0 and want to reach the last index
- At each position i, you can jump nums[i] steps forward
- Your maximum reachable position expands as you progress
- If you can reach or pass the last index, you succeed
- Always take the biggest jump possible at each step

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of jump lengths
2. **Goal**: Can you reach the last index?
3. **Output**: Boolean indicating reachability

#### Phase 2: Key Insight Recognition
- **"What does reaching mean?"** → Can get to index ≥ last position
- **"How to expand reach?"** → maxReach = max(maxReach, i + nums[i])
- **"Why greedy works?"** → Larger jumps only increase reachability
- **"When to stop?"** → When maxReach ≥ last index

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use greedy:
1. Start at position 0, maxReach = 0
2. For each position i:
   - Update maxReach = max(maxReach, i + nums[i])
   - If maxReach ≥ last index: return true
3. Single pass, no need for complex data structures"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return true (already at destination)
- **Single element**: Return true (can reach it)
- **All zeros**: Return false (stuck at position 0)
- **Large nums**: Use long to prevent overflow

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
nums = [2, 3, 1, 1, 4]

Human thinking:
"Let's track maximum reach:

Position 0: maxReach = 0
- Jump nums[0] = 2: maxReach = max(0, 0+2) = 2
- Can reach index 2

Position 1: maxReach = 2
- Jump nums[1] = 3: maxReach = max(2, 1+3) = 4
- Can reach indices 1,2,3,4

Position 2: maxReach = 4
- Jump nums[2] = 1: maxReach = max(4, 2+1) = 4
- Can reach indices 2,3,4

Position 3: maxReach = 4
- Jump nums[3] = 1: maxReach = max(4, 3+1) = 4
- Can reach indices 3,4

Position 4: maxReach = 4
- Jump nums[4] = 4: maxReach = max(4, 4+4) = 8
- Can reach index 4 (last index) ✓

Result: true ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Greedy always maintains furthest possible reach
- **Why it's optimal**: Any solution that reaches further is better
- **Why it's efficient**: Single pass O(N) time

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all jump sequences?"** → Exponential O(2^N)
2. **"What about DP?"** → O(N²) time, O(N) space (unnecessary)
3. **"How to handle zeros?"** → Special case, stuck at position 0
4. **"What about negative jumps?"** → Usually not allowed

### Real-World Analogy
**Like crossing a river with stepping stones:**
- You have stepping stones (array positions)
- From each stone, you can jump nums[i] stones ahead
- You want to know if you can reach the opposite bank
- Always take the longest jump possible to maximize reach
- If you can reach or pass the opposite bank, you succeed
- The greedy strategy ensures you always make the furthest progress

### Human-Readable Pseudocode
```
function canJump(nums):
    if nums.length == 0:
        return true
    
    maxReach = 0
    last = nums.length - 1
    
    for i from 0 to last:
        maxReach = max(maxReach, i + nums[i])
        if maxReach >= last:
            return true
    
    return false
```

### Execution Visualization

### Example: nums = [2, 3, 1, 1, 4]
```
Jump Process:
Position: 0  1  2  3  4
Jump:    2  3  1  1  4
MaxReach: 0  2  4  4  4  8

Step-by-step:
Pos 0: Jump 2 → Reach 2, Max=2
Pos 1: Jump 3 → Reach 4, Max=4
Pos 2: Jump 1 → Reach 3, Max=4
Pos 3: Jump 1 → Reach 4, Max=4
Pos 4: Jump 4 → Reach 8, Max=8 ≥ 4 ✓

Result: true ✓

The greedy path: 0→2→4→8 (success) ✓
```

### Key Visualization Points:
- **Greedy expansion** of maximum reachable position
- **Single pass** through array
- **Early termination** when destination is reachable
- **Monotonic progress**: reach only expands

### Memory Layout Visualization:
```
Variable Evolution:
Position: 0  1  2  3  4
Jump:    2  3  1  1  4
MaxReach: 0  2  4  4  4  8

Decision Process:
Pos 0: Can reach 2, Max=2
Pos 1: Can reach 4, Max=4
Pos 2: Can reach 3, Max=4
Pos 3: Can reach 4, Max=4
Pos 4: Can reach 8, Max=8 ≥ 4 ✓

Final Answer: true ✓
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) iterations
- **Each Iteration**: O(1) operations (max, addition, comparison)
- **Total**: O(N) time, O(1) space
- **Optimal**: Cannot do better than O(N) for this problem
- **vs DP**: O(N²) time, O(N) space (unnecessary)
*/
