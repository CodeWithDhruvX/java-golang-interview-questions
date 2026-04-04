package main

import "fmt"

// 283. Move Zeroes
// Time: O(N), Space: O(1)
func moveZeroes(nums []int) {
	lastNonZeroFoundAt := 0
	
	// Move all non-zero elements to the front
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}
	
	// Fill the remaining positions with zeros
	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	// Test cases
	testCases := [][]int{
		{0, 1, 0, 3, 12},
		{0},
		{1, 2, 3, 4},
		{0, 0, 1, 0, 2, 0, 3},
		{4, 0, 5, 0, 3, 0, 1},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		moveZeroes(nums)
		fmt.Printf("Test Case %d: %v -> After moving zeroes: %v\n", i+1, original, nums)
	}
}



/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two Pointers (Position-Scanner)
- **Position Pointer**: Tracks where the next non-zero element should be placed
- **Scanner Pointer**: Scans through the array to find non-zero elements
- **Key Insight**: Two-pass approach - first collect non-zeros, then fill zeros

## 2. PROBLEM CHARACTERISTICS
- **In-place Modification**: Must modify the original array without extra space
- **Order Preservation**: Non-zero elements must maintain their relative order
- **Zero Targeting**: Specifically moving zeros to the end (not general element removal)

## 3. SIMILAR PROBLEMS
- Remove Duplicates from Sorted Array (LeetCode 26)
- Remove Element (LeetCode 27)
- Move Zeroes (LeetCode 283) - Current problem
- Sort Colors (LeetCode 75) - Dutch National Flag

## 4. KEY OBSERVATIONS
- Zeros are the "targets" to be moved
- Non-zero elements must preserve order (stable)
- Two-pass approach is clean and efficient
- Can also be done in one pass with swapping

## 5. VARIATIONS & EXTENSIONS
- Move all even numbers to front: similar logic
- Move negative numbers to end: change condition
- General partition: move elements satisfying condition to front
- One-pass solution: swap when non-zero found

## 6. INTERVIEW INSIGHTS
- Always clarify: "Should relative order of non-zeros be preserved?"
- Edge cases: empty array, no zeros, all zeros
- Space complexity: O(1) because we modify in-place
- Time complexity: O(N) - single or double pass through array

## 7. COMMON MISTAKES
- Not preserving order of non-zero elements
- Using extra space unnecessarily (new array)
- Off-by-one errors in index calculations
- Forgetting to fill remaining positions with zeros
- Not handling edge cases (empty array, single element)

## 8. OPTIMIZATION STRATEGIES
- Current two-pass solution is already optimal
- One-pass solution with swapping reduces passes but may increase swaps
- For large datasets, consider cache-friendly access patterns
- Trade-off between number of passes vs number of swaps

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing a bookshelf:**
- You have books with some empty spaces (zeros) scattered throughout
- You want to move all the empty spaces to the end
- You must keep the books in their original order
- You can't use extra boxes (must do it in-place)
- First, you slide all books forward to fill gaps
- Then, you mark all remaining spaces as empty

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array with zeros and non-zero elements
2. **Goal**: Move all zeros to the end, keep non-zero order
3. **Output**: Same array modified in-place

#### Phase 2: Key Insight Recognition
- **"Two-phase approach"** → First collect non-zeros, then add zeros
- **Order matters** → Can't just swap arbitrarily
- **Zeros are special** → They're the only elements being moved
- **Position tracking** → Need to remember where to place next non-zero

#### Phase 3: Strategy Development
```
Human thought process:
"I need to gather all the non-zero elements first, keeping their order.
I'll use a pointer 'lastNonZeroFoundAt' to track where the next
non-zero should go. I'll scan through the array and copy each
non-zero element to this position, then advance the pointer.
After I've collected all non-zeros, I'll fill the rest with zeros."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Do nothing
- **No zeros**: Array remains unchanged
- **All zeros**: All positions become zeros (already are)
- **Single element**: Either zero or non-zero, handled correctly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [0, 1, 0, 3, 12]

Human thinking:
"Okay, I need to move all non-zeros to the front.
I'll start with position 0 for the first non-zero I find.

Scanning:
Position 0: It's 0 → skip, keep looking
Position 1: It's 1 → non-zero! Place it at position 0
           Now position 1 is ready for next non-zero
Position 2: It's 0 → skip
Position 3: It's 3 → non-zero! Place it at position 1
           Now position 2 is ready for next non-zero  
Position 4: It's 12 → non-zero! Place it at position 2
           Now position 3 is ready

Non-zeros collected: [1, 3, 12] at positions [0, 1, 2]
Now fill positions 3 and 4 with zeros.

Final result: [1, 3, 12, 0, 0]"
```

