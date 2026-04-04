package main

import (
	"fmt"
	"math"
)

// 53. Maximum Subarray - Divide and Conquer Approach
// Time: O(N log N), Space: O(log N) for recursion stack
func maxSubArrayDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	
	mid := left + (right-left)/2
	
	// Maximum subarray in left half
	leftMax := maxSubArrayHelper(nums, left, mid)
	
	// Maximum subarray in right half
	rightMax := maxSubArrayHelper(nums, mid+1, right)
	
	// Maximum subarray crossing the middle
	crossMax := maxCrossingSubArray(nums, left, mid, right)
	
	return max(leftMax, rightMax, crossMax)
}

func maxCrossingSubArray(nums []int, left, mid, right int) int {
	// Maximum sum starting from mid and going left
	leftSum := math.MinInt32
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	
	// Maximum sum starting from mid+1 and going right
	rightSum := math.MinInt32
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	
	return leftSum + rightSum
}

func max(a, b, c int) int {
	maxVal := a
	if b > maxVal {
		maxVal = b
	}
	if c > maxVal {
		maxVal = c
	}
	return maxVal
}

// Kadane's algorithm for comparison (O(N))
func maxSubArrayKadane(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	maxSoFar := nums[0]
	maxEndingHere := nums[0]
	
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Divide and Conquer with memoization
func maxSubArrayDivideAndConquerMemo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	memo := make(map[string]int)
	return maxSubArrayHelperMemo(nums, 0, len(nums)-1, memo)
}

func maxSubArrayHelperMemo(nums []int, left, right int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", left, right)
	if val, exists := memo[key]; exists {
		return val
	}
	
	if left == right {
		memo[key] = nums[left]
		return nums[left]
	}
	
	mid := left + (right-left)/2
	
	leftMax := maxSubArrayHelperMemo(nums, left, mid, memo)
	rightMax := maxSubArrayHelperMemo(nums, mid+1, right, memo)
	crossMax := maxCrossingSubArray(nums, left, mid, right)
	
	result := max(leftMax, rightMax, crossMax)
	memo[key] = result
	return result
}

// Divide and Conquer with tracking indices
func maxSubArrayDivideAndConquerWithIndices(nums []int) (int, int, int) {
	if len(nums) == 0 {
		return 0, -1, -1
	}
	
	return maxSubArrayHelperWithIndices(nums, 0, len(nums)-1)
}

func maxSubArrayHelperWithIndices(nums []int, left, right int) (int, int, int) {
	if left == right {
		return nums[left], left, right
	}
	
	mid := left + (right-left)/2
	
	leftSum, leftStart, leftEnd := maxSubArrayHelperWithIndices(nums, left, mid)
	rightSum, rightStart, rightEnd := maxSubArrayHelperWithIndices(nums, mid+1, right)
	crossSum, crossStart, crossEnd := maxCrossingSubArrayWithIndices(nums, left, mid, right)
	
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftSum, leftStart, leftEnd
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightSum, rightStart, rightEnd
	} else {
		return crossSum, crossStart, crossEnd
	}
}

func maxCrossingSubArrayWithIndices(nums []int, left, mid, right int) (int, int, int) {
	// Maximum sum starting from mid and going left
	leftSum := math.MinInt32
	sum := 0
	leftIdx := mid
	
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
			leftIdx = i
		}
	}
	
	// Maximum sum starting from mid+1 and going right
	rightSum := math.MinInt32
	sum = 0
	rightIdx := mid + 1
	
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
			rightIdx = i
		}
	}
	
	return leftSum + rightSum, leftIdx, rightIdx
}

// Divide and Conquer for maximum product subarray
func maxProductDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	return maxProductHelper(nums, 0, len(nums)-1)
}

func maxProductHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	
	mid := left + (right-left)/2
	
	// Maximum product in left half
	leftMax := maxProductHelper(nums, left, mid)
	
	// Maximum product in right half
	rightMax := maxProductHelper(nums, mid+1, right)
	
	// Maximum product crossing the middle
	crossMax := maxCrossingProduct(nums, left, mid, right)
	
	return max(leftMax, rightMax, crossMax)
}

