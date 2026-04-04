package main

import "fmt"

// 41. First Missing Positive
// Time: O(N), Space: O(1) - Cyclic Sort approach
func firstMissingPositive(nums []int) int {
	n := len(nums)
	
	// Place each positive number in its correct position
	i := 0
	for i < n {
		correctPos := nums[i] - 1
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			i++
		}
	}
	
	// Find the first position where number doesn't match index+1
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	
	return n + 1 // All numbers are in correct positions
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Cyclic Sort for Missing Elements
- **Index-Value Mapping**: Place numbers at their correct indices
- **In-Place Rearrangement**: Sort without extra space
- **Missing Detection**: Find first index with wrong value
- **Range Constraints**: Leverage 1-to-n range property

## 2. PROBLEM CHARACTERISTICS
- **Missing Positive**: Find smallest positive integer not in array
- **In-Place Operation**: Cannot use extra space
- **Range Utilization**: Numbers in [1, n] can be placed at indices [0, n-1]
- **Linear Time**: Single pass through array

## 3. SIMILAR PROBLEMS
- Find All Numbers Disappeared in an Array (LeetCode 448) - Same technique
- Find the Duplicate Number (LeetCode 287) - Cyclic sort variant
- Missing Number (LeetCode 268) - XOR or cyclic sort
- Set Mismatch (LeetCode 645) - Find wrong and missing numbers

## 4. KEY OBSERVATIONS
- **Index Mapping**: Number i should be at position i-1
- **Range Property**: Only numbers 1 to n can be placed correctly
- **Missing Indicator**: First mismatch reveals missing positive
- **Swap Strategy**: Keep swapping until current position is correct

## 5. VARIATIONS & EXTENSIONS
- **Missing Element**: Find first missing positive
- **Duplicate Detection**: Find repeated numbers
- **Multiple Missing**: Find all missing numbers
- **Set Mismatch**: Find swapped numbers

## 6. INTERVIEW INSIGHTS
- Always clarify: "Array size constraints? Value ranges? Space limits?"
- Edge cases: empty array, all correct, no positive numbers
- Time complexity: O(N) for cyclic sort, O(N) for final scan
- Space complexity: O(1) - in-place operation
- Key insight: use array indices as hash map for numbers 1 to n

## 7. COMMON MISTAKES
- Wrong index calculation (i vs i-1 confusion)
- Not handling negative numbers properly
- Missing edge cases for empty/full arrays
- Infinite loop with duplicate values
- Wrong bounds checking

## 8. OPTIMIZATION STRATEGIES
- **Cyclic Sort**: O(N) time, O(1) space - optimal for range 1..n
- **Hash Set**: O(N) time, O(N) space - simple but uses extra space
- **Boolean Array**: O(N) time, O(N) space - mark visited numbers
- **Bit Manipulation**: O(N) time, O(1) space - for specific cases

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing numbered books on a shelf:**
- You have books numbered 1 to n and n shelf positions
- Each book should go to its matching numbered shelf
- You place books one by one, swapping to correct positions
- The first empty shelf with wrong book reveals the missing number
- Like a librarian organizing books by their call numbers

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers (can be negative, zero, positive)
2. **Goal**: Find smallest positive integer not in array
3. **Constraints**: O(1) space, O(N) time
4. **Output**: Missing positive integer

#### Phase 2: Key Insight Recognition
- **"Index mapping natural"** → Number i belongs at position i-1
- **"Range utilization"** → Only numbers 1..n can be placed correctly
- **"In-place sorting"** → Use array itself as organizing structure
- **"Missing detection"** → First position with wrong value reveals answer

#### Phase 3: Strategy Development
```
Human thought process:
"I need smallest missing positive without extra space.
Brute force: use hash set O(N) space.

Cyclic Sort Approach:
1. Place each number i at position i-1
2. Handle only numbers in range [1, n]
3. Skip negatives, zeros, and numbers > n
4. Keep swapping until current position is correct
5. Scan for first position where value != index+1
6. That index+1 is the missing positive

This gives O(N) time, O(1) space!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 1 (smallest positive)
- **All correct**: Return n+1 (next positive)
- **No positives**: Return 1 (smallest positive)
- **Duplicates**: Skip duplicates in placement

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [3, 4, -1, 1]

Human thinking:
"Cyclic Sort Process:
Step 1: i=0, nums[0]=3
Correct position for 3 is index 2
Swap nums[0] with nums[2]: [-1, 4, 3, 1]
Still at i=0, nums[0]=-1 (out of range) → i++

Step 2: i=1, nums[1]=4
Correct position for 4 is index 3
Swap nums[1] with nums[3]: [-1, 1, 3, 4]
Still at i=1, nums[1]=1
Correct position for 1 is index 0
Swap nums[1] with nums[0]: [1, -1, 3, 4]
Still at i=1, nums[1]=-1 (out of range) → i++

Step 3: i=2, nums[2]=3 (correct position) → i++
Step 4: i=3, nums[3]=4 (correct position) → i++

Final array: [1, -1, 3, 4]
Scan for mismatch:
Index 0: 1 ✓ (correct)
Index 1: -1 ≠ 2 ✗ (missing positive is 2)

Result: 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why index mapping**: Number i naturally belongs at position i-1
- **Why O(1) space**: Use array itself as organizing structure
- **Why O(N) time**: Each element moves at most once
- **Why correct**: First mismatch reveals smallest missing

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → O(N log N) vs O(N), uses extra space
2. **"Should I use hash set?"** → Yes, but violates O(1) space
3. **"What about negative numbers?"** → Skip them, they're irrelevant
4. **"Can I use counting sort?"** → No, range is unbounded
5. **"Why swap instead of assign?"** → Need to preserve other values

### Real-World Analogy
**Like assigning seats in a classroom:**
- Students have assigned seat numbers 1 to n
- You place each student in their correct seat
- If seat is occupied, you swap with the student there
- Students without assigned numbers (negatives) stand aside
- The first empty seat number is the missing student's number
- Like a teacher organizing students by their assigned seats

### Human-Readable Pseudocode
```
function firstMissingPositive(nums):
    n = length(nums)
    i = 0
    
    while i < n:
        correctPos = nums[i] - 1
        
        # Check if current number is in correct range
        # and not already in correct position
        if nums[i] > 0 and nums[i] <= n and nums[i] != nums[correctPos]:
            swap(nums[i], nums[correctPos])
        else:
            i += 1
    
    # Find first position with wrong number
    for i from 0 to n-1:
        if nums[i] != i + 1:
            return i + 1
    
    return n + 1
