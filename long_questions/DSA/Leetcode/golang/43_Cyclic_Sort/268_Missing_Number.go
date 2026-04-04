package main

import "fmt"

// 268. Missing Number
// Time: O(N), Space: O(1) - Cyclic Sort approach
func missingNumber(nums []int) int {
	n := len(nums)
	
	// Place each number in its correct position
	i := 0
	for i < n {
		correctPos := nums[i]
		if correctPos < n && nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			i++
		}
	}
	
	// Find the first position where number doesn't match index
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	
	return n // All numbers are in correct positions, missing number is n
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Cyclic Sort for Missing Elements
- **Index-Value Mapping**: Place numbers at their correct indices
- **In-Place Rearrangement**: Sort without extra space
- **Missing Detection**: Find first index with wrong value
- **Range Constraints**: Leverage 0-to-(n-1) range property

## 2. PROBLEM CHARACTERISTICS
- **Missing Number**: Find missing integer from 0 to n range
- **In-Place Operation**: Cannot use extra space
- **Range Utilization**: Numbers in [0, n-1] can be placed at indices [0, n-1]
- **Linear Time**: Single pass through array

## 3. SIMILAR PROBLEMS
- First Missing Positive (LeetCode 41) - Similar with 1-to-n range
- Find All Numbers Disappeared in an Array (LeetCode 448) - Same technique
- Find the Duplicate Number (LeetCode 287) - Cyclic sort variant
- Set Mismatch (LeetCode 645) - Find wrong and missing numbers

## 4. KEY OBSERVATIONS
- **Index Mapping**: Number i should be at position i
- **Range Property**: Only numbers 0 to n-1 can be placed correctly
- **Missing Indicator**: First mismatch reveals missing number
- **Swap Strategy**: Keep swapping until current position is correct

## 5. VARIATIONS & EXTENSIONS
- **Missing Element**: Find missing number from 0..n
- **Duplicate Detection**: Find repeated numbers
- **Multiple Missing**: Find all missing numbers
- **Set Mismatch**: Find swapped numbers

## 6. INTERVIEW INSIGHTS
- Always clarify: "Array size constraints? Value ranges? Space limits?"
- Edge cases: empty array, single element, all correct
- Time complexity: O(N) for cyclic sort, O(N) for final scan
- Space complexity: O(1) - in-place operation
- Key insight: use array indices as hash map for numbers 0 to n-1

## 7. COMMON MISTAKES
- Wrong index calculation (i vs nums[i] confusion)
- Not handling out-of-range numbers properly
- Missing edge cases for empty/full arrays
- Infinite loop with duplicate values
- Wrong bounds checking

## 8. OPTIMIZATION STRATEGIES
- **Cyclic Sort**: O(N) time, O(1) space - optimal for range 0..n-1
- **XOR Method**: O(N) time, O(1) space - elegant mathematical approach
- **Sum Method**: O(N) time, O(1) space - arithmetic approach
- **Hash Set**: O(N) time, O(N) space - simple but uses extra space

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing numbered parking spots:**
- You have n parking spots numbered 0 to n-1 and n cars
- Each car should park in its matching numbered spot
- One spot will be empty (the missing number)
- You park cars one by one, swapping to correct spots
- The first empty spot reveals the missing car's number
- Like a parking attendant organizing cars by their assigned spots

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of n distinct integers from 0 to n (one missing)
2. **Goal**: Find the missing integer
3. **Constraints**: O(1) space, O(N) time
4. **Output**: Missing integer from 0 to n

#### Phase 2: Key Insight Recognition
- **"Index mapping natural"** → Number i belongs at position i
- **"Range utilization"** → Only numbers 0..n-1 can be placed correctly
- **"In-place sorting"** → Use array itself as organizing structure
- **"Missing detection"** → First position with wrong value reveals answer

#### Phase 3: Strategy Development
```
Human thought process:
"I need missing number from 0..n without extra space.
Brute force: use hash set O(N) space.

Cyclic Sort Approach:
1. Place each number i at position i
2. Handle only numbers in range [0, n-1]
3. Skip numbers equal to n (can't be placed)
4. Keep swapping until current position is correct
5. Scan for first position where value != index
6. That index is the missing number

This gives O(N) time, O(1) space!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 (missing from 0..0)
- **Single element**: Return 1 if [0], return 0 if [1]
- **All correct**: Return n (missing is n)
- **Missing 0**: Return 0 (first position wrong)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [3, 0, 1]

Human thinking:
"Cyclic Sort Process:
Step 1: i=0, nums[0]=3
Correct position for 3 is index 3
But 3 >= n (n=3), can't place → i++

Step 2: i=1, nums[1]=0
Correct position for 0 is index 0
Swap nums[1] with nums[0]: [0, 3, 1]
Still at i=1, nums[1]=3
3 >= n, can't place → i++

Step 3: i=2, nums[2]=1
Correct position for 1 is index 1
Swap nums[2] with nums[1]: [0, 1, 3]
Still at i=2, nums[2]=3
3 >= n, can't place → i++

Final array: [0, 1, 3]
Scan for mismatch:
Index 0: 0 ✓ (correct)
Index 1: 1 ✓ (correct)
Index 2: 3 ≠ 2 ✗ (missing number is 2)

Result: 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why index mapping**: Number i naturally belongs at position i
- **Why O(1) space**: Use array itself as organizing structure
- **Why O(N) time**: Each element moves at most once
- **Why correct**: First mismatch reveals smallest missing

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → O(N log N) vs O(N), uses extra space
2. **"Should I use XOR?"** → Yes, elegant alternative
3. **"What about number n?"** → Can't be placed, skip it
4. **"Can I use sum method?"** → Yes, arithmetic approach
5. **"Why swap instead of assign?"** → Need to preserve other values

### Real-World Analogy
**Like organizing numbered chairs in a classroom:**
- You have n chairs numbered 0 to n-1 and n students
- Each student should sit in their matching numbered chair
- One chair will be empty (the missing student)
- You seat students one by one, swapping to correct chairs
- The first empty chair number is the missing student's number
- Like a teacher organizing students by their assigned chairs

### Human-Readable Pseudocode
```
function missingNumber(nums):
    n = length(nums)
    i = 0
    
    while i < n:
        correctPos = nums[i]
        
        # Check if current number is in correct range
        # and not already in correct position
        if correctPos < n and nums[i] != nums[correctPos]:
            swap(nums[i], nums[correctPos])
        else:
            i += 1
    
    # Find first position with wrong number
    for i from 0 to n-1:
        if nums[i] != i:
            return i
    
    return n