func maxCrossingProduct(nums []int, left, mid, right int) int {
	// For product, we need to consider both positive and negative values
	// Maximum product starting from mid and going left
	maxLeftProd := math.MinInt32
	minLeftProd := math.MaxInt32
	prod := 1
	
	for i := mid; i >= left; i-- {
		prod *= nums[i]
		if prod > maxLeftProd {
			maxLeftProd = prod
		}
		if prod < minLeftProd {
			minLeftProd = prod
		}
	}
	
	// Maximum product starting from mid+1 and going right
	maxRightProd := math.MinInt32
	minRightProd := math.MaxInt32
	prod = 1
	
	for i := mid + 1; i <= right; i++ {
		prod *= nums[i]
		if prod > maxRightProd {
			maxRightProd = prod
		}
		if prod < minRightProd {
			minRightProd = prod
		}
	}
	
	// Consider all combinations
	cross1 := maxLeftProd * maxRightProd
	cross2 := maxLeftProd * minRightProd
	cross3 := minLeftProd * maxRightProd
	cross4 := minLeftProd * minRightProd
	
	return max(cross1, cross2, cross3, cross4)
}

// Divide and Conquer for circular subarray
func maxSubarrayCircularDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// Case 1: Maximum subarray is non-circular (standard Kadane)
	maxKadane := maxSubArrayKadane(nums)
	
	// Case 2: Maximum subarray is circular
	if len(nums) == 1 {
		return maxKadane
	}
	
	// Calculate total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	
	// Find minimum subarray (Kadane on inverted array)
	minKadane := minSubArray(nums)
	
	maxCircular := totalSum - minKadane
	
	if maxCircular == 0 && maxKadane < 0 {
		return maxKadane // All negative numbers
	}
	
	return max(maxKadane, maxCircular)
}

func minSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	minSoFar := nums[0]
	minEndingHere := nums[0]
	
	for i := 1; i < len(nums); i++ {
		minEndingHere = min(nums[i], minEndingHere+nums[i])
		minSoFar = min(minSoFar, minEndingHere)
	}
	
	return minSoFar
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Divide and Conquer for Maximum Subarray
- **Recursive Division**: Split array into halves recursively
- **Crossing Subarray**: Handle subarrays that span the middle
- **Combine Results**: Take maximum of left, right, and crossing subarrays
- **Conquer Strategy**: Solve smaller subproblems and combine solutions

## 2. PROBLEM CHARACTERISTICS
- **Maximum Subarray**: Find contiguous subarray with maximum sum
- **Contiguous Constraint**: Elements must be consecutive in array
- **Optimal Substructure**: Solution can be built from subproblem solutions
- **Overlap Property**: Maximum subarray can cross the divide point

## 3. SIMILAR PROBLEMS
- Maximum Subarray (LeetCode 53) - Same problem
- Maximum Product Subarray - Similar with multiplication
- Circular Subarray Maximum - Wrap-around consideration
- Range Sum Queries - Prefix sum applications

## 4. KEY OBSERVATIONS
- **Divide Natural**: Array can be split at midpoint
- **Three Cases**: Maximum in left half, right half, or crossing middle
- **Crossing Complexity**: Need O(N) time to compute crossing subarray
- **Recursive Structure**: Same problem applied to smaller arrays

## 5. VARIATIONS & EXTENSIONS
- **Standard Divide and Conquer**: O(N log N) time, O(log N) space
- **Kadane's Algorithm**: O(N) time, O(1) space - optimal
- **Circular Subarray**: Handle wrap-around cases
- **Maximum Product**: Similar approach with sign considerations

## 6. INTERVIEW INSIGHTS
- Always clarify: "Array size? Negative numbers? Circular allowed?"
- Edge cases: empty array, all negative, single element
- Time complexity: O(N log N) for D&C, O(N) for Kadane's
- Space complexity: O(log N) recursion stack, O(1) for Kadane's
- Key insight: Kadane's algorithm is more efficient but D&C demonstrates pattern