```

### Execution Visualization

### Example: nums = [3, 4, -1, 1]
```
Cyclic Sort Process:

Initial: [3, 4, -1, 1]
Indices: [0, 1, 2, 3]

Step 1: i=0, nums[0]=3
- Correct position for 3 is index 2
- Swap nums[0] and nums[2]: [-1, 4, 3, 1]
- Still at i=0, nums[0]=-1 (out of range) → i++

Step 2: i=1, nums[1]=4
- Correct position for 4 is index 3
- Swap nums[1] and nums[3]: [-1, 1, 3, 4]
- Still at i=1, nums[1]=1
- Correct position for 1 is index 0
- Swap nums[1] and nums[0]: [1, -1, 3, 4]
- Still at i=1, nums[1]=-1 (out of range) → i++

Step 3: i=2, nums[2]=3 (already correct) → i++
Step 4: i=3, nums[3]=4 (already correct) → i++

Final array: [1, -1, 3, 4]

Missing Detection:
Index 0: nums[0] = 1 ✓ (should be 1)
Index 1: nums[1] = -1 ✗ (should be 2) → Missing positive is 2

Result: 2 ✓
```

### Key Visualization Points:
- **Index Mapping**: Number i → position i-1
- **Swap Strategy**: Keep swapping until correct or out of range
- **Range Filtering**: Only process numbers 1..n
- **Missing Detection**: First mismatch reveals answer

### Cyclic Sort Pattern Visualization:
```
Before: [3, 4, -1, 1]
Target: [1, 2, 3, 4] (if all numbers 1-4 present)

Process:
- 3 moves to index 2
- 4 moves to index 3  
- 1 moves to index 0
- -1 stays (out of range)

After: [1, -1, 3, 4]
Missing: Position 1 should have 2, but has -1
Answer: 2
```

### Time Complexity Breakdown:
- **Cyclic Sort**: O(N) time, O(1) space - optimal for range 1..n
- **Hash Set**: O(N) time, O(N) space - simple but uses extra space
- **Boolean Array**: O(N) time, O(N) space - mark visited numbers
- **Bit Manipulation**: O(N) time, O(1) space - for specific cases

### Alternative Approaches:

#### 1. Hash Set (O(N) time, O(N) space)
```go
func firstMissingPositiveHash(nums []int) int {
    seen := make(map[int]bool)
    
    for _, num := range nums {
        if num > 0 {
            seen[num] = true
        }
    }
    
    for i := 1; i <= len(nums)+1; i++ {
        if !seen[i] {
            return i
        }
    }
    
    return len(nums) + 1
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Uses O(N) extra space

#### 2. Boolean Array (O(N) time, O(N) space)
```go
func firstMissingPositiveBoolean(nums []int) int {
    n := len(nums)
    present := make([]bool, n+1)
    
    for _, num := range nums {
        if num > 0 && num <= n {
            present[num] = true
        }
    }
    
    for i := 1; i <= n; i++ {
        if !present[i] {
            return i
        }
    }
    
    return n + 1
}
```
- **Pros**: Clear logic, easy to understand
- **Cons**: Uses O(N) extra space

#### 3. In-Place Marking (O(N) time, O(1) space)
```go
func firstMissingPositiveMarking(nums []int) int {
    n := len(nums)
    
    // Mark numbers out of range
    for i := 0; i < n; i++ {
        if nums[i] <= 0 || nums[i] > n {
            nums[i] = n + 1
        }
    }
    
    // Mark presence by using negative values
    for i := 0; i < n; i++ {
        num := abs(nums[i])
        if num <= n {
            nums[num-1] = -abs(nums[num-1])
        }
    }
    
    // Find first positive index
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            return i + 1
        }
    }
    
    return n + 1
}
```
- **Pros**: O(1) space, clever use of array
- **Cons**: More complex, destructive to input

### Extensions for Interviews:
- **Multiple Missing**: Find all missing positive numbers
- **Duplicate Detection**: Find all duplicate numbers
- **Set Mismatch**: Find swapped numbers in permutation
- **Range Queries**: Answer range-based queries efficiently
- **Real-world Applications**: Memory management, resource allocation
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
		{1, 2, 3},
		{2, 3, 4},
		{-1, -2, -3},
		{0},
		{1},
		{2},
		{3, 4, 2, 1},
		{1, 1},
		{2, 2},
		{1, 2, 6, 3, 5, 4},
		{1, 2, 0, 2, 1},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		result := firstMissingPositive(nums)
		fmt.Printf("Test Case %d: %v -> First missing positive: %d\n", i+1, original, result)
	}
}
