package main

import "fmt"

// 523. Continuous Subarray Sum
// Time: O(N), Space: O(N)
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	
	// Map to store prefix sum modulo k and its earliest index
	prefixMod := make(map[int]int)
	prefixMod[0] = -1 // Initialize with sum 0 at index -1
	
	currentSum := 0
	
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		
		var mod int
		if k != 0 {
			mod = currentSum % k
			if mod < 0 { // Handle negative numbers
				mod += k
			}
		} else {
			mod = currentSum // When k is 0, we need exact sum of 0
		}
		
		// Check if this modulo has been seen before
		if prevIndex, exists := prefixMod[mod]; exists {
			// Check if the subarray length is at least 2
			if i-prevIndex >= 2 {
				return true
			}
		} else {
			// Store the first occurrence of this modulo
			prefixMod[mod] = i
		}
	}
	
	return false
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Prefix Sum Modulo with Hash Map
- **Prefix Sum Tracking**: Maintain running sum as we iterate
- **Modulo Arithmetic**: Use modulo k to find equal remainders
- **Hash Map Storage**: Store first occurrence of each remainder
- **Subarray Detection**: Equal remainders indicate divisible subarray

## 2. PROBLEM CHARACTERISTICS
- **Continuous Subarray**: Must be contiguous elements
- **Sum Divisibility**: Subarray sum must be divisible by k
- **Length Constraint**: Subarray must have at least 2 elements
- **Multiple Solutions**: Need to find any qualifying subarray

## 3. SIMILAR PROBLEMS
- Subarray Sum Equals K (LeetCode 560) - Find subarrays with exact sum
- Maximum Size Subarray Sum Equals K (LeetCode 325) - Largest subarray with sum k
- Continuous Subarray Sum with Negative Numbers - Handle negative values
- Prefix Sum Applications - Various prefix sum patterns

## 4. KEY OBSERVATIONS
- **Modulo property**: If sum[i] % k == sum[j] % k, then subarray (i+1..j) sum is divisible by k
- **Prefix sum insight**: Use cumulative sums to detect subarray properties
- **Hash map efficiency**: O(1) lookup for remainder existence
- **Zero handling**: Special case when k = 0

## 5. VARIATIONS & EXTENSIONS
- **Negative numbers**: Handle negative values in array
- **Zero k value**: Special case requiring exact zero sum
- **Multiple k values**: Check divisibility by multiple numbers
- **Count solutions**: Count all qualifying subarrays instead of just existence

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can array contain negative numbers? What about k = 0?"
- Edge cases: k = 0, negative numbers, short arrays
- Time complexity: O(N) time, O(k) space (worst case O(N))
- Space optimization: Use hash map for remainder storage

## 7. COMMON MISTAKES
- Not handling k = 0 case properly
- Not checking subarray length constraint (>= 2)
- Using nested loops (O(N²) time)
- Not handling negative numbers correctly
- Forgetting to initialize hash map with remainder 0

## 8. OPTIMIZATION STRATEGIES
- **Single pass**: Process array once with prefix sums
- **Hash map lookup**: O(1) remainder checking
- **Early termination**: Stop when first valid subarray found
- **Space optimization**: Limit hash map size to k entries

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding divisible sections in a number line:**
- You're walking along a number line with values at each position
- You want to find a section where the total sum is divisible by k
- You keep track of your cumulative sum and its remainder when divided by k
- If you encounter the same remainder again, the section between them sums to multiple of k
- This works because the difference between same remainders is divisible by k

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers, integer k
2. **Goal**: Find continuous subarray with sum divisible by k
3. **Constraint**: Subarray must have at least 2 elements
4. **Output**: Boolean indicating existence

#### Phase 2: Key Insight Recognition
- **"Prefix sum modulo"** → Track cumulative sum modulo k
- **"Hash map natural fit"** → Store first occurrence of each remainder
- **"Equal remainder property"** → Same remainder indicates divisible subarray
- **"Length constraint"** → Need to check subarray length >= 2

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find a continuous subarray whose sum is divisible by k.
I'll track the cumulative sum as I go through the array.
For each position, I'll calculate sum % k.
If I've seen this remainder before, the subarray between them sums to multiple of k.
I'll store the first occurrence of each remainder in a hash map.
I also need to ensure the subarray has at least 2 elements."
```

#### Phase 4: Edge Case Handling
- **k = 0**: Need exact sum of 0, different logic
- **Negative numbers**: Handle modulo with negative values
- **Short arrays**: Arrays with < 2 elements return false
- **Large k values**: Hash map handles any k value

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: nums = [23, 2, 4, 6, 7], k = 6

Human thinking:
"I'll track cumulative sum and remainder:

Position 0: sum = 23, remainder = 23 % 6 = 5
Store: {5: 0}

Position 1: sum = 23 + 2 = 25, remainder = 25 % 6 = 1
Store: {5: 0, 1: 1}

Position 2: sum = 25 + 4 = 29, remainder = 29 % 6 = 5
I've seen remainder 5 before at index 0!
Subarray from index 1 to 2: [2, 4]
Length = 2 (>= 2), sum = 6, 6 % 6 = 0 ✓
Found valid subarray, return true!"
```

#### Phase 6: Intuition Validation
- **Why modulo works**: Equal remainders indicate difference divisible by k
- **Why prefix sum**: Captures cumulative information efficiently
- **Why hash map**: O(1) lookup for remainder existence
- **Why length check**: Ensures subarray meets minimum length requirement

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use nested loops?"** → O(N²) time, too slow for large inputs
2. **"Should I use sliding window?"** → Variable window size makes it complex
3. **"What about k = 0?"** → Special case requiring exact zero sum
4. **"Can I optimize space?"** → Hash map is already optimal for this approach

