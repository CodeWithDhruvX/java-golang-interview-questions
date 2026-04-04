package main

import "fmt"

// 136. Single Number
// Time: O(N), Space: O(1)
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// Alternative approach using hash map (O(N) space)
func singleNumberHash(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	
	for num, freq := range count {
		if freq == 1 {
			return num
		}
	}
	
	return -1 // Should not reach here for valid input
}

// XOR properties explanation
func singleNumberWithExplanation(nums []int) int {
	fmt.Printf("XOR Properties:\n")
	fmt.Printf("1. a ^ a = 0 (XOR of same numbers is 0)\n")
	fmt.Printf("2. a ^ 0 = a (XOR with 0 is the number itself)\n")
	fmt.Printf("3. XOR is commutative and associative\n\n")
	
	fmt.Printf("Processing: ")
	result := 0
	for i, num := range nums {
		if i > 0 {
			fmt.Printf(" ^ ")
		}
		fmt.Printf("%d", num)
		result ^= num
	}
	fmt.Printf(" = %d\n\n", result)
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: XOR for Finding Unique Element
- **XOR Properties**: a ^ a = 0, a ^ 0 = a, commutative and associative
- **Pair Cancellation**: XOR of pairs cancels out (a ^ a = 0)
- **Unique Isolation**: Remaining element is the unique one
- **Linear Scan**: Single pass through array

## 2. PROBLEM CHARACTERISTICS
- **Array with Duplicates**: All elements appear twice except one
- **Find Unique**: Identify the element that appears only once
- **Linear Time**: Need O(N) solution
- **Constant Space**: Need O(1) extra space

## 3. SIMILAR PROBLEMS
- Find Missing Number (LeetCode 268) - XOR for missing element
- Find the Duplicate Number (LeetCode 287) - XOR for duplicate
- Two Single Numbers (LeetCode 260) - XOR for two unique numbers
- Bit Manipulation problems using XOR properties

## 4. KEY OBSERVATIONS
- **XOR Identity**: x ^ x = 0 (any number XORed with itself is 0)
- **XOR with Zero**: x ^ 0 = x (any number XORed with 0 is itself)
- **Commutative**: Order doesn't matter: a ^ b ^ c = c ^ b ^ a
- **Associative**: Grouping doesn't matter: (a ^ b) ^ c = a ^ (b ^ c)

## 5. VARIATIONS & EXTENSIONS
- **Multiple Unique Elements**: Find elements appearing odd number of times
- **Missing and Duplicate**: Find both missing and duplicate elements
- **Bit Position Analysis**: Find which bit positions differ
- **Streaming Data**: Handle infinite streams with XOR accumulation

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are there exactly one unique element? Range of numbers?"
- Edge cases: empty array, single element, negative numbers
- Time complexity: O(N) time, O(1) space
- Key insight: XOR properties cancel out duplicates
- Alternative: Hash map approach with O(N) space

## 7. COMMON MISTAKES
- Not understanding XOR properties correctly
- Using wrong operation (AND/OR instead of XOR)
- Not handling edge cases (empty array, single element)
- Assuming numbers are positive
- Forgetting to initialize result to 0

## 8. OPTIMIZATION STRATEGIES
- **XOR Approach**: O(N) time, O(1) space - optimal
- **Hash Map**: O(N) time, O(N) space - alternative
- **Bit Analysis**: O(N) time, O(1) space for specific patterns
- **Early Termination**: Not applicable (need all elements)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like canceling out pairs in a matching game:**
- You have pairs of matching numbers and one unique number
- When you XOR matching pairs, they cancel out (become 0)
- The unique number remains after all cancellations
- XOR is like a "canceling operation" for identical values

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array where every element appears twice except one
2. **Goal**: Find the element that appears only once
3. **Constraint**: Exactly one unique element
4. **Output**: The unique element value

#### Phase 2: Key Insight Recognition
- **"XOR natural fit"** → Cancels out identical pairs
- **"Identity property"** → x ^ x = 0 for any x
- **"Commutative"** → Order doesn't matter
- **"Linear solution"** → Single pass sufficient

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the unique element in an array of pairs.
If I XOR all numbers together:
- Each pair (a, a) becomes a ^ a = 0
- All pairs cancel out to 0
- Only the unique element remains
So result = XOR of all elements!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 or handle appropriately
- **Single element**: Return that element
- **Negative numbers**: XOR works with negative numbers too
- **Large numbers**: Handle potential overflow

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
nums = [2, 2, 1]

Human thinking:
"I'll XOR all elements:
Step 1: result = 0 ^ 2 = 2
Step 2: result = 2 ^ 2 = 0 (pair cancels)
Step 3: result = 0 ^ 1 = 1 (unique remains)

Final result: 1 ✓"

Let me trace another:
nums = [4, 1, 2, 1, 2]

Step 1: result = 0 ^ 4 = 4
Step 2: result = 4 ^ 1 = 5
Step 3: result = 5 ^ 2 = 7
Step 4: result = 7 ^ 1 = 6 (1 pair starts canceling)
Step 5: result = 6 ^ 2 = 4 (2 pair completes cancellation)

Final result: 4 ✓"
```

#### Phase 6: Intuition Validation
- **Why XOR works**: Identical numbers cancel out due to x ^ x = 0
- **Why O(N)**: Single pass through array
- **Why O(1) space**: Only need accumulator variable
- **Why optimal**: Can't do better than linear time

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use hash map?"** → Works but uses O(N) space
2. **"Should I sort first?"** → O(N log N) time, unnecessary
3. **"What about multiple uniques?"** → Need different approach
4. **"Can I optimize further?"** → XOR is already optimal

### Real-World Analogy
**Like finding the odd one out in a line of twins:**
- You have pairs of identical twins and one unique person
- When you pair up twins, they cancel each other out
- The unique person stands alone
- XOR is like the "pairing detector" that finds the unpaired one

### Human-Readable Pseudocode
```
function singleNumber(nums):
    if nums is empty:
        return 0
    
    result = 0
    for num in nums:
        result = result XOR num
    
    return result
```

### Execution Visualization

### Example: nums = [2, 2, 1]
```
XOR Evolution During Processing:
Initial: result = 0

Process 2: result = 0 ^ 2 = 2
Process 2: result = 2 ^ 2 = 0 (pair cancels)
Process 1: result = 0 ^ 1 = 1 (unique remains)

Final result: 1 ✓
```

### Key Visualization Points:
- **Pair Cancellation**: Identical numbers cancel to 0
- **Unique Preservation**: Only unique element remains
- **Order Independence**: XOR order doesn't matter
- **Linear Processing**: Single pass through array

### Memory Layout Visualization:
```
XOR State During Processing:
nums = [4, 1, 2, 1, 2]

Step-by-step XOR evolution:
result = 0
result = 0 ^ 4 = 4
result = 4 ^ 1 = 5
result = 5 ^ 2 = 7
result = 7 ^ 1 = 6 (first 1 starts canceling)
result = 6 ^ 2 = 4 (second 2 completes cancellation)

Final: result = 4 ✓
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) time complexity
- **Constant Space**: O(1) additional space
- **XOR Operations**: Each XOR is O(1) time
- **Total**: O(N) time, O(1) space

