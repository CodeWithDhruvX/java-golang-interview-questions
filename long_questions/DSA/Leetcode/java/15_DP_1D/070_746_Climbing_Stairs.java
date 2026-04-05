import java.util.*;

public class ClimbingStairs {
    
    // 70. Climbing Stairs
    // Time: O(N), Space: O(1)
    public static int climbStairs(int n) {
        if (n <= 0) {
            return 0;
        }
        if (n == 1) {
            return 1;
        }
        if (n == 2) {
            return 2;
        }
        
        int prev2 = 1; // ways to reach step n-2
        int prev1 = 2; // ways to reach step n-1
        int current = 0;
        
        for (int i = 3; i <= n; i++) {
            current = prev1 + prev2;
            prev2 = prev1;
            prev1 = current;
        }
        
        return current;
    }

    // 746. Min Cost Climbing Stairs
    // Time: O(N), Space: O(1)
    public static int minCostClimbingStairs(int[] cost) {
        if (cost == null || cost.length == 0) {
            return 0;
        }
        if (cost.length == 1) {
            return cost[0];
        }
        
        int n = cost.length;
        int[] dp = new int[n + 1];
        dp[0] = 0;
        dp[1] = 0;
        
        for (int i = 2; i <= n; i++) {
            dp[i] = Math.min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);
        }
        