## 7. COMMON MISTAKES
- Wrong crossing subarray calculation (must include elements from both sides)
- Missing base case for single element
- Incorrect recursive boundaries (off-by-one errors)
- Not handling all negative arrays properly
- Wrong maximum of three values logic

## 8. OPTIMIZATION STRATEGIES
- **Standard D&C**: O(N log N) time, O(log N) space - demonstrates pattern
- **Kadane's**: O(N) time, O(1) space - optimal solution
- **Memoization**: O(N^2) time, O(N^2) space - reduces redundant calculations
- **Index Tracking**: O(N log N) time, O(log N) space - with subarray indices

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the most profitable segment of a journey:**
- You have a journey with different profit/loss segments
- You want to find the most profitable contiguous segment
- You can split the journey in half and analyze each part
- The best segment might be entirely in left half, right half, or spanning both
- Like a business analyst finding the most profitable quarter by analyzing halves

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers (can be positive or negative)
2. **Goal**: Find contiguous subarray with maximum sum
3. **Constraints**: Must be consecutive elements, can be empty
4. **Output**: Maximum sum value (and optionally indices)

#### Phase 2: Key Insight Recognition
- **"Divide natural"** → Can split array at midpoint
- **"Three possibilities"** → Max in left, right, or crossing middle
- **"Recursive structure"** → Same problem on smaller arrays
- **"Crossing special"** → Need special handling for middle-spanning subarrays

#### Phase 3: Strategy Development
```
Human thought process:
"I need maximum contiguous subarray.
Brute force: check all O(N²) subarrays.

Divide and Conquer Approach:
1. Split array at midpoint
2. Recursively find max in left half
3. Recursively find max in right half
4. Find max subarray crossing middle (O(N))
5. Return maximum of three results

This gives O(N log N) time, O(log N) space!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 (or handle as specified)
- **Single element**: Return that element
- **All negative**: Return least negative element
- **Large arrays**: Recursion depth O(log N)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]

Human thinking:
"Divide and Conquer Process:
Step 1: Split at middle (index 4)
Left: [-2, 1, -3, 4], Right: [-1, 2, 1, -5, 4]

Step 2: Recursively solve left half
Split left at middle (index 1)
LeftLeft: [-2, 1], LeftRight: [-3, 4]

Step 3: Solve LeftLeft [-2, 1]
Split at middle (index 0)
Base case: single elements
Max(-2) = -2, Max(1) = 1
Crossing: -2 + 1 = -1
Result: max(-2, 1, -1) = 1

Step 4: Solve LeftRight [-3, 4]
Base case: single elements
Max(-3) = -3, Max(4) = 4
Crossing: -3 + 4 = 1
Result: max(-3, 4, 1) = 4

Step 5: Solve Left [-2, 1, -3, 4]
Left max = 1, Right max = 4
Crossing: max from middle left = 1, max from middle right = 4
Crossing sum = 1 + 4 = 5
Result: max(1, 4, 5) = 5

Continue similarly for right half...
Final result: 6 (subarray [4, -1, 2, 1]) ✓"
```

#### Phase 6: Intuition Validation
- **Why divide**: Natural way to break down problem
- **Why three cases**: Maximum must be in one of three regions
- **Why crossing special**: Spans both halves, needs O(N) computation
- **Why recursive**: Same pattern applies to subarrays

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use Kadane's?"** → D&C demonstrates divide pattern, Kadane's is optimal
2. **"Should I check all subarrays?"** → O(N²) vs O(N log N), too slow
3. **"What about crossing subarray?"** → Must include elements from both sides of middle
4. **"Can I use memoization?"** → Possible but still O(N log N) best case
5. **"Why O(N log N)?"** → Each level does O(N) work, O(log N) levels

### Real-World Analogy
**Like finding the most profitable period in stock trading:**
- You have daily stock price changes (profits/losses)
- You want to find the most profitable consecutive period
- You can analyze the year by splitting into halves
- The best period might be entirely in first half, second half, or spanning both
- Like a financial analyst finding the best investment window by analyzing quarters

