package main

import "fmt"

// 448. Find All Numbers Disappeared in an Array
// Time: O(N), Space: O(1) - Cyclic Sort approach
func findDisappearedNumbers(nums []int) []int {
	i := 0
	n := len(nums)
	
	// Place each number in its correct position
	for i < n {
		correctPos := nums[i] - 1 // Numbers are from 1 to n
		if nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			i++
		}
	}
	
	// Find all positions where number doesn't match index+1
	var result []int
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Cyclic Sort for Missing Elements
- **Index-Value Mapping**: Place numbers at their correct indices
- **In-Place Rearrangement**: Sort without extra space
- **Multiple Missing**: Find all indices with wrong values
- **Range Constraints**: Leverage 1-to-n range property

## 2. PROBLEM CHARACTERISTICS
- **Multiple Missing**: Find all missing numbers from 1 to n
- **In-Place Operation**: Cannot use extra space
- **Range Utilization**: Numbers in [1, n] can be placed at indices [0, n-1]
- **Linear Time**: Single pass through array

## 3. SIMILAR PROBLEMS
- First Missing Positive (LeetCode 41) - Similar but find first missing
- Find the Duplicate Number (LeetCode 287) - Cyclic sort variant
- Missing Number (LeetCode 268) - XOR or cyclic sort
- Set Mismatch (LeetCode 645) - Find wrong and missing numbers

## 4. KEY OBSERVATIONS
- **Index Mapping**: Number i should be at position i-1
- **Range Property**: Only numbers 1 to n can be placed correctly
- **Missing Indicators**: Mismatches reveal missing numbers
- **Swap Strategy**: Keep swapping until current position is correct

## 5. VARIATIONS & EXTENSIONS
- **All Missing**: Find all missing numbers
- **Single Missing**: Find first missing positive
- **Duplicate Detection**: Find repeated numbers
- **Set Mismatch**: Find swapped numbers

## 6. INTERVIEW INSIGHTS
- Always clarify: "Array size constraints? Value ranges? Space limits?"
- Edge cases: empty array, all correct, no missing numbers
- Time complexity: O(N) for cyclic sort, O(N) for final scan
- Space complexity: O(1) - in-place operation
- Key insight: use array indices as hash map for numbers 1 to n

## 7. COMMON MISTAKES
- Wrong index calculation (i vs i-1 confusion)
- Not handling duplicates properly
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
- Some books are missing (the disappeared numbers)
- You place books one by one, swapping to correct positions
- Empty shelves reveal the missing book numbers
- Like a librarian finding which numbered books are missing

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers from 1 to n (some may be missing)
2. **Goal**: Find all missing numbers from 1 to n
3. **Constraints**: O(1) space, O(N) time
4. **Output**: List of missing numbers

#### Phase 2: Key Insight Recognition
- **"Index mapping natural"** → Number i belongs at position i-1
- **"Range utilization"** → Only numbers 1..n can be placed correctly
- **"In-place sorting"** → Use array itself as organizing structure
- **"Multiple missing"** → All mismatches reveal missing numbers

#### Phase 3: Strategy Development
```
Human thought process:
"I need all missing numbers without extra space.
Brute force: use hash set O(N) space.

Cyclic Sort Approach:
1. Place each number i at position i-1
2. Handle only numbers in range [1, n]
3. Skip duplicates (already in correct position)
4. Keep swapping until current position is correct
5. Scan for all positions where value != index+1
6. Those index+1 values are the missing numbers

This gives O(N) time, O(1) space!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty list (no numbers)
- **All correct**: Return empty list (no missing)
- **All missing**: Return all numbers 1..n
- **Duplicates**: Skip duplicates in placement

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [4, 3, 2, 7, 8, 2, 3, 1]

Human thinking:
"Cyclic Sort Process:
Step 1: i=0, nums[0]=4
Correct position for 4 is index 3
Swap nums[0] with nums[3]: [7, 3, 2, 4, 8, 2, 3, 1]
Still at i=0, nums[0]=7
Correct position for 7 is index 6
Swap nums[0] with nums[6]: [3, 3, 2, 4, 8, 2, 7, 1]
Still at i=0, nums[0]=3
Correct position for 3 is index 2
Swap nums[0] with nums[2]: [2, 3, 3, 4, 8, 2, 7, 1]
Still at i=0, nums[0]=2
Correct position for 2 is index 1
Swap nums[0] with nums[1]: [3, 2, 3, 4, 8, 2, 7, 1]
Still at i=0, nums[0]=3
Correct position for 3 is index 2
But nums[2]=3 (duplicate) → i++

Continue this process...
Final array: [1, 2, 3, 4, 3, 2, 7, 8]

Missing Detection:
Index 0: 1 ✓
Index 1: 2 ✓
Index 2: 3 ✓
Index 3: 4 ✓
Index 4: 3 ≠ 5 ✗ (missing 5)
Index 5: 2 ≠ 6 ✗ (missing 6)
Index 6: 7 ≠ 7 ✓ (actually 7 is correct)
Index 7: 8 ≠ 8 ✓ (actually 8 is correct)

Result: [5, 6] ✓"
```

#### Phase 6: Intuition Validation
- **Why index mapping**: Number i naturally belongs at position i-1
- **Why O(1) space**: Use array itself as organizing structure
- **Why O(N) time**: Each element moves at most once
- **Why correct**: All mismatches reveal missing numbers

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → O(N log N) vs O(N), uses extra space
2. **"Should I use hash set?"** → Yes, but violates O(1) space
3. **"What about duplicates?"** → Skip them, they're already handled
4. **"Can I use counting sort?"** → No, range is unbounded
5. **"Why swap instead of assign?"** → Need to preserve other values

