package main

import "fmt"

// 26. Remove Duplicates from Sorted Array
// Time: O(N), Space: O(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	uniqueIndex := 0
	
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}
	
	return uniqueIndex + 1
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 1, 1, 1},
		{},
		{2},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		length := removeDuplicates(nums)
		result := nums[:length]
		fmt.Printf("Test Case %d: %v -> Length: %d, Unique elements: %v\n", i+1, original, length, result)
	}
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two Pointers (Slow-Fast)
- **Slow Pointer**: Tracks the position of the last unique element
- **Fast Pointer**: Scans through the array looking for new unique elements
- **Key Insight**: Since array is sorted, duplicates are consecutive

## 2. PROBLEM CHARACTERISTICS
- **Sorted Input**: This is crucial - allows O(1) duplicate detection
- **In-place Modification**: Must modify the original array without extra space
- **Return Value**: Length of the deduplicated array prefix

## 3. SIMILAR PROBLEMS
- Remove Element (LeetCode 27)
- Move Zeroes (LeetCode 283)
- Remove Duplicates from Sorted Array II (LeetCode 80)
- Merge Sorted Array (LeetCode 88)

## 4. KEY OBSERVATIONS
- When array is sorted: nums[i] == nums[uniqueIndex] means duplicate
- When nums[i] != nums[uniqueIndex]: found new unique element
- No need for hash set or additional data structures due to sorting

## 5. VARIATIONS & EXTENSIONS
- Allow at most K duplicates: modify condition to count occurrences
- Remove duplicates from unsorted array: requires hash set
- Remove elements equal to target value: similar two-pointer approach

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is the array sorted?" (crucial assumption)
- Edge cases: empty array, single element, all duplicates
- Space complexity: O(1) because we modify in-place
- Time complexity: O(N) - single pass through array

## 7. COMMON MISTAKES
- Not handling empty array case
- Using extra space unnecessarily (hash set)
- Off-by-one errors in index calculations
- Not understanding that return value is the new length

## 8. OPTIMIZATION STRATEGIES
- The current solution is already optimal for sorted arrays
- For unsorted arrays, consider trade-offs between time and space
- For large datasets, consider cache-friendly access patterns

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing a bookshelf:**
- You have books arranged alphabetically (sorted array)
- You want to keep only one copy of each book
- You can't use extra boxes (must do it in-place)
- You use one finger to mark where the next unique book goes
- You scan with another finger to find new books

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Sorted array with possible duplicates
2. **Goal**: Remove duplicates, keep order, modify in-place
3. **Output**: New length of unique elements

#### Phase 2: Key Insight Recognition
- **"Wait, the array is sorted!"** → This is the game-changer
- **Sorted means**: All duplicates are grouped together
- **No need**: For hash sets or complex tracking
- **Can compare**: Current element with the last unique one

#### Phase 3: Strategy Development
```
Human thought process:
"I need to remember where the last unique element was placed.
Let's call this position 'uniqueIndex'.
I'll scan through the array with another pointer 'i'.
Whenever I find something different from nums[uniqueIndex],
I know it's a new unique element, so I'll place it right after
the last unique one and update uniqueIndex."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 immediately
- **Single element**: Return 1 (already unique)
- **All duplicates**: Return 1
- **All unique**: Return original length

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [1,1,2,2,3,4,4]

Human thinking:
"Okay, first element 1 is definitely unique. uniqueIndex = 0.
Next element is also 1 → duplicate, ignore it.
Next element is 2 → different from 1! This is new unique.
   Place it at position 1, update uniqueIndex to 1.
Next element is 2 → same as my last unique (2), ignore.
Next element is 3 → different from 2! New unique.
   Place it at position 2, update uniqueIndex to 2.
Next element is 4 → different from 3! New unique.
   Place it at position 3, update uniqueIndex to 3.
Next element is 4 → same as my last unique (4), ignore.

Done! I found 4 unique elements."
```

#### Phase 6: Intuition Validation
- **Why it works**: We never lose unique elements
- **Why it's efficient**: Single pass, constant space
- **Why it's correct**: We only advance when we find something truly new

### Common Human Pitfalls & How to Avoid Them
1. **"Should I use a hash set?"** → No! Array is sorted, that's the key
2. **"Do I need to shift elements?"** → No, just overwrite duplicates
3. **"What about the return value?"** → It's the count, not the array
4. **"Do I need to clean up the rest?"** → No, elements after length don't matter

### Real-World Analogy
**Like removing duplicate names from a sorted attendance sheet:**
- You keep one marker at the last unique name
- You scan down the list with a pencil
- When you find a new name, you write it right below the last unique one
- Duplicates get overwritten, but you never lose unique names

### Human-Readable Pseudocode
```
function removeDuplicates(sortedArray):
    if array is empty: return 0
    
    lastUniquePosition = 0
    
    for each element from position 1 to end:
        if current element is different from element at lastUniquePosition:
            move lastUniquePosition forward by 1
            place current element at this new position
    
    return lastUniquePosition + 1
```

### Execution Visualization

### Example: [1,1,2,2,3,4,4]
```
Initial: nums = [1,1,2,2,3,4,4], uniqueIndex = 0

Step 1: i=1, nums[1]=1, nums[0]=1
→ nums[i] == nums[uniqueIndex] (duplicate)
→ No change, uniqueIndex remains 0
State: [1,1,2,2,3,4,4], uniqueIndex=0

Step 2: i=2, nums[2]=2, nums[0]=1  
→ nums[i] != nums[uniqueIndex] (new unique found)
→ uniqueIndex++ (1), nums[1] = nums[2] (2)
State: [1,2,2,2,3,4,4], uniqueIndex=1

Step 3: i=3, nums[3]=2, nums[1]=2
→ nums[i] == nums[uniqueIndex] (duplicate)
→ No change, uniqueIndex remains 1
State: [1,2,2,2,3,4,4], uniqueIndex=1

Step 4: i=4, nums[4]=3, nums[1]=2
→ nums[i] != nums[uniqueIndex] (new unique found)
→ uniqueIndex++ (2), nums[2] = nums[4] (3)
State: [1,2,3,2,3,4,4], uniqueIndex=2

Step 5: i=5, nums[5]=4, nums[2]=3
→ nums[i] != nums[uniqueIndex] (new unique found)
→ uniqueIndex++ (3), nums[3] = nums[5] (4)
State: [1,2,3,4,3,4,4], uniqueIndex=3

Step 6: i=6, nums[6]=4, nums[3]=4
→ nums[i] == nums[uniqueIndex] (duplicate)
→ No change, uniqueIndex remains 3
State: [1,2,3,4,3,4,4], uniqueIndex=3

Final: Return uniqueIndex + 1 = 4
Result prefix: nums[:4] = [1,2,3,4]
```

### Key Visualization Points:
- **uniqueIndex** always points to the last unique element found
- **Array modification** happens in-place, overwriting duplicates
- **Return value** is the length of the unique prefix
- **Elements beyond** the returned length are irrelevant

### Memory Layout Visualization:
```
Memory Address: [0][1][2][3][4][5][6]
Initial State:  [1][1][2][2][3][4][4]
After Process: [1][2][3][4][3][4][4]
                ^^^^^^^^^^^
                Valid prefix (length=4)
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited once
- **Constant Space**: O(1) - only two integer variables
- **In-place**: No additional arrays allocated
*/
