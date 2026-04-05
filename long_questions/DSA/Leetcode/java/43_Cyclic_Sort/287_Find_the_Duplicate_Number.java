import java.util.*;

public class FindTheDuplicateNumber {
    
    // 287. Find the Duplicate Number
    // Time: O(N), Space: O(1) - Floyd's Tortoise and Hare Algorithm
    public static int findDuplicate(int[] nums) {
        // Phase 1: Find the intersection point
        int slow = nums[0];
        int fast = nums[0];
        
        while (true) {
            slow = nums[slow];
            fast = nums[nums[fast]];
            if (slow == fast) {
                break;
            }
        }
        
        // Phase 2: Find the entrance to the cycle
        slow = nums[0];
        while (slow != fast) {
            slow = nums[slow];
            fast = nums[fast];
        }
        
        return slow;
    }

    // Alternative solution using cyclic sort (modifies the array)
    public static int findDuplicateCyclicSort(int[] nums) {
        int i = 0;
        int n = nums.length;
        
        while (i < n) {
            int correctPos = nums[i] - 1;
            if (nums[i] != nums[correctPos]) {
                swap(nums, i, correctPos);
            } else {
                i++;
            }
        }
        
        // The duplicate will be at the position where the number doesn't match index+1
        for (i = 0; i < n; i++) {
            if (nums[i] != i + 1) {
                return nums[i];
            }
        }
        
        return -1; // Should never reach here for valid input
    }
    
    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    // Binary search approach
    public static int findDuplicateBinarySearch(int[] nums) {
        int left = 1;
        int right = nums.length - 1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            int count = 0;
            
            // Count numbers less than or equal to mid
            for (int num : nums) {
                if (num <= mid) {
                    count++;
                }
            }
            
            if (count > mid) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return left;
    }

    // Hash set approach (uses extra space)
    public static int findDuplicateHashSet(int[] nums) {
        Set<Integer> seen = new HashSet<>();
        
        for (int num : nums) {
            if (seen.contains(num)) {
                return num;
            }
            seen.add(num);
        }
        
        return -1; // Should never reach here for valid input
    }

    // Array marking approach (modifies array)
    public static int findDuplicateArrayMarking(int[] nums) {
        for (int num : nums) {
            int index = Math.abs(num);
            if (nums[index] < 0) {
                return index;
            }
            nums[index] = -nums[index];
        }
        
        return -1; // Should never reach here for valid input
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 3, 4, 2, 2},
            {3, 1, 3, 4, 2},
            {1, 1},
            {2, 2, 2, 2, 2},
            {1, 4, 4, 3, 2},
            {5, 4, 3, 2, 1, 5},
            {3, 1, 2, 3, 4, 5},
            {2, 5, 9, 6, 9, 3, 8, 9, 7, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            // Make copies for different approaches
            int[] nums1 = testCases[i].clone();
            int[] nums2 = testCases[i].clone();
            int[] nums3 = testCases[i].clone();
            int[] nums4 = testCases[i].clone();
            int[] nums5 = testCases[i].clone();
            
            int result1 = findDuplicate(nums1);
            int result2 = findDuplicateCyclicSort(nums2);
            int result3 = findDuplicateBinarySearch(nums3);
            int result4 = findDuplicateHashSet(nums4);
            int result5 = findDuplicateArrayMarking(nums5);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  Floyd's: %d\n", result1);
            System.out.printf("  Cyclic: %d\n", result2);
            System.out.printf("  Binary:  %d\n", result3);
            System.out.printf("  HashSet: %d\n", result4);
            System.out.printf("  Marking: %d\n\n", result5);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Cyclic Sort
- **In-Place Sorting**: Sort array without extra space
- **Cycle Detection**: Use array values as indices
- **Position Mapping**: Each number belongs to specific position
- **O(N) Time**: Linear time sorting algorithm

## 2. PROBLEM CHARACTERISTICS
- **Duplicate Detection**: Find duplicate in array with constraints
- **Range Constraints**: Numbers from 1 to N, one duplicate
- **In-Place Requirement**: Cannot use extra space
- **Multiple Solutions**: Different approaches with trade-offs

## 3. SIMILAR PROBLEMS
- Find All Missing Numbers
- First Missing Positive
- Find the Duplicate Number (general case)
- Sort Colors (Dutch National Flag)

## 4. KEY OBSERVATIONS
- Array values can be used as indices
- Each number should be at position value-1
- Duplicate creates a cycle in index mapping
- Time complexity: O(N) vs O(N log N) comparison sort
- Space complexity: O(1) vs O(N) for hash set

## 5. VARIATIONS & EXTENSIONS
- Multiple duplicates
- Missing numbers with duplicates
- Different range constraints
- Custom comparison functions

## 6. INTERVIEW INSIGHTS
- Clarify: "Can we modify the array?"
- Edge cases: single element, no duplicate, large arrays
- Time complexity: O(N) vs O(N log N) vs O(N) with hash
- Space complexity: O(1) vs O(N) for hash set

## 7. COMMON MISTAKES
- Incorrect index calculation (value-1 vs value)
- Infinite loop in swapping logic
- Not handling negative values properly
- Wrong termination conditions
- Missing edge case handling

## 8. OPTIMIZATION STRATEGIES
- Use while loop instead of for for better control
- Handle index bounds carefully
- Minimize array access operations
- Use appropriate swap operations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing numbered boxes:**
- You have boxes numbered 1 to N, each containing a number
- Each number should be in the box with that number-1
- One number is duplicated, creating a cycle
- Cyclic sort swaps boxes until each number is in correct box
- This is like a self-organizing system where items guide themselves

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array with numbers 1 to N, one duplicate
2. **Goal**: Find the duplicate number
3. **Output**: Duplicate number

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N log N) with comparison sort
- **"How to optimize?"** → Use array values as indices
- **"Why cyclic sort?"** → Each number knows its correct position
- **"How to handle cycles?"** → Detect and resolve cycles

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use cyclic sort:
1. For each position i, check if nums[i] is correct
2. If not, swap nums[i] with nums[nums[i]-1]
3. Continue until all numbers are in correct positions
4. The duplicate will be at position where cycle exists
5. Return the duplicate number"
```

#### Phase 4: Edge Case Handling
- **Single element**: Return the only element
- **No duplicate**: Handle gracefully (shouldn't happen per constraints)
- **Large arrays**: Ensure O(N) time complexity
- **Negative numbers**: Use Math.abs() for index calculation

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [1, 3, 4, 2, 2]

Human thinking:
"Let's apply cyclic sort:

Initialize: i=0

Position 0: nums[0]=1, correct position is 1-1=0
- nums[0] is already correct → i=1

Position 1: nums[1]=3, correct position is 3-1=2
- nums[1] != nums[2]=4 → swap nums[1] and nums[2]
- Array becomes: [1, 4, 3, 2, 2]

Position 1 again: nums[1]=4, correct position is 4-1=3
- nums[1] != nums[3]=2 → swap nums[1] and nums[3]
- Array becomes: [1, 2, 3, 4, 2]

Position 2: nums[2]=3, Correct position is 3-1=2
- nums[2] != nums[1]=2 → swap nums[2] and nums[1]
- Array becomes: [1, 3, 2, 4, 2]

Continue this process...
Eventually: [1, 2, 3, 4, 2]

Now check for duplicate:
nums[4]=2, correct position is 2-1=1
nums[1]=3, Correct position is 3-1=2
nums[2]=2, Correct position is 2-1=1

We have a cycle: 2→1→3→2
Duplicate found: 2 ✓

Manual verification:
Array contains two 2's, duplicate is 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Each number moves toward correct position
- **Why it's efficient**: O(N) time vs O(N log N) sorting
- **Why it's correct**: Cycles indicate duplicate location

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort and scan?"** → O(N log N) slower
2. **"What about hash set?"** → Uses O(N) extra space
3. **"How to handle indices?"** → Use value-1 for 0-based indexing
4. **"What about infinite loops?"** → Proper termination conditions

### Real-World Analogy
**Like organizing books by call numbers:**
- You have books with call numbers 1 to N
- Each book should be on shelf with that call number-1
- One call number appears on two books (duplicate)
- Cyclic sort: each book moves to its correct shelf
- Books help organize themselves by following call numbers
- When a book points to already occupied shelf, we found duplicate
- Useful in memory management, cache organization, data deduplication
- Like a self-organizing library system

### Human-Readable Pseudocode
```
function cyclicSort(nums):
    i = 0
    
