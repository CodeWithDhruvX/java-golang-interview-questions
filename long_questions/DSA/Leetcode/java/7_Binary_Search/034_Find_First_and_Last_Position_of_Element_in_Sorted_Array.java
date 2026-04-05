import java.util.Arrays;

public class FindFirstAndLastPositionOfElementInSortedArray {
    
    // 34. Find First and Last Position of Element in Sorted Array
    // Time: O(log N), Space: O(1)
    public static int[] searchRange(int[] nums, int target) {
        return new int[]{findFirstOccurrence(nums, target), findLastOccurrence(nums, target)};
    }

    private static int findFirstOccurrence(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        int result = -1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                result = mid;
                right = mid - 1; // Continue searching left half
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return result;
    }

    private static int findLastOccurrence(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        int result = -1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                result = mid;
                left = mid + 1; // Continue searching right half
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {5, 7, 7, 8, 8, 10},
            {5, 7, 7, 8, 8, 10},
            {},
            {1},
            {1},
            {2, 2, 2, 2, 2},
            {1, 2, 3, 4, 5},
            {1, 2, 3, 4, 5},
            {-3, -2, -1, 0, 1},
            {1, 3, 5, 7, 9}
        };
        
        int[] targets = {8, 6, 0, 1, 0, 2, 3, 6, -1, 4};
        
