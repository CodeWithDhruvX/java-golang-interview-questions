package main

import "fmt"

// 209. Minimum Size Subarray Sum (Variable Size Sliding Window)
// Time: O(N), Space: O(1)
func minSubArrayLen(target int, nums []int) int {
	left := 0
	currentSum := 0
	minLength := len(nums) + 1 // Initialize with a value larger than any possible result
	
	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]
		
		// Shrink the window from the left as much as possible
		for currentSum >= target {
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
			}
			currentSum -= nums[left]
			left++
		}
	}
	
	if minLength == len(nums)+1 {
		return 0 // No subarray found
	}
	
	return minLength
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Variable Size Sliding Window with Sum Tracking
- **Left Pointer**: Window start, moves right when sum >= target
- **Right Pointer**: Window end, expands window by adding elements
- **Current Sum**: Track sum of elements in current window
- **Window Contraction**: Shrink from left while sum >= target

## 2. PROBLEM CHARACTERISTICS
- **Positive Numbers**: All numbers are positive (crucial assumption)
- **Sum Threshold**: Find subarray with sum >= target
- **Minimum Length**: Want shortest valid subarray
- **Contiguous Elements**: Subarray must be continuous

## 3. SIMILAR PROBLEMS
- Maximum Size Subarray Sum (LeetCode 1835)
- Subarray Product Less Than K (LeetCode 713)
- Number of Subarrays with Sum >= K (LeetCode 862)
- Longest Subarray with Sum at Most K

## 4. KEY OBSERVATIONS
- **Positive numbers enable sliding window**: Adding elements increases sum
- **Monotonic property**: Removing elements from left decreases sum
- **Greedy contraction**: Always try to shrink when sum >= target
- **No negative numbers**: Ensures sliding window works correctly

## 5. VARIATIONS & EXTENSIONS
- **Negative numbers**: Requires different approach (prefix sums + hash map)
- **Exact sum target**: Find subarray with sum exactly equal to target
- **Maximum length**: Find longest subarray with sum >= target
- **Multiple queries**: Process multiple targets efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are numbers positive? Can they be zero or negative?"
- Edge cases: empty array, single element, no valid subarray
- Space complexity: O(1) - only pointers and variables
- Time complexity: O(N) - each element visited at most twice

## 7. COMMON MISTAKES
- Not handling case where no subarray meets requirement
- Forgetting to update sum when shrinking window
- Using nested loops instead of sliding window (O(N²))
- Not handling zero or negative numbers correctly
- Off-by-one errors in length calculation

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(1) space
- **Early termination**: Can stop if remaining elements insufficient
- **Parallel processing**: Not applicable due to window dependency
- **Cache optimization**: Sequential memory access pattern

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the shortest sequence of boxes whose total weight meets a target:**
- You have boxes in a line, each with a positive weight
- You need the shortest consecutive sequence that weighs at least target
- Start adding boxes from the left until you meet the target
- Then try to remove boxes from the left while still meeting the target
- Keep track of the shortest sequence that worked
- Continue this process through the entire line

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of positive integers and target sum
2. **Goal**: Find minimum length of contiguous subarray with sum >= target
3. **Output**: Minimum length (0 if no such subarray exists)
4. **Constraint**: All numbers are positive (enables sliding window)

#### Phase 2: Key Insight Recognition
- **"Positive number advantage"** → Adding elements only increases sum
- **"Window contraction"** → Can safely remove from left when sum >= target
- **"Greedy optimization"** → Always try to shrink when possible
- **"Linear time possibility"** → Each element enters and leaves window once

#### Phase 3: Strategy Development
```
Human thought process:
"I'll maintain a window [left, right] and its sum.
I'll expand the window by moving right and adding elements.
When the sum >= target, I'll try to shrink from the left
to make the window as small as possible while still >= target.
I'll keep track of the minimum length I find."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 (no subarray possible)
- **Single element**: Return 1 if element >= target, else 0
- **No valid subarray**: Return 0 if sum of all elements < target
- **Target = 0**: Return 1 (any single element >= 0)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums=[2,3,1,2,4,3], target=7

Human thinking:
"I'll expand my window and track the sum:

left=0, right=0, sum=2 < 7 → expand
left=0, right=1, sum=5 < 7 → expand  
left=0, right=2, sum=6 < 7 → expand
left=0, right=3, sum=8 >= 7 → try shrinking
           length=4, min=4
           Remove nums[0]=2, sum=6 < 7, stop shrinking

Continue expanding:
left=1, right=4, sum=6+4=10 >= 7 → try shrinking
           length=4, min stays 4
           Remove nums[1]=3, sum=7 >= 7, length=3, min=3
           Remove nums[2]=1, sum=6 < 7, stop shrinking

Continue expanding:
left=3, right=5, sum=6+3=9 >= 7 → try shrinking
           length=3, min stays 3
           Remove nums[3]=2, sum=7 >= 7, length=2, min=2
           Remove nums[4]=4, sum=3 < 7, stop shrinking

Final answer: 2"
```

#### Phase 6: Intuition Validation
- **Why sliding window works**: Positive numbers ensure monotonic sum behavior
- **Why contraction works**: Removing from left always decreases sum
- **Why O(N) time**: Each element added and removed at most once
- **Why positive constraint matters**: Negative numbers break monotonic property

### Common Human Pitfalls & How to Avoid Them
1. **"What if numbers can be negative?"** → Sliding window won't work, need different approach
2. **"Should I check all subarrays?"** → That's O(N²), sliding window is O(N)
3. **"Why shrink when sum >= target?"** → To find minimum length
4. **"Can I use prefix sums?"** → Possible but sliding window is simpler for positive numbers

