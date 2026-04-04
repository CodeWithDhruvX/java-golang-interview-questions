package main

import "fmt"

// 33. Search in Rotated Sorted Array
// Time: O(log N), Space: O(1)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			return mid
		}
		
		// Determine which side is sorted
		if nums[left] <= nums[mid] {
			// Left side is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right side is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	
	return -1
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Modified Binary Search for Rotated Array
- **Rotation Point**: Array is rotated at some pivot
- **Sorted Halves**: At least one half is always sorted
- **Target Comparison**: Compare target with sorted half boundaries
- **Range Reduction**: Eliminate half based on target location

## 2. PROBLEM CHARACTERISTICS
- **Rotated Sorted Array**: Originally sorted, then rotated
- **No Duplicates**: All elements are unique
- **Single Search**: Find one occurrence of target
- **Logarithmic Time**: O(log N) solution required

## 3. SIMILAR PROBLEMS
- Find Minimum in Rotated Sorted Array (LeetCode 153)
- Search in Rotated Sorted Array II (with duplicates, LeetCode 81)
- Find Rotation Count
- Search Insert Position in Rotated Array

## 4. KEY OBSERVATIONS
- **One sorted half**: In any rotated array, at least one half is sorted
- **Rotation detection**: Compare left and mid to identify sorted half
- **Target location**: Target must be in sorted half if within range
- **Pivot property**: Pivot is where array "wraps around"

## 5. VARIATIONS & EXTENSIONS
- **With Duplicates**: More complex, may need O(N) in worst case
- **Multiple Rotations**: Handle multiple rotation points
- **Find Rotation Point**: Find index where rotation occurred
- **Search Range**: Find all elements in range

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are there duplicates? How many rotations?"
- Edge cases: empty array, single element, no rotation
- Space complexity: O(1) - only pointers and variables
- Time complexity: O(log N) - modified binary search

## 7. COMMON MISTAKES
- Not correctly identifying which half is sorted
- Wrong boundary conditions for target comparison
- Using standard binary search without rotation handling
- Off-by-one errors in range updates
- Not handling single element arrays

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(log N) time, O(1) space
- **Find pivot first**: Can find rotation point, then normal binary search
- **Recursive version**: Can be implemented recursively
- **Template method**: Use consistent modified binary search pattern

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a page in a book that's been split and reordered:**
- You have a book that was cut at some page and the pieces swapped
- The first part now comes after the second part (rotated)
- Each half is still in correct order within itself
- When you look at the middle, one side will be in proper order
- You can tell which side is sorted by comparing the endpoints
- Search in the sorted half if your target could be there

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Rotated sorted array and target value
2. **Goal**: Find index of target in rotated array
3. **Output**: Index if found, -1 if not found
4. **Constraint**: Array was originally sorted, then rotated

#### Phase 2: Key Insight Recognition
- **"Rotation property"** → One half is always sorted
- **"Sorted half identification"** → Compare left, mid, right values
- **"Target range checking"** → Target must be within sorted half range
- **"Modified binary search"** → Adapt standard binary search logic

#### Phase 3: Strategy Development
```
Human thought process:
"I need to search in a rotated sorted array.
At any point, one half of the array will be sorted.
I can identify which half is sorted by comparing left and mid.
If the left half is sorted and my target is in that range, search there.
Otherwise, search the other half.
This way, I can still eliminate half the elements each time."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return -1 (no elements to search)
- **Single element**: Check if it equals target
- **No rotation**: Standard binary search applies
- **Target not found**: Return -1 when search range exhausted

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums=[4,5,6,7,0,1,2], target=0

Human thinking:
"I'll search for 0 in the rotated array:

Initial: left=0, right=6
mid=3, nums[mid]=7
Check if left half is sorted: nums[0]=4 <= nums[3]=7 ✓
Left half [4,5,6,7] is sorted
Is target (0) in sorted half range [4,7]? No
Search right half: left=4, right=6

Step 2: left=4, right=6
mid=5, nums[mid]=1
Check if left half is sorted: nums[4]=0 <= nums[5]=1 ✓
Left half [0,1] is sorted
Is target (0) in sorted half range [0,1]? Yes!
Search left half: right=4

Step 3: left=4, right=4
mid=4, nums[mid]=0
Found target! Return index 4"
```

#### Phase 6: Intuition Validation
- **Why one half is sorted**: Rotation only affects one boundary
- **Why O(log N) time**: Still eliminate half elements each iteration
- **Why range checking works**: Sorted half allows range comparison
- **Why modification needed**: Standard binary search assumes full array sorted

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just standard binary search?"** → Array is not fully sorted
2. **"How do I know which half is sorted?"** → Compare left and mid values
3. **"What if both halves seem sorted?"** → Can't happen with unique elements
4. **"Can I find the rotation point first?"** → Yes, but this approach is more direct

### Real-World Analogy
**Like finding a chapter in a book where chapters were reordered:**
- You have a book where someone moved the first few chapters to the end
- Within each section, chapters are still in order
- When you open to a random page, you can tell if you're in the "original" or "moved" section
- If your target chapter number fits in the current section's range, search there
- Otherwise, search in the other section

### Human-Readable Pseudocode
```
function searchRotated(array, target):
    left = 0
    right = length(array) - 1
    
    while left <= right:
        middle = left + (right - left) // 2
        
        if array[middle] == target:
            return middle
        
        # Check if left half is sorted
        if array[left] <= array[middle]:
            # Left half is sorted
            if array[left] <= target < array[middle]:
                right = middle - 1  # Search left half
            else:
                left = middle + 1   # Search right half
        else:
            # Right half is sorted
            if array[middle] < target <= array[right]:
                left = middle + 1   # Search right half
            else:
                right = middle - 1  # Search left half
    
    return -1  # Target not found
```

### Execution Visualization

### Example: nums=[4,5,6,7,0,1,2], target=0
```
Array: [4][5][6][7][0][1][2]
Index:   0  1  2  3  4  5  6

Initial: left=0, right=6
mid=3, nums[mid]=7
nums[left]=4 <= nums[mid]=7 → left half sorted
target=0 not in [4,7] → search right half
left=4, right=6

Step 2: left=4, right=6
mid=5, nums[mid]=1
nums[left]=0 <= nums[mid]=1 → left half sorted
target=0 in [0,1] → search left half
right=4

Step 3: left=4, right=4
mid=4, nums[mid]=0
FOUND! Return 4
```

### Example: Target not found: nums=[4,5,6,7,0,1,2], target=3
```
Initial: left=0, right=6
mid=3, nums[mid]=7
left half sorted, target=3 not in [4,7] → search right half
left=4, right=6

Step 2: left=4, right=6
mid=5, nums[mid]=1
left half sorted, target=3 not in [0,1] → search right half
left=6, right=6

Step 3: left=6, right=6
mid=6, nums[mid]=2
right half sorted, target=3 not in [2,2] → search left half
right=5

left=6 > right=5 → NOT FOUND, return -1
```

### Key Visualization Points:
- **Sorted half detection**: Compare nums[left] and nums[mid]
- **Range checking**: Target must be within sorted half bounds
- **Range reduction**: Still eliminate half elements each iteration
- **Rotation handling**: Adapt search based on which half is sorted

### Memory Layout Visualization:
```
Array: [4][5][6][7][0][1][2]
Index:   0  1  2  3  4  5  6
        ^           ^
     left=0      right=6
        mid=3 (value=7)
        left half [0-3] is sorted: [4,5,6,7]
```

### Time Complexity Breakdown:
- **Each iteration**: O(1) - constant work
- **Number of iterations**: O(log N) - halving search space
- **Total time**: O(log N)
- **Space**: O(1) - only pointers and variables

### Alternative Approaches:

#### 1. Find Pivot First (O(log N) time, O(1) space)
```go
func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
    // Find rotation point
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    // Determine which half to search
    rotation := left
    if target >= nums[rotation] && target <= nums[len(nums)-1] {
        // Search right half
        return binarySearch(nums, rotation, len(nums)-1, target)
    } else {
        // Search left half
        return binarySearch(nums, 0, rotation-1, target)
    }
}

func binarySearch(nums []int, left, right, target int) int {
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```
- **Pros**: Clear separation of concerns
- **Cons**: Two binary searches instead of one

#### 2. Linear Search (O(N) time, O(1) space)
```go
func search(nums []int, target int) int {
    for i, num := range nums {
        if num == target {
            return i
        }
    }
    return -1
}
```
- **Pros**: Simple, works for any array
- **Cons**: O(N) time, doesn't leverage sorted property

#### 3. Standard Library (Not directly applicable)
- Standard library binary search assumes fully sorted array
- Would need to find rotation point first

### Extensions for Interviews:
- **With Duplicates**: Handle arrays with duplicate elements
- **Find Rotation Count**: Return index where rotation occurred
- **Search Range**: Find all elements within a range
- **Multiple Targets**: Find all occurrences of target
*/
func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
		{[]int{1}, 0},
		{[]int{1}, 1},
		{[]int{5, 1, 3}, 5},
		{[]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 9},
		{[]int{2, 3, 4, 5, 6, 7, 8, 9, 1}, 1},
		{[]int{6, 7, 0, 1, 2, 3, 4, 5}, 0},
		{[]int{6, 7, 0, 1, 2, 3, 4, 5}, 5},
		{[]int{1, 3, 5}, 2},
	}
	
	for i, tc := range testCases {
		result := search(tc.nums, tc.target)
		fmt.Printf("Test Case %d: nums=%v, target=%d -> Index: %d\n", 
			i+1, tc.nums, tc.target, result)
	}
}