        for (int i = 0; i < testArrays.length; i++) {
            int[] result = searchRange(testArrays[i], targets[i]);
            System.out.printf("Test Case %d: nums=%s, target=%d -> Range: %s\n", 
                i + 1, Arrays.toString(testArrays[i]), targets[i], Arrays.toString(result));
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Modified Binary Search for Range Boundaries
- **Two Binary Searches**: One for first occurrence, one for last
- **Range Tracking**: Find boundaries of target value range
- **Left Search**: Find leftmost occurrence
- **Right Search**: Find rightmost occurrence

## 2. PROBLEM CHARACTERISTICS
- **Sorted Array**: Input array is sorted with possible duplicates
- **Target Range**: Need to find first and last position of target
- **Multiple Occurrences**: Target may appear multiple times
- **Index Range**: Return [firstIndex, lastIndex] or [-1, -1]

## 3. SIMILAR PROBLEMS
- Search Insert Position
- Find First Bad Version
- Search in a 2D Matrix
- Find Peak Element

## 4. KEY OBSERVATIONS
- Standard binary search finds any occurrence
- Need modified version to find boundaries
- First occurrence: continue searching left half when found
- Last occurrence: continue searching right half when found
- Can combine both searches in single pass

## 5. VARIATIONS & EXTENSIONS
- Find count of target elements
- Find range of all elements between two targets
- Handle circular sorted arrays
- Multiple range queries on same array

## 6. INTERVIEW INSIGHTS
- Clarify: "Can array contain duplicates?"
- Clarify: "Should I return 1-based or 0-based indices?"
- Edge cases: empty array, single element, target not found
- Time complexity: O(log N) vs O(N) linear scan

## 7. COMMON MISTAKES
- Using standard binary search (finds any occurrence)
- Not handling empty array case
- Off-by-one errors in index calculations
- Wrong loop termination condition

## 8. OPTIMIZATION STRATEGIES
- Current approach is optimal
- Can combine both searches in single pass
- For multiple queries, array stays sorted
- For very large arrays, consider cache-friendly access

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding all copies of a book in library:**
- You have books arranged by catalog number (sorted array)
- You want to find the first and last copy of a specific book
- You use binary search to quickly narrow down to the right section
- Once you find one copy, you search left and right to find boundaries

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Sorted array (with duplicates) and target value
2. **Goal**: Find first and last index where target appears
3. **Output**: Array [firstIndex, lastIndex] or [-1, -1]

#### Phase 2: Key Insight Recognition
- **"Standard binary search finds any occurrence!"** → Need modification
- **"How to find first?"** → Continue searching left when found
- **"How to find last?"** → Continue searching right when found
- **"Can optimize?"** → Two separate searches are clear and simple

#### Phase 3: Strategy Development
```
Human thought process:
"I'll do two binary searches:
1. First search: find leftmost occurrence
   - When I find target, continue searching left half
2. Second search: Find rightmost occurrence  
   - When I find target, continue searching right half

Both searches are independent and O(log N) each."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return [-1, -1]
- **Single element**: Return [0, 0] if matches, [-1, -1] otherwise
- **Target not found**: Both searches will return -1
- **All elements same as target**: Return [0, n-1]

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [5, 7, 7, 8, 8, 10], target = 8

Human thinking:
"First Search (find first occurrence):
left = 0, right = 5
mid = 2, nums[2] = 7 < 8, search right
left = 3, right = 5
mid = 4, nums[4] = 8 = target, found!
Continue searching left: right = 3
mid = 3, nums[3] = 8 = target, found!
Continue searching left: right = 2
mid = 2, nums[2] = 7 < 8, search right
left = 3, right = 2 → Stop
First occurrence at index 3

Second Search (find last occurrence):
left = 0, right = 5
mid = 2, nums[2] = 7 < 8, search right
left = 3, right = 5
mid = 4, nums[4] = 8 = target, found!
Continue searching right: left = 5
mid = 5, nums[5] = 10 > 8, search left
left = 5, right = 4 → Stop
Last occurrence at index 4

Result: [3, 4]"
```

#### Phase 6: Intuition Validation
- **Why it works**: Binary search guarantees logarithmic time
- **Why it's efficient**: Each search eliminates half the space
- **Why it's correct**: Systematic boundary finding ensures accuracy

### Common Human Pitfalls & How to Avoid Them
1. **"Why not linear scan?"** → Too slow O(N)
2. **"Can I do one search?"** → Complex and error-prone
3. **"What about equal elements?"** → Continue search in appropriate direction
4. **"How to handle not found?"** → Both searches return -1

### Real-World Analogy
**Like finding all red cars in a parking lot:**
- Cars are parked by license plate number (sorted array)
- You want to find the first and last red car (target)
- You quickly go to the middle section of the lot
- When you find a red car, you search left for earlier ones
- You also search right for later ones
- You end up with the first and last red cars

### Human-Readable Pseudocode
```
function searchRange(sortedArray, target):
    // Find first occurrence
    firstIndex = binarySearchModified(sortedArray, target, findFirst=true)
    
    // Find last occurrence
    lastIndex = binarySearchModified(sortedArray, target, findFirst=false)
    
    return [firstIndex, lastIndex]

function binarySearchModified(array, target, findFirst):
    left = 0, right = array.length - 1
    result = -1
    
    while left <= right:
        mid = left + (right - left) / 2
        
        if array[mid] == target:
            result = mid
            if findFirst:
                right = mid - 1  // Search left half
            else:
                left = mid + 1   // Search right half
        else if array[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return result
```

### Execution Visualization

### Example: [5, 7, 7, 8, 8, 10], target = 8
```
Initial: nums = [5, 7, 7, 8, 8, 10], target = 8

First Search (find first occurrence):
left = 0, right = 5, result = -1

Step 1: mid = 2, nums[2] = 7 < 8
→ left = 3, right = 5
Range: [8, 8, 10]

Step 2: mid = 4, nums[4] = 8 = target
→ result = 4, right = 3 (search left)
Range: [8, 8]

Step 3: mid = 3, nums[3] = 8 = target
→ result = 3, right = 2 (search left)
Range: [8]

Step 4: left = 3, right = 2 → Stop
First occurrence = 3

Second Search (find last occurrence):
left = 0, right = 5, result = -1

Step 1: mid = 2, nums[2] = 7 < 8
→ left = 3, right = 5
Range: [8, 8, 10]

Step 2: mid = 4, nums[4] = 8 = target
→ result = 4, left = 5 (search right)
Range: [10]

Step 3: mid = 5, nums[5] = 10 > 8
→ right = 4
Range: [] → Stop
Last occurrence = 4

Final: [3, 4] ✓
```

### Key Visualization Points:
- **Two independent searches** for boundaries
- **Modified binary search** continues in appropriate direction
- **First search** looks left when target found
- **Last search** looks right when target found

### Memory Layout Visualization:
```
Array: [5][7][7][8][8][10]
Index:  0  1  2  3  4  5
Target: 8

First Search:
Step 1: mid=2, nums[2]=7 < 8 → [5][7]|[7][8][8][10]
Step 2: mid=4, nums[4]=8 = target → [5][7]|[7]|[8][8][10]
         Found at 4, search left → [5][7]|[7]|[8]
Step 3: mid=3, nums[3]=8 = target → [5][7]|[7]|[8]
         Found at 3, search left → []
First: index 3 ✓

Second Search:
Step 1: mid=2, nums[2]=7 < 8 → [5][7]|[7][8][8][10]
Step 2: mid=4, nums[4]=8 = target → [5][7]|[7]|[8]|[8][10]
         Found at 4, search right → [10]
Step 3: mid=5, nums[5]=10 > 8 → [] → Stop
Last: index 4 ✓
```

### Time Complexity Breakdown:
- **Two Binary Searches**: 2 × O(log N) = O(log N)
- **Space Complexity**: O(1) - constant extra space
- **Optimal**: Cannot do better than O(log N) for sorted array
- **Alternative**: Single pass O(N) but slower for large arrays
*/