### Real-World Analogy
**Like finding the shortest sequence of payments that reaches a financial goal:**
- You have daily payments (positive amounts) in chronological order
- You need the shortest consecutive period where total payments >= target amount
- Start adding daily payments until you reach the goal
- Then try removing early payments while still meeting the goal
- Keep track of the shortest period that worked
- Continue this process through all payments

### Human-Readable Pseudocode
```
function minSubArrayLength(target, numbers):
    left = 0
    currentSum = 0
    minLength = infinity
    
    for right from 0 to length(numbers)-1:
        currentSum += numbers[right]
        
        # Try to shrink window while sum >= target
        while currentSum >= target:
            windowLength = right - left + 1
            minLength = min(minLength, windowLength)
            currentSum -= numbers[left]
            left += 1
    
    if minLength == infinity:
        return 0  # No valid subarray found
    return minLength
```

### Execution Visualization

### Example: nums=[2,3,1,2,4,3], target=7
```
Array: [2][3][1][2][4][3]
Index:  0  1  2  3  4  5

Initial: left=0, sum=0, minLen=∞

Step 1: right=0, num=2
→ sum=2 < 7, no shrinking
State: left=0, sum=2, minLen=∞

Step 2: right=1, num=3  
→ sum=5 < 7, no shrinking
State: left=0, sum=5, minLen=∞

Step 3: right=2, num=1
→ sum=6 < 7, no shrinking
State: left=0, sum=6, minLen=∞

Step 4: right=3, num=2
→ sum=8 >= 7, try shrinking
   length=4, minLen=4
   remove nums[0]=2, sum=6 < 7, stop
State: left=1, sum=6, minLen=4

Step 5: right=4, num=4
→ sum=10 >= 7, try shrinking
   length=4, minLen=4
   remove nums[1]=3, sum=7 >= 7, length=3, minLen=3
   remove nums[2]=1, sum=6 < 7, stop
State: left=3, sum=6, minLen=3

Step 6: right=5, num=3
→ sum=9 >= 7, try shrinking
   length=3, minLen=3
   remove nums[3]=2, sum=7 >= 7, length=2, minLen=2
   remove nums[4]=4, sum=3 < 7, stop
State: left=5, sum=3, minLen=2

Final: minLen=2
```

### Key Visualization Points:
- **Window expansion**: Always move right, add element to sum
- **Window contraction**: When sum >= target, try to shrink from left
- **Sum tracking**: Maintain current sum of window elements
- **Length calculation**: right - left + 1

### Memory Layout Visualization:
```
Array: [2][3][1][2][4][3]
Index:  0  1  2  3  4  5
        ^           ^
     left=0      right=3
        sum=8 >= 7, try shrinking
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited at most twice
- **Constant Space**: O(1) - only pointers and variables
- **Window Operations**: Each element enters and leaves window once
- **Sum Updates**: O(1) per operation

### Alternative Approaches:

#### 1. Brute Force (O(N²))
```go
func minSubArrayLen(target int, nums []int) int {
    minLength := len(nums) + 1
    
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum >= target {
                minLength = min(minLength, j-i+1)
                break
            }
        }
    }
    
    if minLength == len(nums)+1 {
        return 0
    }
    return minLength
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time complexity

#### 2. Prefix Sum + Binary Search (O(N log N))
```go
func minSubArrayLen(target int, nums []int) int {
    // Build prefix sums
    prefix := make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    
    minLength := len(nums) + 1
    
    for i := 1; i <= len(nums); i++ {
        // Find smallest j where prefix[j] - prefix[i-1] >= target
        targetPrefix := prefix[i-1] + target
        j := lowerBound(prefix, targetPrefix)
        if j <= len(nums) {
            minLength = min(minLength, j - (i-1))
        }
    }
    
    if minLength == len(nums)+1 {
        return 0
    }
    return minLength
}
```
- **Pros**: Works with negative numbers too
- **Cons**: O(N log N) time, more complex

#### 3. Handle Negative Numbers (O(N log N))
```go
func minSubArrayLen(target int, nums []int) int {
    // Use prefix sums + hash map for negative numbers
    prefix := 0
    prefixMap := make(map[int]int)
    prefixMap[0] = -1
    minLength := len(nums) + 1
    
    for i, num := range nums {
        prefix += num
        if prefix-target >= 0 {
            if j, exists := prefixMap[prefix-target]; exists {
                minLength = min(minLength, i-j)
            }
        }
        prefixMap[prefix] = i
    }
    
    if minLength == len(nums)+1 {
        return 0
    }
    return minLength
}
```
- **Pros**: Works with negative numbers
- **Cons**: More complex, O(N log N) due to map operations

### Extensions for Interviews:
- **Exact Target**: Find subarray with sum exactly equal to target
- **Maximum Length**: Find longest subarray with sum >= target
- **Count Subarrays**: Count number of subarrays with sum >= target
- **Negative Numbers**: Handle arrays with negative values
*/
func main() {
	// Test cases
	testCases := []struct {
		target int
		nums   []int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}},
		{4, []int{1, 4, 4}},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}},
		{15, []int{1, 2, 3, 4, 5}},
		{100, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{5, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}},
		{3, []int{1, 1, 1, 1, 1, 1, 1}},
		{1, []int{2, 3, 1, 2, 4, 3}},
	}
	
	for i, tc := range testCases {
		result := minSubArrayLen(tc.target, tc.nums)
		fmt.Printf("Test Case %d: target=%d, nums=%v -> Min length: %d\n", 
			i+1, tc.target, tc.nums, result)
	}
}