```

### Execution Visualization

### Example: nums = [3, 0, 1]
```
Cyclic Sort Process:

Initial: [3, 0, 1]
Indices: [0, 1, 2]
n = 3

Step 1: i=0, nums[0]=3
- Correct position for 3 is index 3
- But 3 >= n, can't place → i++

Step 2: i=1, nums[1]=0
- Correct position for 0 is index 0
- Swap nums[1] and nums[0]: [0, 3, 1]
- Still at i=1, nums[1]=3
- 3 >= n, can't place → i++

Step 3: i=2, nums[2]=1
- Correct position for 1 is index 1
- Swap nums[2] and nums[1]: [0, 1, 3]
- Still at i=2, nums[2]=3
- 3 >= n, can't place → i++

Final array: [0, 1, 3]

Missing Detection:
Index 0: nums[0] = 0 ✓ (should be 0)
Index 1: nums[1] = 1 ✓ (should be 1)
Index 2: nums[2] = 3 ✗ (should be 2) → Missing number is 2

Result: 2 ✓
```

### Key Visualization Points:
- **Index Mapping**: Number i → position i
- **Swap Strategy**: Keep swapping until correct or out of range
- **Range Filtering**: Only process numbers 0..n-1
- **Missing Detection**: First mismatch reveals answer

### Cyclic Sort Pattern Visualization:
```
Before: [3, 0, 1]
Target: [0, 1, 2] (if all numbers 0-2 present)

Process:
- 3 is out of range, skip
- 0 moves to index 0
- 1 moves to index 1

After: [0, 1, 3]
Missing: Position 2 should have 2, but has 3
Answer: 2
```

### Time Complexity Breakdown:
- **Cyclic Sort**: O(N) time, O(1) space - optimal for range 0..n-1
- **XOR Method**: O(N) time, O(1) space - elegant mathematical approach
- **Sum Method**: O(N) time, O(1) space - arithmetic approach
- **Hash Set**: O(N) time, O(N) space - simple but uses extra space

### Alternative Approaches:

#### 1. XOR Method (O(N) time, O(1) space)
```go
func missingNumberXOR(nums []int) int {
    n := len(nums)
    result := n
    
    for i, num := range nums {
        result ^= i ^ num
    }
    
    return result
}
```
- **Pros**: Elegant, O(1) space, no modification of input
- **Cons**: Less intuitive, requires understanding of XOR properties

#### 2. Sum Method (O(N) time, O(1) space)
```go
func missingNumberSum(nums []int) int {
    n := len(nums)
    expectedSum := n * (n + 1) / 2
    actualSum := 0
    
    for _, num := range nums {
        actualSum += num
    }
    
    return expectedSum - actualSum
}
```
- **Pros**: Simple arithmetic, easy to understand
- **Cons**: Potential integer overflow with large n

#### 3. Hash Set (O(N) time, O(N) space)
```go
func missingNumberHash(nums []int) int {
    n := len(nums)
    seen := make(map[int]bool)
    
    for _, num := range nums {
        seen[num] = true
    }
    
    for i := 0; i <= n; i++ {
        if !seen[i] {
            return i
        }
    }
    
    return n
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Uses O(N) extra space

### Extensions for Interviews:
- **Multiple Missing**: Find all missing numbers
- **Duplicate Detection**: Find repeated numbers
- **Set Mismatch**: Find swapped numbers in permutation
- **Range Queries**: Answer range-based queries efficiently
- **Real-world Applications**: Memory management, resource allocation
*/
func main() {
	// Test cases
	testCases := [][]int{
		{3, 0, 1},
		{0, 1},
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
		{0},
		{1},
		{},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
		{2, 0, 3, 1},
		{5, 4, 6, 0, 1, 2, 3},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		result := missingNumber(nums)
		fmt.Printf("Test Case %d: %v -> Missing number: %d\n", i+1, original, result)
	}
}
