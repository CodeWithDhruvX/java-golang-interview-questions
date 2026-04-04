package main

import "fmt"

// 238. Product of Array Except Self
// Time: O(N), Space: O(1) (excluding output array)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	
	// First pass: calculate left products
	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}
	
	// Second pass: calculate right products and multiply with left products
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Prefix/Suffix Product
- **Left Product Pass**: Calculate products of all elements to the left
- **Right Product Pass**: Calculate products of all elements to the right
- **In-place Combination**: Multiply left and right products in result array
- **Space Optimization**: Use output array to store left products

## 2. PROBLEM CHARACTERISTICS
- **Exclusion Product**: Product of all elements except current index
- **No Division**: Cannot use division operator
- **Linear Time**: Need O(N) solution
- **Constant Space**: O(1) extra space excluding output

## 3. SIMILAR PROBLEMS
- Range Sum Query (LeetCode 303) - Prefix sum for range queries
- Subarray Sum Equals K (LeetCode 560) - Prefix sum with hash map
- Continuous Subarray Sum (LeetCode 523) - Prefix sum modulo
- Product of Array Except Self (Variations) - With division allowed

## 4. KEY OBSERVATIONS
- **Two-pass approach**: One for left products, one for right products
- **Output array reuse**: Can use output array to store intermediate results
- **Prefix/suffix concept**: Left and right cumulative products
- **No division constraint**: Must use multiplication approach

## 5. VARIATIONS & EXTENSIONS
- **Division allowed**: Can compute total product and divide by current element
- **Multiple queries**: Precompute for multiple product queries
- **2D version**: Product of array except self in 2D matrix
- **Streaming data**: Handle data that arrives over time

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is division allowed? What about zero elements?"
- Edge cases: empty array, single element, zeros in array
- Time complexity: O(N) time, O(1) space (excluding output)
- Zero handling: Multiple zeros require special consideration

## 7. COMMON MISTAKES
- Using division when not allowed
- Not handling zero elements properly
- Using O(N²) nested loop approach
- Not optimizing space to O(1)
- Forgetting to initialize left/right products

## 8. OPTIMIZATION STRATEGIES
- **Two-pass approach**: Optimal O(N) time, O(1) space
- **Output array reuse**: Store left products in result array
- **Early zero handling**: Check for zeros upfront
- **In-place computation**: Minimize additional memory usage

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like calculating team contributions without one member:**
- You have a team with N members, each has a productivity value
- For each member, you want to know the total productivity of everyone else
- You can't use division (maybe because some values are zero)
- For each position, you multiply productivity of everyone to the left and right
- The result is the team's total productivity excluding that member

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers
2. **Goal**: For each index, product of all other elements
3. **Constraint**: Cannot use division
4. **Output**: Array of products excluding self

#### Phase 2: Key Insight Recognition
- **"Two-pass natural fit"** → Left products then right products
- **"Output array reuse"** → Can store intermediate results
- **"Prefix/suffix concept"** → Cumulative products from both sides
- **"No division constraint"** → Must use multiplication approach

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find product of all elements except current one, without division.
I can do this in two passes:
1. First pass: calculate product of all elements to the left of each position
2. Second pass: multiply by product of all elements to the right of each position
I can store the left products in the result array to save space.
Then multiply with right products in the second pass."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **Single element**: Return [1] (product of empty set)
- **Zero elements**: Handle properly (multiple zeros → all zeros except zero positions)
- **Large arrays**: Handle efficiently with O(1) space

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [1, 2, 3, 4]

Human thinking:
"First pass: calculate left products
Position 0: left of 0 is nothing → product = 1
Position 1: left of 1 is [1] → product = 1
Position 2: left of 2 is [1,2] → product = 2
Position 3: left of 3 is [1,2,3] → product = 6
Result after first pass: [1, 1, 2, 6]

Second pass: multiply by right products
Position 3: right of 3 is nothing → product = 1 → [1,1,2,6]*1 = [1,1,2,6]
Position 2: right of 2 is [4] → product = 4 → [1,1,2,6]*4 = [1,1,8,6]
Position 1: right of 1 is [3,4] → product = 12 → [1,1,8,6]*12 = [1,12,8,6]
Position 0: right of 0 is [2,3,4] → product = 24 → [1,12,8,6]*24 = [24,12,8,6]

Final result: [24, 12, 8, 6]"
```

#### Phase 6: Intuition Validation
- **Why two-pass works**: Separates left and right product calculations
- **Why O(N) time**: Each element processed exactly twice
- **Why O(1) space**: Reuse output array for intermediate storage
- **Why no division**: Multiplication approach handles zeros correctly

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use total product?"** → Division not allowed, zeros cause division by zero
2. **"Should I use nested loops?"** → O(N²) time, too slow for large inputs
3. **"What about zeros?"** → Need special handling, multiple zeros affect all results
4. **"Can I optimize further?"** → Two-pass approach is already optimal

### Real-World Analogy
**Like calculating total team score excluding each player:**
- You have a sports team with N players, each has a score
- For each player, you want to know the team's total score without that player
- You can't just divide total score by player's score (scores might be zero)
- For each player, you sum scores of all other players
- This is equivalent to product in the mathematical sense

### Human-Readable Pseudocode
```
function productExceptSelf(nums):
    n = length(nums)
    result = array of size n
    
    // First pass: left products
    leftProduct = 1
    for i from 0 to n-1:
        result[i] = leftProduct
        leftProduct *= nums[i]
    
    // Second pass: right products
    rightProduct = 1
    for i from n-1 down to 0:
        result[i] *= rightProduct
        rightProduct *= nums[i]
    
    return result
