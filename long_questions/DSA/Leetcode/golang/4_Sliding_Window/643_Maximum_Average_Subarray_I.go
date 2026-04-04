package main

import "fmt"

// 643. Maximum Average Subarray I (Fixed Size Sliding Window)
// Time: O(N), Space: O(1)
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0.0
	}
	
	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	
	maxSum := windowSum
	
	// Slide the window
	for i := k; i < len(nums); i++ {
		windowSum = windowSum - nums[i-k] + nums[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	
	return float64(maxSum) / float64(k)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Fixed Size Sliding Window with Sum Tracking
- **Fixed Window**: Window size is exactly k elements
- **Sum Tracking**: Maintain sum of current window elements
- **Sliding Update**: Remove leftmost element, add new rightmost element
- **Average Calculation**: Divide max sum by k for final answer

## 2. PROBLEM CHARACTERISTICS
- **Fixed Window Size**: Window size is constant (k)
- **Average Maximization**: Find window with maximum average
- **Contiguous Elements**: Subarray must be continuous
- **Sum Optimization**: Maximum average corresponds to maximum sum

## 3. SIMILAR PROBLEMS
- Maximum Average Subarray II (variable window with minimum size)
- Subarray with Maximum Sum (Kadane's algorithm)
- Average of Subarray Minimums (LeetCode 636)
- Find Maximized Subarray Sum (LeetCode 1795)

## 4. KEY OBSERVATIONS
- **Average vs Sum**: Maximizing average is same as maximizing sum for fixed k
- **Fixed window advantage**: Can use sliding window efficiently
- **Linear time possible**: Each element enters and leaves window once
- **No negative impact**: Negative numbers don't break the algorithm

## 5. VARIATIONS & EXTENSIONS
- **Variable window size**: Minimum size k instead of exactly k
- **Multiple queries**: Process multiple k values efficiently
- **Maximum sum instead**: Return maximum sum instead of average
- **Streaming input**: Handle very large arrays efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can k be larger than array length? Handle negative numbers?"
- Edge cases: k equals array length, k equals 1, all negative numbers
- Space complexity: O(1) - only pointers and variables
- Time complexity: O(N) - each element visited once

## 7. COMMON MISTAKES
- Not handling case where k > len(nums)
- Using floating point division too early (precision loss)
- Recomputing sum for each window (O(N²))
- Not handling negative numbers correctly
- Forgetting to convert to float at the end

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(1) space
- **Early termination**: Not applicable (need to check all windows)
- **Parallel processing**: Possible with prefix sums
- **Cache optimization**: Sequential memory access pattern

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the best consecutive k-day performance:**
- You have daily performance numbers for a stock
- You want to find the best k-day consecutive period
- The best period is the one with the highest average performance
- Since all periods have the same length (k days), this is the same as finding the period with the highest total
- Slide a k-day window through the data, tracking the total
- Keep the maximum total you find, then divide by k for the average

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers and window size k
2. **Goal**: Find maximum average of any contiguous subarray of size k
3. **Output**: Maximum average as float
4. **Constraint**: Subarray must have exactly k elements

#### Phase 2: Key Insight Recognition
- **"Average vs Sum optimization"** → For fixed k, max average = max sum/k
- **"Fixed window advantage"** → Can use sliding window efficiently
- **"Linear time possibility"** → Each element enters and leaves window once
- **"Precision handling"** → Use integer for sum, convert to float at end

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the maximum average of any k-length subarray.
Since all subarrays have the same length k, maximizing the average
is the same as maximizing the sum.
I'll use a sliding window of size k:
1. Calculate sum of first k elements
2. Slide the window, updating sum by removing leftmost and adding new rightmost
3. Keep track of the maximum sum found
4. Return maxSum/k as float"
```

#### Phase 4: Edge Case Handling
- **k > len(nums)**: Return 0 or handle as invalid input
- **k == len(nums)**: Return average of entire array
- **k == 1**: Return max element as float
- **All negative numbers**: Still works, max sum will be least negative

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums=[1,12,-5,-6,50,3], k=4

Human thinking:
"I need to find the best 4-day period. Let me slide through:

Window 0-3: [1,12,-5,-6]
   Sum = 1+12-5-6 = 2
   Max so far = 2

Slide to window 1-4: [12,-5,-6,50]
   Remove 1, add 50: New sum = 2 - 1 + 50 = 51
   New max = 51

Slide to window 2-5: [-5,-6,50,3]
   Remove 12, add 3: New sum = 51 - 12 + 3 = 42
   Max stays 51

Best sum was 51, so best average = 51/4 = 12.75"
```

#### Phase 6: Intuition Validation
- **Why sliding window works**: Fixed window size enables efficient updates
- **Why sum optimization works**: For fixed k, max average = max sum/k
- **Why O(N) time**: Each element enters and leaves window exactly once
- **Why precision matters**: Use integer arithmetic until final division

### Common Human Pitfalls & How to Avoid Them
1. **"Why not compute average each time?"** → Floating point operations are slower and can lose precision
2. **"Should I handle negative numbers differently?"** → No, algorithm works the same
3. **"Can I use nested loops?"** → That's O(N²), sliding window is O(N)
4. **"What about k > array length?"** → Handle as edge case

### Real-World Analogy
**Like finding the best consecutive k-month sales performance:**
- You have monthly sales data for a company
- You want to find the best k-month consecutive period
- The best period has the highest average monthly sales
- Since all periods are the same length, just find the period with highest total sales
- Slide a k-month window through the data, tracking total sales
- Keep the maximum total, then divide by k for the average

### Human-Readable Pseudocode
```
function findMaxAverage(numbers, k):
    if k > length(numbers) or k == 0:
        return 0.0
    
    # Calculate sum of first window
    windowSum = 0
    for i from 0 to k-1:
        windowSum += numbers[i]
    
    maxSum = windowSum
    
    # Slide window through array
    for i from k to length(numbers)-1:
        # Remove leftmost element, add new element
        windowSum = windowSum - numbers[i-k] + numbers[i]
        maxSum = max(maxSum, windowSum)
    
    return maxSum / k  # Convert to float division
```

### Execution Visualization

### Example: nums=[1,12,-5,-6,50,3], k=4
```
Array: [1][12][-5][-6][50][3]
Index:  0   1   2   3   4   5

Initial window (0-3): [1,12,-5,-6]
→ Sum = 1 + 12 - 5 - 6 = 2
→ MaxSum = 2

Slide to window (1-4): [12,-5,-6,50]
→ Remove 1, add 50: NewSum = 2 - 1 + 50 = 51
→ MaxSum = 51

Slide to window (2-5): [-5,-6,50,3]
→ Remove 12, add 3: NewSum = 51 - 12 + 3 = 42
→ MaxSum stays 51

Final: MaxAverage = 51 / 4 = 12.75
```

### Key Visualization Points:
- **Fixed window size**: Always exactly k elements
- **Sum tracking**: Maintain current window sum efficiently
- **Sliding update**: Remove leftmost, add rightmost in O(1)
- **Max tracking**: Keep maximum sum found

### Memory Layout Visualization:
```
Array: [1][12][-5][-6][50][3]
Index:  0   1   2   3   4   5
        ^  ^  ^  ^
        0  1  2  3  -> window size=4, sum=2
           ^  ^  ^  ^
           1  2  3  4  -> sum=51 (max)
```

### Time Complexity Breakdown:
- **Initial window**: O(k) to sum first k elements
- **Sliding phase**: O(N-k) slides, each O(1)
- **Total**: O(N) where N = len(nums)
- **Space**: O(1) - only pointers and variables

### Alternative Approaches:

#### 1. Brute Force (O(N²))
```go
func findMaxAverage(nums []int, k int) float64 {
    if k > len(nums) || k == 0 {
        return 0.0
    }
    
    maxSum := 0
    for i := 0; i <= len(nums)-k; i++ {
        sum := 0
        for j := i; j < i+k; j++ {
            sum += nums[j]
        }
        maxSum = max(maxSum, sum)
    }
    
    return float64(maxSum) / float64(k)
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time complexity

#### 2. Prefix Sum (O(N) time, O(N) space)
```go
func findMaxAverage(nums []int, k int) float64 {
    if k > len(nums) || k == 0 {
        return 0.0
    }
    
    // Build prefix sum array
    prefix := make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    
    maxSum := 0
    for i := k; i <= len(nums); i++ {
        sum := prefix[i] - prefix[i-k]
        maxSum = max(maxSum, sum)
    }
    
    return float64(maxSum) / float64(k)
}
```
- **Pros**: O(N) time, useful for multiple queries
- **Cons**: O(N) extra space

#### 3. Kadane's-like Approach (Not applicable)
- Kadane's algorithm is for variable-length subarrays
- Fixed window size requires sliding window approach

### Extensions for Interviews:
- **Variable Window Size**: Minimum size k instead of exactly k
- **Multiple Queries**: Process multiple k values efficiently
- **Maximum Sum**: Return maximum sum instead of average
- **Streaming Input**: Handle very large arrays without storing all elements

### Mathematical Insight:
For fixed window size k:
```
max(Average(subarray)) = max(Sum(subarray)/k) = (1/k) * max(Sum(subarray))
```

Since k is constant, maximizing the average is equivalent to maximizing the sum.
*/
func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4},
		{[]int{5}, 1},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{-1, -2, -3, -4, -5}, 3},
		{[]int{0, 0, 0, 0}, 2},
		{[]int{100, 200, 300, 400, 500}, 5},
	}
	
	for i, tc := range testCases {
		result := findMaxAverage(tc.nums, tc.k)
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Max Average: %.4f\n", 
			i+1, tc.nums, tc.k, result)
	}
}