### Human-Readable Pseudocode
```
function maxSubArray(nums):
    if len(nums) == 0:
        return 0
    
    return maxSubArrayHelper(nums, 0, len(nums) - 1)

function maxSubArrayHelper(nums, left, right):
    if left == right:
        return nums[left]
    
    mid = (left + right) // 2
    
    # Maximum in left half
    leftMax = maxSubArrayHelper(nums, left, mid)
    
    # Maximum in right half
    rightMax = maxSubArrayHelper(nums, mid + 1, right)
    
    # Maximum crossing middle
    crossMax = maxCrossingSubArray(nums, left, mid, right)
    
    return max(leftMax, rightMax, crossMax)

function maxCrossingSubArray(nums, left, mid, right):
    # Max sum from mid going left
    leftSum = -infinity
    sum = 0
    for i from mid down to left:
        sum += nums[i]
        leftSum = max(leftSum, sum)
    
    # Max sum from mid+1 going right
    rightSum = -infinity
    sum = 0
    for i from mid+1 to right:
        sum += nums[i]
        rightSum = max(rightSum, sum)
    
    return leftSum + rightSum
```

### Execution Visualization

### Example: nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
```
Level 0: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
         Split at index 4
    Left: [-2, 1, -3, 4]    Right: [-1, 2, 1, -5, 4]

Level 1: Left: [-2, 1, -3, 4]
         Split at index 1
    LeftLeft: [-2, 1]    LeftRight: [-3, 4]

Level 2: LeftLeft: [-2, 1]
         Split at index 0
    [-2]    [1]
    Max: -2, Max: 1
    Crossing: -2 + 1 = -1
    Result: max(-2, 1, -1) = 1

Level 2: LeftRight: [-3, 4]
         Split at index 2
    [-3]    [4]
    Max: -3, Max: 4
    Crossing: -3 + 4 = 1
    Result: max(-3, 4, 1) = 4

Level 1: Left: [-2, 1, -3, 4]
    Left max = 1, Right max = 4
    Crossing: max left from middle = 1, max right from middle = 4
    Crossing sum = 1 + 4 = 5
    Result: max(1, 4, 5) = 5

Continue similar process for right half...
Final result: 6 ✓
```

### Key Visualization Points:
- **Recursive Splitting**: Each level splits array at midpoint
- **Three Cases**: Left max, right max, crossing max
- **Crossing Calculation**: Expand from middle to both sides
- **Combine Results**: Take maximum of three possibilities

### Divide and Conquer Tree Visualization:
```
        [0..8]
       /      \
    [0..4]    [5..8]
    /    \    /    \
 [0..2][3..4][5..6][7..8]
  /   \  /   \  /   \  /   \
[0][1][2][3][4][5][6][7][8]
```

### Time Complexity Breakdown:
- **Standard D&C**: O(N log N) time, O(log N) space - demonstrates pattern
- **Kadane's**: O(N) time, O(1) space - optimal solution
- **Memoization**: O(N^2) time, O(N^2) space - reduces redundancy
- **Index Tracking**: O(N log N) time, O(log N) space - with indices

### Alternative Approaches:

#### 1. Kadane's Algorithm (O(N) time, O(1) space)
```go
func maxSubArrayKadane(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    maxSoFar := nums[0]
    maxEndingHere := nums[0]
    
    for i := 1; i < len(nums); i++ {
        maxEndingHere = max(nums[i], maxEndingHere + nums[i])
        maxSoFar = max(maxSoFar, maxEndingHere)
    }
    
    return maxSoFar
}
```
- **Pros**: Optimal O(N) time, O(1) space
- **Cons**: Doesn't demonstrate divide and conquer pattern

#### 2. Prefix Sum Approach (O(N²) time, O(N) space)
```go
func maxSubArrayPrefixSum(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    prefix := make([]int, len(nums)+1)
    for i := 1; i <= len(nums); i++ {
        prefix[i] = prefix[i-1] + nums[i-1]
    }
    
    maxSum := math.MinInt32
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            sum := prefix[j+1] - prefix[i]
            maxSum = max(maxSum, sum)
        }
    }
    
    return maxSum
}
```
- **Pros**: Simple concept, easy to understand
- **Cons**: O(N²) time, too slow for large arrays