```

### Execution Visualization

### Example: nums = [1, 2, 3, 4]
```
First Pass (Left Products):
Initial: leftProduct = 1, result = [0, 0, 0, 0]

i=0: result[0] = 1, leftProduct = 1*1 = 1 → result = [1, 0, 0, 0]
i=1: result[1] = 1, leftProduct = 1*2 = 2 → result = [1, 1, 0, 0]
i=2: result[2] = 2, leftProduct = 2*3 = 6 → result = [1, 1, 2, 0]
i=3: result[3] = 6, leftProduct = 6*4 = 24 → result = [1, 1, 2, 6]

Second Pass (Right Products):
Initial: rightProduct = 1, result = [1, 1, 2, 6]

i=3: result[3] *= 1 → 6*1 = 6, rightProduct = 1*4 = 4 → result = [1, 1, 2, 6]
i=2: result[2] *= 4 → 2*4 = 8, rightProduct = 4*3 = 12 → result = [1, 1, 8, 6]
i=1: result[1] *= 12 → 1*12 = 12, rightProduct = 12*2 = 24 → result = [1, 12, 8, 6]
i=0: result[0] *= 24 → 1*24 = 24, rightProduct = 24*1 = 24 → result = [24, 12, 8, 6]

Final result: [24, 12, 8, 6]
```

### Key Visualization Points:
- **Left product accumulation**: Building products from left to right
- **Right product multiplication**: Combining with products from right to left
- **In-place computation**: Using result array for intermediate storage
- **Two-pass efficiency**: Each element processed exactly twice

### Memory Layout Visualization:
```
Array: [1, 2, 3, 4]
         ↓  ↓  ↓  ↓
Left:   1  1  2  6  (products to the left)
Right: 24 12  4  1  (products to the right)
Final: 24 12  8  6  (left * right)

Processing Flow:
Step 1: Store left products in result
Step 2: Multiply by right products in reverse
```

### Time Complexity Breakdown:
- **Two-pass approach**: O(N) time, O(1) space (excluding output)
- **Nested loops**: O(N²) time, O(1) space (too slow)
- **Division approach**: O(N) time, O(1) space (if allowed)
- **Precomputation**: O(N) time, O(N) space for multiple queries

### Alternative Approaches:

#### 1. Division Approach (O(N) time, O(1) space) - If Allowed
```go
func productExceptSelfDivision(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    
    // Count zeros and calculate total product
    zeroCount := 0
    totalProduct := 1
    
    for _, num := range nums {
        if num == 0 {
            zeroCount++
        } else {
            totalProduct *= num
        }
    }
    
    for i, num := range nums {
        if zeroCount > 1 {
            result[i] = 0
        } else if zeroCount == 1 {
            if num == 0 {
                result[i] = totalProduct
            } else {
                result[i] = 0
            }
        } else {
            result[i] = totalProduct / num
        }
    }
    
    return result
}
```
- **Pros**: Simple logic, single pass
- **Cons**: Division not allowed, zero handling complex

#### 2. Prefix and Suffix Arrays (O(N) time, O(N) space)
```go
func productExceptSelfArrays(nums []int) []int {
    n := len(nums)
    left := make([]int, n)
    right := make([]int, n)
    result := make([]int, n)
    
    // Calculate left products
    left[0] = 1
    for i := 1; i < n; i++ {
        left[i] = left[i-1] * nums[i-1]
    }
    
    // Calculate right products
    right[n-1] = 1
    for i := n-2; i >= 0; i-- {
        right[i] = right[i+1] * nums[i+1]
    }
    
    // Combine left and right
    for i := 0; i < n; i++ {
        result[i] = left[i] * right[i]
    }
    
    return result
}
```
- **Pros**: Clear separation of concerns
- **Cons**: Uses O(N) extra space

#### 3. Logarithmic Approach (O(N) time, O(1) space) - For Positive Numbers
```go
func productExceptSelfLog(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    
    // Calculate log sum (for positive numbers only)
    logSum := 0.0
    for _, num := range nums {
        if num <= 0 {
            return productExceptSelf(nums) // Fall back to standard approach
        }
        logSum += math.Log(float64(num))
    }
    
    for i, num := range nums {
        result[i] = int(math.Exp(logSum - math.Log(float64(num))))
    }
    
    return result
}
```
- **Pros**: Mathematical approach
- **Cons**: Only works for positive numbers, floating point precision issues

### Extensions for Interviews:
- **Multiple Queries**: Precompute for efficient multiple product queries
- **2D Matrix**: Product of matrix except each row/column
- **Streaming Data**: Handle data that arrives over time
- **Modulo Arithmetic**: Return results modulo large number
- **Large Numbers**: Handle integer overflow with big integers
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
		{1, 2, 3, 4, 5},
		{0, 0},
		{1, 0},
		{2, 3, 0, 4},
		{-1, -2, -3, -4},
		{1, 1, 1, 1},
		{5},
		{},
	}
	
	for i, nums := range testCases {
		result := productExceptSelf(nums)
		fmt.Printf("Test Case %d: %v -> Product except self: %v\n", 
			i+1, nums, result)
	}
}
