import java.util.Arrays;

public class BinarySearch {
    
    // 704. Binary Search
    // Time: O(log N), Space: O(1)
    public static int search(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return -1;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {-1, 0, 3, 5, 9, 12},
            {-1, 0, 3, 5, 9, 12},
            {1, 2, 3, 4, 5},
            {1, 2, 3, 4, 5},
            {},
            {1},
            {1},
            {-10, -5, 0, 5, 10},
            {-10, -5, 0, 5, 10},
            {2, 4, 6, 8, 10}
        };
        
        int[] targets = {9, 2, 3, 6, 1, 1, 0, -5, 0, 7};
        
        for (int i = 0; i < testArrays.length; i++) {
            int result = search(testArrays[i], targets[i]);
            System.out.printf("Test Case %d: nums=%s, target=%d -> Index: %d\n", 
                i + 1, Arrays.toString(testArrays[i]), targets[i], result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search with Range Tracking
- **Binary Search**: Divide and conquer approach
- **Range Tracking**: Maintain left and right boundaries
- **Mid Calculation**: Prevent overflow with left + (right-left)/2
- **Target Comparison**: Adjust search range based on comparison

## 2. PROBLEM CHARACTERISTICS
- **Sorted Array**: Input array is sorted in ascending order
- **Single Target**: Need to find one specific element
- **Index Return**: Return position or -1 if not found
- **No Duplicates**: Each element appears at most once

## 3. SIMILAR PROBLEMS
- Search Insert Position
- Find First and Last Position of Element
- Search in Rotated Sorted Array
- Find Minimum in Rotated Sorted Array

## 4. KEY OBSERVATIONS
- Sorted array enables binary search O(log N)
- Mid calculation prevents integer overflow
- Can eliminate half of remaining elements each step
- Left/right pointers track current search range

## 5. VARIATIONS & EXTENSIONS
- Find first/last occurrence (with duplicates)
- Search in descending sorted array
- Find closest element to target
- Multiple queries on same array

## 6. INTERVIEW INSIGHTS
- Clarify: "Is array sorted in ascending order?"
- Clarify: "Can array contain duplicates?"
- Edge cases: empty array, single element, target not found
- Why binary search is better than linear search

## 7. COMMON MISTAKES
- Integer overflow in mid calculation
- Off-by-one errors in loop condition
- Not handling empty array case
- Wrong comparison direction for ascending order

## 8. OPTIMIZATION STRATEGIES
- Current solution is optimal
- For multiple queries, array stays sorted
- For cache efficiency, consider branch prediction
- For very large arrays, consider interpolation search

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a name in a phone book:**
- You have a phone book sorted alphabetically (sorted array)
- You want to find a specific person's name (target)
- Instead of reading every page (linear search), you open to the middle
- Based on whether your target comes before or after, you eliminate half
- You keep narrowing down until you find the exact page

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Sorted array and target value
2. **Goal**: Find index of target in array
3. **Output**: Index or -1 if not found

#### Phase 2: Key Insight Recognition
- **"Array is sorted!"** → This is the key insight
- **"Can eliminate half!"** → Binary search principle
- **"How to find middle?"** → left + (right-left)/2
- **"When to stop?"** → When left > right

#### Phase 3: Strategy Development
```
Human thought process:
"I'll maintain left and right pointers for my search range.
I'll repeatedly look at the middle element:
- If it's my target, I'm done!
- If target is smaller, search left half
- If target is larger, search right half
I'll continue until I find it or exhaust the search space."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return -1 immediately
- **Single element**: Check if it matches target
- **Target not found**: Will exit when left > right
- **Integer overflow**: Use left + (right-left)/2 formula

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [-1, 0, 3, 5, 9, 12], target = 9

Human thinking:
"Initialize: left = 0, right = 5

Step 1: mid = 0 + (5-0)/2 = 2
           nums[2] = 3, target = 9
           3 < 9, so target is in right half
           left = 3, right = 5

Step 2: mid = 3 + (5-3)/2 = 4
           nums[4] = 9, target = 9
           Found it! Return index 4"
```

#### Phase 6: Intuition Validation
- **Why it works**: Each step eliminates half the search space
- **Why it's efficient**: Logarithmic time complexity
- **Why it's correct**: Systematic elimination guarantees finding target if exists

### Common Human Pitfalls & How to Avoid Them
1. **"Why not linear search?"** → Too slow O(N)
2. **"What about mid calculation?"** → Prevent overflow with proper formula
3. **"When to stop?"** → When left > right (not left >= right)
4. **"What about duplicates?"** → Standard binary search assumes unique elements

### Real-World Analogy
**Like finding a book in a library:**
- You have books arranged by Dewey Decimal number (sorted array)
- You want to find a specific book (target)
- You go to the middle shelf and section
- If your book's number is higher, you go right
- If lower, you go left
- You keep narrowing down until you find the exact shelf

### Human-Readable Pseudocode
```
function binarySearch(sortedArray, target):
    left = 0
    right = sortedArray.length - 1
    
    while left <= right:
        mid = left + (right - left) / 2
        
        if sortedArray[mid] == target:
            return mid
        else if sortedArray[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return -1
```

### Execution Visualization

### Example: [-1, 0, 3, 5, 9, 12], target = 9
```
Initial: nums = [-1, 0, 3, 5, 9, 12], target = 9
left = 0, right = 5

Step 1: mid = 0 + (5-0)/2 = 2
→ nums[2] = 3, target = 9
→ 3 < 9, search right half
→ left = 3, right = 5
Range: [3, 4, 5] → [5, 9, 12]

Step 2: mid = 3 + (5-3)/2 = 4
→ nums[4] = 9, target = 9
→ Found! Return index 4

Final: Return 4 ✓
```

### Key Visualization Points:
- **Search range** halves each iteration
- **Mid calculation** prevents overflow
- **Comparison** determines which half to eliminate
- **Convergence** happens when target found or range exhausted

### Memory Layout Visualization:
```
Array: [-1, 0, 3, 5, 9, 12]
Index:   0   1   2   3   4   5

Step 1: left=0, right=5, mid=2
         [---|---]
         0   1   2   3   4   5
              ↑
              mid=2 (value=3 < target=9)
              Search right: [3,4,5]

Step 2: left=3, right=5, mid=4
         [-----|-]
         0   1   2   3   4   5
                      ↑
                      mid=4 (value=9 = target=9)
                      Found! ✓
```

### Time Complexity Breakdown:
- **Each Iteration**: Eliminates half of remaining elements
- **Number of Iterations**: O(log N) where N is array length
- **Total Time**: O(log N)
- **Space**: O(1) - only left, right, mid variables
- **Optimal**: Cannot do better for sorted array search
*/
