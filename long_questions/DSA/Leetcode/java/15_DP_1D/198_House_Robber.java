import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class HouseRobber {
    
    // 198. House Robber
    // Time: O(N), Space: O(1)
    public static int rob(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        
        int prev2 = nums[0];
        int prev1 = Math.max(nums[0], nums[1]);
        
        for (int i = 2; i < nums.length; i++) {
            int current = Math.max(prev1, prev2 + nums[i]);
            prev2 = prev1;
            prev1 = current;
        }
        
        return prev1;
    }

    // DP array approach
    public static int robDP(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        
        int[] dp = new int[nums.length];
        dp[0] = nums[0];
        dp[1] = Math.max(nums[0], nums[1]);
        
        for (int i = 2; i < nums.length; i++) {
            dp[i] = Math.max(dp[i - 1], dp[i - 2] + nums[i]);
        }
        
        return dp[nums.length - 1];
    }

    // Recursive with memoization
    public static int robMemo(int[] nums) {
        Map<Integer, Integer> memo = new HashMap<>();
        return robHelper(nums, nums.length - 1, memo);
    }

    private static int robHelper(int[] nums, int index, Map<Integer, Integer> memo) {
        if (index < 0) {
            return 0;
        }
        if (index == 0) {
            return nums[0];
        }
        
        if (memo.containsKey(index)) {
            return memo.get(index);
        }
        
        int result = Math.max(robHelper(nums, index - 1, memo), 
                             robHelper(nums, index - 2, memo) + nums[index]);
        memo.put(index, result);
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 2, 3, 1},
            {2, 7, 9, 3, 1},
            {2, 1, 1, 2},
            {1},
            {1, 2},
            {5, 5, 10, 100, 10, 5},
            {2, 7, 9, 3, 1, 5, 6, 8},
            {100, 1, 1, 100},
            {4, 1, 2, 7, 5, 3, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result1 = rob(testCases[i]);
            int result2 = robDP(testCases[i]);
            int result3 = robMemo(testCases[i]);
            
            System.out.printf("Test Case %d: %s -> Iterative: %d, DP: %d, Memo: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result1, result2, result3);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: 1D Dynamic Programming
- **House Robber Problem**: Cannot rob adjacent houses
- **DP State**: Maximum money up to current house
- **Recurrence**: dp[i] = max(dp[i-1], dp[i-2] + nums[i])
- **Space Optimization**: Use O(1) space instead of O(N) array

## 2. PROBLEM CHARACTERISTICS
- **Linear Arrangement**: Houses in a line, cannot rob adjacent
- **Binary Choice**: For each house, either rob it or skip it
- **Optimal Substructure**: Best solution up to house i depends on previous choices
- **No Circular Constraint**: First and last houses are not adjacent

## 3. SIMILAR PROBLEMS
- House Robber II (circular houses)
- Maximum Sum of Non-Adjacent Elements
- Coin Change with constraints
- Stock Buy/Sell with cooldown

## 4. KEY OBSERVATIONS
- For house i: either rob it (and skip i-1) OR skip it
- This creates dp[i] = max(dp[i-1], dp[i-2] + nums[i])
- Need to handle base cases: first house, second house
- Space can be optimized to track only last two values

## 5. VARIATIONS & EXTENSIONS
- Circular houses (first and last adjacent)
- K houses apart constraint
- Different values for houses
- Multiple robbers working simultaneously

## 6. INTERVIEW INSIGHTS
- Clarify: "Are houses in a circle?"
- Edge cases: empty array, single house, two houses
- Time complexity: O(N) vs O(2^N) brute force
- Space complexity: O(1) vs O(N) array DP

## 7. COMMON MISTAKES
- Not handling first two houses correctly
- Using O(2^N) brute force approach
- Integer overflow for large sums
- Off-by-one errors in DP indexing

## 8. OPTIMIZATION STRATEGIES
- Space optimization using two variables
- Early termination for obvious cases
- Handling negative values correctly
- Modulo arithmetic for large results

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like planning a heist:**
- You have houses in a street, each with some money
- You can't rob two adjacent houses (police will catch you)
- For each house, you decide: rob it OR skip it
- If you rob house i, you must skip house i-1
- If you skip house i, you can consider robbing i-1
- You want to maximize total money stolen

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of house values
2. **Goal**: Maximum money without robbing adjacent houses
3. **Output**: Maximum achievable sum

#### Phase 2: Key Insight Recognition
- **"What's the constraint?"** → Cannot rob adjacent houses
- **"What are my choices?"** → For each house: rob OR skip
- **"What's the recurrence?"** → dp[i] = max(skip i, rob i)
- **"How to optimize?"** → Only need previous two states

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use DP:
1. For each house i:
   - Option 1: Skip house i, take dp[i-1]
   - Option 2: Rob house i, take dp[i-2] + nums[i]
   - Choose maximum of these two options
2. Handle base cases separately
3. Optimize space to track only last two values"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0
- **Single house**: Return its value
- **Two houses**: Return max(nums[0], nums[1])
- **All negative**: Return 0 or maximum (depending on interpretation)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Houses: [1, 2, 3, 1]

Human thinking:
"Let's go house by house:

House 0 (value 1):
- Rob it: take 1, previous houses don't exist
- dp[0] = 1

House 1 (value 2):
- Option 1: Skip house 1, take dp[0] = 1
- Option 2: Rob house 1, take dp[-1] + 2 = 2
- Choose max: dp[1] = max(1, 2) = 2

House 2 (value 3):
- Option 1: Skip house 2, take dp[1] = 2
- Option 2: Rob house 2, take dp[0] + 3 = 1 + 3 = 4
- Choose max: dp[2] = max(2, 4) = 4

House 3 (value 1):
- Option 1: Skip house 3, take dp[2] = 4
- Option 2: Rob house 3, take dp[1] + 1 = 2 + 1 = 3
- Choose max: dp[3] = max(4, 3) = 4

Final answer: 4 ✓

The optimal strategy:
Skip house 0 (1), rob house 1 (2), skip house 2 (3), rob house 3 (1)
Total: 1 + 2 + 1 = 4"
```

#### Phase 6: Intuition Validation
- **Why it works**: Each house decision is optimal given previous choices
- **Why it's efficient**: Each house processed once
- **Why it's correct**: DP ensures no adjacent houses are robbed

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just try all combinations?"** → Exponential O(2^N)
2. **"What about greedy?"** → Local optimum may not be global optimum
3. **"How to handle first house?"** → Special case, no previous house
4. **"What about negative values?"** → May need to skip all houses

### Real-World Analogy
**Like planning a route with gas stations:**
- You have gas stations along a highway (houses with money)
- You can only stop at every other station (can't use adjacent ones)
- For each station, decide: stop there OR drive past it
- If you stop at station i, you must skip station i-1
- If you skip station i, you can consider stopping at i-1
- You want to maximize total gas collected

### Human-Readable Pseudocode
```
function rob(nums):
    if nums.length == 0:
        return 0
    if nums.length == 1:
        return nums[0]
    
    prev2 = nums[0]  // dp[-1] = nums[0]
    prev1 = max(nums[0], nums[1])  // dp[0]
    
    for i from 2 to nums.length-1:
        current = max(prev1, prev2 + nums[i])
        prev2 = prev1
        prev1 = current
    
    return prev1
```

### Execution Visualization

### Example: nums = [1, 2, 3, 1]
```
DP Table Evolution:
i:   -1   0   1   2   3
nums:  -    1   2   3   1
dp:    -    1   2   4   4

Step-by-step:
dp[0] = nums[0] = 1
dp[1] = max(dp[0], dp[-1] + nums[1]) = max(1, 2) = 2
dp[2] = max(dp[1], dp[0] + nums[2]) = max(2, 4) = 4
dp[3] = max(dp[2], dp[1] + nums[3]) = max(4, 3) = 4

Optimal Strategy:
Skip house 0 (1), rob house 1 (2), skip house 2 (3), rob house 3 (1)
Total: 1 + 2 + 1 = 4 ✓
```

### Key Visualization Points:
- **DP recurrence** considers both choices for each house
- **Base cases** handle first houses specially
- **Space optimization** tracks only last two states
- **Maximum operation** chooses optimal local decision

### Memory Layout Visualization:
```
Variable Evolution:
House:   0   1   2   3
nums:    1   2   3   1
prev2:   1   1   2   4
prev1:   1   2   4   4
current: -   1   2   4   4

Decision Process:
House 0: Rob (1) → prev1 = 1
House 1: Rob (2) → prev1 = 2
House 2: Skip (4) → prev1 = 4
House 3: Rob (3) → prev1 = 4

Final Answer: 4 ✓
```

### Time Complexity Breakdown:
- **DP Loop**: O(N) iterations
- **Each Iteration**: O(1) operations (max, addition)
- **Total**: O(N) time, O(1) space
- **Optimal**: Cannot do better than O(N) for this problem
- **vs Brute Force**: O(2^N) time trying all subsets
*/
