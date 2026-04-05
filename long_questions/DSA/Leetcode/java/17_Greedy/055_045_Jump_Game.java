import java.util.*;

public class JumpGame {
    
    // 55. Jump Game
    // Time: O(N), Space: O(1)
    public static boolean canJump(int[] nums) {
        if (nums == null || nums.length == 0) {
            return true;
        }
        
        int farthest = 0;
        
        for (int i = 0; i < nums.length; i++) {
            if (i > farthest) {
                return false; // Can't reach this position
            }
            
            farthest = Math.max(farthest, i + nums[i]);
            
            if (farthest >= nums.length - 1) {
                return true; // Can reach the end
            }
        }
        
        return true;
    }

    // 45. Jump Game II
    // Time: O(N), Space: O(1)
    public static int jump(int[] nums) {
        if (nums == null || nums.length <= 1) {
            return 0;
        }
        
        int jumps = 0;
        int currentEnd = 0;
        int farthest = 0;
        
        for (int i = 0; i < nums.length - 1; i++) {
            farthest = Math.max(farthest, i + nums[i]);
            
            if (i == currentEnd) {
                jumps++;
                currentEnd = farthest;
                
                if (currentEnd >= nums.length - 1) {
                    break;
                }
            }
        }
        
        return jumps;
    }

