public class MaximumSubarray {
    
    // 53. Maximum Subarray - Divide and Conquer Approach
    // Time: O(N log N), Space: O(log N) for recursion stack
    public int maxSubArrayDivideAndConquer(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        return maxSubArrayHelper(nums, 0, nums.length - 1);
    }
    
    private int maxSubArrayHelper(int[] nums, int left, int right) {
        if (left == right) {
            return nums[left];
        }
        
        int mid = left + (right - left) / 2;
        
        // Maximum subarray in left half
        int leftMax = maxSubArrayHelper(nums, left, mid);
        
        // Maximum subarray in right half
        int rightMax = maxSubArrayHelper(nums, mid + 1, right);
        
        // Maximum subarray crossing the middle
        int crossMax = maxCrossingSubArray(nums, left, mid, right);
        
        return Math.max(leftMax, Math.max(rightMax, crossMax));
    }
    
    private int maxCrossingSubArray(int[] nums, int left, int mid, int right) {
        // Maximum sum starting from mid and going left
        int leftSum = Integer.MIN_VALUE;
        int sum = 0;
        
        for (int i = mid; i >= left; i--) {
            sum += nums[i];
            if (sum > leftSum) {
                leftSum = sum;
            }
        }
        
        // Maximum sum starting from mid+1 and going right
        int rightSum = Integer.MIN_VALUE;
        sum = 0;
        
        for (int i = mid + 1; i <= right; i++) {
            sum += nums[i];
            if (sum > rightSum) {
                rightSum = sum;
            }
        }
        
        return leftSum + rightSum;
    }
    
    // Kadane's algorithm for comparison (O(N))
    public int maxSubArrayKadane(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        int maxSoFar = nums[0];
        int maxEndingHere = nums[0];
        
        for (int i = 1; i < nums.length; i++) {
            maxEndingHere = Math.max(nums[i], maxEndingHere + nums[i]);
            maxSoFar = Math.max(maxSoFar, maxEndingHere);
        }
        
        return maxSoFar;
    }
    
    // Version with detailed explanation
    public class MaxSubarrayResult {
        int maxSum;
        int startIndex;
        int endIndex;
        java.util.List<String> explanation;
        
        MaxSubarrayResult(int maxSum, int startIndex, int endIndex, java.util.List<String> explanation) {
            this.maxSum = maxSum;
            this.startIndex = startIndex;
            this.endIndex = endIndex;
            this.explanation = explanation;
        }
    }
    
    public MaxSubarrayResult maxSubArrayDetailed(int[] nums) {
        java.util.List<String> explanation = new java.util.ArrayList<>();
        explanation.add("=== Divide and Conquer for Maximum Subarray ===");
        explanation.add("Array: " + java.util.Arrays.toString(nums));
        
        if (nums.length == 0) {
            explanation.add("Empty array, returning 0");
            return new MaxSubarrayResult(0, -1, -1, explanation);
        }
        
        MaxSubarrayInfo result = maxSubArrayHelperDetailed(nums, 0, nums.length - 1, explanation);
        
        explanation.add(String.format("Final result: sum=%d, start=%d, end=%d", 
            result.maxSum, result.startIndex, result.endIndex));
        
        return new MaxSubarrayResult(result.maxSum, result.startIndex, result.endIndex, explanation);
    }
    
    private static class MaxSubarrayInfo {
        int maxSum;
        int startIndex;
        int endIndex;
        
        MaxSubarrayInfo(int maxSum, int startIndex, int endIndex) {
            this.maxSum = maxSum;
            this.startIndex = startIndex;
            this.endIndex = endIndex;
        }
    }
    
