package main

import "fmt"

// 560. Subarray Sum Equals K
// Time: O(N), Space: O(N)
func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1 // Initialize with sum 0 occurring once
	
	currentSum := 0
	
	for _, num := range nums {
		currentSum += num
		
		// Check if (currentSum - k) exists in prefixSum
		if freq, exists := prefixSum[currentSum-k]; exists {
			count += freq
		}
		
		// Update prefixSum with currentSum
		prefixSum[currentSum]++
	}
	
	return count
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Prefix Sum with Hash Map Counting
- **Prefix Sum Tracking**: Maintain running sum as we iterate
- **Hash Map Counting**: Count occurrences of each prefix sum
- **Difference Lookup**: Find (currentSum - k) in hash map
- **Count Accumulation**: Add frequency of matching prefix sums

## 2. PROBLEM CHARACTERISTICS
- **Subarray Counting**: Count all subarrays with sum exactly k
- **Prefix Sum Property**: Subarray sum = prefix[j] - prefix[i]
- **Hash Map Efficiency**: O(1) lookup for prefix sum frequencies
- **Multiple Solutions**: Need to count all qualifying subarrays

## 3. SIMILAR PROBLEMS
- Continuous Subarray Sum (LeetCode 523) - Subarray sum divisible by k
- Maximum Size Subarray Sum Equals K (LeetCode 325) - Longest subarray with sum k
- Subarray Sum Less Than K - Subarrays with sum less than k
- Range Sum Query - Various prefix sum applications

## 4. KEY OBSERVATIONS
- **Prefix sum formula**: If prefix[j] - prefix[i] = k, then prefix[i] = prefix[j] - k
- **Hash map natural fit**: Store frequencies of prefix sums
- **Single pass efficiency**: Process array once while maintaining counts
- **Count accumulation**: Multiple occurrences add to total count

## 5. VARIATIONS & EXTENSIONS
- **Range constraints**: Subarrays within specific index ranges
- **Length constraints**: Subarrays with minimum/maximum length
- **Multiple k values**: Count subarrays for multiple target sums
- **Negative numbers**: Handle negative values in array and k

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can array contain negative numbers? What about empty subarrays?"
- Edge cases: empty array, single element, all zeros, negative numbers
- Time complexity: O(N) time, O(N) space for hash map
- Count optimization: Use hash map for O(1) prefix sum lookup

## 7. COMMON MISTAKES
- Not initializing hash map with sum 0
- Using nested loops (O(N²) time)
- Not handling negative numbers correctly
- Forgetting to update hash map after lookup
- Not counting subarrays starting from index 0

## 8. OPTIMIZATION STRATEGIES
- **Single pass**: Process array once with prefix sums
- **Hash map counting**: Store frequencies instead of just existence
- **Early initialization**: Initialize with sum 0 occurring once
- **Efficient lookup**: O(1) hash map operations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding balanced transactions:**
- You're tracking daily financial transactions (positive/negative)
- You want to find periods where total net balance equals target amount k
- You keep running totals of your balance
- For each day, you check if you've seen a balance that's exactly k less than current
- Each time you find such a balance, you've found a profitable period

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers, target sum k
2. **Goal**: Count all continuous subarrays with sum exactly k
3. **Output**: Number of qualifying subarrays
4. **Constraint**: Subarrays must be contiguous

#### Phase 2: Key Insight Recognition
- **"Prefix sum property"** → Subarray sum = prefix[j] - prefix[i]
- **"Hash map counting"** → Store frequencies of prefix sums
- **"Difference lookup"** → Find (currentSum - k) in hash map
- **"Count accumulation"** → Multiple matches contribute to total

#### Phase 3: Strategy Development
```
Human thought process:
"I need to count all subarrays that sum to k.
For each position, I'll track the cumulative sum.
If I've seen a prefix sum of (currentSum - k) before,
then the subarray between that position and current position sums to k.
I'll store frequencies of prefix sums in a hash map.
Each time I find a match, I add the frequency to my count."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 (no subarrays)
- **Single element**: Check if element equals k
- **All zeros**: Handle multiple zero subarrays correctly
- **Negative numbers**: Works correctly with prefix sums

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: nums = [1, 1, 1], k = 2

Human thinking:
"I'll track cumulative sum and count subarrays:

Initialize: prefixSum = {0: 1}, currentSum = 0, count = 0

Position 0 (value 1):
currentSum = 0 + 1 = 1
Looking for (1 - 2) = -1 in prefixSum: not found
Update prefixSum: {0: 1, 1: 1}

Position 1 (value 1):
currentSum = 1 + 1 = 2
Looking for (2 - 2) = 0 in prefixSum: found with frequency 1
Add to count: count = 1 (subarray [0,1])
Update prefixSum: {0: 1, 1: 1, 2: 1}

Position 2 (value 1):
currentSum = 2 + 1 = 3
Looking for (3 - 2) = 1 in prefixSum: found with frequency 1
Add to count: count = 2 (subarray [1,2])
Update prefixSum: {0: 1, 1: 1, 2: 1, 3: 1}

Final count = 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why prefix sums work**: Captures cumulative information efficiently
- **Why hash map works**: O(1) lookup for prefix sum frequencies
- **Why difference formula**: currentSum - k finds starting points
- **Why counting frequencies**: Multiple subarrays can end at same position

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use nested loops?"** → O(N²) time, too slow for large inputs
2. **"Should I use sliding window?"** → Doesn't work with negative numbers
3. **"What about initialization?"** → Must initialize with sum 0
4. **"Can I optimize space?"** → Hash map is already optimal

### Real-World Analogy
**Like finding profitable trading periods:**
- You're tracking daily stock price changes
- You want to find periods where total profit equals target k
- You keep running totals of cumulative profit/loss
- For each day, you check if you've seen a cumulative value that's exactly k less
- Each match represents a profitable trading period

### Human-Readable Pseudocode
```
function subarraySum(nums, k):
    count = 0
    prefixSum = map()
    prefixSum[0] = 1  // sum 0 occurs once
    currentSum = 0
    
    for num in nums:
        currentSum += num
        
        // Check if (currentSum - k) exists
        if (currentSum - k) in prefixSum:
            count += prefixSum[currentSum - k]
        
        // Update prefixSum with currentSum
        prefixSum[currentSum]++
    
    return count
