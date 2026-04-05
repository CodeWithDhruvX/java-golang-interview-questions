import java.util.Arrays;

public class MoveZeroes {
    
    // 283. Move Zeroes
    // Time: O(N), Space: O(1)
    public static void moveZeroes(int[] nums) {
        int lastNonZeroFoundAt = 0;
        
        // Move all non-zero elements to the front
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != 0) {
                nums[lastNonZeroFoundAt] = nums[i];
                lastNonZeroFoundAt++;
            }
        }
        
        // Fill the remaining positions with zeros
        for (int i = lastNonZeroFoundAt; i < nums.length; i++) {
            nums[i] = 0;
        }
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {0, 1, 0, 3, 12},
            {0},
            {1, 2, 3, 4},
            {0, 0, 1, 0, 2, 0, 3},
            {4, 0, 5, 0, 3, 0, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = testCases[i].clone();
            int[] original = nums.clone();
            
            moveZeroes(nums);
            System.out.printf("Test Case %d: %s -> After moving zeroes: %s\n", 
                i + 1, Arrays.toString(original), Arrays.toString(nums));
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two Pointers (Collect and Fill)
- **First Pointer**: Collects non-zero elements at the front
- **Second Phase**: Fills remaining positions with zeros
- **Key Insight**: Separate collection of non-zeros from zero placement

## 2. PROBLEM CHARACTERISTICS
- **In-place Modification**: Must modify original array without extra space
- **Order Preservation**: Non-zero elements must maintain relative order
- **Zero Placement**: All zeros must be moved to the end
- **Return Value**: Void function (array modified in-place)

## 3. SIMILAR PROBLEMS
- Remove Duplicates from Sorted Array (LeetCode 26)
- Remove Element (LeetCode 27)
- Sort Colors (LeetCode 75)
- Move Zeros variations with different constraints

## 4. KEY OBSERVATIONS
- Non-zero elements maintain their original relative order
- Zeros can be placed anywhere after all non-zeros
- Two-phase approach: collect non-zeros, then fill zeros
- No need for swapping during collection phase

## 5. VARIATIONS & EXTENSIONS
- Minimize swaps: only swap when necessary
- Count zeros first, then rearrange
- Use single pass with swapping
- Handle different data types (strings, objects)

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I maintain the relative order of non-zeros?"
- Edge cases: all zeros, no zeros, single element
- Space complexity: O(1) because we modify in-place
- Time complexity: O(N) - two passes through array

## 7. COMMON MISTAKES
- Not maintaining relative order of non-zeros
- Using extra space unnecessarily
- Forgetting to fill remaining positions with zeros
- Off-by-one errors in index calculations

## 8. OPTIMIZATION STRATEGIES
- Current solution is optimal for clarity
- Can optimize to single pass with swapping
- For large arrays, consider cache efficiency
- For sparse zeros, counting approach might be better

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing a bookshelf:**
- You want to move all empty books (zeros) to the end
- You first collect all non-empty books and place them at the front
- Then you fill the remaining empty spots with empty books
- Order of non-empty books must be preserved

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array with zeros and non-zeros mixed
2. **Goal**: Move all zeros to the end, keep non-zero order
3. **Output**: Modified array (void return)

#### Phase 2: Key Insight Recognition
- **"I can do this in two phases!"** → Separate concerns
- **Phase 1**: Collect all non-zero elements at the front
- **Phase 2**: Fill the rest with zeros
- **No need**: For complex swapping or additional arrays

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use a pointer 'lastNonZeroFoundAt' to track where the next
non-zero element should go. I'll scan through the array, and whenever
I find a non-zero, I'll place it at this position and advance the pointer.
After collecting all non-zeros, I'll fill the rest with zeros."
```

#### Phase 4: Edge Case Handling
- **All zeros**: Array becomes all zeros (no change)
- **No zeros**: Array remains unchanged
- **Single element**: Either stays the same or becomes zero
- **Empty array**: No operation needed

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [0, 1, 0, 3, 12]

Human thinking:
"Okay, lastNonZeroFoundAt starts at 0.
Position 0: It's 0 → skip it.
Position 1: It's 1 → place it at position 0, advance to 1.
           Array: [1, 1, 0, 3, 12]
Position 2: It's 0 → skip it.
Position 3: It's 3 → place it at position 1, advance to 2.
           Array: [1, 3, 0, 3, 12]
Position 4: It's 12 → place it at position 2, advance to 3.
           Array: [1, 3, 12, 3, 12]

Now fill zeros from position 3 onwards:
Position 3: Set to 0 → [1, 3, 12, 0, 12]
Position 4: Set to 0 → [1, 3, 12, 0, 0]

Done! All zeros moved to the end."
```

#### Phase 6: Intuition Validation
- **Why it works**: We preserve order by placing non-zeros sequentially
- **Why it's efficient**: Two simple linear passes
- **Why it's correct**: All non-zeros collected first, then zeros fill rest

### Common Human Pitfalls & How to Avoid Them
1. **"Should I swap in-place?"** → Two-phase approach is clearer
2. **"Do I need to track zeros?"** → No, just fill remaining positions
3. **"What about order?"** → Collection phase preserves order automatically
4. **"Can I optimize further?"** → Yes, but clarity is often better

### Real-World Analogy
**Like organizing a parking lot:**
- You want to move all empty parking spaces (zeros) to the back
- First, you move all occupied cars to the front spaces
- Then you mark the remaining spaces as empty
- The order of cars is preserved during the move

### Human-Readable Pseudocode
```
function moveZeroes(array):
    nonZeroPosition = 0
    
    // Phase 1: Collect non-zeros
    for each element in array:
        if element is not zero:
            array[nonZeroPosition] = element
            nonZeroPosition++
    
    // Phase 2: Fill zeros
    for i from nonZeroPosition to end of array:
        array[i] = 0
```

### Execution Visualization

### Example: [0, 1, 0, 3, 12]
```
Initial: nums = [0, 1, 0, 3, 12], lastNonZeroFoundAt = 0

Phase 1: Collect Non-Zeros
Step 1: i=0, nums[0]=0 → skip, lastNonZeroFoundAt=0
State: [0, 1, 0, 3, 12], lastNonZeroFoundAt=0

Step 2: i=1, nums[1]=1 → place at position 0, lastNonZeroFoundAt=1
State: [1, 1, 0, 3, 12], lastNonZeroFoundAt=1

Step 3: i=2, nums[2]=0 → skip, lastNonZeroFoundAt=1
State: [1, 1, 0, 3, 12], lastNonZeroFoundAt=1

Step 4: i=3, nums[3]=3 → place at position 1, lastNonZeroFoundAt=2
State: [1, 3, 0, 3, 12], lastNonZeroFoundAt=2

Step 5: i=4, nums[4]=12 → place at position 2, lastNonZeroFoundAt=3
State: [1, 3, 12, 3, 12], lastNonZeroFoundAt=3

Phase 2: Fill Zeros
Step 6: i=3, set to 0 → [1, 3, 12, 0, 12]
Step 7: i=4, set to 0 → [1, 3, 12, 0, 0]

Final: [1, 3, 12, 0, 0]
```

### Key Visualization Points:
- **lastNonZeroFoundAt** tracks where next non-zero goes
- **First phase** overwrites positions with non-zeros
- **Second phase** fills remaining with zeros
- **Order preservation** happens automatically in first phase

### Memory Layout Visualization:
```
Memory Address: [0][1][2][3][4]
Initial State:  [0][1][0][3][12]
After Phase 1: [1][3][12][3][12]
After Phase 2: [1][3][12][0][0]
               ^^^^^^^^^^^
               Non-zeros (preserved order)
```

### Time Complexity Breakdown:
- **Phase 1**: O(N) - single pass to collect non-zeros
- **Phase 2**: O(N) - single pass to fill zeros
- **Total**: O(N) - two linear passes
- **Space**: O(1) - only one integer variable
*/