    public static void main(String[] args) {
        int[][] testCases1 = {
            {2, 3, 1, 1, 4},
            {3, 2, 1, 0, 4},
            {0},
            {1},
            {2, 0, 0},
            {1, 2, 3},
            {2, 5, 0, 0, 1, 1, 0},
            {1, 1, 1, 1, 1},
            {3, 2, 1, 0, 4, 2},
            {2, 0, 1, 0, 1}
        };
        
        int[][] testCases2 = {
            {2, 3, 1, 1, 4},
            {2, 3, 0, 1, 4},
            {0},
            {1},
            {1, 2},
            {2, 3, 1, 1, 4},
            {1, 1, 1, 1},
            {2, 0, 2, 0, 1},
            {3, 2, 1, 0, 4},
            {1, 2, 0, 1, 4}
        };
        
        System.out.println("Jump Game I - Can reach end?");
        for (int i = 0; i < testCases1.length; i++) {
            int[] nums = testCases1[i];
            boolean result = canJump(nums);
            System.out.printf("Test Case %d: %s -> %b\n", 
                i + 1, Arrays.toString(nums), result);
        }
        
        System.out.println("\nJump Game II - Minimum jumps");
        for (int i = 0; i < testCases2.length; i++) {
            int[] nums = testCases2[i];
            int result = jump(nums);
            System.out.printf("Test Case %d: %s -> %d jumps\n", 
                i + 1, Arrays.toString(nums), result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Greedy Range Coverage
- **Jump Game I**: Track farthest reachable position
- **Jump Game II**: Track current jump range boundaries
- **Greedy Strategy**: Always make optimal local choices
- **Range Expansion**: Expand reachable window progressively

## 2. PROBLEM CHARACTERISTICS
- **Array Representation**: Each element shows max jump length from that position
- **Reachability**: Can we reach the end from current position?
- **Optimization**: Minimum number of jumps (Jump Game II)
- **Greedy Validity**: Local optimal choices lead to global optimum

## 3. SIMILAR PROBLEMS
- Gas Station problem
- Candy distribution
- Minimum number of intervals
- Array partitioning problems

## 4. KEY OBSERVATIONS
- **Jump Game I**: If position > farthest reachable, impossible
- **Jump Game II**: When current position reaches end of current jump range, make a jump
- **Farthest Position**: Always track maximum reachable index
- **Early Termination**: Can stop when end is reachable

## 5. VARIATIONS & EXTENSIONS
- Backward approach (from end to start)
- Dynamic programming solution
- BFS approach for minimum jumps
- Variations with obstacles or costs

## 6. INTERVIEW INSIGHTS
- Clarify: "Can I jump beyond array bounds?"
- Edge cases: empty array, single element, all zeros
- Time complexity: O(N) is optimal for greedy approach
- Space complexity: O(1) - only need few variables

## 7. COMMON MISTAKES
- Not handling empty array case
- Off-by-one errors in index calculations
- Not updating farthest position correctly
- Infinite loop in Jump Game II without proper termination

## 8. OPTIMIZATION STRATEGIES
- Greedy approach is optimal for both problems
- Jump Game I: Single pass with farthest tracking
- Jump Game II: Range-based greedy with currentEnd tracking
- No need for DP or recursion

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like crossing a river with stepping stones:**
- Each stone (array element) tells you how far you can jump
- You want to know if you can reach the other side (end of array)
- For minimum jumps, you want to use the fewest, longest jumps possible
- Always jump from the position that gets you farthest

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding (Jump Game I)
1. **Input**: Array of non-negative integers
2. **Goal**: Can you reach the last index?
3. **Output**: Boolean indicating reachability

#### Phase 2: Key Insight Recognition
- **"How far can I get?"** → Track farthest reachable position
- **"When is it impossible?"** → Current position exceeds farthest reachable
- **"When can I stop early?"** → Farthest reaches or exceeds end
- **Greedy choice**: Always extend reach as far as possible

#### Phase 3: Strategy Development (Jump Game I)
```
Human thought process:
"I'll track the farthest position I can reach:
1. Start at position 0, farthest = 0
2. For each position i:
   - If i > farthest, I can't reach this position → return false
   - Update farthest = max(farthest, i + nums[i])
   - If farthest >= last index → return true
3. If I finish the loop, I can reach the end"
```

#### Phase 4: Algorithm Walkthrough (Jump Game I)
```
Example: [2,3,1,1,4]

Human thinking:
"Start: farthest = 0, i = 0
i=0: nums[0]=2, farthest = max(0, 0+2) = 2
i=1: nums[1]=3, farthest = max(2, 1+3) = 4
i=2: nums[2]=1, farthest = max(4, 2+1) = 4
i=3: nums[3]=1, farthest = max(4, 3+1) = 4

Since farthest (4) >= last index (4), I can reach the end ✓"
```

#### Phase 5: Strategy Development (Jump Game II)
```
Human thought process:
"For minimum jumps, I need to be greedy:
1. Track current jump range [currentStart, currentEnd]
2. Within this range, find the farthest position for next jump
3. When I reach currentEnd, make a jump to farthest
4. Repeat until I reach the end"
```

#### Phase 6: Algorithm Walkthrough (Jump Game II)
```
Example: [2,3,1,1,4]

Human thinking:
"Jump 1: Range [0,0], farthest from this range = max(0+2, 0+3, 0+1, 0+1, 0+4) = 4
Jump 2: Range [1,4], but I stop at currentEnd=0 first
Actually, let me trace correctly:

i=0: farthest = max(0, 0+2) = 2
i=0 == currentEnd(0): jumps=1, currentEnd=2

i=1: farthest = max(2, 1+3) = 4  
i=2: farthest = max(4, 2+1) = 4
i=2 == currentEnd(2): jumps=2, currentEnd=4

Since currentEnd(4) >= last index(4), done! Total jumps = 2 ✓"
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use DP?"** → Greedy is more efficient and sufficient
2. **"What about backward approach?"** → Forward is more intuitive
3. **"How to track ranges?"** → Use currentEnd and farthest variables
4. **"When to stop?"** → When currentEnd reaches or exceeds last index

### Real-World Analogy
**Like planning a multi-leg flight journey:**
- Each airport (array position) has connections to other airports
- nums[i] tells you the farthest airport you can reach from here
- Jump Game I: Can you reach your final destination?
- Jump Game II: What's the minimum number of flights needed?

### Human-Readable Pseudocode
```
function canJump(nums):
    farthest = 0
    for i from 0 to nums.length-1:
        if i > farthest:
            return false  // Can't reach this position
        farthest = max(farthest, i + nums[i])
        if farthest >= nums.length-1:
            return true
    return true

function minJumps(nums):
    if nums.length <= 1: return 0
    
    jumps = 0
    currentEnd = 0
    farthest = 0
    
    for i from 0 to nums.length-2:
        farthest = max(farthest, i + nums[i])
        if i == currentEnd:
            jumps++
            currentEnd = farthest
            if currentEnd >= nums.length-1:
                break
    
    return jumps
```

### Execution Visualization

### Example: [2,3,1,1,4] for Jump Game I
```
Position:   0   1   2   3   4
Values:     2   3   1   1   4
Reachable:  [✓] [✓] [✓] [✓] [✓]

Step-by-step:
i=0: farthest = max(0, 0+2) = 2
i=1: farthest = max(2, 1+3) = 4 ← Can reach end!
Result: TRUE ✓
```

### Example: [2,3,1,1,4] for Jump Game II
```
Jump Ranges:
Jump 1: [0] → farthest=2, jumps=1, nextRange=[1,2]
Jump 2: [1,2] → farthest=4, jumps=2, nextRange=[3,4]

Visualization:
[2]  [3,1]  [1,4]
 ↑     ↑↑↑    ↑↑↑
 |     ||||    ||||
Jump1  Jump2   Reached!

Result: 2 jumps ✓
```

### Key Visualization Points:
- **Farthest Position**: Maximum index reachable from current position
- **Range Boundaries**: Current jump defines a range to explore
- **Optimal Jumps**: Always jump from the best position in current range
- **Greedy Choice**: Local optimal leads to global optimal

### Time Complexity Breakdown:
- **Jump Game I**: O(N) - single pass through array
- **Jump Game II**: O(N) - single pass with constant work per element
- **Space**: O(1) - only few integer variables
- **Optimal**: Cannot do better than O(N) for this problem
*/