    while i < nums.length:
        correctPos = nums[i] - 1
        
        if nums[i] != nums[correctPos]:
            swap(nums, i, correctPos)
        else:
            i++
    
    // Find duplicate
    for i from 0 to nums.length-1:
        if nums[i] != i + 1:
            return nums[i]
    
    return -1

function swap(nums, i, j):
    temp = nums[i]
    nums[i] = nums[j]
    nums[j] = temp
```

### Execution Visualization

### Example: nums = [1, 3, 4, 2, 2]
```
Cyclic Sort Process:

Initial: [1, 3, 4, 2, 2]
i=0: nums[0]=1, correctPos=0 → already correct, i=1

i=1: nums[1]=3, correctPos=2
nums[1]=3 != nums[2]=4 → swap(1,2)
Array: [1, 4, 3, 2, 2]

i=1: nums[1]=4, correctPos=3
nums[1]=4 != nums[3]=2 → swap(1,3)
Array: [1, 2, 3, 4, 2]

i=1: nums[1]=2, correctPos=1
nums[1]=2 != nums[0]=1 → swap(1,0)
Array: [2, 1, 3, 4, 2]

Continue process...

Final sorted: [1, 2, 3, 4, 2]

Duplicate detection:
nums[4]=2, correctPos=1 → nums[4]=2 == nums[1]=1 ✓
Found duplicate: 2 ✓

Visualization:
Numbers organize themselves by following value-1 indices
Cycle detection reveals the duplicate
Linear time complexity achieved
```

### Key Visualization Points:
- **Index Mapping**: Value maps to position (value-1)
- **Self-Organization**: Numbers move to correct positions
- **Cycle Detection**: Duplicate creates cycle in index mapping
- **In-Place**: No extra space required

### Memory Layout Visualization:
```
Initial: [1, 3, 4, 2, 2]
Index mapping: 1→0, 3→2, 4→3, 2→1, 2→4

After sorting: [1, 2, 3, 4, 2]
Index mapping: 1→0, 2→1, 3→2, 4→3, 2→4

Cycle: 2→1→3→2 (duplicate at 2)

Cyclic sort uses array values as indices
Self-organizing property enables O(N) sorting
Duplicate detected through cycle analysis
```

### Time Complexity Breakdown:
- **Cyclic Sort**: O(N) time, O(1) space
- **Floyd's Algorithm**: O(N) time, O(1) space
- **Binary Search**: O(N log N) time, O(1) space
- **Hash Set**: O(N) time, O(N) space
- **Optimal**: O(N) time is best possible
- **Trade-offs**: Time vs space vs array modification
*/