#### Phase 6: Intuition Validation
- **Why it works**: We never lose non-zero elements, just move them
- **Why it's efficient**: Two simple linear passes
- **Why it's correct**: Order is preserved because we copy sequentially

### Common Human Pitfalls & How to Avoid Them
1. **"Should I swap zeros with non-zeros?"** → No, that would break order
2. **"Do I need to count zeros first?"** → No, just fill remaining positions
3. **"What if I use one pass?"** → Possible but more complex with swapping
4. **"Do I need to return something?"** → No, modify in-place

### Real-World Analogy
**Like organizing a parking lot:**
- Cars (non-zeros) are scattered with empty spaces (zeros)
- You want all cars at the front, empty spaces at the back
- You can't change the order cars arrived (preserve order)
- First, you drive all cars forward to fill gaps
- Then, you mark all remaining spots as empty

### Human-Readable Pseudocode
```
function moveZeroes(array):
    nextNonZeroPosition = 0
    
    // First pass: collect all non-zero elements
    for each element in array:
        if element is not zero:
            place element at nextNonZeroPosition
            increment nextNonZeroPosition
    
    // Second pass: fill remaining positions with zeros
    for position from nextNonZeroPosition to end:
        set array[position] = 0
```

### Execution Visualization

### Example: [0, 1, 0, 3, 12]
```
Initial: nums = [0, 1, 0, 3, 12], lastNonZeroFoundAt = 0

=== PASS 1: Collect Non-Zeros ===

Step 1: i=0, nums[0]=0
→ It's zero, skip
State: [0, 1, 0, 3, 12], lastNonZeroFoundAt=0

Step 2: i=1, nums[1]=1
→ Non-zero found! nums[0] = nums[1] (1)
→ lastNonZeroFoundAt++ (1)
State: [1, 1, 0, 3, 12], lastNonZeroFoundAt=1

Step 3: i=2, nums[2]=0
→ It's zero, skip
State: [1, 1, 0, 3, 12], lastNonZeroFoundAt=1

Step 4: i=3, nums[3]=3
→ Non-zero found! nums[1] = nums[3] (3)
→ lastNonZeroFoundAt++ (2)
State: [1, 3, 0, 3, 12], lastNonZeroFoundAt=2

Step 5: i=4, nums[4]=12
→ Non-zero found! nums[2] = nums[4] (12)
→ lastNonZeroFoundAt++ (3)
State: [1, 3, 12, 3, 12], lastNonZeroFoundAt=3

=== PASS 2: Fill Zeros ===

Step 6: i=3, fill with zero
State: [1, 3, 12, 0, 12]

Step 7: i=4, fill with zero
State: [1, 3, 12, 0, 0]

Final: [1, 3, 12, 0, 0]
```

### Key Visualization Points:
- **lastNonZeroFoundAt** always points to where next non-zero should go
- **First pass** compacts all non-zeros at the front
- **Second pass** fills the tail with zeros
- **Order preservation** happens because we copy sequentially

### Memory Layout Visualization:
```
Memory Address: [0][1][2][3][4]
Initial State:  [0][1][0][3][12]
After Pass 1:  [1][3][12][3][12]  (non-zeros compacted)
After Pass 2:  [1][3][12][0][0]   (zeros filled)
                ^^^^^^^
                Non-zeros (order preserved)
```

### Time Complexity Breakdown:
- **First Pass**: O(N) - collect non-zeros
- **Second Pass**: O(N) - fill zeros
- **Total**: O(N) - linear time
- **Space**: O(1) - constant extra space
- **Stability**: Yes - preserves relative order

### Alternative One-Pass Approach:
```
for i := 0; i < len(nums); i++ {
    if nums[i] != 0 {
        nums[i], nums[lastNonZeroFoundAt] = nums[lastNonZeroFoundAt], nums[i]
        lastNonZeroFoundAt++
    }
}
```
- **Pros**: Single pass, potentially fewer operations
- **Cons**: More swaps, slightly less intuitive
- **Trade-off**: Passes vs swaps
*/