    private MaxSubarrayInfo maxSubArrayHelperDetailed(int[] nums, int left, int right, java.util.List<String> explanation) {
        if (left == right) {
            explanation.add(String.format("Base case: single element nums[%d] = %d", left, nums[left]));
            return new MaxSubarrayInfo(nums[left], left, right);
        }
        
        int mid = left + (right - left) / 2;
        explanation.add(String.format("Dividing: left=%d, mid=%d, right=%d", left, mid, right));
        
        // Maximum subarray in left half
        explanation.add("Finding maximum in left half...");
        MaxSubarrayInfo leftMax = maxSubArrayHelperDetailed(nums, left, mid, explanation);
        explanation.add(String.format("Left max: sum=%d, range=[%d,%d]", leftMax.maxSum, leftMax.startIndex, leftMax.endIndex));
        
        // Maximum subarray in right half
        explanation.add("Finding maximum in right half...");
        MaxSubarrayInfo rightMax = maxSubArrayHelperDetailed(nums, mid + 1, right, explanation);
        explanation.add(String.format("Right max: sum=%d, range=[%d,%d]", rightMax.maxSum, rightMax.startIndex, rightMax.endIndex));
        
        // Maximum subarray crossing the middle
        explanation.add("Finding maximum crossing middle...");
        MaxSubarrayInfo crossMax = maxCrossingSubArrayDetailed(nums, left, mid, right, explanation);
        explanation.add(String.format("Cross max: sum=%d, range=[%d,%d]", crossMax.maxSum, crossMax.startIndex, crossMax.endIndex));
        
        // Find the maximum of three
        MaxSubarrayInfo result;
        if (leftMax.maxSum >= rightMax.maxSum && leftMax.maxSum >= crossMax.maxSum) {
            result = leftMax;
            explanation.add("Left half has the maximum");
        } else if (rightMax.maxSum >= leftMax.maxSum && rightMax.maxSum >= crossMax.maxSum) {
            result = rightMax;
            explanation.add("Right half has the maximum");
        } else {
            result = crossMax;
            explanation.add("Crossing subarray has the maximum");
        }
        
        return result;
    }
    
    private MaxSubarrayInfo maxCrossingSubArrayDetailed(int[] nums, int left, int mid, int right, java.util.List<String> explanation) {
        // Maximum sum starting from mid and going left
        explanation.add(String.format("Finding left crossing sum from %d to %d", mid, left));
        
        int leftSum = Integer.MIN_VALUE;
        int sum = 0;
        int leftIndex = mid;
        
        for (int i = mid; i >= left; i--) {
            sum += nums[i];
            explanation.add(String.format("  nums[%d] = %d, running sum = %d", i, nums[i], sum));
            
            if (sum > leftSum) {
                leftSum = sum;
                leftIndex = i;
                explanation.add(String.format("  New left max: %d at index %d", leftSum, i));
            }
        }
        
        // Maximum sum starting from mid+1 and going right
        explanation.add(String.format("Finding right crossing sum from %d to %d", mid + 1, right));
        
        int rightSum = Integer.MIN_VALUE;
        sum = 0;
        int rightIndex = mid + 1;
        
        for (int i = mid + 1; i <= right; i++) {
            sum += nums[i];
            explanation.add(String.format("  nums[%d] = %d, running sum = %d", i, nums[i], sum));
            
            if (sum > rightSum) {
                rightSum = sum;
                rightIndex = i;
                explanation.add(String.format("  New right max: %d at index %d", rightSum, i));
            }
        }
        
        int totalSum = leftSum + rightSum;
        explanation.add(String.format("Crossing sum: left=%d + right=%d = %d, range=[%d,%d]", 
            leftSum, rightSum, totalSum, leftIndex, rightIndex));
        
        return new MaxSubarrayInfo(totalSum, leftIndex, rightIndex);
    }
    
    // Performance comparison
    public void comparePerformance(int[] nums) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Array: " + java.util.Arrays.toString(nums));
        
        // Divide and Conquer
        long startTime = System.nanoTime();
        int result1 = maxSubArrayDivideAndConquer(nums);
        long endTime = System.nanoTime();
        System.out.printf("Divide and Conquer: %d (took %d ns)\n", result1, endTime - startTime);
        