```

### Execution Visualization

### Example: nums = [1, 1, 1], k = 2
```
Prefix Sum and Count Evolution:
Initialize: prefixSum = {0: 1}, currentSum = 0, count = 0

i=0, num=1:
currentSum = 1
Looking for (1-2) = -1: not found
Update: prefixSum = {0: 1, 1: 1}

i=1, num=1:
currentSum = 2
Looking for (2-2) = 0: found with frequency 1
count = 1 (subarray [0,1])
Update: prefixSum = {0: 1, 1: 1, 2: 1}

i=2, num=1:
currentSum = 3
Looking for (3-2) = 1: found with frequency 1
count = 2 (subarray [1,2])
Update: prefixSum = {0: 1, 1: 1, 2: 1, 3: 1}

Final count = 2
```

### Key Visualization Points:
- **Prefix sum tracking**: Running cumulative sum
- **Hash map lookup**: Checking for (currentSum - k)
- **Count accumulation**: Adding frequencies to total count
- **Frequency updates**: Updating hash map with current sum

### Memory Layout Visualization:
```
Array:    [1,  1,  1]
Index:      0   1   2
Sum:       1   2   3
Target-k:  -1   0   1
Found:      ✓   ✓   ✓
Count:     0   1   2

Hash Map Evolution:
{0: 1} → {0: 1, 1: 1} → {0: 1, 1: 1, 2: 1} → {0: 1, 1: 1, 2: 1, 3: 1}

Subarrays found:
[0,1]: sum = 1 + 1 = 2 ✓
[1,2]: sum = 1 + 1 = 2 ✓
```

### Time Complexity Breakdown:
- **Hash map approach**: O(N) time, O(N) space
- **Nested loops**: O(N²) time, O(1) space (too slow)
- **Sliding window**: O(N) time, O(1) space (only for positive numbers)
- **Prefix sum array**: O(N²) time, O(N) space

### Alternative Approaches:

#### 1. Nested Loops (O(N²) time, O(1) space)
```go
func subarraySumBrute(nums []int, k int) int {
    count := 0
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum == k {
                count++
            }
        }
    }
    return count
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time, too slow for large inputs

#### 2. Sliding Window (O(N) time, O(1) space) - Only for Positive Numbers
```go
func subarraySumSliding(nums []int, k int) int {
    // Only works for positive numbers
    count := 0
    left := 0
    currentSum := 0
    
    for right := 0; right < len(nums); right++ {
        currentSum += nums[right]
        
        for currentSum > k && left <= right {
            currentSum -= nums[left]
            left++
        }
        
        if currentSum == k {
            count++
        }
    }
    
    return count
}
```
- **Pros**: O(N) time, O(1) space
- **Cons**: Only works for positive numbers

#### 3. Prefix Sum Array with Binary Search (O(N log N) time, O(N) space)
```go
func subarraySumBinarySearch(nums []int, k int) int {
    prefix := make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    
    count := 0
    for i := 1; i <= len(nums); i++ {
        target := prefix[i] - k
        // Binary search for target in prefix[0:i]
        left, right := 0, i-1
        for left <= right {
            mid := (left + right) / 2
            if prefix[mid] == target {
                count++
                break
            } else if prefix[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return count
}
```
- **Pros**: Doesn't use hash map
- **Cons**: O(N log N) time, more complex

### Extensions for Interviews:
- **Range Constraints**: Count subarrays within specific index ranges
- **Length Constraints**: Subarrays with minimum/maximum length
- **Multiple k Values**: Count subarrays for multiple target sums
- **Return Indices**: Return actual indices of qualifying subarrays
- **Modulo Arithmetic**: Count subarrays with sum modulo k equals target
*/
func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1}, 2},
		{[]int{1, 2, 3}, 3},
		{[]int{1, -1, 0}, 0},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{-1, -1, 1}, 0},
		{[]int{3, 4, 7, -2, 2, 1, 4, 2}, 7},
		{[]int{1}, 1},
		{[]int{1}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 1, 2, 1}, 3},
	}
	
	for i, tc := range testCases {
		result := subarraySum(tc.nums, tc.k)
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Subarrays with sum k: %d\n", 
			i+1, tc.nums, tc.k, result)
	}
}