### Real-World Analogy
**Like checking attendance in a classroom:**
- Students have assigned seat numbers 1 to n
- You ask students to sit in their correct seats
- Some students are absent (missing numbers)
- Empty seats reveal which students are missing
- You check each seat to see who's absent
- Like a teacher taking attendance by checking seat assignments

### Human-Readable Pseudocode
```
function findDisappearedNumbers(nums):
    n = length(nums)
    i = 0
    
    while i < n:
        correctPos = nums[i] - 1
        
        # Check if current number is in correct range
        # and not already in correct position
        if nums[i] != nums[correctPos]:
            swap(nums[i], nums[correctPos])
        else:
            i += 1
    
    # Find all positions with wrong numbers
    result = []
    for i from 0 to n-1:
        if nums[i] != i + 1:
            result.append(i + 1)
    
    return result
```

### Execution Visualization

### Example: nums = [4, 3, 2, 7, 8, 2, 3, 1]
```
Cyclic Sort Process:

Initial: [4, 3, 2, 7, 8, 2, 3, 1]
Indices: [0, 1, 2, 3, 4, 5, 6, 7]

After sorting: [1, 2, 3, 4, 3, 2, 7, 8]

Missing Detection:
Index 0: nums[0] = 1 ✓ (should be 1)
Index 1: nums[1] = 2 ✓ (should be 2)
Index 2: nums[2] = 3 ✓ (should be 3)
Index 3: nums[3] = 4 ✓ (should be 4)
Index 4: nums[4] = 3 ✗ (should be 5) → Missing 5
Index 5: nums[5] = 2 ✗ (should be 6) → Missing 6
Index 6: nums[6] = 7 ✓ (should be 7)
Index 7: nums[7] = 8 ✓ (should be 8)

Result: [5, 6] ✓
```

### Key Visualization Points:
- **Index Mapping**: Number i → position i-1
- **Swap Strategy**: Keep swapping until correct or duplicate
- **Range Filtering**: Only process numbers 1..n
- **Multiple Missing**: All mismatches reveal missing numbers

### Cyclic Sort Pattern Visualization:
```
Before: [4, 3, 2, 7, 8, 2, 3, 1]
Target: [1, 2, 3, 4, 5, 6, 7, 8] (if all numbers 1-8 present)

Process:
- 4 moves to index 3
- 3 moves to index 2
- 2 moves to index 1
- 7 moves to index 6
- 8 moves to index 7
- 1 moves to index 0
- Duplicates 2 and 3 stay where they are

After: [1, 2, 3, 4, 3, 2, 7, 8]
Missing: Positions 4 and 5 should have 5 and 6
Answer: [5, 6]
```

### Time Complexity Breakdown:
- **Cyclic Sort**: O(N) time, O(1) space - optimal for range 1..n
- **Hash Set**: O(N) time, O(N) space - simple but uses extra space
- **Boolean Array**: O(N) time, O(N) space - mark visited numbers
- **Bit Manipulation**: O(N) time, O(1) space - for specific cases

### Alternative Approaches:

#### 1. Hash Set (O(N) time, O(N) space)
```go
func findDisappearedNumbersHash(nums []int) []int {
    n := len(nums)
    seen := make(map[int]bool)
    
    for _, num := range nums {
        seen[num] = true
    }
    
    var result []int
    for i := 1; i <= n; i++ {
        if !seen[i] {
            result = append(result, i)
        }
    }
    
    return result
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Uses O(N) extra space

#### 2. Boolean Array (O(N) time, O(N) space)
```go
func findDisappearedNumbersBoolean(nums []int) []int {
    n := len(nums)
    present := make([]bool, n+1)
    
    for _, num := range nums {
        if num > 0 && num <= n {
            present[num] = true
        }
    }
    
    var result []int
    for i := 1; i <= n; i++ {
        if !present[i] {
            result = append(result, i)
        }
    }
    
    return result
}
```
- **Pros**: Clear logic, easy to understand
- **Cons**: Uses O(N) extra space

#### 3. In-Place Marking (O(N) time, O(1) space)
```go
func findDisappearedNumbersMarking(nums []int) []int {
    // Mark presence by making numbers negative
    for i := 0; i < len(nums); i++ {
        index := abs(nums[i]) - 1
        if nums[index] > 0 {
            nums[index] = -nums[index]
        }
    }
    
    var result []int
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            result = append(result, i+1)
        }
    }
    
    return result
}
```
- **Pros**: O(1) space, clever use of array
- **Cons**: Destructive to input, requires absolute value

### Extensions for Interviews:
- **Single Missing**: Find first missing positive
- **Duplicate Detection**: Find all duplicate numbers
- **Set Mismatch**: Find swapped numbers in permutation
- **Range Queries**: Answer range-based queries efficiently
- **Real-world Applications**: Inventory management, error detection
*/
func main() {
	// Test cases
	testCases := [][]int{
		{4, 3, 2, 7, 8, 2, 3, 1},
		{1, 1},
		{2, 2},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 1, 2, 2, 3, 3, 4, 4},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1},
		{},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		result := findDisappearedNumbers(nums)
		fmt.Printf("Test Case %d: %v -> Disappeared numbers: %v\n", i+1, original, result)
	}
}