### Real-World Analogy
**Like finding divisible payment periods:**
- You're tracking daily income and expenses
- You want to find a period where total net income is divisible by weekly budget
- You keep running totals and check when they align with budget cycles
- When the remainder repeats, you've found a period that fits budget constraints
- This helps identify financially balanced time periods

### Human-Readable Pseudocode
```
function checkSubarraySum(nums, k):
    if len(nums) < 2:
        return false
    
    prefixMod = map()
    prefixMod[0] = -1  // sum 0 at index -1
    currentSum = 0
    
    for i from 0 to len(nums)-1:
        currentSum += nums[i]
        
        if k != 0:
            mod = currentSum % k
            if mod < 0: mod += k  // handle negative
        else:
            mod = currentSum  // when k is 0
        
        if mod in prefixMod:
            if i - prefixMod[mod] >= 2:
                return true
        else:
            prefixMod[mod] = i
    
    return false
```

### Execution Visualization

### Example: nums = [23, 2, 4, 6, 7], k = 6
```
Prefix Sum and Remainder Evolution:
i=0: sum=23, mod=23%6=5 → store {5:0}
i=1: sum=25, mod=25%6=1 → store {5:0, 1:1}
i=2: sum=29, mod=29%6=5 → found remainder 5 at index 0
   subarray [1,2], length=2, sum=6, 6%6=0 ✓
   RETURN true

Hash Map Evolution:
{0: -1}  // Initialize
{0: -1, 5: 0}  // After i=0
{0: -1, 5: 0, 1: 1}  // After i=1
Found remainder 5 at index 0, current index = 2
Subarray length = 2 - 0 = 2 (>= 2) ✓
```

### Key Visualization Points:
- **Prefix sum tracking**: Running cumulative sum
- **Modulo calculation**: Computing remainder at each step
- **Hash map lookup**: Checking for previous same remainder
- **Length validation**: Ensuring subarray meets size requirement

### Memory Layout Visualization:
```
Array:    [23,  2,  4,  6,  7]
Index:      0   1   2   3   4
Sum:       23  25  29  35  42
Modulo 6:   5   1   5   5   0

Hash Map:
{0: -1} → {0: -1, 5: 0} → {0: -1, 5: 0, 1: 1}

At index 2: remainder 5 found at index 0
Subarray [1,2]: nums[1] + nums[2] = 2 + 4 = 6
6 % 6 = 0 ✓
```

### Time Complexity Breakdown:
- **Hash map approach**: O(N) time, O(min(N,k)) space
- **Nested loops**: O(N²) time, O(1) space (too slow)
- **Sliding window**: O(N) time, O(1) space (complex for variable window)
- **Prefix sum array**: O(N) time, O(N) space

### Alternative Approaches:

#### 1. Nested Loops (O(N²) time, O(1) space)
```go
func checkSubarraySumBrute(nums []int, k int) bool {
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if j-i+1 >= 2 && (k == 0 && sum == 0 || k != 0 && sum%k == 0) {
                return true
            }
        }
    }
    return false
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time, too slow for large inputs

#### 2. Prefix Sum Array (O(N²) time, O(N) space)
```go
func checkSubarraySumPrefixArray(nums []int, k int) bool {
    prefix := make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            sum := prefix[j+1] - prefix[i]
            if j-i+1 >= 2 && (k == 0 && sum == 0 || k != 0 && sum%k == 0) {
                return true
            }
        }
    }
    return false
}
```
- **Pros**: Clear separation of concerns
- **Cons**: Still O(N²) time, extra space overhead

#### 3. Optimized for Large k (O(N) time, O(k) space)
```go
func checkSubarraySumOptimized(nums []int, k int) bool {
    if k == 0 {
        return checkSubarraySumZero(nums)
    }
    
    seen := make([]bool, k)
    seen[0] = true
    
    sum := 0
    for i, num := range nums {
        sum += num
        mod := ((sum % k) + k) % k // handle negative
        
        if seen[mod] && i > 0 {
            return true
        }
        seen[mod] = true
    }
    
    return false
}

func checkSubarraySumZero(nums []int) bool {
    sum := 0
    sumSet := make(map[int]int)
    sumSet[0] = -1
    
    for i, num := range nums {
        sum += num
        if prevIndex, exists := sumSet[sum]; exists {
            if i-prevIndex >= 2 {
                return true
            }
        } else {
            sumSet[sum] = i
        }
    }
    
    return false
}
```
- **Pros**: More memory efficient for small k
- **Cons**: Limited to reasonable k values

### Extensions for Interviews:
- **Count All Subarrays**: Return count of all qualifying subarrays
- **Return Indices**: Return actual indices of qualifying subarrays
- **Multiple k Values**: Check divisibility by multiple numbers
- **Longest Subarray**: Find longest qualifying subarray
- **Negative k Values**: Handle negative divisor values
*/
func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{23, 2, 4, 6, 7}, 6},
		{[]int{23, 2, 6, 4, 7}, 6},
		{[]int{23, 2, 6, 4, 7}, 13},
		{[]int{0, 0}, 0},
		{[]int{5, 0, 0, 0}, 3},
		{[]int{1, 2, 3}, 5},
		{[]int{1, 2, 12, 1, -1, 2, 1, 2, 10, 1}, 3},
		{[]int{1, 2}, 0},
		{[]int{0, 1, 2, 3}, 0},
		{[]int{5, 1, 2, 3, 1, 2, 3, 1, 2, 3, 5, 5}, 3},
	}
	
	for i, tc := range testCases {
		result := checkSubarraySum(tc.nums, tc.k)
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Continuous subarray sum: %t\n", 
			i+1, tc.nums, tc.k, result)
	}
}