### Alternative Approaches:

#### 1. Hash Map (O(N) time, O(N) space)
```go
func singleNumberHash(nums []int) int {
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }
    
    for num, freq := range count {
        if freq == 1 {
            return num
        }
    }
    
    return -1
}
```
- **Pros**: Intuitive, handles multiple unique elements
- **Cons**: O(N) space, unnecessary for this specific problem

#### 2. Sorting (O(N log N) time, O(1) space)
```go
func singleNumberSort(nums []int) int {
    sort.Ints(nums)
    
    for i := 0; i < len(nums); i += 2 {
        if i+1 < len(nums) && nums[i] != nums[i+1] {
            return nums[i]
        }
    }
    
    return nums[len(nums)-1]
}
```
- **Pros**: Works with sorted data
- **Cons**: O(N log N) time, modifies input

#### 3. Bit Position Analysis (O(N) time, O(1) space)
```go
func singleNumberBitAnalysis(nums []int) int {
    result := 0
    
    for i := 0; i < 32; i++ {
        sum := 0
        for _, num := range nums {
            if (num >> i) & 1 == 1 {
                sum++
            }
        }
        
        if sum%2 == 1 {
            result |= (1 << i)
        }
    }
    
    return result
}
```
- **Pros**: Works for multiple unique elements
- **Cons**: More complex, 32 passes through array

### Extensions for Interviews:
- **Multiple Unique Elements**: Find elements appearing odd number of times
- **Missing and Duplicate**: Find both missing and duplicate elements
- **Streaming Data**: Handle infinite streams with XOR accumulation
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
		{2, 2, 3, 3, 4},
		{0, 1, 1},
		{-1, -1, -2},
		{99, 99, 100},
		{5, 5, 6, 6, 7, 7, 8},
		{10, 10, 10, 10, 15},
	}
	
	for i, nums := range testCases {
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		
		if i == 0 {
			// Show detailed explanation for first test case
			result := singleNumberWithExplanation(nums)
			fmt.Printf("Single number: %d\n\n", result)
		} else {
			result1 := singleNumber(nums)
			result2 := singleNumberHash(nums)
			fmt.Printf("XOR: %d, Hash: %d\n\n", result1, result2)
		}
	}
}