        return dp[n];
    }

    public static void main(String[] args) {
        // Test cases for climbStairs
        int[] testCases1 = {2, 3, 4, 5, 1, 0, 10, 20, 45, 100};
        
        // Test cases for minCostClimbingStairs
        int[][] testCases2 = {
            {10, 15, 20},
            {1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
            {0, 0, 0, 0},
            {1, 2, 3, 4, 5},
            {10, 5, 8, 3, 6, 9, 2},
            {100, 1, 1, 100},
            {1, 2},
            {2, 1},
            {1},
            {0}
        };
        
        System.out.println("Climbing Stairs:");
        for (int i = 0; i < testCases1.length; i++) {
            int n = testCases1[i];
            int result = climbStairs(n);
            System.out.printf("Test Case %d: n=%d -> %d ways\n", 
                i + 1, n, result);
        }
        
        System.out.println("\nMin Cost Climbing Stairs:");
        for (int i = 0; i < testCases2.length; i++) {
            int[] cost = testCases2[i];
            int result = minCostClimbingStairs(cost);
            System.out.printf("Test Case %d: %s -> %d\n", 
                i + 1, Arrays.toString(cost), result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Fibonacci Dynamic Programming
- **Climbing Stairs**: Each step can be reached from n-1 or n-2
- **Fibonacci Sequence**: ways[n] = ways[n-1] + ways[n-2]
- **Space Optimization**: Only need last two values, not full DP array
- **Iterative DP**: Bottom-up approach from base cases

## 2. PROBLEM CHARACTERISTICS
- **Sequential Steps**: Can only move 1 or 2 steps at a time
- **Counting Problem**: Number of distinct ways to reach top
- **Overlapping Subproblems**: ways[n] depends on ways[n-1] and ways[n-2]
- **Optimal Substructure**: Local optimal choices lead to global optimum

## 3. SIMILAR PROBLEMS
- Fibonacci numbers
- House Robber (can't rob adjacent houses)
- Decode Ways (ways to decode message)
- Number of ways to climb stairs with different step sizes

## 4. KEY OBSERVATIONS
- **Base Cases**: f(0)=0, f(1)=1, f(2)=2
- **Recurrence**: f(n) = f(n-1) + f(n-2)
- **Linear Time**: Can compute in O(N) with space optimization
- **Constant Space**: Only need to track last two values

## 5. VARIATIONS & EXTENSIONS
- Different step sizes (1, 2, 3 steps)
- Circular stairs (can wrap around)
- Obstacles on stairs
- Minimum cost climbing stairs

## 6. INTERVIEW INSIGHTS
- Clarify: "Can I take 0 steps? What are the base cases?"
- Edge cases: n=0, n=1, n=2
- Time complexity: O(N) is optimal
- Space complexity: O(1) with optimization, O(N) with DP array

## 7. COMMON MISTAKES
- Not handling base cases correctly
- Using recursion without memoization (exponential time)
- Integer overflow for large n
- Off-by-one errors in recurrence

## 8. OPTIMIZATION STRATEGIES
- Space optimization from O(N) to O(1)
- Matrix exponentiation for O(log N) time
- Precompute Fibonacci numbers
- Use iterative approach instead of recursion

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like climbing a staircase with choices:**
- At each step, you can either take 1 step or 2 steps
- To reach step n, you could have come from n-1 (1 step) or n-2 (2 steps)
- Total ways = ways to reach n-1 + ways to reach n-2
- This is exactly the Fibonacci sequence shifted by 2

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Number of steps n
2. **Goal**: Number of distinct ways to climb to top
3. **Constraints**: Can take 1 or 2 steps at a time
4. **Output**: Total number of ways

#### Phase 2: Key Insight Recognition
- **"How can I reach step n?"** → From n-1 or n-2
- **"What's the pattern?"** → Fibonacci sequence
- **"Can I optimize space?"** → Only need last two values
- **"What are base cases?"** → f(0)=0, f(1)=1, f(2)=2

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use dynamic programming:
1. Base cases: f(0)=0, f(1)=1, f(2)=2
2. For n ≥ 3: f(n) = f(n-1) + f(n-2)
3. Instead of storing all values, just keep last two:
   - prev2 = f(n-2)
   - prev1 = f(n-1)
   - current = prev1 + prev2
4. Update: prev2 = prev1, prev1 = current
5. Continue until n"
```

#### Phase 4: Algorithm Walkthrough
```
Example: n=5

Human thinking:
"Base cases:
f(0)=0, f(1)=1, f(2)=2

Step 3: f(3) = f(2) + f(1) = 2 + 1 = 3
Step 4: f(4) = f(3) + f(2) = 3 + 2 = 5
Step 5: f(5) = f(4) + f(3) = 5 + 3 = 8

So there are 8 ways to climb 5 steps:
1+1+1+1+1, 1+1+1+2, 1+1+2+1, 1+2+1+1,
1+2+2, 2+1+1+1, 2+1+2, 2+2+1 ✓"
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Exponential time without memoization
2. **"What about large n?"** → Use iterative approach to avoid stack overflow
3. **"How to handle base cases?"** → f(0)=0, f(1)=1, f(2)=2
4. **"Can I use matrix multiplication?"** → Yes, for O(log N) time

### Real-World Analogy
**Like planning different routes to reach a destination:**
- Each step is a decision point (take 1 or 2 steps)
- To reach floor n, you could come from floor n-1 or n-2
- Total routes = routes to n-1 + routes to n-2
- This creates a branching pattern like Fibonacci growth

### Human-Readable Pseudocode
```
function climbStairs(n):
    if n <= 0: return 0
    if n == 1: return 1
    if n == 2: return 2
    
    prev2 = 1  // f(1)
    prev1 = 2  // f(2)
    
    for i from 3 to n:
        current = prev1 + prev2  // f(i) = f(i-1) + f(i-2)
        prev2 = prev1
        prev1 = current
    
    return prev1
```

### Execution Visualization

### Example: n=5
```
Step Evolution:
n=0: 0 ways
n=1: 1 way  [1]
n=2: 2 ways  [1+1, 2]
n=3: 3 ways  [1+1+1, 1+2, 2+1]
n=4: 5 ways  [1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2]
n=5: 8 ways  [all combinations above + one more step]

DP Table:
n: 0 1 2 3 4 5
f: 0 1 2 3 5 8

Space-Optimized:
i=3: prev2=1, prev1=2, current=3
i=4: prev2=2, prev1=3, current=5
i=5: prev2=3, prev1=5, current=8

Result: 8 ✓
```

### Key Visualization Points:
- **Fibonacci Pattern**: Each step builds on previous two
- **Space Optimization**: Only need last two values
- **Linear Time**: Single pass from base to n
- **Constant Space**: O(1) with optimization

### Time Complexity Breakdown:
- **Climbing Stairs**: O(N) time, O(1) space
- **Min Cost Climbing Stairs**: O(N) time, O(N) space (or O(1) optimized)
- **Optimal**: Cannot do better than O(N) for this problem
- **Fibonacci**: Can be optimized to O(log N) with matrix exponentiation
*/