#### 3. Dynamic Programming (O(N) time, O(N) space)
```go
func maxSubArrayDP(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    maxSum := nums[0]
    
    for i := 1; i < len(nums); i++ {
        dp[i] = max(nums[i], dp[i-1] + nums[i])
        maxSum = max(maxSum, dp[i])
    }
    
    return maxSum
}
```
- **Pros**: Clear DP formulation
- **Cons**: O(N) space vs O(1) for Kadane's

### Extensions for Interviews:
- **Circular Subarray**: Handle wrap-around cases
- **Maximum Product**: Similar approach with sign considerations
- **Index Tracking**: Return start and end indices
- **Multiple Queries**: Range maximum subarray queries
- **Real-world Applications**: Stock analysis, signal processing, data compression
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Maximum Subarray - Divide and Conquer ===")
	
	testCases := []struct {
		nums       []int
		description string
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, "Standard case"},
		{[]int{1}, "Single element"},
		{[]int{5, 4, -1, 7, 8}, "All positive"},
		{[]int{-1, -2, -3, -4}, "All negative"},
		{[]int{0, 0, 0, 0}, "All zeros"},
		{[]int{-2, -1, -2, -3, -1, -4}, "Mixed negatives"},
		{[]int{2, 3, -2, 4}, "Simple positive"},
		{[]int{-1, 2, 3, -5, 4, 6, -1, 2, -3}, "Complex case"},
		{[]int{100, -1, 100, -1, 100}, "Large positives"},
		{[]int{-2, -3, 4, -1, -2, 1, 5, -3}, "Classic example"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Array: %v\n", tc.nums)
		
		result1 := maxSubArrayDivideAndConquer(tc.nums)
		result2 := maxSubArrayKadane(tc.nums)
		result3 := maxSubArrayDivideAndConquerMemo(tc.nums)
		
		sum, start, end := maxSubArrayDivideAndConquerWithIndices(tc.nums)
		
		fmt.Printf("  Divide & Conquer: %d\n", result1)
		fmt.Printf("  Kadane's: %d\n", result2)
		fmt.Printf("  With Memoization: %d\n", result3)
		fmt.Printf("  With Indices: sum=%d, start=%d, end=%d\n", sum, start, end)
		
		// Test circular version
		circularResult := maxSubarrayCircularDivideAndConquer(tc.nums)
		fmt.Printf("  Circular: %d\n\n", circularResult)
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	largeArray := make([]int, 10000)
	for i := range largeArray {
		if i%3 == 0 {
			largeArray[i] = -1
		} else {
			largeArray[i] = 1
		}
	}
	
	fmt.Printf("Large array test with %d elements\n", len(largeArray))
	
	result := maxSubArrayDivideAndConquer(largeArray)
	fmt.Printf("Divide & Conquer result: %d\n", result)
	
	result = maxSubArrayKadane(largeArray)
	fmt.Printf("Kadane's result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Empty array
	fmt.Printf("Empty array: %d\n", maxSubArrayDivideAndConquer([]int{}))
	
	// Single negative
	fmt.Printf("Single negative: %d\n", maxSubArrayDivideAndConquer([]int{-5}))
	
	// Single positive
	fmt.Printf("Single positive: %d\n", maxSubArrayDivideAndConquer([]int{5}))
	
	// Alternating pattern
	alternating := []int{1, -1, 1, -1, 1, -1, 1, -1}
	fmt.Printf("Alternating: %d\n", maxSubArrayDivideAndConquer(alternating))
	
	// Large values
	largeVals := []int{1000000, -1000000, 1000000, -1000000, 1000000}
	fmt.Printf("Large values: %d\n", maxSubArrayDivideAndConquer(largeVals))
	
	// Test maximum product
	fmt.Println("\n=== Maximum Product Test ===")
	productTest := []int{2, 3, -2, 4}
	fmt.Printf("Max product: %d\n", maxProductDivideAndConquer(productTest))
	
	productTest2 := []int{-2, 0, -1}
	fmt.Printf("Max product with zeros: %d\n", maxProductDivideAndConquer(productTest2))
}
