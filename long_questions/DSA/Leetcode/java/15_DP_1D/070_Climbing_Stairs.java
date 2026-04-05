import java.util.HashMap;
import java.util.Map;

public class ClimbingStairs {
    
    // 70. Climbing Stairs
    // Time: O(N), Space: O(1)
    public static int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }
        
        int prev2 = 1, prev1 = 2;
        
        for (int i = 3; i <= n; i++) {
            int current = prev1 + prev2;
            prev2 = prev1;
            prev1 = current;
        }
        
        return prev1;
    }

    // DP array approach
    public static int climbStairsDP(int n) {
        if (n <= 2) {
            return n;
        }
        
        int[] dp = new int[n + 1];
        dp[1] = 1;
        dp[2] = 2;
        
        for (int i = 3; i <= n; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }
        
        return dp[n];
    }

    // Recursive with memoization
    public static int climbStairsMemo(int n) {
        Map<Integer, Integer> memo = new HashMap<>();
        return climbStairsHelper(n, memo);
    }

    private static int climbStairsHelper(int n, Map<Integer, Integer> memo) {
        if (n <= 2) {
            return n;
        }
        
        if (memo.containsKey(n)) {
            return memo.get(n);
        }
        
        int result = climbStairsHelper(n - 1, memo) + climbStairsHelper(n - 2, memo);
        memo.put(n, result);
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[] testCases = {1, 2, 3, 4, 5, 10, 20, 30, 45, 50};
        
        for (int i = 0; i < testCases.length; i++) {
            int n = testCases[i];
            int result1 = climbStairs(n);
            int result2 = climbStairsDP(n);
            int result3 = climbStairsMemo(n);
            
            System.out.printf("Test Case %d: n=%d -> Iterative: %d, DP: %d, Memo: %d\n", 
                i + 1, n, result1, result2, result3);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: 1D Dynamic Programming
- **Fibonacci-style DP**: Each step depends on previous two steps
- **Optimal Substructure**: Solution to n depends on n-1 and n-2
- **Bottom-Up**: Build solution iteratively from base cases
- **Space Optimization**: Use O(1) space instead of O(N) array

## 2. PROBLEM CHARACTERISTICS
- **Stair Climbing**: Can take 1 or 2 steps at a time
- **Ways to Reach**: Number of distinct ways to reach step n
- **Overlapping Subproblems**: Same subproblems solved repeatedly
- **Optimal Substructure**: Best solution built from optimal sub-solutions

## 3. SIMILAR PROBLEMS
- Fibonacci Numbers
- Minimum Cost Climbing Stairs
- Number of Ways to Decode Message
- Tribonacci Numbers

## 4. KEY OBSERVATIONS
- Ways(n) = Ways(n-1) + Ways(n-2)
- Base cases: Ways(0) = 1, Ways(1) = 1, Ways(2) = 2
- This is exactly Fibonacci sequence shifted by 1
- Space can be optimized to O(1) using two variables

## 5. VARIATIONS & EXTENSIONS
- Different step sizes (1, 2, 3 steps)
- Minimum cost to climb
- Obstacles on stairs
- Circular stairs

## 6. INTERVIEW INSIGHTS
- Clarify: "Can I take 0 steps?" (usually no)
- Edge cases: n = 0, 1, 2
- Time complexity: O(N) vs O(2^N) naive recursion
- Space complexity: O(1) vs O(N) array DP

## 7. COMMON MISTAKES
- Not handling base cases correctly
- Using O(2^N) naive recursion
- Integer overflow for large n
- Off-by-one errors in DP indexing

## 8. OPTIMIZATION STRATEGIES
- Space optimization using two variables
- Matrix exponentiation for very large n
- Iterative vs recursive approaches
- Modulo arithmetic for large results

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like climbing stairs with choices:**
- You're at the bottom (step 0) and want to reach step n
- At each step, you can either take 1 step or 2 steps
- The number of ways to reach step n equals:
  - Ways to reach n-1 (then take 1 step)
  - Plus ways to reach n-2 (then take 2 steps)
- This creates a Fibonacci-like recurrence

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Integer n (number of stairs)
2. **Goal**: Number of distinct ways to climb to top
3. **Output**: Count of ways to reach step n

#### Phase 2: Key Insight Recognition
- **"How can I reach step n?"** → From n-1 with 1 step OR from n-2 with 2 steps
- **"What's the recurrence?"** → Ways(n) = Ways(n-1) + Ways(n-2)
- **"Are there base cases?"** → Ways(0) = 1, Ways(1) = 1
- **"Can I optimize?"** → Only need last two values

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use DP bottom-up:
1. Base cases: Ways(0) = 1, Ways(1) = 1
2. For i from 2 to n:
   - Ways(i) = Ways(i-1) + Ways(i-2)
3. Optimize space: only keep last two values
4. Return Ways(n)"
```

#### Phase 4: Edge Case Handling
- **n = 0**: 1 way (already at top)
- **n = 1**: 1 way (single step)
- **n = 2**: 2 ways (1+1 or 2)
- **Large n**: Use iterative to avoid recursion stack

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
n = 4

Human thinking:
"Let's build up from base cases:
Ways(0) = 1 (base case)
Ways(1) = 1 (base case)
Ways(2) = Ways(1) + Ways(0) = 1 + 1 = 2
Ways(3) = Ways(2) + Ways(1) = 2 + 1 = 3
Ways(4) = Ways(3) + Ways(2) = 3 + 2 = 5

Let me verify:
To reach step 4:
- From step 3 with 1 step: 3 ways
- From step 2 with 2 steps: 2 ways
- Total: 3 + 2 = 5 ways ✓

The 5 ways are:
1. 1+1+1+1 (four 1-steps)
2. 1+1+2 (two 1-steps, one 2-step)
3. 1+2+1 (one 1-step, two 2-steps)
4. 2+1+1 (one 2-step, two 1-steps)
5. 2+2 (two 2-steps)"
```

#### Phase 6: Intuition Validation
- **Why it works**: Each way to n must come from n-1 or n-2
- **Why it's efficient**: Each subproblem solved once
- **Why it's correct**: DP builds optimal from optimal sub-solutions

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just try all combinations?"** → Exponential O(2^N)
2. **"What about recursion?"** → Stack overflow for large n
3. **"How to handle base cases?"** → Ways(0) and Ways(1) are both 1
4. **"What about large n?"** → Use iterative DP

### Real-World Analogy
**Like counting delivery routes:**
- You're delivering packages to floor n
- From each floor, you can send to floor+1 or floor+2
- You want to know how many different delivery routes exist
- Routes to floor n = routes to floor n-1 + routes to floor n-2
- Base case: 1 way to stay at ground floor (already there)
- This creates the same recurrence as climbing stairs

### Human-Readable Pseudocode
```
function climbStairs(n):
    if n <= 1:
        return 1
    if n == 2:
        return 2
    
    prev2 = 1  // Ways(0)
    prev1 = 2  // Ways(1)
    
    for i from 3 to n:
        current = prev1 + prev2
        prev2 = prev1
        prev1 = current
    
    return prev1
```

### Execution Visualization

### Example: n = 4
```
DP Table Building:
i:   0  1  2  3  4
Ways: 1  1  2  3  5

Step-by-step:
Ways(0) = 1 (base)
Ways(1) = 1 (base)
Ways(2) = Ways(1) + Ways(0) = 1 + 1 = 2
Ways(3) = Ways(2) + Ways(1) = 2 + 1 = 3
Ways(4) = Ways(3) + Ways(2) = 3 + 2 = 5

The 5 ways to reach step 4:
1+1+1+1
1+1+2
1+2+1
2+1+1
2+2
```

### Key Visualization Points:
- **DP recurrence** builds from previous two values
- **Space optimization** uses only O(1) memory
- **Bottom-up approach** avoids recursion overhead
- **Fibonacci pattern** shifted by 1 position

### Memory Layout Visualization:
```
Variable Evolution:
Step:    0  1  2  3  4
prev2:   1  1  1  2  3
prev1:   1  2  2  3  5
current:  -  -  2  3  5

Final Answer: 5 ways

Physical Interpretation:
Step 4 can be reached from:
Step 3 (3 ways) + 1 step
Step 2 (2 ways) + 2 steps
Total: 5 distinct routes ✓
```

### Time Complexity Breakdown:
- **DP Loop**: O(N) iterations
- **Each Iteration**: O(1) operations (addition, assignment)
- **Total**: O(N) time, O(1) space
- **Optimal**: Cannot do better than O(N) for this problem
- **vs Recursion**: O(2^N) time, O(N) space for call stack
*/