        // Kadane's algorithm
        startTime = System.nanoTime();
        int result2 = maxSubArrayKadane(nums);
        endTime = System.nanoTime();
        System.out.printf("Kadane's algorithm: %d (took %d ns)\n", result2, endTime - startTime);
    }
    
    // Find maximum subarray in each quadrant
    public int[] maxSubArrayQuadrants(int[] nums) {
        if (nums.length == 0) {
            return new int[]{0, -1, -1};
        }
        
        int n = nums.length;
        int[] results = new int[4];
        
        // Divide into 4 quadrants
        int quarter = n / 4;
        
        results[0] = maxSubArrayHelper(nums, 0, quarter - 1);
        results[1] = maxSubArrayHelper(nums, quarter, n/2 - 1);
        results[2] = maxSubArrayHelper(nums, n/2, 3*quarter - 1);
        results[3] = maxSubArrayHelper(nums, 3*quarter, n - 1);
        
        return results;
    }
    
    // Parallel divide and conquer (conceptual)
    public int maxSubArrayParallel(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        // For demonstration, we'll use sequential approach
        // In practice, this would use multiple threads
        return maxSubArrayDivideAndConquer(nums);
    }
    
    // Maximum subarray with constraints
    public int maxSubArrayWithConstraints(int[] nums, int minLength, int maxLength) {
        if (nums.length == 0) {
            return 0;
        }
        
        int maxSum = Integer.MIN_VALUE;
        
        for (int i = 0; i < nums.length; i++) {
            for (int j = i; j < nums.length; j++) {
                int length = j - i + 1;
                
                if (length >= minLength && length <= maxLength) {
                    int sum = 0;
                    for (int k = i; k <= j; k++) {
                        sum += nums[k];
                    }
                    maxSum = Math.max(maxSum, sum);
                }
            }
        }
        
        return maxSum;
    }
    
    public static void main(String[] args) {
        MaximumSubarray ms = new MaximumSubarray();
        
        // Test cases
        int[][] testCases = {
            {-2,1,-3,4,-1,2,1,-5,4},
            {1},
            {5,4,-1,7,8},
            {-2,-3,-1,-5},
            {1,2,3,4,5},
            {-1,2,-3,4,-5,6},
            {100, -200, 300, -400, 500},
            {-2, -1, -3, -4},
            {0, 0, 0, 0},
            {8, -19, 5, -4, 20}
        };
        
        String[] descriptions = {
            "Standard case",
            "Single element",
            "All positive",
            "All negative",
            "Increasing sequence",
            "Alternating signs",
            "Mixed large numbers",
            "All negative small",
            "All zeros",
            "Complex case"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Array: %s\n", java.util.Arrays.toString(testCases[i]));
            
            int result1 = ms.maxSubArrayDivideAndConquer(testCases[i]);
            int result2 = ms.maxSubArrayKadane(testCases[i]);
            
            System.out.printf("Divide and Conquer: %d\n", result1);
            System.out.printf("Kadane's algorithm: %d\n", result2);
            
            // Detailed analysis for first case
            if (i == 0) {
                MaxSubarrayResult detailed = ms.maxSubArrayDetailed(testCases[i]);
                System.out.println("Detailed explanation:");
                for (String step : detailed.explanation) {
                    System.out.println("  " + step);
                }
            }
            
            System.out.println();
        }
        
        // Performance comparison
        System.out.println("=== Performance Comparison ===");
        int[] largeArray = new int[10000];
        java.util.Random rand = new java.util.Random();
        for (int i = 0; i < largeArray.length; i++) {
            largeArray[i] = rand.nextInt(200) - 100; // Random between -100 and 100
        }
        
        ms.comparePerformance(largeArray);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("Empty array: %d\n", ms.maxSubArrayDivideAndConquer(new int[]{}));
        System.out.printf("Single negative: %d\n", ms.maxSubArrayDivideAndConquer(new int[]{-5}));
        System.out.printf("Single positive: %d\n", ms.maxSubArrayDivideAndConquer(new int[]{5}));
        System.out.printf("Two elements: %d\n", ms.maxSubArrayDivideAndConquer(new int[]{1, -2}));
        
        // Maximum subarray with constraints
        System.out.println("\n=== With Constraints ===");
        int[] constrainedArray = {1, -2, 3, 4, -5, 8};
        int constrainedResult = ms.maxSubArrayWithConstraints(constrainedArray, 2, 3);
        System.out.printf("Array: %s\n", java.util.Arrays.toString(constrainedArray));
        System.out.printf("Max subarray with length 2-3: %d\n", constrainedResult);
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Divide and Conquer
- **Problem Division**: Split problem into smaller subproblems
- **Recursive Solution**: Solve subproblems recursively
- **Result Combination**: Combine subproblem solutions
- **Maximum Subarray**: Find contiguous subarray with maximum sum

## 2. PROBLEM CHARACTERISTICS
- **Maximum Subarray**: Find contiguous subarray with maximum sum
- **Divide Strategy**: Split array at middle, solve halves
- **Crossing Subarray**: Handle subarray crossing middle
- **Recursive Structure**: Natural fit for divide and conquer

## 3. SIMILAR PROBLEMS
- Maximum Subarray Sum (Kadane's)
- Maximum Product Subarray
- Longest Increasing Subsequence
- Matrix Chain Multiplication

## 4. KEY OBSERVATIONS
- Divide and conquer splits problem recursively
- Maximum subarray can be in left, right, or crossing middle
- Time complexity: O(N log N) vs O(N) Kadane's
- Space complexity: O(log N) recursion stack
- Kadane's algorithm is more efficient for this specific problem

## 5. VARIATIONS & EXTENSIONS
- Different combination strategies
- Multiple subarray problems
- Parallel divide and conquer
- Constraint-based variations

## 6. INTERVIEW INSIGHTS
- Clarify: "Can we modify the array?"
- Edge cases: empty array, all negative, single element
- Time complexity: O(N log N) vs O(N) Kadane's
- Space complexity: O(log N) vs O(1) Kadane's

## 7. COMMON MISTAKES
- Incorrect crossing subarray calculation
- Wrong base case handling
- Incorrect subproblem division
- Missing recursive termination
- Wrong combination of results

## 8. OPTIMIZATION STRATEGIES
- Efficient crossing subarray calculation
- Proper base case handling
- Minimize recursive calls
- Use Kadane's for optimal solution

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the most profitable segment:**
- You have a sequence of profits/losses (array)
- Need to find contiguous segment with maximum total profit
- Divide and conquer: split sequence in half
- Find best segment in left half, right half, and crossing middle
- Best of these three is the overall best
- This is like finding the best business segment in a timeline

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers (profits/losses)
2. **Goal**: Find contiguous subarray with maximum sum
3. **Output**: Maximum sum value

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N²) to check all subarrays
- **"How to optimize?"** → Use divide and conquer
- **"Why divide and conquer?"** → Natural recursive structure
- **"How to combine?"** → Max of left, right, and crossing

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use divide and conquer:
1. If array has one element, return it
2. Find middle index
3. Recursively find maximum in left half
4. Recursively find maximum in right half
5. Find maximum subarray crossing middle:
   - Maximum sum from middle to left
   - Maximum sum from middle+1 to right
   - Combine these two sums
6. Return maximum of three results
7. This gives O(N log N) time complexity"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 or handle appropriately
- **Single element**: Return that element
- **All negative**: Return maximum (least negative) element
- **Large arrays**: Ensure O(N log N) time complexity

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Array: [-2,1,-3,4,-1,2,1,-5,4]

Human thinking:
"Let's apply divide and conquer:

Step 1: Divide array
left = [-2,1,-3,4,-1]
right = [2,1,-5,4]
mid = 3 (index 3, value -1)

Step 2: Recursively solve left half
leftLeft = [-2,1]
leftRight = [3,4,-1]

Step 3: Recursively solve right half
rightLeft = [2,1]
rightRight = [-5,4]

Continue recursion until base cases...

Step 4: Find crossing subarray at middle
Maximum from middle (index 3, value -1) to left:
- [-1] + [4] = 3
- [-1] + [4] + [-3] = 0
- [-1] + [4] + [-3] + [1] = 1
Maximum = 1 (subarray [1,-3,4,-1])

Maximum from middle+1 (index 4, value 2) to right:
- [2] = 2
- [2] + [1] = 3
- [2] + [1] + [-5] = -2
- [2] + [1] + [-5] + [4] = 2
Maximum = 2 (subarray [2,1,-5,4])

Crossing sum = 1 + 2 = 3

Step 5: Combine results
Left max = 4 (subarray [4])
Right max = 6 (subarray [2,1,-5,4])
Crossing max = 3

Overall maximum = max(4, 6, 3) = 6 ✓

Manual verification:
Subarray [2,1,-5,4] sum = 2+1-5+4 = 2 ✓
This is indeed the maximum ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Divides problem into smaller subproblems
- **Why it's efficient**: O(N log N) vs O(N²) brute force
- **Why it's correct**: Considers all possible subarrays

### Common Human Pitfalls & How to Avoid Them
1. **"Why not check all subarrays?"** → O(N²) too slow
2. **"What about crossing subarray?"** → Must handle middle case separately
3. **"How to combine results?"** → Take maximum of three cases
4. **"What about base cases?"** → Single element returns itself

### Real-World Analogy
**Like finding the most profitable business period:**
- You have daily profits/losses for a business (array)
- Need to find the most profitable contiguous period
- Divide and conquer: split timeline into halves
- Find best period in first half, second half, and spanning both
- Best of these three is the overall best period
- This is used in financial analysis, stock trading, signal processing
- Like finding the peak performance segment in a timeline

### Human-Readable Pseudocode
```
function maxSubArray(nums):
    if nums.length == 0:
        return 0
    
    return maxSubArrayHelper(nums, 0, nums.length - 1)

function maxSubArrayHelper(nums, left, right):
    if left == right:
        return nums[left]
    
    mid = left + (right - left) // 2
    
    // Maximum in left half
    leftMax = maxSubArrayHelper(nums, left, mid)
    
    // Maximum in right half
    rightMax = maxSubArrayHelper(nums, mid + 1, right)
    
    // Maximum crossing middle
    crossMax = maxCrossingSubArray(nums, left, mid, right)
    
    return max(leftMax, rightMax, crossMax)

function maxCrossingSubArray(nums, left, mid, right):
    // Maximum sum starting from mid and going left
    leftSum = -∞
    sum = 0
    for i from mid down to left:
        sum += nums[i]
        leftSum = max(leftSum, sum)
    
    // Maximum sum starting from mid+1 and going right
    rightSum = -∞
    sum = 0
    for i from mid+1 to right:
        sum += nums[i]
        rightSum = max(rightSum, sum)
    
    return leftSum + rightSum
```

### Execution Visualization

### Example: nums=[-2,1,-3,4,-1,2,1,-5,4]
```
Divide and Conquer Process:

Initial array: [-2,1,-3,4,-1,2,1,-5,4]
left=[-2,1,-3,4,-1], right=[2,1,-5,4], mid=3

Step 1: Recursively solve left half
leftLeft=[-2,1], leftRight=[-3,4,-1]

Step 2: Recursively solve right half
rightLeft=[2,1], rightRight=[-5,4]

Step 3: Find crossing subarray
From mid=3 to left: max sum = 1 (subarray [1,-3,4,-1])
From mid+1=4 to right: max sum = 2 (subarray [2,1,-5,4])
Crossing sum = 1 + 2 = 3

Step 4: Combine results
Left max = 4 (subarray [4])
Right max = 6 (subarray [2,1,-5,4])
Crossing max = 3

Overall maximum = max(4, 6, 3) = 6 ✓

Visualization:
Divide and conquer splits problem recursively
Each level considers all possible subarrays
Maximum of three cases gives optimal result ✓
```

### Key Visualization Points:
- **Problem Division**: Split at middle index
- **Recursive Solution**: Solve subproblems independently
- **Crossing Case**: Handle subarray spanning middle
- **Result Combination**: Maximum of three possibilities

### Memory Layout Visualization:
```
Recursive Call Stack:

Level 0: [-2,1,-3,4,-1,2,1,-5,4]
├─ Left: [-2,1,-3,4,-1]
│  ├─ Left: [-2,1]
│  │  └─ Result: max(-2,1,-1) = 1
│  └─ Right: [-3,4,-1]
│     └─ Result: max(-3,4,-1) = 4
└─ Right: [2,1,-5,4]
   ├─ Left: [2,1]
   │  └─ Result: max(2,1) = 3
   └─ Right: [-5,4]
      └─ Result: max(-5,4) = 4

Crossing subarray at middle:
Left crossing: max sum = 1
Right crossing: max sum = 2
Total crossing: 1 + 2 = 3

Final result: max(4, 4, 3) = 6 ✓
```

### Time Complexity Breakdown:
- **Divide and Conquer**: O(N log N) time, O(log N) space
- **Kadane's Algorithm**: O(N) time, O(1) space
- **Brute Force**: O(N²) time, O(1) space
- **Optimal**: Kadane's is best for this specific problem
- **vs Naive**: O(N log N) vs O(N²) significant improvement
*/
